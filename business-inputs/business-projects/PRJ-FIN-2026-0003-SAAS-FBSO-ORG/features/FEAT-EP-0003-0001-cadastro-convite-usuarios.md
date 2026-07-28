# FEATURE - FEAT-EP-0003-0001: Cadastro e Convite de Usuários

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0003-0001 — Cadastro e Convite de Usuários |
| **Épico** | [EP-0003 — Governança de Acessos e Permissões](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) | **Anterior:** [FEAT-EP-0002-0005 — Auditoria](../FEAT-EP-0002-0005-historico-auditoria-administrativa.md) | **Próximo:** [FEAT-EP-0003-0002 — RBAC](../FEAT-EP-0003-0002-definicao-papeis-permissoes-rbac.md)

**Requisitos BRD Vinculados:** [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Gestão de Usuários e Permissões

---

## Objetivo de Negócio
Permitir que o administrador do tenant cadastre e convide usuários para acessar a plataforma, definindo seus acessos de forma granular.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-FEAT-EP-0003-0001-0024 | Como **Administrador do Tenant**, quero convidar um novo usuário para a plataforma informando nome, e-mail e perfil de acesso | • Formulário com campos: nome completo, e-mail • Sistema valida se e-mail já está cadastrado no mesmo tenant • Convite enviado por e-mail com link para definição de senha |
| US-FEAT-EP-0003-0001-0025 | Como **Administrador do Tenant**, quero visualizar a lista de usuários do meu tenant com seus respectivos papéis, unidades vinculadas e status (ativo, inativo, convite pendente) | • Lista exibe: nome, e-mail, papel principal, unidades vinculadas, status • Filtro por status: Todos, Ativos, Pendentes (convite não aceito), Inativos • Indicador visual para convites ainda não aceitos |
| US-FEAT-EP-0003-0001-0026 | Como **Administrador do Tenant**, quero desativar um usuário para bloquear imediatamente seu acesso à plataforma | • Botão "Desativar" na lista de usuários • Confirmação exigida antes da desativação • Usuário desativado não consegue fazer login • Reativação possível a qualquer momento |
| US-FEAT-EP-0003-0001-0059 | Como **Administrador do Tenant**, quero editar os dados básicos de um usuário (nome completo, e-mail) quando necessário para manter o cadastro de usuários do meu tenant sempre atualizado sem depender do suporte da FBSO | • Tela de edição acessível a partir da lista de usuários • Campos editáveis: nome completo, e-mail • Validação de e-mail duplicado no mesmo tenant • Alteração registrada no histórico de auditoria • Se e-mail alterado, sistema notifica o usuário no novo endereço |
| US-FEAT-EP-0003-0001-0060 | Como **Administrador do Tenant**, quero suspender temporariamente um usuário definindo data de início e data de retorno prevista, com um motivo (férias, licença, afastamento) para gerenciar ausências programadas sem precisar desativar e recriar o usuário depois | • Formulário de suspensão temporária com campos: data de início, data prevista de retorno, motivo (Férias, Licença, Afastamento, Outros) • Validação: data de retorno deve ser posterior à data de início • Bloqueio de acesso efetivo a partir da data de início • Reativação automática na data de retorno • Status exibido como 'Suspenso Temporariamente' com motivo e data de retorno na lista de usuários |
| US-FEAT-EP-0003-0001-0061 | Como **Administrador do Tenant**, quero reativar manualmente um usuário que está suspenso temporariamente antes da data prevista de retorno (ex: funcionário retornou antes do previsto) para restaurar o acesso do usuário imediatamente sem precisar esperar a data de reativação automática | • Botão 'Reativar Agora' visível para usuários com status 'Suspenso Temporariamente' • Confirmação exigida antes da reativação • Reativação imediata: status volta para 'Ativo' e acesso restaurado • Cancelamento da reativação automática agendada • Registro de auditoria diferenciando REATIVAR_MANUAL de REATIVAR_AUTOMATICO |

## Regras de Negócio

- **RN-FEAT-EP-0003-0001-0001:** Convite de usuário expira em 7 dias se não aceito
- **RN-FEAT-EP-0003-0001-0002:** E-mail deve ser único por tenant (não pode haver dois usuários com mesmo e-mail no mesmo tenant)
- **RN-FEAT-EP-0003-0001-0003:** Administrador do tenant não pode desativar a si mesmo

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-05** | Gestão de Usuários e Permissões | [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) / J1: Convite de novo usuário · J3: Revogação de acesso | **FEAT-EP-0003-0001** — Cadastro e Convite de Usuários |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
