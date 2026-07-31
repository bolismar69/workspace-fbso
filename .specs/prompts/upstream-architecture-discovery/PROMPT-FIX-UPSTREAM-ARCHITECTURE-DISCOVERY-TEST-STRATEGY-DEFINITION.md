# PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

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

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual, Architecture Definition (F7), Security Definition (F8), DevOps SRE (F10), ADRs globais.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Pirâmide sem % de cobertura | Definir percentuais mínimos por nível e camada |
| P0 | Solução sem framework de teste | Especificar framework por linguagem/solução |
| P0 | Quality gates ausentes | Definir critérios objetivos de aprovação |
| P1 | Testes de segurança não alinhados | Alinhar SAST/DAST com SECURITY-DEFINITION |
| P2 | Performance testing sem thresholds | Definir thresholds alinhados com SLOs do DEVOPS-SRE |
| P3 | Ambiente de teste não especificado | Documentar provisioning, dados, isolamento |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `senior-qa` | Corrigir estratégia de testes | QA |
| 3 | `test-strategy-design` | Corrigir pirâmide e automação | Estratégia |
| 4 | `performance-testing-review-ai-review` | Corrigir testes de performance | Performance |
| 5 | `security-testing` | Corrigir testes de segurança | Segurança |
| 6 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção da definição de estratégia de testes | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
