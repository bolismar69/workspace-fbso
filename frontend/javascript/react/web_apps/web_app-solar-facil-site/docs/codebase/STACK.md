# STACK — Solar Fácil Site

> Stack tecnológico completo da aplicação frontend.
> Gerado por `acquire-codebase-knowledge` em 2026-07-08.
> Fontes: `package.json`, `tsconfig.json`, `next.config.ts`, `postcss.config.mjs`, `eslint.config.mjs`.

---

## 1. Runtime & Linguagem

| Componente | Tecnologia | Versão | Evidência |
|---|---|---|---|
| **Runtime** | Node.js | ≥20 (target ES2017) | `package.json:engines` (implícito), `tsconfig.json:3` |
| **Linguagem** | TypeScript | ^5 (strict mode) | `package.json:27`, `tsconfig.json` |
| **JS Target** | ES2017 | — | `tsconfig.json:3` |

## 2. Framework & Bibliotecas Core

| Componente | Tecnologia | Versão | Propósito |
|---|---|---|---|
| **Framework** | Next.js (App Router) | 16.2.10 | SSR, roteamento baseado em arquivos, server components |
| **UI Library** | React | 19.2.4 | Biblioteca de componentes declarativos |
| **DOM** | React DOM | 19.2.4 | Renderização no navegador |

## 3. Estilização

| Componente | Tecnologia | Versão | Propósito |
|---|---|---|---|
| **Framework CSS** | Tailwind CSS | ^4 | Utility-first CSS |
| **PostCSS Plugin** | `@tailwindcss/postcss` | ^4 | Integração Tailwind v4 com PostCSS |
| **Design Tokens** | CSS Custom Properties | — | Variáveis `--color-solar-*` em `globals.css` |
| **Ícones** | Lucide React | ^1.23.0 | Biblioteca de ícones SVG |

## 4. Tooling de Desenvolvimento

| Ferramenta | Versão | Propósito |
|---|---|---|
| **ESLint** | ^9 | Linting estático |
| **eslint-config-next** | 16.2.10 | Regras ESLint específicas do Next.js |
| **Prettier** | ^3.9.4 | Formatação de código |
| **prettier-plugin-tailwindcss** | ^0.8.0 | Ordenação automática de classes Tailwind |
| **TypeScript** | ^5 | Type-checking |

## 5. Dependências de Runtime (npm)

| Pacote | Versão | Tipo | Propósito |
|---|---|---|---|
| `next` | 16.2.10 | dependency | Framework full-stack React |
| `react` | 19.2.4 | dependency | Biblioteca UI |
| `react-dom` | 19.2.4 | dependency | Renderizador DOM |
| `lucide-react` | ^1.23.0 | dependency | Ícones SVG |

## 6. Dependências de Desenvolvimento (npm)

| Pacote | Versão | Propósito |
|---|---|---|
| `typescript` | ^5 | Compilador/type-checker |
| `@types/react` | ^19 | Tipos do React |
| `@types/react-dom` | ^19 | Tipos do React DOM |
| `@types/node` | ^20 | Tipos do Node.js |
| `eslint` | ^9 | Linter |
| `eslint-config-next` | 16.2.10 | Config ESLint Next.js |
| `prettier` | ^3.9.4 | Formatador |
| `prettier-plugin-tailwindcss` | ^0.8.0 | Plugin Prettier para Tailwind |
| `tailwindcss` | ^4 | Framework CSS |
| `@tailwindcss/postcss` | ^4 | Plugin PostCSS Tailwind |

## 7. Infraestrutura & Deploy

| Componente | Configuração | Evidência |
|---|---|---|
| **Output Mode** | `standalone` | `next.config.ts:5` |
| **Hospedagem Alvo** | Não-Vercel (DigitalOcean ou similar) | `output: 'standalone'` |
| **CI/CD** | GitHub Actions | `.github/workflows/deploy.yml` |
| **Imagens** | AVIF + WebP | `next.config.ts:9` |

## 8. O que NÃO está no stack

- ❌ **Nenhum gerenciador de estado** (Zustand, Redux, Context API) — usa `useState`/`useCallback` locais
- ❌ **Nenhum cliente HTTP** (axios, ky) — usa `fetch` nativo
- ❌ **Nenhuma biblioteca de formulários** (React Hook Form, Formik) — validação manual
- ❌ **Nenhum framework de teste** (Jest, Vitest, Testing Library) — zero testes automatizados
- ❌ **Nenhuma biblioteca de componentes UI** (Radix, Headless UI) — componentes 100% customizados
- ❌ **Nenhum backend real** — todos os dados são estáticos (constantes + JSON mocks)
- ❌ **Nenhum sistema de tema** (next-themes) — dark mode planejado mas não implementado

---

## 9. Versões e Compatibilidade

| Restrição | Detalhe |
|---|---|
| **Node.js mínimo** | ≥20 (implícito pelo Next.js 16) |
| **React** | 19.2.4 (React 19 — Server Components, Actions) |
| **Next.js** | 16.2.10 (App Router, Server Components por padrão) |
| **Tailwind** | v4 (CSS-first configuration, sem `tailwind.config.ts`) |
| **TypeScript** | strict mode (`tsconfig.json`) |

---

Última atualização: 2026-07-08
