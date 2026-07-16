# Matriz de Indicadores de Negócio (KPIs) para Reporte Executivo (CFO)

```text
[ FINANCEIRO / FLUXO DE CAIXA ] ──► [ COMPLIANCE / RISCO FISCAL ] ──► [ EFICIÊNCIA OPERACIONAL ]
    Preservação de Margem               Zero Rejeição na SEFAZ             Acurácia Cadastral (Destino)
    e Crédito de Fornecedores           e Retenções do Split               e Latência nos Canais
```

------------------------------
## 1. Dimensão Financeira e Fluxo de Caixa

### KPI F1: Índice de Aproveitamento de Créditos do IVA Dual

* Objetivo de Negócio: Medir a eficiência da empresa em capturar os créditos de CBS e IBS nas compras do Lucro Real (Vinculado ao Épico 02.03 — Apropriação de Créditos no Lucro Real) [INDEX].
* Fórmula de Cálculo: (Total de Créditos Efetivamente Apropriados no SAP / Total de Imposto Destacado nas Notas de Entrada) * 100
* Meta Recomendada: ≥ 98% das entradas elegíveis.
* Frequência de Reporte: Mensal (no fechamento fiscal).

### KPI F2: Margem de Lucro Líquida por Unidade Federativa (UF)

* Objetivo de Negócio: Avaliar se a precificação "por fora" adotada na Onda 1 está protegendo a rentabilidade frente às variações do IBS de destino [INDEX].
* Fórmula de Cálculo: (Receita Líquida Real por Estado - Custo do Produto ou Serviço por Estado) / Receita Líquida Real por Estado
* Meta Recomendada: Oscilação máxima de ± 0,5% em relação ao orçamento planejado antes da Reforma.
* Frequência de Reporte: Mensal.

### KPI F3: Índice de Retenção Indevida no Split Payment

* Objetivo de Negócio: Monitorar se as instituições financeiras estão retendo valores a maior, especialmente nas operações com benefícios fiscais de Santana de Parnaíba (Vinculado ao Épico 02.02 — Split Payment Bancário) [INDEX].
* Fórmula de Cálculo: Valor Retido pelos Bancos (CNAB) - Valor Devido de CBS/IBS Destacado nas Notas Fiscais
* Meta Recomendada: R$ 0,00 (Zero divergência).
* Frequência de Reporte: Semanal (pelo time de Tesouraria). [1] 

------------------------------
## 2. Dimensão de Compliance e Risco Fiscal

### KPI C1: Taxa de Rejeição de Notas Fiscais Eletrônicas (NF-e / NFS-e)

* Objetivo de Negócio: Medir a consistência matemática do motor de cálculo corporativo no momento do faturamento (Vinculado ao Épico 02.01 — Faturamento Integrado e Consistência SAP) [INDEX].
* Fórmula de Cálculo: (Número de Notas Rejeitadas por Erro de Imposto ou Arredondamento / Total de Notas Transmitidas) * 100
* Meta Recomendada: < 0,1% (Apenas falhas técnicas de comunicação com o Fisco).
* Frequência de Reporte: Diário / Semanal.

### KPI C2: Risco de Perda de Subvenção de Investimento (Matriz)

* Objetivo de Negócio: Garantir o correto direcionamento contábil dos incentivos fiscais de Santana de Parnaíba no SAP para isenção de IRPJ/CSLL [INDEX].
* Fórmula de Cálculo: (Valor Contabilizado na Reserva de Incentivos / Total de Imposto Economizado por Regimes Especiais) * 100
* Meta Recomendada: 100% de conformidade na escrituração.
* Frequência de Reporte: Mensal.

------------------------------
## 3. Dimensão de Eficiência Operacional e Canais Comerciais

### KPI O1: Índice de Higienização Cadastral de Clientes (Acurácia de Destino)

* Objetivo de Negócio: Medir a saúde da base de dados do CRM para garantir que o princípio do destino seja aplicado sem travar a força de vendas (Vinculado ao Épico 01.01 — Qualificação Geográfica e Onboarding CRM) [INDEX].
* Fórmula de Cálculo: (Clientes Ativos com Código IBGE e CEP Validados / Total de Clientes Ativos na Base) * 100
* Meta Recomendada: 100% dos clientes com faturamento ativo nos últimos 12 meses.
* Frequência de Reporte: Quinzenal.

### KPI O2: Impacto na Conversão Comercial por Latência Fiscal

* Objetivo de Negócio: Garantir que o cálculo centralizado e dinâmico do IVA Dual não degrade o tempo de resposta no e-commerce ou CRM (Vinculado ao Épico 01.02 — Conexão à Inteligência Corporativa de Cálculo) [INDEX].
* Fórmula de Cálculo: Tempo de resposta (SLA) da API do microsserviço de cálculo sob estresse de acessos simulados.
* Meta Recomendada: Média inferior a 100ms por requisição.
* Frequência de Reporte: Semanal / Mensal. [2] 

------------------------------
## Template de Dashboard Executivo (Visão Mensal para o CFO)

| Categoria | Indicador de Negócio | Status Atual | Meta | Tendência | Impacto Financeiro Associado |
|---|---|---|---|---|---|
| Financeiro | F1: Aproveitamento de Créditos (Lucro Real) | 94,5% | ≥ 98% | 🔺 Crescente | R$ X de economia potencial no fluxo de caixa. |
| Financeiro | F3: Divergência de Saldo no Split Payment | R$ 12k | R$ 0 | 🔻 Decrescente | Capital de giro retido temporariamente nos bancos. |
| Compliance | C1: Rejeição de Notas (Layout IVA) | 0,8% | < 0,1% | 🔻 Decrescente | Risco de atraso na entrega física e insatisfação. |
| Operações | O1: Higienização Geográfica (CRM) | 88,2% | 100% | 🔺 Crescente | Evita aplicação de alíquota errada de IBS no destino. |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: finance-metrics-quickref, agile-ba-practices.*
