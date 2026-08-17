# Lista de Épicos da Onda 1: Canais Comerciais e Vendas
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Fase: Onda 1 (Sistemas de Vendas, CRM e Plataformas Comerciais)
- Status: Pronto para Refinamento Funcional
- Responsáveis: Product Managers (PMs) de Vendas e POs de Canais Comerciais

------------------------------
## 1. Objetivo do Documento

* Este documento define os Épicos (Grandes Blocos de Entrega) sob a perspectiva de negócios, necessários para capacitar o ecossistema de vendas da companhia (E-commerce, CRM, Portais B2B e Venda Direta) a operar em total conformidade com o IVA Dual (CBS, IBS e IS), atendendo aos requisitos macro estabelecidos no REQUIREMENTS.md.

------------------------------
## 2. Visão Geral da Jornada Comercial (Onda 1)

```
[ CRM / CADASTRO ] ──────────► [ MENSAGERIA / INTEGRAÇÃO ] ──────────► [ PRECIFICADOR / CHECKOUT ]
    Épico 01:                      Épico 02:                             Épico 03:
 Qualificação e                 Conexão com a Inteligência             Precificação Dinâmica
Saneamento de Clientes              Corporativa                        e Transparência (IVA)
```

------------------------------
## 3. Detalhamento dos Épicos de Negócio

### 🔍 ÉPICO 01: Qualificação Geográfica, Saneamento e Onboarding de Clientes (CRM)

* Descrição de Alto Nível: Adaptar todos os fluxos de captação de clientes, criação de leads e onboarding de contas (B2B e B2C) para coletar, validar e auditar o exato local de destino do consumo (município e estado federativo) antes da emissão de qualquer proposta comercial.
* Justificativa de Negócio: Sob as regras da Reforma Tributária, o IBS é integralmente baseado no princípio do destino. Cadastros com endereços ambíguos ou códigos municipais desatualizados distorcem a margem projetada, expondo a empresa a erros graves de precificação interestadual.
* Requisitos Vinculados (REQUIREMENTS.md): BR-03 (Qualificação Geográfica de Cadastro).
* Capacidades Esperadas (Features do Produto):
* Validação Cadastral em Tempo Real: Motores de busca geográfica que cruzam dados cadastrais informados pelos vendedores com as bases oficiais do IBGE e Correios.
   * Governança de Alíquotas de Origem/Destino: Trava de segurança comercial no CRM impedindo que o time comercial gere propostas para clientes que não possuam o campo de "local de consumo" 100% qualificado.

### ⚙️ ÉPICO 02: Conexão Comercial à Inteligência Corporativa de Cálculo

* Descrição de Alto Nível: Integrar os canais front-end (interfaces de e-commerce, portais de autoatendimento e plataformas de CRM dos vendedores) à regra de cálculo corporativa e unificada de impostos, de modo que toda simulação de vendas consuma a mesma fonte de dados fiscal.
* Justificativa de Negócio: Garantir o alinhamento da "Omnicanalidade Tributária". Um cliente cotando via portal B2B ou diretamente com um vendedor pelo CRM precisa visualizar exatamente o mesmo preço base e a mesma projeção do IVA Dual.
* Requisitos Vinculados (REQUIREMENTS.md): BR-01 (Centralização da Inteligência) e BR-02 (Autonomia do Time Fiscal).
* Capacidades Esperadas (Features do Produto):
* Simulador Unificado em Vendas: Endpoints de simulação imediata de tributos integrados nativamente nas telas onde o vendedor ou cliente monta a proposta comercial.
   * Resiliência de Caixa (Contingência de Canais): Mecanismos de negócios para garantir que, caso haja lentidão na consulta da regra fiscal, as plataformas comerciais operem sob parâmetros de segurança locais sem travar a jornada de compra do cliente.

### 💰 ÉPICO 03: Precificação Dinâmica, Margem Líquida e Transparência ("Por Fora")

* Descrição de Alto Nível: Reestruturar os motores de ofertas, carrinhos de compras e exibições de propostas comerciais para suportar a mecânica de cálculo "por fora" do IVA Dual, demonstrando a decomposição exata do preço base da mercadoria/serviço somado aos novos tributos (CBS, IBS e IS se aplicável).
* Justificativa de Negócio: Proteger a rentabilidade no regime de Lucro Real frente às variações regionais do IBS e cumprir as novas exigências de transparência ao consumidor final na operação nacional.
* Requisitos Vinculados (REQUIREMENTS.md): BR-04 (Transparência), BR-05 (Proteção de Margem) e BR-06 (Garantia de Preço/Token).
* Capacidades Esperadas (Features do Produto):
* Apresentação Transparente do IVA: Telas de checkout e propostas comerciais redesenhadas para discriminar visualmente o Preço Líquido (Base), a parcela do CBS (federal) e o IBS (do município de destino).
   * Painel de Atratividade B2B (Crédito do IVA): Simulador comercial que demonstra para o comprador PJ (Pessoa Jurídica) o valor exato do crédito tributário que ele poderá se apropriar na cadeia não cumulativa do Lucro Real, mitigando o atrito do aumento nominal de preço.
   * Garantia Comercial de Alíquotas: Geração de uma chave de conformidade que congela o preço base e a alíquota municipal calculada por um intervalo regulamentado de horas, impedindo que flutuações de fechamento afetem o combinado com o cliente.

------------------------------
## 4. Critérios de Sucesso para Fechamento da Onda 1 (Comercial)

A alta gestão e as lideranças comerciais considerarão esta lista de Épicos concluída quando:

   1. Acurácia Cadastral: 100% dos novos leads e contas ativas no CRM possuírem o código IBGE do destino preenchido e validado.
   2. Transparência de Margem: O time de vendas (B2B/CRM) conseguir extrair o relatório de lucratividade líquida descontando o IBS de cada estado do país antes de assinar novos contratos.
   3. Conversão Comercial Protegida: A introdução do cálculo "por fora" não gerar aumento na taxa de abandono de carrinhos no e-commerce devido à lentidão ou falta de clareza na exibição do preço final.
