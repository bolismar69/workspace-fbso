# SPRINT-DEVELOPMENT-PLANNING-Frente-1.md — Plano de Desenvolvimento: Sprint 6 — Frente 1

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 6 de 7 — sprint-06-bus-catalogo
- **Frente:** Frente 1 — Recomendados
- **Stack:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Flyway 12.11.0 · Caffeine 3.2.4 · Keycloak 26
- **Data:** 23 de Julho de 2026
- **Branch:** `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-06-bus-catalogo`
- **Status Frente 0:** ✅ 4/4 concluídas. 261 testes (0 failures).

---

## 1. Visão Geral

- **Sprint Goal:** "Estrutura hierárquica de Unidades de Negócio (Matriz/Filial) com CNPJ único entre ativos. Soft delete libera CNPJ para reúso. Catálogo de Produtos/Serviços segmentado por BU com SKU único. Isolamento multi-tenant verificado."
- **Frente 1 Goal:** Resolver 5 débitos técnicos recomendados da auditoria (22 débitos totais). Estes débitos não bloqueiam o início das features M6 mas acumulam risco técnico se ignorados.
- **Tasks planejadas:** 5 (2 já concluídas, 3 pendentes)
- **Ordem de execução:** Sequencial (tasks independentes entre si)
- **Stack detectada:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Flyway 12.11.0 + Caffeine 3.2.4 + Keycloak 26
- **Fonte da stack:** PRD.md §5.1 + ARCHITECTURE.md header

### Status das Tasks

| ID | Tarefa | Débito | Status | Motivo |
|:---|:---|:---|:---:|:---|
| **T-165.DT-130** | V009 migration: RLS em product_service | DT-130 | ⬜ Pendente | — |
| **T-166.DT-131** | Remover hierarchyType de BusinessUnit.java | DT-131 | ✅ Já concluído | hierarchyType removido na F0 (DT-126). isMatrix já no toColumnMap(). |
| **T-167.DT-133** | Atualizar SPRINT-CARD.md | DT-133 | ✅ Já concluído | Docs da sprint atualizados na mesma sessão |
| **T-168.DT-134** | ADR-L08: Query hierárquica WITH RECURSIVE | DT-134 | ⬜ Pendente | — |
| **T-169.DT-137** | Externalizar trusted-proxy-ips no RateLimitFilter | DT-137 | ⬜ Pendente | — |

---

## 2. Dependências entre Tasks

```
Frente 0 (T-161..T-164) ✅
    │
    ├── T-165.DT-130 (V009 RLS)     ← Independente. Depende apenas de V003 (padrão)
    ├── T-168.DT-134 (ADR-L08)      ← Independente. Usa business_unit (já existe)
    └── T-169.DT-137 (RateLimit)    ← Independente. Refatora RateLimitFilter existente
```

As 3 tasks pendentes são independentes entre si — podem ser executadas em qualquer ordem. O padrão recomendado é:

1. **T-165** (migration) — mais crítico (fecha gap de segurança RLS)
2. **T-169** (config) — rápida, melhoria de infraestrutura
3. **T-168** (ADR + query) — documentação + código novo

---

## 3. Plano por Task

### T-165.DT-130 — V009 Migration: RLS em product_service

- **Critério DONE (SPRINT-CARD):** Migration V009 aplicada. `SELECT * FROM product_service` como tenant-A retorna apenas produtos do tenant-A (RLS filtra). U009 reverte.
- **Estimativa:** 1.5h
- **Débito original:** DT-130 — product_service sem tenant_id próprio → sem RLS → gap de isolamento
- **Severidade:** 🟡 High (risco de vazamento cross-tenant via product_service)

#### Abordagem

A tabela `product_service` atualmente NÃO possui coluna `tenant_id`. O isolamento é feito indiretamente via JOIN com `business_unit.tenant_id`. Isso é frágil — se um desenvolvedor esquecer o JOIN, dados de múltiplos tenants vazam.

**Solução:** Desnormalizar `tenant_id` em `product_service` (copiar de `business_unit.tenant_id`) e aplicar o mesmo padrão RLS das outras 4 tabelas (V003).

#### Passos da Migration V009

```sql
-- 1. Adicionar coluna tenant_id (nullable inicialmente)
ALTER TABLE fbso_platform.product_service ADD COLUMN tenant_id UUID;

-- 2. Preencher tenant_id via JOIN com business_unit
UPDATE fbso_platform.product_service ps
SET tenant_id = bu.tenant_id
FROM fbso_platform.business_unit bu
WHERE ps.business_unit_id = bu.id;

-- 3. Tornar NOT NULL após preenchimento
ALTER TABLE fbso_platform.product_service ALTER COLUMN tenant_id SET NOT NULL;

-- 4. Adicionar FK para business_unit.tenant_id (integridade referencial)
--    Nota: não podemos criar FK para tenant.id porque product_service não
--    referencia tenant diretamente. A constraint é garantida pelo JOIN.
--    Alternativa: trigger que valida tenant_id = business_unit.tenant_id

-- 5. ENABLE + FORCE RLS
ALTER TABLE fbso_platform.product_service ENABLE ROW LEVEL SECURITY;
ALTER TABLE fbso_platform.product_service FORCE ROW LEVEL SECURITY;

-- 6. Política de isolamento
CREATE POLICY tenant_isolation ON fbso_platform.product_service
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID)
    WITH CHECK (tenant_id = current_setting('app.current_tenant_id')::UUID);
```

#### Rollback U009

```sql
DROP POLICY IF EXISTS tenant_isolation ON fbso_platform.product_service;
ALTER TABLE fbso_platform.product_service NO FORCE ROW LEVEL SECURITY;
ALTER TABLE fbso_platform.product_service DISABLE ROW LEVEL SECURITY;
ALTER TABLE fbso_platform.product_service DROP COLUMN tenant_id;
```

#### Impacto no código Java

- `ProductService.java` entity: adicionar campo `tenantId` + getter/setter + `toColumnMap()` entry
- `ProductService` entity Javadoc: atualizar nota sobre ausência de tenant_id
- `ProductRepository` (quando criado em M6): usar `hasTenantColumn=true` — o `BaseRepository` aplicará filtro automático

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/resources/db/migration/V009__add_tenant_id_to_product_service.sql` | 🆕 | Migration: adiciona tenant_id, popula, RLS |
| `src/main/resources/db/migration/U009__rollback_tenant_id_product_service.sql` | 🆕 | Rollback: remove RLS, drop column |

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `entity/ProductService.java` | 🔄 | Adicionar campo `tenantId`, getter/setter, `toColumnMap()` |
| `ARCHITECTURE.md` | 🔄 | Atualizar contagem de tabelas RLS: 4→5 (V003/V009), nota em §5.3 |

#### Dependências
- V003 (padrão RLS) — já existe
- V001 (tabela product_service) — já existe
- business_unit com tenant_id populado — já existe

#### Riscos
| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| UPDATE com JOIN falhar se houver product_service órfão (business_unit_id sem BU correspondente) | Baixa | Médio | Verificar antes: `SELECT COUNT(*) FROM product_service WHERE business_unit_id NOT IN (SELECT id FROM business_unit WHERE deleted_dt IS NULL)` |
| V009 quebrar migration chain (V008 já aplicada em dev/staging) | Baixa | Alto | Testar `mvn flyway:migrate` em ambiente local com dados existentes |
| Coluna tenant_id afetar testes existentes | Média | Médio | Rodar `mvn test` após migration; atualizar ProductServiceTest se necessário |

#### Skills aplicáveis
- `313-frameworks-spring-db-migrations-flyway` — Padrões de migration Flyway
- `311-frameworks-spring-jdbc` — JdbcTemplate, BaseRepository
- `124-java-secure-coding` — Isolamento multi-tenant, RLS

---

### T-166.DT-131 — Remover hierarchyType de BusinessUnit.java ✅ JÁ CONCLUÍDO

- **Status:** ✅ Concluído durante a Frente 0 (DT-126)
- **Evidência:** `BusinessUnit.java` não possui campo `hierarchyType`, getter, setter, ou referência em `toColumnMap()`. `isMatrix` está presente em campo (linha 48) e `toColumnMap()` (linha 203).
- **Verificação:** `grep -rn "hierarchyType\|HierarchyType" src/` retorna apenas o Javadoc (linha 28) documentando a remoção — isso é correto.

**Ação necessária:** Nenhuma. Task concluída. Apenas marcar ✅ no SPRINT-CARD.md.

---

### T-167.DT-133 — Atualizar SPRINT-CARD.md ✅ JÁ CONCLUÍDO

- **Status:** ✅ Concluído durante a sessão de atualização de documentos
- **Evidência:** SPRINT-CARD.md atualizado com Frente 0 (4 tasks ✅) + Frente 1 (5 tasks). Header com status 🔄. Métricas atualizadas. Branch name correto.

**Ação necessária:** Nenhuma. Task concluída.

---

### T-168.DT-134 — ADR-L08: Query Hierárquica com PostgreSQL WITH RECURSIVE

- **Critério DONE (SPRINT-CARD):** ADR-L08 documentado. Query `WITH RECURSIVE bu_tree AS (...)` funcional. Teste com 3 níveis de hierarquia.
- **Estimativa:** 1h
- **Débito original:** DT-134 — documentar decisão de query hierárquica antes de implementar M6

#### Abordagem

A tabela `business_unit` tem estrutura de árvore via `parent_id` auto-referenciado. A query hierárquica é necessária para:
- Listar BUs em estrutura de árvore (Matriz → Filiais → Sub-filiais)
- Validar profundidade e integridade da hierarquia
- Navegar a estrutura para herança de configurações

**Decisão (ADR-L08):** Usar PostgreSQL `WITH RECURSIVE` CTE (Common Table Expression) em vez de:
- ❌ Múltiplas queries N+1 (performance degradada com árvores profundas)
- ❌ Carregar todos os registros e montar árvore em memória (não escala)
- ✅ WITH RECURSIVE — única query SQL, otimizada pelo PostgreSQL, sem limite de profundidade

#### ADR-L08 — Template (seguir padrão ARCHITECTURE.md §10)

| ID | Decisão | Justificativa |
|:---|:---|:---|
| **ADR-L08** | PostgreSQL WITH RECURSIVE para queries hierárquicas em business_unit | Performance O(n) em vez de O(n²) com N+1. Única query, transacionalmente consistente. Sem limite de profundidade (RN17-04). PostgreSQL otimiza CTE com índices em parent_id |

#### Implementação: BusinessUnitRepository.findTree()

```java
@Repository
public class BusinessUnitRepository extends BaseRepository<BusinessUnit> {

    public BusinessUnitRepository(JdbcTemplate jdbc) {
        super(jdbc, "business_unit", new BusinessUnitRowMapper(), true);
    }

    /**
     * Retorna a árvore hierárquica completa de BUs do tenant,
     * ordenada com Matriz primeiro, depois filhas por nível.
     *
     * ADR-L08: usa PostgreSQL WITH RECURSIVE para performance O(n).
     */
    public List<BusinessUnit> findTree(UUID tenantId) {
        String sql = """
            WITH RECURSIVE bu_tree AS (
                -- Caso base: Matriz (raiz)
                SELECT *, 0 AS depth, corporate_name AS sort_path
                FROM fbso_platform.business_unit
                WHERE tenant_id = ? AND parent_id IS NULL AND deleted_dt IS NULL

                UNION ALL

                -- Passo recursivo: filhas
                SELECT bu.*, bt.depth + 1,
                       bt.sort_path || ' > ' || bu.corporate_name
                FROM fbso_platform.business_unit bu
                INNER JOIN bu_tree bt ON bu.parent_id = bt.id
                WHERE bu.deleted_dt IS NULL
            )
            SELECT * FROM bu_tree
            ORDER BY sort_path
            """;
        return jdbc.query(sql, rowMapper, tenantId);
    }
}
```

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `repository/BusinessUnitRepository.java` | 🆕 | Repository com findTree() CTE recursiva |
| `repository/rowmapper/BusinessUnitRowMapper.java` | 🆕 | RowMapper para BusinessUnit |

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `ARCHITECTURE.md` §10 | 🔄 | Adicionar ADR-L08 na tabela de ADRs |

#### Dependências
- `BusinessUnit.java` entity — ✅ existe (F0, DT-126)
- `BaseRepository.java` — ✅ existe
- `AuditFieldsRowMapper.java` — ✅ padrão para criar BusinessUnitRowMapper

#### Riscos
| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| CTE com muitos níveis (>100) impactar performance | Baixa | Baixo | RN17-04 permite níveis ilimitados mas árvores reais têm ≤5 níveis. Índice em parent_id cobre |
| RowMapper novo duplicar lógica de AuditFieldsRowMapper | Média | Baixo | Reutilizar AuditFieldsRowMapper via composição (padrão dos outros RowMappers) |

#### Skills aplicáveis
- `030-architecture-adr-general` — Criação de ADR
- `311-frameworks-spring-jdbc` — JdbcTemplate, RowMapper
- `121-java-object-oriented-design` — Design do repository

---

### T-169.DT-137 — Externalizar trusted-proxy-ips no RateLimitFilter

- **Critério DONE (SPRINT-CARD):** IPs configuráveis por ambiente. Sem hardcode. Dev: 127.0.0.1. Prod: IP do proxy reverso.
- **Estimativa:** 30min
- **Débito original:** DT-137 — RateLimitFilter.extractKey() usa apenas request.getRemoteAddr(), ignorando proxies reversos

#### Abordagem

O `RateLimitFilter` atualmente extrai o IP do cliente via `request.getRemoteAddr()`. Quando atrás de um proxy reverso (Nginx, AWS ALB, Cloudflare), `getRemoteAddr()` retorna o IP do proxy, não do cliente real. Todos os clientes atrás do mesmo proxy compartilhariam o mesmo contador de rate limit — quebrando a funcionalidade.

**Solução:** Externalizar a lista de IPs de proxies confiáveis para `application.yml`. Quando a requisição vem de um IP confiável, usar o header `X-Forwarded-For` para obter o IP real do cliente.

#### Padrão a seguir

O projeto já usa externalização para `app.cors.allowed-origins`:

```yaml
# application.yml (linha 103-105)
app:
  cors:
    allowed-origins: ${CORS_ALLOWED_ORIGINS:http://localhost:3000,https://app.fbso.org}
```

```java
// SecurityConfig.java
@Value("${app.cors.allowed-origins:http://localhost:3000,https://app.fbso.org}")
private String allowedOrigins;
```

#### Implementação

**application.yml:**
```yaml
app:
  cors:
    allowed-origins: ${CORS_ALLOWED_ORIGINS:http://localhost:3000,https://app.fbso.org}
  rate-limit:
    trusted-proxy-ips: ${RATE_LIMIT_TRUSTED_PROXY_IPS:127.0.0.1,0:0:0:0:0:0:0:1}
```

**RateLimitFilter.java — alterações:**
- Adicionar campo `private final List<String> trustedProxyIps` injetado via `@Value`
- Modificar construtor para receber `List<String> trustedProxyIps`
- Atualizar `extractKey()` para verificar se `request.getRemoteAddr()` está na lista de confiáveis e, se sim, extrair de `X-Forwarded-For`

```java
// Construtor
public RateLimitFilter(ObjectMapper objectMapper,
                       @Value("${app.rate-limit.trusted-proxy-ips:127.0.0.1}") 
                       List<String> trustedProxyIps) {
    this.cache = Caffeine.newBuilder()...;
    this.objectMapper = objectMapper;
    this.trustedProxyIps = trustedProxyIps;
}

// extractKey atualizado
private String extractKey(HttpServletRequest request) {
    String remoteAddr = request.getRemoteAddr();
    if (trustedProxyIps.contains(remoteAddr)) {
        String forwardedFor = request.getHeader("X-Forwarded-For");
        if (forwardedFor != null && !forwardedFor.isBlank()) {
            // X-Forwarded-For: client, proxy1, proxy2
            return forwardedFor.split(",")[0].trim();
        }
    }
    return remoteAddr;
}
```

**SecurityConfig.java — alteração:**
- Atualizar criação do bean `RateLimitFilter` para injetar a propriedade

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `security/RateLimitFilter.java` | 🔄 | Adicionar trustedProxyIps, atualizar extractKey() |
| `config/SecurityConfig.java` | 🔄 | Atualizar bean RateLimitFilter com @Value |
| `src/main/resources/application.yml` | 🔄 | Adicionar app.rate-limit.trusted-proxy-ips |
| `TECHNICAL-REFERENCE.md` §8 | 🔄 | Adicionar RATE_LIMIT_TRUSTED_PROXY_IPS na tabela de variáveis |

#### Dependências
- Nenhuma — o RateLimitFilter e SecurityConfig já existem

#### Riscos
| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Spoofing de X-Forwarded-For se IP de proxy confiável for comprometido | Baixa | Médio | Apenas IPs de proxy na rede interna/VPC são confiáveis. Dev: localhost. Prod: IP do Nginx/ALB |
| Quebra de compatibilidade com testes existentes | Baixa | Baixo | Testes mockam HttpServletRequest — extractKey() é testável isoladamente |

#### Skills aplicáveis
- `304-frameworks-spring-boot-security` — Spring Security filter chain
- `124-java-secure-coding` — Proteção contra IP spoofing

---

## 4. Ordem de Execução

1. **T-165.DT-130 (V009 RLS)** — ~1.5h
   - Crítico: fecha gap de segurança RLS. Deve ser executado primeiro.
   - Criar V009 + U009 migrations → modificar ProductService.java → rodar migração → rodar testes

2. **T-169.DT-137 (RateLimit Config)** — ~30min
   - Rápido, independente. Melhoria de infraestrutura.
   - Modificar RateLimitFilter + SecurityConfig + application.yml → rodar testes

3. **T-168.DT-134 (ADR-L08 + findTree)** — ~1h
   - Documentação + código novo. Sem dependências externas.
   - Criar BusinessUnitRepository + RowMapper → atualizar ARCHITECTURE.md → rodar testes

**Justificativa da ordem:** Segurança primeiro (T-165), depois infraestrutura (T-169), depois documentação/código (T-168). Tasks são independentes — poderiam ser paralelizadas se houvesse múltiplos desenvolvedores.

---

## 5. Estratégia de Build e Verificação

| Checkpoint | Comando | Esperado |
|:---|:---|:---|
| Build inicial | `./mvnw clean compile` | ✅ SUCCESS |
| Migration V009 | `./mvnw flyway:migrate` | ✅ V009 aplicada sem erro |
| Testes unitários | `./mvnw test` | ✅ Todos passando (261+ testes) |
| Rollback V009 | `./mvnw flyway:undo` | ✅ U009 reverte limpo |
| Re-migrate | `./mvnw flyway:migrate` | ✅ V009 reaplicada |
| Build final | `./mvnw clean verify` | ✅ SUCCESS, testes passando |

### Checkpoints intermediários

1. **Após T-165:** `./mvnw flyway:migrate` + `./mvnw test` → V009 funcional, testes passando
2. **Após T-169:** `./mvnw test` → RateLimitFilterTest (se existir) passando, build OK
3. **Após T-168:** `./mvnw compile` → BusinessUnitRepository compila, ADR documentado

---

## 6. Tasks Concluídas (Sem Ação Necessária)

### T-166.DT-131 — Remover hierarchyType ✅

- **Concluído em:** Frente 0 (DT-126, 23/07/2026)
- **Evidência:** `BusinessUnit.java` — zero referências a `hierarchyType` (campo, getter, setter, toColumnMap). `isMatrix` presente no campo (linha 48) e `toColumnMap()` (linha 203).
- **Verificação:** `grep -rn "hierarchyType\|HierarchyType" src/` → apenas Javadoc documentando remoção ✅

### T-167.DT-133 — Atualizar SPRINT-CARD.md ✅

- **Concluído em:** 23/07/2026 (mesma sessão de atualização de documentos)
- **Evidência:** SPRINT-CARD.md com Frente 0 (4 tasks ✅), Frente 1 (5 tasks), métricas atualizadas, header com status 🔄

---

## Rodapé

🤖 *Documento gerado em 23/07/2026 conforme PROMPT-EXECUTE-SPRINT-TASKS.md Fase 1 — Planejamento do Desenvolvimento. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Flyway 12.11.0 + Caffeine 3.2.4. Skills aplicáveis: 313-frameworks-spring-db-migrations-flyway, 311-frameworks-spring-jdbc, 124-java-secure-coding, 304-frameworks-spring-boot-security, 030-architecture-adr-general, 121-java-object-oriented-design.*

*Próximo passo: Aguardando instruções para iniciar Fase 2 — Implementação.*
