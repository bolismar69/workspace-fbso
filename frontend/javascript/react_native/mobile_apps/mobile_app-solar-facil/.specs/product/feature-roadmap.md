---
title: "Feature Roadmap — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
diataxis: "Reference"
---

# Feature Roadmap — Solar Fácil

## 1. Features Concluídas

| Feature | Status | Observações |
|---|---|---|
| Cadastro de Associados (PF/PJ, Fornecedor/Beneficiado) | ✅ MVP | SQLite + Validação yup |
| Login/Logout | ✅ MVP | CPF/CNPJ + Senha, estado em memória |
| Listagem de Associados | ✅ MVP | ScrollView (migrar para FlatList) |
| Movimentações Mensais | ✅ MVP | Cards + Gráfico Victory |
| Catálogo de Planos | ✅ MVP | Mock JSON |
| FAQ com Accordion | ✅ MVP | Mock JSON |
| Conteúdo Institucional (Saiba Mais) | ✅ MVP | Imagens + Texto |
| Tema Claro/Escuro | ✅ MVP | Segue preferência do sistema |
| Design System Base | ✅ MVP | NativeWind + Temas TypeScript |

## 2. Em Progresso / Planejado

| Feature | Prioridade | Esforço Estimado |
|---|---|---|
| Persistência de Sessão (AsyncStorage/SQLite) | P1 | 2 dias |
| Migração ScrollView → FlashList | P1 | 3 dias |
| Accessibility (labels, roles, touch targets) | P1 | 5 dias |
| Hash de Senhas (bcrypt/scrypt) | P1 | 2 dias |
| Error Boundary Global + Toast/SnackBar | P2 | 3 dias |
| Skeleton Loading + Empty States | P2 | 3 dias |
| Pull-to-Refresh em Listas | P2 | 1 dia |
| Onboarding (3-4 telas de tutorial) | P2 | 3 dias |
| Wizard de Cadastro (steps) | P2 | 5 dias |
| Consolidação de Navegação (8→5-6 tabs) | P3 | 2 dias |
| Animações de Transição entre Telas | P3 | 3 dias |
| Deep Links Funcionais | P3 | 2 dias |
| Testes Unitários (Jest + RNTL) | P2 | 5 dias |
| Testes E2E (Maestro) | P3 | 5 dias |

## 3. Backlog (Futuro)

| Feature | Dependência |
|---|---|
| Integração com Backend Remoto | Backend Java Spring API (`java-spring-api-solar-facil`) |
| Sincronização Offline→Online (fila) | Backend Remoto |
| Push Notifications (FCM/APNs) | Expo Notifications |
| Login Biométrico (Face ID / Fingerprint) | `expo-local-authentication` |
| Gráficos Avançados (projeção, comparação) | Victory ou react-native-gifted-charts |
| Dark Mode Customizado (não apenas sistema) | AppThemeContext |
| Exportação de Dados (PDF/CSV) | `expo-print` |
| Compartilhamento de Economia (redes sociais) | `expo-sharing` |
| Integração com Concessionárias (API ANEEL) | API externa |
| Multi-idioma (i18n) | `expo-localization` + `i18next` |
| Tablet Layout Otimizado | Responsividade |
| Wear OS / watchOS Companion App | N/A (longo prazo) |

## 4. Dívidas Técnicas (ver `docs/codebase/CONCERNS.md`)

| ID | Item | Prioridade |
|---|---|---|
| DT-001 | Senhas em plain text | P1 |
| DT-002 | Sem cobertura de testes | P1 |
| DT-004 | Auth apenas em memória | P1 |
| DT-005 | Duplicação SQLite vs AsyncStorage | P2 |
| DT-008 | Arquivos órfãos no source | P2 |
