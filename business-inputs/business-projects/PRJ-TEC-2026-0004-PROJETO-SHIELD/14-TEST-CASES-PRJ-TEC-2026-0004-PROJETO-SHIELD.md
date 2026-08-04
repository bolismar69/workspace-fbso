# Casos de Teste: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 03-SRS, 13-TEST-PLAN |
| **Solução Técnica** | ms-shield-identity-auth |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. Test Case Catalog

| ID | Feature (SRS) | Precondition | Steps | Expected Result | Postcondition | Priority |
|----|-------------|-------------|-------|----------------|--------------|----------|
| TC-001 | F-01 — Reconhecimento | Redis com mapeamento `escola-alfa.com` → `realm-escola-alfa` | 1. GET /auth/login com header `X-Tenant-Host: escola-alfa.com` 2. Verificar response | 302 redirect para Keycloak `/realms/realm-escola-alfa/...` | Browser redirecionado | High |
| TC-002 | F-01 — Domínio não mapeado | Redis sem mapeamento para `desconhecido.com` | 1. GET /auth/login com header `X-Tenant-Host: desconhecido.com` | 401 com mensagem padronizada, sem detalhes internos | — | High |
| TC-003 | F-02 — Login PKCE | Tenant `escola-alfa` ativo | 1. Seguir redirect 302 do TC-001 2. Login no Keycloak com credenciais válidas 3. Callback com code | 302 para /app; cookies setados com HttpOnly, Secure, SameSite=Strict | Sessão criada no PostgreSQL | High |
| TC-004 | F-02 — Cookies inacessíveis via JS | Sessão ativa após TC-003 | 1. No browser: `document.cookie` 2. No browser: `fetch('/auth/me')` com `credentials: 'include'` | `document.cookie` não retorna tokens; `/auth/me` retorna perfil | — | High |
| TC-005 | F-02 — PKCE state inválido | — | 1. GET /auth/callback?code=xyz&state=invalid | 401 — state não confere (CSRF bloqueado) | — | High |
| TC-006 | F-03 — Perfil do Usuário | Sessão ativa | 1. GET /auth/me com cookie de sessão válido | 200 {user_id, email, roles, tenant_id}. Latência <15ms | — | High |
| TC-007 | F-03 — Refresh Token | Sessão ativa, refresh_token não expirado | 1. POST /auth/refresh com cookie refresh_token | 200; novos cookies setados; expires_in > 0 | Sessão estendida | Medium |
| TC-008 | F-03 — Refresh Token expirado | Sessão expirada (>30min inativa) | 1. POST /auth/refresh com cookie refresh_token expirado | 401; redirecionar para /auth/login | Sessão terminada | Medium |
| TC-009 | F-03 — Logout | Sessão ativa | 1. POST /auth/logout | 200; cookies limpos (Max-Age=0); sessão Keycloak invalidada | Sessão destruída | High |
| TC-010 | F-03 — Logout com Keycloak indisponível | Sessão ativa, Keycloak inacessível | 1. Simular Keycloak down 2. POST /auth/logout | 200; cookies locais limpos mesmo sem Keycloak | Cookies limpos | Medium |
| TC-011 | F-04 — Cross-Tenant Block | Tenant A (escola-alfa) logado, dados Tenant B no banco | 1. GET /api/dados com token Tenant A 2. Query filtra por escola=B | 200 {data: []} — 0 linhas retornadas | Sem vazamento | Critical |
| TC-012 | F-04 — Same-Tenant OK | Tenant A logado, dados Tenant A no banco | 1. GET /api/dados com token Tenant A 2. Query filtra por escola=A | 200 {data: [...]} — dados retornados | — | Critical |
| TC-013 | F-08 — Suspensão de Tenant | Tenant `escola-alfa` ativo, depois suspenso | 1. GET /auth/me → 200 2. Marcar tenant como suspenso 3. GET /auth/me | 403 — acesso bloqueado em <1s | Sessão revogada | High |
| TC-014 | F-08 — Sessões ativas revogadas | Tenant com 10 sessões ativas, depois suspenso | 1. Criar 10 sessões 2. Suspender tenant 3. Todas as 10 chamadas /auth/me | 403 em todas as 10 | Todas sessões revogadas | High |

## 2. Happy Path Cases

| Fluxo | Passos | Resultado Esperado |
|-------|--------|-------------------|
| Acesso normal | Acessar app → redirecionar login → autenticar → usar app → sair | Login transparente, sessão mantida, logout completo |
| Renovação silenciosa | Usar app por 2h sem fechar | Sessão renovada automaticamente, sem novas interações de login |
| Onboarding novo cliente | Provisionar Realm + DNS + Redis → validar acesso | Novo cliente funcional em <4h |

## 3. Edge Cases

| Cenário | Condição | Resultado Esperado |
|---------|---------|-------------------|
| Cache Redis vazio | Primeiro acesso após restart do Redis | Fallback para consulta Keycloak; cache repopulado |
| Dois tenants mesmo domínio | Configuração incorreta | Erro na validação de unicidade; provisionamento bloqueado |
| Cookie corrompido | Cookie alterado manualmente | 401 — sessão inválida; redirecionar login |
| Concorrência de refresh | 2 chamadas simultâneas POST /auth/refresh | Ambas retornam 200; apenas 1 novo token válido |

## 4. Negative Test Cases

| Cenário | Ação | Resultado Esperado |
|---------|------|-------------------|
| Code OIDC reutilizado | /auth/callback com mesmo code 2x | 1ª OK, 2ª 401 (code já consumido) |
| Token expirado manual | POST /auth/me com cookie de token expirado | 401 |
| SQL Injection no redirect_uri | GET /auth/login?redirect_uri=javascript:alert(1) | 400 — input sanitizado |
| Header X-Tenant-Host forjado | Forjar header sem passar pela Cloudflare | Rejeitado — IP de origem não é Cloudflare |

## 5. Gherkin Scenarios

```gherkin
Feature: Login Multi-Tenant
  Como usuário de uma escola
  Quero fazer login no sistema da minha instituição
  Para acessar as funcionalidades do produto

  Scenario: Login bem-sucedido com reconhecimento automático de tenant
    Given que o domínio "escola-alfa.com" está mapeado para o Realm "realm-escola-alfa"
    When eu acesso "https://escola-alfa.com/app"
    Then sou redirecionado para a tela de login da Escola Alfa
    When informo minhas credenciais corretas
    Then sou redirecionado para a aplicação
    And minha sessão está ativa e protegida por cookies HttpOnly

  Scenario: Acesso entre tenants é bloqueado
    Given que estou logado na Escola Alfa
    When tento acessar dados da Escola Beta
    Then recebo uma lista vazia
    And nenhum dado da Escola Beta vazou

  Scenario: Tenant suspenso tem acesso bloqueado imediatamente
    Given que estou logado na Escola Alfa
    When a Escola Alfa é suspensa
    Then meu próximo acesso é bloqueado com erro 403
```

## 6. Traceability TC → FR(SRS) → Test Plan

| TC | FR (SRS) | Seção Test Plan | Status |
|----|---------|----------------|--------|
| TC-001, TC-002 | FR-01 — Extração de domínio | Functional/System | ✅ |
| TC-003, TC-004, TC-005 | FR-03, FR-04 — Login + Cookies | Functional/System + Security | ✅ |
| TC-006, TC-007, TC-008, TC-009, TC-010 | FR-05, FR-06, FR-07 — Sessão | Functional/System | ✅ |
| TC-011, TC-012 | FR-08 — Isolamento | Integration + Security | ✅ |
| TC-013, TC-014 | FR-12 — Suspensão | Functional/System | ✅ |

---

**[STATUS: SUCESSO]** — 14 casos de teste, 4 happy paths, 4 edge cases, 4 negative cases, 3 Gherkin scenarios. Cobertura: 7/12 FRs com TCs diretos.
