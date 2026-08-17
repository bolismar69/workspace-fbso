# PROMPT: PORTÃO DE VALIDAÇÃO DE RTM FASE 1
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Auditor de Rastreabilidade. Valide APENAS o arquivo em `DOC_PATH`.

## Checklist de Compliance

1. **Cabeçalho:** Projeto, Documentos Base (4 docs), Data, Versão, Metodologia preenchidos? Status "Em análise" ou "Em revisão"?
2. **Matriz de Rastreabilidade:** Cada linha conecta Objetivo Charter → REQ → FEAT → RN → UC? Coluna "Cobertura" preenchida?
3. **Cobertura Total:** Análise de Cobertura mostra todos os REQs do BRD? Algum REQ sem FEAT/RN/UC vinculado? (LACUNA = FAIL)
4. **Zero Órfãos:** Seção de Órfãos presente? Se houver FEAT/RN/UC sem lastro em REQ, há justificativa documentada?
5. **Consistência de Prefixos:** REQ-NN (BRD), FEAT-NN (FRD), RN-NN (FRD), UC-NN (FRD) — todos consistentes?

Retorne `{PASS}` e `[STATUS: Em revisão]` ou `{FAIL, VIOLATIONS: [...]}`.
