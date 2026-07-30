# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md` 🆕 — o **backlog consolidado de requisitos "Pronto para TI"**. Lista todos os itens aprovados no DoR (F2) em ordem de prioridade, com rastreamento completo à origem nos documentos de negócio e agrupamento por lote de ingestão.

**Regras de negócio:**
- Apenas itens com status DoR ✅ (aprovado) ingressam no backlog
- Cada item deve linkar para sua origem (Charter/BRD/Epic/Feature/US)
- Priorização segue MoSCoW (Must have, Should have, Could have, Won't have)
- Itens são agrupados por lote de ingestão do INTAKE-LOG (F1)

**Papel no Bloco 0 (Product Definition & Product Backlog & PRD):** Fase 3 de 4. Este artefato consome o INTAKE-LOG (F1) e o DOR-ASSESSMENT (F2) e alimenta o PRD (F4). É a ponte entre a definição de negócio e a especificação técnica.

**Inputs upstream:** INTAKE-LOG (F1), DOR-ASSESSMENT (F2), documentos de negócio (Charter, BRD, Epics, Features, User Stories).

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
Ler `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` (F1 — lotes), `PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md` (F2 — itens aprovados) e documentos de negócio (Charter, BRD, Epics, Features, User Stories) para referência de origem.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de product management, backlog management e priorização de requisitos.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md` com:

1. **Backlog Consolidado** — tabela: ID, Item, Origem (Charter/BRD/Epic/Feature/US), Prioridade (MoSCoW), Status DoR, Lote de Ingestão
2. **Rastreabilidade** — cada item com link markdown válido para sua origem nos documentos de negócio
3. **Resumo por Prioridade** — contagem e percentual por categoria MoSCoW (Must, Should, Could, Won't)
4. **Resumo por Lote** — itens agrupados por lote de ingestão do INTAKE-LOG, com contagem por lote

### Passo 4 — Validação Pós-Geração
Verificar: 100% dos itens aprovados no DoR presentes, rastreabilidade completa (item → origem), priorização MoSCoW aplicada, links markdown válidos.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `product-manager` | Visão de produto e priorização MoSCoW | Produto |
| 2 | `backlog-management` | Estruturação e organização do backlog | Gestão |
| 3 | `requirements-prioritization` | Aplicação da priorização MoSCoW | Requisitos |
| 4 | `business-analyst` | Validação de rastreabilidade item→origem | Análise |
| 5 | `documentation-writer` | Redigir o PRODUCT-BACKLOG-LIST consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da lista de backlog de produto (F3 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
