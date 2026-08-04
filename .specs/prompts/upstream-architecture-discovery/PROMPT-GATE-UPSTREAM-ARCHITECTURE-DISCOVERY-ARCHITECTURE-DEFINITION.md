# PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Arquitetura** para o artefato `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md`. Verifica se a arquitetura do projeto está completa, consistente e cobre todas as soluções do catálogo.

**Princípio fundamental:** Toda solução do catálogo deve aparecer nos diagramas C4 e na matriz de integração. Nenhuma solução pode ficar "desconectada" da arquitetura.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — artefato auditado (F2)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
4. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints globais

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

**Arquivos gerados pelo GENERATE:** `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md`

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP`, `PROJECT-TEAM-CAPACITY`, `PROJECT-STACK` — se fornecidos
Validar que o artefato auditado `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` existe.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — artefato auditado (F2)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
4. `{ARCHITECTURE_GLOBAL}/` — ADRs globais

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Soluções no C4 L1 | Toda solução do catálogo aparece no System Context |
| 1.2 | Soluções no C4 L2 | Containers mapeados para todas as soluções |
| 1.3 | Integrações documentadas | Cada par origem→destino tem protocolo e autenticação |

#### Dimensão 2: Completude Técnica
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Diagramas C4 presentes | L1 e L2 com sintaxe C4 correta |
| 2.2 | Topologia de deploy | Ambientes e infra definidos |
| 2.3 | Comunicação síncrona/assíncrona | Estratégia documentada |
| 2.4 | ADRs de integração | ≥3 ADRs cross-solution |

#### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com blueprints | Soluções seguem blueprints da pasta architecture/ |
| 3.2 | Consistência com PRD Definition | Funcionalidades cross-solution têm integração definida |

#### Dimensão 4: Alinhamento com Stack Corporativa e Time

| # | Verificação | Critério |
|---|---|---|
| 4.1 | Skills mapeados | Skills necessários para esta disciplina estão documentados no artefato |
| 4.2 | Capacidade estimada | Capacidade do time está dimensionada proporcionalmente à complexidade |
| 4.3 | Stack corporativa | Tecnologias propostas constam no `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |
| 4.4 | Tecnologias adicionais | Tecnologias fora do padrão corporativo têm justificativa técnica documentada e aprovada |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE ARQUITETURA: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-ARCH-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o catálogo/ADR/PRD determinava:** [Descrever a referência]
  - **Impacto:** [O risco de solução desconectada ou integração quebrada]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a definição de arquitetura, por favor, responda:
1. Quanto ao **[ID-CONFLITO-ARCH-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A ARQUITETURA ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE ARQUITETURA: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` gerado conforme catálogo e ADRs.
- **AUDITORIA DA IA:** Cobertura completa verificada. Diagramas C4 L1/L2 presentes com todas as soluções. Matriz de integração documentada. ADRs cross-solution registrados. Nenhuma solução desconectada.
- **DIRETRIZ:** Peço que leia a definição de arquitetura para verificar se os diagramas e integrações refletem a arquitetura esperada.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A definição de arquitetura está em compliance e reflete corretamente como as soluções se integram?
2. Deseja enviar mais documentos/arquivos para enriquecer a arquitetura?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e Fase 8 (SECURITY-DEFINITION). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `c4-architecture-c4-architecture` | Validar diagramas C4 | C4 |
| 2 | `architecture-patterns` | Validar padrões arquiteturais | Arquitetura |
| 3 | `architect-review` | Revisão de arquitetura | Arquitetura |
| 4 | `gap-analysis` | Identificar soluções desconectadas | Análise |
| 5 | `senior-architect` | Validação sênior da arquitetura | Arquitetura |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da definição de arquitetura | Time de Arquitetura |
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |
| 3.0 | 30/07/2026 | Atualização F3→F7: orquestração redirecionada para Fase 8 (SECURITY-DEFINITION) | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Artefato auditado (F2) |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) — referência de funcionalidades |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` | Catálogo de Soluções (F8) — cobertura de soluções |
| 4 | `{ARCHITECTURE_GLOBAL}/` | ADRs, blueprints globais |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 7 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| 8 | `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| 9 | `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
