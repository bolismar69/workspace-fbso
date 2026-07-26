# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md`. O agente corretor aplica correções cirúrgicas com base no relatório de falha, preservando seções aprovadas.

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
Ler `SOLUTIONS_CATALOG_SCOPE_FAIL_REPORT.md`, o catálogo atual e documentos de negócio.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Épico/feature sem solução | Adicionar solução(ões) para cobrir o gap |
| P1 | Solução sem classificação | Atribuir tipo e estado |
| P2 | Prioridade indefinida | Atribuir MoSCoW |
| P3 | Inconsistência com TECHNICAL-PLAN | Alinhar nomenclatura ou justificar divergência |

### Passo 3 — Aplicar Correções Cirúrgicas

### Passo 4 — Validar Correções
100% P0 resolvidas, cobertura completa de épicos/features, consistência restaurada.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `architecture-patterns` | Corrigir classificação de soluções | Arquitetura |
| 3 | `system-design` | Refinar design das soluções | Arquitetura |
| 4 | `documentation-writer` | Atualizar catálogo | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção do catálogo de soluções | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
