# Análise de Skills: caveman, ponytail & golang-pro

> **Data:** 2026-07-07
> **Contexto:** Análise comparativa e guia de uso combinado para desenvolvimento de microserviços Go

---

## 1. Comparativo: caveman (JuliusBrussee) vs caveman (mattpocock)

### Metadados

| Dimensão | **JuliusBrussee/caveman** | **mattpocock/skills** |
|---|---|---|
| **1ª aparição** | 4 Abr 2026 | 17 Abr 2026 (13 dias depois) |
| **Installs** | 322.2K | 217.9K |
| **Stars (repo)** | 85.5K | 157.9K (repo inteiro de skills) |
| **Status no repo** | ✅ Ativo e mantido | ❌ **Removido** |
| **SKILL.md** | 77 linhas, completo | **Não existe mais** |
| **Níveis de intensidade** | 6 (lite, full, ultra + 3 wenyan) | ~4 (simplificado) |

### O changelog do Matt Pocock

Commit [`47bde84`](https://github.com/mattpocock/skills/commit/47bde84da032afb2e5058f997f3bbca47d321dbd):

> **Remove the `caveman` and `zoom-out` skills.**
> `caveman` was a duplicate of another skill I was testing and was never meant to be public.

O próprio Matt Pocock afirma que a skill dele era uma duplicata que nunca deveria ter sido pública.

### Diferenças técnicas (quando a skill do Matt ainda existia)

| Regra | Julius | Matt |
|---|---|---|
| **Setas de causalidade (→)** | 🚫 Proíbe — "No causal arrows (→), own token, save nothing" | ✅ Usa — "arrow notation (X → Y) for causality" |
| **Abreviações inventadas** | 🚫 Proíbe (cfg/impl/req/res/fn) | Não menciona |
| **Variantes Wenyan (文言文)** | ✅ 3 níveis (lite/full/ultra) | ❌ Ausente |
| **Níveis de intensidade** | 6 | ~4 |
| **Persistência entre respostas** | ✅ Documentado | ✅ Documentado |
| **Auto-clareza para segurança** | ✅ Documentado | ✅ Documentado |
| **Preservação de idioma do usuário** | ✅ Explícito | Não documentado |
| **Narração de tool calls** | 🚫 Proíbe | Não menciona |

### Veredito

Usar qualquer uma **NÃO** terá o mesmo resultado. Use **JuliusBrussee/caveman**:

```bash
npx skills add https://github.com/juliusbrussee/caveman --skill caveman -g
```

---

## 2. Comparativo: caveman vs ponytail

### Elas fazem a mesma função?

**Não.** Operam em dimensões completamente diferentes:

| Dimensão | `caveman` | `ponytail` |
|---|---|---|
| **O que controla** | **COMO** o agente se comunica | **O QUE** o agente constrói |
| **Domínio** | Estilo de prosa/respostas | Escopo de soluções de código |
| **Problema que ataca** | Token usage excessivo em respostas | Código desnecessário/over-engineering |
| **Métrica** | -65% tokens de output | -54% linhas de código |
| **Mecanismo** | Regras de compressão linguística | Escada de decisão de 7 degraus |
| **Níveis** | 6 (lite, full, ultra + 3 wenyan) | 3 (lite, full, ultra) |
| **Installs** | 322.2K | 35.3K total (10.9K skill principal) |
| **Gatilhos** | "caveman mode", "be brief", "less tokens" | "ponytail", "be lazy", "yagni", "do less" |

### Exemplo concreto da diferença

**Tarefa:** "Adiciona cache pras respostas da API"

| | `caveman` (age no texto) | `ponytail` (age no código) |
|---|---|---|
| **Resultado** | `@lru_cache on fetch. Skip custom cache class.` | `@lru_cache(maxsize=1000)` no fetch — sem classe de cache customizada |
| **Sem a skill** | "Sure! I'd be happy to help you add caching..." | Classe `CacheManager` de 50 linhas com TTL, invalidação, interface abstrata... |

### Elas se complementam?

**Sim — e isso é intencional.** O próprio SKILL.md do `ponytail` declara na seção **Boundaries**:

> *"Ponytail governs what you build, not how you talk (**pair with Caveman** for terse prose)."*

```
┌─────────────────────────────────────────────┐
│               TAREFA DO USUÁRIO             │
├─────────────────────┬───────────────────────┤
│   🏗️ ponytail       │   🗣️ caveman          │
│   "O QUE construir" │   "COMO responder"    │
│   • Escada YAGNI    │   • Compressão de     │
│   • Stdlib primeiro │     prosa             │
│   • Código mínimo   │   • Sem filler        │
│                     │   • Fragmentos        │
├─────────────────────┴───────────────────────┤
│     RESPOSTA: código mínimo + explicação    │
│              ultra-terse                    │
└─────────────────────────────────────────────┘
```

### Veredito

As duas skills **NÃO competem** — elas **se complementam** e foram projetadas para serem usadas juntas.

| Cenário | Recomendação |
|---|---|
| **Tarefa de código** | ponytail + caveman |
| **Pergunta conceitual / prosa** | Só caveman (ponytail: "Do NOT use for non-coding requests") |
| **Code review** | `/ponytail-review` + caveman |
| **Auditoria de código** | `/ponytail-audit` |
| **Apenas quer reduzir tokens** | Só caveman |

---

## 3. Arquitetura das três skills: golang-pro + ponytail + caveman

### Stack instalado

```
✓ caveman       (~/.claude/skills/caveman/)
✓ ponytail      (~/.claude/skills/ponytail/)
✓ golang-pro    (~/.claude/skills/golang-pro/)
```

### Três dimensões ortogonais

```
┌──────────────────────────────────────────────────────────────┐
│                                                              │
│   🏗️  golang-pro     →   DOMÍNIO: "Como fazer em Go"        │
│      15.6K installs      Idioma, padrões, qualidade,         │
│      role: specialist    ferramentas, constraints            │
│                                                              │
│   ✂️  ponytail       →   ESCOPO: "Quanto código escrever"   │
│      10.9K installs      YAGNI, minimalismo,                 │
│      role: scope         7-rung ladder, "só o necessário"   │
│                                                              │
│   🗣️  caveman        →   ESTILO: "Como comunicar"           │
│      322.2K installs     Prosa comprimida,                   │
│      role: style         sem filler, sem narração            │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

### Constraints simultâneas (golang-pro + ponytail)

```
golang-pro MUST DO:           ponytail MUST NOT:
├─ context.Context em         ├─ Interface com 1
│  todas ops bloqueantes         implementação
├─ Erros explícitos           ├─ Factory p/ 1 produto
│  (sem naked returns)        ├─ Config p/ valor
├─ Documentar exports            que nunca muda
├─ Error wrapping c/ %w       ├─ Scaffolding "for later"
├─ gofmt + golangci-lint      ├─ Abstrações não solicitadas
├─ Table-driven tests         └─ Boilerplate
├─ 80%+ coverage
└─ race detector
```

A tensão é produtiva: `ponytail` impede over-engineering da **solução**, `golang-pro` garante que a solução mínima seja **correta**.

---

## 4. Fluxo de desenvolvimento: as 5 fases

### Fase 0 — Kickoff

```bash
# Skills já instaladas. Ativar modos:
# ponytail: full (default) + caveman: full (default)
# golang-pro ativa automaticamente ao detectar código Go
```

### Fase 1 — Análise: entender antes de construir

| Quem age | O que faz |
|---|---|
| **golang-pro** step 1 | Analisa estrutura de módulos, interfaces existentes, padrões de concorrência |
| **ponytail** rungs 1-3 | A feature precisa existir? (YAGNI). Já existe no codebase? Stdlib cobre? |
| **caveman** | Comprime toda a comunicação durante a análise |

### Fase 2 — Design: modelar o mínimo viável

| Quem age | O que faz |
|---|---|
| **golang-pro** step 2 | Cria interfaces pequenas e focadas, composição sobre herança |
| **ponytail** rungs 4-6 | Native platform? Dependência já instalada resolve? Dá pra ser uma linha? |
| **caveman** | Mantém a discussão de design em prosa comprimida |

### Fase 3 — Implementação: código mínimo, Go idiomático

| Quem age | O que faz |
|---|---|
| **golang-pro** steps 3-4 | Escreve Go idiomático: context propagation, error wrapping, `go vet`, `golangci-lint` |
| **ponytail** rung 7 | "Can it be one line? Only then: the minimum code that works." |
| **caveman** full | Explicações em prosa telegráfica, blocos de código intactos |

### Fase 4 — Qualidade: testes, profiling, validação

**ponytail sai de cena.** Qualidade não é "código extra" — é correção.

| Quem age | O que faz |
|---|---|
| **golang-pro** steps 5-6 | `pprof`, benchmarks, table-driven tests c/ `-race`, fuzzing, 80%+ coverage |
| **ponytail** | Não interfere — quality gates não são "bloat" |
| **caveman** | Auto-clareza ativa se houver warning de segurança |

### Fase 5 — Entrega: comunicar o resultado

| Quem age | O que faz |
|---|---|
| **caveman** full | "Rate limit gRPC interceptor. Adaptado do HTTP middleware existente." |
| **ponytail** output | `skipped: [X], add when [Y]` |
| **golang-pro** template | Interface → implementação → teste → explicação de concorrência |

---

## 5. Pipeline completo: do Negócio ao Código

```
┌─────────────────────────────────────────────────────────────────────┐
│                      FASE DE NEGÓCIO (pronta)                        │
│  PROJECT_CHARTER.md  ·  BUSINESS_REQUIREMENTS.md  ·  EPICS.md       │
│  FEATURES.md  ·  USER_STORIES.md                                     │
├─────────────────────────────────────────────────────────────────────┤
│                   DIRETRIZES GLOBAIS (prontas)                       │
│  GLOBAL_SECURITY.md  ·  GLOBAL_ARCHITECT.md                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│            🛑 FASE DE ESPECIFICAÇÃO TÉCNICA (a fazer)               │
│            SPECS.md  ·  TASKS.md  ·  TEST_PLAN.md                   │
│                                                                      │
│            ┌──────────────────────────────────────┐                  │
│            │  caveman:  ❌ OFF                     │                  │
│            │  ponytail: ❌ OFF                     │                  │
│            │  Motivo: documentos de especificação  │                  │
│            │  precisam de precisão total — não     │                  │
│            │  são prosa interativa nem código      │                  │
│            └──────────────────────────────────────┘                  │
│                                                                      │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│            ✅ FASE DE DESENVOLVIMENTO (depois das specs)             │
│                                                                      │
│            ┌──────────────────────────────────────┐                  │
│            │  caveman:  ✅ ON (full)              │                  │
│            │  ponytail: ✅ ON (full)              │                  │
│            │  + golang-pro (standby automático)   │                  │
│            └──────────────────────────────────────┘                  │
│                                                                      │
└─────────────────────────────────────────────────────────────────────┘
```

### Por que NÃO usar durante a Fase de Especificação

#### caveman desligado

| Regra do caveman | Por que SPECS.md/TASKS.md/TEST_PLAN.md viola |
|---|---|
| *"Multi-step sequences where fragment order risk misread"* | TASKS.md lista tarefas com dependências. Fragmentos ambíguos geram retrabalho. |
| *"Compression itself creates technical ambiguity"* | SPECS.md define contratos técnicos. Quem implementa daqui 3 semanas precisa de precisão. |
| *"Code/commits/PRs: write normal"* | Especificações são artefatos permanentes como commits — precisam de precisão total. |

#### ponytail desligado

O SKILL.md do `ponytail` é explícito:

> *"Do NOT use for non-coding requests (general knowledge, prose, translation, summaries, recipes)."*

A escada de 7 degraus é focada em código. Não se aplica a planejamento de tarefas.

### ⚠️ Exceção: YAGNI manual na criação do TASKS.md

Ao quebrar FEATURES.md em TASKS.md, aplicar **manualmente** a pergunta do degrau 1:

> *"Does this need to exist at all?"*

```
FEATURES.md                           TASKS.md
──────────                           ─────────
"Rate limiting nos handlers gRPC" →  T-001: Criar interface RateLimiter
                                     T-002: Implementar token bucket
                                     T-003: Unary interceptor gRPC
                                     T-004: Testes table-driven com -race

Pergunta YAGNI manual: "Tarefa T-005: Dashboard de métricas de rate limit"
→ Não é necessária agora. Removida antes de entrar no backlog.
```

### O gate de ativação

```
                    ESPECIFICAÇÃO                   DESENVOLVIMENTO
                   ──────────────                  ────────────────

  SPECS.md ✅            │                                │
  TASKS.md ✅            │     🚦 GATE DE ATIVAÇÃO        │
  TEST_PLAN.md ✅        │────────────────────────────────│
                         │                                │
                    Time negocios                     Time técnico
                    aprovou as specs               começa a codar
                         │                                │
                    ❌ caveman                       ✅ caveman full
                    ❌ ponytail                      ✅ ponytail full
                                                     ✅ golang-pro
```

---

## 6. Controles durante o desenvolvimento

| Comando | Efeito |
|---|---|
| `/ponytail lite` | Build normal, sugere alternativa mais enxuta |
| `/ponytail full` | Escada YAGNI ativa (default para desenvolvimento) |
| `/ponytail ultra` | Extremista — desafia até o requisito |
| `/caveman lite` | Prosa profissional mas enxuta |
| `/caveman full` | Fragmentos, sem artigos (default) |
| `stop caveman` / `stop ponytail` | Volta ao modo normal |
| `/ponytail-review` | Revisa o diff atual contra a escada |
| `/ponytail-audit` | Audita o repo inteiro por over-engineering |
| `/ponytail-debt` | Track de shortcuts intencionais (comentários `// ponytail:`) |

---

## 7. Framework completo de fases

```
FASE                    SKILLS ATIVAS          QUEM CRIA/EXECUTA
──────────────────────  ─────────────────────  ──────────────────
PROJECT_CHARTER.md      (nenhuma)              Time de negócios
BUSINESS_REQUIREMENTS   (nenhuma)              Time de negócios
EPICS.md                (nenhuma)              Time de negócios
FEATURES.md             (nenhuma)              Time de negócios
USER_STORIES.md         (nenhuma)              Time de negócios
GLOBAL_SECURITY.md      (nenhuma)              Arquiteto/SecOps
GLOBAL_ARCHITECT.md     (nenhuma)              Arquiteto
──────────────────────  ─────────────────────  ──────────────────
SPECS.md                (nenhuma)              Time técnico + AI
TASKS.md                (nenhuma)*             Time técnico + AI
TEST_PLAN.md            (nenhuma)              Time técnico + AI
                        *Aplicar YAGNI manual
                         na quebra de tasks
──────────────────────  ─────────────────────  ──────────────────
🚦 GATE DE ATIVAÇÃO ← specs aprovadas pelo time de negócios
──────────────────────  ─────────────────────  ──────────────────
DESENVOLVIMENTO         caveman:full           Time técnico + AI
                        ponytail:full           + stack golang-pro
                        golang-pro:standby
──────────────────────  ─────────────────────  ──────────────────
CODE REVIEW             /ponytail-review       Time técnico + AI
                        caveman:full
──────────────────────  ─────────────────────  ──────────────────
AUDITORIA PÓS-FEATURE   /ponytail-audit        Time técnico + AI
                        /ponytail-debt
```

---

## 8. Recomendações finais

1. **Use JuliusBrussee/caveman** — é a versão canônica. A versão do mattpocock foi removida pelo próprio autor.

2. **caveman + ponytail são complementares por design** — o autor do ponytail documentou "pair with Caveman" no SKILL.md.

3. **golang-pro + ponytail + caveman formam um stack coeso:**
   - `golang-pro` → código idiomático e correto
   - `ponytail` → código mínimo e necessário
   - `caveman` → comunicação tersa e direta

4. **Ative o stack somente no desenvolvimento** — durante a fase de especificação (SPECS.md, TASKS.md, TEST_PLAN.md), mantenha caveman e ponytail desligados. Especificações precisam de precisão total.

5. **Aplique YAGNI manualmente na criação de tarefas** — o raciocínio do ponytail é útil para eliminar tarefas desnecessárias, mas como raciocínio humano, não como skill ativa.

---

## 9. Prompts para cada fase

### 9.1 DESENVOLVIMENTO — Implementar uma tarefa

- **Skills ativas:** `caveman:full` + `ponytail:full` + `golang-pro` (standby automático)
- **Momento:** Depois do gate de ativação, com SPECS.md/TASKS.md/TEST_PLAN.md aprovados

#### Prompt base (tarefa individual)

```markdown
/prompts:executar-tarefa     ← carrega o prompt template do projeto

Implementar task [T-XXX] do TASKS.md.

Contexto:
- SPECS.md: [caminho]
- TASKS.md: [caminho]
- TEST_PLAN.md: [caminho]
- GLOBAL_SECURITY.md: [caminho]
- GLOBAL_ARCHITECT.md: [caminho]

Antes de começar:
1. Ler a task no TASKS.md, as specs relacionadas no SPECS.md, e os casos de teste no TEST_PLAN.md
2. Verificar GLOBAL_SECURITY.md e GLOBAL_ARCHITECT.md para constraints aplicáveis
3. Subir a escada ponytail: isso precisa existir? Já existe no codebase? Stdlib cobre?
4. Só então implementar

Entregar:
1. Código (interfaces → implementação → testes)
2. go vet + golangci-lint limpos
3. Testes table-driven com -race passando
4. Output ponytail: skipped: [X], add when [Y]
```

#### Variação: implementar feature completa (múltiplas tasks relacionadas)

```markdown
/prompts:executar-tarefa

Implementar feature [FEAT-XXX] do FEATURES.md.

Tasks envolvidas: [T-001, T-002, T-003, ...]
Ordem: [T-001 → T-002 → T-003]

Para cada task:
1. Verificar pré-requisitos da task anterior concluídos
2. Subir a escada ponytail
3. Implementar com constraints golang-pro
4. go vet + golangci-lint + testes -race
5. Confirmar que passa antes de avançar para próxima task

Após todas as tasks:
1. Integration test cobrindo o fluxo completo
2. Atualizar TASKS.md marcando tasks concluídas
```

#### Variação: bug fix

```markdown
/prompts:executar-tarefa

Bug: [descrição do bug]. Reportado em: [issue/ticket].

Contexto: [SPECS.md, TEST_PLAN.md, GLOBAL_SECURITY.md]

Antes de corrigir:
1. Reproduzir o bug (confirmar que existe)
2. grep em todos os callers da função afetada — o fix lazy é o root-cause fix,
   não patch no caller que o ticket menciona
3. Subir escada ponytail: o fix mínimo que cobre todos os callers

Entregar:
1. Root-cause fix (não sintoma)
2. Teste que reproduz o bug e passa com o fix
3. go vet + golangci-lint limpos
4. Se aplicável: comentário // ponytail: com ceiling e upgrade path
```

---

### 9.2 CODE REVIEW

- **Skills ativas:** `/ponytail-review` + `caveman:full`
- **Momento:** Após concluir uma task ou feature, antes de integrar

#### Prompt: review de diff da task concluída

```markdown
/ponytail-review

Revisar diff da task [T-XXX] recém-concluída.

Contra:
- TASKS.md: [T-XXX] — a implementação cobre todos os requisitos?
- SPECS.md: [seção relevante] — o contrato técnico foi respeitado?
- TEST_PLAN.md: [casos de teste relevantes] — todos os casos cobertos?
- GLOBAL_SECURITY.md — viola algum constraint de segurança?
- GLOBAL_ARCHITECT.md — viola algum princípio arquitetural?

Escada ponytail reversa (da implementação pra cima):
1. Rung 7: o código é o mínimo que funciona? Ou tem abstração prematura?
2. Rung 6: algo poderia ser uma linha em vez de várias?
3. Rung 5: usou dependência já instalada? Adicionou dependência nova desnecessária?
4. Rung 4: usou feature nativa da plataforma? (ex: stdlib Go em vez de lib externa)
5. Rung 3: usou stdlib? Ou reimplementou o que já existe?
6. Rung 2: reusou código existente no codebase? Ou duplicou?

Constraints golang-pro:
- context.Context em todas as operações bloqueantes?
- Erros explicitamente tratados (sem _ assignment)?
- Error wrapping com %w?
- Naked returns justificados ou acidentais?
- Goroutines com lifecycle management claro?
- gofmt + golangci-lint limpos?

Segurança (GLOBAL_SECURITY.md):
- Input validation nos trust boundaries?
- Error handling previne data loss?
- Nenhuma informação sensível em logs/erros?

Reportar:
1. Aprovado / Aprovado com ressalvas / Rejeitado
2. Itens que violam a escada (com rung number)
3. Sugestões de simplificação (código específico, não genérico)
4. Comentários // ponytail: ausentes onde deveriam existir
```

#### Prompt: review pré-merge (feature completa)

```markdown
/ponytail-review

Review pré-merge da feature [FEAT-XXX]. Tasks: [T-001...T-00N].

Escopo: diff completo da branch contra main.

Checklist:
1. Cobertura: todas as tasks do TASKS.md implementadas?
2. Ordem: dependências entre tasks respeitadas?
3. Integração: as tasks funcionam juntas? Teste de integração cobre o fluxo completo?
4. Segurança: GLOBAL_SECURITY.md — cada trust boundary validado?
5. Arquitetura: GLOBAL_ARCHITECT.md — padrões e camadas respeitados?
6. Dívida: comentários // ponytail: documentam todos os shortcuts intencionais?
7. Testes: TEST_PLAN.md — todos os casos de teste passam? (-race clean)

Escada ponytail reversa sobre o diff completo:
[mesma escada do prompt anterior]

Reportar:
1. Veredito: merge approved / changes requested / blocked
2. Bloqueantes (security, data loss, arquitetura)
3. Não-bloqueantes (simplificações sugeridas, comentários ausentes)
4. Itens para /ponytail-debt tracking
```

---

### 9.3 AUDITORIA PÓS-FEATURE

**Skills ativas:** `/ponytail-audit` + `/ponytail-debt` + `caveman:full`
**Momento:** Feature completa e integrada, antes do fechamento do ciclo

#### Prompt: auditoria de over-engineering

```markdown
/ponytail-audit

Auditar o repo após a feature [FEAT-XXX].

Foco: over-engineering, código morto, duplicação, abstrações prematuras.

Varredura:
1. Interfaces com uma única implementação → candidate to inline
2. Factories com um único produto → candidate to remove
3. Configs com valores que nunca mudam → candidate to hardcode
4. Scaffolding "for later" não utilizado → candidate to delete
5. Código duplicado entre a feature nova e o codebase existente → candidate to deduplicate
6. Dependências adicionadas que poderiam ser substituídas por stdlib
7. Abstractions that don't carry their weight (camada de serviço com 1 método, etc.)

Contra as specs:
- SPECS.md: algo foi construído além do especificado? (gold-plating)
- TASKS.md: tasks concluídas cobrem exatamente o especificado? (nem mais, nem menos)

Reportar:
1. Arquivos/problemas encontrados (path + rung violado + sugestão)
2. Ordenado por impacto: linhas que podem ser removidas, depois simplificações
3. Falsos positivos: itens que parecem over-engineering mas são justificados
   (ex: interface com 1 impl porque o TEST_PLAN exige mock para testes)
```

#### Prompt: tracking de dívida técnica intencional

```markdown
/ponytail-debt

Catalogar todos os shortcuts intencionais da feature [FEAT-XXX].

Buscar por:
1. Comentários // ponytail: no diff da feature
2. Comentários # ponytail: nos arquivos de script/config
3. TODO/FIXME/HACK que são decisões conscientes de simplicidade, não bugs

Para cada item encontrado:
- Localização (arquivo:linha)
- O que foi simplificado
- Ceiling documentado (ex: "global lock, per-account locks if throughput matters")
- Quando revisitar (condição objetiva: "quando >1000 req/s", "quando >1 tenant")

Gerar:
1. Catálogo em markdown (para commits futuros ou /ponytail-gain)
2. Itens ordenados por risco: maior risco de estourar o ceiling primeiro
3. Itens sem ceiling documentado → flag como incompletos
```

#### Prompt: validação contra TEST_PLAN.md

```markdown
Validar a feature [FEAT-XXX] contra o TEST_PLAN.md.

Skills ativas: caveman:full

1. Executar todos os casos de teste do TEST_PLAN.md para esta feature
2. Para cada caso:
   - Status: PASS / FAIL / NOT IMPLEMENTED
   - Se FAIL: root cause
   - Se NOT IMPLEMENTED: justificativa (não escopo? adiado? esquecido?)
3. Cobertura: % de casos de teste cobertos pela implementação
4. Testes table-driven com -race: todos passam?
5. Fuzzing executado nos parsers/validators?
6. Benchmarks executados? Resultados dentro do esperado?

Reportar:
1. Tabela de resultados (caso → status)
2. Gaps: casos de teste não cobertos (com justificativa ou ação necessária)
3. Métricas: coverage %, race detector status, benchmark results
```

---

### 9.4 Resumo dos prompts por fase (com artefatos de saída)

```
FASE            PROMPT                  ARTEFATO GERADO                  DURAÇÃO
──────────────  ──────────────────────  ──────────────────────────────  ────────────
DESENVOLVIMENTO Prompt 9.1 (task)      DEVELOPER_TASK_[ID]_REPORT.md    10-30 min
                Prompt 9.1 (feature)   DEVELOPER_FEATURE_[ID]_REPORT.md 1-4 horas
                Prompt 9.1 (bug fix)   DEVELOPER_BUGFIX_[ID]_REPORT.md  5-20 min
──────────────  ──────────────────────  ──────────────────────────────  ────────────
CODE REVIEW     Prompt 9.2 (task)      CODE_REVIEW_TASK_[ID]_REPORT.md  5-10 min
                Prompt 9.2 (pré-merge) CODE_REVIEW_FEATURE_[ID]_REPORT  15-30 min
                                       .md
──────────────  ──────────────────────  ──────────────────────────────  ────────────
AUDITORIA       Prompt 9.3 (audit)     AUDIT_FEATURE_[ID]_REPORT.md     10-20 min
                Prompt 9.3 (debt)      DEBT_FEATURE_[ID]_CATALOG.md    5-10 min
                Prompt 9.3 (test plan) TEST_VALIDATION_FEATURE_[ID]     10-15 min
                                       _REPORT.md
```

---

### 9.5 Template de prompt executável (com geração de artefato)

Este template deve ser salvo como `.specs/prompts/PROMPT-EXECUTE-TASK.md` no repositório e carregado via `/prompts:executar-tarefa`:

```markdown
# Prompt: Executar Tarefa de Desenvolvimento

## Skills ativas obrigatórias
- caveman:full
- ponytail:full
- golang-pro (automático ao detectar Go)

## Documentos de referência (INPUT — carregar antes de iniciar)
- SPECS.md
- TASKS.md
- TEST_PLAN.md
- GLOBAL_SECURITY.md
- GLOBAL_ARCHITECT.md
- [Se existir] DEVELOPER_TASK_[T-XXX]_REPORT.md (tasks predecessoras)
- [Se existir] CODE_REVIEW_TASK_[T-XXX]_REPORT.md (review anterior da mesma área)

## Fluxo

### Pré-implementação
1. Ler a task no TASKS.md e as specs relacionadas no SPECS.md
2. Identificar os casos de teste no TEST_PLAN.md que cobrem esta task
3. Verificar constraints aplicáveis em GLOBAL_SECURITY.md e GLOBAL_ARCHITECT.md
4. Se houver tasks predecessoras, ler os DEVELOPER_*_REPORT.md delas
5. Subir a escada ponytail (7 rungs) — registrar decisão de CADA rung

### Implementação
6. Escrever interfaces primeiro (golang-pro step 2)
7. Implementar com context propagation, error wrapping, zero naked returns
8. Rodar go vet ./... e golangci-lint run — registrar output completo
9. Escrever table-driven tests com -race
10. Confirmar coverage ≥80% — registrar percentual exato

### Pós-implementação
11. Executar /ponytail-review no diff
12. Documentar shortcuts com comentários // ponytail:
13. Atualizar TASKS.md: marcar task como concluída

### Geração de artefato (OBRIGATÓRIO)
14. Gerar arquivo `.specs/reports/DEVELOPER_TASK_[T-XXX]_REPORT.md`
    usando o template da Seção 10.1 como estrutura base

## Output esperado
- Código (interfaces → implementação → testes)
- go vet + golangci-lint: clean
- Testes: PASS com -race
- Output ponytail: skipped: [X], add when [Y]
- Resumo caveman: 2-3 linhas do que foi feito
- **ARTEFATO:** `.specs/reports/DEVELOPER_TASK_[T-XXX]_REPORT.md`
```

---

### 9.6 Prompts com instrução de artefato incorporada

#### 9.6.1 DESENVOLVIMENTO — Task individual (com artefato)

```markdown
/prompts:executar-tarefa

Implementar task [T-XXX] do TASKS.md.

INPUT:
- SPECS.md: [caminho]
- TASKS.md: [caminho]
- TEST_PLAN.md: [caminho]
- GLOBAL_SECURITY.md: [caminho]
- GLOBAL_ARCHITECT.md: [caminho]
- [se existir] .specs/reports/DEVELOPER_TASK_[T-PREDECESSOR]_REPORT.md

ARTEFATO DE SAÍDA (obrigatório):
.specs/reports/DEVELOPER_TASK_[T-XXX]_REPORT.md
(estrutura conforme template da Seção 10.1)

Fluxo: Pré-implementação → Implementação → Pós-implementação → Artefato
```

#### 9.6.2 CODE REVIEW — Review de task (com artefato)

```markdown
/ponytail-review

Revisar diff da task [T-XXX].

INPUT:
- .specs/reports/DEVELOPER_TASK_[T-XXX]_REPORT.md (obrigatório)
- SPECS.md: [seção relevante]
- TEST_PLAN.md: [casos de teste relevantes]
- GLOBAL_SECURITY.md
- GLOBAL_ARCHITECT.md

ARTEFATO DE SAÍDA (obrigatório):
.specs/reports/CODE_REVIEW_TASK_[T-XXX]_REPORT.md
(estrutura conforme template da Seção 10.2)

Checklist: escada ponytail reversa (7→1) + constraints golang-pro + GLOBAL_SECURITY.md
```

#### 9.6.3 AUDITORIA — Feature completa (com artefato)

```markdown
/ponytail-audit + /ponytail-debt

Auditar feature [FEAT-XXX] pós-desenvolvimento.

INPUT:
- .specs/reports/DEVELOPER_TASK_[T-*]_REPORT.md (todos da feature)
- .specs/reports/CODE_REVIEW_TASK_[T-*]_REPORT.md (todos da feature)
- SPECS.md
- TASKS.md
- TEST_PLAN.md

ARTEFATOS DE SAÍDA (obrigatório):
.specs/reports/AUDIT_FEATURE_[FEAT-XXX]_REPORT.md
.specs/reports/DEBT_FEATURE_[FEAT-XXX]_CATALOG.md
.specs/reports/TEST_VALIDATION_FEATURE_[FEAT-XXX]_REPORT.md
(estruturas conforme templates da Seção 10.3, 10.4, 10.5)
```

---

### 9.7 Anti-padrões: o que NÃO fazer em cada prompt

| Fase | Anti-padrão | Por que | Correção |
|---|---|---|---|
| **DEV** | "Constrói um sistema completo de [X]" | Viola rung 1 (YAGNI) | "Implementar task T-XXX do TASKS.md" |
| **DEV** | Não referenciar GLOBAL_SECURITY.md | Pode gerar código com vulnerabilidade | Sempre incluir no contexto |
| **DEV** | "Usa a biblioteca [X] para resolver" | Viola rung 3-4-5 | Deixar o agente decidir a ferramenta |
| **DEV** | Não gerar artefato de saída | Próxima fase fica cega | Artefato é obrigatório, não opcional |
| **REVIEW** | "Revisa o código" (genérico) | Sem critérios objetivos | Checklist com rungs e constraints |
| **REVIEW** | Fazer review sem ler DEVELOPER_*_REPORT.md | Perde decisões documentadas | Artefato da fase anterior é INPUT |
| **REVIEW** | Fazer review com ponytail desligado | Review não detecta over-engineering | `/ponytail-review` obrigatório |
| **AUDIT** | Auditar feature pela metade | Dívida não catalogada = dívida esquecida | Audit + debt + test validation juntos |
| **AUDIT** | Pular validação contra TEST_PLAN.md | Gaps de teste viram bugs em produção | Prompt 9.3 (test plan) é obrigatório |
| **GERAL** | Usar caveman durante a especificação | Ambiguidade em specs gera retrabalho | `stop caveman` antes de editar specs |
| **GERAL** | Usar ponytail para tasks não-código | "Do NOT use for non-coding requests" | Skill não foi feita para prosa |
| **GERAL** | Artefato gerado sem dados estruturados | Agente downstream não consegue parse | Seguir template de seções fixas |

---

## 10. Artefatos de saída (reports)

### 10.0 Por que isso faz sentido: o argumento de design

```
┌─────────────────────────────────────────────────────────────────┐
│                    ARTIFACT CHAIN PATTERN                        │
│                                                                  │
│  SPECS.md ────────────┐                                          │
│  TASKS.md ────────────┤                                          │
│  TEST_PLAN.md ────────┤                                          │
│  GLOBAL_SECURITY.md ──┤                                          │
│  GLOBAL_ARCHITECT.md ─┤                                          │
│                        │                                         │
│                        ▼                                         │
│  ┌──────────────────────────────────────────┐                    │
│  │       /prompts:executar-tarefa            │                    │
│  │       Skills: caveman + ponytail          │                    │
│  │              + golang-pro                 │                    │
│  └────────────────────┬─────────────────────┘                    │
│                       │                                          │
│                       ▼                                          │
│  ┌──────────────────────────────────────────┐                    │
│  │  .specs/reports/                          │                    │
│  │  DEVELOPER_TASK_T-001_REPORT.md           │ ← INPUT p/ review │
│  │  ┌────────────────────────────────────┐   │                    │
│  │  │ • Decisão de cada rung da escada   │   │                    │
│  │  │ • Código gerado (paths)            │   │                    │
│  │  │ • Testes executados (-race status) │   │                    │
│  │  │ • go vet/golangci-lint output      │   │                    │
│  │  │ • Coverage % exato                 │   │                    │
│  │  │ • Skipped items + justificativa    │   │                    │
│  │  │ • // ponytail: comments catalog    │   │                    │
│  │  │ • Desvios dos specs originais      │   │                    │
│  │  └────────────────────────────────────┘   │                    │
│  └────────────────────┬─────────────────────┘                    │
│                       │                                          │
│                       ▼                                          │
│  ┌──────────────────────────────────────────┐                    │
│  │       /ponytail-review                    │                    │
│  │       Skills: ponytail-review + caveman   │                    │
│  └────────────────────┬─────────────────────┘                    │
│                       │                                          │
│                       ▼                                          │
│  ┌──────────────────────────────────────────┐                    │
│  │  .specs/reports/                          │                    │
│  │  CODE_REVIEW_TASK_T-001_REPORT.md         │ ← INPUT p/ audit  │
│  │  ┌────────────────────────────────────┐   │                    │
│  │  │ • Escada reversa (7→1) por rung   │   │                    │
│  │  │ • Violações encontradas            │   │                    │
│  │  │ • Sugestões de simplificação       │   │                    │
│  │  │ • Veredito + bloqueantes           │   │                    │
│  │  │ • // ponytail: ausentes            │   │                    │
│  │  └────────────────────────────────────┘   │                    │
│  └────────────────────┬─────────────────────┘                    │
│                       │                                          │
│                       ▼                                          │
│  ┌──────────────────────────────────────────┐                    │
│  │  /ponytail-audit + /ponytail-debt         │                    │
│  │  + validação TEST_PLAN.md                 │                    │
│  └────────────────────┬─────────────────────┘                    │
│                       │                                          │
│                       ▼                                          │
│  ┌──────────────────────────────────────────┐                    │
│  │  .specs/reports/                          │                    │
│  │  AUDIT_FEATURE_FEAT-001_REPORT.md         │ ← arquivo final    │
│  │  DEBT_FEATURE_FEAT-001_CATALOG.md         │                    │
│  │  TEST_VALIDATION_FEATURE_FEAT-001_REPORT  │                    │
│  └──────────────────────────────────────────┘                    │
│                                                                  │
│  Cada artefato é OUTPUT da fase atual e INPUT da fase seguinte.  │
│  A cadeia fecha o ciclo: negócio → specs → código → review →    │
│  auditoria, com rastreabilidade completa em cada etapa.          │
└─────────────────────────────────────────────────────────────────┘
```

#### Os 5 benefícios concretos

| # | Benefício | Quem ganha | Exemplo |
|---|---|---|---|
| 1 | **Pipeline continuity** | Agentes downstream | O review lê o `DEVELOPER_REPORT.md` e sabe exatamente o que foi feito, sem precisar do contexto completo da conversa de desenvolvimento |
| 2 | **Human oversight** | Tech lead / arquiteto | Revisa decisões críticas (ex: "rung 3 — stdlib não cobre, justificativa: `x/net/rate` não é stdlib") sem ler código |
| 3 | **Institutional memory** | Time futuro (3-6 meses) | "Por que o rate limiter não tem TTL configurável?" → `DEVELOPER_TASK_T-003_REPORT.md`: "skipped: TTL config, add when multi-tenant" |
| 4 | **Compliance / audit trail** | SecOps / governance | Cadeia completa: requisito → spec → task → código → review → auditoria, cada passo com artefato assinado |
| 5 | **Debt ceiling tracking** | Manutenção | `DEBT_FEATURE-001_CATALOG.md` lista todos os `// ponytail:` com ceiling documentado. Daqui 3 meses, filtra por "ceiling atingido" |

#### Por que o caveman NÃO compromete os artefatos

O `caveman` opera na **comunicação interativa** (o diálogo durante o desenvolvimento). O artefato é gerado ao final, como output estruturado — mesmo tratamento de commits e PRs, onde o próprio SKILL.md do caveman diz "write normal". O report é um documento permanente, não prosa interativa.

#### Convenção de nomenclatura

```
.specs/reports/
├── DEVELOPER_TASK_[T-XXX]_REPORT.md
├── DEVELOPER_FEATURE_[FEAT-XXX]_REPORT.md
├── DEVELOPER_BUGFIX_[BUG-XXX]_REPORT.md
├── CODE_REVIEW_TASK_[T-XXX]_REPORT.md
├── CODE_REVIEW_FEATURE_[FEAT-XXX]_REPORT.md
├── AUDIT_FEATURE_[FEAT-XXX]_REPORT.md
├── DEBT_FEATURE_[FEAT-XXX]_CATALOG.md
└── TEST_VALIDATION_FEATURE_[FEAT-XXX]_REPORT.md
```

Regras:
- Prefixo: fase do pipeline (`DEVELOPER`, `CODE_REVIEW`, `AUDIT`, `DEBT`, `TEST_VALIDATION`)
- Tipo: granularidade (`TASK`, `FEATURE`, `BUGFIX`)
- ID: identificador exato do TASKS.md ou FEATURES.md
- Sufixo: `REPORT` para análise, `CATALOG` para listagens estruturadas
- Extensão: sempre `.md` (machine-parseable + human-readable)

---

### 10.1 Template: DEVELOPER_TASK_[T-XXX]_REPORT.md

````markdown
# Developer Task Report: [T-XXX]

> **Task:** [T-XXX] — [título da task]
> **Feature:** [FEAT-XXX] — [título da feature]
> **Data:** [YYYY-MM-DD]
> **Skills ativas:** caveman:full, ponytail:full, golang-pro
> **Inputs:** SPECS.md, TASKS.md, TEST_PLAN.md, GLOBAL_SECURITY.md, GLOBAL_ARCHITECT.md
> **Inputs upstream:** [DEVELOPER_TASK_T-PREDECESSOR_REPORT.md ou "N/A"]

---

## 1. Escada Ponytail — Decisão por rung

| Rung | Pergunta | Decisão | Justificativa |
|------|----------|---------|---------------|
| 1 | Precisa existir? (YAGNI) | SIM/NÃO | [justificativa] |
| 2 | Já existe no codebase? | SIM/NÃO | [path do que foi reusado ou justificativa] |
| 3 | Stdlib cobre? | SIM/NÃO | [pacote/função ou justificativa] |
| 4 | Native platform cobre? | SIM/NÃO | [feature nativa ou justificativa] |
| 5 | Dependência já instalada? | SIM/NÃO | [dependência ou justificativa] |
| 6 | Dá pra ser uma linha? | SIM/NÃO | [solução one-liner ou justificativa] |
| 7 | Mínimo que funciona | — | [descrição da solução adotada] |

## 2. Código gerado

### Arquivos criados
- [path] — [descrição]

### Arquivos modificados
- [path] — [o que mudou]

### Interfaces definidas
```go
// [interface definition]
```

## 3. Constraints golang-pro

| Constraint | Status | Detalhe |
|---|---|---|
| context.Context em ops bloqueantes | ✅/❌ | |
| Erros explícitos (sem naked returns) | ✅/❌ | |
| Error wrapping com %w | ✅/❌ | |
| Documentação de exports | ✅/❌ | |
| gofmt | ✅/❌ | |
| golangci-lint | ✅/❌ | [n issues, n fixed] |
| Goroutines com lifecycle | ✅/❌/N/A | |

## 4. Qualidade

| Métrica | Resultado |
|---|---|
| go vet | clean / [n] issues |
| golangci-lint | clean / [n] issues |
| Testes executados | [n] |
| Testes PASS | [n] |
| Race detector | clean / [n] races |
| Coverage | [XX]% |
| Benchmark | [ns/op, allocs/op] ou "N/A" |

## 5. Segurança (GLOBAL_SECURITY.md)

| Constraint aplicável | Status | Detalhe |
|---|---|---|
| [constraint] | ✅/❌/N/A | |
| Input validation nos trust boundaries | ✅/❌/N/A | |
| Error handling previne data loss | ✅/❌/N/A | |

## 6. Arquitetura (GLOBAL_ARCHITECT.md)

| Princípio aplicável | Status | Detalhe |
|---|---|---|
| [princípio] | ✅/❌/N/A | |

## 7. Skipped items (ponytail output)

| Item skipped | Motivo | Adicionar quando |
|---|---|---|
| [X] | [rung N — justificativa] | [condição objetiva] |

## 8. Catálogo de // ponytail: comments

| Arquivo:linha | Comentário | Ceiling | Upgrade path |
|---|---|---|---|
| [path]:[L] | [texto] | [limite] | [ação] |

## 9. Desvios das specs originais

| Spec original | O que foi entregue | Justificativa |
|---|---|---|
| [SPECS.md seção X] | [implementação real] | [rung da escada ou constraint] |

## 10. Resumo (caveman)

[2-3 linhas em prosa comprimida do que foi feito, o que foi pulado, e o próximo passo]
````

---

### 10.2 Template: CODE_REVIEW_TASK_[T-XXX]_REPORT.md

```markdown
# Code Review Report: [T-XXX]

> **Task:** [T-XXX] — [título]
> **Feature:** [FEAT-XXX]
> **Data:** [YYYY-MM-DD]
> **Skills ativas:** /ponytail-review, caveman:full
> **Inputs:** DEVELOPER_TASK_[T-XXX]_REPORT.md, SPECS.md, TEST_PLAN.md, GLOBAL_SECURITY.md, GLOBAL_ARCHITECT.md

---

## 1. Escada Ponytail Reversa (do código para cima)

| Rung | Pergunta (reversa) | Status | Evidência |
|------|--------------------|--------|-----------|
| 7 | O código é o mínimo que funciona? | ✅/❌ | [abstração prematura? layer desnecessária?] |
| 6 | Algo poderia ser uma linha? | ✅/❌ | [trecho que poderia ser inline] |
| 5 | Usou dependência instalada? Nova dep necessária? | ✅/❌ | [dep adicionada desnecessariamente?] |
| 4 | Usou feature nativa? | ✅/❌ | [stdlib não usada onde poderia?] |
| 3 | Stdlib usada onde possível? | ✅/❌ | [reimplementou algo que já existe?] |
| 2 | Reusou código existente? | ✅/❌ | [duplicação detectada?] |
| 1 | Isso precisava existir? (YAGNI) | ✅/❌ | [código que poderia ser deletado?] |

## 2. Violações encontradas

| # | Arquivo:linha | Rung violado | Descrição | Sugestão |
|---|---|---|---|---|
| 1 | [path]:[L] | [N] | [o que viola] | [como corrigir] |

## 3. Sugestões de simplificação

| # | Arquivo:linha | Código atual | Simplificação proposta | Impacto (linhas) |
|---|---|---|---|---|
| 1 | [path]:[L] | [snippet] | [snippet proposto] | -[N] |

## 4. Constraints golang-pro (verificação)

| Constraint | Status | Evidência |
|---|---|---|
| context.Context em ops bloqueantes | ✅/❌ | |
| Erros explícitos | ✅/❌ | |
| Error wrapping com %w | ✅/❌ | |
| Documentação de exports | ✅/❌ | |
| gofmt | ✅/❌ | |
| golangci-lint | ✅/❌ | |
| Goroutines com lifecycle | ✅/❌/N/A | |

## 5. Segurança (GLOBAL_SECURITY.md)

| Constraint | Status | Evidência |
|---|---|---|
| [constraint] | ✅/❌/N/A | |

## 6. Comentários // ponytail: ausentes

| Localização | Shortcut não documentado | Deveria ter |
|---|---|---|
| [path]:[L] | [descrição do atalho] | [comentário sugerido] |

## 7. Veredito

**Status:** APPROVED / CHANGES_REQUESTED / BLOCKED

**Bloqueantes (merge blockers):**
- [item] — [por que]

**Não-bloqueantes (sugestões):**
- [item] — [por que]

**Itens para /ponytail-debt tracking:**
- [item]
```

---

### 10.3 Template: AUDIT_FEATURE_[FEAT-XXX]_REPORT.md

```markdown
# Audit Report: [FEAT-XXX]

> **Feature:** [FEAT-XXX] — [título]
> **Data:** [YYYY-MM-DD]
> **Skills ativas:** /ponytail-audit, caveman:full
> **Inputs:** todos os DEVELOPER_TASK_*_REPORT.md e CODE_REVIEW_*_REPORT.md da feature

---

## 1. Varredura de over-engineering

| # | Arquivo:linha | Tipo | Descrição | Rung violado | Sugestão |
|---|---|---|---|---|---|
| 1 | [path]:[L] | interface 1 impl | [desc] | 7 | inline |
| 2 | [path]:[L] | factory 1 produto | [desc] | 7 | remove |
| 3 | [path]:[L] | config imutável | [desc] | 7 | hardcode |
| 4 | [path]:[L] | scaffolding unused | [desc] | 1 | delete |
| 5 | [path]:[L] | duplicação | [desc] | 2 | deduplicate |
| 6 | [path]:[L] | dep desnecessária | [desc] | 5 | stdlib |

## 2. Impacto estimado

| Categoria | Itens | Linhas removíveis |
|---|---|---|
| Remover | [N] | -[N] |
| Simplificar | [N] | -[N] |
| Deduplicar | [N] | -[N] |
| **Total** | **[N]** | **-[N]** |

## 3. Falsos positivos (over-engineering justificado)

| Arquivo:linha | Parece over-engineering | Justificativa |
|---|---|---|
| [path]:[L] | [desc] | [ex: interface com 1 impl porque TEST_PLAN exige mock] |

## 4. Gold-plating (construído além do SPECS.md)

| Spec original | O que foi entregue a mais | Deveria existir? |
|---|---|---|
| [SPECS.md seção X] | [extra] | SIM/NÃO — [justificativa] |

## 5. Cobertura de specs

| Spec (SPECS.md) | Task (TASKS.md) | Implementado? | Verificado em |
|---|---|---|---|
| [seção X] | [T-XXX] | ✅/❌ | [report] |

## 6. Resumo (caveman)

[2-3 linhas: total de achados, itens bloqueantes, recomendação]
```

---

### 10.4 Template: DEBT_FEATURE_[FEAT-XXX]_CATALOG.md

```markdown
# Debt Catalog: [FEAT-XXX]

> **Feature:** [FEAT-XXX] — [título]
> **Data:** [YYYY-MM-DD]
> **Skills ativas:** /ponytail-debt, caveman:full
> **Inputs:** diff da feature, todos os DEVELOPER_TASK_*_REPORT.md

---

## Catálogo de dívida técnica intencional

| # | Arquivo:linha | Comentário | O que foi simplificado | Ceiling | Upgrade path | Risco |
|---|---|---|---|---|---|---|
| 1 | [path]:[L] | `// ponytail: [texto]` | [desc] | [limite] | [ação] | 🔴🟡🟢 |
| 2 | [path]:[L] | `# ponytail: [texto]` | [desc] | [limite] | [ação] | 🔴🟡🟢 |

## Itens sem ceiling documentado (incompletos)

| # | Arquivo:linha | Comentário | Ação necessária |
|---|---|---|---|
| 1 | [path]:[L] | `// ponytail: [texto]` | Documentar ceiling e upgrade path |

## Ordenação por risco

### 🔴 Alto risco (ceiling próximo ou desconhecido)
- [item] — [por que]

### 🟡 Médio risco
- [item] — [por que]

### 🟢 Baixo risco (ceiling distante, upgrade path claro)
- [item] — [por que]

## Gatilhos de revisão

| Condição | Quando revisitar | Itens afetados |
|---|---|---|
| [ex: >1000 req/s] | [data ou evento] | [ids] |
| [ex: >1 tenant] | [data ou evento] | [ids] |
```

---

### 10.5 Template: TEST_VALIDATION_FEATURE_[FEAT-XXX]_REPORT.md

````markdown
# Test Validation Report: [FEAT-XXX]

> **Feature:** [FEAT-XXX] — [título]
> **Data:** [YYYY-MM-DD]
> **Skills ativas:** caveman:full
> **Inputs:** TEST_PLAN.md, todos os DEVELOPER_TASK_*_REPORT.md

---

## 1. Resultados por caso de teste

| Caso (TEST_PLAN.md) | Task | Status | Evidência |
|---|---|---|---|
| [TC-XXX] | [T-XXX] | ✅ PASS | [test name, arquivo] |
| [TC-XXX] | [T-XXX] | ❌ FAIL | [root cause] |
| [TC-XXX] | — | ⚠️ NOT IMPLEMENTED | [justificativa] |

## 2. Sumário

| Métrica | Valor |
|---|---|
| Total de casos | [N] |
| PASS | [N] ([XX]%) |
| FAIL | [N] ([XX]%) |
| NOT IMPLEMENTED | [N] ([XX]%) |

## 3. Gaps

| Caso não coberto | Impacto | Ação |
|---|---|---|
| [TC-XXX] | [desc] | [implementar / adiar / não escopo] |

## 4. Métricas de qualidade (golang-pro)

| Métrica | Resultado | Threshold |
|---|---|---|
| Coverage | [XX]% | ≥80% |
| Race detector | clean / [N] races | clean |
| go vet | clean / [N] issues | clean |
| golangci-lint | clean / [N] issues | clean |
| Benchmarks | [resultado] | [esperado] |
| Fuzzing | [N] execuções / [N] crashes | 0 crashes |

## 5. Resumo (caveman)

[2-3 linhas: status geral, gaps críticos, recomendação]
````

---

### 10.6 Localização dos artefatos no repositório

```
[PROJETO]/
├── .specs/
│   ├── prompts/
│   │   └── PROMPT-EXECUTE-TASK.md          ← template de prompt base
│   ├── reports/                             ← todos os artefatos gerados
│   │   ├── DEVELOPER_TASK_T-001_REPORT.md
│   │   ├── DEVELOPER_TASK_T-002_REPORT.md
│   │   ├── CODE_REVIEW_TASK_T-001_REPORT.md
│   │   ├── CODE_REVIEW_TASK_T-002_REPORT.md
│   │   ├── AUDIT_FEATURE_FEAT-001_REPORT.md
│   │   ├── DEBT_FEATURE_FEAT-001_CATALOG.md
│   │   └── TEST_VALIDATION_FEATURE_FEAT-001_REPORT.md
│   ├── SPECS.md
│   ├── TASKS.md
│   └── TEST_PLAN.md
├── GLOBAL_SECURITY.md
├── GLOBAL_ARCHITECT.md
└── ...
```
