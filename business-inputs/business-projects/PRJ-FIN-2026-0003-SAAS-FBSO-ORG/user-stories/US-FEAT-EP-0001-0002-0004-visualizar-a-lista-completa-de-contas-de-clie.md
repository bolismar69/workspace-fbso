# User Story: US-FEAT-EP-0001-0002-0004 — visualizar a lista completa de contas de clientes com informações resu

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) ➔ Feature [FEAT-EP-0001-0002](../features/FEAT-EP-0001-0002-visao-de-contas-com-filtros.md) ➔ User Story US-FEAT-EP-0001-0002-0004
- **Prioridade:** Must Have
- **Data-Alvo:** 15/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador FBSO,
- **quero** visualizar a lista completa de contas de clientes com informações resumidas (nome, plano, status, data de criação) para navegar pela base,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Tabela exibe: razão social, plano contratado, status, data de criação, data da última ação**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Lista ordenada por data de criação (mais recentes primeiro)**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Paginação a cada 25 registros**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0001-0002 — Visão de Contas com Filtros](../features/FEAT-EP-0001-0002-visao-de-contas-com-filtros.md) | **Épico:** [EP-0001 — Portal Administrativo Interno](../epics/EP-0001-portal-administrativo-interno.md)
