# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md` — a **definição de arquitetura de dados do projeto** que especifica modelagem, armazenamento, pipelines, integrações e governança de dados. Este documento é a referência técnica para todas as decisões relacionadas a dados no projeto.

**Este documento é independente de tecnologias específicas de banco de dados.** Durante a análise da stack do projeto, identifique o banco de dados utilizado e busque skills relacionados a esse banco para aprimorar as especificações. Caso não encontre skills específicos, utilize skills generalistas de arquitetura de dados, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints, data standards

**Papel no Bloco B (Architecture & Security & Specialists):** Fase 9 de 6. Disciplina: Data Architect / Engenheiro de Dados.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints, data standards) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills necessários para o time de implementação (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time — seniores, plenos, juniores, duração (obter e validar com usuário) |
| `{PROJECT-STACK}` | Stack tecnológica da solução. Baseline corporativa: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão exigem justificativa |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `TECHNICAL_SOLUTION_PATH={TECHNICAL_SOLUTION_PATH}` — base das soluções técnicas
- `TECHNICAL_SOLUTION_NAMES={TECHNICAL_SOLUTION_NAMES}` — soluções do projeto
- `ARCHITECTURE_GLOBAL={ARCHITECTURE_GLOBAL}` — ADRs e blueprints globais
- `SECURITY_GLOBAL={SECURITY_GLOBAL}` — documento de segurança global
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP` — skills do time (se fornecidos)
- `PROJECT-TEAM-CAPACITY` — capacidade do time (se fornecida)
- `PROJECT-STACK` — stack tecnológica; validar contra STACK-PADROES-CORPORATIVOS-FBSO-ORG.md
Validar que `{UPSTREAM_DISCOVERY_PATH}` existe e contém os artefatos upstream.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos upstream:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2) — containers que consomem dados
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3) — criptografia, IAM de dados
4. `{ARCHITECTURE_GLOBAL}/` — ADRs globais, blueprints, data standards (Protobuf, Avro, JSON Schema)

### Passo 2 — Invocar Skills Especializadas
Invocar skills de modelagem de dados, engenharia de dados, banco de dados e governance para projetar a arquitetura de dados completa.

### Passo 2.5 — Apresentar Skills, Capacidade e Stack para Validação Humana

Avaliar e apresentar ao usuário para validação:

1. **PROJECT-TEAM-SKILLS-MAP:** Skills identificados como necessários para implementar a solução nesta disciplina.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

2. **PROJECT-TEAM-CAPACITY:** Capacidade estimada do time nesta disciplina (ex: 2 seniores, 3 plenos).
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

3. **PROJECT-STACK:** Tecnologias identificadas para esta disciplina. Verificar conformidade com `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão corporativo DEVEM ser listadas com justificativa técnica e requerem aprovação explícita do usuário.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md` com:

1. **Modelagem de Dados** — ERD (entidade-relacionamento), schemas (lógicos e físicos), catálogo de entidades por solução, dicionário de dados com tipos, tamanhos, restrições
2. **Estratégia de Armazenamento** — SQL (PostgreSQL, MySQL, SQL Server), NoSQL (MongoDB, Redis, Cassandra), Cache (Redis), Data Warehouse/Lake (Snowflake, BigQuery, S3/Delta Lake)
3. **Pipelines de Dados** — ETL/ELT (batch), streaming (Kafka, Kinesis, Pulsar), processamento em tempo real vs. near-real-time, schedule de jobs
4. **Integrações Inter-Banco** — data services, APIs de dados, CDC (Debezium), replicação, federação de dados
5. **Data Governance** — qualidade de dados (expectativas, validação), linhagem (OpenLineage), catálogo (DataHub/Amundsen), privacidade (LGPD/GDPR — anonimização, pseudonimização, retenção, purge)
6. **Estratégia On-Premise vs Cloud** — comparação de custo/performance, justificativa da escolha, plano de migração (se aplicável)
7. **Tecnologias e Ferramentas** — SGBDs, ferramentas ETL (Airbyte/Fivetran/dbt), plataformas de streaming, orquestradores (Airflow/Dagster/Prefect)

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md`

```markdown
# DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md
## Fase 4 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | {DATA_ATUAL} |
| **Autor** | Data Architect / DB Developer |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md) — PRD Discovery-Level (F1)
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)
- [`DISCOVERY-LEVEL-SECURITY-DEFINITION.md`](DISCOVERY-LEVEL-SECURITY-DEFINITION.md) — Definição de Segurança (F3)

---

## 1. Estratégia de Dados — Visão Macro
- **1.1 Abordagem de Armazenamento:** Tabela: Decisão | Escolha | Rationale (Banco Primário, Cache, Object Storage, Search)
- **1.2 Multi-Tenant — Modelo de Isolamento:** Diagrama ASCII da estrutura de schemas e isolamento

## 2. Entidades Macro — Modelo de Domínio
- **2.1 Entidades Core:** Tabela: Entidade | Descrição | Épico | Volume Estimado (Ano 1)
- **2.2 Diagrama de Entidades Macro:** Diagrama Mermaid erDiagram com todas as entidades, relacionamentos e atributos principais
- **2.3 Estrutura de Schemas:** Tabela: Schema | Propósito | Acesso

## 3. Estratégia de RLS — Row-Level Security (ou mecanismo equivalente)
- **3.1 Mecanismo:** Código SQL de exemplo da policy de isolamento
- **3.2 Tabelas com Isolamento vs. Tabelas Globais:** Classificação

## 4. Indexação e Performance
- **4.1 Índices Críticos:** Tabela: Tabela | Índice | Tipo | Justificativa
- **4.2 Volumes e Crescimento Projetado:** Tabela: Ano | Tenants | Usuários | Produtos | Audit Logs | DB Size Estimado

## 5. Estratégia de Cache
- Tabela: Padrão | TTL | Propósito

## 6. Riscos e Estimativa de Esforço
- **6.1 Riscos de Dados:** Tabela: ID | Risco | Prob. | Impacto | Mitigação
- **6.2 Estimativa de Esforço:** Tabela: Atividade | Complexidade | Esforço (dias) | Responsável

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Data Architecture Discovery-Level | Data Architect / DB Developer |
```

### Passo 4 — Validação Pós-Geração

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-data-engineer` | Supervisão sênior da arquitetura de dados | Engenharia de Dados |
| 2 | `data-engineer` | Projetar pipelines e armazenamento | Engenharia de Dados |
| 3 | `data-modeling` | Modelagem ERD, schemas lógicos/físicos | Modelagem |
| 4 | `database-architect` | Estratégia de SGBDs e otimização | Banco de Dados |
| 5 | `data-engineering-data-pipeline` | Pipelines ETL/ELT e streaming | Pipeline |
| 6 | `data-engineering-data-driven-feature` | Features orientadas a dados | Engenharia de Dados |
| 7 | `sql-pro` | Modelagem e otimização SQL | SQL |
| 8 | `postgres-best-practices` | Boas práticas PostgreSQL | PostgreSQL |
| 9 | `nosql-expert` | Estratégia NoSQL (MongoDB, Redis) | NoSQL |
| 10 | `database-design` | Design de banco de dados | Banco de Dados |
| 11 | `database-migration` | Migração de dados entre sistemas | Migração |
| 12 | `data-quality-frameworks` | Qualidade e governança de dados | Governance |
| 13 | `data-engineer` | Documentação do dicionário de dados | Engenharia de Dados |
| 14 | `mermaid-expert` | Diagramas ERD e fluxo de dados | Diagramas |
| 15 | `documentation-writer` | Redigir o Data Architecture Definition | Documentação |


**Skills generalistas de dados (sempre aplicáveis):**
- `engineering-skills`, `engineering-advanced-skills`
- `senior-data-engineer`, `database-architect`, `database-design`
- `data-engineer`, `data-modeling`, `database`
- `database-migrations`, `database-migrations-sql-migrations`

**Skills tecnológicos de banco de dados (condicionais — buscar ao identificar a stack):**
- Ao identificar um banco de dados específico durante a análise da stack, busque skills relacionados a esse banco para aprimorar as especificações e estimativas
- Caso não encontre skills específicos para o banco identificado, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da definição de arquitetura de dados | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) — requisitos de negócio |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) — containers e integrações |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) — criptografia e IAM |
| 4 | `{ARCHITECTURE_GLOBAL}/` | ADRs, blueprints, data standards |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (se fornecidos) |
| 7 | `{PROJECT-TEAM-CAPACITY}` | Capacidade do time (se fornecida) |
| 8 | `{PROJECT-STACK}` | Stack tecnológica (validar contra padrões corporativos) |
| 9 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
