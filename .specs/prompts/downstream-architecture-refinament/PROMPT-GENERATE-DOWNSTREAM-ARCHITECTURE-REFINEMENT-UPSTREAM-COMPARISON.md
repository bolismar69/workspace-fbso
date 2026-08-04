# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — UPSTREAM-COMPARISON (F12)
## Versão: 1.0 — Relatório Comparativo com ROM Upstream (Condicional)

Atue como um Engineering Manager e Tech Lead especializado em análise comparativa de estimativas e comunicação com comitês de governança.

## OBJETIVO

Gerar um relatório comparativo entre a estimativa PERT downstream (Fase 8) e o ROM upstream (se existir). Este é o ÚNICO vínculo entre os dois roadmaps.

## CONDICIONALIDADE

⚠️ **Esta fase é condicional.** Só executa se:
1. O diretório `upstream-architecture-discovery/` existir
2. O arquivo `DISCOVERY-LEVEL-ROM-ESTIMATE.md` existir dentro dele

Se NÃO existir: gerar documento informando "Upstream discovery não encontrado — sem ROM para comparação. Estimativa PERT é a única baseline disponível." e marcar como COMPLIANCE.

## REGRA CRÍTICA

⚠️ Este relatório NÃO ALTERA a estimativa PERT. A estimativa da Fase 8 já está congelada e aprovada na Barreira B. O propósito é puramente informativo para o comitê.

## INPUTS

1. **BOTTOM-UP-PERT-ESTIMATE.md** (F8) — estimativa downstream (CONGELADA, NÃO ALTERAR)
2. **DISCOVERY-LEVEL-ROM-ESTIMATE.md** (upstream) — se existir
3. **SCOPE-SNAPSHOT.md** (F11) — escopo incluído

## ESTRUTURA DO DOCUMENTO

```markdown
# UPSTREAM-COMPARISON-REPORT — Comparativo ROM Upstream × PERT Downstream

- **ATENÇÃO:** Este relatório é informativo. A estimativa PERT (Fase 8) é a baseline aprovada e NÃO é alterada por esta análise.

## 1. Status do Upstream
[Upstream encontrado: lista de artefatos] OU [Upstream NÃO encontrado — sem comparação possível]

## 2. Tabela Comparativa (se upstream existir)
| Nível | ROM Upstream | PERT Downstream | Desvio |
|:---|---:|---:|---:|
| Projeto | | | ±X% |
| EP-0001 | | | |
| EP-0002 | | | |
| EP-0003 | | | |
| EP-0004 | | | |

## 3. Análise de Convergência/Divergência
- PERT dentro da faixa ROM (±50%): ✅ Convergência esperada — o ROM estava preciso
- PERT fora da faixa ROM: ⚠️ Divergência — analisar causas

## 4. Causas de Divergência (se aplicável)
- Escopo adicional? (US adicionadas após o ROM)
- Complexidade subestimada no ROM?
- Diferenças metodológicas? (top-down vs bottom-up)

## 5. Visualização Comparativa
[Gráfico de barras proporcionais]

## 6. Conclusão
[A estimativa PERT refina/substitui/diverge do ROM? Recomendação para o comitê.]

## 7. Nota de Imutabilidade
A estimativa PERT (Fase 8) permanece inalterada. Este relatório é um artefato de governança para apoiar a decisão do comitê.

🤖 *Documento gerado pelo Engineering Manager / Tech Lead — Fase 12 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

🤖 *Prompt gerador — Fase 12 do Downstream Architecture Refinement · Cross-Check Condicional · Skills: `engineering-skills`, `gap-analysis`, `project-estimation` · Padrões Corporativos FBSO.ORG*
