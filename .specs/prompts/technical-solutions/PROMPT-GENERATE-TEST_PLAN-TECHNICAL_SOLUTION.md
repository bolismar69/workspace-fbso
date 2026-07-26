# PROMPT-GENERATE-TEST_PLAN-ARTEFACT

## Contexto

Este prompt orquestra skills especializadas em qualidade de software e estratégia de testes para gerar ou revisar o artefato `TECHNICAL-SOLUTION-TEST_PLAN.md` na pasta de especificações de uma solução técnica.

O artefato gerado deve ser o **plano de testes completo** da solução — cobrindo estratégia, níveis (unitário, integração, E2E, segurança), cenários por feature e critérios de cobertura.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica (para ferramentas de teste) | `Java 25 + Spring Boot + PostgreSQL` |
| `{SCOPE}` | Escopo da geração | `full`, `delta`, `audit` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 6 parâmetros foram informados.

### Passo 1 — Verificar e Preparar a Estrutura

```
Verificar: {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    │
    ├── NÃO existe → Criar
    │
    └── SIM →
            ├── Ler TECHNICAL-SOLUTION-PRD.md + TECHNICAL-SOLUTION-SPECS.md + TECHNICAL-SOLUTION-ARCHITECTURE.md (se existirem)
            ├── Fallback: docs do projeto (FEATURES, USER-STORIES, DEFINITION_OF_DONE)
            │
            └── TECHNICAL-SOLUTION-TEST_PLAN.md existe?
                  ├── SIM + SCOPE=full → Regenerar
                  ├── SIM + SCOPE=delta → Atualizar com novos cenários
                  ├── SIM + SCOPE=audit → Auditar cobertura vs. features (relatório)
                  └── NÃO → Criar do zero
```

### Passo 2 — Invocar Skills Especializadas

| Ordem | Skill | Responsabilidade |
|---|---|---|
| 1ª | `test-strategy-design` | Estratégia de testes: pirâmide, níveis, ferramentas, cobertura |
| 2ª | `qa-test-planner` | Planejamento de cenários de teste por feature |
| 3ª | `acceptance-criteria` | Vincular cenários aos critérios de aceitação das user stories |
| 4ª | `security-reviewer` | Cenários de teste de segurança (RBAC, Multi-Tenant, OWASP) |

### Passo 3 — Gerar TECHNICAL-SOLUTION-TEST_PLAN.md

Gerar em: `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-TEST_PLAN.md`

#### Estrutura Obrigatória

```markdown
# TECHNICAL-SOLUTION-TEST_PLAN.md — Plano de Testes: {SOLUTION_NAME}

## 1. Estratégia de Testes
- Pirâmide (Unit/Int/E2E/Security)
- Ferramentas por nível
- Metas de cobertura

## 2. Cenários de Teste por Feature
### FXX-XX: {nome}
| ID | Cenário | Nível | Pré-condição | Resultado Esperado | Status |

## 3. Testes de Segurança
- RBAC: cada papel × endpoint proibido
- Multi-Tenant: isolamento entre tenants
- OWASP Top 10 aplicável

## 4. Testes de Performance (NFRs)
- Cenários de carga (dashboard, listas paginadas)

## 5. Testes de Regressão
- Checklist de features já homologadas

## 6. Registro de Alterações
```

### Passo 4 — Validação Pós-Geração (10 verificações)

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `test-strategy-design` | Estratégia, pirâmide, ferramentas, cobertura |
| 2ª | `qa-test-planner` | Cenários de teste por feature |
| 3ª | `acceptance-criteria` | Vinculação cenários ↔ critérios de aceite |
| 4ª | `security-reviewer` | Cenários de segurança (RBAC, Multi-Tenant, OWASP) |

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices.*
