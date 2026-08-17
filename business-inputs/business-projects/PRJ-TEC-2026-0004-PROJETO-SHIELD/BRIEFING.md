# 📋 Briefing de Projeto: Microserviço BFF de Autenticação SaaS (Multi-Tenant)

## 1. Visão Geral do Projeto

* **Nome do Projeto:** Auth-BFF (Backend-For-Frontend de Autenticação Multi-Tenant)
* **Objetivo:** Construir um microserviço de borda (*Edge Service*) responsável por gerenciar todo o ciclo de vida de autenticação, sessão e roteamento multi-tenant da plataforma SaaS. O serviço atuará como intermediário seguro entre os clientes (navegadores), o Keycloak (Multi-Realm) e o Kong API Gateway.
* **Impacto Operacional:** Garantir segurança máxima contra ataques XSS/CSRF, isolar domínios por cliente no Keycloak e abstrair a complexidade de autenticação dos microserviços de negócio (ex: Gestão Escolar).

---

### Elevator Pitch — Projeto Shield

- "Estamos lançando o Projeto Shield, a camada central de Identidade e Segurança do nosso ecossistema de produtos (SaaS, Comunidades de Ensino, Portal da Reforma, entre outros). Ele é responsável por reconhecer automaticamente o domínio de cada cliente, autenticar usuários com total isolamento de dados e garantir uma experiência de entrada simples, fluida e ultra-segura para seus colaboradores e alunos."

---

## 2. Arquitetura e Fluxo da Solução

```
[Cliente (Browser)]
       │
       ▼ (1. HTTPS - cliente-saas.com)
[Cloudflare] -- (Injeta Header: X-Tenant-Host)
       │
       ▼ (2. Roteamento de Entrada)
[Kong API Gateway]
       │
       ▼ (3. Orquestração de Sessão)
[Auth-BFF Service]
       ├── (4. Resolve Realm) ──► [Redis (Cache Mapeamento Host -> Realm)]
       ├── (5. OIDC Flow)     ──► [Keycloak (Multi-Realm)]
       └── (6. Retorna Cookie HttpOnly / Repassa Bearer Token com Tenant Claim)

```

---

## 3. Escopo Funcional & Requisitos Técnicos

### 3.1. Resolução Dinâmica de Tenant & Realm

* **Captura de Origem:** Capturar o header enviado pelo Cloudflare (ex: `X-Tenant-Host: escola-alfa.com` ou `faculdade-beta.com`).
* **Mapeamento de Realm:** Consultar uma camada de cache rápida (Redis) para traduzir o domínio acessado no Realm correspondente do Keycloak (ex: `escola-alfa.com` $\rightarrow$ Realm `realm-escola-alfa`).
* **Fallback & Trativa de Erros:** Retornar resposta padronizada de erro caso o tenant/domínio não seja reconhecido ou esteja inativo.

### 3.2. Fluxo de Autenticação OIDC (OpenID Connect)

* **Fluxo OAuth2:** Implementar o fluxo **Authorization Code Flow com PKCE**.
* **Endpoints a Expor no BFF:**
* `GET /auth/login`: Redireciona o usuário para o formulário de login do Realm específico no Keycloak.
* `GET /auth/callback`: Recebe o `code` do Keycloak, efetua a troca pelos tokens JWT (Access Token, ID Token, Refresh Token).
* `POST /auth/logout`: Invalida a sessão no Keycloak e limpa os cookies do navegador.
* `POST /auth/refresh`: Realiza a renovação silenciosa do Access Token utilizando o Refresh Token.
* `GET /auth/me`: Retorna os dados do perfil do usuário logado e suas permissões/roles para uso do frontend.



### 3.3. Gestão de Sessão & Segurança de Cookies

* **Ocultamento de Tokens (Token Hiding):** Os tokens JWT **nunca** devem ser expostos diretamente para o JavaScript do navegador.
* **Cookies Seguros:** O BFF deve converter o Access Token e Refresh Token em cookies com as seguintes flags ativas:
* `HttpOnly` (Imune a leitura via JS / XSS)
* `Secure` (Apenas tráfego HTTPS)
* `SameSite=Strict` ou `Lax` (Proteção contra CSRF)



---

## 4. Integração com o Kong API Gateway

### Responsabilidades do Kong:

1. **Ponto Único de Entrada (Reverse Proxy):** Receber todas as chamadas públicas da internet.
2. **Encaminhamento para o BFF:** Rotear chamadas de autenticação (`/auth/*`) diretamente para este microserviço BFF.
3. **Validação do Token em Chamadas de Negócio:**
* Nas rotas de negócio (ex: `/api/v1/escolas/*`), o Kong (ou o próprio BFF configurado como Plugin/Filter) interceptará a requisição.
* Transforma o Cookie `HttpOnly` recebido do navegador de volta para o Header `Authorization: Bearer <JWT>` antes de encaminhar a requisição para os microserviços de negócio internos na VPC.



---

## 5. Requisitos Não-Funcionais (RNFs)

* **Performance & Baixa Latência:**
* O tempo de resposta para validação de sessão/token deve ser **$< 15\text{ ms}$**.
* Utilizar cache local ou Redis para armazenar as Chaves Públicas (JWKS) do Keycloak e evitar chamadas HTTP repetitivas ao Keycloak em toda validação.


* **Escalabilidade:** O microserviço deve ser *Stateless* (sem estado em memória local), permitindo auto-scaling horizontal (HPA no Kubernetes) baseado em CPU e tráfego.
* **Observabilidade & Logs:**
* Tracing distribuído (OpenTelemetry/Jaeger).
* Logs estruturados em formato JSON contendo `correlation_id` e `tenant_id` (sem registrar senhas ou dados sensíveis/PII).



---

## 6. Entregáveis do Projeto

1. **Código Fonte do Microserviço (Auth-BFF):** Desenvolvido na stack padrão da empresa (Java/Spring Boot, Node.js/NestJS ou Go).
2. **Dockerfile & Helm Charts/Manifestos Kubernetes:** Prontos para deploy em ambiente de Staging e Production.
3. **Configuração/Declarativo do Kong Gateway:** Arquivo de rotas e plugins do Kong configurados.
4. **Documentação de API (OpenAPI/Swagger):** Documentando os endpoints públicos de autenticação.
5. **Suíte de Testes:** Testes unitários e de integração cobrindo fluxos de sucesso e falha (Realm inválido, token expirado, etc.).

---

## 7. Critérios de Aceite (Definition of Done)

* [ ] Um usuário acessando `escola-alfa.com` é redirecionado para o login do Keycloak no Realm `realm-escola-alfa`.
* [ ] Após autenticar, os tokens JWT são armazenados exclusivamente em Cookies `HttpOnly`.
* [ ] Tentativas de ler o token via `document.cookie` no frontend retornam vazio.
* [ ] O microserviço de negócio de Escolas recebe a requisição interna com o header `Authorization: Bearer <JWT>` válido e populado com o `tenant_id`.
* [ ] O Logout encerra a sessão tanto no BFF quanto no Keycloak.
