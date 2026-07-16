# Glossário do Projeto

> **Programa:** Adequação Corporativa à Reforma Tributária Nacional
> **Código:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
> **Versão:** 1.0
> **Atualizado:** 2026-07-08
> **Objetivo:** Fonte única de verdade para a terminologia do domínio tributário, alinhando o vocabulário entre time de negócios, time técnico e agentes de IA.

---

## 1. Tributos

### CBS — Contribuição sobre Bens e Serviços

Tributo federal criado pela EC 132/2023 que substitui progressivamente PIS, COFINS e IPI. De competência da União, incide sobre o consumo de bens e serviços em todo o território nacional. Aplicado "por fora" (não integra a base de cálculo do produto/serviço).

- **Competência:** Federal (União)
- **Alíquota de teste (2026):** 0,9%
- **Alíquota definitiva:** a definir pelo governo federal
- **Substitui:** PIS, COFINS, IPI (progressivamente até 2027)
- **Sinônimos no código:** `CBS`, `cbs`, `federalVat`, `contribuicaoFederal`

### IBS — Imposto sobre Bens e Serviços

Tributo subnacional criado pela EC 132/2023 que substitui progressivamente ICMS (estadual) e ISS (municipal). De competência compartilhada entre estados e municípios, gerido pelo Comitê Gestor do IBS.

- **Competência:** Subnacional (Estados + Municípios)
- **Alíquota de teste (2026):** 0,1%
- **Alíquota definitiva:** a definir por cada ente federativo, dentro dos limites do Comitê Gestor
- **Substitui:** ICMS, ISS (progressivamente até 2032)
- **Particularidade:** A alíquota varia conforme o **município de destino** do consumo
- **Sinônimos no código:** `IBS`, `ibs`, `subnationalVat`, `impostoSubnacional`

### IVA Dual

Designação do conjunto CBS (federal) + IBS (subnacional) que compõe o modelo brasileiro de Imposto sobre Valor Agregado. "Dual" porque são dois tributos distintos (federal e subnacional) operando sob o mesmo fato gerador, com regras de creditamento e não-cumulatividade similares.

- **Sinônimos no código:** `ivaDual`, `dualVat`, `IvaDual`

### IS — Imposto Seletivo

Tributo de natureza extrafiscal (desestimula o consumo) incidente sobre produtos com externalidades negativas: açúcar, tabaco, álcool, veículos poluentes, entre outros definidos em lei complementar. Não gera crédito tributário para o adquirente.

- **Competência:** Federal (União)
- **Incidência:** Produtos específicos (lista definida em lei complementar)
- **Creditamento:** NÃO permite crédito
- **Sinônimos no código:** `IS`, `impostoSeletivo`, `sinTax`

### Tributos em Extinção (Modelo Antigo)

| Tributo | Sigla | Competência | Incidência | Extinção prevista |
|---------|-------|-------------|------------|-------------------|
| Imposto sobre Circulação de Mercadorias e Serviços | **ICMS** | Estadual | Circulação de mercadorias, transporte interestadual, comunicação | 2029–2032 (redução gradativa) |
| Imposto sobre Serviços de Qualquer Natureza | **ISS** | Municipal | Prestação de serviços | 2029–2032 (redução gradativa) |
| Programa de Integração Social | **PIS** | Federal | Faturamento/receita | 2027 (extinção) |
| Contribuição para o Financiamento da Seguridade Social | **COFINS** | Federal | Faturamento/receita | 2027 (extinção) |
| Imposto sobre Produtos Industrializados | **IPI** | Federal | Produtos industrializados | 2027 (extinção) |
| Diferencial de Alíquotas (interestadual B2C) | **DIFAL** | Estadual | Vendas interestaduais a consumidor final | 2029–2032 (absorvido pelo IBS) |

---

## 2. Regimes Tributários

### Lucro Real

Regime de apuração do Imposto de Renda Pessoa Jurídica (IRPJ) e da Contribuição Social sobre o Lucro Líquido (CSLL) baseado no lucro contábil ajustado (livro fiscal). Obrigatório para empresas com receita bruta anual acima de R$ 78 milhões ou que exerçam atividades específicas. É o regime adotado pela companhia.

- **Relevância para o projeto:** Permite a apropriação de créditos de CBS e IBS sobre insumos, despesas e ativos, essencial para a não-cumulatividade do IVA Dual.
- **Sinônimos no código:** `lucroReal`, `realProfit`, `taxRegimeReal`

### Não-Cumulatividade

Princípio constitucional do IVA Dual (Art. 156-A, §1º, CF) pelo qual o contribuinte pode abater o imposto pago na etapa anterior (entrada) do imposto devido na etapa seguinte (saída), evitando a tributação em cascata.

- **Exemplo:** Empresa compra insumo por R$ 100 com CBS de R$ 9 (9%). Vende produto por R$ 200 com CBS de R$ 18. Pode creditar R$ 9 da entrada e recolher apenas R$ 9.
- **Sinônimos no código:** `naoCumulatividade`, `nonCumulative`, `creditamento`

### Simples Nacional

Regime tributário simplificado para micro e pequenas empresas (receita bruta anual até R$ 4,8 milhões). **Não adotado pela companhia**, mas relevante para qualificação de fornecedores — fornecedores do Simples Nacional NÃO geram crédito de CBS/IBS para a empresa adquirente.

- **Sinônimos no código:** `simplesNacional`, `simples`

### Período Híbrido (2029–2032)

Fase de transição constitucional em que os tributos antigos (ICMS, ISS) coexistem com os novos (CBS, IBS), com redução gradativa dos antigos e aumento progressivo dos novos. Exige dupla apuração e capacidade de operar simultaneamente os dois modelos.

- **Sinônimos no código:** `periodoHibrido`, `hybridPeriod`, `transicao`

---

## 3. Mecanismos e Operações Fiscais

### Split Payment

Mecanismo de recolhimento instantâneo pelo qual CBS e IBS são liquidados automaticamente no momento da transação financeira, separando o valor do tributo da receita líquida da empresa. O adquirente paga o valor total, a instituição financeira retém o percentual correspondente a CBS e IBS e repassa diretamente aos entes públicos.

- **Impacto no capital de giro:** Reduz o fluxo de caixa disponível, pois o tributo não transita mais pela tesouraria da empresa.
- **Sinônimos no código:** `splitPayment`, `splitPagamento`, `liquidacaoInstantanea`

### Creditamento

Direito de abater o valor de CBS/IBS pago na aquisição de insumos, mercadorias, serviços e ativos do valor devido na venda ou faturamento. Essencial para a não-cumulatividade e para a preservação de margens no Lucro Real.

- **Condições para creditamento:**
  1. O fornecedor deve estar em regime que gere crédito (Lucro Real ou Lucro Presumido — NÃO Simples Nacional)
  2. O bem ou serviço deve ser insumo essencial à atividade da empresa
  3. A nota fiscal de entrada deve destacar corretamente CBS e IBS
- **Sinônimos no código:** `creditamento`, `taxCredit`, `apropriacaoCredito`

### Fato Gerador

Evento econômico ou jurídico que faz nascer a obrigação tributária. Para CBS e IBS, o fato gerador é a circulação econômica do bem ou serviço (venda, transferência, importação).

- **Sinônimos no código:** `fatoGerador`, `taxableEvent`

### Base de Cálculo (BC)

Valor sobre o qual a alíquota do tributo é aplicada. Para CBS e IBS, a base de cálculo é o valor da operação (preço do bem ou serviço, incluindo frete, seguro e encargos), sem incluir o próprio tributo (cálculo "por fora").

- **Sinônimos no código:** `baseCalculo`, `baseCalculo`, `taxBase`

### Alíquota

Percentual aplicado sobre a base de cálculo para determinar o valor do tributo devido. No contexto do IBS, a alíquota varia conforme o município de destino do consumo e o tipo de bem/serviço.

- **Tipos de alíquota:**
  - **Alíquota cheia/padrão:** Aplicável à maioria dos bens e serviços
  - **Alíquota reduzida:** Aplicável a setores específicos (saúde, educação, alimentos)
  - **Isenção:** Alíquota zero para setores definidos em lei
- **Sinônimos no código:** `aliquota`, `taxRate`, `rate`

### Subvenção de Investimento (Incentivo Fiscal)

Benefício fiscal concedido por estados ou municípios (ex: Santana de Parnaíba) para atrair investimentos. No Lucro Real, a subvenção corretamente escriturada como reserva de incentivos NÃO integra a base de cálculo do IRPJ/CSLL.

- **Relevância:** A matriz da empresa em Santana de Parnaíba (SP) pode ter incentivos fiscais municipais que precisam ser segregados contabilmente no novo modelo.
- **Sinônimos no código:** `subvencao`, `taxIncentive`, `incentivoFiscal`

### Regime Especial

Regime tributário diferenciado concedido a setores ou operações específicas (ex: setor automotivo, Zona Franca de Manaus, exportação). No contexto da reforma, cada regime especial precisa ser mapeado para determinar seu tratamento no IVA Dual (crédito presumido, isenção, alíquota zero, suspensão).

- **Sinônimos no código:** `regimeEspecial`, `specialRegime`

---

## 4. Códigos e Classificações

### NCM — Nomenclatura Comum do Mercosul

Código de 8 dígitos que classifica mercadorias no comércio internacional e nacional. Base para definição de alíquotas de IPI, ICMS, e futuramente IS. Essencial para o mapeamento de produtos sujeitos ao Imposto Seletivo.

- **Formato:** `XXXX.XX.XX` (8 dígitos)
- **Sinônimos no código:** `NCM`, `ncmCode`

### NBS — Nomenclatura Brasileira de Serviços

Código que classifica serviços para fins fiscais. Análogo à NCM para mercadorias, a NBS será utilizada para a classificação de serviços no IBS municipal.

- **Sinônimos no código:** `NBS`, `nbsCode`

### CClassTrib — Código de Classificação Tributária

Novo código unificado criado pela reforma tributária para classificar bens e serviços de forma padronizada nacionalmente, substituindo a multiplicidade de classificações estaduais e municipais atuais.

- **Sinônimos no código:** `CClassTrib`, `cclassTrib`

### CFOP — Código Fiscal de Operações e Prestações

Código numérico de 4 dígitos que identifica a natureza da operação fiscal (venda, devolução, remessa, retorno, etc.) na nota fiscal eletrônica (NF-e).

- **Faixas relevantes para o projeto:**
  - 5.xxx: Operações interestaduais
  - 6.xxx: Operações interestaduais
  - 1.xxx/2.xxx/3.xxx: Operações dentro do estado
- **Sinônimos no código:** `CFOP`, `cfop`, `fiscalOperationCode`

### Código IBGE de Município

Código numérico de 7 dígitos que identifica unicamente cada município brasileiro. Essencial para o cálculo do IBS, pois a alíquota varia conforme o município de destino.

- **Formato:** 7 dígitos (ex: Santana de Parnaíba/SP = `3547304`)
- **Sinônimos no código:** `codigoIBGE`, `ibgeCode`, `municipioIBGE`

### CNAB — Centro Nacional de Automação Bancária

Padrão de arquivo utilizado para comunicação entre empresas e instituições financeiras (cobrança, pagamento, conciliação). Relevante para o split payment: os arquivos CNAB de retorno conterão os valores retidos de CBS e IBS.

- **Sinônimos no código:** `CNAB`, `cnab`, `bankFile`

### CST / CSOSN — Código de Situação Tributária

Códigos que identificam a situação tributária de um produto ou serviço (tributado, isento, substituição tributária, etc.) na nota fiscal. Serão substituídos progressivamente pelo CClassTrib.

- **CST:** Aplicável a empresas do Lucro Real e Lucro Presumido
- **CSOSN:** Aplicável a empresas do Simples Nacional
- **Sinônimos no código:** `CST`, `CSOSN`, `taxSituationCode`

---

## 5. Termos Geográficos e Comerciais

### Princípio do Destino

Regra constitucional pela qual o IBS pertence ao estado e município onde o bem ou serviço é **consumido**, e não onde é produzido ou faturado. Isso inverte a lógica atual do ICMS (origem) e torna a localização geográfica do comprador o dado mais crítico do sistema.

- **Exemplo:** Empresa em Santana de Parnaíba (SP) vende para cliente em Salvador (BA). O IBS da venda pertence a Salvador, não a Santana de Parnaíba.
- **Sinônimos no código:** `principioDestino`, `destinationPrinciple`

### DIFAL — Diferencial de Alíquotas

No modelo atual (ICMS), é a diferença entre a alíquota interestadual e a alíquota interna do estado de destino, cobrada do remetente em vendas interestaduais a consumidor final não contribuinte (B2C). No IVA Dual, o DIFAL é absorvido pelo IBS — o destino já recebe a alíquota integral.

- **Sinônimos no código:** `DIFAL`, `diferencialAliquotas`

### Operação Interestadual

Transação comercial em que o estabelecimento remetente (vendedor) está em um estado diferente do destinatário (comprador). No modelo IVA Dual, operações interestaduais B2B não geram recolhimento na origem — apenas crédito para o adquirente.

- **Sinônimos no código:** `interestadual`, `interstate`

---

## 6. Termos Específicos do Projeto

### Onda 1 — Ativação Comercial

Primeira fase de entrega do programa (Mês 3-4): ativação das regras de simulação e exibição transparente dos novos tributos nos canais de venda (CRM, e-commerce, portal B2B). Foco em simulação, precificação e experiência do cliente.

### Onda 2 — Ativação Financeira

Segunda fase de entrega do programa (Mês 5-6): consolidação dos processos de faturamento, split payment, apropriação de créditos e conciliação financeira no back-office (ERP, tesouraria). Foco em conformidade fiscal e fluxo de caixa.

### Shadow Run

Operação em modo "sombra" durante 2026 com alíquotas de teste reduzidas (CBS 0,9%, IBS 0,1%): os cálculos e processos são executados e validados, mas sem impacto material no caixa da empresa. Serve como prova de conceito operacional antes das alíquotas cheias.

### Microserviço de Cálculo de Impostos (ms-billing-engine-tax-rates)

Serviço responsável por centralizar o cálculo dos tributos do IVA Dual (CBS, IBS, IS) com pipeline de fases sequenciais e paralelas, aplicando as regras corporativas de precificação e creditamento. Fonte única da verdade para qualquer canal de venda ou faturamento.

### Trava Contábil / Trava Comercial

Mecanismo automatizado que impede a continuidade de uma operação (venda ou faturamento) se uma condição de negócio não for satisfeita:
- **Trava Comercial:** Bloqueia a venda se o cadastro do cliente não tiver código IBGE validado
- **Trava Contábil:** Bloqueia a emissão da NF-e se os parâmetros fiscais divergirem da simulação comercial

---

## 7. Siglas e Abreviações

| Sigla | Expansão | Contexto |
|-------|----------|----------|
| **BC** | Base de Cálculo | Valor sobre o qual o tributo incide |
| **B2B** | Business to Business | Operações entre empresas |
| **B2C** | Business to Consumer | Operações para consumidor final |
| **CBS** | Contribuição sobre Bens e Serviços | Tributo federal do IVA Dual |
| **CFOP** | Código Fiscal de Operações e Prestações | Natureza da operação fiscal |
| **CNAB** | Centro Nacional de Automação Bancária | Padrão de arquivo bancário |
| **CSLL** | Contribuição Social sobre o Lucro Líquido | Tributo federal sobre lucro |
| **CST** | Código de Situação Tributária | Situação tributária na NF-e |
| **DIFAL** | Diferencial de Alíquotas | ICMS interestadual B2C |
| **EC** | Emenda Constitucional | Norma jurídica de nível constitucional |
| **IBS** | Imposto sobre Bens e Serviços | Tributo subnacional do IVA Dual |
| **ICMS** | Imposto sobre Circulação de Mercadorias e Serviços | Tributo estadual (em extinção) |
| **IPI** | Imposto sobre Produtos Industrializados | Tributo federal (em extinção) |
| **IRPJ** | Imposto de Renda Pessoa Jurídica | Tributo federal sobre lucro |
| **IS** | Imposto Seletivo | Tributo extrafiscal federal |
| **ISS** | Imposto sobre Serviços de Qualquer Natureza | Tributo municipal (em extinção) |
| **IVA** | Imposto sobre Valor Agregado | Modelo tributário (VAT) |
| **NBS** | Nomenclatura Brasileira de Serviços | Classificação fiscal de serviços |
| **NCM** | Nomenclatura Comum do Mercosul | Classificação fiscal de mercadorias |
| **NF-e** | Nota Fiscal Eletrônica | Documento fiscal eletrônico |
| **NFS-e** | Nota Fiscal de Serviços Eletrônica | Documento fiscal de serviços |
| **PIS** | Programa de Integração Social | Contribuição federal (em extinção) |
| **COFINS** | Contribuição para o Financiamento da Seguridade Social | Contribuição federal (em extinção) |
| **SEFAZ** | Secretaria da Fazenda | Órgão fiscal estadual |
| **SLO** | Service Level Objective | Meta de desempenho do serviço |
| **UF** | Unidade Federativa | Estado brasileiro |

---

## 8. Referências Legais

| Referência | Tema | Link / Fonte |
|------------|------|--------------|
| **EC 132/2023** | Reforma Tributária — institui CBS e IBS, extingue ICMS, ISS, PIS, COFINS, IPI | [Texto constitucional] |
| **Comitê Gestor do IBS** | Definição de alíquotas, regras de creditamento, calendários de transição | [Acompanhamento contínuo] |
| **LC 214/2025** | Lei Complementar que regulamenta o IBS | [Norma infraconstitucional] |
| **Convênio ICMS 115/03** | Obrigações acessórias do ICMS (em extinção) | [CONFAZ] |
| **NT 2025.002** | Nota Técnica da RFB sobre a transição para CBS | `docs-suporte/` |
| **RTC CFC-RFB** | Material de capacitação sobre a Reforma Tributária | `docs-suporte/*.pdf` |

---

## 9. Convenção de Nomenclatura no Código

Para garantir consistência entre o glossário de negócios e a implementação técnica:

| Contexto | Convenção | Exemplo |
|----------|-----------|---------|
| Structs/types em Go | PascalCase em inglês, com termo do glossário como prefixo/sufixo | `CbsCalculator`, `IbsRate`, `TaxCredit` |
| Campos/variáveis | camelCase em inglês | `cbsRate`, `destinationIbge`, `taxBase` |
| Constantes fiscais | PascalCase com prefixo do tributo | `CbsDefaultRate`, `IbsMinRate` |
| Comentários de negócio (`// ponytail:`) | Português, usando o termo exato do glossário | `// ponytail: creditamento CBS — global lock` |
| Mensagens de erro | Português, usando o termo exato do glossário | `"alíquota IBS não encontrada para município IBGE %s"` |
| Logs | Inglês para busca/filtro, com termo do glossário | `"IBS rate lookup failed for IBGE %s"` |

## 10. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-08 | Criação inicial: tributos, regimes, mecanismos, códigos, termos do projeto, siglas, referências legais, convenção de nomenclatura | Time Técnico |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: domain-modeling, agile-ba-practices.*
