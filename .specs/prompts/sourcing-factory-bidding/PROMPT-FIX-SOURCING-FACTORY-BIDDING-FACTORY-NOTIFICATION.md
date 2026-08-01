# PROMPT-FIX-SOURCING-FACTORY-BIDDING-FACTORY-NOTIFICATION

## Contexto

Este prompt implementa o **FIX do FACTORY-NOTIFICATION** — Fase 7. Acionado quando o GATE encontra NCs.

**Postura do FIX:** Cirúrgico e contido. Corrige APENAS as NCs apontadas pelo GATE. NUNCA reescreve o documento inteiro.

**Padrão:** Arquivos em `notifications/FACTORY-NOTIFICATION-{NOME}.md`. Nomes NUNCA revelam status (selecionada/rejeitada).

**Input:** Relatório de NCs do GATE (IDs [NOTIF-XX], localizações, sugestões).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos |
| `{PROJECT_ID_NAME}` | Identificador do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `discovery` ou `full` |

## Fluxo de Correção

### Passo 0 — Carregar Relatório do GATE
IDs de conflito, localizações e sugestões.

### Passo 1 — Priorizar NCs

| Prioridade | Tipo | Ação |
|---|---|---|
| P0 | Nome de arquivo revela status | Renomear removendo "selecionada"/"rejeitada" do nome |
| P0 | Fábrica sem notificação | Gerar arquivo faltante |
| P1 | Feedback técnico genérico (sem motivo F5) | Incluir motivo específico da F5 |
| P1 | Observações da F5b não incorporadas | Incluir observações da retrospectiva |
| P2 | Tom inadequado (vencedora muito informal, rejeitada muito seco) | Ajustar tom |
| P2 | Data de envio, canal, e-mail ausentes | Preencher dados da F3 |

### Passo 2 — Aplicar Correções Cirúrgicas
Para cada NC: localizar seção exata → aplicar menor correção → verificar confidencialidade.

### Passo 3 — Validar Correções
Todas NCs endereçadas. Se não resolvida → `[NÃO RESOLVIDA]` com justificativa. Verificar que nenhum nome de arquivo revela status.

**Regra de Ouro:** Nunca reescreva notificações do zero. Confidencialidade é inegociável.

## Skills

| 1 | `documentation-writer` | Correção de cartas |
| 2 | `business-analyst` | Linguagem de relacionamento |
| 3 | `code-reviewer` | Revisão de confidencialidade |

🤖 *Fase 7 FIX*
