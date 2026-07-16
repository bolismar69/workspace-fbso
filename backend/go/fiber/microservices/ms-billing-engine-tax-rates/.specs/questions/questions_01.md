# Lacunas e Perguntas — ms-billing-engine-tax-rates

Atualizado em 2026-06-30 (PR #6 merge — Fases 0-1-2 Reforma Tributária concluídas).

## Lacunas Resolvidas

### 1. Fluxo ICMS Próprio Incompleto ✅

**Resolução (2026-06-20):** O fluxo foi completamente reestruturado em `icms.go:23-70`. `getEffectiveTaxConfig()` agora é a primeira etapa do `Calculate()` e alimenta os 3 novos métodos auxiliares: `calcularICMSSimples()`, `calcularICMSOperacaoInterna()`, `calcularICMSOperacaoInterestadual()`. Ver `.specs/skill-output/2026-06-20-192700_icms-proprio-st.md`.

---

### 2. Alíquotas PIS/COFINS Hardcoded ✅

**Resolução (2026-06-20):** As alíquotas agora são carregadas do banco via `GetFederalTaxRule(ctx, regime, cstPis, cstCofins)` com fallback para constantes do pacote (`defaultAliquotaPIS=1.65`, `defaultAliquotaCOFINS=7.6`). Ver `.specs/skill-output/2026-06-20-192730_pis-cofins-db-rates.md`.

---

### 3. Erros Silenciosos nas Goroutines da Fase 2 ✅

**Resolução (2026-06-21):** Implementado channel de erro com buffer (`errChan := make(chan error, len(e.calculators))`) em `engine.go:86`. Erros são coletados após `wg.Wait()` e logados como `slog.Warn`. Resultados parciais das calculadoras que funcionaram são retornados (degradação graciosa). Ver `.specs/skill-output/2026-06-21-005150_health-checks_goroutine-errors.md`.

---

### 4. Cobertura de CSTs PIS/COFINS ✅

**Resolução (2026-06-20):** Todos os CSTs pendentes implementados com testes:
- PIS: 01_02, 03, 04, 05, 06, 49, 50-99, 99 — 8 estratégias
- COFINS: 01_02, 03, 04, 05, 06, 49, 50-99 — 7 estratégias
- 26 testes unitários + 13 testes de integração. Ver `.specs/skill-output/2026-06-20-204235-pis-cofins-cst-completion.md`.

---

### 5. Módulo Reforma Tributária (CBS/IBS/IS) ✅

**Resolução (2026-06-21 15:17):** Implementado `internal/reforma/reforma.go` com `ReformaCalculator` implementando `TaxCalculator`. Adicionado `GetIvaDualRule` ao `TaxRepository` (interface + PostgreSQL + cache Redis). CBS, IBS e IS calculados com suporte a redução de alíquotas (0%, 60%, 100%), imposto seletivo e lookup por município. 7 testes unitários. Wiring no motor bifásico como Fase 2 (paralela). Ver `.specs/skill-output/2026-06-21-151743_reforma-tributaria-cbs-ibs-is.md`.

---

### 6. ICMS Desonerado (F-004) ✅

**Resolução (2026-06-21):** Implementado `internal/legacy/icms_desoneracao.go` com 2 modos: Redução de Base (BR-TAX-CALC-021) e Limitação de Alíquota (BR-TAX-CALC-022). CST validation, motDesICMS (1-12, 90), FCP integrado. 14 testes. Ver `.specs/skill-output/2026-06-21-235900_icms-desonerado-phase-resolver.md`.

---

### 7. Phase Resolution System (F-005) ✅

**Resolução (2026-06-21):** `PhaseResolver` mapeia DataOperacao→Phase (SHADOW_RUN/2026, CBS_PLENA/2027-28, TRANSICAO_SUBNACIONAL/2029-32, IVA_DUAL/2033+). `TaxSelector` com matriz DT-001. 5 testes. Ver `.specs/skill-output/2026-06-21-235900_icms-desonerado-phase-resolver.md`.

---

### 8. IS Pré-Filtro + IBS Circuit Breaker (F-006, F-007) ✅

**Resolução (2026-06-22):** IS extraído como Fase 0 pré-filtro independente (`is_filter.go`, 8 testes). IBS Circuit Breaker com 4 implementações de `IBSRateFetcher` (HTTP, Cached, CircuitBreaker, FallbackDB). Ver `.specs/skill-output/2026-06-22-004300_is-prefilter-ibs-circuitbreaker.md`.

---

### 9. Pipeline SOP-013 7-fases (C-001) + Schema Expandido (C-002) ✅

**Resolução (2026-06-22):** Pipeline reordenado de 3 para 7 fases. `BillingEnginePhased` com `CalculationPhase` genérico. Schema expandido com `ncm_seletivo`, `cbs_rates`, `iss_rates`. 22 testes de pipeline. Ver `.specs/skill-output/2026-06-22-011000_c001-pipeline-sop013.md`.

---

### 10. Fases 0-1-2 Reforma Tributária (PR #6) ✅

**Resolução (2026-06-24/30):** Fase 0 (Rate Limiting, API Versioning, Docker/K8s), Fase 1 Onda 1 (Admin Fiscal, TaxToken, Simulação), Fase 2 Onda 2 (Fornecedores, Créditos). 211+ testes, 25 arquivos de teste. Ver `pull-requests/PR_6-...md`.

---

## Lacunas Abertas

### 8. Documentação de CI/CD

**Contexto:** Não há workflows de CI/CD (GitHub Actions, etc.) configurados para este microserviço Go. O monorepo tem CI para projetos Java, mas não para Go.

**Pergunta:** Há planos para adicionar CI/CD (build, lint, test, deploy) para este serviço?
**Resposta:** Sim, será desenvolvido assim que finalizarmos o desenvolvimento dos microserviços em '~/work/workspace-fbso/backend/go'. Esta no roadmap do monorepo.

---

### 9. CST 99 do COFINS sem estratégia dedicada

**Contexto (2026-06-21):** O `GetCOFINSStrategy()` em `pis_cofins.go:183-199` não possui um `case` para CST 99. Quando CST COFINS é "99", cai no `default` (`COFINS50To99`), que retorna valor zero. Enquanto isso, PIS tem uma estratégia dedicada `PIS99` (`pis_cofins.go:176-177`).

**Impacto:** É o comportamento correto? CST 99 para COFINS deveria ser tratado de forma diferente de 50-98?

**Resposta:** A documentação foi atualizada (`domain.md`) para refletir que COFINS 99 usa fallback 50-99 (valor zero). Pendente confirmação com especialista fiscal se é necessário tratamento diferenciado.

---

### 10. Tabela `reforma_tributaria_rules` não utilizada no código

**Contexto (2026-06-21):** O `data/init.sql` contém a criação da tabela `reforma_tributaria_rules` (linhas 325-337) como "primeiro rascunho", mas o código de produção em `internal/reforma/reforma.go` utiliza apenas `iva_dual_rules` via `GetIvaDualRule()`. A tabela `reforma_tributaria_rules` parece ser um artefato de design inicial que foi substituído pelo modelo `iva_dual_rules`.

**Pergunta:** A tabela `reforma_tributaria_rules` deve ser mantida como referência histórica ou pode ser removida em uma limpeza futura de schema?