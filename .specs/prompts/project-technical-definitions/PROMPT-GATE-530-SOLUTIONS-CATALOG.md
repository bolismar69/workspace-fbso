# PROMPT-GATE-530-SOLUTIONS-CATALOG

## Contexto

Este prompt implementa o **Gate de Validação do Catálogo de Soluções (F13)** para o artefato `530-SOLUTIONS-CATALOG.md`. O agente validador verifica se o catálogo está completo, consistente e cobre todos os épicos/features do projeto de negócio, com base nos insumos das 6 disciplinas do Bloco B.

**Princípio fundamental:** Nenhum épico ou feature do projeto de negócio pode ficar sem solução técnica designada. Cobertura 100% é obrigatória.

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

**Arquivos gerados pelo GENERATE:** `530-SOLUTIONS-CATALOG.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `530-SOLUTIONS-CATALOG.md`, documentos de negócio (Charter, BRD, Epics, Features, RTM).

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
| 3.2 | Responsáveis coerentes | Responsáveis existem no TEAM-SKILLS-MAP |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS OU QUEBRAS DE COBERTURA (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE CATÁLOGO: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-SC-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que os documentos de negócio determinavam:** [Descrever a referência do BRD/Epics/Features]
  - **Impacto:** [O risco de épico/feature sem solução técnica]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o catálogo, por favor, responda:
1. Quanto ao **[ID-CONFLITO-SC-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-530-SOLUTIONS-CATALOG.md)*

---

### ✅ CENÁRIO B: SE O CATÁLOGO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE CATÁLOGO: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `530-SOLUTIONS-CATALOG.md` gerado e estruturado conforme os documentos de negócio.
- **AUDITORIA DA IA:** Cobertura verificada. 100% dos épicos/features mapeados para soluções. Classificação, estado e prioridade documentados. Nenhuma solução órfã detectada.
- **DIRETRIZ:** Peço que leia o catálogo para verificar se as soluções listadas atendem plenamente ao escopo técnico esperado.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O catálogo de soluções está em compliance e cobre adequadamente todos os épicos e features do projeto?
2. Deseja enviar mais documentos/arquivos para enriquecer o catálogo?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e F14 (SOLUTIONS-MATRIX). Se novos inputs → retrocede ao PROMPT-GENERATE).*

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
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
