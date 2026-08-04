# DISCOVERY-LEVEL-ROM-ESTIMATE.md
## Fase 11 — Bloco D: Estimativa & ROM

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | DISCOVERY-LEVEL-ROM-ESTIMATE-v1.2 |
| **Versão** | 1.2 — ROM ±50% (Discovery-Level) — Semanas + Horas |
| **Data** | 02 de agosto de 2026 |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 02/08/2026 |

**Documentos Referenciados:**
- F1-F10: Todas as fases do Upstream Architecture Discovery
- [`DISCOVERY-LEVEL-SPECS.md`](DISCOVERY-LEVEL-SPECS.md) — Consolidação Técnica (F10)

> **Premissa de conversão:** 1 dia = 8 horas · 1 semana = 40 horas · Duração considera 5 dias úteis/semana.

---

## 1. Estimativa ROM Consolidada (±50%)

### 1.1 Esforço de Discovery (Sprint 0 — Análise)

| Disciplina | Atividades | Range/Semanas | Range/Horas | Responsável |
|------------|-----------|:-------------------:|:-----------------:|-------------|
| **Arquitetura** | C4 L1+L2, ADRs, modularização | {0.4-0.4}sem | {16-16}h | Alfredo Salomao |
| **Arquitetura** | Validação GraalVM Native Image POC | {0.2-0.2}sem | {8-8}h | Francisco Oliveira |
| **Segurança** | Kong↔Keycloak config, threat model, compliance docs | {0.6-0.8}sem | {24-32}h | Daniel Bruno Castro |
| **Dados** | Modelagem ER, RLS policies, cache strategy | {0.8-1.0}sem | {32-40}h | William Alves |
| **DevOps/SRE** | CI/CD, Terraform, Ansible, K8s, Istio, observabilidade | {2.2-2.6}sem | {88-104}h | Lucas Silva Neto |
| **Testes** | Estratégia, E2E cenários, quality gates | {0.8-1.0}sem | {32-40}h | Valeria Lucanete |
| **Infra/Cloud** | Provisioning, Cloudflare, DR, custos | {0.9-1.1}sem | {36-44}h | Lucas Silva Neto |
| **Integração** | Contratos API, documentação | {0.2-0.2}sem | {8-8}h | Bolismar Oliveira |
| **Documentação** | Todos os artefatos F1-F11 | {0.4-0.4}sem | {16-16}h | Time |
| **Total Discovery** | 7 profissionais × 1 sprint | **{0.9-1.1}sem** | **{260-308}h** | Time Discovery |

> Duração: 5-7 dias corridos com 7 profissionais em paralelo (1 sprint).

### 1.2 Esforço de Implementação — Por Épico

| Épico | Func. | Complexidade | Time Alocado | Range/Semanas | Range/Horas |
|-------|:-----:|-------------|-------------|:-------------------:|:-----------------:|
| **EP-0001** — Portal Admin Interno | 3 | Média | 2 BE + 1 FE + QA (3.5 pess.) | {3-4}sem | {420-560}h |
| **EP-0002** — Clientes e Assinaturas | 5 | Alta | 2 BE + 1 FE + QA (3.5 pess.) | {5-7}sem | {700-980}h |
| **EP-0003** — Governança e Permissões | 4 | Alta | 2 BE + 1 FE + QA (3.5 pess.) | {4-6}sem | {560-840}h |
| **EP-0004a** — Portal do Cliente e Onboarding | 3 | Média | 1 BE + 1 FE + QA (2.5 pess.) | {3-4}sem | {300-400}h |
| **EP-0004b** — Unidades de Negócio e Catálogo | 2 | Média | 2 BE + 1 FE + QA (3.5 pess.) | {2-3}sem | {280-420}h |
| **Subtotal Implementação** | 17 | — | — | **{17-24}sem** | **{2.260-3.200}h** |

### 1.3 Esforço de Infraestrutura e DevOps (Paralelo à Implementação)

| Atividade | Range/Semanas | Range/Horas | Responsável |
|-----------|:-------------------:|:-----------------:|-------------|
| Terraform (DOKS, DBs, Redis, Spaces) | {2-3}sem | {80-120}h | DevOps |
| Kong + Keycloak configuração | {2-3}sem | {80-120}h | IAM Specialist |
| Pipeline CI/CD (GitHub Actions) | {1-2}sem | {40-80}h | DevOps |
| Observabilidade stack (Prometheus, Loki, Jaeger, Grafana, Elastic) | {2-3}sem | {80-120}h | DevOps |
| Cloudflare configuração (DNS, WAF, SSL) | {1-1}sem | {40-40}h | DevOps |
| Istio + Keda + Karpenter | {2-3}sem | {80-120}h | DevOps |
| **Subtotal Infra/DevOps** | **{10-15}sem** | **{400-600}h** | DevOps + IAM |

> Executado por 2 pessoas (DevOps + IAM Specialist) em paralelo aos épicos de implementação.

### 1.4 Esforço de Testes e Qualidade (Paralelo à Implementação)

| Atividade | Range/Semanas | Range/Horas | Responsável |
|-----------|:-------------------:|:-----------------:|-------------|
| Testes unitários e integração (por épico) | Embutido nos épicos | Embutido nos épicos | Devs + QA |
| E2E — 6 fluxos críticos (Playwright) | {2-3}sem | {80-120}h | QA |
| Testes de performance (k6) | {1-2}sem | {40-80}h | QA |
| Penetration test (pré-lançamento) | {1-1}sem | {40-40}h | QA + IAM |
| **Subtotal QA** | **{4-6}sem** | **{160-240}h** | QA |

> Executado por 1 pessoa (QA Engineer) em paralelo, com pico de esforço no pré-lançamento.

---

## 2. Estimativa ROM Total

### 2.1 Consolidação por Bloco

| Bloco | Range/Semanas | Range/Horas | % do Total | Paralelizável? |
|-------|:-------------------:|:-----------------:|:----------:|:--------------:|
| **Discovery (Sprint 0)** | {1-1}sem | {260-308}h | 8% | Não — precede tudo |
| **Implementação (Épicos)** | {17-24}sem | {2.260-3.200}h | 71% | Parcial — épicos sequenciais |
| **Infra/DevOps** | {10-15}sem | {400-600}h | 13% | Sim — paralelo aos épicos |
| **QA (E2E + Perf + Pentest)** | {4-6}sem | {160-240}h | 5% | Sim — paralelo aos épicos |
| **Buffer / Imprevistos** | {2-3}sem | {100-160}h | 3% | Não |
| **Total Geral** | **{22-26}sem** | **{3.180-4.508}h** | 100% | — |

### 2.2 Cenários ROM (±50%)

| Cenário | Range/Semanas | Range/Horas | Time | Custo Infra (6 meses) |
|---------|:-------------------:|:-----------------:|------|:---------------------:|
| **Otimista (-50%)** | {11-13}sem | {1.590-2.254}h | 4 pessoas | ~$2,400 |
| **Provável (Base)** | **{22-26}sem** | **{3.180-4.508}h** | 5-6 pessoas | ~$5,800 |
| **Pessimista (+50%)** | {33-39}sem | {4.770-6.762}h | 7 pessoas | ~$8,700 |

> **Nota:** Custos de time dependem do modelo de contratação (CLT, PJ, outsourcing) e não estão precificados aqui. Infraestrutura considera 3 ambientes (Dev + Staging + Prod) pelo período do projeto.

### 2.3 Caminho Crítico (Cenário Provável)

| Fase | Range/Semanas | Range/Horas | Time Alocado |
|------|:-------------------:|:-----------------:|-------------|
| Sprint 0 — Discovery | {1-1}sem | {260-308}h | 7 profissionais |
| Setup Infra/DevOps | {3-4}sem | {400-600}h | 2 (DevOps + IAM) |
| EP-0001 — Portal Admin Interno | {3-4}sem | {420-560}h | 3 (2 BE + 1 FE) |
| EP-0002 — Clientes e Assinaturas | {5-7}sem | {700-980}h | 3 (2 BE + 1 FE) |
| EP-0003 — Governança e Permissões | {4-6}sem | {560-840}h | 3 (2 BE + 1 FE) |
| EP-0004a — Portal do Cliente | {3-4}sem | {300-400}h | 2 (1 BE + 1 FE) |
| EP-0004b — BUs e Catálogo | {2-3}sem | {280-420}h | 3 (2 BE + 1 FE) |
| QA (E2E + Perf + Pentest) | Paralelo + {2-2}sem pós | {160-240}h | 1 (QA) |
| Buffer / Imprevistos | {2-3}sem | {100-160}h | — |
| **Total Caminho Crítico** | **{22-26}sem** | **{3.180-4.508}h** | 5-6 pessoas |

### 2.4 Distribuição do Esforço por Disciplina

| Disciplina | Range/Semanas | Range/Horas | % do Total |
|------------|:-------------------:|:-----------------:|:----------:|
| **Desenvolvimento Backend** | — | {1.272-1.803}h | 40% |
| **Desenvolvimento Frontend** | — | {636-902}h | 20% |
| **DevOps/Infra/SRE** | — | {477-676}h | 15% |
| **QA/Testes** | — | {318-451}h | 10% |
| **Segurança/IAM** | — | {254-361}h | 8% |
| **Dados/DB** | — | {159-225}h | 5% |
| **Documentação** | — | {64-90}h | 2% |
| **Total** | **{22-26}sem** | **{3.180-4.508}h** | 100% |

---

## 3. Premissas da Estimativa

| # | Premissa | Impacto se inválida |
|---|----------|---------------------|
| P1 | Time sênior (todos ★★★) com 100% de dedicação | 1.5-2× esforço → {33-52}sem / {4.770-9.016}h |
| P2 | GraalVM Native Image viável (POC Sprint 0) | +{2-4}sem / +{80-160}h para build JVM tradicional |
| P3 | Kong↔Keycloak Service-ID/Token-ID configurável em {2-3}sem | +{2-4}sem / +{80-160}h no setup inicial |
| P4 | DigitalOcean supre todas as necessidades de infra | +{4-6}sem / +{160-240}h para migração de IaC |
| P5 | Escopo limitado ao MVP (13 funcionalidades) | Cada funcionalidade extra: +{1-3}sem / +{40-120}h |
| P6 | Sem dependências externas bloqueantes | Cada integração externa: +{1-4}sem / +{40-160}h |
| P7 | 1 sprint de Discovery suficiente para análise | +{1-2}sem / +{260-308}h (sprints extras de análise) |

---

## 4. Riscos da Estimativa

| ID | Risco | Impacto em Semanas | Impacto em Horas | Mitigação |
|----|-------|:-----------------:|:----------------:|-----------|
| RE1 | Scope creep — funcionalidades além do MVP | +{4-10}sem | +{636-1.803}h | PRD define MVP claro; change request formal |
| RE2 | Complexidade Kong↔Keycloak subestimada | +{2-4}sem | +{80-160}h | IAM Specialist dedicado; POC nas primeiras 48h |
| RE3 | Monólito Modular — refactoring para extração | +{3-6}sem | +{120-240}h | Boundaries de domínio com contratos explícitos |
| RE4 | Curva de aprendizado GraalVM Native Image | +{1-3}sem | +{40-120}h | POC Sprint 0; fallback JVM HotSpot |
| RE5 | Time não dedicado — interrupções e context switching | +{7-13}sem | +{954-2.254}h | Premissa: time 100% dedicado |

---

## 5. Recomendação Técnica para o Comitê de Governança

### Parecer do Time de Arquitetura

Com base na análise de viabilidade técnica conduzida nas 11 fases do Upstream Architecture Discovery, o time de arquitetura emite o seguinte parecer:

**✅ O projeto é tecnicamente viável.**

**Fundamentação:**

1. **Stack madura e corporativa:** Todas as tecnologias propostas seguem os padrões corporativos FBSO (DigitalOcean, Kong↔Keycloak, PostgreSQL RLS, Spring Boot, React/Next.js, K8s/Istio). Não há dependência de tecnologias experimentais ou não validadas.

2. **Arquitetura adequada ao estágio:** O modelo de monólito modular com extração futura é apropriado para a fase Core. Evita a complexidade prematura de microserviços enquanto preserva o caminho de evolução.

3. **Riscos gerenciáveis:** Os 5 principais riscos técnicos têm mitigações documentadas. O risco mais crítico (cross-tenant data access) é mitigado por RLS no PostgreSQL + testes automatizados de isolamento.

4. **Time qualificado:** 7 profissionais seniores com as competências necessárias. O Sprint 0 de Discovery (1 semana) é suficiente para a análise de viabilidade.

5. **Estimativa realista:** ROM de **{22-26}sem / {3.180-4.508}h** para implementação do MVP com time de 5-6 pessoas, com faixa de variação de ±50% conforme premissas documentadas.

**Recomendação:** **GO ✅** — prosseguir para estabelecimento do SQUAD e início da implementação.

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: ROM ±50% em semanas | Tech Lead / Solution Architect |
| 1.1 | 02/08/2026 | Conversão para horas com ranges {min-max} | Tech Lead / Solution Architect |
| 1.2 | 02/08/2026 | Adicionada coluna Range/Semanas lado a lado com Range/Horas em todas as tabelas; premissas e riscos com impacto dual (semanas + horas) | Tech Lead / Solution Architect |

---

🤖 *Upstream Architecture Discovery — Fase 11. Documento final do Discovery — ROM Estimate para o Comitê de Governança.*
