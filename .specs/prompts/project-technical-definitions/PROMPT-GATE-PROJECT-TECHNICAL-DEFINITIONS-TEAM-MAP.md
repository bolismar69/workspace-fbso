# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP

## Contexto

Este prompt implementa o **Gate de Validação da Matriz de Skills** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md`. O agente validador verifica se a matriz de skills está completa, consistente e alinhada com os perfis documentados no `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md`.

**Princípio fundamental:** Todo perfil listado no TEAM-CAPACITY deve ter skills mapeadas no TEAM-MAP. Nenhum perfil pode ficar sem cobertura de competências.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md` (artefato a validar) e `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (referência de perfis).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Perfis
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Todos os perfis cobertos | Cada perfil do TEAM-CAPACITY tem entrada na matriz de skills |
| 1.2 | Sem perfis órfãos | Nenhuma skill mapeada para perfil inexistente no TEAM-CAPACITY |

#### Dimensão 2: Completude da Matriz
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Categorias obrigatórias | Linguagens, Frameworks, Bancos, Cloud, DevOps — todas preenchidas |
| 2.2 | Níveis de proficiência | ★☆☆ a ★★★ definidos para cada skill×perfil |
| 2.3 | Gap analysis | Gaps documentados com recomendações |

#### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com escopo | Skills mapeadas cobrem tecnologias mencionadas no TECHNICAL-PLAN.md |
| 3.2 | Nomes consistentes | Tecnologias usam nomenclatura padronizada (ex: "PostgreSQL", não "Postgres") |

### Passo 3 — Calcular Veredito
- 100% OK → APROVADO
- ≥ 75% OK → RESSALVA
- < 75% OK → REPROVADO

### Passo 4 — Gerar Relatório de Falha (se REPROVADO)
Gerar `TEAM_MAP_SCOPE_FAIL_REPORT.md` com NCs por dimensão, gravidade e ações esperadas.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `skill-audit` | Auditar skills contra TEAM-CAPACITY | Discovery |
| 2 | `gap-analysis` | Identificar lacunas na matriz | Análise |
| 3 | `team-composition-analysis` | Validar composição do time | People |
| 4 | `senior-architect` | Revisão de arquitetura de skills | Arquitetura |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no relatório de falha.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da matriz de skills | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
