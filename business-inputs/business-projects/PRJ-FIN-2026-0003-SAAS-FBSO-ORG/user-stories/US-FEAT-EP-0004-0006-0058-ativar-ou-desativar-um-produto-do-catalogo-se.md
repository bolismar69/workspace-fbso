# User Story: US-FEAT-EP-0004-0006-0058 — ativar ou desativar um produto do catálogo sem excluí-lo definitivamen

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0004](../epics/EP-0004-portal-do-cliente-auto-servico.md) ➔ Feature [FEAT-EP-0004-0006](../features/FEAT-EP-0004-0006-catalogo-de-produtos-e-servicos.md) ➔ User Story US-FEAT-EP-0004-0006-0058
- **Prioridade:** Must Have
- **Data-Alvo:** 15/10/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Cliente,
- **quero** ativar ou desativar um produto do catálogo sem excluí-lo definitivamente,
- **para** controlar quais itens estão em uso.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Cliente está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Botão "Desativar" na lista (para itens ativos)**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Cliente está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Botão "Ativar" na lista (para itens inativos)**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Cliente está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Item desativado não aparece em cadastros futuros, mas mantém histórico**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0004-0006 — Catálogo de Produtos e Serviços](../features/FEAT-EP-0004-0006-catalogo-de-produtos-e-servicos.md) | **Épico:** [EP-0004 — Portal do Cliente Auto Servico](../epics/EP-0004-portal-do-cliente-auto-servico.md)
