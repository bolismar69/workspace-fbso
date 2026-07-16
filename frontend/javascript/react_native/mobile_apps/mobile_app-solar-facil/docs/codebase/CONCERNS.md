---
title: "Preocupações & Dívida Técnica — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["concerns", "tech-debt", "risks", "mobile"]
---

# Preocupações & Dívida Técnica — Solar Fácil

## 1. Dívida Técnica

### 1.1. Alta Prioridade

| ID | Item | Impacto | Localização |
|---|---|---|---|
| DT-001 | Senhas armazenadas em texto plano no SQLite | Segurança — dados sensíveis expostos se o dispositivo for comprometido | `initializeSQLiteDatabase.ts`, `useAssociados.ts` |
| DT-002 | Sem cobertura de testes (unitários, integração, E2E) | Qualidade — regressões não detectadas, refatoração arriscada | Projeto inteiro |
| DT-003 | README-ARQUITETURA.md menciona Redux Toolkit, mas código não usa | Documentação — confusão para novos devs | `README-ARQUITETURA.md` |
| DT-004 | Autenticação persiste apenas em memória (perdida ao fechar o app) | UX — usuário precisa fazer login a cada abertura | `AuthContext.tsx` |

### 1.2. Média Prioridade

| ID | Item | Impacto | Localização |
|---|---|---|---|
| DT-005 | Duplicação de CRUD: versões SQLite e AsyncStorage para as mesmas entidades | Manutenção — duas fontes de verdade, inconsistência potencial | `services/database/` vs `services/storage/` |
| DT-006 | Múltiplos arquivos de backup/cópia (.txt, Copy.ts, Copy2.ts) no source | Manutenção — poluição do diretório, confusão sobre qual é o canonical | `services/database/useAssociadosCopilot*.ts` |
| DT-007 | Hook `useAssociados` abre e fecha conexão a cada operação | Performance — overhead de abrir/fechar banco repetidamente | `useAssociados.ts` |
| DT-008 | Mock services simulam latência com setTimeout, mas não têm tratamento de erro | Resiliência — não testa cenários de falha | `services/mock/` |

### 1.3. Baixa Prioridade

| ID | Item | Impacto |
|---|---|---|
| DT-009 | Código comentado extenso em `_layout.tsx` e `tailwind.config.js` | Legibilidade |
| DT-010 | Ícones comentados e opções de cor inline nos estilos | Legibilidade |
| DT-011 | `any` usado em `inputBorder` e `api.sendBeneficiarioData` | Type safety |
| DT-012 | Sem Error Boundary global | Resiliência |

## 2. Bugs Conhecidos

| ID | Descrição | Severidade |
|---|---|---|
| [TODO] | Não foram documentados bugs conhecidos até o momento | — |

## 3. Riscos de Segurança (OWASP Mobile Top 10)

| Categoria OWASP | Status | Ação Recomendada |
|---|---|---|
| **M1: Improper Credential Usage** | ⚠️ Senhas em plain text | Implementar bcrypt/scrypt + salt |
| **M2: Inadequate Supply Chain Security** | ⚠️ Dependências não auditadas | Rodar `npm audit` regularmente; considerar `socket.dev` |
| **M3: Insecure Authentication** | ⚠️ Sem expiração de sessão | Implementar timeout de sessão |
| **M4: Insufficient Input/Output Validation** | ✅ yup + react-hook-form | Manter validação consistente |
| **M5: Insecure Communication** | ✅ Offline-first (sem rede) | Adicionar SSL pinning quando houver backend |
| **M6: Inadequate Privacy Controls** | ⚠️ Sem política de privacidade clara no app | Adicionar tela de consentimento |
| **M7: Insufficient Binary Protections** | ⚠️ Sem ofuscação de código | Configurar ProGuard (Android) + ofuscação |
| **M8: Security Misconfiguration** | ⚠️ `newArchEnabled: true` sem revisão | Auditar configurações de build |
| **M9: Insecure Data Storage** | ❌ Senhas e dados pessoais em SQLite plain text | Criptografar dados sensíveis |
| **M10: Insufficient Cryptography** | ❌ Sem criptografia de dados em repouso | Usar `expo-secure-store` + criptografia SQLite |

## 4. Performance

### 4.1. Observações

- **Sem virtualização**: [TODO] verificar se listas de associados e movimentações usam `FlatList` com virtualização
- **Imagens**: [TODO] verificar otimização de imagens (tamanho, cache)
- **Bundle size**: [TODO] analisar com `react-native-bundle-visualizer`
- **Animações**: usa Reanimated 3 (UI thread) — boa prática, mas [TODO] verificar uso correto de `useSharedValue`/`useAnimatedStyle`

### 4.2. Oportunidades de Otimização

- Migrar de `ScrollView` para `FlatList`/`FlashList` em listas
- Adicionar `React.memo()` em componentes de lista
- Usar `expo-image` em vez de `Image` (já incluído como dependência)
- Implementar lazy loading de tabs não-visíveis

## 5. Dependências Desatualizadas

| Biblioteca | Versão Atual | [TODO] Verificar Latest |
|---|---|---|
| react | 19.0.0 | [ASK USER] |
| react-native | 0.79.3 | [ASK USER] |
| expo | ~53.0.9 | [ASK USER] |
| @types/react-native | ^0.72.8 | ⚠️ Muito desatualizado vs RN 0.79 |

## 6. Dívida de Documentação

| Item | Status |
|---|---|
| README-ARQUITETURA.md desatualizado (menciona Redux) | ❌ |
| Sem documentação de API (não existe backend) | N/A |
| Sem ADRs para decisões arquiteturais | ❌ |
| Sem glossário de domínio | ❌ |
| Sem documentação de design system | ❌ |

## 7. Arquivos Órfãos (não utilizados)

| Arquivo | Suspeita |
|---|---|
| `screens/associado/AssociadoLoginScreenStorage.tsx.txt` | Backup renomeado para .txt |
| `services/database/initializeDatabase.ts.txt` | Backup renomeado para .txt |
| `services/database/useAssociadosCopilot.ts` | Cópia de desenvolvimento |
| `services/database/useAssociadosCopilotCopy.ts` | Cópia de desenvolvimento |
| `services/database/useAssociadosCopilotCopy2.ts` | Cópia de desenvolvimento |
| `services/database/useAssociadosCopilotGPT.ts` | Cópia de desenvolvimento |
| `services/database/useMovimentacoesCopilot.ts` | Cópia de desenvolvimento |
| `services/database/useMovimentacoesCopilotCopy.ts` | Cópia de desenvolvimento |
| `services/database/initializeSQLiteDatabaseCopy.ts` | Cópia de desenvolvimento |
| `services/database/seedMovimentacoesCopy.ts` | Cópia de desenvolvimento |
| `services/database/initializeDatabaseCopilot.ts` | Cópia de desenvolvimento |

**Total: ~11 arquivos órfãos** no diretório `services/database/` e `screens/associado/`.

## 8. Checklist de Ações Recomendadas

- [ ] Remover arquivos órfãos (`.txt`, `*Copy*.ts`, `*Copilot*.ts`)
- [ ] Unificar estratégia de persistência (escolher SQLite ou AsyncStorage, não ambos)
- [ ] Implementar hash de senhas (bcrypt ou scrypt)
- [ ] Configurar Jest + RNTL
- [ ] Atualizar `@types/react-native` para versão compatível com RN 0.79
- [ ] Configurar Error Boundary global
- [ ] Implementar persistência de sessão de autenticação
- [ ] Criar ADRs para decisões arquiteturais (já documentados)
- [ ] Corrigir README-ARQUITETURA.md (remover menção ao Redux)
- [ ] Rodar `npm audit fix`
