# DISCOVERY-LEVEL-SECURITY-DEFINITION.md
## Fase 3 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | DISCOVERY-LEVEL-SECURITY-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | 02 de agosto de 2026 |
| **Autor** | Security Architect / IAM Specialist |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 02/08/2026 |

**Documentos Vinculados:**
- [`GLOBAL-SECURITY.md`](../../../.specs/security/GLOBAL-SECURITY.md) — Política de Segurança Global (baseline corporativa)
- [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md) — PRD Discovery-Level (F1)
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)
- [`STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`](../../../.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md) — Padrões Corporativos

---

## 1. Postura de Segurança — Visão Macro

### 1.1 Modelo de Defesa em Camadas

A FBSO Platform adota um modelo de **defense-in-depth** com 5 camadas de proteção, da borda ao dado:

```
┌─────────────────────────────────────────────────────────────┐
│ CAMADA 1: Edge/Network — Cloudflare                          │
│ WAF (OWASP Top 10), DDoS Protection, SSL/TLS Termination,   │
│ IP Reputation, Geo-blocking (opcional), Bot Management       │
├─────────────────────────────────────────────────────────────┤
│ CAMADA 2: API Gateway — Kong                                 │
│ JWT Validation (Service-ID/Token-ID via Keycloak),          │
│ Rate Limiting (por tenant, por IP, por endpoint),            │
│ Request Validation, Header Sanitization, CORS enforcement    │
├─────────────────────────────────────────────────────────────┤
│ CAMADA 3: Service Mesh — Istio                               │
│ mTLS entre todos os serviços, Istio AuthorizationPolicy,     │
│ Traffic control (canary, blue-green), Circuit breaking       │
├─────────────────────────────────────────────────────────────┤
│ CAMADA 4: Application — Spring Boot + RBAC                   │
│ Role-based access control (Admin/Gerente/Operador),          │
│ Input validation (Bean Validation), Parameterized queries,   │
│ Audit trail (toda ação administrativa), Rate limiting local  │
├─────────────────────────────────────────────────────────────┤
│ CAMADA 5: Data — PostgreSQL RLS + Criptografia               │
│ Row-Level Security (tenant_id + bu_id), Soft Delete,         │
│ TLS in transit, Encryption at rest (DO Managed),            │
│ Secrets externalizados (variáveis de ambiente / Vault)       │
└─────────────────────────────────────────────────────────────┘
```

### 1.2 Trust Boundary — Kong Gateway

A **trust boundary** primária do sistema é o Kong API Gateway. O racional é:
- Nenhum serviço backend recebe tráfego que não tenha passado pelo Kong
- Kong valida o JWT via Keycloak (protocolo Service-ID/Token-ID) **antes** de encaminhar qualquer requisição
- Kong injeta headers com identidade do usuário e tenant context. O backend **confia** nesses headers — não revalida JWT
- A rede Kubernetes é configurada para que **apenas Kong** tenha acesso externo aos pods do backend (NetworkPolicy)

**Regra de Ouro:** Bypass do Kong é proibido. Qualquer comunicação direta com o backend é bloqueada no nível de rede.

---

## 2. Threat Model — Análise de Ameaças (STRIDE Discovery-Level)

### 2.1 Ameaças Priorizadas

| ID | Ameaça | Categoria STRIDE | Cenário de Ataque | Severidade | Mitigação |
|----|--------|-----------------|-------------------|------------|-----------|
| T1 | **Cross-Tenant Data Access** | Information Disclosure | Usuário do Tenant A consegue visualizar dados do Tenant B por falha no RLS ou bug no tenant context | 🔴 Crítica | RLS enforced no PostgreSQL + tenant_id em toda query + teste de isolamento automatizado |
| T2 | **Intra-Tenant Cross-BU Access** | Information Disclosure | Operador da Filial SP acessa dados fiscais da Filial RJ do mesmo tenant por falha no escopo RBAC | 🔴 Crítica | Escopo de `bu_id` no JWT + verificação no backend + RLS com `tenant_id AND bu_id` |
| T3 | **Privilege Escalation** | Elevation of Privilege | Operador altera papel para Admin via manipulação de payload ou endpoint sem verificação | 🔴 Crítica | RBAC enforced no backend + papéis validados em cada request + alteração de papel restrita a Admin |
| T4 | **Kong Bypass** | Tampering | Atacante descobre IP interno do backend e envia requisição direta, sem passar pelo Kong | 🟠 Alta | NetworkPolicy K8s: apenas Kong → Backend; mTLS Istio entre todos os serviços |
| T5 | **JWT Token Tampering** | Spoofing | Atacante forja JWT com claims adulterados (ex: `role: admin`, `tenant_id: outro`) | 🟠 Alta | Kong valida assinatura JWT via Keycloak antes de qualquer processamento; expiração curta (15 min) + refresh |
| T6 | **Audit Trail Tampering** | Repudiation | Administrador mal-intencionado altera ou remove registros de auditoria para encobrir ação | 🟡 Média | Audit trail em tabela imutável (apenas INSERT, sem UPDATE/DELETE) + replicação para Elastic Stack |
| T7 | **Brute Force / Credential Stuffing** | Spoofing | Atacante tenta múltiplos logins via força bruta no endpoint de autenticação | 🟡 Média | Kong rate limiting (por IP e por tenant) + Cloudflare Bot Management + Keycloak brute-force protection |
| T8 | **SQL Injection via API** | Tampering | Atacante injeta SQL via parâmetros de API mal-validados | 🟡 Média | JPA/Hibernate parameterized queries (nunca concatenação); Bean Validation em todos os inputs |
| T9 | **XSS via Portal do Cliente** | Tampering | Atacante injeta script via campos de texto (ex: nome de produto, razão social) | 🟡 Média | React com sanitização automática (escaping JSX); Content-Security-Policy header; input validation |
| T10 | **DDoS no Onboarding** | Denial of Service | Atacante dispara criação massiva de contas via endpoint de registro para esgotar recursos | 🟢 Baixa | Cloudflare DDoS Protection + Kong rate limiting + CAPTCHA no fluxo de registro (fase futura) |

### 2.2 Trust Zones

| Zona | Componentes | Nível de Confiança | Controles |
|------|------------|-------------------|-----------|
| **Zona 0 — Internet** | Usuários finais, navegadores | ☠️ Zero trust | Todo input é hostil; Cloudflare é primeira barreira |
| **Zona 1 — Edge** | Cloudflare | 🔒 Confiável (validado) | WAF, DDoS, SSL termination |
| **Zona 2 — DMZ** | Kong Gateway | 🔒 Trust boundary | Único ponto de entrada; valida JWT + rate limiting |
| **Zona 3 — Serviços** | Backend, Keycloak | 🔒 Confiável (mTLS) | Istio mTLS; comunicação intra-cluster apenas |
| **Zona 4 — Dados** | PostgreSQL, Redis | 🔒 Alta confiança | Nunca exposto externamente; TLS + RLS + criptografia |

---

## 3. Compliance Regulatório

### 3.1 LGPD — Lei Geral de Proteção de Dados

| Artigo LGPD | Implicação para o Projeto | Ação Requerida |
|-------------|--------------------------|----------------|
| **Art. 7º — Bases legais** | Dados cadastrais de clientes (razão social, CNPJ, e-mail) precisam de base legal | Consentimento/contrato como base legal; registro da finalidade no momento da coleta |
| **Art. 18º — Direitos do titular** | Cliente pode solicitar acesso, correção, exclusão de seus dados | Implementar endpoint de exportação de dados; Soft Delete permite exclusão lógica com rastreabilidade |
| **Art. 46º — Segurança** | Medidas técnicas para proteger dados pessoais | Criptografia em repouso (DO Managed) + em trânsito (TLS 1.3) + RLS + audit trail |
| **Art. 48º — Comunicação de incidentes** | Notificar titulares e ANPD em caso de vazamento | Plano de resposta a incidentes; Elastic Stack para detecção de acessos anômalos |

### 3.2 Frameworks de Segurança Aplicáveis

| Framework | Aplicação | Status |
|-----------|-----------|--------|
| **OWASP Top 10 (2021)** | WAF Cloudflare + práticas de desenvolvimento seguro | ✅ Baseline coberto |
| **ASVS Level 2** | Autenticação, autorização, validação, criptografia | 🎯 Alvo para fase Core |
| **NIST CSF** | Identify, Protect, Detect, Respond, Recover | 📋 Planejado para maturidade |
| **PCI-DSS** | Dados de pagamento — fora do escopo atual (sem processamento financeiro) | 🔮 Futuro (fase de comercialização) |
| **SOC 2** | Controles de segurança para clientes enterprise | 🔮 Futuro (pós-Core, para habilitação de vendas enterprise) |

---

## 4. Estratégia de Segurança por Componente

### 4.1 Cloudflare — Camada de Borda

| Controle | Configuração Discovery-Level |
|----------|------------------------------|
| **WAF Rules** | OWASP Top 10 + regras customizadas para API (block SQLi, XSS, path traversal) |
| **Rate Limiting** | 100 req/min por IP para endpoints públicos; 300 req/min para autenticados |
| **DDoS Protection** | Modo "High" (detecção automática de camada 3/4 e 7) |
| **SSL/TLS** | Full (strict) — Cloudflare ↔ Kong com certificado válido |
| **Bot Management** | Ativar para endpoints de registro/login/onboarding |
| **Security Headers** | HSTS (max-age=1 ano), CSP, X-Frame-Options: DENY, X-Content-Type-Options: nosniff |

### 4.2 Kong API Gateway — Trust Boundary

| Controle | Configuração Discovery-Level |
|----------|------------------------------|
| **JWT Validation** | Plugin JWT — valida assinatura, exp, issuer contra Keycloak (Service-ID/Token-ID) |
| **Rate Limiting** | 30 req/s por tenant; 100 req/s por consumer; burst + quota |
| **CORS** | Whitelist domínios FBSO.ORG apenas (sem wildcard `*` em produção) |
| **Header Sanitization** | Remover headers internos (`X-Tenant-Id`, `X-User-Id`) de requests externos; apenas Kong injeta |
| **IP Restriction** | Painel admin acessível apenas de IPs internos FBSO (via VPN ou whitelist) |
| **Request Size** | Limitar payload a 10 MB por request |

### 4.3 Keycloak — Identity & Access Management

| Controle | Configuração Discovery-Level |
|----------|------------------------------|
| **Realms** | Um realm por tenant (provisioning automático na ativação da conta) |
| **JWT Claims** | `sub`, `tenant_id`, `roles`, `bu_ids`, `modules` — embutidos no JWT |
| **Token Expiry** | Access Token: 15 min; Refresh Token: 8h (rotação automática) |
| **Brute Force** | Keycloak brute-force detection: 5 falhas → bloqueio de 15 min |
| **Password Policy** | Mínimo 12 caracteres, complexidade (maiúscula, minúscula, número, símbolo), histórico de 5 senhas |
| **SAML 2.0** | Configurado para clientes enterprise que exigem SSO corporativo legado (justificativa documentada) |

### 4.4 Backend — ms-fbso-platform-admin

| Controle | Implementação |
|----------|---------------|
| **RBAC Enforcement** | `@PreAuthorize("hasRole('ADMIN')")` em endpoints sensíveis; validação de `bu_id` em toda query |
| **Input Validation** | Bean Validation (`@Valid`, `@NotNull`, `@Size`) + sanitização contra XSS |
| **SQL Protection** | JPA/Hibernate parameterized queries 100% (proibido concatenação de SQL) |
| **Error Handling** | Mensagens genéricas para o cliente; stack traces apenas em logs internos |
| **Audit Trail** | `@Auditable` annotation cross-cutting: toda ação admin registrada com user_id, ação, timestamp, IP |
| **Secrets** | Nenhum secret hardcoded; variáveis de ambiente injetadas via K8s Secrets ou Vault |
| **Log Sanitization** | Nunca logar senhas, tokens, CPFs, ou dados de cartão em logs de aplicação |

### 4.5 PostgreSQL — Camada de Dados

| Controle | Implementação |
|----------|---------------|
| **RLS Multi-Tenant** | Policies por tabela: `tenant_id = current_setting('app.current_tenant_id')` |
| **RLS Intra-Tenant** | Policies com escopo de `bu_id` para tabelas de dados de negócio |
| **TLS** | Conexão backend → PostgreSQL sempre com TLS 1.3 (DO Managed force TLS) |
| **Encryption at Rest** | Ativado por padrão no DO Managed PostgreSQL |
| **Soft Delete** | `deleted_at` + `deleted_by` em todas as tabelas de negócio; sem DELETE físico |
| **Backup** | Backup automático DO: diário com retenção de 7 dias; point-in-time recovery (PITR) |

### 4.6 Rede e Service Mesh

| Controle | Implementação |
|----------|---------------|
| **NetworkPolicy** | Apenas Kong → Backend (porta 8080); Backend → PostgreSQL (5432); Backend → Redis (6379); negar todo resto |
| **Istio mTLS** | Modo STRICT em produção; PERMISSIVE em dev/staging |
| **Istio AuthorizationPolicy** | Allow apenas de namespaces autorizados; deny de namespaces externos |
| **Egress** | Backend não inicia conexões externas (sem egress para internet) |

---

## 5. Checklist SDD — Aplicação ao Projeto

O checklist do `GLOBAL-SECURITY.md` é aplicado ao contexto deste projeto:

### 5.1 Autenticação e Autorização

| Item | Aplicação no Projeto |
|------|---------------------|
| Middleware de sessão | ✅ Kong valida JWT em toda rota antes de encaminhar ao backend |
| RBAC | ✅ `@PreAuthorize` em endpoints; 3 papéis MVP: Admin Tenant, Gerente BU, Operador |
| IDOR | ✅ `bu_id` validado no escopo do usuário logado; query sempre filtrada por `tenant_id` |

### 5.2 Proteção de Dados e Privacidade

| Item | Aplicação no Projeto |
|------|---------------------|
| Criptografia em repouso | ✅ DO Managed PostgreSQL + Redis com criptografia ativada |
| Sanitização de logs | ✅ Configuração de logback para mascarar campos sensíveis (senha, token, CPF) |
| Campos mascarados | ✅ API retorna apenas dados autorizados pelo escopo do usuário |

### 5.3 Validação e Sanitização de Entradas

| Item | Aplicação no Projeto |
|------|---------------------|
| Schema Validation | ✅ Bean Validation em todos os DTOs de entrada |
| SQL Injection | ✅ JPA/Hibernate — sem queries nativas com concatenação |
| XSS | ✅ React com escaping automático + CSP header + input sanitization |

### 5.4 Proteção de Infraestrutura e API

| Item | Aplicação no Projeto |
|------|---------------------|
| Rate Limiting | ✅ Cloudflare + Kong (duas camadas) |
| CORS | ✅ Kong configura whitelist de origens FBSO.ORG |
| Tratamento de erros seguro | ✅ Mensagens genéricas; stack traces apenas em logs internos |

---

## 6. Riscos de Segurança e Estimativa de Esforço

### 6.1 Riscos Específicos de Segurança

| ID | Risco | Prob. | Impacto | Mitigação |
|----|-------|-------|---------|-----------|
| RS1 | Configuração incorreta de RLS — tenant acessa dados de outro tenant | Baixa | 🔴 Crítico | Teste automatizado de isolamento: tentar acessar dados de outro tenant e verificar bloqueio |
| RS2 | Kong mal configurado permite bypass do gateway | Baixa | 🔴 Crítico | NetworkPolicy K8s como segunda camada; teste de penetração validando isolamento |
| RS3 | Segredos (API keys, DB password) expostos em repositório | Média | 🟠 Alto | Gitleaks no CI/CD; Vault para secrets em produção; `.env` no `.gitignore` |
| RS4 | Token JWT com claims excessivas expõe estrutura de permissões | Baixa | 🟡 Médio | Claims mínimas no JWT; permissões detalhadas consultadas no backend via cache |

### 6.2 Estimativa de Esforço de Segurança (Discovery-Level)

| Atividade | Complexidade | Esforço (dias) | Responsável |
|-----------|-------------|----------------|-------------|
| Configuração Kong↔Keycloak Service-ID/Token-ID | Complexa | 2 | Daniel Bruno Castro |
| Definição e implementação de RLS policies | Moderada | 1.5 | William Alves |
| Configuração Cloudflare WAF + Security Headers | Moderada | 1 | Lucas Silva Neto |
| Configuração Istio mTLS + NetworkPolicy | Moderada | 1.5 | Lucas Silva Neto |
| Integração Gitleaks + Semgrep no CI/CD | Leve | 0.5 | Lucas Silva Neto |
| Configuração Keycloak (realms, password policy, brute force) | Moderada | 1 | Daniel Bruno Castro |
| Threat model + documentação | Moderada | 1.5 | Daniel Bruno Castro |
| **Total Security** | — | **~9 dias** | — |

---

## 7. Ferramentas de Verificação Automatizada

| Ferramenta | Pipeline Stage | Verifica |
|------------|---------------|----------|
| **Gitleaks** | Pre-commit + CI (GitHub Actions) | Secrets hardcoded no código |
| **Semgrep** | CI (GitHub Actions) | SAST — vulnerabilidades de código (OWASP, CWE) |
| **npm audit** | CI (GitHub Actions) | Dependências frontend com CVEs conhecidas |
| **Maven OWASP plugin** | CI (GitHub Actions) | Dependências Java com CVEs conhecidas |
| **Terrascan** | CI (GitHub Actions) | IaC security — Terraform com vulnerabilidades |
| **kube-bench** | Deploy | Conformidade CIS para Kubernetes |
| **Penetration Test** | Pré-lançamento (manual) | Teste de penetração completo antes do primeiro deploy em produção |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: Security Definition Discovery-Level. 5 camadas de defesa, 10 ameaças STRIDE, compliance LGPD, estratégia por componente, checklist SDD aplicado, 4 riscos, estimativa de esforço de segurança | Security Architect / IAM Specialist |

---

🤖 *Upstream Architecture Discovery — Fase 3. Documento gerado pelo Security Architect como parte do Bloco B.*
