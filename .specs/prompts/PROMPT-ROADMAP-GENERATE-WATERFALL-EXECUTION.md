# PROMPT: ROADMAP DE EXECUÇÃO E CONSTRUÇÃO — FASE 5 DO WATERFALL (WATERFALL-EXECUTION)
## Versão: 2.3 — WATERFALL Orchestrator v3.3 (6 Fases, 39 Documentos) — Delegação da Construção ao TECHLEAD (PROJECT-TECHNICAL-DEFINITIONS v7.3, modo waterfall-discovery) + Planejamento do Ciclo de Entrega (Sub-fase 1.5) + Janelas de Entrega (096 + Bloco F) + Tooling de Ambiente (IMPLEMENTATION-TOOLING)

Atue como um Arquiteto de Soluções Organizacionais e Gestor de Execução, especializado em metodologia WATERFALL, engenharia de prompts e orquestração de esteiras de desenvolvimento.

## Propósito e Posição no Fluxo

Este roadmap orquestra a **FASE 5 — EXECUÇÃO E CONSTRUÇÃO** do fluxo WATERFALL no **escopo do PM/PO**: gestão do backlog (092), gestão do time (093), entrega das demandas ao TECHLEAD e recepção da atualização do 092. A **construção é delegada ao TECHLEAD** via `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` (v7.1, modo `waterfall-discovery`) — este roadmap cuida do **contrato PM/PO ↔ TECHLEAD** (handoff e recepção).

```
FASE 0 (Bootstrap) → FASE 1 (Negócio) → FASE 2 (Especificação) → FASE 3 (Engenharia) → FASE 4 (Baseline)
    → 🚩 M4: PROJECT BASELINE LOCKED
    → FASE 5 (EXECUÇÃO E CONSTRUÇÃO — ESTE ROADMAP)
    → 🚩 M5: GO-LIVE & HANDOVER
    → FASE 6 (ENCERRAMENTO E OPERAÇÃO — docs 105/110/115)
```

- **Entrada:** M4 travado (baseline de escopo, cronograma e orçamento selados)
- **Saída:** M5 (GO-LIVE & HANDOVER) — com 095-RELATORIO-QUALIDADE alimentado com evidências
- **Posição:** orquestrador da FASE 5 do roadmap master `PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md`

## Regras de Ouro de Execução (NÃO NEGOCIÁVEIS)

1. **M4 TRAVADO:** execução só inicia com todos os documentos F1–F4 em `[STATUS: COMPLIANCE]`. Mudança de escopo pós-M4 passa OBRIGATORIAMENTE pelo 085-PLANO-GESTAO-MUDANCAS
2. **NÃO ADIVINHAR INPUTS:** nenhum prompt infere, busca ou descobre inputs — todo parâmetro é passado explicitamente
3. **TODO GENERATE TEM GATE+FIX:** vale para os triads 092/093 e para os artefatos de ciclo (GATEs delegados dos prompts reusados)
4. **FIX É CIRÚRGICO:** correção apenas nas seções com violações reportadas
5. **HITL OBRIGATÓRIO POR CICLO:** a revisão humana do loop de execução (code review) é obrigatória — nenhum ciclo avança sem aprovação explícita do humano
6. **VOCABULÁRIO WATERFALL:** vetar termos ágeis (Epic/User Story/DoR/Sprint) nos artefatos WATERFALL — usar a tabela de tradução (abaixo)
7. **BASELINE PRESERVADA:** 092 opera o 088, mas o 088 permanece o registro congelado em M4 — mudanças são registradas com rastreabilidade, nunca sobrescritas silenciosamente

### Contrato PM/PO ↔ TECHLEAD (regras de delegação — NÃO NEGOCIÁVEIS)

1. **TECHLEAD PROPOE, PM/PO APLICA:** o TECHLEAD nunca edita 092/093/095/085/088 — devolve o pacote `595-RETURN-PACKAGE-{CICLO-NN}.md`; o PM/PO aplica via GENERATE-092 (modo atualização) e revalida com GATE-092
2. **PACOTE DE DEMANDA EXPLÍCITO:** o handoff entrega snapshot do 092 (CICLO-NN, BL-NN, CR-NN) + docs F1–F4; o TECHLEAD nunca "puxa" arquivos da pasta WATERFALL (Regra 2 preservada)
3. **CICLO SÓ FECHA COM RECEPÇÃO FORMAL:** nenhum ciclo é considerado encerrado antes da recepção do pacote 595 (GATE aprovado) e da aplicação no 092

## VOCABULÁRIO WATERFALL (tabela de tradução obrigatória)

| Termo ágil (PROIBIDO) | Equivalente WATERFALL (usar) |
|---|---|
| Epic | Pacote de trabalho da EAP (060-EAP-WBS) |
| Feature | Funcionalidade `FEAT-NN` (010-FRD) |
| User Story | Caso de Uso `UC-NN` (010-FRD) |
| Definition of Ready (DoR) | GATE de COMPLIANCE do documento de origem |
| Sprint | Ciclo de Entrega `CICLO-NN` (definido pelo 092) |
| Sprint Planning | Planejamento do Ciclo de Entrega `CICLO-NN` (WATERFALL-EXECUTION, Sub-fase 1.5) |
| Product Backlog | 088-PRODUCT-BACKLOG-LIST (operado pelo 092 na FASE 5) |
| Discovery técnico por ciclo | `technical-discovery/590-ciclo-NNN` (PROJECT-TECHNICAL-DEFINITIONS — espelha a `CICLO-NN`) |
| Bloco de trabalho do TECHLEAD | Ciclo de Entrega `CICLO-NN` (092) |

---

## Inputs da Fase 5 — "Time recebe tudo pronto" (UPSTREAM F1–F4)

O time de desenvolvimento recebe, no dia 1 da FASE 5:

| Categoria | Documentos |
|---|---|
| Negócio (F1–F2) | 001-Charter, 002-Stakeholder-Map, 003-Personas-Jornadas, 004-AS-IS/TO-BE, 005-BRD, 010-FRD, 015-RTM-FASE-1 |
| Funcional (F1–F2) | 010-FRD, 016-Protótipos UX/UI, 020-SRS, 025-RTM-FASE-2 |
| Design Técnico (F3) | 030-SAD, 035-HLD, 040-LLD, 041-DEVOPS-SETUP, 042-DATA-SETUP, 043-SEC-SETUP, 044-INFRA-SETUP, 045-TEST-PLAN, 050-TEST-CASES, 060-EAP-WBS |
| Planejamento e Regras (F4) | 062-STAFFING-PLAN, 065-Cronograma, 070-Orçamento, 075-Comunicação, 080-Riscos, 085-Mudanças, 086-Padrões/DoD, 087-CI-CD, 088-Backlog, 090-STRATEGIC-IMPLEMENTATION-AND-DEPLOYMENT-PLAN |

**Regra:** o orquestrador só inicia a FASE 5 com M4 travado e 095-RELATORIO-QUALIDADE com a estrutura de métricas criada na FASE 3 (as evidências são alimentadas ao longo da FASE 5).

---

## Sub-fase 1 — Gestão Diária & Operacional

Triads WATERFALL próprios (em `project-documents-waterfall/`):

| Doc | Prompt | Responsabilidades |
|---|---|---|
| 092-BACKLOG-KANBAN | `PROMPT-GENERATE/GATE/FIX-092-BACKLOG-KANBAN.md` | Revisa/expande o 088 via **Change-Request de Negócio** e **Change-Request Técnico** (formalizadas pelo 085); atualiza **status** (A Fazer → Em Execução → Em Revisão → Concluído/Impedido); define **Ciclos de Entrega** (`CICLO-NN`) de implementação |
| 093-GESTAO-TIMES | `PROMPT-GENERATE/GATE/FIX-093-GESTAO-TIMES.md` | Capacidade, impedimentos (`IMP-NN`) e alocação contra o 062-STAFFING-PLAN |

O par 092 ↔ 093 roda em loop contínuo durante toda a FASE 5 (como no flowchart: `092 → 093 → 092`).

**Modo atualização do 092 (consumo do pacote 595):** na recepção TECHLEAD→PM/PO (sub-fase 3.3), o GENERATE-092 roda em modo atualização com `EXTRA_INPUTS` = `595-RETURN-PACKAGE-{CICLO-NN}.md` — aplica status propostos e CRs Técnicas, mantendo `[STATUS: Em revisão]` até o GATE-092 e a validação humana.

## Sub-fase 1.5 — Planejamento do Ciclo de Entrega (`CICLO-NN`)

Rito formal por ciclo — o equivalente WATERFALL da reunião de Sprint Planning. Roda **uma vez por ciclo**, antes do handoff (3.1): a saída desta sub-fase É o pacote de demanda do ciclo.

### Participantes

| Papel | Responsabilidade |
|---|---|
| **PM/PO** (convoca e preside) | Dono do 092/088 — seleciona e **aplica** as decisões |
| **TECHLEAD** (participa) | **Propõe**: viabilidade técnica, dependências, ordem, skills, expectativa de janelas |
| Especialistas disciplinares (opcional) | Arquitetura/Segurança/Dados, para itens complexos |
| Key Users | ❌ Não participam — entram apenas no UAT |

### Entradas

| Entrada | Uso no rito |
|---|---|
| `088-PRODUCT-BACKLOG-LIST` (congelado em M4) | Candidatos `BL-NN` |
| `092-BACKLOG-KANBAN` | Estado atual dos ciclos |
| `595-RETURN-PACKAGE-{CICLO-NN}` do ciclo **anterior** (recebido na 3.3) | Retroalimentação: débitos `DT-XXX`, pendências, oportunidades |
| `093-GESTAO-TIMES` + `460-TEAM-CAPACITY` | Capacidade disponível |
| `080-PLANO-RISCOS` | Riscos abertos relevantes ao ciclo |
| `085-PLANO-GESTAO-MUDANCAS` | CRs aprovadas a incorporar |
| `096-DEFINICAO-JANELAS-ENTREGA` | Janelas do ciclo |
| `060-EAP-WBS` + `065-CRONOGRAMA-GANTT` | Estimativas e marco do ciclo |

### Pauta (9 pontos)

1. **Contexto** — marco do ciclo (065), CRs pendentes (085), riscos abertos (080)
2. **Retroalimentação do ciclo anterior** — ler o `595-RETURN-PACKAGE-{CICLO-NN}` anterior: débitos técnicos `DT-XXX` reportados, `BL-NN` não concluídos (retorno ao ciclo com justificativa), oportunidades/melhorias propostas pelo TECHLEAD, resultados das janelas (Aceite UAT, evidências, 095 GO/NO-GO). O PM/PO decide o tratamento de cada item ANTES da seleção
3. **Seleção de candidatos** — `BL-NN` do 088 elegíveis (dependências concluídas, prioridade de negócio)
4. **Capacidade × demanda** — cruzar estimativas (060) com capacidade (093/460); o TECHLEAD propõe cortes/ajustes
5. **Escopo do ciclo** — `BL-NN` que entram na `CICLO-NN` + **Meta do Ciclo** (equivalente ao Sprint Goal)
6. **Dependências e ordem** — sequenciamento preliminar (o detalhamento fino em `T-NNN` fica no 580)
7. **Janelas** — expectativa de passagem DEV→QA→UAT→DEPLOY (096)
8. **Riscos do ciclo** — itens do 080 que afetam o ciclo + mitigação
9. **Decisões** — o PM/PO aplica no 092 via `GENERATE-092` (modo planejamento) e valida com `GATE-092` (HITL, 3 perguntas)

### Saída

- **092 atualizado**: `CICLO-NN` criada/atualizada com `BL-NN`, **Meta do Ciclo**, janela alvo e vínculo aos riscos do 080
- **Pacote de demanda do ciclo** (insumo direto do handoff 3.1): snapshot do 092 + docs F1–F4 + cruzamento `CICLO-NN → FEAT-NN/UC-NN (010) → pacote EAP (060) → T-NNN` (a preencher no 580)
- **Registro**: aprovação no GATE-092 + linha no `600-EXECUTION-HISTORY` (TECHLEAD) — **sem documento novo obrigatório** (ata opcional como `EXTRA_INPUTS`)

### Regras (preservam o contrato PM/PO ↔ TECHLEAD)

1. **TECHLEAD propõe, PM/PO aplica:** o TECHLEAD nunca edita o 092 — propõe no rito; o PM/PO aplica via GENERATE-092 e valida com GATE-092
2. **Baseline preservada:** escopo descoberto no planejamento (item que não cabe, dependência órfã) entra via 085 — nunca por edição silenciosa do 088
3. **O rito não decide sozinho:** a reunião alimenta o GATE-092 + aprovação humana, que são quem decidem
4. **Ciclo ≠ janela:** a 1.5 define o ciclo (`CICLO-NN`); a passagem por janelas continua sendo do Bloco F (TECHLEAD)
5. **Loop fechado:** o 595 do ciclo anterior é entrada obrigatória do planejamento — `1.5 → 3.1 → Bloco E/F → 3.3 → 1.5`

## Sub-fase 2 — Janelas de Entrega (DEV → QA → UAT → DEPLOY)

As janelas são **definidas no `096-DEFINICAO-JANELAS-ENTREGA`** (documento WATERFALL, trio em `project-documents-waterfall/`) e **orquestradas pelo TECHLEAD no Bloco F do PROJECT-TECHNICAL-DEFINITIONS v7.1** — o PM/PO não executa janelas.

Loop por ciclo `CICLO-NN`: `DEV (Bloco E) → gate → QA (050 + QA-REVISOR + 095 GO/NO-GO) → UAT (DE-ACORDO/APROVAÇÃO por entrega — Key Users + PM/PO) → DEPLOY (090 + 087, GMUD) → pacote 595 → PM/PO aplica no 092 → próximo `CICLO-NN`.

**Quem propõe/aceita por janela:**
- DEV: TECHLEAD executa e propõe a conclusão; aceite = GATE do ciclo + revisor humano (086)
- QA: TECHLEAD executa e propõe; PM/PO aceita via 095 (GO/NO-GO)
- UAT: PM/PO + Key Users executam; aceite = registro de DE-ACORDO/APROVAÇÃO por entrega (alimenta o 595 e pode ser refletido no 580)
- DEPLOY: TECHLEAD executa (090/087); PM/PO aplica o go/no-go GMUD

> **NOTA:** o 105-TERMO-ACEITE permanece como aceite FINAL do projeto (FASE 6) — nunca gate por entrega. Tratativas de retorno (QA NO-GO → DEV; UAT divergência → 085; DEPLOY bloqueado → IMP-NN) conforme o 096.

> **SETUP DE AMBIENTE E FERRAMENTAS (tooling):** tarefas de infra/ferramentas das janelas DEV/QA (manifests, observabilidade, instalação de middleware/ETL e ferramentas de segurança) são orquestradas pelo roadmap companion **`PROMPT-ROADMAP-GENERATE-IMPLEMENTATION-TOOLING.md`** (v1.0, metodologia-independente), invocado pelo TECHLEAD no Bloco F — o PM/PO não executa tooling, apenas acompanha via 092/095. HMG/PROD sempre via GMUD (090).

## Sub-fase 3 — Contrato de Delegação ao TECHLEAD

A construção é executada pelo TECHLEAD via `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` (v7.1, modo `waterfall-discovery`, **Bloco E**). Este roadmap orquestra apenas o contrato: handoff, acompanhamento e recepção.

### 3.1 Handoff PM/PO → TECHLEAD

1. **Montar o pacote de demanda** via `GENERATE-092` (modo atualização):
   - Snapshot do 092 com a `CICLO-NN` ativa, itens `BL-NN` (status `A Fazer`/`Em Execução`) e `CR-NN` aprovadas do ciclo
   - Docs F1–F4 (a tabela de Inputs acima — pacote técnico mínimo: 010-FRD, 016, 020-SRS, 030-SAD, 035-HLD, 040-LLD, 041–044, 045-TEST-PLAN, 050-TEST-CASES, 060-EAP-WBS, 062, 065, 070, 086, 087, 088, 090)
   - Cruzamento por item: `CICLO-NN → FEAT-NN/UC-NN (010) → pacote EAP (060) → tarefas T-NNN`
2. **Transicionar** os `BL-NN` do ciclo para `Em Execução (delegado ao TECHLEAD)`
3. **Acionar** o TECHLEAD no modo `waterfall-discovery` do PROJECT-TECHNICAL-DEFINITIONS v7.1

### 3.2 Execução no TECHLEAD (resumo — sem duplicar steps)

O TECHLEAD executa o pipeline completo no roadmap dele: Fases migradas/validadas (F1–F19 no modo waterfall), **Bloco E** (contexto base + loop por ciclo — janela DEV) e **Bloco F** (janelas QA/UAT/DEPLOY), encerrando cada ciclo com o **trio 595** (empacotamento do retorno). A tabela de steps vive SOMENTE no v7.1 (Bloco E).

### 3.3 Recepção TECHLEAD → PM/PO

1. Receber o pacote `595-RETURN-PACKAGE-{CICLO-NN}.md` (COMPLIANCE no GATE-595)
2. Aplicar via `GENERATE-092` (modo atualização, `EXTRA_INPUTS` = pacote 595) → `GATE-092` → `FIX-092` se necessário
3. Registrar impedimentos propostos como `IMP-NN` via `GENERATE-093`
4. Alimentar 095-RELATORIO-QUALIDADE com as evidências (seção 4 do pacote)
5. Encaminhar pedidos de mudança ao 085 (escopo/PERT/stack)
6. Revalidação humana (P1–P4 do roadmap master)

### 3.4 Regras do contrato

- **Ownership:** TECHLEAD nunca edita 092/093/095/085/088 — "TECHLEAD propõe, PM/PO aplica"
- **Vocabulário:** tradução na invocação (US → UC-NN do 010; sprint → CICLO-NN do 092) — nunca reescrever prompts reusados
- **Git:** branch por ciclo `feature/sprint-NN-<slug>` executado pelo TECHLEAD no repositório da solução; PR com code review (086) e CI verde (087)
- **Débito técnico:** `DT-XXX` entra no pacote 595 como proposta de CR Técnico (`CR-NN`)

## Sub-fase 4 — Documentação de Suporte e Evidências

| Doc | Quando | Notas |
|---|---|---|
| 095-RELATORIO-QUALIDADE | estrutura criada na F3; **alimentado com evidências ao longo da FASE 5** | evidências chegam via pacote 595 (seção 4), aplicadas na recepção (sub-fase 3.3) |
| 097-MANUAIS-USUARIO | durante a execução | upstream 003/010/016 — telas e jornadas reais |
| 100-MANUAIS-OPERACIONAIS | durante a execução | upstream 041/044/087/090 — runbooks reais |

---

## Integração com WATERFALL-ESTIMATION

Se a EAP mudar via 085 (mudança de escopo aprovada), o orquestrador alerta o humano e oferece **reexecutar o WATERFALL-ESTIMATION modo DOWNSTREAM/REFINEMENT** — nova PERT atualiza 065/070 e, consequentemente, o 088 (via 092). A mudança reflete no **próximo** pacote de demanda entregue ao TECHLEAD (nunca em ciclo em andamento).

## Git Workflow da Execução

- Branch por ciclo: `feature/sprint-NN-<slug>`; PR obrigatório com code review (086) e CI verde (087)
- O commit de encerramento (git add/commit/push/PR do roadmap master) migra para o final da FASE 6

## Localização dos Prompts

```
.specs/prompts/
├── PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md      ← roadmap master (6 fases)
├── PROMPT-ROADMAP-GENERATE-WATERFALL-EXECUTION.md              ← ESTE ROADMAP (FASE 5 — escopo PM/PO)
├── PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md    ← roadmap TECHLEAD (v7.1 — modo waterfall-discovery, Bloco E)
├── project-documents-waterfall/                                ← triads WATERFALL (092, 093 e demais docs)
├── project-technical-definitions/                              ← pipeline TECHLEAD + trio 595-TECHLEAD-RETURN-PACKAGE
├── sprint-artefacts/                                           ← reuso (Bloco E do TECHLEAD): contexto base + artefatos de ciclo
├── sprint-tecnhnical-implementation/                           ← reuso (Bloco E do TECHLEAD): execução, QA-revisor, PR, débito técnico
└── technical-solutions/                                        ← reuso: tarefas por solução técnica
```

---

🤖 *Roadmap gerado pelo Waterfall Orchestrator v3.3. v2.3 (17/08/2026): + Sub-fase 1.5 — Planejamento do Ciclo de Entrega (CICLO-NN) com pauta de 9 pontos e o 595 do ciclo anterior como entrada obrigatória; renomeação global FILA-NN → CICLO-NN (18 arquivos). Histórico v2.2: + tooling de ambiente via IMPLEMENTATION-TOOLING v1.0; referências TECHLEAD atualizadas para v7.1. Histórico v2.1: Janelas de Entrega definidas no 096 e orquestradas pelo TECHLEAD (Bloco F) — aceite por entrega via DE-ACORDO; 105 permanece aceite final (FASE 6).*
