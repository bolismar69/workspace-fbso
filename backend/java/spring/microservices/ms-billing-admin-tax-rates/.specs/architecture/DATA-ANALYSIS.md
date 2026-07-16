# Análise Consolidada de Schema — DATA-ANALYSIS.md

> **Schema:** `billing_tax_rates` (21 tabelas)
> **Microserviço:** ms-billing-admin-tax-rates (DT-1)
> **Data da Análise:** 11 de Julho de 2026
> **Especialistas:** PostgreSQL Code Review · Data Modeling · PostgreSQL Pro DBA
> **Escopo:** Análise multidisciplinar do modelo de dados para o PRJ-FIN-2026-0002

---

## Sumário Executivo

O schema `billing_tax_rates` foi submetido a três análises independentes e complementares — revisão técnica de código (PostgreSQL Code Review), validação de modelagem de dados contra requisitos de negócio (Data Modeling) e análise de performance e operação (PostgreSQL Pro DBA).

**Conclusão geral:** O modelo está **bem estruturado conceitualmente** para o domínio fiscal — padrões como vigência temporal (SCD Tipo 2), wildcard matching hierárquico, polimorfismo de auditoria e multi-tenancy por `empresa_id` são adequados. Entretanto, **a implementação concreta apresenta lacunas significativas** que precisam ser endereçadas antes do desenvolvimento:

- **3 gaps bloqueantes** que impedem o funcionamento correto do sistema
- **6 gaps de alta criticidade** que comprometem integridade, performance ou segurança
- **12 recomendações de melhoria** para elevar o schema ao nível *production-grade*

As três análises convergem em diversos pontos — como a necessidade de índices parciais para registros vigentes, a adoção de ENUMs para colunas de status, e a correção da trigger de fechamento de vigência — o que reforça a validade dessas recomendações.

---

## 1. Gaps Bloqueantes (Impedem Funcionamento)

### 🔴 GAP-01: Trigger `fechar_fim_validade_generica()` Quebrada

**Severidade:** CRÍTICA | **Análises:** Data Modeling, PostgreSQL Pro

**Problema:** A função PL/pgSQL `fechar_fim_validade_generica()` referencia a coluna `codigo_cst` — que **não existe** em nenhuma tabela de regras fiscais. Também usa `fim_validade` em vez de `final_validade`. A trigger **nunca funcionou**, o que significa que:

- Inserções de novas alíquotas **não fecham** a vigência da regra anterior
- Múltiplas linhas com `final_validade IS NULL` podem coexistir para a mesma chave de negócio
- A RN-01 (Conflito de Vigência) **não é garantida** em nível de banco
- O SCD Tipo 2 está **comprometido** — versões históricas não são demarcadas

**Recomendação:** Reescrever triggers específicas por tabela com cláusulas WHERE que reflitam as chaves de negócio reais. Para cargas em lote, desabilitar a trigger e usar UPDATE único no final do batch.

**DDL de Referência (exemplo para `icms_rules`):**
```sql
CREATE OR REPLACE FUNCTION fechar_vigencia_icms_rules()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE icms_rules
    SET final_validade = NEW.inicio_validade - 1
    WHERE empresa_id = NEW.empresa_id
      AND uf_origem = NEW.uf_origem
      AND uf_destino = NEW.uf_destino
      AND final_validade IS NULL
      AND id != NEW.id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```

---

### 🔴 GAP-02: Ausência de Tabela `usuarios`

**Severidade:** CRÍTICA | **Análises:** Data Modeling, PostgreSQL Code Review

**Problema:** O modelo referencia `usuario_id`, `usuario_nome`, `usuario_perfil` em `auditoria_log` e `enviado_por`/`aprovado_por` em `lotes_carga`, mas **não existe tabela `usuarios`**. Consequências:

- Sem integridade referencial para `auditoria_log.usuario_id` — qualquer INT pode ser inserido
- `lotes_carga.enviado_por` e `aprovado_por` são VARCHARs sem FK
- RN-10 a RN-13 (segregação de funções, perfis de acesso, rastreabilidade) **não podem ser implementadas**
- Mudança de nome de usuário quebra consultas históricas nos lotes

**Recomendação:** Criar tabela `usuarios` com: `id SERIAL PK`, `nome VARCHAR`, `email VARCHAR UNIQUE`, `perfil` (ENUM ou FK para `perfis`), `empresa_id FK`, `status`, `criado_em`, `atualizado_em`.

Em `auditoria_log`, manter `usuario_nome` e `usuario_perfil` como snapshots desnormalizados (justificável para imutabilidade de auditoria).

Em `lotes_carga`, substituir `enviado_por`/`aprovado_por` (VARCHAR) por `enviado_por_id`/`aprovado_por_id` (FK → `usuarios.id`).

---

### 🔴 GAP-03: Ausência de Tabelas Mestre de Classificações Fiscais

**Severidade:** CRÍTICA | **Análises:** Data Modeling

**Problema:** NCM, NBS, CClassTrib e CFOP são referenciados como strings VARCHAR nas tabelas de regras fiscais (`product_tax_exceptions.ncm`, `iva_dual_rules.ncm`, `ipi_regras.ncm`), mas **não existem tabelas mestre** para essas classificações. Consequências:

- RN-02 (Integridade de Classificação) — impossível validar que uma classificação existe
- RN-06 (Vinculação Protegida) — impossível detectar alíquotas vinculadas
- RN-07 (Código Único) — impossível prevenir duplicatas
- RN-08 (Formato de Código) — sem validação de formato (8 dígitos NCM, 4 dígitos CFOP)

**Recomendação:** Criar tabelas mestre:
```sql
CREATE TABLE ncm (
    codigo VARCHAR(8) PRIMARY KEY,
    descricao VARCHAR(500) NOT NULL,
    status VARCHAR(10) DEFAULT 'ATIVO',
    criado_em TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE cfop (
    codigo VARCHAR(4) PRIMARY KEY,
    descricao VARCHAR(500) NOT NULL,
    tipo VARCHAR(20), -- 'ENTRADA', 'SAIDA'
    status VARCHAR(10) DEFAULT 'ATIVO',
    criado_em TIMESTAMPTZ DEFAULT now()
);
```

Para NBS e CClassTrib, mesmo padrão adaptado. Adicionar FK das tabelas de regras para as tabelas mestre com `ON DELETE RESTRICT`.

---

## 2. Gaps de Alta Criticidade

### 🟠 GAP-04: Índices Parciais Ausentes para Registros Vigentes

**Severidade:** ALTA | **Análises:** PostgreSQL Code Review, PostgreSQL Pro DBA

**Problema:** O motor de cálculo (DT-3) **sempre** consulta regras com `final_validade IS NULL`. Índices sem a cláusula `WHERE final_validade IS NULL` incluem todo o histórico (50-80% de linhas expiradas), consumindo memória de buffer pool e reduzindo cache hit ratio.

**Recomendação:** Criar índices parciais em **todas** as 13 tabelas de regras fiscais:
```sql
-- icms_rules
CREATE INDEX idx_icms_vigentes ON icms_rules(empresa_id, uf_origem, uf_destino)
    WHERE final_validade IS NULL;

-- iva_dual_rules
CREATE INDEX idx_iva_vigentes ON iva_dual_rules(ncm, uf_destino, municipio_destino_ibge)
    WHERE final_validade IS NULL;

-- iss_rates
CREATE INDEX idx_iss_vigentes ON iss_rates(empresa_id, codigo_ibge, item_lista_servico)
    WHERE final_validade IS NULL;
```

---

### 🟠 GAP-05: Queries de Match Progressivo com `OR` em Múltiplas Colunas

**Severidade:** ALTA | **Análises:** PostgreSQL Pro DBA

**Problema:** `product_tax_exceptions` e `ipi_regras` usam queries com múltiplos `OR`s entre colunas com wildcards. O PostgreSQL não consegue usar índices B-tree eficientemente com `OR` entre colunas diferentes — força BitmapOr ou sequential scan.

**Recomendação:** Reescrever como `UNION ALL`:

```sql
-- Em vez de: WHERE (ncm = $1 OR ncm = $2) AND (uf_destino = $3 OR uf_destino = '*') ...
-- Usar:
SELECT * FROM product_tax_exceptions
WHERE empresa_id = $5 AND ncm = $1 AND uf_destino = $3 AND uf_origem = $4
  AND inicio_validade <= CURRENT_DATE AND final_validade IS NULL
UNION ALL
SELECT * FROM product_tax_exceptions
WHERE empresa_id = $5 AND ncm = $2 AND uf_destino = '*' AND uf_origem = '**'
  AND inicio_validade <= CURRENT_DATE AND final_validade IS NULL
ORDER BY (ncm = $1) DESC, (uf_destino = $3) DESC LIMIT 1;
```

Cada branch do UNION pode usar seu próprio índice.

---

### 🟠 GAP-06: CHECK Constraints Ausentes em Campos de Status

**Severidade:** ALTA | **Análises:** PostgreSQL Code Review (Recomendação #1), PostgreSQL Pro DBA (Recomendação #6)

**Problema:** Todas as colunas `status` são VARCHAR sem restrição. Valores inválidos (`APROVADOO`, `ativos`, `Aprovado`) podem ser inseridos, comprometendo integridade em sistema regulado por SOX.

**Recomendação:** Adotar ENUM types para status estáveis:
```sql
CREATE TYPE status_empresa AS ENUM ('ATIVA', 'INATIVA');
CREATE TYPE status_lote AS ENUM ('EM_VALIDACAO', 'AGUARDANDO_APROVACAO', 'APROVADO', 'REJEITADO');
CREATE TYPE status_lote_item AS ENUM ('ACEITO', 'REJEITADO', 'COM_ALERTA');
CREATE TYPE operacao_auditoria AS ENUM ('CRIACAO', 'EDICAO', 'DESATIVACAO', 'APROVACAO', 'REJEICAO');
CREATE TYPE entidade_auditavel AS ENUM ('ALIQUOTA', 'CLASSIFICACAO', 'REGIME', 'USUARIO', 'LOTE', 'EMPRESA');
```

---

### 🟠 GAP-07: Batch Processing sem Otimização para Carga em Lote

**Severidade:** ALTA | **Análises:** PostgreSQL Pro DBA (Recomendações #3 e #4)

**Problema:** Cada INSERT dispara `fechar_fim_validade_generica()`, gerando um UPDATE adicional. Para um lote de 5.570 municípios, seriam ~11.140 operações (INSERT + UPDATE). Em uma única transação, isso gera WAL massivo (50-100MB) e lock escalation.

**Recomendação:**
1. Desabilitar trigger para a sessão de batch
2. UPDATE único para fechar vigencias: `UPDATE iva_dual_rules SET final_validade = ... WHERE ... AND final_validade IS NULL`
3. INSERT em massa via `COPY` ou multi-row INSERT
4. Processar em batches de 500-1000 linhas com commit intermediário

---

### 🟠 GAP-08: UNIQUE Constraints Faltantes para CNPJ

**Severidade:** ALTA | **Análises:** PostgreSQL Code Review (Recomendação #3), Data Modeling

**Problema:** `empresas.cnpj_raiz`, `tenants.cnpj_completo` e `fornecedores.cnpj` não possuem UNIQUE constraints. CNPJs duplicados invalidariam todo o modelo de multi-tenancy.

**Recomendação:**
```sql
ALTER TABLE empresas ADD CONSTRAINT uk_empresas_cnpj_raiz UNIQUE (cnpj_raiz);
ALTER TABLE tenants ADD CONSTRAINT uk_tenants_cnpj UNIQUE (cnpj_completo);
ALTER TABLE fornecedores ADD CONSTRAINT uk_fornecedores_cnpj UNIQUE (cnpj);
```

---

### 🟠 GAP-09: Rastreabilidade de Auditoria Dependente de Código de Aplicação

**Severidade:** ALTA | **Análises:** PostgreSQL Code Review (Recomendação #7)

**Problema:** A `auditoria_log` depende de código Java/Spring para gerar registros. Operações via SQL direto ou ferramenta externa **não seriam registradas**, violando RN-14 e requisitos SOX.

**Recomendação:** Implementar triggers PL/pgSQL para todas as 8+ tabelas de regras fiscais:
```sql
CREATE OR REPLACE FUNCTION fn_audita_alteracao()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO auditoria_log (
        entidade_tipo, entidade_id, operacao, usuario_id,
        usuario_nome, estado_anterior, estado_novo, data_hora
    ) VALUES (
        TG_TABLE_NAME,
        CASE WHEN TG_OP = 'DELETE' THEN OLD.id ELSE NEW.id END,
        CASE TG_OP
            WHEN 'INSERT' THEN 'CRIACAO'
            WHEN 'UPDATE' THEN 'EDICAO'
            WHEN 'DELETE' THEN 'DESATIVACAO'
        END,
        current_setting('app.usuario_id')::int,
        current_setting('app.usuario_nome'),
        CASE WHEN TG_OP IN ('UPDATE','DELETE') THEN to_jsonb(OLD) ELSE NULL END,
        CASE WHEN TG_OP IN ('INSERT','UPDATE') THEN to_jsonb(NEW) ELSE NULL END,
        now()
    );
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;
```

---

## 3. Recomendações de Melhoria (Production-Grade)

| # | Recomendação | Severidade | Análises | Tabelas Afetadas |
|:---|:---|:---|:---|:---|
| 10 | **CHECK de range para alíquotas** (`0% ≤ valor ≤ 100%`) e vigência (`fim > início`) | MÉDIA | Code Review, Data Modeling | 13 tabelas de regras |
| 11 | **Row-Level Security (RLS)** via `empresa_id` para isolamento multi-tenant | MÉDIA | Code Review | 8 tabelas de regras fiscais |
| 12 | **BRIN index** em `auditoria_log.data_hora` (~180x menor que B-tree) | MÉDIA | PostgreSQL Pro | `auditoria_log` |
| 13 | **Particionamento mensal** (vs. anual) para `auditoria_log` | MÉDIA | Code Review, PostgreSQL Pro | `auditoria_log` |
| 14 | **EXCLUDE constraint** com `btree_gist` para prevenir overlapping de vigência | MÉDIA | PostgreSQL Pro | 13 tabelas de regras |
| 15 | **DOMAIN types** para `codigo_ibge`, `cnpj_raiz`, `cnpj_completo`, `cnae_fiscal` | MÉDIA | PostgreSQL Pro | Várias |
| 16 | **Materialized views** para dashboards (Entrega 4) | MÉDIA | PostgreSQL Pro | Reporting |
| 17 | **Tabela `regime_transition_map`** para mapeamento de correlação entre regimes | MÉDIA | Data Modeling | Nova tabela |
| 18 | **Padronizar `criado_em`/`atualizado_em`** como `TIMESTAMPTZ` | BAIXA | Code Review, Data Modeling | 15 tabelas de regras |
| 19 | **Substituir `ip_origem` VARCHAR → `inet`** para suporte IPv6 | BAIXA | Code Review | `auditoria_log` |
| 20 | **Padronizar nomenclatura** em português (domínio fiscal brasileiro) | BAIXA | Data Modeling | Todas |
| 21 | **`bigserial`** para `lotes_carga.id` e `lotes_carga_itens.id` | BAIXA | Code Review | `lotes_carga`, `lotes_carga_itens` |

---

## 4. Matriz de Convergência

Recomendações sinalizadas por mais de um especialista (maior confiança):

| Recomendação | Code Review | Data Modeling | PostgreSQL Pro |
|:---|:---:|:---:|:---:|
| CHECK constraints em status/alíquotas/vigência | ✅ | ✅ | ✅ |
| Índices parciais `WHERE final_validade IS NULL` | ✅ | — | ✅ |
| ENUM types para status | ✅ | — | ✅ |
| Criar tabela `usuarios` | ✅ | ✅ | — |
| UNIQUE em CNPJ | ✅ | ✅ | — |
| Particionamento `auditoria_log` | ✅ | — | ✅ |
| Padronizar `criado_em`/`atualizado_em` | ✅ | ✅ | — |
| Trigger `fechar_fim_validade_generica()` quebrada | — | ✅ | ✅ |

**Legenda:** ✅ = Sinalizado pelo especialista | — = Não abordado diretamente

---

## 5. Plano de Ação Recomendado

### Fase 0 — Correções Bloqueantes (antes de iniciar desenvolvimento)

1. ✅ Reescrever `fechar_fim_validade_generica()` — triggers por tabela (GAP-01)
2. ✅ Criar tabela `usuarios` + tabela `perfis` (GAP-02)
3. ✅ Criar tabelas mestre de classificações: `ncm`, `cfop`, `nbs`, `cclass_trib` (GAP-03)

### Fase 1 — Fundação (durante V1 de migrations Flyway)

4. ✅ Adicionar ENUM types + CHECK constraints de range (GAP-06)
5. ✅ Adicionar UNIQUE constraints para CNPJ (GAP-08)
6. ✅ Criar índices parciais `WHERE final_validade IS NULL` (GAP-04)
7. ✅ Criar índices para `empresa_id` nas 8 tabelas fiscais
8. ✅ Adicionar CHECK `final_validade > inicio_validade` onde ausente (GAP-04 Data Modeling)

### Fase 2 — Robustez (antes de produção)

9. ✅ Triggers de auditoria automática nas 8+ tabelas (GAP-09)
10. ✅ Otimização de batch processing (GAP-07)
11. ✅ RLS baseado em `empresa_id` (Rec #11)
12. ✅ DOMAIN types para formatos padronizados (Rec #15)

### Fase 3 — Performance e Operação (antes de carga real)

13. ✅ Reescrever queries de match progressivo como UNION ALL (GAP-05)
14. ✅ BRIN index em `auditoria_log.data_hora` (Rec #12)
15. ✅ Particionamento mensal da `auditoria_log` (Rec #13)
16. ✅ EXCLUDE constraints para vigência (Rec #14)
17. ✅ Ajustar autovacuum por tabela (PostgreSQL Pro #9)

### Fase 4 — Analytics (Entrega 4)

18. ✅ Materialized views para dashboards (Rec #16)
19. ✅ Tabela `regime_transition_map` para Período Híbrido (Rec #17)

---

## 6. Sequência de Migrations Flyway (Proposta)

```
V1__criar_enums.sql              — ENUM types para status, operacao, entidade
V2__criar_domains.sql            — DOMAIN types para codigo_ibge, cnpj
V3__criar_empresas.sql           — Tabela empresas
V4__criar_tenants.sql            — Tabela tenants
V5__criar_usuarios.sql           — Tabela usuarios + perfis
V6__criar_classificacoes.sql     — Tabelas ncm, cfop, nbs, cclass_trib
V7__criar_fornecedores.sql       — Tabela fornecedores
V8__criar_lotes.sql              — Tabelas lotes_carga + lotes_carga_itens
V9__criar_auditoria.sql          — Tabela auditoria_log (particionada)
V10__inserir_dados_referencia.sql — INSERTs de NCM, CFOP, CST, empresa default
V11__adicionar_multi_tenancy.sql — ALTER TABLE 8 tabelas (colunas nullable)
V12__preencher_empresa_id.sql    — UPDATE em batches (registros existentes)
V13__not_null_multi_tenancy.sql  — ALTER COLUMN SET NOT NULL
V14__adicionar_fks.sql           — Foreign Keys para empresa_id, usuario_id, classificacoes
V15__adicionar_unique.sql        — UNIQUE constraints (CNPJ, chaves compostas)
V16__adicionar_check.sql         — CHECK constraints (aliquotas, vigencia, status)
V17__adicionar_indices.sql       — Índices parciais + compostos
V18__reescrever_triggers.sql     — Novas triggers de fechamento de vigência
V19__adicionar_triggers_audit.sql — Triggers de auditoria automática
V20__adicionar_rls.sql           — Row-Level Security
```

---

## 7. Referências

- **ERD:** [erd.md](erd.md) — estrutura relacional, 21 tabelas
- **Dicionário de Dados:** [data-dictionary.md](data-dictionary.md) — função de negócio e padrões
- **Regras de Negócio:** [04-FEATURES.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/04-FEATURES.md) — 38 RNs
- **Mapa de Integrações:** [INTEGRATION-MAP.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md)
- **Project Charter:** [01-PROJECT-CHARTER.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/01-PROJECT-CHARTER.md)

---

🤖 *Análise consolidada gerada em 11 de Julho de 2026 por três agentes especialistas independentes: PostgreSQL Code Review, Data Modeling, e PostgreSQL Pro DBA.*
