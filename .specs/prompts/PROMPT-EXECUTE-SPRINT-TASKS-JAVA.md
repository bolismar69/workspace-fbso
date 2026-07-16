# PROMPT-EXECUTE-SPRINT-TASKS

## Contexto

Este prompt orquestra a **execução das tarefas de uma sprint** no contexto do fluxo Spec-Driven Development. Ele consome os artefatos da sprint (`SPRINT-CARD.md`, `SPRINT-TEST-SUITE.md`, `SPRINT-REVIEW.md`) previamente aprovados pelo gate e executa cada tarefa sequencialmente, gerando código, testes e um relatório de execução ao final.

**Princípio fundamental:** A sprint é a unidade atômica de execução. As tarefas são executadas na ordem definida pelo `SPRINT-CARD.md`, respeitando as dependências documentadas. Cada tarefa concluída gera evidência de teste. O relatório final é salvo na pasta da sprint.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.
>
> **Nota:** `BRANCH_NAME` não é mais um parâmetro de entrada. Ele é lido do campo `**Branch:**` no header do `SPRINT-CARD.md` durante a Fase 0 — passo 0. Se o campo não existir no artefato, a execução é abortada.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço) | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SPRINT_NUMBER}` | Número da sprint (1 a N) | `1` |
| `{SPRINT_NAME}` | Nome curto da sprint (kebab-case) | `sprint-01-setup` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |
| `{TASK_IDS}` | IDs das tarefas a executar (opcional — se omitido, executa TODAS as tarefas pendentes da sprint) | `T-001,T-002,T-003` ou vazio para todas |

---

## Documentos de Referência

### Artefatos da Sprint (obrigatórios — fonte direta da execução)

```
SPRINT_DIR = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/sprints/{SPRINT_NAME}/

Ler obrigatoriamente antes de iniciar:
    ├── {SPRINT_DIR}/SPRINT-CARD.md        ← Backlog, estimativas, critérios DONE, riscos
    ├── {SPRINT_DIR}/SPRINT-TEST-SUITE.md  ← Cenários de teste aplicáveis a esta sprint
    └── {SPRINT_DIR}/SPRINT-REVIEW.md      ← O que demonstrar na review (foco do desenvolvimento)
```

### Documentos-Mestre (obrigatórios — baseline de verdade)

```
SPECS_DIR = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/

Ler obrigatoriamente:
    ├── {SPECS_DIR}/SPECS.md               ← Regras de negócio, endpoints, validações
    ├── {SPECS_DIR}/TASKS.md               ← Fonte da verdade das tarefas (rastreabilidade)
    ├── {SPECS_DIR}/TEST_PLAN.md           ← Fonte da verdade dos cenários de teste
    └── {SPECS_DIR}/ARCHITECTURE.md        ← Estrutura de pacotes, ADRs, padrões
    └── {SPECS_DIR}/PRD.md                 ← (se existir)Documento de referencia da origem da demanda (projeto, feature, hotfix,...)
```

### Documentos Transversais (obrigatórios)

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/security/SECURITY.md     ← Regras de segurança (se existir)
    └── {SOLUTION_PATH}/README.md                       ← Build, execução, testes
```

---

## Missão

Implemente as tarefas da **Sprint {SPRINT_NUMBER} — {SPRINT_NAME}** seguindo o backlog definido no `SPRINT-CARD.md`, utilizando o `SPRINT-TEST-SUITE.md` como checklist de qualidade, e gerando ao final o relatório de execução na pasta da sprint.

---

## Fluxo de Execução

### Fase 0 — Pré-implementação

0. **Validar branch de desenvolvimento** — a branch NÃO é mais um parâmetro de entrada. Ela deve ser lida do `SPRINT-CARD.md`:

   ```
   SPRINT_CARD = {SPRINT_DIR}/SPRINT-CARD.md

   a. LER o SPRINT-CARD.md e extrair o campo **Branch:** do header

   b. Se **Branch:** NÃO existir no SPRINT-CARD.md → ERRO CRÍTICO. PARE IMEDIATAMENTE.
      "SPRINT-CARD.md não contém **Branch:**. O artefato foi gerado sem a informação
       da branch de desenvolvimento. Execute novamente PROMPT-GENERATE-SPRINT-ARTEFACTS
       informando BRANCH_NAME. Não é seguro prosseguir sem esta informação."

   c. Se **Branch:** = 'main' ou 'master' → ERRO CRÍTICO. PARE IMEDIATAMENTE.
      "Branch '{branch}' não permitida. Desenvolvimento direto em main/master
       viola políticas de GitOps/GitSecOps. Crie uma branch feature/ ou hotfix/."

   d. Executar `git checkout {BRANCH_NAME}`. Se a branch não existir:
      `git checkout -b {BRANCH_NAME}`

   e. Confirmar: `git branch --show-current` deve retornar {BRANCH_NAME}
   ```
1. **Carregar artefatos da sprint** — ler SPRINT-CARD.md, SPRINT-TEST-SUITE.md, SPRINT-REVIEW.md
2. **Carregar documentos-mestre** — ler SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md, PRD.md (se existir)
3. **Carregar documentos transversais** — SECURITY.md, README.md
4. **Identificar tasks a executar:**
   - Se `{TASK_IDS}` foi informado → executar apenas as tasks listadas
   - Se `{TASK_IDS}` está vazio → executar TODAS as tasks do SPRINT-CARD.md com status pendente
5. **Validar ordem de execução** — verificar dependências no SPRINT-CARD.md e no TASKS.md §3. Tasks com pré-requisitos devem ser executadas depois que seus pré-requisitos forem concluídos
6. **Para cada task, aplicar o checklist ponytail (7 rungs) ANTES de codificar:**

| Rung | Pergunta | Ação |
|:---|:---|:---|
| 1 | Isso precisa existir? (YAGNI — a task é realmente necessária?) | Se NÃO → justificar e pular |
| 2 | Já existe no codebase? (classe, método, utilitário) | Se SIM → reusar, não reescrever |
| 3 | A biblioteca padrão do Java/Spring cobre? | Se SIM → usar stdlib/Spring Boot |
| 4 | Dependência já declarada no `pom.xml` resolve? | Se SIM → usar, nunca adicionar dependência para poucas linhas |
| 5 | O padrão do projeto (ARCHITECTURE.md) já define como fazer? | Se SIM → seguir o padrão |
| 6 | Dá pra ser uma classe/método simples? | Se SIM → manter simples |
| 7 | Só então: código mínimo que funciona | Escrever o mínimo |

### Fase 1 — Implementação (por task)

7. **Para cada task do backlog, executar na ordem:**

   ```
   PARA cada task T-XXX:
       │
       ├── 1. LER o critério DONE no SPRINT-CARD.md
       ├── 2. IMPLEMENTAR seguindo ARCHITECTURE.md (estrutura de pacotes, padrões)
       ├── 3. APLICAR constraints Java/Spring Boot:
       │      - Spring Boot idioms (@Service, @Repository, @Controller, @Aspect)
       │      - Bean Validation (Jakarta) para DTOs
       │      - JDBC Template (não JPA — ADR-L01)
       │      - @Transactional onde necessário
       │      - RFC 7807 para erros (GlobalExceptionHandler)
       │      - Soft Delete + Tenant Filter (BaseRepository)
       │      - @RequiresPermission + @Auditable nos endpoints aplicáveis
       ├── 4. RODAR `mvn compile` — corrigir erros de compilação
       ├── 5. RODAR `mvn test` — verificar se testes existentes não quebraram
       └── 6. MARCAR task como concluída (atualizar mentalmente — o SPRINT-CARD.md é atualizado no final)
   ```

8. **Implementar testes IMEDIATAMENTE após o código:**

   ```
   PARA cada task T-XXX:
       │
       ├── 1. CONSULTAR SPRINT-TEST-SUITE.md — quais cenários se aplicam?
       ├── 2. MAPEAR para TEST_PLAN.md — confirmar IDs e níveis de teste
       ├── 3. IMPLEMENTAR:
       │      - Unit tests (JUnit 5 + Mockito) para Services
       │      - Integration tests (Testcontainers + PostgreSQL) para Repositories
       │      - Security tests para @RequiresPermission (se aplicável)
       ├── 4. EXECUTAR `mvn test -Dtest="**/unit/**"` — confirmar verde
       ├── 5. EXECUTAR `mvn test -Dtest="**/integration/**"` — confirmar verde
       └── 6. VERIFICAR cobertura: `mvn jacoco:check` — meta ≥ 80%
   ```

### Fase 2 — Validação de Qualidade

9. **Executar verificações estáticas:**
   ```
   mvn checkstyle:check    ← Zero warnings
   mvn pmd:check           ← Zero violations
   ```

10. **Executar suite completa da sprint:**
    ```
    mvn test                 ← Todos os testes (unit + integration)
    mvn jacoco:check         ← Cobertura ≥ 80%
    ```

### Fase 3 — Tratamento de Falhas

Se algum teste falhar durante a execução:

1. **Auto-Correção Autônoma (até 3 tentativas):** Analise o stack trace, identifique se o problema está na lógica do código ou na estrutura do teste, e corrija.
2. **Tratamento de Loops:** Se o mesmo erro persistir por **3 tentativas**, PARE imediatamente.
3. **Registro de Impedimento:** Crie `{SPRINT_DIR}/IMPEDIMENT-SPRINT-{SPRINT_NUMBER}.md` com:
   - Task que falhou (T-XXX) e mensagem de erro exata
   - O que foi tentado para corrigir
   - Suspeita do motivo (limitação arquitetural, ambiguidade na SPECS.md, etc.)
   - Propostas adicionais de solução (se houver)
4. **Notifique o humano** e aguarde instruções antes de alterar qualquer outro arquivo.

### Fase 4 — Pós-implementação (Sanity Check)

11. **Limpeza:** Remover código comentado, `System.out.println` de debug, imports não usados, arquivos temporários.
12. **Git Status:** Listar todos os arquivos modificados ou criados (`git status --short`).
13. **Localização:** Validar que os arquivos respeitam o `ARCHITECTURE.md` (estrutura de pacotes, nomes de classe).
14. **Segurança:** Revisão final contra `SECURITY.md` — nenhuma regra violada. Verificar:
    - Nenhuma credencial hardcoded
    - Queries usam PreparedStatement (JDBC parametrizado)
    - Endpoints anotados com `@RequiresPermission` onde aplicável
    - Dados pessoais não expostos em logs ou respostas
15. **Atualizar SPRINT-CARD.md:** Marcar cada task concluída (adicionar `[x]` ou atualizar tabela de progresso).

### Fase 5 — Geração do Relatório de Execução (OBRIGATÓRIO)

16. **Gerar arquivo de report em:**
    ```
    {SPRINT_DIR}/SPRINT-{SPRINT_NUMBER}-EXECUTION-REPORT.md
    ```

#### Estrutura do Relatório

```markdown
# SPRINT-{N}-EXECUTION-REPORT.md — Relatório de Execução: Sprint {N}

[Header: solução, projeto, sprint, stack, data da execução, tasks executadas]

## 1. Resumo da Execução
- Tasks executadas: X/Y
- Tasks com sucesso: X
- Tasks com falha: 0
- Tempo total estimado: Xd
- Tempo total gasto: (preencher)

## 2. Tasks Executadas
| ID | Tarefa | Status | Testes | Cobertura | Observações |
|:---|:---|:---:|:---:|:---:|:---|
| T-XXX | Descrição | ✅ | N/N passando | XX% | — |
| ... | ... | ... | ... | ... | ... |

## 3. Arquivos Criados ou Modificados
| Ação | Arquivo | Task | Descrição da Mudança |
|:---|:---|:---|:---|
| 🆕 | `src/main/java/.../NovoArquivo.java` | T-XXX | Nova classe — propósito |
| 🔄 | `src/main/java/.../Existente.java` | T-YYY | Modificado — razão |

## 4. Evidências de Testes
- Comando executado: `mvn test`
- Total de testes: N
- Status: ✅ 100% PASS
- Cobertura JaCoCo: XX% (linhas), XX% (branchs)
- Cenários do SPRINT-TEST-SUITE.md executados: N/N

## 5. Validação de Segurança
- [ ] Nenhuma credencial ou dado sensível hardcoded
- [ ] Queries usam PreparedStatement (JDBC parametrizado)
- [ ] Endpoints anotados com @RequiresPermission (se aplicável)
- [ ] Dados pessoais mascarados em logs
- [ ] Respostas de erro seguem RFC 7807 (sem stack traces)

## 6. Validação de Arquitetura
- [ ] Estrutura de pacotes segue ARCHITECTURE.md
- [ ] Classes usam anotações Spring idiomáticas (@Service, @Repository, @Controller)
- [ ] BaseRepository usado como template para novos repositories
- [ ] Soft Delete + Tenant Filter respeitados

## 7. Desvios e Observações
- Desvios das specs originais (se houver) e justificativa
- Decisões de design tomadas durante a implementação
- Comentários `// ponytail:` adicionados (atalhos intencionais documentados)

## 8. Próximos Passos
- Tasks restantes na sprint (se houver)
- Pré-requisitos para a próxima sprint
- Recomendações para a review com o PO

## Rodapé
[Indicação de geração por IA, skills utilizados, data/hora]
```

---

## Skills Ativas

| Skill | Modo | Função |
|:---|:---|:---|
| `110-java-maven-best-practices` | automático | Estrutura Maven, `pom.xml`, plugins, convenções |
| `301-frameworks-spring-boot-core` | automático | Configuração Spring Boot, `application.yml`, profiles |
| `302-frameworks-spring-boot-rest` | automático | Controllers REST, DTOs, Bean Validation, RFC 7807 |
| `304-frameworks-spring-boot-security` | automático | JWT Filter, RBAC, TenantContext, aspectos AOP |
| `311-frameworks-spring-jdbc` | automático | JDBC Template, queries parametrizadas, soft delete |
| `130-java-testing-strategies` | automático | Estratégia de testes (unit, integration, security) |
| `131-java-testing-unit-testing` | automático | JUnit 5 + Mockito para services |
| `132-java-testing-integration-testing` | automático | Testcontainers + PostgreSQL para repositories |
| `ponytail` | full | Escada YAGNI de 7 rungs — controle de escopo |
| `caveman` | full | Compressão de prosa interativa (comunicação, não artefatos) |

> ⚠️ **caveman e ponytail NÃO atuam sobre artefatos permanentes** (SPRINT-CARD.md, SPRINT-REVIEW.md, EXECUTION-REPORT.md). O caveman comprime apenas a comunicação interativa; o report final é gerado em prosa normal.

---

## Protocolo de Testes

Durante a execução, para cada task implementada:

```
1. mvn compile              ← Compilar (falha = corrigir antes de testar)
2. mvn test                 ← Suite completa (unit + integration)
3. mvn jacoco:check         ← Meta ≥ 80% linhas
```

Se `mvn test` falhar:
1. Analisar stack trace e corrigir (até 3 tentativas)
2. Se persistir → criar `{SPRINT_DIR}/IMPEDIMENT-SPRINT-{SPRINT_NUMBER}.md`
3. Notificar humano e aguardar

---

## Output Esperado

| Output | Descrição |
|:---|:---|
| Código-fonte | Classes Java implementadas seguindo ARCHITECTURE.md |
| Testes | Unit (JUnit 5 + Mockito) + Integration (Testcontainers) |
| `mvn compile` | BUILD SUCCESS |
| `mvn test` | Todos verdes |
| `mvn jacoco:check` | ≥ 80% coverage |
| `mvn checkstyle:check` | Zero warnings |
| **ARTEFATO** | `{SPRINT_DIR}/SPRINT-{N}-EXECUTION-REPORT.md` |

---

## Anti-Padrões

| ❌ NÃO fazer | ✅ Fazer |
|:---|:---|
| Implementar tasks sem ler SPRINT-CARD.md | Sempre começar pelo SPRINT-CARD.md |
| Pular testes e deixar "para depois" | Escrever testes IMEDIATAMENTE após o código |
| Implementar tasks de sprints futuras | Uma sprint de cada vez, na ordem dos marcos |
| Ignorar o SPRINT-TEST-SUITE.md | Usá-lo como checklist de qualidade |
| Gerar report fora da pasta da sprint | Salvar em `{SPRINT_DIR}/SPRINT-{N}-EXECUTION-REPORT.md` |
| "Vou adicionar X porque vai ser útil depois" | Seguir YAGNI — implementar apenas o que a task pede |
| Usar JPA/Hibernate "porque é mais fácil" | Usar JDBC Template (ADR-L01) — controle total sobre SQL |
| Hardcodar tenant_id ou pular TenantIsolation | Todo repository herda de BaseRepository com Tenant Filter |
| Deixar stack trace vazar em respostas HTTP | GlobalExceptionHandler captura tudo — RFC 7807 |

---

## Observações

1. **O SPRINT-CARD.md é o roteiro.** Cada task tem critério DONE explícito — não considere uma task concluída até que todos os critérios DONE sejam atendidos.

2. **O SPRINT-TEST-SUITE.md é o checklist de qualidade.** Se um cenário de teste falha, a task associada não está concluída, mesmo que o código compile.

3. **Respeitar a ordem das tasks.** O SPRINT-CARD.md lista as tasks na ordem recomendada de execução. Tasks com dependências (ex: T-004 depende de T-001) devem ser executadas em sequência.

4. **Artefatos de sprint são derivados.** Se durante a execução você encontrar uma inconsistência entre os artefatos de sprint e os documentos-mestre, os documentos-mestre prevalecem. Documente a inconsistência no relatório de execução.

5. **Relatório de execução é obrigatório.** Não é opcional. É a evidência de que a sprint foi executada e o ponto de partida para a review com o PO.

6. **Sprints de fundação (1 e 2) são diferentes.** Não espere testes de feature ou endpoints REST — os testes são estruturais (build, migration, segurança).

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 14/07/2026 | Criação inicial: execução de tasks de sprint em Java/Spring Boot, adaptado de PROMPT-EXECUTE-TASK.md. Reports salvos na pasta da sprint. Integração com SPRINT-CARD, SPRINT-TEST-SUITE e SPRINT-REVIEW | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, 110-java-maven-best-practices, 301-frameworks-spring-boot-core, ponytail.*
