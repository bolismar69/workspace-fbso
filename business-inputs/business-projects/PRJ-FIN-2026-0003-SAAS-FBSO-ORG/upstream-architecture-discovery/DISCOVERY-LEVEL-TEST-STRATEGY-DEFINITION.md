# DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md
## Fase 6 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | 02 de agosto de 2026 |
| **Autor** | QA Engineer / Test Specialist |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 02/08/2026 |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md) — PRD Discovery-Level (F1)
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)

---

## 1. Pirâmide de Testes — Visão Macro

```
        ┌───────┐
        │  E2E   │  10% — Fluxos críticos de negócio ponta a ponta
        │        │  Ferramenta: Playwright
        ├───────┤
        │  Integ │  30% — APIs REST, contratos, DB, cache, Kong↔Keycloak
        │        │  Ferramentas: JUnit + Testcontainers + Postman
        ├───────┤
        │ Unit   │  60% — Lógica de negócio, validações, regras de domínio
        │        │  Ferramentas: JUnit 5 + Mockito (backend), Jest (frontend)
        └───────┘
```

### 1.1 Distribuição por Camada

| Camada | Peso | O que cobre | Meta de Cobertura |
|--------|------|------------|-------------------|
| **Unitários** | 60% | Regras de negócio, validações, mapeamentos, utilitários | ≥ 80% (linhas) |
| **Integração** | 30% | APIs REST, repositories, RLS policies, cache, Kong rotas, Keycloak realms | ≥ 70% (cenários) |
| **E2E** | 10% | Fluxos críticos: onboarding, ativação de tenant, suspensão, upgrade de plano, RBAC | 100% fluxos críticos |
| **Performance** | Adicional | Carga em endpoints críticos, tempo de resposta, concorrência | Métricas definidas nos SLOs |
| **Segurança** | Adicional | Isolamento cross-tenant, bypass Kong, privilege escalation, SQL injection | Zero falhas críticas |

---

## 2. Estratégia por Tipo de Teste

### 2.1 Testes Unitários

| Solução | Ferramenta | Foco | Pipeline Stage |
|---------|-----------|------|---------------|
| **Backend** | JUnit 5 + Mockito | Cada módulo interno testado isoladamente. Mocks para dependências cross-module | CI (GitHub Actions) — block on failure |
| **Frontend** | Jest + React Testing Library | Componentes, hooks, utilitários. Mocks para API calls | CI (GitHub Actions) — block on failure |

**Exemplo de assertions críticas:**
- `tenantService.suspend(tenantId)`: verifica status → "suspended", audit log criado
- `rbacService.hasAccess(userId, buId, "write")`: verifica permissão granular
- `planValidator.validateModules(planId)`: verifica consistência de módulos

### 2.2 Testes de Integração

| Tipo | Ferramenta | Escopo |
|------|-----------|--------|
| **API REST** | JUnit 5 + Testcontainers (PostgreSQL + Redis) | Endpoints do backend: CRUD, validação, erros, HTTP status codes |
| **Database** | Testcontainers + PostgreSQL | RLS policies: cross-tenant queries devem retornar vazio; Soft Delete filtra deletados |
| **Cache** | Testcontainers + Redis | Cache hit/miss; invalidação após update |
| **Contratos API** | Postman/Newman | Contratos de API documentados; validação de schemas de request/response |
| **Kong Gateway** | Testes manuais + scripts curl | Rotas, rate limiting, JWT validation, header injection |

**Teste de isolamento cross-tenant (automated):**
```java
@Test
void shouldNotAccessOtherTenantData() {
    setTenantContext(TENANT_A);
    var dataA = businessUnitRepo.findByTenant(TENANT_A);
    
    setTenantContext(TENANT_B);
    var dataB = businessUnitRepo.findByTenant(TENANT_B);
    
    // Tenant B NUNCA deve ver dados do Tenant A
    assertThat(dataB).doesNotContainAnyElementsOf(dataA);
}
```

### 2.3 Testes End-to-End (E2E)

| Fluxo Crítico | Épico | Cenários | Ferramenta |
|---------------|-------|----------|-----------|
| **Onboarding de cliente** | EP-0004 | Cadastro → confirmação de dados → criação primeira BU → boas-vindas | Playwright |
| **Ativação de tenant** | EP-0002 | Admin cria tenant → gera link onboarding → tenant aparece como ativo | Playwright |
| **Suspensão de tenant** | EP-0002 | Admin suspende → usuário do tenant perde acesso em ≤ 5 min | Playwright |
| **Upgrade de plano** | EP-0002 | Tenant Básico → Core → novos módulos visíveis no App Switcher | Playwright |
| **RBAC — Operador restrito** | EP-0003 | Operador BU-SP faz login → vê apenas BU-SP → não acessa BU-RJ | Playwright |
| **Dashboard admin** | EP-0001 | Admin visualiza métricas → filtra por período → métricas consistentes | Playwright |

### 2.4 Testes de Performance

| Cenário | Ferramenta | Target |
|---------|-----------|--------|
| **Carga em endpoints REST** | k6 (Grafana) | 100 VUs simultâneos; P95 < 500ms |
| **Dashboard métricas** | k6 | 50 usuários consultando dashboard simultaneamente |
| **Onboarding concorrente** | k6 | 20 novos tenants fazendo onboarding simultâneo |
| **RLS com volume** | k6 | 100 tenants com 1.000 BUs cada; query com RLS ativo |

### 2.5 Testes de Segurança

| Tipo | Ferramenta | Pipeline Stage |
|------|-----------|---------------|
| **SAST** | Semgrep | CI — block on CRITICAL |
| **Secret Scanning** | Gitleaks | Pre-commit + CI — block on ANY finding |
| **Dependency Scan** | Maven OWASP plugin + npm audit | CI — block on CRITICAL CVEs |
| **IaC Scan** | Terrascan | CI — Terraform security validation |
| **Penetration Test** | Manual + OWASP ZAP | Pré-lançamento (manual) |

---

## 3. Ambientes de Teste

| Ambiente | Propósito | Dados | Acesso |
|----------|-----------|-------|--------|
| **Dev** | Testes de desenvolvedor; integração contínua | Dados sintéticos (faker) | Devs |
| **Staging** | Validação pré-prod; smoke tests; E2E | Dados anonimizados (cópia de prod sanitizada) | QA + Devs |
| **Prod** | Monitoramento contínuo; health checks | Dados reais | Time autorizado |

### 3.1 Quality Gates por Ambiente

| Gate | Dev | Staging | Prod |
|------|-----|---------|------|
| **Unit tests pass** | ✅ Obrigatório | ✅ Obrigatório | ✅ Obrigatório |
| **Integration tests pass** | ✅ Obrigatório | ✅ Obrigatório | ✅ Obrigatório |
| **SAST (Semgrep) zero CRITICAL** | ✅ Obrigatório | ✅ Obrigatório | ✅ Obrigatório |
| **Secrets scan (Gitleaks) clean** | ✅ Obrigatório | ✅ Obrigatório | ✅ Obrigatório |
| **E2E smoke tests pass** | ❌ Opcional | ✅ Obrigatório | ✅ Obrigatório |
| **Performance test pass** | ❌ | ❌ | ✅ Pré-lançamento |
| **Security pentest pass** | ❌ | ❌ | ✅ Pré-lançamento |
| **Human approval** | ❌ | ✅ Manual gate | ✅ Manual gate |

---

## 4. Métricas de Qualidade

| Métrica | Target | Medição |
|---------|--------|---------|
| **Cobertura de código (linhas)** | ≥ 80% backend, ≥ 70% frontend | JaCoCo (backend), Jest coverage (frontend) |
| **Cobertura de cenários E2E** | 100% fluxos críticos | Playwright report |
| **Defeitos escapados para produção** | < 2 por release | Bug tracker |
| **Tempo de execução da suíte CI** | < 10 min | GitHub Actions logs |
| **Tempo de execução E2E** | < 15 min | Playwright report |

---

## 5. Riscos e Estimativa de Esforço

### 5.1 Riscos de Testes

| ID | Risco | Prob. | Impacto | Mitigação |
|----|-------|-------|---------|-----------|
| RT1 | RLS test coverage insuficiente → vazamento cross-tenant em produção | Baixa | 🔴 Crítico | Testes automatizados de isolamento em toda migration |
| RT2 | E2E tests frágeis (flaky) → falsos negativos → desconfiança da suíte | Média | 🟡 Médio | Retry lógico + report de flakiness; manutenção contínua |

### 5.2 Estimativa de Esforço

| Atividade | Complexidade | Esforço (dias) | Responsável |
|-----------|-------------|----------------|-------------|
| Estratégia de testes unitários e integração | Leve | 1 | Valeria Lucanete |
| Configuração Testcontainers (PostgreSQL + Redis) | Leve | 0.5 | Valeria Lucanete |
| Cenários E2E críticos (6 fluxos) | Moderada | 1.5 | Valeria Lucanete |
| Configuração k6 (performance) | Leve | 0.5 | Valeria Lucanete |
| Quality gates pipeline | Leve | 0.5 | Valeria Lucanete |
| Documentação e métricas | Leve | 0.5 | Valeria Lucanete |
| **Total Test Strategy** | — | **~4.5 dias** | — |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: Test Strategy Definition. Pirâmide 60/30/10, 5 tipos de teste, 6 fluxos E2E críticos, 3 ambientes, quality gates, estimativa ~4.5 dias | QA Engineer / Test Specialist |

---

🤖 *Upstream Architecture Discovery — Fase 6. Documento gerado pelo QA Engineer como parte do Bloco B.*
