# PROMPT-GATE-590-TECHNICAL-DISCOVERY

## Contexto

Este prompt implementa o **Gate de Validação dos Contratos Técnicos por Sprint** para os artefatos gerados em `technical-discovery/590-ciclo-NNN/`. Verifica se 100% das tarefas da sprint possuem contratos, se cada contrato referencia User Stories e artefatos base, e se há consistência com SPECS-DEFINITION.

**Princípio fundamental:** Toda tarefa T-NNN da sprint deve ter contratos de API, Dados, Segurança e SRE. Nenhuma tarefa pode ser entregue sem acordo formal entre as disciplinas técnicas.

**Natureza Iterativa:** Ao final da validação bem-sucedida de uma sprint, o gate deve perguntar se o usuário deseja iniciar o Discovery da próxima sprint.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{SPRINT_NUMBER}` | Número da sprint alvo (ex: 01, 02, 03) |
| `{SPRINT_TASKS}` | Lista de TASK-ID da sprint conforme PACKAGE-BACKLOG |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

**Arquivos gerados pelo GENERATE:** `technical-discovery/sprint-{SPRINT_NUMBER}/{CONTRACTS-API, CONTRACTS-DATA, CONTRACTS-SECURITY, CONTRACTS-SRE, DEFINITION-INCREMENTS}-sprint-{SPRINT_NUMBER}.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler todos os 5 contratos da sprint atual, PACKAGE-BACKLOG (F18), SPECS-DEFINITION (F16), artefatos base do Bloco B (ARCHITECTURE, SECURITY, DATA-ARCH, DEVOPS-SRE).

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Tarefas
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Todas as tarefas cobertas | 100% das T-NNN da sprint têm contratos |
| 1.2 | Disciplinas cobertas | Cada tarefa tem contratos de API, Data, Security e SRE |
| 1.3 | Tarefas sem contrato | Nenhuma tarefa da sprint deve ficar sem contrato |

#### Dimensão 2: Rastreabilidade
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Referência a US | Cada contrato referencia ao menos uma US-ID |
| 2.2 | Referência a artefato base | Cada contrato referencia seu artefato base (ARCHITECTURE, DATA-ARCH, SECURITY, DEVOPS-SRE) |
| 2.3 | Referência cruzada entre contratos | Contratos de API e Data referenciam Security; Security referencia API/Data |

#### Dimensão 3: Consistência Técnica
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Consistência com SPECS-DEFINITION | Contratos respeitam a baseline de especificações |
| 3.2 | Consistência com artefato base | Contrato não contradiz o artefato base correspondente |
| 3.3 | Consistência entre contratos | API, Data, Security e SRE não têm contradições entre si |

#### Dimensão 4: Completude dos Contratos
| # | Verificação | Critério |
|---|---|---|
| 4.1 | CONTRACTS-API | Endpoints, request/response, auth, rate limits documentados |
| 4.2 | CONTRACTS-DATA | Schemas, migrations, queries, índices documentados |
| 4.3 | CONTRACTS-SECURITY | Regras IAM, validações, threat model documentados |
| 4.4 | CONTRACTS-SRE | SLOs, dashboards, alertas, runbooks documentados |
| 4.5 | DEFINITION-INCREMENTS | Atualizações retroativas e lições aprendidas documentadas |

### Passo 3 — Emitir Veredito e Pergunta Iterativa

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE DISCOVERY — Sprint {SPRINT_NUMBER}: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-DISCOVERY-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o artefato base/PACKAGE-BACKLOG determinava:** [Descrever a referência]
  - **Impacto:** [O risco de inconsistência entre contratos ou lacuna técnica]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir os contratos da sprint, por favor, responda:
1. Quanto ao **[ID-CONFLITO-DISCOVERY-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-590-TECHNICAL-DISCOVERY.md)*

---

### ✅ CENÁRIO B: SE OS CONTRATOS DA SPRINT ESTIVEREM 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE DISCOVERY — Sprint {SPRINT_NUMBER}: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **SPRINT:** {SPRINT_NUMBER}
- **DOCUMENTOS:**
  - `technical-discovery/sprint-{SPRINT_NUMBER}/CONTRACTS-API-sprint-{SPRINT_NUMBER}.md`
  - `technical-discovery/sprint-{SPRINT_NUMBER}/CONTRACTS-DATA-sprint-{SPRINT_NUMBER}.md`
  - `technical-discovery/sprint-{SPRINT_NUMBER}/CONTRACTS-SECURITY-sprint-{SPRINT_NUMBER}.md`
  - `technical-discovery/sprint-{SPRINT_NUMBER}/CONTRACTS-SRE-sprint-{SPRINT_NUMBER}.md`
  - `technical-discovery/sprint-{SPRINT_NUMBER}/DEFINITION-INCREMENTS-sprint-{SPRINT_NUMBER}.md`
- **AUDITORIA DA IA:** 100% das tarefas com contratos. Cada contrato referencia US e artefato base. Consistência com SPECS-DEFINITION validada. Nenhuma contradição entre contratos detectada.
- **DIRETRIZ:** Peço que leia os contratos da Sprint {SPRINT_NUMBER} para verificar se os acordos técnicos estão adequados.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. Os contratos técnicos da Sprint {SPRINT_NUMBER} estão em compliance e os acordos entre disciplinas estão adequados?
2. Deseja enviar mais documentos/arquivos para enriquecer os contratos?
3. **Deseja iniciar o Discovery da Sprint {NEXT_SPRINT_NUMBER}?** (Sim / Não — Se Sim, o GENERATE-TECHNICAL-DISCOVERY será executado para a próxima sprint. Se Não, o Discovery Técnico é encerrado e o pipeline avança para Barreira D → EXECUTION-HISTORY.)

*(Instrução de Orquestração: Se "Sim, Não" + "Sim" para próxima sprint → retrocede ao GENERATE para Sprint {NEXT_SPRINT_NUMBER}. Se "Sim, Não" + "Não" → [STATUS: COMPLIANCE] e Barreira D. Se novos inputs → retrocede ao GENERATE com inputs atualizados.)*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `api-designer` | Validar contratos de API | API |
| 2 | `data-modeling` | Validar contratos de dados | Dados |
| 3 | `security-auditor` | Validar contratos de segurança | Segurança |
| 4 | `sre-engineer` | Validar contratos de SRE | SRE |
| 5 | `gap-analysis` | Identificar tarefas sem contratos ou contratos incompletos | Análise |
| 6 | `senior-architect` | Supervisão de consistência cross-contrato | Arquitetura |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação iterativa dos contratos técnicos por sprint | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
