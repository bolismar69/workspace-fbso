# PROMPT-GENERATE-SPECS-ARTEFACT

## Contexto

Este prompt orquestra skills especializadas em especificação de solução e documentação técnica para gerar ou revisar o artefato `SPECS.md` na pasta de especificações de uma solução técnica.

O artefato gerado deve ser a **ponte entre os requisitos de negócio e a implementação técnica** — traduzindo user stories, regras de negócio e critérios de aceitação em especificações acionáveis para o time de desenvolvimento.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço ou frontend) | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SOLUTION_TYPE}` | Tipo da solução | `backend`, `frontend`, `batch`, `mobile` |
| `{SCOPE}` | Escopo da geração | `full` (criar do zero), `delta` (atualizar existente), `review` (apenas revisar) |
| `{BRANCH_NAME}` | Nome da branch onde deve ser realizado o desenvolvimento. Negar realizar desenvolvimento direto na branch `main` ou `master` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Antes de qualquer ação, verificar se TODOS os 7 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

### Passo 1 — Verificar e Preparar a Estrutura de Pastas

```
Verificar se existe: {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    │
    ├── NÃO existe → Criar a pasta (mkdir -p)
    │
    └── SIM, existe →
            │
            ├── Verificar se existem PRD.md E ARCHITECTURE.md na pasta
            │     │
            │     ├── AMBOS existem → Ler PRD.md + ARCHITECTURE.md como ponto de partida
            │     │
            │     ├── Apenas PRD.md → Ler PRD.md + documentos do projeto
            │     │
            │     ├── Apenas ARCHITECTURE.md → Ler ARCHITECTURE.md + documentos do projeto
            │     │
            │     └── NENHUM existe → Ler documentos do projeto em {PROJECT_PATH}:
            │           ├── 01-PROJECT-CHARTER-*.md (escopo, entregas, marcos, riscos)
            │           ├── 02-BUSINESS-REQUIREMENTS.md (BRs funcionais, NFRs)
            │           ├── 03-EPICS.md (épicos, jornadas, personas)
            │           ├── 04-FEATURES.md (features, user stories, regras de negócio)
            │           ├── 05-USER-STORYS-*.md (critérios de aceitação detalhados)
            │           ├── TECHNICAL-PLAN.md (stack, ERD, decisões técnicas)
            │           ├── ARCHITECTURE.md (C4, ADRs, pipeline de segurança)
            │           ├── API-CONTRACTS.md (contratos de API — para backend)
            │           └── DEFINITION_OF_DONE.md (critérios de DONE)
            │
            └── Verificar se SPECS.md já existe:
                  │
                  ├── SIM + SCOPE=full → Gerar nova versão (incrementar), preservar histórico
                  ├── SIM + SCOPE=delta → Atualizar apenas seções alteradas
                  ├── SIM + SCOPE=review → Revisar consistência contra docs do projeto
                  └── NÃO → Criar do zero
```

### Passo 2 — Invocar Skills Especializadas

Invocar as skills na ordem abaixo para embasar a geração do artefato:

| Ordem | Skill | Responsabilidade | O que extrair |
|---|---|---|---|
| 1ª | `create-specification` | Criação de especificação funcional e técnica | Estrutura do SPECS.md, mapeamento requisitos → features |
| 2ª | `spec-miner` | Mineração de especificações a partir de documentos existentes | Extrair especificações implícitas dos docs de negócio |
| 3ª | `domain-modeling` | Modelagem de domínio — entidades, regras, ubiquitous language | Glossário de domínio, regras de negócio formais |
| 4ª | `acceptance-criteria` | Critérios de aceitação por feature | Traduzir user stories em critérios verificáveis |
| 5ª | `documentation-writer` | Qualidade textual e consistência cross-documento | Revisão final: clareza, completude, rastreabilidade |

> **Nota:** Se `{SCOPE}=review`, invocar apenas `spec-miner` (para auditar contra docs do projeto) e `documentation-writer` (para verificar qualidade).

### Passo 3 — Gerar ou Atualizar o Artefato SPECS.md

Gerar o arquivo em:
```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md
```

#### Estrutura Obrigatória do Arquivo

```markdown
# SPECS.md — Especificação da Solução: {SOLUTION_NAME}

**Solução:** {SOLUTION_NAME}
**Tipo:** {SOLUTION_TYPE}
**Stack:** {STACK}
**Projeto:** {PROJECT_NAME}
**Versão do SPECS:** {X.0}
**Data:** {data atual}
**Status:** {status}
**Branch:** {BRANCH_NAME}

## 1. Visão Geral da Solução
- Parágrafo descrevendo o propósito da solução no contexto do projeto
- Relação com os épicos e entregas do Project Charter
- O que esta solução implementa (e o que NÃO implementa)

## 2. Requisitos Funcionais
- Tabela mapeando cada BR do projeto para features implementadas por esta solução
- Para cada feature, resumo do comportamento esperado
- Referência cruzada: BR-XXX → Feature FXX-XX → User Stories US-XXX a US-XXX

## 3. Regras de Negócio Implementadas
- Lista de todas as RNs (do FEATURES.md) que esta solução implementa
- Para cada RN: descrição formal, exemplos, casos de borda
- RNs que NÃO são implementadas por esta solução (com justificativa)

## 4. Especificação de APIs (para backend) ou Componentes (para frontend)
### 4.1 [Para Backend] Endpoints REST
- Tabela com: Método, Path, Descrição, RBAC requerido, Request/Response schema
- Regras de validação por endpoint
- Códigos de status HTTP esperados

### 4.1 [Para Frontend] Páginas e Componentes
- Tabela com: Rota, Componente principal, Props/State, API consumida
- Comportamento esperado por estado (loading, empty, error, success)
- Interações do usuário (cliques, formulários, navegação)

## 5. Requisitos Não-Funcionais
- Tabela mapeando cada NFR do BRD para implementação nesta solução
- Métricas objetivas: tempo de resposta, disponibilidade, segurança
- Como cada NFR será verificado (teste, monitoramento, auditoria)

## 6. Modelo de Dados (para backend)
- Lista de entidades/tabelas que esta solução cria ou consome
- Campos essenciais por entidade (nome, tipo, constraints)
- Relacionamentos (FK, índices, unique constraints)
- Entidades fora do escopo (referência futura)

## 7. Critérios de Aceitação por Feature
- Para cada feature: checklist de critérios que definem DONE
- Vinculado aos critérios do DEFINITION_OF_DONE.md
- Evidência esperada para cada critério

## 8. Dependências e Integrações
- Dependências de outras soluções (ex: frontend depende do backend)
- Dependências externas (Keycloak, PostgreSQL, RabbitMQ)
- Contratos de API consumidos ou fornecidos

## 9. Restrições e Premissas Técnicas
- Restrições herdadas do Project Charter (time, prazo, orçamento)
- Premissas técnicas (ex: "Keycloak disponível", "PostgreSQL com tenant isolation")
- Riscos técnicos específicos desta solução

## 10. Glossário da Solução
- Termos técnicos específicos desta solução
- Siglas e acrônimos
- Mapeamento termo de negócio ↔ termo técnico

## 11. Registro de Alterações
- Tabela de versões com data, alteração, autor

## Rodapé
- Indicação de geração por IA, skills utilizados
```

### Passo 4 — Validação Pós-Geração

Após gerar o arquivo, executar as seguintes verificações:

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo no path correto | `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md` existe |
| 2 | Rastreabilidade BR → Feature → US | Tabela em §2 mapeia cada BR para features e user stories |
| 3 | Regras de negócio documentadas | §3 lista todas as RNs aplicáveis com descrição formal |
| 4 | APIs ou Componentes especificados | §4 contém tabela completa de endpoints (backend) ou páginas (frontend) |
| 5 | NFRs mapeados com métricas | §5 vincula cada NFR a uma métrica objetiva e método de verificação |
| 6 | Modelo de dados (backend) | §6 lista entidades com campos, tipos, constraints e relacionamentos |
| 7 | Critérios de aceitação por feature | §7 tem checklist DONE vinculado ao DoD do projeto |
| 8 | Dependências documentadas | §8 lista dependências internas e externas com contratos |
| 9 | Glossário da solução | §10 define termos técnicos e mapeia negócio ↔ técnico |
| 10 | Rodapé de IA | Indicação de geração automatizada + skills utilizados |

---

## Modos de Operação por SCOPE

### SCOPE = full (Criação Completa)

- Executar Passos 0→1→2→3→4 integralmente
- Não existe SPECS.md prévio — criar do zero
- Ler todos os documentos disponíveis (PRD.md, ARCHITECTURE.md, docs do projeto)
- Invocar todas as 5 skills

### SCOPE = delta (Atualização Parcial)

- Executar Passos 0→1→2→3→4
- SPECS.md já existe — atualizar apenas seções que mudaram
- Identificar o delta: quais features/BRs/RNs foram adicionadas ou alteradas?
- Preservar conteúdo não afetado
- Incrementar versão e documentar no Registro de Alterações

### SCOPE = review (Apenas Revisão)

- Executar Passos 0→1 (leitura) → 2 (apenas spec-miner + documentation-writer) → 4
- NÃO modificar o arquivo — gerar relatório de revisão
- Verificar consistência contra docs do projeto (features, BRs, RNs)
- Identificar: seções desatualizadas, itens órfãos, inconsistências
- Emitir relatório com recomendações (não aplicar automaticamente)

---

## Skills Orquestradas

| Ordem | Skill | Propósito | Modos |
|---|---|---|---|
| 1ª | `create-specification` | Estrutura e conteúdo do SPECS.md | full, delta |
| 2ª | `spec-miner` | Extrair especificações de docs existentes | full, delta, review |
| 3ª | `domain-modeling` | Modelagem de entidades e regras de negócio | full, delta |
| 4ª | `acceptance-criteria` | Critérios de aceitação por feature | full, delta |
| 5ª | `documentation-writer` | Revisão de qualidade e consistência | full, delta, review |

---

## Exemplo de Invocação

```
Humano: "Gerar SPECS.md para o backend ms-fbso-platform-admin"

Agente: "Vou precisar de 7 parâmetros:
  - SOLUTION_PATH: caminho da pasta da solução
  - PROJECT_PATH: caminho da pasta do projeto
  - PROJECT_NAME: código do projeto
  - SOLUTION_NAME: nome da solução
  - SOLUTION_TYPE: backend, frontend, batch, ou mobile
  - BRANCH_NAME: nome da branch para desenvolvimento
  - SCOPE: full, delta, ou review"

Humano: "SOLUTION_PATH=/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin
         PROJECT_PATH=/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG
         PROJECT_NAME=PRJ-FIN-2026-0003-SAAS-FBSO-ORG
         SOLUTION_NAME=ms-fbso-platform-admin
         SOLUTION_TYPE=backend
         SCOPE=full
         BRANCH_NAME=feature/platform-admin"

Agente: [Executa Passo 1 → Passo 2 → Passo 3 → Passo 4]
```

---

## Observações

1. **PRD.md + ARCHITECTURE.md são o ponto de partida preferencial.** Se ambos existirem, eles contêm o escopo, entidades, ADRs e estrutura de pacotes. Os documentos do projeto são complementares para rastreabilidade.

2. **SPECS.md é a ponte negócio ↔ técnico.** Ele não repete o PRD.md (que é um guia de entrada) nem o ARCHITECTURE.md (que define o como). O SPECS.md define **o que** precisa ser construído com precisão suficiente para o time começar a codificar.

3. **Rastreabilidade é obrigatória.** Toda especificação deve referenciar o requisito de negócio (BR), a feature (FXX-XX) e a user story (US-XXX) que a originou. Sem rastreabilidade, o SPECS.md vira um documento isolado que ninguém confia.

4. **Modo `review` não modifica o arquivo.** Ele gera um relatório de auditoria que o time pode usar para decidir o que atualizar. Isso evita o risco de uma revisão automática introduzir erros.

5. **Os paths nos documentos devem ser relativos.** Usar paths relativos para que os links funcionem independente da máquina onde o código for clonado.

6. **O SPECS.md evolui com o projeto.** Na Fase 0 (Fundação), ele é um esboço. Conforme as features são implementadas (M2→M6), ele deve ser atualizado com o que foi efetivamente construído (modo `delta`).

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: fluxo de 5 passos, 5 skills orquestradas, 3 modos de operação (full/delta/review), 10 verificações pós-geração | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, create-specification.*
