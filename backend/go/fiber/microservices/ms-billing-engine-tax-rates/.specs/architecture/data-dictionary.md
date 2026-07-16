# Dicionário de Dados — ms-billing-engine-tax-rates

> **Schema:** `billing_tax_rates`
> **Fonte:** `data/init.sql`
> **Atualizado:** 2026-07-02
> **Confiança:** 99% (VERDE) — ver [governance/confidence-report.md](../governance/confidence-report.md)

Documento canônico que descreve a **função de negócio**, o **propósito**, os **padrões de uso** e as **regras de negócio associadas** a cada tabela do schema `billing_tax_rates`. Complementa o [erd.md](erd.md) (que documenta a estrutura relacional e os tipos de colunas) e o [domain.md](../domain/domain.md) (que documenta as regras de negócio por tributo).

---

## Índice

1. [Visão Geral do Schema](#visão-geral-do-schema)
2. [Tabelas do Regime Atual (Pré-Reforma)](#tabelas-do-regime-atual-pré-reforma)
   - [icms_rules](#1-icms_rules)
   - [federal_tax_rules](#2-federal_tax_rules)
   - [product_tax_exceptions](#3-product_tax_exceptions)
   - [tax_equivalence](#4-tax_equivalence)
   - [simples_nacional_rates](#5-simples_nacional_rates)
   - [ipi_regras](#6-ipi_regras)
   - [iss_rates](#7-iss_rates)
3. [Tabelas da Reforma Tributária (IVA Dual)](#tabelas-da-reforma-tributária-iva-dual)
   - [iva_dual_rules](#8-iva_dual_rules)
   - [iva_dual_rules_log](#9-iva_dual_rules_log)
   - [reforma_tributaria_rules](#10-reforma_tributaria_rules-legado)
   - [ncm_seletivo](#11-ncm_seletivo)
   - [cbs_rates](#12-cbs_rates)
   - [cst_reforma](#13-cst_reforma)
4. [Tabelas Operacionais](#tabelas-operacionais)
   - [tax_tokens](#14-tax_tokens)
   - [fornecedor_fiscal](#15-fornecedor_fiscal)
5. [Matriz Tabela × Calculadora](#matriz-tabela--calculadora)
6. [Matriz Tabela × Regra de Negócio](#matriz-tabela--regra-de-negócio)
7. [Padrões Transversais](#padrões-transversais)

---

## Visão Geral do Schema

O schema `billing_tax_rates` contém **15 tabelas** organizadas em três categorias funcionais:

| Categoria | Tabelas | Finalidade |
|---|---|---|
| **Regime Atual (Pré-Reforma)** | `icms_rules`, `federal_tax_rules`, `product_tax_exceptions`, `tax_equivalence`, `simples_nacional_rates`, `ipi_regras`, `iss_rates` | Regras de cálculo dos tributos vigentes (ICMS, PIS/COFINS, IPI, ISS) |
| **Reforma Tributária (IVA Dual)** | `iva_dual_rules`, `iva_dual_rules_log`, `reforma_tributaria_rules`, `ncm_seletivo`, `cbs_rates`, `cst_reforma` | Regras do novo sistema dual CBS/IBS (EC 132/2023, LC 214/2025) |
| **Operacional** | `tax_tokens`, `fornecedor_fiscal` | Suporte a funcionalidades cross-cutting (congelamento de alíquotas, qualificação de fornecedores) |

### Diagrama de Dependência Funcional

```
                    ┌──────────────────────────────────────────────┐
                    │           Motor de Cálculo (SOP-013)         │
                    │         internal/calculator/engine.go        │
                    └──────────────────────────────────────────────┘
                                          │
          ┌──────────────────┬────────────┼────────────┬──────────────────┐
          ▼                  ▼            ▼            ▼                  ▼
┌─────────────────┐ ┌──────────────┐ ┌─────────┐ ┌──────────┐ ┌──────────────┐
│ ICMS Calculator │ │PIS/COFINS    │ │  IPI    │ │   ISS    │ │ Reforma      │
│ (icms.go)       │ │Calculator    │ │(ipi.go) │ │ (iss.go) │ │ CBS/IBS/IS   │
└────────┬────────┘ └──────┬───────┘ └────┬────┘ └────┬─────┘ └──────┬───────┘
         │                 │              │           │               │
         ▼                 ▼              ▼           ▼               ▼
┌─────────────────┐ ┌──────────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────┐
│  icms_rules     │ │federal_tax_  │ │ipi_regras│ │iss_rates │ │iva_dual_rules│
│  product_tax_   │ │rules         │ │          │ │          │ │ncm_seletivo  │
│  exceptions     │ │              │ │          │ │          │ │cst_reforma   │
│  tax_equivalence│ │              │ │          │ │          │ │cbs_rates     │
│  simples_       │ │              │ │          │ │          │ │              │
│  nacional_rates │ │              │ │          │ │          │ │              │
└─────────────────┘ └──────────────┘ └──────────┘ └──────────┘ └──────────────┘

┌──────────────────────┐    ┌──────────────────────┐
│   Token Service      │    │   Credit Engine      │
│ (token/service.go)   │    │ (credit/engine.go)   │
└──────────┬───────────┘    └──────────┬───────────┘
           ▼                           ▼
┌──────────────────────┐    ┌──────────────────────┐
│  tax_tokens          │    │  fornecedor_fiscal   │
│  iva_dual_rules      │    │  iva_dual_rules      │
└──────────────────────┘    └──────────────────────┘
```

---

## Tabelas do Regime Atual (Pré-Reforma)

### 1. `icms_rules`

| Atributo | Valor |
|---|---|
| **Propósito** | Matriz de alíquotas interestaduais de ICMS — define as regras gerais por par (UF origem, UF destino) |
| **Tributo** | ICMS (próprio, DIFAL, ST), FCP |
| **Quem consulta** | `ICMSCalculator` (`internal/legacy/icms.go:442`) via `GetICMSRule(ctx, orig, dest)` |
| **Endpoint afetado** | `POST /v1/calculate` |
| **Regras de negócio** | BR-TAX-CALC-004 (ICMS Próprio), BR-TAX-CALC-005 (ICMS-ST), BR-TAX-CALC-006 (DIFAL), BR-TAX-CALC-021/022 (Desoneração) |

**Função no sistema:** Esta é a tabela de **regras gerais** de ICMS. Para cada par de UFs (origem → destino), define as alíquotas base que serão usadas no cálculo. Funciona como fallback: se existir uma exceção mais específica em `product_tax_exceptions` para o mesmo par (NCM, UF), a exceção prevalece (merge via `getEffectiveTaxConfig()`).

**Padrão de lookup:**
```
SELECT * FROM icms_rules
WHERE uf_origem = $1 AND uf_destino = $2
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
LIMIT 1
```
- Lookup por par (`uf_origem`, `uf_destino`)
- Busca a regra vigente na data atual (controle de validade temporal)
- Merge com `product_tax_exceptions` via `getEffectiveTaxConfig()` no código

**Colunas-chave:**
- `aliquota_interna` — alíquota praticada dentro da UF destino (ex: 18% SP)
- `aliquota_interestadual` — alíquota interestadual (7% Sul/Sudeste, 12% demais)
- `reducao_base` — percentual de redução da base de cálculo
- `percentual_fcp` — adicional do Fundo de Combate à Pobreza (até 2%)
- `mva_padrao` — Margem de Valor Agregado para ST
- `possui_protocolo_st` — flag que habilita ICMS-ST
- `motivo_desoneracao_icms` — código SEFAZ de desoneração (1-12, 90)
- `possui_desoneracao` — flag que ativa o cálculo de ICMS desonerado
- `inicio_validade` / `final_validade` — controle de vigência temporal

**Triggers:** `billing_tax_rates_icms_rules_fim_validade` (fecha vigência da regra anterior), `billing_tax_rates_icms_rules_atualizado_em` (atualiza timestamp)

**Mapeamento Go:** `repository.ICMSRule` struct (`core-lib/repository/entities.go:43`)

---

### 2. `federal_tax_rules`

| Atributo | Valor |
|---|---|
| **Propósito** | Tabela de alíquotas de PIS e COFINS por regime tributário e CST |
| **Tributo** | PIS, COFINS |
| **Quem consulta** | `PISCofinsCalculator` (`internal/legacy/pis_cofins.go:44`) via `GetFederalTaxRule(ctx, regime, cstPis, cstCofins)` |
| **Endpoint afetado** | `POST /v1/calculate` |
| **Regras de negócio** | BR-TAX-CALC-008 (PIS), BR-TAX-CALC-009 (COFINS), "Tese do Século" (exclusão ICMS da base) |

**Função no sistema:** Define as alíquotas nominais de PIS e COFINS aplicáveis a cada combinação de regime tributário (Lucro Real, Lucro Presumido) e CST. A flag `exclui_icms_base` implementa a "Tese do Século" (STF RE 574.706) — quando `true`, o ICMS destacado na nota é excluído da base de cálculo de PIS/COFINS.

**Padrão de lookup:**
```
SELECT * FROM federal_tax_rules
WHERE regime_tributario = $1 AND cst_pis = $2 AND cst_cofins = $3
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
LIMIT 1
```
- Lookup por tripla (`regime`, `cst_pis`, `cst_cofins`)
- Fallback no código para defaults hardcoded (PIS 1.65%, COFINS 7.6%) quando a regra não é encontrada
- Campo `fonte_aliquota` nos detalhes indica se veio do banco ou do fallback

**Colunas-chave:**
- `regime_tributario` — `LUCRO_REAL` ou `LUCRO_PRESUMIDO`
- `cst_pis` / `cst_cofins` — códigos CST (01, 02, 03, 04, 05, 06, 49, 50-99)
- `aliquota_pis` — alíquota nominal (ex: 1.65%)
- `aliquota_cofins` — alíquota nominal (ex: 7.60%)
- `exclui_icms_base` — flag da Tese do Século

**Triggers:** `billing_tax_rates_federal_tax_rules_fim_validade`, `billing_tax_rates_federal_tax_rules_atualizado_em`

**Mapeamento Go:** `repository.FederalTaxRule` struct (`core-lib/repository/entities.go:29`)

---

### 3. `product_tax_exceptions`

| Atributo | Valor |
|---|---|
| **Propósito** | Exceções fiscais por NCM — sobrepõe as regras gerais com configurações específicas de produto |
| **Tributo** | ICMS, PIS, COFINS, ICMS-ST |
| **Quem consulta** | `ICMSCalculator` via `getEffectiveTaxConfig()` que chama `GetProductException(ctx, ncmFull, ncmGroup, ufDestino, regime)` |
| **Endpoint afetado** | `POST /v1/calculate` |
| **Regras de negócio** | BR-TAX-CALC-004, BR-TAX-CALC-005 (ST por NCM), BR-TAX-CONS-013 (CST desoneração) |

**Função no sistema:** Esta tabela implementa o **princípio da exceção**: regras gerais (`icms_rules`, `federal_tax_rules`) podem ser sobrescritas por regras atreladas a um NCM específico. Exemplos: produtos monofásicos (PIS/COFINS CST 04), produtos com ST (MVA específico por NCM), alíquota zero, desonerações setoriais, importados (alíquota interestadual 4%).

**Padrão de lookup (match progressivo):**
```
SELECT * FROM product_tax_exceptions
WHERE (ncm = $1 OR ncm = $2)  -- ncmFull OU ncmGroup (4 dígitos)
  AND (uf_destino = $3 OR uf_destino = '*')
  AND (uf_origem = $4 OR uf_origem = '**')
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
ORDER BY (ncm = $1) DESC, (uf_destino = $3) DESC  -- prioriza match mais específico
LIMIT 1
```
- **Wildcard matching:** `*` em UF destino significa "todas as UFs"; `**` em UF origem significa "todas as origens"
- NCM com 4 dígitos funciona como grupo (ex: `4011` captura todos os subitens de pneus)
- Match mais específico (NCM completo + UF específica) prevalece

**Colunas-chave (além das herdadas de `icms_rules`):**
- `ncm` — NCM completo (8 dígitos) ou grupo (4 dígitos)
- `csosn` — CSOSN para produtos no Simples Nacional
- `mva_st` — MVA específico do produto para ST
- `aliquota_interna_destino` — override da alíquota interna para este NCM
- `aliquota_interestadual` — override (ex: 4% para importados)
- `aliquota_pis_unitario` / `aliquota_cofins_unitario` — para CST 03 (tributação por unidade)
- `regime_tributario_destino` — afeta cálculo de MVA em ST
- `desoneracao_codigo_beneficio_fiscal` — código do benefício fiscal (cBenef no XML)

**Mapeamento Go:** `repository.ProductException` struct (`core-lib/repository/entities.go:87`)

---

### 4. `tax_equivalence`

| Atributo | Valor |
|---|---|
| **Propósito** | Tabela de equivalência CSOSN → CST para empresas do Simples Nacional |
| **Tributo** | ICMS (Simples Nacional) |
| **Quem consulta** | `ICMSCalculator` via `GetEquivalence(ctx, csosn, tipoOperacao)` |
| **Endpoint afetado** | `POST /v1/calculate` |
| **Regras de negócio** | BR-TAX-CALC-007 (ICMS Simples Nacional) |

**Função no sistema:** Empresas do Simples Nacional usam CSOSN (3 dígitos) em vez de CST (2 dígitos). Esta tabela mapeia a equivalência entre os códigos, permitindo que o motor use a mesma lógica de cálculo independentemente do regime. Também indica se a operação permite crédito de ICMS para o comprador.

**Padrão de lookup:**
```
SELECT * FROM tax_equivalence
WHERE csosn = $1 AND tipo_operacao_fiscal = $2
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
LIMIT 1
```

**Colunas-chave:**
- `csosn` — código CSOSN de 3 dígitos
- `cst_equivalente` — CST equivalente (2 dígitos)
- `permite_credito` — flag de permissão de crédito ao comprador
- `tipo_operacao_fiscal` — contexto (entrada/saída) — a equivalência pode mudar

**Mapeamento Go:** `repository.TaxEquivalence` struct (`core-lib/repository/entities.go:60`)

---

### 5. `simples_nacional_rates`

| Atributo | Valor |
|---|---|
| **Propósito** | Tabela progressiva de faixas do Simples Nacional por anexo |
| **Tributo** | ICMS (dentro do Simples Nacional) |
| **Quem consulta** | `ICMSCalculator` via `GetSimplesFaixa(ctx, anexo, rbt12)` |
| **Endpoint afetado** | `POST /v1/calculate` |
| **Regras de negócio** | BR-TAX-CALC-007 (ICMS Simples Nacional) |

**Função no sistema:** O Simples Nacional usa tabela progressiva: a alíquota efetiva depende da receita bruta acumulada (RBT12) e do anexo (I=Comércio, II=Indústria, III a V=Serviços). Esta tabela armazena as faixas e calcula a alíquota efetiva via fórmula `(RBT12 × Alíquota Nominal − Parcela a Deduzir) / RBT12`.

**Padrão de lookup:**
```
SELECT * FROM simples_nacional_rates
WHERE anexo = $1 AND receita_min <= $2 AND receita_max >= $2
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
LIMIT 1
```
- Busca por faixa de RBT12 (between `receita_min` e `receita_max`)
- Índice `idx_rbt12_range` otimiza esta consulta

**Colunas-chave:**
- `anexo` — ANEXO_I a ANEXO_V (classificação da atividade)
- `faixa` — 1 a 6
- `aliquota_nominal` — alíquota da faixa (ex: 4.00%)
- `valor_deduzir` — parcela a deduzir
- `percentual_icms` — percentual do ICMS dentro do Simples (ex: 33.5%)
- `aliquota_repasse_credito` — percentual de crédito que o comprador pode apropriar

**Mapeamento Go:** `repository.SimplesFaixa` struct (`core-lib/repository/entities.go:72`)

---

### 6. `ipi_regras`

| Atributo | Valor |
|---|---|
| **Propósito** | Tabela de regras do IPI com 7 dimensões de lookup (NCM, EX, CRT, tipo operação, perfil comprador, UF, zona especial) |
| **Tributo** | IPI (Imposto sobre Produtos Industrializados) |
| **Quem consulta** | `IPICalculator` (`internal/legacy/ipi.go`) via `GetIPIRegra(ctx, ncm, ex, crt, tipoOp, perfil, uf, zona, data)` |
| **Endpoint afetado** | `POST /v1/calculate` (Fase 1 sequencial — compõe base do ICMS) |
| **Regras de negócio** | BR-TAX-CALC-003 (IPI) |

**Função no sistema:** O IPI é o tributo com a matriz de lookup mais complexa — 7 dimensões independentes que determinam a alíquota aplicável. A tabela segue o princípio de **especificidade decrescente**: campos com `*` (wildcard) são menos específicos e perdem precedência no `ORDER BY`.

**Padrão de lookup (match progressivo com 7 dimensões):**
```
SELECT * FROM ipi_regras
WHERE (ncm = $1 OR ncm = '*')
  AND (ex_ipi = $2 OR ex_ipi = '*')
  AND (crt_emitente = $3 OR crt_emitente = '*')
  AND (tipo_operacao_fiscal = $4 OR tipo_operacao_fiscal = '*')
  AND (perfil_comprador = $5 OR perfil_comprador = '*')
  AND (uf_destino = $6 OR uf_destino = '*')
  AND (zona_especial = $7 OR zona_especial = '*')
  AND CURRENT_DATE BETWEEN inicio_validade AND COALESCE(final_validade, '9999-12-31')
ORDER BY (ncm != '*') DESC, (ex_ipi != '*') DESC, (crt_emitente != '*') DESC,
         (tipo_operacao_fiscal != '*') DESC, (perfil_comprador != '*') DESC,
         (uf_destino != '*') DESC, (zona_especial != '*') DESC
LIMIT 1
```

**Colunas-chave:**
- `ncm` / `ex_ipi` — classificação fiscal na TIPI
- `crt_emitente` — regime tributário do emitente
- `tipo_operacao_fiscal` — natureza da operação
- `perfil_comprador` — classificação do adquirente
- `zona_especial` — Zona Franca de Manaus, ALC, etc.
- `aliquota_ipi` — alíquota ad valorem (%)
- `valor_pauta_ipi` — valor fixo por unidade (ad pauta)
- `cst_ipi` — CST do IPI (2 dígitos)
- `c_enq` — código de enquadramento legal (3 dígitos)

**Mapeamento Go:** `repository.IPIRegra` struct (`core-lib/repository/entities.go:111`)

---

### 7. `iss_rates`

| Atributo | Valor |
|---|---|
| **Propósito** | Alíquotas de ISS por município (código IBGE) e item da Lista de Serviços (LC 116/2003) |
| **Tributo** | ISS (Imposto Sobre Serviços) |
| **Quem consulta** | `ISSCalculator` (`internal/legacy/iss.go`) — usa o código IBGE do município de destino |
| **Endpoint afetado** | `POST /v1/calculate` (Fase 4 paralela) |
| **Regras de negócio** | BR-TAX-CALC-016 (ISS), BR-TAX-CONS-007 (validação 2%-5%) |

**Função no sistema:** O ISS é tributo municipal — a alíquota varia por município. Esta tabela armazena as alíquotas vigentes para cada município brasileiro (código IBGE de 7 dígitos). A validação `2% ≤ alíquota ≤ 5%` (LC 116/2003 art. 8º-A) é aplicada em runtime; alíquotas fora do range geram `slog.Warn` mas não bloqueiam o cálculo.

**Padrão de lookup:**
```
SELECT * FROM iss_rates
WHERE codigo_ibge = $1
  AND (item_lista_servico = $2 OR item_lista_servico IS NULL)
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
ORDER BY item_lista_servico DESC NULLS LAST
LIMIT 1
```

**Colunas-chave:**
- `codigo_ibge` — código IBGE do município (7 dígitos)
- `item_lista_servico` — item da Lista anexa à LC 116/2003 (ex: `1.05` para telecom)
- `aliquota_iss` — alíquota do ISS no município (range: 2%-5%)

**COMMENT ON TABLE:** "Aliquotas ISS por municipio. Range constitucional: [2%, 5%] conforme LC 116/2003 art. 8o-A."

**Mapeamento Go:** `repository.ISSRate` — lido diretamente via query SQL no repository

---

## Tabelas da Reforma Tributária (IVA Dual)

### 8. `iva_dual_rules`

| Atributo | Valor |
|---|---|
| **Propósito** | Tabela mestra do IVA Dual — define as alíquotas de CBS (federal) e IBS (estadual + municipal) para cada tupla (NCM, UF destino, município IBGE) |
| **Tributo** | CBS, IBS, IS |
| **Quem consulta** | `CBSCalculator`, `IBSCalculator` (via `computeIvaDual()` em `reforma.go:62`), `ISFilter` (via `ncm_seletivo`), `CreditEngine` (`credit/engine.go:155`), `TokenService` (`token/service.go:99`), `AdminTaxService` (`admin/service.go:76`), `FallbackIBSClient` (`ibsclient/client.go:310`) |
| **Endpoint afetado** | `POST /v1/calculate`, `POST /v1/simulate`, `POST /v1/token/generate`, `POST /v1/credit/calculate`, `GET /v1/admin/tax-rates/iva-dual` |
| **Regras de negócio** | BR-TAX-CALC-010 (CBS), BR-TAX-CALC-011 (IBS), BR-TAX-CONS-010 (IS), SOP-001, SOP-013, DT-001 |

**Função no sistema:** Esta é a **tabela mais importante da Reforma Tributária**. Segue as premissas definidas em `data/init.sql:351-356`:

1. **TODO NCM deve estar configurado** nesta tabela para gerar o cálculo do imposto
2. **NCM e UF_DESTINO são obrigatórios**
3. **`aliquota_cbs`, `aliquota_ibs_estadual`, `aliquota_ibs_municipal` são obrigatórios**
4. Cada UF destino tem uma linha **sem município** (`municipio_destino_ibge = NULL`) que define as alíquotas padrão estaduais
5. Municípios com regra diferente têm linha específica com **todas as alíquotas preenchidas**

**Padrão de lookup (especificidade decrescente):**
```
SELECT * FROM iva_dual_rules
WHERE ncm = $1
  AND uf_destino = $2
  AND (municipio_destino_ibge = $3 OR municipio_destino_ibge IS NULL)
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
ORDER BY municipio_destino_ibge DESC NULLS LAST  -- regra específica do município prevalece
LIMIT 1
```
- O `ORDER BY municipio_destino_ibge DESC` garante que uma linha com município específico venha **antes** da linha com `NULL` (padrão estadual)
- Índice único `idx_iva_rules_lookup` em `(ncm, uf_destino, COALESCE(municipio_destino_ibge, '0000000'))`

**Cache Redis:** Chave `tax:iva:<ncm>:<uf>:<municipio>` com TTL 24h. `CBSCalculator` e `IBSCalculator` compartilham `computeIvaDual()` — a segunda chamada atinge o cache.

**Colunas-chave:**
- `ncm` — NCM do produto (8 dígitos, obrigatório)
- `uf_destino` — UF de destino do imposto (100% no destino)
- `municipio_destino_ibge` — código IBGE (7 dígitos, NULL = regra padrão estadual)
- `aliquota_cbs` — alíquota CBS federal unificada
- `aliquota_ibs_estadual` — alíquota IBS estadual
- `aliquota_ibs_municipal` — alíquota IBS municipal
- `percentual_reducao` — redução de transição (0, 60, 100)
- `is_imposto_seletivo` — flag de incidência do IS
- `aliquota_is` — alíquota do Imposto Seletivo

**Trigger de auditoria:** `trg_audit_iva_dual_rules` — registra INSERT, UPDATE, DELETE na tabela `iva_dual_rules_log`

**Mapeamento Go:** `repository.IvaDualRule` struct (`core-lib/repository/entities.go:145`)

---

### 9. `iva_dual_rules_log`

| Atributo | Valor |
|---|---|
| **Propósito** | Trilha de auditoria para alterações na tabela `iva_dual_rules` |
| **Tributo** | CBS, IBS (indireto — auditoria) |
| **Quem consulta** | Trigger automático (`trg_audit_iva_dual_rules`); Admin dashboard para consulta de histórico |
| **Endpoint afetado** | `GET /v1/admin/tax-rates/iva-dual` (histórico) |
| **Regras de negócio** | Governança fiscal — rastreabilidade de alterações em alíquotas |

**Função no sistema:** Toda alteração (INSERT, UPDATE, DELETE) na tabela `iva_dual_rules` é automaticamente registrada aqui via trigger PL/pgSQL. Serve como evidência para auditoria fiscal e permite reconstituir o histórico de quem alterou qual alíquota e quando.

**Colunas-chave:**
- `operation_type` — `I` (Insert), `U` (Update), `D` (Delete)
- `changed_by` — `current_user` do PostgreSQL
- `original_id` — FK lógica para a regra original
- Demais colunas são snapshots dos valores no momento da operação

**Trigger:** `trg_audit_iva_dual_rules` (AFTER INSERT OR UPDATE OR DELETE) → `fn_log_iva_dual_rules()`

---

### 10. `reforma_tributaria_rules` (LEGADO)

| Atributo | Valor |
|---|---|
| **Propósito** | Primeira tentativa de tabela da Reforma Tributária — **substituída por `iva_dual_rules`**. Mantida para compatibilidade. |
| **Tributo** | CBS, IBS, IS (legado) |
| **Quem consulta** | Nenhum código ativo — `ReformaCalculator` legado (`internal/reforma/reforma.go`) foi migrado para `CBSCalculator` + `IBSCalculator` |
| **Status** | ⚠️ **Deprecated** — não usar em novos desenvolvimentos |

**Função no sistema:** Esta tabela representava o primeiro rascunho do mapeamento CBS/IBS/IS, com `tipo_aliquota` (padrao, reduzida_60, reduzida_100, seletivo) como classificador. Foi substituída por `iva_dual_rules` que oferece granularidade por UF destino + município IBGE e atende ao princípio do destino.

**Estratégia de migração:** Dados existentes devem ser migrados para `iva_dual_rules`. Novas features usam exclusivamente `iva_dual_rules`.

---

### 11. `ncm_seletivo`

| Atributo | Valor |
|---|---|
| **Propósito** | Catálogo de NCMs sujeitos ao Imposto Seletivo (IS — "Imposto do Pecado") |
| **Tributo** | IS (Imposto Seletivo) |
| **Quem consulta** | `ISFilter` (`internal/legacy/is_filter.go:80`) via `GetNCMSeletivo(ctx, ncm)` |
| **Endpoint afetado** | `POST /v1/calculate` (Fase 0 — pré-filtro) |
| **Regras de negócio** | BR-TAX-INF-005 (NCM Seletivo), BR-TAX-CONS-010 (IS Pré-filtro), SOP-003 |

**Função no sistema:** O ISFilter (Fase 0 do pipeline SOP-013) verifica cada item contra esta tabela. Se o NCM estiver listado, aplica a alíquota da categoria (ex: 50% para bebidas alcoólicas, 100% para cigarros). Se não estiver, IS = 0. Esta tabela é **independente** de `iva_dual_rules` — o IS é resolvido separadamente porque incide sobre categorias específicas de produtos, não sobre todos os NCMs.

**Categorias previstas (EC 132/2023):**
- `BEBIDAS_ALCOOLICAS` — cervejas, vinhos, destilados
- `CIGARROS` — produtos de tabaco
- `REFRIGERANTES` — bebidas açucaradas
- `VEICULOS` — automóveis (previsto)
- `COMBUSTIVEIS_FOSSEIS` — petróleo e derivados (previsto)

**COMMENT ON TABLE:** "NCMs sujeitos ao Imposto Seletivo (IS) — 'Imposto do Pecado'. Gerido pelo Ministerio da Fazenda."

**Mapeamento Go:** `repository.NCMSeletivoRule` struct (`core-lib/repository/entities.go:136`)

---

### 12. `cbs_rates`

| Atributo | Valor |
|---|---|
| **Propósito** | Alíquotas CBS por classe tributária setorial |
| **Tributo** | CBS (Contribuição sobre Bens e Serviços) |
| **Quem consulta** | `CBSCalculator` — fallback quando `iva_dual_rules` não tem entrada para o NCM |
| **Endpoint afetado** | `POST /v1/calculate` (Fase 2 sequencial) |
| **Regras de negócio** | SOP-001 (CBS por classe) |

**Função no sistema:** Fornece alíquotas CBS por classe tributária setorial (ex: `TELECOM`, `GERAL`, `SAUDE`). Funciona como camada de fallback: se o NCM não estiver em `iva_dual_rules`, o sistema consulta a classe tributária do produto e busca a alíquota CBS nesta tabela.

**COMMENT ON TABLE:** "Aliquotas CBS (Contribuicao sobre Bens e Servicos) por classe tributaria. Ref: LC 214/2025."

---

### 13. `cst_reforma`

| Atributo | Valor |
|---|---|
| **Propósito** | Tabela oficial de CST (Código de Situação Tributária) para CBS/IBS conforme LC 214/2025 |
| **Tributo** | CBS, IBS |
| **Quem consulta** | `computeIvaDual()` (`internal/reforma/reforma.go:111`) via `GetCSTReforma(ctx, flags)` |
| **Endpoint afetado** | `POST /v1/calculate`, `POST /v1/simulate` |
| **Regras de negócio** | LC 214/2025, ADR-010, ADR-011 |

**Função no sistema:** Fonte canônica dos CSTs oficiais para o regime IVA Dual. Contém **164 CCTs** (Classificações Tributárias) agrupados em **18 CSTs** de 3 dígitos. O CST é resolvido em runtime com base nos flags de contexto da operação (`CSTFlags`) — o repository seleciona o CCT mais específico que atende aos critérios da transação.

**Estrutura:**
- `cst` (3 dígitos) — código de situação tributária (ex: `000` = tributação integral)
- `cct` (6 dígitos) — classificação tributária única (ex: `000001`)
- 8 flags booleanas que caracterizam a tributação: `exige_tributacao`, `reducao_bc`, `reducao_aliquota`, `transferencia_credito`, `diferimento`, `monofasica`, `credito_presumido`, `ajuste_competencia`
- `percentual_reducao_ibs` / `percentual_reducao_cbs` — reduções específicas
- `tipo_aliquota` — classificação (1-Redução, 2-Padrão, etc.)
- `url_legislacao` — link para o artigo da LC 214/2025
- `simples_nacional` — aplicabilidade ao Simples Nacional

**Mapeamento Go:** `repository.CSTReforma` struct, `repository.CSTFlags` struct (`core-lib/repository/entities.go:160-169`)

---

## Tabelas Operacionais

### 14. `tax_tokens`

| Atributo | Valor |
|---|---|
| **Propósito** | Congelamento temporal de alíquotas — garante previsibilidade fiscal para operações comerciais |
| **Tributo** | CBS, IBS, IS |
| **Quem consulta** | `TokenService` (`internal/token/service.go`), `MemoryTokenStore` (`internal/token/memory_store.go`) |
| **Endpoint afetado** | `POST /v1/token/generate` (gera), `GET /v1/token/{id}` (valida), `POST /v1/calculate` (aceita token via header `X-Tax-Token`) |
| **Regras de negócio** | GAP-002 (TaxToken), SOP-013 fase de shadow run |

**Função no sistema:** Permite que um agente comercial "congele" as alíquotas vigentes em um ponto no tempo, obtendo um token UUID. Este token pode ser usado subsequentemente em chamadas a `/v1/calculate` para garantir que o cálculo use as mesmas alíquotas — essencial para orçamentos, contratos e simulações onde a previsibilidade fiscal é mandatória.

**Fluxo de uso:**
1. `POST /v1/token/generate` com `{ncm, uf_origem, uf_destino, municipio_ibge}` → consulta `iva_dual_rules`, snapshot as alíquotas, gera UUID, salva no store
2. `POST /v1/calculate` com header `X-Tax-Token: <uuid>` → o motor usa as alíquotas do token em vez de consultar `iva_dual_rules` em tempo real
3. `GET /v1/token/{id}` → retorna o token se ainda não expirou

**TTL:** Configurável via `TAX_TOKEN_TTL_MINUTES` (default: 60 minutos)

**Armazenamento:** Em memória (`MemoryTokenStore`) com `sync.RWMutex`; persistência em disco planejada para o futuro

**Colunas-chave:**
- `id` — UUID do token (PK)
- `ncm`, `uf_origem`, `uf_destino`, `municipio_ibge` — tupla de lookup
- `aliquota_cbs`, `aliquota_ibs_estadual`, `aliquota_ibs_municipal`, `aliquota_is` — snapshot
- `expires_at` — timestamp de expiração

---

### 15. `fornecedor_fiscal`

| Atributo | Valor |
|---|---|
| **Propósito** | Cadastro de fornecedores com qualificação fiscal para cálculo de créditos tributários |
| **Tributo** | CBS, IBS (créditos) |
| **Quem consulta** | `CreditEngine` (`internal/credit/engine.go`), `ValidationService` (`internal/supplier/service.go`) |
| **Endpoint afetado** | `POST /v1/supplier/validate`, `GET /v1/supplier/:cnpj`, `POST /v1/credit/calculate` |
| **Regras de negócio** | GAP-007 (BR-08 Qualificação de Fornecedores) |

**Função no sistema:** Antes de calcular créditos tributários para um fornecedor (cash forward na Reforma), o sistema verifica se o fornecedor está qualificado: regime tributário válido, certificado de regularidade fiscal em dia, e flag `permite_credito` ativa. Fornecedores não qualificados têm o crédito bloqueado com `MotivoBloqueio`.

**Fluxo de uso:**
1. `POST /v1/supplier/validate` — cadastra/atualiza o perfil fiscal do fornecedor
2. `GET /v1/supplier/:cnpj` — consulta o status atual
3. `POST /v1/credit/calculate` — verifica `permite_credito` antes de autorizar o cálculo de créditos

**Colunas-chave:**
- `cnpj` — PK (14 dígitos)
- `regime_tributario` — regime do fornecedor
- `certificado_regularidade` — certidão fiscal válida?
- `permite_credito` — flag principal — controla se créditos podem ser calculados
- `status` — estado atual do cadastro (ex: `ATIVO`, `PENDENTE`, `BLOQUEADO`)

---

## Matriz Tabela × Calculadora

| Tabela | ICMS | PIS/COFINS | IPI | ISS | CBS | IBS | IS | FUST/FUNTTEL | Credit Engine | Token Service | Admin Fiscal |
|---|---|---|---|---|---|---|---|---|---|---|---|
| `icms_rules` | ✅ | — | — | — | — | — | — | — | — | — | — |
| `federal_tax_rules` | — | ✅ | — | — | — | — | — | — | — | — | — |
| `product_tax_exceptions` | ✅ | ✅ | — | — | — | — | — | — | — | — | — |
| `tax_equivalence` | ✅ | — | — | — | — | — | — | — | — | — | — |
| `simples_nacional_rates` | ✅ | — | — | — | — | — | — | — | — | — | — |
| `ipi_regras` | — | — | ✅ | — | — | — | — | — | — | — | — |
| `iss_rates` | — | — | — | ✅ | — | — | — | — | — | — | — |
| `iva_dual_rules` | — | — | — | — | ✅ | ✅ | — | — | ✅ | ✅ | ✅ |
| `iva_dual_rules_log` | — | — | — | — | — | — | — | — | — | — | ✅ |
| `reforma_tributaria_rules` | — | — | — | — | ⚠️ | ⚠️ | ⚠️ | — | — | — | — |
| `ncm_seletivo` | — | — | — | — | — | — | ✅ | — | — | — | — |
| `cbs_rates` | — | — | — | — | ✅ | — | — | — | — | — | — |
| `cst_reforma` | — | — | — | — | ✅ | ✅ | — | — | — | — | — |
| `tax_tokens` | — | — | — | — | ✅ | ✅ | ✅ | — | — | ✅ | — |
| `fornecedor_fiscal` | — | — | — | — | — | — | — | — | ✅ | — | — |

Legenda: ✅ = Consulta direta, ⚠️ = Legado/deprecated

---

## Matriz Tabela × Regra de Negócio

| Tabela | BR-TAX-CALC-* | BR-TAX-CONS-* | BR-TAX-INF-* | SOP-* | Outros |
|---|---|---|---|---|---|
| `icms_rules` | 004, 005, 006, 021, 022 | 013 | — | 013 | DT-001 |
| `federal_tax_rules` | 008, 009 | — | — | 013 | Tese do Século |
| `product_tax_exceptions` | 004, 005 | 013 | — | 013 | — |
| `tax_equivalence` | 007 | — | — | 013 | — |
| `simples_nacional_rates` | 007 | — | — | 013 | — |
| `ipi_regras` | 003 | — | — | 013 | TIPI |
| `iss_rates` | 016 | 007 | — | 013 | LC 116/2003 |
| `iva_dual_rules` | 010, 011 | 010 | 001-004 | 001, 013 | DT-001, EC 132/2023, LC 214/2025 |
| `ncm_seletivo` | — | 010 | 005 | 003 | EC 132/2023 |
| `cbs_rates` | 010 | — | — | 001 | LC 214/2025 |
| `cst_reforma` | 010, 011 | — | — | — | LC 214/2025, ADR-010, ADR-011 |
| `tax_tokens` | — | — | — | 013 | GAP-002 |
| `fornecedor_fiscal` | — | — | — | — | GAP-007, BR-08 |

> **Nota:** Para a lista completa de regras de negócio, consulte [domain/domain.md](../domain/domain.md). Os códigos BR-TAX-* são referenciados no [product/feature-roadmap.md](../product/feature-roadmap.md).

---

## Padrões Transversais

### Vigência Temporal

Todas as 15 tabelas de regras fiscais seguem o mesmo padrão de vigência temporal:

```
inicio_validade DATE NOT NULL DEFAULT CURRENT_DATE
final_validade DATE NULL  -- NULL = regra vigente (sem data fim)
```

**Trigger de fechamento:** A função PL/pgSQL `fechar_fim_validade_generica()` é disparada BEFORE INSERT nas tabelas de regras. Quando uma nova regra é inserida para a mesma chave de negócio, a regra anterior tem seu `final_validade` automaticamente preenchido com `(NOVA.inicio_validade - 1)`, garantindo que:
- Apenas uma regra esteja ativa por vez para a mesma chave
- O histórico de regras passadas seja preservado
- Consultas usem `WHERE CURRENT_DATE BETWEEN inicio_validade AND COALESCE(final_validade, '9999-12-31')`

### Wildcard Matching

Várias tabelas usam wildcards para implementar hierarquia de especificidade:

| Wildcard | Significado | Exemplo |
|---|---|---|
| `*` | "Todos" (catch-all) | `uf_destino = '*'` → aplica a todas as UFs |
| `**` | "Todas as origens" | `uf_origem = '**'` → aplica a todas as origens |
| NCM 4 dígitos | Grupo NCM | `ncm = '4011'` → todos os subitens de pneus |

A precedência é sempre do match **mais específico**: NCM completo > NCM grupo > wildcard.

### Triggers Padronizados

Duas funções PL/pgSQL genéricas são reutilizadas via triggers:

1. **`fechar_fim_validade_generica()`** — BEFORE INSERT → fecha `final_validade` da regra anterior
2. **`atualizar_data_atualizacao_generica()`** — BEFORE INSERT OR UPDATE → atualiza `atualizado_em` e preserva `criado_em`

A tabela `iva_dual_rules` tem trigger adicional de auditoria (`fn_log_iva_dual_rules()`) que é específica (não genérica).

### Cache Redis

O `CachedTaxRepository` (decorator pattern em `core-lib/repository/cached_tax_repository.go`) aplica cache Redis com TTL de 24h para as consultas mais frequentes:

| Método | Chave Redis | TTL |
|---|---|---|
| `GetIvaDualRule` | `tax:iva:<ncm>:<uf>:<municipio>` | 24h |
| `GetICMSRule` | `tax:icms:<orig>:<dest>` | 24h |
| `GetFederalTaxRule` | `tax:federal:<regime>:<cstPis>:<cstCofins>` | 24h |

Em caso de miss, a consulta vai ao PostgreSQL e o resultado é armazenado no Redis.

---

## Referências Cruzadas

- **Estrutura relacional (ERD):** [erd.md](erd.md) — diagrama Mermaid e tipos de colunas
- **Regras de negócio por tributo:** [domain/domain.md](../domain/domain.md) — explicação detalhada de cada regra
- **Decisões arquiteturais:** [adrs/INDEX.md](adrs/INDEX.md) — ADR-010 (CST Reforma), ADR-011 (IVA Dual lookup)
- **Schema DDL:** `data/init.sql` — DDL completo com comentários e dados de exemplo
- **Contratos do repository:** `libs/go-native/taxnexus-billing-core-lib/repository/contracts.go` — interface `TaxRepository`
- **Entidades Go:** `libs/go-native/taxnexus-billing-core-lib/repository/entities.go` — structs de domínio
- **Implementação PostgreSQL:** `libs/go-native/taxnexus-billing-core-lib/repository/postgres_repository.go` — queries SQL
- **Cache:** `libs/go-native/taxnexus-billing-core-lib/repository/cached_tax_repository.go` — decorator Redis
