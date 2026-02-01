# Design Doc — CI/CD de Flows Kestra (Incremental + Rollback)

Data: 2026-01-22

## Objetivo

Definir um processo padrão de CI/CD para publicação de flows do Kestra versionados em `orchestration/kestra/**`, com:

- **Validação** (fail-fast) antes de publicar
- **Deploy incremental** (somente o que mudou)
- **Rollback automático** (lote “tudo ou nada” do ponto de vista do pipeline)
- Separação clara entre:
  - **Artefatos** (flows/scripts) em `orchestration/`
  - **Runtime/infra** em `devops/`

## Escopo

Inclui:

- Estratégia e passos do pipeline
- Detecção de mudanças
- Modelo de backup/restore para rollback
- Convenções mínimas para permitir deploy endereçável (namespace/id)

## Fora de escopo (neste doc)

- Implementação completa em GitHub Actions / scripts
- Provisionamento de cluster, Ingress, certificados, observabilidade (isso fica em `devops/`)
- Gestão corporativa de identidade (OIDC/SSO) — inicialmente será Basic Auth

## Premissas e contexto atual

- Hoje: piloto local.
- Alvo: Kestra em cloud na Digital Ocean dentro de um cluster Kubernetes.
- Autenticação inicial: **Basic Auth** (`user:pass`).
- Requisito forte: deploy incremental, com rollback quando qualquer flow do lote falhar.

## Princípios

1. **Git é a fonte da verdade** para definições de flows e scripts.
2. **Não publicar em PR** (PR valida; `main`/release publica).
3. **Validação local antes de deploy** (preferencialmente via CLI do Kestra).
4. **Rollback é responsabilidade do pipeline** (não assumir “transação” do server).
5. **Segredos nunca no repositório** (apenas Secrets/Secret Store do ambiente).

## Terminologia

- **Flow**: definição YAML do Kestra contendo pelo menos `namespace` e `id`.
- **Lote (batch) de deploy**: conjunto de flows alterados em um push/commit.
- **Backup**: cópia da definição atual no server antes de aplicar a nova.

## Estrutura de diretórios (referência)

- `orchestration/kestra/<dominio>/<projeto>/flows/*.yml|yaml`
- `orchestration/kestra/<dominio>/<projeto>/scripts/**`

Observação: scripts podem ser publicados separadamente (depende do seu padrão), mas neste desenho o **objeto “flow” é a unidade mínima** de rollback.

## Convenções mínimas obrigatórias

Para permitir deploy incremental e rollback, cada flow deve ser endereçável. Regras:

- `namespace`: `fbso.<dominio>.<sistema>.<contexto>`
- `id`: `snake_case`

O pipeline deve ser capaz de extrair `namespace` e `id` do YAML.

## Opções de interface com Kestra

### Opção A — CLI (preferida para validação)

Comandos citados:

- Validar: `/app/kestra flow validate --local <pasta>`
- Deploy: `/app/kestra flow updates --no-delete --user='<USER>:<PASSWORD>' <pasta>`

Notas:

- Fixar a versão do CLI (ex.: rodando em container `kestra/kestra:<tag>`), para evitar drift.
- Confirmar a sintaxe exata do comando (`update` vs `updates`) e flags.

### Opção B — API (necessária/útil para backup/restore)

Mesmo que o deploy seja via CLI, o rollback geralmente fica mais simples com API para:

- **Exportar** o estado atual do flow por `(namespace,id)`.
- **Reaplicar** o YAML de backup no rollback.

Se o Kestra oferecer “revisões” com rollback nativo, esta opção pode ser simplificada.

### Opção C — GitHub Actions oficiais (validate-action / deploy-action)

O Kestra disponibiliza Actions prontas para:

- Validar flows (ex.: `kestra-io/validate-action`)
- Publicar flows (ex.: `kestra-io/deploy-action`)

Vantagens:

- Menos código/"glue" no monorepo.
- Padroniza autenticação e chamadas contra o server.
- Reduz risco de drift de comandos do CLI.

Pontos de atenção (importantes para este monorepo):

- **Pinagem de versão**: evitar `@master` / `@develop` em produção; preferir tags/SHAs fixos para reprodutibilidade.
- **Deploy incremental**: as actions normalmente operam por diretório/namespace; o incremental por diff pode exigir um passo anterior para preparar um diretório contendo apenas os flows alterados.
- **Rollback transacional**: as actions de deploy não garantem por si só rollback “tudo ou nada”; ainda é necessário backup/restore (via API/CLI) caso o requisito permaneça.

## Estratégia recomendada (híbrida)

- **Validar** via CLI ou Action oficial (shift-left, fail-fast)
- **Backup/restore** via API (ou CLI equivalente) para construir rollback determinístico
- **Deploy** pode ser via CLI, API ou Action oficial, desde que seja idempotente

Observação: se a Action oficial cobrir validação e deploy de forma confiável, ela deve ser preferida para reduzir manutenção, mantendo backup/restore separado para atender rollback.

## Pipeline proposto

### 1) Pipeline de PR (sem deploy)

Gatilho: PR com mudanças em `orchestration/kestra/**`.

Passos:

1. Identificar flows alterados.
2. Validar flows alterados (CLI local ou Action oficial de validação).
3. (Opcional) Linters adicionais: YAML lint, regras internas de naming/namespace.

Critério de aprovação:

- 100% dos flows alterados validam.

### 2) Pipeline de main/release (deploy incremental)

Gatilho: push em `main` (ou tag release) com mudanças em `orchestration/kestra/**`.

Passos:

1. **Detectar alterações**
   - Coletar lista de arquivos alterados no commit range.
   - Filtrar apenas `**/flows/*.yml|yaml`.
   - Normalizar para conjunto de flows (evitar duplicidade).

2. **Extrair chaves (namespace,id)**
   - Parse leve do YAML para obter `namespace` e `id`.
   - Falhar se não houver `namespace` ou `id`.

3. **Validar (CLI ou Action oficial)**
   - Validar cada flow alterado (ou o diretório contendo apenas os alterados).
   - Se 1 falhar, abortar.

4. **Backup pré-deploy**
   - Para cada flow alterado, buscar a definição atual no server.
   - Persistir em artefato do job (ex.: `backup/<namespace>/<id>.yaml`).
   - Registrar quais flows são “novos” (não existiam no server).

5. **Aplicar deploy incremental**
   - Aplicar flows alterados, um a um, ou em lote.
   - Registrar quais foram aplicados com sucesso.

6. **Rollback automático (se falhar)**
   - Se algum deploy falhar:
     - Para cada flow já aplicado no lote, restaurar a versão de backup.
     - Para flows “novos”, decidir política:
       - (Preferido para consistência) deletar o flow criado no rollback, se a organização permitir.
       - (Alternativa) manter o flow criado, mas marcar incidente (não atende “estado inicial” estrito).
   - Encerrar pipeline com falha.

## Padrão de pipeline (pseudo-YAML)

Objetivo: ilustrar como combinar Actions oficiais do Kestra com incremental + backup/rollback.

Notas:

- Isto é **pseudo-YAML** (exemplo de desenho), não uma implementação final.
- Preferir Actions pinadas por tag/commit SHA (evitar `@master`/`@develop`).
- O rollback depende de existir um meio de **exportar** e **reaplicar** flows (API/CLI). Os comandos abaixo são placeholders.

```yaml
name: Kestra CI/CD (design)

on:
   pull_request:
      paths:
         - 'orchestration/kestra/**'
   push:
      branches: [ main ]
      paths:
         - 'orchestration/kestra/**'

jobs:
   validate_pr:
      if: ${{ github.event_name == 'pull_request' }}
      runs-on: ubuntu-latest
      steps:
         - uses: actions/checkout@v4

         # Valida o conjunto (ou apenas alterados) sem publicar
         - name: Validate flows (PR)
            uses: kestra-io/validate-action@<PINNED>
            with:
               directory: ./orchestration/kestra
               resource: flow
               server: ${{ secrets.KESTRA_HOSTNAME }}
               user: ${{ secrets.KESTRA_USER }}
               password: ${{ secrets.KESTRA_PASSWORD }}

   deploy_main:
      if: ${{ github.event_name == 'push' }}
      runs-on: ubuntu-latest
      steps:
         - uses: actions/checkout@v4
            with:
               fetch-depth: 0

         - name: Build changed_flows/
            run: |
               set -euo pipefail
               rm -rf changed_flows backups
               mkdir -p changed_flows backups

               # Pega apenas arquivos alterados em flows/
               git diff --name-only ${{ github.event.before }} ${{ github.sha }} \
                  | grep -E '^orchestration/kestra/.*/flows/.*\.(ya?ml)$' \
                  | sort -u > /tmp/changed.txt || true

               if [ ! -s /tmp/changed.txt ]; then
                  echo "Sem mudanças em flows. Encerrando."
                  exit 0
               fi

               # Copia mantendo estrutura relativa
               while IFS= read -r f; do
                  mkdir -p "changed_flows/$(dirname "$f")"
                  cp "$f" "changed_flows/$f"
               done < /tmp/changed.txt

         - name: Validate changed flows
            uses: kestra-io/validate-action@<PINNED>
            with:
               directory: ./changed_flows/orchestration/kestra
               resource: flow
               server: ${{ secrets.KESTRA_HOSTNAME }}
               user: ${{ secrets.KESTRA_USER }}
               password: ${{ secrets.KESTRA_PASSWORD }}

         - name: Backup current versions (before deploy)
            run: |
               set -euo pipefail
               # Placeholder:
               # - extrair (namespace,id) de cada YAML alterado
               # - exportar YAML atual do servidor para backups/<namespace>/<id>.yaml
               # Ex.: curl GET /api/v1/flows/{namespace}/{id}
               ./devops/scripts/kestra/backup_flows.sh \
                  --changed-list /tmp/changed.txt \
                  --out backups
            env:
               KESTRA_HOSTNAME: ${{ secrets.KESTRA_HOSTNAME }}
               KESTRA_USER: ${{ secrets.KESTRA_USER }}
               KESTRA_PASSWORD: ${{ secrets.KESTRA_PASSWORD }}

         - name: Deploy changed flows
            id: deploy
            uses: kestra-io/deploy-action@<PINNED>
            with:
               # Observação: se a Action exigir namespace explícito,
               # este passo pode precisar rodar por namespace ou por diretório.
               directory: ./changed_flows/orchestration/kestra
               resource: flow
               server: ${{ secrets.KESTRA_HOSTNAME }}
               user: ${{ secrets.KESTRA_USER }}
               password: ${{ secrets.KESTRA_PASSWORD }}
               delete: false

         - name: Rollback (if deploy failed)
            if: ${{ failure() }}
            run: |
               set -euo pipefail
               # Placeholder:
               # - reaplicar backups para todos os flows do lote
               # - para flows novos, aplicar política (delete ou bootstrap)
               ./devops/scripts/kestra/restore_flows.sh --in backups
            env:
               KESTRA_HOSTNAME: ${{ secrets.KESTRA_HOSTNAME }}
               KESTRA_USER: ${{ secrets.KESTRA_USER }}
               KESTRA_PASSWORD: ${{ secrets.KESTRA_PASSWORD }}
```

## Política para flows novos

Para cumprir o requisito “voltar ao estado inicial do deploy”, para flows novos:

- **Opção 1 (mais correta)**: pipeline deve conseguir **deletar** flow criado, se o lote falhar.
- **Opção 2 (sem delete)**: não cria flows novos em deploy incremental (exige etapa de “bootstrap” separada).

Recomendação: aceitar Opção 1, mas protegida por:

- deploy apenas em `main`
- approvals via GitHub Environments em produção
- logs e trilha de auditoria

## Idempotência e consistência

- Deploy deve ser idempotente: aplicar o mesmo YAML repetidamente não deve degradar o estado.
- Backup deve ser feito **antes** de qualquer alteração do lote.

## Segurança

- Credenciais (Basic Auth, token) ficam em GitHub Secrets/Environment Secrets.
- Recomenda-se migrar de Basic Auth para token/SSO assim que o runtime estiver em cloud.
- Evitar expor endpoint publicamente sem controles (allowlist, TLS, WAF/ingress rules).

## Observabilidade e auditoria

Requisitos mínimos:

- Pipeline registra:
  - commit SHA
  - lista de flows alterados
  - status por flow (validado/aplicado/rollback)
  - logs do servidor (quando aplicável)

## Falhas e cenários

- **Falha na validação**: não faz backup nem deploy.
- **Falha no backup**: aborta deploy (para não perder capacidade de rollback).
- **Falha no deploy do 3º flow**: executa rollback dos 2 aplicados.
- **Falha no rollback**: pipeline falha e gera incidente (pior caso; requer runbook).

## Questões em aberto (para fechar antes de produção)

1. O Kestra expõe endpoints/CLI para exportar e restaurar flows por `(namespace,id)`?
2. Existe mecanismo nativo de revisões/versionamento que permita rollback “nativo”?
3. Política para flows novos: deletar no rollback é aceitável?
4. Publicação de `scripts/`: é necessário upload separado ou basta versionar e referenciar (depende do runtime)?

## Próximos passos

- Confirmar a abordagem de execução:
   - CLI puro, Actions oficiais, ou híbrido.
- Fixar versões (CLI image tag e/ou Action tag/SHA) para reprodutibilidade.
- Definir se backup/restore será por API ou CLI.
- Definir política de flows novos.
- Só então implementar o workflow e scripts com testes de ponta a ponta em ambiente de staging.
