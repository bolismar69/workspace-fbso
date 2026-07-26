# PROMPT-GATE-ARCHITECTURE-SCOPE

## Contexto

Este prompt implementa o **Gate de Alinhamento de Escopo** para o artefato `TECHNICAL-SOLUTION-ARCHITECTURE.md`, conforme definido no fluxo Spec-Driven Development (etapa 2.1 — GE3).

O agente validador verifica se a arquitetura proposta está **aderente ao escopo definido no TECHNICAL-SOLUTION-PRD.md** e **não introduz complexidade ou componentes além do necessário** para atender aos requisitos. O foco aqui é scope creep técnico — arquitetura que vai além do que o TECHNICAL-SOLUTION-PRD.md demanda.

**Princípio fundamental:** O TECHNICAL-SOLUTION-ARCHITECTURE.md define o COMO. Se o COMO for mais complexo ou abrangente do que o OQUE (definido no TECHNICAL-SOLUTION-PRD.md), há scope creep técnico.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 5 parâmetros foram informados.

### Passo 1 — Carregar Documentos Base

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md (artefato a ser validado)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-PRD.md (baseline de escopo — DEVE existir e estar APROVADO)
    └── Documentos de referência do projeto em {PROJECT_PATH}:
          ├── TECHNICAL-PLAN.md (stack, ERD, ADRs globais, sequenciamento)
          ├── TECHNICAL-SOLUTION-ARCHITECTURE.md (visão arquitetural do projeto — C4, ADRs)
          └── INTEGRATION-MAP.md (integrações, fluxos de dados)

Se TECHNICAL-SOLUTION-ARCHITECTURE.md não existir → ERRO: "TECHNICAL-SOLUTION-ARCHITECTURE.md não encontrado. Execute o agente gerador de arquitetura."
Se TECHNICAL-SOLUTION-PRD.md não existir → ERRO: "TECHNICAL-SOLUTION-PRD.md não encontrado. O gate de arquitetura depende do TECHNICAL-SOLUTION-PRD.md aprovado."
```

### Passo 2 — Executar Dimensões de Validação de Escopo

O gate avalia o TECHNICAL-SOLUTION-ARCHITECTURE.md em **5 dimensões**. Para cada dimensão, atribuir um veredito: `APROVADO`, `RESSALVA` ou `REPROVADO`.

#### Dimensão 1: Aderência ao TECHNICAL-SOLUTION-PRD.md (Scope Baseline Técnico)

| # | Verificação | Critério |
|---|---|---|
| 1.1 | Entidades mapeadas | As entidades descritas no TECHNICAL-SOLUTION-ARCHITECTURE.md correspondem exatamente às entidades listadas no TECHNICAL-SOLUTION-PRD.md? Há entidades extras não solicitadas? |
| 1.2 | Funcionalidades cobertas | Cada feature do TECHNICAL-SOLUTION-PRD.md tem um componente/serviço correspondente na arquitetura? Há componentes para features inexistentes? |
| 1.3 | NFRs atendidos | Os requisitos não-funcionais do TECHNICAL-SOLUTION-PRD.md estão refletidos nas decisões arquiteturais? (ex: multi-tenancy, RBAC, auditoria) |
| 1.4 | Escopo declarado | A arquitetura declara explicitamente o que NÃO cobre? Componentes "para uso futuro" estão claramente sinalizados? |

#### Dimensão 2: Simplicidade e KISS (Ponytail Check)

| # | Verificação | Critério |
|---|---|---|
| 2.1 | Package-by-Layer vs Clean Architecture | Se o TECHNICAL-SOLUTION-PRD.md não justifica Clean/Hexagonal, a arquitetura usa package-by-layer? |
| 2.2 | Número de camadas | A arquitetura tem mais camadas do que o estritamente necessário? (suspeito: > 4 camadas) |
| 2.3 | Número de serviços externos | A arquitetura introduz serviços/infraestrutura não exigidos pelo TECHNICAL-SOLUTION-PRD.md? (ex: Kafka sem requisito de streaming) |
| 2.4 | Complexidade de padrões | Há padrões de design que adicionam complexidade sem justificativa no TECHNICAL-SOLUTION-PRD.md? (ex: CQRS, Event Sourcing sem necessidade) |

#### Dimensão 3: Consistência com Technical Plan e ADRs Globais

| # | Verificação | Critério |
|---|---|---|
| 3.1 | Stack tecnológica | A stack do TECHNICAL-SOLUTION-ARCHITECTURE.md é consistente com o TECHNICAL-PLAN.md do projeto? |
| 3.2 | ADRs globais | As decisões arquiteturais respeitam os ADRs globais definidos no projeto? |
| 3.3 | Integrações | As integrações descritas batem com o INTEGRATION-MAP.md? Há integrações não documentadas? |
| 3.4 | Padrões de nomenclatura | A nomenclatura de pacotes, classes e serviços segue o padrão do projeto? |

#### Dimensão 4: Completude Técnica

| # | Verificação | Critério |
|---|---|---|
| 4.1 | Estrutura de pacotes | A árvore de diretórios cobre todos os componentes necessários para as features do TECHNICAL-SOLUTION-PRD.md? |
| 4.2 | Pipeline de segurança | O fluxo de segurança (JWT → Tenant → RBAC → Controller → Service → Repository → Audit) está documentado? |
| 4.3 | Tratamento de erros | A hierarquia de exceções e o GlobalExceptionHandler estão definidos? |
| 4.4 | Estratégia de testes | A pirâmide de testes está documentada com exemplos? |
| 4.5 | ADRs locais | Existem pelo menos 5 ADRs documentando decisões técnicas da solução? |

#### Dimensão 5: Cross-Cutting Concerns e AOP

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Tenant isolation | O mecanismo de isolamento multi-tenant está definido e é adequado ao TECHNICAL-SOLUTION-PRD.md? |
| 5.2 | RBAC | O controle de acesso baseado em papéis está modelado conforme requisitos do TECHNICAL-SOLUTION-PRD.md? |
| 5.3 | Auditoria | O mecanismo de auditoria cobre as operações exigidas pelo TECHNICAL-SOLUTION-PRD.md? |
| 5.4 | Cache e performance | Há estratégias de cache consistentes com os NFRs de performance? Não há over-engineering? |

### Passo 3 — Calcular Veredito Final

```
Para cada dimensão:
    - 100% verificações OK → APROVADO
    - >= 75% verificações OK → RESSALVA
    - < 75% verificações OK → REPROVADO

Veredito final do gate:
    ├── APROVADO: Todas as 5 dimensões APROVADAS
    ├── RESSALVA: Pelo menos 1 dimensão com RESSALVA, nenhuma REPROVADA
    └── REPROVADO: Pelo menos 1 dimensão REPROVADA
```

### Passo 4 — Gerar Relatório de Falha (se REPROVADO ou RESSALVA)

```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE_SCOPE_FAIL_REPORT.md
```

#### Estrutura do Relatório de Falha

```markdown
# ARCHITECTURE_SCOPE_FAIL_REPORT.md — Relatório de Validação de Escopo: TECHNICAL-SOLUTION-ARCHITECTURE.md

[Header com metadados: solução, projeto, stack, data da validação, versão do TECHNICAL-SOLUTION-ARCHITECTURE.md validado]

## 1. Sumário Executivo
- Veredito: REPROVADO | RESSALVA
- Dimensões aprovadas: X/5
- Dimensões com ressalva: Y/5
- Dimensões reprovadas: Z/5
- Total de não-conformidades: N

## 2. Veredito por Dimensão
| Dimensão | Veredito | Verificações OK | Total Verificações | % |
|---|---|---|---|---|
| 1. Aderência ao TECHNICAL-SOLUTION-PRD.md | ... | ... | ... | ... |
| 2. Simplicidade e KISS | ... | ... | ... | ... |
| 3. Consistência com Technical Plan | ... | ... | ... | ... |
| 4. Completude Técnica | ... | ... | ... | ... |
| 5. Cross-Cutting Concerns | ... | ... | ... | ... |

## 3. Não-Conformidades Detalhadas

### 3.1 [Dimensão 1] Aderência ao TECHNICAL-SOLUTION-PRD.md
| ID | Verificação | Status | Evidência | Ação Corretiva |
|---|---|---|---|---|
| NC-001 | 1.1 Entidades mapeadas | REPROVADO | Entidade "NotificationLog" na arquitetura sem correspondência no TECHNICAL-SOLUTION-PRD.md | Remover ou justificar no TECHNICAL-SOLUTION-PRD.md |

### 3.2 [Dimensão 2] Simplicidade e KISS
...

### 3.3 [Dimensão 3] Consistência com Technical Plan
...

### 3.4 [Dimensão 4] Completude Técnica
...

### 3.5 [Dimensão 5] Cross-Cutting Concerns
...

## 4. Itens Suspeitos de Scope Creep Técnico
- Componentes/serviços na arquitetura sem feature correspondente no TECHNICAL-SOLUTION-PRD.md
- Padrões arquiteturais excessivamente complexos para o escopo
- Infraestrutura não justificada (Kafka, Redis, Elasticsearch — se não exigidos)
- "Engenharia de ouro" (gold plating técnico)

## 5. Recomendações para Correção
- Lista priorizada de ações corretivas
- Referências aos documentos que devem ser consultados
- Seções do TECHNICAL-SOLUTION-ARCHITECTURE.md que precisam ser simplificadas ou removidas

## 6. Próximos Passos
1. Encaminhar este relatório ao Agente Arquiteto para correção
2. Após correção, reexecutar este gate (e o gate de TECHNICAL-SOLUTION-PRD.md se houve alteração de escopo)
3. Não prosseguir para TECHNICAL-SOLUTION-SPECS.md até TECHNICAL-SOLUTION-ARCHITECTURE.md ser APROVADO

## Rodapé
- Indicação de geração por IA, skills utilizados
```

### Passo 5 — Validar o Relatório Gerado

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo no path correto | `ARCHITECTURE_SCOPE_FAIL_REPORT.md` existe |
| 2 | Sumário executivo presente | §1 contém veredito e contagem |
| 3 | Tabela de veredito por dimensão | §2 lista as 5 dimensões |
| 4 | Não-conformidades detalhadas | §3 lista cada NC com ação corretiva |
| 5 | Scope creep técnico identificado | §4 lista componentes/padrões suspeitos |
| 6 | Recomendações acionáveis | §5 contém ações priorizadas |
| 7 | Próximos passos claros | §6 indica fluxo de correção |
| 8 | Rodapé de IA | Indicação de geração automatizada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `ponytail` | Revisor de simplicidade — identificar complexidade desnecessária |
| 2ª | `gap-analysis` | Análise de lacunas entre TECHNICAL-SOLUTION-ARCHITECTURE.md e TECHNICAL-SOLUTION-PRD.md |
| 3ª | `architecture-patterns` | Avaliar adequação dos padrões arquiteturais ao escopo |
| 4ª | `documentation-writer` | Qualidade do relatório de falha |

---

## Observações

1. **O TECHNICAL-SOLUTION-PRD.md é a constituição do escopo.** Toda verificação de escopo do TECHNICAL-SOLUTION-ARCHITECTURE.md é feita contra o TECHNICAL-SOLUTION-PRD.md aprovado. Se o TECHNICAL-SOLUTION-PRD.md não passou pelo gate anterior, este gate não deve ser executado.

2. **Simplicidade é requisito, não preferência.** A dimensão 2 (KISS) não é opcional. O skill `ponytail` é invocado como revisor mandatório. Complexidade sem justificativa documentada no TECHNICAL-SOLUTION-PRD.md é tratada como não-conformidade.

3. **Scope creep técnico é tão perigoso quanto scope creep funcional.** Adicionar Kafka, Redis, Elasticsearch ou padrões como CQRS sem justificativa de negócio é scope creep e será reportado como tal.

4. **O relatório de falha referencia o TECHNICAL-SOLUTION-PRD.md.** As ações corretivas devem citar exatamente qual seção do TECHNICAL-SOLUTION-PRD.md justifica (ou não) cada decisão arquitetural.

5. **Revalidação em cascata.** Se a correção do TECHNICAL-SOLUTION-ARCHITECTURE.md alterar o escopo, o gate de TECHNICAL-SOLUTION-PRD.md deve ser reexecutado para garantir consistência bidirecional.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: gate de escopo para TECHNICAL-SOLUTION-ARCHITECTURE.md, 5 dimensões, 20 verificações, foco em scope creep técnico e simplicidade | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, ponytail, gap-analysis.*
