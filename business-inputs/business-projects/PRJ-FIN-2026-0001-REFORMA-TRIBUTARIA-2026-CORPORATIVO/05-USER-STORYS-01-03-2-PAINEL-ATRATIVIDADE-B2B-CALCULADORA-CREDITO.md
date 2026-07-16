# Histórias de Usuário (User Stories) — Feature 03.2
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulos: CRM, Portal B2B (Onda 1)
- Feature Relacionada: 01.03.2 — Painel de Atratividade B2B (Calculadora de Crédito do IVA) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Calculadora de Crédito Tributário para Cliente PJ no CRM

### 1. Descrição da História (Visão de Negócio)

Como Vendedor B2B atuando com clientes corporativos no regime de Lucro Real,
Quero que o CRM exiba automaticamente uma calculadora de crédito tributário ao simular uma venda para um cliente PJ,
Para que eu possa demonstrar ao cliente que o aumento nominal do preço (com IVA por fora) é compensado pelo crédito que ele poderá recuperar na apuração fiscal dele [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Ativação Condicionada ao Perfil do Cliente): A calculadora de crédito só deve ser exibida quando o cliente for classificado como PJ e seu regime tributário for Lucro Real ou Lucro Presumido (regimes que permitem aproveitamento de créditos). Para clientes do Simples Nacional ou Consumidores Finais PF, a calculadora não deve ser exibida.
* RN02 (Cálculo do Crédito Recuperável): O crédito estimado do cliente deve corresponder à soma dos valores de CBS e IBS destacados na proposta, assumindo que o cliente está adquirindo o produto como insumo para sua atividade econômica.
* RN03 (Ressalva Legal Obrigatória): O painel deve conter a ressalva: "* Estimativa de crédito baseada nas alíquotas vigentes e no regime tributário informado pelo cliente. O aproveitamento efetivo depende da escrituração fiscal do adquirente e está sujeito à legislação aplicável."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Vendedor acessa calculadora de crédito para cliente PJ no Lucro Real

* Dado que o vendedor simulou uma proposta de R$ 50.000,00 (preço líquido) para um cliente PJ classificado como "Lucro Real" em Curitiba (PR);
* E o sistema calculou CBS = R$ 4.400,00 e IBS = R$ 8.850,00;
* Quando o vendedor clicar na aba "Visão do Cliente" dentro da simulação;
* Então o sistema deve exibir o "Painel de Atratividade B2B" com:
  - "Preço Nominal Total: R$ 63.250,00"
  - "Crédito Tributário Recuperável Estimado: R$ 13.250,00"
  - "Custo Efetivo Real para sua Empresa: R$ 50.000,00"
* E destacar visualmente que o crédito elimina o impacto dos tributos no custo final do cliente [INDEX].

#### Cenário 2: Vendedor simula para cliente do Simples Nacional — calculadora não exibida

* Dado que o vendedor simulou uma proposta para um cliente PJ optante pelo Simples Nacional;
* Quando o vendedor acessar a tela de simulação;
* Então o painel de calculadora de crédito não deve ser exibido;
* E deve aparecer uma nota informativa: "Cliente optante pelo Simples Nacional — regime não permite aproveitamento integral de créditos de CBS/IBS na mesma modalidade do Lucro Real.";
* E o foco da negociação deve permanecer no preço total "por fora" [INDEX].

------------------------------
## 📝 US-02: Simulador de Cenário Comparativo — Regime Antigo vs. IVA Dual

### 1. Descrição da História (Visão de Negócio)

Como Gerente Comercial de Grandes Contas,
Quero visualizar um comparativo lado a lado entre o custo fiscal do regime antigo (PIS/COFINS/ICMS/ISS) e o novo regime (CBS/IBS) para o mesmo cenário de compra,
Para demonstrar ao cliente corporativo que o novo modelo não aumenta o custo tributário efetivo da operação — apenas altera a forma de cálculo e a transparência [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Cenário Comparativo Completo): O painel deve exibir duas colunas: "Regime Atual (Pré-Reforma)" e "Regime IVA Dual (Pós-Reforma)" com os valores de cada tributo linha a linha, permitindo a comparação visual direta.
* RN02 (Premissas do Cálculo Antigo): Para o regime antigo, o sistema deve assumir a cumulatividade padrão dos tributos (imposto "por dentro" compondo a própria base), refletindo a realidade fiscal anterior à reforma.
* RN03 (Indicador de Economia ou Aumento): O painel deve destacar com cores (verde para economia, vermelho para aumento) a diferença líquida entre os dois regimes na última linha: "Impacto Líquido da Reforma para este Pedido: +R$ XXX ou -R$ XXX".

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Gerente acessa comparativo para negociação de contrato anual

* Dado que o Gerente Comercial está negociando um contrato de fornecimento anual com um cliente PJ em Fortaleza (CE);
* E o pedido simulado tem preço líquido de R$ 100.000,00;
* Quando o gerente clicar em "Comparativo Regime Antigo vs. IVA Dual";
* Então o sistema deve exibir lado a lado:
  - Coluna 1 (Antigo): PIS: R$ 1.650,00 / COFINS: R$ 7.600,00 / ICMS (18%): R$ 21.500,00 / Total Antigo: R$ 130.750,00
  - Coluna 2 (Novo): CBS (8,8%): R$ 8.800,00 / IBS (17,7%): R$ 17.700,00 / Total Novo: R$ 126.500,00
  - Linha final: "Economia estimada no IVA Dual: -R$ 4.250,00 (3,2% menor que o regime antigo)" em destaque verde [INDEX].

#### Cenário 2: Cenário onde o IVA Dual resulta em aumento de carga

* Dado que o pedido simulado é para um cliente em um município com IBS elevado (ex: alíquota municipal máxima);
* Quando o comparativo for gerado;
* Então a linha de "Impacto Líquido" deve aparecer em vermelho indicando o aumento;
* E o sistema deve sugerir ao vendedor: "Considere negociar com o cliente a aplicação do Token de Garantia Fiscal para proteção contra futuras oscilações de alíquota" [INDEX].

------------------------------
## 📝 US-03: Painel de Atratividade no Portal B2B (Autoatendimento do Cliente)

### 1. Descrição da História (Visão de Negócio)

Como Cliente PJ que acessa o Portal B2B de autoatendimento,
Quero visualizar automaticamente o crédito tributário estimado ao montar meu pedido de compra,
Para tomar decisões de compra com clareza sobre o custo efetivo real, sem precisar contatar um vendedor para entender o impacto dos novos tributos [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Exibição Nativa no Portal): O painel de atratividade (crédito estimado e custo efetivo) deve ser exibido como parte nativa da tela de montagem de pedidos no Portal B2B, sem que o cliente precise acionar um botão separado — a informação deve estar sempre visível.
* RN02 (Atualização Dinâmica): O painel deve ser recalculado automaticamente sempre que o cliente alterar a quantidade de itens, adicionar ou remover produtos do carrinho.
* RN03 (Mensagem Comercial Estratégica): O Portal deve exibir uma mensagem de posicionamento: "No regime do IVA Dual, o imposto é calculado por fora e integralmente creditável para sua empresa. O valor que você paga de CBS e IBS retorna como crédito na sua apuração fiscal."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cliente PJ acessa o portal e monta pedido com visualização de crédito

* Dado que um cliente PJ do Lucro Real está logado no Portal B2B;
* E ele adicionou 20 unidades do SKU 98765 ao carrinho (preço líquido total de R$ 15.000,00);
* Quando o portal renderizar a tela de carrinho;
* Então o painel lateral de "Seu Crédito Tributário" deve exibir:
  - "Valor Nominal deste Pedido: R$ 18.975,00"
  - "Crédito de CBS/IBS que você recuperará: R$ 3.975,00"
  - "Custo Efetivo Real: R$ 15.000,00"
* E a mensagem comercial estratégica deve estar visível abaixo do painel [INDEX].

#### Cenário 2: Cliente altera quantidades e vê o crédito atualizar dinamicamente

* Dado que o cliente PJ está com um carrinho montado no Portal B2B;
* E o painel de crédito está exibindo "Crédito Recuperável: R$ 3.975,00";
* Quando o cliente aumentar a quantidade de um item de 20 para 30 unidades;
* Então o painel deve recalcular automaticamente em menos de 2 segundos;
* E exibir o novo valor de crédito proporcional à quantidade atualizada;
* E o custo efetivo real deve refletir apenas o preço líquido dos novos itens [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
