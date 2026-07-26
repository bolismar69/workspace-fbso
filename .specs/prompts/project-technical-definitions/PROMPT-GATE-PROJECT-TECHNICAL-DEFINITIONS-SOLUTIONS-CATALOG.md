# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG

## Contexto

Este prompt implementa o **Gate de Validação do Catálogo de Soluções** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md`. O agente validador verifica se o catálogo está completo, consistente e cobre todos os épicos/features do projeto de negócio.

**Princípio fundamental:** Nenhum épico ou feature do projeto de negócio pode ficar sem solução técnica designada. Cobertura 100% é obrigatória.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md`, documentos de negócio (Charter, BRD, Epics, Features, RTM).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Negócio
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Cobertura de épicos | Cada épico tem pelo menos 1 solução designada |
| 1.2 | Cobertura de features | Cada feature tem solução(ões) mapeada(s) |
| 1.3 | Sem soluções órfãs | Nenhuma solução sem vínculo com épico/feature |

#### Dimensão 2: Completude do Catálogo
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Classificação por tipo | Toda solução tem tipo definido |
| 2.2 | Estado documentado | Toda solução tem estado (existente/criar/planejado) |
| 2.3 | Prioridade definida | MoSCoW atribuído para cada solução |

#### Dimensão 3: Consistência Cross-Documento
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com TECHNICAL-PLAN | Soluções do TECHNICAL-PLAN constam no catálogo |
| 3.2 | Responsáveis coerentes | Responsáveis existem no TEAM-MAP |

### Passo 3 — Calcular Veredito
100% OK → APROVADO | ≥ 75% → RESSALVA | < 75% → REPROVADO

### Passo 4 — Gerar Relatório de Falha (se REPROVADO)
Gerar `SOLUTIONS_CATALOG_SCOPE_FAIL_REPORT.md`.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Identificar lacunas de cobertura | Análise |
| 2 | `architecture-patterns` | Validar padrões arquiteturais | Arquitetura |
| 3 | `system-design` | Validar design do sistema | Arquitetura |
| 4 | `domain-driven-design` | Validar bounded contexts | DDD |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação do catálogo de soluções | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
