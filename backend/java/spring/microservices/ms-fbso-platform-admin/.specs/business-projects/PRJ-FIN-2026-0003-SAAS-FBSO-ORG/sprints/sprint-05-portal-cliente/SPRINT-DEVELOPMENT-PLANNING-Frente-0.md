# SPRINT-DEVELOPMENT-PLANNING-Frente-0.md — Plano de Desenvolvimento: Sprint 5 — Frente 0

- **Solução:** `ms-fbso-platform-admin`
- **Projeto de Negócio:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente
- **Frente:** Frente 0 — Bloqueantes (Pré-Sprint)
- **Stack:** Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway + Caffeine
- **Data do Planejamento:** 2026-07-17
- **Origem:** [SPRINT-CARD.md](./SPRINT-CARD.md) §Frente 0 + [IDENTIFIED-TECHNICAL-DEBT](./IDENTIFIED-TECHNICAL-DEBT-sprint-05-portal-cliente.md)

---

## 1. Visão Geral

| Métrica | Valor |
|:---|---|
| **Sprint Goal** | "Ambiente de desenvolvimento local funcional com Keycloak OIDC + dependências atualizadas e seguras" |
| **Tasks planejadas** | 6 (4 efetivas + 2 NO-OP) |
| **Tasks efetivas** | T-133.DT-095, T-134.DT-045, T-135.DT-068, T-137.DT-099, T-138.DT-100 |
| **Tasks NO-OP** | T-136.DT-096, T-141.DT-098 (já implementadas no código atual) |
| **Ordem de execução** | Sequencial com paralelismo interno |
| **Estimativa total** | ~6h (≈1 dia) |
| **Stack** | Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26 + Flyway 12.11.0 (target) + Caffeine 3.2.4 |

### Descoberta Pré-Planejamento

> ⚠️ **Importante:** Durante a exploração do código para o planejamento, descobriu-se que **DT-096 e DT-098 já estão implementados** no código atual. O `JwtAuthenticationFilter.doFilterInternal()` já extrai `business_unit_ids[]` e `modules[]` via `JwtUtils`, e o `TenantContext` já armazena e expõe esses campos. Estas tasks são marcadas como **NO-OP** — apenas verificação e documentação.

---

## 2. Dependências entre Tasks

```
T-133.DT-095 (docker-compose.yml)     ← Sem dependências — executar primeiro
    │
    └── T-134.DT-045 (Flyway bump)    ← Independente de T-133 (só precisa do pom.xml)
    └── T-135.DT-068 (PG driver bump) ← Independente de T-133 e T-134
         │
         └── T-137.DT-099 (OAuth2 Client) ← Depende de build passar (T-134 + T-135)
              │
              └── T-138.DT-100 (OAuth2 YAML) ← Depende de T-137 (SecurityConfig)
    
T-136.DT-096 (JWT claims)             ← NO-OP — verificar e documentar
T-141.DT-098 (TenantContext consumo)  ← NO-OP — verificar e documentar
```

**Estratégia de paralelismo:**
- T-134 e T-135 podem ser executadas em paralelo (editam o mesmo arquivo pom.xml, mas seções diferentes)
- T-137 e T-138 são sequenciais (T-138 completa a configuração iniciada em T-137)

---

## 3. Plano por Task

### T-133.DT-095 — Criar docker-compose.yml

- **Débito:** DT-095 — Sem docker-compose, desenvolvimento local inviável
- **Critério DONE:** `docker compose up -d` funcional. Keycloak admin acessível em `localhost:8081`. MailHog em `localhost:8025`. PostgreSQL em `localhost:5432`
- **Estimativa:** 3h
- **Dependências:** Nenhuma
- **Skills aplicáveis:** `110-java-maven-best-practices`

#### Abordagem

Criar `docker-compose.yml` na raiz do projeto com 3 serviços:

| Serviço | Imagem | Portas | Propósito |
|:---|:---|:---|:---|
| **postgres** | `postgres:17-alpine` | `5432:5432` | Banco de dados com schema `fbso_platform` |
| **keycloak** | `quay.io/keycloak/keycloak:26.0` | `8081:8080` | Identity Provider — realm `fbso-platform` |
| **mailhog** | `mailhog/mailhog:v1.0.1` | `1025:1025`, `8025:8025` | Captura de emails (dev) |

Configurações críticas:
- PostgreSQL: criar database `fbso_platform`, user `fbso_admin`/`fbso_admin`
- Keycloak: modo `start-dev`, admin `admin`/`admin`, health check
- MailHog: SMTP em 1025, UI em 8025
- Volume para dados PostgreSQL persistirem entre restarts
- Network `fbso-network` compartilhada
- Health checks com `depends_on` condition `service_healthy`

#### Arquivos a criar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `docker-compose.yml` | 🆕 | Compose file com 3 serviços + network + volumes |
| `keycloak/realm-config.json` | 🆕 | Configuração do realm `fbso-platform` para importação automática |

#### Estrutura do `realm-config.json`

O realm `fbso-platform` deve incluir:
- **Realm:** `fbso-platform`, enabled: true
- **Client:** `fbso-platform-admin` (confidential, Authorization Code Flow, PKCE)
  - `client-id`: `fbso-platform-admin`
  - `redirect-uris`: `["http://localhost:3000/*", "http://localhost:8080/*"]`
  - `web-origins`: `["+"]`
- **Roles de realm:** `ADMIN_TENANT`, `MANAGER_BU`, `OPERATOR_BU`, `AUDITOR`
- **Client scope:** `fbso-custom` com mapeadores para claims:
  - `tenant_id` (String → mapeado de atributo de usuário `tenant_id`)
  - `business_unit_ids` (String → mapeado de atributo de usuário `business_unit_ids`)
  - `modules` (String → mapeado de atributo de usuário `modules`)
- **Usuário de teste:** `admin@fbso.org` / `admin` com role `ADMIN_TENANT` e atributos preenchidos

#### Riscos

| Risco | Mitigação |
|:---|:---|
| Realm JSON muito grande (>1000 linhas) | Exportar do Keycloak após criar manualmente via UI, depois versionar o JSON |
| Porta 8080 conflita com a aplicação Spring Boot | Keycloak mapeado para 8081 (host), 8080 (container interno) |
| MailHog não captura emails enviados para produção | Configurar `spring.mail.host=localhost` e `spring.mail.port=1025` no profile dev |

---

### T-134.DT-045 — Bump Flyway 10.22.0 → 12.11.0

- **Débito:** DT-045 (backlog Sprint 3) — Flyway 2 majors atrás
- **Critério DONE:** `mvn flyway:migrate` sucesso. V001-V006 reaplicadas sem erro
- **Estimativa:** 1h
- **Dependências:** Nenhuma (só precisa do pom.xml)
- **Skills aplicáveis:** `110-java-maven-best-practices`, `111-java-maven-dependencies`

#### Abordagem

Alterar 1 property no `pom.xml`:

```xml
<!-- ANTES -->
<flyway.version>10.22.0</flyway.version>

<!-- DEPOIS -->
<flyway.version>12.11.0</flyway.version>
```

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `pom.xml` | 🔄 | Alterar `<flyway.version>` de `10.22.0` para `12.11.0` |

#### Verificação

```bash
mvn flyway:migrate    # Deve aplicar V001-V006 na nova versão sem erro
mvn test              # Testes existentes devem continuar passando
```

#### Riscos

| Risco | Mitigação |
|:---|:---|
| Breaking changes do Flyway 11→12 | Verificar changelog do Flyway. Migrações existentes (V001-V006) usam SQL padrão — improvável quebrar |
| `flyway-database-postgresql` também precisa de bump | Já usa `${flyway.version}` — bump automático |

---

### T-135.DT-068 — Bump PostgreSQL Driver 42.7.10 → 42.7.11

- **Débito:** DT-068 (backlog Sprint 4) — CVE-2026-42198, CVSS 7.5 (DoS)
- **Critério DONE:** `mvn dependency:tree` confirma versão 42.7.11. Build e testes passam
- **Estimativa:** 0.5h
- **Dependências:** Nenhuma
- **Skills aplicáveis:** `110-java-maven-best-practices`, `security-review`

#### Abordagem

O PostgreSQL driver é gerenciado pelo Spring Boot Parent (3.5.14). Para override, adicionar property explícita no `pom.xml`:

```xml
<properties>
    <!-- ... existing properties ... -->
    <postgresql.version>42.7.11</postgresql.version>  <!-- override CVE-2026-42198 -->
</properties>
```

> **Nota:** O Jackson já segue o mesmo padrão de override: `<jackson.version>2.21.4</jackson.version>` no pom.xml. Seguiremos a mesma convenção.

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `pom.xml` | 🔄 | Adicionar `<postgresql.version>42.7.11</postgresql.version>` |

#### Verificação

```bash
mvn dependency:tree | grep postgresql    # Deve mostrar 42.7.11
mvn test                                  # Testes de integração com PostgreSQL devem passar
```

#### Riscos

| Risco | Mitigação |
|:---|:---|
| Mudança de comportamento entre 42.7.10 e 42.7.11 | Patch version — apenas correções de segurança. Baixíssimo risco |
| Incompatibilidade com PostgreSQL 17 | Driver 42.7.x é compatível com PostgreSQL 12-17 |

---

### T-136.DT-096 — Atualizar JwtAuthenticationFilter (claims) — NO-OP

- **Débito:** DT-096 — JWT não extrai claims `modules[]` e `business_unit_ids[]`
- **Status:** ✅ **JÁ IMPLEMENTADO** — Nenhuma ação necessária
- **Estimativa:** 0h (verificação: 15min)
- **Dependências:** Nenhuma

#### Evidência

O código atual já implementa a extração:

```java
// JwtAuthenticationFilter.java:L88-91
var buIds = JwtUtils.getBusinessUnitIds(jwt);   // ← extrai business_unit_ids
var modules = JwtUtils.getModules(jwt);          // ← extrai modules

// JwtAuthenticationFilter.java:L95
TenantContext.set(tenantId, userId, roles, buIds, modules);  // ← popula contexto

// JwtUtils.java:L81-86
public static List<UUID> getBusinessUnitIds(Jwt jwt) {
    List<String> raw = jwt.getClaimAsStringList("business_unit_ids");  // ← claim mapeada
    if (raw == null) return List.of();
    return raw.stream().map(UUID::fromString).toList();
}

// JwtUtils.java:L91-94
public static List<String> getModules(Jwt jwt) {
    List<String> modules = jwt.getClaimAsStringList("modules");  // ← claim mapeada
    return modules != null ? modules : List.of();
}
```

#### Ação

Apenas verificar que o filtro continua funcionando após os bumps (T-134, T-135):

```bash
mvn test -pl . -Dtest="JwtAuthenticationFilterTest"  # Deve passar 6/6
```

---

### T-137.DT-099 — Adicionar spring-boot-starter-oauth2-client + SecurityConfig

- **Débito:** DT-099 — Sem OAuth2 Client, Authorization Code Flow não funciona
- **Critério DONE:** Dependência adicionada. `mvn compile` sucesso. `SecurityConfig` compila sem conflitos
- **Estimativa:** 1.5h
- **Dependências:** T-134 e T-135 (build precisa passar antes de adicionar nova dependência)
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`, `110-java-maven-best-practices`

#### Abordagem

**Passo 1 — Adicionar dependência em `pom.xml`:**

```xml
<!-- NOVO: OAuth2 Client para Authorization Code Flow (Keycloak) -->
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-oauth2-client</artifactId>
</dependency>
```

**Passo 2 — Refatorar `SecurityConfig.java`:**

O `SecurityConfig` atual só configura Resource Server (validação JWT). Para suportar também OAuth2 Client (login), precisamos:

1. **Manter** o Resource Server para validação JWT de APIs
2. **Adicionar** OAuth2 Client para o fluxo de login (Authorization Code Flow)
3. **Permitir** endpoints públicos de auth (`/login`, `/oauth2/**`, `/auth/**`)
4. **Configurar** `HttpSecurity` para ambos os cenários

A abordagem será criar **2 beans `SecurityFilterChain`** com ordens distintas:

```java
@Bean
@Order(1)
public SecurityFilterChain oauth2LoginFilterChain(HttpSecurity http) throws Exception {
    http
        .securityMatcher("/auth/**", "/login", "/oauth2/**")
        .oauth2Login(oauth2 -> oauth2
            .defaultSuccessUrl("/api/v1/auth/me", true)
        )
        .csrf(csrf -> csrf.disable())
        .sessionManagement(session -> session
            .sessionCreationPolicy(SessionCreationPolicy.IF_REQUIRED) // precisa de sessão para OAuth2
        );
    return http.build();
}

@Bean
@Order(2)
public SecurityFilterChain apiFilterChain(HttpSecurity http) throws Exception {
    // Configuração existente do Resource Server para API
    // (mesmo código do securityFilterChain atual)
}
```

> **Nota sobre OAuth2 Client vs Resource Server:** O Resource Server valida tokens JWT (stateless — APIs REST). O OAuth2 Client gerencia o fluxo de login (Authorization Code Flow → redirect Keycloak → callback → token). Ambos coexistem na mesma aplicação com `SecurityFilterChain` separados por `@Order` e `securityMatcher`.

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `pom.xml` | 🔄 | Adicionar `spring-boot-starter-oauth2-client` |
| `config/SecurityConfig.java` | 🔄 | Adicionar `oauth2LoginFilterChain` + refatorar `apiFilterChain` |

#### Verificação

```bash
mvn compile                    # Deve compilar sem erros
mvn test                       # Testes existentes devem continuar passando
```

#### Riscos

| Risco | Mitigação |
|:---|:---|
| Conflito entre Resource Server e OAuth2 Client no mesmo `SecurityFilterChain` | Usar 2 beans separados com `@Order` e `securityMatcher` |
| OAuth2 Client requer sessão (não-stateless) | Apenas para endpoints `/auth/**` e `/login`; API REST continua stateless |
| Testes existentes quebram com nova configuração de segurança | Rodar `mvn test` e corrigir `SecurityConfig` imports se necessário |

---

### T-138.DT-100 — Configurar application.yml OAuth2 Client Keycloak

- **Débito:** DT-100 — Configuração OAuth2 incompleta — login redirect não funciona
- **Critério DONE:** Login redirect para Keycloak funcional. Token JWT validado pelo Resource Server
- **Estimativa:** 1h
- **Dependências:** T-137 (SecurityConfig precisa existir com OAuth2 Client habilitado)
- **Skills aplicáveis:** `304-frameworks-spring-boot-security`

#### Abordagem

Adicionar seção `client:` ao `application.yml` existente (que já tem `resourceserver`):

```yaml
spring:
  security:
    oauth2:
      # ---- Resource Server (API — validação JWT) ----
      resourceserver:
        jwt:
          jwk-set-uri: ${KEYCLOAK_JWKS_URI:http://localhost:8081/realms/fbso-platform/protocol/openid-connect/certs}
          issuer-uri: ${KEYCLOAK_ISSUER_URI:http://localhost:8081/realms/fbso-platform}

      # ---- OAuth2 Client (Login — Authorization Code Flow) ----
      client:
        registration:
          keycloak:
            client-id: ${KEYCLOAK_CLIENT_ID:fbso-platform-admin}
            client-secret: ${KEYCLOAK_CLIENT_SECRET:changeme}
            authorization-grant-type: authorization_code
            scope: openid,profile,email
            redirect-uri: ${KEYCLOAK_REDIRECT_URI:http://localhost:8080/login/oauth2/code/keycloak}
        provider:
          keycloak:
            issuer-uri: ${KEYCLOAK_ISSUER_URI:http://localhost:8081/realms/fbso-platform}
            authorization-uri: ${KEYCLOAK_ISSUER_URI:http://localhost:8081/realms/fbso-platform/protocol/openid-connect/auth}
            token-uri: ${KEYCLOAK_ISSUER_URI:http://localhost:8081/realms/fbso-platform/protocol/openid-connect/token}
            user-info-uri: ${KEYCLOAK_ISSUER_URI:http://localhost:8081/realms/fbso-platform/protocol/openid-connect/userinfo}
            jwk-set-uri: ${KEYCLOAK_ISSUER_URI:http://localhost:8081/realms/fbso-platform/protocol/openid-connect/certs}
            user-name-attribute: preferred_username
```

> **Nota:** As portas do Keycloak mudam de `8080` → `8081` porque no docker-compose o Keycloak expõe `8081:8080`. Isso alinha o `application.yml` com o ambiente dockerizado.

#### Arquivos a modificar

| Arquivo | Tipo | Descrição |
|:---|:---|:---|
| `src/main/resources/application.yml` | 🔄 | Adicionar seção `client:` + corrigir portas de `8080` para `8081` |

#### Verificação

```bash
# Subir ambiente docker
docker compose up -d

# Verificar se Keycloak está acessível
curl http://localhost:8081/realms/fbso-platform/.well-known/openid-configuration

# Iniciar aplicação e verificar redirect
curl -v http://localhost:8080/oauth2/authorization/keycloak
# Deve retornar 302 redirect para http://localhost:8081/realms/fbso-platform/protocol/openid-connect/auth
```

#### Riscos

| Risco | Mitigação |
|:---|:---|
| Variáveis de ambiente não definidas em outros ambientes (staging/prod) | Valores default apontam para docker-compose local (dev). Staging/prod devem definir env vars |
| `client-secret` exposto no application.yml | Usar placeholder `${KEYCLOAK_CLIENT_SECRET:changeme}` — valor real via env var, nunca hardcoded |
| Porta 8081 vs 8080 inconsistente entre docs | Este documento alinha docker-compose (8081) com application.yml (8081) |

---

## 4. Ordem de Execução

| Ordem | Task | Ação | Estimativa |
|:---:|:---|:---|:---:|
| 1 | **T-133.DT-095** | Criar `docker-compose.yml` + `keycloak/realm-config.json` | 3h |
| 2a | **T-134.DT-045** | Bump Flyway 10.22.0→12.11.0 (em paralelo com 2b) | 1h |
| 2b | **T-135.DT-068** | Bump PostgreSQL driver → 42.7.11 (em paralelo com 2a) | 0.5h |
| 3 | **T-136.DT-096** | Verificar JwtAuthenticationFilter — NO-OP (já implementado) | 15min |
| 4 | **T-137.DT-099** | Adicionar OAuth2 Client + atualizar SecurityConfig | 1.5h |
| 5 | **T-138.DT-100** | Configurar application.yml OAuth2 Client | 1h |
| 6 | **T-141.DT-098** | Verificar TenantContext — NO-OP (já implementado) | 15min |

**Total:** ~6h (≈1 dia de trabalho)

---

## 5. Estratégia de Build e Verificação

### Comandos

| Fase | Comando | Esperado |
|:---|:---|:---|
| Compilação | `mvn compile` | BUILD SUCCESS |
| Testes unitários | `mvn test` | Tests run: N, Failures: 0 |
| Testes segurança | `mvn test -Dtest="JwtAuthenticationFilterTest"` | 6/6 passando |
| Verificação de dependências | `mvn dependency:tree` | Flyway 12.11.0, PG 42.7.11 |
| Subir ambiente | `docker compose up -d` | 3 containers healthy |

### Checkpoints

1. **Após T-133:** `docker compose up -d` → 3 containers saudáveis, Keycloak admin acessível
2. **Após T-134+T-135:** `mvn flyway:migrate` + `mvn test` → tudo verde
3. **Após T-136:** `JwtAuthenticationFilterTest` → 6/6 passando
4. **Após T-137+T-138:** `mvn compile` + `mvn test` → tudo verde, aplicação sobe e redireciona para Keycloak

---

## 6. Impacto em Outros Documentos

Após a conclusão da Frente 0, os seguintes documentos precisam ser atualizados:

| Documento | Impacto |
|:---|:---|
| `SPRINT-CARD.md` | Marcar T-133 a T-138 e T-141 como ✅ concluídas |
| `TASKS.md` | Atualizar status das tasks |
| `ARCHITECTURE.md` | Stack atualizado: Flyway 12.11.0, PG driver 42.7.11, OAuth2 Client |
| `PRD.md` | Status da Frente 0 concluída |
| `SPECS.md` | Referência ao docker-compose como ambiente de desenvolvimento padrão |

---

## 7. Observações

1. **DT-096 e DT-098 são NO-OP:** O código já implementa a extração de claims e o armazenamento no TenantContext. O `IDENTIFIED-TECHNICAL-DEBT` foi gerado antes da inspeção detalhada do código — os débitos foram identificados como "ausentes" quando na verdade já estavam implementados.

2. **OAuth2 Client é mínimo na Frente 0:** A implementação completa do fluxo de login (T-057, T-058) será feita na Frente 3. A Frente 0 apenas disponibiliza a dependência e a configuração básica para que a Frente 3 possa construir sobre ela.

3. **docker-compose é o deliverable mais crítico:** Sem ele, nenhum desenvolvedor consegue trabalhar nas features da Sprint 5. Deve ser a primeira task executada.

---

🤖 *Documento gerado em 2026-07-17 como parte da Fase 1 do PROMPT-EXECUTE-SPRINT-TASKS.md. Stack detectada: Java 25 + Spring Boot 3.5.14 + PostgreSQL 17 + Keycloak 26. Skills aplicáveis: 110-java-maven-best-practices, 304-frameworks-spring-boot-security, 111-java-maven-dependencies, security-review.*
