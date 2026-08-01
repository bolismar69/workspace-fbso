# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-BOTTOM-UP-PERT-ESTIMATE (F8)

## Contexto

Gate de Validação da Estimativa Bottom-Up PERT (`BOTTOM-UP-PERT-ESTIMATE.md`). Este é o **gate mais crítico do roadmap** — valida a estimativa independente e aplica as regras DTA.

**Princípios fundamentais:**
1. **Independência:** NENHUMA referência ao ROM upstream nos cálculos
2. **Cobertura:** Todas as US estimadas individualmente com O, ML, P, PERT, σ
3. **DTA:** QA ≥ 25%, Arch ≥ 5%, consistência prazo×horas

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |

**Arquivo auditado:** `BOTTOM-UP-PERT-ESTIMATE.md`
**Referências:** User Stories (todas), F2-F7 (Detail-Level)

## Dimensões de Validação

### Dimensão 1: Independência (CRÍTICA)
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Zero referências ao ROM | Nenhuma menção a `DISCOVERY-LEVEL-ROM-ESTIMATE.md` nos cálculos |
| 1.2 | Cálculo do zero | Cada US com O, ML, P justificados pela complexidade, não por estimativa anterior |
| 1.3 | Se ROM referenciado | Apenas na seção de comparação (se existir), NUNCA como baseline |

### Dimensão 2: Cobertura (CRÍTICA)
| # | Verificação | Critério |
|---|---|---|
| 2.1 | US individuais | Todas as US com linha na estimativa |
| 2.2 | PERT completo | O, ML, P, PERT, σ para cada US |
| 2.3 | IC 95% | Calculado em todos os níveis (US, feature, épico, projeto) |
| 2.4 | Rollup correto | Feature→Épico→Projeto sem erros de soma |

### Dimensão 3: Composição do Esforço
| # | Verificação | Critério |
|---|---|---|
| 3.1 | QA ≥ 25% | Verificado por épico e global |
| 3.2 | Arch ≥ 5% | Verificado global |
| 3.3 | DevOps + Gestão | Percentuais justificados |
| 3.4 | Contingência | 15-25% documentado |

### Dimensão 4: Validação DTA
| # | Verificação | Critério |
|---|---|---|
| 4.1 | QA balanceado | ≥20% por épico |
| 4.2 | Consistência prazo×horas | Divergência ≤50% |
| 4.3 | Outliers | Identificados e justificados |

### Dimensão 5: Completude
| # | Verificação | Critério |
|---|---|---|
| 5.1 | Sumário executivo | Tabela com Dev, Subtotal, Total |
| 5.2 | Análise de riscos | Riscos mapeados com impacto |
| 5.3 | Documentos relacionados | Links para artefatos referenciados |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE

**📊 RELATÓRIO DE AUDITORIA DE ESTIMATIVA: {PROJECT_ID_NAME}**

**🔍 Pontos Conflitantes:**
- **[ID-EST-XX] — [Título]:** [Problema] → [Sugestão]

**🛑 STATUS: [NÃO COMPLIANCE]**

**Regra especial:** Se o conflito for **ID-EST-01 (Contaminação pelo ROM)**, a estimativa deve ser refeita do zero — NÃO é permitido FIX cirúrgico. Isso garante a independência real da estimativa.

### ✅ CENÁRIO B: PRÉ-COMPLIANCE

**📊 RELATÓRIO DE AUDITORIA DE ESTIMATIVA: {PROJECT_ID_NAME}**
- Total US estimadas: {N}
- QA Global: {X}% ≥ 25% ✅
- Arch Global: {Y}% ≥ 5% ✅
- Independência comprovada: ✅
- IC 95% Projeto: {low}h – {high}h

**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**

**❓ 3 Perguntas Obrigatórias:**
1. A estimativa está completa e cobre todas as US do escopo?
2. Deseja enviar mais documentos para refinar a estimativa?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 8 do Downstream Architecture Refinement · Validação DTA + Independência*
