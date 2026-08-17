# PROMPT-GENERATE-460-TEAM-CAPACITY

## Contexto

Este prompt gera o artefato `460-TEAM-CAPACITY.md` — a **tabela de capacidade do time técnico** que registra os profissionais que efetivamente executarão as tarefas de implementação das soluções técnicas do projeto (Bloco A — Fase 6).

Diferente do `450-TEAM-SKILLS-MAP.md` (Fase 5 — que foca em O QUE cada perfil sabe fazer — skills e níveis de proficiência), este documento foca em QUEM está disponível, quantas horas/dia e por qual período.

**Inputs do Bloco 0:** PRODUCT-BACKLOG-LIST.md (F3) e PRD-DEFINITION.md (F4) — referência de escopo do projeto para dimensionar a capacidade necessária.

**Princípio fundamental:** O arquivo é criado com a tabela vazia (sem integrantes). O prompt questiona o humano sobre a existência de inputs. Se fornecidos, os integrantes são preenchidos. Caso contrário, o arquivo é gerado sem integrantes e o preenchimento ocorrerá posteriormente.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Verificar se TODOS os parâmetros foram informados.

### Passo 1 — Questionar o Humano sobre Inputs
Antes de gerar o arquivo, perguntar ao usuário:

> **👥 Time Técnico — Capacidade de Trabalho**
>
> O arquivo `460-TEAM-CAPACITY.md` será criado para registrar os integrantes do time técnico e sua capacidade de trabalho.
>
> Deseja fornecer informações sobre os integrantes do time neste momento?
> - **SIM** → Forneça os dados: Nome, Papel, Nível (★☆☆ a ★★★), Data Inicial, Data Final (se aplicável), Capacidade semanal (horas/dia para cada dia da semana: Seg a Dom), Contato (e-mail)
> - **NÃO** → O arquivo será criado com a tabela vazia. Os integrantes poderão ser preenchidos posteriormente.

### Passo 2 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/460-TEAM-CAPACITY.md` com a estrutura abaixo.

Se o humano forneceu integrantes no Passo 1, preencher as linhas da tabela. Caso contrário, gerar a tabela com a linha de exemplo comentada (ou vazia) e o status `⚠️ TODO — Aguardando definição dos integrantes`.

### Passo 3 — Validação Pós-Geração
Verificar: arquivo no caminho correto, estrutura de colunas completa, legenda presente, instruções de preenchimento documentadas.

---

## Estrutura do Arquivo a ser Gerado

```markdown
# Mapa do Time Técnico — Capacidade de Trabalho

- **Projeto:** {PROJECT_ID_NAME}
- **Versão:** 1.0
- **Data de Criação:** [DATA ATUAL]
- **Última Atualização:** [DATA ATUAL]
- **Status:** ⚠️ TODO — Aguardando definição dos integrantes técnicos que executarão as soluções

---

## Objetivo

Este documento apresenta a tabela de capacidade do time técnico — os profissionais que efetivamente executarão as tarefas de implementação das soluções técnicas do projeto. Ele serve como referência rápida para o Tech Lead e Coordenador alocarem trabalho nos sprints.

As colunas **Nome** e **Contato** devem ser preenchidas tão logo os profissionais sejam designados.

---

## Time Técnico x Capacidade de Trabalho

| Papel | Nível | Data Inicial | Data Final | Capacidade semana | Nome | Contato |
|-------|-------|--------------|------------|--------------------|------|---------|
| (a preencher) | — | — | — | 00/00/00/00/00/00/00 | (a designar) | (a designar) |

---

## Time Técnico x Capacidade de Trabalho Extraordinaria

| Papel | Nível | Data Inicial | Data Final | Capacidade extraordinária | Nome | Contato |
|-------|-------|--------------|------------|---------------------------|------|---------|
| (a preencher se necessário) | — | — | — | 00/00/00/00/00/00/00 | (a designar) | (a designar) |

- **Capacidade semana**: as horas informadas na tabela **Time Técnico x Capacidade de Trabalho Extraordinaria** substituem as horas existentes na tabela **Time Técnico x Capacidade de Trabalho** dentro do periodo informado

---

## Legenda

| Campo | Descrição |
|:---|:---|
| **Nome** | Nome completo do profissional designado (a preencher) |
| **Contato** | E-mail ou canal de comunicação corporativo (a preencher) |
| **Papel** | Função no projeto |
| **Nível** | Proficiência esperada: ★★★ Senior/Especialista/Avançado/Autônomo, ★★☆ Pleno/Intermediário/Produtivo, ★☆☆ Junior/Básico/Assistido |
| **Capacidade semana** | Cada posição representa as horas possiveis de trabalhar em cada dia da semana: 1a-Segunda, 2a-Terça, 3a-Quarta, 4a-Quinta, 5a-Sexta, 6a-Sábado, 7a-Domingo |

---

## Instruções de Preenchimento

1. **Nome e Contato:** Preencher tão logo o profissional seja designado. O Coordenador do Projeto é responsável por manter esta planilha atualizada.
2. **Substituições:** Em caso de substituição de profissional, atualizar a linha correspondente e registrar abaixo.

### Histórico de Alterações

| Data | Alteração | Responsável |
|:---|:---|:---|
| [DATA ATUAL] | Criação do documento (v1.0) | Time de Arquitetura |
```

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `team-composition-analysis` | Analisar composição do time | People |
| 2 | `documentation-writer` | Redigir o TEAM-CAPACITY.md | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 28/07/2026 | Criação inicial: prompt gerador da capacidade do time técnico | Time de Arquitetura |
| 2.0 | 30/07/2026 | Atualização Bloco A (F5-F6): adicionado contexto de fase (Fase 6); adicionados inputs PRODUCT-BACKLOG-LIST (F3) e PRD-DEFINITION (F4) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
