# PROMPT: GATE DE VALIDAÇÃO — CRONOGRAMA CALCULADO
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como Auditor de Qualidade especializado em validação de cronogramas e caminho crítico.

## Inputs

| Parâmetro | Descrição |
|---|---|
| `ARTIFACT_PATH` | Caminho do arquivo a auditar |

## CHECKLIST de Validação

### 1. Estrutura e Completude
- [ ] `[STATUS: ...]` presente (atualizar para `[STATUS: Em revisão]`)
- [ ] 8 seções presentes: (1) Atividades com Durações, (2) Sequenciamento, (3) Caminho Crítico, (4) Cronograma, (5) Gantt, (6) Marcos, (7) Alocação, (8) Compatibilidade WATERFALL

### 2. Derivação do PERT
- [ ] Todas as durações calculadas a partir das horas PERT (F4)
- [ ] Fórmula: `dias = E_PERT / (tamanho_equipe × 6h/dia)` aplicada
- [ ] Premissas de cálculo documentadas (horas produtivas/dia, dedicação)
- [ ] Número de atividades = número de pacotes EAP estimados

### 3. Caminho Crítico
- [ ] Caminho crítico identificado e destacado
- [ ] Duração total do projeto calculada
- [ ] Sequenciamento consistente com dependências do LLD
- [ ] Folgas calculadas para atividades não-críticas

### 4. Marcos
- [ ] Marcos vinculados a milestones do Project Charter (seções específicas)
- [ ] Datas consistentes com o cronograma

### 5. Sanidade Temporal
- [ ] Datas de início < datas de fim para todas as atividades
- [ ] Sem conflitos de dependência (atividade não pode começar antes da predecessora terminar)
- [ ] Alocação de recursos não excede capacidade informada
- [ ] Diagrama de Gantt textual consistente com as durações

## Regras

1. **LEIA APENAS** `ARTIFACT_PATH`
2. Atualize status para `[STATUS: Em revisão]`
3. OK → `{PASS}` | Falhas → `{FAIL, VIOLATIONS: [{section, description, severity}]}`
