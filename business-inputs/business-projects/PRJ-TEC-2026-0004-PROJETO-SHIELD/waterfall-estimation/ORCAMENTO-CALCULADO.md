# ORÇAMENTO CALCULADO — DERIVADO DO PERT: PRJ-TEC-2026-0004-PROJETO-SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Estimativa Base** | WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md v1.0 (929,2h) |
| **Cronograma Base** | CRONOGRAMA-CALCULADO.md v1.0 (39 dias úteis / 7,7 semanas) |
| **Data de Elaboração** | 04/08/2026 |
| **Versão** | 1.0 |

---

### 1. Custos por Recurso

#### 1.1 Recursos Humanos

| Perfil | Horas PERT (E) | Taxa Horária (R$) | Custo (R$) |
|--------|---------------|-------------------|------------|
| Sênior (×2) | 464,6 | R$ 150,00 | R$ 69.690,00 |
| Pleno (×1) | 232,3 | R$ 100,00 | R$ 23.230,00 |
| Júnior (×1) | 232,3 | R$ 60,00 | R$ 13.938,00 |
| **Subtotal RH** | **929,2** | | **R$ 106.858,00** |

#### 1.2 Infraestrutura (Cloud)

| Recurso | Custo Mensal (R$) | Duração | Custo (R$) |
|---------|-------------------|---------|------------|
| DOKS Cluster (3 nodes) | R$ 1.200,00 | 2 meses | R$ 2.400,00 |
| PostgreSQL Managed | R$ 600,00 | 2 meses | R$ 1.200,00 |
| Redis Managed | R$ 400,00 | 2 meses | R$ 800,00 |
| App Platform (SPA) | R$ 200,00 | 2 meses | R$ 400,00 |
| **Subtotal Infra** | | | **R$ 4.800,00** |

#### 1.3 Licenças e Ferramentas

| Ferramenta | Custo | Duração | Custo (R$) |
|-----------|-------|---------|------------|
| Cloudflare Pro | R$ 100,00/mês | 2 meses | R$ 200,00 |
| GitHub Actions (minutos) | R$ 0,00 (free tier) | — | R$ 0,00 |
| **Subtotal Licenças** | | | **R$ 200,00** |

#### 1.4 Consolidação

| Categoria | Custo (R$) | % do Total |
|-----------|------------|-----------|
| Recursos Humanos | R$ 106.858,00 | 95,5% |
| Infraestrutura | R$ 4.800,00 | 4,3% |
| Licenças | R$ 200,00 | 0,2% |
| **Custo Direto Total** | **R$ 111.858,00** | **100%** |

---

### 2. Reserva de Contingência

| Método | Base | Cálculo | Valor (R$) |
|--------|------|---------|------------|
| **Desvio Padrão PERT (1σ)** | σ = 18,4h × R$115/h médio | — | R$ 2.116,00 |
| **% do Custo Direto (15%)** | 15% × R$ 111.858 | — | R$ 16.779,00 |

**Contingência recomendada:** R$ 16.779,00 (15% — margem gerencial para riscos não capturados pelo σ PERT)

| Cenário | Custo Direto | Contingência | **Custo Total** |
|---------|-------------|-------------|-----------------|
| Otimista (−1σ) | R$ 104.742 | R$ 15.711 | R$ 120.453 |
| **Provável** | **R$ 111.858** | **R$ 16.779** | **R$ 128.637** |
| Pessimista (+1σ) | R$ 118.974 | R$ 17.846 | R$ 136.820 |

---

### 3. Curva S — Custo Acumulado

| Período | RH | Infra | Licenças | Mensal | Acumulado |
|---------|-----|-------|---------|--------|-----------|
| Agosto/2026 | R$ 68.511 | R$ 2.600 | R$ 100 | R$ 71.211 | R$ 71.211 |
| Setembro/2026 | R$ 34.563 | R$ 2.200 | R$ 100 | R$ 36.863 | R$ 108.074 |
| Outubro/2026 | R$ 3.784 | R$ 0 | R$ 0 | R$ 3.784 | **R$ 111.858** |

---

### 4. Fluxo de Caixa Projetado

| Período | Entrada (R$) | Saída (R$) | Saldo (R$) |
|---------|-------------|-----------|------------|
| Agosto/2026 | R$ 85.000 | R$ 71.211 | R$ 13.789 |
| Setembro/2026 | R$ 35.000 | R$ 36.863 | −R$ 1.863 |
| Outubro/2026 | R$ 11.858 | R$ 3.784 | R$ 8.074 |
| **Total** | **R$ 131.858** | **R$ 111.858** | **R$ 20.000** |

---

### 5. Comparativo ROM × PERT

| Indicador | ROM (Upstream) | PERT (Downstream) | Variação |
|-----------|---------------|-------------------|----------|
| Horas Totais | 1.072h (±50%) | 929h (±2%) | −13,3% |
| Custo Direto | R$ 123.280 | R$ 111.858 | −9,3% |
| Custo c/ Contingência | — | R$ 128.637 | — |
| Duração | ~9 semanas | ~7,7 semanas | −14,4% |

**Análise:** O refinamento PERT reduziu a estimativa em 13% em relação ao ROM, refletindo a maior precisão do detalhamento dos 22 pacotes EAP/WBS e a remoção de buffers conservadores do ROM.

### 6. Compatibilidade com WATERFALL Doc #13

> Este artefato é consumido como `UPSTREAM_DOC` adicional pelo `PROMPT-GENERATE-ORCAMENTO.md`. As seções 1-5 fornecem dados estruturados para o Documento #13 WATERFALL.
