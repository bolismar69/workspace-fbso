# Testes Automatizados Solar Facil

## Descrição

Guia para criação e execução de testes automatizados no app Solar Facil.

## Prompt

- Você está criando testes automatizados para o app Solar Facil.
- O projeto usa Jest e Testing Library para testes de unidade e integração.

### Regras importantes:

- Todos os componentes React devem ter testes de renderização.
- Fluxos críticos devem ter testes de integração.
- As funções de negócios devem ser testadas isoladamente.
- Os arquivos de teste devem ter extensão `.test.ts` ou `.spec.ts`.
- Os testes ficam na mesma pasta do componente ou função.
- Para rodar os testes, use o comando: `make test` ou `yarn test`.
- Garanta que todos os testes passem antes de qualquer pull request.

## Tags

- testing
- jest
- react-native
