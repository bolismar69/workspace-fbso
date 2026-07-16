# PROMPT-FIX-TASKS-FROM-GATE

## Contexto

Este prompt é acionado quando o **Gate de Viabilidade e Completude Técnica** (PROMPT-GATE-TASKS-TECHNICAL) reprova o `TASKS.md` e gera o relatório `TECHNICAL_TASKS_FAIL_REPORT.md`.

O agente atua como **corretor de planejamento** — lê o relatório de falha e aplica correções no TASKS.md, com foco em **granularidade** (decompor tarefas grandes), **dependências** (resolver ciclos e lacunas) e **cobertura** (garantir que toda feature/US tem tarefas).

**Princípio fundamental:** Um plano de tarefas incorreto produz atrasos e retrabalho. A correção prioriza decompor tarefas grandes e resolver dependências quebradas.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar TODOS os 4 parâmetros.

### Passo 1 — Carregar Artefatos e Relatório de Falha

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL_TASKS_FAIL_REPORT.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TASKS.md (a corrigir)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md
    └── {PROJECT_PATH}/01-PROJECT-CHARTER-*.md (marcos e datas)
```

### Passo 2 — Processar Não-Conformidades por Prioridade

Extrair do `TECHNICAL_TASKS_FAIL_REPORT.md`, com atenção ao §4 (Tarefas Problemáticas) e §5 (Análise de Executabilidade):

| Prioridade | Tipo de NC | Ação |
|---|---|---|
| P0 (Bloqueante) | Dependência cíclica (A → B → A) | **Resolver ciclo** reordenando ou quebrando tarefas |
| P0 (Bloqueante) | Feature/US sem tarefas correspondentes | **Criar** tarefas para cobrir a feature/US |
| P1 (Alta) | Tarefa > 3 dias (granularidade grossa) | **Decompor** tarefa em 2+ subtarefas |
| P1 (Alta) | Tarefa sem estimativa ou com estimativa irreal | **Estimar** ou **reestimar** com base na complexidade |
| P2 (Média) | Tarefa órfã (sem feature/US) | **Vincular** à feature/US correta ou **remover** |
| P2 (Média) | Sequenciamento incorreto (ex: teste antes da implementação) | **Reordenar** tarefas |
| P3 (Baixa) | Tarefa sem responsável, prioridade inconsistente | **Atribuir** ou marcar "A definir", **alinhar** prioridade |

### Passo 3 — Aplicar Correções no TASKS.md

```
Para cada NC no relatório:
    │
    ├── NC de CICLO (P0)?
    │     → Identificar tarefas no ciclo (ex: T-012 → T-015 → T-012)
    │     → Quebrar uma das tarefas para quebrar o ciclo
    │     → Ou reordenar para estabelecer sequência linear
    │     → Atualizar §3 (Dependências) do TASKS.md
    │
    ├── NC de COBERTURA (P0)?
    │     → Feature/US sem tarefas → consultar SPECS.md para entender escopo
    │     → Criar tarefas seguindo ordem natural: modelo → repositório → serviço → controller → testes
    │     → Estimar cada nova tarefa
    │     → Adicionar ao marco correto conforme Project Charter
    │
    ├── NC de GRANULARIDADE (P1)?
    │     → Tarefa > 3 dias → decompor em subtarefas de 0.5-2 dias cada
    │     → Exemplo: "Implementar CRUD de Usuário (5 dias)" →
    │         T-xxx.1: Entidade User + repositório (1 dia)
    │         T-xxx.2: UserService com regras de negócio (1.5 dias)
    │         T-xxx.3: UserController + DTOs (1 dia)
    │         T-xxx.4: Testes unitários e integração (1.5 dias)
    │     → Atualizar dependências entre subtarefas
    │
    ├── NC de ESTIMATIVA (P1)?
    │     → Comparar com tarefas similares no plano
    │     → Ajustar estimativa para faixa realista
    │     → Documentar premissa da estimativa
    │
    └── NC de ORGANIZAÇÃO (P2/P3)?
          → Reordenar, vincular feature/US, atribuir responsável, alinhar prioridade
```

### Passo 4 — Recalcular Progresso e Métricas

Após correções, atualizar:

```markdown
## 1. Visão Geral (atualizar)
- Total de tarefas: {novo total} (eram {antigo total})
- Tarefas por marco: M{X}: {contagem}
```

### Passo 5 — Atualizar Registro de Alterações

```markdown
| v{X+1} | {data} | Correção pós-gate: {N} não-conformidades resolvidas do TECHNICAL_TASKS_FAIL_REPORT.md v{Y}. {D} tarefas decompostas, {C} tarefas criadas, {R} tarefas removidas. | Agente Corretor TASKS/IA |
```

### Passo 6 — Validar Correções

| # | Verificação | Critério |
|---|---|---|
| 1 | Ciclos resolvidos | Nenhuma dependência cíclica no plano |
| 2 | Cobertura completa | Toda feature/US do SPECS.md tem tarefas |
| 3 | Granularidade adequada | Nenhuma tarefa > 3 dias |
| 4 | Estimativas preenchidas | Todas as tarefas têm estimativa realista |
| 5 | Dependências documentadas | §3 reflete sequenciamento correto |
| 6 | Tarefas órfãs removidas | Tarefas sem feature/US foram vinculadas ou removidas |
| 7 | Registro de alterações | Nova versão documentada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `breakdown-epic-pm` | Decompor tarefas grandes em subtarefas |
| 2ª | `writing-plans` | Reestruturar plano com dependências corretas |
| 3ª | `gap-analysis` | Analisar relatório e planejar correções |
| 4ª | `documentation-writer` | Qualidade do TASKS.md corrigido |

---

## Observações

1. **Decomposição é a correção mais comum.** Tarefas grandes (> 3 dias) são a principal causa de reprovação no gate. Aprenda a decompor: uma tarefa de 5 dias geralmente vira 3-4 subtarefas de 1-1.5 dias.

2. **Ciclos de dependência são raros mas bloqueantes.** Se o relatório apontar um ciclo, resolva-o primeiro. Todo o resto é secundário.

3. **Criar tarefas seguindo a estrutura do ARCHITECTURE.md.** Para features sem cobertura, criar tarefas que sigam a estrutura de pacotes documentada: modelo → repositório → serviço → controller → testes.

4. **Não remova tarefas sem verificar.** Se o relatório apontar uma "tarefa órfã", verifique se ela não é uma tarefa de infraestrutura legítima antes de remover. Tarefas de configuração (Docker, CI/CD) podem não estar vinculadas a uma feature específica.

5. **Incrementar versão SEMPRE.** Após correção, reexecutar o gate (FIX → GATE → aprovação).

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: correção de TASKS.md baseada em TECHNICAL_TASKS_FAIL_REPORT.md, foco em decomposição e resolução de dependências | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, breakdown-epic-pm, writing-plans.*
