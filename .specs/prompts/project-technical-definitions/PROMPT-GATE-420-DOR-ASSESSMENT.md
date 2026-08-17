# PROMPT-GATE-420-DOR-ASSESSMENT

## Contexto

Este prompt implementa o **Gate de Validação da Avaliação DoR** para o artefato `420-DOR-ASSESSMENT.md` (Fase 2 — Bloco 0). O agente validador verifica se todos os itens do INTAKE-LOG foram avaliados, se as rejeições estão justificadas e se os itens aprovados atendem integralmente aos critérios DoR.

**Princípio fundamental:** Nenhum item pode transitar do INTAKE-LOG para o backlog sem passar pelo DoR. Itens devolvidos devem ter justificativa clara e acionável.

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

**Arquivos gerados pelo GENERATE:** `420-DOR-ASSESSMENT.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `420-DOR-ASSESSMENT.md` (artefato a validar), `410-INTAKE-LOG.md` (referência de itens a avaliar) e TODOS os documentos de negócio: índices principais + TODOS os arquivos individuais nas subpastas `epics/*.md`, `features/*.md` e `user-stories/*.md`.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Avaliação
| # | Verificação | Critério |
|---|---|---|
| 1.1 | 100% dos itens avaliados | Cada item do INTAKE-LOG possui registro no DoR Assessment (aprovado ou devolvido) |
| 1.2 | Sem itens órfãos | Nenhum item avaliado que não exista no INTAKE-LOG |

#### Dimensão 2: Critérios DoR
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Clareza | Itens aprovados possuem descrição clara e compreensível |
| 2.2 | Testabilidade | Itens aprovados possuem critérios de aceitação ou forma de testar |
| 2.3 | Priorização | Itens aprovados possuem prioridade definida |
| 2.4 | Dependências | Dependências entre itens estão mapeadas |
| 2.5 | Estimativa | Itens aprovados possuem estimativa inicial |

#### Dimensão 3: Justificativa de Devolução
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Motivo claro | Cada item devolvido especifica qual(is) critério(s) DoR não foi(ram) atendido(s) |
| 3.2 | Pendências acionáveis | Matriz de pendências presente com ação necessária e responsável |

#### Dimensão 4: Assinatura PO/PM
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Sign-off formal | Campo de assinatura PO/PM presente (nome e data) |
| 4.2 | Data de aprovação | Data de aprovação preenchida |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

Seu retorno para o usuário humano deve seguir estritamente uma das duas estruturas condicionais abaixo, dependendo do resultado da sua análise:

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS CONFLITOS OU DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO DOR ASSESSMENT: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-DOR-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o documento base determinava:** [Descrever a referência no INTAKE-LOG]
  - **Impacto:** [O risco de item não refinado ingressar no backlog]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o documento, por favor, responda:
1. Quanto ao **[ID-CONFLITO-DOR-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas e de resposta curta para sanar as dúvidas encontradas]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução interna para o orquestrador: O processo pausa aqui e aguarda as respostas do humano. Assim que o humano responder, todo este relatório + as respostas dele serão injetadas no PROMPT-FIX-420-DOR-ASSESSMENT.md)*

---

### ✅ CENÁRIO B: SE O DOCUMENTO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO DOR ASSESSMENT: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `420-DOR-ASSESSMENT.md` gerado conforme as informações fornecidas.
- **AUDITORIA DA IA:** Documento pré-validado. 100% dos itens avaliados, rejeições justificadas, itens aprovados atendem critérios DoR, assinatura PO/PM presente. Nenhum conflito conceitual encontrado pela IA.
- **DIRETRIZ:** Peço que leia o documento para verificar se o mesmo atende plenamente às necessidades de avaliação de prontidão.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O documento está em compliance e reflete corretamente a avaliação DoR de todos os itens?
2. Deseja enviar mais documentos/arquivos para complementar a avaliação?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se o usuário responder "Sim" para a Pergunta 1 e "Não" para as Perguntas 2 e 3, altere o status para [STATUS: COMPLIANCE] e destrave a Fase 3 (Product Backlog List). Se o usuário fornecer novos documentos ou inputs nas Perguntas 2 ou 3, acione imediatamente o fluxo de re-alimentação voltando ao PROMPT-GENERATE-420-DOR-ASSESSMENT).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `requirements-validation` | Validar aplicação dos critérios DoR | Requisitos |
| 2 | `gap-analysis` | Identificar itens não avaliados ou pendências sem justificativa | Análise |
| 3 | `acceptance-criteria` | Verificar testabilidade dos itens aprovados | Qualidade |
| 4 | `business-analyst` | Validar clareza e completude dos itens | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no relatório.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação da avaliação DoR (F2 — Bloco 0) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
