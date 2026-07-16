# Security — Solar Fácil Site

> Definições de segurança: CSP, CORS, OWASP Top 10 frontend, auth, headers.
> Gerado por `documentation-writer` em 2026-07-08.
> Integra conteúdo do `.specs/SECURITY.md` original (2026-07-06).

---

## 1. Headers de Segurança

### 1.1 Implementados

| Header | Valor | Status |
|---|---|---|
| `X-Frame-Options` | `DENY` | ✅ |
| `X-Content-Type-Options` | `nosniff` | ✅ |
| `Referrer-Policy` | `strict-origin-when-cross-origin` | ✅ |

### 1.2 Recomendados (Ausentes)

| Header | Valor Recomendado | Prioridade |
|---|---|---|
| `Content-Security-Policy` | `default-src 'self'; script-src 'self' 'unsafe-inline' https://www.googletagmanager.com; style-src 'self' 'unsafe-inline';` | 🔴 Alta |
| `Strict-Transport-Security` | `max-age=31536000; includeSubDomains` | 🔴 Alta |
| `Permissions-Policy` | `camera=(), microphone=(), geolocation=()` | 🟡 Média |

---

## 2. OWASP Top 10 — Frontend

| Risco | Mitigação | Status |
|---|---|---|
| **A01: Broken Access Control** | Sem autenticação — não aplicável | N/A |
| **A02: Cryptographic Failures** | Depende de HTTPS do servidor | ⚠️ Sem HSTS |
| **A03: Injection** | React escapa JSX por padrão | ✅ |
| **A04: Insecure Design** | Sem validação server-side no formulário | ⚠️ |
| **A05: Security Misconfiguration** | Headers básicos presentes, sem CSP | ⚠️ |
| **A06: Vulnerable Components** | Dependências mínimas (3 runtime) | ✅ |
| **A07: Auth Failures** | Sem autenticação — não aplicável | N/A |
| **A08: Software/Data Integrity** | package-lock.json versionado | ✅ |
| **A09: Logging/Monitoring** | Apenas console.error client-side | ⚠️ |
| **A10: SSRF** | Sem chamadas server-side externas | ✅ |

---

## 3. Segurança do Formulário

| Proteção | Implementação |
|---|---|
| **Honeypot** | Campo hidden `website` — bots preenchem, humanos não |
| **Tempo mínimo** | 3s entre page load e submit |
| **Validação client-side** | Nome, email, telefone, perfil |
| **Validação server-side** | ❌ Ausente — TODO |
| **Rate limiting** | ❌ Ausente — depende do Formspree |
| **CAPTCHA** | ❌ Ausente |

---

## 4. Configuração de Ambiente

```bash
# .env.example
NEXT_PUBLIC_GA_ID=G-XXXXXXXXXX              # Google Analytics 4
NEXT_PUBLIC_FORM_ENDPOINT=https://...        # Endpoint formulário (substituir placeholder)
```

**Regras:**
- ✅ `NEXT_PUBLIC_*` apenas para variáveis expostas ao browser
- ✅ `.env.example` versionado (sem secrets)
- ✅ Valores default seguros no código

---

## 5. Dependências

| Risco | Status |
|---|---|
| **Número de dependências** | 🟢 3 runtime + 10 dev — superfície mínima |
| **Versões fixas** | ⚠️ `^` ranges (sem `package-lock.json` audit regular) |
| **Vulnerabilidades conhecidas** | ❓ Não verificado (`npm audit` pendente) |

---

## 6. TODO — Segurança

- [ ] Adicionar CSP header
- [ ] Adicionar HSTS header
- [ ] Adicionar validação server-side no formulário
- [ ] Rodar `npm audit` e corrigir vulnerabilidades
- [ ] Substituir Formspree placeholder por endpoint real
- [ ] Considerar rate limiting no formulário
- [ ] Migrar `console.error` para serviço de logging
- [ ] Adicionar `Permissions-Policy` header

---

Última atualização: 2026-07-08
