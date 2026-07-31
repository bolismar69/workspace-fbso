# PROMPT-FIX-UPSTREAM-ARCHITECTURE-DISCOVERY-SPECS

## Contexto

Este prompt implementa o **FIX do SPECS Discovery-Level** — Fase do Bloco C/D do Upstream Architecture Discovery.

**Princípio fundamental:** O artefato Discovery-Level deve conter informações suficientes para embasar a análise de viabilidade e estimativa ROM 50%, sem detalhamento excessivo de implementação.

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{UPSTREAM_DISCOVERY_PATH}` | Caminho upstream-architecture-discovery |

**Arquivos gerados pelo GENERATE:** `DISCOVERY-LEVEL-SPECS.md`

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `DISCOVERY-LEVEL-SPECS.md` e artefatos upstream do Discovery.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Completude — seções obrigatórias preenchidas
#### Dimensão 2: Consistência — alinhamento com artefatos upstream do Discovery
#### Dimensão 3: Nível de Detalhe — adequado para ROM 50% (nem raso demais, nem detalhado demais)

### Passo 3 — Emitir Veredito

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE — lista conflitos
### ✅ CENÁRIO B: PRÉ-COMPLIANCE — 3 perguntas obrigatórias

## Skills Utilizados
| 1 | `requirements-validation` | 2 | `gap-analysis` | 3 | `senior-architect` |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — Upstream Architecture Discovery | Time de Arquitetura |

🤖 *Upstream Architecture Discovery — SPECS FIX*
