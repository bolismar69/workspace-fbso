# PROMPT-EXECUTE-1110-FAILURE-HANDLING

## Contexto

Este prompt executa a **Fase de Tratamento de Falhas** do pacote (extraída do `PROMPT-EXECUTE-SPRINT-TASKS.md`, Fase 6). É invocado **condicionalmente** quando uma falha ocorre em qualquer fase de implementação/teste/qualidade: analisa, tenta auto-correção autônoma com limite de 3 tentativas e, se persistir, registra impedimento e pausa para decisão humana.

**Princípios fundamentais:**

1. **Limite rígido:** 3 tentativas de auto-correção — nunca loop infinito.
2. **Impedimento documentado:** falha persistente vira `IMPEDIMENT-SPRINT-{N}.md` com evidência completa.
3. **Humano no controle:** após o impedimento, nenhuma alteração adicional acontece sem instrução humana.

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
| `{FAILURE}` | Falha a tratar: task, mensagem de erro exata e fase de origem | `T-003 — NullPointerException em CalcService.java:42 (Fase 1030)` |
| `{ATTEMPTS}` | Tentativas já realizadas (0 a 3) | `0` |

## Documentos de Referência

```
Ler conforme a falha:
    ├── {CICLO_DIR}/SPRINT-CARD.md          ← Critérios DONE da task
    ├── SPECS_DIR/SPECS.md + TASKS.md       ← Regras de negócio e descrição da task
    ├── SPECS_DIR/ARCHITECTURE.md           ← Padrões esperados
    └── logs de build/teste (saída do comando que falhou)
```

---

## Missão

Tratar a falha `{FAILURE}` do ciclo `{CICLO_NUMBER} — {CICLO_NAME}`: auto-correção autônoma (máx. 3 tentativas) ou, se persistir, registro do impedimento e pausa para decisão humana.

---

## Fluxo de Execução

1. **Auto-Correção Autônoma (até 3 tentativas):**
   - Analisar o stack trace/saída de erro
   - Identificar se o problema está na lógica do código ou na estrutura do teste
   - Corrigir de forma cirúrgica e reexecutar o comando que falhou
2. **Tratamento de Loops:** se o MESMO erro persistir por 3 tentativas → PARE imediatamente (proibido tentar variações infinitas).
3. **Registro de Impedimento:** criar `{CICLO_DIR}/IMPEDIMENT-SPRINT-{CICLO_NUMBER}.md` com:
   - Task que falhou (T-XXX) e mensagem de erro exata
   - O que foi tentado para corrigir (as 3 tentativas)
   - Suspeita do motivo (limitação arquitetural, ambiguidade na SPECS.md, conflito de dependências, etc.)
   - Propostas adicionais de solução (se houver)
4. **Notificar o humano** e aguardar instruções antes de alterar qualquer outro arquivo.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-FAILURE-HANDLING.md` (registro da tratativa) e, quando aplicável, `{CICLO_DIR}/IMPEDIMENT-SPRINT-{N}.md`:

```markdown
# PACKAGE-DEVELOPMENT-FAILURE-HANDLING.md — Falhas: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Falhas Encontradas
| Task | Fase de origem | Mensagem exata | Tentativas | Desfecho |
|:---|:---|:---|:---:|:---|
| T-003 | 1040-IMPLEMENTATION | NullPointerException ... | 1 | ✅ corrigida |
| T-007 | 1100-QUALITY-VALIDATION | coverage 74% < 80% | 3 | ❌ impedimento (ver IMPEDIMENT-SPRINT-{N}.md) |
## 2. Correções Aplicadas
[Tabela: arquivo | correção | revalidação]
## 3. Impedimentos Abertos
[Referência ao IMPEDIMENT-SPRINT-{N}.md + decisão humana pendente]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Revalidar a correção antes de declarar a falha tratada |
| `caveman` | full | Comunicação interativa na notificação ao humano |

> Correção cirúrgica direta — sem skills de stack nesta fase.

---

## Regras de Ouro

1. Máximo de 3 tentativas por falha — PARE e registre impedimento (nunca loop infinito).
2. Correções cirúrgicas: apenas o necessário para resolver a falha.
3. Impedimento sempre com mensagem de erro exata, tentativas e suspeita.
4. Após impedimento: nenhum outro arquivo é alterado sem instrução humana.
5. Desfecho de toda falha registrado no artefato (corrigida OU impedida).
