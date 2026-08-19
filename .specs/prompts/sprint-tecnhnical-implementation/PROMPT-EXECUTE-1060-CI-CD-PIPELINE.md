# PROMPT-EXECUTE-1060-CI-CD-PIPELINE

## Contexto

Este prompt executa a **fase de Implementação do Pipeline CI/CD** do pacote de desenvolvimento (equivalente ao step 3b do Bloco E do TECHLEAD). **Delega a execução** ao prompt especialista `PROMPT-EXECUTE-CI-CD-PIPELINE` (mantido como executor), que materializa os planos `041/087/090` + definições `500/510/086` em pipelines concretos.

**Princípios fundamentais:**

1. **Opcional e condicionado:** roda somente quando `500/087` exigirem pipeline para a solução no ciclo.
2. **Ancoragem documental:** o executor só cria etapas previstas em 041/087/090/510/086 (regra do próprio especialista).
3. **Gate:** pipeline implementado passa pelo `PROMPT-QA-REVISOR-SECURITY` (via Fase 1120) + validação humana antes da PR.

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
| `{CI_PROVIDER}` | Provedor de CI/CD (do 500) | `github-actions` ou `gitlab-ci` |
| `{PROJECT_DOCS_DIR}` | Pasta dos documentos WATERFALL (041/086/087/090) | `.../PRJ-TEC-2026-0004-PROJETO-SHIELD/` |
| `{TECH_DEFS_DIR}` | Pasta das definições TECHLEAD (480/500/510/520/550) | `.../technical-definitions/` |

## Documentos de Referência

```
Ler obrigatoriamente (delegação — o executor consome os mesmos):
    ├── {PROJECT_DOCS_DIR}/041-DEVOPS-SETUP, 087-PLANO-CI-CD-AMBIENTES, 090-STRATEGIC-IMPLEMENTATION..., 086-PADROES-CODIGO
    └── {TECH_DEFS_DIR}/500-DEVOPS-SRE-DEFINITION, 510-TEST-STRATEGY-DEFINITION
```

---

## Missão

Implementar o pipeline CI/CD da solução `{SOLUTION_NAME}` via `PROMPT-EXECUTE-CI-CD-PIPELINE` e registrar o resultado no `PACKAGE-DEVELOPMENT-CI-CD-PIPELINE.md`.

---

## Fluxo de Execução

1. **Condição de entrada:** confirmar nos docs-base (`500`/`087`) que o ciclo exige pipeline para a solução — se não exigir, registrar "não aplicável" e encerrar a fase.
2. **Invocar `PROMPT-EXECUTE-CI-CD-PIPELINE`** (executor mantido) com os parâmetros acima.
3. **Conferir a saída do executor:** arquivos de pipeline gerados (`{SOLUTION_PATH}/.github/workflows/` ou `.gitlab-ci.yml`), mapa `CICD-NN → arquivo → estágios`, validação de sintaxe e relatório TASK-EXECUTED.
4. **Gate:** submeter o pipeline ao `PROMPT-EXECUTE-1130-QUALITY-ASSURANCE-REVIEW` (HITL) e à validação humana.
5. **Registrar** no artefato da fase.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-CI-CD-PIPELINE.md`:

```markdown
# PACKAGE-DEVELOPMENT-CI-CD-PIPELINE.md — Pipeline CI/CD: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Aplicável?
[Sim — exigido por 500/087 | Não — registrado]
## 2. Pipelines Implementados
| CICD-NN | Arquivo | Estágios | Ambiente | Gate |
|:---|:---|:---|:---|:---|
## 3. Validação
[Sintaxe ✅ | actionlint ✅ | gate 1120: APPROVED]
## 4. Rastreabilidade
[CICD-NN → 041/087/090/510 §origem]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Pipeline validado antes de concluir |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

> `PROMPT-EXECUTE-CI-CD-PIPELINE` é um **prompt especialista executor**, não skill — ele próprio aciona as skills de DevOps que julgar necessárias.

---

## Regras de Ouro

1. Roda somente quando 500/087 exigirem — nunca "por boa prática".
2. Pipeline sem ancoragem em 041/087/090/510 é violação (regra do executor).
3. Gate 1120 (HITL) obrigatório antes da PR.
4. Nenhum deploy automático em PROD sem approval gate.
