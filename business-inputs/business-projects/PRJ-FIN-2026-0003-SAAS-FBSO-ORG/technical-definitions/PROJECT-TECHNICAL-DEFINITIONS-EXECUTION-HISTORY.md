# Histórico de Execução — Definições Técnicas do Projeto

- **Baseline de Negócio:** [Project Charter v1.2](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BRD v1.2](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Épicos v1.2](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Features FEAT-EP-](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 2.0
- **Data de Criação:** 25 de Julho de 2026
- **Última Atualização:** 27 de Julho de 2026 (alinhamento com docs de negócio v1.2)

---

## Dashboard de Estado dos Documentos

| # | Documento | Fase | Estado | Data | Versão | Observações |
|:---|:---|:---:|:---|:---|:---:|:---|
| 0 | Bootstrap | 0 | ✅ COMPLIANCE | 25/07 | — | Estrutura criada, TEAM-CAPACITY migrado |
| 1 | TEAM-MAP | 1 | ✅ COMPLIANCE | 26/07 | v1.3 | 11 perfis, 9 categorias, 36+ tecnologias |
| 2 | SOLUTIONS-CATALOG | 2 | ✅ COMPLIANCE | 26/07 | v1.2 | 14 soluções (2 apps + 12 infra) |
| 3 | STACK-MATRIX | 3 | ✅ COMPLIANCE | 26/07 | v1.1 | 14 stacks com versões, 13 ADRs (5 técnicos + 8 negócio/stack) |
| 4 | PRD-DEFINITION | 4 | ✅ COMPLIANCE | 26/07 | v1.0 | 18 features, 58 US, MVP 52 US, ERD migrado do TECHNICAL-PLAN |
| 5 | ARCHITECTURE-DEFINITION | 5 | ✅ COMPLIANCE | 26/07 | v1.1 | C4 L1+L2, 6 seq diagrams, 5 ADRs integração, matriz épico, error handling |
| 6 | SECURITY-DEFINITION | 6 | ✅ COMPLIANCE | 26/07 | v1.1 | 12 ameaças, IAM, DevSecOps, LGPD |
| 7 | SPECS-DEFINITION | 7 | ✅ COMPLIANCE | 26/07 | v1.0 | Convenções Java/React/DB/API/Logging, 11 recursos API, matriz RBAC, pacotes backend, rotas frontend |
| 8 | MILESTONES | 8 | ✅ COMPLIANCE | 26/07 | v1.0 | 7 milestones M1-M7 |
| 9 | SOLUTIONS-MATRIX | 9 | ✅ COMPLIANCE | 26/07 | v1.0 | Tabela-mestra 14 soluções, RACI |
| 10 | Estrutura de Sprints | 10 | ✅ COMPLIANCE | 26/07 | — | 11 sprints (00 a 10) |
| 11 | EXECUTION-HISTORY | 11 | ✅ FINAL | 27/07 | v3.0 | Roadmap completo + registro de consolidação |
| 12 | REPOSITORY-STRUCTURE | 12 | ✅ COMPLIANCE | 27/07 | v1.0 | Árvore de diretórios alvo, convenções `.specs/` |

**Resumo: 12/12 fases concluídas. 10 documentos técnicos gerados. 0 fases pendentes.**

### Consolidação de Documentos (Fase 13 implícita — 27/07/2026)

Os 4 documentos de nível de projeto foram consolidados na pasta `technical-definitions/` e **removidos**:

| Documento Removido | Conteúdo Migrado Para |
|:---|:---|
| `TECHNICAL-PLAN.md` | ERD → PRD-DEFINITION §2.4 · Dir Tree → REPOSITORY-STRUCTURE · Stack/Milestones/Riscos já cobertos por STACK-MATRIX + MILESTONES |
| `ARCHITECTURE.md` | ADRs → STACK-MATRIX §5 · Pacotes/rotas → SPECS-DEFINITION §6.1-6.2 · C4/Auth/Multi-Tenant já cobertos por ARCHITECTURE-DEFINITION |
| `INTEGRATION-MAP.md` | Totalmente absorvido por ARCHITECTURE-DEFINITION (matriz épico §4.3, error handling §6.4) |
| `API-CONTRACTS.md` | 11 recursos + matriz RBAC → SPECS-DEFINITION §3.6-3.7 · Inconsistências resolvidas (camelCase, flat errors, 5min JWT) |
| `TECHNICAL-TEAM-MAP.md` | Duplicata do [TEAM-MAP](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md) (v1.4, mais completo). Removido por redundância. |

---

## Timeline de Execução

```
25/07/2026  ████████████████████████████████████████  Fases 0-2 (Bootstrap + Bloco A parcial)
26/07/2026  ████████████████████████████████████████  Fases 3-11 (Bloco A final → Bloco D)
            │
            ├─▶ Bloco A: People & Solutions ✅
            │     ✅ F1 (TEAM-MAP) → ✅ F2 (CATALOG) → ✅ F3 (STACK) → ✅ F4 (PRD)
            │
            ├─▶ Bloco B: Architecture & Security ✅
            │     ✅ F5 (ARCHITECTURE) → ✅ F6 (SECURITY)
            │
            ├─▶ Bloco C: Specs & Milestones ✅
            │     ✅ F7 (SPECS) → ✅ F8 (MILESTONES)
            │
            └─▶ Bloco D: Matriz, Sprints, Histórico ✅
                  ✅ F9 (MATRIX) → ✅ F10 (SPRINTS) → ✅ F11 (HISTORY)
```

---

## Estrutura Final de Diretórios

```
technical-definitions/
├── PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md      ← este documento
├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md          ← Fase 0
├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md ← Fase 0
├── PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md              ← Fase 1
├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md     ← Fase 2
├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md ← Fase 3
├── PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md        ← Fase 4
├── PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md ← Fase 5
├── PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md   ← Fase 6
├── PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md      ← Fase 7
├── PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md            ← Fase 8
├── PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md      ← Fase 9
└── PROJECT-TECHNICAL-DEFINITIONS-REPOSITORY-STRUCTURE.md ← Fase 12
    ├── sprint-00/    ← Setup & Fundação
    ├── sprint-01/    ← EP-0001 (Dashboard + Contas)
    ├── sprint-02/    ← EP-0001 + EP-0002 início
    ├── sprint-03/    ← EP-0002 (Clientes e Assinaturas)
    ├── sprint-04/    ← EP-0002 final + EP-0003 início
    ├── sprint-05/    ← EP-0003 (RBAC)
    ├── sprint-06/    ← EP-0003 final + EP-04a início
    ├── sprint-07/    ← EP-04a (Portal Cliente)
    ├── sprint-08/    ← EP-04a final + EP-04b início
    ├── sprint-09/    ← EP-04b (BUs + Catálogo)
    └── sprint-10/    ← Homologação final (M7)
```

---

## Principais Decisões Técnicas Registradas

| Decisão | Fase | Detalhe |
|:---|:---:|:---|
| Multi-Tenant RLS + FORCE | 3 | `FORCE ROW LEVEL SECURITY` + `tenant_isolation_policy` via `current_setting('app.current_tenant_id')` |
| OIDC via Kong Gateway | 3, 5 | Authorization Code + PKCE. Kong valida JWT (JWKS local). Backend confia no Gateway. |
| Schemas DB: public / fbso_portal / keycloak | 3 | `public` (DB), `fbso_portal` (RLS app), `keycloak` (Realms) |
| Cloudflare + DigitalOcean | 3, 5 | CDN/WAF na Cloudflare Edge → Kong → DOKS. Custom Hostnames white-label. |
| Keycloak Realms por tenant | 3, 4 | Isolamento via Realms. Personalização de login (logo, cores, fontes). |
| Soft Delete + RLS integrado | 3 | `pg_has_role()` para admin bypass. Expurgo definitivo fora do escopo. |
| Kong como camada de auth | 5, 6 | Sem double-check no backend. Kong+Keycloak = auth centralizada. |
| GraalVM Native Image | 3 | AOT compilation. Fallback JVM via `Dockerfile.jvm`. |

---

## Registro de Decisões e Desvios

| Data | Decisão | Impacto |
|:---|:---|:---|
| 25/07/2026 | Escopo: 2 apps (backend existente + frontend futuro) | 14 soluções totais |
| 25/07/2026 | Bolismar cobre frontend até 01/11 (Tom Santos) | 🔴 Risco M5-M6 |
| 25/07/2026 | Maria Madalena junior ★☆☆ global | 🔴 Plano de mentoria Sprint 0-4 |
| 26/07/2026 | S14 Kong API Gateway adicionado | Auth centralizada, 14 soluções |
| 26/07/2026 | S12 migrado Hostinger → DOKS Secrets | DigitalOcean como plataforma única |

---

## Histórico de Alterações

| Data | Alteração | Responsável |
|:---|:---|:---|
| 25/07/2026 | Criação do dashboard — Fase 0 concluída | Time de Arquitetura |
| 25/07/2026 | Fase 1 concluída — TEAM-MAP (v1.1 → v1.2 → v1.3) | Time de Arquitetura |
| 25/07/2026 | Fase 2 concluída — SOLUTIONS-CATALOG (v1.0 → v1.1 → v1.2) | Time de Arquitetura |
| 26/07/2026 | Fase 3 concluída — STACK-MATRIX (v1.0 → v1.1) | Time de Arquitetura |
| 26/07/2026 | Fase 4 concluída — PRD-DEFINITION (v1.0) | Time de Arquitetura |
| 26/07/2026 | Fase 5 concluída — ARCHITECTURE-DEFINITION (v1.0 → v1.1) | Time de Arquitetura |
| 26/07/2026 | Fase 6 concluída — SECURITY-DEFINITION (v1.0 → v1.1) | Time de Arquitetura |
| 26/07/2026 | Fase 7 concluída — SPECS-DEFINITION (v1.0) | Time de Arquitetura |
| 26/07/2026 | Fase 8 concluída — MILESTONES (v1.0) | Time de Arquitetura |
| 26/07/2026 | Fase 9 concluída — SOLUTIONS-MATRIX (v1.0) | Time de Arquitetura |
| 26/07/2026 | Fase 10 concluída — 11 sprints criadas (00-10) | Time de Arquitetura |
| 26/07/2026 | **Fase 11 FINAL — Roadmap de Definições Técnicas concluído** | Time de Arquitetura |
| 26/07/2026 | Revisão Mermaid global: todos os documentos com ASCII art convertidos para Mermaid (flowchart, gantt, sequence). SOLUTIONS-CATALOG v1.3, STACK-MATRIX v1.2, ARCHITECTURE v1.2, MILESTONES v1.1. | Time de Arquitetura |
| 27/07/2026 | **Fase 12 — REPOSITORY-STRUCTURE criado** (v1.0): árvore de diretórios alvo, convenções `.specs/`, matriz de localização de artefatos | Time de Arquitetura |
| 27/07/2026 | **Fase 13 (implícita) — Consolidação de documentos de projeto**: TECHNICAL-PLAN.md (§2.4 ERD → PRD-DEFINITION, §4.3 Dir Tree → REPOSITORY-STRUCTURE), ARCHITECTURE.md (ADRs → STACK-MATRIX, pacotes/rotas → SPECS-DEFINITION), INTEGRATION-MAP.md (matriz épico/error handling → ARCHITECTURE-DEFINITION), API-CONTRACTS.md (11 recursos + RBAC → SPECS-DEFINITION). 4 arquivos deletados por obsolescência. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Roadmap de 12 fases concluído em 3 dias (25-27 Julho 2026). 10 documentos técnicos gerados, 14 soluções catalogadas, 13 ADRs documentados, 4 documentos de projeto consolidados e removidos.*
