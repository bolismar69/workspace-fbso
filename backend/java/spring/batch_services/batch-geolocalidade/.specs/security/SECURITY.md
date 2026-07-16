---
title: "Segurança — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["security", "reference"]
---

# Segurança — batch-geolocalidade

## Visão Geral

**batch-geolocalidade** é um serviço batch headless executado em ambiente controlado (Kubernetes Cluster interno). Como não expõe endpoints REST, a superfície de ataque é reduzida.

## Superfície de Ataque

| Vetor | Exposição | Mitigação |
|---|---|---|
| **REST API** | ❌ Nenhuma (sem controllers) | — |
| **JDBC (PostgreSQL)** | Credenciais no `application.yaml` ou env vars | Usar Secrets Kubernetes; credenciais nunca em plain text no repositório |
| **File System (CSVs)** | Leitura de arquivos do path configurável | Montar volume como read-only; validar que path não permite path traversal |
| **Dependencies (Maven)** | Vulnerabilidades em bibliotecas | Executar `mvn dependency-check` periodicamente |
| **Spring Actuator** | Não ativo atualmente | Se ativado no futuro, proteger com Spring Security |

## Práticas Recomendadas

### 1. Credenciais de Banco

```yaml
# NUNCA hardcodar senhas no application.yaml
# Usar variáveis de ambiente ou Secrets Kubernetes:
spring:
  datasource:
    password: ${SPRING_DATASOURCE_PASSWORD}  # ✅
    # password: "minha_senha"                 # ❌
```

### 2. Path de Importação

- Montar volume de CSVs como **read-only** no Kubernetes
- Validar que `APP_IMPORT_PATH` não aceita paths relativos ou `..`
- O `FileSystemResource` do Spring resolve paths absolutos — garantir que o path base é seguro

### 3. Dependências

- Manter Spring Boot 3.5.x atualizado (patch de segurança)
- Verificar CVE no driver PostgreSQL
- H2 é usado apenas em testes (scope `test`) — não chega ao artefato de produção

### 4. Princípio do Menor Privilégio

| Componente | Permissão |
|---|---|
| Usuário PostgreSQL (`worker_user`) | Leitura/Escrita nos schemas `spring_batch` e `localidade` |
| Usuário do microserviço consumidor | Somente leitura no schema `localidade` |
| Pod Kubernetes | Leitura do volume de CSVs; sem acesso à internet |

## Checklist OWASP (Resumido)

| Controle | Status |
|---|---|
| A01: Broken Access Control | N/A (sem API) |
| A02: Cryptographic Failures | ✅ TLS para JDBC recomendado em produção |
| A03: Injection | ✅ JPA previne SQL injection; sem entrada de usuário |
| A04: Insecure Design | ✅ Princípio do menor privilégio aplicado |
| A05: Security Misconfiguration | ✅ `spring.batch.job.enabled=false` |
| A06: Vulnerable Components | ⚠️ Verificar com `mvn dependency-check` |
| A07: Auth Failures | N/A (sem autenticação) |
| A08: Software/Data Integrity | ✅ Maven checksums; CSVs validados antes de processar |
| A09: Logging/Monitoring Failures | ✅ SLF4J logging; métricas de batch |
| A10: SSRF | N/A (sem chamadas HTTP) |
