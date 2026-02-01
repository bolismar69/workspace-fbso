# Kestra (Workflows)

Este diretório contém os artefatos versionados do Kestra:

- `flows/`: definições YAML de workflows
- `scripts/`: SQL/Python/Shell e outros ativos referenciados pelos flows

## Domínios

- `enterprise-integrations/`: integrações enterprise (ex.: migração/convivência com Oracle Suite)
- `integration-hub/`: integrações com parceiros e hubs
- `etl/`: pipelines de dados
- `automation/`: automações operacionais

## Convenções

- `namespace`: `fbso.<dominio>.<sistema>.<contexto>`
- `id`: `snake_case`
- Segredos: nunca versionar; usar secrets/env do Kestra e documentar no README do projeto.

## CI/CD (publicação)

A publicação de flows deve ocorrer **preferencialmente via CI/CD**:

- Mudanças em `orchestration/kestra/**` disparam pipeline.
- Pipeline publica/atualiza flows no Kestra via API.
- Autenticação e URL do Kestra são fornecidas por `secrets`/`env` do ambiente.

Deploy de instâncias Kestra (infra, helm/terraform, sizing) fica em `devops/orchestration/kestra/`.

## Design doc

Veja o desenho de CI/CD incremental com rollback em `orchestration/kestra/DESIGN.md`.
