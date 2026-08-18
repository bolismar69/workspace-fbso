# PROMPT-EXECUTE-1030-IMPLEMENTATION

## Contexto

Este prompt executa a **Fase de Implementação** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 2 — passo 9, incluindo o passo 4.1 de varredura CVE/SCA). Implementa cada task do ciclo na ordem do plano, com compilação e testes por task, e atualiza o status no SPRINT-CARD.md.

**Princípios fundamentais:**

1. **Ordem do plano:** segue estritamente o `PACKAGE-DEVELOPMENT-PLANNING.md` (§4).
2. **Stack por task:** confirmar a stack da task (PRD/SPECS/TASKS) e aplicar as skills registradas na pré-implementação.
3. **Build e teste por task:** compilar e testar antes de marcar concluída — sem exceção.
4. **Segurança contínua:** regras do SECURITY.md aplicadas em todo código novo; varredura CVE/SCA opcional pós-build quando 480/510 exigirem.

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
| `{TASK_IDS}` | IDs das tarefas a implementar (opcional — vazio = todas do plano) | `T-001,T-002` |

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-PLANNING.md   ← Ordem e abordagem por task (Fase 1020)
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md ← Stack e skills acionadas (Fase 1010)
    ├── {CICLO_DIR}/SPRINT-CARD.md                    ← Critérios DONE por task
    ├── SPECS_DIR/PRD.md, SPECS.md, TASKS.md          ← Stack e regras de negócio
    ├── SPECS_DIR/ARCHITECTURE.md                     ← Estrutura, padrões, ADRs
    ├── {SOLUTION_PATH}/.specs/security/SECURITY.md   ← Regras de segurança (se existir)
    └── {SOLUTION_PATH}/README.md                     ← Comandos de build/teste
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-PLANNING.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1020-DEVELOPMENT-PLANNING`.

---

## Missão

Implementar as tasks do ciclo `{CICLO_NUMBER} — {CICLO_NAME}` na ordem do plano, compilando e testando por task, e registrar o resultado no `PACKAGE-DEVELOPMENT-IMPLEMENTATION.md`.

---

## Fluxo de Execução

```
PARA cada task T-XXX, na ordem do PACKAGE-DEVELOPMENT-PLANNING.md (§4):

    1. LER o critério DONE no SPRINT-CARD.md

    2. IMPLEMENTAR seguindo ARCHITECTURE.md (estrutura, padrões, nomenclatura)

    3. IDENTIFICAR a(s) stack(s) da task:
       - Consultar PRD.md / SPECS.md / TASKS.md para confirmar a stack
       - Obter as skills registradas na pré-implementação (PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md §6)

       3.1 FALLBACK (somente se a task envolver componente NÃO coberto pela
           pré-implementação — reaplicar o protocolo do 1010, passos 6.1/6.3):

           ├── 6.1 (ordem de detecção da stack):
           │      1º {STACK} → 2º PRD.md → 3º SPECS.md → 4º ARCHITECTURE.md
           │      → 5º TASKS.md → 6º README.md
           │
           ├── 6.3 (seleção de skills para o componente novo):
           │      - Invocar `001-skills-inventory`
           │      - Cruzar o componente com as skills correspondentes
           │        (tabela de referência em PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md §Skills)
           │      - ⚠️ Dúvida → QUESTIONAR o humano:
           │        "Para a stack {stack_detectada}, identificamos os componentes
           │         X, Y, Z. Skills sugeridas: A, B, C. Confirma?"
           │
           └── 6.4 Registrar a decisão no PACKAGE-DEVELOPMENT-IMPLEMENTATION.md
                  (§2 da task — skills novas acionadas e por quê)

    3.1 APLICAR as melhores práticas das skills acionadas:
       - Padrões idiomáticos da linguagem
       - Convenções do framework
       - Padrões do ARCHITECTURE.md
       - Regras de segurança do SECURITY.md

    4. EXECUTAR a compilação/build:
       - Comando do README.md, ou inferido pelo gerenciador:
         · pom.xml → `mvn compile`
         · build.gradle → `gradle build`
         · go.mod → `go build ./...`
         · package.json → `npm run build`
         · Cargo.toml → `cargo build`
         · pyproject.toml → `python -m compileall .`
       - Corrigir erros de compilação antes de prosseguir

    4.1 (OPCIONAL — se o projeto exigir varredura de dependências, conforme
         480-SECURITY-DEFINITION e 510-TEST-STRATEGY): executar o
         `PROMPT-EXECUTE-CVE-SCA-SCAN` pós-build. Reprovado → corrigir ou
         catalogar débito DT-XXX antes de prosseguir.

    5. EXECUTAR o comando de teste do projeto:
       - README.md, ou inferido pelo gerenciador:
         · Maven → `mvn test`
         · Gradle → `gradle test`
         · Go → `go test ./...`
         · npm/pnpm/yarn → `npm test`
         · Cargo → `cargo test`
         · pytest → `pytest`
       - Verificar que testes existentes não quebraram

    6. MARCAR a task como concluída no SPRINT-CARD.md
```

**Falhas durante a implementação** → encaminhar ao `PROMPT-EXECUTE-1070-FAILURE-HANDLING` (não improvisar correções em loop infinito).

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-IMPLEMENTATION.md`:

```markdown
# PACKAGE-DEVELOPMENT-IMPLEMENTATION.md — Implementação: Ciclo {N}
[Header: solução, projeto, ciclo, stack, data]
## 1. Resumo
- Tasks implementadas: X/Y
- Ordem seguida: [referência ao plano §4]
## 2. Por Task
### T-XXX — [Nome]
- **Abordagem aplicada:** [do plano]
- **Arquivos criados/modificados:** [tabela: Ação | Arquivo | Descrição]
- **Build:** `[comando]` → [resultado]
- **Testes:** `[comando]` → [resultado]
- **CVE/SCA (se aplicável):** [resultado da varredura]
- **Status:** ✅ concluída / ❌ falha (→ 1070)
## 3. Desvios do Plano
[Desvios vs PACKAGE-DEVELOPMENT-PLANNING.md + justificativa]
## Rodapé
[Indicação de geração por IA, skills utilizados, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| Skills da stack | herdadas | Selecionadas na 1010 (`PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md` §6) — guiam a implementação idiomática de cada task |
| `ponytail` | full | Escada YAGNI de 7 rungs no código gerado |
| `security-review` | automático | Regras de segurança (OWASP Top 10) em todo código novo |
| `code-review` | automático | Revisão de qualidade do código gerado |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

> `PROMPT-EXECUTE-CVE-SCA-SCAN` (passo 4.1) é um **prompt**, não skill — invocado quando 480/510 exigirem varredura.

---

## Regras de Ouro

1. NUNCA implementar fora da ordem do plano (desvio = registrar e justificar).
2. NENHUMA task é concluída sem build e teste verdes.
3. Stack/skill em dúvida → perguntar ao humano.
4. Falha persistente → 1070 (impedimento + decisão humana), nunca loop infinito.
5. Código novo segue ARCHITECTURE.md + SECURITY.md sempre.
