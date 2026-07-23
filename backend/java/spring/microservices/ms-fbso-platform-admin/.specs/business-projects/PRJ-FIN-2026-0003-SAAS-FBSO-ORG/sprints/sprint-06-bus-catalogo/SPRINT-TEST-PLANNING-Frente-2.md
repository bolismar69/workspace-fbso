# SPRINT-TEST-PLANNING-Frente-2.md — Plano de Testes: Sprint 6 — Frente 2

- **Solução:** `ms-fbso-platform-admin`
- **Sprint:** 6 — sprint-06-bus-catalogo
- **Frente:** Frente 2 — M6 Features
- **Data:** 23 de Julho de 2026

## 1. Visão Geral
- **Tasks implementadas:** T-069 a T-074 (6 tasks)
- **Cenários mapeados:** 20 (12 unit + 5 integration + 3 security)
- **Meta:** ≥ 80% cobertura. 288→~310 testes.

## 2. Mapeamento Task → Cenários

| Task | ID | Descrição | Nível |
|:---|:---|:---|:---|
| T-070 | TC-F2-070-001 | create com CNPJ válido → 201 | Unit |
| T-070 | TC-F2-070-002 | create com CNPJ duplicado → DuplicateCnpjException | Unit |
| T-070 | TC-F2-070-003 | create com CNPJ inválido → BusinessException | Unit |
| T-070 | TC-F2-070-004 | create com parentId inativo → BusinessException | Unit |
| T-070 | TC-F2-070-005 | update não altera CNPJ | Unit |
| T-070 | TC-F2-070-006 | deactivate (soft delete) | Unit |
| T-073 | TC-F2-073-001 | create produto com SKU único → OK | Unit |
| T-073 | TC-F2-073-002 | create produto com SKU duplicado → BusinessException | Unit |
| T-073 | TC-F2-073-003 | create produto sem SKU → OK | Unit |
| T-073 | TC-F2-073-004 | fiscalMappingStatus = NOT_MAPPED | Unit |
| T-073 | TC-F2-073-005 | deactivate produto (soft delete) | Unit |
| T-069 | TC-F2-069-001 | existsByCnpj retorna true/false | Unit |

## 3. Comandos
- Unit: `./mvnw test`
- Coverage: `./mvnw jacoco:check`

🤖 *Gerado em 23/07/2026.*
