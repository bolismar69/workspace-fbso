# PROMPT-GATE-SPECS-TECHNICAL

## Contexto

Este prompt implementa o **Gate de Viabilidade e Completude Técnica** para o artefato `SPECS.md`, conforme definido no fluxo Spec-Driven Development (etapa 3.1 — GT1).

O agente validador atua como um **revisor técnico** — verificando se as especificações são completas, consistentes, tecnicamente viáveis e prontas para serem decompostas em tarefas. Diferente do gate de escopo (etapa 2.1), o foco aqui é **qualidade técnica e integridade da especificação**.

**Princípio fundamental:** O SPECS.md é a ponte entre requisitos e implementação. Uma especificação incompleta ou inconsistente gera tarefas mal definidas, retrabalho e bugs.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SOLUTION_TYPE}` | Tipo da solução | `backend`, `frontend`, `batch`, `mobile` |
| `{BRANCH_NAME}` | Nome da branch onde deve ser realizado o desenvolvimento. Negar realizar desenvolvimento direto na branch `main` ou `master` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 6 parâmetros foram informados.

### Passo 1 — Carregar Documentos Base

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md (artefato a ser validado)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md (baseline de escopo — APROVADO)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md (baseline técnica — APROVADO)
    └── Documentos de referência do projeto:
          ├── 02-BUSINESS-REQUIREMENTS.md (BRs e NFRs)
          ├── 04-FEATURES.md (features, user stories, regras de negócio)
          └── DEFINITION_OF_DONE.md (critérios de DONE)

Se SPECS.md não existir → ERRO: "SPECS.md não encontrado. Execute o agente gerador de SPECS."
Se PRD.md ou ARCHITECTURE.md não existirem → ERRO: "PRD.md e ARCHITECTURE.md aprovados são pré-requisitos para validação do SPECS.md."
```

### Passo 2 — Executar Dimensões de Validação Técnica

O gate avalia o SPECS.md em **6 dimensões**. Para cada dimensão, atribuir um veredito: `APROVADO`, `RESSALVA` ou `REPROVADO`.

#### Dimensão 1: Rastreabilidade e Cobertura

| # | Verificação | Critério |
|---|---|---|
| 1.1 | Matriz BR → Feature → US | Cada requisito de negócio (BR) está mapeado para uma feature e user story no SPECS.md? |
| 1.2 | Cobertura de features | Todas as features do PRD.md atribuídas a esta solução estão especificadas? |
| 1.3 | Features órfãs | Existem features no SPECS.md sem correspondência no PRD.md? |
| 1.4 | Cobertura de user stories | Todas as US do FEATURES.md para esta solução têm especificação correspondente? |

#### Dimensão 2: Completude da Especificação Funcional

| # | Verificação | Critério |
|---|---|---|
| 2.1 | Regras de negócio formais | Cada RN está descrita com: descrição formal, exemplos e casos de borda? |
| 2.2 | APIs documentadas (backend) | Cada endpoint tem: método, path, RBAC, request/response schema e códigos de status? |
| 2.3 | Componentes documentados (frontend) | Cada página/componente tem: rota, props/state, API consumida, estados (loading/empty/error/success)? |
| 2.4 | Comportamento de borda | Casos de erro, edge cases e cenários de falha estão especificados? |

#### Dimensão 3: Consistência com Arquitetura

| # | Verificação | Critério |
|---|---|---|
| 3.1 | Entidades vs modelo de dados | As entidades no §6 do SPECS.md correspondem às entidades do ARCHITECTURE.md? |
| 3.2 | APIs vs estrutura de pacotes | Os endpoints do SPECS.md são suportados pela estrutura de pacotes do ARCHITECTURE.md? |
| 3.3 | NFRs vs ADRs | Os NFRs do SPECS.md são endereçados pelos ADRs do ARCHITECTURE.md? |
| 3.4 | Dependências vs integrações | As dependências listadas no §8 batem com o INTEGRATION-MAP.md? |

#### Dimensão 4: Requisitos Não-Funcionais (NFRs)

| # | Verificação | Critério |
|---|---|---|
| 4.1 | Métricas objetivas | Cada NFR tem uma métrica quantificável? (ex: "p95 < 200ms", não "rápido") |
| 4.2 | Método de verificação | Cada NFR especifica COMO será verificado? (teste, monitoramento, auditoria) |
| 4.3 | Cobertura de NFRs | Todos os NFRs do BRD aplicáveis a esta solução estão no §5? |
| 4.4 | NFRs de segurança | RBAC, Multi-Tenant, OWASP, LGPD estão cobertos com cenários específicos? |

#### Dimensão 5: Critérios de Aceitação

| # | Verificação | Critério |
|---|---|---|
| 5.1 | Checklist por feature | Cada feature tem um checklist de critérios de aceitação no §7? |
| 5.2 | Vinculação ao DoD | Os critérios estão vinculados ao DEFINITION_OF_DONE.md? |
| 5.3 | Evidência esperada | Cada critério especifica qual evidência comprova seu cumprimento? |
| 5.4 | Testabilidade | Os critérios são objetivos e testáveis? (evitar "funciona corretamente") |

#### Dimensão 6: Qualidade e Consistência Documental

| # | Verificação | Critério |
|---|---|---|
| 6.1 | Glossário da solução | O §10 define termos técnicos e mapeia negócio ↔ técnico? |
| 6.2 | Restrições e premissas | O §9 lista restrições (tempo, prazo, orçamento) e premissas técnicas? |
| 6.3 | Consistência terminológica | Os mesmos conceitos têm os mesmos nomes em todo o documento? |
| 6.4 | Referências cruzadas | Links para PRD.md, ARCHITECTURE.md e docs do projeto estão corretos? |

### Passo 3 — Calcular Veredito Final

```
Para cada dimensão:
    - 100% verificações OK → APROVADO
    - >= 75% verificações OK → RESSALVA
    - < 75% verificações OK → REPROVADO

Veredito final do gate:
    ├── APROVADO: Todas as 6 dimensões APROVADAS
    ├── RESSALVA: Pelo menos 1 dimensão com RESSALVA, nenhuma REPROVADA
    └── REPROVADO: Pelo menos 1 dimensão REPROVADA
```

### Passo 4 — Gerar Relatório de Falha (se REPROVADO ou RESSALVA)

```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL_SPECS_FAIL_REPORT.md
```

#### Estrutura do Relatório de Falha

```markdown
# TECHNICAL_SPECS_FAIL_REPORT.md — Relatório de Validação Técnica: SPECS.md

[Header com metadados: solução, tipo, projeto, data, versão do SPECS.md validado]

## 1. Sumário Executivo
- Veredito: REPROVADO | RESSALVA
- Dimensões aprovadas: X/6
- Dimensões com ressalva: Y/6
- Dimensões reprovadas: Z/6
- Total de não-conformidades: N
- Bloqueia TASKS.md? SIM (se REPROVADO) | NÃO (se RESSALVA, mas requer acompanhamento)

## 2. Veredito por Dimensão
| Dimensão | Veredito | Verificações OK | Total Verificações | % |
|---|---|---|---|---|
| 1. Rastreabilidade e Cobertura | ... | ... | ... | ... |
| 2. Completude Funcional | ... | ... | ... | ... |
| 3. Consistência com Arquitetura | ... | ... | ... | ... |
| 4. Requisitos Não-Funcionais | ... | ... | ... | ... |
| 5. Critérios de Aceitação | ... | ... | ... | ... |
| 6. Qualidade Documental | ... | ... | ... | ... |

## 3. Não-Conformidades Detalhadas

### 3.1 [Dimensão 1] Rastreabilidade e Cobertura
| ID | Verificação | Status | Evidência | Ação Corretiva |
|---|---|---|---|---|
| NC-001 | 1.1 Matriz BR→Feature→US | REPROVADO | BR-007 sem feature mapeada | Adicionar mapeamento de BR-007 para feature correspondente |

### 3.2 [Dimensão 2] Completude Funcional
...

### 3.3 [Dimensão 3] Consistência com Arquitetura
...

### 3.4 [Dimensão 4] Requisitos Não-Funcionais
...

### 3.5 [Dimensão 5] Critérios de Aceitação
...

### 3.6 [Dimensão 6] Qualidade Documental
...

## 4. Bloqueios Técnicos Identificados
- Itens que IMPEDEM a geração de TASKS.md (não-conformidades críticas)
- Itens que permitem prosseguir com RESSALVA (não-conformidades não-bloqueantes)

## 5. Análise de Viabilidade Técnica
- Features com especificação insuficiente para implementação
- APIs sem contrato completo (faltando request/response schema)
- Regras de negócio ambíguas ou contraditórias
- Dependências externas não resolvidas

## 6. Recomendações para Correção
- Lista priorizada de ações corretivas (bloqueantes primeiro)
- Referências aos documentos que devem ser consultados
- Seções do SPECS.md que precisam ser reescritas ou complementadas

## 7. Próximos Passos
1. Encaminhar este relatório ao Agente de Specs para correção
2. Após correção, reexecutar este gate
3. Não prosseguir para TASKS.md até SPECS.md ser APROVADO (ou RESSALVA com acompanhamento)

## Rodapé
- Indicação de geração por IA, skills utilizados
```

### Passo 5 — Validar o Relatório Gerado

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo no path correto | `TECHNICAL_SPECS_FAIL_REPORT.md` existe |
| 2 | Sumário executivo | §1 contém veredito e indica se bloqueia TASKS.md |
| 3 | Veredito por dimensão | §2 lista as 6 dimensões |
| 4 | Não-conformidades detalhadas | §3 lista cada NC com ID, verificação, evidência e ação corretiva |
| 5 | Bloqueios identificados | §4 classifica NCs em bloqueantes vs não-bloqueantes |
| 6 | Análise de viabilidade | §5 avalia se as especificações são suficientes para implementar |
| 7 | Recomendações acionáveis | §6 contém ações corretivas priorizadas |
| 8 | Próximos passos claros | §7 indica o fluxo de correção |
| 9 | Rodapé de IA | Indicação de geração automatizada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `spec-miner` | Auditar especificações contra docs de negócio |
| 2ª | `acceptance-criteria` | Validar critérios de aceitação — são testáveis e completos? |
| 3ª | `domain-modeling` | Verificar consistência do modelo de domínio e regras de negócio |
| 4ª | `gap-analysis` | Identificar lacunas entre SPECS.md, PRD.md e ARCHITECTURE.md |
| 5ª | `documentation-writer` | Revisão final de qualidade e consistência do relatório |

---

## Observações

1. **Este gate é o ponto de não-retorno técnico.** Após SPECS.md aprovado, o time começa a decompor em tarefas (TASKS.md) e cenários de teste (TEST_PLAN.md). Erros aqui custam caro downstream.

2. **RESSALVA permite prosseguir, mas exige acompanhamento.** Não-conformidades não-bloqueantes são documentadas como RESSALVA e devem ser resolvidas antes do gate final de arquitetura (etapa 6).

3. **Viabilidade técnica é binária para itens bloqueantes.** Se uma feature não tem especificação suficiente para ser implementada, o gate REPROVA, independentemente das outras dimensões.

4. **Consistência cruzada é obrigatória.** O SPECS.md deve ser consistente com PRD.md (escopo), ARCHITECTURE.md (como) e FEATURES.md (user stories). Inconsistências entre documentos são tratadas como não-conformidades.

5. **O relatório de falha classifica criticidade.** Nem toda não-conformidade bloqueia o fluxo. O §4 do relatório separa claramente o que é bloqueante do que é recomendação.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: gate técnico para SPECS.md, 6 dimensões, 24 verificações, classificação bloqueante vs não-bloqueante | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, spec-miner, gap-analysis.*
