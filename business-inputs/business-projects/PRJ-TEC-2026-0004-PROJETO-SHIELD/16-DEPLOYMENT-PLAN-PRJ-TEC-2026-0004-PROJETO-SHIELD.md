# Plano de Deployment: PROJETO SHIELD
## [STATUS: Em revisão]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 10-SAD, 11-HLD, 12-LLD |
| **Stack** | DOKS + Kong + Shield + Istio + Argo CD + GitHub Actions |
| **Data** | 03/08/2026 | **Versão** | 2.0 — Revisão Integração | **Metodologia** | WATERFALL |

---

## 1. Deployment Strategy

**Estratégia:** RollingUpdate com GitOps (Argo CD). Kong + Shield são deployados como unidades acopladas — a configuração de rotas Kong que apontam para o Shield faz parte do deploy.

- **Dev:** Deploy automático a cada push na branch `develop`
- **Staging:** Deploy automático a cada merge na `main`
- **Prod:** Deploy via tag de release (`v*`) com approval manual no Argo CD

## 2. Environment Inventory

| Ambiente | URL | Recursos | Owner |
|----------|-----|---------|-------|
| Dev | `shield-dev.fbso.org` | DOKS 1 node + App Platform (SPA) | DevOps |
| Staging | `shield-staging.fbso.org` | DOKS 3 nodes + App Platform (SPA) | DevOps |
| Prod | `shield.fbso.org` | DOKS 3 nodes + App Platform (SPA) + HA PostgreSQL + Redis | DevOps |

## 3. Pre-deployment Checklist

| Item | Responsável | Critério |
|------|------------|---------|
| Testes do Filter Contract passam (FK-01 a FK-09) | QA | CI verde |
| Kong rotas `/api/*` → Shield `POST /internal/session/validate` configuradas | DevOps | Kong config aplicado |
| Kong rotas `/auth/callback` → Shield configuradas | DevOps | Kong config aplicado |
| SAST sem Critical/High | DevOps | Semgrep scan limpo |
| Secret scan limpo | DevOps | Gitleaks zero leaks |
| Docker image Shield publicado no GHCR | CI | Tag `v*` presente |

## 4. Deployment Steps

| Step | Ação | Responsável | Rollback |
|------|------|------------|---------|
| 1 | Criar tag de release e push | Tech Lead | Deletar tag |
| 2 | CI build + test + push GHCR | CI pipeline | — |
| 3 | CI atualiza GitOps repo (Shield image tag + Kong config) | CI pipeline | Reverter commit |
| 4 | Argo CD sync — Kong config (rotas + plugins) | Argo CD | Argo CD rollback |
| 5 | Argo CD sync — Shield BFF (RollingUpdate) | Argo CD | `kubectl rollout undo` |
| 6 | Smoke tests: (a) SPA carrega, (b) chamada API sem cookie → 302 Keycloak, (c) cookie válido → 200 com dados | QA | — |
| 7 | Verificar: `document.cookie` não expõe JWT (TC-005) | QA | — |
| 8 | Monitorar 15min: latência p95 <15ms, erro <0.1%, KEDA scaling | DevOps | Rollback se erro >1% |
| 9 | Notificar times consumidores | PM | — |

## 5. Kong Configuration (novo)

A configuração das rotas Kong é parte do deploy e versionada no repositório GitOps:

```yaml
# kong-routes.yaml (versionado no GitOps repo)
routes:
  - name: api-proxy
    paths: ["/api/*"]
    plugins:
      - name: shield-session-filter
        config:
          validate_url: "http://shield.shield-system.svc:8080/internal/session/validate"
          inject_jwt: true
  - name: auth-callback
    paths: ["/auth/callback"]
    plugins:
      - name: shield-auth-callback
        config:
          callback_url: "http://shield.shield-system.svc:8080/auth/callback"
```

## 6-9. (Seções mantidas — Database Migration, Rollback, Communication, Go-Live Runbook — conforme versão 1.0)

---

**[STATUS: SUCESSO]** — Plano de deploy atualizado com Kong configuration versionada e smoke tests do Filter Contract.
