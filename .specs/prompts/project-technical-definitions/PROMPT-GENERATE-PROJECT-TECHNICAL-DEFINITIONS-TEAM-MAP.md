# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md` — a **matriz de skills** do time técnico do projeto. Diferente do `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (que responde QUEM está disponível e quantas horas/dia), este documento foca em O QUE cada perfil sabe fazer e seu nível de proficiência técnica.

O artefato é complementar aos outros documentos de time e serve como referência para alocação de tarefas, identificação de gaps de competências e planejamento de capacitação.

**Inputs upstream:** `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (perfis e pessoas) + documentos de negócio (para entender escopo técnico necessário).

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio | `/home/user/work/business-inputs/business-projects` |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions | Derivado do PROJECT_PATH |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Verificar se TODOS os parâmetros foram informados.

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` para conhecer os perfis existentes, e documentos de negócio (Charter, BRD, Features) para entender as demandas técnicas.

### Passo 2 — Invocar Skills Especializadas
Invocar skills para mapear competências necessárias vs. disponíveis, gerar matriz de skills com níveis de proficiência e identificar gaps.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md` com:
- Matriz Perfil × Tecnologia × Nível (★☆☆ a ★★★)
- Skills por categoria (Linguagens, Frameworks, Bancos, Cloud, DevOps, Segurança, Frontend, Mobile)
- Gap analysis: skills necessárias vs. disponíveis
- Recomendações de contratação/capacitação
- Referência ao TEAM-CAPACITY.md para nomes e contatos

### Passo 4 — Validação Pós-Geração
Verificar: arquivo no caminho correto, matriz preenchida para todos os perfis, gaps documentados, referência ao TEAM-CAPACITY presente.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `team-composition-analysis` | Analisar composição do time e identificar gaps | People |
| 2 | `gap-analysis` | Análise de gaps de competências vs. necessidades | Análise |
| 3 | `skill-audit` | Auditoria de skills existentes no time | Discovery |
| 4 | `engineering-skills` | Validar skills de engenharia necessárias | Engenharia |
| 5 | `senior-architect` | Validar skills de arquitetura necessárias | Arquitetura |
| 6 | `senior-pm` | Validar skills de gestão necessárias | PM |
| 7 | `documentation-writer` | Redigir o TEAM-MAP.md consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da matriz de skills do time técnico | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
