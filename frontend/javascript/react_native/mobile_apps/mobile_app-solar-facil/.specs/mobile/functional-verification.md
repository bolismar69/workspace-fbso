---
title: "Verificação Funcional Mobile — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
status: "code-only"
note: "⚠️ Verificação funcional não executada — simulador/device indisponível. Fluxos documentados com base na análise de código."
---

# Verificação Funcional Mobile — Solar Fácil

## ⚠️ Status: Não Executada

Testes automatizados e manuais não executados por indisponibilidade de simulador/device e ausência de frameworks de teste configurados (Jest, RNTL, Detox, Maestro).

## Fluxos de Usuário Documentados (Análise de Código)

### Fluxo 1: Splash → Home
```
App abre → Splash Screen (estático, fundo branco, logotipo 200px)
  → RootLayout monta providers (ReactQuery → AppTheme → Auth → Database)
  → DatabaseProvider auto-inicializa SQLite
  → HomeScreen renderiza (tela pública)
```

**Componentes envolvidos:** `_layout.tsx`, `DatabaseContext.tsx`, `AuthContext.tsx`, `HomeScreen.tsx`

### Fluxo 2: Cadastro de Associado
```
Tab "User" → cadastro.tsx → FormCadastroAssociado
  → useFormValidation (yup schema)
  → Campos: nome, email, CPF/CNPJ, senha, tipoPessoa, tipoAssociado, endereço, etc.
  → Submit → useMutationAssociadoInsertRecord → SQLite
  → [TODO] Feedback de sucesso/erro
```

**Componentes envolvidos:** `cadastro.tsx`, `FormCadastroAssociado`, `useFormValidation`, `useMutationAssociadoInsertRecord`

**Validações esperadas:**
- CPF/CNPJ válido (formatado com máscara)
- Email em formato válido
- Senha obrigatória
- Campos condicionais: PF (dataNascimento) vs PJ (dataAbertura, razaoSocial)

### Fluxo 3: Login
```
Tab "Login" → login.tsx → AssociadoLoginScreen
  → Formulário: CPF/CNPJ + senha
  → useQueryAssociadosSearchByCpfCnpjSenha → SQLite
  → Se encontrado: AuthContext.login(userID, userName, associado)
  → Tab muda para "Logout" (ícone flash-off-outline)
  → Se não encontrado: [TODO] mensagem de erro
```

**Componentes envolvidos:** `login.tsx`, `AssociadoLoginScreen`, `AuthContext`

### Fluxo 4: Visualização de Movimentações
```
Tab "..." → movimentacao.tsx → MovimentacaoComCardScreen
  → useQueryMovimentacoesSearchByAssociadoId(associado.id) → SQLite
  → Cards de movimentação com:
      - Valor total, valor economizado, percentual economizado
      - Status de pagamento (Pago/Pendente) com cor de fundo
      - Data de vencimento e pagamento
  → Gráfico Victory com dados de economia
```

**Componentes envolvidos:** `movimentacao.tsx`, `MovimentacaoComCardScreen`, `useQueryMovimentacoesSearchByAssociadoId`

### Fluxo 5: Lista de Associados
```
Tab "Lista" → listatodos.tsx → AssociadoListaTodosScreen
  → useQueryAssociadosSearchAll → SQLite
  → Lista de AssociadoItem (nome, CPF/CNPJ, status)
  → [TODO] Ações: editar, excluir
```

### Fluxo 6: Planos Comerciais
```
Tab "Plano" → planos.tsx → PlanosScreen
  → servicePlans → mockPlans.json
  → Cards de CardPlan com nome, descrição, economia estimada
```

### Fluxo 7: FAQ
```
Tab "FAQ" → faq.tsx → FAQScreen
  → serviceFAQs → mockFAQs.json
  → FaqAccordion com pergunta (header) e resposta (expansível)
```

### Fluxo 8: Saiba Mais
```
Tab "Saiba" → saibamais.tsx → SaibaMaisScreen
  → Conteúdo institucional (imagens, texto)
```

### Fluxo 9: Logout
```
Tab "Logout" → toque → AuthContext.logout()
  → Estado limpo (isLoggedIn=false, userID=null, associado=undefined)
  → Tab volta a mostrar "Login"
```

## Cobertura de Cenários Alternativos

| Cenário | Status | Observação |
|---|---|---|
| **API offline** | N/A | App é offline-first — sem chamadas de rede |
| **Timeout** | N/A | Sem chamadas de rede |
| **Dados inválidos** | ✅ | yup valida antes de submeter |
| **CPF/CNPJ duplicado** | ✅ | SQLite UNIQUE constraint — [RUNTIME] Verificar mensagem de erro |
| **Lista vazia** | ❌ | Sem empty state |
| **Erro de banco** | ⚠️ | Console.error — não mostrado ao usuário |
| **Deep link** | ⚠️ | Scheme configurado — [RUNTIME] Verificar |

## Testes Automatizados

| Framework | Status |
|---|---|
| Jest | ❌ Não configurado |
| React Native Testing Library | ❌ Não configurado |
| Detox (E2E) | ❌ Não configurado |
| Maestro (E2E) | ❌ Não configurado |

## Checklist para Verificação Futura

- [ ] Configurar Jest + RNTL
- [ ] Escrever testes unitários para validadores (CPF, CNPJ, RG)
- [ ] Escrever testes de componente para inputs e formulários
- [ ] Escrever testes de integração para fluxos de cadastro e login
- [ ] Configurar Detox ou Maestro para E2E
- [ ] Executar verificação manual em iOS e Android
- [ ] Testar cenários de erro (dados inválidos, constraint violations)
- [ ] Testar deep links
