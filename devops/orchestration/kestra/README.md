# Deploy / Runtime - Kestra

Este diretório contém **apenas** definições de infraestrutura e runtime do Kestra (instâncias, sizing, rede, observabilidade, backup, secrets, etc.).

## Princípios

- Definições de flows e scripts ficam em `orchestration/kestra/`.
- A **publicação de flows** deve ser feita via CI/CD (GitHub Actions), usando credenciais e endpoints do ambiente.
- Segredos nunca são versionados; o runtime provê secret store/variáveis.

## Itens típicos aqui

- Helm charts / manifests Kubernetes
- Terraform / IaC
- Integração com observabilidade (logs/metrics/tracing)
- Gestão de secrets no ambiente (ex.: Vault/Secrets Manager)

## CI/CD

O pipeline de publicação dos flows usa scripts em `devops/scripts/kestra/`.
