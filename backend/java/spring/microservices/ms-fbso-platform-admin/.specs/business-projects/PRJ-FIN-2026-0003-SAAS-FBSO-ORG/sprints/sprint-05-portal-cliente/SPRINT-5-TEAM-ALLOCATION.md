# SPRINT-5-TEAM-ALLOCATION.md — Matriz de Alocação Técnica

- **Sprint:** 5 de 7 — Portal do Cliente e Onboarding
- **Base:** [TECHNICAL-TEAM-MAP.md](../../../../../../../../business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/TECHNICAL-TEAM-MAP.md) v1.5
- **Status time:** ⚠️ 10 papéis — todos aguardando designação (`<TODO>`)
- **Risco R2 (Project Charter):** 🔴 Crítico — Equipe reduzida não consegue entregar no prazo
- **Data:** 22 de Julho de 2026
- **Versão:** 1.0

---

## 1. Capacidade Instalada (11 dias úteis — 15/09 a 30/09/2026)

| Papel | h/dia | Total Sprint (88h max) | Foco na Sprint 5 |
|:---|:---:|:---:|:---|
| Tech Lead / Líder Técnico | 8h | 88h | Code review, mentoria, decisões técnicas, gestão dívida técnica |
| Desenvolvedor Full-Stack | 8h | 88h | OnboardingService (T-060), DashboardClientService (T-063) |
| Desenvolvedor Frontend | 8h | 88h | **4 tasks frontend** (T-157..T-160) — Next.js, onboarding wizard, dashboard UI |
| Desenvolvedor Backend | 8h | 88h | Auth endpoints (T-058, T-059), OnboardingController (T-061), DashboardController (T-064), /auth/me (T-066), Frentes 1-2 |
| QA / Test Engineer | 8h | 88h | 38 cenários de teste (33 + 5 UI), T-067, T-068, T-147 |
| Analista de Homologação (Negócio) | 8h | 88h | Validação RN13/RN14, cenários onboarding, massa de dados |
| Desenvolvedor Banco de Dados | 8h | 88h | Keyset pagination (T-139), migration V007 (T-142), UTC timestamps (T-151), BU Matriz (T-062) |
| Arquiteto de Solução | 4h | 44h | Diagramas estado (T-143, T-145), ADRs, revisão arquitetural |
| DevOps Engineer | 4h | 44h | springdoc bump (T-152), CORS (T-153), CI/CD pipeline |
| Especialista IAM / Keycloak | 4h | 44h | 🔑 **Caminho crítico:** T-057 (Keycloak realm), T-065 (JWT claims), T-148 (JwtConverter) |
| **Total** | **68h/d** | **748h** | |

---

## 2. Matriz RACI — 40 Tasks × 10 Papéis

**Legenda:** R = Responsible (executa) | A = Accountable (aprova) | C = Consulted | I = Informed

### 2.1 Frente 0 — Bloqueantes (✅ Concluída)

| Task | IAM | DevBack | DevFull | DevFront | DevBD | QA | Homolog | Arq | DevOps | TechLead |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| T-133..T-138 | I | R | I | I | I | I | I | C | C | A |

> Frente 0 executada por Agente IA. DevOps consultado para docker-compose (T-133). Arquiteto consultado para SecurityConfig (T-137).

### 2.2 Frente 1 — Recomendados (10 tasks)

| Task | DT | IAM | DevBack | DevFull | DevFront | DevBD | QA | Homolog | Arq | DevOps | TechLead |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| T-139 keyset | DT-023 | — | C | — | — | **R** | I | — | — | — | A |
| T-140 corrigir contagem | DT-097 | — | I | — | — | — | I | — | — | — | **R** |
| T-141 TenantContext→JWT | DT-098 | C | I | — | — | — | — | — | — | — | A |
| T-142 migration V007 | DT-107 | — | C | — | — | **R** | I | — | C | — | A |
| T-143 máquina estados | DT-108 | — | C | C | — | — | — | — | **R** | — | A |
| T-144 rate limit Filter | DT-110 | — | **R** | — | — | — | I | — | C | — | A |
| T-145 diagrama onboarding | DT-124 | — | C | C | C | — | — | — | **R** | — | A |
| T-146 ExceptionHandler 401 | DT-121 | C | **R** | — | — | — | I | — | — | — | A |
| T-147 +5 cenários teste | DT-106 | I | I | — | — | — | **R** | C | — | — | A |
| T-148 JwtConverter | DT-102 | **R** | C | — | — | — | — | — | — | — | A |

### 2.3 Frente 2 — Desejáveis (8 tasks)

| Task | DT | IAM | DevBack | DevFull | DevFront | DevBD | QA | Homolog | Arq | DevOps | TechLead |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| T-149 AuditFieldsRowMapper | DT-086 | — | **R** | — | — | C | — | — | — | — | A |
| T-150 ObjectMapper injetável | DT-089 | — | **R** | — | — | — | — | — | — | — | A |
| T-151 UTC OffsetDateTime | DT-090 | — | C | — | — | **R** | — | — | — | — | A |
| T-152 springdoc bump | DT-092 | — | I | — | — | — | — | — | — | **R** | A |
| T-153 CORS externalizado | DT-093 | — | C | — | — | — | — | — | C | **R** | A |
| T-154 atualizar riscos | DT-101 | — | I | — | — | — | — | — | — | — | **R** |
| T-155 header SPECS.md | DT-112 | — | — | — | — | — | — | — | — | — | **R** |
| T-156 recalcular progresso | DT-113 | — | — | — | — | — | — | — | — | — | **R** |

### 2.4 Frente 3a — Backend (12 tasks)

| Task | Feature | IAM | DevBack | DevFull | DevFront | DevBD | QA | Homolog | Arq | DevOps | TechLead |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| T-057 Keycloak realm | F04-01 | **R** | C | I | I | — | I | — | C | C | A |
| T-058 Auth endpoints | F04-01 | C | **R** | I | C | — | I | — | — | — | A |
| T-059 Rate limiting | F04-01 | — | **R** | — | — | — | I | — | C | — | A |
| T-060 OnboardingService | F04-02 | — | C | **R** | C | C | I | C | C | — | A |
| T-061 OnboardingController | F04-02 | — | **R** | C | C | — | I | — | — | — | A |
| T-062 BU Matriz | F04-02 | — | **R** | C | — | **R** | I | — | — | — | A |
| T-063 DashboardClientService | F04-03 | — | C | **R** | C | C | I | — | — | — | A |
| T-064 DashboardController | F04-03 | — | **R** | — | C | — | I | — | — | — | A |
| T-065 JWT claims→TenantContext | F04-04 | **R** | C | — | I | — | I | — | — | — | A |
| T-066 GET /auth/me | F04-04 | C | **R** | — | C | — | I | — | — | — | A |
| T-067 Testes unitários M5 | QA | I | C | C | — | — | **R** | C | — | — | A |
| T-068 Testes integração M5 | QA | I | C | C | — | C | **R** | C | — | C | A |

### 2.5 Frente 3b — Frontend (4 tasks) 🆕

| Task | Feature | IAM | DevBack | DevFull | DevFront | DevBD | QA | Homolog | Arq | DevOps | TechLead |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| T-157 Bootstrap Next.js | F04-01 | I | C | C | **R** | — | I | — | C | C | A |
| T-158 Auth UI | F04-01 | C | C | — | **R** | — | I | — | — | — | A |
| T-159 Onboarding wizard UI | F04-02 | — | C | C | **R** | — | I | C | C | — | A |
| T-160 Dashboard cliente UI | F04-03 | — | C | C | **R** | — | I | — | — | — | A |

---

## 3. Cronograma Dia a Dia (Plano de Ocupação)

### Semana 1 (15/09 — 19/09): Setup + Auth + Início Onboarding

| Dia | IAM (4h) | DevBack (8h) | DevFull (8h) | DevFront (8h) | DevBD (8h) | QA (8h) | Homolog (8h) | Arq (4h) | DevOps (4h) | TechLead (8h) |
|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|
| **Seg 15** | T-057 Keycloak realm | T-144 rate limit | T-060 OnboardingService | T-157 Bootstrap Next.js | T-139 keyset | T-147 cenários teste | Escrever cenários RN14 | T-143 máquina estados | T-152 springdoc | Coordenação + review |
| **Ter 16** | T-057 (cont.) | T-146 ExceptionHandler | T-060 (cont.) | T-157 (cont.) | T-142 migration V007 | T-147 (cont.) | Validar realm-config.json | T-145 diagrama onboarding | T-153 CORS | Code review PRs |
| **Qua 17** | T-148 JwtConverter | T-058 Auth endpoints | T-060 (cont.) | T-158 Auth UI | T-062 (apoio BU Matriz) | Revisar suite teste | Massa dados onboarding | Revisar ADRs | CI/CD pipeline | Revisão arquitetural |
| **Qui 18** | T-065 JWT claims | T-058 (cont.) | T-063 DashboardService | T-158 (cont.) | T-151 UTC timestamps | T-067 (início) | Testar fluxo onboarding | Revisão segurança | Monitor staging | Planning Sprint 6 |
| **Sex 19** | T-065 (cont.) | T-059 Rate limiting | T-063 (cont.) | T-159 Onboarding wizard | Apoio T-062 | T-067 (cont.) | Validar dashboard | Folga/review | Deploy staging | Retrospectiva semana |

**Milestone Fim Semana 1:** Keycloak realm pronto ✅ · Auth endpoints funcionais ✅ · Onboarding wizard UI iniciado · Dashboard service em andamento

### Semana 2 (22/09 — 30/09): Onboarding + Dashboard + Testes

| Dia | IAM (4h) | DevBack (8h) | DevFull (8h) | DevFront (8h) | DevBD (8h) | QA (8h) | Homolog (8h) | Arq (4h) | DevOps (4h) | TechLead (8h) |
|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|
| **Seg 22** | Suporte Keycloak | T-061 OnboardingController | Finalizar T-063 | T-159 (cont.) | Frente 2 (T-149) | T-067 (cont.) | Validar onboarding | Revisar PRs | Monitor staging | Code review |
| **Ter 23** | T-148 (finalizar) | T-061 (cont.) | T-064 DashboardController | T-159 (finalizar) | Frente 2 (T-150) | T-068 integração | Homologar auth | Documentar ADRs | Logs + alertas | Revisão quality gate |
| **Qua 24** | Suporte claims JWT | T-064 (cont.) | Apoio QA | T-160 Dashboard UI | Apoio T-068 | T-068 (cont.) | Homologar onboarding | Revisão final | Backup + restore | Demo parcial PO |
| **Qui 25** | Validação final OAuth2 | T-066 /auth/me | Correções PO | T-160 (finalizar) | Correções DB | T-068 (finalizar) | UAT onboarding | Aceite arquitetural | Smoke test staging | Correções UAT |
| **Sex 26** | Documentar realm | Correções UAT | Correções UAT | Testes responsivos | Verificar índices | Suite regressão | UAT dashboard | — | Tag release | Review final |
| **Seg 29** | — | Deploy HML | Deploy HML | Deploy HML | — | Testes E2E | Homologação final | — | Deploy produção | Aceite formal |
| **Ter 30** | — | Plantão | Plantão | Plantão | — | Relatório QA | Relatório homolog | — | Monitor pós-deploy | **Fim Sprint 5** ✅ |

**Milestone Fim Semana 2:** Onboarding 100% funcional ✅ · Dashboard cliente operacional ✅ · Testes passando ✅ · PO aprovou ✅ · Deploy HML ✅

---

## 4. Caminho Crítico

```
T-057 (IAM, 2d) ──→ T-058 (DevBack, 2d) ──→ T-065 (IAM, 1.5d) ──→ T-066 (DevBack, 1d)
        ↓                                            ↓
   [bloqueia T-058]                           [bloqueia T-066]
```

**Duração do caminho crítico:** 6.5 dias-homem (≈ 8 dias wall-clock com IAM 4h/dia)

**Riscos do caminho crítico:**
- IAM Specialist com apenas 4h/dia: T-057 (16h = 4 dias) + T-065 (12h = 3 dias) = 7 dias úteis só de IAM
- Se T-057 atrasar, toda a cadeia auth atrasa (T-058, T-065, T-066)
- Frontend auth UI (T-158) depende de T-057 concluído para testar redirect

**Mitigação:** Iniciar T-057 durante a Frente 1 (antes de 15/09). Considerar bump do IAM para 8h/dia durante a Sprint 5.

---

## 5. Riscos de Alocação

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| 🔴 Time não preenchido — 10 `<TODO>` | Alta | Crítico | Preencher nomes antes de 15/09. Se impossível, reduzir escopo (cortar Frente 3b frontend, manter apenas backend) |
| 🟡 IAM Specialist 4h/dia insuficiente | Média | Alto | Iniciar T-057 na Frente 1. Considerar bump 4h→8h na Sprint 5 |
| 🟡 Frontend sem design system | Média | Médio | Tailwind CSS cobre tokens. MSW mock permite dev paralelo. Evoluir design system na Sprint 6 |
| 🟡 Dev Frontend sem backend real | Média | Médio | MSW mock handlers (T-157) simulam todas as APIs. Testes integrados só na Semana 2 |
| 🟢 QA ocioso na Semana 1 | Baixa | Baixo | QA escreve cenários e prepara suite antes do código ficar pronto (shift-left testing) |
| 🟢 Analista Homologação sem tasks formais | Baixa | Baixo | Atribuir validação RN13/RN14, massa dados, revisão realm-config.json |

---

## 6. Métricas de Ocupação

| Papel | Tasks Atribuídas | Carga Estimada | % Ocupação | Folga |
|:---|:---|---:|:---:|:---|
| Tech Lead | 6 (T-140, T-154, T-155, T-156 + review) | ~30h | 34% | ⚠️ Subutilizado — pode absorver tasks de outros papéis |
| Dev Full-Stack | 3 (T-060, T-063, apoio) | ~40h | 45% | Pode apoiar Dev Backend em picos |
| Dev Frontend | 4 (T-157..T-160) | ~64h | 73% | Alocado exclusivamente ao frontend |
| Dev Backend | 8 (T-058, T-059, T-061, T-064, T-066, T-144, T-146, T-149, T-150) | ~60h | 68% | Boa ocupação. Cuidado com dependências do IAM |
| QA Engineer | 4 (T-067, T-068, T-147 + regressão) | ~52h | 59% | Shift-left: escrever cenários antes do código |
| Analista Homologação | 0 tasks formais | ~20h | 23% | ⚠️ Subutilizado — atribuir validação RN13/RN14 formal |
| Dev BD | 5 (T-139, T-142, T-062 apoio, T-149, T-151) | ~32h | 36% | Pode absorver mais tasks de migração/índices |
| Arquiteto | 3 (T-143, T-145, revisões) | ~24h | 55% | 4h/dia bem dimensionado para escopo |
| DevOps | 3 (T-152, T-153, CI/CD) | ~16h | 36% | 4h/dia adequado. Pode apoiar deploy staging |
| Especialista IAM | 3 (T-057, T-065, T-148) | ~28h | 64% | ⚠️ 4h/dia é gargalo. Carga real: 28h = 7 dias wall-clock |
| **Total** | **40 tasks** | **~366h** | **49% média** | **382h de folga total** |

> ⚠️ **Alerta de Subutilização:** Tech Lead (34%), Analista Homologação (23%) e Dev BD (36%) têm folga significativa. Em caso de ausência de profissionais preenchidos, esses papéis podem absorver tasks de Dev Backend ou QA.

---

## 7. Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 22/07/2026 | Criação inicial: matriz RACI 40 tasks × 10 papéis, cronograma dia a dia, caminho crítico, métricas de ocupação. Base: TECHNICAL-TEAM-MAP.md v1.5 + SPRINT-CARD.md revisado. | Agente IA |

---

🤖 *Documento gerado como parte da reavaliação da Sprint 5 com time técnico v1.5. Todos os nomes de profissionais estão como `<TODO>` — preencher antes do início da Frente 3 (15/09/2026).*
