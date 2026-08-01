# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — ARCHITECTURE-DEFINITION (F2)
## Versão: 1.1 — Arquitetura Detail-Level (C4 L2/L3 + ADRs Detalhados) — Independente de Tecnologia

Atue como um Engenheiro de Sistemas e Arquiteto de Soluções, especializado em processos de Downstream Architecture Refinement e design detalhado de sistemas distribuídos.

## OBJETIVO

Produzir a definição de arquitetura em nível de implementação: C4 Level 2 (Container) e Level 3 (Component), ADRs detalhados com diagramas de sequência, padrões de código e matriz de integração refinada.

**Este documento é independente de tecnologias específicas.** Durante a análise da stack tecnológica do projeto, identifique as tecnologias utilizadas e busque skills relacionados a essas tecnologias para aprimorar as estimativas. Caso não encontre skills específicos, utilize skills generalistas de arquitetura e engenharia de sistemas.

## INPUTS

1. **Docs de negócio:** Features e User Stories
2. **PRD Detail-Level** (F1)
3. **Stack tecnológica:** `technical-definitions/PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md` (se existir)
4. **Repositórios de código:** `TECHNICAL_SOLUTION_PATH`

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-ARCHITECTURE-DEFINITION — Arquitetura Detail-Level

## 1. C4 Level 2 — Container Diagram
[Diagrama: containers identificados + tecnologias detectadas + protocolos de comunicação]

## 2. C4 Level 3 — Component Diagrams
[Para cada serviço principal: decomposição em componentes internos, responsabilidades, interfaces]

## 3. ADRs Detalhados
[5-8 ADRs com: contexto, decisão, alternativas consideradas, trade-offs, diagrama de sequência, consequências]
- ADR-001: Estratégia de Multi-Tenancy
- ADR-002: Gateway↔IAM Service-ID Validation
- ADR-003: Autorização via Claims Injection
- ADR-004: Estratégia de Audit Log
- ADR-005: API Versioning

## 4. Padrões de Código e Estrutura
[Convenções de projeto, estrutura de pacotes, design patterns aplicáveis, princípios SOLID]

## 5. Matriz de Integração
[Serviço → Serviço: protocolo, autenticação, formato, contratos de API]

## 6. Riscos Arquiteturais
```

### Skills Recomendados

**Skills generalistas de arquitetura (sempre aplicáveis):**
- `senior-architect`, `engineering-skills`, `engineering-advanced-skills`
- `architecture`, `software-architecture`, `architecture-patterns`
- `c4-container`, `c4-component`, `architecture-decision-records`
- `system-design`, `microservices-patterns`

**Skills tecnológicos (condicionais — buscar ao identificar a stack):**
- Ao identificar uma tecnologia específica durante a análise da stack, procure skills relacionados a essa tecnologia para aprimorar as estimativas
- Caso não encontre skills específicos para a tecnologia identificada, utilize os skills generalistas listados acima como referência

🤖 *Prompt gerador — Fase 2 do Downstream Architecture Refinement · Independente de Tecnologia*
