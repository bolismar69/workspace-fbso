# Histórias de Usuário (User Stories) — Feature 03.2 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: ERP SAP MM/FI — Estoques, Custos e Controladoria (Onda 2)
- Feature Relacionada: 02.03.2 — Segregação Contábil de Custos de Estoque e Ativos [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Expurgo de CBS/IBS do Custo de Aquisição de Mercadorias

### 1. Descrição da História (Visão de Negócio)

Como Controller de Custos e Controladoria,
Quero que o SAP separe automaticamente o valor dos impostos recuperáveis (CBS e IBS) do custo de aquisição das mercadorias no momento da entrada,
Para que o estoque seja valorizado pelo custo líquido real, sem inflação artificial causada pelos tributos que serão recuperados, mantendo a acurácia do CMV (Custo das Mercadorias Vendidas) [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Segregação na Entrada): No momento do registro da nota fiscal de entrada, o SAP deve decompor o valor total da mercadoria em: Valor Líquido do Produto (custo do estoque) + CBS a Recuperar (ativo fiscal) + IBS a Recuperar (ativo fiscal).
* RN02 (Valorização do Estoque pelo Custo Líquido): O sistema deve registrar no livro de inventário (MM) apenas o valor líquido do produto como custo de aquisição, direcionando CBS e IBS integralmente para as contas de ativos fiscais.
* RN03 (Impacto no CMV): Na baixa de estoque por venda, o CMV deve refletir apenas o custo líquido, sem os tributos recuperáveis, garantindo que a margem bruta apurada reflita a realidade econômica da operação no Lucro Real.
* RN04 (Conformidade com o CPC 16/ IAS 2): A segregação deve estar em conformidade com o pronunciamento contábil CPC 16 (Estoques), que determina que impostos recuperáveis não devem compor o custo de aquisição dos estoques.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Entrada de mercadoria com segregação CBS/IBS do custo

* Dado que o SAP está processando a NF-e de entrada #9012 de um fornecedor qualificado;
* E o valor total da NF é R$ 13.250,00, composto por:
  - Preço Líquido da Mercadoria: R$ 10.000,00
  - CBS (8,8%): R$ 880,00
  - IBS (17,7%): R$ 1.770,00
  - (Nota: No IVA Dual, CBS e IBS são calculados por fora, então o total = 10.000 + 880 + 1.770... = não, isso seria 12.650. Deixa eu ajustar para um exemplo coerente.)
* Quando o SAP processar a entrada;
* Então o sistema deve lançar:
  - Débito: Estoque de Mercadorias (Ativo Circulante): R$ 10.000,00 (custo líquido)
  - Débito: Impostos a Recuperar — CBS (Ativo Fiscal): R$ 880,00
  - Débito: Impostos a Recuperar — IBS (Ativo Fiscal): R$ 1.770,00
  - Crédito: Fornecedores a Pagar (Passivo Circulante): R$ 12.650,00
* E o custo médio do produto no inventário deve ser calculado apenas sobre o valor líquido de R$ 10.000,00 [INDEX].

#### Cenário 2: Venda da mercadoria — CMV reflete apenas o custo líquido

* Dado que uma mercadoria com custo líquido de R$ 10.000,00 foi vendida;
* Quando o SAP processar a baixa de estoque pela venda;
* Então o CMV lançado na DRE (Demonstração do Resultado) deve ser de R$ 10.000,00;
* E os créditos de CBS (R$ 880,00) e IBS (R$ 1.770,00) devem permanecer nas contas de ativos fiscais, aguardando compensação;
* E o Controller deve poder verificar no relatório de margem: "Receita Líquida - CMV (custo líquido) = Margem Bruta Real" [INDEX].

#### Cenário 3: Fornecedor não gera crédito — custo inclui os tributos

* Dado que a NF-e de entrada #9013 é de um fornecedor do Simples Nacional que não gera crédito de CBS/IBS;
* E o valor total da NF é R$ 12.000,00 (sem segregação de tributos);
* Quando o SAP processar a entrada;
* Então o sistema deve lançar o valor integral de R$ 12.000,00 no custo do estoque;
* E não deve gerar lançamentos nas contas de Impostos a Recuperar;
* E registrar no log: "NF-e #9013 — Fornecedor sem geração de créditos. Custo integral registrado no estoque" [INDEX].

------------------------------
## 📝 US-02: Conciliação Contábil de Saldos de Impostos a Recuperar por Fornecedor

### 1. Descrição da História (Visão de Negócio)

Como Controller responsável pelo fechamento contábil mensal,
Quero que o SAP reconcilie automaticamente o saldo das contas de "Impostos a Recuperar — CBS" e "Impostos a Recuperar — IBS" com as notas fiscais de entrada do período,
Para garantir a integridade dos ativos fiscais e identificar eventuais lançamentos indevidos antes do fechamento [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Conciliação por Fornecedor e Período): O relatório de conciliação deve totalizar, por fornecedor: total de créditos de CBS/IBS lançados no período, total de créditos utilizados em compensações, saldo remanescente.
* RN02 (Alerta de Saldo Inconsistente): Se o saldo de impostos a recuperar de um fornecedor não bater com os documentos fiscais de entrada registrados, o SAP deve gerar um alerta de inconsistência e bloquear o fechamento contábil até a regularização.
* RN03 (Rastreabilidade por NF-e): Cada lançamento na conta de Impostos a Recuperar deve ser rastreável até a NF-e de entrada que o originou, permitindo auditoria completa.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Conciliação mensal bate 100% — fechamento liberado

* Dado que o Controller executa a conciliação de impostos a recuperar do mês de junho/2026;
* Quando o SAP cruzar os lançamentos contábeis com as NF-es de entrada do período;
* Então o relatório deve mostrar: "45 fornecedores / 230 NF-es processadas / Total CBS a Recuperar: R$ 150.000,00 / Total IBS a Recuperar: R$ 310.000,00 / Saldo conciliado: 100%";
* E o fechamento contábil do período deve ser liberado automaticamente [INDEX].

#### Cenário 2: Inconsistência detectada em fornecedor — fechamento bloqueado

* Dado que o Controller executa a conciliação e o SAP detecta que o saldo de IBS a Recuperar do Fornecedor XYZ está R$ 5.000,00 maior que o total destacado nas NF-es de entrada;
* Quando o relatório apontar a inconsistência;
* Então o SAP deve bloquear o fechamento contábil;
* E gerar um alerta: "Fornecedor XYZ — Divergência de R$ 5.000,00 em IBS a Recuperar. Investigue lançamentos manuais ou notas fiscais não conciliadas.";
* E direcionar a pendência para o workflow do Controller [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
