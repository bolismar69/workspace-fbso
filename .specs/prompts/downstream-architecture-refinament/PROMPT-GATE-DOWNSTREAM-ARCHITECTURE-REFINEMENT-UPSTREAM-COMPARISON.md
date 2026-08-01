# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-UPSTREAM-COMPARISON (F12)

## Contexto

Gate de Validação do Relatório Comparativo com Upstream (`UPSTREAM-COMPARISON-REPORT.md`). Fase **condicional** — só executa se upstream existir.

**Princípios fundamentais:**
1. Relatório é informativo — NÃO altera a estimativa PERT
2. Se upstream não existe: gate automaticamente satisfeito
3. Se upstream existe: relatório deve comparar sem sugerir alterações no PERT

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

**Arquivo auditado:** `UPSTREAM-COMPARISON-REPORT.md`
**Referências:** F8 (PERT — congelada), ROM upstream (se existir)

## Condicionalidade

⚠️ Se `upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md` NÃO existir:
- Emitir `✅ STATUS: [COMPLIANCE — FASE NÃO APLICÁVEL]`
- Motivo: "Upstream discovery não encontrado — sem ROM para comparação"
- Não executar as dimensões abaixo

## Dimensões de Validação (apenas se upstream existir)

### Dimensão 1: Não-Alteração do PERT
| # | Verificação | Critério |
|---|---|---|
| 1.1 | PERT inalterado | Nenhum número do PERT foi modificado |
| 1.2 | Caráter informativo | Relatório deixa explícito que é informativo |

### Dimensão 2: Completude da Comparação
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Tabela comparativa | ROM vs PERT por épico e total |
| 2.2 | Análise de desvio | Percentual documentado e justificado |
| 2.3 | Causas de divergência | Se PERT fora da faixa ROM, causas analisadas |

### Dimensão 3: Visualização
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Gráfico comparativo | Visualização proporcional |
| 3.2 | Conclusão clara | Refina, substitui ou diverge? |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE COMPARAÇÃO: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-COMP-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: UPSTREAM NÃO EXISTE
**✅ STATUS: [COMPLIANCE — FASE NÃO APLICÁVEL]**
Motivo: Upstream discovery não encontrado.

### ✅ CENÁRIO C: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. O relatório comparativo está completo e claro para o comitê?
2. Deseja enviar mais documentos para enriquecer a comparação?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 12 do Downstream Architecture Refinement · Cross-Check Condicional*
