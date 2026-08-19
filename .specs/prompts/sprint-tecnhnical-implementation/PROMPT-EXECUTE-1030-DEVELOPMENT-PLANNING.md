# PROMPT-EXECUTE-1030-DEVELOPMENT-PLANNING

## Contexto

Este prompt executa a **Fase de Planejamento do Desenvolvimento** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 1 — passo 8). Gera o plano de desenvolvimento do ciclo ANTES de iniciar a implementação: o que será implementado, em qual ordem, com quais dependências e qual a abordagem técnica de cada task.

**Princípios fundamentais:**

1. **Plano antes do código:** nenhuma task é implementada sem este artefato.
2. **Fonte da verdade:** o plano deriva do SPRINT-CARD.md (goal, backlog, DONE) e do TASKS.md (§2 tarefas, §3 dependências) — nunca cria conteúdo novo.
3. **Artefato único por ciclo:** gerado UMA VEZ; se o plano mudar durante a execução, atualizar o próprio arquivo.

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
| `{TASK_IDS}` | IDs das tarefas do plano (opcional — vazio = todas pendentes) | `T-001,T-002` |

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/SPRINT-CARD.md                    ← Goal, backlog, DONE criteria, riscos
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md ← tasks selecionadas, ordem validada, stack, skills (Fase 1010)
    └── SPECS_DIR/TASKS.md                            ← Fonte da verdade das tarefas (§2) e dependências (§3)
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1020-PRE-IMPLEMENTATION`.

---

## Missão

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-PLANNING.md` — o plano de desenvolvimento do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`, cobrindo visão geral, dependências, plano por task, ordem de execução e estratégia de build/verificação.

---

## Fluxo de Execução

1. **Carregar** SPRINT-CARD.md (goal, backlog, DONE, riscos), PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md (tasks selecionadas, ordem validada, stack, skills) e TASKS.md §2/§3 (dependências).
2. **Gerar o artefato** com a estrutura obrigatória:

```markdown
# PACKAGE-DEVELOPMENT-PLANNING.md — Plano de Desenvolvimento: Ciclo {N}
[Header: solução, projeto, ciclo, stack, data]

## 1. Visão Geral
- Goal do Ciclo: [do SPRINT-CARD.md]
- Tasks a implementar: X
- Ordem de execução: [sequencial | paralela | mista]
- Stack: [stack detectada + versões]

## 2. Dependências entre Tasks
[Diagrama ou lista mostrando ordem e dependências entre T-XXX]
[Formato: T-001 → (T-002, T-003) → T-005 → T-007]

## 3. Plano por Task
### T-XXX — [Nome da Tarefa]
- **Critério DONE:** [do SPRINT-CARD.md]
- **Estimativa:** Xd
- **Abordagem:** [estratégia de implementação]
- **Arquivos a criar:** [tabela: Arquivo | Tipo | Descrição]
- **Arquivos a modificar:** [tabela: Arquivo | Tipo | Descrição]
- **Dependências:** [tasks que devem estar concluídas antes]
- **Riscos:** [riscos específicos desta task]
- **Skills aplicáveis:** [skills da stack que guiam a implementação]
[Repetir para cada task]

## 4. Ordem de Execução
[Lista numerada com a ordem exata, justificando escolhas]

## 5. Estratégia de Build e Verificação
- Comando de build: `[comando]`
- Comando de teste rápido: `[comando]`
- Checkpoints: [momentos em que o build deve ser verificado]

## Rodapé
[Indicação de geração por IA, data/hora]
```

3. **Validar** contra os docs-base: toda task do plano existe no SPRINT-CARD; toda dependência está no TASKS.md §3; skills batem com as registradas na pré-implementação.
4. **Apresentar ao humano** para validação (`[STATUS: COMPLIANCE]` só com aprovação).

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| Skills da stack | herdadas | Registradas na 1010 (`PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md` §6) — preenchem "Skills aplicáveis" de cada task do plano |
| `ponytail` | full | Rung 1 (YAGNI) na análise de necessidade de cada task |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. NUNCA implementar sem este artefato aprovado.
2. Plano deriva de SPRINT-CARD + TASKS — nunca inventar task, dependência ou estimativa.
3. Artefato gerado UMA VEZ antes da primeira task; mudanças no plano = atualizar o arquivo (não recriar).
4. Artefato em `{CICLO_DIR}/` com `[STATUS: Em análise]` inicial; COMPLIANCE apenas com validação humana.
