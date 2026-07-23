# SPRINT-DEVELOPMENT-PLANNING-Frente-1.md — Plano de Desenvolvimento: Sprint 5 — Frente 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 1 — Recomendados (🟡 Must + Should)
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Keycloak 26 · Flyway 12.11.0 · Caffeine 3.2.4 · JWT (Nimbus)
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-05-portal-cliente`
- **Data do planejamento:** 2026-07-23

---

## 1. Visão Geral

- **Objetivo da Frente 1:** Resolver 10 débitos técnicos recomendados (Must + Should) antes de iniciar as features da Frente 3. Estes débitos são pré-requisitos ou facilitadores diretos das tarefas T-057 a T-068.
- **Tasks a implementar:** 10 (8 Must + 1 Should + 1 NO-OP)
- **Ordem de execução:** Mista — tasks de documentação e código podem ser paralelizadas. Tasks de BD devem vir antes das tasks de backend que dependem delas.
- **Estimativa total:** ~17.5h (≈3d serial, ≈1.5d paralelo com time completo)
- **Papéis envolvidos:** Dev BD, Tech Lead, Arquiteto, Dev Backend, Especialista IAM, QA Engineer

---

## 2. Dependências entre Tasks

```
                    ┌──────────────────────────────────────────────┐
                    │          T-140.DT-097 (Docs: contagem)       │
                    │          T-143.DT-108 (Docs: máquina estados) │
                    │          T-145.DT-124 (Docs: diag. onboarding)│
                    │          T-147.DT-106 (Docs: +5 cenários)    │
                    │                                              │
                    │  PARALELIZÁVEIS — sem dependências de código │
                    └──────────────────────────────────────────────┘

T-142.DT-107 (BD: is_matrix) ──→ T-139.DT-023 (keyset pagination)
                                      │
                                      ▼
                              T-144.DT-110 (Rate Limit Filter)
                                      │
                                      ▼
                              T-146.DT-121 (ExceptionHandler 401)
                                      │
                                      ▼
                              T-148.DT-102 (JwtConverter — Should)

T-141.DT-098 (TenantContext→JWT) — NO-OP (já implementado)
```

**Legenda:** `──→` = dependência sequencial · `PARALELIZÁVEIS` = sem dependências entre si

---

## 3. Análise PonteTail (7 Rungs) por Task

> **Checklist aplicado antes de planejar cada task.** Se alguma task falhar no Rung 1 (YAGNI), ela é removida do plano. Se passar, segue para estimativa.

### Resumo da Análise

| Task | R1-YAGNI | R2-JáExiste | R3-Stdlib | R4-DepExistente | R5-Padrão | R6-Simples | R7-Mínimo | Veredito |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---|
| T-139.DT-023 | ✅ | ❌ | ✅ JDBC puro | ✅ JdbcTemplate | ✅ BaseRepository | ✅ 1 método | ✅ | Prosseguir |
| T-140.DT-097 | ✅ | ❌ (números errados) | N/A (docs) | N/A | N/A | ✅ | ✅ | Prosseguir |
| T-141.DT-098 | — | ✅ (NO-OP) | — | — | — | — | — | Pular |
| T-142.DT-107 | ✅ | ❌ | ✅ BOOLEAN | ✅ Flyway 12.11.0 | ✅ V001-V006 | ✅ 1 migration | ✅ | Prosseguir |
| T-143.DT-108 | ✅ | ❌ | N/A (docs) | N/A | ✅ ARCHITECTURE.md | ✅ | ✅ | Prosseguir |
| T-144.DT-110 | ✅ | ❌ | ✅ Caffeine 3.2.4 | ✅ Caffeine no pom.xml | ✅ Filter pattern | ✅ 1 Filter | ✅ | Prosseguir |
| T-145.DT-124 | ✅ | ❌ | N/A (docs) | N/A | ✅ ARCHITECTURE.md | ✅ | ✅ | Prosseguir |
| T-146.DT-121 | ✅ | ❌ | ✅ Spring Security | ✅ spring-boot-starter-security | ✅ GlobalExceptionHandler | ✅ 1 handler | ✅ | Prosseguir |
| T-147.DT-106 | ✅ | ❌ | N/A (docs) | N/A | ✅ SPRINT-TEST-SUITE.md | ✅ | ✅ | Prosseguir |
| T-148.DT-102 | ✅ | ❌ (parcial — DT-076) | ✅ JwtAuthenticationConverter | ✅ spring-boot-starter-oauth2-resource-server | ✅ SecurityConfig | ✅ 1 converter | ✅ | Prosseguir |

---

## 4. Plano por Task

---

### T-139.DT-023 — Implementar `findAllKeyset()` no `BaseRepository`

- **Critério DONE:** Keyset funcional. Testes com >10k registros.
- **Estimativa:** 3h
- **Papel:** Dev BD
- **Prioridade:** Must
- **RACI:** Dev BD (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Adicionar método `findAllKeyset(UUID lastId, int pageSize, String sortColumn)` ao `BaseRepository` usando o padrão SQL `WHERE id > :lastId ORDER BY id LIMIT :size`. Este padrão evita o problema de performance da paginação offset com grandes volumes (>10k registros), pois o índice B-tree no `id` (chave primária) é usado diretamente, sem precisar pular linhas.

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/com/fbso/platform/admin/repository/common/BaseRepository.java` | 🔄 | Adicionar método `findAllKeyset()` com sanitização de coluna e suporte a tenant filter. Reutilizar `sanitizeColumn()`, `tenantClause()`, `buildParams()` existentes. |

**Arquivos a criar:** Nenhum

**Dependências:** Nenhuma (usa infraestrutura existente do `BaseRepository`)

**Riscos:**
- Keyset pagination exige que a coluna de ordenação tenha índice. A coluna `id` (PK) já é indexada. Se outras colunas forem usadas para ordenação, verificar existência de índice.
- Compatibilidade com `findAll()` existente — não substituir, adicionar como método complementar.

**Skills aplicáveis:** `311-frameworks-spring-jdbc`, `121-java-object-oriented-design`

**PonteTail — Decisão Rung 3:** A biblioteca padrão (JDBC + Spring JdbcTemplate) cobre totalmente. Não adicionar dependência externa.

---

### T-140.DT-097 — Corrigir contagem de cenários: 21→28

- **Critério DONE:** Números consistentes entre SPRINT-CARD.md e SPRINT-TEST-SUITE.md
- **Estimativa:** 0.5h
- **Papel:** Tech Lead
- **Prioridade:** Must
- **RACI:** Tech Lead (R+A)

**Abordagem:**
Atualizar duas referências numéricas: (1) `SPRINT-CARD.md` — métricas da sprint declaram 21 cenários, corrigir para 28; (2) `SPRINT-TEST-SUITE.md` — cabeçalho declara 21, corpo tem 28 (soma: 9+10+5+4=28).

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `sprints/sprint-05-portal-cliente/SPRINT-CARD.md` | 🔄 | Atualizar métrica "Cenários de teste" de 21→28 |
| `sprints/sprint-05-portal-cliente/SPRINT-TEST-SUITE.md` | 🔄 | Atualizar cabeçalho de 21→28 |
| `sprints/sprint-05-portal-cliente/README.md` | 🔄 | Atualizar referência de 28→33 (após T-147) |

**Dependências:** T-147.DT-106 (adição de +5 cenários) — executar T-140 ANTES de T-147 para ter baseline correta

**Riscos:** Baixo — apenas correção numérica.

**Skills aplicáveis:** Nenhuma (correção documental)

---

### T-141.DT-098 — Conectar TenantContext ao JwtAuthenticationFilter

- **Status:** ✅ **NO-OP** — Código já implementado
- **Critério DONE:** Já atendido — `TenantContext` armazena e expõe `getBusinessUnitIds()` e `getModules()`. Getters são chamados em T-065 e T-066.
- **Evidência:** `TenantContext.java` (linhas 42-58) — campos `businessUnitIds` e `modules` populados desde a Sprint 4. `JwtAuthenticationFilter.java` — extrai claims via `JwtUtils.getBusinessUnitIds()` e `JwtUtils.getModules()` (6 testes passando).

**Economia:** ~1.5h (NO-OP). Esta task foi identificada pela auditoria antes da inspeção detalhada do código.

---

### T-142.DT-107 — Adicionar campo `is_matrix` à entidade `BusinessUnit`

- **Critério DONE:** Migration V007 aplicada. Entidade `BusinessUnit` atualizada com campo `is_matrix`. RowMapper atualizado. T-062 pode usar a flag.
- **Estimativa:** 1h
- **Papel:** Dev BD
- **Prioridade:** Must
- **RACI:** Dev BD (R), Dev Backend (C), Arquiteto (C), Tech Lead (A)

**Abordagem:**
Criar migration Flyway V007 que adiciona coluna `is_matrix BOOLEAN NOT NULL DEFAULT false` à tabela `fbso_platform.business_unit`. Atualizar a entidade `BusinessUnit.java` com o novo campo e getter/setter. Atualizar `toColumnMap()` para incluir o campo. Atualizar o `BusinessUnitRowMapper` (se existir) ou o mapeamento inline.

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/resources/db/migration/V007__add_is_matrix_to_business_unit.sql` | 🆕 | `ALTER TABLE fbso_platform.business_unit ADD COLUMN is_matrix BOOLEAN NOT NULL DEFAULT false` |
| `src/main/java/com/fbso/platform/admin/entity/BusinessUnit.java` | 🔄 | Adicionar campo `private boolean isMatrix` + getter/setter. Atualizar `toColumnMap()`. |
| `src/main/java/com/fbso/platform/admin/repository/rowmapper/UserPermissionRowMapper.java` | 🔄 | Se o RowMapper de BusinessUnit existir, adicionar mapeamento da coluna `is_matrix` |

**Arquivos a criar:** V007 migration

**Dependências:** T-134.DT-045 (Flyway 12.11.0 — ✅ concluído na Frente 0)

**Riscos:**
- `DEFAULT false` garante que BUs existentes não sejam marcadas incorretamente como matriz.
- Verificar se há outras queries ou services que precisam ser atualizados para usar `is_matrix` em vez de `parent_id IS NULL`.

**Skills aplicáveis:** `313-frameworks-spring-db-migrations-flyway`, `311-frameworks-spring-jdbc`

**PonteTail — Decisão Rung 2:** A lógica atual usa `parent_id=NULL` para identificar a Matriz. Isso é frágil — qualquer BU sem parent pode ser confundida. A flag explícita `is_matrix` é semanticamente clara e evita ambiguidade.

---

### T-143.DT-108 — Documentar máquina de estados de `TenantStatus`

- **Critério DONE:** Diagrama de estados documentado no ARCHITECTURE.md. Validação de transições implementada no `OnboardingService` (T-060).
- **Estimativa:** 1.5h
- **Papel:** Arquiteto de Solução
- **Prioridade:** Must
- **RACI:** Arquiteto (R), Dev Backend (C), Dev Full-Stack (C), Tech Lead (A)

**Abordagem:**
Documentar a máquina de estados finita (FSM) para `TenantStatus` no `ARCHITECTURE.md`. O enum atual tem: `PENDING_SETUP`, `PENDING_ONBOARDING`, `ACTIVE`, `SUSPENDED`, `CANCELED`. As transições válidas são:

```
PENDING_SETUP ──→ PENDING_ONBOARDING (admin conclui setup inicial)
PENDING_ONBOARDING ──→ ACTIVE (4 passos onboarding concluídos — RN14-04)
PENDING_ONBOARDING ──→ SUSPENDED (admin suspende durante onboarding)
ACTIVE ──→ SUSPENDED (admin suspende tenant ativo)
SUSPENDED ──→ ACTIVE (admin reativa tenant)
ACTIVE ──→ CANCELED (admin cancela tenant)
SUSPENDED ──→ CANCELED (admin cancela tenant suspenso)
```

Implementar validação no `OnboardingService` (ou em um `TenantStateValidator`) que rejeita transições inválidas lançando `InvalidStatusTransitionException` (já existe no codebase).

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `ARCHITECTURE.md` | 🔄 | Adicionar seção "Máquina de Estados — TenantStatus" com diagrama Mermaid stateDiagram |
| `src/main/java/com/fbso/platform/admin/enums/TenantStatus.java` | 🔄 | (Opcional) Adicionar método `canTransitionTo(TenantStatus target): boolean` |
| `src/main/java/com/fbso/platform/admin/service/OnboardingService.java` | 🔄 | Validar transições usando `TenantStatus.canTransitionTo()` antes de mudar status |

**Arquivos a criar:** Nenhum (documentação no ARCHITECTURE.md existente)

**Dependências:** Nenhuma (tarefa de design/documentação)

**Riscos:**
- Se a validação for colocada no `OnboardingService`, ela só cobre transições durante onboarding. Transições fora do fluxo de onboarding (ex: admin suspendendo tenant) precisam de validação em outro lugar. Considerar validador centralizado.

**Skills aplicáveis:** `030-architecture-adr-general`, `033-architecture-diagrams`

**PonteTail — Decisão Rung 5:** O `ARCHITECTURE.md` já define o padrão para documentação de design (§2 do prompt). O diagrama usa Mermaid (padrão do projeto desde a consolidação C4 v2.0).

---

### T-144.DT-110 — Implementar rate limiting via Filter + Caffeine

- **Critério DONE:** Filter funcional. Mensagem exibe tempo restante de bloqueio. Testes passam. Parte integrante de T-059.
- **Estimativa:** 3h
- **Papel:** Dev Backend
- **Prioridade:** Must
- **RACI:** Dev Backend (R), Arquiteto (C), Tech Lead (A)

**Abordagem:**
Criar `RateLimitFilter` que intercepta requisições para `POST /api/v1/auth/login`. Usar Caffeine Cache (já no classpath — Sprint 4) com política `expireAfterWrite(15, TimeUnit.MINUTES)`. Chave do cache: `username` (extraído do body ou IP se username não disponível). Contador de tentativas. Na 5ª falha → entrada no cache com tempo de expiração. Resposta: `429 Too Many Requests` + RFC 7807 body com campo `detail` informando tempo restante.

**Decisão de design (DT-110):** Filter (não @Aspect). Justificativa: (1) Filter é mais simples e idiomático no Spring Security; (2) Rate limiting é uma preocupação transversal de infraestrutura, não de negócio — Filter é o lugar correto; (3) Trigger para migrar para Redis: quando `INSTANCE_COUNT > 1` em produção (risco documentado no SPRINT-CARD.md).

**Arquivos a criar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/com/fbso/platform/admin/security/RateLimitFilter.java` | 🆕 | Filter com Caffeine Cache. Chave: IP ou username. 5 tentativas → bloqueio 15min. Resposta RFC 7807. |

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/com/fbso/platform/admin/config/SecurityConfig.java` | 🔄 | Adicionar `RateLimitFilter` ANTES do `JwtAuthenticationFilter` no `apiFilterChain` para `POST /auth/login` |
| `src/test/java/com/fbso/platform/admin/security/RateLimitFilterTest.java` | 🆕 | Testes: 5 falhas → bloqueio, reset após 15min, mensagem com tempo restante |

**Dependências:** Nenhuma (Caffeine já está no pom.xml desde a Sprint 4 Frente 0)

**Riscos:**
- Em cluster, o cache Caffeine é local (por instância). Um atacante pode distribuir tentativas entre instâncias. Mitigação: Fase 1 single-instance. Trigger documentado para Redis.
- O Filter precisa ler o body da requisição para extrair username, o que consome o `InputStream`. Usar `ContentCachingRequestWrapper` ou extrair username do `HttpServletRequest.getParameter()`.

**Skills aplicáveis:** `304-frameworks-spring-boot-security`, `124-java-secure-coding`

**PonteTail — Decisão Rung 4:** Caffeine 3.2.4 já está declarado no `pom.xml` desde a Sprint 4 (T-133.DT-095). Não adicionar nova dependência.

---

### T-145.DT-124 — Criar diagrama de estados do onboarding

- **Critério DONE:** Diagrama documentado no ARCHITECTURE.md. 4 estados + transições + edge cases cobertos.
- **Estimativa:** 2h
- **Papel:** Arquiteto de Solução
- **Prioridade:** Must
- **RACI:** Arquiteto (R), Dev Backend (C), Dev Full-Stack (C), Dev Frontend (C), Tech Lead (A)

**Abordagem:**
Documentar o diagrama de estados do fluxo de onboarding antes da implementação do `OnboardingService` (T-060). Os 4 passos são: Step-1 (Dados do Tenant), Step-2 (CNPJ + Regime Tributário → cria 1ª BU Matriz), Step-3 (Configurações Fiscais), Step-4 (Confirmação → Tenant → ACTIVE). O onboarding é retomável: o tenant pode sair e voltar, retomando do último passo concluído.

Estados do onboarding (distintos do `TenantStatus`):
```
NOT_STARTED → STEP_1_DONE → STEP_2_DONE → STEP_3_DONE → COMPLETED
```

Edge cases a documentar:
1. Tenant sai no meio do step 2 → retoma step 2
2. Step 2 falha (CNPJ inválido) → permanece em STEP_1_DONE
3. Tentar pular step (ex: acessar step 3 sem step 2) → 422
4. Retomar após 30 dias de inatividade → sem expiração (Fase 0)
5. Admin pode resetar onboarding? → Não na Fase 0

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `ARCHITECTURE.md` | 🔄 | Adicionar seção "Máquina de Estados — Onboarding" com diagrama Mermaid e tabela de edge cases |

**Arquivos a criar:** Nenhum (documentação inline no ARCHITECTURE.md)

**Dependências:** T-143.DT-108 (máquina de estados TenantStatus) — compartilham o mesmo destino no ARCHITECTURE.md. Podem ser feitas em sequência ou em paralelo (seções diferentes).

**Riscos:**
- Se o diagrama não capturar um edge case, o `OnboardingService` (T-060) pode implementar comportamento incorreto. Revisão cruzada com Dev Backend é essencial.

**Skills aplicáveis:** `030-architecture-adr-general`, `033-architecture-diagrams`

---

### T-146.DT-121 — Adicionar `@ExceptionHandler(AuthenticationException.class)` → 401

- **Critério DONE:** 401 JSON padronizado (RFC 7807). Token inválido/expirado não gera 500.
- **Estimativa:** 0.5h
- **Papel:** Dev Backend
- **Prioridade:** Must
- **RACI:** Dev Backend (R), Especialista IAM (C), Tech Lead (A)

**Abordagem:**
Adicionar handler no `GlobalExceptionHandler` para `org.springframework.security.core.AuthenticationException`. Atualmente, `AuthenticationException` não tem handler específico e cai no `handleGenericException()` (500). O handler deve retornar 401 com body RFC 7807 padronizado, seguindo o mesmo formato dos handlers existentes.

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/com/fbso/platform/admin/exception/GlobalExceptionHandler.java` | 🔄 | Adicionar `@ExceptionHandler(AuthenticationException.class)` → 401 RFC 7807. Manter `handleSecurityException` existente para outros casos. |

**Dependências:** Nenhuma

**Riscos:** Baixo — adição de um handler. Verificar precedência com `handleSecurityException` (que cobre `SecurityException`) e `handleSpringAccessDenied` (que cobre `AccessDeniedException`). `AuthenticationException` é uma subclasse diferente — não há conflito.

**Skills aplicáveis:** `304-frameworks-spring-boot-security`, `126-java-exception-handling`

**PonteTail — Decisão Rung 2:** O `GlobalExceptionHandler` já tem 10 handlers seguindo o mesmo padrão RFC 7807. Adicionar mais um segue o padrão existente.

---

### T-147.DT-106 — Adicionar 5 cenários de teste ausentes

- **Critério DONE:** 33 cenários no total (28+5). Suite atualizada no SPRINT-TEST-SUITE.md.
- **Estimativa:** 2h
- **Papel:** QA Engineer
- **Prioridade:** Must
- **RACI:** QA Engineer (R), Analista de Homologação (C), Tech Lead (A)

**Abordagem:**
Adicionar 5 cenários de teste identificados como ausentes pela auditoria (DT-106):
1. **Timeout de sessão (60min)** — RN13-02: Após 60min de inatividade, token expira e próximo request retorna 401.
2. **Validação complexidade senha (RN13-01)** — Senha deve ter 8+ caracteres, letra+número. Testar boundary: 7 chars, 8 chars sem número, 8 chars válidos.
3. **Passo 3 do onboarding** — Cenário atualmente inexistente. Cobrir: sucesso, falha, retomada.
4. **Segurança F04-03 (Dashboard Cliente)** — Isolamento entre tenants: cliente tenant A não vê dados do tenant B.
5. **Segurança F04-03 (Dashboard Cliente)** — Usuário sem token acessa dashboard → 401.

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `sprints/sprint-05-portal-cliente/SPRINT-TEST-SUITE.md` | 🔄 | Adicionar 5 cenários. Atualizar resumo: 28→33. |
| `sprints/sprint-05-portal-cliente/SPRINT-CARD.md` | 🔄 | Atualizar métricas: 28→33 cenários |

**Dependências:** T-140.DT-097 (corrigir contagem baseline 21→28 primeiro)

**Riscos:** Baixo — apenas documentação de cenários. A implementação dos testes ocorrerá na Fase 3-4.

**Skills aplicáveis:** `133-java-testing-acceptance-tests`

---

### T-148.DT-102 — Consolidar dupla decodificação JWT via `JwtAuthenticationConverter`

- **Critério DONE:** 1 decodificação por request. Testes de performance passam.
- **Estimativa:** 2h
- **Papel:** Especialista IAM
- **Prioridade:** Should
- **RACI:** Especialista IAM (R), Dev Backend (C), Tech Lead (A)

**Abordagem:**
Atualmente, cada requisição decodifica o JWT duas vezes: (1) `BearerTokenAuthenticationFilter` do Spring Security (Resource Server) valida o token e popula `JwtAuthenticationToken`; (2) `JwtAuthenticationFilter` customizado extrai claims (`tenant_id`, `roles`, `modules[]`, `business_unit_ids[]`) do token bruto (`request.getHeader("Authorization")`). A consolidação usa `JwtAuthenticationConverter` — um converter customizado que, durante a primeira decodificação (feita pelo Spring), já extrai todas as claims e popula o `TenantContext`. O `JwtAuthenticationFilter` customizado é então simplificado ou removido.

**Estratégia (baseada no DT-076 — Sprint 4):**
1. Criar `FbsoJwtAuthenticationConverter` que implementa `Converter<Jwt, AbstractAuthenticationToken>`
2. No `convert()`, extrair claims (`tenant_id`, `roles`, `modules[]`, `business_unit_ids[]`) do objeto `Jwt` (já decodificado pelo Spring)
3. Popular `TenantContext` com os valores extraídos
4. Configurar o converter no `SecurityConfig.apiFilterChain()`: `.jwt(jwt -> jwt.jwtAuthenticationConverter(fbsoJwtAuthenticationConverter))`
5. Remover ou simplificar `JwtAuthenticationFilter` (não precisa mais re-extrair do header)

**Arquivos a modificar:**

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/java/com/fbso/platform/admin/security/FbsoJwtAuthenticationConverter.java` | 🆕 | Converter que extrai claims do `Jwt` decodificado e popula `TenantContext` |
| `src/main/java/com/fbso/platform/admin/config/SecurityConfig.java` | 🔄 | Configurar `.jwtAuthenticationConverter(fbsoJwtAuthenticationConverter)` no Resource Server |
| `src/main/java/com/fbso/platform/admin/security/JwtAuthenticationFilter.java` | 🔄 | Simplificar ou remover extração duplicada de claims |

**Dependências:** Nenhuma técnica, mas é Should (não Must) — pode ser postergada se o tempo for crítico.

**Riscos:**
- `JwtAuthenticationFilter` atualmente faz mais do que extrair claims (ex: validação adicional). Refatoração deve preservar toda a lógica existente.
- Impacto nos 6 testes existentes de `JwtAuthenticationFilterTest` — devem ser atualizados para refletir a nova arquitetura.

**Skills aplicáveis:** `304-frameworks-spring-boot-security`, `145-java-refactoring-high-performance`

**PonteTail — Decisão Rung 2:** O Spring Security já fornece `JwtAuthenticationConverter`. O código atual duplica a decodificação. Consolidar reduz CPU e complexidade.

---

## 5. Ordem de Execução

A ordem recomendada considera dependências técnicas e a capacidade de paralelização com o time técnico v1.5:

| # | Task | Papel | Depende de | Paralelizável com |
|:---:|:---|:---|:---|:---|
| 1 | **T-140.DT-097** — Corrigir contagem cenários | Tech Lead | — | Todas as outras |
| 2 | **T-143.DT-108** — Máquina estados TenantStatus | Arquiteto | — | 3, 4, 5, 7 |
| 3 | **T-145.DT-124** — Diagrama onboarding | Arquiteto | 2 (seções diferentes no mesmo doc) | 4, 5, 7 |
| 4 | **T-142.DT-107** — Campo is_matrix + V007 | Dev BD | — | 2, 3, 5, 7 |
| 5 | **T-147.DT-106** — +5 cenários teste | QA Engineer | 1 (baseline corrigida) | 2, 3, 4, 7 |
| 6 | **T-139.DT-023** — Keyset pagination | Dev BD | 4 (V007 pode rodar antes, mas keyset é independente) | — |
| 7 | **T-146.DT-121** — ExceptionHandler 401 | Dev Backend | — | 2, 3, 4, 5 |
| 8 | **T-144.DT-110** — Rate Limit Filter | Dev Backend | — (Caffeine já no classpath) | — |
| 9 | **T-148.DT-102** — JwtConverter (Should) | Especialista IAM | — | — |
| — | **T-141.DT-098** — NO-OP | — | — | — |

### Justificativa da Ordem:

1. **T-140 primeiro** porque é rápido (30min) e corrige a baseline numérica que T-147 depende.
2. **T-143 + T-145** podem rodar em paralelo com **T-142 + T-146 + T-147** — documentação vs. código.
3. **T-139** (keyset) e **T-144** (rate limit) são independentes entre si e podem rodar em paralelo após as tasks de documentação.
4. **T-148** (JwtConverter) é Should — executa por último, se houver capacidade do Especialista IAM.

### Diagrama de Gantt (≈1.5d com time completo):

```
DIA 1 (manhã)          DIA 1 (tarde)           DIA 2 (manhã)
├──────────────────────┼───────────────────────┼──────────────────────┤
│ T-140 (0.5h) TL     │ T-147 (2h) QA         │ T-144 (3h) DevBack   │
│ T-143 (1.5h) Arq    │ T-145 (2h) Arq        │ T-148 (2h) IAM       │
│ T-142 (1h) DevBD    │ T-139 (3h) DevBD      │                      │
│ T-146 (0.5h) DevBack│                       │                      │
└──────────────────────┴───────────────────────┴──────────────────────┘
```

---

## 6. Estratégia de Build e Verificação

### Comandos

| Propósito | Comando |
|:---|:---|
| Compilação | `mvn compile` |
| Testes unitários | `mvn test` |
| Testes específicos (rate limit) | `mvn test -Dtest="RateLimitFilterTest"` |
| Testes específicos (keyset) | `mvn test -Dtest="BaseRepositoryTest"` |
| Migração Flyway | `mvn flyway:migrate` |
| Validação migration | `mvn flyway:validate` |
| Verificação dependências | `mvn dependency:tree` |

### Checkpoints

| # | Momento | O que verificar |
|:---:|:---|:---|
| CP-1 | Após T-142 (V007) | `mvn flyway:migrate` executa V007 sem erro. `mvn compile` OK. |
| CP-2 | Após T-139 (keyset) | `mvn compile` OK. `BaseRepositoryTest` passa. |
| CP-3 | Após T-144 + T-146 | `mvn compile` OK. `RateLimitFilterTest` + `GlobalExceptionHandler` OK. |
| CP-4 | Após T-148 (JwtConverter) | `mvn test` — todos os 213+ testes passando. `JwtAuthenticationFilterTest` (6) intactos. |
| CP-5 | Final da Frente 1 | `mvn test` completo. Zero regressões. |

### Meta de Qualidade

- **Build:** ✅ `mvn compile` sem erros
- **Testes existentes:** 213+ testes — zero regressões
- **Novos testes:** T-139 (keyset) e T-144 (rate limit) — cobertura ≥ 80%
- **Checkstyle:** Novos arquivos com zero warnings (não agravar as 711 violações existentes — DT-114)

---

## 7. Observações

1. **T-141 NO-OP** — Documentado no plano para rastreabilidade, mas não consome tempo de execução.
2. **Paralelismo real depende do time preenchido** — O SPRINT-5-TEAM-ALLOCATION.md indica que todos os 10 papéis ainda estão com `<TODO>`. Sem time preenchido, a execução serial leva ~3d.
3. **T-148 é Should** — Se o Especialista IAM estiver sobrecarregado (T-057 é prioridade máxima), esta task pode ser movida para a Frente 2 ou postergada para Sprint 6.
4. **T-143 + T-145 são pré-requisitos de design para T-060** — O `OnboardingService` (Frente 3) depende destes diagramas para implementar corretamente as transições de estado.
5. **Caffeine já está no classpath** desde a Sprint 4 Frente 0 — T-144 não requer nova dependência.

---

🤖 *Documento gerado em 2026-07-23 como parte da Fase 1 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 + Caffeine 3.2.4. Skills utilizados no planejamento: 311-frameworks-spring-jdbc, 304-frameworks-spring-boot-security, 313-frameworks-spring-db-migrations-flyway, 121-java-object-oriented-design, 126-java-exception-handling, 145-java-refactoring-high-performance, 030-architecture-adr-general, 033-architecture-diagrams.*
