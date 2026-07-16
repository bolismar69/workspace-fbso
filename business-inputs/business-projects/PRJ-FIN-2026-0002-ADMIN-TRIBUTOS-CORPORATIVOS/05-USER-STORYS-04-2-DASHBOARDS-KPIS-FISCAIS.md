# Histórias de Usuário (User Stories) — Feature 04.2
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** Funcionalidade Transversal — Dashboards Gerenciais de KPIs Fiscais (Entrega 4)
* **Feature Relacionada:** [04.2 — Dashboards Gerenciais de KPIs Fiscais](./04-FEATURES.md#feature-042-dashboards-gerenciais-de-kpis-fiscais) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Dashboard de Cobertura Fiscal — Mapa Geográfico de Alíquotas

### 1. Descrição da História (Visão de Negócio)

Como CFO,
Quero visualizar em um mapa do Brasil a cobertura de alíquotas de IBS por município, com indicação clara de quais municípios onde a companhia opera já possuem alíquota cadastrada e quais ainda estão descobertos,
Para identificar rapidamente lacunas de cobertura que representam risco fiscal e定向 o esforço do time de Finanças [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Escala de Cores): Municípios com alíquota cadastrada e vigente = verde; municípios com alíquota programada (data futura) = azul; municípios com operação ativa mas sem alíquota cadastrada = vermelho; municípios sem operação = cinza.
* RN02 (Filtro por Tributo): O dashboard permite alternar entre IBS, ICMS (regime legado) e "Visão Híbrida" (mostra ambos os regimes simultaneamente com cores diferentes).
* RN03 (Drill-Down): Ao clicar em um município, o sistema exibe: nome do município, código IBGE, faturamento estimado, alíquotas vigentes (com valores e vigências) ou a indicação "Sem alíquota cadastrada".

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: CFO identifica municípios descobertos em SP

* Dado que o CFO acessa o Dashboard de Cobertura Fiscal e seleciona "Tributo = IBS";
* Quando o mapa carrega, o estado de SP exibe diversos municípios em verde, mas 3 municípios onde a companhia tem operação ativa aparecem em vermelho;
* Ao clicar em um município vermelho, o sistema exibe: "São José dos Campos/SP — Código IBGE 3549904 — Faturamento estimado: R$ 320.000/mês — ⚠️ Sem alíquota de IBS cadastrada";
* Então o CFO encaminha o alerta ao Gerente Fiscal para priorizar o cadastro [INDEX].

------------------------------
## 📝 US-02: Dashboard de Atividade e Governança

### 1. Descrição da História (Visão de Negócio)

Como Gerente Fiscal,
Quero visualizar em um painel a atividade mensal de gestão tributária (volume de alterações por tipo e por tributo), os indicadores de governança (cobertura de trilha, conflitos prevenidos) e o status das aprovações de alto impacto pendentes,
Para monitorar a produtividade do time e a saúde dos controles internos em uma única tela [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Indicadores em Tempo Real): Os dados do dashboard refletem a situação atual do portal no momento do carregamento. Uma alteração feita 1 minuto antes de o dashboard ser carregado já aparece nos indicadores.
* RN05 (Comparativo Mensal): O gráfico de volume de alterações exibe os últimos 6 meses para permitir análise de tendência (crescente = adoção do portal aumentando; estável = operação madura).
* RN06 (Alertas de Pendências): Se houver alterações de alto impacto aguardando aprovação há mais de 3 dias úteis, o dashboard exibe um alerta no topo com o número de itens pendentes e link para a fila de aprovação.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Monitoramento mensal da atividade

* Dado que o Gerente Fiscal acessa o Dashboard de Atividade;
* Quando o dashboard carrega;
* Então exibe: "Julho/2026: 47 alterações (32 criações, 10 edições, 5 desativações)";
* E o gráfico de barras mostra a evolução mensal dos últimos 6 meses;
* E o indicador de cobertura de trilha exibe "100% — 47/47 alterações com registro" (KPI G1);
* E o indicador de conflitos prevenidos exibe "3 conflitos bloqueados no mês" (KPI G2) [INDEX].

#### Cenário 2: Alerta de aprovações pendentes

* Dado que existem 2 alterações de alto impacto aguardando aprovação há 5 dias úteis;
* Quando o Gerente Fiscal acessa o Dashboard de Atividade;
* Então o topo do dashboard exibe o alerta "⚠️ 2 alterações de alto impacto aguardando aprovação há mais de 3 dias úteis. [Ir para a fila de aprovação]";
* E ao clicar no link, é direcionado para a fila de aprovação (Feature 03.1) [INDEX].

------------------------------
## 📝 US-03: Dashboard de Vencimentos — Alíquotas a Expirar

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero visualizar uma linha do tempo das próximas 12 semanas com as alíquotas que expirarão no período, agrupadas por semana e por tributo,
Para planejar proativamente as renovações e evitar que alíquotas expirem sem substituição, causando descobertura fiscal [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN07 (Janela de 12 Semanas): O dashboard exibe as alíquotas com data de término de vigência nos próximos 90 dias, agrupadas por semana.
* RN08 (Ordenação por Urgência): As semanas são exibidas da esquerda para a direita (semana atual + 11 seguintes). Alíquotas que expiram na semana atual são destacadas em vermelho.
* RN09 (Ação Rápida): Ao clicar em uma alíquota na linha do tempo, o sistema oferece as opções: "Renovar" (abre o formulário de edição para estender a vigência) ou "Substituir" (abre o formulário de nova alíquota pré-preenchido com os dados da alíquota atual).

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Identificação de alíquotas críticas na semana atual

* Dado que 5 alíquotas expiram na semana atual (3 de ICMS, 2 de ISS);
* Quando o Analista Fiscal acessa o Dashboard de Vencimentos;
* Então a coluna da semana atual deve exibir "5 alíquotas" em vermelho;
* E ao expandir a semana, as 5 alíquotas são listadas com tributo, UF/município, valor e data exata de expiração;
* E cada uma tem o botão "Renovar" ao lado [INDEX].

#### Cenário 2: Ação rápida de renovação

* Dado que uma alíquota de CBS expira em 7 dias;
* Quando o Analista clica em "Renovar" na linha do tempo;
* Então o sistema abre o formulário de edição com a data de fim pré-preenchida como "31/12/2027" (sugestão de extensão padrão de 1 ano);
* E o Analista ajusta a data se necessário e clica em "Salvar";
* E a alíquota permanece vigente e some do Dashboard de Vencimentos [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
