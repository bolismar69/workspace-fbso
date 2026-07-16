---
title: "Segurança — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# Segurança — Solar Fácil

## 1. OWASP Mobile Top 10 (2024) — Avaliação

| # | Risco | Status | Ação |
|---|---|---|---|
| M1 | Improper Credential Usage | ❌ | Senhas em plain text — implementar bcrypt/scrypt + salt |
| M2 | Inadequate Supply Chain Security | ⚠️ | Executar `npm audit` regularmente; revisar dependências |
| M3 | Insecure Authentication/Authorization | ❌ | Auth apenas em memória; sem timeout de sessão; sem biometria |
| M4 | Insufficient Input/Output Validation | ✅ | yup + react-hook-form validam inputs |
| M5 | Insecure Communication | N/A | App offline-first (sem rede); [FUTURO] SSL pinning + HTTPS |
| M6 | Inadequate Privacy Controls | ⚠️ | Política de privacidade existe (README), mas não integrada ao app |
| M7 | Insufficient Binary Protections | ❌ | Sem ofuscação (ProGuard/R8); sem Root/Jailbreak detection |
| M8 | Security Misconfiguration | ⚠️ | `newArchEnabled: true` sem auditoria; logs em produção (`console.log`) |
| M9 | Insecure Data Storage | ❌ | Dados sensíveis (CPF, email, endereço) em plain text no SQLite |
| M10 | Insufficient Cryptography | ❌ | Sem criptografia de dados em repouso; sem Keychain/Keystore |

## 2. Autenticação

### Estado Atual
- Login via CPF/CNPJ + senha (plain text)
- Estado mantido em memória (React Context)
- Sem expiração de sessão
- Sem logout automático por inatividade

### Recomendações

```typescript
// 1. Hash de senha
import bcrypt from "react-native-bcrypt";

const salt = await bcrypt.genSalt(10);
const hash = await bcrypt.hash(senha, salt);
// Armazenar hash no SQLite, nunca plain text

// 2. Biometria (expo-local-authentication)
import * as LocalAuthentication from "expo-local-authentication";

const result = await LocalAuthentication.authenticateAsync({
  promptMessage: "Autentique-se para acessar o Solar Fácil",
});

// 3. Timeout de sessão
const SESSION_TIMEOUT = 30 * 60 * 1000; // 30 minutos
useEffect(() => {
  const timer = setTimeout(() => logout(), SESSION_TIMEOUT);
  const resetTimer = () => { /* reset on user activity */ };
  return () => clearTimeout(timer);
}, []);
```

## 3. Armazenamento Seguro

### Estado Atual
- Senhas, CPF, email, endereço → SQLite plain text
- Auth state → memória (volátil)

### Recomendações

```typescript
// Keychain (iOS) / Keystore (Android)
import * as SecureStore from "expo-secure-store";

// Salvar token/hash
await SecureStore.setItemAsync("authToken", token);

// Recuperar
const token = await SecureStore.getItemAsync("authToken");

// SQLite encryption (SQLCipher adapter)
// [TODO]: Avaliar expo-sqlite com SQLCipher
```

## 4. Ofuscação de Código

### Android (ProGuard/R8)

```properties
# android/app/proguard-rules.pro
-keep class com.facebook.hermes.** { *; }
-keep class com.facebook.react.** { *; }
```

### iOS

Habilitar ofuscação via EAS Build profile.

## 5. Comunicação de Rede

### Estado Atual
N/A — app é offline-first.

### Futuro (com backend)
- SSL Pinning via `expo-network` ou `react-native-ssl-pinning`
- Certificate Transparency verification
- HTTPS obrigatório — nunca HTTP

## 6. Logging & Monitoramento

### Recomendações

```typescript
// Remover console.log em produção
if (__DEV__) {
  console.log("Debug info");
}

// [TODO]: Implementar logging estruturado
// [TODO]: Crash reporting (Sentry via expo-sentry)
```

## 7. Checklist de Segurança

- [ ] Implementar hash de senhas (bcrypt/scrypt)
- [ ] Migrar dados sensíveis para Keychain/Keystore
- [ ] Configurar ofuscação de código Android/iOS
- [ ] Remover console.log de builds de produção
- [ ] Adicionar timeout de sessão
- [ ] Implementar biometria para login rápido
- [ ] Executar `npm audit` e corrigir vulnerabilidades
- [ ] Configurar SSL pinning (quando backend existir)
- [ ] Auditoria de dependências (Socket.dev ou Snyk)
- [ ] Adicionar política de privacidade integrada ao app
