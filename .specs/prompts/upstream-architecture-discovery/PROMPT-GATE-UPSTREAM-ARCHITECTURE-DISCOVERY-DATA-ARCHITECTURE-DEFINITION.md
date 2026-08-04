# PROMPT-GATE-UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Arquitetura de Dados** para o artefato `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md`. Verifica se a arquitetura de dados do projeto está completa, consistente com as definições de arquitetura e segurança, e cobre armazenamento, pipelines, governança e tecnologias.

**Princípio fundamental:** Todo dado produzido ou consumido por uma solução deve ter sua origem, armazenamento, pipeline e governança documentados. Nenhum fluxo de dados pode ficar sem rastreabilidade.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — artefato auditado (F4)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints, data standards) |
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

**Arquivos gerados pelo GENERATE:** `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md`

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP`, `PROJECT-TEAM-CAPACITY`, `PROJECT-STACK` — se fornecidos
Validar que o artefato auditado `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` existe.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` — artefato auditado (F4)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
5. `{ARCHITECTURE_GLOBAL}/` — ADRs globais e data standards

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Dados
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Entidades catalogadas | Todas as entidades de negócio mapeadas no dicionário de dados |
| 1.2 | Armazenamento por solução | Toda solução tem storage strategy definida (SQL/NoSQL/Cache/File) |
| 1.3 | Pipelines documentados | ETL/ELT e streaming com origem, destino e schedule |
| 1.4 | Integrações inter-banco | Data services e APIs de dados documentados |

#### Dimensão 2: Completude Técnica
| # | Verificação | Critério |
|---|---|---|
| 2.1 | ERD presente | Diagrama entidade-relacionamento com entidades e relacionamentos |
| 2.2 | Dicionário de dados | Tipos, tamanhos, restrições, descrições |
| 2.3 | Estratégia On-Prem vs Cloud | Comparação e justificativa documentadas |
| 2.4 | Tecnologias especificadas | SGBDs, ferramentas ETL, plataformas de streaming |

#### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com ARCHITECTURE-DEFINITION | Containers que consomem dados têm storage mapeado |
| 3.2 | Alinhamento com SECURITY-DEFINITION | Criptografia em repouso, IAM de dados, privacidade |

#### Dimensão 4: Governança
| # | Verificação | Critério |
|---|---|---|
| 4.1 | Data Governance | Qualidade, linhagem e catálogo definidos |
| 4.2 | Privacidade | LGPD/GDPR — anonimização, retenção, purge |
| 4.3 | Rastreabilidade | Linhagem de dados documentada |

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

#### 📊 RELATÓRIO DE AUDITORIA DE ARQUITETURA DE DADOS: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-DATA-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que a Architecture/Security/PRD Definition determinava:** [Descrever a referência]
  - **Impacto:** [O risco de dados sem governança ou pipeline quebrado]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a definição de arquitetura de dados, por favor, responda:
1. Quanto ao **[ID-CONFLITO-DATA-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A ARQUITETURA DE DADOS ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE ARQUITETURA DE DADOS: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `UPSTREAM-ARCHITECTURE-DISCOVERY-DATA-ARCHITECTURE-DEFINITION.md` gerado conforme Architecture Definition, Security Definition e PRD Definition.
- **AUDITORIA DA IA:** Cobertura completa verificada. ERD presente com entidades de todas as soluções. Storage strategy definida. Pipelines documentados. Data Governance coberto. Consistente com Architecture (containers) e Security (criptografia, IAM de dados).
- **DIRETRIZ:** Peço que leia a definição de arquitetura de dados para verificar se os modelos, pipelines e estratégias de armazenamento refletem as necessidades do projeto.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A definição de arquitetura de dados está em compliance e reflete corretamente como os dados são armazenados e processados?
2. Deseja enviar mais documentos/arquivos para enriquecer a arquitetura de dados?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e Fase 10 (DEVOPS-SRE-DEFINITION). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-data-engineer` | Validar arquitetura de dados | Engenharia de Dados |
| 2 | `data-engineer` | Validar pipelines e armazenamento | Engenharia de Dados |
| 3 | `data-modeling` | Validar modelos e ERD | Modelagem |
| 4 | `database-architect` | Validar estratégia de SGBDs | Banco de Dados |
| 5 | `data-quality-frameworks` | Validar governança | Governance |
| 6 | `gap-analysis` | Identificar gaps de dados | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: gate de validação da definição de arquitetura de dados | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md` | Artefato auditado (F4) |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |
| 7 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (obter e validar com usuário) |
| 8 | `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time (obter e validar com usuário) |
| 9 | `{PROJECT-STACK}` | Stack tecnológica. Baseline: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
