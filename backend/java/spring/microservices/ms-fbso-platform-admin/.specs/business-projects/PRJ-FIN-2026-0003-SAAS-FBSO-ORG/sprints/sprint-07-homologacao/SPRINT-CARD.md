# SPRINT-CARD: Sprint 7 — Integração, Testes e Homologação

- **Sprint:** 7 de 7 🚀
- **Marco:** M7 — Homologação
- **Datas:** 15/10/2026 → 30/10/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [TEST_PLAN.md](../../TEST_PLAN.md) · [SPECS.md](../../SPECS.md)

---

## 🎯 Sprint Goal

**"Suite de regressão 100% verde. Dashboard p95 ≤ 3s com 1000 tenants. OpenAPI 3.0 documentando 37 endpoints. Zero vulnerabilidades críticas/high em SAST. Deploy staging validado com smoke tests. Deploy produção com rollback plan documentado. LGPD: zero dados pessoais em logs. Go-live em 30/10."**

> 🚀 **Sprint final. Tudo converge aqui.**

---

## 📋 Sprint Backlog

| ID | Tarefa | Est. | Critério DONE |
|:---|:---|:---:|:---|
| **T-071** | Testes de regressão: bateria completa M2-M6. Automatizar como suite única | 2d | Suite executa sem falhas. Zero quebras em features já homologadas |
| **T-072** | Testes de performance: dashboard ≤3s (p95) com 1000 tenants. Listas paginadas 10k registros. JMeter ou k6 | 2d | p95 ≤ 3s. Listas ≤ 1s. Relatório de carga |
| **T-073** | Criar/atualizar `fbso-platform-api.yaml` (OpenAPI 3.0): 37 endpoints, schemas, exemplos, erros (BR-NFR04) | 2d | OpenAPI válido. 37 endpoints. Schemas com exemplos |
| **T-074** | Documentar `README.md` raiz + README.md por pacote (controller, service, repository, security) (BR-NFR04) | 1.5d | Quickstart. README por pacote com responsabilidade |
| **T-075** | SAST scan (Semgrep / SonarQube): zero críticas/high. Verificar OWASP Top 10: SQL injection, broken auth, XSS, CSRF | 2d | 0 críticos/high. PreparedStatement. JSON escapado |
| **T-076** | Verificar LGPD: dados pessoais mascarados em logs. 100% entidades com soft delete (NFR-LGPD) | 1d | Logs sem dados pessoais. 11 entidades com deleted_dt |
| **T-077** | Deploy staging (K8s): manifests, smoke test, health checks (liveness, readiness). BR-NFR01 (99,5%) | 2d | Staging operacional. Health check UP. Smoke test passa |
| **T-078** | UAT com Product Owner: demonstração 18 features. Correção de bugs. Aprovação formal (DoD) | 2d | PO valida. Bugs corrigidos. Termo de aceite |
| **T-079** | Deploy produção (K8s): promote staging → prod. Validação pós-go-live (KPIs). Rollback plan. Monitoramento | 2d | Produção operacional. KPIs monitorados. Rollback testado |

**Total:** 9 tarefas · ~16.5 dias-homem

---

## 📦 Entregáveis da Sprint

| Entregável | Descrição | Task |
|:---|:---|:---:|
| Suite de regressão | Script automatizado com unit + integração + segurança | T-071 |
| Relatório de performance | JMeter/k6: p95 ≤ 3s dashboard, ≤ 1s listas | T-072 |
| OpenAPI 3.0 | `fbso-platform-api.yaml` — 37 endpoints documentados | T-073 |
| Documentação | README.md raiz + por pacote | T-074 |
| Relatório SAST | 0 críticos/high. OWASP verificado | T-075 |
| Relatório LGPD | Zero dados pessoais em logs. Soft delete 100% | T-076 |
| Deploy staging | K8s com health checks e smoke tests | T-077 |
| UAT aprovado | PO assina termo de aceite | T-078 |
| **Deploy produção** 🚀 | Go-live com monitoramento e rollback plan | T-079 |

---

## ✅ Definition of Done (Sprint-Level)

- [ ] `mvn test` (regressão completa) → 100% verde
- [ ] JaCoCo ≥ 80% (linhas), ≥ 70% (branchs)
- [ ] p95 dashboard ≤ 3s com 1000 tenants (JMeter)
- [ ] `fbso-platform-api.yaml` válido (sem erros de schema)
- [ ] 37 endpoints documentados com exemplos
- [ ] README.md raiz com quickstart (como rodar, testar, deploy)
- [ ] SAST: 0 vulnerabilidades críticas/high
- [ ] LGPD: logs sem dados pessoais, 11 entidades com deleted_dt
- [ ] Staging: health check UP, smoke tests passam
- [ ] UAT: PO aprova (termo de aceite assinado)
- [ ] Produção: KPIs monitorados, rollback testado

---

## ⚠️ Riscos e Bloqueadores (Sprint Final)

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Teste de regressão revelar bug crítico em feature de sprint anterior | Média | Alto | Reservar 2 dias de buffer para correções (já embutido em T-078) |
| Performance abaixo da meta (p95 > 3s) | Média | Alto | Otimização de queries (índices), cache, HikariCP tuning. Testar desde o início da sprint |
| SAST encontrar vulnerabilidade crítica | Baixa | Alto | Rodar scan preliminar na Sprint 6 para antecipar correções |
| UAT revelar funcionalidade faltante | Média | Médio | Alinhar expectativas com PO durante a Sprint 6 |
| Deploy produção falhar | Baixa | Crítico | Rollback plan documentado e testado em staging |

---

## 🔗 Dependências

- **Pré-requisitos:** TODAS as sprints anteriores (1 a 6) concluídas.
- **Dependências externas:** K8s cluster (staging + prod), Keycloak produção, PostgreSQL produção, SMTP produção.

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Regressão | 146/146 cenários verdes |
| Cobertura JaCoCo | ≥ 80% linhas |
| Performance p95 dashboard | ≤ 3s |
| SAST | 0 críticos/high |
| OpenAPI endpoints | 37/37 documentados |
| Soft delete entidades | 11/11 |
| Deploy staging → produção | Sem rollback |

---

## 🚀 Plano de Go-Live

```
D-5 (23/10): Suite de regressão verde
D-4 (24/10): Performance OK. SAST zerado
D-3 (27/10): Deploy staging. Smoke tests
D-2 (28/10): UAT com PO. Aprovação
D-1 (29/10): Deploy produção (horário comercial)
D-0 (30/10): Validação pós-go-live. Monitoramento 24h
```

---

🤖 *Gerado a partir de TASKS.md v2.0. Sprint final — go-live em 30/10. Tudo converge aqui.*
