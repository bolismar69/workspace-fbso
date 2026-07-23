# Histórias de Usuário (User Stories) — Feature 02.2
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulos: CRM, E-commerce, Portal B2B (Onda 1)
- Feature Relacionada: 01.02.2 — Resiliência de Vendas (Contingência Local) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Ativação Automática de Fallback por Indisponibilidade do Motor Central

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Plataformas Comerciais,
Quero que os canais de vendas ativem automaticamente uma tabela local de alíquotas contingenciais quando o motor centralizado de cálculo estiver indisponível,
Para que a operação comercial não seja interrompida e o cliente possa concluir sua compra mesmo durante falhas técnicas [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Disparador do Fallback): O modo de contingência deve ser ativado automaticamente quando o motor centralizado retornar timeout (3s para CRM, 500ms para E-commerce) ou erro HTTP 5xx em 3 tentativas consecutivas.
* RN02 (Tabela de Alíquotas Médias Contingenciais): A tabela local deve conter alíquotas médias de CBS e IBS por UF, atualizadas semanalmente pela Controladoria. As alíquotas devem ser gravadas localmente em cada canal e servidas sem dependência de rede externa.
* RN03 (Sinalização de Modo de Contingência): Sempre que o canal operar em modo de contingência, deve ser exibido um indicador visual para o usuário interno (vendedor) e uma notificação automática para o time de TI e Controladoria.
* RN04 (Retorno Automático ao Modo Normal): O sistema deve verificar a disponibilidade do motor central a cada 60 segundos e desativar o modo de contingência automaticamente quando o motor voltar a responder.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Motor central fica indisponível durante simulação no CRM

* Dado que o vendedor está montando uma proposta comercial no CRM;
* E o motor centralizado de cálculo está fora do ar (erro 503 em 3 tentativas consecutivas);
* Quando o vendedor clicar em "Simular Impostos";
* Então o CRM deve detectar a falha, ativar o modo de contingência automaticamente;
* E aplicar as alíquotas médias da tabela local: CBS = 9,0% e IBS = 17,5% (média nacional);
* E exibir um banner laranja no topo da tela: "⚠️ Operando em modo de contingência fiscal. Os valores exatos serão consolidados no faturamento.";
* E notificar o time de TI e a Controladoria via canal de alertas [INDEX].

#### Cenário 2: Motor central se recupera e canal retorna ao modo normal

* Dado que o CRM está operando em modo de contingência há 12 minutos;
* E o monitor de disponibilidade detecta que o motor central voltou a responder com sucesso;
* Quando o ciclo de verificação de 60 segundos for executado;
* Então o CRM deve desativar automaticamente o modo de contingência;
* E remover o banner laranja de contingência;
* E restaurar as chamadas diretas ao motor centralizado para novas simulações;
* E registrar no log: "Modo de contingência encerrado. Duração: 12 minutos. Pedidos afetados: 7" [INDEX].

------------------------------
## 📝 US-02: Notificação ao Cliente e Política de Ajuste Posterior

### 1. Descrição da História (Visão de Negócio)

Como Product Manager de E-commerce,
Quero que o cliente seja informado de forma transparente quando sua compra for concluída sob alíquotas contingenciais,
Para que não haja surpresa ou reclamação quando o valor final do imposto for ajustado no faturamento [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Mensagem de Transparência para o Cliente): Durante o modo de contingência, o checkout deve exibir a mensagem: "Os tributos exibidos são estimativas baseadas em alíquotas médias. O valor final será consolidado no momento do faturamento, podendo variar para mais ou para menos."
* RN02 (Status do Pedido em Contingência): Pedidos concluídos sob alíquotas contingenciais devem receber o status especial "Aguardando Auditoria Fiscal" e ser direcionados para uma fila prioritária de revisão pela Controladoria em até 24 horas úteis.
* RN03 (Limite de Variação Tolerável): Se a diferença entre o valor contingencial cobrado e o valor real calculado no faturamento for superior a 5% do preço líquido, a empresa absorverá a diferença excedente, protegendo o cliente.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cliente finaliza compra no e-commerce durante contingência

* Dado que o e-commerce está operando em modo de contingência;
* E um cliente está no checkout com carrinho de R$ 500,00 (preço líquido);
* Quando o cliente finalizar o pagamento;
* Então o sistema deve exibir no resumo do pedido: "Tributos Estimados (sujeito a ajuste): CBS ~R$ 45,00 / IBS ~R$ 87,50";
* E atribuir o status "Aguardando Auditoria Fiscal" ao pedido;
* E enviar um e-mail automático ao cliente informando que o valor final dos tributos será confirmado em até 24 horas úteis [INDEX].

#### Cenário 2: Ajuste pós-contingência excede o limite de variação tolerável

* Dado que um pedido foi concluído em contingência com IBS estimado de R$ 87,50;
* E a Controladoria apurou que o IBS real devido para o destino do cliente é de R$ 110,00 (diferença de +25,7%);
* Quando o analista fiscal processar o ajuste no SAP;
* Então o sistema deve aplicar o limite de variação de 5% sobre o preço líquido (R$ 25,00);
* E o cliente deve pagar apenas R$ 87,50 + R$ 25,00 = R$ 112,50 (em vez de R$ 110,00 que seria o valor real... na verdade o cliente paga o teto e a empresa absorve R$ 22,50);
* E o valor excedente absorvido pela empresa deve ser lançado em conta de despesa operacional "Absorção de Variação Fiscal — Contingência" [INDEX].

------------------------------
## 📝 US-03: Painel de Auditoria Pós-Contingência e Reconciliação

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Controladoria,
Quero acessar um painel que liste todos os pedidos concluídos sob modo de contingência com a comparação entre valores estimados e valores reais,
Para auditar o impacto financeiro de cada evento de contingência e reportar ao CFO o custo operacional da indisponibilidade [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Rastreabilidade por Evento): Cada ativação de modo de contingência deve gerar um ID de evento único, associando todos os pedidos concluídos durante aquele período de indisponibilidade.
* RN02 (Conciliação Obrigatória): Todo pedido em status "Aguardando Auditoria Fiscal" deve ser reconciliado pela Controladoria em até 24 horas úteis, com o preenchimento obrigatório do campo "Parecer Fiscal".
* RN03 (Métricas de Impacto Financeiro): O painel deve totalizar: (a) número de pedidos afetados, (b) valor total de imposto estimado, (c) valor total de imposto real, (d) diferença absorvida pela empresa.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Controladoria audita evento de contingência de 12 minutos

* Dado que ocorreu um evento de contingência (ID EVT-20260623-001) com duração de 12 minutos;
* E 7 pedidos foram concluídos sob alíquotas contingenciais;
* Quando o Gerente de Controladoria acessar o painel de auditoria e selecionar o evento EVT-20260623-001;
* Então o painel deve exibir a lista dos 7 pedidos com: número do pedido, valor estimado, valor real, diferença, status de conciliação;
* E totalizar na barra de sumário: "Impacto Financeiro do Evento: R$ 145,30 absorvidos pela empresa em 7 pedidos";
* E permitir a exportação do relatório para PDF e Excel [INDEX].

#### Cenário 2: Conciliação de pedido pendente de auditoria

* Dado que o pedido #98765 está com status "Aguardando Auditoria Fiscal" há 18 horas;
* Quando o Analista Fiscal clicar em "Reconciliar Pedido" e preencher o Parecer Fiscal: "IBS real superior ao estimado. Cliente protegido pelo limite de variação. Diferença de R$ 18,50 absorvida pela empresa.";
* Então o sistema deve alterar o status para "Reconciliado";
* E registrar o ID do analista, data/hora e texto do parecer;
* E atualizar as métricas do painel em tempo real [INDEX].
