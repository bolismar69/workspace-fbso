# PROMPT: ROADMAP DE EXECUÇÃO MACRO E GUIA DE ORQUESTRAÇÃO DE DOCUMENTOS
## Versão: 5.0 — Integrada com Bootstrap Inteligente, Validação Soberana Humana (Human-in-the-Loop) e Git Workflow Automatizado

Atue como um Especialista em Gestão de Processos (BPM) e Arquiteto de Soluções Ágeis, especializado em Auditoria de Escopo de Projetos e Engenharia de Prompts.

Preciso que você crie um roadmap de execução detalhado e um guia de orquestração para o seguinte processo: Criação, revisão, evolução e validação dos documentos base de um projeto.

Objetivo Principal: Garantir que todos os documentos estejam criados, revisados e 100% alinhados conceitualmente entre si (rastreabilidade vertical de ponta a ponta), mitigando desvios de escopo (scope creep) para garantir o sucesso do projeto.

Regra Crítica de Execução (Gating Rule): O processo é estritamente sequencial. Nenhuma fase subsequente pode ser iniciada sem a aprovação formal, soberana e explícita do usuário humano na fase anterior.

---

## VARIÁVEIS DE ENTRADA E BOOTSTRAP DO PROJETO (FASE 0)

Antes de iniciar qualquer fase de geração de documentos, o prompt deve obrigatoriamente executar o ritual de bootstrap descrito abaixo. Esta fase garante que o escopo do projeto está corretamente parametrizado, que a estrutura de diretórios existe e que o estado atual dos artefatos é conhecido.

### Tabela de Inputs Obrigatórios

| Variável | Descrição | Exemplo |
|---|---|---|
| `PROJECT_PATH` | Caminho base onde os projetos de negócio residem | `/home/bolismar/work/workspace-fbso/business-inputs/business-projects` |
| `PROJECT_ID` | Identificador único do projeto (ID corporativo) | `PRJ-FIN-2026-0003` |
| `PROJECT_NAME` | Nome curto do produto/projeto | `SAAS-FBSO-ORG` |
| `PROJECT_DOCUMENTS_INPUTS` | Lista de caminhos para documentos brutos de entrada (atas, PDFs, especificações) usados como insumo para geração | `[]` (ex: `[/tmp/ata-reuniao.md, /tmp/especificacao-v1.pdf]`) |
| `PROJECT_PROMPT_INPUTS` | Lista de caminhos para prompts auxiliares ou contextos adicionais a serem carregados | `[]` (ex: `[/tmp/contexto-tecnico.md]`) |
| `PROMPT_BRANCH` | Nome da branch Git onde as alterações serão salvas. **Não pode** ser `main`, `master` ou `develop`. | `feature/PRJ-FIN-2026-0003-docs` |

### Validação do PROMPT_BRANCH

**Regra de Bloqueio (Gating Rule):** O processo NÃO pode ser iniciado se `PROMPT_BRANCH` for um dos valores proibidos. Se o usuário informar `main`, `master` ou `develop`, exiba a mensagem:

> ⛔ **Branch Inválida:** `{PROMPT_BRANCH}` é uma branch protegida. O processo de documentação NÃO pode ser executado diretamente em branches protegidas.
>
> Por favor, informe um nome de branch de trabalho (ex: `feature/PRJ-FIN-2026-0003-docs`, `docs/atualizacao-prompts`).

Repita a solicitação até receber um nome de branch válido.

### Variáveis Derivadas (calculadas automaticamente)

A partir dos inputs acima, o prompt deve computar:

```
PROJECT_ID_NAME            = PROJECT_ID + "-" + PROJECT_NAME
PROJECT_COMPLETE_PATH_NAME = PROJECT_PATH + "/" + PROJECT_ID_NAME
```

**Exemplo concreto:**
```
PROJECT_ID_NAME            = "PRJ-FIN-2026-0003-SAAS-FBSO-ORG"
PROJECT_COMPLETE_PATH_NAME = "/home/bolismar/work/workspace-fbso/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG"
```

### Workflow de Bootstrap (Execução Obrigatória)

Execute os passos abaixo em ordem estrita. Não prossiga para a Fase 1 sem completar todos eles.

---

#### Passo 0.1 — Solicitar Inputs ao Usuário

Se alguma das 6 variáveis da tabela de inputs não tiver sido fornecida no contexto, pergunte ao usuário de forma clara e objetiva:

> "Para iniciar o Roadmap de Documentos, preciso das seguintes informações:
> 1. **PROJECT_PATH** — Caminho base dos projetos (ex: `/home/bolismar/work/workspace-fbso/business-inputs/business-projects`)
> 2. **PROJECT_ID** — ID do projeto (ex: `PRJ-FIN-2026-0003`)
> 3. **PROJECT_NAME** — Nome do produto (ex: `SAAS-FBSO-ORG`)
> 4. **PROJECT_DOCUMENTS_INPUTS** — Documentos de entrada (deixe vazio `[]` se não houver)
> 5. **PROJECT_PROMPT_INPUTS** — Prompts auxiliares (deixe vazio `[]` se não houver)
> 6. **PROMPT_BRANCH** — Nome da branch Git para salvar as alterações (ex: `feature/PRJ-FIN-2026-0003-docs`). **Não pode** ser `main`, `master` ou `develop`."

**Após coletar os inputs, validar PROMPT_BRANCH:** Se o valor for `main`, `master` ou `develop`, aplicar a Regra de Bloqueio da seção "Validação do PROMPT_BRANCH" e solicitar novamente.

---

#### Passo 0.2 — Exibir Caminho Derivado e Solicitar Confirmação

Após receber os inputs, compute `PROJECT_COMPLETE_PATH_NAME` e `PROJECT_ID_NAME` e exiba ao usuário:

> **📁 Caminho do Projeto:** `{PROJECT_COMPLETE_PATH_NAME}`
> **🏷️ Identificador:** `{PROJECT_ID_NAME}`
> **🌿 Branch Git:** `{PROMPT_BRANCH}`
>
> Confirma que estas informações estão corretas?
> - **SIM** → Prosseguir para criação/verificação da estrutura de diretórios
> - **NÃO** → Solicitar correção dos inputs e repetir o Passo 0.2

**Regra:** Não avance sem a confirmação explícita do humano.

---

#### Passo 0.3 — Criar Estrutura de Diretórios

Uma vez confirmado, execute:

```bash
mkdir -p {PROJECT_COMPLETE_PATH_NAME}/user-stories/
mkdir -p {PROJECT_COMPLETE_PATH_NAME}/epics/
mkdir -p {PROJECT_COMPLETE_PATH_NAME}/features/
```

Este comando:
- Cria a pasta do projeto se não existir (equivalente a `mkdir -p` no topo)
- Cria a subpasta `user-stories/` para os artefatos modulares da Fase 5
- Cria a subpasta `epics/` para os artefatos modulares da Fase 3 (arquivos individuais por épico)
- Cria a subpasta `features/` para os artefatos modulares da Fase 4 (arquivos individuais por feature)
- É idempotente — não tem efeito colateral se as pastas já existirem

---

#### Passo 0.4 — Verificar Status dos Arquivos de Projeto

Verifique na ordem a existência de cada artefato do roadmap e reporte o status:

| Arquivo | Caminho Esperado | Status |
|---|---|---|
| Project Charter | `{PROJECT_COMPLETE_PATH_NAME}/01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` | ✅ Existe / ❌ Não existe |
| BRD | `{PROJECT_COMPLETE_PATH_NAME}/02-BRD-{PROJECT_ID_NAME}.md` | ✅ Existe / ❌ Não existe |
| Epics (índice) | `{PROJECT_COMPLETE_PATH_NAME}/03-EPICS-{PROJECT_ID_NAME}.md` | ✅ Existe / ❌ Não existe |
| Epics (pasta) | `{PROJECT_COMPLETE_PATH_NAME}/epics/` | ✅ Existe com N arquivos / ❌ Vazia / ❌ Não existe |
| Features | `{PROJECT_COMPLETE_PATH_NAME}/04-FEATURES-{PROJECT_ID_NAME}.md` | ✅ Existe / ❌ Não existe |
| Features (pasta) | `{PROJECT_COMPLETE_PATH_NAME}/features/` | ✅ Existe com N arquivos / ❌ Vazia / ❌ Não existe |
| User Stories (pasta) | `{PROJECT_COMPLETE_PATH_NAME}/user-stories/` | ✅ Existe com N arquivos / ❌ Vazia / ❌ Não existe |
| Matriz RTM | `{PROJECT_COMPLETE_PATH_NAME}/05-USER-STORIES-{PROJECT_ID_NAME}.md` | ✅ Existe / ❌ Não existe |

**Lógica de decisão com base no status:**
- Se **todos** os arquivos estão marcados como ❌ Não existe → Projeto novo. Iniciar da Fase 1 (Project Charter).
- Se **alguns** arquivos existem → Projeto em andamento. Apresentar o status e perguntar ao humano de qual fase deseja continuar (o processo sequencial será retomado a partir dali, respeitando as dependências entre fases).
- Se **todos** os arquivos existem → Projeto completo. Perguntar se deseja revisar alguma fase específica ou iniciar um novo ciclo de evolução.

---

#### Passo 0.5 — Apresentar Resumo e Iniciar

Exiba um resumo final antes de iniciar a primeira fase pendente:

> **📊 Resumo do Projeto:** `{PROJECT_ID_NAME}`
> **📁 Localização:** `{PROJECT_COMPLETE_PATH_NAME}`
> **📝 Próxima Fase:** [Fase N — Nome da Fase]
> **📄 Artefatos Existentes:** [X de 6]
>
> Iniciando a [Fase N]...

--------------------------------------------------------------------------------
MECANISMO DE ORQUESTRAÇÃO DINÂMICA (LOOPS DE VALIDAÇÃO SOBERANA)
--------------------------------------------------------------------------------
Toda fase do projeto deve rodar sob um ecossistema trifásico de prompts (Gerador, Auditor/Portão e Corretor), mas com controle final obrigatório do Humano. O fluxo segue estritamente esta máquina de estados:

1. Geração / Evolução: A IA recebe os inputs disponíveis e executa o prompt gerador (`PROMPT-GENERATE-[FASE].md`). 
   - NOTA DE ARQUITETURA (FASE 5): Nesta fase final, o gerador adota uma estrutura modular desacoplada, gerando um índice central (`USER-STORIES-{PROJECT_ID_NAME}.md`) e múltiplos arquivos atômicos na pasta `/user-stories/`.
2. Auditoria Interna da IA: O artefato (ou o repositório modular na Fase 5) é enviado para o portão (`PROMPT-GATE-[FASE].md`).
   - SE A IA ENCONTRAR ERROS: Emite o status `[NÃO COMPLIANCE]`, coleta o feedback do humano, aciona o `PROMPT-FIX-[FASE].md` de forma cirúrgica (reparando apenas o arquivo atômico afetado) e retorna ao passo 2.
   - SE A IA NÃO ENCONTRAR ERROS: Avança para o passo 3 (Portão de Validação Humana).
3. Portão de Validação Humana (Pré-Compliance): A IA apresenta o documento e emite o status `[PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]`, fazendo 3 perguntas obrigatórias sobre: Aderência ao negócio, novos documentos e novos inputs textuais.
4. Lógica de Decisão Baseada nas Respostas do Humano:
   - CENÁRIO DE SUCESSO (Aprovação): Se o humano validar o documento e NÃO enviar novos arquivos ou inputs, a fase é dada por encerrada (`[STATUS: COMPLIANCE]`), o arquivo é congelado e a próxima fase é destravada.
   - CENÁRIO DE RETROCESSO (Evolução Incremental): Se o humano fornecer novos documentos ou novas informações de texto, o orquestrador DEVE retroceder ao passo 1 (`PROMPT-GENERATE-[FASE].md`), injetando o documento gerado até o momento + os novos insumos para uma atualização incremental e viva (sem descartar o histórico).

--------------------------------------------------------------------------------
MATRIZ DE RASTREABILIDADE AUTOMATIZADA (FASE 5 vs FASE 1)
--------------------------------------------------------------------------------
Antes de aprovar a conclusão da Fase 5 (User Stories), a IA deverá obrigatoriamente construir e validar uma Matriz de Rastreabilidade com os seguintes critérios de checagem cruzada:
1. Mapeamento de Dependência: Toda User Story criada na Fase 5 deve, obrigatoriamente, estar vinculada retroativamente: User Story -> Feature -> Epic -> Requisito de Negócio (BRD) -> Objetivo/Premissa do Project Charter.
2. Identificação de Órfãos: Identificar e alertar o usuário se existir alguma User Story que NÃO possua um objetivo correspondente no Project Charter (Evitando Escopo Oculto).
3. Verificação de Cobertura: Garantir que 100% dos Objetivos Estratégicos definidos no Project Charter foram atendidos por, pelo menos, uma User Story (Evitando Escopo Negligenciado).
4. Prompt de Auditoria Interna: A IA deve executar uma auto-análise comparando o conjunto de arquivos modulares da pasta `user-stories/` e o arquivo central `USER-STORIES-{PROJECT_ID_NAME}.md` contra o arquivo '01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md', gerando um relatório de conformidade (Pass/Fail) baseado em consistência conceitual, regras de negócio e termos técnicos equivalentes. O portão (`project-documents/PROMPT-GATE-USER-STORIES.md`) deve validar cada arquivo individual e a integridade dos links markdown ativos na matriz RTM.

--------------------------------------------------------------------------------
FASES DO ROADMAP COM CHECKPOINTS E PIPELINES DE PROMPTS
--------------------------------------------------------------------------------

1. Project Charter
   - Objetivo: Definição do escopo de alto nível, objetivos, premissas, restrições e governança com 14 seções macro.
   - Inputs: Documentos brutos originais do usuário, atas e questionamentos ativos.
   - Responsáveis: Project Management / Product Owner.
   - Entregáveis: Arquivo '01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md' validado pelo humano.
   - Pipeline Sequencial de Tarefas: Executar `project-documents/PROMPT-GENERATE-PROJECT-CHARTER.md` $\rightarrow$ Validar via `project-documents/PROMPT-GATE-PROJECT-CHARTER.md`. Aplicar loops de correção (`PROMPT-FIX`) ou loops de retrocesso por novos insumos conforme o Mecanismo de Orquestração até o aceite final e explícito do humano.

2. Business Requirements Document (BRD)
   - Objetivo: Traduzir o Project Charter em requisitos de negócio detalhados e regras de atendimento.
   - Inputs: '01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md' (validado e congelado) + Entrevistas com stakeholders.
   - Checkpoint de Rastreabilidade: Validar se cada requisito de negócio atende diretamente a pelo menos um Objetivo ou Premissa do Project Charter.
   - Responsáveis: Business Analyst / Product Owner.
   - Entregáveis: Arquivo '02-BRD-{PROJECT_ID_NAME}.md' validado pelo humano e rastreado.
   - Pipeline Sequencial de Tarefas: Executar `project-documents/PROMPT-GENERATE-BRD.md` $\rightarrow$ Validar via `project-documents/PROMPT-GATE-BRD.md`. Aplicar loops de correção ou retrocesso por novos insumos até o aceite humano final.

3. Epics (Épicos)
   - Objetivo: Agrupar os requisitos de negócio em grandes blocos de entrega de valor / funcionalidades macro.
   - Inputs: '02-BRD-{PROJECT_ID_NAME}.md' (validado e congelado).
   - Checkpoint de Rastreabilidade: Garantir que os Épicos cobrem a totalidade dos Requisitos de Negócio sem criar escopos extras não mapeados no BRD.
   - Responsáveis: Product Owner / Product Manager.
   - Entregáveis: Arquivo índice '03-EPICS-{PROJECT_ID_NAME}.md' aprovado pelo humano + coleção de arquivos atômicos individuais `/epics/EP-NNNN-{nome-slug}.md` com detalhamento completo de cada épico e matriz BRD×Épico×Jornada.
   - Pipeline Sequencial de Tarefas: Executar `project-documents/PROMPT-GENERATE-EPICS.md` $\rightarrow$ Validar via `project-documents/PROMPT-GATE-EPICS.md`. Aplicar loops de correção ou retrocesso por novos insumos até o aceite humano final.

4. Features (Funcionalidades)
   - Objetivo: Decompor os Épicos em funcionalidades menores, tangíveis e implementáveis pelo time de desenvolvimento.
   - Inputs: '03-EPICS-{PROJECT_ID_NAME}.md' (validado e congelado).
   - Checkpoint de Rastreabilidade: Vincular formalmente o ID da Feature ao ID do Épico de origem.
   - Responsáveis: Product Owner / Tech Lead.
   - Entregáveis: Arquivo índice '04-FEATURES-{PROJECT_ID_NAME}.md' aprovado pelo humano + coleção de arquivos atômicos individuais `/features/FEAT-EP-{EEEE}-{SSSS}-{nome-slug}.md` com detalhamento completo de cada feature, user stories, regras de negócio e matriz BRD×Épico/Jornada×Feature.
   - Pipeline Sequencial de Tarefas: Executar `project-documents/PROMPT-GENERATE-FEATURES.md` $\rightarrow$ Validar via `project-documents/PROMPT-GATE-FEATURES.md`. Aplicar loops de correção ou retrocesso por novos insumos até o aceite humano final.

5. User Stories (Histórias de Usuário) & Validação de Repositório Modular
   - Objetivo: Refinar as Features no formato ágil clássico com critérios de aceite exaustivos (Gherkin) em arquivos individuais separados por ID, consolidando a árvore de rastreabilidade em um índice central vivo.
   - Inputs: '04-FEATURES-{PROJECT_ID_NAME}.md' (validado e congelado) e o '01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md' original de negócio.
   - Checkpoint de Rastreabilidade Mestre: Execução da Matriz Automatizada Bidirecional. O portão lerá o arquivo 'USER-STORIES-{PROJECT_ID_NAME}.md' e validará a existência e conformidade de cada arquivo físico contido na pasta `/user-stories/`.
   - Responsáveis: Product Owner / Equipe de Desenvolvimento.
   - Entregáveis: Arquivo central '05-USER-STORIES-{PROJECT_ID_NAME}.md' com links markdown ativos + coleção de arquivos atômicos individuais `/user-stories/US-FEAT-{codigo-feature}-{SSSS}-{nome-da-user-story}.md` validados, onde `SSSS` é sequencial global (0001-9999).
   - Pipeline Sequencial de Tarefas: Executar `project-documents/PROMPT-GENERATE-USER-STORIES.md` para criar o repositório descentralizado. Enviar a estrutura para o `project-documents/PROMPT-GATE-USER-STORIES.md`. Se houver links quebrados ou falhas conceituais, rodar o `project-documents/PROMPT-FIX-USER-STORIES.md` de forma isolada no arquivo com defeito até o status COMPLIANCE FINAL aprovado pelo humano.

--------------------------------------------------------------------------------
MODELO DA MATRIZ DE RASTREABILIDADE DE ESCOPO (RTM)
--------------------------------------------------------------------------------
**Projeto:** Implementação do Portal de Autoatendimento do Cliente  
**Status de Auditoria:** 100% Alinhado (Pass)

| ID Obj. (Fase 1: Charter) | Descrição do Objetivo | ID Req. (Fase 2: BRD) | Descrição do Requisito de Negócio | ID Épico (Fase 3) | ID Feature (Fase 4) | ID User Story (Fase 5) | Descrição da User Story (História) | Status de Validação |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **OBJ-01** | Reduzir chamados de suporte em 30% | **REQ-01.1** | Permitir que o cliente altere seus dados cadastrais sozinho. | **EP-0001** | **FEAT-EP-0001-0001** | **US-FEAT-EP-0001-0001-0001** | Como cliente, quero atualizar meu e-mail para receber notificações corretas. | ✅ Aprovado |
| **OBJ-01** | Reduzir chamados de suporte em 30% | **REQ-01.2** | Disponibilizar download de 2ª via de fatura em PDF. | **EP-0001** | **FEAT-EP-0001-0002** | **US-FEAT-EP-0001-0002-0002** | Como cliente, quero baixar o PDF da fatura para realizar o pagamento. | ✅ Aprovado |
| **OBJ-01** | Reduzir chamados de suporte em 30% | **REQ-01.3** | Notificar cliente sobre faturas em atraso via e-mail. | **EP-0001** | **FEAT-EP-0001-0003** | **US-FEAT-EP-0001-0003-0003** | Como cliente, quero receber alertas de fatura vencida para evitar multas. | ✅ Aprovado |
| OBJ-02 | Aumentar a segurança dos dados | REQ-02.1 | Implementar autenticação em duas etapas (2FA). | EP-0002 | FEAT-EP-0002-0001 | US-FEAT-EP-0002-0001-0004 | Como usuário, quero ativar o 2FA via SMS para proteger meus dados. | ✅ Aprovado |
| OBJ-02 | Aumentar a segurança dos dados | REQ-02.2 | Implementar política de senhas fortes com expiração. | EP-0002 | FEAT-EP-0002-0002 | US-FEAT-EP-0002-0002-0005 | Como usuário, quero ser forçado a trocar minha senha a cada 90 dias. | ✅ Aprovado |


--------------------------------------------------------------------------------
FINALIZAÇÃO DO PROCESSO — GIT WORKFLOW (COMMIT, PUSH, PR, MERGE)
--------------------------------------------------------------------------------

Quando TODAS as 5 fases do roadmap estiverem concluídas e o usuário humano confirmar explicitamente que as alterações na documentação estão finalizadas, o orquestrador DEVE executar o pipeline Git descrito abaixo. Esta etapa é obrigatória e automatizada — o usuário NÃO precisa executar comandos Git manualmente.

### Gatilho de Ativação

O pipeline Git é acionado quando o usuário indicar que a documentação está concluída com frases como:
- "Finalizei as alterações"
- "Está tudo pronto"
- "Pode commitar"
- "Finalizar o processo"
- Ou qualquer variação que indique conclusão dos trabalhos

### Pré-condições

Antes de executar qualquer comando Git, valide:
1. `PROMPT_BRANCH` está definida e NÃO é `main`, `master` ou `develop`
2. O diretório atual é um repositório Git
3. O working tree está limpo ou possui alterações a serem commitadas

### Pipeline de Comandos (Execução Sequencial)

Execute os comandos abaixo em ordem estrita. Se qualquer comando falhar, interrompa o pipeline e reporte o erro ao usuário.

#### Passo F.1 — Git Add e Commit

```bash
git add -A
git commit -m "docs: atualização dos prompts e documentos do projeto — ${PROJECT_ID_NAME}

- Project Charter, BRD, Epics, Features, User Stories e Matriz RTM
- Gerado pelo Roadmap de Documentos v4.0
- Branch: ${PROMPT_BRANCH}

Co-Authored-By: Claude <noreply@anthropic.com>"
```

**Regra:** Se não houver alterações para commitar (working tree limpo), informe o usuário e pule para o Passo F.4 (Cleanup).

#### Passo F.2 — Git Push

```bash
git push origin ${PROMPT_BRANCH}
```

**Tratamento de erro:** Se a branch remota já existir (erro de non-fast-forward), pergunte ao usuário:
> ⚠️ A branch remota `${PROMPT_BRANCH}` já existe. Deseja forçar o push (`--force`)?
> - **SIM** → Executar `git push origin ${PROMPT_BRANCH} --force`
> - **NÃO** → Abortar o pipeline. O usuário deverá resolver manualmente.

#### Passo F.3 — Criar, Mergear e Fechar Pull Request via `gh`

```bash
gh pr create \
  --base main \
  --head ${PROMPT_BRANCH} \
  --title "docs: atualização de documentação — ${PROJECT_ID_NAME}" \
  --body "## 📄 Atualização de Documentação

**Projeto:** ${PROJECT_ID_NAME}
**Branch:** ${PROMPT_BRANCH}

### Documentos Atualizados
- 01-PROJECT-CHARTER-${PROJECT_ID_NAME}.md
- 02-BRD-${PROJECT_ID_NAME}.md
- 03-EPICS-${PROJECT_ID_NAME}.md
- epics/*.md
- 04-FEATURES-${PROJECT_ID_NAME}.md
- features/*.md
- 05-USER-STORIES-{PROJECT_ID_NAME}.md
- user-stories/*.md

### Checklist
- [x] Documentos gerados e validados
- [x] Rastreabilidade vertical verificada
- [x] Aprovação humana em todas as fases

🤖 Generated with [Claude Code](https://claude.com/claude-code)"
```

**Após criar a PR com sucesso**, faça o merge imediato:

```bash
gh pr merge --merge --delete-branch
```

**Flags do merge:**
- `--merge`: Usa merge commit (preserva o histórico completo; alternativa: `--squash` para achatar commits)
- `--delete-branch`: Remove a branch remota após o merge bem-sucedido

> ⚠️ **Nota sobre `--merge` vs `--squash` vs `--rebase`:**
> - `--merge` (padrão escolhido): Preserva o histórico real de commits. Ideal para documentação, pois mantém a trilha de auditoria de quem gerou o quê.
> - `--squash`: Achata todos os commits em um só. Use se o histórico intermediário não for relevante.
> - `--rebase`: Reaplica commits sem merge commit. Use se o histórico linear for exigido pelo projeto.
>
> Se o projeto tiver regras específicas de merge, o usuário pode solicitar a troca da estratégia.

#### Passo F.4 — Cleanup Local

Após o merge bem-sucedido, faça checkout para a branch base e remova a branch local:

```bash
git checkout main
git branch -d ${PROMPT_BRANCH}
```

### Resumo Final

Exiba um sumário ao usuário após a conclusão:

> **🎉 Processo de Documentação Finalizado!**
>
> | Etapa | Status |
> |---|---|
> | Commit | ✅ Realizado |
> | Push | ✅ Enviado para `origin/${PROMPT_BRANCH}` |
> | Pull Request | ✅ Criada e mergeada |
> | Branch Remota | ✅ Deletada |
> | Branch Local | ✅ Deletada |
> | Base | ✅ Voltou para `main` |
>
> **📁 Artefatos gerados em:** `${PROJECT_COMPLETE_PATH_NAME}/`

### Tratamento de Falhas

| Falha | Ação |
|---|---|
| `git commit` falha (nada a commitar) | Informar usuário, pular para Passo F.4 |
| `git push` falha (branch remota existe) | Perguntar sobre `--force` |
| `gh pr create` falha (PR já existe) | Informar usuário, perguntar se deseja fazer merge manual da PR existente |
| `gh pr merge` falha (conflitos) | Reportar conflitos, abortar. Branch NÃO será deletada. Usuário resolve conflitos manualmente |
| Qualquer outro erro | Interromper pipeline, reportar erro, NÃO deletar branches |

### Regra de Segurança (Gating Rule Final)

**A branch local `${PROMPT_BRANCH}` NUNCA deve ser deletada se o merge falhar.** Isso garante que o trabalho do usuário nunca seja perdido. A deleção da branch local (Passo F.4) só ocorre após confirmação de que o merge foi bem-sucedido.

---

## LOCALIZAÇÃO DOS PROMPTS

Os prompts de geração, gate e correção de cada fase estão na pasta `project-documents/`:

```
.specs/prompts/project-documents/
├── FLOWCHART-ROADMAP-GENERATE-PROJECT-DOCUMENTS.md
├── PROMPT-GENERATE-PROJECT-CHARTER.md              ← Fase 1
├── PROMPT-GATE-PROJECT-CHARTER.md
├── PROMPT-FIX-PROJECT-CHARTER.md
├── PROMPT-GENERATE-BRD.md                          ← Fase 2
├── PROMPT-GATE-BRD.md
├── PROMPT-FIX-BRD.md
├── PROMPT-GENERATE-EPICS.md                        ← Fase 3
├── PROMPT-GATE-EPICS.md
├── PROMPT-FIX-EPICS.md
├── PROMPT-GENERATE-FEATURES.md                     ← Fase 4
├── PROMPT-GATE-FEATURES.md
├── PROMPT-FIX-FEATURES.md
├── PROMPT-GENERATE-USER-STORIES.md                 ← Fase 5
├── PROMPT-GATE-USER-STORIES.md
└── PROMPT-FIX-USER-STORIES.md
```

**Total:** 1 orquestrador + 5 geradores + 5 gates + 5 fixers = **16 prompts** (+ 1 flowchart).
