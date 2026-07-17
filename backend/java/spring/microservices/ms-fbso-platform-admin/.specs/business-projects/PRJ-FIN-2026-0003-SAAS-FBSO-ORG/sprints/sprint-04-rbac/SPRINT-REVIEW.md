# SPRINT-REVIEW: Sprint 4 — Governança de Acessos (RBAC)

- **Sprint:** 4 de 7
- **Data da Review:** 15/09/2026
- **Participantes:** Time Técnico, Tech Lead, **Product Owner** 🎯
- **Features:** 4 (F03-01 a F03-04)

---


## 🎯 O Que Demonstrar

### 1. Gestão de Usuários (F03-01)

- [ ] **Convidar:** Admin Tenant convida usuário com email → status INVITE_PENDING
- [ ] **Validação:** Tentar mesmo email no mesmo tenant → 409
- [ ] **Autodesativação:** Admin tentar desativar a si mesmo → 422 "Um administrador não pode desativar a si mesmo"
- [ ] **Lista:** Usuários exibidos com nome, email, papel, status, BUs vinculadas

### 2. Matriz RBAC — Demonstração por Papel (F03-02)

- [ ] **Login como Admin Tenant:** Acessa tudo — dashboard, tenants, planos, usuários, permissões, BUs, produtos
- [ ] **Login como Gerente BU:** Vê BUs e produtos da sua unidade. Pode criar/editar. Não vê tenants, planos, auditoria
- [ ] **Login como Operador BU:** Apenas leitura de BUs e produtos. Tentar editar → 403
- [ ] **Login como Auditor:** Apenas leitura de auditoria. Tentar criar qualquer coisa → 403

> 🎬 **Script:** "Vou fazer login como cada um dos 4 papéis. Reparem que o menu lateral muda. O Operador tenta criar um produto... 403."

### 3. Vinculação Usuário × Unidade (F03-03)

- [ ] **Atribuir:** Vincular usuário a BU-A com papel Gerente
- [ ] **Isolamento:** Usuário BU-A vê apenas produtos da BU-A
- [ ] **Tentativa de acesso:** Usuário BU-A tenta acessar produto BU-B por ID direto → 404
- [ ] **Imediato:** Alterar permissão → efeito na próxima requisição

### 4. 403 Amigável (F03-04)

- [ ] **Acesso direto:** Operador digita URL `/admin/plans/create` → 403
- [ ] **Formato:** `{"title":"Acesso negado","detail":"Você não tem permissão para acessar esta área.","status":403}`
- [ ] **Sem vazamento:** 403 (nunca 404 — não revela existência do recurso)
- [ ] **Sem stack trace:** Nenhum detalhe técnico na resposta

---

## 📋 Pontos de Verificação (PO)

| Verificação | Status |
|:---|:---:|
| Convidar usuário → email enviado | ⬜ |
| Email duplicado → bloqueado | ⬜ |
| Admin não desativa a si mesmo | ⬜ |
| Admin Tenant acessa tudo | ⬜ |
| Gerente BU edita apenas sua BU | ⬜ |
| Operador BU apenas lê | ⬜ |
| Auditor apenas lê auditoria | ⬜ |
| Usuário sem BU → não acessa | ⬜ |
| Usuário sem módulo → não acessa | ⬜ |
| 403 amigável em PT-BR | ⬜ |
| Acesso direto URL proibida → 403 (não 404) | ⬜ |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## ➡️ Próximo Passo

**Sprint 5 — Portal do Cliente** (15/09 → 30/09): Login Keycloak, Onboarding guiado 4 passos, Dashboard do cliente, App Switcher.

---

🤖 *Checklist de review da Sprint 4. O teste parametrizado da matriz RN10-01 deve estar 100% verde.*
