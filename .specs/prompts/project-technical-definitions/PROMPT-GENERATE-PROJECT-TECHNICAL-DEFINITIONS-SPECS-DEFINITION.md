# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` — a **baseline de especificações técnicas** que TODAS as soluções do projeto devem seguir. Este documento garante consistência cross-solution em convenções, padrões e restrições técnicas.

**Inputs upstream:** `PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md` (Fase 5) + `PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md` (Fase 6) + `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` (Fase 4) + `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md` (Fase 3) + `architecture/blueprint/`.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{ARCHITECTURE_GLOBAL}` | Caminho da pasta de arquitetura global |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler Architecture Definition, Security Definition, PRD Definition, Stack Matrix, blueprints (pom.xml templates, Dockerfiles, application.yml, Checkstyle, Sonar configs).

### Passo 2 — Invocar Skills Especializadas
Invocar skills de API, banco de dados, mensageria, observabilidade e padrões de código.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` com:
- Convenções de nomenclatura (pastas, pacotes, endpoints, filas, tópicos)
- Padrões de API (URL design, versionamento, paginação, erros, headers)
- Padrões de banco de dados (nomenclatura, auditoria, soft delete, multi-tenant)
- Padrões de mensageria (formato de eventos, schema registry, dead letter)
- Padrões de logging e observabilidade (formato, níveis, tracing context)
- Restrições técnicas cross-solution (timeouts, rate limits, tamanhos máximos)
- Decisões de design que TODAS as soluções devem seguir
- Referência aos blueprints da pasta `architecture/blueprint/`

### Passo 4 — Validação Pós-Geração
Verificar: todas as seções preenchidas, padrões referenciam blueprints, restrições documentadas.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `api-documentation` | Padrões de documentação de API | API |
| 2 | `openapi-spec-generation` | Especificação OpenAPI cross-solution | API |
| 3 | `701-technologies-openapi` | Padrões OpenAPI | API |
| 4 | `database-design` | Padrões de design de banco | DB |
| 5 | `postgres-best-practices` | Boas práticas PostgreSQL | DB |
| 6 | `data-modeling` | Modelagem de dados cross-solution | DB |
| 7 | `observability-engineer` | Padrões de observabilidade | Observabilidade |
| 8 | `181-java-observability-logging` | Padrões de logging para Java | Observabilidade |
| 9 | `182-java-observability-metrics-micrometer` | Métricas com Micrometer | Observabilidade |
| 10 | `183-observability-tracing-opentelemetry` | Tracing com OpenTelemetry | Observabilidade |
| 11 | `coding-guidelines` | Diretrizes de código cross-solution | Código |
| 12 | `clean-code` | Princípios de código limpo | Código |
| 13 | `java-maven-best-practices` | Padrões Maven cross-solution | Java |
| 14 | `documentation-writer` | Redigir o Specs Definition | Documentação |

> **🔄 Flexibilidade:** Skills de linguagens/frameworks específicos devem ser trocados conforme stacks do projeto. Justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da baseline de especificações | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
