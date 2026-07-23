# PROMPT-FIX-PRD-FROM-GATE

## Contexto

Este prompt é acionado quando o **Gate de Alinhamento de Escopo** (PROMPT-GATE-PRD-SCOPE) reprova o `PRD.md` e gera o relatório `PRD_SCOPE_FAIL_REPORT.md`.

O agente atua como **corretor de escopo** — lê o relatório de falha, consome cada não-conformidade reportada e corrige o PRD.md para que ele passe no gate na próxima execução.

**Princípio fundamental:** Este prompt NÃO reescreve o PRD.md do zero. Ele aplica correções cirúrgicas baseadas no relatório de falha, preservando o conteúdo que já está correto.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 4 parâmetros foram informados.

### Passo 1 — Carregar Artefatos e Relatório de Falha

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD_SCOPE_FAIL_REPORT.md (relatório do gate)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md (artefato a ser corrigido)
    └── Documentos de referência em {PROJECT_PATH}:
          ├── 01-PROJECT-CHARTER-*.md
          ├── 02-BUSINESS-REQUIREMENTS.md
          ├── 03-EPICS.md
          ├── 04-FEATURES.md
          └── DEFINITION_OF_DONE.md

Se PRD_SCOPE_FAIL_REPORT.md não existir → ERRO: "Relatório de falha não encontrado. Execute o gate primeiro."
Se PRD.md não existir → ERRO: "PRD.md não encontrado."
```

### Passo 2 — Processar Não-Conformidades por Prioridade

Ler o `PRD_SCOPE_FAIL_REPORT.md` e extrair:

1. **Veredito geral** (§1 do relatório) — REPROVADO ou RESSALVA
2. **Dimensões reprovadas** (§2) — quais dimensões tiveram < 75%
3. **Não-conformidades detalhadas** (§3) — cada NC com ID, verificação, evidência e ação corretiva
4. **Itens de scope creep** (§4) — o que precisa ser removido
5. **Recomendações** (§5) — lista priorizada de correções

As correções devem seguir a ordem de prioridade:

| Prioridade | Tipo de NC | Ação |
|---|---|---|
| P0 (Crítica) | Scope creep confirmado (features/BRs sem origem nos docs de negócio) | **Remover** do PRD.md |
| P1 (Alta) | Feature/entrega do Project Charter não mapeada no PRD.md | **Adicionar** seção correspondente |
| P2 (Média) | Rastreabilidade incompleta (matriz BR→Feature→US) | **Completar** mapeamentos |
| P3 (Baixa) | Documentação incompleta (fora de escopo, premissas) | **Complementar** seções |

### Passo 3 — Aplicar Correções no PRD.md

Para cada não-conformidade no relatório, aplicar a correção correspondente:

```
Para cada NC no relatório (§3):
    │
    ├── NC é de REMOÇÃO (scope creep)?
    │     → Localizar a feature/BR/requisito no PRD.md
    │     → Remover ou mover para seção "Fora de Escopo"
    │     → Atualizar matriz de rastreabilidade
    │
    ├── NC é de ADIÇÃO (item faltante)?
    │     → Consultar o documento de negócio referenciado na ação corretiva
    │     → Adicionar seção/feature/BR no local apropriado do PRD.md
    │     → Atualizar matriz de rastreabilidade
    │
    ├── NC é de COMPLEMENTO (incompleto)?
    │     → Identificar a seção incompleta
    │     → Expandir com informações dos docs de negócio
    │     → Adicionar exemplos, casos de borda, referências
    │
    └── NC é de CORREÇÃO (inconsistente)?
          → Identificar a inconsistência (ex: prioridade diferente do BRD)
          → Alinhar com o documento de referência
          → Documentar a decisão se houver justificativa para divergência
```

### Passo 4 — Atualizar Registro de Alterações

Adicionar entrada no registro de alterações do PRD.md:

```markdown
| v{X+1} | {data} | Correção pós-gate: {N} não-conformidades resolvidas do PRD_SCOPE_FAIL_REPORT.md v{Y}. NCs: {lista de IDs} | Agente Corretor PRD/IA |
```

### Passo 5 — Validar Correções

Antes de considerar concluído, verificar:

| # | Verificação | Critério |
|---|---|---|
| 1 | Todas as NCs endereçadas | Cada NC do relatório tem uma correção aplicada ou justificativa documentada |
| 2 | Scope creep removido | Itens listados no §4 do relatório foram removidos ou movidos para "Fora de Escopo" |
| 3 | Adições referenciadas | Novas seções/features adicionadas referenciam os docs de negócio corretos |
| 4 | Matriz de rastreabilidade atualizada | Matriz BR→Feature→US reflete as correções |
| 5 | Registro de alterações | Nova versão documentada com referência ao relatório de falha |
| 6 | Consistência cross-documento | PRD.md corrigido é consistente com Project Charter, BRD, Épicos e Features |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `gap-analysis` | Analisar relatório de falha e planejar correções |
| 2ª | `requirements-validation` | Garantir que correções mantêm alinhamento com baseline |
| 3ª | `agile-ba-practices` | Boas práticas de BA na escrita do PRD.md |
| 4ª | `documentation-writer` | Qualidade e consistência do PRD.md corrigido |

---

## Observações

1. **Correções cirúrgicas, não reescrita.** O objetivo é resolver as não-conformidades apontadas no relatório, não reescrever o PRD.md. Se o relatório indicar que uma reescrita completa é necessária, usar o PROMPT-GENERATE-PRD-ARTEFACT com `SCOPE=full`.

2. **Scope creep confirmado deve ser removido.** Se o gate identificou features/BRs sem origem nos documentos de negócio, a ação padrão é REMOVER. Se houver justificativa de negócio para manter, documentar a decisão e solicitar atualização dos docs de negócio.

3. **O relatório de falha é o contrato.** Cada correção deve ser rastreável a uma NC específica do relatório. Não fazer alterações não solicitadas no PRD.md durante a correção.

4. **Incrementar versão SEMPRE.** Toda correção gera uma nova versão do PRD.md com entrada no registro de alterações referenciando o relatório de falha.

5. **Após correção, reexecutar o gate.** O fluxo natural é: FIX PRD → GATE PRD → (se APROVADO) prosseguir. Não considerar o PRD.md como aprovado até o gate confirmar.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: correção de PRD.md baseada em PRD_SCOPE_FAIL_REPORT.md, 4 níveis de prioridade, 6 verificações pós-correção | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, gap-analysis, requirements-validation.*
