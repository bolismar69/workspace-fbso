# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP

## Contexto

Este prompt implementa o **Gate de Validação da Matriz de Skills** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md`. O agente validador verifica se a matriz de skills está completa, consistente e alinhada com os perfis documentados no `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md`.

**Princípio fundamental:** Todo perfil listado no TEAM-CAPACITY deve ter skills mapeadas no TEAM-SKILLS-MAP. Nenhum perfil pode ficar sem cobertura de competências.

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

**Arquivos gerados pelo GENERATE:** `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md` (artefato a validar), `PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md` (referência de perfis), `PRODUCT-BACKLOG-LIST.md` (F3) e `PRD-DEFINITION.md` (F4 — Bloco 0) para validar alinhamento com escopo do projeto.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Perfis
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Todos os perfis cobertos | Cada perfil do TEAM-CAPACITY tem entrada na matriz de skills |
| 1.2 | Sem perfis órfãos | Nenhuma skill mapeada para perfil inexistente no TEAM-CAPACITY |

#### Dimensão 2: Completude da Matriz
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Categorias obrigatórias | Linguagens, Frameworks, Bancos, Cloud, DevOps — todas preenchidas |
| 2.2 | Níveis de proficiência | ★☆☆ a ★★★ definidos para cada skill×perfil |
| 2.3 | Gap analysis | Gaps documentados com recomendações |

#### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhamento com escopo | Skills mapeadas cobrem tecnologias mencionadas no TECHNICAL-PLAN.md |
| 3.2 | Nomes consistentes | Tecnologias usam nomenclatura padronizada (ex: "PostgreSQL", não "Postgres") |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

Seu retorno para o usuário humano deve seguir estritamente uma das duas estruturas condicionais abaixo, dependendo do resultado da sua análise:

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS CONFLITOS OU DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE SKILLS: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-TM-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o documento base determinava:** [Descrever a referência do TEAM-CAPACITY ou TECHNICAL-PLAN]
  - **Impacto:** [O risco de alocação incorreta ou gap de competências]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir o documento, por favor, responda:
1. Quanto ao **[ID-CONFLITO-TM-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas e de resposta curta para sanar as dúvidas encontradas]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução interna para o orquestrador: O processo pausa aqui e aguarda as respostas do humano. Assim que o humano responder, todo este relatório + as respostas dele serão injetadas no PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md)*

---

### ✅ CENÁRIO B: SE O DOCUMENTO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE SKILLS: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP.md` gerado conforme as informações fornecidas.
- **AUDITORIA DA IA:** Documento pré-validado. Cobertura de perfis completa, matriz de skills preenchida, gaps documentados. Nenhum conflito conceitual encontrado pela IA.
- **DIRETRIZ:** Peço que leia o documento para verificar se o mesmo atende plenamente às suas necessidades de alocação técnica.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. O documento está em compliance e perfeitamente alinhado com o TEAM-CAPACITY e as necessidades técnicas do projeto?
2. Deseja enviar mais documentos/arquivos para enriquecer a matriz de skills?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se o usuário responder "Sim" para a Pergunta 1 e "Não" para as Perguntas 2 e 3, altere o status para [STATUS: COMPLIANCE] e destrave a Fase 6 (TEAM-CAPACITY). Se o usuário fornecer novos documentos ou inputs nas Perguntas 2 ou 3, acione imediatamente o fluxo de re-alimentação voltando ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `skill-audit` | Auditar skills contra TEAM-CAPACITY | Discovery |
| 2 | `gap-analysis` | Identificar lacunas na matriz | Análise |
| 3 | `team-composition-analysis` | Validar composição do time | People |
| 4 | `senior-architect` | Revisão de arquitetura de skills | Arquitetura |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no relatório.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da matriz de skills | Time de Arquitetura |
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário (alinhado ao roadmap de documentos de negócio) | Time de Arquitetura |
| 3.0 | 30/07/2026 | Atualização Bloco A (F5-F6): adicionados inputs PRODUCT-BACKLOG-LIST (F3) e PRD-DEFINITION (F4); atualizada referência de fase (F2→F6) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
