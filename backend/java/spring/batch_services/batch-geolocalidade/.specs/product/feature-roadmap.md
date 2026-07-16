---
title: "Feature Roadmap — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["product", "roadmap", "reference"]
---

# Feature Roadmap — batch-geolocalidade

## Features Concluídas

| Feature | Data | Descrição |
|---|---|---|
| Importação de Municípios | 2026-03-20 | Step 1: leitura de DTB_Municipios.csv, criação em cascata de UF/Regiões/Municípios |
| Importação de Distritos | 2026-03-20 | Step 2: leitura de DTB_Distritos.csv com FK para Municípios via proxy JPA |
| Importação de Subdistritos | 2026-03-20 | Step 3: leitura de DTB_Subdistritos.csv com FK para Distritos via proxy JPA |
| Dual Schema PostgreSQL | 2026-03-20 | Separação `spring_batch` / `localidade` com criação idempotente |
| LoadTestRunner determinístico | 2026-03-20 | Runner condicional com exit code 0/1 |
| Cache de entidades no processor | 2026-03-20 | HashMap para UF, Região Intermediária e Região Imediata |

## Features Planejadas

| Feature | Prioridade | Descrição |
|---|---|---|
| Spring Actuator Health Check | Média | Expor `/actuator/health` para monitoramento |
| REST endpoint para disparar Jobs | Média | `POST /api/v1/jobs/importacao-geolocalidade/launch` |
| REST endpoint para status de Jobs | Média | `GET /api/v1/jobs/importacao-geolocalidade/executions/{id}` |
| Migração Flyway | Alta | Substituir `ddl-auto: update` por migrations versionadas |
| Índices em colunas FK | Alta | Adicionar índices em `uf_id`, `regiao_intermediaria_id`, etc. |
| Suporte a atualização incremental | Baixa | Importar apenas municípios/distritos novos ou alterados |
| Métricas Prometheus | Média | Expor métricas de execução (duração, counts) via Micrometer |

## Dívidas Técnicas

| ID | Descrição | Impacto | Prioridade |
|---|---|---|---|
| DT-01 | Uso de `ddl-auto: update` em produção (recomendado: `validate` + Flyway) | Risco de perda de dados em migrations automáticas | Alta |
| DT-02 | Ausência de índices em colunas FK | Performance degradada em queries com JOIN | Alta |
| DT-03 | Mapa `UF_SIGLAS` hardcoded no `MunicipioProcessor` | Se o IBGE criar nova UF, requer alteração de código | Baixa |
| DT-04 | Cache em memória sem TTL | Se o Job for muito longo, pode consumir memória desnecessária | Baixa |
