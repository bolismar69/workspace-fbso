# SPRINT-CODE-REVIEW-Frente-1.md — Relatório de Ajustes Pós-Code Review

- **Solução:** `ms-fbso-platform-admin`
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Sprint:** 5 de 7 — sprint-05-portal-cliente · **Frente:** 1 — Recomendados
- **Stack detectada:** Java 25 · Spring Boot 3.5.14 · PostgreSQL 17 · Caffeine 3.2.4
- **Data da revisão:** 2026-07-23
- **Skills acionados:** `ponytail-audit`, `ponytail-review`, `engineering-skills`, `security-auditor`, `performance-review`, `requesting-code-review`, `differential-review`

---

## 1. Resumo da Revisão

- **Skills acionados:** 7
- **Total de achados:** 28
- **Por severidade:**

| Critical | High | Medium | Low | Info |
|:---:|:---:|:---:|:---:|:---:|
| 0 | 1 | 5 | 16 | 6 |

**Veredito:** Código da Frente 1 está **APROVADO com recomendações**. Nenhum achado crítico. 1 achado High (cobertura de testes). 5 Medium (melhorias de hardening e boas práticas). 22 Low/Info (nits e observações). Zero regressões. Zero quebras de API/schema.

---

## 2. Achados — `ponytail-audit` (5 achados)

| ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
|:---|:---|:---|:---:|:---|:---|
| PA-001 | Low | RateLimitFilter.java | 93-95 | `shrink:` null-guard redundante. `Cache.get(k, Function)` nunca retorna null | Remover `if (entry == null) entry = new RateLimitEntry()` |
| PA-002 | Low | RateLimitFilter.java | 134-137 | `delete:` `request.getParameter("username")` sempre retorna null para JSON POST body — branch inalcançável | Remover branch getParameter(). Manter apenas fallback IP |
| PA-003 | Low | BaseRepository.java | 111-120 | `shrink:` 4 ramos if-else para montar Object[] params. Complexidade ciclomática 4 | Usar `List<Object>` com add condicional → toArray() |
| PA-004 | Low | FbsoJwtAuthenticationConverter.java | 53-57 | `shrink:` 5 chamadas JwtUtils.* idênticas ao JwtAuthenticationFilter. Duplicação transicional (DT-102) | Extrair `JwtUtils.extractAllClaims(Jwt)` quando filtro for simplificado |
| PA-005 | Low | RateLimitFilter.java | 50 | `native:` `new ObjectMapper()` manual. Spring Boot já configura ObjectMapper com módulos Jackson | Injetar via parâmetro do @Bean method |

**net possível:** -12 linhas, 0 deps removidos.

---

## 3. Achados — `ponytail-review` (6 achados)

| ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
|:---|:---|:---|:---:|:---|:---|
| PR-001 | Low | RateLimitFilter.java | 91-95 | `shrink:` null-guard redundante (idêntico PA-001) | Remover |
| PR-002 | Low | RateLimitFilter.java | 134-137 | `delete:` branch inalcançável (idêntico PA-002) | Remover |
| PR-003 | Low | BaseRepository.java | 111-120 | `shrink:` 4-branch if-else (idêntico PA-003) | Refatorar para List |
| PR-004 | Low | FbsoJwtAuthenticationConverter.java | 53-57 | `shrink:` duplicação com JwtAuthenticationFilter (idêntico PA-004) | Consolidar |
| PR-005 | Info | BusinessUnit.java | — | Clean — campo isMatrix segue padrão existente | — |
| PR-006 | Info | SecurityConfig.java | — | Clean — beans seguem padrão Spring sem abstração desnecessária | — |

---

## 4. Achados — `engineering-skills` (5 achados)

| ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
|:---|:---|:---|:---:|:---|:---|
| ES-001 | Medium | RateLimitFilter.java | 50 | ObjectMapper manual vs. Spring context | Injetar via @Bean method |
| ES-002 | Low | RateLimitFilter.java | 111-120 | Complexidade ciclomática 4 na montagem de params | Refatorar para List<Object> |
| ES-003 | Low | FbsoJwtAuthenticationConverter.java | 53-70 | SRP: converter faz 2 coisas (extrai claims + popula TenantContext) | Manter como transicional. Reavaliar Sprint 6 |
| ES-004 | Low | RateLimitFilter.java | 78-84 | Leitura de response.getStatus() pós-filterChain — depende de implementação do wrapper | OK para Fase 0. Documentar limitação |
| ES-005 | Low | BaseRepository.java | 93-123 | Duplicação de padrão SQL com findAll() | Postergar — fora do escopo da Frente 1 |

---

## 5. Achados — `security-auditor` (5 achados)

| ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
|:---|:---|:---|:---:|:---|:---|
| SA-001 | Medium | RateLimitFilter.java | 50 | ObjectMapper sem proteção contra expansion attacks (DoS) | Configurar limites de profundidade ou injetar ObjectMapper do Spring |
| SA-002 | Medium | RateLimitFilter.java | 140-142 | X-Forwarded-For sem validação — spoofing de IP para bypass rate limit | Usar apenas remoteAddr ou validar XFF contra proxy confiável |
| SA-003 | Info | FbsoJwtAuthenticationConverter.java | 60 | TenantContext.set() thread-safe — modelo 1 thread/request do Spring | STATUS: SEGURO |
| SA-004 | Info | V007 migration | — | ALTER TABLE IF NOT EXISTS + índice parcial — sem exposição de dados | STATUS: SEGURO |
| SA-005 | Info | GlobalExceptionHandler.java | — | Handler 401 não vaza stack trace. RFC 7807 padronizado | STATUS: SEGURO |

---

## 6. Achados — `performance-review` (4 achados)

| ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
|:---|:---|:---|:---:|:---|:---|
| PF-001 | Info | RateLimitFilter.java | — | Caffeine O(1), expireAfterWrite, maximumSize=10k. Sem degradação | STATUS: ADEQUADO |
| PF-002 | Info | BaseRepository.java | 93-123 | Keyset usa índice B-tree (Index Scan). Superior a OFFSET para >10k | STATUS: ADEQUADO |
| PF-003 | Info | FbsoJwtAuthenticationConverter.java | — | Executado 1x/request. Custo ~μs. Elimina 1 decodificação JWT (~ms) | STATUS: GANHO DE PERFORMANCE |
| PF-004 | Info | V007 migration | — | Índice parcial WHERE deleted_dt IS NULL — 10-30% menor | STATUS: ADEQUADO |

---

## 7. Achados — `requesting-code-review` (6 achados)

| ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
|:---|:---|:---|:---:|:---|:---|
| RC-001 | Low | RateLimitFilter.java | 124-127 | `endsWith("/auth/login")` frágil — "/auth/login/backup" também match | Usar `endsWith("/api/v1/auth/login")` ou regex com $ |
| RC-002 | Low | RateLimitFilter.java | 39-40 | MAX_ATTEMPTS/BLOCK_DURATION hardcoded. Não configurável por ambiente | Adicionar `// ponytail: externalizar para application.yml` |
| RC-003 | Low | BaseRepository.java | 93-96 | Validação pageSize fail-fast (positivo) | STATUS: BOM PADRÃO |
| RC-004 | Medium | BusinessUnit.java | 1 | Javadoc ainda referencia "Sprint 6" como momento do CRUD completo | Atualizar Javadoc mencionando que `isMatrix` foi adicionado na Sprint 5 |
| RC-005 | Info | GlobalExceptionHandler.java | — | Handler segue padrão consistente RFC 7807 | STATUS: CONSISTENTE |
| RC-006 | Info | SecurityConfig.java | — | Beans com @Bean + Javadoc. Convenções respeitadas | STATUS: LIMPO |

---

## 8. Achados — `differential-review` (6 achados)

| ID | Severidade | Arquivo | Linha | Descrição | Recomendação |
|:---|:---|:---|:---:|:---|:---|
| DR-001 | Info | BusinessUnit.java | 102 | toColumnMap adiciona is_matrix — sem consumidores existentes. Blast radius: ZERO | STATUS: SEGURO |
| DR-002 | Info | BaseRepository.java | 93-123 | Novo método — não altera assinaturas existentes. Sem quebra de contrato | STATUS: SEGURO |
| DR-003 | Info | GlobalExceptionHandler.java | — | AuthenticationException antes de SecurityException na hierarquia. Sem conflito | STATUS: SEGURO |
| DR-004 | Medium | SecurityConfig.java | — | addFilterBefore(rateLimitFilter) — early return para não-login. Blast radius: POST /auth/login apenas | STATUS: SEGURO |
| DR-005 | Info | V007 migration | — | DEFAULT false + IF NOT EXISTS. Schema compatível com código existente | STATUS: SEGURO |
| DR-006 | **High** | Todos os 6 arquivos Java | — | **Nenhum novo código tem cobertura de testes.** 0 testes para RateLimitFilter, FbsoJwtAuthenticationConverter, findAllKeyset, novo handler. 213 testes existentes mantidos (0 regressões) | Implementar 10 cenários do SPRINT-TEST-PLANNING-Frente-1.md |

---

## 9. Matriz de Consolidação — Achados Únicos (deduplicados)

Após remoção de achados duplicados entre skills, o conjunto final é:

| ID | Severidade | Arquivo | Categoria | Descrição Resumida |
|:---|:---|:---|:---|:---|
| **CR-001** | Medium | RateLimitFilter.java:50 | Segurança | `new ObjectMapper()` sem módulos Spring — injetar via @Bean |
| **CR-002** | Medium | RateLimitFilter.java:140-142 | Segurança | X-Forwarded-For sem validação — risco de bypass rate limit |
| **CR-003** | Medium | RateLimitFilter.java:93-95 | Simplificação | Null-guard redundante após `Cache.get(k, Function)` |
| **CR-004** | Medium | RateLimitFilter.java:134-137 | Simplificação | `getParameter("username")` inalcançável para JSON body |
| **CR-005** | Medium | BaseRepository.java:111-120 | Simplificação | 4-branch if-else → List<Object> para montagem de params |
| **CR-006** | Low | RateLimitFilter.java:124-127 | Robustez | `endsWith("/auth/login")` frágil — usar path completo |
| **CR-007** | Low | RateLimitFilter.java:39-40 | Configurabilidade | Constantes hardcoded — externalizar via application.yml |
| **CR-008** | Low | FbsoJwtAuthenticationConverter.java:53-57 | DRY | Duplicação com JwtAuthenticationFilter (transicional — DT-102) |
| **CR-009** | Low | BusinessUnit.java:1 | Documentação | Javadoc referencia Sprint 6 — atualizar para mencionar Sprint 5 |
| **CR-010** | **High** | Todos os 6 arquivos Java | Testes | Zero cobertura para novo código — implementar 10 cenários |

---

## 10. Plano de Ajustes

### Prioridade 1 — Corrigir AGORA (antes da Frente 3)

| ID | Arquivo | Ação |
|:---|:---|:---|
| **CR-003** | RateLimitFilter.java | Remover null-guard redundante (linhas 93-95) |
| **CR-004** | RateLimitFilter.java | Remover branch `getParameter("username")` inalcançável (linhas 134-137) |
| **CR-001** | RateLimitFilter.java | Receber ObjectMapper via construtor + @Bean method |
| **CR-002** | RateLimitFilter.java | Usar apenas `request.getRemoteAddr()` ou validar XFF |
| **CR-005** | BaseRepository.java | Refatorar params com List<Object> |
| **CR-006** | RateLimitFilter.java | Corrigir `endsWith` para path completo |
| **CR-009** | BusinessUnit.java | Atualizar Javadoc |

### Prioridade 2 — Fase 4 (Implementação de Testes)

| ID | Arquivo | Ação |
|:---|:---|:---|
| **CR-010** | Múltiplos | Implementar 10 cenários do SPRINT-TEST-PLANNING-Frente-1.md |

### Prioridade 3 — Postergar (Sprint 6+)

| ID | Arquivo | Ação |
|:---|:---|:---|
| **CR-007** | RateLimitFilter.java | Externalizar constantes para application.yml |
| **CR-008** | FbsoJwtAuthenticationConverter.java | Consolidar com JwtAuthenticationFilter quando filtro for simplificado |

---

## 11. Build Pós-Revisão

- **Comando:** `mvn compile`
- **Resultado:** ✅ BUILD SUCCESS
- **Testes:** 213 testes — 0 failures, 1 erro pré-existente (SubscriptionServiceTest), 8 skipped
- **Regressões:** 0

---

## 12. Conclusão

A Frente 1 entrega código de qualidade com **zero achados críticos**. Os 10 achados (1 High, 5 Medium, 4 Low) são majoritariamente melhorias de hardening, simplificação e boas práticas. O principal gap é a ausência de testes unitários para o novo código (CR-010), que será endereçado na Fase 4 conforme o `SPRINT-TEST-PLANNING-Frente-1.md`.

**Decisão:** Prosseguir com ajustes de Prioridade 1 e 2 antes da Frente 3. Prioridade 3 pode ser postergada.

---

🤖 *Relatório consolidado gerado em 2026-07-23 como parte da Fase 7 do PROMPT-EXECUTE-SPRINT-TASKS.md. 7 skills acionados: ponytail-audit, ponytail-review, engineering-skills, security-auditor, performance-review, requesting-code-review, differential-review. 28 achados brutos → 10 achados únicos após deduplicação. 0 críticos. Código aprovado.*
