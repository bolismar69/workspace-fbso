# SPRINT-TEST-SUITE: Sprint 7 — Integração, Testes e Homologação

- **Sprint:** 7 de 7 🚀
- **Origem:** [TEST_PLAN.md](../../TEST_PLAN.md) §5, §6, §7 + todas as seções anteriores
- **Features:** Todas (F01-01 a F04-06 — 18 features)
- **Total de cenários:** 146 (regressão completa) + performance + OWASP + smoke

> 🚀 Sprint final. Esta suite consolida TODOS os testes do projeto. Execute sem exceção.

---


## 1. Regressão Completa (T-078)

Executar a suite completa conforme [TEST_PLAN.md §6.2](../../TEST_PLAN.md):

```bash
#!/bin/bash
echo "=== REGRESSAO COMPLETA: ms-fbso-platform-admin ==="

# 1. Testes unitarios (Sprints 1-6)
./mvnw test -Dtest="**/unit/**" && echo "✅ Unitarios" || exit 1

# 2. Testes de integracao (Sprints 1-6)
./mvnw test -Dtest="**/integration/**" -Dspring.profiles.active=test && echo "✅ Integracao" || exit 1

# 3. Testes de seguranca (Sprints 2-6)
./mvnw test -Dtest="**/security/**" && echo "✅ Seguranca" || exit 1

# 4. Testes E2E
docker compose -f infra/docker/docker-compose.test.yml up -d
./mvnw test -Dtest="**/e2e/**" && echo "✅ E2E" || exit 1

# 5. Cobertura
./mvnw jacoco:check && echo "✅ Cobertura >= 80%" || exit 1

# 6. Qualidade
./mvnw checkstyle:check pmd:check && echo "✅ Qualidade" || exit 1

echo "=== REGRESSAO COMPLETA ==="
```

### Checklist de Regressão por Feature

| Feature | Marco | Cenários | Status |
|:---|:---|:---:|:---:|
| F01-01 — Dashboard Admin | M2 | 7 | ⬜ |
| F01-02 — Lista de Contas | M2 | 4 | ⬜ |
| F01-03 — Alertas Dashboard | M2 | 5 | ⬜ |
| F02-01 — Criar Tenant | M3 | 7 | ⬜ |
| F02-02 — Transições Status | M3 | 9 | ⬜ |
| F02-03 — Planos | M3 | 7 | ⬜ |
| F02-04 — Assinaturas | M3 | 9 | ⬜ |
| F02-05 — Auditoria | M3 | 8 | ⬜ |
| F03-01 — Gestão Usuários | M4 | 7 | ⬜ |
| F03-02 — Matriz Permissões | M4 | 9 | ⬜ |
| F03-03 — Vinculação U×M | M4 | 6 | ⬜ |
| F03-04 — Acesso Condicional | M4 | 5 | ⬜ |
| F04-01 — Login | M5 | 9 | ⬜ |
| F04-02 — Onboarding | M5 | 10 | ⬜ |
| F04-03 — Dashboard Cliente | M5 | 5 | ⬜ |
| F04-04 — App Switcher | M5 | 4 | ⬜ |
| F04-05 — Business Units | M6 | 9 | ⬜ |
| F04-06 — Produtos | M6 | 7 | ⬜ |

**Total: 127 cenários por feature + 19 transversais = 146 cenários**

---

## 2. Testes de Performance (T-079)

| ID | Descrição | Carga | Meta | Status |
|:---|:---|:---|:---|:---:|
| TC-PERF-001 | Dashboard summary 1000 tenants, 50 usuários | 50 threads, ramp-up 10s | p95 ≤ 3s, p99 ≤ 5s, 0 erros | ⬜ |
| TC-PERF-002 | Dashboard evolution 12 meses, 20 usuários | 20 threads | p95 ≤ 3s | ⬜ |
| TC-PERF-003 | Lista paginada 1000 tenants, 30 usuários | 30 threads | p95 ≤ 2s | ⬜ |
| TC-PERF-004 | Criação simultânea 50 tenants | 50 threads, cada POST | 50 criados, 0 duplicatas, 0 deadlocks | ⬜ |
| TC-PERF-005 | Upsert permissão 1000 usuários, 10 BUs | 10 threads | p95 ≤ 2s | ⬜ |

### Testes de Concorrência

| ID | Descrição | Resultado Esperado | Status |
|:---|:---|:---|:---:|
| TC-RACE-001 | 2 assinaturas simultâneas mesmo tenant | Uma 201, outra 409. 1 ativa | ⬜ |
| TC-RACE-002 | 2 atualizações simultâneas preço plano | Versão incrementada 2x. Sem perda | ⬜ |
| TC-RACE-003 | Suspensão + reativação simultâneas | Estado final consistente. Auditoria OK | ⬜ |

---

## 3. Testes de Segurança OWASP (T-082)

| ID | Categoria | Descrição | Status |
|:---|:---|:---|:---:|
| TC-SEC-OWASP-001 | A1 — SQL Injection | Busca textual: `' OR 1=1 --` | ⬜ |
| TC-SEC-OWASP-002 | A1 — SQL Injection | UUID: `'; SELECT pg_sleep(10); --` | ⬜ |
| TC-SEC-OWASP-003 | A2 — Broken Auth | JWT assinatura inválida | ⬜ |
| TC-SEC-OWASP-004 | A2 — Broken Auth | JWT `alg: none` | ⬜ |
| TC-SEC-OWASP-005 | A2 — Broken Auth | Replay attack pós-logout | ⬜ |
| TC-SEC-OWASP-007 | A3 — XSS | corporate_name=`<script>alert('XSS')</script>` | ⬜ |
| TC-SEC-OWASP-008 | A3 — XSS | search=`<script>alert('XSS')</script>` | ⬜ |
| TC-SEC-OWASP-009 | A5 — Broken Access Control | Method smuggling | ⬜ |
| TC-SEC-OWASP-010 | A5 — Broken Access Control | IDOR — alterar permissão de outro | ⬜ |
| TC-SEC-OWASP-011 | A6 — Security Misconfig | Stack traces em erros | ⬜ |
| TC-SEC-OWASP-012 | A6 — Security Misconfig | Headers de segurança | ⬜ |
| TC-SEC-OWASP-013 | A8 — CSRF | CORS origem não autorizada | ⬜ |
| TC-SEC-OWASP-015 | A9 — Vuln Components | `mvn dependency-check:check` | ⬜ |
| TC-SEC-OWASP-016 | A10 — Insufficient Logging | Ações suspeitas em audit_log | ⬜ |

---

## 4. Smoke Tests Pós-Deploy (T-084)

| ID | Verificação | Comando | Resultado | Status |
|:---|:---|:---|:---|:---:|
| SMOKE-01 | Health check | `curl /actuator/health` | `{"status":"UP"}` | ⬜ |
| SMOKE-02 | Autenticação | `curl /api/v1/tenants -H "Authorization: Bearer <jwt>"` | 200 | ⬜ |
| SMOKE-03 | Erro 401 sem token | `curl /api/v1/tenants` | 401 | ⬜ |
| SMOKE-04 | Erro 403 sem permissão | `curl /api/v1/plans -H "Authorization: Bearer <jwt-operator>"` | 403 | ⬜ |
| SMOKE-05 | RFC 7807 | `curl /api/v1/tenants/99999` | 404 RFC 7807 | ⬜ |
| SMOKE-06 | Banco conectado | `curl /actuator/health` | "db" = "UP" | ⬜ |

---

## 5. Verificação LGPD (T-083)

| ID | Verificação | Critério | Status |
|:---|:---|:---|:---:|
| LGPD-01 | Logs sem dados pessoais | Nome, email, CNPJ mascarados em logs | ⬜ |
| LGPD-02 | Soft delete 11/11 entidades | Toda tabela com coluna `deleted_dt` | ⬜ |
| LGPD-03 | Auditoria preserva dados excluídos | Registro em audit_log com previous_value | ⬜ |
| LGPD-04 | Resposta HTTP não expõe dados outros tenants | Verificado em TC-SEC-DL-* | ⬜ |

---

## 📊 Resumo Final

| Nível | Cenários |
|:---|:---:|
| Unit (regressão) | 70 |
| Integração (regressão) | 45 |
| E2E (regressão) | 16 |
| Segurança (regressão + OWASP) | 25 |
| Performance (JMeter/k6) | 5 |
| Smoke (pós-deploy) | 6 |
| LGPD | 4 |
| **Total** | **171** |

---

## 🚀 Critério de Go / No-Go para Produção

| Check | Go | No-Go |
|:---|:---:|:---:|
| Regressão | 100% verde | Qualquer falha |
| Cobertura | ≥ 80% | < 80% |
| Performance p95 | ≤ 3s | > 3s |
| SAST | 0 críticos/high | ≥ 1 crítico/high |
| UAT PO | Aprovado | Não aprovado |
| Smoke staging | 6/6 | < 6/6 |

---

🤖 *Extraído de TEST_PLAN.md v2.0. Execute esta suite completa antes do go-live. Sem exceções.*
