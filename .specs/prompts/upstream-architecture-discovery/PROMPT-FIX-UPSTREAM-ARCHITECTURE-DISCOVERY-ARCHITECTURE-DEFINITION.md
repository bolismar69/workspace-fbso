# PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F7) — artefato auditado (a ser corrigido)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — Matriz Solução×Disciplina (F9)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` — Consolidação Técnica (F10)
6. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints, data standards globais

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho da pasta de arquitetura global |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP` — skills do time (obter e validar com usuário)
- `PROJECT-TEAM-CAPACITY` — capacidade esperada do time (obter e validar com usuário)
- `PROJECT-STACK` — stack tecnológica (obter e validar com usuário; baseline `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`)
Validar que o artefato auditado existe: `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md`.

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano) e confirmar a leitura dos artefatos:
1. `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` — artefato auditado (a ser corrigido)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — Matriz Solução×Disciplina (F9)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` — Consolidação Técnica (F10)
6. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints, data standards globais
7. `{PROJECT-TEAM-SKILLS-MAP}` — skills do time (obter e validar com usuário)
8. `{PROJECT-TEAM-CAPACITY}` — capacidade esperada do time (obter e validar com usuário)
9. `{PROJECT-STACK}` — stack tecnológica (baseline `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`)

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Solução ausente dos diagramas C4 | Adicionar solução ao L1 e L2 |
| P0 | Integração crítica não documentada | Documentar protocolo, autenticação, fluxo |
| P1 | Diagrama C4 incompleto | Completar elementos faltantes |
| P2 | ADR de integração insuficiente | Adicionar ADR com ID, decisão, alternativas |
| P1 | Stack fora do padrão corporativo sem justificativa | Documentar justificativa técnica ou substituir por tecnologia do padrão FBSO |
| P2 | Skills não mapeados para esta disciplina | Adicionar skills necessários ao PROJECT-TEAM-SKILLS-MAP |
| P2 | Capacidade do time não dimensionada | Estimar capacidade (seniores/plenos/juniores) para esta disciplina |
| P3 | Inconsistência com blueprints | Alinhar com padrão do architecture/ |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `c4-architecture-c4-architecture` | Corrigir diagramas C4 | C4 |
| 3 | `architecture-patterns` | Corrigir padrões arquiteturais | Arquitetura |
| 4 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da definição de arquitetura | Time de Arquitetura |
| 2.0 | 30/07/2026 | Atualização F3→F7: alinhamento com nova numeração do Bloco B | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` | Artefato auditado (F7) — a ser corrigido |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` | Catálogo de Soluções (F8) |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` | Matriz Solução×Disciplina (F9) |
| 5 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` | Consolidação Técnica (F10) |
| 6 | `{ARCHITECTURE_GLOBAL}/` | ADRs, blueprints, data standards globais |
| 7 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 8 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 9 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| 10 | `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| 11 | `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

---

🤖 *Upstream Architecture Discovery — Fase 7 FIX*
