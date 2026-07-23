# PROMPT-GENERATE-PRD-ARTEFACT

## Contexto

Este prompt orquestra skills especializadas em análise de negócio, requisitos e product management para gerar ou revisar o artefato `PRD.md` (Product Requirements Document) na pasta de especificações de uma solução técnica.

O artefato gerado é a **porta de entrada para a documentação de negócio** — um sumário executivo de alto nível que referencia os documentos-fonte do projeto (Project Charter, Business Requirements, Épicos, Features, User Stories) e serve como baseline de escopo para toda a cadeia downstream: ARCHITECTURE.md → SPECS.md → TASKS.md → TEST_PLAN.md.

**Princípio fundamental:** O PRD.md não duplica os documentos de negócio. Ele os **resume, referencia e contextualiza** para o time técnico, estabelecendo o escopo que será validado pelo GATE-PRD-SCOPE e desdobrado em arquitetura e especificações.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica (microsserviço ou frontend) | `/home/user/work/backend/go/fiber/microservices/ms-billing-engine-tax-rates` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-billing-engine-tax-rates` |
| `{SCOPE}` | Escopo da geração | `full` (criar do zero), `delta` (atualizar existente), `review` (apenas revisar) |
| `{BRANCH_STRATEGY}` | Estratégia de branching do projeto: `branch-por-sprint` (múltiplas sprints, uma branch `feature/sprint-NN-<slug>` por sprint) ou `branch-unica` (sprint único, branch `feature/<nome-projeto>`). A seção de branching do PRD.md gerado deve documentar a estratégia completa com tabela, workflow git e comandos. |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Antes de qualquer ação, verificar se TODOS os 6 parâmetros foram informados. Se algum estiver ausente, perguntar ao humano antes de prosseguir.

### Passo 1 — Verificar e Preparar a Estrutura de Pastas

```
Verificar se existe: {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/
    │
    ├── NÃO existe → Criar a pasta (mkdir -p)
    │
    └── SIM, existe →
            │
            ├── Verificar se existe PRD.md na pasta
            │     │
            │     ├── SIM + SCOPE=full → Gerar nova versão (incrementar), preservar histórico
            │     ├── SIM + SCOPE=delta → Atualizar apenas seções alteradas
            │     ├── SIM + SCOPE=review → Auditar consistência contra docs do projeto (relatório)
            │     └── NÃO → Criar do zero
            │
            └── Verificar documentos do projeto em {PROJECT_PATH}:
                  │
                  ├── TODOS os 6 documentos-base existem?
                  │     ├── 01-PROJECT-CHARTER-*.md (escopo, entregas D1-D7, marcos M1-M7, riscos, stakeholders)
                  │     ├── 02-BUSINESS-REQUIREMENTS.md (BRs funcionais, NFRs por bloco)
                  │     ├── 03-EPICS.md (épicos, jornadas, personas)
                  │     ├── 04-FEATURES.md (features, user stories, regras de negócio)
                  │     ├── MATRIZ-KPI.md (KPIs e critérios de sucesso)
                  │     └── DEFINITION_OF_DONE.md (critérios de DONE)
                  │
                  ├── PARCIAL → ALERTA: "Documentos ausentes: {lista}. A geração será limitada ao disponível."
                  │
                  └── NENHUM → ERRO: "Nenhum documento de projeto encontrado em {PROJECT_PATH}."
```

### Passo 2 — Invocar Skills Especializadas

Invocar as skills na ordem abaixo para embasar a geração do artefato:

| Ordem | Skill | Responsabilidade | O que extrair |
|---|---|---|---|
| 1ª | `prd` | Estrutura e conteúdo do Product Requirements Document | Template PRD, seções obrigatórias, nível de detalhe |
| 2ª | `requirements-elicitation` | Leitura e interpretação de requisitos de negócio | Extrair BRs, NFRs, objetivos estratégicos dos docs-fonte |
| 3ª | `stakeholder-analysis` | Identificação e mapeamento de stakeholders | Quem são os stakeholders, papéis, responsabilidades |
| 4ª | `gap-analysis` | Análise de lacunas entre docs do projeto e PRD existente | Delta entre versão atual e docs-fonte (modo delta/review) |
| 5ª | `documentation-writer` | Qualidade textual e consistência cross-documento | Revisão final: clareza, completude, rastreabilidade, links |

> **Nota:** Se `{SCOPE}=review`, invocar apenas `gap-analysis` (para auditar contra docs do projeto) e `documentation-writer` (para verificar qualidade e links).

### Passo 3 — Gerar ou Atualizar o Artefato PRD.md

Gerar o arquivo em:
```
{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md
```

#### Estrutura Obrigatória do Arquivo

```markdown
# Product Requirements Document (PRD)
## {TÍTULO DO PROGRAMA/PROJETO}

**Código:** {PROJECT_NAME}
**Versão do PRD:** {X.0}
**Data:** {data atual}
**Status:** {status — ex: "Em Definição", "Aprovado", "Fase 1 implementada"}
**Tipo:** Resumo de Alto Nível — porta de entrada para a documentação de negócio completa
**Situação impplementação:** {a ser preenchido pelos processos de desenvolvimento}

> ⚠️ **Aviso de Leitura:** Este documento é um sumário executivo. Todas as especificações detalhadas,
> regras de negócio, critérios de aceite e user stories completas residem nos documentos-fonte
> referenciados na Seção 4.

---

## 1. Visão Geral do Produto

- Contexto de negócio: qual problema o programa resolve, por que ele existe
- Abrangência: dimensões cobertas (ex: Comercial, Financeira, Estratégica)
- Tabela-resumo das dimensões com escopo de cada uma
- Referência ao documento-fonte: [01-PROJECT-CHARTER-*.md](path relativo)

---

## 2. Objetivos Estratégicos de Negócio

- Lista numerada de 4-8 objetivos que direcionam TODAS as iniciativas do programa
- Cada objetivo: 1 frase clara com verbo de ação
- Referência ao documento-fonte: [01-PROJECT-CHARTER-*.md](path relativo)

---

## 3. Requisitos de Negócio (Resumo)

- Tabela-resumo de TODOS os BRs (ex: BR-01 a BR-09), organizados por bloco temático
- Colunas: ID, Requisito (nome), Essência (1 frase)
- Blocos alinhados com as ondas/fases do programa
- Referência ao documento-fonte: [02-BUSINESS-REQUIREMENTS.md](path relativo)

---

## 4. Hierarquia da Documentação e Rastreabilidade

- Diagrama ASCII da cadeia: Project Charter → Requirements → Épicos → Features → User Stories
- Subseções por onda/fase com tabelas de documentos:
  - 4.1 Documentos da Onda 1 (se aplicável)
  - 4.2 Documentos da Onda 2 (se aplicável)
  - 4.N Governança e Métricas (MATRIZ-KPI.md, README.md)
  - 4.N+1 Base de Conhecimento Técnico (docs-suporte, se existirem)
- Referência ao documento-fonte: [README.md](path relativo) (índice completo)

---

## 5. Cronograma Macro

- Tabela com: Período, Fase, Foco
- Alinhado com os marcos M1-M7 do Project Charter
- Inclui fases de transição e shadow run (se aplicável)
- Referência ao documento-fonte: [01-PROJECT-CHARTER-*.md — Seção Cronograma](path relativo)

---

## 6. KPIs e Critérios de Sucesso

- KPIs agrupados por dimensão (Financeira, Compliance, Operacional)
- Tabela com: KPI (código + nome), Meta (valor quantificável)
- Critérios globais de sucesso (lista de bullet points)
- Referência aos documentos-fonte: [MATRIZ-KPI.md](path relativo) e [01-PROJECT-CHARTER-*.md](path relativo)

---

## 7. Principais Riscos de Negócio

- Tabela com: Risco, Impacto (Crítico/Alto/Médio), Mitigação
- 4-6 riscos principais do programa
- Referência ao documento-fonte: [01-PROJECT-CHARTER-*.md — Seção Riscos](path relativo)

---

## 8. Estatísticas do Projeto

- Tabela-resumo com contagens: BRs, Épicos, Features, User Stories, Regras de Negócio, KPIs
- Extraído por contagem real dos documentos-fonte (não estimado)

---

## 9. Stakeholders Principais

- Tabela com: Papel, Responsabilidade
- 5-8 stakeholders extraídos do Project Charter + STAKEHOLDER-MAP.md (se existir)
- Referência ao documento-fonte: [01-PROJECT-CHARTER-*.md](path relativo)

---

## 10. Referências Cruzadas

- Tabela mapeando CADA seção deste PRD ao documento-fonte correspondente
- Colunas: Seção do PRD, Documento-fonte (com link relativo)
- Isso garante rastreabilidade completa e auditável

---

## 11. Estratégia de Branching

> 🚫 **Regra de ouro:** Nenhum commit deste projeto pode ser feito diretamente em `main`. Todo desenvolvimento passa por branches `feature/*`.

**Se o projeto tem múltiplas sprints,** documentar a estratégia **uma branch por sprint**:

- Tabela de mapeamento Sprint → Branch (ex: `feature/sprint-03-portal-admin`)
- Ciclo de vida: CRIAR → DESENVOLVER → PR + REVIEW → MERGE NO MAIN → DELETAR
- Workflow de comandos git para início, durante e final da sprint
- Seção de hotfix para sprints já mergeadas

**Se o projeto tem sprint único,** documentar a branch única (ex: `feature/<nome-do-projeto>`).

> 📖 O design document canônico da estratégia (se existir) deve ser referenciado aqui.

---

## Rodapé
- Indicação de geração por IA, skills utilizados
- Link para o próximo passo: "Consulte os documentos-fonte listados acima para especificações completas"
```

### Passo 4 — Validação Pós-Geração

Após gerar o arquivo, executar as seguintes verificações:

| # | Verificação | Critério |
|---|---|---|
| 1 | Arquivo no path correto | `{SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md` existe |
| 2 | Header com metadados | Código, versão, data, status e tipo declarados no header |
| 3 | Visão geral presente | §1 contextualiza o programa e suas dimensões |
| 4 | Objetivos estratégicos | §2 lista 4-8 objetivos com verbo de ação, referenciando Project Charter |
| 5 | BRs resumidos | §3 tabela cobre TODOS os BRs do 02-BUSINESS-REQUIREMENTS.md, organizados por bloco |
| 6 | Hierarquia documentada | §4 tem diagrama ASCII + tabelas por onda com links para docs-fonte |
| 7 | Cronograma macro | §5 tabela alinhada com marcos M1-M7 do Project Charter |
| 8 | KPIs e critérios | §6 lista KPIs quantificáveis agrupados por dimensão + critérios globais |
| 9 | Riscos de negócio | §7 tabela com 4-6 riscos, impacto e mitigação |
| 10 | Estatísticas do projeto | §8 contagens reais extraídas dos documentos-fonte |
| 11 | Stakeholders | §9 lista 5-8 stakeholders com papéis e responsabilidades |
| 12 | Referências cruzadas | §10 mapeia cada seção ao documento-fonte com link relativo |
| 13 | Links relativos funcionais | Todos os links usam paths relativos (ex: `../../../../../../../business-inputs/...`) |
| 14 | Aviso de leitura | Bloco ⚠️ no header indicando que este é um sumário — detalhes nos docs-fonte |
| 15 | Rodapé de IA | Indicação de geração automatizada + skills utilizados |
| 16 | Estratégia de branching | Seção de Estratégia de Branching documenta a estratégia (tabela Sprint→Branch ou branch única) com workflow de comandos git |

---

## Modos de Operação por SCOPE

### SCOPE = full (Criação Completa)

- Executar Passos 0→1→2→3→4 integralmente
- Não existe PRD.md prévio — criar do zero
- Ler TODOS os documentos disponíveis em `{PROJECT_PATH}`
- Invocar todas as 5 skills
- Contar estatísticas reais (BRs, Épicos, Features, USs) a partir dos docs-fonte

### SCOPE = delta (Atualização Parcial)

- Executar Passos 0→1→2→3→4
- PRD.md já existe — atualizar apenas seções que mudaram
- Identificar o delta: quais BRs, features, épicos ou KPIs foram adicionados/alterados?
- Preservar conteúdo não afetado
- Incrementar versão (ex: 1.0 → 1.1) e documentar alterações no próprio PRD.md
- Atualizar estatísticas (§8) com contagens reais

### SCOPE = review (Apenas Revisão)

- Executar Passos 0→1 (leitura) → 2 (apenas gap-analysis + documentation-writer) → 4
- NÃO modificar o arquivo — gerar relatório de revisão
- Verificar consistência contra docs do projeto:
  - Todos os BRs referenciados? Nenhum BR faltando ou sobrando?
  - Links para docs-fonte estão funcionando?
  - Estatísticas (§8) batem com contagens reais dos docs-fonte?
  - Cronograma (§5) está atualizado com os marcos do Project Charter?
- Emitir relatório com recomendações (não aplicar automaticamente)

---

## Skills Orquestradas

| Ordem | Skill | Propósito | Modos |
|---|---|---|---|
| 1ª | `prd` | Estrutura e template do PRD.md | full, delta |
| 2ª | `requirements-elicitation` | Extrair BRs, NFRs e objetivos dos docs-fonte | full, delta |
| 3ª | `stakeholder-analysis` | Identificar stakeholders do Project Charter | full, delta |
| 4ª | `gap-analysis` | Auditar PRD existente vs docs-fonte | delta, review |
| 5ª | `documentation-writer` | Revisão de qualidade, links e consistência | full, delta, review |

---

## Exemplo de Invocação

```
Humano: "Gerar PRD.md para o ms-billing-engine-tax-rates"

Agente: "Vou precisar de 6 parâmetros:
  - SOLUTION_PATH: caminho da pasta do microsserviço
  - PROJECT_PATH: caminho da pasta do projeto de negócio
  - PROJECT_NAME: código do projeto
  - SOLUTION_NAME: nome do microsserviço
  - BRANCH_STRATEGY: estratégia de branching (ex: 'branch por sprint' ou branch única)
  - SCOPE: full, delta, ou review"

Humano: "SOLUTION_PATH=/home/user/work/backend/go/fiber/microservices/ms-billing-engine-tax-rates
         PROJECT_PATH=/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
         PROJECT_NAME=PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO
         SOLUTION_NAME=ms-billing-engine-tax-rates
         SCOPE=full
         BRANCH_STRATEGY=branch-por-sprint"

Agente: [Executa Passo 1 → Passo 2 → Passo 3 → Passo 4]
```

---

## Observações

1. **PRD.md é um sumário, não uma duplicata.** O PRD.md NÃO repete o conteúdo completo dos documentos de negócio. Cada seção resume os pontos principais e referencia o documento-fonte exato (com link relativo e indicação de seção) onde os detalhes podem ser encontrados.

2. **Rastreabilidade é o requisito mais importante.** O GATE-PRD-SCOPE valida 5 dimensões, e 3 delas (D1: Aderência ao Project Charter, D2: Consistência com BRs, D3: Consistência com Épicos/Features) dependem de rastreabilidade precisa. A Seção 10 (Referências Cruzadas) é a evidência dessa rastreabilidade.

3. **Estatísticas devem ser contadas, não estimadas.** A Seção 8 (Estatísticas do Projeto) deve conter contagens reais obtidas pela leitura dos documentos-fonte. Contar BRs no 02-BUSINESS-REQUIREMENTS.md, Features no 04-FEATURES.md, User Stories nos arquivos 05-USER-STORIES-*.md.

4. **Os paths nos links devem ser relativos.** Usar paths relativos (ex: `../../../../../../../business-inputs/business-projects/...`) para que os links funcionem independente da máquina onde o código for clonado.

5. **O PRD.md evolui com o projeto.** Na Fase 0 (Fundação), ele é um esboço baseado no Project Charter. Conforme as ondas são detalhadas (Onda 1, Onda 2), ele deve ser atualizado com novos BRs, features e KPIs (modo `delta`).

6. **O aviso de leitura é obrigatório.** O bloco ⚠️ no header é essencial para estabelecer expectativas corretas: este documento é um resumo, não a especificação completa. Isso evita que o time técnico tome decisões baseadas apenas no PRD.md sem consultar os docs-fonte.

7. **O PRD.md é o primeiro artefato da cadeia.** Ele alimenta ARCHITECTURE.md (que define o COMO), que alimenta SPECS.md (que define O QUE construir), que alimenta TASKS.md e TEST_PLAN.md. Um PRD.md incorreto contamina toda a cadeia. Por isso o GATE-PRD-SCOPE é o primeiro gate do fluxo.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.1 | 16/07/2026 | Padronização de branching: `{BRANCH_NAME}` substituído por `{BRANCH_STRATEGY}`. Header do template sem linha `Branch:`. Adicionada §11 (Estratégia de Branching) com tabela Sprint→Branch, ciclo de vida e workflow git. Validação #16 adicionada. | Time Técnico |
| 1.0 | 13/07/2026 | Criação inicial: fluxo de 5 passos, 5 skills orquestradas, 3 modos de operação (full/delta/review), 15 verificações pós-geração, estrutura de 10 seções baseada nos PRDs existentes | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, prd, requirements-elicitation.*
