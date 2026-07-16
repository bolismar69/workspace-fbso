# Detalhamento de Features do Programa — Onda 1 e Onda 2
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Status: Pronto para Detalhamento de User Stories
- Responsáveis: Product Managers (PMs) e Product Owners (POs) Corporativos
- Referência: [03-EPICS.md](./03-EPICS.md)

------------------------------
## 1. Objetivo do Documento

Este documento detalha as Features (capacidades funcionais de produto) derivadas dos Épicos definidos em [03-EPICS.md](./03-EPICS.md). Cada Feature está vinculada a um Épico e descreve o objetivo de negócio, o comportamento funcional esperado e as regras de negócio aplicáveis.

A numeração adota o padrão **Onda.Épico.Feature** (`01.01.1` = Onda 1, Épico 1, Feature 1), garantindo rastreabilidade completa da diretriz estratégica até a User Story.

**Features por Onda:** 7 (Onda 1) + 9 (Onda 2) = **16 Features** no total.

------------------------------
## 2. Onda 1 — Canais Comerciais e Vendas

- **Fase:** Sistemas de Vendas, CRM e Plataformas Comerciais
- **Épicos:** 01.01 (Qualificação e Onboarding CRM), 01.02 (Conexão à Inteligência Corporativa), 01.03 (Precificação Dinâmica e Transparência)
- **Responsáveis:** Product Managers (PMs) de Vendas e POs de Canais Comerciais

### 🔍 ÉPICO 01.01: Qualificação Geográfica, Saneamento e Onboarding de Clientes (CRM)

#### Feature 01.01.1: Validação Cadastral Geográfica em Tempo Real

* Objetivo de Negócio: Garantir a exatidão do local de consumo para correta aplicação da alíquota do IBS baseada no princípio do destino [INDEX].
* Descrição Funcional: O sistema deve validar de forma mandatória os dados geográficos de novos clientes ou alterações cadastrais na entrada (onboarding). O preenchimento do CEP deve acionar uma validação que preenche automaticamente o Estado, o Município e o Código IBGE oficial, impedindo digitação manual desses campos.
* Regras de Negócio e Restrições:
   * Se o endereço for de uma zona rural ou sem CEP específico, a interface deve exigir o preenchimento manual guiado por lista fechada de municípios homologados pelo IBGE, bloqueando caracteres livres.
   * O cadastro só será marcado como "Apto para Venda" se possuir o ID geográfico do IBGE vinculado.

#### Feature 01.01.2: Governança e Trava Comercial de Vendas (CRM)

* Objetivo de Negócio: Blindar a empresa contra erros de precificação causados por cadastros desatualizados de clientes antigos.
* Descrição Funcional: Mecanismo de bloqueio operacional nas telas de criação de oportunidades, propostas e orçamentos dentro do CRM. O sistema deve varrer o cadastro do cliente selecionado; se os dados de localização (Código IBGE) não estiverem higienizados conforme a Reforma Tributária, o botão de "Gerar Proposta" deve ficar desabilitado.
* Regras de Negócio e Restrições:
   * O sistema deve exibir um aviso claro na tela: "Proposta bloqueada: Cadastro do cliente necessita de atualização geográfica para enquadramento no IVA Dual".
   * Deve existir um fluxo de aprovação de exceção para que o Gerente Comercial libere a trava apenas em casos de contingência jurídica comprovada.

------------------------------

### ⚙️ ÉPICO 01.02: Conexão Comercial à Inteligência Corporativa de Cálculo

#### Feature 01.02.1: Simulador Unificado em Telas de Proposta (Omnicanalidade)

* Objetivo de Negócio: Garantir que o vendedor e o cliente vejam o mesmo impacto tributário em qualquer canal [INDEX].
* Descrição Funcional: Disponibilização de um botão "Simular Impostos" na montagem do carrinho de compras (E-commerce) e na tela de propostas (CRM). Ao ser acionado, o sistema envia o cenário comercial para a inteligência de cálculo e retorna o valor projetado de CBS, IBS e IS sem criar um registro fiscal definitivo.
* Regras de Negócio e Restrições:
   * O tempo de resposta para o cliente na tela do e-commerce não pode prejudicar a experiência de compra (alvo de mercado de alta performance).
   * A simulação deve considerar o perfil do cliente (Contribuinte PJ no Lucro Real vs. Consumidor Final PF), pois isso altera o comportamento do cálculo de negócios.

#### Feature 01.02.2: Resiliência de Vendas (Contingência Local)

* Objetivo de Negócio: Evitar a perda de vendas na ponta caso a inteligência central de cálculo fique temporariamente indisponível.
* Descrição Funcional: Mecanismo de fallback de negócios para os canais comerciais. Se a consulta ao motor de impostos falhar por timeout ou indisponibilidade, a plataforma de vendas deve ativar uma política de contingência comercial.
* Regras de Negócio e Restrições:
   * O sistema adota uma tabela estática de alíquotas médias contingenciais gravada localmente no canal para permitir que o cliente conclua o pedido.
   * O pedido entra com o status "Aguardando Auditoria Fiscal" e o cliente é notificado discretamente na tela de que as condições fiscais finais serão consolidadas no faturamento.

------------------------------

### 💰 ÉPICO 01.03: Precificação Dinâmica, Margem Líquida e Transparência ("Por Fora")

#### Feature 01.03.1: Interface Visual de Checkout com Decomposição do IVA

* Objetivo de Negócio: Atender às leis de transparência fiscal e educar o mercado sobre a precificação "por fora" [INDEX].
* Descrição Funcional: Redesenho completo do resumo financeiro do carrinho de compras e da proposta impressa. O valor principal do produto/serviço deve passar a representar o Preço Líquido (sem imposto por dentro). Abaixo dele, devem ser criadas linhas explícitas destacando o CBS (Federal), o IBS (Destino) e o valor totalizado a pagar.
* Regras de Negócio e Restrições:
   * Exemplo visual obrigatório para o layout:
     * Valor do Item: R$ 1.000,00
     * CBS (Federal - X%): R$ 88,00
     * IBS (Destino - Y%): R$ 177,00
     * Total do Pedido: R$ 1.265,00

#### Feature 01.03.2: Painel de Atratividade B2B (Calculadora de Crédito do IVA)

* Objetivo de Negócio: Reduzir o atrito comercial do aumento nominal de preço para clientes corporativos (PJ).
* Descrição Funcional: Inclusão de um painel informativo exclusivo para vendas corporativas (B2B) no CRM e no portal de clientes. Ao simular a venda para uma empresa que opera no Lucro Real, o sistema deve calcular e exibir o valor do crédito que aquele cliente irá recuperar na entrada dele.
* Regras de Negócio e Restrições:
   * O painel deve exibir uma mensagem comercial de impacto: "Preço Nominal: R$ 1.265,00. Crédito Recuperável estimado para sua empresa: R$ 265,00. Custo Efetivo Real: R$ 1.000,00".

#### Feature 01.03.3: Chave de Garantia Comercial (Token de Validade Fiscal)

* Objetivo de Negócio: Impedir prejuízos causados por alterações de alíquotas entre a negociação comercial e a emissão da nota [INDEX].
* Descrição Funcional: Mecanismo que gera um identificador único (Token) para cada proposta comercial emitida e aceita. Esse token congela as premissas de preço base e as alíquotas simuladas de IBS/CBS por uma janela regulamentada de tempo (ex: 48 horas).
* Regras de Negócio e Restrições:
   * Se o faturamento (Onda 2) receber o pedido dentro do prazo de validade do token, ele deve respeitar o valor acordado, absorvendo eventuais variações diárias de impostos na margem da empresa. Se o prazo expirar, o pedido exige uma nova simulação obrigatória.

------------------------------
## 3. Onda 2 — Finanças, Faturamento e ERP

- **Fase:** Sistemas de Finanças, Escrituração, Faturamento e Governança SAP
- **Épicos:** 02.01 (Faturamento Integrado e Consistência SAP), 02.02 (Split Payment Bancário), 02.03 (Apropriação de Créditos no Lucro Real)
- **Responsáveis:** Product Managers (PMs) de Finanças/Controladoria e POs de ERP/SAP

### 🧾 ÉPICO 02.01: Faturamento Integrado e Consistência de Emissão (SAP)

#### Feature 02.01.1: Validação de Faturamento Pré-Emissão e Trava Contábil

* Objetivo de Negócio: Eliminar o risco de rejeição de Notas Fiscais na SEFAZ por divergência de valores e garantir a conformidade do faturamento no Lucro Real.
* Descrição Funcional: Mecanismo automatizado de auditoria interna executado no módulo SAP SD (Vendas e Distribuição) antes da transmissão do documento fiscal (NF-e ou NFS-e nacional). O sistema deve confrontar as alíquotas e o cálculo do IVA Dual enviados pela frente de vendas com o cálculo atualizado da inteligência corporativa.
* Regras de Negócio e Restrições:
   * Se houver qualquer divergência matemática de centavos (decorrente de regras de arredondamento) ou alteração de alíquota na cidade de destino, o SAP deve bloquear a emissão da nota automaticamente.
   * O documento entra em status "Bloqueio Fiscal" e uma tarefa é gerada para a Controladoria avaliar e autorizar o ajuste da cotação ou a absorção da diferença na margem da empresa.

#### Feature 02.01.2: Motor de Conversão do ISS para IBS (Serviços Santana de Parnaíba)

* Objetivo de Negócio: Adaptar a operação da matriz para o fim do imposto municipal tradicional (ISS) e início do IBS de serviços baseado no destino.
* Descrição Funcional: Parametrização das regras de determinação de impostos no SAP para faturamento de serviços emitidos pela matriz em Santana de Parnaíba (SP). O sistema deve desativar a retenção de ISS com base na competência da cidade de origem e ativar a aplicação das novas alíquotas de IBS do local de consumo do cliente.
* Regras de Negócio e Restrições:
   * O sistema deve descontinuar a emissão de notas via sistema legado da prefeitura de Santana de Parnaíba e integrar o fluxo diretamente ao padrão da NFS-e Nacional.

#### Feature 02.01.3: Automação de Benefícios e Regimes Especiais (Transição Santana de Parnaíba)

* Objetivo de Negócio: Blindar e aplicar os incentivos fiscais e regimes especiais vigentes da matriz sem intervenção manual do faturista.
* Descrição Funcional: Inteligência de negócio mapeada nas tabelas de condições de impostos do SAP. Ao identificar que a operação é elegível a um regime especial ativo da matriz, o sistema aplica reduções de base de cálculo ou alíquotas diferenciadas sobre o IBS, conforme as regras de transição permitidas pela legislação.
* Regras de Negócio e Restrições:
   * O sistema deve validar se o benefício fiscal local permanece aplicável em vendas interestaduais ou se a regra do destino (alíquota cheia do estado/município comprador) anula o incentivo de origem.

------------------------------

### 💸 ÉPICO 02.02: Operação e Governança do Mecanismo de Split Payment (Tesouraria)

#### Feature 02.02.1: Liquidação e Conciliação Financeira Segregada (Split Bancário)

* Objetivo de Negócio: Adaptar as contas a receber para a retenção automática do imposto na fonte pelas instituições financeiras.
* Descrição Funcional: Reformulação do fluxo de caixa e do processamento de arquivos de retorno bancário (CNAB) no módulo SAP FI (Finanças). O sistema deve reconhecer que os pagamentos efetuados pelos clientes via boleto, PIX ou cartão de crédito entrarão na conta da empresa já descontados os valores de CBS e IBS [INDEX].
* Regras de Negócio e Restrições:
   * Ao processar a baixa do título de um cliente, o SAP deve realizar um lançamento duplo automático: dar baixa total na duplicata do cliente, direcionar o valor líquido para a conta corrente principal e lançar a parcela do imposto retido em uma conta de compensação fiscal de "Impostos Retidos na Fonte - Split".

#### Feature 02.02.2: Ajuste de Split para Operações Incentivadas

* Objetivo de Negócio: Evitar que os bancos retenham impostos a maior na fonte em notas emitidas com benefícios fiscais pela matriz.
* Descrição Funcional: Integração de dados entre o faturamento e as instruções de cobrança bancária. O sistema deve gerar os metadados do boleto ou PIX contendo as tags exatas do valor recalculado do CBS/IBS (após o benefício fiscal aplicado pela Feature 02.01.3).
* Regras de Negócio e Restrições:
   * Garante que a rede bancária processe o split apenas sobre o valor tributário reduzido/incentivado, protegendo o caixa líquido e o capital de giro da empresa na matriz.

#### Feature 02.02.3: Painel de Auditoria e Reconciliação do Split Payment

* Objetivo de Negócio: Identificar falhas de retenção ou desvios cometidos pelas instituições financeiras.
* Descrição Funcional: Relatório de conferência e auditoria interna na tesouraria. O sistema deve cruzar o valor de CBS/IBS destacado na Nota Fiscal (visão faturamento) com o valor efetivamente retido e informado pelo banco no arquivo de liquidação (visão tesouraria).
* Regras de Negócio e Restrições:
   * O painel deve apontar discrepâncias em um relatório de exceções (ex: "Nota Fiscal 123 - Banco reteve R$ 10,00 a maior"). Qualquer erro gera uma notificação automatizada para o time de contas a receber acionar a instituição financeira parceira.

------------------------------

### ⚖️ ÉPICO 02.03: Otimização de Custos e Apropriação de Créditos (Suprimentos e Controladoria)

#### Feature 02.03.1: Auditoria Fiscal de Entrada e Bloqueio de Créditos (Procure-to-Pay)

* Objetivo de Negócio: Garantir que a empresa só se aproprie de créditos legítimos no Lucro Real, evitando riscos de autuação por fornecedores irregulares.
* Descrição Funcional: Mecanismo de checagem automatizada no recebimento físico e fiscal de mercadorias no módulo SAP MM (Compras). O sistema deve condicionar a liberação do crédito de CBS/IBS à validação de que o imposto destacado pelo fornecedor foi devidamente processado e recolhido pelo mecanismo de Split Payment na origem.
* Regras de Negócio e Restrições:
   * Se o fornecedor for optante por regimes simplificados ou houver inconformidade na guia de recolhimento, o SAP impede o lançamento do imposto na conta de "Impostos a Recuperar", direcionando o valor temporariamente para uma conta de "Créditos em Análise".

#### Feature 02.03.2: Segregação Contábil de Custos de Estoque e Ativos

* Objetivo de Negócio: Refletir o custo real dos produtos estocados nos armazéns, expurgando o impacto dos impostos recuperáveis.
* Descrição Funcional: Ajuste na inteligência contábil de avaliação de inventário no SAP. O sistema deve extrair o valor total do CBS e IBS calculados por fora e direcioná-los integralmente para as contas de ativos fiscais, registrando no custo de estoque do produto puramente o seu valor líquido.
* Regras de Negócio e Restrições:
   * Esta segregação é vital para manter a acurácia do Custo das Mercadorias Vendidas (CMV) e evitar a inflação artificial dos estoques da empresa em nível nacional.

#### Feature 02.03.3: Escrituração de Reserva de Incentivos (Subvenção no Lucro Real)

* Objetivo de Negócio: Proteger o ganho financeiro dos regimes especiais de Santana de Parnaíba contra a tributação do IRPJ e da CSLL.
* Descrição Funcional: Rotina de escrituração contábil acionada no fechamento fiscal mensal. O sistema deve identificar o montante total de IBS que deixou de ser recolhido em virtude dos benefícios da matriz e realizar uma transferência contábil interna.
* Regras de Negócio e Restrições:
   * O valor economizado deve ser retirado da conta de resultado de impostos e creditado em uma conta de patrimônio líquido dedicada: "Reserva de Incentivos Fiscais - Subvenção para Investimento", cumprindo os requisitos legais de isenção de imposto de renda sobre o benefício.

### Critérios de Sucesso para Homologação da Onda 2 (Definition of Done - DoD)

O comitê financeiro considerará as Features da Onda 2 prontas para produção quando:

1. Consistência Contábil: O balancete de testes do SAP demonstrar que os impostos gerados pelo faturamento interestadual foram destinados às contas corretas por estado/município de destino sem erros de arredondamento.
2. Baixa Automatizada: O sistema realizar a baixa de 100% dos títulos de teste no banco através do fluxo de Split Payment, liquidando as contas de compensação sem intervenção manual.
3. Blindagem de Santana de Parnaíba: A auditoria contábil validar que os lançamentos dos benefícios fiscais da matriz foram devidamente isolados nas contas de reserva de subvenção, protegendo o Lucro Real.

------------------------------
## 4. Matriz de Rastreabilidade das Features

| Feature | Épico | Onda | User Stories |
|:---|:---|:---|:---|
| 01.01.1 — Validação Cadastral Geográfica em Tempo Real | 01.01 | Onda 1 | [US](./05-USER-STORYS-01-01-1-VALIDACAO-CADASTRAL-GEOGRAFICA-TEMPO-REAL.md) ✅ |
| 01.01.2 — Governança e Trava Comercial de Vendas (CRM) | 01.01 | Onda 1 | [US](./05-USER-STORYS-01-01-2-TRAVA-COMERCIAL-NO-CRM-POR-FALTA-DE-HIGIENIZACAO-CADASTRAL.md) ✅ |
| 01.02.1 — Simulador Unificado Omnicanal | 01.02 | Onda 1 | [US](./05-USER-STORYS-01-02-1-SIMULADOR-UNIFICADO-OMNICANAL.md) ✅ |
| 01.02.2 — Resiliência de Vendas / Contingência Local | 01.02 | Onda 1 | [US](./05-USER-STORYS-01-02-2-RESILIENCIA-VENDAS-CONTINGENCIA-LOCAL.md) ✅ |
| 01.03.1 — Interface Visual de Checkout com Decomposição do IVA | 01.03 | Onda 1 | [US](./05-USER-STORYS-01-03-1-INTERFACE-VISUAL-CHECKOUT-DECOMPOSICAO-IVA.md) ✅ |
| 01.03.2 — Painel de Atratividade B2B / Calculadora de Crédito | 01.03 | Onda 1 | [US](./05-USER-STORYS-01-03-2-PAINEL-ATRATIVIDADE-B2B-CALCULADORA-CREDITO.md) ✅ |
| 01.03.3 — Chave de Garantia / Token de Validade Fiscal | 01.03 | Onda 1 | [US](./05-USER-STORYS-01-03-3-CHAVE-DE-GARANTIA-TOKEN-DE-VALIDADE-FISCAL.md) ✅ |
| 02.01.1 — Validação de Faturamento Pré-Emissão e Trava Contábil | 02.01 | Onda 2 | [US](./05-USER-STORYS-02-01-1-VALIDACAO-FATURAMENTO-PRE-EMISSAO-TRAVA-CONTABIL.md) ✅ |
| 02.01.2 — Motor de Conversão do ISS para IBS | 02.01 | Onda 2 | [US](./05-USER-STORYS-02-01-2-MOTOR-CONVERSAO-ISS-IBS.md) ✅ |
| 02.01.3 — Automação de Benefícios e Regimes Especiais | 02.01 | Onda 2 | [US](./05-USER-STORYS-02-01-3-AUTOMACAO-BENEFICIOS-REGIMES-ESPECIAIS.md) ✅ |
| 02.02.1 — Liquidação e Conciliação Segregada (Split) | 02.02 | Onda 2 | [US](./05-USER-STORYS-02-02-1-LIQUIDACAO-CONCILIACAO-SEGREGADA-SPLIT.md) ✅ |
| 02.02.2 — Ajuste de Split para Operações Incentivadas | 02.02 | Onda 2 | [US](./05-USER-STORYS-02-02-2-AJUSTE-SPLIT-OPERACOES-INCENTIVADAS.md) ✅ |
| 02.02.3 — Painel de Auditoria e Reconciliação do Split | 02.02 | Onda 2 | [US](./05-USER-STORYS-02-02-3-PAINEL-AUDITORIA-RECONCILIACAO-SPLIT.md) ✅ |
| 02.03.1 — Auditoria Fiscal de Entrada e Bloqueio de Créditos | 02.03 | Onda 2 | [US](./05-USER-STORYS-02-03-1-AUDITORIA-FISCAL-ENTRADA-BLOQUEIO-CREDITOS.md) ✅ |
| 02.03.2 — Segregação Contábil de Custos de Estoque e Ativos | 02.03 | Onda 2 | [US](./05-USER-STORYS-02-03-2-SEGREGACAO-CONTABIL-CUSTOS-ESTOQUE-ATIVOS.md) ✅ |
| 02.03.3 — Escrituração de Reserva de Incentivos (Subvenção) | 02.03 | Onda 2 | [US](./05-USER-STORYS-02-03-3-ESCRITURACAO-RESERVA-INCENTIVOS.md) ✅ |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 013-agile-feature, agile-ba-practices.*
