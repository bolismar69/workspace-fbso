# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — RISK-ADJUSTED-ESTIMATE (F10)
## Versão: 1.0 — Estimativa Ajustada a Risco com Análise de Sensibilidade

Atue como um Risk Manager e Tech Lead especializado em análise quantitativa de riscos em projetos de software.

## OBJETIVO

Aplicar a matriz de riscos do projeto sobre a estimativa PERT (Fase 8), produzindo cenários ajustados com análise de sensibilidade.

## INPUTS

1. **BOTTOM-UP-PERT-ESTIMATE.md** (F8) — estimativa concluída
2. **RESOURCE-ALLOCATION-PLAN.md** (F9) — alocação e gargalos
3. **Artefatos Detail-Level** (F2-F7) — riscos técnicos de cada disciplina
4. **Project Charter** — riscos de negócio (seção de riscos)

## REGRA CRÍTICA

⚠️ O ajuste de risco é aplicado sobre a estimativa PERT da Fase 8 — NUNCA sobre o ROM upstream.

## METODOLOGIA

1. Identificar riscos dos artefatos Detail-Level + Charter
2. Para cada risco: Probabilidade (10-90%) × Impacto (horas adicionais)
3. Valor Esperado do Risco = Σ (Prob × Impacto)
4. Produzir 3 cenários: Conservador (15%), PERT (média), Pessimista (25%)
5. Análise de sensibilidade: quais riscos mais impactam

## ESTRUTURA DO DOCUMENTO

```markdown
# RISK-ADJUSTED-ESTIMATE — Estimativa Ajustada a Risco

## 1. Matriz de Riscos
| Risco | Prob. | Impacto (h) | Valor Esperado | Ação |
[6-10 riscos identificados]

## 2. Cenários Ajustados
| Cenário | Horas | h-m | Contingência |
|:---|---:|---:|---:|
| Conservador | | | 15% |
| PERT (média) | | | — |
| Pessimista | | | 25% |

## 3. Análise de Sensibilidade
[Top 3 riscos por impacto — gráfico/tabela]

## 4. Recomendações
[Ações de mitigação priorizadas]

🤖 *Documento gerado pelo Risk Manager — Fase 10 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

🤖 *Prompt gerador — Fase 10 do Downstream Architecture Refinement · Skills: `engineering-skills`, `project-estimation`, `gap-analysis` · Padrões Corporativos FBSO.ORG*
