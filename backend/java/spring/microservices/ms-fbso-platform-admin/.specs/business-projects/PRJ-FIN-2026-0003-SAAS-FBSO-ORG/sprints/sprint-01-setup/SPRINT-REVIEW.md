# SPRINT-REVIEW: Sprint 1 — Setup e Fundação

- **Sprint:** 1 de 7
- **Data da Review:** 07/08/2026
- **Participantes:** Time Técnico, Tech Lead
- **PO:** Opcional (sprint de infraestrutura — sem features visíveis)

> ⚠️ Esta sprint é de **fundação técnica**. Não há interface visível para o Product Owner. A review é focada em validação técnica.

---

## 🎯 O Que Demonstrar

### 1. Build e Compilação

- [ ] Executar `mvn clean install` ao vivo
- [ ] Mostrar BUILD SUCCESS
- [ ] Mostrar JAR gerado em `target/`

### 2. Migrations Flyway

- [ ] Executar `docker compose up postgres` (se ainda não estiver rodando)
- [ ] Executar `mvn flyway:migrate`
- [ ] Conectar ao PostgreSQL e listar tabelas: `\dt fbso_platform.*`
- [ ] Mostrar as 11 tabelas criadas

```
Esperado:
 fbso_platform.tenant
 fbso_platform.plan
 fbso_platform.plan_module
 fbso_platform.subscription
 fbso_platform."user"
 fbso_platform.user_permission
 fbso_platform.resource_action
 fbso_platform.role_resource
 fbso_platform.business_unit
 fbso_platform.product_service
 fbso_platform.audit_log
```

### 3. Índices Únicos Parciais

- [ ] Mostrar índices criados (V002):
  - `unique_cnpj_active` em `business_unit`
  - `unique_email_active` em `user`
  - `unique_sku_active` em `product_service`

### 4. Estrutura de Pacotes

- [ ] Mostrar árvore de diretórios `src/main/java/com/fbso/platform/admin/`
- [ ] Confirmar 14 pacotes top-level (47 diretórios total) conforme ARCHITECTURE.md §2

### 5. Docker

- [ ] Build da imagem: `docker build -t fbso-platform-admin:latest .`
- [ ] Verificar tamanho: `docker images fbso-platform-admin`

---

## 📋 Pontos de Verificação Técnica

| Verificação | Status |
|:---|:---:|
| `mvn clean install` passa | ⬜ |
| `mvn flyway:migrate` cria 11 tabelas | ⬜ |
| `mvn compile` sem erros | ⬜ |
| `BaseRepository` testado unitariamente | ⬜ |
| JaCoCo configurado com meta 80% | ⬜ |
| Dockerfile funcional | ⬜ |
| Estrutura de pacotes segue ARCHITECTURE.md | ⬜ |

---

## 🚧 Bloqueios Identificados

| Bloqueio | Ação | Responsável |
|:---|:---|:---|
| (preencher na review) | | |

---

## ➡️ Próximo Passo

**Sprint 2 — Segurança** (07/08 → 15/08): JWT Filter, TenantContext, @RequiresPermission, @Auditable, GlobalExceptionHandler. **Esta sprint é pré-requisito crítico — nenhum endpoint de negócio é implementado antes dela.**

---


🤖 *Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): esclarecida contagem pacotes (14 top-level, 47 diretórios total).*
