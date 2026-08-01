# PROMPT-GATE-SOURCING-FACTORY-BIDDING-ESTIMATE-RETROSPECTIVE-PIB (F5b)

## Contexto

Gate de Validação da Análise Retrospectiva PIB (`ESTIMATE-RETROSPECTIVE-PIB.md`). Fase condicional — só executa se F5 resultou em 0 aprovadas.

**Princípio fundamental:** A retrospectiva complementa a F5 com análise qualitativa. Não altera vereditos — apenas enriquece o feedback para realinhamento.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `discovery` ou `full` |

**Arquivo auditado:** `ESTIMATE-RETROSPECTIVE-PIB.md`
**Referências:** F5 (ESTIMATE-VALIDATION.md), CSVs das fábricas, Baseline PIB

## Condicionalidade

⚠️ Se F5 aprovou ≥ 1 fábrica → **gate não executa**. Emitir `✅ [COMPLIANCE — FASE NÃO APLICÁVEL]`.

## Dimensões de Validação

### Dimensão 1: Completude da Análise
| # | Verificação | Critério |
|---|---|---|
| 1.1 | PIB por épico | Tabela com baseline, melhor fábrica, desvio para cada épico |
| 1.2 | Flat estimates | CV entre épicos calculado para cada fábrica; alertas emitidos |
| 1.3 | QA/Arch fixo | Verificação de valores absolutos idênticos entre épicos |
| 1.4 | Comentários | Análise de qualidade dos textos; duplicatas identificadas |
| 1.5 | Independência | Comparação cross-fábrica; valores idênticos documentados |

### Dimensão 2: Não-Alteração de Vereditos
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Vereditos F5 preservados | Nenhum veredito da F5 foi alterado |
| 2.2 | Caráter complementar | Documento deixa explícito que é análise complementar |

### Dimensão 3: Acionabilidade
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Recomendações específicas | Cada observação tem ação concreta para realinhamento |
| 3.2 | Fábricas afetadas listadas | Para cada problema, quais fábricas são afetadas |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE RETROSPECTIVA**
**🔍 Pontos Conflitantes:** [ID-RETRO-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: FASE NÃO APLICÁVEL
**✅ STATUS: [COMPLIANCE — FASE NÃO APLICÁVEL]**
Motivo: Pelo menos 1 fábrica aprovada na F5.

### ✅ CENÁRIO C: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A análise retrospectiva cobre todas as dimensões e identifica problemas relevantes?
2. Deseja enviar mais dados ou evidências para enriquecer a análise?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 5b do Sourcing & Factory Bidding. Condicional: 0 aprovadas.*
