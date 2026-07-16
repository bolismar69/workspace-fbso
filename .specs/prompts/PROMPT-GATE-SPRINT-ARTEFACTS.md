# PROMPT-GATE-SPRINT-ARTEFACTS

## Contexto

Este prompt implementa o **Gate de Qualidade de Artefatos de Sprint** — a validação formal que verifica se os 3 artefatos (`SPRINT-CARD.md`, `SPRINT-TEST-SUITE.md`, `SPRINT-REVIEW.md`) foram gerados corretamente e estão **consistentes entre si e com os documentos-mestre**.

O gate é executado **após a geração** dos artefatos (via `PROMPT-GENERATE-SPRINT-ARTEFACTS`) e **antes do início da execução da sprint**. Se encontrar falhas, gera o relatório `SPRINT_ARTEFACTS_FAIL_REPORT.md` na pasta da sprint para correção.

**Princípio fundamental:** Artefatos de sprint inconsistentes geram retrabalho na execução. Este gate previne que o time comece uma sprint com tarefas mal definidas, testes faltantes ou critérios de aceitação ambíguos.

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
| `{BRANCH_NAME}` | Nome da branch onde deve ser realizado o desenvolvimento. Negar realizar desenvolvimento direto na branch `main` ou `master` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 7 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

### Passo 1 — Carregar Artefatos e Documentos-Mestre

```
SPRINT_DIR   = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/sprints/{SPRINT_NAME}/
SPECS_DIR    = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/

Ler obrigatoriamente:
    ├── {SPRINT_DIR}/SPRINT-CARD.md       ← Artefato a validar
    ├── {SPRINT_DIR}/SPRINT-TEST-SUITE.md ← Artefato a validar
    ├── {SPRINT_DIR}/SPRINT-REVIEW.md     ← Artefato a validar
    ├── {SPECS_DIR}/TASKS.md              ← Baseline para verificação de tarefas
    ├── {SPECS_DIR}/TEST_PLAN.md          ← Baseline para verificação de testes
    ├── {SPECS_DIR}/SPECS.md              ← Baseline para verificação de RNs e endpoints
    └── {SPECS_DIR}/ARCHITECTURE.md       ← Baseline para estrutura e padrões

Se SPRINT-CARD.md não existir → ERRO: "SPRINT-CARD.md não encontrado em {SPRINT_DIR}. Execute o gerador de artefatos de sprint."
Se SPRINT-TEST-SUITE.md não existir → ERRO: "SPRINT-TEST-SUITE.md não encontrado. Execute o gerador de artefatos de sprint."
Se SPRINT-REVIEW.md não existir → ERRO: "SPRINT-REVIEW.md não encontrado. Execute o gerador de artefatos de sprint."
```

### Passo 2 — Executar Dimensões de Validação

O gate avalia os 3 artefatos em **6 dimensões**. Para cada dimensão, atribuir um veredito: `APROVADO`, `RESSALVA` ou `REPROVADO`.

#### Dimensão 1: Completude dos Artefatos

| # | Verificação | Critério |
|:---|:---|:---|
| 1.1 | SPRINT-CARD.md tem todas as seções | Goal, Backlog, Features, DoD, Riscos, Dependências, Métricas presentes |
| 1.2 | SPRINT-TEST-SUITE.md tem todas as seções | Header, cenários por feature, resumo, RNs cobertas presentes |
| 1.3 | SPRINT-REVIEW.md tem todas as seções | O que demonstrar, pontos de verificação, bloqueios, próximo passo presentes |
| 1.4 | Sprint backlog cobre 100% das tarefas | Número de tarefas no SPRINT-CARD.md = número de tarefas no TASKS.md para esta sprint |
| 1.5 | Todas as features da sprint têm cenários | Cada feature listada no SPRINT-CARD.md tem cenários correspondentes no SPRINT-TEST-SUITE.md |

#### Dimensão 2: Consistência com TASKS.md

| # | Verificação | Critério |
|:---|:---|:---|
| 2.1 | IDs de tarefas preservados | T-XXX no SPRINT-CARD.md batem exatamente com TASKS.md |
| 2.2 | Estimativas preservadas | Cada tarefa tem a mesma estimativa (em dias) do TASKS.md |
| 2.3 | Critérios DONE preservados | O critério DONE de cada tarefa no SPRINT-CARD.md é idêntico ou equivalente ao do TASKS.md |
| 2.4 | Features referenciadas corretamente | Cada tarefa referencia o Feature ID (FXX-XX) e User Stories (US-XXX) corretos do TASKS.md |
| 2.5 | Prioridades MoSCoW preservadas | Must/Should/Could no SPRINT-CARD.md batem com TASKS.md |
| 2.6 | Sem tarefas órfãs ou extras | Nenhuma tarefa no SPRINT-CARD.md que não existe no TASKS.md. Nenhuma tarefa do TASKS.md ausente no SPRINT-CARD.md |

#### Dimensão 3: Consistência com TEST_PLAN.md

| # | Verificação | Critério |
|:---|:---|:---|
| 3.1 | IDs de cenários preservados | TC-XXX-XXX no SPRINT-TEST-SUITE.md batem exatamente com TEST_PLAN.md |
| 3.2 | Nível de teste preservado | Unit/Integração/E2E/Segurança de cada cenário bate com TEST_PLAN.md |
| 3.3 | Features cobertas corretamente | Cenários extraídos pertencem às features desta sprint (não de sprints futuras ou passadas) |
| 3.4 | Sem cenários órfãos | Nenhum cenário no SPRINT-TEST-SUITE.md sem correspondência no TEST_PLAN.md |
| 3.5 | Cenários transversais incluídos | Testes de segurança e isolamento multi-tenant aplicáveis estão incluídos |

#### Dimensão 4: Consistência com SPECS.md

| # | Verificação | Critério |
|:---|:---|:---|
| 4.1 | RNs referenciadas corretamente | RNs listadas no SPRINT-CARD.md e SPRINT-TEST-SUITE.md batem com SPECS.md §3 |
| 4.2 | Endpoints REST corretos | Endpoints mencionados batem com SPECS.md §4.1 para as features da sprint |
| 4.3 | Critérios de aceitação alinhados | SPRINT-REVIEW.md reflete os critérios DONE do SPECS.md §7 |
| 4.4 | Matriz de cobertura preservada | Feature → RN → US → Task consistente com SPECS.md §3.3 |

#### Dimensão 5: Consistência Cruzada entre os 3 Artefatos

| # | Verificação | Critério |
|:---|:---|:---|
| 5.1 | Features no CARD batem com TEST-SUITE | As mesmas features aparecem nos dois artefatos |
| 5.2 | RNs no CARD batem com TEST-SUITE | As mesmas RNs são referenciadas nos dois artefatos |
| 5.3 | Sprint Goal consistente | O goal no SPRINT-CARD.md é refletido nos cenários do SPRINT-TEST-SUITE.md e nos itens de demonstração do SPRINT-REVIEW.md |
| 5.4 | DONE criteria → REVIEW | Os critérios DONE do SPRINT-CARD.md são verificáveis nos checkboxes do SPRINT-REVIEW.md |
| 5.5 | Dependências → REVIEW | As dependências listadas no CARD são coerentes com o "Próximo Passo" da sprint anterior no REVIEW |
| 5.6 | Sem conteúdo duplicado ou contraditório | Nenhuma informação nos 3 artefatos que se contradiz |

#### Dimensão 6: Qualidade dos Scripts de Demo (SPRINT-REVIEW.md)

| # | Verificação | Critério |
|:---|:---|:---|
| 6.1 | Scripts presentes para features visíveis | A partir da Sprint 3, cada feature com interface tem 🎬 Script: |
| 6.2 | Scripts são acionáveis | Cada script descreve ações concretas (o que clicar, o que digitar, o que esperar) |
| 6.3 | Checkboxes de verificação completos | Todos os itens de demonstração têm checkbox |
| 6.4 | Sprint 1 e 2 tratadas corretamente | Sprints de fundação não exigem scripts de demo para PO, mas têm checklist de validação técnica |

### Passo 3 — Calcular Veredito Final

```
Para cada dimensão:
    - 100% verificações OK → APROVADO
    - >= 75% verificações OK → RESSALVA
    - < 75% verificações OK → REPROVADO

Veredito final do gate:
    ├── APROVADO: Todas as 6 dimensões APROVADAS
    ├── RESSALVA: Pelo menos 1 dimensão com RESSALVA, nenhuma REPROVADA
    └── REPROVADO: Pelo menos 1 dimensão REPROVADA
```

### Passo 4 — Gerar Relatório de Falha (se REPROVADO ou RESSALVA)

Se o veredito for **REPROVADO** ou **RESSALVA**, gerar o relatório:

```
{SPRINT_DIR}/SPRINT_ARTEFACTS_FAIL_REPORT.md
```

#### Estrutura do Relatório de Falha

```markdown
# SPRINT_ARTEFACTS_FAIL_REPORT.md — Relatório de Validação: Sprint {N}

[Header: solução, projeto, sprint, stack, data da validação]

## 1. Sumário Executivo
- Veredito: REPROVADO | RESSALVA
- Dimensões aprovadas: X/6
- Dimensões com ressalva: Y/6
- Dimensões reprovadas: Z/6
- Total de não-conformidades: N
- Artefatos com falhas: [lista]

## 2. Veredito por Dimensão
| Dimensão | Veredito | Verificações OK | Total Verificações | % |
|:---|:---|:---:|:---:|:---:|
| 1. Completude | ... | ... | 5 | ... |
| 2. Consistência TASKS.md | ... | ... | 6 | ... |
| 3. Consistência TEST_PLAN.md | ... | ... | 5 | ... |
| 4. Consistência SPECS.md | ... | ... | 4 | ... |
| 5. Consistência Cruzada | ... | ... | 6 | ... |
| 6. Qualidade Scripts Demo | ... | ... | 4 | ... |

## 3. Não-Conformidades Detalhadas

### 3.1 [Dimensão 1] Completude
| ID | Verificação | Artefato | Status | Evidência | Ação Corretiva |
|:---|:---|:---|:---|:---|:---|
| NC-001 | 1.4 Cobertura tarefas | SPRINT-CARD.md | REPROVADO | Sprint backlog lista 12 tarefas, mas TASKS.md tem 15 para este marco | Adicionar tarefas T-XXX, T-YYY, T-ZZZ ao sprint backlog |

### 3.2 [Dimensão 2] Consistência com TASKS.md
...

### 3.3 [Dimensão 3] Consistência com TEST_PLAN.md
...

### 3.4 [Dimensão 4] Consistência com SPECS.md
...

### 3.5 [Dimensão 5] Consistência Cruzada
...

### 3.6 [Dimensão 6] Qualidade Scripts Demo
...

## 4. Artefatos Afetados
| Artefato | NCs | Severidade |
|:---|:---:|:---|
| SPRINT-CARD.md | N | Alta/Média/Baixa |
| SPRINT-TEST-SUITE.md | N | ... |
| SPRINT-REVIEW.md | N | ... |

## 5. Recomendações para Correção
[Lista priorizada de ações corretivas, agrupada por artefato]

## 6. Próximos Passos
1. Encaminhar este relatório ao Agente Corretor de Artefatos de Sprint
2. Após correção, reexecutar este gate
3. Somente iniciar a execução da sprint após APROVADO

## Rodapé
[Indicação de geração por IA, skills utilizados]
```

### Passo 5 — Validar o Relatório Gerado

| # | Verificação | Critério |
|:---|:---|:---|
| 1 | Arquivo no path correto | `SPRINT_ARTEFACTS_FAIL_REPORT.md` existe na pasta da sprint |
| 2 | Sumário executivo presente | §1 contém veredito e contagem |
| 3 | Tabela de veredito por dimensão | §2 lista as 6 dimensões com métricas |
| 4 | Não-conformidades detalhadas | §3 lista cada NC com ID, artefato, evidência e ação corretiva |
| 5 | Artefatos afetados identificados | §4 agrupa NCs por artefato com severidade |
| 6 | Recomendações acionáveis | §5 contém ações priorizadas |
| 7 | Próximos passos claros | §6 indica fluxo de correção e revalidação |
| 8 | Rodapé de IA | Indicação de geração automatizada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|:---|:---|:---|
| 1ª | `gap-analysis` | Identificar lacunas entre artefatos de sprint e documentos-mestre |
| 2ª | `test-strategy-design` | Validar cobertura e granularidade dos cenários de teste |
| 3ª | `acceptance-criteria` | Validar critérios de aceitação e scripts de demo |
| 4ª | `ponytail` | Revisor de simplicidade — artefatos estão enxutos ou há ruído? |
| 5ª | `documentation-writer` | Qualidade do relatório de falha |

---

## Observações

1. **Gate é mandatório antes de iniciar a sprint.** Nenhuma tarefa deve ser executada antes que os artefatos passem neste gate. Artefatos inconsistentes são a principal causa de retrabalho durante a sprint.

2. **Rastreabilidade é a chave.** A dimensão 2, 3 e 4 verificam consistência com os 3 documentos-mestre. IDs de tarefas, cenários e RNs devem ser idênticos — não "parecidos", IDÊNTICOS.

3. **Consistência cruzada (Dimensão 5) é o diferencial.** Não basta cada artefato estar correto isoladamente — eles precisam contar a mesma história. Se o SPRINT-CARD.md lista uma feature mas o SPRINT-TEST-SUITE.md não tem cenários para ela, há uma falha de consistência.

4. **SPRINT_ARTEFACTS_FAIL_REPORT.md é gerado na pasta da sprint**, junto com os artefatos que precisam ser corrigidos. Isso mantém o contexto localizado.

5. **Sprints de fundação (1 e 2) têm regras diferentes.** A Dimensão 6 reconhece que Sprints 1 e 2 não têm features visíveis para o PO. Seus SPRINT-REVIEW.md são focados em validação técnica.

6. **Revalidação obrigatória.** Após a correção via `PROMPT-FIX-SPRINT-ARTEFACTS`, este gate deve ser reexecutado. O ciclo é: GENERATE → GATE → (se falhar) FIX → GATE → (APROVADO) EXECUTAR SPRINT.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 14/07/2026 | Criação inicial: gate de qualidade para artefatos de sprint, 6 dimensões, 30 verificações, foco em consistência cruzada e rastreabilidade | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, gap-analysis, test-strategy-design, acceptance-criteria, ponytail, documentation-writer.*
