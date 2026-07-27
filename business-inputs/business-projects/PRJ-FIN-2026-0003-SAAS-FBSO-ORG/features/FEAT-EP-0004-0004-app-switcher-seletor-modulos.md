# FEATURE - FEAT-EP-0004-0004: App Switcher (Seletor de Módulos)

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0004-0004 — App Switcher (Seletor de Módulos) |
| **Épico** | [EP-0004 — Experiência do Cliente e Autoatendimento](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | **Anterior:** [FEAT-EP-0004-0003 — Dashboard](../FEAT-EP-0004-0003-dashboard-cliente.md) | **Próximo:** [FEAT-EP-0004-0005 — Unidades de Negócio](../FEAT-EP-0004-0005-gestao-unidades-negocio.md)

**Requisitos BRD Vinculados:** [BR-08](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — App Switcher

---

## Objetivo de Negócio
Fornecer o mecanismo de navegação entre módulos da plataforma, permitindo que o cliente alterne entre diferentes produtos sem sair do portal.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-047 | Como **Cliente**, quero ver um seletor de aplicativos no topo do portal que exiba os módulos disponíveis no meu plano para navegar entre os produtos que contratei | • App Switcher posicionado no canto superior, ao lado do logo da FBSO Platform • Exibe apenas módulos que o usuário tem permissão para acessar • Módulo ativo aparece destacado visualmente |
| US-048 | Como **Cliente**, quero que ao selecionar um módulo diferente no App Switcher, o menu lateral e o conteúdo da tela se adaptem imediatamente ao módulo escolhido | • Troca de módulo atualiza menu lateral em tempo real • Conteúdo da tela é redirecionado para o dashboard do módulo selecionado • Transição fluida, sem recarregamento completo da página |
| US-049 | Como **Cliente com apenas um módulo contratado**, quero que o Seletor de Módulo exiba o nome do meu módulo mesmo que eu não tenha outras opções, para que eu saiba em qual produto estou | • Seletor de Módulo visível mesmo com um único módulo • Exibe o nome do módulo ativo sem dropdown de seleção • Indica visualmente que novos módulos podem ser adicionados no futuro |

## Regras de Negócio

- **RN-FEAT-EP-0004-0004-0001:** Lista de módulos no App Switcher é determinada pela interseção entre: módulos do plano contratado e módulos que o usuário tem permissão
- **RN-FEAT-EP-0004-0004-0002:** Na Fase 0, haverá um módulo placeholder chamado "FBSO Platform" visível para todos os clientes
- **RN-FEAT-EP-0004-0004-0003:** A troca de módulo mantém o contexto da Unidade de Negócio selecionada

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-08** | App Switcher (Seletor de Aplicativos) | [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) / J4: Navegação com App Switcher | **FEAT-EP-0004-0004** — App Switcher |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
