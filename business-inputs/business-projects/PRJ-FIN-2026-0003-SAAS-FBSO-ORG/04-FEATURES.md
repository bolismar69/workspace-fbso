# Funcionalidades do Projeto: FBSO Platform — Portal Administrativo SaaS

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-FBSO-PLATFORM-2026-001 |
| **Versão** | 1.1 — Revisada conforme Docs Review (15/07/2026) |
| **Data** | 13 de julho de 2026 |
| **Origem** | EPICS-FBSO-PLATFORM-2026-001 v1.1 |
| **Status** | Aprovado |

---

## Visão Geral das Funcionalidades

| ID | Funcionalidade | Épico | Prioridade | User Stories | Data-Alvo |
|----|---------------|-------|------------|-------------|-----------|
| F01-01 | Dashboard de Métricas Operacionais | EP-01 | Must Have | 3 | **15/08/2026** |
| F01-02 | Visão de Contas com Filtros | EP-01 | Must Have | 2 | **15/08/2026** |
| F01-03 | Alertas e Indicadores de Atenção | EP-01 | Should Have | 2 | **15/08/2026** |
| F02-01 | Cadastro e Ativação de Contas de Clientes | EP-02 | Must Have | 4 | **31/08/2026** |
| F02-02 | Gestão de Status do Tenant | EP-02 | Must Have | 3 | **31/08/2026** |
| F02-03 | Configuração de Planos Comerciais | EP-02 | Must Have | 4 | **31/08/2026** |
| F02-04 | Vinculação e Gestão de Assinaturas | EP-02 | Must Have | 3 | **31/08/2026** |
| F02-05 | Histórico de Auditoria Administrativa | EP-02 | Must Have | 2 | **31/08/2026** |
| F03-01 | Cadastro e Convite de Usuários | EP-03 | Must Have | 3 | **15/09/2026** |
| F03-02 | Definição de Papéis e Permissões (RBAC) | EP-03 | Must Have | 4 | **15/09/2026** |
| F03-03 | Vinculação Usuário × Unidade × Módulo | EP-03 | Must Have | 3 | **15/09/2026** |
| F03-04 | Controle de Visibilidade de Menus e Ações | EP-03 | Must Have | 3 | **15/09/2026** |
| F04-01 | Autenticação e Recuperação de Senha | EP-04 | Must Have | 3 | **30/09/2026** |
| F04-02 | Onboarding Guiado de Primeiro Acesso | EP-04 | Must Have | 5 | **30/09/2026** |
| F04-03 | Dashboard do Cliente | EP-04 | Should Have | 2 | **30/09/2026** | ⚠️ D5 (Portal do Cliente) cobre autenticação + onboarding + menu. F04-03 (Dashboard) é bônus se tempo permitir. |
| F04-04 | App Switcher (Seletor de Módulos) | EP-04 | Must Have | 3 | **30/09/2026** |
| F04-05 | Gestão de Unidades de Negócio | EP-04 | Must Have | 5 | **15/10/2026** |
| F04-06 | Catálogo de Produtos e Serviços | EP-04 | Must Have | 4 | **15/10/2026** |

**Total: 18 funcionalidades | 58 user stories**

### Cronograma de Entregas por Funcionalidade

| Data-Alvo | Marco | Épico | Funcionalidades |
|-----------|-------|-------|----------------|
| **15/08/2026** | M2 | EP-01 | F01-01, F01-02, F01-03 |
| **31/08/2026** | M3 | EP-02 | F02-01, F02-02, F02-03, F02-04, F02-05 |
| **15/09/2026** | M4 | EP-03 | F03-01, F03-02, F03-03, F03-04 |
| **30/09/2026** | M5 | EP-04a | F04-01, F04-02, F04-03, F04-04 |
| **15/10/2026** | M6 | EP-04b | F04-05, F04-06 |
| **30/10/2026** | M7 | Todos | Homologação final (D1-D7) |

---

## EP-01: Portal Administrativo Interno

---

### F01-01: Dashboard de Métricas Operacionais

**Objetivo de Negócio:** Prover visibilidade em tempo real da operação do SaaS para o time interno da FBSO.ORG, permitindo acompanhar a saúde da base de clientes e tomar decisões operacionais com agilidade.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-001 | Como **Administrador FBSO**, quero visualizar os indicadores principais da operação em uma tela de dashboard para ter uma visão rápida da saúde do SaaS | • Dashboard carrega com os indicadores atualizados em até 3 segundos • Indicadores exibidos: total de contas ativas, total de contas por status, total de contas por plano • Cada indicador é clicável e leva à lista filtrada correspondente |
| US-002 | Como **Líder Comercial**, quero filtrar as métricas do dashboard por período (últimos 7, 30, 90 dias, mês atual, ano atual) para analisar tendências de crescimento | • Filtro de período disponível no topo do dashboard • Ao alterar o período, todos os indicadores são recalculados • Gráfico de evolução da base reflete o período selecionado |
| US-003 | Como **Diretoria**, quero visualizar um gráfico de evolução da base de clientes ao longo do tempo para acompanhar o crescimento do SaaS | • Gráfico de linhas ou barras exibe a quantidade de novas contas por mês • Gráfico permite alternar entre visão de contas totais e novas contas • Período do gráfico segue o filtro aplicado no dashboard |

#### Regras de Negócio

- **RN01-01:** Métricas consideram apenas tenants com status diferente de "Excluído" (soft delete)
- **RN01-02:** Período padrão do dashboard ao carregar: mês atual
- **RN01-03:** Indicadores que exibem "zero" devem ser apresentados com o número 0, nunca em branco

---

### F01-02: Visão de Contas com Filtros

**Objetivo de Negócio:** Permitir que o time administrativo localize rapidamente qualquer conta de cliente e visualize suas informações principais.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-004 | Como **Administrador FBSO**, quero visualizar a lista completa de contas de clientes com informações resumidas (nome, plano, status, data de criação) para navegar pela base | • Tabela exibe: razão social, plano contratado, status, data de criação, data da última ação • Lista ordenada por data de criação (mais recentes primeiro) • Paginação a cada 25 registros |
| US-005 | Como **Administrador FBSO**, quero buscar uma conta específica por nome ou razão social para localizar rapidamente um cliente | • Campo de busca textual no topo da lista • Busca filtra em tempo real (a partir de 3 caracteres digitados) • Resultados exibem correspondências parciais (ex: "Super" encontra "Supermercado Bom Preço") |

#### Regras de Negócio

- **RN02-01:** Contas com status "Excluído" (soft delete) não aparecem na lista padrão
- **RN02-02:** A busca não diferencia maiúsculas de minúsculas

---

### F01-03: Alertas e Indicadores de Atenção

**Objetivo de Negócio:** Destacar proativamente situações que exigem ação do time administrativo, evitando que problemas passem despercebidos.

**Prioridade:** Should Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-006 | Como **Administrador FBSO**, quero ver indicadores de alerta no dashboard para contas que precisam de atenção (ex: onboarding incompleto há mais de 48h, assinatura suspensa) | • Alertas aparecem como cards coloridos (amarelo: atenção; vermelho: crítico) no topo do dashboard • Cada alerta exibe a quantidade de contas na situação e é clicável • Ao clicar, direciona para a lista filtrada das contas naquela situação |
| US-007 | Como **Administrador FBSO**, quero que o sistema destaque visualmente na lista de contas aquelas com status irregular para identificação rápida durante a navegação | • Contas suspensas exibem ícone ou cor de destaque na lista • Ao passar o cursor, tooltip explica o motivo da suspensão (se registrado) |

#### Regras de Negócio

- **RN03-01:** Alerta de "onboarding incompleto" dispara após 48 horas da criação da conta
- **RN03-02:** Alertas são visíveis para todos os usuários do time interno; não há personalização por usuário nesta fase

---

## EP-02: Gestão de Clientes e Assinaturas

---

### F02-01: Cadastro e Ativação de Contas de Clientes

**Objetivo de Negócio:** Permitir que o time interno crie contas de clientes de forma estruturada, com todos os dados corporativos necessários, gerando automaticamente o convite de acesso ao portal.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-008 | Como **Administrador FBSO**, quero criar uma nova conta de cliente (Tenant) preenchendo razão social, nome fantasia e segmento de mercado para registrar o cliente na plataforma | • Formulário com campos obrigatórios: razão social, segmento • Campos opcionais: nome fantasia, observações • Ao salvar, Tenant é criado com status "Pendente Onboarding" • Sistema gera link único de ativação para o cliente |
| US-009 | Como **Administrador FBSO**, quero que o sistema envie automaticamente um e-mail de boas-vindas ao cliente com o link de ativação da conta após a criação do Tenant | • E-mail disparado automaticamente após criação do Tenant • E-mail contém link único e instruções de primeiro acesso • Link expira em 7 dias (renovável pelo administrador) |
| US-010 | Como **Administrador FBSO**, quero editar os dados cadastrais de um cliente (razão social, nome fantasia, segmento) para manter as informações sempre atualizadas | • Tela de edição acessível a partir da lista de contas • Campos editáveis: razão social, nome fantasia, segmento, observações • Alterações registradas no histórico de auditoria |
| US-011 | Como **Administrador FBSO**, quero reenviar o e-mail de ativação caso o cliente não tenha recebido ou o link tenha expirado | • Botão "Reenviar Convite" na tela de detalhes do Tenant • Disponível apenas para contas com status "Pendente Onboarding" • Gera novo link e novo prazo de 7 dias |

#### Regras de Negócio

- **RN04-01:** Toda criação de Tenant gera registro de auditoria com: administrador responsável, data/hora, dados iniciais
- **RN04-02:** Razão social é validada como obrigatória; sistema alerta se já existir Tenant ativo com mesma razão social
- **RN04-03:** Link de ativação é único e de uso único por Tenant

---

### F02-02: Gestão de Status do Tenant

**Objetivo de Negócio:** Controlar o ciclo de vida de cada conta de cliente, permitindo ativar, suspender e reativar conforme a situação operacional ou comercial.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-012 | Como **Administrador FBSO**, quero alterar o status de uma conta de cliente entre os estados: Pendente Onboarding, Ativo, Suspenso, Inativo para refletir a situação real da conta | • Seletor de status na tela de detalhes do Tenant • Transições permitidas seguem regra de negócio RN05-01 • Ao alterar para Suspenso, campo de motivo é obrigatório |
| US-013 | Como **Administrador FBSO**, quero que ao suspender uma conta, todos os usuários daquele tenant tenham o acesso ao portal bloqueado imediatamente | • Bloqueio efetivo em até 5 minutos após a suspensão • Usuários logados recebem mensagem de sessão encerrada na próxima ação • Status do tenant atualizado em tempo real para "Suspenso" |
| US-014 | Como **Administrador FBSO**, quero visualizar o histórico de mudanças de status de cada conta (quando foi ativada, suspensa, reativada e por quem) | • Linha do tempo de status na tela de detalhes do Tenant • Cada evento exibe: status anterior → novo status, responsável, data/hora, motivo (se aplicável) |

#### Regras de Negócio

- **RN05-01:** Transições de status permitidas:
  - Pendente Onboarding → Ativo (quando cliente completa onboarding)
  - Ativo → Suspenso (inadimplência ou solicitação administrativa)
  - Suspenso → Ativo (reativação)
  - Ativo → Inativo (encerramento de contrato)
  - Inativo → Ativo (recontratação)
- **RN05-02:** Suspensão exige motivo registrado (campo obrigatório)
- **RN05-03:** Reativação de conta suspensa restaura as permissões anteriores dos usuários

---

### F02-03: Configuração de Planos Comerciais

**Objetivo de Negócio:** Permitir que o time de produto crie e gerencie os planos comerciais do SaaS de forma autônoma, definindo preços, módulos incluídos e recorrências.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-015 | Como **Gestor de Produto**, quero cadastrar um novo plano comercial definindo nome, descrição, valor mensal e recorrência disponível (mensal, trimestral, anual) | • Formulário de cadastro com todos os campos • Valores monetários formatados em Real (R$) • Recorrências selecionáveis via checkboxes (pode marcar mais de uma) |
| US-016 | Como **Gestor de Produto**, quero definir quais módulos/produtos da plataforma um plano inclui (ex: Tributali-Engine, Storekeeper Portal) para controlar o que cada cliente pode acessar | • Lista de módulos disponíveis com checkbox ao lado de cada um • Módulos marcados são incluídos no plano • Plano "Full Suite" inclui todos os módulos automaticamente |
| US-017 | Como **Gestor de Produto**, quero editar um plano existente (nome, preço, módulos) mantendo o histórico de versões anteriores | • Edição de plano gera nova versão • Clientes já vinculados permanecem na versão contratada até upgrade • Não é possível excluir plano, apenas desativá-lo |
| US-018 | Como **Administrador FBSO**, quero desativar um plano comercial para que ele não esteja mais disponível para novas contratações, sem afetar clientes que já o possuem | • Plano desativado não aparece como opção em novas assinaturas • Clientes ativos no plano desativado continuam com acesso normal • Plano aparece na lista administrativa com indicador "Descontinuado" |

#### Regras de Negócio

- **RN06-01:** Um plano não pode ser excluído se houver clientes ativos vinculados a ele
- **RN06-02:** Alteração de preço de plano não afeta assinaturas já contratadas (vale o preço da data de contratação)
- **RN06-03:** Deve existir pelo menos um plano ativo no sistema

---

### F02-04: Vinculação e Gestão de Assinaturas

**Objetivo de Negócio:** Vincular clientes a planos com controle de vigência, permitindo upgrade/downgrade e acompanhamento do status da assinatura.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-019 | Como **Administrador FBSO**, quero vincular um cliente a um plano comercial definindo data de início, vigência e status da assinatura | • Seletor de plano (apenas planos ativos listados) • Campos: data de início, data de término (ou "indeterminado"), status (Ativa, Suspensa, Cancelada) • Ao ativar assinatura, módulos do plano são liberados para o tenant |
| US-020 | Como **Líder Comercial**, quero realizar upgrade ou downgrade de plano de um cliente, mantendo o histórico da assinatura anterior | • Ao trocar de plano, assinatura anterior é finalizada com data de término • Nova assinatura é criada com data de início igual ao dia seguinte ao término da anterior • Histórico exibe todas as assinaturas do cliente em ordem cronológica |
| US-021 | Como **Administrador FBSO**, quero suspender a assinatura de um cliente, o que deve bloquear o acesso dele aos módulos do plano | • Suspensão da assinatura bloqueia acesso aos módulos em até 5 minutos • Status da assinatura alterado para "Suspensa" • Reativação da assinatura restaura acesso aos mesmos módulos |

#### Regras de Negócio

- **RN07-01:** Um tenant pode ter apenas uma assinatura ativa por vez
- **RN07-02:** Upgrade/downgrade não pode deixar o tenant sem assinatura ativa durante a transição
- **RN07-03:** Data de término de assinatura é opcional (planos sem data de término = vigência contínua)

---

### F02-05: Histórico de Auditoria Administrativa

**Objetivo de Negócio:** Registrar todas as ações administrativas realizadas pelo time interno, garantindo rastreabilidade e conformidade.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-022 | Como **Administrador FBSO**, quero que toda ação de criação, alteração de status, mudança de plano e edição de dados de tenant seja automaticamente registrada em um histórico de auditoria | • Registro inclui: tipo da ação, administrador responsável, data/hora, dados anteriores e novos (quando aplicável) • Histórico acessível na tela de detalhes do Tenant • Histórico não pode ser editado ou apagado |
| US-023 | Como **Auditor Interno**, quero filtrar o histórico de auditoria por período e por tipo de ação para localizar eventos específicos | • Filtros disponíveis: período (data inicial e final), tipo de ação (criação, suspensão, alteração de plano) • Resultados ordenados do mais recente para o mais antigo |

#### Regras de Negócio

- **RN08-01:** Auditoria cobre 100% das ações administrativas (criação, edição, alteração de status, mudança de plano, alteração de permissões)
- **RN08-02:** Registros de auditoria são imutáveis — não podem ser editados ou excluídos por nenhum usuário

---

## EP-03: Governança de Acessos e Permissões

---

### F03-01: Cadastro e Convite de Usuários

**Objetivo de Negócio:** Permitir que o administrador do tenant cadastre e convide usuários para acessar a plataforma, definindo seus acessos de forma granular.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-024 | Como **Administrador do Tenant**, quero convidar um novo usuário para a plataforma informando nome, e-mail e perfil de acesso | • Formulário com campos: nome completo, e-mail • Sistema valida se e-mail já está cadastrado no mesmo tenant • Convite enviado por e-mail com link para definição de senha |
| US-025 | Como **Administrador do Tenant**, quero visualizar a lista de usuários do meu tenant com seus respectivos papéis, unidades vinculadas e status (ativo, inativo, convite pendente) | • Lista exibe: nome, e-mail, papel principal, unidades vinculadas, status • Filtro por status: Todos, Ativos, Pendentes (convite não aceito), Inativos • Indicador visual para convites ainda não aceitos |
| US-026 | Como **Administrador do Tenant**, quero desativar um usuário para bloquear imediatamente seu acesso à plataforma | • Botão "Desativar" na lista de usuários • Confirmação exigida antes da desativação • Usuário desativado não consegue fazer login • Reativação possível a qualquer momento |

#### Regras de Negócio

- **RN09-01:** Convite de usuário expira em 7 dias se não aceito
- **RN09-02:** E-mail deve ser único por tenant (não pode haver dois usuários com mesmo e-mail no mesmo tenant)
- **RN09-03:** Administrador do tenant não pode desativar a si mesmo

---

### F03-02: Definição de Papéis e Permissões (RBAC)

**Objetivo de Negócio:** Estabelecer os papéis padrão da plataforma com conjuntos de permissões bem definidos, garantindo que cada usuário acesse apenas o que seu perfil permite.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-027 | Como **Administrador do Tenant**, quero atribuir um dos papéis padrão (Admin, Gerente, Operador, Auditor) a cada usuário para definir seu nível de acesso na plataforma | • Seletor de papel no cadastro e na edição do usuário • Descrição de cada papel disponível como tooltip • Alteração de papel registrada em auditoria |
| US-028 | Como **Gestor de Produto**, quero que cada papel tenha um conjunto predefinido de permissões: Admin do Tenant (acesso total), Gerente de Unidade (gerencia sua unidade), Operador (executa tarefas), Auditor (apenas visualiza) | • Permissões mapeadas conforme tabela de papéis (RN10-01) • Permissões não são customizáveis por tenant nesta fase • Papéis são os mesmos para todos os tenants |
| US-029 | Como **Administrador do Tenant**, quero que ao atribuir o papel "Admin do Tenant" a um usuário, ele automaticamente tenha acesso a todas as Unidades de Negócio e todos os módulos do tenant | • Admin do Tenant vê todas as unidades de negócio no seletor • Admin do Tenant vê todos os módulos contratados no App Switcher • Não é necessário configurar permissões individuais para Admin |
| US-030 | Como **Administrador do Tenant**, quero que ao atribuir o papel "Auditor" a um usuário, ele possa visualizar todos os dados das unidades permitidas mas não possa criar, editar ou excluir nada | • Botões de criação/edição/exclusão não visíveis para Auditor • Menus de configuração não aparecem para Auditor • Tentativa de acesso direto a funcionalidades de escrita é bloqueada |

#### Regras de Negócio

- **RN10-01:** Tabela de Permissões por Papel:

| Funcionalidade | Admin Tenant | Gerente BU | Operador BU | Auditor [Fase Futura] |
|---------------|-------------|-----------|------------|---------|
| Dashboard | Ver | Ver | Ver | Ver |
| Unidades de Negócio | Criar, Editar, Ver | Ver (apenas sua) | Ver (apenas sua) | Ver |
| Catálogo de Produtos | Criar, Editar, Ver, Excluir | Criar, Editar, Ver | Ver | Ver |
| Usuários e Permissões | Criar, Editar, Ver, Excluir | — | — | — |
| Planos e Assinaturas | Ver (apenas seu plano) | — | — | Ver |
| Configurações Fiscais* | — | — | — | — |

> *Funcionalidades fiscais não fazem parte do escopo desta fase.

---

### F03-03: Vinculação Usuário × Unidade de Negócio × Módulo

**Objetivo de Negócio:** Permitir controle granular de acesso, definindo exatamente quais unidades de negócio e quais módulos cada usuário pode acessar.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-031 | Como **Administrador do Tenant**, quero definir quais Unidades de Negócio um usuário pode acessar (uma, várias ou todas) para restringir seu escopo de atuação | • No cadastro/edição do usuário, lista de Unidades de Negócio com checkbox • Permite selecionar "Todas" ou unidades específicas • Para Admin do Tenant, "Todas" é fixo e não pode ser alterado |
| US-032 | Como **Administrador do Tenant**, quero definir quais módulos/produtos um usuário pode acessar (ex: apenas Storekeeper, apenas Tributali-Engine, ou ambos) para restringir sua visão da plataforma | • No cadastro/edição do usuário, lista de módulos contratados pelo tenant • Cada módulo com checkbox (marcado = acesso permitido) • Usuário sem acesso a um módulo não o vê no App Switcher |
| US-033 | Como **Administrador do Tenant**, quero alterar as vinculações de um usuário a qualquer momento (adicionar/remover unidade, adicionar/remover módulo) com efeito imediato | • Alterações salvas têm efeito na próxima ação do usuário • Se usuário estiver logado e tiver acesso a uma unidade removida, a sessão é ajustada • Registro de auditoria gerado para cada alteração |

#### Regras de Negócio

- **RN11-01:** Um usuário deve ter pelo menos uma Unidade de Negócio vinculada para acessar o portal (exceto Admin do Tenant, que tem acesso implícito a todas)
- **RN11-02:** Um usuário deve ter pelo menos um módulo vinculado para acessar o portal
- **RN11-03:** A lista de módulos disponíveis para vinculação é determinada pelos módulos incluídos no plano contratado pelo tenant

---

### F03-04: Controle de Visibilidade de Menus e Ações

**Objetivo de Negócio:** Garantir que a interface do portal se adapte dinamicamente às permissões do usuário, ocultando menus, botões e funcionalidades que ele não está autorizado a acessar.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-034 | Como **Usuário do Portal**, quero que o menu lateral exiba apenas as opções correspondentes às minhas permissões (papel + módulo ativo) para ter uma interface limpa e focada no meu trabalho | • Menu lateral renderizado dinamicamente conforme permissões • Itens de menu sem permissão não aparecem (não ficam desabilitados — simplesmente não renderizam) • Ao trocar de módulo no App Switcher, menu se adapta |
| US-035 | Como **Usuário do Portal**, quero que botões de ação (Criar, Editar, Excluir) apareçam apenas se eu tiver permissão para executar aquela ação | • Botão "Novo" visível apenas para quem tem permissão de criação • Botão "Editar" visível apenas para quem tem permissão de edição • Botão "Excluir" visível apenas para quem tem permissão de exclusão |
| US-036 | Como **Usuário do Portal**, quero que ao tentar acessar uma área não permitida diretamente (via URL ou atalho), o sistema me redirecione para uma tela de "Acesso Negado" com explicação amigável | • Tela de acesso negado com mensagem clara e não-técnica • Exibe: "Você não tem permissão para acessar esta área. Se precisar de acesso, contate o administrador da sua conta." • Não exibe detalhes técnicos ou caminhos internos |

#### Regras de Negócio

- **RN12-01:** Ocultação de menu é a primeira camada (UX); o bloqueio por permissão no acesso direto é a camada de segurança — ambas devem ser implementadas
- **RN12-02:** Usuários veem o nome do módulo ativo no topo do portal, ao lado do logo da FBSO Platform

---

## EP-04: Experiência do Cliente e Autoatendimento

---

### F04-01: Autenticação e Recuperação de Senha

**Objetivo de Negócio:** Oferecer ao cliente uma experiência segura e fluida de acesso ao portal, com fluxo de recuperação de senha autônomo.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-037 | Como **Cliente**, quero fazer login no portal usando meu e-mail e senha para acessar minha conta | • Tela de login com campos: e-mail e senha • Mensagens de erro genéricas (não revelar se o e-mail existe ou não) • Após login bem-sucedido, redirecionar para o dashboard ou onboarding (se primeiro acesso) |
| US-038 | Como **Cliente**, quero recuperar minha senha caso eu a esqueça, recebendo um link de redefinição por e-mail | • Opção "Esqueci minha senha" na tela de login • Usuário informa e-mail e recebe link de redefinição • Link expira em 1 hora • Nova senha deve atender critérios mínimos de complexidade |
| US-039 | Como **Cliente**, quero que ao errar a senha repetidas vezes, minha conta seja temporariamente bloqueada por segurança | • Após 5 tentativas consecutivas com erro, conta bloqueada por 15 minutos • Mensagem informa o tempo restante de bloqueio • Administrador do tenant pode desbloquear manualmente |

#### Regras de Negócio

- **RN13-01:** Senha deve ter no mínimo 8 caracteres, incluindo letra e número
- **RN13-02:** Sessão expira após 60 minutos de inatividade
- **RN13-03:** Link de redefinição de senha é de uso único

---

### F04-02: Onboarding Guiado de Primeiro Acesso

**Objetivo de Negócio:** Conduzir o cliente por um fluxo simples e guiado no primeiro acesso, garantindo que ele configure o essencial para começar a usar a plataforma sem precisar de ajuda do suporte.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-040 | Como **Cliente no primeiro acesso**, quero ser recebido por um fluxo guiado de onboarding que me conduza passo a passo pelas configurações iniciais obrigatórias | • Ao detectar primeiro login, sistema inicia automaticamente o onboarding • Barra de progresso visível (Passo 1 de 4, Passo 2 de 4...) • Não é possível pular etapas obrigatórias • Cliente pode salvar e continuar depois (retoma de onde parou) |
| US-041 | Como **Cliente no onboarding**, quero confirmar e complementar meus dados cadastrais (razão social, nome fantasia, segmento) para garantir que as informações estão corretas | • Dados pré-preenchidos com informações fornecidas pelo time FBSO.ORG • Cliente confirma ou edita cada campo • Avançar para o próximo passo salva os dados |
| US-042 | Como **Cliente no onboarding**, quero cadastrar minha primeira Unidade de Negócio (CNPJ matriz, regime tributário, endereço) para começar a operar na plataforma | • Formulário de cadastro de Unidade de Negócio integrado ao fluxo • Campos: CNPJ, razão social, regime tributário (Simples, Lucro Real, Lucro Presumido), endereço • Validação de CNPJ informa se formato é válido (validação de dígitos) |
| US-043 | Como **Cliente no onboarding**, quero visualizar um resumo do meu plano contratado (nome do plano, módulos incluídos, valor) para entender o que está disponível para mim | • Card com informações do plano: nome, módulos incluídos (ícones e nomes), valor mensal • Informação de que novos módulos podem ser contratados futuramente • Botão "Começar a usar" para finalizar o onboarding |
| US-044 | Como **Cliente**, quero ser recebido com uma tela de boas-vindas após concluir o onboarding, com orientações sobre os próximos passos | • Tela de boas-vindas com: saudação personalizada, resumo do que foi configurado • Sugestões de próximos passos: "Convide seu time", "Cadastre seus produtos" • Botão "Ir para o Portal" que leva ao dashboard do cliente |

#### Regras de Negócio

- **RN14-01:** Onboarding é obrigatório no primeiro acesso; não pode ser pulado
- **RN14-02:** Primeira Unidade de Negócio cadastrada no onboarding é automaticamente definida como Matriz
- **RN14-03:** Onboarding só é considerado concluído quando todos os passos obrigatórios são finalizados
- **RN14-04:** Tenant só muda para status "Ativo" após conclusão do onboarding

---

### F04-03: Dashboard do Cliente

**Objetivo de Negócio:** Oferecer ao cliente uma tela inicial com visão geral de sua conta e acessos rápidos às principais funcionalidades.

**Prioridade:** Should Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-045 | Como **Cliente**, quero visualizar um dashboard com informações resumidas da minha conta: unidades de negócio ativas, total de produtos cadastrados e meu plano atual | • Dashboard com cards informativos: Unidades Ativas, Produtos no Catálogo, Plano Contratado • Cada card é clicável e leva à respectiva área de gestão |
| US-046 | Como **Cliente**, quero ver notificações e lembretes relevantes no meu dashboard (ex: "Complete seu cadastro de produtos", "Convite de usuário pendente") | • Área de notificações visível no dashboard • Cada notificação com link para a ação relacionada • Notificações podem ser dispensadas pelo usuário |

#### Regras de Negócio

- **RN15-01:** Dashboard do cliente adapta-se ao módulo ativo no App Switcher (métricas diferentes por módulo)
- **RN15-02:** Na Fase 0 (este projeto), haverá um dashboard genérico que será expandido quando os módulos forem ativados

---

### F04-04: App Switcher (Seletor de Módulos)

**Objetivo de Negócio:** Fornecer o mecanismo de navegação entre módulos da plataforma, permitindo que o cliente alterne entre diferentes produtos sem sair do portal. É a peça central do modelo multi-produto da FBSO Platform.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-047 | Como **Cliente**, quero ver um seletor de aplicativos no topo do portal que exiba os módulos disponíveis no meu plano para navegar entre os produtos que contratei | • App Switcher posicionado no canto superior, ao lado do logo da FBSO Platform • Exibe apenas módulos que o usuário tem permissão para acessar • Módulo ativo aparece destacado visualmente |
| US-048 | Como **Cliente**, quero que ao selecionar um módulo diferente no App Switcher, o menu lateral e o conteúdo da tela se adaptem imediatamente ao módulo escolhido | • Troca de módulo atualiza menu lateral em tempo real • Conteúdo da tela é redirecionado para o dashboard do módulo selecionado • Transição fluida, sem recarregamento completo da página |
| US-049 | Como **Cliente com apenas um módulo contratado**, quero que o Seletor de Módulo exiba o nome do meu módulo mesmo que eu não tenha outras opções, para que eu saiba em qual produto estou | • Seletor de Módulo visível mesmo com um único módulo (na Fase 0, menu de navegação; expande para App Switcher visual quando houver 2+ produtos) • Exibe o nome do módulo ativo sem dropdown de seleção • Indica visualmente que novos módulos podem ser adicionados no futuro |

#### Regras de Negócio

- **RN16-01:** Lista de módulos no App Switcher é determinada pela interseção entre: módulos do plano contratado e módulos que o usuário tem permissão
- **RN16-02:** Na Fase 0, haverá um módulo placeholder chamado "FBSO Platform" visível para todos os clientes (nome documentado no glossário do Project Charter)
- **RN16-03:** A troca de módulo mantém o contexto da Unidade de Negócio selecionada

---

### F04-05: Gestão de Unidades de Negócio

**Objetivo de Negócio:** Permitir que o cliente cadastre e gerencie suas filiais e CNPJs de forma autônoma, com estrutura hierárquica clara (Matriz/Filial).

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-050 | Como **Cliente**, quero visualizar a lista das minhas Unidades de Negócio organizadas hierarquicamente (Matriz no topo, filiais recuadas abaixo) para entender a estrutura da minha empresa | • Visualização em cards ou lista com recuo visual para filiais • Cada card exibe: razão social, CNPJ (mascarado), regime tributário, status • Indicador visual de Matriz vs. Filial |
| US-051 | Como **Cliente**, quero cadastrar uma nova Unidade de Negócio (filial) informando CNPJ, razão social, regime tributário e definindo a qual unidade ela se vincula (Matriz ou outra filial) | • Formulário com campos obrigatórios: CNPJ, razão social, regime tributário • Seletor de unidade pai (Matriz ou filial existente) • Validação de duplicidade de CNPJ ativo para o mesmo tenant |
| US-052 | Como **Cliente**, quero editar os dados de uma Unidade de Negócio (razão social, regime tributário, endereço) para manter as informações atualizadas | • Tela de edição acessível a partir do card da unidade • Campos editáveis: razão social, regime tributário, endereço • CNPJ não pode ser alterado após o cadastro |
| US-053 | Como **Cliente**, quero desativar uma Unidade de Negócio que não está mais em operação, mantendo seu histórico no sistema | • Botão "Desativar" no card da unidade • Confirmação exigida antes da desativação • Unidade desativada não aparece nos seletores para novos cadastros • Dados históricos permanecem acessíveis para consulta |
| US-054 | Como **Cliente**, quero usar o seletor de Unidade de Negócio no topo do portal para alternar entre minhas filiais e visualizar os dados específicos de cada uma | • Seletor dropdown no topo do portal, ao lado do App Switcher • Exibe apenas unidades que o usuário tem permissão para acessar • Ao trocar de unidade, dados exibidos nas telas são filtrados automaticamente |

#### Regras de Negócio

- **RN17-01:** CNPJ deve ser único entre Unidades de Negócio ativas do mesmo tenant (soft delete libera o CNPJ para reúso)
- **RN17-02:** Uma unidade desativada não pode ser definida como "pai" de novas filiais
- **RN17-03:** A primeira unidade cadastrada (durante onboarding) é automaticamente a Matriz
- **RN17-04:** Não há limite de níveis hierárquicos (Matriz → Filial → Sub-filial)
- **RN17-05:** Seletor de Unidade de Negócio reflete apenas as unidades que o usuário tem permissão (Admin vê todas; Gerente e Operador veem apenas as vinculadas)

---

### F04-06: Catálogo de Produtos e Serviços

**Objetivo de Negócio:** Permitir que o cliente cadastre e gerencie seu portfólio comercial (produtos e serviços), preparando a base de dados para o futuro mapeamento fiscal do Tributali-Engine.

**Prioridade:** Must Have

#### User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-055 | Como **Cliente**, quero cadastrar um novo produto ou serviço informando nome, SKU/código interno, tipo (Produto ou Serviço) e descrição | • Formulário com campos: nome (obrigatório), SKU (opcional), tipo (obrigatório: Produto ou Serviço), descrição (opcional) • Item criado com status "Ativo" por padrão • Vinculado automaticamente à Unidade de Negócio ativa no seletor |
| US-056 | Como **Cliente**, quero visualizar a lista de produtos e serviços cadastrados para minha Unidade de Negócio, com busca por nome ou SKU | • Lista filtrada pela Unidade de Negócio selecionada no seletor • Campo de busca textual que filtra por nome ou SKU • Colunas: Nome, SKU, Tipo, Status, Indicador de Mapeamento Fiscal |
| US-057 | Como **Cliente**, quero editar as informações de um produto ou serviço (nome, SKU, tipo, descrição) para manter o catálogo atualizado | • Tela de edição acessível a partir da lista • Todos os campos do cadastro são editáveis • Alterações aplicadas imediatamente após salvar |
| US-058 | Como **Cliente**, quero ativar ou desativar um produto do catálogo sem excluí-lo definitivamente, para controlar quais itens estão em uso | • Botão "Desativar" na lista (para itens ativos) • Botão "Ativar" na lista (para itens inativos) • Item desativado não aparece em cadastros futuros, mas mantém histórico |

#### Regras de Negócio

- **RN18-01:** Catálogo é segmentado por Unidade de Negócio — cada unidade tem seu próprio catálogo
- **RN18-02:** SKU é opcional, mas se informado deve ser único por Unidade de Negócio
- **RN18-03:** Indicador de "Mapeamento Fiscal" nesta fase exibe "Não mapeado" para todos os itens (placeholder para integração futura com Tributali-Engine)
- **RN18-04:** Exclusão de produtos segue política de soft delete (desativação lógica, não remoção física)

---

## Matriz de Cobertura: Entregas do Project Charter × Funcionalidades

| Entrega (Project Charter) | Funcionalidades Relacionadas |
|---------------------------|------------------------------|
| D1 — Portal Administrativo Interno | F01-01, F01-02, F01-03 |
| D2 — Módulo de Gestão de Contas | F02-01, F02-02, F02-05 |
| D3 — Módulo de Planos e Assinaturas | F02-03, F02-04 |
| D4 — Módulo de Usuários e Permissões | F03-01, F03-02, F03-03, F03-04 |
| D5 — Portal do Cliente | F04-01, F04-02, F04-03, F04-04 |
| D6 — Cadastro de Unidades de Negócio | F04-05 |
| D7 — Catálogo de Produtos/Serviços | F04-06 |

```
24/07    15/08     31/08     15/09     30/09     15/10     30/10
  │────────│─────────│─────────│─────────│─────────│─────────│
  M1       M2        M3        M4        M5        M6        M7
  ▼        ▼         ▼         ▼         ▼         ▼         ▼
Kickoff  EP-01     EP-02     EP-03     EP-04a    EP-04b   Aceite
         Portal    Clientes  Acessos   Portal    BUs+     Final
         Admin     Planos    Permiss   Cliente   Catálogo
         F01-01    F02-01    F03-01    F04-01    F04-05
         F01-02    F02-02    F03-02    F04-02    F04-06
         F01-03    F02-03    F03-03    F04-03
                   F02-04    F03-04    F04-04
                   F02-05
```



---

## Matriz de Priorização (MoSCoW)

| Prioridade | Funcionalidades | Quantidade |
|-----------|----------------|------------|
| **Must Have** | F01-01, F01-02, F02-01, F02-02, F02-03, F02-04, F02-05, F03-01, F03-02, F03-03, F03-04, F04-01, F04-02, F04-04, F04-05, F04-06 | 16 |
| **Should Have** | F01-03, F04-03 | 2 |
| **Could Have** | — | 0 |
| **Won't Have (esta fase)** | Funcionalidades dos módulos Tributali-Engine e Storekeeper Portal | — |

---

> **Este documento é estritamente de negócio.** Critérios de aceitação descrevem comportamento esperado do ponto de vista do usuário. Detalhes de implementação técnica serão tratados nos documentos de engenharia (FRS, Especificações Técnicas).

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: agile-ba-practices, acceptance-criteria, breakdown-feature-prd. Revisão 1.1 baseada no Docs Review (15/07/2026) — skills: caveman, caveman-review.*
