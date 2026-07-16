---
title: "Manifesto de Permissões — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
platforms: ["iOS", "Android"]
---

# Manifesto de Permissões — Solar Fácil

## 1. iOS (Info.plist via app.json)

| Permissão | Chave | Justificativa | Status |
|---|---|---|---|
| Galeria de Fotos | `NSPhotoLibraryUsageDescription` | `expo-image-picker` para upload de foto do associado | Necessário |
| Câmera | `NSCameraUsageDescription` | `expo-image-picker` para tirar foto | Necessário |
| Criptografia | `ITSAppUsesNonExemptEncryption: false` | App não usa criptografia customizada | Declarado |

## 2. Android (AndroidManifest.xml)

| Permissão | Nível | Justificativa |
|---|---|---|
| `INTERNET` | Normal | OTA Updates, debug (Metro bundler) |
| `READ_EXTERNAL_STORAGE` | Dangerous | Acesso a galeria para image picker |
| `VIBRATE` | Normal | Haptics feedback |

## 3. Permissões Futuras

| Permissão | Quando | Justificativa |
|---|---|---|
| `ACCESS_FINE_LOCATION` | Backend com geolocalização | Mapear concessionárias próximas |
| `POST_NOTIFICATIONS` (Android 13+) | Push notifications | Alertas de vencimento, economia |
| `USE_BIOMETRIC` | Login biométrico | Face ID / Fingerprint |
| `CAMERA` | Se adicionar leitura de código de barras | Leitura de conta de luz |

## 4. Princípios

- **Mínimo privilégio:** solicitar apenas permissões estritamente necessárias
- **Just-in-time:** solicitar permissão no momento do uso, não no startup
- **Clareza:** explicar ao usuário por que a permissão é necessária antes de solicitar
