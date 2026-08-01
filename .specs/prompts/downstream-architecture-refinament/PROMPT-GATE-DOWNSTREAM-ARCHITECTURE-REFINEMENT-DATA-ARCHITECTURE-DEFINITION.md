# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DATA-ARCHITECTURE-DEFINITION (F4)

## Contexto

Gate de Validação da Arquitetura de Dados Detail-Level (`DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md`). Verifica modelo de dados, multi-tenancy, query patterns e migrations.

**Princípio fundamental:** Modelo de dados suporta todos os casos de uso das US e estratégia de isolamento multi-tenant é completa.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |

**Arquivo auditado:** `DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION.md`
**Referências:** F2 (Arquitetura), F3 (Segurança — RLS)

## Dimensões de Validação

### Dimensão 1: Modelo de Dados
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Entidades completas | Atributos, tipos, constraints, índices, FKs |
| 1.2 | Cobertura de features | Tabelas para todas as features (tenants, planos, RBAC, auditoria, BUs, produtos) |

### Dimensão 2: Multi-Tenancy
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Estratégia de isolamento | Documentada e justificada |
| 2.2 | Políticas de isolamento | Por tabela, com verificação |

### Dimensão 3: Performance e Ciclo de Vida
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Query patterns | Índices e otimizações para queries críticas |
| 3.2 | Particionamento | Estratégia para tabelas de alto volume |
| 3.3 | Migrations | Versionamento, baseline, rollback |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE DADOS: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-DATA-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. O modelo de dados está completo e cobre todos os casos de uso?
2. Deseja enviar mais documentos de dados?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 4 do Downstream Architecture Refinement*
