# User Stories: Visão de Contas com Filtros

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** F01-02 — Visão de Contas com Filtros
- **Épico:** EP-01 — Portal Administrativo Interno
- **Prioridade:** Must Have
- **Data-Alvo:** 15/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES.md](../04-FEATURES.md)

---

## Objetivo de Negócio

Permitir que o time administrativo localize rapidamente qualquer conta de cliente e visualize suas informações principais.

---

## User Stories

### US-004 — Lista de Contas com Informações Resumidas

**Como** Administrador FBSO, **quero** visualizar a lista completa de contas de clientes com informações resumidas (nome, plano, status, data de criação) **para** navegar pela base.

**Critérios de Aceitação:**
- Tabela exibe: razão social, plano contratado, status, data de criação, data da última ação
- Lista ordenada por data de criação (mais recentes primeiro)
- Paginação a cada 25 registros
- Lista vazia (zero contas) exibe mensagem 'Nenhuma conta cadastrada' com link para criar primeira conta
- Erro de carregamento exibe mensagem com opção 'Tentar novamente'

### US-005 — Busca de Conta por Nome ou Razão Social

**Como** Administrador FBSO, **quero** buscar uma conta específica por nome ou razão social **para** localizar rapidamente um cliente.

**Critérios de Aceitação:**
- Campo de busca textual no topo da lista
- Busca com debounce de 300ms dispara após 3 caracteres digitados. Filtragem é case-insensitive e insensitive a acentos.
- Resultados exibem correspondências parciais (ex: "Super" encontra "Supermercado Bom Preço")
- Busca também permite filtrar por status do tenant e plano contratado via dropdowns complementares ao campo de texto

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN02-01** | Contas com status "Excluído" (soft delete) não aparecem na lista padrão |
| **RN02-02** | A busca não diferencia maiúsculas de minúsculas |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Lista de contas com paginação e ordenação correta | Print da lista com 25+ registros |
| F2 | Busca textual funcional com correspondência parcial | Teste com termo de 3 caracteres |
| F3 | Contas excluídas (soft delete) não visíveis na lista padrão | Verificação com tenant soft-deleted |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
