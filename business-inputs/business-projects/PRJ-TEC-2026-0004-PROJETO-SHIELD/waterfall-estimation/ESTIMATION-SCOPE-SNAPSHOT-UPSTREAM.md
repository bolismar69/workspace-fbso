# SCOPE SNAPSHOT — UPSTREAM/DISCOVERY: PRJ-TEC-2026-0004-PROJETO-SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Estimativa Vinculada** | WATERFALL-ESTIMATION-UPSTREAM-ROM.md v1.0 |
| **Data de Congelamento** | 04/08/2026 |
| **Versão** | 1.0 |
| **Modo** | UPSTREAM/DISCOVERY |

---

## SCOPE SNAPSHOT — UPSTREAM/DISCOVERY
Um Scope Snapshot (retrato ou recorte de escopo) no contexto de Upstream/Discovery é um registro ou ponto de parada que define claramente os limites, as hipóteses e os objetivos de um problema antes de o time iniciar o desenvolvimento prático. Ele serve para alinhar o que será investigado, testado e construído.

### Conceitos Principais
- **Scope Snapshot (Recorte de Escopo):** Uma foto do momento que sintetiza o entendimento atual da dor do usuário, as fronteiras da solução proposta e o que está dentro ou fora do projeto.
- **Upstream (Rio Acima):** A fase inicial do fluxo de trabalho onde ideias, dados e hipóteses são triados, avaliados e amadurecidos antes de virarem tarefas de execução.
- **Discovery (Descoberta):** O processo de pesquisa, entrevistas e testes rápidos para validar se o problema mapeado no upstream realmente faz sentido e gera valor.

### Para que Serve
- **Reduzir Incertezas:** Evita que a equipe gaste tempo e dinheiro construindo produtos desalinhados das reais necessidades do cliente.
- **Alinhamento Coletivo:** Garante que produto, design, engenharia e negócio compartilhem a mesma visão sobre o que será validado.
- **Critério de Saída:** Funciona como um marco para indicar quando a ideia está madura o suficiente para descer para o fluxo de desenvolvimento (downstream)

---

### 1. Itens de Escopo Estimados

| ID | Componente | Fonte WATERFALL | Seção | Versão Doc | Status |
|----|-----------|----------------|-------|-----------|--------|
| C01 | Cloudflare Edge | 06-HLD | §1 (System Context), §2 (Container Diagram) | v2.0 | ✅ Estimado |
| C02 | SPA Frontend | 06-HLD | §1, §2 | v2.0 | ✅ Estimado |
| C03 | Kong API Gateway | 06-HLD | §1, §2, §4 (Integration Topology) | v2.0 | ✅ Estimado |
| C04 | ms-shield-identity-auth (BFF) | 06-HLD | §1, §2, §4 | v2.0 | ✅ Estimado |
| C05 | Keycloak Multi-Realm | 06-HLD | §1, §2 | v2.0 | ✅ Estimado |
| C06 | PostgreSQL RLS | 06-HLD | §1, §2 | v2.0 | ✅ Estimado |
| C07 | Redis | 06-HLD | §1, §2, §4 | v2.0 | ✅ Estimado |
| C08 | Istio Service Mesh | 06-HLD §1; 05-SAD §1 | HLD: System Context; SAD: Architectural Overview | v2.0 | ✅ Estimado |
| C09 | DOKS Infrastructure | 06-HLD | §1, §3 (Technology Stack) | v2.0 | ✅ Estimado |
| C10 | Observability Stack | 06-HLD §3; 05-SAD §1 | HLD: Tech Stack; SAD: Overview | v2.0 | ✅ Estimado |
| C11 | CI/CD Pipeline | 06-HLD | §3 (Technology Stack) | v2.0 | ✅ Estimado |
| C12 | Integração Microserviços | 06-HLD | §1, §2, §4 | v2.0 | ✅ Estimado |

**Total de componentes estimados:** 12

---

### 2. Exclusões Explícitas (NÃO Estimado)

| Item | Motivo da Exclusão | Fonte da Decisão |
|------|-------------------|-----------------|
| Desenvolvimento dos microserviços de negócio | Fora do escopo do Shield — apenas integração | 01-Charter §3.1 |
| Migração de dados legados | Não previsto no escopo do projeto | 01-Charter §3 |
| Portal administrativo (backoffice) | Não listado nas entregas do Charter | 01-Charter §3.2 |
| Google for Education / Microsoft 365 / GOV.BR | Visão de longo prazo — fora do escopo atual | 01-Charter §2.1 |
| Conformidade LGPD (jurisdição) | Depende de parecer do Jurídico | 02-BRD BC-04 |
| Onboarding self-service (4h) | Meta de longo prazo — não no MVP | 01-Charter §2.1 |
| SSO unificado entre produtos | Meta de longo prazo — não no MVP | 01-Charter §2.1 |

---

### 3. Matriz de Rastreabilidade (Escopo × Documento Fonte)

| Escopo | 01-Charter | 02-BRD | 05-SAD | 06-HLD |
|--------|-----------|--------|--------|--------|
| C01 Cloudflare | §3.1 (Infra Segurança) | BC-03 (Proteção Borda) | — | §1, §2 |
| C02 SPA Frontend | §3.1 (Portal Acesso) | REQ-01, REQ-10 | — | §1, §2 |
| C03 Kong Gateway | §3.1 (Infra Segurança) | REQ-03 (Proteção Credenciais) | §1 (ADR-001) | §1, §2, §4 |
| C04 BFF (Shield) | §3.1 (Plataforma Identidade) | REQ-01, REQ-04, REQ-05 | §1 (ADR-001) | §1, §2, §4 |
| C05 Keycloak | §3.1 (Ambiente Isolado) | REQ-01, REQ-02, REQ-08 | §1 | §1, §2 |
| C06 PostgreSQL | §3.1 (Ambiente Isolado) | REQ-02, REQ-07 | §1 | §1, §2 |
| C07 Redis | §3.1 (Infra Segurança) | REQ-05 (Latência <15ms) | §1 | §1, §2, §4 |
| C08 Istio | §3.1 (Infra Segurança) | REQ-02 (Isolamento) | §1 | §1 |
| C09 DOKS | §3.1 (Infra Segurança) | BC-02 (Provedor Exclusivo) | §1 | §1, §3 |
| C10 Observabilidade | — | REQ-07 (Auditoria) | §1 | §3 |
| C11 CI/CD | — | — | §1 (ADR-001) | §3 |
| C12 Integração MS | §3.1 (Portal Acesso) | REQ-04, REQ-10 | §1 | §1, §2, §4 |

---

### 4. Versões dos Documentos Fonte

| Documento | Versão | Data | Status |
|-----------|--------|------|--------|
| 01-PROJECT-CHARTER | v2.0 | 03/08/2026 | COMPLIANCE |
| 02-BRD | v2.0 | 03/08/2026 | COMPLIANCE |
| 05-SAD | v2.0 | 03/08/2026 | COMPLIANCE |
| 06-HLD | v2.0 | 03/08/2026 | COMPLIANCE |

---

### 5. Premissas de Escopo

| Premissa | Impacto se Inválida |
|----------|---------------------|
| O escopo do Shield se limita à plataforma de identidade — não inclui desenvolvimento dos microserviços de negócio | +Escopo significativo (centenas de horas) |
| Apenas 1 domínio/cliente no MVP; multi-domínio é configuração adicional | Multiplicar esforço Cloudflare e Keycloak por N domínios |
| Stack 100% dentro da baseline corporativa FBSO — sem tecnologias exóticas | Risco de incompatibilidade e retrabalho |
| Equipe de 4 pessoas (2 seniores, 1 pleno, 1 júnior) dedicadas 100% | Se shared, prazo estica proporcionalmente |
| GraalVM Native compila sem issues de reflection/serialization | +80h no BFF se fallback para JVM |
| Microserviços de negócio já aceitam header JWT padrão FBSO | +40h se adaptação de contratos necessária |
