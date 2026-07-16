# Lista de Épicos da Onda 2: Finanças, Faturamento e ERP
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Fase: Onda 2 (Sistemas de Finanças, Escrituração, Faturamento e Governança SAP)
- Status: Pronto para Refinamento Funcional
- Responsáveis: Product Managers (PMs) de Finanças/Controladoria e POs de ERP/SAP

------------------------------
## 1. Objetivo do Documento

Este documento define os Épicos (Grandes Blocos de Entrega) sob a perspectiva de negócios, necessários para capacitar as áreas de Faturamento, Controladoria, Tesouraria e Suprimentos no ERP SAP a processar a liquidação, escrituração e apropriação de créditos do IVA Dual (CBS, IBS e IS), atendendo aos requisitos estabelecidos no REQUIREMENTS.md.

------------------------------
## 2. Visão Geral do Back-Office (Onda 2)

```text
[ EMISSÃO E CONFIRMAÇÃO ] ───────► [ GESTÃO DE CAIXA / RETENÇÃO ] ───────► [ ESCRITURAÇÃO E ATIVOS ]
       Épico 01:                              Épico 02:                               Épico 03:
 Faturamento e Consistência                Mecanismo de Split                  Apropriação de Créditos
      Matemática (SAP)                     Payment Bancário                     no Lucro Real (SAP)
```

------------------------------
## 3. Detalhes dos Épicos de Negócio

### 🧾 ÉPICO 01: Faturamento Integrado e Consistência de Emissão (SAP)

* Descrição de Alto Nível: Garantir que o processo de faturamento físico e de serviços (matriz em Santana de Parnaíba) envie dados de cálculo idênticos aos simulados em vendas para a emissão dos novos layouts de documentos fiscais, eliminando rejeições regulatórias e garantindo a aplicação de incentivos vigentes.
* Justificativa de Negócio: No regime de Lucro Real, divergências de centavos ou inconsistências nas regras de transição federativas entre a proposta comercial e a nota emitida travam a expedição de mercadorias e geram risco de autuação fiscal de faturamento.
* Requisitos Vinculados (REQUIREMENTS.md): BR-07 (Unicidade Matemática entre Pedido e Nota Fiscal).
* Capacidades Esperadas (Features do Produto):
* Validação de Faturamento Pré-Emissão: Mecanismo de negócio que bloqueia a geração da nota fiscal se o valor calculado na liquidação divergir das premissas de preço e alíquotas de destino fechadas em vendas.
   * Conversão do ISS em IBS (Matriz): Regra de faturamento dedicada para a matriz em Santana de Parnaíba, descontinuando o modelo de retenção de ISS na competência municipal e ativando a cobrança do IBS de serviços por destino.
   * Governança de Benefícios Fiscais e Regimes Especiais (Santana de Parnaíba): Inteligência de negócio que identifica transações elegíveis a incentivos fiscais locais (como reduções de base de cálculo, alíquotas diferenciadas ou diferimentos). Garante a aplicação dessas exceções nas operações originadas na matriz, aplicando as regras de transição estipuladas pela reforma tributária de forma automatizada.

### 💸 ÉPICO 02: Operação e Governança do Mecanismo de Split Payment (Tesouraria)

* Descrição de Alto Nível: Adaptar o processo de contas a receber e tesouraria para suportar a segregação automatizada do fluxo de caixa no ato do pagamento do cliente, destinando os impostos diretamente aos entes públicos e o valor líquido para a empresa, considerando reduções por incentivos.
* Justificativa de Negócio: Esta é a mudança operacional mais crítica da reforma. A empresa precisa gerenciar o fluxo de caixa sabendo que o dinheiro do imposto (CBS e IBS) será retido na fonte pela rede bancária, alterando radicalmente as regras de conciliação e concorrência de saldos.
* Requisitos Vinculados (REQUIREMENTS.md): BR-09 (Viabilização do Mecanismo de Split Payment).
* Capacidades Esperadas (Features do Produto):
* Conciliação Financeira Segregada (Split): Processo de liquidação capaz de baixar títulos de cobrança (boletos, cartões, PIX) identificando a parcela líquida recebida e a parcela de imposto retida pelo banco.
   * Split Reduzido para Regimes Especiais: Parametrização financeira que calcula a instrução de retenção bancária com base no valor incentivado real do documento fiscal emitido por Santana de Parnaíba, mitigando o risco de retenção a maior na fonte e protegendo o caixa líquido.
   * Auditoria Bancária do IVA: Relatório de controle financeiro para cruzar o valor retido na fonte pela instituição bancária com o valor devido de CBS/IBS calculado pelo sistema.

### ⚖️ ÉPICO 03: Otimização de Custos e Apropriação de Créditos (Suprimentos e Controladoria)

* Descrição de Alto Nível: Estruturar as regras de entrada de mercadorias e serviços no SAP para mapear, auditar e apropriar 100% dos créditos não cumulativos do IVA Dual gerados pela cadeia de fornecedores da empresa nacionalmente.
* Justificativa de Negócio: No Lucro Real, a eficiência de margem depende da captura integral dos créditos. Compras efetuadas de fornecedores que não repassarem o imposto (devido ao Split Payment ou regimes especiais) não geram crédito, o que encarece o custo do produto e penaliza a rentabilidade.
* Requisitos Vinculados (REQUIREMENTS.md): BR-08 (Rastreabilidade de Créditos no Lucro Real).
* Capacidades Esperadas (Features do Produto):
* Auditoria de Crédito na Entrada (Procure-to-Pay): Regra de negócio na escrituração fiscal que condiciona o aproveitamento do crédito de CBS/IBS à validação da regularidade e recolhimento do imposto por parte do fornecedor.
   * Segregação de Custo e Tributo Recuperável: Atualização da inteligência contábil de estoques para separar o valor do imposto recuperável do custo real do produto estocado em todos os armazéns do país.
   * Contabilização de Reserva de Incentivos (Lucro Real): Mecanismo de escrituração contábil que isola e direciona o ganho financeiro obtido pelos regimes especiais de Santana de Parnaíba para contas específicas de subvenção para investimento, mantendo a blindagem de exclusão da base de cálculo do IRPJ e da CSLL.

------------------------------
## 4. Critérios de Sucesso para Fechamento da Onda 2 (Financeira)
As diretorias de Controladoria e Finanças considerarão esta lista de Épicos concluída e o programa pronto para encerramento quando:

   1. Divergência Zero: 100% dos livros fiscais gerados no SAP baterem centavo por centavo com os documentos fiscais emitidos nacionalmente.
   2. Fluxo de Caixa Conciliado: O processo de conciliação bancária de recebíveis rodar de forma automatizada sob o modelo de Split Payment sem gerar pendências inexplicáveis de saldo, inclusive nas notas fiscais faturadas com alíquotas incentivadas.
   3. Margem de Lucro Real Protegida: A empresa comprovar, no primeiro fechamento fiscal, que capturou todos os créditos tributários previstos em lei sobre as compras efetuadas e registrou os incentivos da matriz conforme as exigências regulatórias vigentes.
