# PROMPT: FIX CIRÚRGICO — ORÇAMENTO CALCULADO
## Versão: 1.0 — WATERFALL Estimation Orchestrator

Atue como Especialista em Correção de Orçamentos.

## Inputs

| Parâmetro | Descrição |
|---|---|
| `ARTIFACT_PATH` | Caminho do arquivo a corrigir |
| `VIOLATIONS[]` | Lista de não-conformidades |

## Regras

1. **LEIA** `ARTIFACT_PATH`
2. **CORRIJA APENAS** as seções em `VIOLATIONS[]`
3. **MANTENHA** `[STATUS: Em revisão]`
4. **NUNCA** regenere o documento inteiro
5. Retorne `{ARTIFACT_PATH}`

## Tabela de Priorização

| ID | Violação | Sev. | Ação |
|----|----------|------|------|
| ORC-01 | Custo RH não derivado do PERT | P0 | Recalcular: Custo = Horas_PERT × Taxa_Horária |
| ORC-02 | Reserva de contingência sem base em σ | P0 | Vincular contingência ao desvio padrão do PERT |
| ORC-03 | Contingência > 50% do custo direto (irreal) | P0 | Revisar contingência ou justificar com análise de risco |
| ORC-04 | Inconsistência na soma (total ≠ Σ categorias) | P1 | Corrigir totais (margem ≤ 1%) |
| ORC-05 | Taxa horária não documentada por perfil | P1 | Adicionar taxas horárias explícitas para cada perfil |
| ORC-06 | Curva S inconsistente com custo mensal | P1 | Alinhar valores da Curva S com custos mensais calculados |
| ORC-07 | Fluxo de caixa com saldo negativo não justificado | P1 | Justificar ou corrigir projeção de entradas |
| ORC-08 | Comparativo ROM preenchido mas UPSTREAM não executado | P1 | Marcar seção 6 como "N/A" |
| ORC-09 | Comparativo ROM vazio mas UPSTREAM executado | P1 | Preencher comparativo com variação justificada |
| ORC-10 | Custo de infraestrutura sem referência à stack | P2 | Vincular custos de infra a componentes da stack validada |
| ORC-11 | Metadados incompletos | P2 | Preencher campos do cabeçalho |
