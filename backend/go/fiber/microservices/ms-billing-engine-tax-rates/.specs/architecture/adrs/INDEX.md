# Architecture Decision Records — ms-billing-engine-tax-rates

Registro canônico de todas as decisões arquiteturais do microsserviço de cálculo de tributos sobre faturamento.

**Formato:** [ADR — Architecture Decision Record](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions)
**Convenção de numeração:** ADR-NNN sequencial, imutável. Status: `Proposto` → `Aceito` → `Substituído` ou `Obsoleto`.

---

## Índice Cronológico

| ADR | Data | Status | Título | Documento Fonte |
|-----|------|--------|--------|-----------------|
| [ADR-001](./adr-001.md) | 2026-06-20 | ✅ Aceito | Motor de Cálculo Unificado — `BillingEnginePhased` SOP-013 7-fases | `architecture.md`, `engineering/code-analysis.md` |
| [ADR-002](./adr-002.md) | 2026-06-21 | ✅ Aceito | Phase Resolution System como Mecanismo de Transição Temporal (2026–2033) | `PRJ-FIN-2026-0001/ARCHITECTURE.md` §3.2 |
| [ADR-003](./adr-003.md) | 2026-06-21 | ✅ Aceito | Cache Redis + Circuit Breaker como Camada de Resiliência Regulatória | `PRJ-FIN-2026-0001/ARCHITECTURE.md` §3.3 |
| [ADR-004](./adr-004.md) | 2026-06-20 | ✅ Aceito | Strategy Pattern para Cálculo de PIS/COFINS por CST | `engineering/code-analysis.md` §3 |
| [ADR-005](./adr-005.md) | 2026-06-21 | ✅ Aceito | DDD com Interface `TaxCalculator` e Injeção de Dependência Manual | `architecture.md` §1, `domain/domain.go` |
| [ADR-006](./adr-006.md) | 2026-06-21 | ✅ Aceito | W3C Trace Context — Rastreabilidade Distribuída Ponta a Ponta | `middleware/requestid.go`, `architecture.md` §2 |
| [ADR-007](./adr-007.md) | 2026-06-24 | ✅ Aceito | API Versioning — Prefixo `/v1/` com Rotas Legacy Deprecadas | `PRJ-FIN-2026-0001/ARCHITECTURE.md` §3.4, `main.go` |
| [ADR-008](./adr-008.md) | 2026-06-24 | ✅ Aceito | Deploy Docker + Kubernetes — Artefatos de Produção (GAP-010) | `Dockerfile`, `deploy/k8s/`, `docker-compose.yaml` |
| [ADR-009](./adr-009.md) | 2026-06-24 | ✅ Aceito | Rate Limiting — Proteção contra Abuso/DoS (Fase 0) | `main.go`, `middleware/ratelimit.go` |
| [ADR-010](./adr-010.md) | 2026-07-01 | ✅ Aceito | Tabela `cst_reforma` como Fonte Oficial de CST para CBS/IBS | `FEATURE-DT03/ARCHITECTURE.md` §ADR-001 |
| [ADR-011](./adr-011.md) | 2026-07-01 | ✅ Aceito | CST como Campo Calculado pelo Motor (não Input do Consumidor) | `FEATURE-DT03/ARCHITECTURE.md` §ADR-002 |
| [ADR-012](./adr-012.md) | 2026-07-01 | ✅ Aceito | CCT como Metadado de Auditoria (não Chave Primária) | `FEATURE-DT03/ARCHITECTURE.md` §ADR-003 |

---

## Status Legend

| Símbolo | Significado |
|---------|-------------|
| 🔨 Proposto | Em discussão, ainda não implementado |
| ✅ Aceito | Implementado e em produção (ou branch ativo) |
| ⚠️ Substituído | Substituído por ADR mais recente (ver link) |
| ❌ Obsoleto | Não mais aplicável |

---

## Como Adicionar um Novo ADR

1. Criar arquivo `adr-NNN.md` nesta pasta (próximo número sequencial)
2. Seguir o template: Título, Status, Data, Contexto, Decisão, Alternativas, Consequências
3. Atualizar este índice com a nova entrada
4. Referenciar o ADR no documento de arquitetura relevante

---

> **Última atualização:** 2026-07-01
> **Total de ADRs:** 12 (12 ✅ aceitos)
