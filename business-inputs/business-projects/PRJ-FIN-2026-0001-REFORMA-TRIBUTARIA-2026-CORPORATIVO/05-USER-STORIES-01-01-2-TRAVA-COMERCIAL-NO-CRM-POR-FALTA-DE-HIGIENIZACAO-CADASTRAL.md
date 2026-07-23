# Histórias de Usuário (User Stories) — Feature 01.2
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: CRM / Força de Vendas (Onda 1)
- Feature Relacionada: 01.01.2 — Governança e Trava Comercial de Vendas (CRM) [INDEX]
- Status: Pronto para Dese
nvolvimento Técnico

------------------------------
## 📝 US-01: Bloqueio Mandatório na Geração de Propostas por Inconsistência Geográfica

### 1. Descrição da História (Visão de Negócio)

Como Gerente Comercial da Operação Nacional,
Quero que o sistema impeça a geração ou emissão de propostas comerciais para clientes que não possuam o código IBGE ou o endereço de destino higienizados,
Para que a empresa não assuma compromissos comerciais com alíquotas de IBS incorretas e evite perdas de margem de lucro [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Campos Mandatórios do IVA Dual): Para que um cliente seja considerado "Higienizado", seu registro no CRM deve conter obrigatoriamente preenchidos os campos: CEP, Estado (UF), Município e Código IBGE de Destino (Local de Consumo) [INDEX].
* RN02 (Momento do Bloqueio): A varredura do cadastro do cliente deve ocorrer no instante em que o vendedor clicar na ação de "Avançar", "Salvar Proposta" ou "Gerar PDF do Orçamento".

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cliente com cadastro pendente de higienização geográfica

* Dado que o vendedor está montando uma oportunidade comercial no CRM para um cliente PJ;
* E o cadastro deste cliente possui o campo "Código IBGE de Destino" em branco ou inválido;
* Quando o vendedor clicar no botão "Gerar Proposta Comercial";
* Então o sistema deve bloquear a ação, mantendo o status da proposta como "Rascunho Bloqueado";
* E exibir um aviso impeditivo na tela com a mensagem: "Ação Interrompida: O cadastro deste cliente não possui a qualificação geográfica mandatória para o cálculo do IVA Dual. Atualize o endereço de destino no módulo de Onboarding antes de prosseguir." [INDEX]

#### Cenário 2: Cliente com cadastro 100% higienizado

* Dado que o vendedor está montando uma oportunidade no CRM para um cliente de fora do estado de São Paulo;
* E o cadastro deste cliente possui todos os campos geográficos e o Código IBGE devidamente validados;
* Quando o vendedor clicar no botão "Gerar Proposta Comercial";
* Então o sistema deve liberar a operação, calcular o imposto por fora baseado no destino e avançar a proposta para o status "Emitida" [INDEX].

------------------------------
## 📝 US-02: Alerta Visual Preventivo na Tela da Oportunidade ComerciaI

### 1. Descrição da História (Visão de Negócio)

Como Vendedor Interno (Inside Sales),
Quero visualizar um indicador claro de status fiscal na tela da oportunidade logo no início do atendimento,
Para que eu possa solicitar a atualização cadastral ao cliente antes de gastar tempo estruturando a negociação de preços.

### 2. Regras de Negócio (Business Rules)

* RN01 (Indicador Visual de Risco): O CRM deve exibir uma tag de status visual na barra superior da oportunidade (Ex: Verde para Higienizado / Vermelho para Pendente de Ajuste).
* RN02 (Ação Direta de Correção): O indicador de risco deve conter um link de atalho que direcione o vendedor direto para a tela de saneamento de dados do cliente.

### 3. Critérios de Aceite no Padrão BDD

### Cenário 1: Visualização de conta pendente de higienização

* Dado que o vendedor abre uma oportunidade comercial histórica para realizar uma nova venda;
* Quando a tela do CRM carregar as informações do cliente;
* Então o sistema deve avaliar os dados geográficos e exibir em destaque no topo da tela uma tag vermelha com o texto: ⚠️ Cadastro Inapto para IVA Dual;
* E disponibilizar um botão associado escrito "Atualizar Dados Geográficos do Destino".

------------------------------
## 📝 US-03: Alçada de Exceção e Liberação Temporária por Contingência

### 1. Descrição da História (Visão de Negócio)

Como Diretor Comercial ou Líder do Comitê Fiscal,
Quero possuir uma credencial de liberação especial para aprovar propostas travadas por motivos de força maior ou liminares jurídicas,
Para que operações estratégicas e urgentes da companhia não fiquem paralisadas por falhas sistêmicas externas [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Perfil de Aprovação Máxima): Apenas usuários com perfil de segurança "Diretoria Comercial" ou "Super-User Fiscal" podem liberar a trava cadastral.
* RN02 (Rastreabilidade/Auditoria): Toda liberação por exceção obriga o preenchimento de uma justificativa de negócio em campo de texto livre e gera um log de auditoria permanente no sistema.
* RN03 (Validade do Bypass): A liberação temporária expira em 5 dias úteis, exigindo que o cadastro seja regularizado antes do faturamento final (Onda 2) [INDEX].

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Vendedor solicita liberação de exceção para proposta travada

* Dado que um vendedor teve sua proposta bloqueada pelo sistema devido à falta de código IBGE de um cliente em zona rural;
* Quando o vendedor acionar a opção "Solicitar Liberação por Contingência";
* Então o sistema deve disparar uma notificação interna para o fluxo de aprovação da Diretoria;
* E congelar a oportunidade até a deliberação do gestor.

#### Cenário 2: Diretor aprova a exceção e justifica o bypass

* Dado que o Diretor Comercial recebeu a solicitação de liberação temporária da proposta bloqueada;
* Quando o Diretor clicar em "Aprovar Exceção", preencher a justificativa "Operação sob liminar jurídica de isenção — processo nº X" e confirmar;
* Então o sistema deve destravar temporariamente o botão "Gerar Proposta" para aquela negociação específica;
* E registrar no banco de dados o ID do Diretor, a data da ação e o texto de justificativa para futuras auditorias fiscais do Lucro Real [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
