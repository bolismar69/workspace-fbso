# PROMPT-FIX-490-DATA-ARCHITECTURE-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `490-DATA-ARCHITECTURE-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{ARCHITECTURE_GLOBAL}` | Caminho da pasta de arquitetura global (ADRs, blueprints, data standards) |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual, Architecture Definition (F7), Security Definition (F8), ADRs globais e data standards.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Entidade crítica sem storage definido | Mapear storage strategy para a entidade |
| P0 | Pipeline de dados sem origem/destino | Documentar origem, destino, schedule, formato |
| P0 | Dado sensível sem governança de privacidade | Adicionar controles LGPD/GDPR |
| P1 | ERD incompleto | Adicionar entidades e relacionamentos faltantes |
| P2 | Dicionário de dados incompleto | Preencher tipos, tamanhos, descrições |
| P3 | Inconsistência com ARCHITECTURE/SECURITY | Alinhar containers de dados e criptografia |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `senior-data-engineer` | Corrigir arquitetura de dados | Engenharia de Dados |
| 3 | `data-modeling` | Corrigir modelos e ERD | Modelagem |
| 4 | `database-architect` | Corrigir estratégia de SGBDs | Banco de Dados |
| 5 | `data-quality-frameworks` | Corrigir governança | Governance |
| 6 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção da definição de arquitetura de dados | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
