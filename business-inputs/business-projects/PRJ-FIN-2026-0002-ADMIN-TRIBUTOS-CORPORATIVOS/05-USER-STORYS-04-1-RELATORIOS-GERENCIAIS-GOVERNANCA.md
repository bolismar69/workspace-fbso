# Histórias de Usuário (User Stories) — Feature 04.1
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** Funcionalidade Transversal — Relatórios Gerenciais de Governança (Entrega 4)
* **Feature Relacionada:** [04.1 — Relatórios Gerenciais de Governança](./04-FEATURES.md#feature-041-relatórios-gerenciais-de-governança) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Geração Automática do Relatório Mensal de Governança

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Controladoria,
Quero que o portal gere automaticamente, no primeiro dia útil de cada mês, um relatório sumarizando todas as alterações em tabelas fiscais do mês anterior,
Para apresentar ao Comitê Fiscal e à Controladoria sem que ninguém precise compilar manualmente dados de múltiplas fontes [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Agendamento): O relatório é gerado automaticamente às 07:00 do primeiro dia útil de cada mês, cobrindo o período do primeiro ao último dia do mês anterior.
* RN02 (Conteúdo do Relatório): O relatório contém: (a) sumário executivo com totais de criações, edições e desativações no período; (b) tabela de alterações por tributo; (c) tabela de alterações por usuário responsável; (d) seção de alterações de alto impacto aprovadas; (e) seção de alterações de alto impacto rejeitadas; (f) lista de conflitos de vigência detectados no período.
* RN03 (Distribuição Automática): O relatório é enviado por e-mail para a lista de distribuição cadastrada (Controladoria, Gerente Fiscal, Controller, Comitê Fiscal) e fica disponível para download no portal por 24 meses.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Geração do relatório de julho/2026

* Dado que é 1º de agosto de 2026 (primeiro dia útil);
* Quando o sistema executa a rotina automática de geração;
* Então o relatório "Relatório Mensal de Governança — Julho 2026" deve ser gerado em PDF;
* E o Controller deve recebê-lo por e-mail até as 08:00;
* E ao acessar o portal, o relatório deve estar disponível na seção "Relatórios" [INDEX].

#### Cenário 2: Conteúdo do relatório confere com a trilha de auditoria

* Dado que o relatório de julho/2026 foi gerado;
* Quando o Controller confere o total de alterações do relatório (ex: 47 alterações em julho) com o total de registros na Linha do Tempo para o mesmo período;
* Então os números devem ser idênticos — o relatório e a trilha de auditoria compartilham a mesma fonte de dados [INDEX].

------------------------------
## 📝 US-02: Geração de Relatório Ad Hoc para Período Específico

### 1. Descrição da História (Visão de Negócio)

Como Controller,
Quero gerar um relatório de governança para um período específico que não corresponde ao mês calendário — por exemplo, para atender a uma solicitação de auditoria externa ou investigar um incidente,
Para obter a mesma qualidade de sumarização do relatório mensal, mas para qualquer intervalo de datas que eu definir [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Período Customizável): O relatório ad hoc permite selecionar qualquer data de início e data de fim. O período mínimo é de 1 dia; o máximo é de 12 meses.
* RN05 (Mesmo Formato): O relatório ad hoc segue exatamente o mesmo formato e conteúdo do relatório mensal automático (RN02).
* RN06 (Registro de Geração): Toda geração de relatório ad hoc é registrada na trilha de auditoria com: usuário solicitante, data/hora e período solicitado.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Relatório ad hoc para auditoria externa

* Dado que a Auditoria Externa solicitou um relatório de governança do período 15/03/2026 a 15/06/2026;
* Quando o Controller acessa "Relatórios > Gerar Relatório Ad Hoc", define o período e clica em "Gerar";
* Então o sistema deve gerar o relatório no mesmo formato do relatório mensal, cobrindo exatamente o período solicitado;
* E registrar na trilha: "Relatório ad hoc gerado por [Controller] para o período 15/03/2026 a 15/06/2026" [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
