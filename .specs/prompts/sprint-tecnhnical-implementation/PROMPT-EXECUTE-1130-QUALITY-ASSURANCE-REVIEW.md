# PROMPT-EXECUTE-1130-QUALITY-ASSURANCE-REVIEW

## Contexto

Este prompt executa a **fase de Revisão de QA/Segurança (HITL)** do pacote de desenvolvimento — o gate humano de aceite do código do ciclo. Ele **delega a execução** ao prompt atual `PROMPT-QA-REVISOR-SECURITY` (mantido como executor) e registra o veredito no artefato da fase.

**Princípios fundamentais:**

1. **Gate humano obrigatório:** nenhum ciclo avança sem o veredito `[STATUS: APPROVED]`.
2. **Limite de 3 loops:** FAILED → correção no 1030/1080 → re-invocar o gate (máx. 3 tentativas).
3. **O revisor não corrige:** o `PROMPT-QA-REVISOR-SECURITY` só encontra falhas e documenta — a correção volta ao executor (1030).

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
| `{REFERENCIA}` | Arquivo de referência do gate (`TEST_PLAN.md` ou `SECURITY.md`) | `SPECS_DIR/TEST_PLAN.md` |
| `{LOOP_ATUAL}` | Número da tentativa do gate (1 a 3) | `1` |

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-IMPLEMENTATION.md  ← Código implementado (Fase 1030)
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-CODE-REVIEW.md     ← Ajustes do review (Fase 1080)
    └── {REFERENCIA}                                       ← TEST_PLAN.md ou SECURITY.md (critérios)
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-CODE-REVIEW.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1120-CODE-REVIEW`.

---

## Missão

Executar o gate humano de QA/segurança do ciclo `{CICLO_NUMBER} — {CICLO_NAME}` via `PROMPT-QA-REVISOR-SECURITY`, com até 3 loops, e registrar o veredito no `PACKAGE-DEVELOPMENT-QUALITY-ASSURANCE-REVIEW.md`.

---

## Fluxo de Execução

1. **Invocar `PROMPT-QA-REVISOR-SECURITY`** (executor mantido) com:
   - Código implementado: conteúdo do `PACKAGE-DEVELOPMENT-IMPLEMENTATION.md` + diff do ciclo
   - Arquivo de referência: `{REFERENCIA}`
   - Loop atual: `{LOOP_ATUAL}` de 3
2. **Interpretar o veredito:**
   - `[STATUS: APPROVED]` → registrar e prosseguir para a Fase 1090 (orquestrador).
   - `[STATUS: FAILED]` → o revisor gera o relatório de falhas (`FEEDBACK_ERRORS`): encaminhar as correções ao `PROMPT-EXECUTE-1040-IMPLEMENTATION` (código) ou ao 1080 (ajustes), conforme o tipo da falha, e **re-invocar este gate** com `{LOOP_ATUAL}` + 1.
3. **Estouro de loops:** após a 3ª tentativa com FAILED → encaminhar ao `PROMPT-EXECUTE-1110-FAILURE-HANDLING` (impedimento + decisão humana).

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-QUALITY-ASSURANCE-REVIEW.md`:

```markdown
# PACKAGE-DEVELOPMENT-QUALITY-ASSURANCE-REVIEW.md — Gate QA/Segurança: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Veredito
[✅ APPROVED no loop X | ❌ FAILED no loop X de 3]
## 2. Referência Utilizada
[{REFERENCIA} — TEST_PLAN.md ou SECURITY.md]
## 3. Falhas Encontradas (se FAILED)
| Loop | Erro | Diretriz violada | Encaminhamento |
|:---:|:---|:---|:---|
| 1 | [erro] | [regra do SECURITY/TEST_PLAN] | → 1030 |
## 4. Histórico de Loops
[1..3 com veredito e correções aplicadas entre loops]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Veredito registrado antes de declarar a fase concluída |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

> `PROMPT-QA-REVISOR-SECURITY` é um **prompt executor**, não skill — a fase delega a ele o trabalho de revisão.

---

## Regras de Ouro

1. Gate é HITL — nenhum ciclo avança sem `[STATUS: APPROVED]`.
2. Máximo de 3 loops; estouro → 1070 (impedimento).
3. O revisor NUNCA corrige código — falhas voltam ao executor (1030/1080).
4. Cada loop registrado no artefato com veredito e encaminhamento.
