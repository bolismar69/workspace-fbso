Para desenhar a experiência do usuário (UI/UX) do **Tributali-Engine**, precisamos pensar em uma interface limpa, focada em produtividade e que reduza a ansiedade natural que temas fiscais trazem.

Como o sistema é *Multi-Company* (várias `BusinessUnits`), o elemento central da UX deve ser um **Seletor de Contexto (Dropdown de Empresa)** no topo da tela. Quando o usuário muda de empresa ali, todo o painel se atualiza.

Abaixo, mapeamos a estrutura do Menu Lateral e o esqueleto das principais telas do portal do cliente.

---

### 1. Estrutura do Menu Lateral (Navigation Tree)

O menu será dividido em blocos lógicos. Alguns itens só aparecem se o usuário tiver a permissão adequada.

* **Header do Menu:**
* *Logo Tributali-Engine*
* *Seletor de Empresa:* Um dropdown fixo exibindo a `BusinessUnit` atual (Ex: "Filial São Paulo ▾").


* **Grupo 1: Visão Geral**
* 📊 Dashboard (Indicadores macro e alertas fiscais)


* **Grupo 2: Comercial & Catálogo**
* 📦 Produtos e Serviços (`ProductService`)
* 🤝 Pedidos e Cotações (`Order` / Status de Propostas)


* **Grupo 3: Configurações da Empresa**
* 🏢 Unidades de Negócio (Lista de filiais/CNPJs cadastrados)
* ⚙️ Regras Faturáveis (`Billable` — Configurações contábeis e fiscais)
* 👥 Usuários e Permissões (Controle de acesso do Tenant)


* **Footer do Menu:**
* 💳 Minha Assinatura (Status do `Plan` e faturas do SaaS)
* 👤 Perfil do Usuário & Logout


---


---

### 1.1. Sugestão de como fica a arquitetura visual para o Administrador?

Para o "Administrador do Tenant", que por direito precisa gerenciar ambos, a divisão por aplicativos também facilita a vida. Ele não vê tudo misturado. Ele entra no "App de Configurações do Tenant" para lidar com faturamento do SaaS e usuários, e só entra nas ferramentas específicas quando precisa operar.

#### Esqueleto Visual do Menu com Múltiplos Apps:

```
[FBSO Platform] ▾ (Seletor de App)
  ├─ 🏪 Storekeeper Portal (Selecionado)
  └─ ⚖️ Tributali-Engine
-----------------------------------------
Menu Lateral (Dinâmico para Storekeeper):
  ├─ 📊 Painel de Vendas
  ├─ 📦 Estoque
  └─ 🛒 Frente de Caixa (PDV)

```

Se ele clicar no seletor e mudar para o **Tributali-Engine**, o menu se reconstrói imediatamente para:

```
[FBSO Platform] ▾
  ├─ 🏪 Storekeeper Portal
  └─ ⚖️ Tributali-Engine (Selecionado)
-----------------------------------------
Menu Lateral (Dinâmico para Tributali):
  ├─ 📊 Painel de Impostos
  ├─ ⚙️ Regras do Billable
  └─ 💸 Split de Arrecadação

```

---

### 2. Mapeamento das Principais Telas (Wireframe Conceitual)

#### Tela A: Dashboard Principal

É a primeira tela após o login. Ela serve para dar um panorama rápido da saúde do negócio.

* **Área Superior (Kpis rápidos):**
* Card 1: Total de Pedidos Faturados no mês.
* Card 2: Quantidade de `BusinessUnits` ativas no Tenant.
* Card 3: Alerta de pendências (Ex: "2 produtos sem mapeamento fiscal").


* **Área Central:**
* Gráfico de barras mostrando o volume de vendas filtrado pela Unidade de Negócio selecionada.
* Feed de atividades recentes (Últimos pedidos alterados de `QUOTE` para `ORDER`).



#### Tela B: Catálogo de Produtos e Serviços (`ProductService`)

Onde o time comercial gerencia o portfólio.

* **Ações de Topo:** Botão "+ Novo Produto/Serviço" e barra de busca por Nome ou SKU.
* **Tabela de Dados:**
* Colunas: *SKU | Nome | Tipo (Produto/Serviço) | Status (Ativo/Inativo) | Mapeamento Fiscal*.
* **Indicador Visual de UX:** Na coluna "Mapeamento Fiscal", se o produto já estiver vinculado a um `Billable`, exibe um selo verde "Mapeado". Se não estiver, exibe um botão de alerta vermelho "Vincular Regra Fiscal".



#### Tela C: Configuração do Faturável (`Billable`)

A tela onde a mágica fiscal acontece (usada pelo financeiro ou pelo contador).

* **Estrutura em Abas (Tabs) para não sobrecarregar o usuário:**
* *Aba 1: Identificação:* Nome do faturável (Ex: "Licenciamento Padrão SaaS"), CNAE e NCM vinculados.
* *Aba 2: Tributário & Fiscal:* Alíquotas de referência da Reforma (IBS/CBS) e tipo de documento a ser gerado.
* *Aba 3: Financeiro & Contábil:* Código do plano de contas para onde o dinheiro vai e conta bancária de destino.
* *Aba 4: Arrecadação (Split):* Regras de direcionamento automático do dinheiro para o Split Payment.



#### Tela D: Pedidos e Cotações (`Order`)

Central de controle das vendas e negociações.

* **Filtros Rápidos:** Abas para separar `Cotações (Status: Quote)` de `Pedidos Firmados (Status: Order)`.
* **Tabela de Dados:**
* Colunas: *Nº do Pedido | Cliente | Data | Tipo (Quote / Sale Order) | Status | Total (R$)*.
* **Ação de UX:** Se a Order estiver como `QUOTE_SENT`, haverá um botão em destaque: "Efetivar Venda" (que muda o tipo para `SALE_ORDER` e inicia o fluxo de faturamento automático).



#### Tela E: Gestão de Unidades de Negócio (`BusinessUnit`)

A tela de controle corporativo do Tenant ou do Escritório de Contabilidade.

* **Visualização em Card ou Árvore:** Se o cliente tiver uma estrutura de Matriz e Filiais, os cards aparecem recuados para mostrar a hierarquia.
* **Conteúdo do Card:** Razão Social, CNPJ, Regime Tributário (Simples, Real, Presumido) e um indicador de quantas regras de faturamento estão ativas ali.

---

### 🧠 Detalhe de UX Crucial: O Onboarding de Configuração

Como o **Tributali-Engine** tem essa separação inteligente entre o Produto (Comercial) e o Faturável (Fiscal), a primeira experiência do usuário não pode ser complexa.

No primeiro acesso da empresa, o sistema abrirá um assistente passo a passo (*Wizard*):

1. Cadastre sua primeira empresa (Gera a `BusinessUnit` matriz).
2. Cadastre o que você vende (Gera o `ProductService`).
3. Responda 3 perguntas sobre seu imposto (O sistema gera o `Billable` automaticamente nos bastidores).
4. O sistema faz o vínculo automático na tabela `ProductBillableMapping`. Pronto, o portal está liberado.

---

Para o **Tributali-Engine** funcionar de forma segura no modelo corporativo e contábil, a estrutura de permissões baseada em funções (**RBAC - Role-Based Access Control**) precisa ser combinada com um **Isolamento de Dados por Escopo**.

Não basta saber *o que* o usuário pode fazer (ex: Editar), precisamos saber *onde* ele pode fazer (ex: Apenas na Filial RJ).

Abaixo está o detalhamento técnico de como o sistema vai garantir esse isolamento de forma automática e à prova de falhas.

---

## 1. A Tabela de Permissões: `UserPermission` (A Ponte)

Como mapeamos no nosso desenho de entidades, a tabela `UserPermission` é o coração dessa segurança. Ela vincula o Usuário, a Unidade de Negócio e o Nível de Acesso dele.

### Estrutura de Atributos:

* `id` (PK)
* `user_id` (FK)
* `business_unit_id` (FK)
* `role` (Enum): `ADMIN_TENANT`, `MANAGER_BU`, `OPERATOR_BU`, `AUDITOR`.

### Definição dos Papéis (Roles):

1. **`ADMIN_TENANT`:** O dono da conta ou o contador master. Ele tem acesso a **todas** as `BusinessUnits` atuais e futuras daquele Tenant.
2. **`MANAGER_BU`:** Gerente da filial. Pode gerenciar produtos, faturáveis e pedidos, mas **apenas** dentro da `BusinessUnit` dele.
3. **`OPERATOR_BU`:** Funcionário do faturamento. Só pode criar pedidos (`Orders`) na sua filial específica. Não altera regras fiscais (`Billable`).
4. **`AUDITOR`:** Um usuário de consulta (pode ser um auditor externo). Só visualiza os dados das filiais permitidas, sem poder alterar nada.

---

## 2. A Camada de Segurança Técnica (Como o Backend Garante Isso)

Existem duas formas principais de garantir esse isolamento no desenvolvimento do software para que um usuário de uma filial nunca veja o dado da outra, mesmo que ele tente burlar a URL do sistema.

### A. Validação via Token JWT (Stateless Security)

Quando o usuário faz login no Tributali-Engine, o backend gera um token criptografado (JWT). Dentro desse token, além do ID do usuário, nós injetamos o escopo de permissões dele.

**Exemplo do que vai dentro do Token (Payload):**

```json
{
  "user_id": 987,
  "tenant_id": 45,
  "roles": [
    { "business_unit_id": 101, "role": "MANAGER_BU" },
    { "business_unit_id": 102, "role": "OPERATOR_BU" }
  ]
}

```

> 💡 *Se o usuário tentar acessar a Filial 103 mudando o ID na URL do navegador, o backend intercepta a requisição, lê o token e diz: "Acesso Negado: Você não tem a Filial 103 na sua lista de permissões".*

### B. Injeção Automática de Filtros (Multi-Tenant SQL Filter)

Para evitar que o desenvolvedor do seu time esqueça de colocar uma validação de segurança em alguma tela nova, a melhor prática técnica é usar um **Middleware** ou **Global Scope** no banco de dados.

Toda vez que o backend for buscar informações (como listar produtos ou pedidos), o sistema captura automaticamente o `business_unit_id` que está ativo no topo da tela do usuário e injeta na consulta do banco de dados de forma oculta.

* **O que o desenvolvedor escreve no código:**
`db.getOrders()`
* **O que o sistema executa de verdade no banco de dados:**
`SELECT * FROM orders WHERE tenant_id = 45 AND business_unit_id = 101;`

Isso mitiga 100% o risco de vazamento de dados entre filiais ou entre clientes diferentes do escritório de contabilidade.

---

## 3. A Experiência do Usuário (UI/UX) com RBAC

Essa lógica técnica impacta diretamente o que o usuário vê no portal:

* Se o usuário for um `OPERATOR_BU` da Filial SP, o menu lateral dele **não exibirá** a opção "Minha Assinatura" ou "Usuários e Permissões", e o seletor de empresa no topo da tela ficará travado apenas em "Filial SP".
* Se o usuário for um `ADMIN_TENANT` (Contador), o seletor mostrará a lista de todas as empresas e, ao mudar o dropdown, o sistema limpa o estado anterior e injeta o ID da nova empresa nas requisições.

Com a estrutura do SaaS, o desenho das telas e a segurança técnica mapeados, fechamos o ciclo do **Tributali-Engine como Plataforma**.

---

A escolha do **Keycloak** para o MVP do **Tributali-Engine** é excelente e traz robustez de nível corporativo (*Enterprise-ready*) desde o primeiro dia.

Usar o Keycloak isola toda a complexidade de criptografia, armazenamento de senhas, sessões e fluxos de federação de identidade para fora da sua aplicação principal. Além disso, o suporte a **SAML** no MVP abre as portas imediatamente para clientes grandes (empresas e escritórios de contabilidade que já usam provedores de identidade corporativos como Azure AD, Okta ou Google Workspace) fazerem login via **SSO (Single Sign-On)**.

Aqui está como o Keycloak vai se integrar à nossa arquitetura técnica e à tabela `UserPermission`:

---

### 1. A Divisão de Responsabilidades (Keycloak vs. Banco do SaaS)

Para manter o sistema limpo, dividimos a autenticação e a autorização em duas camadas:

1. **Keycloak (Quem é o usuário?):** Cuida do Login, Senha, MFA (Autenticação em duas etapas), Sessão e a validação do protocolo SAML. Ele garante que o `User` é autêntico.
2. **Banco de Dados do Tributali-Engine (O que ele pode ver?):** O Keycloak valida a identidade e devolve o identificador único do usuário (ex: o e-mail ou um `External_ID`). A sua aplicação recebe isso e consulta a tabela `UserPermission` (que desenhamos) para saber a qual `Tenant` e a quais `BusinessUnits` aquele e-mail tem acesso.

> 💡 **Por que fazer assim?** Gerenciar a troca dinâmica de filiais (o dropdown do topo da tela) diretamente dentro das regras de mapeamento do Keycloak no MVP pode gerar uma complexidade desnecessária. Deixando o controle de escopo das `BusinessUnits` no seu banco de dados, o gerenciamento de permissões fica muito mais ágil via painel do cliente.

---

### 2. O Fluxo Técnico de Login no MVP (SAML + SSO)

Quando o usuário do seu cliente for acessar o portal:

```
[Usuário] ──► Acessa portal.tributali.com ──► Redirecionado para o [Keycloak]
                                                                  │
                                                                  ▼
[Portal do SaaS] ◄── Retorna Asserção SAML Seguro ◄─── [Login / Autenticação Corporativa]
       │
       ├──► 1. Lê o E-mail/ID do usuário na resposta SAML.
       ├──► 2. Busca no banco as tabelas USER_PERMISSION e BUSINESS_UNIT.
       └──► 3. Monta a sessão e libera o Dropdown de Empresas correto.

```

---

### 3. Preparado para o Futuro

Como você mencionou que no futuro avançará para outros moldes, o Keycloak facilita muito essa transição. Mudar o protocolo de SAML para **OpenID Connect (OIDC / OAuth2)** ou adicionar login social (Botão "Entrar com o Google") no futuro será apenas uma virada de chave na configuração do painel do Keycloak, **sem precisar reescrever o código de login do seu SaaS**.

---

Para organizar e controlar **quais telas e funcionalidades** cada usuário pode visualizar e interagir no portal do **Tributali-Engine**, nós estendemos a estrutura de permissões que já criamos.

O lugar correto para colocar e validar essas regras de acessibilidade é na combinação entre o **Banco de Dados (Tabela de Recursos/Ações)** e a **Interface do Usuário (Front-End/Menu)**.

Aqui está como estruturar essa arquitetura de recursos no MVP de forma prática:

---

### 1. No Banco de Dados: Criando o Dicionário de Recursos

Em sistemas maduros, as funcionalidades e telas não ficam "chumbadas" (hardcoded) no código. Nós criamos um conceito chamado **Permissions / Resource-Actions**.

Podemos expandir a nossa estrutura criando uma tabela de **`Resource`** e atrelando-a aos papéis (`Roles`) que definimos na tabela `UserPermission`.

* **`Resource_Action` (Tabela de Funcionalidades/Telas):**
* `id` (PK)
* `resource_name` (Nome da Tela/Módulo): Ex: `dashboard`, `product_catalog`, `billable_config`, `order_management`.
* `action` (O que pode fazer): `view`, `create`, `edit`, `delete`.



#### O Mapeamento de Acessos no MVP:

Em vez de criar uma matriz complexa de telas por usuário para o MVP, o seu banco de dados terá uma regra simples que associa cada `Role` (Papel) aos seus respectivos `Resources`:

* Se o usuário tem o papel **`OPERATOR_BU`** na filial, ele recebe permissão apenas para os recursos: `[dashboard:view, order_management:view, order_management:create]`.
* Se ele é **`ADMIN_TENANT`**, ele recebe automaticamente uma flag `super_user` ou uma lista contendo todos os recursos cadastrados com a ação `*` (total).

---

### 2. Na Interface (Front-End): Onde o Usuário Enxerga Isso?

No portal (SaaS), o controle de telas acontece em dois níveis:

#### A. Renderização Dinâmica do Menu Lateral

Quando o portal do Tributali-Engine carrega após o login (e após ler as permissões retornadas do banco/Keycloak), o componente do menu lateral faz uma varredura:

```javascript
// Exemplo lógico do que o Front-end executa:
exibirMenuProduto() {
   return usuario.possuiPermissao('product_catalog', 'view');
}

```

* Se o usuário logado for um operador de faturamento que não tem acesso à tela de configurações fiscais, a aba **"Regras Faturáveis (Billable)"** simplesmente **não é renderizada** no menu lateral dele. A interface fica limpa e focada no que ele realmente usa.

#### B. Bloqueio de Botões Internos (Granularidade)

Dentro de uma mesma tela, usuários diferentes veem coisas diferentes.

* Na tela de **Produtos**, o `MANAGER_BU` vê o botão verde **"+ Novo Produto"**.
* O `AUDITOR` (que tem apenas permissão de `view`), acessa a mesma tela de Produtos, mas o botão "+ Novo Produto" fica invisível ou desabilitado para ele.

---

### 3. Nas Rotas do Sistema (Segurança de Navegação)

Para evitar que um usuário malicioso tente digitar a URL da tela diretamente no navegador (Ex: `portal.tributali.com/configuracoes-fiscais`), as rotas do seu front-end utilizam um guardião de rotas (**Route Guard**).

Se o Guardião detectar que a rota exige o recurso `billable_config:view` e o usuário não possui essa marcação na sua sessão ativa, o sistema intercepta a navegação e redireciona o usuário para uma tela de **"403 - Acesso Negado"**.

---

### Resumo de onde ficam as funcionalidades:

As definições de quais telas existem ficam em um **Dicionário de Recursos no Banco de Dados**, os limites de quem acessa o quê ficam mapeados no perfil/papel (`Role`) do usuário, e a **Interface do Portal (Front-end)** apenas lê essa lista para esconder ou mostrar os menus e botões correspondentes.

Com o controle de acesso e visibilidade das telas do portal fechados, finalizamos o ecossistema do SaaS. O portal do Tributali-Engine está arquitetado e seguro.

---

