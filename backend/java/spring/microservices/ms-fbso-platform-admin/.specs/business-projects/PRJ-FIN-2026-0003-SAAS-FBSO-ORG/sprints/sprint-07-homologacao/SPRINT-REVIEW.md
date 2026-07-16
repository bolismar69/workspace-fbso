# SPRINT-REVIEW: Sprint 7 — Homologação e Go-Live

- **Sprint:** 7 de 7 🚀
- **Data da Review:** 30/10/2026
- **Participantes:** Time Técnico, Tech Lead, **Product Owner**, Stakeholders
- **Features:** Todas (18 features — F01-01 a F04-06)

> 🚀 **Review final. Go-live em 30/10.** Esta é a revisão de homologação — o sistema está completo.

---

## 🎯 O Que Demonstrar

### 1. Suite de Regressão (T-071)

- [ ] Executar `./mvnw test` (regressão completa) ao vivo OU mostrar relatório da última execução
- [ ] **Resultado esperado:** 146/146 cenários verdes
- [ ] Mostrar relatório JaCoCo: ≥ 80% linhas, ≥ 70% branchs

### 2. Performance (T-072)

- [ ] Dashboard admin com 1000 tenants: p95 ≤ 3s
- [ ] Listas paginadas (1000 registros): p95 ≤ 1s
- [ ] Mostrar relatório JMeter/k6
- [ ] Testes de concorrência: zero race conditions

### 3. OpenAPI 3.0 (T-073)

- [ ] Abrir `fbso-platform-api.yaml` no Swagger UI
- [ ] Mostrar 37 endpoints documentados
- [ ] Demonstrar exemplos de request/response
- [ ] Códigos de erro documentados por endpoint

### 4. Documentação (T-074)

- [ ] Mostrar `README.md` raiz com quickstart
- [ ] Mostrar README.md por pacote (controller, service, repository, security)
- [ ] "Novo desenvolvedor consegue rodar o projeto em ≤30min?"

### 5. Segurança (T-075, T-076)

- [ ] Relatório SAST (Semgrep/SonarQube): 0 críticos/high
- [ ] OWASP Top 10 verificado
- [ ] LGPD: demonstrar logs sem dados pessoais
- [ ] Soft delete: `\dt fbso_platform.*` — 11 tabelas com `deleted_dt`

### 6. Deploy Staging (T-077)

- [ ] Smoke test staging: `curl /actuator/health` → `{"status":"UP"}`
- [ ] Autenticação funcional em staging
- [ ] Deploy em K8s: mostrar `kubectl get pods`

### 7. Demonstração Completa para o PO (T-078)

- [ ] **Dashboard Admin:** Métricas, lista de contas, alertas
- [ ] **Gestão:** Criar tenant, suspender, reativar. Criar/editar/desativar plano. Assinar/upgrade
- [ ] **RBAC:** Login com cada papel, demonstrar acesso condicional
- [ ] **Onboarding:** Fluxo completo PENDING→ACTIVE em 4 passos
- [ ] **BUs:** Hierarquia Matriz/Filial, CNPJ único, reúso
- [ ] **Catálogo:** Criar/listar/desativar produtos, indicador "Não mapeado"
- [ ] **Auditoria:** Consulta com filtros, imutabilidade

---

## 📋 Checklist de Go / No-Go

| Verificação | Status | Go? |
|:---|:---:|:---:|
| Regressão 100% verde | ⬜ | |
| Cobertura ≥ 80% | ⬜ | |
| Dashboard p95 ≤ 3s | ⬜ | |
| SAST 0 críticos/high | ⬜ | |
| OpenAPI 37/37 endpoints | ⬜ | |
| LGPD logs OK | ⬜ | |
| Soft delete 11/11 entidades | ⬜ | |
| Smoke staging 6/6 | ⬜ | |
| UAT PO aprovado | ⬜ | |
| Rollback testado | ⬜ | |
| Monitoramento configurado | ⬜ | |
| **GO-LIVE** 🚀 | | |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## 🚀 Pós-Go-Live (30/10)

- [ ] Validar KPIs: uptime, latência p95, erros 4xx/5xx
- [ ] Monitorar alertas por 24h
- [ ] Rollback plan em mãos se necessário
- [ ] Agendar retrospectiva do projeto

---

## 🎉 Encerramento do Projeto

**18 features. 79 tarefas. 7 sprints. 14 semanas.**

| Marco | Data | Status |
|:---|:---|:---:|
| Sprint 1 — Setup | 07/08 | ⬜ |
| Sprint 2 — Segurança | 15/08 | ⬜ |
| Sprint 3 — Portal Admin + Contas | 31/08 | ⬜ |
| Sprint 4 — RBAC | 15/09 | ⬜ |
| Sprint 5 — Portal Cliente | 30/09 | ⬜ |
| Sprint 6 — BUs e Catálogo | 15/10 | ⬜ |
| Sprint 7 — Homologação | 30/10 | ⬜ |

---

🤖 *Checklist de review da Sprint 7. Última revisão antes do go-live. 🚀*
