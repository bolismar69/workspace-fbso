# Casos de Teste: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 03-SRS, 08-TEST-PLAN |
| **Solução Técnica** | ms-shield-identity-auth |
| **Data** | 03/08/2026 | **Versão** | 2.0 — Revisão Integração | **Metodologia** | WATERFALL |

---

## 1. Test Case Catalog

| ID | Cenário | Precondition | Steps | Expected Result | Priority |
|----|---------|-------------|-------|----------------|----------|
| TC-001 | **SPA carrega sem sessão → API interceptada → redirect Keycloak** | Redis com mapeamento `escola-alfa.com` → `realm-escola-alfa`. Zero cookies | 1. SPA faz GET /api/v1/alunos (sem cookie) 2. Kong encaminha para Shield | 302 → Keycloak /realms/realm-escola-alfa/auth | High |
| TC-002 | **Host não mapeado → 401** | Redis sem mapeamento para `desconhecido.com` | 1. SPA faz GET /api/v1/alunos com header Host: desconhecido.com | 401 — "Domínio não configurado", sem detalhes internos | High |
| TC-003 | **Login completo via Keycloak → cookie setado → JWT no Redis** | Tenant ativo. Zero cookies | 1. Seguir redirect 302 do TC-001 2. Login no Keycloak 3. Callback processado pelo Shield | 302 → escola-alfa.com. Cookie SHIELD_SESSION setado: HttpOnly, Secure, SameSite=Strict. JWT armazenado no Redis | High |
| TC-004 | **Cookie válido → Kong+Shield injetam JWT → MS recebe Authorization** | Sessão ativa (TC-003) | 1. SPA faz GET /api/v1/alunos (com cookie SHIELD_SESSION) 2. Kong → Shield valida 3. Shield recupera JWT do Redis → injeta header 4. MS recebe requisição | MS recebe GET /api/v1/alunos com header `Authorization: Bearer <JWT>`. Claims: tenant_id, roles, user_id | High |
| TC-005 | **Frontend NUNCA vê o JWT** | Sessão ativa | 1. No browser: `document.cookie` 2. No browser: inspecionar todas as responses de API | `document.cookie` não retorna SHIELD_SESSION (HttpOnly). Nenhuma response contém JWT no body ou headers | Critical |
| TC-006 | **Cookie expirado → refresh OK → novo JWT** | Sessão com access_token expirado, refresh_token válido | 1. SPA faz GET /api/v1/alunos (cookie com sessão expirada) 2. Shield tenta refresh | 200. Novo JWT armazenado no Redis. MS recebe Authorization header com novo token | Medium |
| TC-007 | **Cookie expirado → refresh falhou → redirect Keycloak** | Sessão totalmente expirada (>30min) | 1. SPA faz GET /api/v1/alunos (cookie expirado) 2. Shield tenta refresh → Keycloak retorna 400 | 302 → Keycloak /auth. Cookie de sessão removido | Medium |
| TC-008 | **Logout → sessão removida do Redis** | Sessão ativa | 1. SPA remove cookie local 2. Próxima chamada API → Shield detecta cookie ausente → 302 Keycloak. Sessão anterior expira no Redis por TTL | Usuário redirecionado para login. JWT anterior inacessível | High |
| TC-009 | **Cross-Tenant Block — MS recebe JWT tenant A, query dados tenant B** | Tenant A logado, dados Tenant B no banco | 1. Shield injeta JWT com tenant_id=A 2. MS faz GET /api/v1/dados?escola=B 3. MS executa SET LOCAL app.current_tenant = 'A' 4. PostgreSQL RLS filtra | 200 {data: []} — 0 linhas retornadas. Sem vazamento | Critical |
| TC-010 | **Same-Tenant OK — MS recebe JWT tenant A, query dados tenant A** | Tenant A logado | 1. Shield injeta JWT tenant_id=A 2. MS faz GET /api/v1/dados?escola=A | 200 {data: [...]} — dados do tenant A retornados | Critical |
| TC-011 | **JWT forjado/tampered → Kong rejeita** | — | 1. Enviar GET /api/v1/alunos com header `Authorization: Bearer <jwt_invalido>` (sem passar pelo Shield) | 401 — assinatura JWT inválida. Kong rejeita antes de chegar ao MS | High |
| TC-012 | **Tenant suspenso → Shield bloqueia** | Tenant `escola-alfa` suspenso | 1. SPA com cookie SHIELD_SESSION do tenant suspenso 2. Shield valida → tenant_id está suspenso | 403 — acesso bloqueado. Cookie invalidado | High |
| TC-013 | **KEDA escala Shield sob carga** | 200+ req/s simultâneas | 1. Disparar 200+ req/s para /api/v1/alunos com cookies válidos 2. Observar KEDA ScaledObject | Shield escala 2→N pods. Zero erros 5xx. Latência p95 <15ms | Medium |
| TC-014 | **Redis indisponível → Shield degrada com cache local** | Redis fora do ar | 1. Derrubar Redis 2. SPA faz GET /api/v1/alunos com cookie válido | Shield valida JWT localmente (cache em memória, TTL curto). Alerta Prometheus disparado. 200 OK | Medium |

---

## 2. Happy Path Cases

| Fluxo | Gatilho | Resultado Esperado |
|-------|---------|-------------------|
| Acesso normal | SPA faz chamada API. Kong+Shield injetam JWT. MS responde | Usuário não percebe autenticação — transparente |
| Primeiro acesso | SPA chamada API sem cookie → redirect Keycloak → login → callback → cookie setado | Usuário vê tela de login uma vez. Depois sessão automática |
| Renovação silenciosa | Cookie com JWT expirado → Shield faz refresh → novo JWT no Redis | Sessão estendida sem intervenção do usuário |

## 3. Edge Cases

| Cenário | Condição | Resultado Esperado |
|---------|---------|-------------------|
| Cache Redis vazio | Primeiro acesso após restart | Host resolvido via fallback Keycloak; cache repopulado |
| Cookie corrompido | SHIELD_SESSION alterado manualmente | 302 → Keycloak login (cookie inválido = sem sessão) |
| Concorrência de refresh | 2 chamadas simultâneas com cookie expirado | Ambas recebem novo JWT; apenas 1 session_id prevalece |
| SPA em múltiplas abas | 2 abas da mesma escola | Compartilham cookie SHIELD_SESSION — ambas funcionam |

## 4. Negative Test Cases

| Cenário | Ação | Resultado Esperado |
|---------|------|-------------------|
| Header Authorization forjado | Enviar JWT falso diretamente para MS (bypass Kong) | Topologia impede — MS só recebe tráfego do Kong (network policy) |
| Tentativa de ler cookie via JS | `document.cookie` no console | Cookie SHIELD_SESSION não aparece (HttpOnly) |
| Replay de callback OIDC | /auth/callback com mesmo code 2x | 1ª OK, 2ª 401 (code já consumido) |
| SQL Injection no host header | Header X-Tenant-Host com SQL injection | Input sanitizado antes da consulta Redis |

## 5. Gherkin Scenarios

```gherkin
Feature: Autenticação Transparente via Kong+Shield
  Como SPA frontend
  Quero fazer chamadas de API normalmente
  Para que a autenticação seja gerenciada pelo Kong+Shield sem meu conhecimento

  Scenario: Primeira visita — API interceptada, login via Keycloak
    Given que não tenho cookie SHIELD_SESSION
    When faço GET /api/v1/alunos
    Then recebo 302 para a tela de login do Keycloak da minha escola
    When autentico com sucesso
    Then recebo cookie SHIELD_SESSION HttpOnly
    And minhas próximas chamadas API recebem os dados normalmente

  Scenario: Sessão ativa — JWT injetado transparentemente
    Given que tenho cookie SHIELD_SESSION válido
    When faço GET /api/v1/alunos
    Then recebo 200 com os dados da minha escola
    And em nenhum momento vejo o JWT no frontend

  Scenario: Isolamento entre tenants
    Given que estou logado na Escola Alfa
    When o MS consulta dados da Escola Beta
    Then o PostgreSQL RLS retorna 0 linhas
    And nenhum dado da Escola Beta vazou
```

## 6. Traceability TC → SRS → Test Plan

| TC | Componente Shield (LLD) | Seção Test Plan | Status |
|----|------------------------|----------------|--------|
| TC-001, TC-002 | SessionFilter, TenantResolver | Functional/System | ✅ |
| TC-003, TC-004, TC-005 | SessionFilter, SessionStore, JWTInjector | Functional + Security | ✅ |
| TC-006, TC-007 | SessionFilter, KeycloakClient | Integration | ✅ |
| TC-008 | SessionStore.delete() | Functional | ✅ |
| TC-009, TC-010 | PostgreSQL RLS (via MS) | Integration + Security | ✅ |
| TC-011 | Kong JWTInjector | Security | ✅ |
| TC-012 | SessionFilter.validate() | Functional | ✅ |
| TC-013 | KEDA ScaledObject | Performance | ✅ |
| TC-014 | SessionStore (fallback local) | Integration | ✅ |

---

**[STATUS: SUCESSO]** — 14 casos de teste alinhados com arquitetura Kong Filter. Nenhum TC chama /auth/* diretamente. TC-005 verifica propriedade crítica: frontend nunca vê o JWT.
