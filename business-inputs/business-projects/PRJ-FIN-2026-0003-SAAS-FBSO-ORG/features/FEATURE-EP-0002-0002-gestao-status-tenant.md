# FEATURE - EP-0002-0002: Gestão de Status do Tenant

| Campo | Detalhe |
|-------|---------|
| **Feature** | EP-0002-0002 — Gestão de Status do Tenant |
| **Épico** | [EP-0002 — Gestão de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | **Anterior:** [EP-0002-0001 — Cadastro e Ativação](../FEATURE-EP-0002-0001-cadastro-ativacao-contas-clientes.md) | **Próximo:** [EP-0002-0003 — Planos Comerciais](../FEATURE-EP-0002-0003-configuracao-planos-comerciais.md)

**Requisitos BRD Vinculados:** [BR-02](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Ativação e Gestão de Contas

---

## Objetivo de Negócio
Controlar o ciclo de vida de cada conta de cliente, permitindo ativar, suspender e reativar conforme a situação operacional ou comercial.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-012 | Como **Administrador FBSO**, quero alterar o status de uma conta de cliente entre os estados: Pendente Onboarding, Ativo, Suspenso, Inativo para refletir a situação real da conta | • Seletor de status na tela de detalhes do Tenant • Transições permitidas seguem regra de negócio RN-FEAT-EP-0002-0002-0001 • Ao alterar para Suspenso, campo de motivo é obrigatório |
| US-013 | Como **Administrador FBSO**, quero que ao suspender uma conta, todos os usuários daquele tenant tenham o acesso ao portal bloqueado imediatamente | • Bloqueio efetivo em até 5 minutos após a suspensão • Usuários logados recebem mensagem de sessão encerrada na próxima ação • Status do tenant atualizado em tempo real para "Suspenso" |
| US-014 | Como **Administrador FBSO**, quero visualizar o histórico de mudanças de status de cada conta (quando foi ativada, suspensa, reativada e por quem) | • Linha do tempo de status na tela de detalhes do Tenant • Cada evento exibe: status anterior → novo status, responsável, data/hora, motivo (se aplicável) |

## Regras de Negócio

- **RN-FEAT-EP-0002-0002-0001:** Transições de status permitidas: Pendente Onboarding → Ativo (quando cliente completa onboarding); Ativo → Suspenso (inadimplência ou solicitação administrativa); Suspenso → Ativo (reativação); Ativo → Inativo (encerramento de contrato); Inativo → Ativo (recontratação)
- **RN-FEAT-EP-0002-0002-0002:** Suspensão exige motivo registrado (campo obrigatório)
- **RN-FEAT-EP-0002-0002-0003:** Reativação de conta suspensa restaura as permissões anteriores dos usuários

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-02** | Ativação e Gestão de Contas | [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) / J2: Suspensão por inadimplência | **EP-0002-0002** — Gestão de Status do Tenant |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
