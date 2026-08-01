# PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-NOTIFICATION

## Contexto

Este prompt implementa o **GENERATE do FACTORY-NOTIFICATION** — Fase 7 do Sourcing & Factory Bidding.

**Propósito:** Gerar notificações formais para todas as fábricas participantes informando o resultado do processo de seleção: carta de seleção para a vencedora, feedback técnico para as rejeitadas.

**Inputs upstream:** FACTORY-COMPARISON (F6) + ESTIMATE-VALIDATION (F5) + FACTORY-DISTRIBUTION (F3) + ESTIMATE-RETROSPECTIVE-PIB (F5b, se aplicável).

## Padrão de Geração

### Pasta de Destino
As notificações são geradas em: `{SOURCING_BIDDING_PATH}/notifications/`

### Padrão de Nomenclatura
```
FACTORY-NOTIFICATION-{NOME-DA-FABRICA}.md
```
**Exemplos:** `FACTORY-NOTIFICATION-STEFANINI.md`, `FACTORY-NOTIFICATION-CAPGEMINI.md`

> ⚠️ **Confidencialidade:** O status da fábrica (selecionada, rejeitada, segundo colocado) NUNCA aparece no nome do arquivo. Essa informação é confidencial e consta apenas dentro do conteúdo do arquivo.

### Conteúdo de Cada Notificação
- **Fábrica selecionada:** Carta de Seleção com próximos passos (contrato, kickoff, acesso)
- **Segundo colocado:** Carta de Agradecimento — proposta aprovada tecnicamente, manter contato para futuras oportunidades
- **Fábricas rejeitadas:** Feedback técnico com motivo específico da rejeição (F5) e recomendação para futuras participações

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `discovery` ou `full` |
| `{SOURCING_BIDDING_PATH}` | Pasta sourcing-factory-bidding-{mode} |

## Fluxo de Execução

### Passo 0 — Validar Parâmetros
### Passo 1 — Carregar FACTORY-COMPARISON (F6) + ESTIMATE-VALIDATION (F5) + FACTORY-DISTRIBUTION (F3)
### Passo 2 — Criar pasta `notifications/` se não existir
### Passo 3 — Para cada fábrica na F3, gerar arquivo `FACTORY-NOTIFICATION-{NOME}.md` com o conteúdo adequado ao status (selecionada/segundo/rejeitada)
### Passo 4 — Validação Pós-Geração: verificar que 100% das fábricas têm notificação e que nenhum nome de arquivo revela o status

## Skills Utilizados

| 1 | `documentation-writer` | Redação das cartas de notificação |
| 2 | `business-analyst` | Linguagem adequada ao relacionamento com fornecedores |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 7 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 7 GENERATE*
