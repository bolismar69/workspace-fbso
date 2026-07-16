---
mode: "ask"
applyTo: "**"
---

Description: Guia de boas práticas, comandos e fluxos para o desenvolvimento do app Solar Facil.
prompt: [
Você está desenvolvendo o app Solar Facil usando React Native com Expo, Redux Toolkit, Tailwind CSS e TypeScript.
Versões em uso:
- Node JS: v20.19.2
- npm: 10.8.2
- npx: 10.8.2
- Expo: ~53.0.9
- React Native: ^0.72.8
- TypeScript: ~5.8.3
- Tailwind CSS: ^3.4.17
As principais convenções de código são:,

- Arquivos de imagem em kebab-case.
- Arquivos de codigo em PascalCase,
- Componentes em './src/components'.
- Screens em './src/screens'.
- Serviços de API em './src/services'.
- Hooks personalizados em './src/hooks'.
- Styles em ./src/styles/,
- Constantes em ./src/constants ,
- API em ./src/apis,
  variaveis e objetos programáticos em camelCase,
  O fluxo GitFlow é obrigatório:,
- develop → main → feature/fix/chore.
  Os commits seguem Conventional Commits.
  Linters: ESLint e Prettier devem ser executados antes de qualquer PR.
  Testes: rodar 'make test'.
  Builds: usar 'eas build --platform all'.
  Documentação completa em './docs/VSCODE_PROMPT.md'.
  tags: [development, react-native, expo, best-practices
