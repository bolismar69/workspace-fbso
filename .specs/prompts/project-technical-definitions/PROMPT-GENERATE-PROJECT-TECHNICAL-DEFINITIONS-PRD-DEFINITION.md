# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` 🆕 — a **baseline de requisitos de produto no nível do projeto**. Este documento é a referência normativa que cada PRD.md de solução individual vai especializar, garantindo que todos os times partam da mesma interpretação dos requisitos de negócio.

**Relação com outros artefatos:** O PRD Definition está para os PRDs de solução assim como o GLOBAL-SECURITY.md está para os SECURITY.md de solução — é a referência que cada solução especializa, não substitui.

**Inputs upstream:** `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` (Fase 2) + documentos de negócio (Charter, BRD, Epics, Features, User Stories, RTM) + TECHNICAL-PLAN.md (referência).

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler documentos de negócio (Charter, BRD, Epics, Features, User Stories, RTM) + Catálogo de Soluções (Fase 2) + TECHNICAL-PLAN.md (referência).

### Passo 2 — Invocar Skills Especializadas
Invocar skills de produto, requisitos e análise de negócio para extrair e consolidar todos os requisitos, mapear para soluções e definir MVP global.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` com:
1. **Visão do Produto** — Sistema completo como produto unificado
2. **Matriz Requisito → Solução** — Cada BR/Feature/US mapeado para solução(ões)
3. **MVP Global** — Entrega mínima viável cross-solution
4. **Requisitos Funcionais Cross-Solution** — Fluxos multi-solução
5. **Requisitos Não-Funcionais Globais** — Performance, disponibilidade, escalabilidade
6. **Restrições de Produto** — O que NÃO está no escopo
7. **Glossário de Domínio Unificado** — Termos canônicos para todos os times
8. **Critérios de Aceitação Cross-Solution** — Cenários end-to-end

### Passo 4 — Validação Pós-Geração
Verificar: todas as seções preenchidas, matriz de cobertura 100%, MVP definido, restrições documentadas.

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

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
