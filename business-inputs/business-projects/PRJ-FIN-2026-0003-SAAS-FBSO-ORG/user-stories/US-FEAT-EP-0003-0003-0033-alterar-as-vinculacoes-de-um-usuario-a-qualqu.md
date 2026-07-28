# User Story: US-FEAT-EP-0003-0003-0033 — alterar as vinculações de um usuário a qualquer momento (adicionar/rem

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0003](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md) ➔ Feature [FEAT-EP-0003-0003](../features/FEAT-EP-0003-0003-vinculacao-usuario-unidade-modulo.md) ➔ User Story US-FEAT-EP-0003-0003-0033
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** alterar as vinculações de um usuário a qualquer momento (adicionar/remover unidade, adicionar/remover módulo) com efeito imediato,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Alterações salvas têm efeito na próxima ação do usuário**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Se usuário estiver logado e tiver acesso a uma unidade removida, a sessão é ajustada**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador do Tenant está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Registro de auditoria gerado para cada alteração**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0003-0003 — Vinculação Usuário × Unidade × Módulo](../features/FEAT-EP-0003-0003-vinculacao-usuario-unidade-modulo.md) | **Épico:** [EP-0003 — Gestao de Usuarios e Permissoes RBAC](../epics/EP-0003-gestao-de-usuarios-e-permissoes-rbac.md)
