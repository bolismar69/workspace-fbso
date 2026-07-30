# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST

## Contexto

Este prompt implementa o **Gate de Validação do Backlog de Produto** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md` (Fase 3 — Bloco 0). O agente validador verifica se todos os itens aprovados no DoR estão presentes, com rastreabilidade completa, priorização aplicada e links markdown válidos.

**Princípio fundamental:** Todo item aprovado no DoR Assessment deve constar no backlog. Todo item no backlog deve ter rastreabilidade inequívoca até sua origem nos documentos de negócio.

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

**Arquivos gerados pelo GENERATE:** `PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md` (artefato a validar), `PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md` (referência de itens aprovados) e `PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md` (referência de lotes).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Completude do Backlog
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Itens aprovados presentes | 100% dos itens com DoR ✅ no DOR-ASSESSMENT estão listados no backlog |
| 1.2 | Itens devolvidos ausentes | Nenhum item com DoR 🔄 (devolvido) consta no backlog |
| 1.3 | Sem itens órfãos | Nenhum item no backlog sem correspondência no DOR-ASSESSMENT |

#### Dimensão 2: Rastreabilidade
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Origem documentada | Cada item possui campo "Origem" preenchido (Charter/BRD/Epic/Feature/US) |
| 2.2 | Links markdown válidos | Cada link de rastreabilidade segue sintaxe markdown `[texto](caminho)` e aponta para arquivo existente |
| 2.3 | Lote referenciado | Cada item possui referência ao lote de ingestão do INTAKE-LOG |

#### Dimensão 3: Priorização
| # | Verificação | Critério |
|---|---|---|
| 3.1 | MoSCoW aplicado | Todos os itens possuem prioridade MoSCoW (Must/Should/Could/Won't) |
| 3.2 | Resumo por prioridade | Tabela de contagem e percentual por categoria MoSCoW presente |

#### Dimensão 4: Agrupamento
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Resumo por lote | Itens agrupados por lote de ingestão com contagens |
| 4.2 | Consistência com INTAKE-LOG | Lotes referenciados existem no INTAKE-LOG |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

Seu retorno para o usuário humano deve seguir estritamente uma das duas estruturas condicionais abaixo, dependendo do resultado da sua análise:

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS CONFLITOS OU DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO PRODUCT BACKLOG LIST: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-PB-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o documento base determinava:** [Descrever a referência no DOR-ASSESSMENT ou INTAKE-LOG]
  - **Impacto:** [O risco de backlog incompleto ou mal priorizado]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o documento, por favor, responda:
1. Quanto ao **[ID-CONFLITO-PB-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas e de resposta curta para sanar as dúvidas encontradas]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução interna para o orquestrador: O processo pausa aqui e aguarda as respostas do humano. Assim que o humano responder, todo este relatório + as respostas dele serão injetadas no PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md)*

---

### ✅ CENÁRIO B: SE O DOCUMENTO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO PRODUCT BACKLOG LIST: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md` gerado conforme as informações fornecidas.
- **AUDITORIA DA IA:** Documento pré-validado. Todos os itens aprovados no DoR presentes, rastreabilidade completa com links markdown válidos, priorização MoSCoW aplicada, resumos por prioridade e por lote. Nenhum conflito conceitual encontrado pela IA.
- **DIRETRIZ:** Peço que leia o documento para verificar se o mesmo atende plenamente às necessidades de planejamento do backlog técnico.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O documento está em compliance e reflete corretamente o backlog de produto "Pronto para TI"?
2. Deseja enviar mais documentos/arquivos para complementar ou reordenar o backlog?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se o usuário responder "Sim" para a Pergunta 1 e "Não" para as Perguntas 2 e 3, altere o status para [STATUS: COMPLIANCE] e destrave a Fase 4 (PRD Definition). Se o usuário fornecer novos documentos ou inputs nas Perguntas 2 ou 3, acione imediatamente o fluxo de re-alimentação voltando ao PROMPT-GENERATE-PRODUCT-BACKLOG-LIST).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `requirements-validation` | Validar presença e rastreabilidade dos itens | Requisitos |
| 2 | `gap-analysis` | Identificar itens ausentes ou links quebrados | Análise |
| 3 | `product-manager` | Validar priorização MoSCoW | Produto |
| 4 | `backlog-management` | Verificar estrutura e agrupamento do backlog | Gestão |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no relatório.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação do backlog de produto (F3 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
