# PROMPT-GENERATE-SECURITY-ARTEFACT

## Contexto

Este prompt orquestra skills especializadas em segurança da informação, threat modeling e secure coding para gerar ou revisar o artefato `SECURITY.md` na pasta de especificações de uma solução técnica.

O artefato gerado é o **plano de segurança específico da solução** — ele traduz os requisitos de segurança do PRD.md e as decisões arquiteturais do ARCHITECTURE.md em controles concretos, verificáveis e acionáveis, tendo como referência normativa obrigatória o `GLOBAL-SECURITY.md` (Política e Checklist de Segurança do Projeto).

**Princípio fundamental:** O SECURITY.md não repete o GLOBAL-SECURITY.md. Ele o **especializa** para a solução específica — aplicando as regras de ouro, o checklist SDD e as ferramentas de verificação ao contexto concreto da stack, das integrações e do modelo de ameaças da solução.

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
- `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md` existe? Se não, alertar: "PRD.md não encontrado. Recomenda-se gerar o PRD.md (Fase 1) antes do SECURITY.md."
- `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md` existe? Se não, alertar: "ARCHITECTURE.md não encontrado. Recomenda-se gerar o ARCHITECTURE.md (Fase 2) antes do SECURITY.md."

### Passo 1 — Verificar e Preparar a Estrutura de Pastas

```
Verificar se existe: {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    │
    ├── NÃO existe → Criar a pasta (mkdir -p)
    │
    └── SIM, existe →
            │
            ├── Carregar documentos de referência obrigatórios:
            │     ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md
            │     │   (escopo, entidades, funcionalidades, NFRs de segurança declarados)
            │     ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md
            │     │   (pipeline de segurança por requisição, cross-cutting concerns,
            │     │    RBAC, multi-tenancy, camadas e pontos de entrada)
            │     └── {SECURITY_GLOBAL} (GLOBAL-SECURITY.md)
            │         (Regras de Ouro, Checklist SDD, ferramentas de verificação)
            │
            ├── Verificar se SECURITY.md já existe:
            │     ├── SIM → Ler versão atual, preservar changelog
            │     └── NÃO → Gerar do zero
            │
            └── Carregar documentos complementares (se existirem):
                  ├── {PROJECT_PATH}/TECHNICAL-PLAN.md (stack, ERD, ADRs)
                  └── {PROJECT_PATH}/ARCHITECTURE.md (decisões arquiteturais do projeto)
```

### Passo 2 — Invocar Skills Especializadas

Invocar obrigatoriamente nesta ordem:

| Ordem | Skill | Propósito |
|---|---|---|
| 1 | `security-auditor` | Auditar PRD.md e ARCHITECTURE.md extraindo todos os requisitos e decisões de segurança implícitos e explícitos |
| 2 | `threat-modeling-expert` | Construir o modelo de ameaças (STRIDE) específico da solução: atores, superfícies de ataque, fluxos de dados sensíveis |
| 3 | `security-best-practices` | Mapear as melhores práticas de segurança aplicáveis à stack `{STACK}` (ex: Spring Security, OWASP para Java) |
| 4 | `security-reviewer` | Revisar os controles propostos contra o GLOBAL-SECURITY.md e verificar cobertura das regras de ouro |
| 5 | `documentation-writer` | Redigir o SECURITY.md consolidado com todas as seções obrigatórias |

### Passo 3 — Gerar o Arquivo SECURITY.md

Gerar o arquivo em: `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SECURITY.md`

#### Template de Seções Obrigatórias

O SECURITY.md DEVE conter as 12 seções abaixo:

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
| 4.1 | Arquivo no caminho correto | `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SECURITY.md` existe |
| 4.2 | 12 seções obrigatórias | Todas as 12 seções estão presentes com conteúdo |
| 4.3 | Referência ao GLOBAL-SECURITY.md | O arquivo cita explicitamente o GLOBAL-SECURITY.md como referência normativa |
| 4.4 | Threat Model preenchido | Seção 2 contém matriz STRIDE com pelo menos 5 componentes |
| 4.5 | Matriz RBAC | Seção 3 contém tabela de papéis × permissões |
| 4.6 | OWASP Top 10 coberto | Seção 6 cobre as 10 categorias com status individual |
| 4.7 | Controles acionáveis | Cada controle é descrito com: o quê, quem implementa, como verificar |
| 4.8 | Stack mencionada | O documento referencia a stack `{STACK}` nas recomendações |
| 4.9 | Integração com ARCHITECTURE.md | Seção 8 referencia o pipeline de segurança definido no ARCHITECTURE.md |
| 4.10 | Checklist SDD | Seção 10 contém checklist derivado do GLOBAL-SECURITY.md |
| 4.11 | Changelog atualizado | Seção 12 registra a geração/revisão atual |

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 21/07/2026 | Criação inicial: prompt gerador de SECURITY.md com 12 seções obrigatórias, integração com GLOBAL-SECURITY.md e threat model STRIDE | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Foram utilizados os skills: security-auditor, threat-modeling-expert, security-best-practices, security-reviewer, documentation-writer.*
