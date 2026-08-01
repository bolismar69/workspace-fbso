# PROMPT-GENERATE-SOURCING-FACTORY-BIDDING-FACTORY-DISTRIBUTION

## Contexto

Este prompt implementa o **GENERATE do FACTORY-DISTRIBUTION** para o processo de Sourcing & Factory Bidding (Fase 3).

**Propósito:** Registra as fábricas participantes com dados de contato (Canal, E-mail, Telefone), controla envio do RFQ e prazos de resposta. Colunas: Fábrica, Data Envio, Canal, E-mail, Telefone, Prazo, Status.

**Modo de operação:** Adapta-se ao `SOURCING_BIDDING_MODE` definido no Bootstrap (`discovery` ou `full`).

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{SOURCING_BIDDING_MODE}` | Modo: `discovery` ou `full` |
| `{SOURCING_BIDDING_PATH}` | Pasta sourcing-factory-bidding-{mode} |
| `{ESTIMATES_PATH}` | Pasta de estimativas recebidas |

## Fluxo de Execução

### Passo 0 — Validar Parâmetros e Modo
Confirmar `{PROJECT_PATH}`, `{PROJECT_ID_NAME}`, `{SOURCING_BIDDING_MODE}`.

### Passo 1 — Carregar Artefatos Base
- RFQ-PACKAGE.md (F1) — lista de artefatos enviados
- ESTIMATION-SCHEMA.csv (F2) — template enviado

### Passo 2 — Invocar Skills Especializadas
- `project-estimation` — Gestão do processo de distribuição
- `documentation-writer` — Registro de fábricas

### Passo 3 — GENERATE o Artefato

**Especificações do Artefato:**

1. **Tabela de fábricas** com colunas: `# | Fábrica | Data Envio | Canal | E-mail | Telefone | Prazo | Status`
2. **Mínimo 2 fábricas** cadastradas
3. **E-mail e Telefone** preenchidos com ⚠️ se não informados (placeholder para time operacional)
4. **Material Enviado:** lista RFQ-PACKAGE.md + ESTIMATION-SCHEMA.csv + artefatos upstream
5. **Instruções para o Time Operacional** com passos claros

### Passo 4 — Validação Pós-GENERATE
Verificar: ≥ 2 fábricas, colunas Canal/E-mail/Telefone presentes, prazos definidos.

## Skills Utilizados

| 1 | `project-estimation` | Gestão do processo de distribuição | 2 | `documentation-writer` | Registro de fábricas |

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 31/07/2026 | Criação inicial — Fase 3 Sourcing & Factory Bidding | Time de Arquitetura |

🤖 *Sourcing & Factory Bidding — Fase 3 GENERATE*
