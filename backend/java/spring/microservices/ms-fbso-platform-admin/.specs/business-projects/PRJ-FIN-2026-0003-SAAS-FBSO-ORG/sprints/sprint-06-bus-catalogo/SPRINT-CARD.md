# SPRINT-CARD: Sprint 6 — Unidades de Negócio e Catálogo

- **Sprint:** 6 de 7
- **Marco:** M6 (EP-04b)
- **Datas:** 30/09/2026 → 15/10/2026
- **Duração:** 11 dias úteis
- **Responsável:** Agente IA
- **Status:** 🔄 Em Execução — Frentes 0+1 concluídas ✅ (9/18 tasks, 50%). 288 testes (0 failures). V009 RLS product_service. ADR-L08 WITH RECURSIVE. RateLimitFilter trusted-proxy-ips. Próximo: M6 Features (9 tasks).
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) v3.8 · [SPECS.md](../../SPECS.md) v2.7 · [TEST_PLAN.md](../../TEST_PLAN.md) v3.3
- **Débitos Técnicos:** [IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo.md) — 22 débitos (4 críticos → Frente 0, 5 recomendados → Frente 1, 8 nits → Sprint 7)
- **Execução Frente 0:** [SPRINT-6-EXECUTION-REPORT-Frente-0.md](./SPRINT-6-EXECUTION-REPORT-Frente-0.md)

---

> 🚫 **BRANCH OBRIGATÓRIA:** Toda implementação deste sprint DEVE usar exclusivamente a branch `PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-06-bus-catalogo`. Antes de começar, execute:
> ```bash
> git checkout PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-06-bus-catalogo
> git branch --show-current  # deve exibir: PRJ-FIN-2026-0003-java-ms-fbso-platform-admin-sprint-06-bus-catalogo
> ```
> 📖 Detalhes completos: [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

## 🎯 Sprint Goal

**"Estrutura hierárquica de Unidades de Negócio (Matriz/Filial) com CNPJ único entre ativos. Soft delete libera CNPJ para reúso. Catálogo de Produtos/Serviços segmentado por BU com SKU único. Isolamento multi-tenant verificado: tenant-A não vê dados de tenant-B."**

---

## 📋 Frente 0 — Correções Pré-Sprint (Bloqueantes) ✅ | 23/07/2026

> **Status:** ✅ 4/4 concluídas. [Relatório de Execução](./SPRINT-6-EXECUTION-REPORT-Frente-0.md)

| ID | Tarefa | Débito | Est. | Critério DONE | Status |
|:---|:---|:---:|:---:|:---|:---:|
| **T-161.DT-126** | Reescrever `BusinessUnit.java` — alinhar 16 campos com schema V001+V007 | DT-126 | 2h | Entity compila. `toColumnMap()` com 16 colunas | ✅ |
| **T-162.DT-127** | Criar `ProductService.java` entity — 6 colunas de domínio | DT-127 | 1h | Entity compila. Pronta para ProductRepository | ✅ |
| **T-163.DT-128** | Implementar `validateBusinessUnitTenant()` no PermissionService | DT-128 | 1h | IDOR cross-tenant bloqueado. TenantIsolationException | ✅ |
| **T-164.DT-129** | Criar `CnpjValidator.java` — algoritmo unificado numérico+alfanumérico (IN RFB 2.119/2022) | DT-129 | 1.5h | 45 testes. 11 CNPJs reais validados. OnboardingService atualizado | ✅ |

**Entregáveis F0:** 3 arquivos criados (ProductService.java, CnpjValidator.java, CnpjValidatorTest.java) + 6 modificados. 261 testes (0 failures). Code review: 7 skills, 12 findings, HIGH-1 DRY corrigido.

---

## 📋 Frente 1 — Recomendados ✅ | 23/07/2026

> **Status:** ✅ 5/5 concluídas. [Relatório de Execução](./SPRINT-6-EXECUTION-REPORT-Frente-1.md)

| ID | Tarefa | Débito | Est. | Critério DONE | Status |
|:---|:---|:---:|:---:|:---|:---:|
| **T-165.DT-130** | Criar V009 migration: RLS em `product_service` (tenant_id + POLICY) | DT-130 | 1.5h | RLS filtra product_service por tenant | ✅ |
| **T-166.DT-131** | Remover `hierarchyType` de BusinessUnit.java | DT-131 | 15min | Entity sem referência a hierarchyType (já removido na F0) | ✅ |
| **T-167.DT-133** | Atualizar SPRINT-CARD.md — branch name, métricas, .bak cleanup | DT-133 | 30min | Concluído na sessão de docs | ✅ |
| **T-168.DT-134** | ADR-L08: Query hierárquica PostgreSQL WITH RECURSIVE + BusinessUnitRepository | DT-134 | 1h | ADR documentado. CTE funcional. 5 testes. | ✅ |
| **T-169.DT-137** | Externalizar `trusted-proxy-ips` no RateLimitFilter | DT-137 | 30min | IPs configuráveis por ambiente. 6 testes. | ✅ |

**Entregáveis F1:** 10 arquivos criados, 6 modificados. 288 testes (0 failures). ADR-L08 documentado. V009 fecha gap RLS.

---

## 📋 Sprint Backlog — M6 Features (EP-04b)

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---|:---:|:---|
| **T-069** | Entidade BusinessUnit (tenant_id, parent_id FK auto-relacionamento, cnpj, corporate_name, tax_regime, address, status) + Repository com queries hierárquicas | F04-05 | 1.5d | Estrutura recursiva. CNPJ índice parcial unique_cnpj_active. FindAll hierárquico |
| **T-070** | `BusinessUnitService`: CRUD com hierarquia. CNPJ único ativos (RN17-01). CNPJ imutável. Soft delete libera reúso. Unidade desativada não pode ser pai (RN17-02). Sem limite níveis (RN17-04) | F04-05 | 2d | Duplicado → 409. CNPJ alterado → 400. Soft delete libera. parent_id inativo → 422 |
| **T-071** | `BusinessUnitController`: CRUD `/api/v1/business-units` + `POST /{id}/deactivate`. `@RequiresPermission`. Seletor BU: listar BUs permitidas (RN17-05) | F04-05 | 1.5d | Admin todas. Gerente/Operador apenas autorizadas. Seletor funcional |
| **T-072** | Entidade ProductService (business_unit_id FK, name, sku, type PRODUCT/SERVICE, description, status) + Repository. SKU único por BU ativo (RN18-02, índice parcial) | F04-06 | 1d | SKU único por BU. Vinculação automática à BU ativa |
| **T-073** | `ProductService`: CRUD. Indicador "Não mapeado" (placeholder fiscal — RN18-03). Soft delete (RN18-04). SKU opcional, se informado único por BU. Catálogo segmentado (RN18-01) | F04-06 | 2d | CRUD funcional. "Não mapeado" no response. Soft delete preserva histórico |
| **T-074** | `ProductController`: CRUD `/api/v1/products` + `POST /{id}/deactivate`, `/activate`. `@RequiresPermission`. Filtro por BU ativa no seletor | F04-06 | 1d | Admin/Manager: create/edit/deactivate. Operator: view. Desativado não aparece |
| **T-075** | Testes unitários M6: `BusinessUnitService`, `ProductService`. RN17-01, RN17-04, RN18-02, RN18-04 | F04-05, F04-06 | 1.5d | ≥ 80%. Soft delete + reúso CNPJ |
| **T-076** | Testes integração M6: CRUD BU + Product com Testcontainers. CNPJ único, hierarquia, SKU único, soft delete | F04-05, F04-06 | 2d | PostgreSQL real. Cross-tenant query não vaza BU |
| **T-077** | Testes isolamento multi-tenant: tenant-A não vê dados tenant-B em BU ou Produto. 2 tenants + dados sobrepostos | F04-05, F04-06 | 1.5d | Zero dados cross-tenant |

**Total:** 9 tarefas (features) + 4 (Frente 0 ✅) + 5 (Frente 1 ⬜) = 18 tarefas · ~18 dias-homem

---

## 📦 Features Entregues

| Feature | Descrição | RNs Cobertas |
|:---|:---|:---|
| **F04-05** | Unidades de Negócio | RN17-01, RN17-02, RN17-03, RN17-04, RN17-05 |
| **F04-06** | Catálogo de Produtos/Serviços | RN18-01, RN18-02, RN18-03, RN18-04 |

---

## 📊 Estrutura Hierárquica (Referência)

```
Matriz (parent_id = NULL)
  ├── Filial A (parent_id = Matriz.id)
  │     └── Filial A-1 (parent_id = FilialA.id)
  └── Filial B (parent_id = Matriz.id)
```

- Sem limite de níveis (RN17-04)
- CNPJ único entre ativos do tenant (RN17-01)
- Soft delete libera CNPJ para reúso (RN17-01)
- Unidade desativada não pode ser "pai" (RN17-02)
- Seletor BU filtra por permissão do usuário (RN17-05)

---

## ✅ Definition of Done (Sprint-Level)

- [x] **Frente 0:** BusinessUnit.java alinhada com V001+V007 (16 campos) ✅
- [x] **Frente 0:** ProductService.java entity criada ✅
- [x] **Frente 0:** validateBusinessUnitTenant() — IDOR cross-tenant bloqueado ✅
- [x] **Frente 0:** CnpjValidator — algoritmo unificado CNPJ alfanumérico (IN RFB 2.119/2022) ✅
- [x] **Frente 1:** V009 migration — RLS em product_service (5ª tabela, gap fechado) ✅
- [x] **Frente 1:** ADR-L08 — BusinessUnitRepository.findTree() com WITH RECURSIVE ✅
- [x] **Frente 1:** RateLimitFilter — trusted-proxy-ips externalizado ✅
- [ ] CRUD BusinessUnit funcional com hierarquia recursiva
- [ ] CNPJ único entre ativos do mesmo tenant (409 se duplicado)
- [ ] CNPJ imutável após cadastro (400 se tentar alterar)
- [ ] Soft delete de BU libera CNPJ para reúso
- [ ] Unidade INACTIVE não pode ser `parent_id` (422)
- [ ] CRUD ProductService funcional com segmentação por BU
- [ ] SKU único por BU ativo (índice parcial)
- [ ] Indicador "Não mapeado" no response de produto
- [ ] **Teste de isolamento: tenant-A não vê dados de tenant-B em NENHUM endpoint**

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Query hierárquica recursiva com performance ruim (N+1) | Média | Médio | CTE recursiva no PostgreSQL. Alternativa: carregar flat + montar árvore em memória |
| Soft delete + reúso de CNPJ causar race condition | Baixa | Alto | Índice parcial `WHERE deleted_dt IS NULL` garante unicidade no banco |
| Isolamento multi-tenant falhar em queries manuais (nativeQuery) | Média | Crítico | T-070 testa cross-tenant para TODOS os endpoints. Proibir `@Query(nativeQuery=true)` sem `tenant_id` |

---

## 🔗 Dependências

- **Pré-requisitos:** Sprint 2 (TenantIsolationAspect). Sprint 3 (TenantRepository). Sprint 5 (primeira BU criada no onboarding).
- **Sucessor:** Sprint 7 (Homologação) — executa testes de regressão sobre todas as features.

---

## 📊 Métricas da Sprint

| Métrica | Meta | Atual |
|:---|:---:|:---:|
| Tasks completadas | 18/18 | 9/18 (50%) |
| Entidades | 2 (BusinessUnit, ProductService) | 2 entities criadas (F0) |
| Repositories | 1 (BusinessUnitRepository) | 1 criado (F1) |
| RNs implementadas | 9 | — |
| Cenários de teste | 17 | 45 CnpjValidator + 27 F1 = 72 |
| Isolamento multi-tenant | 100% endpoints | validateBusinessUnitTenant + RLS V009 ✅ |
| Testes totais | — | 288 (0 failures) |

---

🤖 *Gerado a partir de TASKS.md v3.9. Frentes 0+1 concluídas em 23/07/2026: 9/18 tasks (50%). V009 fecha gap RLS product_service. ADR-L08 documentado. Próximo: M6 Features (T-069 a T-077).*
