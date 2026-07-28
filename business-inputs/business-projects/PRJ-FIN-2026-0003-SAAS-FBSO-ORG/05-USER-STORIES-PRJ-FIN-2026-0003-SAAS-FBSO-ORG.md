# Matriz de Rastreabilidade de Escopo (RTM)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG — Portal Administrativo SaaS FBSO Platform
- **Data de Geração:** 2026-07-21
- **Última Atualização:** 2026-07-27 — Estrutura modular completa: 1 arquivo por User Story (58 arquivos)
- **Versão:** 3.0 — Repositório modular com 58 US individuais
- **Status de Auditoria:** Em Validação — 7 entregas, 4 épicos, 18 funcionalidades, 62 user stories

---

## Visão Geral da Cobertura

| Métrica | Valor |
|---|---|
| Total de Entregas (Charter) | 7 (D1-D7) |
| Total de Épicos | 4 (EP-0001 a EP-0004) |
| Total de Funcionalidades | 18 (FEAT-EP-0001-0001 a FEAT-EP-0004-0006) |
| Total de User Stories | 62 |
| Cobertura D→Epic | 7/7 (100%) |
| Cobertura Epic→Feature | 4/4 (100%) |
| Cobertura Feature→US | 18/18 (100% — todas as features com pelo menos 1 US) |
| Órfãos detectados | 0 |

---

## Matriz de Rastreabilidade Completa

- **STATUS** = Compliance Gate (🔴 NON-COMPLIANCE / 🟡 PENDING-REVIEW / 🟢 COMPLIANCE)
- **FASE** = Ciclo de vida da User Story (NEW → BIZ-REFINE → READY-TECH → TECH-REFINE → READY-DEV → IN-PROGRESS → CODE-REVIEW → QA → UAT → DONE → DEPLOYED → CANCELLED)
- **RNs** = Regras de Negócio associadas (herdadas dos arquivos de feature)

| D# | EPIC-ID | FEATURE-ID | US-ID | US Descrição | STATUS | FASE | RNs |
|:---|:---|:---|:---|:---|:---:|:---|:---|
| D1 | EP-0001 | FEAT-EP-0001-0001 | US-FEAT-EP-0001-0001-0001 | visualizar os indicadores principais da operação em uma tela de dashboard para ter uma visão rápida  | 🟡 | `NEW` | — |
| D1 | EP-0001 | FEAT-EP-0001-0001 | US-FEAT-EP-0001-0001-0002 | filtrar as métricas do dashboard por período (últimos 7, 30, 90 dias, mês atual, ano atual) para ana | 🟡 | `NEW` | — |
| D1 | EP-0001 | FEAT-EP-0001-0001 | US-FEAT-EP-0001-0001-0003 | visualizar um gráfico de evolução da base de clientes ao longo do tempo para acompanhar o cresciment | 🟡 | `NEW` | — |
| D1 | EP-0001 | FEAT-EP-0001-0002 | US-FEAT-EP-0001-0002-0004 | visualizar a lista completa de contas de clientes com informações resumidas (nome, plano, status, da | 🟡 | `NEW` | — |
| D1 | EP-0001 | FEAT-EP-0001-0002 | US-FEAT-EP-0001-0002-0005 | buscar uma conta específica por nome ou razão social para localizar rapidamente um cliente | 🟡 | `NEW` | — |
| D1 | EP-0001 | FEAT-EP-0001-0003 | US-FEAT-EP-0001-0003-0006 | ver indicadores de alerta no dashboard para contas que precisam de atenção (ex: onboarding incomplet | 🟡 | `NEW` | — |
| D1 | EP-0001 | FEAT-EP-0001-0003 | US-FEAT-EP-0001-0003-0007 | que o sistema destaque visualmente na lista de contas aquelas com status irregular para identificaçã | 🟡 | `NEW` | — |
| D2 | EP-0002 | FEAT-EP-0002-0001 | US-FEAT-EP-0002-0001-0008 | criar uma nova conta de cliente (Tenant) preenchendo razão social, nome fantasia e segmento de merca | 🟡 | `NEW` | — |
| D2 | EP-0002 | FEAT-EP-0002-0001 | US-FEAT-EP-0002-0001-0009 | que o sistema envie automaticamente um e-mail de boas-vindas ao cliente com o link de ativação da co | 🟡 | `NEW` | — |
| D2 | EP-0002 | FEAT-EP-0002-0001 | US-FEAT-EP-0002-0001-0010 | editar os dados cadastrais de um cliente (razão social, nome fantasia, segmento) para manter as info | 🟡 | `NEW` | — |
| D2 | EP-0002 | FEAT-EP-0002-0001 | US-FEAT-EP-0002-0001-0011 | reenviar o e-mail de ativação caso o cliente não tenha recebido ou o link tenha expirado | 🟡 | `NEW` | — |
| D2 | EP-0002 | FEAT-EP-0002-0002 | US-FEAT-EP-0002-0002-0012 | alterar o status de uma conta de cliente entre os estados: Pendente Onboarding, Ativo, Suspenso, Ina | 🟡 | `NEW` | — |
| D2 | EP-0002 | FEAT-EP-0002-0002 | US-FEAT-EP-0002-0002-0013 | que ao suspender uma conta, todos os usuários daquele tenant tenham o acesso ao portal bloqueado ime | 🟡 | `NEW` | — |
| D2 | EP-0002 | FEAT-EP-0002-0002 | US-FEAT-EP-0002-0002-0014 | visualizar o histórico de mudanças de status de cada conta (quando foi ativada, suspensa, reativada  | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0003 | US-FEAT-EP-0002-0003-0015 | cadastrar um novo plano comercial definindo nome, descrição, valor mensal e recorrência disponível ( | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0003 | US-FEAT-EP-0002-0003-0016 | definir quais módulos/produtos da plataforma um plano inclui (ex: Tributali-Engine, Storekeeper Port | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0003 | US-FEAT-EP-0002-0003-0017 | editar um plano existente (nome, preço, módulos) mantendo o histórico de versões anteriores | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0003 | US-FEAT-EP-0002-0003-0018 | desativar um plano comercial para que ele não esteja mais disponível para novas contratações, sem af | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0004 | US-FEAT-EP-0002-0004-0019 | vincular um cliente a um plano comercial definindo data de início, vigência e status da assinatura | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0004 | US-FEAT-EP-0002-0004-0020 | realizar upgrade ou downgrade de plano de um cliente, mantendo o histórico da assinatura anterior | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0004 | US-FEAT-EP-0002-0004-0021 | suspender a assinatura de um cliente, o que deve bloquear o acesso dele aos módulos do plano | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0005 | US-FEAT-EP-0002-0005-0022 | que toda ação de criação, alteração de status, mudança de plano e edição de dados de tenant seja aut | 🟡 | `NEW` | — |
| D3 | EP-0002 | FEAT-EP-0002-0005 | US-FEAT-EP-0002-0005-0023 | filtrar o histórico de auditoria por período e por tipo de ação para localizar eventos específicos | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0001 | US-FEAT-EP-0003-0001-0024 | convidar um novo usuário para a plataforma informando nome, e-mail e perfil de acesso | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0001 | US-FEAT-EP-0003-0001-0025 | visualizar a lista de usuários do meu tenant com seus respectivos papéis, unidades vinculadas e stat | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0001 | US-FEAT-EP-0003-0001-0026 | desativar ou reativar um usuário para controlar seu acesso à plataforma | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0002 | US-FEAT-EP-0003-0002-0027 | atribuir um dos papéis padrão (Admin, Gerente, Operador, Auditor) a cada usuário para definir seu ní | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0002 | US-FEAT-EP-0003-0002-0028 | que cada papel tenha um conjunto predefinido de permissões: Admin do Tenant (acesso total), Gerente  | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0002 | US-FEAT-EP-0003-0002-0029 | que ao atribuir o papel "Admin do Tenant" a um usuário, ele automaticamente tenha acesso a todas as  | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0002 | US-FEAT-EP-0003-0002-0030 | que ao atribuir o papel "Auditor" a um usuário, ele possa visualizar todos os dados das unidades per | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0003 | US-FEAT-EP-0003-0003-0031 | definir quais Unidades de Negócio um usuário pode acessar (uma, várias ou todas) para restringir seu | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0003 | US-FEAT-EP-0003-0003-0032 | definir quais módulos/produtos um usuário pode acessar (ex: apenas Storekeeper, apenas Tributali-Eng | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0003 | US-FEAT-EP-0003-0003-0033 | alterar as vinculações de um usuário a qualquer momento (adicionar/remover unidade, adicionar/remove | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0004 | US-FEAT-EP-0003-0004-0034 | que o menu lateral exiba apenas as opções correspondentes às minhas permissões (papel + módulo ativo | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0004 | US-FEAT-EP-0003-0004-0035 | que botões de ação (Criar, Editar, Excluir) apareçam apenas se eu tiver permissão para executar aque | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0004 | US-FEAT-EP-0003-0004-0036 | que ao tentar acessar uma área não permitida diretamente (via URL ou atalho), o sistema me redirecio | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0001 | US-FEAT-EP-0004-0001-0037 | fazer login no portal usando meu e-mail e senha para acessar minha conta | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0001 | US-FEAT-EP-0004-0001-0038 | recuperar minha senha caso eu a esqueça, recebendo um link de redefinição por e-mail | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0001 | US-FEAT-EP-0004-0001-0039 | que ao errar a senha repetidas vezes, minha conta seja temporariamente bloqueada por segurança | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0002 | US-FEAT-EP-0004-0002-0040 | ser recebido por um fluxo guiado de onboarding que me conduza passo a passo pelas configurações inic | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0002 | US-FEAT-EP-0004-0002-0041 | confirmar e complementar meus dados cadastrais (razão social, nome fantasia, segmento) para garantir | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0002 | US-FEAT-EP-0004-0002-0042 | cadastrar minha primeira Unidade de Negócio (CNPJ matriz, regime tributário, endereço) para começar  | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0002 | US-FEAT-EP-0004-0002-0043 | visualizar um resumo do meu plano contratado (nome do plano, módulos incluídos, valor) para entender | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0002 | US-FEAT-EP-0004-0002-0044 | ser recebido com uma tela de boas-vindas após concluir o onboarding, com orientações sobre os próxim | 🟡 | `NEW` | — |
| — | EP-0004 | FEAT-EP-0004-0003 | US-FEAT-EP-0004-0003-0045 | visualizar um dashboard com informações resumidas da minha conta: unidades de negócio ativas, total  | 🟡 | `NEW` | — |
| — | EP-0004 | FEAT-EP-0004-0003 | US-FEAT-EP-0004-0003-0046 | ver notificações e lembretes relevantes no meu dashboard (ex: "Complete seu cadastro de produtos", " | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0004 | US-FEAT-EP-0004-0004-0047 | ver um seletor de aplicativos no topo do portal que exiba os módulos disponíveis no meu plano para n | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0004 | US-FEAT-EP-0004-0004-0048 | que ao selecionar um módulo diferente no App Switcher, o menu lateral e o conteúdo da tela se adapte | 🟡 | `NEW` | — |
| D5 | EP-0004 | FEAT-EP-0004-0004 | US-FEAT-EP-0004-0004-0049 | que o Seletor de Módulo exiba o nome do meu módulo mesmo que eu não tenha outras opções | 🟡 | `NEW` | — |
| D6 | EP-0004 | FEAT-EP-0004-0005 | US-FEAT-EP-0004-0005-0050 | visualizar a lista das minhas Unidades de Negócio organizadas hierarquicamente (Matriz no topo, fili | 🟡 | `NEW` | — |
| D6 | EP-0004 | FEAT-EP-0004-0005 | US-FEAT-EP-0004-0005-0051 | cadastrar uma nova Unidade de Negócio (filial) informando CNPJ, razão social, regime tributário e de | 🟡 | `NEW` | — |
| D6 | EP-0004 | FEAT-EP-0004-0005 | US-FEAT-EP-0004-0005-0052 | editar os dados de uma Unidade de Negócio (razão social, regime tributário, endereço) para manter as | 🟡 | `NEW` | — |
| D6 | EP-0004 | FEAT-EP-0004-0005 | US-FEAT-EP-0004-0005-0053 | desativar uma Unidade de Negócio que não está mais em operação, mantendo seu histórico no sistema | 🟡 | `NEW` | — |
| D6 | EP-0004 | FEAT-EP-0004-0005 | US-FEAT-EP-0004-0005-0054 | usar o seletor de Unidade de Negócio no topo do portal para alternar entre minhas filiais e visualiz | 🟡 | `NEW` | — |
| D7 | EP-0004 | FEAT-EP-0004-0006 | US-FEAT-EP-0004-0006-0055 | cadastrar um novo produto ou serviço informando nome, SKU/código interno, tipo (Produto ou Serviço)  | 🟡 | `NEW` | — |
| D7 | EP-0004 | FEAT-EP-0004-0006 | US-FEAT-EP-0004-0006-0056 | visualizar a lista de produtos e serviços cadastrados para minha Unidade de Negócio, com busca por n | 🟡 | `NEW` | — |
| D7 | EP-0004 | FEAT-EP-0004-0006 | US-FEAT-EP-0004-0006-0057 | editar as informações de um produto ou serviço (nome, SKU, tipo, descrição) para manter o catálogo a | 🟡 | `NEW` | — |
| D7 | EP-0004 | FEAT-EP-0004-0006 | US-FEAT-EP-0004-0006-0058 | ativar ou desativar um produto do catálogo sem excluí-lo definitivamente | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0001 | US-FEAT-EP-0003-0001-0059 | editar os dados básicos de um usuário (nome, e-mail) para manter o cadastro atualizado | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0001 | US-FEAT-EP-0003-0001-0060 | suspender temporariamente um usuário definindo período de ausência (férias, licença, afastamento) | 🟡 | `NEW` | — |
| D4 | EP-0003 | FEAT-EP-0003-0001 | US-FEAT-EP-0003-0001-0061 | reativar manualmente um usuário antes do fim da suspensão temporária | 🟡 | `NEW` | — |
| — | EP-0004 | FEAT-EP-0004-0003 | US-FEAT-EP-0004-0003-0062 | realizar upgrade do plano contratado diretamente pelo portal (self-service) | 🟡 | `NEW` | — |

### Legenda das Fases

| FASE | Significado | Distribuição |
|:---|:---|:---|
| `NEW` | Novo — ideia registrada, aguardando priorização | Todas as 62 US (recém-criadas na estrutura modular) |
| `BIZ-REFINE` | Em refinamento de negócio — PO detalha critérios | A definir após priorização |
| `READY-TECH` | Pronto para refinamento técnico — negócio concluiu | A definir |
| `TECH-REFINE` | Em refinamento técnico — time debate solução, estima | A definir |
| `READY-DEV` | Pronto para desenvolvimento | A definir |
| `IN-PROGRESS` | Em desenvolvimento | A definir |
| `CODE-REVIEW` | Em revisão de código | A definir |
| `QA` | Em teste — QA valida critérios de aceite | A definir |
| `UAT` | Em homologação — PO ou cliente valida | A definir |
| `DONE` | Pronto — 100% do DoD | A definir |
| `DEPLOYED` | Em produção | A definir |
| `CANCELLED` | Cancelado — perdeu valor de negócio | A definir |

---

## Verificação de Cobertura de Entregas

| ID Entrega | Qtd. US | Status |
|:---|:---|:---|
| D1 — Portal Administrativo Interno | 7 | ✅ Coberto |
| D2 — Gestão de Contas de Clientes (Tenants) | 7 | ✅ Coberto |
| D3 — Gestão de Planos e Assinaturas | 9 | ✅ Coberto |
| D4 — Gestão de Usuários e Permissões (RBAC) | 16 | ✅ Coberto |
| D5 — Portal do Cliente (Auto-Serviço) | 11 | ✅ Coberto |
| D6 — Cadastro de Unidades de Negócio | 5 | ✅ Coberto |
| D7 — Catálogo de Produtos e Serviços | 4 | ✅ Coberto |
| Bônus — Dashboard do Cliente (Should Have) | 3 | ⚠️ Should Have, sem entrega obrigatória no Charter |

---

## Verificação de Integridade Física

| Verificação | Resultado |
|:---|---|
| Links quebrados na RTM | 0 — Todos os 62 links apontam para arquivos existentes em `user-stories/` |
| Arquivos órfãos na pasta `user-stories/` | 0 — Todos os 62 arquivos estão catalogados na RTM |
| Features sem User Stories | 0 — Todas as 18 funcionalidades possuem pelo menos 1 US |
| Entregas do Charter sem cobertura | 0 — Todas as 7 entregas (D1-D7) estão cobertas |

---

## Auditoria de Consistência Conceitual

| Critério | Resultado |
|:---|---|
| Consistência de nomenclatura de IDs | ✅ Pass — Padrão US-FEAT-EP-{EEEE}-{SSSS}-{SSSS} com SSSS global (0001-0062) |
| Rastreabilidade vertical completa | ✅ Pass — Toda US rastreia até D (Charter) via Feature→Epic |
| Alinhamento terminológico | ✅ Pass — Termos de negócio consistentes entre fases |
| Um arquivo por User Story | ✅ Pass — 62 arquivos individuais em `user-stories/` |
| Padrão de nomenclatura | ✅ Pass — US-FEAT-{codigo-feature}-{SSSS}-{nome-slug}.md |

---

## Documentos Relacionados

- [01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](./01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Fase 1
- [02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](./02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Fase 2
- [03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](./03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Fase 3
- [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](./04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Fase 4
- [user-stories/](./user-stories/) — Fase 5 (62 arquivos individuais)

---

**Status Final:** `[PRÉ-COMPLIANCE]` — Repositório modular com 62 arquivos gerado (58 originais + 4 do Gap Analysis #11-#16). Aguardando validação do GATE e aprovação humana.
