# SPRINT-REVIEW: Sprint 5 — Portal do Cliente e Onboarding

- **Sprint:** 5 de 7
- **Data da Review:** 30/09/2026
- **Participantes:** Time Técnico, Tech Lead, **Product Owner** 🎯
- **Features:** 4 (F04-01 a F04-04)
- **Progresso:** Frente 0 ✅ (7/7 — 17/07/2026). Próximo: Frente 1 (10 tarefas)

---

## 🏗️ Infraestrutura (Frente 0) ✅ CONCLUÍDA

- [x] **docker-compose.yml:** PostgreSQL 17 + Keycloak 26 + MailHog — `docker compose up -d` funcional
- [x] **Keycloak realm:** `fbso-platform` com 4 roles (ADMIN_TENANT, MANAGER_BU, OPERATOR_BU, AUDITOR), client `fbso-platform-admin` (Auth Code + PKCE), 3 custom claims (tenant_id, business_unit_ids, modules)
- [x] **Dependências:** Flyway 10.22.0→12.11.0, PostgreSQL driver 42.7.10→42.7.11 (CVE-2026-42198 fix), OAuth2 Client adicionado
- [x] **SecurityConfig:** 2 SecurityFilterChain beans — OAuth2 Login (@Order(1)) + API Resource Server (@Order(2))
- [x] **application.yml:** Client registration + provider Keycloak configurados. Portas alinhadas com docker-compose (8081)
- [x] **Build:** ✅ `mvn compile` + `mvn test` (213 testes, 0 falhas)

---

## 🎯 O Que Demonstrar

### 1. Login e Autenticação (F04-01)

- [ ] **Login:** E-mail + senha → redireciona para dashboard do cliente
- [ ] **Senha incorreta:** Mensagem "Email ou senha incorretos"
- [ ] **Rate limiting:** 5 tentativas incorretas → bloqueio 15min com contagem regressiva
- [ ] **Recuperação:** "Esqueci minha senha" → email com link → redefinir → login com nova senha
- [ ] **Sessão:** 60min de inatividade → logout automático

### 2. Onboarding Guiado — Fluxo Completo (F04-02)

- [ ] **Passo 1:** Confirmar dados cadastrais (razão social, segmento) → barra 25%
- [ ] **Passo 2:** Cadastrar CNPJ da Matriz → validação CNPJ → barra 50%
- [ ] **Passo 3:** (placeholder/configuração adicional) → barra 75%
- [ ] **Passo 4 / Complete:** Finalizar → barra 100% → **tenant transita para ACTIVE**
- [ ] **Validação:** Tentar pular do passo 1 para o passo 3 → redirecionado para passo 2

> 🎬 **Script:** "Este é o primeiro acesso do tenant 'Novo Mercado'. Ele está PENDING_ONBOARDING. Vou fazer o onboarding completo: confirmar dados, cadastrar Matriz, finalizar. O tenant muda para ACTIVE e o dashboard do cliente fica disponível."

### 3. Dashboard do Cliente (F04-03)

- [ ] **Cards:** Unidades Ativas (1 — a Matriz), Produtos no Catálogo (0), Plano Contratado
- [ ] **Notificações:** Lembretes e alertas com link para ação
- [ ] **Isolamento:** Cliente Tenant-A não vê dados do Tenant-B

### 4. App Switcher e /auth/me (F04-04)

- [ ] **GET /auth/me:** Retorna nome, email, papel, tenant_id, modules[], business_unit_ids[]
- [ ] **App Switcher:** Visível mesmo com 1 módulo (exibe "FBSO Platform")
- [ ] **JWT claims:** modules[] e business_unit_ids[] presentes no token

---

## 📋 Pontos de Verificação (PO)

| Verificação | Status |
|:---|:---:|
| Login funcional (email + senha) | ⬜ |
| Senha incorreta → mensagem clara | ⬜ |
| Bloqueio após 5 tentativas (15min) | ⬜ |
| Recuperação de senha → link → redefinir | ⬜ |
| Onboarding 4 passos em ordem | ⬜ |
| Barra de progresso visível (25%→50%→75%→100%) | ⬜ |
| Primeira BU = Matriz (parent_id=NULL) | ⬜ |
| Tenant ACTIVE após conclusão | ⬜ |
| Dashboard cliente com cards | ⬜ |
| GET /auth/me → dados corretos | ⬜ |
| App Switcher visível (1 módulo placeholder) | ⬜ |
| Cliente não vê dados de outro tenant | ⬜ |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## ➡️ Próximo Passo

**Sprint 6 — Unidades de Negócio e Catálogo** (30/09 → 15/10): CRUD de BUs hierárquicas (Matriz/Filial), Catálogo de Produtos/Serviços segmentado por BU.

---

🤖 *Checklist de review da Sprint 5. O fluxo de onboarding completo é a demonstração principal.*
