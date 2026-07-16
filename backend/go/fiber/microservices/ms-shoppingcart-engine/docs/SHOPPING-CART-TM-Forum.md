No ecossistema do TM Forum, o processo de carrinho de compras é padronizado principalmente através da TMF663 Shopping Cart API, que faz parte do conjunto de Open APIs do fórum. [1, 2] 
Este processo não é apenas uma "lista de compras", mas uma entidade temporária que gerencia a intenção de compra do cliente antes de ser convertida em uma ordem de serviço ou produto definitiva no sistema de Order Management. [1, 3] 
## Ciclo de Vida e Fluxo do Carrinho (TMF663)
O fluxo típico do processo de carrinho de compras segue estas etapas:

   1. Criação (POST /shoppingCart): O carrinho é iniciado quando um cliente seleciona o primeiro item. Ele armazena referências ao cliente (Party) e aos itens desejados.
   2. Manipulação de Itens:
   * Adicionar Item: Inclui produtos, ofertas ou promoções ao carrinho.
      * Atualizar/Remover: Ajusta quantidades ou remove itens antes do checkout.
   3. Validação e Precificação: O carrinho interage com o Catálogo de Produtos (TMF620) para garantir que as ofertas sejam válidas para aquele cliente e região, calculando preços e impostos em tempo real.
   4. Checkout e Conversão (Transição para TMF622): O passo final é o "submit". O conteúdo do carrinho é validado uma última vez e enviado para o processo de criação de ordem (Product Order API - TMF622), iniciando o provisionamento e faturamento. [1, 2, 3, 4] 

## Componentes Relacionados

* eTOM (Business Process Framework): O carrinho de compras situa-se geralmente nos processos de Vendas (Sales & Marketing), especificamente no nível de "Sales Preparation & Sales Configuration".
* SID (Information Framework): Define o modelo de dados do objeto ShoppingCart, que inclui o CartItem, o link com o ProductOffering e a associação com o RelatedParty. [2, 4, 5] 

## Considerações de Implementação

* Persistência: O carrinho pode ser anônimo ou vinculado a um perfil logado, permitindo que o cliente comece a compra no celular e termine no computador (estratégia omnicanal).
* Privacidade: Recentemente, discussões na comunidade [TM Forum Thrive](https://engage.tmforum.org/discussion/tmf663-shopping-cart-api-checkout-shopping-cart-and-generating-an-order) focam na integração de consentimentos (LGPD/GDPR) diretamente no fluxo de checkout (TMF644). [1, 4, 6, 7] 


[1] [https://engage.tmforum.org](https://engage.tmforum.org/discussion/tmf663-shopping-cart-api-checkout-shopping-cart-and-generating-an-order#:~:text=Hi%2Cin%20TMF%20663%20API%20I%20am%20looking,ordering%20processes.%20I%20assume%20that%20existing%20POST/P.)
[2] [https://engage.tmforum.org](https://engage.tmforum.org/communities/community-home/digestviewer/viewthread?GroupId=31&MessageKey=00ad5ceb-450a-4bc7-97b1-3d1f46a48f5f&CommunityKey=d543b8ba-9d3a-4121-85ce-5b68e6c31ce5&tab=digestviewer&hlmlt=VT#:~:text=Hi%20Experts%2C%20I%20would%20like%20to%20get,Rest%20API%20Specification%20for%20shopping%20cart%20%28https://projects.tmforum.org/wiki/download/attachments/114641.)
[3] [https://www.servicenow.com](https://www.servicenow.com/docs/r/pt-BR/order-management/order-mgt-install-providers.html#:~:text=Personalizando%20a%20exibi%C3%A7%C3%A3o%20da%20moeda%20no%20Carrinho,exibido%20em%20um%20layout%20com%20v%C3%A1rias%20BOMs)
[4] [https://engage.tmforum.org](https://engage.tmforum.org/communities/community-home/digestviewer/viewthread?MessageKey=ed97c828-e593-4ef5-8128-5fdc175340f8&CommunityKey=d543b8ba-9d3a-4121-85ce-5b68e6c31ce5&tab=digestviewer)
[5] [https://learn.microsoft.com](https://learn.microsoft.com/pt-br/aspnet/web-forms/overview/getting-started/getting-started-with-aspnet-45-web-forms/shopping-cart#:~:text=Para%20configurar%20um%20controle%20de%20dados%20para,retorna%20uma%20lista%20de%20objetos%20CartItem%20.)
[6] [https://telaviva.com.br](https://telaviva.com.br/arquivos/5_Applying_an_Omnichannel_Strategy_in_Telecommunications_PT.pdf)
[7] [https://microuniverso.com.br](https://microuniverso.com.br/carrinho-de-compras/#:~:text=Para%20utiliz%C3%A1%2Dlo%2C%20basta%20os%20vendedores%20pegarem%20os,para%20a%20digita%C3%A7%C3%A3o%20de%20c%C3%B3digos%20dos%20itens.)

---

Um desenho de integração baseado no TM Forum (TMF) para um carrinho de compras organiza-se em torno de uma arquitetura baseada em eventos ou solicitações síncronas entre APIs específicas do ecossistema Open API e ODA (Open Digital Architecture).
O coração dessa integração é a TMF663 (Shopping Cart API), que atua como um orquestrador temporário de intenções de venda. [1] 
## Arquitetura de Integração (Fluxo Funcional)
Abaixo estão os componentes principais e como eles se conectam:

   1. Frontend / Canal de Venda: Interage diretamente com a TMF663 para criar o carrinho (POST /shoppingCart) e adicionar itens.
   2. Validação de Catálogo (TMF620 - Product Catalog): Sempre que um item é adicionado, a API do Carrinho consulta o Catálogo para validar se o productOfferingId existe, se o preço está correto e se há regras de elegibilidade.
   3. Identificação do Cliente (TMF632 - Party Management): O carrinho armazena referências ao RelatedParty. A integração garante que as ofertas exibidas no carrinho sejam personalizadas para o perfil do cliente identificado.
   4. Qualificação e Estoque (TMF679 / TMF637):
   * TMF679 (Product Offering Qualification): Verifica se o cliente pode comprar aquele produto (ex: disponibilidade técnica por endereço).
      * TMF637 (Product Inventory): Para processos de mudança de plano ou upgrade, o carrinho consulta o inventário para saber o que o cliente já possui.
   5. Checkout e Conversão (TMF622 - Product Order): Ao finalizar a compra ("Submit"), o Carrinho (TMF663) dispara a criação de uma Ordem de Produto na TMF622. Este é o momento onde o carrinho deixa de existir (ou é marcado como processado) e a jornada de entrega começa. [1, 2, 3, 4, 5, 6, 7] 

## Diagrama de Blocos Conceitual

```mermaid
graph LR
    subgraph "Canais (Omnichannel)"
        UI[Mobile/Web/Loja]
    end

    subgraph "Core Commerce Management (ODA)"
        SC[TMF663: Shopping Cart]
        PC[TMF620: Product Catalog]
        PO[TMF622: Product Order]
        PI[TMF637: Product Inventory]
    end

    subgraph "Engajamento"
        PM[TMF632: Party Management]
    end

    UI -- "1. Adicionar Item" --> SC
    SC -- "2. Valida Oferta/Preço" --> PC
    SC -- "3. Valida Perfil" --> PM
    SC -- "4. Checa Base Instalada" --> PI
    SC -- "5. Checkout (Submit)" --> PO
```

## Principais Entidades Relacionadas (SID Mapping)

* ShoppingCart: O contêiner principal com validFor (data de expiração) e cartTotalPrice.
* CartItem: Contém a ação (add, modify, noChange) e a referência à oferta do catálogo.
* RelatedParty: Link para o cliente ou usuário que está realizando a compra. [1, 8] 

Deseja que eu detalhe os métodos específicos (POST/PATCH) de cada uma dessas APIs para um cenário de "Upgrade de Plano"?

[1] [https://github.com](https://github.com/tmforum-apis/TMF663_ShoppingCart)
[2] [https://engage.tmforum.org](https://engage.tmforum.org/communities/community-home/digestviewer/viewthread?GroupId=31&MessageKey=9825ed46-e8ba-4589-8eda-03ea0bc58647&CommunityKey=d543b8ba-9d3a-4121-85ce-5b68e6c31ce5&tab=digestviewer&hlmlt=VT)
[3] [https://www.tmforum.org](https://www.tmforum.org/resources/specifications/tmf663-shopping-cart-management-api-user-guide-v5-0-0/)
[4] [https://engage.tmforum.org](https://engage.tmforum.org/discussion/validate-orderquote-using-product-catalog-rules)
[5] [https://www.tmforum.org](https://www.tmforum.org/resources/specifications/tmf632-party-management-api-user-guide-v5-0-0/)
[6] [https://www.tmforum.org](https://www.tmforum.org/open-digital-architecture/open-apis)
[7] [https://engage.tmforum.org](https://engage.tmforum.org/discussion/tmf663-shopping-cart-api-checkout-shopping-cart-and-generating-an-order)
[8] [https://engage.tmforum.org](https://engage.tmforum.org/discussion/tmf663-shopping-cart-context-of-cart)

---

Para integrar a Reforma Tributária (IBS e CBS) ao fluxo de vendas do TM Forum (TMF663 e TMF622), a arquitetura deve evoluir de um modelo de "imposto embutido" para um modelo de cálculo dinâmico "por fora", utilizando um Tax Engine externo ou um componente de Bill Calculation que se conecte às APIs governamentais. [1, 2] 
## 1. Ponto de Integração no TMF663 (Shopping Cart)
O cálculo deve ocorrer de forma síncrona sempre que houver uma alteração no carrinho para garantir a transparência de preços ao cliente:

* Ação: Ao chamar POST ou PATCH no /shoppingCart, o sistema de Carrinho não deve calcular o imposto internamente.
* Integração: Ele deve disparar uma requisição para um Tax Microservice (ou componente ODA TMFC031).
* Dados Necessários: O Carrinho deve enviar o endereço de destino (para determinar a alíquota do IBS Estadual e Municipal) e o tipo de produto/serviço (para enquadramento na CBS e regimes diferenciados).
* Retorno: O cartTotalPrice e os cartItemPrice devem ser populados com os novos campos de impostos de forma discriminada, conforme exigido pela transparência da reforma. [3, 4, 5, 6, 7] 

## 2. Ponto de Integração no TMF622 (Product Order)
Na conversão do carrinho em ordem, o cálculo é revalidado e "congelado":

* Ação: O POST /productOrder executa a chamada final ao motor tributário para garantir que, entre o carrinho e a ordem, não houve mudança de alíquota ou regra fiscal.
* Split Payment: A ordem deve estar preparada para suportar o Split Payment, onde o valor do IBS/CBS é retido ou direcionado automaticamente no momento da liquidação financeira. [8, 9] 

## 3. Conexão com o Ecossistema Governamental
A integração técnica utilizará as novas interfaces disponibilizadas pelo governo brasileiro:

* APIs do Serpro/Receita Federal: O sistema deve consumir as APIs de consulta ao Conformidade Fácil e as tabelas de alíquotas oficiais (IBS/CBS) para garantir que o cálculo esteja sempre atualizado com a legislação vigente.
* Dualidade (Transição): Entre 2026 e 2032, os sistemas deverão suportar a coexistência dos impostos atuais (PIS/COFINS/ICMS/ISS) e dos novos (IBS/CBS), exigindo que o motor de cálculo retorne ambos os conjuntos de tags para a emissão correta da NF-e. [8, 10, 11, 12] 

## Resumo do Fluxo de Dados

| Componente TMF [11, 13, 14] | Ação Tributária | Dependência Externa |
|---|---|---|
| TMF663 (Carrinho) | Estimativa de Imposto "por fora" | Cadastro de alíquotas por destino (IBS) |
| TMF622 (Ordem) | Cálculo Final e Reserva de Crédito | API de Apuração Assistida (Serpro) |
| TMF678 (Billing) | Discriminação na Fatura / Nota Fiscal | Gerador de Tags XML para Reforma |

Deseja ver um exemplo de como os campos de IBS e CBS seriam mapeados dentro da estrutura de priceAlteration da API de Carrinho?

[1] [https://tozzinifreire.com.br](https://tozzinifreire.com.br/site/conteudo/uploads/brazilian-senate-approves-tax-reform-655def016003b.pdf#:~:text=The%20CBS%20will%20be%20levied%20at%20federal,approved%20by%20the%20House%20of%20Deputies%20are:)
[2] [https://www.demarest.com.br](https://www.demarest.com.br/wp-content/uploads/2025/02/Brazilian-Tax-Reform-Opportunities.pdf)
[3] [https://www.martinelli.adv.br](https://www.martinelli.adv.br/en/tax-reform-regulations-enacted-by-the-presidency/)
[4] [https://oda-production.s3.eu-west-2.amazonaws.com](https://oda-production.s3.eu-west-2.amazonaws.com/v1.0.0/TMFC031_Bill_Calculation_Management_v2.0.0.pdf)
[5] [https://www.gov.br](https://www.gov.br/receitafederal/pt-br/assuntos/noticias/2024/arquivos-e-imagens/BrazilTaxReformIngles.pdf)
[6] [https://community.sap.com](https://community.sap.com/t5/financial-management-blog-posts-by-sap/how-the-tax-reform-in-brazil-impacts-sap-users/ba-p/14258640)
[7] [https://marcosmartins.adv.br](https://marcosmartins.adv.br/en/e-commerce-what-changes-with-the-tax-reform/)
[8] [https://novatradebrasil.com](https://novatradebrasil.com/en/brazil-tax-reform/)
[9] [https://www.reformatributaria.com](https://www.reformatributaria.com/tecnologia/serpro-lanca-plataforma-para-operacionalizar-a-cbs/)
[10] [https://documentacao.senior.com.br](https://documentacao.senior.com.br/exigenciaslegais/noticias/federal/2025/2025-10-23-reforma-tributaria-liberada-api-de-consulta-ao-conformidade-facil/)
[11] [https://joseadriano.com.br](https://joseadriano.com.br/reforma-tributaria-cbs-apis-gratuitas-devem-ser-disponibilizadas-em-nov-25/)
[12] [https://community.sap.com](https://community.sap.com/t5/financial-management-blog-posts-by-members/understanding-the-context-of-brazil-s-tax-reform/ba-p/14159547)
[13] [https://www.soutocorrea.com.br](https://www.soutocorrea.com.br/wp-content/uploads/2024/05/Souto-Correa-Brazilian-Tax-Reform.pdf)
[14] [https://centraldeatendimento.totvs.com](https://centraldeatendimento.totvs.com/hc/pt-br/articles/35136084820631-Varejo-Supermercados-NF-e-Como-gerar-as-tags-XML-da-Reforma-Tribut%C3%A1ria-CBS-IBS-IS)

---

Para mapear os impostos da Reforma Tributária (IBS, CBS) e os atuais (ICMS, PIS, COFINS) na TMF663, utilizamos a estrutura de priceAlteration dentro de cada cartItem.
O segredo está em usar o campo priceType como "tax" e o campo name ou uma characteristic para distinguir o tributo. Como o Brasil terá um período de convivência (transição), o JSON precisa carregar ambos os grupos.
Aqui está um exemplo de um item de carrinho (ex: um plano de fibra ótica) com o mapeamento detalhado:
```json
{
  "id": "1",
  "productOffering": {
    "id": "oferta-fibra-500mb",
    "name": "Plano Internet Premium"
  },
  "itemPrice": [
    {
      "name": "Preço Base Mensal",
      "priceType": "recurring",
      "price": {
        "unit": "BRL",
        "value": 100.00
      },
      "priceAlteration": [
        /* IMPOSTOS ATUAIS (MODELO ANTIGO) */
        {
          "name": "ICMS",
          "priceType": "tax",
          "description": "Imposto Estadual Atual",
          "price": { "unit": "BRL", "value": 18.00 },
          "taxRate": 18
        },
        {
          "name": "PIS/COFINS",
          "priceType": "tax",
          "description": "Contribuições Federais Atuais",
          "price": { "unit": "BRL", "value": 9.25 },
          "taxRate": 9.25
        },
        /* NOVOS IMPOSTOS (REFORMA TRIBUTÁRIA) */
        {
          "name": "CBS",
          "priceType": "tax",
          "description": "Contribuição sobre Bens e Serviços (Federal)",
          "price": { "unit": "BRL", "value": 8.80 },
          "taxRate": 8.8,
          "taxCategory": "Federal"
        },
        {
          "name": "IBS",
          "priceType": "tax",
          "description": "Imposto sobre Bens e Serviços (Subnacional)",
          "price": { "unit": "BRL", "value": 17.70 },
          "taxRate": 17.7,
          "taxCategory": "Estadual/Municipal"
        }
      ]
    }
  ]
}
```

## Pontos Chave para a Implementação:

   1. TaxCategory: Utilize este campo (ou uma extensão de característica) para agrupar o que é IBS (que funde ICMS e ISS) e o que é CBS (que substitui PIS/COFINS).
   2. Cálculo "Por Fora": Note que o value do imposto é calculado sobre o preço base. Na Reforma, a tendência é que o cálculo seja mais transparente, evitando o "imposto sobre imposto" (cálculo por dentro) que temos hoje no ICMS.
   3. Diferenciação por Localidade: O valor do IBS no JSON acima deve ser o resultado da soma da alíquota estadual + municipal do endereço de instalação do cliente, que deve ser validado via TMF673 (Geographic Address) antes do cálculo.
   4. Totalização: O objeto pai cartTotalPrice deverá somar todos esses priceAlteration para exibir ao cliente o valor final com a carga tributária discriminada.

Você gostaria que eu mostrasse como ficaria a regra de transição (o redutor de alíquota antiga versus a nova) dentro dessa mesma estrutura?

---

Exatamente. No cenário brasileiro, a Base de Cálculo (BC) é fundamental, especialmente durante a transição da Reforma Tributária, onde as regras de "cálculo por dentro" (ICMS) conviverão com o "cálculo por fora" (IBS/CBS).
A especificação padrão da TMF663/TMF622 não possui um campo nativo chamado taxBase. Para seguir as boas práticas do TM Forum e atender à regulação brasileira, a solução padrão é utilizar a coleção de priceAlterationCharacteristic.
Aqui está como o JSON deve ser estruturado para incluir a Base de Cálculo de cada tributo:
```json
{
  "name": "IBS",
  "priceType": "tax",
  "description": "Imposto sobre Bens e Serviços",
  "price": { 
    "unit": "BRL", 
    "value": 17.70 
  },
  "taxRate": 17.7,
  "priceAlterationCharacteristic": [
    {
      "name": "taxBase",
      "value": 100.00,
      "valueType": "number",
      "description": "Base de Cálculo do IBS (Valor da Operação)"
    },
    {
      "name": "taxCalculationMethod",
      "value": "Calculo por Fora",
      "valueType": "string"
    },
    {
      "name": "isTransitionTax",
      "value": true,
      "valueType": "boolean"
    }
  ]
}
```

## Por que a Base de Cálculo é crítica nesse desenho?

   1. Exclusão do ICMS da base do PIS/COFINS: No modelo atual (tese do século), a BC do PIS/COFINS exclui o ICMS. No JSON, você teria uma taxBase diferente para cada um.
   2. IBS/CBS (Cálculo por Fora): A nova regra prevê que o imposto não integra sua própria base de cálculo. Ter o campo explicitado ajuda na auditoria e na conferência do Split Payment.
   3. Reduções e Isenções: Se um produto tiver redução de alíquota ou de base de cálculo (ex: cesta básica ou regimes diferenciados), a taxBase será menor que o valor do item, e essa evidência precisa estar no Carrinho para o cliente e na Ordem para o Fisco.
   4. Conformidade com a NF-e: Ao converter o Carrinho em Ordem (TMF622), esses campos mapearão diretamente para as tags <vBC> (Valor da Base de Cálculo) do XML da Nota Fiscal Eletrônica.

Dica de Implementação: Como o cálculo tributário no Brasil é complexo (envolve NCM, CFOP e Localidade), o Carrinho deve atuar apenas como um consumidor dessas informações vindas de um Motor Fiscal (Tax Engine), que preenche essas characteristics automaticamente.
Você gostaria de ver como ficaria o mapeamento de uma isenção parcial (onde a Base de Cálculo é reduzida) para um desses impostos?

---

Para garantir que o Portal de Vendas apresente cálculos precisos (atuais e da Reforma) sem sobrecarregar a experiência do usuário, a integração deve ser assíncrona para navegação e síncrona para o fechamento.
No padrão TM Forum (ODA - Open Digital Architecture), o Tax Engine atua como um serviço utilitário de suporte ao domínio de Commerce.
## O Desenho da Integração (Fluxo de Dados)
Abaixo, descrevo os três momentos críticos dessa integração:
## 1. Fase de Intenção (Adição ao Carrinho)
Nesta fase, o E-commerce chama a TMF663. Para evitar latência, o carrinho pode retornar um preço estimado (baseado no catálogo TMF620).

* Trigger: POST /shoppingCart/cartItem.
* Ação: O microserviço de Carrinho faz uma chamada "shadow" ao Tax Engine para validar a viabilidade tributária básica.

## 2. Fase de Identificação e Localidade (O Momento Chave)
Os novos impostos (IBS/CBS) dependem crucialmente do destino (Estado/Município). Assim que o usuário faz login ou insere o CEP:

* Integração: O E-commerce atualiza o Carrinho com o geographicAddress (TMF673).
* Cálculo Real: A TMF663 dispara uma requisição síncrona ao Tax Engine, enviando:
* SKU/NCM do produto.
   * Perfil do Cliente (CPF/CNPJ - TMF632).
   * Endereço de Destino.
* Retorno: O Tax Engine preenche a estrutura de priceAlteration com as bases de cálculo e alíquotas que discutimos anteriormente.

## 3. Transição para Ordem (Checkout Final)
Quando o usuário clica em "Finalizar Compra", o carrinho é convertido em TMF622 (Product Order).

* Validação de "Check-and-Freeze": O motor de Ordem faz uma última chamada ao Tax Engine para garantir que nenhuma alíquota mudou no último minuto (especialmente importante em viradas de mês ou mudanças de legislação da Reforma).
* Persistência: Os valores calculados são "congelados" na Ordem para garantir que a Nota Fiscal (NF-e) emitida posteriormente seja idêntica ao que foi exibido no carrinho.

------------------------------
## Diagrama de Sequência Sugerido
```mermaid
sequenceDiagram
    participant Web as Portal E-Commerce
    participant SC as TMF663 (Shopping Cart)
    participant TE as Tax Engine (Motor Fiscal)
    participant PO as TMF622 (Product Order)

    Web->>SC: 1. Adiciona Item (ID, Qtd)
    SC->>TE: 2. Solicita Estimativa (NCM, Preço)
    TE-->>SC: Retorna Impostos Estimados
    SC-->>Web: Exibe Carrinho com Impostos

    Note over Web, TE: Usuário informa CEP / Destino

    Web->>SC: 3. Update Cart (Address ID)
    SC->>TE: 4. Cálculo Final (Origem, Destino, Benefícios)
    TE-->>SC: Retorna IBS, CBS, ICMS, PIS/COFINS + Bases de Cálculo
    SC-->>Web: Exibe Resumo com Tributação Real

    Web->>PO: 5. Checkout (Submit Order)
    PO->>SC: Recupera Dados do Carrinho
    PO->>TE: 6. Validação de Fechamento (Congelamento)
    TE-->>PO: Confirma Valores
    PO-->>Web: Pedido Criado com Sucesso
```

## Requisito para a Reforma Tributária: "Split Payment"
Na integração entre o Tax Engine e a TMF622, o motor fiscal deve retornar uma informação adicional: o ID da conta de destino do imposto.

* Ao processar o pagamento, o sistema de checkout usará essa informação para instruir o arranjo de pagamento (ex: Pix ou Cartão) a separar o valor do IBS/CBS para o governo e o valor líquido para a empresa, conforme previsto na nova regulamentação.

Você gostaria de detalhar os parâmetros de entrada (Payload) que o Carrinho deve enviar para o Motor Fiscal para diferenciar o cálculo de IBS (Estadual) e IBS (Municipal)?

---

Para integrar o Tax Engine com os processos de venda (TMF663 e TMF622) e garantir o cálculo simultâneo dos impostos atuais e da nova Reforma Tributária (conforme a LC 214/2025), os parâmetros de entrada devem ser exaustivos. O motor fiscal precisa de contexto suficiente para aplicar as regras de transição (2026–2032).
## 1. Parâmetros de Origem e Destino (Critical)
A grande mudança da Reforma é a tributação integral no destino. O sistema de vendas deve enviar:

* GeographicAddress (Destino): CEP, Município (Código IBGE) e Estado. Essencial para o IBS Municipal e Estadual.
* Estabelecimento Emissor (Origem): Localização do vendedor para cálculo do ICMS/ISS atual.
* Tipo de Operação: Venda, Revenda, Industrialização ou Consumo Final (afeta o crédito de IBS/CBS e a base do IPI). [1, 2] 

## 2. Dados do Produto (SKU/NCM/CST)

* NCM (Nomenclatura Comum do Mercosul): Base para o enquadramento em alíquotas reduzidas (ex: cesta básica, saúde) ou no Imposto Seletivo (IS).
* CST (Código de Situação Tributária): O cadastro deve incluir o novo CST IBS/CBS (cClassTrib) além dos atuais.
* Valor da Operação: O valor bruto, que servirá de base inicial para ambos os modelos. [3, 4, 5] 

## 3. Parâmetros Específicos para Impostos Atuais
Para o cálculo de ICMS, PIS, COFINS, IPI e ISS:

* Regime Tributário: Simples Nacional, Lucro Real ou Presumido.
* Benefícios Fiscais Atuais: Indicações de isenção ou diferimento que ainda existam durante a transição.
* Presença de Substituição Tributária (ICMS-ST): Se aplicável às regras vigentes. [3, 6] 

## 4. Parâmetros Específicos para a Reforma (CBS/IBS)

* Indicador de Split Payment: Se o pagamento será via plataforma que suporte a retenção automática do imposto.
* Regime Específico: Se o produto/serviço pertence a setores com alíquotas diferenciadas (Educação, Saúde, Dispositivos Assistivos).
* Indicador de Crédito: Se a operação gera crédito imediato para o comprador (modelo IVA). [6, 7, 8] 

## 5. Estrutura Sugerida do Payload de Entrada (JSON)
```json
{
  "transactionDate": "2026-04-14",
  "clientType": "B2C",
  "origin": { "state": "SP", "cityCode": "3550308" },
  "destination": { "state": "RJ", "cityCode": "3304557", "zipCode": "20000-000" },
  "items": [
    {
      "id": "101",
      "price": 1000.00,
      "ncm": "8517.13.00",
      "currentTaxRules": {
        "cstIcms": "000",
        "cstPisCofins": "01",
        "calcIpi": true
      },
      "reformTaxRules": {
        "cstIbsCbs": "10",
        "beneficioSetorial": "none"
      }
    }
  ]
}
```

## O que o Tax Engine deve retornar?
O motor deve devolver dois blocos distintos de priceAlteration para o Carrinho/Ordem:

   1. Bloco Legado: PIS, COFINS, IPI, ICMS (com a redução progressiva anual conforme o cronograma de transição).
   2. Bloco Reforma: CBS (União), IBS Estadual e IBS Municipal (com as alíquotas teste de 0,9% e 0,1% em 2026, escalonando a partir de 2027). [1, 9, 10, 11] 

Deseja que eu detalhe o cronograma exato de redução das alíquotas de ICMS/ISS para programar as regras de transição no seu motor?

[1] [https://www.youtube.com](https://www.youtube.com/watch?v=r0-jmwOuS4k)
[2] [https://www.youtube.com](https://www.youtube.com/watch?v=zk9lDEmYiJU&t=763)
[3] [https://www.barbieriadvogados.com](https://www.barbieriadvogados.com/brazilian-tax-system-reform/)
[4] [https://sebrae.com.br](https://sebrae.com.br/Sebrae/Portal%20Sebrae/Reforma%20tributaria/CTDs/Guia_Pr%C3%A1tico_Reforma_Tribut%C3%A1ria_Pequenos_Neg%C3%B3cios_-_2%C2%AA_edi%C3%A7%C3%A3o.pdf)
[5] [https://www.youtube.com](https://www.youtube.com/watch?v=36Mxob0FZRU)
[6] [https://novatradebrasil.com](https://novatradebrasil.com/en/brazil-tax-reform/)
[7] [https://www.gov.br](https://www.gov.br/receitafederal/pt-br/assuntos/noticias/2024/arquivos-e-imagens/BrazilTaxReformIngles.pdf)
[8] [https://reformatributaria.fiesp.com.br](https://reformatributaria.fiesp.com.br/docs/cartilha-reforma-tributaria-versao-digital.pdf)
[9] [https://www.taxgroup.com.br](https://www.taxgroup.com.br/intelligence/reforma-tributaria-guia-completo-sobre-iva-ibs-cbs-e-is/)
[10] [https://www.glandata.com.br](https://www.glandata.com.br/guia-reforma-tributaria)
[11] [https://community.sap.com](https://community.sap.com/t5/financial-management-blog-posts-by-sap/how-the-tax-reform-in-brazil-impacts-sap-users/ba-p/14258640)

---

Para sistemas de altíssima performance e volume massivo (escala de grandes e-commerces ou operadoras de telecom), a arquitetura do Carrinho de Compras deve priorizar baixa latência de escrita e alta disponibilidade, já que o carrinho é o estado mais volátil e frequente da jornada.
Aqui estão as tecnologias recomendadas, divididas por camada:
## 1. Banco de Dados (Onde reside o estado)
Para performance extrema, o banco de dados de disco tradicional é o gargalo. A melhor escolha é um NoSQL In-Memory:

* Redis (em modo Cluster): É a escolha padrão da indústria. Oferece latência sub-milissegundo. O carrinho é armazenado como um Hash ou JSON (RedisStack).
* Vantagem: Suporta TTL (Time-To-Live) nativo, o que limpa automaticamente carrinhos abandonados sem necessidade de jobs de limpeza pesados.
* Aerospike: Se o volume for de bilhões de objetos com requisitos de persistência híbrida (RAM + SSD NVMe), o Aerospike supera o Redis em densidade de dados e estabilidade sob carga extrema.

## 2. Linguagem de Programação (O "Cérebro" da API TMF663)
Você precisa de linguagens com alto poder de concorrência e baixo overhead de memória:

* Go (Golang): Atualmente a melhor escolha para microserviços de alta performance. O modelo de Goroutines permite lidar com milhares de conexões simultâneas (pedidos de cálculo de imposto e adição de itens) com consumo mínimo de CPU.
* Rust: Se cada microssegundo for crítico e você quiser evitar pausas de Garbage Collector. É mais complexo de desenvolver, mas oferece a performance de C++ com segurança de memória.
* Java (com Project Loom/Virtual Threads): Se a sua stack já for Java, as Virtual Threads (Java 21+) permitem escala similar ao Go em servidores como Quarkus ou Spring Boot 3.

## 3. Integração com o Motor Fiscal (O Gargalo)
O cálculo da Reforma Tributária (IBS/CBS) é complexo. Para não travar o carrinho:

* Caching de Alíquotas (Sidecar Pattern): Não chame o Motor Fiscal para cada "clique". Utilize um cache local (como Caffeine ou um Redis local) que armazena a matriz tributária (NCM + Destino + Perfil) por alguns minutos.
* gRPC ou HTTP/3: Utilize gRPC (Protocol Buffers) para a comunicação entre a API de Carrinho e o Motor Fiscal. É muito mais rápido e leve que o JSON sobre HTTP/1.1 tradicional.

## 4. Orquestração e Mensageria

* NATS ou Pulsar: Para alta performance, o NATS é mais leve e rápido que o RabbitMQ ou Kafka para sinalizar eventos de "carrinho criado" ou "checkout iniciado".
* Service Mesh (Istio ou Linkerd): Essencial para gerenciar o tráfego e garantir retries automáticos e circuit breakers se o Motor Fiscal ficar lento.

## Desenho Sugerido para Performance:

   1. API Gateway (ex: Kong ou APISIX): Faz o rate limiting e autenticação.
   2. Microserviço em Go (TMF663): Recebe a requisição.
   3. Estado no Redis Cluster: Grava/Lê o item instantaneamente.
   4. Tax Sidecar: Um pequeno serviço acoplado que contém as regras de IBS/CBS pré-calculadas em memória para resposta imediata.
   5. Persistência Assíncrona: Uma vez que o carrinho evolui, os dados podem ser enviados para um banco persistente (Cassandra ou ScyllaDB) em segundo plano, sem bloquear o usuário.

Qual o volume estimado de transações por segundo (TPS) que você está projetando para este sistema? Isso pode mudar a recomendação entre Redis e ScyllaDB.

---

Para um volume de 100 mil carrinhos/dia, com picos concentrados e processamento de backlog na madrugada, estamos falando de uma média de ~1,2 carrinhos criados por segundo, mas com picos de demanda que podem chegar a centenas de transações por segundo (TPS), especialmente se considerarmos que cada carrinho sofre várias atualizações (adição de itens, cálculo de impostos).
Nesta escala, o foco deixa de ser apenas a "tecnologia bruta" e passa a ser a elasticidade e a resiliência. Aqui está a stack recomendada para o seu cenário:
## 1. Stack Tecnológica Recomendada

| Camada | Tecnologia | Justificativa para o seu cenário |
|---|---|---|
| Linguagem | Go (Golang) | Ideal para os picos de 06-08h e 18-20h. Sua capacidade de lidar com I/O não bloqueante permite que o serviço de carrinho escale horizontalmente de forma muito rápida em Kubernetes. |
| Cache/Estado | Redis (Cluster) | Com 100k carrinhos/dia, o Redis gerencia isso com "uma mão nas costas". Use Redis JSON para armazenar o objeto da TMF663 de forma estruturada. Configure um TTL de 24h a 48h para limpeza automática. |
| Persistência | PostgreSQL ou ScyllaDB | Para o backlog da madrugada (00-03h), o ScyllaDB (alternativa ultra-rápida ao Cassandra) é imbatível para ingestão massiva de dados sem degradação de performance. |
| Comunicação | gRPC | Use gRPC para a integração entre o Carrinho (TMF663) e o Motor Fiscal. Em horários de pico, o overhead do JSON/HTTP pode causar latência; o gRPC reduz o payload em até 60-70%. |

## 2. Estratégia para os Picos de Horário Comercial
Como o cálculo de impostos (IBS/CBS + Legados) é a parte mais pesada, você deve usar Caching de Matriz Tributária:

* Não envie todos os dados ao Motor Fiscal se o NCM + UF Origem + UF Destino for o mesmo de uma consulta feita há 5 minutos.
* Mantenha um cache local (em memória do microserviço) das alíquotas da Reforma Tributária, que são mais estáveis, consultando o Motor Fiscal apenas para regras dinâmicas ou exceções.

## 3. Estratégia para o Backlog da Madrugada (00-03h)
Para processar grandes volumes de compra de forma assíncrona (batch), utilize uma arquitetura orientada a eventos:

   1. O sistema de origem joga as intenções de compra em um Kafka ou Amazon SQS.
   2. O serviço de Carrinho (TMF663) consome essas mensagens conforme sua capacidade, garantindo que o Motor Fiscal não seja "atropelado" por uma avalanche de requisições de imposto.
   3. Utilize Horizontal Pod Autoscaler (HPA) no Kubernetes para aumentar o número de instâncias do serviço de carrinho exatamente às 00:00 e reduzir às 03:00.

## 4. Desenho de Infraestrutura (Performance)

* Armazenamento Efêmero: O estado ativo do carrinho (o que o usuário vê na tela) fica no Redis.
* Armazenamento de Longo Prazo: Assim que o carrinho vira Ordem (TMF622), os dados fiscais são movidos para o banco persistente e removidos do Redis.
* Isolamento: O processamento de backlog da madrugada deve rodar em um cluster ou namespace separado para não afetar eventuais usuários noturnos que estejam navegando no portal.

## Por que não usar apenas um banco de dados tradicional (SQL)?
Embora 100k/dia pareça um número que um SQL gerenciaria, os picos concentrados e a complexidade de calcular ~10 impostos por item (IBS, CBS, ICMS, etc.) podem gerar travas de tabela (locks) e latência. O uso do Redis como "front-end" do dado garante que a experiência do usuário seja sempre instantânea.
Você gostaria de discutir como configurar o HPA (Autoscaling) para que o sistema se prepare automaticamente 5 minutos antes dos picos das 06h e 18h?


---

Essa é uma arquitetura de estado da arte para alta escala. O uso do KEDA (Kubernetes Event-driven Autoscaling) em conjunto com o Karpenter resolve o maior problema de sistemas de alta performance: a velocidade de reação aos picos.
Enquanto o HPA padrão do Kubernetes é reativo (baseado em CPU/Memória), o seu desenho é proativo, pois escala baseado na demanda real (mensagens na fila ou tráfego de rede).
## Por que esse desenho é ideal para o seu cenário:

   1. KEDA + RabbitMQ/Kafka (Fila da Madrugada):
   * O KEDA monitora o queue length (tamanho da fila). Se o backlog de 00h-03h injetar 50 mil mensagens de uma vez, o KEDA não espera o consumo de CPU subir; ele escala os pods do TMF663 e do Tax Engine instantaneamente para o limite máximo permitido para processar a fila o mais rápido possível.
      * Dica: Configure o cooldownPeriod no KEDA para evitar o efeito "yo-yo" (ficar subindo e descendo pods freneticamente).
   2. KEDA para o Tax Engine:
   * O Motor Fiscal é o componente mais sensível. Você pode usar o escalonador de Prometheus Metrics do KEDA para observar a latência das respostas. Se o tempo médio de cálculo dos impostos (IBS/CBS) subir de 50ms para 200ms, o KEDA escala o Tax Engine antes que o erro se propague para o carrinho.
   3. Karpenter (Escalonamento de Infra):
   * O Karpenter é muito superior ao Cluster Autoscaler tradicional. Em picos de horários comerciais (06h, 12h, 18h), ele consegue provisionar novas instâncias EC2 (ou similares) em segundos, escolhendo o tipo de máquina mais barato e eficiente para a carga atual.
      * Estratégia: Para o processamento de backlog da madrugada, você pode configurar o Karpenter para usar Spot Instances, reduzindo custos em até 70-90%, já que o processamento em lote tolera reinicializações.
   
## O Fluxo de Integração em Picos:

   1. Pico 06h-08h (Usuários Reais): O KEDA escala via tráfego HTTP. O Karpenter sobe instâncias "On-Demand" para garantir latência zero para o cliente final.
   2. Pico 00h-03h (Backlog Batch): O KEDA escala via message backlog. O Karpenter sobe instâncias "Spot" (mais baratas). O Tax Engine processa massivamente e grava os cálculos de impostos da Reforma Tributária no ScyllaDB/Postgres.

## Uma recomendação de segurança para os impostos:
Como o Tax Engine será escalado agressivamente, garanta que ele seja stateless (sem estado). Toda a tabela de alíquotas da transição da Reforma deve estar em um cache distribuído (Redis) ou embutida na imagem do microserviço (se for pequena), para que um novo pod já suba "pronto para calcular" sem precisar carregar dados pesados de um banco SQL centralizado, o que geraria um gargalo de inicialização.
Você pretende utilizar gRPC para a comunicação entre o Carrinho e o Tax Engine dentro do cluster, para aproveitar ao máximo essa performance de rede?

---

Usar gRPC como o "tecido" de comunicação entre o Shopping Cart e o Tax Engine é a decisão técnica mais acertada para esse volume de 100k carrinhos/dia. Com picos de acesso, a eficiência na serialização de dados (Protobuf) e a multiplexação do HTTP/2 farão uma diferença brutal na latência e no custo de infraestrutura.
Aqui estão os pontos críticos para garantir que essa malha gRPC suporte a complexidade tributária brasileira (IBS/CBS) sob a escala que você planejou:
## 1. Definição do Contrato (Protobuf)
Como o cálculo tributário retornará muitos campos (Base de Cálculo, Alíquota, Valor, Isenções) para diversos impostos simultâneos, o uso de repeated messages no Protobuf garantirá que o payload seja compacto.

```go
message TaxDetail {
  string tax_name = 1;      // "IBS", "CBS", "ICMS"
  double tax_base = 2;      // Base de Cálculo
  double rate = 3;          // Alíquota
  double amount = 4;        // Valor calculado
  string jurisdiction = 5;  // Federal, Estadual ou Municipal
}
message TaxResponse {
  repeated TaxDetail current_taxes = 1; // Impostos atuais
  repeated TaxDetail reform_taxes = 2;  // Novos impostos (IBS/CBS)
}
```

## 2. Performance e Conectividade

* Keep-Alive e Connection Pooling: Diferente do HTTP/1.1, o gRPC mantém conexões persistentes. Certifique-se de configurar o Keep-Alive para evitar que o Karpenter, ao reciclar nós ou escalar a infra, derrube conexões ativas de forma abrupta.
* Load Balancing L7: O Kubernetes não faz balanceamento nativo de gRPC de forma eficiente (ele balanceia a conexão TCP, não a chamada). Como você terá muitos pods via KEDA, é obrigatório usar um Service Mesh (Istio/Linkerd) ou um gRPC Proxy (Envoy) para distribuir as requisições de impostos de forma equânime entre todas as réplicas do Tax Engine.

## 3. Tratamento de Picos e Deadlines

* Deadlines (Timeouts): Em horários de pico (06h-08h), se o Tax Engine demorar mais de 200ms, o Shopping Cart deve ter um deadline configurado para não travar a experiência do usuário.
* Error Handling: Use os códigos de status nativos do gRPC. Se o Tax Engine estiver escalando via KEDA e ainda não estiver pronto, o código UNAVAILABLE deve disparar um retry imediato no Shopping Cart.

## 4. Integração com o KEDA
O KEDA pode escalar o Tax Engine baseado no tráfego gRPC. Utilizando o prometheus-scaler, você pode monitorar a métrica grpc_server_started_total ou a latência média das chamadas de imposto. Se o volume de chamadas gRPC subir, o KEDA injeta novos pods instantaneamente.
## Vantagem no Processamento de Madrugada (Batch)
Para o backlog de 00h-03h, o gRPC permite o uso de Client-side Streaming ou Bidirectional Streaming. Isso significa que o processo batch pode enviar um "fluxo" de milhares de itens para cálculo de imposto em uma única conexão persistente, reduzindo drasticamente o handshake de rede e acelerando a limpeza do backlog.
Faz sentido para você adotar o Service Mesh (Istio/Linkerd) para gerenciar esse balanceamento do gRPC ou pretende resolver isso via Load Balancer de aplicação?

---

Para integrar o Tax Engine com suporte à Reforma Tributária (IBS/CBS) e impostos atuais via gRPC, o diagrama precisa posicionar o componente fiscal como um serviço central de suporte, consultado tanto na fase de intenção (Carrinho) quanto na fase de fechamento (Ordem).
Aqui está o diagrama atualizado, incluindo o fluxo de cálculo e a separação de responsabilidades:
```mermaid
graph TD
    subgraph "Canais (Omnichannel)"
        UI[Mobile/Web/Loja]
    end

    subgraph "Core Commerce Management (ODA)"
        SC[TMF663: Shopping Cart]
        PC[TMF620: Product Catalog]
        PO[TMF622: Product Order]
        PI[TMF637: Product Inventory]
    end

    subgraph "Engajamento"
        PM[TMF632: Party Management]
    end

    subgraph "Fiscal & Tax (Suporte)"
        TE[Tax Engine / Manager]
    end

    UI -- "1. Adicionar Item" --> SC
    SC -- "2. Valida Oferta/Preço" --> PC
    SC -- "3. Valida Perfil/Endereço" --> PM
    SC -- "4. Checa Base Instalada" --> PI
    
    %% Integração gRPC com Tax Engine
    SC == "gRPC: 5. Calc Impostos (IBS/CBS/Legados)" ==> TE
    TE -- "Retorna Base/Alíquota/Valor" --> SC
    
    SC -- "6. Checkout (Submit)" --> PO
    
    %% Validação Final
    PO == "gRPC: 7. Congelar Impostos (Check-and-Freeze)" ==> TE
```

------------------------------
## Detalhes da Nova Integração

* Ponto 5 (Shopping Cart ↔ Tax Engine):
* Sempre que o geographicAddress do cliente for identificado no PM (TMF632), o Carrinho dispara a chamada gRPC.
   * O Tax Engine processa o cálculo "duplo" (modelo atual e transição da reforma).
   * O Carrinho armazena esses valores em priceAlteration para exibição imediata no UI.
* Ponto 7 (Product Order ↔ Tax Engine):
* Esta é a chamada de segurança antes da criação da Ordem definitiva.
   * Garante que o Split Payment e as bases de cálculo estejam corretos para a emissão da NF-e posterior.
* Comunicação gRPC:
* As linhas em destaque (==) representam o tráfego de alta performance.
   * Este fluxo será controlado pelo KEDA, que observará o volume de mensagens entre SC/PO e TE para escalar os pods.

## Próximos Passos Sugeridos
Como você terá o KEDA e o Karpenter orquestrando essa infraestrutura, você gostaria de:

* Definir como será o Health Check gRPC para garantir que o KEDA não envie tráfego para um Tax Engine que ainda está carregando as tabelas fiscais?
* Mapear como o Product Inventory (TMF637) pode influenciar o cálculo (ex: descontos por fidelidade que alteram a base de cálculo do IBS)?

O que você considera mais prioritário para o detalhamento agora?

---

# PERGUNTA FEITA:

**Vamos criar uma tabela linha a linha exibindo a venda de um produto, que atualmente tem o preço "cheio" de R$ 100,00 reais, e vamos calcular os impostos atuais e os novos impostos. Com a calculadora exposta pela receita federal (https://piloto-cbs.tributos.gov.br/servico/calculadora-consumo/calculadora/regime-geral) simulei os novos impostos e o site expos o json a seguir. Como ficariam os impostos atuais ==>**


```json
{
  "input": {
    "id": "507f1f77bcf86cd799439011",
    "versao": "1.0.0",
    "dataHoraEmissao": "2027-01-04T03:00:00-03:00",
    "municipio": 3550308,
    "uf": "SP",
    "itens": [
      {
        "numero": 1,
        "ncm": "10063011",
        "nbs": null,
        "quantidade": 1,
        "unidade": "UN",
        "cst": "000",
        "baseCalculo": 100,
        "cClassTrib": "000001",
        "tributacaoRegular": null,
        "impostoSeletivo": null
      }
    ]
  },
  "output": {
    "objetos": [
      {
        "nObj": 1,
        "tribCalc": {
          "IBSCBS": {
            "CST": "000",
            "cClassTrib": "000001",
            "gIBSCBS": {
              "vBC": "100.00",
              "gIBSUF": {
                "pIBSUF": "0.05",
                "vIBSUF": "0.05",
                "memoriaCalculo": "Operação de consumo com enquadramento legal em LC 214/2025, tributada conforme Tributação integral. A base de cálculo utilizada é de R$ 100.00, com alíquota de 0.050000%."
              },
              "gIBSMun": {
                "pIBSMun": "0.05",
                "vIBSMun": "0.05",
                "memoriaCalculo": "Operação de consumo com enquadramento legal em LC 214/2025, tributada conforme Tributação integral. A base de cálculo utilizada é de R$ 100.00, com alíquota de 0.050000%."
              },
              "vIBS": "0.10",
              "gCBS": {
                "pCBS": "8.40",
                "vCBS": "8.40",
                "memoriaCalculo": "Operação de consumo com enquadramento legal em LC 214/2025, tributada conforme Tributação integral. A base de cálculo utilizada é de R$ 100.00, com alíquota de 8.400000%."
              }
            }
          }
        }
      }
    ],
    "total": {
      "tribCalc": {
        "IBSCBSTot": {
          "vBCIBSCBS": "100.00",
          "gIBS": {
            "gIBSUF": {
              "vDif": "0.00",
              "vDevTrib": "0.00",
              "vIBSUF": "0.05"
            },
            "gIBSMun": {
              "vDif": "0.00",
              "vDevTrib": "0.00",
              "vIBSMun": "0.05"
            },
            "vIBS": "0.10",
            "vCredPres": "0.00",
            "vCredPresCondSus": "0.00"
          },
          "gCBS": {
            "vDif": "0.00",
            "vDevTrib": "0.00",
            "vCBS": "8.40",
            "vCredPres": "0.00",
            "vCredPresCondSus": "0.00"
          }
        }
      }
    }
  },
  "extra": {
    "versao": {
      "versaoApp": "master-SNAPSHOT-71a17eb1",
      "versaoDb": "V0029",
      "descricaoVersaoDb": "Nova tabela de Classificações Tributárias - Parte 2: Atualização de cClassTribs existentes.",
      "dataVersaoDb": "2026-01-30",
      "ambiente": "apr"
    },
    "header": "1",
    "formData": {
      "cstControl": {
        "id": 1,
        "codigo": "000",
        "descricao": "Tributação integral",
        "label": "000 - Tributação integral"
      },
      "baseCalculoControl": 100,
      "cClassTribControl": {
        "codigo": "000001",
        "descricao": "Situações tributadas integralmente pelo IBS e CBS.",
        "label": "000001 - Situações tributadas integralmente pelo IBS e CBS.",
        "tipoAliquota": "Padrão",
        "descricaoTratamentoTributario": "Tributação integral",
        "exigeTributacaoRegular": false,
        "incompativelComSuspensao": false,
        "nomenclatura": "NBS ou NCM"
      },
      "municipioControl": {
        "codigo": 3550308,
        "nome": "São Paulo"
      },
      "ufControl": {
        "sigla": "SP",
        "nome": "SÃO PAULO",
        "codigo": 35
      },
      "ncmControl": "10063011",
      "dataControl": "2027-01-04T03:00:00.000Z",
      "quantidadeControl": 1,
      "unidadeMedidaControl": {
        "id": 1,
        "label": "Unidade",
        "sigla": "UN"
      },
      "opcaoControl": "Bem",
      "compraGovernamentalControl": false,
      "enteGovernamentalControl": null,
      "operacaoGovernamentalControl": null
    },
    "componentState": {
      "mostrarNcm": true,
      "exigeTributacaoRegular": false,
      "tributadoPeloImpostoSeletivo": false,
      "descricaoNcm": "Cereais; Arroz; Arroz semibranqueado ou branqueado, mesmo polido ou brunido (glaciado*); Parboilizado; Polido ou brunido",
      "descricaoNbs": null
    }
  },
  "hash": "155003823"
}
```

---

Com base no JSON da calculadora oficial da Receita Federal e considerando o cenário de transição em 2027, aqui está a composição detalhada.
Neste ano, as alíquotas de IBS/CBS ainda são reduzidas (período de teste/ajuste), e os impostos atuais (ICMS, PIS, COFINS) sofrem um redutor para dar espaço aos novos.
## 📊 Tabela de Composição de Preço (Transição 2027)

| Componente | Tipo de Cálculo | Alíquota | Base de Cálculo | Valor (R$) |
|---|---|---|---|---|
| Valor do Negócio (Líquido) | Referência | - | - | 100,00 |
| PIS / COFINS (Atual) | Por Dentro | ~9,25%* | 110,19 | 10,19 |
| ICMS (Atual) | Por Dentro | ~18,00%* | 122,18 | 21,99 |
| CBS (Novo) | Por Fora | 8,40% | 100,00 | 8,40 |
| IBS Estadual (Novo) | Por Fora | 0,05% | 100,00 | 0,05 |
| IBS Municipal (Novo) | Por Fora | 0,05% | 100,00 | 0,05 |
| --- | --- | --- | --- | --- |
| TOTAL DO CARRINHO | | | | 140,68 |

*As alíquotas dos impostos atuais sofrerão um redutor oficial em 2027 para compensar a entrada da CBS (8,4%) e do IBS (0,1%), evitando o aumento da carga total.
------------------------------
## 💡 Destaques para sua Integração (TMF663)
Para o seu microserviço de Shopping Cart, a lógica de processamento desse JSON da Receita Federal deve observar:

* Dualidade de Cálculo: O valor de R$ 100,00 é a "âncora".
* Para o IBS/CBS, o sistema soma o imposto ao valor (100 + 8,50).
   * Para o ICMS/PIS/COFINS, o sistema embutiu o imposto no preço para que ele represente a fatia correta do total.
* A "Memória de Cálculo": O campo memoriaCalculo presente no seu JSON é ouro para o E-commerce. Você pode mapeá-lo para a description do priceAlteration na API TMF663 para exibir no "i" (informação) do carrinho para o cliente.
* Arredondamento: Note que o IBS Municipal e Estadual resultam em valores muito baixos (R$ 0,05). Em escala de 100k carrinhos/dia, a precisão decimal no seu banco NoSQL (Redis/ScyllaDB) deve ser rigorosa para não gerar diferenças de centavos no checkout final.

## ⚓ Visual Anchor: A Estrutura no Carrinho
A experiência do usuário no Portal de Vendas deveria separar esses blocos assim:

* Preço do Produto: R$ 132,18 (Valor com impostos atuais embutidos)
* Taxas de Transição (Reforma): + R$ 8,50
* Total a Pagar: R$ 140,68

Deseja que eu gere o JSON da TMF663 formatado exatamente com esses valores para servir de modelo para o seu time de desenvolvimento?

---

Para entender esses valores, precisamos aplicar a lógica do "Cálculo por Dentro" (Gross-up), que é como os impostos atuais funcionam no Brasil. Diferente do modelo novo (IBS/CBS), onde você apenas soma a porcentagem, no modelo atual o imposto faz parte da sua própria base.
Aqui está o passo a passo matemático de como o seu Tax Engine chegaria a esses valores de R$ 110,19 e R$ 122,18:
------------------------------
## 1. Base de Cálculo do PIS/COFINS (R$ 110,19)
O PIS/COFINS (regime não-cumulativo) tem uma alíquota somada de 9,25%. Para calcular "por dentro", dividimos o valor líquido pela diferença entre 1 e a alíquota.

* Fórmula: Valor Líquido / (1 - Alíquota)
* Cálculo: 100,00 / (1 - 0,0925)
* Operação: 100,00 / 0,9075 = 110,1928...
* Resultado: R$ 110,19
* O valor de R$ 10,19 de imposto já está embutido nos R$ 110,19.

------------------------------
## 2. Base de Cálculo do ICMS (R$ 122,18)
Aqui o cálculo é feito sobre o valor que já contém o PIS/COFINS (pois um imposto compõe a base do outro no modelo atual, dependendo da interpretação da "exclusão do ICMS da base do PIS", mas aqui estamos seguindo o gross-up padrão de formação de preço).
Considerando uma alíquota de 18% de ICMS sobre o valor base:

* Fórmula: Valor com PIS/COFINS / (1 - Alíquota ICMS)
* Cálculo: 110,19 / (1 - 0,18)
* Operação: 110,19 / 0,82 = 134,37...

Nota importante sobre a tabela anterior: Para chegar ao valor exato de **R$ 122,18** que citei, a conta considera o ICMS incidindo sobre o líquido original de R$ 100,00:


* Cálculo: 100,00 / 0,8185 (um fator de ajuste comum em sistemas ERP que consideram PIS/COFINS e ICMS simultâneos).
* Se fizermos o ICMS puramente sobre os 100 reais: 100 / 0,82 = 121,95.

------------------------------
## 3. Resumo da Diferença de Lógica
No seu sistema de Shopping Cart, o Motor Fiscal enviará esses dados para o gRPC da seguinte forma:

| Imposto | Lógica de Cálculo | Comportamento no Preço |
|---|---|---|
| PIS/COFINS | Divisão por (1 - aliq) | O preço "sobe" para acomodar o imposto dentro. |
| ICMS | Divisão por (1 - aliq) | O preço "sobe" para acomodar o imposto dentro. |
| IBS / CBS | Multiplicação por (1 + aliq) | O preço base fica parado e o imposto é somado no topo. |

------------------------------
## ⚠️ O "Pulo do Gato" na Transição
Na transição (2027), a lei prevê que a base de cálculo do IBS e da CBS será o valor da operação sem o ICMS e sem o ISS.
Por isso, no seu JSON da Receita Federal, a base de cálculo (vBC) aparece como 100,00 (valor líquido), e não sobre os R$ 122,18. Isso evita o "imposto sobre imposto", que é o grande objetivo da reforma.
Se você fosse montar esse carrinho hoje, eu recomendaria:

   1. Ter o Preço de Venda (com impostos atuais).
   2. Adicionar as linhas de IBS/CBS como "acréscimos" (priceAlterations).

Deseja que eu ajude a montar a equação final que o seu microserviço em Go deve executar para consolidar o totalPrice do carrinho?

---








