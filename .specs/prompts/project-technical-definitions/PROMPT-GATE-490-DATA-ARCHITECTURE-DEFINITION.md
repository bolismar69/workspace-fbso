# PROMPT-GATE-490-DATA-ARCHITECTURE-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Arquitetura de Dados** para o artefato `490-DATA-ARCHITECTURE-DEFINITION.md`. Verifica se a arquitetura de dados do projeto está completa, consistente com as definições de arquitetura e segurança, e cobre armazenamento, pipelines, governança e tecnologias.

**Princípio fundamental:** Todo dado produzido ou consumido por uma solução deve ter sua origem, armazenamento, pipeline e governança documentados. Nenhum fluxo de dados pode ficar sem rastreabilidade.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints, data standards) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

**Arquivos gerados pelo GENERATE:** `490-DATA-ARCHITECTURE-DEFINITION.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `490-DATA-ARCHITECTURE-DEFINITION.md`, Architecture Definition (F7), Security Definition (F8), PRD Definition (F4), ADRs globais e data standards.

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
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-490-DATA-ARCHITECTURE-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A ARQUITETURA DE DADOS ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE ARQUITETURA DE DADOS: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `490-DATA-ARCHITECTURE-DEFINITION.md` gerado conforme Architecture Definition, Security Definition e PRD Definition.
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

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
