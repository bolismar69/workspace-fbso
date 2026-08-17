# ESTIMATIVA ROM UPSTREAM/DISCOVERY (±50%): PRJ-TEC-2026-0004-PROJETO-SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 02-BRD, 05-SAD, 06-HLD |
| **Data de Elaboração** | 04/08/2026 |
| **Versão** | 1.0 |
| **Modo** | UPSTREAM/DISCOVERY — ROM ±50% |
| **Metodologia** | Bottom-Up por Componente Arquitetural × 5 Dimensões DTA |
| **Time** | 2 Seniores, 1 Pleno, 1 Júnior — Disponibilidade: 3 meses |

---

## ROM UPSTREAM/DISCOVERY (±50%)

Uma estimativa ROM (Rough Order of Magnitude ou Ordem de Grandeza Aproximada) no contexto de Upstream ou Product Discovery é uma projeção inicial e de alto nível de custo, tempo ou esforço feita no início de uma ideia, quando ainda há muitas incertezas e o escopo não está detalhado. Ela serve para validar a viabilidade preliminar do projeto antes de investir em pesquisas aprofundadas ou desenvolvimento

### Características da Estimativa ROM no Discovery
- **Fase inicial:** Ocorre antes de qualquer especificação técnica profunda ou planejamento fechado
- **Alta margem de variação:** Possui uma faixa de precisão ampla (frequentemente de -25% a +75%), pois baseia-se em premissas e não em fatos consolidados.
- **Foco em decisão rápida:** Ajuda a diretoria ou o time de produtos a decidir se vale a pena avançar com o Discovery daquela oportunidade ou descartá-la imediatamente.

### Como Funciona no Ciclo de Produto
- **Triagem de ideias:** No Upstream, várias propostas entram como hipóteses ou desejos de melhoria.
- **Aproximação de custo:** Aplica-se a ROM para mensurar se a ideia custará "dias, semanas ou meses" de trabalho bruto.
- **Refinamento progressivo:** Conforme o Product Discovery avança e as dores do usuário são validadas, a estimativa ROM deixa de existir e é substituída por estimativas reais e refinadas na etapa de Downstream.

---

### 1. Escopo Estimado

**Componentes do HLD considerados nesta estimativa:**

| ID | Componente/Container | Descrição | Fonte (HLD) |
|----|---------------------|-----------|-------------|
| C01 | Cloudflare Edge | WAF, DNS, SSL termination, proxy de API, header X-Tenant-Host | HLD §1, §2 |
| C02 | SPA Frontend | Aplicação estática HTML/JS/CSS — formulários de login, zero secrets | HLD §1, §2 |
| C03 | Kong API Gateway | Routing `/auth/*` e `/api/*`, validação JWT, rate limiting, plugin de autenticação | HLD §1, §2, §4 |
| C04 | ms-shield-identity-auth (BFF) | Quarkus Native — validação de sessão, injeção JWT, mapeamento Host→Realm | HLD §1, §2, §4 |
| C05 | Keycloak Multi-Realm | Provisionamento de realms, fluxos OIDC/PKCE, temas visuais, integração Kong | HLD §1, §2 |
| C06 | PostgreSQL RLS | Schema `fbso_portal`, RLS multi-tenant, pool de conexões, SET LOCAL tenant | HLD §1, §2 |
| C07 | Redis | Session store (SHIELD_SESSION), cache Host→Realm, cache JWKS | HLD §1, §2, §4 |
| C08 | Istio Service Mesh | mTLS entre serviços, traffic control, sidecar injection | HLD §1; SAD §1 |
| C09 | DOKS Infrastructure | Cluster Kubernetes 1.30, networking, node pools, Terraform provisioning | HLD §1, §3 |
| C10 | Observability Stack | Prometheus + Grafana + OpenTelemetry + Loki — métricas, logs, tracing | HLD §3; SAD §1 |
| C11 | CI/CD Pipeline | GitHub Actions (build, SAST, secrets) + Argo CD (GitOps deploy) | HLD §3 |
| C12 | Integração Microserviços | Contratos de API, headers JWT, testes de integração com ms-escolas-core, ms-reforma-core, ms-saas-core | HLD §1, §2, §4 |

**Exclusões explícitas:**
- Desenvolvimento dos microserviços de negócio (ms-escolas-core, ms-reforma-core, ms-saas-core) — apenas integração com o Shield
- Migração de dados de sistemas legados
- Desenvolvimento de portal administrativo (backoffice)
- Integração com Google for Education, Microsoft 365, GOV.BR (visão de longo prazo — Charter §2.1)
- Conformidade LGPD com jurisdição de dados (depende do Jurídico — BRD BC-04)

---

### 2. Matriz de Componentes × Dimensões (Horas)

| ID | Componente | Dev (h) | QA (h) | Arch (h) | DevOps (h) | Gestão (h) | Total (h) |
|----|-----------|---------|--------|----------|------------|------------|-----------|
| C01 | Cloudflare Edge | 16 | 8 | 4 | 8 | 4 | **40** |
| C02 | SPA Frontend | 40 | 16 | 4 | 8 | 8 | **76** |
| C03 | Kong API Gateway | 40 | 16 | 8 | 16 | 8 | **88** |
| C04 | ms-shield-identity-auth (BFF) | 120 | 40 | 16 | 16 | 16 | **208** |
| C05 | Keycloak Multi-Realm | 80 | 32 | 16 | 24 | 12 | **164** |
| C06 | PostgreSQL RLS | 48 | 16 | 16 | 8 | 8 | **96** |
| C07 | Redis | 32 | 12 | 8 | 8 | 4 | **64** |
| C08 | Istio Service Mesh | 24 | 8 | 8 | 24 | 4 | **68** |
| C09 | DOKS Infrastructure | 16 | 8 | 8 | 32 | 4 | **68** |
| C10 | Observability Stack | 24 | 8 | 4 | 24 | 4 | **64** |
| C11 | CI/CD Pipeline | 16 | 8 | 4 | 24 | 4 | **56** |
| C12 | Integração Microserviços | 40 | 16 | 8 | 8 | 8 | **80** |
| **TOTAL** | | **496** | **188** | **104** | **200** | **84** | **1.072** |

---

### 3. ROM Consolidado

#### 3.1 Horas Totais

| Cenário | Fator | Horas |
|---------|-------|-------|
| **Mínimo (ROM min)** | ×0.50 | 536 h |
| **Provável** | ×1.00 | **1.072 h** |
| **Máximo (ROM max)** | ×1.50 | 1.608 h |

#### 3.2 Faixa de Confiança

```
ROM = 1.072h × (1 ± 0.50)
FAIXA: [536h — 1.608h]
```

#### 3.3 Conversão Financeira

| Perfil | Horas Estimadas | Taxa Horária (R$) | Custo (R$) |
|--------|----------------|-------------------|------------|
| Sênior (×2) | 536 | R$ 150,00 | R$ 80.400,00 |
| Pleno (×1) | 268 | R$ 100,00 | R$ 26.800,00 |
| Júnior (×1) | 268 | R$ 60,00 | R$ 16.080,00 |
| **Custo Total (Provável)** | **1.072** | **R$ 115,00 (médio)** | **R$ 123.280,00** |

**Premissa de alocação:** 50% sênior, 25% pleno, 25% júnior — reflete perfil de projeto com alta complexidade arquitetural (BFF nativo, mTLS, RLS).

#### 3.4 Validação DTA Interna

| Métrica | Valor | Limite | Status |
|---------|-------|--------|--------|
| QA / Dev | 188 / 496 = **37,9%** | ≥ 25% | ✅ |
| Arch / Total | 104 / 1.072 = **9,7%** | ≥ 5% | ✅ |

---

### 4. Premissas por Componente

| ID | Componente | Premissa | Impacto se inválida |
|----|-----------|----------|---------------------|
| C01 | Cloudflare | Configuração de WAF e DNS é declarativa; 1 domínio por cliente | Multiplicar esforço por N domínios |
| C02 | SPA Frontend | Framework JS padrão FBSO; sem lógica de negócio no frontend | +Complexidade se houver estado local |
| C03 | Kong | Plugin de autenticação usa Service-ID/Token-ID padrão FBSO; sem customização de lua | +40h se lua customizado |
| C04 | BFF | GraalVM Native compila sem issues; libs Quarkus são compatible | +80h se fallback para JVM |
| C05 | Keycloak | Realm template cobre 80% dos casos; temas usam estrutura padrão | +24h por customização atípica |
| C06 | PostgreSQL | RLS policies são declarativas por tenant_id; sem stored procedures complexas | +32h se lógica de isolamento complexa |
| C07 | Redis | Dados em memória apenas — sem persistência; sem replicação cross-region | +16h se necessário failover |
| C08 | Istio | mTLS usa configuração padrão Istio; sem políticas L7 complexas | +24h se regras L7 customizadas |
| C09 | DOKS | Cluster provisionado via Terraform; node pool único | +32h se multi-pool ou multi-region |
| C10 | Observabilidade | Stack padrão FBSO; sem dashboards customizados por cliente | +16h se dashboards por tenant |
| C11 | CI/CD | Pipeline padrão FBSO; Semgrep e Gitleaks já configurados | +8h se SAST customizado |
| C12 | Integração | Microserviços já expõem APIs REST; aceitam header JWT padrão | +40h se adaptação de contratos |

---

### 5. Riscos e Fatores de Ajuste

| Risco | Probabilidade | Impacto na Estimativa | Fator de Ajuste |
|-------|--------------|----------------------|-----------------|
| GraalVM Native — incompatibilidade de libs (reflection, serialization) | 🟡 Média | Retrabalho no BFF, possível fallback para JVM | +15% (+161h) |
| Complexidade de RLS — políticas de isolamento mais complexas que o previsto | 🟡 Média | Retrabalho no modelo de dados e nas queries | +10% (+107h) |
| Keycloak — curva de aprendizado (time sem exp prévia em multi-realm) | 🟡 Média | Provisionamento e temas levam mais tempo | +10% (+107h) |
| Integração com microserviços legados — contratos de API não padronizados | 🟢 Baixa | Retrabalho nos contratos de integração | +5% (+54h) |
| Prazo BRD (6 semanas) vs estimativa (~9 semanas com 4 pessoas) — pressão sobre escopo | 🔴 Alta | Corte de escopo ou horas extras; risco de qualidade | −20% escopo ou +30% prazo |
| Dependência externa — Cloudflare e DOKS são SaaS; mudanças de API | 🟢 Baixa | Adaptação a mudanças de API dos provedores | +5% (+54h) |

**Fator de ajuste ponderado:** +12% (+129h) considerando probabilidades.

---

### 6. Análise de Alocação e Prazo

| Cenário | Horas | Equipe | Duração Estimada |
|---------|-------|--------|-----------------|
| **Mínimo (ROM min)** | 536 h | 4 pessoas × 6h/dia | **~4,5 semanas** (22 dias) |
| **Provável** | 1.072 h | 4 pessoas × 6h/dia | **~9 semanas** (45 dias) |
| **Máximo (ROM max)** | 1.608 h | 4 pessoas × 6h/dia | **~13,5 semanas** (67 dias) |

**Premissas de alocação:**
- 6 horas produtivas/dia/pessoa (descontando reuniões, cerimônias, overhead)
- Equipe 100% dedicada ao projeto
- 5 dias/semana

**⚠️ Alerta de Prazo:** O BRD (BC-05) estabelece prazo máximo de **6 semanas**. A estimativa provável (9 semanas) excede esse limite em 50%. Opções:
1. **Reduzir escopo:** Cortar componentes não-críticos (ex: Observabilidade customizada, Istio avançado)
2. **Aumentar equipe:** +1 sênior reduziria prazo para ~7 semanas
3. **Aceitar ROM max como provável:** Trabalhar com buffer e negociar prazo com stakeholders

---

### 7. Recomendação para Governança

**Resumo executivo:** O Projeto SHIELD entrega a plataforma de identidade centralizada da FBSO.ORG com 12 componentes arquiteturais. A estimativa ROM aponta esforço provável de **1.072 horas** (R$ 123 mil), com faixa de **536h a 1.608h**. A duração estimada de 9 semanas (com equipe de 4) excede o prazo de 6 semanas definido no BRD — requerendo decisão sobre escopo ou capacidade. A arquitetura é robusta (QA 38%, Arch 10%) e 100% alinhada à baseline corporativa.

**Confiança da estimativa:** Média — Baseada em HLD consolidado com ADRs e stack padronizada, mas com incertezas em GraalVM Native (reflection) e complexidade de RLS multi-tenant.

**Próximo passo recomendado:** **GO condicionado** — Aprovar financiamento com a condição de renegociar prazo (6→9 semanas) OU reduzir escopo (cortar C08-Istio e C10-Observabilidade para fase 2).

**Decisão solicitada ao Comitê:** [_] GO — Aprovado  [_] NO-GO — Rejeitado  [_] HOLD — Pendente

---

### 8. Declaração de Metodologia

> **Metodologia ROM aplicada:** Bottom-Up por componente arquitetural do HLD × 5 dimensões DTA (Dev, QA, Arch, DevOps, Gestão). Fórmula ROM = Provável × (1 ± 0.50). Validação DTA interna: QA ≥ 25% Dev, Arch ≥ 5% Total. Esta estimativa é de ORDEM DE GRANDEZA (±50%) — adequada para decisão GO/NO-GO de governança, NÃO para compromisso de prazo ou orçamento. O detalhamento para compromisso será feito no modo DOWNSTREAM/REFINEMENT (PERT ±15-25%).
