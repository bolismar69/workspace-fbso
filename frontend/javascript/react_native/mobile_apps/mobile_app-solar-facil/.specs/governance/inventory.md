---
title: "Inventário do Projeto — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# Inventário do Projeto — Solar Fácil

## 1. Código Fonte

| Pasta | Arquivos | Linhas (est.) | Cobertura de Testes |
|---|---|---|---|
| `src/app/` | 9 | ~300 | 0% |
| `src/components/forms/` | 6 | ~800 | 0% |
| `src/components/inputs/` | 9 | ~600 | 0% |
| `src/components/lists/` | 1 | ~50 | 0% |
| `src/components/` (outros) | 7 | ~400 | 0% |
| `src/screens/associado/` | 6 | ~1200 | 0% |
| `src/screens/beneficiado/` | 2 | ~300 | 0% |
| `src/screens/fornecedor/` | 2 | ~300 | 0% |
| `src/screens/general/` | 4 | ~500 | 0% |
| `src/screens/movimentacao/` | 1 | ~200 | 0% |
| `src/context/` | 4 | ~200 | 0% |
| `src/hooks/queries/` | 5 | ~200 | 0% |
| `src/hooks/mutations/` | 7 | ~300 | 0% |
| `src/hooks/` (outros) | 1 | ~30 | 0% |
| `src/services/database/` | 4 (+11 órfãos) | ~500 | 0% |
| `src/services/storage/` | 5 | ~400 | 0% |
| `src/services/mock/` | 2 | ~100 | 0% |
| `src/services/` (outros) | 4 | ~100 | 0% |
| `src/types/` | 12 | ~200 | N/A |
| `src/styles/` | 2 | ~250 | N/A |
| `src/utils/` | 8 | ~200 | 0% |
| `src/constants/` | 1 | ~30 | N/A |
| `src/assets/` | 10 | N/A | N/A |
| `src/mocks/` | 4 | ~100 | N/A |

**Total Aproximado:** ~75 arquivos fonte, ~6500 linhas de código, **0% de cobertura de testes**.

## 2. Documentação

| Pasta | Arquivos | Status |
|---|---|---|
| `.specs/` | ~45 | ✅ Gerado (SPEC-MINING 2026-07-08) |
| `docs/codebase/` | 7 | ✅ Gerado |
| `README.md` | 1 | ✅ Existe |
| `README-ARQUITETURA.md` | 1 | ⚠️ Desatualizado |

## 3. Cobertura de Testes

| Tipo | Framework | Status |
|---|---|---|
| Unitários | Jest + RNTL | ❌ Não configurado |
| Integração | Jest + RNTL | ❌ Não configurado |
| E2E | Detox / Maestro | ❌ Não configurado |
| Acessibilidade | eslint-plugin-react-native-a11y | ❌ Não configurado |
| Lint | ESLint | ✅ Configurado |

## 4. Build & Deploy

| Ambiente | Status |
|---|---|
| Desenvolvimento | ✅ `npx expo start` |
| Build Preview | ✅ `eas build --profile preview` |
| Build Produção | ✅ `eas build --profile production` |
| OTA Updates | ✅ `expo-updates` configurado |
| CI/CD | ❌ GitHub Actions não configurado |
