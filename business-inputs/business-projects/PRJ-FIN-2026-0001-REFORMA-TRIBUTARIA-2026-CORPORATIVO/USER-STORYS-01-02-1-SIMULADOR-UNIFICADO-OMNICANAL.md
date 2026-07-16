# Histórias de Usuário (User Stories) — Feature 02.1
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulos: CRM, E-commerce, Portal B2B (Onda 1)
- Feature Relacionada: 01.02.1 — Simulador Unificado em Telas de Proposta (Omnicanalidade) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Simulação de Impostos no CRM Durante Montagem da Proposta

### 1. Descrição da História (Visão de Negócio)

Como Vendedor da Força de Vendas Nacional,
Quero acionar um botão de "Simular Impostos" na tela de montagem da proposta comercial no CRM,
Para que eu possa visualizar o impacto exato de CBS, IBS e IS no preço final antes de apresentar a proposta ao cliente [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Requisição à Inteligência Corporativa): O botão "Simular Impostos" deve enviar os dados do cenário comercial (itens, quantidades, preço base, código IBGE de destino, perfil do cliente) para o motor centralizado de cálculo de impostos e exibir o retorno em tempo real.
* RN02 (Simulação Não Persistente): A simulação é volátil e não gera registro fiscal ou compromisso comercial — apenas informa o vendedor para tomada de decisão.
* RN03 (Perfil do Cliente Impacta o Cálculo): A simulação deve considerar se o cliente é Contribuinte PJ no Lucro Real (com direito a crédito) ou Consumidor Final PF (sem crédito), pois isso altera a exibição do preço e as orientações comerciais.
* RN04 (Indisponibilidade do Motor): Se o motor centralizado não responder em até 3 segundos, o CRM deve exibir uma mensagem de fallback orientando o vendedor a utilizar a tabela de alíquotas médias de contingência.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Vendedor simula impostos para cliente PJ no Lucro Real

* Dado que o vendedor está montando uma proposta no CRM para um cliente PJ em Belo Horizonte (MG);
* E o cadastro do cliente possui Código IBGE validado e perfil "Lucro Real";
* E o carrinho contém 10 itens totalizando R$ 10.000,00 de preço líquido base;
* Quando o vendedor clicar no botão "Simular Impostos";
* Então o motor de cálculo deve retornar em menos de 2 segundos o valor projetado de CBS, IBS e Total;
* E o CRM deve exibir no resumo: Preço Líquido: R$ 10.000,00 / CBS (Federal): R$ 880,00 / IBS (MG - Belo Horizonte): R$ 1.770,00 / Total: R$ 12.650,00;
* E indicar que o cliente PJ poderá se apropriar de créditos de CBS e IBS [INDEX].

#### Cenário 2: Vendedor simula impostos para Consumidor Final PF

* Dado que o vendedor está montando uma proposta para um cliente PF em Salvador (BA);
* Quando o vendedor clicar em "Simular Impostos";
* Então o sistema deve exibir o preço total "por fora" com CBS e IBS;
* E não deve exibir a seção de crédito recuperável (exclusiva para PJ);
* E destacar que o valor total inclui os tributos de destino conforme legislação de transparência fiscal [INDEX].

#### Cenário 3: Timeout do motor de cálculo

* Dado que o vendedor acionou a simulação de impostos para um cenário complexo com 50 itens;
* Quando o motor centralizado não responder em até 3 segundos;
* Então o CRM deve exibir: "⚠️ Motor de Cálculo temporariamente indisponível. Utilize a tabela de alíquotas médias de contingência para estimativa.";
* E disponibilizar um botão de acesso rápido à tabela de contingência local [INDEX].

------------------------------
## 📝 US-02: Simulação de Impostos no Carrinho do E-commerce

### 1. Descrição da História (Visão de Negócio)

Como Cliente Final navegando no E-commerce da empresa,
Quero visualizar a estimativa dos impostos CBS e IBS sobre o valor do meu carrinho antes de finalizar a compra,
Para que eu tenha total transparência da composição do preço e não enfrente surpresas no momento do pagamento [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Gatilho Automático no Carrinho): A simulação de impostos deve ser disparada automaticamente sempre que o cliente alterar a quantidade de itens no carrinho ou informar o CEP de entrega.
* RN02 (SLA de Performance Estrito): O tempo de resposta da simulação para o cliente no e-commerce não pode ultrapassar 100ms, sob risco de degradação da experiência de compra e aumento da taxa de abandono de carrinho.
* RN03 (Cache de Alíquotas por Destino): Para atingir o SLA de performance, o e-commerce deve utilizar uma camada de cache de alíquotas por código IBGE com TTL de 2 horas, invocando o motor centralizado apenas em caso de cache miss.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cliente informa CEP de entrega e visualiza impostos no carrinho

* Dado que um cliente anônimo navega no e-commerce e adicionou produtos ao carrinho;
* Quando o cliente informar o CEP de entrega "30130-110" (Belo Horizonte - MG) no carrinho;
* Então o sistema deve consultar a camada de cache de alíquotas por Código IBGE (3106200);
* E atualizar dinamicamente o resumo do carrinho exibindo: Subtotal (Produtos) / CBS (Federal) / IBS (Destino) / Total;
* E o tempo total da atualização deve ser inferior a 100ms [INDEX].

#### Cenário 2: Cliente altera CEP para outro estado e vê o imposto mudar

* Dado que o cliente visualizou o carrinho com entrega em MG (IBS de MG);
* Quando o cliente alterar o CEP de entrega para "01310-100" (São Paulo - SP);
* Então o sistema deve invalidar o cache do Código IBGE anterior e consultar as alíquotas de SP;
* E o valor do IBS no resumo deve ser recalculado automaticamente, refletindo a alíquota do município de São Paulo;
* E o valor da CBS (federal) deve permanecer inalterado [INDEX].

------------------------------
## 📝 US-03: Consistência Cross-Canal — Mesmo Preço em CRM, Portal B2B e E-commerce

### 1. Descrição da História (Visão de Negócio)

Como Diretor de Canais e Experiência do Cliente,
Quero garantir que um mesmo produto, para um mesmo cliente e destino, apresente exatamente o mesmo preço calculado independentemente do canal utilizado (CRM do vendedor, Portal B2B ou E-commerce),
Para que a empresa mantenha a credibilidade comercial e evite disputas de preço entre canais [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Fonte Única de Cálculo): Todos os canais (CRM, Portal B2B, E-commerce) devem consumir exatamente o mesmo endpoint da inteligência corporativa de cálculo de impostos, sendo vedada qualquer lógica de cálculo de tributos embarcada localmente nos canais.
* RN02 (Parâmetros Idênticos de Simulação): A requisição de simulação enviada por qualquer canal deve conter o mesmo payload padronizado: ID_Cliente, Código_IBGE_Destino, Itens (SKU, Quantidade, Preço_Base), Perfil_Cliente (PJ/PF).
* RN03 (Teste de Consistência Automatizado): Um robô de auditoria deve executar diariamente simulações idênticas em todos os canais para um conjunto de cenários predefinidos e alertar o time de TI em caso de divergência de qualquer valor.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Mesma simulação no CRM e no Portal B2B retorna valores idênticos

* Dado que um cliente PJ acessa o Portal B2B e simula a compra de 5 unidades do SKU 12345 para entrega em Recife (PE);
* E simultaneamente, um vendedor interno simula a mesma composição no CRM para o mesmo cliente;
* Quando ambos os canais enviarem a requisição de simulação para o motor corporativo;
* Então o valor de CBS, IBS e Total retornado deve ser absolutamente idêntico em ambos os canais, centavo por centavo;
* E o timestamp da simulação deve ser registrado em cada canal para auditoria [INDEX].

#### Cenário 2: Robô de auditoria detecta divergência cross-canal

* Dado que o robô de auditoria executa sua rotina diária de simulações cross-canal às 06:00;
* Quando o robô identificar que o Portal B2B retornou IBS de R$ 177,00 e o E-commerce retornou IBS de R$ 178,50 para o mesmo cenário;
* Então o sistema de monitoramento deve gerar um alerta de severidade "Alta" para o time de TI;
* E registrar a divergência em log com os parâmetros exatos de cada requisição para diagnóstico;
* E notificar o PMO Corporativo sobre o risco de inconsistência comercial [INDEX].
