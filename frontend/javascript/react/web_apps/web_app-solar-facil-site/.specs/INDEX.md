# Solar Fácil Site — Índice da Documentação

> Documentação técnica completa gerada pela orquestração de 10 skills.
> Pipeline: `PROMPT-MINING-FRONTEND-SPECIFICATION.md` (Passos 0→11).
> Gerado em 2026-07-08. Confiança média: 🟡 75%.

---

## 📁 Mapa da Documentação

### docs/codebase/ — Codebase Knowledge (8 arquivos)

| # | Documento | Propósito |
|---|---|---|
| 1 | [STACK.md](../docs/codebase/STACK.md) | Stack tecnológico completo |
| 2 | [STRUCTURE.md](../docs/codebase/STRUCTURE.md) | Estrutura de diretórios e entry points |
| 3 | [ARCHITECTURE.md](../docs/codebase/ARCHITECTURE.md) | Camadas, padrões e fluxo de dados |
| 4 | [CONVENTIONS.md](../docs/codebase/CONVENTIONS.md) | Naming, formatação, error handling |
| 5 | [INTEGRATIONS.md](../docs/codebase/INTEGRATIONS.md) | APIs externas, dependências |
| 6 | [TESTING.md](../docs/codebase/TESTING.md) | Frameworks, estratégia, cobertura |
| 7 | [CONCERNS.md](../docs/codebase/CONCERNS.md) | Dívida técnica, bugs, riscos |
| 8 | [.codebase-scan.txt](../docs/codebase/.codebase-scan.txt) | Raw scan output |

### domain/ — Domínio (1 arquivo)

| # | Documento | Propósito |
|---|---|---|
| 1 | [domain.md](domain/domain.md) | Glossário + ubiquitous language |

### api/ — API Design (1 arquivo)

| # | Documento | Propósito |
|---|---|---|
| 1 | [solar-facil-api.yaml](api/solar-facil-api.yaml) | OpenAPI 3.1 — endpoint de contato |

### architecture/ — Arquitetura (8 arquivos)

| # | Documento | Propósito |
|---|---|---|
| 1 | [architecture.md](architecture/architecture.md) | Visão arquitetural consolidada |
| 2 | [c4-context.md](architecture/c4-context.md) | C4 Nível 1 — Contexto |
| 3 | [c4-containers.md](architecture/c4-containers.md) | C4 Nível 2 — Containers |
| 4 | [c4-components.md](architecture/c4-components.md) | C4 Nível 3 — Componentes React |
| 5 | [integrations.md](architecture/integrations.md) | Integrações e dependências |
| 6 | [adrs/INDEX.md](architecture/adrs/INDEX.md) | Índice de ADRs |
| 7 | [adrs/adr-001.md](architecture/adrs/adr-001.md) | ADR-001: Next.js 16 + React 19 |
| 8 | [adrs/adr-002.md](architecture/adrs/adr-002.md) | ADR-002: Dados estáticos |
| 9 | [adrs/adr-003.md](architecture/adrs/adr-003.md) | ADR-003: Tailwind CSS v4 |
| 10 | [adrs/adr-004.md](architecture/adrs/adr-004.md) | ADR-004: Output standalone |

### engineering/ — Engenharia (2 arquivos)

| # | Documento | Propósito |
|---|---|---|
| 1 | [api-guidelines.md](engineering/api-guidelines.md) | Padrões de consumo de API |
| 2 | [code-analysis.md](engineering/code-analysis.md) | Análise de fluxo de código |

### product/ — Produto (3 arquivos)

| # | Documento | Propósito |
|---|---|---|
| 1 | [product.md](product/product.md) | Visão do produto e personas |
| 2 | [requirements.md](product/requirements.md) | RFs + RNFs (EARS + MoSCoW) |
| 3 | [feature-roadmap.md](product/feature-roadmap.md) | Roadmap + dívidas técnicas |

### design/ — Design System (3 arquivos)

| # | Documento | Propósito |
|---|---|---|
| 1 | [DESIGN.md](design/DESIGN.md) | Design system document |
| 2 | [design-tokens.md](design/design-tokens.md) | Tokens de design |
| 3 | [components.md](design/components.md) | Catálogo de componentes |

### frontend/ — Análises de Frontend (5 arquivos)

| # | Documento | Propósito |
|---|---|---|
| 1 | [audit-report.md](frontend/audit-report.md) | Auditoria técnica de UI |
| 2 | [ux-critique.md](frontend/ux-critique.md) | Avaliação heurística (Nielsen) |
| 3 | [visual-inspection-report.md](frontend/visual-inspection-report.md) | Inspeção visual |
| 4 | [accessibility-audit.md](frontend/accessibility-audit.md) | Auditoria WCAG 2.1/2.2 |
| 5 | [functional-verification.md](frontend/functional-verification.md) | Verificação funcional |

### governance/ — Governança (2 arquivos)

| # | Documento | Propósito |
|---|---|---|
| 1 | [inventory.md](governance/inventory.md) | Inventário do projeto |
| 2 | [confidence-report.md](governance/confidence-report.md) | Score de confiança |

### security/ — Segurança (1 arquivo)

| # | Documento | Propósito |
|---|---|---|
| 1 | [SECURITY.md](security/SECURITY.md) | Headers, OWASP, auth |

### Documentos Legados (.specs/ raiz)

Estes documentos foram criados pelo spec-miner original (2026-07-05/06) e são preservados como referência:

| # | Documento | Status |
|---|---|---|
| — | ARCHITECTURE.md | ✅ Migrado para architecture/ |
| — | COMPONENTS.md | ✅ Migrado para design/components.md |
| — | DATA-MODEL.md | ✅ Migrado para domain/ |
| — | DESIGN.md | ✅ Migrado para design/DESIGN.md |
| — | DESIGN-SYSTEM.md | ✅ Migrado para design/design-tokens.md |
| — | PRODUCT.md | ✅ Migrado para product/product.md |
| — | REQUIREMENTS.md | ✅ Migrado para product/requirements.md |
| — | SECURITY.md | ✅ Migrado para security/SECURITY.md |

---

## 📊 Score de Confiança

| Área | Confiança |
|---|---|
| Stack | 🟢 100% |
| Estrutura | 🟢 100% |
| Arquitetura | 🟢 95% |
| Design System | 🟢 90% |
| Domínio | 🟢 90% |
| Requisitos | 🟢 85% |
| Integrações | 🟢 85% |
| Segurança | 🟡 70% |
| Acessibilidade | 🟡 60% |
| Performance | 🟡 60% |
| UX | 🟡 75% |
| Testes | 🔴 0% |
| **Média** | 🟡 **75%** |

---

## 🧭 Como Usar Esta Documentação

### Para Onboarding
1. [product.md](product/product.md) — entenda o produto
2. [domain/domain.md](domain/domain.md) — aprenda os termos
3. [docs/codebase/STRUCTURE.md](../docs/codebase/STRUCTURE.md) — conheça o código
4. [architecture/architecture.md](architecture/architecture.md) — veja a arquitetura

### Para Desenvolvimento
1. [engineering/code-analysis.md](engineering/code-analysis.md) — fluxo de código
2. [design/components.md](design/components.md) — componentes disponíveis
3. [docs/codebase/CONVENTIONS.md](../docs/codebase/CONVENTIONS.md) — padrões

### Para Design
1. [design/DESIGN.md](design/DESIGN.md) — design system
2. [design/design-tokens.md](design/design-tokens.md) — tokens
3. [frontend/ux-critique.md](frontend/ux-critique.md) — UX heurística

### Para Segurança
1. [security/SECURITY.md](security/SECURITY.md) — configurações

---

Última atualização: 2026-07-08
