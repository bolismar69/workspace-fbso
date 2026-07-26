# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX

## Contexto

Este prompt implementa o **Gate de Validação da Matriz de Stacks** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md`. Verifica se todas as soluções têm stack definida, versões especificadas e compatibilidade cross-solution garantida.

**Princípio fundamental:** Nenhuma solução sem stack definida. Compatibilidade de versões entre soluções que se integram é obrigatória.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{ARCHITECTURE_GLOBAL}` | Caminho da pasta de arquitetura global |

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md`, `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md`, ADRs e blueprints globais.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Todas as soluções cobertas | Cada solução do catálogo tem stack definida |
| 1.2 | Dimensões obrigatórias | Linguagem, Framework, Banco preenchidos para cada solução |

#### Dimensão 2: Precisão Técnica
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Versões especificadas | Tecnologias têm versão (não "latest") |
| 2.2 | Justificativas | Escolhas não-óbvias têm justificativa |
| 2.3 | Referências a ADRs | Stacks referenciam ADRs relevantes |

#### Dimensão 3: Compatibilidade Cross-Solution
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Consistência de versões | Serviços integrados usam versões compatíveis |
| 3.2 | Alinhamento com blueprints | Stacks seguem blueprints da pasta architecture/ |

### Passo 3 — Calcular Veredito
100% OK → APROVADO | ≥ 75% → RESSALVA | < 75% → REPROVADO

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `tech-stack-evaluator` | Validar escolhas de stack | Avaliação |
| 2 | `gap-analysis` | Identificar gaps de cobertura | Análise |
| 3 | `senior-architect` | Validar consistência arquitetural | Arquitetura |
| 4 | `java-architect` | Validar stack Java | Java |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da matriz de stacks | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
