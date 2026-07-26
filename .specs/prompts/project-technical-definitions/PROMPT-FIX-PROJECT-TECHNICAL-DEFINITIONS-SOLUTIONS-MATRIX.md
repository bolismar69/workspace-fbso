# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md`. O agente corretor aplica correções cirúrgicas com base no relatório de falha.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório de Falha e Artefatos
Ler `SOLUTIONS_MATRIX_SCOPE_FAIL_REPORT.md`, a matriz atual, Catálogo de Soluções, Stack Matrix, TEAM-CAPACITY.md.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Solução ausente da matriz | Adicionar linha completa para a solução |
| P0 | Stack inconsistente | Alinhar com Stack Matrix (Fase 3) |
| P1 | Campo obrigatório vazio | Preencher com dado da fase correspondente |
| P2 | Indicador não calculado | Calcular cobertura de skills e risco |
| P3 | RACI incompleto | Completar matriz RACI |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `reference-builder` | Corrigir referências cruzadas | Mapeamento |
| 3 | `team-composition-analysis` | Corrigir alocação de time | People |
| 4 | `documentation-writer` | Atualizar matriz | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da matriz de soluções | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
