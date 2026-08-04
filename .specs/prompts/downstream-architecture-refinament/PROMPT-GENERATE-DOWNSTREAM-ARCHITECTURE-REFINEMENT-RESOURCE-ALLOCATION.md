# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — RESOURCE-ALLOCATION (F9)
## Versão: 1.0 — Plano de Alocação de Recursos baseado na Estimativa PERT

Atue como um Tech Lead e Resource Manager especializado em alocação de times técnicos.

## OBJETIVO

Produzir um plano de alocação de recursos baseado exclusivamente na estimativa PERT da Fase 8, identificando capacidade, duração projetada e gargalos.

## INPUTS

1. **BOTTOM-UP-PERT-ESTIMATE.md** (F8) — estimativa concluída e aprovada
2. **TEAM-SKILLS-MAP:** {obter a partir do contexto do projeto, e questionar o usuario sobre skills necessários para serem foco da solução}
3. **TEAM-CAPACITY:** {obter a partir do contexto do projeto, e questionar o usuario sobre as capacidades esperadas para o time necessário para o foco da solução}
4. **Project Charter** — time previsto e restrições

## REGRA CRÍTICA

⚠️ Este plano é derivado EXCLUSIVAMENTE da estimativa PERT da Fase 8. Não usar ROM upstream ou qualquer outra fonte.

## ESTRUTURA DO DOCUMENTO

```markdown
# RESOURCE-ALLOCATION-PLAN — Plano de Alocação de Recursos

## 1. Time Necessário

> 📌 Como este roadmap é para estimativa, a tabela lista os **Papéis Necessários** (sem nomeação individual — trata-se de dimensionamento de perfis para a estimativa). As colunas de Disponibilidade, Atuação Principal, Épicos e Observações são mantidas pois são relevantes para o dimensionamento da estimativa.

[Tabela: Papel | Disponibilidade % | Atuação Principal | Épicos | Observações]

## 2. Capacidade Mensal
- Capacidade bruta: N pessoas × 160h = X h/mês
- Capacidade efetiva (com cargas parciais): Y h/mês

## 3. Projeção de Duração
| Cenário | Horas | Capacidade Efetiva | Duração |
|:---|---:|---:|---:|
| Conservador (15%) | | | |
| PERT | | | |
| Pessimista (25%) | | | |

## 4. Alocação por Épico
[Tabela: Épico | Horas PERT | % do Total | Perfil Necessário | Duração Estimada]

## 5. Gargalos Identificados
[Ex: Frontend sem dev dedicado, RBAC requer especialista Keycloak, etc.]

## 6. Recomendações
[Reforços, treinamentos, ajustes de cronograma]

🤖 *Documento gerado pelo Tech Lead / Resource Manager — Fase 9 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

### Skills Recomendados
- `engineering-skills`, `engineering-advanced-skills`
- `project-estimation`, `context-manager`, `senior-architect`

🤖 *Prompt gerador — Fase 9 do Downstream Architecture Refinement · Skills: `engineering-skills`, `engineering-advanced-skills`, `project-estimation`, `context-manager`, `senior-architect` · Padrões Corporativos FBSO.ORG*
