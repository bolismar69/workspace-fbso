# Contexto:
  - Este prompt orquestra **11 skills especializadas** para minerar, criar e atualizar a documentação técnica completa de **aplicações mobile** (React Native, Expo, Flutter, iOS nativo, Android nativo, Kotlin Multiplatform) no diretório `.specs/`.
  - As skills são organizadas em **3 camadas**: descoberta (codebase + domínio), design & arquitetura (API + arquitetura + mobile design), e qualidade (documentação + verificação).
  - A estrutura de pastas alvo (descrita abaixo) expande a referência canônica original com artefatos específicos de mobile: design system nativo, auditoria de UI/UX mobile, acessibilidade em plataformas móveis (TalkBack/VoiceOver), inspeção visual via simulador/device, e verificação funcional com ferramentas de teste mobile.
  - A documentação existente em `.specs/` deve ser preservada e incrementada, nunca sobrescrita sem auditoria de delta.
  - Use o README.md da raiz da solução para instruções de build, execução e testes.

---

# 🎯 Skills Orquestradas

## Camada 1 — Descoberta e Fundamentos

| Ordem | Skill | Responsabilidade | Pasta(s) Principal(is) |
|---|---|---|---|
| 0ª | `acquire-codebase-knowledge` | Scanner Python + 7 docs estruturados (STACK, STRUCTURE, ARCHITECTURE, CONVENTIONS, INTEGRATIONS, TESTING, CONCERNS) | `docs/codebase/` |
| 1ª | `domain-modeling` | Glossário de domínio, ubiquitous language, termos de negócio | `domain/` |

## Camada 2 — Design e Arquitetura

| Ordem | Skill | Responsabilidade | Pasta(s) Principal(is) |
|---|---|---|---|
| 2ª | `api-designer` | Contrato OpenAPI 3.1 dos endpoints consumidos pelo mobile, modelagem de recursos, versionamento, segurança de API | `api/` |
| 3ª | `architecture-designer` | C4 (Context→Containers→Components), ADRs, integrações, dicionário de dependências, arquitetura de navegação, padrões mobile (MVVM, Clean Architecture, etc.) | `architecture/` |
| 4ª | `impeccable` | Auditoria técnica de UI mobile (`audit`), avaliação heurística de UX adaptada para touch/gesture (`critique`), geração de DESIGN.md (`document`), extração de design tokens nativos | `mobile/`, `design/` |
| 5ª | `react-native-expert` | Revisão de padrões React Native/Expo, performance de renderização (FlatList, FlashList), uso de Reanimated, navegação, boas práticas de componentes nativos | `mobile/`, `engineering/` |
| 6ª | Inspeção Visual Mobile | Inspeção visual com screenshots de simulador/device — detecção de problemas de layout em diferentes telas/resoluções, safe areas, notch, gestos, responsividade cross-device | `mobile/` |
| 7ª | Auditoria de Acessibilidade Mobile | Auditoria de acessibilidade para plataformas móveis: TalkBack (Android), VoiceOver (iOS), touch targets (≥48dp), contraste em dispositivos móveis, hierarquia de acessibilidade, labels semânticas | `mobile/` |

## Camada 3 — Qualidade e Consistência

| Ordem | Skill | Responsabilidade | Pasta(s) Principal(is) |
|---|---|---|---|
| 8ª | Mobile App Functional Testing | Verificação funcional com ferramentas de teste mobile (Detox, Maestro, Jest + React Native Testing Library) — fluxos de usuário, formulários, navegação entre telas, deep links, estados offline | `mobile/` |
| 9ª | `documentation-writer` | Diátaxis (Tutorials/How-to/Reference/Explanation), qualidade textual, consistência cross-documento, INDEX.md, CHANGELOG.md | **Todas as pastas** (revisão final) |

---

# ⚙️ Parâmetros de Entrada (preencher antes de executar)

> **Instrução:** No momento de invocar este prompt, o humano deve informar os valores abaixo. Se algum arquivo não existir, marcar como `N/D` (não disponível) e criar no passo apropriado.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_ROOT}` | Caminho absoluto da raiz da solução | `/home/user/work/mobile_app-solar-facil` |
| `{SPECS_DIR}` | Caminho relativo da pasta de especificações | `.specs` |
| `{SOLUTION_NAME}` | Nome da solução/aplicação | `mobile_app-solar-facil` |
| `{SOLUTION_DESCRIPTION}` | Descrição curta do propósito da solução | `App mobile React Native (Expo) para simulação de energia solar e cálculo de economia (Solar as a Service)` |
| `{APP_TYPE}` | Tipo de aplicação mobile | `react-native`, `expo`, `flutter`, `ios-native`, `android-native`, `kotlin-multiplatform`, `capacitor`, `pwa` |
| `{LANGUAGE}` | Linguagem/framework principal | `typescript/react-native` |
| `{STYLING}` | Método de estilização | `nativewind`, `stylesheet`, `tamagui`, `styled-components`, `restyle`, `unistyles`, `tailwind` |
| `{NAVIGATION}` | Biblioteca de navegação | `expo-router`, `react-navigation`, `react-native-navigation`, `flutter-navigator`, `swiftui-navigation` |
| `{STATE_MANAGEMENT}` | Gerenciamento de estado | `context-api`, `zustand`, `redux-toolkit`, `jotai`, `mobx`, `riverpod`, `bloc` |
| `{BACKEND_API}` | URL base da API consumida (se aplicável) | `http://localhost:8080` |
| `{METRO_PORT}` | Porta do Metro bundler (React Native/Expo) | `8081` |
| `{PLATFORM_TARGET}` | Plataformas alvo da aplicação | `ios`, `android`, `cross-platform`, `ios+android` |
| `{TESTING_FRAMEWORK}` | Framework de teste principal | `jest+rntl`, `detox`, `maestro`, `appium`, `xctest`, `espresso`, `flutter-test` |
| `{MIN_IOS_VERSION}` | Versão mínima do iOS suportada | `15.0` |
| `{MIN_ANDROID_SDK}` | SDK mínimo do Android suportado | `24` (Android 7.0) |
| `{PROJECT_ID}` | Identificador do projeto corporativo (se aplicável) | `PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO` |
| `{SCOPE}` | Escopo da documentação a ser gerada | `full`, `discovery-only`, `design-only`, `mobile-only`, `qa-only`, `delta` |
| `{TARGET_FILES}` | Lista específica de arquivos a criar/atualizar (se `{SCOPE}=delta`) | `mobile/audit-report.md, design/DESIGN.md` |

---

# 📁 Estrutura de Pastas Alvo (Referência Canônica — Mobile)

> Esta estrutura expande a referência canônica original com pastas específicas de aplicações mobile. Pastas/arquivos ausentes são **oportunidade de criação**; existentes são **base para atualização incremental**.

```
{SPECS_DIR}/
├── INDEX.md                                        ← Índice central de toda a documentação
├── CHANGELOG.md                                    ← Histórico de mudanças da documentação
│
├── domain/
│   └── domain.md                                   ← Glossário de domínio + ubiquitous language
│
├── api/
│   └── {api-name}.yaml                             ← Contrato OpenAPI 3.1 (endpoints consumidos pelo app)
│
├── architecture/
│   ├── architecture.md                             ← Visão arquitetural geral da solução mobile
│   ├── c4-context.md                               ← Diagrama C4 — Nível 1 (Contexto: usuários mobile, APIs, serviços externos)
│   ├── c4-containers.md                            ← Diagrama C4 — Nível 2 (Containers: app bundle, Metro bundler, backend, push notifications, analytics)
│   ├── c4-components.md                            ← Diagrama C4 — Nível 3 (Componentes: Screens → Components → Hooks → Services)
│   ├── navigation-architecture.md                  ← Arquitetura de navegação: rotas, deep links, tabs, stacks, modals
│   ├── platform-architecture.md                    ← Arquitetura específica por plataforma: diferenças iOS/Android, módulos nativos
│   ├── integrations.md                             ← Integrações externas, APIs, serviços de push, analytics, crash reporting
│   └── adrs/                                       ← ADRs específicos da solução mobile
│       ├── INDEX.md                                ← Índice cronológico dos ADRs
│       ├── adr-001.md                              ← Ex: Escolha Expo SDK 53 + React Native 0.79
│       ├── adr-002.md                              ← Ex: Escolha NativeWind 4 para estilização
│       └── ...
│
├── engineering/
│   ├── api-guidelines.md                           ← Padrões de consumo de API, erros, offline-first, cache strategy
│   ├── code-analysis.md                            ← Análise técnica de fluxo de componentes/hooks/services
│   ├── state-management.md                         ← Estratégia de gerenciamento de estado, estrutura de stores/contexts
│   └── performance-guidelines.md                   ← Otimização de renderização (FlatList, memo, Reanimated), tamanho de bundle, imagens
│
├── product/
│   ├── product.md                                  ← Descrição do produto e visão
│   ├── requirements.md                             ← Especificação de requisitos (RF/RNF + MoSCoW)
│   └── feature-roadmap.md                          ← Roadmap de features + dívidas técnicas
│
├── design/                                         ← Design System Mobile (populado pelo `impeccable`)
│   ├── DESIGN.md                                   ← Design system document (gerado por `impeccable document`)
│   ├── design-tokens.md                            ← Tokens de design (cores, tipografia, spacing, shadows, animations)
│   ├── components.md                               ← Catálogo de componentes React Native/Flutter com variantes
│   └── animations.md                               ← Catálogo de animações e transições (Reanimated, Moti, Lottie, etc.)
│
├── mobile/                                         ← Análises Específicas de Mobile
│   ├── audit-report.md                             ← Relatório de auditoria técnica (`impeccable audit` mobile)
│   ├── ux-critique.md                              ← Avaliação heurística de UX mobile com foco em touch/gesture (`impeccable critique`)
│   ├── visual-inspection-report.md                 ← Inspeção visual com screenshots multi-dispositivo/resolução
│   ├── accessibility-audit.md                      ← Auditoria de acessibilidade mobile (TalkBack, VoiceOver, touch targets)
│   ├── functional-verification.md                  ← Verificação funcional com ferramentas de teste mobile
│   ├── performance-report.md                       ← Performance mobile (render, bundle size, memória, consumo de bateria) — opcional
│   ├── offline-support-report.md                   ← Comportamento offline, cache strategy, sync de dados — opcional
│   └── store-metadata.md                           ← Metadados das lojas (App Store, Google Play): descrições, keywords, screenshots
│
├── governance/
│   ├── inventory.md                                ← Inventário do projeto + cobertura de testes
│   └── confidence-report.md                        ← Relatório de confiança da documentação (%)
│
├── security/
│   ├── SECURITY.md                                 ← Definições de segurança (SSL pinning, autenticação, armazenamento seguro, OWASP Mobile Top 10)
│   └── permissions-manifest.md                     ← Permissões solicitadas por plataforma (iOS plist, AndroidManifest)
│
├── business-projects/
│   ├── README.md                                   ← Índice de projetos corporativos
│   └── {PROJECT_ID}/
│       ├── PRD.md                                  ← Product Requirements Document
│       ├── ARCHITECTURE.md                         ← Arquitetura específica do projeto
│       ├── SPECS.md                                ← Especificação detalhada
│       ├── TASKS.md                                ← Checklist de tarefas com DoD
│       └── TEST_PLAN.md                            ← Plano de testes
│
├── features/
│   └── {FEATURE_NAME}/
│       ├── SPECS.md                                ← Especificação da feature
│       ├── ARCHITECTURE.md                         ← Arquitetura da feature
│       ├── TASKS.md                                ← Tarefas da feature
│       └── TEST_PLAN.md                            ← Plano de testes da feature
│
├── pull-requests/
│   └── PR_{N}_{branch-name}.md                     ← Documentação de pull request
│
├── questions/
│   └── questions_{NNN}.md                          ← Lacunas, perguntas e histórico de respostas
│
└── skill-output/
    └── {AAAA-MM-DD-HHMISS}-mobile-spec-mining.md   ← Relatório de execução do pipeline
```

### Diretório Auxiliar: `docs/codebase/` (gerado pelo `acquire-codebase-knowledge`)

```
docs/codebase/
├── STACK.md                                        ← Linguagem, runtime, frameworks, dependências npm/cocoapods/gradle
├── STRUCTURE.md                                    ← Layout de diretórios, entry points, arquivos de configuração (app.json, eas.json, etc.)
├── ARCHITECTURE.md                                 ← Camadas (screens→components→hooks→services), padrões, fluxo de dados
├── CONVENTIONS.md                                  ← Naming, formatação, error handling, imports, padrões React Native/Expo
├── INTEGRATIONS.md                                 ← APIs externas, push notifications (FCM/APNs), analytics, crash reporting (Sentry/Crashlytics)
├── TESTING.md                                      ← Frameworks, organização de testes, mocking strategy, device/simulator testing
└── CONCERNS.md                                     ← Dívida técnica, bugs conhecidos, riscos, gargalos de performance, dependências desatualizadas
```

---

# 🧭 Mapeamento Skill → Artefato

> Tabela de responsabilidade: qual skill é **dona** (primary owner) de cada artefato. O `documentation-writer` atua como revisor de **todos** os artefatos ao final.

## Artefatos Genéricos

| Artefato | Skill Dona | Ação Típica |
|---|---|---|
| `docs/codebase/STACK.md` | `acquire-codebase-knowledge` | Rodar `scan.py`, preencher template STACK |
| `docs/codebase/STRUCTURE.md` | `acquire-codebase-knowledge` | Mapear diretórios, entry points, arquivos de config |
| `docs/codebase/ARCHITECTURE.md` | `acquire-codebase-knowledge` | Documentar camadas e fluxo de dados |
| `docs/codebase/CONVENTIONS.md` | `acquire-codebase-knowledge` | Extrair padrões de código mobile |
| `docs/codebase/INTEGRATIONS.md` | `acquire-codebase-knowledge` | Mapear APIs, push notifications, analytics, crash reporting |
| `docs/codebase/TESTING.md` | `acquire-codebase-knowledge` | Documentar frameworks e estratégias de teste mobile |
| `docs/codebase/CONCERNS.md` | `acquire-codebase-knowledge` | Listar dívidas, bugs, riscos de segurança mobile |
| `domain/domain.md` | `domain-modeling` | Criar/atualizar glossário, resolver ambiguidades |
| `api/{api-name}.yaml` | `api-designer` | Gerar/atualizar contrato OpenAPI 3.1 |
| `architecture/architecture.md` | `architecture-designer` | Escrever visão geral, padrões mobile, trade-offs |
| `architecture/c4-*.md` | `architecture-designer` | Gerar diagramas C4 nos níveis aplicáveis (Mermaid) |
| `architecture/navigation-architecture.md` | `architecture-designer` | Documentar árvore de navegação, deep links, rotas |
| `architecture/platform-architecture.md` | `architecture-designer` | Diferenças iOS/Android, módulos nativos, bridging |
| `architecture/integrations.md` | `architecture-designer` | Mapear dependências externas e libs mobile |
| `architecture/adrs/*.md` | `architecture-designer` | Registrar decisões arquiteturais (ADR) |
| `engineering/api-guidelines.md` | `api-designer` → `documentation-writer` | Padrões de consumo de API, offline-first, retry, cache |
| `engineering/code-analysis.md` | `documentation-writer` | Análise de fluxo de componentes/hooks/services |
| `engineering/state-management.md` | `react-native-expert` → `documentation-writer` | Documentar estrutura de estado global/local |
| `engineering/performance-guidelines.md` | `react-native-expert` → `documentation-writer` | Otimizações de renderização e bundle |
| `product/product.md` | `documentation-writer` | Descrição do produto (Explanation — Diátaxis) |
| `product/requirements.md` | `documentation-writer` | Requisitos (EARS + MoSCoW) |
| `product/feature-roadmap.md` | `documentation-writer` | Roadmap + dívidas técnicas |
| `governance/inventory.md` | `documentation-writer` | Inventário e cobertura |
| `governance/confidence-report.md` | `documentation-writer` | Score de confiança da documentação |
| `security/SECURITY.md` | `documentation-writer` | SSL pinning, auth, storage seguro, OWASP Mobile Top 10 |
| `security/permissions-manifest.md` | `documentation-writer` | Permissões iOS plist + Android Manifest |
| `INDEX.md` | `documentation-writer` | Índice central (atualizado a cada ciclo) |
| `CHANGELOG.md` | `documentation-writer` | Registro de mudanças da doc |

## Artefatos Específicos de Mobile

| Artefato | Skill Dona | Ação Típica |
|---|---|---|
| `design/DESIGN.md` | `impeccable document` | Gerar documentação do design system nativo a partir do código |
| `design/design-tokens.md` | `impeccable extract` → `documentation-writer` | Extrair tokens de cor (OKLCH), tipografia nativa, spacing, shadows/elevation |
| `design/components.md` | `impeccable document` → `documentation-writer` | Catalogar componentes mobile com variantes, props e platform specifics |
| `design/animations.md` | `impeccable document` → `documentation-writer` | Catalogar animações (Reanimated, Moti, Lottie) e transições entre telas |
| `mobile/audit-report.md` | `impeccable audit` | Auditoria técnica mobile: a11y, performance, responsividade cross-device |
| `mobile/ux-critique.md` | `impeccable critique` | Avaliação heurística mobile com foco em touch, gestos, affordance |
| `mobile/visual-inspection-report.md` | Inspeção Visual Mobile | Screenshots multi-dispositivo, safe areas, notch, problemas de layout |
| `mobile/accessibility-audit.md` | Auditoria de Acessibilidade Mobile | TalkBack/VoiceOver, touch targets (≥48dp), contraste, labels semânticas |
| `mobile/functional-verification.md` | Mobile App Functional Testing | Verificação de fluxos com Detox/Maestro/Jest + RNTL |
| `mobile/performance-report.md` | `react-native-expert` | Render performance, bundle size, startup time, consumo de memória |
| `mobile/offline-support-report.md` | `react-native-expert` | Estratégia offline, cache, sincronização — opcional |
| `mobile/store-metadata.md` | `documentation-writer` | Metadados App Store / Google Play |

---

# 🔁 Protocolo de Execução (Passo a Passo)

## Passo 0 — Leitura do Cenário Atual

1. Confirmar que `{SOLUTION_ROOT}` e `{SPECS_DIR}` estão acessíveis.
2. Confirmar que o projeto mobile pode ser executado (simulador iOS/Android ou device físico disponível) — necessário para Passos 5, 6, 7, e 8.
3. Ler `{SPECS_DIR}/INDEX.md` (se existir) para entender a estrutura atual.
4. Ler `{SPECS_DIR}/CHANGELOG.md` (se existir) para conhecer o histórico.
5. Fazer scan dos arquivos existentes em cada subpasta e comparar com a estrutura canônica acima.
6. Gerar um **delta inicial**: lista de arquivos existentes (📄), ausentes (❌), e desatualizados (🔄).
7. Detectar automaticamente:
   - Framework e versão (`package.json` → `dependencies`, `pubspec.yaml`, `Podfile`, `build.gradle`)
   - Método de estilização (NativeWind config, StyleSheet, Tamagui, styled-components)
   - Porta do Metro bundler (`package.json` scripts, `app.json`, `metro.config.js`)
   - Tipo de aplicação (Expo managed, Expo bare, React Native CLI, Flutter, nativo)
   - Navegação (Expo Router, React Navigation, Navigator 2.0, etc.)
   - Plataformas alvo (iOS, Android, ambas)
   - Ferramentas de teste (Jest, Detox, Maestro, RNTL — inferir de `package.json` e `devDependencies`)
8. Confirmar com o humano o escopo antes de prosseguir (especialmente se `{SCOPE}=full`).

---

## Passo 1 — Codebase Knowledge Discovery (`acquire-codebase-knowledge`) 🔍

**Objetivo:** Gerar base sólida de 7 documentos estruturados sobre stack, estrutura, arquitetura, convenções, integrações, testes e concerns.

### Pré-condição:
- `{SOLUTION_ROOT}` acessível com `package.json`, `pubspec.yaml`, `Podfile` ou equivalente.
- Python 3.8+ disponível para rodar `scan.py`.

### Ações:
1. **Rodar o scanner:**
   ```bash
   python3 ~/.claude/skills/acquire-codebase-knowledge/scripts/scan.py --output docs/codebase/.codebase-scan.txt
   ```
2. **Ler** arquivos de intenção: `README.md`, `package.json`, `app.json`, `eas.json`, `tsconfig.json`, e quaisquer `PRD.md`, `SPEC.md` ou `ROADMAP.md` existentes.
3. **Ler** arquivos de configuração específicos mobile:
   - `app.json` / `app.config.ts` — configuração Expo (scheme, plugins, splash, updates)
   - `eas.json` — configuração de build EAS (profiles, channels, credentials)
   - `metro.config.js` — configuração do Metro bundler
   - `tailwind.config.js` — se usar NativeWind/Tailwind
   - `ios/Podfile` — dependências nativas iOS
   - `android/build.gradle` — configuração Android
4. **Preencher** os 7 templates na ordem:
   1. `STACK.md` — linguagem, runtime, frameworks, dependências npm/CocoaPods/Gradle, versões de SDK
   2. `STRUCTURE.md` — layout de diretórios, entry points (`App.tsx`, `index.js`, `_layout.tsx`), arquivos de configuração chave
   3. `ARCHITECTURE.md` — camadas (screens→components→hooks→services), padrões (Container/Presenter, hooks composition, Context API), fluxo de dados, navegação
   4. `CONVENTIONS.md` — naming, formatação, error handling, imports, padrões React Native/Expo, organização de estilos
   5. `INTEGRATIONS.md` — APIs externas, push notifications (FCM/APNs), analytics (Firebase/Amplitude), crash reporting (Sentry/Crashlytics), armazenamento local (AsyncStorage/SQLite/MMKV)
   6. `TESTING.md` — frameworks (Jest, RNTL, Detox, Maestro), organização de testes, mocking strategy (MSW, Nock, mocks manuais), device/simulator testing
   7. `CONCERNS.md` — dívida técnica, bugs conhecidos, riscos de segurança OWASP Mobile Top 10, gargalos de performance, dependências desatualizadas
5. **Validar** cada doc contra `references/inquiry-checkpoints.md`.
6. **Marcar** desconhecidos como `[TODO]` e questões de intenção como `[ASK USER]`.

### Validação:
- [ ] `scan.py` executado sem erros.
- [ ] Todos os 7 arquivos existem em `docs/codebase/`.
- [ ] Nenhuma seção obrigatória vazia.
- [ ] Configurações específicas mobile (`app.json`, `eas.json`, `Podfile`, `build.gradle`) analisadas.
- [ ] Desconhecidos usam `[TODO]`, não suposições.
- [ ] `[ASK USER]` compilados para apresentação ao final.

---

## Passo 2 — Domain Modeling (`domain-modeling`) 🏛️

**Objetivo:** Garantir que o glossário de domínio e a linguagem ubíqua estejam corretos e completos **antes** de documentar API e arquitetura.

### Pré-condição:
- Passo 1 concluído (entendimento da codebase).

### Ações:
1. **Ler** `domain/domain.md` (se existir) e extrair todos os termos definidos.
2. **Cross-referenciar** com o código-fonte em `{SOLUTION_ROOT}`:
   - Nomes de interfaces TypeScript, tipos, enums, constantes.
   - Termos usados em nomes de props, estados, variáveis, e parâmetros de navegação.
   - Comentários e docstrings que definem conceitos de negócio.
   - Models de banco de dados local (SQLite, Realm, WatermelonDB).
   - Tipos de requisição/resposta de API.
3. **Identificar ambiguidades**: termos iguais com significados diferentes, termos ausentes no glossário mas usados no código.
4. **Resolver com o humano**: para cada ambiguidade, perguntar "no código usa-se X, mas o glossário define Y — qual é o correto?"
5. **Atualizar** `domain/domain.md`:
   - Adicionar termos ausentes.
   - Corrigir definições inconsistentes.
   - Manter agrupamento por subdomínio.
6. **Criar** `domain/domain.md` se não existir, usando o template:

```markdown
---
title: "Domínio — {SOLUTION_NAME}"
version: "1.0"
date_created: "{AAAA-MM-DD}"
last_updated: "{AAAA-MM-DD}"
owner: "Time de Engenharia"
tags: ["domain", "glossary", "ubiquitous-language"]
---

# Glossário de Domínio — {SOLUTION_NAME}

## 1. Introdução
{Escopo do domínio: o que esta solução modela e quais as fronteiras do negócio.}

## 2. Termos de Domínio
### 2.1. {Subdomínio 1}
| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| ... | ... | ... | ... |

## 3. Relações entre Conceitos
{Diagrama ou descrição de como os principais conceitos se relacionam.}

## 4. Regras de Negócio Fundamentais
{Lista das regras mais estáveis que todo o time deve conhecer.}
```

### Validação:
- [ ] Todo termo do glossário aparece no código com o mesmo significado.
- [ ] Nenhum termo do código (interface, type, enum) está ausente do glossário.
- [ ] Ambiguidades resolvidas e registradas.

---

## Passo 3 — API Design (`api-designer`) 🔌

**Objetivo:** Produzir/atualizar o contrato OpenAPI 3.1 dos endpoints consumidos pelo aplicativo mobile.

### Pré-condição:
- `domain/domain.md` deve estar atualizado.
- `docs/codebase/INTEGRATIONS.md` disponível como referência.

### Ações:
1. **Ler** o código-fonte em `{SOLUTION_ROOT}` para descobrir chamadas de API:
   - Hooks/services que usam `fetch`, `axios`, ou clients HTTP.
   - Interfaces TypeScript de request/response.
   - URLs, métodos HTTP, headers (incluindo auth tokens).
   - Códigos de erro e formato de resposta.
   - Estratégias de retry, timeout, e cache.
2. **Ler** `api/{api-name}.yaml` (se existir) para baseline.
3. **Comparar** endpoints documentados vs. endpoints consumidos:
   - Chamadas documentadas mas não usadas → marcar como `x-unused: true`.
   - Chamadas usadas mas não documentadas → adicionar à spec.
4. **Gerar/atualizar** `api/{api-name}.yaml` com:
   - `openapi: 3.1.0`
   - `info` (title, version, description)
   - `servers` (local, staging, production)
   - `paths` (todos os endpoints consumidos com métodos HTTP, parâmetros, bodies, responses)
   - `components/schemas` (todos os DTOs com exemplos)
   - `components/securitySchemes` (JWT Bearer, API Key, etc.)
   - `components/responses` (erros padrão)
   - `tags` (agrupamento por recurso)
5. **Validar** a spec gerada:
   ```bash
   npx @redocly/cli lint {SPECS_DIR}/api/{api-name}.yaml
   ```
6. **Atualizar** `engineering/api-guidelines.md` com padrões específicos mobile:
   - Estratégia offline-first (cache, fila de requisições pendentes)
   - Tratamento de network errors e timeout
   - Retry com exponential backoff
   - Estrutura de hooks customizados para queries/mutations (React Query, SWR)

### Validação:
- [ ] Spec OpenAPI passa na validação `@redocly/cli lint` sem erros.
- [ ] Todas as chamadas de API do código estão documentadas.
- [ ] Schemas de request/response batem com as interfaces TypeScript.
- [ ] Exemplos de request/response incluídos para cada endpoint.
- [ ] Padrões mobile (offline-first, cache, retry) documentados em `api-guidelines.md`.

---

## Passo 4 — Architecture Design (`architecture-designer`) 🏗️

**Objetivo:** Produzir/atualizar a documentação arquitetural nos níveis C4 aplicáveis, integrações, arquitetura de navegação e ADRs.

### Pré-condição:
- `domain/domain.md` e `api/{api-name}.yaml` atualizados.
- `docs/codebase/ARCHITECTURE.md` disponível como referência.

### Ações:

#### 4.1 — C4 Model (níveis aplicáveis a mobile)

| Nível | Arquivo | Conteúdo |
|---|---|---|
| **Context** | `c4-context.md` | Sistema no ecossistema: usuários mobile, sistemas externos (APIs, push notifications, auth, analytics) |
| **Containers** | `c4-containers.md` | Containers de deploy: app iOS (App Store), app Android (Google Play), Metro bundler, API backend, push notification service, crash reporting, analytics |
| **Components** | `c4-components.md` | Árvore de componentes mobile: App → Navigators → Screens → Components → Hooks → Services |

Cada arquivo deve conter:
1. Diagrama **Mermaid** (obrigatório).
2. Descrição textual de cada elemento do diagrama.
3. Fluxos de dados principais (sequência de chamadas, fluxo de navegação).

#### 4.2 — Arquitetura de Navegação

1. **Mapear** a estrutura completa de navegação:
   - Rotas (Expo Router, React Navigation stacks/tabs/drawers)
   - Deep links e universal links
   - Parâmetros de rota e tipos
   - Fluxos: autenticação → onboarding → home → feature
   - Modals, bottom sheets, e navegação condicional
2. **Gerar** `architecture/navigation-architecture.md` com:
   - Diagrama de árvore de navegação (Mermaid flowchart)
   - Tabela de rotas com parâmetros e guards
   - Deep link scheme e universal link configuration
   - Fluxos de navegação por user story

#### 4.3 — Arquitetura de Plataforma

1. **Documentar** diferenças específicas por plataforma:
   - Módulos nativos e bridging (React Native)
   - Platform-specific code (`Platform.OS`, `.ios.ts` vs `.android.ts`)
   - APIs nativas acessadas (câmera, geolocalização, biometria, etc.)
   - Configurações de build por plataforma
2. **Gerar** `architecture/platform-architecture.md`.

#### 4.4 — Integrações

1. **Mapear** todas as dependências externas no código:
   - APIs REST consumidas (fetch/axios).
   - Push notifications (Firebase Cloud Messaging, Apple Push Notification service).
   - Analytics (Firebase Analytics, Amplitude, Mixpanel).
   - Crash reporting (Sentry, Crashlytics).
   - Mapas e geolocalização.
   - Pacotes npm/CocoaPods/Gradle críticos.
   - Serviços de autenticação (OAuth, JWT, Firebase Auth).
2. **Gerar/atualizar** `architecture/integrations.md`:
   - Para cada integração: endpoint/URL, propósito, contrato, timeout, retry policy.
   - Diagrama de dependências entre serviços.
   - Tabela de dependências com versão e propósito.

#### 4.5 — ADRs

1. **Identificar** decisões arquiteturais merecedoras de ADR:
   - Escolha de framework (Expo managed vs. bare vs. React Native CLI vs. Flutter).
   - Escolha de navegação (Expo Router vs. React Navigation).
   - Estratégia de estilização (NativeWind vs. StyleSheet vs. Tamagui vs. Restyle).
   - Estratégia de estado (Context vs. Zustand vs. Redux Toolkit vs. Jotai).
   - Estratégia de armazenamento local (AsyncStorage vs. SQLite vs. MMKV vs. WatermelonDB).
   - Estratégia de deploy (EAS Build vs. Fastlane vs. manual).
   - Estratégia de atualizações OTA (Expo Updates vs. CodePush).
2. **Criar/atualizar** ADRs em `architecture/adrs/` usando template padrão.
3. **Atualizar** `architecture/adrs/INDEX.md` com a lista cronológica.

#### 4.6 — Visão Geral de Arquitetura

1. **Gerar/atualizar** `architecture/architecture.md` consolidando:
   - Resumo executivo (1 parágrafo).
   - Padrão arquitetural adotado (com justificativa).
   - Árvore de componentes principal (comentada).
   - Stack tecnológica completa (tabela com versões).
   - Princípios de design (component-based, hooks pattern, container/presenter).
   - Cross-cutting concerns (auth, logging, error handling, loading/error/empty states, offline mode).

### Validação:
- [ ] Todos os diagramas Mermaid renderizam corretamente.
- [ ] C4 Components cobre a árvore principal de componentes e navegação.
- [ ] `navigation-architecture.md` documenta todas as rotas e fluxos.
- [ ] Integrações cobrem todas as dependências externas e pacotes críticos.
- [ ] ADRs seguem template e têm status definido.

---

## Passo 5 — Mobile UI/UX Audit (`impeccable` + `react-native-expert`) 🎨

**Objetivo:** Realizar auditoria técnica, avaliação heurística de UX adaptada para mobile, e extrair design system nativo da aplicação.

### Pré-condição:
- Aplicação mobile executável (simulador iOS/Android ou device físico disponível) — para `audit` com screenshots.
- Passos 1-4 concluídos (contexto técnico estabelecido).

### Ações:

#### 5.1 — Executar `impeccable document` (DESIGN.md)

1. Rodar `node .agents/skills/impeccable/scripts/context.mjs` para detectar PRODUCT.md/DESIGN.md existentes.
2. Se `NO_PRODUCT_MD`, criar `product/product.md` antes de prosseguir.
3. Carregar `reference/product.md` (app UI → design SERVE o produto).
4. **Gerar** `design/DESIGN.md` com:
   - Design tokens (cores OKLCH, tipografia nativa — SF Pro/iOS, Roboto/Android, spacing, shadows/elevation, border-radius)
   - Componentes existentes com variantes (Button, Card, Input, Modal, BottomSheet, etc.)
   - Padrões de layout mobile (SafeAreaView, KeyboardAvoidingView, ScrollView vs. FlatList)
   - Estados: default, pressed, focused, disabled, loading, error, empty
   - Responsive breakpoints (small phone, large phone, tablet)
   - Animações e transições (Reanimated, Moti, layout animations)
5. **Extrair** `design/design-tokens.md` e `design/components.md`:
   - Tokens extraídos diretamente do código (tailwind.config.js, NativeWind theme, StyleSheet constants)
   - Componentes catalogados com props TypeScript e variantes
6. **Gerar** `design/animations.md` (se aplicável):
   - Animações de entrada/saída de telas
   - Micro-interações (button press feedback com Haptics)
   - Animações Lottie, Moti, Reanimated

#### 5.2 — Executar `impeccable audit` (Relatório Técnico Mobile)

1. Carregar `reference/audit.md`.
2. **Auditar** a aplicação nos seguintes critérios:
   - **Acessibilidade mobile:** contraste (≥4.5:1 body, ≥3:1 large text), touch targets (≥48x48dp Android, ≥44x44pt iOS), labels de acessibilidade, suporte a font scaling
   - **Performance:** tempo de startup, render performance (FlatList virtualization, memo), tamanho de bundle/assets, uso de memória
   - **Responsividade cross-device:** comportamento em small phone (iPhone SE 375pt), standard (iPhone 14 390pt), large phone (iPhone Pro Max 430pt), tablet (iPad 768pt+)
   - **Qualidade de código:** componentes, props typing, hooks usage, regras do ESLint, tratamento de plataforma (Platform.OS)
   - **Estados:** loading, empty, error, edge cases, comportamento offline, gestão de permissões
   - **Segurança mobile:** armazenamento seguro de tokens, SSL pinning, ofuscação de código, proteção de dados sensíveis
3. **Gerar** `mobile/audit-report.md` com:
   - Checklist de aprovação/reprovação por critério
   - Issues encontrados com severidade (P0/P1/P2)
   - Screenshots de problemas detectados (capturados do simulador/device)
   - Recomendações de correção com exemplos de código

#### 5.3 — Executar `impeccable critique` (UX Heurística Mobile)

1. Carregar `reference/critique.md`.
2. **Avaliar** a aplicação usando heurísticas de Nielsen **adaptadas para mobile**:
   - **Visibility of system status:** feedback tátil (Haptics), indicadores de loading, status bar, notificações in-app
   - **Match between system and real world:** ícones compreensíveis, linguagem natural, metáforas familiares
   - **User control and freedom:** gesto de voltar (Android back, iOS swipe), undo, navegação clara
   - **Consistency and standards:** padrões iOS (Human Interface Guidelines) e Android (Material Design), comportamento esperado de gestos
   - **Error prevention:** validação de formulários em tempo real, confirmação de ações destrutivas, prevenção de toques acidentais
   - **Recognition rather than recall:** navegação visível (bottom tabs, drawer), ações contextuais acessíveis
   - **Flexibility and efficiency of use:** atalhos, deep links, gestos avançados, pull-to-refresh, swipe actions
   - **Aesthetic and minimalist design:** densidade de informação adequada ao tamanho da tela, hierarquia visual clara
   - **Help users recognize, diagnose, and recover from errors:** mensagens de erro claras, retry actions, suporte offline
   - **Help and documentation:** onboarding, tooltips, empty states educativos
3. **Gerar** `mobile/ux-critique.md` com:
   - Score por heurística (1-5)
   - Pontos fortes e fracos
   - Evidências (screenshots do simulador/device)
   - Recomendações priorizadas

#### 5.4 — Revisão Técnica (`react-native-expert`)

1. **Revisar** padrões específicos React Native/Expo:
   - Uso correto de FlatList/FlashList (keyExtractor, getItemLayout, windowSize)
   - Performance de Reanimated (useSharedValue, useAnimatedStyle)
   - Gerenciamento de memória (cleanup de listeners, timers, subscriptions)
   - Uso correto de React Navigation (prevenção de memory leaks, tipo seguro de rotas)
   - Tratamento de platform-specific code
2. **Atualizar** `engineering/performance-guidelines.md` com recomendações.
3. **Atualizar** `engineering/state-management.md` com a estrutura de stores/contexts documentada.

### Validação:
- [ ] `design/DESIGN.md` gerado com tokens e componentes mobile.
- [ ] `design/design-tokens.md` extraído do código (NativeWind theme ou StyleSheet constants).
- [ ] `design/components.md` cataloga componentes com variantes e platform specifics.
- [ ] `mobile/audit-report.md` cobre a11y mobile, performance, responsividade cross-device, e qualidade.
- [ ] `mobile/ux-critique.md` cobre as 10 heurísticas adaptadas para mobile.
- [ ] Screenshots do simulador/device incluídos nos reports.
- [ ] `engineering/performance-guidelines.md` e `engineering/state-management.md` atualizados.

---

## Passo 6 — Inspeção Visual Mobile 👁️

**Objetivo:** Inspecionar visualmente a aplicação rodando no simulador/device, detectar problemas de layout, adaptação a diferentes tamanhos de tela e consistência visual.

### Pré-condição:
- Aplicação rodando no simulador iOS/Android ou device físico via Metro bundler (`npx expo start` ou `npx react-native start`).
- Capacidade de capturar screenshots do simulador/device.

### Ações:
1. **Iniciar** o Metro bundler e o app no simulador:
   ```bash
   # Expo
   npx expo start --ios     # ou --android
   # React Native CLI
   npx react-native run-ios # ou run-android
   ```
2. **Capturar screenshots** em múltiplos dispositivos/resoluções:
   | Nome | Dispositivo | Resolução (pts) | Densidade |
   |---|---|---|---|
   | iPhone SE | iPhone SE (3rd gen) | 375×667 | @2x |
   | iPhone Standard | iPhone 14 | 390×844 | @3x |
   | iPhone Pro Max | iPhone 15 Pro Max | 430×932 | @3x |
   | iPad | iPad (10th gen) | 744×1133 | @2x |
   | Android Small | Pixel 6a | 393×851 | 2.75x |
   | Android Large | Pixel 7 Pro | 412×915 | 3.5x |
   | Android Tablet | Samsung Galaxy Tab | 800×1280 | 2x |
3. **Detectar** problemas específicos mobile:
   - **Safe Areas:** Conteúdo invadindo notch/Dynamic Island, safe area insets incorretos
   - **Keyboard Avoidance:** Formulários não ajustando quando teclado aparece (KeyboardAvoidingView)
   - **Layout:** Element overflow em telas pequenas, texto truncado
   - **Responsivo:** Quebra de layout em tablets vs. phones, orientação portrait vs. landscape
   - **Toque:** Touch targets pequenos (<44pt iOS, <48dp Android), elementos muito próximos
   - **Acessibilidade visual:** Contraste, font scaling (acessibilidade de tamanho de fonte do sistema)
   - **Consistência entre plataformas:** Diferenças iOS vs. Android (sombras, bordas, tipografia, navegação)
   - **Status Bar:** Cores, transparência, conteúdo atrás da status bar
   - **Home Indicator:** Conteúdo atrás do home indicator (iPhones sem botão)
4. **Navegar** por todos os estados da aplicação:
   - Splash screen
   - Tela de onboarding (se existir)
   - Tela de login/autenticação
   - Telas principais com conteúdo
   - Estado de loading (skeleton, spinner)
   - Estado de lista vazia
   - Estado de erro (com retry)
   - Modal e bottom sheets
   - Comportamento offline (se aplicável)
   - Deep link / universal link
   - Permissões do sistema (câmera, localização, notificações)
5. **Gerar** `mobile/visual-inspection-report.md` com:
   - Screenshots por dispositivo para cada tela
   - Issues detectados com severidade (P1/P2/P3)
   - Localização no código-fonte de cada issue (arquivo + linha)
   - Screenshots comparativos iOS vs. Android (se aplicável)
   - Recomendações de correção com exemplos de código

### Validação:
- [ ] Screenshots capturados em todos os dispositivos listados.
- [ ] Todos os estados da aplicação inspecionados.
- [ ] Safe areas, notch, keyboard avoidance verificados.
- [ ] Issues mapeados para arquivos fonte específicos.
- [ ] Relatório inclui tabela de severidade e comparação cross-platform.

---

## Passo 7 — Auditoria de Acessibilidade Mobile ♿

**Objetivo:** Auditoria completa de acessibilidade para plataformas móveis, identificando barreiras para usuários com deficiência e fornecendo remediação.

### Pré-condição:
- Aplicação rodando no simulador/device (iOS com VoiceOver habilitado, Android com TalkBack habilitado).

### Ações:
1. **Confirmar escopo:** Nível de conformidade (WCAG 2.1 AA adaptado para mobile / diretrizes específicas de plataforma), telas alvo, fluxos de usuário chave.
2. **Executar verificações automatizadas:**
   - iOS: Accessibility Inspector (Xcode) — audit de todas as telas
   - Android: Accessibility Scanner (Google) — scan de todas as telas
   - React Native: `eslint-plugin-react-native-a11y` (lint de accessibility props)
3. **Realizar verificações manuais:**
   - **Leitor de tela (VoiceOver/TalkBack):**
     - Todos os elementos interativos têm labels descritivas
     - Ordem de foco é lógica e segue o fluxo visual
     - Imagens decorativas marcadas como `accessible={false}`
     - Agrupamentos de elementos (accessibilityRole, accessibilityLabel)
   - **Touch targets:**
     - Todos os elementos interativos ≥ 44×44pt (iOS) / 48×48dp (Android)
     - Espaçamento mínimo entre elementos tocáveis
   - **Contraste de cores:**
     - ≥ 4.5:1 para texto normal
     - ≥ 3:1 para texto grande (≥18pt ou ≥14pt bold)
     - Verificar em diferentes modos (claro/escuro)
   - **Font scaling:**
     - App responde corretamente ao Dynamic Type (iOS) / Font Size (Android)
     - Texto não é cortado quando escala em 200%
   - **Navegação alternativa:**
     - App é navegável com switch control (iOS) / switch access (Android)
     - Gestos alternativos disponíveis para ações complexas
   - **Conteúdo dinâmico:**
     - Mudanças de tela são anunciadas pelo leitor de tela
     - Animações respeitam `prefers-reduced-motion`
4. **Mapear findings para diretrizes de acessibilidade mobile:**

   | Finding | Plataforma | Diretriz | Severidade | Impacto no Usuário |
   |---|---|---|---|---|
   | Botão sem accessibilityLabel | iOS + Android | WCAG 4.1.2 / HIG A11y | P1 | Leitores de tela não identificam a ação |
   | Touch target < 44pt | iOS | HIG A11y / WCAG 2.5.5 | P1 | Usuários com destreza reduzida não conseguem tocar |
   | Contraste insuficiente | Android | Material A11y / WCAG 1.4.3 | P2 | Usuários com baixa visão não leem o texto |

5. **Gerar** `mobile/accessibility-audit.md` com:
   - Sumário executivo (N issues por severidade e plataforma)
   - Lista detalhada de violações com mapeamento para diretrizes
   - Screenshots evidenciando problemas
   - Recomendações de remediação com exemplos de código React Native/Expo
   - Checklist de verificação para re-teste após correções
   - Guia rápido: accessibility props essenciais (accessibilityLabel, accessibilityRole, accessibilityState, accessibilityHint)

### Validação:
- [ ] Todos os critérios de acessibilidade mobile aplicáveis foram verificados.
- [ ] Cada finding mapeado para uma diretriz (WCAG + HIG/Material Design).
- [ ] Remediações incluem exemplos de código React Native/Flutter/nativo.
- [ ] Leitor de tela testado em todas as telas (VoiceOver iOS + TalkBack Android).
- [ ] Font scaling e contraste verificados em ambas as plataformas.

---

## Passo 8 — Verificação Funcional Mobile ✅

**Objetivo:** Verificar funcionalmente os fluxos principais da aplicação mobile, validando comportamento real em simulador/device.

### Pré-condição:
- Aplicação rodando no simulador/device.
- Ferramentas de teste configuradas (`{TESTING_FRAMEWORK}` definido nos parâmetros).

### Ações:
1. **Mapear fluxos de usuário** a partir do código, navegação e docs existentes:
   - Fluxo 1: Splash → Onboarding → Home
   - Fluxo 2: Autenticação (login, cadastro, recuperação de senha, biometria)
   - Fluxo 3: CRUD principal da aplicação (criar, listar, editar, deletar)
   - Fluxo 4: Navegação entre tabs/seções
   - Fluxo 5: Tratamento de erros e estados offline
   - Fluxo 6: Deep links e universal links
   - Fluxo 7: Permissões do sistema (câmera, localização, notificações)
2. **Executar cada fluxo:**
   - Abrir o app no simulador/device
   - Navegar, preencher campos, interagir com elementos
   - Capturar screenshots de cada etapa
   - Verificar estados intermediários (loading, disabled, skeleton)
   - Verificar resultado final esperado
   - Testar cenários de erro (API offline, timeout, dados inválidos)
3. **Executar testes automatizados** (se framework disponível):
   - **React Native Testing Library + Jest:** Testes unitários e de integração de componentes
     ```bash
     npx jest --coverage
     ```
   - **Detox (E2E):** Testes end-to-end no simulador
     ```bash
     npx detox build --configuration ios.sim.debug
     npx detox test --configuration ios.sim.debug
     ```
   - **Maestro (E2E):** Testes end-to-end com YAML flows
     ```bash
     maestro test .maestro/flows/
     ```
4. **Documentar** cada fluxo em `mobile/functional-verification.md`:
   - Passos executados
   - Screenshots de cada etapa
   - Resultado esperado vs. obtido
   - Plataformas testadas (iOS, Android)
   - Problemas encontrados (se houver) com localização no código

### Validação:
- [ ] Todos os fluxos principais verificados em ambas as plataformas.
- [ ] Screenshots capturados em cada etapa.
- [ ] Comportamento real corresponde ao documentado nos requisitos.
- [ ] Cenários de erro e offline testados.
- [ ] Resultados de testes automatizados incluídos (se framework disponível).

---

## Passo 9 — Documentation Writer (`documentation-writer`) 📝

**Objetivo:** Aplicar o framework Diátaxis para revisar, unificar e expandir toda a documentação, garantindo qualidade textual e consistência cross-documento.

### Pré-condição:
- Passos 1 a 8 concluídos (toda a documentação base existe).

### Ações:

#### 9.1 — Classificar Cada Artefato no Framework Diátaxis

| Quadrante | Pergunta-chave | Artefatos `.specs/` + `docs/codebase/` |
|---|---|---|
| **Tutorial** (learning) | "Como eu começo?" | `engineering/code-analysis.md`, `docs/codebase/STRUCTURE.md`, `engineering/state-management.md` |
| **How-to** (problem) | "Como resolvo X?" | `engineering/api-guidelines.md`, `docs/codebase/CONVENTIONS.md`, `engineering/performance-guidelines.md`, `mobile/visual-inspection-report.md`, `mobile/accessibility-audit.md` |
| **Reference** (information) | "O que é X?" | `domain/domain.md`, `api/*.yaml`, `docs/codebase/STACK.md`, `docs/codebase/INTEGRATIONS.md`, `design/DESIGN.md`, `design/design-tokens.md`, `design/components.md`, `design/animations.md`, `product/requirements.md`, `product/feature-roadmap.md`, `security/permissions-manifest.md`, `mobile/store-metadata.md` |
| **Explanation** (understanding) | "Por que X?" | `architecture/architecture.md`, `architecture/adrs/*.md`, `architecture/c4-*.md`, `architecture/navigation-architecture.md`, `architecture/platform-architecture.md`, `docs/codebase/ARCHITECTURE.md`, `docs/codebase/CONCERNS.md`, `product/product.md`, `governance/confidence-report.md`, `mobile/audit-report.md`, `mobile/ux-critique.md`, `mobile/performance-report.md`, `mobile/offline-support-report.md` |

#### 9.2 — Revisar Cada Artefato

Para cada arquivo, verificar:
1. **Clareza**: Linguagem simples, sem jargão desnecessário. Acrônimos definidos na primeira ocorrência.
2. **Precisão**: Código, comandos e exemplos são executáveis e corretos. Comandos específicos de plataforma (`expo`, `npx react-native`, `fastlane`) são verificados.
3. **Consistência**: Terminologia alinhada com `domain/domain.md`. Estilo consistente entre arquivos.
4. **Completude**: O arquivo responde à pergunta-chave do seu quadrante Diátaxis.
5. **Cross-referências**: Links entre documentos funcionam e fazem sentido.
6. **Cobertura multi-plataforma**: iOS e Android são cobertos adequadamente quando aplicável.

#### 9.3 — Criar Artefatos Ausentes (Product, Governance, Security, Store)

**Product:**
- `product/product.md` — Explicação (Diátaxis): visão do produto mobile, proposta de valor, personas mobile.
- `product/requirements.md` — Referência (Diátaxis): RFs, RNFs, regras de negócio com EARS + MoSCoW.
- `product/feature-roadmap.md` — Referência (Diátaxis): features planejadas, concluídas, dívidas técnicas.

**Governance:**
- `governance/inventory.md` — Referência: tabela de todos os módulos, arquivos, cobertura de testes.
- `governance/confidence-report.md` — Explicação: score % de confiança em cada área.

**Security:**
- `security/SECURITY.md` — Referência: SSL pinning, autenticação, armazenamento seguro (Keychain/Keystore), OWASP Mobile Top 10, código de ofuscação, ProGuard/R8.
- `security/permissions-manifest.md` — Referência: tabela de permissões iOS (Info.plist) e Android (AndroidManifest.xml) com justificativa de uso.

**Store:**
- `mobile/store-metadata.md` — Referência: nome do app, descrição curta/longa, keywords, screenshots obrigatórios, política de privacidade, categorias, classificação etária.

#### 9.4 — Atualizar INDEX.md e CHANGELOG.md

1. **Atualizar** `INDEX.md`:
   - Listar todos os arquivos existentes com links e descrição de 1 linha.
   - Agrupar por pasta/seção (incluir `docs/codebase/`).
   - Incluir data da última atualização geral.
   - Incluir score de confiança (de `governance/confidence-report.md`).

2. **Atualizar** `CHANGELOG.md`:
   - Registrar data, arquivos criados/atualizados, e resumo das mudanças.
   - Seguir formato padrão (Criados / Atualizados / Removidos).

### Validação:
- [ ] Todo artefato está classificado num quadrante Diátaxis.
- [ ] Links cross-documento funcionam (entre `.specs/` e `docs/codebase/`).
- [ ] INDEX.md lista todos os arquivos existentes em ambos os diretórios.
- [ ] CHANGELOG.md registra as mudanças desta execução.
- [ ] Terminologia consistente com `domain/domain.md`.
- [ ] Cobertura iOS e Android adequada quando aplicável.
- [ ] Nenhum `[TODO]` ou `[ASK USER]` sem resposta.

---

## Passo 10 — Sanity Check Final 🩺

1. **Links**: Verificar todos os links internos em `INDEX.md` e cross-referências entre documentos.
2. **Mermaid**: Confirmar que diagramas Mermaid em `architecture/c4-*.md`, `architecture/navigation-architecture.md` têm sintaxe válida.
3. **OpenAPI**: Rodar `npx @redocly/cli lint {SPECS_DIR}/api/{api-name}.yaml` e corrigir erros.
4. **Consistência de domínio**: Escolher 5 termos aleatórios do `domain/domain.md` e verificar que aparecem com o mesmo significado nos outros documentos.
5. **Screenshots**: Verificar que todos os screenshots nos reports de mobile são válidos e referenciados corretamente.
6. **Cobertura de plataforma**: Verificar que ambos iOS e Android são cobertos quando aplicável (navigation-architecture, platform-architecture, visual-inspection).
7. **Cobertura**: Comparar estrutura final com a estrutura canônica (Passo 0) e reportar:
   - Artefatos criados: N
   - Artefatos atualizados: N
   - Artefatos ainda ausentes: N (com justificativa)
8. **Git diff**: Executar `git diff --stat {SPECS_DIR}/ docs/codebase/` para auditar todas as mudanças.

---

## Passo 11 — Geração do Relatório de Execução

1. **Nome do arquivo**: `{AAAA-MM-DD-HHMMSS}-mobile-spec-mining.md`
   - Use o timestamp **real do momento da criação**: `date +%Y-%m-%d-%H%M%S`.
2. **Local**: `{SPECS_DIR}/skill-output/`
3. **Conteúdo**:

```markdown
# 📑 Relatório de Mineração de Especificações Mobile (MOBILE-SPEC-MINING)

* **Data e Hora:** {AAAA-MM-DD HH:MM:SS} (GMT-3)
* **Skills orquestradas:** acquire-codebase-knowledge → domain-modeling → api-designer → architecture-designer → impeccable → react-native-expert → Inspeção Visual Mobile → Auditoria de Acessibilidade Mobile → Mobile App Functional Testing → documentation-writer
* **Solução:** {SOLUTION_NAME} — {SOLUTION_DESCRIPTION}
* **Tipo de Aplicação:** {APP_TYPE}
* **Framework:** {LANGUAGE} | **Styling:** {STYLING} | **Navegação:** {NAVIGATION}
* **Plataformas:** {PLATFORM_TARGET}
* **Escopo:** {SCOPE}

---

## 📋 Resumo da Execução

{1 parágrafo descrevendo o que foi feito, quais skills foram usadas e em qual ordem.}

## 📁 Artefatos Processados

| Ação | Arquivo | Skill | Mudança |
|---|---|---|---|
| 🆕 | `docs/codebase/STACK.md` | acquire-codebase-knowledge | Criado — stack completa documentada |
| 🔄 | `domain/domain.md` | domain-modeling | Atualizado — N novos termos |
| ... | ... | ... | ... |

## 📊 Cobertura Final

| Pasta | Arquivos Esperados | Arquivos Existentes | Cobertura |
|---|---|---|---|
| docs/codebase/ | 7 | N | X% |
| domain/ | 1 | N | X% |
| api/ | 1 | N | X% |
| architecture/ | N | N | X% |
| engineering/ | 4 | N | X% |
| product/ | 3 | N | X% |
| design/ | 4 | N | X% |
| mobile/ | 8 | N | X% |
| governance/ | 2 | N | X% |
| security/ | 2 | N | X% |

## ⚠️ Artefatos Ausentes (com justificativa)

- `...` — justificativa

## 🔍 Descobertas Notáveis

{Lista das 5-10 descobertas mais importantes do pipeline}

## 🔗 Links Gerados

- INDEX.md: `{SPECS_DIR}/INDEX.md`
- CHANGELOG.md: `{SPECS_DIR}/CHANGELOG.md`
- Codebase docs: `docs/codebase/`

---

🤖 *Documentação gerada pela orquestração de skills de IA: acquire-codebase-knowledge → domain-modeling → api-designer → architecture-designer → impeccable → react-native-expert → Inspeção Visual Mobile → Auditoria de Acessibilidade Mobile → Mobile App Functional Testing → documentation-writer.*
```

---

# 🔄 Modos de Execução

O prompt suporta 7 modos, controlados pelo parâmetro `{SCOPE}`:

| Modo | `{SCOPE}` | Comportamento |
|---|---|---|
| **Completo** | `full` | Executa Passos 0→11 integralmente, criando/atualizando todos os artefatos |
| **Descoberta apenas** | `discovery-only` | Executa Passos 0→1 + Passo 11 (codebase knowledge + domain) |
| **Design apenas** | `design-only` | Executa Passos 0 + Passos 2→6 + Passo 11 (API + arquitetura + mobile design) |
| **Mobile apenas** | `mobile-only` | Executa Passos 0 + Passos 5→8 + Passo 11 (impeccable + inspeção visual + a11y mobile + testing) |
| **Qualidade apenas** | `qa-only` | Executa Passos 0 + Passos 9→11 (documentation-writer review + sanity check) |
| **Delta** | `delta` | Executa apenas para os arquivos listados em `{TARGET_FILES}`, usando a skill dona correspondente + `documentation-writer` como revisor |
| **Revisão** | `review` | Executa apenas Passo 9 (documentation-writer) + Passo 10 + Passo 11, sem criar novos artefatos |

---

# ⚠️ Regras de Ouro

1. **Nunca sobrescrever sem auditoria**: antes de modificar um arquivo existente, fazer diff do conteúdo atual vs. novo e reportar no relatório final.
2. **Domínio primeiro, mobile depois**: os Passos 1-4 (fundamentos) são pré-requisito para os Passos 5-8 (mobile específico). Se o domínio não estiver estável, não faz sentido auditar UI/UX mobile.
3. **Código é a verdade**: se houver conflito entre documentação existente e código, o código vence. Reportar a divergência e alinhar a documentação.
4. **App deve estar executável para Passos 5, 6, 7, 8**: iniciar o Metro bundler e o app no simulador/device antes desses passos. Se não for possível executar, pular passos que exigem runtime e registrar no relatório.
5. **Humano decide ambiguidades**: o `domain-modeling` pode identificar termos conflitantes, mas apenas o humano pode resolvê-los.
6. **Mermaid é obrigatório**: todo diagrama arquitetural deve ser renderizável como Mermaid (não usar imagens estáticas).
7. **OpenAPI validado**: a spec da API deve passar na validação `@redocly/cli lint` sem erros.
8. **Screenshots são evidência**: todo report de mobile (`audit-report.md`, `visual-inspection-report.md`, `functional-verification.md`, `accessibility-audit.md`) deve incluir screenshots do simulador/device como evidência.
9. **Diátaxis como camada final**: o `documentation-writer` revisa todos os artefatos, mas não redefine decisões técnicas dos passos anteriores — apenas melhora clareza, consistência e organização.
10. **INDEX.md é sagrado**: sempre atualizar o índice central ao final de cada execução, incluindo links para `docs/codebase/`.
11. **Marcar incertezas explicitamente**: usar `[TODO]` para fatos não verificáveis e `[ASK USER]` para decisões que requerem intenção humana.
12. **iOS + Android sempre**: documentação deve cobrir ambas as plataformas quando aplicável. Diferenças significativas entre plataformas devem ser explicitamente documentadas em `architecture/platform-architecture.md`.
13. **OWASP Mobile Top 10**: a auditoria de segurança deve cobrir os riscos específicos de aplicações mobile (M1: Improper Credential Usage, M2: Inadequate Supply Chain Security, M3: Insecure Authentication/Authorization, M4: Insufficient Input/Output Validation, M5: Insecure Communication, M6: Inadequate Privacy Controls, M7: Insufficient Binary Protections, M8: Security Misconfiguration, M9: Insecure Data Storage, M10: Insufficient Cryptography).

---

# 📚 Referências

- **Diátaxis Framework**: [https://diataxis.fr/](https://diataxis.fr/)
- **C4 Model**: [https://c4model.com/](https://c4model.com/)
- **OpenAPI 3.1**: [https://spec.openapis.org/oas/v3.1.0](https://spec.openapis.org/oas/v3.1.0)
- **Mermaid.js**: [https://mermaid.js.org/](https://mermaid.js.org/)
- **ADR Template**: [https://adr.github.io/madr/](https://adr.github.io/madr/)
- **WCAG 2.2**: [https://www.w3.org/TR/WCAG22/](https://www.w3.org/TR/WCAG22/)
- **Apple Human Interface Guidelines — Accessibility**: [https://developer.apple.com/design/human-interface-guidelines/accessibility](https://developer.apple.com/design/human-interface-guidelines/accessibility)
- **Android Accessibility — Material Design**: [https://m3.material.io/foundations/accessible-design](https://m3.material.io/foundations/accessible-design)
- **OWASP Mobile Top 10 (2024)**: [https://owasp.org/www-project-mobile-top-10/](https://owasp.org/www-project-mobile-top-10/)
- **Nielsen's Heuristics**: [https://www.nngroup.com/articles/ten-usability-heuristics/](https://www.nngroup.com/articles/ten-usability-heuristics/)
- **iOS Accessibility Inspector**: [https://developer.apple.com/library/archive/documentation/Accessibility/Conceptual/AccessibilityMacOSX/OSXAXTestingApps.html](https://developer.apple.com/library/archive/documentation/Accessibility/Conceptual/AccessibilityMacOSX/OSXAXTestingApps.html)
- **Android Accessibility Scanner**: [https://play.google.com/store/apps/details?id=com.google.android.apps.accessibility.auditor](https://play.google.com/store/apps/details?id=com.google.android.apps.accessibility.auditor)
- **EARS Format**: `references/ears-format.md`
- **OKLCH Color**: [https://oklch.com/](https://oklch.com/)
- **React Native Testing Library**: [https://callstack.github.io/react-native-testing-library/](https://callstack.github.io/react-native-testing-library/)
- **Detox (E2E Mobile)**: [https://wix.github.io/Detox/](https://wix.github.io/Detox/)
- **Maestro (E2E Mobile)**: [https://maestro.mobile.dev/](https://maestro.mobile.dev/)
- **React Navigation**: [https://reactnavigation.org/](https://reactnavigation.org/)
- **Expo Router**: [https://docs.expo.dev/router/introduction/](https://docs.expo.dev/router/introduction/)
- **NativeWind**: [https://www.nativewind.dev/](https://www.nativewind.dev/)
- **Reanimated**: [https://docs.swmansion.com/react-native-reanimated/](https://docs.swmansion.com/react-native-reanimated/)
- **iOS Human Interface Guidelines**: [https://developer.apple.com/design/human-interface-guidelines](https://developer.apple.com/design/human-interface-guidelines)
- **Android Material Design 3**: [https://m3.material.io/](https://m3.material.io/)
