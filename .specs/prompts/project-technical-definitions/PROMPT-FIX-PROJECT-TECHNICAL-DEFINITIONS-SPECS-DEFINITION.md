# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório de falha.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório de Falha e Artefatos
Ler `SPECS_DEFINITION_SCOPE_FAIL_REPORT.md`, o documento atual, Stack Matrix, blueprints.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Seção obrigatória vazia | Preencher seção completa com padrões e exemplos |
| P1 | Padrão contradiz Stack Matrix | Alinhar com a stack definida |
| P1 | Padrão viola regra de segurança | Alinhar com Security Definition |
| P2 | Padrão sem exemplo | Adicionar exemplo concreto |
| P3 | Restrição sem valor | Especificar valor numérico |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `coding-guidelines` | Corrigir padrões de código | Código |
| 3 | `api-documentation` | Corrigir padrões de API | API |
| 4 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da baseline de especificações | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
