---
title: "Inventário — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["inventory", "governance", "reference"]
---

# Inventário do Projeto — TaxNexus Portal

## Sumário

| Métrica | Valor |
|---|---|
| Total de arquivos fonte | 5 (3 TSX + 1 TS + 1 CSS) |
| Linhas de código (aprox.) | ~350 |
| Componentes React | 2 (App, TaxSimulator) |
| Hooks customizados | 1 (useTaxService) |
| Testes | 0 |
| Cobertura de testes | 0% |

## Módulos

### `src/main.tsx`
- **Tipo:** Entry point
- **Linhas:** 10
- **Dependências:** React, ReactDOM, App, index.css
- **Cobertura:** 0%

### `src/App.tsx`
- **Tipo:** Shell component
- **Linhas:** 60
- **Estado:** `cnpj`, `authenticated`
- **Responsabilidades:** Pseudo-auth, layout condicional, delegação para TaxSimulator
- **Cobertura:** 0%

### `src/components/TaxSimulator.tsx`
- **Tipo:** Feature component
- **Linhas:** 161
- **Estado:** `selectedState`, `selectedCity`, `ncm`, `saldo`, `chartData`, `rawResponse`
- **Dependências:** Recharts (BarChart, ResponsiveContainer, etc.), useTaxService
- **Responsabilidades:** Formulário, chamada API, renderização condicional, gráfico
- **Cobertura:** 0%

### `src/hooks/useTaxService.ts`
- **Tipo:** Custom hook
- **Linhas:** 72
- **Estado:** `loading`
- **API chamada:** POST `http://localhost:8080/v1/tax/calculate`
- **Cobertura:** 0%

### `src/index.css`
- **Tipo:** Estilos globais
- **Linhas:** 112
- **Tokens:** Design tokens CSS (light/dark mode)
- **Nota:** Contém estilos residuais do template Vite

### `src/App.css`
- **Tipo:** Estilos (órfão)
- **Linhas:** 185
- **Status:** Não importado por App.tsx atual — residual do template Vite

## Assets

| Arquivo | Tipo | Usado? |
|---|---|---|
| `public/favicon.svg` | Favicon | Sim |
| `public/icons.svg` | SVG sprites | Não (template residual) |
| `src/assets/hero.png` | Imagem | Não (template residual) |
| `src/assets/react.svg` | SVG | Não (template residual) |
| `src/assets/vite.svg` | SVG | Não (template residual) |

## Configuração

| Arquivo | Propósito |
|---|---|
| `package.json` | Dependências e scripts (React 19, Vite 8, Recharts 3) |
| `vite.config.ts` | Config Vite + plugin React |
| `tsconfig.json` | TypeScript base config |
| `tsconfig.app.json` | TS config para código fonte |
| `tsconfig.node.json` | TS config para tooling |
| `eslint.config.js` | ESLint flat config |
| `Dockerfile` | Build multi-stage (Node 18 → Nginx) |
| `nginx.conf` | Nginx SPA config (porta 5173, try_files) |
| `index.html` | HTML shell |

## Dependências de Produção

| Pacote | Versão | Usado? |
|---|---|---|
| react | ^19.2.4 | ✅ |
| react-dom | ^19.2.4 | ✅ |
| recharts | ^3.8.0 | ✅ |
| lucide-react | ^0.577.0 | ❌ (não utilizado) |

## Arquivos para Remoção

| Arquivo | Motivo |
|---|---|
| `src/AppCopy.tsx.txt` | Template Vite original — mantido como referência mas não compila |
| `src/App.css` | Não importado pelo App.tsx atual |
| `src/assets/hero.png` | Template Vite |
| `src/assets/react.svg` | Template Vite |
| `src/assets/vite.svg` | Template Vite |

## Diretórios Vazios (Scaffolding)

| Diretório | Propósito planejado |
|---|---|
| `src/api/` | Clientes HTTP |
| `src/services/` | Lógica de negócio |
| `src/store/` | Estado global (Zustand/Context) |
| `src/pages/` | Páginas (React Router) |

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
