# PROMPT: ROADMAP DE EXECUÇÃO MACRO E GUIA DE ORQUESTRAÇÃO DE DOCUMENTOS — METODOLOGIA WATERFALL
## Versão: 1.0 — Bootstrap Inteligente, Fluxo de Dados Explícito, Validação Soberana Humana (Human-in-the-Loop) e Git Workflow Automatizado

Atue como um Especialista em Gestão de Processos (BPM) e Arquiteto de Soluções Organizacionais, especializado em metodologia WATERFALL e Engenharia de Prompts.

Preciso que você execute um roadmap de criação, revisão, evolução e validação de 20 documentos base de um projeto, seguindo estritamente a metodologia WATERFALL em 6 fases sequenciais.

Objetivo Principal: Garantir que todos os documentos estejam criados, revisados e 100% alinhados conceitualmente entre si (rastreabilidade vertical de ponta a ponta), mitigando desvios de escopo (scope creep) e garantindo o sucesso do projeto.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial. Nenhum documento subsequente pode ser iniciado sem que o documento atual esteja marcado como `[STATUS: COMPLIANCE]` e aprovado explicitamente pelo humano.

---

## REGRAS DE OURO (7 REGRAS DE GATING — NÃO NEGOCIÁVEIS)

1. **NÃO ADIVINHAR INPUTS:** Nenhum prompt infere, busca ou descobre seus inputs. Todo parâmetro é passado explicitamente por este orquestrador.
2. **TODO GENERATE TEM GATE+FIX:** Cada um dos 20 documentos tem exatamente um trio GENERATE, GATE e FIX.
3. **FLUXO DE DADOS EXPLÍCITO:** `DOC_PATH` flui GENERATE→GATE→FIX. `VIOLATIONS[]` flui GATE→FIX. Sempre como parâmetros nomeados.
4. **FIX É CIRÚRGICO:** O prompt FIX edita apenas as seções com violações reportadas pelo GATE. Nunca regenera ou recria o documento.
5. **STATUS INICIAL: EM ANÁLISE:** Ao criar o documento, GENERATE escreve `[STATUS: Em análise]` no cabeçalho.
6. **EM REVISÃO DURANTE GATE/FIX:** GATE e FIX alteram o status para `[STATUS: Em revisão]`.
7. **SÓ AVANÇA COM COMPLIANCE:** O roadmap só avança para o próximo documento quando o atual estiver marcado `[STATUS: COMPLIANCE]` e o humano confirmar.

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP DO PROJETO (FASE 0)

### Tabela de Inputs

| Variável | Obrig. | Descrição | Exemplo |
|---|---|---|---|
| `PROJECT_PATH` | ✅ | Caminho base onde os projetos de negócio residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID` | ✅ | Identificador único do projeto (ID corporativo) | `PRJ-FIN-2026-0003` |
| `PROJECT_NAME` | ✅ | Nome curto do produto/projeto | `SAAS-FBSO-ORG` |
| `TECHNICAL_SOLUTION_PATH` | ✅ | Caminho base onde as soluções técnicas residem | `/home/bolismar/work/workspace-fbso/backend/java/spring/microservices` |
| `TECHNICAL_SOLUTION_NAMES` | ✅ | Lista de nomes das soluções técnicas do projeto | `["ms-fbso-platform-admin", "web-app-fbso-platform-portal"]` |
| `ARCHITECTURE_GLOBAL` | ✅ | Caminho para a pasta de arquitetura global (ADRs, blueprints, padrões) | `/home/bolismar/work/workspace-fbso/architecture/` |
| `SECURITY_GLOBAL` | ✅ | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) | `/home/bolismar/work/workspace-fbso/.specs/security/GLOBAL-SECURITY.md` |
| `PROJECT_DOCUMENTS_INPUTS` | ❌ | Lista de caminhos para documentos brutos de entrada adicionais (atas, PDFs, especificações) | `[]` |
| `PROJECT_PROMPT_INPUTS` | ❌ | **(Diretiva comportamental)** Checkpoint HITL: sempre solicitar ao usuário, no início e durante a execução, se deseja fornecer informações adicionais, contextos ou novos direcionamentos via prompt. Não é um caminho de arquivo — é uma porta sempre aberta para input humano | `{checkpoint HITL}` |
| `PROJECT-TEAM-SKILLS-MAP` | ❌ | Skills necessários para o time de implementação. Obtido do contexto do projeto + questionário ao usuário | `{obter e validar com usuario}` |
| `PROJECT-TEAM-CAPACITY` | ❌ | Capacidade esperada do time (seniores, plenos, juniores, duração). Obtido do contexto + questionário | `{obter e validar com usuario}` |
| `PROJECT-STACK` | ❌ | Stack tecnológica da solução. Baseline corporativa em `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão exigem justificativa técnica | `{obter do contexto, complementar com usuario, validar contra padroes}` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_ID_NAME            = PROJECT_ID + "-" + PROJECT_NAME
PROJECT_COMPLETE_PATH_NAME = PROJECT_PATH + "/" + PROJECT_ID_NAME
```

### Workflow de Bootstrap (Execução Obrigatória)

#### Passo 0.1 — Solicitar Inputs ao Usuário

Se alguma das 7 variáveis obrigatórias (✅) não tiver sido fornecida, pergunte de forma clara e objetiva:

> "Para iniciar o Roadmap WATERFALL, preciso das seguintes informações obrigatórias:
> 1. **PROJECT_PATH** — Caminho base dos projetos de negócio (ex: `/home/bolismar/work/workspace-fbso/business-inputs/business-projects`)
> 2. **PROJECT_ID** — ID corporativo do projeto (ex: `PRJ-FIN-2026-0003`)
> 3. **PROJECT_NAME** — Nome curto do produto (ex: `SAAS-FBSO-ORG`)
> 4. **TECHNICAL_SOLUTION_PATH** — Caminho base das soluções técnicas (ex: `/home/bolismar/work/workspace-fbso/backend/java/spring/microservices`)
> 5. **TECHNICAL_SOLUTION_NAMES** — Lista de nomes das soluções técnicas (ex: `["ms-fbso-platform-admin", "web-app-fbso-platform-portal"]`)
> 6. **ARCHITECTURE_GLOBAL** — Caminho da pasta de arquitetura global (ex: `/home/bolismar/work/workspace-fbso/architecture/`)
> 7. **SECURITY_GLOBAL** — Caminho do GLOBAL-SECURITY.md (ex: `/home/bolismar/work/workspace-fbso/.specs/security/GLOBAL-SECURITY.md`)
>
> Opcionais (pressione Enter para pular):
> 8. **PROJECT_DOCUMENTS_INPUTS** — Documentos brutos de entrada adicionais (deixe `[]` se não houver)
> 9. **PROJECT-STACK** — Stack tecnológica. Posso extrair do contexto do projeto, mas preciso que você valide contra a baseline corporativa em `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`
> 10. **PROJECT-TEAM-SKILLS-MAP** — Skills necessários para o time (deixe em branco para eu inferir do contexto)
> 11. **PROJECT-TEAM-CAPACITY** — Capacidade do time: seniores, plenos, juniores e duração prevista"

**NOTA SOBRE `PROJECT_PROMPT_INPUTS`:** Esta NÃO é uma variável de arquivo. É um checkpoint HITL (Human-in-the-Loop). No início de CADA fase e durante a execução, o orquestrador DEVE perguntar ao usuário:
> "Antes de prosseguir para a [Fase X — Documento Y]: deseja fornecer informações adicionais, novos contextos, direcionamentos ou ajustes de escopo?"

Esta porta NUNCA se fecha — o humano pode injetar novos inputs a qualquer momento.

#### Passo 0.2 — Validar Stack Contra Baseline Corporativa

Se `PROJECT-STACK` foi fornecida:
1. Ler `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`
2. Comparar cada tecnologia informada com a baseline
3. Para tecnologias dentro do padrão → ✅ Aprovado
4. Para tecnologias fora do padrão → ⚠️ Solicitar justificativa técnica ao usuário

> **🛠️ Validação de Stack:**
> - ☑️ Java 25 — ✅ Padrão corporativo
> - ☑️ Spring Boot 4.x — ✅ Padrão corporativo
> - ⚠️ MongoDB — ❌ Fora do padrão (padrão é PostgreSQL). Justificativa técnica?

Se `PROJECT-STACK` NÃO foi fornecida: inferir do contexto do projeto (briefing, documentos de entrada) e apresentar para validação do usuário.

#### Passo 0.3 — Solicitar e Validar PROJECT-TEAM-SKILLS-MAP

Se não fornecido, perguntar:
> "Com base no escopo do projeto e nas tecnologias identificadas, sugiro os seguintes skills para o time:
> - [lista inferida]
>
> Confirma? Deseja adicionar ou remover algum?"

#### Passo 0.4 — Solicitar e Validar PROJECT-TEAM-CAPACITY

Se não fornecido, perguntar:
> "Para dimensionar o cronograma e orçamento, preciso da capacidade do time:
> - Quantos seniores?
> - Quantos plenos?
> - Quantos juniores?
> - Duração prevista do projeto?"

#### Passo 0.5 — Exibir Caminhos Derivados e Solicitar Confirmação

> **📁 Caminho do Projeto:** `{PROJECT_COMPLETE_PATH_NAME}`
> **🏷️ Identificador:** `{PROJECT_ID_NAME}`
> **⚙️ Soluções Técnicas:** `{TECHNICAL_SOLUTION_NAMES}` ({N} solução(ões))
> **🏗️ Arquitetura Global:** `{ARCHITECTURE_GLOBAL}`
> **🛡️ Segurança Global:** `{SECURITY_GLOBAL}`
> **🛠️ Stack:** `{PROJECT-STACK}`
> **👥 Time:** `{PROJECT-TEAM-CAPACITY}`
>
> Confirma que estas informações estão corretas?
> - **SIM** → Prosseguir para criação/verificação da estrutura de diretórios
> - **NÃO** → Solicitar correção dos inputs e repetir

**Regra:** Não avance sem a confirmação explícita do humano.

#### Passo 0.6 — Criar Estrutura de Diretórios

```bash
mkdir -p {PROJECT_COMPLETE_PATH_NAME}
```

#### Passo 0.7 — Verificar Status dos Arquivos

Verificar existência e status de cada um dos 20 documentos. Para arquivos existentes, ler o cabeçalho e buscar por `[STATUS: COMPLIANCE]`.

| # | Arquivo | Status |
|---|---|---|
| 1 | `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` | ✅/❌ |
| 2 | `02-BRD-{PROJECT_ID_NAME}.md` | ✅/❌ |
| ... | ... (todos os 20) | ... |
| 20 | `20-LICOES-APRENDIDAS-{PROJECT_ID_NAME}.md` | ✅/❌ |

**Decisão:**
- Todos ❌ → Iniciar da Fase 1, Documento 1 (PROJECT-CHARTER)
- Parcial → Iniciar do primeiro documento sem `[STATUS: COMPLIANCE]`
- Todos ✅ → Perguntar: revisar, novo ciclo ou encerrar

#### Passo 0.8 — Apresentar Resumo e Iniciar

> **📊 Resumo do Projeto:** `{PROJECT_ID_NAME}`
> **📁 Localização:** `{PROJECT_COMPLETE_PATH_NAME}`
> **⚙️ Soluções Técnicas:** `{TECHNICAL_SOLUTION_NAMES}`
> **🛠️ Stack Validada:** `{PROJECT-STACK}`
> **👥 Time:** `{PROJECT-TEAM-CAPACITY}`
> **📝 Próxima Fase:** Fase X, Documento Y — {NOME}
> **📄 Artefatos Existentes:** X de 20 ({Y} com COMPLIANCE)
>
> Iniciando a Fase X...

---

## FASES DO ROADMAP WATERFALL

### FASE 1 — INICIALIZAÇÃO

| # | Documento | Arquivo |
|---|---|---|
| 1 | PROJECT-CHARTER | `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` |

### FASE 2 — REQUISITOS

| # | Documento | Arquivo |
|---|---|---|
| 2 | BRD | `02-BRD-{PROJECT_ID_NAME}.md` |
| 3 | SRS | `03-SRS-{PROJECT_ID_NAME}.md` |
| 4 | RTM | `04-RTM-{PROJECT_ID_NAME}.md` |

### FASE 3 — DESIGN E ARQUITETURA

| # | Documento | Arquivo |
|---|---|---|
| 5 | SAD | `05-SAD-{PROJECT_ID_NAME}.md` |
| 6 | HLD | `06-HLD-{PROJECT_ID_NAME}.md` |
| 7 | LLD | `07-LLD-{PROJECT_ID_NAME}.md` |

### FASE 4 — TESTES E QUALIDADE

| # | Documento | Arquivo |
|---|---|---|
| 8 | TEST-PLAN | `08-TEST-PLAN-{PROJECT_ID_NAME}.md` |
| 9 | TEST-CASES | `09-TEST-CASES-{PROJECT_ID_NAME}.md` |
| 15 | Relatório de Qualidade | `10-RELATORIO-QUALIDADE-{PROJECT_ID_NAME}.md` |

### FASE 5 — PLANEJAMENTO (pós-arquitetura e testes)

| # | Documento | Arquivo |
|---|---|---|
| 5 | EAP/WBS | `11-EAP-WBS-{PROJECT_ID_NAME}.md` |
| 6 | Cronograma/Gantt | `12-CRONOGRAMA-GANTT-{PROJECT_ID_NAME}.md` |
| 7 | Orçamento | `13-ORCAMENTO-{PROJECT_ID_NAME}.md` |
| 8 | Plano de Comunicação | `14-PLANO-COMUNICACAO-{PROJECT_ID_NAME}.md` |
| 9 | Plano de Riscos | `15-PLANO-RISCOS-{PROJECT_ID_NAME}.md` |

### FASE 6 — IMPLANTAÇÃO E ENCERRAMENTO

| # | Documento | Arquivo |
|---|---|---|
| 16 | DEPLOYMENT-PLAN | `16-DEPLOYMENT-PLAN-{PROJECT_ID_NAME}.md` |
| 17 | Manuais de Usuário | `17-MANUAIS-USUARIO-{PROJECT_ID_NAME}.md` |
| 18 | Manuais Operacionais | `18-MANUAIS-OPERACIONAIS-{PROJECT_ID_NAME}.md` |
| 19 | Termo de Aceite | `19-TERMO-ACEITE-{PROJECT_ID_NAME}.md` |
| 20 | Lições Aprendidas | `20-LICOES-APRENDIDAS-{PROJECT_ID_NAME}.md` |

---

## LOOP DE EXECUÇÃO POR DOCUMENTO

Para cada documento `{DOC}` na ordem sequencial acima:

### STEP 1: Computar inputs para GENERATE (por documento)

Cada GENERATE recebe APENAS as variáveis de domínio relevantes ao seu escopo, conforme a Matriz de Roteamento abaixo. NUNCA passar variáveis que o documento não utiliza.

```
DOC_PATH              = {PROJECT_COMPLETE_PATH_NAME}/{NN}-{DOC-SLUG}-{PROJECT_ID_NAME}.md
UPSTREAM_DOCS         = [lista de paths de documentos upstream já em COMPLIANCE]
EXTRA_INPUTS          = PROJECT_DOCUMENTS_INPUTS (documentos brutos de entrada)
SKILLS                = [lista de skills para este documento — vide tabela abaixo]

# Variáveis de domínio — passar APENAS as marcadas com ✅ para este documento:
{ARCHITECTURE_GLOBAL}  → somente se documento está na coluna ARCH
{SECURITY_GLOBAL}      → somente se documento está na coluna SEC
{TECHNICAL_SOLUTIONS}  → somente se documento está na coluna SOL
{TEAM_SKILLS}          → somente se documento está na coluna TEAM
{TEAM_CAPACITY}        → somente se documento está na coluna TEAM
{PROJECT-STACK}        → somente se documento está na coluna STACK
```

### STEP 2: Invocar GENERATE

Invocar `project-documents-waterfall/PROMPT-GENERATE-{DOC-SLUG}.md` passando **explicitamente** os parâmetros computados no STEP 1:
- `DOC_PATH`, `PROJECT_ID_NAME`, `UPSTREAM_DOCS`, `EXTRA_INPUTS`, `SKILLS`
- **+ variáveis de domínio** conforme a Matriz de Roteamento (apenas as marcadas com ✅ para este documento):
  - `ARCHITECTURE_GLOBAL` (coluna ARCH), `SECURITY_GLOBAL` (coluna SEC), `TECHNICAL_SOLUTIONS` (coluna SOL), `TEAM_SKILLS` + `TEAM_CAPACITY` (coluna TEAM), `PROJECT-STACK` (coluna STACK)

**Antes de cada GENERATE:** Checkpoint HITL — perguntar ao usuário se deseja fornecer novos inputs (diretiva `PROJECT_PROMPT_INPUTS`):
> "Antes de gerar [Documento X]: deseja fornecer informações adicionais, novos contextos ou ajustes de escopo?"

### STEP 3: Invocar GATE

Invocar `project-documents-waterfall/PROMPT-GATE-{DOC-SLUG}.md` passando **explicitamente**:
- `DOC_PATH` (recebido do GENERATE)
- `CHECKLIST` (definido no próprio prompt GATE)

GATE deve:
- Ler apenas o arquivo em `DOC_PATH`
- Atualizar status para `[STATUS: Em revisão]`
- Retornar `{PASS}` ou `{FAIL, VIOLATIONS: [{section, description, severity}]}`

### STEP 4a: Se GATE retornar FAIL → Invocar FIX

Invocar `project-documents-waterfall/PROMPT-FIX-{DOC-SLUG}.md` passando **explicitamente**:
- `DOC_PATH` (mesmo arquivo)
- `VIOLATIONS[]` (lista exata de não-conformidades do GATE)

FIX deve:
- Editar **apenas** as seções em `VIOLATIONS[]`
- Manter status como `[STATUS: Em revisão]`
- Retornar `{DOC_PATH}`
- → Voltar ao STEP 3 (re-executar GATE)

### STEP 4b: Se GATE retornar PASS → Validação Humana + Checkpoint HITL

Apresentar o documento e fazer 4 perguntas:

> **P1:** "O conteúdo deste documento está aderente às necessidades do projeto?"
> **P2:** "Existem novos documentos de entrada que devem ser incorporados a esta fase?"
> **P3:** "Há novas informações textuais, mudanças de escopo ou ajustes técnicos a serem considerados?"
> **P4 (HITL):** "Deseja fornecer novas informações, contextos adicionais ou direcionamentos antes de prosseguir para o próximo documento?"

- Se humano aprovar (SIM para P1, NÃO para P2/P3/P4): `[STATUS: COMPLIANCE]`, documento congelado, próximo documento liberado
- Se humano fornecer novos inputs (P2, P3 ou P4): Voltar ao STEP 2 (re-executar GENERATE com novo contexto)
- **A porta HITL NUNCA se fecha:** mesmo após aprovação, o orquestrador deve aceitar novos inputs e reabrir qualquer fase

---

## MATRIZ DE SKILLS POR DOCUMENTO

| # | Documento | Skills Primárias | Fallback |
|---|---|---|---|
| 1 | PROJECT-CHARTER | `draft-project-charter`, `senior-pm` | ✅ |
| 2 | BRD | `brd-creation`, `business-analyst`, `requirements-elicitation` | ✅ |
| 3 | SRS | `frs-creation`, `requirements-engineering` | ✅ |
| 4 | RTM | `requirements-modeling`, `requirements-validation` | ✅ |
| 5 | EAP/WBS | `decomposition-planning-roadmap`, `project-estimation` | ✅ |
| 6 | Cronograma/Gantt | `roadmap-planning`, `project-estimation` | ✅ |
| 7 | Orçamento | `project-estimation` | ✅ |
| 8 | Plano de Comunicação | `stakeholder-analysis`, `stakeholder-map` | ✅ |
| 9 | Plano de Riscos | `risk-manager`, `risk-management-specialist` | ✅ |
| 5 | SAD | `software-architecture`, `architecture-designer`, `architecture-patterns` | ✅ |
| 6 | HLD | `c4-container`, `system-design`, `architecture-decision-records` | ✅ |
| 7 | LLD | `c4-component`, `ddd-tactical-patterns`, `database-designer` | ✅ |
| 8 | TEST-PLAN | `test-strategy-design`, `qa-test-planner` | ✅ |
| 9 | TEST-CASES | `test-case-creation`, `acceptance-criteria` | ✅ |
| 15 | Relatório de Qualidade | `quality-documentation-manager`, `qa` | ✅ |
| 16 | DEPLOYMENT-PLAN | `deployment-engineer`, `devops-rollout-plan` | ✅ |
| 17 | Manuais de Usuário | `documentation-generation-doc-generate`, `docs-writer` | ✅ |
| 18 | Manuais Operacionais | `documentation-generation-doc-generate` | ✅ |
| 19 | Termo de Aceite | `contract-and-proposal-writer` | ✅ |
| 20 | Lições Aprendidas | — | ✅ |

---

## MATRIZ DE ROTEAMENTO DE DOMÍNIO (quais variáveis cada documento recebe)

Cada linha mostra quais variáveis de domínio o GENERATE daquele documento deve receber. Variáveis não marcadas NÃO devem ser passadas.

| # | Documento | ARCH | SEC | SOL | TEAM | STACK | Fundamentação |
|---|---|---|---|---|---|---|---|
| 1 | PROJECT-CHARTER | — | — | — | ✅ | ✅ | Documenta time e stack no plano de negócio (Seção 11 Orçamento + Seção 5 RACI) |
| 2 | BRD | — | — | — | — | — | Requisitos de negócio puros — sem detalhes técnicos |
| 3 | SRS | — | — | — | — | — | Especificação funcional — sem acoplamento com stack |
| 4 | RTM | — | — | — | — | — | Matriz de rastreabilidade — puramente relacional |
| 5 | EAP/WBS | — | — | — | ✅ | — | Decomposição de trabalho por perfil de time |
| 6 | Cronograma/Gantt | — | — | — | ✅ | — | Alocação de recursos e durações por capacidade |
| 7 | Orçamento | — | — | — | ✅ | ✅ | Custo por recurso (RH × stack) |
| 8 | Plano de Comunicação | — | — | — | — | — | Stakeholders e canais — sem dependência técnica |
| 9 | Plano de Riscos | — | — | — | — | — | Riscos de projeto — sem dependência técnica |
| 5 | SAD | ✅ | ✅ | ✅ | — | ✅ | 6 visões arquiteturais — requer ARCH global, SEC global, stack e soluções |
| 6 | HLD | ✅ | ✅ | ✅ | — | ✅ | Design de alto nível — requer ARCH (ADRs), SEC (decisões) e stack (tecnologias) |
| 7 | LLD | ✅ | — | ✅ | — | ✅ | Design de baixo nível — requer ARCH (padrões), stack (frameworks) e soluções |
| 8 | TEST-PLAN | — | ✅ | ✅ | — | ✅ | Estratégia de testes — requer stack (ferramentas), SEC (testes segurança) e soluções |
| 9 | TEST-CASES | — | — | ✅ | — | — | Casos de teste por feature/solução |
| 15 | Relatório de Qualidade | — | — | — | — | — | Métricas de qualidade — sem dependência técnica direta |
| 16 | DEPLOYMENT-PLAN | ✅ | ✅ | ✅ | — | ✅ | Deploy — requer ARCH (topologia), SEC (secure deploy) e stack |
| 17 | Manuais de Usuário | — | — | ✅ | — | — | Documentação por solução |
| 18 | Manuais Operacionais | ✅ | — | ✅ | — | ✅ | Runbooks — requer ARCH (visão ops), stack e soluções |
| 19 | Termo de Aceite | — | — | — | — | — | Aceite formal — sem dependência técnica |
| 20 | Lições Aprendidas | — | — | — | — | — | Retrospectiva — sem dependência técnica |

**Legenda das colunas:**
- **ARCH** = `ARCHITECTURE_GLOBAL` — ADRs, blueprints, padrões globais de arquitetura
- **SEC** = `SECURITY_GLOBAL` — GLOBAL-SECURITY.md (regras de ouro, checklist SDD)
- **SOL** = `TECHNICAL_SOLUTIONS` + `TECHNICAL_SOLUTION_NAMES` — lista de soluções técnicas do projeto
- **TEAM** = `PROJECT-TEAM-SKILLS-MAP` + `PROJECT-TEAM-CAPACITY` — skills e capacidade do time
- **STACK** = `PROJECT-STACK` — stack tecnológica validada

**Regra:** O orquestrador DEVE consultar esta matriz antes de cada STEP 2 e passar APENAS as variáveis marcadas com ✅ para o documento corrente. Variáveis não marcadas NÃO devem ser injetadas no GENERATE.

---

## MAPEAMENTO DE UPSTREAM DOCS DE CADA DOCUMENTO

Cada GENERATE recebe em `UPSTREAM_DOCS` a lista de documentos anteriores que já estão em COMPLIANCE e que este documento deve referenciar:

| # | Documento | UPSTREAM_DOCS (paths relativos) |
|---|---|---|
| 1 | PROJECT-CHARTER | `[]` (documento raiz) |
| 2 | BRD | `[01-PROJECT-CHARTER]` |
| 3 | SRS | `[01-PROJECT-CHARTER, 02-BRD]` |
| 4 | RTM | `[01-PROJECT-CHARTER, 02-BRD, 03-SRS]` |
| 5 | EAP/WBS | `[01-PROJECT-CHARTER, 07-LLD, 09-TEST-CASES]` |
| 6 | Cronograma/Gantt | `[01-PROJECT-CHARTER, 11-EAP-WBS, 09-TEST-CASES]` |
| 7 | Orçamento | `[01-PROJECT-CHARTER, 11-EAP-WBS, 06-Cronograma, 08-TEST-PLAN]` |
| 8 | Plano de Comunicação | `[01-PROJECT-CHARTER]` |
| 9 | Plano de Riscos | `[01-PROJECT-CHARTER]` |
| 5 | SAD | `[01-PROJECT-CHARTER, 02-BRD, 03-SRS]` |
| 6 | HLD | `[01-PROJECT-CHARTER, 05-SAD]` |
| 7 | LLD | `[01-PROJECT-CHARTER, 05-SAD, 06-HLD]` |
| 8 | TEST-PLAN | `[01-PROJECT-CHARTER, 03-SRS, 05-SAD, 07-LLD]` |
| 9 | TEST-CASES | `[01-PROJECT-CHARTER, 03-SRS, 08-TEST-PLAN]` |
| 15 | Relatório de Qualidade | `[08-TEST-PLAN, 09-TEST-CASES]` |
| 16 | DEPLOYMENT-PLAN | `[01-PROJECT-CHARTER, 05-SAD, 06-HLD, 07-LLD]` |
| 17 | Manuais de Usuário | `[01-PROJECT-CHARTER, 03-SRS]` |
| 18 | Manuais Operacionais | `[01-PROJECT-CHARTER, 05-SAD, 16-DEPLOYMENT-PLAN]` |
| 19 | Termo de Aceite | `[01-PROJECT-CHARTER, 08-TEST-PLAN, 10-RELATORIO-QUALIDADE]` |
| 20 | Lições Aprendidas | `[todos os 19 documentos anteriores]` |

---

## EFEITOS CASCATA

Quando um documento já em COMPLIANCE é modificado, todos os documentos downstream que dependem dele devem ser regenerados e revalidados. A ordem de dependência agora segue: **Requisitos → Arquitetura → Testes → Planejamento → Deploy**.

| Se modificar... | Impacta (regenerar + revalidar)... |
|---|---|
| 01-PROJECT-CHARTER | Todos os 19 documentos downstream |
| 02-BRD | 03-SRS, 04-RTM, 05-SAD, 06-HLD, 07-LLD, 08-TEST-PLAN, 09-TEST-CASES, 05-EAP, 06-Cronograma, 07-Orçamento, 16-DEPLOYMENT-PLAN, 17-Manuais, 19-Termo |
| 05-SAD | 06-HLD, 07-LLD, 08-TEST-PLAN, 16-DEPLOYMENT-PLAN, 18-Manuais Ops |
| 07-LLD | 08-TEST-PLAN, 09-TEST-CASES, **05-EAP**, **06-Cronograma**, **07-Orçamento**, 16-DEPLOYMENT-PLAN |
| 08-TEST-PLAN | 09-TEST-CASES, 15-Relatório, **05-EAP**, **06-Cronograma**, **07-Orçamento**, 19-Termo Aceite |
| 09-TEST-CASES | 15-Relatório, **05-EAP**, **06-Cronograma** |
| 05-EAP/WBS | 06-Cronograma, 07-Orçamento |
| (regra geral) | Todos os documentos listados como UPSTREAM_DOCS do modificado e seus dependentes transitivos |

**Destaque:** A alteração mais impactante é no 07-LLD, que agora cascateia para Planejamento (EAP, Cronograma, Orçamento) via Testes. Modificações no design de baixo nível ou nos planos de teste **exigem reestimativa de esforço, prazo e custo**.

**Ação:** Alertar o humano, listar documentos afetados, perguntar: (A) regeneração completa, ou (B) marcar como "potencialmente desatualizados."

---

## FINALIZAÇÃO — GIT WORKFLOW

Quando os 20 documentos estiverem em COMPLIANCE e o humano confirmar a conclusão, executar o pipeline Git abaixo. O nome da branch de trabalho é derivado automaticamente:

```
WORK_BRANCH = "feature/" + PROJECT_ID_NAME + "-waterfall-docs"
```

Exemplo: `feature/PRJ-FIN-2026-0003-SAAS-FBSO-ORG-waterfall-docs`

**Regra de segurança:** Validar que `WORK_BRANCH` não é `main`, `master` ou `develop`. Se houver colisão, adicionar sufixo numérico.

### Passo F.1 — Git Add e Commit

```bash
git add -A
git commit -m "docs: documentação WATERFALL completa — ${PROJECT_ID_NAME}

- 20 documentos WATERFALL gerados e validados
- Status: COMPLIANCE em todos os documentos
- Gerado pelo Waterfall Orchestrator v1.0
- Branch: ${WORK_BRANCH}

Co-Authored-By: Claude <noreply@anthropic.com>"
```

**Regra:** Se não houver alterações para commitar, informar e encerrar.

### Passo F.2 — Git Push

```bash
git push origin ${WORK_BRANCH}
```

Se falhar (branch remota existe): perguntar sobre `--force`.

### Passo F.3 — Criar e Mergear PR

```bash
gh pr create --base main --head ${WORK_BRANCH} --title "docs: documentação WATERFALL — ${PROJECT_ID_NAME}" --body "Documentação WATERFALL completa para ${PROJECT_ID_NAME}. 20 documentos validados."
gh pr merge --merge --delete-branch
```

### Passo F.4 — Cleanup Local

```bash
git checkout main
git branch -d ${WORK_BRANCH}
```

**Regra de segurança:** Nunca deletar branch local se o merge falhar.

---

## LOCALIZAÇÃO DOS PROMPTS

```
.specs/prompts/project-documents-waterfall/
├── PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS-WATERFALL.md   ← ESTE ORQUESTRADOR
├── PROMPT-GENERATE-PROJECT-CHARTER.md
├── PROMPT-GATE-PROJECT-CHARTER.md
├── PROMPT-FIX-PROJECT-CHARTER.md
├── PROMPT-GENERATE-BRD.md
├── PROMPT-GATE-BRD.md
├── PROMPT-FIX-BRD.md
├── (total de 61 prompts)
└── PROMPT-FIX-LICOES-APRENDIDAS.md
```

---

🤖 *Roadmap gerado pelo Waterfall Orchestrator v1.0. Skills: draft-project-charter, senior-pm, brainstorming.*
