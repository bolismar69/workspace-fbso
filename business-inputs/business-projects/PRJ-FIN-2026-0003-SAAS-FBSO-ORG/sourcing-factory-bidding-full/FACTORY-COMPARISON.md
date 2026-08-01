# FACTORY-COMPARISON — Matriz Comparativa de Fábricas (Full Mode)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo (62 US · 18 Features · 4 Épicos)
- **Data:** 31/07/2026
- **Baseline PERT:** 7,300h / 46 h-m (c/ 20% contingência)
- **Status:** ⚠️ NENHUMA FÁBRICA APROVADA — RFQ será reaberto

---

## 1. Ranking por Horas Totais

| # | Fábrica | Total Horas | vs PERT | QA% | Arch% | Prazo | Time | Consistência | PIB 🆕 | Veredito |
|:---|:---|---:|---:|---:|---:|---:|---:|:---|:---:|:---|
| 1 | Capgemini | 11,680h | +60% | 11.0% | 11.0% | 3m | 15p | 🔴 +62% | 0.40 🔴 | 🔴 Rejeitada |
| 2 | Deloitte | 11,680h | +60% | 11.0% | 11.0% | 3m | 15p | 🔴 +62% | 0.40 🔴 | 🔴 Rejeitada |
| 3 | Infosys | 11,680h | +60% | 11.0% | 11.0% | 3m | 15p | 🔴 +62% | 0.40 🔴 | 🔴 Rejeitada |
| 4 | Overlabs | 27,680h | +279% | 4.6% | 4.6% | 4m | 45p | ✅ -4% | 0.00 🔴 | 🔴 Rejeitada |
| 5 | Stefanini | 28,030h | +284% | 4.6% | 4.6% | 3m | 15p | 🔴 +289% | 0.00 🔴 | 🔴 Rejeitada |
| 6 | CI&T | 51,680h | +608% | 2.5% | 2.5% | 6m | 45p | ⚠️ +20% | 0.00 🔴 | 🔴 Rejeitada |
| 7 | TOTVS | 91,680h | +1,156% | 1.4% | 1.4% | 6m | 80p | ⚠️ +19% | 0.00 🔴 | 🔴 Rejeitada |
| 8 | TCS | 131,680h | +1,704% | 1.0% | 1.0% | 9m | 93p | ✅ -2% | 0.00 🔴 | 🔴 Rejeitada |

---

## 2. Visualização Comparativa

```
Horas Totais por Fábrica (escala log)

PERT Baseline  ██ 7,300h
Capgemini      ███ 11,680h
Deloitte       ███ 11,680h  ← IDÊNTICOS
Infosys        ███ 11,680h
Overlabs       ███████ 27,680h
Stefanini      ███████ 28,030h
CI&T           █████████████ 51,680h
TOTVS          ██████████████████████ 91,680h
TCS            █████████████████████████████████ 131,680h
```

### Nota sobre PIB (Proximity to Internal Baseline) 🆕

O PIB Score mede a proximidade da estimativa da fábrica em relação à baseline interna PERT Downstream (7,300h). Fórmula: `PIB = 1 − (|Factory − 7,300| / 7,300)`. Escala: 1.0 = match exato, 0.0 = 2× ou mais de desvio. **Nenhuma fábrica atingiu PIB ≥ 0.50** — todas as 8 estão com PIB Score abaixo do threshold de alerta, confirmando o descolamento generalizado da baseline interna.

---

## 3. Análise por Cluster

### Cluster A — Estimativas Próximas ao PERT (11,680h)
**Capgemini, Deloitte, Infosys**

| Característica | Valor |
|:---|---:|
| Horas | 11,680h (+60% vs PERT) |
| QA / Arch | 11.0% / 11.0% (insuficiente) |
| Prazo declarado | 3 meses com 15 pessoas |
| Prazo calculado | 4.9 meses (divergência +62%) |

⚠️ **Alerta:** As 3 fábricas submeteram valores **idênticos** em todas as colunas. Possível:
- Uso do mesmo template/preenchimento automático
- Coordenação entre fábricas
- Erro de interpretação do schema (valores placeholder não substituídos)

### Cluster B — Estimativas Intermediárias (27-28k)
**Overlabs, Stefanini**

| Característica | Overlabs | Stefanini |
|:---|---:|---:|
| Horas | 27,680h | 28,030h |
| QA / Arch | 4.6% / 4.6% | 4.6% / 4.6% |
| Consistência Prazo | ✅ -4% | 🔴 +289% |

Overlabs foi a única com consistência Prazo×Horas dentro da margem aceitável, mas QA/Arch insuficientes.

### Cluster C — Superestimadas (51-131k)
**CI&T, TOTVS, TCS**

Superestimaram em 6-17× o PERT baseline. TCS com 131,680h e 93 pessoas é inviável para um projeto deste porte.

---

## 4. Problemas Sistêmicos Identificados

| Problema | Fábricas Afetadas | Impacto |
|:---|:---|:---|
| QA/Arch como overhead fixo (320h) em vez de proporcional | Todas (8/8) | 🔴 Eliminatório |
| Comentários genéricos sem racional de estimativa | Todas (8/8) | 🟡 Transparência |
| Valores idênticos entre fábricas (possível template compartilhado) | Capgemini, Deloitte, Infosys (3/8) | 🔴 Independência |
| Divergência Prazo×Horas >50% | Stefanini, Capgemini, Deloitte, Infosys (4/8) | 🔴 Consistência |

---

## 5. Recomendação

**NENHUMA fábrica atende aos critérios DTA.** Recomenda-se:

1. **Reabrir RFQ** com instruções explícitas:
   - `horas_qa` deve ser ≥ 25% de `total_horas`
   - `horas_arch` deve ser ≥ 5% de `total_horas`
   - `total_horas` = soma de todas as colunas de horas (validado automaticamente)
   - Preencher `comentarios` com o racional detalhado da estimativa (metodologia, premissas, justificativas por épico)

2. **Adicionar fórmula de validação** no CSV para evitar erros de preenchimento

3. **Investigação:** Contatar Capgemini, Deloitte e Infosys sobre a identidade dos valores submetidos

4. **Prazo:** Reabrir RFQ com novo prazo de 5 dias úteis

---

## 6. Comparação com Rodada Discovery

| Métrica | Discovery (ROM) | Full (PERT) |
|:---|---:|:---|
| Baseline | 6,080h (ROM ±50%) | 7,300h (PERT ±20%) |
| Fábricas aprovadas | 2/8 (Stefanini, CI&T) | **0/8** |
| Principal motivo de rejeição | QA/Arch + Prazo irreal | QA/Arch (sistêmico) |
| Menor estimativa | Stefanini 16,000h | Capgemini/Deloitte/Infosys 11,680h |

> ℹ️ **Nota:** A rodada Discovery usou escopo macro (épicos) e ROM ±50%. A rodada Full usa escopo completo (62 US) e PERT ±20% como baseline. A rejeição unânime na rodada Full sugere que as fábricas não adaptaram suas metodologias de estimativa ao nível de detalhamento solicitado.

---

🤖 *Matriz Comparativa — Fase 6 do Sourcing & Factory Bidding (Full Mode). Nenhuma fábrica aprovada. RFQ será reaberto.*
