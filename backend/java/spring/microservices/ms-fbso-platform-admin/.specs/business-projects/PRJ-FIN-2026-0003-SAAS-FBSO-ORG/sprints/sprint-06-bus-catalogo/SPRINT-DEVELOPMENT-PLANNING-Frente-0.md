# SPRINT-DEVELOPMENT-PLANNING-Frente-0.md — Plano de Desenvolvimento: Sprint 6 — Frente 0

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 6 de 7 — sprint-06-bus-catalogo
- **Frente:** Frente 0 — Correções Pré-Sprint (Bloqueantes)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Flyway 12.11.0 + Caffeine 3.2.4 + Keycloak 26
- **Data:** 23 de Julho de 2026
- **Origem:** [IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo.md) — 4 débitos bloqueantes (DT-126 a DT-129)
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-06-bus-catalogo`

---

## 1. Visão Geral

- **Sprint Goal (Frente 0):** Corrigir os 4 débitos bloqueantes que impedem o início do desenvolvimento das features M6 (Unidades de Negócio e Catálogo de Produtos).
- **Tasks a implementar:** 4 (T-161.DT-126 a T-164.DT-129)
- **Ordem de execução:** Paralela — as 4 tasks são independentes entre si
- **Estimativa total:** ~5.5h (≈1 dia)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JDBC Template (não JPA) + Flyway 12.11.0

### Stack Detectada

| Componente | Versão | Fonte |
|:---|:---|:---|
| Linguagem | Java 25 | PRD.md §5.1 |
| Framework | Spring Boot 3.5.14 | pom.xml |
| Persistência | JDBC Template (ADR-L01) | BaseRepository.java |
| Migrations | Flyway 12.11.0 | pom.xml |
| Cache | Caffeine 3.2.4 | pom.xml |
| Segurança | Spring Security + Keycloak 26 (JWT RS256) | SecurityConfig.java |
| Testes | JUnit 5 + Mockito + Testcontainers 1.21.4 | pom.xml |

### Skills Aplicáveis

| Skill | Task(s) | Justificativa |
|:---|:---|:---|
| `121-java-object-oriented-design` | T-161, T-162 | Design de entidades com BaseEntity, toColumnMap(), padrão de campos |
| `311-frameworks-spring-jdbc` | T-161, T-162, T-163 | RowMapper, BaseRepository, JdbcTemplate queries |
| `124-java-secure-coding` | T-163, T-164 | Validação de tenant isolation (IDOR prevention), validação de CNPJ |
| `126-java-exception-handling` | T-163 | TenantIsolationException, integração com GlobalExceptionHandler |
| `110-java-maven-best-practices` | T-164 | Estrutura de utilitários, convenções de nomenclatura |
| `ponytail` | Todas | Checklist YAGNI de 7 rungs — controle de escopo |

---

## 2. Dependências entre Tasks

```
T-161.DT-126 (BusinessUnit entity)  ─── independente ───┐
T-162.DT-127 (ProductService entity) ─── independente ───┤
T-163.DT-128 (validateBusinessUnit)  ─── independente ───┤── podem ser executadas em paralelo
T-164.DT-129 (CnpjValidator)         ─── independente ───┘
```

**Todas as 4 tasks são independentes** — sem dependências entre si. Podem ser executadas em qualquer ordem ou em paralelo.

### Dependências Externas (já satisfeitas)

| Task | Depende de | Status |
|:---|:---|:---|
| T-161 | BaseEntity, V001 schema, V007 migration | ✅ Existentes |
| T-162 | BaseEntity, V001 schema (product_service) | ✅ Existentes |
| T-163 | PermissionService, TenantContext | ✅ Existentes |
| T-164 | OnboardingService (para substituir isValidCnpj) | ✅ Existente |

---

## 3. Plano por Task

### T-161.DT-126 — Reescrever BusinessUnit.java

- **Critério DONE:** Entity compila. `toColumnMap()` contém apenas colunas existentes no DB. `BaseRepository.save()` e `update()` funcionam sem SQL error.
- **Estimativa:** 2h
- **Abordagem:**
  1. Analisar schema V001 da tabela `business_unit` (colunas reais)
  2. Reescrever a entity alinhando 100% com o DB:
     - Adicionar campos faltantes: `corporateName` (String), `taxRegime` (String → TaxRegime enum), `street`, `number`, `complement`, `neighborhood`, `city`, `state`, `zipCode`, `status` (String)
     - Renomear `name` → `corporateName`
     - Remover `hierarchyType` (não existe no DB)
     - Manter `isMatrix` (V007 adicionou a coluna)
     - Manter `parentId`, `tenantId`, `cnpj`
  3. Atualizar `toColumnMap()` com todas as 16 colunas de domínio
  4. Atualizar getters/setters
  5. Verificar que `BaseRepository.save()` e `update()` geram SQL com colunas corretas

- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/com/fbso/platform/admin/entity/BusinessUnit.java` | 🔄 | Reescrever com campos corretos: adicionar corporateName, taxRegime, street, number, complement, neighborhood, city, state, zipCode, status. Remover name (→corporateName), hierarchyType. Manter isMatrix. Atualizar toColumnMap() |

- **Dependências:** Nenhuma (task independente)
- **Riscos:**
  - `toColumnMap()` com coluna errada → SQL error no save/update → mitigado pela verificação contra V001
  - `TaxRegime` enum já existe (SIMPLES_NACIONAL, LUCRO_REAL, LUCRO_PRESUMIDO) → usar como tipo do campo taxRegime
  - `Address.java` existe mas é código morto → NÃO usar; manter campos de endereço inline como está no schema
- **Skills aplicáveis:** `121-java-object-oriented-design`, `311-frameworks-spring-jdbc`

---

### T-162.DT-127 — Criar ProductService.java Entity

- **Critério DONE:** Entity compila. `toColumnMap()` cobre todas as colunas do V001. Pronta para ser usada pelo ProductRepository.
- **Estimativa:** 1h
- **Abordagem:**
  1. Analisar schema V001 da tabela `product_service`:
     - Colunas de domínio: `business_unit_id` (UUID FK), `name` (VARCHAR 255), `sku` (VARCHAR 50, nullable), `type` (VARCHAR 20, default 'SERVICE'), `description` (TEXT, nullable), `status` (VARCHAR 30, default 'ACTIVE')
     - Colunas de auditoria: `created_dt`, `updated_dt`, `created_by`, `updated_by`, `deleted_dt`, `deleted_by`
  2. Criar classe seguindo o padrão das entities existentes (Plan.java, User.java como referência):
     - `extends BaseEntity`
     - Campos privados com getters/setters
     - `getId()` / `setId()` para infraestrutura BaseRepository
     - `toColumnMap()` com `LinkedHashMap` (apenas colunas de domínio)
     - `toString()` para debugging
  3. Usar `ProductType` enum existente (PRODUCT, SERVICE) para o campo `type`
  4. **NOTA:** `product_service` NÃO tem `tenant_id` — o isolamento é via JOIN com `business_unit`. O repository usará `hasTenantColumn=false`

- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/com/fbso/platform/admin/entity/ProductService.java` | 🆕 | Entidade Javabean extends BaseEntity: id, businessUnitId, name, sku, type (ProductType), description, status. toColumnMap() com 6 colunas |

- **Dependências:** Nenhuma (task independente)
- **Riscos:**
  - `ProductType` enum: verificar se tem valores PRODUCT e SERVICE → ✅ confirmado
  - Schema sem tenant_id: entity NÃO terá campo tenantId → repository usará `hasTenantColumn=false`
- **Skills aplicáveis:** `121-java-object-oriented-design`, `311-frameworks-spring-jdbc`

---

### T-163.DT-128 — Implementar validateBusinessUnitTenant()

- **Critério DONE:** Teste: tenant-A tenta atribuir role em BU do tenant-B → `TenantIsolationException`. Teste: mesmo tenant → sucesso.
- **Estimativa:** 1h
- **Abordagem:**
  1. Localizar o método `assignRole()` em `PermissionService.java` (linha ~217)
  2. Implementar `validateBusinessUnitTenant(UUID businessUnitId)`:
     - Query: `SELECT tenant_id FROM fbso_platform.business_unit WHERE id = ? AND deleted_dt IS NULL`
     - Executar via `JdbcTemplate.queryForObject()`
     - Comparar `tenant_id` retornado com `TenantContext.getTenantId()`
     - Se diferentes → lançar `TenantIsolationException` (já existe, mapeada para 403)
     - Se BU não encontrada → lançar `TenantNotFoundException` (já existe, mapeada para 404)
  3. Integrar no método `assignRole()` — chamar `validateBusinessUnitTenant(businessUnitId)` antes de `permissionRepo.assign()`
  4. Seguir o padrão de `validateUserTenant()` já existente no mesmo serviço

- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/com/fbso/platform/admin/service/PermissionService.java` | 🔄 | Adicionar método `validateBusinessUnitTenant(UUID)` + integrar no `assignRole()` |

- **Dependências:** Nenhuma (task independente — BusinessUnit já existe como entidade e tabela)
- **Riscos:**
  - `JdbcTemplate` precisa ser injetado no PermissionService (verificar se já está disponível)
  - `TenantIsolationException` já existe e é mapeada para 403 no GlobalExceptionHandler ✅
  - `TenantNotFoundException` já existe ✅
- **Skills aplicáveis:** `124-java-secure-coding`, `126-java-exception-handling`, `311-frameworks-spring-jdbc`

---

### T-164.DT-129 — Criar CnpjValidator

- **Critério DONE:** CNPJ "00.000.000/0000-00" → inválido. CNPJ "11.222.333/0001-81" → válido (dígitos corretos). `OnboardingService.completeStep2()` usa `CnpjValidator.isValid()`.
- **Estimativa:** 1.5h
- **Abordagem:**
  1. Criar classe `CnpjValidator` no pacote `utils/`
  2. Implementar algoritmo oficial de validação de CNPJ:
     - Remover caracteres não numéricos (regex `[^0-9]`)
     - Verificar comprimento = 14 dígitos
     - Rejeitar sequências com todos dígitos iguais (ex: 00.000.000/0000-00)
     - Calcular 1º dígito verificador (pesos 5,4,3,2,9,8,7,6,5,4,3,2)
     - Calcular 2º dígito verificador (pesos 6,5,4,3,2,9,8,7,6,5,4,3,2)
     - Comparar com dígitos informados
  3. Método público: `public static boolean isValid(String cnpj)`
  4. Substituir `OnboardingService.isValidCnpj()` (método privado, linha 134) por chamada a `CnpjValidator.isValid()`
  5. Adicionar validação de formato com máscara: `public static boolean isValidFormatted(String cnpj)` — regex `^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`

- **Arquivos a criar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/com/fbso/platform/admin/utils/CnpjValidator.java` | 🆕 | Classe utilitária com validação de CNPJ (formato + dígitos verificadores) |

- **Arquivos a modificar:**
  | Arquivo | Tipo | Descrição |
  |:---|:---|:---|
  | `src/main/java/com/fbso/platform/admin/service/OnboardingService.java` | 🔄 | Substituir método privado `isValidCnpj()` por `CnpjValidator.isValid()` |

- **Dependências:** Nenhuma (task independente)
- **Riscos:**
  - CNPJ com máscara vs sem máscara: aceitar ambos os formatos e normalizar internamente
  - O algoritmo de dígitos verificadores é determinístico — teste com CNPJs conhecidos válidos
  - `OnboardingService.isValidCnpj()` é `private` — pode ser removido sem impacto em outras classes
- **Skills aplicáveis:** `124-java-secure-coding`, `110-java-maven-best-practices`

---

## 4. Ordem de Execução

1. **T-161.DT-126** — BusinessUnit.java (2h) — maior complexidade, maior risco
2. **T-162.DT-127** — ProductService.java (1h) — rápida, padrão bem definido
3. **T-163.DT-128** — validateBusinessUnitTenant (1h) — depende de JdbcTemplate disponível
4. **T-164.DT-129** — CnpjValidator (1.5h) — requer algoritmo de dígitos verificadores

> 💡 As 4 tasks são independentes e podem ser executadas em paralelo. A ordem acima é sugerida para execução sequencial (uma pessoa), priorizando a task de maior risco (T-161) primeiro.

---

## 5. Estratégia de Build e Verificação

- **Comando de build:** `./mvnw clean compile`
- **Comando de teste rápido:** `./mvnw test -pl . -Dtest="**/unit/**/*Test"` (apenas unitários para feedback rápido)
- **Comando de teste completo:** `./mvnw verify` (unit + integration, requer Docker)
- **Checkpoints:**
  1. Após T-161: `./mvnw compile` — verificar que BusinessUnit compila sem erros
  2. Após T-162: `./mvnw compile` — verificar que ProductService compila sem erros
  3. Após T-163: `./mvnw compile` — verificar que PermissionService compila
  4. Após T-164: `./mvnw compile` — verificar que OnboardingService compila com CnpjValidator
  5. Checkpoint final: `./mvnw test` — todos os 227 testes existentes continuam passando

---

## 6. Critérios de Aceite da Frente 0

- [ ] `BusinessUnit.java` compila e `toColumnMap()` contém apenas colunas do V001
- [ ] `ProductService.java` criada e compila seguindo padrão BaseEntity
- [ ] `validateBusinessUnitTenant()` implementado e integrado no `assignRole()`
- [ ] `CnpjValidator.isValid()` valida CNPJs com dígitos verificadores corretos
- [ ] `OnboardingService` usa `CnpjValidator.isValid()` em vez do método privado
- [ ] `./mvnw clean compile` — BUILD SUCCESS
- [ ] `./mvnw test` — 227 testes passando (1 erro pré-existente permitido: SubscriptionServiceTest)

---

🤖 *Documento gerado em 23/07/2026. Fase 1 do PROMPT-EXECUTE-SPRINT-TASKS. Skills acionadas: 121-java-object-oriented-design, 311-frameworks-spring-jdbc, 124-java-secure-coding, 126-java-exception-handling, 110-java-maven-best-practices, ponytail.*
