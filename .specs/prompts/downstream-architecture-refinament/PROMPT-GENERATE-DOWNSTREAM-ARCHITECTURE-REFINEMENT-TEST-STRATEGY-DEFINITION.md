# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — TEST-STRATEGY-DEFINITION (F6)
## Versão: 1.1 — Estratégia de Testes Detail-Level (Matriz de Cobertura + Casos de Teste) — Independente de Tecnologia

Atue como um QA Lead e Test Specialist especializado em estratégias de teste para aplicações multi-tenant e sistemas distribuídos.

## OBJETIVO

Produzir a estratégia de testes em nível de implementação: matriz de cobertura por US, casos de teste de aceitação, estratégia de automação, quality gates.

**Este documento é independente de tecnologias específicas de teste.** Durante a análise da stack do projeto, identifique as ferramentas e frameworks de teste utilizados e busque skills relacionados. Caso não encontre skills específicos, utilize skills generalistas de qualidade e teste.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **Segurança Detail-Level** (F3)
3. **Todas as User Stories** (cenários de aceitação)

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-TEST-STRATEGY-DEFINITION — Estratégia de Testes Detail-Level

## 1. Pirâmide de Testes Refinada
- Testes Unitários: framework e ferramentas (meta ≥80%)
- Testes de Integração: framework e ferramentas (meta ≥60%)
- Testes E2E: ferramentas (fluxos críticos)
- Testes de Performance: ferramentas (carga, stress)
- Testes de Segurança: SAST, DAST, Secret Scanning

## 2. Matriz de Cobertura por US
| US-ID | Unitário | Integração | E2E | Performance | Segurança | Responsável |
[Uma linha por US]

## 3. Casos de Teste de Aceitação
[Baseados nos cenários de aceitação das US — 3-5 casos por feature]

## 4. Estratégia de Automação
- CI: testes unitários + integração a cada commit/PR
- Nightly: E2E + performance smoke tests
- Release: E2E completo + performance completo + security scan

## 5. Quality Gates
| Gate | Critério | Bloqueia? |
| PR | Unit ≥80%, Integ ≥60%, SAST limpo | Sim |
| Staging | E2E críticos passam, Perf smoke OK | Sim |
| Release | E2E 100%, Security 0 críticas | Sim |

## 6. Testes de Isolamento Multi-Tenant
[Cenários de verificação de isolamento entre tenants]

## 7. Riscos de Qualidade
```

### Skills Recomendados

**Skills generalistas de teste e qualidade (sempre aplicáveis):**
- `senior-qa`, `testing-patterns`, `test-strategy-design`
- `test-master`, `testing-qa`, `test-driven-development`
- `e2e-testing`, `e2e-testing-patterns`
- `test-automator`, `test-case-creation`
- `qa`, `qa-test-planner`, `acceptance-criteria`
- `performance-testing-review-multi-agent-review`

**Skills tecnológicos de teste (condicionais — buscar ao identificar a stack):**
- Ao identificar um framework ou ferramenta de teste específica durante a análise da stack, busque skills relacionados a essa tecnologia para aprimorar as especificações de teste
- Caso não encontre skills específicos para a ferramenta identificada, utilize os skills generalistas listados acima como referência

🤖 *Prompt gerador — Fase 6 do Downstream Architecture Refinement · Independente de Tecnologia*
