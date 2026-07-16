# PROMPT-GATE-TEST_PLAN-TECHNICAL

## Contexto

Este prompt implementa o **Gate de Viabilidade e Completude Técnica** para o artefato `TEST_PLAN.md`, conforme definido no fluxo Spec-Driven Development (etapa 3.1 — GT3).

O agente validador verifica se o plano de testes cobre **todas as features especificadas**, contempla **todos os níveis da pirâmide de testes** (unitário, integração, E2E, segurança) e define **cenários testáveis com critérios objetivos de aprovação**.

**Princípio fundamental:** Um plano de testes incompleto resulta em cobertura insuficiente, bugs em produção e regressões não detectadas. Este gate é a última barreira antes da execução do desenvolvimento.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica (para validação de ferramentas de teste) | `Java 25 + Spring Boot + PostgreSQL` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 5 parâmetros foram informados.

### Passo 1 — Carregar Documentos Base

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TEST_PLAN.md (artefato a ser validado)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md (especificações — APROVADO)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TASKS.md (tarefas — APROVADO)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md (estratégia de testes — APROVADO)
    └── Documentos de referência:
          ├── {PROJECT_PATH}/04-FEATURES.md (features, user stories, critérios de aceitação)
          └── {PROJECT_PATH}/DEFINITION_OF_DONE.md (critérios de DONE)

Se TEST_PLAN.md não existir → ERRO: "TEST_PLAN.md não encontrado."
Se SPECS.md ou TASKS.md não estiverem APROVADOS → ERRO: "SPECS.md e TASKS.md aprovados são pré-requisitos."
```

### Passo 2 — Executar Dimensões de Validação Técnica

O gate avalia o TEST_PLAN.md em **6 dimensões**. Veredito por dimensão: `APROVADO`, `RESSALVA` ou `REPROVADO`.

#### Dimensão 1: Cobertura de Features e User Stories

| # | Verificação | Critério |
|---|---|---|
| 1.1 | Feature coverage | Cada feature do SPECS.md tem cenários de teste no TEST_PLAN.md? |
| 1.2 | User story coverage | Cada user story tem pelo menos 1 cenário de teste? |
| 1.3 | Cenários órfãos | Existem cenários de teste para features/US não existentes no SPECS.md? |
| 1.4 | Cobertura de regras de negócio | Cada RN do SPECS.md tem cenários de teste que validam casos de borda? |

#### Dimensão 2: Pirâmide de Testes

| # | Verificação | Critério |
|---|---|---|
| 2.1 | Testes unitários | Há cenários de teste unitário para serviços, repositórios e utilitários? |
| 2.2 | Testes de integração | Há cenários de teste de integração para APIs, banco de dados e mensageria? |
| 2.3 | Testes E2E | Há cenários de teste end-to-end para fluxos críticos de negócio? |
| 2.4 | Testes de segurança | Há cenários para RBAC, Multi-Tenant, OWASP Top 10? |
| 2.5 | Proporção da pirâmide | A distribuição segue a pirâmide? (muitos unitários, alguns integração, poucos E2E) |

#### Dimensão 3: Qualidade dos Cenários de Teste

| # | Verificação | Critério |
|---|---|---|
| 3.1 | Estrutura completa | Cada cenário tem: ID, descrição, nível, pré-condição, passos, resultado esperado? |
| 3.2 | Pré-condições explícitas | As pré-condições são específicas e reproduzíveis? (ex: "usuário autenticado como ADMIN no tenant X") |
| 3.3 | Resultados esperados objetivos | Os resultados esperados são verificáveis? (evitar "funciona", "está correto") |
| 3.4 | Cobertura de estados | Cada cenário cobre os estados relevantes? (loading, empty, error, success, edge cases) |

#### Dimensão 4: Cobertura de Segurança

| # | Verificação | Critério |
|---|---|---|
| 4.1 | RBAC por papel | Para cada endpoint, há cenários testando acesso com papel autorizado E negando acesso sem o papel? |
| 4.2 | Multi-Tenant | Há cenários verificando que o Tenant A não acessa dados do Tenant B? |
| 4.3 | OWASP Top 10 | SQL Injection, XSS, CSRF, Broken Access Control estão cobertos? |
| 4.4 | Dados sensíveis | Há cenários validando que dados sensíveis não vazam em respostas ou logs? |

#### Dimensão 5: Cobertura de NFRs e Performance

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Testes de carga | Há cenários de teste de carga para endpoints críticos? |
| 5.2 | Métricas de performance | Os cenários especificam métricas objetivas? (ex: "p95 < 200ms com 100 req/s") |
| 5.3 | Testes de disponibilidade | Há cenários para verificar disponibilidade e resiliência? |
| 5.4 | Testes de concorrência | Há cenários para operações concorrentes e race conditions? |

#### Dimensão 6: Testes de Regressão e Organização

| # | Verificação | Critério |
|---|---|---|
| 6.1 | Suite de regressão | Há checklist de regressão para features já homologadas? |
| 6.2 | Organização por feature | Os cenários estão agrupados por feature (FXX-XX)? |
| 6.3 | Status rastreável | Cada cenário tem status (planejado/implementado/aprovado/falhou)? |
| 6.4 | Ferramentas e frameworks | As ferramentas de teste listadas são compatíveis com a stack ({STACK})? |

### Passo 3 — Calcular Veredito Final

```
Para cada dimensão:
    - 100% verificações OK → APROVADO
    - >= 75% verificações OK → RESSALVA
    - < 75% verificações OK → REPROVADO

Veredito final:
    ├── APROVADO: Todas as 6 dimensões APROVADAS
    ├── RESSALVA: Pelo menos 1 dimensão com RESSALVA, nenhuma REPROVADA
    └── REPROVADO: Pelo menos 1 dimensão REPROVADA
```

### Passo 4 — Gerar Relatório de Falha (se REPROVADO ou RESSALVA)

```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL_TEST_PLAN_FAIL_REPORT.md
```

#### Estrutura do Relatório de Falha

```markdown
# TECHNICAL_TEST_PLAN_FAIL_REPORT.md — Relatório de Validação Técnica: TEST_PLAN.md

[Header: solução, stack, projeto, data, versão do TEST_PLAN.md validado]

## 1. Sumário Executivo
- Veredito: REPROVADO | RESSALVA
- Dimensões aprovadas: X/6
- Dimensões com ressalva: Y/6
- Dimensões reprovadas: Z/6
- Total de não-conformidades: N
- Bloqueia DESENVOLVIMENTO? SIM (se REPROVADO) | NÃO (se RESSALVA, mas requer acompanhamento)

## 2. Veredito por Dimensão
| Dimensão | Veredito | Verificações OK | Total | % |
|---|---|---|---|---|
| 1. Cobertura de Features e US | ... | ... | ... | ... |
| 2. Pirâmide de Testes | ... | ... | ... | ... |
| 3. Qualidade dos Cenários | ... | ... | ... | ... |
| 4. Cobertura de Segurança | ... | ... | ... | ... |
| 5. Cobertura de NFRs | ... | ... | ... | ... |
| 6. Regressão e Organização | ... | ... | ... | ... |

## 3. Não-Conformidades Detalhadas

### 3.1 [Dimensão 1] Cobertura de Features e User Stories
| ID | Verificação | Status | Evidência | Ação Corretiva |
|---|---|---|---|---|
| NC-001 | 1.1 Feature coverage | REPROVADO | Feature F06-03 sem cenários de teste | Adicionar cenários para F06-03 |

[... repetir para cada dimensão]

## 4. Mapa de Cobertura (Features × Cenários)
| Feature | User Stories | Cenários Unit | Cenários Int | Cenários E2E | Cenários Sec | Cobertura % |
|---|---|---|---|---|---|---|
| F01-01 | US-001, US-002 | 3 | 2 | 1 | 2 | 100% |
| F06-03 | US-045 | 0 | 0 | 0 | 0 | 0% ⚠️ |

## 5. Gaps de Segurança
- Features sem cenários de RBAC
- Features sem cenários de isolamento Multi-Tenant
- Vulnerabilidades OWASP não cobertas

## 6. Recomendações para Correção
- Lista priorizada de ações corretivas (bloqueantes primeiro)
- Features que precisam de cenários adicionais
- Cenários com estrutura incompleta que precisam ser detalhados
- Sugestões de cenários de segurança ausentes

## 7. Próximos Passos
1. Encaminhar ao Agente de Testes para correção
2. Após correção, reexecutar este gate
3. Não prosseguir para DESENVOLVIMENTO até TEST_PLAN.md ser APROVADO

## Rodapé
- Indicação de geração por IA, skills utilizados
```

### Passo 5 — Validar o Relatório Gerado

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo no path correto | `TECHNICAL_TEST_PLAN_FAIL_REPORT.md` existe |
| 2 | Sumário executivo | §1 contém veredito e indica se bloqueia desenvolvimento |
| 3 | Veredito por dimensão | §2 lista as 6 dimensões |
| 4 | Não-conformidades detalhadas | §3 lista cada NC com ação corretiva |
| 5 | Mapa de cobertura | §4 apresenta tabela Feature × Cenários com % de cobertura |
| 6 | Gaps de segurança | §5 identifica features sem cenários de segurança |
| 7 | Recomendações acionáveis | §6 contém ações corretivas priorizadas |
| 8 | Próximos passos | §7 indica fluxo de correção |
| 9 | Rodapé de IA | Indicação de geração automatizada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `test-strategy-design` | Validar estratégia de testes e pirâmide |
| 2ª | `qa-test-planner` | Auditar cobertura de cenários por feature |
| 3ª | `security-reviewer` | Verificar cobertura de cenários de segurança |
| 4ª | `gap-analysis` | Identificar lacunas de cobertura vs SPECS.md |
| 5ª | `acceptance-criteria` | Verificar alinhamento cenários ↔ critérios de aceite |
| 6ª | `documentation-writer` | Qualidade do relatório de falha |

---

## Observações

1. **Último gate antes do desenvolvimento.** O TEST_PLAN.md é o último artefato validado antes do time começar a codificar. Um plano de testes fraco resulta em bugs escapando para produção.

2. **Cobertura de segurança não é opcional.** Mesmo que o projeto não tenha um requisito explícito de segurança, cenários de RBAC, Multi-Tenant e OWASP básico são obrigatórios (§4 do relatório detalha gaps).

3. **O mapa de cobertura é a evidência visual.** A tabela Feature × Cenários no §4 do relatório permite identificar rapidamente features com cobertura zero ou insuficiente.

4. **Cenários órfãos indicam desalinhamento.** Cenários de teste para features inexistentes sugerem que o TEST_PLAN.md foi gerado a partir de um SPECS.md desatualizado.

5. **Reexecução após correção.** Assim como os outros gates, este gate deve ser reexecutado após cada ciclo de correção do TEST_PLAN.md.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: gate técnico para TEST_PLAN.md, 6 dimensões, 25 verificações, mapa de cobertura Feature × Cenários | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, test-strategy-design, security-reviewer.*
