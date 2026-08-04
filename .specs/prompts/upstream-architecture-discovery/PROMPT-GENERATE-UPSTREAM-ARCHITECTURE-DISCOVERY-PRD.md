# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD

## Contexto

> 📐 **Discovery-Level / Upstream Architecture / Análise de Viabilidade:** Este prompt gera a versão Discovery-Level (high-level) do PRD para análise de viabilidade e estimativa ROM 50%. Para a versão detalhada, usar `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md`.

Este prompt gera o artefato `DISCOVERY-LEVEL-PRD.md` 🆕 — o **PRD Discovery-Level** baseado nos Épicos definidos pelo Negócio. Diferente do PRD detalhado (F4 do tech-defs), este documento foca na visão macro do produto para embasar a análise de viabilidade técnica.

**Papel no Bloco 0 (Product Definition Discovery-Level):** Fase 1 de 1. Este artefato é o ponto de partida do Discovery Técnico. Ele consolida a visão do produto a partir dos Épicos e alimenta todas as 6 disciplinas do Bloco B.

**Inputs upstream:** *(documentos de negócio)*
1. `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` — Project Charter
2. `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` — Business Requirements Document
3. `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` — Índice de Épicos
4. `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` — Arquivos individuais de épicos

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Documentos brutos adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `TECHNICAL_SOLUTION_PATH={TECHNICAL_SOLUTION_PATH}` — base das soluções técnicas
- `TECHNICAL_SOLUTION_NAMES={TECHNICAL_SOLUTION_NAMES}` — soluções do projeto
- `ARCHITECTURE_GLOBAL={ARCHITECTURE_GLOBAL}` — arquitetura global
- `SECURITY_GLOBAL={SECURITY_GLOBAL}` — segurança global
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
Validar que `{PROJECT_COMPLETE_PATH_NAME}` existe e contém os documentos de negócio (Charter, BRD, Épicos).
Criar `{UPSTREAM_DISCOVERY_PATH}` se não existir.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes documentos de negócio:
1. `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` — Project Charter
2. `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` — BRD
3. `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` — Índice de Épicos
4. `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` — Épicos individuais
### Passo 2 — Invocar Skills
Invocar skills de produto, análise de negócio e arquitetura para extrair visão macro.
### Passo 3 — Gerar o Artefato
Gerar `DISCOVERY-LEVEL-PRD.md` com:
1. **Visão do Produto (Resumo)** — 1-2 parágrafos sobre o produto
2. **Épicos × Soluções (Macro)** — Matriz de quais soluções cada épico demanda
3. **MVP Macro** — O que é essencial para o primeiro delivery
4. **Restrições Conhecidas** — O que NÃO está no escopo ou é restrito
5. **Glossário Inicial** — Termos canônicos do domínio
---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-PRD.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-PRD.md`

```markdown
# DISCOVERY-LEVEL-PRD.md — PRD Discovery-Level
## Fase 1 — Bloco 0: Product Definition Discovery-Level

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Produto** | {NOME_DO_PRODUTO} |
| **Documento** | DISCOVERY-LEVEL-PRD-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | {DATA_ATUAL} |
| **Tipo** | Documento de Negócio — briefing executivo do PM/PO/Analista de Negócios para o time de TI |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Vinculados:**
- [`01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md`](../01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md) — Project Charter
- [`02-BRD-{PROJECT_ID_NAME}.md`](../02-BRD-{PROJECT_ID_NAME}.md) — Business Requirements Document
- [`03-EPICS-{PROJECT_ID_NAME}.md`](../03-EPICS-{PROJECT_ID_NAME}.md) — Índice de Épicos
- [`epics/*.md`](../epics/) — Épicos individuais

---

## 1. Visão do Produto (Resumo Executivo para o Time de TI)
- 1-3 parágrafos explicando o que é o produto, por que ele importa agora e a visão de longo prazo
- Zero citações técnicas — linguagem 100% negócio

## 2. Épicos × Soluções (Matriz Macro de Cobertura)
- Tabela: Épico | Objetivo de Negócio | Funcionalidades | Soluções Demandadas | Prioridade | Data-Alvo
- Diagrama de dependências entre épicos (formato árvore/ASCII)

## 3. MVP Macro — O Que é Essencial para o Primeiro Delivery
- Agrupado por blocos temáticos (ex: Bloco 1: Operação Interna, Bloco 2: Segurança, Bloco 3: Experiência)
- Tabela por funcionalidade: # | Funcionalidade | Épico | Por que é essencial
- Cronograma macro do MVP: Período | Marco | Entregas | Épicos

## 4. Restrições Conhecidas (Constraints)
- **4.1 Fora do Escopo Desta Fase (Out of Scope):** Tabela: Item | Justificativa | Quando será tratado
- **4.2 Restrições de Negócio:** Tabela: Restrição (R1, R2, ...) | Descrição | Impacto
- **4.3 Restrições de Mercado e Regulatórias:** Tabela: Restrição | Descrição

## 5. Glossário Inicial — Termos Canônicos do Domínio
- Tabela: Termo | Definição | Fonte (Charter §X, BRD §Y, EP-XXXX)
- Termos extraídos dos documentos de negócio vinculados

## 6. Público-Alvo e Personas
- Tabela: Persona | Descrição | Principais Necessidades | Épicos Relacionados

## 7. Métricas de Sucesso do Negócio
- Tabela: Métrica | Situação Atual | Meta | Épico Relacionado

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: PRD Discovery-Level para análise de viabilidade | Product Owner / Analista de Negócios |
```

### Passo 4 — Validação Pós-Geração

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `product-discovery` | Validar oportunidades de produto, mapear premissas, testar problem-solution fit | Discovery |
| 2 | `product-manager` | Visão de produto high-level | Produto |
| 3 | `business-analyst` | Análise dos épicos de negócio | Negócio |
| 4 | `senior-architect` | Alinhamento com soluções técnicas | Arquitetura |
| 5 | `requirements-engineering` | Estruturação de requisitos macro | Requisitos |
| 6 | `documentation-writer` | Redigir o PRD Discovery-Level | Documentação |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador do PRD Discovery-Level (F1 — Bloco 0) | Time de Arquitetura |

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` | Project Charter |
| 2 | `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` | Business Requirements Document |
| 3 | `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` | Índice de Épicos |
| 4 | `{PROJECT_COMPLETE_PATH_NAME}/epics/*.md` | Épicos individuais |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Upstream Architecture Discovery — Fase 1*
