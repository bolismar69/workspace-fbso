# Histórias de Usuário (User Stories) — Feature 01.1 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: ERP SAP SD — Faturamento e Distribuição (Onda 2)
- Feature Relacionada: 02.01.1 — Validação de Faturamento Pré-Emissão e Trava Contábil [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Bloqueio Automático de Emissão de Nota Fiscal por Divergência de Cálculo

### 1. Descrição da História (Visão de Negócio)

Como Analista de Faturamento do SAP,
Quero que o sistema bloqueie automaticamente a emissão de uma Nota Fiscal eletrônica (NF-e/NFS-e) quando houver qualquer divergência matemática entre o valor simulado na venda e o cálculo atualizado no momento do faturamento,
Para eliminar o risco de rejeição pela SEFAZ e garantir conformidade centavo por centavo no Lucro Real [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Confronto Obrigatório Pré-Emissão): Antes de transmitir qualquer documento fiscal, o SAP SD deve executar uma rotina de validação que confronta: (a) alíquotas e valores de CBS/IBS/IS do pedido de venda (originados da Onda 1), com (b) alíquotas e valores recalculados pela inteligência corporativa no momento do faturamento.
* RN02 (Tolerância Zero para Divergência): Não é admitida qualquer divergência de centavos entre os valores simulados e os valores de faturamento. Até diferenças de arredondamento (R$ 0,01) devem ser tratadas como bloqueio.
* RN03 (Token de Garantia Prevalece): Se o pedido possuir um Token de Garantia Fiscal (Feature 03.3 da Onda 1) válido, o SAP deve herdar os valores do token e ignorar o recálculo atualizado, respeitando o congelamento de 48 horas.
* RN04 (Status de Bloqueio Fiscal): Documentos bloqueados recebem o status "Bloqueio Fiscal — Aguardando Análise da Controladoria" e não podem ser transmitidos até a deliberação do analista responsável.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Pedido sem divergência — faturamento liberado automaticamente

* Dado que o SAP SD recebeu uma ordem de faturamento do pedido #45678 originado do CRM;
* E o pedido contém os valores simulados: Preço Líquido R$ 10.000,00 / CBS R$ 880,00 / IBS R$ 1.770,00;
* Quando a rotina de validação pré-emissão for executada e confrontar com o recálculo do motor corporativo;
* E o recálculo retornar exatamente os mesmos valores (R$ 880,00 e R$ 1.770,00);
* Então o SAP deve liberar a emissão da NF-e automaticamente;
* E registrar no log: "Pedido #45678 — Validação Pré-Emissão OK. Divergência: R$ 0,00" [INDEX].

#### Cenário 2: Pedido com divergência de IBS — bloqueio fiscal acionado

* Dado que o SAP recebeu o pedido #45679 com IBS simulado de R$ 1.770,00 para Belo Horizonte (MG);
* E a alíquota municipal de IBS foi atualizada pelo Comitê Gestor entre a simulação e o faturamento;
* Quando o recálculo corporativo retornar IBS de R$ 1.810,00 (diferença de R$ 40,00);
* Então o SAP deve bloquear a emissão da NF-e;
* E atribuir o status "Bloqueio Fiscal — Aguardando Análise da Controladoria";
* E gerar uma tarefa no workflow da Controladoria com os detalhes da divergência: "IBS Simulado: R$ 1.770,00 / IBS Recalculado: R$ 1.810,00 / Diferença: R$ 40,00";
* E notificar o Analista de Faturamento via alerta no SAP [INDEX].

#### Cenário 3: Pedido com Token de Garantia válido — bloqueio não se aplica

* Dado que o pedido #45680 possui um Token de Garantia Fiscal válido (expira em 24h);
* E a alíquota de IBS do destino foi alterada após a geração do token;
* Quando a rotina de validação pré-emissão for executada;
* Então o SAP deve detectar a presença do token válido;
* E herdar os valores congelados do token, ignorando o recálculo atualizado;
* E liberar a emissão da NF-e com os valores originais da proposta comercial [INDEX].

------------------------------
## 📝 US-02: Workflow de Análise e Liberação de Documentos Bloqueados

### 1. Descrição da História (Visão de Negócio)

Como Analista de Controladoria responsável pela revisão fiscal,
Quero receber uma notificação no meu workflow do SAP sempre que um documento for bloqueado por divergência de cálculo,
Para que eu possa analisar a causa raiz, decidir entre absorver a diferença na margem ou solicitar recotação comercial, e liberar o faturamento dentro do SLA de 4 horas [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (SLA de Resolução): Todo documento em "Bloqueio Fiscal" deve ser analisado e resolvido em até 4 horas úteis. Documentos que ultrapassarem esse prazo devem escalar automaticamente para o Gerente de Controladoria.
* RN02 (Opções de Resolução): O analista pode: (a) "Autorizar Absorção" — a diferença é absorvida na margem da empresa e o faturamento segue com o valor recalculado; (b) "Devolver para Recotação" — o pedido retorna ao CRM para o vendedor renegociar com o cliente; (c) "Forçar Valores Originais" — apenas para casos com justificativa jurídica documentada.
* RN03 (Registro de Decisão): Qualquer que seja a opção escolhida, o sistema deve registrar: ID do analista, data/hora, opção selecionada, justificativa em campo de texto livre, e impacto financeiro da decisão.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Analista autoriza absorção da divergência na margem

* Dado que o documento #NF-98765 está bloqueado com divergência de IBS de R$ 40,00;
* Quando o Analista de Controladoria selecionar "Autorizar Absorção" e justificar: "Diferença dentro do limite aceitável de variação. Absorvido na margem do período.";
* Então o SAP deve liberar a emissão da NF-e com o valor recalculado (R$ 1.810,00);
* E lançar a diferença de R$ 40,00 na conta contábil "Ajuste de Margem — Variação de Alíquota";
* E registrar a decisão completa no log de auditoria [INDEX].

#### Cenário 2: SLA de 4 horas expira e escala para gerência

* Dado que o documento #NF-98766 está bloqueado há 3 horas e 58 minutos;
* E ainda não foi analisado por nenhum analista;
* Quando o relógio do sistema atingir 4 horas de bloqueio;
* Então o SAP deve alterar a prioridade do documento para "Crítica";
* E enviar uma notificação push e e-mail ao Gerente de Controladoria: "Documento #NF-98766 excedeu SLA de 4 horas em bloqueio fiscal. Ação imediata necessária.";
* E registrar a escalação no log de auditoria [INDEX].
