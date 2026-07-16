# Mapa do Time Técnico — Alocação e Competências

- **Projeto:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
- **Programa Pai:** PRJ-FIN-2026-0001 — Adequação Corporativa à Reforma Tributária Nacional
- **Referência:** [PLANO-TECNICO.md](./PLANO-TECNICO.md) — Seção 11 (Estrutura Proposta de Times)
- **Data de Criação:** 11 de Julho de 2026
- **Versão:** 1.0
- **Status:** Aguardando preenchimento de nomes e contatos

---

## 1. Objetivo

Este documento é a planilha de alocação do time técnico. Ele relaciona todos os papéis necessários ao projeto, as competências requeridas, a frente de atuação (DT-1, DT-2 ou DT-3), as fases em que cada profissional atua e a carga horária prevista.

As colunas **Nome** e **Contato** devem ser preenchidas pelo PM ou Tech Lead responsável tão logo os profissionais sejam designados.

---

## 2. Legenda

| Campo | Descrição |
|:---|:---|
| **Nome** | Nome completo do profissional designado (a preencher) |
| **Contato** | E-mail ou ramal corporativo (a preencher) |
| **Papel** | Função no projeto |
| **Skill Primária** | Principal competência técnica exigida para o papel |
| **Nível** | Proficiência esperada: ★★★ Avançado/Autônomo, ★★☆ Intermediário/Produtivo, ★☆☆ Básico/Assistido. Ver [PLANO-TECNICO.md](./PLANO-TECNICO.md) Seção 11.2 para a legenda completa. |
| **Atuação (DT)** | Frente técnica: DT-1 (Backend Java), DT-2 (Frontend React), DT-3 (Go Refactor), Transversal, Gestão ou Qualidade |
| **Fases** | Em quais fases do projeto o profissional atua: F0 = Fundação, F1 = Definições, F2 = Desenvolvimento, F3 = Integração/Cleanup, F4 = Go-Live. Ver tabela de fases abaixo. |
| **Horas/dia** | Carga horária diária prevista com base no percentual de alocação × jornada padrão de 8h |
| **Horas Totais** | Estimativa de horas totais no projeto (a recalcular quando as durações dos sprints forem definidas) |

### 2.1 Correspondência de Fases

| Código | Fase | Sprints | Duração Estimada |
|:---|:---|:---|:---|
| **F0** | Fundação | Sprint 0 | 1-2 semanas |
| **F1** | Design e Definições Transversais | Sprint 1 | 1-2 semanas |
| **F2** | Desenvolvimento Paralelo | Sprints 2..N | 8-12 semanas (estimado) |
| **F3** | Integração, Cleanup e Revogação de Grants | Sprints N+1..M | 4-8 semanas (estimado) |
| **F4** | Estabilização e Go-Live | Sprints M+1.. | 2-4 semanas (estimado) |

> **Nota:** As durações acima são estimativas iniciais. As horas totais serão recalculadas quando o calendário de sprints for definido pelo PM.

---

## 3. Time DT-1 — Backend Java/Spring

### Tech Lead Java/Spring (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Tech Lead Java/Spring | Spring Boot 4.0.1 / Java 21 | ★★★ | DT-1 | F0, F1, F2, F3, F4 | F0-F1: 8h (100%) / F2-F4: 8h → 4h | ⬜ |

**Competências requeridas:** Java 21 (★★★), Spring Boot 4.0.1 (★★★), Spring Data JDBC (★★★), PostgreSQL/Flyway (★★★), GraalVM Native Image (★★★), Keycloak/SAML 2.0 (★★★), OpenAPI/Swagger (★★★), Maven (★★★), Testes JUnit 5/Mockito/Testcontainers (★★★), Docker/Kubernetes (★★☆), Checkstyle (★★☆)

### Dev Java/Spring Sênior (2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Dev Java/Spring Sênior | Spring Boot 4.0.1 / Java 21 | ★★★ | DT-1 | F0, F1, F2, F3, F4 | F0-F1: 6,5h (80%) / F2-F4: 8h → 4h | ⬜ |
|  |  | Dev Java/Spring Sênior | Spring Boot 4.0.1 / Java 21 | ★★★ | DT-1 | F0, F1, F2, F3, F4 | F0-F1: 6,5h (80%) / F2-F4: 8h → 4h | ⬜ |

**Competências requeridas:** Java 21 (★★★), Spring Boot 4.0.1 (★★★), Spring Data JDBC (★★★), PostgreSQL (★★☆), GraalVM Native Image (★★☆), Docker/Kubernetes (★★☆), Keycloak/SAML 2.0 (★★☆), OpenAPI (★★☆), Maven (★★★), Testes JUnit 5/Mockito/Testcontainers (★★★), Checkstyle (★★☆)

### Dev Java/Spring Pleno (1-2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Dev Java/Spring Pleno | Java 21 / Spring Boot 4.0.1 | ★★☆ | DT-1 | F0, F1, F2, F3, F4 | F0-F1: 4h (50%) / F2-F4: 8h → 4h | ⬜ |
|  |  | Dev Java/Spring Pleno | Java 21 / Spring Boot 4.0.1 | ★★☆ | DT-1 | F0, F1, F2, F3, F4 | F0-F1: 4h (50%) / F2-F4: 8h → 4h | ⬜ |

**Competências requeridas:** Java 21 (★★☆), Spring Boot 4.0.1 (★★☆), Spring Data JDBC (★★☆), PostgreSQL (★★☆), GraalVM Native Image (★☆☆), Docker/Kubernetes (★☆☆), Keycloak (★☆☆), OpenAPI (★★☆), Maven (★★☆), Testes JUnit 5/Mockito (★★☆), Checkstyle (★★☆)

---

## 4. Time DT-2 — Frontend React + UX/UI

### Tech Lead Frontend/React (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Tech Lead Frontend/React | React 19 / TypeScript | ★★★ | DT-2 | F0, F1, F2, F3, F4 | F0-F1: 8h (100%) / F2-F4: 8h → 4h | ⬜ |

**Competências requeridas:** React 19 (★★★), TypeScript (★★★), Vite (★★★), Design System/Component Library (★★☆), MSW (★★★), OpenAPI/Geração de Tipos (★★★), Consumo de APIs REST (★★★), Keycloak/SAML 2.0 frontend (★★☆), Testes Vitest/RTL/Playwright (★★★), Acessibilidade WCAG 2.1 AA (★★☆), CSS/Design Tokens (★★☆), DevOps Frontend (★★☆), Figma (★☆☆)

### Dev React Sênior (1-2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Dev React Sênior | React 19 / TypeScript | ★★★ | DT-2 | F0, F1, F2, F3, F4 | F0: 6,5h (80%) / F1: 4h (50%) / F2-F4: 8h → 4h | ⬜ |
|  |  | Dev React Sênior | React 19 / TypeScript | ★★★ | DT-2 | F0, F1, F2, F3, F4 | F0: 6,5h (80%) / F1: 4h (50%) / F2-F4: 8h → 4h | ⬜ |

**Competências requeridas:** React 19 (★★★), TypeScript (★★★), Vite (★★☆), Design System (★★☆), MSW (★★☆), OpenAPI/Geração de Tipos (★★☆), Consumo de APIs REST (★★★), Keycloak/SAML 2.0 frontend (★★☆), Testes Vitest/RTL/Playwright (★★★), Acessibilidade WCAG 2.1 AA (★★☆), CSS/Design Tokens (★★☆), DevOps Frontend (★★☆)

### Dev React Pleno (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Dev React Pleno | React 19 / TypeScript | ★★☆ | DT-2 | F0, F1, F2, F3, F4 | F0-F1: 4h (50%) / F2-F4: 8h → 4h | ⬜ |

**Competências requeridas:** React 19 (★★☆), TypeScript (★★☆), Vite (★★☆), Design System (★☆☆), MSW (★★☆), OpenAPI/Geração de Tipos (★★☆), Consumo de APIs REST (★★☆), Keycloak/SAML 2.0 frontend (★☆☆), Testes Vitest/RTL (★★☆), Acessibilidade WCAG 2.1 AA (★☆☆), CSS/Design Tokens (★★☆), DevOps Frontend (★☆☆)

### UX/UI Designer (1-2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | UX/UI Designer | Design System / Figma | ★★★ | DT-2 | F0, F1, F2, F3 | F0-F1: 8h (100%) / F2: 2,5h (30%) / F3: 1,5h (20%) | ⬜ |
|  |  | UX/UI Designer | UX Research / Testes de Usabilidade | ★★★ | DT-2 | F0, F1, F2, F3 | F0-F1: 8h (100%) / F2: 2,5h (30%) / F3: 1,5h (20%) | ⬜ |

**Competências requeridas:** Design System/Component Library (★★★), Acessibilidade WCAG 2.1 AA (★★★), Figma/Prototipação (★★★), UX Research/Testes de Usabilidade (★★★), CSS/Design Tokens/Temas (★★★)

---

## 5. Time DT-3 — Refatoração Go/Fiber

### Tech Lead Go/Fiber (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Tech Lead Go/Fiber | Go 1.21+ / Fiber Framework | ★★★ | DT-3 | F2, F3, F4 | F2: 1,5h (20%) / F3: 8h (100%) / F4: 6,5h (80%) | ⬜ |

**Competências requeridas:** Go 1.21+ (★★★), Fiber Framework (★★★), PostgreSQL (★★★), Refatoração/Análise de Código Legado (★★★), Testes de Regressão Go/testing (★★★), Docker/Kubernetes (★★☆), OpenAPI/Contratos de API (★★☆), Grants e Permissionamento DB (★★☆)

### Dev Go/Fiber Sênior (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Dev Go/Fiber Sênior | Go 1.21+ / Fiber Framework | ★★★ | DT-3 | F3, F4 | F3: 8h (100%) / F4: 6,5h (80%) | ⬜ |

**Competências requeridas:** Go 1.21+ (★★★), Fiber Framework (★★★), PostgreSQL (★★☆), Refatoração/Análise de Código Legado (★★☆), Testes de Regressão Go/testing (★★☆), Docker/Kubernetes (★★☆), OpenAPI/Contratos de API (★★☆), Grants e Permissionamento DB (★★☆)

---

## 6. Times Transversais

### Arquiteto de Solução (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Arquiteto de Solução | Arquitetura de Microsserviços / SAML 2.0 | ★★★ | Transversal | F0, F1, F2, F3, F4 | F0-F1: 8h (100%) / F2-F3: 2,5h (30%) / F4: 1,5h (20%) | ⬜ |

**Competências requeridas:** Arquitetura de Microsserviços (★★★), Desenho de APIs REST/Contratos (★★★), SAML 2.0/Keycloak (★★★), Modelagem de Dados/PostgreSQL (★★☆), Kubernetes (★★☆), CI/CD Pipelines (★★☆), Gestão de Segredos (★★☆), Segurança de Aplicações (★★☆), Docker/Containers (★★☆), GraalVM Native Image (★★☆), IaC (★★☆), Monitoramento (★★☆)

### DBA (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | DBA | Modelagem de Dados / PostgreSQL | ★★★ | Transversal | F1, F2, F3, F4 | F1: 6,5h (80%) / F2: 1,5h (20%) / F3: 4h (50%) / F4: 1,5h (20%) | ⬜ |

**Competências requeridas:** Modelagem de Dados/PostgreSQL (★★★), Grants, Roles e Segurança de DB (★★★), Monitoramento e Observabilidade (★☆☆)

### DevOps Engineer (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | DevOps Engineer | Kubernetes / CI/CD Pipelines | ★★★ | Transversal | F0, F1, F2, F3, F4 | F0: 2,5h (30%) / F1: 8h (100%) / F2: 2,5h (30%) / F3: 6,5h (80%) / F4: 4h (50%) | ⬜ |

**Competências requeridas:** Kubernetes (★★★), CI/CD Pipelines (★★★), Gestão de Segredos Vault/Secrets Manager (★★★), Docker/Containers (★★★), IaC (★★★), Monitoramento e Observabilidade (★★★), Arquitetura de Microsserviços (★★☆), Desenho de APIs REST (★★☆), Grants/Roles/DB (★★☆), SAML 2.0/Keycloak (★★☆), GraalVM Native Image (★★☆)

### DevSecOps Engineer (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | DevSecOps Engineer | Segurança de Aplicações / SAML 2.0 | ★★★ | Transversal | F0, F1, F2, F3, F4 | F0: 1,5h (20%) / F1: 8h (100%) / F2: 1,5h (20%) / F3: 6,5h (80%) / F4: 4h (50%) | ⬜ |

**Competências requeridas:** Segurança de Aplicações SAST/SCA (★★★), SAML 2.0/Keycloak (★★★), Grants, Roles e Segurança de DB (★★★), Gestão de Segredos (★★★), Kubernetes (★★☆), CI/CD Pipelines (★★☆), Docker/Containers (★★☆), IaC (★★☆), Monitoramento e Observabilidade (★★☆), Arquitetura de Microsserviços (★★☆)

### Infraestrutura / K8S Engineer (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Infraestrutura / K8S Engineer | Kubernetes / Docker | ★★★ | Transversal | F0, F1, F2, F3, F4 | F0: 1,5h (20%) / F1: 8h (100%) / F2: 1,5h (20%) / F3: 4h (50%) / F4: 4h (50%) | ⬜ |

**Competências requeridas:** Kubernetes (★★★), Docker/Containers (★★★), IaC (★★★), Monitoramento e Observabilidade (★★★), Arquitetura de Microsserviços (★★☆), CI/CD Pipelines (★★☆), GraalVM Native Image (★★☆)

---

## 7. Gestão

### Product Manager — PM (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Product Manager (PM) | Gestão de Projetos Ágeis | ★★★ | Gestão | F0, F1, F2, F3, F4 | F0-F1: 8h (100%) / F2-F4: 4h (50%) | ⬜ |

**Competências requeridas:** Gestão de Projetos Ágeis Scrum/Kanban (★★★), Gestão de Stakeholders (★★★), Comunicação e Facilitação (★★★), Domínio de Negócio Tributário (★★☆), Análise de Requisitos (★★☆)

### Product Owner — PO (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Product Owner (PO) | Domínio de Negócio Tributário | ★★★ | Gestão | F0, F1, F2, F3, F4 | F0: 4h (50%) / F1: 8h (100%) / F2-F4: 4h (50%) | ⬜ |

**Competências requeridas:** Domínio de Negócio Tributário (★★★), Gestão de Stakeholders (★★★), Análise de Requisitos/Critérios de Aceite (★★★), Comunicação e Facilitação (★★★), Gestão de Projetos Ágeis (★★☆), Homologação/UAT (★★☆)

---

## 8. Qualidade

### QA / Test Engineer (1-2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | QA / Test Engineer | Testes Automatizados E2E | ★★★ | Qualidade | F1, F2, F3, F4 | F1: 2,5h (30%) / F2: 4h (50%) / F3: 8h (100%) / F4: 6,5h (80%) | ⬜ |
|  |  | QA / Test Engineer | Testes Manuais / Exploratórios | ★★★ | Qualidade | F1, F2, F3, F4 | F1: 2,5h (30%) / F2: 4h (50%) / F3: 8h (100%) / F4: 6,5h (80%) | ⬜ |

**Competências requeridas:** Testes Automatizados E2E (★★★), Testes Manuais/Exploratórios (★★★), Testes de Regressão (★★★), Homologação/UAT (★★☆), Análise de Requisitos/Critérios de Aceite (★★☆), Gestão de Projetos Ágeis (★★☆), Domínio de Negócio Tributário (★☆☆)

### Time de Homologação — Negócio (1-2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Atuação (DT) | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---|:---:|:---:|
|  |  | Analista de Homologação (Negócio) | Domínio de Negócio Tributário | ★★★ | Qualidade | F1, F2, F3, F4 | F1: 1,5h (20%) / F2: 1,5h (20%) / F3: 2,5h (30%) / F4: 8h (100%) | ⬜ |
|  |  | Analista de Homologação (Negócio) | Domínio de Negócio Tributário | ★★★ | Qualidade | F1, F2, F3, F4 | F1: 1,5h (20%) / F2: 1,5h (20%) / F3: 2,5h (30%) / F4: 8h (100%) | ⬜ |

**Competências requeridas:** Domínio de Negócio Tributário (★★★), Testes Manuais/Exploratórios (★★★), Homologação/UAT (★★★), Testes de Regressão (★★★), Análise de Requisitos/Critérios de Aceite (★★★), Testes Automatizados E2E (★☆☆)

> **Nota:** O time de homologação é idealmente composto por usuários-chave do time de Finanças (analistas fiscais seniores, controller), com dedicação mais intensa na Fase 4 (Go-Live).

---

## 9. Tabela Resumo

### 9.1 Headcount por Frente

| Frente | Papéis | Vagas | Fases com Dedicação Máxima |
|:---|:---|:---:|:---|
| **DT-1 — Backend Java** | Tech Lead, Dev Sênior (2), Dev Pleno (1-2) | 4-5 | F2, F3 (100% desenvolvimento e integração) |
| **DT-2 — Frontend React** | Tech Lead, Dev Sênior (1-2), Dev Pleno (1), UX/UI Designer (1-2) | 4-6 | F2 (100% desenvolvimento) |
| **DT-3 — Go Refactor** | Tech Lead, Dev Sênior (1) | 2 | F3 (100% cleanup e revogação) |
| **Transversais** | Arquiteto, DBA, DevOps, DevSecOps, Infra/K8S | 5 | F1 (100% definições), F3 (50-80% integração e revogação) |
| **Gestão** | PM, PO | 2 | F0-F1 (100%) |
| **Qualidade** | QA/Test (1-2), Homologação (1-2) | 2-4 | F3 (100% QA), F4 (100% Homologação) |
| **Total** | | **19-24** | |

### 9.2 Alocação Média por Fase

| Fase | DT-1 (4-5p) | DT-2 (4-6p) | DT-3 (2p) | Transversais (5p) | Gestão (2p) | Qualidade (2-4p) | Total Ativo |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| **F0 — Fundação** | 2,9 pessoas | 2,9 pessoas | — | 1,2 pessoas | 1,5 pessoas | — | **8,5** |
| **F1 — Definições** | 2,9 pessoas | 2,0 pessoas | — | 4,6 pessoas | 2,0 pessoas | 0,8 pessoas | **12,3** |
| **F2 — Desenvolvimento** | 4,5 pessoas | 3,7 pessoas | 0,2 pessoas | 1,2 pessoas | 1,0 pessoas | 1,4 pessoas | **12,0** |
| **F3 — Integração/Cleanup** | 4,5 pessoas | 3,6 pessoas | 2,0 pessoas | 3,2 pessoas | 1,0 pessoas | 2,6 pessoas | **16,9** |
| **F4 — Go-Live** | 2,3 pessoas | 1,8 pessoas | 1,6 pessoas | 2,2 pessoas | 1,0 pessoas | 3,4 pessoas | **12,3** |

> **Cálculo:** Pessoas ativas = soma de (quantidade de vagas × percentual de alocação na fase). Ex: DT-1 na F0 = (1 TL × 100%) + (2 Sênior × 80%) + (1,5 Pleno × 50%) = 1 + 1,6 + 0,75 ≈ 2,9 pessoas equivalentes em tempo integral.

### 9.3 Marcos de Alocação

| Marco | Fase | O que acontece |
|:---|:---|:---|
| **Pré-Sprint 0** | — | PM e Arquiteto agendam sessões de definição com times transversais para o Sprint 1 |
| **Sprint 0 (F0)** | Fundação | Todos os times mobilizados. DT-3 ainda não alocado. UX/UI em carga máxima. |
| **Sprint 1 (F1)** | Definições | **Pico dos times transversais** (4,6 pessoas). UX/UI conclui e obtém aprovação. DT-1 e DT-2 finalizam definições de arquitetura, infra e segurança. |
| **Sprints 2..N (F2)** | Desenvolvimento | **Pico de DT-1 e DT-2**. Times transversais em baixa (suporte). Tech Lead DT-3 inicia alinhamentos. QA começa preparação de testes. |
| **Sprints N+1..M (F3)** | Integração/Cleanup | **Fase mais intensa do projeto** (16,9 pessoas). DT-3 em carga máxima. QA em 100%. DevOps e DevSecOps executam revogação de grants. |
| **Sprints M+1.. (F4)** | Go-Live | Homologação em 100%. Times de desenvolvimento em redução (50%). Monitoramento e rollback plan. |

---

## 10. Instruções de Preenchimento

1. **Nome e Contato:** Preencher tão logo o profissional seja designado. O PM é responsável por manter esta planilha atualizada.
2. **Horas Totais:** A coluna será preenchida quando as durações dos sprints forem definidas. Fórmula: `Horas/dia × dias úteis na fase × semanas da fase`.
3. **Nível:** A proficiência indicada é a **esperada** para o papel. Se o profissional designado tiver nível diferente, revisar o plano de capacitação ou mentoria.
4. **Substituições:** Em caso de substituição de profissional, atualizar a linha correspondente e registrar a alteração no histórico abaixo.

### 10.1 Histórico de Alterações

| Data | Alteração | Responsável |
|:---|:---|:---|
| 11/07/2026 | Criação do documento (v1.0) | — |

---

🤖 *Documento gerado com apoio de Claude Code (Anthropic), em 11 de Julho de 2026, com base na Seção 11 do PLANO-TECNICO.md v1.2.*
