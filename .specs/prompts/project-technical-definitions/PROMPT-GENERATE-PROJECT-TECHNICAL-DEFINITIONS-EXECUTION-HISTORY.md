# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY

## Contexto

Este prompt gerencia o artefato `PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md` — o **dashboard de controle** que registra o estado de CADA documento gerado em TODAS as fases do roadmap. Diferente dos outros artefatos, este NÃO tem gate/fix próprios — ele é **atualizado incrementalmente** após cada fase ser concluída.

**Princípio fundamental:** O Execution History é o single source of truth para o status do roadmap. Qualquer pessoa deve conseguir abrir este arquivo e saber exatamente em que estado está cada documento.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Esquema de Estados

```
CREATED → GATE ⟷ FIX → PRE-COMPLIANCE → COMPLIANCE
  │         │              │
  │         └─ NÃO-COMPLIANCE (com FAIL_REPORT)
  │
  └─ (estado inicial após geração)
```

---

## Fluxo de Execução

### Criação Inicial (executado uma vez após o Bootstrap)

Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-EXECUTION-HISTORY.md` com:

**1. Tabela-Mestra de Documentos:**

| Documento | Fase | Caminho | Estado | Data Criação | Última Atualização | NCs Pendentes | Veredito Gate |
|---|---|---|---|---|---|---|---|
| PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md | — | .../TEAM-CAPACITY.md | MIGRADO | — | — | — | — |
| PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md | — | .../TEAM-CAPACITY-EXCEPTIONS.md | TEMPLATE | — | — | — | — |
| PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md | 1 | .../TEAM-MAP.md | PENDENTE | — | — | — | — |
| PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md | 2 | .../SOLUTIONS-CATALOG.md | PENDENTE | — | — | — | — |
| ... (todos os 11 artefatos + capacity exceptions) |

**2. Registro de Ciclos Gate→Fix:**

| Timestamp | Documento | Ciclo | NCs Encontradas | NCs Resolvidas | Veredito |
|---|---|---|---|---|---|
| (preenchido a cada iteração) |

**3. Indicadores (atualizados a cada fase):**
- Total de documentos: N
- % COMPLIANCE: X%
- % Pendente: Y%
- NCs abertas: Z
- Fase atual: N

**4. Timeline Visual:**
```
Fase 1  [████████████] COMPLIANCE ✅
Fase 2  [████████░░░░] GATE ⚠️
Fase 3  [░░░░░░░░░░░░] PENDENTE
...
```

### Atualização Incremental (executado após cada fase)

Após cada fase ser concluída (COMPLIANCE), atualizar:
1. Linha do documento na tabela-mestra (estado, datas, veredito)
2. Registro de ciclos Gate→Fix (se houve)
3. Indicadores (% compliance, NCs)
4. Timeline visual

### Atualização após Gate (executado quando um gate emite veredito)

Após cada execução de gate (independentemente do veredito):
1. Atualizar estado do documento (GATE, NÃO-COMPLIANCE, PRE-COMPLIANCE)
2. Se REPROVADO: registrar NCs encontradas, link para FAIL_REPORT
3. Atualizar indicadores

---

## Observações

- Este artefato NÃO tem gate próprio — sua correção é a própria execução das fases
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

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
