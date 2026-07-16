---
title: "Especificação de Requisitos — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["product", "requirements", "reference"]
---

# Especificação de Requisitos — batch-geolocalidade

## Requisitos Funcionais (RF)

| ID | Requisito | MoSCoW | Status |
|---|---|---|---|
| RF-01 | O sistema DEVE importar dados dos arquivos CSV da DTB/IBGE | Must | ✅ Implementado |
| RF-02 | O sistema DEVE processar a hierarquia completa: UF → Região Intermediária → Região Imediata → Município → Distrito → Subdistrito | Must | ✅ Implementado |
| RF-03 | O sistema DEVE respeitar a ordem de dependência: Municípios antes de Distritos, Distritos antes de Subdistritos | Must | ✅ Implementado |
| RF-04 | O sistema DEVE pular as 7 primeiras linhas de metadados dos CSVs do IBGE | Must | ✅ Implementado |
| RF-05 | O sistema DEVE suportar encoding UTF-8 nos arquivos de entrada | Must | ✅ Implementado |
| RF-06 | O sistema DEVE tolerar colunas extras nos CSVs (vírgula no final da linha) | Must | ✅ Implementado |
| RF-07 | O sistema DEVE derivar a sigla da UF a partir do código IBGE | Must | ✅ Implementado |
| RF-08 | O sistema DEVE encerrar com exit code 0 (sucesso) ou 1 (falha) | Must | ✅ Implementado |
| RF-09 | O sistema DEVE usar dois schemas PostgreSQL: `spring_batch` (metadata) e `localidade` (negócio) | Must | ✅ Implementado |
| RF-10 | O sistema DEVE validar a existência dos arquivos CSV antes de iniciar o Job | Should | ✅ Implementado |
| RF-11 | O sistema NÃO DEVE executar Jobs automaticamente no startup (`spring.batch.job.enabled=false`) | Must | ✅ Implementado |
| RF-12 | O sistema DEVE expor endpoint REST para disparar Jobs sob demanda | Could | ❌ Não implementado |
| RF-13 | O sistema DEVE expor health check via Spring Actuator | Could | ❌ Não implementado |

## Requisitos Não-Funcionais (RNF)

| ID | Requisito | MoSCoW | Status |
|---|---|---|---|
| RNF-01 | O sistema DEVE processar ~16.600 registros (5.570 municípios + 10.407 distritos + 684 subdistritos) em menos de 5 minutos | Must | ✅ Implementado |
| RNF-02 | O sistema DEVE usar chunks de 100 registros por transação | Must | ✅ Implementado |
| RNF-03 | O sistema DEVE usar `getReferenceById()` em vez de `findById()` para FKs, para evitar SELECTs desnecessários | Must | ✅ Implementado |
| RNF-04 | O sistema DEVE usar cache em memória para UF, Região Intermediária e Região Imediata durante o processamento | Should | ✅ Implementado |
| RNF-05 | O sistema DEVE ser executável como Job Kubernetes (pod efêmero) | Must | ✅ Implementado |
| RNF-06 | O sistema DEVE suportar configuração via variáveis de ambiente | Must | ✅ Implementado |
| RNF-07 | O sistema DEVE rodar em Java 21 | Must | ✅ Implementado |
