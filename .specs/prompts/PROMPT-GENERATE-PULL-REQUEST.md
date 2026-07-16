# Contexto:
  - Obter a documentação do relatorio de implementação `TASK-EXECUTED-{AAAA-MM-DD-HHMMSS}-[nome-da-feature-em-kebab-case].md` com o objetivo de atualizar o repositorio de codigo local e remoto, e realizar a abertura de pull-request

# Missão:
  - Realizar comando git (add, commit, push, gh , e demais necessarios) para atualizar os repositorios local e remoto
  - Criar pull-request vinculado a atualização que foi realizada para o repositorio remoto

# Protocolo de Execução:
  - executar comandos `git` para atualizar repositorio local, realizando os registros que vinculem a alteração realizada
  - executar comandos `gh` para atualizar repositorio remoto, realizando os registros que vinculem a alteração realizada
    - criar pull-request para vincular a atualização do repositorio remoto
  - Anexar no pull-request o documento `TASK-EXECUTED-{AAAA-MM-DD-HHMMSS}-[nome-da-feature-em-kebab-case].md`
  - Documentar a açao realizada, criando na pasta `@.specs/pull-requests/` arquivo tendo como nome `"PR_"+{codigo-da-pull-request}+"-"+{tipo-escopo-identificador-da-request}+"__"+{nome-da-feature-em-kebab-case}+".md"` (ex: `PR_6-feat-PRJ-FIN-2026-0001__reforma-tributaria-fases-0-1-2.md` )

---

# CONTEXTO:
  - Você finalizou a fase de desenvolvimento. O relatório documental detalhado desta implementação já foi gerado e está salvo na pasta `.specs/skill-output` com o padrão de nomenclatura `{AAAA-MM-DD-HHMMSS}-[nome-da-feature-em-kebab-case].md` ( ex: `2026-06-25-073636-reforma-tributaria-fases-0-1-2.md` ). 
  - Sua tarefa agora é localizar esse arquivo específico de documentação, extrair os dados dele e executar o fluxo de atualização dos repositórios local e remoto, além de abrir a Pull Request.

### 1. MAPEAMENTO DA DOCUMENTAÇÃO EXISTENTE
1. Localize o arquivo `{TASK-EXECUTED-MD}` existente que segue o padrão de nome `{AAAA-MM-DD-HHMMSS}-[nome-da-feature-em-kebab-case].md`, que encontra-se na pasta `.specs/skill-output`
2. Leia o conteúdo deste arquivo. Você usará as informações contidas nele (métricas, arquivos alterados, fases e pendências) para estruturar o relatório final no terminal e para alimentar a Pull Request.

### 2. FLUXO DE GIT & GITHUB CLI
Execute os comandos necessários no terminal para:
1. Criar e alternar para uma nova branch com nomenclatura semântica (ex: `feature/[nome-da-feature-em-kebab-case]`).
2. Verificar o status atual e adicionar ao stage (`git add`) apenas os arquivos modificados que pertencem ao escopo descrito no documento mapeado (evitando stage de arquivos não relacionados de outros projetos do monorepo).
3. Realizar o commit utilizando o padrão Conventional Commits (ex: `feat(PRJ/TASK/ISSUE/others-XXXX): <mensagem>` ).
4. Fazer o push da nova branch configurando o tracking upstream para o repositório remoto (`origin`).
5. Criar uma Pull Request (PR) utilizando a GitHub CLI (`gh pr create`) direcionada para a branch principal (ex: `main`), passando o caminho do arquivo `{TASK-EXECUTED-MD}` localizado como argumento do parâmetro `--body-file`.

### 3. RELATÓRIO DO PROCESSO DE PULL REQUEST (SAÍDA NO TERMINAL)
Após a execução bem-sucedida dos comandos acima, gere no terminal o relatório consolidado do processo em formato Markdown. Siga estritamente o modelo de mapeamento e estrutura abaixo, preenchendo os colchetes com os dados reais do Git e do arquivo documental que você localizou:

# 📌 Pull Request #[Número da PR] — [Título da PR]

* **URL:** <[URL real gerada pela gh CLI]>
* **Branch:** `[nome-da-sua-nova-branch]` → `[branch-destino]`
* **Data de criação:** [Data atual em formato AAAA-MM-DD]
* **Repositório:** [Organização/Nome do repositório]
* **Status:** 🟢 Aberta
* **Projeto:** [Caminho relativo para o PRD.md]
* **Sumário de implementação:** [Caminho relativo para o arquivo TASK-EXECUTED-*.md localizado]

---

## 📊 Estatísticas

| Métrica | Valor |
|---|---|
| Arquivos alterados | **[Quantidade baseada no git diff]** |
| Linhas adicionadas | **+[Quantidade]** |
| Linhas removidas | **−[Quantidade]** |
| Pacotes/Módulos novos ou afetados | **[Quantidade]** |
| Endpoints REST / Rotas | **[Quantidade]** |
| Schemas OpenAPI / Contratos | **[Quantidade]** |
| Tabelas SQL novas / Migrations | **[Quantidade]** |
| Testes unitários | **[Quantidade]** (100% passing) |
| GAPs/Requisitos implementados | **[Quantidade]** |
| TASKS.md concluído | **[Concluídas]/[Total]** ✅ |
| Tamanho imagem Docker (se aplicável) | **[Tamanho ou N/A]** |

---

## 🛠️ Ações Realizadas

### 1. Análise pré-PR ([Data])

| Aspecto | Estado |
|---|---|
| Branch anterior | [Nome] ⚠️ |
| Remote verificado | [URL do Remote] ✅ |
| Arquivos no working tree | [Quantidade total detectada] |

### 2. Criação da Feature Branch
Exiba o comando exato utilizado:
```bash
git checkout -b [nome-da-branch]
```

### 3. Stage Seletivo (Apenas Escopo do Projeto)
Exiba os comandos de `git add` utilizados para isolar o escopo do projeto.

#### Arquivos incluídos ([Quantidade])

| Grupo / Pasta | Tipo | Descrição |
|---|---|---|
| [Ex: .specs/] | [🆕 Novo / 🔄 Refactor] | [Breve descrição baseado no documento] |

#### Arquivos excluídos/ignorados ([Quantidade] — se houver outros projetos alterados no workspace)

| Grupo / Pasta | Motivo |
|---|---|
| [Nome] | [Contexto diferente / Fora do escopo desta feature] |

### 4. Commit
Exiba o comando executado:
```bash
git commit -m "[Mensagem usada]"
```

| Campo | Valor |
|---|---|
| Tipo | [feat/fix/docs/etc] |
| Escopo | [ID do projeto] |
| Hash | [Hash curto gerado pelo git log -1] |
| Co-autoria | `Co-Authored-By: Claude <noreply@anthropic.com>` |

### 5. Push
Exiba o comando executado:
```bash
git push -u origin [nome-da-branch]
```

### 6. Criação da Pull Request
Exiba o comando exato da CLI usado para criar a PR (garantindo o uso do --body-file correto):
```bash
gh pr create --base ... --head ... --title ... --body-file "[Caminho do TASK-EXECUTED-*.md]"
```

### 7. Resultado
✅ **PR #[Número] criada**: [Link da URL gerada]

---

## 📦 Detalhamento por Fase
(Replique a tabela de fases/GAPs extraída de dentro do arquivo TASK-EXECUTED-*.md localizado)

| FASE / Onda | Feature / GAP | Artefatos Criados ou Modificados |
|---|---|---|
| [Fase] | [Código do Requisito] | [Caminho dos arquivos] |

---

## 🔒 Segurança
(Replique as informações de segurança do documento de origem)

| Controle | Status |
|---|---|
| [Mecanismo] | [Status] |

### ⚠️ Vulnerabilidade ou Dívida Técnica pendente (Se aplicável)
- [Transcrever observações de risco de segurança encontradas no documento original]

---

## 📄 Documentos Vinculados

| Documento | Localização |
|---|---|
| PRD | [Caminho] |
| Tasks | [Caminho do TASK-EXECUTED-*.md localizado] (100% ✅) |
| Test Plan | [Caminho] |

---

## 🧪 Evidências de Testes
(Insira o log de sucesso dos testes que consta na documentação ou capture a saída do terminal)

---

## 📋 Dívidas Técnicas Remanescentes

| DT | Descrição | Impacto |
|---|---|---|
| [Código] | [Descrição da pendência futura] | [Baixo/Médio/Alto] |

---

## 🔗 Links

- **Pull Request:** <https://github.com/bolismar69/workspace-fbso/pull/6>
- **Repositório:** <https://github.com/bolismar69/workspace-fbso>
- **Branch:** `feature/reforma-tributaria-2026-fases-0-1-2`
- **Commit:** `83f6905`
- **Projeto:** `PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO`

---

🤖 *Registro gerado em 2026-06-25. Histórico completo da criação da PR #6 para consulta humana e IA.*
