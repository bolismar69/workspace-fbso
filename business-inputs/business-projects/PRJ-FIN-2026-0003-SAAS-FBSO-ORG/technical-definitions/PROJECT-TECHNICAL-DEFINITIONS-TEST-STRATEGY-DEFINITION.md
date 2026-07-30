# PROJECT-TECHNICAL-DEFINITIONS-TEST-STRATEGY-DEFINITION — Estratégia de Testes

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Fase:** F11 — Bloco B · **Disciplina:** Test Specialist
- **Versão:** 1.0 · **Data:** 30/07/2026 · **Status:** CREATED

---

## 1. Pirâmide de Testes
| Nível | Cobertura | Ferramenta |
|:---|:---:|:---|
| Unitários | ≥ 80% | JUnit 5, Mockito (Java) / Jest (Frontend) |
| Integração | ≥ 60% | Spring Boot Test, Testcontainers |
| E2E | Fluxos críticos | Playwright |
| Aceitação | Gherkin specs | Cucumber |

## 2. Automação
- **CI gate:** Testes unitários + integração em todo PR
- **Nightly:** E2E completo + testes de performance
- **Pre-release:** Aceitação + security scan

## 3. Performance
- **Carga:** k6 — 100, 500, 1000 VUs graduais
- **Stress:** k6 — até ponto de quebra
- **Benchmark:** JMH para componentes críticos (RBAC engine, multi-tenant queries)

## 4. Segurança
- **SAST:** SonarQube + Semgrep no CI
- **DAST:** OWASP ZAP scan semanal no Staging
- **Dependency:** Dependabot + Snyk

## 5. Quality Gates
| Gate | Critério |
|:---|:---|
| PR | Unit tests pass + coverage não cai + Semgrep sem HIGH/CRITICAL |
| Staging | E2E pass + k6 baseline OK + ZAP sem HIGH |
| Release | Todos os gates + sign-off QA |

🤖 *F11 — TEST-STRATEGY-DEFINITION · Roadmap v5.0*
