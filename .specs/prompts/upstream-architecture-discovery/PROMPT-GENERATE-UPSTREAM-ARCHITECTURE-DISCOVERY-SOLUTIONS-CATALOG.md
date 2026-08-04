# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-CATALOG
## Contexto
> 📐 **Discovery-Level:** Versão macro do catálogo de soluções para análise de viabilidade e ROM 50%.

Este prompt gera `DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — catálogo macro de soluções (nomes, tipos, propósito high-level). Sem detalhamento de stacks ou implementação.

**Papel no Bloco C (Catálogo, Matriz & Consolidação):** Fase 8 de 3. Consome os 6 artefatos do Bloco B e alimenta a Matriz (F9) e SPECS (F10).

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
6. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — Estratégia de Testes (F6)
7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7)


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
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — DevOps/SRE (F5)
6. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — Testes (F6)
7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Infra/Cloud (F7)

### Passo 2 — Invocar Skills
### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` com: nomes, tipos, propósito, complexidade estimada.

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md`

```markdown
# DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md
## Fase 8 — Bloco C: Catálogo, Matriz & Consolidação Discovery-Level

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-SOLUTIONS-CATALOG-v1.0 |
| **Versão** | 1.0 — Discovery-Level |
| **Data** | {DATA_ATUAL} |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Referência:** Fases F1-F7 do Bloco B

---

## Catálogo Macro de Soluções
- Tabela: ID | Solução | Tipo | Propósito High-Level | Épicos Atendidos | Status
- IDs no formato SOL-001, SOL-002, ...

### Resumo
- Tabela de métricas: Total de soluções, por categoria (aplicação, infraestrutura, automação, storage), por status (existentes, novas, a configurar)

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Catálogo Discovery-Level | Tech Lead |
```

### Passo 4 — Validação Pós-Geração

## Skills Utilizados
| 1 | `senior-architect` | Catalogação de soluções | 2 | `cloud-architect` | Tipos de serviço cloud |
| 3 | `documentation-writer` | Redigir catálogo |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — F8 Bloco C Discovery-Level | Time de Arquitetura |
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
| 8 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 9 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Upstream Architecture Discovery — Fase 8*
