# PROMPT-GENERATE-ARCHITECTURE-ARTEFACT

## Contexto

Este prompt orquestra skills especializadas em arquitetura de solução e engenharia de software para gerar ou revisar o artefato `TECHNICAL-SOLUTION-ARCHITECTURE.md` na pasta de especificações de uma solução técnica.

O artefato gerado deve seguir o padrão **package-by-layer** (simplicidade > pureza arquitetural) com aspectos AOP para cross-cutting concerns, conforme preconizado pelo skill `ponytail` como referência de simplificação de soluções.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço) | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Antes de qualquer ação, verificar se TODOS os 5 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

### Passo 1 — Verificar e Preparar a Estrutura de Pastas

```
Verificar se existe: {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    │
    ├── NÃO existe → Criar a pasta (mkdir -p)
    │
    └── SIM, existe →
            │
            ├── Verificar se existe TECHNICAL-SOLUTION-PRD.md na pasta
            │     │
            │     ├── SIM → Ler TECHNICAL-SOLUTION-PRD.md como ponto de partida principal
            │     │         (contém escopo, entidades, ADRs, estrutura de pacotes esperada)
            │     │
            │     └── NÃO → Ler documentos do projeto em {PROJECT_PATH}:
            │           ├── 01-PROJECT-CHARTER-*.md (escopo, entregas, marcos)
            │           ├── 02-BUSINESS-REQUIREMENTS.md (requisitos funcionais e NFRs)
            │           ├── 03-EPICS.md (épicos e jornadas de usuário)
            │           ├── 04-FEATURES.md (features, user stories, regras de negócio)
            │           ├── TECHNICAL-PLAN.md (stack, ERD, ADRs, sequenciamento)
            │           ├── TECHNICAL-SOLUTION-ARCHITECTURE.md (visão arquitetural do projeto — C4, ADRs)
            │           ├── INTEGRATION-MAP.md (integrações, fluxos de dados)
            │           └── API-CONTRACTS.md (contratos de API)
            │
            └── Ler também (sempre, independente de TECHNICAL-SOLUTION-PRD.md):
                  ├── {PROJECT_PATH}/TECHNICAL-PLAN.md (stack, ERD, ADRs globais)
                  └── {PROJECT_PATH}/TECHNICAL-SOLUTION-ARCHITECTURE.md (decisões arquiteturais do projeto)
```

### Passo 2 — Invocar Skills Especializadas

Invocar as skills na ordem abaixo para embasar a geração do artefato:

| Ordem | Skill | Responsabilidade | O que extrair |
|---|---|---|---|
| 1ª | `architecture-patterns` | Clean Architecture, Hexagonal, DDD — patterns para estruturar o backend | Estrutura de camadas, regras de dependência, design de repository, testes |
| 2ª | `engineering-skills` | Práticas de engenharia de software (estrutura de pacotes, tratamento de erros, logging) | Padrão de pacotes, hierarchy de exceções, estratégia de testes |
| 3ª | `030-architecture-adr-general` | Architecture Decision Records — documentar decisões técnicas | ADRs locais da solução com justificativa e trade-offs |
| 4ª | `ponytail` | Simplificação de soluções — eliminar complexidade desnecessária | Revisão final: a estrutura está simples o suficiente? Há algo que pode ser removido? |

> **Regra:** O skill `ponytail` deve ser invocado como **revisor final** do artefato gerado. Se ele sugerir simplificações, aplicá-las antes de salvar o arquivo.

### Passo 3 — Gerar ou revisar o Artefato TECHNICAL-SOLUTION-ARCHITECTURE.md

Gerar ou revisar o arquivo em:
```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md
```

#### Estrutura Obrigatória do Arquivo

```markdown
# TECHNICAL-SOLUTION-ARCHITECTURE.md — Arquitetura da Solução: {SOLUTION_NAME}

[Header com metadados: solução, stack, projeto, versão, data, status, origem]

## 1. Estilo Arquitetural
- Package-by-Layer com Spring Boot (explicar por que, comparar com Clean/Hexagonal)
- Diagrama ASCII da estrutura de camadas
- ADR documentando a decisão de simplificação

## 2. Estrutura de Pacotes
- Árvore completa de diretórios com ~30-40 classes mapeadas
- descrição de 1 linha por pacote

## 3. Pipeline de Segurança por Requisição
- Diagrama do fluxo: JWT Filter → TenantContext → RBAC → Controller → Service → Repository → Audit
- 7 etapas numeradas com descrição do que cada uma faz

## 4. Design dos Aspectos Cross-Cutting
- @RequiresPermission (anotação + uso no controller)
- @Auditable (anotação + uso no service)
- TenantIsolationAspect (automático — intercepta @Repository)

## 5. Design de Persistência
- BaseRepository com Soft Delete + Tenant Filter
- Índices únicos parciais (SQL de exemplo das migrations)

## 6. Tratamento de Erros (RFC 7807)
- Hierarchy de exceções
- GlobalExceptionHandler com exemplos de código

## 7. Estratégia de Testes
- Pirâmide (Unit/Int/E2E)
- Exemplo de teste de service com mock
- Exemplo de teste de isolamento Multi-Tenant com Testcontainers

## 8. Decisões de Design (ADRs Locais)
- Tabela com 5-7 ADRs: ID, Decisão, Justificativa
- Pelo menos 1 ADR documentando por que package-by-layer e não Clean Architecture

## 9. Registro de Alterações
- Tabela de versões com data, alteração, autor

## Rodapé
- Indicação de geração por IA, skills utilizados
```

### Passo 4 — Validação Pós-Geração

Após gerar ou revisar o arquivo, executar as seguintes verificações:

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo criado no path correto | `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md` existe |
| 2 | Estrutura de pacotes mapeada | Árvore de diretórios com pelo menos 8 pacotes e 30 classes |
| 3 | Pipeline de segurança documentado | 7 etapas do fluxo de segurança numeradas |
| 4 | Aspectos cross-cutting detalhados | @RequiresPermission, @Auditable, TenantIsolationAspect com exemplos de código |
| 5 | BaseRepository com Soft Delete | Template method documentado |
| 6 | Pelo menos 5 ADRs locais | Tabela com ID, Decisão, Justificativa |
| 7 | ADR de simplificação presente | Um ADR documentando package-by-layer vs Clean Architecture |
| 8 | Estratégia de testes | Pirâmide + 2 exemplos de código |
| 9 | Rodapé de IA | Indicação de geração automatizada + skills utilizados |
| 10 | Revisão do ponytail aplicada | Se o ponytail sugeriu simplificações, elas foram incorporadas |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `architecture-patterns` | Padrões arquiteturais (Clean, Hexagonal, DDD) — base teórica |
| 2ª | `engineering-skills` | Práticas de engenharia — estrutura de pacotes, erros, testes |
| 3ª | `030-architecture-adr-general` | ADRs — documentar decisões com justificativa |
| 4ª | `ponytail` | Revisor de simplicidade — eliminar complexidade desnecessária |

---

## Exemplo de Invocação

```
Humano: "Gerar TECHNICAL-SOLUTION-ARCHITECTURE.md para o microsserviço ms-fbso-platform-admin"

Agente: "Vou precisar de 5 parâmetros:
  - SOLUTION_PATH: caminho da pasta do microsserviço
  - PROJECT_PATH: caminho da pasta do projeto de negócio
  - PROJECT_NAME: código do projeto
  - SOLUTION_NAME: nome do microsserviço
  - STACK: stack tecnológica"

Humano: "SOLUTION_PATH=/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin
         PROJECT_PATH=/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG
         PROJECT_NAME=PRJ-FIN-2026-0003-SAAS-FBSO-ORG
         SOLUTION_NAME=ms-fbso-platform-admin
         STACK=Java 25 + Spring Boot + PostgreSQL"

Agente: [Executa Passo 1 → Passo 2 → Passo 3 → Passo 4]
```

---

## Observações

1. **TECHNICAL-SOLUTION-PRD.md é o ponto de partida preferencial.** Se existir, ele contém o escopo exato do que o microsserviço implementa, as entidades, ADRs e a estrutura de pacotes esperada. Os documentos do projeto são complementares.

2. **O skill `ponytail` é obrigatório como revisor final.** Ele garante que a arquitetura proposta é a mais simples possível para o contexto (time, prazo, escopo). Se o ponytail sugerir simplificações, elas DEVEM ser aplicadas.

3. **Package-by-layer é o padrão default.** A menos que o TECHNICAL-SOLUTION-PRD.md ou os documentos do projeto justifiquem explicitamente uma arquitetura mais complexa (Clean, Hexagonal), o padrão é package-by-layer com aspectos AOP para cross-cutting concerns.

4. **Nunca sobrescrever um TECHNICAL-SOLUTION-ARCHITECTURE.md existente sem auditoria de delta.** Se o arquivo já existir, gerar uma nova versão (incrementar versão) e documentar no Registro de Alterações o que mudou e por quê.

5. **Os paths nos documentos devem ser relativos.** Usar paths relativos (ex: `../../../../../../../business-inputs/...`) para que os links funcionem independente da máquina onde o código for clonado.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: fluxo de 4 passos, 4 skills orquestradas, 10 verificações pós-geração, referência ao ponytail como revisor de simplicidade | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices.*
