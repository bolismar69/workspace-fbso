# PROMPT-GATE-595-TECHLEAD-RETURN-PACKAGE

## Contexto

Este prompt implementa o **Gate de Validação do Pacote de Retorno do TECHLEAD** para o artefato `595-RETURN-PACKAGE-{CICLO-NN}.md` (Bloco E — modo waterfall-discovery). O agente validador verifica se o pacote está completo, rastreável e respeita o contrato PM/PO ↔ TECHLEAD.

**Princípio fundamental:** o pacote PROPOE, nunca APLICA. Nenhum arquivo do PM/PO (092/093/095/085/088) pode ter sido alterado pelo TECHLEAD, e todo item do snapshot do 092 deve ter status proposto.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{CICLO_NN}` | Identificador do Ciclo de Entrega (`CICLO-NN`) do 092 |
| `{SNAPSHOT_092}` | Caminho do snapshot do 092-BACKLOG-KANBAN recebido no pacote de demanda |

**Arquivo gerado pelo GENERATE:** `595-RETURN-PACKAGE-{CICLO_NN}.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `595-RETURN-PACKAGE-{CICLO_NN}.md` (artefato a validar) e `{SNAPSHOT_092}` (referência de itens do ciclo).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura do Snapshot
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Status por item | 100% dos `BL-NN` do snapshot 092 possuem status proposto e % concluído — **[595-01]** |
| 1.2 | Itens extras | Nenhum `BL-NN` do pacote sem correspondência no snapshot 092 — **[595-01]** |

#### Dimensão 2: Débito Técnico
| # | Verificação | Critério |
|---|---|---|
| 2.1 | CR vinculada | Todo `DT-XXX` possui proposta de CR Técnico (`CR-NN`) com justificativa para o 085 — **[595-02]** |
| 2.2 | Origem documentada | Todo `DT-XXX` referencia sua fonte (`IDENTIFIED-TECHNICAL-DEBT.md`/`SPRINT-REVIEW.md`) — **[595-02]** |

#### Dimensão 3: Evidências
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Caminhos absolutos | Toda evidência da seção 4 aponta para caminho absoluto — **[595-03]** |
| 3.2 | Evidências existem | Toda evidência aponta para arquivo/PR existente — **[595-03]** |

#### Dimensão 4: Impedimentos e Mudanças
| # | Verificação | Critério |
|---|---|---|
| 4.1 | IMP completo | Todo `IMP-NN` tem nome, descrição, impacto no ciclo e solução sugerida — **[595-04]** |
| 4.2 | Mudança formal | Todo pedido de mudança da seção 6 referencia o processo do 085 (e PERT/065/070 quando aplicável) — **[595-05]** |

#### Dimensão 5: Contrato e Vocabulário
| # | Verificação | Critério |
|---|---|---|
| 5.1 | Veto a termos ágeis | Nenhum termo ágil (Sprint, User Story, DoR, Epic) no pacote — **[595-06]** |
| 5.2 | Ownership preservado | Nenhum arquivo do PM/PO (092/093/095/085/088) foi alterado pelo TECHLEAD — **[595-07]** |

#### Dimensão 6: Janelas de Entrega e Aceite (096)
| # | Verificação | Critério |
|---|---|---|
| 6.1 | Coluna Janela | Todo `BL-NN` tem janela `JAN-*-NN` preenchida, coerente com o 600-EXECUTION-HISTORY — **[595-08]** |
| 6.2 | Evidências por janela | Seção 4 agrupa evidências por janela (DEV/QA/UAT/DEPLOY) e todas apontam para arquivos existentes — **[595-09]** |
| 6.3 | Aceite UAT por entrega | Todo `BL-NN` que saiu da UAT tem registro de DE-ACORDO/APROVAÇÃO por entrega (Key Users + PM/PO) — **[595-10]** |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS CONFLITOS OU DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO RETURN PACKAGE: [CICLO-NN — Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[595-NN] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o contrato determinava:** [Referência ao snapshot 092 ou à regra violada]
  - **Impacto:** [Risco para a aplicação do pacote pelo PM/PO]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
1. Quanto ao **[595-NN]**, qual é a definição correta a ser aplicada?

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução interna para o orquestrador: o processo pausa e aguarda as respostas do humano; em seguida, relatório + respostas são injetadas no PROMPT-FIX-595-TECHLEAD-RETURN-PACKAGE.md)*

---

### ✅ CENÁRIO B: SE O DOCUMENTO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DO RETURN PACKAGE: [CICLO-NN — Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `595-RETURN-PACKAGE-{CICLO-NN}.md` gerado conforme as informações fornecidas.
- **AUDITORIA DA IA:** Documento pré-validado. Cobertura total do snapshot 092, débito técnico com CR vinculada, evidências com caminhos válidos agrupadas por janela, coluna "Janela" coerente com o 600, aceite UAT (DE-ACORDO) presente por entrega, impedimentos completos, pedidos de mudança formais, vocabulário WATERFALL e ownership preservado. Nenhum conflito conceitual encontrado pela IA.
- **DIRETRIZ:** Peço que leia o pacote para verificar se o mesmo atende plenamente ao contrato PM/PO ↔ TECHLEAD antes de enviá-lo ao PM/PO.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O pacote está em compliance e pode ser enviado ao PM/PO?
2. Deseja enviar mais documentos/arquivos para complementar o pacote?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se o usuário responder "Sim" para a Pergunta 1 e "Não" para as Perguntas 2 e 3, altere o status para [STATUS: COMPLIANCE] e entregue o pacote ao PM/PO — a recepção segue o item 3.3 da WATERFALL-EXECUTION v2.0. Se o usuário fornecer novos documentos ou inputs nas Perguntas 2 ou 3, acione o fluxo de re-alimentação voltando ao PROMPT-GENERATE-595-TECHLEAD-RETURN-PACKAGE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `requirements-validation` | Validar cobertura e rastreabilidade do pacote | Requisitos |
| 2 | `gap-analysis` | Identificar itens ausentes ou vínculos quebrados | Análise |
| 3 | `backlog-management` | Verificar status e máquina de estados do 092 | Gestão |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no relatório.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 16/08/2026 | Criação inicial: gate de validação do pacote de retorno (Bloco E — modo waterfall-discovery) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
