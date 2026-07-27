# FEATURE - FEAT-EP-0003-0003: Vinculação Usuário × Unidade de Negócio × Módulo

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0003-0003 — Vinculação Usuário × Unidade de Negócio × Módulo |
| **Épico** | [EP-0003 — Governança de Acessos e Permissões](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) | **Anterior:** [FEAT-EP-0003-0002 — RBAC](../FEAT-EP-0003-0002-definicao-papeis-permissoes-rbac.md) | **Próximo:** [FEAT-EP-0003-0004 — Controle de Visibilidade](../FEAT-EP-0003-0004-controle-visibilidade-menus-acoes.md)

**Requisitos BRD Vinculados:** [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Gestão de Usuários e Permissões

---

## Objetivo de Negócio
Permitir controle granular de acesso, definindo exatamente quais unidades de negócio e quais módulos cada usuário pode acessar.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-031 | Como **Administrador do Tenant**, quero definir quais Unidades de Negócio um usuário pode acessar (uma, várias ou todas) para restringir seu escopo de atuação | • No cadastro/edição do usuário, lista de Unidades de Negócio com checkbox • Permite selecionar "Todas" ou unidades específicas • Para Admin do Tenant, "Todas" é fixo e não pode ser alterado |
| US-032 | Como **Administrador do Tenant**, quero definir quais módulos/produtos um usuário pode acessar (ex: apenas Storekeeper, apenas Tributali-Engine, ou ambos) para restringir sua visão da plataforma | • No cadastro/edição do usuário, lista de módulos contratados pelo tenant • Cada módulo com checkbox (marcado = acesso permitido) • Usuário sem acesso a um módulo não o vê no App Switcher |
| US-033 | Como **Administrador do Tenant**, quero alterar as vinculações de um usuário a qualquer momento (adicionar/remover unidade, adicionar/remover módulo) com efeito imediato | • Alterações salvas têm efeito na próxima ação do usuário • Se usuário estiver logado e tiver acesso a uma unidade removida, a sessão é ajustada • Registro de auditoria gerado para cada alteração |

## Regras de Negócio

- **RN-FEAT-EP-0003-0003-0001:** Um usuário deve ter pelo menos uma Unidade de Negócio vinculada para acessar o portal (exceto Admin do Tenant, que tem acesso implícito a todas)
- **RN-FEAT-EP-0003-0003-0002:** Um usuário deve ter pelo menos um módulo vinculado para acessar o portal
- **RN-FEAT-EP-0003-0003-0003:** A lista de módulos disponíveis para vinculação é determinada pelos módulos incluídos no plano contratado pelo tenant

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-05** | Gestão de Usuários e Permissões | [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) / J1: Convite de novo usuário · J2: Restrição de acesso entre filiais | **FEAT-EP-0003-0003** — Vinculação Usuário × Unidade × Módulo |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
