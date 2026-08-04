# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md` — a **definição de estratégia de testes do projeto** que especifica a pirâmide de testes, automação, performance, segurança, ambientes e quality gates. Este documento é a referência normativa para toda a atividade de testes no projeto.

**Este documento é independente de tecnologias específicas de teste.** Durante a análise da stack do projeto, identifique as ferramentas e frameworks de teste utilizados e busque skills relacionados. Caso não encontre skills específicos, utilize skills generalistas de qualidade e teste, utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

**Relação com DEVOPS-SRE-DEFINITION:** Enquanto o DEVOPS-SRE-DEFINITION (F10) define o pipeline CI/CD e a esteira de deploy, este documento detalha **o que é testado, como é testado e quais critérios determinam a qualidade** em cada etapa desse pipeline.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7)

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
Validar que `{UPSTREAM_DISCOVERY_PATH}` existe e contém os artefatos upstream.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos upstream:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` — Definição de Arquitetura (F2) — contratos de API
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` — Definição de Segurança (F3) — testes de segurança
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` — Definição DevOps/SRE (F5) — pipeline e SLOs
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` — Definição Infra/Cloud (F7) — ambientes de teste

### Passo 2 — Invocar Skills Especializadas
Invocar skills de QA, automação de testes, performance, segurança e qualidade para projetar a estratégia de testes completa.

### Passo 2.5 — Apresentar Skills, Capacidade e Stack para Validação Humana

Avaliar e apresentar ao usuário para validação:

1. **PROJECT-TEAM-SKILLS-MAP:** Skills identificados como necessários para implementar a solução nesta disciplina.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

2. **PROJECT-TEAM-CAPACITY:** Capacidade estimada do time nesta disciplina (ex: 2 seniores, 3 plenos).
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

3. **PROJECT-STACK:** Tecnologias identificadas para esta disciplina. Verificar conformidade com `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão corporativo DEVEM ser listadas com justificativa técnica e requerem aprovação explícita do usuário.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-TEST-STRATEGY-DEFINITION.md` com:

1. **Pirâmide de Testes** — unitários (cobertura esperada por camada), integração (contratos, banco, mensageria), E2E (fluxos críticos), aceitação (BDD/Cucumber), percentuais de cobertura por tipo
2. **Estratégia de Automação** — frameworks por linguagem (JUnit/Jest/Pytest para unitários, Playwright/Cypress para E2E, REST Assured/Postman para API), pipeline de testes no CI/CD, paralelismo
3. **Testes de Performance** — carga (k6/JMeter/Gatling), stress (picos), soak (longa duração), benchmarks (comparação entre versões), thresholds alinhados com SLOs do DevOps SRE
4. **Testes de Segurança** — SAST (SonarQube/Semgrep) no pipeline, SCA (OWASP Dependency Check/Dependabot), DAST (OWASP ZAP/Burp), penetration testing (manual/automated), alinhados com SECURITY-DEFINITION
5. **Ambientes de Teste** — isolamento (dev/staging/QA), dados de teste (anonimizados, LGPD), massa de dados (sintética vs. produção), provisionamento via IaC
6. **Quality Gates** — critérios mínimos para passar em cada nível (unit: >=80% cobertura, integração: contratos verificados, E2E: fluxos críticos verdes), mutation testing (Stryker/PIT), cobertura por solução
7. **Ferramentas** — por solução e linguagem, justificativa de escolha, integração com CI/CD

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md`

```markdown
# DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md
## Fase 6 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | {DATA_ATUAL} |
| **Autor** | QA Engineer / Test Specialist |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md) — PRD Discovery-Level (F1)
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md) — Definição de Arquitetura (F2)

---

## 1. Pirâmide de Testes — Visão Macro
- Diagrama ASCII da pirâmide com pesos (ex: 60% unitários, 30% integração, 10% E2E)
- **1.1 Distribuição por Camada:** Tabela: Camada | Peso | O que cobre | Meta de Cobertura

## 2. Estratégia por Tipo de Teste
- **2.1 Testes Unitários:** Tabela: Solução | Ferramenta | Foco | Pipeline Stage
- **2.2 Testes de Integração:** Tabela: Tipo | Ferramenta | Escopo
- **2.3 Testes End-to-End (E2E):** Tabela: Fluxo Crítico | Épico | Cenários | Ferramenta
- **2.4 Testes de Performance:** Tabela: Cenário | Ferramenta | Target
- **2.5 Testes de Segurança:** Tabela: Tipo | Ferramenta | Pipeline Stage

## 3. Ambientes de Teste
- Tabela: Ambiente | Propósito | Dados | Acesso
- **3.1 Quality Gates por Ambiente:** Tabela: Gate | Dev | Staging | Prod

## 4. Métricas de Qualidade
- Tabela: Métrica | Target | Medição

## 5. Riscos e Estimativa de Esforço
- **5.1 Riscos de Testes:** Tabela: ID | Risco | Prob. | Impacto | Mitigação
- **5.2 Estimativa de Esforço:** Tabela: Atividade | Complexidade | Esforço (dias) | Responsável

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Test Strategy Definition Discovery-Level | QA Engineer / Test Specialist |
```

### Passo 4 — Validação Pós-Geração

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `senior-qa` | Supervisão sênior de QA | QA |
| 2 | `qa-test-planner` | Planejamento da estratégia de testes | Planejamento |
| 3 | `test-strategy-design` | Design da pirâmide e automação | Estratégia |
| 4 | `tdd-guide` | Diretrizes de TDD e qualidade | TDD |
| 5 | `e2e-testing-patterns` | Padrões de testes E2E | E2E |
| 6 | `k6-load-testing` | Testes de carga com k6 | Performance |
| 7 | `playwright-expert` | Automação de testes E2E com Playwright | Automação |
| 8 | `testing-patterns` | Padrões gerais de teste | Testes |
| 9 | `unit-testing-test-generate` | Geração de testes unitários | Unitários |
| 10 | `security-testing` | Testes de segurança (SAST/DAST) | Segurança |
| 11 | `performance-testing-review-ai-review` | Revisão de testes de performance | Performance |
| 12 | `mermaid-expert` | Diagramas da pirâmide e pipeline | Diagramas |
| 13 | `documentation-writer` | Redigir o Test Strategy Definition | Documentação |

**Skills generalistas de teste e qualidade (sempre aplicáveis):**
- `engineering-skills`, `engineering-advanced-skills`
- `senior-qa`, `testing-patterns`, `test-strategy-design`
- `test-master`, `testing-qa`, `test-driven-development`
- `e2e-testing`, `e2e-testing-patterns`
- `test-automator`, `test-case-creation`
- `qa`, `qa-test-planner`, `acceptance-criteria`
- `performance-testing-review-multi-agent-review`

**Skills tecnológicos de teste (condicionais — buscar ao identificar a stack):**
- Ao identificar um framework ou ferramenta de teste específica durante a análise da stack, busque skills relacionados a essa tecnologia para aprimorar as especificações de teste
- Caso não encontre skills específicos para a ferramenta identificada, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da definição de estratégia de testes | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` | Definição de Arquitetura (F2) — contratos de API |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SECURITY-DEFINITION.md` | Definição de Segurança (F3) — testes de segurança |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md` | Definição DevOps/SRE (F5) — pipeline e SLOs |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md` | Definição Infra/Cloud (F7) — ambientes |
| 5 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 6 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (se fornecidos) |
| 7 | `{PROJECT-TEAM-CAPACITY}` | Capacidade do time (se fornecida) |
| 8 | `{PROJECT-STACK}` | Stack tecnológica (validar contra padrões corporativos) |
| 9 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
