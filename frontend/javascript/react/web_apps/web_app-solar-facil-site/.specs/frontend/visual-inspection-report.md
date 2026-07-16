# Visual Inspection Report — Solar Fácil Site

> Inspeção visual com screenshots em 4 viewports. Detecta problemas de layout, responsividade e consistência visual.
> Gerado por `web-design-reviewer` em 2026-07-08.

---

## ⚠️ Limitação

**A inspeção visual com Playwright requer o navegador Chrome instalado**, indisponível neste ambiente (falta de permissão `sudo` para `npx playwright install chrome`).

Este relatório contém:
- Checklist de inspeção manual para executar quando o Chrome estiver disponível
- Análise estática de possíveis problemas com base no código Tailwind
- Procedimento de inspeção documentado

---

## 1. Procedimento de Inspeção (a executar)

### Viewports Alvo

| Nome | Largura | Dispositivo |
|---|---|---|
| Mobile | 375px | iPhone SE |
| Tablet | 768px | iPad |
| Desktop | 1280px | Standard PC |
| Wide | 1920px | Large display |

### Checklist por Viewport

```markdown
[ ] Header: logo visível, nav links ou hamburger menu
[ ] Hero: headline sem overflow, CTAs clicáveis (min 44px touch target)
[ ] Calculadora: inputs alinhados, botão "Calcular" visível
[ ] Resultado: valor de economia legível, plano sugerido destacado
[ ] Planos: 3 cards lado a lado (desktop) ou empilhados (mobile)
[ ] FAQ: accordion expande sem quebrar layout
[ ] Formulário: campos alinhados, labels acima dos inputs
[ ] Footer: links organizados, sem overflow horizontal
```

---

## 2. Análise Estática de Layout (código)

### 2.1 Breakpoints Configurados

Tailwind v4 com breakpoints padrão:
- `sm: 640px` → 1 coluna
- `md: 768px` → 2 colunas
- `lg: 1024px` → grid completo
- `xl: 1280px` → largura máxima

### 2.2 Possíveis Problemas de Layout

| Problema Potencial | Localização | Causa Provável |
|---|---|---|
| Header fixo sobrepondo conteúdo | `Header.tsx` | `position: fixed, height: 64px` — verificar `pt-16` no main |
| Cards de plano em grid | `PlansSection.tsx` | 3 cards em `md:grid-cols-3` — verificar overflow em 768px |
| Calculadora em mobile | `ConsumerCalculator.tsx` | Input + button em linha — precisam empilhar em <640px |
| Formulário mobile | `ContactForm.tsx` | Labels + inputs — precisam ser full-width |
| Tabela comparativa mobile | `PlansComparisonTable.tsx` | Tabela horizontal pode precisar de scroll ou card view |

### 2.3 Estados a Inspecionar

| Estado | Página | O que verificar |
|---|---|---|
| Loading | Homepage (planos) | Skeleton placeholder aparece antes dos cards |
| Empty | Planos | Mensagem "Nenhum plano" se array vazio |
| Error | Calculadora | Input vazio → erro "Informe um valor" |
| Success | Contato | SuccessScreen após submit bem-sucedido |
| Outlier | Calculadora | Valor <50 → mensagem de outlier |
| Mobile menu | Header | Menu toggle → links visíveis |

---

## 3. Recomendações para Inspeção Visual

1. Instalar Chrome: `npx playwright install chrome` (requer sudo)
2. Rodar: `npx playwright test e2e/visual-inspection.spec.ts`
3. Capturar screenshots nos 4 viewports e 5 estados
4. Comparar com o design system documentado em `design/DESIGN.md`

---

Última atualização: 2026-07-08
