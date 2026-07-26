# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md`. O agente corretor aplica correções cirúrgicas com base no relatório de falha.

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
Ler `MILESTONES_SCOPE_FAIL_REPORT.md`, o documento atual, Project Charter, Catálogo de Soluções.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Marco do Charter sem milestone | Criar milestone técnico correspondente |
| P0 | Solução sem milestone de entrega | Adicionar solução a milestone existente |
| P1 | Dependência não documentada | Documentar ordem de construção |
| P2 | Risco não mapeado | Adicionar risco e mitigação |
| P3 | Critério de aceitação vago | Tornar mensurável |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `roadmap-planning` | Corrigir milestones | Roadmap |
| 3 | `risk-manager` | Adicionar riscos faltantes | Risco |
| 4 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção dos milestones | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
