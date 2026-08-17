# PROMPT-FIX-590-TECHNICAL-DISCOVERY

## Contexto

Este prompt é acionado quando o gate reprova os contratos de uma sprint em `technical-discovery/590-ciclo-NNN/`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve os arquivos do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance nos contratos afetados.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{SPRINT_NUMBER}` | Número da sprint alvo (ex: 01, 02, 03) |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), os 5 contratos da sprint, PACKAGE-BACKLOG, artefatos base do Bloco B.

### Passo 2 — Processar NCs por Prioridade

| Prioridade | Tipo de NC | Contrato Afetado | Ação Corretiva |
|---|---|---|---|
| P0 | Tarefa da sprint sem contrato | Todos | Criar contrato(s) faltante(s) para a tarefa |
| P0 | Endpoint sem auth definido | CONTRACTS-API | Adicionar esquema de autenticação |
| P0 | Tabela sem migration | CONTRACTS-DATA | Adicionar operação de schema/migration |
| P0 | Controle IAM ausente | CONTRACTS-SECURITY | Adicionar regra IAM ou validação |
| P0 | SLO não definido | CONTRACTS-SRE | Adicionar SLO para o serviço |
| P1 | Contrato não referencia US-ID | Qualquer | Adicionar referência à US correspondente |
| P1 | Contrato não referencia artefato base | Qualquer | Adicionar referência ao artefato base |
| P2 | Referência cruzada entre contratos ausente | Qualquer | Adicionar link entre contratos |
| P2 | Especificação contradiz SPECS-DEFINITION | Qualquer | Alinhar com a baseline de especificações |
| P3 | Seção obrigatória vazia | Qualquer | Preencher seção conforme modelo |
| P3 | DEFINITION-INCREMENTS incompleto | DEFINITION-INCREMENTS | Adicionar atualizações retroativas e lições |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar NCs | Análise |
| 2 | `api-designer` | Corrigir contratos de API | API |
| 3 | `data-modeling` | Corrigir contratos de dados | Dados |
| 4 | `security-auditor` | Corrigir contratos de segurança | Segurança |
| 5 | `sre-engineer` | Corrigir contratos de SRE | SRE |
| 6 | `documentation-writer` | Atualizar documentos | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção dos contratos técnicos por sprint | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
