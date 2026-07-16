# Contexto:
  - Este prompt orquestra **4 skills especializadas** para minerar, criar e atualizar a documentação técnica completa da solução no diretório `.specs/`.
  - As skills são invocadas em sequência lógica: primeiro entende-se o domínio, depois desenha-se a API, depois a arquitetura, e por fim aplica-se o framework Diátaxis para unificar qualidade e consistência.
  - A estrutura de pastas alvo (descrita abaixo) é a referência canônica — o prompt pode enriquecê-la com novos artefatos quando o contexto da solução exigir.
  - A documentação existente em `.specs/` deve ser preservada e incrementada, nunca sobrescrita sem auditoria de delta.
  - Use o README.md da raiz da solução para instruções de build, execução e testes.

---

# 🎯 Skills Orquestradas

| Ordem | Skill | Responsabilidade | Pasta(s) Principal(is) |
|---|---|---|---|
| 1ª | `domain-modeling` | Glossário de domínio, ubiquitous language, ADRs de domínio | `domain/` |
| 2ª | `api-designer` | Contrato OpenAPI 3.1, modelagem de recursos, versionamento, segurança de API | `api/` |
| 3ª | `architecture-designer` | C4 (Context→Containers→Components→Code), ADRs, ERD, dicionário de dados, integrações | `architecture/` |
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
| `{SCOPE}` | Escopo da documentação a ser gerada | `full`, `domain-only`, `api-only`, `architecture-only`, `delta` |
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
| `domain/domain.md` | `domain-modeling` | Criar/atualizar glossário, resolver ambiguidades de linguagem |
| `api/tax-rates-api.yaml` | `api-designer` | Gerar/atualizar contrato OpenAPI 3.1 completo |
| `architecture/architecture.md` | `architecture-designer` | Escrever visão geral, padrões, trade-offs |
| `architecture/c4-*.md` | `architecture-designer` | Gerar diagramas C4 nos 4 níveis (Mermaid) |
| `architecture/erd.md` | `architecture-designer` | Modelar entidades e relacionamentos |
| `architecture/data-dictionary.md` | `architecture-designer` | Documentar tabelas, colunas, tipos, regras |
| `architecture/integrations.md` | `architecture-designer` | Mapear dependências externas e libs |
| `architecture/adrs/*.md` | `architecture-designer` | Registrar decisões arquiteturais (ADR) |
| `engineering/api-guidelines.md` | `api-designer` → `documentation-writer` | Padrões de API REST, erros, segurança |
| `engineering/code-analysis.md` | `code-documenter` → `documentation-writer` | Análise de fluxo de handlers/services |
| `product/product.md` | `documentation-writer` | Descrição do produto (Explanation — Diátaxis) |
| `product/requirements.md` | `documentation-writer` | Requisitos (Reference — Diátaxis) |
| `product/feature-roadmap.md` | `documentation-writer` | Roadmap (Reference — Diátaxis) |
| `design/*.md` | `documentation-writer` | Documentação de design system (se aplicável) |
| `governance/inventory.md` | `documentation-writer` | Inventário e cobertura |
| `governance/confidence-report.md` | `documentation-writer` | Score de confiança da documentação |
| `security/SECURITY.md` | `documentation-writer` | Definições de segurança (se aplicável) |
| `INDEX.md` | `documentation-writer` | Índice central (atualizado a cada ciclo) |
| `CHANGELOG.md` | `documentation-writer` | Registro de mudanças da doc |
| `business-projects/{ID}/*.md` | `documentation-writer` | PRD, SPECS, TASKS, TEST_PLAN, ARCHITECTURE |
| `features/{NAME}/*.md` | `documentation-writer` | SPECS, ARCHITECTURE, TASKS, TEST_PLAN |

---

# 🔁 Protocolo de Execução (Passo a Passo)

## Passo 0 — Leitura do Cenário Atual

1. Confirmar que `{SOLUTION_ROOT}` e `{SPECS_DIR}` estão acessíveis.
2. Ler `{SPECS_DIR}/INDEX.md` (se existir) para entender a estrutura atual.
3. Ler `{SPECS_DIR}/CHANGELOG.md` (se existir) para conhecer o histórico.
4. Fazer scan dos arquivos existentes em cada subpasta e comparar com a estrutura canônica acima.
5. Gerar um **delta inicial**: lista de arquivos existentes (📄), ausentes (❌), e desatualizados (🔄).
6. Confirmar com o humano o escopo antes de prosseguir (especialmente se `{SCOPE}=full`).

---

## Passo 1 — Domain Modeling (`domain-modeling`) 🏛️

**Objetivo:** Garantir que o glossário de domínio e a linguagem ubíqua estejam corretos e completos **antes** de documentar API e arquitetura.

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

## Passo 3 — Architecture Design (`architecture-designer`) 🏗️

**Objetivo:** Produzir/atualizar a documentação arquitetural completa nos 4 níveis C4, ERD, dicionário de dados, integrações e ADRs.

### Pré-condição:
- `domain/domain.md` e `api/tax-rates-api.yaml` atualizados (a arquitetura referencia ambos).

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

## Passo 4 — Documentation Writer (`documentation-writer`) 📝

**Objetivo:** Aplicar o framework Diátaxis para revisar, unificar e expandir toda a documentação, garantindo qualidade textual e consistência cross-documento.

### Pré-condição:
- Passos 1, 2, 3 concluídos (a documentação técnica base existe).

### Ações:

#### 4.1 — Classificar Cada Artefato no Framework Diátaxis

| Quadrante | Pergunta-chave | Artefatos `.specs/` |
|---|---|---|
| **Tutorial** (learning) | "Como eu começo?" | `engineering/code-analysis.md` (fluxo guiado) |
| **How-to** (problem) | "Como resolvo X?" | `engineering/api-guidelines.md`, `architecture/integrations.md` |
| **Reference** (information) | "O que é X?" | `domain/domain.md`, `api/*.yaml`, `architecture/erd.md`, `architecture/data-dictionary.md`, `product/requirements.md`, `product/feature-roadmap.md` |
| **Explanation** (understanding) | "Por que X?" | `architecture/architecture.md`, `architecture/adrs/*.md`, `architecture/c4-*.md`, `product/product.md`, `governance/confidence-report.md` |

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

**Security** (se aplicável):
- `security/SECURITY.md` — Referência: threat model, práticas de segurança, OWASP, checklists.

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
4. **Consistência de domínio**: Escolher 5 termos aleatórios do `domain/domain.md` e verificar que aparecem com o mesmo significado nos outros documentos.
5. **Cobertura**: Comparar estrutura final com a estrutura canônica (Passo 0) e reportar:
   - Artefatos criados: N
   - Artefatos atualizados: N
   - Artefatos ainda ausentes: N (com justificativa)
6. **Git diff**: Executar `git diff --stat {SPECS_DIR}/` para auditar todas as mudanças.

---

## Passo 6 — Geração do Relatório de Execução

1. **Nome do arquivo**: `{AAAA-MM-DD-HHMMSS}-spec-mining.md`
   - Use o timestamp **real do momento da criação**: `date +%Y-%m-%d-%H%M%S`.
2. **Local**: `{SPECS_DIR}/skill-output/`
3. **Conteúdo**:

```markdown
# 📑 Relatório de Mineração de Especificações (SPEC-MINING)

* **Data e Hora:** {AAAA-MM-DD HH:MM:SS} (GMT-3)
* **Skills orquestradas:** domain-modeling → api-designer → architecture-designer → documentation-writer
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
| domain/ | 1 | 1 | 100% |
| api/ | 1 | 1 | 100% |
| architecture/ | 10 | 8 | 80% |
| engineering/ | 2 | 2 | 100% |
| product/ | 3 | 3 | 100% |
| design/ | 3 | 0 | 0% (não aplicável) |
| governance/ | 2 | 2 | 100% |
| security/ | 1 | 1 | 100% |
| business-projects/ | 5+ | 5+ | 100% |

## �⚠️ Artefatos Ausentes (com justificativa)

- `architecture/c4-code-class.md` — pendente: aguardando estabilização do refactor do pacote `internal/reforma/`
- `design/*` — não aplicável: solução backend sem UI

## 🔗 Links Gerados

- INDEX.md atualizado: `{SPECS_DIR}/INDEX.md`
- CHANGELOG.md atualizado: `{SPECS_DIR}/CHANGELOG.md`

---

🤖 *Documentação gerada pela orquestração de skills de IA: domain-modeling → api-designer → architecture-designer → documentation-writer.*
```

---

# 🔄 Modos de Execução

O prompt suporta 4 modos, controlados pelo parâmetro `{SCOPE}`:

| Modo | `{SCOPE}` | Comportamento |
|---|---|---|
| **Completo** | `full` | Executa Passos 0→6 integralmente, criando/atualizando todos os artefatos |
| **Domínio apenas** | `domain-only` | Executa apenas Passo 0 + Passo 1 + Passo 6 |
| **API apenas** | `api-only` | Executa Passo 0 + Passo 2 + Passo 6 (requer domínio estável) |
| **Arquitetura apenas** | `architecture-only` | Executa Passo 0 + Passo 3 + Passo 6 (requer domínio + API estáveis) |
| **Delta** | `delta` | Executa apenas para os arquivos listados em `{TARGET_FILES}`, usando a skill dona correspondente + `documentation-writer` como revisor |
| **Revisão** | `review` | Executa apenas Passo 4 (documentation-writer) + Passo 5 + Passo 6, sem criar novos artefatos |

---

# ⚠️ Regras de Ouro

1. **Nunca sobrescrever sem auditoria**: antes de modificar um arquivo existente, fazer diff do conteúdo atual vs. novo e reportar no relatório final.
2. **Domínio primeiro**: o Passo 1 (`domain-modeling`) é pré-requisito para os Passos 2 e 3. Se o domínio não estiver estável, não faz sentido documentar API e arquitetura.
3. **Código é a verdade**: se houver conflito entre documentação existente e código, o código vence. Reportar a divergência e alinhar a documentação.
4. **Humano decide ambiguidades**: o `domain-modeling` pode identificar termos conflitantes, mas apenas o humano pode resolvê-los.
5. **Mermaid é obrigatório**: todo diagrama arquitetural deve ser renderizável como Mermaid (não usar imagens estáticas).
6. **OpenAPI validado**: a spec da API deve passar na validação `@redocly/cli lint` sem erros.
7. **Diátaxis como camada final**: o `documentation-writer` revisa todos os artefatos, mas não redefine decisões técnicas dos passos anteriores — apenas melhora clareza, consistência e organização.
8. **INDEX.md é sagrado**: sempre atualizar o índice central ao final de cada execução.

---

# �📚 Referências

- **Diátaxis Framework**: [https://diataxis.fr/](https://diataxis.fr/)
- **C4 Model**: [https://c4model.com/](https://c4model.com/)
- **OpenAPI 3.1**: [https://spec.openapis.org/oas/v3.1.0](https://spec.openapis.org/oas/v3.1.0)
- **Mermaid.js**: [https://mermaid.js.org/](https://mermaid.js.org/)
- **ADR Template**: [https://adr.github.io/madr/](https://adr.github.io/madr/)
- **EARS Format**: `references/ears-format.md` (via `spec-miner` skill)
- **Arc42**: [https://docs.arc42.org/](https://docs.arc42.org/)
