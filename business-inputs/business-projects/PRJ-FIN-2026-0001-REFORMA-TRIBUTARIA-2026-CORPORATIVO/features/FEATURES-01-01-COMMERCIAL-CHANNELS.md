# Detalhamento de Features da Onda 01: Canais Comerciais e Vendas
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Fase: Onda 1 (Sistemas de Vendas, CRM e Plataformas Comerciais)
- Épico: 01.01 (Qualificação Geográfica e Onboarding CRM), 01.02 (Conexão à Inteligência Corporativa de Cálculo), 01.03 (Precificação Dinâmica e Transparência)
- Status: Pronto para Detalhamento de User Stories
- Responsáveis: Product Managers (PMs) de Vendas e POs de Canais Comerciais

------------------------------
## 🔍 ÉPICO 01.01: Qualificação Geográfica, Saneamento e Onboarding de Clientes (CRM)

### Feature 01.01.1: Validação Cadastral Geográfica em Tempo Real

* Objetivo de Negócio: Garantir a exatidão do local de consumo para correta aplicação da alíquota do IBS baseada no princípio do destino [INDEX].
* Descrição Funcional: O sistema deve validar de forma mandatória os dados geográficos de novos clientes ou alterações cadastrais na entrada (onboarding). O preenchimento do CEP deve acionar uma validação que preenche automaticamente o Estado, o Município e o Código IBGE oficial, impedindo digitação manual desses campos.
* Regras de Negócio e Restrições:
* Se o endereço for de uma zona rural ou sem CEP específico, a interface deve exigir o preenchimento manual guiado por lista fechada de municípios homologados pelo IBGE, bloqueando caracteres livres.
   * O cadastro só será marcado como "Apto para Venda" se possuir o ID geográfico do IBGE vinculado.

### Feature 01.01.2: Governança e Trava Comercial de Vendas (CRM)

* Objetivo de Negócio: Blindar a empresa contra erros de precificação causados por cadastros desatualizados de clientes antigos.
* Descrição Funcional: Mecanismo de bloqueio operacional nas telas de criação de oportunidades, propostas e orçamentos dentro do CRM. O sistema deve varrer o cadastro do cliente selecionado; se os dados de localização (Código IBGE) não estiverem higienizados conforme a Reforma Tributária, o botão de "Gerar Proposta" deve ficar desabilitado.
* Regras de Negócio e Restrições:
* O sistema deve exibir um aviso claro na tela: "Proposta bloqueada: Cadastro do cliente necessita de atualização geográfica para enquadramento no IVA Dual".
   * Deve existir um fluxo de aprovação de exceção para que o Gerente Comercial libere a trava apenas em casos de contingência jurídica comprovada.

------------------------------
## ⚙️ ÉPICO 01.02: Conexão Comercial à Inteligência Corporativa de Cálculo

### Feature 01.02.1: Simulador Unificado em Telas de Proposta (Omnicanalidade)

* Objetivo de Negócio: Garantir que o vendedor e o cliente vejam o mesmo impacto tributário em qualquer canal [INDEX].
* Descrição Funcional: Disponibilização de um botão "Simular Impostos" na montagem do carrinho de compras (E-commerce) e na tela de propostas (CRM). Ao ser acionado, o sistema envia o cenário comercial para a inteligência de cálculo e retorna o valor projetado de CBS, IBS e IS sem criar um registro fiscal definitivo.
* Regras de Negócio e Restrições:
* O tempo de resposta para o cliente na tela do e-commerce não pode prejudicar a experiência de compra (alvo de mercado de alta performance).
   * A simulação deve considerar o perfil do cliente (Contribuinte PJ no Lucro Real vs. Consumidor Final PF), pois isso altera o comportamento do cálculo de negócios.

### Feature 01.02.2: Resiliência de Vendas (Contingência Local)

* Objetivo de Negócio: Evitar a perda de vendas na ponta caso a inteligência central de cálculo fique temporariamente indisponível.
* Descrição Funcional: Mecanismo de fallback de negócios para os canais comerciais. Se a consulta ao motor de impostos falhar por timeout ou indisponibilidade, a plataforma de vendas deve ativar uma política de contingência comercial.
* Regras de Negócio e Restrições:
* O sistema adota uma tabela estática de alíquotas médias contingenciais gravada localmente no canal para permitir que o cliente conclua o pedido.
   * O pedido entra com o status "Aguardando Auditoria Fiscal" e o cliente é notificado discretamente na tela de que as condições fiscais finais serão consolidadas no faturamento.

------------------------------
## 💰 ÉPICO 01.03: Precificação Dinâmica, Margem Líquida e Transparência ("Por Fora")

### Feature 01.03.1: Interface Visual de Checkout com Decomposição do IVA

* Objetivo de Negócio: Atender às leis de transparência fiscal e educar o mercado sobre a precificação "por fora" [INDEX].
* Descrição Funcional: Redesenho completo do resumo financeiro do carrinho de compras e da proposta impressa. O valor principal do produto/serviço deve passar a representar o Preço Líquido (sem imposto por dentro). Abaixo dele, devem ser criadas linhas explícitas destacando o CBS (Federal), o IBS (Destino) e o valor totalizado a pagar.
* Regras de Negócio e Restrições:
* Exemplo visual obrigatório para o layout:
   * Valor do Item: R$ 1.000,00
      * CBS (Federal - X%): R$ 88,00
      * IBS (Destino - Y%): R$ 177,00
      * Total do Pedido: R$ 1.265,00
   
### Feature 01.03.2: Painel de Atratividade B2B (Calculadora de Crédito do IVA)

* Objetivo de Negócio: Reduzir o atrito comercial do aumento nominal de preço para clientes corporativos (PJ).
* Descrição Funcional: Inclusão de um painel informativo exclusivo para vendas corporativas (B2B) no CRM e no portal de clientes. Ao simular a venda para uma empresa que opera no Lucro Real, o sistema deve calcular e exibir o valor do crédito que aquele cliente irá recuperar na entrada dele.
* Regras de Negócio e Restrições:
* O painel deve exibir uma mensagem comercial de impacto: "Preço Nominal: R$ 1.265,00. Crédito Recuperável estimado para sua empresa: R$ 265,00. Custo Efetivo Real: R$ 1.000,00".

### Feature 01.03.3: Chave de Garantia Comercial (Token de Validade Fiscal)

* Objetivo de Negócio: Impedir prejuízos causados por alterações de alíquotas entre a negociação comercial e a emissão da nota [INDEX].
* Descrição Funcional: Mecanismo que gera um identificador único (Token) para cada proposta comercial emitida e aceita. Esse token congela as premissas de preço base e as alíquotas simuladas de IBS/CBS por uma janela regulamentada de tempo (ex: 48 horas).
* Regras de Negócio e Restrições:
* Se o faturamento (Onda 2) receber o pedido dentro do prazo de validade do token, ele deve respeitar o valor acordado, absorvendo eventuais variações diárias de impostos na margem da empresa. Se o prazo expirar, o pedido exige uma nova simulação obrigatória.

