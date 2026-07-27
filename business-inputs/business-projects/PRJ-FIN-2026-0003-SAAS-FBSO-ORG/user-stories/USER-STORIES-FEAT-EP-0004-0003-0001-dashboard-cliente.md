# User Stories: Dashboard do Cliente

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** F04-03 — Dashboard do Cliente
- **Épico:** EP-04 — Experiência do Cliente e Autoatendimento
- **Prioridade:** Should Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES.md](../04-FEATURES.md)

> **Nota:** F04-03 é classificado como Should Have (bônus se tempo permitir). Caso não seja implementado, o destino pós-login/onboarding será o Seletor de Módulo (F04-04).

---

## Objetivo de Negócio

Oferecer ao cliente uma tela inicial com visão geral de sua conta e acessos rápidos às principais funcionalidades.

---

## User Stories

### US-045 — Dashboard com Informações Resumidas da Conta

**Como** Cliente, **quero** visualizar um dashboard com informações resumidas da minha conta: unidades de negócio ativas, total de produtos cadastrados e meu plano atual.

**Critérios de Aceitação:**
- Dashboard com cards informativos: Unidades Ativas, Produtos no Catálogo, Plano Contratado
- Cada card é clicável e leva à respectiva área de gestão
- Card 'Produtos no Catálogo' exibe '0' ou 'Em breve' até que F04-06 (Catálogo de Produtos) esteja implementada (data-alvo 15/10/2026).

### US-046 — Notificações e Lembretes no Dashboard

**Como** Cliente, **quero** ver notificações e lembretes relevantes no meu dashboard (ex: "Complete seu cadastro de produtos", "Convite de usuário pendente").

**Critérios de Aceitação:**
- Área de notificações visível no dashboard
- Cada notificação com link para a ação relacionada
- Notificações podem ser dispensadas pelo usuário

### Casos de Borda

- Dashboard sem dados (primeiro acesso pós-onboarding) exibe cards com valor zero e mensagens de orientação.
- Em caso de falha no carregamento, dashboard exibe mensagem 'Não foi possível carregar os dados. Tente novamente.' com botão de recarregar.


---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN15-01** | Dashboard do cliente adapta-se ao módulo ativo no Seletor de Módulo (App Switcher na fase multi-produto) (métricas diferentes por módulo) |
| **RN15-02** | Na Fase 0 (este projeto), haverá um dashboard genérico que será expandido quando os módulos forem ativados |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Cards informativos exibem dados corretos e são clicáveis | Print do dashboard com navegação para cada card |
| F2 | Notificações exibidas e dispensáveis | Teste de exibição e dispensa de notificação |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*

---
👷 *Revisão técnica realizada pelo Agente: CaveMan em 15/07/2026, conforme User Story Review. Foram utilizados os skills: caveman-review.*
