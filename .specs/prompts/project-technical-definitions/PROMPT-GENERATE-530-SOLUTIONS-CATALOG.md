# PROMPT-GENERATE-530-SOLUTIONS-CATALOG

## Contexto

Este prompt gera o artefato `530-SOLUTIONS-CATALOG.md` — o **catálogo exaustivo de soluções técnicas** do projeto. Cataloga serviços, bancos de dados, integrações, filas/mensageria, processos ETL, data services, frontends, mobile apps — toda solução técnica necessária para entregar o escopo do projeto.

**Inputs upstream (Bloco C — F13):** Este artefato recebe insumos cumulativos de todos os blocos anteriores:
- **Inputs Globais do Roadmap:** `PROJECT_PATH`, `PROJECT_ID_NAME`, `TECHNICAL_SOLUTION_PATH`, `TECHNICAL_SOLUTION_NAMES`, `ARCHITECTURE_GLOBAL`, `SECURITY_GLOBAL`
- **Bloco 0 (Product Def & Backlog & PRD):** `410-INTAKE-LOG.md`, `420-DOR-ASSESSMENT.md`, `430-PRODUCT-BACKLOG-LIST.md`, `440-PRD-DEFINITION.md`
- **Bloco A (People & Solutions):** `450-TEAM-SKILLS-MAP.md` (Discovery Team skills) + `460-TEAM-CAPACITY.md` (disponibilidade e horas)
- **Bloco B (6 Disciplinas Técnicas):** `470-ARCHITECTURE-DEFINITION.md` (padrões técnicos, C4, ADRs) + `480-SECURITY-DEFINITION.md` (threat model, IAM, compliance) + `490-DATA-ARCHITECTURE-DEFINITION.md` (modelagem, storage, pipelines) + `500-DEVOPS-SRE-DEFINITION.md` (CI/CD, IaC, observabilidade, SLOs) + `510-TEST-STRATEGY-DEFINITION.md` (pirâmide de testes, automação) + `520-INFRA-CLOUD-DEFINITION.md` (topologia, networking, DR)
- **Documentos de Negócio:** Charter, BRD, Epics, Features (para escopo e requisitos)

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler documentos de negócio (Charter, BRD, Epics, Features), TECHNICAL-PLAN.md (referência) e TEAM-SKILLS-MAP.md (para atribuição).

### Passo 2 — Invocar Skills Especializadas
Invocar skills para analisar domínio, identificar soluções necessárias, classificar por tipo e mapear relação com épicos/features.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/530-SOLUTIONS-CATALOG.md` com:
- Lista exaustiva de soluções (nome, descrição, propósito)
- Classificação por tipo (backend, frontend, mobile, batch, integração, banco, infra)
- Relação com épicos/features do projeto de negócio
- Propriedade (time/papel responsável)
- Estado atual (existente, a criar, planejado)
- Prioridade (MoSCoW)

### Passo 4 — Validação Pós-Geração
Verificar: arquivo no caminho correto, todas as soluções catalogadas, classificação consistente, relação com negócio documentada.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `architecture-patterns` | Identificar padrões arquiteturais das soluções | Arquitetura |
| 2 | `system-design` | Projetar sistema completo a partir dos docs de negócio | Arquitetura |
| 3 | `domain-driven-design` | Análise de domínio para identificar bounded contexts | DDD |
| 4 | `domain-analysis` | Identificar domínios e subdomínios | Domínio |
| 5 | `microservices-architect` | Definir granularidade de serviços | Arquitetura |
| 6 | `backend-architect` | Identificar soluções de backend necessárias | Especialidade |
| 7 | `frontend-dev-guidelines` | Identificar soluções de frontend necessárias | Especialidade |
| 8 | `data-engineering-data-pipeline` | Identificar pipelines e ETL necessários | Dados |
| 9 | `database-architect` | Identificar bancos de dados necessários | Dados |
| 10 | `documentation-writer` | Redigir o catálogo consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador do catálogo de soluções técnicas | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
