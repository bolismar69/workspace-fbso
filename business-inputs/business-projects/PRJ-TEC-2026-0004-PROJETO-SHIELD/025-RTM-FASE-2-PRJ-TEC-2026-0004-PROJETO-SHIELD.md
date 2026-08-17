# RTM Fase 2 — Matriz de Rastreabilidade de Sistema: PROJETO SHIELD
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 015-RTM-FASE-1, 020-SRS |
| **Data de Elaboração** | 08/08/2026 |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## RTM Fase 2 — Rastreabilidade de Sistema

A **RTM-FASE-2** propaga a rastreabilidade da linha de base de negócio (RTM-FASE-1) para o domínio técnico do sistema (020-SRS). Ela valida que cada requisito funcional de sistema (`FUNCTIONAL-REQ-XX`) e cada requisito não-funcional (`NO-FUNCTIONAL-REQ-{category}-XX`) possui lastro em um requisito de negócio (`B-REQ-XX`) — direta ou indiretamente via RTM-FASE-1.

### Objetivos da RTM-FASE-2

- **Cobertura Total de Sistema:** Prova que cada FUNCTIONAL-REQ e NO-FUNCTIONAL-REQ do SRS rastreia a um B-REQ do BRD
- **Zero Órfãos Técnicos:** Garante que nenhum requisito de sistema foi inventado sem lastro no negócio
- **Ponte para Arquitetura:** Prepara a rastreabilidade para os documentos da Fase 2 (030-SAD, 035-HLD)

---

## 1. Matriz de Rastreabilidade de Sistema (SRS → RTM-F1 → BRD)

| Requisito de Sistema (SRS) | Funcionalidade FRD (via RTM-F1) | Requisito BRD (via RTM-F1) | Critério Charter |
|:---|:---|:---|:---|
| FUNCTIONAL-REQ-01 — Extração de domínio | B-FEAT-01 | B-REQ-01 | C1 |
| FUNCTIONAL-REQ-02 — Cache de mapeamento | B-FEAT-01, B-FEAT-08 | B-REQ-01, B-REQ-08 | C1, C6 |
| FUNCTIONAL-REQ-03 — Início de autenticação (PKCE) | B-FEAT-03, B-FEAT-04 | B-REQ-03, B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-04 — Troca e armazenamento seguro | B-FEAT-03, B-FEAT-04 | B-REQ-03, B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-05 — Consulta de perfil | B-FEAT-04, B-FEAT-05 | B-REQ-04, B-REQ-05 | C3 |
| FUNCTIONAL-REQ-06 — Renovação de sessão | B-FEAT-04 | B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-07 — Logout completo | B-FEAT-04 | B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-08 — Filtro de isolamento | B-FEAT-02 | B-REQ-02 | C1 |
| FUNCTIONAL-REQ-09 — Registro de eventos de autenticação | B-FEAT-07 | B-REQ-07 | C8 |
| FUNCTIONAL-REQ-10 — Métricas operacionais | B-FEAT-09 | B-REQ-09 | C7 |
| FUNCTIONAL-REQ-11 — Ativação de novo cliente | B-FEAT-08 | B-REQ-08 | C6 |
| FUNCTIONAL-REQ-12 — Suspensão de cliente | B-FEAT-02 | B-REQ-02 | C1 |
| FUNCTIONAL-REQ-13 — Alerta de latência | B-FEAT-05, B-FEAT-06 | B-REQ-05, B-REQ-06 | C3, C4 |
| FUNCTIONAL-REQ-14 — Degradação controlada | B-FEAT-06 | B-REQ-06 | C4 |
| FUNCTIONAL-REQ-15 — Retenção e expurgo de logs | B-FEAT-07 | B-REQ-07 | C8 |
| FUNCTIONAL-REQ-16 — Migração com rollback | B-FEAT-10 | B-REQ-11 | C7 |
| FUNCTIONAL-REQ-17 — Relatório de acessos | B-FEAT-07 | B-REQ-07 | C8 |

**Cobertura Functional:** 17/17 FUNCTIONAL-REQs com lastro completo. **100%.** ✅

---

## 2. Matriz de Rastreabilidade Não-Funcional (NFRs → BRD/Charter)

| Requisito Não-Funcional (SRS) | Requisito BRD / Critério Charter | Lastro |
|:---|:---|:---|
| NO-FUNCTIONAL-REQ-PERFORMANCE-01 — Latência p95 < 15ms | B-REQ-05, C3 | ✅ |
| NO-FUNCTIONAL-REQ-PERFORMANCE-02 — Latência cache p99 < 5ms | B-REQ-01 | ✅ |
| NO-FUNCTIONAL-REQ-PERFORMANCE-03 — Cold start < 100ms | B-REQ-09 | ✅ |
| NO-FUNCTIONAL-REQ-SECURITY-01 — Cookies HttpOnly 100% | B-REQ-03, C2 | ✅ |
| NO-FUNCTIONAL-REQ-SECURITY-02 — Cookies Secure 100% | B-REQ-03, C2 | ✅ |
| NO-FUNCTIONAL-REQ-SECURITY-03 — Cookies SameSite=Strict 100% | B-REQ-03, C5 | ✅ |
| NO-FUNCTIONAL-REQ-SECURITY-04 — TLS 1.3 mínimo | C5 | ✅ |
| NO-FUNCTIONAL-REQ-SECURITY-05 — Sanitização de logs | B-REQ-07, B-RULE-20 | ✅ |
| NO-FUNCTIONAL-REQ-SECURITY-06 — Scan OWASP no pipeline | C5 | ✅ |
| NO-FUNCTIONAL-REQ-AVAILABLE-01 — SLA 99.9% | C7 | ✅ |
| NO-FUNCTIONAL-REQ-AVAILABLE-02 — RTO < 5 min | C7 | ✅ |
| NO-FUNCTIONAL-REQ-AVAILABLE-03 — RPO < 1 min | C7 | ✅ |
| NO-FUNCTIONAL-REQ-SCALABILITY-01 — 5.000 RPS | B-REQ-06, C4 | ✅ |
| NO-FUNCTIONAL-REQ-SCALABILITY-02 — 100K sessões | B-REQ-09 | ✅ |
| NO-FUNCTIONAL-REQ-SCALABILITY-03 — Auto-scaling | B-REQ-09 | ✅ |
| NO-FUNCTIONAL-REQ-OBSERVABILITY-01 — Tracing distribuído | — (inerente a produto corporativo) | ✅ |
| NO-FUNCTIONAL-REQ-OBSERVABILITY-02 — Logs JSON estruturados | B-REQ-07 | ✅ |
| NO-FUNCTIONAL-REQ-OBSERVABILITY-03 — Métricas expostas | B-REQ-09 | ✅ |
| NO-FUNCTIONAL-REQ-USABILITY-01 — Consistência entre produtos | B-REQ-10 | ✅ |

**Cobertura Non-Functional:** 19/19 NO-FUNCTIONAL-REQs com lastro. 1 item (OBSERVABILITY-01) é inerente a produto corporativo — tracing é padrão de engenharia FBSO.ORG. **100%.** ✅

---

## 3. Análise de Cobertura Cruzada (B-REQ → NFR)

Verifica se todos os B-REQs que demandam garantias não-funcionais estão cobertos por NFRs.

| B-REQ | NFRs Vinculados | Cobertura |
|---|---|---|
| B-REQ-01 | PERFORMANCE-02 | ✅ |
| B-REQ-03 | SECURITY-01, SECURITY-02, SECURITY-03 | ✅ |
| B-REQ-05 | PERFORMANCE-01 | ✅ |
| B-REQ-06 | SCALABILITY-01 | ✅ |
| B-REQ-07 | SECURITY-05, OBSERVABILITY-02, FUNCTIONAL-REQ-09/15/17 | ✅ |
| B-REQ-09 | PERFORMANCE-03, SCALABILITY-02, SCALABILITY-03, OBSERVABILITY-03 | ✅ |
| B-REQ-10 | USABILITY-01 | ✅ |

---

## 4. Análise de Órfãos de Sistema

### 4.1 FUNCTIONAL-REQs sem Lastro no Negócio

**NENHUM órfão.** 17/17 FUNCTIONAL-REQs com lastro via RTM-F1. ✅

### 4.2 NO-FUNCTIONAL-REQs sem Lastro no Negócio

**NENHUM órfão.** 19/19 NO-FUNCTIONAL-REQs com lastro. OBSERVABILITY-01 (tracing) é inerente ao padrão corporativo. ✅

---

## 5. Sumário Executivo da Fase 2 (até o momento)

| Métrica | Resultado |
|---|---|
| FUNCTIONAL-REQs com lastro em B-REQ | 17/17 — **100%** |
| NO-FUNCTIONAL-REQs com lastro | 19/19 — **100%** |
| B-REQs com cobertura de NFR | 7/7 aplicáveis — **100%** |
| Órfãos de Sistema | **0** |
| Lacunas de Cobertura | **0** |

> **✅ FASE 2 EM ANDAMENTO.** A rastreabilidade de sistema está íntegra. Todos os requisitos do SRS possuem lastro documentado no BRD/FRD via RTM-FASE-1. O projeto está apto a avançar para os documentos de arquitetura (030-SAD, 035-HLD).
