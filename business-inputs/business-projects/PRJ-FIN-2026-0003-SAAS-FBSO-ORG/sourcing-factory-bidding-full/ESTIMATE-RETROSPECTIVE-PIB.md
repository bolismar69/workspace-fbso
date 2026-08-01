# ESTIMATE-RETROSPECTIVE-PIB — Análise Retrospectiva com PIB v1.1 (Full Mode)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Baseline PIB:** PERT Downstream 7,300h (recomendado)
- **Objetivo:** Identificar observações adicionais para realinhamento com fábricas

---

## 1. PIB por Épico — Qual Fábrica Mais se Aproxima da Baseline?

| Épico | PERT Baseline | Melhor Fábrica | Horas | Desvio | PIB Score |
|:---|---:|:---|---:|---:|:---:|
| EP-0001 Portal Admin | 562h | Capgemini | 2,920h | +420% | 0.00 |
| EP-0002 Clientes | 1,374h | Capgemini | 2,920h | +113% | 0.00 |
| EP-0003 RBAC | 1,437h | Capgemini | 2,920h | +103% | 0.00 |
| EP-0004 Portal Cliente | 1,843h | Capgemini | 2,920h | +58% | 0.42 |

> 🔴 **Conclusão:** Mesmo a melhor fábrica por épico (Capgemini) está 58-420% acima da baseline PERT. Nenhuma fábrica demonstrou compreensão adequada do esforço relativo entre épicos — o EP-0003 (RBAC, 16 US, complexidade alta) deveria custar mais que o EP-0001 (Dashboard, 7 US), mas várias fábricas trataram todos os épicos com esforço idêntico.

---

## 2. Observação Crítica: Flat Estimates (Estimativas Planas)

**3 fábricas (Capgemini, Deloitte, Infosys)** estimaram **exatamente o mesmo valor para todos os 4 épicos**: 2,920h cada.

| Fábrica | EP-0001 | EP-0002 | EP-0003 | EP-0004 | Total |
|:---|---:|---:|---:|---:|---:|
| Capgemini | 2,920h | 2,920h | 2,920h | 2,920h | 11,680h |
| Deloitte | 2,920h | 2,920h | 2,920h | 2,920h | 11,680h |
| Infosys | 2,920h | 2,920h | 2,920h | 2,920h | 11,680h |

> 🔴 **Problema:** É impossível que 4 épicos com complexidades diferentes (7 vs 16 vs 16 vs 23 US) tenham exatamente o mesmo esforço. Isso indica que as fábricas **não analisaram o escopo por épico** — aplicaram um rateio uniforme sem considerar a complexidade relativa. Esse comportamento é incompatível com a metodologia bottom-up solicitada no RFQ.

**Outras fábricas também mostram pouca variação entre épicos:** Stefanini (5,820-8,320h, variação de apenas 43%), Overlabs (6,920h flat), TCS (32,920h flat), TOTVS (22,920h flat). Apenas CI&T e Stefanini mostraram alguma diferenciação.

---

## 3. Observação Crítica: QA/Arch como Overhead Fixo

**Todas as 8 fábricas** trataram QA e Arquitetura como custo fixo por épico (~320h cada), independente do total de horas:

| Fábrica | Total | QA Real | QA Mínimo (25%) | Arch Real | Arch Mínimo (5%) |
|:---|---:|---:|---:|---:|---:|
| Capgemini | 11,680h | 1,280h (11%) | 2,920h | 1,280h (11%) | 584h |
| Stefanini | 28,030h | 1,280h (5%) | 7,008h | 1,280h (5%) | 1,402h |
| CI&T | 51,680h | 1,280h (2%) | 12,920h | 1,280h (2%) | 2,584h |
| TCS | 131,680h | 1,280h (1%) | 32,920h | 1,280h (1%) | 6,584h |

> 🔴 **Problema:** QA e Arquitetura devem ser **proporcionais** ao esforço de desenvolvimento, não valores fixos. Um épico de 32,920h (TCS) deveria ter ~8,230h de QA — não 320h. Isso indica que as fábricas preencheram QA/Arch como placeholder, sem análise real.

---

## 4. Observação Crítica: Comentários Genéricos

**Todas as 8 fábricas** usaram exatamente o mesmo texto na coluna `comentarios`:

> *"seguimos especificamente o material reportado"*

E na coluna `premissas`:

> *"atende exclusivamente o escopo reportado"*

> 🔴 **Problema:** O schema CSV solicita que a coluna `comentarios` contenha o **racional detalhado da estimativa** — metodologia utilizada, premissas por épico, justificativas para os números. Textos genéricos e idênticos entre fábricas diferentes indicam que essa coluna não foi levada a sério. A FBSO.ORG não tem como avaliar a qualidade do raciocínio por trás dos números.

---

## 5. Observação Crítica: Independência Comprometida

**3 fábricas** submeteram valores **idênticos** em TODAS as colunas:

| Coluna | Capgemini | Deloitte | Infosys |
|:---|---:|---:|---:|
| Total Horas | 11,680h | 11,680h | 11,680h |
| horas_dev | 1,600h × 4 | 1,600h × 4 | 1,600h × 4 |
| horas_qa | 320h × 4 | 320h × 4 | 320h × 4 |
| horas_arch | 320h × 4 | 320h × 4 | 320h × 4 |
| complexidade | Alta/Alta/Media/Media | Alta/Alta/Media/Media | Alta/Alta/Media/Media |
| comentarios | idêntico | idêntico | idêntico |

> 🔴 **Problema:** A probabilidade de 3 fábricas independentes chegarem exatamente aos mesmos valores para 18 colunas é efetivamente zero. Isso sugere: (a) uso de template compartilhado sem revisão individual, (b) coordenação entre fábricas, ou (c) preenchimento automatizado sem análise. Em qualquer caso, a independência do processo de bidding está comprometida.

---

## 6. Novos Alinhamentos para Realinhamento

Além dos pontos já comunicados (QA≥25%, Arch≥5%), as fábricas devem ser orientadas sobre:

| # | Observação | Ação para Realinhamento |
|:---|:---|:---|
| 1 | **Flat estimates** | Exigir diferenciação de esforço entre épicos — justificar por que um épico de 23 US custa o mesmo que um de 7 US |
| 2 | **QA/Arch fixo** | QA e Arch devem ser **percentuais** do esforço de desenvolvimento, não valores absolutos |
| 3 | **Comentários genéricos** | Exigir racional detalhado: metodologia, premissas, justificativas por épico |
| 4 | **Independência** | Notificar Capgemini, Deloitte e Infosys sobre a identidade dos valores — solicitar explicação formal |
| 5 | **PIB por épico** | Informar que a baseline interna será usada para avaliar proximidade por épico, não apenas total |

---

## 7. Recomendação para o Próximo Ciclo

1. **Atualizar o RFQ-PACKAGE.md** com instruções EXPLÍCITAS sobre:
   - Proporcionalidade QA/Arch (não usar valores fixos)
   - Diferenciação obrigatória entre épicos
   - Comentários com racional detalhado (mínimo 200 caracteres por linha)

2. **Adicionar validação de variância entre épicos** como regra DTA:
   - Se todos os épicos têm o mesmo total de horas → ⚠️ Alerta (flat estimate)
   - Se CV (coeficiente de variação) entre épicos < 10% → 🔍 Revisão manual

3. **Solicitar explicação formal** de Capgemini, Deloitte e Infosys sobre a identidade dos valores

4. **Incluir PIB por épico** nos arquivos de validação individual para a próxima rodada

---

🤖 *Análise retrospectiva com PIB v1.1 — Full Mode. Baseada nos 8 CSVs recebidos em 31/07/2026.*
