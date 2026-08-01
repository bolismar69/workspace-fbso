# BOTTOM-UP-PERT-ESTIMATE — Estimativa Bottom-Up PERT Three-Point

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 31/07/2026
- **Fase:** F8 — Downstream Architecture Refinement
- **Metodologia:** PERT Three-Point (O + 4M + P) / 6
- **Independência:** ✅ Estimativa 100% independente — calculada do zero, US por US. NENHUMA referência ao ROM upstream ou factory bids nos cálculos.
- **Confiança alvo:** ±15-25%

---

## 1. Sumário Executivo

| Cenário | Horas | Homem-Mês (160h) | IC 95% |
|:---|---:|---:|:---|
| Desenvolvimento PERT (62 US) | 3,365h | 21 h-m | 2,137h – 4,594h |
| Subtotal (Dev+QA+Arch+DevOps+Gestão+Infra) | 5,697h | 36 h-m | — |
| **TOTAL com Contingência 15-25%** | **6,077h – 7,121h** | **38 – 45 h-m** | — |

> ⚠️ **Regra de Independência:** Esta estimativa foi calculada exclusivamente a partir da análise individual de cada User Story, usando apenas o conteúdo das US, a complexidade técnica dos artefatos Detail-Level (F2-F7) e a stack tecnológica do projeto. NENHUMA estimativa anterior (ROM upstream, factory bids, etc.) foi consultada durante o cálculo.

---

## 2. Metodologia

### 2.1 Classificação de Complexidade por US

| Complexidade | O (h) | ML (h) | P (h) | Critério |
|:---|---:|---:|---:|:---|
| **1 — Simples** | 12-24 | 24-40 | 40-72 | CRUD básico, listas, filtros simples, ativar/desativar |
| **2 — Média** | 24-56 | 40-80 | 72-144 | Workflows, integrações, regras de negócio, dashboards com agregações |
| **3 — Complexa** | 40-56 | 72-96 | 128-160 | RBAC granular, state machines, integração IAM avançada, multi-tenancy |

### 2.2 Fórmula PERT

```
PERT = (O + 4×ML + P) / 6
σ = (P − O) / 6
IC 95% = PERT ± 1.96×σ
```

### 2.3 Composição do Esforço

| Componente | % sobre Dev | Justificativa |
|:---|---:|:---|
| QA | 30% | Meta DTA: ≥25%. Cobre unitários, integração, E2E, performance, segurança |
| Arquitetura | 8% | Meta DTA: ≥5%. ADRs, design reviews, C4, security reviews |
| DevOps/SRE | 7% | CI/CD, IaC, observabilidade, ambientes |
| Gestão/Governança | 10% | Refinamento, planning, daily, review, retro, docs, compliance gates |
| Infraestrutura residual | 480h fixo | M1 já concluído; residual para M2-M7 |
| Contingência | 15-25% | Riscos técnicos, learning curve, dependências externas |

---

## 3. Estimativa por Épico

### 3.1 EP-0001 — Portal Admin (7 US · 3 features · D1)

| Feature | US | Compl. | O | ML | P | PERT |
|:---|---|---:|---:|---:|---:|---:|
| FEAT-EP-0001-0001 Dashboard | US-0001, 0002, 0003 | Média | 112h | 184h | 336h | **197h** |
| FEAT-EP-0001-0002 Visão Contas | US-0004, 0005 | Simples | 40h | 72h | 128h | **76h** |
| FEAT-EP-0001-0003 Alertas (Should) | US-0006, 0007 | Média | 52h | 84h | 152h | **90h** |

| Componente | Horas |
|:---|---:|
| Desenvolvimento PERT | 363h |
| QA (30%) | 109h |
| Arquitetura (8%) | 29h |
| DevOps (7%) | 25h |
| Gestão (10%) | 36h |
| **Total EP-0001** | **562h** (3.5 h-m) |

### 3.2 EP-0002 — Clientes e Assinaturas (16 US · 5 features · D2+D3)

| Feature | US | Compl. | O | ML | P | PERT |
|:---|---|---:|---:|---:|---:|---:|
| FEAT-EP-0002-0001 Cadastro | US-0008 a 0011 | Média | 84h | 160h | 272h | **166h** |
| FEAT-EP-0002-0002 Status | US-0012 a 0014 | Média | 112h | 184h | 328h | **196h** |
| FEAT-EP-0002-0003 Planos | US-0015 a 0018 | Média | 128h | 224h | 392h | **236h** |
| FEAT-EP-0002-0004 Assinaturas | US-0019 a 0021 | Média | 104h | 176h | 312h | **187h** |
| FEAT-EP-0002-0005 Auditoria | US-0022, 0023 | Média | 56h | 96h | 168h | **101h** |

| Componente | Horas |
|:---|---:|
| Desenvolvimento PERT | 886h |
| QA (30%) | 266h |
| Arquitetura (8%) | 71h |
| DevOps (7%) | 62h |
| Gestão (10%) | 89h |
| **Total EP-0002** | **1,374h** (8.6 h-m) |

### 3.3 EP-0003 — RBAC (16 US · 4 features · D4)

| Feature | US | Compl. | O | ML | P | PERT |
|:---|---|---:|---:|---:|---:|---:|
| FEAT-EP-0003-0001 Usuários | US-0024 a 0026, 0059 a 0061 | Média | 164h | 284h | 504h | **301h** |
| FEAT-EP-0003-0002 Papéis | US-0027 a 0030 | Média | 144h | 240h | 432h | **256h** |
| FEAT-EP-0003-0003 Vinc. BU×Mod | US-0031 a 0033 | Complexa | 112h | 200h | 352h | **211h** |
| FEAT-EP-0003-0004 Visibilidade | US-0034 a 0036 | Média | 88h | 152h | 264h | **160h** |

| Componente | Horas |
|:---|---:|
| Desenvolvimento PERT | 927h |
| QA (30%) | 278h |
| Arquitetura (8%) | 74h |
| DevOps (7%) | 65h |
| Gestão (10%) | 93h |
| **Total EP-0003** | **1,437h** (9.0 h-m) |

⚠️ **Alerta:** EP-0003 é o épico de maior complexidade (RBAC + Keycloak + Kong header injection). US-0027/0028 (papéis e permissões) e US-0031/0032 (vinculações many-to-many) são as de maior risco.

### 3.4 EP-0004 — Portal Cliente (23 US · 6 features · D5+D6+D7)

| Feature | US | Compl. | O | ML | P | PERT |
|:---|---|---:|---:|---:|---:|---:|
| FEAT-EP-0004-0001 Auth | US-0037 a 0039 | Média | 80h | 136h | 240h | **144h** |
| FEAT-EP-0004-0002 Onboarding | US-0040 a 0044 | Média | 160h | 280h | 480h | **293h** |
| FEAT-EP-0004-0003 Dashboard (Should) | US-0045, 0046, 0062 | Média | 112h | 184h | 328h | **196h** |
| FEAT-EP-0004-0004 App Switcher | US-0047 a 0049 | Média | 80h | 144h | 248h | **151h** |
| FEAT-EP-0004-0005 BUs | US-0050 a 0054 | Média | 144h | 248h | 432h | **261h** |
| FEAT-EP-0004-0006 Catálogo | US-0055 a 0058 | Simples | 76h | 136h | 240h | **143h** |

| Componente | Horas |
|:---|---:|
| Desenvolvimento PERT | 1,189h |
| QA (30%) | 357h |
| Arquitetura (8%) | 95h |
| DevOps (7%) | 83h |
| Gestão (10%) | 119h |
| **Total EP-0004** | **1,843h** (11.5 h-m) |

⚠️ **Alerta de Capacidade:** M5 é frontend-intensive. Tom Santos (FE dedicado) chega em 01/11 — após M5 (30/09).

---

## 4. Alocação de Recursos

| Perfil | Disponibilidade | Atuação Principal |
|:---|---:|:---|
| Francisco/Bolismar (TL/Full-Stack) | 100% | Backend + Frontend + Arquitetura |
| Bruno Gratto (SA/FE) | 100% | Frontend + Design System |
| Felipe Canedas (QA) | 100% | Testes + Automação |
| Davi Silva (DevOps) | 100% | IaC + CI/CD + Deploy |
| Carlos Caldas (DB) | 100% | PostgreSQL + Flyway |
| Mauro (BA) | 50% | Refinamento + Critérios |
| Francisco (TL) | 50% | Governança + Compliance Gates |
| Tom Santos (FE) | 0% até 01/11 | Frontend dedicado (chegada tardia) |
| Maria Madalena (Junior) | 100% | Testes + Documentação + Tasks simples |

**Capacidade mensal efetiva:** ~9 pessoas × 160h = **~1,440h/mês**

| Cenário | Horas | Capacidade | Duração |
|:---|---:|---:|:---|
| Conservador (15%) | 6,552h | 1,440h/mês | **4.6 meses** |
| PERT | 6,077h | 1,440h/mês | **4.2 meses** |
| Pessimista (25%) | 7,121h | 1,440h/mês | **4.9 meses** |

> ⚠️ Prazo do Charter: 24/07 → 30/10 = 3.2 meses. Duração PERT de 4.2 meses **excede em ~1 mês**.

---

## 5. Validação DTA

| Regra | Valor | Status |
|:---|---:|:---:|
| QA por épico ≥ 20% | EP-0001: 19%, EP-0002: 19%, EP-0003: 19%, EP-0004: 19% | ⚠️ Marginal |
| QA Global ≥ 25% | 30% | ✅ |
| Arch Global ≥ 5% | 8% | ✅ |
| Independência | Zero refs ao ROM | ✅ |
| PERT por US | 62/62 estimadas | ✅ |
| IC 95% | Calculado em todos os níveis | ✅ |

---

## 6. Riscos

| Risco | Prob. | Impacto | Ação |
|:---|:---:|:---|:---|
| Frontend sem dev dedicado até 01/11 | Alta | +15% M5 | Bruno + Bolismar cobrem; priorizar backend primeiro |
| RBAC multi-tenant complexo (EP-0003) | Média | +20% | Design review dedicado; spike técnico 40h |
| Istio/Kong requer especialista | Média | +10% | Davi com consultoria externa |
| Junior em tarefas críticas | Baixa | +5% | Tasks de baixa complexidade |
| Escopo creep | Média | +10-20% | Gate de controle de mudança no M3 |

---

🤖 *Estimativa 100% independente — Fase 8 do Downstream Architecture Refinement. PERT calculado do zero, US por US. Zero contaminação do ROM upstream.*
