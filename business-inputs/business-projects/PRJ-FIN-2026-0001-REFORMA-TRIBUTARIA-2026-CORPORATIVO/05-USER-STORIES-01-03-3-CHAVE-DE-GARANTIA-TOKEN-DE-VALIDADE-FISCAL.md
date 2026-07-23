# Histórias de Usuário (User Stories) — Feature 03.3
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulos: CRM, E-commerce e Faturamento ERP SAP (Integração Onda 1 -> Onda 2)
- Feature Relacionada: 01.03.3 — Chave de Garantia Comercial / Token de Validade Fiscal [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Geração do Token de Garantia Fiscal no Fechamento da Proposta Comercial

### 1. Descrição da História (Visão de Negócio)

Como Product Manager dos Canais Comerciais,
Quero que o sistema gere um identificador único estruturado (Token) no momento do aceite de uma proposta comercial ou fechamento de carrinho,
Para que as premissas de preço líquido, alíquota de IBS/CBS calculadas por fora e dados geográficos fiquem congelados e protegidos contra flutuações fiscais diárias [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Payload de Dados do Token): O Token gerado não deve ser apenas uma hash aleatória, mas deve carregar ou referenciar a matriz de dados: ID_Cliente, Código_IBGE_Destino, Preço_Líquido_Base, Alíquota_CBS_Aplicada, Alíquota_IBS_Aplicada e Timestamp_Geracao [INDEX].
* RN02 (Janela de Tolerância Comercial): A janela padrão de validade e congelamento do token será de 48 horas a partir do timestamp de geração, servindo como a política de tolerância para que o cliente realize o pagamento ou o fluxo envie o pedido ao SAP.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cliente conclui o pedido no E-commerce ou aceita proposta no CRM

* Dado que um cliente PJ do Lucro Real aceitou uma proposta comercial contendo alíquota interestadual de IBS calculada por fora;
* Quando o sistema de vendas (CRM/E-commerce) alterar o status da oportunidade para "Ganha/Fechada";
* Então o motor corporativo de cálculo deve emitir uma assinatura digital (Token de Validade Fiscal) associada àquele pedido;
* E gravar no registro do pedido a data/hora exata da expiração (Timestamp de Geração + 48 horas).

------------------------------
## 📝 US-02: Validação do Token no Faturamento (SAP) dentro do Prazo de Validade

### 1. Descrição da História (Visão de Negócio)

Como Analista de Faturamento do ERP SAP,
Quero que o processo de faturamento valide a integridade do Token recebido dos canais de vendas,
Para que a nota fiscal seja emitida com os exatos valores tributários e comerciais prometidos ao cliente, mantendo consistência absoluta de centavos [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Consistência Matemática Mandatória): Se o token estiver válido, o módulo SAP SD (Vendas e Distribuição) deve herdar e forçar as condições de imposto (CBS/IBS) originais do token, ignorando qualquer atualização diária de alíquotas que tenha ocorrido no intervalo entre a venda e o faturamento [INDEX].
* RN02 (Absorção de Variações na Margem): Eventuais aumentos de alíquotas decretados pelo governo nesse intervalo de 48h serão absorvidos como custo pela empresa, protegendo o preço final acordado com o cliente.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Pedido chega ao faturamento dentro da janela de 48 horas

* Dado que o módulo SAP SD recebeu uma ordem de faturamento originada do e-commerce;
* E o pedido veio acompanhado de um "Token de Validade Fiscal" cujo prazo de expiração ainda não foi atingido;
* Quando o processo de faturamento pré-emissão for acionado;
* Então o SAP deve validar o Token com sucesso e travar os valores de CBS e IBS exatamente iguais aos simulados na venda [INDEX];
* E liberar a emissão da nota fiscal eletrônica com 100% de consistência com a proposta comercial.

------------------------------
## 📝 US-03: Rejeição de Token Expirado no SAP com Recálculo Obrigatório

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Controladoria e Risco Fiscal,
Quero que o SAP barre o faturamento se o pedido apresentar um Token vencido,
Para que a empresa não emita notas fiscais com alíquotas defasadas que gerem autuações ou perdas financeiras crônicas.

### 2. Regras de Negócio (Business Rules)

* RN01 (Estouro de Janela Comercial): Pedidos que passarem mais de 48 horas na fila de aprovação financeira, análise de crédito ou problemas logísticos perdem o direito ao congelamento fiscal.
* RN02 (Reprecificação Forçada): Ao rejeitar o token, o sistema é obrigado a realizar uma nova requisição à inteligência corporativa para aplicar as alíquotas de destino atualizadas na data do faturamento [INDEX].

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Pedido com token vencido tenta ser faturado no SAP

* Dado que um pedido comercial ficou travado por 72 horas em análise de crédito financeira interna;
* E o seu "Token de Validade Fiscal" associado já ultrapassou o horário limite de expiração;
* Quando o sistema tentar rodar o faturamento automático deste pedido no SAP;
* Então o ERP deve rejeitar a ordem de faturamento, atribuindo o status de "Bloqueio por Token Expirado";
* E forçar o recálculo do IVA Dual atualizado na data corrente, atualizando o preço final do pedido [INDEX];
* E notificar o vendedor de que o pedido precisa de revalidação comercial com o cliente devido ao estouro do prazo fiscal.

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
