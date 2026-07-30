# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-DEVOPS-SRE-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{ARCHITECTURE_GLOBAL}` | Caminho da pasta de arquitetura global |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual, Architecture Definition (F7), Security Definition (F8), ADRs globais.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Pipeline CI/CD sem rollback | Adicionar estratégia de rollback automatizado |
| P0 | Solução sem observabilidade | Adicionar logging, métricas e tracing para a solução |
| P0 | SLOs não definidos | Definir error budgets, latency e availability targets |
| P1 | IaC incompleto | Documentar estrutura de módulos e state management |
| P2 | Runbook de incidente ausente | Criar procedimentos S1-S4 com escalação |
| P3 | Alerting sem thresholds | Definir canais e thresholds por métrica |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `senior-devops` | Corrigir estratégia DevOps | DevOps |
| 3 | `sre-engineer` | Corrigir SLOs e runbooks | SRE |
| 4 | `observability-engineer` | Corrigir observabilidade | Observabilidade |
| 5 | `kubernetes-specialist` | Corrigir topologia K8s | K8s |
| 6 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção da definição de DevOps e SRE | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
