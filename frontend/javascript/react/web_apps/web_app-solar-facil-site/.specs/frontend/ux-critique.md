# UX Critique — Solar Fácil Site

> Avaliação heurística de UX (10 heurísticas de Nielsen).
> Gerado por `impeccable critique` (análise do código + PRODUCT.md) em 2026-07-08.

---

## 1. Visibility of System Status ⭐⭐⭐⭐ (4/5)

**O design sempre deve manter os usuários informados sobre o que está acontecendo.**

- ✅ **Skeleton loading** — placeholder durante carregamento de planos
- ✅ **Estados de submit** — `isSubmitting` desabilita botão e mostra feedback visual
- ✅ **Mensagens de erro** — em português claro, sem technical jargon
- ✅ **Success screen** — confirmação pós-envio do formulário
- ⚠️ **Calculadora** — sem indicador visual de "calculando..." (resultado é instantâneo porque é local)

## 2. Match Between System and Real World ⭐⭐⭐⭐⭐ (5/5)

**O sistema deve falar a língua do usuário.**

- ✅ **Português natural** — toda a interface em pt-BR, tom informal e acolhedor
- ✅ **Conceitos familiares** — "conta de luz", "desconto", "economia" — não jargão técnico
- ✅ **Valores em reais** — `formatBRL()` formata moeda brasileira corretamente
- ✅ **Metáfora da "Casa Conectada"** — cada seção é um "cômodo"
- ✅ **Personas reais** — Consumidor, Fornecedor, Cooperativa — papéis compreensíveis

## 3. User Control and Freedom ⭐⭐⭐⭐ (4/5)

**Usuários frequentemente cometem erros e precisam de uma "saída de emergência".**

- ✅ **Reset da calculadora** — botão para limpar e recalcular
- ✅ **Navegação clara** — Header fixo, breadcrumbs, links no footer
- ✅ **Sem fidelidade** — FAQ enfatiza "cancele a qualquer momento, sem multas"
- ⚠️ **Formulário** — sem botão "Limpar" explícito (usuário limpa campo a campo)

## 4. Consistency and Standards ⭐⭐⭐⭐⭐ (5/5)

**Os usuários não devem se perguntar se palavras ou ações diferentes significam a mesma coisa.**

- ✅ **Design system documentado** — cores, tipografia, componentes consistentes
- ✅ **Inter é a única fonte** — hierarquia por peso, não por família
- ✅ **Botões padronizados** — 3 variantes (primary, secondary, outline), comportamento consistente
- ✅ **Estados previsíveis** — hover, focus, disabled padronizados
- ✅ **Convenções de código** — TypeScript strict, ESLint, Prettier

## 5. Error Prevention ⭐⭐⭐ (3/5)

**Melhor que boas mensagens de erro é um design que previne o erro.**

- ✅ **Validação inline** — campos validados durante digitação
- ✅ **Honeypot + timer** — anti-spam no formulário
- ✅ **Placeholder com formato** — ajuda o usuário a preencher corretamente
- ⚠️ **Validação apenas client-side** — sem proteção server-side
- ⚠️ **Sem confirmação antes de submit** — usuário pode enviar sem revisar

## 6. Recognition Rather Than Recall ⭐⭐⭐⭐ (4/5)

**Minimize a carga de memória do usuário tornando objetos, ações e opções visíveis.**

- ✅ **Header fixo** — navegação sempre visível
- ✅ **JourneySummary** — contexto da calculadora visível no formulário
- ✅ **Planos visíveis na homepage** — não precisa navegar para `/planos`
- ✅ **FAQ inline** — respostas expandem no contexto, sem abrir nova página
- ⚠️ **Breadcrumb** — existe mas uso limitado (3 páginas)

## 7. Flexibility and Efficiency of Use ⭐⭐⭐ (3/5)

**Atalhos — ocultos para novatos — podem acelerar a interação para usuários experientes.**

- ✅ **CTA direto no Hero** — "Quero Economizar" sem scroll
- ✅ **Abas consumidor/fornecedor** — troca rápida entre modos da calculadora
- ⚠️ **Sem atalhos de teclado** — navegação apenas por Tab/click
- ⚠️ **Sem autocomplete** — campos não usam `autocomplete` attributes

## 8. Aesthetic and Minimalist Design ⭐⭐⭐⭐⭐ (5/5)

**Diálogos não devem conter informação irrelevante ou raramente necessária.**

- ✅ **Design principles claros** — 5 princípios documentados em PRODUCT.md
- ✅ **Anti-referências explícitas** — o que NÃO fazer está documentado
- ✅ **Regra "Uma Voz"** — ≤15% de cor primária na tela
- ✅ **Flat-by-Default** — sombras apenas como resposta a estado
- ✅ **3 dependências de runtime** — zero bibliotecas desnecessárias

## 9. Help Users Recognize, Diagnose, and Recover from Errors ⭐⭐⭐⭐ (4/5)

**Mensagens de erro devem ser expressas em linguagem simples, indicar o problema e sugerir solução.**

- ✅ **Mensagens em pt-BR** — sem inglês ou "portunhol"
- ✅ **Erros específicos** — "Nome deve ter no mínimo 2 caracteres", não "Erro de validação"
- ✅ **Sugestão de alternativa** — "Fale conosco pelo WhatsApp" quando formulário falha
- ✅ **Outlier com explicação** — calculadora explica POR QUE o valor não pode ser calculado
- ⚠️ **Erro de rede** — mensagem genérica, sem sugerir "tente novamente"

## 10. Help and Documentation ⭐⭐⭐ (3/5)

**Embora seja melhor que o sistema possa ser usado sem documentação, pode ser necessário fornecer ajuda.**

- ✅ **FAQ integrada** — 6 perguntas frequentes respondidas
- ✅ **Seção "Como Funciona"** — 3 passos visuais (Produtor → Cooperativa → Consumidor)
- ✅ **README.md** — instruções de dev padronizadas
- ⚠️ **Sem tooltips** — campos do formulário sem ajuda contextual
- ⚠️ **Sem página de ajuda dedicada** — FAQ apenas em `/planos`

---

## Resumo

| Heurística | Score |
|---|---|
| 1. Visibility of System Status | ⭐⭐⭐⭐ 4/5 |
| 2. Match Between System and Real World | ⭐⭐⭐⭐⭐ 5/5 |
| 3. User Control and Freedom | ⭐⭐⭐⭐ 4/5 |
| 4. Consistency and Standards | ⭐⭐⭐⭐⭐ 5/5 |
| 5. Error Prevention | ⭐⭐⭐ 3/5 |
| 6. Recognition Rather Than Recall | ⭐⭐⭐⭐ 4/5 |
| 7. Flexibility and Efficiency of Use | ⭐⭐⭐ 3/5 |
| 8. Aesthetic and Minimalist Design | ⭐⭐⭐⭐⭐ 5/5 |
| 9. Help Recognize, Diagnose, Recover | ⭐⭐⭐⭐ 4/5 |
| 10. Help and Documentation | ⭐⭐⭐ 3/5 |
| **Média** | **⭐ 4.0/5** |

---

## Top 3 Recomendações

1. **Adicionar validação server-side** (Erro Prevention — #5)
2. **Adicionar tooltips/ ajuda contextual** nos formulários (Help — #10)
3. **Adicionar autocomplete e atalhos de teclado** (Efficiency — #7)

---

Última atualização: 2026-07-08
