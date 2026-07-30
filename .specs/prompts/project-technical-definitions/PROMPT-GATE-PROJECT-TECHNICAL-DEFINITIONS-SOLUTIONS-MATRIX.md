# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX

## Contexto

Este prompt implementa o **Gate de Validação da Matriz de Soluções (F14)** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md`. Verifica se a matriz-mestra está completa, consistente e todas as referências cruzadas com as 6 disciplinas do Bloco B e o Catálogo (F13) são válidas.

**Princípio fundamental:** Toda solução do catálogo deve ter: responsável, repositório, stack, time e status definidos na matriz. Nenhum campo pode ficar vazio.

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

**Arquivos gerados pelo GENERATE:** `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md`

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

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE MATRIZ: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-MAT-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o Catálogo/Stack Matrix/TEAM-CAPACITY determinava:** [Descrever a referência]
  - **Impacto:** [O risco de alocação incorreta ou referência quebrada]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a matriz de soluções, por favor, responda:
1. Quanto ao **[ID-CONFLITO-MAT-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md)*

---

### ✅ CENÁRIO B: SE A MATRIZ ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE MATRIZ: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md` gerado consolidando todas as fases anteriores.
- **AUDITORIA DA IA:** Completude verificada. Todas as soluções com responsável, repositório, stack, time e status. Consistência cruzada com Catálogo, Stack Matrix, TEAM-CAPACITY e Milestones validada. Indicadores de cobertura e risco calculados. RACI preenchido.
- **DIRETRIZ:** Peço que leia a matriz-mestra para verificar se a alocação de recursos e responsabilidades está correta.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A matriz de soluções está em compliance e a alocação de recursos está correta?
2. Deseja enviar mais documentos/arquivos para enriquecer a matriz?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e F15 (SOLUTIONS-STACK-MATRIX). Se novos inputs → retrocede ao PROMPT-GENERATE).*

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
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
