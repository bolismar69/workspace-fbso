# Histórias de Usuário (User Stories) — Feature 02.1
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** M6 — Administração de Acessos e Perfis (Entrega 2)
* **Feature Relacionada:** [02.1 — Administração de Acessos e Perfis](./04-FEATURES.md#feature-021-administração-de-acessos-e-perfis) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Cadastro de Usuários e Atribuição de Perfil

### 1. Descrição da História (Visão de Negócio)

Como Administrador Fiscal,
Quero cadastrar novos usuários no portal e atribuir a cada um o perfil de acesso correspondente à sua função (Administrador Fiscal, Analista Fiscal ou Auditor/Controller),
Para que o time de Finanças possa operar o portal dentro dos perímetros de acesso definidos pela política de segregação de funções [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Três Perfis Disponíveis): O sistema oferece exatamente três perfis: Administrador Fiscal, Analista Fiscal e Auditor/Controller. Não é possível criar perfis customizados.
* RN02 (Cadastro Exclusivo pelo Administrador Fiscal): Apenas usuários com perfil Administrador Fiscal podem acessar o módulo de administração de usuários e realizar cadastros.
* RN03 (Dados Obrigatórios): Nome completo, e-mail corporativo e perfil são campos obrigatórios para o cadastro de um novo usuário.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cadastro de novo Analista Fiscal

* Dado que o Administrador Fiscal está na tela de "Administração de Usuários";
* Quando preencher nome "Maria Silva", e-mail "maria.silva@companhia.com.br", selecionar perfil "Analista Fiscal" e clicar em "Cadastrar";
* Então o sistema deve criar o usuário e enviar um e-mail de boas-vindas com link para definir a senha inicial;
* E o novo usuário deve aparecer na lista de usuários ativos com perfil "Analista Fiscal" [INDEX].

#### Cenário 2: Tentativa de cadastro com e-mail duplicado

* Dado que já existe um usuário com e-mail "maria.silva@companhia.com.br";
* Quando o Administrador tentar cadastrar outro usuário com o mesmo e-mail;
* Então o sistema deve exibir: "Já existe um usuário cadastrado com o e-mail maria.silva@companhia.com.br." [INDEX].

------------------------------
## 📝 US-02: Gestão de Perfis e Permissões — Matriz de Acesso

### 1. Descrição da História (Visão de Negócio)

Como Gerente Fiscal (Administrador Fiscal),
Quero visualizar e gerenciar a matriz de permissões de cada perfil de acesso,
Para assegurar que as permissões estejam alinhadas com a política de segregação de funções e com os requisitos de controles internos (COSO, Lei das S.A.) [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Matriz de Permissões Imutável por Usuário): As permissões de cada perfil são fixas e definidas pelo Comitê Fiscal e pela Controladoria. Nenhum usuário, incluindo o Administrador Fiscal, pode alterar as permissões de um perfil — apenas consultar a matriz.
* RN05 (Visibilidade da Matriz): A matriz de permissões é visível para os perfis Administrador Fiscal e Auditor/Controller. O perfil Analista Fiscal não tem acesso a esta visualização.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Consulta à matriz de permissões

* Dado que o Administrador Fiscal acessa "Administração de Usuários > Matriz de Permissões";
* Quando o sistema exibe a tabela com 3 perfis (colunas) × 6 ações (linhas);
* Então a matriz deve mostrar ✅ onde o perfil tem permissão e ❌ onde não tem;
* E não deve haver nenhum botão de edição nesta tela — apenas consulta [INDEX].

------------------------------
## 📝 US-03: Segregação de Funções — Bloqueio de Autoaprovação

### 1. Descrição da História (Visão de Negócio)

Como Controller,
Quero ter a garantia de que o sistema impede que um Administrador Fiscal aprove uma alteração de alto impacto que ele mesmo propôs,
Para que o princípio de segregação de funções seja aplicado de forma automática e inviolável, sem depender de controles manuais [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN06 (Bloqueio de Autoaprovação): Quando um Administrador Fiscal tenta aprovar uma alteração de alto impacto em que ele próprio é o proponente, o sistema bloqueia a ação e exige que outro Administrador Fiscal ou Controller realize a aprovação.
* RN07 (Alerta de Administrador Único): Se existir apenas um Administrador Fiscal cadastrado no portal, o sistema exibe um alerta permanente no módulo de administração: "⚠️ Apenas um Administrador Fiscal cadastrado. A segregação de funções exige ao menos dois administradores para o fluxo de aprovação de alto impacto."

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Bloqueio de autoaprovação

* Dado que o Administrador Fiscal "Carlos" propôs uma alteração de alíquota que ultrapassa o patamar de materialidade;
* E a alteração está no status "Pendente de Aprovação";
* Quando o próprio Carlos tentar aprovar a alteração;
* Então o sistema deve exibir: "Você não pode aprovar uma alteração que você mesmo propôs. A aprovação deve ser realizada por outro Administrador Fiscal ou Controller.";
* E a alteração deve permanecer no status "Pendente de Aprovação" [INDEX].

#### Cenário 2: Alerta de administrador único

* Dado que existe apenas um usuário com perfil Administrador Fiscal no portal;
* Quando esse administrador acessar o módulo de administração de usuários;
* Então o sistema deve exibir o alerta "⚠️ Apenas um Administrador Fiscal cadastrado. A segregação de funções exige ao menos dois administradores para o fluxo de aprovação de alto impacto.";
* E o alerta deve permanecer visível até que um segundo Administrador Fiscal seja cadastrado [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
