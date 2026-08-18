# PROMPT-EXECUTE-1010-TECHNICAL-DEBT-AUDIT

## Contexto

Este prompt executa a **fase de Auditoria de Débito Técnico** do pacote de desenvolvimento. Roda em **dois modos**: (a) **pré-ciclo** — auditoria completa antes de iniciar a implementação (prevenção); (b) **catálogo** — recebimento dos `DT-XXX` gerados pelas fases 1070 (impedimentos) e 1080 (achados Medium/Low), mantendo o documento de débitos atualizado. **Delega a execução** ao prompt atual `PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT` (mantido como executor).

**Princípios fundamentais:**

1. **Prevenção, não correção tardia:** identificar débitos ANTES de iniciar o ciclo evita propagação.
2. **Código DT-XXX imutável:** cada débito recebe um código permanente — chave primária para rastreamento externo (Jira/Trello/GitHub).
3. **Decisão humana:** o agente NUNCA decide unilateralmente o que corrigir — o documento de débitos é apresentado ao time.

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
| `{MODE}` | `pre-ciclo` (auditoria completa) ou `catalogo` (receber DT-XXX das fases 1070/1080) | `pre-ciclo` |

## Documentos de Referência

```
Ler obrigatoriamente (modo pre-ciclo):
    ├── {CICLO_DIR}/SPRINT-CARD.md, SPRINT-TEST-SUITE.md, SPRINT-REVIEW.md
    ├── SPECS_DIR/PRD.md, SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md
    └── {SOLUTION_PATH}/src/ (código-fonte a escanear)

Ler no modo catálogo:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-FAILURE-HANDLING.md  (1070)
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-CODE-REVIEW.md       (1080)
    └── IDENTIFIED-TECHNICAL-DEBT-{CICLO_NAME}.md (documento de débitos)
```

---

## Missão

Auditar/atualizar o débito técnico do ciclo `{CICLO_NUMBER} — {CICLO_NAME}` via `PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT` no modo `{MODE}`, e registrar a operação no `PACKAGE-DEVELOPMENT-TECHNICAL-DEBT-AUDIT.md`.

---

## Fluxo de Execução

### Modo `pre-ciclo`

1. **Invocar `PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT`** (executor mantido) com os parâmetros do ciclo.
2. **Apresentar o documento** `IDENTIFIED-TECHNICAL-DEBT-{CICLO_NAME}.md` ao time para decisão explícita (tratar agora vs. depois — HITL).
3. **Registrar** no artefato da fase: débitos encontrados, decisão por débito e códigos `DT-XXX` aceitos para o ciclo.

### Modo `catalogo`

1. **Coletar** os `DT-XXX` gerados em 1070 (impedimentos) e 1080 (achados Medium/Low).
2. **Registrar** cada um no `IDENTIFIED-TECHNICAL-DEBT-{CICLO_NAME}.md` com rastreabilidade (origem, impacto, decisão humana).
3. **Registrar** no artefato da fase a lista de débitos incorporados.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-TECHNICAL-DEBT-AUDIT.md`:

```markdown
# PACKAGE-DEVELOPMENT-TECHNICAL-DEBT-AUDIT.md — Débito Técnico: Ciclo {N}
[Header: solução, projeto, ciclo, data, modo]
## 1. Modo
[pre-ciclo | catalogo]
## 2. Débitos Identificados / Incorporados
| DT-XXX | Descrição | Origem (skill/fase) | Impacto | Decisão |
|:---|:---|:---|:---|:---|
| DT-014 | Query N+1 em RelatorioService | performance-review (1080) | Médio | tratar no CICLO-03 |
## 3. Documento de Débitos
[Referência: IDENTIFIED-TECHNICAL-DEBT-{CICLO_NAME}.md — estado atualizado]
## 4. Decisões Humanas
[O que tratar agora vs. depois — registradas]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| Skills do executor | herdadas | `PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT` emprega suas 9 skills de auditoria (detalhadas no próprio prompt — ver seção Skills dele) |
| `verification-before-completion` | automático | Códigos DT-XXX e decisões conferidos antes de concluir |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. Código `DT-XXX` é imutável — nunca reutilizar ou renumerar.
2. Decisão de tratar débito é SEMPRE humana.
3. Modo `catalogo` nunca apaga histórico — incorpora novos débitos ao documento.
4. Toda incorporação rastreia origem (fase/skill) e decisão.
