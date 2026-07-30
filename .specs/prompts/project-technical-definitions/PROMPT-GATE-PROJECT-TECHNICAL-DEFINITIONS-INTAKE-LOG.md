# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG

## Contexto

Este prompt implementa o **Gate de Validação do Registro de Ingestão** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` (Fase 1 — Bloco 0). O agente validador verifica se todos os lotes de requisitos estão versionados, rastreáveis, com cobertura completa e metadados preenchidos.

**Princípio fundamental:** Todo documento de negócio (Charter, BRD, Epics, Features, User Stories) deve estar coberto por pelo menos um lote de ingestão. Nenhum requisito pode ingressar no pipeline técnico sem registro formal.

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

**Arquivos gerados pelo GENERATE:** `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` (artefato a validar) e TODOS os documentos de negócio como referência: índices principais + TODOS os arquivos individuais nas subpastas `epics/*.md` (4 arquivos), `features/*.md` (18 arquivos) e `user-stories/*.md` (62 arquivos).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Versionamento
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Números de versão sequenciais | Cada lote possui `v1`, `v2`, ... sem saltos ou duplicatas |
| 1.2 | Data de ingestão | Cada lote possui data de recebimento preenchida |

#### Dimensão 2: Rastreabilidade
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Origem documentada | Cada lote referencia documentos de negócio específicos (Charter, BRD, Epics, Features, US) |
| 2.2 | Tipo de ciclo | Cada lote identifica se é Waterfall ou Ágil (Scrum/Kanban/OKR) |

#### Dimensão 3: Cobertura
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Cobertura total | 100% dos documentos de negócio (Charter, BRD, Epics, Features, US) estão cobertos por pelo menos um lote |
| 3.2 | Matriz de cobertura | Matriz/gráfico de cobertura presente e preenchido |

#### Dimensão 4: Metadados
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Origem/Responsável | Cada lote possui origem (PO, PM, Stakeholder) e responsável |
| 4.2 | Status | Cada lote possui status claro: Recebido, Em Refinamento, Pronto para TI, etc. |
| 4.3 | Escopo | Descrição de escopo presente para cada lote |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

Seu retorno para o usuário humano deve seguir estritamente uma das duas estruturas condicionais abaixo, dependendo do resultado da sua análise:

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS CONFLITOS OU DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO INTAKE-LOG: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-IL-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o documento base determinava:** [Descrever a referência nos documentos de negócio]
  - **Impacto:** [O risco de perda de rastreabilidade ou cobertura]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o documento, por favor, responda:
1. Quanto ao **[ID-CONFLITO-IL-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas e de resposta curta para sanar as dúvidas encontradas]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução interna para o orquestrador: O processo pausa aqui e aguarda as respostas do humano. Assim que o humano responder, todo este relatório + as respostas dele serão injetadas no PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md)*

---

### ✅ CENÁRIO B: SE O DOCUMENTO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO INTAKE-LOG: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` gerado conforme as informações fornecidas.
- **AUDITORIA DA IA:** Documento pré-validado. Lotes versionados, rastreabilidade com documentos de negócio estabelecida, cobertura total, metadados preenchidos. Nenhum conflito conceitual encontrado pela IA.
- **DIRETRIZ:** Peço que leia o documento para verificar se o mesmo atende plenamente às necessidades de registro de ingestão.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O documento está em compliance e reflete corretamente todos os lotes de requisitos recebidos?
2. Deseja enviar mais documentos/arquivos para incluir novos lotes ou complementar os existentes?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se o usuário responder "Sim" para a Pergunta 1 e "Não" para as Perguntas 2 e 3, altere o status para [STATUS: COMPLIANCE] e destrave a Fase 2 (DoR Assessment). Se o usuário fornecer novos documentos ou inputs nas Perguntas 2 ou 3, acione imediatamente o fluxo de re-alimentação voltando ao PROMPT-GENERATE-INTAKE-LOG).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Análise de gaps de cobertura entre lotes e documentos | Análise |
| 2 | `requirements-validation` | Validação de rastreabilidade dos lotes | Requisitos |
| 3 | `business-analyst` | Verificação de escopo e completude dos lotes | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no relatório.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação do registro de ingestão (F1 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
