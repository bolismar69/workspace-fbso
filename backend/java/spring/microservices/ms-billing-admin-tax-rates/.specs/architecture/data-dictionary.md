# Dicionário de Dados — ms-billing-admin-tax-rates

> **Schema:** `billing_tax_rates` (compartilhado com `ms-billing-engine-tax-rates`)
> **Fonte:** `data/init.sql` existente + novas tabelas de administração (DT-1) + tabelas independentes IVA Dual (ADR-003)
> **Atualizado:** 2026-07-12 (24 tabelas documentadas: 15 existentes + 6 admin + 3 IVA Dual independentes)
> **Decisão arquitetural:** [ADR-003](adrs/ADR-003-tax-table-strategy.md) — Estratégia de tabelas independentes para CBS, IBS, IS
> **Referência de integração:** [INTEGRATION-MAP.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md) — Seção 5

> ⚠️ **ATENÇÃO:** A tabela `iva_dual_rules` (CBS+IBS+IS unificados) está sendo **redefinida em 3 tabelas independentes** (`aliquotas_cbs`, `aliquotas_ibs`, `aliquotas_is`) conforme [ADR-003](adrs/ADR-003-tax-table-strategy.md). As novas implementações devem usar as tabelas independentes. A VIEW `v_aliquotas_iva` provê compatibilidade retroativa para o motor de cálculo (DT-3).

Documento canônico que descreve a **função de negócio**, o **propósito**, os **padrões de uso** e as **regras de negócio associadas** a cada tabela do schema `billing_tax_rates`. Complementa o [erd.md](erd.md) (que documenta a estrutura relacional e os tipos de colunas).

**Perspectiva:** Este dicionário é escrito do ponto de vista do **microserviço de administração tributária (DT-1)** — o componente responsável por criar, editar, desativar e auditar as tabelas fiscais. O motor de cálculo (DT-3) consulta estas mesmas tabelas em modo leitura para realizar os cálculos de impostos.

---

## Índice

1. [Visão Geral do Schema](#visão-geral-do-schema)
2. [Tabelas do Regime Atual (Pré-Reforma)](#tabelas-do-regime-atual-pré-reforma)
   - [icms_rules](#1-icms_rules)
   - [aliquotas_pis_cofins](#2-aliquotas_pis_cofins)
   - [product_tax_exceptions](#3-product_tax_exceptions)
   - [equivalencia_csosn_cst](#4-equivalencia_csosn_cst)
   - [faixas_simples_nacional](#5-faixas_simples_nacional)
   - [ipi_regras](#6-ipi_regras)
   - [aliquotas_iss](#7-aliquotas_iss)
3. [Tabelas da Reforma Tributária (IVA Dual)](#tabelas-da-reforma-tributária-iva-dual)
   - [iva_dual_rules](#8-iva_dual_rules)
   - [iva_dual_rules_log](#9-iva_dual_rules_log)
   - [reforma_tributaria_rules](#10-reforma_tributaria_rules-legado)
   - [ncm_seletivo](#11-ncm_seletivo)
   - [aliquotas_cbs](#12-aliquotas_cbs)
   - [cst_reforma](#13-cst_reforma)
4. [Tabelas Operacionais (DT-3)](#tabelas-operacionais-dt-3)
   - [tax_tokens](#14-tax_tokens)
   - [fornecedor_fiscal](#15-fornecedor_fiscal)
5. [Novas Tabelas — Administração (DT-1)](#novas-tabelas--administração-dt-1)
   - [empresas](#16-empresas)
   - [estabelecimentos](#17-estabelecimentos)
   - [fornecedores](#18-fornecedores)
   - [lotes_carga](#19-lotes_carga)
   - [lotes_carga_itens](#20-lotes_carga_itens)
   - [auditoria](#21-auditoria)
6. [🆕 Colunas Multi-Tenancy e Rastreabilidade](#-colunas-multi-tenancy-e-rastreabilidade)
7. [Matriz Tabela × Operação](#matriz-tabela--operação)
8. [Padrões Transversais](#padrões-transversais)

---

## Visão Geral do Schema

O schema `billing_tax_rates` contém **24 tabelas** organizadas em cinco categorias funcionais:

| Categoria | Tabelas | Finalidade | Administrada por |
|---|---|---|---|
| **Regime Atual (Pré-Reforma)** | `icms_rules`, `aliquotas_pis_cofins`, `product_tax_exceptions`, `equivalencia_csosn_cst`, `faixas_simples_nacional`, `ipi_regras`, `aliquotas_iss` | Regras dos tributos vigentes (ICMS, PIS/COFINS, IPI, ISS) | **DT-1** |
| **Reforma Tributária (IVA Dual)** | `iva_dual_rules` (⚠️ em redefinição), `iva_dual_rules_log`, `reforma_tributaria_rules` (legado), `ncm_seletivo`, `cst_reforma` | Regras do novo sistema dual CBS/IBS (EC 132/2023, LC 214/2025) — estrutura original unificada | **DT-1** (exceto legado e referência) |
| **🆕 IVA Dual — Tabelas Independentes (ADR-003)** | `aliquotas_cbs`, `aliquotas_ibs`, `aliquotas_is`, `v_aliquotas_iva` | **Nova estratégia:** CBS, IBS e IS em tabelas independentes com chaves naturais corretas. VIEW `v_aliquotas_iva` como contrato de interface para DT-3 | **DT-1** |
| **Operacional (DT-3)** | `tax_tokens`, `fornecedor_fiscal` | Suporte a cálculo de impostos (congelamento de alíquotas, qualificação de fornecedores) | **DT-3** |
| **🆕 Administração (DT-1)** | `empresas`, `estabelecimentos`, `fornecedores`, `lotes_carga`, `lotes_carga_itens`, `auditoria` | Infraestrutura multi-tenancy, carga em lote com aprovação, auditoria unificada | **DT-1** |

### Diagrama de Dependência Funcional (Perspectiva DT-1 — Administração)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                      Portal de Gestão Tributária (DT-2)                      │
│                    React 19 + Vite — Interface do Usuário                    │
└──────────────────────────────────┬──────────────────────────────────────────┘
                                   │ REST/JSON (API-CONTRACTS.md)
                                   ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                Microserviço de Administração Tributária (DT-1)               │
│                     Java 21 + Spring Boot 4.0.1                              │
│                                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐    │
│  │  Recursos REST:                                                      │    │
│  │  /api/v1/aliquotas  /api/v1/classificacoes  /api/v1/regimes         │    │
│  │  /api/v1/usuarios   /api/v1/auditoria       /api/v1/lotes           │    │
│  │  /api/v1/aprovacoes /api/v1/relatorios       /api/v1/empresas       │    │
│  └─────────────────────────────────────────────────────────────────────┘    │
│                                   │                                          │
│         ┌─────────────────────────┼─────────────────────────┐               │
│         ▼                         ▼                         ▼               │
│  ┌──────────────┐    ┌─────────────────────┐    ┌──────────────────────┐    │
│  │ CRUD         │    │ Carga em Lote       │    │ Auditoria &          │    │
│  │ Alíquotas    │    │ (staging → aprovação│    │ Relatórios           │    │
│  │ Classif.     │    │  → efetivação)      │    │ Dashboards           │    │
│  │ Regimes      │    │                     │    │ Governança           │    │
│  └──────┬───────┘    └──────────┬──────────┘    └───────────┬──────────┘    │
│         │                       │                           │               │
└─────────┼───────────────────────┼───────────────────────────┼───────────────┘
          │                       │                           │
          ▼                       ▼                           ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                      PostgreSQL — Schema billing_tax_rates                   │
│                                                                              │
│  ┌────────────────────┐  ┌────────────────────┐  ┌────────────────────┐     │
│  │ Tabelas Existentes │  │ 🆕 Novas Tabelas   │  │ Tabelas Op. (DT-3) │     │
│  │ (+ colunas 🆕)     │  │ DT-1               │  │                    │     │
│  │                    │  │                    │  │                    │     │
│  │ icms_rules         │  │ empresas           │  │ tax_tokens         │     │
│  │ aliquotas_pis_cofins  │  │ estabelecimentos            │  │ fornecedor_fiscal  │     │
│  │ product_tax_except │  │ fornecedores       │  │                    │     │
│  │ equivalencia_csosn_cst    │  │ lotes_carga        │  │                    │     │
│  │ simples_nac_rates  │  │ lotes_carga_itens  │  │                    │     │
│  │ ipi_regras         │  │ auditoria      │  │                    │     │
│  │ aliquotas_iss          │  │                    │  │                    │     │
│  │ iva_dual_rules     │  │                    │  │                    │     │
│  │ ncm_seletivo       │  │                    │  │                    │     │
│  │ aliquotas_cbs          │  │                    │  │                    │     │
│  │ cst_reforma        │  │                    │  │                    │     │
│  └────────────────────┘  └────────────────────┘  └────────────────────┘     │
│                                                                              │
│  🆕 Colunas Multi-Tenancy em TODAS as tabelas de regras fiscais:             │
│     empresa_id  •  tenant_id  •  origem_cadastro  •  lote_origem_id         │
│     lote_item_origem_id                                                      │
└─────────────────────────────────────────────────────────────────────────────┘
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

### 2. `aliquotas_pis_cofins`

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
SELECT * FROM aliquotas_pis_cofins
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

**Triggers:** `billing_tax_rates_aliquotas_pis_cofins_fim_validade`, `billing_tax_rates_aliquotas_pis_cofins_atualizado_em`

**Mapeamento Go:** `repository.FederalTaxRule` struct (`core-lib/repository/entities.go:29`)

---

### 3. `product_tax_exceptions` 🆕 VIEW LEGADA (ADR-003)

> 🆕 **Status:** Esta entidade foi **dividida em 2 tabelas independentes** conforme [ADR-003](adrs/ADR-003-tax-table-strategy.md): `excecoes_icms` e `excecoes_pis_cofins`. A tabela original unificada violava o princípio de "cada imposto com sua própria tabela" — ICMS e PIS/COFINS têm colunas, chaves de lookup e ciclos de vida distintos (PIS/COFINS extinguem em 2027; ICMS em 2029-2032). `product_tax_exceptions` agora é uma **VIEW** (`SELECT * FROM v_excecoes_fiscais`) para compatibilidade com DT-3.

| Atributo | Valor |
|---|---|
| **Propósito** | 🆕 VIEW legada — redireciona para `v_excecoes_fiscais` (FULL OUTER JOIN de `excecoes_icms` + `excecoes_pis_cofins`) |
| **Tributo** | ICMS, PIS, COFINS, ICMS-ST (legado unificado — VIEW read-only) |
| **Quem consulta** | `ICMSCalculator` via `getEffectiveTaxConfig()` — **apenas SELECT** |
| **Endpoint afetado** | `POST /v1/calculate` |
| **Nota** | ⚠️ VIEW não é updatable (FULL OUTER JOIN). DT-1 escreve em `excecoes_icms` e `excecoes_pis_cofins`; DT-3 apenas lê |

### 🆕 3a. `excecoes_icms`

| Atributo | Valor |
|---|---|
| **Propósito** | 🆕 Tabela independente de exceções de ICMS por NCM — substitui a porção ICMS de `product_tax_exceptions` |
| **Tributo** | ICMS, ICMS-ST |
| **Quem administra** | DT-1 via `POST/PUT/DELETE /api/v1/excecoes/icms` |

**Função no sistema:** Sobrepõe as regras gerais de `icms_rules` para NCMs específicos. Exemplos: MVA específico por NCM para ST, protocolo ST setorial, alíquota interestadual 4% para importados, desonerações setoriais. Colunas são exclusivamente de ICMS — sem campos de PIS/COFINS.

**Colunas-chave:** `ncm`, `uf_origem`, `uf_destino`, `cst_icms`, `csosn`, `mva_st`, `aliquota_interna_destino`, `aliquota_interestadual`, `percentual_fcp`, `reducao_base`, `possui_protocolo_st`, `motivo_desoneracao_icms`, `possui_desoneracao`, `desoneracao_codigo_beneficio_fiscal`, `regime_tributario_destino`

### 🆕 3b. `excecoes_pis_cofins`

| Atributo | Valor |
|---|---|
| **Propósito** | 🆕 Tabela independente de exceções de PIS/COFINS por NCM — substitui a porção PIS/COFINS de `product_tax_exceptions` |
| **Tributo** | PIS, COFINS |
| **Quem administra** | DT-1 via `POST/PUT/DELETE /api/v1/excecoes/pis-cofins` |

**Função no sistema:** Sobrepõe as regras gerais de `aliquotas_pis_cofins` para NCMs específicos. Exemplos: produtos monofásicos (CST 04), alíquota zero, tributação por unidade (CST 03). Colunas são exclusivamente de PIS/COFINS — sem campos de ICMS.

**Colunas-chave:** `ncm`, `uf_origem`, `uf_destino`, `cst_pis`, `cst_cofins`, `aliquota_pis_unitario`, `aliquota_cofins_unitario`

---

### 4. `equivalencia_csosn_cst`

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
SELECT * FROM equivalencia_csosn_cst
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

### 5. `faixas_simples_nacional`

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
SELECT * FROM faixas_simples_nacional
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

### 7. `aliquotas_iss`

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
SELECT * FROM aliquotas_iss
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

### 8. `iva_dual_rules` ⚠️ EM REDEFINIÇÃO (ADR-003)

> ⚠️ **Status:** Esta tabela está sendo **redefinida em 3 tabelas independentes** conforme [ADR-003 — Opção C](adrs/ADR-003-tax-table-strategy.md). CBS, IBS e IS possuem chaves naturais e ciclos de vida distintos — forçá-los na mesma tabela causa duplicação de CBS e acoplamento de evolução. **Novas implementações devem usar `aliquotas_cbs`, `aliquotas_ibs` e `aliquotas_is`.** A VIEW `v_aliquotas_iva` supre o motor de cálculo (DT-3) com a mesma assinatura de consulta. `iva_dual_rules` é mantida durante o período de migração.

| Atributo | Valor |
|---|---|
| **Propósito** | ⚠️ [EM REDEFINIÇÃO] Tabela original do IVA Dual — unificava CBS, IBS e IS. Substituída por `aliquotas_cbs` + `aliquotas_ibs` + `aliquotas_is` |
| **Tributo** | CBS, IBS, IS (legado unificado) |
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

### 12. `aliquotas_cbs` 🆕 REDEFINIDA (ADR-003)

> 🆕 **Status:** Esta tabela foi **redefinida** conforme [ADR-003 — Opção C](adrs/ADR-003-tax-table-strategy.md). Deixou de ser uma tabela de fallback setorial e passou a ser a **tabela primária de CBS**, absorvendo a porção CBS que antes residia em `iva_dual_rules`. Chave natural: `(empresa_id, ncm)`.

| Atributo | Valor |
|---|---|
| **Propósito** | 🆕 Tabela primária de alíquotas CBS — substitui a porção CBS de `iva_dual_rules` e absorve o antigo fallback setorial |
| **Tributo** | CBS (Contribuição sobre Bens e Serviços) |
| **Quem consulta** | `CBSCalculator` — consulta primária via `GetCBSRate(ctx, ncm, empresa_id)` |
| **Endpoint afetado** | `POST /v1/calculate` (Fase 2 sequencial), `POST /v1/simulate` |
| **Regras de negócio** | SOP-001 (CBS por classe), BR-TAX-CALC-010 (CBS) |

**Função no sistema:** Tabela canônica de alíquotas CBS. Cada NCM tem **uma única linha** (não mais 27 como em `iva_dual_rules`), pois CBS é imposto federal uniforme — não varia por UF ou município. A classe tributária setorial (`TELECOM`, `GERAL`, `SAUDE`, etc.) determina a alíquota aplicável. Para o motor de cálculo (DT-3), a VIEW `v_aliquotas_iva` faz JOIN com `aliquotas_ibs` e `aliquotas_is` entregando o mesmo formato que `iva_dual_rules` provia.

**Chave natural:** `(empresa_id, ncm)` — uma alíquota CBS por NCM, sem dependência de UF.

**Padrão de lookup:**
```
SELECT * FROM aliquotas_cbs
WHERE empresa_id = $1 AND ncm = $2
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
LIMIT 1
```

**Colunas-chave:**
- `ncm` — NCM do produto (8 dígitos, obrigatório)
- `classe_tributaria` — classe setorial (TELECOM, GERAL, SAUDE, etc.)
- `aliquota_cbs` — alíquota CBS federal unificada (%)

**Diferencial vs `iva_dual_rules`:**
- ✅ 1 linha por NCM (vs 27 linhas com CBS duplicada)
- ✅ CBS e IBS evoluem independentemente (alterar CBS não afeta IBS)
- ✅ Lookup simplificado: 2 parâmetros (empresa + NCM) vs 4 parâmetros

---

### 🆕 12b. `aliquotas_ibs`

> 🆕 **Nova tabela** conforme [ADR-003 — Opção C](adrs/ADR-003-tax-table-strategy.md). Substitui a porção IBS de `iva_dual_rules`. Chave natural: `(empresa_id, ncm, uf_destino, municipio_destino_ibge)`.

| Atributo | Valor |
|---|---|
| **Propósito** | 🆕 Tabela primária de alíquotas IBS — substitui a porção IBS de `iva_dual_rules` |
| **Tributo** | IBS (Imposto sobre Bens e Serviços — estadual + municipal) |
| **Quem consulta** | `IBSCalculator` — consulta primária via `GetIBSRate(ctx, ncm, uf, municipio, empresa_id)` |
| **Endpoint afetado** | `POST /v1/calculate` (Fase 3 paralela), `POST /v1/simulate` |
| **Regras de negócio** | BR-TAX-CALC-011 (IBS), SOP-013 |

**Função no sistema:** Tabela canônica de alíquotas IBS. Segue o princípio do destino: o imposto pertence ao local de consumo. Cada combinação de NCM + UF destino + município (opcional) tem suas alíquotas estadual e municipal. Quando `municipio_destino_ibge IS NULL`, a alíquota é a padrão estadual (aplica-se a todos os municípios da UF). Quando preenchido, é a regra específica do município.

**Chave natural:** `(empresa_id, ncm, uf_destino, municipio_destino_ibge)` — permite que cada município tenha alíquota própria.

**Padrão de lookup (especificidade decrescente):**
```
SELECT * FROM aliquotas_ibs
WHERE empresa_id = $1 AND ncm = $2 AND uf_destino = $3
  AND (municipio_destino_ibge = $4 OR municipio_destino_ibge IS NULL)
  AND inicio_validade <= CURRENT_DATE
  AND (final_validade IS NULL OR final_validade >= CURRENT_DATE)
ORDER BY municipio_destino_ibge DESC NULLS LAST
LIMIT 1
```

**Colunas-chave:**
- `ncm` — NCM do produto (8 dígitos, obrigatório)
- `uf_destino` — UF de destino do imposto (princípio do destino)
- `municipio_destino_ibge` — código IBGE (7 dígitos, NULL = regra padrão estadual)
- `aliquota_ibs_estadual` — alíquota IBS estadual (%)
- `aliquota_ibs_municipal` — alíquota IBS municipal (%)
- `percentual_reducao` — redução de transição (0, 60, 100)

**Diferencial vs `iva_dual_rules`:**
- ✅ CBS não está acoplada — IBS evolui independentemente
- ✅ Sem colunas CBS/IS com NULLs artificiais
- ✅ Migração de dados focada: apenas IBS é afetado por mudanças legislativas estaduais/municipais

---

### 🆕 12c. `aliquotas_is`

> 🆕 **Nova tabela** conforme [ADR-003 — Opção C](adrs/ADR-003-tax-table-strategy.md). Substitui a porção IS de `iva_dual_rules`. Chave natural: `(empresa_id, ncm)`.

| Atributo | Valor |
|---|---|
| **Propósito** | 🆕 Tabela primária de alíquotas do Imposto Seletivo — substitui a porção IS de `iva_dual_rules` |
| **Tributo** | IS (Imposto Seletivo — "Imposto do Pecado") |
| **Quem consulta** | `ISFilter` — consulta primária via `GetISRate(ctx, ncm, empresa_id)` |
| **Endpoint afetado** | `POST /v1/calculate` (Fase 0 — pré-filtro) |
| **Regras de negócio** | BR-TAX-INF-005 (NCM Seletivo), BR-TAX-CONS-010 (IS Pré-filtro), SOP-003 |

**Função no sistema:** Define as alíquotas do Imposto Seletivo por NCM e categoria de produto. Diferente de CBS e IBS, o IS é um imposto extrafiscal — incide apenas sobre categorias específicas (bebidas alcoólicas, cigarros, refrigerantes, etc.). A tabela `ncm_seletivo` permanece como catálogo de referência dos NCMs sujeitos ao IS; `aliquotas_is` é a tabela de alíquotas efetivas.

**Chave natural:** `(empresa_id, ncm)` — o IS é determinado apenas pelo NCM do produto, sem dependência de UF ou município.

**Categorias previstas (EC 132/2023):**
- `BEBIDAS_ALCOOLICAS` — cervejas, vinhos, destilados
- `CIGARROS` — produtos de tabaco
- `REFRIGERANTES` — bebidas açucaradas
- `VEICULOS` — automóveis (previsto)
- `COMBUSTIVEIS_FOSSEIS` — petróleo e derivados (previsto)

**Colunas-chave:**
- `ncm` — NCM do produto (8 dígitos)
- `categoria` — categoria do IS (BEBIDAS_ALCOOLICAS, CIGARROS, etc.)
- `aliquota_is` — alíquota do Imposto Seletivo (%)

**Diferencial vs `iva_dual_rules`:**
- ✅ Sem NULLs artificiais (em `iva_dual_rules`, `aliquota_is` era NULL para 95%+ das linhas)
- ✅ Categorias podem ser adicionadas sem afetar CBS e IBS
- ✅ Tabela enxuta: apenas NCMs sujeitos ao IS têm linhas

---

### 🆕 12d. `v_aliquotas_iva` (VIEW Materializada)

> 🆕 **Nova VIEW** conforme [ADR-003 — Opção C](adrs/ADR-003-tax-table-strategy.md). Contrato de interface entre DT-1 (administração) e DT-3 (motor de cálculo).

| Atributo | Valor |
|---|---|
| **Propósito** | 🆕 VIEW materializada que faz JOIN de `aliquotas_cbs` + `aliquotas_ibs` + `aliquotas_is`, entregando a mesma assinatura de consulta que `iva_dual_rules` provia |
| **Tributo** | CBS, IBS, IS (visão unificada para o motor) |
| **Quem consulta** | `CBSCalculator`, `IBSCalculator`, `ISFilter` — via `computeIvaDual()` (sem alterações no código do motor) |
| **Endpoint afetado** | `POST /v1/calculate`, `POST /v1/simulate` (transparente para o motor) |

**Função no sistema:** Garante **backward-compatibility** durante e após a migração de `iva_dual_rules` → tabelas independentes. O motor de cálculo (DT-3) continua consultando com a mesma query — a VIEW traduz para as 3 tabelas subjacentes. A administração (DT-1) escreve diretamente nas tabelas independentes.

**Estrutura:** `FULL OUTER JOIN` entre `aliquotas_cbs`, `aliquotas_ibs` e `aliquotas_is` sobre `(ncm, empresa_id)`. Filtra apenas registros vigentes (`final_validade IS NULL OR >= CURRENT_DATE`).

**Refresh:** `REFRESH MATERIALIZED VIEW CONCURRENTLY v_aliquotas_iva` — executado após cada aprovação de lote de carga.

**Cache Redis:** Chaves mantidas com mesmo formato: `tax:iva:<ncm>:<uf>:<municipio>` → a VIEW populada garante cache hit.

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

## Novas Tabelas — Administração (DT-1)

> As 6 tabelas abaixo são **criadas pelo DT-1** e não existem no schema original do motor de cálculo. Elas provêm a infraestrutura de multi-tenancy, carga em lote com aprovação e auditoria unificada.

### 16. `empresas`

| Atributo | Valor |
|---|---|
| **Propósito** | Raiz do multi-tenancy — cada empresa do grupo econômico é cadastrada aqui |
| **Quem administra** | DT-1 via `POST/PUT/GET /api/v1/empresas` (perfil: ADMINISTRADOR_FISCAL) |
| **Quem consulta** | DT-1 (portal); DT-3 (motor — somente leitura) |

**Função no sistema:** As tabelas de regras fiscais são segmentadas por `empresa_id`. Uma empresa pode ter múltiplos estabelecimentos (`estabelecimentos`). O cadastro de empresas é pré-requisito para qualquer operação de gestão tributária — alíquotas, classificações e regimes são sempre vinculados a uma empresa.

**Colunas-chave:**
- `cnpj_raiz` — CNPJ raiz (8 primeiros dígitos)
- `razao_social` / `nome_fantasia` — identificação corporativa
- `status` — `ATIVA`, `INATIVA`

### 17. `estabelecimentos`

| Atributo | Valor |
|---|---|
| **Propósito** | Estabelecimentos (filiais, unidades) dentro de uma empresa |
| **Quem administra** | DT-1 via `GET /api/v1/empresas/{id}/estabelecimentos` |
| **Quem consulta** | DT-1 (portal); DT-3 (motor — somente leitura) |

**Função no sistema:** Permite granularidade de alíquotas por CNPJ completo quando necessário. Cada tenant pertence a uma empresa (`empresa_id` FK). Nas tabelas de regras fiscais, `tenant_id` é nullable — quando `NULL`, a regra aplica-se a todos os estabelecimentos da empresa.

### 18. `fornecedores`

| Atributo | Valor |
|---|---|
| **Propósito** | Cadastro mestre de fornecedores com qualificação fiscal expandida |
| **Quem administra** | DT-1 via `POST/PUT/GET /api/v1/fornecedores` |
| **Quem consulta** | DT-1 (portal); DT-3 (motor — somente leitura) |

**Função no sistema:** Complementa a tabela operacional `fornecedor_fiscal` (gerenciada pelo DT-3) adicionando dimensões de negócio: CNAE principal, segmentação por empresa do grupo, e status de cadastro (`ATIVO`, `PENDENTE`, `BLOQUEADO`). A reconciliação entre `fornecedores` (negócio) e `fornecedor_fiscal` (operacional) é feita por job diário (ver [INTEGRATION-MAP.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md) Seção 8.1).

**Colunas-chave:**
- `empresa_id` — FK para `empresas`
- `cnpj` — CNPJ do fornecedor (14 dígitos)
- `regime_tributario` — Lucro Real, Lucro Presumido, Simples Nacional
- `cnae_principal` — CNAE fiscal principal (7 dígitos)

### 19. `lotes_carga`

| Atributo | Valor |
|---|---|
| **Propósito** | Cabeçalho dos lotes de importação de alíquotas — staging para aprovação |
| **Quem administra** | DT-1: `POST /api/v1/lotes` (envio), `POST /api/v1/lotes/{id}/aprovar` (aprovação) |
| **Quem consulta** | DT-1 (portal — acompanhamento de status) |

**Função no sistema:** Cada arquivo de carga enviado pelo time fiscal gera um registro aqui. O lote passa pelos status `EM_VALIDACAO` → `AGUARDANDO_APROVACAO` → `APROVADO` ou `REJEITADO`. Os dados **não são efetivados** nas tabelas finais de alíquotas até que um Administrador Fiscal aprove o lote. Este mecanismo previne que cargas inválidas ou não revisadas contaminem as regras vigentes.

**Fluxo de status:**
```
EM_VALIDACAO ──► AGUARDANDO_APROVACAO ──► APROVADO (itens ACEITOS efetivados)
                                     └──► REJEITADO (nenhum item efetivado)
```

**Colunas-chave:**
- `empresa_id` — empresa destinatária
- `tributo` — IBS, CBS, ICMS, etc.
- `status` — controla o ciclo de vida do lote
- `total_linhas` / `linhas_aceitas` / `linhas_rejeitadas` / `linhas_com_alertas` — contadores
- `enviado_por` / `aprovado_por` — rastreabilidade de responsáveis
- `justificativa` — fornecida no envio

### 20. `lotes_carga_itens`

| Atributo | Valor |
|---|---|
| **Propósito** | Linhas individuais do arquivo de carga — cada linha é validada e classificada |
| **Quem administra** | DT-1 (processamento automático na validação; consulta via `GET /api/v1/lotes/{id}/itens`) |

**Função no sistema:** Para cada linha do arquivo de carga, o sistema aplica as validações de negócio (RN-01 a RN-05) e classifica o item como `ACEITO`, `REJEITADO` ou `COM_ALERTA`. O `conteudo_original` é preservado em JSONB para auditoria e reprocessamento. Após a aprovação do lote, os itens `ACEITOS` são efetivados nas tabelas finais, e as colunas `entidade_criada_tipo` e `entidade_criada_id` registram o vínculo.

**Colunas-chave:**
- `lote_id` — FK para `lotes_carga`
- `numero_linha` — posição no arquivo original
- `conteudo_original` — JSONB com todos os campos da linha
- `status` — `ACEITO`, `REJEITADO`, `COM_ALERTA`
- `motivo_rejeicao` — descrição da RN violada
- `entidade_criada_tipo` / `entidade_criada_id` — rastreabilidade pós-efetivação

### 21. `auditoria`

| Atributo | Valor |
|---|---|
| **Propósito** | Trilha de auditoria unificada para **todas** as entidades gerenciadas pelo DT-1 |
| **Quem administra** | DT-1 (registro automático em toda operação de escrita; consulta via `GET /api/v1/auditoria`) |
| **Quem consulta** | DT-1 (portal — perfis Administrador Fiscal, Auditor/Controller); DT-3 (motor — somente leitura) |

**Função no sistema:** Substitui o modelo fragmentado anterior (onde apenas `iva_dual_rules` possuía tabela de log). Agora, toda operação de CRIACAO, EDICAO, DESATIVACAO, APROVACAO ou REJEICAO em qualquer entidade (alíquotas, classificações, regimes, usuários, lotes, empresas) gera automaticamente um registro imutável nesta tabela.

**Regras de imutabilidade (RN-14):**
- Nenhum perfil de usuário pode alterar ou excluir registros desta tabela
- Triggers de banco revogam `UPDATE` e `DELETE` para todos os usuários
- Retenção: mínimo de 5 anos; partição por ano para gerenciamento de volume

**Colunas-chave:**
- `entidade_tipo` / `entidade_id` — identificam o que foi alterado
- `operacao` — `CRIACAO`, `EDICAO`, `DESATIVACAO`, `APROVACAO`, `REJEICAO`
- `usuario_id` / `usuario_nome` / `usuario_perfil` — quem realizou a ação
- `estado_anterior` / `estado_novo` — snapshots JSONB completos (diff)
- `justificativa` — fornecida pelo usuário
- `ip_origem` — endereço IP da requisição
- `data_hora` — timestamp exato (RN-15)

---

## 🆕 Colunas Multi-Tenancy e Rastreabilidade

As seguintes colunas são **adicionadas pelo DT-1** às 8 tabelas de regras fiscais existentes (listadas abaixo) para suportar multi-tenancy e rastreabilidade de origem dos cadastros:

| Coluna | Tipo | Obrigatória | Descrição |
|:---|:---|:---|:---|
| `empresa_id` | integer FK → `empresas.id` | Sim (novos registros) | Empresa do grupo econômico à qual a regra se aplica |
| `tenant_id` | integer FK → `estabelecimentos.id` | Não (nullable) | Estabelecimento específico (null = aplica a todos) |
| `origem_cadastro` | varchar(10) | Sim (default: `MANUAL`) | `MANUAL` (formulário) ou `LOTE` (carga em lote) |
| `lote_origem_id` | integer FK → `lotes_carga.id` | Não (nullable) | Se origem = LOTE, aponta para o lote |
| `lote_item_origem_id` | integer FK → `lotes_carga_itens.id` | Não (nullable) | Se origem = LOTE, aponta para o item específico |

**Tabelas que recebem estas colunas:**
`icms_rules`, `aliquotas_pis_cofins`, `product_tax_exceptions`, `aliquotas_iss`, `ipi_regras`, `iva_dual_rules`, `ncm_seletivo`, `aliquotas_cbs` (🆕 tabela primária CBS), `aliquotas_ibs` (🆕 tabela primária IBS), `aliquotas_is` (🆕 tabela primária IS)

**Estratégia de migração:**
1. Colunas adicionadas como **nullable** (para não impactar registros existentes)
2. Registros existentes recebem `empresa_id` = 1 (empresa padrão), `origem_cadastro` = `MANUAL`
3. Após migração validada, colunas `empresa_id` e `origem_cadastro` passam a ser **NOT NULL** para novos registros
4. `equivalencia_csosn_cst`, `faixas_simples_nacional`, `cst_reforma` e `reforma_tributaria_rules` **não recebem** estas colunas (tabelas de referência normativa ou legadas)

---

## Matriz Tabela × Operação

Operações de administração disponíveis para cada tabela via portal (DT-1) e API:

| Tabela | Criar | Editar | Desativar | Consultar | Carga Lote | Aprovar | Auditar | Admin por |
|---|---|---|---|---|---|---|---|
| `icms_rules` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| `aliquotas_pis_cofins` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| `product_tax_exceptions` | — | — | — | ✅ | — | — | — | VIEW legada (read-only) |
| 🆕 `excecoes_icms` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| 🆕 `excecoes_pis_cofins` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| `equivalencia_csosn_cst` | ✅ | ✅ | ✅ | ✅ | — | — | ✅ | DT-1 |
| `faixas_simples_nacional` | ✅ | ✅ | ✅ | ✅ | — | — | ✅ | DT-1 |
| `ipi_regras` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| `aliquotas_iss` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| `iva_dual_rules` | ⚠️ | ⚠️ | ⚠️ | ✅ | ⚠️ | ⚠️ | ✅ | DT-1 (⚠️ em redefinição — ADR-003) |
| `iva_dual_rules_log` | — | — | — | ✅ | — | — | — | Trigger |
| `reforma_tributaria_rules` | ⚠️ | ⚠️ | ⚠️ | ⚠️ | — | — | — | Legado |
| `ncm_seletivo` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| `cst_reforma` | — | — | — | ✅ | — | — | — | Ref. |
| 🆕 **`aliquotas_cbs`** | **✅** | **✅** | **✅** | **✅** | **✅** | **✅** | **✅** | **DT-1 (tabela primária CBS — ADR-003)** |
| 🆕 **`aliquotas_ibs`** | **✅** | **✅** | **✅** | **✅** | **✅** | **✅** | **✅** | **DT-1 (tabela primária IBS — ADR-003)** |
| 🆕 **`aliquotas_is`** | **✅** | **✅** | **✅** | **✅** | **✅** | **✅** | **✅** | **DT-1 (tabela primária IS — ADR-003)** |
| 🆕 **`v_aliquotas_iva`** | **—** | **—** | **—** | **✅** | **—** | **—** | **—** | **VIEW (refresh automático)** |
| `tax_tokens` | — | — | — | — | — | — | — | DT-3 |
| `fornecedor_fiscal` | — | — | — | ✅ | — | — | — | DT-3 |
| 🆕 `empresas` | ✅ | ✅ | ✅ | ✅ | — | — | ✅ | DT-1 |
| 🆕 `estabelecimentos` | ✅ | ✅ | ✅ | ✅ | — | — | ✅ | DT-1 |
| 🆕 `fornecedores` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | DT-1 |
| 🆕 `lotes_carga` | ✅¹ | — | — | ✅ | ✅ | ✅ | ✅ | DT-1 |
| 🆕 `lotes_carga_itens` | — | — | — | ✅ | ✅² | — | ✅ | DT-1 |
| 🆕 `auditoria` | — | — | — | ✅ | — | — | ✅³ | DT-1 |

Legenda: ✅ = Operação suportada, ⚠️ = Legado/deprecated/em transição, — = Não se aplica, ¹ = Via upload de arquivo, ² = Validação automática na carga, ³ = Registro automático e imutável

> ⚠️ **ADR-003:** A partir da vigência desta decisão, operações de **criação/edição/carga** de CBS, IBS e IS devem usar as tabelas independentes (`aliquotas_cbs`, `aliquotas_ibs`, `aliquotas_is`). `iva_dual_rules` permanece acessível para consulta durante o período de migração. A VIEW `v_aliquotas_iva` supre o motor de cálculo (DT-3).

## Matriz Tabela × Regra de Negócio

| Tabela | BR-TAX-CALC-* | BR-TAX-CONS-* | BR-TAX-INF-* | SOP-* | Outros |
|---|---|---|---|---|---|
| `icms_rules` | 004, 005, 006, 021, 022 | 013 | — | 013 | DT-001 |
| `aliquotas_pis_cofins` | 008, 009 | — | — | 013 | Tese do Século |
| `product_tax_exceptions` | — | — | — | — | VIEW legada (ADR-003) |
| 🆕 `excecoes_icms` | 004, 005 | 013 | — | 013 | — |
| 🆕 `excecoes_pis_cofins` | 004, 005 | 013 | — | 013 | — |
| `equivalencia_csosn_cst` | 007 | — | — | 013 | — |
| `faixas_simples_nacional` | 007 | — | — | 013 | — |
| `ipi_regras` | 003 | — | — | 013 | TIPI |
| `aliquotas_iss` | 016 | 007 | — | 013 | LC 116/2003 |
| `iva_dual_rules` | 010, 011 | 010 | 001-004 | 001, 013 | ⚠️ Em redefinição — ADR-003 |
| `ncm_seletivo` | — | 010 | 005 | 003 | EC 132/2023 (catálogo de referência) |
| 🆕 **`aliquotas_cbs`** | **010** | **—** | **—** | **001** | **ADR-003, LC 214/2025 (tabela primária CBS)** |
| 🆕 **`aliquotas_ibs`** | **011** | **010** | **001-004** | **001, 013** | **ADR-003, EC 132/2023, LC 214/2025 (tabela primária IBS)** |
| 🆕 **`aliquotas_is`** | **—** | **010** | **005** | **003** | **ADR-003, EC 132/2023 (tabela primária IS)** |
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

A tabela `iva_dual_rules` possui trigger adicional de auditoria legado (`fn_log_iva_dual_rules()`) que alimenta `iva_dual_rules_log`. Este mecanismo é mantido por compatibilidade; novas operações do DT-1 registram auditoria na tabela `auditoria` (ver [Auditoria Unificada](#-auditoria-unificada)).

### Cache Redis

O `CachedTaxRepository` (decorator pattern em `core-lib/repository/cached_tax_repository.go`) aplica cache Redis com TTL de 24h para as consultas mais frequentes:

| Método | Chave Redis | TTL |
|---|---|---|
| `GetIvaDualRule` | `tax:iva:<ncm>:<uf>:<municipio>` | 24h |
| `GetICMSRule` | `tax:icms:<orig>:<dest>` | 24h |
| `GetFederalTaxRule` | `tax:federal:<regime>:<cstPis>:<cstCofins>` | 24h |

Em caso de miss, a consulta vai ao PostgreSQL e o resultado é armazenado no Redis.

### 🆕 Multi-Tenancy e Rastreabilidade de Origem

O DT-1 introduz 5 colunas transversais em todas as tabelas de regras fiscais (ver [🆕 Colunas Multi-Tenancy e Rastreabilidade](#-colunas-multi-tenancy-e-rastreabilidade)). O comportamento no portal é:

- **Segmentação por empresa:** Cada operação de consulta, criação ou edição é filtrada por `empresa_id` — o portal só exibe dados da empresa à qual o usuário logado pertence
- **Origem do cadastro:** Toda alíquota carrega `origem_cadastro` = `MANUAL` (formulário) ou `LOTE` (carga em lote). Esta informação é exibida no Painel de Alíquotas e na trilha de auditoria
- **Rastreabilidade de lote:** Se `origem_cadastro = LOTE`, os campos `lote_origem_id` e `lote_item_origem_id` permitem navegar do registro final até o arquivo e a linha específica que o originou

### 🆕 Auditoria Unificada

A tabela `auditoria` consolida a trilha de auditoria que antes era fragmentada (apenas `iva_dual_rules_log` para o IVA Dual). Características:

- **Cobertura total:** Todas as operações de CRIACAO, EDICAO, DESATIVACAO, APROVACAO e REJEICAO em qualquer entidade são registradas
- **Imutabilidade:** Triggers de banco revogam `UPDATE` e `DELETE` na tabela. Nenhum usuário (incluindo Administrador Fiscal) pode alterar registros de auditoria
- **Snapshots completos:** `estado_anterior` e `estado_novo` (JSONB) preservam o registro completo antes e depois de cada alteração
- **Retenção:** Mínimo de 5 anos; partição por ano (`auditoria_2026`, `auditoria_2027`, ...)
- **Substituição do modelo antigo:** O trigger `trg_audit_iva_dual_rules` (que alimenta `iva_dual_rules_log`) é mantido por compatibilidade durante o período de transição. Novas operações usam exclusivamente `auditoria`

---

## ⚠️ VIEWs de Compatibilidade Legada (Estratégia de Deploy)

> **Seção crítica para deploy.** O motor `ms-billing-engine-tax-rates` (DT-3, Go) referencia tabelas pelos nomes originais em inglês. VIEWs com nomes legados são criadas sobre as novas tabelas em português para manter o DT-3 funcional sem alterações de código.

### Motivação

O schema `billing_tax_rates` foi originalmente criado pelo time do motor de cálculo (Go) com nomenclatura em inglês. As novas tabelas do DT-1 (Java/Spring Boot) adotam nomenclatura em português — alinhada ao domínio fiscal brasileiro. Para evitar refatoração do motor legado, são criadas VIEWs simples (updatable) que redirecionam os nomes antigos para as novas tabelas.

### Mapeamento VIEW → Tabela

| VIEW (nome legado — DT-3 consulta) | Tabela Real (nome português — DT-1 escreve) | Tipo |
|---|---|---|
| `federal_tax_rules` | `aliquotas_pis_cofins` | Simples (updatable) |
| `cbs_rates` | `aliquotas_cbs` | Simples (updatable) |
| `ibs_rates` | `aliquotas_ibs` | Simples (updatable) |
| `is_rates` | `aliquotas_is` | Simples (updatable) |
| `iss_rates` | `aliquotas_iss` | Simples (updatable) |
| `simples_nacional_rates` | `faixas_simples_nacional` | Simples (updatable) |
| `tax_equivalence` | `equivalencia_csosn_cst` | Simples (updatable) |
| `tenants` | `estabelecimentos` | Simples (updatable) |
| `auditoria_log` | `auditoria` | Simples (updatable) |
| `iva_dual_rules` | `v_aliquotas_iva` | Redireciona para VIEW materializada (read-only) |
| `product_tax_exceptions` | `v_excecoes_fiscais` | FULL OUTER JOIN (read-only) |

### DDL das VIEWs

```sql
-- Migration: V30__criar_views_compatibilidade.sql
-- ⚠️ 11 VIEWs — todas devem existir antes do deploy do DT-3

-- 9 VIEWs updatable (SELECT * simples):
CREATE VIEW federal_tax_rules      AS SELECT * FROM aliquotas_pis_cofins;
CREATE VIEW cbs_rates              AS SELECT * FROM aliquotas_cbs;
CREATE VIEW ibs_rates              AS SELECT * FROM aliquotas_ibs;
CREATE VIEW is_rates               AS SELECT * FROM aliquotas_is;
CREATE VIEW iss_rates              AS SELECT * FROM aliquotas_iss;
CREATE VIEW simples_nacional_rates AS SELECT * FROM faixas_simples_nacional;
CREATE VIEW tax_equivalence        AS SELECT * FROM equivalencia_csosn_cst;
CREATE VIEW tenants                AS SELECT * FROM estabelecimentos;
CREATE VIEW auditoria_log          AS SELECT * FROM auditoria;

-- 2 VIEWs read-only (JOIN/MVIEW — DT-3 apenas SELECT):
CREATE VIEW iva_dual_rules         AS SELECT * FROM v_aliquotas_iva;
CREATE VIEW product_tax_exceptions AS SELECT * FROM v_excecoes_fiscais;
```

### Requisitos de Deploy

| # | Requisito | Criticidade |
|---|-----------|------------|
| 1 | VIEWs na **mesma transação** das tabelas | 🔴 CRÍTICO |
| 2 | Ordem: tabelas → dados → VIEWs | 🔴 CRÍTICO |
| 3 | `iva_dual_rules` (VIEW) antes de dropar tabela `iva_dual_rules` original | 🔴 CRÍTICO |
| 4 | `product_tax_exceptions` e `iva_dual_rules` **NÃO** são updatable | 🟠 ALTO |
| 5 | DT-3 **não** precisa de redeploy | ✅ |
| 6 | Script de rollback: `DROP VIEW` restaura nomes | 🟡 MÉDIO |

### Verificação Pós-Deploy

```sql
SELECT table_name, is_insertable_into, is_updatable
FROM information_schema.views
WHERE table_schema = 'billing_tax_rates'
  AND table_name IN ('federal_tax_rules', 'cbs_rates', 'ibs_rates', 'is_rates',
                     'iss_rates', 'simples_nacional_rates', 'tax_equivalence',
                     'iva_dual_rules', 'product_tax_exceptions',
                     'tenants', 'auditoria_log')
ORDER BY table_name;
```

### Rollback

```sql
DROP VIEW IF EXISTS federal_tax_rules, cbs_rates, ibs_rates, is_rates,
                    iss_rates, simples_nacional_rates, tax_equivalence,
                    iva_dual_rules, product_tax_exceptions,
                    tenants, auditoria_log;
```

---

## Referências Cruzadas

- **Estrutura relacional (ERD):** [erd.md](erd.md) — diagrama Mermaid, 24 tabelas e relacionamentos
- **ADR-003 — Estratégia de Tabelas:** [ADR-003-tax-table-strategy.md](adrs/ADR-003-tax-table-strategy.md) — Decisão de separar CBS, IBS e IS em tabelas independentes
- **Análise de Engenharia:** [engineering-skills.md](engineering-skills.md) — 8 perspectivas de engenharia + banco de dados
- **Análise Prévia (DB):** [DATA-ANALYSIS.md](DATA-ANALYSIS.md) — 3 especialistas DB, 21 recomendações
- **Mapa de Integrações:** [INTEGRATION-MAP.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md) — visão completa de componentes, canais e segurança
- **Contrato de API:** [API-CONTRACTS.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/API-CONTRACTS.md) — endpoints REST que operam sobre estas tabelas
- **Regras de negócio:** [04-FEATURES.md](../../../../../business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/04-FEATURES.md) — RN-01 a RN-38
- **Schema DDL:** `data/init.sql` (existente) + migrations Flyway (novas tabelas e colunas)
- **Entidades Java:** `src/main/java/com/fbso/.../model/` — JPA entities e DTOs
- **Repositórios:** `src/main/java/com/fbso/.../repository/` — Spring Data JDBC repositories

---

🤖 *Dicionário de Dados adaptado para o ms-billing-admin-tax-rates (DT-1) em 12 de Julho de 2026. Atualizado conforme ADR-003 (tabelas independentes IVA Dual).*
🤖 *Documento gerado por Inteligência Artificial. Agentes: Arquiteto de Dados (skills: `sql-pro`, `database-optimizer`), Especialista PostgreSQL (skill: `postgresql`), Senior Architect (engineering skill: `senior-architect`). Skills aplicadas: `postgresql`, `sql-pro`, `database-optimizer`, `senior-architect`, `senior-backend`. Ferramenta: Claude Code (Anthropic).*
