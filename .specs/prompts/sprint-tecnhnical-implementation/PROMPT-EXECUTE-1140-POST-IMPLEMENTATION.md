# PROMPT-EXECUTE-1140-POST-IMPLEMENTATION

## Contexto

Este prompt executa a **Fase de Pós-implementação (Sanity Check)** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 8 — passos 19 a 23). Faz a varredura final antes do relatório: limpeza, estado do git, conformidade com a arquitetura, revisão de segurança e atualização do backlog do ciclo.

**Princípios fundamentais:**

1. **Ciclo limpo:** sem código comentado, prints de debug, imports não usados ou arquivos temporários.
2. **Conformidade verificada:** estrutura e convenções conforme ARCHITECTURE.md; regras do SECURITY.md revalidadas.
3. **Backlog atualizado:** cada task marcada como concluída no SPRINT-CARD.md antes do relatório.

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
Ler obrigatoriamente:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-CODE-REVIEW.md  ← Veredito do review (Fase 1080)
    ├── {CICLO_DIR}/SPRINT-CARD.md                      ← Backlog do ciclo (marcar tasks)
    ├── SPECS_DIR/ARCHITECTURE.md                       ← Estrutura e convenções
    └── {SOLUTION_PATH}/.specs/security/SECURITY.md     ← Regras de segurança
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-CODE-REVIEW.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1120-CODE-REVIEW`.

---

## Missão

Executar o sanity check final do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`: limpeza, git status, validação de arquitetura e segurança, e atualização do SPRINT-CARD.md — registrando no `PACKAGE-DEVELOPMENT-POST-IMPLEMENTATION.md`.

---

## Fluxo de Execução

1. **Limpeza:** remover código comentado, prints de debug, imports não usados e arquivos temporários.
2. **Git Status:** listar todos os arquivos modificados ou criados (`git status --short`).
3. **Localização:** validar que os arquivos respeitam o `ARCHITECTURE.md` (estrutura de diretórios, nomes de arquivo, convenções).
4. **Segurança:** revisão final contra `SECURITY.md` — nenhuma regra violada. Verificar:
   - Nenhuma credencial ou segredo hardcoded
   - Queries parametrizadas (prevenção de SQL injection)
   - Autorização implementada onde exigido (RBAC, scopes)
   - Dados sensíveis não expostos em logs ou respostas de erro
5. **Atualizar SPRINT-CARD.md:** marcar cada task concluída no backlog do ciclo.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-POST-IMPLEMENTATION.md`:

```markdown
# PACKAGE-DEVELOPMENT-POST-IMPLEMENTATION.md — Pós-implementação: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Limpeza
[Itens removidos: comentários, prints, imports, temporários]
## 2. Estado do Git
[Lista `git status --short` dos arquivos do ciclo]
## 3. Validação de Arquitetura
| Item | Regra (ARCHITECTURE.md) | Status |
|:---|:---|:---:|
| Estrutura de diretórios | §2 | ✅ |
| Nomenclatura | §convenções | ✅ |
## 4. Revisão Final de Segurança
- [ ] Nenhuma credencial hardcoded
- [ ] Queries parametrizadas
- [ ] Autorização implementada onde exigido
- [ ] Dados sensíveis não expostos
## 5. SPRINT-CARD Atualizado
[Tabela: task | status ✅/❌ | observações]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `security-review` | automático | Revisão final de segurança contra SECURITY.md |
| `verification-before-completion` | automático | Checklist completo antes de declarar a fase concluída |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. Nenhuma violação de SECURITY.md pode permanecer — corrigir ou escalar ao humano.
2. SPRINT-CARD.md é atualizado SEMPRE (tasks marcadas antes do relatório).
3. Estrutura fora do ARCHITECTURE.md = corrigir localização ou registrar desvio justificado.
4. Artefato em `{CICLO_DIR}/`; conclusão da fase libera a geração do relatório (1100).
