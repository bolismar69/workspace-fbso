# SPRINT-REVIEW: Sprint 2 — Segurança Cross-Cutting

- **Sprint:** 2 de 7
- **Data da Review:** 15/08/2026
- **Participantes:** Time Técnico, Tech Lead
- **PO:** Recomendado (entender os mecanismos de segurança)

---

> 🚫 **Branch:** `feature/java-fbso-platform-admin` ([PRD §8.4](../../PRD.md#84-branch-de-desenvolvimento))

## 🎯 O Que Demonstrar

### 1. Autenticação JWT (T-009, T-010)

- [ ] **Sem token:** `curl /api/v1/tenants` → 401
- [ ] **Token válido:** `curl -H "Authorization: Bearer <jwt>" /api/v1/tenants` → 200
- [ ] **Token expirado:** JWT com `exp` passado → 401
- [ ] **Token adulterado:** assinatura inválida → 401
- [ ] **Health check:** `curl /actuator/health` → 200 (sem token)

### 2. TenantContext + PostgreSQL RLS (T-011, T-015.1)

- [ ] Demonstrar log com `tenant_id` extraído do JWT
- [ ] Mostrar SQL gerado com `WHERE tenant_id = ?` no BaseRepository
- [ ] **Prova de isolamento (RLS):** query com tenant-A → 3 registros. Query com tenant-B → 2 registros. Nenhum vazamento.
- [ ] **Prova de bloqueio (RLS):** Tentar INSERT com tenant_id forjado → rejeitado pelo PostgreSQL
- [ ] `SHOW app.current_tenant_id` → UUID do tenant da sessão atual

### 3. RBAC — Bloqueio por Papel (T-013)

- [ ] Admin FBSO acessa `GET /tenants` → 200 ✅
- [ ] Operador tenta `POST /tenants` → 403 ❌
- [ ] Auditor tenta `PATCH /plans` → 403 ❌
- [ ] Mostrar resposta 403 amigável: `{"title":"Acesso negado","detail":"Você não tem permissão para executar esta operação.","status":403}`

### 4. Auditoria Assíncrona (T-014)

- [ ] Executar operação com `@Auditable` → sucesso imediato
- [ ] Consultar `audit_log` → registro presente (gravado async)
- [ ] Mostrar campos: previous_value, new_value, actor_id, timestamp

### 5. Tratamento de Erros (T-015)

- [ ] Provocar `BusinessException` → 422 RFC 7807
- [ ] Provocar `MethodArgumentNotValidException` → 400 com fields[]
- [ ] Provocar `Exception` genérica → 500 sem stack trace
- [ ] Verificar: todas as mensagens em PT-BR

---

## 📋 Pontos de Verificação

| Verificação | Status |
|:---|:---:|
| 401 sem token | ⬜ |
| 401 token inválido/expirado | ⬜ |
| 200 com token válido | ⬜ |
| TenantContext populado corretamente | ⬜ |
| TenantContext limpo após requisição | ⬜ |
| PostgreSQL RLS habilitado em 5 tabelas (V003) | ✅ |
| Politica tenant_isolation ativa em todas as tabelas com tenant_id | ✅ |
| TenantAwareDataSource configura app.current_tenant_id no pool HikariCP | ✅ |
| INSERT cross-tenant rejeitado pelo PostgreSQL | ✅ (estrutural — requer PostgreSQL para validacao real) |
| SELECT sem WHERE filtrado automaticamente pelo RLS | ✅ (estrutural — requer PostgreSQL para validacao real) |
| @RequiresPermission bloqueia (403) | ⬜ |
| @RequiresPermission permite (200/201) | ⬜ |
| 403 amigável (RFC 7807, PT-BR) | ⬜ |
| @Auditable grava async | ⬜ |
| GlobalExceptionHandler RFC 7807 | ⬜ |
| Zero stack traces em respostas HTTP | ⬜ |
| Zero dados pessoais em logs | ⬜ |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## ➡️ Próximo Passo

**Sprint 3 — Portal Admin + Contas/Planos** (15/08 → 31/08): Dashboard com métricas, CRUD de Tenants, Planos, Assinaturas. **Primeiro sprint com entrega visível para o PO.**

---


🤖 *Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): removido T-012 da seção RLS (substituído por T-015.1).*
