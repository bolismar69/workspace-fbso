# PROMPT-GENERATE-600-EXECUTION-HISTORY

## Contexto

Este prompt gerencia o artefato `600-EXECUTION-HISTORY.md` — o **dashboard de controle** que registra o estado de CADA documento gerado em TODAS as 19 fases + artefatos migrados do roadmap. Diferente dos outros artefatos, este NÃO tem gate/fix próprios — ele é **atualizado incrementalmente** após cada fase ser concluída.

**Posicionamento:** O EXECUTION-HISTORY é um artefato **standalone**, executado após a Barreira D (Bloco D 100% COMPLIANCE). Ele não pertence a nenhum bloco — é o fechamento do pipeline completo.

**Princípio fundamental:** O Execution History é o single source of truth para o status do roadmap. Qualquer pessoa deve conseguir abrir este arquivo e saber exatamente em que estado está cada um dos 20 artefatos.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Criação Inicial (executado uma vez após o Bootstrap)

Gerar `{TECHNICAL_DEFINITIONS_PATH}/600-EXECUTION-HISTORY.md` com:

**1. Tabela-Mestra de Documentos (20 artefatos):**

| Documento | Bloco | Fase | Caminho | Estado | Data Criação | Última Atualização | NCs Pendentes | Veredito Gate |
|---|---|---|---|---|---|---|---|---|
| **Bloco 0 — Product Def & Backlog & PRD** | | | | | | | | |
| INTAKE-LOG.md | 0 | F1 | .../INTAKE-LOG.md | PENDENTE | — | — | — | — |
| DOR-ASSESSMENT.md | 0 | F2 | .../DOR-ASSESSMENT.md | PENDENTE | — | — | — | — |
| PRODUCT-BACKLOG-LIST.md | 0 | F3 | .../PRODUCT-BACKLOG-LIST.md | PENDENTE | — | — | — | — |
| PRD-DEFINITION.md | 0 | F4 | .../PRD-DEFINITION.md | PENDENTE | — | — | — | — |
| **Bloco A — People & Solutions** | | | | | | | | |
| TEAM-SKILLS-MAP.md | A | F5 | .../TEAM-SKILLS-MAP.md | PENDENTE | — | — | — | — |
| TEAM-CAPACITY.md | A | F6 | .../TEAM-CAPACITY.md | MIGRADO | — | — | — | — |
| TEAM-CAPACITY-EXCEPTIONS.md | A | — | .../TEAM-CAPACITY-EXCEPTIONS.md | TEMPLATE | — | — | — | — |
| **Bloco B — Architecture & Security & Specialists** | | | | | | | | |
| ARCHITECTURE-DEFINITION.md | B | F7 | .../ARCHITECTURE-DEFINITION.md | PENDENTE | — | — | — | — |
| SECURITY-DEFINITION.md | B | F8 | .../SECURITY-DEFINITION.md | PENDENTE | — | — | — | — |
| DATA-ARCHITECTURE-DEFINITION.md | B | F9 | .../DATA-ARCHITECTURE-DEFINITION.md | PENDENTE | — | — | — | — |
| DEVOPS-SRE-DEFINITION.md | B | F10 | .../DEVOPS-SRE-DEFINITION.md | PENDENTE | — | — | — | — |
| TEST-STRATEGY-DEFINITION.md | B | F11 | .../TEST-STRATEGY-DEFINITION.md | PENDENTE | — | — | — | — |
| INFRA-CLOUD-DEFINITION.md | B | F12 | .../INFRA-CLOUD-DEFINITION.md | PENDENTE | — | — | — | — |
| **Bloco C — Catálogo, Matriz, Stack, Specs & Milestones** | | | | | | | | |
| SOLUTIONS-CATALOG.md | C | F13 | .../SOLUTIONS-CATALOG.md | PENDENTE | — | — | — | — |
| SOLUTIONS-MATRIX.md | C | F14 | .../SOLUTIONS-MATRIX.md | PENDENTE | — | — | — | — |
| SOLUTIONS-STACK-MATRIX.md | C | F15 | .../SOLUTIONS-STACK-MATRIX.md | PENDENTE | — | — | — | — |
| SPECS-DEFINITION.md | C | F16 | .../SPECS-DEFINITION.md | PENDENTE | — | — | — | — |
| MILESTONES.md | C | F17 | .../MILESTONES.md | PENDENTE | — | — | — | — |
| **Bloco D — Sprints — Technical Discovery** | | | | | | | | |
| 580-SPRINT-BACKLOG.md | D | F18 | technical-discovery/580-SPRINT-BACKLOG.md | PENDENTE | — | — | — | — |
| DISCOVERY TÉCNICO (590-sprint-NNN/) | D | F19 | technical-discovery/590-sprint-NNN/ | PENDENTE | — | — | — | — |

**2. Registro de Ciclos Gate→Fix:**

| Timestamp | Documento | Bloco | Ciclo | NCs Encontradas | NCs Resolvidas | Veredito |
|---|---|---|---|---|---|---|
| (preenchido a cada iteração) |

**3. Indicadores (atualizados a cada fase):**
- Total de documentos: 20
- % COMPLIANCE: X%
- % Pendente: Y%
- NCs abertas: Z
- Fase atual: N
- Bloco atual: 0 / A / B / C / D

**4. Timeline Visual por Bloco:**
```
Bloco 0 [████████████████████] 4/4 COMPLIANCE ✅
Bloco A [████████░░░░░░░░░░░░] 2/3 GATE ⚠️
Bloco B [░░░░░░░░░░░░░░░░░░░░] 0/6 PENDENTE
Bloco C [░░░░░░░░░░░░░░░░░░░░] 0/5 PENDENTE
Bloco D [░░░░░░░░░░░░░░░░░░░░] 0/2 PENDENTE
```

### Atualização Incremental (executado após cada fase)

Após cada fase ser concluída (COMPLIANCE), atualizar:
1. Linha do documento na tabela-mestra (estado, datas, veredito)
2. Registro de ciclos Gate→Fix (se houve)
3. Indicadores (% compliance, NCs)
4. Timeline visual por bloco

### Atualização após Gate (executado quando um gate emite veredito)

Após cada execução de gate (independentemente do veredito):
1. Atualizar estado do documento (GATE, NÃO-COMPLIANCE, PRE-COMPLIANCE)
2. Se NÃO COMPLIANCE: registrar NCs encontradas (relatório inline do gate)
3. Atualizar indicadores

---

## Observações

- Este artefato NÃO tem gate próprio — sua correção é a própria execução das fases
- Posicionado como artefato standalone após a Barreira D (Bloco D 100% COMPLIANCE)
- A revisão humana ao final do roadmap deve inspecionar este documento para confirmar que todos os artefatos estão COMPLIANCE
- O Execution History pode ser consultado a qualquer momento para saber "em que pé está o roadmap"
- Se o roadmap for interrompido e retomado, este arquivo é a referência para determinar o ponto de continuação

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `changelog-automation` | Automatizar registro de mudanças | Tracking |
| 2 | `documentation-writer` | Manter documento atualizado | Documentação |
| 3 | `track-management` | Gerenciar tracking de estado | Tracking |
| 4 | `kpi-dashboard-design` | Desenhar indicadores visuais | Métricas |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador do execution history | Time de Arquitetura |
| 2.0 | 30/07/2026 | Reformulação: posicionamento standalone pós-Barra D; tracking expandido para 20 artefatos com INTAKE-LOG, DOR-ASSESSMENT, PRODUCT-BACKLOG-LIST, DATA-ARCHITECTURE, DEVOPS-SRE, TEST-STRATEGY, INFRA-CLOUD, SPRINT-BACKLOG, TECHNICAL-DISCOVERY | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
