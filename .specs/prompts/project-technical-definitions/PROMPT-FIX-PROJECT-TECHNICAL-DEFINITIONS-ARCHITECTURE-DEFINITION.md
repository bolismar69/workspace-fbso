# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório de falha.

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

### Passo 1 — Carregar Relatório de Falha e Artefatos
Ler `ARCHITECTURE_DEFINITION_SCOPE_FAIL_REPORT.md`, o documento atual, catálogo de soluções e ADRs.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Solução ausente dos diagramas C4 | Adicionar solução ao L1 e L2 |
| P0 | Integração crítica não documentada | Documentar protocolo, autenticação, fluxo |
| P1 | Diagrama C4 incompleto | Completar elementos faltantes |
| P2 | ADR de integração insuficiente | Adicionar ADR com ID, decisão, alternativas |
| P3 | Inconsistência com blueprints | Alinhar com padrão do architecture/ |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `c4-architecture-c4-architecture` | Corrigir diagramas C4 | C4 |
| 3 | `architecture-patterns` | Corrigir padrões arquiteturais | Arquitetura |
| 4 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da definição de arquitetura | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
