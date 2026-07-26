# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP

## Contexto

Este prompt é acionado quando o `PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md` reprova o artefato `PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md` com veredito `REPROVADO`.

O agente corretor atua como **cirurgião de skills** — aplica correções pontuais na matriz com base no relatório de falha, preservando todas as seções aprovadas. **Nunca reescreve o documento do zero.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 1 — Carregar Artefatos e Relatório de Falha
Ler `TEAM_MAP_SCOPE_FAIL_REPORT.md` (relatório de falha), `PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md` (artefato a corrigir) e `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (referência).

### Passo 2 — Processar Não-Conformidades por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Perfil sem skills mapeadas | Adicionar matriz completa de skills para o perfil |
| P1 | Categoria obrigatória ausente | Preencher categoria com tecnologias relevantes |
| P2 | Nível de proficiência indefinido | Atribuir nível (★☆☆ a ★★★) com justificativa |
| P3 | Gap não documentado | Documentar gap com recomendação |

### Passo 3 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção afetada, aplicar correção, preservar conteúdo não afetado.

### Passo 4 — Validar Correções
Verificar: 100% P0 resolvidas, todos os perfis cobertos, gaps documentados, consistência mantida.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório de falha e priorizar | Análise |
| 2 | `skill-audit` | Corrigir skills ausentes ou incompletas | Discovery |
| 3 | `team-composition-analysis` | Validar correções de composição | People |
| 4 | `documentation-writer` | Atualizar TEAM-MAP.md com correções | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da matriz de skills | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
