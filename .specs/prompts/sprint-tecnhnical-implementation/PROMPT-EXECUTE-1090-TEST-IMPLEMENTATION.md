# PROMPT-EXECUTE-1090-TEST-IMPLEMENTATION

## Contexto

Este prompt executa a **Fase de Implementação dos Testes** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 4 — passo 11). Implementa os testes do ciclo **imediatamente após o código**, seguindo o plano de testes, e valida a cobertura.

**Princípios fundamentais:**

1. **Testes imediatos:** escritos logo após cada código — nunca no fim do ciclo.
2. **Cobertura por nível:** unitários, integração e segurança conforme o plano (3.1/3.2/3.3).
3. **Meta objetiva:** cobertura ≥ 80% (ou a meta definida em ARCHITECTURE.md/TEST_PLAN.md).

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{CICLO_DIR}` | Pasta do ciclo | `.../sprints/sprint-01-setup/` |
| `{CICLO_NUMBER}` | Número do ciclo | `1` |
| `{CICLO_NAME}` | Nome curto do ciclo (kebab-case) | `sprint-01-setup` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-TEST-PLANNING.md  ← Mapeamento e estratégia (Fase 1040)
    ├── {CICLO_DIR}/SPRINT-TEST-SUITE.md                  ← Cenários aplicáveis ao ciclo
    ├── SPECS_DIR/TEST_PLAN.md                            ← IDs e níveis de teste
    └── SPECS_DIR/ARCHITECTURE.md                         ← Padrões de testes por camada
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-TEST-PLANNING.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1070-TEST-PLANNING`.

---

## Missão

Implementar os testes do ciclo `{CICLO_NUMBER} — {CICLO_NAME}` conforme o plano, executar unitários/integração/segurança e comprovar a cobertura ≥ 80%, registrando tudo no `PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md`.

---

## Fluxo de Execução

```
PARA cada task T-XXX:

    1. CONSULTAR SPRINT-TEST-SUITE.md — quais cenários se aplicam?
    2. MAPEAR para TEST_PLAN.md — confirmar IDs e níveis de teste

    3. IMPLEMENTAR os testes seguindo as práticas da stack:
       - Unitários (JUnit 5 + Mockito | pytest | Jest | Go testing)
         → lógica de negócio e validações
       - Integração (Testcontainers | Docker Compose | Go integration)
         → repositórios, acesso a dados e APIs
       - Segurança (se aplicável ao ciclo)
         → RBAC, isolamento multi-tenant, rate limiting

    4. EXECUTAR testes unitários — confirmar verde
    5. EXECUTAR testes de integração — confirmar verde
    6. VERIFICAR cobertura de código — meta ≥ 80%:
       · Java: `mvn jacoco:check`
       · Go: `go test -cover ./...`
       · JS/TS: `jest --coverage`
       · Python: `pytest --cov`
       · Rust: `cargo tarpaulin`
```

**Falhas nos testes** → encaminhar ao `PROMPT-EXECUTE-1110-FAILURE-HANDLING`.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md`:

```markdown
# PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md — Testes: Ciclo {N}
[Header: solução, projeto, ciclo, stack, data]
## 1. Resumo
- Testes implementados: X (unit Y | integration Z | security W)
- Cenários cobertos: N/N (SPRINT-TEST-SUITE.md)
## 2. Por Task
| Task | Cenários | Nível | Ferramenta | Status |
|:---|:---|:---|:---|:---:|
| T-XXX | TC-XXX-001 | Unit | [framework] | ✅ verde |
## 3. Evidências
- Comando: `[comando]` → [resultado]
- Total de testes: N | Passando: N | Falhando: 0
## 4. Cobertura
| Métrica | Valor | Meta | Status |
|:---|:---|:---|:---:|
| Linhas | XX% | ≥ 80% | ✅/❌ |
| Branches | XX% | ≥ 80% | ✅/❌ |
## 5. Desvios do Plano
[Desvios vs PACKAGE-DEVELOPMENT-TEST-PLANNING.md + justificativa]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| Skills de teste da stack | herdadas | `131-java-testing-unit-testing`, `golang-testing`, `javascript-typescript-jest`, `pytest-coverage` — práticas de escrita de testes por stack (tabela 6.3 do original) |
| `verification-before-completion` | automático | Testes verdes e cobertura comprovados antes de declarar a fase concluída |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. Testes seguem o plano aprovado (1070) — desvio registrado e justificado.
2. Nenhum teste é considerado pronto sem executar e passar.
3. Cobertura abaixo da meta = tratar antes de prosseguir (corrigir testes ou escalar ao humano).
4. Falha persistente → 1070 (impedimento + decisão humana).
5. IDs dos cenários (TEST_PLAN) preservados no registro.
