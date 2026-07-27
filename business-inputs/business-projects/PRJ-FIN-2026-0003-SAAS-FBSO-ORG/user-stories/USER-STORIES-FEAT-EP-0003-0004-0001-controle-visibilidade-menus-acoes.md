# User Stories: Controle de Visibilidade de Menus e Ações

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0003-0004 — Controle de Visibilidade de Menus e Ações
- **Épico:** EP-0003 — Governança de Acessos e Permissões
- **Prioridade:** Must Have
- **Data-Alvo:** 15/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Garantir que a interface do portal se adapte dinamicamente às permissões do usuário, ocultando menus, botões e funcionalidades que ele não está autorizado a acessar.

---

## User Stories

### US-034 — Menu Lateral Dinâmico por Permissões

**Como** Usuário do Portal, **quero** que o menu lateral exiba apenas as opções correspondentes às minhas permissões (papel + módulo ativo) **para** ter uma interface limpa e focada no meu trabalho.

**Critérios de Aceitação:**
- Menu lateral renderizado dinamicamente conforme permissões
- Itens de menu sem permissão não aparecem (não ficam desabilitados — simplesmente não renderizam)
- Ao trocar de módulo no App Switcher, menu se adapta

### US-035 — Visibilidade Condicional de Botões de Ação

**Como** Usuário do Portal, **quero** que botões de ação (Criar, Editar, Excluir) apareçam apenas se eu tiver permissão para executar aquela ação.

**Critérios de Aceitação:**
- Botão "Novo" visível apenas para quem tem permissão de criação
- Botão "Editar" visível apenas para quem tem permissão de edição
- Botão "Excluir" visível apenas para quem tem permissão de exclusão

### US-036 — Tela de Acesso Negado para Rotas Não Autorizadas

**Como** Usuário do Portal, **quero** que ao tentar acessar uma área não permitida diretamente (via URL ou atalho), o sistema me redirecione para uma tela de "Acesso Negado" com explicação amigável.

**Critérios de Aceitação:**
- Tela de acesso negado com mensagem clara e não-técnica
- Exibe: "Você não tem permissão para acessar esta área. Se precisar de acesso, contate o administrador da sua conta."
- Não exibe detalhes técnicos ou caminhos internos
- Chamadas de API (fetch/AJAX) para endpoints não autorizados retornam HTTP 403 Forbidden com corpo JSON padronizado: {"error": "acesso_negado", "message": "Você não tem permissão para acessar este recurso."} — sem redirecionamento para página HTML.

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0003-0004-0001** | Ocultação de menu é a primeira camada (UX); o bloqueio por permissão no acesso direto é a camada de segurança — ambas devem ser implementadas |
| **RN-FEAT-EP-0003-0004-0002** | Usuários veem o nome do módulo ativo no topo do portal, ao lado do logo da FBSO Platform |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Menu lateral renderizado conforme permissões do usuário + módulo ativo | Print do menu para cada combinação papel × módulo |
| F2 | Botões de ação condicionais à permissão do usuário | Checklist de papel × botão verificado |
| F3 | Acesso direto a URL não autorizada redireciona para tela amigável | Tentativa de acesso a rota proibida |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
