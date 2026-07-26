# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação do PRD Definition** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md`. Verifica se a baseline de produto cobre todos os requisitos de negócio, mapeia corretamente para soluções e define MVP global consistente.

**Princípio fundamental:** 100% dos requisitos de negócio (BRD, Features, User Stories) devem estar mapeados para pelo menos uma solução. Nenhum requisito órfão é permitido.

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
Ler `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md`, documentos de negócio (Charter, BRD, Epics, Features, User Stories, RTM), `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md`.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Requisitos
| # | Verificação | Critério |
|---|---|---|
| 1.1 | BRs cobertos | Cada BR do BRD mapeado para ≥1 solução |
| 1.2 | Features cobertas | Cada feature mapeada para solução(ões) |
| 1.3 | User Stories cobertas | Cada US referenciada na matriz |
| 1.4 | Sem órfãos | Nenhum requisito sem solução designada |

#### Dimensão 2: Completude do PRD
| # | Verificação | Critério |
|---|---|---|
| 2.1 | 8 seções presentes | Todas as seções obrigatórias preenchidas |
| 2.2 | MVP Global definido | Escopo do MVP claro e mensurável |
| 2.3 | Glossário unificado | Termos canônicos definidos |

#### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com Charter | PRD não contradiz premissas do Charter |
| 3.2 | Soluções referenciadas existem | Toda solução na matriz existe no Catálogo |
| 3.3 | Termos consistentes | Glossário usa terminologia do BRD |

### Passo 3 — Calcular Veredito
### Passo 4 — Gerar Relatório de Falha (se REPROVADO)

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `requirements-validation` | Validar cobertura de requisitos | Requirements |
| 2 | `gap-analysis` | Identificar requisitos órfãos | Análise |
| 3 | `prd` | Validar estrutura do PRD | Product |
| 4 | `product-manager` | Validar MVP global | Product |
| 5 | `business-analyst` | Validar alinhamento com negócio | Business |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação do PRD Definition | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
