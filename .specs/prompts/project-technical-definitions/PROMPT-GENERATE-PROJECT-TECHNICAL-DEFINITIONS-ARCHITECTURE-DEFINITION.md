# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md` — a **definição de arquitetura do projeto** que especifica como todas as soluções técnicas se integram. Este documento **absorve e substitui** o INTEGRATION-MAP.md existente (criado ad-hoc) e define o "C4 Level 1 e 2 do projeto inteiro".

**Inputs upstream:** `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` (Fase 2) + `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` (Fase 4) + `architecture/` (ADRs, blueprints, data standards globais) + TECHNICAL-PLAN.md (referência).

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
Ler Catálogo de Soluções (Fase 2), PRD Definition (Fase 4), ADRs globais, blueprints (Java, Go, React), data standards (Protobuf, Avro), TECHNICAL-PLAN.md e INTEGRATION-MAP.md (referências).

### Passo 2 — Invocar Skills Especializadas
Invocar skills de C4, arquitetura, integração, mensageria e infra para projetar a arquitetura do sistema completo.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md` com:
- Diagrama C4 Level 1 (System Context): todas as soluções + sistemas externos
- Diagrama C4 Level 2 (Container): comunicação entre serviços
- Matriz de integração: origem → destino → protocolo → autenticação
- Topologia de deploy (containers, rede, ambientes Dev/Staging/Prod)
- Estratégia de comunicação: síncrono (REST/gRPC) vs. assíncrono (mensageria/eventos)
- Diagramas de sequência para fluxos cross-solution críticos
- ADRs de integração (decisões que afetam múltiplas soluções)

### Passo 4 — Validação Pós-Geração
Verificar: diagramas C4 presentes, matriz de integração completa, topologia definida, ADRs documentados.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `c4-architecture-c4-architecture` | Diagramas C4 Level 1-2 | C4 |
| 2 | `c4-context` | System Context diagram | C4 |
| 3 | `c4-container` | Container diagram | C4 |
| 4 | `architecture-patterns` | Padrões arquiteturais cross-solution | Arquitetura |
| 5 | `architecture-decision-records` | ADRs de integração | ADR |
| 6 | `create-architectural-decision-record` | Criar ADRs formais | ADR |
| 7 | `api-design-principles` | Design de APIs entre soluções | Integração |
| 8 | `api-patterns` | Padrões de API REST/gRPC | Integração |
| 9 | `openapi-spec-generation` | Contratos de API cross-solution | Integração |
| 10 | `event-sourcing-architect` | Estratégia de eventos entre serviços | Mensageria |
| 11 | `saga-orchestration` | Orquestração de sagas cross-solution | Mensageria |
| 12 | `kubernetes-architect` | Topologia de deploy K8s | Infra |
| 13 | `docker-expert` | Containerização das soluções | Infra |
| 14 | `cloud-design-patterns` | Padrões de cloud | Cloud |
| 15 | `domain-driven-design` | Bounded contexts cross-solution | DDD |
| 16 | `deployment-pipeline-design` | Pipeline de deploy cross-solution | DevOps |
| 17 | `mermaid-expert` | Diagramas Mermaid para o documento | Diagramas |
| 18 | `documentation-writer` | Redigir o Architecture Definition | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da definição de arquitetura | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
