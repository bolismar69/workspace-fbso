# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório de falha.

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
Ler `PRD_DEFINITION_SCOPE_FAIL_REPORT.md`, o PRD Definition atual, documentos de negócio e catálogo de soluções.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Requisito órfão (sem solução) | Mapear para solução existente ou justificar exclusão |
| P1 | Seção obrigatória ausente | Preencher seção completa |
| P2 | MVP Global vago | Detalhar escopo com critérios mensuráveis |
| P3 | Termo inconsistente | Alinhar com glossário do BRD |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções
100% P0 resolvidas, cobertura 100%, 8 seções completas.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `requirements-engineering` | Corrigir requisitos faltantes | Requirements |
| 3 | `prd-development` | Corrigir seções do PRD | Product |
| 4 | `documentation-writer` | Atualizar PRD Definition | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção do PRD Definition | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
