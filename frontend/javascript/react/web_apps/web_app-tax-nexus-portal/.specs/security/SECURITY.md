---
title: "Segurança — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["security", "owasp", "frontend"]
---

# Definições de Segurança — TaxNexus Portal

## 1. Estado Atual da Segurança

O TaxNexus Portal encontra-se em **estágio MVP**, com diversas questões de segurança identificadas que precisam ser resolvidas antes de qualquer exposição a ambientes de produção.

## 2. Autenticação e Autorização

### Implementação Atual (⚠️ Insegura)

- **Pseudo-auth client-side:** Validação de CNPJ com `cnpj.length >= 14`
- **Sem tokens:** Nenhum JWT, OAuth, ou API key
- **Sem sessão:** `authenticated` é um `useState` — reset ao recarregar a página
- **Sem proteção de rota:** Toda a lógica de auth está no client

📍 `src/App.tsx:8-14`

### Recomendações

| Prioridade | Ação |
|---|---|
| P0 | Implementar OAuth 2.0 com PKCE (Authorization Code Flow) |
| P0 | Backend deve emitir e validar JWT (RS256, 15min access + 7d refresh) |
| P1 | Armazenar refresh token em httpOnly cookie (não localStorage) |
| P1 | Implementar CSRF protection (double-submit cookie pattern) |

## 3. Comunicação com o Backend

### Riscos Identificados

| Risco | Severidade | Descrição |
|---|---|---|
| API sem autenticação | 🔴 Crítica | `POST /v1/tax/calculate` não exige token |
| Sem HTTPS | 🔴 Crítica | `http://localhost:8080` — em produção, tráfego em texto plano |
| URL hardcoded | 🟡 Média | Sem proxy reverso ou API gateway |
| Sem rate limiting | 🟡 Média | Frontend não implementa throttling de requisições |

### Recomendações

- Usar HTTPS em produção (TLS 1.3)
- Adicionar `Authorization: Bearer <jwt>` em todas as chamadas
- Configurar Vite proxy para desenvolvimento (`vite.config.ts → server.proxy`)
- Implementar exponential backoff no cliente HTTP

## 4. Segurança do Frontend (OWASP Top 10 para Client-Side)

| OWASP | Status | Observação |
|---|---|---|
| **A01: Broken Access Control** | ⚠️ Vulnerável | Auth client-side sem verificação backend |
| **A02: Cryptographic Failures** | ⚠️ Vulnerável | Sem HTTPS, sem criptografia de dados sensíveis |
| **A03: Injection** | ✅ OK | Sem `dangerouslySetInnerHTML`, sem `eval()` |
| **A04: Insecure Design** | ⚠️ Vulnerável | Pseudo-auth é inseguro por design |
| **A05: Security Misconfiguration** | ⚠️ Atenção | Nginx sem `server_tokens` (OK), mas sem CSP/HSTS |
| **A06: Vulnerable Components** | ✅ OK | `npm audit` necessário — verificar dependências |
| **A07: Auth Failures** | ⚠️ Vulnerável | Sem auth real, sem proteção contra brute force |
| **A08: Software/Data Integrity** | ✅ OK | `npm ci` com `package-lock.json` |
| **A09: Logging & Monitoring** | ⚠️ Atenção | Apenas `console.error` — sem reporting de erros |
| **A10: SSRF** | ✅ N/A | Frontend puro, sem server-side requests |

## 5. Headers de Segurança Recomendados (Nginx)

```nginx
# A adicionar ao nginx.conf antes do deploy em produção
add_header Content-Security-Policy "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline' https://cdn.jsdelivr.net; connect-src 'self' https://api.taxnexus.com;" always;
add_header X-Frame-Options "DENY" always;
add_header X-Content-Type-Options "nosniff" always;
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
add_header Permissions-Policy "camera=(), microphone=(), geolocation=()" always;
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains; preload" always;
```

## 6. Dependências — Verificação de Vulnerabilidades

### Comando

```bash
npm audit --production
```

### Status

Não executado neste ciclo de documentação. **Recomendação:** Executar antes de cada deploy e integrar ao CI/CD.

## 7. Dados Sensíveis

### Identificados no Código

| Dado | Local | Risco |
|---|---|---|
| CNPJ | `TaxRequest.cnpj` | Trafega em texto plano para o backend |
| Código IBGE | `TaxRequest.ibge` | Baixo — dado público |
| Saldo Remanescente | `TaxRequest.saldo_remanescente` | Financeiro — deve ser protegido em trânsito |
| ID Cadastro Único | `TaxResponse.callback.id_cadastro_unico` | Rastreabilidade corporativa — possível PII |

## 8. Checklist de Segurança para Produção

- [ ] Autenticação OAuth 2.0 / JWT implementada
- [ ] HTTPS com TLS 1.3 configurado
- [ ] Headers de segurança no Nginx (CSP, HSTS, X-Frame-Options)
- [ ] `npm audit` limpo (0 vulnerabilidades HIGH/CRITICAL)
- [ ] Variáveis de ambiente para URLs de API (sem hardcode)
- [ ] Rate limiting no backend
- [ ] CORS configurado para origens específicas
- [ ] Logs de erro enviados para serviço de monitoramento
- [ ] Threat model documentado e revisado
- [ ] Testes de segurança automatizados no CI/CD

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
