# Histórias de Usuário (User Stories) — Feature 03.3 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: ERP SAP FI — Controladoria e Escrituração Fiscal (Onda 2)
- Feature Relacionada: 02.03.3 — Escrituração de Reserva de Incentivos (Subvenção no Lucro Real) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Identificação e Cálculo do Montante de IBS Economizado por Benefícios Fiscais

### 1. Descrição da História (Visão de Negócio)

Como Controller Tributário da matriz,
Quero que o SAP calcule automaticamente o montante total de IBS que deixou de ser recolhido em cada período devido aos benefícios fiscais aplicados nas operações da matriz de Santana de Parnaíba,
Para que a empresa possa contabilizar corretamente a reserva de incentivos fiscais e garantir a isenção de IRPJ e CSLL sobre esse valor [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Cálculo do IBS Economizado): Para cada operação incentivada, o SAP deve calcular: IBS Economizado = IBS que seria devido sem benefício − IBS efetivamente recolhido com benefício. Esse delta é o ganho fiscal a ser direcionado para a reserva de incentivos.
* RN02 (Segregação por Benefício): O cálculo deve ser segregado por tipo de benefício fiscal aplicado (redução de base de cálculo, alíquota diferenciada, diferimento, isenção), permitindo rastreabilidade para auditoria fiscal.
* RN03 (Acumulador Mensal): O SAP deve acumular o IBS economizado ao longo do mês e disponibilizar o total para a rotina de escrituração da reserva no fechamento contábil.
* RN04 (Conciliação com as NF-es Emitidas): Os valores de IBS economizado devem ser conciliáveis com as informações complementares registradas em cada NF-e emitida com benefício fiscal.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Fechamento mensal — SAP totaliza IBS economizado por benefício

* Dado que a matriz de Santana de Parnaíba emitiu 150 NF-es com benefícios fiscais no mês de junho/2026;
* Quando o Controller executar a rotina de fechamento mensal;
* Então o SAP deve calcular e exibir:
  - "IBS Economizado — Redução de Base de Cálculo (Lei Mun. X): R$ 85.000,00"
  - "IBS Economizado — Alíquota Diferenciada (Regime Y): R$ 42.000,00"
  - "IBS Economizado — Diferimento (Regime Z): R$ 23.000,00"
  - "Total de IBS Economizado no Período: R$ 150.000,00"
* E cada valor deve ser rastreável clicando para ver a lista de NF-es que contribuíram [INDEX].

#### Cenário 2: Operação sem benefício não gera IBS economizado

* Dado que uma NF-e foi emitida pela matriz sem aplicação de qualquer benefício fiscal;
* Quando a rotina de cálculo for executada;
* Então o SAP deve registrar IBS Economizado = R$ 0,00 para aquela operação;
* E a NF-e não deve ser listada no relatório de incentivos fiscais [INDEX].

------------------------------
## 📝 US-02: Transferência Contábil Automática para Reserva de Incentivos (Subvenção para Investimento)

### 1. Descrição da História (Visão de Negócio)

Como Controller responsável pela escrituração fiscal da matriz,
Quero que o SAP realize automaticamente a transferência contábil do IBS economizado da conta de resultado de impostos para a conta de Reserva de Incentivos Fiscais no Patrimônio Líquido,
Para cumprir os requisitos legais que permitem a exclusão desse valor da base de cálculo do IRPJ e da CSLL [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Lançamento Contábil da Reserva): No fechamento mensal, o SAP deve realizar o lançamento:
  - Débito: Despesa de IBS (Resultado) — R$ XXX (reduzindo a despesa de IBS no resultado)
  - Crédito: Reserva de Incentivos Fiscais — Subvenção para Investimento (Patrimônio Líquido) — R$ XXX
* RN02 (Condição para Constituição da Reserva): A transferência só pode ser realizada se a empresa tiver lucro líquido no período. Caso contrário, o valor permanece na conta de resultado e a constituição da reserva é postergada para o próximo período com lucro.
* RN03 (Controle de Limite Legal): O SAP deve controlar o saldo acumulado da reserva de incentivos e alertar se a destinação exceder o limite legal permitido para subvenções para investimento.
* RN04 (Documentação de Suporte): O lançamento contábil deve ser vinculado a um documento interno de "Constituição de Reserva de Incentivos" que compile as NF-es e os benefícios que geraram o IBS economizado no período.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Fechamento com lucro — reserva constituída automaticamente

* Dado que a empresa apurou lucro líquido de R$ 500.000,00 no mês de junho/2026;
* E o total de IBS economizado por benefícios fiscais no período foi de R$ 150.000,00;
* Quando o Controller executar o fechamento contábil;
* Então o SAP deve realizar automaticamente o lançamento:
  - Débito: Despesa de IBS (Conta de Resultado): R$ 150.000,00
  - Crédito: Reserva de Incentivos Fiscais — Subvenção para Investimento (PL): R$ 150.000,00
* E gerar o documento interno "Reserva de Incentivos — Junho/2026" vinculando as 150 NF-es que geraram o benefício;
* E registrar: "Reserva constituída conforme Lei Complementar XYZ. IRPJ e CSLL não incidem sobre este montante." [INDEX].

#### Cenário 2: Fechamento com prejuízo — reserva não constituída

* Dado que a empresa apurou prejuízo líquido no mês de junho/2026;
* E o IBS economizado foi de R$ 150.000,00;
* Quando o Controller executar o fechamento contábil;
* Então o SAP deve detectar que não há lucro para absorver a constituição da reserva;
* E manter o valor na conta de resultado de IBS;
* E gerar um alerta informativo: "Constituição da reserva de incentivos postergada — empresa sem lucro líquido no período.";
* E o valor de R$ 150.000,00 deve ser acumulado para constituição no próximo período com lucro [INDEX].

------------------------------
## 📝 US-03: Relatório de Conformidade para Auditoria Fiscal e Blindagem do IRPJ/CSLL

### 1. Descrição da História (Visão de Negócio)

Como CFO e Controller Tributário,
Quero que o SAP gere um relatório anual de conformidade documentando toda a movimentação da Reserva de Incentivos Fiscais,
Para apresentar à auditoria externa e à Receita Federal a correta segregação dos incentivos, blindando a empresa contra autuações de IRPJ e CSLL sobre os valores subvencionados [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Trilha de Auditoria Completa): O relatório deve conter: (a) saldo inicial da reserva no período, (b) lista de todas as constituições mensais com os respectivos documentos de suporte, (c) valores utilizados/transferidos da reserva, (d) saldo final, (e) total de IRPJ e CSLL não tributados em virtude da subvenção.
* RN02 (Formato para ECF — Escrituração Contábil Fiscal): O relatório deve ser exportável no formato compatível com a ECF (Escrituração Contábil Fiscal) da Receita Federal, facilitando o preenchimento da declaração anual.
* RN03 (Assinatura de Responsabilidade): O relatório gerado deve conter campos para assinatura digital do Controller e do CFO, atestando a veracidade das informações prestadas ao Fisco.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Geração do relatório anual para auditoria externa

* Dado que a auditoria externa solicitou o relatório de incentivos fiscais do exercício de 2026;
* Quando o Controller acessar a funcionalidade de relatórios fiscais e selecionar "Relatório de Reserva de Incentivos — 2026";
* Então o SAP deve gerar um PDF contendo:
  - Saldo Inicial em 01/01/2026: R$ 0,00
  - Constituições mensais: Jan (não houve) ... Jun (R$ 150.000,00) ... Dez (a definir)
  - Saldo Final em 31/12/2026: R$ XXX
  - IRPJ/CSLL não tributados: R$ XXX (34% sobre o saldo)
* E disponibilizar exportação no formato ECF (XML);
* E incluir os campos de assinatura digital para Controller e CFO [INDEX].

#### Cenário 2: Relatório vazio para período sem benefícios

* Dado que o Controller solicita o relatório de um mês em que não houve operações incentivadas;
* Quando o SAP processar a consulta;
* Então o sistema deve gerar um relatório informando: "Período sem constituição de reserva de incentivos — nenhuma operação incentivada identificada.";
* E o saldo da reserva deve permanecer inalterado [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
