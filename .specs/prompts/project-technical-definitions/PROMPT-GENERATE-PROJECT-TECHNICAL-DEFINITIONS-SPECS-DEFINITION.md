# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION (F16)

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` — a **consolidação técnica enxuta** do projeto. Diferentemente de uma baseline exaustiva, este documento **sumariza e referencia** os artefatos técnicos já produzidos nos blocos anteriores, funcionando como um índice navegável que evita duplicação de conteúdo.

Cada seção contém aproximadamente 1 parágrafo de sumário executivo seguido de `→ ver [ARTEFATO]` para o detalhamento completo. O documento NÃO repete o conteúdo dos artefatos referenciados — apenas os consolida.

**Inputs upstream (Bloco C — F16):** Consolida todos os artefatos dos blocos anteriores:
- **Bloco 0 (Product Def & Backlog & PRD):** `INTAKE-LOG.md`, `DOR-ASSESSMENT.md`, `PRODUCT-BACKLOG-LIST.md`, `PRD-DEFINITION.md`
- **Bloco A (People & Solutions):** `TEAM-SKILLS-MAP.md`, `TEAM-CAPACITY.md`
- **Bloco B (6 Disciplinas Técnicas):** `ARCHITECTURE-DEFINITION.md`, `DATA-ARCHITECTURE-DEFINITION.md`, `SECURITY-DEFINITION.md`, `DEVOPS-SRE-DEFINITION.md`, `TEST-STRATEGY-DEFINITION.md`, `INFRA-CLOUD-DEFINITION.md`
- **Bloco C (Catálogo, Matriz, Stack):** `SOLUTIONS-CATALOG.md` (F13), `SOLUTIONS-MATRIX.md` (F14), `SOLUTIONS-STACK-MATRIX.md` (F15)

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
Carregar TODOS os artefatos dos Blocos 0, A, B e C (listados em Inputs upstream). Identificar seções-chave de cada artefato que merecem destaque na consolidação.

### Passo 2 — Verificar Disponibilidade dos Artefatos
Confirmar que 100% dos artefatos referenciados existem nos caminhos esperados. Artefato faltante = seção não pode ser gerada.

### Passo 3 — Gerar o Artefato (Estrutura Enxuta)
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md` com as seguintes 10 seções. **Cada seção deve conter ~1 parágrafo de sumário + `→ ver [ARTEFATO]` para detalhes. NÃO duplicar conteúdo dos artefatos referenciados.**

1. **Convenções Cross-Solution** — Padrões de nomenclatura, estrutura de projetos e convenções de código que se aplicam a todas as soluções. → ver `ARCHITECTURE-DEFINITION.md`

2. **Padrões de API** — Design de endpoints, versionamento, paginação, formato de erros, headers e contratos OpenAPI. → ver `ARCHITECTURE-DEFINITION.md` §matriz-integração

3. **Padrões de Dados** — Modelagem de dados, schemas, estratégias de migração e políticas de armazenamento. → ver `DATA-ARCHITECTURE-DEFINITION.md`

4. **Padrões de Segurança** — Controles de acesso, IAM, criptografia, threat model e compliance aplicáveis. → ver `SECURITY-DEFINITION.md`

5. **Padrões DevOps/SRE** — Pipelines CI/CD, IaC, observabilidade, SLOs e runbooks. → ver `DEVOPS-SRE-DEFINITION.md`

6. **Estratégia de Testes** — Pirâmide de testes, automação, integração contínua e critérios de qualidade. → ver `TEST-STRATEGY-DEFINITION.md`

7. **Topologia de Infra** — Computação, rede, armazenamento, disaster recovery e provisionamento. → ver `INFRA-CLOUD-DEFINITION.md`

8. **Stacks por Solução** — Mapeamento de tecnologias por solução com versões e justificativas. → ver `SOLUTIONS-STACK-MATRIX.md`

9. **Decisões Técnicas Transversais** — ADRs e decisões de design que impactam múltiplas soluções, extraídas da `ARCHITECTURE-DEFINITION.md` e dos blueprints.

10. **Restrições e Limites Técnicos** — Timeouts, rate limits, tamanhos máximos, limites de concorrência e outras restrições cross-solution.

### Passo 4 — Validação Pós-Geração
Verificar: 100% dos artefatos referenciados existem, links markdown válidos, sumários corretos (não duplicam conteúdo), todas as 10 seções preenchidas.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-architect` | Consolidar decisões arquiteturais | Arquitetura |
| 2 | `reference-builder` | Construir referências cruzadas entre artefatos | Mapeamento |
| 3 | `technical-design-doc-creator` | Redigir sumários técnicos concisos | Documentação |
| 4 | `documentation-writer` | Redigir a Specs Definition consolidada | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência. Priorizar skills de síntese e consolidação sobre skills de geração de conteúdo detalhado.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da baseline de especificações | Time de Arquitetura |
| 2.0 | 30/07/2026 | Reformulação: de baseline exaustiva para consolidação técnica enxuta (10 seções com sumário + referência) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
