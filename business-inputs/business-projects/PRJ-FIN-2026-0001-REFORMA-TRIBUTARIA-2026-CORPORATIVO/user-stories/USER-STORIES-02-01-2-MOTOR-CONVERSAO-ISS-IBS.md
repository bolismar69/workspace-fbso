# Histórias de Usuário (User Stories) — Feature 01.2 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: ERP SAP SD — Faturamento de Serviços (Onda 2)
- Feature Relacionada: 02.01.2 — Motor de Conversão do ISS para IBS (Serviços Santana de Parnaíba) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Descontinuação da Retenção de ISS e Ativação do IBS para Serviços

### 1. Descrição da História (Visão de Negócio)

Como Analista de Faturamento de Serviços da Matriz em Santana de Parnaíba,
Quero que o SAP desative automaticamente a lógica de retenção de ISS na competência municipal de origem e ative a aplicação do IBS com base no destino do serviço,
Para que as notas fiscais de serviço da matriz estejam em conformidade com a Reforma Tributária e o ISS legado não seja mais retido nas operações [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Data de Corte para Transição): A partir da data definida pelo calendário regulatório federal para início da cobrança do IBS sobre serviços, o SAP deve desativar a condição de imposto "ISS" em todas as operações de serviço da matriz e ativar a condição "IBS — Serviços".
* RN02 (Migração da NFS-e Municipal para Nacional): O fluxo de emissão deve migrar do webservice da prefeitura de Santana de Parnaíba para o padrão nacional da NFS-e integrado ao ambiente da SEFAZ/Receita Federal.
* RN03 (Alíquota de IBS por Destino do Tomador): A alíquota de IBS aplicada deve ser a do município e estado de domicílio do tomador do serviço (princípio do destino), e não mais a alíquota do município prestador (origem).
* RN04 (Convivência Temporária): Durante o Período Híbrido (2029-2032), o SAP deve ser capaz de emitir tanto notas com ISS (legado) quanto com IBS (novo), selecionando o regime correto com base na data da operação e na natureza do serviço.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Faturamento de serviço após a data de corte — IBS aplicado no destino

* Dado que a data de corte para início do IBS sobre serviços já foi atingida;
* E a matriz de Santana de Parnaíba (SP) emite uma NFS-e para um cliente tomador em Florianópolis (SC);
* E o valor do serviço é R$ 20.000,00;
* Quando o analista de faturamento processar a ordem de serviço no SAP;
* Então o sistema deve ignorar a condição de ISS (origem SP);
* E aplicar o IBS do destino: alíquota de Florianópolis (SC);
* E emitir a NFS-e Nacional com o IBS destacado conforme legislação do IVA Dual [INDEX].

#### Cenário 2: Período Híbrido — serviço tributado pelo regime antigo (data retroativa)

* Dado que o sistema está operando em 2030 durante o Período Híbrido;
* E uma ordem de serviço com data retroativa a 2028 precisa ser faturada;
* Quando o analista informar a data de competência como 15/03/2028 (anterior à data de corte);
* Então o SAP deve identificar que a data é anterior à transição;
* E aplicar o ISS legado (alíquota de Santana de Parnaíba) em vez do IBS;
* E emitir a NFS-e Municipal (modelo antigo) para a prefeitura de Santana de Parnaíba [INDEX].

#### Cenário 3: Serviço interestadual — alíquota de IBS do destino corretamente aplicada

* Dado que um serviço de consultoria é prestado pela matriz para um cliente em Recife (PE);
* E a alíquota de IBS para serviços em Recife é diferente da alíquota de IBS para serviços em São Paulo;
* Quando o faturamento for processado;
* Então o SAP deve consultar a tabela de alíquotas de IBS municipal e aplicar especificamente a alíquota de Recife (PE);
* E o valor do IBS recolhido deve ser destinado ao município e estado de destino (PE), não ao de origem (SP) [INDEX].

------------------------------
## 📝 US-02: Bloqueio de Dupla Tributação ISS + IBS na Mesma Operação

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Controladoria e Risco Fiscal,
Quero que o SAP impeça ativamente que uma mesma operação de serviço seja tributada simultaneamente por ISS (regime antigo) e IBS (regime novo),
Para evitar dupla tributação que gere prejuízo financeiro e risco de autuação fiscal [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Exclusividade de Regime): Uma operação de serviço só pode ser tributada por ISS OU por IBS, nunca por ambos simultaneamente. O SAP deve validar essa exclusividade antes da emissão da nota.
* RN02 (Regra de Determinação de Regime): O regime aplicável é determinado pela data de competência da operação: datas anteriores à transição → ISS; datas posteriores → IBS.
* RN03 (Bloqueio com Alerta Explícito): Se o SAP detectar que ambas as condições de imposto (ISS e IBS) estão ativas para a mesma operação, deve bloquear a emissão e exibir o alerta: "ERRO: Conflito de regimes tributários. A operação possui ISS e IBS simultaneamente ativos. Corrija as condições de imposto antes de prosseguir."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: SAP detecta ISS e IBS ativos simultaneamente e bloqueia emissão

* Dado que um faturista está processando uma NFS-e de serviço;
* E por erro de parametrização, tanto a condição de ISS (2%) quanto de IBS (17,7%) estão ativas na mesma operação;
* Quando o faturista tentar emitir a nota;
* Então o SAP deve exibir o alerta de conflito de regimes e bloquear a emissão;
* E instruir o faturista a revisar as condições de imposto ou contatar a Controladoria [INDEX].

#### Cenário 2: Operação com regime único — emissão liberada

* Dado que o SAP validou a operação de serviço e confirmou que apenas a condição de IBS está ativa;
* Quando o faturista prosseguir com a emissão;
* Então o sistema deve liberar a NFS-e Nacional normalmente, sem bloqueios;
* E registrar no log de auditoria: "Operação validada — Regime exclusivo: IBS. Sem conflito com ISS" [INDEX].
