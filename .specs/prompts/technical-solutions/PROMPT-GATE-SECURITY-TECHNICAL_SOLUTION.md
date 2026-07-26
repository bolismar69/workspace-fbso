# PROMPT-GATE-SECURITY-SCOPE

## Contexto

Este prompt implementa o **Gate de Validação de Segurança** para o artefato `TECHNICAL-SOLUTION-SECURITY.md`, conforme definido no fluxo Spec-Driven Development — Fase 3 do Roadmap de Solução Técnica.

O agente validador atua como um **guardião de conformidade de segurança** — verificando se o TECHNICAL-SOLUTION-SECURITY.md está aderente ao GLOBAL-SECURITY.md (Política e Checklist de Segurança), ao threat model derivado do TECHNICAL-SOLUTION-PRD.md e às decisões de arquitetura do TECHNICAL-SOLUTION-ARCHITECTURE.md. O foco é **conformidade normativa e cobertura de riscos** — o TECHNICAL-SOLUTION-SECURITY.md deve especializar as regras globais para a solução específica sem enfraquecê-las.

**Princípio fundamental:** O GLOBAL-SECURITY.md define as regras de ouro inegociáveis. O TECHNICAL-SOLUTION-SECURITY.md as especializa. Se uma regra de ouro for violada ou omitida, o gate reprova automaticamente.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |
| `{SECURITY_GLOBAL}` | Caminho absoluto para o GLOBAL-SECURITY.md | `/home/user/work/.specs/security/GLOBAL-SECURITY.md` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 6 parâmetros foram informados. Se algum estiver ausente, perguntar antes de prosseguir.

### Passo 1 — Carregar Documentos Base

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-SECURITY.md (artefato a ser validado)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-PRD.md (baseline de escopo)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md (pipeline de segurança, cross-cutting)
    └── {SECURITY_GLOBAL} (GLOBAL-SECURITY.md — política e checklist de segurança do projeto)

Se TECHNICAL-SOLUTION-SECURITY.md não existir → ERRO: "TECHNICAL-SOLUTION-SECURITY.md não encontrado. Execute primeiro o agente gerador de segurança."
Se TECHNICAL-SOLUTION-PRD.md não existir → ERRO: "TECHNICAL-SOLUTION-PRD.md não encontrado. O gate de segurança depende do TECHNICAL-SOLUTION-PRD.md para identificar os ativos a proteger."
Se TECHNICAL-SOLUTION-ARCHITECTURE.md não existir → ERRO: "TECHNICAL-SOLUTION-ARCHITECTURE.md não encontrado. O gate de segurança depende do TECHNICAL-SOLUTION-ARCHITECTURE.md para validar o pipeline de segurança."
Se GLOBAL-SECURITY.md não existir → ERRO CRÍTICO: "GLOBAL-SECURITY.md não encontrado. Esta é a referência normativa mestra — o gate não pode prosseguir sem ela."
```

### Passo 2 — Executar Dimensões de Validação

O gate avalia o TECHNICAL-SOLUTION-SECURITY.md em **5 dimensões** independentes. Para cada dimensão, atribuir um veredito: `APROVADO`, `RESSALVA` ou `REPROVADO`.

---

#### Dimensão 1: Conformidade com GLOBAL-SECURITY.md (Regras de Ouro e Checklist SDD)

**Regra especial:** Violação de qualquer Regra de Ouro = REPROVADO direto na dimensão, independentemente das demais verificações.

| # | Verificação | Critério |
|---|---|---|
| 1.1 | Regra de Ouro #1: Menor Privilégio | O TECHNICAL-SOLUTION-SECURITY.md documenta que todo novo endpoint/recurso é privado por padrão? O RBAC está mapeado com papéis e permissões explícitos? |
| 1.2 | Regra de Ouro #2: Zero Hardcoded Secrets | O TECHNICAL-SOLUTION-SECURITY.md define a estratégia de gestão de segredos (Vault, env vars, secrets manager)? Proíbe explicitamente hardcoded secrets? |
| 1.3 | Regra de Ouro #3: Não Confiar no Input | O TECHNICAL-SOLUTION-SECURITY.md especifica schema validation, sanitização contra XSS/SQL Injection e parameterized queries? |
| 1.4 | Checklist SDD: Autenticação/Autorização | Middleware de sessão, RBAC, proteção IDOR documentados com controles específicos? |
| 1.5 | Checklist SDD: Proteção de Dados | Criptografia em repouso (bcrypt/argon2), proteção contra vazamento em logs, campos mascarados? |
| 1.6 | Checklist SDD: Validação de Entradas | Schema validation, sanitização, parameterized queries — ferramentas e bibliotecas nomeadas? |
| 1.7 | Checklist SDD: Proteção de Infraestrutura | Rate limiting, CORS restrito, tratamento de erros seguro (sem stack traces)? |

#### Dimensão 2: Cobertura do Threat Model (STRIDE)

| # | Verificação | Critério |
|---|---|---|
| 2.1 | STRIDE completo | Os 6 tipos de ameaça (Spoofing, Tampering, Repudiation, Info Disclosure, DoS, Elevation of Privilege) estão analisados para cada componente/superfície? |
| 2.2 | Atores e agentes de ameaça | Estão identificados: atacante externo, insider, tenant malicioso, serviço comprometido? |
| 2.3 | Superfícies de ataque | APIs REST, filas de mensageria, banco de dados, logs, interfaces externas — todos listados? |
| 2.4 | Matriz de risco | Probabilidade × Impacto calculados para cada ameaça? Riscos priorizados (top 5-10)? |
| 2.5 | Controles mitigadores | Toda ameaça de alta severidade tem controle mitigador correspondente e verificável? |

#### Dimensão 3: Alinhamento com TECHNICAL-SOLUTION-ARCHITECTURE.md

| # | Verificação | Critério |
|---|---|---|
| 3.1 | Pipeline de segurança por requisição | O fluxo de autenticação/autorização (JWT → RBAC → IDOR → Tenant) descrito no TECHNICAL-SOLUTION-ARCHITECTURE.md está refletido no TECHNICAL-SOLUTION-SECURITY.md? |
| 3.2 | Cross-cutting concerns | Os aspectos AOP de segurança (logging, auditoria, rate limiting) do TECHNICAL-SOLUTION-ARCHITECTURE.md estão cobertos? |
| 3.3 | Tenant isolation | Se multi-tenant, a estratégia de isolamento (RLS, schema-per-tenant, discriminator column) está documentada e alinhada? |
| 3.4 | Estrutura de pacotes de segurança | Os pacotes/namespaces de segurança definidos na arquitetura (ex: `security.filter`, `security.config`) estão referenciados? |

#### Dimensão 4: Proteção de Dados e Privacidade

| # | Verificação | Critério |
|---|---|---|
| 4.1 | Criptografia em repouso | Algoritmos e escopo definidos (BD, backups, logs)? Hashing de senhas com bcrypt/argon2? |
| 4.2 | Criptografia em trânsito | TLS 1.3 ou superior? mTLS entre serviços internos? Certificate pinning onde aplicável? |
| 4.3 | Mascaramento e retenção | Campos sensíveis (CPF, cartão, tokens) mascarados em logs e APIs? Política de retenção e expurgo definida? |
| 4.4 | Conformidade regulatória | LGPD, PCI, SOC2 — requisitos aplicáveis mapeados com controles correspondentes? |

#### Dimensão 5: Acionabilidade e Cobertura Operacional

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Controles acionáveis | Cada controle responde: o quê, quem implementa, como verificar, qual a evidência? |
| 5.2 | Pipeline DevSecOps | SAST, Secret Scanning, Dependency Scanning — ferramentas nomeadas, frequência definida, gates no CI/CD? |
| 5.3 | OWASP Top 10 | As 10 categorias cobertas com status individual (✅ Mitigado, ⚠️ Parcial, ❌ Exposto)? |
| 5.4 | Resposta a incidentes | Canais de reporte (conforme GLOBAL-SECURITY.md), procedimentos de contenção e escalação? |
| 5.5 | Security ADRs | Pelo menos 3 decisões de segurança documentadas com ID, contexto, alternativas e justificativa? |

---

### Passo 3 — Calcular Veredito

**Veredito por dimensão** (calculado sobre as verificações daquela dimensão):

| Condição | Veredito |
|---|---|
| 100% das verificações OK | ✅ APROVADO |
| ≥ 75% das verificações OK | ⚠️ RESSALVA |
| < 75% das verificações OK | ❌ REPROVADO |

**Regra especial — Regras de Ouro:** Se QUALQUER verificação das Regras de Ouro (1.1, 1.2, 1.3) falhar, a Dimensão 1 é REPROVADA automaticamente, independentemente da porcentagem.

**Veredito final do gate:**

| Condição | Veredito | Bloqueia Fase 4 (TECHNICAL-SOLUTION-SPECS.md)? |
|---|---|---|
| 5 dimensões APROVADAS | ✅ APROVADO (COMPLIANCE) | Não |
| ≥1 RESSALVA, 0 REPROVADA | ⚠️ RESSALVA (PRÉ-COMPLIANCE) | Não (com observações) |
| ≥1 REPROVADA | ❌ REPROVADO (NÃO-COMPLIANCE) | SIM |

### Passo 4 — Gerar Relatório de Falha (se REPROVADO)

Gerar: `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SECURITY_SCOPE_FAIL_REPORT.md`

```markdown
# 🚨 Relatório de Falha — Gate de Segurança (TECHNICAL-SOLUTION-SECURITY.md)

* **Data e Hora:** [AAAA-MM-DD HH:MM:SS]
* **Projeto:** {PROJECT_NAME}
* **Solução:** {SOLUTION_NAME}
* **Gate:** GATE-SECURITY-SCOPE
* **Veredito:** ❌ REPROVADO

---

## 📊 Sumário Executivo

| Dimensão | Veredito | Verificações OK/Total | % |
|---|---|---|---|
| 1. Conformidade GLOBAL-SECURITY.md | [APROVADO/RESSALVA/REPROVADO] | [N]/[M] | [X]% |
| 2. Cobertura Threat Model (STRIDE) | [APROVADO/RESSALVA/REPROVADO] | [N]/[M] | [X]% |
| 3. Alinhamento com TECHNICAL-SOLUTION-ARCHITECTURE.md | [APROVADO/RESSALVA/REPROVADO] | [N]/[M] | [X]% |
| 4. Proteção de Dados e Privacidade | [APROVADO/RESSALVA/REPROVADO] | [N]/[M] | [X]% |
| 5. Acionabilidade e Cobertura Operacional | [APROVADO/RESSALVA/REPROVADO] | [N]/[M] | [X]% |

**Total de Não-Conformidades:** [N]
**Regras de Ouro Violadas:** [Sim/Não]

---

## 🔴 Não-Conformidades por Dimensão

### Dimensão 1: Conformidade com GLOBAL-SECURITY.md
[Para cada verificação reprovada:]
- **NC-1.1:** [Descrição do problema]
  - **Gravidade:** [P0/P1/P2/P3]
  - **Evidência:** [Trecho ou omissão no TECHNICAL-SOLUTION-SECURITY.md]
  - **Ação Esperada:** [O que precisa ser corrigido]

### Dimensão 2: Cobertura do Threat Model
[...]

### Dimensão 3: Alinhamento com TECHNICAL-SOLUTION-ARCHITECTURE.md
[...]

### Dimensão 4: Proteção de Dados e Privacidade
[...]

### Dimensão 5: Acionabilidade e Cobertura Operacional
[...]

---

## 🛡️ Regras de Ouro (Atenção Especial)

| Regra | Status | Observação |
|---|---|---|
| 1. Menor Privilégio | [✅ Cumprida / ❌ Violada] | [Detalhe] |
| 2. Zero Hardcoded Secrets | [✅ Cumprida / ❌ Violada] | [Detalhe] |
| 3. Não Confiar no Input | [✅ Cumprida / ❌ Violada] | [Detalhe] |

---

## 🛠️ Recomendações para Correção

1. [Recomendação priorizada 1]
2. [Recomendação priorizada 2]
...

## 📋 Próximos Passos

1. Executar `PROMPT-FIX-SECURITY-TECHNICAL_SOLUTION.md` com este relatório
2. Revalidar com `PROMPT-GATE-SECURITY-TECHNICAL_SOLUTION.md`
3. Se aprovado, avançar para Fase 4 (TECHNICAL-SOLUTION-SPECS.md)
```

### Passo 5 — Validar Relatório Gerado

| # | Verificação |
|---|---|
| 5.1 | Relatório gravado no caminho correto |
| 5.2 | Todas as 5 dimensões têm veredito |
| 5.3 | NCs listadas com gravidade, evidência e ação esperada |
| 5.4 | Regras de Ouro com status individual |
| 5.5 | Recomendações priorizadas e acionáveis |

---

## Skills Utilizados

> **📌 Nota sobre Skills:** A tabela abaixo lista os skills **recomendados** para a validação do gate de segurança. O agente validador deve usá-los como ponto de partida, mas tem autonomia para selecionar outros skills identificados como mais aderentes às necessidades específicas da solução, stack tecnológica ou domínio de negócio. A ordem sugerida reflete as dependências lógicas entre as atividades de validação.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `security-auditor` | Auditar conformidade com GLOBAL-SECURITY.md e regras de ouro (Dimensão 1) | Conformidade |
| 2 | `security-audit` | Auditoria complementar de gaps de segurança e verificação de aderência normativa | Conformidade |
| 3 | `skill-security-auditor` | Validação cruzada de requisitos de segurança entre TECHNICAL-SOLUTION-SECURITY.md, TECHNICAL-SOLUTION-PRD.md e GLOBAL-SECURITY.md | Conformidade |
| 4 | `gap-analysis` | Identificar lacunas entre TECHNICAL-SOLUTION-SECURITY.md e documentos de referência (TECHNICAL-SOLUTION-PRD.md, TECHNICAL-SOLUTION-ARCHITECTURE.md) | Análise |
| 5 | `threat-modeling-expert` | Validar completude e rigor do threat model STRIDE: atores, superfícies, matriz de risco (Dimensão 2) | Threat Model |
| 6 | `senior-security` | Validação sênior: revisar criticidade, controles estratégicos e cobertura de ameaças de alto impacto | Estratégia |
| 7 | `security-best-practices` | Verificar aderência dos controles às boas práticas de segurança para a stack `{STACK}` | Boas Práticas |
| 8 | `api-security-best-practices` | Validar controles de segurança de API: rate limiting, CORS, input validation, JWT/OAuth2 (Dimensão 5) | API Security |
| 9 | `api-security-testing` | Verificar se o plano de testes de segurança de API está adequadamente coberto no pipeline | API Security |
| 10 | `security-scanning-security-sast` | Validar configuração do SAST no pipeline DevSecOps: ferramenta nomeada, frequência, gates (Dimensão 5) | Pipeline |
| 11 | `security-scanning-security-dependencies` | Verificar se a gestão de dependências (SCA) está documentada com ferramenta e política (Dimensão 5) | Pipeline |
| 12 | `security-scanning-security-hardening` | Validar hardening de infraestrutura: secrets management, containers, WAF, patching (Dimensão 5) | Infra |
| 13 | `security-reviewer` | Revisão final de cobertura e acionabilidade dos controles; verificar checklist SDD completo | Revisão |
| 14 | `security-review` | Revisão complementar: consistência cross-documento, aderência ao TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md | Revisão |
| 15 | `engineering-skills` | Validar que todos os controles são acionáveis: o quê, quem implementa, como verificar, evidência | Qualidade |

> **🔄 Flexibilidade:** Se durante a validação o agente identificar que um skill diferente dos listados acima é mais adequado para uma verificação específica (ex: um skill de conformidade regulatória como LGPD/PCI para a Dimensão 4, um skill de segurança em nuvem para revisão de infraestrutura, ou um skill especializado na stack `{STACK}`), ele deve substituí-lo e justificar a escolha no relatório de falha (`SECURITY_SCOPE_FAIL_REPORT.md`).

---

## Observações

- As 3 Regras de Ouro do GLOBAL-SECURITY.md são **inegociáveis**. Violação de qualquer uma = REPROVAÇÃO automática da Dimensão 1.
- O gate é binário com escape de ressalva: APROVADO ou RESSALVA permitem avançar; REPROVADO bloqueia a Fase 4 (TECHNICAL-SOLUTION-SPECS.md).
- O relatório de falha (`SECURITY_SCOPE_FAIL_REPORT.md`) é o contrato de correção — o `PROMPT-FIX-SECURITY-TECHNICAL_SOLUTION.md` o utiliza como entrada única para as correções.
- Se TECHNICAL-SOLUTION-PRD.md ou TECHNICAL-SOLUTION-ARCHITECTURE.md forem alterados após o TECHNICAL-SOLUTION-SECURITY.md ter sido aprovado, o gate DEVE ser reexecutado (efeito cascata).

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 21/07/2026 | Criação inicial: gate de segurança com 5 dimensões, regras de ouro inegociáveis e relatório de falha estruturado | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência: security-auditor, security-audit, skill-security-auditor, gap-analysis, threat-modeling-expert, senior-security, security-best-practices, api-security-best-practices, api-security-testing, security-scanning-security-sast, security-scanning-security-dependencies, security-scanning-security-hardening, security-reviewer, security-review, engineering-skills. Outros skills podem ser utilizados conforme aderência à necessidade específica.*
