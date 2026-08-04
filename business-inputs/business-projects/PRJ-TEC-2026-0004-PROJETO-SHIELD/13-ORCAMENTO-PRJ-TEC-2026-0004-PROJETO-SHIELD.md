# Orçamento: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 11-EAP-WBS, 06-Cronograma, 08-TEST-PLAN |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. Custos por Pacote EAP (Horas)

| Pacote EAP | Horas Estimadas |
|------------|----------------|
| 1.1 Infraestrutura e Ambiente Base | 64h |
| 1.2 Motor de Identidade por Cliente | 72h |
| 1.3 Portal de Acesso (Auth-BFF) | 104h |
| 1.4 Camada de Isolamento de Dados | 40h |
| 1.5 Observabilidade | 32h |
| 1.6 Testes e Homologação | 84h |
| 1.7 Entrega e Documentação | 24h |
| **Subtotal** | **420h** |

## 2. Custos por Perfil (Horas)

| Perfil | Horas | Custo/Hora Estimado | Custo Total |
|--------|-------|--------------------|-------------|
| Arquiteto / Tech Lead | 96h | R$ 180 | R$ 17.280 |
| DevOps / Platform Engineer | 160h | R$ 150 | R$ 24.000 |
| Dev Backend (Quarkus/Java) | 160h | R$ 130 | R$ 20.800 |
| Especialista IAM / Keycloak | 104h | R$ 160 | R$ 16.640 |
| DBA / DB Engineer | 64h | R$ 140 | R$ 8.960 |
| QA / Test Engineer | 120h | R$ 120 | R$ 14.400 |
| Dev Frontend | 48h | R$ 120 | R$ 5.760 |
| Project Manager (PM) | 120h | R$ 150 | R$ 18.000 |
| Product Owner (PO) | 120h | R$ 150 | R$ 18.000 |
| **Total Pessoal** | **992h** | — | **R$ 143.840** |

## 3. Custos de Infraestrutura Cloud (6 semanas)

| Recurso | Custo Mensal Estimado | 6 Semanas (1.5 mês) |
|---------|----------------------|---------------------|
| DOKS (3 nós) | R$ 2.800 | R$ 4.200 |
| PostgreSQL Managed (HA) | R$ 1.500 | R$ 2.250 |
| Redis Managed | R$ 700 | R$ 1.050 |
| Load Balancers | R$ 400 | R$ 600 |
| Cloudflare Pro | R$ 200 | R$ 300 |
| **Total Infraestrutura** | **R$ 5.600** | **R$ 8.400** |

## 4. Curva S (Horas Acumuladas)

| Semana | Horas na Semana | Horas Acumuladas | % do Total |
|--------|----------------|-----------------|------------|
| S1 (04-08/08) | 168h | 168h | 17% |
| S2 (11-15/08) | 184h | 352h | 35% |
| S3 (18-22/08) | 176h | 528h | 53% |
| S4 (25-29/08) | 168h | 696h | 70% |
| S5 (01-05/09) | 184h | 880h | 89% |
| S6 (08-12/09) | 112h | 992h | 100% |

## 5. Reserva de Contingência

| Tipo | Horas | % do Total | Valor |
|------|-------|-----------|-------|
| Curva de aprendizado (GraalVM Native) | 24h | 2.4% | R$ 3.120 |
| Complexidade Keycloak Multi-Realm | 32h | 3.2% | R$ 5.120 |
| Complexidade políticas de isolamento | 16h | 1.6% | R$ 2.240 |
| Imprevistos gerais (10%) | 99h | 10% | R$ 14.384 |
| **Total Contingência** | **171h** | **17.2%** | **R$ 24.864** |

## 6. Resumo Consolidado

| Categoria | Horas | Custo |
|-----------|-------|-------|
| Pessoal (992h) | 992h | R$ 143.840 |
| Infraestrutura Cloud | — | R$ 8.400 |
| Reserva de Contingência (171h) | 171h | R$ 24.864 |
| **Total do Projeto** | **1.163h** | **R$ 177.104** |

---

**[STATUS: SUCESSO]** — Orçamento completo. Total: 1.163h (992h + 171h contingência), R$ 177k.
