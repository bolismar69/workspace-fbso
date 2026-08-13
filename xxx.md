# Plano: Criar Projeto SHIELD no Jira + Confluence (Waterfall)

## Contexto

O projeto PRJ-TEC-2026-0004-PROJETO-SHIELD é um microserviço IAM multi-tenant (Auth-BFF) com metodologia **Waterfall** documentada em 20 documentos oficiais + 7 artefatos de estimativa. O projeto já passou por estimativas UPSTREAM (ROM ±50%) e DOWNSTREAM (PERT ±15%), está em fase de GO/NO-GO.

## Ferramentas Disponíveis

| Ferramenta | Status | Uso |
|-----------|--------|-----|
| **MCP Jira** `read:jira-work`, `write:jira-work` | ✅ | Criar projeto, issues, links |
| **MCP Confluence** `read/write pages` | ✅ | Upload da documentação |
| **REST API + Token** | ✅ | Fallback para operações não cobertas pelo MCP |

---

## Decisões de Design (6 confirmadas)

1. **Granularidade Alta** — Epic→Story→Task→Subtask
2. **Workflow com Homologação** — `Backlog → To Do → In Progress → Code Review → Homologação → Done`
3. **Stories como Entregáveis** — Epic(Fase) → Story(Entregável WBS) → Task(Pacote WBS)
4. **Kanban + Fix Versions** — Board Kanban com filtro por fase + milestones via Versions
5. **Dependências entre Epics** — `blocks` entre fases + Fix Versions
6. **7 Fases Waterfall** — incluindo fases pré-desenvolvimento

---

## PARTE 1: Estrutura Confluence

### 1.1 Localização

**NÃO criar novo espaço.** Usar o espaço existente `NEPF` (Negócio e Produto Funcional), seguindo o padrão já estabelecido:

```
📁 Projetos de Negócio (folder 229526)
  ├── 📁 PRJ-FIN-2026-0003-SAAS-FBSO-ORG (existente)
  └── 📁 PRJ-TEC-2026-0004-PROJETO-SHIELD (NOVO)
        └── (27 páginas organizadas por fase Waterfall)
```

### 1.2 Árvore de Páginas (27 páginas, 7 pastas)

```
📁 PRJ-TEC-2026-0004-PROJETO-SHIELD/
│
├── 📄 Overview do Projeto
│   (resumo executivo, links Jira/GitHub, status GO/NO-GO)
│
├── 📁 F1-Negócios & Discovery/
│   ├── 📄 01 - Project Charter v2.0
│   ├── 📄 02 - BRD v2.0
│   ├── 📄 03 - SRS v1.0
│   ├── 📄 04 - RTM v1.0
│   ├── 📄 05 - SAD v2.0
│   └── 📄 06 - HLD v2.0
│
├── 📁 F2-Estimativa UPSTREAM/
│   ├── 📄 ROM ±50%
│   ├── 📄 Scope Snapshot UPSTREAM
│   └── 📄 Governance ROM Report (GO/NO-GO preliminar)
│
├── 📁 F3-Detalhamento Técnico/
│   ├── 📄 07 - LLD v3.0
│   ├── 📄 08 - Test Plan v3.0
│   ├── 📄 09 - Test Cases v2.0
│   ├── 📄 10 - Relatório de Qualidade
│   └── 📄 11 - EAP/WBS
│
├── 📁 F4-Estimativa DOWNSTREAM/
│   ├── 📄 PERT ±15%
│   ├── 📄 Scope Snapshot DOWNSTREAM
│   ├── 📄 Cronograma Calculado (PERT)
│   └── 📄 Orçamento Calculado (PERT)
│
├── 📁 F5-GO/NO-GO & Planejamento/
│   ├── 📄 Decisão GO/NO-GO (comitê)
│   ├── 📄 12 - Cronograma Gantt
│   └── 📄 13 - Orçamento
│
├── 📁 F6-Execução Técnica/
│   ├── 📄 14 - Plano de Comunicação
│   ├── 📄 15 - Plano de Riscos
│   ├── 📄 16 - Deployment Plan v2.0
│   ├── 📄 17 - Manuais de Usuário
│   ├── 📄 18 - Manuais Operacionais
│   ├── 📄 19 - Termo de Aceite
│   └── 📄 20 - Lições Aprendidas
│
└── 📁 ADRs/
    ├── 📄 ADR-01: GraalVM Native
    ├── 📄 ADR-02: BFF como único cliente Keycloak
    ├── 📄 ADR-03: Istio + Kong
    ├── 📄 ADR-04: Redis Cache
    ├── 📄 ADR-05: PostgreSQL RLS
    ├── 📄 ADR-06: GitOps (Argo CD)
    └── 📄 ADR-07: Cookies HttpOnly
```

### 1.3 Integração Confluence → Jira

Cada página de documento no Confluence terá na descrição um link para o Epic/Task correspondente no Jira:
```
🔗 Jira: [SHIELD-XX] Nome da Issue
```

Cada Epic/Story/Task no Jira terá na descrição um link para a pasta ou página no Confluence:
```
📚 Docs: https://bolismar69.atlassian.net/wiki/spaces/NEPF/pages/...
```

---

## PARTE 2: Estrutura Jira

### 2.1 Projeto

- **Nome:** `PRJ-TEC-2026-0004 - PROJETO SHIELD`
- **Key:** `SHIELD`
- **Tipo:** `software` (company-managed, template `jira-core-project-management`)
- **Lead:** FRAN Oliveira (712020:bed45ac6-4894-4b1d-909c-2a27c5c4653a)

### 2.2 Workflow Customizado (6 estados)

```
Backlog → To Do → In Progress → Code Review → Homologação → Done
              ↑                                            │
              └────────── Rejeitado ←──────────────────────┘
```

### 2.3 Epics (8 Fases Waterfall)

| # | Epic | Fix Version | Bloqueia | Conteúdo |
|---|------|-------------|----------|----------|
| F0 | Governança & Setup | V0-Setup | — | Kickoff, cerimônias, reports |
| F1 | Negócios & Discovery | V1-Discovery | F2 | Docs 01→06 |
| F2 | Estimativa UPSTREAM | V2-Upstream | F3 | ROM ±50% + Scope + Governance |
| F3 | Detalhamento Técnico | V3-Detalhamento | F4 | Docs 07→11 |
| F4 | Estimativa DOWNSTREAM | V4-Downstream | F5 | PERT ±15% + Scope + Cronograma + Orçamento |
| F5 | GO/NO-GO & Planejamento | V5-Planejamento | F6 | Decisão + Docs 12-13 |
| F6 | Execução Técnica | V6-Execucao | F7 | WBS D1→D7 (7 Stories, 23 Tasks) |
| F7 | Documentação & Encerramento | V7-Encerramento | — | Docs 14→20 |

### 2.4 Fix Versions (8)

`V0-Setup`, `V1-Discovery`, `V2-Upstream`, `V3-Detalhamento`, `V4-Downstream`, `V5-Planejamento`, `V6-Execucao`, `V7-Encerramento`

### 2.5 Components (7 — apenas F6)

`infra-producao`, `motor-identidade`, `portal-acesso`, `isolamento-dados`, `observabilidade`, `homologacao-seguranca`, `go-live`

### 2.6 Issues da F6 (Execução Técnica)

| Story (D#) | Component | Tasks |
|-------------|-----------|-------|
| D1: Infraestrutura Produção | infra-producao | 1.1.1 DOKS+Istio, 1.1.2 CI/CD+ArgoCD |
| D2: Motor de Identidade | motor-identidade | 1.2.1 Keycloak, 1.2.2 Kong+Shield, 1.2.3 Cloudflare |
| D3: Portal de Acesso | portal-acesso | 1.3.1 BFF Login, 1.3.2 Session, 1.3.3 Redis |
| D4: Isolamento de Dados | isolamento-dados | 1.4.1 RLS, 1.4.2 Cross-Tenant Tests |
| D5: Observabilidade | observabilidade | 1.5.1 Grafana, 1.5.2 Loki/Jaeger |
| D6: Homologação Segurança | homologacao-seguranca | 1.6.1→1.6.6 (6 Tasks) |
| D7: Go-Live | go-live | 1.7.1→1.7.5 (5 Tasks) |
| Riscos Críticos | — | R01 GraalVM, R02 Keycloak, R03 Cross-Tenant, R08 LGPD |

**Total Jira: ~55-60 issues**

### 2.7 Dependências (Issue Links)

```
F0 ──blocks──→ F1 ──blocks──→ F2 ──blocks──→ F3 ──blocks──→ F4 ──blocks──→ F5 ──blocks──→ F6 ──blocks──→ F7
```

Tasks **GATE** ao final de cada fase passam por `Homologação` → `Done` para liberar o bloqueio.

---

## PARTE 3: Implementação (via MCP)

### Passo 1: Confluence — Criar pasta do projeto
- Criar folder `PRJ-TEC-2026-0004-PROJETO-SHIELD` dentro de `Projetos de Negócio` (229526)
- Criar página Overview com resumo + links

### Passo 2: Confluence — Criar pastas por fase (7)
- Criar folders F1 a F7 + ADRs dentro da pasta do projeto

### Passo 3: Confluence — Upload dos documentos (27 páginas)
- Para cada arquivo `.md` na pasta do projeto, criar página no Confluence
- Conteúdo convertido de Markdown para storage format (HTML)
- Cada página linka para o Jira correspondente

### Passo 4: Jira — Criar projeto `SHIELD`
- Usar `createJiraIssue` com projectTypeKey `software`

### Passo 5: Jira — Criar Components e Fix Versions
- 7 Components + 8 Fix Versions

### Passo 6: Jira — Criar Epics com dependências
- 8 Epics com descrições linkando Confluence
- `createIssueLink` para configurar `blocks` entre fases

### Passo 7: Jira — Criar Stories e Tasks
- 11 Stories + 23 Tasks dentro dos Epics
- Cada issue com link para doc correspondente no Confluence

### Passo 8: Jira — Configurar workflow
- Customizar para 6 estados com transições

### Passo 9: Verificação
- Board Kanban com 6 colunas
- 8 Epics em ordem com dependências
- 27 páginas no Confluence sob NEPF > Projetos de Negócio > SHIELD
- Links bidirecionais Jira ↔ Confluence

---

## Estimativa

| Passo | Itens | Tempo |
|-------|-------|-------|
| 1-3: Confluence (pastas + docs) | 1 folder + 7 subfolders + 27 páginas | 25-30 min |
| 4-5: Jira setup (projeto + components + versions) | 1 projeto + 15 itens | 5 min |
| 6-7: Jira issues (Epics + Stories + Tasks + links) | ~55 issues + links | 25-30 min |
| 8-9: Workflow + verificação | 1 workflow | 10 min |
| **Total** | **~88 itens** | **~70 min** |

---

## Verificação Final

1. ✅ `https://bolismar69.atlassian.net/wiki/spaces/NEPF/folder/229526` → Projetos de Negócio contém pasta SHIELD
2. ✅ Dentro da pasta SHIELD: 7 subpastas + Overview + 27 páginas
3. ✅ `https://bolismar69.atlassian.net/jira/software/projects/SHIELD/boards` → Board Kanban 6 colunas
4. ✅ 8 Epics com dependências `blocks` entre fases
5. ✅ Filtrar por Fix Version V6-Execucao → 7 Stories + 23 Tasks
6. ✅ Links bidirecionais Jira ↔ Confluence funcionando

---
