# PROMPT-GATE-570-MILESTONES (F17)

## Contexto

Este prompt implementa o **Gate de Validação dos Milestones (F17)** para o artefato `570-MILESTONES.md`. Verifica se o roadmap técnico cobre todas as soluções, está alinhado com o Project Charter, e se a SPECS-DEFINITION (F16) foi devidamente incorporada como insumo primário.

**Princípio fundamental:** Todo milestone do Project Charter deve ter milestones técnicos correspondentes. Nenhuma solução do catálogo pode ficar sem milestone de entrega. A SPECS-DEFINITION deve ser a base técnica para o sequenciamento.

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

**Arquivos gerados pelo GENERATE:** `570-MILESTONES.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `570-MILESTONES.md`, Project Charter, Catálogo de Soluções, PRD Definition, Specs Definition.

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

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE MILESTONES: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-MIL-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o Charter/Catálogo determinava:** [Descrever a referência]
  - **Impacto:** [O risco de atraso ou desalinhamento com negócio]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir os milestones, por favor, responda:
1. Quanto ao **[ID-CONFLITO-MIL-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-570-MILESTONES.md)*

---

### ✅ CENÁRIO B: SE OS MILESTONES ESTIVEREM 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE MILESTONES: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `570-MILESTONES.md` gerado conforme Charter e catálogo.
- **AUDITORIA DA IA:** Alinhamento verificado. Marcos do Charter cobertos com milestones técnicos. Todas as soluções com milestone de entrega. Dependências, riscos e critérios de aceitação documentados. Capacidade validada contra TEAM-CAPACITY.
- **DIRETRIZ:** Peço que leia os milestones para verificar se o roadmap técnico está alinhado com suas expectativas de prazo e sequenciamento.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. Os milestones técnicos estão em compliance e o sequenciamento proposto é viável?
2. Deseja enviar mais documentos/arquivos para enriquecer o planejamento?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e F18 (PACKAGE-BACKLOG). Se novos inputs → retrocede ao PROMPT-GENERATE).*

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
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
