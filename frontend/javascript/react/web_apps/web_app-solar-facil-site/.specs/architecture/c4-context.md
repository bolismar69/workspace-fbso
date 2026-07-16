# C4 Context — Solar Fácil Site

> Nível 1: O sistema no ecossistema — atores (usuários) e sistemas externos.
> Gerado por `architecture-designer` em 2026-07-08.

---

## Diagrama de Contexto (Mermaid)

```mermaid
C4Context
    title Solar Fácil Site — Diagrama de Contexto (Nível 1)

    Person(consumidor, "Consumidor", "Pessoa física ou pequena empresa<br/>que quer reduzir a conta de luz")
    Person(fornecedor, "Fornecedor/Produtor", "Quem já tem painéis solares<br/>e quer rentabilizar o excedente")
    Person(cooperativa, "Cooperativa", "Parceiro B2B que opera<br/>a distribuição de energia")

    System(site, "Solar Fácil Site", "Portal web Next.js<br/>Calculadora, Planos, Contato")

    System_Ext(formspree, "Formspree", "Serviço de formulário<br/>(placeholder)")
    System_Ext(ga4, "Google Analytics 4", "Analytics e tracking<br/>de eventos")
    System_Ext(gfonts, "Google Fonts", "Fonte Inter — servida<br/>via next/font (build time)")

    Rel(consumidor, site, "Calcula economia, vê planos,<br/>envia contato", "HTTPS")
    Rel(fornecedor, site, "Calcula ganho, vê planos,<br/>envia contato", "HTTPS")
    Rel(cooperativa, site, "Consulta informações<br/>de integração", "HTTPS")

    Rel(site, formspree, "Envia dados do formulário", "HTTPS/POST")
    Rel(site, ga4, "Dispara eventos de tracking", "HTTPS/gtag.js")
    Rel(site, gfonts, "Obtém fonte no build time", "next/font (offline em prod)")
```

## Descrição dos Elementos

### Atores

| Ator | Descrição | Necessidade Principal |
|---|---|---|
| **Consumidor** | Pessoa física ou pequena empresa. Acessa principalmente via mobile, durante horário comercial. Não é especialista em energia. | "Quanto vou economizar sem instalar nada?" |
| **Fornecedor/Produtor** | Já tem painéis solares instalados. Perfil mais técnico. Quer rentabilizar excedente. | "Quanto posso ganhar compartilhando meu excedente?" |
| **Cooperativa** | Parceiro B2B. Opera a distribuição regulamentada pela ANEEL. | "Como integrar com a plataforma?" |

### Sistemas Externos

| Sistema | Tipo | Status | Integração |
|---|---|---|---|
| **Formspree** | Serviço de formulário HTTPS | ⚠️ Placeholder | POST multipart/form-data |
| **Google Analytics 4** | Analytics client-side | ✅ Condicional (env var) | `window.gtag()` |
| **Google Fonts** | CDN de fontes | ✅ Build time | `next/font/google` (self-hosted em prod) |

### Fluxos de Dados Principais

1. **Consumidor → Site**: Navegação, input na calculadora, submissão de formulário
2. **Fornecedor → Site**: Navegação, input na calculadora (modo fornecedor), submissão de formulário
3. **Site → Formspree**: POST com FormData (nome, email, telefone, perfil, mensagem)
4. **Site → GA4**: Eventos `cta_click`, `calculator_use`, `faq_open`, `lead_capture`

---

Última atualização: 2026-07-08
