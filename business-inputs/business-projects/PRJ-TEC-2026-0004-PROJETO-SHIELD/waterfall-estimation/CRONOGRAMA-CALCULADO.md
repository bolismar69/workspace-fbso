# CRONOGRAMA CALCULADO — DERIVADO DO PERT: PRJ-TEC-2026-0004-PROJETO-SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Estimativa Base** | WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md v1.0 (929,2h PERT) |
| **Capacidade do Time** | 2 Seniores, 1 Pleno, 1 Júnior (24h/dia combinadas) |
| **Data de Elaboração** | 04/08/2026 |
| **Versão** | 1.0 |

---

### 1. Lista de Atividades com Durações

> **Fórmula:** dias = E_PERT / (equipe × 6h/dia). Equipe por atividade varia conforme especialidade.

| ID | Atividade | Pacote EAP | E PERT (h) | Equipe | Dias |
|----|----------|-----------|-----------|--------|------|
| A01 | Provisionamento DOKS | 1.1.1 | 56,7 | 1 Sen (DevOps) | 9,5 |
| A02 | Rede e Segurança (Istio) | 1.1.2 | 46,7 | 1 Sen (DevOps) | 7,8 |
| A03 | Kong API Gateway | 1.1.3 | 43,0 | 1 Sen (DevOps) | 7,2 |
| A04 | Pipeline CI/CD | 1.1.4 | 28,3 | 1 Pl (DevOps) | 4,7 |
| A05 | Keycloak Provisionamento | 1.2.1 | 76,0 | 1 Sen (IAM) + 1 Jr | 6,3 |
| A06 | Temas e OIDC | 1.2.2 | 37,3 | 1 Sen (IAM) | 6,2 |
| A07 | Cloudflare Integração | 1.2.3 | 19,7 | 1 Pl (DevOps) | 3,3 |
| A08 | Reconhecimento Domínio | 1.3.1 | 45,7 | 1 Sen (Dev) + 1 Pl | 3,8 |
| A09 | Login e Logout | 1.3.2 | 79,3 | 1 Sen (Dev) + 1 Pl | 6,6 |
| A10 | Sessão e Cookies | 1.3.3 | 47,0 | 1 Sen (Dev) | 7,8 |
| A11 | Cache Host→Realm | 1.3.4 | 24,0 | 1 Pl (Dev) | 4,0 |
| A12 | Modelagem RLS | 1.4.1 | 57,7 | 1 Sen (DBA) + 1 Jr | 4,8 |
| A13 | Pool de Conexões | 1.4.2 | 12,3 | 1 Pl (DBA) | 2,1 |
| A14 | Métricas e Dashboards | 1.5.1 | 38,7 | 1 Sen (DevOps) | 6,5 |
| A15 | Agregação de Logs | 1.5.2 | 20,0 | 1 Pl (DevOps) | 3,3 |
| A16 | Tracing Distribuído | 1.5.3 | 20,3 | 1 Pl (DevOps) | 3,4 |
| A17 | Testes de Isolamento | 1.6.1 | 52,7 | 1 Pl (QA) + 1 Jr | 4,4 |
| A18 | Testes de Carga | 1.6.2 | 41,3 | 1 Pl (QA) | 6,9 |
| A19 | Testes de Segurança | 1.6.3 | 57,3 | 1 Sen (QA) + 1 Jr | 4,8 |
| A20 | Homologação e Aceite | 1.6.4 | 36,7 | 1 Pl (QA) + PO | 3,1 |
| A21 | Documentação API | 1.7.1 | 12,3 | 1 Pl (Dev) | 2,1 |
| A22 | Manuais Operacionais | 1.7.2 | 16,7 | 1 Sen (TechLead) | 2,8 |
| A23 | Go-Live | 1.7.3 | 25,7 | 1 Sen (DevOps) + PM | 2,1 |

**Premissas:** 6h produtivas/dia/pessoa. Equipe dedicada. Dias úteis (5/semana).

---

### 2. Sequenciamento e Dependências

| Atividade | Depende de | Tipo | Folga (dias) |
|-----------|-----------|------|-------------|
| A02 (Istio) | A01 (DOKS) | FS | 0 |
| A03 (Kong) | A02 (Istio) | FS | 0 |
| A04 (CI/CD) | A01 (DOKS) | FS | 3 |
| A05 (Keycloak) | A01 (DOKS) | FS | 0 |
| A06 (OIDC) | A05 (Keycloak) | FS | 0 |
| A07 (Cloudflare) | A05 (Keycloak) | FS | 3 |
| A08 (Domínio) | A03 (Kong), A05 (Keycloak) | FS | 0 |
| A09 (Login) | A08 (Domínio) | FS | 0 |
| A10 (Sessão) | A09 (Login) | FS | 0 |
| A11 (Cache) | A08 (Domínio) | FS | 2 |
| A12 (RLS) | A01 (DOKS) | FS | 4 |
| A13 (Pool) | A12 (RLS) | FS | 0 |
| A14 (Métricas) | A01 (DOKS), A03 (Kong) | FS | 5 |
| A15 (Logs) | A14 (Métricas) | FS | 0 |
| A16 (Tracing) | A14 (Métricas) | FS | 0 |
| A17 (Iso Test) | A09 (Login), A12 (RLS) | FS | 0 |
| A18 (Carga) | A17 (Iso Test) | FS | 0 |
| A19 (Sec Test) | A18 (Carga) | FS | 0 |
| A20 (Homolog) | A19 (Sec Test) | FS | 0 |
| A21 (Doc API) | A09 (Login) | FS | 6 |
| A22 (Manuais) | A20 (Homolog) | FS | 0 |
| A23 (Go-Live) | A20 (Homolog), A22 (Manuais) | FS | 0 |

---

### 3. Caminho Crítico

**Duração total:** **39 dias úteis (~7,7 semanas)**

```
A01(DOKS) → A05(Keycloak) → A06(OIDC) → A08(Domínio) → A09(Login) → A10(Sessão)
    ↓                                                                           ↓
A02(Istio) → A03(Kong) ─────────────────────────────────────────────┘           ↓
                                                                                ↓
A17(IsoTest) → A18(Carga) → A19(SecTest) → A20(Homolog) → A22(Manuais) → A23(GoLive)
```

### 4. Cronograma (Datas)

| Atividade | Início | Fim | Dias |
|-----------|--------|-----|------|
| A01 DOKS | 04/08/2026 | 15/08/2026 | 10 |
| A05 Keycloak | 18/08/2026 | 26/08/2026 | 7 |
| A09 Login/Logout | 27/08/2026 | 05/09/2026 | 8 |
| A17-A19 Testes | 08/09/2026 | 24/09/2026 | 13 |
| A20 Homologação | 25/09/2026 | 29/09/2026 | 5 |
| A23 Go-Live | 30/09/2026 | 01/10/2026 | 2 |

### 5. Marcos (Milestones)

| Marco | Data | Charter |
|-------|------|---------|
| M1: Kickoff | 04/08/2026 | — |
| M2: Infraestrutura OK | 15/08/2026 | D1 |
| M3: Motor ID funcional | 26/08/2026 | D2 |
| M4: Portal Acesso funcional | 05/09/2026 | D3 |
| M5: Testes concluídos | 29/09/2026 | D6 |
| M6: **Go-Live** | **01/10/2026** | D7 |

### 6. Diagrama de Gantt (Textual)

```
ATIVIDADE         | S1 | S2 | S3 | S4 | S5 | S6 | S7 | S8 |
A01 DOKS          | ██ | ██ |    |    |    |    |    |    |
A02 Istio         |    | ██ |    |    |    |    |    |    |
A05 Keycloak      |    | ██ | ██ |    |    |    |    |    |
A08-A10 Portal    |    |    | ██ | ██ |    |    |    |    |
A12-A13 RLS       |    |    | ██ |    |    |    |    |    |
A17-A19 Testes    |    |    |    | ██ | ██ | ██ |    |    |
A20 Homologação   |    |    |    |    |    |    | ██ |    |
A23 Go-Live       |    |    |    |    |    |    |    | ██ |
CRÍTICO           | ██ | ██ | ██ | ██ | ██ | ██ | ██ | ██ |
```

### 7. Compatibilidade com WATERFALL Doc #12

> Este artefato é consumido como `UPSTREAM_DOC` adicional pelo `PROMPT-GENERATE-CRONOGRAMA-GANTT.md`. As seções 1-6 fornecem dados estruturados para o Documento #12 WATERFALL.
