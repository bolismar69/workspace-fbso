# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md` — a **matriz-mestra** que consolida todas as definições anteriores em uma tabela única: solução → responsável → repositório → stack → perfis → status.

**Inputs upstream:** Todas as fases anteriores (1-8) + `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` + `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md`.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler TODOS os artefatos das fases 1-8 + TEAM-CAPACITY.md + TEAM-CAPACITY-EXCEPTIONS.md.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de mapeamento, portfolio e documentação.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-MATRIX.md` com:
- **Tabela-mestra:** Solução | Tipo | Repositório/Pasta | Stack | Tech Lead | Time Alocado | Perfis Necessários | Status
- Referência ao TEAM-CAPACITY.md para nomes e contatos
- Referência ao TEAM-CAPACITY-EXCEPTIONS.md para exceções de capacidade
- Capacidade alocada vs. necessária por solução
- Indicadores: cobertura de skills (%), risco de gargalo, status geral
- Matriz RACI simplificada (Responsável, Autoridade, Consultado, Informado)

### Passo 4 — Validação Pós-Geração
Verificar: todas as soluções na matriz, responsáveis definidos, stacks referenciadas, capacidade calculada.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `reference-builder` | Construir matriz de referência cruzada | Mapeamento |
| 2 | `team-composition-analysis` | Analisar alocação de time por solução | People |
| 3 | `project-manager` | Validar alocação de recursos | PM |
| 4 | `track-management` | Estruturar tracking de soluções | Portfolio |
| 5 | `stakeholder-map` | Mapear stakeholders por solução | Stakeholder |
| 6 | `documentation-writer` | Redigir a matriz consolidada | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da matriz de soluções | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
