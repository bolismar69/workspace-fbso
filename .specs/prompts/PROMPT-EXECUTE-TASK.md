# Prompt: Executar Tarefa de Desenvolvimento

- **Versão:** 2.0.0
- **Atualizado:** 2026-07-07
- **Template de artefato:** `.specs/modelos/MODELO-GENERATE-IMPLEMENTATION-REPORT.md`
- **Templates de report detalhados:** `analise-skills-caveman-ponytail-golang.md` — Seção 10

---

## Skills ativas obrigatórias

| Skill | Modo | Função |
|---|---|---|
| `caveman` | `full` | Compressão de prosa interativa (comunicação durante o desenvolvimento) |
| `ponytail` | `full` | Escada YAGNI de 7 rungs — controle de escopo do código gerado |
| `golang-pro` | automático | Ativado ao detectar código Go — padrões idiomáticos, qualidade, constraints |

- ⚠️ **caveman e ponytail NÃO atuam sobre artefatos permanentes** (SPECS.md, TASKS.md, reports). O caveman comprime apenas a comunicação interativa; o report final é gerado em prosa normal, como commits e PRs.

---

## Documentos de referência (INPUT — carregar antes de iniciar)

### Documentos de especificação (obrigatórios)
- `.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/SPECS.md`
- `.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TASKS.md`
- `.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TEST_PLAN.md`
- `.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/ARCHITECTURE.md`
- `.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/PRD.md`

### Documentos globais (obrigatórios)
- `.specs/security/SECURITY.md`
- `.specs/architecture/architecture.md`

### Documentação técnica da aplicação
- `.specs/` — documentação técnica completa (base de verdade atual da solução)
- `README.md` — instruções genéricas para execução e testes da aplicação

### Reports upstream (se existirem)
- `.specs/reports/DEVELOPER_TASK_[T-PREDECESSOR]_REPORT.md` — tasks predecessoras concluídas
- `.specs/reports/CODE_REVIEW_TASK_[T-XXX]_REPORT.md` — review anterior da mesma área de código

---

## Missão

Implemente a **Fase [solicitar qual fase de TASKS.md]** seguindo estritamente as especificações contidas no `SPECS.md` e utilizando o `TASKS.md` como roteiro de execução passo a passo.

---

## Fluxo

### Fase 0 — Pré-implementação

1. **Carregar documentos de referência** — ler SPECS.md, TASKS.md, TEST_PLAN.md, SECURITY.md, ARCHITECTURE.md, PRD.md
2. **Ler reports upstream** — se existirem `DEVELOPER_TASK_*_REPORT.md` de tasks predecessoras, carregar para contexto
3. **Identificar task atual** — localizar a task no TASKS.md e os casos de teste correspondentes no TEST_PLAN.md
4. **Subir a escada ponytail (7 rungs)** — registrar a decisão de CADA rung:

| Rung | Pergunta | Ação |
|------|----------|------|
| 1 | Isso precisa existir? (YAGNI) | Se NÃO → justificar e pular |
| 2 | Já existe no codebase? | Se SIM → reusar, não reescrever |
| 3 | Stdlib do Go cobre? | Se SIM → usar stdlib |
| 4 | Native platform cobre? | Se SIM → usar feature nativa |
| 5 | Dependência já instalada resolve? | Se SIM → usar, nunca adicionar dep nova para poucas linhas |
| 6 | Dá pra ser uma linha? | Se SIM → uma linha |
| 7 | Só então: código mínimo que funciona | Escrever o mínimo |

### Fase 1 — Implementação

5. **Escrever interfaces primeiro** — contratos antes da implementação (golang-pro step 2)
6. **Implementar com constraints golang-pro:**
   - `context.Context` em todas as operações bloqueantes
   - Erros explícitos (sem naked returns)
   - Error wrapping com `%w`
   - Documentar todos os exports
   - `gofmt` + `golangci-lint`
7. **Rodar `go vet ./...`** antes de prosseguir
8. **Rodar `golangci-lint run`** e corrigir todos os issues

### Fase 2 — Testes (Desenvolvimento Orientado a Testes)

9. **Escrever table-driven tests com `-race`** para cada trecho de código ou rota criada
10. **Executar testes imediatamente** — referenciar TEST_PLAN.md para os casos de teste
11. **Confirmar coverage ≥ 80%**
12. **Confirmar race detector: clean**

### Fase 3 — Pós-implementação

13. **Executar `/ponytail-review`** no diff — validar escada reversa (7→1)
14. **Documentar shortcuts intencionais** com comentários `// ponytail:` (ceiling + upgrade path)
15. **Atualizar TASKS.md** — marcar task como `[x]` após 100% dos testes passarem

### Fase 4 — Geração de artefato (OBRIGATÓRIO)

16. **Gerar arquivo de report** em `.specs/reports/DEVELOPER_TASK_[T-XXX]_REPORT.md`

    O arquivo deve seguir o template detalhado na **Seção 10.1** de `analise-skills-caveman-ponytail-golang.md`, cobrindo:
    - Decisão de cada rung da escada ponytail
    - Código gerado (paths, interfaces)
    - Constraints golang-pro (status de cada um)
    - Qualidade (go vet, golangci-lint, testes, coverage, race detector, benchmarks)
    - Segurança (constraints do SECURITY.md)
    - Arquitetura (constraints do ARCHITECTURE.md)
    - Skipped items (ponytail output)
    - Catálogo de `// ponytail:` comments
    - Desvios das specs originais
    - Resumo caveman (2-3 linhas)

    **Alternativamente**, use o modelo simplificado em `.specs/modelos/MODELO-GENERATE-IMPLEMENTATION-REPORT.md` para reports de execução rápida.

---

## Protocolo de Testes (fallback)

Se durante a execução do `TEST_PLAN.md` algum teste falhar:

1. **Auto-Correção Autônoma:** Analise o stack trace, identifique se o problema está na lógica do código ou na estrutura do teste, e corrija.
2. **Tratamento de Loops:** Se o mesmo erro persistir por **3 tentativas**, PARE imediatamente.
3. **Registro de Impedimento:** Crie `.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/IMPEDIMENT-FASE-[N].md` com:
   - Teste que quebrou e mensagem de erro exata
   - O que foi tentado para corrigir
   - Suspeita do motivo (limitação arquitetural, ambiguidade na SPEC, etc.)
   - Propostas adicionais de solução (se houver)
4. **Questionamentos ao Humano:** Se aplicável, crie `QUESTIONS.md` com perguntas que auxiliarão na retomada.
5. **Alerta:** Notifique o usuário no chat e aguarde instruções antes de alterar qualquer outro arquivo.

---

## Protocolo de Checagem Pós-Implementação (Sanity Check)

Após concluir todas as tasks da fase:

1. **Limpeza:** Remover código comentado, `fmt.Println` de debug, arquivos temporários.
2. **Git Status:** Listar arquivos modificados ou criados.
3. **Localização:** Validar que os arquivos respeitam o `ARCHITECTURE.md` e `.specs/architecture/architecture.md`.
4. **Segurança:** Revisão final contra `.specs/security/SECURITY.md` — nenhuma regra violada.
5. **Evidência:** Confirmar que o `TASKS.md` está com todas as tasks da fase marcadas como `[x]`.
6. **Pronto para Commit:** Se tudo 100%, preparar mensagem resumindo os arquivos alterados.

---

## Output esperado

| Output | Descrição |
|---|---|
| Código | Interfaces → implementação → testes |
| `go vet ./...` | clean |
| `golangci-lint run` | clean |
| Testes | PASS com `-race` |
| Coverage | ≥ 80% |
| Resumo caveman | 2-3 linhas em prosa comprimida |
| **ARTEFATO** | `.specs/reports/DEVELOPER_TASK_[T-XXX]_REPORT.md` |

---

## Anti-padrões

| ❌ NÃO fazer | ✅ Fazer |
|---|---|
| "Constrói um sistema completo de [X]" | "Implementar task T-XXX do TASKS.md" |
| Implementar sem referenciar SECURITY.md | Sempre incluir SECURITY.md no contexto |
| "Usa a biblioteca [X] para resolver" | Deixar o agente decidir via escada ponytail |
| Não gerar artefato de saída | Artefato é **obrigatório**, não opcional |
| Usar caveman para escrever o report | Report é artefato permanente — prosa normal |
| Implementar tasks futuras antes das atuais | Uma task de cada vez, na ordem do TASKS.md |

---

**Comece agora lendo os documentos de referência, suba a escada ponytail para a primeira task, e execute. Forneça feedback de progresso a cada etapa concluída.**

---

🤖 *Prompt template version 2.0.0 — gerido pelo time técnico.*
