# PROMPT-GENERATE-510-TEST-STRATEGY-DEFINITION

## Contexto

Este prompt gera o artefato `510-TEST-STRATEGY-DEFINITION.md` — a **definição de estratégia de testes do projeto** que especifica a pirâmide de testes, automação, performance, segurança, ambientes e quality gates. Este documento é a referência normativa para toda a atividade de testes no projeto.

**Relação com DEVOPS-SRE-DEFINITION:** Enquanto o DEVOPS-SRE-DEFINITION (F10) define o pipeline CI/CD e a esteira de deploy, este documento detalha **o que é testado, como é testado e quais critérios determinam a qualidade** em cada etapa desse pipeline.

**Inputs upstream:** `470-ARCHITECTURE-DEFINITION.md` (Fase 7) + `480-SECURITY-DEFINITION.md` (Fase 8) + `500-DEVOPS-SRE-DEFINITION.md` (Fase 10) + `520-INFRA-CLOUD-DEFINITION.md` (Fase 12 — referência para ambientes de teste).

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler Architecture Definition (F7 — topologia de serviços, contratos de API para testes de integração), Security Definition (F8 — requisitos para testes de segurança), DevOps SRE (F10 — pipeline CI/CD, SLOs para testes de performance), Infra Cloud (F12 — ambientes de teste disponíveis), ADRs globais.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de QA, automação de testes, performance, segurança e qualidade para projetar a estratégia de testes completa.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/510-TEST-STRATEGY-DEFINITION.md` com:

1. **Pirâmide de Testes** — unitários (cobertura esperada por camada), integração (contratos, banco, mensageria), E2E (fluxos críticos), aceitação (BDD/Cucumber), percentuais de cobertura por tipo
2. **Estratégia de Automação** — frameworks por linguagem (JUnit/Jest/Pytest para unitários, Playwright/Cypress para E2E, REST Assured/Postman para API), pipeline de testes no CI/CD, paralelismo
3. **Testes de Performance** — carga (k6/JMeter/Gatling), stress (picos), soak (longa duração), benchmarks (comparação entre versões), thresholds alinhados com SLOs do DevOps SRE
4. **Testes de Segurança** — SAST (SonarQube/Semgrep) no pipeline, SCA (OWASP Dependency Check/Dependabot), DAST (OWASP ZAP/Burp), penetration testing (manual/automated), alinhados com SECURITY-DEFINITION
5. **Ambientes de Teste** — isolamento (dev/staging/QA), dados de teste (anonimizados, LGPD), massa de dados (sintética vs. produção), provisionamento via IaC
6. **Quality Gates** — critérios mínimos para passar em cada nível (unit: >=80% cobertura, integração: contratos verificados, E2E: fluxos críticos verdes), mutation testing (Stryker/PIT), cobertura por solução
7. **Ferramentas** — por solução e linguagem, justificativa de escolha, integração com CI/CD

### Passo 4 — Validação Pós-Geração
Verificar: pirâmide definida com % de cobertura por tipo, ferramentas especificadas por solução/linguagem, quality gates documentados com métricas objetivas, testes de segurança alinhados com SECURITY-DEFINITION, testes de performance alinhados com DEVOPS-SRE (SLOs).

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

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt gerador da definição de estratégia de testes | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
