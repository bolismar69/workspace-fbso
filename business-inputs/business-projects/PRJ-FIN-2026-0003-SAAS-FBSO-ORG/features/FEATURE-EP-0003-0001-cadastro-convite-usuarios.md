# FEATURE - EP-0003-0001: Cadastro e Convite de Usuários

| Campo | Detalhe |
|-------|---------|
| **Feature** | EP-0003-0001 — Cadastro e Convite de Usuários |
| **Épico** | [EP-0003 — Governança de Acessos e Permissões](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) | **Anterior:** [EP-0002-0005 — Auditoria](../FEATURE-EP-0002-0005-historico-auditoria-administrativa.md) | **Próximo:** [EP-0003-0002 — RBAC](../FEATURE-EP-0003-0002-definicao-papeis-permissoes-rbac.md)

**Requisitos BRD Vinculados:** [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Gestão de Usuários e Permissões

---

## Objetivo de Negócio
Permitir que o administrador do tenant cadastre e convide usuários para acessar a plataforma, definindo seus acessos de forma granular.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-024 | Como **Administrador do Tenant**, quero convidar um novo usuário para a plataforma informando nome, e-mail e perfil de acesso | • Formulário com campos: nome completo, e-mail • Sistema valida se e-mail já está cadastrado no mesmo tenant • Convite enviado por e-mail com link para definição de senha |
| US-025 | Como **Administrador do Tenant**, quero visualizar a lista de usuários do meu tenant com seus respectivos papéis, unidades vinculadas e status (ativo, inativo, convite pendente) | • Lista exibe: nome, e-mail, papel principal, unidades vinculadas, status • Filtro por status: Todos, Ativos, Pendentes (convite não aceito), Inativos • Indicador visual para convites ainda não aceitos |
| US-026 | Como **Administrador do Tenant**, quero desativar um usuário para bloquear imediatamente seu acesso à plataforma | • Botão "Desativar" na lista de usuários • Confirmação exigida antes da desativação • Usuário desativado não consegue fazer login • Reativação possível a qualquer momento |

## Regras de Negócio

- **RN-FEAT-EP-0003-0001-0001:** Convite de usuário expira em 7 dias se não aceito
- **RN-FEAT-EP-0003-0001-0002:** E-mail deve ser único por tenant (não pode haver dois usuários com mesmo e-mail no mesmo tenant)
- **RN-FEAT-EP-0003-0001-0003:** Administrador do tenant não pode desativar a si mesmo

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-05** | Gestão de Usuários e Permissões | [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) / J1: Convite de novo usuário · J3: Revogação de acesso | **EP-0003-0001** — Cadastro e Convite de Usuários |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
