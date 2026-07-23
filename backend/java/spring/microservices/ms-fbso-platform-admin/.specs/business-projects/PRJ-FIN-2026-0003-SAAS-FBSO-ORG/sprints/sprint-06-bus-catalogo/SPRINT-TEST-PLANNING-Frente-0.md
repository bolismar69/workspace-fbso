# SPRINT-TEST-PLANNING-Frente-0.md — Plano de Testes: Sprint 6 — Frente 0

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 6 de 7 — sprint-06-bus-catalogo
- **Frente:** Frente 0 — Correções Pré-Sprint (Bloqueantes)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + JUnit 5 + Mockito
- **Data:** 23 de Julho de 2026
- **Origem:** [SPRINT-DEVELOPMENT-PLANNING-Frente-0.md](./SPRINT-DEVELOPMENT-PLANNING-Frente-0.md)

---

## 1. Visão Geral

- **Tasks implementadas:** 4 (T-161 a T-164)
- **Cenários de teste mapeados:** 9 (5 CnpjValidator + 2 PermissionService + 2 regressão)
- **Meta de cobertura:** Manter ≥ 80% linhas — as 4 tasks não adicionam nova lógica de negócio complexa
- **Ferramentas:** JUnit 5 + Mockito (unitários)

### Estratégia de Teste para Frente 0

A Frente 0 é uma correção pré-sprint — não adiciona features, apenas corrige débitos bloqueantes. A estratégia de teste é:

1. **CnpjValidator (T-164):** Única classe com lógica nova e complexa — teste unitário completo com CNPJs válidos e inválidos
2. **PermissionService.validateBusinessUnitTenant (T-163):** Teste unitário do método de validação
3. **BusinessUnit (T-161) + ProductService (T-162):** Entidades sem lógica — cobertas pelos testes de compilação e pelos testes de integração existentes
4. **Regressão:** Suite completa de testes existentes (227) deve continuar passando

---

## 2. Mapeamento Task → Cenários de Teste

| Task | Cenário(s) | Nível | Ferramenta | Status |
|:---|:---|:---|:---|:---:|
| T-164 | TC-F0-001: CNPJ válido com máscara → true | Unit | JUnit 5 | ⬜ |
| T-164 | TC-F0-002: CNPJ válido sem máscara → true | Unit | JUnit 5 | ⬜ |
| T-164 | TC-F0-003: CNPJ com dígitos verificadores errados → false | Unit | JUnit 5 | ⬜ |
| T-164 | TC-F0-004: CNPJ todos dígitos iguais → false | Unit | JUnit 5 | ⬜ |
| T-164 | TC-F0-005: CNPJ curto/longo/nulo → false | Unit | JUnit 5 | ⬜ |
| T-164 | TC-F0-006: format() e strip() utilitários | Unit | JUnit 5 | ⬜ |
| T-163 | TC-F0-007: validateBusinessUnitTenant — mesmo tenant → OK | Unit | JUnit 5 + Mockito | ⬜ |
| T-163 | TC-F0-008: validateBusinessUnitTenant — tenant diferente → TenantIsolationException | Unit | JUnit 5 + Mockito | ⬜ |
| T-161/162 | TC-F0-009: Regressão — 227 testes existentes passando | Unit+Int | Surefire+Failsafe | ⬜ |

---

## 3. Estratégia por Nível de Teste

### 3.1 Testes Unitários — CnpjValidator

- **Ferramenta:** JUnit 5 (sem Mockito — validator é estático e sem dependências)
- **Padrão:** AAA (Arrange-Act-Assert)
- **Localização:** `src/test/java/com/fbso/platform/admin/unit/utils/CnpjValidatorTest.java`
- **O que mockar:** Nada — validator é puro, sem dependências externas
- **CNPJs de teste:**
  - Válido com máscara: `"11.222.333/0001-81"`
  - Válido sem máscara: `"11222333000181"`
  - Dígitos errados: `"11222333000199"`
  - Todos dígitos iguais: `"00.000.000/0000-00"`
  - Curto: `"123"`
  - Nulo: `null`
  - Branco: `"   "`

### 3.2 Testes Unitários — PermissionService.validateBusinessUnitTenant

- **Ferramenta:** JUnit 5 + Mockito
- **Localização:** Estender `PermissionServiceTest.java` existente
- **O que mockar:** `JdbcTemplate` (para a query SQL)
- **Cenários:**
  1. BU pertence ao tenant → sem exceção
  2. BU de outro tenant → `TenantIsolationException`
  3. BU inexistente → `TenantIsolationException`

### 3.3 Testes de Regressão

- Suite completa: `./mvnw test`
- Meta: 227/228 passando (1 erro pré-existente DT-136 tolerado)
- OnboardingServiceTest deve passar com o novo CnpjValidator

---

## 4. Ordem de Execução dos Testes

1. **CnpjValidatorTest** (T-164) — rápido, sem dependências, feedback imediato
2. **PermissionServiceTest** (T-163) — requer Mockito, ~1s
3. **Suite de regressão** (`./mvnw test`) — verifica que nada quebrou

---

## 5. Comandos de Execução

- Unit — CnpjValidator: `./mvnw test -Dtest="CnpjValidatorTest"`
- Unit — PermissionService: `./mvnw test -Dtest="PermissionServiceTest"`
- Regressão completa: `./mvnw test`
- Coverage: `./mvnw jacoco:check` (meta: ≥ 80%)

---

## 6. Ações Manuais ou Externas

Nenhuma ação manual necessária — todos os testes são automatizados.

---

🤖 *Documento gerado em 23/07/2026. Fase 3 do PROMPT-EXECUTE-SPRINT-TASKS para Sprint 6 Frente 0.*
