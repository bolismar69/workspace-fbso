# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ROM-ESTIMATE
## Contexto
> 📐 **Discovery-Level:** Consolidação da estimativa ROM +-50% para o comitê de governança.

Este prompt gera `DISCOVERY-LEVEL-ROM-ESTIMATE.md` — documento que consolida a estimativa Rough Order of Magnitude com faixa de +-50%. Inclui matriz de esforço por solução, premissas, riscos e faixa de valores. Este é o output final do Discovery para decisão GO/NO-GO.

**Papel no Bloco D (Estimativa & ROM):** Fase 11 de 1. Consome todos os artefatos anteriores.

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
10. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` — Consolidação Técnica (F10)

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Documentos brutos adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |

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
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP` — skills do time (se fornecidos; embasam estimativa de esforço)
- `PROJECT-TEAM-CAPACITY` — capacidade do time (se fornecida; impacta faixa ROM)
Validar que `{UPSTREAM_DISCOVERY_PATH}` existe e contém os artefatos upstream (F1-F10).

### Passo 1 — Carregar Documentos Base
Confirmar leitura de TODOS os artefatos Discovery-Level:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — Dados (F4)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — DevOps/SRE (F5)
6. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — Testes (F6)
7. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Infra/Cloud (F7)
8. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo (F8)
9. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — Matriz (F9)
10. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` — SPECS (F10)

### Passo 2 — Invocar Skills
Invocar skills de arquitetura, produto e análise para consolidar estimativas.

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ROM-ESTIMATE.md` com:
1. **Matriz de Esforço por Solução** — solução, complexidade, esforço estimado (homem-mês), faixa (min-máx)
2. **Premissas** — assumptions que embasam a estimativa
3. **Riscos e Mitigações** — riscos identificados com impacto na estimativa
4. **Faixa de Valores** — ROM consolidado (min, provável, máx) com +-50%
5. **Recomendação Técnica** — parecer do time de arquitetura para o comitê

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-ROM-ESTIMATE.md`

```markdown
# DISCOVERY-LEVEL-ROM-ESTIMATE.md
## Fase 11 — Bloco D: Estimativa & ROM

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-ROM-ESTIMATE-v1.0 |
| **Versão** | 1.0 — ROM ±50% (Discovery-Level) |
| **Data** | {DATA_ATUAL} |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Referenciados:**
- F1-F10: Todas as fases do Upstream Architecture Discovery
- [`DISCOVERY-LEVEL-SPECS.md`](DISCOVERY-LEVEL-SPECS.md) — Consolidação Técnica (F10)

> **Premissa de conversão:** 1 dia = 8 horas · 1 semana = 40 horas · Duração considera 5 dias úteis/semana.

---

## 1. Estimativa ROM Consolidada (±50%)
- **1.1 Esforço de Discovery (Sprint 0 — Análise):** Tabela: Disciplina | Atividades | Range/Semanas | Range/Horas | Responsável
- **1.2 Esforço de Implementação — Por Épico:** Tabela: Épico | Func. | Complexidade | Time Alocado | Range/Semanas | Range/Horas
- **1.3 Esforço de Infraestrutura e DevOps (Paralelo à Implementação):** Tabela: Atividade | Range/Semanas | Range/Horas | Responsável
- **1.4 Esforço de Testes e Qualidade (Paralelo à Implementação):** Tabela: Atividade | Range/Semanas | Range/Horas | Responsável

## 2. Estimativa ROM Total
- **2.1 Consolidação por Bloco:** Tabela: Bloco | Range/Semanas | Range/Horas | % do Total | Paralelizável?
- **2.2 Cenários ROM (±50%):** Tabela: Cenário | Range/Semanas | Range/Horas | Time | Custo Infra
- **2.3 Caminho Crítico (Cenário Provável):** Tabela: Fase | Range/Semanas | Range/Horas | Time Alocado
- **2.4 Distribuição do Esforço por Disciplina:** Tabela: Disciplina | Range/Semanas | Range/Horas | % do Total

## 3. Premissas da Estimativa
- Tabela: # | Premissa | Impacto se inválida

## 4. Riscos da Estimativa
- Tabela: ID | Risco | Impacto em Semanas | Impacto em Horas | Mitigação

## 5. Recomendação Técnica para o Comitê de Governança
- **Parecer do Time de Arquitetura:** ✅ O projeto é tecnicamente viável (ou não)
- **Fundamentação:** 5 pontos justificando a recomendação
- **Recomendação:** GO ✅ ou NO-GO ❌

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: ROM ±50% Discovery-Level | Tech Lead / Solution Architect |
```

### Passo 4 — Validação Pós-Geração

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `product-discovery` | Mapear premissas e validar cenários de estimativa | Discovery |
| 2 | `senior-architect` | Estimativa de esforço técnico | Arquitetura |
| 3 | `cloud-architect` | Custos de infra | Cloud |
| 4 | `senior-devops` | Esforço DevOps/SRE | DevOps |
| 5 | `gap-analysis` | Análise de riscos | Análise |
| 6 | `documentation-writer` | Documento executivo | Documentação |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — F11 Bloco D Discovery-Level | Time de Arquitetura |
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
| 9 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` | Matriz Solução×Disciplina (F9) |
| 10 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` | Consolidação Técnica (F10) |
| 11 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 12 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 13 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (se fornecidos; embasam estimativa de esforço) |
| 14 | `{PROJECT-TEAM-CAPACITY}` | Capacidade do time (se fornecida; impacta faixa ROM) |

---

🤖 *Upstream Architecture Discovery — Fase 11*
