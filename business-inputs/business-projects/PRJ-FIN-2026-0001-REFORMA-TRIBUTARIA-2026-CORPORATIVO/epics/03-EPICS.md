# Lista de Épicos do Programa — Onda 1 e Onda 2
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Status: Pronto para Refinamento Funcional
- Responsáveis: Product Managers (PMs) e Product Owners (POs) Corporativos
- Referência: [02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)

------------------------------
## 1. Objetivo do Documento

Este documento define os Épicos (Grandes Blocos de Entrega) sob a perspectiva de negócios, necessários para capacitar tanto o ecossistema de vendas (E-commerce, CRM, Portais B2B e Venda Direta) quanto as áreas de Faturamento, Controladoria, Tesouraria e Suprimentos no ERP SAP a operar em total conformidade com o IVA Dual (CBS, IBS e IS), atendendo aos requisitos macro estabelecidos no [REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md).

A numeração adota o padrão **Onda.Épico** (`01.01` = Onda 1, Épico 1), garantindo rastreabilidade consistente com a hierarquia de Features e User Stories.

------------------------------
## 2. Visão Geral da Jornada do Programa

```text
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                               ONDA 1: CANAIS COMERCIAIS E VENDAS                          │
│  [ CRM / CADASTRO ] ──────► [ MENSAGERIA / INTEGRAÇÃO ] ──────► [ PRECIFICADOR / CHECKOUT ]│
│     Épico 01.01:                  Épico 01.02:                        Épico 01.03:         │
│  Qualificação e                Conexão com a Inteligência          Precificação Dinâmica   │
│ Saneamento de Clientes              Corporativa                   e Transparência (IVA)    │
└─────────────────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                             ONDA 2: FINANÇAS, FATURAMENTO E ERP                           │
│ [ EMISSÃO E CONFIRMAÇÃO ] ──────► [ GESTÃO DE CAIXA / RETENÇÃO ] ──────► [ ESCRITURAÇÃO ] │
│     Épico 02.01:                       Épico 02.02:                        Épico 02.03:    │
│  Faturamento e Consistência         Mecanismo de Split                  Apropriação de    │
│       Matemática (SAP)              Payment Bancário                 Créditos no Lucro Real│
└─────────────────────────────────────────────────────────────────────────────────────────┘
```

------------------------------
## 3. Onda 1 — Canais Comerciais e Vendas

> **Fase:** Sistemas de Vendas, CRM e Plataformas Comerciais
> **Responsáveis:** Product Managers (PMs) de Vendas e POs de Canais Comerciais

### 🔍 ÉPICO 01.01: Qualificação Geográfica, Saneamento e Onboarding de Clientes (CRM)

* Descrição de Alto Nível: Adaptar todos os fluxos de captação de clientes, criação de leads e onboarding de contas (B2B e B2C) para coletar, validar e auditar o exato local de destino do consumo (município e estado federativo) antes da emissão de qualquer proposta comercial.
* Justificativa de Negócio: Sob as regras da Reforma Tributária, o IBS é integralmente baseado no princípio do destino. Cadastros com endereços ambíguos ou códigos municipais desatualizados distorcem a margem projetada, expondo a empresa a erros graves de precificação interestadual.
* Requisitos Vinculados ([REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)): BR-03 (Qualificação Geográfica de Cadastro).
* Capacidades Esperadas (Features do Produto):
   * Validação Cadastral em Tempo Real: Motores de busca geográfica que cruzam dados cadastrais informados pelos vendedores com as bases oficiais do IBGE e Correios.
   * Governança de Alíquotas de Origem/Destino: Trava de segurança comercial no CRM impedindo que o time comercial gere propostas para clientes que não possuam o campo de "local de consumo" 100% qualificado.

### ⚙️ ÉPICO 01.02: Conexão Comercial à Inteligência Corporativa de Cálculo

* Descrição de Alto Nível: Integrar os canais front-end (interfaces de e-commerce, portais de autoatendimento e plataformas de CRM dos vendedores) à regra de cálculo corporativa e unificada de impostos, de modo que toda simulação de vendas consuma a mesma fonte de dados fiscal.
* Justificativa de Negócio: Garantir o alinhamento da "Omnicanalidade Tributária". Um cliente cotando via portal B2B ou diretamente com um vendedor pelo CRM precisa visualizar exatamente o mesmo preço base e a mesma projeção do IVA Dual.
* Requisitos Vinculados ([REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)): BR-01 (Centralização da Inteligência) e BR-02 (Autonomia do Time Fiscal).
* Capacidades Esperadas (Features do Produto):
   * Simulador Unificado em Vendas: Endpoints de simulação imediata de tributos integrados nativamente nas telas onde o vendedor ou cliente monta a proposta comercial.
   * Resiliência de Caixa (Contingência de Canais): Mecanismos de negócios para garantir que, caso haja lentidão na consulta da regra fiscal, as plataformas comerciais operem sob parâmetros de segurança locais sem travar a jornada de compra do cliente.

### 💰 ÉPICO 01.03: Precificação Dinâmica, Margem Líquida e Transparência ("Por Fora")

* Descrição de Alto Nível: Reestruturar os motores de ofertas, carrinhos de compras e exibições de propostas comerciais para suportar a mecânica de cálculo "por fora" do IVA Dual, demonstrando a decomposição exata do preço base da mercadoria/serviço somado aos novos tributos (CBS, IBS e IS se aplicável).
* Justificativa de Negócio: Proteger a rentabilidade no regime de Lucro Real frente às variações regionais do IBS e cumprir as novas exigências de transparência ao consumidor final na operação nacional.
* Requisitos Vinculados ([REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)): BR-04 (Transparência), BR-05 (Proteção de Margem) e BR-06 (Garantia de Preço/Token).
* Capacidades Esperadas (Features do Produto):
   * Apresentação Transparente do IVA: Telas de checkout e propostas comerciais redesenhadas para discriminar visualmente o Preço Líquido (Base), a parcela do CBS (federal) e o IBS (do município de destino).
   * Painel de Atratividade B2B (Crédito do IVA): Simulador comercial que demonstra para o comprador PJ (Pessoa Jurídica) o valor exato do crédito tributário que ele poderá se apropriar na cadeia não cumulativa do Lucro Real, mitigando o atrito do aumento nominal de preço.
   * Garantia Comercial de Alíquotas: Geração de uma chave de conformidade que congela o preço base e a alíquota municipal calculada por um intervalo regulamentado de horas, impedindo que flutuações de fechamento afetem o combinado com o cliente.

### Critérios de Sucesso para Fechamento da Onda 1 (Comercial)

A alta gestão e as lideranças comerciais considerarão esta lista de Épicos concluída quando:

1. Acurácia Cadastral: 100% dos novos leads e contas ativas no CRM possuírem o código IBGE do destino preenchido e validado.
2. Transparência de Margem: O time de vendas (B2B/CRM) conseguir extrair o relatório de lucratividade líquida descontando o IBS de cada estado do país antes de assinar novos contratos.
3. Conversão Comercial Protegida: A introdução do cálculo "por fora" não gerar aumento na taxa de abandono de carrinhos no e-commerce devido à lentidão ou falta de clareza na exibição do preço final.

------------------------------
## 4. Onda 2 — Finanças, Faturamento e ERP

> **Fase:** Sistemas de Finanças, Escrituração, Faturamento e Governança SAP
> **Responsáveis:** Product Managers (PMs) de Finanças/Controladoria e POs de ERP/SAP

### 🧾 ÉPICO 02.01: Faturamento Integrado e Consistência de Emissão (SAP)

* Descrição de Alto Nível: Garantir que o processo de faturamento físico e de serviços (matriz em Santana de Parnaíba) envie dados de cálculo idênticos aos simulados em vendas para a emissão dos novos layouts de documentos fiscais, eliminando rejeições regulatórias e garantindo a aplicação de incentivos vigentes.
* Justificativa de Negócio: No regime de Lucro Real, divergências de centavos ou inconsistências nas regras de transição federativas entre a proposta comercial e a nota emitida travam a expedição de mercadorias e geram risco de autuação fiscal de faturamento.
* Requisitos Vinculados ([REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)): BR-07 (Unicidade Matemática entre Pedido e Nota Fiscal).
* Capacidades Esperadas (Features do Produto):
   * Validação de Faturamento Pré-Emissão: Mecanismo de negócio que bloqueia a geração da nota fiscal se o valor calculado na liquidação divergir das premissas de preço e alíquotas de destino fechadas em vendas.
   * Conversão do ISS em IBS (Matriz): Regra de faturamento dedicada para a matriz em Santana de Parnaíba, descontinuando o modelo de retenção de ISS na competência municipal e ativando a cobrança do IBS de serviços por destino.
   * Governança de Benefícios Fiscais e Regimes Especiais (Santana de Parnaíba): Inteligência de negócio que identifica transações elegíveis a incentivos fiscais locais (como reduções de base de cálculo, alíquotas diferenciadas ou diferimentos). Garante a aplicação dessas exceções nas operações originadas na matriz, aplicando as regras de transição estipuladas pela reforma tributária de forma automatizada.

### 💸 ÉPICO 02.02: Operação e Governança do Mecanismo de Split Payment (Tesouraria)

* Descrição de Alto Nível: Adaptar o processo de contas a receber e tesouraria para suportar a segregação automatizada do fluxo de caixa no ato do pagamento do cliente, destinando os impostos diretamente aos entes públicos e o valor líquido para a empresa, considerando reduções por incentivos.
* Justificativa de Negócio: Esta é a mudança operacional mais crítica da reforma. A empresa precisa gerenciar o fluxo de caixa sabendo que o dinheiro do imposto (CBS e IBS) será retido na fonte pela rede bancária, alterando radicalmente as regras de conciliação e concorrência de saldos.
* Requisitos Vinculados ([REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)): BR-09 (Viabilização do Mecanismo de Split Payment).
* Capacidades Esperadas (Features do Produto):
   * Conciliação Financeira Segregada (Split): Processo de liquidação capaz de baixar títulos de cobrança (boletos, cartões, PIX) identificando a parcela líquida recebida e a parcela de imposto retida pelo banco.
   * Split Reduzido para Regimes Especiais: Parametrização financeira que calcula a instrução de retenção bancária com base no valor incentivado real do documento fiscal emitido por Santana de Parnaíba, mitigando o risco de retenção a maior na fonte e protegendo o caixa líquido.
   * Auditoria Bancária do IVA: Relatório de controle financeiro para cruzar o valor retido na fonte pela instituição bancária com o valor devido de CBS/IBS calculado pelo sistema.

### ⚖️ ÉPICO 02.03: Otimização de Custos e Apropriação de Créditos (Suprimentos e Controladoria)

* Descrição de Alto Nível: Estruturar as regras de entrada de mercadorias e serviços no SAP para mapear, auditar e apropriar 100% dos créditos não cumulativos do IVA Dual gerados pela cadeia de fornecedores da empresa nacionalmente.
* Justificativa de Negócio: No Lucro Real, a eficiência de margem depende da captura integral dos créditos. Compras efetuadas de fornecedores que não repassarem o imposto (devido ao Split Payment ou regimes especiais) não geram crédito, o que encarece o custo do produto e penaliza a rentabilidade.
* Requisitos Vinculados ([REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)): BR-08 (Rastreabilidade de Créditos no Lucro Real).
* Capacidades Esperadas (Features do Produto):
   * Auditoria de Crédito na Entrada (Procure-to-Pay): Regra de negócio na escrituração fiscal que condiciona o aproveitamento do crédito de CBS/IBS à validação da regularidade e recolhimento do imposto por parte do fornecedor.
   * Segregação de Custo e Tributo Recuperável: Atualização da inteligência contábil de estoques para separar o valor do imposto recuperável do custo real do produto estocado em todos os armazéns do país.
   * Contabilização de Reserva de Incentivos (Lucro Real): Mecanismo de escrituração contábil que isola e direciona o ganho financeiro obtido pelos regimes especiais de Santana de Parnaíba para contas específicas de subvenção para investimento, mantendo a blindagem de exclusão da base de cálculo do IRPJ e da CSLL.

### Critérios de Sucesso para Fechamento da Onda 2 (Financeira)

As diretorias de Controladoria e Finanças considerarão esta lista de Épicos concluída e o programa pronto para encerramento quando:

1. Divergência Zero: 100% dos livros fiscais gerados no SAP baterem centavo por centavo com os documentos fiscais emitidos nacionalmente.
2. Fluxo de Caixa Conciliado: O processo de conciliação bancária de recebíveis rodar de forma automatizada sob o modelo de Split Payment sem gerar pendências inexplicáveis de saldo, inclusive nas notas fiscais faturadas com alíquotas incentivadas.
3. Margem de Lucro Real Protegida: A empresa comprovar, no primeiro fechamento fiscal, que capturou todos os créditos tributários previstos em lei sobre as compras efetuadas e registrou os incentivos da matriz conforme as exigências regulatórias vigentes.

------------------------------
## 5. Matriz de Rastreabilidade dos Épicos

| Épico | Onda | Requisitos Vinculados | Features |
|:---|:---|:---|:---|
| 01.01 — Qualificação Geográfica e Onboarding CRM | Onda 1 | BR-03 | 01.01.1, 01.01.2 |
| 01.02 — Conexão à Inteligência Corporativa | Onda 1 | BR-01, BR-02 | 01.02.1, 01.02.2 |
| 01.03 — Precificação Dinâmica e Transparência | Onda 1 | BR-04, BR-05, BR-06 | 01.03.1, 01.03.2, 01.03.3 |
| 02.01 — Faturamento Integrado e Consistência (SAP) | Onda 2 | BR-07 | 02.01.1, 02.01.2, 02.01.3 |
| 02.02 — Split Payment Bancário (Tesouraria) | Onda 2 | BR-09 | 02.02.1, 02.02.2, 02.02.3 |
| 02.03 — Apropriação de Créditos no Lucro Real | Onda 2 | BR-08 | 02.03.1, 02.03.2, 02.03.3 |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 012-agile-epic, agile-ba-practices.*
