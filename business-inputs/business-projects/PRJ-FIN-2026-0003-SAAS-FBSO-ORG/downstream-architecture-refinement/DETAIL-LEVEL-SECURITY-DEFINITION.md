# DETAIL-LEVEL-SECURITY-DEFINITION — Segurança Detail-Level

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 31/07/2026
- **Fase:** F3 — Downstream Architecture Refinement
- **Padrões FBSO:** Keycloak IAM, Kong Gateway, Kong↔Keycloak Service-ID/Token-ID
- **Referência:** [GLOBAL-SECURITY.md](../../../.specs/security/GLOBAL-SECURITY.md), [Arquitetura Detail-Level (F2)](./DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md)

---

## 1. Threat Model (STRIDE por Componente)

### 1.1 Matriz STRIDE

| Componente | Spoofing | Tampering | Repudiation | Info Disclosure | DoS | Elevation of Privilege |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|
| **Kong Gateway** | ⚠️ Service-ID spoofing | ⚠️ Header injection maliciosa | ✅ Log de requests | ⚠️ JWT exposto em logs | ⚠️ Rate limiting bypass | ⚠️ Bypass de validação JWT |
| **Keycloak IAM** | ⚠️ Credential stuffing | ⚠️ Token manipulation | ✅ Audit log | ⚠️ User enumeration | ⚠️ Brute force login | ⚠️ Privilege escalation via roles |
| **Backend API** | ✅ Já validado pelo Kong | ⚠️ SQL Injection | ⚠️ Ações sem audit | ⚠️ Cross-tenant data leakage | ⚠️ Query pesada sem limite | ⚠️ Bypass de @PreAuthorize |
| **Frontend Portal** | ✅ OIDC via Keycloak | ⚠️ XSS | ✅ Client-side logging | ⚠️ Token storage | ⚠️ Bundle size/loading | ⚠️ Route guard bypass |
| **PostgreSQL** | ✅ JDBC auth | ⚠️ Direct DB access | ✅ Triggers audit_log | ⚠️ RLS bypass | ⚠️ Connection exhaustion | ⚠️ Escalação de privilégio SQL |
| **Redis** | ✅ Internal network | ⚠️ Cache poisoning | ❌ Sem audit nativo | ⚠️ Dados em memória | ⚠️ Memory exhaustion | ✅ Internal only |

> **Legenda STRIDE:**
> - ✅ **Mitigado** — O controle de segurança já está implementado pelo padrão corporativo ou pela arquitetura
> - ⚠️ **Requer Atenção** — Existe risco que demanda mitigação ativa (descrita nas ameaças críticas abaixo)
> - ❌ **Não Mitigado / Não Aplicável** — O componente não possui controle nativo para esta dimensão; risco deve ser aceito ou mitigado em outra camada

### 1.2 Ameaças Críticas e Mitigações

#### T1 — Cross-Tenant Data Leakage (Info Disclosure — Backend + PostgreSQL)
| Atributo | Valor |
|:---|:---|
| **Severidade** | 🔴 Crítica |
| **Cenário** | Tenant A acessa dados do Tenant B por bug no RLS ou tenant_id incorreto |
| **Mitigação** | RLS FORCE em todas as tabelas; `TenantAwareDataSource` seta `app.current_tenant_id`; teste automatizado de isolamento por tenant; Kong header `X-Tenant-ID` como fonte única |
| **Verificação** | Teste de penetração: autenticar como Tenant A, tentar acessar recursos do Tenant B → 403/404 |

#### T2 — JWT Theft + Replay (Spoofing — Frontend + Kong)
| Atributo | Valor |
|:---|:---|
| **Severidade** | 🔴 Crítica |
| **Cenário** | Atacante rouba JWT do localStorage e faz replay de requests |
| **Mitigação** | JWT com expiração curta (15min access + 1h refresh); refresh token rotation; Kong valida `exp` e `iat`; HTTPS em toda cadeia |
| **Verificação** | Tentar usar JWT expirado → Kong retorna 401 |

#### T3 — Bypass de @PreAuthorize (Elevation of Privilege — Backend)
| Atributo | Valor |
|:---|:---|
| **Severidade** | 🟠 Alta |
| **Cenário** | Novo endpoint adicionado sem anotação de segurança → acesso não autorizado |
| **Mitigação** | `denyAll` por padrão no SecurityConfig; `@PreAuthorize` obrigatório em todos os controllers; code review checklist |
| **Verificação** | Scan automatizado: todo endpoint REST tem @PreAuthorize |

#### T4 — SQL Injection (Tampering — Backend)
| Atributo | Valor |
|:---|:---|
| **Severidade** | 🟠 Alta |
| **Cenário** | Query construída com concatenação de string + input do usuário |
| **Mitigação** | Spring Data JDBC com prepared statements; nunca concatenar SQL; Semgrep rule para detectar concatenação SQL |
| **Verificação** | SAST (Semgrep) no CI; code review obrigatório para queries nativas |

#### T5 — Brute Force Login (DoS — Keycloak)
| Atributo | Valor |
|:---|:---|
| **Severidade** | 🟡 Média |
| **Cenário** | Atacante tenta múltiplas senhas para enumerar contas ou quebrar senha |
| **Mitigação** | Keycloak: bloqueio após 5 tentativas (15min); Kong: rate limiting por IP; Cloudflare: DDoS protection |
| **Verificação** | 6+ tentativas de login com senha errada → conta bloqueada temporariamente |

---

## 2. Matriz de Controles (OWASP ASVS L1 + L2)

| ASVS ID | Categoria | Nível | Controle | Implementação no Projeto |
|:---|:---|:---:|:---|:---|
| **V2.1.1** | Authentication | L1 | Credenciais não em URL/body de GET | Todas as requisições de auth via POST com body JSON |
| **V2.1.2** | Authentication | L1 | Senhas ≥ 8 caracteres | Keycloak password policy: mínimo 8, complexidade (upper+lower+digit+special) |
| **V2.2.1** | Authentication | L2 | MFA disponível | Keycloak OTP via Authenticator App — configurável por realm |
| **V2.3.1** | Authentication | L1 | Bloqueio após tentativas | Keycloak brute force detection: 5 tentativas → bloqueio 15min |
| **V2.4.1** | Authentication | L2 | JWT expiração ≤ 15min | Access token: 15min; Refresh token: 1h com rotação |
| **V3.1.1** | Session Mgmt | L1 | Logout invalida token | Keycloak logout endpoint + backchannel logout |
| **V3.2.1** | Session Mgmt | L2 | Token revogável | Keycloak Admin API: revogar tokens de usuário desativado |
| **V4.1.1** | Access Control | L1 | Deny by default | Spring Security: `.anyRequest().denyAll()` + `@PreAuthorize` explícito |
| **V4.1.2** | Access Control | L1 | RBAC para endpoints | PermissionEvaluator + claims do Kong: roles + permissions |
| **V4.1.3** | Access Control | L2 | Autorização por recurso | Tenant ID validation: usuário só acessa recursos do seu tenant |
| **V4.2.1** | Access Control | L2 | IDOR prevention | UUIDs aleatórios (gen_random_uuid()); acesso validado contra tenant_id |
| **V5.1.1** | Input Validation | L1 | Input sanitization | Jakarta Bean Validation em todos os DTOs; @Valid em controllers |
| **V5.2.1** | Input Validation | L1 | SQL Injection prevention | Prepared statements (Spring Data JDBC); Semgrep SAST |
| **V5.3.1** | Input Validation | L2 | XSS prevention | Next.js auto-escaping de output; CSP header via Cloudflare |
| **V6.1.1** | Cryptography | L1 | TLS em trânsito | HTTPS em toda cadeia (Cloudflare → Kong → Backend) |
| **V6.2.1** | Cryptography | L2 | Dados sensíveis criptografados em repouso | AES-256 para campos sensíveis (documentos); PostgreSQL TDE opcional |
| **V7.1.1** | Logging | L1 | Logs sem dados sensíveis | SLF4J: nunca loga JWT, senha, token, CPF/CNPJ |
| **V7.2.1** | Logging | L2 | Audit trail imutável | Tabela `audit_log` com trigger PostgreSQL; append-only |
| **V8.1.1** | Data Protection | L1 | Cache sem dados sensíveis | Redis: apenas métricas agregadas e contadores; sem PII |
| **V8.2.1** | Data Protection | L2 | Right to erasure (LGPD) | Soft delete + anonimização; endpoint de exclusão de dados |
| **V9.1.1** | Communication | L1 | HTTPS only | HSTS via Cloudflare; Kong: redirect HTTP→HTTPS |

---

## 3. Especificação IAM (Keycloak)

### 3.1 Estrutura de Realms

```
Master Realm (admin interno)
  └── fbso-admin (realm do portal administrativo)
        ├── Client: fbso-admin-portal (public, Authorization Code + PKCE)
        ├── Client: fbso-admin-api (confidential, client_credentials — Kong↔Keycloak)
        └── Roles: ROLE_FBSO_ADMIN, ROLE_FBSO_OPERATOR, ROLE_FBSO_AUDITOR

Para cada Tenant (provisionado dinamicamente):
  └── tenant-{uuid} (realm do tenant)
        ├── Client: tenant-portal (public, Authorization Code + PKCE)
        ├── Client: tenant-api (confidential, client_credentials — backend→Keycloak)
        └── Roles: ROLE_TENANT_ADMIN, ROLE_TENANT_MANAGER, ROLE_TENANT_OPERATOR, ROLE_TENANT_AUDITOR
```

### 3.2 Protocol Mappers (JWT Claims)

| Mapper | Claim | Tipo | Valor |
|:---|:---|:---|:---|
| Tenant ID | `tenant_id` | String | UUID do tenant |
| User Roles | `realm_access.roles` | JSON Array | `["ROLE_TENANT_ADMIN", "ROLE_TENANT_MANAGER"]` |
| Permissions | `permissions` | JSON Array | `["MANAGE_USERS", "VIEW_DASHBOARD", "EDIT_PRODUCTS"]` |
| Business Unit IDs | `business_unit_ids` | JSON Array | `["uuid-bu-001", "uuid-bu-002"]` |
| User ID | `sub` | String | UUID do usuário Keycloak |
| Email | `email` | String | Email do usuário |

### 3.3 Fluxos de Autenticação

| Fluxo | Client | Grant Type | Uso |
|:---|:---|:---|:---|
| Portal Admin Login | fbso-admin-portal | Authorization Code + PKCE | Admin FBSO acessa portal |
| Portal Cliente Login | tenant-portal | Authorization Code + PKCE | Cliente acessa seu portal |
| Kong↔Keycloak Validation | fbso-admin-api / tenant-api | client_credentials | Kong valida Service-ID/Token-ID |
| Backend→Keycloak Admin | realm-admin-cli | client_credentials (service account) | Provisioning de realms |

### 3.4 Política de Senhas

| Parâmetro | Valor |
|:---|:---|
| Comprimento mínimo | 8 caracteres |
| Complexidade | Upper + Lower + Digit + Special |
| Histórico | Não repetir últimas 5 senhas |
| Expiração | 90 dias (admin) / 180 dias (tenant) |
| Bloqueio | 5 tentativas → 15 minutos |
| Recuperação | Link por email, expira em 30 minutos, uso único |

---

## 4. Matriz RBAC Granular

### 4.1 Roles × Permissions × Resources

| Permissão | FBSO Admin | FBSO Operador | Tenant Admin | Gerente | Operador | Auditor |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|
| **Tenants** | | | | | | |
| Criar/Ativar/Suspender Tenant | ✅ | ✅ | ❌ | ❌ | ❌ | ❌ |
| Visualizar todos Tenants | ✅ | ✅ | ❌ | ❌ | ❌ | ✅ |
| Visualizar próprio Tenant | — | — | ✅ | ✅ | ✅ | ✅ |
| **Planos** | | | | | | |
| Criar/Editar/Desativar Plano | ✅ | ❌ | ❌ | ❌ | ❌ | ❌ |
| Visualizar Planos | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Upgrade/Downgrade (self-service) | — | — | ✅ | ❌ | ❌ | ❌ |
| **Usuários** | | | | | | |
| Convidar/Editar/Desativar Usuário | ✅ | ✅ | ✅ | ✅ | ❌ | ❌ |
| Atribuir Papéis | ✅ | ✅ | ✅ | ❌ | ❌ | ❌ |
| Suspensão Temporária | ✅ | ✅ | ✅ | ✅ | ❌ | ❌ |
| **Unidades de Negócio** | | | | | | |
| Criar/Editar/Desativar BU | ❌ | ❌ | ✅ | ✅ | ❌ | ❌ |
| Visualizar BU | ❌ | ❌ | ✅ | ✅ | ✅ | ✅ |
| **Produtos** | | | | | | |
| Criar/Editar Produto | ❌ | ❌ | ✅ | ✅ | ✅ | ❌ |
| Ativar/Desativar Produto | ❌ | ❌ | ✅ | ✅ | ❌ | ❌ |
| **Dashboard** | | | | | | |
| Dashboard Admin (todas contas) | ✅ | ✅ | ❌ | ❌ | ❌ | ❌ |
| Dashboard Cliente (própria conta) | — | — | ✅ | ✅ | ✅ | ✅ |
| **Auditoria** | | | | | | |
| Visualizar Logs Admin | ✅ | ❌ | ❌ | ❌ | ❌ | ✅ |
| Visualizar Logs Tenant | — | — | ✅ | ❌ | ❌ | ✅ |

---

## 5. Data Protection

### 5.1 Criptografia

| Camada | Método | Escopo |
|:---|:---|:---|
| Em trânsito | TLS 1.3 | Toda comunicação externa e interna |
| Em repouso — banco | AES-256 (PostgreSQL TDE) | Todo o schema `fbso_portal` |
| Em repouso — campos sensíveis | AES-256 (aplicação) | Documentos, tokens de API de terceiros |
| Em repouso — backups | AES-256 (DigitalOcean Spaces) | Backups diários do PostgreSQL |

### 5.2 RLS Policies (PostgreSQL)

```sql
-- Todas as tabelas do schema fbso_portal com RLS FORCE
ALTER TABLE fbso_portal.tenants ENABLE ROW LEVEL SECURITY;
ALTER TABLE fbso_portal.tenants FORCE ROW LEVEL SECURITY;

CREATE POLICY tenant_isolation ON fbso_portal.tenants
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::uuid);
```

### 5.3 LGPD — Direitos do Titular

| Direito | Implementação |
|:---|:---|
| **Acesso** | Endpoint `GET /api/v1/users/{id}/data-export` — exporta todos os dados do usuário |
| **Correção** | Endpoint `PUT /api/v1/users/{id}` — edita dados cadastrais |
| **Exclusão (Right to Erasure)** | Soft delete: `deleted_at` + anonimização de PII (nome → "USUARIO_EXCLUIDO", email → hash) |
| **Portabilidade** | Export JSON estruturado conforme formato LGPD |
| **Consentimento** | Registro de consentimento em tabela `consent_log` com timestamp e finalidade |

---

## 6. Riscos de Segurança

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---|:---|
| Cross-tenant data leakage via RLS bypass | Baixa | Crítico | RLS FORCE; testes automatizados de isolamento; pentest trimestral |
| JWT theft de localStorage (XSS) | Média | Crítico | CSP header; expiração curta; refresh rotation; httpOnly para refresh token |
| SQL Injection em query nativa | Baixa | Alto | Prepared statements; Semgrep SAST no CI; code review obrigatório |
| Endpoint sem @PreAuthorize (privilege escalation) | Média | Alto | denyAll default; scan automatizado de annotations; code review checklist |
| Keycloak credential stuffing | Média | Médio | Brute force protection; rate limiting no Kong; MFA opcional |
| Secrets em código ou logs | Baixa | Alto | Gitleaks no CI; nunca logar .env ou secrets; Kubernetes Secrets em produção |

---

🤖 *Documento gerado pelo Security Architect — Fase 3 do Downstream Architecture Refinement. Padrões FBSO: Keycloak IAM, Kong Gateway, Service-ID/Token-ID. Alinhado com GLOBAL-SECURITY.md.*
