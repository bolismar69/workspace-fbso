---
title: "Avaliação Heurística de UX Mobile — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
method: "Análise estática de código + heurísticas de Nielsen adaptadas para mobile"
note: "Avaliação realizada apenas via análise de código — sem observação de usuários ou interação com o app"
---

# Avaliação Heurística de UX Mobile — Solar Fácil

## 1. Visibility of System Status — ⭐⭐⭐ (3/5)

**O usuário sabe o que está acontecendo?**

| Observação | Avaliação |
|---|---|
| StatusBar configurada com `style: "auto"` | ✅ Boa — adapta-se ao tema |
| React Query provê `isLoading`/`isFetching` | ✅ Boa — loading detectável |
| Sem indicador visual de sincronização/salvamento | ❌ Usuário não sabe se dados foram salvos |
| Sem feedback tátil em ações (Haptics instalado mas uso não detectado) | ⚠️ Oportunidade |
| Tab de Login mostra estado (Logout vs Login) dinamicamente | ✅ Boa |

**Recomendações:**
- Adicionar `Haptics.notificationAsync()` após salvar/excluir
- Adicionar Toast/SnackBar de confirmação ("Cadastro salvo com sucesso")
- Implementar indicador de salvamento automático

## 2. Match Between System and Real World — ⭐⭐⭐⭐ (4/5)

**A linguagem e os ícones são familiares ao usuário?**

| Observação | Avaliação |
|---|---|
| Ícones de tab intuitivos: solar-power, currency-exchange, person, book-open | ✅ Excelente |
| Terminologia de domínio correta: associado, beneficiado, fornecedor, movimentação | ✅ Consistente com o negócio |
| Ícone de login muda: flash-outline (deslogado) → flash-off-outline (logado) | ✅ Boa metáfora |
| Nomes de tabs: Solar, Plano, Saiba, FAQ, Login, User, Lista | ⚠️ "User" é genérico — poderia ser "Cadastro" |

## 3. User Control and Freedom — ⭐⭐ (2/5)

**O usuário pode desfazer ações e navegar livremente?**

| Observação | Avaliação |
|---|---|
| Navegação por 8 tabs — sempre visível e acessível | ✅ Boa |
| Sem botão "Voltar" ou gesto de voltar em fluxos de formulário | ❌ Usuário preso em formulários longos |
| Sem confirmação de logout (logout é imediato) | ❌ Logout acidental possível |
| Sem "Cancelar" em formulários de cadastro | ❌ Usuário não pode abortar cadastro |
| Sem gesto "swipe back" (iOS) em stacks | N/A (não há stacks aninhadas) |

**Recomendações:**
- Adicionar confirmação "Deseja sair?" no logout
- Adicionar botão Cancelar/Voltar em formulários
- Implementar salvamento de rascunho em formulários longos

## 4. Consistency and Standards — ⭐⭐⭐ (3/5)

**O app segue padrões de plataforma (iOS HIG, Material Design)?**

| Observação | Avaliação |
|---|---|
| SafeAreaView usado para respeitar áreas seguras | ✅ Bom |
| Edge-to-edge no Android habilitado | ✅ Bom |
| Componentes nativos para inputs (DateTimePicker, Picker) | ✅ Bom |
| Cores e estilos inconsistentes entre Theme e Tailwind | ⚠️ `secondary` = `#A4DE02` (Tailwind) vs `#A5C9CA` (Theme) |
| Estados de botão não padronizados (sem pressed/disabled) | ❌ Inconsistente |
| Formulários sem padrão de layout | ⚠️ Cada formulário tem layout próprio |

## 5. Error Prevention — ⭐⭐⭐ (3/5)

**O app previne erros antes que aconteçam?**

| Observação | Avaliação |
|---|---|
| Validação de formulários com yup (em tempo real via react-hook-form) | ✅ Excelente |
| CPF/CNPJ validados com algoritmos específicos | ✅ Bom |
| Máscaras de input (CPF, CNPJ, telefone) — `react-native-text-input-mask` | ✅ Bom |
| Sem confirmação para ações destrutivas (delete, logout) | ❌ Risco de ações acidentais |
| Campos obrigatórios provavelmente sinalizados | ⚠️ [RUNTIME] Verificar se labels indicam obrigatoriedade |

## 6. Recognition Rather than Recall — ⭐⭐⭐⭐ (4/5)

**O usuário reconhece opções em vez de lembrar?**

| Observação | Avaliação |
|---|---|
| Tabs sempre visíveis com ícones e labels | ✅ Excelente |
| Planos e FAQs visíveis sem necessidade de login | ✅ Boa descoberta |
| Dados mock carregados sem ação do usuário | ✅ Bom |
| Sem breadcrumbs ou indicador de localização | ⚠️ [RUNTIME] Avaliar se usuário se perde entre tabs |

## 7. Flexibility and Efficiency of Use — ⭐⭐ (2/5)

**O app oferece atalhos e eficiência para usuários experientes?**

| Observação | Avaliação |
|---|---|
| Sem atalhos de teclado/gesto | ❌ |
| Sem pull-to-refresh | ❌ |
| Sem swipe actions em listas (editar, excluir) | ❌ |
| Sem deep links para telas específicas | ⚠️ Scheme configurado, mas sem implementação |
| Formulário de cadastro é extenso (~30 campos) | ❌ Sem step-by-step ou progress indicator |

## 8. Aesthetic and Minimalist Design — ⭐⭐⭐ (3/5)

**A interface é limpa e focada?**

| Observação | Avaliação |
|---|---|
| Paleta de cores coerente (verdes + amarelo) | ✅ Tema de sustentabilidade |
| 8 tabs na barra inferior — potencialmente poluído | ⚠️ Muitas opções para um app de escopo médio |
| Informação densa nos formulários (30 campos em uma tela) | ❌ Sobrecarga cognitiva |
| Cabeçalho com fundo amarelo consistente | ✅ Identidade visual |

**Recomendações:**
- Consolidar tabs: "Saiba" + "FAQ" → "Ajuda"
- Dividir cadastro em steps (wizard)
- Esconder tabs avançadas para usuários não logados

## 9. Help Users Recognize, Diagnose, and Recover from Errors — ⭐⭐ (2/5)

**O app ajuda o usuário a resolver erros?**

| Observação | Avaliação |
|---|---|
| Mensagens de erro de validação (yup) — específicas e em português | ✅ Bom |
| Sem tratamento de erros de banco na UI | ❌ Erro silencioso para o usuário |
| Sem mensagem de "Nenhum resultado encontrado" em buscas | ❌ |
| Sem opção "Tentar novamente" em caso de falha | ❌ |

## 10. Help and Documentation — ⭐⭐ (2/5)

**O app oferece ajuda contextual?**

| Observação | Avaliação |
|---|---|
| Tela de FAQ com perguntas frequentes (mockadas) | ✅ Bom |
| Tela "Saiba Mais" com conteúdo institucional | ✅ Bom |
| Sem onboarding/tutorial para novos usuários | ❌ |
| Sem tooltips ou ajuda contextual em campos complexos | ❌ |
| Sem link para suporte/contato visível | ❌ |

---

## Sumário

| Heurística | Score |
|---|---|
| 1. System Status | ⭐⭐⭐ |
| 2. Real World Match | ⭐⭐⭐⭐ |
| 3. User Control | ⭐⭐ |
| 4. Consistency | ⭐⭐⭐ |
| 5. Error Prevention | ⭐⭐⭐ |
| 6. Recognition | ⭐⭐⭐⭐ |
| 7. Flexibility | ⭐⭐ |
| 8. Aesthetic | ⭐⭐⭐ |
| 9. Error Recovery | ⭐⭐ |
| 10. Help | ⭐⭐ |

**Média:** ⭐⭐⭐ (2.8/5) — Abaixo do recomendado para produção

## Top 5 Recomendações

1. **Adicionar onboarding** — Tutorial para novos usuários (3-4 telas com swipe)
2. **Consolidar navegação** — Reduzir de 8 para 5-6 tabs, esconder tabs restritas
3. **Wizard de cadastro** — Dividir formulário de 30 campos em 4-5 etapas
4. **Confirmações e feedback** — Toast de sucesso, confirmação de logout, confirmação de delete
5. **Error recovery** — Error boundary, retry automático, mensagens de erro amigáveis
