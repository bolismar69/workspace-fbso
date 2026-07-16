# CHANGELOG — Solar Fácil (.specs/)

Histórico de mudanças da documentação técnica.

---

## 2026-07-08 — SPEC-MINING Pipeline Inicial (full)

### Criados

#### docs/codebase/ (7 arquivos)
- `STACK.md` — Stack tecnológica completa
- `STRUCTURE.md` — Estrutura de diretórios e arquivos
- `ARCHITECTURE.md` — Camadas, padrões, fluxo de dados
- `CONVENTIONS.md` — Naming, imports, estilização, TypeScript
- `INTEGRATIONS.md` — SQLite, AsyncStorage, mock services
- `TESTING.md` — Status e recomendações
- `CONCERNS.md` — Dívida técnica, bugs, riscos

#### .specs/domain/ (1 arquivo)
- `domain.md` — Glossário de domínio (30+ termos), regras de negócio, ambiguidades resolvidas

#### .specs/api/ (1 arquivo)
- `solar-facil-api.yaml` — Schema OpenAPI 3.1 dos dados locais (SQLite + tipos TypeScript)

#### .specs/architecture/ (12 arquivos)
- `architecture.md` — Visão arquitetural consolidada
- `c4-context.md` — C4 Nível 1 (Contexto)
- `c4-containers.md` — C4 Nível 2 (Containers)
- `c4-components.md` — C4 Nível 3 (Componentes)
- `navigation-architecture.md` — Árvore de navegação e rotas
- `platform-architecture.md` — Diferenças iOS/Android
- `integrations.md` — Mapa de dependências
- `adrs/INDEX.md` — Índice cronológico
- `adrs/adr-001.md` — Expo SDK 53 + RN 0.79 New Architecture
- `adrs/adr-002.md` — NativeWind 4
- `adrs/adr-003.md` — Context API + React Query
- `adrs/adr-004.md` — SQLite offline-first

#### .specs/engineering/ (4 arquivos)
- `api-guidelines.md` — Padrões de consumo de dados
- `code-analysis.md` — Análise de fluxo de código
- `state-management.md` — Estratégia de gerenciamento de estado
- `performance-guidelines.md` — Diretrizes de otimização

#### .specs/product/ (3 arquivos)
- `product.md` — Visão do produto
- `requirements.md` — RFs + RNFs com MoSCoW
- `feature-roadmap.md` — Roadmap de features

#### .specs/design/ (4 arquivos)
- `DESIGN.md` — Design system document
- `design-tokens.md` — Tokens extraídos do código
- `components.md` — Catálogo de componentes
- `animations.md` — Catálogo de animações

#### .specs/mobile/ (8 arquivos)
- `audit-report.md` — Auditoria técnica (code-only)
- `ux-critique.md` — Avaliação heurística de UX
- `visual-inspection-report.md` — Inspeção visual (code-only)
- `accessibility-audit.md` — Auditoria de acessibilidade (code-only)
- `functional-verification.md` — Verificação funcional (code-only)
- `performance-report.md` — Análise de performance
- `offline-support-report.md` — Estratégia offline
- `store-metadata.md` — Metadados das lojas

#### .specs/governance/ (2 arquivos)
- `inventory.md` — Inventário do projeto
- `confidence-report.md` — Score de confiança (61%)

#### .specs/security/ (2 arquivos)
- `SECURITY.md` — OWASP Mobile Top 10, recomendações
- `permissions-manifest.md` — Permissões iOS/Android

#### Raiz .specs/
- `INDEX.md` — Índice central
- `CHANGELOG.md` — Este arquivo

### Atualizados
Nenhum (documentação inicial).

### Removidos
Nenhum.

### Notas
- Pipeline executado sem simulador/device — auditoria visual, acessibilidade e funcional são code-only
- Sem PROJECT_ID — pasta `business-projects/` vazia
- Score de confiança da documentação: 61%
- 5 questões em aberto no `domain/domain.md` → [ASK USER]
