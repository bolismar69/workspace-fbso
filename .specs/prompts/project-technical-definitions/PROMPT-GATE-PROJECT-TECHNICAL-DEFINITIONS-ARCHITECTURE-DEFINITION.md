# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Arquitetura** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md`. Verifica se a arquitetura do projeto está completa, consistente e cobre todas as soluções do catálogo.

**Princípio fundamental:** Toda solução do catálogo deve aparecer nos diagramas C4 e na matriz de integração. Nenhuma solução pode ficar "desconectada" da arquitetura.

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

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md`, Catálogo de Soluções, PRD Definition, ADRs globais.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Soluções no C4 L1 | Toda solução do catálogo aparece no System Context |
| 1.2 | Soluções no C4 L2 | Containers mapeados para todas as soluções |
| 1.3 | Integrações documentadas | Cada par origem→destino tem protocolo e autenticação |

#### Dimensão 2: Completude Técnica
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Diagramas C4 presentes | L1 e L2 com sintaxe C4 correta |
| 2.2 | Topologia de deploy | Ambientes e infra definidos |
| 2.3 | Comunicação síncrona/assíncrona | Estratégia documentada |
| 2.4 | ADRs de integração | ≥3 ADRs cross-solution |

#### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com blueprints | Soluções seguem blueprints da pasta architecture/ |
| 3.2 | Consistência com PRD Definition | Funcionalidades cross-solution têm integração definida |

### Passo 3 — Calcular Veredito
### Passo 4 — Gerar Relatório de Falha (se REPROVADO)

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `c4-architecture-c4-architecture` | Validar diagramas C4 | C4 |
| 2 | `architecture-patterns` | Validar padrões arquiteturais | Arquitetura |
| 3 | `architect-review` | Revisão de arquitetura | Arquitetura |
| 4 | `gap-analysis` | Identificar soluções desconectadas | Análise |
| 5 | `senior-architect` | Validação sênior da arquitetura | Arquitetura |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da definição de arquitetura | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
