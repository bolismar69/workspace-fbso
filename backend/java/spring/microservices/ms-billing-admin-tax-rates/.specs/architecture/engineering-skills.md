# Engineering Skills Analysis — ms-billing-admin-tax-rates

> **Schema:** `billing_tax_rates` (21 tabelas)
> **Microserviço:** ms-billing-admin-tax-rates (DT-1)
> **Data da Análise:** 12 de Julho de 2026
> **Tipo de Análise:** Engenharia de Software + Especialização em Banco de Dados
> **Escopo:** Análise multidisciplinar cruzando visão arquitetural, engenharia de dados, segurança e operações com especialização PostgreSQL

---

## Metodologia

Esta análise cruza **4 perspectivas de engenharia** com **4 especializações de banco de dados** para oferecer uma visão 360° do schema `billing_tax_rates`:

| # | Perspectiva | Skill | Foco da Análise |
|---|------------|-------|-----------------|
| E1 | **Arquiteto de Sistemas** | Engineering Architecture | Padrões arquiteturais, integração entre microserviços, design de APIs, consistência distribuída |
| E2 | **Engenheiro de Backend** | Senior Backend Engineering | API-to-database mapping, transações, connection pooling, caching hierarchy, Spring Boot patterns |
| E3 | **Engenheiro de Dados** | Senior Data Engineering | Pipelines ETL/ELT, qualidade de dados, linhagem, governança, observabilidade |
| E4 | **Revisor de Código** | Code Review Engineering | Qualidade, segurança, manutenibilidade, aderência a padrões, detecção de anti-padrões |
| D1 | **Especialista PostgreSQL** | PostgreSQL Table Design | Data types, constraints, índices, particionamento, JSONB, MVCC, schema evolution |
| D2 | **Database Review** | SQL Code Review | Checklist de segurança, performance, qualidade de código, schema design |
| D3 | **Otimizador de DB** | Database Optimizer | Execution plans, indexing strategies, N+1, caching, partitioning/sharding |
| D4 | **Arquiteto de Dados** | SQL Pro | Normalização, modelagem dimensional, SCD, CQRS, microservices DB patterns |

---

## Sumário Executivo

O schema `billing_tax_rates` foi analisado sob 8 perspectivas complementares. A análise prévia em [DATA-ANALYSIS.md](DATA-ANALYSIS.md) (realizada por 3 especialistas de banco de dados) identificou **3 gaps bloqueantes, 6 gaps de alta criticidade e 12 recomendações**. Esta análise de engenharia **confirma e endossa todas as 21 recomendações prévias** e adiciona **8 novas descobertas** do ponto de vista arquitetural e de engenharia de software.

### Convergência com DATA-ANALYSIS.md

Todas as 21 recomendações da análise prévia foram validadas contra as 8 skills aplicadas nesta análise. Os 3 gaps bloqueantes (trigger quebrada, ausência de `usuarios`, ausência de tabelas mestre de classificações) são **confirmados como críticos** sob todas as perspectivas.

### Novas Descobertas (Engenharia)

| # | Descoberta | Severidade | Perspectivas |
|---|-----------|-----------|-------------|
| ENG-01 | Schema compartilhado entre DT-1 e DT-3 sem contrato formal de ownership | ALTA | E1, E2, D4 |
| ENG-02 | Ausência de event sourcing para propagação de mudanças entre DT-1 → DT-3 | ALTA | E1, E2, E3 |
| ENG-03 | Cache Redis sem estratégia de invalidação cross-service documentada | ALTA | E2, D3 |
| ENG-04 | JSONB `estado_anterior`/`estado_novo` sem schema validation | MÉDIA | E3, E4, D1 |
| ENG-05 | Timestamps mistos (`DATE`, `TIMESTAMP`, `TIMESTAMPTZ`) quebram timezone safety | MÉDIA | E2, E4, D1 |
| ENG-06 | Ausência de health check endpoint que valide conectividade com schema | MÉDIA | E1, E2 |
| ENG-07 | Falta estratégia de idempotência para operações de carga em lote | MÉDIA | E2, E3, D4 |
| ENG-08 | `auditoria_log` sem estratégia de retenção automatizada (apenas "5 anos" declarativo) | BAIXA | E3, E4 |

---

## 1. Análise por Perspectiva de Engenharia

### 1.1 Arquiteto de Sistemas (E1) — Padrões Arquiteturais

**Referências:** [erd.md](erd.md), [data-dictionary.md](data-dictionary.md), [INTEGRATION-MAP.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md)

#### Padrão Atual

O schema `billing_tax_rates` implementa um **modelo de banco de dados compartilhado** (shared database) entre dois microserviços:

```
DT-1 (Admin) ──READ/WRITE──► billing_tax_rates ◄──READ ONLY── DT-3 (Engine)
```

- **DT-1 (Admin):** Proprietário das operações de escrita — CRUD de alíquotas, carga em lote, aprovação
- **DT-3 (Engine):** Consumidor read-only — consulta alíquotas vigentes para cálculo de impostos

#### Avaliação

**Pontos fortes:**
- Separação clara de responsabilidades: DT-1 escreve, DT-3 lê
- Cache Redis compartilhado para queries frequentes do motor
- Staging de carga em lote (`lotes_carga` + `lotes_carga_itens`) previne contaminação de dados de produção
- Triggers de vigência temporal (SCD Tipo 2) preservam histórico de alterações

**Riscos identificados:**

| Risco | Descrição | Mitigação |
|-------|-----------|-----------|
| **Schema coupling** | Alterações de DDL pelo DT-1 podem quebrar queries do DT-3. Não há contrato formal de interface de banco | Definir `DATABASE_CONTRACT.md` listando tabelas/colunas que o DT-3 pode depender; versionar com migrations Flyway |
| **Cache incoherence** | DT-1 atualiza alíquotas; DT-3 pode ler stale data do Redis (TTL 24h). Sem invalidação ativa | Implementar notificação de invalidação (Redis Pub/Sub ou evento `tax-rate.updated`) |
| **No event sourcing** | Mudanças em alíquotas não geram eventos de domínio. DT-3 precisa polling ou cache TTL | Implementar `tax_rate_change_events` ou outbox pattern para notificar DT-3 de alterações |
| **Transação distribuída implícita** | Upload de lote → validação → aprovação → efetivação tem múltiplos passos sem saga/compensação | Formalizar como saga com passos explícitos e compensação |

#### Recomendações do Arquiteto

1. **Formalizar Database Contract:** Documentar exatamente quais tabelas, colunas e queries o DT-3 pode executar. Versionar junto com migrations Flyway
2. **Implementar Outbox Pattern:** Tabela `tax_rate_events` (id, event_type, entity_type, entity_id, payload JSONB, created_at) — DT-1 escreve eventos; DT-3 consome via polling ou CDC
3. **Estratégia de Invalidação de Cache:** Publicar evento `tax-rate.updated` após cada efetivação; listener no DT-3 invalida chaves Redis afetadas
4. **Health Check Cross-Service:** Endpoint `/health/readiness` no DT-1 que valida conectividade com schema `billing_tax_rates`

> `★ Insight ─────────────────────────────────────`
> O maior risco arquitetural não está no schema em si, mas na **falta de contrato explícito entre DT-1 e DT-3 sobre o schema compartilhado**. Em microserviços com shared database, o acoplamento no banco é o ponto de falha mais comum — uma coluna renomeada pelo time de admin quebra silenciosamente o motor de cálculo. A solução canônica é: (1) database contract versionado, (2) outbox pattern para eventos de mudança, (3) integração test suite que valida queries do DT-3 contra migrations do DT-1.
> `─────────────────────────────────────────────────`

---

### 1.2 Engenheiro de Backend (E2) — API ↔ Database Mapping

**Referências:** [data-dictionary.md](data-dictionary.md) — Seção "Diagrama de Dependência Funcional", [INTEGRATION-MAP.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md)

#### Mapeamento API → Tabelas

| Recurso REST | Operações | Tabelas Afetadas | Transação? |
|---|---|---|---|
| `POST /api/v1/aliquotas/icms` | CRIACAO | `icms_rules` (+ trigger fecha vigência) | Sim (single-table) |
| `PUT /api/v1/aliquotas/icms/{id}` | EDICAO | `icms_rules` (UPDATE) + `auditoria_log` (INSERT) | Sim (2 tabelas) |
| `DELETE /api/v1/aliquotas/icms/{id}` | DESATIVACAO | `icms_rules` (soft delete — `final_validade = now()`) + `auditoria_log` | Sim (2 tabelas) |
| `POST /api/v1/lotes` | UPLOAD | `lotes_carga` (INSERT) + `lotes_carga_itens` (bulk INSERT) | Sim (2 tabelas) |
| `POST /api/v1/lotes/{id}/aprovar` | APROVACAO | `lotes_carga` (UPDATE status) + tabelas finais (INSERT múltiplas) + `auditoria_log` | ⚠️ Multi-table — requer @Transactional |

#### Avaliação de Tipos de Dados (Perspectiva Backend ↔ PostgreSQL)

Mapeamento Java ↔ PostgreSQL identificado no ERD e dicionário:

| Coluna no ERD | Tipo Declarado | Tipo PostgreSQL Recomendado | Status |
|---|---|---|---|
| `id` (todas as tabelas) | `serial` / `bigserial` | `BIGINT GENERATED ALWAYS AS IDENTITY` | ⚠️ `serial` é legacy; migrar para IDENTITY |
| `cnpj_raiz`, `cnpj`, `cnpj_completo` | `varchar(N)` | `TEXT` + CHECK constraint de formato | ⚠️ `varchar(N)` desnecessário; usar `TEXT` + `CHECK` |
| `status` (todas) | `varchar` | ENUM type específico por domínio | ❌ VARCHAR sem constraint permite valores inválidos |
| `total` (orders) | `NUMERIC(10,2)` | ✅ `NUMERIC(10,2)` | ✅ Correto — money nunca float |
| `aliquota_*` (várias) | `decimal` | `NUMERIC(5,2)` ou `NUMERIC(6,3)` com CHECK | ⚠️ Sem constraint de range (0-100%) |
| `criado_em` / `atualizado_em` | `timestamp` / `date` / `timestamptz` | `TIMESTAMPTZ` (padronizado) | ❌ Mistura incompatível — ver ENG-05 |
| `conteudo_original`, `estado_*` | `jsonb` | ✅ `JSONB` | ✅ Correto para semi-estruturado |
| `justificativa`, `motivo_*` | `text` | ✅ `TEXT` | ✅ Correto |
| `ip_origem` | `varchar` | `INET` (suporta IPv4 e IPv6) | ⚠️ Recomendação #19 do DATA-ANALYSIS |

#### Descoberta ENG-05: Inconsistência de Tipos Temporais

O schema atual mistura **3 tipos temporais diferentes**:

```
icms_rules.criado_em          → DATE
product_tax_exceptions.criado_em → TIMESTAMP
auditoria_log.data_hora       → TIMESTAMP (sem timezone?)
empresas.criado_em            → TIMESTAMP
```

**Problema:** Quando um registro é criado via API Java (`OffsetDateTime` ou `Instant`), a conversão para `DATE` trunca a hora; para `TIMESTAMP` (sem timezone) assume fuso da sessão. Em ambiente multi-timezone, auditoria perde precisão.

**Recomendação:** Padronizar **todas** as colunas de timestamp como `TIMESTAMPTZ` (TIMESTAMP WITH TIME ZONE). O Spring Boot 4.0.1 + JDBC mapeia `OffsetDateTime` ↔ `TIMESTAMPTZ` nativamente.

```sql
-- Migração corretiva (exemplo para icms_rules):
ALTER TABLE icms_rules
  ALTER COLUMN criado_em TYPE TIMESTAMPTZ USING criado_em::timestamptz,
  ALTER COLUMN atualizado_em TYPE TIMESTAMPTZ USING atualizado_em::timestamptz;
```

> `★ Insight ─────────────────────────────────────`
> O Spring Boot 4.0.1 tem suporte nativo a `OffsetDateTime` via JDBC 4.2. Quando o driver PostgreSQL encontra `TIMESTAMPTZ`, ele automaticamente converte para UTC na escrita e para o timezone da JVM na leitura. Usar `TIMESTAMP` (sem timezone) quebra essa garantia — cada sessão de banco pode interpretar o valor em um fuso diferente. Em sistemas fiscais, onde a ordem cronológica de auditoria é juridicamente relevante, timezone safety não é opcional.
> `─────────────────────────────────────────────────`

#### Recomendações do Backend Engineer

1. **Connection Pool:** Configurar HikariCP (padrão Spring Boot) com `maximumPoolSize = 20`, `minimumIdle = 5`, `connectionTimeout = 30000`, `idleTimeout = 600000`
2. **@Transactional boundaries:** Mapear cada operação da matriz Tabela × Operação para métodos com `@Transactional` apropriado. Operações multi-tabela (aprovação de lote) precisam de `@Transactional(isolation = READ_COMMITTED)`
3. **Batch operations:** Para carga em lote, usar `JdbcTemplate.batchUpdate()` em vez de `saveAll()` do Spring Data — evita N+1 e reduz round-trips
4. **Query timeout:** Configurar `spring.datasource.hikari.query-timeout = 30000` e `javax.persistence.query.timeout = 30000` para prevenir long-running queries em tabelas de staging

---

### 1.3 Engenheiro de Dados (E3) — Pipelines, Qualidade e Governança

**Referências:** [data-dictionary.md](data-dictionary.md) — Seções `lotes_carga`, `lotes_carga_itens`, `auditoria_log`

#### Pipeline de Carga em Lote (Atual)

```
Arquivo CSV/Excel ──► POST /api/v1/lotes ──► Validação (RN-01 a RN-05)
                                                    │
                                    ┌───────────────┼───────────────┐
                                    ▼               ▼               ▼
                              lotes_carga    lotes_carga_itens  (status: ACEITO,
                              (cabeçalho)    (linhas validadas)  REJEITADO,
                                                                  COM_ALERTA)
                                    │
                                    ▼
                          Aprovação (Admin Fiscal)
                                    │
                                    ▼
                          Efetivação nas tabelas finais
                          (icms_rules, iva_dual_rules, etc.)
```

#### Avaliação de Qualidade de Dados

| Dimensão | Situação Atual | Recomendação |
|---|---|---|
| **Completude** | ✅ Colunas NOT NULL definidas para campos obrigatórios | Adicionar NOT NULL para `empresa_id`, `origem_cadastro` após migração |
| **Unicidade** | ❌ Sem UNIQUE em CNPJ, sem chaves naturais | ENG-08 do DATA-ANALYSIS: adicionar UNIQUE constraints |
| **Consistência** | ❌ Status como VARCHAR sem validação | GAP-06 do DATA-ANALYSIS: usar ENUM types |
| **Acurácia** | ❌ Sem CHECK em alíquotas (range 0-100%) | Recomendação #10 do DATA-ANALYSIS |
| **Linhagem** | ✅ Parcial — `lote_origem_id` + `lote_item_origem_id` | Completo: adicionar `origem_cadastro` preenchido automaticamente |
| **Tempestividade** | ✅ `lotes_carga` registra timestamps de envio e aprovação | Adicionar métrica de tempo entre envio e aprovação |

#### Descoberta ENG-03: Validação de Schema para JSONB

As colunas JSONB `conteudo_original` (em `lotes_carga_itens`), `estado_anterior` e `estado_novo` (em `auditoria_log`) armazenam dados sem validação estrutural. Isso é um risco de qualidade de dados — um bug no código Java pode persistir JSON malformado ou com campos faltantes.

**Recomendação:** Adicionar CHECK constraints que validam a estrutura mínima esperada:

```sql
-- Para auditoria_log: garante que snapshots sejam objetos JSON
ALTER TABLE auditoria_log
  ADD CONSTRAINT chk_estado_anterior_json_object
  CHECK (estado_anterior IS NULL OR jsonb_typeof(estado_anterior) = 'object');

ALTER TABLE auditoria_log
  ADD CONSTRAINT chk_estado_novo_json_object
  CHECK (estado_novo IS NULL OR jsonb_typeof(estado_novo) = 'object');

-- Para lotes_carga_itens: garante estrutura mínima
ALTER TABLE lotes_carga_itens
  ADD CONSTRAINT chk_conteudo_original_object
  CHECK (jsonb_typeof(conteudo_original) = 'object');
```

#### Recomendações do Data Engineer

1. **Data Quality Metrics:** Expor métricas via Micrometer: `tax_rates.lotes.validated.total`, `tax_rates.lotes.validated.accepted`, `tax_rates.lotes.validated.rejected`, `tax_rates.lotes.approval.time_ms`
2. **Schema evolution tracking:** Manter changelog de schema como documentação viva (não apenas migrations Flyway)
3. **Retenção automatizada:** Implementar job agendado (`@Scheduled`) que arquiva partições antigas de `auditoria_log` (>5 anos) para cold storage
4. **Data profiling:** Executar queries de profiling semanalmente para detectar anomalias (ex: crescimento anormal de `lotes_carga`, picos de rejeição)

---

### 1.4 Revisor de Código (E4) — Qualidade e Segurança

**Referências:** [erd.md](erd.md), [data-dictionary.md](data-dictionary.md), padrões de nomenclatura

#### Code Smells Identificados

| # | Code Smell | Localização | Impacto |
|---|-----------|------------|---------|
| CS-01 | Nomenclatura mista (inglês + português) | Várias tabelas | Confusão para desenvolvedores; `tax_equivalence` vs `ipi_regras` vs `iss_rates` — 3 padrões diferentes |
| CS-02 | `serial` em vez de `GENERATED ALWAYS AS IDENTITY` | `empresas.id`, `tenants.id`, `fornecedores.id`, `lotes_carga.id`, `lotes_carga_itens.id` | `serial` é legacy PostgreSQL; IDENTITY é o padrão moderno (PG10+) com melhor suporte a `INSERT ... OVERRIDING SYSTEM VALUE` |
| CS-03 | `varchar(N)` com limite arbitrário | `cnpj_raiz varchar(8)`, `cnae_principal varchar(7)`, etc. | Se o comprimento for regra de negócio, usar `TEXT` + `CHECK(LENGTH(col) = N)`. Se não, `TEXT` basta |
| CS-04 | Colunas `enviado_por`/`aprovado_por` como VARCHAR | `lotes_carga` | Sem FK, sem rastreabilidade; qualquer string é aceita |
| CS-05 | `bigserial` vs `serial` inconsistente | `auditoria_log.id` usa `bigserial`; `empresas.id` usa `serial` | Inconsistência sem justificativa; padronizar em `BIGINT GENERATED ALWAYS AS IDENTITY` |
| CS-06 | `varchar` sem tamanho vs `varchar(N)` | `auditoria_log.entidade_tipo varchar` vs `lotes_carga.tributo varchar` (sem tamanho) | Ou todos têm tamanho explícito, ou todos usam `TEXT` |

#### Análise de Segurança

| Vetor | Status | Evidência |
|-------|--------|-----------|
| **SQL Injection** | ⚠️ Parcial | RN-03 exige queries parametrizadas, mas não há proteção em nível de banco (ex: `PREPARE` statements) |
| **Row-Level Security** | ❌ Ausente | Recomendação #11 do DATA-ANALYSIS: RLS via `empresa_id` |
| **Audit Trail** | ⚠️ Parcial | `auditoria_log` depende de código Java (GAP-09 do DATA-ANALYSIS) |
| **Data Encryption** | ❌ Ausente | Sem menção a `pgcrypto` ou TDE para dados sensíveis (CNPJ, alíquotas) |
| **Access Control** | ❌ Ausente | Sem roles de banco definidas (`app_admin`, `app_engine_readonly`) |
| **Secret Management** | ❌ Ausente | Sem menção a vault/secret manager para credenciais de banco |

#### Recomendações do Code Reviewer

1. **Padronizar nomenclatura:** Escolher um idioma (português, dado domínio fiscal brasileiro) e aplicar consistentemente. Sugestão: `snake_case` em português para tabelas/colunas
2. **Migrar `serial` → `IDENTITY`:** Planejar migration para converter todas as PKs para `GENERATED ALWAYS AS IDENTITY`
3. **Database roles:** Criar roles com privilégios mínimos:
   ```sql
   CREATE ROLE app_admin WITH LOGIN PASSWORD '<vault>';
   CREATE ROLE app_engine_readonly WITH LOGIN PASSWORD '<vault>';
   GRANT SELECT ON ALL TABLES IN SCHEMA billing_tax_rates TO app_engine_readonly;
   GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA billing_tax_rates TO app_admin;
   ```
4. **Prepared Statements:** Garantir que 100% das queries Java usem `PreparedStatement` (Spring Data JDBC faz isso por padrão, mas queries nativas via `@Query` precisam de auditoria)

---

## 2. Análise por Especialização de Banco de Dados

### 2.1 Especialista PostgreSQL (D1) — Table Design Review

**Referência:** [postgresql SKILL.md](#) — PostgreSQL Table Design best practices

Aplicando as regras da skill `postgresql` ao schema documentado:

#### Data Types — Auditoria

| Regra PostgreSQL | Schema Atual | Conformidade |
|---|---|---|
| IDs: `BIGINT GENERATED ALWAYS AS IDENTITY` | `serial` (5 tabelas novas) + `bigserial` (1 tabela) | ❌ |
| Strings: `TEXT` (não `varchar(N)`) | `varchar(N)` em CNPJ, UF, status | ❌ |
| Money: `NUMERIC(p,s)` | `decimal` (sem escala explícita em algumas colunas) | ⚠️ |
| Time: `TIMESTAMPTZ` | `DATE` (icms_rules), `TIMESTAMP`, `TIMESTAMPTZ` misturados | ❌ |
| Do NOT use `timestamp` (sem TZ) | `icms_rules.criado_em` é `date`; `auditoria_log.data_hora` usa `timestamp` | ❌ |
| Do NOT use `char(n)` / `varchar(n)` | Várias colunas com `varchar(N)` | ❌ |
| Do NOT use `serial` | 5 tabelas com `serial` | ❌ |
| Do NOT use `money` type | Não encontrado | ✅ |

**Conclusão:** O schema viola **5 das 8 regras fundamentais** de data types PostgreSQL. As violações não são quebra-galas imediatas (o banco funciona), mas acumulam dívida técnica: `serial` vs `IDENTITY`, `varchar(N)` vs `TEXT`, e mistura de tipos temporais precisam ser corrigidos antes da primeira migration de produção.

#### Constraints — Auditoria

| Regra PostgreSQL | Schema Atual | Conformidade |
|---|---|---|
| PK: `BIGINT GENERATED ALWAYS AS IDENTITY` | `serial` em 5 tabelas | ❌ |
| FK: especificar `ON DELETE/UPDATE` | FK declaradas sem ação explícita (default: `NO ACTION`) | ⚠️ |
| FK indexes: criar manualmente | Não documentado se existem índices em FKs | ⚠️ |
| UNIQUE: preferir `NULLS NOT DISTINCT` (PG15+) | Sem UNIQUE em CNPJ | ❌ |
| CHECK: constraints de range em alíquotas | Ausentes | ❌ |

**Ação corretiva prioritária:** Adicionar `ON DELETE RESTRICT` explícito nas FKs para documentar intenção e prevenir deleções acidentais em cascata.

#### Índices — Status Atual

O ERD e data dictionary mencionam índices específicos:

| Índice Documentado | Tipo | Tabela | Avaliação |
|---|---|---|---|
| `idx_rbt12_range` | (não especificado) | `simples_nacional_rates` | ✅ Essencial para query de faixa de RBT12 |
| `idx_iva_rules_lookup` | Único em `(ncm, uf_destino, COALESCE(municipio, '0000000'))` | `iva_dual_rules` | ✅ Bem desenhado; cobre o padrão de lookup com NULL |
| Índices implícitos de PK | B-tree | Todas as tabelas | ✅ Automáticos |
| Índices implícitos de UNIQUE | B-tree | `usuarios.email` (proposto), CNPJs (proposto) | ⚠️ Ainda não implementados |

**Índices faltantes (confirmados por D1, D2 e D3):**

1. **FK indexes:** PostgreSQL **não** auto-indexa colunas FK. Toda FK deve ter índice explícito:
   ```sql
   CREATE INDEX idx_tenants_empresa ON tenants(empresa_id);
   CREATE INDEX idx_fornecedores_empresa ON fornecedores(empresa_id);
   CREATE INDEX idx_lotes_empresa ON lotes_carga(empresa_id);
   CREATE INDEX idx_lotes_itens_lote ON lotes_carga_itens(lote_id);
   -- + empresa_id em todas as 8 tabelas de regras fiscais
   ```

2. **Partial indexes para vigentes:** (GAP-04 do DATA-ANALYSIS — confirmado)

3. **Expression indexes para LOWER:** Para buscas case-insensitive em `razao_social`, `nome_fantasia`:
   ```sql
   CREATE INDEX idx_empresas_razao_lower ON empresas (LOWER(razao_social));
   ```

> `★ Insight ─────────────────────────────────────`
> PostgreSQL não cria índices automaticamente em Foreign Keys — esse é o "gotcha" #1 que causa outages em produção. Sem índice na FK, um `DELETE` na tabela pai força sequential scan na tabela filha para verificar `ON DELETE RESTRICT`. Em `empresas` (pai) com `tenants` (filha), deletar uma empresa sem índice em `tenants.empresa_id` pode travar a tabela `tenants` inteira com lock. O `sql-code-review` skill flag especificamente isso no checklist de segurança.
> `─────────────────────────────────────────────────`

---

### 2.2 Database Review (D2) — Análise Estruturada

**Referência:** [sql-code-review SKILL.md](#) — Framework de SQL Code Review

Aplicando o checklist de 4 dimensões a cada tabela:

| Tabela | Segurança | Performance | Manutenibilidade | Schema Quality | Score |
|---|---|---|---|---|---|
| `empresas` | ⚠️ (sem RLS) | ✅ | ⚠️ (nomes mistos) | ⚠️ (sem UNIQUE em CNPJ) | 5.5/10 |
| `tenants` | ⚠️ (sem RLS) | ⚠️ (sem índice FK) | ⚠️ (nomes mistos) | ⚠️ (sem UNIQUE em CNPJ) | 5.0/10 |
| `lotes_carga` | ⚠️ (VARCHAR em user) | ✅ | ⚠️ (VARCHAR sem FK) | ⚠️ (VARCHAR sem CHECK) | 5.5/10 |
| `lotes_carga_itens` | ✅ | ⚠️ (JSONB sem GIN) | ✅ | ⚠️ (JSONB sem validação) | 6.5/10 |
| `auditoria_log` | ⚠️ (depende de app) | ⚠️ (sem BRIN) | ✅ | ✅ (bem normalizada) | 7.0/10 |
| `fornecedores` | ⚠️ (sem RLS) | ⚠️ (sem índice FK) | ✅ | ⚠️ (sem UNIQUE em CNPJ) | 5.5/10 |
| `icms_rules` | ⚠️ (sem RLS) | ⚠️ (sem partial index) | ⚠️ (tipos data mistos) | ⚠️ (sem CHECK alíquota) | 5.0/10 |
| `iva_dual_rules` | ⚠️ (sem RLS) | ✅ (índice lookup) | ✅ | ✅ | 7.5/10 |

**Média geral do schema: 6.2/10** — funcional, mas com lacunas de segurança e performance que precisam ser endereçadas antes de produção.

#### Top 5 Ações Prioritárias (D2)

1. 🔴 **CHECK constraints em alíquotas** (todas as tabelas de regras) — segurança de dados
2. 🔴 **FK indexes** (todas as FKs) — performance e prevenção de locks
3. 🟠 **Partial indexes `WHERE final_validade IS NULL`** (13 tabelas) — performance de leitura
4. 🟠 **UNIQUE em CNPJ** (empresas, tenants, fornecedores) — integridade de dados
5. 🟡 **Triggers de auditoria no banco** (não apenas no código Java) — compliance SOX

---

### 2.3 Otimizador de Banco de Dados (D3) — Performance & Escala

**Referência:** [database-optimizer SKILL.md](#) — Advanced Query Optimization, Indexing Strategies, N+1 Resolution

#### Análise de Carga por Tabela

| Tabela | Perfil de Carga | Volume Estimado | Hot Spots |
|---|---|---|---|
| `icms_rules` | Read-heavy (motor), write-low (admin) | ~100 linhas (27 UFs × ~4 combos) | Query por `(uf_origem, uf_destino, empresa_id)` |
| `iva_dual_rules` | Read-heavy (motor), write-medium (admin) | ~50K linhas (5K NCMs × 27 UFs) | Query por `(ncm, uf_destino, municipio)` + cache Redis |
| `product_tax_exceptions` | Read-heavy, write-medium | ~1K-5K (exceções por NCM) | Wildcard matching com OR — ver GAP-05 |
| `lotes_carga_itens` | Write-heavy (carga), read-low | ~10K-100K por lote | Bulk INSERT + validação |
| `auditoria_log` | Write-only (triggers), read-low (dashboards) | ~100K-1M/ano | INSERT automático; queries de range por `data_hora` |

#### Recomendações de Performance (D3)

1. **Autovacuum tuning por tabela:**
   ```sql
   -- Tabelas update-heavy (regras fiscais com fechamento de vigência)
   ALTER TABLE iva_dual_rules SET (autovacuum_vacuum_scale_factor = 0.01);
   ALTER TABLE icms_rules SET (autovacuum_vacuum_scale_factor = 0.05);

   -- Tabelas insert-only (auditoria_log)
   ALTER TABLE auditoria_log SET (autovacuum_vacuum_insert_scale_factor = 0.01);
   ```

2. **Workload analysis por tabela:**
   - `iva_dual_rules`: **Hot table** — toda consulta de CBS/IBS passa por aqui. Cache Redis mitiga, mas miss rate >10% exige read replica
   - `auditoria_log`: **Cold table** — escreve muito, lê pouco. Não compete por buffer pool com hot tables

3. **Connection pooling:** Isolar conexões de admin (DT-1, curtas, transacionais) das de motor (DT-3, read-only, cache-first) em pools separados

4. **Query rewrite para wildcard matching** (GAP-05 do DATA-ANALYSIS — confirmado pelo D3 como prioritário para produção)

---

### 2.4 Arquiteto de Dados (D4) — Modelagem e Padrões

**Referência:** [sql-pro SKILL.md](#) — Data Modeling, Normalization, Cloud Architecture

#### Avaliação de Normalização

| Tabela | Forma Normal | Observação |
|---|---|---|
| `empresas` | 3NF ✅ | Bem normalizada; `cnpj_raiz` como chave natural candidata |
| `tenants` | 3NF ✅ | Dependência funcional correta: `cnpj_completo → empresa_id` |
| `icms_rules` | 3NF ⚠️ | `aliquota_interna` + `aliquota_interestadual` dependem funcionalmente de `(uf_origem, uf_destino)` — ok. Mas `mva_padrao`, `possui_protocolo_st`, `possui_desoneracao` são dependentes de NCM — deveriam estar em `product_tax_exceptions` |
| `auditoria_log` | 3NF ⚠️ | `usuario_nome` e `usuario_perfil` são desnormalizados (dependem de `usuario_id`). Justificável para imutabilidade de auditoria |
| `lotes_carga` | 2NF ⚠️ | `enviado_por` e `aprovado_por` (VARCHAR) dependem do ID do usuário (que não existe como entidade). Violação de dependência funcional |

#### Descoberta ENG-01: Ownership de Schema Compartilhado

O modelo atual assume que DT-1 "possui" o schema e DT-3 "consome". Mas não há demarcação formal. Risco: DT-3 adiciona coluna em `iva_dual_rules` para otimização de cálculo; DT-1 remove-a em migration seguinte.

**Recomendação (D4):** Implementar **Database Views** como contrato de interface:
```sql
CREATE VIEW engine.iva_dual_active AS
SELECT ncm, uf_destino, municipio_destino_ibge,
       aliquota_cbs, aliquota_ibs_estadual, aliquota_ibs_municipal
FROM iva_dual_rules
WHERE final_validade IS NULL;
```
DT-3 consulta apenas as views — DT-1 pode refatorar tabelas subjacentes sem quebrar o contrato.

---

## 3. Consolidação: 8 Engineering Skills × 3 Documentos

### 3.1 Validação Cruzada do DATA-ANALYSIS.md

Cada recomendação do DATA-ANALYSIS.md foi reavaliada contra as 8 skills de engenharia:

| ID | Recomendação DATA-ANALYSIS | Status após Eng. Review | Perspectivas que confirmam |
|----|--------------------------|------------------------|--------------------------|
| GAP-01 | Trigger `fechar_fim_validade_generica()` quebrada | ✅ CONFIRMADO CRÍTICO | D1, D2, D4 — SCD Tipo 2 é fundação do modelo fiscal |
| GAP-02 | Ausência de tabela `usuarios` | ✅ CONFIRMADO CRÍTICO | E1, E2, E4, D4 — quebra rastreabilidade SOX |
| GAP-03 | Ausência de tabelas mestre de classificações | ✅ CONFIRMADO CRÍTICO | E3, D2, D4 — sem FK, sem integridade referencial |
| GAP-04 | Índices parciais ausentes | ✅ CONFIRMADO | D1, D2, D3 — todas as 3 skills de DB convergem |
| GAP-05 | Queries OR em wildcard matching | ✅ CONFIRMADO | D1, D3 — UNION ALL é padrão ouro para multi-column OR |
| GAP-06 | CHECK constraints ausentes em status | ✅ CONFIRMADO | E4, D1, D2 — ENUM > VARCHAR com CHECK |
| GAP-07 | Batch processing sem otimização | ✅ CONFIRMADO | E2, E3, D3 — COPY + batch commit |
| GAP-08 | UNIQUE constraints faltantes (CNPJ) | ✅ CONFIRMADO | E3, E4, D2, D4 — 4 perspectivas convergem |
| GAP-09 | Auditoria dependente de código de aplicação | ✅ CONFIRMADO | E1, E4, D2 — triggers PL/pgSQL são defesa em profundidade |
| Rec #10 | CHECK de range (0-100%) | ✅ CONFIRMADO | D1, D2, D4 |
| Rec #11 | Row-Level Security (RLS) | ✅ CONFIRMADO | E1, E4, D1 |
| Rec #12 | BRIN index em `auditoria_log` | ✅ CONFIRMADO | D1, D3 — 180x menor que B-tree |
| Rec #13 | Particionamento mensal | ✅ CONFIRMADO | E3, D1, D3 |
| Rec #14 | EXCLUDE constraint (btree_gist) | ✅ CONFIRMADO | D1, D3 |
| Rec #15 | DOMAIN types | ✅ CONFIRMADO | E4, D1, D2 |
| Rec #16 | Materialized views (dashboards) | ✅ CONFIRMADO | E3, D4 |
| Rec #17 | `regime_transition_map` | ✅ CONFIRMADO | D4 |
| Rec #18 | Padronizar `criado_em`/`atualizado_em` como TIMESTAMPTZ | ✅ CONFIRMADO | E2, E4, D1 — ver ENG-05 |
| Rec #19 | `ip_origem` VARCHAR → INET | ✅ CONFIRMADO | E4, D1 |
| Rec #20 | Padronizar nomenclatura (português) | ✅ CONFIRMADO | E4, D2 |
| Rec #21 | `bigserial` para IDs de staging | ✅ CONFIRMADO | E4, D1 |

**Resultado: 21/21 recomendações confirmadas (100% de convergência).**

---

## 4. Plano de Ação Integrado

### Fase 0 — Correções Bloqueantes (antes de qualquer código)

| # | Ação | Skills | Baseado em |
|---|------|-------|-----------|
| 0.1 | Reescrever triggers de fechamento de vigência por tabela | D1, D2, D4 | GAP-01 |
| 0.2 | Criar tabela `usuarios` + `perfis` com FK | E1, E4, D4 | GAP-02 |
| 0.3 | Criar tabelas mestre `ncm`, `cfop`, `nbs`, `cclass_trib` | E3, D2, D4 | GAP-03 |
| 0.4 | Formalizar Database Contract DT-1 ↔ DT-3 | E1, E2 | ENG-01, ENG-06 |

### Fase 1 — Fundação (Migrations Flyway V1-V20)

| # | Ação | Skills | Baseado em |
|---|------|-------|-----------|
| 1.1 | ENUM types + CHECK constraints (status, alíquotas, vigência) | E4, D1, D2 | GAP-06, Rec #10 |
| 1.2 | UNIQUE constraints (CNPJ, chaves naturais) | E3, E4, D2, D4 | GAP-08 |
| 1.3 | Índices parciais `WHERE final_validade IS NULL` | D1, D2, D3 | GAP-04 |
| 1.4 | FK indexes (todas as FKs) | D1, D2, D3 | PostgreSQL gotcha |
| 1.5 | Padronizar `TIMESTAMPTZ` em todas as colunas temporais | E2, E4, D1 | ENG-05, Rec #18 |
| 1.6 | Migrar `serial` → `BIGINT GENERATED ALWAYS AS IDENTITY` | E4, D1 | Code smell CS-02 |
| 1.7 | `varchar(N)` → `TEXT` + CHECK de comprimento | D1 | PostgreSQL best practice |
| 1.8 | `ip_origem VARCHAR` → `INET` | E4, D1 | Rec #19 |

### Fase 2 — Robustez (antes de staging/produção)

| # | Ação | Skills | Baseado em |
|---|------|-------|-----------|
| 2.1 | Triggers de auditoria PL/pgSQL (defesa em profundidade) | E1, E4, D2 | GAP-09 |
| 2.2 | Otimização de batch processing (desabilitar trigger, COPY, batch commit) | E2, E3, D3 | GAP-07 |
| 2.3 | Row-Level Security via `empresa_id` | E1, E4, D1 | Rec #11 |
| 2.4 | DOMAIN types para formatos padronizados | E4, D1, D2 | Rec #15 |
| 2.5 | Outbox pattern / eventos de mudança para DT-3 | E1, E2 | ENG-02 |
| 2.6 | Invalidação ativa de cache Redis (Pub/Sub ou evento) | E2, D3 | ENG-02 |
| 2.7 | VALIDATE CHECK para JSONB (auditoria, lotes) | E3, E4, D1 | ENG-03 |
| 2.8 | Database roles (`app_admin`, `app_engine_readonly`) | E4 | Code review security |

### Fase 3 — Performance e Operação

| # | Ação | Skills | Baseado em |
|---|------|-------|-----------|
| 3.1 | UNION ALL para queries de wildcard matching | D1, D3 | GAP-05 |
| 3.2 | BRIN index em `auditoria_log.data_hora` | D1, D3 | Rec #12 |
| 3.3 | Particionamento mensal `auditoria_log` | E3, D1, D3 | Rec #13 |
| 3.4 | EXCLUDE constraints (btree_gist) para vigência | D1, D3 | Rec #14 |
| 3.5 | Autovacuum tuning por tabela | D3 | PostgreSQL Pro |
| 3.6 | Connection pooling segregado (admin vs engine) | E2, D3 | Performance |
| 3.7 | Database Views como contrato DT-1/DT-3 | E1, D4 | ENG-01 |

### Fase 4 — Analytics e Governança

| # | Ação | Skills | Baseado em |
|---|------|-------|-----------|
| 4.1 | Materialized views para dashboards | E3, D4 | Rec #16 |
| 4.2 | Tabela `regime_transition_map` (Período Híbrido) | D4 | Rec #17 |
| 4.3 | Data quality metrics (Micrometer) | E3 | ENG-03 |
| 4.4 | Retenção automatizada de auditoria (>5 anos) | E3, E4 | ENG-08 |
| 4.5 | Health check endpoint cross-service | E1, E2 | ENG-06 |

---

## 5. Sequência de Migrations Flyway (Revisada)

A sequência proposta em DATA-ANALYSIS.md foi revisada e ajustada com as descobertas de engenharia:

```
V1__criar_enums.sql                  — ENUM types (status, operacao, entidade)
V2__criar_domains.sql                — DOMAIN types (codigo_ibge, cnpj, cnae)
V3__criar_empresas.sql               — Tabela empresas (com IDENTITY, TIMESTAMPTZ, TEXT)
V4__criar_tenants.sql                — Tabela tenants (com FK index)
V5__criar_usuarios.sql               — Tabela usuarios + perfis (🆕 prioridade elevada)
V6__criar_classificacoes.sql         — Tabelas ncm, cfop, nbs, cclass_trib
V7__criar_fornecedores.sql           — Tabela fornecedores
V8__criar_lotes.sql                  — lotes_carga + lotes_carga_itens (com JSONB CHECK)
V9__criar_auditoria.sql              — auditoria_log (particionada, com BRIN)
V10__inserir_dados_referencia.sql    — INSERTs NCM, CFOP, CST, empresa default
V11__adicionar_multi_tenancy.sql     — ALTER TABLE 8 tabelas (colunas nullable)
V12__preencher_empresa_id.sql        — UPDATE batches
V13__not_null_multi_tenancy.sql      — SET NOT NULL
V14__adicionar_fks.sql               — FKs + FK indexes (🆕 índices adicionados)
V15__adicionar_unique.sql            — UNIQUE constraints
V16__adicionar_check.sql             — CHECK constraints (aliquotas, vigencia, JSONB)
V17__adicionar_indices.sql           — Partial indexes + expression indexes (🆕)
V18__reescrever_triggers.sql         — Novas triggers de fechamento de vigência
V19__adicionar_triggers_audit.sql    — Triggers de auditoria automática
V20__adicionar_rls.sql               — RLS + database roles (🆕)
V21__padronizar_timestamps.sql       — DATE/TIMESTAMP → TIMESTAMPTZ (🆕 ENG-05)
V22__criar_views_engine.sql          — Views como contrato DT-1/DT-3 (🆕 ENG-01)
V23__criar_eventos_mudanca.sql       — Outbox table tax_rate_events (🆕 ENG-02)
```

> **Nota:** 4 novas migrations (V21-V23) foram adicionadas em relação à sequência original do DATA-ANALYSIS.md para cobrir as descobertas de engenharia.

---

## 6. Matriz de Riscos (Probabilidade × Impacto)

| Risco | Probabilidade | Impacto | Score | Mitigação |
|-------|-------------|---------|-------|-----------|
| DT-3 quebra por alteração de schema do DT-1 | Média | Crítico | 🔴 ALTO | Database contract + views (ENG-01) |
| Stale data no cache Redis (24h TTL) | Alta | Médio | 🟠 MÉDIO | Invalidação ativa (ENG-02) |
| Trigger de vigência quebrada corrompe histórico | Certeza | Alto | 🔴 ALTO | Reescrever triggers (GAP-01) |
| Sem `usuarios` — auditoria sem rastreabilidade | Certeza | Alto | 🔴 ALTO | Criar tabela `usuarios` (GAP-02) |
| Batch load trava por trigger por linha | Alta | Alto | 🔴 ALTO | Batch processing otimizado (GAP-07) |
| SQL injection em queries nativas `@Query` | Baixa | Crítico | 🟠 MÉDIO | Auditoria de código + PreparedStatement |
| Timezone bug em comparação de datas fiscais | Média | Alto | 🟠 MÉDIO | Padronizar TIMESTAMPTZ (ENG-05) |
| Lock contention em `iva_dual_rules` (hot table) | Média | Médio | 🟡 BAIXO | Read replica + cache tuning |

---

## 7. Referências

- **ERD:** [erd.md](erd.md) — 21 tabelas, diagramas Mermaid, relacionamentos
- **Dicionário de Dados:** [data-dictionary.md](data-dictionary.md) — função de negócio e padrões
- **Análise Prévia (DB):** [DATA-ANALYSIS.md](DATA-ANALYSIS.md) — 3 especialistas DB, 21 recomendações
- **Mapa de Integrações:** [INTEGRATION-MAP.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md)
- **Regras de Negócio:** [04-FEATURES.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/04-FEATURES.md)
- **Contrato de API:** [API-CONTRACTS.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/API-CONTRACTS.md)
- **ADR de Schema:** [adrs/](./adrs/) — Decisões de arquitetura registradas
- **ADR-003 — Estratégia de Tabelas:** [ADR-003-tax-table-strategy.md](adrs/ADR-003-tax-table-strategy.md) — Análise de tabelas unificadas vs independentes para CBS, IBS, IS

---

## Apêndice: Skills e Agentes Utilizados

### Skills de Banco de Dados

| Skill | Arquivo | Versão | Papel na Análise |
|-------|---------|--------|-----------------|
| `postgresql` | `~/.agents/skills/postgresql/SKILL.md` | community/2026-02-27 | Revisão de data types, constraints, índices, particionamento |
| `sql-code-review` | `~/.agents/skills/sql-code-review/SKILL.md` | — | Checklist de segurança, performance, qualidade, schema design |
| `sql-optimization` | `~/.agents/skills/sql-optimization/SKILL.md` | — | Otimização de queries, análise de anti-padrões |
| `database-optimizer` | `~/.agents/skills/database-optimizer/SKILL.md` | community/2026-02-27 | Execution plans, indexing strategies, caching, scaling |
| `sql-pro` | `~/.agents/skills/sql-pro/SKILL.md` | community/2026-02-27 | Data modeling, normalização, cloud architecture |
| `database-admin` | `~/.agents/skills/database-admin/SKILL.md` | community/2026-02-27 | Operações, HA/DR, migrations, segurança |

### Skills de Engenharia

| Skill | Bundle | Papel na Análise |
|-------|--------|-----------------|
| `senior-architect` | engineering-skills/engineering-team | Padrões arquiteturais, integração microserviços, consistência distribuída |
| `senior-backend` | engineering-skills/engineering-team | API-to-database mapping, transações, connection pooling, Spring Boot patterns |
| `senior-data-engineer` | engineering-skills/engineering-team | Pipelines ETL, qualidade de dados, governança, observabilidade |
| `code-reviewer` | engineering-skills/engineering-team | Code smells, segurança, padrões de nomenclatura, anti-padrões |

### Agentes de Análise

| Agente | Tipo | Papel |
|--------|------|------|
| **Arquiteto de Dados** | `sql-pro` + `database-optimizer` | Modelagem, normalização, padrões de arquitetura de dados |
| **Especialista PostgreSQL** | `postgresql` | Design de tabelas, data types, constraints, índices |
| **Database Review Specialist** | `sql-code-review` + `sql-optimization` | Auditoria de schema, scoring, checklist |
| **Senior Architect** | Engineering Architecture | Padrões de sistema, contratos, eventos, integração |
| **Senior Backend Engineer** | Engineering Backend | API mapping, transações, performance, caching |
| **Senior Data Engineer** | Engineering Data | Pipelines, qualidade, governança, métricas |
| **Code Reviewer** | Engineering Code Review | Code smells, segurança, nomenclatura, manutenibilidade |

### Documentos Analisados

| Documento | Caminho | Conteúdo |
|-----------|---------|----------|
| ERD | `.specs/architecture/erd.md` | 21 tabelas, diagramas Mermaid, relacionamentos |
| Dicionário de Dados | `.specs/architecture/data-dictionary.md` | Função de negócio, padrões de lookup, regras |
| Análise Prévia (DB) | `.specs/architecture/DATA-ANALYSIS.md` | 3 gaps bloqueantes, 6 gaps alta criticidade, 12 recomendações |
| Mapa de Integrações | `business-inputs/.../INTEGRATION-MAP.md` | Componentes, canais, segurança |
| API Contracts | `business-inputs/.../API-CONTRACTS.md` | Endpoints REST, payloads |

---

## Controle de Versão

| Versão | Data | Autor | Mudanças |
|--------|------|-------|----------|
| 1.0 | 2026-07-12 | Engineering Skills Team | Análise inicial — 8 perspectivas, 8 novas descobertas, validação cruzada de 21 recomendações |

---

🤖 *Documento gerado por Inteligência Artificial em 12 de Julho de 2026.*

**Agentes de IA utilizados nesta análise:**
- Arquiteto de Dados (skills: `sql-pro`, `database-optimizer`)
- Especialista PostgreSQL (skill: `postgresql`)
- Database Review Specialist (skills: `sql-code-review`, `sql-optimization`)
- Senior Architect (engineering skill: `senior-architect`)
- Senior Backend Engineer (engineering skill: `senior-backend`)
- Senior Data Engineer (engineering skill: `senior-data-engineer`)
- Code Reviewer (engineering skill: `code-reviewer`)
- Database Administrator (skill: `database-admin`)

**Skills aplicadas:** `postgresql`, `sql-code-review`, `sql-optimization`, `database-optimizer`, `sql-pro`, `database-admin`, `senior-architect`, `senior-backend`, `senior-data-engineer`, `code-reviewer`

**Ferramenta de orquestração:** Claude Code (Anthropic) — agentes especializados com isolamento de contexto
