# User Story: US-FEAT-EP-0001-0003-0006 — ver indicadores de alerta no dashboard para contas que precisam de ate

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0001](../epics/EP-0001-portal-administrativo-interno.md) ➔ Feature [FEAT-EP-0001-0003](../features/FEAT-EP-0001-0003-alertas-e-indicadores-de-atencao.md) ➔ User Story US-FEAT-EP-0001-0003-0006
- **Prioridade:** Should Have
- **Data-Alvo:** 15/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador FBSO,
- **quero** ver indicadores de alerta no dashboard para contas que precisam de atenção (ex: onboarding incompleto há mais de 48h, assinatura suspensa),
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Alertas aparecem como cards coloridos (amarelo: atenção; vermelho: crítico) no topo do dashboard**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Cada alerta exibe a quantidade de contas na situação e é clicável**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Administrador FBSO está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Ao clicar, direciona para a lista filtrada das contas naquela situação**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0001-0003 — Alertas e Indicadores de Atenção](../features/FEAT-EP-0001-0003-alertas-e-indicadores-de-atencao.md) | **Épico:** [EP-0001 — Portal Administrativo Interno](../epics/EP-0001-portal-administrativo-interno.md)
