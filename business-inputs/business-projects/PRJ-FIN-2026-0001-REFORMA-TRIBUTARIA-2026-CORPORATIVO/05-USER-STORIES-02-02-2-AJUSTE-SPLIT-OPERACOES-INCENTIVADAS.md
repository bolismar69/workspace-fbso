# Histórias de Usuário (User Stories) — Feature 02.2 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulos: ERP SAP SD/FI — Faturamento, Cobrança e Tesouraria (Onda 2)
- Feature Relacionada: 02.02.2 — Ajuste de Split para Operações Incentivadas [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Geração de Instrução de Cobrança com Valor de Split Reduzido (Benefício Fiscal)

### 1. Descrição da História (Visão de Negócio)

Como Analista de Faturamento e Cobrança,
Quero que o SAP gere as instruções de cobrança bancária (boleto, PIX, cartão) com as tags de split payment refletindo o valor real do CBS e IBS após a aplicação dos benefícios fiscais da matriz,
Para que o banco retenha apenas o valor correto e reduzido do imposto, protegendo o capital de giro da empresa e evitando retenção a maior na fonte [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Metadados de Split Após Benefício): As tags de split payment enviadas na instrução de cobrança devem conter os valores de CBS e IBS já reduzidos pelos benefícios fiscais aplicados na Feature 01.3, e não os valores cheios que seriam devidos sem o benefício.
* RN02 (Validação Pré-Geração do Boleto): Antes de gerar o boleto ou instrução PIX, o SAP deve executar uma validação cruzada: "Valor do Split (CBS + IBS) ≤ Valor Total do Título - Valor Líquido da Empresa". Se a soma do split for maior que o devido, a geração deve ser bloqueada.
* RN03 (Campos Informativos no Boleto): O boleto deve conter no campo de instruções ao sacado uma linha informativa sobre o benefício aplicado: "Valores de tributos reduzidos conforme regime especial da matriz."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Boleto gerado com split reduzido após benefício fiscal

* Dado que a NF-e #6001 foi emitida pela matriz de Santana de Parnaíba com benefício fiscal de redução de 40% na base de cálculo do IBS;
* E o IBS devido após benefício é R$ 1.062,00 (em vez de R$ 1.770,00 cheio);
* E o valor total da NF é R$ 11.942,00;
* Quando o SAP gerar a instrução de cobrança bancária (boleto);
* Então as tags de split payment devem conter: CBS = R$ 880,00 / IBS = R$ 1.062,00;
* E o valor líquido a receber pela empresa deve ser R$ 10.000,00;
* E a instrução ao banco deve solicitar a retenção de apenas R$ 1.942,00 (CBS + IBS reduzido);
* E o boleto deve conter a linha informativa sobre o benefício fiscal aplicado [INDEX].

#### Cenário 2: Bloqueio de geração de boleto com split inconsistente

* Dado que o SAP está prestes a gerar um boleto para a NF-e #6002;
* E por erro de parametrização, as tags de split indicam CBS + IBS = R$ 5.000,00;
* E o valor total da NF é R$ 12.650,00, resultando em valor líquido para a empresa de apenas R$ 7.650,00 (quando deveria ser R$ 10.000,00);
* Quando o sistema executar a validação pré-geração;
* Então o SAP deve detectar a inconsistência: "Split (R$ 5.000,00) incompatível com o valor líquido esperado (R$ 10.000,00)";
* E bloquear a geração do boleto;
* E notificar o Analista de Faturamento para revisar as condições de imposto aplicadas [INDEX].

------------------------------
## 📝 US-02: Conciliação do Split com Benefício — Garantia de que o Banco Reterá o Valor Correto

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Tesouraria,
Quero que o SAP compare o valor de split que o banco efetivamente reteve com o valor correto (pós-benefício) destacado na Nota Fiscal,
Para identificar imediatamente se o banco reteve imposto a maior e acionar a instituição financeira para correção [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Conferência Tripla): A reconciliação do split para operações incentivadas deve cruzar três fontes: (a) valor de CBS/IBS destacado na NF-e (pós-benefício), (b) valor enviado nas tags de split da instrução de cobrança, (c) valor efetivamente retido pelo banco no arquivo de retorno CNAB.
* RN02 (Alerta de Retenção Indevida): Se o banco reter valor superior ao destacado na NF-e (ignorando as tags de split), o SAP deve gerar um alerta de severidade "Alta" e uma pendência de estorno no workflow da Tesouraria com SLA de 24 horas.
* RN03 (Relatório Gerencial de Eficiência do Split): O SAP deve gerar um relatório mensal para o CFO com o total de imposto que deixou de ser retido na fonte em virtude dos benefícios fiscais, demonstrando o ganho de capital de giro obtido.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Banco reteve o valor correto (split reduzido) — conciliação OK

* Dado que a NF-e #6001 possui IBS pós-benefício de R$ 1.062,00;
* E as tags de split enviadas ao banco indicavam IBS = R$ 1.062,00;
* Quando o arquivo de retorno CNAB mostrar que o banco reteve exatamente R$ 1.062,00 de IBS;
* Então a conferência tripla deve retornar "OK — sem divergências";
* E o relatório gerencial deve registrar o ganho de capital de giro: "Benefício fiscal gerou redução de R$ 708,00 na retenção de IBS nesta operação" [INDEX].

#### Cenário 2: Banco reteve valor cheio (ignorou split reduzido) — alerta gerado

* Dado que a NF-e #6003 possui IBS pós-benefício de R$ 1.062,00;
* Mas o banco reteve R$ 1.770,00 (valor cheio, ignorando o benefício);
* Quando o SAP detectar a divergência de R$ 708,00 na conferência tripla;
* Então o sistema deve gerar alerta "Alta — Retenção Indevida de IBS";
* E criar pendência no workflow: "NF-e #6003 — Banco reteve R$ 708,00 a maior de IBS. Benefício fiscal não reconhecido pela instituição financeira. Solicitar estorno.";
* E bloquear o fechamento contábil daquele título até a resolução [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
