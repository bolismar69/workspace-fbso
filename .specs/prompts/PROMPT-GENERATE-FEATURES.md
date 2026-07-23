# PROMPT: GERADOR DE ESPECIFICAÇÃO DE FUNCIONALIDADES (FEATURES PRD)
## Arquivo: PROMPT-GENERATE-FEATURES.md
## Versão: 1.0 — Layout Puro e Mapeamento de Transição Ágil

Atue como um Analista de Negócios Ágil Sênior (Agile BA) e Product Owner, aplicando as competências de `agile-ba-practices`, `acceptance-criteria`, `breakdown-feature-prd` e os filtros pragmáticos corporativos de `caveman` e `caveman-review`. Sua missão é ler o documento de Épicos congelado da Fase 3 e desdobrá-lo em uma **Especificação de Funcionalidades (Features)** exaustiva.

### 🛑 DIRETRIZES CRÍTICAS DE HERANÇA E ESCOPO:
1. **Taxonomia e Rastreabilidade de IDs:** Siga rigorosamente a taxonomia padrão. As funcionalidades devem utilizar o formato `F[ÉPICO]-[SEQUENCIAL]` (Ex: `F01-01`). As User Stories associadas de rascunho devem utilizar `US-[SEQUENCIAL]` (Ex: `US-001`). As Regras de Negócio específicas devem utilizar `RN[ÉPICO]-[SEQUENCIAL]` (Ex: `RN01-01`).
2. **Abstração Técnica Absoluta:** O foco é o comportamento esperado sob a ótica do usuário e do negócio. Não inclua referências a arquitetura técnica, bancos de dados, microsserviços ou tecnologias front-end/back-end.
3. **Mapeamento Exigido:** O documento deve conter obrigatoriamente: Visão Geral de Funcionalidades, Cronograma por Funcionalidade, Matriz de Cobertura contra o Project Charter, Diagrama de Linha de Tempo ASCII e a Matriz de Priorização MoSCoW.
4. **Análise de Conformidade Interna (Caveman Review):** Aplique um filtro mental pragmático durante a geração. Elimine qualquer termo rebuscado de tecnologia que mascare a falta de uma regra de negócio real.

---

### INSTRUÇÕES DE EXECUÇÃO:
1. **Inputs:** Consuma o arquivo `03-EPICS-{PROJECT_ID_NAME}.md` (Fase 3), o `02-BRD-{PROJECT_ID_NAME}.md` (Fase 2) e o `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` (Fase 1).
2. **Decomposição:** Identifique as funcionalidades necessárias para cobrir as jornadas e requisitos descritos, estimando o total de funcionalidades e histórias.
3. **Formatação:** Retorne o documento preenchendo estritamente o layout limpo abaixo com base nos dados do projeto do usuário.

---

### ESTRUTURA E LAYOUT PADRÃO DE SAÍDA (TEMPLATE PURO):

#### # Funcionalidades do Projeto: [Inserir Nome do Projeto]
##### Document ID: FEATURES-[SIGLA-DO-PROJETO]-2026-001

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-[ÁREA]-2026-[NÚMERO]-[NOME-DO-PROJETO] |
| **Documento** | FEATURES-[SIGLA-DO-PROJETO]-2026-001 |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades |
| **Data** | [Data Atual] |
| **Origem** | `EPICS-[SIGLA-DO-PROJETO]-2026-001` (Versão Validada) |
| **Status** | Em Revisão / Aguardando Validação |

---

## Visão Geral das Funcionalidades

[Construir tabela cruzando todas as funcionalidades geradas: ID, Funcionalidade, Épico Pai, Prioridade MoSCoW, Quantidade de User Stories estimadas, Data-Alvo alinhada aos marcos].

**Total: [X] funcionalidades | [Y] user stories**

### Cronograma de Entregas por Funcionalidade

[Construir tabela ordenando cronologicamente por: Data-Alvo, Marco, Épico, Funcionalidades].

---

## [ID_EPICO]: [Nome do Épico]

---

### [ID_FUNCIONALIDADE]: [Nome da Funcionalidade]

**Objetivo de Negócio:** [Explicar resumidamente o valor comercial, a dor que alivia e a necessidade operacional desta funcionalidade específica].

**Prioridade:** [Must Have / Should Have / Could Have]

#### User Stories (Rascunho Inicial)

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| [US-XXX] | Como **[Persona]**, quero [ação/funcionalidade] para [benefício/valor de negócio] | • [Critério de comportamento esperado de tela 1]<br>• [Critério de comportamento esperado de tela 2] |

#### Regras de Negócio

- **[RNXX-XX]:** [Descrever restrições operacionais específicas da tela, validações lógicas, regras de exibição e valores padrão].

---

## Matriz de Cobertura: Entregas do Project Charter × Funcionalidades

[Construir tabela mapeando cada ID de Entrega do Charter (D1, D2...) contra a lista de IDs de Funcionalidades relacionadas criadas].

[Desenhar a linha de tempo em formato ASCII consolidando as datas, Marcos M1 a M7, os Épicos correspondentes e listando verticalmente abaixo de cada marco todas as suas respectivas funcionalidades FXX-XX].


---

## Matriz de Priorização (MoSCoW)

[Construir tabela resumida contendo: Prioridade, Lista de Funcionalidades agregadas e Quantidade Total por linha].

---
`[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]`

