# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX

## Contexto

Este prompt implementa o **Gate de Validação da Matriz de Soluções** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md`. Verifica se a matriz-mestra está completa, consistente e todas as referências cruzadas são válidas.

**Princípio fundamental:** Toda solução do catálogo deve ter: responsável, repositório, stack, time e status definidos na matriz. Nenhum campo pode ficar vazio.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md`, Catálogo de Soluções, Stack Matrix, TEAM-CAPACITY.md, Milestones.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Completude
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Soluções cobertas | Toda solução do catálogo está na matriz |
| 1.2 | Campos preenchidos | Tipo, Repositório, Stack, Tech Lead — sem vazios |
| 1.3 | Perfis definidos | Perfis necessários listados para cada solução |

#### Dimensão 2: Consistência de Referências
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Stack consistente | Stack na matriz = Stack Matrix (Fase 3) |
| 2.2 | Responsáveis existem | Tech Leads referenciados existem no TEAM-CAPACITY |
| 2.3 | Status alinhado | Status reflete milestones (Fase 8) |

#### Dimensão 3: Indicadores
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Cobertura de skills | % calculado para cada solução |
| 3.2 | Riscos de gargalo | Sinalizados quando capacidade < necessária |
| 3.3 | RACI preenchido | Matriz RACI com Responsável e Autoridade definidos |

### Passo 3 — Calcular Veredito
### Passo 4 — Gerar Relatório de Falha (se REPROVADO)

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Identificar campos vazios | Análise |
| 2 | `team-composition-analysis` | Validar alocação de time | People |
| 3 | `project-manager` | Validar viabilidade da matriz | PM |
| 4 | `reference-builder` | Validar referências cruzadas | Mapeamento |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da matriz de soluções | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
