# PROMPT-GATE-PRD-SCOPE

## Contexto

Este prompt implementa o **Gate de Alinhamento de Escopo** para o artefato `PRD.md` (Product Requirements Document), conforme definido no fluxo Spec-Driven Development (etapa 2.1 — GE1).

O agente validador atua como um **guardião de escopo** — verificando se o PRD.md mantém aderência estrita ao Project Charter e aos documentos de negócio do projeto, prevenindo scope creep antes que os requisitos incorretos contaminem a arquitetura, especificações e tarefas.

**Princípio fundamental:** O PRD.md é a tradução fiel do escopo de negócio para o time técnico. Qualquer desvio de escopo neste ponto propaga-se por toda a cadeia downstream (ARCHITECTURE → SPECS → TASKS → TEST_PLAN → Código).

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{BRANCH_NAME}` | Nome da branch onde deve ser realizado o desenvolvimento. Negar realizar desenvolvimento direto na branch `main` ou `master` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 5 parâmetros foram informados. Se algum estiver ausente, perguntar antes de prosseguir.

### Passo 1 — Carregar Documentos Base

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md (artefato a ser validado)
    └── Documentos de referência do projeto em {PROJECT_PATH}:
          ├── 01-PROJECT-CHARTER-*.md (escopo aprovado, entregas D1-D7, marcos M1-M7)
          ├── 02-BUSINESS-REQUIREMENTS.md (BRs funcionais, NFRs)
          ├── 03-EPICS.md (épicos, jornadas, personas)
          ├── 04-FEATURES.md (features, user stories, regras de negócio)
          └── DEFINITION_OF_DONE.md (critérios de DONE)

Se PRD.md não existir → ERRO: "PRD.md não encontrado. Execute primeiro o agente gerador de PRD."
Se documentos do projeto não existirem → ALERTA: "Documentos de referência ausentes. A validação será limitada."
```

### Passo 2 — Executar Dimensões de Validação de Escopo

O gate avalia o PRD.md em **5 dimensões** independentes. Para cada dimensão, atribuir um veredito: `APROVADO`, `RESSALVA` ou `REPROVADO`.

#### Dimensão 1: Aderência ao Project Charter (Scope Baseline)

| # | Verificação | Critério |
|---|---|---|
| 1.1 | Entregas D1-D7 | Cada entrega listada no PRD.md corresponde a uma entrega do Project Charter? Há entregas no PRD.md não previstas no Charter? |
| 1.2 | Marcos M1-M7 | Os marcos do PRD.md respeitam o sequenciamento e datas do Project Charter? |
| 1.3 | Escopo declarado | O PRD.md declara explicitamente o que está FORA de escopo? Há itens ambíguos que podem ser interpretados como scope creep? |
| 1.4 | Objetivos do projeto | Os objetivos listados no PRD.md são consistentes com os do Project Charter? Há objetivos adicionais não aprovados? |

#### Dimensão 2: Consistência com Business Requirements (BRs)

| # | Verificação | Critério |
|---|---|---|
| 2.1 | Cobertura de BRs | Todos os BRs do 02-BUSINESS-REQUIREMENTS.md estão referenciados no PRD.md? |
| 2.2 | BRs órfãos | Existem BRs no PRD.md que não constam nos documentos de negócio? (scope creep) |
| 2.3 | NFRs mapeados | Os requisitos não-funcionais do BRD estão refletidos no PRD.md? |
| 2.4 | Priorização | A priorização (MoSCoW) do PRD.md é consistente com o BRD? |

#### Dimensão 3: Consistência com Épicos e Features

| # | Verificação | Critério |
|---|---|---|
| 3.1 | Épicos cobertos | Cada épico do 03-EPICS.md está representado no PRD.md? |
| 3.2 | Features mapeadas | As features do 04-FEATURES.md estão corretamente atribuídas às entregas no PRD.md? |
| 3.3 | User Stories referenciadas | O PRD.md referencia as user stories corretas para cada feature? |
| 3.4 | Personas e jornadas | As personas e jornadas do PRD.md são consistentes com os épicos? |

#### Dimensão 4: Rastreabilidade e Completude

| # | Verificação | Critério |
|---|---|---|
| 4.1 | Matriz de rastreabilidade | Existe matriz BR → Feature → US → Entrega? Está completa e sem lacunas? |
| 4.2 | Regras de negócio | Todas as RNs do FEATURES.md estão documentadas no PRD.md? |
| 4.3 | Critérios de aceite | Cada user story tem critérios de aceite documentados e vinculados ao DoD? |
| 4.4 | Dependências | Dependências entre features e entre soluções estão documentadas? |

#### Dimensão 5: Não-Escopo e Anti-Scope Creep

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Seção de fora-de-escopo | O PRD.md tem uma seção explícita de "Fora de Escopo"? |
| 5.2 | Gold plating | Há funcionalidades "nice-to-have" não solicitadas nos docs de negócio? |
| 5.3 | Premissas não validadas | O PRD.md declara premissas? Elas são razoáveis e validadas? |
| 5.4 | Escopo futuro | Menções a "fases futuras" são claramente separadas do escopo atual? |

### Passo 3 — Calcular Veredito Final

```
Para cada dimensão:
    - 100% verificações OK → APROVADO
    - >= 75% verificações OK → RESSALVA (aprova com observações)
    - < 75% verificações OK → REPROVADO

Veredito final do gate:
    ├── APROVADO: Todas as 5 dimensões APROVADAS
    ├── RESSALVA: Pelo menos 1 dimensão com RESSALVA, nenhuma REPROVADA
    └── REPROVADO: Pelo menos 1 dimensão REPROVADA
```

### Passo 4 — Gerar Relatório de Falha (se REPROVADO ou RESSALVA)

Se o veredito for **REPROVADO** ou **RESSALVA**, gerar o relatório:

```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD_SCOPE_FAIL_REPORT.md
```

#### Estrutura do Relatório de Falha

```markdown
# PRD_SCOPE_FAIL_REPORT.md — Relatório de Validação de Escopo: PRD.md

[Header com metadados: solução, projeto, data da validação, versão do PRD validado]

## 1. Sumário Executivo
- Veredito: REPROVADO | RESSALVA
- Dimensões aprovadas: X/5
- Dimensões com ressalva: Y/5
- Dimensões reprovadas: Z/5
- Total de não-conformidades: N

## 2. Veredito por Dimensão
| Dimensão | Veredito | Verificações OK | Total Verificações | % |
|---|---|---|---|---|
| 1. Aderência ao Project Charter | ... | ... | ... | ... |
| 2. Consistência com BRs | ... | ... | ... | ... |
| 3. Consistência com Épicos/Features | ... | ... | ... | ... |
| 4. Rastreabilidade e Completude | ... | ... | ... | ... |
| 5. Não-Escopo e Anti-Scope Creep | ... | ... | ... | ... |

## 3. Não-Conformidades Detalhadas

### 3.1 [Dimensão 1] Aderência ao Project Charter
| ID | Verificação | Status | Evidência | Ação Corretiva |
|---|---|---|---|---|
| NC-001 | 1.1 Entregas D1-D7 | REPROVADO | Entrega D4 não mapeada no PRD | Adicionar seção para entrega D4 |

### 3.2 [Dimensão 2] Consistência com BRs
...

### 3.3 [Dimensão 3] Consistência com Épicos/Features
...

### 3.4 [Dimensão 4] Rastreabilidade e Completude
...

### 3.5 [Dimensão 5] Não-Escopo e Anti-Scope Creep
...

## 4. Itens Suspeitos de Scope Creep
- Lista de requisitos/features no PRD.md sem correspondência nos docs de negócio
- Funcionalidades "gold plating" identificadas
- Premissas não validadas que expandem escopo

## 5. Recomendações para Correção
- Lista priorizada de ações corretivas (ordenadas por gravidade)
- Referências aos documentos de negócio que devem ser consultados
- Seções do PRD.md que precisam ser reescritas

## 6. Próximos Passos
1. Encaminhar este relatório ao Agente PRD para correção
2. Após correção, reexecutar este gate
3. Não prosseguir para ARCHITECTURE.md até PRD.md ser APROVADO

## Rodapé
- Indicação de geração por IA, skills utilizados
```

### Passo 5 — Validar o Relatório Gerado

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo no path correto | `PRD_SCOPE_FAIL_REPORT.md` existe na pasta correta |
| 2 | Sumário executivo presente | §1 contém veredito e contagem de dimensões |
| 3 | Tabela de veredito por dimensão | §2 lista as 5 dimensões com % de aprovação |
| 4 | Não-conformidades detalhadas | §3 lista cada NC com ID, verificação, status, evidência e ação corretiva |
| 5 | Scope creep identificado | §4 lista itens sem correspondência nos docs de negócio |
| 6 | Recomendações acionáveis | §5 contém ações corretivas priorizadas |
| 7 | Próximos passos claros | §6 indica o fluxo de correção |
| 8 | Rodapé de IA | Indicação de geração automatizada + skills utilizados |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `gap-analysis` | Análise de lacunas entre PRD.md e docs de negócio |
| 2ª | `requirements-validation` | Validação de requisitos contra baseline de escopo |
| 3ª | `stakeholder-alignment-checker` | Verificar alinhamento do PRD com stakeholders |
| 4ª | `documentation-writer` | Qualidade e clareza do relatório de falha |

---

## Observações

1. **Gate binário com escape de ressalva.** O gate não é puramente binário. RESSALVAS permitem que o fluxo prossiga com anotações, mas requere que as não-conformidades sejam endereçadas antes do gate final de arquitetura (etapa 6).

2. **Scope creep é contagioso.** Um requisito fora de escopo no PRD.md gera features desnecessárias no SPECS.md, tarefas no TASKS.md e cenários de teste no TEST_PLAN.md. Este gate é a primeira e mais importante linha de defesa.

3. **O relatório de falha é o contrato de correção.** O PRD_SCOPE_FAIL_REPORT.md não é apenas um apontamento de erros — ele contém as ações corretivas exatas e referências aos docs de negócio que o Agente PRD deve usar para corrigir.

4. **Revalidação obrigatória.** Após correção do PRD.md, este gate DEVE ser reexecutado para garantir que as não-conformidades foram resolvidas e que novas não foram introduzidas.

5. **Independência do validador.** O agente GATE não corrige o PRD.md — ele apenas valida e reporta. A correção é responsabilidade do Agente PRD (via PROMPT-FIX-PRD-FROM-GATE.md).

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: gate de escopo para PRD.md, 5 dimensões de validação, 20 verificações, relatório PRD_SCOPE_FAIL_REPORT.md | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, gap-analysis, requirements-validation.*
