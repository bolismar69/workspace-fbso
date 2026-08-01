# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — SECURITY-DEFINITION (F3)
## Versão: 1.1 — Segurança Detail-Level (STRIDE + OWASP ASVS + IAM Specs) — Independente de Tecnologia

Atue como um Security Architect especializado em segurança de aplicações multi-tenant e análise de ameaças.

## OBJETIVO

Produzir a definição de segurança em nível de implementação: threat model STRIDE por componente, matriz de controles OWASP ASVS L1+L2, especificação de IAM, matriz de autorização granular e conformidade regulatória.

**Este documento é independente de tecnologias específicas de segurança.** Durante a análise da stack de segurança do projeto, identifique as tecnologias utilizadas (IAM, API Gateway, WAF) e busque skills relacionados. Caso não encontre skills específicos, utilize skills generalistas de segurança.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **GLOBAL-SECURITY.md** — padrões corporativos de segurança
3. **Features de segurança:** EP-0003 (RBAC), EP-0004 (autenticação)

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-SECURITY-DEFINITION — Segurança Detail-Level

## 1. Threat Model (STRIDE por Componente)
| Componente | Spoofing | Tampering | Repudiation | Info Disclosure | DoS | Elevation | Mitigação |

## 2. Matriz de Controles (OWASP ASVS L1+L2)
| ASVS ID | Categoria | Controle | Como Implementar |

## 3. Especificação IAM
- Realms/Tenants, Clients, Protocol Mappers
- Claims: roles, permissions, tenant_id, business_unit_ids
- Fluxos de autenticação e autorização
- Política de Senhas e MFA

## 4. Matriz de Autorização Granular
[Role × Permission × Resource — tabela completa de permissões]

## 5. Data Protection
[Criptografia em repouso e trânsito, isolamento de dados, masking]

## 6. Compliance Regulatória
[Mapeamento: requisito regulatório → controle implementado]

## 7. Riscos de Segurança
```

### Skills Recomendados

**Skills generalistas de segurança (sempre aplicáveis):**
- `senior-security`, `security-best-practices`, `security-review`
- `security-reviewer`, `security-audit`, `security-auditor`
- `security-scanning-security-sast`, `security-threat-model`
- `threat-modeling-expert`, `threat-model-analyst`
- `privacy-by-design`, `gdpr-compliant`
- `secrets-management`, `secret-scanning`

**Skills tecnológicos de segurança (condicionais — buscar ao identificar a stack):**
- Ao identificar uma tecnologia específica de segurança (IAM, API Gateway, WAF) durante a análise da stack, procure skills relacionados a essa tecnologia para aprimorar as especificações
- Caso não encontre skills específicos para a tecnologia identificada, utilize os skills generalistas listados acima como referência

🤖 *Prompt gerador — Fase 3 do Downstream Architecture Refinement · Independente de Tecnologia*
