# Plano de Deployment: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 10-SAD, 11-HLD, 12-LLD |
| **Stack** | DOKS + Istio + Kong + Argo CD + GitHub Actions |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. Deployment Strategy

**Estratégia:** RollingUpdate com GitOps (Argo CD)

- **Dev:** Deploy automático a cada push na branch `develop`
- **Staging:** Deploy automático a cada merge na `main`
- **Prod:** Deploy via tag de release (`v*`) com approval manual no Argo CD

## 2. Environment Inventory

| Ambiente | URL | Recursos | Owner |
|----------|-----|---------|-------|
| Dev | `shield-dev.fbso.org` | DOKS 1 node (2vCPU/4GB) | DevOps |
| Staging | `shield-staging.fbso.org` | DOKS 3 nodes (4vCPU/8GB) | DevOps |
| Prod | `shield.fbso.org` | DOKS 3 nodes (4vCPU/8GB) + HA PostgreSQL + Redis | DevOps |

## 3. Pre-deployment Checklist

| Item | Responsável | Critério |
|------|------------|---------|
| Todos os tests passam (Unit + Integration + Security) | QA | CI verde |
| SAST sem Critical/High findings | DevOps | Semgrep scan limpo |
| Secret scan limpo | DevOps | Gitleaks zero leaks |
| Docker image publicado no GHCR | CI pipeline | Tag `v*` presente |
| GitOps repo atualizado com nova tag | CI pipeline | Commit automático |
| Change Request aprovado (se aplicável) | PM + PO | Aprovação no PR |

## 4. Deployment Steps

| Step | Ação | Responsável | Rollback |
|------|------|------------|---------|
| 1 | Criar tag de release `vX.Y.Z` e push | Tech Lead | Deletar tag |
| 2 | CI build + test + push GHCR | CI pipeline | — |
| 3 | CI atualiza GitOps repo com nova image tag | CI pipeline | Reverter commit |
| 4 | Argo CD detecta mudança e inicia sync | Argo CD | Argo CD rollback |
| 5 | RollingUpdate: novo pod ready → old pod terminated | Kubernetes | `kubectl rollout undo` |
| 6 | Smoke tests pós-deploy (health + /auth/login) | QA | — |
| 7 | Monitorar métricas por 15min (latência, erro, sessões) | DevOps | Rollback se erro >1% |
| 8 | Notificar times consumidores no Slack | PM | — |

## 5. Database Migration Plan

| Migração | Descrição | Script | Rollback |
|----------|-----------|--------|---------|
| M001 | Schema `shield` inicial — user_sessions + audit_events | Flyway V1__create_sessions.sql | `DROP SCHEMA shield CASCADE` |
| M002 | RLS policies por tenant | Flyway V2__rls_policies.sql | `DROP POLICY` por tabela |
| M003 | Índices para queries de sessão e auditoria | Flyway V3__add_indexes.sql | `DROP INDEX` |

## 6. Rollback Plan

| Gatilho | Ação | Tempo |
|--------|------|-------|
| Smoke test falha | `kubectl rollout undo deployment/ms-shield-identity-auth` | < 2min |
| Erro >1% por 5min | Rollback automático (Argo CD sync para tag anterior) | < 3min |
| Cross-Tenant leak detectado | Rollback imediato + war room | < 5min |
| DB migration falha | Flyway repair + rollback script | < 10min |

## 7. Communication Plan (Deploy)

| Público | Quando | Canal | Mensagem |
|---------|-------|------|---------|
| Times consumidores | 24h antes do deploy | Slack #produto | "Deploy Shield vX.Y.Z agendado para [data]. Sem impacto esperado." |
| Equipe Shield | Início do deploy | Slack #projeto-shield | "Iniciando deploy vX.Y.Z em produção" |
| Times consumidores | Deploy concluído | Slack #produto | "Shield vX.Y.Z deployed. Changelog: [link]" |
| Diretoria | Após Go-Live | E-mail | Relatório de deploy + métricas iniciais |

## 8. Validation & Smoke Tests

| Teste | Comando/Script | Critério |
|-------|--------------|---------|
| Health check | `curl -s https://shield.fbso.org/health` | `{"status":"UP"}` |
| Login flow | Cypress smoke test | 302 → Keycloak → callback OK |
| /auth/me | `curl -s -b cookies.txt https://shield.fbso.org/auth/me` | 200 + perfil |
| Cross-tenant | Script QA: token A → dados B | 200 {data: []} |
| Métricas | Prometheus: `http_requests_total{app="shield"}` | Métricas fluindo |
| KEDA | `kubectl get scaledobject -n shield-system` | Ready=True |

## 9. Go-Live Runbook

| Horário | Ação | Responsável |
|---------|------|------------|
| D-1 17:00 | Comunicar times sobre deploy | PM |
| D-day 08:00 | CI build + push tag release | Tech Lead |
| D-day 08:30 | Argo CD sync — staging validation | DevOps |
| D-day 09:00 | Argo CD sync — produção | DevOps |
| D-day 09:05 | Smoke tests | QA |
| D-day 09:20 | Monitoramento 15min — decisão GO/NO-GO | DevOps + Tech Lead |
| D-day 09:30 | Comunicar conclusão | PM |
| D-day 10:00 | Termo de aceite | PO |

---

**[STATUS: SUCESSO]** — Plano com 3 ambientes, 8 passos de deploy, 3 migrações DB, rollback <5min, 6 smoke tests, runbook Go-Live.
