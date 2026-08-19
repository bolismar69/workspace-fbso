# PROMPT-EXECUTE-1070-TEST-PLANNING

## Contexto

Este prompt executa a **Fase de Planejamento dos Testes** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 3 — passo 10). Gera o plano de testes do ciclo ANTES de implementar os testes: quais cenários cobrir, com quais ferramentas, em qual ordem e com qual meta de cobertura. Deriva do SPRINT-TEST-SUITE.md (o que testar) e detalha o COMO testar.

**Princípios fundamentais:**

1. **Plano antes dos testes:** nenhum teste é escrito sem este artefato.
2. **Derivado, não inventado:** cenários vêm do SPRINT-TEST-SUITE.md e TEST_PLAN.md — o plano organiza e detalha, nunca cria cenários novos.
3. **Ações manuais explícitas:** todo cenário não automatizável vira instrução passo-a-passo acionável.

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

## Documentos de Referência

```
Ler obrigatoriamente:
    ├── {CICLO_DIR}/SPRINT-TEST-SUITE.md              ← Cenários de teste aplicáveis ao ciclo
    ├── {CICLO_DIR}/PACKAGE-DEVELOPMENT-IMPLEMENTATION.md ← Tasks implementadas (Fase 1030)
    ├── SPECS_DIR/TEST_PLAN.md                        ← Fonte da verdade dos cenários (IDs)
    ├── SPECS_DIR/ARCHITECTURE.md                     ← Padrões de testes por camada
    └── {SOLUTION_PATH}/README.md                     ← Comandos de execução de testes
```

> ⚠️ Se `PACKAGE-DEVELOPMENT-IMPLEMENTATION.md` não existir → **PARE**: execute primeiro o `PROMPT-EXECUTE-1040-IMPLEMENTATION`.

---

## Missão

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-TEST-PLANNING.md` — o plano de testes do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`, com mapeamento task→cenário, estratégia por nível, ordem de execução, comandos e ações manuais.

---

## Fluxo de Execução

1. **Carregar** SPRINT-TEST-SUITE.md, PACKAGE-DEVELOPMENT-IMPLEMENTATION.md, TEST_PLAN.md, ARCHITECTURE.md e README.md.
2. **Mapear** tasks implementadas → cenários de teste (IDs do TEST_PLAN preservados para rastreabilidade bidirecional).
3. **Gerar o artefato** com a estrutura obrigatória:

```markdown
# PACKAGE-DEVELOPMENT-TEST-PLANNING.md — Plano de Testes: Ciclo {N}
[Header: solução, projeto, ciclo, stack, data]

## 1. Visão Geral
- Tasks implementadas: X
- Cenários de teste mapeados: Y (do SPRINT-TEST-SUITE.md)
- Meta de cobertura: ≥ 80% (padrão)
- Ferramentas: [JUnit 5 + Mockito + Testcontainers | pytest | Jest | Go testing | ...]

## 2. Mapeamento Task → Cenários de Teste
| Task | Cenário(s) | Nível | Ferramenta | Status |
|:---|:---|:---|:---|:---:|
| T-XXX | TC-XXX-001 | Unit | [framework] | ⬜ |

## 3. Estratégia por Nível de Teste
### 3.1 Testes Unitários
[Ferramenta, padrão (AAA/Given-When-Then/table-driven), localização, o que mockar, o que NÃO mockar]
### 3.2 Testes de Integração
[Ferramenta (Testcontainers/Docker Compose), localização, o que usar real, dados de seed]
### 3.3 Testes de Segurança (se aplicável)
[Foco: RBAC, multi-tenant isolation, rate limiting, JWT validation; localização]

## 4. Ordem de Execução dos Testes
[1. Unitários → 2. Integração → 3. Segurança — com justificativa]

## 5. Comandos de Execução
- Unit: `[comando]` | Integration: `[comando]` | Coverage: `[comando]` | Lint/Quality: `[comando]`

## 6. Ações Manuais ou Externas
> Obrigatório quando aplicável (cenário não automatizável). Para cada ação:
### Ação X: [Título Descritivo]
- **Cenário(s) relacionado(s):** [IDs do SPRINT-TEST-SUITE.md]
- **Quem executa:** [Humano — papel | Sistema externo — nome/URL]
- **Pré-condições:** [o que precisa estar pronto]
- **Ambiente:** [dev | staging | CI | produção]
- **Passo a passo:** [1..N — verbos no imperativo, comandos copiáveis, URLs/portas explícitas]
- **Resultado esperado:** [o que observar para confirmar sucesso]
- **Se falhar:** [ação corretiva ou contato]
- **Evidência a coletar:** [ ] screenshot [ ] log [ ] resposta HTTP

## 7. Provenientes de Testes de Validação de Qualidade:
| Task | Mensagem exata | Suspeita | Proposta solução |

## 8. Provenientes de Code Review:
| Task | Mensagem exata | Suspeita | Proposta solução | Skills |

## Rodapé
[Indicação de geração por IA, data/hora]
```

4. **Validar:** todo cenário rastreia ao TEST_PLAN.md; IDs preservados; ações manuais têm instruções copiáveis.
5. **Apresentar ao humano** para validação.

> ⚠️ Artefato gerado UMA VEZ, após a implementação e antes dos testes. Se o escopo de testes mudar durante a execução, atualizar este arquivo.

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| Skills de teste da stack | herdadas | Referência do original (tabela 6.3): `131-java-testing-unit-testing`, `golang-testing`, `javascript-typescript-jest`, `pytest-coverage` — informam ferramentas e padrões por nível |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

---

## Regras de Ouro

1. NUNCA escrever teste sem este artefato aprovado.
2. Cenários derivados do SPRINT-TEST-SUITE/TEST_PLAN — nunca inventados.
3. IDs originais do TEST_PLAN.md preservados (rastreabilidade bidirecional).
4. Ações manuais com verbos no imperativo, comandos copiáveis e evidência definida.
5. Artefato em `{CICLO_DIR}/` com `[STATUS: Em análise]`; COMPLIANCE com validação humana.
