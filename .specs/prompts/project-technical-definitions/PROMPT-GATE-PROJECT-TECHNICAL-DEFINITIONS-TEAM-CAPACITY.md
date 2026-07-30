# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY

## Contexto

Este prompt implementa o **Gate de Validação da Capacidade do Time** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (Bloco A — Fase 6). O agente validador verifica se a tabela de capacidade está consistente com a matriz de skills definida na Fase 5 (TEAM-SKILLS-MAP) e alinhada com o escopo do Bloco 0.

**Princípios fundamentais:**
1. Todo Papel listado no `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md` deve ter uma entrada correspondente no TEAM-CAPACITY. Nenhum papel do Discovery Team pode ficar sem alocação de capacidade.
2. Se existir algum Papel no TEAM-CAPACITY que NÃO esteja no TEAM-SKILLS-MAP, isso deve ser alertado como potencial inconsistência — o humano decide se o papel extra é intencional ou deve ser removido.

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

**Arquivos gerados pelo GENERATE:** `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (artefato a validar), `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md` (referência de papéis do Discovery Team), `PRODUCT-BACKLOG-LIST.md` (F3) e `PRD-DEFINITION.md` (F4 — Bloco 0) para validar alinhamento da capacidade com o escopo do projeto.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Papéis (SKILLS-MAP → CAPACITY)
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Papéis do Discovery Team cobertos | Cada um dos 9 papéis do TEAM-SKILLS-MAP tem linha correspondente no TEAM-CAPACITY |
| 1.2 | Papéis com nome preenchido | Se o TEAM-SKILLS-MAP já tem nomes definidos, verificar se esses nomes aparecem no TEAM-CAPACITY |

#### Dimensão 2: Consistência Reversa (CAPACITY → SKILLS-MAP)
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Sem papéis órfãos no CAPACITY | Se existir um Papel no TEAM-CAPACITY que não corresponde a nenhum papel do TEAM-SKILLS-MAP, alertar o humano |
| 2.2 | Capacidade preenchida | Para cada linha com Nome preenchido, a capacidade semanal está definida |

#### Dimensão 3: Completude
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Colunas obrigatórias | Papel, Nível, Capacidade semana — estrutura presente |
| 3.2 | Legenda e instruções | Seções de legenda e instruções de preenchimento presentes |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE CAPACIDADE: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-TC-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o TEAM-SKILLS-MAP determinava:** [Descrever a referência]
  - **Impacto:** [O risco de papel sem alocação ou inconsistência]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ⚠️ Alertas de Papéis Órfãos (se aplicável):
- **Papel no CAPACITY sem correspondência no SKILLS-MAP:** [Nome do Papel] — este papel existe na tabela de capacidade mas não está listado no Discovery Team do TEAM-SKILLS-MAP. Deseja mantê-lo (papel adicional intencional) ou removê-lo?

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o documento, por favor, responda:
1. Quanto ao **[ID-CONFLITO-TC-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md)*

---

### ✅ CENÁRIO B: SE O DOCUMENTO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE CAPACIDADE: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` gerado conforme o TEAM-SKILLS-MAP.
- **AUDITORIA DA IA:** Cobertura verificada. Todos os papéis do Discovery Team têm entrada no TEAM-CAPACITY. Sem papéis órfãos. Estrutura de colunas completa.
- **DIRETRIZ:** Peço que leia o documento para verificar se a alocação de capacidade está correta e se todos os integrantes necessários estão contemplados.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O documento de capacidade está em compliance e reflete corretamente a alocação do time?
2. Deseja enviar mais documentos/arquivos para enriquecer a tabela de capacidade?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e Bloco B (F7 — ARCHITECTURE-DEFINITION). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `team-composition-analysis` | Validar composição e alocação do time | People |
| 2 | `gap-analysis` | Identificar papéis sem capacidade | Análise |
| 3 | `project-manager` | Validar alocação de recursos | PM |
| 4 | `senior-pm` | Supervisão de capacidade | PM |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 28/07/2026 | Criação inicial: gate de validação da capacidade do time (Fase 2) | Time de Arquitetura |
| 2.0 | 30/07/2026 | Atualização Bloco A (F5-F6): adicionados inputs PRODUCT-BACKLOG-LIST (F3) e PRD-DEFINITION (F4); atualizadas referências de fase (Fase 2→Fase 6, orquestração Fase 3→Bloco B) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
