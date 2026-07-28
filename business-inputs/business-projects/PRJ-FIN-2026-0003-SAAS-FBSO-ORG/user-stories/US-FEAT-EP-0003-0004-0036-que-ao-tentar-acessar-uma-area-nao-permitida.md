# User Story: US-FEAT-EP-0003-0004-0036 — que ao tentar acessar uma área não permitida diretamente (via URL ou a

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0004](../features/FEAT-EP-0003-0004-controle-de-visibilidade-de-menus-e-acoes.md) ➔ User Story US-FEAT-EP-0003-0004-0036
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Usuário do Portal,
- **quero** que ao tentar acessar uma área não permitida diretamente (via URL ou atalho), o sistema me redirecione para uma tela de "Acesso Negado" com explicação amigável,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Usuário do Portal está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Tela de acesso negado com mensagem clara e não-técnica**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Usuário do Portal está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Exibe: "Você não tem permissão para acessar esta área. Se precisar de acesso, contate o administrador da sua conta."**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Usuário do Portal está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Não exibe detalhes técnicos ou caminhos internos**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0004 — Controle de Visibilidade de Menus e Ações](../features/FEAT-EP-0003-0004-controle-de-visibilidade-de-menus-e-acoes.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
