# Segurança e Configuração: Solar Fácil Site

> **Especificação de engenharia reversa** — padrões de segurança, configuração de ambiente e deploy.
> Gerado por `/spec-miner` em 2026-07-05. Revisado para pt-BR em 2026-07-06.

---

## 1. Headers de Segurança HTTP

**Arquivo**: `next.config.ts:13-24`

Todas as rotas recebem estes headers:

| Header | Valor | Propósito |
|--------|-------|-----------|
| `X-Frame-Options` | `DENY` | Previne clickjacking — o site não pode ser embutido em iframes |
| `X-Content-Type-Options` | `nosniff` | Previne sniffing de MIME-type |
| `Referrer-Policy` | `strict-origin-when-cross-origin` | Limita informações de referrer enviadas cross-origin |

**Headers ausentes** (não configurados):
- `Content-Security-Policy` — nenhum header CSP definido
- `Strict-Transport-Security` — sem header HSTS (importante para produção)
- `Permissions-Policy` — sem política de permissões

---

## 2. Segurança do Formulário

### Medidas Anti-Spam

**Arquivo**: `src/hooks/useContactForm.ts`

| Medida | Implementação | Linha |
|--------|---------------|-------|
| **Portão de tempo** | Bloqueia envios < 3 segundos do carregamento da página | `:8, :48-53` |
| **Campo honeypot** | Campo `website` oculto (tabIndex -1, posicionado fora da tela em -9999px left). Bots que o preenchem disparam falso sucesso. | `:9, :47-53` |
| **Falso sucesso** | Quando anti-spam é acionado, retorna tela de sucesso sem enviar dados — bots não conseguem distinguir | `:51-53` |

### Validação

**Arquivo**: `src/lib/validation.ts`

| Campo | Regra | Tipo |
|-------|-------|------|
| `name` | Obrigatório, mín 2 caracteres | Client-side |
| `email` | Obrigatório, regex `/^[^\s@]+@[^\s@]+\.[^\s@]+$/` | Client-side |
| `phone` | Opcional, 10–11 dígitos se informado | Client-side |
| `profile` | Obrigatório, não-vazio | Client-side |
| `message` | Opcional, máx 1000 caracteres | Client-side |

> ⚠️ **Constatação**: Toda validação é apenas client-side. Não existe validação server-side. O endpoint do formulário (`NEXT_PUBLIC_FORM_ENDPOINT`) é externo (Formspree ou similar), mas não há API route, server action ou camada de middleware de validação neste projeto Next.js.

### Endpoint do Formulário

**Arquivo**: `src/hooks/useContactForm.ts:7`

```typescript
const FORM_ENDPOINT = process.env.NEXT_PUBLIC_FORM_ENDPOINT || 'https://formspree.io/f/placeholder';
```

Os dados são enviados como `FormData` via `fetch()`. O endpoint é configurável por variável de ambiente. O padrão é um placeholder — o formulário falhará silenciosamente para uma URL placeholder do Formspree se não for configurado.

---

## 3. Variáveis de Ambiente

**Arquivo**: `.env.example`

| Variável | Padrão | Propósito | Exposição |
|----------|--------|-----------|-----------|
| `NEXT_PUBLIC_GA_ID` | `G-XXXXXXXXXX` | ID do Google Analytics 4 | Pública (client) |
| `NEXT_PUBLIC_FORM_ENDPOINT` | `https://formspree.io/f/your-form-id` | Endpoint de envio do formulário | Pública (client) |
| `NEXT_PUBLIC_SITE_URL` | `https://www.solarfacil.com.br` | URL canônica do site | Pública (client) |

Todas as variáveis de ambiente têm prefixo `NEXT_PUBLIC_` — são incorporadas em tempo de build e visíveis no JavaScript do lado do cliente. Isso é apropriado para esses valores, mas significa que **não há segredos server-side** no projeto.

---

## 4. Segurança de Conteúdo

### Configuração de Imagens
**Arquivo**: `next.config.ts:9-11`
- Formatos modernos habilitados: AVIF + WebP
- Nenhum padrão de imagem remota configurado (apenas imagens locais)

### robots.txt
**Arquivo**: `src/app/robots.ts`
- Permite todos os crawlers (`userAgent: '*'`, `allow: '/'`)
- URL do sitemap incluída

### Segurança de Metadados
**Arquivo**: `src/app/layout.tsx:14-35`
- `metadataBase` definido para URL canônica — previne spoofing de URL OpenGraph
- Robots: `index: true, follow: true`

---

## 5. Configuração de Deploy

### Output de Build
**Arquivo**: `next.config.ts:5`
```typescript
output: 'standalone'
```
Produz um build autocontido adequado para hospedagem não-Vercel (DigitalOcean, etc.). O output standalone inclui um server.js mínimo, todos os assets estáticos e o runtime do Next.js.

### TypeScript
**Arquivo**: `tsconfig.json`
- `strict: true` — verificação completa de tipos
- `noEmit: true` — Next.js cuida da compilação
- `moduleResolution: "bundler"` — resolução moderna
- `skipLibCheck: true` — builds mais rápidos (pula verificação de tipos em node_modules)

### ESLint
**Arquivo**: `eslint.config.mjs`
- `eslint-config-next/core-web-vitals` — regras de Core Web Vitals
- `eslint-config-next/typescript` — regras específicas TypeScript
- Ignora: `.next`, `out`, `build`, `next-env.d.ts`

---

## 6. Scripts de Terceiros

| Script | Origem | Carregamento | Condicional |
|--------|--------|-------------|-------------|
| Google Analytics 4 | `googletagmanager.com` | Assíncrono | Apenas se `GA_MEASUREMENT_ID` estiver configurado |
| Google Fonts (Inter) | `next/font/google` | Otimizado em build | Sempre (auto-hospeda arquivos de fonte em build via Next.js) |

Nenhum outro script de terceiros é carregado. Sem scripts CDN, sem gerenciadores de tags além do GA4.

---

## 7. Avaliação de Segurança

### Pontos Fortes
- ✅ Proteção contra clickjacking (X-Frame-Options: DENY)
- ✅ Proteção contra MIME-sniffing (X-Content-Type-Options)
- ✅ Privacidade de referrer (Referrer-Policy)
- ✅ Anti-spam duplo (portão de tempo + honeypot)
- ✅ TypeScript strict mode completo
- ✅ Sem scripts inline exceto GA4 (necessário)
- ✅ Auto-hospedagem de fontes via next/font (sem chamada CDN Google Fonts em runtime)
- ✅ Sem `dangerouslySetInnerHTML` exceto JSON-LD (seguro, conteúdo estático)

### Lacunas e Recomendações
- ❌ **Sem header CSP** — Adicionar Content-Security-Policy para restringir origens de script
- ❌ **Sem HSTS** — Adicionar Strict-Transport-Security para produção
- ❌ **Sem validação server-side** — Validação do formulário de contato é apenas client-side; adicionar API route ou server action
- ❌ **Número WhatsApp fixo** (`5511999999999`) — placeholder, precisa do número real
- ❌ **Links App Store** são placeholders — precisam de IDs reais
- ⚠️ **Endpoint de formulário placeholder** — Envios de formulário falham silenciosamente se não configurado
- ⚠️ **Sem rate limiting** — Formulário de contato tem portão de tempo mas sem limitação de taxa baseada em IP
