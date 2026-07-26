# Contexto:
  - Este prompt orquestra **10 skills especializadas** para minerar, criar e atualizar a documentação técnica completa de **aplicações web frontend** no diretório `.specs/`.
  - As skills são organizadas em **3 camadas**: descoberta (codebase + domínio), design & arquitetura (API + arquitetura + frontend), e qualidade (documentação + verificação).
  - A estrutura de pastas alvo (descrita abaixo) expande a referência canônica original com artefatos específicos de frontend: design system, auditoria de UI/UX, acessibilidade, inspeção visual, e verificação funcional.
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
| 2ª | `api-designer` | Contrato OpenAPI 3.1 dos endpoints consumidos pelo frontend, modelagem de recursos, versionamento, segurança de API | `api/` |
| 3ª | `architecture-designer` | C4 (Context→Containers→Components), ADRs, integrações, dicionário de dependências | `architecture/` |
| 4ª | `impeccable` | Auditoria técnica de UI (`audit`), avaliação heurística de UX (`critique`), geração de DESIGN.md (`document`) | `frontend/`, `design/` |
| 5ª | `web-design-reviewer` | Inspeção visual com Playwright — screenshots, detecção de problemas de layout, responsividade, contraste | `frontend/` |
| 6ª | `accessibility-compliance-accessibility-audit` | Auditoria WCAG 2.1/2.2, barreiras de acessibilidade, recomendações de remediação | `frontend/` |

## Camada 3 — Qualidade e Consistência

| Ordem | Skill | Responsabilidade | Pasta(s) Principal(is) |
|---|---|---|---|
| 7ª | `webapp-testing` | Verificação funcional com Playwright — fluxos de usuário, formulários, navegação | `frontend/` |
| 8ª | `documentation-writer` | Diátaxis (Tutorials/How-to/Reference/Explanation), qualidade textual, consistência cross-documento, INDEX.md, CHANGELOG.md | **Todas as pastas** (revisão final) |

---

# ⚙️ Parâmetros de Entrada (preencher antes de executar)

> **Instrução:** No momento de invocar este prompt, o humano deve informar os valores abaixo. Se algum arquivo não existir, marcar como `N/D` (não disponível) e criar no passo apropriado.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_ROOT}` | Caminho absoluto da raiz da solução | `/home/user/work/web_app-tax-nexus-portal` |
| `{SPECS_DIR}` | Caminho relativo da pasta de especificações | `.specs` |
| `{SOLUTION_NAME}` | Nome da solução/aplicação | `web_app-tax-nexus-portal` |
| `{SOLUTION_DESCRIPTION}` | Descrição curta do propósito da solução | `Portal frontend React para simulação da Reforma Tributária 2026 (Tax as a Service)` |
| `{APP_TYPE}` | Tipo de aplicação web | `spa`, `ssr`, `ssg`, `pwa`, `micro-frontend`, `hybrid` |
| `{LANGUAGE}` | Linguagem/framework principal | `typescript/react` |
| `{STYLING}` | Método de estilização | `tailwind`, `css-modules`, `styled-components`, `scss`, `css` |
| `{BACKEND_API}` | URL base da API consumida (se aplicável) | `http://localhost:8080` |
| `{DEV_SERVER_PORT}` | Porta do servidor de desenvolvimento | `5173` |
| `{PROJECT_ID}` | Identificador do projeto corporativo (se aplicável) | `PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO` |
| `{SCOPE}` | Escopo da documentação a ser gerada | `full`, `discovery-only`, `design-only`, `frontend-only`, `qa-only`, `delta` |
| `{TARGET_FILES}` | Lista específica de arquivos a criar/atualizar (se `{SCOPE}=delta`) | `frontend/audit-report.md, design/DESIGN.md` |

---

# 📁 Estrutura de Pastas Alvo (Referência Canônica — Frontend)

> Esta estrutura expande a referência canônica original com pastas específicas de frontend. Pastas/arquivos ausentes são **oportunidade de criação**; existentes são **base para atualização incremental**.

```
{SPECS_DIR}/
├── INDEX.md                                        ← Índice central de toda a documentação
├── CHANGELOG.md                                    ← Histórico de mudanças da documentação
│
├── domain/
│   └── domain.md                                   ← Glossário de domínio + ubiquitous language
│
├── api/
│   └── {api-name}.yaml                             ← Contrato OpenAPI 3.1 (endpoints consumidos)
│
├── architecture/
│   ├── architecture.md                             ← Visão arquitetural geral da solução
│   ├── c4-context.md                               ← Diagrama C4 — Nível 1 (Contexto)
│   ├── c4-containers.md                            ← Diagrama C4 — Nível 2 (Containers)
│   ├── c4-components.md                            ← Diagrama C4 — Nível 3 (Componentes React)
│   ├── integrations.md                             ← Integrações externas, APIs, dependências
│   └── adrs/                                       ← ADRs específicos da solução
│       ├── INDEX.md                                ← Índice cronológico dos ADRs
│       ├── adr-001.md                              ← Ex: Escolha React 19 + Vite 8
│       └── ...
│
├── engineering/
│   ├── api-guidelines.md                           ← Padrões de consumo de API, erros, observabilidade
│   └── code-analysis.md                            ← Análise técnica de fluxo de código
│
├── product/
│   ├── product.md                                  ← Descrição do produto e visão
│   ├── requirements.md                             ← Especificação de requisitos (RF/RNF + MoSCoW)
│   └── feature-roadmap.md                          ← Roadmap de features + dívidas técnicas
│
├── design/                                         ← Design System (populado pelo impeccable)
│   ├── DESIGN.md                                   ← Design system document (gerado por `impeccable document`)
│   ├── design-tokens.md                            ← Tokens de design (cores, tipografia, spacing, shadows)
│   └── components.md                               ← Catálogo de componentes React com variantes
│
├── frontend/                                       ← Análises Específicas de Frontend
│   ├── audit-report.md                             ← Relatório de auditoria técnica (`impeccable audit`)
│   ├── ux-critique.md                              ← Avaliação heurística de UX (`impeccable critique`)
│   ├── visual-inspection-report.md                 ← Inspeção visual com screenshots (`web-design-reviewer`)
│   ├── accessibility-audit.md                      ← Auditoria WCAG (`accessibility-compliance-accessibility-audit`)
│   ├── functional-verification.md                  ← Verificação funcional Playwright (`webapp-testing`)
│   └── performance-report.md                       ← Performance (LCP, CLS, bundle size) — opcional
│
├── governance/
│   ├── inventory.md                                ← Inventário do projeto + cobertura de testes
│   └── confidence-report.md                        ← Relatório de confiança da documentação (%)
│
├── security/
│   └── SECURITY.md                                 ← Definições de segurança (CSP, CORS, OWASP, auth)
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
    └── {AAAA-MM-DD-HHMISS}-frontend-spec-mining.md ← Relatório de execução do pipeline
```

### Diretório Auxiliar: `docs/codebase/` (gerado pelo `acquire-codebase-knowledge`)

```
docs/codebase/
├── STACK.md                                        ← Linguagem, runtime, frameworks, dependências
├── STRUCTURE.md                                    ← Layout de diretórios, entry points, key files
├── ARCHITECTURE.md                                 ← Camadas, padrões, fluxo de dados
├── CONVENTIONS.md                                  ← Naming, formatação, error handling, imports
├── INTEGRATIONS.md                                 ← APIs externas, bancos, auth, monitoring
├── TESTING.md                                      ← Frameworks, organização, mocking strategy
└── CONCERNS.md                                     ← Dívida técnica, bugs, riscos, gargalos
```

---

# 🧭 Mapeamento Skill → Artefato

> Tabela de responsabilidade: qual skill é **dona** (primary owner) de cada artefato. O `documentation-writer` atua como revisor de **todos** os artefatos ao final.

## Artefatos Genéricos

| Artefato | Skill Dona | Ação Típica |
|---|---|---|
| `docs/codebase/STACK.md` | `acquire-codebase-knowledge` | Rodar `scan.py`, preencher template STACK |
| `docs/codebase/STRUCTURE.md` | `acquire-codebase-knowledge` | Mapear diretórios, entry points |
| `docs/codebase/ARCHITECTURE.md` | `acquire-codebase-knowledge` | Documentar camadas e fluxo de dados |
| `docs/codebase/CONVENTIONS.md` | `acquire-codebase-knowledge` | Extrair padrões de código |
| `docs/codebase/INTEGRATIONS.md` | `acquire-codebase-knowledge` | Mapear APIs, auth, monitoring |
| `docs/codebase/TESTING.md` | `acquire-codebase-knowledge` | Documentar frameworks e estratégias |
| `docs/codebase/CONCERNS.md` | `acquire-codebase-knowledge` | Listar dívidas, bugs, riscos |
| `domain/domain.md` | `domain-modeling` | Criar/atualizar glossário, resolver ambiguidades |
| `api/{api-name}.yaml` | `api-designer` | Gerar/atualizar contrato OpenAPI 3.1 |
| `architecture/architecture.md` | `architecture-designer` | Escrever visão geral, padrões, trade-offs |
| `architecture/c4-*.md` | `architecture-designer` | Gerar diagramas C4 nos níveis aplicáveis (Mermaid) |
| `architecture/integrations.md` | `architecture-designer` | Mapear dependências externas e libs |
| `architecture/adrs/*.md` | `architecture-designer` | Registrar decisões arquiteturais (ADR) |
| `engineering/api-guidelines.md` | `api-designer` → `documentation-writer` | Padrões de consumo de API, erros, segurança |
| `engineering/code-analysis.md` | `documentation-writer` | Análise de fluxo de componentes/hooks |
| `product/product.md` | `documentation-writer` | Descrição do produto (Explanation — Diátaxis) |
| `product/requirements.md` | `documentation-writer` | Requisitos (EARS + MoSCoW) |
| `product/feature-roadmap.md` | `documentation-writer` | Roadmap + dívidas técnicas |
| `governance/inventory.md` | `documentation-writer` | Inventário e cobertura |
| `governance/confidence-report.md` | `documentation-writer` | Score de confiança da documentação |
| `security/SECURITY.md` | `documentation-writer` | CSP, CORS, OWASP, auth |
| `INDEX.md` | `documentation-writer` | Índice central (atualizado a cada ciclo) |
| `CHANGELOG.md` | `documentation-writer` | Registro de mudanças da doc |

## Artefatos Específicos de Frontend

| Artefato | Skill Dona | Ação Típica |
|---|---|---|
| `design/DESIGN.md` | `impeccable document` | Gerar documentação do design system a partir do código |
| `design/design-tokens.md` | `impeccable extract` → `documentation-writer` | Extrair tokens de cor, tipografia, spacing, shadows |
| `design/components.md` | `impeccable document` → `documentation-writer` | Catalogar componentes React com variantes e props |
| `frontend/audit-report.md` | `impeccable audit` | Auditoria técnica: a11y, performance, responsividade |
| `frontend/ux-critique.md` | `impeccable critique` | Avaliação heurística de UX com scoring |
| `frontend/visual-inspection-report.md` | `web-design-reviewer` | Screenshots, problemas de layout, contraste |
| `frontend/accessibility-audit.md` | `accessibility-compliance-accessibility-audit` | Violações WCAG, severidade, remediação |
| `frontend/functional-verification.md` | `webapp-testing` | Verificação de fluxos com Playwright |
| `frontend/performance-report.md` | `impeccable optimize` | LCP, CLS, bundle size (opcional) |

---

# 🔁 Protocolo de Execução (Passo a Passo)

## Passo 0 — Leitura do Cenário Atual

1. Confirmar que `{SOLUTION_ROOT}` e `{SPECS_DIR}` estão acessíveis.
2. Confirmar que o servidor de desenvolvimento pode ser iniciado na porta `{DEV_SERVER_PORT}` (necessário para Passos 5, 6, e 7).
3. Ler `{SPECS_DIR}/INDEX.md` (se existir) para entender a estrutura atual.
4. Ler `{SPECS_DIR}/CHANGELOG.md` (se existir) para conhecer o histórico.
5. Fazer scan dos arquivos existentes em cada subpasta e comparar com a estrutura canônica acima.
6. Gerar um **delta inicial**: lista de arquivos existentes (📄), ausentes (❌), e desatualizados (🔄).
7. Detectar automaticamente:
   - Framework e versão (`package.json` → `dependencies`)
   - Método de estilização (tailwind.config, CSS modules, styled-components)
   - Porta do dev server (`vite.config.ts`, `package.json` scripts)
   - Tipo de aplicação (SPA, SSR, SSG — inferir de `next.config`, `vite.config`, `package.json`)
8. Confirmar com o humano o escopo antes de prosseguir (especialmente se `{SCOPE}=full`).

---

## Passo 1 — Codebase Knowledge Discovery (`acquire-codebase-knowledge`) 🔍

**Objetivo:** Gerar base sólida de 7 documentos estruturados sobre stack, estrutura, arquitetura, convenções, integrações, testes e concerns.

### Pré-condição:
- `{SOLUTION_ROOT}` acessível com `package.json` ou equivalente.
- Python 3.8+ disponível para rodar `scan.py`.

### Ações:
1. **Rodar o scanner:**
   ```bash
   python3 ~/.claude/skills/acquire-codebase-knowledge/scripts/scan.py --output docs/codebase/.codebase-scan.txt
   ```
2. **Ler** arquivos de intenção: `README.md`, `package.json`, e quaisquer `PRD.md`, `SPEC.md` ou `ROADMAP.md` existentes.
3. **Preencher** os 7 templates na ordem:
   1. `STACK.md` — linguagem, runtime, frameworks, todas as dependências
   2. `STRUCTURE.md` — layout de diretórios, entry points, arquivos chave
   3. `ARCHITECTURE.md` — camadas, padrões (component-based, hooks, Context API, etc.), fluxo de dados
   4. `CONVENTIONS.md` — naming, formatação, error handling, imports, padrões React
   5. `INTEGRATIONS.md` — APIs externas, auth, monitoring, CDN
   6. `TESTING.md` — frameworks, organização de testes, mocking strategy
   7. `CONCERNS.md` — dívida técnica, bugs, riscos de segurança, gargalos de performance
4. **Validar** cada doc contra `references/inquiry-checkpoints.md`.
5. **Marcar** desconhecidos como `[TODO]` e questões de intenção como `[ASK USER]`.

### Validação:
- [ ] `scan.py` executado sem erros.
- [ ] Todos os 7 arquivos existem em `docs/codebase/`.
- [ ] Nenhuma seção obrigatória vazia.
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
   - Termos usados em nomes de props, estados, e variáveis.
   - Comentários e docstrings que definem conceitos de negócio.
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

**Objetivo:** Produzir/atualizar o contrato OpenAPI 3.1 dos endpoints consumidos pelo frontend.

### Pré-condição:
- `domain/domain.md` deve estar atualizado.
- `docs/codebase/INTEGRATIONS.md` disponível como referência.

### Ações:
1. **Ler** o código-fonte em `{SOLUTION_ROOT}` para descobrir chamadas de API:
   - Hooks que usam `fetch`, `axios`, ou clients HTTP.
   - Interfaces TypeScript de request/response.
   - URLs, métodos HTTP, headers.
   - Códigos de erro e formato de resposta.
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
6. **Atualizar** `engineering/api-guidelines.md` se novos padrões forem identificados.

### Validação:
- [ ] Spec OpenAPI passa na validação `@redocly/cli lint` sem erros.
- [ ] Todas as chamadas de API do código estão documentadas.
- [ ] Schemas de request/response batem com as interfaces TypeScript.
- [ ] Exemplos de request/response incluídos para cada endpoint.

---

## Passo 4 — Architecture Design (`architecture-designer`) 🏗️

**Objetivo:** Produzir/atualizar a documentação arquitetural nos níveis C4 aplicáveis, integrações e ADRs.

### Pré-condição:
- `domain/domain.md` e `api/{api-name}.yaml` atualizados.
- `docs/codebase/ARCHITECTURE.md` disponível como referência.

### Ações:

#### 4.1 — C4 Model (níveis aplicáveis a frontend)

| Nível | Arquivo | Conteúdo |
|---|---|---|
| **Context** | `c4-context.md` | Sistema no ecossistema: atores (usuários), sistemas externos (APIs, auth) |
| **Containers** | `c4-containers.md` | Containers de deploy: SPA bundle, Nginx, CDN, API backend, auth provider |
| **Components** | `c4-components.md` | Árvore de componentes React: App → Pages → Components, hooks, providers |

Cada arquivo deve conter:
1. Diagrama **Mermaid** (obrigatório).
2. Descrição textual de cada elemento do diagrama.
3. Fluxos de dados principais (sequência de chamadas).

#### 4.2 — Integrações

1. **Mapear** todas as dependências externas no código:
   - APIs REST consumidas (fetch/axios).
   - CDNs e serviços de terceiros.
   - Pacotes npm críticos (Recharts, Lucide, etc.).
   - Serviços de autenticação (OAuth, JWT).
2. **Gerar/atualizar** `architecture/integrations.md`:
   - Para cada integração: endpoint/URL, propósito, contrato, timeout, retry policy.
   - Diagrama de dependências entre serviços.
   - Tabela de dependências npm com versão e propósito.

#### 4.3 — ADRs

1. **Identificar** decisões arquiteturais merecedoras de ADR:
   - Escolha de framework (React 19 vs. alternativas).
   - Escolha de bundler (Vite 8 vs. Webpack/Turbopack).
   - Estratégia de estilização (Tailwind vs. CSS Modules vs. styled-components).
   - Estratégia de estado (local state vs. Context vs. Zustand/Redux).
   - Estratégia de deploy (Docker multi-stage vs. Vercel/Netlify).
2. **Criar/atualizar** ADRs em `architecture/adrs/` usando template padrão.
3. **Atualizar** `architecture/adrs/INDEX.md` com a lista cronológica.

#### 4.4 — Visão Geral de Arquitetura

1. **Gerar/atualizar** `architecture/architecture.md` consolidando:
   - Resumo executivo (1 parágrafo).
   - Padrão arquitetural adotado (com justificativa).
   - Árvore de componentes principal (comentada).
   - Stack tecnológica completa (tabela).
   - Princípios de design (component-based, hooks pattern, etc.).
   - Cross-cutting concerns (auth, logging, error handling, loading states).

### Validação:
- [ ] Todos os diagramas Mermaid renderizam corretamente.
- [ ] C4 Components cobre a árvore principal de componentes.
- [ ] Integrações cobrem todas as dependências externas e pacotes críticos.
- [ ] ADRs seguem template e têm status definido.

---

## Passo 5 — Frontend UI/UX Audit (`impeccable`) 🎨

**Objetivo:** Realizar auditoria técnica, avaliação heurística de UX, e extrair design system da aplicação.

### Pré-condição:
- Aplicação em execução em `http://localhost:{DEV_SERVER_PORT}` (para `audit` com screenshots).
- Passos 1-4 concluídos (contexto técnico estabelecido).

### Ações:

#### 5.1 — Executar `impeccable document` (DESIGN.md)

1. Rodar `node .agents/skills/impeccable/scripts/context.mjs` para detectar PRODUCT.md/DESIGN.md existentes.
2. Se `NO_PRODUCT_MD`, criar `product/product.md` antes de prosseguir.
3. Carregar `reference/product.md` (app UI → design SERVE o produto).
4. **Gerar** `design/DESIGN.md` com:
   - Design tokens (cores OKLCH, tipografia, spacing, shadows, border-radius)
   - Componentes existentes com variantes
   - Padrões de layout e grid
   - Estados: default, hover, focus, disabled, loading, error, empty
   - Responsive breakpoints

#### 5.2 — Executar `impeccable audit` (Relatório Técnico)

1. Carregar `reference/audit.md`.
2. **Auditar** a aplicação nos seguintes critérios:
   - **Acessibilidade:** contraste (≥4.5:1 body, ≥3:1 large text), focus states, labels, alt text, ARIA
   - **Performance:** bundle size, render blocking, imagem não otimizada
   - **Responsividade:** comportamento em 375px, 768px, 1280px, 1920px
   - **Qualidade de código:** componentes, props typing, hooks usage
   - **Estados:** loading, empty, error, edge cases
3. **Gerar** `frontend/audit-report.md` com:
   - Checklist de aprovação/reprovação por critério
   - Issues encontrados com severidade (P0/P1/P2)
   - Screenshots de problemas detectados
   - Recomendações de correção

#### 5.3 — Executar `impeccable critique` (UX Heurística)

1. Carregar `reference/critique.md`.
2. **Avaliar** a aplicação usando heurísticas de Nielsen:
   - Visibility of system status
   - Match between system and real world
   - User control and freedom
   - Consistency and standards
   - Error prevention
   - Recognition rather than recall
   - Flexibility and efficiency of use
   - Aesthetic and minimalist design
   - Help users recognize, diagnose, and recover from errors
   - Help and documentation
3. **Gerar** `frontend/ux-critique.md` com:
   - Score por heurística (1-5)
   - Pontos fortes e fracos
   - Recomendações priorizadas

### Validação:
- [ ] `design/DESIGN.md` gerado com tokens e componentes.
- [ ] `frontend/audit-report.md` cobre a11y, performance, responsividade, e qualidade.
- [ ] `frontend/ux-critique.md` cobre as 10 heurísticas de Nielsen.
- [ ] Screenshots incluídos nos reports.

---

## Passo 6 — Inspeção Visual (`web-design-reviewer`) 👁️

**Objetivo:** Inspecionar visualmente a aplicação rodando, detectar problemas de layout, responsividade e consistência visual.

### Pré-condição:
- Aplicação rodando em `http://localhost:{DEV_SERVER_PORT}`.
- Playwright MCP disponível (`browser_navigate`, `browser_snapshot`, `browser_take_screenshot`, `browser_resize`).

### Ações:
1. **Navegar** para `http://localhost:{DEV_SERVER_PORT}`.
2. **Capturar screenshots** em 4 viewports:
   | Nome | Largura | Dispositivo Representativo |
   |---|---|---|
   | Mobile | 375px | iPhone SE |
   | Tablet | 768px | iPad |
   | Desktop | 1280px | Standard PC |
   | Wide | 1920px | Large display |
3. **Detectar** problemas:
   - **Layout:** Element overflow, overlap, quebra de grid/flex
   - **Responsivo:** Quebra em breakpoints específicos, touch targets pequenos
   - **Acessibilidade visual:** Contraste, focus visível, texto cortado
   - **Consistência:** Fontes misturadas, cores inconsistentes, espaçamento irregular
4. **Navegar** por todos os estados da aplicação:
   - Tela de login
   - Tela principal com formulário
   - Estado de loading (botão "Processando...")
   - Estado de resultado (cards + gráfico)
   - Estado de erro (se possível simular)
5. **Gerar** `frontend/visual-inspection-report.md` com:
   - Screenshots antes/depois para cada viewport
   - Issues detectados com severidade (P1/P2/P3)
   - Localização no código-fonte de cada issue
   - Recomendações de correção

### Validação:
- [ ] Screenshots capturados em todos os 4 viewports.
- [ ] Todos os estados da aplicação inspecionados.
- [ ] Issues mapeados para arquivos fonte específicos.
- [ ] Relatório inclui tabela de severidade.

---

## Passo 7 — Auditoria de Acessibilidade (`accessibility-compliance-accessibility-audit`) ♿

**Objetivo:** Auditoria WCAG 2.1/2.2 completa, identificando barreiras de acessibilidade e fornecendo remediação.

### Pré-condição:
- Aplicação rodando em `http://localhost:{DEV_SERVER_PORT}`.
- Playwright MCP disponível.

### Ações:
1. **Confirmar escopo:** WCAG nível (A/AA/AAA), páginas alvo, fluxos de usuário chave.
2. **Executar scans automatizados** com ferramentas de a11y (axe-core, Lighthouse).
3. **Realizar verificações manuais:**
   - Navegação por teclado (Tab, Enter, Escape, arrow keys)
   - Leitor de tela (VoiceOver/NVDA)
   - Ordem de foco (focus order lógica?)
   - Contraste de cores (todas as combinações texto/fundo)
   - Redimensionamento de texto (200% sem perda de conteúdo)
4. **Mapear findings para critérios WCAG:**
   | Finding | WCAG SC | Severidade | Impacto no Usuário |
   |---|---|---|---|
5. **Gerar** `frontend/accessibility-audit.md` com:
   - Sumário executivo (N issues por severidade)
   - Lista detalhada de violações com WCAG SC mapping
   - Recomendações de remediação com exemplos de código
   - Checklist de verificação para re-teste

### Validação:
- [ ] Todos os critérios WCAG nível AA aplicáveis foram verificados.
- [ ] Cada finding mapeado para um WCAG Success Criterion.
- [ ] Remediações incluem exemplos de código.
- [ ] Navegação por teclado testada em todos os estados.

---

## Passo 8 — Verificação Funcional (`webapp-testing`) ✅

**Objetivo:** Verificar funcionalmente os fluxos principais da aplicação com Playwright, validando comportamento real.

### Pré-condição:
- Aplicação rodando em `http://localhost:{DEV_SERVER_PORT}`.
- Playwright MCP disponível (`browser_navigate`, `browser_click`, `browser_fill_form`, `browser_snapshot`).

### Ações:
1. **Mapear fluxos de usuário** a partir do código e dos docs existentes:
   - Fluxo 1: Login com CNPJ → Acesso ao simulador
   - Fluxo 2: Preencher formulário → Clicar Simular → Ver resultados
   - Fluxo 3: Trocar CNPJ (logout) → Retornar à tela de login
   - Fluxo 4: Tratamento de erros (CNPJ inválido, API offline)
2. **Executar cada fluxo** com Playwright:
   - Navegar, preencher campos, clicar botões
   - Capturar screenshots de cada etapa
   - Verificar estados intermediários (loading, disabled)
   - Verificar resultado final esperado
3. **Documentar** cada fluxo em `frontend/functional-verification.md`:
   - Passos executados
   - Screenshots de cada etapa
   - Resultado esperado vs. obtido
   - Problemas encontrados (se houver)

### Validação:
- [ ] Todos os fluxos principais verificados.
- [ ] Screenshots capturados em cada etapa.
- [ ] Comportamento real corresponde ao documentado nos requisitos.

---

## Passo 9 — Documentation Writer (`documentation-writer`) 📝

**Objetivo:** Aplicar o framework Diátaxis para revisar, unificar e expandir toda a documentação, garantindo qualidade textual e consistência cross-documento.

### Pré-condição:
- Passos 1 a 8 concluídos (toda a documentação base existe).

### Ações:

#### 9.1 — Classificar Cada Artefato no Framework Diátaxis

| Quadrante | Pergunta-chave | Artefatos `.specs/` + `docs/codebase/` |
|---|---|---|
| **Tutorial** (learning) | "Como eu começo?" | `engineering/code-analysis.md`, `docs/codebase/STRUCTURE.md` |
| **How-to** (problem) | "Como resolvo X?" | `engineering/api-guidelines.md`, `docs/codebase/CONVENTIONS.md`, `frontend/visual-inspection-report.md` |
| **Reference** (information) | "O que é X?" | `domain/domain.md`, `api/*.yaml`, `docs/codebase/STACK.md`, `docs/codebase/INTEGRATIONS.md`, `design/DESIGN.md`, `design/design-tokens.md`, `design/components.md`, `product/requirements.md`, `product/feature-roadmap.md`, `frontend/accessibility-audit.md` |
| **Explanation** (understanding) | "Por que X?" | `architecture/architecture.md`, `architecture/adrs/*.md`, `architecture/c4-*.md`, `docs/codebase/ARCHITECTURE.md`, `docs/codebase/CONCERNS.md`, `product/product.md`, `governance/confidence-report.md`, `frontend/audit-report.md`, `frontend/ux-critique.md` |

#### 9.2 — Revisar Cada Artefato

Para cada arquivo, verificar:
1. **Clareza**: Linguagem simples, sem jargão desnecessário. Acrônimos definidos na primeira ocorrência.
2. **Precisão**: Código, comandos e exemplos são executáveis e corretos.
3. **Consistência**: Terminologia alinhada com `domain/domain.md`. Estilo consistente entre arquivos.
4. **Completude**: O arquivo responde à pergunta-chave do seu quadrante Diátaxis.
5. **Cross-referências**: Links entre documentos funcionam e fazem sentido.

#### 9.3 — Criar Artefatos Ausentes (Product, Governance, Security)

**Product:**
- `product/product.md` — Explicação (Diátaxis): visão do produto, proposta de valor, personas.
- `product/requirements.md` — Referência (Diátaxis): RFs, RNFs, regras de negócio com EARS + MoSCoW.
- `product/feature-roadmap.md` — Referência (Diátaxis): features planejadas, concluídas, dívidas técnicas.

**Governance:**
- `governance/inventory.md` — Referência: tabela de todos os módulos, arquivos, cobertura de testes.
- `governance/confidence-report.md` — Explicação: score % de confiança em cada área.

**Security:**
- `security/SECURITY.md` — Referência: CSP, CORS, OWASP Top 10 frontend, headers de segurança, dependências.

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
- [ ] Nenhum `[TODO]` ou `[ASK USER]` sem resposta.

---

## Passo 10 — Sanity Check Final 🩺

1. **Links**: Verificar todos os links internos em `INDEX.md` e cross-referências entre documentos.
2. **Mermaid**: Confirmar que diagramas Mermaid em `architecture/c4-*.md` têm sintaxe válida.
3. **OpenAPI**: Rodar `npx @redocly/cli lint {SPECS_DIR}/api/{api-name}.yaml` e corrigir erros.
4. **Consistência de domínio**: Escolher 5 termos aleatórios do `domain/domain.md` e verificar que aparecem com o mesmo significado nos outros documentos.
5. **Screenshots**: Verificar que todos os screenshots nos reports de frontend são válidos e referenciados corretamente.
6. **Cobertura**: Comparar estrutura final com a estrutura canônica (Passo 0) e reportar:
   - Artefatos criados: N
   - Artefatos atualizados: N
   - Artefatos ainda ausentes: N (com justificativa)
7. **Git diff**: Executar `git diff --stat {SPECS_DIR}/ docs/codebase/` para auditar todas as mudanças.

---

## Passo 11 — Geração do Relatório de Execução

1. **Nome do arquivo**: `{AAAA-MM-DD-HHMMSS}-frontend-spec-mining.md`
   - Use o timestamp **real do momento da criação**: `date +%Y-%m-%d-%H%M%S`.
2. **Local**: `{SPECS_DIR}/skill-output/`
3. **Conteúdo**:

```markdown
# 📑 Relatório de Mineração de Especificações Frontend (FRONTEND-SPEC-MINING)

* **Data e Hora:** {AAAA-MM-DD HH:MM:SS} (GMT-3)
* **Skills orquestradas:** acquire-codebase-knowledge → domain-modeling → api-designer → architecture-designer → impeccable → web-design-reviewer → accessibility-compliance-accessibility-audit → webapp-testing → documentation-writer
* **Solução:** {SOLUTION_NAME} — {SOLUTION_DESCRIPTION}
* **Tipo de Aplicação:** {APP_TYPE}
* **Framework:** {LANGUAGE} | **Styling:** {STYLING}
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
| engineering/ | 2 | N | X% |
| product/ | 3 | N | X% |
| design/ | 3 | N | X% |
| frontend/ | 6 | N | X% |
| governance/ | 2 | N | X% |
| security/ | 1 | N | X% |

## ⚠️ Artefatos Ausentes (com justificativa)

- `...` — justificativa

## 🔍 Descobertas Notáveis

{Lista das 5-10 descobertas mais importantes do pipeline}

## 🔗 Links Gerados

- INDEX.md: `{SPECS_DIR}/INDEX.md`
- CHANGELOG.md: `{SPECS_DIR}/CHANGELOG.md`
- Codebase docs: `docs/codebase/`

---

🤖 *Documentação gerada pela orquestração de skills de IA: acquire-codebase-knowledge → domain-modeling → api-designer → architecture-designer → impeccable → web-design-reviewer → accessibility-compliance-accessibility-audit → webapp-testing → documentation-writer.*
```

---

# 🔄 Modos de Execução

O prompt suporta 7 modos, controlados pelo parâmetro `{SCOPE}`:

| Modo | `{SCOPE}` | Comportamento |
|---|---|---|
| **Completo** | `full` | Executa Passos 0→11 integralmente, criando/atualizando todos os artefatos |
| **Descoberta apenas** | `discovery-only` | Executa Passos 0→1 + Passo 11 (codebase knowledge + domain) |
| **Design apenas** | `design-only` | Executa Passos 0 + Passos 2→6 + Passo 11 (API + arquitetura + frontend design) |
| **Frontend apenas** | `frontend-only` | Executa Passos 0 + Passos 5→8 + Passo 11 (impeccable + web-design-reviewer + a11y + testing) |
| **Qualidade apenas** | `qa-only` | Executa Passos 0 + Passos 9→11 (documentation-writer review + sanity check) |
| **Delta** | `delta` | Executa apenas para os arquivos listados em `{TARGET_FILES}`, usando a skill dona correspondente + `documentation-writer` como revisor |
| **Revisão** | `review` | Executa apenas Passo 9 (documentation-writer) + Passo 10 + Passo 11, sem criar novos artefatos |

---

# ⚠️ Regras de Ouro

1. **Nunca sobrescrever sem auditoria**: antes de modificar um arquivo existente, fazer diff do conteúdo atual vs. novo e reportar no relatório final.
2. **Domínio primeiro, frontend depois**: os Passos 1-4 (fundamentos) são pré-requisito para os Passos 5-8 (frontend específico). Se o domínio não estiver estável, não faz sentido auditar UI/UX.
3. **Código é a verdade**: se houver conflito entre documentação existente e código, o código vence. Reportar a divergência e alinhar a documentação.
4. **App deve estar rodando para Passos 5, 6, 7, 8**: iniciar o dev server em `{DEV_SERVER_PORT}` antes desses passos. Se não for possível iniciar, pular passos que exigem runtime e registrar no relatório.
5. **Humano decide ambiguidades**: o `domain-modeling` pode identificar termos conflitantes, mas apenas o humano pode resolvê-los.
6. **Mermaid é obrigatório**: todo diagrama arquitetural deve ser renderizável como Mermaid (não usar imagens estáticas).
7. **OpenAPI validado**: a spec da API deve passar na validação `@redocly/cli lint` sem erros.
8. **Screenshots são evidência**: todo report de frontend (`audit-report.md`, `visual-inspection-report.md`, `functional-verification.md`) deve incluir screenshots como evidência.
9. **Diátaxis como camada final**: o `documentation-writer` revisa todos os artefatos, mas não redefine decisões técnicas dos passos anteriores — apenas melhora clareza, consistência e organização.
10. **INDEX.md é sagrado**: sempre atualizar o índice central ao final de cada execução, incluindo links para `docs/codebase/`.
11. **Marcar incertezas explicitamente**: usar `[TODO]` para fatos não verificáveis e `[ASK USER]` para decisões que requerem intenção humana.

---

# 📚 Referências

- **Diátaxis Framework**: [https://diataxis.fr/](https://diataxis.fr/)
- **C4 Model**: [https://c4model.com/](https://c4model.com/)
- **OpenAPI 3.1**: [https://spec.openapis.org/oas/v3.1.0](https://spec.openapis.org/oas/v3.1.0)
- **Mermaid.js**: [https://mermaid.js.org/](https://mermaid.js.org/)
- **ADR Template**: [https://adr.github.io/madr/](https://adr.github.io/madr/)
- **WCAG 2.2**: [https://www.w3.org/TR/WCAG22/](https://www.w3.org/TR/WCAG22/)
- **Nielsen's Heuristics**: [https://www.nngroup.com/articles/ten-usability-heuristics/](https://www.nngroup.com/articles/ten-usability-heuristics/)
- **EARS Format**: `references/ears-format.md`
- **OKLCH Color**: [https://oklch.com/](https://oklch.com/)
- **Playwright MCP**: [https://github.com/microsoft/playwright-mcp](https://github.com/microsoft/playwright-mcp)
- **Arc42**: [https://docs.arc42.org/](https://docs.arc42.org/)
