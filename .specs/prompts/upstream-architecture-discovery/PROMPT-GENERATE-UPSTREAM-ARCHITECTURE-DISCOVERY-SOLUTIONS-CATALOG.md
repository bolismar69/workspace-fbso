# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SOLUTIONS-CATALOG
## Contexto
> 📐 **Discovery-Level:** Versão macro do catálogo de soluções para análise de viabilidade e ROM 50%.

Este prompt gera `DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — catálogo macro de soluções (nomes, tipos, propósito high-level). Sem detalhamento de stacks ou implementação.

**Papel no Bloco C (Catálogo, Matriz & Consolidação):** Fase 8 de 3. Consome os 6 artefatos do Bloco B e alimenta a Matriz (F9) e SPECS (F10).

**Inputs upstream:** DISCOVERY-LEVEL-PRD (F1) + 6 artefatos do Bloco B (F2-F7).

## Parâmetros de Entrada
(8 parâmetros padrão do roadmap + `{UPSTREAM_DISCOVERY_PATH}`)

## Fluxo de Execução
### Passo 0 — Validação · ### Passo 1 — Carregar PRD + 6 artefatos Bloco B
### Passo 2 — Invocar Skills · ### Passo 3 — Gerar catálogo macro com: nomes, tipos, propósito, complexidade estimada
### Passo 4 — Validação Pós-Geração

## Skills Utilizados
| 1 | `senior-architect` | Catalogação de soluções | 2 | `cloud-architect` | Tipos de serviço cloud |
| 3 | `documentation-writer` | Redigir catálogo |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — F8 Bloco C Discovery-Level | Time de Arquitetura |
🤖 *Upstream Architecture Discovery — Fase 8*
