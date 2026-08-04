# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-MATRIX
## Contexto
> 📐 **Discovery-Level:** Matriz macro solução×disciplina×complexidade para ROM 50%.

Este prompt gera `DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — matriz que cruza soluções com disciplinas técnicas e classifica complexidade (Baixa/Média/Alta) para embasar estimativas.

**Papel no Bloco C:** Fase 9 de 3. Consome PRD, 6 artefatos Bloco B e Catálogo (F8).

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
6. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — Estratégia de Testes (F6)
7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7)
8. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)


## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Documentos brutos adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `TECHNICAL_SOLUTION_PATH={TECHNICAL_SOLUTION_PATH}` — base das soluções técnicas
- `TECHNICAL_SOLUTION_NAMES={TECHNICAL_SOLUTION_NAMES}` — soluções do projeto
- `ARCHITECTURE_GLOBAL={ARCHITECTURE_GLOBAL}` — arquitetura global
- `SECURITY_GLOBAL={SECURITY_GLOBAL}` — segurança global
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
Validar que `{UPSTREAM_DISCOVERY_PATH}` existe e contém os artefatos upstream.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos upstream:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD (F1)
2-7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-*-DEFINITION.md` — Bloco B (F2-F7)
8. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo (F8)
### Passo 2 — Invocar Skills
### Passo 3 — Gerar matriz cruzando solução × disciplina × complexidade (Baixa/Média/Alta) + riscos identificados + gaps de conhecimento
---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md`

```markdown
# DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md
## Fase 9 — Bloco C: Catálogo, Matriz & Consolidação Discovery-Level

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-SOLUTIONS-MATRIX-v1.0 |
| **Versão** | 1.0 — Discovery-Level |
| **Data** | {DATA_ATUAL} |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Referência:** Fases F1-F8

---

## Matriz Solução × Disciplina × Complexidade
- Tabela cruzada: Solução | Arquitetura | Segurança | Dados | DevOps/SRE | Testes | Infra/Cloud | Complexidade Média
- Cada célula com indicador de complexidade: 🔴 Alta, 🟡 Média, 🟢 Baixa
- IDs referenciando o catálogo (SOL-001, SOL-002, ...)

### Legenda
- Descrição de cada nível de complexidade

### Distribuição de Complexidade
- Tabela: Nível | Soluções | %

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Matriz Discovery-Level | Tech Lead |
```

### Passo 4 — Validação Pós-Geração

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-architect` | Cruzamento solução×disciplina e avaliação de complexidade | Arquitetura |
| 2 | `gap-analysis` | Identificação de gaps e riscos entre disciplinas | Análise |
| 3 | `documentation-writer` | Redigir a matriz Discovery-Level | Documentação |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — F9 Bloco C Discovery-Level | Time de Arquitetura |
## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` | Definição de Dados (F4) |
| 5 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` | Definição DevOps/SRE (F5) |
| 6 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` | Estratégia de Testes (F6) |
| 7 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` | Definição Infra/Cloud (F7) |
| 8 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` | Catálogo de Soluções (F8) |
| 9 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 10 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Upstream Architecture Discovery — Fase 9*
