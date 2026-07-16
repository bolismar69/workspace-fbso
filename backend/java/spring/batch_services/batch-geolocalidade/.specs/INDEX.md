---
title: "Índice da Documentação — batch-geolocalidade"
version: "1.0"
last_updated: "2026-07-08"
confidence_score: "85%"
---

# Índice da Documentação — batch-geolocalidade

Serviço **Spring Batch** para importação de dados do IBGE (DTB) a partir de arquivos CSV e persistência em PostgreSQL via JPA/Hibernate.

Stack: **Java 21**, **Spring Boot 3.5.12**, **Spring Batch 5.x**, **PostgreSQL 16**.

---

## 📁 Domain

| Arquivo | Descrição |
|---|---|
| [domain/domain.md](domain/domain.md) | Glossário de domínio, ubiquitous language, hierarquia geopolítica brasileira |

## 📁 API

| Arquivo | Descrição |
|---|---|
| [api/batch-geolocalidade-api.yaml](api/batch-geolocalidade-api.yaml) | Contrato OpenAPI 3.1 (serviço headless — endpoints planejados) |

## 📁 Architecture

| Arquivo | Descrição |
|---|---|
| [architecture/architecture.md](architecture/architecture.md) | Visão arquitetural geral, stack, diretórios, princípios |
| [architecture/c4-context.md](architecture/c4-context.md) | C4 — Nível 1: Contexto do sistema |
| [architecture/c4-containers.md](architecture/c4-containers.md) | C4 — Nível 2: Containers |
| [architecture/c4-components.md](architecture/c4-components.md) | C4 — Nível 3: Componentes |
| [architecture/c4-code-class.md](architecture/c4-code-class.md) | C4 — Nível 4: Código/Classes |
| [architecture/erd.md](architecture/erd.md) | Diagrama Entidade-Relacionamento |
| [architecture/data-dictionary.md](architecture/data-dictionary.md) | Dicionário de dados (tabelas, colunas, tipos, regras) |
| [architecture/integrations.md](architecture/integrations.md) | Integrações externas, dependências Maven, env vars |
| [architecture/adrs/INDEX.md](architecture/adrs/INDEX.md) | Índice cronológico dos ADRs |
| [architecture/adrs/adr-001.md](architecture/adrs/adr-001.md) | ADR-0001: Chaves Naturais (Código IBGE) como PK |
| [architecture/adrs/adr-002.md](architecture/adrs/adr-002.md) | ADR-0002: Padrão Dual Schema |

## 📁 Engineering

| Arquivo | Descrição |
|---|---|
| [engineering/api-guidelines.md](engineering/api-guidelines.md) | Padrões de API, exit codes, endpoints planejados |
| [engineering/code-analysis.md](engineering/code-analysis.md) | Análise de fluxo de código: bootstrap, chunk processing, cache, erros |

## 📁 Product

| Arquivo | Descrição |
|---|---|
| [product/product.md](product/product.md) | Descrição do produto, visão, personas |
| [product/requirements.md](product/requirements.md) | Especificação de requisitos (RF/RNF + MoSCoW) |
| [product/feature-roadmap.md](product/feature-roadmap.md) | Roadmap de features + dívidas técnicas |

## 📁 Governance

| Arquivo | Descrição |
|---|---|
| [governance/inventory.md](governance/inventory.md) | Inventário do projeto, arquivos, cobertura |
| [governance/confidence-report.md](governance/confidence-report.md) | Relatório de confiança da documentação (85%) |

## 📁 Security

| Arquivo | Descrição |
|---|---|
| [security/SECURITY.md](security/SECURITY.md) | Definições de segurança, superfície de ataque, OWASP checklist |

---

## 📁 Outros

| Arquivo | Descrição |
|---|---|
| [CHANGELOG.md](CHANGELOG.md) | Histórico de mudanças da documentação |

---

## 🔢 Stats

- **Artefatos totais:** 19 arquivos de documentação
- **Diagramas Mermaid:** 9 (C4, ERD, fluxogramas, class diagrams)
- **ADRs:** 2
- **Cobertura C4:** 4/4 níveis ✅
- **Confiança:** 85% 🟢

---

*Última atualização: 2026-07-08 — Mineração inicial completa (SCOPE=full).*
