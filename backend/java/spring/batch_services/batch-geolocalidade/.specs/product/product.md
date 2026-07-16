---
title: "Produto — batch-geolocalidade"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["product", "explanation"]
---

# Produto — batch-geolocalidade

## Visão do Produto

**batch-geolocalidade** é um serviço de processamento batch responsável por importar e manter atualizada a base de dados geopolíticos do Brasil (Divisão Territorial Brasileira — DTB) a partir de arquivos oficiais do IBGE.

## Proposta de Valor

- **Dados oficiais e atualizados**: Fonte IBGE, a referência nacional para dados territoriais
- **Processamento confiável**: Spring Batch garante transacionalidade, restartabilidade e rastreabilidade
- **Integração simplificada**: Tabelas padronizadas consumidas por outros microserviços (ex: `ms-geolocalidade`)
- **Operação automatizada**: Executável como CronJob Kubernetes, sem intervenção manual

## Personas

| Persona | Descrição | Interesse |
|---|---|---|
| **Operador DevOps** | Responsável por implantar e monitorar o batch em produção | Execução confiável, logs claros, exit codes determinísticos |
| **Desenvolvedor Backend** | Consome os dados de localidade em outros serviços | Dados normalizados, schema previsível, compatibilidade com códigos IBGE |
| **Analista de Dados** | Consulta as tabelas para análises geoespaciais | Dados completos e atualizados |

## Contexto de Negócio

Este serviço faz parte do ecossistema **FBSO** de gestão tributária. A localização geográfica é essencial para:
- Determinar alíquotas de tributos estaduais e municipais (ICMS, ISS, IBS)
- Identificar regimes especiais por região
- Validar inscrições estaduais e municipais (IE, IM)
