# PROMPT-EXECUTE-1110-ARTIFACTS-UPDATE

## Contexto

Este prompt executa a **Fase de Atualização de Artefatos** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 10 — passo 25, OBRIGATÓRIA). Atualiza os artefatos do ciclo e os documentos-mestre do projeto com o resultado real da execução — a última fase do pacote de desenvolvimento.

**Princípios fundamentais:**

1. **Documentos-mestre são a fonte da verdade** — artefatos de ciclo são derivados; em conflito, os mestres prevalecem.
2. **Atualizar versão e data** no header de TODO documento modificado.
3. **Não apagar histórico** — checkboxes marcados (✅/❌), nunca remover linhas.
4. **Consistência cruzada:** TASKS.md ↔ sprints/README.md ↔ SPRINT-CARD.md devem bater.

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

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md ← Resultado consolidado (Fase 1100)
    ├── {CICLO_DIR}/SPRINT-CARD.md + SPRINT-TEST-SUITE.md + SPRINT-REVIEW.md ← Artefatos do ciclo
    └── SPECS_DIR/TASKS.md, SPECS.md, TEST_PLAN.md, ARCHITECTURE.md, PRD.md ← Docs-mestre
    └── {CICLO_DIR}/../README.md (sprints/README.md) ← Índice de sprints
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1100-EXECUTION-REPORT`.

---

## Missão

Atualizar os artefatos do ciclo e os documentos-mestre do projeto com o resultado real do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`, registrando as alterações no `PACKAGE-DEVELOPMENT-ARTIFACTS-UPDATE.md`.

---

## Fluxo de Execução

```
APÓS o relatório de execução (1100), atualizar:

1. ARTEFATOS DO CICLO ({CICLO_DIR}/):
   ├── SPRINT-CARD.md
   │   - Marcar tasks ✅/❌ no backlog; atualizar DoD (checkboxes)
   │   - Atualizar Métricas (tasks, endpoints, RNs, cenários, cobertura)
   │   - Tasks adiadas: documentar
   ├── SPRINT-TEST-SUITE.md
   │   - Marcar cenários ✅/❌; adicionar cenários descobertos na execução
   │   - Atualizar o resumo (total executado)
   └── SPRINT-REVIEW.md
       - Marcar itens demonstrados ✅; atualizar Métricas da Review
       - Documentar bloqueios; preencher Pontos de Verificação (PO)

2. DOCUMENTOS-MESTRE ({SPECS_DIR}/):
   ├── TASKS.md       - Status por task ✅/❌; Progresso no header (ex: 35/99);
   │                    tasks novas criadas na execução → numeração sequencial
   ├── SPECS.md       - Versão/data; Status; RNs formalizadas; endpoints alterados (§4.1/§4.2)
   ├── TEST_PLAN.md   - Versão/data; cenários executados marcados; Status;
   │                    ferramentas de teste novas documentadas
   ├── ARCHITECTURE.md - Versão/data; Status; ADRs novas; pacotes/diretórios novos (§2)
   └── PRD.md         - Versão/data; Status; escopo alterado documentado

3. ÍNDICE DE SPRINTS (sprints/README.md):
   └── Matriz de Rastreabilidade (status/progresso/data por fase)
   └── Tabela de Progresso (datas reais de início/fim)
   └── Versões dos docs-mestre referenciados
   └── Footer com resumo atualizado; novas fases (ex: frentes de débito técnico) na matriz
```

> ⚠️ **Ciclo não totalmente concluído:** registrar tasks pendentes como observações e propor encaminhamento (mover para o próximo ciclo, criar nova task, etc.).

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-ARTIFACTS-UPDATE.md`:

```markdown
# PACKAGE-DEVELOPMENT-ARTIFACTS-UPDATE.md — Atualização de Artefatos: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Artefatos do Ciclo Atualizados
| Artefato | Alterações | Versão/Data |
|:---|:---|:---|
| SPRINT-CARD.md | tasks ✅/❌, DoD, métricas | v1.1 — data |
## 2. Documentos-Mestre Atualizados
| Documento | Alterações | Versão/Data |
|:---|:---|:---|
| TASKS.md | status T-XXX..T-YYY, progresso 35/99 | v1.4 — data |
## 3. Índice de Sprints
[Alterações na matriz, progresso e footer]
## 4. Tasks Pendentes (se houver)
[Tabela: task | motivo | encaminhamento proposto]
## 5. Consistência Cruzada
[TASKS.md ↔ sprints/README.md ↔ SPRINT-CARD.md — status do conferimento]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Consistência cruzada (TASKS ↔ README ↔ CARD) conferida antes de declarar concluída |
| `caveman` | full | Comunicação interativa (nunca nos artefatos) |

---

## Regras de Ouro

1. Docs-mestre prevalecem sobre artefatos do ciclo — em conflito, os mestres vencem.
2. Versão e data atualizadas em TODO documento modificado.
3. NUNCA apagar histórico — marcar ✅/❌, nunca remover linhas.
4. Consistência cruzada conferida ao final (TASKS ↔ README ↔ CARD).
5. Tasks pendentes sempre com encaminhamento proposto — nunca silenciadas.
