# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` — a **baseline de requisitos de produto no nível do projeto**, focada no **PRD de Negócio**. Este documento é a referência normativa que cada PRD.md de solução individual vai especializar, garantindo que todos os times partam da mesma interpretação dos requisitos de negócio.

**Relação com outros artefatos:** O PRD Definition está para os PRDs de solução assim como o GLOBAL-SECURITY.md está para os SECURITY.md de solução — é a referência que cada solução especializa, não substitui.

> **IMPORTANTE — PRD DE NEGÓCIO CONGELADO APÓS BARREIRA 0:**
> Este documento é **CONGELADO** após a Barreira 0. Não será reaberto para alterações de escopo. O `SPECS-DEFINITION` (Fase 16 — Bloco C) fará a consolidação técnica detalhada a partir deste PRD.

**Inputs upstream (Fase 4 — Bloco 0 — Product Definition & Product Backlog & PRD):**
- **Inputs Globais do Roadmap** (todas as 8 variáveis)
- **INTAKE-LOG.md (F1):** Log de demandas iniciais e triagem
- **DOR-ASSESSMENT.md (F2):** Definition of Ready — avaliação de maturidade
- **PRODUCT-BACKLOG-LIST.md (F3):** Backlog priorizado do produto
- **Documentos de Negócio:** Charter, BRD, Epics, Features, User Stories, RTM

**Blocos posteriores (NÃO são inputs deste prompt — vêm após o Bloco 0):**
- Bloco A (F5-F6) — People & Solutions: Usa este PRD como referência
- Bloco B (F7-F15) — Architecture, Security, Specs: Usa este PRD como referência
- Bloco C (F16+) — Technical Consolidation: SPECS-DEFINITION detalha tecnicamente

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

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler documentos de negócio (Charter, BRD, Epics, Features, User Stories, RTM) + Catálogo de Soluções (Fase 2) + TECHNICAL-PLAN.md (referência).

### Passo 2 — Invocar Skills Especializadas
Invocar skills de produto, requisitos e análise de negócio para extrair e consolidar todos os requisitos, mapear para soluções e definir MVP global.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` com:
1. **Visão do Produto** — Sistema completo como produto unificado, com foco no valor de negócio
2. **Matriz Requisito → Entrega** — Cada BR/Feature/US mapeado para entregas de negócio
3. **MVP Global** — Entrega mínima viável com critérios de negócio mensuráveis
4. **Requisitos Funcionais de Negócio** — Funcionalidades essenciais descritas em linguagem de negócio
5. **Restrições de Produto** — O que NÃO está no escopo (decisões de negócio)
6. **Glossário de Domínio Unificado** — Termos canônicos para todos os times

> Requisitos técnicos detalhados (NFRs, cenários de aceitação técnicos, fluxos cross-solution detalhados) pertencem ao `SPECS-DEFINITION.md` (Fase 16 — Bloco C).

### Passo 4 — Validação Pós-Geração
Verificar: todas as seções preenchidas, matriz de cobertura 100% do backlog, MVP definido, restrições documentadas.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `prd` | Gerar PRD no nível do projeto | Product |
| 2 | `prd-development` | Desenvolver PRD com todas as seções | Product |
| 3 | `product-manager` | Visão de produto e MVP global | Product |
| 4 | `product-strategist` | Estratégia de produto cross-solution | Product |
| 5 | `requirements-engineering` | Engenharia de requisitos | Requirements |
| 6 | `requirements-elicitation` | Extrair requisitos dos docs de negócio | Requirements |
| 7 | `requirements-prioritization` | Priorizar requisitos (MoSCoW) | Requirements |
| 8 | `business-analyst` | Análise de negócio para traduzir BRD→PRD | Business |
| 9 | `business-analysis-planning` | Planejamento da análise | Business |
| 10 | `user-stories` | Referência às user stories existentes | US |
| 11 | `acceptance-criteria` | Critérios de aceitação cross-solution | US |
| 12 | `stakeholder-analysis` | Alinhamento com stakeholders | Stakeholder |
| 13 | `domain-analysis` | Análise de domínio unificado | Domínio |
| 14 | `roadmap-planning` | Alinhamento com roadmap de negócio | Roadmap |
| 15 | `documentation-writer` | Redigir o PRD Definition consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador do PRD Definition | Time de Arquitetura |
| 2.0 | 30/07/2026 | Migração para Bloco 0 (F4): PRD de Negócio congelado após Barreira 0; inputs alterados para INTAKE-LOG/DOR-ASSESSMENT/BACKLOG; seções reorientadas para foco em negócio | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
