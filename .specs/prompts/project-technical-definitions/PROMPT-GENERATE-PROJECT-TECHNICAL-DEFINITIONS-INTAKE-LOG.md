# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` 🆕 — o **registro de lotes de ingestão de requisitos** do Negócio para a TI. Documenta cada onda/lote de requisitos recebido, sua origem, data, responsável e escopo.

**Regra de versionamento:** Cada novo lote de ingestão recebe um número de versão (`v1`, `v2`, ...). Em projetos Waterfall, tipicamente há um único lote. Em projetos Ágeis (Scrum/Kanban/OKR), múltiplas ondas são registradas ao longo do ciclo de vida do projeto.

**Papel no Bloco 0 (Product Definition & Product Backlog & PRD):** Fase 1 de 4. Este artefato é o ponto de entrada formal dos requisitos de negócio no pipeline técnico. Ele alimenta o DoR Assessment (F2) e o Product Backlog List (F3).

**Inputs upstream:** Documentos de negócio — Charter, BRD, Epics, Features, User Stories.

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
Verificar se TODOS os parâmetros obrigatórios foram informados.

### Passo 1 — Carregar Documentos Base
Ler TODOS os documentos de negócio do projeto: índices principais (`01-PROJECT-CHARTER-*.md`, `02-BRD-*.md`, `03-EPICS-*.md`, `04-FEATURES-*.md`, `05-USER-STORIES-*.md`) + documentos complementares (`DEFINITION_OF_DONE.md`, `MATRIZ-KPI.md`, `STAKEHOLDER-MAP.md`, `GLOSSARY.md`) + TODOS os arquivos individuais nas subpastas `epics/*.md`, `features/*.md` e `user-stories/*.md` para identificar o escopo completo de cada lote de ingestão.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de análise de negócio, product management e stakeholder analysis para estruturar os lotes.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` com:

1. **Registro de Lotes** — tabela: Versão, Data, Origem (PO/PM/Stakeholder), Tipo (Waterfall/Ágil), Escopo (descrição), Status
2. **Detalhamento por Lote** — para cada lote: lista de documentos de negócio associados, User Stories/Features/Epics incluídos
3. **Matriz de Cobertura** — gráfico/matriz de quais documentos de negócio foram cobertos em qual lote
4. **Histórico de Alterações** — changelog do próprio INTAKE-LOG

### Passo 4 — Validação Pós-Geração
Verificar: lotes versionados sequencialmente, rastreabilidade com docs de negócio, cobertura completa de 100% dos documentos.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `business-analyst` | Análise dos documentos de negócio para identificar escopo | Análise |
| 2 | `product-manager` | Visão de produto e ondas de entrega | Produto |
| 3 | `stakeholder-analysis` | Identificação de stakeholders por lote | Stakeholder |
| 4 | `requirements-engineering` | Engenharia de requisitos para estruturação dos lotes | Requisitos |
| 5 | `documentation-writer` | Redigir o INTAKE-LOG consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador do registro de lotes de ingestão (F1 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
