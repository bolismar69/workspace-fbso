# EP-0001: Portal Administrativo Interno

| Campo | Detalhe |
|-------|---------|
| **Épico** | EP-0001 — Portal Administrativo Interno |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Épicos (Estrutura Modular v4.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 e `01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Próximo:** [EP-0002 — Gestão de Clientes e Assinaturas](../EP-0002-gestao-de-clientes-e-assinaturas.md)

---

## 1. Nome do Épico
**Portal Administrativo Interno — Painel de Controle FBSO.ORG**

**Requisitos BRD Vinculados:** [BR-01](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Dashboard Administrativo

## 2. Objetivo (Goal)

- **Problema:** A FBSO.ORG não possui uma interface centralizada para visualizar e acompanhar a operação do SaaS. Os dados sobre clientes, planos e status das contas estão dispersos ou inexistem, impossibilitando a tomada de decisão operacional ágil e baseada em dados. O time administrativo não tem como saber, em uma única tela, quantos clientes estão ativos, quais planos foram contratados ou se há contas que precisam de atenção.
- **Solução:** Construir um painel de controle (dashboard) administrativo que concentre as principais métricas operacionais do SaaS e permita ao time interno da FBSO.ORG navegar pela base de clientes com filtros e visões consolidadas por período, plano e status.
- **Impacto:** Time interno ganha autonomia e agilidade para acompanhar a saúde da operação; redução do tempo gasto buscando informações em fontes dispersas; base para decisões comerciais e operacionais informadas.

## 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Administrador FBSO.ORG** | Colaborador do time interno responsável pela operação do SaaS | Visão rápida da base de clientes; identificar contas com problemas; métricas de crescimento |
| **Líder Comercial** | Responsável pela estratégia de vendas e relacionamento com clientes | Acompanhar adoção de planos; identificar oportunidades de upgrade; visão de contas por status |
| **Diretoria** | Sócios e diretores da FBSO.ORG | Visão macro da operação; indicadores de crescimento; saúde financeira da base |

## 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Acompanhamento diário da operação**
1. Administrador acessa o portal administrativo
2. Visualiza dashboard com métricas do dia: total de contas ativas, novas contas, contas suspensas
3. Identifica visualmente contas que precisam de atenção (ex: onboarding incompleto, assinatura vencida)
4. Clica em um indicador para ver a lista detalhada de contas naquela situação
> 🏷️ Atende [BR-01](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 2: Análise de crescimento por plano**
1. Líder comercial acessa o dashboard e aplica filtro por período (ex: últimos 30 dias)
2. Visualiza distribuição de clientes por plano (Básico, Core, Full Suite)
3. Identifica plano com maior crescimento no período
4. Visualiza os dados filtrados para compartilhar em apresentação de resultados (exportação de relatórios disponível em fase futura)
> 🏷️ Atende [BR-01](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

## 5. Requisitos de Negócio (Business Requirements)

### Requisitos Funcionais

- Exibição de indicadores-chave: total de contas ativas, contas novas no período, contas suspensas, distribuição por plano
- Filtro por período para todos os indicadores (dia, semana, mês, trimestre, ano)
- Visão de contas agrupadas por status (ativo, inativo, onboarding pendente, suspenso)
- Lista de contas com busca por nome, status e plano
- Indicadores visuais de alerta para situações críticas (ex: contas suspensas por inadimplência, onboarding incompleto há mais de X dias)
- Gráfico de evolução da base de clientes ao longo do tempo
- Estrutura preparada para receber métricas financeiras futuras (MRR, Churn) quando a comercialização for ativada

### Requisitos Não-Funcionais

- Dashboard deve carregar em até 3 segundos com volume de dados previsto para o primeiro ano
- Métricas devem refletir dados atualizados (atraso máximo de atualização a definir)
- Acesso restrito a usuários do time interno FBSO.ORG com permissão administrativa

## 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Tempo para obter visão completa da operação (da tela de login à informação desejada) | ≤ 30 segundos |
| Satisfação do time administrativo com o dashboard | Nota ≥ 4,0 / 5,0 |
| Redução de perguntas operacionais para o time técnico | ≥ 50% de redução |

## 7. Fora do Escopo (Out of Scope)

- Métricas financeiras detalhadas (MRR, Churn Rate, LTV) — disponíveis apenas na fase de comercialização
- Dashboards customizáveis por usuário — nesta fase, dashboard único para todos os administradores
- Exportação de relatórios em PDF/Excel — funcionalidade futura
- Previsões ou análises preditivas — funcionalidade futura

## 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Alto** | É a principal ferramenta de trabalho do time interno; sem ela, a operação do SaaS é inviável em escala |

---

## Matriz de Rastreabilidade BRD → Este Épico

| BRD | Requisito Funcional | Este Épico | Jornada(s) que Realizam |
|:---|:---|:---|:---|
| **BR-01** | Dashboard Administrativo | **EP-0001** — Portal Administrativo Interno | J1: Acompanhamento diário da operação · J2: Análise de crescimento por plano |

---

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Próximo:** [EP-0002 — Gestão de Clientes e Assinaturas](../EP-0002-gestao-de-clientes-e-assinaturas.md)

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: breakdown-epic-pm, agile-ba-practices. Estrutura modular v4.0.*

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE ÉPICOS]
