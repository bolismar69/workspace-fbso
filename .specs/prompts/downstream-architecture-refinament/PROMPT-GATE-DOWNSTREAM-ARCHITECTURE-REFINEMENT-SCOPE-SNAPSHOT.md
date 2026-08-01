# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SCOPE-SNAPSHOT (F11)

## Contexto

Gate de Validação do Scope Snapshot (`SCOPE-SNAPSHOT.md`). Verifica se a foto do escopo cobre 100% das US estimadas e está corretamente congelada.

**Princípio fundamental:** Snapshot é uma foto imutável do escopo no momento da estimativa. NÃO contém planejamento de sprints ou contratos.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base |
| `{PROJECT_ID_NAME}` | ID do projeto |

**Arquivo auditado:** `SCOPE-SNAPSHOT.md`
**Referências:** F8 (PERT), RTM (05-USER-STORIES)

## Dimensões de Validação

### Dimensão 1: Cobertura
| # | Verificação | Critério |
|---|---|---|
| 1.1 | 100% US | Todas as US da estimativa listadas |
| 1.2 | Consistência com RTM | US IDs batem com a RTM oficial |
| 1.3 | Sem órfãos | Nenhuma US fora do snapshot |

### Dimensão 2: Formato
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Sem planejamento | NÃO contém sprints, tarefas ou contratos |
| 2.2 | Data de congelamento | Registrada explicitamente |
| 2.3 | Hash do escopo | Checksum para detecção de mudanças |

### Dimensão 3: Integridade
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Nota de imutabilidade | Snapshot não deve ser alterado sem change request |
| 3.2 | Referência à estimativa | Link para F8 (BOTTOM-UP-PERT-ESTIMATE) |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE SNAPSHOT: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-SNAP-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. O snapshot reflete exatamente o escopo que foi estimado?
2. Deseja enviar mais documentos de escopo?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 11 do Downstream Architecture Refinement*
