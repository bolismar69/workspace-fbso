# Definition of Done (DoD)

> **Programa:** Adequação Corporativa à Reforma Tributária Nacional
> **Código:** PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
> **Versão:** 1.0
> **Atualizado:** 2026-07-08
> **Objetivo:** Estabelecer critérios objetivos, verificáveis e não-ambíguos para determinar quando uma task, feature ou fase está concluída. Este documento é o contrato compartilhado entre time técnico, Product Owner e agentes de IA.

---

## 1. Princípios

1. **Binário, não subjetivo.** Todo critério é verificável por máquina ou por checklist objetivo. Nenhum critério depende de "parece bom" ou "acho que está pronto."
2. **Acumulativo.** Uma task só está DONE quando TODOS os critérios aplicáveis são satisfeitos — não quando "a maioria" está.
3. **Independente de quem executa.** Os mesmos critérios valem para tasks implementadas por humanos e por agentes de IA. O artefato de saída (`DEVELOPER_*_REPORT.md`) é a evidência.
4. **Fases têm DoD cumulativo.** A DoD de uma feature inclui a DoD de todas as suas tasks. A DoD de uma fase inclui a DoD de todas as suas features.
5. **Skills não substituem a DoD.** `ponytail`, `caveman` e `golang-pro` são ferramentas de execução — a DoD é o critério de aceitação.

---

## 2. Níveis de DoD

```
DoD de TASK     →  menor unidade de entrega (1 task do TASKS.md)
DoD de FEATURE  →  conjunto de tasks relacionadas (1 feature do FEATURES.md)
DoD de FASE     →  conjunto de features (1 fase/onda do programa)
```

Cada nível herda os critérios do nível anterior e adiciona os seus próprios.

---

## 3. DoD de TASK

Uma task individual (ex: `T-001` do `TASKS.md`) está **DONE** quando TODOS os critérios abaixo são satisfeitos.

### 3.1 Código

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| C1 | Código implementa exatamente o especificado na task do TASKS.md — nem mais, nem menos | Diff contra a spec da task | `DEVELOPER_TASK_*_REPORT.md` — Seção 9 (Desvios) vazia ou com justificativa aprovada |
| C2 | Código compila sem erros | `go build ./...` | Exit code 0 |
| C3 | `go vet ./...` limpo | `go vet ./...` | Zero issues reportados |
| C4 | `golangci-lint run` limpo nos arquivos alterados | `golangci-lint run ./path/...` | Zero novos issues (issues pré-existentes em arquivos não alterados são aceitáveis) |
| C5 | `gofmt` aplicado em todos os arquivos | `gofmt -d .` | Zero diffs |
| C6 | Nenhum `fmt.Println`, `log.Print` de debug, código comentado, ou arquivo temporário | Busca manual + revisão do diff | Diff limpo de artefatos de desenvolvimento |
| C7 | Constantes mágicas extraídas e nomeadas | Revisão do diff | Valores com significado de negócio têm nome simbólico |
| C8 | Imports organizados (stdlib → third-party → interno) | `goimports -l .` | Zero arquivos com imports desorganizados |

### 3.2 Design e Arquitetura

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| D1 | Interfaces definidas antes da implementação (contract-first) | Revisão dos arquivos | `DEVELOPER_TASK_*_REPORT.md` — Seção 2 (Interfaces definidas) |
| D2 | Nenhuma interface com 1 única implementação (a menos que justificado por mock de teste) | `/ponytail-review` — rung 7 reverso | Justificativa documentada no report se aplicável |
| D3 | Nenhuma factory com 1 único produto | `/ponytail-review` — rung 7 reverso | — |
| D4 | Nenhuma abstração não solicitada pela spec | `/ponytail-review` — rung 1 reverso | — |
| D5 | Código respeita GLOBAL_ARCHITECT.md e ARCHITECTURE.md do projeto | Revisão contra os documentos | `DEVELOPER_TASK_*_REPORT.md` — Seção 6 |

### 3.3 Constraints golang-pro

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| G1 | `context.Context` em todas as operações bloqueantes (I/O, rede, lock) | Revisão do diff | Nenhuma função bloqueante sem `ctx context.Context` como primeiro parâmetro |
| G2 | Erros explicitamente tratados — zero `_` assignment para `error` sem justificativa | `golangci-lint` + revisão | Zero ocorrências de `_ = err` ou `val, _ := func()` |
| G3 | Error wrapping com `%w` em todas as fronteiras de camada | Revisão do diff | `fmt.Errorf("context: %w", err)` em todos os pontos de propagação |
| G4 | Zero naked returns em funções com mais de 1 retorno | Revisão do diff | `return nil, err` — nunca `return` |
| G5 | Goroutines com lifecycle management explícito (ctx, done channel, ou errgroup) | Revisão do diff | Toda goroutine tem caminho de cancelamento claro |
| G6 | Zero `panic` para tratamento de erro normal | Revisão do diff | `panic` reservado para invariantes quebrados (bug, não erro de runtime) |
| G7 | Documentação de todos os exports (funções, tipos, constantes) | `go doc` | Godoc renderiza para cada símbolo exportado |
| G8 | Configuração via functional options ou env vars — nunca hardcoded | Revisão do diff | Valores de config vêm de `os.Getenv` ou `Option` |

### 3.4 Testes

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| T1 | Table-driven tests implementados para toda lógica não-trivial | `go test -v ./path/...` | Testes com `t.Run()` e múltiplos casos |
| T2 | Testes passam com flag `-race` | `go test -race ./...` | Exit code 0, zero races detectadas |
| T3 | Coverage ≥ 80% no diff (arquivos novos ou alterados) | `go test -cover ./path/...` | Coverage report |
| T4 | Casos de teste do TEST_PLAN.md referentes a esta task estão implementados e passando | Checklist contra TEST_PLAN.md | `TEST_VALIDATION_*_REPORT.md` — casos marcados ✅ |
| T5 | Cenário de erro testado (não apenas happy path) | Revisão dos testes | Pelo menos 1 caso de teste com entrada inválida ou condição de erro |
| T6 | Fuzzing aplicado a parsers, validators e funções que processam input externo | `go test -fuzz=. -fuzztime=10s` | Zero crashes |

### 3.5 Segurança

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| S1 | Input validation em todos os trust boundaries (API handlers, message consumers) | Revisão do diff | Validação antes de qualquer processamento |
| S2 | Nenhuma credencial, token, chave de API ou dado sensível em hardcode | `grep -rE '(password|secret|token|key|api_key)' --include='*.go'` | Zero ocorrências em código fonte |
| S3 | Erros retornados ao cliente não expõem detalhes internos (stack trace, queries SQL, paths) | Revisão das mensagens de erro | Mensagens genéricas para o cliente, detalhes no log interno |
| S4 | Nenhuma violação do SECURITY.md | Checklist contra SECURITY.md | `DEVELOPER_TASK_*_REPORT.md` — Seção 5 (todas ✅) |
| S5 | Dados fiscais (cálculo de impostos, alíquotas) validados antes de persistência | Revisão do diff | Validação de ranges, tipos e integridade dos dados fiscais |

### 3.6 Documentação e Rastreabilidade

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| R1 | `DEVELOPER_TASK_[T-XXX]_REPORT.md` gerado em `.specs/reports/` | Arquivo existe | Report completo seguindo template da Seção 10.1 |
| R2 | `// ponytail:` comments documentados para todos os shortcuts intencionais, com ceiling e upgrade path | Busca por shortcuts não documentados | `DEVELOPER_TASK_*_REPORT.md` — Seção 8 |
| R3 | TASKS.md atualizado: task marcada como `[x]` | Verificar arquivo TASKS.md | Checkbox marcado |
| R4 | CHANGELOG.md atualizado (se o projeto tiver) com a entrada da task | Verificar CHANGELOG.md | Linha adicionada |

### 3.7 Revisão

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| V1 | `/ponytail-review` executado no diff — sem bloqueantes | Executar review | `CODE_REVIEW_TASK_*_REPORT.md` — veredito APPROVED |
| V2 | Escada ponytail reversa (7→1) validada — sem violações não justificadas | `/ponytail-review` | Todos os rungs ✅ ou com justificativa documentada |
| V3 | Nenhum código duplicado com o codebase existente | `/ponytail-review` — rung 2 reverso | — |

---

## 4. DoD de FEATURE

Uma feature (conjunto de tasks relacionadas do FEATURES.md) está **DONE** quando, ALÉM de todas as suas tasks individuais atenderem à DoD de TASK:

| # | Critério adicional de FEATURE | Como verificar | Evidência |
|---|-------------------------------|----------------|-----------|
| F1 | Todas as tasks da feature estão DONE (DoD de TASK satisfeita para cada uma) | Checklist contra TASKS.md | Todas as checkboxes `[x]` |
| F2 | Teste de integração cobre o fluxo completo da feature (end-to-end dentro do serviço) | `go test -race ./... -run Integration` | Teste de integração PASS |
| F3 | Contratos de API (gRPC proto / OpenAPI) atualizados se a feature expõe ou modifica endpoints | Verificar `.specs/api/` | Spec de API sincronizada com a implementação |
| F4 | Nenhuma regressão: testes de features já concluídas continuam passando | `go test -race ./...` | Suite completa verde |
| F5 | `CODE_REVIEW_FEATURE_[FEAT-XXX]_REPORT.md` gerado com veredito APPROVED | Arquivo existe | Report completo |
| F6 | PO / PM validou os critérios de aceite da feature | Demonstração para o PO | Aprovação registrada |
| F7 | Dependências entre tasks respeitadas (ordem de execução no TASKS.md) | Verificar ordem das tasks | Tarefas concluídas na sequência especificada |

---

## 5. DoD de FASE

Uma fase/onda do programa está **DONE** quando, ALÉM de todas as suas features atenderem à DoD de FEATURE:

| # | Critério adicional de FASE | Como verificar | Evidência |
|---|----------------------------|----------------|-----------|
| P1 | Todas as features da fase estão DONE | Checklist contra FEATURES.md e TASKS.md | Fase 100% concluída no TASKS.md |
| P2 | `AUDIT_FEATURE_[FASE]_REPORT.md` gerado — sem bloqueantes | Arquivo existe | Zero itens 🔴 |
| P3 | `DEBT_FEATURE_[FASE]_CATALOG.md` gerado — todos os `// ponytail:` catalogados com ceiling | Arquivo existe | Nenhum item sem ceiling documentado |
| P4 | `TEST_VALIDATION_FEATURE_[FASE]_REPORT.md` gerado — 100% dos casos do TEST_PLAN.md cobertos e passando | Arquivo existe | 100% PASS ou NOT IMPLEMENTED com justificativa aprovada |
| P5 | Sanity Check executado (limpeza, git status, localização, segurança, evidência) | Protocolo do PROMPT-EXECUTE-TASK.md | Todos os 6 passos ✅ |
| P6 | Métricas de qualidade da fase dentro dos thresholds: coverage ≥ 80%, race detector clean, golangci-lint clean | Relatórios agregados | Dashboard de qualidade da fase |
| P7 | KPIs de negócio da fase (MATRIZ-KPI.md) medidos e dentro da meta | Medição dos KPIs aplicáveis | Dashboard executivo atualizado |

---

## 6. Uso com Agentes de IA

### 6.1 Como o agente usa a DoD

O agente de IA deve:
1. Carregar este arquivo junto com os documentos de referência no início da execução
2. Para cada task, percorrer a checklist da Seção 3 e marcar cada item como ✅ ou ❌
3. Registrar o resultado no `DEVELOPER_TASK_*_REPORT.md`
4. Se algum item falhar, a task NÃO está DONE — corrigir e re-verificar

### 6.2 Relação com as skills

| Skill | Papel na DoD |
|-------|--------------|
| `golang-pro` | Garante os critérios G1–G8 (constraints Go) e T1–T6 (testes) |
| `ponytail` | Garante os critérios D2–D4 (design mínimo) e V1–V3 (revisão) |
| `caveman` | Comprime a comunicação durante a execução, mas NÃO atua sobre os artefatos (reports, specs) |
| `/ponytail-review` | Verifica os critérios V1–V3 (escada reversa) |
| `/ponytail-audit` | Verifica os critérios de fase P2–P4 |

### 6.3 Como o humano usa a DoD

- **Tech Lead:** audita `DEVELOPER_TASK_*_REPORT.md` contra esta DoD — itens ❌ são bloqueantes para merge
- **QA Lead:** verifica T1–T6 e P4 (validação contra TEST_PLAN.md)
- **PO:** verifica F6 (critérios de aceite da feature)
- **PMO:** verifica P1–P7 (conclusão da fase)

---

## 7. O que NÃO faz parte da DoD

Estes itens são importantes mas NÃO são critério de DONE para uma task:

| Item | Onde é tratado |
|------|----------------|
| Deploy em staging/produção | `ENVIRONMENTS.md` + pipeline de CI/CD |
| Monitoramento e alertas configurados | `DEPLOYMENT.md` + DevOps |
| Treinamento de usuários | Plano de change management (fora do escopo técnico) |
| Documentação de usuário final | Material de treinamento (fora do escopo técnico) |
| Performance sob carga real | Testes de carga em staging (fase de homologação) |

---

## 8. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-08 | Criação inicial: DoD de TASK (28 critérios), FEATURE (7 critérios), FASE (7 critérios), uso com agentes de IA | Time Técnico |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: agile-ba-practices, acceptance-criteria.*
