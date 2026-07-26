# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Baseline de Especificações** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md`. Verifica se os padrões e convenções são completos, consistentes e aplicáveis a todas as soluções.

**Princípio fundamental:** Toda solução do catálogo deve conseguir implementar seguindo apenas esta baseline + seus blueprints. Nenhum padrão ambíguo ou contraditório é permitido.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md`, Stack Matrix, Architecture Definition, blueprints.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Completude
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Seções obrigatórias | API, DB, Mensageria, Logging, Restrições — todas preenchidas |
| 1.2 | Exemplos concretos | Cada padrão tem exemplo de código/config |
| 1.3 | Referências a blueprints | Decisões referenciam arquivos em architecture/blueprint/ |

#### Dimensão 2: Consistência Técnica
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Consistência com Stack Matrix | Padrões não contradizem stacks definidas |
| 2.2 | Consistência com Security | Padrões respeitam regras de segurança |
| 2.3 | Consistência interna | Sem contradições entre seções |

#### Dimensão 3: Aplicabilidade
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Cobertura de stacks | Padrões cobrem todas as linguagens/frameworks da Stack Matrix |
| 3.2 | Restrições mensuráveis | Timeouts, rate limits, tamanhos — valores numéricos |

### Passo 3 — Calcular Veredito
### Passo 4 — Gerar Relatório de Falha (se REPROVADO)

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `code-review` | Revisar padrões de código | Qualidade |
| 2 | `api-documentation` | Validar padrões de API | API |
| 3 | `database-design` | Validar padrões de banco | DB |
| 4 | `observability-engineer` | Validar padrões de observabilidade | Observabilidade |
| 5 | `gap-analysis` | Identificar padrões faltantes | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da baseline de especificações | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
