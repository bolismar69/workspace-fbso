# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-INFRA-CLOUD-DEFINITION (F7)

## Contexto

Gate de Validação da Infra/Cloud Detail-Level (`DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md`). Verifica sizing, custos, topologia e DR.

**Princípio fundamental:** Sizing suporta volumes projetados, custos são realistas, DR tem RPO/RTO definidos.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |

**Arquivo auditado:** `DETAIL-LEVEL-INFRA-CLOUD-DEFINITION.md`
**Referências:** F2 (Arquitetura), F4 (Dados — volumes), F5 (DevOps)

## Dimensões de Validação

### Dimensão 1: Sizing
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Recursos dimensionados | Compute, banco, cache, gateway |
| 1.2 | Justificativa | Cada tier com justificativa baseada em volumes |
| 1.3 | Ambientes | Dev, Staging, Produção definidos |

### Dimensão 2: Custos e Rede
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Custos mensais | Detalhados por provedor e serviço |
| 2.2 | Topologia de rede | Diagrama com fluxo de acesso |
| 2.3 | Segurança de rede | Firewall, WAF, DDoS |

### Dimensão 3: Disaster Recovery
| # | Verificação | Critério |
|---|---|---|
| 3.1 | RPO/RTO | Definidos e justificados |
| 3.2 | Estratégia de backup | Ferramenta, frequência, retenção |
| 3.3 | DR Drill | Periodicidade definida |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE INFRA: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-INFRA-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A infraestrutura está dimensionada para os volumes projetados?
2. Deseja enviar mais documentos de infra/cloud?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 7 do Downstream Architecture Refinement*
