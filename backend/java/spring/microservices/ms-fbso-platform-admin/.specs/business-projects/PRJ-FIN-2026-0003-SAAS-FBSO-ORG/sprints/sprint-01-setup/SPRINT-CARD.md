# SPRINT-CARD: Sprint 1 — Setup e Fundação

- **Sprint:** 1 de 7
- **Marco:** Pre-M2 — Setup
- **Datas:** 24/07/2026 → 07/08/2026
- **Duração:** 11 dias úteis
- **Responsável:** A definir
- **Documentos-mestre:** [TASKS.md](../../TASKS.md) · [SPECS.md](../../SPECS.md) · [ARCHITECTURE.md](../../ARCHITECTURE.md)

---

## 🎯 Sprint Goal

**"Scaffold Maven compila com `mvn clean install`. 11 tabelas criadas no PostgreSQL via Flyway. Estrutura de pacotes completa compila sem erros. BaseRepository funcional com Soft Delete + Tenant Filter."**

---

## 📋 Sprint Backlog

| ID | Tarefa | Prio. | Est. | Status | Critério DONE |
|:---|:---|:---:|:---:|:---:|:---|
| **T-001** | Scaffold Maven: `pom.xml` com Spring Boot, Security, JDBC, Flyway, Testcontainers, Micrometer, Jakarta Validation | Must | 2d | ✅ | `mvn compile` BUILD SUCCESS. 13 arquivos compilados com Java 25 |
| **T-002** | `application.yml` (dev, staging, prod profiles) com datasource PostgreSQL, Keycloak JWKS URI, server port, logging | Must | 1d | ✅ | 3 profiles operacionais. Dev conecta PostgreSQL Docker. Logging configurado por profile |
| **T-003** | `Dockerfile` (GraalVM Native) + `Dockerfile.jvm` (fallback) + `.dockerignore` | Must | 1d | ✅ | Build nativo e JVM configurados. .dockerignore criado |
| **T-004** | Migration V001: schema `fbso_platform` + 11 tabelas Core com auditoria | Must | 3d | ✅ | 11 tabelas com PKs, FKs, NOT NULL, 6 campos de auditoria |
| **T-005** | Migration V002: índices únicos parciais + índices de desempenho | Must | 1d | ✅ | 4 índices parciais + 8 índices de desempenho criados |
| **T-006** | `BaseEntity.java` + `Address.java` (VO) + 8 enums (TenantStatus, TenantSegment, Recurrence, SubscriptionStatus, UserStatus, Role, TaxRegime, ProductType) | Must | 1d | ✅ | 11 tabelas compatíveis. 8 enums implementados |
| **T-007** | `BaseRepository.java` (JDBC Template): findAll, findById, save, update, softDelete com Soft Delete + Tenant Filter | Must | 2d | ✅ | CRUD injeta `WHERE deleted_dt IS NULL` e `tenant_id = ?`. 7/7 testes passando |
| **T-008** | Estrutura de pacotes: controller, service, repository, entity, dto/request, dto/response, enums, exception, security/aspect, security/annotation, config, common, utils | Must | 1d | ✅ | 47 diretórios. `mvn compile` sem erros. Segue ARCHITECTURE.md §2 |

**Total:** 8/8 tarefas concluídas · ~12 dias-homem

---

## 📦 Entregáveis da Sprint

1. **`pom.xml`** funcional com todas as dependências
2. **`application.yml`** com 3 profiles (dev, staging, prod)
3. **`Dockerfile`** + **`Dockerfile.jvm`** para builds nativo e JVM
4. **Migrations Flyway** V001 (11 tabelas) + V002 (índices)
5. **Entidades base**: `BaseEntity.java`, `Address.java`, 8 enums
6. **`BaseRepository.java`**: template JDBC com Soft Delete + Tenant Filter
7. **Estrutura de pacotes** completa com classes esqueleto

---

## ✅ Definition of Done (Sprint-Level)

- [x] `mvn clean install` executa sem erros
- [x] `mvn flyway:migrate` cria as 11 tabelas no PostgreSQL
- [x] `mvn compile` passa sem erros em todos os módulos
- [x] `BaseRepository` testado unitariamente (métodos CRUD + soft delete)
- [x] JaCoCo configurado com meta mínima 80%
- [x] `Dockerfile` gera imagem funcional
- [x] Estrutura de pacotes segue ARCHITECTURE.md §2

---

## ⚠️ Riscos e Bloqueadores

| Risco | Prob. | Impacto | Mitigação |
|:---|:---:|:---:|:---|
| Flyway migration com erro de sintaxe PostgreSQL | Média | Alto | Testar migrations em PostgreSQL Docker local antes de commitar |
| Conflito de versões no `pom.xml` (Spring Boot + dependências) | Média | Médio | Usar `spring-boot-starter-parent` como BOM. Verificar compatibilidade com Java 25 |
| `BaseRepository` com SQL genérico que não cobre casos específicos | Baixa | Alto | Implementar como abstract class com métodos protected para subclasses customizarem |
| GraalVM Native Image falha em runtime | Baixa | Baixo | `Dockerfile.jvm` como fallback. GraalVM é otimização, não bloqueante |

---

## 🔗 Dependências

- **Pré-requisitos:** PostgreSQL 17 disponível (Docker). JDK 25 instalado.
- **Sucessor:** Sprint 2 (Segurança) — depende de T-001, T-006, T-007, T-008 estarem prontos.
- **Features que dependem desta sprint:** TODAS. Nenhum endpoint funciona sem BaseRepository + migrations.

---

## 📊 Métricas da Sprint

| Métrica | Meta |
|:---|:---:|
| Tasks completadas | 8/8 |
| Cobertura JaCoCo | ≥ 80% (linhas) |
| Tempo de build (`mvn clean install`) | ≤ 30s |
| Tamanho imagem Docker (nativo) | < 200MB |

---

🤖 *Gerado a partir de TASKS.md v2.3. Revisão Caveman em 15/07/2026 (DOCS-SERVICE-SPRINTS-CAVEMAN-REVIEW.md): TASKS ref v2.0→v2.3. Documento da sprint — complementa, não substitui, os documentos-mestre.*
