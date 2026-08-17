# PROMPT-GATE-560-SPECS-DEFINITION (F16)

## Contexto

Este prompt implementa o **Gate de Validação da Consolidação Técnica (F16)** para o artefato `560-SPECS-DEFINITION.md`. Verifica se a consolidação enxuta referencia 100% dos artefatos dos blocos anteriores, se os links markdown são válidos e se não há duplicação de conteúdo.

**Princípio fundamental:** O SPECS-DEFINITION é um sumário navegável, não uma baseline. Toda decisão técnica deve ter um `→ ver [ARTEFATO]` apontando para o documento original. Nenhum bloco de conteúdo deve duplicar o artefato referenciado.

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

**Arquivos gerados pelo GENERATE:** `560-SPECS-DEFINITION.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `560-SPECS-DEFINITION.md` e TODOS os artefatos referenciados (Blocos 0, A, B, C) para verificar existência e consistência das referências.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Cobertura de Artefatos
| # | Verificação | Critério |
|---|---|---|
| 1.1 | 100% artefatos referenciados | Cada artefato dos Blocos 0, A, B, C tem seção correspondente |
| 1.2 | Artefatos existem | Caminhos dos arquivos referenciados são válidos |
| 1.3 | Sem lacunas | Nenhuma disciplina do Bloco B sem referência |

#### Dimensão 2: Qualidade dos Sumários
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Sumários concisos | Cada seção tem ~1 parágrafo (não listas extensas) |
| 2.2 | Sem duplicação | Nenhum bloco de conteúdo copiado dos artefatos fonte |
| 2.3 | Links markdown válidos | Todo `→ ver [ARTEFATO]` é link clicável e válido |

#### Dimensão 3: Estrutura
| # | Verificação | Critério |
|---|---|---|
| 3.1 | 10 seções presentes | Todas as seções obrigatórias estão preenchidas |
| 3.2 | Ordem correta | Sequência: Convenções → API → Dados → Segurança → DevOps/SRE → Testes → Infra → Stacks → Decisões → Restrições |

### Passo 3 — Emitir Veredito

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE CONSOLIDAÇÃO TÉCNICA: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-SPC-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o artefato fonte determinava:** [Descrever a referência]
  - **Impacto:** [O risco de referência quebrada ou duplicação]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a consolidação técnica, por favor, responda:
1. Quanto ao **[ID-CONFLITO-SPC-01]**, qual é a correção a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-560-SPECS-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A CONSOLIDAÇÃO ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE CONSOLIDAÇÃO TÉCNICA: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `560-SPECS-DEFINITION.md` gerado como consolidação técnica enxuta.
- **AUDITORIA DA IA:** 100% dos artefatos dos Blocos 0, A, B e C referenciados. Links markdown válidos. Nenhuma duplicação de conteúdo detectada. Sumários concisos conforme especificado.
- **DIRETRIZ:** Peço que leia a consolidação técnica para verificar se os sumários representam corretamente as decisões dos artefatos fonte.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A consolidação técnica está em compliance e os sumários estão corretos?
2. Deseja enviar mais documentos/arquivos para enriquecer a consolidação?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e F17 (MILESTONES). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `reference-builder` | Validar referências cruzadas | Mapeamento |
| 2 | `gap-analysis` | Identificar lacunas de cobertura | Análise |
| 3 | `senior-architect` | Validar completude arquitetural | Arquitetura |
| 4 | `code-review` | Verificar duplicação de conteúdo | Qualidade |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da baseline de especificações | Time de Arquitetura |
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |
| 3.0 | 30/07/2026 | Reformulação: validação de consolidação enxuta — 100% artefatos, links markdown, sem duplicação | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
