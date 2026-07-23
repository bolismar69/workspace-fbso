# Épicos do Projeto: FBSO Platform — Portal Administrativo SaaS

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | EPICS-FBSO-PLATFORM-2026-001 |
| **Versão** | 1.1 — Revisada conforme Docs Review (15/07/2026) |
| **Data** | 13 de julho de 2026 |
| **Origem** | BRD-FBSO-PLATFORM-2026-001 v1.1 e PC-FBSO-PLATFORM-2026-001 v1.1 |
| **Status** | Aprovado |

---

## Visão Geral dos Épicos

| ID | Épico | Objetivo de Negócio | Func. | Prioridade | Data-Alvo |
|----|-------|---------------------|-------|------------|-----------|
| **EP-01** | Portal Administrativo Interno | Time FBSO.ORG gerencia a operação SaaS com visibilidade em tempo real | 3 | Must Have | **15/08/2026** |
| **EP-02** | Gestão de Clientes e Assinaturas | Estruturar a operação comercial: contas, planos e ciclo de vida do cliente | 5 | Must Have | **31/08/2026** |
| **EP-03** | Governança de Acessos e Permissões | Garantir segurança e isolamento de dados entre clientes e entre filiais | 4 | Must Have | **15/09/2026** |
| **EP-04** | Experiência do Cliente e Autoatendimento | Cliente realiza onboarding, gerencia seus dados e navega entre módulos | 6 | Must Have | **15/10/2026** |

### Cronograma de Épicos

| Data-Alvo | Marco | Épicos | Funcionalidades |
|-----------|-------|--------|----------------|
| 15/08/2026 | M2 | EP-01 — Portal Administrativo Interno | F01-01, F01-02, F01-03 |
| 31/08/2026 | M3 | EP-02 — Gestão de Clientes e Assinaturas | F02-01 a F02-05 |
| 15/09/2026 | M4 | EP-03 — Governança de Acessos e Permissões | F03-01 a F03-04 |
| 30/09/2026 | M5 | EP-04 — Portal do Cliente e Onboarding | F04-01 a F04-04 |
| 15/10/2026 | M6 | EP-04 — Unidades de Negócio e Catálogo | F04-05, F04-06 |
| 30/10/2026 | M7 | Aceite Final — homologação completa | Todas |

### Mapa de Dependências entre Épicos

```
24/07    15/08     31/08     15/09     30/09     15/10     30/10
  │────────│─────────│─────────│─────────│─────────│─────────│
  M1       M2        M3        M4        M5        M6        M7
  ▼        ▼         ▼         ▼         ▼         ▼         ▼
Kickoff  EP-01     EP-02     EP-03     EP-04a    EP-04b   Aceite
         Portal    Clientes  Acessos   Portal    BUs+     Final
         Admin     Planos    Permiss   Cliente   Catálogo
```

> **Nota:** EP-01 é pré-requisito para EP-02 (dependência sequencial conforme milestones M2→M3 do Project Charter). EP-03 depende da existência de Tenants e Usuários. EP-04 depende de toda a camada de governança (EP-03). EP-04 é entregue em duas etapas: Portal do Cliente e Onboarding (M5) e Unidades de Negócio + Catálogo (M6). D6 e D7 podem parcialmente paralelizar com validação de D5.

---

## EP-01: Portal Administrativo Interno

### 1. Nome do Épico
**Portal Administrativo Interno — Painel de Controle FBSO.ORG**

### 2. Objetivo (Goal)

- **Problema:** A FBSO.ORG não possui uma interface centralizada para visualizar e acompanhar a operação do SaaS. Os dados sobre clientes, planos e status das contas estão dispersos ou inexistem, impossibilitando a tomada de decisão operacional ágil e baseada em dados. O time administrativo não tem como saber, em uma única tela, quantos clientes estão ativos, quais planos foram contratados ou se há contas que precisam de atenção.
- **Solução:** Construir um painel de controle (dashboard) administrativo que concentre as principais métricas operacionais do SaaS e permita ao time interno da FBSO.ORG navegar pela base de clientes com filtros e visões consolidadas por período, plano e status.
- **Impacto:** Time interno ganha autonomia e agilidade para acompanhar a saúde da operação; redução do tempo gasto buscando informações em fontes dispersas; base para decisões comerciais e operacionais informadas.

### 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Administrador FBSO.ORG** | Colaborador do time interno responsável pela operação do SaaS | Visão rápida da base de clientes; identificar contas com problemas; métricas de crescimento |
| **Líder Comercial** | Responsável pela estratégia de vendas e relacionamento com clientes | Acompanhar adoção de planos; identificar oportunidades de upgrade; visão de contas por status |
| **Diretoria** | Sócios e diretores da FBSO.ORG | Visão macro da operação; indicadores de crescimento; saúde financeira da base |

### 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Acompanhamento diário da operação**
1. Administrador acessa o portal administrativo
2. Visualiza dashboard com métricas do dia: total de contas ativas, novas contas, contas suspensas
3. Identifica visualmente contas que precisam de atenção (ex: onboarding incompleto, assinatura vencida)
4. Clica em um indicador para ver a lista detalhada de contas naquela situação

**Jornada 2: Análise de crescimento por plano**
1. Líder comercial acessa o dashboard e aplica filtro por período (ex: últimos 30 dias)
2. Visualiza distribuição de clientes por plano (Básico, Core, Full Suite)
3. Identifica plano com maior crescimento no período
4. Visualiza os dados filtrados para compartilhar em apresentação de resultados (exportação de relatórios disponível em fase futura)

### 5. Requisitos de Negócio (Business Requirements)

#### Requisitos Funcionais

- Exibição de indicadores-chave: total de contas ativas, contas novas no período, contas suspensas, distribuição por plano
- Filtro por período para todos os indicadores (dia, semana, mês, trimestre, ano)
- Visão de contas agrupadas por status (ativo, inativo, onboarding pendente, suspenso)
- Lista de contas com busca por nome, status e plano
- Indicadores visuais de alerta para situações críticas (ex: contas suspensas por inadimplência, onboarding incompleto há mais de X dias)
- Gráfico de evolução da base de clientes ao longo do tempo
- Estrutura preparada para receber métricas financeiras futuras (MRR, Churn) quando a comercialização for ativada

#### Requisitos Não-Funcionais

- Dashboard deve carregar em até 3 segundos com volume de dados previsto para o primeiro ano
- Métricas devem refletir dados atualizados (atraso máximo de atualização a definir)
- Acesso restrito a usuários do time interno FBSO.ORG com permissão administrativa

### 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Tempo para obter visão completa da operação (da tela de login à informação desejada) | ≤ 30 segundos |
| Satisfação do time administrativo com o dashboard | Nota ≥ 4,0 / 5,0 |
| Redução de perguntas operacionais para o time técnico | ≥ 50% de redução |

### 7. Fora do Escopo (Out of Scope)

- Métricas financeiras detalhadas (MRR, Churn Rate, LTV) — disponíveis apenas na fase de comercialização
- Dashboards customizáveis por usuário — nesta fase, dashboard único para todos os administradores
- Exportação de relatórios em PDF/Excel — funcionalidade futura
- Previsões ou análises preditivas — funcionalidade futura

### 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Alto** | É a principal ferramenta de trabalho do time interno; sem ela, a operação do SaaS é inviável em escala |

---

## EP-02: Gestão de Clientes e Assinaturas

### 1. Nome do Épico
**Gestão de Clientes e Assinaturas — Ciclo de Vida do Tenant**

### 2. Objetivo (Goal)

- **Problema:** A FBSO.ORG precisa gerenciar o ciclo de vida completo dos seus clientes SaaS — da ativação inicial à suspensão por inadimplência ou upgrade de plano. Sem uma ferramenta centralizada, essas operações seriam manuais, propensas a erros e impossíveis de auditar. Além disso, a oferta comercial (planos e preços) precisa ser configurável pelo time de produto sem dependência de desenvolvimento técnico.
- **Solução:** Criar os módulos de gestão de contas de clientes (Tenants) e de planos/assinaturas, permitindo que o time administrativo ative, suspenda e reative contas, enquanto o time de produto configura os planos comerciais e o time comercial vincula clientes aos planos contratados.
- **Impacto:** Operação comercial estruturada e auditável; redução de erros manuais; autonomia do time de produto para criar e ajustar ofertas; base de dados de clientes consolidada para futura automação de faturamento.

### 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Administrador FBSO.ORG** | Gerencia a base de clientes no dia a dia | Ativar/suspender contas rapidamente; ver histórico de ações; identificar contas problemáticas |
| **Gestor de Produto** | Define e ajusta a oferta comercial | Criar novos planos; alterar preços e módulos incluídos; versionar planos |
| **Líder Comercial** | Acompanha a carteira de clientes | Verificar plano de cada cliente; identificar oportunidades de upgrade; acompanhar renovações |

### 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Ativação de novo cliente (venda consultiva — fase futura)**
1. Vendedor fecha contrato com cliente corporativo e gera uma Ordem de Serviço (processo comercial fora do escopo desta fase; simulado manualmente pelo time administrativo)
2. Administrador acessa o módulo de gestão de contas e cria o Tenant vinculado ao plano contratado
3. Sistema gera link de onboarding e o dispara para o e-mail do cliente
4. Administrador acompanha status: "Aguardando Onboarding" → "Ativo"
5. Caso o cliente não complete o onboarding em X dias, sistema alerta o administrador

**Jornada 2: Suspensão por inadimplência**
1. Administrador identifica cliente com pagamento pendente (via alerta ou busca)
2. Acessa a conta do cliente e altera status para "Suspenso"
3. Sistema bloqueia acesso dos usuários do tenant a partir daquele momento
4. Ação registrada em auditoria com identificação do administrador, data e motivo
5. Futuramente, processo será automatizado via integração com gateway de pagamento (fora do escopo)

**Jornada 3: Upgrade de plano**
1. Cliente solicita upgrade do plano Básico para o Core
2. Comercial acessa a conta, seleciona novo plano e define nova data de vigência
3. Sistema atualiza os módulos disponíveis para o tenant imediatamente
4. No próximo login, o cliente visualiza os novos módulos no App Switcher

### 5. Requisitos de Negócio (Business Requirements)

#### Requisitos Funcionais

- Criação de conta de cliente (Tenant) com dados corporativos: razão social, nome fantasia, segmento
- Ativação, suspensão e reativação de contas com registro de motivo
- Visualização do status de cada conta e tempo em cada status
- Cadastro de planos comerciais: nome, descrição, valor, recorrência (mensal, trimestral, anual)
- Definição de módulos/produtos incluídos em cada plano
- Vinculação de cliente a plano com data de início, vigência e status da assinatura
- Troca de plano (upgrade/downgrade) com registro de data de alteração
- Histórico de todas as ações administrativas: quem fez, o que fez, quando fez
- Estrutura de plano preparada para acoplar faturamento real no futuro (sem processar cobranças)

#### Requisitos Não-Funcionais

- Registro de auditoria imutável para todas as ações administrativas (audit trail)
- Bloqueio de acesso ao portal do cliente em até 5 minutos após suspensão da conta
- Validação de unicidade de razão social por tenant
- Planos inativos (descontinuados) não podem ser vinculados a novas assinaturas

### 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Tempo para ativar uma nova conta (da decisão ao link de onboarding) | ≤ 2 minutos |
| Tempo para suspender conta e bloquear acesso | ≤ 5 minutos |
| Erros em configuração de plano (ex: módulo errado liberado) | Zero não conformidades |
| Cobertura de auditoria | 100% das ações administrativas |

### 7. Fora do Escopo (Out of Scope)

- Processamento de cobranças e faturamento real — estrutura de dados preparada, sem execução financeira
- Renovação automática de assinaturas — será manual nesta fase
- Integração com gateways de pagamento
- Período de trial gratuito automático — pode ser simulado via ativação manual
- Gestão de contratos e documentos legais — funcionalidade futura

### 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Crítico** | Sem gestão de clientes e assinaturas, não há operação SaaS. É o coração do modelo de negócio. |

---

## EP-03: Governança de Acessos e Permissões

### 1. Nome do Épico
**Governança de Acessos e Permissões — RBAC Multi-Tenant e Multi-Unidade**

### 2. Objetivo (Goal)

- **Problema:** Em um SaaS multi-produto com múltiplos clientes (Tenants) e múltiplas filiais (Unidades de Negócio) por cliente, a segurança dos dados é um requisito inegociável. Um operador de caixa de uma filial não pode acessar dados fiscais de outra filial; um auditor externo não pode alterar configurações; e um cliente não pode, sob nenhuma hipótese, visualizar dados de outro cliente. Sem um sistema robusto de permissões, a FBSO Platform estaria exposta a riscos legais, fiscais e de reputação.
- **Solução:** Implementar controle de acesso baseado em papéis (RBAC) que permita ao administrador do tenant definir precisamente o que cada usuário pode ver e fazer — em qual módulo, em qual unidade de negócio e com qual nível de permissão (Admin, Gerente, Operador, Auditor).
- **Impacto:** Segurança e conformidade; isolamento total de dados entre tenants e entre unidades de negócio; flexibilidade para o cliente gerenciar sua própria equipe; base para o modelo de App Switcher com visibilidade condicional por módulo.

### 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Administrador do Tenant** | Dono da conta ou contador master | Convidar usuários; definir quem acessa o quê; revogar acessos; visão completa de todas as unidades |
| **Gerente de Unidade (Manager BU)** | Responsável por uma filial específica | Gerenciar produtos e operações apenas na sua unidade; não pode alterar regras fiscais |
| **Operador de Unidade (Operator BU)** | Funcionário operacional (ex: faturamento, caixa) | Executar tarefas na sua unidade específica; sem acesso a configurações sensíveis |
| **Auditor (fase futura)** | Auditor externo ou interno — previsto para fase posterior ao MVP | Visualizar dados das unidades autorizadas; não pode criar, editar ou excluir nada |

### 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Administrador convida novo usuário para o time**
1. Administrador do tenant acessa "Gestão de Usuários"
2. Clica em "Convidar Usuário" e preenche e-mail, nome e papel (ex: Gerente)
3. Seleciona quais Unidades de Negócio o usuário poderá acessar (ex: Filial SP e Filial RJ)
4. Seleciona quais módulos o usuário pode acessar (ex: apenas Storekeeper Portal)
5. Sistema envia convite por e-mail; usuário define senha no primeiro acesso
6. Ao fazer login, o usuário vê apenas as unidades e módulos autorizados

**Jornada 2: Restrição de acesso entre filiais**
1. João (Operador da Filial SP) faz login no portal
2. O seletor de Unidade de Negócio exibe apenas "Filial SP"
3. João acessa "Catálogo de Produtos" e vê apenas produtos da Filial SP
4. João não vê a opção "Configurações Fiscais" pois não tem permissão
5. João não vê o App Switcher (só tem acesso a um módulo)

**Jornada 3: Revogação de acesso**
1. Administrador remove o acesso de um usuário que saiu da empresa
2. Usuário não consegue mais fazer login (acesso bloqueado imediatamente)
3. Ação registrada em auditoria

### 5. Requisitos de Negócio (Business Requirements)

#### Requisitos Funcionais

- Cadastro de usuários vinculados a um Tenant (convite por e-mail)
- Definição de papéis de acesso (MVP): Admin do Tenant, Gerente de Unidade, Operador. Papel "Auditor" documentado e com schema previsto para fase futura.
- Vinculação de usuário a uma ou mais Unidades de Negócio
- Vinculação de usuário a um ou mais módulos/produtos da plataforma
- Cada papel define um conjunto de permissões: visualizar, criar, editar, excluir por área funcional
- Usuário sem permissão para um módulo não vê o módulo no App Switcher
- Usuário sem permissão para uma funcionalidade não vê o menu ou botão correspondente
- Usuário sem permissão para uma Unidade de Negócio não acessa seus dados
- Administrador do tenant pode revogar acesso a qualquer momento
- Registro de auditoria para criação, alteração e revogação de permissões

#### Requisitos Não-Funcionais

- Verificação de permissão em todas as operações — não basta esconder o menu, o acesso deve ser barrado também se o usuário tentar acessar diretamente
- Bloqueio de acesso em até 5 minutos após revogação
- Senhas e credenciais gerenciadas com segurança (política de complexidade mínima)

### 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Incidentes de acesso não autorizado entre unidades de negócio | Zero |
| Incidentes de acesso não autorizado entre tenants | Zero |
| Tempo para configurar permissões de um novo usuário | ≤ 3 minutos |
| Cobertura de verificação de permissão | 100% das operações |

### 7. Fora do Escopo (Out of Scope)

- Integração com provedores de identidade corporativos (SSO / SAML / Azure AD) — funcionalidade futura
- Autenticação em duas etapas (MFA) — funcionalidade futura
- Permissões customizáveis por cliente (além dos papéis padrão) — funcionalidade futura
- Login social (Google, LinkedIn) — funcionalidade futura
- Delegação temporária de permissões — funcionalidade futura

### 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Crítico** | Segurança de dados é pré-requisito legal e de mercado. Sem RBAC robusto, o SaaS não pode operar com clientes reais. |

---

## EP-04: Experiência do Cliente e Autoatendimento

### 1. Nome do Épico
**Experiência do Cliente e Autoatendimento — Portal, Onboarding e Catálogo**

### 2. Objetivo (Goal)

- **Problema:** Para que a FBSO Platform seja percebida como um produto profissional e escalável, o cliente precisa de uma experiência de entrada (onboarding) fluida e autônoma, um portal onde possa gerenciar seus dados sem depender do time FBSO.ORG, e ferramentas para cadastrar o que é essencial ao seu negócio: suas filiais (Unidades de Negócio) e seu portfólio de produtos. Sem essa camada de autoatendimento, cada novo cliente geraria uma carga manual insustentável para o time interno.
- **Solução:** Construir o portal do cliente com fluxo de onboarding guiado, App Switcher para navegação entre módulos, gestão de unidades de negócio com estrutura hierárquica (Matriz/Filial) e catálogo de produtos/serviços com classificação por tipo e status — tudo adaptado ao plano e às permissões do usuário.
- **Impacto:** Redução drástica do esforço de ativação de clientes; experiência de produto profissional desde o primeiro acesso; base de dados comercial estruturada para futura ativação dos módulos fiscais e de varejo.

### 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Cliente Administrador** | Dono ou responsável pela conta da empresa no SaaS | Fazer onboarding rápido; cadastrar filiais; gerenciar catálogo; convidar usuários do time |
| **Cliente Operador** | Funcionário que usa o portal no dia a dia | Acessar funcionalidades da sua unidade; realizar tarefas operacionais; ver apenas o que lhe compete |
| **Contador (futuro)** | Profissional contábil que gerencia múltiplos clientes — funcionalidade multi-tenant NÃO está no escopo desta fase | Alternar entre clientes (unidades de negócio de diferentes tenants); visão multi-empresa (fase futura) |

### 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Primeiro acesso e onboarding**
1. Cliente recebe e-mail com link de ativação da conta
2. Clica no link, define sua senha e faz login pela primeira vez
3. Sistema identifica que é primeiro acesso e inicia o fluxo de onboarding
4. **Passo 1:** Confirma dados cadastrais da empresa
5. **Passo 2:** Cadastra a primeira Unidade de Negócio (CNPJ matriz, regime tributário)
6. **Passo 3:** Recebe orientação sobre os módulos disponíveis no seu plano
7. **Passo 4:** É direcionado ao portal com uma tela de boas-vindas
8. Portal liberado; menus adaptados ao plano e permissões

**Jornada 2: Cadastro de filiais (Unidades de Negócio)**
1. Administrador do tenant acessa "Unidades de Negócio"
2. Visualiza a matriz cadastrada no onboarding
3. Clica em "+ Nova Unidade" para cadastrar uma filial
4. Preenche CNPJ, endereço, regime tributário e define o vínculo hierárquico (Filial de: Matriz)
5. Sistema valida CNPJ duplicado ativo e exibe confirmação
6. Nova unidade aparece na lista, com indicador visual de hierarquia (recuo para filiais)

**Jornada 3: Cadastro de portfólio de produtos**
1. Gerente de unidade acessa "Catálogo de Produtos e Serviços"
2. Visualiza lista de itens cadastrados para sua unidade
3. Clica em "+ Novo Produto" e preenche: nome, SKU, tipo (Produto ou Serviço), status
4. Sistema salva o produto e o exibe na lista
5. Indicador visual mostra que o produto ainda não possui mapeamento fiscal (preparação para Tributali-Engine)

**Jornada 4: Navegação com App Switcher (visão futura)**
1. Cliente com plano Full Suite faz login
2. App Switcher no topo mostra: "Storekeeper Portal" (selecionado) e "Tributali-Engine"
3. Clica em "Tributali-Engine" e o menu lateral se transforma para mostrar opções fiscais
4. Volta ao Storekeeper e o menu retorna para opções de varejo

### 5. Requisitos de Negócio (Business Requirements)

#### Requisitos Funcionais — Portal do Cliente e Onboarding

- Tela de login e recuperação de senha
- Fluxo de primeiro acesso (onboarding): detecção automática de primeiro login
- Passos do onboarding: confirmação de dados, cadastro da primeira Unidade de Negócio, orientação sobre o plano
- Dashboard do cliente pós-onboarding com visão adaptada ao plano e ao módulo ativo
- Área de perfil do usuário: nome, e-mail, alteração de senha

#### Requisitos Funcionais — App Switcher

- Seletor de módulos visível no topo do portal, ao lado da identidade visual da FBSO Platform
- Exibe apenas os módulos que o usuário tem permissão para acessar
- Ao alternar de módulo, o menu lateral e o conteúdo se adaptam ao módulo selecionado
- Na Fase 0 (este projeto), haverá apenas um módulo visível (placeholder para os futuros)

#### Requisitos Funcionais — Unidades de Negócio

- Cadastro de Unidade de Negócio: CNPJ, razão social, nome fantasia, regime tributário, endereço
- Estrutura hierárquica Matriz/Filial com indicador visual de relacionamento
- Validação de CNPJ duplicado por tenant (CNPJ ativo não pode ser cadastrado duas vezes)
- Listagem de unidades vinculadas ao tenant, agrupadas por hierarquia
- Indicador de status (ativo/inativo) por unidade

#### Requisitos Funcionais — Catálogo de Produtos/Serviços

- Cadastro de item: nome, SKU/código, tipo (Produto, Serviço), descrição, status (ativo/inativo)
- Listagem de itens por Unidade de Negócio com busca por nome ou SKU
- Indicador visual de status: Ativo, Inativo
- Indicador visual "Preparado para mapeamento fiscal" (placeholder para fase futura)
- Edição e ativação/desativação de itens

#### Requisitos Não-Funcionais

- Onboarding completo em até 4 passos; tempo total ≤ 10 minutos
- Portal do cliente responsivo (desktop principal; tablet aceitável)
- Interface em português (Brasil)
- Mensagens de erro em linguagem clara e não-técnica

### 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Clientes que completam onboarding sem ajuda do suporte | ≥ 80% |
| Tempo médio de onboarding completo (login → portal liberado) | ≤ 10 minutos |
| Abandono durante onboarding | ≤ 15% |
| Satisfação com a experiência de primeiro acesso | Nota ≥ 4,0 / 5,0 |
| Unidades de Negócio cadastradas com sucesso na primeira tentativa | ≥ 95% |

### 7. Fora do Escopo (Out of Scope)

- Funcionalidades específicas dos módulos Tributali-Engine e Storekeeper Portal
- Mapeamento fiscal de produtos (NCM, NBS, CNAE, alíquotas IBS/CBS)
- Emissão de pedidos e faturas
- Customização visual do portal por tenant (white label)
- Aplicativo mobile
- Portal em outros idiomas além de português (Brasil)

### 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Crítico** | O portal do cliente é a face visível do SaaS. A experiência de onboarding é o primeiro contato do cliente com o produto. Um onboarding ruim gera abandono e sobrecarrega o suporte. |

---

## Sumário de Cobertura do Escopo

| Entrega do Project Charter | Épico(s) que cobrem |
|---------------------------|---------------------|
| D1 — Portal Administrativo Interno | EP-01 |
| D2 — Módulo de Gestão de Contas | EP-02 |
| D3 — Módulo de Planos e Assinaturas | EP-02 |
| D4 — Módulo de Usuários e Permissões | EP-03 |
| D5 — Portal do Cliente | EP-04 |
| D6 — Cadastro de Unidades de Negócio | EP-04 |
| D7 — Catálogo de Produtos/Serviços | EP-04 |

---

> **Este documento é estritamente de negócio.** Detalhamentos técnicos e decisões de implementação serão tratados em documentos complementares.

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: breakdown-epic-pm, agile-ba-practices. Revisão 1.1 baseada no Docs Review (15/07/2026) — skills: caveman, caveman-review.*
