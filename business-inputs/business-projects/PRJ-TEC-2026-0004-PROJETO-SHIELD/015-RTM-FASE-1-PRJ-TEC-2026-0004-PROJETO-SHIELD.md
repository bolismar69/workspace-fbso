# RTM Fase 1 — Matriz de Rastreabilidade de Negócio: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 001-PROJECT-CHARTER, 002-STAKEHOLDER-MAP, 005-BRD, 010-FRD |
| **Data de Elaboração** | 08/08/2026 |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

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
| Premissa 1 (Charter Seção 7) | B-REQ-11 | B-FEAT-10 — Transição Transparente | B-RULE-23, B-RULE-24 | B-UC-04 | ✅ |

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
- B-RULE-10 a B-RULE-24: definidas no 010-FRD, todas vinculadas a B-FEAT ✅

**NENHUM órfão encontrado.** 24/24 regras com lastro. ✅

---

## 4. Sumário Executivo da Fase 1

| Métrica | Resultado |
|---|---|
| Critérios do Charter cobertos por B-REQ | 8/8 — **100%** |
| B-REQs cobertos por B-FEAT | 11/11 — **100%** |
| B-FEATs com lastro em B-REQ | 10/10 — **100%** |
| Casos de Uso com lastro em B-FEAT | 7/7 — **100%** |
| Regras de Negócio com lastro | 25/25 — **100%** |
| Órfãos (Gold-Plating) | **0** |
| Lacunas de Cobertura | **0** |

> **✅ LINHA DE BASE DE ESCOPO FUNCIONAL SELADA.** A Fase 1 está completa e íntegra. Todos os artefatos de negócio (001-Charter, 002-Stakeholder-Map, 005-BRD, 010-FRD) estão 100% rastreáveis entre si. O projeto está apto a avançar para a Fase 2 — Especificação de Sistema e Arquitetura Macro.

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 08/08/2026 | Baseline inicial de rastreabilidade de negócio — Fase 1 selada | Time de Negócios / Orquestrador WATERFALL v2.0 |
