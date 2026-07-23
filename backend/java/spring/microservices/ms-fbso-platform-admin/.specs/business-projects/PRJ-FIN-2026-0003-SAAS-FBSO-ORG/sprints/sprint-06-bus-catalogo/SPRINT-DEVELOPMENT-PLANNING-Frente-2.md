# SPRINT-DEVELOPMENT-PLANNING-Frente-2.md — Plano de Desenvolvimento: Sprint 6 — Frente 2 (M6 Features)

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 6 de 7 — sprint-06-bus-catalogo
- **Frente:** Frente 2 — Sprint Backlog M6 Features (EP-04b)
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Flyway 12.11.0 · Caffeine 3.2.4
- **Data:** 23 de Julho de 2026
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-06-bus-catalogo`
- **Pré-requisitos:** Frente 0 ✅ (4/4) + Frente 1 ✅ (5/5). 288 testes (0 failures).

---

## 1. Visão Geral

- **Sprint Goal:** "Estrutura hierárquica de Unidades de Negócio (Matriz/Filial) com CNPJ único entre ativos. Catálogo de Produtos/Serviços segmentado por BU com SKU único. Isolamento multi-tenant verificado."
- **Frente 2 Goal:** Implementar CRUD completo para BusinessUnit e ProductService com todas as RNs, endpoints REST, e bateria completa de testes.
- **Tasks:** 9 (T-069 a T-077)
- **Ordem de execução:** Sequencial com dependências: Repository → Service → Controller → Tests
- **Features:** F04-05 (Unidades de Negócio), F04-06 (Catálogo de Produtos)
- **RNs:** RN17-01 a RN17-05, RN18-01 a RN18-04 (9 RNs)

### Aproveitamento de Frentes Anteriores

| Artefato | Origem | Status |
|:---|:---|:---:|
| `BusinessUnit.java` (16 campos) | F0 (DT-126) | ✅ Pronto |
| `ProductService.java` (+tenantId) | F0 (DT-127) + F1 (DT-130) | ✅ Pronto |
| `BusinessUnitRepository.java` (findTree, findChildren) | F1 (DT-134) | ✅ Parcial |
| `BusinessUnitRowMapper.java` | F1 (DT-134) | ✅ Pronto |
| `CnpjValidator.java` (alfanumérico) | F0 (DT-129) | ✅ Pronto |
| `validateBusinessUnitTenant()` | F0 (DT-128) | ✅ Pronto |
| V009 RLS product_service | F1 (DT-130) | ✅ Pronto |

---

## 2. Dependências entre Tasks

```
Frentes 0+1 (pré-requisitos) ✅
    │
    ├── T-069 (BusinessUnitRepository — complementar)
    │       │
    │       └── T-070 (BusinessUnitService)
    │               │
    │               └── T-071 (BusinessUnitController)
    │
    ├── T-072 (ProductRepository)
    │       │
    │       └── T-073 (ProductService)
    │               │
    │               └── T-074 (ProductController)
    │
    └── T-075 (Testes unitários) ← depende de T-070, T-073
            │
            └── T-076 (Testes integração) ← depende de T-071, T-074
                    │
                    └── T-077 (Testes isolamento) ← depende de T-076
```

---

## 3. Plano por Task

### T-069 — BusinessUnitRepository (complementar)

- **Critério DONE:** `findByCnpj()`, `existsByCnpj()` queries. Demais CRUD via BaseRepository herdado.
- **Estimativa:** 0.5d (reduzido — `findTree()` e `findChildren()` já existem da F1)
- **Abordagem:** Adicionar queries específicas de CNPJ ao repository existente. O BaseRepository já provê `findAll()`, `findById()`, `save()`, `update()`, `softDelete()`, `count()`.

#### Métodos a adicionar

```java
// Verifica se CNPJ já existe no tenant (para RN17-01)
public boolean existsByCnpj(String cnpj, UUID tenantId)

// Busca BU por CNPJ (para validação de duplicidade)
public Optional<BusinessUnit> findByCnpj(String cnpj, UUID tenantId)
```

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `repository/BusinessUnitRepository.java` | 🔄 | +existsByCnpj(), +findByCnpj() |

#### Skills: `311-frameworks-spring-jdbc`, `121-java-object-oriented-design`

---

### T-070 — BusinessUnitService: CRUD com Hierarquia

- **Critério DONE:** CNPJ duplicado → 409. CNPJ alterado → 400. Soft delete libera. parent_id inativo → 422. Sem limite níveis (RN17-04).
- **Estimativa:** 2d
- **RNs:** RN17-01 (CNPJ único), RN17-02 (pai inativo bloqueado), RN17-03 (primeira BU = Matriz), RN17-04 (níveis ilimitados), RN17-05 (seletor por permissão)
- **Abordagem:** Seguir padrão `TenantService` — construtor com repository, métodos com `@Auditable` e `@Transactional`.

#### RNs e Validações

| RN | Validação | Exceção |
|:---|:---|:---|
| RN17-01 | CNPJ único entre ativos do tenant. `existsByCnpj()` antes de create. Soft delete libera reúso. | `DuplicateCnpjException` → 409 |
| RN17-01 | CNPJ imutável após cadastro. Se `request.cnpj != entity.cnpj` → erro. | `BusinessException` → 400 |
| RN17-02 | `parent_id` deve referir BU ativa (`status = ACTIVE`, `deleted_dt IS NULL`) | `BusinessException` → 422 |
| RN17-03 | Primeira BU do tenant = `isMatrix = true`. Onboarding chama com flag. | — |
| RN17-04 | Sem validação de profundidade — CTE recursiva não impõe limite. | — |
| RN17-05 | Admin vê todas BUs. Manager/Operator: `findByUserPermissions(userId)`. | — |

#### Métodos do Service

```
create(BusinessUnitCreateRequest)  → @Auditable(CREATED)
update(id, BusinessUnitUpdateRequest) → @Auditable(UPDATED)
deactivate(id)                     → @Auditable(DEACTIVATED)
findById(id)
findAll(filters)
findTree()                         → delega para repo.findTree()
getAccessibleBUs(userId)           → RN17-05 (seletor)
```

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `service/BusinessUnitService.java` | 🆕 | CRUD + validações RN17-01 a RN17-05 |
| `dto/request/BusinessUnitCreateRequest.java` | 🆕 | Record: cnpj, corporateName, taxRegime, parentId, address... |
| `dto/request/BusinessUnitUpdateRequest.java` | 🆕 | Record: corporateName, taxRegime, parentId, address... (cnpj NÃO incluso) |
| `dto/response/BusinessUnitResponse.java` | 🆕 | Record: id, tenantId, parentId, cnpj, corporateName, children[], isMatrix... |

#### Skills: `121-java-object-oriented-design`, `126-java-exception-handling`, `124-java-secure-coding`

---

### T-071 — BusinessUnitController

- **Critério DONE:** Admin todas. Gerente/Operador apenas autorizadas. Seletor funcional. `@RequiresPermission` em todos os endpoints.
- **Estimativa:** 1.5d
- **Abordagem:** Seguir padrão `TenantController` — `@RestController`, `@RequestMapping`, `@RequiresPermission`.

#### Endpoints

| Método | Path | RBAC | Descrição |
|:---|:---|:---|:---|
| `GET` | `/api/v1/business-units` | Admin/Manager/Operator | Listar (hierárquico). Admin=todas, Manager/Operator=autorizadas (RN17-05) |
| `GET` | `/api/v1/business-units/{id}` | Admin/Manager/Operator | Detalhes + sub-árvore |
| `POST` | `/api/v1/business-units` | Admin/Manager | Criar (CNPJ validado) |
| `PATCH` | `/api/v1/business-units/{id}` | Admin/Manager | Atualizar (CNPJ imutável) |
| `POST` | `/api/v1/business-units/{id}/deactivate` | Admin | Soft delete. Libera CNPJ para reúso. |
| `GET` | `/api/v1/business-units/tree` | Admin/Manager/Operator | Árvore completa (findTree) |

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `controller/BusinessUnitController.java` | 🆕 | 6 endpoints REST |

#### Skills: `302-frameworks-spring-boot-rest`, `304-frameworks-spring-boot-security`

---

### T-072 — ProductRepository

- **Critério DONE:** SKU único por BU. `hasTenantColumn=true` (V009). Métodos: `findByBusinessUnit()`, `existsBySku()`.
- **Estimativa:** 0.5d (reduzido — entity já existe)
- **Abordagem:** Extender `BaseRepository<ProductService>` com `hasTenantColumn=true`. Adicionar queries específicas.

#### Métodos

```java
public List<ProductService> findByBusinessUnit(UUID buId)
public boolean existsBySku(String sku, UUID businessUnitId)
public List<ProductService> findAllByTenant(UUID tenantId, int page, int size)
```

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `repository/ProductRepository.java` | 🆕 | Repository com queries de catálogo |
| `repository/rowmapper/ProductServiceRowMapper.java` | 🆕 | RowMapper para ProductService (7 colunas + audit) |

#### Skills: `311-frameworks-spring-jdbc`, `121-java-object-oriented-design`

---

### T-073 — ProductService (Service): CRUD

- **Critério DONE:** CRUD funcional. "Não mapeado" no response. Soft delete preserva histórico. SKU opcional, único por BU.
- **Estimativa:** 2d
- **RNs:** RN18-01 (catálogo por BU), RN18-02 (SKU único), RN18-03 (indicador "Não mapeado"), RN18-04 (soft delete)
- **Abordagem:** Seguir padrão `PlanService`. Indicador `fiscalMappingStatus = NOT_MAPPED` como default.

#### Validações

| RN | Validação | Exceção |
|:---|:---|:---|
| RN18-01 | `findByBusinessUnit()` — segmentação automática | — |
| RN18-02 | SKU se informado, único por BU ativo (`existsBySku()`). Índice parcial PostgreSQL. | `BusinessException` → 409 |
| RN18-03 | Campo `fiscalMappingStatus` default `NOT_MAPPED` no response | — |
| RN18-04 | Soft delete — `deactivate()` seta `deleted_dt`, produto não aparece em queries | — |

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `service/ProductService.java` | 🆕 | CRUD + validações RN18-01 a RN18-04 |
| `dto/request/ProductCreateRequest.java` | 🆕 | Record: businessUnitId, name, sku?, type, description? |
| `dto/request/ProductUpdateRequest.java` | 🆕 | Record: name, sku?, type, description? |
| `dto/response/ProductResponse.java` | 🆕 | Record: id, businessUnitId, name, sku, type, fiscalMappingStatus, status |

#### Skills: `121-java-object-oriented-design`, `126-java-exception-handling`

---

### T-074 — ProductController

- **Critério DONE:** Admin/Manager: create/edit/deactivate. Operator: view. Desativado não aparece.
- **Estimativa:** 1d
- **Abordagem:** Seguir padrão `PlanController`.

#### Endpoints

| Método | Path | RBAC | Descrição |
|:---|:---|:---|:---|
| `GET` | `/api/v1/products` | Admin/Manager/Operator | Listar (filtro `?business_unit_id=`) |
| `GET` | `/api/v1/products/{id}` | Admin/Manager/Operator | Detalhes |
| `POST` | `/api/v1/products` | Admin/Manager | Criar (SKU validado) |
| `PATCH` | `/api/v1/products/{id}` | Admin/Manager | Atualizar |
| `POST` | `/api/v1/products/{id}/deactivate` | Admin/Manager | Soft delete |

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `controller/ProductController.java` | 🆕 | 5 endpoints REST |

#### Skills: `302-frameworks-spring-boot-rest`, `304-frameworks-spring-boot-security`

---

### T-075 — Testes Unitários M6

- **Critério DONE:** ≥ 80% cobertura. RN17-01, RN17-04, RN18-02, RN18-04 testadas.
- **Estimativa:** 1.5d
- **Abordagem:** JUnit 5 + Mockito. Padrão AAA. Mocks de repositories.

#### Cenários

| Classe | Cenários | RNs |
|:---|:---:|:---|
| `BusinessUnitServiceTest` | create (CNPJ único, CNPJ duplicado→409, CNPJ inválido→400), update (CNPJ imutável→400, parentId inativo→422), deactivate (soft delete, libera CNPJ), findTree | RN17-01, RN17-02, RN17-04 |
| `ProductServiceTest` | create (SKU único, SKU duplicado→409), update, deactivate (soft delete), fiscalMappingStatus default | RN18-02, RN18-04, RN18-03 |
| `BusinessUnitRepositoryTest` | existsByCnpj, findByCnpj | — |
| `ProductRepositoryTest` | findByBusinessUnit, existsBySku | — |

#### Skills: `131-java-testing-unit-testing`, `130-java-testing-strategies`

---

### T-076 — Testes Integração M6

- **Critério DONE:** PostgreSQL real. Cross-tenant query não vaza BU. CNPJ único, hierarquia, SKU único, soft delete.
- **Estimativa:** 2d
- **Abordagem:** Testcontainers + PostgreSQL 17. `BaseIntegrationTest`. REST Assured para endpoints.

#### Cenários

| Classe | Cenários |
|:---|:---:|
| `BusinessUnitServiceIT` | CRUD com PostgreSQL real. Hierarquia: Matriz→Filial→Sub-filial. CNPJ reúso pós-soft-delete. parentId inativo→422. |
| `ProductServiceIT` | CRUD com PostgreSQL real. SKU único por BU. Segmentação por BU. |
| `BusinessUnitControllerIT` | 6 endpoints. @RequiresPermission (Admin/Manager/Operator). |
| `ProductControllerIT` | 5 endpoints. @RequiresPermission (Admin/Manager/Operator). |

#### Skills: `132-java-testing-integration-testing`, `322-frameworks-spring-boot-testing-integration-tests`

---

### T-077 — Testes Isolamento Multi-Tenant

- **Critério DONE:** Zero dados cross-tenant. 2 tenants + dados sobrepostos. Cada endpoint testado.
- **Estimativa:** 1.5d
- **Abordagem:** Testcontainers com 2 tenants. Tenant-A cria BUs e produtos. Tenant-B cria BUs e produtos. Verificar que Tenant-A nunca vê dados de Tenant-B.

#### Cenários

| ID | Descrição |
|:---|:---|
| TC-S6-MT-001 | Tenant-A não vê BUs de Tenant-B em `GET /business-units` |
| TC-S6-MT-002 | Tenant-A não vê produtos de Tenant-B em `GET /products` |
| TC-S6-MT-003 | Tenant-A tenta `POST product` em BU de Tenant-B → 404 |
| TC-S6-MT-004 | Tenant-A tenta `GET /business-units/{id-tenant-B}` → 404 |
| TC-S6-MT-005 | RLS V009: `SELECT * FROM product_service` como Tenant-A → apenas produtos Tenant-A |

#### Skills: `132-java-testing-integration-testing`, `124-java-secure-coding`

---

## 4. Ordem de Execução

```
1. T-069  BusinessUnitRepository (complementar)     ~0.5d
2. T-070  BusinessUnitService                        ~2d
3. T-071  BusinessUnitController                     ~1.5d
4. T-072  ProductRepository + RowMapper             ~0.5d
5. T-073  ProductService (service)                   ~2d
6. T-074  ProductController                          ~1d
7. T-075  Testes unitários                           ~1.5d
8. T-076  Testes integração                          ~2d
9. T-077  Testes isolamento multi-tenant             ~1.5d
```

**Justificativa:** BUs antes de Produtos (dependência conceitual — produto pertence a BU). Repository → Service → Controller → Testes (dependência técnica). Testes de isolamento por último (dependem de tudo funcionando).

---

## 5. Estratégia de Build e Verificação

| Checkpoint | Comando | Esperado |
|:---|:---|:---|
| Após T-069 | `./mvnw compile` | ✅ |
| Após T-071 | `./mvnw compile` | ✅ |
| Após T-074 | `./mvnw compile` | ✅ |
| Após T-075 | `./mvnw test` | ✅ Todos passando |
| Após T-076 | `./mvnw verify` | ✅ Testcontainers OK |
| Após T-077 | `./mvnw verify` | ✅ Isolamento confirmado |
| Final | `./mvnw clean verify` | ✅ BUILD SUCCESS |

---

## Rodapé

🤖 *Documento gerado em 23/07/2026 conforme PROMPT-EXECUTE-SPRINT-TASKS.md Fase 1. Stack: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Flyway 12.11.0 + Caffeine 3.2.4. 9 tasks, 2 features (F04-05, F04-06), 9 RNs, ~12.5 dias-homem estimados.*
