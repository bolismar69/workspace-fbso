# Contexto:
  - Este prompt é **genérico** — adapta-se a qualquer demanda: projeto, feature, issue, incidente, PoC, hotfix, etc.
  - Antes de executar, o humano deve fornecer os parâmetros da seção `# ⚙️ Parâmetros de Entrada`.
  - A documentação técnica em `.specs/` é a base de verdade atual da solução.
  - Use o README.md da raiz da solução sistêmica para instruções de build, execução e testes.

---

# ⚙️ Parâmetros de Entrada (preencher antes de executar)

> **Instrução:** No momento de invocar este prompt, o humano deve informar os valores abaixo. Se algum arquivo não existir, marcar como `N/D` (não disponível) e pular as etapas que o referenciam.

| Parâmetro | Descrição | Observação ou Exemplo |
|---|---|---|
| `{DEMAND_ID}` | Identificador único da demanda | `PRJ-FIN-2026-0001`, `ISSUE-422`, `INC-2026-0091`, `POC-ML-001`, , `HOTFIX-158499` |
| `{DEMAND_TYPE}` | Tipo da demanda | Lista de possibilidades: `business-projects`, `features`, `issues`, `incidents`, `hotfixs`, `pocs`, podendo se indicar opção diferente dessas |
| `{DEMAND_TITLE}` | Título descritivo curto | `Reforma Tributária 2026 Corporativo` |
| `{SOLUTION_ROOT}` | Caminho absoluto da raiz da solução sistêmica | `/home/user/work/ms-billing-engine-tax-rates` |
| `{SPECS_DIR}` | Caminho relativo da pasta de especificações (a partir de `{SOLUTION_ROOT}`) | `.specs` |
| `{SPECS_FILE}` | Caminho para o arquivo de especificações | `{SPECS_DIR}/{DEMAND_TYPE}/{DEMAND_ID}*/SPECS.md` |
| `{TASKS_FILE}` | Caminho para o arquivo de tasks/checklist | `{SPECS_DIR}/{DEMAND_TYPE}/{DEMAND_ID}*/TASKS.md` |
| `{TEST_PLAN_FILE}` | Caminho para o plano de testes | `{SPECS_DIR}/{DEMAND_TYPE}/{DEMAND_ID}*/TEST_PLAN.md` |
| `{ARCHITECTURE_FILE}` | Caminho para o documento de arquitetura | `{SPECS_DIR}/{DEMAND_TYPE}/{DEMAND_ID}*/ARCHITECTURE.md` |
| `{SECURITY_FILE_GLOBAL}` | Caminho para regras de segurança globais | `/home/user/work/.specs/security/SECURITY.md` |
| `{SECURITY_FILE_PROJECT}` | Caminho para regras de segurança do projeto (se existir) | `{SPECS_DIR}/security/SECURITY.md` |
| `{SKILL_OUTPUT_DIR}` | Pasta onde o relatório de execução será salvo | `{SPECS_DIR}/skill-output` |
| `{LANGUAGE}` | Linguagem/framework principal | `go`, `java`, `python`, `typescript`, `rust` |
| `{TEST_COMMAND}` | Comando para rodar a suíte de testes | `go test -count=1 ./... && go vet ./...` |
| `{BUILD_COMMAND}` | Comando para build de verificação | `go build ./...` |
| `{LINT_COMMAND}` | Comando para verificação estática (se diferente do test command) | `go vet ./...` |

---

# Missão:
> Processar as TASKS da demanda `{DEMAND_ID}` (`{DEMAND_TITLE}`), executando os itens do checklist `{TASKS_FILE}`, gerando ao final um documento de execução (TASK-EXECUTED) com evidências de testes e sumário das alterações.

---

# 🔁 Protocolo de Execução (Passo a Passo)

## Passo 1 — Leitura do Cenário
1. Confirmar que todos os parâmetros de entrada foram preenchidos.
2. Ler `{SPECS_FILE}` para entender o escopo da demanda (se `N/D`, pular).
3. Ler `{TASKS_FILE}` para identificar quais checkboxes ainda estão `[ ]` (não concluídos).
4. Ler `{TEST_PLAN_FILE}` para conhecer os cenários de teste exigidos (se `N/D`, definir testes baseados nas regras de negócio da `{SPECS_FILE}`).
5. Ler `{ARCHITECTURE_FILE}` para respeitar a estrutura de diretórios e padrões de design (se `N/D`, inferir dos padrões existentes no código).
6. Ler `{SECURITY_FILE_GLOBAL}` para conhecer as regras de segurança aplicáveis.

## Passo 2 — Execução Sequencial
1. Execute **um item de cada vez** do checklist `{TASKS_FILE}`.
2. **Não pule etapas.** Não implemente itens futuros antes de concluir os anteriores.
3. Para cada item implementado:
   - Escreva o código seguindo os padrões do projeto (DDD, interfaces, injeção de dependência, ou o padrão vigente em `{ARCHITECTURE_FILE}`).
   - Consulte `{TEST_PLAN_FILE}` e escreva os testes correspondentes **imediatamente**.
   - Execute `{TEST_COMMAND}` a cada item concluído.
   - Só marque o checkbox `[✅]` no `{TASKS_FILE}` após **100% de testes passando**.

## Passo 3 — Tratamento de Falhas
Se algum teste falhar:
1. **Auto-Correção Autônoma** (até 3 tentativas): Analise stack trace, corrija código ou teste.
2. **Loops Infinitos**: Se falhar 3x seguidas no mesmo erro, **PARE**.
3. **Registro de Impedimento**: Crie `{SPECS_DIR}/{DEMAND_TYPE}/{DEMAND_ID}*/IMPEDIMENT.md` com:
   - Item que falhou + mensagem de erro exata.
   - O que foi tentado para corrigir.
   - Suspeita da causa raiz.
   - Propostas alternativas de solução.
4. **Questionamentos**: Se aplicável, crie `{SPECS_DIR}/{DEMAND_TYPE}/{DEMAND_ID}*/QUESTIONS.md` com perguntas para retomada.
5. **Alerta**: Notifique o humano e aguarde instruções.

## Passo 4 — Sanity Check (Pós-Implementação)
Após TODOS os itens concluídos:
1. **Limpeza**: Remova código comentado, prints de debug, arquivos temporários.
2. **Git Status**: Liste apenas arquivos modificados/criados. Valide contra `{ARCHITECTURE_FILE}`.
3. **Security Scan**: Revise cada arquivo alterado contra `{SECURITY_FILE_GLOBAL}` (e `{SECURITY_FILE_PROJECT}` se existir):
   - [ ] Nenhuma credencial hardcoded.
   - [ ] Inputs sanitizados via schemas/validadores.
   - [ ] `{LINT_COMMAND}` — zero warnings.
   - [ ] RBAC aplicado em endpoints sensíveis (se aplicável à `{LANGUAGE}`).
   - [ ] Princípio do menor privilégio.
4. **Evidência**: Confirme que `{TASKS_FILE}` está com todos os checkboxes `[✅]`.
5. **Build Check**: Execute `{BUILD_COMMAND}` (e build de container se aplicável).

## Passo 5 — Geração do Documento de Execução (TASK-EXECUTED)
1. **Nome do arquivo**: `AAAA-MM-DD-HHMMSS_[slug-da-demanda-em-kebab-case].md`
   - Exemplo: `2026-06-25-073636-reforma-tributaria-fases-0-1-2.md`
   - Use o timestamp **real do momento da criação** (comando: `date +%Y-%m-%d-%H%M%S`).
2. **Local**: `{SKILL_OUTPUT_DIR}/`
3. **Conteúdo mínimo obrigatório**:

````markdown
# 📑 Relatório de Execução de Tarefa (TASK-EXECUTED)

* **Data e Hora da Conclusão:** AAAA-MM-DD HH:MM:SS (GMT-3)
* **Skills utilizadas:** [listar skills]
* **Demanda:** {DEMAND_ID} — {DEMAND_TITLE}
* **Tipo:** {DEMAND_TYPE}
* **Fase/Escopo/Feature/Issue Concluído:** [descrever]

---

## 🛠️ 1. Resumo do Desenvolvimento Realizado
[Parágrafo técnico denso: principais artefatos criados/alterados. Citar linguagem ({LANGUAGE}).]

## 🗂️ 2. Arquivos Modificados ou Criados
| Ação | Arquivo | Mudança |
|:---|:---|:---|
| 🆕 / 🔄 | `caminho/relativo` | Descrição concisa |

## 🧪 3. Evidências e Resultados dos Testes ({TEST_PLAN_FILE})
* **Comando Executado:** `{TEST_COMMAND}`
* **Total de Testes Rodados:** [N]
* **Status Final:** 🟩 100% PASSOU / 🟥 FALHOU
* **Saída Sumarizada do Terminal:**
  ```text
  [colar output real dos testes]
  ```

### Cobertura de Testes por Item
| Item | Descrição | Cenários de Teste | Testes Implementados |
|:---|:---|:---|:---|
| [ID] | [nome] | [N] cenários | N ✅ |

## 🔒 4. Validação de Segurança e Qualidade ({SECURITY_FILE_GLOBAL})
* [✅/❌] Nenhuma credencial hardcoded.
* [✅/❌] Inputs sanitizados.
* [✅/❌] `{LINT_COMMAND}` — zero warnings.
* [✅/❌] RBAC implementado em endpoints sensíveis (se aplicável).
* [✅/❌] Rate limiting com anti-spoof (se aplicável).
* [✅/❌] Métricas protegidas (se aplicável).
* [✅/❌] Princípio do Menor Privilégio.

## Documentação Atualizada
| Documento | Atualização |
|:---|:---|

## Dívidas Técnicas Resolvidas
| ID | Descrição | Item relacionado |
|:---|:---|:---|

## Dívidas Técnicas Remanescentes
| ID | Descrição | Impacto |
|:---|:---|:---|

---
🤖 *Documentação gerada de forma automatizada pelo agente de desenvolvimento de IA.*
````

4. **Validação final**: O arquivo deve conter TODAS as seções acima preenchidas.

## Passo 6 — Confirmação de Prontidão
Emita uma mensagem curta resumindo:
- Demanda `{DEMAND_ID}` concluída.
- Tipo: `{DEMAND_TYPE}`.
- Número de itens implementados.
- Total de testes passando.
- Localização do arquivo TASK-EXECUTED gerado.
- Status: **pronto para commit e pull request**.

---

# 📋 Exemplo de Invocação

> **Humano:** "Executar o Prompt #1 com os parâmetros:"
>
> - `{DEMAND_ID}` = `PRJ-FIN-2026-0001`
> - `{DEMAND_TYPE}` = `business-projects`
> - `{DEMAND_TITLE}` = `Reforma Tributária 2026 Corporativo`
> - `{SOLUTION_ROOT}` = `/home/user/work/ms-billing-engine-tax-rates`
> - `{SPECS_DIR}` = `.specs`
> - `{SPECS_FILE}` = `.specs/business-projects/PRJ-FIN-2026-0001-*/SPECS.md`
> - `{TASKS_FILE}` = `.specs/business-projects/PRJ-FIN-2026-0001-*/TASKS.md`
> - `{TEST_PLAN_FILE}` = `.specs/business-projects/PRJ-FIN-2026-0001-*/TEST_PLAN.md`
> - `{ARCHITECTURE_FILE}` = `.specs/business-projects/PRJ-FIN-2026-0001-*/ARCHITECTURE.md`
> - `{SECURITY_FILE_GLOBAL}` = `.specs/security/SECURITY.md`
> - `{SECURITY_FILE_PROJECT}` = `N/D`
> - `{SKILL_OUTPUT_DIR}` = `.specs/skill-output`
> - `{LANGUAGE}` = `go`
> - `{TEST_COMMAND}` = `go test -count=1 ./... && go vet ./...`
> - `{BUILD_COMMAND}` = `go build ./...`
> - `{LINT_COMMAND}` = `go vet ./...`

---

# Skills Recomendadas
`golang-pro` (se Go), `java-architect` (se Java), `python-pro` (se Python), `typescript-pro` (se TypeScript), `code-documenter`, `gap-analysis`, `spec-miner`, `code-reviewer`, `fullstack-guardian`, `test-driven-development`, `verification-before-completion`
