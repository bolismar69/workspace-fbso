# FEATURE - FEAT-EP-0004-0001: Autenticação e Recuperação de Senha

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0004-0001 — Autenticação e Recuperação de Senha |
| **Épico** | [EP-0004 — Experiência do Cliente e Autoatendimento](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | **Anterior:** [FEAT-EP-0003-0004 — Controle de Visibilidade](../FEAT-EP-0003-0004-controle-visibilidade-menus-acoes.md) | **Próximo:** [FEAT-EP-0004-0002 — Onboarding](../FEAT-EP-0004-0002-onboarding-guiado-primeiro-acesso.md)

**Requisitos BRD Vinculados:** [BR-06](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Portal do Cliente com Autenticação

---

## Objetivo de Negócio
Oferecer ao cliente uma experiência segura e fluida de acesso ao portal, com fluxo de recuperação de senha autônomo.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-037 | Como **Cliente**, quero fazer login no portal usando meu e-mail e senha para acessar minha conta | • Tela de login com campos: e-mail e senha • Mensagens de erro genéricas (não revelar se o e-mail existe ou não) • Após login bem-sucedido, redirecionar para o dashboard ou onboarding (se primeiro acesso) |
| US-038 | Como **Cliente**, quero recuperar minha senha caso eu a esqueça, recebendo um link de redefinição por e-mail | • Opção "Esqueci minha senha" na tela de login • Usuário informa e-mail e recebe link de redefinição • Link expira em 1 hora • Nova senha deve atender critérios mínimos de complexidade |
| US-039 | Como **Cliente**, quero que ao errar a senha repetidas vezes, minha conta seja temporariamente bloqueada por segurança | • Após 5 tentativas consecutivas com erro, conta bloqueada por 15 minutos • Mensagem informa o tempo restante de bloqueio • Administrador do tenant pode desbloquear manualmente |

## Regras de Negócio

- **RN-FEAT-EP-0004-0001-0001:** Senha deve ter no mínimo 8 caracteres, incluindo letra e número
- **RN-FEAT-EP-0004-0001-0002:** Sessão expira após 60 minutos de inatividade
- **RN-FEAT-EP-0004-0001-0003:** Link de redefinição de senha é de uso único

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-06** | Portal do Cliente com Autenticação | [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) / J1: Primeiro acesso e onboarding | **FEAT-EP-0004-0001** — Autenticação e Recuperação de Senha |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
