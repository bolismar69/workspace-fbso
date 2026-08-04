# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SPECS
## Contexto
> 📐 **Discovery-Level:** Consolidação técnica high-level — sumariza descobertas do Bloco B para embasar ROM.

Este prompt gera `DISCOVERY-LEVEL-SPECS.md` — documento enxuto que consolida as descobertas técnicas das 6 disciplinas em um resumo executivo para o comitê de governança.

**Papel no Bloco C:** Fase 10 de 3. É a saída final do Bloco C antes da estimativa ROM.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Definição de Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
6. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — Estratégia de Testes (F6)
7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7)
8. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
9. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — Matriz Solução×Disciplina (F9)


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
| `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

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
- `PROJECT-STACK` — stack tecnológica; validar conformidade com padrões corporativos
Validar que `{UPSTREAM_DISCOVERY_PATH}` existe e contém todos os artefatos upstream (F1-F9).

### Passo 1 — Carregar Documentos Base
Confirmar leitura de TODOS os artefatos Discovery-Level:
1-7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-*-DEFINITION.md` — Bloco B (F2-F7)
8. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD (F1)
9. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo (F8)
10. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — Matriz (F9)

### Passo 2 — Invocar Skills
Invocar skills de discovery, arquitetura e análise para síntese executiva.

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` — consolidação técnica high-level sumarizando descobertas do Bloco B para embasar ROM.

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-SPECS.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-SPECS.md`

```markdown
# DISCOVERY-LEVEL-SPECS.md
## Fase 10 — Bloco C: Catálogo, Matriz & Consolidação Discovery-Level

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-SPECS-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Consolidação Técnica) |
| **Data** | {DATA_ATUAL} |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Referenciados (Bloco B):**
- F1: [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md)
- F2-F7: [`DISCOVERY-LEVEL-*-DEFINITION.md`] — Bloco B (6 disciplinas)
- F8: [`DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md`](DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md)
- F9: [`DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md`](DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md)

---

## 1. Sumário Técnico Executivo
- 1-2 parágrafos consolidando: nome do produto, arquitetura, infraestrutura, time, épicos e prazo macro

## 2. Consolidação por Disciplina
- **2.1 Arquitetura (F2):** Tabela de decisões-chave e principais riscos
- **2.2 Segurança (F3):** Tabela de decisões-chave (modelo, threat model, compliance, trust boundary)
- **2.3 Dados (F4):** Tabela de decisões-chave (banco, isolamento, cache, volumes)
- **2.4 DevOps/SRE (F5):** Tabela de decisões-chave (CI/CD, IaC, observabilidade, SLOs)
- **2.5 Testes (F6):** Tabela de decisões-chave (pirâmide, cobertura, quality gates)
- **2.6 Infra/Cloud (F7):** Tabela de decisões-chave (provedor, orquestração, custos, DR)

## 3. Matriz de Consistência Cross-Disciplina
- Tabela: Par | Verificação | Status (✅)
- 10+ pares de disciplinas verificando consistência entre si

## 4. Premissas para ROM
- Lista numerada de premissas que embasam a estimativa

## 5. Riscos Consolidados (Top 5)
- Tabela: # | Risco | Disciplina | Severidade

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Consolidação técnica Discovery-Level | Tech Lead |
```

### Passo 4 — Validação Pós-Geração

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `discovery-process` | Síntese das descobertas técnicas em um documento executivo consolidado | Discovery |
| 2 | `senior-architect` | Consolidação arquitetural cross-discipline | Arquitetura |
| 3 | `gap-analysis` | Identificação de gaps e inconsistências entre disciplinas | Análise |
| 4 | `documentation-writer` | Documento executivo para o comitê | Documentação |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — F10 Bloco C Discovery-Level | Time de Arquitetura |
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
| 9 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` | Matriz (F9) |
| 10 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 11 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 12 | `{PROJECT-STACK}` | Stack tecnológica (baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`) |

---

🤖 *Upstream Architecture Discovery — Fase 10*
