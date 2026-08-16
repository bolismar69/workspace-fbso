# PROMPT: GERADOR DE PLANO DE GESTÃO DE MUDANÇAS
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Gerente de Projetos Sênior e Especialista em Gestão de Mudanças. Sua missão é criar o **Plano de Gestão de Mudanças** que define como mudanças de escopo, cronograma e orçamento serão solicitadas, classificadas, avaliadas e aprovadas após a baseline do projeto.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `TEAM_SKILLS` | Skills mapeados para o time de implementação (`PROJECT-TEAM-SKILLS-MAP`) — base para composição do CCB |
| `UPSTREAM_DOCS` | Lista: `[{PROJECT_COMPLETE_PATH_NAME}/001-PROJECT-CHARTER-{PROJECT_ID_NAME}.md, {PROJECT_COMPLETE_PATH_NAME}/060-EAP-WBS-{PROJECT_ID_NAME}.md, {PROJECT_COMPLETE_PATH_NAME}/080-PLANO-RISCOS-{PROJECT_ID_NAME}.md]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["change-management", "senior-pm"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** 001-PROJECT-CHARTER (escopo, entregas e partes interessadas decisórias), 060-EAP-WBS (pacotes de trabalho impactáveis) e 080-PLANO-RISCOS (interação mudança × risco)
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback (5 Seções)

```
# Plano de Gestão de Mudanças: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 001-PROJECT-CHARTER, 060-EAP-WBS, 080-PLANO-RISCOS |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## O que é um Plano de Gestão de Mudanças?

O **Plano de Gestão de Mudanças (Change Management Plan)** define o processo formal para solicitar, avaliar, aprovar e registrar qualquer mudança em escopo, cronograma, orçamento ou requisitos após a linha de base do projeto. Ele protege a baseline contra o scope creep e garante que decisões sejam tomadas com impacto conhecido.

### Conexão com o Pipeline

- **UPSTREAM:** Consome o escopo e as partes interessadas decisórias do 001-PROJECT-CHARTER, os pacotes de trabalho da 060-EAP-WBS e o inventário de riscos da 080-PLANO-RISCOS
- **DOWNSTREAM:** Rege toda alteração subsequente — qualquer mudança aprovada dispara os efeitos cascata (regeneração + revalidação) dos documentos impactados

---

## 1. Formulário de Solicitação de Mudança (Change Request)

| Campo | Descrição |
|-------|-----------|
| **ID da Solicitação** | `CCR-{NN}` — sequencial |
| **Data** | {DATA DA SOLICITAÇÃO} |
| **Solicitante** | {nome, papel, área} |
| **Tipo de Mudança** | Escopo / Cronograma / Orçamento / Requisito / Correção / Outro |
| **Descrição** | O quê, por quê, onde (pacote de trabalho, requisito, entrega afetada) |
| **Justificativa de Negócio** | {benefício esperado, urgência} |
| **Impacto Estimado** | {escopo, prazo, custo, risco, qualidade} |
| **Anexos** | {documentos de apoio} |

---

## 2. Composição do Change Control Board (CCB)

| Papel | Membro (Nome/Cargo) | Voto | Responsabilidade |
|-------|---------------------|------|------------------|
| Presidente do CCB | {Sponsor ou PMO} | Decisivo | Convocar, arbitrar, aprovar/rejeitar |
| Gestor do Projeto (PM) | {PM} | Consultivo | Apresentar análise de impacto |
| Representante de Negócio | {PO / Sponsor de negócio} | Decisivo | Avaliar valor e alinhamento estratégico |
| Representante Técnico | {Líder técnico / Arquiteto} | Consultivo | Avaliar viabilidade e esforço |
| Representante de Qualidade | {QA} | Consultivo | Avaliar impacto em testes e qualidade |
| Representante Financeiro | {controller} | Decisivo (orçamento) | Avaliar impacto financeiro |

> **NOTA:** O CCB é montado a partir do mapa de stakeholders do Charter e das skills do time disponível. Para mudanças de baixo impacto, um subconjunto do CCB pode deliberar.

---

## 3. Classificação de Impacto

### 3.1 Critérios de Classificação

| Nível | Critérios (atende a pelo menos um) |
|-------|-------------------------------------|
| **ALTO** | Impacta marco crítico do Charter / altera a linha de base de escopo (RTM) / custo acima de {X}% do orçamento / afeta requisito regulatório ou de segurança |
| **MÉDIO** | Altera cronograma sem quebrar marcos / custo entre {X}% e {Y}% / afeta pacote de trabalho específico da EAP |
| **BAIXO** | Sem impacto em escopo, prazo ou custo / correção de erro documental / ajuste menor de procedimento |

### 3.2 Matriz de Classificação

| Critério | Baixo | Médio | Alto |
|----------|-------|-------|------|
| Impacto em Escopo | Sem alteração | {faixa} | {faixa} |
| Impacto em Prazo | ≤ {X} dias | {faixa} | > {Y} dias / marco |
| Impacto em Custo | ≤ {X}% | {faixa} | > {Y}% |
| Impacto em Risco | Sem novos riscos | Novo risco gerenciável | Novo risco crítico |

---

## 4. Fluxo de Aprovação (Workflow)

```
[SOLICITANTE] Preenche o Formulário CCR
        │
        ▼
[PM] Recebe, registra no Change Log e avalia completude
        │
        ▼
[PM + Time] Análise de Impacto (escopo, prazo, custo, risco, qualidade)
        │
        ▼
[PM] Classifica: ALTO / MÉDIO / BAIXO
        │
        ├── BAIXO ──► Aprovação direta do PM ──► Implementar ──► Registrar
        │
        ├── MÉDIO ──► CCB (subconjunto) delibera ──► Aprovar / Rejeitar / Devolver
        │
        └── ALTO ──► CCB completo delibera ──► Sponsor decide ──► Aprovar / Rejeitar
                        │
                        ▼
              [APROVADO] Atualizar baseline + efeitos cascata (regenerar downstream)
              [REJEITADO] Notificar solicitante com justificativa
              [DEVOLVIDO] Solicitante ajusta e reenvia
```

### 4.1 Prazos de Decisão

| Classificação | Prazo Máximo de Decisão |
|---------------|--------------------------|
| BAIXO | {X} dias úteis |
| MÉDIO | {X} dias úteis |
| ALTO | {X} dias úteis (ou reunião extraordinária) |

---

## 5. Registro de Mudanças (Change Log)

| CCR | Data | Solicitante | Tipo | Classificação | Status | Aprovado por | Decisão | Data Decisão | Itens Impactados |
|-----|------|-------------|------|---------------|--------|--------------|---------|--------------|------------------|
| CCR-01 | {data} | {nome} | Escopo | MÉDIO | Aprovado | CCB | Aprovar | {data} | 060-EAP-WBS, 070-ORCAMENTO |
| CCR-02 | {data} | {nome} | Cronograma | BAIXO | Aprovado | PM | Aprovar | {data} | 065-CRONOGRAMA-GANTT |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Baseline inicial do plano de gestão de mudanças | Time de Planejamento |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se as 5 seções estiverem completas, o CCB estiver definido com papéis e responsabilidades, e o fluxo de aprovação cobrir as 3 classificações de impacto.
