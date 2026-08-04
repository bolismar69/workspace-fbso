# Cronograma e Diagrama de Gantt: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 05-EAP-WBS |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. Cronograma por Semana

| Semana | Período | Fase | Atividades Principais | Equipe |
|--------|---------|------|----------------------|--------|
| **S1** | 04-08/08 | Infraestrutura + Motor ID | Provisionamento DOKS (1.1.1), Rede/Istio (1.1.2), Kong (1.1.3), Pipeline (1.1.4), Início Keycloak (1.2.1) | DevOps (40h), IAM (16h) |
| **S2** | 11-15/08 | Motor ID + Início Portal | Keycloak Realms (1.2.1), Temas OIDC (1.2.2), Cloudflare (1.2.3), Início Dev BFF (1.3.1) | IAM (32h), DevOps (16h), Dev Backend (16h) |
| **S3** | 18-22/08 | Portal de Acesso + Isolamento | Login/Logout (1.3.2), Sessão (1.3.3), Cache (1.3.4), Isolamento DB (1.4.1), Pool (1.4.2) | Dev Backend (40h), DBA (20h) |
| **S4** | 25-29/08 | Portal + Observabilidade | Finalizar BFF, Métricas (1.5.1), Logs (1.5.2), Tracing (1.5.3) | Dev Backend (24h), DevOps (24h), DBA (8h) |
| **S5** | 01-05/09 | Testes + Documentação | Isolamento (1.6.1), Carga (1.6.2), Segurança (1.6.3), API docs (1.7.1) | QA (40h), IAM (8h), Dev Backend (8h) |
| **S6** | 08-12/09 | Homologação + Go-Live | Homologação (1.6.4), Manuais (1.7.2), Go-Live (1.7.3) | QA (16h), DevOps (8h), PM (8h), PO (8h) |

## 2. Dependências entre Atividades

| Atividade | Depende de | Tipo |
|-----------|-----------|------|
| 1.3.1 Reconhecimento de Domínio | 1.1.3 Kong, 1.1.4 Pipeline | Finish-to-Start |
| 1.3.2 Login/Logout | 1.2.2 OIDC, 1.3.1 Domínio | Finish-to-Start |
| 1.3.3 Sessão/Cookies | 1.3.2 Login/Logout | Finish-to-Start |
| 1.4.1 Isolamento DB | 1.2.1 Keycloak Realms | Finish-to-Start |
| 1.5.1 Métricas | 1.3 Portal (parcial) | Start-to-Start |
| 1.6.1 Testes Isolamento | 1.3 Portal + 1.4 Isolamento | Finish-to-Start |
| 1.6.2 Testes Carga | 1.5 Observabilidade | Finish-to-Start |
| 1.7.3 Go-Live | 1.6.4 Homologação | Finish-to-Start |

## 3. Caminho Crítico

```
1.1 Infra (S1) → 1.2 Keycloak (S1-S2) → 1.3 Portal BFF (S2-S4) → 1.6 Testes (S5) → 1.7 Go-Live (S6)
```

**Duração total pelo caminho crítico: 6 semanas (30 dias úteis)**

## 4. Marcos (Milestones)

| Marco | Data | Vinculado ao Charter |
|-------|------|---------------------|
| M1: Kickoff | 04/08 | M1 — Kickoff |
| M2: Infra + Keycloak Prontos | 15/08 | M2 — Base Pronta |
| M3: Portal BFF + Isolamento Prontos | 29/08 | M3 — Plataforma Desenvolvida |
| M4: Observabilidade Ativa | 05/09 | M4 — Monitoramento Ativo |
| M5: Testes Concluídos | 12/09 | M5 — Homologação Concluída |
| M6: Go-Live | 15/09 | M6 — Go-Live |

## 5. Diagrama de Gantt

```
              S1       S2       S3       S4       S5       S6
              04-08    11-15    18-22    25-29    01-05    08-12
1.1 Infra     ████████
1.2 Keycloak  ████████████████
1.3 Portal            ████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
1.4 Isolamento                ████████████
1.5 Observab.                         ████████
1.6 Testes                                    ████████████████
1.7 Docs/Go-Live                                      ████▒▒▒▒████

M1 Kickoff    █
M2 Base       ░░█
M3 Plataforma ░░░░░░░░█
M4 Monitor    ░░░░░░░░░░░░█
M5 Homolog    ░░░░░░░░░░░░░░░░░░█
M6 Go-Live    ░░░░░░░░░░░░░░░░░░░░░░█
```

---

**[STATUS: SUCESSO]** — Cronograma 6 semanas, caminho crítico identificado, 6 marcos alinhados com o Charter.
