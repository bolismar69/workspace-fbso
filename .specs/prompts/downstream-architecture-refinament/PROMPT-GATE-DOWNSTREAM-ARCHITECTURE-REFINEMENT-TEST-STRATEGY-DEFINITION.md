# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-TEST-STRATEGY-DEFINITION (F6)

## Contexto

Gate de Validação da Estratégia de Testes Detail-Level (`DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md`). Verifica matriz de cobertura, casos de teste e quality gates.

**Princípio fundamental:** Cada US tem cobertura de testes definida e quality gates são aplicáveis.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |

**Arquivo auditado:** `DETAIL-LEVEL-TEST-STRATEGY-DEFINITION.md`
**Referências:** F2 (Arquitetura), F3 (Segurança), US files

## Dimensões de Validação

### Dimensão 1: Cobertura
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Matriz US×Teste | Cada US com ≥1 tipo de teste |
| 1.2 | Pirâmide balanceada | Unit > Integ > E2E proporcional |
| 1.3 | Testes multi-tenant | Cenários de isolamento entre tenants |

### Dimensão 2: Automação e Gates
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Estratégia CI | Testes automatizados no pipeline |
| 2.2 | Quality gates | PR, Staging, Release com critérios |
| 2.3 | Ferramentas definidas | Framework por camada de teste |

### Dimensão 3: Aceitação
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Casos de aceitação | Baseados nos cenários das US |
| 3.2 | Cobertura de features | Pelo menos 1 caso por feature |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE TESTES: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-TEST-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A estratégia de testes cobre todas as US e features?
2. Deseja enviar mais documentos de teste/qualidade?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 6 do Downstream Architecture Refinement*
