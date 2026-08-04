# Manuais de Usuário: Plataforma Shield — Integração para Times de Produto
## [STATUS: Em revisão]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 03-SRS |
| **Solução Técnica** | ms-shield-identity-auth |
| **Data** | 03/08/2026 | **Versão** | 2.0 — Revisão Integração | **Metodologia** | WATERFALL |

---

## 1. Getting Started — Para Times de Produto

A Plataforma Shield gerencia autenticação de forma **transparente para o frontend**. Seu produto **não precisa implementar fluxo de login** — apenas fazer chamadas de API normalmente.

### O que seu frontend SPA precisa fazer:

1. **Nada relacionado a autenticação.** A SPA faz chamadas de API como `GET /api/v1/alunos` sem se preocupar com sessão
2. **Enviar cookies automaticamente** — usar `credentials: 'include'` no fetch/axios
3. **Tratar redirect 302** — se a API retornar 302, o navegador segue o redirect para o Keycloak (tela de login). Após autenticação, o navegador volta para a URL original
4. **NUNCA tentar ler ou armazenar tokens JWT** — o JWT é injetado pelo Kong no header `Authorization` da requisição interna. O frontend não tem acesso a ele

### Checklist de Integração (4 passos)

| Passo | Ação | Responsável |
|-------|------|------------|
| 1 | Configurar fetch/axios com `credentials: 'include'` | Dev Frontend |
| 2 | Tratar respostas 302 como fluxo normal (navegador segue redirect) | Dev Frontend |
| 3 | Testar: sem cookie → API redireciona para login → autenticar → API retorna dados | QA + Dev Frontend |
| 4 | Verificar: `document.cookie` NÃO mostra tokens (confirmação de segurança) | QA |

### O que seu produto NÃO deve fazer:

- ❌ Redirecionar manualmente para `/auth/login`
- ❌ Chamar `/auth/me` para obter perfil (as claims estão no JWT injetado)
- ❌ Chamar `/auth/logout` (basta remover o cookie local; a sessão expira no Redis)
- ❌ Armazenar client_secret do Keycloak (o Shield é o único que conhece)
- ❌ Tentar extrair JWT de cookies ou headers de resposta

---

## 2. Fluxo de Autenticação (visão do frontend)

```
1. SPA carrega → faz GET /api/v1/alunos
2. Se resposta = 200 com dados → usuário está autenticado (transparente)
3. Se resposta = 302 → navegador segue para tela de login Keycloak
4. Usuário autentica → navegador volta para a URL original
5. SPA recarrega → faz GET /api/v1/alunos → 200 com dados
```

**O frontend não implementa nenhuma lógica de autenticação.** Apenas faz chamadas de API e deixa o navegador seguir redirects.

---

## 3. FAQ

| Pergunta | Resposta |
|----------|---------|
| Preciso implementar tela de login? | Não. O Keycloak fornece a tela de login com o tema visual da escola |
| Como obtenho os dados do usuário (nome, email, roles)? | As claims do JWT (tenant_id, roles, user_id, email) são injetadas no header `Authorization`. Seu backend as extrai do token |
| Preciso de client_secret do Keycloak? | Não. O Shield é o único componente que conhece esse segredo |
| O que acontece quando a sessão expira? | A API retorna 302 → Keycloak. Se ainda houver refresh token válido, a renovação é silenciosa (usuário não percebe) |
| Funciona em múltiplas abas? | Sim. O cookie SHIELD_SESSION é compartilhado entre abas do mesmo domínio |

---

**[STATUS: SUCESSO]** — Manual alinhado com arquitetura Kong Filter. Times de produto integram em 4 passos, sem implementar lógica de autenticação.
