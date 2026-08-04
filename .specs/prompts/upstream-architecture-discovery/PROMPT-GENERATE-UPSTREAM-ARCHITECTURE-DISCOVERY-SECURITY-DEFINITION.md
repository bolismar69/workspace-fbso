# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md` — a **definição de segurança do projeto** que especializa o GLOBAL-SECURITY.md para o contexto específico do projeto, definindo regras e premissas de segurança que se aplicam a TODAS as soluções.

**Relação com GLOBAL-SECURITY.md:** O SECURITY-DEFINITION não repete o GLOBAL-SECURITY.md — ele o **especializa** para este projeto, aplicando as regras de ouro, checklist SDD e ferramentas de verificação ao contexto concreto das soluções, stacks e integrações do projeto.

**Este documento é independente de tecnologias específicas de segurança.** Durante a análise da stack de segurança do projeto, identifique as tecnologias utilizadas (IAM, API Gateway, WAF) e busque skills relacionados. Caso não encontre skills específicos, utilize skills generalistas de segurança, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
4. `{SECURITY_GLOBAL}` — GLOBAL-SECURITY.md (regras de ouro corporativas)

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills necessários para o time de implementação (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time — seniores, plenos, juniores, duração (obter e validar com usuário) |
| `{PROJECT-STACK}` | Stack tecnológica da solução. Baseline corporativa: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão exigem justificativa |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `TECHNICAL_SOLUTION_PATH={TECHNICAL_SOLUTION_PATH}` — base das soluções técnicas
- `TECHNICAL_SOLUTION_NAMES={TECHNICAL_SOLUTION_NAMES}` — soluções do projeto
- `ARCHITECTURE_GLOBAL={ARCHITECTURE_GLOBAL}` — ADRs e blueprints globais
- `SECURITY_GLOBAL={SECURITY_GLOBAL}` — documento de segurança global
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP` — skills do time (se fornecidos)
- `PROJECT-TEAM-CAPACITY` — capacidade do time (se fornecida)
- `PROJECT-STACK` — stack tecnológica; validar contra STACK-PADROES-CORPORATIVOS-FBSO-ORG.md
Validar que `{SECURITY_GLOBAL}` aponta para arquivo existente e que `{UPSTREAM_DISCOVERY_PATH}` contém os artefatos upstream.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos upstream:
1. `{SECURITY_GLOBAL}` — GLOBAL-SECURITY.md (referência normativa mestra)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2) — superfícies de ataque
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)

### Passo 2 — Invocar Skills Especializadas
Invocar skills de segurança para construir threat model macro, definir controles cross-solution, estratégia IAM, políticas de secrets e pipeline DevSecOps global.

### Passo 2.5 — Apresentar Skills, Capacidade e Stack para Validação Humana

Avaliar e apresentar ao usuário para validação:

1. **PROJECT-TEAM-SKILLS-MAP:** Skills identificados como necessários para implementar a solução nesta disciplina.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

2. **PROJECT-TEAM-CAPACITY:** Capacidade estimada do time nesta disciplina (ex: 2 seniores, 3 plenos).
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

3. **PROJECT-STACK:** Tecnologias identificadas para esta disciplina. Verificar conformidade com `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão corporativo DEVEM ser listadas com justificativa técnica e requerem aprovação explícita do usuário.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-SECURITY-DEFINITION.md` com:
- Threat model do projeto (nível macro — atores, superfícies entre soluções)
- Regras de ouro do GLOBAL-SECURITY.md aplicadas ao contexto do projeto
- Estratégia de IAM cross-solution (Keycloak realms, clients, flows)
- Política de secrets management (Vault, env vars)
- Requisitos de criptografia em trânsito entre soluções (mTLS, TLS 1.3)
- Pipeline DevSecOps global (SAST, SCA, Secret Scanning para todas as soluções)
- Matriz de conformidade regulatória (LGPD, PCI, SOC2 — quais soluções são afetadas)

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-SECURITY-DEFINITION.md`

```markdown
# DISCOVERY-LEVEL-SECURITY-DEFINITION.md
## Fase 3 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-SECURITY-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | {DATA_ATUAL} |
| **Autor** | Security Architect / IAM Specialist |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Vinculados:**
- [`GLOBAL-SECURITY.md`](../../../.specs/security/GLOBAL-SECURITY.md) — Política de Segurança Global (baseline corporativa)
- [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md) — PRD Discovery-Level (F1)
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)
- [`STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`](../../../.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md) — Padrões Corporativos

---

## 1. Postura de Segurança — Visão Macro
- **1.1 Modelo de Defesa em Camadas:** Diagrama ASCII com 5 camadas (Edge → API Gateway → Service Mesh → Application → Data)
- **1.2 Trust Boundary — API Gateway:** Regra de ouro e justificativa do ponto único de entrada

## 2. Threat Model — Análise de Ameaças (STRIDE Discovery-Level)
- **2.1 Ameaças Priorizadas:** Tabela: ID | Ameaça | Categoria STRIDE | Cenário de Ataque | Severidade | Mitigação
- **2.2 Trust Zones:** Tabela: Zona | Componentes | Nível de Confiança | Controles

## 3. Compliance Regulatório
- **3.1 LGPD (ou regulação aplicável):** Tabela: Artigo | Implicação para o Projeto | Ação Requerida
- **3.2 Frameworks de Segurança Aplicáveis:** Tabela: Framework | Aplicação | Status

## 4. Estratégia de Segurança por Componente
- **4.1 Edge/CDN/WAF:** Tabela de controles e configurações
- **4.2 API Gateway — Trust Boundary:** Tabela de controles (JWT, Rate Limiting, CORS, Header Sanitization, IP Restriction, Request Size)
- **4.3 IAM — Identity Provider:** Tabela de controles (Realms, JWT Claims, Token Expiry, Brute Force, Password Policy)
- **4.4 Backend — Aplicação:** Tabela de controles (RBAC Enforcement, Input Validation, SQL Protection, Error Handling, Audit Trail, Secrets, Log Sanitization)
- **4.5 Banco de Dados — Camada de Dados:** Tabela de controles (RLS, TLS, Encryption at Rest, Soft Delete, Backup)
- **4.6 Rede e Service Mesh:** Tabela de controles (NetworkPolicy, mTLS, AuthorizationPolicy, Egress)

## 5. Checklist SDD — Aplicação ao Projeto
- Subseções 5.1-5.4 com checklist aplicado: Autenticação/Autorização, Proteção de Dados, Validação de Entradas, Proteção de Infra/API
- Cada item: Item | Aplicação no Projeto

## 6. Riscos de Segurança e Estimativa de Esforço
- **6.1 Riscos Específicos de Segurança:** Tabela: ID | Risco | Prob. | Impacto | Mitigação
- **6.2 Estimativa de Esforço de Segurança:** Tabela: Atividade | Complexidade | Esforço (dias) | Responsável

## 7. Ferramentas de Verificação Automatizada
- Tabela: Ferramenta | Pipeline Stage | Verifica

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Security Definition Discovery-Level | Security Architect / IAM Specialist |
```

### Passo 4 — Validação Pós-Geração

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


**Skills generalistas de segurança (sempre aplicáveis):**
- `engineering-skills`, `engineering-advanced-skills`
- `senior-security`, `security-best-practices`, `security-review`
- `security-reviewer`, `security-audit`, `security-auditor`
- `security-scanning-security-sast`, `security-threat-model`
- `threat-modeling-expert`, `threat-model-analyst`
- `privacy-by-design`, `gdpr-compliant`
- `secrets-management`, `secret-scanning`

**Skills tecnológicos de segurança (condicionais — buscar ao identificar a stack):**
- Ao identificar uma tecnologia específica de segurança (IAM, API Gateway, WAF) durante a análise da stack, procure skills relacionados a essa tecnologia para aprimorar as especificações
- Caso não encontre skills específicos para a tecnologia identificada, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da definição de segurança | Time de Arquitetura |
| 2.0 | 30/07/2026 | Atualização F4→F8: referência ARCHITECTURE-DEFINITION atualizada para Fase 7; adicionada referência cruzada aos artefatos F9-F12 (Bloco B) | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{SECURITY_GLOBAL}` | GLOBAL-SECURITY.md — regras de ouro corporativas |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) — funcionalidades a proteger |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) — superfícies de ataque |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` | Catálogo de Soluções (F8) |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (se fornecidos) |
| 7 | `{PROJECT-TEAM-CAPACITY}` | Capacidade do time (se fornecida) |
| 8 | `{PROJECT-STACK}` | Stack tecnológica (validar contra padrões corporativos) |
| 9 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
