# Histórias de Usuário (User Stories) — Feature 01.3
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** M3 — Gestão de Classificações e Regimes (Entrega 1)
* **Feature Relacionada:** [01.3 — Gestão de Classificações e Regimes](./04-FEATURES.md#feature-013-gestão-de-classificações-e-regimes) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Cadastro e Manutenção de Classificações Fiscais (NCM, NBS, CClassTrib, CFOP)

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero cadastrar e manter as classificações fiscais de produtos e serviços (NCM, NBS, CClassTrib, CFOP) em um repositório centralizado,
Para garantir que todas as alíquotas estejam associadas a códigos fiscais válidos e que novas classificações possam ser adicionadas rapidamente quando o portfólio da companhia se expandir [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Tipos de Classificação): O sistema suporta quatro tipos: NCM (8 dígitos), NBS (código de serviços), CClassTrib (código unificado da reforma) e CFOP (4 dígitos).
* RN02 (Código Único por Tipo): Não pode haver duas classificações do mesmo tipo com o mesmo código.
* RN03 (Formato de Código): NCM deve ter 8 dígitos numéricos; CFOP deve ter 4 dígitos numéricos. O sistema rejeita formatos inválidos com mensagem explicativa.
* RN04 (Descritivo Obrigatório): Toda classificação deve ter uma descrição em português claro, que será exibida nos formulários de cadastro de alíquotas e nos relatórios.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cadastro de novo NCM

* Dado que o Analista Fiscal acessa o Módulo 3 — Classificações Fiscais;
* Quando selecionar "Novo NCM", preencher código "2202.10.00" e descrição "Águas minerais e gasosas, incluindo refrigerantes";
* Então o sistema deve confirmar o cadastro e o NCM deve aparecer na lista de classificações disponíveis para associação a alíquotas no Módulo 2 [INDEX].

#### Cenário 2: Bloqueio de código duplicado

* Dado que já existe o NCM "2202.10.00" cadastrado;
* Quando o Analista tentar cadastrar outro NCM com o mesmo código;
* Então o sistema deve exibir: "O código 2202.10.00 já está cadastrado como NCM — Águas minerais e gasosas, incluindo refrigerantes." [INDEX].

#### Cenário 3: Rejeição de formato inválido

* Dado que o Analista está cadastrando um CFOP;
* Quando preencher código "123" (3 dígitos) e tentar salvar;
* Então o sistema deve exibir: "CFOP deve ter exatamente 4 dígitos. Formato esperado: XXXX." [INDEX].

------------------------------
## 📝 US-02: Gestão de Regimes Tributários

### 1. Descrição da História (Visão de Negócio)

Como Administrador Fiscal,
Quero gerenciar os regimes tributários aplicáveis à companhia e aos fornecedores (Lucro Real, Lucro Presumido, Simples Nacional),
Para que o cadastro de alíquotas reflita corretamente as regras de cada regime, incluindo creditamento e transições [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN05 (Regime Padrão): O sistema exige que exatamente um regime tributário esteja marcado como "padrão". Este regime é pré-selecionado nos formulários de cadastro de alíquotas.
* RN06 (Impacto no Creditamento): Cada regime possui um indicador booleano "Gera Crédito" (Lucro Real = Sim; Simples Nacional = Não). Este indicador é exibido como informativo nos dashboards e relatórios.
* RN07 (Proteção contra Desativação): Um regime não pode ser desativado se houver alíquotas vigentes a ele associadas.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Alteração do regime padrão

* Dado que o regime "Lucro Real" está marcado como padrão;
* Quando o Administrador Fiscal acessar "Regimes Tributários", selecionar "Lucro Presumido" e clicar em "Definir como Padrão";
* Então o sistema deve desmarcar "Lucro Real", marcar "Lucro Presumido" e exibir "Regime padrão alterado para Lucro Presumido.";
* E os novos cadastros de alíquotas devem vir com "Lucro Presumido" pré-selecionado [INDEX].

#### Cenário 2: Tentativa de desativar regime com alíquotas vinculadas

* Dado que o regime "Simples Nacional" possui 15 alíquotas vigentes associadas;
* Quando o Administrador tentar desativá-lo;
* Então o sistema deve exibir: "Este regime possui 15 alíquotas vigentes associadas. Reassocie ou desative as alíquotas antes de desativar o regime.";
* E listar as alíquotas vinculadas com links para cada uma [INDEX].

------------------------------
## 📝 US-03: Visualização de Alíquotas Vinculadas a uma Classificação ou Regime

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero visualizar, a partir de uma classificação fiscal ou regime, todas as alíquotas a eles vinculadas,
Para entender rapidamente o impacto de uma alteração ou desativação antes de executá-la [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN08 (Navegação Bidirecional): Da tela de detalhes de uma classificação fiscal, o sistema lista todas as alíquotas associadas. De cada alíquota, há um link para a classificação associada. O mesmo vale para regimes.
* RN09 (Indicador de Impacto): Ao tentar desativar uma classificação ou regime, o sistema exibe o número de alíquotas impactadas e permite filtrar por tipo de impacto (alíquotas que ficarão sem classificação ou sem regime).

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Consulta de alíquotas vinculadas a um NCM

* Dado que o NCM "2202.10.00" está associado a 3 alíquotas: IS (federal), ICMS (SP) e ICMS (MG);
* Quando o Analista acessa a tela de detalhes do NCM "2202.10.00";
* Então o sistema deve listar as 3 alíquotas vinculadas com tributo, valor, UF e vigência;
* E cada alíquota da lista deve ter um link para sua tela de edição no Módulo 2 [INDEX].

#### Cenário 2: Aviso de impacto antes da desativação

* Dado que o NCM "2202.10.00" possui 3 alíquotas vinculadas;
* Quando o Analista clicar em "Desativar" para este NCM;
* Então o sistema deve exibir: "Esta classificação possui 3 alíquotas vigentes associadas. Reassocie ou desative as alíquotas antes de desativar a classificação.";
* E listar as 3 alíquotas com opção de desassociar ou editar cada uma diretamente [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
