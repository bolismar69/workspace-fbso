# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — PRD (F1)
## Versão: 1.1 — PRD Detail-Level — Documento de Negócio para TI

Atue como um Product Manager, Product Owner (PO) e Business Analyst da área de Negócios. Você representa a visão do negócio e está preparando um documento para apresentar o projeto ao time de TI.

## OBJETIVO

Produzir um **PRD Detail-Level criado pelo Negócio (PO/PM/Funcional)** que apresenta um resumo completo do projeto para o time de TI. Este documento serve como a ponte oficial negócio→TI e referencia todos os documentos de projeto existentes.

**Este roadmap é independente.** Este documento NÃO deve buscar referências na pasta `upstream-architecture-discovery/`. Toda informação necessária está nos documentos de negócio do projeto.

## INPUTS (Exclusivamente Documentos de Projeto)

1. **Project Charter:** `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md`
2. **BRD:** `02-BRD-{PROJECT_ID_NAME}.md`
3. **Épicos:** `03-EPICS-{PROJECT_ID_NAME}.md` e `epics/*.md`
4. **Features:** `04-FEATURES-{PROJECT_ID_NAME}.md` e `features/*.md`
5. **User Stories:** `05-USER-STORIES-{PROJECT_ID_NAME}.md` e `user-stories/*.md`
6. **Documentos de apoio (NÃO obrigatórios — verificar existência):**
   - `DEFINITION_OF_DONE.md` — se existir, referenciar na seção de restrições
   - `GLOSSARY.md` — se existir, incorporar termos ao glossário do PRD
   - `MATRIZ-KPI.md` — se existir, referenciar nas restrições de negócio (métricas de sucesso)
   - `STAKEHOLDER-MAP.md` — se existir, referenciar nas personas
   - `PRODUCT-BACKLOG-LIST.md` — se existir, referenciar no escopo
   - **Regra:** Verificar a existência de cada um. Se existir, referenciar. Se não existir, prosseguir sem ele — não são obrigatórios.

## REGRA CRÍTICA

⚠️ **Não buscar NENHUMA referência em `upstream-architecture-discovery/`.** Este roadmap é independente. Se o upstream foi executado ou não, é irrelevante para este documento. O PRD Detail-Level é gerado exclusivamente a partir dos documentos de negócio do projeto.

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-PRD — Product Requirements Document (Detail-Level)

- **Autor:** Negócio (PO/PM/Funcional)
- **Audiência:** Time de TI
- **Data:** {data_atual}
- **Documentos Referenciados:** Charter, BRD, Épicos, Features, User Stories

## 1. Visão do Produto
[Resumo executivo do produto para TI — o que é, por que existe, para quem]

## 2. Personas e Jornadas
[Por persona: necessidades, dores, cenários de uso, US relacionadas]

## 3. Escopo por Entrega (D1-D7)
[Por entrega: objetivo de negócio, features incluídas, US, valor esperado, dependências]

## 4. Matriz US × Jornada
[Tabela: cada US mapeada para sua jornada de negócio]

## 5. Restrições de Negócio
[LGPD, compliance, SLA, restrições de mercado, premissas de negócio]

## 6. Glossário de Negócio
[Termos de domínio que TI precisa conhecer]

## 7. Referências aos Documentos de Projeto
[Lista completa de documentos de negócio referenciados, com paths relativos]
```

### Skills Recomendados
- `business-analyst`, `agile-ba-practices`
- `product-manager`, `prd-development`, `prd`
- `discovery-process`
- `documentation-writer`, `documentation`

🤖 *Prompt gerador — Fase 1 do Downstream Architecture Refinement · Documento de Negócio Independente*
