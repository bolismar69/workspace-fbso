# Histórias de Usuário (User Stories) — Feature 02.1 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: ERP SAP FI — Tesouraria e Contas a Receber (Onda 2)
- Feature Relacionada: 02.02.1 — Liquidação e Conciliação Financeira Segregada (Split Bancário) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Processamento de Retorno Bancário com Identificação de Split Payment

### 1. Descrição da História (Visão de Negócio)

Como Analista de Tesouraria responsável pela conciliação bancária diária,
Quero que o SAP processe os arquivos de retorno bancário (CNAB) identificando automaticamente a parcela de CBS e IBS retida pelo banco na fonte via Split Payment,
Para que a baixa dos títulos de clientes seja realizada corretamente, segregando a receita líquida da empresa dos valores retidos e destinados aos entes públicos [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Lançamento Duplo Automático na Baixa): Ao processar o pagamento de um título, o SAP deve realizar dois lançamentos contábeis simultâneos: (a) baixa do título do cliente pelo valor total (líquido + impostos), (b) direcionamento da parcela de CBS e IBS retida para a conta de compensação "Impostos Retidos na Fonte — Split Payment".
* RN02 (Identificação via Tags no Arquivo CNAB): O sistema deve identificar as parcelas de CBS e IBS no arquivo de retorno bancário por meio das tags/metadados de split payment enviadas na instrução de cobrança (boleto, PIX, cartão).
* RN03 (Conciliação por Documento Fiscal): Cada lançamento de split deve ser vinculado ao número da Nota Fiscal de origem, permitindo rastreabilidade completa entre o documento emitido, o valor retido e o título liquidado.
* RN04 (Tratamento de Divergência de Retenção): Se o valor retido pelo banco for diferente do valor destacado na Nota Fiscal, o SAP deve lançar a diferença em uma conta de "Variação de Split Payment a Conciliar" e gerar uma pendência para o time de Tesouraria.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Pagamento via boleto com split — baixa processada corretamente

* Dado que um cliente pagou um boleto de R$ 12.650,00 referente à NF-e #5001;
* E a NF-e #5001 possui CBS destacado de R$ 880,00 e IBS de R$ 1.770,00;
* E o banco processou o split e reteve R$ 2.650,00 (CBS + IBS);
* Quando o SAP processar o arquivo de retorno bancário CNAB;
* Então o sistema deve identificar as tags de split e realizar os lançamentos:
  - Débito: Banco Conta Movimento — R$ 10.000,00 (valor líquido recebido)
  - Débito: Impostos Retidos na Fonte — Split (CBS) — R$ 880,00
  - Débito: Impostos Retidos na Fonte — Split (IBS) — R$ 1.770,00
  - Crédito: Clientes — R$ 12.650,00 (baixa total do título)
* E vincular os lançamentos de split à NF-e #5001 [INDEX].

#### Cenário 2: Divergência entre valor retido e valor destacado na NF

* Dado que a NF-e #5002 possui IBS destacado de R$ 1.770,00;
* E o banco reteve R$ 1.800,00 (R$ 30,00 a maior);
* Quando o SAP processar o retorno bancário e detectar a divergência;
* Então o sistema deve lançar:
  - Débito: Impostos Retidos na Fonte — Split (IBS) — R$ 1.770,00 (valor correto da NF)
  - Débito: Variação de Split Payment a Conciliar — R$ 30,00 (diferença a investigar)
* E gerar automaticamente uma pendência no workflow da Tesouraria: "NF-e #5002 — Banco reteve R$ 30,00 a maior de IBS. Acionar instituição financeira para estorno.";
* E bloquear o encerramento contábil do período até a resolução da pendência [INDEX].

------------------------------
## 📝 US-02: Conciliação Automatizada de Títulos com Split Payment

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Tesouraria,
Quero que o SAP reconcilie automaticamente os títulos de cobrança com os pagamentos recebidos via split payment,
Para eliminar o processamento manual de centenas de baixas diárias e garantir o fechamento contábil tempestivo [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Motor de Conciliação Automática): O SAP deve cruzar automaticamente os pagamentos recebidos no extrato bancário com os títulos em aberto no contas a receber, utilizando como chave de conciliação: CNPJ do pagador, valor total do título, número da NF-e e ID do split payment.
* RN02 (Tolerância de Data de Pagamento): Se o pagamento for identificado com atraso de até 3 dias úteis em relação ao vencimento, a conciliação automática deve considerar os encargos financeiros padrão (multa e juros) e lançá-los em contas separadas de receita financeira.
* RN03 (Pendências Não Conciliadas Automaticamente): Pagamentos que não puderem ser conciliados automaticamente em até 24 horas devem ser direcionados para a fila de trabalho manual do analista de tesouraria, com as possíveis causas sugeridas pelo sistema.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Conciliação automática bem-sucedida de lote de boletos

* Dado que o SAP recebeu o arquivo de retorno bancário com 200 pagamentos de clientes;
* E todos os pagamentos possuem as tags de split payment e referenciam NF-es válidas no sistema;
* Quando o motor de conciliação automática for executado;
* Então o sistema deve conciliar 100% dos pagamentos com seus respectivos títulos;
* E gerar o relatório de conciliação: "200 títulos processados — 200 conciliados automaticamente — 0 pendências";
* E o tempo total de processamento do lote não deve exceder 5 minutos [INDEX].

#### Cenário 2: Pagamento não conciliado — encaminhado para análise manual

* Dado que um pagamento de R$ 5.000,00 foi identificado no extrato;
* E não há nenhum título em aberto com o CNPJ do pagador e valor compatível;
* Quando o motor de conciliação automática for executado;
* Então o sistema deve classificar o pagamento como "Não Conciliado — Aguardando Análise";
* E sugerir possíveis causas: "Pagamento sem título correspondente", "CNPJ do pagador difere do CNPJ do sacado", "Valor divergente";
* E direcionar para a fila de trabalho do analista de tesouraria [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
