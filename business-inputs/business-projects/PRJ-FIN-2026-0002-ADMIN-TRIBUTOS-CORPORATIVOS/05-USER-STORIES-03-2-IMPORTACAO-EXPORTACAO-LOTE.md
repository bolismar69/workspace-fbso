# Histórias de Usuário (User Stories) — Feature 03.2
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** M5 — Importação e Exportação em Lote (Entrega 3)
* **Feature Relacionada:** [03.2 — Importação e Exportação de Alíquotas em Lote](./04-FEATURES.md#feature-032-importação-e-exportação-de-alíquotas-em-lote) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Download de Template e Upload de Planilha de Alíquotas

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero baixar um template padronizado, preencher as alíquotas em lote e fazer o upload da planilha no portal,
Para processar centenas ou milhares de alíquotas de uma só vez — como quando o Comitê Gestor do IBS publica as alíquotas municipais — em vez de cadastrar uma a uma manualmente [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Template Padronizado): O portal disponibiliza um arquivo template no formato Excel (.xlsx) com as colunas obrigatórias: Tributo, UF, Código IBGE, NCM (opcional), Data Início, Data Fim (opcional), Valor (%), Justificativa. A primeira linha do arquivo contém os nomes das colunas e uma segunda linha de exemplo.
* RN02 (Validação de Estrutura): No upload, o sistema verifica se o arquivo possui as colunas obrigatórias com os nomes exatos. Se a estrutura divergir, o upload é rejeitado com mensagem indicando as colunas faltantes ou excedentes.
* RN03 (Formatos Aceitos): Apenas arquivos nos formatos .xlsx e .csv são aceitos.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Download do template e upload bem-sucedido

* Dado que o Analista Fiscal acessa "Importação em Lote";
* Quando clica em "Baixar Template" e o sistema gera o arquivo "template_aliquotas.xlsx";
* E o Analista preenche 100 linhas com alíquotas de IBS para 100 municípios de SP;
* E faz o upload do arquivo preenchido;
* Então o sistema deve aceitar o arquivo e iniciar o processamento, exibindo uma barra de progresso [INDEX].

#### Cenário 2: Rejeição de arquivo com estrutura inválida

* Dado que o Analista tenta fazer upload de um arquivo que não contém a coluna "Código IBGE";
* Quando o upload é iniciado;
* Então o sistema deve rejeitar o arquivo e exibir: "Estrutura do arquivo inválida. Coluna obrigatória ausente: 'Código IBGE'. Baixe o template padrão e preencha conforme o formato esperado." [INDEX].

------------------------------
## 📝 US-02: Processamento com Validação Linha a Linha e Relatório de Resultados

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero que o portal processe cada linha da planilha aplicando as mesmas validações de negócio do cadastro manual e, ao final, me apresente um relatório claro de quais linhas foram aceitas e quais foram rejeitadas — com o motivo específico da rejeição,
Para que eu possa corrigir apenas os registros problemáticos e reenviá-los, em vez de revisar manualmente centenas de alíquotas [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Validação Individual): Cada linha da planilha é validada contra as mesmas regras do cadastro manual (conflito de vigência, integridade de classificação, consistência de transição, datas). Linhas que violam regras são rejeitadas individualmente, sem interromper o processamento das demais.
* RN05 (Relatório de Processamento): Ao final do processamento, o sistema exibe: total de linhas processadas, total de aceitas, total de rejeitadas e uma tabela detalhada das rejeições com: número da linha no arquivo original, conteúdo da linha, motivo da rejeição.
* RN06 (Arquivo de Rejeitados): O portal gera um arquivo .xlsx contendo apenas as linhas rejeitadas, com uma coluna adicional "Motivo da Rejeição", para que o Analista possa corrigi-las e reenviá-las.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Processamento com aceitação parcial

* Dado que o Analista fez upload de uma planilha com 100 alíquotas de IBS para municípios de SP;
* E 92 linhas estão corretas e 8 linhas possuem conflitos de vigência com alíquotas já existentes;
* Quando o processamento é concluído;
* Então o sistema deve exibir: "Processamento concluído: 100 linhas processadas, 92 aceitas, 8 rejeitadas.";
* E na tabela de rejeições, para cada linha rejeitada, exibir o número da linha, o município e o motivo: "Conflito de vigência com alíquota existente para o mesmo tributo e município";
* E disponibilizar o download do arquivo de rejeitados [INDEX].

#### Cenário 2: Processamento com 100% de aceitação

* Dado que o Analista fez upload de uma planilha com 50 alíquotas, todas válidas;
* Quando o processamento é concluído;
* Então o sistema deve exibir: "Processamento concluído: 50 linhas processadas, 50 aceitas, 0 rejeitadas.";
* E cada linha aceita deve gerar um registro individual na trilha de auditoria (tipo "Criação em Lote") [INDEX].

#### Cenário 3: Histórico de cargas em lote

* Dado que 3 operações de importação em lote foram realizadas neste mês;
* Quando o Gerente Fiscal acessa a Linha do Tempo e filtra por "Operação = Importação em Lote";
* Então o sistema deve listar as 3 operações, cada uma com: data, usuário, arquivo original (link para download) e resumo do resultado (X aceitas, Y rejeitadas) [INDEX].

------------------------------
## 📝 US-03: Exportação de Tabelas Vigentes para Conformidade e Auditoria

### 1. Descrição da História (Visão de Negócio)

Como Controller,
Quero exportar as tabelas de alíquotas vigentes em formato de planilha, com opção de incluir ou não o histórico de alterações,
Para anexar a documentação fiscal atualizada em relatórios de conformidade, auditorias externas e prestações de contas a órgãos reguladores [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN07 (Exportação Filtrada): A exportação respeita os filtros atualmente aplicados no Painel de Alíquotas — o usuário filtra primeiro, depois exporta.
* RN08 (Opções de Exportação): O usuário pode escolher entre "Apenas Vigentes" (padrão), "Vigentes + Programadas" (inclui alíquotas com data de início futura) ou "Completo" (inclui também expiradas, para auditoria histórica).
* RN09 (Formato de Saída): O arquivo é gerado em formato .xlsx com colunas: Tributo, UF, Código IBGE, Município, NCM, Data Início, Data Fim, Valor (%), Status, Última Alteração (data e usuário).

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Exportação de alíquotas vigentes de IBS para SP

* Dado que o Controller aplicou os filtros "Tributo = IBS" e "UF = SP" no Painel de Alíquotas;
* Quando clica em "Exportar" e seleciona "Apenas Vigentes";
* Então o sistema deve gerar um arquivo "aliquotas_IBS_SP_vigentes_[data].xlsx";
* E o arquivo deve conter apenas as alíquotas de IBS vigentes para municípios de SP [INDEX].

#### Cenário 2: Exportação completa para auditoria anual

* Dado que a Auditoria Interna solicitou a base completa de alíquotas para o exercício de 2026;
* Quando o Controller seleciona "Exportar > Completo" e define o período de 01/01/2026 a 31/12/2026;
* Então o sistema deve gerar um arquivo incluindo alíquotas vigentes, programadas e expiradas no período;
* E a coluna "Última Alteração" deve conter a data e o usuário da alteração mais recente de cada alíquota [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
