# PROMPT-EXECUTE-1150-EXECUTION-REPORT

## Contexto

Este prompt executa a **Fase de Geração do Relatório de Execução** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 9 — passo 24, OBRIGATÓRIA). Consolida o resultado do ciclo em um relatório completo: resumo, stack/skills, tasks, arquivos, evidências de testes, validação de segurança e arquitetura, desvios e próximos passos.

**Princípios fundamentais:**

1. **Obrigatório:** o ciclo não está encerrado sem este relatório.
2. **Evidência real:** todo comando e resultado registrado vem das fases anteriores (1010 a 1090).
3. **Honestidade:** desvios, dificuldades e decisões são documentados — o relatório é a memória do ciclo.

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
Ler obrigatoriamente (fases anteriores do ciclo — fonte das evidências):
    ├── PACKAGE-DEVELOPMENT-PRE-IMPLEMENTATION.md     (1010)
    ├── PACKAGE-DEVELOPMENT-PLANNING.md               (1020)
    ├── PACKAGE-DEVELOPMENT-IMPLEMENTATION.md         (1030)
    ├── PACKAGE-DEVELOPMENT-TEST-PLANNING.md          (1040)
    ├── PACKAGE-DEVELOPMENT-TEST-IMPLEMENTATION.md    (1050)
    ├── PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md     (1060)
    ├── PACKAGE-DEVELOPMENT-FAILURE-HANDLING.md       (1070, se existir)
    ├── PACKAGE-DEVELOPMENT-CODE-REVIEW.md            (1080)
    └── PACKAGE-DEVELOPMENT-POST-IMPLEMENTATION.md    (1090)
```

> ⚠️ Se qualquer artefato obrigatório das fases 1010–1090 não existir → **PARE** e sinalize ao orquestrador (1000).

---

## Missão

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md` — o relatório de execução do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`, consolidando as evidências das fases anteriores.

---

## Fluxo de Execução

1. **Carregar os artefatos das fases 1010–1090** (tabela acima).
2. **Consolidar** no relatório (estrutura obrigatória):

```markdown
# PACKAGE-DEVELOPMENT-EXECUTION-REPORT.md — Relatório de Execução: Ciclo {N}
[Header: solução, projeto, ciclo, stack detectada, data da execução, tasks executadas]

## 1. Resumo da Execução
- Tasks executadas: X/Y
- Tasks com sucesso: X
- Tasks com falha: 0
- Tempo total estimado: Xd (do SPRINT-CARD.md)
- Tempo total gasto: (preencher ao final)

## 2. Stack e Skills Utilizadas
- Stack detectada: [linguagem + framework + banco + infra]
- Fonte da stack: [parâmetro {STACK} | PRD.md | SPECS.md | inferido do README.md]
- Skills acionadas: [lista com justificativa]

## 3. Tasks Executadas
| ID | Tarefa | Status | Testes | Cobertura | Observações |
|:---|:---|:---:|:---:|:---:|:---|
| T-XXX | Descrição | ✅ | N/N passando | XX% | — |

## 4. Arquivos Criados ou Modificados
| Ação | Arquivo | Task | Descrição da Mudança |
|:---|:---|:---|:---|
| 🆕 | `src/main/.../NovoArquivo.ext` | T-XXX | Nova classe/função/módulo — propósito |

## 5. Evidências de Testes
- Comando de build: `[comando]` → [resultado]
- Comando de teste: `[comando]` → [resultado]
- Total de testes: N | Status: ✅ 100% PASS
- Cobertura: XX% (linhas), XX% (branches)
- Cenários do SPRINT-TEST-SUITE.md executados: N/N

## 6. Validação de Segurança
- [ ] Nenhuma credencial ou dado sensível hardcoded
- [ ] Queries parametrizadas (injection prevention)
- [ ] Controles de acesso implementados onde exigido
- [ ] Dados pessoais não expostos em logs/respostas
- [ ] Respostas de erro não expõem stack traces

## 7. Validação de Arquitetura
- [ ] Estrutura de diretórios segue ARCHITECTURE.md
- [ ] Convenções de nomenclatura respeitadas
- [ ] Padrões documentados nas ADRs seguidos

## 8. Desvios e Observações
- Desvios das specs e justificativa
- Decisões de design tomadas
- Comentários `// ponytail:` adicionados (atalhos intencionais)
- Achados Critical/High remanescentes pós-2º ciclo de review (se houver)
- Dificuldades encontradas e como foram resolvidas

## 9. Próximos Passos
- Tasks restantes no ciclo (se houver)
- Pré-requisitos para o próximo ciclo (conforme SPRINT-REVIEW.md)
- Recomendações para a review com o PO

## Rodapé
[Indicação de geração por IA, skills utilizados, data/hora da geração]
```

3. **Apresentar ao humano** para validação (`[STATUS: COMPLIANCE]` só com aprovação).

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Conferir evidências das fases antes de assinar o relatório |
| `caveman` | full | Comunicação interativa — NÃO atua no relatório (artefato permanente; prosa normal, conforme o original) |

---

## Regras de Ouro

1. Relatório OBRIGATÓRIO — ciclo sem relatório não está encerrado.
2. Todas as evidências vêm dos artefatos das fases anteriores (nunca re-executar testes para "confirmar").
3. Desvios documentados com honestidade — nada de mascarar falhas ou ajustes.
4. Artefato em `{CICLO_DIR}/`; após COMPLIANCE, libera a Fase 1110.
