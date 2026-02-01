# Kestra - Integration Hub

Workflows e ativos para integração com parceiros, hubs e conectores.

## Estrutura sugerida por projeto

- `<parceiro_ou_conector>/flows/`
- `<parceiro_ou_conector>/scripts/`

## Regras

- Preferir namespaces `fbso.integration_hub.*`
- Padronizar contratos e mapeamentos como scripts/arquivos versionados
- Publicação via CI/CD sempre que possível
