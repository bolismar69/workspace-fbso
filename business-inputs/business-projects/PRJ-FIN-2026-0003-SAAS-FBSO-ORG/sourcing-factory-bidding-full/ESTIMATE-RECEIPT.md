# ESTIMATE-RECEIPT — Registro de Recebimento de Estimativas (Full Mode)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Prazo Final:** 07/08/2026

---

## 1. Procedimento de Recebimento

1. Fábrica envia CSV preenchido por email
2. Time operacional salva na pasta `estimates/` com o padrão: `ESTIMATION-SCHEMA-{FABRICA}.csv`
3. Registrar na tabela abaixo
4. Validar formato (colunas corretas, separador `;`)

---

## 2. Registro de Recebimento

| # | Fábrica | Arquivo | Data Recebimento | Total Horas | Status |
|:---|:---|:---|:---|:---:|:---|
| 1 | Stefanini | `ESTIMATION-SCHEMA-STEFANINI.csv` | 31/07/2026 | 28,030h | 🔴 Rejeitada (QA+Arch+Prazo) |
| 2 | Capgemini | `ESTIMATION-SCHEMA-CAPGEMINI.csv` | 31/07/2026 | 11,680h | 🔴 Rejeitada (QA+Prazo) |
| 3 | CI&T | `ESTIMATION-SCHEMA-CIET.csv` | 31/07/2026 | 51,680h | 🔴 Rejeitada (QA+Arch) |
| 4 | TOTVS | `ESTIMATION-SCHEMA-TOTVS.csv` | 31/07/2026 | 91,680h | 🔴 Rejeitada (QA+Arch) |
| 5 | Deloitte | `ESTIMATION-SCHEMA-DELOITTE.csv` | 31/07/2026 | 11,680h | 🔴 Rejeitada (QA+Prazo) |
| 6 | Infosys | `ESTIMATION-SCHEMA-INFOSYS.csv` | 31/07/2026 | 11,680h | 🔴 Rejeitada (QA+Prazo) |
| 7 | TCS | `ESTIMATION-SCHEMA-TCS.csv` | 31/07/2026 | 131,680h | 🔴 Rejeitada (QA+Arch+Outlier) |
| 8 | Overlabs | `ESTIMATION-SCHEMA-OVERLABS.csv` | 31/07/2026 | 27,680h | 🔴 Rejeitada (QA+Arch) |

---

## 3. Status do Recebimento

| Métrica | Valor |
|:---|---:|
| Total enviadas | 8 |
| Recebidas | 0 |
| Pendentes | 8 |
| Prazo final | 07/08/2026 |

---

🤖 *Estimate Receipt — Fase 4 do Sourcing & Factory Bidding (Full Mode)*
