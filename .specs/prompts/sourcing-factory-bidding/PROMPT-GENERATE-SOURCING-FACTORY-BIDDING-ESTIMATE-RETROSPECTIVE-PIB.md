# PROMPT: GENERATE — SOURCING-FACTORY-BIDDING — ESTIMATE-RETROSPECTIVE-PIB (F5b)
## Versão: 1.0 — Análise Retrospectiva com PIB (Condicional: 0 Aprovadas)

Atue como um Engineering Manager e Tech Lead especializado em análise forense de estimativas e detecção de anomalias em processos de sourcing.

## OBJETIVO

Produzir uma análise retrospectiva aprofundada quando **0 fábricas são aprovadas** na Fase 5. Esta fase complementa a validação DTA com uma análise qualitativa que identifica problemas sistêmicos, padrões suspeitos e observações que devem ser comunicadas às fábricas no realinhamento.

## CONDICIONALIDADE

⚠️ **Esta fase é condicional.** Só executa se o resultado da Fase 5 for **0 fábricas aprovadas**. Se pelo menos 1 fábrica foi aprovada, esta fase é **pulada** e o roadmap avança diretamente para F6.

## INPUTS

1. **ESTIMATE-VALIDATION.md** (F5) — resultado da validação DTA
2. **CSVs das fábricas** em `estimates/ESTIMATION-SCHEMA-{FAB}.csv` — dados brutos
3. **Baseline PIB** conforme o modo:
   - `agile-discovery` → `upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md`
   - `agile-refinement` → `downstream-architecture-refinement/BOTTOM-UP-PERT-ESTIMATE.md`
   - `waterfall-discovery` → `waterfall-estimation/WATERFALL-ESTIMATION-UPSTREAM-ROM.md`
   - `waterfall-refinement` → `waterfall-estimation/WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md`
4. **DTA-VALIDATION-STANDARDS.md** — regra PIB §2.6

## DIMENSÕES DE ANÁLISE

### Dimensão 1: PIB por Épico
Calcular PIB Score individualmente para cada épico. Identificar qual fábrica mais se aproxima da baseline em cada épico. Se a melhor fábrica ainda está muito distante (>100% de desvio), documentar como alerta sistêmico.

### Dimensão 2: Detecção de Flat Estimates
Calcular o coeficiente de variação (CV) entre épicos para cada fábrica:
```
CV = desvio_padrão(horas_epico) / média(horas_epico)
```
Se CV < 10% para uma fábrica → ⚠️ Alerta: estimativa plana. A fábrica aplicou o mesmo valor para todos os épicos independente da complexidade.

### Dimensão 3: QA/Arch como Overhead Fixo
Verificar se `horas_qa` e `horas_arch` são valores absolutos idênticos em todos os épicos (ex: 320h para cada épico). Se sim → ⚠️ Alerta: tratados como overhead fixo em vez de percentual do esforço.

### Dimensão 4: Qualidade dos Comentários
Analisar a coluna `comentarios` de cada fábrica:
- Texto genérico (< 50 caracteres) → ⚠️
- Texto idêntico entre fábricas diferentes → 🔴
- Sem menção a metodologia ou premissas → ⚠️

### Dimensão 5: Independência do Processo
Comparar valores entre fábricas. Se duas ou mais fábricas têm:
- Mesmo `total_horas`
- Mesmos valores em `horas_dev`, `horas_qa`, `horas_arch` por épico
- Mesmos textos em `comentarios` e `premissas`

→ 🔴 Possível violação de independência. Documentar com evidências.

## ESTRUTURA DO DOCUMENTO

```markdown
# ESTIMATE-RETROSPECTIVE-PIB — Análise Retrospectiva com PIB

## 1. PIB por Épico
[Tabela: Épico | Baseline | Melhor Fábrica | Horas | Desvio | PIB Score]

## 2. Flat Estimates
[Lista de fábricas com CV < 10% + evidência dos valores por épico]

## 3. QA/Arch como Overhead Fixo
[Tabela: Fábrica | Total | QA Real | QA Mínimo | Arch Real | Arch Mínimo]

## 4. Comentários Genéricos
[Evidências de textos genéricos ou idênticos entre fábricas]

## 5. Independência Comprometida
[Fábricas com valores idênticos + evidências lado a lado]

## 6. Recomendações para Realinhamento
[Lista de ações adicionais além das já documentadas na F5]
```

## REGRA CRÍTICA

⚠️ Esta análise NÃO altera os vereditos da F5. Ela adiciona **observações qualitativas** que enriquecem o feedback às fábricas. O resultado da F5 (aprovada/rejeitada) permanece inalterado.

## SKILLS RECOMENDADOS

- `analyst-estimates` — Cross-source variance analysis
- `gap-analysis` — Detecção de anomalias e padrões
- `project-estimation` — Contexto de estimativas

🤖 *Prompt gerador — Fase 5b do Sourcing & Factory Bidding. Condicional: 0 aprovadas.*
