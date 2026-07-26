# PROMPT-GENERATE-SECURITY-ARTEFACT

## Contexto

Este prompt orquestra skills especializadas em segurança da informação, threat modeling e secure coding para gerar ou revisar o artefato `TECHNICAL-SOLUTION-SECURITY.md` na pasta de especificações de uma solução técnica.

O artefato gerado é o **plano de segurança específico da solução** — ele traduz os requisitos de segurança do TECHNICAL-SOLUTION-PRD.md e as decisões arquiteturais do TECHNICAL-SOLUTION-ARCHITECTURE.md em controles concretos, verificáveis e acionáveis, tendo como referência normativa obrigatória o `GLOBAL-SECURITY.md` (Política e Checklist de Segurança do Projeto).

**Princípio fundamental:** O TECHNICAL-SOLUTION-SECURITY.md não repete o GLOBAL-SECURITY.md. Ele o **especializa** para a solução específica — aplicando as regras de ouro, o checklist SDD e as ferramentas de verificação ao contexto concreto da stack, das integrações e do modelo de ameaças da solução.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço ou frontend) | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |
| `{SECURITY_GLOBAL}` | Caminho absoluto para o documento de política de segurança global (GLOBAL-SECURITY.md) | `/home/user/work/.specs/security/GLOBAL-SECURITY.md` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Antes de qualquer ação, verificar se TODOS os 6 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

Validar também:
- `{SECURITY_GLOBAL}` aponta para um arquivo existente? Se não, alertar: "GLOBAL-SECURITY.md não encontrado no caminho informado. A geração prosseguirá sem a referência normativa global."
- `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-PRD.md` existe? Se não, alertar: "TECHNICAL-SOLUTION-PRD.md não encontrado. Recomenda-se gerar o TECHNICAL-SOLUTION-PRD.md (Fase 1) antes do TECHNICAL-SOLUTION-SECURITY.md."
- `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md` existe? Se não, alertar: "TECHNICAL-SOLUTION-ARCHITECTURE.md não encontrado. Recomenda-se gerar o TECHNICAL-SOLUTION-ARCHITECTURE.md (Fase 2) antes do TECHNICAL-SOLUTION-SECURITY.md."

### Passo 1 — Verificar e Preparar a Estrutura de Pastas

```
Verificar se existe: {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    │
    ├── NÃO existe → Criar a pasta (mkdir -p)
    │
    └── SIM, existe →
            │
            ├── Carregar documentos de referência obrigatórios:
            │     ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-PRD.md
            │     │   (escopo, entidades, funcionalidades, NFRs de segurança declarados)
            │     ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md
            │     │   (pipeline de segurança por requisição, cross-cutting concerns,
            │     │    RBAC, multi-tenancy, camadas e pontos de entrada)
            │     └── {SECURITY_GLOBAL} (GLOBAL-SECURITY.md)
            │         (Regras de Ouro, Checklist SDD, ferramentas de verificação)
            │
            ├── Verificar se TECHNICAL-SOLUTION-SECURITY.md já existe:
            │     ├── SIM → Ler versão atual, preservar changelog
            │     └── NÃO → Gerar do zero
            │
            └── Carregar documentos complementares (se existirem):
                  ├── {PROJECT_PATH}/TECHNICAL-PLAN.md (stack, ERD, ADRs)
                  └── {PROJECT_PATH}/TECHNICAL-SOLUTION-ARCHITECTURE.md (decisões arquiteturais do projeto)
```

### Passo 2 — Invocar Skills Especializadas

> **📌 Nota sobre Skills:** A tabela abaixo lista os skills **recomendados** para esta fase. O agente deve usá-los como ponto de partida, mas tem autonomia para selecionar outros skills identificados como mais aderentes às necessidades específicas da solução, stack tecnológica ou domínio de negócio. A ordem sugerida reflete as dependências lógicas entre as atividades (auditoria → threat model → controles → revisão → documentação).

Invocar preferencialmente nesta ordem:

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `security-auditor` | Auditar TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md extraindo todos os requisitos e decisões de segurança implícitos e explícitos | Auditoria |
| 2 | `security-audit` | Auditoria complementar de gaps de segurança nos artefatos de entrada e verificação de conformidade normativa | Auditoria |
| 3 | `skill-security-auditor` | Validação cruzada de requisitos de segurança entre TECHNICAL-SOLUTION-PRD.md, TECHNICAL-SOLUTION-ARCHITECTURE.md e GLOBAL-SECURITY.md | Auditoria |
| 4 | `threat-modeling-expert` | Construir o modelo de ameaças (STRIDE) específico da solução: atores, superfícies de ataque, fluxos de dados sensíveis | Threat Model |
| 5 | `senior-security` | Supervisão sênior: validar criticidade dos ativos, revisar threat model e definir controles estratégicos | Estratégia |
| 6 | `security-best-practices` | Mapear as melhores práticas de segurança aplicáveis à stack `{STACK}` (ex: Spring Security, OWASP para Java) | Boas Práticas |
| 7 | `api-security-best-practices` | Especificar controles de segurança para APIs REST: rate limiting, CORS, input validation, autenticação JWT/OAuth2 | API Security |
| 8 | `api-security-testing` | Planejar estratégia de testes de segurança de API: DAST, fuzzing, testes de penetração de endpoints | API Security |
| 9 | `security-scanning-security-sast` | Definir configuração de SAST no pipeline DevSecOps: ferramenta, regras, frequência, gates no CI/CD | Pipeline |
| 10 | `security-scanning-security-dependencies` | Definir gestão de dependências e SCA (Software Composition Analysis): ferramenta, política de atualização, SLSA | Pipeline |
| 11 | `security-scanning-security-hardening` | Especificar hardening de infraestrutura: containers, secrets management, patching, WAF, security groups | Infra |
| 12 | `security-reviewer` | Revisar os controles propostos contra o GLOBAL-SECURITY.md e verificar cobertura das regras de ouro e checklist SDD | Revisão |
| 13 | `security-review` | Revisão complementar de completude, consistência cross-documento e aderência ao TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md | Revisão |
| 14 | `engineering-skills` | Garantir que todos os controles são acionáveis e verificáveis pela equipe de engenharia (o quê, quem, como verificar) | Qualidade |
| 15 | `documentation-writer` | Redigir o TECHNICAL-SOLUTION-SECURITY.md consolidado com as 12 seções obrigatórias, changelog e referências cruzadas | Documentação |

> **🔄 Flexibilidade:** Se durante a execução o agente identificar que um skill diferente dos listados acima é mais adequado para uma atividade específica (ex: um skill especializado na stack `{STACK}`, um skill de conformidade regulatória como LGPD/PCI, ou um skill de segurança em nuvem), ele deve substituí-lo e justificar a escolha no changelog do TECHNICAL-SOLUTION-SECURITY.md.

### Passo 3 — Gerar o Arquivo TECHNICAL-SOLUTION-SECURITY.md

Gerar o arquivo em: `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-SECURITY.md`

#### Template de Seções Obrigatórias

O TECHNICAL-SOLUTION-SECURITY.md DEVE conter as 12 seções abaixo:

---

**1. Visão Geral de Segurança**
- Contexto da solução e sensibilidade dos dados tratados
- Nível de criticidade (baixo, médio, alto, crítico)
- Resumo dos controles principais

**2. Threat Model (STRIDE)**
- **Atores e Agentes de Ameaça:** quem pode atacar o sistema?
- **Superfícies de Ataque:** APIs, endpoints, filas, banco de dados, logs, interfaces externas
- **Fluxos de Dados Sensíveis:** diagrama ou tabela de quais dados trafegam entre componentes
- **Matriz STRIDE:** para cada componente/superfície, classificar Spoofing, Tampering, Repudiation, Information Disclosure, Denial of Service, Elevation of Privilege
- **Riscos Priorizados:** top 5-10 riscos com severidade e probabilidade

**3. Autenticação e Autorização**
- Mecanismo de autenticação (JWT, OAuth2, SAML, API Key)
- Política de senhas e credenciais
- RBAC: matriz de papéis × permissões (admin, user, manager, auditor, etc.)
- Row-Level Security (RLS) / Multi-Tenant Isolation
- Proteção contra IDOR (Insecure Direct Object References)

**4. Proteção de Dados e Privacidade**
- Criptografia em repouso (algoritmos, escopo — BD, backups, logs)
- Criptografia em trânsito (TLS 1.3, mTLS, cert pinning)
- Dados sensíveis: política de mascaramento, retenção e expurgo
- Conformidade regulatória aplicável (LGPD, PCI, SOC2)

**5. Segurança de API e Comunicação**
- Rate Limiting por endpoint e por tenant
- Política de CORS (domínios autorizados, métodos, headers)
- Input Validation: sanitização contra XSS, SQL Injection, Command Injection
- Content Security Policy (CSP) e headers de segurança HTTP
- Proteção contra CSRF

**6. Cobertura OWASP Top 10**
- Para cada uma das 10 categorias OWASP: status (✅ Mitigado, ⚠️ Parcial, ❌ Exposto) e controles aplicados

**7. Gestão de Dependências (SCA — Software Composition Analysis)**
- Ferramenta de análise de dependências utilizada (ex: `npm audit`, `pip audit`, Snyk, Dependabot)
- Política de atualização de dependências com vulnerabilidades conhecidas
- SLSA framework / Supply Chain Levels
- Assinatura e verificação de artefatos

**8. Pipeline de Segurança (DevSecOps)**
- SAST (Static Application Security Testing): ferramenta, frequência, regras
- Secret Scanning: ferramenta, padrões detectados, prevenção de hardcoded secrets
- DAST (Dynamic Application Security Testing): ferramenta, escopo
- Container/Image Scanning (se aplicável)
- Git hooks e pre-commit checks de segurança

**9. Segurança de Infraestrutura**
- Configuração de rede (VPC, security groups, WAF)
- Gestão de segredos (vault, secrets manager, variáveis de ambiente)
- Hardening de containers e imagens base
- Política de patching e atualização de SO

**10. Checklist de Verificação (SDD — Security-Driven Development)**
- Checklist de segurança para desenvolvimento derivado do GLOBAL-SECURITY.md
- Validar: autenticação, autorização, proteção de dados, validação de entradas, proteção de infraestrutura
- Itens marcados como ✅ Aplicável, ⚠️ Parcial, ❌ Não Aplicável

**11. Resposta a Incidentes e Monitoramento**
- Alertas de segurança configurados (WAF, IDS, anomaly detection)
- Logs de auditoria: o que é logado, onde, por quanto tempo
- Plano de resposta a incidentes (contatos, procedimentos, escalação)
- Política de divulgação de vulnerabilidades

**12. Changelog**
- Registro de alterações do documento com versão, data, alteração e autor

---

### Passo 4 — Validação Pós-Geração

Antes de reportar sucesso, executar as verificações abaixo:

| # | Verificação | Critério de Sucesso |
|---|---|---|
| 4.1 | Arquivo no caminho correto | `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-SECURITY.md` existe |
| 4.2 | 12 seções obrigatórias | Todas as 12 seções estão presentes com conteúdo |
| 4.3 | Referência ao GLOBAL-SECURITY.md | O arquivo cita explicitamente o GLOBAL-SECURITY.md como referência normativa |
| 4.4 | Threat Model preenchido | Seção 2 contém matriz STRIDE com pelo menos 5 componentes |
| 4.5 | Matriz RBAC | Seção 3 contém tabela de papéis × permissões |
| 4.6 | OWASP Top 10 coberto | Seção 6 cobre as 10 categorias com status individual |
| 4.7 | Controles acionáveis | Cada controle é descrito com: o quê, quem implementa, como verificar |
| 4.8 | Stack mencionada | O documento referencia a stack `{STACK}` nas recomendações |
| 4.9 | Integração com TECHNICAL-SOLUTION-ARCHITECTURE.md | Seção 8 referencia o pipeline de segurança definido no TECHNICAL-SOLUTION-ARCHITECTURE.md |
| 4.10 | Checklist SDD | Seção 10 contém checklist derivado do GLOBAL-SECURITY.md |
| 4.11 | Changelog atualizado | Seção 12 registra a geração/revisão atual |

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 21/07/2026 | Criação inicial: prompt gerador de TECHNICAL-SOLUTION-SECURITY.md com 12 seções obrigatórias, integração com GLOBAL-SECURITY.md e threat model STRIDE | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência: security-auditor, security-audit, skill-security-auditor, threat-modeling-expert, senior-security, security-best-practices, api-security-best-practices, api-security-testing, security-scanning-security-sast, security-scanning-security-dependencies, security-scanning-security-hardening, security-reviewer, security-review, engineering-skills, documentation-writer. Outros skills podem ser utilizados conforme aderência à necessidade específica.*
