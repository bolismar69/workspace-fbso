# Sprint 00 — Setup & Fundação: Perguntas em Aberto

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data:** 27 de Julho de 2026
- **Status:** ✅ 8/8 perguntas respondidas — Sprint 00 em andamento
- **Origem:** Planejamento técnico original — 8 perguntas herdadas + 7 itens de setup

---

## 1. Perguntas Técnicas em Aberto (8 — Herdadas do Planejamento)

Estas 8 perguntas foram identificadas no planejamento técnico original (§9) com prazo "Sprint 0 (F0)". Precisam ser respondidas antes do início do desenvolvimento.

| # | Pergunta | Responsável Sugerido | Resposta |
|---|----------|----------------------|----------|
| 1 | **Banco de dados — naming:** Qual a convenção de nomenclatura para tabelas e colunas? `snake_case`? Prefixos por módulo? | Carlos (DB) | a convenção para banco de dados esta no arquivo 'PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION' na sessçao "### 2.3 Banco de Dados (PostgreSQL — S03)" |
| 2 | **Migrations:** Flyway ou Liquibase? Versionamento semântico ou sequencial? | Carlos (DB) + Francisco (TL) | Flyway, versionamento sequencial. |
| 3 | **OpenAPI — codegen:** Gerar interfaces Java (Spring) e tipos TypeScript a partir do OpenAPI YAML? Qual ferramenta? | Francisco (TL) + Bolismar (FS) | Vamos usar o `openapi-generator-cli` ( Java: `npx @openapitools/openapi-generator-cli generate -i openapi.yaml -g java -o ./generated-java` e depois usando o plug oficial `openapi-generator-maven-plugin` no `pom.xml`, e depois executando um `mvn clean compile` vamos gerar automaticamente as interfaces do Controller (Server Stub) e os Models ) (TypeScript: `npx @openapitools/openapi-generator-cli generate -i openapi.yaml -g typescript-fetch -o ./generated-ts` , depois usar o `npm install -D openapi-typescript typescript` , e por fim adicionar no `package.json` a entrada `"generate-api": "openapi-typescript ./openapi.yaml -o ./src/types/api.d.ts"` ) |
| 4 | **Testes — cobertura mínima:** Qual o percentual de cobertura esperado (JaCoCo para backend, Vitest/Istanbul para frontend)? | Felipe (QA) + Francisco (TL) | >= 85% |
| 5 | **CI/CD:** GitHub Actions, GitLab CI, ou Jenkins? Pipeline de deploy para staging e produção? | Davi (DevOps) | GitHub Actions. Sim.  |
| 6 | **Monitoramento:** Ferramenta de observabilidade (APM)? Log aggregation? Alertas? | Davi (DevOps) | APM=OpenTelemetry + Grafana. Log Agreggation=Loki. Alertas=Alert-Manager |
| 7 | **Multi-Tenant — Connection Pool:** Estratégia de connection pooling para banco compartilhado entre tenants? | Carlos (DB) + Francisco (TL) | Pool unico compartilhado (HikariCP, `maximum-pool-size=20`). Sem segregação por tenant - o isolamento é garantido por RLS, não pelo pool. |
| 8 | **Keycloak — Alta disponibilidade:** Keycloak em cluster ou instância única no MVP? Plano de failover? | Gertrudes (IAM) + Davi (DevOps) | Instancia unica no MVP |

---

## 2. Entregáveis do Sprint 00 (7 — Setup & Fundação)

Itens de infraestrutura e scaffold que devem ser concluídos antes do Sprint 01 (desenvolvimento de features).

### 2.1 OpenAPI YAML (`fbso-platform-api.yaml`)

**Responsável:** Tech Lead Backend (Francisco)
**Revisor:** Tech Lead Frontend (Bolismar até 01/11, depois Tom)

- [ ] Gerar especificação OpenAPI 3.0+ com os 11 recursos definidos no [SPECS-DEFINITION.md §3.6](./PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md)
- [ ] Validar schemas de request/response para cada endpoint
- [ ] Definir exemplos de request/response para a coleção Postman/Bruno
- [ ] Localização: `backend/java/spring/microservices/ms-fbso-platform-admin/.specs/api/fbso-platform-api.yaml`
- [ ] Cópia de consumo: `frontend/javascript/react/web_apps/web_app-fbso-platform-portal/.specs/api/fbso-platform-api.yaml`

**Critério de aceitação:** OpenAPI YAML válido, aprovado por ambos os times (backend + frontend).

**RESPOSTA FRANCISCO OLIVEIRA ->** o arquivo `backend/java/spring/microservices/ms-fbso-platform-admin/.specs/api/fbso-platform-api.yaml` ainda não existe mas boa parte da documentação para cria-lo e seu código ja estão concentrados na pasta da solução técnica ``backend/java/spring/microservices/ms-fbso-platform-admin/`.

---

### 2.2 Schema PostgreSQL + Migrations Flyway

**Responsável:** Carlos (DB) + Francisco (TL)
**Referência:** [PRD-DEFINITION.md §2.4](./PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md) (ERD com 10 entidades Core)

- [ ] Criar migrations `V001__create_tenants_table.sql` até `V00N` para as 10 entidades da Fase 0
- [ ] Criar migrations reversas (`U001` a `U00N`) para rollback
- [ ] Implementar índices únicos parciais com Soft Delete (CNPJ, email, SKU)
- [ ] Configurar políticas RLS com `FORCE ROW LEVEL SECURITY`
- [ ] Definir schemas: `public`, `fbso_portal` (RLS), `keycloak`
- [ ] Criar usuários de banco por sistema (app, keycloak, grafana)
- [ ] Localização: `backend/java/spring/microservices/ms-fbso-platform-admin/src/main/resources/db/migration/`

**RESPOSTA FRANCISCO OLIVEIRA ->** as migrações vamos concentar na pasta `/home/bolismar/work/workspace-fbso/data_engineering/databases/db-postgresql/schema_fbso_platform/`. Ja exitem migrations desenvolvidas na pasta `backend/java/spring/microservices/ms-fbso-platform-admin/src/main/resources/db/migration/` porem vamos move-las para a pasta correta no monorepo para armazenar objetos de banco de dados.

**Critério de aceitação:** `docker compose up` aplica todas as migrations sem erro. Schemas e tabelas criados conforme ERD.

---

### 2.3 Scaffold Backend (`ms-fbso-platform-admin`)

**Responsável:** Francisco (TL) + Bolismar (FS)
**Referência:** [SPECS-DEFINITION.md §6.1](./PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) (estrutura de pacotes)

- [ ] Criar projeto Maven com Spring Boot 3.5.14 + Java 25 LTS
- [ ] Configurar `pom.xml` com dependências: Spring Web, Spring Security, Spring Data JDBC, Flyway, PostgreSQL Driver, OTel, Micrometer, Lombok, JUnit 5, Testcontainers
- [ ] Criar estrutura de pacotes `com.fbso.platform.admin.*` conforme SPECS-DEFINITION §6.1
- [ ] Criar `Dockerfile` (GraalVM Native Image) + `Dockerfile.jvm` (fallback)
- [ ] Configurar `application.yml` com profiles (dev, staging, prod)
- [ ] Configurar Maven Wrapper (`./mvnw`)
- [ ] Localização: `backend/java/spring/microservices/ms-fbso-platform-admin/`

**Critério de aceitação:** `./mvnw clean compile` sem erros. Projeto compila e inicializa com Spring Boot.

**RESPOSTA FRANCISCO OLIVEIRA ->** o projeto ja esta criado em `backend/java/spring/microservices/ms-fbso-platform-admin/` com grande parte do desenvolvimento realizado.

---

### 2.4 Scaffold Frontend (`web_app-fbso-platform-portal`)

**Responsável:** Bolismar (FS)
**Referência:** [SPECS-DEFINITION.md §6.2](./PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) (estrutura de rotas)

- [ ] Criar projeto Next.js 15 + React 19 + TypeScript 5.7+
- [ ] Configurar Tailwind CSS 4 + Zustand 5 + SWR 2 + Zod 3
- [ ] Criar estrutura de rotas conforme SPECS-DEFINITION §6.2: `(auth)`, `(onboarding)`, `(admin)`, `(portal)`
- [ ] Criar estrutura de componentes: `layout/`, `dashboard/`, `common/`
- [ ] Configurar `lib/auth.ts` (integração Keycloak), `lib/api-client.ts`, `lib/permissions.ts`
- [ ] Configurar ESLint com `jsx-a11y` + `strict: true` no `tsconfig.json`
- [ ] Localização: `frontend/javascript/react/web_apps/web_app-fbso-platform-portal/`

**Critério de aceitação:** `npm run dev` inicia o servidor Next.js. Páginas placeholder renderizam sem erro.

**RESPOSTA FRANCISCO OLIVEIRA ->** o projeto do frontend tem apenas a pasta do projeto criada `frontend/javascript/react/web_apps/web_app-fbso-platform-portal/`

---

### 2.5 Docker Compose (Ambiente Dev Local)

**Responsável:** Davi (DevOps)
**Referência:** [ARCHITECTURE-DEFINITION.md §5.1](./PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md) (topologia dev)

- [ ] Criar `docker-compose.yml` com serviços: PostgreSQL 17 Alpine, Keycloak 26.0, MailHog
- [ ] Criar `realm-config.json` com realm `fbso-admin`, client `fbso-portal`, roles e protocol mappers
- [ ] Configurar rede `fbso-network` (bridge)
- [ ] Configurar volumes para persistência de dados (PGDATA)
- [ ] Variáveis de ambiente via `.env` (nunca hardcoded)
- [ ] Localização: `infra/docker/`

**Critério de aceitação:** `docker compose up -d` sobe todos os serviços. Keycloak acessível em `:8081` com realm funcional. PostgreSQL com schemas criados.

**RESPOSTA FRANCISCO OLIVEIRA ->** arquivos `docker-compose.yml`, `realm-config.json` foram criados na pasta `backend/java/spring/microservices/ms-fbso-platform-admin/`

---

### 2.6 MSW Mock (Frontend)

**Responsável:** Bolismar (FS)
**Dependência:** OpenAPI YAML (item 2.1) concluído

- [ ] Configurar MSW 2 no projeto frontend
- [ ] Gerar handlers a partir do OpenAPI YAML (codegen ou manual)
- [ ] Cobrir endpoints dos épicos EP-0001 (Dashboard Admin) e EP-0002 (Clientes e Assinaturas) — prioridade para Sprint 01
- [ ] Dados mock realistas (massa de dados fornecida pelo BA Mauro)
- [ ] Localização: `frontend/javascript/react/web_apps/web_app-fbso-platform-portal/src/mocks/handlers/`

**Critério de aceitação:** Frontend renderiza dados mock para endpoints EP-0001 e EP-0002 sem backend em execução.

**RESPOSTA FRANCISCO OLIVEIRA ->** o projeto do frontend tem apenas a pasta do projeto criada `frontend/javascript/react/web_apps/web_app-fbso-platform-portal/`

---

### 2.7 Responder às 8 Perguntas em Aberto

**Responsável:** Time completo (ver Seção 1)
**Prazo:** Antes do início do Sprint 01

- [x] Cada pergunta respondida com decisão e justificativa
- [x] Decisões registradas como ADRs quando aplicável
- [x] Respostas atualizadas neste documento (Seção 1, coluna "Resposta")

**Critério de aceitação:** 8/8 perguntas respondidas. Decisões comunicadas ao time.

---

## 3. Sequenciamento Revisado (baseado no estado real do projeto)

O backend (`ms-fbso-platform-admin`) já possui código, migrations e Docker Compose. O frontend está no ponto zero. O sequenciamento foi ajustado para refletir essa realidade.

```
Dia 1:   2.7 Auditar backend existente — levantar gaps vs. SPECS-DEFINITION §6.1
         ──▶ Mapear: o que já está implementado, o que falta, o que precisa ser ajustado

Dia 1-2: 2.5 Mover Docker Compose de backend/ → infra/docker/
         ──▶ Atualizar paths e variáveis. Validar com `docker compose up -d`

Dia 2:   2.2 Mover migrations de backend/ → data_engineering/databases/db-postgresql/schema_fbso_platform/
         ──▶ Atualizar `spring.flyway.locations` no application.yml

Dia 2-3: 2.1 Extrair OpenAPI YAML dos controllers Spring existentes
         ──▶ SpringDoc ou extração manual. Validar contra SPECS-DEFINITION §3.6

Dia 3-5: 2.4 Scaffold Frontend (Next.js 15 + React 19 + Tailwind)
         ──▶ Estrutura de rotas conforme SPECS-DEFINITION §6.2

Dia 5-6: 2.6 Configurar MSW Mock baseado no OpenAPI YAML
         ──▶ Cobrir EP-0001 e EP-0002

Dia 6-7: Validação integrada: docker compose up → backend compila → frontend renderiza mock
```

### Estado Atual por Entregável

| Item | Estado | Localização Atual | Localização Alvo |
|---|---|---|---|
| **2.1** OpenAPI | 🔴 Não existe | — | `backend/.../.specs/api/` |
| **2.2** Migrations | 🟡 Existem, serão movidas | `backend/.../db/migration/` | `data_engineering/databases/db-postgresql/schema_fbso_platform/` |
| **2.3** Backend | 🟢 Código existente | `backend/.../ms-fbso-platform-admin/` | ✅ Localização correta |
| **2.4** Frontend | 🔴 Apenas pasta | `frontend/.../web_app-fbso-platform-portal/` | ✅ Localização correta |
| **2.5** Docker Compose | 🟡 Existe, será movido | `backend/.../ms-fbso-platform-admin/` | `infra/docker/` |
| **2.6** MSW Mock | 🔴 Não existe | — | `frontend/.../src/mocks/handlers/` |
| **2.7** Perguntas | 🟢 8/8 respondidas | — | — |

---

## 4. Referências

| Documento | Relevância para o Sprint 00 |
|---|---|
| [SPECS-DEFINITION.md](./PROJECT-TECHNICAL-DEFINITIONS-SPECS-DEFINITION.md) | Convenções de código, API, DB que os scaffolds devem seguir |
| [PRD-DEFINITION.md](./PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md) | ERD (§2.4) — base para migrations |
| [ARCHITECTURE-DEFINITION.md](./PROJECT-TECHNICAL-DEFINITIONS-ARCHITECTURE-DEFINITION.md) | Topologia dev (§5.1) — base para Docker Compose |
| [STACK-MATRIX.md](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Versões exatas de cada tecnologia |
| [REPOSITORY-STRUCTURE.md](./PROJECT-TECHNICAL-DEFINITIONS-REPOSITORY-STRUCTURE.md) | Localização de cada artefato no workspace |
| [MILESTONES.md](./PROJECT-TECHNICAL-DEFINITIONS-MILESTONES.md) | Roadmap M1-M7 com datas-alvo |
| [TEAM-MAP.md](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md) | Quem faz o quê — matriz de skills |
| [TEAM-CAPACITY.md](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md) | Disponibilidade de horas por profissional |
