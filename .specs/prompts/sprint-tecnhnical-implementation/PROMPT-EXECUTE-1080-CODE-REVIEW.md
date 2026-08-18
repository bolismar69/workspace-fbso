# PROMPT-EXECUTE-1080-CODE-REVIEW

## Contexto

Este prompt executa a **Fase de Code Review** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 7 — passos 14 a 23). Aciona as **7 auditorias** sobre o código do ciclo, consolida os achados, executa os ajustes e controla o **loop de retorno** à Fase de Testes (máx. 2 ciclos).

**Princípios fundamentais:**

1. **Sete lentes independentes:** ponytail-audit, ponytail-review, engineering-skills, security-audit, performance-review, requesting-code-review e differential-review.
2. **Consolidação única:** todos os achados agrupados por arquivo e severidade no relatório.
3. **Ajuste com revalidação:** cada correção é validada com build; falha reverte o ajuste.
4. **Loop controlado:** retorno à Fase 1040 (Test Planning) limitado a 2 ciclos completos.

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
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |
| `{LOOP_COUNT}` | Número do loop de retorno atual (1 ou 2) | `1` |

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md ← Resultado da qualidade (Fase 1060)
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-IMPLEMENTATION.md      ← Arquivos implementados (Fase 1030)
    ├── {CICLO_DIR}/SPRINT-CARD.md                             ← Critérios DONE
    ├── SPECS_DIR/ARCHITECTURE.md + SPECS.md                   ← Padrões e regras
    └── {SOLUTION_PATH}/.specs/security/SECURITY.md            ← Regras de segurança
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-QUALITY-VALIDATION.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1060-QUALITY-VALIDATION`.

---

## Missão

Auditar o código do ciclo `{CICLO_NUMBER} — {CICLO_NAME}` com as 7 auditorias, consolidar achados, executar os ajustes e registrar o resultado no `PACKAGE-DEVELOPMENT-CODE-REVIEW.md` (+ `SPRINT-CODE-REVIEW-{fase}.md` quando houver achados).

---

## Fluxo de Execução

1. **Auditoria `ponytail-audit`:** código morto (YAGNI), duplicação (DRY), dependências desnecessárias, desvios do ARCHITECTURE.md, complexidade, código não idiomático → findings com severidade.
2. **Revisão `ponytail-review`:** qualidade/legibilidade, aderência ao ARCHITECTURE.md, segurança (SECURITY.md), bordas não tratadas, consistência.
3. **Auditoria `engineering-skills`:** SOLID/DRY/KISS, coesão/acoplamento, padrões, eficiência, tratamento de erros, qualidade dos testes.
4. **Auditoria `security-audit`:** OWASP Top 10, exposição de dados, autorização/autenticação, configurações inseguras, validação de entrada, criptografia.
5. **Revisão `performance-review`:** N+1, alocações, bloqueios em fluxos assíncronos, caching, complexidade algorítmica, pools.
6. **Revisão `requesting-code-review`:** legibilidade, convenções, testabilidade, edge cases, simplificação, documentação.
7. **Revisão `differential-review`** (sobre o `git diff` do ciclo): regressões de segurança, blast radius, cobertura das linhas alteradas, quebra de contratos/schema.

8. **Consolidar achados:**
   - SE houver findings: gerar `{CICLO_DIR}/SPRINT-CODE-REVIEW-{fase}.md` (estrutura do original: resumo por severidade, seções por auditoria com IDs PA/PR/ES/SA/PF/RC/DR, plano de ajustes §9, execução §10, build pós-ajustes §11) e prosseguir ao passo 9.
   - SENÃO (zero achados relevantes): pular para a conclusão do artefato.

9. **Executar ajustes** a partir do plano (§9), na ordem:
   - Aplicar a correção → marcar concluído (§10) → executar build → falhou? reverter, marcar ❌ e documentar.

10. **Controle do loop de retorno:**
    - Após os ajustes: retornar ao `PROMPT-EXECUTE-1040-TEST-PLANNING` (o plano de testes pode precisar de atualização) e reexecutar 1040 → 1050 → 1060 → 1080.
    - **Máximo de 2 ciclos completos.** Após o 2º ciclo, achados Critical/High não resolvidos são registrados no relatório de execução (§8 do 1100) e o fluxo prossegue; Medium/Low viram débito técnico documentado.

> ⚠️ **Por que voltar à Fase de Testes?** Ajustes de code review podem alterar assinaturas, introduzir classes, remover código morto ou mudar fluxos de exceção — os testes precisam ser revalidados para manter a qualidade.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-CODE-REVIEW.md` (+ `SPRINT-CODE-REVIEW-{fase}.md` quando houver achados):

```markdown
# PACKAGE-DEVELOPMENT-CODE-REVIEW.md — Code Review: Ciclo {N}
[Header: solução, projeto, ciclo, stack, data, loop {LOOP_COUNT}/2]
## 1. Resumo da Revisão
- Auditorias acionadas: [7 skills]
- Total de achados: N
| Critical | High | Medium | Low |
|:---:|:---:|:---:|:---:|
| X | Y | Z | W |
## 2. Achados Consolidados
[Referência ao SPRINT-CODE-REVIEW-{fase}.md (seções por auditoria)]
## 3. Ajustes Executados
| ID | Arquivo | Ação | Resultado |
|:---|:---|:---|:---:|
| SA-001 | Controller.java | @RequiresPermission adicionado | ✅ Compila |
## 4. Loop de Retorno
- Retorno à 1040: [sim/não — nº do loop]
- Achados Critical/High restantes: [lista ou "nenhum"]
- Débitos Medium/Low documentados: [lista]
## 5. Veredito
[✅ REVIEW APROVADA (prosseguir à 1090) | ↩️ RETORNAR À 1040 (loop {n}/2)]
## Rodapé
[Indicação de geração por IA, skills utilizados, data/hora]
```

---

## Skills

### As 7 auditorias da Fase 7 do original

| Skill | Foco |
|:---|:---|
| `ponytail-audit` | Código morto (YAGNI), duplicação (DRY), dependências desnecessárias, desvios do ARCHITECTURE.md, complexidade, código não idiomático |
| `ponytail-review` | Legibilidade, aderência ao ARCHITECTURE.md, segurança (SECURITY.md), bordas não tratadas, consistência |
| `engineering-skills` | SOLID/DRY/KISS, coesão/acoplamento, padrões, eficiência, tratamento de erros, qualidade dos testes |
| `security-audit` | OWASP Top 10, exposição de dados, authz/authn, configurações inseguras, validação de entrada, criptografia |
| `performance-review` | Queries N+1, alocações, bloqueios em fluxos assíncronos, caching, complexidade algorítmica, pools |
| `requesting-code-review` | Legibilidade, convenções, testabilidade, edge cases, simplificação, documentação |
| `differential-review` | Regressões de segurança no diff, blast radius, cobertura das linhas alteradas, quebra de contratos/schema |

### Transversais

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `code-review` | automático | Transversal de qualidade (reforço das 7) |
| `security-review` | automático | Transversal de segurança |
| `verification-before-completion` | automático | Build verde pós-ajuste antes de concluir |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. As 7 auditorias são obrigatórias — nenhuma pode ser pulada.
2. Achados sem severidade/arquivo/linha/recomendação são inválidos.
3. Correção só é concluída com build verde; falha reverte o ajuste.
4. Loop de retorno limitado a 2 ciclos; excedente = registrar no relatório e prosseguir (Critical/High) ou débito técnico (Medium/Low).
5. Relatórios sempre em `{CICLO_DIR}/`; decisão final é do humano.
