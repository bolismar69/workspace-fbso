---
title: "Catálogo de Componentes — TaxNexus Portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
source: "src/ (App.tsx, TaxSimulator.tsx, useTaxService.ts)"
---

# Catálogo de Componentes React — TaxNexus Portal

O portal tem uma hierarquia simples de componentes: **App → TaxSimulator**, com o hook customizado `useTaxService` fornecendo a lógica de API.

---

## Estrutura de Arquivos

```
src/
├── main.tsx              ← Entry point (React 19 + Vite 8)
├── App.tsx               ← Container principal + lógica de autenticação
├── App.css               ← Estilos do template Vite (hero, counter, next-steps)
├── index.css             ← Design tokens globais + resets
├── components/
│   └── TaxSimulator.tsx  ← Formulário + resultados + gráfico
├── hooks/
│   └── useTaxService.ts  ← Hook de chamada à API backend
├── pages/                ← Diretório vazio (futuro)
├── services/             ← Diretório vazio (futuro)
├── store/                ← Diretório vazio (futuro)
├── api/                  ← Diretório vazio (futuro)
└── assets/               ← Diretório vazio (futuro)
```

---

## App (`App.tsx`)

**Tipo:** Container / Controller component
**Responsabilidade:** Gerenciar estado de autenticação e renderizar a view correta

### Props
Nenhuma (componente raiz).

### Estado Interno

| Estado | Tipo | Default | Descrição |
|---|---|---|---|
| `cnpj` | `string` | `''` | CNPJ digitado pelo usuário |
| `authenticated` | `boolean` | `false` | Se o CNPJ foi validado (length ≥ 14) |

### Variantes

| Variante | Condição | Renderização |
|---|---|---|
| **Login** | `!authenticated` | Card centralizado com formulário CNPJ |
| **Autenticado** | `authenticated` | Header + componente `<TaxSimulator>` |

### Eventos

| Evento | Handler | Comportamento |
|---|---|---|
| Submit CNPJ | `handleLogin` | Se `cnpj.length >= 14`, seta `authenticated = true`; senão, `alert("CNPJ inválido")` |
| Sair | `() => setAuthenticated(false)` | Volta à tela de login, limpa estado |

### Estrutura Visual

**View Login:**
```
<div.min-h-screen.bg-gray-100.py-10>
  <div.max-w-md.mx-auto.bg-white.p-8.rounded-lg.shadow-md>
    <h1>TaxNexus TaaS</h1>
    <form>
      <label>CNPJ do Contribuinte</label>
      <input type="text" placeholder="00000000000000" />
      <button type="submit">ACESSAR PORTAL</button>
    </form>
  </div>
</div>
```

**View Autenticado:**
```
<div.container.mx-auto>
  <header.flex.justify-between>
    <h1>TaxNexus Simulator v1.0</h1>
    <button>Sair / Trocar CNPJ</button>
  </header>
  <TaxSimulator cnpj={cnpj} />
</div>
```

---

## TaxSimulator (`components/TaxSimulator.tsx`)

**Tipo:** Feature component (formulário + resultados)
**Responsabilidade:** Coletar parâmetros, chamar API, exibir resultados comparativos

### Props

| Prop | Tipo | Obrigatório | Descrição |
|---|---|---|---|
| `cnpj` | `string` | Sim | CNPJ do contribuinte logado |

### Estado Interno

| Estado | Tipo | Default | Descrição |
|---|---|---|---|
| `selectedState` | `string` | `''` | Código IBGE da UF selecionada |
| `selectedCity` | `string` | `''` | Código IBGE da cidade selecionada |
| `ncm` | `string` | `'62011100'` | Código NCM do produto |
| `saldo` | `number` | `0` | Saldo remanescente de créditos |
| `chartData` | `any[] \| null` | `null` | Dados formatados para o gráfico |
| `rawResponse` | `any \| null` | `null` | Resposta bruta da API |

### Dados Estáticos

**Estados disponíveis:** SP (35), RJ (33), AM (13)
**Cidades:** São Paulo (3550308), Rio de Janeiro (3304557), Manaus (1302603)

### Subcomponentes

#### InfoBar (CNPJ)
```
<div.bg-gray-50.border-l-4.border-blue-500>
  Empresa Conectada: <span.font-bold>{cnpj}</span>
</div>
```

#### Select Estado
- `<select>` com options dos 3 estados
- Evento: `onChange` → `setSelectedState`

#### Select Cidade
- `<select>` filtrado por `selectedState`
- **Disabled** quando `!selectedState`
- Evento: `onChange` → `setSelectedCity`

#### Input NCM
- `<input type="text">` com default `'62011100'`
- Sem validação de formato

#### Input Saldo Remanescente
- `<input type="number">`
- Container com fundo `yellow-50` + borda `yellow-100`
- Sem validação de valor mínimo

#### Botão Simular
- Texto: `'Processando...'` (loading) ou `'SIMULAR REFORMA TRIBUTÁRIA'` (default)
- **Disabled** quando `loading || !selectedCity`
- Hover: `bg-blue-700`
- Disabled: `bg-gray-400`

#### Cards de Resultado
- **Renderização condicional:** `{rawResponse && (...)}`

**Card Legado (Sistema Atual):**
```
<div.bg-orange-50.border.border-orange-200.rounded-xl.p-5>
  <h3>Sistema Atual (Legado)</h3>
  <ul>
    <li>PIS: {rawResponse.pis}</li>
    <li>COFINS: {rawResponse.cofins}</li>
    <li>ICMS: {rawResponse.icms}</li>
    <li>IPI: {rawResponse.ipi}</li>
  </ul>
</div>
```

**Card Reforma (IVA Dual):**
```
<div.bg-blue-50.border.border-blue-200.rounded-xl.p-5>
  <h3>Pós-Reforma (IVA Dual)</h3>
  <ul>
    <li>CBS: {rawResponse.cbs_calculada}</li>
    <li>IBS: {rawResponse.ibs_calculado}</li>
    <li>Seletivo: {rawResponse.imposto_seletivo}</li>
    <li>IPVA/ITCMD Novo: {rawResponse.ipva_novo + rawResponse.itcmd_novo}</li>
  </ul>
</div>
```

#### Gráfico de Barras (Recharts)
```
<div.h-80.bg-white.rounded-lg.border.shadow-inner>
  <h3>Comparativo de Transição</h3>
  <ResponsiveContainer>
    <BarChart data={chartData}>
      <CartesianGrid strokeDasharray="3 3" vertical={false} />
      <XAxis dataKey="name" />
      <YAxis />
      <Tooltip formatter={BRL} />
      <Legend />
      <Bar dataKey="cbs" fill="#2563eb" stackId="a" />
      <Bar dataKey="ibs" fill="#059669" stackId="a" />
    </BarChart>
  </ResponsiveContainer>
</div>
```
- **Dados:** 2026 (Transição) e 2027 (Projeção com multiplicador ×1.02)
- **Cores:** CBS = `#2563eb` (azul), IBS = `#059669` (verde)
- **Formato:** Barras empilhadas (`stackId="a"`)

### Estados do Componente

| Estado | Disparo | UI |
|---|---|---|
| **Default** | Renderização inicial | Formulário vazio, botão disabled |
| **Pronto** | Cidade selecionada | Botão azul "SIMULAR" habilitado |
| **Loading** | Click em Simular | Botão cinza "Processando...", disabled |
| **Resultado** | API retorna dados | Cards laranja/azul + gráfico visíveis |
| **Erro (API)** | API falha | ❌ Não implementado — botão volta ao normal silenciosamente |
| **Empty (sem estado)** | Nenhuma UF selecionada | Select de cidade disabled |

---

## useTaxService (`hooks/useTaxService.ts`)

**Tipo:** Custom Hook
**Responsabilidade:** Encapsular chamada à API backend de cálculo tributário

### Interface

```typescript
interface TaxCalculationRequest {
  cnpj: string;
  ncm: string;
  ibge: string;
  saldo_remanescente: number;
}

interface UseTaxServiceReturn {
  calculateTax: (data: TaxCalculationRequest) => Promise<any>;
  loading: boolean;
}
```

### Comportamento
- **Endpoint:** `POST http://localhost:8080/v1/tax/calculate`
- **Headers:** `Content-Type: application/json`
- **Timeout:** ❌ Não implementado
- **Error handling:** ❌ Não implementado (sem try/catch)
- **Loading state:** ✅ Gerencia `loading` boolean durante a chamada

---

## Dependências de Componentes

```
App
├── useState (React)
├── TaxSimulator (./components/TaxSimulator)
│   ├── useState (React)
│   ├── useTaxService (../hooks/useTaxService)
│   │   └── fetch (Web API)
│   ├── BarChart (recharts)
│   ├── Bar (recharts)
│   ├── XAxis (recharts)
│   ├── YAxis (recharts)
│   ├── Tooltip (recharts)
│   ├── CartesianGrid (recharts)
│   ├── ResponsiveContainer (recharts)
│   └── Legend (recharts)
└── App.css, index.css (estilos globais)
```

---

## Estados Não Implementados

| Componente | Estado Faltante | Impacto |
|---|---|---|
| App | Error na validação CNPJ (inline, não alert) | P1 — experiência pobre |
| TaxSimulator | Erro na chamada API | P0 — falha silenciosa |
| TaxSimulator | Timeout da API | P1 — loading eterno |
| TaxSimulator | Validação inline dos campos | P1 — erros só no submit |
| TaxSimulator | Skeleton/placeholder durante load | P2 — layout shift |
| useTaxService | Retry policy | P2 — sem resiliência |

---

🤖 *Catálogo gerado por mineração de especificações frontend (PROMPT-MINING-FRONTEND-SPECIFICATION).*
