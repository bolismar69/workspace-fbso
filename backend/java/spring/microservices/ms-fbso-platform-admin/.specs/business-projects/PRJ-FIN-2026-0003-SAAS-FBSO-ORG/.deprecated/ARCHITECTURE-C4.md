# ARCHITECTURE-C4.md — Modelo C4: ms-fbso-platform-admin

- Microserviço: ms-fbso-platform-admin
- Stack: Java 25 + Spring Boot + PostgreSQL
- Escopo C4: Contexto (L1), Containers (L2), Componentes (L3)
- Origem funcional: ARCHITECTURE.md e PRD.md do projeto PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- Versão: 1.1
- Data: 16 de Julho de 2026
- Complemento: [ARCHITECTURE-C4-DEPLOYMENT.md](./ARCHITECTURE-C4-DEPLOYMENT.md) — visão de infraestrutura e deployment

---

## 1. Objetivo

Este documento traduz a arquitetura atual do ms-fbso-platform-admin para o modelo C4, mantendo aderência à estrutura package-by-layer e aos mecanismos de segurança e isolamento multi-tenant já definidos.

---

## 2. C4 Level 1 — System Context

### 2.1 Visão

O sistema ms-fbso-platform-admin é o backend administrativo SaaS multi-tenant da plataforma FBSO, responsável por gestão de tenants, planos, assinaturas, usuários, permissões, unidades de negócio, produtos/serviços, onboarding, auditoria e dashboards. Ele é consumido pelo frontend `web_app-fbso-platform-portal` (SPA React/Next.js), que por sua vez serve os usuários finais (Admin FBSO, Gestor de Unidade, Auditor).

### 2.2 Diagrama (L1)

```mermaid
C4Context
  title C4 L1 - Contexto do Sistema ms-fbso-platform-admin

  Person(admin, "Admin FBSO", "Administra tenants, planos, assinaturas e governança")
  Person(manager, "Gestor de Unidade", "Opera entidades de negócio do tenant")
  Person(auditor, "Auditor", "Consulta trilhas de auditoria e conformidade")

  System(system, "ms-fbso-platform-admin", "API administrativa multi-tenant (Spring Boot)")

  System_Ext(frontend, "FBSO Platform Portal", "SPA Web (React/Next.js) — consome a Admin API")
  System_Ext(keycloak, "Keycloak", "Identity Provider para autenticação/autorização JWT")
  System_Ext(postgres, "PostgreSQL", "Banco transacional com RLS, índices e constraints")
  System_Ext(smtp, "SMTP/Email Service", "Envio de convites, ativações e notificações")
  System_Ext(obs, "Observabilidade", "Logs, métricas e monitoramento operacional")

  Rel(admin, frontend, "Usa interface administrativa", "HTTPS")
  Rel(manager, frontend, "Usa interface do tenant", "HTTPS")
  Rel(auditor, frontend, "Consulta auditoria", "HTTPS")

  Rel(frontend, system, "Consome API REST", "HTTPS/JSON")

  Rel(system, keycloak, "Valida e interpreta JWT", "HTTPS/OIDC")
  Rel(system, postgres, "Lê/escreve dados de domínio", "JDBC")
  Rel(system, smtp, "Envia e-mails transacionais", "SMTP")
  Rel(system, obs, "Publica logs e métricas", "OTel/HTTP")
```

---

## 3. C4 Level 2 — Container Diagram

### 3.1 Visão

Dentro do boundary do sistema, o principal container é uma API Spring Boot que expõe endpoints REST e encapsula as regras de negócio. O armazenamento é feito em PostgreSQL com Row-Level Security para isolamento de tenant.

### 3.2 Diagrama (L2)

```mermaid
C4Container
  title C4 L2 - Containers do ms-fbso-platform-admin

  Person(admin, "Admin FBSO")
  Person(manager, "Gestor de Unidade")
  Person(auditor, "Auditor")

  System_Ext(frontend, "FBSO Platform Portal", "SPA Web")
  System_Ext(keycloak, "Keycloak", "IdP JWT")
  System_Ext(smtp, "SMTP/Email Service", "E-mails transacionais")
  System_Ext(obs, "Observabilidade", "Logs/Métricas")

  System_Boundary(sys, "ms-fbso-platform-admin") {
    Container(api, "Admin API", "Java 25, Spring Boot", "Expõe endpoints REST e orquestra casos de uso")
    ContainerDb(db, "PostgreSQL", "PostgreSQL 17", "Persistência de domínio, auditoria, RLS multi-tenant")
  }

  Rel(admin, frontend, "Usa interface administrativa", "HTTPS")
  Rel(manager, frontend, "Usa interface do tenant", "HTTPS")
  Rel(auditor, frontend, "Consulta auditoria", "HTTPS")

  Rel(frontend, api, "Consome API REST", "HTTPS/JSON")

  Rel(api, keycloak, "Valida token e permissões", "OIDC/JWT")
  Rel(api, db, "CRUD, consultas e soft delete", "JDBC")
  Rel(api, smtp, "Envia e-mails transacionais", "SMTP")
  Rel(api, obs, "Emite logs, métricas e eventos", "HTTP/OTel")
```

---

## 4. C4 Level 3 — Component Diagram (Admin API)

### 4.1 Visão

Este nível mapeia os principais componentes internos do container Admin API, refletindo os pacotes existentes: controller, dto, service, repository, security, entity/enums, exception, config e common.

### 4.2 Diagrama (L3)

```mermaid
C4Component
  title C4 L3 - Componentes internos da Admin API

  ContainerDb(db, "PostgreSQL", "PostgreSQL 17", "RLS + dados de domínio")
  System_Ext(keycloak, "Keycloak", "Validação JWT")

  Container_Boundary(api, "Admin API (Spring Boot)") {
    Component(controller, "Controllers", "Spring Web (@RestController)", "Endpoints REST dos 11 recursos")
    Component(dto, "DTOs", "Java + Bean Validation", "Contratos de request/response")
    Component(service, "Services", "Spring Service", "Casos de uso e regras de negócio")
    Component(repository, "Repositories", "JdbcTemplate + BaseRepository", "Persistência, filtros tenant e soft delete")

    Component(jwtFilter, "JwtAuthenticationFilter", "Security Filter", "Valida token, extrai claims e contexto")
    Component(rbacAspect, "RbacAspect", "AOP", "Aplica @RequiresPermission")
    Component(auditAspect, "AuditAspect", "AOP", "Registra operações auditáveis")
    Component(annotations, "@RequiresPermission + @Auditable", "Custom Annotations", "Contratos declarativos para RBAC e Auditoria")

    Component(entityEnums, "Entity + Enums", "Modelo de domínio", "Entidades e tipos de negócio")
    Component(exception, "GlobalExceptionHandler + Exceptions", "Spring MVC Advice", "Padroniza erros (RFC 7807)")
    Component(config, "Config", "Spring Config", "Security, Web, DB, Flyway, TenantAwareDataSource")
    Component(common, "Common", "Shared Kernel", "BaseEntity, Address e utilitários comuns")
    Component(utils, "Utils", "Validators + Helpers", "CNPJ validation, JWT parsing, Date utilities")
  }

  Rel(controller, dto, "Recebe/retorna")
  Rel(controller, service, "Invoca casos de uso")
  Rel(service, entityEnums, "Manipula domínio")
  Rel(service, repository, "Persiste e consulta")
  Rel(repository, db, "Executa SQL", "JDBC")

  Rel(jwtFilter, keycloak, "Valida JWT", "OIDC/JWT")
  Rel(jwtFilter, config, "Propaga tenant_id para sessão DB")
  Rel(config, db, "Set app.current_tenant_id", "Session setting")

  Rel(rbacAspect, annotations, "Lê @RequiresPermission")
  Rel(rbacAspect, controller, "Intercepta métodos com @RequiresPermission")
  Rel(auditAspect, annotations, "Lê @Auditable")
  Rel(auditAspect, service, "Intercepta métodos com @Auditable")
  Rel(auditAspect, repository, "Grava trilha de auditoria")

  Rel(controller, exception, "Mapeia erros de negócio e segurança")
  Rel(config, controller, "Configura pipeline HTTP")
  Rel(config, jwtFilter, "Registra cadeia de segurança")
  Rel(common, service, "Tipos e utilitários compartilhados")
  Rel(common, repository, "Base classes compartilhadas")
  Rel(utils, jwtFilter, "JWT parsing helpers")
  Rel(utils, service, "CNPJ validation, Date utilities")
  Rel(utils, controller, "Bean Validation helpers")
```

---

## 5. Mapeamento do Desenho Original para C4

| Bloco no desenho atual | Representação C4 |
|:---|:---|
| controller/ + dto/ | Componentes Controllers e DTOs (L3) |
| service/ | Componente Services (L3) |
| repository/ (JDBC Template) | Componente Repositories (L3) |
| security/ (JWT Filter, RbacAspect, AuditAspect) | Componentes JwtAuthenticationFilter, RbacAspect, AuditAspect (L3) |
| security/annotation/ (@RequiresPermission, @Auditable) | Componente @RequiresPermission + @Auditable (L3) |
| entity/ + enums/ | Componente Entity + Enums (L3) |
| PostgreSQL RLS | ContainerDb PostgreSQL e relação Config -> DB (L2/L3) |
| config/ (Security, Web, DB, Flyway, TenantAwareDataSource) | Componente Config (L3) |
| exception/ | Componente Exception Handler (L3) |
| common/ (BaseEntity, Address) | Componente Common (L3) |
| utils/ (CnpjValidator, JwtUtils, DateUtils) | Componente Utils (L3) |

---

## 6. Notas de Evolução

- Se surgirem novos módulos (ex.: billing, notificações, integrações fiscais), recomenda-se expandir com C4 L3 por bounded context.
- Quando houver múltiplos deployables (ex.: APIs separadas por domínio), adicionar C4 L2 com novos containers internos.
- Para operações e ambiente, complementar com [ARCHITECTURE-C4-DEPLOYMENT.md](./ARCHITECTURE-C4-DEPLOYMENT.md).

---

## 7. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.1 | 16/07/2026 | Revisão de alinhamento C4: Adicionado Frontend (`web_app-fbso-platform-portal`) e SMTP/Email Service como sistemas externos no L1/L2. L3: Adicionados componentes `security/annotation/` (@RequiresPermission, @Auditable) e `utils/` (CnpjValidator, JwtUtils, DateUtils). Corrigida posição do TenantAwareDataSource (config/, não security/). Detalhado Config com WebConfig e TenantAwareDataSource. Atualizada tabela de mapeamento (§5) com annotations e utils. Adicionado cross-reference para C4-DEPLOYMENT.md e changelog. | Arquiteto/IA |
| 1.0 | 16/07/2026 | Criação inicial: C4 L1 (Contexto), L2 (Containers), L3 (Componentes), tabela de mapeamento, notas de evolução. | Time Técnico |
