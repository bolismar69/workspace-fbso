# User Stories: Alertas e Indicadores de Atenção

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0001-0003 — Alertas e Indicadores de Atenção
- **Épico:** EP-0001 — Portal Administrativo Interno
- **Prioridade:** Should Have
- **Data-Alvo:** 15/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Destacar proativamente situações que exigem ação do time administrativo, evitando que problemas passem despercebidos.

---

## User Stories

### US-006 — Indicadores de Alerta no Dashboard

**Como** Administrador FBSO, **quero** ver indicadores de alerta no dashboard para contas que precisam de atenção (ex: onboarding incompleto há mais de 48h, assinatura suspensa).

**Critérios de Aceitação:**
- Alertas aparecem como cards coloridos (amarelo: atenção; vermelho: crítico) no topo do dashboard
- Cada alerta exibe a quantidade de contas na situação e é clicável
- Ao clicar, direciona para a lista filtrada das contas naquela situação
- Quando não há alertas ativos, seção de alertas exibe mensagem 'Nenhum alerta no momento'
- Alertas são recalculados a cada carregamento do dashboard (não atualizam em tempo real nesta fase)

### US-007 — Destaque Visual de Contas Irregulares

**Como** Administrador FBSO, **quero** que o sistema destaque visualmente na lista de contas aquelas com status irregular **para** identificação rápida durante a navegação.

**Critérios de Aceitação:**
- Contas suspensas exibem ícone ou cor de destaque na lista
- Ao passar o cursor, tooltip explica o motivo da suspensão. Se motivo não registrado, exibe 'Motivo não informado.'

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0001-0003-0001** | Alerta de "onboarding incompleto" dispara após 48 horas da criação da conta |
| **RN-FEAT-EP-0001-0003-0002** | Alertas são visíveis para todos os usuários do time interno; não há personalização por usuário nesta fase |
| **RN-FEAT-EP-0001-0003-0003** | Classificação de alertas por cor: Amarelo (atenção) = onboarding incompleto há mais de 48h. Vermelho (crítico) = assinatura suspensa, conta suspensa. Esta classificação é fixa nesta fase. |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Cards de alerta coloridos no dashboard com contagem e link para lista filtrada | Print do dashboard com alertas ativos |
| F2 | Destaque visual em contas suspensas na lista com tooltip explicativo | Print da lista com conta suspensa em destaque |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
