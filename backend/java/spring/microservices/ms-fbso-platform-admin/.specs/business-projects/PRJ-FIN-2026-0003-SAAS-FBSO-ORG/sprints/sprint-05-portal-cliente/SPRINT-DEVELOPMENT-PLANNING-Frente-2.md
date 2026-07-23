# SPRINT-DEVELOPMENT-PLANNING-Frente-2.md — Plano de Desenvolvimento: Sprint 5 — Frente 2

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 2 — Desejáveis (🔵 Could Have)
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Keycloak 26 · Flyway 12.11.0 · Caffeine 3.2.4 · JWT (Nimbus)
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-05-portal-cliente`
- **Data do planejamento:** 2026-07-23

---

## 1. Visão Geral

- **Objetivo da Frente 2:** Resolver 8 débitos técnicos desejáveis (todos Could Have) — melhorias de qualidade de código, configuração, documentação e consistência. Estas tarefas não bloqueiam features mas reduzem dívida técnica acumulada e melhoram a manutenibilidade do codebase.
- **Tasks a implementar:** 8 (8 Could)
- **Ordem de execução:** Paralela — todas as tasks são independentes entre si. Tasks de código (T-149, T-150, T-151, T-152, T-153) e tasks de documentação (T-154, T-155, T-156) podem ser executadas simultaneamente.
- **Estimativa total:** ~4.5h (≈0.5 dia serial, ≈2h paralelo)
- **Papéis envolvidos:** Dev Backend, DevOps, Dev BD, Tech Lead

---

## 2. Dependências entre Tasks

```
┌──────────────────────────────────────────────────────────────┐
│ TODAS AS 8 TASKS SÃO INDEPENDENTES ENTRE SI                  │
│                                                              │
│ T-149.DT-086 (AuditFieldsRowMapper)   — código isolado       │
│ T-150.DT-089 (ObjectMapper Injection) — código isolado       │
│ T-151.DT-090 (OffsetDateTime UTC)     — código isolado       │
│ T-152.DT-092 (Springdoc bump)         — pom.xml isolado      │
│ T-153.DT-093 (CORS externalização)    — config + yml isolado │
│ T-154.DT-101 (SPRINT-CARD riscos)     — documentação isolada │
│ T-155.DT-112 (SPECS.md header)        — documentação isolada │
│ T-156.DT-113 (TASKS.md progresso)     — documentação isolada │
│                                                              │
│ ⇒ EXECUÇÃO 100% PARALELIZÁVEL                                │
└──────────────────────────────────────────────────────────────┘
```

**Justificativa:** Nenhuma task da Frente 2 depende de outra. As 5 tasks de código alteram arquivos diferentes. As 3 tasks de documentação alteram artefatos diferentes. Execução paralela total é possível.

---

## 3. Análise PonteTail (7 Rungs) por Task

> **Checklist aplicado antes de planejar cada task.** Se alguma task falhar no Rung 1 (YAGNI), ela é removida do plano.

### Resumo da Análise

| Task | R1-YAGNI | R2-JáExiste | R3-Stdlib | R4-DepExistente | R5-Padrão | R6-Simples | R7-Mínimo | Veredito |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---|
| T-149.DT-086 | ✅ | ❌ (duplicado em 4 arquivos) | ✅ JDBC RowMapper | ✅ spring-jdbc | ✅ RowMapper pattern | ✅ 1 helper class | ✅ | Prosseguir |
| T-150.DT-089 | ✅ | ❌ (`new ObjectMapper()`) | ✅ Jackson ObjectMapper | ✅ spring-boot-starter-web | ✅ @Autowired | ✅ 1 injeção | ✅ | Prosseguir |
| T-151.DT-090 | ✅ | ❌ (sem UTC explícito) | ✅ java.time | ✅ Java 25 | ✅ Padrão UTC | ✅ 7 linhas | ✅ | Prosseguir |
| T-152.DT-092 | ✅ | ❌ (v2.8.8) | ✅ Maven BOM | ✅ springdoc-openapi | ✅ pom.xml | ✅ 1 versão | ✅ | Prosseguir |
| T-153.DT-093 | ✅ | ❌ (hardcoded) | ✅ Spring @Value | ✅ spring-boot | ✅ Config externalizada | ✅ 1 property | ✅ | Prosseguir |
| T-154.DT-101 | ✅ | ❌ (parcial — já tem refs) | N/A (docs) | N/A | ✅ SPRINT-CARD.md | ✅ | ✅ | Prosseguir |
| T-155.DT-112 | ✅ | ❌ (não tem campo) | N/A (docs) | N/A | ✅ SPECS.md header | ✅ 1 linha | ✅ | Prosseguir |
| T-156.DT-113 | ✅ | ❌ (contagem antiga) | N/A (docs) | N/A | ✅ TASKS.md | ✅ | ✅ | Prosseguir |

---

## 4. Plano por Task

---

### T-149.DT-086 — Extrair helper `AuditFieldsRowMapper`

- **Critério DONE:** Helper extraído. 4 RowMappers usam helper. Zero duplicação dos 6 campos de auditoria.
- **Estimativa:** 1.5h
- **Papel:** Dev Backend
- **Prioridade:** Could
- **RACI:** Dev Backend (R), Tech Lead (A)

**Abordagem:**
Criar uma classe utilitária `AuditFieldsRowMapper` com um método estático `mapAuditFields(ResultSet rs, BaseEntity entity)` que popula os 6 campos de auditoria (`created_dt`, `updated_dt`, `created_by`, `updated_by`, `deleted_dt`, `deleted_by`) a partir do `ResultSet`. Os 4 RowMappers que mapeiam entidades com campos de auditoria (TenantRowMapper, PlanRowMapper, SubscriptionRowMapper, UserRowMapper) passarão a delegar para este helper em vez de repetir as 6 linhas de mapeamento.

**Evidência da duplicação (código idêntico em 4 arquivos):**

```java
// Padrão repetido em TenantRowMapper:38-43, PlanRowMapper:29-34,
// SubscriptionRowMapper:26-31, UserRowMapper:30-35:
t.setCreatedDt(rs.getObject("created_dt", java.time.OffsetDateTime.class));
t.setUpdatedDt(rs.getObject("updated_dt", java.time.OffsetDateTime.class));
t.setCreatedBy(rs.getObject("created_by", UUID.class));
t.setUpdatedBy(rs.getObject("updated_by", UUID.class));
t.setDeletedDt(rs.getObject("deleted_dt", java.time.OffsetDateTime.class));
t.setDeletedBy(rs.getObject("deleted_by", UUID.class));
```

**RowMappers que NÃO precisam do helper:**
- `AuditEntryRowMapper` — mapeia `audit_log`, que tem estrutura diferente (não é BaseEntity)
- `UserPermissionRowMapper` — mapeia `user_permission`, que não tem campos de auditoria

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../repository/rowmapper/AuditFieldsRowMapper.java` | 🆕 | Helper com método `static void mapAuditFields(ResultSet rs, BaseEntity entity)` |

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../repository/rowmapper/TenantRowMapper.java` | 🔄 | Substituir 6 linhas de auditoria por chamada ao helper |
| `src/main/java/.../repository/rowmapper/PlanRowMapper.java` | 🔄 | Substituir 6 linhas de auditoria por chamada ao helper |
| `src/main/java/.../repository/rowmapper/SubscriptionRowMapper.java` | 🔄 | Substituir 6 linhas de auditoria por chamada ao helper |
| `src/main/java/.../repository/rowmapper/UserRowMapper.java` | 🔄 | Substituir 6 linhas de auditoria por chamada ao helper |

**Dependências:** Nenhuma

**Riscos:**
- Baixo — mudança mecânica. O helper é uma extração de código idêntico. Se o ResultSet não tiver alguma coluna de auditoria, o comportamento é o mesmo de antes (exception do JDBC).
- **Mitigação:** Rodar `mvn test` após a refatoração para garantir que os RowMappers continuam funcionando.

**Skills aplicáveis:** `121-java-object-oriented-design`, `141-java-refactoring-with-modern-features`

**PonteTail — Decisão Rung 2:** O código de mapeamento de auditoria está duplicado 4 vezes (24 linhas idênticas no total). Extrair para helper reduz para 4 chamadas de 1 linha cada. DRY aplicado sem adicionar complexidade.

---

### T-150.DT-089 — Injete `ObjectMapper` do Spring no `AuditAspect`

- **Critério DONE:** `@Autowired ObjectMapper` ou injeção por construtor. Serialização consistente com o resto da aplicação.
- **Estimativa:** 0.5h
- **Papel:** Dev Backend
- **Prioridade:** Could
- **RACI:** Dev Backend (R), Tech Lead (A)

**Abordagem:**
Substituir `private static final ObjectMapper objectMapper = new ObjectMapper()` (linha 40 do AuditAspect.java) por injeção via construtor. O `ObjectMapper` gerenciado pelo Spring já está configurado com os módulos Jackson apropriados (JavaTimeModule, etc.) e é injetado no `SecurityConfig` (linha 62). Usar a mesma instância no `AuditAspect` garante consistência na serialização JSON (ex: formato de datas, nulos).

**Situação atual (linha 40):**
```java
private static final ObjectMapper objectMapper = new ObjectMapper();
```

**Situação desejada:**
```java
private final ObjectMapper objectMapper;

public AuditAspect(JdbcTemplate jdbc, TaskExecutor taskExecutor, ObjectMapper objectMapper) {
    this.jdbc = jdbc;
    this.taskExecutor = taskExecutor;
    this.objectMapper = objectMapper;
}
```

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../security/aspect/AuditAspect.java` | 🔄 | Remover `static final ObjectMapper`, adicionar parâmetro no construtor |

**Dependências:** Nenhuma (Spring gerencia o bean `ObjectMapper` automaticamente via `JacksonAutoConfiguration`)

**Riscos:**
- Baixíssimo — o `ObjectMapper` do Spring é um superconjunto do `new ObjectMapper()` (tem mais módulos registrados). A serialização existente não quebra; apenas fica mais consistente.
- **Mitigação:** Rodar testes que exercem auditoria (`AuditAspectTest`) para confirmar que a serialização permanece funcional.

**Skills aplicáveis:** `301-frameworks-spring-boot-core`, `145-java-refactoring-high-performance`

**PonteTail — Decisão Rung 4:** Spring Boot já provê `ObjectMapper` auto-configurado com `JavaTimeModule` e outras configurações. Criar instância manual bypassa essas configurações e pode produzir JSON inconsistente com os controllers.

---

### T-151.DT-090 — Substituir `OffsetDateTime.now()` → `OffsetDateTime.now(ZoneOffset.UTC)`

- **Critério DONE:** Todos os timestamps em UTC independente do fuso horário da JVM.
- **Estimativa:** 0.5h
- **Papel:** Dev BD
- **Prioridade:** Could
- **RACI:** Dev BD (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Substituir todas as 7 chamadas a `OffsetDateTime.now()` por `OffsetDateTime.now(ZoneOffset.UTC)` em 3 arquivos. Isso garante que os timestamps sejam consistentes independentemente do fuso horário configurado na JVM ou no sistema operacional. Adicionar `import static java.time.ZoneOffset.UTC` para reduzir verbosidade.

> ⚠️ **Por que isso importa:** `OffsetDateTime.now()` usa o fuso horário padrão da JVM (`user.timezone`). Se o ambiente de produção tiver timezone diferente do ambiente de desenvolvimento, os timestamps podem divergir, causando bugs sutis em queries temporais e relatórios.

**Localização de todas as ocorrências:**

| Arquivo | Linha | Contexto |
|:---|:---:|:---|
| `BaseEntity.java` | 33 | `this.createdDt = OffsetDateTime.now();` |
| `BaseEntity.java` | 34 | `this.updatedDt = OffsetDateTime.now();` |
| `BaseEntity.java` | 117 | `this.deletedDt = OffsetDateTime.now();` |
| `BaseRepository.java` | 148 | `OffsetDateTime.now()` em `softDelete()` |
| `BaseRepository.java` | 170 | `OffsetDateTime now = OffsetDateTime.now();` em `save()` |
| `BaseRepository.java` | 225 | `OffsetDateTime now = OffsetDateTime.now();` em `update()` |
| `AuditAspect.java` | 119 | `OffsetDateTime.now()` em `writeAuditLog()` |

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../common/BaseEntity.java` | 🔄 | 3× `OffsetDateTime.now()` → `OffsetDateTime.now(ZoneOffset.UTC)` |
| `src/main/java/.../repository/common/BaseRepository.java` | 🔄 | 3× `OffsetDateTime.now()` → `OffsetDateTime.now(ZoneOffset.UTC)` |
| `src/main/java/.../security/aspect/AuditAspect.java` | 🔄 | 1× `OffsetDateTime.now()` → `OffsetDateTime.now(ZoneOffset.UTC)` |

**Dependências:** Nenhuma

**Riscos:**
- Moderado — se os testes existentes assumirem timezone local em asserções de data/hora, podem falhar após a mudança para UTC.
- **Mitigação:** Rodar `mvn test` completo. Se houver testes que comparam `OffsetDateTime` com timezone local, atualizá-los para usar UTC ou usar `isCloseTo()` com tolerância.

**Skills aplicáveis:** `121-java-object-oriented-design`, `126-java-exception-handling`

**PonteTail — Decisão Rung 5:** O padrão do projeto (ARCHITECTURE.md §6.3, SPECS.md §6.3) define que `created_dt` e `updated_dt` são `TIMESTAMPTZ` — o PostgreSQL armazena tudo em UTC. A aplicação deve emitir timestamps em UTC para consistência banco↔app.

---

### T-152.DT-092 — Bump `springdoc-openapi` 2.8.8→2.8.16

- **Critério DONE:** Swagger UI funcional. `mvn compile` OK. Sem breaking changes.
- **Estimativa:** 0.5h
- **Papel:** DevOps
- **Prioridade:** Could
- **RACI:** DevOps (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Alterar a versão no `<version>` do `pom.xml` de `2.8.8` para `2.8.16`. O springdoc-openapi v2.8.x é compatível com Spring Boot 3.5.x e a mudança é um patch/bugfix — sem breaking changes esperadas. Verificar se o Swagger UI (`/swagger-ui.html`) continua funcional após o bump.

**Situação atual (pom.xml linha 95):**
```xml
<version>2.8.8</version>
```

**Situação desejada:**
```xml
<version>2.8.16</version>
```

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `pom.xml` | 🔄 | Alterar `<version>` de springdoc-openapi de 2.8.8 para 2.8.16 |

**Dependências:** Nenhuma

**Riscos:**
- Baixo — springdoc 2.8.x é linha estável. Bump de patch (8→16) tipicamente contém apenas bugfixes.
- **Mitigação:** Rodar `mvn compile` e verificar que o Swagger UI carrega sem erros no console.

**Skills aplicáveis:** `112-java-maven-plugins`

**PonteTail — Decisão Rung 4:** springdoc-openapi já está declarado no pom.xml. O bump é apenas atualização de versão de uma dependência existente.

---

### T-153.DT-093 — Externalizar CORS origins para `application.yml`

- **Critério DONE:** Origens CORS configuráveis por ambiente via `app.cors.allowed-origins`.
- **Estimativa:** 0.5h
- **Papel:** DevOps
- **Prioridade:** Could
- **RACI:** DevOps (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Extrair a lista hardcoded de origens CORS do `SecurityConfig.java` para `application.yml` usando `@Value` com valor default. Isso segue o princípio 12-Factor App (configuração no ambiente) e permite que dev, staging e produção tenham origens diferentes sem alterar código.

**Situação atual (SecurityConfig.java linhas 187-190):**
```java
configuration.setAllowedOrigins(List.of(
    "http://localhost:3000",   // dev frontend
    "https://app.fbso.org"     // prod frontend
));
```

**Situação desejada — `application.yml`:**
```yaml
app:
  cors:
    allowed-origins: http://localhost:3000,https://app.fbso.org
```

**Situação desejada — `SecurityConfig.java`:**
```java
@Value("${app.cors.allowed-origins:http://localhost:3000,https://app.fbso.org}")
private String allowedOrigins;

// no corsConfigurationSource():
configuration.setAllowedOrigins(
    Arrays.asList(allowedOrigins.split(","))
);
```

**Decisão de design:** Usar `String` com split por vírgula (em vez de `List<String>` com `@Value`) porque Spring Boot não suporta injeção direta de `List<String>` sem `@ConfigurationProperties`. Para uma property simples, split é mais direto e não requer classe de configuração adicional.

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/.../config/SecurityConfig.java` | 🔄 | Substituir `List.of(...)` hardcoded por `@Value` + split |
| `src/main/resources/application.yml` | 🔄 | Adicionar `app.cors.allowed-origins` |

**Dependências:** Nenhuma

**Riscos:**
- Baixo — o `@Value` tem valor default idêntico ao hardcoded atual. Se a property não for definida, comportamento permanece o mesmo.
- **Mitigação:** Rodar `mvn compile` para validar injeção do `@Value`.

**Skills aplicáveis:** `301-frameworks-spring-boot-core`, `304-frameworks-spring-boot-security`

**PonteTail — Decisão Rung 5:** O `application.yml` já segue o padrão Spring Boot de externalização de configuração. Extrair CORS para lá é consistente com JWT, datasource e outras configs que já são externalizadas.

---

### T-154.DT-101 — Atualizar tabela de riscos no SPRINT-CARD.md

- **Critério DONE:** Seção de riscos atualizada com referências corretas aos débitos que os mitigam.
- **Estimativa:** 0.5h
- **Papel:** Tech Lead
- **Prioridade:** Could
- **RACI:** Tech Lead (R+A)

**Abordagem:**
Revisar a tabela de riscos do SPRINT-CARD.md (§⚠️ Riscos e Bloqueadores) e garantir que cada risco referencie corretamente o débito técnico que o trata. A tabela atual já referencia DT-095 (docker-compose), DT-110 (rate limiting), DT-124 (diagrama onboarding), DT-045 (Flyway), DT-068 (PG driver), DT-099+DT-100 (OAuth2). A tarefa consiste em:

1. Confirmar que todas as referências estão corretas pós-execução das Frentes 0 e 1
2. Atualizar o status de mitigação (ex: "Frente 0 concluída ✅" onde aplicável)
3. Adicionar nota sobre Frente 1 (rate limiting implementado T-144, diagramas documentados T-143+T-145)
4. Verificar se novos riscos surgiram durante as Frentes 0-1

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `sprints/sprint-05-portal-cliente/SPRINT-CARD.md` | 🔄 | Atualizar coluna Mitigação com status das Frentes 0-1 |

**Dependências:** T-143, T-144, T-145 (já concluídos na Frente 1)

**Riscos:** Baixo — apenas atualização documental.

**Skills aplicáveis:** Nenhuma (atualização documental)

---

### T-155.DT-112 — Atualizar header do SPECS.md

- **Critério DONE:** Header do SPECS.md reflete o progresso atual da Sprint 5.
- **Estimativa:** 0.25h
- **Papel:** Tech Lead
- **Prioridade:** Could
- **RACI:** Tech Lead (R+A)

**Abordagem:**
Atualizar o campo `Situação implementação:` no header do SPECS.md para refletir o progresso atual: Sprints 1-4 concluídas ✅, Sprint 5 Frentes 0-1 concluídas ✅ (16/40 tasks, 40%), Frente 2 em execução. Ajustar também a linha de status para indicar "Próximo: Sprint 5 Frente 3 — Features Portal do Cliente (16 tarefas)".

**Situação atual (SPECS.md linha 9):**
```
- **Situação implementação:** Em Execução — Sprints 1-4 concluídas ✅.
  Sprint 5 Frentes 0-1 concluídas ✅ (16/40 tasks, 40%).
  18 endpoints REST. 213 testes. RBAC DB-backed com matriz RN10-01
  100% validada. Rate limiting + keyset pagination + JWT converter implementados.
```

**Alteração:** Adicionar referência "Próximo: Sprint 5 Frente 3 — Features Portal do Cliente (16 tarefas: T-057..T-068 + T-157..T-160)" e atualizar status da Frente 2 quando concluída.

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `SPECS.md` | 🔄 | Atualizar `Situação implementação` e adicionar indicação do próximo passo |

**Dependências:** Nenhuma (atualização de baseline)

**Riscos:** Baixo — apenas atualização documental.

**Skills aplicáveis:** Nenhuma (atualização documental)

---

### T-156.DT-113 — Recalcular progresso no TASKS.md

- **Critério DONE:** Métricas de progresso atualizadas e precisas no TASKS.md.
- **Estimativa:** 0.25h
- **Papel:** Tech Lead
- **Prioridade:** Could
- **RACI:** Tech Lead (R+A)

**Abordagem:**
Recalcular o progresso no header do TASKS.md após a conclusão das Frentes 0, 1 e 2 da Sprint 5:

- **Frente 0:** 6 tarefas concluídas ✅
- **Frente 1:** 10 tarefas concluídas ✅
- **Frente 2:** 8 tarefas (esta frente)
- **Frentes 0+1+2:** 6 + 10 + 8 = 24 tarefas de débito técnico
- **Frente 3 (features):** 16 tarefas pendentes
- **Total Sprint 5:** 40 tarefas

**Cálculo do progresso atualizado:**
- Tasks concluídas antes da Sprint 5: 88 (Sprints 1-4)
- Sprint 5 Frente 0: +6 = 94
- Sprint 5 Frente 1: +10 = 104
- Sprint 5 Frente 2 (após concluída): +8 = 112
- Progresso: 112/167 ≈ 67%
- **Atualizar:** header para refletir 104/167 (62%) durante a Frente 2, depois 112/167 (67%) ao final.

**Seções a atualizar no TASKS.md:**

| Seção | Atualização |
|:---|:---|
| Header `Situação implementação` | Atualizar progresso (104/167 → 112/167) |
| Header `Progresso Atual` (§1) | Atualizar contagem e percentual |
| Sprint 5 — Frente 2 (§2) | Marcar status ⬜→✅, atualizar para 8/8 |

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `TASKS.md` | 🔄 | Atualizar header + métricas §1 + status Frente 2 |

**Dependências:** Deve ser a ÚLTIMA task da Frente 2 — reflete o estado final após todas as outras tasks.

**Riscos:** Baixo — apenas atualização de métricas. Consistência cruzada deve ser verificada com SPRINT-CARD.md.

**Skills aplicáveis:** Nenhuma (atualização documental)

---

## 5. Ordem de Execução

Todas as 8 tasks são independentes. A ordem recomendada agrupa por tipo (código vs. documentação) para facilitar a execução:

| # | Task | Papel | Depende de | Paralelizável com |
|:---:|:---|:---|:---|:---|
| 1 | **T-149.DT-086** — AuditFieldsRowMapper | Dev Backend | — | 2, 3, 4, 5 |
| 2 | **T-150.DT-089** — ObjectMapper Injection | Dev Backend | — | 1, 3, 4, 5 |
| 3 | **T-151.DT-090** — OffsetDateTime UTC | Dev BD | — | 1, 2, 4, 5 |
| 4 | **T-152.DT-092** — Springdoc bump | DevOps | — | 1, 2, 3, 5 |
| 5 | **T-153.DT-093** — CORS externalização | DevOps | — | 1, 2, 3, 4 |
| 6 | **T-154.DT-101** — SPRINT-CARD riscos | Tech Lead | — | 7, 8 |
| 7 | **T-155.DT-112** — SPECS.md header | Tech Lead | — | 6, 8 |
| 8 | **T-156.DT-113** — TASKS.md progresso | Tech Lead | 1-7 concluídas | — (última) |

### Justificativa da Ordem:

1. **T-149 a T-153 (código)** podem ser executadas simultaneamente — alteram arquivos diferentes.
2. **T-154 a T-156 (documentação)** podem ser executadas em paralelo com as tasks de código.
3. **T-156 é a última** — consolida o progresso após todas as outras tasks concluídas.

### Diagrama de Gantt (≈2h paralelo):

```
HORA 1                      HORA 2
├──────────────────────────┼──────────────────────────┤
│ T-149 (1.5h) DevBack     │                          │
│ T-150 (0.5h) DevBack     │ T-151 (0.5h) DevBD       │
│ T-152 (0.5h) DevOps      │ T-153 (0.5h) DevOps      │
│ T-154 (0.5h) TL          │ T-155 (0.25h) TL         │
│                          │ T-156 (0.25h) TL          │
└──────────────────────────┴──────────────────────────┘
```

---

## 6. Estratégia de Build e Verificação

### Comandos

| Propósito | Comando |
|:---|:---|
| Compilação | `mvn compile` |
| Testes unitários (completos) | `mvn test` |
| Testes específicos (RowMappers) | `mvn test -Dtest="*RowMapper*"` |
| Testes específicos (AuditAspect) | `mvn test -Dtest="AuditAspectTest"` |
| Swagger UI (verificação manual) | Acessar `http://localhost:8080/swagger-ui.html` após startup |
| Verificação de timestamps (UTC) | `grep -r "OffsetDateTime.now()" src/main/java/` deve retornar VAZIO |
| Verificação de duplicação (audit fields) | `grep -r "setCreatedDt\|setUpdatedDt\|setCreatedBy\|setUpdatedBy\|setDeletedDt\|setDeletedBy" src/main/java/**/rowmapper/` deve retornar APENAS o helper |

### Checkpoints

| # | Momento | O que verificar |
|:---:|:---|:---|
| CP-1 | Após T-149 (helper) | `mvn compile` OK. `mvn test` — RowMapper tests passam. |
| CP-2 | Após T-150+T-151 (ObjectMapper+UTC) | `mvn compile` OK. Zero `OffsetDateTime.now()` sem UTC. |
| CP-3 | Após T-152 (springdoc) | `mvn compile` OK. `mvn dependency:tree` confirma versão 2.8.16. |
| CP-4 | Após T-153 (CORS) | `mvn compile` OK. `@Value` injetado corretamente. |
| CP-5 | Após T-154+T-155 (docs) | Consistência entre SPECS.md, TASKS.md e SPRINT-CARD.md. |
| CP-6 | Final (após T-156) | `mvn test` — todos os 213+ testes passando. Zero regressões. |

### Meta de Qualidade

- **Build:** ✅ `mvn compile` sem erros
- **Testes existentes:** 213+ testes — zero regressões
- **Checkstyle:** Novos arquivos com zero warnings (não agravar as 711 violações existentes — DT-114)
- **UTC:** Zero chamadas a `OffsetDateTime.now()` sem `ZoneOffset.UTC`
- **DRY:** Zero duplicação dos 6 campos de auditoria nos RowMappers

---

## 7. Observações

1. **Todas as tasks são Could Have** — Se o tempo for crítico e a Frente 3 precisar começar, estas tasks podem ser postergadas para a Sprint 6 sem impacto nas features.
2. **T-149 é a maior task** (1.5h) — Extrair o helper e atualizar 4 RowMappers. É uma refatoração mecânica, sem mudança de comportamento.
3. **T-151 (UTC) é a mais crítica** — Embora Could, timestamps inconsistentes entre ambientes são bugs difíceis de detectar. Recomenda-se executar mesmo que outras tasks sejam postergadas.
4. **T-152 (springdoc) e T-153 (CORS)** são tarefas DevOps rápidas (0.5h cada) e podem ser executadas em qualquer momento.
5. **T-156 deve ser a última** — O recálculo de progresso depende de todas as outras tasks estarem concluídas para refletir o estado final correto.
6. **Todas as tasks de documentação (T-154, T-155, T-156)** atualizam artefatos que também serão modificados na Fase 10 (Atualização de Artefatos pós-execução). Coordenar para evitar conflitos de merge.

---

🤖 *Documento gerado em 2026-07-23 como parte da Fase 1 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4. Skills utilizados no planejamento: 121-java-object-oriented-design, 141-java-refactoring-with-modern-features, 301-frameworks-spring-boot-core, 304-frameworks-spring-boot-security, 112-java-maven-plugins.*
