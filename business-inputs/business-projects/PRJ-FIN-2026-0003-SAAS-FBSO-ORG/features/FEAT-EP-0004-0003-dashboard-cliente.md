# FEATURE - FEAT-EP-0004-0003: Dashboard do Cliente

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0004-0003 — Dashboard do Cliente |
| **Épico** | [EP-0004 — Experiência do Cliente e Autoatendimento](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | **Anterior:** [FEAT-EP-0004-0002 — Onboarding](../FEAT-EP-0004-0002-onboarding-guiado-primeiro-acesso.md) | **Próximo:** [FEAT-EP-0004-0004 — App Switcher](../FEAT-EP-0004-0004-app-switcher-seletor-modulos.md)

**Requisitos BRD Vinculados:** [BR-06](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Portal do Cliente

---

## Objetivo de Negócio
Oferecer ao cliente uma tela inicial com visão geral de sua conta e acessos rápidos às principais funcionalidades.

**Prioridade:** Should Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-FEAT-EP-0004-0003-0045 | Como **Cliente**, quero visualizar um dashboard com informações resumidas da minha conta: unidades de negócio ativas, total de produtos cadastrados e meu plano atual | • Dashboard com cards informativos: Unidades Ativas, Produtos no Catálogo, Plano Contratado • Cada card é clicável e leva à respectiva área de gestão |
| US-FEAT-EP-0004-0003-0046 | Como **Cliente**, quero ver notificações e lembretes relevantes no meu dashboard (ex: "Complete seu cadastro de produtos", "Convite de usuário pendente") | • Área de notificações visível no dashboard • Cada notificação com link para a ação relacionada • Notificações podem ser dispensadas pelo usuário |
| US-FEAT-EP-0004-0003-0062 | Como **Administrador do Tenant**, quero realizar o upgrade do meu plano contratado diretamente pelo portal, selecionando um plano superior e visualizando os novos módulos que ficarão disponíveis para expandir as funcionalidades da minha conta de forma autônoma, sem precisar contatar o time comercial da FBSO | • Área 'Meu Plano' no dashboard do cliente exibe apenas planos superiores (upgrade) com comparação clara de nome, valor e módulos • Confirmação de upgrade encerra assinatura atual, cria nova assinatura e libera novos módulos imediatamente • Time comercial é notificado sobre o upgrade • Downgrade e cancelamento NÃO disponíveis via self-service (apenas via contato comercial) • Histórico de assinaturas exibe transição com indicação 'UPGRADE_PLANO_SELF_SERVICE' |

## Regras de Negócio

- **RN-FEAT-EP-0004-0003-0001:** Dashboard do cliente adapta-se ao módulo ativo no App Switcher (métricas diferentes por módulo)
- **RN-FEAT-EP-0004-0003-0002:** Na Fase 0 (este projeto), haverá um dashboard genérico que será expandido quando os módulos forem ativados

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-06** | Portal do Cliente com Autenticação | [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) / J1: Primeiro acesso e onboarding | **FEAT-EP-0004-0003** — Dashboard do Cliente |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
