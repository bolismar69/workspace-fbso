# OBJETIVO DO PRODUTO # SAAS: Nome: FBSO Platform

- *Suíte de Produtos (Multi-Product Suite)
- (SaaS Único / Plataforma "Core" com Módulos)

- O que um sistema para Supermercados, Farmácias ou Lojistas(*ERP/frente de caixa*)  ou TAx-Engine  precisa?

* Cadastro de Clientes e Usuários (Keycloak, que você já definiu).
* Cadastro de Unidades de Negócio (`BusinessUnits` / Filiais).
* Catálogo de Produtos e Serviços (`ProductService`).
* Emissão de Pedidos e Faturas (`Order` e `Invoice`).
* Meios de Pagamento (`Billing_Info` e `Transaction_Payment`).

### A Arquitetura Ideal para a FBSO.ORG: O Modelo "Lego"

A melhor abordagem técnica é criar um **SaaS Base Único (o "Core" da FBSO.ORG)** e tratar os produtos como **Módulos/Licenças** que o cliente pode ativar na assinatura dele.

#### Como isso se reflete no Banco de Dados e no Portal?

No nosso modelo atual, o `PLAN` e a `SUBSCRIPTION` do `Tenant` passam a ditar quais recursos (ou "SaaS") ele contratou.

1. **A tabela `PLAN` ganha uma lista de módulos permitidos:**
* **Plano Tributali Core:** Libera apenas a Engine Fiscal (`Billable`, `Split_Payment`).
* **Plano Commerce Core (Lojistas/Supermercados):** Libera controle de estoque, frente de caixa (PDV) e integrações comerciais, ocultando a complexidade fiscal pesada.
* **Plano Full Suite:** Libera ambos.


2. **No Front-end (Portal do Cliente):**
O menu lateral que desenhamos passa a ser dinâmico com base no plano contratado daquele módulo.
* Se o cliente contratou o *Tributali-Engine*, o menu foca em Regras Faturáveis e Cálculos.
* Se contratou o *FBSO Commerce*, o menu se adapta para exibir Vendas de Balcão, Estoque e Frente de Caixa.

### O Grande Trunfo: Venda Cruzada (Cross-selling)

Ao unificar a base tecnológica sob o guarda-chuva da **FBSO.ORG**, você cria uma máquina de vendas imbatível:

* Se um supermercado contrata o seu SaaS comercial de varejo (`FBSO Commerce`) e, no futuro, precisa se adequar perfeitamente ao Split Payment da Reforma Tributária, você não precisa convencê-lo a migrar de sistema ou integrar APIs complexas.
* O vendedor da FBSO.ORG simplesmente diz: *"Basta clicar aqui, fazer o upgrade do plano, e o módulo Tributali-Engine será ativado na sua conta hoje mesmo, usando os mesmos produtos que você já cadastrou"*.

---


### 1. A Solução de UX: O "Launcher de Aplicativos" (App Switcher)

Em vez de exibir todas as abas de todos os produtos em um único menu lateral gigante, o portal deve ter um **App Switcher** (um seletor de sistema, parecido com o "menu de 9 pontos" do Google onde você muda do Gmail para o Drive).

Ao fazer o login na plataforma da FBSO.ORG, o usuário não entra "direto" em uma tela cheia de abas. Ele entra em uma tela de boas-vindas neutra ou no aplicativo padrão dele.

* **Menu de Navegação:** No canto superior esquerdo (ao lado do logo da FBSO), haverá um seletor.
* **Se o usuário escolher "Storekeeper Portal (Vendas & Loja)":**
O menu lateral muda instantaneamente para focar em *PDV, Estoque, Clientes, Caixa e Relatórios de Venda*. Toda a complexidade tributária pesada (IBS/CBS, regras contábeis do `Billable`) fica **100% oculta**. Ele opera como um lojista puro.
* **Se o usuário mudar para "Tributali-Engine (Fiscal/Reforma)":**
A interface inteira se transforma para exibir o menu de *Mapeamento Tributário, Regras do Billable, Alíquotas e Split de Impostos*. O caixa da loja física some da tela.

---

### 2. A Solução de Segurança: Vinculando Permissões aos "Apps"

Para garantir que um lojista de balcão (frentista de caixa ou repositor de estoque) não clique por engano no Tributali-Engine e mude uma regra fiscal da empresa, nós estendemos a nossa tabela de permissões que desenhamos anteriormente.

Nós adicionamos o conceito de **Licença de Aplicativo** à permissão do usuário:

* **Usuário: "João (Operador de Caixa)"**
* Vinculado ao App: `Storekeeper Portal` (Regra: Operador)
* Vinculado ao App: `Tributali-Engine` ──► **Acesso Bloqueado**
* *Resultado na tela:* Para o João, o botão de alternar para o Tributali-Engine sequer existe. Ele nem sabe que a empresa contratou essa ferramenta.


* **Usuário: "Maria (Contadora/Administradora)"**
* Vinculado ao App: `Storekeeper Portal` (Regra: Visualizador/Auditor)
* Vinculado ao App: `Tributali-Engine` (Regra: Admin)
* *Resultado na tela:* Maria vê o seletor no topo e pode navegar entre os mundos comercial e fiscal com segurança.



---

### 3. Como fica a arquitetura visual para o Administrador?

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

### Conclusão: O Melhor dos Dois Mundos

Adotando essa estratégia de design de produto:

1. **Tecnicamente (Backend & Banco de Dados):** É um SaaS Único. Extremamente barato de manter, fácil de atualizar, sem duplicação de código.
2. **Visualmente (UX / Front-end):** Parecem dois softwares completamente separados e independentes. O usuário do comércio trabalha em paz, e o usuário do fiscal trabalha isolado.

Essa engenharia de UX mata o problema da complexidade e mantém a FBSO.ORG altamente profissional e escalável.



---


## modulo: Tributali-Engine
O objetivo do produto precisa ser explicado em uma única frase impactante para o cliente.

- **Objetivo do SaaS:** > *"Simplificar a transição e a gestão da Reforma Tributária para empresas e contadores, transformando a complexidade fiscal e o split payment em automação financeira inteligente e segura."*

- **Objetivo:** Simplificar a transição e a gestão da Reforma Tributária (IBS/CBS) para empresas e escritórios de contabilidade, transformando a complexidade fiscal e o mecanismo de Split Payment em automação financeira inteligente através de uma estrutura multi-tenant e multi-company.

---

## **FBSO Platform - Modulo: Tributali-Engine como Produto SaaS**.

Para que o portal seja escalável, seguro e ofereça uma excelente experiência para os clientes (empresas e contadores), precisamos desenhar a **Arquitetura de Solução Técnica e os Módulos do Portal**.

---

## 1. Arquitetura da Solução Técnica (The SaaS Stack)

Como o FBSO Platform - Modulo: Tributali-Engine lidará com múltiplos CNPJs por cliente (`BusinessUnits`) e precisa ser extremamente rápido, a arquitetura recomendada segue o modelo **Multi-Tenant com Isolamento Lógico**:

* **Banco de Dados Único (Shared Database):** Uma tabela de banco de dados para todo o sistema, onde **todas** as tabelas comerciais e operacionais possuem a coluna `tenant_id` e `business_unit_id`. Toda consulta ao banco (`SQL Query`) obrigatoriamente filtra por esse ID para garantir que um cliente jamais veja os dados de outro.
* **Autenticação Stateless (JWT):** O usuário faz login, recebe um token criptografado contendo quais `business_unit_ids` ele tem permissão para acessar. O backend valida esse token a cada clique.

---

## 2. Estrutura do Portal (Front-End / Interfaces)

O portal do FBSO Platform - Modulo: Tributali-Engine precisa ser dividido em três visões de sistema (módulos de interface):

### A. Portal Administrativo (Seu Painel de Controle)

Onde o seu time gerencia o SaaS.

* **Dashboard de Métricas:** MRR (Receita Recorrente Mensal), Churn (Cancelamentos) e novos Tenants criados.
* **Gestão de Contas:** Ativação, suspensão de Tenants por inadimplência, e upgrade de planos.
* **Módulo de Configuração de "Orders/Quotes":** Onde seu time de vendas pode gerar um pedido em estado de `QUOTE` para um cliente corporativo.

### B. Portal do Cliente / Dashboard do Usuário

A tela que o cliente final (ou contador) enxerga ao fazer login.

* **Onboarding Automatizado:** Um passo a passo onde o usuário cria o `Tenant` e preenche os dados da primeira `BusinessUnit` (CNPJ, endereço, regime tributário atual).
* **Gestão de Unidades de Negócio:** Tela para cadastrar novas filiais ou novas empresas clientes (se for um contador).
* **Catálogo de Produtos/Serviços:** Tela para a empresa cadastrar o seu portfólio comercial (`ProductService`).
* **Gestão de Usuários e Permissões:** Onde o dono do Tenant convida funcionários e define quem pode ver o quê (Ex: "João só vê a Filial RJ").

---

## 3. O Fluxo de Ativação do Produto (Onboarding Técnico)

Para o SaaS funcionar como uma solução fluida, o fluxo técnico desde o site até o painel deve ser automatizado:

```
[Site Institucional] ──► Escolha do Plano ──► [Checkout / Cadastro]
                                                   │
                                                   ▼
                                        [Criação do Tenant no Banco]
                                                   │
                                                   ▼
                                   [Geração Automática da BusinessUnit 1]
                                                   │
                                                   ▼
                                        [Liberação do Portal]

```

Se o cliente vier por uma venda consultiva (Vendedor), o fluxo começa com a `Order` em estado `QUOTE_SENT`. Quando o cliente aceita, o sistema dispara um e-mail com um link único para ele realizar esse mesmo onboarding acima, já com o plano pré-configurado.

---

## Nem todo cliente do **FBSO Platform - Modulo: Tributali-Engine** vai querer usar o gateway de pagamento integrado do sistema para transacionar as suas vendas. Grandes empresas e escritórios de contabilidade geralmente já possuem seus próprios ERPs (como SAP, Totvs, Omie) ou contas bancárias de relacionamento onde os recebimentos (PIX, Boletos, Cartão) já acontecem por fora.

Para o FBSO Platform - Modulo: Tributali-Engine, o que importa não é *mudar o processo bancário do cliente*, mas sim **capturar o fato gerador (o pagamento recebido)** para poder:

1. Calcular o IBS/CBS de forma precisa com base nos itens do pedido.
2. Alimentar a inteligência da calculadora e os gráficos do painel.
3. Permitir a conciliação contábil e fiscal.

---

### Como o Sistema vai Suportar Isso?

Para que o portal seja essa solução flexível, o fluxo de caixa precisa aceitar duas origens:

#### 1. Origem Automatizada (Via Gateway/Webhook)

O cliente usa o split e as cobranças do Tributali-Engine. O gateway avisa o sistema via webhook, e o sistema grava o `TRANSACTION_PAYMENT` automaticamente como "Pago".

#### 2. Origem Externa / Input Direto (Via API ou Tela)

O cliente processa o pagamento no banco dele e avisa o Tributali-Engine. Para suportar isso, nós criamos duas portas de entrada no SaaS:

* **Via Interface (Manual):** Na tela de detalhes da `Invoice` (Fatura) ou `Order` (Pedido), o usuário clica em um botão **"Registrar Pagamento Manual"**, onde ele digita o valor recebido, a data do pagamento e a forma de pagamento (Ex: "Pago via Transferência Bancária").
* **Via API (Integração):** O ERP do cliente (ex: Totvs) faz uma requisição HTTP POST para a API do Tributali-Engine dizendo: *"O pedido ID 5543 foi pago hoje no valor de R$ 1.000,00"*.

---

### O Ajuste Fino na Modelagem (`TRANSACTION_PAYMENT`)

Para que o banco de dados diferencie o que passou pelo seu motor financeiro do que foi apenas "informado/inputado" pelo cliente, adicionamos um campo de controle na tabela de pagamentos:

* **`TRANSACTION_PAYMENT`**
* `id` (PK)
* `invoice_id` (FK)
* `amount_paid` (Decimal)
* `paid_at` (Datetime)
* **`origin` (Enum): `INTERNAL_GATEWAY` ou `EXTERNAL_INPUT**` *(Crucial para sabermos se o dinheiro foi processado por nós ou apenas registrado pelo cliente)*
* `external_reference_id` (String - Opcional: Para guardar o ID do pagamento que veio do sistema do cliente ou o NSU da maquininha de cartão).



### O Impacto no Split Payment

Essa sua observação traz um alerta técnico muito importante sobre a Reforma Tributária:

* Se a origem for `INTERNAL_GATEWAY`, a Tributali-Engine **executa** o Split Payment bancário (retém o imposto na fonte e envia para o governo).
* Se a origem for `EXTERNAL_INPUT`, o sistema **não consegue reter o dinheiro** (pois o dinheiro caiu direto na conta do cliente por fora). Nesse caso, a Tributali-Engine atua como uma **Engine de Cálculo e Auditoria**: ela registra o débito do imposto, calcula quanto o cliente deve recolher e gera uma guia de arrecadação compensatória (ou o relatório correspondente) para que o financeiro dele pague o governo depois.

Essa flexibilidade transforma o seu SaaS em um sistema agnóstico, capaz de plugar em qualquer empresa, independente do banco ou ERP que ela já utilize.

---

## DELETE LOGICO

Para sistemas corporativos e, principalmente, de auditoria fiscal como o **Tributali-Engine**, a rastreabilidade absoluta dos dados é um requisito obrigatório.

Se um usuário apagar um produto ou um pedido, o contador ou o auditor fiscal precisa saber exatamente quem fez isso, quando fez e quais eram os dados originais.

Analisando o cenário de um SaaS Multi-Tenant de alta performance, a melhor prática de mercado é utilizar a estratégia de **Deleção Lógica na Própria Tabela (Soft Delete)** combinada com um **Histórico de Auditoria (Shadowing/Audit Log)**, descartando a ideia de mover os dados para outra tabela.

Aqui está o porquê e como estruturar isso da forma mais eficiente:

---

### 1. Por que NÃO mover os registros para uma tabela de "Deletados"?

Mover dados entre tabelas em tempo de execução (Ex: tirar de `ORDER` e mandar para `ORDER_DELETED`) quebra a **integridade referencial** do banco de dados (as chaves estrangeiras `FK`).

Se você mover uma `ORDER` para a tabela de excluídos, a tabela `INVOICE` que apontava para ela vai ficar com um ponteiro quebrado (órfão), gerando erros graves no sistema ou obrigando você a mover a cascata inteira de tabelas relacionadas, o que é um pesadelo de manutenção.

---

### 2. A Estratégia Recomendada: Soft Delete + Shadow Columns

A melhor prática consiste em manter o registro na tabela original, mas adicionar uma "marcação de invisibilidade". Para isso, estendemos os seus campos de auditoria padrão.

#### Seus campos padrão em TODAS as tabelas:

* `id` (PK)
* `created_dt` (Datetime)
* `updated_dt` (Datetime)
* `created_by` (FK para o `USER.id` ou ID do Keycloak)
* `updated_by` (FK para o `USER.id` ou ID do Keycloak)

#### Os campos adicionados para Deleção Lógica:

* **`deleted_dt` (Datetime / Nullable):** Se estiver `NULL`, o registro está ativo. Se tiver uma data/hora, significa que foi deletado.
* **`deleted_by` (FK / Nullable):** Guarda o ID do usuário que realizou a exclusão.

> 💡 **Como o sistema se comporta?**
> Toda vez que o portal fizer um `SELECT` para listar produtos ou pedidos na tela, o backend (através do mesmo Middleware/Global Scope que filtra o `Tenant`) injeta automaticamente a regra: `WHERE deleted_dt IS NULL`. Para o usuário comum, o dado sumiu. Para o auditor, o dado continua lá.

---

### 3. Resolvendo o problema dos Índices Únicos (`UNIQUE`)

O maior desafio do Soft Delete no mundo real é quando a tabela possui um campo único. No nosso modelo, a `BUSINESS_UNIT` tem o campo `cnpj` como `UNIQUE`.

Se o usuário cadastrar o CNPJ X, depois der soft-delete nele, o registro continua no banco. Se ele tentar cadastrar o mesmo CNPJ X de novo na semana seguinte, o banco de dados vai dar um erro de "Duplicidade", mesmo o usuário achando que tinha deletado o anterior.

#### Como resolver isso de forma elegante:

Em vez de usar uma coluna booleana (`is_deleted = true/false`), nós usamos a coluna `deleted_dt`. Nas tabelas que exigem unicidade (como o CNPJ), você cria um **Índice Único Composto** que ignora os deletados ou inclui a data de deleção:

```sql
-- Exemplo de melhor prática no PostgreSQL:
CREATE UNIQUE INDEX unique_cnpj_active 
ON business_unit (tenant_id, cnpj) 
WHERE deleted_dt IS NULL;

```

Isso garante que o CNPJ só precisa ser único entre as empresas que estão **ativas**. Se uma empresa for deletada logicamente, o CNPJ dela é liberado para ser cadastrado novamente no sistema sem conflitos, mantendo o histórico da antiga guardado com segurança.

---

