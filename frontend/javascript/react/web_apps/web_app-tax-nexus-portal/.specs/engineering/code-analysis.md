---
title: "Análise de Código — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["code-analysis", "tutorial", "flow"]
---

# Análise Técnica de Fluxo de Código — TaxNexus Portal

## 1. Entry Point: `src/main.tsx`

**Tipo:** Tutorial (Diátaxis) — "Como a aplicação inicia"

```typescript
// src/main.tsx
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
```

**Fluxo:**
1. `createRoot` obtém o elemento `#root` do `index.html`
2. Renderiza `<App />` dentro de `<StrictMode>` (detecta side-effects indesejados em desenvolvimento)
3. `index.css` é carregado como estilo global (design tokens CSS com light/dark mode)

**Observações:**
- Sem providers (Router, Store, Theme) — arquitetura mínima para MVP
- `!` non-null assertion no `getElementById` — TypeScript assume que `#root` existe

---

## 2. Shell da Aplicação: `src/App.tsx`

**Tipo:** Tutorial (Diátaxis) — "Como a autenticação e o layout funcionam"

### Estrutura

```
App.tsx
├── Estado: cnpj (string), authenticated (boolean)
├── Tela de Login (quando !authenticated)
│   ├── Form com input CNPJ
│   └── Validação: cnpj.length >= 14 → setAuthenticated(true)
└── Tela Principal (quando authenticated)
    ├── Header: "TaxNexus Simulator v1.0" + Botão Sair
    └── <TaxSimulator cnpj={cnpj} />
```

### Pseudo-auth por CNPJ

```typescript
const handleLogin = (e: React.FormEvent) => {
  e.preventDefault();
  if (cnpj.length >= 14) {
    setAuthenticated(true);
  } else {
    alert("Por favor, insira um CNPJ válido.");
  }
};
```

**Observações de segurança:**
- Validação puramente client-side — bypass trivial
- Sem chamada ao backend para verificar existência do CNPJ
- `alert()` como único feedback de validação — UX rudimentar
- Nenhum token/sessão persistido — refresh da página perde o estado

---

## 3. Componente de Simulação: `src/components/TaxSimulator.tsx`

**Tipo:** Tutorial + Reference (Diátaxis)

### Props

```typescript
interface TaxSimulatorProps {
  cnpj: string;  // Recebido do App.tsx após pseudo-auth
}
```

### Estado Local

| Estado | Tipo | Default | Propósito |
|---|---|---|---|
| `selectedState` | string | `''` | UF selecionada (ex: '35' = SP) |
| `selectedCity` | string | `''` | Código IBGE do município |
| `ncm` | string | `'62011100'` | Código NCM |
| `saldo` | number | `0` | Saldo remanescente de créditos |
| `chartData` | any[] \| null | `null` | Dados formatados para Recharts |
| `rawResponse` | any \| null | `null` | Resposta bruta da API |

### Dados Hardcoded

**Estados disponíveis:** SP (35), RJ (33), AM (13)

**Cidades disponíveis:** Apenas capitais — São Paulo (3550308), Rio de Janeiro (3304557), Manaus (1302603)

### Fluxo de Simulação

```
handleSimulate()
  ├── 1. Chama calculateTax({ cnpj, ncm, ibge: selectedCity, saldo_remanescente: saldo })
  ├── 2. Aguarda resposta do backend
  ├── 3. Se sucesso:
  │     ├── rawResponse = data.calculation
  │     └── chartData = [
  │           { name: '2026 (Transição)', cbs, ibs },
  │           { name: '2027 (Projeção)', cbs: cbs*1.02, ibs: ibs*1.02 }
  │         ]
  └── 4. Renderização condicional:
        ├── Cards: Sistema Atual (PIS, COFINS, ICMS, IPI) vs. Pós-Reforma (CBS, IBS, IS, IPVA+ITCMD)
        └── Gráfico: BarChart Recharts (barras empilhadas CBS + IBS)
```

**Nota importante:** O valor de 2027 é uma **projeção simplificada** (2% de aumento sobre 2026). Não reflete o cálculo real da regra de transição completa (que envolveria saldo remanescente).

### Formato de Moeda

Valores formatados como `pt-BR` BRL:
```typescript
value.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })
```

### Gráfico Recharts

- **Tipo:** `BarChart` com barras empilhadas (`stackId="a"`)
- **Eixo X:** Ano (2026 Transição, 2027 Projeção)
- **Eixo Y:** Valor em R$
- **Tooltip:** Formatado como moeda BRL
- **Cores:** CBS = `#2563eb` (azul), IBS = `#059669` (verde)

---

## 4. Hook de API: `src/hooks/useTaxService.ts`

**Tipo:** How-to (Diátaxis) — "Como consumir a API de cálculo"

### Interface pública

```typescript
const { calculateTax, loading } = useTaxService();
```

### Implementação

```typescript
const calculateTax = async (payload: TaxRequest): Promise<TaxResponse | null> => {
  setLoading(true);
  try {
    const response = await fetch('http://localhost:8080/v1/tax/calculate', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
      body: JSON.stringify(payload),
    });
    if (!response.ok) {
      const errorData = await response.json();
      console.error("Erro do servidor:", errorData);
      return null;
    }
    return await response.json();
  } catch (error) {
    console.error("Erro na chamada Fetch:", error);
    return null;
  } finally {
    setLoading(false);
  }
};
```

**Padrão:** `useState` para loading + `try/catch/finally` com `fetch` nativo.

**Limitações identificadas:**
- URL hardcoded (`localhost:8080`)
- Sem `AbortController` — não cancela requisições em unmount
- `errorData` lido como JSON mas pode não ser JSON (ex: erro de rede antes do response)
- Sem tipagem para `errorData`

---

## 5. Design Tokens: `src/index.css`

**Tipo:** Reference (Diátaxis)

O arquivo define um sistema de design tokens via CSS custom properties:

| Token | Light | Dark |
|---|---|---|
| `--text` | `#6b6375` | `#9ca3af` |
| `--text-h` | `#08060d` | `#f3f4f6` |
| `--bg` | `#fff` | `#16171d` |
| `--border` | `#e5e4e7` | `#2e303a` |
| `--accent` | `#aa3bff` | `#c084fc` |

**Nota:** O `App.tsx` atual usa classes Tailwind (`bg-gray-100`, `text-blue-900`, etc.) — os design tokens do `index.css` são remanescentes do template Vite original. Tailwind **não** está listado como dependência no `package.json`, sugerindo que o CSS é servido via CDN ou os tokens nativos estão em transição para Tailwind.

---

## 6. Arquivos Órfãos/Residuais

| Arquivo | Status | Descrição |
|---|---|---|
| `src/AppCopy.tsx.txt` | Template original | Versão inicial do Vite template (contador). Renomeado para `.txt` para não compilar. |
| `src/App.css` | Residual | Estilos do template Vite original. Não importado por `App.tsx` atual. |
| `src/assets/react.svg` | Residual | Logo React do template Vite |
| `src/assets/vite.svg` | Residual | Logo Vite do template Vite |
| `src/assets/hero.png` | Residual | Imagem hero do template Vite |

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
