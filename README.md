
---

[![CI Generic Backend](https://github.com/bolismar69/workspace-fbso/actions/workflows/trigger-ci-generic-backend.yml/badge.svg)](https://github.com/bolismar69/workspace-fbso/actions/workflows/trigger-ci-generic-backend.yml)

---

![example event parameter]  (https://github.com/bolismar69/workspace-fbso/actions/workflows/trigger-ci-generic-backend.yml/badge.svg)

---

#### PUSH

![example event parameter](https://github.com/bolismar69/workspace-fbso/actions/workflows/trigger-ci-generic-backend.yml/badge.svg?event=push)

---

#### FEATURE

![example branch parameter](https://github.com/bolismar69/workspace-fbso/actions/workflows/trigger-ci-generic-backend.yml/badge.svg?branch=feature-1)

---




# 🗃️ Workspace Central FBSO - Monorepo de Engenharia


### referencias

Permissões do Workflow - Material

Documentação

https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#permissions

https://docs.github.com/en/actions/security-guides/automatic-token-authentication



----

# EXEMPLO DE SERVICE CONTAINERS

Service Containers - Material

Repositório

https://github.com/KubeDev/Review-Filmes

Pipeline Inicial


name: Teste Integração
on: 
    workflow_dispatch:  
jobs:
    build-test:
        runs-on: ubuntu-latest
        steps:
            - name: Checkout
              uses: actions/checkout@v4
            - name: Setup .NET
              uses: actions/setup-dotnet@v4
              with:
                dotnet-version: 8.0.x
            - name: Teste de Integração
              env:
                ConnectionStrings__DefaultConnection: "Host=localhost;Database=review;Username=review;Password=Passw0rd2023!"
              run: dotnet test ./src/Review-Filmes.Test.Integration/Review-Filmes.Test.Integration.csproj --verbosity normal 

Documentação

https://docs.github.com/en/actions/using-containerized-services

---

# Executando Self Hosted na Sua Máquina

Repositório Base:
https://github.com/kubedev/self-runner

Documentação:
https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/about-self-hosted-runners#communication-requirements


---

# Criando as proprias Actions - Introdução

Documentação:
https://docs.github.com/en/actions/creating-actions/about-custom-actions#types-of-actions

### Estrutura de uma Action

Documentação

https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions

https://feathericons.com/

Exemplo de Action

https://github.com/marketplace/actions/checkout


#### EXEMPLO ACTION COM JAVASCRIPT --- usa NCCC ( NPM ) PARA CONCENTAR TODO CODIGO NUM ARQUIVO UNICO
Action com Javascript

Action com Javascript

const core = require('@actions/core');
const github = require('@actions/github');

async function run() {
  try {
    // Obtem os inputs da action
    const token = core.getInput('GITHUB_TOKEN');
    const issueId = core.getInput('issue-id');
    const comment = core.getInput('comment');

    // Inicializa o cliente do GitHub
    const octokit = github.getOctokit(token);

    // Obtem o contexto do repositório
    const context = github.context;
    const { owner, repo } = context.repo;

    // Adiciona um comentário ao issue ou pull request
    const response = await octokit.rest.issues.createComment({
      owner: owner,
      repo: repo,
      issue_number: issueId,
      body: comment,
    });

    // Obtenha o ID do comentário criado
    const commentId = response.data.id;

    // Define o output com o ID do comentário
    core.setOutput('comment-id', commentId);

    console.log('Comentário adicionado com sucesso! ID do comentário:', commentId);
  } catch (error) {
    core.setFailed(`Erro ao adicionar comentário: ${error.message}`);
  }
}

run();

### ACTION COMPOSITE

Action Composite

Arquivo Action.yml

name: 'Workflow Evento'
author: 'Fabricio Veronez'
description: 'Adiciona um comentário a um issue ou pull request'
inputs:
  GITHUB_TOKEN:
    description: 'Token do GitHub'
    required: true
  issue-id:
    description: 'ID do issue ou pull request'
    required: true
  comment:
    description: 'Comentário a ser adicionado'
    required: true
outputs:
  comment-id:
    description: 'ID do comentário criado'
    value: ${{ steps.comment-action.outputs.comment-id }}
runs:
  using: "composite"
  steps:
    - name: Adicionando comentário
      shell: bash
      id: comment-action
      env:
        GITHUB_TOKEN: ${{ inputs.GITHUB_TOKEN }}
        ISSUE_ID: ${{ inputs.issue-id }}
        COMMENT: ${{ inputs.comment }}
      run: |
        GITHUB_TOKEN=$GITHUB_TOKEN
        ISSUE_ID=$ISSUE_ID
        COMMENT=$COMMENT

        # Inicialize o cliente do GitHub
        response=$(curl -s -X POST -H "Authorization: token ${GITHUB_TOKEN}" \
            -H "Accept: application/vnd.github.v3+json" \
            https://api.github.com/repos/${GITHUB_REPOSITORY}/issues/${ISSUE_ID}/comments \
            -d "{\"body\": \"${COMMENT}\"}")

        # Obtenha o ID do comentário criado
        comment_id=$(echo $response | jq -r .id)

        echo 

        # Define o output com o ID do comentário

        echo "comment-id=$(echo $comment_id)" >> $GITHUB_OUTPUT

        echo "Comentário adicionado com sucesso! ID do comentário: $comment_id"


###  ACTION DOCKER 
--- so usar com agentes de execução linux
--- performance baixa

#### EXEMPLO ACTION DOCKER
> Action Docker

```bash
# Arquivo main.sh

#!/bin/bash

# Obtenha os inputs da action
GITHUB_TOKEN=$1
ISSUE_ID=$2
COMMENT=$3
GITHUB_REPOSITORY=$4

# Inicialize o cliente do GitHub
response=$(curl -s -X POST -H "Authorization: token ${GITHUB_TOKEN}" \
    -H "Accept: application/vnd.github.v3+json" \
    https://api.github.com/repos/${GITHUB_REPOSITORY}/issues/${ISSUE_ID}/comments \
    -d "{\"body\": \"${COMMENT}\"}")

# Obtenha o ID do comentário criado
comment_id=$(echo $response | jq -r .id)

# Define o output com o ID do comentário
echo "comment-id=${comment_id}" >> $GITHUB_OUTPUT

echo "Comentário adicionado com sucesso! ID do comentário: $comment_id"
```
 
```dockerfile
Arquivo Dockerfile

# Use uma imagem base do Linux. Aqui, usamos o Debian por ser uma escolha popular.
FROM debian:latest

# Instale as dependências necessárias
RUN apt-get update && \
    apt-get install -y curl jq && \
    rm -rf /var/lib/apt/lists/*

# Copie o script shell para o container
COPY main.sh /main.sh

# Torne o script executável
RUN chmod +x /main.sh

# Defina o script como ponto de entrada
ENTRYPOINT ["/main.sh"]
```

### PUBLICAR ACTIONS NO GITHUB MARKETPLACE

Publicação no GitHub Actions Marketplace

Documentação

https://docs.github.com/en/actions/creating-actions/publishing-actions-in-github-marketplace


---

# GLOSSARIO:

