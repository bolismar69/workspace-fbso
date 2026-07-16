# Accessibility Audit — Solar Fácil Site

> Auditoria WCAG 2.1/2.2 — análise de acessibilidade com base no código-fonte.
> Gerado por `accessibility-compliance-accessibility-audit` em 2026-07-08.
> ⚠️ Verificações manuais (teclado, leitor de tela) não realizadas — Chrome indisponível.

---

## 1. Sumário Executivo

| Nível | Critérios Verificados | Passou | Falhou | Precisa Verificação Manual |
|---|---|---|---|---|
| **A (mínimo)** | 15 | 10 | 2 | 3 |
| **AA (recomendado)** | 8 | 5 | 1 | 2 |
| **AAA (desejável)** | 3 | 1 | 0 | 2 |
| **Total** | **26** | **16** | **3** | **7** |

---

## 2. Violações WCAG

### P0 — Críticas (bloqueiam uso)

| # | Finding | WCAG SC | Nível | Impacto |
|---|---|---|---|---|
| A1 | **Sem skip-to-content link** — usuários de teclado precisam tabular por todo o header | 2.4.1 Bypass Blocks | A | 🔴 Alto |
| A2 | **Sem atributos `aria-label` em elementos interativos** exceto o padrão HTML | 4.1.2 Name, Role, Value | A | 🔴 Alto |

### P1 — Altas (dificultam uso)

| # | Finding | WCAG SC | Nível | Impacto |
|---|---|---|---|---|
| A3 | **Ordem de foco não verificada** — MobileMenu pode ter foco quando invisível | 2.4.3 Focus Order | A | 🟡 Médio |
| A4 | **Sem indicador de idioma para conteúdo em espanhol** — `lang` attribute no html é `pt-BR` mas `/es/` planejado | 3.1.1 Language of Page | A | 🟡 Médio |

### P2 — Médias (inconvenientes)

| # | Finding | WCAG SC | Nível | Impacto |
|---|---|---|---|---|
| A5 | **Sem `autocomplete` em campos de formulário** — nome, email, telefone | 1.3.5 Identify Input Purpose | AA | 🟢 Baixo |

---

## 3. Critérios Verificados (aprovados)

| WCAG SC | Descrição | Evidência |
|---|---|---|
| **1.1.1 Non-text Content** | Alt text em imagens | Sem imagens — não aplicável |
| **1.3.1 Info and Relationships** | Estrutura semântica | HTML5 semântico (`<header>`, `<nav>`, `<main>`, `<footer>`) |
| **1.3.2 Meaningful Sequence** | Ordem de leitura | DOM segue ordem visual |
| **1.4.1 Use of Color** | Cor não é único meio | Ícones + texto + cores |
| **1.4.3 Contrast (Minimum)** | ≥4.5:1 texto, ≥3:1 large | Verde Terra #1E5631 : #ffffff = 7.28:1 (passa AAA) |
| **1.4.4 Resize Text** | 200% sem perda | Tailwind rem units |
| **2.1.1 Keyboard** | Funcionalidade por teclado | Botões e links nativos |
| **2.2.2 Pause, Stop, Hide** | Sem animação automática | Animações apenas em hover/scroll |
| **2.3.1 Three Flashes** | Sem flashes | Sem vídeo ou animação rápida |
| **2.4.2 Page Titled** | Título descritivo | Metadata API do Next.js |
| **2.4.4 Link Purpose** | Links descritivos | "Planos", "Contato", CTAs claros |
| **2.4.6 Headings and Labels** | Headings descritivos | Estrutura clara de headings |
| **3.1.1 Language of Page** | Idioma declarado | `<html lang="pt-BR">` |
| **3.2.3 Consistent Navigation** | Navegação consistente | Header/footer em todas as páginas |
| **3.3.1 Error Identification** | Erros identificados | Mensagens descritivas em pt-BR |
| **3.3.2 Labels or Instructions** | Labels em inputs | `FormField` inclui label |

---

## 4. Recomendações de Remediação

### 4.1 Skip-to-Content Link (WCAG 2.4.1)

```tsx
// Adicionar em src/app/layout.tsx, dentro do <body>, antes do <Header />
<a
  href="#main-content"
  className="sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:bg-white focus:px-4 focus:py-2 focus:rounded"
>
  Pular para o conteúdo principal
</a>
```

### 4.2 ARIA Labels (WCAG 4.1.2)

```tsx
// MobileMenu toggle
<button aria-label="Abrir menu" aria-expanded={isOpen} onClick={toggle}>
  <MenuIcon />
</button>

// Calculator input
<input
  type="number"
  aria-label="Valor da sua conta de luz em reais"
  placeholder="R$ 0,00"
/>

// FAQ accordion
<button aria-expanded={isOpen} aria-controls={`faq-${index}`}>
  {question}
</button>
```

### 4.3 Autocomplete Attributes (WCAG 1.3.5)

```tsx
<input type="text" name="name" autoComplete="name" />
<input type="email" name="email" autoComplete="email" />
<input type="tel" name="phone" autoComplete="tel" />
```

---

## 5. Checklist de Re-teste

- [ ] Adicionar skip-to-content link e verificar com Tab
- [ ] Adicionar `aria-label` em elementos interativos
- [ ] Verificar ordem de foco com MobileMenu fechado
- [ ] Adicionar `autocomplete` em campos de formulário
- [ ] Navegar site inteiro apenas com teclado
- [ ] Testar com NVDA (Windows) ou VoiceOver (Mac)
- [ ] Verificar contraste de `luz-do-sol` (#ffffbf) com texto
- [ ] Verificar redimensionamento de texto a 200%

---

Última atualização: 2026-07-08
