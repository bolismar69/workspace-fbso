# User Story: US-FEAT-EP-0004-0001-0038 — recuperar minha senha caso eu a esqueça, recebendo um link de redefini

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0004](../epics/EP-0004-portal-do-cliente-auto-servico.md) ➔ Feature [FEAT-EP-0004-0001](../features/FEAT-EP-0004-0001-autenticacao-e-recuperacao-de-senha.md) ➔ User Story US-FEAT-EP-0004-0001-0038
- **Prioridade:** Must Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Cliente,
- **quero** recuperar minha senha caso eu a esqueça, recebendo um link de redefinição por e-mail,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Cliente está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Opção "Esqueci minha senha" na tela de login**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Cliente está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Usuário informa e-mail e recebe link de redefinição**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Cliente está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Link expira em 1 hora**.

### Cenário 4: [Fluxo Alternativo 4]
- **Dado que** Cliente está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Nova senha deve atender critérios mínimos de complexidade**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0004-0001 — Autenticação e Recuperação de Senha](../features/FEAT-EP-0004-0001-autenticacao-e-recuperacao-de-senha.md) | **Épico:** [EP-0004 — Portal do Cliente Auto Servico](../epics/EP-0004-portal-do-cliente-auto-servico.md)
