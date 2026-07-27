# User Stories: Dashboard de Métricas Operacionais

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0001-0001 — Dashboard de Métricas Operacionais
- **Épico:** EP-0001 — Portal Administrativo Interno
- **Prioridade:** Must Have
- **Data-Alvo:** 15/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Prover visibilidade em tempo real da operação do SaaS para o time interno da FBSO.ORG, permitindo acompanhar a saúde da base de clientes e tomar decisões operacionais com agilidade.

---

## User Stories

### US-001 — Dashboard de Indicadores Principais

**Como** Administrador FBSO, **quero** visualizar os indicadores principais da operação em uma tela de dashboard **para** ter uma visão rápida da saúde do SaaS.

**Critérios de Aceitação:**
- Dashboard carrega com os indicadores atualizados em até 3 segundos
- Indicadores exibidos: total de contas ativas, total de contas por status, total de contas por plano, taxa de conversão de onboarding (contas que completaram onboarding / total de contas criadas no período)
- Cada indicador é clicável e leva à lista filtrada correspondente
- Dashboard com base de clientes vazia exibe todos os indicadores com valor zero (RN-FEAT-EP-0001-0001-0003)
- Em caso de falha no carregamento dos dados, dashboard exibe mensagem informativa com opção de tentar novamente
- Indicadores exibem skeleton/loader durante o carregamento inicial

### US-002 — Filtro de Métricas por Período

**Como** Líder Comercial, **quero** filtrar as métricas do dashboard por período (dia, últimos 7, 30, 90 dias, mês atual, trimestre, ano atual) **para** analisar tendências de crescimento.

**Critérios de Aceitação:**
- Filtro de período disponível no topo do dashboard
- Ao alterar o período, todos os indicadores são recalculados
- Gráfico de evolução da base reflete o período selecionado

### US-003 — Gráfico de Evolução da Base

**Como** Diretoria, **quero** visualizar um gráfico de evolução da base de clientes ao longo do tempo **para** acompanhar o crescimento do SaaS.

**Critérios de Aceitação:**
- Gráfico de linhas exibe a quantidade de novas contas por mês
- Gráfico permite alternar entre visão de contas totais e novas contas
- Período do gráfico segue o filtro aplicado no dashboard

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0001-0001-0001** | Métricas consideram apenas tenants com status diferente de "Excluído" (desativação lógica) |
| **RN-FEAT-EP-0001-0001-0002** | Período padrão do dashboard ao carregar: mês atual |
| **RN-FEAT-EP-0001-0001-0003** | Indicadores que exibem "zero" devem ser apresentados com o número 0, nunca em branco |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Dashboard funcional exibindo contas ativas, distribuição por plano e status de tenants | Ata de demo com PO |
| F2 | Filtro por período funcional em todos os indicadores | Checklist de períodos testados |
| F3 | Gráfico de evolução da base operacional | Print do gráfico com dados reais |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
