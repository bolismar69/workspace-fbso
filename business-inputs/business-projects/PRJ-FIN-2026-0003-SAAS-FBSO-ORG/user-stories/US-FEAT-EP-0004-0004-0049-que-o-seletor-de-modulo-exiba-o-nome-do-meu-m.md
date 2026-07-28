# User Story: US-FEAT-EP-0004-0004-0049 — que o Seletor de Módulo exiba o nome do meu módulo mesmo que eu não te

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0004](../epics/EP-0004-portal-do-cliente-auto-servico.md) ➔ Feature [FEAT-EP-0004-0004](../features/FEAT-EP-0004-0004-app-switcher-seletor-de-modulos.md) ➔ User Story US-FEAT-EP-0004-0004-0049
- **Prioridade:** Must Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Cliente com apenas um módulo contratado,
- **quero** que o Seletor de Módulo exiba o nome do meu módulo mesmo que eu não tenha outras opções,
- **para** que eu saiba em qual produto estou.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Cliente com apenas um módulo contratado está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Seletor de Módulo visível mesmo com um único módulo**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Cliente com apenas um módulo contratado está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Exibe o nome do módulo ativo sem dropdown de seleção**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Cliente com apenas um módulo contratado está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Indica visualmente que novos módulos podem ser adicionados no futuro**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0004-0004 — App Switcher (Seletor de Módulos)](../features/FEAT-EP-0004-0004-app-switcher-seletor-de-modulos.md) | **Épico:** [EP-0004 — Portal do Cliente Auto Servico](../epics/EP-0004-portal-do-cliente-auto-servico.md)
