# FACTORY-DISTRIBUTION.md — Registro de Distribuição
## Sourcing & Factory Bidding — Fase 3 — Bloco B

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FACTORY-DISTRIBUTION-v1.0 |
| **Versão** | 1.0 — Discovery-Level |
| **Data** | 03 de agosto de 2026 |
| **Modo** | `discovery` |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 03/08/2026 |

---

## 1. Fábricas Participantes

| # | Fábrica | Contato | Data de Envio | Prazo de Resposta | Status |
|---|---------|---------|:------------:|:-----------------:|--------|
| 1 | `CAPGEMINI` | `estimates@capgemini.com` | `03/08/2026` | `{data}` | ⏳ Em validação |
| 2 | `INFOSYS` | `estimates@infosys.com` | `03/08/2026` | `{data}` | ⏳ Em validação |
| 3 | `STEFANINI` | `estimates@stefanini.com` | `03/08/2026` | `{data}` | ⏳ Em validação |
| 4 | `TOTVS` | `estimates@totvs.com` | `03/08/2026` | `{data}` | ⏳ Em validação |

> **Nota operacional:** Recomenda-se um mínimo de 3 fábricas para garantir competição e comparabilidade estatística (mediana cross-fábrica para detecção de outliers).

---

## 2. Pacote Enviado

| Item | Arquivo | Descrição |
|------|---------|-----------|
| 1 | `RFQ-PACKAGE.md` | Pacote completo: visão do projeto, escopo (4 épicos, 18 funcionalidades), stack, instruções |
| 2 | `ESTIMATION-SCHEMA.csv` | Template CSV (16 colunas) para preenchimento da estimativa |
| 3 | `../upstream-architecture-discovery/` | Artefatos técnicos completos (11 documentos F1-F11) |

---

## 3. Instruções de Preenchimento

Cada fábrica recebeu o pacote RFQ com as seguintes orientações:

1. Preencher `ESTIMATION-SCHEMA.csv` com uma linha por épico (EP-0001 a EP-0004)
2. Todas as 5 colunas de horas (`horas_dev`, `horas_qa`, `horas_arch`, `horas_devops`, `horas_gestao`) devem ser preenchidas
3. `total_horas` deve ser a soma exata das 5 colunas
4. `prazo_entrega_meses` deve ser consistente com `total_horas / (time_estimado_pessoas × 160h)` (divergência > 50% → REJEITADA)
5. `time_estimado_pessoas` e `valor_estimado` são obrigatórios
6. QA global deve ser ≥ 25% do total de horas; QA < 10% → REJEITADA
7. Arquitetura global deve ser ≥ 5% do total; Arch < 2% → REJEITADA
8. Nomear arquivo como `ESTIMATION-SCHEMA-{NOME-DA-FABRICA}.csv`
9. Enviar para o canal indicado pela FBSO.ORG até a data de prazo

---

## 4. Registro de Comunicações

| Data | Fábrica | Canal | Resumo |
|------|---------|-------|--------|
| | | | |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 03/08/2026 | Criação inicial: registro de fábricas participantes e distribuição do RFQ | PMO |

---

🤖 *Sourcing & Factory Bidding — Fase 3. Registro de distribuição do pacote RFQ.*
