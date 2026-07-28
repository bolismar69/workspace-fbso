# FEATURE - FEAT-EP-0003-0002: Definição de Papéis e Permissões (RBAC)

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0003-0002 — Definição de Papéis e Permissões (RBAC) |
| **Épico** | [EP-0003 — Governança de Acessos e Permissões](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) | **Anterior:** [FEAT-EP-0003-0001 — Cadastro e Convite](../FEAT-EP-0003-0001-cadastro-convite-usuarios.md) | **Próximo:** [FEAT-EP-0003-0003 — Vinculação Usuário×Unidade×Módulo](../FEAT-EP-0003-0003-vinculacao-usuario-unidade-modulo.md)

**Requisitos BRD Vinculados:** [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Gestão de Usuários e Permissões

---

## Objetivo de Negócio
Estabelecer os papéis padrão da plataforma com conjuntos de permissões bem definidos, garantindo que cada usuário acesse apenas o que seu perfil permite.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-FEAT-EP-0003-0002-0027 | Como **Administrador do Tenant**, quero atribuir um dos papéis padrão (Admin, Gerente, Operador, Auditor) a cada usuário para definir seu nível de acesso na plataforma | • Seletor de papel no cadastro e na edição do usuário • Descrição de cada papel disponível como tooltip • Alteração de papel registrada em auditoria |
| US-FEAT-EP-0003-0002-0028 | Como **Gestor de Produto**, quero que cada papel tenha um conjunto predefinido de permissões: Admin do Tenant (acesso total), Gerente de Unidade (gerencia sua unidade), Operador (executa tarefas), Auditor (apenas visualiza) | • Permissões mapeadas conforme tabela de papéis (RN-FEAT-EP-0003-0002-0001) • Permissões não são customizáveis por tenant nesta fase • Papéis são os mesmos para todos os tenants |
| US-FEAT-EP-0003-0002-0029 | Como **Administrador do Tenant**, quero que ao atribuir o papel "Admin do Tenant" a um usuário, ele automaticamente tenha acesso a todas as Unidades de Negócio e todos os módulos do tenant | • Admin do Tenant vê todas as unidades de negócio no seletor • Admin do Tenant vê todos os módulos contratados no App Switcher • Não é necessário configurar permissões individuais para Admin |
| US-FEAT-EP-0003-0002-0030 | Como **Administrador do Tenant**, quero que ao atribuir o papel "Auditor" a um usuário, ele possa visualizar todos os dados das unidades permitidas mas não possa criar, editar ou excluir nada | • Botões de criação/edição/exclusão não visíveis para Auditor • Menus de configuração não aparecem para Auditor • Tentativa de acesso direto a funcionalidades de escrita é bloqueada |

## Regras de Negócio

- **RN-FEAT-EP-0003-0002-0001:** Tabela de Permissões por Papel:

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

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-05** | Gestão de Usuários e Permissões | [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) / J1: Convite de novo usuário | **FEAT-EP-0003-0002** — Definição de Papéis e Permissões (RBAC) |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
