# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX

## Contexto

Este prompt implementa o **Gate de Validação da Matriz de Stacks (F15)** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md`. Verifica se todas as soluções têm stack definida, versões especificadas e compatibilidade cross-solution garantida, com base nas 6 disciplinas do Bloco B.

**Princípio fundamental:** Nenhuma solução sem stack definida. Compatibilidade de versões entre soluções que se integram é obrigatória.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

**Arquivos gerados pelo GENERATE:** `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md`

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

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE STACKS: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-SM-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o catálogo/ADR determinava:** [Descrever a referência]
  - **Impacto:** [O risco de incompatibilidade ou stack incorreta]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a matriz de stacks, por favor, responda:
1. Quanto ao **[ID-CONFLITO-SM-01]**, qual é a stack correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md)*

---

### ✅ CENÁRIO B: SE A MATRIZ ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE STACKS: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md` gerado conforme o catálogo e ADRs.
- **AUDITORIA DA IA:** Cobertura completa verificada. Stacks definidas para todas as soluções com versões específicas. Compatibilidade cross-solution validada. Nenhuma solução sem stack detectada.
- **DIRETRIZ:** Peço que leia a matriz para verificar se as stacks definidas atendem às suas expectativas técnicas.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A matriz de stacks está em compliance e as tecnologias/versões escolhidas estão corretas?
2. Deseja enviar mais documentos/arquivos para enriquecer a definição de stacks?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e F16 (SPECS-DEFINITION). Se novos inputs → retrocede ao PROMPT-GENERATE).*

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
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
