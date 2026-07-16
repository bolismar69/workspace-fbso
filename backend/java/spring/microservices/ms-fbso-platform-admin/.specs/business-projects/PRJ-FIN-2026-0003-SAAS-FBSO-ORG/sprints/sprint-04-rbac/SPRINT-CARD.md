# SPRINT-CARD: Sprint 4 — Governança de Acessos (RBAC)

- **Sprint:** 4 de 7
- **Marco:** M4 (EP-03)
- **Datas:** 31/08/2026 → 15/09/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [SPECS.md](../../SPECS.md)

---

> 🚫 **BRANCH OBRIGATÓRIA:** Toda implementação deste sprint DEVE usar exclusivamente a branch `feature/sprint-04-rbac`. Antes de começar, execute:
> ```bash
> git checkout feature/sprint-04-rbac
> git branch --show-current  # deve exibir: feature/sprint-04-rbac
> ```
> 📖 Detalhes completos: [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

## 🎯 Sprint Goal

**"4 papéis (Admin Tenant, Gerente BU, Operador BU, Auditor) aplicados com matriz de permissões RN10-01. Gestão de usuários com convite por e-mail. Vinculação Usuário × Unidade × Módulo. Bloqueio de acesso direto com 403 amigável em PT-BR."**

---

## 📋 Sprint Backlog

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---|:---:|:---|
| **T-039** | Entidade User + UserRepository. `findByEmailAndTenant`. Email único por tenant ativo (RN09-02, índice parcial) | F03-01 | 1d | UserRepository funcional. Soft delete respeitado |
| **T-040** | `UserService`: convite (email único RN09-02), desativar (não autodesativar RN09-03), reativar. Convite expira 7 dias (RN09-01) | F03-01 | 1.5d | Duplicado → 409. Autodesativar → 422 |
| **T-041** | `UserController`: CRUD `/api/v1/users` + `POST /{id}/deactivate`. `@RequiresPermission(USER, ...)` | F03-01 | 1d | Lista exibe nome, email, role, status, BUs vinculadas |
| **T-042** | Entidades ResourceAction + RoleResource. Seed data com matriz RN10-01 (4 roles × 7+ resources) | F03-02 | 1d | Seed carrega. Matriz completa. `findByRole` retorna recursos |
| **T-043** | Entidade UserPermission (user_id, business_unit_id, role). UNIQUE(user_id, business_unit_id) | F03-02, F03-03 | 1.5d | Tabela ponte com constraint. Admin tenant acesso implícito a todas BUs |
| **T-044** | `PermissionService`: atribuir/revogar, vincular BU, gerenciar módulos. Admin acesso implícito (RN11-01, RN11-02) | F03-02, F03-03 | 2d | Sem BU → acesso negado. Sem módulo → acesso negado. Efeito imediato (RN11-03) |
| **T-045** | `PermissionController`: `GET /users/{uid}/permissions`, `PUT /users/{uid}/permissions`. `@RequiresPermission(PERMISSION, ...)` | F03-02, F03-03 | 1d | GET retorna atuais. PUT atualiza vínculos. Auditoria registrada |
| **T-046** | Integrar `RbacAspect` com `RoleResource` do banco. Cache de matriz (TTL 5min) | F03-04 | 1.5d | RBAC funcional. Cache melhora performance. 403 (não 404 — RN12-01) |
| **T-047** | Garantir 403 padrão: `{"title":"Acesso negado","detail":"Você não tem permissão para acessar esta área.","status":403}` (RN12-02) | F03-04 | 0.5d | Resposta sempre nesse formato. PT-BR |
| **T-048** | Testes unitários M4: `UserService`, `PermissionService`. RN09-03, RN10-01, RN11-01, RN11-02 | F03-01 a F03-03 | 1.5d | ≥ 80%. Cada RN positivo+negativo |
| **T-049** | Testes segurança RBAC: cada papel × endpoint proibido → 403. Teste parametrizado com Testcontainers | F03-02, F03-04 | 2d | Matriz RN10-01 validada como teste parametrizado com 20+ combinações |

**Total:** 11 tarefas · ~15 dias-homem

---

## 📦 Features Entregues

| Feature | Descrição | RNs Cobertas |
|:---|:---|:---|
| **F03-01** | Gestão de Usuários | RN09-01, RN09-02, RN09-03 |
| **F03-02** | Matriz de Permissões RBAC | RN10-01 |
| **F03-03** | Vinculação Usuário × Unidade × Módulo | RN11-01, RN11-02, RN11-03 |
| **F03-04** | Acesso Condicional (403 Amigável) | RN12-01, RN12-02 |

---

## 📊 Matriz RN10-01 (Referência Rápida)

| Papel | Dashboard | Tenants | Plans | Users | Permissions | BUs | Products | Audit |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| **Admin Tenant** | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| **Gerente BU** | ❌ | ❌ | ❌ | ❌ | ❌ | ✅ edit | ✅ edit | ❌ |
| **Operador BU** | ❌ | ❌ | ❌ | ❌ | ❌ | ✅ view | ✅ view | ❌ |
| **Auditor** | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ✅ view |

---

## ✅ Definition of Done (Sprint-Level)

- [ ] CRUD Usuários funcional (convidar, desativar, reativar)
- [ ] Admin não pode desativar a si mesmo (RN09-03 → 422)
- [ ] Seed data da matriz RN10-01 carregada corretamente
- [ ] `@RequiresPermission` integrado com banco (não hardcoded)
- [ ] Teste parametrizado: 20+ combinações papel × endpoint proibido → 403
- [ ] Resposta 403 sempre no formato padrão (PT-BR, RFC 7807)
- [ ] Usuário sem BU vinculada → 403 (RN11-01)
- [ ] Usuário sem módulo → 403 (RN11-02)

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Cache de permissões servir dado stale | Média | Alto | TTL curto (5min). Invalidação na alteração de permissões |
| Seed data RBAC não carregar em todos os ambientes | Baixa | Alto | Migration V003 com seed data. Teste em dev e CI |
| Integração RbacAspect × banco causar N+1 queries | Média | Médio | Cache em memória (Caffeine). Query única `findByRole` |

---

## 🔗 Dependências

- **Pré-requisitos:** Sprint 2 (RbacAspect, @RequiresPermission). Sprint 3 (UserRepository).
- **Sucessor:** Sprint 5 (Portal Cliente) — depende de UserRepository e PermissionRepository.

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Tasks completadas | 11/11 |
| Combinações papel×endpoint testadas | 20+ |
| Cenários de teste RBAC | 19 |
| Resposta 403 padrão | 100% endpoints |

---

🤖 *Gerado a partir de TASKS.md v2.0. A matriz RN10-01 é o coração desta sprint — validar cada combinação.*
