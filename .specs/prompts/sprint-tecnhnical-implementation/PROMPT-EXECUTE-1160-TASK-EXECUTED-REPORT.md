# PROMPT-EXECUTE-1160-TASK-EXECUTED-REPORT

## Contexto

Este prompt executa a **fase de Relatório TASK-EXECUTED** do pacote de desenvolvimento — gera o documento `TASK-EXECUTED-{timestamp}-[feature].md` (evidência histórica da implementação, usada como corpo da PR). **Delega a execução** ao prompt atual `PROMPT-GENERATE-IMPLEMENTATION-REPORT` (mantido como executor), alimentando-o com os artefatos do ciclo.

**Princípios fundamentais:**

1. **Evidência para a PR:** o TASK-EXECUTED é o anexo oficial da Pull Request (Fase 1150).
2. **Fonte consolidada:** o executor consome os `PACKAGE-DEVELOPMENT-*.md` do ciclo — nunca re-executa nada.
3. **Artefato distinto do 1100:** este é o relatório por demanda/feature (`TASK-EXECUTED`); o 1100 é o relatório do ciclo (`PACKAGE-DEVELOPMENT-EXECUTION-REPORT`).

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{CICLO_DIR}` | Pasta do ciclo | `.../sprints/sprint-01-setup/` |
| `{CICLO_NUMBER}` | Número do ciclo | `1` |
| `{CICLO_NAME}` | Nome curto do ciclo (kebab-case) | `sprint-01-setup` |

## Documentos de Referência

```
Ler obrigatoriamente (fontes do relatório):
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md  ← Relatório do ciclo (Fase 1100)
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-IMPLEMENTATION.md    ← Arquivos criados/modificados (Fase 1030)
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md ← Evidências de testes (Fase 1050)
    └── {CICLO_DIR}/PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md  ← Lint/qualidade (Fase 1060)
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1150-EXECUTION-REPORT`.

---

## Missão

Gerar o `TASK-EXECUTED-*.md` do ciclo `{CICLO_NUMBER} — {CICLO_NAME}` via `PROMPT-GENERATE-IMPLEMENTATION-REPORT`, e registrar o caminho no `PACKAGE-DEVELOPMENT-TASK-EXECUTED-REPORT.md`.

---

## Fluxo de Execução

1. **Consolidar** as fontes: execução (1100), arquivos (1030), testes (1050), qualidade (1060).
2. **Invocar `PROMPT-GENERATE-IMPLEMENTATION-REPORT`** (executor mantido) com os dados consolidados.
3. **Conferir a saída:** arquivo `{AAAA-MM-DD-HHMMSS}-TASK-EXECUTED-[feature].md` criado em `{SOLUTION_PATH}/.specs/skill-output/` com: resumo do desenvolvimento, arquivos modificados/criados, evidências de testes, validação de segurança, débitos resolvidos e surgidos.
4. **Registrar** o caminho no artefato da fase.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-TASK-EXECUTED-REPORT.md`:

```markdown
# PACKAGE-DEVELOPMENT-TASK-EXECUTED-REPORT.md — TASK-EXECUTED: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Arquivo Gerado
[Caminho completo do TASK-EXECUTED-*.md]
## 2. Conteúdo Resumido
- Desenvolvimento: [parágrafo]
- Arquivos: [X criados / Y modificados]
- Testes: [N/N passando, cobertura XX%]
## 3. Débitos
- Resolvidos: [DT-XXX...]
- Surgidos: [DT-XXX...]
## 4. Próximo Passo
[PR via PROMPT-EXECUTE-1180-PULL-REQUEST]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| Skills do executor | herdadas | `PROMPT-GENERATE-IMPLEMENTATION-REPORT` emprega: `code-documenter`, `gap-analysis`, `spec-miner`, `code-reviewer`, `fullstack-guardian` |
| `verification-before-completion` | automático | Conferir o TASK-EXECUTED gerado antes de concluir |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. NUNCA re-executar testes para o relatório — consolidar as evidências das fases.
2. Nome do arquivo segue estritamente o padrão do executor (`{AAAA-MM-DD-HHMMSS}-TASK-EXECUTED-[feature].md`).
3. Arquivo salvo em `{SOLUTION_PATH}/.specs/skill-output/` — nunca em diretórios globais.
4. Sem TASK-EXECUTED não há PR (Fase 1150 depende dele).
