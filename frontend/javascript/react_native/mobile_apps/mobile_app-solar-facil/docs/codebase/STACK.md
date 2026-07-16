---
title: "Stack Tecnológica — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["stack", "tecnologia", "mobile", "expo", "react-native"]
---

# Stack Tecnológica — Solar Fácil

## 1. Linguagem & Runtime

| Camada | Tecnologia | Versão | Notas |
|---|---|---|---|
| Linguagem | TypeScript | ~5.8.3 | Strict mode habilitado |
| Runtime | Node.js | 20.19.2 | Gerenciado via nvm |
| Framework Mobile | React Native | 0.79.3 | New Architecture habilitada (`newArchEnabled: true`) |
| Plataforma | Expo (managed) | ~53.0.9 | Expo SDK 53, Expo Router 5, EAS Build |
| Navegador (Web) | react-native-web | ~0.20.0 | Suporte web via Expo |

## 2. Navegação

| Biblioteca | Versão | Propósito |
|---|---|---|
| expo-router | ~5.0.6 | File-based routing com typed routes |
| @react-navigation/native | ^7.1.10 | Navegação nativa base |
| @react-navigation/native-stack | ^7.3.14 | Stack navigator |
| @react-navigation/bottom-tabs | ^7.3.14 | Bottom tab navigator |
| @react-navigation/stack | ^7.3.3 | Stack navigator (legado) |
| @react-navigation/elements | ^2.3.8 | Elementos de UI para navegação |
| react-native-screens | ~4.11.1 | Screens nativas para navegação |
| react-native-tab-view | ^4.1.2 | Tab view para navegação |
| react-native-pager-view | ^6.8.1 | Pager view para tabs |

## 3. Estilização & Design

| Biblioteca | Versão | Propósito |
|---|---|---|
| nativewind | ^4.0.1 | Tailwind CSS para React Native |
| tailwindcss | ^3.4.17 | Framework CSS utilitário |
| expo-blur | ~14.1.4 | Efeitos de blur |
| expo-system-ui | ~5.0.7 | Configuração de UI do sistema |
| @expo/vector-icons | ^14.1.0 | Ícones (Ionicons, MaterialIcons, SimpleLineIcons) |
| react-native-vector-icons | ^10.2.0 | Ícones nativos |
| expo-symbols | ~0.4.4 | SF Symbols (iOS) |

### Paleta de Cores (Tailwind)

| Token | Valor Hex | Uso |
|---|---|---|
| primary | #1E5631 | Verde escuro — cor principal, botões, títulos |
| secondary | #A4DE02 | Verde lima — destaques |
| accent | #FFD700 | Dourado — acentos |
| background | #F6F6F6 | Fundo base |
| neutral | #FFFFFF | Branco neutro |

## 4. Gerenciamento de Estado

| Biblioteca | Versão | Propósito |
|---|---|---|
| React Context API | built-in | Auth (AuthContext), Tema (AppThemeContext), Database (DatabaseContext) |
| @tanstack/react-query | ^5.83.0 | Gerenciamento de queries/mutations assíncronas para SQLite |
| React useState/useEffect | built-in | Estado local de componentes |

## 5. Formulários & Validação

| Biblioteca | Versão | Propósito |
|---|---|---|
| react-hook-form | ^7.56.4 | Gerenciamento de formulários |
| yup | ^1.6.1 | Schemas de validação |
| @hookform/resolvers | ^5.0.1 | Integração react-hook-form + yup |

## 6. Banco de Dados & Armazenamento

| Biblioteca | Versão | Propósito |
|---|---|---|
| expo-sqlite | ~15.2.13 | Banco de dados SQLite local (tabelas: associados, movimentacoes) |
| @react-native-async-storage/async-storage | 2.1.2 | Armazenamento key-value (fallback/legado) |

## 7. Animações & Gestos

| Biblioteca | Versão | Propósito |
|---|---|---|
| react-native-reanimated | ~3.17.4 | Animações de alta performance |
| moti | ^0.30.0 | Animações declarativas sobre Reanimated |
| react-native-gesture-handler | ~2.24.0 | Gestos touch (swipe, pinch, etc.) |
| expo-haptics | ~14.1.4 | Feedback tátil (vibração) |

## 8. Rede & Dados

| Biblioteca | Versão | Propósito |
|---|---|---|
| axios | ^1.9.0 | Cliente HTTP (mock services) |

## 9. Mídia & Sistema

| Biblioteca | Versão | Propósito |
|---|---|---|
| expo-image | ~2.2.0 | Imagens otimizadas |
| expo-image-picker | ~16.1.4 | Seleção de imagens |
| expo-web-browser | ~14.1.6 | Abertura de URLs externas |
| expo-linking | ~7.1.5 | Deep linking |
| expo-constants | ~17.1.6 | Constantes do sistema |
| expo-updates | ~0.28.17 | OTA updates via EAS |
| expo-splash-screen | ~0.30.8 | Splash screen |
| expo-status-bar | ~2.2.3 | Status bar |
| expo-font | ~13.3.1 | Fontes customizadas |
| expo-dev-client | ~5.2.4 | Dev client para builds custom |

## 10. Build & CI/CD

| Ferramenta | Versão | Propósito |
|---|---|---|
| EAS CLI | >= 16.14.1 | Build, submit e updates |
| EAS Build | — | Builds development/preview/production |
| Expo Updates | — | OTA updates via `u.expo.dev` |

## 11. Gráficos

| Biblioteca | Versão | Propósito |
|---|---|---|
| victory-native | ^41.17.3 | Gráficos e visualização de dados |

## 12. Desenvolvimento

| Ferramenta | Versão | Propósito |
|---|---|---|
| ESLint | ^9.25.0 | Linting |
| eslint-config-expo | ~9.2.0 | Config base Expo |
| @typescript-eslint/* | ^8.33.0 | Linting TypeScript |

## 13. Plataformas Alvo

| Plataforma | Versão Mínima | Notas |
|---|---|---|
| iOS | 15.0 | Suporte a tablet |
| Android | SDK 24 (Android 7.0) | Edge-to-edge habilitado |
| Web | — | Via react-native-web (bundler: metro) |

## 14. Observações

- **Não há backend remoto configurado**: a aplicação opera em modo offline-first com SQLite local e serviços mock.
- **New Architecture**: o projeto usa a New Architecture do React Native (`newArchEnabled: true` no app.json).
- **Typed Routes**: Expo Router com `typedRoutes: true` para navegação type-safe.
- **Expo Updates**: configurado para OTA updates via EAS (`runtimeVersion: appVersion`).
