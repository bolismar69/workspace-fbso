# SPRINT-2-EXECUTION-REPORT-T015.1.md — Relatório de Execução: T-015.1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** `PRJ-FIN-2026-0003-SAAS-FBSO-ORG`
- **Sprint:** 2 de 7 — Segurança Cross-Cutting (task incremental)
- **Task:** T-015.1 — PostgreSQL Row-Level Security
- **Stack:** Java 25 + Spring Boot + PostgreSQL 17 + HikariCP
- **Data:** 14 de Julho de 2026

---

## 1. Resumo da Execução

| Métrica | Valor |
|:---|---|
| Task executada | T-015.1 |
| Status | ✅ Concluída |
| Testes | 33/33 passando |
| Novos testes | +11 (TenantAwareDataSourceTest: 6, RLSIsolationTest: 5) |
| Arquivos criados | 5 |
| Arquivos modificados | 2 |
| Build | ✅ SUCCESS |

---

## 2. Stack e Skills

- **Stack detectada:** Java 25 + Spring Boot + PostgreSQL (via PRD.md)
- **Skills aplicadas:** `313-frameworks-spring-db-migrations-flyway`, `304-frameworks-spring-boot-security`, `security-review`

---

## 3. Arquivos Criados ou Modificados

| Ação | Arquivo | Descrição |
|:---|:---|:---|
| 🆕 | `src/main/resources/db/migration/V003__enable_rls.sql` | RLS + políticas tenant_isolation em 5 tabelas |
| 🆕 | `src/main/resources/db/migration/U003__disable_rls.sql` | Rollback — remove RLS das 5 tabelas |
| 🆕 | `src/main/java/.../config/TenantAwareDataSource.java` | DataSource proxy — configura app.current_tenant_id em cada getConnection() |
| 🆕 | `src/main/java/.../config/DataSourceConfig.java` | BeanPostProcessor que encapsula HikariCP com TenantAwareDataSource |
| 🆕 | `src/test/java/.../unit/config/TenantAwareDataSourceTest.java` | 6 testes unitários do proxy DataSource |
| 🆕 | `src/test/java/.../integration/security/RLSIsolationTest.java` | 5 testes estruturais da migration V003 |
| 🔄 | `src/main/java/.../security/TenantContext.java` | Adicionado getTenantIdQuietly() para o DataSource proxy |
| 🔄 | `src/main/java/.../security/JwtAuthenticationFilter.java` | Javadoc atualizado — delega RLS ao TenantAwareDataSource |

---

## 4. Evidências de Testes

- **Build:** `./mvnw compile` → ✅ SUCCESS
- **Testes:** `./mvnw test -Djacoco.skip=true` → 33/33 passando, 0 falhas, 1 skip (CI-only)
- **Testes novos:** 11 (TenantAwareDataSourceTest: 6, RLSIsolationTest: 5)
- **Cenários SPRINT-TEST-SUITE.md:** TC-S2-022 a TC-S2-026 cobertos

### Cobertura de Cenários RLS

| ID | Descrição | Status |
|:---|:---|:---:|
| TC-S2-022 | V003: RLS habilitado em 5 tabelas | ✅ |
| TC-S2-023 | Política tenant_isolation criada | ✅ |
| TC-S2-024 | TenantAwareDataSource configura app.current_tenant_id | ✅ |
| TC-S2-025 | Conexão sem tenant_id → RESET (Admin FBSO) | ✅ |
| TC-S2-026 | SQLException no SET não propaga | ✅ |
| TC-INFRA-022 | V003 existe e habilita RLS | ✅ |
| TC-INFRA-023 | Políticas com USING + WITH CHECK | ✅ |
| TC-INFRA-024 | TenantAwareDataSource getConnection() seta tenant_id | ✅ |
| TC-INFRA-025 | U003 rollback existe | ✅ |
| TC-INFRA-026 | Tabelas globais sem RLS | ✅ |

---

## 5. Validação de Segurança

- [x] Nenhuma credencial hardcoded
- [x] SQL injection: tenant_id passado via string format (UUID validado pelo JWT — não há entrada de usuário)
- [x] Defesa em profundidade: RLS (banco) + BaseRepository (app) + Teste Isolamento (detecção)
- [x] Conexão sem tenant_id → RESET (não vaza tenant anterior da pool)
- [x] Admin FBSO (tenant_id=null) → acesso global (RLS bypass)

---

## 6. Validação de Arquitetura

- [x] Migration segue padrão VNNN__description.sql do Flyway
- [x] TenantAwareDataSource segue padrão DelegatingDataSource (Spring Framework)
- [x] BeanPostProcessor é static → sem dependência circular
- [x] ADR-L07 documentado em ARCHITECTURE.md §4.3 e §8

---

## 7. Desvios e Observações

1. **5 tabelas, não 11:** A especificação original mencionava "11 tabelas com tenant_id", mas a migration V001 real tem apenas 5 tabelas com coluna `tenant_id` (subscription, user, business_unit, product_service, audit_log). As outras 6 tabelas são globais ou herdam isolamento via JOIN.

2. **DataSource proxy vs SET no Filter:** A abordagem inicial de configurar `app.current_tenant_id` diretamente no `JwtAuthenticationFilter` foi descartada porque o `SET` afeta apenas uma conexão do pool HikariCP. A solução com `TenantAwareDataSource` (DataSource proxy) garante que cada `getConnection()` configure a variável de sessão corretamente, independentemente de qual conexão do pool for usada.

3. **JaCoCo incompatível com Java 25:** A versão 0.8.12 do JaCoCo não suporta class file major version 69 (Java 25). Testes executados com `-Djacoco.skip=true`. Atualização do JaCoCo necessária para relatórios de cobertura.

---

## 8. Próximos Passos

- **T-015.1 concluída.** Pipeline de segurança agora tem 3 camadas de isolamento multi-tenant.
- **Próxima:** Sprint 3 — Portal Admin + Contas/Planos (T-016 a T-038)
- **Recomendação:** Quando o PostgreSQL estiver disponível via Testcontainers, executar os testes de violação RLS com banco real (INSERT cross-tenant → POLICY violation)

---

🤖 *Relatório de execução gerado em 14/07/2026 pelo Agente de Execução. Task T-015.1 executada conforme PROMPT-EXECUTE-SPRINT-TASKS.md.*
