# SPRINT-REVIEW: Sprint 6 — Unidades de Negócio e Catálogo

- **Sprint:** 6 de 7
- **Data da Review:** 15/10/2026
- **Participantes:** Time Técnico, Tech Lead, **Product Owner** 🎯
- **Features:** 2 (F04-05, F04-06)
- **Status:** 🔄 Em Execução — Frente 0 concluída ✅ (4/4). 261 testes (0 failures). CnpjValidator com algoritmo alfanumérico (IN RFB 2.119/2022).
- **Débitos Técnicos:** [IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo.md](./IDENTIFIED-TECHNICAL-DEBT-sprint-06-bus-catalogo.md) — 22 débitos (4🔴→F0 ✅, 5🟡→F1, 8🔵→Sprint 7)
- **Execução F0:** [SPRINT-6-EXECUTION-REPORT-Frente-0.md](./SPRINT-6-EXECUTION-REPORT-Frente-0.md)

---

## ✅ Frente 0 — Correções Pré-Sprint (Concluída)

- [x] **BusinessUnit.java** reescrita com 16 campos (V001+V007) — `corporateName`, `taxRegime`, endereço completo, `status`
- [x] **ProductService.java** entity criada — 6 colunas de domínio, `ProductType` enum
- [x] **validateBusinessUnitTenant()** — bloqueia IDOR cross-tenant no `assignRole()` do PermissionService
- [x] **CnpjValidator** — algoritmo unificado numérico+alfanumérico (IN RFB 2.119/2022). 45 testes. 11 CNPJs reais validados (6 alfanuméricos)
- [x] **Code Review:** 7 skills, 12 findings, HIGH-1 DRY corrigido

## 🎯 O Que Demonstrar

### 1. Unidades de Negócio — Hierarquia (F04-05)

- [ ] **Criar Matriz:** CNPJ válido, parent_id=NULL
- [ ] **Criar Filial:** parent_id=Matriz.id → aparece na árvore hierárquica
- [ ] **Tree view:** GET /business-units retorna estrutura aninhada (Matriz > Filial A > Filial A-1)
- [ ] **CNPJ duplicado:** Tentar criar com mesmo CNPJ → 409 "CNPJ já cadastrado"
- [ ] **CNPJ imutável:** Tentar PATCH alterando CNPJ → 400 "CNPJ não pode ser alterado"
- [ ] **Desativar:** Filial desativada não aparece em seletores. Tentar criar nova filial com parent_id desativado → 422
- [ ] **Reúso CNPJ:** Soft delete da BU → criar nova BU com mesmo CNPJ → OK

> 🎬 **Script:** "Esta é a Matriz com CNPJ XX. Vou criar uma filial vinculada a ela. Agora vou criar uma sub-filial da filial. A árvore hierárquica é ilimitada. Vou desativar uma filial e tentar usá-la como 'pai' — bloqueado."

### 2. Catálogo de Produtos (F04-06)

- [ ] **Criar produto:** Vinculado à BU ativa no seletor. Nome, SKU (opcional), tipo (PRODUCT/SERVICE)
- [ ] **SKU duplicado:** Mesmo SKU na mesma BU → 409 "SKU já cadastrado para esta unidade"
- [ ] **SKU em BU diferente:** Mesmo SKU em BU diferente → OK
- [ ] **Indicador "Não mapeado":** Produto sem mapeamento fiscal exibe badge "Não mapeado"
- [ ] **Desativar/Reativar:** Produto desativado não aparece no catálogo. Reativar volta a exibir
- [ ] **Segmentação:** GET /products?business_unit_id=A → apenas produtos da BU-A

### 3. Isolamento Multi-Tenant (T-077) ⭐

- [ ] **Demonstrar:** Login como Tenant-A → vê BUs e produtos do Tenant-A
- [ ] **Demonstrar:** Login como Tenant-B → vê BUs e produtos do Tenant-B
- [ ] **Prova:** Tenant-A tenta acessar BU de Tenant-B por ID direto → 404
- [ ] **Admin FBSO:** Login como Admin FBSO → visão global de todos os tenants

> 🎬 **Script:** "Este é o teste mais importante da sprint. Tenant-A tem 3 BUs e 10 produtos. Tenant-B tem 2 BUs e 5 produtos. Como Tenant-A, só vejo os meus dados. Como Tenant-B, só vejo os dele. Nenhum vazamento."

---

## 📋 Pontos de Verificação (PO)

| Verificação | Status |
|:---|:---:|
| Criar BU Matriz (CNPJ) | ⬜ |
| Criar BU Filial (hierarquia) | ⬜ |
| Tree view hierárquica funcional | ⬜ |
| CNPJ duplicado → bloqueado | ⬜ |
| CNPJ não pode ser alterado | ⬜ |
| Soft delete + reúso CNPJ | ⬜ |
| BU desativada não pode ser "pai" | ⬜ |
| Criar produto vinculado à BU | ⬜ |
| SKU duplicado mesma BU → bloqueado | ⬜ |
| Indicador "Não mapeado" visível | ⬜ |
| Catálogo segmentado por BU | ⬜ |
| **Tenant-A não vê dados de Tenant-B** ⭐ | ⬜ |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## ➡️ Próximo Passo

**Sprint 7 — Homologação** (15/10 → 30/10): Testes de regressão, performance (p95 ≤ 3s), SAST scan, OpenAPI 37 endpoints, deploy staging → produção. **Sprint final — go-live em 30/10.**

---

🤖 *Checklist de review da Sprint 6. O isolamento multi-tenant é o destaque — zero vazamento entre tenants.*
