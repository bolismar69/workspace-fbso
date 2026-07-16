# PROMPT-EXECUTE-SPRINT-TASKS

## Contexto

Este prompt orquestra a **execução das tarefas de uma sprint** no contexto do fluxo Spec-Driven Development. Ele consome os artefatos da sprint (`SPRINT-CARD.md`, `SPRINT-TEST-SUITE.md`, `SPRINT-REVIEW.md`) previamente aprovados pelo gate e executa cada tarefa sequencialmente, gerando código, testes e um relatório de execução ao final.

**Princípios fundamentais:**

1. **Stack-agnóstico:** Este prompt não pressupõe linguagem, framework ou banco de dados específicos. A stack é descoberta a partir dos documentos do projeto (`PRD.md`, `SPECS.md`, `TASKS.md`, `ARCHITECTURE.md`) e as skills apropriadas são acionadas dinamicamente.
2. **Sprint como unidade atômica:** As tarefas são executadas na ordem definida pelo `SPRINT-CARD.md`, respeitando as dependências documentadas.
3. **Localidade dos artefatos:** Todo output (relatórios, impedimentos) é salvo na pasta da sprint, não em diretórios globais.

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
| `{STACK}` | Stack tecnológica principal (opcional — se omitido, será auto-detectada dos documentos do projeto) | `Java 25 + Spring Boot + PostgreSQL` ou vazio para auto-detecção |
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
    ├── {SPECS_DIR}/PRD.md                 ← Stack, escopo, entidades, ADRs (se existir)
    ├── {SPECS_DIR}/SPECS.md               ← Regras de negócio, endpoints, validações
    ├── {SPECS_DIR}/TASKS.md               ← Fonte da verdade das tarefas (rastreabilidade)
    ├── {SPECS_DIR}/TEST_PLAN.md           ← Fonte da verdade dos cenários de teste
    └── {SPECS_DIR}/ARCHITECTURE.md        ← Estrutura de diretórios, ADRs, padrões
```

### Documentos Transversais (obrigatórios)

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/security/SECURITY.md     ← Regras de segurança (se existir)
    └── {SOLUTION_PATH}/README.md                       ← Build, execução, testes, comandos
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
2. **Carregar documentos-mestre** — ler PRD.md (se existir), SPECS.md, TASKS.md, TEST_PLAN.md, ARCHITECTURE.md
3. **Carregar documentos transversais** — SECURITY.md, README.md
4. **Identificar tasks a executar:**
   - Se `{TASK_IDS}` foi informado → executar apenas as tasks listadas
   - Se `{TASK_IDS}` está vazio → executar TODAS as tasks do SPRINT-CARD.md com status pendente
5. **Validar ordem de execução** — verificar dependências no SPRINT-CARD.md e no TASKS.md §3. Tasks com pré-requisitos devem ser executadas depois que seus pré-requisitos forem concluídos
6. **Descobrir a stack do projeto e acionar as skills corretas:**

   ```
   ├── 6.1 IDENTIFICAR a stack a partir de:
   │      1º {STACK} (parâmetro informado — se preenchido, usar diretamente)
   │      2º PRD.md (campo "Stack" no header ou §1)
   │      3º SPECS.md (campo "Stack" no header ou §1.1)
   │      4º ARCHITECTURE.md (campo "Stack" no header)
   │      5º TASKS.md (campo "Stack" no header)
   │      6º README.md (instruções de build — ex: mvn, go, npm, cargo)
   │
   ├── 6.2 DECOMPOR a stack em componentes:
   │      - Linguagem(s): Java, Go, Python, TypeScript, Rust, Kotlin, C#, PHP, Ruby
   │      - Framework(s): Spring Boot, Quarkus, Micronaut, Fiber, Gin, Express, FastAPI, Next.js, Django, Rails, Laravel
   │      - Banco(s): PostgreSQL, MySQL, MongoDB, Redis, Elasticsearch
   │      - Infraestrutura: Docker, K8s, RabbitMQ, Kafka
   │
   ├── 6.3 SELECIONAR as skills apropriadas para cada componente.
   │      Utilizar o skill `001-skills-inventory` para listar as skills disponíveis.
   │      Cruzar cada componente da stack com as skills correspondentes.
   │
   │      Tabela de referência (exemplos — NÃO exaustiva):
   │
   │      | Componente | Skills prováveis |
   │      |---|---|
   │      | Linguagem: Java | `110-java-maven-best-practices`, `121-java-object-oriented-design`, `126-java-exception-handling`, `130-java-testing-strategies`, `141-java-refactoring-with-modern-features` |
   │      | Linguagem: Go | `golang-pro`, `golang-performance`, `golang-testing`, `golang-error-handling` |
   │      | Linguagem: TypeScript | `typescript-pro`, `javascript-typescript-jest` |
   │      | Linguagem: Python | `python-pro`, `pytest-coverage` |
   │      | Linguagem: Rust | `rust-engineer` |
   │      | Framework: Spring Boot | `301-frameworks-spring-boot-core`, `302-frameworks-spring-boot-rest`, `304-frameworks-spring-boot-security`, `311-frameworks-spring-jdbc` |
   │      | Framework: Quarkus | `401-frameworks-quarkus-core`, `402-frameworks-quarkus-rest` |
   │      | Framework: Micronaut | `501-frameworks-micronaut-core`, `502-frameworks-micronaut-rest` |
   │      | Framework: Fiber/Gin (Go) | `golang-pro` (já cobre padrões idiomáticos) |
   │      | Banco: PostgreSQL | `postgres-pro`, `postgresql-optimization` |
   │      | Banco: MongoDB | `mongodb-query-optimizer`, `mongodb-schema-design` |
   │      | Testes: Qualquer stack | `131-java-testing-unit-testing` (Java), `golang-testing` (Go), `javascript-typescript-jest` (JS/TS), `pytest-coverage` (Python) |
   │      | Segurança: Qualquer stack | `security-review`, `124-java-secure-coding` (Java), `gdpr-compliant` |
   │
   │      ⚠️ Se houver dúvida sobre quais skills acionar para um componente,
   │         PERGUNTAR ao humano: "Para a stack {stack_detectada}, identificamos
   │         os componentes X, Y, Z. Skills sugeridas: A, B, C. Confirma?"
   │
   └── 6.4 REGISTRAR a decisão: quais skills foram acionadas e por quê.
          Este registro fará parte do relatório de execução (§9).
   ```

7. **Para cada task, aplicar o checklist ponytail (7 rungs) ANTES de codificar:**

| Rung | Pergunta | Ação |
|:---|:---|:---|
| 1 | Isso precisa existir? (YAGNI — a task é realmente necessária?) | Se NÃO → justificar e pular |
| 2 | Já existe no codebase? (classe, função, módulo, utilitário) | Se SIM → reusar, não reescrever |
| 3 | A biblioteca padrão da linguagem cobre? | Se SIM → usar stdlib, não importar lib externa |
| 4 | Dependência já declarada no projeto resolve? | Se SIM → usar, nunca adicionar dependência nova para poucas linhas |
| 5 | O padrão do projeto (ARCHITECTURE.md) já define como fazer? | Se SIM → seguir o padrão documentado |
| 6 | Dá pra ser uma unidade simples? (função, classe, módulo enxuto) | Se SIM → manter simples |
| 7 | Só então: código mínimo que funciona | Escrever o mínimo |

### Fase 1 — Planejamento do Desenvolvimento

8. **Gerar o artefato `SPRINT-DEVELOPMENT-PLANNING.md` ANTES de iniciar a implementação:**

   ```
   {SPRINT_DIR}/SPRINT-DEVELOPMENT-PLANNING.md
   ```

   Este artefato documenta o plano de desenvolvimento da sprint — o que será implementado, em qual ordem, com quais dependências e qual a abordagem técnica para cada task.

   #### Estrutura do SPRINT-DEVELOPMENT-PLANNING.md

   ```markdown
   # SPRINT-DEVELOPMENT-PLANNING.md — Plano de Desenvolvimento: Sprint {N}

   [Header: solução, projeto, sprint, stack, data]

   ## 1. Visão Geral
   - Sprint Goal: [do SPRINT-CARD.md]
   - Tasks a implementar: X
   - Ordem de execução: [sequencial | paralela | mista]
   - Stack: [stack detectada + versões]

   ## 2. Dependências entre Tasks
   [Diagrama ou lista mostrando a ordem e dependências entre T-XXX]
   [Formato: T-001 → (T-002, T-003, T-004, T-006, T-008) → T-005 → T-007]

   ## 3. Plano por Task

   ### T-XXX — [Nome da Tarefa]
   - **Critério DONE:** [do SPRINT-CARD.md]
   - **Estimativa:** Xd
   - **Abordagem:** [estratégia de implementação]
   - **Arquivos a criar:**
     | Arquivo | Tipo | Descrição |
     |:---|:---|:---|
     | `src/main/.../Classe.java` | 🆕 | Propósito da classe |
   - **Arquivos a modificar:**
     | Arquivo | Tipo | Descrição |
     |:---|:---|:---|
   - **Dependências:** [tasks que devem estar concluídas antes]
   - **Riscos:** [riscos específicos desta task]
   - **Skills aplicáveis:** [skills da stack que guiam a implementação]

   ### T-YYY — ...
   [Repetir para cada task da sprint]

   ## 4. Ordem de Execução
   [Lista numerada com a ordem exata de implementação, justificando escolhas]

   ## 5. Estratégia de Build e Verificação
   - Comando de build: `[comando]`
   - Comando de teste rápido: `[comando]`
   - Checkpoints: [momentos em que o build deve ser verificado]

   ## Rodapé
   [Indicação de geração por IA, data/hora]
   ```

   > ⚠️ Este artefato é gerado UMA VEZ, antes de iniciar a implementação da primeira task. Ele serve como roteiro para toda a Fase 1. Se o plano mudar durante a execução, atualizar este arquivo.

### Fase 2 — Implementação (por task)

9. **Para cada task do backlog, executar na ordem definida no SPRINT-DEVELOPMENT-PLANNING.md:**

   ```
   PARA cada task T-XXX:
       │
       ├── 1. LER o critério DONE no SPRINT-CARD.md
       │
       ├── 2. IMPLEMENTAR seguindo ARCHITECTURE.md (estrutura de diretórios, padrões,
       │      convenções de nomenclatura)
       │
       ├── 3. IDENTIFICAR a(s) stack(s) envolvidas nesta task específica:
       │      - Consultar PRD.md e/ou SPECS.md e/ou TASKS.md para confirmar a stack
       │      - Obter as melhores skills para realizar o desenvolvimento
       │      - Quando houver dúvida sobre qual stack ou skill usar, QUESTIONAR o humano
       │
       ├── 3.1 APLICAR as melhores práticas definidas pelas skills acionadas:
       │      - Seguir padrões idiomáticos da linguagem
       │      - Respeitar convenções do framework
       │      - Aplicar padrões de projeto documentados no ARCHITECTURE.md
       │      - Seguir regras de segurança do SECURITY.md
       │
       ├── 4. EXECUTAR o comando de compilação/build do projeto:
       │      - Descobrir o comando a partir do README.md, ou
       │      - Inferir pelo gerenciador de dependências:
       │        · pom.xml → `mvn compile`
       │        · build.gradle → `gradle build`
       │        · go.mod → `go build ./...`
       │        · package.json → `npm run build`
       │        · Cargo.toml → `cargo build`
       │        · pyproject.toml → `python -m compileall .`
       │      - Corrigir erros de compilação antes de prosseguir
       │
       ├── 5. EXECUTAR o comando de teste do projeto:
       │      - Descobrir o comando a partir do README.md, ou
       │      - Inferir pelo gerenciador:
       │        · Maven → `mvn test`
       │        · Gradle → `gradle test`
       │        · Go → `go test ./...`
       │        · npm/pnpm/yarn → `npm test`
       │        · Cargo → `cargo test`
       │        · pytest → `pytest`
       │      - Verificar se testes existentes não quebraram
       │
       └── 6. MARCAR task como concluída (atualizar status no SPRINT-CARD.md)
   ```

### Fase 3 — Planejamento dos Testes

10. **Gerar o artefato `SPRINT-TEST-PLANNING.md` ANTES de iniciar a implementação dos testes:**

    ```
    {SPRINT_DIR}/SPRINT-TEST-PLANNING.md
    ```

    Este artefato documenta o plano de testes da sprint — quais cenários cobrir, com quais ferramentas, em qual ordem e com qual meta de cobertura. Ele deriva do `SPRINT-TEST-SUITE.md` (o que testar) e detalha o COMO testar.

    #### Estrutura do SPRINT-TEST-PLANNING.md

    ```markdown
    # SPRINT-TEST-PLANNING.md — Plano de Testes: Sprint {N}

    [Header: solução, projeto, sprint, stack, data]

    ## 1. Visão Geral
    - Tasks implementadas: X
    - Cenários de teste mapeados: Y (do SPRINT-TEST-SUITE.md)
    - Meta de cobertura: ≥ 80% (padrão)
    - Ferramentas: [ex: JUnit 5 + Mockito + Testcontainers, pytest, Jest, Go testing]

    ## 2. Mapeamento Task → Cenários de Teste
    | Task | Cenário(s) | Nível | Ferramenta | Status |
    |:---|:---|:---|:---|:---:|
    | T-XXX | TC-XXX-001 | Unit | [framework] | ⬜ |
    | T-XXX | TC-XXX-002 | Integration | [framework] | ⬜ |

    ## 3. Estratégia por Nível de Teste

    ### 3.1 Testes Unitários
    - **Ferramenta:** [JUnit 5 + Mockito | pytest | Jest | Go testing | ...]
    - **Padrão:** [AAA (Arrange-Act-Assert) | Given-When-Then | table-driven]
    - **Localização:** `src/test/.../unit/`
    - **O que mockar:** [repositories, serviços externos, clock]
    - **O que NÃO mockar:** [entidades, DTOs, value objects]

    ### 3.2 Testes de Integração
    - **Ferramenta:** [Testcontainers + PostgreSQL | Docker Compose | ...]
    - **Localização:** `src/test/.../integration/`
    - **O que usar real:** [banco de dados, fila, cache]
    - **Dados de seed:** [estratégia para popular dados de teste]

    ### 3.3 Testes de Segurança (se aplicável)
    - **Foco:** [RBAC, multi-tenant isolation, rate limiting, JWT validation]
    - **Localização:** `src/test/.../security/`

    ## 4. Ordem de Execução dos Testes
    [Lista Priorizada:
     1. Testes unitários (sem dependências externas — rodam primeiro)
     2. Testes de integração (dependem de container/DB)
     3. Testes de segurança (dependem de contexto autenticado)]

    ## 5. Comandos de Execução
    - Unit: `[comando]`
    - Integration: `[comando]`
    - Coverage: `[comando]`
    - Lint/Quality: `[comando]`

    ## 6. Ações Manuais ou Externas

    > ⚠️ **Obrigatório quando aplicável.** Se qualquer cenário de teste NÃO puder ser
    > executado automaticamente pelo agente (ex: requer intervenção humana, acesso a
    > sistema externo não disponível no ambiente, configuração manual de infraestrutura,
    > validação visual), esta seção DEVE ser preenchida com instruções detalhadas.

    Para cada ação manual ou externa necessária, documentar:

    ### Ação X: [Título Descritivo]

    - **Cenário(s) relacionado(s):** [IDs dos cenários do SPRINT-TEST-SUITE.md]
    - **Quem executa:** [Humano — papel/função | Sistema externo — nome/URL]
    - **Pré-condições:** [o que precisa estar pronto antes de executar esta ação]
    - **Ambiente:** [dev | staging | CI | produção]

    **Passo a passo:**

    1. [Instrução concreta e acionável — o que fazer, não o que validar]
    2. [Ex: "Execute o comando `docker compose up postgres keycloak` na raiz do projeto"]
    3. [Ex: "Acesse http://localhost:8080/realms/fbso-platform e faça login com admin/admin"]
    4. [Ex: "Execute `curl -X POST http://localhost:8080/api/v1/tenants ...`"]

    **Resultado esperado:**

    - [O que o executor deve observar para confirmar que a ação foi bem-sucedida]
    - [Ex: "O endpoint retorna HTTP 201 com JSON contendo `status: PENDING_ONBOARDING`"]
    - [Ex: "O log do serviço exibe `Audit: TENANT CREATED — tenant=...`"]

    **Se falhar:**

    - [Ação corretiva ou contato para suporte]
    - [Ex: "Verificar se o Keycloak está rodando: `docker ps | grep keycloak`"]
    - [Ex: "Se o token JWT não for aceito, verificar se o realm `fbso-platform` foi importado"]

    **Evidência a coletar:**

    - [ ] Screenshot da tela/terminal
    - [ ] Log gerado
    - [ ] Resposta HTTP (headers + body)

    ### Ação Y: ...

    > **Formato das instruções:**
    > - Usar **verbos no imperativo** (Execute, Acesse, Verifique, Configure)
    > - Comandos devem ser **copiáveis** (blocos de código com syntax highlighting)
    > - URLs, portas e credenciais devem ser **explícitas** (não usar placeholders)
    > - Se a ação exigir uma ordem específica, **numerar** os passos

    ## Rodapé
    [Indicação de geração por IA, data/hora]
    ```

    > ⚠️ Este artefato é gerado UMA VEZ, após concluir a implementação de todas as tasks e antes de iniciar os testes. Se o escopo de testes mudar durante a execução, atualizar este arquivo.

### Fase 4 — Implementação dos Testes

11. **Implementar testes IMEDIATAMENTE após o código, seguindo o SPRINT-TEST-PLANNING.md:**

   ```
   PARA cada task T-XXX:
       │
       ├── 1. CONSULTAR SPRINT-TEST-SUITE.md — quais cenários se aplicam?
       ├── 2. MAPEAR para TEST_PLAN.md — confirmar IDs e níveis de teste
       │
       ├── 3. IMPLEMENTAR os testes seguindo as práticas da stack:
       │      - Testes unitários (ex: JUnit 5 + Mockito, pytest, Jest, Go testing)
       │        para lógica de negócio e validações
       │      - Testes de integração (ex: Testcontainers, Docker Compose, Go integration)
       │        para repositórios, acesso a dados e APIs
       │      - Testes de segurança (se aplicável à sprint)
       │        para RBAC, isolamento multi-tenant, rate limiting
       │
       ├── 4. EXECUTAR testes unitários — confirmar verde
       ├── 5. EXECUTAR testes de integração — confirmar verde
       └── 6. VERIFICAR cobertura de código — meta ≥ 80%
              - Comando depende da stack:
                · Java: `mvn jacoco:check`
                · Go: `go test -cover ./...`
                · JS/TS: `jest --coverage`
                · Python: `pytest --cov`
                · Rust: `cargo tarpaulin`
   ```

### Fase 5 — Validação de Qualidade

12. **Executar verificações estáticas e de estilo:**

    ```
    - Linter/formatador da stack:
      · Java: `mvn checkstyle:check pmd:check`
      · Go: `go vet ./... && golangci-lint run`
      · JS/TS: `eslint . && prettier --check .`
      · Python: `ruff check . && mypy .`
      · Rust: `cargo clippy && cargo fmt --check`

    - Zero warnings. Zero violations.
    ```

13. **Executar suite completa da sprint:**

    ```
    - Comando de teste completo (unit + integration)
    - Cobertura ≥ 80% (meta padrão; ajustar se ARCHITECTURE.md ou TEST_PLAN.md definirem meta diferente)
    ```

### Fase 6 — Tratamento de Falhas

Se algum teste falhar durante a execução:

1. **Auto-Correção Autônoma (até 3 tentativas):** Analise o stack trace/saída de erro, identifique se o problema está na lógica do código ou na estrutura do teste, e corrija.
2. **Tratamento de Loops:** Se o mesmo erro persistir por **3 tentativas**, PARE imediatamente.
3. **Registro de Impedimento:** Crie `{SPRINT_DIR}/IMPEDIMENT-SPRINT-{SPRINT_NUMBER}.md` com:
   - Task que falhou (T-XXX) e mensagem de erro exata
   - O que foi tentado para corrigir
   - Suspeita do motivo (limitação arquitetural, ambiguidade na SPECS.md, conflito de dependências, etc.)
   - Propostas adicionais de solução (se houver)
4. **Notifique o humano** e aguarde instruções antes de alterar qualquer outro arquivo.

### Fase 7 — Code Review (PonyTail)

14. **Executar auditoria PonyTail (`ponytail-audit`):**

    ```
    - Acionar o skill `ponytail-audit` sobre todo o código implementado na sprint
    - O skill analisa o código em busca de:
      · Código morto ou não utilizado (YAGNI — Rung 1)
      · Duplicação de código (DRY — Rung 2)
      · Dependências desnecessárias (Rungs 3 e 4)
      · Desvios dos padrões documentados no ARCHITECTURE.md (Rung 5)
      · Complexidade excessiva (Rung 6 — poderia ser mais simples?)
      · Código não idiomático para a stack detectada
    - Output: lista de findings com severidade (Critical/High/Medium/Low),
      arquivo, linha, descrição e recomendação
    ```

15. **Executar revisão PonyTail (`ponytail-review`):**

    ```
    - Acionar o skill `ponytail-review` sobre o código implementado na sprint
    - O skill revisa:
      · Qualidade do código (legibilidade, clareza, manutenibilidade)
      · Aderência aos padrões de projeto do ARCHITECTURE.md
      · Segurança (alinhamento com SECURITY.md)
      · Cobertura de bordas não tratadas (null safety, exceções, timeouts)
      · Consistência com o restante do codebase
    - Output: lista de pontos de melhoria com justificativa e sugestão de refatoração
    ```

16. **Consolidar achados e gerar relatório de ajustes:**

    ```
    SE houver findings do ponytail-audit OU pontos do ponytail-review:
    
      ├── 16.1 GERAR o relatório consolidado:
      │       {SPRINT_DIR}/PONYTAIL-REPORT-ADJUST-SPRINT-{SPRINT_NUMBER}.md
      │
      ├── 16.2 CONSOLIDAR os achados de ambas as auditorias no relatório,
      │       agrupando por arquivo e priorizando por severidade
      │
      └── 16.3 PROSSEGUIR para o passo 17 (executar ajustes)
    
    SENÃO (zero achados relevantes):
    
      └── PULAR para a Fase 8 (Pós-implementação — Sanity Check)
    ```

    #### Estrutura do `PONYTAIL-REPORT-ADJUST-SPRINT-{N}.md`

    ```markdown
    # PONYTAIL-REPORT-ADJUST-SPRINT-{N} — Relatório de Ajustes Pós-Code Review

    [Header: solução, projeto, sprint, stack detectada, data da revisão]

    ## 1. Resumo da Revisão
    - Skills acionados: `ponytail-audit`, `ponytail-review`
    - Total de achados: N
    - Por severidade:
      | Critical | High | Medium | Low |
      |:---:|:---:|:---:|:---:|
      | X | Y | Z | W |

    ## 2. Achados — `ponytail-audit`
    | ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
    |:---|:---|:---|:---:|:---|:---|
    | PA-001 | High | `src/.../Classe.java` | 42 | Código não utilizado — método `getX()` sem referências | Remover método |
    | PA-002 | Medium | `src/.../Servico.java` | 87 | Duplicação com `Util.java:120` | Extrair para método compartilhado |
    | ... | ... | ... | ... | ... | ... |

    ## 3. Achados — `ponytail-review`
    | ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
    |:---|:---|:---|:---:|:---|:---|
    | PR-001 | Medium | `src/.../Controller.java` | 55 | Falta validação de entrada — possível NPE | Adicionar `@NotNull` / null guard |
    | PR-002 | Low | `src/.../Service.java` | 120 | Nome do método não segue convenção do projeto | Renomear para `findByTenantId` |
    | ... | ... | ... | ... | ... | ... |

    ## 4. Plano de Ajustes
    [Lista priorizada de ajustes a executar, agrupados por arquivo e ordenados
     por severidade (Critical → High → Medium → Low)]

    ### Arquivo: `src/main/.../Classe.java`
    - [ ] PA-001 (High): Remover método `getX()` não utilizado
    - [ ] PR-003 (Medium): Adicionar tratamento de exceção no método `processar()`

    ### Arquivo: `src/main/.../Servico.java`
    - [ ] PA-002 (Medium): Extrair lógica duplicada para método compartilhado
    ...

    ## 5. Execução dos Ajustes
    [Registro do que foi alterado, arquivo por arquivo, com a justificativa e o
     resultado do build pós-ajuste]

    | ID | Arquivo | Ação | Resultado |
    |:---|:---|:---|:---:|
    | PA-001 | `Classe.java` | Método `getX()` removido | ✅ Compila |
    | PA-002 | `Servico.java`, `Util.java` | Extraído `metodoCompartilhado()` | ✅ Compila |
    | ... | ... | ... | ... |

    ## 6. Build Pós-Ajustes
    - Comando: `[build command]`
    - Resultado: ✅ SUCCESS / ❌ FAILURE
    - Testes rápidos: N/N passando

    ## Rodapé
    [Indicação de geração por IA, skills utilizados, data/hora da geração]
    ```

17. **Executar ajustes a partir do relatório:**

    ```
    PARA cada achado no PONYTAIL-REPORT-ADJUST-SPRINT-{N}.md, na ordem do plano (§4):
        │
        ├── 1. APLICAR a correção recomendada no código-fonte
        ├── 2. MARCAR o achado como concluído no relatório (§5)
        ├── 3. EXECUTAR build/compilação para validar
        │      - Se falhar → reverter ajuste, marcar como ❌ e documentar motivo
        └── 4. ATUALIZAR a tabela de execução (§5) com o resultado
    ```

18. **Retornar à Fase 3 — Planejamento dos Testes:**

    ```
    Após concluir todos os ajustes de código:
    
    ├── 18.1 VOLTAR para a Fase 3 — Planejamento dos Testes (passo 10)
    │       - O SPRINT-TEST-PLANNING.md pode precisar ser atualizado
    │         para refletir as mudanças de código
    │
    ├── 18.2 REEXECUTAR Fases 3, 4 e 5 (planejamento → testes → qualidade)
    │       para garantir que os ajustes não quebraram testes existentes
    │       e que novos cenários de borda sejam cobertos
    │
    └── 18.3 CONTROLE DE CICLOS:
            - Máximo de 2 ciclos completos (Fase 3→4→5→7)
            - Se após o 2º ciclo ainda houver achados Critical/High não resolvidos,
              registrar no SPRINT-EXECUTION-REPORT.md (§8) e prosseguir
            - Achados Medium/Low remanescentes são documentados como
              dívida técnica na seção de observações do relatório final
    ```

    > ⚠️ **Por que voltar à Fase 3?** Ajustes de código motivados pelo PonyTail podem:
    > - Alterar assinaturas de métodos → testes unitários precisam ser atualizados
    > - Introduzir novos métodos/classes → novos testes podem ser necessários
    > - Remover código morto → testes obsoletos devem ser removidos
    > - Alterar fluxos de exceção → cenários de borda precisam ser cobertos
    >
    > O ciclo de retorno garante que o código ajustado mantenha a mesma qualidade
    > de testes do código original.

### Fase 8 — Pós-implementação (Sanity Check)

19. **Limpeza:** Remover código comentado, prints de debug, imports não usados, arquivos temporários.
20. **Git Status:** Listar todos os arquivos modificados ou criados (`git status --short`).
21. **Localização:** Validar que os arquivos respeitam o `ARCHITECTURE.md` (estrutura de diretórios, nomes de arquivo, convenções).
22. **Segurança:** Revisão final contra `SECURITY.md` — nenhuma regra violada. Verificar:
    - Nenhuma credencial ou segredo hardcoded
    - Queries parametrizadas (SQL injection prevention)
    - Autorização implementada onde exigido (RBAC, scopes)
    - Dados sensíveis não expostos em logs ou respostas de erro
23. **Atualizar SPRINT-CARD.md:** Marcar cada task concluída no backlog da sprint.

### Fase 9 — Geração do Relatório de Execução (OBRIGATÓRIO)

24. **Gerar arquivo de report em:**
    ```
    {SPRINT_DIR}/SPRINT-{SPRINT_NUMBER}-EXECUTION-REPORT.md
    ```

#### Estrutura do Relatório

```markdown
# SPRINT-{N}-EXECUTION-REPORT.md — Relatório de Execução: Sprint {N}

[Header: solução, projeto, sprint, stack detectada, data da execução, tasks executadas]

## 1. Resumo da Execução
- Tasks executadas: X/Y
- Tasks com sucesso: X
- Tasks com falha: 0
- Tempo total estimado: Xd (do SPRINT-CARD.md)
- Tempo total gasto: (preencher ao final)

## 2. Stack e Skills Utilizadas
- Stack detectada: [linguagem + framework + banco + infra]
- Fonte da stack: [parâmetro {STACK} | PRD.md | SPECS.md | inferido do README.md]
- Skills acionadas: [lista das skills com justificativa para cada uma]

## 3. Tasks Executadas
| ID | Tarefa | Status | Testes | Cobertura | Observações |
|:---|:---|:---:|:---:|:---:|:---|
| T-XXX | Descrição | ✅ | N/N passando | XX% | — |
| ... | ... | ... | ... | ... | ... |

## 4. Arquivos Criados ou Modificados
| Ação | Arquivo | Task | Descrição da Mudança |
|:---|:---|:---|:---|
| 🆕 | `src/main/.../NovoArquivo.ext` | T-XXX | Nova classe/função/módulo — propósito |
| 🔄 | `src/main/.../Existente.ext` | T-YYY | Modificado — razão |

## 5. Evidências de Testes
- Comando de build executado: `[comando]` → [resultado]
- Comando de teste executado: `[comando]` → [resultado]
- Total de testes: N
- Status: ✅ 100% PASS
- Cobertura: XX% (linhas), XX% (branchs)
- Cenários do SPRINT-TEST-SUITE.md executados: N/N

## 6. Validação de Segurança
- [ ] Nenhuma credencial ou dado sensível hardcoded
- [ ] Queries usam parametrização (proteção contra injection)
- [ ] Controles de acesso implementados onde exigido pela sprint
- [ ] Dados pessoais não expostos em logs ou respostas HTTP
- [ ] Respostas de erro não expõem stack traces ou detalhes internos

## 7. Validação de Arquitetura
- [ ] Estrutura de diretórios segue ARCHITECTURE.md
- [ ] Convenções de nomenclatura respeitadas
- [ ] Padrões de projeto documentados nas ADRs foram seguidos

## 8. Desvios e Observações
- Desvios das specs originais (se houver) e justificativa
- Decisões de design tomadas durante a implementação
- Comentários `// ponytail:` adicionados (atalhos intencionais documentados)
- Dificuldades encontradas e como foram resolvidas

## 9. Próximos Passos
- Tasks restantes na sprint (se houver)
- Pré-requisitos para a próxima sprint (conforme SPRINT-REVIEW.md)
- Recomendações para a review com o PO

## Rodapé
[Indicação de geração por IA, skills utilizados, data/hora da geração]
```

---

## Skills (Acionamento Dinâmico)

> **Este prompt não declara skills fixas.** As skills são descobertas dinamicamente na **Fase 0, passo 6** com base na stack do projeto. O skill `001-skills-inventory` é usado como ponto de partida para listar as skills disponíveis.

### Protocolo de Seleção de Skills

```
1. INVOCAR `001-skills-inventory` para listar skills disponíveis
2. CRUZAR cada componente da stack com as skills:
   - Nome da skill contém o componente? (ex: "java", "golang", "spring", "quarkus")
   - Descrição da skill menciona o componente?
3. PRIORIZAR skills específicas da stack sobre skills genéricas
4. SE houver ambiguidade (múltiplas skills para o mesmo componente) →
   PERGUNTAR ao humano qual skill prefere
5. SE não houver skill para um componente →
   Usar conhecimento geral da stack e documentar a ausência no relatório (§8)
6. REGISTRAR a seleção final no relatório de execução (§2)
```

### Skills Transversais (acionadas em TODAS as stacks)

| Skill | Modo | Função |
|:---|:---|:---|
| `ponytail` | full | Escada YAGNI de 7 rungs — controle de escopo do código gerado |
| `caveman` | full | Compressão de prosa interativa (comunicação durante o desenvolvimento) |
| `security-review` | automático | Revisão de segurança — OWASP Top 10, boas práticas |
| `code-review` | automático | Revisão de qualidade do código gerado |
| `verification-before-completion` | automático | Verificar que a implementação funciona antes de declarar concluída |

> ⚠️ **caveman e ponytail NÃO atuam sobre artefatos permanentes** (SPRINT-CARD.md, SPRINT-TEST-SUITE.md, SPRINT-REVIEW.md, EXECUTION-REPORT.md). O caveman comprime apenas a comunicação interativa; o report final é gerado em prosa normal.

---

## Protocolo de Testes

Durante a execução, para cada task implementada:

```
1. BUILD/COMPILE   ← Compilar/buildar (falha = corrigir antes de testar)
2. UNIT TEST       ← Testes unitários
3. INTEGRATION TEST ← Testes de integração (se aplicável)
4. COVERAGE        ← Verificar cobertura ≥ 80%
```

Os comandos específicos são descobertos na **Fase 0, passo 6** a partir do README.md ou inferidos pelo gerenciador de dependências do projeto.

Se os testes falharem:
1. Analisar saída de erro e corrigir (até 3 tentativas)
2. Se persistir → criar `{SPRINT_DIR}/IMPEDIMENT-SPRINT-{SPRINT_NUMBER}.md`
3. Notificar humano e aguardar

---

## Output Esperado

| Output | Descrição |
|:---|:---|
| Código-fonte | Seguindo ARCHITECTURE.md e melhores práticas da stack |
| Testes | Unit + Integration + Security (conforme SPRINT-TEST-SUITE.md) |
| Build/Compile | ✅ Sucesso |
| Testes | ✅ Todos verdes |
| Cobertura | ≥ 80% (meta padrão) |
| Linter/Formatter | ✅ Zero warnings |
| **ARTEFATO** | `{SPRINT_DIR}/SPRINT-{N}-EXECUTION-REPORT.md` |
| **ARTEFATO** | `{SPRINT_DIR}/PONYTAIL-REPORT-ADJUST-SPRINT-{N}.md` (se houver achados) |
| **ARTEFATO** | `{SPRINT_DIR}/SPRINT-DEVELOPMENT-PLANNING.md` |
| **ARTEFATO** | `{SPRINT_DIR}/SPRINT-TEST-PLANNING.md` |

---

## Anti-Padrões

| ❌ NÃO fazer | ✅ Fazer |
|:---|:---|
| Implementar tasks sem ler SPRINT-CARD.md | Sempre começar pelo SPRINT-CARD.md |
| Presumir a stack sem verificar os docs do projeto | Detectar stack de PRD.md, SPECS.md ou README.md |
| Pular testes e deixar "para depois" | Escrever testes IMEDIATAMENTE após o código |
| Implementar tasks de sprints futuras | Uma sprint de cada vez, na ordem dos marcos |
| Ignorar o SPRINT-TEST-SUITE.md | Usá-lo como checklist de qualidade |
| Gerar report fora da pasta da sprint | Salvar em `{SPRINT_DIR}/SPRINT-{N}-EXECUTION-REPORT.md` |
| "Vou adicionar X porque vai ser útil depois" | Seguir YAGNI — implementar apenas o que a task pede |
| Adicionar dependência nova para funcionalidade trivial | Usar o que já existe no projeto (stdlib, deps existentes) |
| Usar padrão diferente do ARCHITECTURE.md "porque é melhor" | Seguir o padrão documentado; se precisar mudar, propor ADR |
| Deixar stack trace ou detalhes internos em respostas de erro | Tratar erros conforme definido no SPECS.md e SECURITY.md |
| Pular a Fase 7 (Code Review PonyTail) "porque o código já está bom" | `ponytail-audit` e `ponytail-review` são obrigatórios — sempre há algo a melhorar |
| Ignorar achados do PonyTail e prosseguir direto para o Sanity Check | Aplicar os ajustes e revalidar com Fases 3→4→5 (máx. 2 ciclos) |
| Fazer ajustes do PonyTail sem reexecutar os testes | Sempre voltar à Fase 3 após ajustes para garantir que nada quebrou |

---

## Observações

1. **O SPRINT-CARD.md é o roteiro.** Cada task tem critério DONE explícito — não considere uma task concluída até que todos os critérios DONE sejam atendidos.

2. **O SPRINT-TEST-SUITE.md é o checklist de qualidade.** Se um cenário de teste falha, a task associada não está concluída, mesmo que o código compile.

3. **Respeitar a ordem das tasks.** O SPRINT-CARD.md lista as tasks na ordem recomendada de execução. Tasks com dependências devem ser executadas em sequência.

4. **Artefatos de sprint são derivados.** Se durante a execução você encontrar uma inconsistência entre os artefatos de sprint e os documentos-mestre, os documentos-mestre prevalecem. Documente a inconsistência no relatório de execução (§8).

5. **Relatório de execução é obrigatório.** Não é opcional. É a evidência de que a sprint foi executada e o ponto de partida para a review com o PO.

6. **Na dúvida sobre a stack, PERGUNTE.** É melhor confirmar "este projeto é Java 25 + Spring Boot + PostgreSQL, correto?" do que assumir errado e gerar retrabalho.

7. **Sprints de fundação (ex: Setup, Segurança) são diferentes.** Elas podem não ter features de negócio ou endpoints REST. Os testes tendem a ser estruturais (build, migração, segurança). O SPRINT-CARD.md sinaliza isso.

8. **Code Review com PonyTail é obrigatório.** A Fase 7 (`ponytail-audit` + `ponytail-review`) deve ser executada em TODAS as sprints, independentemente do tamanho. Se zero achados forem encontrados, o relatório pode ser omitido, mas a auditoria deve ser executada e registrada no relatório de execução.

9. **Ciclo de ajustes tem limite.** Após a Fase 7, o retorno à Fase 3 é esperado para revalidar testes. Porém, o ciclo Fase 3→4→5→7 não deve exceder 2 iterações. Achados remanescentes após 2 ciclos devem ser documentados como dívida técnica no relatório final.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 4.0 | 14/07/2026 | Adicionada Fase 7 — Code Review (PonyTail) com `ponytail-audit` e `ponytail-review`. Geração de `PONYTAIL-REPORT-ADJUST-SPRINT-{N}.md`. Ciclo de retorno à Fase 3 para revalidação de testes pós-ajustes. Fases 7→8 (Sanity Check), 8→9 (Relatório). Total de 9 fases, 24 passos | Time de Arquitetura |
| 3.0 | 14/07/2026 | Adicionados artefatos de planejamento: SPRINT-DEVELOPMENT-PLANNING.md (Fase 1) e SPRINT-TEST-PLANNING.md (Fase 3). Fases renumeradas: 0→5→8. Total de 8 fases, 19 passos | Time de Arquitetura |
| 2.0 | 14/07/2026 | Generalização para stack-agnóstico. Stack detectada dos documentos do projeto (PRD.md, SPECS.md). Skills acionadas dinamicamente via `001-skills-inventory`. Comandos de build/teste/coverage inferidos do gerenciador de dependências. Removidas todas as referências hardcoded a Java/Spring Boot/Maven. Adicionada §2 (Stack e Skills) ao relatório de execução | Time de Arquitetura |
| 1.0 | 14/07/2026 | Criação inicial: execução de tasks de sprint em Java/Spring Boot, adaptado de PROMPT-EXECUTE-TASK.md | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, ponytail.*
