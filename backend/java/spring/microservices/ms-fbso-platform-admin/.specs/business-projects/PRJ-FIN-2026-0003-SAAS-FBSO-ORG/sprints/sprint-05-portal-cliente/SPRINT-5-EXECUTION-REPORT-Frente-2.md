# SPRINT-5-EXECUTION-REPORT-Frente-2.md — Relatório de Execução: Sprint 5 — Frente 2

- **Solução:** `ms-fbso-platform-admin`
- **Projeto de Negócio:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 2 — Desejáveis (🔵 Could Have)
- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4
- **Data da execução:** 2026-07-23
- **Tasks executadas:** T-149.DT-086 a T-156.DT-113 (8 tasks)

---

## 1. Resumo da Execução

| Métrica | Valor |
|:---|---|
| **Tasks executadas** | 8/8 (100%) |
| **Tasks com sucesso** | 8 |
| **Tasks com falha** | 0 |
| **Tempo estimado** | ~4.5h |
| **Tempo efetivo** | ~1.5h (paralelização: todas as 8 tasks são independentes) |
| **Build** | ✅ SUCCESS |
| **Testes unitários** | 213 testes executados, 0 failures, 1 pre-existing error (SubscriptionServiceTest), 8 skipped |
| **Code Review** | ✅ 7 skills acionados, 1 achado Low (não bloqueante) |

---

## 2. Stack e Skills Utilizadas

- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4 + JWT (Nimbus)
- **Fonte da stack:** PRD.md §5.1 + ARCHITECTURE.md header
- **Skills aplicáveis:**
  - `121-java-object-oriented-design` — AuditFieldsRowMapper helper (T-149)
  - `141-java-refactoring-with-modern-features` — DRY refactoring RowMappers (T-149)
  - `301-frameworks-spring-boot-core` — ObjectMapper injection, CORS @Value (T-150, T-153)
  - `304-frameworks-spring-boot-security` — SecurityConfig refactoring (T-153)
  - `112-java-maven-plugins` — springdoc bump (T-152)

### Skills da Fase 7 (Code Review)

| Skill | Achados | Veredito |
|:---|:---:|:---|
| `ponytail-audit` | 0 | Lean already. Ship. (-4 linhas net) |
| `ponytail-review` | 1 Low | Javadoc órfão no SecurityConfig (não bloqueante) |
| `engineering-skills` | 0 | SOLID/DRY/KISS ✅ |
| `security-audit` | 0 | OWASP Top 10 ✅ |
| `performance-review` | 0 | Neutro/Positivo ✅ |
| `requesting-code-review` | 0 | Nomenclatura e convenções ✅ |
| `differential-review` | 0 | Blast radius baixo ✅ |

---

## 3. Tasks Executadas

| ID | Tarefa | DT | Status | Build | Observações |
|:---|:---|:---|:---:|:---:|:---|
| **T-149.DT-086** | Extrair helper `AuditFieldsRowMapper` | DT-086 | ✅ | ✅ | Helper estático criado. 4 RowMappers refatorados. 24 linhas duplicadas eliminadas. |
| **T-150.DT-089** | Injete `ObjectMapper` do Spring no `AuditAspect` | DT-089 | ✅ | ✅ | `new ObjectMapper()` removido. Injeção por construtor. `AuditAspectTest` atualizado. |
| **T-151.DT-090** | Substituir `OffsetDateTime.now()` → `now(ZoneOffset.UTC)` | DT-090 | ✅ | ✅ | 13 ocorrências corrigidas em 9 arquivos (7 planejadas + 6 adicionais descobertas na Fase 5). |
| **T-152.DT-092** | Bump `springdoc-openapi` 2.8.8→2.8.16 | DT-092 | ✅ | ✅ | Versão atualizada no pom.xml. Compilação OK. |
| **T-153.DT-093** | Externalizar CORS origins para `application.yml` | DT-093 | ✅ | ✅ | `app.cors.allowed-origins` com env var `CORS_ALLOWED_ORIGINS`. Default mantido. |
| **T-154.DT-101** | Atualizar mitigação de riscos no SPRINT-CARD.md | DT-101 | ✅ | N/A | 3 riscos marcados como eliminados. 4 mitigações atualizadas com ✅. Novo risco JWT adicionado. |
| **T-155.DT-112** | Atualizar header do SPECS.md | DT-112 | ✅ | N/A | Header atualizado: 24/40 (60%). Adicionado "Próximo: Frente 3". |
| **T-156.DT-113** | Recalcular progresso no TASKS.md | DT-113 | ✅ | N/A | Progresso: 95→105/167 (57%→63%). Frentes 0-1-2 marcadas ✅. |

---

## 4. Arquivos Criados ou Modificados

### 🆕 Criados (3 source + 3 artefatos)

| Ação | Arquivo | Task | Descrição da Mudança |
|:---|:---|:---|:---|
| 🆕 | `src/main/java/.../repository/rowmapper/AuditFieldsRowMapper.java` | T-149 | Helper estático `mapAuditFields(ResultSet, BaseEntity)` — 6 campos de auditoria |
| 🆕 | `SPRINT-DEVELOPMENT-PLANNING-Frente-2.md` | Fase 1 | Plano de desenvolvimento com análise PonteTail |
| 🆕 | `SPRINT-TEST-PLANNING-Frente-2.md` | Fase 3 | Plano de testes: 8 tasks mapeadas |
| 🆕 | `SPRINT-CODE-REVIEW-Frente-2.md` | Fase 7 | Code review: 7 skills, 1 achado Low |

### 🔄 Modificados (19 source + 3 artefatos)

| Ação | Arquivo | Task | Descrição da Mudança |
|:---|:---|:---|:---|
| 🔄 | `BaseEntity.java` | T-151 | 3× `now()` → `now(ZoneOffset.UTC)` |
| 🔄 | `BaseRepository.java` | T-151 | 3× `now()` → `now(ZoneOffset.UTC)` |
| 🔄 | `AuditAspect.java` | T-150, T-151 | ObjectMapper injetado + 1× UTC |
| 🔄 | `AuditAspectTest.java` | T-150 | Construtor atualizado com ObjectMapper |
| 🔄 | `TenantRowMapper.java` | T-149 | 6 linhas → `AuditFieldsRowMapper.mapAuditFields()` |
| 🔄 | `PlanRowMapper.java` | T-149 | 6 linhas → `AuditFieldsRowMapper.mapAuditFields()` |
| 🔄 | `SubscriptionRowMapper.java` | T-149 | 6 linhas → `AuditFieldsRowMapper.mapAuditFields()` |
| 🔄 | `UserRowMapper.java` | T-149 | 6 linhas → `AuditFieldsRowMapper.mapAuditFields()` |
| 🔄 | `SecurityConfig.java` | T-153 | CORS externalizado via `@Value` |
| 🔄 | `application.yml` | T-153 | Nova seção `app.cors.allowed-origins` |
| 🔄 | `pom.xml` | T-152 | springdoc 2.8.8→2.8.16 |
| 🔄 | `DashboardService.java` | T-151 | 1× `now()` → `now(ZoneOffset.UTC)` |
| 🔄 | `UserService.java` | T-151 | 1× `now()` → `now(ZoneOffset.UTC)` |
| 🔄 | `SubscriptionService.java` | T-151 | 2× `now()` → `now(ZoneOffset.UTC)` |
| 🔄 | `Subscription.java` | T-151 | 1× `now()` → `now(ZoneOffset.UTC)` |
| 🔄 | `AuditEntry.java` | T-151 | 1× `now()` → `now(ZoneOffset.UTC)` |
| 🔄 | `SPRINT-CARD.md` | T-154 | Riscos atualizados. Métricas: 24/40 (60%). |
| 🔄 | `SPECS.md` | T-155 | Header: 24/40 (60%). Próximo: Frente 3. |
| 🔄 | `TASKS.md` | T-156 | Progresso: 105/167 (63%). Frentes 0-1-2 ✅. |

**Total:** 22 arquivos (3 🆕 source + 3 🆕 artefatos + 16 🔄)

---

## 5. Evidências de Testes

| Comando | Resultado |
|:---|:---|
| `mvn compile` | ✅ BUILD SUCCESS (7.0s) |
| `mvn test` (full) | ✅ 213 testes: 0 failures, 1 pre-existing error, 8 skipped |

**Erro Pré-Existente (NÃO causado pela Frente 2):**

```
SubscriptionServiceTest > create > shouldCreateWithLockedPrice: ERROR
  IllegalStateException: TenantContext não inicializado
```

> Documentado no [SPRINT-5-EXECUTION-REPORT-Frente-1.md](./SPRINT-5-EXECUTION-REPORT-Frente-1.md) §5. Correção recomendada na Frente 3.

### Verificações Estáticas Pós-Implementação

| Verificação | Comando | Resultado |
|:---|:---|:---|
| UTC zero | `grep -rn "OffsetDateTime.now()" src/main/java/` | ✅ VAZIO |
| DRY RowMappers | `grep -rn "setCreatedDt\|..." src/main/java/**/rowmapper/` | ✅ Apenas `AuditFieldsRowMapper.java` |
| ObjectMapper manual | `grep "new ObjectMapper()" ...AuditAspect.java` | ✅ VAZIO |
| Credenciais hardcoded | Revisão manual do diff | ✅ Zero |
| SQL injection | Revisão de concatenação SQL | ✅ Zero |

---

## 6. Validação de Segurança

- [x] Nenhuma credencial ou dado sensível hardcoded — CORS usa `@Value` com default seguro
- [x] Queries usam parametrização — `findAllKeyset()` com `?` placeholders e `sanitizeColumn()`
- [x] Controles de acesso mantidos — sem alterações em RBAC ou JWT
- [x] Dados pessoais não expostos em logs — `AuditAspect.writeAuditLog()` loga apenas metadados
- [x] Respostas de erro seguem RFC 7807 — sem stack traces
- [x] `springdoc-openapi` 2.8.16 sem CVEs conhecidas

---

## 7. Validação de Arquitetura

- [x] Estrutura de diretórios segue ARCHITECTURE.md §2
  - `AuditFieldsRowMapper` em `repository/rowmapper/` ✅
  - `RateLimitFilter`, `FbsoJwtAuthenticationConverter` em `security/` ✅ (Frente 1)
- [x] Convenções de nomenclatura respeitadas
  - `AuditFieldsRowMapper` — sufixo RowMapper descritivo
  - `mapAuditFields(ResultSet, BaseEntity)` — verbo + objeto
- [x] Padrões de projeto documentados nas ADRs foram seguidos:
  - **ADR-L01 (JDBC Template):** BaseRepository.findAllKeyset usa JdbcTemplate
  - **ADR-L04 (RFC 7807):** GlobalExceptionHandler segue formato padronizado
  - **ADR-L06 (Package-by-Layer):** Novos arquivos nos pacotes corretos
- [x] Nenhuma quebra de API — zero endpoints alterados
- [x] Nenhuma dependência nova — springdoc apenas atualização de versão

---

## 8. Desvios e Observações

### Desvios do Planejamento Original

| Task | Planejado | Realizado | Justificativa |
|:---|:---|:---|:---|
| T-151.DT-090 | 7 ocorrências em 3 arquivos | 13 ocorrências em 9 arquivos | Fase 5 (grep) revelou 6 usos adicionais em Services/Entities — corrigidos para completude |
| T-150.DT-089 | Injeção via `@Autowired` | Injeção por construtor | Padrão Spring recomendado (constructor injection > field injection) |

### Decisões de Design Tomadas

1. **`AuditFieldsRowMapper` como `final class` + construtor `private`:** Padrão para classes utilitárias — impede instanciação e herança indevida.

2. **`AuditEntryRowMapper` NÃO usa o helper:** A tabela `audit_log` tem estrutura diferente (campo `timestamp` em vez de `created_dt`) e não estende `BaseEntity`. Correto mantê-lo independente.

3. **CORS via `@Value` com split por vírgula:** Alternativa mais simples que `@ConfigurationProperties` para uma única property. Trade-off: não suporta valores com vírgulas nos origins (improvável).

4. **6 UTC adicionais descobertos na Fase 5:** Demonstra o valor da validação automatizada pós-implementação como sanity check.

---

## 9. Próximos Passos

1. **Fase 10 — Atualização de Artefatos:** Atualizar SPRINT-CARD.md, SPRINT-TEST-SUITE.md, SPRINT-REVIEW.md, TASKS.md, SPECS.md, TEST_PLAN.md, ARCHITECTURE.md, PRD.md, sprints/README.md.

2. **Frente 3 — Features:** 16 tarefas (T-057..T-068 backend + T-157..T-160 frontend). Início recomendado após preenchimento do time técnico (10 papéis com `<TODO>` no TECHNICAL-TEAM-MAP.md).

3. **Corrigir teste pré-existente:** `SubscriptionServiceTest.shouldCreateWithLockedPrice` — adicionar `TenantContext.set(...)` no `@BeforeEach`.

---

🤖 *Relatório gerado em 2026-07-23 como parte da Fase 9 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4. 8/8 tasks concluídas (100%).*
