# PROMPT-GATE-440-PRD-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação do PRD Definition** para o artefato `440-PRD-DEFINITION.md` (PRD de Negócio — Bloco 0, Fase 4). Verifica se a baseline de produto cobre 100% do backlog (F3), mapeia corretamente as entregas de negócio e define MVP global consistente.

**Princípio fundamental:** 100% dos requisitos de negócio (BRD, Features, User Stories) devem estar mapeados para pelo menos uma solução. Nenhum requisito órfão é permitido.

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

**Arquivos gerados pelo GENERATE:** `440-PRD-DEFINITION.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `440-PRD-DEFINITION.md`, TODOS os documentos de negócio: índices principais + TODOS os arquivos individuais nas subpastas `epics/*.md` (4 arquivos), `features/*.md` (18 arquivos) e `user-stories/*.md` (62 arquivos), INTAKE-LOG, DOR-ASSESSMENT, PRODUCT-BACKLOG-LIST (artefatos do Bloco 0), `530-SOLUTIONS-CATALOG.md`.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Requisitos e Backlog
| # | Verificação | Critério |
|---|---|---|
| 1.1 | BRs cobertos | Cada BR do BRD mapeado para ≥1 solução |
| 1.2 | Features cobertas | Cada feature mapeada para solução(ões) |
| 1.3 | User Stories cobertas | Cada US referenciada na matriz |
| 1.4 | Backlog coberto | 100% dos itens do PRODUCT-BACKLOG-LIST (F3) contemplados no PRD |
| 1.5 | Sem órfãos | Nenhum requisito sem solução designada |

#### Dimensão 2: Completude do PRD
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Seções obrigatórias presentes | Visão, Matriz, MVP, Requisitos de Negócio, Restrições e Glossário — todas preenchidas |
| 2.2 | MVP Global definido | Escopo do MVP claro e mensurável |
| 2.3 | Glossário unificado | Termos canônicos definidos |

#### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com Charter | PRD não contradiz premissas do Charter |
| 3.2 | Soluções referenciadas existem | Toda solução na matriz existe no Catálogo |
| 3.3 | Termos consistentes | Glossário usa terminologia do BRD |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE PRD: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-PRD-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que os documentos de negócio determinavam:** [Descrever a referência do BRD/Features/US]
  - **Impacto:** [O risco de requisito sem cobertura ou MVP mal definido]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o PRD Definition, por favor, responda:
1. Quanto ao **[ID-CONFLITO-PRD-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-440-PRD-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE O PRD ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE PRD: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `440-PRD-DEFINITION.md` gerado conforme os documentos de negócio.
- **AUDITORIA DA IA:** Cobertura de requisitos verificada. 100% dos BRs, Features e User Stories mapeados para soluções. MVP Global definido, 6 seções obrigatórias. Nenhum requisito órfão detectado.
- **DIRETRIZ:** Peço que leia o PRD Definition para verificar se a baseline de produto atende às suas expectativas.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O PRD Definition está em compliance e reflete corretamente a baseline de produto do projeto?
2. Deseja enviar mais documentos/arquivos para enriquecer a definição de produto?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e Bloco A (F5-TEAM-SKILLS-MAP). Se novos inputs → retrocede ao PROMPT-GENERATE).*

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
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |
| 3.0 | 30/07/2026 | Migração para Bloco 0 (F4): adicionado critério de cobertura de backlog (F3); atualizada orquestração para Bloco A (F5) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
