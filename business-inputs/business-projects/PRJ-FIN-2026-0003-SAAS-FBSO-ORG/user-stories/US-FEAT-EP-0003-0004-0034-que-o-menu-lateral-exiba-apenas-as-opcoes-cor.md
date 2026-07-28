# User Story: US-FEAT-EP-0003-0004-0034 — que o menu lateral exiba apenas as opções correspondentes às minhas pe

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0004](../features/FEAT-EP-0003-0004-controle-de-visibilidade-de-menus-e-acoes.md) ➔ User Story US-FEAT-EP-0003-0004-0034
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Usuário do Portal,
- **quero** que o menu lateral exiba apenas as opções correspondentes às minhas permissões (papel + módulo ativo) para ter uma interface limpa e focada no meu trabalho,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Usuário do Portal está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Menu lateral renderizado dinamicamente conforme permissões**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Usuário do Portal está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Itens de menu sem permissão não aparecem (não ficam desabilitados — simplesmente não renderizam)**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Usuário do Portal está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Ao trocar de módulo no App Switcher, menu se adapta**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0004 — Controle de Visibilidade de Menus e Ações](../features/FEAT-EP-0003-0004-controle-de-visibilidade-de-menus-e-acoes.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
