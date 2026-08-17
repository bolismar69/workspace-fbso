# PROMPT: ROADMAP DE ESTIMATIVA WATERFALL — UPSTREAM/DISCOVERY (ROM ±50%) + DOWNSTREAM/REFINEMENT (PERT ±15-25%)
## Versão: 1.0 — Bootstrap Inteligente, Fluxo de Dados Explícito, Validação Soberana Humana (HITL) e Git Workflow Automatizado

Atue como um Especialista em Gestão de Processos (BPM), PMO e Arquiteto de Soluções, especializado em metodologia WATERFALL e estimativas de projeto (ROM, PERT, Three-Point Estimation).

Preciso que você execute um roadmap de **cálculo de estimativas** em 2 modos independentes, integrados à sequência de documentos WATERFALL:

- **Modo UPSTREAM/DISCOVERY (ROM ±50%):** Estimativa de ordem de grandeza após a arquitetura de alto nível (HLD), para decisão GO/NO-GO do Comitê de Governança.
- **Modo DOWNSTREAM/REFINEMENT (PERT ±15-25%):** Estimativa de precisão após o design detalhado (LLD) e a decomposição de trabalho (EAP/WBS), para firmar compromisso de prazo, capacidade e custo.

**Base técnica:**
- `.specs/standards/DTA-Engine-de-Bidding-e-Estimativas.md` — Schema original e regras de ouro
- `.specs/standards/DTA-VALIDATION-STANDARDS.md` — Regras, fórmulas e padrões de validação (§1 ROM, §2 DTA, §5 Internal Baseline)
- `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` — Baseline corporativa de stack

**Regra Crítica de Execução (Gating Rule):** O processo é estritamente sequencial. Nenhum artefato subsequente pode ser iniciado sem que o artefato atual esteja marcado como `[STATUS: COMPLIANCE]` e aprovado explicitamente pelo humano.

---

## REGRAS DE OURO (7 REGRAS — NÃO NEGOCIÁVEIS)

1. **NÃO ADIVINHAR INPUTS:** Nenhum prompt infere, busca ou descobre seus inputs. Todo parâmetro é passado explicitamente por este orquestrador.
2. **TODO GENERATE TEM GATE+FIX:** Cada artefato tem exatamente um trio GENERATE, GATE e FIX.
3. **FLUXO DE DADOS EXPLÍCITO:** `ARTIFACT_PATH` flui GENERATE→GATE→FIX. `VIOLATIONS[]` flui GATE→FIX. Sempre como parâmetros nomeados.
4. **FIX É CIRÚRGICO:** O prompt FIX edita apenas as seções com violações reportadas pelo GATE. Nunca regenera ou recria o artefato.
5. **STATUS INICIAL: EM ANÁLISE:** Ao criar o artefato, GENERATE escreve `[STATUS: Em análise]` no cabeçalho.
6. **EM REVISÃO DURANTE GATE/FIX:** GATE e FIX alteram o status para `[STATUS: Em revisão]`.
7. **SÓ AVANÇA COM COMPLIANCE:** O roadmap só avança para o próximo artefato quando o atual estiver marcado `[STATUS: COMPLIANCE]` e o humano confirmar.

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP (FASE 0)

### Tabela de Inputs

| Variável | Obrig. | Descrição | Exemplo |
|---|---|---|---|
| `PROJECT_PATH` | ✅ | Caminho base onde os projetos de negócio residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID_NAME` | ✅ | Identificador completo do projeto | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `WATERFALL_ESTIMATION_MODE` | ❌ | Modo de estimativa: `upstream-discovery` ou `downstream-refinement`. Se não informado, o Bootstrap detecta e pergunta | `upstream-discovery` |
| `PROJECT_DOCUMENTS_INPUTS` | ❌ | Lista de caminhos para documentos brutos de entrada adicionais | `[]` |
| `PROJECT_PROMPT_INPUTS` | ❌ | **(Diretiva comportamental)** Checkpoint HITL: sempre solicitar ao usuário, no início e durante a execução, se deseja fornecer informações adicionais | `{checkpoint HITL}` |
| `PROJECT-TEAM-SKILLS-MAP` | ❌ | Skills necessários para o time de implementação | `{obter e validar com usuario}` |
| `PROJECT-TEAM-CAPACITY` | ❌ | Capacidade esperada do time (seniores, plenos, juniores, duração) | `{obter e validar com usuario}` |
| `PROJECT-STACK` | ❌ | Stack tecnológica da solução. Baseline corporativa em `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` | `{obter do contexto, complementar com usuario, validar contra padroes}` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
WATERFALL_ESTIMATION_PATH     = PROJECT_COMPLETE_PATH_NAME + "/waterfall-estimation"
WATERFALL_DOCS_PATH           = PROJECT_COMPLETE_PATH_NAME
```

### Modos de Operação

| Modo | Pasta | Gatilho | Precisão | Finalidade |
|---|---|---|---|---|
| **Mode 1 — `upstream-discovery`** | `waterfall-estimation/` | Após Doc #06 (HLD) COMPLIANCE | ROM ±50% | GO/NO-GO Governance |
| **Mode 2 — `downstream-refinement`** | `waterfall-estimation/` | Após Doc #07 (LLD) + Doc #11 (EAP/WBS) COMPLIANCE | PERT ±15-25% | Compromisso de prazo, capacidade e custo |

### Pré-condições por Modo

| Modo | Documentos WATERFALL que DEVEM existir e estar COMPLIANCE |
|---|---|
| `upstream-discovery` | `01-PROJECT-CHARTER`, `02-BRD`, `05-SAD`, `06-HLD` |
| `downstream-refinement` | `03-SRS`, `04-RTM`, `07-LLD`, `11-EAP-WBS` |

---

## ARQUITETURA DE FASES

O roadmap é organizado em **7 fases** distribuídas em **2 modos independentes**:

```
FASE 0: BOOTSTRAP (detecta modo + pergunta ao humano)
  │
  ├──▶ MODO UPSTREAM/DISCOVERY (ROM ±50%)
  │      Fase 1 → Fase 2 → Fase 3
  │      ⛔ Barreira UPSTREAM → GO/NO-GO Governance
  │
  └──▶ MODO DOWNSTREAM/REFINEMENT (PERT ±15-25%)
         Fase 4 → Fase 5 → Fase 6 → Fase 7
         ⛔ Barreira DOWNSTREAM → Alimenta Docs #12 Cronograma + #13 Orçamento
```

**Os modos são independentes.** Um projeto pode executar apenas o UPSTREAM, apenas o DOWNSTREAM, ou ambos em sequência (UPSTREAM primeiro para GO/NO-GO, depois DOWNSTREAM para detalhamento).

---

## FASES DO ROADMAP

### Fase 0 — Bootstrap Inteligente

Workflow:
1. **Solicitar inputs ao usuário** (se não fornecidos)
2. **Auditar documentos WATERFALL existentes:**
   - Verificar existência e status COMPLIANCE de: 01-Charter, 02-BRD, 05-SAD, 06-HLD
   - Verificar existência e status COMPLIANCE de: 03-SRS, 04-RTM, 07-LLD, 11-EAP/WBS
3. **Detectar modo(s) disponível(is):**
   - Se 01, 02, 05, 06 todos COMPLIANCE → **Mode 1 (`upstream-discovery`) disponível**
   - Se 03, 04, 07, 11 todos COMPLIANCE → **Mode 2 (`downstream-refinement`) disponível**
   - Se ambos os conjuntos completos → **Ambos disponíveis** — perguntar qual executar (ou ambos em sequência)
   - Se nenhum → informar quais docs faltam e encerrar
4. **Humano decide** qual modo (`WATERFALL_ESTIMATION_MODE`)
5. **Se modo `downstream-refinement`, validar que o UPSTREAM_DOCS inclui o EAP/WBS** — o PERT depende dos pacotes de trabalho decompostos
6. **Coletar `PROJECT-TEAM-SKILLS-MAP`, `PROJECT-TEAM-CAPACITY` e `PROJECT-STACK`:**
   - Apresentar stack corporativa FBSO como baseline
   - Questionar sobre skills adicionais
   - Questionar sobre capacidade e senioridade do time
   - **Aguardar validação do usuário**
7. **Criar estrutura:** `mkdir -p {WATERFALL_ESTIMATION_PATH}`
8. **Auditar artefatos de estimativa existentes** no diretório `waterfall-estimation/`
9. **Apresentar resumo e iniciar** a primeira fase pendente

---

### MODO UPSTREAM/DISCOVERY — ROM ±50%

**Gatilho:** Documentos #01 (Charter), #02 (BRD), #05 (SAD), #06 (HLD) estão COMPLIANCE.
**Finalidade:** Submeter ao Comitê de Governança para decisão GO/NO-GO.

#### Fase 1 — WATERFALL-ESTIMATION-UPSTREAM-ROM.md

**Objetivo:** Calcular a estimativa ROM (Rough Order of Magnitude ±50%) baseada na arquitetura de alto nível (HLD) e no contexto de negócio (Charter + BRD).

**Metodologia:**
- **Bottom-Up por componente arquitetural:** Cada container/serviço do HLD é uma unidade de estimativa
- **Fórmula ROM:** `ROM = Estimativa_Provável × (1 ± 0.50)`
- **Faixa:** `[ROM_min = 0.50 × Provável, ROM_max = 1.50 × Provável]`
- **Dimensões por componente:** Desenvolvimento, QA, Arquitetura, DevOps/SRE, Gestão
- **Premissas documentadas:** Cada componente lista explicitamente as premissas assumidas
- **Riscos de estimativa:** Fatores que podem empurrar a estimativa para o limite superior da faixa

**Inputs consumidos:**
- `01-PROJECT-CHARTER` → escopo macro, milestones, orçamento macro
- `02-BRD` → requisitos de negócio, restrições
- `05-SAD` → 6 visões arquiteturais, decisões de design
- `06-HLD` → containers, matriz de integração, ADRs, topologia
- `PROJECT-STACK` → tecnologias e frameworks
- `PROJECT-TEAM-CAPACITY` → seniores, plenos, juniores

**Output esperado:**
- Seções: (1) Escopo Estimado, (2) Matriz de Componentes × Dimensões, (3) ROM Consolidado, (4) Premissas por Componente, (5) Riscos e Fatores de Ajuste, (6) Faixa de Confiança, (7) Recomendação para Governança
- Estimativa em horas e valor financeiro (usando taxas padrão ou informadas)
- NÃO substitui documentos WATERFALL — é um artefato complementar

**Pipeline:** `PROMPT-GENERATE-WATERFALL-ESTIMATION-UPSTREAM-ROM.md` → Gate → Fix → COMPLIANCE

#### Fase 2 — ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md

**Objetivo:** Congelar formalmente o escopo que foi considerado na estimativa ROM. Este snapshot serve como evidência para o comitê de governança do que está (e do que NÃO está) coberto pela estimativa.

**Metodologia:**
- Listar todos os componentes/containers do HLD que foram estimados
- Listar explicitamente o que está FORA do escopo
- Vincular cada item de escopo ao documento WATERFALL de origem (Charter §X, BRD §Y, HLD §Z)
- Documentar versão e data de cada documento fonte usado

**Inputs consumidos:**
- `01-PROJECT-CHARTER`, `02-BRD`, `05-SAD`, `06-HLD`
- `WATERFALL-ESTIMATION-UPSTREAM-ROM.md` (Fase 1 — gerado acima)

**Output esperado:**
- Seções: (1) Itens de Escopo Estimados, (2) Exclusões Explícitas, (3) Matriz de Rastreabilidade (escopo × documento fonte), (4) Versões dos Documentos Fonte, (5) Premissas de Escopo

**Pipeline:** `PROMPT-GENERATE-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md` → Gate → Fix → COMPLIANCE

#### Fase 3 — GOVERNANCE-ROM-REPORT.md

**Objetivo:** Produzir relatório executivo sintético para o Comitê de Governança decidir GO/NO-GO.

**Metodologia:**
- Sumário executivo (máximo 1 página) com recomendação
- Tabela de custo total (ROM min, provável, max)
- Timeline macro (meses ou sprints de alto nível)
- Principais riscos e mitigadores
- Premissas críticas que, se alteradas, invalidam a estimativa
- Decisão solicitada ao comitê: GO (aprovar financiamento) ou NO-GO (arquivar/redirecionar)

**Inputs consumidos:**
- `WATERFALL-ESTIMATION-UPSTREAM-ROM.md` (Fase 1)
- `ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md` (Fase 2)
- `01-PROJECT-CHARTER` (para milestones e orçamento macro)

**Output esperado:**
- Seções: (1) Sumário Executivo, (2) Escopo e Premissas, (3) Estimativa Financeira (ROM), (4) Timeline Macro, (5) Riscos e Mitigadores, (6) Recomendação, (7) Decisão Solicitada (GO/NO-GO)

**Pipeline:** `PROMPT-GENERATE-WATERFALL-ESTIMATION-GOVERNANCE-ROM-REPORT.md` → Gate → Fix → COMPLIANCE

**⛔ Barreira UPSTREAM:** Após Fase 3 COMPLIANCE, o relatório é submetido ao Comitê de Governança para decisão GO/NO-GO.

---

### MODO DOWNSTREAM/REFINEMENT — PERT ±15-25%

**Gatilho:** Documentos #03 (SRS), #04 (RTM), #07 (LLD), #11 (EAP/WBS) estão COMPLIANCE.
**Finalidade:** Alimentar com precisão cirúrgica os Documentos #12 (Cronograma/Gantt) e #13 (Orçamento).

#### Fase 4 — WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md

**Objetivo:** Calcular a estimativa PERT (Program Evaluation and Review Technique ±15-25%) baseada no design detalhado (LLD) e nos pacotes de trabalho (EAP/WBS).

**Metodologia:**
- **Bottom-Up por pacote de trabalho da EAP/WBS:** Cada pacote é estimado individualmente
- **Three-Point Estimation por pacote:**
  - `O = Otimista` (melhor cenário)
  - `M = Mais Provável` (cenário esperado)
  - `P = Pessimista` (pior cenário)
  - `E = (O + 4M + P) / 6` (PERT ponderado)
  - `σ = (P - O) / 6` (desvio padrão)
- **Consolidação por fase WATERFALL:** Agrupar estimativas por fase (Requisitos, Design, Testes, Planejamento, Deploy)
- **Cálculo do caminho crítico:** Baseado nas dependências do LLD e sequenciamento da EAP
- **Dimensões por pacote:** Desenvolvimento, QA, Arquitetura, DevOps/SRE, Gestão
- **Validação DTA interna:** QA ≥ 25% do Dev, Arquitetura ≥ 5% do total
- **Estimativa independente:** O PERT é calculado do zero — NÃO usa ROM upstream como baseline ou ponto de partida

**Inputs consumidos:**
- `03-SRS` → requisitos funcionais e não-funcionais detalhados
- `04-RTM` → rastreabilidade requisitos ↔ componentes
- `07-LLD` → APIs, tabelas, diagramas de classe/sequência, integrações
- `11-EAP/WBS` → pacotes de trabalho decompostos (nível 3+)
- `PROJECT-STACK` → tecnologias e frameworks
- `PROJECT-TEAM-SKILLS-MAP` → skills disponíveis
- `PROJECT-TEAM-CAPACITY` → seniores, plenos, juniores, duração

**Output esperado:**
- Seções: (1) Escopo Estimado, (2) Matriz de Pacotes EAP × Dimensões com Three-Point, (3) PERT Consolidado por Fase WATERFALL, (4) Caminho Crítico, (5) Desvio Padrão e Faixa de Confiança, (6) Validação DTA Interna (QA ≥ 25%, Arch ≥ 5%), (7) Premissas por Pacote, (8) Independência (declaração explícita de que esta estimativa não usou ROM upstream)

**Pipeline:** `PROMPT-GENERATE-WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md` → Gate → Fix → COMPLIANCE

#### Fase 5 — ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md

**Objetivo:** Congelar formalmente o escopo detalhado considerado na estimativa PERT. Este snapshot é a baseline contra a qual qualquer mudança de escopo será medida.

**Metodologia:**
- Listar todos os pacotes EAP/WBS estimados, com seus IDs e descrições
- Listar explicitamente o que está FORA do escopo
- Vincular cada pacote ao documento WATERFALL de origem (SRS §X, RTM §Y, LLD §Z)
- Documentar versão e data de cada documento fonte usado
- Incluir a declaração de independência da estimativa

**Inputs consumidos:**
- `03-SRS`, `04-RTM`, `07-LLD`, `11-EAP/WBS`
- `WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md` (Fase 4 — gerado acima)

**Output esperado:**
- Seções: (1) Pacotes EAP Estimados, (2) Exclusões Explícitas, (3) Matriz de Rastreabilidade (pacote × documento fonte), (4) Versões dos Documentos Fonte, (5) Independência da Estimativa, (6) Premissas de Escopo

**Pipeline:** `PROMPT-GENERATE-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md` → Gate → Fix → COMPLIANCE

#### Fase 6 — CRONOGRAMA-CALCULADO.md

**Objetivo:** Produzir cronograma detalhado derivado da estimativa PERT, pronto para ser consumido pelo Documento #12 (Cronograma/Gantt) do WATERFALL.

**Metodologia:**
- **Derivar durações do PERT:** Cada pacote EAP tem duração = E (PERT) / (tamanho_equipe × horas_dia)
- **Sequenciamento:** Usar dependências do LLD e da EAP/WBS
- **Caminho crítico:** Identificado e destacado
- **Marcos (Milestones):** Vinculados aos milestones do Project Charter
- **Alocação de recursos:** Por perfil (sênior, pleno, júnior) conforme capacidade informada
- **Formato compatível:** Estrutura compatível com o template do Doc #12 WATERFALL

**Inputs consumidos:**
- `WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md` (Fase 4)
- `ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md` (Fase 5)
- `01-PROJECT-CHARTER` (milestones)
- `PROJECT-TEAM-CAPACITY` (tamanho e senioridade da equipe)

**Output esperado:**
- Seções: (1) Lista de Atividades com Durações, (2) Sequenciamento e Dependências, (3) Caminho Crítico, (4) Cronograma (datas), (5) Diagrama de Gantt (textual), (6) Marcos (Milestones), (7) Alocação de Recursos por Período

**Pipeline:** `PROMPT-GENERATE-WATERFALL-ESTIMATION-CRONOGRAMA-CALCULADO.md` → Gate → Fix → COMPLIANCE

#### Fase 7 — ORCAMENTO-CALCULADO.md

**Objetivo:** Produzir orçamento detalhado derivado da estimativa PERT, pronto para ser consumido pelo Documento #13 (Orçamento) do WATERFALL.

**Metodologia:**
- **Custo por recurso:** Horas PERT × taxa horária por perfil (sênior/pleno/júnior)
- **Custo de infraestrutura:** Derivado do LLD (containers, bancos, cloud)
- **Custo de licenças:** Conforme stack validada
- **Curva S:** Custo acumulado ao longo do tempo
- **Reserva de contingência:** Baseada no desvio padrão do PERT: `Contingência = f(σ, nível_confiança)`
- **Fluxo de caixa projetado:** Por mês/sprint
- **Formato compatível:** Estrutura compatível com o template do Doc #13 WATERFALL

**Inputs consumidos:**
- `WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md` (Fase 4)
- `CRONOGRAMA-CALCULADO.md` (Fase 6)
- `PROJECT-STACK` (para licenças e custos de infra)
- `PROJECT-TEAM-CAPACITY` (para taxas horárias)

**Output esperado:**
- Seções: (1) Custos por Pacote EAP, (2) Custos por Recurso (RH, Infra, Licenças, Serviços), (3) Curva S (Custo Acumulado), (4) Reserva de Contingência (baseada em σ PERT), (5) Fluxo de Caixa Projetado, (6) Comparativo com Faixa ROM (se UPSTREAM executado)

**Pipeline:** `PROMPT-GENERATE-WATERFALL-ESTIMATION-ORCAMENTO-CALCULADO.md` → Gate → Fix → COMPLIANCE

**⛔ Barreira DOWNSTREAM:** Após Fase 7 COMPLIANCE, os artefatos `CRONOGRAMA-CALCULADO.md` e `ORCAMENTO-CALCULADO.md` são injetados como `UPSTREAM_DOCS` adicionais nos GENERATEs WATERFALL #12 (Cronograma) e #13 (Orçamento).

---

## LOOP DE EXECUÇÃO POR ARTEFATO

Para cada artefato `{ARTIFACT}` na ordem sequencial do modo corrente:

### STEP 1: Computar inputs para GENERATE

Cada GENERATE recebe APENAS as variáveis de domínio relevantes ao seu escopo. NUNCA passar variáveis que o artefato não utiliza.

```
ARTIFACT_PATH          = {WATERFALL_ESTIMATION_PATH}/{ARTIFACT-SLUG}.md
UPSTREAM_DOCS          = [lista de paths de documentos WATERFALL upstream já em COMPLIANCE]
INTERNAL_UPSTREAM      = [lista de artefatos de estimativa já gerados neste roadmap]
WATERFALL_DOCS_PATH    = {PROJECT_COMPLETE_PATH_NAME}
SKILLS                 = [lista de skills para este artefato — vide tabela abaixo]
```

### STEP 2: Invocar GENERATE

Invocar `waterfall-estimation/PROMPT-GENERATE-WATERFALL-ESTIMATION-{ARTIFACT-SLUG}.md` passando **explicitamente** os parâmetros computados no STEP 1:
- `ARTIFACT_PATH`, `PROJECT_ID_NAME`, `UPSTREAM_DOCS`, `INTERNAL_UPSTREAM`, `SKILLS`
- **+ variáveis de domínio** conforme a Matriz de Roteamento:
  - `PROJECT-STACK`, `PROJECT-TEAM-SKILLS-MAP`, `PROJECT-TEAM-CAPACITY`

**Antes de cada GENERATE:** Checkpoint HITL — perguntar ao usuário:
> "Antes de gerar [{ARTIFACT}]: deseja fornecer informações adicionais, novos contextos ou ajustes de escopo?"

### STEP 3: Invocar GATE

Invocar `waterfall-estimation/PROMPT-GATE-WATERFALL-ESTIMATION-{ARTIFACT-SLUG}.md` passando **explicitamente**:
- `ARTIFACT_PATH` (recebido do GENERATE)
- `CHECKLIST` (definido no próprio prompt GATE)

GATE deve:
- Ler apenas o arquivo em `ARTIFACT_PATH`
- Atualizar status para `[STATUS: Em revisão]`
- Retornar `{PASS}` ou `{FAIL, VIOLATIONS: [{section, description, severity}]}`

### STEP 4a: Se GATE retornar FAIL → Invocar FIX

Invocar `waterfall-estimation/PROMPT-FIX-WATERFALL-ESTIMATION-{ARTIFACT-SLUG}.md` passando **explicitamente**:
- `ARTIFACT_PATH` (mesmo arquivo)
- `VIOLATIONS[]` (lista exata de não-conformidades do GATE)

FIX deve:
- Editar **apenas** as seções em `VIOLATIONS[]`
- Manter status como `[STATUS: Em revisão]`
- Retornar `{ARTIFACT_PATH}`
- → Voltar ao STEP 3 (re-executar GATE)

### STEP 4b: Se GATE retornar PASS → Validação Humana + Checkpoint HITL

Apresentar o artefato e fazer as perguntas padrão:

> **P1:** "O conteúdo deste artefato está aderente às necessidades do projeto?"
> **P2:** "Existem novos documentos de entrada que devem ser incorporados?"
> **P3:** "Há novas informações textuais, mudanças de escopo ou ajustes técnicos?"
> **P4 (HITL):** "Deseja fornecer novas informações antes de prosseguir?"

- Se humano aprovar (SIM para P1, NÃO para P2/P3/P4): `[STATUS: COMPLIANCE]`, artefato congelado, próximo liberado
- Se humano fornecer novos inputs: Voltar ao STEP 2 (re-executar GENERATE com novo contexto)

---

## MATRIZ DE SKILLS POR ARTEFATO

| # | Artefato | Skills Primárias | Fallback |
|---|---|---|---|
| 1 | UPSTREAM-ROM | `project-estimation`, `estimate-builder`, `senior-architect` | ✅ |
| 2 | SCOPE-SNAPSHOT-UPSTREAM | `gap-analysis`, `business-analyst` | ✅ |
| 3 | GOVERNANCE-ROM-REPORT | `presentation-creation`, `senior-pm` | ✅ |
| 4 | DOWNSTREAM-PERT | `project-estimation`, `afrexai-construction-estimator`, `senior-architect` | ✅ |
| 5 | SCOPE-SNAPSHOT-DOWNSTREAM | `gap-analysis`, `business-analyst` | ✅ |
| 6 | CRONOGRAMA-CALCULADO | `roadmap-planning`, `project-estimation`, `senior-pm` | ✅ |
| 7 | ORCAMENTO-CALCULADO | `project-estimation`, `ads-budget`, `senior-pm` | ✅ |

---

## INTEGRAÇÃO COM O ROADMAP WATERFALL

### Upstream Integration Point (após Doc #06 HLD)

```
WATERFALL Docs (Fase 3 — Design e Arquitetura)
  #05-SAD ✅ → #06-HLD ✅
                     │
                     ▼
         ┌─────────────────────────┐
         │ WATERFALL-ESTIMATION    │
         │ Modo: upstream-discovery│
         │ Fases 1 → 2 → 3        │
         │ Output: ROM + GO/NO-GO  │
         └────────────┬────────────┘
                      │
         GO ✅ → Continua WATERFALL (#07-LLD)
         NO-GO ❌ → Projeto Cancelado / Arquivado
```

### Downstream Integration Point (após Doc #07 LLD + Doc #11 EAP/WBS)

```
WATERFALL Docs (Fase 3 + Fase 5)
  #07-LLD ✅ → #11-EAP/WBS ✅
                │
                ▼
    ┌─────────────────────────┐
    │ WATERFALL-ESTIMATION    │
    │ Modo: downstream-       │
    │ refinement              │
    │ Fases 4 → 5 → 6 → 7    │
    │ Output: PERT + Crono +  │
    │ Orçamento               │
    └────────────┬────────────┘
                 │
    Cronograma-Calculado → alimenta Doc #12
    Orcamento-Calculado  → alimenta Doc #13
```

### Atualização do UPSTREAM_DOCS nos GENERATEs WATERFALL

Os GENERATEs #12 (Cronograma) e #13 (Orçamento) passam a receber um `UPSTREAM_DOCS` adicional:

| Documento WATERFALL | UPSTREAM_DOCS Atualizado |
|---|---|
| #12 Cronograma/Gantt | `[01-Charter, 11-EAP/WBS, 09-TEST-CASES, **CRONOGRAMA-CALCULADO**]` |
| #13 Orçamento | `[01-Charter, 11-EAP/WBS, 12-Cronograma, 08-TEST-PLAN, **ORCAMENTO-CALCULADO**]` |

---

## ESTRUTURA DE DIRETÓRIOS GERADA

```
business-inputs/business-projects/{PROJECT_ID_NAME}/
├── 01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md
├── ... (demais 19 documentos WATERFALL)
│
└── waterfall-estimation/
    ├── WATERFALL-ESTIMATION-UPSTREAM-ROM.md         (F1)
    ├── ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md        (F2)
    ├── GOVERNANCE-ROM-REPORT.md                     (F3)
    ├── WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md      (F4)
    ├── ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md      (F5)
    ├── CRONOGRAMA-CALCULADO.md                      (F6)
    └── ORCAMENTO-CALCULADO.md                       (F7)
```

**Nota:** Os modos compartilham o mesmo diretório `waterfall-estimation/`. Se ambos os modos forem executados, os 7 artefatos coexistem no mesmo diretório. Não há conflito porque cada artefato tem nome único.

---

## EFEITOS CASCATA

Artefatos de estimativa downstream que dependem de artefatos modificados:

| Se modificar... | Impacta (regenerar + revalidar)... |
|---|---|
| F1 (UPSTREAM-ROM) | F2 (SCOPE-SNAPSHOT), F3 (GOVERNANCE-REPORT) |
| F2 (SCOPE-SNAPSHOT-UPSTREAM) | F3 (GOVERNANCE-REPORT) |
| F4 (DOWNSTREAM-PERT) | F5 (SCOPE-SNAPSHOT-DOWNSTREAM), F6 (CRONOGRAMA), F7 (ORCAMENTO) |
| F5 (SCOPE-SNAPSHOT-DOWNSTREAM) | F6 (CRONOGRAMA), F7 (ORCAMENTO) |
| F6 (CRONOGRAMA-CALCULADO) | F7 (ORCAMENTO), Doc WATERFALL #12 (Cronograma) |
| F7 (ORCAMENTO-CALCULADO) | Doc WATERFALL #13 (Orçamento) |

**Efeito cruzado WATERFALL → Estimativa:** Se um documento WATERFALL upstream for modificado, todos os artefatos de estimativa que o consomem devem ser regenerados:
- Modificação no `06-HLD` → F1-F3 devem ser regenerados
- Modificação no `07-LLD` ou `11-EAP/WBS` → F4-F7 devem ser regenerados

---

## FINALIZAÇÃO — GIT WORKFLOW

Quando todos os artefatos do modo corrente estiverem em COMPLIANCE e o humano confirmar a conclusão, executar o pipeline Git abaixo:

```
WORK_BRANCH = "feature/" + PROJECT_ID_NAME + "-waterfall-estimation-" + WATERFALL_ESTIMATION_MODE
```

### Passo F.1 — Git Add e Commit

```bash
git add -A
git commit -m "docs: estimativa WATERFALL ${WATERFALL_ESTIMATION_MODE} — ${PROJECT_ID_NAME}

- Artefatos de estimativa gerados e validados
- Modo: ${WATERFALL_ESTIMATION_MODE}
- Status: COMPLIANCE em todos os artefatos
- Gerado pelo Waterfall Estimation Orchestrator v1.0
- Branch: ${WORK_BRANCH}

Co-Authored-By: Claude <noreply@anthropic.com>"
```

### Passo F.2 — Git Push e PR

```bash
git push origin ${WORK_BRANCH}
gh pr create --base main --head ${WORK_BRANCH} \
  --title "docs: estimativa WATERFALL ${WATERFALL_ESTIMATION_MODE} — ${PROJECT_ID_NAME}" \
  --body "Estimativa WATERFALL ${WATERFALL_ESTIMATION_MODE} para ${PROJECT_ID_NAME}. Artefatos validados."
gh pr merge --merge --delete-branch
```

### Passo F.3 — Cleanup Local

```bash
git checkout main
git branch -d ${WORK_BRANCH}
```

---

## LOCALIZAÇÃO DOS PROMPTS

```
.specs/prompts/
├── PROMPT-ROADMAP-GENERATE-WATERFALL-ESTIMATION.md   ← ESTE ORQUESTRADOR
│
└── waterfall-estimation/
    ├── PROMPT-GENERATE-WATERFALL-ESTIMATION-UPSTREAM-ROM.md
    ├── PROMPT-GATE-WATERFALL-ESTIMATION-UPSTREAM-ROM.md
    ├── PROMPT-FIX-WATERFALL-ESTIMATION-UPSTREAM-ROM.md
    ├── PROMPT-GENERATE-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md
    ├── PROMPT-GATE-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md
    ├── PROMPT-FIX-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-UPSTREAM.md
    ├── PROMPT-GENERATE-WATERFALL-ESTIMATION-GOVERNANCE-ROM-REPORT.md
    ├── PROMPT-GATE-WATERFALL-ESTIMATION-GOVERNANCE-ROM-REPORT.md
    ├── PROMPT-FIX-WATERFALL-ESTIMATION-GOVERNANCE-ROM-REPORT.md
    ├── PROMPT-GENERATE-WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md
    ├── PROMPT-GATE-WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md
    ├── PROMPT-FIX-WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md
    ├── PROMPT-GENERATE-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md
    ├── PROMPT-GATE-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md
    ├── PROMPT-FIX-WATERFALL-ESTIMATION-SCOPE-SNAPSHOT-DOWNSTREAM.md
    ├── PROMPT-GENERATE-WATERFALL-ESTIMATION-CRONOGRAMA-CALCULADO.md
    ├── PROMPT-GATE-WATERFALL-ESTIMATION-CRONOGRAMA-CALCULADO.md
    ├── PROMPT-FIX-WATERFALL-ESTIMATION-CRONOGRAMA-CALCULADO.md
    ├── PROMPT-GENERATE-WATERFALL-ESTIMATION-ORCAMENTO-CALCULADO.md
    ├── PROMPT-GATE-WATERFALL-ESTIMATION-ORCAMENTO-CALCULADO.md
    ├── PROMPT-FIX-WATERFALL-ESTIMATION-ORCAMENTO-CALCULADO.md
    └── (21 prompts: 7 GENERATE + 7 GATE + 7 FIX)
```

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 04/08/2026 | Criação inicial: roadmap de estimativa WATERFALL com 2 modos (upstream-discovery ROM ±50%, downstream-refinement PERT ±15-25%), 7 fases, integração com docs WATERFALL #12 e #13. Design baseado nos roadmaps Upstream-Discovery e Downstream-Refinement existentes. | Time de Arquitetura |

---

🤖 *Roadmap gerado pelo Waterfall Estimation Orchestrator v1.0. Skills: project-estimation, estimate-builder, senior-pm.*
