# Contexto:
  - Este prompt orquestra **8 skills especializadas** para minerar, criar e atualizar a documentação técnica completa da solução no diretório `.specs/`.
  - As skills são invocadas em sequência lógica: primeiro faz-se o scan automatizado da codebase e extraem-se especificações do código existente, depois modela-se o domínio, desenha-se a API, documenta-se o código-fonte, projeta-se a arquitetura, audita-se a segurança, e por fim aplica-se o framework Diátaxis para unificar qualidade e consistência.
  - A estrutura de pastas alvo (descrita abaixo) é a referência canônica — o prompt pode enriquecê-la com novos artefatos quando o contexto da solução exigir.
  - A documentação existente em `.specs/` deve ser preservada e incrementada, nunca sobrescrita sem auditoria de delta.
  - Use o README.md da raiz da solução para instruções de build, execução e testes.

---

# 🎯 Skills Orquestradas

| Ordem | Skill | Responsabilidade | Pasta(s) Principal(is) |
|---|---|---|---|
| 0ª | `acquire-codebase-knowledge` | Scan automatizado da codebase: stack, estrutura, arquitetura, convenções, integrações, testes, concerns | `docs/codebase/` (insumo para todos os passos seguintes) |
| 0.5ª | `spec-miner` | Engenharia reversa de especificações: extrai requisitos em formato EARS, mapeia dependências, identifica lógica de negócio implícita | `specs/` (especificação reversa) |
| 1ª | `domain-modeling` | Glossário de domínio, ubiquitous language, ADRs de domínio | `domain/` |
| 2ª | `api-designer` | Contrato OpenAPI 3.1, modelagem de recursos, versionamento, segurança de API | `api/` |
| 2.5ª | `code-documenter` | Documentação inline (docstrings, JSDoc), análise de fluxo de código, validação de exemplos | `engineering/` |
| 3ª | `architecture-designer` | C4 (Context→Containers→Components→Code), ADRs, ERD, dicionário de dados, integrações | `architecture/` |
| 3.5ª | `security-audit` | Auditoria de segurança: recon, hunt, validate, report. Threat model, vulnerabilidades, OWASP Top 10 | `security/` |
| 4ª | `documentation-writer` | Diátaxis (Tutorials/How-to/Reference/Explanation), qualidade textual, consistência cross-documento | **Todas as pastas** (revisão final) |

---

# ⚙️ Parâmetros de Entrada (preencher antes de executar)

> **Instrução:** No momento de invocar este prompt, o humano deve informar os valores abaixo. Se algum arquivo não existir, marcar como `N/D` (não disponível) e criar no passo apropriado.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_ROOT}` | Caminho absoluto da raiz da solução | `/home/user/work/ms-billing-engine-tax-rates` |
| `{SPECS_DIR}` | Caminho relativo da pasta de especificações | `.specs` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-billing-engine-tax-rates` |
| `{SOLUTION_DESCRIPTION}` | Descrição curta do propósito da solução | `Microsserviço de cálculo de tributos sobre faturamento (IPI, ICMS, PIS, COFINS, ISS, CBS, IBS, IS)` |
| `{LANGUAGE}` | Linguagem/framework principal | `go/fiber` |
| `{DATABASE}` | Banco de dados principal | `PostgreSQL 16` |
| `{PROJECT_ID}` | Identificador do projeto corporativo (se aplicável) | `PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO` |
| `{SCOPE}` | Escopo da documentação a ser gerada | `full`, `scan-only`, `domain-only`, `api-only`, `code-analysis-only`, `architecture-only`, `security-only`, `delta`, `review` |
| `{TARGET_FILES}` | Lista específica de arquivos a criar/atualizar (se `{SCOPE}=delta`) | `domain/domain.md, api/tax-rates-api.yaml` |

---

# 📁 Estrutura de Pastas Alvo (Referência Canônica)

> Esta estrutura é o template de saída. O prompt deve tratar pastas/arquivos ausentes como **oportunidade de criação** e pastas/arquivos existentes como **base para atualização incremental**. Novos artefatos podem ser sugeridos quando o contexto da solução exigir (ex: `security/SECURITY.md`, `architecture/c4-containers.md`).

```
{SPECS_DIR}/
├── INDEX.md                                        ← Índice central de toda a documentação
├── CHANGELOG.md                                    ← Histórico de mudanças da documentação
│
├── domain/
│   └── domain.md                                   ← Glossário de domínio + ubiquitous language
│
├── api/
│   └── tax-rates-api.yaml                          ← Contrato OpenAPI 3.1 completo
│
├── architecture/
│   ├── architecture.md                             ← Visão arquitetural geral da solução
│   ├── c4-context.md                               ← Diagrama C4 — Nível 1 (Contexto)
│   ├── c4-containers.md                            ← Diagrama C4 — Nível 2 (Containers)
│   ├── c4-components.md                            ← Diagrama C4 — Nível 3 (Componentes)
│   ├── c4-code-class.md                            ← Diagrama C4 — Nível 4 (Código/Classes)
│   ├── erd.md                                      ← Diagrama Entidade-Relacionamento
│   ├── data-dictionary.md                          ← Dicionário de dados (tabelas, colunas, tipos, regras)
│   ├── integrations.md                             ← Integrações externas, dependências, libs
│   └── adrs/                                       ← ADRs específicos da solução
│       ├── INDEX.md                                ← Índice cronológico dos ADRs
│       ├── adr-001.md                              ← ADR #1
│       ├── adr-002.md                              ← ADR #2
│       └── ...
│
├── engineering/
│   ├── api-guidelines.md                           ← Padrões de API, erros, observabilidade
│   └── code-analysis.md                            ← Análise técnica de fluxo de código
│
├── product/
│   ├── product.md                                  ← Descrição do produto e visão
│   ├── requirements.md                             ← Especificação de requisitos (RF/RNF + MoSCoW)
│   └── feature-roadmap.md                          ← Roadmap de features + dívidas técnicas
│
├── design/                                         ← (Opcional — se a solução tiver UI)
│   ├── design.md                                   ← Visão geral de design
│   ├── design-system.md                            ← Design system / componentes visuais
│   └── components.md                               ← Catálogo de componentes de UI
│
├── governance/
│   ├── inventory.md                                ← Inventário do projeto + cobertura de testes
│   └── confidence-report.md                        ← Relatório de confiança da documentação (%)
│
├── security/                                       ← (Opcional — se existirem regras específicas)
│   └── SECURITY.md                                 ← Definições de segurança da solução
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
    └── {AAAA-MM-DD-HHMISS}-{nome-da-task}.md       ← Registros de execução de skills
```

---

# 🧭 Mapeamento Skill → Artefato

> Tabela de responsabilidade: qual skill é **dona** (primary owner) de cada artefato. O `documentation-writer` atua como revisor de **todos** os artefatos ao final.

| Artefato | Skill Dona | Ação Típica |
|---|---|---|
| `docs/codebase/STACK.md` | `acquire-codebase-knowledge` | Scan automatizado: linguagens, frameworks, versões |
| `docs/codebase/STRUCTURE.md` | `acquire-codebase-knowledge` | Estrutura de diretórios, naming conventions |
| `docs/codebase/ARCHITECTURE.md` | `acquire-codebase-knowledge` | Padrão arquitetural detectado, pontos de entrada |
| `docs/codebase/CONVENTIONS.md` | `acquire-codebase-knowledge` | Convenções de código, padrões do projeto |
| `docs/codebase/INTEGRATIONS.md` | `acquire-codebase-knowledge` | Dependências externas, serviços, APIs |
| `docs/codebase/TESTING.md` | `acquire-codebase-knowledge` | Frameworks de teste, cobertura, padrões |
| `docs/codebase/CONCERNS.md` | `acquire-codebase-knowledge` | Tech debt, bugs, riscos de segurança detectados |
| `specs/{project}_reverse_spec.md` | `spec-miner` | Extrair especificação reversa em EARS do código |
| `domain/domain.md` | `domain-modeling` | Criar/atualizar glossário, resolver ambiguidades de linguagem |
| `api/tax-rates-api.yaml` | `api-designer` | Gerar/atualizar contrato OpenAPI 3.1 completo |
| `architecture/architecture.md` | `architecture-designer` | Escrever visão geral, padrões, trade-offs |
| `architecture/c4-*.md` | `architecture-designer` | Gerar diagramas C4 nos 4 níveis (Mermaid) |
| `architecture/erd.md` | `architecture-designer` | Modelar entidades e relacionamentos |
| `architecture/data-dictionary.md` | `architecture-designer` | Documentar tabelas, colunas, tipos, regras |
| `architecture/integrations.md` | `architecture-designer` | Mapear dependências externas e libs |
| `architecture/adrs/*.md` | `architecture-designer` | Registrar decisões arquiteturais (ADR) |
| `engineering/api-guidelines.md` | `api-designer` → `documentation-writer` | Padrões de API REST, erros, segurança |
| `engineering/code-analysis.md` | `code-documenter` | Análise de fluxo de handlers/services, docstrings, cobertura |
| `product/product.md` | `documentation-writer` | Descrição do produto (Explanation — Diátaxis) |
| `product/requirements.md` | `documentation-writer` | Requisitos (Reference — Diátaxis) |
| `product/feature-roadmap.md` | `documentation-writer` | Roadmap (Reference — Diátaxis) |
| `design/*.md` | `documentation-writer` | Documentação de design system (se aplicável) |
| `governance/inventory.md` | `documentation-writer` | Inventário e cobertura |
| `governance/confidence-report.md` | `documentation-writer` | Score de confiança da documentação |
| `security/SECURITY.md` | `security-audit` | Threat model, vulnerabilidades, OWASP, checklists |
| `security/threat-model.md` | `security-audit` | Modelo de ameaças, superfície de ataque, severidade |
| `INDEX.md` | `documentation-writer` | Índice central (atualizado a cada ciclo) |
| `CHANGELOG.md` | `documentation-writer` | Registro de mudanças da doc |
| `business-projects/{ID}/*.md` | `documentation-writer` | PRD, SPECS, TASKS, TEST_PLAN, ARCHITECTURE |
| `features/{NAME}/*.md` | `documentation-writer` | SPECS, ARCHITECTURE, TASKS, TEST_PLAN |

---

# 🔁 Protocolo de Execução (Passo a Passo)

## Passo 0 — Codebase Scan (`acquire-codebase-knowledge`) 🔍

**Objetivo:** Fazer scan automatizado da codebase para produzir 7 documentos factuais que servirão de insumo para todos os passos seguintes. Este passo substitui a leitura manual do cenário — os documentos gerados são a fonte canônica de fatos sobre o código.

### Ações:

1. **Executar o script de scan** a partir da raiz do projeto:
   ```bash
   python3 "$SKILL_ROOT/scripts/scan.py" --output {SOLUTION_ROOT}/docs/codebase/.codebase-scan.txt
   ```
   O script detecta automaticamente: linguagens (25+), frameworks, plataformas CI/CD (10+), containers, orquestração, métricas de código, configs de segurança e marcadores de teste de performance.

2. **Populatar os 7 documentos** em `docs/codebase/` com base nas evidências coletadas:
   - `STACK.md` — Linguagens, frameworks, versões, banco de dados, runtime.
   - `STRUCTURE.md` — Estrutura de diretórios, naming conventions, organização de pacotes.
   - `ARCHITECTURE.md` — Padrão arquitetural detectado, pontos de entrada, camadas.
   - `CONVENTIONS.md` — Convenções de código, padrões de commit, estilo.
   - `INTEGRATIONS.md` — Dependências externas, APIs chamadas, message brokers, serviços cloud.
   - `TESTING.md` — Frameworks de teste, cobertura, padrões de teste.
   - `CONCERNS.md` — Tech debt, bugs conhecidos, riscos de segurança, code smells.

3. **Validar** que todos os 7 documentos existem e cada afirmação é rastreável a um arquivo fonte, config ou saída de terminal.

4. **Marcar incertezas**: use `[TODO]` para fatos não verificáveis e `[ASK USER]` para decisões que dependem de intenção humana.

5. **Ler** `{SPECS_DIR}/INDEX.md` e `{SPECS_DIR}/CHANGELOG.md` (se existirem) para entender o estado atual da documentação.

6. **Gerar delta inicial**: comparar estrutura atual de `{SPECS_DIR}/` com a estrutura canônica e listar arquivos existentes (📄), ausentes (❌), e desatualizados (🔄).

7. **Confirmar com o humano** o escopo antes de prosseguir (especialmente se `{SCOPE}=full`).

### Validação:
- [ ] Script de scan executado com sucesso.
- [ ] 7 documentos em `docs/codebase/` populados com evidências verificáveis.
- [ ] Delta inicial de `{SPECS_DIR}/` gerado.
- [ ] Escopo confirmado com o humano.

---

## Passo 0.5 — Engenharia Reversa de Especificações (`spec-miner`) ⛏️

**Objetivo:** Extrair especificações do código existente usando engenharia reversa, documentando requisitos observados em formato EARS e identificando lógica de negócio implícita. Este passo complementa os 7 documentos do Passo 0 com uma visão comportamental do sistema.

### Pré-condição:
- Passo 0 concluído (os 7 docs de `docs/codebase/` fornecem o contexto inicial para o `spec-miner`).

### Ações:

1. **Definir escopo da análise**: com base no delta do Passo 0 e no `{SCOPE}`, identificar quais partes do sistema analisar (sistema completo ou features específicas).

2. **Explorar a estrutura** usando Glob, Grep, Read:
   - Encontrar pontos de entrada (handlers, rotas, controllers, CLI commands).
   - Mapear dependências entre módulos/pacotes.
   - Localizar marcadores de tech debt (`TODO`, `FIXME`, `HACK`, `XXX`).
   - Descobrir uso de configurações e variáveis de ambiente.
   - Mapear rotas de API e message handlers.

3. **Rastrear fluxos de dados** e caminhos de requisição para entender o comportamento do sistema.

4. **Documentar requisitos observados** em formato EARS (Easy Approach to Requirements Syntax):
   | Tipo | Padrão | Exemplo |
   |------|--------|---------|
   | Ubíquo | The `<system>` shall `<action>`. | The API shall return JSON responses. |
   | Event-driven | When `<trigger>`, the `<system>` shall `<action>`. | When a request lacks an auth token, the system shall return HTTP 401. |
   | State-driven | While `<state>`, the `<system>` shall `<action>`. | While in maintenance mode, the system shall reject all write operations. |
   | Optional | Where `<feature>` is supported, the `<system>` shall `<action>`. | Where caching is enabled, the system shall store responses for 60s. |

5. **Identificar lógica de negócio implícita**: regras implementadas em código mas não documentadas, validações, cálculos, fluxos de decisão.

6. **Gerar** `specs/{solution_name}_reverse_spec.md` contendo:
   - Stack tecnológico e arquitetura (observados, não inferidos).
   - Estrutura de módulos/diretórios.
   - Requisitos observados (formato EARS).
   - Observações não-funcionais (performance, segurança, resiliência).
   - Critérios de aceitação inferidos.
   - Incertezas e perguntas.
   - Recomendações.

7. **Distinguir fatos de inferências**: toda observação deve ser ancorada em evidência de código (caminho do arquivo + linha). Inferências devem ser marcadas como `[INFERIDO]`.

### Validação:
- [ ] Todos os pontos de entrada mapeados.
- [ ] Requisitos EARS ancorados em evidências de código.
- [ ] Lógica de negócio implícita identificada e documentada.
- [ ] Incertezas listadas para resolução humana.

---

## Passo 1 — Domain Modeling (`domain-modeling`) 🏛️

**Objetivo:** Garantir que o glossário de domínio e a linguagem ubíqua estejam corretos e completos **antes** de documentar API e arquitetura.

### Pré-condição:
- Passos 0 e 0.5 concluídos (os documentos de scan e a especificação reversa fornecem o vocabulário técnico e de negócio inicial).

### Ações:
1. **Ler** `domain/domain.md` (se existir) e extrair todos os termos definidos.
2. **Cross-referenciar** com o código-fonte em `{SOLUTION_ROOT}`:
   - Nomes de structs, enums, constantes, tipos.
   - Termos usados em nomes de endpoints, parâmetros, e mensagens de erro.
   - Comentários e docstrings que definem conceitos de negócio.
3. **Identificar ambiguidades**: termos iguais com significados diferentes, termos diferentes para o mesmo conceito, termos ausentes no glossário mas usados no código.
4. **Resolver com o humano**: para cada ambiguidade, perguntar "no código usa-se X, mas o glossário define Y — qual é o correto?"
5. **Atualizar** `domain/domain.md`:
   - Adicionar termos ausentes.
   - Corrigir definições inconsistentes.
   - Remover termos obsoletos (marcar como `[DEPRECATED]` antes de remover).
   - Manter agrupamento por subdomínio (ex: tributos, CSTs, CFOPs, NCMs, regimes).
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
| Termo | Definição | Sinônimos | Código (struct/enum) |
|---|---|---|---|
| ... | ... | ... | ... |

### 2.2. {Subdomínio 2}
...

## 3. Relações entre Conceitos
{Diagrama ou descrição de como os principais conceitos se relacionam.}

## 4. Regras de Negócio Fundamentais
{Lista das regras mais estáveis que todo o time deve conhecer.}
```

### Validação:
- [ ] Todo termo do glossário aparece no código com o mesmo significado.
- [ ] Nenhum termo do código (struct, enum, constante pública) está ausente do glossário.
- [ ] Ambiguidades resolvidas e registradas.

---

## Passo 2 — API Design (`api-designer`) 🔌

**Objetivo:** Produzir/atualizar o contrato OpenAPI 3.1 completo, alinhado com o domínio definido no Passo 1.

### Pré-condição:
- `domain/domain.md` deve estar atualizado (os nomes de recursos e parâmetros da API derivam dos termos de domínio).

### Ações:
1. **Ler** o código-fonte em `{SOLUTION_ROOT}` para descobrir endpoints existentes:
   - Handlers, rotas, middlewares.
   - Estruturas de request/response (schemas, DTOs).
   - Códigos de erro e formato de resposta.
   - Autenticação e autorização (JWT, RBAC, scopes).
2. **Ler** `api/tax-rates-api.yaml` (se existir) para baseline.
3. **Comparar** endpoints documentados vs. endpoints implementados:
   - Rotas documentadas mas não implementadas → marcar como `x-unimplemented: true`.
   - Rotas implementadas mas não documentadas → adicionar à spec.
   - Schemas divergentes → alinhar com a implementação real.
4. **Gerar/atualizar** `api/tax-rates-api.yaml` com:
   - `openapi: 3.1.0`
   - `info` (title, version, description)
   - `servers` (local, staging, production)
   - `paths` (todos os endpoints com métodos HTTP, parâmetros, bodies, responses)
   - `components/schemas` (todos os DTOs com exemplos)
   - `components/securitySchemes` (JWT Bearer, API Key, etc.)
   - `components/responses` (erros padrão: 400, 401, 403, 404, 422, 500)
   - `tags` (agrupamento por recurso)
5. **Validar** a spec gerada:
   ```bash
   npx @redocly/cli lint {SPECS_DIR}/api/tax-rates-api.yaml
   ```
6. **Atualizar** `engineering/api-guidelines.md` se novos padrões forem identificados.

### Validação:
- [ ] Spec OpenAPI passa na validação `@redocly/cli lint` sem erros.
- [ ] Todos os endpoints do código estão documentados.
- [ ] Schemas de request/response batem com os DTOs do código.
- [ ] Exemplos de request/response incluídos para cada endpoint.
- [ ] Autenticação documentada (securitySchemes + security por endpoint).

---

## Passo 2.5 — Documentação de Código (`code-documenter`) 📄

**Objetivo:** Gerar e validar documentação inline de código (docstrings, JSDoc), analisar fluxos de código, e produzir o artefato `engineering/code-analysis.md` com cobertura de handlers, services e repositórios.

### Pré-condição:
- Passos 1 e 2 concluídos (o domínio e a API fornecem o vocabulário para documentar o código).

### Ações:

1. **Detectar** linguagens e frameworks no código-fonte:
   - Identificar arquivos não documentados (funções/métodos sem docstrings).
   - Identificar o formato de documentação apropriado (Google-style, NumPy, JSDoc, Javadoc).

2. **Documentar código inline**:
   - Adicionar/atualizar docstrings em funções públicas, classes e métodos.
   - Documentar parâmetros, valores de retorno, exceções e efeitos colaterais.
   - Usar terminologia alinhada com `domain/domain.md`.

3. **Analisar fluxos de código** para `engineering/code-analysis.md`:
   - Mapear fluxos principais: entry point → handler → service → repository → response.
   - Documentar middlewares, interceptors, filters.
   - Identificar padrões de erro e estratégias de retry.
   - Rastrear dependências entre camadas.

4. **Validar exemplos de código**:
   - Python: `python -m doctest file.py` ou `pytest --doctest-modules`
   - TypeScript: `tsc --noEmit`
   - OpenAPI: `npx @redocly/cli lint openapi.yaml`
   - Corrigir exemplos que falham na validação.

5. **Gerar relatório de cobertura**: percentual de código documentado, funções sem docstrings, qualidade da documentação.

### Validação:
- [ ] Docstrings seguem formato consistente (Google, NumPy, ou JSDoc conforme linguagem).
- [ ] Exemplos de código compilam/executam sem erro.
- [ ] `engineering/code-analysis.md` cobre os fluxos principais.
- [ ] Terminologia alinhada com `domain/domain.md`.
- [ ] Relatório de cobertura gerado.

---

## Passo 3 — Architecture Design (`architecture-designer`) 🏗️

**Objetivo:** Produzir/atualizar a documentação arquitetural completa nos 4 níveis C4, ERD, dicionário de dados, integrações e ADRs.

### Pré-condição:
- `domain/domain.md`, `api/tax-rates-api.yaml`, e `specs/{solution_name}_reverse_spec.md` atualizados (a arquitetura referencia todos eles).

### Ações:

#### 3.1 — C4 Model (4 níveis)

| Nível | Arquivo | Conteúdo |
|---|---|---|
| **Context** | `c4-context.md` | Sistema no ecossistema: atores externos (usuários, sistemas), fluxos de entrada/saída |
| **Containers** | `c4-containers.md` | Containers de deploy: API, banco, cache, message broker, serviços externos |
| **Components** | `c4-components.md` | Componentes internos de cada container: handlers, services, repositories, calculators |
| **Code/Class** | `c4-code-class.md` | Diagrama de classes/pacotes principais: interfaces, implementações, dependências |

Cada arquivo deve conter:
1. Diagrama **Mermaid** (obrigatório).
2. Descrição textual de cada elemento do diagrama.
3. Fluxos de dados principais (sequência de chamadas).

Template base para cada nível C4:
```markdown
---
title: "C4 — {Nível} — {SOLUTION_NAME}"
level: "{Context|Containers|Components|Code}"
---

# C4 — {Nível}: {Título}

## Diagrama

```mermaid
C4Context / C4Container / C4Component / classDiagram
  ... 
```

## Elementos

| Nome | Tipo | Responsabilidade | Tecnologia |
|---|---|---|---|
| ... | ... | ... | ... |

## Fluxos Principais

### Fluxo: {Nome}
1. {Passo 1}
2. {Passo 2}
...
```

#### 3.2 — ERD + Dicionário de Dados

1. **Analisar** `data/init.sql`, arquivos de migration, ou entities do ORM.
2. **Gerar/atualizar** `architecture/erd.md`:
   - Diagrama Mermaid `erDiagram` com todas as tabelas, colunas, PKs, FKs, índices.
   - Relacionamentos com cardinalidade.
3. **Gerar/atualizar** `architecture/data-dictionary.md`:
   - Para cada tabela: propósito, colunas (nome, tipo, nullable, default, descrição), índices, triggers.
   - Regras de negócio implementadas via constraints/triggers.

#### 3.3 — Integrações

1. **Mapear** todas as dependências externas no código:
   - APIs REST chamadas (HTTP clients).
   - Bancos de dados (PostgreSQL, Redis).
   - Message brokers (Kafka, RabbitMQ).
   - Serviços cloud (AWS, Azure, GCP).
   - Libs locais compartilhadas (`libs/go-native/...`).
2. **Gerar/atualizar** `architecture/integrations.md`:
   - Para cada integração: endpoint/endereço, propósito, contrato (request/response), timeout, retry policy, circuit breaker.
   - Diagrama de dependências entre serviços.

#### 3.4 — ADRs

1. **Identificar** decisões arquiteturais merecedoras de ADR:
   - Escolha de framework/biblioteca.
   - Padrão de arquitetura (DDD, Clean, Hexagonal).
   - Estratégia de versionamento de API.
   - Decisões de banco de dados (particionamento, índices).
2. **Criar/atualizar** ADRs em `architecture/adrs/` usando template:
```markdown
---
title: "ADR-{NNNN}: {Título da Decisão}"
status: "Proposed | Accepted | Deprecated | Superseded"
date: "{AAAA-MM-DD}"
deciders: "{Stakeholders}"
tags: ["architecture", "{domínio}"]
---

# ADR-{NNNN}: {Título da Decisão}

## Contexto
{Por que esta decisão é necessária? Quais as forças em jogo?}

## Decisão
{O que foi decidido. Seja específico.}

## Alternativas Consideradas
| Alternativa | Prós | Contras | Motivo da Rejeição |
|---|---|---|---|
| ... | ... | ... | ... |

## Consequências
### Positivas
- ...

### Negativas
- ...

## Referências
- ...
```
3. **Atualizar** `architecture/adrs/INDEX.md` com a lista cronológica.

#### 3.5 — Visão Geral de Arquitetura

1. **Gerar/atualizar** `architecture/architecture.md` consolidando:
   - Resumo executivo (1 parágrafo).
   - Padrão arquitetural adotado (com justificativa).
   - Estrutura de diretórios do projeto (tree com comentários).
   - Stack tecnológica completa (tabela).
   - Princípios de design (SOLID, DDD, etc.).
   - Cross-cutting concerns (logging, tracing, metrics, auth).

### Validação:
- [ ] Todos os diagramas Mermaid renderizam corretamente.
- [ ] ERD cobre todas as tabelas do `data/init.sql`.
- [ ] Integrações cobrem todas as dependências externas.
- [ ] ADRs seguem template e têm status definido.
- [ ] Consistência entre `c4-*.md`, `erd.md`, e `integrations.md`.

---

## Passo 3.5 — Auditoria de Segurança (`security-audit`) 🛡️

**Objetivo:** Auditar a segurança da solução, identificando vulnerabilidades exploráveis com impacto real, produzindo threat model e preenchendo a pasta `security/`.

### Pré-condição:
- Passo 3 concluído (a documentação arquitetural — C4, integrações, ERD — fornece o mapa de superfície de ataque).

### Ações:

#### 3.5.1 — Setup

1. **Estabelecer diretório de saída**: `{SPECS_DIR}/security/` (criar se não existir).
2. **Verificar execuções anteriores**: se `security/findings.json` existir, ler para evitar duplicação de achados e focar em novas áreas.
3. **Artefatos a produzir**:
   - `security/SECURITY.md` — Relatório de segurança legível para humanos.
   - `security/threat-model.md` — Modelo de ameaças, superfície de ataque, severidade.
   - `security/findings.json` — Achados estruturados (machine-readable).

#### 3.5.2 — Fase 1: Recon (Mapeamento)

1. **Ler a documentação arquitetural** gerada nos passos anteriores:
   - `architecture/c4-context.md` — atores externos, fluxos de entrada/saída.
   - `architecture/c4-containers.md` — containers de deploy, superfícies expostas.
   - `architecture/integrations.md` — dependências externas, APIs chamadas.
   - `docs/codebase/CONCERNS.md` — tech debt e riscos já identificados.
2. **Identificar**:
   - Superfície de ataque (APIs públicas, endpoints, ports abertos).
   - Fluxos de autenticação e autorização.
   - Pontos de entrada de dados (forms, file uploads, webhooks).
   - Dependências externas com vulnerabilidades conhecidas.
3. **Escrever** `security/architecture.md` (resumo da superfície de ataque para alimentar as fases seguintes).

#### 3.5.3 — Fase 2: Hunt (Caça a Vulnerabilidades)

Varrer o código-fonte usando classes de ataque:

| Classe de Ataque | Foco | Referência |
|---|---|---|
| **Injection** | SQL, NoSQL, Command, LDAP, XPath | OWASP A03:2021 |
| **Auth & Session** | JWT, OAuth2, sessões, password reset | OWASP A07:2021 |
| **Access Control** | RBAC, ABAC, IDOR, privilege escalation | OWASP A01:2021 |
| **Data Exposure** | Logs, error messages, debug endpoints, backups | OWASP A02:2021 |
| **Input Validation** | XSS, SSTI, open redirect, file upload | OWASP A03:2021 |
| **SSRF & Network** | Server-side request forgery, internal service exposure | OWASP A10:2021 |
| **Business Logic** | Race conditions, workflow bypass, negative amounts | OWASP A04:2021 |
| **Dependencies** | CVEs em bibliotecas, licenças, versões desatualizadas | OWASP A06:2021 |

#### 3.5.4 — Fase 3: Validate (Verificação)

Para cada achado:
1. **Construir cenário de ataque concreto**: quem é o atacante, o que faz, o que obtém.
2. **Confirmar com evidência de código** (arquivo + linha).
3. **Classificar severidade**:
   - **CRITICAL**: RCE não-autenticado, dump completo de BD, account takeover sem credenciais.
   - **HIGH**: RCE autenticado, SQLi com exfiltração, XSS stored para todos usuários, auth bypass.
   - **MEDIUM**: XSS condicional, CSRF com mudança de estado, information disclosure de secrets.
   - **LOW**: Information disclosure de dados não-secretos, DoS com esforço sustentado.
   - **INFORMATIONAL**: Observação sem exploit standalone.

#### 3.5.5 — Fase 4: Report (Documentação)

1. **Escrever** `security/SECURITY.md`:
   - Resumo executivo (1 parágrafo).
   - Threat model: atores, superfície de ataque, cenários.
   - Achados priorizados por severidade.
   - Recomendações de remediação.
   - Checklist OWASP Top 10 (marcar ✅/⚠️/❌ para cada item).
   - Hardening notes (defense-in-depth, não vulnerabilidades).

2. **Escrever** `security/threat-model.md`:
   - Diagrama de fluxo de dados com trust boundaries.
   - Atores e seus níveis de acesso.
   - Cenários de ataque por ator.
   - Controles existentes e lacunas.

3. **Gerar** `security/findings.json` com achados estruturados (formato JSON Schema para consumo automatizado).

### Validação:
- [ ] `security/SECURITY.md` cobre OWASP Top 10.
- [ ] `security/threat-model.md` inclui diagrama de trust boundaries.
- [ ] Todo achado tem cenário de ataque concreto e evidência de código.
- [ ] Severidade segue critérios impacto × likelihood.
- [ ] `security/findings.json` é JSON válido.

---

## Passo 4 — Documentation Writer (`documentation-writer`) 📝

**Objetivo:** Aplicar o framework Diátaxis para revisar, unificar e expandir toda a documentação, garantindo qualidade textual e consistência cross-documento.

### Pré-condição:
- Passos 0, 0.5, 1, 2, 2.5, 3, e 3.5 concluídos (toda a documentação técnica base existe).

### Ações:

#### 4.1 — Classificar Cada Artefato no Framework Diátaxis

| Quadrante | Pergunta-chave | Artefatos `.specs/` |
|---|---|---|
| **Tutorial** (learning) | "Como eu começo?" | `engineering/code-analysis.md` (fluxo guiado), `docs/codebase/STRUCTURE.md` |
| **How-to** (problem) | "Como resolvo X?" | `engineering/api-guidelines.md`, `architecture/integrations.md`, `security/SECURITY.md` (remediação) |
| **Reference** (information) | "O que é X?" | `domain/domain.md`, `api/*.yaml`, `architecture/erd.md`, `architecture/data-dictionary.md`, `product/requirements.md`, `product/feature-roadmap.md`, `docs/codebase/STACK.md`, `docs/codebase/TESTING.md`, `security/findings.json` |
| **Explanation** (understanding) | "Por que X?" | `architecture/architecture.md`, `architecture/adrs/*.md`, `architecture/c4-*.md`, `product/product.md`, `governance/confidence-report.md`, `specs/{solution_name}_reverse_spec.md`, `security/threat-model.md`, `docs/codebase/CONCERNS.md` |

#### 4.2 — Revisar Cada Artefato

Para cada arquivo, verificar:
1. **Clareza**: Linguagem simples, sem jargão desnecessário. Acrônimos definidos na primeira ocorrência.
2. **Precisão**: Código, comandos e exemplos são executáveis e corretos.
3. **Consistência**: Terminologia alinhada com `domain/domain.md`. Estilo consistente entre arquivos.
4. **Completude**: O arquivo responde à pergunta-chave do seu quadrante Diátaxis.
5. **Cross-referências**: Links entre documentos funcionam e fazem sentido.

#### 4.3 — Criar Artefatos Ausentes (Product, Governance, Design, Security)

**Product:**
- `product/product.md` — Explicação (Diátaxis): visão do produto, proposta de valor, personas.
- `product/requirements.md` — Referência (Diátaxis): RFs, RNFs, regras de negócio com MoSCoW.
- `product/feature-roadmap.md` — Referência (Diátaxis): features planejadas, concluídas, dívidas técnicas.

**Governance:**
- `governance/inventory.md` — Referência: tabela de todos os módulos, arquivos, cobertura de testes.
- `governance/confidence-report.md` — Explicação: score % de confiança em cada área da documentação.

**Design** (se aplicável):
- `design/design.md` — Explicação: visão geral do design system.
- `design/design-system.md` — Referência: tokens, cores, tipografia, espaçamento.
- `design/components.md` — Referência: catálogo de componentes.

**Security** (gerado pelo `security-audit` — Passo 3.5):
- `security/SECURITY.md` — Referência: threat model, práticas de segurança, OWASP, checklists.
- `security/threat-model.md` — Explicação: modelo de ameaças, trust boundaries.
- `security/findings.json` — Referência: achados estruturados (machine-readable).

**docs/codebase/** (gerado pelo `acquire-codebase-knowledge` — Passo 0):
- `docs/codebase/STACK.md` — Referência: stack tecnológico.
- `docs/codebase/STRUCTURE.md` — Tutorial: estrutura do projeto.
- `docs/codebase/ARCHITECTURE.md` — Explicação: arquitetura detectada.
- `docs/codebase/CONVENTIONS.md` — Referência: convenções de código.
- `docs/codebase/INTEGRATIONS.md` — How-to: dependências externas.
- `docs/codebase/TESTING.md` — Referência: frameworks e padrões de teste.
- `docs/codebase/CONCERNS.md` — Explicação: tech debt e riscos.

**specs/** (gerado pelo `spec-miner` — Passo 0.5):
- `specs/{solution_name}_reverse_spec.md` — Explicação: especificação reversa em EARS.

#### 4.4 — Atualizar INDEX.md e CHANGELOG.md

1. **Atualizar** `INDEX.md`:
   - Listar todos os arquivos existentes com links e descrição de 1 linha.
   - Agrupar por pasta/seção.
   - Incluir data da última atualização geral.
   - Incluir score de confiança (de `governance/confidence-report.md`).

2. **Atualizar** `CHANGELOG.md`:
   - Registrar data, arquivos criados/atualizados, e resumo das mudanças.
   - Seguir o formato:
```markdown
## {AAAA-MM-DD} — {Resumo curto}
### Criados
- `caminho/arquivo.md` — descrição
### Atualizados
- `caminho/arquivo.md` — o que mudou
### Removidos
- `caminho/arquivo.md` — motivo
```

### Validação:
- [ ] Todo artefato está classificado num quadrante Diátaxis.
- [ ] Links cross-documento funcionam.
- [ ] INDEX.md lista todos os arquivos existentes.
- [ ] CHANGELOG.md registra as mudanças desta execução.
- [ ] Terminologia consistente com `domain/domain.md`.

---

## Passo 5 — Sanity Check Final 🩺

1. **Links**: Verificar todos os links internos em `INDEX.md` e cross-referências entre documentos.
2. **Mermaid**: Confirmar que diagramas Mermaid em `architecture/c4-*.md` e `architecture/erd.md` têm sintaxe válida.
3. **OpenAPI**: Rodar `npx @redocly/cli lint {SPECS_DIR}/api/tax-rates-api.yaml` e corrigir erros.
4. **Docstrings**: Verificar que exemplos de código em docstrings compilam/executam (Passo 2.5).
5. **Segurança**: Validar `security/findings.json` como JSON válido e verificar que achados CRITICAL/HIGH têm cenários de ataque concretos.
6. **Consistência de domínio**: Escolher 5 termos aleatórios do `domain/domain.md` e verificar que aparecem com o mesmo significado nos outros documentos.
7. **Cobertura**: Comparar estrutura final com a estrutura canônica (Passo 0) e reportar:
   - Artefatos criados: N
   - Artefatos atualizados: N
   - Artefatos ainda ausentes: N (com justificativa)
8. **docs/codebase/**: Verificar que os 7 documentos existem e não contêm `[TODO]` não resolvidos.
9. **Git diff**: Executar `git diff --stat {SPECS_DIR}/` para auditar todas as mudanças.

---

## Passo 6 — Geração do Relatório de Execução

1. **Nome do arquivo**: `{AAAA-MM-DD-HHMMSS}-spec-mining.md`
   - Use o timestamp **real do momento da criação**: `date +%Y-%m-%d-%H%M%S`.
2. **Local**: `{SPECS_DIR}/skill-output/`
3. **Conteúdo**:

```markdown
# 📑 Relatório de Mineração de Especificações (SPEC-MINING)

* **Data e Hora:** {AAAA-MM-DD HH:MM:SS} (GMT-3)
* **Skills orquestradas:** acquire-codebase-knowledge → spec-miner → domain-modeling → api-designer → code-documenter → architecture-designer → security-audit → documentation-writer
* **Solução:** {SOLUTION_NAME} — {SOLUTION_DESCRIPTION}
* **Escopo:** {SCOPE}

---

## �​ Resumo da Execução

{1 parágrafo descrevendo o que foi feito, quais skills foram usadas e em qual ordem.}

## 📁 Artefatos Processados

| Ação | Arquivo | Skill | Mudança |
|---|---|---|---|
| 🆕 | `domain/domain.md` | domain-modeling | Criado — glossário com 45 termos |
| 🔄 | `api/tax-rates-api.yaml` | api-designer | Atualizado — 3 novos endpoints |
| 🔄 | `architecture/c4-context.md` | architecture-designer | Atualizado — novo ator externo |
| 🆕 | `architecture/c4-containers.md` | architecture-designer | Criado — 5 containers |
| ... | ... | ... | ... |

## 📊 Cobertura Final

| Pasta | Arquivos Esperados | Arquivos Existentes | Cobertura |
|---|---|---|---|
| docs/codebase/ | 7 | 7 | 100% |
| specs/ | 1 | 1 | 100% |
| domain/ | 1 | 1 | 100% |
| api/ | 1 | 1 | 100% |
| architecture/ | 10 | 8 | 80% |
| engineering/ | 2 | 2 | 100% |
| product/ | 3 | 3 | 100% |
| design/ | 3 | 0 | 0% (não aplicável) |
| governance/ | 2 | 2 | 100% |
| security/ | 3 | 3 | 100% |
| business-projects/ | 5+ | 5+ | 100% |

## �⚠️ Artefatos Ausentes (com justificativa)

- `architecture/c4-code-class.md` — pendente: aguardando estabilização do refactor do pacote `internal/reforma/`
- `design/*` — não aplicável: solução backend sem UI

## 🔗 Links Gerados

- INDEX.md atualizado: `{SPECS_DIR}/INDEX.md`
- CHANGELOG.md atualizado: `{SPECS_DIR}/CHANGELOG.md`

---

🤖 *Documentação gerada pela orquestração de 8 skills de IA: acquire-codebase-knowledge → spec-miner → domain-modeling → api-designer → code-documenter → architecture-designer → security-audit → documentation-writer.*
```

---

# 🔄 Modos de Execução

O prompt suporta 8 modos, controlados pelo parâmetro `{SCOPE}`:

| Modo | `{SCOPE}` | Comportamento |
|---|---|---|
| **Completo** | `full` | Executa Passos 0→6 integralmente, criando/atualizando todos os artefatos |
| **Scan + Spec** | `scan-only` | Executa apenas Passo 0 + Passo 0.5 + Passo 6 (scan de codebase + engenharia reversa) |
| **Domínio apenas** | `domain-only` | Executa Passo 0 + Passo 0.5 + Passo 1 + Passo 6 |
| **API apenas** | `api-only` | Executa Passo 0 + Passo 0.5 + Passo 1 + Passo 2 + Passo 6 (requer domínio estável) |
| **Código apenas** | `code-analysis-only` | Executa Passo 0 + Passo 1 + Passo 2 + Passo 2.5 + Passo 6 (documentação inline de código) |
| **Arquitetura apenas** | `architecture-only` | Executa Passo 0 + Passo 0.5 + Passo 1 + Passo 2 + Passo 3 + Passo 6 (requer domínio + API estáveis) |
| **Segurança apenas** | `security-only` | Executa Passo 0 + Passo 3 + Passo 3.5 + Passo 6 (requer arquitetura estável) |
| **Delta** | `delta` | Executa apenas para os arquivos listados em `{TARGET_FILES}`, usando a skill dona correspondente + `documentation-writer` como revisor |
| **Revisão** | `review` | Executa apenas Passo 4 (documentation-writer) + Passo 5 + Passo 6, sem criar novos artefatos |

---

# ⚠️ Regras de Ouro

1. **Nunca sobrescrever sem auditoria**: antes de modificar um arquivo existente, fazer diff do conteúdo atual vs. novo e reportar no relatório final.
2. **Scan primeiro**: os Passos 0 (`acquire-codebase-knowledge`) e 0.5 (`spec-miner`) são pré-requisitos para todos os passos seguintes. Sem o scan automatizado e a especificação reversa, os passos de domínio, API e arquitetura trabalham com informações incompletas.
3. **Domínio antes de API e Arquitetura**: o Passo 1 (`domain-modeling`) é pré-requisito para os Passos 2 e 3. Se o domínio não estiver estável, não faz sentido documentar API e arquitetura.
4. **Código é a verdade**: se houver conflito entre documentação existente e código, o código vence. Reportar a divergência e alinhar a documentação.
5. **Humano decide ambiguidades**: o `domain-modeling` e o `spec-miner` podem identificar termos conflitantes ou lógica implícita, mas apenas o humano pode resolvê-los.
6. **Mermaid é obrigatório**: todo diagrama arquitetural deve ser renderizável como Mermaid (não usar imagens estáticas).
7. **OpenAPI validado**: a spec da API deve passar na validação `@redocly/cli lint` sem erros.
8. **Docstrings validados**: exemplos de código em docstrings devem compilar/executar (Passo 2.5).
9. **Segurança com evidência**: todo achado de segurança deve ter cenário de ataque concreto e evidência de código (Passo 3.5).
10. **Diátaxis como camada final**: o `documentation-writer` revisa todos os artefatos, mas não redefine decisões técnicas dos passos anteriores — apenas melhora clareza, consistência e organização.
11. **INDEX.md é sagrado**: sempre atualizar o índice central ao final de cada execução.

---

# �📚 Referências

- **Diátaxis Framework**: [https://diataxis.fr/](https://diataxis.fr/)
- **C4 Model**: [https://c4model.com/](https://c4model.com/)
- **OpenAPI 3.1**: [https://spec.openapis.org/oas/v3.1.0](https://spec.openapis.org/oas/v3.1.0)
- **Mermaid.js**: [https://mermaid.js.org/](https://mermaid.js.org/)
- **ADR Template**: [https://adr.github.io/madr/](https://adr.github.io/madr/)
- **EARS Format**: `references/ears-format.md` (via `spec-miner` skill)
- **Arc42**: [https://docs.arc42.org/](https://docs.arc42.org/)
- **OWASP Top 10 (2021)**: [https://owasp.org/www-project-top-ten/](https://owasp.org/www-project-top-ten/)
- **CVSS 4.0**: [https://www.first.org/cvss/v4-0/](https://www.first.org/cvss/v4-0/)
- **Acquire Codebase Knowledge**: `skills/acquire-codebase-knowledge/SKILL.md`
- **Spec Miner**: `skills/spec-miner/SKILL.md`
- **Code Documenter**: `skills/code-documenter/SKILL.md`
- **Security Audit**: `skills/security-audit/SKILL.md`
