# PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT

## Contexto

Este prompt orquestra a **identificação e catalogação de débitos técnicos** antes do início de uma nova sprint, como parte do fluxo Spec-Driven Development. Ele atua como uma **auditoria técnica multidisciplinar** que combina **9 skills complementares** para varrer documentação, código, arquitetura, complexidade, duplicação, atalhos intencionais, qualidade, dívida técnica estrutural e segurança de dependências, produzindo um documento único e acionável de débitos técnicos.

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

Este prompt orquestra **9 skills** em paralelo, cada uma cobrindo um ângulo diferente. Cada skill produz seus achados no seu formato nativo, que são depois consolidados no documento único.

| # | Skill | Ângulo | O que encontra |
|:---:|:---|:---|:---|
| 1 | `code-reviewer` | Qualidade de Código | Violações SOLID, code smells, más práticas de arquitetura, padrões incorretos — regras universais (segurança, async, recursos, exceções, performance) + regras específicas por linguagem (14 linguagens suportadas) |
| 2 | `caveman-review` | Consistência | Inconsistências entre documentação e código, métricas divergentes entre artefatos, versões desatualizadas, referências cruzadas quebradas |
| 3 | `superpowers:brainstorming` | Decisões | Decisões arquiteturais questionáveis, trade-offs não documentados, premissas inválidas, gaps no planejamento que afetam a sprint atual |
| 4 | `ponytail-review` | Complexidade | Código morto, abstrações desnecessárias (YAGNI), dependências que podem ser substituídas por stdlib/nativas, lógica que pode ser simplificada |
| 5 | `ponytail-debt` | Atalhos Intencionais | Coleta todos os comentários `ponytail:` do código em um ledger de dívida, sinalizando marcadores sem trigger de upgrade como risco de apodrecimento |
| 6 | `code-review` | Bugs & Vulnerabilidades | Bugs, vulnerabilidades de segurança, anti-padrões, problemas de performance via CodeRabbit CLI — revisão automatizada com severidade (Critical/Warning/Info) |
| 7 | `jscpd` + `dry-refactoring` | Duplicação | Código duplicado (copy-paste) detectado via jscpd em 220+ linguagens com métricas de % duplicação; `dry-refactoring` propõe estratégias de eliminação (extract function/module/constant/base class) |
| 8 | `tech-debt` | Dívida Estrutural | Categorização sistemática (Code, Architecture, Test, Dependency, Documentation, Infrastructure) com framework de priorização: `(Impact + Risk) × (6 − Effort)`, plano de remediação faseado |
| 9 | `security-review` | Segurança & Dependências | Data-flow tracing entre arquivos, CVEs em dependências, licenças incompatíveis, pacotes desatualizados, secrets expostos, vulnerabilidades de injeção — 8 linguagens suportadas |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 6 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

Se `{STACK}` não for informado, auto-detectar a partir de `PRD.md` e `ARCHITECTURE.md`.

### Passo 1 — Carregar Contexto Completo

```
1. Ler TODOS os documentos-mestre (PRD, SPECS, TASKS, TEST_PLAN, ARCHITECTURE)
2. Ler TODOS os artefatos da sprint atual (SPRINT-CARD, SPRINT-TEST-SUITE, SPRINT-REVIEW, etc.)
3. **Buscar TODOS os IDENTIFIED-TECHNICAL-DEBT-*.md de sprints anteriores:**
   - Localizar em `{SPRINT_DIR}/../sprint-*/IDENTIFIED-TECHNICAL-DEBT-*.md`
   - Extrair débitos com status diferente de "Concluído"/"Resolvido"
   - Montar o Backlog de Débitos Técnicos (seção §Backlog)
   - Estes débitos NÃO ganham novo DT-XXX — mantêm o código original
4. Escanear a estrutura completa do código-fonte:
   - Listar todos os pacotes e classes
   - Listar todas as migrations Flyway
   - Listar todos os testes e sua cobertura
   - Verificar configurações (application.yml, pom.xml/build.gradle)
5. Se existir DOCS-SPRINT-CAVEMAN-REVIEW.md de sprint anterior, carregar como baseline
```

### Passo 2 — Executar as 9 Skills em Paralelo

Cada skill é invocada como um agente independente, com seu próprio escopo. Os agentes 1-5 são baseados em skills locais (zero dependências externas). Os agentes 6-9 requerem verificação de pré-requisitos.

**Verificação de pré-requisitos (Agentes 6-9):**

```
Agente 6 (code-review): Verificar se CodeRabbit CLI está instalado e autenticado.
  → coderabbit --version && coderabbit auth status
  → Se NÃO disponível: executar mesmo assim com revisão manual de código
    (análise estática sem CLI, usando padrões conhecidos de bugs/security)

Agente 7 (jscpd + dry-refactoring): Verificar se Node.js está disponível.
  → node --version && npx jscpd --version
  → Se NÃO disponível: executar mesmo assim com grep manual de duplicação
    (buscar blocos idênticos >10 linhas em arquivos do mesmo pacote)

Agente 8 (tech-debt): Sem dependências externas.
  → Análise puramente baseada em leitura de código e métricas estáticas.

Agente 9 (security-review): Sem dependências externas obrigatórias.
  → Análise de dados de build files (pom.xml, build.gradle, package.json, etc.).
  → Se disponível, usar ferramentas nativas: mvn dependency:analyze, npm audit, etc.
```

#### Agente 1: code-reviewer

```
Escopo: Código-fonte completo (src/main/) + Configurações
Prompt do agente:
  "Execute uma revisão de qualidade de código completa em {SOLUTION_PATH}/src/
   usando as regras do code-reviewer (rules/universal.md + language-specific).

   Carregue nesta ordem:
   1. rules/universal.md — segurança, async, recursos, exceções, performance (sempre)
   2. languages/{lang}.md — regras específicas da linguagem do projeto

   Identifique:
   1. VIOLAÇÕES SOLID:
      - Single Responsibility: classes/funções com múltiplas responsabilidades
      - Open/Closed: modificações em classes estáveis em vez de extensão
      - Liskov Substitution: subtipos que quebram contratos
      - Interface Segregation: interfaces gordas com métodos não usados
      - Dependency Inversion: dependências de concreto em vez de abstração

   2. CODE SMELLS:
      - Métodos longos (>30 linhas), god classes (>500 linhas, >20 métodos)
      - Parâmetros excessivos (>5), flags booleanas como parâmetro
      - Código comentado, nomes não descritivos, números mágicos
      - Mutable state desnecessário, null handling ausente

   3. PADRÕES DE ARQUITETURA:
      - Violações de camadas (ex: controller acessando repository direto sem service)
      - Convenções do projeto não seguidas (ex: BaseRepository não usado)
      - Configurações incorretas ou duplicadas em application.yml/properties

   4. SEGURANÇA (universal rules):
      - SQL injection, XSS, path traversal, command injection
      - Insecure deserialization, weak cryptography
      - Secrets hardcoded, variáveis de ambiente não validadas

   5. TESTES:
      - Cobertura insuficiente em caminhos críticos (@Transactional, @Auditable)
      - Testes sem asserts reais, testes que dormem (Thread.sleep)
      - Falta de testes de integração para repositories e serviços

   Para cada achado, indique: localização exata (arquivo:linha), severidade
   (🔴 Critical / 🟡 Warning / 🔵 Info), regra violada (ex: SOLID-SRP, SMELL-long-method),
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

#### Agente 5: ponytail-debt

```
Escopo: Comentários `ponytail:` no código-fonte
Prompt do agente:
  "Execute um ponytail-debt scan no código em {SOLUTION_PATH}/src/.

   1. Grep por todos os comentários marcados com 'ponytail:':
      grep -rnE '(#|//|--) ?ponytail:' {SOLUTION_PATH}/src/ \
        --exclude-dir=node_modules --exclude-dir=.git --exclude-dir=target

   2. Para cada marcador encontrado, extraia:
      - Arquivo:linha
      - O que foi simplificado (contexto do comentário)
      - Ceiling (limite nomeado após 'ponytail:')
      - Trigger de upgrade (quando revisitar)

   3. Classifique cada marcador:
      - 🟢 has-trigger: tem trigger de upgrade definido
      - 🔴 no-trigger: NÃO tem trigger — risco de apodrecimento (vira 'later means never')

   4. Opcional: use git blame -L<linha>,<linha> para identificar o autor
      de cada marcador (para follow-up com o time)

   Output final:
   - Arquivo agrupado por diretório
   - '<N> markers, <M> with no trigger.'
   - Se nada encontrado: 'No ponytail: debt. Clean ledger.'"
```

#### Agente 6: code-review (CodeRabbit)

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

#### Agente 7: jscpd + dry-refactoring

```
Escopo: Duplicação de código (copy-paste) em todo o source
Pré-requisito: Node.js (node --version). Se indisponível, fallback para grep manual.
Prompt do agente:
  "Execute uma análise de duplicação de código em {SOLUTION_PATH}/src/.

   1. DETECÇÃO (jscpd):
      Se Node.js disponível:
        npx jscpd --reporters ai --min-lines 10 --min-tokens 50 \
          --ignore '**/node_modules/**,**/target/**,**/dist/**' \
          {SOLUTION_PATH}/src/
      → Output: lista de clones (arquivo:linhas ~ arquivo:linhas) + % duplicação

      Se Node.js NÃO disponível (fallback manual):
        - Buscar blocos idênticos >10 linhas entre arquivos do mesmo pacote
        - Identificar lógica repetida em múltiplos serviços/handlers
        - Verificar constantes e configurações duplicadas

   2. CATEGORIZAÇÃO (dry-refactoring — escolher estratégia para cada clone):
      | Estratégia | Quando usar | Exemplo |
      |-----------|-------------|---------|
      | Extract function | Bloco de lógica repetido | Função de validação duplicada em 3 services |
      | Extract module/utility | Lógica compartilhada entre domínios diferentes | DateUtils em múltiplos pacotes |
      | Extract constant/config | Dados/configuração repetidos | Valores default duplicados |
      | Template/base class | Estrutura de classe repetida | Repository custom methods idênticos |

   3. PRIORIZAÇÃO:
      - Clones com maior número de linhas primeiro (maior impacto)
      - Clones entre módulos não relacionados → sinalizam missing shared utility
      - Clones entre testes → sinalizam missing test helper
      - Use --min-lines 10 para filtrar ruído

   4. MÉTRICAS:
      - % total de duplicação no projeto
      - Top 5 arquivos mais duplicados
      - Número de clones por categoria de estratégia

   Output: lista de clones com estratégia sugerida, priorizados por impacto."
```

#### Agente 8: tech-debt

```
Escopo: Categorização e priorização de dívida técnica estrutural
Prompt do agente:
  "Execute uma análise completa de dívida técnica em {SOLUTION_PATH}/src/
   usando o framework de categorização e priorização do tech-debt.

   1. CATEGORIZAR cada achado:

   | Categoria | O que inclui | Indicadores |
   |-----------|-------------|-------------|
   | Code debt | Lógica duplicada, más abstrações, números mágicos | Bugs, desenvolvimento lento |
   | Architecture debt | Monólitos que deveriam ser split, data store errado | Limites de escala |
   | Test debt | Baixa cobertura, testes flaky, falta de integração | Regressões em produção |
   | Dependency debt | Bibliotecas desatualizadas, sem manutenção | Vulnerabilidades de segurança |
   | Documentation debt | Runbooks ausentes, READMEs desatualizados, tribal knowledge | Dor no onboarding |
   | Infrastructure debt | Deploys manuais, sem monitoring, sem IaC | Incidentes, recovery lento |

   2. PRIORIZAR cada item (framework quantitativo):
      - Impact: Quanto isso desacelera o time? (1-5)
      - Risk: O que acontece se não corrigirmos? (1-5)
      - Effort: Quão difícil é a correção? (1-5, invertido — menor esforço = maior prioridade)

      Priority = (Impact + Risk) × (6 − Effort)

   3. Para CADA achado, incluir:
      - Categoria + localização exata (arquivo:linha)
      - Métrica quantificada (ex: complexidade=14, 85 linhas duplicadas em 3 arquivos)
      - Priority score (2-50) + severidade (🔴 ≥40, 🟡 25-39, 🔵 <25)
      - Impacto em velocidade (horas/perda por mês)
      - ROI estimado da correção (horas investidas vs. horas economizadas)
      - Se é impeditivo para a Sprint {SPRINT_NUMBER}

   4. OUTPUT: plano de remediação faseado que pode ser executado junto com feature work:
      - Fase 1 (Imediato — antes da sprint): itens com priority ≥40 e bloqueantes
      - Fase 2 (Durante a sprint): itens com priority 25-39
      - Fase 3 (Sprints futuras): itens com priority <25

   Output agrupado por categoria, ordenado por priority score decrescente."
```

#### Agente 9: security-review

```
Escopo: Segurança de código, dependências, licenças e supply chain
Prompt do agente:
  "Execute uma auditoria de segurança completa em {SOLUTION_PATH}/.

   1. DEPENDENCY AUDIT (fast wins primeiro):
      - Varrer build files (pom.xml, build.gradle, package.json, Cargo.toml, go.mod)
      - Identificar dependências com CVEs conhecidas (CVSS ≥7.0 → crítico)
      - Sinalizar pacotes deprecated, sem manutenção (>12 meses sem update)
      - Sugerir upgrade path: versão atual → versão recomendada (com breaking changes)

   2. VULNERABILITY SCAN (data-flow tracing):
      Traçar fluxos de dados entre arquivos para detectar:
      - Injection: SQL, NoSQL, Command, LDAP, XPath — onde user input chega a sinks perigosos
      - XSS: output não sanitizado em respostas HTML/JSON
      - Path traversal: file paths construídos com input do usuário
      - Insecure deserialization: ObjectMapper sem TypeSafe, pickle.loads, yaml.load
      - SSRF: URLs controladas pelo usuário usadas em HTTP clients

   3. SECRETS & EXPOSURE:
      - API keys, tokens, senhas hardcoded no código
      - Credenciais em arquivos de configuração (application.yml, .env)
      - Chaves privadas, certificados, connection strings expostos
      - Variáveis de ambiente lidas sem validação

   4. AUTH & ACCESS CONTROL:
      - Endpoints públicos sem autenticação (quando deveriam ter)
      - Falta de autorização por role/scope
      - JWT sem validação de exp, secret fraco, alg=none
      - Password reset sem rate limiting

   5. LICENSES & SUPPLY CHAIN:
      - Licenças problemáticas: GPL viral, AGPL, EUPL, SSPL, BUSL, Commons Clause
      - Dependências de maintainer único (bus factor = 1)
      - Dependências com histórico de compromise (event-stream, left-pad, etc.)
      - Forks não-oficiais, repositórios com <100 stars

   6. FERRAMENTAS COMPLEMENTARES (se disponíveis):
      - Maven: mvn dependency:analyze (unused/undeclared)
      - Gradle: gradle dependencies, gradle dependencyUpdates
      - npm: npm audit --json
      - Python: pip-audit, safety check
      - Geral: OWASP Dependency-Check, Trivy, Snyk CLI

   Para CADA achado, indicar:
   - Componente + versão atual + versão recomendada (se aplicável)
   - CVE ID, CVSS score, exploitability (se CVE)
   - Severidade: 🔴 Critical (CVE≥9 ou GPL em SaaS) | 🟡 High (CVE≥7 ou major lag>2) | 🔵 Medium (minor lag ou licença ambígua)
   - Se é impeditivo para a Sprint {SPRINT_NUMBER}
   - Ação corretiva específica (ex: 'bump postgresql 42.7.1→42.7.5')
   - Self-verify: reconfirmar cada finding para eliminar falsos positivos"
```

### Passo 3 — Consolidar Achados no Documento Único

Consolidar TODOS os achados dos 9 agentes + backlog de sprints anteriores em um único documento:

```
ARQUIVO: {SPRINT_DIR}/IDENTIFIED-TECHNICAL-DEBT-{SPRINT_NAME}.md
```

O documento deve seguir esta estrutura:

```markdown
# IDENTIFIED-TECHNICAL-DEBT-{SPRINT_NAME}

- **Sprint alvo:** {SPRINT_NUMBER} de N — {SPRINT_NAME}
- **Data da análise:** YYYY-MM-DD
- **Skills executadas:** code-reviewer, caveman-review, superpowers:brainstorming, ponytail-review, ponytail-debt, code-review (CodeRabbit), jscpd+dry-refactoring, tech-debt, security-review
- **Stack:** {STACK}
- **Total de achados:** X (🔴 Y críticos, 🟡 Z riscos, 🔵 W nits)
- **Impeditivos para iniciar a sprint:** S sim, N não

---

## Resumo Executivo

[Parágrafo de 3-5 linhas sumarizando os achados mais graves e a recomendação geral
sobre iniciar ou não a sprint sem correções prévias.]

---

## Backlog de Débitos Técnicos (Sprints Anteriores)

> **Instrução:** Antes de catalogar os novos débitos, o agente deve **obrigatoriamente**:
> 1. Buscar TODOS os arquivos `IDENTIFIED-TECHNICAL-DEBT-*.md` de sprints anteriores em `{SPRINT_DIR}/../`.
> 2. Para cada arquivo encontrado, extrair os débitos cujo status **não seja "Concluído"** ou "Resolvido".
> 3. Consolidar abaixo como backlog ativo, com referência ao documento original.

Débitos técnicos identificados em sprints anteriores que **permanecem não resolvidos** e são candidatos a tratamento na sprint atual ou futuras.

| DT-XXX | Sprint Origem | Descrição | Severidade | Bloqueante? | Status | Resolução (do doc original ou revisada) |
|:---|:---|:---|:---:|:---:|:---|:---|
| DT-001 | Sprint 2 | [Descrição copiada do documento original] | 🔴 | SIM | Pendente | [Resumo da resolução original ou **revisado**: nova resolução se contexto mudou] |
| DT-002 | Sprint 2 | [Descrição copiada do documento original] | 🟡 | NÃO | Pendente | [Resumo da resolução] |
| DT-005 | Sprint 1 | [Descrição] | 🔵 | NÃO | Pendente | [Resumo — ou `↗ ver DT-005 em IDENTIFIED-TECHNICAL-DEBT-sprint-01-*.md` se detalhamento for extenso] |

> **Regra de referência:** Se o detalhamento completo do débito for extenso (>5 linhas), fazer referência ao documento original em vez de copiar: `↗ ver DT-XXX em {arquivo-original}`. Isso mantém o documento enxuto. Apenas revise o resumo de resolução se o contexto técnico tiver mudado desde a sprint original.

**Total em backlog:** N débitos pendentes de sprints anteriores.

---

---

## Matriz de Débitos Técnicos

> **Esta matriz consolida TODOS os débitos — tanto os novos (descobertos nesta sprint) quanto os do backlog (sprints anteriores).**
>
> **Legenda das colunas:**
> - **ID:** DT-XXX (Débito Técnico, numeração sequencial e IMUTÁVEL)
> - **Sprint Origem:** Em qual sprint o débito foi identificado (Sprint 1, 2, 3, ...). Débitos do backlog mantêm o DT-XXX original.
> - **Severidade:** 🔴 Crítico (bloqueante) | 🟡 Risco (deve ser tratado) | 🔵 Nit (desejável)
> - **Skill:** Qual skill identificou (CR=caveman-review, PONY=ponytail-review, PDBT=ponytail-debt, ARCH=brainstorming, CREV=code-reviewer, CODE=code-review, JSCPD=jscpd+dry-refactoring, DEBT=tech-debt, SEC=security-review). Débitos de backlog: `BACKLOG`.
> - **Complexidade:** H (Alta, >4h) | M (Média, 1-4h) | L (Baixa, <1h)
> - **Bloqueante?:** SIM (impede o início/incremento da sprint) | NÃO (pode ser tratado depois)
> - **Efeito se não tratado:** O que acontece se este débito for ignorado

| ID | Sprint Origem | Arquivo/Artefato | Achado | Severidade | Skill | Complexidade | Bloqueante? | Efeito se não tratado |
|:---|:---|:---|:---|:---:|:---:|:---:|:---:|:---|
| DT-001 | Sprint 2 | `Arquivo.java:L42` | [Descrição concisa — backlog] | 🔴 | BACKLOG | M | SIM | [Consequência concreta] |
| DT-050 | Sprint 3 | `Service.java:L128` | [Descrição concisa — novo débito] | 🟡 | CREV | H | NÃO | [Consequência concreta] |
| DT-051 | Sprint 3 | ... | ... | ... | ... | ... | ... | ... |

> **Ordenação:** Débitos bloqueantes primeiro (🔴), depois por sprint origem (mais antigos primeiro — risco de apodrecimento), depois por severidade.

---

## Achados por Skill

### code-reviewer (N achados)

[Agrupados por categoria: SOLID Violations, Code Smells, Architecture Patterns, Security, Tests]
Para cada achado: regra violada (ex: SOLID-SRP, SMELL-long-method, SEC-injection).

### caveman-review (N achados)

[Agrupados por artefato: SPRINT-CARD, SPRINT-TEST-SUITE, código, etc.]

### superpowers:brainstorming (N achados)

[Agrupados por decisão arquitetural questionada]

### ponytail-review (N achados)

[Agrupados por tag: delete, stdlib, native, yagni, shrink]
Terminar com: `net: -N lines possible.`

### ponytail-debt (N achados)

[Agrupados por arquivo. Cada marcador: ceiling, trigger de upgrade, risco de apodrecimento.]
Terminar com: `<N> markers, <M> with no trigger.`

### code-review (N achados)

[Agrupados por severidade CodeRabbit: Critical, Warning, Info]
Para cada Critical/Warning: indicar se foi verificado manualmente (não é falso positivo).
Se CodeRabbit CLI indisponível: indicar "Revisão manual — CodeRabbit CLI não disponível."

### jscpd + dry-refactoring (N achados)

[Agrupados por estratégia: extract function, extract module, extract constant, template/base class]
Para cada clone: % duplicação, estratégia sugerida, estimativa de esforço.
Resumo: % total de duplicação do projeto, Top 5 arquivos mais duplicados.

### tech-debt (N achados)

[Agrupados por categoria: Code Debt, Architecture Debt, Test Debt, Dependency Debt, Documentation Debt, Infrastructure Debt]
Para cada achado: priority score (2-50), métrica quantificada, impacto em velocidade (horas/mês), e ROI estimado da correção.
Incluir plano de remediação faseado: Fase 1 (≥40), Fase 2 (25-39), Fase 3 (<25).

### security-review (N achados)

[Agrupados por categoria: Dependency CVEs, Injection Vulnerabilities, Secrets & Exposure, Auth & Access Control, Licenses & Supply Chain]
Para cada CVE: ID, CVSS score, componente, versão fixa, exploitability.
Para cada licença problemática: tipo de licença, por que é problemática para este projeto.
Resumo: N dependências auditadas, X vulnerabilidades (Y críticas), Z licenças problemáticas.

---

## Recomendações Prioritárias

> **Esta seção lista débitos que DEVEM ou DEVERIAM ser tratados na sprint atual.**
> Inclui tanto débitos recém-descobertos (Agentes 1-9) quanto débitos do backlog (sprints anteriores) que se tornaram críticos para a sprint atual.

### 🔴 Bloqueantes (impeditivos — devem ser corrigidos ANTES de iniciar a sprint)

Débitos que **impedem** o início ou o avanço da Sprint {SPRINT_NUMBER}. Sem correção, o desenvolvimento para ou produz código quebrado.

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa | Responsável |
|:---|:---|:---|:---|:---:|:---|
| T-087.DT-001 | DT-001 | Sprint 2 | [O que fazer — backlog revisado] | Xh | A definir |
| T-088.DT-050 | DT-050 | Sprint 3 | [O que fazer — novo débito] | Xh | A definir |

### 🟡 Recomendados (devem ser tratados — podem ser incluídos no backlog da sprint)

Débitos que **não bloqueiam** o início da sprint mas que, se ignorados, acumulam risco técnico significativo.

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa | Sprint sugerida |
|:---|:---|:---|:---|:---:|:---|
| T-099.DT-002 | DT-002 | Sprint 2 | [O que fazer — backlog revisado] | Xh | Sprint {SPRINT_NUMBER} |
| T-100.DT-051 | DT-051 | Sprint 3 | [O que fazer — novo débito] | Xh | Sprint {SPRINT_NUMBER} |

### 🔵 Desejáveis (nice-to-have — se houver capacidade)

| ID (TASKS.md) | DT-XXX | Sprint Origem | Ação Corretiva | Estimativa |
|:---|:---|:---|:---|:---:|
| T-105.DT-005 | DT-005 | Sprint 1 | [O que fazer — backlog antigo] | Xh |

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

## Débitos Técnicos Elegíveis para Sprints Futuras

> **Esta seção lista débitos que NÃO serão tratados na sprint atual, mas permanecem no radar para sprints futuras.**
> Inclui tanto débitos do backlog antigo quanto débitos recém-descobertos que o time decidiu postergar.
> A classificação segue o mesmo formato da Matriz de Débitos Técnicos.

Débitos que o time decidiu **explicitamente postergar** — seja por baixa severidade, alta complexidade, ou dependência de outros fatores. Estes itens devem ser reavaliados no próximo ciclo de `PROMPT-GENERATE-IDENTIFIED-TECHNICAL-DEBT`.

| DT-XXX | Sprint Origem | Descrição | Severidade | Bloqueante? | Skill | Complexidade | Sprint Sugerida | Justificativa do Adiamento |
|:---|:---|:---|:---:|:---:|:---:|:---:|:---|:---|
| DT-005 | Sprint 1 | [Descrição] | 🔵 | NÃO | BACKLOG | L | Sprint 5 | Baixo impacto — dívida cosmética |
| DT-052 | Sprint 3 | [Descrição — novo débito postergado] | 🟡 | NÃO | DEBT | H | Sprint 4 | Alta complexidade — requer refactor prévio do módulo X |
| DT-053 | Sprint 3 | [Descrição] | 🔵 | NÃO | PONY | M | Sprint N+2 | Depende da entrega do EP-03 |

> **Regra de reavaliação:** Na próxima execução deste prompt (próxima sprint), os débitos desta seção devem ser relidos e reclassificados — um débito que era 🟡 pode se tornar 🔴 se o contexto mudou.

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

🤖 *Análise gerada em YYYY-MM-DD. {N} achados consolidados a partir de 9 skills. Documento base para decisão do time.*
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
         (12 🔴 críticos, 22 🟡 riscos, 13 🔵 nits) por 9 skills.
         8 são impeditivos para iniciar a Sprint 3.
         Documento salvo em: .../sprint-03-portal-admin/IDENTIFIED-TECHNICAL-DEBT-sprint-03-portal-admin.md
         
         Quais itens da lista acima serão tratados na Sprint 3?
         O que pode ir para as demais sprints?"

Humano: [Decide quais tratar agora vs. depois]

Agente: [Registra decisões no documento, executa Passo 5 (propagação de impacto),
         atualiza docs-mestre e artefatos da sprint, executa Passo 6 (verificação)]
```

---

🤖 *Prompt v3.1. Integra 9 skills no fluxo Spec-Driven Development para identificação sistemática de débitos técnicos pré-sprint. v3.1 adiciona: code-reviewer (SOLID + 14 linguagens), ponytail-debt (ledger de atalhos), jscpd+dry-refactoring (duplicação), tech-debt (priorização quantitativa), security-review (data-flow tracing) — substituindo 2 skills inexistentes e 1 meta-índice. Regras de Numeração (§5.0): código DT imutável, task sequencial contínua, rastreabilidade DT-XXX↔T-XXX.*
