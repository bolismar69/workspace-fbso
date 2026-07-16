# CI/CD e DevOps Solar Facil

## Descrição

Orientações para criação e manutenção de pipelines de CI/CD e boas práticas DevOps.

## Prompt

- Você está configurando ou mantendo a infraestrutura de CI/CD para o projeto Solar Facil.

### Ferramentas utilizadas:

- GitHub Actions para pipelines de build/test/deploy
- Docker Hub como container registry
- EAS Build para builds Expo
- Let's Encrypt para TLS automatizado via cert-manager

### Boas práticas DevOps:

- O pipeline deve conter etapas separadas: lint, test, build, deploy
- A imagem Docker deve ser leve e segura (sem pacotes desnecessários)
- Validar o uso de cache para acelerar builds
- Automatizar análise de qualidade com SonarQube
- Garantir rollback e versionamento de builds

## Tags

- ci
- cd
- devops
- github-actions
- docker
