# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-DEVOPS-SRE-DEFINITION (F5)

## Contexto

Gate de Validação do DevOps/SRE Detail-Level (`DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md`). Verifica pipeline specs, IaC, observabilidade, SLOs e runbooks.

**Princípio fundamental:** Pipeline cobre todos os ambientes, SLOs têm SLIs mensuráveis, e observabilidade cobre toda a stack.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |

**Arquivo auditado:** `DETAIL-LEVEL-DEVOPS-SRE-DEFINITION.md`
**Referências:** F2 (Arquitetura), F3 (Segurança)

## Dimensões de Validação

### Dimensão 1: Pipeline e Deploy
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Pipelines por ambiente | Dev, Staging, Produção com workflows |
| 1.2 | IaC | Templates para todos os recursos |
| 1.3 | Estratégia de deploy | Blue-green, canary ou similar definido |

### Dimensão 2: Observabilidade
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Métricas | Ferramenta + alert rules definidas |
| 2.2 | Logs | Estratégia de coleta e consulta |
| 2.3 | Tracing | Propagação de trace entre serviços |

### Dimensão 3: SLOs e Runbooks
| # | Verificação | Critério |
|---|---|---|
| 3.1 | SLOs com SLIs | Cada SLO tem SLI mensurável e janela |
| 3.2 | Runbooks | Procedimentos para falhas comuns |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE DEVOPS/SRE: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-DEVOPS-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A estratégia DevOps/SRE cobre todos os ambientes e cenários de falha?
2. Deseja enviar mais documentos de infra/SRE?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 5 do Downstream Architecture Refinement*
