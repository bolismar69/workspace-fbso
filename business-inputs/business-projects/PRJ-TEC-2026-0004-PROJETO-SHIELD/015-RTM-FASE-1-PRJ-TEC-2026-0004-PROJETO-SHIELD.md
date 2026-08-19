# RTM Fase 1 — Matriz de Rastreabilidade de Negócio: PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 001-PROJECT-CHARTER, 002-STAKEHOLDER-MAP, 003-PERSONAS-JORNADAS, 004-MAPEAMENTO-AS-IS-TO-BE, 005-BRD, 010-FRD |
| **Data de Elaboração** | 08/08/2026 |
| **Versão** | 1.7 — Aprovação humana (19/08/2026, P1=SIM/P2–P4=NÃO — atalho OK) — documento congelado em COMPLIANCE |
| **Metodologia** | WATERFALL |

---

### Siglas definidas no documento

- **B-PERSONA-** _(Business Persona)_: Persona de negócio — origem 003-PERSONAS-JORNADAS
- **B-JOURNEY-** _(Business Journey)_: Jornada de negócio — origem 003-PERSONAS-JORNADAS
- **B-PROCESS-** _(Business Process)_: Processo de negócio (AS-IS/TO-BE) — origem 004-MAPEAMENTO-AS-IS-TO-BE
- **B-GAP-ANALYSIS-** _(Business Gap Analysis)_: Gap AS-IS → TO-BE — origem 004-MAPEAMENTO-AS-IS-TO-BE
- **B-REQ-** / **B-FEAT-** / **B-RULE-** / **B-UC-** / **B-REQ-SECURITY-**: siglas definidas no 005-BRD e no 010-FRD

---

## RTM Fase 1 — Rastreabilidade de Negócio

A **RTM-FASE-1** é o instrumento de governança que sela a **Linha de Base de Escopo Funcional** ao final da Fase 1. Ela atua como validador contratual de negócio antes que qualquer documento técnico (020-SRS, 030-SAD, 035-HLD, 040-LLD) seja iniciado.

### Objetivos da RTM-FASE-1

- **Cobertura Total:** Prova que cada Critério de Sucesso do Charter (C1-C8) tem cobertura de requisitos de negócio (B-REQ), e cada B-REQ tem funcionalidades (B-FEAT), regras (B-RULE) e casos de uso (B-UC) correspondentes
- **Zero Órfãos:** Garante que nenhum B-FEAT, B-RULE ou B-UC foi criado sem lastro em um requisito de negócio explícito do BRD
- **Análise de Impacto (CCR):** Se um requisito de negócio mudar, esta matriz aponta imediatamente quais funcionalidades, regras e casos de uso são impactados

---

## 1. Matriz de Rastreabilidade Primária (Charter → BRD → FRD)

| Critério Charter | Requisito BRD | Funcionalidade FRD | Regras Vinculadas | Casos de Uso | Cobertura |
|:---|:---|:---|:---|:---|:---|
| C1 — Segurança entre Clientes | B-REQ-01 | B-FEAT-01 — Reconhecimento pelo Domínio | B-RULE-01, B-RULE-10, B-RULE-11 | B-UC-01 | ✅ |
| C1 — Segurança entre Clientes | B-REQ-02 | B-FEAT-02 — Isolamento de Ambientes | B-RULE-02, B-RULE-06, B-RULE-08, B-RULE-12, B-RULE-13 | B-UC-03 | ✅ |
| C2 — Proteção de Credenciais | B-REQ-03 | B-FEAT-03 — Proteção de Credenciais | B-RULE-03, B-RULE-14 | B-UC-01 | ✅ |
| C5 — Cobertura a Ataques Cibernéticos | B-REQ-03 | B-FEAT-03 — Proteção de Credenciais | B-RULE-03, B-RULE-14 | B-UC-01 | ✅ |
| C5 — Cobertura a Ataques Cibernéticos | B-REQ-10 | B-FEAT-04 — Portal de Acesso Padronizado | B-RULE-04, B-RULE-05, B-RULE-15, B-RULE-16 | B-UC-01 | ✅ |
| C6 — Tempo para Adicionar Novo Cliente | B-REQ-04 | B-FEAT-04 — Portal de Acesso Padronizado | B-RULE-04, B-RULE-05, B-RULE-15, B-RULE-16 | B-UC-01 | ✅ |
| C6 — Tempo para Adicionar Novo Cliente | B-REQ-08 | B-FEAT-08 — Ativação de Novo Cliente | B-RULE-07, B-RULE-21, B-RULE-22 | B-UC-02 | ✅ |
| C3 — Velocidade de Resposta | B-REQ-05 | B-FEAT-05 — Resposta Rápida | B-RULE-17 | B-UC-01 | ✅ |
| C4 — Capacidade de Atender Picos | B-REQ-06 | B-FEAT-06 — Suporte a Picos | B-RULE-18 | B-UC-05 | ✅ |
| C4 — Capacidade de Atender Picos | B-REQ-09 | B-FEAT-09 — Adaptação ao Crescimento | B-RULE-25 | B-UC-06 | ✅ |
| C7 — Disponibilidade da Plataforma | B-REQ-09 | B-FEAT-09 — Adaptação ao Crescimento | B-RULE-25 | B-UC-06 | ✅ |
| C7 — Disponibilidade da Plataforma | B-REQ-11 | B-FEAT-10 — Transição Transparente | B-RULE-23, B-RULE-24 | B-UC-04 | ✅ |
| C8 — Rastreabilidade de Acessos | B-REQ-07 | B-FEAT-07 — Registro de Auditoria | B-RULE-19, B-RULE-20 | B-UC-07 | ✅ |
| Premissas 1 e 2 (Charter Seção 7) | B-REQ-11 | B-FEAT-10 — Transição Transparente | B-RULE-23, B-RULE-24 | B-UC-04 | ✅ |

---

## 2. Análise de Cobertura

### 2.1 Requisitos do BRD → FRD

| Requisito BRD | FEATs Vinculados | Regras Vinculadas | UCs Vinculados | Status |
|---|---|---|---|---|
| B-REQ-01 | B-FEAT-01 | B-RULE-01, B-RULE-10, B-RULE-11 | B-UC-01 | ✅ Coberto |
| B-REQ-02 | B-FEAT-02 | B-RULE-02, B-RULE-06, B-RULE-08, B-RULE-12, B-RULE-13 | B-UC-03 | ✅ Coberto |
| B-REQ-03 | B-FEAT-03 | B-RULE-03, B-RULE-14 | B-UC-01 | ✅ Coberto |
| B-REQ-04 | B-FEAT-04 | B-RULE-04, B-RULE-05, B-RULE-15, B-RULE-16 | B-UC-01 | ✅ Coberto |
| B-REQ-05 | B-FEAT-05 | B-RULE-17 | B-UC-01 | ✅ Coberto |
| B-REQ-06 | B-FEAT-06 | B-RULE-18 | B-UC-05 | ✅ Coberto |
| B-REQ-07 | B-FEAT-07 | B-RULE-19, B-RULE-20 | B-UC-07 | ✅ Coberto |
| B-REQ-08 | B-FEAT-08 | B-RULE-07, B-RULE-21, B-RULE-22 | B-UC-02 | ✅ Coberto |
| B-REQ-09 | B-FEAT-09 | B-RULE-25 | B-UC-06 | ✅ Coberto |
| B-REQ-10 | B-FEAT-04 | B-RULE-04, B-RULE-05, B-RULE-15, B-RULE-16 | B-UC-01 | ✅ Coberto |
| B-REQ-11 | B-FEAT-10 | B-RULE-23, B-RULE-24 | B-UC-04 | ✅ Coberto |

**Resultado:** 11/11 B-REQs cobertos por pelo menos um B-FEAT. **100% de cobertura.**

### 2.2 Critérios de Sucesso do Charter → BRD

| Critério Charter | B-REQs Vinculados | Cobertura |
|---|---|---|
| C1 — Segurança entre Clientes | B-REQ-01, B-REQ-02 | ✅ 2 requisitos |
| C2 — Proteção de Credenciais | B-REQ-03 | ✅ 1 requisito |
| C3 — Velocidade de Resposta | B-REQ-05 | ✅ 1 requisito |
| C4 — Capacidade de Atender Picos | B-REQ-06, B-REQ-09 | ✅ 2 requisitos |
| C5 — Cobertura a Ataques Cibernéticos | B-REQ-03, B-REQ-10 | ✅ 2 requisitos |
| C6 — Tempo para Adicionar Novo Cliente | B-REQ-04, B-REQ-08 | ✅ 2 requisitos |
| C7 — Disponibilidade da Plataforma | B-REQ-09, B-REQ-11 | ✅ 2 requisitos |
| C8 — Rastreabilidade de Acessos | B-REQ-07 | ✅ 1 requisito |

**Resultado:** 8/8 critérios do Charter cobertos por pelo menos um B-REQ. **100% de cobertura.**

### 2.3 Contexto de Origem — B-JOURNEY / B-PERSONA / B-PROCESS / B-GAP-ANALYSIS

Jornadas, personas, processos e gaps que fundamentam cada requisito de negócio.

**Regra de Cobertura (determinística):**
- ✅ = as 4 colunas estão preenchidas **E** cada ID citado existe no documento de origem (003/004) e pertence ao B-REQ da linha
- ⚠️ = alguma coluna vazia ou ID não localizado
- ❌ = nenhum contexto (B-REQ sem lastro em 003/004)

| Requisito BRD | B-JOURNEY | B-PERSONA | B-PROCESS | B-GAP-ANALYSIS | Cobertura |
|---|---|---|---|---|---|
| B-REQ-01 | B-JOURNEY-01 (etapa 1) | B-PERSONA-01, B-PERSONA-02, B-PERSONA-03 | B-PROCESS-01 | B-GAP-ANALYSIS-01 | ✅ |
| B-REQ-02 | B-JOURNEY-01 (etapa 3) | B-PERSONA-01 | B-PROCESS-01 | B-GAP-ANALYSIS-02 | ✅ |
| B-REQ-03 | B-JOURNEY-01 (etapa 2) | B-PERSONA-01, B-PERSONA-02, B-PERSONA-03 | B-PROCESS-01 | B-GAP-ANALYSIS-03 | ✅ |
| B-REQ-04 | B-JOURNEY-01 (etapas 2/4) | B-PERSONA-01, B-PERSONA-02, B-PERSONA-03 | B-PROCESS-01 | B-GAP-ANALYSIS-01 | ✅ |
| B-REQ-05 | B-JOURNEY-01 (etapa 3) | B-PERSONA-02, B-PERSONA-03 | B-PROCESS-01 | B-GAP-ANALYSIS-04 | ✅ |
| B-REQ-06 | B-JOURNEY-01 (etapa 3) | B-PERSONA-03 | B-PROCESS-01 | B-GAP-ANALYSIS-04 | ✅ |
| B-REQ-07 | B-JOURNEY-03 (etapa 2) | B-PERSONA-01 | B-PROCESS-03 | B-GAP-ANALYSIS-06 | ✅ |
| B-REQ-08 | B-JOURNEY-02 (etapas 1–3) | B-PERSONA-01 | B-PROCESS-02 | B-GAP-ANALYSIS-05 | ✅ |
| B-REQ-09 | B-JOURNEY-01 (etapa 3) | B-PERSONA-02, B-PERSONA-03 | B-PROCESS-01 | B-GAP-ANALYSIS-04 | ✅ |
| B-REQ-10 | B-JOURNEY-01 (etapa 1) | B-PERSONA-01, B-PERSONA-02, B-PERSONA-03 | B-PROCESS-01 | B-GAP-ANALYSIS-01 | ✅ |
| B-REQ-11 | B-JOURNEY-04 (etapas 1–4) | B-PERSONA-04 | B-PROCESS-04 | B-GAP-ANALYSIS-07 | ✅ |

**Cobertura reversa (zero-órfãos de contexto):** todo B-JOURNEY (4/4), B-PERSONA (4/4), B-PROCESS (4/4) e B-GAP-ANALYSIS (7/7) dos 003/004 aparece em pelo menos uma linha desta seção. ✅

> **NOTA:** O reuso de um mesmo B-GAP-ANALYSIS/B-PROCESS/B-JOURNEY por múltiplos B-REQs é legítimo (um gap pode derivar vários requisitos) e não penaliza a Cobertura. O `B-REQ-11` possui contexto completo de origem via B-JOURNEY-04/B-PERSONA-04/B-PROCESS-04/B-GAP-ANALYSIS-07, espelhando o B-UC-04 do 010-FRD. Os vínculos `B-REQ-02`, `B-REQ-09` e `B-REQ-10` → `B-JOURNEY-01` são indiretos e legítimos: a jornada relaciona os critérios C1/C4 (001, Seção 6) e a Seção 6 do 005-BRD mapeia esses B-REQs aos mesmos critérios — encadeamento registrado nesta NOTA. Os `B-REQ-SECURITY-*` são tratados na subseção 2.4.

---

### 2.4 Contexto de Referência — B-REQ-SECURITY (facetas de segurança/compliance)

Os requisitos de segurança e compliance (005-BRD, subseção 1.3) são **facetas** dos requisitos de negócio e restrições a que estão vinculados — **não geram linha própria** na matriz primária (Seção 1) nem em 2.1/2.2, e **não contam como órfãos**.

| ID | Requisito de Segurança/Compliance | Regulação/Política | Referência de Negócio Vinculada | Cobertura |
|---|---|---|---|---|
| B-REQ-SECURITY-01 | Isolamento estrito entre ambientes de clientes | LGPD (segurança dos dados), política interna de segurança | B-REQ-02 | ✅ |
| B-REQ-SECURITY-02 | Credenciais jamais expostas fora do ambiente seguro de autenticação | Política interna de segurança | B-REQ-03 | ✅ |
| B-REQ-SECURITY-03 | Registro de auditoria de acessos sem dados sensíveis, com retenção mínima de 6 meses | LGPD (responsabilização e auditoria) | B-REQ-07 | ✅ |
| B-REQ-SECURITY-04 | Conformidade com LGPD para dados processados em nuvem sem datacenter no Brasil | LGPD — diretrizes do Jurídico | B-LIMIT-04 | ✅ |

**Cobertura:** 4/4 B-REQ-SECURITY com referência de negócio registrada no 005-BRD (subseção 1.3). ✅

---

## 3. Análise de Órfãos (Gold-Plating)

### 3.1 Funcionalidades sem Lastro no BRD

| Funcionalidade FRD | Origem BRD | Status |
|---|---|---|
| B-FEAT-01 | B-REQ-01 | ✅ Vinculado |
| B-FEAT-02 | B-REQ-02 | ✅ Vinculado |
| B-FEAT-03 | B-REQ-03 | ✅ Vinculado |
| B-FEAT-04 | B-REQ-04, B-REQ-10 | ✅ Vinculado |
| B-FEAT-05 | B-REQ-05 | ✅ Vinculado |
| B-FEAT-06 | B-REQ-06 | ✅ Vinculado |
| B-FEAT-07 | B-REQ-07 | ✅ Vinculado |
| B-FEAT-08 | B-REQ-08 | ✅ Vinculado |
| B-FEAT-09 | B-REQ-09 | ✅ Vinculado |
| B-FEAT-10 | B-REQ-11 | ✅ Vinculado |

**NENHUM órfão encontrado.** 10/10 funcionalidades com lastro em pelo menos um B-REQ. ✅

### 3.2 Casos de Uso sem Lastro em Funcionalidade

| Caso de Uso | Funcionalidade Vinculada | Status |
|---|---|---|
| B-UC-01 | B-FEAT-01, B-FEAT-03, B-FEAT-04, B-FEAT-05 | ✅ Vinculado |
| B-UC-02 | B-FEAT-08 | ✅ Vinculado |
| B-UC-03 | B-FEAT-02 | ✅ Vinculado |
| B-UC-04 | B-FEAT-10 | ✅ Vinculado |
| B-UC-05 | B-FEAT-06 | ✅ Vinculado |
| B-UC-06 | B-FEAT-09 | ✅ Vinculado |
| B-UC-07 | B-FEAT-07 | ✅ Vinculado |

**NENHUM órfão encontrado.** 7/7 casos de uso com lastro em pelo menos um B-FEAT. ✅

### 3.3 Regras de Negócio sem Lastro

Total de regras: B-RULE-01 a B-RULE-25 (25 regras).
- B-RULE-01 a B-RULE-09: definidas no 005-BRD, todas vinculadas a B-REQ ✅
- B-RULE-10 a B-RULE-25: definidas no 010-FRD, todas vinculadas a B-FEAT ✅

**NENHUM órfão encontrado.** 25/25 regras com lastro. ✅

---

## 4. Sumário Executivo da Fase 1

| Métrica | Resultado |
|---|---|
| Critérios do Charter cobertos por B-REQ | 8/8 — **100%** |
| B-REQs cobertos por B-FEAT | 11/11 — **100%** |
| B-FEATs com lastro em B-REQ | 10/10 — **100%** |
| Casos de Uso com lastro em B-FEAT | 7/7 — **100%** |
| Regras de Negócio com lastro | 25/25 — **100%** |
| B-JOURNEY com lastro em B-REQ (2.3) | 4/4 — **100%** |
| B-PERSONA com lastro em B-REQ (2.3) | 4/4 — **100%** |
| B-PROCESS com lastro em B-REQ (2.3) | 4/4 — **100%** |
| B-GAP-ANALYSIS com lastro em B-REQ (2.3) | 7/7 — **100%** |
| B-REQs com Contexto de Origem completo (2.3) | 11/11 — **100%** |
| Órfãos (Gold-Plating) | **0** |
| Lacunas de Cobertura | **0** |
| B-REQ-SECURITY com referência registrada (2.4) | 4/4 — **100%** |

> **✅ LINHA DE BASE DE ESCOPO FUNCIONAL SELADA.** A Fase 1 está completa e íntegra. Todos os artefatos de negócio (001-Charter, 002-Stakeholder-Map, 003-Personas-Jornadas, 004-As-Is/To-Be, 005-BRD, 010-FRD) estão 100% rastreáveis entre si. O projeto está apto a avançar para a Fase 2 — Especificação de Sistema e Arquitetura Macro.

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 08/08/2026 | Baseline inicial de rastreabilidade de negócio — Fase 1 selada | Time de Negócios / Orquestrador WATERFALL v2.0 |
| 1.1 | 19/08/2026 | Atualização (skill waterfall-business-documents): Documentos Base incluem 003 e 004; seção 2.3 Contexto de Origem adicionada; contagem de regras corrigida (25/25); sumário com os 6 artefatos | Time de Negócios / skill waterfall-business-documents |
| 1.2 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO) — Fase 1 congelada: 7/7 documentos em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
| 1.3 | 19/08/2026 | Correção cirúrgica (review FASE 1): linha 2.3 do B-REQ-11 alinhada à regra determinística (❌ justificado — origem nas Premissas 1 e 2 do Charter; 004 B-GAP-ANALYSIS-04 estendido para derivar B-REQ-09); subseção 2.4 (B-REQ-SECURITY × Referência) adicionada; Sumário Executivo com B-REQ-SECURITY | Time de Negócios / skill waterfall-business-documents |
| 1.4 | 19/08/2026 | Correção cirúrgica (review FASE 1, F1 — decisão humana): B-REQ-11 passou a ter contexto completo de origem — B-JOURNEY-04/B-PERSONA-04 (003) e B-PROCESS-04/B-GAP-ANALYSIS-07 (004); linha 2.3 ✅; cobertura reversa 4/4, 4/4, 4/4 e 7/7; Sumário 11/11 | Time de Negócios / skill waterfall-business-documents |
| 1.5 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO — atalho OK) — documento congelado em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
| 1.6 | 19/08/2026 | Correção cirúrgica (update pós-selo, F2/F3/F5/F6): linha do B-REQ-11 com Premissas 1 e 2 (Seção 1); NOTA 2.3 documenta vínculos indiretos (B-REQ-02/09/10 → B-JOURNEY-01 via C1/C4); título com subtítulo do produto; campo Versão do cabeçalho alinhado | Time de Negócios / skill waterfall-business-documents |
| 1.7 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO — atalho OK) — correções do update pós-selo aprovadas; documento congelado em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
