# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RISK-ADJUSTED-ESTIMATE (F10)

## Contexto

Gate de Validação da Estimativa Ajustada a Risco (`RISK-ADJUSTED-ESTIMATE.md`). Verifica matriz de riscos, cenários e análise de sensibilidade.

**Princípio fundamental:** Ajuste aplicado sobre a estimativa PERT (F8), nunca sobre ROM upstream. Análise de sensibilidade identifica riscos de maior impacto.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

**Arquivo auditado:** `RISK-ADJUSTED-ESTIMATE.md`
**Referências:** F8 (PERT), F9 (Resource Allocation), artefatos Detail-Level

## Dimensões de Validação

### Dimensão 1: Fidelidade
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Base é PERT | Ajuste sobre F8, não sobre ROM |
| 1.2 | Riscos dos artefatos | Riscos extraídos de F2-F7 e Charter |

### Dimensão 2: Quantificação
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Probabilidade × Impacto | Cada risco com valores numéricos |
| 2.2 | Valor esperado | Σ(Prob × Impacto) calculado |
| 2.3 | 3 cenários | Conservador, PERT, Pessimista |

### Dimensão 3: Análise
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Sensibilidade | Top 3 riscos por impacto |
| 3.2 | Recomendações | Ações de mitigação priorizadas |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE RISCO: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-RISK-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A matriz de riscos cobre todas as ameaças relevantes ao projeto?
2. Deseja enviar mais documentos de risco?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 10 do Downstream Architecture Refinement*
