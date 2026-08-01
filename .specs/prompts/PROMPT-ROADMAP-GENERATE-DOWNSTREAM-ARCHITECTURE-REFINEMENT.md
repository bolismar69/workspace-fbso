# PROMPT: ROADMAP DE DOWNSTREAM ARCHITECTURE REFINEMENT
## Versão: 2.0 — Análise de Viabilidade + Estimativa Bottom-Up PERT Independente + Scope Snapshot

Atue como um Especialista em Gestão de Processos (BPM), Arquiteto de Soluções Ágeis e Tech Lead, especializado em Refinamento Técnico, Design Detalhado e Estimativas de Precisão.

Preciso que você crie um roadmap de execução detalhado e um guia de orquestração para o seguinte processo: **Análise de viabilidade técnica e estimativa detalhada do projeto pelo time de TI** na fase de Downstream Engineering — quando o Negócio já definiu Project Charter, BRD, Épicos, Features e User Stories e o time de TI precisa realizar sua própria análise de viabilidade, produzir design detalhado das disciplinas técnicas e gerar uma estimativa independente e precisa (PERT ±15-25%) para compromisso de prazo e orçamento.

**Este roadmap é focado em estimativa e análise de viabilidade — não em planejamento de sprints ou entregas.** O refinamento técnico do que será feito em cada sprint e os contratos técnicos (API, Data, Security, SRE) são responsabilidade de outro roadmap (`PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md`). Este roadmap entrega: design detalhado das 6 disciplinas + estimativa PERT independente + snapshot do escopo estimado + cross-check com upstream (se existir).

**Este roadmap é independente.** Ele pode ser executado:
- **Após o Upstream Architecture Discovery** — como refinamento natural do ROM ±50% para PERT ±15-25%
- **Sem o Upstream Architecture Discovery** — quando a área de negócios quer uma estimativa detalhada direta, sem ter passado pela fase de discovery

**A estimativa bottom-up PERT é completamente independente de estimativas anteriores.** O cálculo é feito do zero, US por US, sem usar ou depender de nenhuma estimativa pré-existente. O único vínculo com o upstream discovery acontece na **Fase 12 (Cross-Check Report)**, ao final do roadmap, onde o orquestrador verifica se existe um ROM upstream e gera um relatório comparativo.

Objetivo Principal: Produzir **Análise de Viabilidade Técnica + Design Detalhado das 6 Disciplinas + Estimativa Bottom-Up PERT independente + Scope Snapshot + Cross-Check com Upstream** com qualidade suficiente para o time de TI validar a viabilidade do projeto e o comitê de governança aprovar prazo e orçamento com confiança de ±15-25%.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial em todas as fases. Nenhuma fase subsequente pode ser iniciada sem a aprovação formal, soberana e explícita do usuário humano na fase anterior.

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP (FASE 0)

### Tabela de Inputs

| Variável | Obrig. | Descrição | Exemplo |
|---|---|---|---|
| `PROJECT_PATH` | ✅ | Caminho base onde os projetos de negócio residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID_NAME` | ✅ | Identificador completo do projeto (ID + Nome) | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `TECHNICAL_SOLUTION_PATH` | ✅ | Caminho base onde as soluções técnicas residem | `/home/bolismar/work/workspace-fbso/backend/java/spring/microservices` |
| `TECHNICAL_SOLUTION_NAMES` | ✅ | Lista de nomes das soluções técnicas do projeto | `["ms-fbso-platform-admin", "web-app-fbso-platform-portal"]` |
| `ARCHITECTURE_GLOBAL` | ✅ | Caminho para a pasta de arquitetura global (ADRs, blueprints, padrões) | `/home/bolismar/work/workspace-fbso/architecture/` |
| `SECURITY_GLOBAL` | ✅ | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) | `/home/bolismar/work/workspace-fbso/.specs/security/GLOBAL-SECURITY.md` |
| `PROJECT_DOCUMENTS_INPUTS` | ❌ | Lista de caminhos para documentos brutos de entrada adicionais | `[]` |
| `PROJECT_PROMPT_INPUTS` | ❌ | Lista de caminhos para prompts auxiliares ou contextos adicionais | `[]` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME       = PROJECT_PATH + "/" + PROJECT_ID_NAME
DOWNSTREAM_REFINEMENT_PATH       = PROJECT_COMPLETE_PATH_NAME + "/downstream-architecture-refinement"
UPSTREAM_DISCOVERY_PATH          = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
TECHNICAL_DEFINITIONS_PATH       = PROJECT_COMPLETE_PATH_NAME + "/technical-definitions"
```

### Pré-condição Obrigatória

Apenas **um** roadmap deve ter sido executado antes deste:

1. **`PROMPT-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md`** — executado até a **Fase 5 (USER-STORIES)**. Os seguintes artefatos devem existir:
   - `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md`
   - `02-BRD-{PROJECT_ID_NAME}.md`
   - `03-EPICS-{PROJECT_ID_NAME}.md`
   - `04-FEATURES-{PROJECT_ID_NAME}.md`
   - `05-USER-STORIES-{PROJECT_ID_NAME}.md`
   - `epics/*.md` (arquivos individuais de épicos)
   - `features/*.md` (arquivos individuais de features)
   - `user-stories/*.md` (arquivos individuais de user stories)

### Pré-condição Opcional (Zero Impacto na Estimativa)

O roadmap **`PROMPT-ROADMAP-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY.md`** é **opcional**. Se tiver sido executado, os artefatos Discovery-Level (em `upstream-architecture-discovery/`) serão usados apenas como **referência complementar** para o design detalhado — nunca como entrada para a estimativa. A estimativa bottom-up PERT é sempre calculada do zero, independentemente da existência ou não do ROM upstream.

**Regra de Independência da Estimativa:** O Bloco B (Fases 8-10) não pode, sob nenhuma circunstância, usar valores do ROM upstream como entrada, baseline, ponto de partida ou referência para os cálculos PERT. Cada US é estimada individualmente com three-point estimation pura. O ROM upstream só é consultado na Fase 12 (Cross-Check Report), após a estimativa já estar concluída e congelada.

---

## ARQUITETURA DE FASES

O roadmap é organizado em **12 fases** agrupadas em **4 blocos**:

```
FASE 0: BOOTSTRAP (sequencial)
  │
  ├─▶ BLOCO A: Architecture Deep-Dive (Detail-Level)
  │     Fase 1 → Fase 2 → Fase 3 → Fase 4 → Fase 5 → Fase 6 → Fase 7
  │     ⛔ Barreira A
  │
  ├─▶ BLOCO B: Bottom-Up Estimation — INDEPENDENTE (PERT ±15-25%)
  │     Fase 8 → Fase 9 → Fase 10
  │     ⛔ Barreira B
  │
  ├─▶ BLOCO C: Scope Snapshot + Cross-Check
  │     Fase 11 → Fase 12 (condicional: só executa se upstream existir)
  │     ⛔ Barreira C
  │
  └─▶ GATE DE ESTIMATIVA: ESTIMATE-READY 📊
        ├── Estimativa aprovada → Time de TI valida viabilidade
        └── Pendências → Retorna ao bloco com gaps
```

---

## FASES DO ROADMAP

### Fase 0 — Bootstrap Inteligente

Workflow:
1. Solicitar inputs ao usuário (se não fornecidos)
2. **Validar pré-condição de negócio:** Verificar existência de Charter, BRD, Épicos, Features, User Stories (docs de negócio)
3. **Detectar upstream (informativo):** Verificar se o diretório `upstream-architecture-discovery/` existe. Se sim, listar os artefatos encontrados e informar que serão usados como referência complementar para o design (Bloco A), mas NÃO para a estimativa (Bloco B). Se não, informar que o refinamento prossegue sem referência upstream — todos os artefatos Detail-Level serão gerados do zero.
4. **Invocar `project-document-discovery`:** Classificar o projeto em 10 dimensões (4-signal algorithm) para determinar o escopo proporcional de documentos no refinement
5. **Invocar `discovery-process`:** Estruturar o ciclo de refinamento (framing → synthesis → experiments) adaptado ao escopo detalhado
6. **Auditar User Stories:** Contar US por épico/feature, verificar distribuição, identificar gaps de cobertura
7. Exibir caminhos derivados e solicitar confirmação
8. Criar estrutura: `mkdir -p {DOWNSTREAM_REFINEMENT_PATH}`
9. Auditar artefatos existentes no diretório `downstream-architecture-refinement/`
10. Apresentar resumo da situação atual e iniciar a primeira fase pendente

### Fase 1 — DETAIL-LEVEL-PRD.md 🆕
PRD Detail-Level — **alinhamento negócio↔TI com escopo completo**. Visão do produto consolidada com as 62 US, personas detalhadas, jornadas mapeadas por feature, restrições de negócio, glossário estendido. Referencia os documentos de negócio e o PRD Discovery-Level. **Mais detalhado que o Discovery-Level PRD:** inclui mapeamento US↔Jornada, priorização por entrega, e critérios de aceite de negócio sumarizados.

Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-PRD.md` → Gate → Fix → COMPLIANCE

### Fase 2 — DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md 🆕
Solution Architect — **C4 Level 2 (Container) e Level 3 (Component)** para as soluções S01 e S02. ADRs detalhados com diagramas de sequência, matriz de integração refinada (com contratos de API), estratégia de multi-tenancy (RLS + discriminator column), padrões de código (packages, naming, design patterns). Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 3 — DETAIL-LEVEL-SECURITY-DEFINITION.md 🆕
Security Architect — **threat model detalhado** (STRIDE por componente), matriz de controles (OWASP ASVS L1+L2), especificação de IAM (Keycloak realms, OIDC clients, protocol mappers, JWT claims), política de senhas, RBAC granular (role↔permission matrix), data protection (AES-256, RLS policies), compliance (LGPD artigo por artigo). Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SECURITY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 4 — DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md 🆕
Data Architect — **ERD completo** (todas as tabelas com colunas, tipos, constraints, índices), estratégia de particionamento (audit_log por mês), política de retenção (5 anos PostgreSQL + S3), migration strategy (Flyway versionado com baseline), query patterns otimizadas (índices para dashboard, filtros de tenant). Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DATA-ARCHITECTURE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 5 — DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md 🆕
DevOps/SRE Architect — **pipeline specs detalhadas** (GitHub Actions workflow por ambiente), IaC templates (Terraform DOKS + Ansible playbooks), observabilidade stack (Prometheus alert rules, Grafana dashboards, Loki log queries, Jaeger sampling strategy, OpenTelemetry instrumentation), SLOs com SLIs (99.9% backend p99<500ms, 99.5% frontend LCP<2.5s), estratégia de deploy (blue-green, canary). Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DEVOPS-SRE-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 6 — DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md 🆕
Test Specialist — **matriz de cobertura por US** (62 US × tipo de teste), casos de teste de aceitação (baseados nos cenários Gherkin das US), estratégia de automação (JUnit 5 + Mockito + Testcontainers para backend, Jest + Testing Library para frontend, Playwright para E2E, k6 para carga), quality gates por ambiente (PR/Staging/Release). Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-TEST-STRATEGY-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 7 — DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md 🆕
Infra/Cloud Specialist — **sizing detalhado** (DOKS node pools, PostgreSQL tiers, Redis tiers), cálculo de custos mensais (DigitalOcean + Cloudflare + Kong), topologia de rede (VPC, subnets, security groups), disaster recovery (RPO 1h, RTO 4h, procedimentos), estratégia de backup (PostgreSQL WAL + pg_dump, retenção 30 dias). Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-INFRA-CLOUD-DEFINITION.md` → Gate → Fix → COMPLIANCE

### Fase 8 — BOTTOM-UP-PERT-ESTIMATE.md 🆕 ⭐
**O coração do downstream — estimativa bottom-up PERT.** Cada uma das 62 User Stories é estimada individualmente com three-point estimation (O/ML/P). Skills específicos de estimativa são usados (`project-estimation`, referências `bottom-up-estimation.md` e `three-point-estimation-pert.md`). O resultado substitui o ROM ±50% do upstream por uma estimativa PERT ±15-25%.

Conteúdo:
- Estimativa individual por US (complexidade, O, ML, P, PERT, σ, IC 95%)
- Rollup Feature → Épico → Projeto
- Composição do esforço: Dev + QA (≥25%) + Arch (≥5%) + DevOps + Gestão + Contingência
- Comparação com ROM upstream e factory bids
- Validação DTA (QA balanceado, consistência prazo×horas, outliers)

Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-BOTTOM-UP-PERT-ESTIMATE.md` → Gate → Fix → COMPLIANCE

### Fase 9 — RESOURCE-ALLOCATION-PLAN.md 🆕
**Plano de alocação de recursos** baseado na estimativa PERT. Capacidade do time (9 pessoas, cargas parciais), projeção de duração (PERT ÷ capacidade mensal), perfil alocado por milestone, identificação de gargalos (frontend M5, RBAC M4), recomendações de reforço.

Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RESOURCE-ALLOCATION.md` → Gate → Fix → COMPLIANCE

### Fase 10 — RISK-ADJUSTED-ESTIMATE.md 🆕
**Estimativa ajustada a risco.** Aplica a matriz de riscos do projeto (6 riscos identificados) sobre a estimativa PERT, produzindo 3 cenários (Conservador 15% / PERT / Pessimista 25%), com confidence intervals e análise de sensibilidade (quais riscos mais impactam o prazo).

Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RISK-ADJUSTED-ESTIMATE.md` → Gate → Fix → COMPLIANCE

### Fase 11 — SCOPE-SNAPSHOT.md 🆕 📸
**Foto do escopo estimado — registro do que foi incluído na estimativa.** Este documento é um snapshot, não um plano de sprints. Sua função é registrar de forma imutável qual era o escopo no momento da estimativa, para que futuras revisões possam identificar o que mudou.

⚠️ **Importante:** Este documento **NÃO dispara planejamento de sprints ou entregas.** O refinamento técnico das sprints e a geração de contratos (API, Data, Security, SRE) são responsabilidade do roadmap `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md`.

Conteúdo:
- Lista completa das US incluídas na estimativa (ID, descrição, feature, épico, entrega)
- Contagem por entrega (D1-D7), épico e feature
- US agrupadas por complexidade (Simples/Média/Complexa)
- Data de congelamento do escopo
- Hash/checksum do conjunto de US estimadas (para detectar mudanças futuras)
- Nota explícita: "Este documento é um snapshot do escopo em [data]. O planejamento de sprints será feito no roadmap PROJECT-TECHNICAL-DEFINITIONS."

Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SCOPE-SNAPSHOT.md` → Gate → Fix → COMPLIANCE

### Fase 12 — UPSTREAM-COMPARISON-REPORT.md 🔗 (condicional)
**Cross-Check com Upstream — o único vínculo entre os dois roadmaps.** Esta fase é **condicional**: só executa se o diretório `upstream-architecture-discovery/` existir e contiver o arquivo `DISCOVERY-LEVEL-ROM-ESTIMATE.md`.

⚠️ **Regra crítica:** Esta fase só inicia **após a estimativa PERT estar concluída, congelada e aprovada** na Barreira B. O ROM upstream NUNCA é consultado durante o cálculo da estimativa. O relatório é puramente comparativo e informativo.

Conteúdo do relatório:
- Tabela comparativa: ROM upstream vs PERT downstream (horas totais, por épico)
- Análise de convergência/divergência (desvio percentual)
- Se PERT dentro da faixa ROM (±50%): convergência esperada
- Se PERT fora da faixa ROM: análise de causas (escopo adicional? complexidade subestimada no ROM? novas US?)
- Gráfico visual de comparação (barras proporcionais)
- Conclusão: a estimativa PERT refina, substitui ou diverge do ROM?
- Se upstream não existir: documento registra "Upstream discovery não encontrado — sem ROM para comparação"

**Importante:** Este relatório NÃO altera a estimativa PERT. Seu propósito é fornecer ao comitê uma visão da evolução da precisão (ROM ±50% → PERT ±15-25%) e destacar quaisquer divergências significativas que mereçam atenção.

Pipeline: `PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-UPSTREAM-COMPARISON.md` → Gate → Fix → COMPLIANCE

### Gate de Estimativa — ESTIMATE-READY 📊

Após a Barreira C, o orquestrador apresenta o **Resumo Executivo da Análise de Viabilidade e Estimativa**:

1. **Escopo Estimado:** Snapshot das US incluídas (Fase 11)
2. **Design Detalhado:** 6 disciplinas validadas pelo time de TI (Bloco A)
3. **Estimativa Bottom-Up PERT:** Faixa com ±15-25% de confiança, resource allocation, riscos (Bloco B)
4. **Cross-Check com Upstream:** Relatório comparativo (se aplicável) (Fase 12)
5. **Parecer de Viabilidade:** Recomendação do time de TI sobre a viabilidade do projeto

O Comitê decide:
- **Estimate Accepted ✅:** Estimativa aprovada como baseline de prazo e orçamento. Dispara:
  - `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md` — para refinamento técnico e planejamento de sprints
  - Revisão da estimativa após 2 sprints de execução (refinar para ±10%)
- **Pendências ⚠️:** Retorna ao bloco com gaps para ajuste

---

## MECANISMO DE ORQUESTRAÇÃO DINÂMICA

Toda fase (1-12) deve rodar sob o ecossistema trifásico de prompts (Gerador, Auditor/Portão e Corretor), com controle final obrigatório do Humano:

1. **Geração / Evolução:** A IA recebe os inputs disponíveis e executa o prompt gerador da fase
2. **Auditoria Interna da IA:** O artefato é enviado para o gate.
   - SE A IA ENCONTRAR ERROS: Emite o status `[NÃO COMPLIANCE]`, apresenta as falhas encontradas com sugestões de tratativa, faz as 3 perguntas obrigatórias do Portão de Validação Humana e aguarda o direcionamento do humano. Com base nas respostas, aciona o FIX de forma cirúrgica e retorna ao passo 2.
   - SE A IA NÃO ENCONTRAR ERROS (100% OK): Emite o status `[PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]`, apresenta as 3 perguntas obrigatórias e aguarda o direcionamento humano.
3. **Portão de Validação Humana (3 perguntas obrigatórias):**
   1. O documento está em compliance com a sua necessidade e perfeitamente alinhado com os documentos base (negócio + upstream discovery)?
   2. Deseja enviar mais documentos/arquivos para enriquecer este artefato?
   3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?
4. **Lógica de Decisão Baseada nas Respostas do Humano:**
   - CENÁRIO DE SUCESSO (Aprovação): Se o humano validar e NÃO enviar novos arquivos ou inputs (Sim, Não, Não), a fase é dada por encerrada (`[STATUS: COMPLIANCE]`), o arquivo é congelado e a próxima fase é destravada.
   - CENÁRIO DE RETROCESSO (Evolução Incremental): Se o humano fornecer novos documentos ou novas informações, o orquestrador DEVE retroceder ao passo 1 (GENERATE), injetando o documento gerado até o momento + os novos insumos para uma atualização incremental.

---

## REGRAS DE BLOQUEIO (GATING RULES)

### Barreiras de Bloco

| Barreira | Posição | Validação | Regra Especial |
|---|---|---|---|
| ⛔ Barreira A | Após Bloco A (F7) | 7 disciplinas OK. Consistência horizontal entre os 7 artefatos. **Detalhamento suficiente para o time de TI validar viabilidade** — cada artefato deve conter análise técnica que permita identificar riscos e complexidades. Se upstream existir, artefatos podem referenciá-lo como baseline, mas devem ir além. | Artefato muito similar ao Discovery-Level (sem refinamento real) = NÃO COMPLIANCE |
| ⛔ Barreira B | Após Bloco B (F10) | **PERT com todas as US estimadas individualmente.** QA ≥ 25%. Arch ≥ 5%. Consistência Prazo×Horas validada. Outliers identificados. **Estimativa 100% independente — não pode ter usado ROM upstream como baseline.** Confiança alvo: ±15-25%. | US sem estimativa individual = NÃO COMPLIANCE. PERT sem IC 95% = NÃO COMPLIANCE. Evidência de contaminação pelo ROM = NÃO COMPLIANCE |
| ⛔ Barreira C | Após Bloco C (F12) | Scope Snapshot cobre 100% das US estimadas. Se upstream existe: relatório comparativo gerado. Se upstream não existe: F12 pulada, barreira satisfeita. Relatório NÃO altera estimativa PERT. | Scope Snapshot incompleto = NÃO COMPLIANCE |

### Consistência Horizontal (Bloco A)

A Barreira A deve validar que os 7 artefatos Detail-Level são consistentes entre si:

- ARCHITECTURE ↔ SECURITY: controles de segurança implementam padrões arquiteturais definidos
- ARCHITECTURE ↔ DATA: modelo de dados suporta a arquitetura de componentes
- ARCHITECTURE ↔ DEVOPS-SRE: pipeline de deploy viabiliza a topologia
- ARCHITECTURE ↔ INFRA-CLOUD: sizing de infra suporta a arquitetura
- SECURITY ↔ INFRA-CLOUD: controles de rede e IAM consistentes
- TEST-STRATEGY ↔ ARCHITECTURE: matriz de testes cobre os componentes

### Validação DTA da Estimativa (Bloco B)

A Barreira B aplica as regras DTA de validação de estimativas:

| Regra | Critério | Ação se não atender |
|:---|:---|:---|
| **QA Balanceado** | QA ≥ 20% por épico | ⚠️ Risco de Débito Técnico |
| **QA Global** | QA ≥ 25% do total de horas | ⚠️ Risco de Subinvestimento em Qualidade |
| **Arquitetura/SRE** | Arch ≥ 5% do total geral de horas | ⚠️ Risco de Subinvestimento Técnico |
| **PERT por US** | Todas as US com O, ML, P, PERT, σ | ❌ NÃO COMPLIANCE |
| **Consistência Prazo×Horas** | `prazo_calculado = total_horas / (time_estimado × 160h)`. Divergência >50% → ⚠️ | 🔍 Revisão manual |
| **Independência da Estimativa** | Bloco B não pode ter usado valores do ROM upstream como baseline ou ponto de partida. A estimativa é pura e independente. | ❌ NÃO COMPLIANCE — refazer Bloco B |
| **Cross-Check (Fase 12)** | Se upstream existe: relatório comparativo gerado e aprovado. Se não existe: fase pulada. | ⚠️ Informativo — não altera PERT |

---

## ESTRUTURA DE DIRETÓRIOS GERADA

```
business-inputs/business-projects/{PROJECT_ID_NAME}/
└── downstream-architecture-refinement/
    ├── DETAIL-LEVEL-PRD.md                       (F1)  🆕
    ├── DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md    (F2)  🆕
    ├── DETAIL-LEVEL-SECURITY-DEFINITION.md        (F3)  🆕
    ├── DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md (F4) 🆕
    ├── DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md      (F5)  🆕
    ├── DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md    (F6)  🆕
    ├── DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md      (F7)  🆕
    ├── BOTTOM-UP-PERT-ESTIMATE.md                 (F8)  🆕 ⭐
    ├── RESOURCE-ALLOCATION-PLAN.md                (F9)  🆕
    ├── RISK-ADJUSTED-ESTIMATE.md                  (F10) 🆕
    ├── SCOPE-SNAPSHOT.md                          (F11) 🆕 📸
    └── UPSTREAM-COMPARISON-REPORT.md              (F12) 🔗 (condicional)
```

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes. A diferença fundamental vs Upstream é o uso de skills de estimativa detalhada (`project-estimation`) e skills técnicos específicos (`java-*`, `spring-boot-*`, `c4-*`).

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `superpowers:brainstorming` | Brainstorming inicial da arquitetura de refinamento | Orquestração |
| 2 | `superpowers:executing-plans` | Execução do plano de fases com gates | Orquestração |
| 3 | `superpowers:writing-plans` | Escrita e refino do plano de execução | Orquestração |
| 4 | `superpowers:verification-before-completion` | Verificação de completude antes de cada COMPLIANCE | Qualidade |
| 5 | `project-estimation` | **Estimativa bottom-up com PERT three-point** — o core do downstream. Usa referências `bottom-up-estimation.md` e `three-point-estimation-pert.md` | Estimativa ⭐ |
| 6 | `discovery-process` | Estrutura o ciclo de refinamento: framing → synthesis → experiments | Discovery |
| 7 | `project-document-discovery` | Classifica o projeto em 10 dimensões e determina escopo proporcional | Discovery |
| 8 | `senior-architect` | Design detalhado de solução (C4 L2/L3) | Arquitetura |
| 9 | `c4-code` | Diagramas C4 Level 3 (Component) para serviços Java | Arquitetura |
| 10 | `c4-component` | Diagramas C4 Level 2 (Container) refinados | Arquitetura |
| 11 | `java-spring-boot` | Padrões Spring Boot para design detalhado | Tecnologia |
| 12 | `java-architect` | Padrões de arquitetura Java (packages, design patterns) | Tecnologia |
| 13 | `cloud-architect` | Topologia de infra detalhada com sizing | Infraestrutura |
| 14 | `senior-devops` | Pipeline specs, IaC detalhado, SLOs | DevOps |
| 15 | `senior-security` | Threat model STRIDE, OWASP ASVS, IAM specs | Segurança |
| 16 | `senior-data-engineer` | ERD completo, query patterns, particionamento | Dados |
| 17 | `senior-qa` | Matriz de cobertura por US, casos de teste | Qualidade |
| 18 | `gap-analysis` | Análise de gaps entre upstream e downstream | Análise |
| 19 | `documentation-writer` | Documentação do roadmap e execution history | Documentação |
| 20 | `context-manager` | Gestão de contexto entre fases longas (62 US) | Contexto |

### Skills Exclusivos do Downstream (não usados no Upstream)

| Skill | Motivo |
|:---|:---|
| `project-estimation` | **Diferencial principal** — estimativa bottom-up PERT substitui ROM manual |
| `c4-code` | Diagramas C4 L3 para componentes internos dos microservices |
| `c4-component` | Refinamento dos containers macro para nível de componente |
| `java-spring-boot` | Padrões Spring Boot para controllers, services, repositories |
| `java-architect` | Padrões de design Java (package structure, design patterns) |
| `context-manager` | Necessário para processar 62 US + 18 features sem perda de contexto |

---

## Localização dos Prompts das Fases

Os prompts de geração, gate e correção de cada fase estão na pasta `downstream-architecture-refinament/`:

```
.specs/prompts/downstream-architecture-refinament/
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-PRD.md                     🆕 F1
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-PRD.md                         🆕 F1
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-PRD.md                          🆕 F1
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-ARCHITECTURE-DEFINITION.md  🆕 F2
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-ARCHITECTURE-DEFINITION.md      🆕 F2
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-ARCHITECTURE-DEFINITION.md       🆕 F2
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SECURITY-DEFINITION.md      🆕 F3
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SECURITY-DEFINITION.md          🆕 F3
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SECURITY-DEFINITION.md           🆕 F3
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DATA-ARCHITECTURE-DEFINITION.md 🆕 F4
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DATA-ARCHITECTURE-DEFINITION.md     🆕 F4
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DATA-ARCHITECTURE-DEFINITION.md      🆕 F4
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DEVOPS-SRE-DEFINITION.md        🆕 F5
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DEVOPS-SRE-DEFINITION.md            🆕 F5
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DEVOPS-SRE-DEFINITION.md             🆕 F5
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-TEST-STRATEGY-DEFINITION.md     🆕 F6
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-TEST-STRATEGY-DEFINITION.md         🆕 F6
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-TEST-STRATEGY-DEFINITION.md          🆕 F6
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-INFRA-CLOUD-DEFINITION.md       🆕 F7
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-INFRA-CLOUD-DEFINITION.md           🆕 F7
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-INFRA-CLOUD-DEFINITION.md            🆕 F7
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-BOTTOM-UP-PERT-ESTIMATE.md      🆕 F8 ⭐
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-BOTTOM-UP-PERT-ESTIMATE.md           🆕 F8 ⭐
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-BOTTOM-UP-PERT-ESTIMATE.md            🆕 F8 ⭐
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RESOURCE-ALLOCATION.md          🆕 F9
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RESOURCE-ALLOCATION.md              🆕 F9
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RESOURCE-ALLOCATION.md               🆕 F9
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RISK-ADJUSTED-ESTIMATE.md       🆕 F10
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RISK-ADJUSTED-ESTIMATE.md           🆕 F10
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RISK-ADJUSTED-ESTIMATE.md            🆕 F10
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SCOPE-SNAPSHOT.md               🆕 F11 📸
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SCOPE-SNAPSHOT.md                   🆕 F11 📸
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SCOPE-SNAPSHOT.md                    🆕 F11 📸
├── PROMPT-GENERATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-UPSTREAM-COMPARISON.md          🆕 F12 🔗
├── PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-UPSTREAM-COMPARISON.md              🆕 F12 🔗
├── PROMPT-FIX-DOWNSTREAM-ARCHITECTURE-REFINEMENT-UPSTREAM-COMPARISON.md               🆕 F12 🔗
└── ... (36 prompts no total: 12 GENERATE + 12 GATE + 12 FIX)
```

---

## DIFERENÇAS FUNDAMENTAIS: UPSTREAM vs DOWNSTREAM

| Dimensão | Upstream (Discovery) | Downstream (Refinement) |
|:---|:---|:---|
| **Quando executar** | Pré-GO/NO-GO (decisão de investimento) | A qualquer momento (pré ou pós-GO) quando precisar de estimativa detalhada |
| **Dependência** | Requer só docs de negócio (Charter+BRD+Épicos) | Requer docs de negócio completos (Charter+BRD+Épicos+Features+US) |
| **Upstream é pré-requisito?** | — | **Não.** Opcional. Se existir, vira referência para design (Bloco A) e comparação (Fase 13) |
| **Input** | Charter + BRD + Épicos (escopo macro) | Charter + BRD + Épicos + Features + US (escopo completo) |
| **Arquitetura** | C4 L1 (System Context), containers macro | C4 L2/L3 (Container/Component), ADRs detalhados |
| **Estimativa** | ROM ±50% (manual, por solução) | **PERT Bottom-Up ±15-25% (US individuais, independente)** |
| **Estimativa depende de anterior?** | N/A (primeira estimativa) | **Não. 100% independente.** Calculada do zero, US por US |
| **Skills de estimativa** | Nenhum específico | `project-estimation` + referências `bottom-up-estimation.md` + `three-point-estimation-pert.md` |
| **Skills técnicos** | Genéricos (`senior-*`) | Específicos (`java-spring-boot`, `c4-code`, `c4-component`) |
| **Output** | 11 artefatos Discovery-Level | 12 artefatos Detail-Level + Scope Snapshot + Cross-Check Report |
| **Gate final** | GO/NO-GO (decisão de investimento) | ESTIMATE-READY (estimativa aprovada como baseline) |
| **Confiança** | ±50% (suficiente para aprovar orçamento) | ±15-25% (suficiente para compromisso de prazo) |
| **Diretório** | `upstream-architecture-discovery/` | `downstream-architecture-refinement/` |
| **Vínculo entre roadmaps** | — | **Apenas Fase 13 (Cross-Check Report)** — condicional, não altera estimativa |

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial: roadmap de Downstream Architecture Refinement com 13 fases em 5 blocos + Gate READY-FOR-EXECUTION. | Time de Arquitetura |
| 1.1 | 31/07/2026 | Independência do upstream: removida pré-condição obrigatória, adicionado Bloco D com Fase 13 (Cross-Check Report) condicional, estimativa 100% independente. | Time de Arquitetura |
| 2.0 | 31/07/2026 | **Foco em viabilidade e estimativa:** Removido Bloco C (Sprint-Ready Contracts). F11 renomeada para SCOPE-SNAPSHOT (foto do escopo, sem planejamento de sprints). F12 absorve Cross-Check (era F13). Gate renomeado para ESTIMATE-READY. Removida pasta contracts/. 12 fases em 4 blocos, 36 prompts. O planejamento de sprints e contratos técnicos são responsabilidade do PROJECT-TECHNICAL-DEFINITIONS. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados na seção Skills Utilizados. Este roadmap é focado em **análise de viabilidade e estimativa** pelo time de TI. É **independente** do `PROMPT-ROADMAP-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY.md` — o único vínculo é a Fase 12 (Cross-Check Report condicional). O planejamento de sprints e contratos técnicos são responsabilidade do `PROMPT-ROADMAP-GENERATE-PROJECT-TECHNICAL-DEFINITIONS.md`.*
