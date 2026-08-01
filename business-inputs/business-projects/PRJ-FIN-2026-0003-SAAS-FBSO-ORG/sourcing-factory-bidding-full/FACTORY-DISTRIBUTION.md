# FACTORY-DISTRIBUTION — Registro de Distribuição do RFQ (Full Mode)

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Modo:** Full — Projeto Completo
- **Data:** 31/07/2026
- **Status:** RFQ ENVIADO

---

## 1. Fábricas Participantes

| Fábrica | Data Envio | Canal | E-mail | Telefone | Prazo | Status |
|:---|:---|:---|:---|:---|:---|:---|
| Stefanini | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |
| Capgemini | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |
| CI&T | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |
| TOTVS | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |
| Deloitte | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |
| Infosys | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |
| TCS | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |
| Overlabs | 31/07/2026 | — | — | — | 07/08/2026 | 📤 Enviado |

---

## 2. Pacote Enviado

| Item | Arquivo |
|:---|:---|
| RFQ Package | `RFQ-PACKAGE.md` |
| Estimation Schema | `ESTIMATION-SCHEMA.csv` |
| Baseline PERT (referência) | `../downstream-architecture-refinement/BOTTOM-UP-PERT-ESTIMATE.md` |
| Scope Snapshot (62 US) | `../downstream-architecture-refinement/SCOPE-SNAPSHOT.md` |

---

## 3. Prazos

| Marco | Data |
|:---|:---|
| Data de Envio | 31/07/2026 |
| Prazo para Dúvidas | 02/08/2026 (2 dias úteis) |
| **Prazo para Entrega das Estimativas** | **07/08/2026 (5 dias úteis)** |
| Validação DTA | 08/08/2026 |
| Comparação e Ranking | 08/08/2026 |
| Notificação às Fábricas | 08/08/2026 |

---

## 4. Instruções de Preenchimento

1. Baixar `ESTIMATION-SCHEMA.csv`
2. Substituir `{FABRICA}` pelo nome da fábrica
3. Preencher todas as colunas numéricas (horas_dev, horas_qa, horas_arch, horas_devops, horas_gestao)
4. `total_horas` = soma das horas (dev+qa+arch+devops+gestao)
5. `prazo_entrega_meses` = prazo total em meses
6. `time_estimado_pessoas` = tamanho do time necessário
7. `comentarios` = racional detalhado da estimativa (metodologia, premissas, justificativas)
8. Enviar CSV preenchido por email para o PMO FBSO.ORG

---

## 5. Regras de Validação (DTA)

| Regra | Critério |
|:---|:---|
| QA ≥ 25% | QA global ≥ 25% do total de horas |
| QA por épico ≥ 20% | Cada épico com QA ≥ 20% |
| Arch ≥ 5% | Arquitetura ≥ 5% do total |
| Consistência Prazo×Horas | `total_horas / (time × 160)` vs `prazo_meses` — divergência >50% → rejeitada |
| Formato | Colunas obrigatórias preenchidas |
| Outliers | Total de horas dentro de ±50% da mediana cross-fábrica |

---

🤖 *Factory Distribution — Fase 3 do Sourcing & Factory Bidding (Full Mode)*
