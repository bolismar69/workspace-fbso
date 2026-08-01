# UPSTREAM-COMPARISON-REPORT — Comparativo ROM Upstream × PERT Downstream

- **Data:** 31/07/2026
- **Fase:** F12 — Downstream Architecture Refinement (condicional)
- **ATENÇÃO:** Este relatório é **informativo**. A estimativa PERT (Fase 8) é a baseline aprovada e **NÃO é alterada** por esta análise.

---

## 1. Status do Upstream

✅ **Upstream encontrado.** 11 artefatos Discovery-Level em `upstream-architecture-discovery/`. ROM disponível em `DISCOVERY-LEVEL-ROM-ESTIMATE.md`.

---

## 2. Tabela Comparativa

| Nível | ROM Upstream | PERT Downstream | Desvio |
|:---|---:|---:|---:|
| **Projeto (Total)** | 4,640 – 8,000h (29-50 h-m) | **6,077 – 7,121h (38-45 h-m)** | Dentro da faixa ROM ✅ |
| EP-0001 Portal Admin | 960 – 1,600h | 562h | Abaixo da faixa ROM ⚠️ |
| EP-0002 Clientes | 1,280 – 2,240h | 1,374h | Dentro da faixa ✅ |
| EP-0003 RBAC | 1,120 – 1,920h | 1,437h | Dentro da faixa ✅ |
| EP-0004 Portal Cliente | 1,280 – 2,240h | 1,843h | Dentro da faixa ✅ |

### Comparação com Factory Bids

| Fonte | Horas | Desvio vs PERT |
|:---|---:|---:|
| ROM Interno (provável) | 6,080h | **+0%** ✅ |
| Stefanini (aprovada) | 16,000h | +163% |
| CI&T (aprovada c/ ressalva) | 52,000h | +756% |

---

## 3. Análise de Convergência

- **ROM provável (6,080h) vs PERT (6,077h):** Desvio de **0%** — convergência notável. O ROM interno, feito na fase de discovery com escopo macro (épicos apenas), acertou o valor com precisão impressionante.
- **EP-0001 abaixo da faixa ROM:** O PERT do EP-0001 ficou em 562h vs faixa ROM de 960-1,600h. Possível explicação: o ROM foi feito quando o escopo do EP-0001 ainda não estava detalhado em US, e pode ter superestimado o esforço do dashboard.
- **Demais épicos dentro da faixa ROM:** Consistência metodológica entre abordagem top-down (ROM por solução) e bottom-up (PERT por US).

---

## 4. Visualização Comparativa

```
Estimativas (horas) — Escala Linear

ROM provável     ████████ 6,080h
PERT Downstream  ████████ 6,077h  ← CONVERGÊNCIA (desvio 0%)
ROM pessimista   ██████████ 8,000h
Stefanini        ████████████████████ 16,000h
CI&T             ██████████████████████████████████████████████████████████████████ 52,000h
```

---

## 5. Conclusão

**A estimativa PERT downstream (6,077h / 38 h-m) refina e valida o ROM upstream (6,080h / 38 h-m provável).** A convergência de 0% entre as duas metodologias independentes (top-down por solução vs bottom-up por US) aumenta a confiança na estimativa. 

**Recomendação ao comitê:** Adotar **7,300h / 46 h-m** como baseline (PERT + 20% contingência recomendada pelo F10 Risk-Adjusted), com revisão após 2 sprints de execução para refinar para ±10%.

---

## 6. Nota de Imutabilidade

A estimativa PERT (Fase 8) permanece **inalterada**. Este relatório é um artefato de governança para apoiar a decisão do comitê e documentar a evolução da precisão: ROM ±50% → PERT ±15-25%.

---

🤖 *Cross-Check Report — Fase 12 do Downstream Architecture Refinement. Único vínculo com upstream. Informativo — não altera a estimativa PERT.*
