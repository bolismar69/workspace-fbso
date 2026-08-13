# PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-ESTIMATE-VALIDATION

## Contexto

Este prompt implementa o **GENERATE do ESTIMATE-VALIDATION** para o processo de Sourcing & Factory Bidding (Fase 5).

**Propósito:** Valida cada estimativa recebida contra as regras definidas em `.specs/standards/DTA-VALIDATION-STANDARDS.md` §2: QA Balanceado, Arquitetura/SRE, Consistência Prazo×Horas, Outliers, Formato.

**Modo de operação:** Adapta-se ao `SOURCING_BIDDING_MODE` definido no Bootstrap (`agile-discovery`, `agile-refinement`, `waterfall-discovery` ou `waterfall-refinement`).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `agile-discovery`, `agile-refinement`, `waterfall-discovery` ou `waterfall-refinement` |
| `{SOURCING_BIDDING_PATH}` | Pasta sourcing-factory-bidding-{mode} |
| `{ESTIMATES_PATH}` | Pasta de estimativas recebidas |

## Fluxo de Execução

### Passo 0 — Validar Parâmetros e Modo
Confirmar `{PROJECT_PATH}`, `{PROJECT_ID_NAME}`, `{SOURCING_BIDDING_MODE}`.

### Passo 1 — Carregar Artefatos Base
- CSVs das fábricas em `estimates/ESTIMATION-SCHEMA-{FAB}.csv`
- Baseline PIB conforme modo: `DISCOVERY-LEVEL-ROM-ESTIMATE.md` (agile-discovery) ou `BOTTOM-UP-PERT-ESTIMATE.md` (full)
- DTA-VALIDATION-STANDARDS.md §2 (regras de validação)

### Passo 2 — Invocar Skills Especializadas
- `estimate-builder-qmohd` — QA gate validation
- `analyst-estimates` — Cross-source variance + PIB
- `gap-analysis` — Detecção de outliers

### Passo 3 — GENERATE o Artefato

**Especificações do Artefato:**

1. **Seção 1 — Regras:** Tabela com regras DTA+PIB: QA Balanceado, QA Global, Arquitetura, Formato, Consistência Prazo×Horas, Outliers, PIB
2. **Seção 2 — Resultados por Fábrica:** Tabela com `Fábrica | Total Horas | QA% | Arch% | Prazo | PIB Score | Veredito`. Uma linha por fábrica.
3. **Seção 3 — PIB por Épico:** Tabela cross-fábrica com baseline ROM por épico e PIB individual
4. **Seção 4 — Análise PIB Total:** Desvio da baseline, PIB Score, Nota para cada fábrica
5. **Arquivos individuais:** Gerar `estimates/ESTIMATE-VALIDATION-{FABRICA}.md` para cada fábrica com racional de não-compliance
6. **Veredito claro:** APROVADA / REJEITADA / APROVADA COM RESSALVA para cada fábrica

### Passo 4 — Validação Pós-GENERATE
Verificar: 100% fábricas validadas, PIB calculado, baseline do modo correto, vereditos explícitos.

## Skills Utilizados

| 1 | `analyst-estimates` | Análise e validação de estimativas | 2 | `estimate-builder` | Verificação estrutural |
| 3 | `estimate-builder-qmohd` | Validação de qualidade | 4 | `afrexai-construction-estimator` | Metodologia de verificação |
| 5 | `gap-analysis` | Detecção de outliers |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 5 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 5 GENERATE*
