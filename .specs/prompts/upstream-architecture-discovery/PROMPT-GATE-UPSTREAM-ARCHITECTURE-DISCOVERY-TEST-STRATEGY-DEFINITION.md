# PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Estratégia de Testes** para o artefato `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md`. Verifica se a estratégia de testes do projeto está completa, cobre a pirâmide de testes, automação, performance, segurança e quality gates, e está alinhada com a arquitetura, segurança e DevOps.

**Princípio fundamental:** Toda solução deve ter sua estratégia de testes definida na pirâmide. Nenhum código pode chegar a produção sem quality gates objetivos e mensuráveis.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — artefato auditado (F6)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7)
6. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints globais

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

**Arquivos gerados pelo GENERATE:** `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md`

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP`, `PROJECT-TEAM-CAPACITY`, `PROJECT-STACK` — se fornecidos
Validar que o artefato auditado `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` existe.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` — artefato auditado (F6)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
5. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7)
6. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints globais

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Pirâmide de Testes
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Níveis definidos | Unitários, integração, E2E, aceitação — todos documentados |
| 1.2 | % de cobertura | Cobertura esperada por nível e por camada |
| 1.3 | Frameworks especificados | Frameworks por linguagem com justificativa |

#### Dimensão 2: Automação e Pipeline
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Testes no CI/CD | Pipeline de testes integrado ao CI/CD |
| 2.2 | Paralelismo | Estratégia de execução paralela definida |
| 2.3 | Ambientes de teste | Isolamento, dados anonimizados, provisionamento |

#### Dimensão 3: Testes Especializados
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Testes de Performance | Carga, stress, soak definidos com thresholds |
| 3.2 | Testes de Segurança | SAST, DAST, SCA, penetration testing |
| 3.3 | Alinhamento com SLOs | Thresholds de performance alinhados com DEVOPS-SRE |

#### Dimensão 4: Quality Gates
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Gatilhos definidos | Critérios mínimos para cada nível |
| 4.2 | Métricas objetivas | Coverage, mutation testing, falhas máximas |
| 4.3 | Consistência com SECURITY | Testes de segurança alinhados com SECURITY-DEFINITION |

#### Dimensão 5: Alinhamento com Stack Corporativa e Time

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Skills mapeados | Skills necessários para esta disciplina estão documentados no artefato |
| 5.2 | Capacidade estimada | Capacidade do time está dimensionada proporcionalmente à complexidade |
| 5.3 | Stack corporativa | Tecnologias propostas constam no `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |
| 5.4 | Tecnologias adicionais | Tecnologias fora do padrão corporativo têm justificativa técnica documentada e aprovada |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE ESTRATÉGIA DE TESTES: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-TEST-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que a Architecture/Security/DevOps determinava:** [Descrever a referência]
  - **Impacto:** [O risco de qualidade — cobertura insuficiente, testes sem automação, security testing ausente]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a definição de estratégia de testes, por favor, responda:
1. Quanto ao **[ID-CONFLITO-TEST-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A ESTRATÉGIA DE TESTES ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE ESTRATÉGIA DE TESTES: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md` gerado conforme Architecture Definition, Security Definition e DevOps SRE.
- **AUDITORIA DA IA:** Pirâmide definida com % de cobertura. Ferramentas especificadas por solução e linguagem. Quality gates documentados. Testes de segurança alinhados com SECURITY-DEFINITION. Testes de performance alinhados com DEVOPS-SRE (SLOs).
- **DIRETRIZ:** Peço que leia a definição de estratégia de testes para verificar se a pirâmide, automação e quality gates atendem às expectativas de qualidade.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A definição de estratégia de testes está em compliance e reflete corretamente as necessidades de qualidade do projeto?
2. Deseja enviar mais documentos/arquivos para enriquecer a estratégia de testes?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Este documento não possui uma próxima fase de technical-definition (Bloco B completo). O fluxo segue para o Bloco C (F13-F17 — Solutions Catalog, Matrix, Stack Matrix). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-qa` | Validar estratégia de testes | QA |
| 2 | `test-strategy-design` | Validar pirâmide e automação | Estratégia |
| 3 | `performance-testing-review-ai-review` | Validar testes de performance | Performance |
| 4 | `security-testing` | Validar testes de segurança | Segurança |
| 5 | `gap-analysis` | Identificar gaps de teste | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação da definição de estratégia de testes | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` | Artefato auditado (F6) |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` | Definição DevOps/SRE (F5) |
| 5 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` | Definição Infra/Cloud (F7) |
| 6 | `{ARCHITECTURE_GLOBAL}/` | ADRs, blueprints globais |
| 7 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 8 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 9 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| 10 | `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| 11 | `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
