# Catálogo de Componentes — Solar Fácil Site

> Catálogo completo de 25 componentes React com props, estados e variantes.
> Gerado por `impeccable document` + análise do código-fonte em 2026-07-08.
> Fontes: `src/components/`, `.impeccable/design.json`.

---

## Shared Components (7)

### 1. Button
- **Arquivo**: `src/components/shared/Button.tsx`
- **Variantes**: `primary`, `secondary`, `outline`
- **Estados**: default, hover, focus-visible, disabled
- **Props**: `variant`, `href?`, `onClick?`, `disabled?`, `children`

### 2. SectionWrapper
- **Arquivo**: `src/components/shared/SectionWrapper.tsx`
- **Server Component**: Sim
- **Props**: `id?: string`, `background?: 'white' | 'yellow'`, `children`

### 3. MetricCard
- **Arquivo**: `src/components/shared/MetricCard.tsx`
- **Props**: `value: string`, `label: string`, `icon?: string`

### 4. Breadcrumb
- **Arquivo**: `src/components/shared/Breadcrumb.tsx`
- **Props**: `items: { label: string; href: string }[]`

### 5. Skeleton
- **Arquivo**: `src/components/shared/Skeleton.tsx`
- **Props**: Nenhuma (placeholder puro)

### 6. JsonLd
- **Arquivo**: `src/components/shared/JsonLd.tsx`
- **Server Component**: Sim
- **Props**: `type: string`, `data: Record<string, unknown>`

### 7. AnalyticsProvider
- **Arquivo**: `src/components/shared/AnalyticsProvider.tsx`
- **Client Component**: Sim
- **Props**: `measurementId?: string`

---

## Layout Components (3)

### 8. Header
- **Arquivo**: `src/components/layout/Header.tsx`
- **Server Component**: Sim (contém MobileMenu Client)
- **Contém**: Logo, Nav links, CTA button, MobileMenu toggle

### 9. Footer
- **Arquivo**: `src/components/layout/Footer.tsx`
- **Server Component**: Sim
- **Contém**: Logo, links, redes sociais, copyright

### 10. MobileMenu
- **Arquivo**: `src/components/layout/MobileMenu.tsx`
- **Client Component**: Sim
- **Estados**: open, closed, animating

---

## Homepage Components (9)

### 11. HeroSection
- **Arquivo**: `src/components/home/HeroSection.tsx`
- **Server Component**: Sim
- **Contém**: Headline, subtítulo, CTAs (consumidor + fornecedor)

### 12. CalculatorSection
- **Arquivo**: `src/components/home/CalculatorSection.tsx`
- **Server Component**: Sim (wrapper, estado nos filhos)
- **Contém**: Abas consumidor/fornecedor

### 13. ConsumerCalculator
- **Arquivo**: `src/components/home/ConsumerCalculator.tsx`
- **Client Component**: Sim
- **Estados**: default, input, loading (calculando), result, error, outlier
- **Hook**: `useCalculator('consumer')`

### 14. ProviderCalculator
- **Arquivo**: `src/components/home/ProviderCalculator.tsx`
- **Client Component**: Sim
- **Estados**: default, input, loading, result, error, outlier
- **Hook**: `useCalculator('provider')`

### 15. ProofSection
- **Arquivo**: `src/components/home/ProofSection.tsx`
- **Server Component**: Sim
- **Dados**: `METRICS` constant (3 métricas)

### 16. HowItWorksSection
- **Arquivo**: `src/components/home/HowItWorksSection.tsx`
- **Server Component**: Sim
- **Dados**: `HOW_IT_WORKS_STEPS` constant (3 passos)

### 17. PlansSection
- **Arquivo**: `src/components/home/PlansSection.tsx`
- **Client Component**: Sim
- **Estados**: loading (Skeleton), loaded (PlanCards), error, empty
- **Hook**: `usePlans()`

### 18. DifferentiatorsSection
- **Arquivo**: `src/components/home/DifferentiatorsSection.tsx`
- **Server Component**: Sim
- **Dados**: `DIFFERENTIATORS` constant (4 diferenciais)

### 19. FinalCtaSection
- **Arquivo**: `src/components/home/FinalCtaSection.tsx`
- **Server Component**: Sim
- **Contém**: Headline, subtítulo, CTAs

---

## Plans Page Components (4)

### 20. PlanCard
- **Arquivo**: `src/components/plans/PlanCard.tsx`
- **Props**: `plan: Plan`, `highlight?: boolean`
- **Estados**: default, hover (elevação)

### 21. PlansComparisonTable
- **Arquivo**: `src/components/plans/PlansComparisonTable.tsx`
- **Server Component**: Sim
- **Dados**: `PLANS` constant

### 22. ProviderHighlight
- **Arquivo**: `src/components/plans/ProviderHighlight.tsx`
- **Server Component**: Sim
- **Contém**: Benefícios para fornecedores, CTA

### 23. FaqAccordion
- **Arquivo**: `src/components/plans/FaqAccordion.tsx`
- **Client Component**: Sim
- **Estados**: collapsed (default), expanded
- **Hook**: `useFaqAccordion()`

---

## Contact Page Components (5)

### 24. ContactForm
- **Arquivo**: `src/components/contact/ContactForm.tsx`
- **Client Component**: Sim
- **Estados**: default, validating, submitting, success, error
- **Hook**: `useContactForm()`

### 25. FormField
- **Arquivo**: `src/components/contact/FormField.tsx`
- **Props**: `label`, `name`, `type`, `value`, `error?`, `required?`, `placeholder?`, `onChange`
- **Estados**: default, focus, error

### 26. JourneySummary
- **Arquivo**: `src/components/contact/JourneySummary.tsx`
- **Client Component**: Sim
- **Props**: `profile?`, `message?`

### 27. DirectChannels
- **Arquivo**: `src/components/contact/DirectChannels.tsx`
- **Client Component**: Sim
- **Contém**: WhatsApp, Email, Redes Sociais

### 28. SuccessScreen
- **Arquivo**: `src/components/contact/SuccessScreen.tsx`
- **Client Component**: Sim
- **Contém**: Mensagem de confirmação, próximos passos

---

## Hooks (7)

| Hook | Tipo | Dependências |
|---|---|---|
| `useCalculator` | State + Analytics | `calculator.ts`, `analytics.ts` |
| `useContactForm` | State + Fetch | `validation.ts`, `fetch` |
| `usePlans` | Data Fetching | `servicePlans.ts` |
| `useFAQs` | Data Fetching | `serviceFAQs.ts` |
| `useConcessionarias` | Data Fetching | `serviceConcessionarias.ts` |
| `useConsumoMedio` | Data Fetching | `serviceConsumoMedio.ts` |
| `useFaqAccordion` | UI State | — |

---

## Resumo

| Categoria | Count | Client | Server |
|---|---|---|---|
| Shared | 7 | 4 | 3 |
| Layout | 3 | 1 | 2 |
| Homepage | 9 | 3 | 6 |
| Plans | 4 | 2 | 2 |
| Contact | 5 | 5 | 0 |
| **Total** | **28** | **15** | **13** |

---

Última atualização: 2026-07-08
