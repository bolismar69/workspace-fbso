# Detalhamento de Features da Onda 02: Finanças, Faturamento e ERP
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Fase: Onda 2 (Sistemas de Finanças, Escrituração, Faturamento e Governança SAP)
- Épico: 02.01 (Faturamento Integrado e Consistência de Emissão SAP), 02.02 (Operação e Governança do Split Payment), 02.03 (Otimização de Custos e Apropriação de Créditos)
- Status: Pronto para Detalhamento de User Stories
- Responsáveis: Product Managers (PMs) de Finanças/Controladoria e POs de ERP/SAP

------------------------------
## 🧾 ÉPICO 02.01: Faturamento Integrado e Consistência de Emissão (SAP)

### Feature 02.01.1: Validação de Faturamento Pré-Emissão e Trava Contábil

* Objetivo de Negócio: Eliminar o risco de rejeição de Notas Fiscais na SEFAZ por divergência de valores e garantir a conformidade do faturamento no Lucro Real.
* Descrição Funcional: Mecanismo automatizado de auditoria interna executado no módulo SAP SD (Vendas e Distribuição) antes da transmissão do documento fiscal (NF-e ou NFS-e nacional). O sistema deve confrontar as alíquotas e o cálculo do IVA Dual enviados pela frente de vendas com o cálculo atualizado da inteligência corporativa.
* Regras de Negócio e Restrições:
* Se houver qualquer divergência matemática de centavos (decorrente de regras de arredondamento) ou alteração de alíquota na cidade de destino, o SAP deve bloquear a emissão da nota automaticamente.
   * O documento entra em status "Bloqueio Fiscal" e uma tarefa é gerada para a Controladoria avaliar e autorizar o ajuste da cotação ou a absorção da diferença na margem da empresa.

### Feature 02.01.2: Motor de Conversão do ISS para IBS (Serviços Santana de Parnaíba)

* Objetivo de Negócio: Adaptar a operação da matriz para o fim do imposto municipal tradicional (ISS) e início do IBS de serviços baseado no destino.
* Descrição Funcional: Parametrização das regras de determinação de impostos no SAP para faturamento de serviços emitidos pela matriz em Santana de Parnaíba (SP). O sistema deve desativar a retenção de ISS com base na competência da cidade de origem e ativar a aplicação das novas alíquotas de IBS do local de consumo do cliente.
* Regras de Negócio e Restrições:
* O sistema deve descontinuar a emissão de notas via sistema legado da prefeitura de Santana de Parnaíba e integrar o fluxo diretamente ao padrão da NFS-e Nacional.

### Feature 02.01.3: Automação de Benefícios e Regimes Especiais (Transição Santana de Parnaíba)

* Objetivo de Negócio: Blindar e aplicar os incentivos fiscais e regimes especiais vigentes da matriz sem intervenção manual do faturista.
* Descrição Funcional: Inteligência de negócio mapeada nas tabelas de condições de impostos do SAP. Ao identificar que a operação é elegível a um regime especial ativo da matriz, o sistema aplica reduções de base de cálculo ou alíquotas diferenciadas sobre o IBS, conforme as regras de transição permitidas pela legislação.
* Regras de Negócio e Restrições:
* O sistema deve validar se o benefício fiscal local permanece aplicável em vendas interestaduais ou se a regra do destino (alíquota cheia do estado/município comprador) anula o incentivo de origem.

------------------------------
## 💸 ÉPICO 02.02: Operação e Governança do Mecanismo de Split Payment (Tesouraria)

### Feature 02.02.1: Liquidação e Conciliação Financeira Segregada (Split Bancário)

* Objetivo de Negócio: Adaptar as contas a receber para a retenção automática do imposto na fonte pelas instituições financeiras.
* Descrição Funcional: Reformulação do fluxo de caixa e do processamento de arquivos de retorno bancário (CNAB) no módulo SAP FI (Finanças). O sistema deve reconhecer que os pagamentos efetuados pelos clientes via boleto, PIX ou cartão de crédito entrarão na conta da empresa já descontados os valores de CBS e IBS [INDEX].
* Regras de Negócio e Restrições:
* Ao processar a baixa do título de um cliente, o SAP deve realizar um lançamento duplo automático: dar baixa total na duplicata do cliente, direcionar o valor líquido para a conta corrente principal e lançar a parcela do imposto retido em uma conta de compensação fiscal de "Impostos Retidos na Fonte - Split".

### Feature 02.02.2: Ajuste de Split para Operações Incentivadas

* Objetivo de Negócio: Evitar que os bancos retenham impostos a maior na fonte em notas emitidas com benefícios fiscais pela matriz.
* Descrição Funcional: Integração de dados entre o faturamento e as instruções de cobrança bancária. O sistema deve gerar os metadados do boleto ou PIX contendo as tags exatas do valor recalculado do CBS/IBS (após o benefício fiscal aplicado pela Feature 02.01.3).
* Regras de Negócio e Restrições:
* Garante que a rede bancária processe o split apenas sobre o valor tributário reduzido/incentivado, protegendo o caixa líquido e o capital de giro da empresa na matriz.

### Feature 02.02.3: Painel de Auditoria e Reconciliação do Split Payment

* Objetivo de Negócio: Identificar falhas de retenção ou desvios cometidos pelas instituições financeiras.
* Descrição Funcional: Relatório de conferência e auditoria interna na tesouraria. O sistema deve cruzar o valor de CBS/IBS destacado na Nota Fiscal (visão faturamento) com o valor efetivamente retido e informado pelo banco no arquivo de liquidação (visão tesouraria).
* Regras de Negócio e Restrições:
* O painel deve apontar discrepâncias em um relatório de exceções (ex: "Nota Fiscal 123 - Banco reteve R$ 10,00 a maior"). Qualquer erro gera uma notificação automatizada para o time de contas a receber acionar a instituição financeira parceira.

------------------------------
## ⚖️ ÉPICO 02.03: Otimização de Custos e Apropriação de Créditos (Suprimentos e Controladoria)

### Feature 02.03.1: Auditoria Fiscais de Entrada e Bloqueio de Créditos (Procure-to-Pay)

* Objetivo de Negócio: Garantir que a empresa só se aproprie de créditos legítimos no Lucro Real, evitando riscos de autuação por fornecedores irregulares.
* Descrição Funcional: Mecanismo de checagem automatizada no recebimento físico e fiscal de mercadorias no módulo SAP MM (Compras). O sistema deve condicionar a liberação do crédito de CBS/IBS à validação de que o imposto destacado pelo fornecedor foi devidamente processado e recolhido pelo mecanismo de Split Payment na origem.
* Regras de Negócio e Restrições:
* Se o fornecedor for optante por regimes simplificados ou houver inconformidade na guia de recolhimento, o SAP impede o lançamento do imposto na conta de "Impostos a Recuperar", direcionando o valor temporariamente para uma conta de "Créditos em Análise".

### Feature 02.03.2: Segregação Contábil de Custos de Estoque e Ativos

* Objetivo de Negócio: Refletir o custo real dos produtos estocados nos armazéns, expurgando o impacto dos impostos recuperáveis.
* Descrição Funcional: Ajuste na inteligência contábil de avaliação de inventário no SAP. O sistema deve extrair o valor total do CBS e IBS calculados por fora e direcioná-los integralmente para as contas de ativos fiscais, registrando no custo de estoque do produto puramente o seu valor líquido.
* Regras de Negócio e Restrições:
* Esta segregação é vital para manter a acurácia do Custo das Mercadorias Vendidas (CMV) e evitar a inflação artificial dos estoques da empresa em nível nacional.

### Feature 02.03.3: Escrituração de Reserva de Incentivos (Subvenção no Lucro Real)

* Objetivo de Negócio: Proteger o ganho financeiro dos regimes especiais de Santana de Parnaíba contra a tributação do IRPJ e da CSLL.
* Descrição Funcional: Rotina de escrituração contábil acionada no fechamento fiscal mensal. O sistema deve identificar o montante total de IBS que deixou de ser recolhido em virtude dos benefícios da matriz e realizar uma transferência contábil interna.
* Regras de Negócio e Restrições:
* O valor economizado deve ser retirado da conta de resultado de impostos e creditado em uma conta de patrimônio líquido dedicada: "Reserva de Incentivos Fiscais - Subvenção para Investimento", cumprindo os requisitos legais de isenção de imposto de renda sobre o benefício.

------------------------------
### 5. Critérios de Sucesso para Homologação da Onda 2 (Definition of Done - DoD)
O comitê financeiro considerará as Features da Onda 2 prontas para produção quando:

   1. Consistência Contábil: O balancete de testes do SAP demonstrar que os impostos gerados pelo faturamento interestadual foram destinados às contas corretas por estado/município de destino sem erros de arredondamento.
   2. Baixa Automatizada: O sistema realizar a baixa de 100% dos títulos de teste no banco através do fluxo de Split Payment, liquidando as contas de compensação sem intervenção manual.
   3. Blindagem de Santana de Parnaíba: A auditoria contábil validar que os lançamentos dos benefícios fiscais da matriz foram devidamente isolados nas contas de reserva de subvenção, protegendo o Lucro Real.
