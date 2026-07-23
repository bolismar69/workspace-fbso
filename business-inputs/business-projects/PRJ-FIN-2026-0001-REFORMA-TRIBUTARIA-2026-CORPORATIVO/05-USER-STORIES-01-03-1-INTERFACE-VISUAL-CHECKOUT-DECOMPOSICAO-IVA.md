# Histórias de Usuário (User Stories) — Feature 03.1
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulos: E-commerce, Portal B2B, CRM (Onda 1)
- Feature Relacionada: 01.03.1 — Interface Visual de Checkout com Decomposição do IVA [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Redesenho do Resumo Financeiro do Carrinho com Segregação CBS e IBS

### 1. Descrição da História (Visão de Negócio)

Como Cliente do E-commerce,
Quero visualizar claramente no resumo do carrinho a distinção entre o preço líquido do produto e os tributos CBS (federal) e IBS (destino) calculados por fora,
Para entender exatamente quanto estou pagando pelo produto e quanto corresponde aos impostos da reforma tributária [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Layout Obrigatório de Decomposição): O resumo financeiro do carrinho deve seguir obrigatoriamente a estrutura visual:
  - Linha 1: "Valor dos Produtos/Serviços: R$ X.XXX,XX" (preço líquido, sem imposto)
  - Linha 2: "CBS — Contribuição Federal (X,X%): R$ XXX,XX"
  - Linha 3: "IBS — Imposto Estadual/Municipal (Y,Y% - [Nome do Município/UF]): R$ XXX,XX"
  - Linha 4 (se aplicável): "IS — Imposto Seletivo: R$ XXX,XX"
  - Linha final em destaque: "Total do Pedido: R$ X.XXX,XX"
* RN02 (Tooltip Educativo): Cada linha de imposto deve possuir um ícone "ℹ️" que, ao ser clicado ou tocado, exiba um tooltip com explicação resumida sobre o que é aquele tributo e por que ele é cobrado.
* RN03 (Segregação Visual): As linhas de tributos devem usar cor secundária (cinza) para diferenciar visualmente do preço base (preto) e do total (destaque verde ou azul).

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cliente visualiza decomposição completa do IVA no carrinho

* Dado que o cliente adicionou produtos ao carrinho no valor líquido de R$ 1.000,00;
* E informou CEP de entrega em Santana de Parnaíba (SP);
* Quando o carrinho for renderizado com os impostos calculados;
* Então o resumo financeiro deve exibir:
  - "Valor dos Produtos: R$ 1.000,00"
  - "CBS — Contribuição Federal (8,8%): R$ 88,00"
  - "IBS — Imposto Estadual/Municipal (17,7% - Santana de Parnaíba/SP): R$ 177,00"
  - "Total do Pedido: R$ 1.265,00" (em destaque)
* E cada linha de tributo deve ter o ícone "ℹ️" com tooltip explicativo funcional [INDEX].

#### Cenário 2: Produto com Imposto Seletivo (IS) aplicável

* Dado que o carrinho contém um produto da categoria "bebidas alcoólicas" sujeito ao IS;
* E o valor líquido total do carrinho é R$ 500,00;
* Quando o resumo for exibido;
* Então deve aparecer uma linha adicional: "IS — Imposto Seletivo (15,0%): R$ 75,00";
* E o Total do Pedido deve refletir a soma de todos os tributos: CBS + IBS + IS [INDEX].

------------------------------
## 📝 US-02: Layout da Proposta Comercial Impressa com Discriminação Fiscal

### 1. Descrição da História (Visão de Negócio)

Como Vendedor Externo que apresenta propostas presenciais a clientes corporativos,
Quero gerar um PDF de proposta comercial que discrimine claramente o preço líquido da mercadoria e os tributos CBS e IBS por fora,
Para que o cliente PJ visualize com clareza o valor do crédito tributário que poderá recuperar e a proposta esteja em conformidade com a legislação de transparência fiscal [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Template Obrigatório do PDF): A proposta impressa deve incluir uma seção intitulada "Composição do Valor — Regime IVA Dual (Reforma Tributária)" contendo a tabela de decomposição de preço base e tributos, seguindo a mesma estrutura do checkout digital.
* RN02 (Identificação do Destino Fiscal): O PDF deve conter uma linha informativa: "Alíquotas aplicadas conforme local de consumo: [Município]/[UF] — Código IBGE: [código]".
* RN03 (Rodapé Fiscal Obrigatório): Toda proposta impressa deve conter no rodapé a observação: "Os valores de CBS e IBS nesta proposta são estimativas baseadas nas alíquotas vigentes na data de emissão. O valor final será consolidado no faturamento comToken de Garantia Fiscal de 48 horas."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Vendedor gera PDF de proposta para cliente PJ em Minas Gerais

* Dado que o vendedor concluiu a simulação de uma proposta de R$ 25.000,00 (preço líquido) para um cliente PJ em Uberlândia (MG);
* Quando o vendedor clicar em "Gerar Proposta em PDF";
* Então o PDF gerado deve conter a seção "Composição do Valor — Regime IVA Dual" com a decomposição completa;
* E exibir a linha de identificação: "Alíquotas aplicadas conforme local de consumo: Uberlândia/MG — Código IBGE: 3170206";
* E incluir o rodapé fiscal padrão sobre o Token de Garantia Fiscal de 48 horas [INDEX].

#### Cenário 2: Proposta para cliente com múltiplos itens e diferentes NCMs

* Dado que a proposta contém 3 itens com NCMs diferentes, sendo um deles sujeito ao IS;
* Quando o PDF for gerado;
* Então a decomposição deve ser apresentada por item individual e totalizada ao final;
* E o IS deve ser destacado apenas nos itens aos quais se aplica;
* E o total geral deve consolidar todos os tributos [INDEX].

------------------------------
## 📝 US-03: Conformidade Visual com a Legislação de Transparência Fiscal

### 1. Descrição da História (Visão de Negócio)

Como Diretor Jurídico e de Compliance,
Quero que todas as interfaces de cliente (web, mobile e impressos) estejam em conformidade com os requisitos legais de exibição da carga tributária da Reforma Tributária,
Para mitigar riscos de autuações por omissão de informação fiscal e reforçar o posicionamento da empresa como referência em transparência [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Exibição Obrigatória em Todas as Telas de Fechamento): Qualquer tela ou página que apresente o valor total a pagar pelo cliente deve conter a decomposição CBS/IBS, sem exceção. Isso inclui: carrinho, checkout, página de confirmação de pedido, e-mail de confirmação, proposta em PDF e portal de acompanhamento de pedidos.
* RN02 (Percentual de Carga Tributária Total): Além da decomposição por tributo, deve ser exibido o percentual total da carga tributária sobre o valor líquido: "Carga Tributária Total: XX,X% sobre o valor do produto/serviço".
* RN03 (Auditoria Periódica de Conformidade Visual): O time de Compliance deve realizar uma auditoria visual trimestral em todos os canais para verificar se as interfaces mantêm a formatação de decomposição tributária conforme a legislação vigente.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Todas as telas de fechamento exibem decomposição fiscal

* Dado que um cliente navega pelo e-commerce, adiciona itens ao carrinho e avança para o checkout;
* Quando o cliente passar pelas telas de Carrinho → Checkout → Confirmação;
* Então todas essas telas devem exibir a decomposição de CBS e IBS no resumo financeiro;
* E o e-mail de confirmação enviado após o pagamento deve conter a mesma decomposição;
* E a carga tributária total (percentual) deve estar visível em cada uma dessas interfaces [INDEX].

#### Cenário 2: Auditoria de conformidade detecta canal fora do padrão

* Dado que o time de Compliance executa a auditoria visual trimestral;
* Quando for identificado que o Portal B2B não está exibindo a linha do IBS separadamente (apenas CBS);
* Então o time deve registrar uma não conformidade no sistema de governança;
* E acionar o Product Manager do Portal B2B com SLA de correção de 48 horas;
* E após a correção, gerar evidência visual (screenshot) anexada ao registro de conformidade [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
