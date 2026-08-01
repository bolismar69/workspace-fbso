# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-SECURITY-DEFINITION (F3)

## Contexto

Gate de Validação da Segurança Detail-Level (`DETAIL-LEVEL-SECURITY-DEFINITION.md`). Verifica threat model STRIDE, OWASP ASVS, IAM specs e compliance.

**Princípio fundamental:** Threat model cobre todos os componentes da arquitetura e controles são acionáveis.

## Parâmetros

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | ID completo do projeto |
| `{SECURITY_GLOBAL}` | GLOBAL-SECURITY.md |

**Arquivo auditado:** `DETAIL-LEVEL-SECURITY-DEFINITION.md`
**Referências:** F2 (Arquitetura), GLOBAL-SECURITY.md

## Dimensões de Validação

### Dimensão 1: Threat Model
| # | Verificação | Critério |
|---|---|---|
| 1.1 | STRIDE por componente | Cada componente da arquitetura analisado |
| 1.2 | Mitigações definidas | Cada ameaça tem controle correspondente |

### Dimensão 2: Controles e Compliance
| # | Verificação | Critério |
|---|---|---|
| 2.1 | OWASP ASVS L1+L2 | Matriz de controles preenchida |
| 2.2 | IAM specs | Realms, clients, claims, fluxos documentados |
| 2.3 | Matriz RBAC | Role×Permission×Resource completa |
| 2.4 | Compliance regulatório | Requisitos mapeados para controles |

### Dimensão 3: Consistência
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Alinhado com F2 | Controles implementam padrões arquiteturais |
| 3.2 | Alinhado com GLOBAL-SECURITY | Padrões corporativos seguidos |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE
**📊 RELATÓRIO DE AUDITORIA DE SEGURANÇA: {PROJECT_ID_NAME}**
**🔍 Pontos Conflitantes:** [ID-SEC-XX]
**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**
**❓ 3 Perguntas Obrigatórias:**
1. A definição de segurança cobre todas as ameaças e compliance?
2. Deseja enviar mais documentos de segurança?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 3 do Downstream Architecture Refinement*
