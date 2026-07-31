# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-MATRIX
## Contexto
> 📐 **Discovery-Level:** Matriz macro solução×disciplina×complexidade para ROM 50%.

Este prompt gera `DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — matriz que cruza soluções com disciplinas técnicas e classifica complexidade (Baixa/Média/Alta) para embasar estimativas.

**Papel no Bloco C:** Fase 9 de 3. Consome PRD, 6 artefatos Bloco B e Catálogo (F8).

**Inputs upstream:** DISCOVERY-LEVEL-PRD (F1) + Bloco B (F2-F7) + SOLUTIONS-CATALOG (F8).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{UPSTREAM_DISCOVERY_PATH}` | Caminho upstream-architecture-discovery |

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
### Passo 1 — Carregar Documentos Base
Ler PRD (F1) + 6 artefatos Bloco B (F2-F7) + Catálogo (F8).
### Passo 2 — Invocar Skills
### Passo 3 — Gerar matriz cruzando solução × disciplina × complexidade (Baixa/Média/Alta) + riscos identificados + gaps de conhecimento
### Passo 4 — Validação Pós-Geração

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-architect` | Cruzamento solução×disciplina e avaliação de complexidade | Arquitetura |
| 2 | `gap-analysis` | Identificação de gaps e riscos entre disciplinas | Análise |
| 3 | `documentation-writer` | Redigir a matriz Discovery-Level | Documentação |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — F9 Bloco C Discovery-Level | Time de Arquitetura |
🤖 *Upstream Architecture Discovery — Fase 9*
