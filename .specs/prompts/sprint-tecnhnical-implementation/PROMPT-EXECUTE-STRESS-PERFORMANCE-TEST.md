# PROMPT-EXECUTE-STRESS-PERFORMANCE-TEST

## Contexto

Este prompt **executa testes de estresse e performance** da solução técnica, operacionalizando os thresholds da `510-TEST-STRATEGY-DEFINITION` (carga/stress/soak, ferramentas k6/JMeter/Gatling) e os SLOs/SLIs da `500-DEVOPS-SRE-DEFINITION` (p50/p95/p99, error budgets, burn rate). O veredito é objetivo: o teste **passa** somente se os indicadores medidos respeitarem os thresholds documentados.

**Princípios fundamentais:**

1. **Thresholds vêm dos documentos:** os números aceitáveis de latência/erro são os SLOs do `500` e os thresholds do `510` — nunca chute do agente.
2. **Ambiente explícito:** o teste roda contra um ambiente definido (DEV/QA/HMG) alinhado às janelas de entrega (096/Bloco F).
3. **Evidência mensurável:** o relatório compara medição × threshold, item a item, com veredito automático por indicador.
4. **Gate obrigatório:** resultado validado pelo `PROMPT-QA-REVISOR-SECURITY` e pela validação humana antes de liberar a janela.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|:---|:---|:---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/go/fiber/microservices/ms-billing-engine-tax-rates` |
| `{PROJECT_NAME}` | Nome/código do projeto de negócio | `PRJ-TEC-2026-0004-PROJETO-SHIELD` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-billing-engine-tax-rates` |
| `{TECH_DEFS_DIR}` | Pasta das definições técnicas TECHLEAD (500/510) | `.../technical-definitions/` |
| `{PROJECT_DOCS_DIR}` | Pasta dos documentos WATERFALL (045/050) | `.../PRJ-TEC-2026-0004-PROJETO-SHIELD/` |
| `{PERF_TOOL}` | Ferramenta de performance (se omitido, usar a da 510) | `k6`, `jmeter`, `gatling` |
| `{TARGET_ENV}` | Ambiente alvo do teste | `QA` |
| `{SCENARIOS}` | Casos de teste de performance a executar (IDs do 050, se omitido executar os cenários de performance mapeados na 510) | `TC-PERF-001,TC-PERF-002` |

---

## Documentos de Referência (obrigatórios — fonte da verdade)

```
Ler obrigatoriamente antes de executar:

    ├── {TECH_DEFS_DIR}/510-TEST-STRATEGY-DEFINITION.md  ← Testes de performance: carga/stress/soak, ferramentas, thresholds
    ├── {TECH_DEFS_DIR}/500-DEVOPS-SRE-DEFINITION.md     ← SLOs/SLIs (p50/p95/p99, availability, error budgets, burn rate)
    ├── {PROJECT_DOCS_DIR}/045-TEST-PLAN-*.md            ← Requisitos de ambiente de teste e estratégia de dados de teste
    └── {PROJECT_DOCS_DIR}/050-TEST-CASES-*.md           ← Cenários de performance/estresse aplicáveis
```

> ⚠️ Se `510`/`500` não existirem ou não definirem thresholds de performance → **PARE** e solicite os valores ao humano (teste sem threshold não é teste). NUNCA inferir sem ancoragem documental.

---

## Missão

Executar os cenários de estresse/performance de `{SOLUTION_NAME}` no ambiente `{TARGET_ENV}` com `{PERF_TOOL}`, e emitir veredito objetivo por indicador comparando as medições com os SLOs/thresholds documentados.

---

## Fluxo de Execução

### Fase 0 — Preparação

0. **Validar ambiente alvo** (janela de entrega correspondente no 096/Bloco F; ambiente `{TARGET_ENV}` disponível conforme 045).
1. **Carregar documentos-base** e extrair:
   - De `510`: tipos de teste aplicáveis (carga, stress, soak, benchmark), ferramenta padrão e thresholds por indicador.
   - De `500`: SLOs por serviço — latência p50/p95/p99, disponibilidade (9s), error budget e thresholds de burn rate.
   - De `045`/`050`: cenários e massa de dados (sintética/anônima, LGPD).

### Fase 1 — Roteiro de Teste

2. **Gerar o roteiro** (arquivo do `{PERF_TOOL}` + plano): cenários, carga base, rampa, pico (stress), duração do soak, ambiente, coleta de métricas. Cada cenário referencia o ID do 050 e o SLO que valida.

### Fase 2 — Execução

3. **Executar o teste** contra `{TARGET_ENV}` e capturar as métricas: latência (p50/p95/p99), taxa de erro, throughput, consumo de recursos (CPU/memória) e qualquer indicador de burn rate.

### Fase 3 — Relatório com Veredito

4. **Gerar o relatório** com comparação item a item:

```markdown
# Relatório de Estresse/Performance — {SOLUTION_NAME}
- Ferramenta: {PERF_TOOL} | Ambiente: {TARGET_ENV} | Data: {DATA}

| Indicador | Medido | Threshold (SLO) | Veredito |
|:---|:---|:---|:---|
| p95 latência | 380ms | ≤ 500ms (500-DEVOPS-SRE §SLO) | ✅ PASSOU |
| Taxa de erro no pico | 1,8% | ≤ 1,0% (500 §error budget) | ❌ FALHOU |
| p99 latência (stress) | 2,1s | ≤ 1,5s (510 §performance) | ❌ FALHOU |
```

**Veredito final:** `[STATUS: PASS]` se TODOS os indicadores respeitarem os thresholds; `[STATUS: FAIL]` se qualquer indicador violar o SLO.

### Fase 4 — Tratamento e Gate

5. **Se `[STATUS: FAIL]`:** registrar o déficit de performance no documento de débitos (`DT-XXX`, com o indicador e o threshold violado) e escalar ao humano — correção de performance é decisão do time (otimizar agora vs. renegociar SLO vs. adiar janela).
6. **Gate:** submeter o relatório ao `PROMPT-QA-REVISOR-SECURITY` e à validação humana. Só com `[STATUS: PASS]` + aprovação humana a janela de entrega pode prosseguir.
7. **Registrar evidências** no relatório de implementação da sprint (seção de testes) e alimentar `095-RELATORIO-QUALIDADE` com o resultado.

---

## Regras de Ouro

1. NUNCA executar teste de performance sem thresholds documentados (500/510).
2. NUNCA usar dados reais de produção na massa de teste (LGPD — usar sintético/anônimo conforme 045).
3. Veredito é objetivo: medição × threshold. Não existe "passou com ressalva" para violação de SLO.
4. Todo `[STATUS: FAIL]` gera `DT-XXX` rastreável e decisão humana registrada.
5. Relatório final anexado às evidências da janela (595-RETURN-PACKAGE / 095).
