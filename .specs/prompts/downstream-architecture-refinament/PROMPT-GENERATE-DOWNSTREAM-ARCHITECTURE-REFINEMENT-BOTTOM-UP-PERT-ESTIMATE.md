# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — BOTTOM-UP-PERT-ESTIMATE (F8)
## Versão: 1.0 — Estimativa Bottom-Up PERT Three-Point por User Story

Atue como um Especialista em Estimativas de Projetos de Software e Tech Lead, especializado em estimativas bottom-up com PERT three-point e validação DTA.

## OBJETIVO

Produzir uma estimativa detalhada bottom-up usando PERT three-point (O + 4M + P) / 6 para **cada User Story individualmente**, consolidando por feature → épico → projeto, com composição completa do esforço (Dev + QA + Arquitetura + DevOps + Gestão + Contingência).

## REGRA CRÍTICA DE INDEPENDÊNCIA

⚠️ **Esta estimativa é 100% independente.** Você NÃO deve consultar, referenciar ou usar como baseline NENHUMA estimativa anterior (ROM upstream, factory bids, etc.). Cada US é estimada do zero usando apenas:
- O conteúdo da User Story (descrição, critérios de aceitação, cenários Gherkin)
- A complexidade técnica da feature à qual pertence
- O conhecimento da stack tecnológica do projeto (identificada nos artefatos Detail-Level F2-F7 e na matriz de stacks)

Se o arquivo `upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md` existir, **ignore-o completamente** durante o cálculo. Ele só será consultado na Fase 12 (Cross-Check), após esta estimativa estar concluída.

## INPUTS

1. **User Stories:** Ler TODOS os arquivos em `user-stories/`
2. **Features:** `04-FEATURES-{PROJECT_ID_NAME}.md` e `features/*.md`
3. **Épicos:** `03-EPICS-{PROJECT_ID_NAME}.md`
4. **Artefatos Detail-Level (Bloco A):** F1-F7 (PRD, Arquitetura, Segurança, Dados, DevOps, Testes, Infra) — use para entender complexidades técnicas
5. **TEAM-SKILLS-MAP:** {obter a partir do contexto do projeto, e questionar o usuario sobre skills necessários para serem foco da solução}
6. **TEAM-CAPACITY:** {obter a partir do contexto do projeto, e questionar o usuario sobre as capacidades esperadas para o time necessário para o foco da solução}
7. **Stack tecnológica:** {obter a partir do contexto do projeto, e questionar o usuario sobre todas as tecnologias a serem foco da solução}

## METODOLOGIA

### 1. Classificação de Complexidade por US

| Complexidade | Critério | O (h) | ML (h) | P (h) | Exemplos |
|:---|:---|---:|---:|---:|:---|
| **1 — Simples** | CRUD básico, listas, filtros, telas simples, sem integração externa | 12-24 | 24-40 | 40-72 | Cadastrar/editar registros, busca textual, ativar/desativar |
| **2 — Média** | Workflows, integrações, regras de negócio, autenticação, notificações | 24-56 | 40-80 | 72-144 | Dashboards com agregações, onboarding guiado, filtros avançados, emails transacionais |
| **3 — Complexa** | Orquestração multi-serviço, RBAC granular, state machines, segurança avançada, performance crítica | 40-56 | 72-96 | 128-160 | Matriz de permissões, integração Keycloak avançada, multi-tenancy RLS, upgrade/downgrade de planos |

### 2. Cálculo PERT por US

Para cada US:
```math
PERT = (O + 4×ML + P) / 6
σ = (P - O) / 6
IC 95% = PERT ± 1.96×σ
```

### 3. Rollup e Composição

Após estimar cada US individualmente:
1. Somar PERT por feature
2. Somar por épico
3. Aplicar percentuais sobre o total de desenvolvimento:
   - QA: 30% (meta DTA: ≥25%)
   - Arquitetura: 8% (meta DTA: ≥5%)
   - DevOps/SRE: 7%
   - Gestão/Governança: 10%
   - Infraestrutura residual: 480h fixo (M1 já concluído)
   - Contingência: 15-25% sobre o subtotal

## ESTRUTURA DO DOCUMENTO

```markdown
# BOTTOM-UP-PERT-ESTIMATE — Estimativa Bottom-Up PERT Three-Point

- **Projeto:** {PROJECT_ID_NAME}
- **Data:** {data_atual}
- **Metodologia:** PERT Three-Point (O + 4M + P) / 6
- **Independência:** Estimativa 100% independente — calculada do zero, US por US
- **Confiança alvo:** ±15-25%

## 1. Sumário Executivo

| Cenário | Horas | Homem-Mês (160h) | IC 95% |
|:---|---:|---:|:---|
| Desenvolvimento PERT | {dev_h}h | {dev_hm} h-m | {ic_low}h – {ic_high}h |
| Subtotal (Dev+QA+Arch+DevOps+Gestão+Infra) | {subtotal}h | {subtotal_hm} h-m | — |
| TOTAL com Contingência 15-25% | {low}h – {high}h | {low_hm} – {high_hm} h-m | — |

## 2. Metodologia

[Explicar PERT, classificação de complexidade, composição do esforço, regra de independência]

## 3. Estimativa por Épico

### 3.1 EP-0001 — [Nome]
[Para cada feature, tabela: US ID | Descrição | Compl. | O | ML | P | PERT | σ]
[Resumo do épico: tabela Dev | QA | Arch | DevOps | Gestão | Total]

### 3.2 EP-0002 — [Nome]
[Idem]

[...]

## 4. Alocação de Recursos
[Time disponível, capacidade mensal, projeção de duração]

## 5. Análise de Riscos
[6 riscos com probabilidade, impacto e ação]

## 6. Validação DTA
[QA≥25%, Arch≥5%, consistência prazo×horas, independência comprovada]

## 7. Documentos Relacionados

🤖 *Documento gerado pelo Tech Lead / Especialista em Estimativas — Fase 8 do Downstream Architecture Refinement · Skills utilizados: [lista de skills efetivamente acionados] · Padrões Corporativos FBSO.ORG*
```

## VALIDAÇÃO DTA (pré-gate)

Antes de enviar ao gate, verificar:
- [ ] Todas as US estimadas individualmente (O, ML, P, PERT, σ)
- [ ] QA ≥ 25% do total
- [ ] Arch ≥ 5% do total
- [ ] IC 95% calculado para cada nível (US, feature, épico, projeto)
- [ ] NENHUMA referência ao ROM upstream nos cálculos
- [ ] Consistência prazo×horas validada
- [ ] Outliers identificados e justificados

## SKILLS RECOMENDADOS

- `engineering-skills`, `engineering-advanced-skills`
- `project-estimation` — core da estimativa
- Referências: `bottom-up-estimation.md`, `three-point-estimation-pert.md`
- `gap-analysis` — identificação de outliers e riscos
- `context-manager` — gestão de contexto para 62 US

## NOTAS

- Se o projeto tiver menos ou mais US que 62, ajustar automaticamente
- A classificação de complexidade deve ser justificada para cada US
- US "Should Have" são estimadas mas marcadas como opcionais no sumário
- O documento final é um snapshot — não deve ser alterado após aprovação na Barreira B

🤖 *Prompt gerador — Fase 8 do Downstream Architecture Refinement · PERT Three-Point Bottom-Up · Skills: `engineering-skills`, `engineering-advanced-skills`, `project-estimation`, `gap-analysis`, `context-manager` · Padrões Corporativos FBSO.ORG*
