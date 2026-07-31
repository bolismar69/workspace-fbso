# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ROM-ESTIMATE
## Contexto
> 📐 **Discovery-Level:** Consolidação da estimativa ROM +-50% para o comitê de governança.

Este prompt gera `ROM-ESTIMATE.md` — documento que consolida a estimativa Rough Order of Magnitude com faixa de +-50%. Inclui matriz de esforço por solução, premissas, riscos e faixa de valores. Este é o output final do Discovery para decisão GO/NO-GO.

**Papel no Bloco D (Estimativa & ROM):** Fase 11 de 1. Consome todos os artefatos anteriores.

**Inputs upstream:** Todos os artefatos Discovery-Level (F1-F10).

**Seções do artefato:**
1. **Matriz de Esforço por Solução** — solução, complexidade, esforço estimado (homem-mês), faixa (min-máx)
2. **Premissas** — assumptions que embasam a estimativa
3. **Riscos e Mitigações** — riscos identificados com impacto na estimativa
4. **Faixa de Valores** — ROM consolidado (min, provável, máx) com +-50%
5. **Recomendação Técnica** — parecer do time de arquitetura para o comitê

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `product-discovery` | Mapear premissas e validar cenários de estimativa | Discovery |
| 2 | `senior-architect` | Estimativa de esforço técnico | Arquitetura |
| 3 | `cloud-architect` | Custos de infra | Cloud |
| 4 | `senior-devops` | Esforço DevOps/SRE | DevOps |
| 5 | `gap-analysis` | Análise de riscos | Análise |
| 6 | `documentation-writer` | Documento executivo | Documentação |

## Registro de Alterações
| 1.0 | 30/07/2026 | Criação inicial — F11 Bloco D Discovery-Level | Time de Arquitetura |
🤖 *Upstream Architecture Discovery — Fase 11*
