# SECURITY.md — Plano de Segurança da Solução: ms-fbso-platform-admin

- **Microserviço:** `ms-fbso-platform-admin`
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Caffeine Cache + Spring Security + Keycloak
- **Projeto de Negócio:** [PRJ-FIN-2026-0003-SAAS-FBSO-ORG](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/)
- **Versão:** 1.1
- **Data:** 21 de Julho de 2026
- **Situação implementação:** Em Execução
- **Status:** [STATUS: COMPLIANCE] — Validado via GATE-SECURITY-SCOPE em 21/07/2026. 5 dimensões validadas (1 APROVADO, 4 RESSALVAS). 5 NCs corrigidas.
- **Origem:** [PRD.md](./PRD.md) + [ARCHITECTURE.md](./ARCHITECTURE.md) + [GLOBAL-SECURITY.md](../../../../../../.specs/security/GLOBAL-SECURITY.md)
- **Referência Normativa:** [GLOBAL-SECURITY.md](../../../../../../.specs/security/GLOBAL-SECURITY.md) — Regras de Ouro e Checklist SDD

---

## 1. Visão Geral de Segurança

### 1.1 Contexto da Solução

O `ms-fbso-platform-admin` é o **backend administrativo SaaS multi-tenant** da plataforma FBSO.ORG. Ele gerencia dados sensíveis de múltiplos tenants (clientes corporativos), incluindo cadastros de empresas (CNPJ), usuários, permissões RBAC, unidades de negócio e catálogo de produtos.

### 1.2 Nível de Criticidade

**ALTO** — O sistema:
- Armazena dados corporativos de múltiplos clientes (multi-tenant)
- Gerencia identidades e permissões de usuários
- Contém trilha de auditoria com requisitos fiscais
- É a fundação sobre a qual módulos futuros (Tributali-Engine, Storekeeper) serão acoplados

### 1.3 Resumo dos Controles

| Camada | Controle |
|:---|:---|
| **Autenticação** | JWT (Keycloak) — RS256, validação de issuer, expiração |
| **Autorização** | RBAC DB-backed com `@RequiresPermission` (AOP) — 4 papéis |
| **Isolamento Multi-Tenant** | 3 camadas: PostgreSQL RLS (FORCE) + BaseRepository (tenant_id) + Teste de isolamento |
| **Auditoria** | `@Auditable` (AOP assíncrono) — 100% ações administrativas |
| **Proteção de Dados** | TLS 1.3, bcrypt (Keycloak), soft delete, mascaramento de PII em logs |
| **Segurança de API** | Rate limiting, CORS restrito, Bean Validation, RFC 7807 |
| **Pipeline DevSecOps** | SAST (Semgrep), Secret Scanning, Dependency Scanning (Dependabot) |

### 1.4 Pipeline de Segurança — Fluxo Unificado

O fluxo completo de segurança por requisição, alinhado com [ARCHITECTURE.md §4](./ARCHITECTURE.md):

```
1. JwtAuthenticationFilter  →  Valida JWT RS256 (Keycloak), extrai claims (tenant_id, user_id, roles),
                                seta TenantContext + app.current_tenant_id na sessão PostgreSQL
                                ❌ Falha → 401 Unauthorized
2. RbacAspect               →  Lê @RequiresPermission, consulta PermissionService (DB-backed),
                                valida contra matriz RN10-01 (resource_action + role_resource)
                                ❌ Falha → 403 Forbidden
3. Controller               →  Valida DTO (@Valid + Bean Validation), converte DTO → entity
                                ❌ Falha → 400 Bad Request
4. Service                  →  Executa lógica de negócio, valida regras (47 RNs — PRD §6.6)
                                ❌ Falha → 422 Unprocessable Entity
5. BaseRepository           →  Adiciona AND tenant_id = ? + AND deleted_dt IS NULL (automático)
6. PostgreSQL RLS (FORCE)   →  POLÍTICA tenant_isolation: USING (tenant_id = current_setting(...))
                                Bloqueia queries sem filtro — nível de banco, impossível burlar
7. AuditAspect              →  Intercepta @Auditable, captura snapshot before/after, grava em
                                audit_log (ASSÍNCRONO — não bloqueia a resposta)
8. ✅ HTTP Response         →  JSON + RFC 7807 (erros) + mensagens PT-BR (i18n)
```

> 💡 Esta pipeline está alinhada com o [ARCHITECTURE.md §4](./ARCHITECTURE.md). Cada etapa é descrita em detalhes nas seções abaixo.

---

## 2. Threat Model (STRIDE)

### 2.1 Atores e Agentes de Ameaça

| Ator | Motivação | Capacidade |
|:---|:---|:---|
| **Atacante externo não autenticado** | Acesso não autorizado, data breach | Força bruta em endpoints, exploração de vulnerabilidades |
| **Tenant malicioso** | Acessar dados de outro tenant | Envio de `tenant_id` arbitrário em requisições |
| **Usuário interno com privilégios elevados** | Abuso de permissões | Acesso a dados além do necessário |
| **Operador/Admin legítimo comprometido** | Roubo de credenciais | Uso de token JWT válido para ações não autorizadas |
| **Serviço/dependência comprometida** | Supply chain attack | Injeção de código via dependência vulnerável |

### 2.2 Superfícies de Ataque

| Superfície | Exposição | Risco |
|:---|:---|:---|
| REST API (`/api/v1/*`) | Externa (via frontend SPA) | Injeção, força bruta, IDOR, XSS |
| Keycloak (IdP) | Externa | Roubo de token, session hijacking |
| PostgreSQL (porta 5432) | Interna (VPC) | SQL injection via aplicação |
| Logs e métricas | Interna (observabilidade) | Vazamento de PII em logs |
| Container image (Docker/GraalVM) | Registry | Imagem com vulnerabilidades conhecidas |

### 2.3 Fluxos de Dados Sensíveis

| Fluxo | Dados | Proteção |
|:---|:---|:---|
| Frontend → Admin API | JWT (header), DTOs (body) | TLS 1.3, JWT RS256 |
| Admin API → Keycloak | Validação de token (OIDC) | TLS 1.3 |
| Admin API → PostgreSQL | Queries com `tenant_id` | JDBC (rede interna), RLS no banco |
| Admin API → SMTP | E-mails de convite/onboarding | TLS (STARTTLS) |
| Admin API → Observabilidade | Logs, métricas, tracing | Sem PII nos logs |

### 2.4 Matriz STRIDE

| Componente | S | T | R | I | D | E |
|:---|:---|:---|:---|:---|:---|:---|
| **REST API** | JWT RS256 | Bean Validation | RFC 7807 errors + AuditAspect | TLS 1.3 | Rate limiting | RBAC (@RequiresPermission) |
| **Keycloak IdP** | OIDC standard | Token signing RS256 | Audit log no IdP | TLS 1.3 | HA configurado | — |
| **PostgreSQL** | TLS (rede interna) | RLS (FORCE) | AuditAspect (app-level) | TLS + RLS | Connection pool (HikariCP) | RLS POLICY tenant_isolation |
| **BaseRepository** | — | tenant_id automático | — | — | — | IDOR prevention |
| **Observabilidade** | — | — | — | Sem PII nos logs | — | — |

### 2.5 Riscos Priorizados (Top 7)

| # | Risco | Severidade | Probabilidade | Impacto |
|:---|:---|:---|:---|:---|
| 1 | Cross-tenant data access (quebra de isolamento) | CRÍTICA | Baixa | Altíssimo |
| 2 | IDOR — acesso a recurso de outro tenant | ALTA | Média | Alto |
| 3 | Token JWT roubado/reutilizado | ALTA | Baixa | Alto |
| 4 | SQL Injection via query parameter | ALTA | Baixa | Alto |
| 5 | Vazamento de PII em logs | MÉDIA | Média | Médio |
| 6 | Dependência vulnerável (supply chain) | MÉDIA | Média | Médio |
| 7 | Força bruta em endpoints de autenticação | BAIXA | Média | Baixo |

---

## 3. Autenticação e Autorização

### 3.1 Mecanismo de Autenticação

| Aspecto | Implementação |
|:---|:---|
| **Provedor** | Keycloak (Identity Provider externo) |
| **Protocolo** | OAuth2 / OpenID Connect (OIDC) — Authorization Code Flow |
| **Token** | JWT assinado RS256 |
| **Validação** | `JwtAuthenticationFilter` (OncePerRequestFilter) — valida assinatura, expiração (`exp`), emissor (`iss`) |
| **Claims utilizadas** | `sub` (user_id), `tenant_id`, `roles`, `preferred_username` |
| **Sessão** | Stateless — sem sessão no servidor. `TenantContext` via ThreadLocal por request |

### 3.2 Política de Senhas e Credenciais

Delegado ao Keycloak:
- Mínimo 8 caracteres, letra + número (RN13-01 — PRD §6.6)
- Link de redefinição de senha de uso único (RN13-03)
- Sessão expira após 60 min de inatividade (RN13-02 — PRD §6.6)
- **O backend NUNCA gerencia senhas** — apenas valida JWT (ADR-04, PRD §5.2)

### 3.3 RBAC — Matriz de Papéis × Permissões

| Papel | Escopo | Permissões Principais |
|:---|:---|:---|
| **ADMIN_TENANT** | Todo o tenant | CRUD em todas as entidades, gestão de usuários, convites, planos, assinaturas |
| **MANAGER_BU** | Unidade(s) de Negócio vinculada(s) | CRUD em produtos/serviços da BU, dashboard da BU, leitura de permissões |
| **OPERATOR_BU** | Unidade(s) de Negócio vinculada(s) | Leitura de dashboard, leitura de catálogo, sem permissão de escrita |
| **AUDITOR** | Todo o tenant (futuro) | Leitura de trilha de auditoria (`GET /audit`). Previsto para fase futura (Charter §3.1, item 4) |

**Implementação:**
- `@RequiresPermission(resource, action)` — anotação AOP em controllers/services
- `RbacAspect` + `PermissionService` — consulta banco (`user_permission` + `role_resource`)
- Matriz RN10-01 nas tabelas `resource_action` (8 resources × 4 actions) + `role_resource` (seed V004)
- Sem cache TTL — alterações de permissão têm efeito imediato (RN11-03 — PRD §6.6)

### 3.4 Row-Level Security / Multi-Tenant Isolation

| Camada | Mecanismo | Tipo |
|:---|:---|:---|
| **1. PostgreSQL RLS** | `FORCE ROW LEVEL SECURITY` + `POLICY tenant_isolation USING (tenant_id = current_setting('app.current_tenant_id')::UUID)` em 5 tabelas (subscription, user, user_permission, business_unit, audit_log) | Preventiva (banco) |
| **2. BaseRepository** | `AND tenant_id = ?` automático em todas as queries | Preventiva (aplicação) |
| **3. Teste de Isolamento** | Testcontainers: tenant-A consulta → nunca vê dados de tenant-B | Detectiva |

### 3.5 Proteção contra IDOR

- Toda query inclui `tenant_id` (Camada 1 + Camada 2 acima)
- Tentativa de acessar recurso de outro tenant → retorna **vazio** (não 403 — PRD §6.1)
- `tenant_id` do JWT NUNCA é aceito como parâmetro de entrada — sempre do contexto de autenticação

### 3.6 Estratégia de Cache e Performance

> **Fonte:** [ARCHITECTURE.md §5.4](./ARCHITECTURE.md) — Estratégia de Cache — Caffeine.

| Aspecto | Configuração | Responsável | Verificação |
|:---|:---|:---|:---|
| **Biblioteca** | Caffeine Cache (`spring-boot-starter-cache`) | Backend | Actuator metrics `/cache.*` |
| **TTL padrão** | 5 minutos (`expireAfterWrite=5m`) | Backend | Teste de integração: cache hit/miss |
| **Tamanho máximo** | 10.000 entradas (`maximumSize=10000`) | Backend | Métrica `cache.size` |
| **Evicção** | LRU (Least Recently Used) | Backend | Monitoramento de evicções |
| **NÃO cacheado** | Permissões RBAC — consulta sempre em tempo real (RN11-03: "Sem cache TTL") | Backend | `PermissionService` sem `@Cacheable` |
| **Cacheado** | Planos ativos, módulos de plano, catálogo de produtos (baixa mutabilidade), dashboards (agregados TTL curto) | Backend | `@Cacheable` nos services correspondentes |
| **Invalidação** | `@CacheEvict(allEntries = true)` nos endpoints de create/update/soft delete | Backend | Teste: após POST, GET retorna dado atualizado |
| **ADR implícito** | Cache local (não distribuído) — simplicidade Fase 0. Redis reavaliado Fase 1 (>3 instâncias PRD) | Tech Lead | Revisão de arquitetura Fase 1 |

---

## 4. Proteção de Dados e Privacidade

### 4.1 Criptografia em Repouso

| Dado | Algoritmo | Escopo |
|:---|:---|:---|
| Senhas de usuários | bcrypt (Keycloak) | Identity Provider |
| Dados em disco (BD) | AES-256-XTS (criptografia de volume) + TLS | PostgreSQL — Infra/DevOps |
| Backups | AES-256-GCM (criptografia de bucket/volume) | Infra/DevOps |
| Logs em repouso | AES-256-XTS (criptografia de volume do nó K8s) | Infra/DevOps |

### 4.2 Criptografia em Trânsito

| Conexão | Protocolo |
|:---|:---|
| Frontend ↔ Admin API | TLS 1.3 |
| Admin API ↔ Keycloak | TLS 1.3 (OIDC) |
| Admin API ↔ PostgreSQL | TLS 1.2+ (rede interna VPC) |
| Admin API ↔ Admin API (inter-service) | mTLS — postergado (Fase 1). Atualmente serviço monolítico sem comunicação inter-service |
| Admin API ↔ SMTP | TLS (STARTTLS) |

### 4.3 Dados Sensíveis — Mascaramento e Retenção

| Campo | Política |
|:---|:---|
| **CNPJ** | Armazenado em claro (necessário para validação fiscal). Mascarado em logs: `***123456/0001**` |
| **E-mail** | Armazenado em claro. Mascarado em logs: `u***r@domain.com` |
| **Senhas** | NUNCA armazenadas pelo backend. Gerenciadas pelo Keycloak (bcrypt) |
| **Tokens JWT** | NUNCA logados. Validados e descartados após extração de claims |
| **Campos de auditoria** | `created_by`, `updated_by`, `deleted_by` — armazenados, não mascarados |

**Política de expurgo:** Soft delete mantém dados por 5 anos (requisito fiscal). Auditoria (`audit_log`) — append-only, sem expurgo na Fase 0.

### 4.4 Conformidade Regulatória

| Regulação | Aplicabilidade | Status |
|:---|:---|:---|
| **LGPD** | Dados de pessoas físicas (usuários) | Em conformidade: soft delete (direito ao esquecimento), mascaramento de PII, política de retenção documentada |
| **PCI DSS** | Não aplicável — sem processamento de pagamentos nesta fase (Gateway delegado ao Keycloak/Stripe futuro) | N/A |
| **SOC 2** | Não aplicável na Fase 0 | Roadmap Fase 1 |

---

## 5. Segurança de API e Comunicação

### 5.1 Rate Limiting

| Aspecto | Configuração |
|:---|:---|
| **Por IP** | 100 req/min (endpoints públicos: `/actuator/health`, `/error`) |
| **Por tenant** | 500 req/min (endpoints autenticados) |
| **Por endpoint sensível** | 20 req/min (`POST /users/invite`, `POST /auth/*`) |
| **Implementação** | Spring Filter + Bucket4j — **Sprint 6** (backlog). Enquanto não implementado, a proteção contra força bruta é delegada ao Keycloak (lockout após 5 tentativas) e ao WAF (PRD). Ver OWASP A05 (§6) |

### 5.2 Política de CORS

```
Access-Control-Allow-Origin: https://admin.fbso.org (PRD)
Access-Control-Allow-Origin: http://localhost:3000 (DEV)
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Authorization, Content-Type
Access-Control-Allow-Credentials: true
Access-Control-Max-Age: 3600
```

Configurado em `WebConfig.java` — **origens restritas, nunca `*` em produção**.

### 5.3 Input Validation

| Camada | Mecanismo |
|:---|:---|
| **Controller (DTO)** | `@Valid` + Bean Validation annotations (`@NotBlank`, `@Email`, `@Size`, `@CNPJ`) |
| **Service** | Validação de regras de negócio — 47 RNs documentadas (PRD §6.6) |
| **Repository** | Parameterized queries (JDBC Template) — **ZERO concatenação manual de SQL** |
| **Headers** | JWT validado pelo framework (Spring Security + `JwtAuthenticationFilter`) |

### 5.4 Content Security Policy (CSP) e Headers HTTP

Headers configurados via Spring Security:

| Header | Valor |
|:---|:---|
| `X-Content-Type-Options` | `nosniff` |
| `X-Frame-Options` | `DENY` |
| `X-XSS-Protection` | `0` (descontinuado — proteção via CSP) |
| `Strict-Transport-Security` | `max-age=31536000; includeSubDomains` (HSTS) |
| `Cache-Control` | `no-store` (para endpoints autenticados) |

### 5.5 Proteção contra CSRF

Desabilitada para API REST (stateless com JWT). CSRF não se aplica a APIs que não usam cookies de sessão.

---

## 6. Cobertura OWASP Top 10

| # | Categoria | Status | Controles Aplicados |
|:---|:---|:---|:---|
| **A01** | Broken Access Control | ✅ Mitigado | RBAC DB-backed (@RequiresPermission), RLS (FORCE), tenant_id em toda query, IDOR prevention |
| **A02** | Cryptographic Failures | ✅ Mitigado | TLS 1.3, bcrypt (Keycloak), senhas NUNCA no backend, JWT RS256 |
| **A03** | Injection | ✅ Mitigado | JDBC Template (parameterized queries), Bean Validation, sem concatenação SQL |
| **A04** | Insecure Design | ✅ Mitigado | Threat model STRIDE documentado, ADRs de segurança, 3 camadas de defesa |
| **A05** | Security Misconfiguration | ⚠️ Parcial | Spring Security config documentado, CORS restrito, HSTS. Rate limiting pendente (Sprint 6) |
| **A06** | Vulnerable Components | ⚠️ Parcial | Dependabot ativo, stack versionada. Falta política formal de atualização (ver §7) |
| **A07** | Auth Failures | ✅ Mitigado | JWT stateless, sem sessão no servidor, 60min timeout (Keycloak), lockout após 5 tentativas (Keycloak) |
| **A08** | Software/Data Integrity | ✅ Mitigado | GraalVM Native Image (assinatura), Docker image scanning, Maven checksum |
| **A09** | Security Logging Failures | ✅ Mitigado | AuditAspect — 100% ações, RFC 7807 errors, logs estruturados (request_id, tenant_id, user_id), sem PII em logs |
| **A10** | SSRF | ✅ Mitigado | Sem chamadas HTTP arbitrárias iniciadas pelo backend na Fase 0. Comunicação apenas com serviços conhecidos (Keycloak, SMTP) |

---

## 7. Gestão de Dependências (SCA)

### 7.1 Ferramenta de Análise

| Ferramenta | Escopo | Frequência |
|:---|:---|:---|
| **Dependabot** | `pom.xml` — dependências Maven | A cada push / semanal |
| **Semgrep (SAST)** | Código-fonte Java | A cada push / PR |
| **Trivy / Docker Scout** | Imagem Docker (GraalVM Native Image) | A cada build |

### 7.2 Política de Atualização

| Severidade da CVE | Ação | Prazo |
|:---|:---|:---|
| **CRÍTICA (CVSS ≥ 9.0)** | Atualização imediata | 24h |
| **ALTA (CVSS 7.0–8.9)** | Atualização na próxima sprint | 2 semanas |
| **MÉDIA (CVSS 4.0–6.9)** | Planejar atualização | 4 semanas |
| **BAIXA (CVSS < 4.0)** | Monitorar, atualizar com a próxima versão major | Conforme roadmap |

Exemplo: PostgreSQL driver 42.7.11 (CVE-2026-42198, DoS CVSS 7.5) — atualizado na Sprint 5 Frente 0.

### 7.3 SLSA Framework

Nível atual: **SLSA L1** (build automatizado, origem do código rastreável).
Objetivo Fase 1: SLSA L2 (build assinado, proveniência atestada).

---

## 8. Pipeline de Segurança (DevSecOps)

### 8.1 SAST — Static Application Security Testing

| Ferramenta | Configuração | Frequência | Regras |
|:---|:---|:---|:---|
| **Semgrep** | `semgrep scan --config=auto` | A cada push / PR | Regras padrão Java + OWASP + custom |
| **SonarQube** | Integração CI/CD | A cada PR | Quality Gate: zero bugs críticos, zero vulnerabilidades |

### 8.2 Secret Scanning

| Ferramenta | Padrões Detectados | Gatilho |
|:---|:---|:---|
| **Semgrep Secrets** | API keys, tokens, senhas, connection strings | Pre-commit hook + CI |
| **Gitleaks** | Padrões customizados (FBSO-specific) | Pre-commit hook |

**Regra absoluta:** NUNCA hardcode secrets. Todos os segredos via variáveis de ambiente (`application.yml` → `${ENV_VAR}`).

### 8.3 DAST — Dynamic Application Security Testing

Planejado para Sprint 7 (próximo ao M7 — Aceite Final). Escopo: OWASP ZAP scan automatizado contra ambiente HML.

### 8.4 Container/Image Scanning

| Ferramenta | Gatilho |
|:---|:---|
| **Trivy** | A cada build de imagem Docker/GraalVM |
| **Docker Scout** | Push para registry |

### 8.5 Git Hooks e Pre-Commit

```bash
# pre-commit
./mvnw spotless:check       # formatação
./mvnw pmd:check             # análise estática
gitleaks detect --verbose    # secrets
```

---

## 9. Segurança de Infraestrutura

### 9.1 Configuração de Rede

| Aspecto | DEV | HML | PRD |
|:---|:---|:---|:---|
| **API exposta** | localhost:8080 | VPC interna | VPC + WAF |
| **PostgreSQL** | localhost:5432 | VPC interna | Serviço gerenciado + VPC |
| **Keycloak** | localhost:8081 | VPC interna | Cluster dedicado |
| **WAF** | Não | Não | Sim (AWS WAF / Cloudflare) |

### 9.2 Gestão de Segredos

| Ambiente | Mecanismo |
|:---|:---|
| **DEV** | `.env` (não versionado) + Docker Compose |
| **HML** | Kubernetes Secrets |
| **PRD** | HashiCorp Vault / AWS Secrets Manager |

Segredos gerenciados: `DB_PASSWORD`, `KEYCLOAK_CLIENT_SECRET`, `SMTP_PASSWORD`, `JWT_ISSUER_URI`.

### 9.3 Hardening de Containers

- **Imagem base:** `chainguard/graalvm-native` (distroless, sem shell)
- **Usuário:** não-root (`USER 1001`)
- **Filesystem:** read-only (exceto `/tmp`)
- **Capabilities:** drop ALL, sem `NET_RAW`

### 9.4 Política de Patching

| Componente | Frequência |
|:---|:---|
| **Imagem base** | Semanal (Dependabot + rebuild) |
| **Dependências Maven** | Ver §7.2 |
| **Keycloak** | Acompanhar releases estáveis (mensal) |

---

## 10. Checklist de Verificação (SDD)

> Derivado do [GLOBAL-SECURITY.md](../../../../../../.specs/security/GLOBAL-SECURITY.md) — Regras de Ouro e Checklist SDD.

### Regras de Ouro (Inegociáveis)

| # | Regra | Status | Evidência |
|:---|:---|:---|:---|
| **1** | Princípio do Menor Privilégio | ✅ Implementado | RBAC AOP (§3.3), 4 papéis, `@RequiresPermission`, RLS FORCE, Zero permissão por default |
| **2** | Zero Hardcoded Secrets | ✅ Implementado | Todas credenciais via `application.yml` → `${ENV_VAR}`. Gitleaks no pre-commit (§8.2) |
| **3** | Não Confiar no Input do Usuário | ✅ Implementado | Bean Validation (§5.3), JDBC Template (parameterized queries), `TenantContext` do JWT (não do request body) |

### Checklist SDD

| # | Categoria | Item | Status |
|:---|:---|:---|:---|
| **A1** | Autenticação | Middleware de sessão/token em todas as rotas? | ✅ `JwtAuthenticationFilter` em todas as rotas autenticadas |
| **A2** | Autorização | RBAC validado para cada endpoint? | ✅ `@RequiresPermission` + `RbacAspect` |
| **A3** | Autorização | Proteção IDOR (recurso pertence ao tenant/usuario)? | ✅ tenant_id em toda query (§3.5) |
| **D1** | Dados | Senhas hasheadas com bcrypt/argon2? | ✅ bcrypt (Keycloak) — §4.1 |
| **D2** | Dados | Logs NUNCA registram dados sensíveis? | ✅ Filtro de PII em logs (§4.3) |
| **D3** | Dados | Campos sensíveis mascarados na API? | ✅ CNPJ e e-mail mascarados em logs (§4.3) |
| **V1** | Validação | Schema validation no payload? | ✅ `@Valid` + Bean Validation em DTOs (§5.3) |
| **V2** | Validação | Queries 100% parameterizadas? | ✅ JDBC Template — zero concatenação (§5.3) |
| **V3** | Validação | Sanitização contra XSS? | ✅ CSP headers + validação de input (§5.4) |
| **I1** | Infra | Rate limiting em endpoints expostos? | ⚠️ Pendente Sprint 6 (§5.1) |
| **I2** | Infra | CORS restrito a domínios autorizados? | ✅ `WebConfig.java` — sem `*` em PRD (§5.2) |
| **I3** | Infra | Tratamento de erros seguro (sem stack traces)? | ✅ RFC 7807 — `GlobalExceptionHandler` (§5.3) |

---

## 11. Resposta a Incidentes e Monitoramento

### 11.1 Alertas de Segurança

| Alerta | Gatilho | Canal |
|:---|:---|:---|
| **Taxa de erro 4xx elevada** | > 50 erros/min por tenant | Slack #alerts + PagerDuty |
| **Taxa de erro 5xx** | Qualquer 5xx em PRD | PagerDuty (imediato) |
| **Tentativas de IDOR** | Log pattern: tenant_id mismatch | Revisão semanal |
| **Falha de validação JWT** | > 10 falhas/min do mesmo IP | Slack #security |
| **Vulnerabilidade crítica (CVE)** | Dependabot/Semgrep alert CRITICAL | PagerDuty (§7.2) |

### 11.2 Logs de Auditoria

| Aspecto | Configuração |
|:---|:---|
| **O que é logado** | Ação (CREATE/UPDATE/DELETE), entidade, usuário, tenant, timestamp, IP, request_id |
| **Onde** | `audit_log` (PostgreSQL) + stdout (JSON estruturado para observabilidade) |
| **Retenção** | 5 anos (banco), 30 dias (logs estruturados) |
| **Proteção** | Append-only, sem UPDATE/DELETE na tabela `audit_log` |

### 11.3 Plano de Resposta a Incidentes

| Fase | Ação | Responsável |
|:---|:---|:---|
| **Detecção** | Alerta automático (PagerDuty) ou reporte manual | Monitoramento / Equipe |
| **Contenção** | Isolar serviço afetado (K8s cordon), revogar tokens comprometidos (Keycloak admin) | DevOps / Tech Lead |
| **Erradicação** | Corrigir vulnerabilidade, aplicar patch, rebuild + deploy | Equipe de Desenvolvimento |
| **Recuperação** | Restaurar backups (se necessário), reativar serviço, validar isolamento | DevOps |
| **Post-Mortem** | Documentar causa raiz, atualizar threat model, adicionar teste de regressão | Tech Lead |

### 11.4 Divulgação de Vulnerabilidades

Conforme GLOBAL-SECURITY.md: reportar vulnerabilidades para **org-fbso@gmail.com**. Não abrir issue pública. Prazo de resposta: 48h.

---

### 11.5 Security ADRs (Architecture Decision Records)

> **Fonte:** [ARCHITECTURE.md §9](./ARCHITECTURE.md) — Decisões de Design (ADRs Locais). Abaixo os ADRs com impacto direto em segurança.

| ADR | Decisão | Contexto | Alternativas | Justificativa |
|:---|:---|:---|:---|:---|
| **ADR-01** (global) / ADR-L07 (local) | Shared Database + tenant_id | Múltiplos tenants no mesmo banco | Schema-per-tenant, database-per-tenant | Simplicidade operacional; RLS (FORCE) garante isolamento no nível do banco |
| **ADR-04** (global) | Keycloak como IdP | Autenticação centralizada | Auth0, Cognito, Spring Security local | Open source, self-hosted, suporte OIDC/SAML, sem custo de licença |
| **ADR-07** (global) | JWT Stateless | Sem sessão no servidor | Sessão HTTP, OAuth2 stateful | Escalabilidade horizontal, sem afinidade de sessão, TenantContext via ThreadLocal |
| **ADR-L02** (local) | Aspectos AOP para cross-cutting | RBAC e Auditoria sem poluir services | Anotações + interceptors, filtros manuais | Separação de concerns, zero risco de esquecimento, código de negócio limpo |
| **ADR-L03** (local) | Auditoria assíncrona | Gravação de audit_log não bloqueia resposta | Síncrono, fila de mensagens | Performance (não bloqueia o request); trade-off: perda em crash aceitável Fase 0 |
| **ADR-L07** (local) | PostgreSQL RLS (FORCE) | Isolamento multi-tenant no banco | Aplicação-only (BaseRepository), schema-per-tenant | Defesa em profundidade (3 camadas); garantia no nível do banco impossível de burlar |

> 💡 **Referência completa:** [ARCHITECTURE.md §9](./ARCHITECTURE.md) contém os 7 ADRs locais (ADR-L01 a ADR-L07) com justificativas detalhadas. Os ADRs acima são o subconjunto com impacto direto em segurança.

---

## 12. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.1 | 21/07/2026 | **GATE-SECURITY-SCOPE COMPLIANCE:** Validação em 5 dimensões (1 APROVADO, 4 RESSALVAS). 5 NCs corrigidas: §1.4 pipeline unificado (fluxo 8 etapas), §3.6 estratégia de cache Caffeine (TTL/invalidação/escopo), §4.1 algoritmos criptografia (AES-256-XTS/GCM + logs), §4.2 TLS 1.2+ PG + mTLS postergado, §11.5 Security ADRs (6 decisões com contexto/alternativas/justificativa). Status: COMPLIANCE. | Agente GATE-SECURITY-SCOPE/IA |
| 1.0 | 21/07/2026 | Criação inicial: 12 seções — visão geral, threat model STRIDE, autenticação/autorização (JWT+RBAC+RLS), proteção de dados (TLS+bcrypt+LGPD), segurança de API (rate limit+CORS+Bean Validation), OWASP Top 10, SCA (Dependabot+SLSA), pipeline DevSecOps (SAST+Secret Scanning), infraestrutura (WAF+Vault+hardening), checklist SDD (GLOBAL-SECURITY.md), resposta a incidentes, changelog. Status: aguardando GATE. | Agente Arquiteto de Segurança/IA |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Segurança/Claude. Foram utilizados os skills: security-auditor, threat-modeling-expert, security-best-practices, security-reviewer, documentation-writer. Referência normativa: GLOBAL-SECURITY.md.*
