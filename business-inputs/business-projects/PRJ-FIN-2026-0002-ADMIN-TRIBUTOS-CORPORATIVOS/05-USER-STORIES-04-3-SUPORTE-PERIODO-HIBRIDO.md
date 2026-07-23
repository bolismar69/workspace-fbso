# Histórias de Usuário (User Stories) — Feature 04.3
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** Funcionalidade Transversal — Suporte ao Período Híbrido (Entrega 4)
* **Feature Relacionada:** [04.3 — Suporte ao Período Híbrido — Dupla Gestão de Regimes](./04-FEATURES.md#feature-043-suporte-ao-período-híbrido--dupla-gestão-de-regimes) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Visão Dual de Regimes no Painel e nos Formulários

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero visualizar em todas as telas do portal um indicador claro de qual regime cada alíquota pertence — "LEGADO" (ICMS, ISS, PIS, COFINS, IPI), "IVA DUAL" (CBS, IBS, IS) ou "TRANSITÓRIO" (regras específicas do período de convivência) —
Para operar com segurança durante o Período Híbrido (2029–2032), quando os dois modelos tributários coexistirão [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Indicador Visual de Regime): Toda alíquota exibe um selo de regime: "LEGADO" (cor cinza), "IVA DUAL" (cor azul), "TRANSITÓRIO" (cor laranja). O selo é visível no Painel de Alíquotas, nos formulários de cadastro/edição e nos relatórios.
* RN02 (Filtro por Regime): O Painel de Alíquotas Vigentes (Feature 01.1) inclui o filtro "Regime" com as três opções, permitindo isolar apenas as alíquotas de um regime específico.
* RN03 (Alerta Visual no Cadastro): Ao cadastrar uma nova alíquota durante o Período Híbrido, o formulário exibe um lembrete: "Período Híbrido vigente. Verifique se esta alíquota substitui ou complementa uma alíquota do regime legado."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Filtragem de alíquotas apenas do IVA Dual

* Dado que estamos em 2030 (Período Híbrido) e o portal contém alíquotas dos dois regimes;
* Quando o Analista Fiscal aplica o filtro "Regime = IVA DUAL" no Painel de Alíquotas;
* Então o sistema exibe apenas as alíquotas de CBS, IBS e IS;
* E todas exibem o selo azul "IVA DUAL" ao lado do nome do tributo [INDEX].

#### Cenário 2: Lembrete no cadastro durante Período Híbrido

* Dado que a data atual está dentro do Período Híbrido (2029–2032);
* Quando um Analista Fiscal abre o formulário de "Nova Alíquota";
* Então o topo do formulário exibe: "⚠️ Período Híbrido vigente. Verifique se esta alíquota substitui ou complementa uma alíquota do regime legado.";
* E o campo "Regime" vem pré-selecionado com "IVA DUAL" [INDEX].

------------------------------
## 📝 US-02: Mapeamento de Correlação entre Tributos Antigos e Novos

### 1. Descrição da História (Visão de Negócio)

Como Gerente Fiscal,
Quero definir e visualizar relações de substituição entre tributos do regime antigo e do novo regime (ex: "ICMS interestadual 12% para SP → MG será substituído por IBS estadual X% a partir de 01/01/2029"),
Para que o portal possa alertar sobre substituições ausentes e o time de Finanças tenha um mapa claro da transição [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Criação de Correlação): O Administrador Fiscal pode criar uma correlação entre um tributo legado e um tributo do IVA Dual, especificando: tributo de origem (legado), tributo de destino (IVA Dual), região geográfica, data de transição e justificativa (ex: "Conforme cronograma constitucional — redução de ICMS e início de IBS").
* RN05 (Indicador de Cobertura de Transição): O Dashboard de Cobertura Fiscal (Feature 04.2) exibe, para cada tributo legado, o percentual de alíquotas que já possuem correlação de substituição mapeada.
* RN06 (Alerta de Transição): Com 90 dias de antecedência de cada data de transição cadastrada, o portal gera um alerta no Painel e envia notificação ao Gerente Fiscal e ao Comitê Fiscal.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Criação de correlação ICMS → IBS

* Dado que o Administrador Fiscal acessa "Período Híbrido > Mapeamento de Correlação";
* Quando seleciona "Tributo Origem = ICMS", "UF Origem = SP", "UF Destino = MG", "Tributo Destino = IBS", "Data de Transição = 01/01/2029", preenche a justificativa e clica em "Salvar";
* Então o sistema cria a correlação e exibe: "Correlação criada: ICMS SP→MG (12%) será substituído por IBS a partir de 01/01/2029.";
* E o indicador de cobertura de transição para ICMS é atualizado [INDEX].

#### Cenário 2: Alerta de transição próxima

* Dado que existe uma correlação com data de transição em 01/01/2029;
* E a data atual é 03/10/2028 (90 dias antes);
* Quando o sistema executa a verificação diária de alertas;
* Então gera um alerta: "⏰ Transição programada: ICMS SP→MG será substituído por IBS em 01/01/2029 (90 dias). Verifique as alíquotas de IBS para MG.";
* E notifica o Gerente Fiscal e o Comitê Fiscal por e-mail [INDEX].

------------------------------
## 📝 US-03: Desativação Progressiva de Tributos em Extinção

### 1. Descrição da História (Visão de Negócio)

Como Administrador Fiscal,
Quero desativar em lote todas as alíquotas de um tributo do regime legado a partir de uma data específica, seguindo o cronograma constitucional de transição,
Para executar a descontinuação de forma controlada, com registro de auditoria individual para cada alíquota e aprovação em duas etapas [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN07 (Desativação em Lote por Tributo e Data): O Administrador Fiscal seleciona um tributo legado (ex: ICMS), uma data de desativação e opcionalmente uma UF (para desativação escalonada por estado). O sistema lista todas as alíquotas que serão afetadas antes de executar.
* RN08 (Aprovação Obrigatória): A desativação em lote é considerada uma operação de alto impacto e requer aprovação de um segundo Administrador Fiscal (fluxo de duas etapas, Feature 03.1), independentemente do valor financeiro envolvido.
* RN09 (Auditoria Individual): Cada alíquota desativada gera um registro de auditoria individual, com a justificativa "Desativação em lote — [tributo] extinto conforme cronograma constitucional. Operação [ID do lote] aprovada por [Administrador]."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Desativação em lote de ISS municipal

* Dado que é 01/01/2032 e o ISS foi totalmente substituído pelo IBS;
* Quando o Administrador Fiscal acessa "Período Híbrido > Desativação Progressiva", seleciona "Tributo = ISS", "Data = 01/01/2032" e clica em "Simular";
* Então o sistema lista as 23 alíquotas de ISS que serão desativadas, com município, valor e vigência;
* E exibe: "23 alíquotas de ISS serão desativadas. Esta operação requer aprovação de outro Administrador Fiscal.";
* Após aprovação, cada uma das 23 alíquotas recebe um registro de auditoria individual e muda o status para "Expirada" [INDEX].

#### Cenário 2: Desativação escalonada por UF

* Dado que a transição do ICMS é escalonada: primeiro SP, depois MG, depois demais estados;
* Quando o Administrador seleciona "Tributo = ICMS", "UF = SP" e "Data = 01/01/2029";
* Então o sistema lista apenas as alíquotas de ICMS para SP (excluindo outros estados);
* E após a aprovação, apenas as alíquotas de SP são desativadas, preservando as dos demais estados [INDEX].

------------------------------
## 📝 US-04: Visão Comparativa de Impacto entre Regimes (Informativo)

### 1. Descrição da História (Visão de Negócio)

Como Comitê Fiscal,
Quero visualizar uma comparação informativa de como uma mesma operação seria tributada no regime antigo e no novo regime,
Para avaliar impactos da transição no planejamento tributário e tomar decisões informadas sobre estratégias de precificação [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN10 (Simulação Comparativa): O usuário informa os parâmetros da operação (tributo, UF origem, UF destino, valor da operação) e o sistema calcula e exibe lado a lado: o valor do tributo no regime legado e o valor no IVA Dual, com a diferença absoluta e percentual.
* RN11 (Finalidade Exclusivamente Informativa): A visão comparativa não altera alíquotas, não gera registros fiscais e não se integra com a calculadora corporativa. É uma ferramenta de análise para o Comitê Fiscal.
* RN12 (Acesso Restrito): A visão comparativa é acessível apenas aos perfis Administrador Fiscal e Auditor/Controller.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Comparação de carga tributária ICMS vs. IBS

* Dado que o Comitê Fiscal acessa "Período Híbrido > Comparação de Regimes";
* Quando seleciona: "Operação = Venda interestadual", "UF Origem = SP", "UF Destino = MG", "Valor = R$ 100.000" e clica em "Comparar";
* Então o sistema exibe:
  * Regime LEGADO (ICMS): Alíquota interestadual 12% = R$ 12.000
  * Regime IVA DUAL (IBS): Alíquota destino MG = [X]% = R$ [Y]
  * Diferença: R$ [Y - 12.000] ([+/- Z]%)
* E exibe a nota: "Esta é uma simulação informativa. As alíquotas oficiais são as cadastradas no portal e utilizadas pela calculadora corporativa." [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
