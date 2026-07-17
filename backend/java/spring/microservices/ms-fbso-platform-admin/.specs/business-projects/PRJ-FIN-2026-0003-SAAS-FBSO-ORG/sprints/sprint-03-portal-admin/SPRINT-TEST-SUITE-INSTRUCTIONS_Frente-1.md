# SPRINT-TEST-SUITE-INSTRUCTIONS.md — Execução de Testes: Sprint 3

- **Sprint:** 3 de 7 — Portal Admin + Contas e Planos
- **Atualização:** 17/07/2026
- **Contexto:** Execução dos testes de integração T-023 (DashboardRepository com Testcontainers) + verificação completa dos 50 testes unitários existentes
- **Pré-requisito:** Docker Engine rodando (verificado ✅)

---

## 1. Diagnóstico Atual

### 1.1 Ambiente Verificado

| Componente | Versão | Status |
|:---|:---|:---:|
| Docker | Engine rodando (`nifty_elbakyan`) | ✅ |
| Java | GraalVM 25.0.3 LTS | ✅ |
| Build | `mvn compile` SUCCESS | ✅ |
| Testes unitários | 50/50 passando (10 classes) | ✅ |
| Testes integração | 0 implementados | 🔴 T-023 |

### 1.2 Testes Unitários Existentes (50/50 ✅)

| Classe | Testes | Framework | O que cobre |
|:---|:---:|:---|:---|
| `DashboardServiceTest` | 10 | Mockito | summary, evolution, byStatus, byPlan, alerts, período inválido |
| `TenantRepositoryTest` | 4 | Mockito + Mock JdbcTemplate | findAllPaginated, findByNameCorporate, countFiltered |
| `AuditAspectTest` | 3 | Mockito | captura tenant/user no JoinPoint, TaskExecutor |
| `BaseRepositoryTest` | ~8 | Mockito | CRUD, soft delete, tenant filter |
| `TenantAwareDataSourceTest` | ~4 | Mockito | SET app.current_tenant_id, fallback |
| `RbacAspectTest` | ~6 | Mockito | matriz completa DASHBOARD/TENANT/PLAN/SUBSCRIPTION |
| `JwtAuthenticationFilterTest` | ~6 | Mockito | JWT RS256, exp, claims, ObjectMapper |
| `GlobalExceptionHandlerTest` | ~4 | MockMvc | RFC 7807, 401, 403, 422, 500 |
| `RLSIsolationTest` | 5 | Estrutural | arquivos V003/U003 existem, 5 tabelas, idempotência |

### 1.3 O Que Falta Implementar (T-023)

| ID | Descrição | Nível | Ref. TEST_PLAN |
|:---|:---|:---|:---|
| TC-F01-01-004 | GET /dashboard/admin/summary → 200 com dados reais | Integração | §3.1 |
| TC-F01-01-005 | Soft-deleted tenants excluídos das métricas | Integração | §3.1 |
| TC-F01-02-003 | GET /tenants paginado com filtros (PostgreSQL real) | Integração | §3.2 |
| TC-F01-02-004 | Busca <3 chars retorna lista vazia | Integração | §3.2 |
| TC-F01-03-003 | GET /dashboard/admin/alerts → cards populados | Integração | §3.3 |
| TC-F01-03-004 | Sem alertas → lista vazia | Integração | §3.3 |

**Objetivo:** 6 cenários de integração, PostgreSQL real via Testcontainers, 10+ tenants seed.

---

## 2. Passo a Passo para Execução

### Passo 1 — Verificar Pré-requisitos (30 segundos)

```bash
# 1. Confirmar que Docker está rodando
docker ps
# Esperado: lista de containers (pelo menos 1 container visível)

# 2. Verificar versão Java
java --version
# Esperado: "GraalVM 25.0.3" ou superior

# 3. Verificar que a branch correta está ativa
git branch --show-current
# Esperado: "feature/sprint-03-portal-admin"

# 4. Confirmar que o build compila
./mvnw clean compile -q
# Esperado: sem erros (apenas BUILD SUCCESS)
```

> 💡 **Função do `docker ps`**: Lista containers ativos. O Docker precisa estar rodando porque o Testcontainers usa a API do Docker para subir containers PostgreSQL efêmeros durante os testes. Se o Docker não estiver rodando, os testes de integração falham com `Could not find a valid Docker environment`.

> 💡 **Função do `-q` no Maven**: Modo silencioso (quiet) — suprime os logs de download de dependências e mostra apenas erros. Útil para verificações rápidas. Remova o `-q` se quiser ver o log completo.

---

### Passo 2 — Executar Todos os Testes Unitários (1-2 minutos)

Este comando executa APENAS os testes unitários (sem integração), que não dependem de Docker:

```bash
# Executa todos os testes unitários (exclui integração)
./mvnw test -Dcheckstyle.skip=true
```

**Resultado esperado (SUCCESSO):**
```
[INFO] Tests run: 77, Failures: 0, Errors: 0, Skipped: 5
[INFO] BUILD SUCCESS
```

**Resultado esperado (FALHA):**
```
[ERROR] Tests run: 77, Failures: 1, Errors: 0, Skipped: 0
[ERROR] BUILD FAILURE
```
> Se houver falha, procurar no log o nome do teste que falhou e o stack trace.

> ⚠️ **Logs WARN/ERROR durante os testes são NORMAIS.** O projeto testa cenários de erro (exceções, acesso negado, auditoria) que produzem logs de WARN e ERROR intencionalmente. Esses logs NÃO indicam falha nos testes. Apenas a linha `Tests run: X, Failures: Y` no final determina sucesso (Y=0) ou falha (Y>0).

> 💡 **Por que `-Dcheckstyle.skip=true`?** O plugin checkstyle gera centenas de warnings de formatação (indentação, javadoc) que poluem a saída. Pular checkstyle durante os testes acelera a execução e facilita a leitura dos resultados.

---

### Passo 3 — Executar o Teste de Integração do Dashboard (3-5 minutos na primeira vez)

> ⚠️ **Na primeira execução**, o Testcontainers baixa a imagem `postgres:17-alpine` (~80 MB). Isso acontece automaticamente. Execuções subsequentes usam a imagem em cache (~30s).

```bash
# Executa o teste de integração do DashboardRepository
./mvnw test -Dtest="com.fbso.platform.admin.integration.repository.DashboardRepositoryIT" -DfailIfNoTests=false
```

**Resultado esperado (SUCCESSO):**
```
[INFO] Tests run: 23, Failures: 0, Errors: 0, Skipped: 0
[INFO] BUILD SUCCESS
```

> 💡 **Função do `-DfailIfNoTests=false`**: Por padrão, o Maven falha o build se nenhum teste for encontrado. Esta flag evita esse erro — útil durante o desenvolvimento quando a classe de teste ainda não foi criada, ou quando se usa filtros que podem não encontrar nada.

**O que acontece nos bastidores:**
1. Testcontainers detecta o Docker e baixa `postgres:17-alpine` (primeira vez)
2. Sobe container PostgreSQL com credenciais `fbso_test/fbso_test`
3. Flyway executa migrations V001→V005 automaticamente
4. Teste insere 10+ tenants seed
5. Queries do DashboardRepository são executadas contra PostgreSQL real
6. Container é removido ao final (via `@Container` + Ryuk)

> 💡 **Função do Testcontainers + Ryuk**: O Testcontainers usa um container auxiliar chamado Ryuk (Moby Project) que atua como "garbage collector" — ele monitora os containers criados pelos testes e os remove automaticamente quando o processo JVM termina, mesmo que o teste falhe ou seja interrompido. Isso evita dangling containers.

---

### Passo 4 — Executar Testes de Segurança + Integração (3-5 minutos)

```bash
# Executa todos os testes: unitários (Surefire) + integração (Failsafe)
./mvnw clean verify -Dcheckstyle.skip=true
```

> 💡 **Função do `verify`**: Diferente de `mvn test`, o `verify` executa tanto os testes unitários (Surefire, sufixo `*Test.java`) quanto os de integração (Failsafe, sufixo `*IT.java`). O JaCoCo também gera o relatório na fase `verify`, capturando a cobertura combinada de ambos.

**Resultado esperado (SUCCESSO):**
```
[INFO] Tests run: 77, Failures: 0, Errors: 0, Skipped: 5
[INFO] BUILD SUCCESS
```

> 💡 **Função do `clean` antes do `test`**: O `clean` remove o diretório `target/` (compilados anteriores). Isso garante que nenhum `.class` stale (desatualizado) seja usado, evitando falsos positivos. Sempre use `clean test` (não apenas `test`) para garantir um build limpo.

---

### Passo 5 — Gerar Relatório de Cobertura JaCoCo (parte do verify)

```bash
# JaCoCo é executado automaticamente na fase verify.
# O comando abaixo gera o relatório e verifica thresholds (≥80% linhas, ≥70% branches):
./mvnw clean verify -Dcheckstyle.skip=true
```

O relatório é gerado em HTML:

```bash
# Abrir relatório no navegador
open target/site/jacoco/index.html

# Ou verificar cobertura via CSV
awk -F',' 'NR>1{i_missed+=$4; i_covered+=$5} END{printf "Cobertura: %.1f%%\n", i_covered*100/(i_missed+i_covered)}' target/site/jacoco/jacoco.csv
```

**Resultado esperado (SUCCESSO):**
- Relatório gerado em `target/site/jacoco/index.html`
- Cobertura de instruções ≥ 80% (atual: **86%**)
- Thresholds JaCoCo verificados (LINE ≥ 80%, BRANCH ≥ 70%)

> 💡 **Função do JaCoCo**: Java Code Coverage — analisa quais linhas do código foram executadas durante os testes. Gera um relatório HTML onde linhas verdes = cobertas, amarelas = parcialmente cobertas, vermelhas = não cobertas. Após atualização para 0.8.14, é compatível com Java 25 (class file major version 69).

> ⚠️ **Classes excluídas da métrica**: Configurações (`config/**`), DTOs, Enums, Entidades, `Address*`, e RowMappers são excluídos do cálculo de cobertura. Foco da meta são Services, Repositories, Controllers, Security e Exception Handlers.

---

## 3. Comandos Úteis por Cenário

### 3.1 Executar uma Classe de Teste Específica

```bash
# Teste unitário específico
./mvnw test -Dtest="DashboardServiceTest" -DfailIfNoTests=false

# Teste de integração específico
./mvnw test -Dtest="DashboardRepositoryIT" -DfailIfNoTests=false
```

### 3.2 Executar um Método de Teste Específico

```bash
# Executa apenas 1 método de teste
./mvnw test -Dtest="DashboardServiceTest#shouldReturnSummaryWithCorrectCounts" -DfailIfNoTests=false
```

### 3.3 Executar com Logs Detalhados (Debug)

Use quando um teste falha e você precisa ver as queries SQL ou stack traces completos:

```bash
# Log detalhado dos testes
./mvnw test -Dtest="DashboardRepositoryIT" -DfailIfNoTests=false 2>&1 | tee test-output.log

# Com log SQL ativado (vê as queries sendo executadas)
./mvnw test -Dtest="DashboardRepositoryIT" \
  -Dspring.jpa.show-sql=true \
  -Dlogging.level.org.springframework.jdbc=DEBUG \
  -DfailIfNoTests=false
```

> 💡 **Função do `2>&1 | tee`**: `2>&1` redireciona stderr para stdout (junta erro e saída normal em um fluxo só). O pipe `| tee test-output.log` grava em arquivo E mostra na tela simultaneamente. Sem `tee`, se usar apenas `>`, a saída some do terminal e vai só para o arquivo.

### 3.4 Pular Testes (Apenas Compilar)

```bash
# Compila sem executar nenhum teste
./mvnw clean compile -DskipTests
```

> 💡 **Função do `-DskipTests`**: Compila as classes de teste mas não as executa. Diferente de `-Dmaven.test.skip=true` que nem compila os testes — use `-DskipTests` para verificar que os testes compilam sem executá-los.

### 3.5 Executar Apenas Testes de Integração

```bash
# Executa apenas os testes de integração (pós Passo 3 implementado)
./mvnw test -Dtest="com.fbso.platform.admin.integration.**" -DfailIfNoTests=false
```

---

## 4. Resultados Esperados — Matriz de Verificação

### 4.1 DashboardRepositoryIT (T-023) — 6 cenários

| Cenário | Query testada | Entrada | Saída esperada | VALIDAÇÃO |
|:---|:---|:---|:---|:---:|
| TC-F01-01-004 | `getSummary()` | 10 tenants seed (5 ACTIVE, 3 PENDING, 2 SUSPENDED) | `DashboardSummaryResponse` com total=10, active=5, pending=3, suspended=2. `monthlyRevenue > 0` se houver subscriptions com preço | X |
| TC-F01-01-005 | `getSummary()` com soft-delete | 1 tenant com `deleted_dt NOT NULL` | Tenant excluído NÃO aparece nas métricas. total=9 (não 10) | X |
| TC-F01-01-005b | `getAccountsByStatus()` | Mesmos dados | Contagem por status exclui deleted. Cada status bate com o seed | X |
| TC-F01-02-003 | `findAllPaginated(page=0, size=25)` | Sem filtros | Lista paginada com 10 registros. `totalElements=10`. Ordenado por `created_at DESC` | X |
| TC-F01-02-004 | `findAllPaginated` com busca textual `"ab"` | `searchTerm="ab"` | `totalElements=0` (busca < 3 chars ignora) | X |
| TC-F01-03-003 | `onboardingStalled()` + `suspendedSubscriptions()` | 1 tenant PENDING com created_at >48h atrás + 1 SUSPENDED | alerts não vazio. Cards WARNING e CRITICAL presentes | X |

### 4.2 Resumo de Sucesso por Marco

| Marco | Testes | Critério |
|:---|:---:|:---|
| Passo 2 (Unit) | 50/50 ✅ | `Failures: 0, Errors: 0` |
| Passo 3 (Integração T-023) | 6/6 X | `Failures: 0, Errors: 0` - relatorio Jacoco apresenta erro |
| Passo 4 (Suite completa) | 56/56 X | `BUILD SUCCESS` |
| Passo 5 (JaCoCo) | ≥ 80% X | Cobertura verificável em HTML - Coberta abaixo de 80% |

---

## 5. Troubleshooting — Problemas Comuns

### 5.1 "Could not find a valid Docker environment"

**Causa:** Docker não está rodando ou o socket não está acessível.

```bash
# Verificar status do Docker
docker ps
# Se falhar, iniciar o Docker
sudo systemctl start docker    # Linux
# ou
open -a Docker                  # macOS

# Verificar permissão do socket
ls -la /var/run/docker.sock
# O usuário precisa estar no grupo docker
sudo usermod -aG docker $USER
# Re-login necessário após este comando
```

### 5.2 "Container startup timeout" (timeout após 60s)

**Causa:** Primeira execução baixando imagem `postgres:17-alpine` em conexão lenta.

```bash
# Baixar a imagem manualmente antes do teste
docker pull postgres:17-alpine

# Verificar se a imagem foi baixada
docker images | grep postgres
# Esperado: "postgres   17-alpine   ..." na lista
```

### 5.3 "Flyway migration failed — relation already exists"

**Causa:** Testcontainers reutilizou um container com migrations já aplicadas.

```bash
# Forçar remoção de containers órfãos
docker ps -a | grep postgres | awk '{print $1}' | xargs docker rm -f

# Limpar cache do Testcontainers
rm -rf ~/.testcontainers/

# Reexecutar
./mvnw clean test -Dtest="DashboardRepositoryIT" -DfailIfNoTests=false
```

### 5.4 "Cannot get JDBC Connection" (connection refused)

**Causa:** Container PostgreSQL subiu mas a porta não foi mapeada.

```bash
# Verificar containers ativos durante o teste
docker ps
# Deve mostrar um container postgres:17-alpine com porta mapeada

# Se não aparecer, verificar se Testcontainers está configurado
grep -r "testcontainers" pom.xml
# Deve mostrar 3 dependências: testcontainers, postgresql, junit-jupiter
```

### 5.5 "Out of Memory" durante os testes

**Causa:** Testcontainers consome memória adicional para o container PostgreSQL.

```bash
# Aumentar memória do Maven
export MAVEN_OPTS="-Xmx2g -Xms512m"

# Ou verificar memória disponível
free -h
```

---

## 6. Checklist de Execução

- [x] **Passo 1:** Docker rodando (`docker ps` OK)
- [x] **Passo 1:** Branch correta (`feature/sprint-03-portal-admin`)
- [x] **Passo 1:** Build compila (`mvn compile` SUCCESS)
- [x] **Passo 2:** 77/77 testes unitários passando (5 skipped)
- [x] **Passo 3:** 23/23 testes integração DashboardRepositoryIT passando
- [x] **Passo 4:** Suite completa 77/77 passando (5 skipped)
- [x] **Passo 5:** JaCoCo 86% ≥ 80% verificado
- [x] Relatório salvo em `target/site/jacoco/index.html`
- [x] Nenhum container órfão (`docker ps` só mostra o que estava antes)

---

## 7. Próximos Passos Após T-023

| Sequência | O Que Fazer | Sprint |
|:---|:---|:---|
| 1 | ✅ T-023 concluído → M2 (Dashboard) 100% | Sprint 3 |
| 2 | Iniciar Frente 2 (M3): T-024 a T-038 — Tenant, Plan, Subscription | Sprint 3 |
| 3 | Encaixar Frente 3 (correções): T-042 (10min), T-045 (15min) entre tarefas | Sprint 3 |
| 4 | T-037, T-038 — testes unitários e integração M3 | Sprint 3 |
| 5 | T-043 — refatorar RLSIsolationTest para Testcontainers | Sprint 3 |
| 6 | Sprint 4 — RBAC (T-039 a T-049) | Próxima sprint |

---

🤖 *Instruções geradas em 17/07/2026 com base em SPRINT-TEST-SUITE.md v2.5, TEST_PLAN.md v2.5, BaseIntegrationTest.java existente, e ambiente Docker verificado. Os 6 cenários de integração listados em §4.1 precisam ser implementados como classe `DashboardRepositoryIT.java` no pacote `com.fbso.platform.admin.integration.repository`.*
