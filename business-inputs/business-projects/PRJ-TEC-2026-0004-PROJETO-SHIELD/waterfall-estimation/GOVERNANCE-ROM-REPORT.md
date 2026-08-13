# RELATÓRIO DE GOVERNANÇA — GO/NO-GO: PRJ-TEC-2026-0004-PROJETO-SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Estimativa Base** | WATERFALL-ESTIMATION-UPSTREAM-ROM.md v1.0 |
| **Data** | 04/08/2026 |
| **Versão** | 1.0 |
| **Destinatário** | Comitê de Governança — FBSO.ORG |

---

### 1. Sumário Executivo

O Projeto SHIELD entrega a **plataforma de identidade centralizada da FBSO.ORG** — porta de entrada única para todos os produtos do ecossistema. Com 12 componentes arquiteturais e stack 100% alinhada à baseline corporativa, o esforço estimado é de **1.072 horas (R$ 123 mil)** com faixa ROM de 536h a 1.608h. A duração provável de **9 semanas** excede o prazo de 6 semanas do BRD, exigindo decisão sobre escopo ou capacidade. **Recomendação: GO condicionado** à renegociação do prazo.

---

### 2. Escopo e Premissas

**Escopo coberto pela estimativa (12 componentes):**
- Cloudflare Edge (WAF, DNS, proxy)
- SPA Frontend (portal de acesso)
- Kong API Gateway (roteamento, validação JWT)
- ms-shield-identity-auth BFF (Quarkus Native, validação de sessão)
- Keycloak Multi-Realm (OIDC, PKCE, temas)
- PostgreSQL RLS (isolamento multi-tenant)
- Redis (session store, cache)
- Istio Service Mesh (mTLS)
- DOKS Infrastructure (Kubernetes, Terraform)
- Observability Stack (Prometheus, Grafana, OpenTelemetry)
- CI/CD Pipeline (GitHub Actions, Argo CD)
- Integração com Microserviços de Negócio

**Premissas críticas (se alteradas, invalidam a estimativa):**
1. Stack dentro da baseline corporativa FBSO — sem tecnologias não aprovadas
2. GraalVM Native compila sem issues de reflection/serialization
3. Equipe de 4 pessoas dedicadas 100% ao projeto
4. Microserviços de negócio já aceitam header JWT padrão FBSO
5. Apenas 1 domínio/cliente no MVP

---

### 3. Estimativa Financeira (ROM ±50%)

| Cenário | Horas | Custo Estimado (R$) |
|---------|-------|---------------------|
| Mínimo (−50%) | 536 h | R$ 61.640 |
| **Provável** | **1.072 h** | **R$ 123.280** |
| Máximo (+50%) | 1.608 h | R$ 184.920 |

**Composição do custo (Provável):**

| Dimensão | Horas | % do Total |
|----------|-------|-----------|
| Desenvolvimento | 496 h | 46,3% |
| QA | 188 h | 17,5% |
| Arquitetura | 104 h | 9,7% |
| DevOps/SRE | 200 h | 18,7% |
| Gestão | 84 h | 7,8% |

---

### 4. Timeline Macro

| Fase | Duração Estimada | Esforço |
|------|-----------------|---------|
| Infraestrutura e Ambiente | 2 semanas | C01, C08, C09, C11 |
| Motor de Identidade (Keycloak) | 3 semanas | C05 |
| Portal de Acesso (BFF + SPA) | 4 semanas | C02, C03, C04, C07 |
| Isolamento de Dados (RLS) | 2 semanas | C06 |
| Observabilidade | 1,5 semanas | C10 |
| Integração e Testes | 3 semanas | C12, QA transversal |
| Deploy e Go-Live | 1 semana | C11, C09 |
| **Total (com sobreposição)** | **~9 semanas** | **1.072 h** |

**Marcos (do Charter):**

| Marco | Data Prevista |
|-------|--------------|
| Kickoff | Imediato (04/08/2026) |
| Infraestrutura provisionada | Semana 2 |
| Motor de Identidade funcional | Semana 5 |
| Portal de Acesso funcional | Semana 7 |
| Go-Live | Semana 9 (~06/10/2026) |

---

### 5. Riscos e Mitigadores

| Risco | Severidade | Mitigador |
|-------|-----------|-----------|
| GraalVM Native — incompatibilidade de reflection/serialization | 🔴 Alta | Prova de conceito na Semana 1; fallback para JVM se necessário (+80h) |
| Prazo BRD (6 sem) vs estimativa (9 sem) | 🔴 Alta | Negociar prazo OU cortar Istio/Observabilidade para fase 2 (−132h) |
| Keycloak Multi-Realm — curva de aprendizado | 🟡 Média | Alocar sênior dedicado; usar realm template padrão FBSO |
| Complexidade RLS — políticas mais complexas que previsto | 🟡 Média | Projetar RLS policies na Semana 1; validar com DBA sênior |
| Integração com MS legados — contratos não padronizados | 🟢 Baixa | Mapear contratos na Semana 1; adaptar header JWT se necessário |
| Dependência SaaS (Cloudflare, DOKS) | 🟢 Baixa | Configuração declarativa; monitorar changelogs |

---

### 6. Recomendação

**Decisão recomendada:** 🟢 **GO condicionado**

**Justificativa:**
O projeto tem arquitetura madura (HLD consolidado, ADRs documentados), stack 100% alinhada à baseline corporativa, e estimativa com validação DTA robusta (QA 38%, Arch 10%). O único ponto de atenção é o prazo: 9 semanas vs 6 semanas do BRD. Recomenda-se aprovar com a condição de renegociar o prazo para 9 semanas OU reduzir escopo (cortar Istio e Observabilidade avançada para fase 2, economizando ~132h e ~1 semana).

**Condições para GO:**
1. Renegociar prazo com stakeholders: 6 → 9 semanas
2. Se prazo de 6 semanas for inegociável: cortar C08 (Istio) e C10 (Observabilidade) para fase 2
3. Prova de conceito GraalVM Native na Semana 1 para mitigar risco de incompatibilidade

---

### 7. Decisão do Comitê

| Campo | Preenchimento |
|-------|--------------|
| **Decisão** | [_] GO — Aprovado  [_] NO-GO — Rejeitado  [_] HOLD — Pendente |
| **Condições** | |
| **Responsável** | |
| **Data** | __/__/____ |
| **Assinatura** | _______________________ |
