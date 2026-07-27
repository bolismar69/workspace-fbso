# PROMPT: GERADOR DE ESPECIFICAÇÃO DE FUNCIONALIDADES (FEATURES PRD)
## Arquivo: PROMPT-GENERATE-FEATURES.md
## Versão: 3.0 — Padrão de ID FEAT-EP-{EEEE}-{SSSS} com Arquivos Individuais por Feature

Atue como um Analista de Negócios Ágil Sênior (Agile BA) e Product Owner, aplicando as competências de `agile-ba-practices`, `acceptance-criteria`, `breakdown-feature-prd` e os filtros pragmáticos corporativos de `caveman` e `caveman-review`. Sua missão é ler o documento de Épicos congelado da Fase 3 e desdobrá-lo em uma **Especificação de Funcionalidades (Features)** exaustiva, estruturada de forma modular: um arquivo índice central + arquivos individuais por feature na pasta `features/`.

### 🛑 DIRETRIZES CRÍTICAS DE HERANÇA E ESCOPO:
1. **Estrutura Modular (NOVO — v2.0):** O documento de Features passa a ser entregue em dois artefatos complementares:
   - **Arquivo Índice:** `04-FEATURES-{PROJECT_ID_NAME}.md` — resumo executivo com visão consolidada, cronograma, matriz de cobertura Charter×Features, matriz MoSCoW e diagrama de linha de tempo.
   - **Arquivos Individuais:** `features/FEAT-EP-{EEEE}-{SSSS}-{nome-slugificado}.md` — um arquivo por feature com objetivo de negócio, user stories, regras de negócio e matriz de rastreabilidade específica.
2. **Taxonomia e Rastreabilidade de IDs (ATUALIZADO — v3.0):** As funcionalidades devem utilizar o formato `FEAT-EP-{EEEE}-{SSSS}` onde `{EEEE}` é o código do épico associado com 4 dígitos e `{SSSS}` é o sequencial dentro do épico com 4 dígitos (reiniciando para cada épico). Ex: `FEAT-EP-0001-0001`, `FEAT-EP-0001-0002`, `FEAT-EP-0002-0001`. As User Stories associadas mantêm o formato `US-[SEQUENCIAL]` (ex: `US-001`). As Regras de Negócio passam a utilizar o formato `RN-FEAT-EP-{EEEE}-{SSSS}-{SSSS}` (ex: `RN-FEAT-EP-0001-0001-0001`).
3. **Abstracão Técnica Absoluta:** O foco é o comportamento esperado sob a ótica do usuário e do negócio. Não inclua referências a arquitetura técnica, bancos de dados, microsserviços ou tecnologias front-end/back-end.
4. **Mapeamento Exigido (no Índice):** O documento deve conter obrigatoriamente: Visão Geral de Funcionalidades com links para arquivos individuais, Cronograma por Funcionalidade, Matriz de Cobertura contra o Project Charter, Diagrama de Linha de Tempo ASCII e Matriz de Priorização MoSCoW.
5. **Análise de Conformidade Interna (Caveman Review):** Aplique um filtro mental pragmático durante a geração. Elimine qualquer termo rebuscado de tecnologia que mascare a falta de uma regra de negócio real.
6. **Links Cruzados Obrigatórios:**
   - O índice deve linkar para cada arquivo individual: `[FEAT-EP-0001-0001](features/FEAT-EP-0001-0001-nome.md)`
   - Cada arquivo individual deve linkar de volta para o índice: `[Índice de Features](../04-FEATURES-{PROJECT_ID_NAME}.md)`
   - Cada arquivo individual deve linkar para o épico associado: `[EP-0001](../epics/EP-0001-nome.md)`
   - Cada arquivo individual deve referenciar o BRD: `[BR-XX](../02-BRD-{PROJECT_ID_NAME}.md)`

---
### INSTRUÇÕES DE EXECUÇÃO:
1. **Inputs:** Consuma o arquivo `03-EPICS-{PROJECT_ID_NAME}.md` (Fase 3), o `02-BRD-{PROJECT_ID_NAME}.md` (Fase 2) e o `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` (Fase 1).
2. **Criar Estrutura de Diretórios:** Execute `mkdir -p {PROJECT_COMPLETE_PATH_NAME}/features/` para garantir que a pasta de artefatos modulares exista.
3. **Decomposição:** Identifique as funcionalidades necessárias para cobrir as jornadas e requisitos descritos. Atribua códigos sequenciais `FEAT-EP-{EEEE}-{SSSS}` a cada feature, onde `{EEEE}` é o código do épico pai e `{SSSS}` é o sequencial reiniciado para cada épico.
4. **Geração dos Arquivos Individuais:** Para cada feature, crie um arquivo na pasta `features/` com o nome `FEAT-EP-{EEEE}-{SSSS}-{nome-slugificado}.md` seguindo o template de detalhamento (objetivo, user stories, regras de negócio, matriz de rastreabilidade).
5. **Geração do Arquivo Índice:** Crie o arquivo `04-FEATURES-{PROJECT_ID_NAME}.md` na raiz do projeto com o conteúdo consolidado.
6. **Formatação:** Retorne os documentos finais estruturados com os layouts abaixo.

---
### ESTRUTURA E LAYOUT DO ARQUIVO ÍNDICE (`04-FEATURES-{PROJECT_ID_NAME}.md`):

```
# Funcionalidades do Projeto: [Inserir Nome do Projeto]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | FEATURES-{PROJECT_ID_NAME} |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | [Data Atual] |
| **Origem** | `03-EPICS-{PROJECT_ID_NAME}.md` v[X.X] |
| **Status** | Em Revisão / Aguardando Validação |

---

## Visão Geral das Funcionalidades

[Construir tabela cruzando todas as funcionalidades: ID com link para arquivo individual, Funcionalidade, Épico Pai, Prioridade MoSCoW, User Stories, Data-Alvo].

| ID | Funcionalidade | Épico | Prioridade | User Stories | Data-Alvo |
|----|---------------|-------|------------|-------------|-----------|
| **FEAT-EP-0001-0001** | [Nome da Feature 1](features/FEAT-EP-0001-0001-nome-1.md) | EP-0001 | Must Have | N | **DD/MM/AAAA** |

**Total: [X] funcionalidades | [Y] user stories**

> 📄 **Detalhamento completo** de cada feature disponível na pasta [`features/`](features/).

### Mapeamento de Numeração (se aplicável)

[Tabela de ponte entre IDs antigos (ex: F01-01) e novos (FEAT-EP-0001-0001), útil durante a transição].

### Cronograma de Entregas por Funcionalidade

[Construir tabela ordenada cronologicamente por: Data-Alvo, Marco, Épico, Funcionalidades].

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Features

| BRD | Requisito Funcional | Épico/Jornada | Feature |
|:---|:---|:---|:---|
| **BR-01** | [Descrição] | EP-0001 / J1: [Nome] | [FEAT-EP-0001-0001](features/FEAT-EP-0001-0001-nome.md) |

---

## Matriz de Cobertura: Entregas do Project Charter × Features

| Entrega (Project Charter) | Funcionalidades Relacionadas |
|---------------------------|------------------------------|
| D1 — [Nome da Entrega] | [FEAT-EP-0001-0001](features/...), [FEAT-EP-0001-0002](features/...) |

[Diagrama de Linha de Tempo ASCII consolidando datas, Marcos, Épicos e Features].

---

## Matriz de Priorização (MoSCoW)

| Prioridade | Funcionalidades | Quantidade |
|-----------|----------------|------------|
| **Must Have** | [Lista de FP IDs] | N |
| **Should Have** | [Lista de FP IDs] | N |
| **Could Have** | — | 0 |
| **Won't Have (esta fase)** | [Descrição do excluído] | — |

---

> **Este documento é um índice resumido.** O detalhamento completo de cada feature — incluindo objetivo de negócio, user stories, critérios de aceitação e regras de negócio — está nos arquivos individuais da pasta [`features/`](features/).

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: agile-ba-practices, acceptance-criteria, breakdown-feature-prd. Estrutura modular v2.0.*
```

---
### ESTRUTURA E LAYOUT DE CADA ARQUIVO INDIVIDUAL (`features/FEAT-EP-{EEEE}-{SSSS}-{nome-slugificado}.md`):

> **Nota:** O slug do nome deve ser derivado do título da feature, em kebab-case. Ex: "Dashboard de Métricas Operacionais" → `dashboard-metricas-operacionais`.

```
# FEATURE - FEAT-EP-{EEEE}-{SSSS}: [Inserir Nome da Feature]

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-{EEEE}-{SSSS} — [Nome da Feature] |
| **Épico** | [EP-{EEEE} — Nome do Épico](../epics/EP-{EEEE}-nome-do-epico.md) |
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | FEATURES-{PROJECT_ID_NAME} |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades |
| **Data** | [Data Atual] |
| **Origem** | `03-EPICS-{PROJECT_ID_NAME}.md` e `02-BRD-{PROJECT_ID_NAME}.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-{PROJECT_ID_NAME}.md`](../04-FEATURES-{PROJECT_ID_NAME}.md) | **Épico:** [EP-{EEEE}](../epics/EP-{EEEE}-nome-do-epico.md) | **Anterior:** [FEAT-EP-{EEEE}-{SSSS} — Nome](../FEAT-EP-{EEEE}-{SSSS}-anterior.md) | **Próximo:** [FEAT-EP-{EEEE}-{SSSS} — Nome](../FEAT-EP-{EEEE}-{SSSS}-proximo.md)

**Requisitos BRD Vinculados:** [BR-XX](../02-BRD-{PROJECT_ID_NAME}.md) — [Nome do Requisito]

---

## Objetivo de Negócio
[Explicar resumidamente o valor comercial, a dor que alivia e a necessidade operacional desta funcionalidade específica].

**Prioridade:** [Must Have / Should Have / Could Have]

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-XXX | Como **[Persona]**, quero [ação/funcionalidade] para [benefício/valor de negócio] | • [Critério 1]<br>• [Critério 2] |

## Regras de Negócio

- **RN-FEAT-EP-{EEEE}-{SSSS}-{SSSS}:** [Descrever restrições operacionais específicas, validações lógicas, regras de exibição e valores padrão].

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-XX** | [Descrição] | [EP-{EEEE}](../epics/EP-{EEEE}-nome.md) / J[N]: [Nome da Jornada] | **FEAT-EP-{EEEE}-{SSSS}** — [Nome da Feature] |

---

> 📄 **Índice de Features:** [`04-FEATURES-{PROJECT_ID_NAME}.md`](../04-FEATURES-{PROJECT_ID_NAME}.md) | **Épico:** [EP-{EEEE}](../epics/EP-{EEEE}-nome-do-epico.md)
```

---
### PÓS-GERAÇÃO:

Após criar todos os arquivos individuais e o índice, execute uma verificação de integridade:
1. Confirme que todos os links no índice apontam para arquivos existentes na pasta `features/`
2. Confirme que todos os arquivos individuais têm link de volta para o índice
3. Confirme que cada arquivo referencia corretamente o épico associado e os BRDs aplicáveis
4. Confirme que a matriz consolidada no índice cobre 100% das features
5. Confirme que o formato `FEAT-EP-{EEEE}-{SSSS}` é consistente em todos os lugares

Ao final de cada arquivo gerado (índice e individuais), insira a tag:
`[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]`
