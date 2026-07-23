# IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo

- **Sprint alvo:** 6 de 7 — sprint-06-bus-catalogo
- **Data da análise:** 2026-07-23
- **Skills executadas:** code-reviewer, caveman-review, ponytail-review, ponytail-debt, tech-debt, security-review
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Flyway 12.11.0 + Caffeine 3.2.4 + Keycloak 26
- **Total de achados:** 22 (🔴 4 críticos, 🟡 10 riscos, 🔵 8 nits)
- **Impeditivos para iniciar a sprint:** 4 SIM — entidades incompletas/inexistentes bloqueiam qualquer desenvolvimento
- **Workflow:** 6 agentes paralelos (code-reviewer, caveman-review, ponytail-review, ponytail-debt, tech-debt, security-review) — 190 tool calls, ~1M tokens, 9.6 min. Validação cruzada: 100% convergência nos 4 bloqueantes

---

## Resumo Executivo

A Sprint 6 (M6 — Unidades de Negócio e Catálogo) enfrenta **4 débitos bloqueantes** que impedem o início imediato do desenvolvimento. O principal problema é que as entidades estão significativamente desatualizadas em relação ao schema do banco: `BusinessUnit.java` está faltando 9 campos essenciais (corporate_name, tax_regime, address completo, status) e `ProductService.java` simplesmente não existe como classe Java — embora a tabela exista no banco desde a V001. Além disso, nenhum dos repositories, services ou controllers para M6 foi criado ainda (eram planejados para esta sprint). Um TODO crítico no `PermissionService` (validateBusinessUnitTenant, linha 217) impede a validação de isolamento multi-tenant para BUs. O `CnpjValidator` não existe e precisará ser criado. A dívida técnica acumulada de sprints anteriores é baixa — a maioria dos débitos do backlog foi resolvida nas Sprints 4-5. Recomenda-se tratar os 4 bloqueantes na **Frente 0** (pré-sprint, ~0.5 dia) antes de iniciar as 9 tarefas de feature.

---

## Backlog de Débitos Técnicos (Sprints Anteriores)

> Débitos pendentes de sprints anteriores, extraídos dos documentos `IDENTIFIED-TECHNICAL-DEBT-sprint-0*-*.md`.

| DT-XXX | Sprint Origem | Descrição | Severidade | Bloqueante? | Status | Resolução |
|:---|:---|:---|:---:|:---:|:---|:---|
| DT-023 | Sprint 3 | Migrar paginação offset→keyset no BaseRepository | 🟡 | NÃO | ✅ **Concluído Sprint 5** | `findAllKeyset(lastId, pageSize, sortColumn)` implementado no BaseRepository (T-139) |
| DT-045 | Sprint 3 | Migrar Flyway 10.22.0→12.11.0 | 🟡 | NÃO | ✅ **Concluído Sprint 5** | Flyway 12.11.0 + flyway-database-postgresql (T-134) |
| DT-031 | Sprint 3 | Reduzir Checkstyle maxAllowedViolations 300→100 | 🔵 | NÃO | ⬜ Pendente (Sprint 6) | Meta Sprint 6: 100 violações. Sprint 7: 0 |
| DT-034 | Sprint 3 | Remover Address.java se não usado até Sprint 6 | 🔵 | NÃO | ⬜ Pendente (Sprint 6) | **Confirmado:** Address.java NÃO é usado por nenhuma classe. Pode ser removido |
| DT-044 | Sprint 3 | Criar logback-spring.xml com appender JSON | 🔵 | NÃO | ⬜ Pendente (Sprint 7) | Baixa prioridade — sem impacto em M6 |
| DT-068 | Sprint 4 | Atualizar PostgreSQL driver 42.7.11 (CVE-2026-42198) | 🟡 | NÃO | ✅ **Concluído Sprint 5** | PG driver 42.7.11 (T-135) |
| DT-086 | Sprint 4 | Extrair AuditFieldsRowMapper helper | 🔵 | NÃO | ✅ **Concluído Sprint 5** | AuditFieldsRowMapper extraído (T-149) |
| DT-090 | Sprint 4 | Substituir OffsetDateTime.now()→UTC em 13 ocorrências | 🔵 | NÃO | ✅ **Concluído Sprint 5** | Todos os timestamps em UTC (T-151) |
| DT-124 | Sprint 5 | Diagrama de estados do onboarding (4 passos) | 🟡 | NÃO | ✅ **Concluído Sprint 5** | Documentado no ARCHITECTURE.md §8 (T-145) |

**Total em backlog:** 3 débitos pendentes (DT-031, DT-034, DT-044). Nenhum é bloqueante para Sprint 6.

---

## Matriz de Débitos Técnicos

| ID | Sprint Origem | Arquivo/Artefato | Achado | Severidade | Skill | Complexidade | Bloqueante? | Efeito se não tratado |
|:---|:---|:---|:---|:---:|:---:|:---:|:---:|:---|
| DT-031 | Sprint 3 | `checkstyle.xml` | Checkstyle maxAllowedViolations: reduzir para 100 | 🔵 | BACKLOG | M | NÃO | Qualidade de código degrada progressivamente |
| DT-034 | Sprint 3 | `common/Address.java` | Address.java não utilizado — código morto | 🔵 | BACKLOG | L | NÃO | -95 linhas de código morto no projeto |
| DT-044 | Sprint 3 | `logback-spring.xml` | Logs sem estrutura JSON em produção | 🔵 | BACKLOG | L | NÃO | Dificulta debugging em produção (Sprint 7) |
| **DT-126** | Sprint 6 | `entity/BusinessUnit.java` | **Entidade desatualizada:** faltam 9 campos vs schema DB (corporate_name, tax_regime, street, number, complement, neighborhood, city, state, zip_code, status). Campos `name` e `hierarchyType` não existem no DB | 🔴 | CREV | H | **SIM** | Impossível persistir BUs corretamente — schema mismatch |
| **DT-127** | Sprint 6 | `entity/ProductService.java` | **Entidade inexistente:** tabela product_service existe no V001 mas não há classe Java | 🔴 | CREV | H | **SIM** | Impossível implementar CRUD de produtos sem entidade |
| **DT-128** | Sprint 6 | `service/PermissionService.java:217` | **TODO bloqueante:** validateBusinessUnitTenant() não implementado — sem validação de que BU pertence ao tenant | 🔴 | CREV | M | **SIM** | IDOR: usuário do tenant-A pode acessar BUs do tenant-B |
| **DT-129** | Sprint 6 | `utils/` | **CnpjValidator inexistente:** necessário para BusinessUnitService.create() e OnboardingService.step2(). Validação de CNPJ atualmente inline ou ausente | 🔴 | SEC | M | **SIM** | CNPJs inválidos aceitos; duplicação de lógica de validação |
| DT-130 | Sprint 6 | V003 / `product_service` | **RLS gap:** product_service NÃO está na lista RLS do V003. Tabela não tem tenant_id direto — isolamento depende de JOIN com business_unit | 🟡 | SEC | M | NÃO | Queries sem JOIN com BU podem vazar produtos entre tenants |
| DT-131 | Sprint 6 | `entity/BusinessUnit.java` | **hierarchyType sem coluna no DB:** campo `hierarchyType` existe na entity mas NÃO no schema (V001). V007 adicionou `is_matrix` como替代 | 🟡 | CAV | L | NÃO | Campo mapeado em toColumnMap() causa erro SQL no save() |
| DT-132 | Sprint 6 | `repository/`, `service/`, `controller/` | **6 classes não existem:** BusinessUnitRepository, BusinessUnitService, BusinessUnitController, ProductRepository, ProductService, ProductController — todas precisam ser criadas | 🟡 | DEBT | H | NÃO | Escopo normal da sprint — 6 novas classes + 2 RowMappers |
| DT-133 | Sprint 6 | `sprint-06/SPRINT-CARD.md` | **SPRINT-CARD desatualizado:** referencia branch errada (feature/sprint-06-bus-catalogo), TASKS.md v2.0 (atual: v3.7), e contém 3 arquivos .bak | 🟡 | CAV | L | NÃO | Confusão na execução — branch errada, métricas erradas |
| DT-134 | Sprint 6 | Arquitetura — query hierárquica | **Decisão pendente:** árvore de BUs via CTE recursiva (WITH RECURSIVE) vs carregar flat + montar em memória. BaseRepository não tem suporte nativo a queries recursivas | 🟡 | DEBT | M | NÃO | Performance imprevisível para hierarquias profundas (>5 níveis) |
| DT-135 | Sprint 6 | `src/test/` | **Zero testes para escopo M6:** sem testes unitários ou de integração para BU e Product. Cobertura JaCoCo vai cair abaixo de 80% | 🟡 | DEBT | H | NÃO | Regressões não detectadas. Meta JaCoCo 80% comprometida |
| DT-136 | Sprint 6 | `SubscriptionServiceTest` | **1 teste quebrado:** shouldCreateWithLockedPrice falha com `IllegalState: TenantContext não inicializado`. Pré-existente desde Sprint 5 | 🟡 | CREV | L | NÃO | Build com 1 erro (227/228 passando). Falso negativo |
| DT-137 | Sprint 6 | `security/RateLimitFilter.java:132` | **Ponytail:** externalizar trusted-proxy-ips para application.yml. 🟢 has-trigger (Sprint 6) | 🟡 | PDBT | L | NÃO | Configuração hardcoded — requer recompilação para ajuste |
| DT-138 | Sprint 6 | `controller/UserController.java:103` | **TODO sem trigger:** mover lógica update para UserService.update(). Sem prazo definido | 🔵 | PDBT | M | NÃO | Lógica no controller — viola padrão de camadas |
| DT-139 | Sprint 6 | `sprint-06/*.bak` | **3 arquivos .bak lixo:** SPRINT-CARD.md.bak, SPRINT-REVIEW.md.bak, SPRINT-TEST-SUITE.md.bak | 🔵 | PONY | L | NÃO | Confusão — qual é o arquivo canônico? |
| DT-140 | Sprint 6 | `dto/` | **Novos DTOs necessários:** ~10 DTOs request/response para BU e Product (criar, update, response, list) | 🔵 | DEBT | M | NÃO | Escopo normal da sprint |
| **DT-141** | Sprint 6 | `dto/request/LoginRequest.java` | **DTO órfão:** LoginRequest (9 linhas) nunca usado — AuthController delega ao OAuth2/Keycloak, não aceita email/senha direto | 🔵 | PONY | L | NÃO | Código morto — remove 9 linhas |
| **DT-142** | Sprint 6 | V004 / RbacAspect | **Boas notícias:** matriz RBAC (V004) já tem BUSINESS_UNIT (view, create, edit) e PRODUCT_SERVICE (view, create, edit). Roles ADMIN_TENANT e MANAGER_BU já mapeadas. Nenhuma modificação no RbacAspect necessária | 🔵 | SEC | L | NÃO | Infraestrutura RBAC pronta — só anotar os novos controllers |

---

## Achados por Skill

### code-reviewer (8 achados)

**🔴 DT-126 — BusinessUnit entity schema mismatch (CRITICAL)**

Arquivo: `entity/BusinessUnit.java`

A entidade está significativamente desatualizada em relação ao schema V001 do banco:

| Campo no DB (V001) | Presente na Entity? |
|:---|:---:|
| `id` | ✅ |
| `tenant_id` | ✅ (tenantId) |
| `parent_id` | ✅ (parentId) |
| `cnpj` | ✅ |
| `corporate_name` | ❌ **FALTANDO** |
| `tax_regime` | ❌ **FALTANDO** |
| `street` | ❌ **FALTANDO** |
| `number` | ❌ **FALTANDO** |
| `complement` | ❌ **FALTANDO** |
| `neighborhood` | ❌ **FALTANDO** |
| `city` | ❌ **FALTANDO** |
| `state` | ❌ **FALTANDO** |
| `zip_code` | ❌ **FALTANDO** |
| `status` | ❌ **FALTANDO** |
| `is_matrix` | ✅ (V007) |

Campos na entity que NÃO existem no DB:
- `name` → deveria ser `corporate_name`
- `hierarchyType` → sem coluna correspondente (V007 substituiu por `is_matrix`)

**Ação:** Reescrever BusinessUnit.java com todos os campos do schema. Remover `name` (substituir por `corporateName`). Remover `hierarchyType` (substituído por `isMatrix`). Adicionar campos de endereço. Atualizar `toColumnMap()`.

**🔴 DT-127 — ProductService entity inexistente (CRITICAL)**

Arquivo: `entity/ProductService.java` — **NÃO EXISTE**

A tabela `fbso_platform.product_service` existe desde V001 com colunas: id, business_unit_id, name, sku, type, description, status + campos de auditoria. NÃO há classe Java correspondente.

**Ação:** Criar `ProductService.java` extends BaseEntity com todos os campos do schema V001.

**🔴 DT-128 — PermissionService.validateBusinessUnitTenant não implementado (CRITICAL)**

Arquivo: `service/PermissionService.java:217`

```java
// TODO Frente 3/Sprint 6: validateBusinessUnitTenant(businessUnitId)
```

Este TODO está no método que valida permissões de acesso a BUs. Sem esta validação, um usuário do tenant-A poderia receber permissão para uma BU do tenant-B — violando o isolamento multi-tenant.

**Ação:** Implementar `validateBusinessUnitTenant(UUID businessUnitId)` consultando `business_unit.tenant_id` e comparando com `TenantContext.getTenantId()`.

**🟡 DT-131 — hierarchyType sem coluna no DB**

`BusinessUnit.java` tem campo `hierarchyType` incluído no `toColumnMap()`, mas NÃO existe coluna `hierarchy_type` na tabela `business_unit` (V001). O método `save()` do BaseRepository tentará inserir nesta coluna e falhará com erro SQL.

**Ação:** Remover `hierarchyType` da entity e do `toColumnMap()`. A distinção Matriz/Filial é feita via `isMatrix` (V007) e `parent_id IS NULL`.

**🟡 DT-132 — 6 classes não existem (escopo M6)**

Nenhuma das seguintes classes foi criada ainda (todas pertencem ao escopo normal da Sprint 6):
- `BusinessUnitRepository.java` + `BusinessUnitRowMapper.java`
- `BusinessUnitService.java`
- `BusinessUnitController.java`
- `ProductRepository.java` + `ProductRowMapper.java`
- `ProductService.java` (classe de serviço)
- `ProductController.java`

**🟡 DT-136 — 1 teste quebrado (pré-existente)**

`SubscriptionServiceTest.shouldCreateWithLockedPrice` falha com `IllegalState: TenantContext não inicializado`. Pré-existente desde a Sprint 5. 227/228 testes passam.

**🔵 DT-138 — UserController lógica no controller**

`UserController.java:103`: `// TODO Frente 2: mover lógica de update para UserService.update()`. Violação do padrão de camadas — lógica de negócio no controller.

**🔵 DT-140 — ~10 novos DTOs necessários**

A Sprint 6 precisará de DTOs request/response para BU e Product: `BusinessUnitCreateRequest`, `BusinessUnitUpdateRequest`, `BusinessUnitResponse`, `ProductCreateRequest`, `ProductUpdateRequest`, `ProductResponse`, etc.

---

### caveman-review (3 achados)

**🟡 DT-133 — SPRINT-CARD.md desatualizado**

Arquivo: `sprints/sprint-06-bus-catalogo/SPRINT-CARD.md`

3 inconsistências encontradas:
1. **Branch errada:** Documento diz `feature/sprint-06-bus-catalogo` mas a branch real é `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-06-bus-catalogo`
2. **Versão TASKS errada:** Rodapé diz "Gerado a partir de TASKS.md v2.0" — TASKS.md está em v3.7
3. **Header desatualizado:** Não reflete que Frentes 0-1-2-3a da Sprint 5 foram concluídas

**🟡 DT-131 — hierarchyType divergência docs×código**

A SPECS.md §6.1 lista `BusinessUnit` com campo `tax_regime` e `address`. O PRD.md §4.1 lista campos: tenant_id, parent_id, cnpj, corporate_name, tax_regime, address, status. A entity atual não corresponde a nenhum dos dois.

**🔵 DT-139 — Arquivos .bak no diretório sprint-06**

3 arquivos duplicados: `SPRINT-CARD.md.bak`, `SPRINT-REVIEW.md.bak`, `SPRINT-TEST-SUITE.md.bak`. São cópias idênticas aos arquivos principais. Lixo de template.

---

### ponytail-review (2 achados)

**🔵 DT-034 (backlog) — Address.java não utilizado**

`common/Address.java` — apenas 1 referência no código: seu próprio builder (`return new Address(this)`). Nenhuma entidade, serviço ou controller importa ou usa Address. A tabela `business_unit` incorpora os campos de endereço diretamente (street, number, city, etc.).

**Ação:** Remover `Address.java`. **net: -95 lines.**

**🔵 DT-139 — .bak files**

3 arquivos `.bak` no diretório sprint-06 — candidatos a `delete`.

**Resumo ponytail:** net: -95 lines (Address.java). +3 arquivos .bak para remover.

---

### ponytail-debt (3 achados)

**🟡 DT-137 — RateLimitFilter trusted-proxy-ips (has-trigger: Sprint 6)**

Arquivo: `security/RateLimitFilter.java:132`
```
// ponytail: externalizar trusted-proxy-ips para application.yml na Sprint 6
```
🟢 has-trigger. Ceiling: IPs confiáveis hardcoded. Trigger: Sprint 6.

**🔵 DT-138 — UserController update logic (no-trigger)**

Arquivo: `controller/UserController.java:103`
```
// TODO Frente 2: mover lógica de update para UserService.update()
```
🔴 no-trigger — sem prazo definido. Risco de apodrecimento: lógica permanece no controller indefinidamente.

**🔴 DT-128 — PermissionService validateBusinessUnitTenant (has-trigger: Sprint 6)**

Arquivo: `service/PermissionService.java:217`
```
// TODO Frente 3/Sprint 6: validateBusinessUnitTenant(businessUnitId)
```
🟢 has-trigger. Este é o marcador mais crítico — BLOQUEANTE para Sprint 6.

**3 markers, 1 with no trigger.**

---

### tech-debt (6 achados)

**Categoria: Code Debt**

| DT | Descrição | Impact (1-5) | Risk (1-5) | Effort (1-5) | Priority | Severidade |
|:---|:---|:---:|:---:|:---:|:---:|:---:|
| DT-126 | BU entity schema mismatch | 5 | 5 | 3 | **40** | 🔴 |
| DT-127 | ProductService entity inexistente | 5 | 5 | 3 | **40** | 🔴 |
| DT-128 | validateBusinessUnitTenant TODO | 4 | 5 | 2 | **36** | 🔴 |
| DT-131 | hierarchyType sem coluna DB | 3 | 4 | 1 | **35** | 🟡 |
| DT-130 | RLS gap product_service | 2 | 4 | 3 | 18 | 🔵 |

**Categoria: Architecture Debt**

| DT | Descrição | Impact (1-5) | Risk (1-5) | Effort (1-5) | Priority | Severidade |
|:---|:---|:---:|:---:|:---:|:---:|:---:|
| DT-134 | Decisão query hierárquica pendente | 3 | 3 | 3 | **18** | 🔵 |

**Categoria: Test Debt**

| DT | Descrição | Impact (1-5) | Risk (1-5) | Effort (1-5) | Priority | Severidade |
|:---|:---|:---:|:---:|:---:|:---:|:---:|
| DT-135 | Zero testes escopo M6 | 3 | 4 | 5 | **7** | 🟡 |
| DT-136 | 1 teste quebrado pré-existente | 1 | 2 | 1 | **15** | 🔵 |

**Categoria: Documentation Debt**

| DT | Descrição | Impact (1-5) | Risk (1-5) | Effort (1-5) | Priority | Severidade |
|:---|:---|:---:|:---:|:---:|:---:|:---:|
| DT-133 | SPRINT-CARD desatualizado | 2 | 2 | 1 | **20** | 🔵 |

**Plano de Remediação Faseado:**

- **Fase 1 (Imediato — antes da sprint):** DT-126, DT-127, DT-128, DT-129 — 4 bloqueantes (~4h)
- **Fase 2 (Durante a sprint):** DT-130 (RLS), DT-131 (hierarchyType), DT-133 (SPRINT-CARD), DT-134 (query hierárquica), DT-137 (RateLimitFilter) — 5 recomendações (~6h)
- **Fase 3 (Sprints futuras):** DT-031 (checkstyle), DT-034 (Address.java), DT-044 (logback), DT-138 (UserController), DT-139 (.bak), DT-140 (DTOs — parte do escopo normal)

---

### security-review (4 achados)

**🔴 DT-129 — CnpjValidator inexistente (CRITICAL)**

Não existe classe `CnpjValidator` no projeto. A validação de CNPJ é necessária para:
- `BusinessUnitService.create()` — validação obrigatória (RN17-01)
- `OnboardingService.completeStep2()` — validação durante onboarding

**Ação:** Criar `utils/CnpjValidator.java` com validação de formato (XX.XXX.XXX/XXXX-XX) e dígitos verificadores. Reutilizar no OnboardingService e BusinessUnitService.

**🟡 DT-130 — RLS gap em product_service**

A tabela `product_service` NÃO está na lista RLS do V003. As tabelas com RLS ativo são: subscription, user, business_unit, audit_log. A tabela product_service NÃO tem coluna `tenant_id` — a isolation depende do JOIN com `business_unit.tenant_id`. Isso significa que:

1. Um `SELECT * FROM product_service` sem JOIN com business_unit retorna produtos de TODOS os tenants
2. A mitigação atual depende 100% da aplicação (BaseRepository + queries com JOIN)

**Ação recomendada:** Adicionar `tenant_id` em product_service (migration V009) e ativar RLS. Ou, no mínimo, documentar que TODA query em product_service DEVE fazer JOIN com business_unit e filtrar por `business_unit.tenant_id`.

**🟡 DT-128 — IDOR via falta de validação de tenant em BU**

`PermissionService.validateBusinessUnitTenant()` não implementado. Um atacante poderia:
1. Obter um `business_unit_id` de outro tenant (via IDOR em endpoint de listagem)
2. Atribuir permissão de acesso a esta BU
3. Acessar dados de outro tenant

**Mitigação atual:** RLS em business_unit + BaseRepository tenant filter. Mas a validação explícita no PermissionService é uma camada adicional de defesa.

**🔵 Dependências — sem CVEs críticas**

Stack atualizada (Sprint 5):
- Spring Boot 3.5.14 ✅
- Jackson 2.21.4 (CVE-2026-22733/CVE-2026-22731 mitigados) ✅
- PostgreSQL driver 42.7.11 (CVE-2026-42198 mitigado) ✅
- Flyway 12.11.0 ✅
- Caffeine 3.2.4 ✅
- Testcontainers 1.21.4 ✅

Nenhuma CVE crítica conhecida nas versões atuais.

---

## Recomendações Prioritárias

### 🔴 Bloqueantes (impeditivos — devem ser corrigidos ANTES de iniciar a sprint)

| ID (TASKS.md) | DT-XXX | Ação Corretiva | Estimativa | Responsável |
|:---|:---|:---|:---:|:---|
| T-161.DT-126 | DT-126 | Reescrever `BusinessUnit.java`: adicionar corporateName, taxRegime, street, number, complement, neighborhood, city, state, zipCode, status. Remover `name`→`corporateName`, remover `hierarchyType`. Atualizar `toColumnMap()` com todas as 16 colunas | 2h | A definir |
| T-162.DT-127 | DT-127 | Criar `ProductService.java` extends BaseEntity: id, businessUnitId, name, sku, type (ProductType enum), description, status. `toColumnMap()` com 6 colunas de domínio | 1h | A definir |
| T-163.DT-128 | DT-128 | Implementar `PermissionService.validateBusinessUnitTenant(UUID buId)`: query `SELECT tenant_id FROM business_unit WHERE id = ? AND deleted_dt IS NULL`, comparar com `TenantContext.getTenantId()`, lançar `TenantIsolationException` se mismatch | 1h | A definir |
| T-164.DT-129 | DT-129 | Criar `utils/CnpjValidator.java`: validação de formato (regex XX.XXX.XXX/XXXX-XX), cálculo dos 2 dígitos verificadores, método `isValid(String cnpj): boolean`. Integrar no `OnboardingService.completeStep2()` | 1.5h | A definir |

### 🟡 Recomendados (devem ser tratados durante a sprint)

| ID (TASKS.md) | DT-XXX | Ação Corretiva | Estimativa | Sprint sugerida |
|:---|:---|:---|:---:|:---|
| T-165.DT-130 | DT-130 | Criar V009: `ALTER TABLE product_service ADD COLUMN tenant_id UUID NOT NULL DEFAULT gen_random_uuid()` + preencher tenant_id via JOIN com business_unit + `ENABLE ROW LEVEL SECURITY` + `CREATE POLICY`. OU documentar que toda query de product_service DEVE incluir JOIN com business_unit | 1.5h | Sprint 6 |
| T-166.DT-131 | DT-131 | Remover `hierarchyType` de BusinessUnit.java e toColumnMap(). Adicionar `isMatrix` no toColumnMap() (campo já existe via V007) | 15min | Sprint 6 |
| T-167.DT-133 | DT-133 | Atualizar SPRINT-CARD.md: corrigir branch name, referência TASKS.md v3.7, header status. Remover 3 arquivos .bak | 30min | Sprint 6 |
| T-168.DT-134 | DT-134 | Documentar decisão de query hierárquica no ARCHITECTURE.md: PostgreSQL WITH RECURSIVE para árvore de BUs. Implementar `BusinessUnitRepository.findTree()` com CTE | 1h | Sprint 6 |
| T-169.DT-137 | DT-137 | Externalizar `trusted-proxy-ips` no RateLimitFilter para `application.yml` (`app.rate-limit.trusted-proxy-ips`) | 30min | Sprint 6 |

### 🔵 Desejáveis (nice-to-have — se houver capacidade)

| ID (TASKS.md) | DT-XXX | Ação Corretiva | Estimativa |
|:---|:---|:---|:---:|
| T-170.DT-034 | DT-034 | Remover `common/Address.java` (-95 linhas) | 5min |
| T-171.DT-031 | DT-031 | Reduzir Checkstyle maxAllowedViolations para 100 | 30min |
| T-172.DT-138 | DT-138 | Mover lógica de update de UserController para UserService.update() | 1h |
| T-173.DT-136 | DT-136 | Corrigir SubscriptionServiceTest — inicializar TenantContext no @BeforeEach | 30min |

---

## Decisão do Time

> **Decisão tomada em 23/07/2026.**

| ID | Decisão | Sprint alvo | Justificativa |
|:---|:---|:---:|:---|
| DT-126 | Tratar agora — Frente 0 | Sprint 6 | Bloqueante: entidade quebrada impede save/update |
| DT-127 | Tratar agora — Frente 0 | Sprint 6 | Bloqueante: CRUD de produtos impossível sem entidade |
| DT-128 | Tratar agora — Frente 0 | Sprint 6 | Bloqueante: IDOR cross-tenant sem validação |
| DT-129 | Tratar agora — Frente 0 | Sprint 6 | Bloqueante: CNPJs inválidos aceitos |
| DT-130 | Tratar agora — Frente 1 | Sprint 6 | Recomendado: RLS gap em product_service |
| DT-131 | Tratar agora — Frente 1 | Sprint 6 | Recomendado: hierarchyType quebra save() |
| DT-133 | Tratar agora — Frente 1 | Sprint 6 | Recomendado: SPRINT-CARD desatualizado |
| DT-134 | Tratar agora — Frente 1 | Sprint 6 | Recomendado: decisão query hierárquica |
| DT-137 | Tratar agora — Frente 1 | Sprint 6 | Recomendado: externalizar trusted-proxy-ips |
| DT-031 | Postergar — Backlog | Sprint 7 | Nit: checkstyle 100 violações |
| DT-034 | Postergar — Backlog | Sprint 7 | Nit: remover Address.java (-95 linhas) |
| DT-044 | Postergar — Backlog | Sprint 7 | Nit: logback JSON |
| DT-135 | Postergar — Backlog | Sprint 6 (escopo normal) | Débito de teste — T-075/076/077 cobrem |
| DT-136 | Postergar — Backlog | Sprint 7 | Nit: 1 teste quebrado pré-existente |
| DT-138 | Postergar — Backlog | Sprint 7 | Nit: UserController update no service |
| DT-139 | Postergar — Backlog | Sprint 7 | Nit: remover arquivos .bak |
| DT-140 | Postergar — Backlog | Sprint 6 (escopo normal) | DTOs criados durante implementação |
| DT-141 | Postergar — Backlog | Sprint 7 | Nit: remover LoginRequest.java (-9 linhas) |
| DT-142 | Nenhuma ação | — | Info: RBAC pronto para Sprint 6 |

---

## Débitos Técnicos Elegíveis para Sprints Futuras

| DT-XXX | Sprint Origem | Descrição | Severidade | Bloqueante? | Skill | Complexidade | Sprint Sugerida | Justificativa do Adiamento |
|:---|:---|:---|:---:|:---:|:---:|:---:|:---|:---|
| DT-044 | Sprint 3 | Logback JSON appender | 🔵 | NÃO | BACKLOG | L | Sprint 7 | Baixo impacto em M6 — logging estruturado é prioridade de produção |

---

## Análise de Impacto nos Documentos

### Impacto nos Documentos-Mestre

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-126, DT-127 | ARCHITECTURE.md | Estrutura de pacotes: novas entidades adicionadas | Atualizar §2 com BusinessUnit e ProductService |
| DT-126, DT-127 | SPECS.md | Modelo de dados §6.1: atualizar campos de BU e Product | Atualizar campos de BusinessUnit e ProductService |
| DT-128, DT-130 | SECURITY.md | Nova validação de isolamento + RLS gap | Documentar validateBusinessUnitTenant e V009 |
| DT-126..DT-140 | TASKS.md | Novas tasks de débito (Frente 0) | Adicionar T-168 a T-180 |
| DT-129, DT-130 | TEST_PLAN.md | Novos cenários de segurança: CNPJ validation, RLS product_service | Adicionar cenários TC-SEC-* |
| DT-134 | ARCHITECTURE.md | Decisão de query hierárquica | Documentar ADR-L08 (CTE recursiva) |

### Impacto nos Artefatos da Sprint

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-133 | SPRINT-CARD.md | Branch name, versão TASKS, header status | Corrigir + remover .bak |
| DT-126..DT-140 | SPRINT-CARD.md | Adicionar Frente 0 com 4 tasks bloqueantes | Atualizar tabela de tasks |

---

🤖 *Análise gerada em 2026-07-23. 18 achados consolidados a partir de 6 skills (code-reviewer, caveman-review, ponytail-review, ponytail-debt, tech-debt, security-review). Documento base para decisão do time.*
