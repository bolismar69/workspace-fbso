# SPRINT-REVIEW: Sprint 3 — Portal Admin + Contas e Planos

- **Sprint:** 3 de 7
- **Status:** 🔄 Em andamento (iniciada em 16/07/2026)
- **Data da Review:** 31/08/2026
- **Participantes:** Time Técnico, Tech Lead, **Product Owner** 🎯
- **Features:** 8 (F01-01 a F02-05)

> 🎯 **Primeira review com demonstração para o PO.** Prepare o ambiente com dados seed para uma demo fluida. Sprint em andamento — checklist de review será preenchido na demonstração de 31/08.

---


## 🎯 O Que Demonstrar

### 1. Dashboard Administrativo (F01-01)

- [ ] **Métricas:** Cards com total de contas ativas, por status (ACTIVE, PENDING, SUSPENDED), por plano
- [ ] **Período:** Alternar entre 7d, 30d, mês atual — métricas recalculam
- [ ] **Performance:** Dashboard carrega em ≤3s
- [ ] **Soft delete:** Tenant excluído (soft) não aparece nas métricas

> 🎬 **Script:** "Temos 50 tenants. 35 ativos, 10 pendentes de onboarding, 5 suspensos. Distribuídos em 3 planos. Vou filtrar pelos últimos 30 dias..."

### 2. Lista de Contas (F01-02)

- [ ] **Paginação:** 25 registros por página. Navegar entre páginas
- [ ] **Busca textual:** Digitar "Mercado" → filtra em tempo real (≥3 caracteres)
- [ ] **Filtros:** Filtrar por status (ACTIVE, PENDING) e por plano
- [ ] **Ordenação:** Por data de criação (mais recentes primeiro)

### 3. Criar Tenant (F02-01)

- [ ] **Criação:** Preencher formulário → status PENDING_ONBOARDING
- [ ] **Validação:** Tentar criar com razão social duplicada → mensagem clara
- [ ] **E-mail:** Demonstrar envio de e-mail de ativação (mock SMTP)
- [ ] **Reenvio:** Botão "Reenviar convite" funcional

### 4. Transições de Status (F02-02)

- [ ] **Ciclo completo:** PENDING → ACTIVE → SUSPENDED (com motivo) → ACTIVE
- [ ] **Validação:** Tentar ACTIVE → PENDING → 422 com mensagem explicativa
- [ ] **Timeline:** Histórico de status visível para cada tenant

> 🎬 **Script:** "Este tenant estava PENDING. Vou ativá-lo. Agora vou suspender por 'Inadimplência'. O motivo fica registrado. Agora reativo..."

### 5. Planos (F02-03)

- [ ] **Criação:** Plano com nome, preço (R$), recorrência, módulos
- [ ] **Edição:** Alterar preço → nova versão. Assinantes existentes mantêm preço antigo
- [ ] **Desativação:** Tentar desativar plano com assinantes → bloqueado
- [ ] **Último plano:** Tentar desativar o último plano ativo → 422 (RN06-03)

### 6. Assinaturas (F02-04)

- [ ] **Contratar:** Tenant assina plano → ACTIVE
- [ ] **Duplicação:** Tentar segunda assinatura ativa → 409
- [ ] **Upgrade:** Mudar de plano Básico → Avançado. Timeline mostra histórico
- [ ] **Suspensão:** Suspender assinatura → módulos bloqueados

### 7. Auditoria (F02-05)

- [ ] **Consulta:** GET /audit com filtros de período, ação, entidade
- [ ] **Paginação:** 25 registros, ordenados por timestamp DESC
- [ ] **Imutabilidade:** Tentar DELETE → 403. Tentar UPDATE → 403

---

## 📋 Pontos de Verificação (PO)

| Verificação | Status |
|:---|:---:|
| Dashboard carrega com métricas corretas | ⬜ |
| Cards do dashboard são clicáveis (levam à lista filtrada) | ⬜ |
| Lista de tenants com paginação e busca textual | ⬜ |
| Criar tenant → status PENDING, email enviado | ⬜ |
| Suspender tenant exige motivo | ⬜ |
| Reativar tenant restaura acesso | ⬜ |
| Criar/Editar/Desativar planos | ⬜ |
| Assinar tenant, upgrade, suspender assinatura | ⬜ |
| Consultar auditoria com filtros | ⬜ |
| Mensagens de erro em PT-BR e claras | ⬜ |

---

## 📊 Métricas da Review

| Métrica | Meta | Resultado |
|:---|:---:|:---:|
| Dashboard p95 latency | ≤ 3s | |
| Endpoints REST funcionais | 17 | |
| RNs implementadas | 20 | |
| Cenários de teste passando | 55/55 | |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## ➡️ Próximo Passo

**Sprint 4 — RBAC** (31/08 → 15/09): Gestão de Usuários, Matriz de Permissões, Vinculação Usuário × Unidade × Módulo, Acesso Condicional (403).

---

🤖 *Sprint 3 iniciada em 16/07/2026. Checklist de review a ser executado na demonstração para o PO em 31/08/2026.*
