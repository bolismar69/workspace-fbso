# PROMPT-FIX-SPRINT-ARTEFACTS

## Contexto

Este prompt é acionado quando o **Gate de Qualidade de Artefatos de Sprint** (`PROMPT-GATE-SPRINT-ARTEFACTS`) reprova os artefatos e gera o relatório `SPRINT_ARTEFACTS_FAIL_REPORT.md`.

O agente atua como **corretor de artefatos de sprint** — lê o relatório de falha e aplica correções cirúrgicas nos 3 arquivos (`SPRINT-CARD.md`, `SPRINT-TEST-SUITE.md`, `SPRINT-REVIEW.md`), realinhando-os com os documentos-mestre (`TASKS.md`, `TEST_PLAN.md`, `SPECS.md`).

**Princípio fundamental:** A correção NUNCA altera os documentos-mestre. Se uma inconsistência revelar um erro nos documentos-mestre, isso deve ser reportado como ação separada — mas a correção dos artefatos de sprint sempre se alinha aos documentos-mestre, nunca o contrário.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço) | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SPRINT_NUMBER}` | Número da sprint (1 a N) | `1` |
| `{SPRINT_NAME}` | Nome curto da sprint (kebab-case) | `sprint-01-setup` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 6 parâmetros foram informados.

### Passo 1 — Carregar Relatório de Falha, Artefatos e Documentos-Mestre

```
SPRINT_DIR   = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/sprints/{SPRINT_NAME}/
SPECS_DIR    = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/

Ler obrigatoriamente:
    ├── {SPRINT_DIR}/SPRINT_ARTEFACTS_FAIL_REPORT.md  ← Relatório de falha (fonte das NCs)
    ├── {SPRINT_DIR}/SPRINT-CARD.md                   ← Artefato a corrigir
    ├── {SPRINT_DIR}/SPRINT-TEST-SUITE.md             ← Artefato a corrigir
    ├── {SPRINT_DIR}/SPRINT-REVIEW.md                 ← Artefato a corrigir
    ├── {SPECS_DIR}/TASKS.md                          ← Baseline de referência
    ├── {SPECS_DIR}/TEST_PLAN.md                      ← Baseline de referência
    └── {SPECS_DIR}/SPECS.md                          ← Baseline de referência

Se SPRINT_ARTEFACTS_FAIL_REPORT.md não existir → ERRO: "Relatório de falha não encontrado. Execute o gate primeiro."
```

### Passo 2 — Processar Não-Conformidades por Prioridade

Extrair do `SPRINT_ARTEFACTS_FAIL_REPORT.md` todas as NCs e classificá-las:

| Prioridade | Tipo de NC | Ação |
|:---|:---|:---|
| **P0 (Crítica)** | Tarefa do TASKS.md ausente no SPRINT-CARD.md | **Adicionar** tarefa ao backlog com ID, estimativa e critério DONE do TASKS.md |
| **P0 (Crítica)** | Feature sem cenários de teste no SPRINT-TEST-SUITE.md | **Extrair** cenários do TEST_PLAN.md e adicionar à suite |
| **P0 (Crítica)** | IDs de tarefa ou cenário divergentes dos docs-mestre | **Corrigir** IDs para igualar exatamente aos docs-mestre |
| **P1 (Alta)** | Critério DONE divergente entre CARD e TASKS.md | **Alinhar** critério DONE com TASKS.md |
| **P1 (Alta)** | RNs referenciadas incorretamente (SPECS.md §3) | **Corrigir** referências de RN conforme SPECS.md |
| **P1 (Alta)** | Inconsistência cruzada: feature no CARD sem correspondência no TEST-SUITE ou REVIEW | **Completar** o artefato faltante |
| **P2 (Média)** | Seção obrigatória ausente ou incompleta em qualquer artefato | **Completar** a seção conforme template |
| **P2 (Média)** | Script de demo faltante para feature visível (Sprint ≥3) | **Adicionar** 🎬 Script: com ações concretas |
| **P2 (Média)** | Estimativa de tarefa divergente do TASKS.md | **Corrigir** estimativa |
| **P3 (Baixa)** | Sprint Goal vago ou não mensurável | **Reescrever** goal com verbo imperativo + métrica |
| **P3 (Baixa)** | Checkboxes faltantes no SPRINT-REVIEW.md | **Adicionar** checkboxes |
| **P3 (Baixa)** | Links quebrados para documentos-mestre | **Corrigir** paths relativos |

### Passo 3 — Aplicar Correções por Artefato

```
Para cada NC no relatório, agrupada por artefato afetado:
    │
    ├── NC no SPRINT-CARD.md?
    │     ├── Tarefa ausente → Adicionar linha na tabela de Sprint Backlog
    │     ├── ID/Tarefa incorreto → Corrigir para bater com TASKS.md
    │     ├── Critério DONE divergente → Alinhar com TASKS.md
    │     ├── Feature/Rn mal referenciada → Corrigir referência
    │     ├── Seção obrigatória faltando → Adicionar seção conforme template
    │     ├── Sprint Goal vago → Reescrever com formato "Verbo. Métrica."
    │     └── Dependências incorretas → Corrigir conforme TASKS.md §3
    │
    ├── NC no SPRINT-TEST-SUITE.md?
    │     ├── Feature sem cenários → Extrair cenários do TEST_PLAN.md
    │     ├── ID de cenário divergente → Corrigir para ID exato do TEST_PLAN.md
    │     ├── Nível de teste incorreto → Corrigir (Unit/Int/E2E/Seg)
    │     ├── Cenário de sprint errada → Remover (se futuro) ou adicionar (se faltando)
    │     ├── RN não coberta → Adicionar cenário que cobre a RN
    │     └── Tabela de resumo incorreta → Recalcular contagens
    │
    └── NC no SPRINT-REVIEW.md?
          ├── Feature sem item de demonstração → Adicionar checklist + 🎬 Script:
          ├── Script de demo vago → Reescrever com ações concretas
          ├── Checkbox faltando → Adicionar checkbox
          ├── Próximo passo incorreto → Corrigir com dados da próxima sprint
          └── Seção obrigatória faltando → Adicionar conforme template
```

### Passo 4 — Verificar Consistência Cruzada Pós-Correção

Após aplicar todas as correções, reexecutar as verificações da **Dimensão 5** (Consistência Cruzada) do gate:

| # | Verificação | Ação se falhar |
|:---|:---|:---|
| 5.1 | Features no CARD batem com TEST-SUITE | Sincronizar features entre os dois artefatos |
| 5.2 | RNs no CARD batem com TEST-SUITE | Sincronizar RNs entre os dois artefatos |
| 5.3 | Sprint Goal → cenários → demo coerentes | Ajustar goal ou adicionar cenários/demo |
| 5.4 | DONE criteria → REVIEW checkboxes | Adicionar checkboxes faltantes no REVIEW |
| 5.5 | Dependências coerentes com sprint anterior | Corrigir referências |

### Passo 5 — Atualizar Registro de Alterações

Em cada artefato modificado, adicionar entrada no registro de alterações (se o artefato tiver essa seção) ou adicionar um comentário de rodapé:

```markdown
> ⚠️ Correção pós-gate: {N} NCs resolvidas do SPRINT_ARTEFACTS_FAIL_REPORT.md v{Y}. NCs: {lista de IDs}. Data: {data}.
```

### Passo 6 — Validar Correções

| # | Verificação | Critério |
|:---|:---|:---|
| 1 | Todas as NCs endereçadas | Cada NC do relatório tem correção aplicada ou justificativa documentada |
| 2 | Nenhuma alteração nos docs-mestre | TASKS.md, TEST_PLAN.md, SPECS.md não foram modificados |
| 3 | IDs preservados | IDs de tarefas (T-XXX), cenários (TC-XXX) e RNs (RNXX-XX) batem com docs-mestre |
| 4 | Consistência cruzada restaurada | Verificações da Dimensão 5 passam |
| 5 | Artefatos prontos para re-gate | Arquivos corrigidos e salvos nos paths corretos |
| 6 | Registro de correção presente | Cada artefato modificado indica que foi corrigido pós-gate |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|:---|:---|:---|
| 1ª | `gap-analysis` | Analisar relatório de falha e planejar correções |
| 2ª | `breakdown-plan` | Corrigir sprint backlog e dependências |
| 3ª | `test-strategy-design` | Corrigir cobertura e organização de cenários de teste |
| 4ª | `acceptance-criteria` | Corrigir critérios de aceitação e scripts de demo |
| 5ª | `ponytail` | Revisor final — as correções mantiveram os artefatos enxutos? |
| 6ª | `documentation-writer` | Qualidade dos artefatos corrigidos |

---

## Observações

1. **Documentos-mestre são IMUTÁVEIS neste fluxo.** Se uma NC revelar um erro no TASKS.md, TEST_PLAN.md ou SPECS.md, a correção NÃO deve ser feita aqui. Documente a descoberta no relatório de correção e siga o fluxo de correção do documento-mestre específico (`PROMPT-FIX-TASKS-TECHNICAL_SOLUTION.md`, etc.).

2. **Corrigir o artefato, não reescrevê-lo.** As correções devem ser cirúrgicas — alterar apenas o que o relatório de falha apontou. Não reescrever seções inteiras que não foram questionadas.

3. **Consistência cruzada é o ponto mais frágil.** A maioria das NCs tende a ser de inconsistência entre os 3 artefatos (Dimensão 5). Após corrigir cada artefato individualmente, SEMPRE verificar se eles continuam consistentes entre si.

4. **Prioridade segue severidade do gate.** P0 (REPROVADO) > P1 (RESSALVA alta) > P2 (RESSALVA média) > P3 (RESSALVA baixa). Corrigir primeiro o que impede a aprovação.

5. **Após correção, reexecutar o gate.** O ciclo é: FIX → GATE. Se APROVADO, a sprint pode começar. Se ainda REPROVADO ou RESSALVA, gerar novo FAIL_REPORT e repetir o FIX.

6. **Sprints de fundação (1 e 2) têm regras diferentes para demo.** Não adicionar scripts de demo para PO em Sprints 1 e 2 — elas são de validação técnica. O SPRINT-REVIEW.md dessas sprints foca em verificações de build, migration e segurança.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 14/07/2026 | Criação inicial: correção de artefatos de sprint baseada em SPRINT_ARTEFACTS_FAIL_REPORT.md, 3 níveis de prioridade (P0-P3), correções agrupadas por artefato | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, gap-analysis, breakdown-plan, test-strategy-design, acceptance-criteria, ponytail, documentation-writer.*
