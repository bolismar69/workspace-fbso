# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` — o **catálogo exaustivo de soluções técnicas** do projeto. Cataloga serviços, bancos de dados, integrações, filas/mensageria, processos ETL, data services, frontends, mobile apps — toda solução técnica necessária para entregar o escopo do projeto.

**Inputs upstream:** Documentos de negócio (Charter, BRD, Epics, Features) + TECHNICAL-PLAN.md (referência) + PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md (para atribuição de responsabilidades).

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio | `/home/user/work/business-inputs/business-projects` |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions | Derivado |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler documentos de negócio (Charter, BRD, Epics, Features), TECHNICAL-PLAN.md (referência) e TEAM-MAP.md (para atribuição).

### Passo 2 — Invocar Skills Especializadas
Invocar skills para analisar domínio, identificar soluções necessárias, classificar por tipo e mapear relação com épicos/features.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` com:
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
