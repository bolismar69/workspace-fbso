# PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT

## Contexto

Este prompt orquestra a **identificação e catalogação de débitos técnicos** antes do início de uma nova sprint, como parte do fluxo Spec-Driven Development. Ele atua como uma **auditoria técnica multidisciplinar** que combina **7 skills complementares** para varrer documentação, código, arquitetura, complexidade, dependências e qualidade, produzindo um documento único e acionável de débitos técnicos.

**Princípios fundamentais:**

1. **Prevenção, não correção tardia:** Identificar débitos técnicos ANTES de iniciar a sprint evita que decisões erradas de sprints anteriores se propagem e contaminem o novo desenvolvimento.
2. **Multidisciplinaridade:** Cada skill cobre um ângulo diferente — engenharia, consistência docs×código, decisões arquiteturais, complexidade, qualidade de código, dívida técnica estrutural e segurança de dependências. A combinação dos 7 ângulos produz uma visão completa e sem pontos cegos.
3. **Decisão humana no loop:** O documento de débitos é apresentado ao time para decisão explícita sobre o que tratar agora vs. depois. O agente NUNCA decide unilateralmente o que corrigir.
4. **Rastreabilidade completa:** Cada débito aceito para correção é rastreado até os documentos de negócio e técnicos impactados, garantindo que o impacto seja visível em toda a cadeia.
5. **Código DT imutável:** Cada débito técnico recebe um código `DT-XXX` no documento `IDENTIFIED-TECHNICAL-DEBT-{SPRINT_NAME}.md`. Este código é **permanente e imutável** — uma vez atribuído, NUNCA é alterado, mesmo que o débito seja movido para outra sprint. O código DT-XXX serve como chave primária para rastreamento em ferramentas externas (Jira, Trello, GitHub Issues) e mudanças de numeração quebrariam links, dashboards e integrações.

**Quando executar:** Regularmente antes de iniciar uma nova sprint. Também pode ser executado sob demanda quando o time identifica indícios de acúmulo de débito técnico.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço) | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SPRINT_NUMBER}` | Número da sprint atual (a que está para iniciar) | `3` |
| `{SPRINT_NAME}` | Nome curto da sprint atual (kebab-case) | `sprint-03-portal-admin` |
| `{STACK}` | Stack tecnológica principal (opcional — se omitido, auto-detectado) | `Java 25 + Spring Boot + PostgreSQL` |

---

## Documentos de Referência

### Documentos-Mestre (baseline de verdade)

```
SPECS_DIR = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/

Ler obrigatoriamente:
    ├── {SPECS_DIR}/PRD.md                 ← Stack, escopo, entidades, ADRs
    ├── {SPECS_DIR}/SPECS.md               ← Regras de negócio, endpoints, validações
    ├── {SPECS_DIR}/TASKS.md               ← Fonte da verdade das tarefas (todas as sprints)
    ├── {SPECS_DIR}/TEST_PLAN.md           ← Fonte da verdade dos cenários de teste
    └── {SPECS_DIR}/ARCHITECTURE.md        ← Estrutura, ADRs, padrões, diagramas
```

### Artefatos da Sprint Atual

```
SPRINT_DIR = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/sprints/{SPRINT_NAME}/

Ler se existirem:
    ├── {SPRINT_DIR}/SPRINT-CARD.md
    ├── {SPRINT_DIR}/SPRINT-TEST-SUITE.md
    ├── {SPRINT_DIR}/SPRINT-REVIEW.md
    ├── {SPRINT_DIR}/SPRINT-DEVELOPMENT-PLANNING.md (ou DRAFT)
    └── {SPRINT_DIR}/DOCS-SPRINT-CAVEMAN-REVIEW.md (se review anterior existir)
```

### Documentos de Negócio (se existirem no projeto)

```
BUSINESS_DIR = {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/

Ler se existirem:
    ├── 01-PROJECT-CHARTER-*.md
    ├── 02-BUSINESS-REQUIREMENTS.md
    ├── 03-EPICS.md
    ├── 04-FEATURES.md
    ├── 05-USER-STORYS.md
    └── DEFINITION_OF_DONE.md
```

### Código-Fonte

```
SOURCE_DIR = {SOLUTION_PATH}/src/

Escanear:
    ├── main/java/   ← Código de produção (estrutura de pacotes, classes, padrões)
    ├── main/resources/  ← Configurações, migrations, application.yml
    └── test/java/   ← Testes existentes, cobertura, padrões
```

---

## Skills Empregadas

Este prompt orquestra **7 skills** em paralelo, cada uma cobrindo um ângulo diferente. Cada skill produz seus achados no seu formato nativo, que são depois consolidados no documento único.

| # | Skill | Ângulo | O que encontra |
|:---:|:---|:---|:---|
| 1 | `engineering-skills` | Engenharia | Violações de padrões de código, más práticas de arquitetura, dependências incorretas, configurações erradas, falta de cobertura de testes, vulnerabilidades de segurança |
| 2 | `caveman-review` | Consistência | Inconsistências entre documentação e código, métricas divergentes entre artefatos, versões desatualizadas, referências cruzadas quebradas |
| 3 | `superpowers:brainstorming` | Decisões | Decisões arquiteturais questionáveis, trade-offs não documentados, premissas inválidas, gaps no planejamento que afetam a sprint atual |
| 4 | `ponytail-review` | Complexidade | Código morto, abstrações desnecessárias (YAGNI), dependências que podem ser substituídas por stdlib/nativas, lógica que pode ser simplificada |
| 5 | `code-review` | Qualidade | Bugs, vulnerabilidades de segurança, anti-padrões, problemas de performance via CodeRabbit CLI — revisão automatizada com severidade (Critical/Warning/Info) |
| 6 | `codebase-cleanup-tech-debt` | Dívida Estrutural | Código duplicado, complexidade ciclomática, god classes, métodos longos, dependências circulares, cobertura de testes insuficiente, documentação ausente — com métricas quantificadas e análise de ROI |
| 7 | `codebase-cleanup-deps-audit` | Dependências | Vulnerabilidades em dependências (CVEs), licenças incompatíveis, pacotes desatualizados (major/minor), supply chain risks, dependências transitivas problemáticas |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 6 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

Se `{STACK}` não for informado, auto-detectar a partir de `PRD.md` e `ARCHITECTURE.md`.

### Passo 1 — Carregar Contexto Completo

```
1. Ler TODOS os documentos-mestre (PRD, SPECS, TASKS, TEST_PLAN, ARCHITECTURE)
2. Ler TODOS os artefatos da sprint atual (SPRINT-CARD, SPRINT-TEST-SUITE, SPRINT-REVIEW, etc.)
3. Escanear a estrutura completa do código-fonte:
   - Listar todos os pacotes e classes
   - Listar todas as migrations Flyway
   - Listar todos os testes e sua cobertura
   - Verificar configurações (application.yml, pom.xml/build.gradle)
4. Se existir DOCS-SPRINT-CAVEMAN-REVIEW.md de sprint anterior, carregar como baseline
```

### Passo 2 — Executar as 7 Skills em Paralelo

Cada skill é invocada como um agente independente, com seu próprio escopo. Os agentes 1-4 são baseados em skills locais (zero dependências externas). Os agentes 5-7 requerem verificação de pré-requisitos.

**Verificação de pré-requisitos (Agentes 5-7):**

```
Agente 5 (code-review): Verificar se CodeRabbit CLI está instalado e autenticado.
  → coderabbit --version && coderabbit auth status
  → Se NÃO disponível: executar mesmo assim com revisão manual de código
    (análise estática sem CLI, usando padrões conhecidos de bugs/security)

Agente 6 (codebase-cleanup-tech-debt): Sem dependências externas.
  → Análise puramente baseada em leitura de código e métricas estáticas.

Agente 7 (codebase-cleanup-deps-audit): Verificar arquivos de build.
  → pom.xml, build.gradle, package.json, Cargo.toml, etc.
  → Se disponível, usar ferramentas nativas: mvn dependency:analyze, npm audit, etc.
```

#### Agente 1: engineering-skills

```
Escopo: Código-fonte + Configurações + Build
Prompt do agente:
  "Usando as skills do engineering-skills (senior-backend, senior-architect,
   code-reviewer, senior-security), faça uma análise completa do código em
   {SOLUTION_PATH}/src/ e identifique:
   1. Violações de padrões de arquitetura (ex: camadas puladas, responsabilidades erradas)
   2. Código que não segue as convenções do projeto (ex: BaseRepository não usado)
   3. Vulnerabilidades de segurança (ex: SQL injection, falta de validação)
   4. Configurações incorretas ou duplicadas
   5. Dependências ausentes ou incorretas no build file
   6. Cobertura de testes insuficiente ou testes mal escritos
   7. Código de sprints anteriores com bugs ou incompleto
   Para cada achado, indique: localização exata (arquivo:linha), severidade,
   e se é impeditivo para a Sprint {SPRINT_NUMBER}."
```

#### Agente 2: caveman-review

```
Escopo: Documentação × Código (triangulação)
Prompt do agente:
  "Execute um caveman-review completo triangulando:
   1. Documentos-mestre × Artefatos da sprint (métricas, versões, referências)
   2. Artefatos da sprint × Código real (o que a doc diz vs. o que o código tem)
   3. Artefatos da sprint × Artefatos da sprint (consistência cruzada)
   Use o formato caveman (L<linha>: <severidade> <problema>. <correção>.)
   Para cada achado, classifique como: 🔴 bug, 🟡 risk, 🔵 nit.
   Foco especial em números divergentes (RNs, endpoints, cenários de teste)
   e em código que contradiz o que a documentação afirma."
```

#### Agente 3: superpowers:brainstorming

```
Escopo: Decisões arquiteturais e de planejamento
Prompt do agente:
  "Analise as decisões arquiteturais e de planejamento documentadas em
   ARCHITECTURE.md, PRD.md, e SPRINT-DEVELOPMENT-PLANNING.md.
   Identifique:
   1. Decisões que se mostraram incorretas à luz do código existente
   2. Premissas inválidas (ex: 'vamos usar X' mas o código usa Y)
   3. Gaps no planejamento que afetam a Sprint {SPRINT_NUMBER}
   4. Riscos arquiteturais não documentados
   5. Trade-offs que deveriam ser revistos antes de iniciar a sprint
   Para cada achado, indique: qual decisão está em questão, qual era a
   premissa original, e por que ela precisa ser revista agora."
```

#### Agente 4: ponytail-review

```
Escopo: Complexidade e over-engineering no código
Prompt do agente:
  "Execute um ponytail-review no código em {SOLUTION_PATH}/src/.
   Use o formato ponytail (tags: delete, stdlib, native, yagni, shrink).
   Identifique:
   1. Código morto ou não utilizado
   2. Abstrações com uma única implementação (YAGNI)
   3. Dependências que replicam funcionalidades nativas da plataforma
   4. Lógica que pode ser simplificada sem perder funcionalidade
   5. Classes/métodos que existem mas nunca são chamados
   Foco em código das sprints anteriores (Sprints 1 a {SPRINT_NUMBER}-1).
   Termine com 'net: -N lines possible.'"
```

#### Agente 5: code-review (CodeRabbit)

```
Escopo: Bugs, vulnerabilidades e qualidade de código
Pré-requisito: CodeRabbit CLI (coderabbit --version). Se indisponível, fallback para revisão manual.
Prompt do agente:
  "Execute uma revisão de código completa em {SOLUTION_PATH}/src/.

   Se CodeRabbit CLI disponível:
     - Executar: coderabbit review --agent -t all
     - Agrupar por severidade: Critical, Warning, Info
     - Para cada Critical/Warning: verificar se é falso positivo
     - Criar task list para issues confirmados

   Se CodeRabbit CLI NÃO disponível (fallback manual):
     - Revisar diffs entre HEAD e main/master para código das Sprints 1-{SPRINT_NUMBER}
     - Identificar: bugs, security issues, anti-padrões, race conditions, null handling
     - Usar padrões conhecidos: SQL injection, XSS, path traversal, insecure deserialization
     - Para cada achado, indicar: arquivo:linha, severidade, evidência, fix sugerido

   Foco em código novo/modificado das sprints anteriores.
   NÃO revisar código de dependências externas (node_modules, .m2, etc.)."
```

#### Agente 6: codebase-cleanup-tech-debt

```
Escopo: Dívida técnica estrutural com métricas quantificadas e análise de ROI
Prompt do agente:
  "Execute uma análise completa de dívida técnica estrutural em {SOLUTION_PATH}/src/.

   Categorias a inspecionar:

   1. CODE DEBT — Duplicação, Complexidade, Estrutura:
      - Código duplicado (copy-paste, lógica repetida em múltiplos arquivos)
      - Complexidade ciclomática alta (>10) — identifique hotspots
      - Métodos longos (>50 linhas) e god classes (>500 linhas, >20 métodos)
      - Dependências circulares entre pacotes/módulos
      - Feature envy (métodos que usam mais dados de outras classes que da própria)

   2. ARCHITECTURE DEBT — Design e Tecnologia:
      - Abstrações ausentes ou com vazamento (leaky abstractions)
      - Fronteiras arquiteturais violadas (ex: controller acessando DAO direto)
      - Componentes monolíticos que deveriam ser modularizados
      - Uso de APIs deprecadas ou padrões legados

   3. TESTING DEBT — Cobertura e Qualidade:
      - Caminhos críticos sem teste (identifique pelo @Auditable, @Transactional)
      - Testes frágeis (dependentes de ambiente, sleep, ordem de execução)
      - Falta de testes de integração (Testcontainers não usado apesar de configurado)
      - Cobertura abaixo da meta (JaCoCo <80% em classes da sprint atual)

   4. DOCUMENTATION DEBT:
      - APIs públicas sem JavaDoc/OpenAPI
      - Lógica complexa sem comentário explicativo
      - ADRs ausentes para decisões arquiteturais já implementadas

   5. INFRASTRUCTURE DEBT:
      - Passos manuais de deploy (verificar Dockerfile, docker-compose)
      - Ausência de health checks ou métricas de produção
      - Scripts de seed/rollback inexistentes ou desatualizados

   Para CADA achado, quantificar:
   - Localização exata (arquivo:linha)
   - Métrica (ex: complexidade=14, duplicação=85 linhas em 3 arquivos)
   - Impacto em velocidade (horas/perda por mês)
   - ROI estimado da correção (horas investidas vs. horas economizadas)
   - Se é impeditivo para a Sprint {SPRINT_NUMBER}

   Output agrupado por categoria, ordenado por severidade (Critical → High → Medium → Low)."
```

#### Agente 7: codebase-cleanup-deps-audit

```
Escopo: Segurança de dependências, licenças e supply chain
Prompt do agente:
  "Execute uma auditoria completa de dependências em {SOLUTION_PATH}/.

   1. INVENTÁRIO:
      - Listar TODAS as dependências diretas do build file (pom.xml/build.gradle/etc.)
      - Identificar dependências transitivas críticas (as trazidas por frameworks)
      - Mapear escopo de cada dependência (compile, runtime, test, provided)

   2. VULNERABILIDADES (CVEs):
      - Verificar dependências com vulnerabilidades conhecidas (CVE/CVSS ≥ 7.0 = crítico)
      - Para cada CVE: ID, severidade, componente afetado, versão fixa, exploitability
      - Priorizar: RCE > auth bypass > information disclosure > DoS
      - Indicar se há exploit público conhecido

   3. LICENÇAS:
      - Identificar licenças de TODAS as dependências
      - Sinalizar licenças problemáticas: GPL viral, AGPL, EUPL, SSPL, BUSL, Commons Clause
      - Verificar compatibilidade com o modelo de negócio (SaaS B2B)
      - Dependências sem licença declarada → risco legal

   4. VERSÕES DESATUALIZADAS:
      - Identificar pacotes com major version lag (>1 major atrás)
      - Identificar pacotes com minor version lag (>3 minors atrás)
      - Pacotes sem update nos últimos 12 meses (abandonados?)
      - Sugerir upgrade path: current → recommended (com breaking changes list)

   5. SUPPLY CHAIN:
      - Dependências de maintainers únicos (bus factor = 1)
      - Dependências com histórico de compromise (ex: event-stream, left-pad)
      - Dependências de repositórios com poucas estrelas/contribuidores
      - Verificar se há dependências de forks não-oficiais

   6. FERRAMENTAS (se disponíveis):
      - Maven: mvn dependency:analyze (unused/undeclared), mvn versions:display-dependency-updates
      - Gradle: gradle dependencies, gradle dependencyUpdates
      - npm: npm audit --json
      - Python: pip-audit, safety check
      - Geral: OWASP Dependency-Check, Trivy, Snyk CLI

   Para CADA achado, indicar:
   - Componente + versão atual + versão recomendada
   - Severidade (🔴 Critical: CVE≥9 ou GPL em SaaS | 🟡 High: CVE≥7 ou major lag>2 | 🔵 Medium: minor lag ou licença ambígua)
   - Se é impeditivo para a Sprint {SPRINT_NUMBER}
   - Ação corretiva específica (ex: 'bump postgresql 42.7.1→42.7.5')"
```

### Passo 3 — Consolidar Achados no Documento Único

Consolidar TODOS os achados dos 4 agentes em um único documento:

```
ARQUIVO: {SPRINT_DIR}/IDENTIFIED-TECHNICAL-DEBT-{SPRINT_NAME}.md
```

O documento deve seguir esta estrutura:

```markdown
# IDENTIFIED-TECHNICAL-DEBT-{SPRINT_NAME}

- **Sprint alvo:** {SPRINT_NUMBER} de N — {SPRINT_NAME}
- **Data da análise:** YYYY-MM-DD
- **Skills executadas:** engineering-skills, caveman-review, superpowers:brainstorming, ponytail-review, code-review, codebase-cleanup-tech-debt, codebase-cleanup-deps-audit
- **Stack:** {STACK}
- **Total de achados:** X (🔴 Y críticos, 🟡 Z riscos, 🔵 W nits)
- **Impeditivos para iniciar a sprint:** S sim, N não

---

## Resumo Executivo

[Parágrafo de 3-5 linhas sumarizando os achados mais graves e a recomendação geral
sobre iniciar ou não a sprint sem correções prévias.]

---

## Matriz de Débitos Técnicos

> **Legenda das colunas:**
> - **ID:** DT-XXX (Débito Técnico, numeração sequencial)
> - **Severidade:** 🔴 Crítico (bloqueante) | 🟡 Risco (deve ser tratado) | 🔵 Nit (desejável)
> - **Skill:** Qual skill identificou (CR=caveman-review, PONY=ponytail-review, ARCH=brainstorming, ENG=engineering-skills, CODE=code-review, DEBT=codebase-cleanup-tech-debt, DEPS=codebase-cleanup-deps-audit)
> - **Complexidade:** H (Alta, >4h) | M (Média, 1-4h) | L (Baixa, <1h)
> - **Bloqueante?:** SIM (impede o início/incremento da sprint) | NÃO (pode ser tratado depois)
> - **Efeito se não tratado:** O que acontece se este débito for ignorado

| ID | Arquivo/Artefato | Achado | Severidade | Skill | Complexidade | Bloqueante? | Efeito se não tratado |
|:---|:---|:---|:---:|:---:|:---:|:---:|:---|
| DT-001 | `Arquivo.java:L42` | [Descrição concisa do problema] | 🔴 | PONY | M | SIM | [Consequência concreta] |
| DT-002 | ... | ... | ... | ... | ... | ... | ... |

---

## Achados por Skill

### engineering-skills (N achados)

[Agrupados por categoria: Arquitetura, Segurança, Configuração, Testes, etc.]

### caveman-review (N achados)

[Agrupados por artefato: SPRINT-CARD, SPRINT-TEST-SUITE, código, etc.]

### superpowers:brainstorming (N achados)

[Agrupados por decisão arquitetural questionada]

### ponytail-review (N achados)

[Agrupados por tag: delete, stdlib, native, yagni, shrink]
Terminar com: `net: -N lines possible.`

### code-review (N achados)

[Agrupados por severidade CodeRabbit: Critical, Warning, Info]
Para cada Critical/Warning: indicar se foi verificado manualmente (não é falso positivo).
Se CodeRabbit CLI indisponível: indicar "Revisão manual — CodeRabbit CLI não disponível."

### codebase-cleanup-tech-debt (N achados)

[Agrupados por categoria: Code Debt, Architecture Debt, Testing Debt, Documentation Debt, Infrastructure Debt]
Para cada achado: incluir métrica quantificada (ex: complexidade=14, 85 linhas duplicadas em 3 arquivos),
impacto em velocidade (horas/mês), e ROI estimado da correção.

### codebase-cleanup-deps-audit (N achados)

[Agrupados por categoria: Vulnerabilidades, Licenças, Versões Desatualizadas, Supply Chain]
Para cada CVE: ID, CVSS score, componente, versão fixa.
Para cada licença problemática: tipo de licença, por que é problemática para este projeto.
Resumo: N dependências auditadas, X vulnerabilidades (Y críticas), Z licenças problemáticas.

---

## Plano de Ação Recomendado

### Correções Pré-Sprint (impeditivos — devem ser feitos ANTES de iniciar)

| ID (TASKS.md) | Débito | Ação | Estimativa | Responsável |
|:---|:---|:---|:---:|:---|
| T-087.DT-001 | DT-001 | [O que fazer] | Xh | A definir |

### Correções Durante a Sprint (não-bloqueantes — podem ser incluídas no backlog)

| ID (TASKS.md) | Débito | Ação | Estimativa | Sprint sugerida |
|:---|:---|:---|:---:|:---|
| T-099.DT-017 | DT-017 | [O que fazer] | Xh | Sprint {SPRINT_NUMBER} |

### Correções Pós-Sprint (débito de longo prazo — sprints futuras)

| ID (TASKS.md) | Débito | Ação | Estimativa | Sprint sugerida |
|:---|:---|:---|:---:|:---|
| T-150.DT-020 | DT-020 | [O que fazer] | Xh | Sprint N+1 |

---

## Decisão do Time

> **Esta seção deve ser preenchida APÓS a revisão do time.**
> O agente deve apresentar o documento ao time e fazer a pergunta explícita:
> "Quais itens da lista acima serão tratados na Sprint {SPRINT_NUMBER}?
>  O que pode ir para as demais sprints?"

| ID | Decisão | Sprint alvo | Justificativa |
|:---|:---|:---:|:---|
| DT-XXX | Tratar agora | {SPRINT_NUMBER} | [Por que] |
| DT-YYY | Postergar | Sprint N+1 | [Por que] |

---

## Análise de Impacto nos Documentos

> **Esta seção deve ser preenchida APÓS a decisão do time,**
> para cada débito que foi aceito para correção na sprint atual.

### Impacto nos Documentos de Negócio

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-XXX | BUSINESS-REQUIREMENTS.md | Nova restrição técnica vira BR | Adicionar BR-XX |
| DT-XXX | 04-FEATURES.md | Feature existente ganha nova RN | Atualizar RN XX-YY |
| — | — | Sem impacto | — |

### Impacto nos Documentos-Mestre

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-XXX | ARCHITECTURE.md | Novo ADR ou revisão de ADR existente | Adicionar/atualizar ADR |
| DT-XXX | SPECS.md | Nova RN ou endpoint | Adicionar RN/corrigir contagem |
| DT-XXX | TASKS.md | Nova tarefa ou revisão de estimativa | Adicionar task/corrigir estimativa |
| DT-XXX | TEST_PLAN.md | Novos cenários de teste | Adicionar cenários |
| DT-XXX | PRD.md | Atualização de escopo ou stack | Atualizar seção X |

### Impacto nos Artefatos da Sprint

| ID | Documento | Impacto | Ação necessária |
|:---|:---|:---|:---|
| DT-XXX | SPRINT-CARD.md | Nova task ou métrica alterada | Adicionar linha na tabela de tasks |
| DT-XXX | SPRINT-TEST-SUITE.md | Novos cenários de teste | Adicionar cenários |
| DT-XXX | SPRINT-REVIEW.md | Novo item de demonstração | Adicionar checklist |

---

🤖 *Análise gerada em YYYY-MM-DD. {N} achados consolidados a partir de 7 skills. Documento base para decisão do time.*
```

### Passo 4 — Apresentar ao Time para Decisão

Após gerar o documento `IDENTIFIED-TECHNICAL-DEBT-{SPRINT_NAME}.md`:

1. **Apresentar o resumo executivo** — destacar quantos itens são impeditivos e por quê
2. **Fazer a pergunta explícita:**
   > "Quais itens da lista acima serão tratados na Sprint {SPRINT_NUMBER}? O que pode ir para as demais sprints?"
3. **Aguardar resposta do time** — NUNCA prosseguir sem decisão explícita
4. **Registrar as decisões** na seção "Decisão do Time" do documento

### Passo 5 — Propagar Impacto nos Documentos (APÓS decisão do time)

#### 5.0 Regras de Numeração de Tasks (APLICAR ANTES DE TUDO)

> **ESTA É A REGRA MAIS IMPORTANTE DO PASSO 5. NUNCA VIOLAR.**

Quando débitos técnicos são promovidos a tarefas de sprint, siga estas 3 regras:

**Regra 1 — Numeração sempre ao final (APPEND only):** Toda nova tarefa/task criada para atender um débito técnico recebe um ID sequencial **sempre a partir do último ID existente** no `TASKS.md`. Nunca inserir tasks de débito no meio da numeração existente — isso evita renumeração em cascata de tasks já consolidadas.

```
Algoritmo:
  1. Ler o TASKS.md e encontrar o MAIOR T-XXX já existente
  2. A primeira nova task de débito será T-{(maior) + 1}
  3. As demais seguem sequencialmente: T-{maior+2}, T-{maior+3}, ...

Exemplo:
  Última task no TASKS.md = T-086
  → 12 débitos da Frente 0 viram: T-087, T-088, T-089, ..., T-098
  → 7 débitos da Frente 3 viram: T-099, T-100, T-101, ..., T-105
  → NENHUMA task existente (T-001 a T-086) é renumerada
```

**Regra 2 — Código DT embutido no ID da task:** O ID da task incorpora o código do débito técnico no formato `T-XXX.DT-YYY`. Isso garante rastreabilidade visual imediata — ao ver o ID, sabe-se instantaneamente qual débito originou a task, sem precisar abrir o documento.

```
Formato: T-{novo-sequencial}.DT-{código-do-débito}
Exemplos:
  T-087.DT-001   ← Task #87, atende o débito técnico DT-001
  T-088.DT-002   ← Task #88, atende o débito técnico DT-002
  T-099.DT-017   ← Task #99, atende o débito técnico DT-017
  T-100.DT-025   ← Task #100, atende o débito técnico DT-025
```

O título da task também repete o código DT para reforço:

```
Formato completo: "T-XXX.DT-YYY: DT-YYY — <descrição curta>"
Exemplo: "T-087.DT-001: DT-001 — Atualizar Spring Boot 3.5.1→3.5.14 + Jackson 2.21.4"
```

**Regra 3 — Códigos DT são IMUTÁVEIS (external tool integration):** Os códigos `DT-XXX` gerados no documento de débitos técnicos são **permanentes e imutáveis**. Uma vez atribuídos, NUNCA são renumerados ou alterados. Isso é crítico porque:

- Ferramentas externas (Jira, Trello, Linear, GitHub Issues) podem ter tickets vinculados a estes códigos
- O formato `T-087.DT-001` permite buscar por `DT-001` em qualquer ferramenta e encontrar todos os tickets relacionados
- Dashboards e relatórios de débito técnico referenciam `DT-XXX` como chave primária
- Alterar um código `DT-XXX` quebraria links, rastreabilidade e integrações
- Se um débito for movido de sprint, ele MANTÉM seu `DT-XXX` original — apenas ganha um novo `T-XXX` na sprint destino

```
Exemplo de imutabilidade:
  DT-020 foi identificado na Sprint 3 mas postergado para Sprint 5
  → O código DT-020 NUNCA muda
  → Na Sprint 5, ele será "T-150.DT-020: DT-020 — ..." com um novo T-150
  → O link DT-020 no Jira continua válido independente da sprint
```

**Resumo visual do formato:**

```
TASKS.md (antes):
  T-084 | Deploy staging (K8s)...
  T-085 | UAT com Product Owner...
  T-086 | Deploy produção (K8s)...

IDENTIFIED-TECHNICAL-DEBT:
  DT-001 | Spring Boot CVEs | 🔴 | DEPS | M | SIM | ...
  DT-002 | AuditAspect quebrado | 🔴 | ARCH | M | SIM | ...

TASKS.md (depois — APPEND no final):
  ...
  T-086 | Deploy produção (K8s)...               ← última task original
  T-087.DT-001 | DT-001 — Atualizar Spring Boot...  ← NOVA (Frente 0)
  T-088.DT-002 | DT-002 — Refatorar AuditAspect...  ← NOVA (Frente 0)
  ...
  T-098.DT-012 | DT-012 — Criar exceções...         ← NOVA (Frente 0)
  T-099.DT-017 | DT-017 — Decidir V004...            ← NOVA (Frente 3)
  ...
  T-105.DT-046 | DT-046 — Atualizar Testcontainers... ← NOVA (Frente 3)
```

> ⚠️ **IMPORTANTE:** Este formato torna a rastreabilidade **visual e documental**. Basta olhar o ID `T-087.DT-001` para saber: (a) é a 87ª task do projeto, (b) atende o débito DT-001. Para detalhes do débito, abrir `IDENTIFIED-TECHNICAL-DEBT-*.md` e buscar por `DT-001`. Nenhuma ferramenta externa quebra porque `DT-001` é imutável.

#### 5.1 Atualização de Documentos

Para cada débito que o time decidiu tratar na sprint atual:

```
1. Identificar a cadeia de impacto:
   Débito Técnico → O que muda no código? → O que muda na arquitetura?
   → O que muda nas regras de negócio? → O que muda nas tasks?
   → O que muda nos testes? → O que muda na documentação?

2. Aplicar as Regras de Numeração (§5.0) ANTES de editar qualquer documento:
   a. Encontrar o MAIOR T-XXX no TASKS.md (buscar por todas as tasks numeradas)
   b. Atribuir T-{maior+1}.DT-{YYY} para o primeiro débito (Regra 1 + Regra 2)
   c. Continuar sequencialmente: T-{maior+2}.DT-{YYY}, T-{maior+3}.DT-{YYY}, ...
   d. NUNCA alterar os códigos DT-XXX no documento de débitos (Regra 3)
   e. NUNCA inserir tasks no meio da numeração — sempre APPEND ao final

3. Atualizar documentos na ordem correta (de cima para baixo):
   a. Documentos de Negócio (BUSINESS-REQUIREMENTS, EPICS, FEATURES, USER-STORYS)
      — se o débito introduzir novas restrições ou regras de negócio
   b. ARCHITECTURE.md — se houver mudança arquitetural (ADR, padrão, diagrama)
   c. PRD.md — se houver mudança de escopo, stack ou entidades
   d. SPECS.md — se houver novas RNs, endpoints ou validações
   e. TASKS.md — se houver novas tarefas ou alteração de estimativas
   f. TEST_PLAN.md — se houver novos cenários de teste
   g. Artefatos da Sprint (SPRINT-CARD, SPRINT-TEST-SUITE, SPRINT-REVIEW,
      SPRINT-DEVELOPMENT-PLANNING) — adicionar as novas tasks/cenários

4. Para cada documento atualizado:
   - Incrementar a versão (ex: v1.7 → v1.8)
   - Adicionar entrada no changelog explicando o que mudou e por quê
   - Atualizar referências cruzadas nos artefatos da sprint
```

### Passo 6 — Verificação Final

```
1. Verificar que TODOS os documentos impactados foram atualizados
2. Verificar que as referências cruzadas entre docs-mestre e artefatos da sprint
   permanecem consistentes
3. Verificar que as novas tasks têm IDs únicos e não conflitam com tasks existentes
4. Executar build e testes para garantir que correções no código não quebraram nada
5. Reportar resumo final:
   - N débitos identificados
   - M débitos aceitos para correção na sprint atual
   - K documentos atualizados
   - BUILD: SUCCESS/FAILURE
```

---

## Regras de Ouro

1. **NUNCA decidir unilateralmente:** O agente cataloga e recomenda. O time decide. A seção "Decisão do Time" só é preenchida após resposta explícita do humano.
2. **NUNCA alterar docs-mestre sem registrar:** Toda alteração em documento-mestre gera bump de versão + changelog.
3. **SEMPRE propagar impacto de baixo para cima:** Negócio → Arquitetura → Specs → Tasks → Testes. Nunca o contrário.
4. **SEMPRE verificar build após correções no código:** `mvn test` (ou equivalente da stack) deve passar antes de considerar o passo concluído.
5. **O documento de débitos é vivo:** Se novos débitos forem descobertos durante a sprint, adicionar ao documento com a data de descoberta.
6. **Prioridade é segurança > correção > complexidade > dívida estrutural:** Débitos de segurança (SQL injection, CVEs, falta de validação) são sempre 🔴 e bloqueantes. Débitos de dependências com CVEs ≥9 são bloqueantes. Débitos de complexidade (YAGNI, shrink) e dívida estrutural de longo prazo raramente são bloqueantes.
7. **Código DT é imutável e permanente:** O código `DT-XXX` atribuído no documento `IDENTIFIED-TECHNICAL-DEBT-*.md` NUNCA é alterado. Ele serve como chave primária para ferramentas externas (Jira, Trello, GitHub Issues). Se um débito mudar de sprint, apenas seu `T-XXX` (task ID) muda — o `DT-XXX` permanece idêntico.
8. **Task de débito usa formato `T-XXX.DT-YYY`:** Toda task originada de débito técnico recebe um ID no formato `T-{último+1}.DT-{código}` (ex: `T-087.DT-001`, `T-088.DT-002`), garantindo rastreabilidade visual e documental imediata entre o backlog e o documento de débitos.
9. **Numeração sempre APPEND (nunca insert no meio):** Tasks de débito são SEMPRE adicionadas ao final do TASKS.md com um novo número sequencial. NUNCA inserir no meio da numeração existente — isso evita renumeração em cascata que quebraria referências em ferramentas externas.

---

## Exemplo de Uso

```
Humano: "Inicie o PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT para a Sprint 3"

Agente: "Vou executar a auditoria técnica multidisciplinar.
         Parâmetros necessários:
         - SOLUTION_PATH: ?
         - PROJECT_NAME: ?
         - SOLUTION_NAME: ?
         - SPRINT_NUMBER: 3
         - SPRINT_NAME: ?
         - STACK: ? (ou auto-detectar)"

Humano: "SOLUTION_PATH=/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin
         PROJECT_NAME=PRJ-FIN-2026-0003-SAAS-FBSO-ORG
         SOLUTION_NAME=ms-fbso-platform-admin
         SPRINT_NAME=sprint-03-portal-admin
         STACK=auto"

Agente: [Executa Passos 0-3, gera IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md]
        "Análise concluída. 47 débitos técnicos identificados
         (12 🔴 críticos, 22 🟡 riscos, 13 🔵 nits) por 7 skills.
         8 são impeditivos para iniciar a Sprint 3.
         Documento salvo em: .../sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md
         
         Quais itens da lista acima serão tratados na Sprint 3?
         O que pode ir para as demais sprints?"

Humano: [Decide quais tratar agora vs. depois]

Agente: [Registra decisões no documento, executa Passo 5 (propagação de impacto),
         atualiza docs-mestre e artefatos da sprint, executa Passo 6 (verificação)]
```

---

🤖 *Prompt v3.0. Integra 7 skills no fluxo Spec-Driven Development para identificação sistemática de débitos técnicos pré-sprint. v3.0 adiciona: Regras de Numeração (§5.0) — código DT imutável, task sequencial contínua, rastreabilidade DT-XXX↔T-XXX para integração com ferramentas externas.*
