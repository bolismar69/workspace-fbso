---
title: "Feature Roadmap — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["roadmap", "features", "tech-debt"]
---

# Feature Roadmap — TaxNexus Portal

## Features Concluídas

| Feature | Status | Data | Descrição |
|---|---|---|---|
| FTR-001: Pseudo-auth CNPJ | ✅ Done | 2026-03 | Login simplificado com validação local de 14 dígitos |
| FTR-002: Formulário de simulação | ✅ Done | 2026-03 | Seletores encadeados UF→Cidade + NCM + Saldo |
| FTR-003: Integração API cálculo | ✅ Done | 2026-03 | POST /v1/tax/calculate com fetch nativo |
| FTR-004: Cards comparativos | ✅ Done | 2026-03 | Exibição lado a lado: sistema atual vs. IVA Dual |
| FTR-005: Gráfico de transição | ✅ Done | 2026-03 | BarChart Recharts com 2026 vs. 2027 |
| FTR-006: Dockerização | ✅ Done | 2026-03 | Build multi-stage Node→Nginx, porta 5173 |

## Features Planejadas

| Feature | Prioridade | Esforço | Descrição |
|---|---|---|---|
| FTR-007: Autenticação real | P1 | M | JWT/OAuth com refresh token |
| FTR-008: Configuração por ambiente | P1 | P | Variáveis VITE_* para URL da API |
| FTR-009: Tratamento de erros UX | P1 | M | Toasts de erro, mensagens amigáveis |
| FTR-010: Histórico de simulações | P2 | G | Lista de simulações anteriores por CNPJ |
| FTR-011: Múltiplos cenários | P2 | G | Comparação lado a lado de 2+ cenários |
| FTR-012: Exportação de relatório | P2 | M | PDF/CSV com resultado da simulação |
| FTR-013: Testes unitários | P2 | G | Vitest + React Testing Library |
| FTR-014: Router (React Router) | P3 | M | Navegação entre páginas (home, histórico, relatórios) |
| FTR-015: Estado global | P3 | M | Zustand ou Context API para estado compartilhado |
| FTR-016: Design responsivo | P3 | G | Layout mobile-friendly |
| FTR-017: Acessibilidade (a11y) | P3 | M | WCAG 2.1 AA |
| FTR-018: Monitoramento (RUM) | P3 | M | Sentry/Datadog para erros e performance |

## Dívidas Técnicas

| ID | Descrição | Severidade | Esforço | Feature Relacionada |
|---|---|---|---|---|
| TD-001 | URL da API hardcoded — `localhost:8080` fixo | 🔴 Alta | P | FTR-008 |
| TD-002 | Sem autenticação na chamada à API | 🔴 Alta | G | FTR-007 |
| TD-003 | Sem `AbortController` — requisições não canceladas em unmount | 🟡 Média | P | — |
| TD-004 | Pseudo-auth client-side (bypass trivial) | 🔴 Alta | G | FTR-007 |
| TD-005 | Sem timeout configurado no fetch | 🟡 Média | P | FTR-009 |
| TD-006 | Sem testes automatizados (0% coverage) | 🟡 Média | G | FTR-013 |
| TD-007 | `lucide-react` não utilizado — dependência órfã | 🟢 Baixa | P | — |
| TD-008 | Arquivos residuais do template Vite (AppCopy.tsx.txt, App.css, assets) | 🟢 Baixa | P | — |
| TD-009 | Tailwind usado em classes mas não listado em dependências (CDN?) | 🟡 Média | P | — |
| TD-010 | Estados e cidades hardcoded (apenas 3 UFs) | 🟡 Média | M | — |

## Marcos do Projeto

| Marco | Data Prevista | Features |
|---|---|---|
| 🚀 MVP v1.0 | 2026-03 (concluído) | FTR-001 a FTR-006 |
| 🔒 v1.1 — Segurança | 2026-Q3 | FTR-007, FTR-008, FTR-009 |
| 📊 v1.2 — Analytics | 2026-Q4 | FTR-010, FTR-011, FTR-012 |
| 🧪 v1.3 — Qualidade | 2027-Q1 | FTR-013, FTR-016, FTR-017 |
| 🏗️ v2.0 — Plataforma | 2027-Q2 | FTR-014, FTR-015, FTR-018 |

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
