# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md` — a **definição de arquitetura de dados do projeto** que especifica modelagem, armazenamento, pipelines, integrações e governança de dados. Este documento é a referência técnica para todas as decisões relacionadas a dados no projeto.

**Relação com ARCHITECTURE-DEFINITION:** Enquanto o ARCHITECTURE-DEFINITION (F7) define a topologia de containers e integrações entre soluções, este documento detalha **como os dados fluem, são armazenados, transformados e governados** dentro e entre essas soluções.

**Inputs upstream:** `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` (Fase 4) + `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md` (Fase 5) + `PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md` (Fase 7) + `PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md` (Fase 8) + `{ARCHITECTURE_GLOBAL}` (ADRs, blueprints, data standards).

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints, data standards) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler PRD Definition (F4 — requisitos de negócio que demandam dados), Team Skills Map (F5 — perfis de banco/dados do time), Architecture Definition (F7 — containers e integrações que consomem dados), Security Definition (F8 — requisitos de criptografia, IAM de dados), ADRs globais, blueprints, data standards (Protobuf, Avro, JSON Schema).

### Passo 2 — Invocar Skills Especializadas
Invocar skills de modelagem de dados, engenharia de dados, banco de dados e governance para projetar a arquitetura de dados completa.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-DATA-ARCHITECTURE-DEFINITION.md` com:

1. **Modelagem de Dados** — ERD (entidade-relacionamento), schemas (lógicos e físicos), catálogo de entidades por solução, dicionário de dados com tipos, tamanhos, restrições
2. **Estratégia de Armazenamento** — SQL (PostgreSQL, MySQL, SQL Server), NoSQL (MongoDB, Redis, Cassandra), Cache (Redis), Data Warehouse/Lake (Snowflake, BigQuery, S3/Delta Lake)
3. **Pipelines de Dados** — ETL/ELT (batch), streaming (Kafka, Kinesis, Pulsar), processamento em tempo real vs. near-real-time, schedule de jobs
4. **Integrações Inter-Banco** — data services, APIs de dados, CDC (Debezium), replicação, federação de dados
5. **Data Governance** — qualidade de dados (expectativas, validação), linhagem (OpenLineage), catálogo (DataHub/Amundsen), privacidade (LGPD/GDPR — anonimização, pseudonimização, retenção, purge)
6. **Estratégia On-Premise vs Cloud** — comparação de custo/performance, justificativa da escolha, plano de migração (se aplicável)
7. **Tecnologias e Ferramentas** — SGBDs, ferramentas ETL (Airbyte/Fivetran/dbt), plataformas de streaming, orquestradores (Airflow/Dagster/Prefect)

### Passo 4 — Validação Pós-Geração
Verificar: ERD presente com entidades de todas as soluções, storage strategy definida por tipo de dado, pipelines documentados com schedule, data governance coberto (qualidade + linhagem + privacidade), consistência com ARCHITECTURE-DEFINITION (containers que consomem dados) e SECURITY-DEFINITION (criptografia em repouso, IAM de dados).

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

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da definição de arquitetura de dados | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
