# SPRINT-CARD: Sprint 6 — Unidades de Negócio e Catálogo

- **Sprint:** 6 de 7
- **Marco:** M6 (EP-04b)
- **Datas:** 30/09/2026 → 15/10/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [SPECS.md](../../SPECS.md)

---

> 🚫 **BRANCH OBRIGATÓRIA:** Toda implementação deste sprint DEVE usar exclusivamente a branch `feature/sprint-06-bus-catalogo`. Antes de começar, execute:
> ```bash
> git checkout feature/sprint-06-bus-catalogo
> git branch --show-current  # deve exibir: feature/sprint-06-bus-catalogo
> ```
> 📖 Detalhes completos: [PRD.md §8.4](../../PRD.md#84-estratégia-de-branching--uma-branch-por-sprint)

## 🎯 Sprint Goal

**"Estrutura hierárquica de Unidades de Negócio (Matriz/Filial) com CNPJ único entre ativos. Soft delete libera CNPJ para reúso. Catálogo de Produtos/Serviços segmentado por BU com SKU único. Isolamento multi-tenant verificado: tenant-A não vê dados de tenant-B."**

---

## 📋 Sprint Backlog

| ID | Tarefa | Feature | Est. | Critério DONE |
|:---|:---|:---|:---:|:---|
| **T-062** | Entidade BusinessUnit (tenant_id, parent_id FK auto-relacionamento, cnpj, corporate_name, tax_regime, address, status) + Repository com queries hierárquicas | F04-05 | 1.5d | Estrutura recursiva. CNPJ índice parcial unique_cnpj_active. FindAll hierárquico |
| **T-063** | `BusinessUnitService`: CRUD com hierarquia. CNPJ único ativos (RN17-01). CNPJ imutável. Soft delete libera reúso. Unidade desativada não pode ser pai (RN17-02). Sem limite níveis (RN17-04) | F04-05 | 2d | Duplicado → 409. CNPJ alterado → 400. Soft delete libera. parent_id inativo → 422 |
| **T-064** | `BusinessUnitController`: CRUD `/api/v1/business-units` + `POST /{id}/deactivate`. `@RequiresPermission`. Seletor BU: listar BUs permitidas (RN17-05) | F04-05 | 1.5d | Admin todas. Gerente/Operador apenas autorizadas. Seletor funcional |
| **T-065** | Entidade ProductService (business_unit_id FK, name, sku, type PRODUCT/SERVICE, description, status) + Repository. SKU único por BU ativo (RN18-02, índice parcial) | F04-06 | 1d | SKU único por BU. Vinculação automática à BU ativa |
| **T-066** | `ProductService`: CRUD. Indicador "Não mapeado" (placeholder fiscal — RN18-03). Soft delete (RN18-04). SKU opcional, se informado único por BU. Catálogo segmentado (RN18-01) | F04-06 | 2d | CRUD funcional. "Não mapeado" no response. Soft delete preserva histórico |
| **T-067** | `ProductController`: CRUD `/api/v1/products` + `POST /{id}/deactivate`, `/activate`. `@RequiresPermission`. Filtro por BU ativa no seletor | F04-06 | 1d | Admin/Manager: create/edit/deactivate. Operator: view. Desativado não aparece |
| **T-068** | Testes unitários M6: `BusinessUnitService`, `ProductService`. RN17-01, RN17-04, RN18-02, RN18-04 | F04-05, F04-06 | 1.5d | ≥ 80%. Soft delete + reúso CNPJ |
| **T-069** | Testes integração M6: CRUD BU + Product com Testcontainers. CNPJ único, hierarquia, SKU único, soft delete | F04-05, F04-06 | 2d | PostgreSQL real. Cross-tenant query não vaza BU |
| **T-070** | Testes isolamento multi-tenant: tenant-A não vê dados tenant-B em BU ou Produto. 2 tenants + dados sobrepostos | F04-05, F04-06 | 1.5d | Zero dados cross-tenant |

**Total:** 9 tarefas · ~14 dias-homem

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

| Métrica | Meta |
|:---|:---:|
| Tasks completadas | 9/9 |
| Entidades | 2 (BusinessUnit, ProductService) |
| RNs implementadas | 9 |
| Cenários de teste | 17 |
| Isolamento multi-tenant | 100% endpoints |

---

🤖 *Gerado a partir de TASKS.md v2.0. O isolamento multi-tenant (T-070) é o teste mais crítico desta sprint.*
