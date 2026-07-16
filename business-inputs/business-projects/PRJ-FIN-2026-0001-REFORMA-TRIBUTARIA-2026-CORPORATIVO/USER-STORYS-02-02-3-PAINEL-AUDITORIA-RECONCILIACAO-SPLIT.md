# Histórias de Usuário (User Stories) — Feature 02.3 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulos: ERP SAP FI — Tesouraria e Controladoria (Onda 2)
- Feature Relacionada: 02.02.3 — Painel de Auditoria e Reconciliação do Split Payment [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Painel de Conferência — Cruzamento NF-e × Retenção Bancária

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Controladoria e Auditoria Interna,
Quero acessar um painel que cruza automaticamente o valor de CBS/IBS destacado em cada Nota Fiscal com o valor efetivamente retido pelo banco no Split Payment,
Para identificar discrepâncias de retenção em tempo real e acionar as instituições financeiras antes que o erro se acumule [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Visão Consolidada por Período): O painel deve permitir filtrar por período (diário, semanal, mensal), por instituição financeira (banco), por UF de destino e por tipo de operação (normal vs. incentivada).
* RN02 (Indicador Visual de Divergência): Cada lançamento deve exibir um indicador colorido: 🟢 verde para conciliação OK (diferença = R$ 0,00), 🟡 amarelo para diferença de até R$ 5,00, 🔴 vermelho para diferença superior a R$ 5,00.
* RN03 (Drill-Down ao Detalhe da NF-e): Ao clicar em um lançamento com divergência, o usuário deve visualizar o detalhamento completo: número da NF-e, data de emissão, valor destacado de CBS, valor destacado de IBS, banco responsável, valor retido de CBS, valor retido de IBS, diferença apurada.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Controller filtra painel por banco e identifica divergências

* Dado que o Gerente de Controladoria acessa o painel de auditoria do Split Payment;
* E seleciona o filtro: Banco "XYZ", Período "Últimos 7 dias";
* Quando o painel carregar os lançamentos;
* Então devem ser exibidos 150 lançamentos processados pelo Banco XYZ na semana;
* E o resumo deve mostrar: "145 OK (🟢) / 3 com divergência leve (🟡) / 2 com divergência crítica (🔴)";
* E as 2 divergências críticas devem estar destacadas no topo da lista para ação imediata [INDEX].

#### Cenário 2: Controller investiga divergência crítica via drill-down

* Dado que o painel exibe uma divergência crítica de R$ 150,00 para a NF-e #7890;
* Quando o Controller clicar no lançamento para expandir o detalhe;
* Então o sistema deve exibir:
  - NF-e: #7890 — Emitida em 22/06/2026
  - CBS Destacado: R$ 880,00 | CBS Retido pelo Banco: R$ 880,00 ✅
  - IBS Destacado: R$ 1.770,00 | IBS Retido pelo Banco: R$ 1.920,00 ❌
  - Diferença IBS: +R$ 150,00 (Banco reteve a maior)
  - Instituição: Banco XYZ — Agência 1234
* E disponibilizar o botão "Gerar Notificação para o Banco" [INDEX].

------------------------------
## 📝 US-02: Geração Automática de Notificação de Divergência para Instituição Financeira

### 1. Descrição da História (Visão de Negócio)

Como Analista de Contas a Receber,
Quero que o SAP gere automaticamente uma notificação padronizada para a instituição financeira quando for detectada uma divergência de split payment superior a R$ 5,00,
Para agilizar o processo de contestação e estorno do valor retido indevidamente [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Gatilho de Notificação Automática): Divergências de split superiores a R$ 5,00 devem disparar automaticamente a geração de um documento de notificação (PDF) contendo: dados da NF-e, valor destacado, valor retido, diferença apurada e solicitação formal de estorno.
* RN02 (Template Padronizado por Banco): Cada instituição financeira parceira pode ter um template de contestação específico, que deve ser cadastrado e selecionado automaticamente pelo SAP com base no banco responsável pela retenção.
* RN03 (SLA de Acompanhamento): Após o envio da notificação, o SAP deve monitorar se o estorno foi processado pelo banco em até 5 dias úteis. Se o prazo expirar sem resolução, a pendência escala para o Gerente de Tesouraria.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Sistema gera notificação para divergência de IBS

* Dado que o painel detectou uma divergência de R$ 150,00 de IBS retido a maior pelo Banco XYZ na NF-e #7890;
* Quando o Analista clicar em "Gerar Notificação para o Banco";
* Então o SAP deve gerar um PDF com o template do Banco XYZ contendo todos os dados da divergência;
* E registrar a notificação no sistema com o protocolo "NOT-SPLIT-20260623-001";
* E iniciar o contador de SLA de 5 dias úteis para acompanhamento;
* E enviar o PDF por e-mail para o canal de contestação do banco cadastrado no sistema [INDEX].

#### Cenário 2: SLA de 5 dias expira sem resposta do banco — escalação

* Dado que a notificação NOT-SPLIT-20260623-001 foi enviada ao Banco XYZ há 5 dias úteis;
* E o banco ainda não processou o estorno de R$ 150,00;
* Quando o relógio do sistema marcar o vencimento do SLA;
* Então o SAP deve escalar a pendência para o Gerente de Tesouraria;
* E enviar um e-mail automático: "NOT-SPLIT-20260623-001 — SLA de 5 dias expirado. Banco XYZ não respondeu à contestação de R$ 150,00 referente à NF-e #7890.";
* E alterar o status da notificação para "Escalado — Aguardando Intervenção Gerencial" [INDEX].

------------------------------
## 📝 US-03: Relatório Gerencial de Eficiência do Split Payment para o CFO

### 1. Descrição da História (Visão de Negócio)

Como CFO da companhia,
Quero receber um relatório executivo mensal consolidando a performance do Split Payment em todas as instituições financeiras parceiras,
Para monitorar a acurácia das retenções, identificar bancos com alta taxa de divergência e decidir sobre a manutenção ou substituição de parceiros financeiros [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Consolidação por Banco): O relatório deve apresentar, para cada instituição financeira: volume total de transações processadas, valor total de CBS/IBS retido, número de divergências, percentual de acurácia, tempo médio de resolução de pendências.
* RN02 (Ranking de Eficiência): Os bancos devem ser ranqueados por percentual de acurácia, destacando os 3 melhores (verde) e os que estiverem abaixo de 98% (vermelho).
* RN03 (Distribuição Automática): O relatório deve ser gerado automaticamente no 1º dia útil de cada mês e distribuído por e-mail para CFO, Gerente de Tesouraria e Controller.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Geração do relatório mensal de eficiência do Split Payment

* Dado que é o 1º dia útil do mês de julho de 2026;
* Quando o job agendado de geração de relatório for executado;
* Então o SAP deve consolidar os dados de junho/2026 e gerar o relatório "Eficiência do Split Payment — Junho/2026";
* E o relatório deve mostrar:
  - Banco XYZ: 2.450 transações / R$ 3,2M retidos / 99,1% acurácia / 22 divergências resolvidas
  - Banco ABC: 1.890 transações / R$ 2,1M retidos / 96,3% acurácia / 70 divergências (5 pendentes)
* E destacar o Banco ABC em vermelho com a recomendação: "Revisar parceria — acurácia abaixo de 98%";
* E enviar o PDF por e-mail para a lista de distribuição executiva [INDEX].
