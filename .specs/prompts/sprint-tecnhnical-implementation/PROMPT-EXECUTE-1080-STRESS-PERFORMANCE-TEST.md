# PROMPT-EXECUTE-1080-STRESS-PERFORMANCE-TEST

## Contexto

Este prompt executa a **fase de Teste de Estresse/Performance** do pacote de desenvolvimento (equivalente ao step 4a do Bloco E do TECHLEAD). **Delega a execução** ao prompt especialista `PROMPT-EXECUTE-STRESS-PERFORMANCE-TEST` (mantido como executor), que roda os cenários de performance contra os thresholds da `510` e os SLOs da `500`, com veredito objetivo.

**Princípios fundamentais:**

1. **Opcional e condicionado:** roda quando `510`/`050` tiverem cenários de performance para o ciclo.
2. **Veredito objetivo (do executor):** medição × threshold — não existe "passou com ressalva".
3. **Falha de SLO:** vira `DT-XXX` (via Fase 1130) + decisão humana (otimizar, renegociar SLO ou adiar janela).

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/go/fiber/microservices/ms-billing-engine-tax-rates` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-billing-engine-tax-rates` |
| `{CICLO_DIR}` | Pasta do ciclo | `.../sprints/sprint-01-setup/` |
| `{CICLO_NUMBER}` | Número do ciclo | `1` |
| `{TECH_DEFS_DIR}` | Pasta das definições TECHLEAD (500/510) | `.../technical-definitions/` |
| `{PROJECT_DOCS_DIR}` | Pasta dos documentos WATERFALL (045/050) | `.../PRJ-TEC-2026-0004-PROJETO-SHIELD/` |
| `{PERF_TOOL}` | Ferramenta (se omitido, a da 510) | `k6`, `jmeter`, `gatling` |
| `{TARGET_ENV}` | Ambiente alvo do teste | `QA` |
| `{SCENARIOS}` | Casos de performance (IDs do 050; vazio = os mapeados na 510) | `TC-PERF-001,TC-PERF-002` |

## Documentos de Referência

```
Ler obrigatoriamente (delegação — o executor consome os mesmos):
    ├── {TECH_DEFS_DIR}/510-TEST-STRATEGY-DEFINITION, 500-DEVOPS-SRE-DEFINITION
    └── {PROJECT_DOCS_DIR}/045-TEST-PLAN, 050-TEST-CASES
```

---

## Missão

Executar os testes de estresse/performance da solução `{SOLUTION_NAME}` via `PROMPT-EXECUTE-STRESS-PERFORMANCE-TEST` e registrar o veredito no `PACKAGE-DEVELOPMENT-STRESS-PERFORMANCE-TEST.md`.

---

## Fluxo de Execução

1. **Condição de entrada:** confirmar em `510`/`050` que há cenários de performance para o ciclo — se não houver, registrar "não aplicável" e encerrar.
2. **Invocar `PROMPT-EXECUTE-STRESS-PERFORMANCE-TEST`** (executor mantido) com os parâmetros acima.
3. **Conferir a saída do executor:** relatório com indicador × threshold (p50/p95/p99, erro, burn rate) e veredito `[STATUS: PASS/FAIL]`.
4. **Se FAIL:** registrar o déficit via `PROMPT-EXECUTE-1010-TECHNICAL-DEBT-AUDIT` (modo `catalogo`) e escalar ao humano (otimizar vs. renegociar SLO vs. adiar janela).
5. **Registrar** no artefato da fase.

---

## Saída

Gerar `{CICLO_DIR}/PACKAGE-DEVELOPMENT-STRESS-PERFORMANCE-TEST.md`:

```markdown
# PACKAGE-DEVELOPMENT-STRESS-PERFORMANCE-TEST.md — Stress/Performance: Ciclo {N}
[Header: solução, projeto, ciclo, data]
## 1. Aplicável?
[Sim — cenários em 510/050 | Não — registrado]
## 2. Resultado
- Ferramenta: {PERF_TOOL} | Ambiente: {TARGET_ENV}
| Indicador | Medido | Threshold (SLO) | Veredito |
|:---|:---|:---|:---|
## 3. Débitos / Escalamentos
[DT-XXX via 1130 + decisão humana registrada]
## 4. Veredito Final
[✅ PASS | ❌ FAIL (→ decisão humana)]
## Rodapé
[Indicação de geração por IA, data/hora]
```

---

## Skills

| Skill | Modo | Uso na fase |
|:---|:---|:---|
| `verification-before-completion` | automático | Veredito conferido antes de concluir |
| `caveman` | full | Comunicação interativa (nunca em artefatos permanentes) |

> `PROMPT-EXECUTE-STRESS-PERFORMANCE-TEST` é um **prompt especialista executor**, não skill.

---

## Regras de Ouro

1. Roda somente com cenários de performance em 510/050.
2. Nunca executar sem thresholds documentados (regra do executor).
3. Nunca usar dados reais de produção na massa de teste (LGPD).
4. FAIL de SLO = DT-XXX + decisão humana — nunca seguir em silêncio.
