# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-PRD

## Contexto

> 📐 **Discovery-Level / Upstream Architecture / Análise de Viabilidade:** Este prompt gera a versão Discovery-Level (high-level) do PRD para análise de viabilidade e estimativa ROM 50%. Para a versão detalhada, usar `PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md`.

Este prompt gera o artefato `DISCOVERY-LEVEL-PRD.md` 🆕 — o **PRD Discovery-Level** baseado nos Épicos definidos pelo Negócio. Diferente do PRD detalhado (F4 do tech-defs), este documento foca na visão macro do produto para embasar a análise de viabilidade técnica.

**Papel no Bloco 0 (Product Definition Discovery-Level):** Fase 1 de 1. Este artefato é o ponto de partida do Discovery Técnico. Ele consolida a visão do produto a partir dos Épicos e alimenta todas as 6 disciplinas do Bloco B.

**Inputs upstream:** Documentos de negócio — Project Charter, BRD, Épicos (índice + arquivos individuais em `epics/*.md`).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{UPSTREAM_DISCOVERY_PATH}` | Caminho da pasta upstream-architecture-discovery |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Documentos brutos adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Prompts auxiliares |

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
### Passo 1 — Carregar Documentos Base
Ler TODOS os documentos de negócio: Project Charter, BRD, Épicos (índice + `epics/*.md`).
### Passo 2 — Invocar Skills
Invocar skills de produto, análise de negócio e arquitetura para extrair visão macro.
### Passo 3 — Gerar o Artefato
Gerar `DISCOVERY-LEVEL-PRD.md` com:
1. **Visão do Produto (Resumo)** — 1-2 parágrafos sobre o produto
2. **Épicos × Soluções (Macro)** — Matriz de quais soluções cada épico demanda
3. **MVP Macro** — O que é essencial para o primeiro delivery
4. **Restrições Conhecidas** — O que NÃO está no escopo ou é restrito
5. **Glossário Inicial** — Termos canônicos do domínio
### Passo 4 — Validação Pós-Geração

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `product-manager` | Visão de produto high-level | Produto |
| 2 | `business-analyst` | Análise dos épicos de negócio | Negócio |
| 3 | `senior-architect` | Alinhamento com soluções técnicas | Arquitetura |
| 4 | `requirements-engineering` | Estruturação de requisitos macro | Requisitos |
| 5 | `documentation-writer` | Redigir o PRD Discovery-Level | Documentação |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador do PRD Discovery-Level (F1 — Bloco 0) | Time de Arquitetura |

🤖 *Upstream Architecture Discovery — Fase 1*
