# PROMPT-GATE-TASKS-TECHNICAL

## Contexto

Este prompt implementa o **Gate de Viabilidade e Completude Técnica** para o artefato `TASKS.md`, conforme definido no fluxo Spec-Driven Development (etapa 3.1 — GT2).

O agente validador verifica se o plano de tarefas é **completo, sequenciado corretamente, estimado de forma realista e alinhado com as especificações**. O foco aqui é garantir que o time de desenvolvimento receba tarefas acionáveis, sem ambiguidades e com dependências claras.

**Princípio fundamental:** Tarefas mal definidas ou com dependências incorretas geram bloqueios, retrabalho e atrasos. O TASKS.md é o plano de execução — se ele estiver errado, a execução estará errada.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{BRANCH_NAME}` | Nome da branch onde deve ser realizado o desenvolvimento. Negar realizar desenvolvimento direto na branch `main` ou `master` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 5 parâmetros foram informados.

### Passo 1 — Carregar Documentos Base

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TASKS.md (artefato a ser validado)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md (especificações — APROVADO)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md (escopo — APROVADO)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md (arquitetura — APROVADO)
    └── Documentos de referência:
          ├── {PROJECT_PATH}/01-PROJECT-CHARTER-*.md (marcos M1-M7 e datas)
          └── {PROJECT_PATH}/04-FEATURES.md (features e user stories)

Se TASKS.md não existir → ERRO: "TASKS.md não encontrado."
Se SPECS.md não existir ou não estiver APROVADO → ERRO: "SPECS.md aprovado é pré-requisito."
```

### Passo 2 — Executar Dimensões de Validação Técnica

O gate avalia o TASKS.md em **5 dimensões**. Veredito por dimensão: `APROVADO`, `RESSALVA` ou `REPROVADO`.

#### Dimensão 1: Cobertura e Rastreabilidade

| # | Verificação | Critério |
|---|---|---|
| 1.1 | Cobertura de features | Cada feature do SPECS.md tem tarefas correspondentes no TASKS.md? |
| 1.2 | Cobertura de user stories | Cada user story está coberta por pelo menos uma tarefa? |
| 1.3 | Tarefas órfãs | Existem tarefas sem feature/US correspondente no SPECS.md? |
| 1.4 | Rastreabilidade | Cada tarefa referencia a feature (FXX-XX) e a user story (US-XXX) de origem? |

#### Dimensão 2: Granularidade e Acionabilidade

| # | Verificação | Critério |
|---|---|---|
| 2.1 | Tamanho das tarefas | Nenhuma tarefa deve exceder 3 dias de estimativa. Tarefas > 3 dias precisam ser decompostas. |
| 2.2 | Clareza do enunciado | Cada tarefa descreve claramente O QUE fazer (não COMO)? O enunciado é acionável sem consulta adicional? |
| 2.3 | Tarefas atômicas | Cada tarefa tem uma única responsabilidade? Tarefas "E" (fazer X E Y) são suspeitas. |
| 2.4 | Critérios de conclusão | Cada tarefa tem critérios objetivos de DONE? (ex: "Testes passando", "PR aprovado") |

#### Dimensão 3: Sequenciamento e Dependências

| # | Verificação | Critério |
|---|---|---|
| 3.1 | Dependências documentadas | O §3 do TASKS.md documenta dependências entre tarefas? |
| 3.2 | Dependências cíclicas | Existem ciclos de dependência? (A → B → A) |
| 3.3 | Ordem lógica | O sequenciamento respeita a ordem natural? (ex: modelo → repositório → serviço → controller → testes) |
| 3.4 | Marcos do projeto | As tarefas estão alocadas nos marcos M1-M7 conforme Project Charter? |

#### Dimensão 4: Estimativas e Alocação

| # | Verificação | Critério |
|---|---|---|
| 4.1 | Estimativas preenchidas | Todas as tarefas têm estimativa (horas/dias/pontos)? |
| 4.2 | Estimativas realistas | As estimativas são proporcionais à complexidade? (ex: "CRUD simples" ≠ 5 dias) |
| 4.3 | Responsável definido | Cada tarefa tem um responsável ou está marcada como "A definir"? |
| 4.4 | Priorização MoSCoW | A prioridade (Must/Should/Could/Won't) está alinhada com o PRD.md e SPECS.md? |

#### Dimensão 5: Alinhamento com Arquitetura

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Tarefas de infraestrutura | As tarefas refletem a estrutura de pacotes do ARCHITECTURE.md? (ex: tarefa para criar pacote `repository`) |
| 5.2 | Tarefas de segurança | Há tarefas para implementar o pipeline de segurança (JWT, RBAC, Tenant)? |
| 5.3 | Tarefas de testes | Há tarefas explícitas para testes unitários, integração e segurança? |
| 5.4 | Tarefas de cross-cutting | Aspectos como auditoria, logging e tratamento de erros têm tarefas correspondentes? |

### Passo 3 — Calcular Veredito Final

```
Para cada dimensão:
    - 100% verificações OK → APROVADO
    - >= 75% verificações OK → RESSALVA
    - < 75% verificações OK → REPROVADO

Veredito final:
    ├── APROVADO: Todas as 5 dimensões APROVADAS
    ├── RESSALVA: Pelo menos 1 dimensão com RESSALVA, nenhuma REPROVADA
    └── REPROVADO: Pelo menos 1 dimensão REPROVADA
```

### Passo 4 — Gerar Relatório de Falha (se REPROVADO ou RESSALVA)

```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL_TASKS_FAIL_REPORT.md
```

#### Estrutura do Relatório de Falha

```markdown
# TECHNICAL_TASKS_FAIL_REPORT.md — Relatório de Validação Técnica: TASKS.md

[Header: solução, projeto, data, versão do TASKS.md validado]

## 1. Sumário Executivo
- Veredito: REPROVADO | RESSALVA
- Dimensões aprovadas: X/5
- Dimensões com ressalva: Y/5
- Dimensões reprovadas: Z/5
- Total de não-conformidades: N
- Bloqueia TEST_PLAN.md? SIM (se REPROVADO) | NÃO (se RESSALVA)

## 2. Veredito por Dimensão
| Dimensão | Veredito | Verificações OK | Total | % |
|---|---|---|---|---|
| 1. Cobertura e Rastreabilidade | ... | ... | ... | ... |
| 2. Granularidade e Acionabilidade | ... | ... | ... | ... |
| 3. Sequenciamento e Dependências | ... | ... | ... | ... |
| 4. Estimativas e Alocação | ... | ... | ... | ... |
| 5. Alinhamento com Arquitetura | ... | ... | ... | ... |

## 3. Não-Conformidades Detalhadas

### 3.1 [Dimensão 1] Cobertura e Rastreabilidade
| ID | Verificação | Status | Evidência | Ação Corretiva |
|---|---|---|---|---|
| NC-001 | 1.1 Cobertura de features | REPROVADO | Feature F06-03 sem tarefas | Criar tarefas para F06-03 |

[... repetir para cada dimensão]

## 4. Tarefas Problemáticas
- Tarefas > 3 dias que precisam ser decompostas
- Tarefas com dependências cíclicas ou mal definidas
- Tarefas órfãs (sem feature/US correspondente)
- Tarefas com estimativas inconsistentes (muito baixas ou muito altas)

## 5. Análise de Executabilidade
- O plano de tarefas é executável dentro dos marcos do projeto?
- Há gargalos de dependência que podem atrasar a entrega?
- A carga de trabalho por marco está balanceada?

## 6. Recomendações para Correção
- Lista priorizada de ações corretivas
- Tarefas que precisam ser decompostas (com sugestão de granularidade)
- Dependências que precisam ser corrigidas ou adicionadas
- Reestimativas sugeridas para tarefas com estimativas inconsistentes

## 7. Próximos Passos
1. Encaminhar ao Agente de Tarefas para correção
2. Após correção, reexecutar este gate
3. Não prosseguir para TEST_PLAN.md até TASKS.md ser APROVADO

## Rodapé
- Indicação de geração por IA, skills utilizados
```

### Passo 5 — Validar o Relatório Gerado

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo no path correto | `TECHNICAL_TASKS_FAIL_REPORT.md` existe |
| 2 | Sumário executivo | §1 contém veredito e indica se bloqueia TEST_PLAN.md |
| 3 | Veredito por dimensão | §2 lista as 5 dimensões |
| 4 | Não-conformidades detalhadas | §3 lista cada NC com ação corretiva |
| 5 | Tarefas problemáticas | §4 identifica tarefas específicas com problemas |
| 6 | Análise de executabilidade | §5 avalia viabilidade do plano dentro dos marcos |
| 7 | Recomendações acionáveis | §6 contém ações corretivas e sugestões de decomposição |
| 8 | Rodapé de IA | Indicação de geração automatizada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `breakdown-epic-pm` | Avaliar granularidade e decomposição das tarefas |
| 2ª | `writing-plans` | Verificar estrutura do plano, dependências e sequenciamento |
| 3ª | `gap-analysis` | Identificar lacunas de cobertura vs SPECS.md |
| 4ª | `documentation-writer` | Qualidade do relatório de falha |

---

## Observações

1. **Granularidade é o erro mais comum.** Tarefas muito grandes (> 3 dias) são a principal causa de atrasos. O gate é implacável nesse critério: toda tarefa > 3 dias deve ser decomposta.

2. **Dependências cíclicas são bloqueantes.** Um ciclo de dependência (A → B → A) torna o plano inexequível. O gate REPROVA automaticamente se encontrar ciclos.

3. **Estimativas irreais são tão perigosas quanto falta de estimativas.** Uma tarefa "CRUD de entidade" estimada em 5 dias ou "Implementar pipeline de segurança completo" estimada em 4 horas são sinais de que as estimativas não foram pensadas.

4. **Rastreabilidade bidirecional.** Toda tarefa deve apontar para sua feature/US de origem, e toda feature/US deve ter tarefas que a implementam. A ausência em qualquer direção é não-conformidade.

5. **O TASKS.md evolui com a execução.** O modo `delta` do agente gerador de TASKS deve ser usado para manter o plano atualizado. Este gate deve ser reexecutado sempre que houver alteração significativa no plano.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: gate técnico para TASKS.md, 5 dimensões, 20 verificações, foco em granularidade e dependências | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, breakdown-epic-pm, gap-analysis.*
