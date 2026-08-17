# PROMPT-GENERATE-420-DOR-ASSESSMENT

## Contexto

Este prompt gera o artefato `420-DOR-ASSESSMENT.md` 🆕 — a **avaliação de Definition of Ready (DoR)** dos requisitos recebidos. Aplica um checklist formal a cada item do INTAKE-LOG para determinar se está suficientemente refinado para seguir para o backlog técnico.

**Critérios DoR avaliados por item:**
- O requisito é **claro e compreensível**?
- É **testável** (critérios de aceitação definidos)?
- Está **priorizado**?
- **Dependências** estão mapeadas?
- Possui **estimativa inicial**?

**Papel no Bloco 0 (Product Definition & Product Backlog & PRD):** Fase 2 de 4. Este artefato consome o INTAKE-LOG (F1) e alimenta o Product Backlog List (F3). Somente itens aprovados no DoR progridem para o backlog.

**Inputs upstream:** INTAKE-LOG (F1), documentos de negócio (Charter, BRD, Epics, Features, User Stories).

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
Ler `410-INTAKE-LOG.md` (F1) e TODOS os documentos de negócio do projeto: índices principais + TODOS os arquivos individuais nas subpastas `epics/*.md`, `features/*.md` e `user-stories/*.md` para ter a relação completa de itens a avaliar no DoR.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de análise de negócio, engenharia de requisitos, validação e critérios de aceitação.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/420-DOR-ASSESSMENT.md` com:

1. **Checklist DoR de Negócio** — para cada requisito: claro? testável? priorizado? dependências mapeadas? estimado?
2. **Itens Aprovados** — requisitos que passaram no DoR (status ✅), prontos para o backlog (F3)
3. **Itens Devolvidos** — requisitos que voltam para refinamento com justificativa (status 🔄), incluindo o critério não atendido
4. **Matriz de Pendências** — para cada item devolvido: o que falta, ação necessária, responsável pelo refinamento
5. **Assinatura PO/PM** — registro de aprovação formal com campo para nome e data

### Passo 4 — Validação Pós-Geração
Verificar: 100% dos itens do INTAKE-LOG avaliados, itens devolvidos têm justificativa clara, itens aprovados atendem todos os critérios DoR.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `business-analyst` | Análise de clareza e completude dos requisitos | Análise |
| 2 | `requirements-engineering` | Aplicação dos critérios DoR aos itens | Requisitos |
| 3 | `requirements-validation` | Validação de cada item contra o checklist | Requisitos |
| 4 | `acceptance-criteria` | Verificação de testabilidade e critérios de aceitação | Qualidade |
| 5 | `gap-analysis` | Identificação de pendências nos itens devolvidos | Análise |
| 6 | `documentation-writer` | Redigir o DOR-ASSESSMENT consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da avaliação DoR (F2 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
