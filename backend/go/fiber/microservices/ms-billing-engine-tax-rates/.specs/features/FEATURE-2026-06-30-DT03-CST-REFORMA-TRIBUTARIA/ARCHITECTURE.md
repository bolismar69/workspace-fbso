# Architecture Decision — DT-03: Tabela CST Reforma Tributária

**Feature:** FEATURE-2026-06-30-DT03-CST-REFORMA-TRIBUTARIA
**Versão:** 1.0
**Data:** 01 de Julho de 2026

> ⚠️ **Aviso de Leitura:** Este documento registra as decisões arquiteturais para a resolução da DT-03. Ele **complementa** o [ARCHITECTURE.md do projeto-base](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/ARCHITECTURE.md) (v2.0, Fases 0-1-2 concluídas) e a [arquitetura do microserviço](../../architecture/architecture.md).
>
> 📋 **Registro Canônico:** As ADRs desta feature estão registradas no catálogo global. Consulte:
> - [ADR-010 — Tabela `cst_reforma`](../../architecture/adrs/adr-010.md)
> - [ADR-011 — CST como Campo Calculado](../../architecture/adrs/adr-011.md)
> - [ADR-012 — CCT como Metadado](../../architecture/adrs/adr-012.md)
> - [Índice completo de ADRs](../../architecture/adrs/INDEX.md)

---

## ADR-001: Tabela `cst_reforma` como Fonte da Verdade para CST

**Status:** Proposto
**Data:** 2026-06-30

### Contexto

O código atual (`internal/reforma/reforma.go`) usa constantes hardcoded `cstPadrao = "01"` e `cstIsento = "04"` para todos os cenários de CBS/IBS. A LC 214/2025 define 18 CSTs oficiais com 164 CCTs (Classificações Tributárias) detalhadas. Precisamos de uma fonte de verdade para mapear cenários fiscais → CST oficial.

### Decisão

**Criar a tabela `cst_reforma`** como fonte canônica de CST para CBS/IBS, populada com os 164 registros oficiais da RFB.

### Alternativas Consideradas

| Alternativa | Prós | Contras |
|-------------|------|---------|
| **A) Tabela no PostgreSQL** (escolhida) | Atualizável sem deploy, cacheável via Redis, auditável, consistente com padrão existente (`iva_dual_rules`) | Requer migration SQL |
| B) Constantes em Go (switch/case) | Sem dependência de DB, rápido | 164 cases hardcoded, requer deploy para atualizar, inconsistente com padrão do projeto |
| C) Arquivo JSON embutido (`embed.FS`) | Sem dependência de DB | Não auditável, requer deploy para atualizar, inconsistente com padrão |

### Consequências

- **Positivas:**
  - Time Fiscal pode revisar/atualizar CSTs sem deploy (via Admin Fiscal — GAP-001)
  - Cache Redis automático via `CachedTaxRepository` (padrão existente)
  - Consistente com arquitetura atual: regras fiscais em PostgreSQL + cache Redis
  - Habilita DT-04 (créditos) e split payment (GAP-006)

- **Negativas:**
  - +1 consulta ao banco por cálculo de CBS e IBS (mitigada por cache Redis)
  - 164 INSERTs no `data/init.sql`

---

## ADR-002: CST como Campo Calculado (não Input)

**Status:** Proposto
**Data:** 2026-06-30

### Contexto

Há ambiguidade sobre se o CST deve ser:
- (a) Informado pelo consumidor no request (como ICMS/PIS/COFINS)
- (b) Calculado pelo motor com base nas regras fiscais

### Decisão

**O CST para CBS/IBS é calculado pelo motor**, não informado pelo consumidor. A lógica é:

1. Consultar `iva_dual_rules` → obter alíquotas e percentual de redução
2. Consultar `cst_reforma` → mapear cenário fiscal para o CST oficial
3. Retornar CST no response

### Justificativa

- CBS/IBS são tributos "por fora" — o CST reflete a **classificação legal da operação**, não uma opção do emitente
- Diferente de ICMS/PIS/COFINS (onde CST/CSOSN depende do regime do emitente), o CST de CBS/IBS é determinado pela **natureza da operação** (NCM + UF + tipo de produto)
- A tabela oficial da RFB vincula CST a características objetivas (ex: NCM de cigarro → monofásica → CST 400)

---

## ADR-003: CCT como Metadado (não Chave)

**Status:** Proposto
**Data:** 2026-06-30

### Contexto

A tabela oficial tem 2 níveis: CST (3 dígitos, 18 valores) e CCT (6 dígitos, 164 valores). O CCT é uma subclassificação detalhada que referencia artigos específicos da LC 214/2025.

### Decisão

**Usar CST (3 dígitos) como chave primária para o campo `CST` na resposta da API. Armazenar CCT como coluna de metadado** na tabela `cst_reforma` para referência e auditoria.

### Justificativa

- O CST (3 dígitos) é o código que vai no documento fiscal (NF-e, NFC-e, CT-e)
- O CCT (6 dígitos) é uma classificação expandida usada para distinguir sub-regras dentro do mesmo CST
- O consumidor da API precisa do CST para preencher o documento fiscal; o CCT é relevante apenas para auditoria interna
- A tabela mapeia CCT→CST de forma N:1 (múltiplos CCTs para o mesmo CST)

### Estrutura Resultante

```
cst_reforma
├── cst (CHAR 3)          ← retornado na API → TributoItem.CST
├── cct (CHAR 6) UNIQUE   ← metadado de auditoria
├── flags booleanas       ← lógica de negócio (monofásica, diferimento...)
└── url_legislacao        ← fundamento legal
```

---

## Diagrama de Contexto (C4 — Nível 2: Container)

```
┌──────────────────────────────────────────────────────────┐
│              ms-billing-engine-tax-rates                  │
│                                                          │
│  ┌────────────────────────────────────────────┐          │
│  │          BillingEnginePhased (SOP-013)      │          │
│  │                                            │          │
│  │  F2 (CBS): CBSCalculator                   │          │
│  │    ├─ repo.GetIvaDualRule()  ← iva_dual    │          │
│  │    └─ repo.GetCSTReforma()   ← cst_reforma │          │
│  │                                            │          │
│  │  F4 (IBS): IBSCalculator                   │          │
│  │    ├─ repo.GetIvaDualRule()  ← iva_dual    │          │
│  │    └─ repo.GetCSTReforma()   ← cst_reforma │          │
│  └────────────────────────────────────────────┘          │
│                         │                                │
│            ┌────────────┴────────────┐                   │
│            ▼                         ▼                   │
│  ┌──────────────────┐    ┌──────────────────┐            │
│  │   PostgreSQL     │    │      Redis       │            │
│  │                  │    │                  │            │
│  │ ┌──────────────┐ │    │  cache:cst:      │            │
│  │ │ iva_dual     │ │    │  {ncm}:{uf}:     │            │
│  │ │ _rules       │ │    │  {flags}         │            │
│  │ ├──────────────┤ │    │                  │            │
│  │ │ cst_reforma  │ │    │                  │            │
│  │ │ (NOVO)       │ │    │                  │            │
│  │ └──────────────┘ │    │                  │            │
│  └──────────────────┘    └──────────────────┘            │
└──────────────────────────────────────────────────────────┘
```

---

## Referências

- [ARCHITECTURE.md do projeto-base](../../business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/ARCHITECTURE.md) — Seção 6.3 (DT-03)
- [architecture.md do microserviço](../../architecture/architecture.md) — Pipeline SOP-013, injeção de dependência
- [c4-context.md do microserviço](../../architecture/c4-context.md) — Diagrama de contexto
- [integrations.md do microserviço](../../architecture/integrations.md) — Repository pattern, cache Redis
- [SPECS.md](./SPECS.md) — Especificação completa da feature
- [TASKS.md](./TASKS.md) — Tarefas de implementação
- [TEST_PLAN.md](./TEST_PLAN.md) — Plano de testes
