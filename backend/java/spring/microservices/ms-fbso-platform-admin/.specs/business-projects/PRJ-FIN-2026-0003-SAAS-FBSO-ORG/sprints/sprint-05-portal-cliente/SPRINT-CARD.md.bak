# SPRINT-CARD: Sprint 5 — Portal do Cliente e Onboarding

- **Sprint:** 5 de 7
- **Marco:** M5 (EP-04a)
- **Datas:** 15/09/2026 → 30/09/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [SPECS.md](../../SPECS.md)

---

> 🚫 **BRANCH OBRIGATÓRIA:** Toda implementação deste sprint DEVE usar exclusivamente a branch `feature/sprint-05-portal-cliente`. Antes de começar, execute:
> ```bash
> git checkout feature/sprint-05-portal-cliente
> git branch --show-current  # deve exibir: feature/sprint-05-portal-cliente
> ```
> 📖 Detalhes completos: [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

## 🎯 Sprint Goal

**"Login via Keycloak OIDC funcional com recuperação de senha e rate limiting. Onboarding guiado em 4 passos — primeira BU vira Matriz, tenant transita para ACTIVE ao concluir. Dashboard do cliente com cards de unidades, produtos e plano. Endpoint /auth/me retorna dados do usuário logado."**

---

## 📋 Sprint Backlog

| ID | Tarefa | Feature | Prio. | Est. | Critério DONE |
|:---|:---|:---|:---:|:---:|:---|
| **T-050** | Configurar Keycloak realm `fbso-platform` + client Authorization Code Flow. Mapear claims (tenant_id, roles, business_unit_ids, modules). Exportar `realm-config.json` | F04-01 | Must | 2d | Realm funcional. JWT com claims corretas. Config versionada |
| **T-051** | `POST /api/v1/auth/login` (delega Keycloak). Recuperação: `POST /auth/forgot-password` (expira 1h — RN13-03), `POST /auth/reset-password` (complexidade RN13-01) | F04-01 | Must | 2d | Login funcional. Link reset 1h. Senha 8+ chars, letra+número. Sessão 60min inatividade (RN13-02) |
| **T-052** | Rate limiting: 5 tentativas → bloqueio 15min (US-039, RN13-02). Implementar via @Aspect ou filter + Caffeine | F04-01 | Must | 1.5d | 5 falhas → bloqueado. Mensagem exibe tempo restante |
| **T-053** | `OnboardingService`: 4 passos obrigatórios (US-040 a US-044). Ordem obrigatória (RN14-01). Primeira BU = Matriz (RN14-02). Tenant → ACTIVE ao concluir (RN14-04). Retomável | F04-02 | Must | 2d | 4 passos. Ordem validada. Status salvo permite retomar |
| **T-054** | `OnboardingController`: `GET /onboarding/status`, `PATCH /step-1`, `POST /step-2`, `POST /complete`. `@RequiresPermission` | F04-02 | Must | 1.5d | 4 endpoints. Step-2 valida CNPJ. Complete só se todos passos OK (RN14-03) |
| **T-055** | Criação primeira BU como Matriz (step-2). Reutilizar `BusinessUnitService.create()` com flag `isMatrix=true`. CNPJ validado | F04-02 | Must | 1d | BU Matriz com parent_id=NULL. CNPJ válido |
| **T-056** | `DashboardClientService`: cards — unidades ativas, produtos catálogo, plano contratado, notificações. Cards clicáveis | F04-03 | Should | 1.5d | Cards com dados resumidos. Link para área |
| **T-057** | `DashboardClientController`: `GET /dashboard/client/summary`, `/notifications`. `@RequiresPermission`. Dashboard genérico Fase 0 (RN15-01, RN15-02) | F04-03 | Should | 1d | Dados do cliente autenticado. Filtrado por tenant_id |
| **T-058** | Atualizar `JwtAuthenticationFilter`: claims `modules[]` + `business_unit_ids[]` no `TenantContext`. Placeholder "FBSO Platform" (RN16-01, RN16-02) | F04-04 | Must | 1.5d | JWT retorna modules[]. App Switcher viável com 1 módulo |
| **T-059** | `GET /api/v1/auth/me`: retorna dados do usuário logado (id, name, email, role, business_unit_ids, modules[], tenant_id, status onboarding). Sem `@RequiresPermission` | F04-04 | Must | 1d | GET /auth/me funcional. Stateless (dados do token). 401 se sem token |
| **T-060** | Testes unitários M5: `OnboardingService`, `DashboardClientService`, fluxo auth (login, recovery, rate-limit) | F04-01 a F04-04 | Must | 1.5d | ≥ 80%. Rate limit testado (5 tentativas, bloqueio, reset 15min) |
| **T-061** | Testes integração M5: fluxo onboarding completo (PENDING→ACTIVE). Rate limiting PostgreSQL real. Teste expiração link reset | F04-01, F04-02 | Must | 2d | Onboarding completo testado. Tenant → ACTIVE |

**Total:** 12 tarefas · ~18.5 dias-homem (10 Must + 2 Should)

---

## 📦 Features Entregues

| Feature | Descrição | RNs Cobertas | Prio. |
|:---|:---|:---|:---:|
| **F04-01** | Login e Autenticação | RN13-01, RN13-02, RN13-03 | Must |
| **F04-02** | Onboarding Guiado | RN14-01, RN14-02, RN14-03, RN14-04 | Must |
| **F04-03** | Dashboard do Cliente | RN15-01, RN15-02 | Should |
| **F04-04** | App Switcher | RN16-01, RN16-02 | Must |

---

## ✅ Definition of Done (Sprint-Level)

- [ ] Login Keycloak funcional (Authorization Code Flow)
- [ ] Recuperação de senha com link (expira 1h)
- [ ] Rate limiting: 5 tentativas → bloqueio 15min
- [ ] Onboarding 4 passos: ordem validada, não pula etapas
- [ ] Primeira BU vira Matriz (parent_id=NULL)
- [ ] Tenant transita PENDING_ONBOARDING → ACTIVE ao concluir
- [ ] Dashboard cliente: cards com unidades, produtos, plano
- [ ] `GET /auth/me` retorna dados do token (stateless)
- [ ] JWT inclui claims `modules[]` e `business_unit_ids[]`

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Keycloak não disponível para desenvolvimento local | Média | Crítico | Keycloak container no Docker Compose. Realm exportado como `realm-config.json` |
| Rate limiting com Caffeine não funcionar em cluster (stateless) | Baixa | Médio | Fase 0: single instance. Fase 1: migrar para Redis se necessário |
| Onboarding com estado inconsistente (passo 2 falha, passo 1 salvo) | Média | Alto | `@Transactional` no `OnboardingService`. Rollback se qualquer passo falhar |

---

## 🔗 Dependências

- **Pré-requisitos:** Sprint 2 (JwtAuthenticationFilter). Sprint 3 (TenantRepository). Sprint 4 (UserRepository). Keycloak provisionado.
- **Sucessor:** Sprint 6 (BUs e Catálogo) — depende de BusinessUnit criada durante onboarding.

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Tasks completadas | 12/12 |
| Must Have | 10 |
| Should Have | 2 (F04-03) |
| Cenários de teste | 21 |

---

🤖 *Gerado a partir de TASKS.md v2.0. Keycloak é dependência crítica — provisionar antes de começar.*
