# Feature Roadmap — Solar Fácil Site

> Roadmap de features + dívidas técnicas.
> Gerado por `documentation-writer` em 2026-07-08.

---

## 1. Features Concluídas ✅

| Feature | Entregue | RFs |
|---|---|---|
| Calculadora de Economia (Consumidor) | ✅ | RF-CAL-001..005,009,010 |
| Calculadora de Ganho (Fornecedor) | ✅ | RF-CAL-006..008 |
| Formulário de Contato com validação | ✅ | RF-CON-001..011,013 |
| Página de Planos + FAQ | ✅ | RF-PLN-001..008 |
| Header + Footer + Mobile Menu | ✅ | RF-NAV-001..005 |
| SEO (robots.txt, sitemap, metadata) | ✅ | RF-SEO-001..004 |
| Google Analytics 4 condicional | ✅ | RF-ACT-001..003 |
| Output standalone | ✅ | RNF-DEP-001 |
| Design System documentado | ✅ | — |

---

## 2. Features Planejadas (Curto Prazo)

| Feature | Prioridade | RFs |
|---|---|---|
| Validação server-side no formulário | 🔴 Alta | RF-CON-012 |
| Headers CSP + HSTS | 🔴 Alta | RNF-SEC-002,003 |
| Substituir placeholders (WhatsApp, App Store, Formspree) | 🔴 Alta | — |
| Conectar eventos GA4 pendentes (faq_open, lead_capture) | 🟡 Média | RF-ACT-004,005 |
| Remover delay simulado dos serviços | 🟡 Média | — |

---

## 3. Features Planejadas (Médio Prazo)

| Feature | Prioridade | RFs |
|---|---|---|
| API Backend real (substituir mocks) | 🟡 Média | — |
| Dark mode (next-themes) | 🟢 Baixa | — |
| Internacionalização (/es/) | 🟢 Baixa | RNF-I18N-002 |
| Testes automatizados (Vitest + Playwright) | 🔴 Alta | — |
| Skip-to-content + ARIA labels | 🟡 Média | — |

---

## 4. Dívidas Técnicas

| # | Dívida | Impacto | Esforço |
|---|---|---|---|
| D1 | Zero testes automatizados | 🔴 | 3-5 dias |
| D2 | Validação apenas client-side | 🔴 | 1-2 dias |
| D3 | Sem CSP/HSTS headers | 🟡 | 0.5 dia |
| D4 | Placeholders em produção | 🟡 | 0.5 dia |
| D5 | Duplicação de dados (constants vs mocks) | 🟢 | 1 dia |
| D6 | console.error em produção | 🟢 | 0.5 dia |
| D7 | Delay artificial 500ms | 🟢 | 0.5 dia |

---

Última atualização: 2026-07-08
