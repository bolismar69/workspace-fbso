# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-ARCHITECTURE-DEFINITION (F2)

## Contexto

Gate de Validação da Arquitetura Detail-Level (`DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md`). Verifica C4 L2/L3, ADRs, padrões e matriz de integração.

**Princípio fundamental:** Refinamento real além do Discovery-Level — specs acionáveis para implementação.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho das soluções técnicas |

**Arquivo auditado:** `DETAIL-LEVEL-ARCHITECTURE-DEFINITION.md`
**Referências:** F1 (PRD), stack matrix (se existir)

## Dimensões de Validação

### Dimensão 1: Completude Arquitetural
| # | Verificação | Critério |
|---|---|---|
| 1.1 | C4 L2 presente | Containers com tecnologias e protocolos |
| 1.2 | C4 L3 para serviços principais | Componentes internos de cada serviço |
| 1.3 | ADRs detalhados | ≥5 ADRs com contexto, decisão, trade-offs, diagramas |
| 1.4 | Matriz de integração | Cada par origem→destino com protocolo e autenticação |

### Dimensão 2: Padrões e Estrutura
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Padrões de código | Estrutura de pacotes, naming, design patterns |
| 2.2 | Estratégia multi-tenancy | Isolamento documentado |
| 2.3 | API versioning | Estratégia definida |

### Dimensão 3: Independência Tecnológica
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Generalista | Não presume stack específica sem justificativa |
| 3.2 | Se tecnologia específica citada | Skills relacionados foram buscados |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE ARQUITETURA: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-ARCH-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A arquitetura está completa e cobre todas as integrações?
2. Deseja enviar mais documentos/ADRs/blueprints?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 2 do Downstream Architecture Refinement*
