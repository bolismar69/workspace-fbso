# Glossário e Regras de Domínio — ms-billing-engine-tax-rates

Gerado pelo agente **Spec Miner** em 2026-06-20 com evidência de código. Atualizado em 2026-06-22 após pipeline SOP-013 7-fases (C-001) e separação CBS/IBS.

## Tributos Implementados

| Sigla | Nome | Status | Fonte |
|-------|------|--------|-------|
| IPI | Imposto sobre Produtos Industrializados |   Implementado | `internal/legacy/ipi.go` |
| ICMS | Imposto sobre Circulação de Mercadorias e Serviços |   Implementado | `internal/legacy/icms.go` |
| ICMS-ST | ICMS Substituição Tributária |   Implementado | `internal/legacy/icms.go:129-225, 281-320` |
| ICMS DIFAL | Diferencial de Alíquotas (EC 87/2015) |   Implementado | `internal/legacy/icms.go:323-361` |
| PIS | Programa de Integração Social |   Implementado | `internal/legacy/pis_cofins.go` |
| COFINS | Contribuição para Financiamento da Seguridade Social |   Implementado | `internal/legacy/pis_cofins.go` |
| ISS | Imposto Sobre Serviços de Qualquer Natureza |   Implementado (2026-06-21) | `internal/legacy/iss.go` |
| FUST | Fundo de Universalização dos Serviços de Telecom |   Implementado (2026-06-21) | `internal/legacy/fust.go` |
| FUNTTEL | Fundo para o Desenvolvimento Tecnológico das Telecom |   Implementado (2026-06-21) | `internal/legacy/funttel.go` |
| CBS/IBS/IS | Reforma Tributária (IBS/CBS/IS) |   Implementado (2026-06-21) | `internal/reforma/reforma.go`, `data/init.sql:358-371` (schema) |

## Regras de Negócio Principais

### 1. Ordem de Execução dos Cálculos — Pipeline SOP-013 (C-001)

**Regra:** O motor executa os tributos em 7 fases porque há dependências assimétricas entre eles. A ordem é mandatória conforme SOP-013:

| Fase | Modo | Tributo(s) | Dependência |
|------|------|-----------|-------------|
| F0 | Sequencial | IS | Pré-filtro — NCM seletivo (BR-TAX-CONS-010) |
| F1 | Sequencial | IPI | Nenhuma (tributo extrafiscal). IPI_VALOR → ICMS |
| F2 | Sequencial | CBS | Nenhuma ("por fora", não compõe base de outros) |
| F3 | Sequencial | ICMS | IPI (base de cálculo). ICMS → PIS/COFINS (exclusão) |
| F4 | Paralela | IBS + ISS + PIS/COFINS | ICMS (VALOR_EXCLUSAO_ICMS), IPI |
| F5 | Sequencial | FUST | ICMS + PIS + COFINS (base líquida) |
| F6 | Sequencial | FUNTTEL | Mesma base que FUST |

**Dependências assimétricas que forçam esta ordem:**
- IPI → ICMS: IPI compõe a base do ICMS para consumidor final
- ICMS → PIS/COFINS: ICMS destacado é excluído da base (STF, "Tese do Século")
- ICMS + PIS + COFINS → FUST/FUNTTEL: base líquida após impostos principais
- IS → CBS: IS é pré-filtro obrigatório (BR-TAX-CONS-010)

**Injeção inter-fase:** Após cada fase, `injectTributoValues()` injeta os tributos calculados nos detalhes do input com chaves padronizadas. O ICMS é injetado com duas chaves: `ITEM_ICMS_VALOR` (FUST/FUNTTEL) e `VALOR_EXCLUSAO_ICMS` (PIS/COFINS).

**Fonte:** `internal/calculator/engine.go:1-330`

### 2. IPI — Imposto sobre Produtos Industrializados

**Regra de prioridade de configuração:**
1. Detalhes do item (overrides inline) — se `aliquota`, `cst` e `c_enq` estiverem preenchidos, pula o banco
2. Regra do banco (`ipi_regras`) — busca por NCM + CRT + tipo de operação + perfil comprador + UF + zona especial
3. Fallback: usa detalhes do item como valores inline

**Rateio de despesas:** Frete, seguro, outras despesas e desconto são rateados proporcionalmente entre os itens do documento.

**Modalidades de cálculo:**
- **Ad Valorem:** `valor_item × (aliquota / 100)`
- **Ad Pauta:** Valor fixo por unidade (definido na regra fiscal)

**Fonte:** `internal/legacy/ipi.go:1-124`

### 3. ICMS — Regime Normal (Fluxo Completo)

O fluxo foi reestruturado em 2026-06-20 para integrar `getEffectiveTaxConfig()` e implementar regras gerais de ST.

**Fluxo de decisão por operação:**

```
Calculate()
├── getEffectiveTaxConfig() → merge icms_rules + product_tax_exceptions
├── Se Simples Nacional → calcularICMSSimples()
└── Senão (Regime Normal):
    ├── Operação Interna (UF origem == UF destino)
    │   ├── CST 010 ou protocolo ST + MVA? → ICMS-ST
    │   │   └── Base ST = Valor × (1 + MVA/100), Aliq = AliquotaInterna
    │   └── Senão → ICMS Próprio (com redução de base e FCP)
    │       └── Base = Valor × (1 − ReducaoBase/100), Aliq = AliquotaInterna
    └── Operação Interestadual (UF origem != UF destino)
        ├── ICMS Interestadual (sempre calculado)
        │   └── Base = Valor × (1 − ReducaoBase/100), Aliq = AliquotaInterestadual
        ├── Se destino final (consumidor) → DIFAL
        │   └── DIFAL = Base × (AliquotaInterna − AliquotaInterestadual)
        └── Se protocolo ST + MVA → ICMS-ST
            └── Base ST = Valor × (1 + MVA/100), Aliq = AliquotaInterna
```

**Merge de configuração:** `getEffectiveTaxConfig()` em `icms.go:362-400` faz merge entre `icms_rules` e `product_tax_exceptions`, com override inteligente (campos não-zero da exceção sobrescrevem a regra base).

**Fonte:** `internal/legacy/icms.go:23-70, 72-127, 129-225, 227-321`

### 4. ICMS-ST — Substituição Tributária (Regras Gerais)

Implementado em 2026-06-20 com suporte a ST interno e interestadual sem dependência de exceção de produto.

**Regra ST Interno:** Acionado quando CST=010 OU `config.PossuiProtocoloST && MVA > 0`. MVA pode vir do config efetivo (regra geral) ou de override no item (CST 010).

**Regra ST Interestadual:** Acionado quando `config.PossuiProtocoloST && MVA > 0` (após ICMS interestadual). MVA via `getEffectiveTaxConfig()` — se exceção de produto tiver MVA, é injetada no config.

```
Base ST = Valor do Item × (1 + MVA/100)
ICMS-ST = Base ST × Aliquota Interna Destino
```

**Fonte:** `internal/legacy/icms.go:129-225, 281-320`

### 5. ICMS DIFAL — EC 87/2015

Atualizado em 2026-06-20 para receber `ICMSConfig` (mergeado) em vez de `*ICMSRule`.

**Regra:** Para operações interestaduais com consumidor final (não contribuinte), calcula-se o diferencial de alíquotas:

```
DIFAL = Base × (Aliquota Interna Destino − Aliquota Interestadual)
```

**Condições de aplicação:**
- Operação interestadual
- Consumidor final (não contribuinte)
- Alíquota interna do destino > alíquota interestadual
- Não aplicável a Simples Nacional

**Fonte:** `internal/legacy/icms.go:323-361`

### 6. ICMS — Simples Nacional

**Regra de equivalência CSOSN → CST:** Via tabela `tax_equivalence`, o código CSOSN é convertido para o CST equivalente.

**Cálculo da alíquota efetiva:**

```
Alíquota Efetiva = ((RBT12 × Aliquota Nominal) − Valor a Deduzir) / RBT12
ICMS = Alíquota Efetiva × Percentual ICMS no Anexo
```

Os parâmetros (`rbt12`, `aliquota nominal`, `valor a deduzir`, `perc_icms_anexo`) vêm da tabela `simples_nacional_rates` por faixa de RBT12.

**Fonte:** `internal/legacy/icms.go:107-120`

### 7. PIS/COFINS — Strategy por CST (com Integração ao Banco)

Migrado em 2026-06-20 para consultar alíquotas do banco (`federal_tax_rules`) com fallback para defaults.

**Resolução de alíquotas:**
```
Banco (GetFederalTaxRule) > Default hardcoded (PIS 1.65%, COFINS 7.6%)
```

**Exclusão do ICMS da base ("Tese do Século"):** Controlada pela flag `ExcluiICMSBase` da `federal_tax_rules`:
- `ExcluiICMSBase = true` → base = valor_item − ICMS (comportamento padrão, "Tese do Século")
- `ExcluiICMSBase = false` → base = valor_item (sem exclusão)

**Validação:** 13 cenários de teste em `pis_cofins_calculate_test.go` com mock repository cobrindo: ExcluiICMSBase true/false, fallback quando regra não encontrada, fallback em erro de banco, alíquotas customizadas, edge cases (base negativa, ICMS zero, CST vazio).

**Constantes de fallback:**
```go
const (
    defaultAliquotaPIS    = 1.65
    defaultAliquotaCOFINS = 7.6
)
```

**Estratégias por CST:**

| CST | Estratégia | Descrição | Cálculo |
|-----|-----------|-----------|---------|
| 01, 02 | `PIS01_02` / `COFINS01_02` | Tributado (Ad Valorem) | `(Valor Item × Aliquota) − Exclusão ICMS` |
| 03 | `PIS03` / `COFINS03` | Tributado (por unidade) | `Quantidade × Aliquota` via `CalcTax` |
| 04 | `PIS04` / `COFINS04` | Monofásico — tributo concentrado no produtor/importador | Valor zero |
| 05 | `PIS05` / `COFINS05` | Substituição Tributária — tributo já recolhido | Valor zero |
| 06 | `PIS06` / `COFINS06` | Alíquota Zero | Valor zero |
| 49 | `PIS49` / `COFINS49` | Outras Operações de Saída | Valor zero |
| 50-99 | `PIS50To99` / `COFINS50To99` | Operações de crédito, suspensão, regimes especiais | Valor zero |
| 99 | `PIS99` / COFINS50To99 | PIS: Outras operações. COFINS: CST 99 não possui estratégia dedicada — cai no fallback `COFINS50To99` (valor zero) | Valor zero |

**Observabilidade:** Campo `fonte_aliquota` nos `MoreTextDetails` indica se alíquota veio do banco ou do default. Log `slog.Warn` quando regra federal não encontrada.

**Fonte:** `internal/legacy/pis_cofins.go:1-196`, `internal/legacy/pis_strategies.go:1-98`, `internal/legacy/cofins_strategies.go:1-90`

### 8. Reforma Tributária — CBS, IBS, IS (Regime IVA Dual)

Implementado em 2026-06-21, refatorado em 2026-06-22 (C-001) com separação CBS/IBS para o pipeline SOP-013. A Reforma Tributária brasileira substitui PIS/COFINS/IPI (federais) e ICMS (estadual) pelo sistema de IVA Dual:

- **CBS** (Contribuição sobre Bens e Serviços) — esfera federal, unificada. Calculado pelo `CBSCalculator` (Fase 2 sequencial)
- **IBS** (Imposto sobre Bens e Serviços) — esfera estadual + municipal, soma das alíquotas. Calculado pelo `IBSCalculator` (Fase 4 paralela)
- **IS** (Imposto Seletivo) — "imposto do pecado". Calculado pelo `ISFilter` (Fase 0, pré-filtro)

**Separação CBS/IBS (C-001):** O `ReformaCalculator` original (legado) calculava CBS e IBS juntos. Para o pipeline SOP-013, foi dividido em `CBSCalculator` e `IBSCalculator` — CBS executa sequencialmente na Fase 2 ("por fora", antes do ICMS), IBS executa em paralelo na Fase 4. Ambos compartilham a função interna `computeIvaDual()` que consulta `GetIvaDualRule` — o cache Redis evita dupla consulta ao banco.

**Regra de lookup:** `GetIvaDualRule` prioriza regra específica do município sobre a regra padrão estadual (`ORDER BY municipio_destino_ibge DESC`).

**Redução de alíquotas (transição):**
```
fatorReducao = 1 − (percentualReducao / 100)
aliquotaEfetiva = aliquotaNominal × fatorReducao
```

**Cálculo:**
```
CBS = valor_item × (aliquotaCBS × fatorReducao / 100)
IBS = valor_item × ((aliquotaIBSEstadual + aliquotaIBSMunicipal) × fatorReducao / 100)
IS  = valor_item × (aliquotaIS / 100)  [ISFilter, Fase 0 — não sofre redução]
```

**CSTs:** Valores provisórios `01` (normal) e `04` (com redução) — aguardando tabela oficial da RFB.

**Cache:** Redis com chave `tax:iva:<ncm>:<uf>:<municipio>` (TTL 24h).

**Fonte:** `internal/reforma/reforma.go:1-178`, `internal/reforma/cbs_calculator.go:1-64`, `internal/reforma/ibs_calculator.go:1-65`, `repository/postgres_repository.go:359-401`

### 9. ISS — Imposto Sobre Serviços de Qualquer Natureza

**Regra de incidência:** Apenas itens classificados como serviço (detalhe `ITEM_LISTA_SERVICO` preenchido). Mercadorias não pagam ISS.

**Validação de alíquota municipal (BR-TAX-CONS-007):**
```
2% ≤ Aliquota_ISS ≤ 5%
```
Se a alíquota estiver fora do range, o cálculo prossegue com `slog.Warn` — a correção é responsabilidade do Tax Compliance Officer.

**Cálculo (BR-TAX-CALC-016):**
```
ISS = Preço_Serviço × Aliquota_Municipal / 100
```

**Retenção na fonte:** Se a flag `ISS_RETIDO` estiver presente, o ISS foi retido pelo tomador e é apenas informado, não recolhido.

**Classificação de serviços (BR-TAX-DEF-007):** Serviços de telecomunicação = item 1.05 da Lista de Serviços anexa à LC 116/2003.

**Fonte:** `internal/legacy/iss.go:1-140`

### 10. FUST e FUNTTEL — Contribuições Setoriais de Telecom

**Regra de classificação (BR-TAX-INF-007):** Apenas SCM (Serviço de Comunicação Multimídia) e STFC (Serviço Telefônico Fixo Comutado) pagam FUST e FUNTTEL. SVA (Serviço de Valor Adicionado) é isento.

**FUST (BR-TAX-CALC-019) — Lei 9.998/2000:**
```
Base_FUST = Valor_Serviço − ICMS − PIS − COFINS
FUST = Base_FUST × 1%
Se Base_FUST < 0 → FUST = 0
```

**FUNTTEL (BR-TAX-CALC-020) — Lei 10.052/2000:**
```
Base_FUNTTEL = Base_FUST (mesma base)
FUNTTEL = Base_FUNTTEL × 0,5%
```

**Dependência de pipeline:** FUST e FUNTTEL são executados na Fase 3 (pós-paralela) porque dependem dos valores de ICMS, PIS e COFINS já calculados na Fase 2. A engine injeta `ITEM_ICMS_VALOR`, `ITEM_PIS_VALOR` e `ITEM_COFINS_VALOR` nos detalhes do input antes da Fase 3.

**Fonte:** `internal/legacy/fust.go:1-120`, `internal/legacy/funttel.go:1-100`, `internal/legacy/telecom.go:1-90`

### 11. Motor de Cálculo Multi-Fase (SOP-013 / C-001)

**Regra:** O motor executa tributos em 7 fases porque as dependências entre eles são assimétricas. A arquitetura usa `CalculationPhase` genérica com modos `Sequential` e `Parallel`:

1. **Fase 0 (Sequencial):** IS — pré-filtro obrigatório (NCM seletivo), não sofre redução
2. **Fase 1 (Sequencial):** IPI — compõe base do ICMS para consumidor final
3. **Fase 2 (Sequencial):** CBS — "por fora", não compõe base de outros
4. **Fase 3 (Sequencial):** ICMS — sequencial ANTES do PIS/COFINS (Tese do Século)
5. **Fase 4 (Paralela):** IBS + ISS + PIS/COFINS — independentes entre si
6. **Fase 5 (Sequencial):** FUST — depende de ICMS + PIS + COFINS
7. **Fase 6 (Sequencial):** FUNTTEL — mesma base que FUST

**Construtores:**
- `BillingEnginePhased(phases ...)` — principal (C-001), fases arbitrárias
- `BillingEngineFull(pre, calcs, post)` — compatibilidade retroativa (3 fases)
- `BillingEngineOrdered(pre, calcs ...)` — compatibilidade retroativa (2 fases)
- `BillingEngine(calcs ...)` — compatibilidade retroativa (1 fase paralela)

**Injeção inter-fase:** `injectTributoValues()` injeta automaticamente os valores calculados nos detalhes do input após cada fase. Chaves duplas para ICMS: `ITEM_ICMS_VALOR` + `VALOR_EXCLUSAO_ICMS`.

**Propagação de erros:** Fases sequenciais propagam erro (interrompem pipeline). Fases paralelas coletam e logam erros (pipeline continua).

**Fonte:** `internal/calculator/engine.go:1-350`, `cmd/api/main.go:86-140`

### 12. ICMS Desonerado — Redução de Base e Limitação de Alíquota (F-004)

Implementado em 2026-06-21 conforme SOP-017, BR-TAX-CALC-021/022.

**Regra de validação CST (BR-TAX-CONS-013):** Apenas CSTs {20, 30, 40, 41, 50, 70, 90} permitem desoneração. CST 00 (tributação integral) NÃO permite. Simples Nacional é excluído (usa regime próprio CSOSN).

**Motivos de desoneração SEFAZ (motDesICMS):** 13 códigos oficiais (1-12, 90) conforme Ajuste SINIEF 07/2005. Default = 9 (Outros).

**Modo 1 — Redução de Base (BR-TAX-CALC-021):**
```
Base_Reduzida = Valor × (1 − PctRedução/100)
ICMS = Base_Reduzida × Alíquota/100
vICMSDeson = (Valor × Alíquota) − ICMS
```

**Modo 2 — Limitação de Alíquota (BR-TAX-CALC-022):**
```
Índice = 1 − (AliqAlvo/AliqNominal)
Base_Reduzida = Valor × (AliqAlvo/AliqNominal)
ICMS = Base_Reduzida × Alíquota/100
vICMSDeson = (Valor × Alíquota) − ICMS
```

**Prioridade de modo:** Se `aliquota_alvo > 0` → Limitação. Senão, se `percentual_reducao > 0` → Redução de Base.

**Integração FCP:** O FCP é calculado sobre a base reduzida quando presente.

**Fonte:** `internal/legacy/icms_desoneracao.go:1-310`, `internal/legacy/icms.go:177-228` (integração via `calcularICMSOperacaoInterna`)

### 13. Phase Resolution System — Fases da Reforma Tributária (F-005)

Implementado em 2026-06-21 conforme BR-TAX-INF-001 a 004, DT-001, BR-TAX-ACT-005/006.

**Resolução de fase (PhaseResolver.Resolve):**
```
DataOperacao.Year() ∈ {2026}            → SHADOW_RUN
DataOperacao.Year() ∈ {2027, 2028}      → CBS_PLENA
DataOperacao.Year() ∈ {2029, 2030, 2031, 2032} → TRANSICAO_SUBNACIONAL
DataOperacao.Year() ≥ 2033              → IVA_DUAL
```

**Características por fase (DT-001):**

| Fase | CBS | IBS | PIS/COFINS | ICMS | ISS | IPI |
|------|-----|-----|-----------|------|-----|-----|
| SHADOW_RUN (2026) | Shadow | Shadow | Ativo | Ativo | Ativo | Ativo |
| CBS_PLENA (2027-28) | Ativo | Shadow | Extinto | Ativo | Ativo | Ativo |
| TRANSICAO (2029-32) | Ativo | Ativo | Extinto | Ativo* | Ativo* | Ativo |
| IVA_DUAL (2033+) | Ativo | Ativo | Extinto | Extinto | Extinto | Ativo |

*Com fator de redução progressivo: 2029=25%, 2030=50%, 2031=75%, 2032=100%.

**Shadow Tax (BR-TAX-ACT-005):** Na fase Shadow Run (2026), CBS e IBS são calculados e registrados nos detalhes, mas NÃO compõem `total_impostos`. O log `slog.Info` registra `shadow_total` separadamente.

**Extinção IVA Dual (BR-TAX-ACT-006):** A partir de 2033, tributos legados (PIS, COFINS, ICMS, ISS) são zerados (BaseCalculo=0, Aliquota=0, Valor=0) com tag `IVA_DUAL_EXTINTO` nos text details.

**Integração ao motor:** `ProcessWithPhase()` em `engine.go` recebe `CalculatorFilter` e aplica shadow tax exclusion, subnational reduction e legacy tax extinction no pós-processamento.

**TaxSelector:** Consulta `ShouldIncludeInTotal(dataOperacao, tributo)` para determinar se um tributo compõe o total a pagar. Aplica matriz DT-001 completa.

**Fonte:** `internal/phase/phase.go:1-198`, `internal/phase/tax_selector.go:1-160`, `internal/calculator/engine.go:79-205` (ProcessWithPhase), `cmd/api/main.go:48-51` (wiring)

## Glossário de Domínio

| Termo | Sigla | Definição |
|-------|------|-----------|
| NCM | Nomenclatura Comum do Mercosul | Código de 8 dígitos que classifica mercadorias |
| CFOP | Código Fiscal de Operações e Prestações | Classifica o tipo de operação fiscal |
| CST | Código de Situação Tributária | Indica a tributação pelo ICMS (origem) |
| CSOSN | Código de Situação da Operação no Simples Nacional | Equivalente ao CST para optantes do Simples |
| CRT | Código de Regime Tributário | 1=Simples Nacional, 3=Normal |
| MVA | Margem de Valor Agregado | Percentual aplicado na base de cálculo do ICMS-ST |
| DIFAL | Diferencial de Alíquotas | Diferença entre alíquota interna e interestadual |
| FCP | Fundo de Combate à Pobreza | Adicional de até 2% na alíquota do ICMS |
| RBT12 | Receita Bruta Total 12 meses | Base para definição da faixa do Simples Nacional |
| TIPI | Tabela de Incidência do IPI | Classificação fiscal para IPI (usa NCM + EX) |
| EX | Exceção na TIPI | Código adicional ao NCM para classificação do IPI |
| ISS | Imposto Sobre Serviços de Qualquer Natureza | Tributo municipal sobre prestação de serviços (LC 116/2003) |
| SCM | Serviço de Comunicação Multimídia | Banda larga, dados — sujeito a FUST/FUNTTEL |
| STFC | Serviço Telefônico Fixo Comutado | Telefonia fixa — sujeito a FUST/FUNTTEL |
| SVA | Serviço de Valor Adicionado | Streaming, antivírus, suporte técnico — isento de FUST/FUNTTEL |
| FUST | Fundo de Universalização dos Serviços de Telecomunicações | Contribuição setorial de 1% (Lei 9.998/2000) |
| FUNTTEL | Fundo para o Desenvolvimento Tecnológico das Telecomunicações | Contribuição setorial de 0,5% (Lei 10.052/2000) |
| motDesICMS | Motivo da Desoneração do ICMS | Código SEFAZ (1-12, 90) que classifica o tipo de benefício fiscal aplicado (Ajuste SINIEF 07/2005) |
| vICMSDeson | Valor do ICMS Desonerado | Diferença entre o ICMS que seria devido sem desoneração e o ICMS efetivamente calculado |
| Shadow Run | Fase de Sombra da Reforma Tributária | CBS/IBS calculados em 2026 mas não compõem total a pagar (testes e homologação) |
| IVA Dual | Sistema Dual de IVA | CBS (federal) + IBS (subnacional) em vigor a partir de 2033, substituindo 5 tributos legados |
| DT-001 | Matriz de Transição Fiscal | Tabela que define quais tributos estão ativos por fase da reforma tributária |
