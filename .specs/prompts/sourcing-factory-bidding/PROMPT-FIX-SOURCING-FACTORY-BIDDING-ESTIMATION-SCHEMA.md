# PROMPT-FIX-SOURCING-FACTORY-BIDDING-ESTIMATION-SCHEMA
## Contexto
Este prompt implementa o **FIX do ESTIMATION-SCHEMA** — Fase 2. Acionado quando o GATE encontra NCs.
**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as NCs apontadas. NUNCA reescreve o documento inteiro.
## Fluxo de Correção
### Passo 0 — Carregar Relatório do GATE (IDs de conflito, localizações, sugestões)
### Passo 1 — Priorizar: P0 Bloqueante, P1 Importante, P2 Menor
### Passo 2 — Aplicar Correção Cirúrgica (menor mudança possível que resolve a NC)
### Passo 3 — Validar (todas NCs endereçadas; se não resolvida → [NÃO RESOLVIDA] com justificativa)
**Regra de Ouro:** Nunca reescreva o documento do zero.
## Skills
| 1 | `documentation-writer` | Correção cirúrgica | 2 | `code-reviewer` | Revisão das correções |
🤖 *Fase 2 FIX*
