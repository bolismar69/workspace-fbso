---
title: "Verificação Funcional — TaxNexus Portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["functional-verification", "playwright", "testing", "user-flows"]
verification_method: "static-code-analysis"
browser_available: false
---

# Verificação Funcional — TaxNexus Portal

**⚠️ Método degradado:** Análise estática dos fluxos de usuário a partir do código-fonte. Execução com Playwright indisponível (browser não pôde ser instalado neste ambiente). Os fluxos abaixo documentam o comportamento **esperado** conforme implementado no código.

---

## Fluxos de Usuário Mapeados

### Fluxo 1: Login com CNPJ → Acesso ao Simulador

**Arquivos:** `App.tsx:8-15`

| Passo | Ação | Estado Esperado | Verificação |
|---|---|---|---|
| 1 | Usuário acessa `http://localhost:5173` | Tela de login renderizada | ✅ |
| 2 | Usuário digita CNPJ (14+ dígitos) | Texto aparece no input | ✅ |
| 3 | Usuário clica "ACESSAR PORTAL" | `handleLogin` disparado | ✅ |
| 4a | CNPJ.length >= 14 | `authenticated = true`, renderiza `<TaxSimulator>` | ✅ |
| 4b | CNPJ.length < 14 | `alert("CNPJ inválido")` | ⚠️ (antipattern) |
| 5 | Tela do simulador visível | Header com CNPJ + formulário | ✅ |

**Edge cases testados (código):**
- CNPJ vazio + submit → `alert()` exibido
- CNPJ "12345678901234" (14 chars) → acesso permitido
- CNPJ "1234567890123" (13 chars) → `alert()` exibido

**⚠️ Issues encontrados:**
- Qualquer string de 14 caracteres passa como CNPJ válido (sem validação de formato/dígitos)
- `alert()` bloqueia a UI e não é acessível

---

### Fluxo 2: Preencher Formulário → Simular → Ver Resultados

**Arquivos:** `TaxSimulator.tsx:30-55`, `hooks/useTaxService.ts`

| Passo | Ação | Estado Esperado | Verificação |
|---|---|---|---|
| 1 | Selecionar Estado (ex: SP) | `selectedState = '35'`, select cidade habilitado | ✅ |
| 2 | Selecionar Cidade (ex: São Paulo) | `selectedCity = '3550308'`, botão SIMULAR habilitado | ✅ |
| 3 | Ajustar NCM (default: 62011100) | Input editável com valor padrão | ✅ |
| 4 | Digitar Saldo Remanescente | Input number aceita valor | ✅ |
| 5 | Clicar "SIMULAR REFORMA TRIBUTÁRIA" | Botão muda para "Processando..." + disabled | ✅ |
| 6 | POST `http://localhost:8080/v1/tax/calculate` | Request enviado com body JSON | ✅ |
| 7a | API retorna 200 com `calculation` | Cards laranja/azul + gráfico renderizados | ✅ |
| 7b | API retorna erro | ❌ Sem tratamento — loading termina silenciosamente | ❌ |
| 7c | API timeout (>30s) | ❌ Sem AbortController — loading eterno | ❌ |
| 8 | Resultados visíveis | Cards: PIS, COFINS, ICMS, IPI \| CBS, IBS, Seletivo, IPVA/ITCMD | ✅ |
| 9 | Gráfico de barras visível | 2026 (Transição) vs 2027 (Projeção ×1.02) | ✅ |

**Dados enviados na request:**
```json
{
  "cnpj": "valor do input de login",
  "ncm": "62011100",
  "ibge": "3550308",
  "saldo_remanescente": 0
}
```

**⚠️ Issues encontrados:**
- Timeout não tratado — se API estiver offline, loading é eterno
- Erro da API não exibe feedback ao usuário
- Projeção 2027 hardcoded com multiplicador ×1.02 sem indicação de estimativa

---

### Fluxo 3: Trocar CNPJ (Logout) → Retornar à Tela de Login

**Arquivos:** `App.tsx:45-49`

| Passo | Ação | Estado Esperado | Verificação |
|---|---|---|---|
| 1 | Usuário está na tela do simulador | Header + formulário visíveis | ✅ |
| 2 | Clicar "Sair / Trocar CNPJ" | `setAuthenticated(false)` | ✅ |
| 3 | Tela de login renderizada | CNPJ em branco, formulário resetado | ✅ |

**⚠️ Issues encontrados:**
- Sem confirmação antes de sair ("Tem certeza? Dados da simulação serão perdidos")
- Estado da simulação atual é perdido sem warning

---

### Fluxo 4: Tratamento de Erros e Edge Cases

| Cenário | Comportamento Esperado | Implementado? | Severidade |
|---|---|---|---|
| API offline (conexão recusada) | Mensagem de erro: "Serviço indisponível. Tente novamente." | ❌ Não | 🔴 P0 |
| API retorna 4xx | Mensagem específica por código de erro | ❌ Não | 🔴 P0 |
| API retorna 5xx | Mensagem: "Erro interno. Nossa equipe foi notificada." | ❌ Não | 🔴 P0 |
| CNPJ com letras | Validação em tempo real rejeita | ❌ Não (aceita qualquer texto) | 🟠 P1 |
| NCM com letras | Validação em tempo real rejeita | ❌ Não (aceita qualquer texto) | 🟠 P1 |
| Saldo negativo | Input `type="number"` permite negativo | ⚠️ Permitido pelo HTML | 🟡 P2 |
| Estado sem cidade selecionada | Botão disabled | ✅ Sim | ✅ |
| Duplo clique em SIMULAR | Botão disabled durante loading previne | ✅ Sim | ✅ |
| Refresh durante simulação | Estado perdido (React state) | ⚠️ Esperado para SPA | 🟢 P3 |
| CNPJ muito longo (30+ chars) | Input aceita (sem `maxLength`) | ⚠️ Sem limite | 🟡 P2 |

---

### Fluxo 5: Responsividade entre Dispositivos

**Verificação por viewport (esperada, não testada):**

| Viewport | Comportamento Esperado | Riscos Identificados |
|---|---|---|
| 375px (iPhone SE) | Formulário single-column, cards empilhados | Selects `md:grid-cols-3` colapsam — verificar se `grid-cols-1` é o default |
| 768px (iPad) | Grid 2 colunas para selects? | Classe `md:grid-cols-3` aplica 3 colunas em 768px — pode ficar apertado |
| 1280px (Desktop) | Layout ótimo: formulário + cards lado a lado | ✅ |
| 1920px (Large) | Container `max-w-4xl` (896px) centralizado | Muito espaço vazio nas laterais |

---

## Resumo da Verificação

| Fluxo | Status | Issues Críticos |
|---|---|---|
| F1: Login | ⚠️ Passa, mas frágil | `alert()`, sem validação real de CNPJ |
| F2: Simular | ⚠️ Happy path passa | Sem error handling da API, sem timeout |
| F3: Logout | ✅ Passa | Sem confirmação de descarte |
| F4: Erros | ❌ Reprovado | Nenhum cenário de erro tratado |
| F5: Responsivo | ⚠️ Não testado | Riscos em mobile e wide desktop |

**Happy path:** ✅ Funcional — o fluxo principal (CNPJ → simular → ver resultados) funciona.
**Error path:** ❌ Não implementado — qualquer falha resulta em experiência quebrada.

---

## Recomendações para Teste Real (Playwright)

Script Playwright sugerido para quando o browser estiver disponível:

```typescript
// F1: Login
await page.goto('http://localhost:5173');
await page.fill('input[placeholder*="0000"]', '12345678901234');
await page.click('button:has-text("ACESSAR PORTAL")');
await expect(page.locator('text=TaxNexus Simulator v1.0')).toBeVisible();

// F2: Simular
await page.selectOption('select >> nth=0', '35'); // SP
await page.selectOption('select >> nth=1', '3550308'); // São Paulo
await page.click('button:has-text("SIMULAR")');
await expect(page.locator('text=Processando...')).toBeVisible();
// Aguardar resultado (depende da API)
await expect(page.locator('text=Sistema Atual')).toBeVisible({ timeout: 30000 });

// F3: Logout
await page.click('text=Sair / Trocar CNPJ');
await expect(page.locator('text=TaxNexus TaaS')).toBeVisible();

// F4: Erro CNPJ
await page.fill('input[placeholder*="0000"]', '123');
await page.click('button:has-text("ACESSAR PORTAL")');
// ⚠️ Isso vai disparar alert() — Playwright precisa de dialog handler
```

---

🤖 *Verificação funcional gerada por mineração de especificações frontend (PROMPT-MINING-FRONTEND-SPECIFICATION). Testes com Playwright indisponíveis — browser não pôde ser instalado neste ambiente.*
