# PROMPT: GERADOR DE STAKEHOLDER MAP (REGISTRO COMPLETO DE PARTES INTERESSADAS)
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Analista de Negócios Sênior e Especialista em Gestão de Stakeholders. Sua missão é criar um **Stakeholder Map** completo para o projeto, expandindo a versão simplificada do Project Charter.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `TEAM_SKILLS` | Skills mapeados para o time de implementação (`PROJECT-TEAM-SKILLS-MAP`) |
| `TEAM_CAPACITY` | Capacidade do time (`PROJECT-TEAM-CAPACITY`) |
| `UPSTREAM_DOCS` | Lista: `[{PROJECT_COMPLETE_PATH_NAME}/001-PROJECT-CHARTER-{PROJECT_ID_NAME}.md]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["stakeholder-analysis", "stakeholder-map", "agile-ba-practices"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** o PROJECT-CHARTER em `UPSTREAM_DOCS[0]` — extraia a lista de stakeholders da Seção 5 e expanda cada um com informações detalhadas
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback (5 Seções)

```
# Stakeholder Map: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento Base** | 001-PROJECT-CHARTER-{PROJECT_ID_NAME}.md |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## O que é um Stakeholder Map?

O **Stakeholder Map (Mapa de Partes Interessadas)** é o registro canônico de todas as pessoas, grupos e organizações que têm interesse ou influência sobre o projeto. Ele expande a Seção 5 do Project Charter com informações detalhadas de contato, responsabilidades decisórias, canais de comunicação e caminhos de escalação.

### O que contém

- **Identificação completa** de cada stakeholder com nome, papel, posição organizacional e contato
- **Matriz RACI por fase do projeto** — granularidade por atividade, não apenas por entrega macro
- **Canais de comunicação** com frequência, participantes e artefatos de saída
- **Caminho de escalação** para cada tipo de impedimento (negócio, operacional, regulatório)

### Conexão com o Pipeline

- **UPSTREAM:** Consome a lista de stakeholders e RACI macro do 001-PROJECT-CHARTER
- **DOWNSTREAM:** Alimenta 005-BRD (mapeamento detalhado de stakeholders), 075-PLANO-COMUNICACAO (canais e frequência)

---

## 1. Identificação das Partes Interessadas

### 1.1 Patrocinadores Executivos (Sponsors)

| Papel | Nome | Posição | Decide sobre | Contato |
|-------|------|---------|--------------|---------|
| Sponsor — {cargo} | `<nome>` | {posição hierárquica} | {decisões sob sua alçada} | `<email>` |

### 1.2 {Categoria — ex: Governança, Compliance, etc.}

| Papel | Nome | Posição | Decide sobre | Contato |
|-------|------|---------|--------------|---------|
| {papel} | `<nome>` | {posição} | {decisões} | `<email>` |

### 1.3 {Categoria — ex: Usuários Finais, Time de Negócio}

| Papel | Nome | Posição | Decide sobre | Contato |
|-------|------|---------|--------------|---------|
| {papel} | `<nome>` | {posição} | {decisões} | `<email>` |

### 1.4 Lideranças de Negócio Impactadas

| Papel | Nome | Posição | Decide sobre | Contato |
|-------|------|---------|--------------|---------|
| {papel} | `<nome>` | {posição} | {decisões} | `<email>` |

### 1.5 Execução e Governança do Projeto

| Papel | Nome | Posição | Decide sobre | Contato |
|-------|------|---------|--------------|---------|
| Product Owner (PO) | `<nome>` | {posição} | Priorização, critérios de aceite, trade-offs escopo×prazo, aceitação formal | `<email>` |
| PMO | `<nome>` | {posição} | Alinhamento cronograma, arbitragem de conflitos, escalação | `<email>` |

> **NOTA:** As categorias acima (1.1 a 1.5) são sugestivas. Adapte as categorias ao contexto do projeto. O importante é que cada stakeholder identificado no Charter apareça aqui com informações completas.

---

## 2. Matriz RACI por Fase do Projeto

**Legenda:** R = Responsible (executa) | A = Accountable (aprova/responde) | C = Consulted (consultado) | I = Informed (informado)

### 2.1 Fase 1 — Iniciação e Requisitos de Negócio

| Atividade | {Stakeholder 1} | {Stakeholder 2} | ... |
|-----------|-----------------|-----------------|-----|
| {atividade} | R/A/C/I | ... | ... |

### 2.2 Fase 2 — Especificação de Sistema e Arquitetura

| Atividade | {Stakeholder 1} | {Stakeholder 2} | ... |
|-----------|-----------------|-----------------|-----|

### 2.3 Fase 3 — Engenharia Detalhada e Qualidade

| Atividade | {Stakeholder 1} | {Stakeholder 2} | ... |
|-----------|-----------------|-----------------|-----|

### 2.4 Fase 4 — Planejamento e Baseline

| Atividade | {Stakeholder 1} | {Stakeholder 2} | ... |
|-----------|-----------------|-----------------|-----|

### 2.5 Fase 5 — Encerramento e Operação

| Atividade | {Stakeholder 1} | {Stakeholder 2} | ... |
|-----------|-----------------|-----------------|-----|

---

## 3. Canais de Comunicação e Frequência

| Fórum | Participantes | Frequência | Objetivo | Artefato de Saída |
|-------|---------------|------------|----------|-------------------|
| **Comitê Executivo do Projeto** | {stakeholders} | Mensal | Aprovar direcionamento, liberar recursos | Dashboard Executivo |
| **Reunião de {tema}** | {stakeholders} | {frequência} | {objetivo} | {artefato} |

---

## 4. Caminho de Escalação (Escalation Path)

```
[IMPEDIMENTO DE NEGÓCIO]
    │
    ├─ 1. Resolver com Product Owner (PO)
    │
    ├─ 2. Se requer decisão de negócio:
    │      Escalar para {Comitê/Sponsor}
    │
    └─ 3. Se impacta orçamento ou diretriz estratégica:
           Escalar para PMO → Sponsor Principal

[IMPEDIMENTO OPERACIONAL]
    │
    ├─ 1. Resolver com {responsável operacional}
    │
    ├─ 2. Se requer mudança de processo:
    │      Escalar para {instância superior}
    │
    └─ 3. Se impacta adoção ou satisfação:
           Escalar para PO → PMO → Sponsor

[IMPEDIMENTO REGULATÓRIO]
    │
    ├─ 1. {instância técnica} analisa impacto
    │
    ├─ 2. Emite parecer com recomendação
    │
    └─ 3. Se impacto crítico:
           PO convoca Comitê Executivo extraordinário
```

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Criação inicial a partir do Project Charter | Time de Negócios |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se as 5 seções estiverem completas, todos os stakeholders do Charter estiverem expandidos, e o escalation path cobrir os 3 tipos de impedimento.
