# Plano de Gerenciamento de Riscos: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documento Base** | 01-PROJECT-CHARTER |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. Registro de Riscos

| ID | Descrição | Categoria | P | I | Score | Trigger | Estratégia | Owner |
|----|-----------|-----------|---|---|-------|---------|-----------|-------|
| R01 | Curva de aprendizado em compilação nativa (GraalVM) — builds falhando ou lentos no CI/CD | Técnico | Alta (4) | Alto (4) | **16** | Primeiro build nativo falha ou excede 30min | Mitigate — 24h de buffer; mentoria do Arquiteto; build em container padronizado | DevOps |
| R02 | Complexidade de configuração multi-cliente no Keycloak — Realms não isolados corretamente | Técnico | Média (3) | Alto (4) | **12** | Teste de isolamento entre Realms falha na Fase 2 | Mitigate — Especialista IAM 104h; PoC antes do desenvolvimento; template de Realm validado | IAM Specialist |
| R03 | Vazamento de dados entre clientes (Cross-Tenant Data Leak) — token da Escola A acessa dados da Escola B | Segurança | Média (3) | Crítico (5) | **15** | QA detecta acesso cruzado em teste de isolamento | Mitigate — RLS como barreira primária; QA 120h com testes cross-tenant; revisão de segurança antes do Go-Live | Tech Lead + QA |
| R04 | Indisponibilidade de serviços gerenciados da DigitalOcean — cluster ou banco fora do ar | Infraestrutura | Baixa (2) | Alto (4) | **8** | Alerta de downtime no DOKS ou PostgreSQL | Mitigate — Multi-AZ; backups automáticos; DR runbook documentado | DevOps |
| R05 | Performance das políticas de isolamento no banco — queries com RLS lentas sob carga | Técnico | Média (3) | Médio (3) | **9** | Latência de queries >100ms em teste de carga na Fase 4 | Mitigate — DBA 64h; testes de carga específicos para RLS; índices otimizados | DBA |
| R06 | Atraso na adaptação dos produtos consumidores — times de produto não integram a nova plataforma a tempo | Externo | Média (3) | Médio (3) | **9** | Time de produto reporta impedimento na integração | Mitigate — Dev Frontend alocado desde a Fase 1; contrato de API publicado na Fase 2; comunicação quinzenal | PM + PO |
| R07 | Customização visual excessiva de temas por cliente — solicitações de design fora do escopo | Negócio | Baixa (2) | Médio (3) | **6** | Cliente solicita tema personalizado durante pilot | Accept — Escopo limita a tema padrão; customizações como Change Request separado | PO |
| R08 | Não conformidade LGPD por dados em nuvem fora do Brasil — DigitalOcean sem datacenter BR | Compliance | Média (3) | Alto (4) | **12** | Jurídico aponta risco antes do Go-Live | Mitigate — Criptografia em repouso e trânsito; termo de tratamento de dados; política de exclusão de PII dos logs | Tech Lead + Jurídico |
| R09 | Exaustão de conexões de banco sob carga multi-cliente — sem pool de conexões | Técnico | Média (3) | Médio (3) | **9** | Erros de "too many connections" em teste de carga | Mitigate — PgBouncer configurado; pool sizing validado em teste de carga | DBA + DevOps |
| R10 | Escalabilidade insuficiente em pico de acesso — KEDA não escala pods rápido o bastante | Técnico | Baixa (2) | Alto (4) | **8** | Erros 5xx em teste de carga com 200+ req/s | Mitigate — Configuração de KEDA validada na Fase 5; métricas de trigger calibradas; teste de carga antecipado | DevOps + QA |

## 2. Matriz de Probabilidade × Impacto

```
IMPACTO
Crítico (5) │           │           │  R03      │           │
Alto (4)    │           │  R04,R10  │  R01,R08  │  R02      │
Médio (3)   │           │  R07      │  R05,R06  │           │
Baixo (2)   │           │           │  R09      │           │
            └───────────┴───────────┴───────────┴───────────┘
              Baixa (1)   Baixa (2)   Média (3)   Alta (4)
                             PROBABILIDADE
```

**Zona Vermelha (Score ≥ 12):** R01 (16), R03 (15), R02 (12), R08 (12) — **atenção imediata**
**Zona Amarela (Score 8-9):** R04 (8), R05 (9), R06 (9), R09 (9), R10 (8) — **monitoramento ativo**
**Zona Verde (Score < 8):** R07 (6) — **observação**

## 3. Plano de Resposta

| Risco | Estratégia | Ação | Responsável | Prazo |
|-------|-----------|------|-------------|-------|
| R01 — Curva GraalVM | Mitigate | Build nativo em container padronizado; CI/CD com cache de camadas; mentoria de 8h do Arquiteto | DevOps + Tech Lead | Semana 1 |
| R02 — Complexidade Keycloak | Mitigate | PoC de 2 Realms com isolamento validado; template documentado | IAM Specialist | Semana 2 |
| R03 — Vazamento Cross-Tenant | Mitigate | RLS em 100% das tabelas; testes cross-tenant automatizados; revisão de segurança | Tech Lead + QA + DBA | Semana 4 |
| R04 — Indisponibilidade DO | Mitigate | Multi-AZ; backups diários; DR runbook; alertas configurados | DevOps | Semana 1 |
| R08 — LGPD/Nuvem BR | Mitigate | Termo de tratamento assinado; criptografia em repouso; política de PII em logs | Tech Lead + Jurídico | Semana 2 |

## 4. Plano de Contingência

| Risco | Trigger | Ação de Contingência |
|-------|---------|---------------------|
| R01 — Build nativo falha | Pipeline de CI quebra por >24h | Fallback para build JVM (não-nativo) até correção; +72h no cronograma |
| R03 — Cross-Tenant detectado | QA encontra acesso cruzado | War Room imediata; congelamento de deploy; revisão completa de políticas de isolamento |
| R04 — DOKS indisponível | Downtime >2h | Acionar suporte DO; comunicar times consumidores; avaliar ambiente de contingência |
| R08 — Bloqueio LGPD | Jurídico veta Go-Live | Plano de migração para nuvem com datacenter BR; renegociação de prazo com Diretoria |

## 5. Riscos Residuais (após mitigação)

| Risco Original | Risco Residual | Nível Após Mitigação |
|---------------|---------------|---------------------|
| R01 — GraalVM | Possibilidade de build nativo ainda apresentar problemas pontuais em edge cases | Baixo (4) |
| R03 — Cross-Tenant | RLS cobre queries diretas; ataques sofisticados por inferência temporal permanecem como risco residual | Baixo (6) |
| R08 — LGPD | Risco regulatório residual por ausência de datacenter no Brasil — mitigado por medidas contratuais e técnicas | Médio (6) |

## 6. Monitoramento

| Indicador | Threshold | Frequência de Revisão | Responsável |
|-----------|----------|----------------------|-------------|
| Tempo de build GraalVM no CI | < 30min | A cada build | DevOps |
| Isolamento entre Realms Keycloak | 0 falhas | Semanal (Fases 1-3) | IAM Specialist |
| Testes cross-tenant (QA) | 100% bloqueados | Diário (Fase 5) | QA |
| Disponibilidade DOKS/PostgreSQL | 99.9% | Contínuo (alertas) | DevOps |
| Latência queries RLS | p95 < 50ms | Semanal (Fases 3-5) | DBA |
| Burndown de horas da equipe | ≤ desvio de 10% | Semanal | PM |

---

**[STATUS: SUCESSO]** — Plano com 10 riscos registrados (4 vermelhos, 5 amarelos, 1 verde), 5 planos de resposta, 4 contingências, 3 riscos residuais, 6 indicadores de monitoramento.
