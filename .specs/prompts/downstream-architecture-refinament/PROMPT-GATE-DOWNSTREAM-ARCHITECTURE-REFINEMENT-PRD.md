# PROMPT-GATE-DOWNSTREAM-ARCHITECTURE-REFINEMENT-PRD (F1)

## Contexto

Gate de Validação do PRD Detail-Level (`DETAIL-LEVEL-PRD.md`). Verifica se o documento de negócio atende aos critérios para apresentação ao time de TI.

**Princípio fundamental:** Documento criado pelo Negócio (PO/PM/Funcional), independente de upstream discovery, com referências completas aos documentos de projeto.

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |

**Arquivo auditado:** `DETAIL-LEVEL-PRD.md`

## Dimensões de Validação

### Dimensão 1: Independência
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Sem referências ao upstream | Nenhuma menção a `upstream-architecture-discovery/` |
| 1.2 | Fontes corretas | Referencia exclusivamente docs de negócio (Charter, BRD, Épicos, Features, US) |

### Dimensão 2: Completude de Negócio
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Visão do produto | Clara e em linguagem de negócio |
| 2.2 | Personas e jornadas | Todas as personas cobertas com cenários de uso |
| 2.3 | Escopo por entrega | D1-D7 mapeados com features e valor de negócio |
| 2.4 | Matriz US×Jornada | Todas as US mapeadas para pelo menos uma jornada |
| 2.5 | Restrições de negócio | Regulatório, compliance, SLA documentados |

### Dimensão 3: Referências
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Docs de projeto referenciados | Charter, BRD, Épicos, Features, US com paths |
| 3.2 | Glossário | Termos de domínio definidos |
| 3.3 | Docs de apoio (NÃO obrigatórios) | DEFINITION_OF_DONE.md, GLOSSARY.md, MATRIZ-KPI.md, STAKEHOLDER-MAP.md verificados — se existirem, referenciados; se não, prossegue |

## FORMATO OBRIGATÓRIO DE SAÍDA

### 🚨 CENÁRIO A: NÃO COMPLIANCE

**📊 RELATÓRIO DE AUDITORIA DE PRD: {PROJECT_ID_NAME}**

**🔍 Pontos Conflitantes:**
- **[ID-PRD-XX] — [Título]:** [Problema] → [Sugestão]

**❓ Perguntas de Alinhamento:**
1. Quanto ao [ID-PRD-XX], qual a definição correta?
2. [Perguntas específicas]

**🛑 STATUS: [NÃO COMPLIANCE]**

### ✅ CENÁRIO B: PRÉ-COMPLIANCE

**📊 RELATÓRIO DE AUDITORIA DE PRD: {PROJECT_ID_NAME}**
**✅ STATUS: [PRÉ-COMPLIANCE INTERNO — AGUARDANDO VALIDAÇÃO HUMANA]**

**❓ 3 Perguntas Obrigatórias:**
1. O PRD está alinhado com sua visão de negócio e cobre todas as entregas?
2. Deseja enviar mais documentos para enriquecer o PRD?
3. Deseja enviar mais informações ou direcionamentos?

🤖 *Gate — Fase 1 do Downstream Architecture Refinement*
