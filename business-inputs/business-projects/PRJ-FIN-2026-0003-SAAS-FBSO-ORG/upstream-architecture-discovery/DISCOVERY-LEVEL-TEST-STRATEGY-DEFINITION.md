# DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION — Estratégia de Testes (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F6 — Bloco B · **Disciplina:** Test Specialist
- **Versão:** 1.0 · **Data:** 30/07/2026 · **Status:** CREATED

## 1. Pirâmide de Testes (Macro)
| Nível | Cobertura Alvo | Ferramenta |
|:---|:---:|:---|
| Unitários | ≥ 80% | JUnit 5 (BE) / Jest (FE) |
| Integração | ≥ 60% | Spring Boot Test + Testcontainers |
| E2E | Fluxos críticos | Playwright |
| Performance | Smoke test por release | k6 |

## 2. Ambientes de Teste
| Ambiente | Dados | Isolamento |
|:---|:---|:---|
| Dev | Mock/anonymized | Por feature branch |
| Staging | Anonymized (produção-like) | Compartilhado |
| CI | Docker Compose efêmero | Isolado por PR |

## 3. Quality Gates
| Gate | Critério |
|:---|:---|
| PR | Unit tests pass + coverage estável + Semgrep limpo |
| Staging | E2E pass + k6 smoke OK |
| Release | Todos os gates + sign-off QA |

## 4. Riscos de Qualidade
| Risco | Mitigação |
|:---|:---|
| Cobertura insuficiente de isolamento multi-tenant | Testes específicos de RLS por tenant |
| Frontend sem dev dedicado → poucos testes E2E | Playwright com cenários críticos priorizados |
| Time reduzido → QA sobrecarregado | Automação máxima; QA foco em exploratório |

## 5. Estimativa de Esforço
| Atividade | Esforço |
|:---|:---:|
| Setup de frameworks de teste | 0.5-1 homem-mês |
| Testes unitários + integração (base) | 1-2 homem-mês |
| Testes E2E (Playwright — cenários críticos) | 1 homem-mês |
| **Total Testes** | **2.5-4 homem-mês** |

🤖 *Upstream Architecture Discovery — Fase 6*
