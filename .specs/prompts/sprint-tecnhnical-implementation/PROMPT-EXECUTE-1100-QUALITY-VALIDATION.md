# PROMPT-EXECUTE-1100-QUALITY-VALIDATION

## Contexto

Este prompt executa a **Fase de Validação de Qualidade** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 5 — passos 12 e 13). Executa as verificações estáticas e de estilo e a suíte completa de testes do ciclo, garantindo zero violações e cobertura na meta.

**Princípios fundamentais:**

1. **Zero tolerância:** zero warnings, zero violations de lint/estilo.
2. **Suíte completa:** unit + integration rodando juntos, cobertura ≥ 80% (meta padrão; ajustar se ARCHITECTURE.md/TEST_PLAN.md definirem outra).
3. **Evidência no artefato:** todo comando e resultado registrado.

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
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md ← Testes implementados (Fase 1050)
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-TEST-PLANNING.md       ← Comandos e metas (Fase 1040)
    └── SPECS_DIR/ARCHITECTURE.md + TEST_PLAN.md               ← Meta de cobertura e padrões
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1090-TEST-IMPLEMENTATION`.

---

## Missão

Validar a qualidade do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`: linter/formatador com zero violações e suíte completa verde com cobertura na meta, registrando evidências no `PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md`.

---

## Fluxo de Execução

1. **Executar verificações estáticas e de estilo** (linter/formatador da stack):

```
· Java:   mvn checkstyle:check pmd:check
· Go:     go vet ./... && golangci-lint run
· JS/TS:  eslint . && prettier --check .
· Python: ruff check . && mypy .
· Rust:   cargo clippy && cargo fmt --check
```

   **Zero warnings. Zero violations.**

2. **Executar a suíte completa do ciclo:**

```
- Comando de teste completo (unit + integration)
- Cobertura ≥ 80% (meta padrão; ajustar se ARCHITECTURE.md ou TEST_PLAN.md definirem meta diferente)
```

3. **Tratar falhas:** qualquer violação ou teste quebrado → encaminhar ao `PROMPT-EXECUTE-1110-FAILURE-HANDLING` (auto-correção máx. 3 tentativas, depois impedimento).
4. **Gerar o artefato de saída** com as evidências.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md`:

```markdown
# PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md — Qualidade: Ciclo {N}
[Header: solução, projeto, ciclo, stack, data]
## 1. Verificações Estáticas
| Ferramenta | Comando | Resultado | Violações |
|:---|:---|:---|:---:|
| Checkstyle/PMD (Java) | `mvn checkstyle:check pmd:check` | ✅ | 0 |
| eslint + prettier (JS/TS) | `eslint . && prettier --check .` | ✅ | 0 |
| ... | ... | ... | ... |
## 2. Suíte Completa
- Comando: `[comando]` → [resultado]
- Testes: N/N passando
- Cobertura: linhas XX% | branches XX% (meta: ≥ 80%)
## 3. Violações Corrigidas (se houve)
[Tabela: violação | arquivo | correção | revalidação]
## 4. Veredito
[✅ QUALIDADE APROVADA | ❌ REPROVADA (→ 1070)]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Validar zero violações e suíte verde antes do veredito |
| `code-review` | automático | Transversal — qualidade do código validado |
| `security-review` | automático | Transversal — regras de segurança validadas |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

> Lint/formatador e suíte são ferramentas CLI da stack (não skills) — comandos na seção Fluxo.

---

## Regras de Ouro

1. Zero warnings e zero violations — não existe "violação aceitável".
2. Meta de cobertura vem de ARCHITECTURE.md/TEST_PLAN.md (default ≥ 80%) — nunca reduzir sem aprovação humana.
3. Falha persistente → 1070, nunca mascarar ou pular verificação.
4. Toda evidência registrada com comando e resultado reais.
