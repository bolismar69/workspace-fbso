# ADR-003: Estratégia de Tabelas de Impostos — Unificadas vs Independentes

> **ADR Number:** 003
> **Data:** 2026-07-12
> **Status:** Proposta (Pending Review)
> **Decisores:** Arquiteto de Dados, Especialista PostgreSQL, Engineering Architect
> **Escopo:** Schema `billing_tax_rates` — tabelas de configuração de impostos corporativos multi-tenant
> **Referências:** [erd.md](../erd.md), [data-dictionary.md](../data-dictionary.md), [DATA-ANALYSIS.md](../DATA-ANALYSIS.md), [engineering-skills.md](../engineering-skills.md)

---

## Contexto

O schema `billing_tax_rates` precisa armazenar configurações de **8 impostos** que se dividem em dois regimes com ciclos de vida sobrepostos:

```
┌──────────────────────────────────────────────────────────────────────────────┐
│                         LINHA DO TEMPO DE TRANSIÇÃO                           │
├──────────────────────────────────────────────────────────────────────────────┤
│  2026          2027          2029          2032          2033                 │
│   │             │             │             │             │                   │
│   │ PIS/COFINS  │             │             │             │                   │
│   │ ICMS        │─────────────│─────────────│  extinto    │                   │
│   │ ISS         │─────────────│─────────────│  extinto    │                   │
│   │ IPI         │             │             │             │ (mantido)         │
│   │ CBS (0.1%)  │ CBS (pleno) │             │             │                   │
│   │ IBS (0.9%)  │             │ IBS (cresc) │ IBS (pleno) │                   │
│   │ IS          │             │             │             │                   │
│   │             │             │             │             │                   │
│   └─ Shadow Run └─ CBS Live ──└─ Híbrido ───└─ Full IVA ──                   │
└──────────────────────────────────────────────────────────────────────────────┘
```

### O Problema

O design atual adota uma **tabela unificada** para o novo regime (`iva_dual_rules` — CBS + IBS + IS) enquanto o regime antigo usa tabelas majoritariamente independentes (`icms_rules`, `ipi_regras`, `iss_rates`). A questão é:

> **Devemos manter a configuração dos impostos em tabelas unificadas (`iva_dual_rules`), ou separá-las em tabelas independentes por tributo (`cbs_rates`, `ibs_rates`, `is_rates`)?**

---

## Estado Atual: Mapeamento Imposto → Tabela

### Regime Atual (Pré-Reforma)

```
┌─────────────────────────────────────────────────────────────────┐
│                     REGIME ATUAL (PRÉ-REFORMA)                   │
├──────────┬──────────────────────────────────────────────────────┤
│ PIS      │ ──► federal_tax_rules (com COFINS)                   │
│          │ ──► product_tax_exceptions (com ICMS, COFINS)        │
├──────────┼──────────────────────────────────────────────────────┤
│ COFINS   │ ──► federal_tax_rules (com PIS)                      │
│          │ ──► product_tax_exceptions (com ICMS, PIS)           │
├──────────┼──────────────────────────────────────────────────────┤
│ IPI      │ ──► ipi_regras                         ⬅ Independente│
├──────────┼──────────────────────────────────────────────────────┤
│ ICMS     │ ──► icms_rules                         ⬅ Independente│
│          │ ──► product_tax_exceptions (com PIS, COFINS)         │
│          │ ──► tax_equivalence                                  │
│          │ ──► simples_nacional_rates                           │
├──────────┼──────────────────────────────────────────────────────┤
│ ISS      │ ──► iss_rates                          ⬅ Independente│
└──────────┴──────────────────────────────────────────────────────┘
```

### Novo Regime (Reforma Tributária)

```
┌─────────────────────────────────────────────────────────────────┐
│                    NOVO REGIME (REFORMA TRIBUTÁRIA)              │
├──────────┬──────────────────────────────────────────────────────┤
│ CBS      │ ──► iva_dual_rules (com IBS, IS)  ⬅ Unificado       │
│          │ ──► cbs_rates (fallback setorial)                    │
├──────────┼──────────────────────────────────────────────────────┤
│ IBS      │ ──► iva_dual_rules (com CBS, IS)  ⬅ Unificado       │
├──────────┼──────────────────────────────────────────────────────┤
│ IS       │ ──► iva_dual_rules (com CBS, IBS) ⬅ Unificado       │
│          │ ──► ncm_seletivo (catálogo NCMs)                     │
├──────────┼──────────────────────────────────────────────────────┤
│ CBS+IBS  │ ──► cst_reforma (tabela de referência)               │
│ (legado) │ ──► reforma_tributaria_rules     ⬅ Deprecated       │
└──────────┴──────────────────────────────────────────────────────┘
```

---

## Análise: Dimensões de Comparação

### 1. Dimensionalidade de Lookup

Este é o fator técnico mais importante. Cada imposto tem **chaves naturais diferentes** — o que define a unicidade de uma linha:

| Imposto | Chave Natural (Natural Key) | Dimensões | Exemplo |
|---------|---------------------------|-----------|---------|
| **CBS** | `(ncm, classe_tributaria)` | 2 | NCM 8471.30.19 + classe TELECOM → 8.8% |
| **IBS** | `(ncm, uf_destino, municipio_destino_ibge)` | 3 | NCM 8471.30.19 + SP + 3550308 (São Paulo) → 5.2% + 2.8% |
| **IS** | `(ncm)` | 1 | NCM 2402.20.00 (cigarros) → 100% |
| **ICMS** | `(uf_origem, uf_destino)` | 2 | SP → MG → 12% interestadual |
| **IPI** | `(ncm, ex_ipi, crt_emitente, tipo_op, perfil, uf, zona)` | 7 | — |
| **ISS** | `(codigo_ibge, item_lista_servico)` | 2 | 3550308 + 1.05 (telecom) → 2.5% |

> `★ Insight ─────────────────────────────────────`
> Quando duas entidades têm chaves naturais diferentes, forçá-las na mesma tabela viola o princípio de **coesão funcional** — a dependência funcional entre atributos não-chave e a chave primária fica quebrada. No caso de `iva_dual_rules`, `aliquota_cbs` depende apenas de `ncm`, enquanto `aliquota_ibs_estadual` depende de `(ncm, uf_destino, municipio)`. Isso significa que uma tabela unificada **sempre** terá duplicação de dados de CBS (repetidos para cada UF/município) ou NULLs artificiais.
> `─────────────────────────────────────────────────`

**Consequência prática da tabela unificada:**

```sql
-- 1 NCM × 27 UFs × 1 município padrão = 27 linhas
-- A aliquota_cbs é a MESMA em todas as 27 linhas
-- Se CBS mudar de 8.8% para 9.2% → UPDATE em 27 linhas (para apenas 1 NCM)

-- Com tabela independente:
-- 1 NCM = 1 linha em cbs_rates
-- Se CBS mudar → UPDATE em 1 linha
```

Com 5.000 NCMs × 27 UFs = **135.000 linhas** em `iva_dual_rules`, das quais a CBS está duplicada 27× para cada NCM. Uma mudança de alíquota CBS afeta 5.000 linhas em vez de 1.

---

### 2. Ciclo de Vida Independente

Cada imposto tem seu próprio **calendário legislativo**:

| Imposto | Início | Extinção | Duração | Eventos de Mudança |
|---------|--------|----------|---------|-------------------|
| PIS/COFINS | Vigente | **2027** | ~1 ano | Extinção gradual |
| ICMS | Vigente | **2029-2032** | 3-6 anos | Redução gradual por UF |
| ISS | Vigente | **2029-2032** | 3-6 anos | Redução gradual por município |
| CBS | 2026 (shadow) | Indefinido | Permanente | Ajustes setoriais periódicos |
| IBS | 2026 (shadow) | Indefinido | Permanente | Ajustes estaduais/municipais |
| IS | 2026 | Indefinido | Permanente | Inclusão de novas categorias |

**Em uma tabela unificada**, quando o ICMS é extinto em 2032, **nenhuma linha é removida** de `iva_dual_rules` — apenas a aplicação para de consultar `icms_rules`. Mas o IBS continua. Os dois regimes operam em tabelas **diferentes** durante o período híbrido.

O problema da unificação **não é entre regime antigo e novo** (eles já estão em tabelas separadas), mas **entre CBS, IBS e IS dentro do novo regime** — que têm granularidades e ciclos de mudança distintos.

---

### 3. Multi-Tenancy e Segmentação

Com `empresa_id` segmentando todas as tabelas de regras fiscais:

| Cenário | Tabela Unificada | Tabelas Independentes |
|---------|-----------------|----------------------|
| Empresa A tem CBS setorial TELECOM (8.8%) | 1 linha por NCM×UF com `aliquota_cbs = 8.8` | 1 linha em `cbs_rates` com `aliquota_cbs = 8.8` |
| Empresa B tem CBS setorial GERAL (12%) | 1 linha por NCM×UF com `aliquota_cbs = 12.0` | 1 linha em `cbs_rates` com `aliquota_cbs = 12.0` |
| Ambas na mesma UF para IBS | IBS igual para ambas (não segmenta por setor) | IBS independente da CBS |

Na tabela unificada, a segmentação por `empresa_id` **força a duplicação de dados que não variam por empresa**. Se a CBS é uniforme para todas as empresas do setor TELECOM, mas cada empresa precisa de sua própria linha (porque `empresa_id` é parte da chave), a duplicação multiplica.

---

### 4. Complexidade de Desenvolvimento

| Aspecto | Tabela Unificada (`iva_dual_rules`) | Tabelas Independentes (`cbs_rates`, `ibs_rates`, `is_rates`) |
|---------|--------------------------------------|--------------------------------------------------------------|
| **Engine (DT-3)** | 1 query — `computeIvaDual()` | 3 queries — precisa compor resultado de 3 tabelas |
| **Admin CRUD (DT-1)** | 1 endpoint `/aliquotas/iva-dual` que gerencia 3 impostos juntos | 3 endpoints `/aliquotas/cbs`, `/aliquotas/ibs`, `/aliquotas/is` |
| **Carga em Lote** | 1 arquivo CSV com colunas para 3 impostos | 3 arquivos CSV independentes (1 por imposto) |
| **Validação (RN-01 a RN-05)** | Validação cruzada complexa (ex: CBS requer UF? Não, mas a tabela exige) | Validação focada — cada imposto valida só suas regras |
| **Cache Redis** | 1 chave `tax:iva:<ncm>:<uf>:<municipio>` | 3 chaves `tax:cbs:<ncm>`, `tax:ibs:<ncm>:<uf>:<munic>`, `tax:is:<ncm>` |
| **Migração de Dados** | Complexa — afetar CBS afeta IBS na mesma tabela | Isolada — cada imposto migra independentemente |

---

### 5. Performance de Consulta (Motor de Cálculo)

**Cenário: Cálculo de CBS + IBS + IS para um item (NCM = 8471.30.19, UF = SP, Município = 3550308)**

| Abordagem | Queries | Rows Scanned | Cache Keys |
|-----------|---------|-------------|------------|
| **Unificada** | 1 query | 1-3 rows (índice lookup) | 1 key |
| **Independente** | 3 queries (paralelizáveis) | 1 row cada | 3 keys |
| **Independente + VIEW** | 1 query na VIEW (JOIN de 3 tabelas) | 3 rows (3 index lookups) | 1 key (materialized) |

**Conclusão de performance:** A diferença é insignificante com os índices corretos. A tabela unificada vence por 1 query a menos, mas a diferença é de microssegundos — irrelevante comparado ao custo do cálculo tributário em si. Com Redis cache (TTL 24h), a diferença desaparece completamente após o primeiro acesso.

---

## Estratégias Comparadas

### Opção A: Tabela Totalmente Unificada (`iva_dual_rules` — Status Quo)

```
┌─────────────────────────────────────────────────────────────┐
│  iva_dual_rules                                             │
│  ┌───────────────────────────────────────────────────────┐  │
│  │ ncm | uf_destino | municipio | cbs% | ibs_e% | ibs_m% │  │
│  │     |            |           |      | is_flag | is%   │  │
│  ├───────────────────────────────────────────────────────┤  │
│  │ Todos os impostos do IVA Dual em uma única tabela     │  │
│  │ Chave: (empresa_id, ncm, uf_destino, municipio)       │  │
│  └───────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

| Prós | Contras |
|------|---------|
| ✅ 1 query no motor de cálculo | ❌ Duplicação de CBS (27× por NCM) |
| ✅ 1 endpoint de administração | ❌ CBS não precisa de UF, mas a PK exige |
| ✅ Simples de entender (1 tabela = IVA Dual) | ❌ Mudança de CBS afeta 135K+ linhas |
| ✅ Backward-compatible com design atual | ❌ IS com NULLs artificiais quando `is_flag = false` |
| | ❌ Carga em lote força arquivo multi-imposto |
| | ❌ Cada imposto não pode evoluir independentemente |
| | ❌ Validações cruzadas complexas (RNs conflitantes) |

### Opção B: Tabelas Totalmente Independentes

```
┌──────────────────────┐  ┌──────────────────────────────────┐  ┌──────────────────────┐
│  cbs_rates           │  │  ibs_rates                       │  │  is_rates            │
│  ┌────────────────┐  │  │  ┌───────────────────────────┐  │  │  ┌────────────────┐  │
│  │ ncm | classe   │  │  │  │ ncm | uf | municipio      │  │  │  │ ncm | categoria│  │
│  │     | cbs%     │  │  │  │     |    | ibs_e% | ibs_m% │  │  │  │     | is%      │  │
│  └────────────────┘  │  │  └───────────────────────────┘  │  │  └────────────────┘  │
│  PK: (empresa, ncm)  │  │  PK: (empresa, ncm, uf, munic)  │  │  PK: (empresa, ncm)  │
└──────────────────────┘  └──────────────────────────────────┘  └──────────────────────┘
```

| Prós | Contras |
|------|---------|
| ✅ Cada imposto tem sua chave natural correta | ❌ Engine precisa de 3 queries (ou 1 VIEW) |
| ✅ Sem duplicação — CBS: 1 linha por NCM | ❌ 3 tabelas para administrar em vez de 1 |
| ✅ Mudança de CBS = UPDATE em 1 linha (não 27) | ❌ Maior complexidade no CRUD (3 endpoints) |
| ✅ Cada imposto evolui independentemente | ❌ Carga em lote: 3 arquivos separados |
| ✅ IS sem NULLs artificiais | |
| ✅ Validações isoladas e simples | |
| ✅ Migrações independentes (Flyway por imposto) | |
| ✅ Alinhado com o padrão do regime atual (ICMS, IPI, ISS são independentes) | |

### Opção C: Híbrida — Tabelas Independentes + VIEW Unificada (Recomendada)

```
┌──────────────────────────────────────────────────────────────────────────────────┐
│                          CAMADA DE APRESENTAÇÃO (VIEWs)                           │
│  ┌────────────────────────────────────────────────────────────────────────────┐  │
│  │  v_iva_dual_active  (materialized view para o motor de cálculo)             │  │
│  │  SELECT c.ncm, i.uf_destino, i.municipio,                                   │  │
│  │         c.aliquota_cbs, i.aliquota_ibs_estadual, i.aliquota_ibs_municipal,  │  │
│  │         s.aliquota_is                                                        │  │
│  │  FROM cbs_rates c                                                           │  │
│  │  LEFT JOIN ibs_rates i ON c.ncm = i.ncm AND c.empresa_id = i.empresa_id     │  │
│  │  LEFT JOIN is_rates s ON c.ncm = s.ncm AND c.empresa_id = s.empresa_id      │  │
│  └────────────────────────────────────────────────────────────────────────────┘  │
│                                                                                  │
│                          CAMADA DE DADOS (TABELAS)                                │
│  ┌──────────────┐    ┌─────────────────────────────┐    ┌──────────────────────┐  │
│  │ cbs_rates    │    │ ibs_rates                   │    │ is_rates             │  │
│  │              │    │                             │    │                      │  │
│  │ ncm (PK)     │    │ ncm (PK)                    │    │ ncm (PK)             │  │
│  │ classe_trib  │    │ uf_destino (PK)             │    │ categoria            │  │
│  │ aliquota_cbs │    │ municipio_destino_ibge (PK) │    │ aliquota_is          │  │
│  │ empresa_id   │    │ aliquota_ibs_estadual       │    │ empresa_id           │  │
│  │ vigencia     │    │ aliquota_ibs_municipal      │    │ vigencia             │  │
│  │              │    │ empresa_id                  │    │                      │  │
│  │              │    │ vigencia                    │    │                      │  │
│  └──────────────┘    └─────────────────────────────┘    └──────────────────────┘  │
└──────────────────────────────────────────────────────────────────────────────────┘
```

| Prós | Contras |
|------|---------|
| ✅ Todos os prós das tabelas independentes | ❌ VIEW materializada precisa de refresh |
| ✅ Engine (DT-3) continua com 1 query (na VIEW) | ❌ Maior número de objetos no banco |
| ✅ Admin (DT-1) pode usar tabelas individuais ou VIEW | ❌ Migração do design atual (`iva_dual_rules` → 3 tabelas) |
| ✅ Contrato de interface via VIEW — desacopla storage de consumo | |
| ✅ Transição suave: a VIEW garante backward-compatibility | |
| ✅ Performance: materialized view elimina custo do JOIN para o motor | |

---

## Recomendação

### Adotar **Opção C — Tabelas Independentes + VIEW Unificada**

**Justificativa técnica (6 fundamentos):**

1. **Chaves naturais diferentes** — CBS depende de `(ncm, classe_tributaria)`, IBS de `(ncm, uf, municipio)`, IS de `(ncm)`. Forçá-las na mesma tabela viola a 3FN e causa duplicação de dados.

2. **Ciclos de vida distintos** — Cada imposto tem seu próprio calendário legislativo. CBS pode mudar (ajuste setorial) sem afetar IBS. Na tabela unificada, uma mudança de CBS força rewrite de todas as linhas de IBS junto.

3. **Consistência com o regime atual** — `icms_rules`, `ipi_regras` e `iss_rates` já são independentes. Unificar apenas o novo regime cria inconsistência no schema: por que ICMS merece tabela própria mas IBS (seu substituto) não?

4. **Multi-tenancy sem duplicação** — Na tabela unificada, `empresa_id` na PK força duplicação de CBS (que não varia por UF) para cada UF de IBS. Com tabelas independentes, CBS tem 1 linha por empresa+NCM, IBS tem 1 linha por empresa+NCM+UF+município.

5. **Evolução independente** — Se o IS ganhar uma nova categoria (`JOGOS_DE_AZAR`) em 2028, apenas `is_rates` é alterado. Na tabela unificada, adicionar coluna `nova_categoria_is` afeta todas as linhas de CBS e IBS.

6. **A VIEW garante backward-compatibility** — O motor de cálculo (`computeIvaDual()`) não precisa ser refatorado imediatamente. A VIEW `v_iva_dual_active` entrega o mesmo formato que `iva_dual_rules` entrega hoje.

### Plano de Migração: `iva_dual_rules` → 3 Tabelas Independentes

```
Fase 0 (agora):      Criar cbs_rates, ibs_rates, is_rates
                     Criar VIEW v_iva_dual_active (mantém contrato atual)
                     
Fase 1 (transição):  DT-1 escreve nas 3 tabelas independentes
                     Trigger replica na VIEW (ou a VIEW consulta as 3)
                     DT-3 lê da VIEW (sem mudanças no código)
                     
Fase 2 (estabilização): iva_dual_rules marcada como deprecated
                     Dados migrados para as 3 tabelas independentes
                     
Fase 3 (limpeza):    DROP TABLE iva_dual_rules (após validar que VIEW supre todas as queries)
```

### Sequência de Migrations (revisada com nomenclatura português)

```
V24__criar_aliquotas_cbs.sql              — Tabela CBS independente
V25__criar_aliquotas_ibs.sql              — Tabela IBS independente
V26__criar_aliquotas_is.sql               — Tabela IS independente
V27__criar_view_aliquotas_iva.sql         — VIEW materializada v_aliquotas_iva
V28__popular_tabelas_independentes.sql     — Migração: iva_dual_rules → aliquotas_cbs/ibs/is
V29__deprecate_iva_dual_rules.sql         — Marcar iva_dual_rules como legado
V30__criar_views_compatibilidade.sql      — 🔴 VIEWs com nomes legados para DT-3
```

### Convenção de Nomenclatura (ADR-003 Apêndice)

**Decisão:** Tabelas novas do DT-1 seguem nomenclatura em **português** (domínio fiscal brasileiro). Tabelas legadas do DT-3 mantêm nomes originais (inglês/misto) por backward-compatibility.

| Nome Original (DT-3 legado) | Novo Nome (DT-1) | Estratégia |
|---|---|---|
| `cbs_rates` | `aliquotas_cbs` | Tabela em PT + VIEW `cbs_rates` → `aliquotas_cbs` |
| `ibs_rates` | `aliquotas_ibs` | Tabela em PT + VIEW `ibs_rates` → `aliquotas_ibs` |
| `is_rates` | `aliquotas_is` | Tabela em PT + VIEW `is_rates` → `aliquotas_is` |
| `iva_dual_rules` | `v_aliquotas_iva` | VIEW materializada + VIEW `iva_dual_rules` → `v_aliquotas_iva` |
| `tenants` | `estabelecimentos` | Tabela em PT + VIEW `tenants` → `estabelecimentos` |
| `auditoria_log` | `auditoria` | Tabela em PT + VIEW `auditoria_log` → `auditoria` |

> ⚠️ **VIEWs de compatibilidade (V30):** 6 VIEWs updatable são criadas com os nomes originais em inglês, redirecionando para as tabelas em português. Isso garante que o DT-3 (Go) continue funcionando sem nenhuma alteração de código. Ver [erd.md#-views-de-compatibilidade-legada](../erd.md) para DDL completo e requisitos de deploy.

### Impacto nos Componentes

| Componente | Mudança Necessária | Complexidade |
|-----------|-------------------|-------------|
| **DT-1 Admin Service** | 1 endpoint → 3 endpoints específicos por imposto + nomes PT | Média |
| **DT-1 Carga em Lote** | 1 template CSV → 3 templates (ou CSV com seções) | Média |
| **DT-3 Engine** | `computeIvaDual()` → VIEWs transparentes (query idêntica) | **Nenhuma** |
| **DT-3 Cache Redis** | 1 chave → 3 chaves (mais granular, melhor invalidação) | Baixa |
| **DT-2 Portal** | 1 tela "IVA Dual" → 3 telas (CBS, IBS, IS) | Média |
| **Testes** | Testes por imposto (mais isolados, mais simples) | Baixa |
| **Deploy** | Ordem: tabelas → dados → VIEWs (mesma transação) | Crítico |

---

## Riscos e Mitigações

| Risco | Probabilidade | Impacto | Mitigação |
|-------|-------------|---------|-----------|
| Motor de cálculo quebra por mudança na VIEW | Baixa | Alto | VIEW materializada mantém formato idêntico; testes de integração DT-1 ↔ DT-3 |
| Perda de dados na migração `iva_dual_rules` → 3 tabelas | Baixa | Crítico | Migração em 2 fases: populate shadow tables → validate row counts → switch |
| Performance da VIEW com JOIN de 3 tabelas | Média | Médio | Materialized view com refresh a cada carga de lote; índices nas 3 tabelas |
| Resistência da equipe (já acostumada com `iva_dual_rules`) | Média | Baixo | A VIEW garante que ninguém precise mudar código no curto prazo |

---

## Alternativas Consideradas e Rejeitadas

### Alternativa 1: Manter `iva_dual_rules` como está

**Rejeitada porque:** A duplicação de CBS (27× para cada NCM), a impossibilidade de evolução independente dos impostos, e a inconsistência com o padrão já adotado para ICMS/IPI/ISS tornam esta abordagem insustentável a longo prazo. O custo de corrigir agora (Fase 0 de documentação) é ordens de grandeza menor que corrigir em 2028 com o sistema em produção.

### Alternativa 2: `iva_dual_rules` + tabelas satélite por imposto

**Rejeitada porque:** Introduz complexidade desnecessária. Se já existem tabelas satélite, `iva_dual_rules` vira uma tabela de JOIN redundante. Melhor eliminar completamente e usar apenas as independentes.

### Alternativa 3: Uma tabela por regime (antigo vs novo) em vez de por imposto

**Rejeitada porque:** O regime antigo já está fragmentado por imposto. Unificar só o novo regime criaria inconsistência. E o período híbrido (2027-2032) exigiria queries que cruzam os dois regimes — ter tabelas separadas por imposto facilita desativar um imposto sem afetar os outros.

---

## Decisão

**Adotar tabelas independentes por imposto (`cbs_rates`, `ibs_rates`, `is_rates`) com uma VIEW materializada unificada (`v_iva_dual_active`) como contrato de interface para o motor de cálculo.**

Esta decisão:
- Alinha-se com o padrão já estabelecido pelas tabelas do regime atual
- Respeita as diferentes chaves naturais e ciclos de vida de cada imposto
- Mantém backward-compatibility via VIEW para o motor de cálculo
- Permite que cada imposto evolua independentemente durante o período de transição 2026-2032
- Elimina duplicação de dados (CBS repetida 27× por NCM)
- Reduz o blast radius de mudanças legislativas (alterar CBS não afeta IBS)

---

## Controle de Versão

| Versão | Data | Autor | Mudanças |
|--------|------|-------|----------|
| 1.0 | 2026-07-12 | Data Architect + PostgreSQL Specialist + Engineering Architect | Proposta inicial |

---

🤖 *Documento gerado por Inteligência Artificial em 12 de Julho de 2026.*

**Agentes de IA utilizados nesta análise:**
- Arquiteto de Dados (skills: `sql-pro`, `database-optimizer`)
- Especialista PostgreSQL (skill: `postgresql`)
- Senior Architect (engineering skill: `senior-architect`)
- Senior Backend Engineer (engineering skill: `senior-backend`)

**Skills aplicadas:** `postgresql`, `sql-pro`, `database-optimizer`, `senior-architect`, `senior-backend`

**Ferramenta de orquestração:** Claude Code (Anthropic)
