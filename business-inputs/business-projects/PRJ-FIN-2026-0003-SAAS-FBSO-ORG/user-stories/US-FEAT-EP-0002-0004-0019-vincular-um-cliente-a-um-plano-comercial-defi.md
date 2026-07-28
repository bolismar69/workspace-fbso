# User Story: US-FEAT-EP-0002-0004-0019 — vincular um cliente a um plano comercial definindo data de início, vig

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) ➔ Feature [FEAT-EP-0002-0004](../features/FEAT-EP-0002-0004-vinculacao-e-gestao-de-assinaturas.md) ➔ User Story US-FEAT-EP-0002-0004-0019
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador FBSO,
- **quero** vincular um cliente a um plano comercial definindo data de início, vigência e status da assinatura,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Seletor de plano (apenas planos ativos listados)**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Campos: data de início, data de término (ou "indeterminado"), status (Ativa, Suspensa, Cancelada)**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Ao ativar assinatura, módulos do plano são liberados para o tenant**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0002-0004 — Vinculação e Gestão de Assinaturas](../features/FEAT-EP-0002-0004-vinculacao-e-gestao-de-assinaturas.md) | **Épico:** [EP-0002 — Gestao de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)
