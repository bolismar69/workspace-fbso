# PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD

## Contexto

Este prompt implementa o **GATE do PRD Discovery-Level** para o artefato `DISCOVERY-LEVEL-PRD.md` (Fase 1 — Bloco 0).

**Princípio fundamental:** O PRD Discovery-Level deve cobrir todos os Épicos definidos pelo Negócio com uma visão macro suficiente para embasar a análise de viabilidade técnica.

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{UPSTREAM_DISCOVERY_PATH}` | Caminho upstream-architecture-discovery |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de soluções técnicas |
| `{ARCHITECTURE_GLOBAL}` | Caminho da arquitetura global |
| `{SECURITY_GLOBAL}` | Caminho do GLOBAL-SECURITY.md |

**Arquivos gerados pelo GENERATE:** `DISCOVERY-LEVEL-PRD.md`

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `DISCOVERY-LEVEL-PRD.md` e documentos de negócio (Charter, BRD, Épicos incluindo `epics/*.md`).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Épicos
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Cobertura total | 100% dos Épicos referenciados no PRD |
| 1.2 | Soluções mapeadas | Cada Épico vinculado a ≥1 solução |

#### Dimensão 2: Completude
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Seções presentes | Visão, Matriz, MVP Macro, Restrições, Glossário |
| 2.2 | MVP definido | Escopo do MVP Macro claro |

### Passo 3 — Emitir Veredito

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE — lista conflitos com ID, descrição, impacto, sugestão
### ✅ CENÁRIO B: PRÉ-COMPLIANCE — 3 perguntas obrigatórias de validação humana

*(Instrução de Orquestração: Se "Sim, Não, Não" → COMPLIANCE e Bloco B (F2). Se novos inputs → retrocede ao GENERATE).*

## Skills Utilizados

| Ordem | Skill | Propósito |
|---|---|---|
| 1 | `requirements-validation` | Validar cobertura de épicos |
| 2 | `gap-analysis` | Identificar épicos órfãos |
| 3 | `product-manager` | Validar MVP Macro |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial | Time de Arquitetura |

🤖 *Upstream Architecture Discovery — Fase 1 GATE*
