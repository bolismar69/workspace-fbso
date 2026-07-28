# User Story: US-FEAT-EP-0004-0003-0062 — realizar upgrade do plano contratado diretamente pelo portal

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0004](../epics/EP-0004-portal-do-cliente-auto-servico.md) ➔ Feature [FEAT-EP-0004-0003](../features/FEAT-EP-0004-0003-dashboard-do-cliente.md) ➔ User Story US-FEAT-EP-0004-0003-0062
- **Prioridade:** Should Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.0 — Gap Analysis #16 (27/07/2026)
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Administrador do Tenant,
- **quero** realizar o upgrade do meu plano contratado diretamente pelo portal, selecionando um plano superior e visualizando os novos módulos que ficarão disponíveis,
- **para** expandir as funcionalidades da minha conta de forma autônoma, sem precisar contatar o time comercial da FBSO.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal — Visualizar planos disponíveis para upgrade]
- **Dado que** o Administrador do Tenant está autenticado no portal e acessa a área "Meu Plano" no dashboard,
- **Quando** visualiza as opções de planos disponíveis,
- **Então** o sistema deve: **exibir apenas planos de nível superior ao plano atual (upgrade), com comparação clara de: nome do plano, valor mensal, módulos incluídos em cada um, destacando os módulos NOVOS que o upgrade adicionaria**.

### Cenário 2: [Fluxo Principal — Confirmar upgrade]
- **Dado que** o Administrador do Tenant seleciona um plano superior e revisa as mudanças,
- **Quando** confirma o upgrade,
- **Então** o sistema deve: **registrar a solicitação de upgrade, encerrar a assinatura atual com data de término = hoje, criar nova assinatura com o plano selecionado com data de início = hoje, e liberar imediatamente o acesso aos novos módulos do plano**.

### Cenário 3: [Notificação ao time comercial]
- **Dado que** um upgrade de plano é confirmado pelo Administrador do Tenant,
- **Quando** o sistema processa a alteração,
- **Então** o sistema deve: **notificar o time comercial da FBSO sobre o upgrade realizado (tenant, plano anterior → novo plano, data, valor) para follow-up comercial e faturamento futuro**.

### Cenário 4: [Restrição — Apenas upgrade, não downgrade]
- **Dado que** o Administrador do Tenant acessa a área "Meu Plano",
- **Quando** visualiza as opções disponíveis,
- **Então** o sistema deve: **NÃO exibir planos inferiores ao atual. Downgrade ou cancelamento devem ser solicitados via contato com o time comercial da FBSO**.

### Cenário 5: [Histórico de assinaturas]
- **Dado que** um upgrade foi realizado,
- **Quando** o Administrador do Tenant acessa o histórico de assinaturas,
- **Então** o sistema deve: **exibir o histórico completo: assinatura anterior (plano, período, status "Encerrada"), assinatura atual (plano, data de início, status "Ativa"), com indicação de que o upgrade foi realizado via self-service pelo portal**.

### Cenário 6: [Auditoria]
- **Dado que** um upgrade de plano é concluído,
- **Quando** o sistema registra a operação,
- **Então** o sistema deve: **registrar no histórico de auditoria: administrador do tenant responsável, data/hora, plano anterior, novo plano, módulos adicionados e tipo da ação (UPGRADE_PLANO_SELF_SERVICE)**.

---

## 3. Regras de Negócio de Tela Relacionadas

- **RN-FEAT-EP-0004-0003-0062-01:** Apenas usuários com perfil `admin-do-tenant` podem realizar upgrade de plano pelo portal
- **RN-FEAT-EP-0004-0003-0062-02:** Somente upgrade é permitido via self-service (plano superior). Downgrade e cancelamento exigem contato com o time comercial
- **RN-FEAT-EP-0004-0003-0062-03:** O upgrade tem efeito imediato — novos módulos são liberados no momento da confirmação
- **RN-FEAT-EP-0004-0003-0062-04:** A cobrança proporcional (pro-rata) do novo plano será tratada na fase de faturamento — nesta fase, o registro da transição é suficiente
- **RN-FEAT-EP-0004-0003-0062-05:** O upgrade não afeta usuários, permissões ou dados já existentes no tenant — apenas adiciona acesso aos novos módulos

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0004-0003 — Dashboard do Cliente](../features/FEAT-EP-0004-0003-dashboard-do-cliente.md) | **Épico:** [EP-0004 — Portal do Cliente Auto Servico](../epics/EP-0004-portal-do-cliente-auto-servico.md)
