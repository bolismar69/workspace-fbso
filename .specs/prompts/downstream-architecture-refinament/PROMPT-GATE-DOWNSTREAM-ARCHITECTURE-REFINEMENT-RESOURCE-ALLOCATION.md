# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-RESOURCE-ALLOCATION (F9)

## Contexto

Gate de Validação do Plano de Alocação de Recursos (`RESOURCE-ALLOCATION-PLAN.md`). Verifica se a alocação é derivada exclusivamente da estimativa PERT (F8).

**Princípio fundamental:** Alocação 100% baseada na estimativa PERT, com capacidade realista e gargalos identificados.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

**Arquivo auditado:** `RESOURCE-ALLOCATION-PLAN.md`
**Referência obrigatória:** F8 (PERT Estimate) — única fonte de horas

## Dimensões de Validação

### Dimensão 1: Fidelidade ao PERT
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Horas derivadas do PERT | Todos os números vêm da F8 |
| 1.2 | Sem fonte externa | Nenhuma referência a ROM upstream ou factory bids |

### Dimensão 2: Realismo
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Capacidade efetiva | Calculada com cargas parciais |
| 2.2 | Duração projetada | Consistente com capacidade × horas |
| 2.3 | Gargalos | Identificados e com recomendação |

### Dimensão 3: Completude
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Papéis necessários | Todos os papéis/perfis necessários listados (sem nomeação individual — estimativa) |
| 3.2 | Alocação por épico | Distribuição de horas documentada |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE ALOCAÇÃO: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-RES-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A alocação de recursos reflete a realidade do time?
2. Deseja enviar mais documentos de capacidade/time?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 9 do Downstream Architecture Refinement*
