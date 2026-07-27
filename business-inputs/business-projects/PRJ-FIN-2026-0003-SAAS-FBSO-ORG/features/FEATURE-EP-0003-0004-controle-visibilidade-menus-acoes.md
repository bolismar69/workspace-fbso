# FEATURE - EP-0003-0004: Controle de Visibilidade de Menus e Ações

| Campo | Detalhe |
|-------|---------|
| **Feature** | EP-0003-0004 — Controle de Visibilidade de Menus e Ações |
| **Épico** | [EP-0003 — Governança de Acessos e Permissões](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) | **Anterior:** [EP-0003-0003 — Vinculação Usuário×Unidade×Módulo](../FEATURE-EP-0003-0003-vinculacao-usuario-unidade-modulo.md) | **Próximo:** [EP-0004-0001 — Autenticação](../FEATURE-EP-0004-0001-autenticacao-recuperacao-senha.md)

**Requisitos BRD Vinculados:** [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Gestão de Usuários e Permissões

---

## Objetivo de Negócio
Garantir que a interface do portal se adapte dinamicamente às permissões do usuário, ocultando menus, botões e funcionalidades que ele não está autorizado a acessar.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-034 | Como **Usuário do Portal**, quero que o menu lateral exiba apenas as opções correspondentes às minhas permissões (papel + módulo ativo) para ter uma interface limpa e focada no meu trabalho | • Menu lateral renderizado dinamicamente conforme permissões • Itens de menu sem permissão não aparecem (não ficam desabilitados — simplesmente não renderizam) • Ao trocar de módulo no App Switcher, menu se adapta |
| US-035 | Como **Usuário do Portal**, quero que botões de ação (Criar, Editar, Excluir) apareçam apenas se eu tiver permissão para executar aquela ação | • Botão "Novo" visível apenas para quem tem permissão de criação • Botão "Editar" visível apenas para quem tem permissão de edição • Botão "Excluir" visível apenas para quem tem permissão de exclusão |
| US-036 | Como **Usuário do Portal**, quero que ao tentar acessar uma área não permitida diretamente (via URL ou atalho), o sistema me redirecione para uma tela de "Acesso Negado" com explicação amigável | • Tela de acesso negado com mensagem clara e não-técnica • Exibe: "Você não tem permissão para acessar esta área. Se precisar de acesso, contate o administrador da sua conta." • Não exibe detalhes técnicos ou caminhos internos |

## Regras de Negócio

- **RN-FEAT-EP-0003-0004-0001:** Ocultação de menu é a primeira camada (UX); o bloqueio por permissão no acesso direto é a camada de segurança — ambas devem ser implementadas
- **RN-FEAT-EP-0003-0004-0002:** Usuários veem o nome do módulo ativo no topo do portal, ao lado do logo da FBSO Platform

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-05** | Gestão de Usuários e Permissões | [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md) / J2: Restrição de acesso entre filiais · J3: Revogação de acesso | **EP-0003-0004** — Controle de Visibilidade de Menus e Ações |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0003](../epics/EP-0003-governanca-de-acessos-e-permissoes.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
