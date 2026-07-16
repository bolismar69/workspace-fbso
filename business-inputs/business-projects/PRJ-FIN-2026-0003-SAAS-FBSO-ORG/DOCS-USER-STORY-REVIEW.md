# User Stories Review: 18 arquivos — EP-01 a EP-04
## Revisor: Caveman | Data: 15/07/2026 | Escopo: 05-USER-STORYS-*.md (18 arquivos)

---

## Sumário

18 arquivos, ~1,270 linhas, 58 user stories. Cobertura completa das 18 funcionalidades do Features v1.1. **3 problemas críticos cross-epic, 12 issues de alinhamento, ~40 edge cases ausentes.**

---

## 1. Problemas Críticos Cross-Épico

### 🔴 CRIT-01: RBAC — 4 papéis persistem nos User Stories

Charter v1.1, BRD v1.1, Epics v1.1, Features v1.1 definem **3 papéis MVP** (Admin, Gerente, Operador). Auditor = fase futura. User stories ignoram essa decisão:

| Arquivo | Linha | Violação |
|---------|-------|----------|
| `05-USER-STORYS-02-5` | 32 | US-023 ator "Auditor Interno" |
| `05-USER-STORYS-03-2` | 23 | US-027 lista "Admin, Gerente, Operador, Auditor" |
| `05-USER-STORYS-03-2` | 32 | US-028 descreve Auditor como papel ativo |
| `05-USER-STORYS-03-2` | 48-55 | US-030 — história COMPLETA de implementação do Auditor (3 ACs) |
| `05-USER-STORYS-03-2` | 80 | Feature AC F1 exige "4 papéis" |
| `05-USER-STORYS-03-2` | 63 | Tabela RN10-01 sem anotação `[Fase Futura]` |

`Fix:` US-030 marcar como "Won't Have (esta fase)" ou "Backlog — Fase Futura". US-027 e US-028: adicionar "(MVP: Admin, Gerente, Operador)" e nota "Auditor previsto para fase futura". US-023: trocar ator para "Administrador FBSO". RN10-01: adicionar `[Fase Futura]` na coluna Auditor.

---

### 🔴 CRIT-02: Atores não-RBAC como personas de US

User stories usam personas que **não são papéis RBAC** do sistema:

| Persona US | Onde | É papel RBAC? |
|-----------|------|---------------|
| Gestor de Produto | F02-03 (US-015, 016, 017) | ❌ Stakeholder no RACI, não papel de sistema |
| Líder Comercial | F02-04 (US-020), F01-01 (US-002) | ❌ Stakeholder no RACI, não papel de sistema |
| Diretoria | F01-01 (US-003) | ❌ Stakeholder, não papel de sistema |
| Auditor Interno | F02-05 (US-023) | ❌ Explicitamente fase futura |
| Administrador FBSO | Diversos | ⚠️ OK para EP-01 (portal interno). Para EP-02/03, equivalente a "Admin do Tenant"? |

`Fix:` Documentar que "Administrador FBSO" é perfil interno do time FBSO.ORG, distinto dos papéis RBAC do tenant. Para Gestor de Produto e Líder Comercial: ou criar papéis correspondentes no RBAC, ou reatribuir US a "Administrador FBSO" com nota de que a funcionalidade será usada por esses stakeholders.

---

### 🔴 CRIT-03: Backend enforcement de permissões sem user story

Epics v1.1 (linha 270) e Features v1.1 RN12-01 (linha 333) exigem: "não basta esconder o menu, o acesso deve ser barrado também se o usuário tentar acessar diretamente."

US-034 e US-035 cobrem apenas frontend (ocultação de menus/botões). US-036 cobre acesso via URL no navegador. **Nenhuma US cobre verificação de permissões em chamadas de API** (403 Forbidden em JSON, middleware/interceptor).

`Fix:` Expandir US-036 para incluir "Chamadas de API não autorizadas retornam HTTP 403 com mensagem JSON padronizada" ou criar US-036b específica para backend enforcement.

---

## 2. EP-01: Portal Administrativo Interno (3 arquivos)

### 05-USER-STORYS-01-1-DASHBOARD-METRICAS-OPERACIONAIS.md (74 linhas)
US-001, US-002, US-003

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 27-28 | **Taxa de conversão de onboarding ausente.** Charter D1 exige este indicador. Não listado em US-001. |
| 2 | 🟡 | 32 | Filtros "dia" e "trimestre" ausentes. EP-01 epic (linha 88) exige: "dia, semana, mês, trimestre, ano". US-002 lista apenas: "7, 30, 90 dias, mês atual, ano atual". |
| 3 | 🟡 | 44 | "Gráfico de linhas ou barras" — "ou" ambíguo. Definir qual. |
| 4 | 🔵 | 26 | "Indicadores atualizados em até 3 segundos" — "atualizados" ambíguo: latência de dados ou tempo de carregamento? |
| 5 | 🔵 | — | Sem AC para: estado vazio (zero contas), erro de carregamento, loading state, período sem dados. |

### 05-USER-STORYS-01-2-VISAO-CONTAS-FILTROS.md (64 linhas)
US-004, US-005

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 35-36 | "Busca filtra em tempo real (a partir de 3 caracteres)" — debounced? A cada tecla? |
| 2 | 🟡 | — | EP-01 epic (linha 89) menciona busca por status e plano. US-005 só cobre nome/razão social. |
| 3 | 🔵 | 37 | RN02-02 cobre case-insensitive mas não acentos (ex: "João" vs "Joao"). |
| 4 | 🔵 | — | Sem AC para: estado vazio, busca sem resultados, erro de carregamento, paginação durante busca. |

### 05-USER-STORYS-01-3-ALERTAS-INDICADORES-ATENCAO.md (62 linhas)
US-006, US-007

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 26 | "amarelo: atenção; vermelho: crítico" — sem regra de negócio mapeando situações para cores. |
| 2 | 🟡 | 36 | "(se registrado)" — sem fallback se motivo não registrado. |
| 3 | 🔵 | — | Sem AC para: estado sem alertas, atualização em tempo real, conflito de múltiplos status. |

---

## 3. EP-02: Gestão de Clientes e Assinaturas (5 arquivos)

### 05-USER-STORYS-02-1-CADASTRO-ATIVACAO-CONTAS-CLIENTES.md (84 linhas)
US-008, US-009, US-010, US-011

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 26 | "segmento" vago — free-text ou dropdown? Valores permitidos? |
| 2 | 🟡 | 37 | "renovável" — mesmo link ou novo link? US-011 gera novo link, conflito potencial. |
| 3 | 🔵 | — | Sem rate limiting no reenvio de convite (spam). Sem tratamento de email service down. |

### 05-USER-STORYS-02-2-GESTAO-STATUS-TENANT.md (73 linhas)
US-012, US-013, US-014

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 53 | RN05-01 não prevê transição "Pendente Onboarding → Inativo" nem "Suspenso → Inativo". |
| 2 | 🟡 | 35 | "Bloqueio efetivo em até 5 minutos" — como verificar SLA? |
| 3 | 🟡 | 36 | "próxima ação" — refresh, API call, ou WebSocket? |
| 4 | 🔵 | — | Sem AC para: transição inválida, confirmação antes de suspender, idempotência de suspensão. |

### 05-USER-STORYS-02-3-CONFIGURACAO-PLANOS-COMERCIAIS.md (83 linhas)
US-015, US-016, US-017, US-018

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 22,31 | Ver CRIT-02: "Gestor de Produto" não é papel RBAC. |
| 2 | 🟡 | 46 | US-017 diz "Não é possível excluir plano" mas RN06-01 implica que exclusão é possível sem clientes ativos. Contradição. |
| 3 | 🟡 | 37 | "Full Suite" hardcoded ou designável? Comportamento ao adicionar novos módulos? |
| 4 | 🔵 | — | Sem validação de preço (negativo, zero). Sem tratamento de zero recorrências selecionadas. |

### 05-USER-STORYS-02-4-VINCULACAO-GESTAO-ASSINATURAS.md (74 linhas)
US-019, US-020, US-021

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 31 | Ver CRIT-02: "Líder Comercial" não é papel RBAC. |
| 2 | 🔴 | 36 vs 56 | US-020 AC: "data de início = dia seguinte ao término" vs RN07-02: "não pode deixar tenant sem assinatura ativa". Contradição. |
| 3 | 🟡 | — | Estados de tenant vs assinatura não cross-referenciados. Tenant "Ativo" + Assinatura "Suspensa" = ? |
| 4 | 🔵 | — | Sem AC para: pro-rata upgrade/downgrade, upgrade para plano descontinuado. |

### 05-USER-STORYS-02-5-HISTORICO-AUDITORIA-ADMINISTRATIVA.md (63 linhas)
US-022, US-023

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 32 | Ver CRIT-01: US-023 ator "Auditor Interno" — papel de fase futura. |
| 2 | 🟡 | 35 | Filtro "tipo de ação" incompleto vs RN08-01. Faltam "edição" e "alteração de permissões". |
| 3 | 🟡 | 26 | "quando aplicável" para dados anteriores/novos — subjetivo. Enumerar ações que sempre registram diff. |
| 4 | 🔵 | — | Sem retenção/purge policy. Sem filtro por administrador responsável. Sem consideração GDPR. |

---

## 4. EP-03: Governança de Acessos e Permissões (4 arquivos)

### 05-USER-STORYS-03-1-CADASTRO-CONVITE-USUARIOS.md (75 linhas)
US-024, US-025, US-026

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 23 vs 25 | US menciona "perfil de acesso" mas ACs do formulário só listam nome e e-mail. |
| 2 | 🟡 | 43 | "Reativação possível a qualquer momento" — sem restrição de quem pode reativar. |
| 3 | 🔵 | — | Sem AC para: convite expirado, paginação na lista, preservação de vinculações na desativação. |

### 05-USER-STORYS-03-2-DEFINICAO-PAPEIS-PERMISSOES-RBAC.md (90 linhas)
US-027, US-028, US-029, US-030

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 23,32,48,80 | Ver CRIT-01: 4 papéis em todo o arquivo. US-030 implementação completa do Auditor. |
| 2 | 🔴 | 63 | RN10-01 coluna Auditor sem `[Fase Futura]`. |
| 3 | 🟡 | — | Sem proteção contra rebaixamento do último Admin (admin lockout). |
| 4 | 🔵 | — | Sem AC para bootstrap do primeiro admin do tenant. |

### 05-USER-STORYS-03-3-VINCULACAO-USUARIO-UNIDADE-MODULO.md (74 linhas)
US-031, US-032, US-033

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 44 | "efeito na próxima ação do usuário" — vago: próximo login, request, clique? |
| 2 | 🟡 | 37 | US-032 só menciona frontend (não vê no App Switcher). Sem AC de bloqueio backend (CRIT-03). |
| 3 | 🔵 | — | Sem AC para: remoção de módulo do plano afetando usuários, concorrência, operações em lote. |

### 05-USER-STORYS-03-4-CONTROLE-VISIBILIDADE-MENUS-ACOES.md (73 linhas)
US-034, US-035, US-036

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 41-46 | US-036 cobre URL no navegador, **não cobre chamadas de API** (fetch/AJAX). Ver CRIT-03. |
| 2 | 🟡 | 35-37 | Só cobre botões Criar/Editar/Excluir. Ações como Desativar, Reativar, Aprovar não cobertas. |
| 3 | 🔵 | 55 | RN12-02 sem US correspondente. |

---

## 5. EP-04: Experiência do Cliente (6 arquivos)

### 05-USER-STORYS-04-1-AUTENTICACAO-RECUPERACAO-SENHA.md (75 linhas)
US-037, US-038, US-039

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 26 | Mensagem de erro genérica não especificada. E para conta bloqueada? |
| 2 | 🟡 | 47 | "Administrador do tenant pode desbloquear" — onde? Que tela? |
| 3 | 🔵 | — | Sem AC para: login com tenant suspenso, usuário desativado, sessão expirada, sessões concorrentes. |

### 05-USER-STORYS-04-2-ONBOARDING-GUIADO-PRIMEIRO-ACESSO.md (95 linhas)
US-040, US-041, US-042, US-043, US-044

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 64 | US-044 redireciona para dashboard (F04-03), que é **Should Have / bônus**. Se F04-03 não implementado, sem destino. |
| 2 | 🟡 | 47 | Validação de CNPJ — apenas dígitos ou consulta à Receita Federal? |
| 3 | 🟡 | 28 | "Não é possível pular etapas obrigatórias" — todas são obrigatórias? |
| 4 | 🔵 | — | Sem AC para: abandono/retomada, navegação entre passos, falha de rede durante salvamento. |

### 05-USER-STORYS-04-3-DASHBOARD-CLIENTE.md (62 linhas)
US-045, US-046

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 44 | RN15-01 usa "App Switcher" — deveria ser "Seletor de Módulo" (Fase 0). |
| 2 | 🟡 | 26 | Card "Produtos no Catálogo" antes do catálogo existir (30/09 vs 15/10). |
| 3 | 🔵 | — | Sem AC para: estado vazio, loading, erro, F04-03 não implementado. |

### 05-USER-STORYS-04-4-APP-SWITCHER-SELETOR-MODULOS.md (74 linhas)
US-047, US-048, US-049

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🔴 | 40,44 | US-049: "App Switcher" deveria ser "Seletor de Módulo". Features v1.1 já revisou — US file desatualizado. |
| 2 | 🟡 | 44 | Falta parenteses: "(na Fase 0, menu de navegação; expande para App Switcher visual quando houver 2+ produtos)". |
| 3 | 🔵 | 46 | "Indica visualmente que novos módulos podem ser adicionados" — como? Ícone? Tooltip? |

### 05-USER-STORYS-04-5-GESTAO-UNIDADES-NEGOCIO.md (96 linhas)
US-050, US-051, US-052, US-053, US-054

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 63 | "ao lado do App Switcher" — deveria ser "ao lado do Seletor de Módulo". |
| 2 | 🟡 | 26 | "cards ou lista" — ambíguo. |
| 3 | 🟡 | — | Sem validação contra ciclos hierárquicos (A → B → A). |
| 4 | 🔵 | — | Sem AC para: desativar matriz, filhas de unidade desativada, CNPJ de unidade já desativada. |

### 05-USER-STORYS-04-6-CATALOGO-PRODUTOS-SERVICOS.md (85 linhas)
US-055, US-056, US-057, US-058

| # | Severidade | Linha | Problema |
|---|-----------|------|----------|
| 1 | 🟡 | 45 | "Todos os campos do cadastro são editáveis" — inclusive tipo (Produto/Serviço)? Implicações para mapeamento fiscal futuro. |
| 2 | 🟡 | 28 | "Unidade de Negócio ativa no seletor" — qual seletor? Explicitar: Seletor de Unidade de Negócio (US-054). |
| 3 | 🔵 | — | Sem AC para: estado vazio, busca sem resultados, SKU duplicado, reativação de produto. |

---

## 6. Problemas Sistêmicos (Todos os Épicos)

### 6.1 Estados de erro, vazio e loading — Ausentes

**Nenhum** dos 18 arquivos inclui ACs sistemáticos para:
- **Empty state:** O que o usuário vê com zero registros (dashboard sem contas, catálogo vazio, lista de usuários vazia)
- **Error state:** Comportamento com falha de API (timeout, 500, rede indisponível)
- **Loading state:** Skeleton, spinner, ou indicador de carregamento

### 6.2 Auditoria sem especificação padronizada

Múltiplos arquivos mencionam "registro de auditoria" (US-022, US-027, US-033) mas sem padronização dos campos mínimos: usuário responsável, data/hora, ação, valores anteriores, valores novos.

### 6.3 Dependências entre features não documentadas

| Dependência | Risco |
|-------------|-------|
| US-044 (boas-vindas) → F04-03 (dashboard) | ALTO — dashboard é Should Have |
| US-045 (card produtos) → F04-06 (catálogo) | MÉDIO — datas diferentes (30/09 vs 15/10) |
| US-041 (dados pre-preenchidos) → F02-01 (cadastro tenant) | BAIXO — já entregue em M3 |

### 6.4 Machine de estados cross-entidade não definida

| Entidade | Estados |
|----------|---------|
| Tenant | Pendente Onboarding, Ativo, Suspenso, Inativo |
| Assinatura | Ativa, Suspensa, Cancelada |
| Usuário | Ativo, Inativo, Convite Pendente |

Não há AC que defina comportamento quando estados de entidades diferentes conflitam (ex: Tenant Ativo + Assinatura Suspensa).

---

## 7. Matriz de Conformidade

| Feature | Arquivo US | US count OK? | RBAC OK? | Terminologia OK? | Edge cases |
|---------|-----------|-------------|----------|-----------------|------------|
| F01-01 | 01-1 | ✅ 3 | ⚠️ personas internas | ✅ | ❌ 6 ausentes |
| F01-02 | 01-2 | ✅ 2 | ⚠️ personas internas | ✅ | ❌ 6 ausentes |
| F01-03 | 01-3 | ✅ 2 | ⚠️ personas internas | ✅ | ❌ 5 ausentes |
| F02-01 | 02-1 | ✅ 4 | ✅ Admin FBSO | ✅ | ❌ 4 ausentes |
| F02-02 | 02-2 | ✅ 3 | ✅ Admin FBSO | ✅ | ❌ 5 ausentes |
| F02-03 | 02-3 | ✅ 4 | 🔴 Gestor de Produto | ✅ | ❌ 4 ausentes |
| F02-04 | 02-4 | ✅ 3 | 🔴 Líder Comercial | ✅ | ❌ 4 ausentes |
| F02-05 | 02-5 | ✅ 2 | 🔴 Auditor Interno | ✅ | ❌ 4 ausentes |
| F03-01 | 03-1 | ✅ 3 | ✅ Admin Tenant | ✅ | ❌ 4 ausentes |
| F03-02 | 03-2 | ✅ 4 | 🔴 4 papéis + US Auditor | ✅ | ❌ 3 ausentes |
| F03-03 | 03-3 | ✅ 3 | ✅ Admin Tenant | ✅ | ❌ 4 ausentes |
| F03-04 | 03-4 | ✅ 3 | ✅ | ✅ | ❌ 4 ausentes |
| F04-01 | 04-1 | ✅ 3 | ✅ | ✅ | ❌ 4 ausentes |
| F04-02 | 04-2 | ✅ 5 | ✅ | ✅ | ❌ 5 ausentes |
| F04-03 | 04-3 | ✅ 2 | ✅ | 🟡 App Switcher | ❌ 4 ausentes |
| F04-04 | 04-4 | ✅ 3 | ✅ | 🔴 App Switcher | ❌ 4 ausentes |
| F04-05 | 04-5 | ✅ 5 | ✅ | 🟡 App Switcher | ❌ 5 ausentes |
| F04-06 | 04-6 | ✅ 4 | ✅ | ⚠️ ref vaga | ❌ 5 ausentes |

---

## 8. Ações Recomendadas (Priorizadas)

### Bloco A — Crítico (afeta escopo MVP)

1. **A-1: Remover/Migrar US-030 (Auditor) para fase futura.** Mover para backlog com nota "Won't Have (esta fase)". Atualizar US-027 e US-028 para 3 papéis MVP + nota Auditor futuro.
2. **A-2: Corrigir atores não-RBAC.** US-023 (Auditor Interno → Administrador FBSO). US-015/016/017 (Gestor de Produto → Administrador FBSO com nota). US-020 (Líder Comercial → Administrador FBSO com nota).
3. **A-3: Adicionar backend enforcement US.** Expandir US-036 para cobrir verificação de permissões em chamadas de API (HTTP 403 + JSON).
4. **A-4: Resolver dependência US-044 → F04-03.** Definir destino fallback se dashboard não implementado. Sugestão: tela de boas-vindas estática ou Seletor de Módulo.

### Bloco B — Alto (desalinhamentos com docs revisados)

5. **B-1: Adicionar "taxa de conversão de onboarding"** na US-001. Charter D1 exige este indicador.
6. **B-2: Corrigir terminologia App Switcher → Seletor de Módulo** em US-049 (04-4), US-054 (04-5), RN15-01 (04-3).
7. **B-3: Completar filtros de período** na US-002. Adicionar "dia" e "trimestre".
8. **B-4: Adicionar RN10-01 `[Fase Futura]`** na coluna Auditor em 03-2.
9. **B-5: Resolver contradição US-020 vs RN07-02** (gap de 1 dia no upgrade).
10. **B-6: Resolver contradição US-017 vs RN06-01** (exclusão de plano).

### Bloco C — Médio (edge cases e clareza)

11. **C-1: Adicionar ACs de estado vazio** em todas as US com listas/cards (US-001, 004, 006, 012, 019, 025, 045, 050, 056).
12. **C-2: Adicionar ACs de erro e loading** em todas as US com carregamento remoto.
13. **C-3: Definir máquina de estados cross-entidade** (tenant × assinatura). Adicionar RN.
14. **C-4: Adicionar proteção contra admin lockout** (RN: "Não é permitido rebaixar/desativar o último Admin do Tenant").
15. **C-5: Padronizar campos de auditoria** em todas as US que mencionam "registro de auditoria".
16. **C-6: Esclarecer validação de CNPJ** (dígitos vs Receita Federal) em US-042 e US-051.
17. **C-7: Completar filtros de auditoria** em US-023 (adicionar "edição" e "alteração de permissões").

---

## Resumo: Nota Caveman

| Dimensão | Nota | Comentário |
|----------|------|------------|
| Cobertura (58 US × 18 features) | ★★★★★ | Completa. Todas features cobertas. |
| Alinhamento RBAC com docs v1.1 | ★★☆☆☆ | 6 arquivos com 4 papéis ou atores inválidos. |
| Terminologia (App Switcher) | ★★★☆☆ | 3 arquivos desatualizados vs Features v1.1. |
| Edge cases | ★★☆☆☆ | ~40 ausentes. Estados vazio/erro/loading sistematicamente omitidos. |
| Consistência interna | ★★★☆☆ | 2 contradições cross-US (RN06-01, RN07-02). |
| Segurança (backend enforcement) | ★★☆☆☆ | Sem US para middleware de permissões. |

**Nota geral: 2.8/5** — User stories cobrem o escopo, mas falham em 3 dimensões críticas: alinhamento com decisão RBAC (3 papéis), ausência de backend enforcement, e omissão sistemática de edge cases. As 18 ações recomendadas são pré-requisito para início de desenvolvimento.

---

> Revisão caveman: 18 arquivos, 1,270 linhas, 58 US. 3 CRIT, 12 HIGH, 7 MED. 17 ações priorizadas.
