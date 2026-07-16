# Histórias de Usuário (User Stories) — Feature 01.3 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: ERP SAP SD/MM — Condições de Impostos e Benefícios Fiscais (Onda 2)
- Feature Relacionada: 02.01.3 — Automação de Benefícios e Regimes Especiais (Transição Santana de Parnaíba) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Aplicação Automática de Benefícios Fiscais da Matriz sem Intervenção Manual

### 1. Descrição da História (Visão de Negócio)

Como Controller Fiscal responsável pela matriz de Santana de Parnaíba,
Quero que o SAP aplique automaticamente os regimes especiais e benefícios fiscais vigentes na matriz (reduções de base de cálculo, alíquotas diferenciadas, diferimentos) às operações elegíveis,
Para eliminar a dependência de intervenção manual do faturista e reduzir o risco de perda de incentivos por erro operacional [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Motor de Regras Condicionais): A aplicação do benefício deve ser acionada automaticamente quando a operação atender simultaneamente a três condições: (a) ser originada na matriz de Santana de Parnaíba, (b) o NCM/NBS do produto ou serviço estar na lista de itens elegíveis ao benefício, (c) a natureza da operação (CFOP) ser compatível com o regime especial.
* RN02 (Validação de Vigência Temporal): O SAP deve verificar a data de vigência do benefício fiscal na tabela de regimes especiais e só aplicá-lo se a data da operação estiver dentro do período de validade.
* RN03 (Cálculo em Cascata): O benefício deve ser aplicado na ordem correta: primeiro a redução de base de cálculo (se houver), depois a alíquota diferenciada sobre a base reduzida, resultando no valor final de IBS/CBS da operação incentivada.
* RN04 (Rastreabilidade do Benefício no Documento Fiscal): A nota fiscal emitida deve conter no campo de "Informações Complementares" o código do benefício fiscal aplicado e o valor do imposto que deixou de ser recolhido.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Operação elegível a redução de base de cálculo — benefício aplicado automaticamente

* Dado que a matriz de Santana de Parnaíba está emitindo uma NF-e de produto com NCM elegível a redução de 40% na base de cálculo do IBS;
* E o benefício está vigente conforme tabela de regimes especiais;
* E a operação tem CFOP de venda interestadual compatível;
* Quando o faturamento for processado;
* Então o SAP deve identificar automaticamente a elegibilidade;
* E reduzir a base de cálculo do IBS em 40% (de R$ 10.000,00 para R$ 6.000,00);
* E aplicar a alíquota de IBS do destino sobre R$ 6.000,00;
* E registrar no documento fiscal: "Benefício Fiscal Aplicado: Redução de Base de Cálculo — Lei Mun. SNP nº XXXX/2026. IBS economizado: R$ XXX,XX" [INDEX].

#### Cenário 2: Produto não elegível — benefício não aplicado

* Dado que a matriz está faturando um produto cujo NCM não consta na lista de itens elegíveis ao benefício;
* Quando o faturamento for processado;
* Então o SAP não deve aplicar nenhuma redução ou alíquota diferenciada;
* E o IBS deve ser calculado sobre a base cheia do produto;
* E nenhum registro de benefício deve aparecer no documento fiscal [INDEX].

#### Cenário 3: Benefício expirado — sistema rejeita aplicação

* Dado que o benefício fiscal "Redução XYZ" expirou em 31/12/2029;
* E a data da operação de faturamento é 15/01/2030;
* Quando o SAP verificar a tabela de regimes especiais;
* Então o sistema deve detectar que o benefício está fora do período de vigência;
* E não aplicar a redução, calculando o IBS sobre a base cheia;
* E registrar em log: "Benefício XYZ não aplicado — fora do período de vigência (expirou em 31/12/2029)" [INDEX].

------------------------------
## 📝 US-02: Validação de Benefício de Origem em Vendas Interestaduais

### 1. Descrição da História (Visão de Negócio)

Como Controller Tributário da matriz,
Quero que o SAP valide se um benefício fiscal concedido pelo município de Santana de Parnaíba permanece aplicável quando a venda é interestadual,
Para evitar que a empresa aplique indevidamente um incentivo de origem em uma operação cujo IBS é integralmente devido ao destino [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Princípio do Destino vs. Incentivo de Origem): Benefícios fiscais concedidos pelo município de origem (Santana de Parnaíba) podem não ser reconhecidos pelo estado/município de destino da mercadoria ou serviço. O SAP deve verificar a aplicabilidade do benefício no destino antes de aplicá-lo.
* RN02 (Regra de Bloqueio Interestadual): Se a operação for interestadual e o benefício fiscal for de origem (concedido apenas pela legislação de Santana de Parnaíba/SP), o SAP deve alertar que o destino pode não reconhecer o benefício e exigir o recolhimento integral do IBS.
* RN03 (Cenário de Risco — Alerta ao Faturista): Quando houver dúvida sobre a aplicabilidade do benefício no destino, o SAP deve exibir um alerta amarelo (não bloqueante, mas informativo): "⚠️ Atenção: O benefício fiscal aplicado é de origem. Verifique se o estado/município de destino reconhece este incentivo."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Venda interestadual — benefício de origem não reconhecido pelo destino

* Dado que a matriz de Santana de Parnaíba está faturando uma venda interestadual para um cliente em Goiânia (GO);
* E o produto se enquadra em um benefício fiscal municipal de Santana de Parnaíba (redução de ISS/IBS);
* Quando o SAP processar o faturamento;
* Então o sistema deve detectar que se trata de operação interestadual com benefício de origem;
* E exibir o alerta amarelo sobre a possível não aceitação pelo destino;
* E registrar no log: "Operação interestadual — benefício de origem aplicado. Destino: GO. Risco: não reconhecimento pelo fisco de destino" [INDEX].

#### Cenário 2: Venda dentro do estado — benefício aplicado sem alerta

* Dado que a matriz está faturando uma venda para um cliente na cidade de São Paulo (SP) — mesma UF;
* E o produto é elegível ao benefício fiscal municipal de Santana de Parnaíba;
* Quando o SAP processar o faturamento;
* Então o sistema deve aplicar o benefício normalmente;
* E não deve exibir o alerta de risco interestadual [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
