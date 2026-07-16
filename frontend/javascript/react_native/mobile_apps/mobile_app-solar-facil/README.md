# Welcome to your Expo app 👋

This is an [Expo](https://expo.dev) project created with [`create-expo-app`](https://www.npmjs.com/package/create-expo-app).

## Get started

pre-reques:
nvm install v20
nvm use 20
node -v
npm -v

1. Install dependencies

   ```bash
   npm install
   ```

2. Start the app

   ```bash
   npx expo start
   ```
   ou
   ```bash
   npx expo start --clear
   ```

3. Em caos de testes, é melhor limpa a pasta .expo, para ela ser gerada novamente ( reprocessamentos podem deixar sujeiras nessa pasta ). Execute:
   ```bash
   rm -rf .expo .expo-shared && npx expo start --clear
   ```

4. Para executar com simulaodor de celular tem que criar a variavel de ambiente
   ```bash
   export ANDROID_HOME=/home/bolismar/Android/Sdk
   rm -rf .expo .expo-shared && npx expo start --clear
   ```

In the output, you'll find options to open the app in a

- [development build](https://docs.expo.dev/develop/development-builds/introduction/)
- [Android emulator](https://docs.expo.dev/workflow/android-studio-emulator/)
- [iOS simulator](https://docs.expo.dev/workflow/ios-simulator/)
- [Expo Go](https://expo.dev/go), a limited sandbox for trying out app development with Expo

You can start developing by editing the files inside the **app** directory. This project uses [file-based routing](https://docs.expo.dev/router/introduction).

## Get a fresh project

When you're ready, run:

```bash
npm run reset-project
```

This command will move the starter code to the **app-example** directory and create a blank **app** directory where you can start developing.

## Learn more

To learn more about developing your project with Expo, look at the following resources:

- [Expo documentation](https://docs.expo.dev/): Learn fundamentals, or go into advanced topics with our [guides](https://docs.expo.dev/guides).
- [Learn Expo tutorial](https://docs.expo.dev/tutorial/introduction/): Follow a step-by-step tutorial where you'll create a project that runs on Android, iOS, and the web.

## Join the community

Join our community of developers creating universal apps.

- [Expo on GitHub](https://github.com/expo/expo): View our open source platform and contribute.
- [Discord community](https://chat.expo.dev): Chat with Expo users and ask questions.

# INSTALACAO TAILWINDCSS ATRAVES DO NATIVEWIND

npx expo install nativewind@^4.0.1 react-native-reanimated tailwindcss

---

# Energia Solar Fácil

Aplicativo desenvolvido em **React Native** com **Expo** e estilizado com **Tailwind CSS (NativeWind)**.

## 🚀 Tecnologias

- React Native
- Expo Go
- NativeWind (Tailwind CSS)
- Node.js v20
- EAS Build (para builds de produção)

## 📦 Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/react-native-expo-app-solar-facil.git
   cd react-native-expo-app-solar-facil
   ```

2. Repositorios associados ao projeto:
repositorios associados:    
frontend - app -> git clone https://github.com/bolismar/react-native-expo-app-solar-facil.git
frontend - web -> git clone https://github.com/bolismar69/react-native-expo-app-solar-facil.git
backend - api -> git clone https://github.com/bolismar69/java-spring-api-solar-facil.git



# VARIAVEL PARA O ANDROID

export ANDROID_HOME=/home/bolismar/Android/Sdk

# Reinstalar as dependencias

## 1.Excluir arquivos antigos:

rm -rf node_modules package-lock.json
rm -rf .expo .expo-shared node_modules package-lock.json

## 2. Reinstalar as dependências:

npm install

## 3. Verificar por problemas:

npm audit fix

# para atender ao processo de validacao de dados e UPLOAD

npx expo install expo-image-picker
npm install react-native-text-input-mask
npm install axios
npm install victory-native
npm install yup
npm install @hookform/resolvers
npm install react-hook-form

# Tailwind CSS via NativeWind para estilização no React Native

##

npm install nativewind
npm install react-hook-form yup @hookform/resolvers
npm install react-navigation react-native-screens react-native-safe-area-context react-native-gesture-handler react-native-reanimated
npm install react-native-svg victory-native

# Para atender a controle de Query e Mutation no acesso a dados
npm install @tanstack/react-query

# Google Play Store - Geração imagens

Enviar:
.aab.
Capturas de tela (mínimo 2 para cada dispositivo).
Ícone: 512x512.
Feature Graphic: 1024x500.

## Tabela com cores:

https://www.homehost.com.br/blog/tutoriais/tabela-de-cores-html/
https://www.w3schools.com/cssref/css3_pr_background-origin.php

npm install -g expo-cli

# VALIDAR CDIGO

## identificar codigo nao usado

px eslint . --ext .ts,.tsx


# novas instalacoes para botoes com animcaao
npx expo install react-native-vector-icons
npx expo install moti react-native-reanimated

npm install @react-navigation/native @react-navigation/stack react-native-screens react-native-safe-area-context react-native-gesture-handler react-native-reanimated

npm install @react-navigation/stack

# PARA ICONES
## Obs: O Expo já inclui react-native-vector-icons como dependência.
npx expo install @expo/vector-icons

@react-navigation/native
@react-navigation/stack
react-native-screens
react-native-safe-area-context
react-native-gesture-handler
react-native-reanimated

@react-navigation/native
@react-navigation/stack
react-native-screens
react-native-safe-area-context
react-native-gesture-handler
react-native-reanimated


-- REACT NAVIGATION
npm install @react-navigation/native
npm install react-native-screens react-native-safe-area-context react-native-gesture-handler react-native-reanimated

-- TIPOS DO TYPESCRIPT
npm install --save-dev @types/react-navigation @types/react-navigation-stack


OK para - npm install @react-navigation/native
ok para - npm install react-native-screens react-native-safe-area-context react-native-gesture-handler react-native-reanimated


npm install react-native-tab-view react-native-pager-view
npm install @react-native-picker/picker
npm install @react-native-community/datetimepicker

-- PARA BANCO DE DADOS
=> FileSystem from "expo-file-system";
=> AsyncStorage from "@react-native-async-storage/async-storage";
npx expo install expo-sqlite


# TODO - PARA WEB
- Criar uma versão web-only sem moti, com fallback para animações simples.
- Incluir Platform.select() para adaptar animações ou renderizações.
- Gerar mocks de backend/API para testar fluxos mais completos em mobile.

# ============ 
## FORCAR CLONAR REPOSITORIO

1. Clone o repositório remoto em uma nova pasta
´´´bash
git clone git@github.com:bolismar69/react-native-expo-app-solar-facil.git temp-solar-facil
´´´

2. Copie o conteúdo da sua pasta atual para a nova pasta (sem a pasta .git).
´´´bash
rsync -av --exclude=".git" ~/projetos/react-native-expo-app-solar-facil/ ~/temp-solar-facil/
´´´

3. Vá para a nova pasta clonada
´´´bash
cd ~/temp-solar-facil
´´´

4. Confirme os arquivos copiados
´´´bash
git status 
´´´

5. Faça o commit e o push forçado para o GitHub
´´´bash
git add .
git commit -m "Atualização completa do projeto (sincronização forçada)"
git push origin main --force
´´

# GERAIS GERAIS

## site de permissionarias de energia eletrica
https://www.arsesp.sp.gov.br/paginas/energia/concessionarias_cooperaticas.aspx

## relatorios publicos sobre energia eletrica
https://www.gov.br/aneel/pt-br/centrais-de-conteudos/relatorios-e-indicadores/distribuicao/relatorios-distribuicao

## sites com fotos para ideias 
https://br.freepik.com/fotos-vetores-gratis/energia-solar-economia
https://br.freepik.com/fotos-vetores-gratis/economia-energia-placa-solar-limpa
https://www.istockphoto.com/br/fotos/economia-de-energia-solar?page=2



# ESTUDAR

## funcionalidades a inserir

1. usar as funções de select via native-picker
https://github.com/react-native-picker/picker


2. react-native-dropdown-picker
https://hossein-zare.github.io/react-native-dropdown-picker-website/docs/advanced/search


3. avaliar o uso de SQLite
https://www.youtube.com/watch?v=BJEACwKXWf8
https://www.sqlite.org/

# SUBIR APLICACAO NO EXPO.DEV

1. Crie uma conta na expo.dev, e crie projeto ( importante ter em mãos após criar o proeto, do ID do Projeto no expo.dev)

2. Na pasta do projeto, digitar os comandos, para inicializar nosso projeto para uma compilação de desenvolvimento, vamos entrar no diretório do projeto e executar o seguinte comando para instalar a biblioteca:
```bash
npx expo install expo-dev-client

```

3. Inicializar uma compilação de desenvolvimento:

3.1. Precisamos instalar a ferramenta EAS Command Line Interface (CLI) como uma dependência global em nossa máquina local:
```bash
npm install -g eas-cli

```

3.2. Faça login ou crie uma conta na Expo:
```bash
eas login

```

3.3. Inicializar e vincular o projeto ao EAS. Isso vai criar o arquivo eas.json na pasta raiz do projeto:
```bash
eas init

```

```
Foi criado o projeto:
✔ Would you like to create a project for @bolismar69/solar-facil? … yes
✔ Created @bolismar69/solar-facil
✔ Project successfully linked (ID: eaf32393-0c03-4e40-9005-fd5955794e5a) (modified app.json)
```

3.4. Abrir o arquivo app.json e verificar se foi criada a sessão "extra":"eas" com o projectId. Se não foi criado, crie conforme padrão abaixo:
```json
{
  "extra": {
    "eas": {
      "projectId": "0cd3da2d-xxx-xxx-xxx-xxxxxxxxxx"
    }
  }
}
```

3.5. Configurar projeto para EAS Build:
```bash
eas build:configure

```

```
Ao executar, este comando:
1. Solicita a seleção de uma plataforma: Android, iOS ou Todas. Como estamos criando aplicativos para Android e iOS, vamos selecionar Todas.
2. Cria eas.json na raiz do diretório do nosso projeto com a seguinte configuração:
```

```json (eas.json)
{
  "cli": {
    "version": ">= 14.2.0",
    "appVersionSource": "remote"
  },
  "build": {
    "development": {
      "developmentClient": true,
      "distribution": "internal"
    },
    "preview": {
      "distribution": "internal"
    },
    "production": {
      "autoIncrement": true
    }
  },
  "submit": {
    "production": {}
  }
}
```

3.6. Compilar e subi a aplicação no expo.dev

3.6.1. Build de Desenvolvimento
```bash
eas build --profile development --platform all

```

3.6.2. Build de Preview
```bash
eas build --profile preview --platform all

```

3.6.3. Build de Produção
```bash
eas build --profil e production --platform all

```

3.7. Subir para Expo.dev
```
Após o build ser concluído, o Expo automaticamente sobe o build para o Expo.dev. Você pode acessar o link gerado no terminal para visualizar o build.
```

3.8. Publicar a Aplicação (Se você deseja publicar a aplicação para que ela esteja disponível no Expo Go, use o comando):

3.8.1. para versões mais antigas do expo-cli:
```bash
npx expo publish

```

3.8.2. para versões mais recentes do expo-cli:
```bash
eas update --branch development
ou
eas update --branch preview
ou
eas update --branch production
ou
```

Será gerado um URL para ser usado no app Expo-go:
```
https://expo.dev/accounts/bolismar69/projects/solar-facil/fingerprints/dcaccc1eb230fa0790905c347db74b0e20de10d2
```






# GERAR ZIP DA APLICACAO
```bash
zip -r solar-facil.zip . -x "node_modules/*" "dist/*" ".expo/*"
```

