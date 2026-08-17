# PROMPT-GATE-470-ARCHITECTURE-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Arquitetura** para o artefato `470-ARCHITECTURE-DEFINITION.md`. Verifica se a arquitetura do projeto está completa, consistente e cobre todas as soluções do catálogo.

**Princípio fundamental:** Toda solução do catálogo deve aparecer nos diagramas C4 e na matriz de integração. Nenhuma solução pode ficar "desconectada" da arquitetura.

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
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

**Arquivos gerados pelo GENERATE:** `470-ARCHITECTURE-DEFINITION.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `470-ARCHITECTURE-DEFINITION.md`, Catálogo de Soluções, PRD Definition, ADRs globais.

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
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-470-ARCHITECTURE-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A ARQUITETURA ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE ARQUITETURA: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `470-ARCHITECTURE-DEFINITION.md` gerado conforme catálogo e ADRs.
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

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
