# DISCOVERY-LEVEL-SECURITY-DEFINITION — Definição de Segurança (Discovery)
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG · **Fase:** F3 — Bloco B · **Disciplina:** Security Architect
- **Versão:** 1.0 · **Data:** 30/07/2026 · **Status:** CREATED

## 1. Threat Model (High-Level)
| Ameaça | Superfície | Severidade |
|:---|:---|:---:|
| Vazamento de dados entre tenants | Banco de dados (queries sem filtro tenant_id) | Crítico |
| Roubo de sessão / token JWT | Frontend (XSS, token storage) | Alto |
| Acesso não autorizado a endpoints admin | API Gateway (bypass de autenticação) | Alto |
| Injeção SQL / XSS | Backend API (input não sanitizado) | Médio |
| Exposição de dados pessoais (LGPD) | Audit logs, backups | Crítico |

## 2. Requisitos de Compliance
| Requisito | Regulação | Ação |
|:---|:---|:---|
| LGPD — Dados pessoais | Lei 13.709/2018 | Anonimização em não-prod; soft delete; auditoria de acessos |
| Autenticação forte | ISO 27001 | OIDC + PKCE; MFA para admins |
| Logs de auditoria | PCI-DSS (referência) | Toda ação administrativa registrada em audit_log |

## 3. Estratégia de Segurança (Macro)
| Camada | Controle |
|:---|:---|
| Rede | Cloudflare → Kong API Gateway como único ponto de entrada público |
| Aplicação | Microserviços confiam no Kong (trust boundary); RBAC baseado nas roles injetadas pelo Kong no header JWT |
| Dados | Criptografia em repouso (AES-256); RLS para isolamento multi-tenant |
| Identidade | Keycloak como IdP; realms por ambiente |

## 4. Riscos de Segurança
| Risco | Severidade | Mitigação |
|:---|:---:|:---|
| Configuração incorreta do RLS expõe dados | Crítico | Testes automatizados de isolamento por tenant |
| Keycloak mal configurado permite escalação de privilégio | Alto | Revisão de realms/roles por security specialist |
| Dependências com vulnerabilidades conhecidas | Médio | Dependabot + Snyk no CI; atualização semanal |

## 5. Estimativa de Esforço
| Atividade | Complexidade | Esforço |
|:---|:---:|:---:|
| Configuração Keycloak (realms, OIDC, roles) | Média | 1-2 homem-mês |
| Implementação RLS + TenantAwareDataSource | Média | 1 homem-mês |
| Kong + OIDC Plugin | Média | 0.5-1 homem-mês |
| Auditoria e compliance (LGPD) | Média | 0.5-1 homem-mês |
| **Total Segurança** | — | **3-5 homem-mês** |

🤖 *Upstream Architecture Discovery — Fase 3*
