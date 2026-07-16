---
title: "Arquitetura de Plataforma — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["platform", "ios", "android", "cross-platform"]
---

# Arquitetura de Plataforma — Solar Fácil

## 1. Visão Geral

O Solar Fácil é um app **cross-platform** usando Expo managed workflow. A mesma base de código TypeScript/React Native compila para iOS e Android, com adaptações específicas por plataforma quando necessário.

## 2. Configurações por Plataforma

### 2.1. iOS

| Configuração | Valor | Local |
|---|---|---|
| Versão Mínima | iOS 15.0 | `app.json` → `expo.ios` |
| Suporte a Tablet | Sim (`supportsTablet: true`) | `app.json` |
| Bundle Identifier | `com.janeves.solarfacil` | `app.json` |
| Criptografia | Non-exempt (`ITSAppUsesNonExemptEncryption: false`) | `app.json` |
| Edge-to-Edge | N/A (iOS) | — |

### 2.2. Android

| Configuração | Valor | Local |
|---|---|---|
| SDK Mínimo | 24 (Android 7.0) | [TODO] Expo SDK 53 default |
| Edge-to-Edge | Habilitado (`edgeToEdgeEnabled: true`) | `app.json` |
| Package | `com.janeves.solarfacil` | `app.json` |
| Adaptive Icon | `solar-facil-adaptive-icon.png` (fundo #ffffff) | `app.json` |

## 3. APIs Nativas Utilizadas

| API | iOS | Android | Biblioteca |
|---|---|---|---|
| SQLite | ✅ | ✅ | `expo-sqlite` |
| Imagens | ✅ | ✅ | `expo-image`, `expo-image-picker` |
| Haptics | ✅ | ✅ | `expo-haptics` |
| Fontes | ✅ | ✅ | `expo-font` |
| Safe Area | ✅ | ✅ | `react-native-safe-area-context` |
| Gestos | ✅ | ✅ | `react-native-gesture-handler` |
| Animações (UI Thread) | ✅ | ✅ | `react-native-reanimated` |
| Web Browser | ✅ | ✅ | `expo-web-browser` |
| Status Bar | ✅ | ✅ | `expo-status-bar` |
| Blur | ✅ | ✅ | `expo-blur` |
| Sistema UI | ✅ | ✅ | `expo-system-ui` |
| Splash Screen | ✅ | ✅ | `expo-splash-screen` |
| DateTime Picker | ✅ | ✅ | `@react-native-community/datetimepicker` |
| Picker | ✅ | ✅ | `@react-native-picker/picker` |
| SF Symbols | ✅ | ❌ (iOS only) | `expo-symbols` |
| Updates (OTA) | ✅ | ✅ | `expo-updates` |

## 4. Diferenças de Comportamento

### 4.1. Safe Areas

- **iOS**: `SafeAreaView` cobre notch, Dynamic Island e home indicator (iPhones sem botão físico)
- **Android**: `SafeAreaView` cobre status bar e navigation bar (edge-to-edge)

### 4.2. Animações

- **iOS**: Reanimated roda na UI thread nativamente
- **Android**: Reanimated requer `react-native-reanimated/plugin` no Babel config

### 4.3. Ícones

- **iOS**: Suporte a SF Symbols via `expo-symbols`
- **Android**: Material Icons via `@expo/vector-icons`

### 4.4. Font Scaling

- **iOS**: Dynamic Type — usuário pode aumentar fonte nas configurações do sistema
- **Android**: Font Size nas configurações de acessibilidade
- **App**: NativeWind + sistema de temas não tratam font scaling explicitamente [TODO]

## 5. Módulos Nativos & Bridging

O projeto usa **Expo managed workflow**, então não há módulos nativos customizados diretos. Toda comunicação com código nativo passa pelos packages Expo:

```
JS Thread ←→ Bridge ←→ Native Modules (Expo packages)
```

### Plugins Expo (app.json)

| Plugin | Propósito |
|---|---|
| `expo-router` | File-based routing |
| `expo-splash-screen` | Splash screen config (imagem 200px, fundo branco) |
| `expo-sqlite` | Banco de dados SQLite |

## 6. Build & Deploy

### 6.1. Perfis EAS Build

| Perfil | iOS | Android | Uso |
|---|---|---|---|
| Development | Sim | Sim | Desenvolvimento com dev client |
| Preview | Sim | Sim | Testes internos (distribuição interna) |
| Production | Sim | Sim | Publicação nas lojas |

### 6.2. Expo Updates

| Branch | Propósito |
|---|---|
| development | Atualizações de desenvolvimento |
| preview | Atualizações de preview |
| production | Atualizações de produção |

## 7. Permissões

### 7.1. iOS (Info.plist)

| Permissão | Chave | Necessária? |
|---|---|---|
| Galeria de Fotos | `NSPhotoLibraryUsageDescription` | Sim (image picker) |
| Câmera | `NSCameraUsageDescription` | Sim (image picker) |

### 7.2. Android (AndroidManifest.xml)

| Permissão | Necessária? |
|---|---|
| `READ_EXTERNAL_STORAGE` | Sim (image picker) |
| `INTERNET` | Sim (OTA updates, debug) |

## 8. Observações

- **New Architecture**: habilitada em ambas plataformas (`newArchEnabled: true`) — usa Fabric renderer + TurboModules
- **react-native-web**: o projeto tem suporte a web via `expo-web-browser` e `react-native-web`, mas o foco primário é mobile
- [TODO]: Testar comportamento com font scaling em ambas plataformas
- [TODO]: Configurar ProGuard/R8 para Android (ofuscação) e ofuscação iOS
- [TODO]: Verificar compatibilidade de SF Symbols com versões antigas do iOS
