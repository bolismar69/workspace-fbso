# Histórias de Usuário (User Stories) — Feature 01.1
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** M1 — Painel de Alíquotas Vigentes (Entrega 1)
* **Feature Relacionada:** [01.1 — Painel de Alíquotas Vigentes](./04-FEATURES.md#feature-011-painel-de-alíquotas-vigentes) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Consulta ao Painel de Alíquotas com Filtros

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero consultar o Painel de Alíquotas Vigentes com filtros por tributo, região, período e status,
Para localizar rapidamente as alíquotas aplicáveis a uma operação específica sem percorrer planilhas ou abrir chamados técnicos [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Filtros Combináveis): Os filtros de tributo, UF, município (código IBGE), período de vigência e status (vigente, programada, expirada) devem ser combináveis entre si, permitindo refinar a consulta progressivamente.
* RN02 (Ordenação Padrão): A lista de alíquotas deve ser ordenada por: primeiro alíquotas com alertas ativos (prestes a expirar, conflitos), depois por tributo (ordem alfabética) e então por data de início de vigência (mais recente primeiro).
* RN03 (Paginação): O painel deve exibir resultados paginados, com 50 alíquotas por página e indicador do total de registros encontrados.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Consulta com filtro de tributo e UF

* Dado que o Analista Fiscal acessa o Painel de Alíquotas Vigentes;
* Quando selecionar o filtro "Tributo = IBS" e "UF = SP";
* Então o sistema deve exibir apenas as alíquotas de IBS vigentes para municípios do estado de São Paulo;
* E exibir o total de registros encontrados no cabeçalho da tabela [INDEX].

#### Cenário 2: Consulta sem filtros — visão geral

* Dado que o Analista Fiscal acessa o Painel de Alíquotas Vigentes;
* Quando não aplicar nenhum filtro;
* Então o sistema deve exibir todas as alíquotas vigentes, ordenadas primeiro pelos alertas ativos;
* E as alíquotas com indicador de "a expirar em 30 dias" devem aparecer no topo da lista com destaque visual (ícone ⚠️) [INDEX].

------------------------------
## 📝 US-02: Indicadores Visuais de Integridade das Alíquotas

### 1. Descrição da História (Visão de Negócio)

Como Gerente Fiscal,
Quero visualizar indicadores de integridade no Painel de Alíquotas que sinalizem automaticamente alíquotas prestes a expirar, conflitos de vigência e ausência de substituição no Período Híbrido,
Para agir proativamente antes que uma expiração ou conflito cause impacto nas operações de venda e faturamento [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Alerta de Expiração): Alíquotas cuja data de término de vigência esteja a 30 dias ou menos devem exibir o indicador "⚠️ A expirar" na coluna de status, independentemente dos filtros aplicados.
* RN05 (Alerta de Conflito): Quando existirem duas ou mais alíquotas para o mesmo tributo e mesma região geográfica com períodos de vigência que se sobreponham, ambas devem exibir o indicador "🔴 Conflito" e um link para visualizar a alíquota conflitante.
* RN06 (Alerta de Transição — Período Híbrido): Durante o Período Híbrido (2029–2032), alíquotas do regime antigo (ICMS, ISS, PIS, COFINS, IPI) que não possuem alíquota substituta mapeada no novo regime devem exibir o indicador "🟡 Sem substituição".

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Alerta de alíquota prestes a expirar

* Dado que existe uma alíquota de CBS com vigência até 15/08/2026;
* E a data atual é 20/07/2026 (menos de 30 dias para expiração);
* Quando o Gerente Fiscal acessa o Painel de Alíquotas;
* Então essa alíquota deve exibir o indicador "⚠️ A expirar em 26 dias" com cor laranja [INDEX].

#### Cenário 2: Alerta de conflito entre alíquotas sobrepostas

* Dado que existe uma alíquota de IBS para o município 3547304 (Santana de Parnaíba/SP) vigente de 01/01/2026 a 31/12/2026;
* E foi cadastrada uma nova alíquota de IBS para o mesmo município vigente de 01/06/2026 a 30/06/2027;
* Quando o Gerente Fiscal acessa o Painel de Alíquotas;
* Então ambas as alíquotas devem exibir o indicador "🔴 Conflito" na coluna de status;
* E ao clicar no indicador, o sistema deve exibir as duas alíquotas lado a lado com a interseção de períodos destacada [INDEX].

------------------------------
## 📝 US-03: Acesso ao Painel por Perfil de Usuário

### 1. Descrição da História (Visão de Negócio)

Como Controller (perfil Auditor/Controller),
Quero acessar o Painel de Alíquotas Vigentes com capacidade de consulta e exportação, mas sem qualquer permissão de alteração,
Para obter as informações fiscais necessárias para conciliação e auditoria sem risco de modificar dados acidentalmente [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN07 (Acesso Universal de Consulta): O Painel de Alíquotas Vigentes é acessível a todos os três perfis (Administrador Fiscal, Analista Fiscal, Auditor/Controller) com as mesmas capacidades de visualização e filtro.
* RN08 (Restrição de Ação por Perfil): O perfil Auditor/Controller não visualiza botões ou links de criação, edição ou desativação de alíquotas no painel. O perfil Analista Fiscal visualiza os botões de criação e edição, mas não de desativação de alíquotas criadas por outros (a desativação segue regras do fluxo de aprovação).

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Auditor acessa o painel sem controles de edição

* Dado que um usuário com perfil Auditor/Controller está autenticado no portal;
* Quando acessa o Painel de Alíquotas Vigentes;
* Então o sistema deve exibir todas as alíquotas com os mesmos filtros disponíveis aos demais perfis;
* E não deve exibir nenhum botão de "Nova Alíquota", "Editar" ou "Desativar";
* E deve exibir o botão "Exportar" para extração dos dados em planilha [INDEX].

#### Cenário 2: Analista Fiscal visualiza controles de cadastro

* Dado que um usuário com perfil Analista Fiscal está autenticado;
* Quando acessa o Painel de Alíquotas Vigentes;
* Então o sistema deve exibir o botão "Nova Alíquota" no topo do painel;
* E deve exibir o botão "Editar" em cada linha de alíquota;
* E deve exibir o botão "Desativar" apenas nas alíquotas que o próprio analista cadastrou [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
