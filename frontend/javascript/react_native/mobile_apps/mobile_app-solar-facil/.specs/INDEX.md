---
title: "Índice Central de Documentação — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
score: "61%"
---

# Índice Central de Documentação — Solar Fácil

Este índice referencia toda a documentação técnica gerada pelo pipeline **MOBILE-SPEC-MINING**.

**Última atualização:** 2026-07-08
**Confiança da documentação:** 61% (ver `governance/confidence-report.md`)

---

## 📚 docs/codebase/ — Conhecimento da Codebase

| Arquivo | Descrição |
|---|---|
| [STACK.md](../docs/codebase/STACK.md) | Stack tecnológica completa (linguagem, runtime, frameworks, dependências) |
| [STRUCTURE.md](../docs/codebase/STRUCTURE.md) | Estrutura de diretórios, entry points, rotas, componentes |
| [ARCHITECTURE.md](../docs/codebase/ARCHITECTURE.md) | Camadas, padrões arquiteturais, fluxo de dados, cross-cutting concerns |
| [CONVENTIONS.md](../docs/codebase/CONVENTIONS.md) | Naming, imports, estilização, formulários, TypeScript, git |
| [INTEGRATIONS.md](../docs/codebase/INTEGRATIONS.md) | APIs, SQLite, AsyncStorage, serviços mock, dependências externas |
| [TESTING.md](../docs/codebase/TESTING.md) | Status de testes, frameworks recomendados, estratégia de mocking |
| [CONCERNS.md](../docs/codebase/CONCERNS.md) | Dívida técnica, bugs, riscos de segurança, arquivos órfãos |

---

## 🏛️ domain/ — Domínio

| Arquivo | Descrição |
|---|---|
| [domain.md](domain/domain.md) | Glossário de domínio, ubiquitous language, regras de negócio |

---

## 🔌 api/ — Contratos de Dados

| Arquivo | Descrição |
|---|---|
| [solar-facil-api.yaml](api/solar-facil-api.yaml) | Schema SQLite + Contrato de dados (OpenAPI 3.1 para referência) |

---

## 🏗️ architecture/ — Arquitetura

| Arquivo | Descrição |
|---|---|
| [architecture.md](architecture/architecture.md) | Visão geral, stack, princípios, decisões, cross-cutting concerns |
| [c4-context.md](architecture/c4-context.md) | C4 Nível 1: Contexto do sistema (atores, sistemas externos) |
| [c4-containers.md](architecture/c4-containers.md) | C4 Nível 2: Containers (iOS/Android app, SQLite, Metro bundler) |
| [c4-components.md](architecture/c4-components.md) | C4 Nível 3: Componentes (screens, hooks, services) |
| [navigation-architecture.md](architecture/navigation-architecture.md) | Árvore de navegação, rotas, deep links, fluxos |
| [platform-architecture.md](architecture/platform-architecture.md) | Diferenças iOS/Android, APIs nativas, build/deploy |
| [integrations.md](architecture/integrations.md) | Mapa de dependências, contratos, políticas de retry/cache |
| [adrs/INDEX.md](architecture/adrs/INDEX.md) | Índice cronológico dos ADRs |
| [adrs/adr-001.md](architecture/adrs/adr-001.md) | ADR-001: Expo SDK 53 + React Native 0.79 New Architecture |
| [adrs/adr-002.md](architecture/adrs/adr-002.md) | ADR-002: NativeWind 4 para estilização |
| [adrs/adr-003.md](architecture/adrs/adr-003.md) | ADR-003: Context API + React Query para estado |
| [adrs/adr-004.md](architecture/adrs/adr-004.md) | ADR-004: SQLite offline-first como fonte de dados |

---

## ⚙️ engineering/ — Engenharia

| Arquivo | Descrição |
|---|---|
| [api-guidelines.md](engineering/api-guidelines.md) | Padrões de consumo de dados, queries/mutations, offline-first, erros |
| [code-analysis.md](engineering/code-analysis.md) | Análise de fluxo de componentes, hooks, serviços |
| [state-management.md](engineering/state-management.md) | Estratégia de estado: Context API + React Query |
| [performance-guidelines.md](engineering/performance-guidelines.md) | Otimização de renderização, virtualização, imagens, bundle |

---

## 📦 product/ — Produto

| Arquivo | Descrição |
|---|---|
| [product.md](product/product.md) | Visão do produto, proposta de valor, personas |
| [requirements.md](product/requirements.md) | Requisitos funcionais (EARS) e não-funcionais + MoSCoW |
| [feature-roadmap.md](product/feature-roadmap.md) | Features concluídas, planejadas e backlog |

---

## 🎨 design/ — Design System

| Arquivo | Descrição |
|---|---|
| [DESIGN.md](design/DESIGN.md) | Design system document (tokens, componentes, layout, ícones) |
| [design-tokens.md](design/design-tokens.md) | Tokens extraídos do código (cores, fontes, spacing, shadows) |
| [components.md](design/components.md) | Catálogo de componentes React Native com variantes |
| [animations.md](design/animations.md) | Catálogo de animações e transições (Reanimated, Moti) |

---

## 📱 mobile/ — Auditoria Mobile

| Arquivo | Descrição |
|---|---|
| [audit-report.md](mobile/audit-report.md) | Auditoria técnica: a11y, performance, responsividade, código, segurança |
| [ux-critique.md](mobile/ux-critique.md) | Avaliação heurística de UX mobile (10 heurísticas de Nielsen) |
| [visual-inspection-report.md](mobile/visual-inspection-report.md) | Inspeção visual (⚠️ code-only — simulador indisponível) |
| [accessibility-audit.md](mobile/accessibility-audit.md) | Auditoria de acessibilidade (⚠️ code-only — VoiceOver/TalkBack não testados) |
| [functional-verification.md](mobile/functional-verification.md) | Verificação funcional (⚠️ code-only — fluxos documentados, não testados) |
| [performance-report.md](mobile/performance-report.md) | Análise de performance mobile |
| [offline-support-report.md](mobile/offline-support-report.md) | Estratégia offline-first, cache, sincronização futura |
| [store-metadata.md](mobile/store-metadata.md) | Metadados App Store e Google Play |

---

## 🔒 security/ — Segurança

| Arquivo | Descrição |
|---|---|
| [SECURITY.md](security/SECURITY.md) | OWASP Mobile Top 10, autenticação, armazenamento, ofuscação |
| [permissions-manifest.md](security/permissions-manifest.md) | Permissões iOS (Info.plist) e Android (AndroidManifest) |

---

## 📊 governance/ — Governança

| Arquivo | Descrição |
|---|---|
| [inventory.md](governance/inventory.md) | Inventário do projeto: arquivos, linhas, cobertura de testes |
| [confidence-report.md](governance/confidence-report.md) | Score de confiança da documentação por área (média: 61%) |

---

## 📝 CHANGELOG

Ver [CHANGELOG.md](CHANGELOG.md)

---

## 📂 Outras Pastas

| Pasta | Descrição |
|---|---|
| `business-projects/` | Projetos corporativos vinculados (vazio — sem PROJECT_ID) |
| `features/` | Especificações por feature (vazio — features não documentadas individualmente) |
| `pull-requests/` | Documentação de pull requests |
| `questions/` | Lacunas e perguntas — [ASK USER] |
| `skill-output/` | Relatórios de execução do pipeline SPEC-MINING |

---

🤖 *Documentação gerada pela orquestração de skills: acquire-codebase-knowledge → domain-modeling → api-designer → architecture-designer → impeccable → react-native-expert → Inspeção Visual Mobile → Auditoria de Acessibilidade Mobile → Mobile App Functional Testing → documentation-writer.*
