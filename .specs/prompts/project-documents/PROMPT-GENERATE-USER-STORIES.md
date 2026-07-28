# PROMPT: GERADOR MODULAR DE USER STORIES E ATUALIZADOR DA MATRIZ MESTRE (RTM)
## Arquivo: PROMPT-GENERATE-USER-STORIES.md
## Versão: 2.0 — Arquitetura de Documentação Decoplada (Modular)

Atue como um Product Owner Sênior e Especialista em Engenharia de Backlog Ágil, aplicando as competências de `014-agile-user-story`, `agile-ba-practices`, `acceptance-criteria`, `program-manager`, `senior-pm`, `delivery-manager`, `agile-coach`, `scrum-master` e os filtros pragmáticos corporativos de `caveman` e `caveman-review`. Sua missão é ler as funcionalidades da Fase 4 e realizar o detalhamento modular do backlog.

### 🛑 DIRETRIZES CRÍTICAS DA ARQUITETURA MODULAR:
1. **Geração de Arquivos Individuais:** Você NÃO deve gerar um único arquivo contendo todas as histórias do projeto. Para cada funcionalidade fornecida, gere documentos isolados e dedicados para cada User Story mapeada, seguindo o padrão de nomenclatura de arquivo: `US-FEAT-{codigo-feature}-{SSSS}-{nome-da-user-story}.md`, onde `SSSS` é uma sequência numérica **global** iniciando em `0001` e seguindo até `9999`, independente da feature ou épico a que a história pertence.
2. **Abstração Técnica de Negócio:** Mantenha o foco estrito na experiência do usuário final, jornadas corporativas e regras de tela funcionais. É terminantemente proibido o vazamento de jargões técnicos de TI (como tabelas SQL, frameworks front-end, endpoints de API ou infraestrutura de nuvem).
3. **Escrita Comportamental Exaustiva:** Cada arquivo individual de história deve conter critérios de aceitação refinados em cenários dinâmicos baseados no formato de negócio Gherkin (**Dado que**, **Quando**, **Então**), incluindo tratamentos de exceção de tela e estados secundários (ex: base de dados vazia, travamento ou carregamento por loaders).
4. **Atualização da Matriz Mestre Centralizada (RTM Final):** Paralelamente à criação dos arquivos de histórias, você deve gerar ou atualizar o arquivo central autônomo `05-USER-STORIES-{PROJECT_ID_NAME}.md`. A matriz segue o padrão enxugado com colunas `D# | EPIC-ID | FEATURE-ID | US-ID | US Descrição | STATUS | FASE | RNs`, onde `STATUS` é o Compliance Gate (🔴🟡🟢) e `FASE` é o ciclo de vida da User Story (NEW → DEPLOYED). O arquivo individual de cada história segue o padrão `US-FEAT-{codigo-feature}-{SSSS}-{nome-da-user-story}.md` onde `SSSS` é sequencial global 0001-9999.
5. **Análise de Conformidade Interna (Caveman Review):** Aplique um filtro mental pragmático durante a geração. Elimine qualquer termo rebuscado de tecnologia que mascare a falta de uma regra de negócio real.

---

### INSTRUÇÕES DE EXECUÇÃO:
1. **Consumo de Inputs:** Analise o arquivo de Funcionalidades (`04-FEATURES-{PROJECT_ID_NAME}.md`) congelado na Fase 4 e use as bases anteriores (Épicos, BRD e Charter) para garantir a consistência conceitual vertical.
2. **Processamento:** Itere sobre as capacidades mapeadas e produza as saídas estruturadas divididas em duas partes, conforme os templates puros detalhados abaixo.

---

### PARTE 1: ESTRUTURA E LAYOUT DO ARQUIVO INDIVIDUAL DA HISTÓRIA
(Gerar um arquivo contendo esta anatomia para cada User Story identificada. Caminho sugerido: `/user-stories/US-FEAT-{codigo-feature}-{SSSS}-{nome-da-user-story}.md`)

```markdown
# User Story: US-FEAT-{codigo-feature}-{SSSS} (ex: US-FEAT-EP-0001-0001-0001) — [Nome Curto da História]
<!-- O código da US segue o padrão US-FEAT-{codigo-feature}-{SSSS} onde SSSS é sequencial global (0001-9999). O arquivo físico é nomeado US-FEAT-{codigo-feature}-{SSSS}-{nome-da-user-story}.md -->

- **Projeto:** PRJ-[ÁREA]-2026-[NÚMERO]-[NOME-DO-PROJETO]
- **Mapeamento Ágil:** Épico [ID_ÉPICO] ➔ Feature [ID_FEATURE] ➔ User Story {codigo-US} (ex: US-FEAT-EP-0001-0001-0001)
- **Prioridade:** [Must Have / Should Have / Could Have]
- **Data-Alvo:** [Data de Entrega Planejada]
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** Em Revisão / Aguardando Validação Humana

---

## 1. Declaração da História (User Story Statement)

- **Como** [Persona de Usuário / Papel Corporativo],
- **quero** [realizar uma ação funcional ou interagir com uma capacidade de tela específica],
- **para** [obter o respectivo valor de negócio ou sanar a dor operacional descrita].

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal de Sucesso Comercial/Operacional]
- **Dado que** [contexto inicial de negócio, regras preexistentes ou permissões de perfil do usuário],
- **Quando** [o usuário aciona o gatilho, clica ou interage com a interface do portal],
- **Então** [o sistema deve processar os dados e retornar a resposta esperada pelo negócio].

### Cenário 2: [Fluxo Alternativo, de Exceção, Tratamento de Alerta ou Base Vazia]
- **Dado que** [contexto de ausência de registros ou indisponibilidade de critérios de negócio],
- **Quando** [o usuário tenta executar ou visualizar a funcionalidade em tela],
- **Então** [o sistema deve mitigar o cenário apresentando placeholders funcionais ou informativos claros].

---

## 3. Regras de Negócio de Tela Relacionadas
- **[ID_RN]:** [Descrever restrições operacionais específicas da tela, validações lógicas corporativas, valores preenchidos por padrão e comportamentos derivados do BRD].
```

---

### PARTE 2: LAYOUT DO ARQUIVO CENTRAL DE INDEXAÇÃO DA MATRIZ MESTRE
(Gerar ou atualizar o arquivo unificado. Caminho sugerido: `/05-USER-STORIES-{PROJECT_ID_NAME}.md`)

O arquivo central deve seguir a estrutura padronizada abaixo, com matriz enxugada (apenas IDs, sem colunas descritivas de Feature e Épico) e colunas de rastreabilidade `STATUS` (Compliance Gate) e `FASE` (Ciclo de Vida).

```markdown
# Matriz de Rastreabilidade de Escopo (RTM)

- **Projeto:** {PROJECT_ID_NAME} — [Descrição curta do produto]
- **Data de Geração:** [DATA]
- **Última Atualização:** [DATA]
- **Versão:** 1.0
- **Status de Auditoria:** [PASS / FAIL]

---

## Visão Geral da Cobertura

| Métrica | Valor |
|---|---|
| Total de Entregas (Charter) | [N] (D1-D[N]) |
| Total de Épicos | [N] (EP-0001 a EP-00NN) |
| Total de Funcionalidades | [N] |
| Total de User Stories | [N] |
| Cobertura D→Epic | [N]/[N] (100%) |
| Cobertura Epic→Feature | [N]/[N] (100%) |
| Cobertura Feature→US | [N]/[N] (100%) |
| Órfãos detectados | 0 |

---

## Matriz de Rastreabilidade Completa

> **STATUS** = Compliance Gate (🔴 NON-COMPLIANCE / 🟡 PENDING-REVIEW / 🟢 COMPLIANCE)
> **FASE** = Ciclo de vida da User Story (NEW → BIZ-REFINE → READY-TECH → TECH-REFINE → READY-DEV → IN-PROGRESS → CODE-REVIEW → QA → UAT → DONE → DEPLOYED → CANCELLED)
> **RNs** = Regras de Negócio associadas (preencher a partir dos arquivos de feature)

| D# | EPIC-ID | FEATURE-ID | US-ID | US Descrição | STATUS | FASE | RNs |
|:---|:---|:---|:---|:---|:---:|:---|:---|
| D1 | EP-0001 | FEAT-EP-0001-0001 | US-FEAT-EP-0001-0001-0001 | [Descrição curta da história] | 🟢 | `NEW` | — |

### Legenda das Fases

| FASE | Significado |
|:---|:---|
| `NEW` | Novo — ideia registrada, aguardando priorização |
| `BIZ-REFINE` | Em refinamento de negócio — PO detalha critérios |
| `READY-TECH` | Pronto para refinamento técnico — negócio concluiu |
| `TECH-REFINE` | Em refinamento técnico — time debate solução, estima |
| `READY-DEV` | Pronto para desenvolvimento — entra na fila do Sprint Backlog |
| `IN-PROGRESS` | Em desenvolvimento — codificando + testes unitários |
| `CODE-REVIEW` | Em revisão de código — PR aberto |
| `QA` | Em teste — QA valida critérios de aceite |
| `UAT` | Em homologação — PO ou cliente valida |
| `DONE` | Pronto — 100% do DoD cumprido |
| `DEPLOYED` | Em produção — disponível ao usuário final |
| `CANCELLED` | Cancelado — perdeu valor de negócio ou premissa inviável |

---

## Verificação de Cobertura de Entregas

| ID Entrega | Qtd. US | Status |
|:---|:---|:---|
| D1 — [Nome da Entrega] | [N] | ✅ Coberto |

---

## Verificação de Integridade Física

| Verificação | Resultado |
|:---|---|
| Links quebrados na RTM | 0 |
| Arquivos órfãos na pasta `user-stories/` | 0 |
| Features sem User Stories | 0 |
| Entregas do Charter sem cobertura | 0 |

---

## Documentos Relacionados

- [01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md](./01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md) — Fase 1
- [02-BRD-{PROJECT_ID_NAME}.md](./02-BRD-{PROJECT_ID_NAME}.md) — Fase 2
- [03-EPICS-{PROJECT_ID_NAME}.md](./03-EPICS-{PROJECT_ID_NAME}.md) — Fase 3
- [04-FEATURES-{PROJECT_ID_NAME}.md](./04-FEATURES-{PROJECT_ID_NAME}.md) — Fase 4
- [user-stories/](./user-stories/) — Fase 5 ([N] arquivos individuais)

---

**Status Final:** `[COMPLIANCE]` — Matriz validada. Projeto 100% rastreado do Charter às User Stories.
```

---
`[STATUS: SUCESSO - ARQUIVOS INDIVIDUAIS E MATRIZ RTM GERADOS PARA AUDITORIA]`
