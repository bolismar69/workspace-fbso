# CONCERNS — Solar Fácil Site

> Dívida técnica, bugs conhecidos, riscos de segurança e gargalos de performance.
> Gerado por `acquire-codebase-knowledge` em 2026-07-08.
> Fontes: análise do código-fonte, `.specs/INDEX.md`, `.specs/SECURITY.md`.

---

## 1. Dívida Técnica

### 1.1 Críticas (Devem ser resolvidas antes de produção)

| # | Item | Impacto | Localização |
|---|---|---|---|
| D1 | **ZERO testes automatizados** — 57 arquivos fonte sem cobertura | Risco de regressão em cada deploy | Todo o projeto |
| D2 | **Placeholders em produção** — WhatsApp, App Store, Google Play URLs falsas | Usuários reais encontrarão links quebrados | `constants.ts:130-132` |
| D3 | **Endpoint de formulário placeholder** — `formspree.io/f/placeholder` | Formulário de contato não funciona | `useContactForm.ts:7` |
| D4 | **Sem headers CSP/HSTS** — vulnerável a XSS, clickjacking mitigado mas não prevenido | Segurança | `next.config.ts` |

### 1.2 Altas (Devem ser resolvidas no próximo ciclo)

| # | Item | Impacto | Localização |
|---|---|---|---|
| D5 | **Duplicação de dados** — `PLANS` em `constants.ts` e `mockPlans.json` com estruturas diferentes | Inconsistência de dados, manutenção duplicada | `constants.ts` vs `mocks/mockPlans.json` |
| D6 | **Serviços com delay simulado** — `setTimeout(500ms)` em todos os services | UX artificial — ou migrar para API real ou remover delay | `services/service*.ts` |
| D7 | **Sem dark mode** — comentado como "milestone futuro" em `globals.css:38-40` | Funcionalidade esperada por usuários | `globals.css` |
| D8 | **Eventos GA4 não conectados** — `faq_open` e `lead_capture` definidos mas não disparados | Dados de analytics incompletos | `analytics.ts` vs componentes |

### 1.3 Médias (Planejadas)

| # | Item | Impacto | Localização |
|---|---|---|---|
| D9 | **Sem gerenciador de estado** — `useState` local funciona agora, mas se a app crescer... | Refatoração futura necessária | Todos os hooks |
| D10 | **Validação apenas client-side** — formulário sem validação server-side | Segurança (bypassável) | `useContactForm.ts`, `validation.ts` |
| D11 | **Sem i18n real** — PRODUCT.md menciona `/es/` para espanhol mas não há implementação | Funcionalidade internacional pendente | Todo o projeto |
| D12 | **`console.error` em produção** — logs de erro visíveis no console do usuário | Profissionalismo, debug | `services/service*.ts`, hooks |

## 2. Bugs Conhecidos

| # | Bug | Severidade | Evidência |
|---|---|---|---|
| B1 | Nenhum bug documentado | — | Análise estática apenas — sem testes para identificar bugs |

> ⚠️ **Nota**: Sem testes automatizados, bugs só seriam descobertos por inspeção manual ou em produção. Este é o risco #1 do projeto.

## 3. Riscos de Segurança

### 3.1 Identificados

| Risco | Severidade | Descrição | Mitigação Atual |
|---|---|---|---|
| **XSS** | Alta | Sem CSP header definido | Nenhuma (React mitiga XSS básico via escape) |
| **Clickjacking** | Média | Sem frame protections além de `X-Frame-Options: DENY` | Header configurado |
| **Form spam** | Média | Formulário sem CAPTCHA | Honeypot + tempo mínimo (3s) |
| **Data exposure** | Baixa | `NEXT_PUBLIC_GA_ID` exposto ao cliente | Por design (GA4 requer client-side) |
| **MITM** | Baixa | Sem HSTS header | Depende de HTTPS do servidor |

### 3.2 Recomendações Imediatas

1. **Adicionar header CSP** em `next.config.ts` — mínimo: `default-src 'self'; script-src 'self' 'unsafe-inline' https://www.googletagmanager.com`
2. **Adicionar header HSTS** — `Strict-Transport-Security: max-age=31536000; includeSubDomains`
3. **Substituir endpoint placeholder** do formulário
4. **Adicionar validação server-side** no formulário (API Route ou Server Action)

## 4. Gargalos de Performance

### 4.1 Identificados

| Gargalo | Severidade | Descrição |
|---|---|---|
| **Delay artificial** | Média | `setTimeout(500ms)` em todos os services — UX desnecessariamente lenta |
| **Tailwind v4 build** | Baixa | Compilação CSS pode ser pesada em cold builds |
| **Sem image optimization real** | Baixa | Sem imagens no site atualmente (PRODUCT.md menciona fotografia) |

### 4.2 Oportunidades de Otimização

- **Remover delay simulado** se não houver plano de migração para API real em <3 meses
- **Adicionar `loading.tsx`** para streaming SSR nas páginas (Next.js App Router)
- **Adicionar `prefetch`** para links de navegação entre páginas

## 5. Perguntas em Aberto [ASK USER]

Estas questões foram identificadas no `.specs/INDEX.md` e permanecem sem resposta:

- [ ] **Qual é o endpoint real do formulário?** O padrão é um placeholder do Formspree.
- [ ] **Qual é o número real de WhatsApp?** `5511999999999` é claramente placeholder.
- [ ] **Os links App Store / Google Play são reais?** Ambas as URLs apontam para IDs de app placeholder.
- [ ] **Há uma API backend planejada?** Atualmente todos os dados são estáticos.
- [ ] **Quando o modo escuro será implementado?** Comentado como "milestone futuro".
- [ ] **Os datasets JSON mock ou as constantes em código são a fonte da verdade?** Há duplicação.
- [ ] **A camada de serviços será substituída por chamadas de API reais?** O delay de 500ms sugere que sim.
- [ ] **Qual é o cronograma de internacionalização?** PRODUCT.md menciona `/es/` para espanhol.
- [ ] **Qual política CSP deve ser adotada?** Atualmente nenhum header CSP está definido.

## 6. Score de Saúde do Projeto

| Área | Score | Nota |
|---|---|---|
| **Qualidade de Código** | 🟢 8/10 | TypeScript strict, convenções claras, boa organização |
| **Cobertura de Testes** | 🔴 0/10 | Zero testes automatizados |
| **Segurança** | 🟡 5/10 | Headers básicos, sem CSP/HSTS, sem validação server-side |
| **Documentação** | 🟢 9/10 | `.specs/` rico e detalhado, múltiplos formatos |
| **Dívida Técnica** | 🟡 5/10 | Placeholders críticos, sem testes, duplicação de dados |
| **Performance** | 🟢 8/10 | App estática leve, sem APIs externas, Tailwind otimizado |
| **Acessibilidade** | 🟡 6/10 | Planejada WCAG AA, mas sem auditoria formal |
| **Geral** | 🟡 5.9/10 | Bom código, mas débitos críticos de testes e segurança |

---

Última atualização: 2026-07-08
