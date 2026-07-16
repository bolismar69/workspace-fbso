---
title: "Auditoria de Acessibilidade WCAG — TaxNexus Portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["accessibility", "wcag", "a11y", "audit"]
audit_method: "static-code-analysis"
wcag_level_target: "AA"
browser_available: false
---

# Auditoria de Acessibilidade WCAG 2.1/2.2 — TaxNexus Portal

**Método:** Análise estática de código-fonte. Scans automatizados e verificação com leitor de tela não disponíveis (browser não pôde ser instalado neste ambiente). Issues reportados com base em violações identificáveis no código.

**Alvo WCAG:** Nível AA (conforme definido em PRODUCT.md).

---

## Sumário Executivo

| Severidade | Quantidade | Descrição |
|---|---|---|
| 🔴 Crítico | 3 | Violações que bloqueiam completamente usuários de tecnologia assistiva |
| 🟠 Alto | 4 | Barreiras significativas que dificultam o uso |
| 🟡 Médio | 4 | Issues que degradam a experiência mas têm workarounds |
| 🟢 Baixo | 2 | Melhorias recomendadas para excelência em a11y |
| **Total** | **13** | |

**Conformidade estimada:** ~40% dos critérios WCAG AA aplicáveis atendidos.

---

## Findings por Critério WCAG

### 1. Perceivable

#### 1.1.1 Non-text Content (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Gráfico Recharts sem alternativa textual | 🔴 Crítico | `TaxSimulator.tsx:131-153` | Usuários de leitor de tela não acessam os dados do gráfico. O `<BarChart>` renderiza SVG sem `role="img"` ou `aria-label` |

**Remediation:**
```tsx
<BarChart data={chartData} role="img" aria-label="Gráfico comparativo CBS e IBS: 2026 Transição e 2027 Projeção">
```
OU adicionar uma tabela escondida com os mesmos dados (`sr-only`) para leitores de tela.

---

#### 1.3.1 Info and Relationships (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Input CNPJ sem label programático | 🔴 Crítico | `App.tsx:25-31` | `<label>` e `<input>` não estão associados via `htmlFor`/`id` |
| Input NCM sem `<label>` associado por `htmlFor` | 🟠 Alto | `TaxSimulator.tsx:83-84` | Label "NCM:" usa `className` mas não `htmlFor` |
| Input Saldo sem `<label>` associado por `htmlFor` | 🟠 Alto | `TaxSimulator.tsx:90-91` | Label "Saldo Remanescente..." sem `htmlFor` |
| Resultados usam `<ul>` mas sem heading hierarchy adequada | 🟡 Médio | `TaxSimulator.tsx:109-126` | Cards de resultado com `<h3>` dentro de `<ul>` — hierarquia de headings quebrada |

**Remediation (CNPJ):**
```tsx
<label htmlFor="cnpj" className="block text-sm font-medium text-gray-700">
  CNPJ do Contribuinte
</label>
<input 
  id="cnpj"
  type="text" 
  className="mt-1 block w-full p-3 border rounded-md shadow-sm"
  placeholder="00.000.000/0000-00"
  value={cnpj}
  onChange={(e) => setCnpj(e.target.value)}
/>
```

---

#### 1.3.3 Sensory Characteristics (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Cores são o ÚNICO diferenciador entre cards Legado vs Reforma | 🟠 Alto | `TaxSimulator.tsx:107,118` | Usuários daltônicos podem não distinguir laranja de azul. Os títulos "Sistema Atual" e "Pós-Reforma" estão presentes, mas a associação de cor como significado não tem fallback |

**Remediation:** ✅ Os títulos textuais já servem como fallback. Adicionar ícones distintos (📋 para legado, 🔄 para reforma) como reforço não baseado em cor.

---

#### 1.4.3 Contrast — Minimum (Level AA)

| Finding | Severity | Location | Contraste | Status |
|---|---|---|---|---|
| Botão disabled (`bg-gray-400`) texto branco | 🟡 Médio | `TaxSimulator.tsx:97` | ~2.9:1 | ❌ Falha AA (mín 4.5:1) |

**Remediation:** Usar `text-gray-600` no estado disabled OU escurecer o background para `bg-gray-500`.

---

#### 1.4.4 Resize Text (Level AA)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Sem verificação prática | 🟡 Médio | Global | O layout usa classes Tailwind com `rem`/`em`, o que é bom, mas o `#root { width: 1126px }` fixo em `index.css:54` pode causar overflow com zoom de 200% |

---

### 2. Operable

#### 2.1.1 Keyboard (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| `alert()` é uma modal que requer clique do mouse para dispensar | 🟡 Médio | `App.tsx:13` | Usuários de teclado precisam usar mouse para fechar o alert |
| Selects e inputs são nativos (operáveis por teclado) | ✅ | — | Bom — sem custom components que bloqueiam keyboard nav |

---

#### 2.4.3 Focus Order (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Ordem de foco presumivelmente correta (DOM order = visual order) | ✅ | — | O fluxo é linear: CNPJ → submit → selects → inputs → simular |

---

#### 2.4.7 Focus Visible (Level AA)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Sem `:focus-visible` ou `outline` customizado em inputs e selects | 🔴 Crítico | `TaxSimulator.tsx:68,76,84,91` | Navegação por teclado é impossível — usuário não sabe onde está |
| Botão "ACESSAR PORTAL" sem focus style customizado | 🟠 Alto | `App.tsx:34-38` | Depende do outline default do browser, que pode ser suprimido por CSS reset |

**Remediation (todos os inputs):**
```css
input:focus-visible, select:focus-visible, button:focus-visible {
  outline: 2px solid #2563eb;
  outline-offset: 2px;
}
```

---

### 3. Understandable

#### 3.2.2 Labels or Instructions (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Placeholder "00000000000000" não ensina o formato esperado | 🟢 Baixo | `App.tsx:28` | Usuário pode não saber se deve incluir pontuação |
| Sem indicação de campos obrigatórios (Estado, Cidade) | 🟢 Baixo | `TaxSimulator.tsx:68,76` | Usuário pode não perceber que cidade é required até o botão ficar disabled |

---

#### 3.3.1 Error Identification (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Erro de CNPJ usa `alert()` — não descreve o erro especificamente nem sugere correção | 🔴 (já contado) | `App.tsx:13` | "Por favor, insira um CNPJ válido" não diz O QUE está errado (formato? dígitos? quantidade?) |

---

#### 3.3.2 Error Suggestion (Level AA)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Nenhuma sugestão de correção ao digitar CNPJ inválido | 🟠 Alto | `App.tsx:10-14` | Usuário precisa adivinhar o formato correto |

**Remediation:**
```tsx
{cnpj && cnpj.length < 14 && (
  <p className="text-red-600 text-sm mt-1" role="alert">
    CNPJ deve ter 14 dígitos numéricos. Formato: 00.000.000/0000-00
  </p>
)}
```

---

### 4. Robust

#### 4.1.2 Name, Role, Value (Level A)

| Finding | Severity | Location | Impact |
|---|---|---|---|
| Selects e inputs usam elementos nativos — role/name implícito | ✅ | — | Bom — elementos HTML nativos têm semântica built-in |
| Gráfico Recharts não expõe role/name para leitores de tela | 🔴 (já contado) | `TaxSimulator.tsx:131` | SVG do Recharts não tem ARIA labels |
| Botão "Sair/Trocar CNPJ" é `<button>` nativo — ✅ | ✅ | `App.tsx:45-49` | — |
| Botão "ACESSAR PORTAL" é `<button type="submit">` nativo — ✅ | ✅ | `App.tsx:34-38` | — |

---

## Checklist de Verificação (para re-teste com browser)

- [ ] Navegação completa por teclado: Tab → preencher CNPJ → Tab → ACESSAR → Tab → Estado → Tab → Cidade → Tab → NCM → Tab → Saldo → Tab → SIMULAR
- [ ] Focus indicator visível em TODOS os elementos interativos
- [ ] Leitor de tela anuncia: "CNPJ do Contribuinte, edit text" (não apenas "edit text")
- [ ] Leitor de tela anuncia loading state: "Processando..." quando API é chamada
- [ ] Leitor de tela anuncia resultado: navegação por headings (H3 nos cards)
- [ ] Contraste verificado com ferramenta (axe-core / Lighthouse):
  - [ ] Body text vs background ≥ 4.5:1
  - [ ] Placeholder text vs background ≥ 4.5:1
  - [ ] Botão disabled vs background ≥ 3:1
- [ ] Zoom 200%: todo conteúdo visível sem scroll horizontal
- [ ] Gráfico: alternativa textual disponível para leitores de tela
- [ ] `prefers-reduced-motion`: animações respeitam a preferência do OS
- [ ] `prefers-color-scheme: dark`: todos os contrastes mantêm ≥ 4.5:1

---

## Priorização para Remediação

| Ordem | Finding | Critério WCAG | Esforço |
|---|---|---|---|
| 1 | Label programático no input CNPJ | 1.3.1 A | 5 min |
| 2 | Focus indicators visíveis | 2.4.7 AA | 10 min |
| 3 | Labels programáticos nos campos do simulador | 1.3.1 A | 10 min |
| 4 | Substituir `alert()` por erro inline | 3.3.1 A + 3.3.2 AA | 20 min |
| 5 | ARIA label no gráfico Recharts | 1.1.1 A | 10 min |
| 6 | Contraste botão disabled | 1.4.3 AA | 5 min |
| 7 | Heading hierarchy nos cards de resultado | 1.3.1 A | 5 min |
| 8 | Placeholder descritivo CNPJ | 3.2.2 A | 2 min |

**Tempo total estimado para conformidade AA:** ~1 hora de desenvolvimento.

---

🤖 *Auditoria WCAG gerada por mineração de especificações frontend (PROMPT-MINING-FRONTEND-SPECIFICATION). Verificações manuais e scans automatizados indisponíveis — browser não pôde ser instalado neste ambiente.*
