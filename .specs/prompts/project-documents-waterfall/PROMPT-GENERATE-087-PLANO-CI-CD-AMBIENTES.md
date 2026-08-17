# PROMPT: GERADOR DE PLANO DE CI/CD E AMBIENTES (087-PLANO-CI-CD-AMBIENTES)
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Engenheiro DevOps Sênior, especializado em estratégia de branches, pipelines de integração contínua e gestão de ambientes, no contexto da metodologia WATERFALL.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 030-SAD, 035-HLD, 041-DEVOPS-SETUP, 044-INFRA-SETUP]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["deployment-pipeline-design", "cicd-automation-workflow-automate", "devops-engineer"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** o 041-DEVOPS-SETUP (pipelines DED-NN e ambientes) e o 044-INFRA-SETUP (topologia IDD-NN) — este plano operacionaliza o que os setups definiram; o 030-SAD/035-HLD fornecem as decisões e a topologia de origem
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Use o prefixo padronizado: **CICD-NN** (pipelines e configurações de ambiente)
6. Aplique a tabela VOCABULÁRIO WATERFALL abaixo em todo o documento
7. Ao final, retorne `{DOC_PATH}` confirmando a criação

## VOCABULÁRIO WATERFALL (obrigatório — não usar vocabulário ágil)

| Termo ágil (PROIBIDO) | Equivalente WATERFALL (usar) |
|---|---|
| Epic | Pacote de trabalho da EAP (060-EAP-WBS) |
| Feature | Funcionalidade `FEAT-NN` (010-FRD) |
| User Story | Caso de Uso `UC-NN` (010-FRD) |
| Definition of Ready (DoR) | GATE de COMPLIANCE do documento de origem |
| Sprint | Ciclo de entrega (FASE 5 — EXECUÇÃO E CONSTRUÇÃO) |
| Product Backlog | 088-PRODUCT-BACKLOG-LIST (FASE 4) |

## Template de Fallback (6 Seções)

```
# Plano de CI/CD e Ambientes: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 030-SAD, 035-HLD, 041-DEVOPS-SETUP, 044-INFRA-SETUP |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## Plano de CI/CD e Ambientes

O **Plano de CI/CD e Ambientes** operacionaliza a esteira DevOps (041) sobre a infraestrutura (044): estratégia de branches, pipelines por ambiente, proteções e automação de deploy. É o guia de execução técnica que o time da FASE 5 segue em cada ciclo de entrega.

### O que contém

- **Estratégia de Branches:** fluxo git, proteções e nomenclatura
- **Pipelines (CICD-NN):** estágios, gates e aprovações por ambiente
- **Ambientes:** DEV/QA/HMG/PROD com provisionamento (044) e proteções
- **Automação de Deploy:** alinhada ao 041 (IaC, rollback)

### Conexão com o Pipeline

- **UPSTREAM:** Consome pipelines do 041-DEVOPS-SETUP, topologia do 044-INFRA-SETUP e decisões do 030-SAD/035-HLD
- **DOWNSTREAM:** Alimenta 090-STRATEGIC-IMPLEMENTATION-AND-DEPLOYMENT-PLAN (estratégia de deploy em produção), 092-BACKLOG-KANBAN (execução dos ciclos) e 100-MANUAIS-OPERACIONAIS (runbooks de ambientes)

---

## 1. Estratégia de Branches

| Item | Regra |
|------|-------|
| Branch principal | {main/master protegida} |
| Branch por ciclo | {feature/ciclo-NN-<slug> — padrão do git workflow da FASE 5} |
| Merge | {PR obrigatório com code review (086) e CI verde} |
| Proteções | {aprovações mínimas, bloqueio de push direto} |

```mermaid
flowchart LR
    Main["main"] --> Feature["feature/ciclo-01-x"]
    Feature -->|"PR + CI verde + review"| Main
```

---

## 2. Pipelines (CICD-NN)

| ID | Pipeline | Estágios | Gates/Aprovações | Checagens do 086 |
|----|----------|----------|------------------|------------------|
| CICD-01 | CI — build e testes | {lint → test → build} | {automático} | {STD-NN automatizados} |
| CICD-02 | CD — deploy | {deploy → smoke test} | {aprovação manual por ambiente} | {SAST/SCA (043)} |

---

## 3. Ambientes

| Ambiente | Provisionamento (044-IDD) | Pipeline (CICD-NN) | Proteções |
|----------|---------------------------|--------------------|-----------|
| DEV | {IDD-NN} | CICD-02 | {deploy livre} |
| QA | {IDD-NN} | CICD-02 | {gate de testes (050)} |
| HMG/UAT | {IDD-NN} | CICD-02 | {aprovação de negócio} |
| PROD | {IDD-NN} | CICD-02 | {GMUD (090) obrigatória} |

---

## 4. Automação de Deploy

| Item | Especificação (alinhada ao 041) |
|------|----------------------------------|
| IaC | {módulos do 041/044} |
| Rollback | {estratégia automatizada} |
| Smoke Test | {checagens pós-deploy} |

---

## 5. Rastreabilidade

| Item CICD | Origem (030/035/041/044) | Consumidores Previstos | Status |
|-----------|---------------------------|------------------------|--------|
| CICD-01 | {DED-NN do 041} | 090, 092, 100 | ✅ Vinculado |

> **REGRA DE OURO:** Nenhum pipeline ou ambiente pode existir sem lastro no DEVOPS-SETUP (041) ou no INFRA-SETUP (044).

---

## 6. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Criação inicial a partir de SAD/HLD, DEVOPS-SETUP e INFRA-SETUP | Time de Engenharia |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se as 6 seções estiverem completas, a estratégia de branches cobrir proteções e nomenclatura, os pipelines cobrirem CI e CD com gates, os 4 ambientes estiverem definidos com provisionamento e proteções, e a rastreabilidade não tiver órfãos.
