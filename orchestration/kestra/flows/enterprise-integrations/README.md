# Kestra - Enterprise Integrations

Workflows e automações para integrações enterprise e convivência/migração de stacks legadas.

## Estrutura sugerida por projeto

- `<projeto>/flows/`
- `<projeto>/scripts/`

## Regras

- Preferir namespaces `fbso.enterprise_integrations.*`
- Scripts versionados em `scripts/` (SQL/Python/Shell)
- Segredos apenas via secret store do Kestra
