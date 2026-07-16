---
title: "C4 — Containers — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
diagram_type: "C4 — Nível 2 (Container)"
---

# C4 — Nível 2: Containers

## Diagrama de Containers

```mermaid
C4Container
    title Container diagram for Solar Fácil

    Person(associado, "Associado", "Usuário do app Solar Fácil")

    System_Boundary(solar_facil_system, "Solar Fácil — Ecossistema Mobile") {
        Container(ios_app, "iOS App", "Expo/React Native", "Aplicativo iOS distribuído via App Store. Build gerenciado pelo EAS.")
        Container(android_app, "Android App", "Expo/React Native", "Aplicativo Android distribuído via Google Play Store. Build gerenciado pelo EAS.")
        Container(metro_bundler, "Metro Bundler", "JavaScript", "Empacotador JavaScript para desenvolvimento local. Porta 8081.")
        ContainerDb(sqlite, "SQLite Database", "expo-sqlite", "Banco de dados local com tabelas: associados, movimentacoes. Arquivo: solarfacil.db")
        Container(async_storage, "AsyncStorage", "React Native", "Armazenamento key-value legado para dados de associados, beneficiados, fornecedores e movimentações")
        Container(mock_data, "Mock Data", "JSON Files", "Dados estáticos de referência: concessionárias, FAQs, planos comerciais, consumo médio")
    }

    System_Ext(expo_services, "Expo Services", "OTA Updates, EAS Build, EAS Submit")
    System_Ext(app_store, "Apple App Store", "Distribuição iOS")
    System_Ext(google_play, "Google Play Store", "Distribuição Android")

    Rel(associado, ios_app, "Usa", "iOS")
    Rel(associado, android_app, "Usa", "Android")
    Rel(ios_app, sqlite, "Lê/Escreve dados transacionais")
    Rel(android_app, sqlite, "Lê/Escreve dados transacionais")
    Rel(ios_app, async_storage, "Lê/Escreve (legado)")
    Rel(android_app, async_storage, "Lê/Escreve (legado)")
    Rel(ios_app, mock_data, "Lê dados de referência")
    Rel(android_app, mock_data, "Lê dados de referência")
    Rel(metro_bundler, ios_app, "Serve JS bundle (dev)", "localhost:8081")
    Rel(metro_bundler, android_app, "Serve JS bundle (dev)", "localhost:8081")
    Rel(ios_app, expo_services, "OTA Updates", "HTTPS")
    Rel(android_app, expo_services, "OTA Updates", "HTTPS")
    Rel(expo_services, app_store, "Publica build")
    Rel(expo_services, google_play, "Publica build")
```

## Descrição dos Containers

| Container | Tecnologia | Responsabilidade |
|---|---|---|
| **iOS App** | Expo/React Native (Swift bridge) | Aplicativo nativo iOS com New Architecture habilitada |
| **Android App** | Expo/React Native (Kotlin bridge) | Aplicativo nativo Android com edge-to-edge e New Architecture |
| **Metro Bundler** | JavaScript (Node.js) | Desenvolvimento local — hot reload, HMR na porta 8081 |
| **SQLite Database** | expo-sqlite (~15.2.13) | Fonte primária de dados — tabelas `associados` e `movimentacoes` |
| **AsyncStorage** | @react-native-async-storage (2.1.2) | Armazenamento key-value (legado em migração para SQLite) |
| **Mock Data** | JSON estático | Dados de referência (concessionárias, FAQs, planos) |

## Comunicação entre Containers

### Desenvolvimento Local

```
Metro Bundler :8081 ──JS Bundle──► iOS Simulator / Android Emulator
                                  │
                                  ├── SQLite (local)
                                  ├── AsyncStorage (local)
                                  └── Mock Data (JSON local)
```

### Produção

```
iOS/Android App (JS bundle empacotado)
  ├── SQLite (local)
  ├── AsyncStorage (local)
  └── Mock Data (JSON local)

Expo Updates (expo.dev) ──OTA──► App (atualiza JS bundle sem nova build)
```
