# Catálogo de Regras de Negócio: Cálculo de Impostos Corporativos

**Código:** RULES-CATALOG-FIN-00001  
**Versão:** 1.1  
**Política vinculada:** [POLICE-FIN-00001](./POLICE-FIN-00001-CALCULO-DE-IMPOSTOS-CORPORATIVOS.md)  
**Domínio:** Tributário — Telecomunicações  
**Última Atualização:** 2026-06-21  
**Owner:** Tax Compliance Officer + Engineering Lead  
**Cadência de Revisão:** Semestral, alinhada à política-mãe  

---

## Glossário de Termos de Negócio (SBVR)

| Termo | Definição |
|:---|:---|
| **Lucro Real** | Regime tributário em que a base de cálculo do IRPJ/CSLL é o lucro contábil ajustado por adições e exclusões do LALUR, com compensação de prejuízos fiscais limitada a 30% do lucro líquido do período. |
| **Lucro Presumido** | Regime tributário em que a base de cálculo do IRPJ/CSLL é obtida pela aplicação de percentuais de presunção legal (8% para comércio/indústria; 32% para serviços) sobre a receita bruta. |
| **Regime Cumulativo (PIS/COFINS)** | Modalidade em que a alíquota total é 3,65% sobre a receita bruta, sem direito a crédito sobre insumos. |
| **Regime Não-Cumulativo (PIS/COFINS)** | Modalidade em que a alíquota total é 9,25% sobre a receita bruta, com direito a crédito sobre insumos diretamente relacionados à atividade-fim. Empresas de TELECOM são obrigadas a este regime. |
| **Cálculo "por fora"** | Método em que o imposto não integra sua própria base de cálculo. Aplicável à CBS na Reforma Tributária. |
| **Princípio do Destino** | Regra do IBS em que a alíquota aplicável é a do município/estado onde o bem ou serviço é consumido, não onde é produzido. |
| **Não-Cumulatividade Plena** | Modelo da CBS em que todo insumo gera crédito financeiro, independentemente de pertinência ao objeto social da empresa. |
| **DIFAL** | Diferencial de Alíquota do ICMS: diferença entre a alíquota interna do estado de destino e a alíquota interestadual, aplicável em vendas interestaduais para consumidor final. |
| **ICMS-ST** | Substituição Tributária do ICMS: imposto recolhido antecipadamente pelo remetente (indústria/importador) em nome de toda a cadeia de comercialização. |
| **FCP** | Fundo de Combate à Pobreza: adicional de até 2% sobre a base de cálculo do ICMS em operações específicas, conforme legislação estadual. |
| **Shadow Run** | Execução paralela de cálculo (não compõe o valor a pagar) para validação e calibração do novo regime tributário antes da entrada em produção. |
| **Hybrid Mode** | Período de coexistência entre o regime Legacy (ICMS/ISS/PIS/COFINS) e o regime da Reforma (CBS/IBS) durante a transição 2026–2033. |
| **PER/DCOMP** | Pedido Eletrônico de Restituição / Declaração de Compensação: mecanismo para compensar créditos excedentes de um tributo federal com débitos de outro. |
| **RBT12** | Receita Bruta Acumulada dos últimos 12 meses, utilizada como parâmetro para determinação da alíquota efetiva no Simples Nacional. |
| **MVA** | Margem de Valor Agregado: percentual de lucro presumido pelo fisco para compor a base de cálculo do ICMS-ST. |
| **NCM** | Nomenclatura Comum do Mercosul: código de 8 dígitos que classifica mercadorias e determina alíquotas de IPI e incidência do IS. |
| **CEST** | Código Especificador da Substituição Tributária: identifica produtos sujeitos ao regime de ICMS-ST. |
| **CFOP** | Código Fiscal de Operações e Prestações: define a natureza jurídica da operação (venda, devolução, transferência, etc.). |
| **CST** | Código de Situação Tributária: identifica a tributação aplicável a um item (tributado, isento, com crédito, etc.). |

---

## Sumário de Regras

| Categoria | Código | Quantidade |
|:---|:---|:---|
| Structural Rules (Definições e Fatos) | `BR-TAX-DEF` | 11 |
| Derivation Rules (Cálculos e Fórmulas) | `BR-TAX-CALC` | 22 |
| Constraint Rules (Restrições e Condições) | `BR-TAX-CONS` | 13 |
| Inference Rules (Classificações e Deduções) | `BR-TAX-INF` | 7 |
| Action Rules (Disparadores e Workflows) | `BR-TAX-ACT` | 7 |
| **Total** | | **60** |

---

# PARTE I — STRUCTURAL RULES (Definições e Fatos)

Regras que estabelecem a terminologia oficial do domínio tributário e os relacionamentos entre conceitos.

---

### BR-TAX-DEF-001 — Definição de Lucro Real

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o regime de Lucro Real seja definido como o regime em que a base de cálculo do IRPJ e da CSLL é o lucro contábil ajustado por adições e exclusões do LALUR, com compensação de prejuízos fiscais limitada a 30% do lucro líquido do período._

**Fonte:** POLICE-FIN-00001 §4.3 (IRPJ/CSLL)  
**Autoridade para alteração:** Tax Compliance Officer  
**Volatilidade:** Estável  
**Aplicação:** Motor de Período — Engine de Lucratividade  

---

### BR-TAX-DEF-002 — Definição de Lucro Presumido

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o regime de Lucro Presumido seja definido como o regime em que a base de cálculo é a receita bruta multiplicada pelo percentual de presunção legal de 8% para atividades de comércio e indústria, e 32% para atividades de serviços._

**Fonte:** POLICE-FIN-00001 §4.3 (IRPJ/CSLL)  
**Autoridade para alteração:** Tax Compliance Officer  
**Volatilidade:** Estável  
**Aplicação:** Motor de Período — Engine de Lucratividade  

---

### BR-TAX-DEF-003 — Definição de Regime Cumulativo (PIS/COFINS)

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o Regime Cumulativo de PIS/COFINS seja definido como a modalidade em que a alíquota total é 3,65% (0,65% PIS + 3% COFINS) sobre a receita bruta, sem direito a crédito sobre insumos._

**Fonte:** POLICE-FIN-00001 §4.3 (PIS/COFINS); Lei 9.718/1998  
**Autoridade para alteração:** Tax Compliance Officer  
**Volatilidade:** Volátil (será extinto em 2027)  
**Aplicação:** Motor de Faturamento — Engine Legacy  

---

### BR-TAX-DEF-004 — Definição de Regime Não-Cumulativo (PIS/COFINS)

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o Regime Não-Cumulativo de PIS/COFINS seja definido como a modalidade em que a alíquota total é 9,25% (1,65% PIS + 7,6% COFINS) sobre a receita bruta, com direito a crédito sobre insumos diretamente relacionados à atividade-fim._

**Fonte:** POLICE-FIN-00001 §4.3 (PIS/COFINS); Lei 10.833/2003  
**Autoridade para alteração:** Tax Compliance Officer  
**Volatilidade:** Volátil (será extinto em 2027)  
**Aplicação:** Motor de Faturamento — Engine Legacy  

---

### BR-TAX-DEF-005 — Definição de Cálculo "por fora"

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o cálculo "por fora" seja definido como o método em que o imposto não integra sua própria base de cálculo, sendo aplicado como um percentual sobre a base líquida._

**Fonte:** POLICE-FIN-00001 §4.3 (CBS)  
**Autoridade para alteração:** Tax Compliance Officer  
**Volatilidade:** Estável  
**Aplicação:** Motor de Faturamento — Engine Reforma (CBS)  

**Exemplo:**
- Positivo: Base = R$ 1.000,00; CBS 10% → CBS = R$ 100,00; Total = R$ 1.100,00
- Negativo (se fosse "por dentro"): Base = R$ 1.000,00; CBS 10% → Base recalculada = R$ 1.000 / 0,90 = R$ 1.111,11

---

### BR-TAX-DEF-006 — Princípio do Destino (IBS)

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o Princípio do Destino do IBS seja definido como a regra em que a alíquota aplicável é a do município e estado onde o bem ou serviço é consumido pelo adquirente final, independentemente do local de produção ou prestação._

**Fonte:** POLICE-FIN-00001 §4.3 (IBS); EC 132/2023  
**Autoridade para alteração:** Comitê Gestor do IBS (externo)  
**Volatilidade:** Estável (constitucional)  
**Aplicação:** Motor de Faturamento — Engine Reforma (IBS)  

---

### BR-TAX-DEF-007 — Classificação de Serviços de Telecomunicação no ISS

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que serviços de telecomunicação sejam classificados no item 1.05 da Lista de Serviços anexa à Lei Complementar 116/2003._

**Fonte:** POLICE-FIN-00001 §4.3 (ISS); LC 116/2003  
**Autoridade para alteração:** Congresso Nacional  
**Volatilidade:** Estável  
**Aplicação:** Motor de Faturamento — Engine Legacy (ISS)  

---

### BR-TAX-DEF-008 — Definição de FCP

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o FCP (Fundo de Combate à Pobreza) seja definido como o adicional de até 2% calculado sobre a base de ICMS em operações específicas, conforme legislação de cada estado._

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS)  
**Autoridade para alteração:** Assembleias Legislativas Estaduais  
**Volatilidade:** Moderada  
**Aplicação:** Motor de Faturamento — Engine Legacy (ICMS)  

---

### BR-TAX-DEF-009 — Definição de ICMS-ST

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o ICMS-ST (Substituição Tributária) seja definido como o regime em que o imposto é recolhido antecipadamente pelo remetente (indústria ou importador) em nome de todos os contribuintes da cadeia de comercialização subsequente._

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS); Lei Complementar 87/1996  
**Autoridade para alteração:** CONFAZ  
**Volatilidade:** Moderada  
**Aplicação:** Motor de Faturamento — Engine Legacy (ICMS)  

---

### BR-TAX-DEF-010 — Definição de FUST

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o FUST (Fundo de Universalização dos Serviços de Telecomunicações) seja definido como a contribuição setorial de 1% incidente sobre a Receita Operacional Líquida (Valor Bruto − ICMS − PIS − COFINS) de serviços de telecomunicação (SCM, STFC), excluídos os Serviços de Valor Adicionado (SVA)._

**Fonte:** POLICE-FIN-00001 §4.3 (FUST/FUNTTEL); Lei 9.998/2000  
**Autoridade para alteração:** Congresso Nacional  
**Volatilidade:** Estável  
**Aplicação:** Motor de Faturamento — Engine Telecom  

---

### BR-TAX-DEF-011 — Definição de FUNTTEL

**Categoria:** Structural  
**Declaração (SBVR):** _É obrigatório que o FUNTTEL (Fundo para o Desenvolvimento Tecnológico das Telecomunicações) seja definido como a contribuição setorial de 0,5% incidente sobre a mesma base do FUST (Receita Operacional Líquida), aplicável exclusivamente a serviços de telecomunicação (SCM, STFC)._

**Fonte:** POLICE-FIN-00001 §4.3 (FUST/FUNTTEL); Lei 10.052/2000  
**Autoridade para alteração:** Congresso Nacional  
**Volatilidade:** Estável  
**Aplicação:** Motor de Faturamento — Engine Telecom  

---

# PARTE II — DERIVATION RULES (Cálculos e Fórmulas)

Regras matemáticas que o motor de cálculo deve implementar. Cada regra é diretamente traduzível para uma função em Go.

---

### BR-TAX-CALC-001 — Base de Cálculo IRPJ/CSLL (Lucro Real)

**Categoria:** Derivation  
**Declaração (SBVR):** _A Base de Cálculo do IRPJ e da CSLL no Lucro Real deve ser calculada como: Lucro Contábil + Adições LALUR − Exclusões LALUR − Compensação de Prejuízos Fiscais._

**Fórmula:**  
```
Base_IRPJ_CSLL = Lucro_Contabil + SUM(Adicoes_LALUR) - SUM(Exclusoes_LALUR) - MIN(Compensacao_Prejuizos, Lucro_Liquido_Ajustado × 0,30)
```

**Fonte:** POLICE-FIN-00001 §4.3 (IRPJ/CSLL)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestIRPJ_BaseCalculo_LucroReal_ComAjustesLALUR`  

---

### BR-TAX-CALC-002 — Base de Cálculo IRPJ/CSLL (Lucro Presumido)

**Categoria:** Derivation  
**Declaração (SBVR):** _A Base de Cálculo do IRPJ e da CSLL no Lucro Presumido deve ser calculada como: Receita Bruta × Percentual de Presunção._

**Fórmula:**  
```
Base_Presumida = Receita_Bruta × Percentual_Presuncao
ONDE:
  Percentual_Presuncao = 8%   SE CNAE ∈ COMERCIO || INDUSTRIA
  Percentual_Presuncao = 32%  SE CNAE ∈ SERVICOS
```

**Fonte:** POLICE-FIN-00001 §4.3 (IRPJ/CSLL)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestIRPJ_BaseCalculo_LucroPresumido_Servico_32pct`  

---

### BR-TAX-CALC-003 — Valor IRPJ

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do IRPJ deve ser calculado como: Base de Cálculo × 15%._

**Fórmula:**  
```
IRPJ_Valor = Base_IRPJ × 0,15
```

**Fonte:** POLICE-FIN-00001 §4.1  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestIRPJ_Valor_Aliquota15pct`  

---

### BR-TAX-CALC-004 — Adicional de IRPJ

**Categoria:** Derivation  
**Declaração (SBVR):** _O adicional de IRPJ de 10% deve ser calculado sobre a parcela da base de cálculo que exceder R$ 20.000,00 multiplicado pelo número de meses do período de apuração._

**Fórmula:**  
```
IRPJ_Adicional = MAX(0, Base_IRPJ - 20000 × n_meses) × 0,10
```

**Fonte:** POLICE-FIN-00001 §4.3 (IRPJ/CSLL); Lei 9.249/1995 art. 3º  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestIRPJ_Adicional_ExcedeLimiteMensal`  

**Exemplo:**
- Trimestral (n=3): Base = R$ 200.000,00 → Limite = R$ 60.000,00 → Excesso = R$ 140.000,00 → Adicional = R$ 14.000,00

---

### BR-TAX-CALC-005 — Valor CSLL

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor da CSLL deve ser calculado como: Base de Cálculo × 9%._

**Fórmula:**  
```
CSLL_Valor = Base_CSLL × 0,09
```

**Fonte:** POLICE-FIN-00001 §4.1  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestCSLL_Valor_Aliquota9pct`  

**Exceção:** Instituições financeiras e equiparadas: alíquota de 20% (Lei 13.169/2015), não aplicável a TELECOM.

---

### BR-TAX-CALC-006 — Valor PIS Cumulativo

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do PIS no regime cumulativo deve ser calculado como: Receita Bruta × 0,65%._

**Fórmula:**  
```
PIS_Cumulativo = Receita_Bruta × 0,0065
```

**Fonte:** POLICE-FIN-00001 §4.1; Lei 9.718/1998  
**Prioridade:** Média (extinto em 2027)  
**Enforcement:** Automático  
**Testabilidade:** `TestPIS_Cumulativo_Aliquota065pct`  

---

### BR-TAX-CALC-007 — Valor PIS Não-Cumulativo

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do PIS no regime não-cumulativo deve ser calculado como: Base de Créditos × 1,65%._

**Fórmula:**  
```
PIS_NaoCumulativo = (Receita_Bruta - Creditos_Entrada) × 0,0165
```

**Fonte:** POLICE-FIN-00001 §4.1; Lei 10.833/2003  
**Prioridade:** Alta (TELECOM é obrigada a este regime)  
**Enforcement:** Automático  
**Testabilidade:** `TestPIS_NaoCumulativo_ComCreditos`  

---

### BR-TAX-CALC-008 — Valor COFINS Cumulativo

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor da COFINS no regime cumulativo deve ser calculado como: Receita Bruta × 3%._

**Fórmula:**  
```
COFINS_Cumulativo = Receita_Bruta × 0,03
```

**Fonte:** POLICE-FIN-00001 §4.1; Lei 9.718/1998  
**Prioridade:** Média (extinto em 2027)  
**Enforcement:** Automático  
**Testabilidade:** `TestCOFINS_Cumulativo_Aliquota3pct`  

---

### BR-TAX-CALC-009 — Valor COFINS Não-Cumulativo

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor da COFINS no regime não-cumulativo deve ser calculado como: Base de Créditos × 7,6%._

**Fórmula:**  
```
COFINS_NaoCumulativo = (Receita_Bruta - Creditos_Entrada) × 0,076
```

**Fonte:** POLICE-FIN-00001 §4.1; Lei 10.833/2003  
**Prioridade:** Alta (TELECOM é obrigada a este regime)  
**Enforcement:** Automático  
**Testabilidade:** `TestCOFINS_NaoCumulativo_ComCreditos`  

---

### BR-TAX-CALC-010 — Base de Cálculo ICMS

**Categoria:** Derivation  
**Declaração (SBVR):** _A base de cálculo do ICMS deve ser composta por: Valor da Operação + Frete + Seguro + Outras Despesas Acessórias + IPI (quando o adquirente for consumidor final)._

**Fórmula:**  
```
Base_ICMS = Valor_Operacao + Frete + Seguro + Outras_Despesas + IPI (se ConsumidorFinal)
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS)  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestICMS_BaseCalculo_CompoeFreteSeguroIPI`  

---

### BR-TAX-CALC-011 — Valor ICMS Próprio

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do ICMS próprio deve ser calculado como: Base de Cálculo × Alíquota aplicável._

**Fórmula:**  
```
ICMS_Proprio = Base_ICMS × Aliquota_ICMS
```

**Fonte:** POLICE-FIN-00001 §4.1  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestICMS_Proprio_CalculoPadrao`  

---

### BR-TAX-CALC-012 — DIFAL (Diferencial de Alíquota do ICMS)

**Categoria:** Derivation  
**Declaração (SBVR):** _O DIFAL deve ser calculado como a diferença entre o ICMS calculado com a alíquota interna do estado de destino e o ICMS calculado com a alíquota interestadual._

**Fórmula:**  
```
DIFAL = (Base_ICMS × AliqInternaDestino) - (Base_ICMS × AliqInterestadual)
```

**Condições de Aplicação:**
1. `UF_Origem ≠ UF_Destino` (operação interestadual)
2. `Consumidor_Final = true` (não é revenda)
3. `AliqInternaDestino > AliqInterestadual` (há diferença a recolher)

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS); EC 87/2015  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestDIFAL_SP_Para_BA_ConsumidorFinal`  

**Exemplo:**
- Venda SP → BA (consumidor final): Base = R$ 1.000,00
- ICMS Origem = R$ 1.000 × 7% = R$ 70,00
- ICMS Destino = R$ 1.000 × 20,5% = R$ 205,00
- DIFAL = R$ 205,00 − R$ 70,00 = **R$ 135,00** (devidos à BA)

---

### BR-TAX-CALC-013 — Base de Cálculo ICMS-ST

**Categoria:** Derivation  
**Declaração (SBVR):** _A base de cálculo do ICMS-ST deve ser calculada como: Valor da Operação × (1 + MVA)._

**Fórmula:**  
```
Base_ICMS_ST = Valor_Operacao × (1 + MVA/100)
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestICMS_ST_BaseCalculo_ComMVA`  

---

### BR-TAX-CALC-014 — Valor ICMS-ST

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do ICMS-ST a recolher deve ser calculado como: (Base ST × Alíquota Interna do Destino) − ICMS Próprio._

**Fórmula:**  
```
ICMS_ST_Valor = (Base_ICMS_ST × AliqInternaDestino) - ICMS_Proprio
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestICMS_ST_Valor_STMenosProprio`  

**Exemplo:**
- Valor = R$ 1.000,00; MVA = 40%; AliqInterna = 18%; AliqInterestadual = 12%
- Base ST = R$ 1.000 × 1,40 = R$ 1.400,00
- ICMS Próprio = R$ 1.000 × 12% = R$ 120,00
- ICMS ST = R$ 1.400 × 18% − R$ 120,00 = R$ 252,00 − R$ 120,00 = **R$ 132,00**

---

### BR-TAX-CALC-021 — ICMS Desonerado — Redução de Base de Cálculo

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do ICMS com redução de base de cálculo deve ser calculado aplicando o percentual de redução sobre o valor do item para obter a base reduzida, e então aplicando a alíquota nominal sobre esta base._

**Fórmula:**  
```
Base_Reduzida = Valor_Item × (1 − Percentual_Reducao / 100)
ICMS_Desonerado = Base_Reduzida × Aliquota_ICMS / 100
vICMSDeson = (Valor_Item × Aliquota_ICMS / 100) − ICMS_Desonerado
Valor_Final_Item = Valor_Item − vICMSDeson
```

**Condições de Aplicação:**
1. `CST_ICMS ∈ {20, 30, 40, 41, 50, 70, 90}` — apenas CSTs que permitem desoneração
2. `motDesICMS ∈ {1..12, 90}` — código de motivo válido

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS Desonerado)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestICMS_Desonerado_ReducaoBase`  

**Exemplo:**
- Valor = R$ 1.000,00; Alíquota = 18%; Redução = 20%
- Base Reduzida = 1.000 × (1 − 0,20) = R$ 800,00
- ICMS = 800 × 18% = R$ 144,00
- vICMSDeson = (1.000 × 18%) − 144 = 180 − 144 = **R$ 36,00**
- Valor Final = 1.000 − 36 = **R$ 964,00**

---

### BR-TAX-CALC-022 — ICMS Desonerado — Limitação de Alíquota Efetiva

**Categoria:** Derivation  
**Declaração (SBVR):** _Quando a legislação define uma carga tributária efetiva máxima (alíquota alvo), o índice de redução deve ser calculado como 1 − (Alíquota_Alvo / Alíquota_Nominal), e a base reduzida como Valor × (Alíquota_Alvo / Alíquota_Nominal)._

**Fórmula:**  
```
Indice_Reducao = 1 − (Aliquota_Alvo / Aliquota_Nominal)
Base_Reduzida = Valor_Item × (Aliquota_Alvo / Aliquota_Nominal)
ICMS_Desonerado = Base_Reduzida × Aliquota_Nominal / 100
vICMSDeson = (Valor_Item × Aliquota_Nominal / 100) − ICMS_Desonerado
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS Desonerado)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestICMS_Desonerado_LimitacaoAliquota`  

**Exemplo:**
- Valor = R$ 1.000,00; Alíquota Nominal = 18%; Alíquota Alvo = 8%
- Índice Redução = 1 − (8/18) = 55,56%
- Base Reduzida = 1.000 × (8/18) = R$ 444,44
- ICMS = 444,44 × 18% = R$ 80,00
- vICMSDeson = 180 − 80 = **R$ 100,00**

---

### BR-TAX-CALC-015 — Valor IPI

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do IPI deve ser calculado como: Valor do Produto × Alíquota TIPI determinada pelo NCM do item._

**Fórmula:**  
```
IPI_Valor = Valor_Produto × Aliquota_TIPI(NCM)
```

**Fonte:** POLICE-FIN-00001 §4.3 (IPI)  
**Prioridade:** Média  
**Enforcement:** Automático  
**Testabilidade:** `TestIPI_Valor_PorNCM`  

**Exceção:** IPI não incide sobre serviços de telecomunicação puros. Incide apenas sobre fornecimento de equipamentos e produtos industrializados.

---

### BR-TAX-CALC-016 — Valor ISS

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do ISS deve ser calculado como: Preço do Serviço × Alíquota Municipal (2% a 5%)._

**Fórmula:**  
```
ISS_Valor = Preco_Servico × Aliquota_Municipal
ONDE:
  Aliquota_Municipal ∈ [2%, 5%]
```

**Fonte:** POLICE-FIN-00001 §4.3 (ISS); LC 116/2003  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestISS_Valor_AliquotaMunicipal`  

---

### BR-TAX-CALC-017 — Valor CBS

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor da CBS deve ser calculado como: Base × Alíquota CBS, utilizando o método "por fora" (a CBS não integra sua própria base de cálculo)._

**Fórmula:**  
```
CBS_Valor = Base × Aliquota_CBS
ONDE:
  Aliquota_CBS = 0,001  SE Fase = "Shadow Run" (2026)
  Aliquota_CBS = DEFINIDA_PELO_GOVERNO  SE Fase ≥ "CBS Plena" (2027+)
  // NOTA: Alíquota setorial TELECOM estimada em 10,5% a 12%
```

**Fonte:** POLICE-FIN-00001 §4.3 (CBS); EC 132/2023  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Status:** ⚠️ Alíquota plena aguardando definição do Governo Federal  
**Testabilidade:** `TestCBS_NaoCompoePropriaBase`  

---

### BR-TAX-CALC-018 — Valor IBS

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do IBS deve ser calculado como: Base × Alíquota do Destino (soma da parcela estadual + municipal), obtida via API do Comitê Gestor._

**Fórmula:**  
```
IBS_Valor = Base × Aliquota_Total_Destino
ONDE:
  Aliquota_Total_Destino = Aliquota_Estadual_Destino + Aliquota_Municipal_Destino
  Aliquota_Total_Destino = 0,009  SE Fase = "Shadow Run" (2026)
```

**Fonte:** POLICE-FIN-00001 §4.3 (IBS); EC 132/2023; LC 214/2025 (prevista)  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Status:** ⚠️ API do Comitê Gestor ainda não publicada  
**Testabilidade:** `TestIBS_Valor_ComAliquotaDestino`  

---

### BR-TAX-CALC-019 — Valor FUST

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do FUST deve ser calculado como: Receita Operacional Líquida × 1%, onde a Receita Operacional Líquida é o Valor Bruto do Serviço deduzido do ICMS, PIS e COFINS incidentes._

**Fórmula:**  
```
Base_FUST = Valor_Servico - ICMS_Valor - PIS_Valor - COFINS_Valor
FUST_Valor = Base_FUST × 0,01
```

**Fonte:** POLICE-FIN-00001 §4.3 (FUST/FUNTTEL); Lei 9.998/2000  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestFUST_Calculo_BaseLiquida`  

**Exemplo:**
- SCM = R$ 70,00; ICMS = R$ 17,50; PIS = R$ 0,45; COFINS = R$ 2,10
- Base FUST = 70,00 − 17,50 − 0,45 − 2,10 = **R$ 49,95**
- FUST = 49,95 × 1% = **R$ 0,50**

---

### BR-TAX-CALC-020 — Valor FUNTTEL

**Categoria:** Derivation  
**Declaração (SBVR):** _O valor do FUNTTEL deve ser calculado como: Receita Operacional Líquida × 0,5%, utilizando a mesma base de cálculo do FUST._

**Fórmula:**  
```
Base_FUNTTEL = Base_FUST (mesma base)
FUNTTEL_Valor = Base_FUNTTEL × 0,005
```

**Fonte:** POLICE-FIN-00001 §4.3 (FUST/FUNTTEL); Lei 10.052/2000  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestFUNTTEL_Calculo_BaseLiquida`  

**Exemplo:**
- Base = R$ 49,95 → FUNTTEL = 49,95 × 0,5% = **R$ 0,25**

---

# PARTE III — CONSTRAINT RULES (Restrições e Condições)

Regras que impõem limites, proibições e condições obrigatórias sobre os cálculos.

---

### BR-TAX-CONS-001 — Limite de Compensação de Prejuízos Fiscais

**Categoria:** Constraint — Range  
**Declaração (SBVR):** _É obrigatório que a compensação de prejuízos fiscais de exercícios anteriores não exceda 30% do lucro líquido ajustado do período._

**Fórmula da Restrição:**  
```
Compensacao_Maxima = Lucro_Liquido_Ajustado × 0,30
Compensacao_Efetiva = MIN(Prejuizo_Acumulado, Compensacao_Maxima)
```

**Fonte:** POLICE-FIN-00001 §4.3 (IRPJ/CSLL); Lei 8.981/1995 art. 42  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestIRPJ_CompensacaoPrejuizo_NaoExcede30pct`  

---

### BR-TAX-CONS-002 — Obrigatoriedade de Regime Não-Cumulativo para TELECOM

**Categoria:** Constraint — Mandatory  
**Declaração (SBVR):** _É obrigatório que empresas do setor de telecomunicações utilizem o regime não-cumulativo de PIS/COFINS._

**Fonte:** POLICE-FIN-00001 §4.3 (PIS/COFINS); Lei 10.833/2003  
**Prioridade:** Alta  
**Enforcement:** Automático (validação no input)  
**Testabilidade:** `TestPISCOFINS_Telecom_ValidaRegimeNaoCumulativo`  

---

### BR-TAX-CONS-003 — Não Incidência de IPI sobre Serviços de Telecom

**Categoria:** Constraint — Prohibited  
**Declaração (SBVR):** _É proibido calcular IPI sobre serviços de telecomunicação puros._

**Fonte:** POLICE-FIN-00001 §4.3 (IPI)  
**Prioridade:** Média  
**Enforcement:** Automático  
**Testabilidade:** `TestIPI_ServicoTelecom_NaoIncide`  

**Condição de Aplicação:** Se o item for classificado como serviço (não mercadoria), `IPI_Valor = 0`.

---

### BR-TAX-CONS-004 — Alíquota Interestadual de ICMS por Região

**Categoria:** Constraint — Conditional  
**Declaração (SBVR):** _É obrigatório que a alíquota interestadual de ICMS seja 7% quando a origem estiver nas regiões Sul ou Sudeste e o destino em outra região; e 12% nos demais casos._

**Regra:**  
```
AliqInterestadual = 4%   SE Origem_Mercadoria == 1 (Importada) AND Conteudo_Importacao > 40%
AliqInterestadual = 7%   SE UF_Origem ∈ {Sul, Sudeste} AND UF_Destino ∉ {Sul, Sudeste}
                          // Nota: Espírito Santo (ES) é tratado como destino de 7% nas saídas
                          // do Sul/Sudeste, equiparando-se a Norte/Nordeste/Centro-Oeste
AliqInterestadual = 12%  CASO CONTRARIO
                          // Inclusive para ES quando é origem (segue regra Sul/Sudeste)
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS); Resolução Senado 22/1989; Resolução Senado 13/2012  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestICMS_AliquotaInterestadual_SPparaBA_7pct`  

---

### BR-TAX-CONS-005 — DIFAL Zero para Operação Intra-Estadual

**Categoria:** Constraint — Conditional  
**Declaração (SBVR):** _É obrigatório que o DIFAL seja zero quando a operação for intra-estadual._

**Regra:**  
```
IF UF_Origem == UF_Destino THEN DIFAL = 0
```

**Fonte:** POLICE-FIN-00001 §5.3  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestDIFAL_IntraEstadual_RetornaZero`  

---

### BR-TAX-CONS-006 — DIFAL Zero para Não-Consumidor Final

**Categoria:** Constraint — Conditional  
**Declaração (SBVR):** _É obrigatório que o DIFAL seja zero quando o destinatário não for consumidor final (operação de revenda ou industrialização)._

**Regra:**  
```
IF Consumidor_Final == false THEN DIFAL = 0
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS); EC 87/2015  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestDIFAL_Revenda_RetornaZero`  

---

### BR-TAX-CONS-007 — Limite de Alíquota de ISS

**Categoria:** Constraint — Range  
**Declaração (SBVR):** _É obrigatório que a alíquota de ISS esteja no intervalo entre 2% e 5%, conforme legislação municipal._

**Regra:**  
```
2% ≤ Aliquota_ISS ≤ 5%
```

**Fonte:** POLICE-FIN-00001 §4.3 (ISS); LC 116/2003 art. 8º-A  
**Prioridade:** Média  
**Enforcement:** Automático (validação na entrada de dados)  
**Testabilidade:** `TestISS_AliquotaForaDoIntervalo_RetornaErro`  

---

### BR-TAX-CONS-008 — CBS Não Integra Própria Base

**Categoria:** Constraint — Mandatory  
**Declaração (SBVR):** _É obrigatório que a CBS seja calculada "por fora", não integrando sua própria base de cálculo._

**Regra:**  
```
CBS_Base_Nao_Ajustada // A base usada para calcular CBS NÃO é acrescida da própria CBS
```

**Fonte:** POLICE-FIN-00001 §4.3 (CBS); EC 132/2023  
**Prioridade:** Crítica  
**Enforcement:** Automático (design da fórmula)  
**Testabilidade:** `TestCBS_NaoRealimentaBase`  

---

### BR-TAX-CONS-009 — TTL Máximo do Cache de Alíquota IBS

**Categoria:** Constraint — Range  
**Declaração (SBVR):** _É obrigatório que o cache da alíquota IBS para uma jurisdição de destino tenha Time-To-Live máximo de 24 horas. Expirado o TTL, é obrigatória nova consulta à API do Comitê Gestor._

**Regra:**  
```
TTL_Cache_IBS ≤ 24h
IF Cache.Expired(IBS_Key) THEN Reconsultar_API_Comite_Gestor()
```

**Fonte:** POLICE-FIN-00001 §4.3 (IBS); SOP-002  
**Prioridade:** Alta  
**Enforcement:** Automático (Redis TTL)  
**Testabilidade:** `TestIBS_Cache_ExpiraEm24h`  

---

### BR-TAX-CONS-010 — IS é Pré-Filtro Obrigatório Antes da CBS

**Categoria:** Constraint — Mandatory  
**Declaração (SBVR):** _É obrigatório que a verificação de incidência do Imposto Seletivo seja executada antes do cálculo da CBS para cada item._

**Regra:**  
```
Ordem de Cálculo: IS → IPI → CBS → ICMS → IBS → ISS → PIS/COFINS → FUST → FUNTTEL
// IS: primeiro por ser pré-filtro (pode zerar base de itens específicos)
// IPI: "por fora", compõe base do ICMS para consumidor final
// CBS: "por fora", não compõe base de outros
// ICMS: "por dentro", seu valor destacado é excluído da base de PIS/COFINS (STF)
// IBS: "por fora", alíquota do destino
// ISS: sobre serviços, não interage com ICMS
// PIS/COFINS: base exclui ICMS destacado
// FUST/FUNTTEL: em cascata, base = bruto − ICMS − PIS − COFINS
```

**Fonte:** POLICE-FIN-00001 §4.3 (IS); SOP-003; SOP-013  
**Prioridade:** Crítica  
**Enforcement:** Automático (ordenação no pipeline de cálculo)  
**Testabilidade:** `TestPipeline_OrdemCalculo_IS_Antes_CBS`  

---

### BR-TAX-CONS-011 — PER/DCOMP Exige Aprovação do Tax Compliance Officer

**Categoria:** Constraint — Authorization  
**Declaração (SBVR):** _É obrigatório que a compensação cruzada de tributos federais via PER/DCOMP seja aprovada previamente pelo Tax Compliance Officer._

**Regra:**  
```
IF Tipo_Compensacao == "CRUZADA" THEN requer Aprovacao(Tax_Compliance_Officer)
```

**Fonte:** POLICE-FIN-00001 §9.0 (SOP-006)  
**Prioridade:** Alta  
**Enforcement:** Manual (workflow de aprovação)  
**Testabilidade:** `TestPERDCOMP_SemAprovacao_RetornaErro`  

---

### BR-TAX-CONS-012 — FUST/FUNTTEL Exclusivo para SCM/STFC

**Categoria:** Constraint — Prohibited  
**Declaração (SBVR):** _É proibido calcular FUST e FUNTTEL sobre Serviços de Valor Adicionado (SVA). Estas contribuições incidem exclusivamente sobre serviços de telecomunicação propriamente ditos (SCM, STFC)._

**Regra:**  
```
IF Natureza_Servico == "SVA" THEN FUST = 0 AND FUNTTEL = 0
// SVA inclui: streaming, antivírus, suporte técnico, e-mail hospedado, etc.
// SCM e STFC são serviços de telecom para fins destas contribuições
```

**Fonte:** POLICE-FIN-00001 §4.3 (FUST/FUNTTEL); Lei 9.998/2000; Lei 10.052/2000  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestFUST_SVA_NaoIncide`  

---

### BR-TAX-CONS-013 — CSTs Válidos para Desoneração de ICMS

**Categoria:** Constraint — Conditional  
**Declaração (SBVR):** _É obrigatório que a desoneração de ICMS somente seja aplicada quando o CST do item for 20, 30, 40, 41, 50, 70 ou 90. CST 00 (tributado integralmente) não permite desoneração._

**Regra:**  
```
IF CST_ICMS ∈ {20, 30, 40, 41, 50, 70, 90} THEN Desoneracao_Permitida = true
IF CST_ICMS == "00" THEN Desoneracao_Permitida = false
// CSTs 20 e 70: com redução de base
// CSTs 30, 40, 41, 50: isentos/não tributados/suspensos
// CST 90: outras situações (inclui desoneração genérica)
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS Desonerado)  
**Prioridade:** Alta  
**Enforcement:** Automático (validação pré-cálculo)  
**Testabilidade:** `TestDesoneracao_CSTInvalido_RetornaErro`  

---

# PARTE IV — INFERENCE RULES (Classificações e Deduções)

Regras que inferem novo conhecimento a partir de fatos existentes.

---

### BR-TAX-INF-001 — Classificação de Fase: Shadow Run (2026)

**Categoria:** Inference — Classification  
**Declaração (SBVR):** _Se a data da operação estiver contida no ano de 2026, então a fase tributária é classificada como "Shadow Run"._

**Regra:**  
```
IF Ano(Data_Operacao) == 2026 THEN Fase = "SHADOW_RUN"
```

**Fonte:** POLICE-FIN-00001 §5.1  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestPhaseResolver_2026_RetornaShadowRun`  

---

### BR-TAX-INF-002 — Classificação de Fase: CBS Plena (2027)

**Categoria:** Inference — Classification  
**Declaração (SBVR):** _Se a data da operação estiver contida no ano de 2027, então a fase tributária é classificada como "CBS Plena"._

**Regra:**  
```
IF Ano(Data_Operacao) == 2027 THEN Fase = "CBS_PLENA"
```

**Fonte:** POLICE-FIN-00001 §5.1  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestPhaseResolver_2027_RetornaCBSPlena`  

---

### BR-TAX-INF-003 — Classificação de Fase: Transição Subnacional (2029–2032)

**Categoria:** Inference — Classification  
**Declaração (SBVR):** _Se a data da operação estiver contida no intervalo de 2029 a 2032, então a fase tributária é classificada como "Transição Subnacional"._

**Regra:**  
```
IF Ano(Data_Operacao) >= 2029 AND Ano(Data_Operacao) <= 2032 THEN Fase = "TRANSICAO_SUBNACIONAL"
```

**Fonte:** POLICE-FIN-00001 §5.1  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestPhaseResolver_2030_RetornaTransicaoSubnacional`  

---

### BR-TAX-INF-004 — Classificação de Fase: IVA Dual (2033+)

**Categoria:** Inference — Classification  
**Declaração (SBVR):** _Se a data da operação for igual ou posterior a 2033, então a fase tributária é classificada como "IVA Dual"._

**Regra:**  
```
IF Ano(Data_Operacao) >= 2033 THEN Fase = "IVA_DUAL"
```

**Fonte:** POLICE-FIN-00001 §5.1  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestPhaseResolver_2033_RetornaIVADual`  

---

### BR-TAX-INF-005 — Incidência de Imposto Seletivo (IS)

**Categoria:** Inference — Eligibility  
**Declaração (SBVR):** _Se o NCM do item constar na tabela oficial de produtos sujeitos ao Imposto Seletivo e a flag isento_is for falsa, então o IS é aplicável._

**Regra:**  
```
IF NCM(Item) ∈ Tabela_NCM_Seletivo AND Item.isento_is == false THEN IS_Aplicavel = true
IF Item.isento_is == true THEN IS_Aplicavel = false  // Override manual
```

**Fonte:** POLICE-FIN-00001 §4.3 (IS)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Status:** ⚠️ Tabela oficial de NCMs sujeitos ao IS ainda não publicada  
**Testabilidade:** `TestIS_BebidaAlcoolica_Incide; TestIS_TelecomPuro_NaoIncide`  

**Produtos preliminarmente sujeitos (lista não exaustiva):**
- Bebidas alcoólicas (NCM 2203 a 2208)
- Tabaco (NCM 2402, 2403)
- Açúcar (NCM 1701)
- Veículos poluentes
- Combustíveis fósseis

---

### BR-TAX-INF-006 — Classificação de Saldo Tributário

**Categoria:** Inference — Status  
**Declaração (SBVR):** _Se o saldo de créditos menos débitos for positivo, então o status é "Saldo a Recuperar". Se negativo, o status é "Valor a Recolher"._

**Regra:**  
```
Saldo = Creditos_Periodo - Debitos_Periodo
IF Saldo > 0 THEN Status = "SALDO_A_RECUPERAR"  // Crédito acumulado para próximo período
IF Saldo < 0 THEN Status = "VALOR_A_RECOLHER"    // Gera guia de pagamento (DARF/GNRE/GPS)
IF Saldo == 0 THEN Status = "NEUTRO"
```

**Fonte:** POLICE-FIN-00001 §9.0  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestSaldo_CreditoMaiorQueDebito_RetornaARecuperar`  

---

### BR-TAX-INF-007 — Classificação de Serviço para FUST/FUNTTEL

**Categoria:** Inference — Eligibility  
**Declaração (SBVR):** _Se a natureza do serviço for SCM (Serviço de Comunicação Multimídia) ou STFC (Serviço Telefônico Fixo Comutado), então FUST e FUNTTEL são aplicáveis. Se a natureza for SVA (Serviço de Valor Adicionado), então FUST e FUNTTEL não são aplicáveis._

**Regra:**  
```
IF Natureza_Servico ∈ {SCM, STFC} THEN FUST_Aplicavel = true AND FUNTTEL_Aplicavel = true
IF Natureza_Servico == "SVA" THEN FUST_Aplicavel = false AND FUNTTEL_Aplicavel = false
```

**Fonte:** POLICE-FIN-00001 §4.3 (FUST/FUNTTEL)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestFUST_SCM_Incide; TestFUST_SVA_NaoIncide`  

---

# PARTE V — ACTION RULES (Disparadores e Workflows)

Regras que definem o que deve acontecer quando determinadas condições são satisfeitas.

---

### BR-TAX-ACT-001 — Circuit Breaker da API IBS

**Categoria:** Action — Trigger  
**Declaração (SBVR):** _Quando a API do Comitê Gestor do IBS falhar 3 vezes em um intervalo de 60 segundos, o sistema deve abrir o circuit breaker, utilizar a última alíquota cacheada com TTL estendido, e notificar imediatamente o Tax Compliance Officer._

**Regra:**  
```
WHEN API_IBS.Errors(3, within=60s)
THEN
  Circuit_Breaker.State = OPEN
  Usar_Ultima_Aliquota_Cacheada()
  Notificar("Tax Compliance Officer", "Circuit Breaker IBS ABERTO — usando cache de fallback")
  Schedule_Retry(interval=5min)
```

**Fonte:** POLICE-FIN-00001 §10 (R2), §4.3 (IBS), SOP-002  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestIBS_CircuitBreaker_AbreApos3Falhas_UsaCache`  

---

### BR-TAX-ACT-002 — Investigação de Divergência no Shadow Run

**Categoria:** Action — Trigger  
**Declaração (SBVR):** _Quando a variação entre o Shadow (CBS/IBS) e o Legacy (PIS+COFINS / ICMS+ISS) exceder 10% na reconciliação mensal, o Controller e o Engineering Lead devem investigar item a item._

**Regra:**  
```
WHEN ABS(Shadow_Total - Legacy_Total) / Legacy_Total > 0,10
THEN
  Investigar_Item_A_Item()
  Gerar_Relatorio_Divergencia()
  Agendar_Reuniao(Controller, Engineering_Lead)
```

**Fonte:** POLICE-FIN-00001 SOP-004  
**Prioridade:** Alta  
**Enforcement:** Semi-automático (alerta automático + investigação manual)  
**Testabilidade:** `TestShadowRun_DivergenciaMaior10pct_DisparaAlerta`  

---

### BR-TAX-ACT-003 — Escalonamento Nível 3 (Divergência > R$ 100.000)

**Categoria:** Action — Escalation  
**Declaração (SBVR):** _Quando uma divergência de cálculo exceder R$ 100.000 ou houver risco de autuação fiscal, o Comitê Fiscal deve reportar ao CFO e ao Departamento Jurídico em 12 horas, com comunicação ao Conselho de Administração em 5 dias úteis._

**Regra:**  
```
WHEN Divergencia > 100000 OR Risco_Autuacao == true
THEN
  Notificar(CFO, Juridico, prazo=12h)
  Comunicar(Conselho_Administracao, prazo=5_dias_uteis)
  Abrir_Processo_Investigacao()
```

**Fonte:** POLICE-FIN-00001 §3.3 (Nível 3)  
**Prioridade:** Crítica  
**Enforcement:** Manual (workflow de escalonamento)  
**Testabilidade:** `TestEscalonamento_DivergenciaMaior100k_NotificaCFO`  

---

### BR-TAX-ACT-004 — Atualização de Alíquota por Nova Legislação

**Categoria:** Action — Workflow  
**Declaração (SBVR):** _Quando uma nova legislação que altere alíquotas for publicada, o Tax Compliance Officer e o Engineering Lead devem concluir o deploy da atualização em produção em no máximo 5 dias úteis._

**Workflow:**  
```
AFTER Publicacao_Nova_Legislacao
  1. Tax Compliance Officer registra em tax_rate_change_log (Dia 0)
  2. Engineering Lead cria branch tax-update/YYYY-MM-DD-descricao (Dia 0)
  3. Atualizar tabelas SQL de alíquotas (Dia 0–1)
  4. Atualizar models/constants.go se aplicável (Dia 0–1)
  5. PR com dupla aprovação (Engineering + Tax) (Dia 1–2)
  6. Deploy em staging → teste de regressão (Dia 2–3)
  7. Deploy em produção (Dia 3–5)
  8. Comunicar Financeiro sobre vigência e impacto (Dia 5)
```

**Fonte:** POLICE-FIN-00001 SOP-005  
**Prioridade:** Crítica  
**Enforcement:** Manual (workflow de desenvolvimento)  
**Testabilidade:** `TestAtualizacaoAliquota_CicloCompleto_5DiasUteis`  

---

### BR-TAX-ACT-005 — Shadow Run Não Compõe Total a Pagar

**Categoria:** Action — Trigger  
**Declaração (SBVR):** _Quando a fase for Shadow Run, os valores de CBS e IBS calculados não devem compor o total_a_pagar da operação, sendo armazenados separadamente como shadow_tax_result._

**Regra:**  
```
WHEN Fase == "SHADOW_RUN"
THEN
  CBS_Valor → shadow_tax_result.cbs  // Não soma em total_a_pagar
  IBS_Valor → shadow_tax_result.ibs  // Não soma em total_a_pagar
  PIS_Valor + COFINS_Valor + ICMS_Valor + ISS_Valor → total_a_pagar
```

**Fonte:** POLICE-FIN-00001 §5.2, SOP-001  
**Prioridade:** Crítica  
**Enforcement:** Automático  
**Testabilidade:** `TestShadowRun_CBS_NaoCompoeTotalAPagar`  

---

### BR-TAX-ACT-006 — Extinção de Tributos na Fase IVA Dual

**Categoria:** Action — Workflow  
**Declaração (SBVR):** _Quando a fase for IVA Dual (2033+), os tributos PIS, COFINS, ICMS e ISS são considerados extintos e não devem ser calculados._

**Regra:**  
```
WHEN Fase == "IVA_DUAL"
THEN
  PIS_Valor = 0    // Extinto
  COFINS_Valor = 0 // Extinto
  ICMS_Valor = 0   // Extinto
  ISS_Valor = 0    // Extinto
  // Apenas CBS, IBS, IPI, IRPJ, CSLL, CPP, FGTS e IS são calculados
```

**Fonte:** POLICE-FIN-00001 §5.1, §5.2  
**Prioridade:** Média (vigência em 2033)  
**Enforcement:** Automático  
**Testabilidade:** `TestIVADual_TributosLegacy_NaoCalculados`  

---

### BR-TAX-ACT-007 — Abatimento do ICMS Desonerado no Valor da Nota

**Categoria:** Action — Trigger  
**Declaração (SBVR):** _Quando a desoneração de ICMS for aplicada, o valor do ICMS desonerado (vICMSDeson) deve ser subtraído do valor total da nota fiscal, para que o benefício seja repassado ao adquirente._

**Regra:**  
```
WHEN Desoneracao_Aplicada == true
THEN
  vICMSDeson = (Valor_Item × Aliquota_Nominal) − ICMS_Efetivo
  Valor_Final_Item = Valor_Item − vICMSDeson
  // O campo vICMSDeson do XML da NF-e deve ser preenchido
  // O campo motDesICMS deve ser preenchido com o código oficial
```

**Fonte:** POLICE-FIN-00001 §4.3 (ICMS Desonerado)  
**Prioridade:** Alta  
**Enforcement:** Automático  
**Testabilidade:** `TestDesoneracao_AbateValorTotal`  

---

# PARTE VI — TABELAS DE DECISÃO

## DT-001: Seleção de Tributos por Fase e Natureza da Operação

```
┌──────────────────────────────────────────────────────────────────────────────┐
│ Decision Table: Seleção de Tributos por Fase e Natureza da Operação          │
│ Ref: BR-TAX-INF-001 a 004 + BR-TAX-CONS-005 a 006                           │
├──────────────────────────┬────────────┬────────────┬────────────┬────────────┤
│ CONDITIONS               │ Shadow Run │ CBS Plena  │ Transição  │ IVA Dual   │
│                          │ (2026)     │ (2027)     │ (2029-32)  │ (2033+)    │
├──────────────────────────┼────────────┼────────────┼────────────┼────────────┤
│ Fase                     │ SHADOW_RUN │ CBS_PLENA  │ TRANSICAO  │ IVA_DUAL   │
├──────────────────────────┼────────────┼────────────┼────────────┼────────────┤
│ ACTIONS                  │            │            │            │            │
├──────────────────────────┼────────────┼────────────┼────────────┼────────────┤
│ Calcular PIS             │     X      │            │            │            │
│ Calcular COFINS          │     X      │            │            │            │
│ Calcular CBS             │  (shadow)  │     X      │     X      │     X      │
│ Calcular ICMS            │     X      │     X      │  X × fator │            │
│ Calcular DIFAL           │     X*     │     X*     │  X* × fator│            │
│ Calcular ICMS-ST         │     X*     │     X*     │  X* × fator│            │
│ Calcular IBS             │  (shadow)  │  (shadow)  │     X      │     X      │
│ Calcular ISS             │     X      │     X      │  X × fator │            │
│ Calcular IS              │     X**    │     X**    │     X**    │     X**    │
│ Calcular IPI             │     X**    │     X**    │     X**    │     X**    │
│ Calcular FUST            │     X***   │     X***   │     X***   │     X***   │
│ Calcular FUNTTEL         │     X***   │     X***   │     X***   │     X***   │
│ Calcular IRPJ/CSLL       │     X      │     X      │     X      │     X      │
│ Calcular CPP/FGTS        │     X      │     X      │     X      │     X      │
│                          │            │            │            │            │
│ * = se interestadual e   │            │            │            │            │
│     consumidor final     │            │            │            │            │
│ ** = se aplicável ao item│            │            │            │            │
│ *** = apenas serviços de │            │            │            │            │
│   telecom (SCM/STFC);    │            │            │            │            │
│   não incide sobre SVA   │            │            │            │            │
│ fator = % remanescente   │            │            │            │            │
│   do regime legacy       │            │            │            │            │
└──────────────────────────┴────────────┴────────────┴────────────┴────────────┘
```

## DT-002: Determinação de Aplicabilidade do DIFAL

```
┌──────────────────────────────────────────────────────────────────┐
│ Decision Table: Aplicabilidade do DIFAL                          │
│ Ref: BR-TAX-CALC-012, BR-TAX-CONS-005, BR-TAX-CONS-006          │
├──────────────────────────┬──────────┬──────────┬──────────┬─────┤
│ CONDITIONS               │    R1    │    R2    │    R3    │ R4  │
├──────────────────────────┼──────────┼──────────┼──────────┼─────┤
│ UF_Origem ≠ UF_Destino   │    Y     │    Y     │    Y     │  N  │
│ Consumidor_Final = true  │    Y     │    Y     │    N     │  -  │
│ AliqIntDest > AliqInter  │    Y     │    N     │    -     │  -  │
├──────────────────────────┼──────────┼──────────┼──────────┼─────┤
│ ACTIONS                  │          │          │          │     │
├──────────────────────────┼──────────┼──────────┼──────────┼─────┤
│ Calcular DIFAL           │    X     │          │          │     │
│ DIFAL = 0                │          │    X     │    X     │  X  │
│ Apenas ICMS Próprio      │          │    X     │    X     │  X  │
│ Responsável Recolhimento │ Remetente│   N/A    │   N/A    │ N/A │
└──────────────────────────┴──────────┴──────────┴──────────┴─────┘
```

## DT-003: Escalonamento por Divergência de Cálculo

```
┌──────────────────────────────────────────────────────────────────────────┐
│ Decision Table: Escalonamento por Divergência de Cálculo                 │
│ Ref: BR-TAX-ACT-003, POLICE-FIN-00001 §3.3                              │
├──────────────────────────┬──────────────┬──────────────┬────────────────┤
│ CONDITIONS               │   Nível 1    │   Nível 2    │    Nível 3     │
├──────────────────────────┼──────────────┼──────────────┼────────────────┤
│ Valor_Divergencia        │  ≤ R$ 10.000 │ R$10K–100K   │  > R$ 100.000  │
│ Risco_Autuacao           │      N       │      N       │       Y        │
│ Risco_Criminal           │      N       │      N       │       N        │
├──────────────────────────┼──────────────┼──────────────┼────────────────┤
│ ACTIONS                  │              │              │                │
├──────────────────────────┼──────────────┼──────────────┼────────────────┤
│ Responsável              │ Engineering  │ Tax Compliance│ Comitê Fiscal │
│                          │ Lead         │ Officer       │                │
│ Prazo Resposta           │     48h      │     24h       │      12h       │
│ Notificar CFO            │              │      X        │       X        │
│ Notificar Jurídico       │              │              │       X        │
│ Notificar Conselho       │              │              │    5 dias      │
│ Documentar Causa Raiz    │      X       │      X        │       X        │
└──────────────────────────┴──────────────┴──────────────┴────────────────┘
```

**Nível 4 — Risco Criminal (Lei 8.137/90):** Adicionado como caso especial. CFO aciona Conselho e assessoria jurídica externa imediatamente, independentemente do valor.

## DT-004: Seleção de Alíquota PIS/COFINS por Regime

```
┌──────────────────────────────────────────────────────────────────┐
│ Decision Table: Seleção de Alíquota PIS/COFINS                  │
│ Ref: BR-TAX-DEF-003, BR-TAX-DEF-004, BR-TAX-CONS-002           │
├──────────────────────────┬──────────────────┬───────────────────┤
│ CONDITIONS               │   Cumulativo     │  Não-Cumulativo   │
├──────────────────────────┼──────────────────┼───────────────────┤
│ Regime_Tributario        │  CUMULATIVO      │ NAO_CUMULATIVO    │
│ Setor                    │  qualquer        │ TELECOM (fixo)    │
├──────────────────────────┼──────────────────┼───────────────────┤
│ ACTIONS                  │                  │                   │
├──────────────────────────┼──────────────────┼───────────────────┤
│ Aliquota_PIS             │      0,65%       │       1,65%       │
│ Aliquota_COFINS          │      3,00%       │       7,60%       │
│ Aliquota_Total           │      3,65%       │       9,25%       │
│ Direito_Credito          │      Não         │       Sim         │
│ CST_PIS_Padrao           │       01         │        01         │
│ CST_COFINS_Padrao        │       01         │        01         │
│ CST_Credito              │      N/A         │        50         │
└──────────────────────────┴──────────────────┴───────────────────┘
```

## DT-005: Sistema Tributário por Fase (Máquina de Estados)

```
┌────────────────────────────────────────────────────────────────────────────┐
│ Decision Table: Máquina de Estados do Sistema Tributário                    │
│ Ref: BR-TAX-INF-001 a 004, §5.2                                            │
├────────────────────────┬───────────┬───────────┬─────────────┬─────────────┤
│ CONDITIONS             │ Shadow Run│ CBS Plena │ Transição   │ IVA Dual    │
├────────────────────────┼───────────┼───────────┼─────────────┼─────────────┤
│ Ano                    │   2026    │   2027    │ 2029–2032   │   2033+     │
├────────────────────────┼───────────┼───────────┼─────────────┼─────────────┤
│ ACTIONS                │           │           │             │             │
├────────────────────────┼───────────┼───────────┼─────────────┼─────────────┤
│ Motor Legacy Ativo     │    Sim    │  Parcial  │   Parcial   │    Não      │
│ Motor Reforma Ativo    │  Shadow   │  Parcial  │   Parcial   │    Sim      │
│ PIS/COFINS             │  Ativo    │  Extinto  │   Extinto   │   Extinto   │
│ CBS                    │  Shadow   │  Ativo    │   Ativo     │   Ativo     │
│ ICMS                   │  Ativo    │  Ativo    │  Redução    │   Extinto   │
│ IBS                    │  Shadow   │  Shadow   │   Ativo     │   Ativo     │
│ ISS                    │  Ativo    │  Ativo    │  Redução    │   Extinto   │
│ FUST/FUNTTEL           │  Ativo    │  Ativo    │   Ativo     │   Ativo     │
│ Recolher Legacy        │    Sim    │  Parcial  │   Parcial   │    Não      │
│ Recolher Reforma       │    Não    │  Parcial  │   Parcial   │    Sim      │
└────────────────────────┴───────────┴───────────┴─────────────┴─────────────┘
```

---

# PARTE VII — MATRIZ DE RASTREABILIDADE

## Regra → Requisito Técnico → Teste → Componente

| Regra | Requisito Técnico | Função/Método Esperado | Teste Unitário | Componente |
|:---|:---|:---|:---|:---|
| `BR-TAX-CALC-001` | Calcular base IRPJ Lucro Real com ajustes LALUR | `IRPJCalculator.CalcBaseLucroReal(lucroContabil, adicoes, exclusoes, prejuizo) float64` | `TestIRPJ_BaseLucroReal_ComAjustesLALUR` | `irpj_calculator.go` |
| `BR-TAX-CALC-002` | Calcular base IRPJ Lucro Presumido | `IRPJCalculator.CalcBaseLucroPresumido(receitaBruta, cnae) float64` | `TestIRPJ_BaseLucroPresumido_Servico` | `irpj_calculator.go` |
| `BR-TAX-CALC-003` | Calcular valor IRPJ | `IRPJCalculator.Calculate(base) float64` | `TestIRPJ_Valor` | `irpj_calculator.go` |
| `BR-TAX-CALC-004` | Calcular adicional IRPJ | `IRPJCalculator.CalcAdicional(base, nMeses) float64` | `TestIRPJ_Adicional_ExcedeLimite` | `irpj_calculator.go` |
| `BR-TAX-CALC-005` | Calcular valor CSLL | `CSLLCalculator.Calculate(base) float64` | `TestCSLL_Valor` | `csll_calculator.go` |
| `BR-TAX-CALC-006` | Calcular PIS Cumulativo | `PISCalculator.CalcCumulativo(receitaBruta) float64` | `TestPIS_Cumulativo` | `pis_cofins_calculator.go` |
| `BR-TAX-CALC-007` | Calcular PIS Não-Cumulativo | `PISCalculator.CalcNaoCumulativo(receitaBruta, creditos) float64` | `TestPIS_NaoCumulativo_ComCreditos` | `pis_cofins_calculator.go` |
| `BR-TAX-CALC-008` | Calcular COFINS Cumulativo | `COFINSCalculator.CalcCumulativo(receitaBruta) float64` | `TestCOFINS_Cumulativo` | `pis_cofins_calculator.go` |
| `BR-TAX-CALC-009` | Calcular COFINS Não-Cumulativo | `COFINSCalculator.CalcNaoCumulativo(receitaBruta, creditos) float64` | `TestCOFINS_NaoCumulativo` | `pis_cofins_calculator.go` |
| `BR-TAX-CALC-010` | Compor base de cálculo do ICMS | `ICMSCalculator.CalcBase(valor, frete, seguro, despesas, ipi, consumidorFinal) float64` | `TestICMS_BaseCalculo_CompoeAcessorias` | `icms_calculator.go` |
| `BR-TAX-CALC-011` | Calcular ICMS Próprio | `ICMSCalculator.Calculate(base, aliquota) float64` | `TestICMS_Proprio` | `icms_calculator.go` |
| `BR-TAX-CALC-012` | Calcular DIFAL | `ICMSCalculator.CalcularDIFAL(base, aliqInternaDest, aliqInterestadual) float64` | `TestDIFAL_SP_Para_BA` | `icms_calculator.go` |
| `BR-TAX-CALC-013` | Calcular base ICMS-ST | `ICMSCalculator.CalcBaseST(valor, mva) float64` | `TestICMS_ST_BaseCalculo` | `icms_st_calculator.go` |
| `BR-TAX-CALC-014` | Calcular valor ICMS-ST | `ICMSCalculator.CalcValorST(baseST, aliqInterna, icmsProprio) float64` | `TestICMS_ST_Valor` | `icms_st_calculator.go` |
| `BR-TAX-CALC-015` | Calcular IPI por NCM | `IPICalculator.Calculate(valor, ncm) float64` | `TestIPI_PorNCM` | `ipi_calculator.go` |
| `BR-TAX-CALC-016` | Calcular ISS | `ISSCalculator.Calculate(precoServico, aliquotaMunicipal) float64` | `TestISS_Valor` | `iss_calculator.go` |
| `BR-TAX-CALC-017` | Calcular CBS "por fora" | `CBSCalculator.Calculate(base, aliquota) float64` | `TestCBS_NaoCompoePropriaBase` | `cbs_calculator.go` |
| `BR-TAX-CALC-018` | Calcular IBS por destino | `IBSCalculator.Calculate(base, aliquotaDestino) float64` | `TestIBS_Valor_Destino` | `ibs_calculator.go` |
| `BR-TAX-CONS-001` | Validar limite de 30% de compensação | `IRPJCalculator.ValidateCompensacao(compensacao, lucroAjustado) error` | `TestIRPJ_Compensacao_NaoExcede30pct` | `irpj_calculator.go` |
| `BR-TAX-CONS-002` | Forçar regime não-cumulativo TELECOM | `RegimeValidator.Validate(cnae, regime) error` | `TestValidacao_Telecom_ExigeNaoCumulativo` | `regime_validator.go` |
| `BR-TAX-CONS-004` | Resolver alíquota interestadual ICMS | `ICMSResolver.GetAliqInterestadual(ufOrigem, ufDestino) float64` | `TestICMS_Aliquota_SP_BA_7pct` | `icms_resolver.go` |
| `BR-TAX-CONS-005` | DIFAL zero intra-estadual | `ICMSCalculator.CalcularDIFAL(...)` retorna 0 se UF igual | `TestDIFAL_IntraEstadual_Zero` | `icms_calculator.go` |
| `BR-TAX-CONS-007` | Validar range alíquota ISS | `ISSValidator.ValidateAliquota(aliq) error` | `TestISS_Aliquota_ForaRange_Erro` | `iss_validator.go` |
| `BR-TAX-CONS-009` | TTL cache IBS ≤ 24h | `IBSCache.Set(ibgeCode, rate, ttl=24h)` | `TestIBS_Cache_TTL_24h` | `ibs_cache.go` |
| `BR-TAX-CONS-010` | IS antes da CBS no pipeline | `TaxPipeline.Execute()` garante ordem IS→CBS→IBS | `TestPipeline_Ordem_IS_CBS_IBS` | `tax_pipeline.go` |
| `BR-TAX-INF-001` a `004` | Resolver fase tributária | `PhaseResolver.Resolve(date time.Time) Phase` | `TestPhaseResolver_2026_2033` | `phase_resolver.go` |
| `BR-TAX-INF-005` | Determinar incidência de IS | `ISFilter.IsApplicable(ncm string, isento bool) bool` | `TestIS_Aplicavel_Bebida; TestIS_NaoAplicavel_Telecom` | `is_filter.go` |
| `BR-TAX-INF-006` | Classificar saldo tributário | `TaxPeriod.ClosePeriod(creditos, debitos) SaldoStatus` | `TestSaldo_CreditoMaior_ARecuperar` | `tax_period.go` |
| `BR-TAX-ACT-001` | Circuit breaker API IBS | `IBSClient.GetRate(ibgeCode) (float64, error)` com circuit breaker | `TestIBS_CircuitBreaker_Abre` | `ibs_client.go` |
| `BR-TAX-ACT-005` | Shadow run não compõe total | `TaxResponseBuilder.Build(...)` separa shadow de production | `TestShadowRun_CBS_Nao_TotalAPagar` | `tax_response_builder.go` |
| `BR-TAX-DEF-010` | Definição FUST | `FUSTCalculator` — configurado com alíquota 1% | `TestFUST_Definicao` | `fust_calculator.go` |
| `BR-TAX-DEF-011` | Definição FUNTTEL | `FUNTTELCalculator` — configurado com alíquota 0,5% | `TestFUNTTEL_Definicao` | `funttel_calculator.go` |
| `BR-TAX-CALC-019` | Calcular FUST | `FUSTCalculator.Calculate(valorServico, icms, pis, cofins) float64` | `TestFUST_Calculo_BaseLiquida` | `fust_calculator.go` |
| `BR-TAX-CALC-020` | Calcular FUNTTEL | `FUNTTELCalculator.Calculate(baseFUST) float64` | `TestFUNTTEL_Calculo_BaseLiquida` | `funttel_calculator.go` |
| `BR-TAX-CALC-021` | ICMS Desonerado — Redução Base | `ICMSCalculator.CalcDesoneradoReducao(valor, aliquota, pctReducao) (icms, vICMSDeson float64)` | `TestICMS_Desonerado_ReducaoBase` | `icms_calculator.go` |
| `BR-TAX-CALC-022` | ICMS Desonerado — Limitação Alíquota | `ICMSCalculator.CalcDesoneradoLimite(valor, aliqNominal, aliqAlvo) (icms, vICMSDeson float64)` | `TestICMS_Desonerado_LimitacaoAliquota` | `icms_calculator.go` |
| `BR-TAX-CONS-012` | FUST/FUNTTEL só SCM/STFC | `TelecomValidator.IsTelecomService(natureza) bool` | `TestFUST_SVA_NaoIncide` | `telecom_validator.go` |
| `BR-TAX-CONS-013` | CSTs válidos para desoneração | `ICMSValidator.ValidateCSTDesoneracao(cst) error` | `TestDesoneracao_CSTInvalido_RetornaErro` | `icms_validator.go` |
| `BR-TAX-INF-007` | SCM/SVA para FUST/FUNTTEL | `TelecomClassifier.Classify(natureza) ServiceType` | `TestFUST_SCM_Incide` | `telecom_classifier.go` |
| `BR-TAX-ACT-007` | Abater ICMS Desonerado | `TaxResponseBuilder.ApplyDesoneracao(item, vICMSDeson)` | `TestDesoneracao_AbateValorTotal` | `tax_response_builder.go` |

---

# PARTE VIII — VALIDAÇÃO DO CATÁLOGO

## Resumo de Completude

| Dimensão | Status | Detalhe |
|:---|:---|:---|
| Todas as condições especificadas | ✅ | 60 regras com condições explícitas |
| Caso default definido | ✅ | DT-005 cobre todas as fases até 2033+ |
| Exceções documentadas | ✅ | Cada regra lista exceções quando aplicável |
| Fonte identificada | ✅ | Toda regra referencia a política e a legislação |
| Exemplos positivos/negativos | ✅ | Regras críticas incluem exemplos numéricos |
| Testabilidade | ✅ | Toda regra possui ID de teste correspondente |
| Rastreabilidade ponta a ponta | ✅ | Matriz Regra→Requisito→Teste→Componente completa |
| Conflitos detectados | 2 | C1: CBS "por fora" × crédito; C2: Shadow vs Produção |

## Conflitos Conhecidos

| ID | Regras Envolvidas | Descrição | Resolução |
|:---|:---|:---|:---|
| C1 | BR-TAX-CALC-017 × §9.2 | CBS "por fora" simplifica a base, mas a não-cumulatividade plena exige tracking de créditos sobre insumos | Resolvido com dois passos: (1) CBS sobre base bruta, (2) apuração de créditos separadamente no fechamento do período |
| C2 | BR-TAX-ACT-005 × SOP-001 | Shadow Run calcula CBS/IBS em paralelo com PIS/COFINS; risco de double-counting em relatórios | Resolvido com flag `shadow_tax_result` isolando os valores de shadow do `total_a_pagar` |

## Gaps Bloqueantes (Dependências Externas)

| ID | Gap | Regras Impactadas | Bloqueio | Previsão |
|:---|:---|:---|:---|:---|
| G1 | Alíquota CBS setorial TELECOM não definida | BR-TAX-CALC-017 | Cálculo em produção (2027) | A definir pelo Min. Fazenda |
| G2 | API Comitê Gestor IBS não publicada | BR-TAX-CALC-018, BR-TAX-ACT-001 | Cálculo em produção | LC 214/2025 pendente |
| G3 | Tabela NBS não mapeada para c_class_trib | BR-TAX-CALC-017, BR-TAX-CALC-018 | Classificação automática | A publicar pelo Gov. Federal |
| G4 | Lista oficial NCMs sujeitos ao IS | BR-TAX-INF-005 | Cálculo do IS em produção | A publicar pelo Min. Fazenda |
| G5 | Cronograma oficial de redução ICMS/ISS | DT-001 (Fase 3) | Cálculo exato 2029–2032 | A definir pelo CONFAZ |

---

## Controle de Versão

| Versão | Data | Autor | Alterações |
|:---|:---|:---|:---|
| 1.0 | 2026-06-21 | Comitê Fiscal + Engineering Lead | Versão inicial. 50 regras de negócio em 5 categorias, 5 tabelas de decisão, matriz de rastreabilidade. Vinculado à POLICE-FIN-00001 v1.0. |
| 1.1 | 2026-06-21 | Comitê Fiscal + Engineering Lead | Adicionadas 10 regras: FUST/FUNTTEL (BR-TAX-DEF-010/011, BR-TAX-CALC-019/020, BR-TAX-CONS-012, BR-TAX-INF-007), ICMS Desonerado (BR-TAX-CALC-021/022, BR-TAX-CONS-013, BR-TAX-ACT-007), alíquota ICMS 4% para importados (BR-TAX-CONS-004 atualizada). Pipeline order corrigido (BR-TAX-CONS-010). DTs atualizadas. Total: 60 regras. Vinculado à POLICE-FIN-00001 v1.1. |

---

**Próxima Revisão Obrigatória:** 2026-12-21 (alinhada à política-mãe), ou na ocorrência de qualquer um dos eventos: publicação da LC 214/2025 (IBS), definição da alíquota CBS setorial, publicação da lista oficial de NCMs do IS.

---

_Documento classificado como **Confidencial — Uso Interno**. Vinculado à POLICE-FIN-00001. Distribuição controlada pelo Tax Compliance Officer._
