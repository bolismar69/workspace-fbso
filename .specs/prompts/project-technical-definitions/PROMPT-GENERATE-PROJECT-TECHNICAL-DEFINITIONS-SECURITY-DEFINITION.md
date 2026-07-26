# PROMPT-GENERATE-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION

## Contexto

Este prompt gera o artefato `PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md` — a **definição de segurança do projeto** que especializa o GLOBAL-SECURITY.md para o contexto específico do projeto, definindo regras e premissas de segurança que se aplicam a TODAS as soluções.

**Relação com GLOBAL-SECURITY.md:** O SECURITY-DEFINITION não repete o GLOBAL-SECURITY.md — ele o **especializa** para este projeto, aplicando as regras de ouro, checklist SDD e ferramentas de verificação ao contexto concreto das soluções, stacks e integrações do projeto.

**Inputs upstream:** `PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md` (Fase 5) + `PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md` (Fase 4) + `GLOBAL-SECURITY.md` + `PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md` (Fase 2).

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{SECURITY_GLOBAL}` | Caminho do GLOBAL-SECURITY.md |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Validar que `{SECURITY_GLOBAL}` aponta para arquivo existente.

### Passo 1 — Carregar Documentos Base
Ler GLOBAL-SECURITY.md (referência normativa mestra), Architecture Definition (superfícies de ataque), PRD Definition (funcionalidades a proteger), Catálogo de Soluções.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de segurança para construir threat model macro, definir controles cross-solution, estratégia IAM, políticas de secrets e pipeline DevSecOps global.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md` com:
- Threat model do projeto (nível macro — atores, superfícies entre soluções)
- Regras de ouro do GLOBAL-SECURITY.md aplicadas ao contexto do projeto
- Estratégia de IAM cross-solution (Keycloak realms, clients, flows)
- Política de secrets management (Vault, env vars)
- Requisitos de criptografia em trânsito entre soluções (mTLS, TLS 1.3)
- Pipeline DevSecOps global (SAST, SCA, Secret Scanning para todas as soluções)
- Matriz de conformidade regulatória (LGPD, PCI, SOC2 — quais soluções são afetadas)

### Passo 4 — Validação Pós-Geração
Verificar: threat model preenchido, regras de ouro referenciadas, IAM definido, pipeline documentado, matriz de compliance preenchida.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `security-auditor` | Auditar GLOBAL-SECURITY.md e extrair requisitos | Auditoria |
| 2 | `security-audit` | Auditoria complementar de gaps | Auditoria |
| 3 | `threat-modeling-expert` | Construir threat model macro do projeto | Threat Model |
| 4 | `threat-model-analyst` | Analisar ameaças cross-solution | Threat Model |
| 5 | `senior-security` | Supervisão sênior de segurança | Estratégia |
| 6 | `security-best-practices` | Boas práticas de segurança para todas as stacks | Boas Práticas |
| 7 | `api-security-best-practices` | Segurança de APIs cross-solution | API Security |
| 8 | `security-scanning-security-sast` | Pipeline SAST global | Pipeline |
| 9 | `security-scanning-security-dependencies` | SCA global | Pipeline |
| 10 | `security-scanning-security-hardening` | Hardening cross-solution | Infra |
| 11 | `secrets-management` | Estratégia de secrets management | DevSecOps |
| 12 | `gdpr-compliant` | Conformidade LGPD | Compliance |
| 13 | `auth-implementation-patterns` | Estratégia IAM (Keycloak) | IAM |
| 14 | `security-reviewer` | Revisão final contra GLOBAL-SECURITY.md | Revisão |
| 15 | `documentation-writer` | Redigir o Security Definition | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da definição de segurança | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
