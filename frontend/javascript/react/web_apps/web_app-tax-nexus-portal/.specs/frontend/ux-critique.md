---
title: "Avaliação Heurística de UX — TaxNexus Portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["ux", "heuristics", "nielsen", "critique"]
critique_method: "static-code-analysis"
browser_available: false
---

# Avaliação Heurística de UX — TaxNexus Portal

**Método:** Análise das 10 heurísticas de Nielsen sobre o código-fonte e a estrutura da aplicação.
**Alvo:** Fluxo completo — Login (CNPJ) → Simulador → Resultados.
**⚠️ DEGRADED: single-context (no sub-agent tool available for dual-assessment; browser unavailable for visual inspection).**

---

## Design Health Score

| # | Heuristic | Score | Key Issue |
|---|-----------|-------|-----------|
| 1 | Visibility of System Status | **3**/4 | Loading state no botão, mas sem indicador de progresso durante o cálculo |
| 2 | Match System / Real World | **4**/4 | Terminologia fiscal correta (CNPJ, NCM, CBS, IBS); formatação BRL nativa |
| 3 | User Control and Freedom | **3**/4 | "Sair/Trocar CNPJ" sempre visível; mas sem confirmação ao descartar simulação |
| 4 | Consistency and Standards | **3**/4 | Padrão consistente de inputs/selects; `alert()` quebra convenção web moderna |
| 5 | Error Prevention | **2**/4 | Sem máscara CNPJ; NCM aceita qualquer texto; sem validação em tempo real |
| 6 | Recognition Rather Than Recall | **3**/4 | CNPJ visível no header pós-login; select de cidade depende de estado (bom), mas sem ajuda contextual |
| 7 | Flexibility and Efficiency | **2**/4 | Sem atalhos de teclado; sem autocomplete nos campos; formulário linear sem bulk actions |
| 8 | Aesthetic and Minimalist Design | **3**/4 | Layout limpo e focado; gráfico Recharts bem integrado; leve excesso de cores (6 tons diferentes) |
| 9 | Error Recovery | **1**/4 | `alert()` não sugere correção; sem tratamento de timeout; sem feedback de erro da API |
| 10 | Help and Documentation | **1**/4 | Nenhum tooltip, nenhuma ajuda contextual, nenhum link de documentação |
| **Total** | | **25/40** | **Acceptable** — interface funcional, mas precisa de trabalho significativo em error recovery e ajuda |

**Rating bands:** 36-40 Excellent | 28-35 Good | 20-27 Acceptable | 12-19 Poor | 0-11 Critical

---

## Anti-Patterns Verdict

**Não parece AI-generated.** A interface tem propósito claro e domínio específico (tributário brasileiro) que resiste ao "slop test". As cores semânticas (laranja = legado, azul = reforma) ensinam o modelo mental sem parecerem decorativas.

**LLM assessment:** O layout é funcional e direto — um formulário vertical que conduz a cards comparativos e um gráfico. Não há ornamentação desnecessária. A fraqueza não é estética, é de usabilidade: falta prevenção de erros e ajuda contextual.

**Deterministic scan:** N/A (detector CLI não disponível neste ambiente).

---

## Overall Impression

O TaxNexus Portal é **funcional e semanticamente correto** — um controller tributário reconhece imediatamente o domínio. A interface faz o mínimo necessário bem: formulário → cálculo → visualização. O problema é que ela para no "mínimo" — faltam camadas de polimento que transformam uma ferramenta funcional em uma ferramenta profissional: prevenção de erros, recuperação, e ajuda contextual.

O single biggest opportunity: **error handling**. O `alert()` para CNPJ inválido é o sintoma de uma filosofia "happy path only". Uma ferramenta fiscal precisa ser resilient — o usuário vai errar, e a interface precisa ajudar, não punir.

---

## What's Working

1. **Linguagem de domínio impecável.** CBS, IBS, NCM, CNPJ, UF, saldo remanescente — todos os termos são corretos e familiares ao público-alvo. A formatação `pt-BR` com `toLocaleString` é o detalhe que faz a diferença para profissionais fiscais
2. **Modelo mental ensinado via cor.** Laranja (quente, "velho") para sistema atual e azul (frio, "novo") para pós-reforma é uma escolha intuitiva que ajuda o usuário a navegar a comparação sem precisar ler labels
3. **Fluxo linear sem distrações.** CNPJ → localidade → NCM → saldo → resultado. Sem modais, sem popups, sem cross-selling. A ferramenta respeita o foco do usuário

---

## Priority Issues

### [P1] Error handling é frágil e não ajuda o usuário

**Why it matters:** `alert()` é uma janela modal do browser — bloqueia a interação, não pode ser estilizada, e é ignorada por alguns leitores de tela. Quando a API falha (timeout, erro 500), o usuário não recebe feedback algum — o botão simplesmente volta de "Processando..." para "SIMULAR" sem explicação

**Fix:**
1. Criar componente `<FormError>` com `role="alert"` e `aria-live="polite"` para erros inline
2. Adicionar `catch` no `useTaxService` que retorna mensagem de erro amigável
3. Mostrar erro da API no componente de resultado com ação sugerida ("Tente novamente", "Verifique o CNPJ")

**Suggested command:** `$impeccable harden` para error states, `$impeccable clarify` para mensagens

### [P1] Prevenção de erros inexistente

**Why it matters:** O formulário aceita qualquer input sem validação — CNPJ com letras, NCM inexistente, saldo negativo. O usuário só descobre o erro após submit (e às vezes nem isso, se a API retornar 200 com dados estranhos)

**Fix:**
1. Máscara de CNPJ: `__.___.___/____-__` com validação de dígitos verificadores
2. NCM: validar contra lista dos 8 dígitos, ou ao menos regex `\d{8}`
3. Saldo: impedir valores negativos no input number (`min="0"`)
4. Selects: desabilitar "Simular" até todos os campos obrigatórios preenchidos

**Suggested command:** `$impeccable harden`

### [P2] Sem atalhos para usuários frequentes

**Why it matters:** Um controller tributário pode fazer dezenas de simulações por dia. Cada simulação requer: selecionar UF → selecionar cidade → digitar NCM → digitar saldo → clicar. Não há autocomplete, nem valores padrão inteligentes, nem histórico

**Fix:**
1. Adicionar `autocomplete` attributes nos inputs
2. Salvar última UF/Cidade/NCM no `localStorage` para pré-preenchimento
3. Permitir Enter para submeter o formulário após preencher todos os campos

**Suggested command:** `$impeccable onboard` para fluxo de power user

### [P2] Zero ajuda contextual em um domínio complexo

**Why it matters:** O sistema tributário brasileiro é notoriamente complexo. Um CFO que não é especialista fiscal pode não saber o que é "saldo remanescente de créditos" ou qual NCM se aplica ao seu produto

**Fix:**
1. Tooltips com ícone `?` ao lado de "Saldo Remanescente" e "NCM"
2. Link para tabela NCM ou autocomplete com descrição do produto
3. Footer com "Precisa de ajuda? Consulte a documentação da Reforma Tributária"

**Suggested command:** `$impeccable clarify` para copy, `$impeccable layout` para posicionamento de tooltips

---

## Persona Red Flags

### Alex (Power User — Controller Tributário)
- ❌ Sem atalhos de teclado; cada simulação requer 5 interações manuais com mouse
- ❌ Sem histórico de simulações anteriores; precisa redigitar tudo do zero
- ❌ Sem autocomplete ou valores padrão; formulário é "burro" entre sessões
- **Risco de abandono:** Médio — a ferramenta funciona, mas o atrito para uso repetido é alto

### Jordan (First-Timer — CFO explorando a ferramenta)
- ❌ Sem explicação do que é "NCM" ou "Saldo Remanescente" — assume conhecimento prévio
- ❌ Placeholder "00000000000000" não ensina o formato esperado do CNPJ
- ❌ Sem indicação visual de progresso — o usuário não sabe quantos passos faltam
- **Risco de abandono:** Alto — Jordan pode desistir no primeiro campo que não entender

### Sam (Accessibility-Dependent — Usuário de leitor de tela)
- ❌ Input de CNPJ sem `<label>` associado via `htmlFor` — leitor de tela anuncia "edit text" sem contexto
- ❌ Sem focus indicators visíveis — usuário de teclado não sabe onde está
- ❌ `alert()` pode não ser anunciado por todos os leitores de tela
- ❌ Gráfico Recharts sem descrição textual alternativa — dados completamente invisíveis
- **Risco de abandono:** Crítico — Sam não consegue passar da tela de login

---

## Minor Observations

- O título da página (`<title>`) é "web_app-tax-nexus-portal" — deveria ser "TaxNexus Portal"
- O botão "ACESSAR PORTAL" está em all-caps — considerar sentence case para legibilidade
- A transição entre login e simulador é instantânea — um fade sutil (150ms) melhoraria a percepção de contexto
- O gráfico mostra "2027 (Projeção)" mas isso não está rotulado em lugar nenhum — o rótulo do eixo X diz apenas "2027 (Projeção)" mas o valor é multiplicado por 1.02 sem aviso

---

## Questions to Consider

- "O que acontece quando a API está offline? O usuário tem feedback ou fica esperando infinitamente?"
- "E se o usuário quiser comparar DOIS cenários lado a lado (ex: SP vs RJ para o mesmo produto)?"
- "Um gráfico de barras é a melhor visualização para comparar apenas 2 pontos no tempo?"
- "Deveria existir um modo 'avançado' com mais parâmetros (alíquota interestadual, substituição tributária)?"

---

## Recommended Actions

1. **[P1] `$impeccable harden`** — Error handling completo: substituir `alert()`, adicionar máscara CNPJ, validação em tempo real, tratamento de timeout
2. **[P1] `$impeccable harden`** — Acessibilidade: labels programáticos, focus indicators, alt text para gráfico
3. **[P2] `$impeccable clarify`** — Tooltips e ajuda contextual para NCM e Saldo Remanescente
4. **[P2] `$impeccable onboard`** — Salvar preferências no localStorage, valores padrão inteligentes
5. **[P3] `$impeccable polish`** — Título da página, placeholders descritivos, fade na transição de tela

> Re-run `$impeccable critique` after fixes to see your score improve.

---

🤖 *Avaliação heurística gerada por mineração de especificações frontend (PROMPT-MINING-FRONTEND-SPECIFICATION). Análise por código — browser e sub-agentes não disponíveis neste ambiente.*
