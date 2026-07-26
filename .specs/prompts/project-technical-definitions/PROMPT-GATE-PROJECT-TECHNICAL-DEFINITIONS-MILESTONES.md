# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-MILESTONES

## Contexto

Este prompt implementa o **Gate de Validação dos Milestones** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md`. Verifica se o roadmap técnico cobre todas as soluções, está alinhado com o Project Charter e respeita as dependências entre soluções.

**Princípio fundamental:** Todo milestone do Project Charter deve ter milestones técnicos correspondentes. Nenhuma solução do catálogo pode ficar sem milestone de entrega.

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
Ler `PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md`, Project Charter, Catálogo de Soluções, PRD Definition, Specs Definition.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Alinhamento com Negócio
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Marcos do Charter cobertos | Cada M1-M7 tem milestone técnico correspondente |
| 1.2 | Features cobertas | Features do projeto têm milestone de entrega |
| 1.3 | User Stories referenciadas | US críticas aparecem nos critérios de aceitação |

#### Dimensão 2: Cobertura de Soluções
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Soluções cobertas | Toda solução do catálogo aparece em ≥1 milestone |
| 2.2 | Dependências documentadas | Ordem de construção reflete dependências reais |

#### Dimensão 3: Viabilidade
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Riscos documentados | Cada milestone tem riscos e mitigações |
| 3.2 | Capacidade factível | Estimativa não excede capacidade do time |
| 3.3 | Critérios de aceitação | Mensuráveis e verificáveis |

### Passo 3 — Calcular Veredito
### Passo 4 — Gerar Relatório de Falha (se REPROVADO)

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `roadmap-planning` | Validar roadmap | Roadmap |
| 2 | `gap-analysis` | Identificar milestones faltantes | Análise |
| 3 | `project-manager` | Validar viabilidade | PM |
| 4 | `risk-manager` | Validar riscos | Risco |
| 5 | `senior-pm` | Supervisão de planejamento | PM |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação dos milestones | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
