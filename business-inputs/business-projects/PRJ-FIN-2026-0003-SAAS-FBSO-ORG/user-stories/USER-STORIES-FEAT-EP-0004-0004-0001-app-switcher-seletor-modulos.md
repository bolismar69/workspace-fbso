# User Stories: App Switcher (Seletor de Módulos)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** F04-04 — App Switcher (Seletor de Módulos)
- **Épico:** EP-04 — Experiência do Cliente e Autoatendimento
- **Prioridade:** Must Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES.md](../04-FEATURES.md)

---

## Objetivo de Negócio

Fornecer o mecanismo de navegação entre módulos da plataforma, permitindo que o cliente alterne entre diferentes produtos sem sair do portal. É a peça central do modelo multi-produto da FBSO Platform.

---

## User Stories

### US-047 — Seletor de Módulos no Topo do Portal

**Como** Cliente, **quero** ver um seletor de aplicativos no topo do portal que exiba os módulos disponíveis no meu plano **para** navegar entre os produtos que contratei.

**Critérios de Aceitação:**
- App Switcher posicionado no canto superior, ao lado do logo da FBSO Platform
- Exibe apenas módulos que o usuário tem permissão para acessar
- Módulo ativo aparece destacado visualmente

### US-048 — Troca de Módulo com Adaptação do Menu

**Como** Cliente, **quero** que ao selecionar um módulo diferente no App Switcher, o menu lateral e o conteúdo da tela se adaptem imediatamente ao módulo escolhido.

**Critérios de Aceitação:**
- Troca de módulo atualiza menu lateral em tempo real
- Conteúdo da tela é redirecionado para o dashboard do módulo selecionado
- Transição fluida, sem recarregamento completo da página

### US-049 — Seletor de Módulo com Módulo Único (Placeholder)

**Como** Cliente com apenas um módulo contratado, **quero** que o Seletor de Módulo exiba o nome do meu módulo mesmo que eu não tenha outras opções, para que eu saiba em qual produto estou.

**Critérios de Aceitação:**
- Seletor de Módulo visível mesmo com um único módulo (na Fase 0, exibido como menu de navegação; expande para App Switcher visual quando houver 2+ produtos ativos)
- Exibe o nome do módulo ativo sem dropdown de seleção
- Exibe badge ou tooltip 'Mais módulos em breve' ao lado do nome do módulo ativo, indicando que novos produtos podem ser contratados futuramente

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN16-01** | Lista de módulos no App Switcher é determinada pela interseção entre: módulos do plano contratado e módulos que o usuário tem permissão |
| **RN16-02** | Na Fase 0, haverá um módulo placeholder chamado "FBSO Platform" visível para todos os clientes |
| **RN16-03** | A troca de módulo mantém o contexto da Unidade de Negócio selecionada |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | App Switcher visível no topo com módulos do plano do usuário | Print com diferentes combinações plano × permissão |
| F2 | Troca de módulo atualiza menu lateral sem recarregar a página | Teste de transição entre módulos |
| F3 | Placeholder "FBSO Platform" visível para cliente com 1 módulo | Print do App Switcher com módulo único |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*

---
👷 *Revisão técnica realizada pelo Agente: CaveMan em 15/07/2026, conforme User Story Review. Foram utilizados os skills: caveman-review.*
