# PROJECT-TECHNICAL-DEFINITIONS-TEAM-SKILLS-MAP — Matriz de Skills do Discovery Team

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 3.1 — Discovery Team (TOGAF Upstream Architecture)
- **Data de Criação:** 25 de Julho de 2026
- **Última Atualização:** 30 de Julho de 2026 (alinhamento com Bloco 0 v5.0)
- **Status:** Em Revisão / Aguardando Validação (Fase 5 — Bloco A)
- **Fase:** F5 — Bloco A (People & Solutions)
- **Bloco 0 — Artefatos Upstream:** [INTAKE-LOG](./PROJECT-TECHNICAL-DEFINITIONS-INTAKE-LOG.md) (F1) · [DOR-ASSESSMENT](./PROJECT-TECHNICAL-DEFINITIONS-DOR-ASSESSMENT.md) (F2) · [PRODUCT-BACKLOG-LIST](./PROJECT-TECHNICAL-DEFINITIONS-PRODUCT-BACKLOG-LIST.md) (F3) · [PRD-DEFINITION](./PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION.md) (F4 — escopo do produto)
- **Documento Complementar:** [PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md) (nomes, contatos, horas/semana)

---

## 1. Objetivo

Este documento apresenta a **matriz de competências técnicas** do time executor do projeto. Enquanto o [TEAM-CAPACITY](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md) responde **QUEM** está disponível e **quantas horas/dia**, este documento foca em **O QUE** cada perfil sabe fazer e seu **nível de proficiência**.

Ele serve como:
- **Referência de alocação:** Tech Lead consulta a matriz para designar a pessoa certa para cada tarefa
- **Detector de gaps:** Identifica competências necessárias que o time não possui
- **Plano de capacitação:** Base para definir treinamentos e contratações

---

## 2. Discovery Team — Papéis da Upstream Architecture (TOGAF Preliminary Phase)

Os papéis abaixo compõem o time fixo de **Discovery Técnico** responsável pela fase de Upstream Architecture (Blocos A e B do roadmap). Cada papel representa uma competência essencial para a análise e definição da arquitetura do projeto.

| # | Papel | Responsabilidade no Discovery | Profissional Designado | Nível |
|---|-------|-------------------------------|------------------------|-------|
| 1 | Engenheiro de Sistemas | Visão sistêmica, integração entre componentes, requisitos não-funcionais | Francisco Oliveira (Tech Lead) | ★★★ |
| 2 | Arquiteto de Soluções | Definição da arquitetura macro, C4, ADRs, padrões cross-solution | Bruno Gratto (Solution Architect) | ★★★ |
| 3 | Arquiteto de Banco de Dados | Modelagem de dados, multi-tenant, schema design, migrações | Carlos Caldas (DB Developer) | ★★★ |
| 4 | Arquiteto de DevOps/SRE | Infraestrutura, CI/CD, containerização, observabilidade, deploy | Davi Silva (DevOps) | ★★★ |
| 5 | Arquiteto de Segurança | Threat model, IAM, secrets management, compliance (LGPD, PCI, SOC2) | Gertrudes Paiva (IAM Specialist) | ★★★ |
| 6 | Arquiteto/Especialista de Testes | Estratégia de testes, tipos de teste, ferramentas, qualidade | Felipe Canedas (QA Engineer) | ★★★ |
| 7 | Arquiteto/Especialista de Cloud-Provider | Serviços cloud, networking, custos, multi-cloud | Davi Silva (DevOps) | ★★★ |
| 8 | Líder Técnico / Tech Lead | Coordenação técnica, code review, mentoria, decisões de design | Francisco Oliveira (Tech Lead) | ★★★ |
| 9 | Especialista em Integrações/APIs | Contratos de API, mensageria, integração entre soluções | Bolismar Oliveira (Full-Stack) | ★★★ |

> ⚠️ **Nota:** Alguns profissionais acumulam mais de um papel no Discovery Team (ex: Francisco cobre Engenharia de Sistemas + Tech Lead; Davi cobre DevOps/SRE + Cloud-Provider). O [TEAM-CAPACITY](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md) detalha a disponibilidade horária de cada um.

---

## 3. Contexto Técnico do Projeto

### 3.1 Stack Tecnológica Definida

| Camada | Tecnologia | Relevância |
|:---|:---|:---|
| **Backend** | Java 25 LTS + Spring Boot 3.5+ (Monolítico Modular REST) | Core da aplicação |
| **Build** | Maven, Oracle GraalVM 25.0.3+9.1 (Native Image AOT) | Compilação e containerização |
| **Banco de Dados** | PostgreSQL 17 (Multi-Tenant lógico via `tenant_id`, Soft Delete) | Persistência |
| **IAM** | Keycloak 26 (SAML 2.0 + OAuth 2.0/OIDC + JWT) | Autenticação e autorização |
| **Frontend** | React + Next.js + Tailwind CSS | Interface web (admin + portal cliente) |
| **Containerização** | Docker, Docker Compose, Kubernetes | Deploy e orquestração |
| **Mensageria (futuro)** | RabbitMQ | Comunicação assíncrona entre módulos |
| **Observabilidade** | OpenTelemetry, Micrometer, Grafana, CloudWatch (via AWS) | Logs, métricas, tracing, dashboards |
| **Cloud/Hosting** | AWS, Hostinger, Digital-Ocean | Infraestrutura e hospedagem |
| **Segurança** | RBAC, JWT validation, CORS, Rate Limiting, OWASP Top 10 | Postura de segurança |

### 3.2 Escopo Técnico (18 funcionalidades, 58 user stories)

| Épico | Funcionalidades | Complexidade Técnica |
|:---|:---|:---|
| EP-0001 — Portal Admin Interno | Dashboard, Filtros, Alertas | Média (queries analíticas, gráficos) |
| EP-0002 — Clientes e Assinaturas | Tenants, Planos, Assinaturas, Auditoria | Alta (multi-tenant, integridade referencial) |
| EP-0003 — RBAC | Usuários, Papéis, Permissões, Visibilidade | Alta (modelo de segurança, regras granulares) |
| EP-0004 — Portal do Cliente | Autenticação, Onboarding, Catálogo, App Switcher | Muito Alta (UX, fluxos complexos, Keycloak SSO) |

---

## 4. Matriz de Competências

### 4.1 Matriz Perfil × Tecnologia × Proficiência

**Legenda de Proficiência:**
- ★★★ Senior/Especialista — Autônomo, define padrões, mentoriza
- ★★☆ Pleno/Intermediário — Produtivo, segue padrões, resolve problemas típicos
- ★☆☆ Junior/Básico — Assistido, executa tarefas bem definidas
- `—` Sem exposição conhecida

| Tecnologia / Domínio | TL<br>Francisco | FS<br>Bolismar | FE<br>Tom | BE<br>Maria | QA<br>Felipe | BA<br>Mauro | DB<br>Carlos | SA<br>Bruno | DevOps<br>Davi | IAM<br>Gertrudes | BE<br>Judith |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| **Linguagens** | | | | | | | | | | | |
| Java 25 LTS | ★★★ | ★★★ | — | ★☆☆ | ★☆☆ | — | — | ★☆☆ | — | — | ★★★ |
| JavaScript/TypeScript | ★★☆ | ★★★ | ★★★ | ★☆☆ | ★☆☆ | — | — | ★☆☆ | — | — | — |
| SQL (ANSI/PostgreSQL) | ★★★ | ★★★ | ★☆☆ | ★☆☆ | ★☆☆ | ★☆☆ | ★★★ | ★☆☆ | — | — | ★☆☆ |
| HTML/CSS | ★★☆ | ★★★ | ★★★ | — | ★☆☆ | — | — | ★☆☆ | — | — | — |
| Shell Script (Bash) | ★★★ | ★★★ | — | ★☆☆ | ★☆☆ | — | ★☆☆ | ★☆☆ | ★★★ | — | ★☆☆ |
| **Frameworks Backend** | | | | | | | | | | | |
| Spring Boot 3.5+ | ★★★ | ★★★ | — | ★☆☆ | ★☆☆ | — | — | ★☆☆ | — | — | ★★★ |
| Spring Security | ★★★ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | — | — | ★★★ |
| Spring Data JDBC / JPA | ★★★ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | — | — | ★★★ |
| Spring Validation | ★★★ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | — | — | ★★★ |
| JUnit 5 + Mockito | ★★★ | ★★★ | — | ★☆☆ | ★★★ | — | — | ★☆☆ | — | — | ★★★ |
| **Frameworks Frontend** | | | | | | | | | | | |
| React 19+ | ★★☆ | ★★★ | ★★★ | — | ★☆☆ | — | — | ★☆☆ | — | — | — |
| Next.js (App Router) | ★★☆ | ★★★ | ★★★ | — | — | — | — | ★☆☆ | — | — | — |
| Tailwind CSS | ★☆☆ | ★★★ | ★★★ | — | — | — | — | ★☆☆ | — | — | — |
| **Banco de Dados** | | | | | | | | | | | |
| PostgreSQL 17 | ★★★ | ★★★ | — | ★☆☆ | ★☆☆ | ★☆☆ | ★★★ | ★☆☆ | — | — | ★☆☆ |
| Flyway Migrations | ★★★ | ★★★ | — | ★☆☆ | — | — | ★★★ | ★☆☆ | — | — | ★☆☆ |
| Multi-Tenant Design | ★★☆ | ★★★ | — | ★☆☆ | — | — | ★☆☆ | ★☆☆ | — | — | ★☆☆ |
| **IAM & Segurança** | | | | | | | | | | | |
| Keycloak 26 | ★★☆ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | — | ★★★ | — |
| SAML 2.0 | ★☆☆ | ★★☆ | — | ★☆☆ | — | — | — | ★☆☆ | — | ★★★ | — |
| OAuth 2.0 / OIDC | ★★☆ | ★★★ | ★☆☆ | ★☆☆ | — | — | — | ★☆☆ | — | ★★★ | — |
| JWT / Token Validation | ★★★ | ★★★ | ★★☆ | ★☆☆ | ★☆☆ | — | — | ★☆☆ | — | ★★★ | ★☆☆ |
| RBAC Design | ★★★ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | — | ★★★ | — |
| OWASP Top 10 | ★★★ | ★★★ | ★★☆ | ★☆☆ | ★★★ | — | — | ★☆☆ | — | ★★★ | ★☆☆ |
| **DevOps & Infra** | | | | | | | | | | | |
| Docker / Docker Compose | ★★★ | ★★★ | ★★☆ | ★☆☆ | ★☆☆ | — | — | ★☆☆ | ★★★ | — | ★☆☆ |
| Kubernetes | ★☆☆ | ★★★ | — | — | — | — | — | ★☆☆ | ★★★ | — | — |
| Kong API Gateway | ★★★ | ★★☆ | — | — | — | — | — | — | ★★★ | — | — |
| GraalVM Native Image | ★★☆ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | — | — | — |
| CI/CD (GitHub Actions) | ★★★ | ★★★ | ★★☆ | ★☆☆ | ★☆☆ | — | — | ★☆☆ | ★★★ | — | ★☆☆ |
| **Cloud & Hosting** | | | | | | | | | | | |
| AWS | ★★☆ | ★★★ | — | — | — | — | — | ★☆☆ | ★★★ | — | — |
| Hostinger | — | — | — | — | — | — | — | — | ★★★ | — | — |
| Digital-Ocean | — | — | — | — | — | — | — | — | ★★★ | — | — |
| **Qualidade & Testes** | | | | | | | | | | | |
| Testes Unitários | ★★★ | ★★★ | ★★☆ | ★☆☆ | ★★★ | — | — | ★☆☆ | — | — | ★★★ |
| Testes de Integração | ★★★ | ★★★ | — | ★☆☆ | ★★★ | — | — | ★☆☆ | — | — | ★★☆ |
| Testes E2E (Playwright) | ★★☆ | ★★★ | ★★★ | — | ★★★ | ★★★ | — | ★☆☆ | — | — | — |
| Testes de Segurança (SAST) | ★★☆ | ★★★ | — | ★☆☆ | ★★★ | — | — | ★☆☆ | — | — | — |
| Testes de Performance (JMeter) | ★★☆ | ★★★ | — | — | ★★★ | — | — | ★☆☆ | — | — | — |
| **Observabilidade** | | | | | | | | | | | |
| OpenTelemetry | ★★☆ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | ★★★ | — | — |
| Micrometer | ★★☆ | ★★★ | — | ★☆☆ | — | — | — | ★☆☆ | ★★★ | — | ★☆☆ |
| Grafana | — | ★★☆ | — | — | — | — | — | — | ★★★ | — | — |
| Logging (SLF4J/Logback) | ★★★ | ★★★ | — | ★☆☆ | ★★☆ | — | ★★★ | ★☆☆ | ★★★ | — | ★★★ |

### 4.2 Resumo por Papel

| Papel | Profissional | Tecnologias Core (★★★) | Carga |
|:---|:---|:---|:---:|
| Tech Lead | Francisco Oliveira | Java, Spring Boot, SQL, JUnit, Docker, CI/CD, RBAC, OWASP, Logging, Kong API Gateway | 8h/d |
| Full-Stack | Bolismar Oliveira | Java, Spring Boot, React, Tailwind, SQL, JUnit, Docker, Kubernetes, AWS, CI/CD, OpenTelemetry, Micrometer, Bash, RBAC, OWASP, Playwright, Logging, Multi-Tenant Design, Keycloak, OAuth/OIDC, Spring Security, SAST, JMeter, GraalVM, JWT, Flyway | 8h/d |
| Frontend | Tom Santos | React, Next.js, Tailwind, TypeScript, Playwright (início 01/11/2026) | 8h/d |
| Backend (Jr.) | Maria Madalena | Nenhuma tecnologia em nível ★★★ — perfil junior em capacitação | 8h/d |
| QA Engineer | Felipe Canedas | JUnit, Playwright, SAST, JMeter, OWASP — domina seu escopo. Demais tecnologias requer apoio de outros profissionais | 8h/d |
| BA (Homologação) | Mauro Hanashiro | Playwright (testes de sistema frontend — ótima experiência). SQL básico (★☆☆). Postman básico. | 8h/d |
| DB Developer | Carlos Caldas | PostgreSQL, SQL, Flyway, Logging — conhecimento limitado a banco de dados | 8h/d |
| Solution Architect | Bruno Gratto | Conhecimento conceitual e arquitetural sólido. Proficiência prática básica (★☆☆) nas tecnologias. | 8h/d |
| DevOps | Davi Silva | Docker, Kubernetes, AWS, Hostinger, Digital-Ocean, CI/CD, Kong API Gateway, OpenTelemetry, Micrometer, Grafana, Bash, Logging — hiperespecialista em infraestrutura e cloud. Demais tecnologias: desconhece. | 8h/d |
| IAM Specialist | Gertrudes Paiva | Keycloak, SAML 2.0, OAuth 2.0/OIDC, JWT, RBAC, OWASP — excelência em seu foco. Demais tecnologias: desconhece. | 8h/d |
| Backend (temp.) | Judith Campos | Java, Spring Boot, JUnit, Spring Security, Logging (27-31/07/2026). SQL básico (★☆☆). | temporário |

> ⚠️ **Nota sobre carga parcial:** Arquiteto, DevOps e Especialista IAM atuam com dedicação parcial (4h/dia efetivas para o projeto), com picos em fases de setup (F0) e menor intensidade em fases de desenvolvimento (F1-F2), conforme [TEAM-CAPACITY](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md).

---

## 5. Skills por Categoria

### 5.1 Linguagens de Programação

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| Java 25 LTS | ✅ Forte | Francisco, Bolismar, Judith | Maria é junior (★☆☆) — precisa de mentoria |
| JavaScript/TypeScript | ✅ Forte | Bolismar, Tom | Backend team tem exposição limitada. Tom só inicia em 01/11 |
| SQL (ANSI/PostgreSQL) | ✅ Forte | Francisco, Bolismar, Carlos | — |
| HTML/CSS | ⚠️ Depende de Tom | Bolismar, Tom | Tom só inicia em 01/11/2026 |
| Shell Script (Bash) | ✅ Forte | Francisco, Bolismar, Davi | — |

### 5.2 Frameworks Backend

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| Spring Boot 3.5+ | ✅ Forte | Francisco, Bolismar, Judith | Maria (★☆☆) — pareamento necessário |
| Spring Security | ✅ Forte | Francisco, Bolismar, Judith | — |
| Spring Data JDBC/JPA | ✅ Forte | Francisco, Bolismar, Judith | — |
| Spring Validation | ✅ Forte | Francisco, Bolismar, Judith | — |
| JUnit 5 + Mockito | ✅ Forte | Francisco, Bolismar, Felipe, Judith | — |

### 5.3 Frameworks Frontend

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| React 19+ | ⚠️ Depende de Tom | Bolismar, Tom | Tom só inicia em 01/11. Bolismar cobre full-stack até lá |
| Next.js (App Router) | ⚠️ Limitada | Bolismar, Tom | Sem redundancy além de Bolismar até novembro |
| Tailwind CSS | ⚠️ Depende de Tom | Bolismar, Tom | Idem — Bolismar cobre até chegada de Tom |

### 5.4 Banco de Dados

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| PostgreSQL 17 | ✅ Forte | Francisco, Bolismar, Carlos | — |
| Flyway Migrations | ✅ Coberto | Francisco, Bolismar, Carlos | — |
| Multi-Tenant Design | ✅ Coberto | Bolismar | Francisco e demais em ★★☆/★☆☆. Bruno cobre conceitualmente. |

### 5.5 IAM & Segurança

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| Keycloak 26 | ✅ Coberto | Bolismar, Gertrudes | — |
| SAML 2.0 | 🔴 Crítico | Gertrudes | Apenas 1 pessoa com ★★★. Bolismar em ★★☆ como backup parcial. |
| OAuth 2.0 / OIDC | ✅ Coberto | Bolismar, Gertrudes | — |
| JWT / Token Validation | ✅ Forte | Francisco, Bolismar, Gertrudes | — |
| RBAC Design | ✅ Forte | Francisco, Bolismar, Gertrudes | — |
| OWASP Top 10 | ✅ Forte | Francisco, Bolismar, Felipe, Gertrudes | — |

### 5.6 DevOps & Infraestrutura

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| Docker / Docker Compose | ✅ Forte | Francisco, Bolismar, Davi | — |
| Kubernetes | ✅ Coberto | Bolismar, Davi | Sem terceiro backup |
| GraalVM Native Image | ✅ Coberto | Bolismar | Davi não cobre (—). Francisco em ★★☆ como suporte |
| CI/CD (GitHub Actions) | ✅ Forte | Francisco, Bolismar, Davi | — |

### 5.7 Cloud & Hosting

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| AWS | ✅ Coberto | Bolismar, Davi | Francisco em ★★☆ como backup |
| Hostinger | ⚠️ Concentrada | Davi | Apenas DevOps cobre |
| Digital-Ocean | ⚠️ Concentrada | Davi | Apenas DevOps cobre |

### 5.8 Qualidade & Testes

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| Testes Unitários | ✅ Forte | Francisco, Bolismar, Felipe, Judith | — |
| Testes de Integração | ✅ Forte | Francisco, Bolismar, Felipe | — |
| Testes E2E (Playwright) | ✅ Forte | Bolismar, Tom, Felipe, Mauro | Mauro contribui com experiência em testes de sistema frontend |
| Testes de Segurança (SAST) | ✅ Coberto | Bolismar, Felipe | — |
| Testes de Performance (JMeter) | ✅ Coberto | Bolismar, Felipe | — |

### 5.9 Observabilidade

| Skill | Cobertura | Profissionais ★★★ | Gaps |
|:---|:---|:---|:---|
| OpenTelemetry | ✅ Coberto | Bolismar, Davi | Francisco em ★★☆ como backup |
| Micrometer | ✅ Coberto | Bolismar, Davi | — |
| Grafana | ⚠️ Concentrada | Davi | Bolismar em ★★☆ como suporte para dashboards |
| Logging (SLF4J/Logback) | ✅ Forte | Francisco, Bolismar, Carlos, Davi, Judith | — |

---

## 6. Gap Analysis: Competências Necessárias vs. Disponíveis

### 6.1 Gaps Críticos (🔴 — Risco Alto)

| Gap | Severidade | Impacto | Mitigação |
|:---|:---:|:---|:---|
| **SAML 2.0 — único especialista** | 🔴 Crítica | Projeto depende de 1 pessoa (Gertrudes) para integração SSO corporativo. Bolismar cobre ★★☆ como backup parcial, mas conhecimento profundo está concentrado. | Capacitar Francisco (Tech Lead) em SAML 2.0 até M2 (15/08/2026). Documentar configurações no ADR. |
| **Frontend até 01/11/2026** | 🔴 Crítica | Tom (Frontend ★★★) só inicia em novembro. Até lá, Bolismar é o único full-stack disponível para frontend. Sem ele, features EP-0001 a EP-0004a atrasam. | Bolismar cobre frontend até outubro. Priorizar componentes reutilizáveis para reduzir retrabalho quando Tom chegar. |
| **Maria Madalena — perfil 100% junior** | 🔴 Crítica | Backend developer junior (★☆☆ em todas as tecnologias). Não entrega tarefas complexas sem supervisão. Representa risco de produtividade para o time backend. | Pareamento obrigatório com Francisco (TL) ou Bolismar (FS) nas primeiras 4 sprints. Tasks designadas devem ser bem delimitadas e revisadas. Plano de evolução: atingir ★★☆ em Java/Spring Boot até M3 (31/08/2026). |

### 6.2 Gaps Moderados (🟡 — Risco Médio)

| Gap | Severidade | Impacto | Mitigação |
|:---|:---:|:---|:---|
| **Davi Silva — hiperespecialista DevOps** | 🟡 Média | DevOps domina infra/deploy/cloud (★★★) mas desconhece totalmente backend (Java, Spring, BD) e frontend. Não consegue fazer troubleshooting de aplicação, apenas de infraestrutura. | OK para o papel. Garantir que incidents de aplicação tenham sempre um dev backend de plantão além do DevOps. Runbooks de troubleshooting devem cobrir a interface DevOps↔Dev. |
| **Carlos Caldas — escopo limitado a banco** | 🟡 Média | Conhecimento exclusivo em PostgreSQL/SQL/Flyway. Não contribui em outras camadas. | OK — manter foco em banco. Garantir que tasks de DB sejam priorizadas para ele. |
| **Felipe Canedas — dependência externa** | 🟡 Média | QA domina seu escopo (★★★) mas precisa de apoio para atividades fora dele (ex: configurar ambiente Docker, analisar logs). | Bolismar ou Francisco como ponto focal de suporte ao QA para infra/testes. |
| **Hostinger / Digital-Ocean — concentrados no DevOps** | 🟡 Média | Apenas Davi conhece esses provedores. Se ambos forem usados em produção, não há redundancy. | Documentar configurações de DNS, rede e deploy para cada provider. Capacitar Bolismar (já tem AWS ★★★) nos outros providers se necessário. |

### 6.3 Gaps Menores (🟢 — Risco Baixo)

| Gap | Severidade | Impacto | Mitigação |
|:---|:---:|:---|:---|
| **Bruno Gratto — perfil conceitual** | 🟢 Baixa | Arquiteto com conhecimento conceitual/arquitetural sólido mas proficiência prática básica. Ideal para decisões de design, não para implementação. | Usar Bruno para revisão de arquitetura, ADRs e code review de design. Não alocar tasks de implementação. |
| **Gertrudes Paiva — hiperespecialista** | 🟢 Baixa | Excelência em IAM mas desconhece outras tecnologias. Perfil esperado para o papel. | Manter foco total em Keycloak/SAML/OAuth. Não requer conhecimento de outras camadas. |
| **Judith Campos — disponibilidade limitada** | 🟢 Baixa | Backend Senior mas disponível apenas 27-31/07/2026 (5 dias). | Alocar para tasks de alto impacto e curta duração: revisão de código, setup de testes, documentação. |
| **Grafana — concentrado no DevOps** | 🟢 Baixa | Apenas Davi com ★★★ em Grafana. Bolismar cobre ★★☆ para dashboards básicos. | OK. Configuração inicial feita pelo DevOps. Time consome dashboards prontos. |

---

## 7. Recomendações

### 7.1 Contratação / Realocação

| # | Recomendação | Prazo | Prioridade |
|:---|:---|:---|:---:|
| R1 | Contratar ou realocar **segundo Frontend Developer** para cobrir o período até 01/11/2026 (Tom) | Imediato | 🔴 Alta |
| R2 | Plano de aceleração para **Maria Madalena**: mentoria dedicada para atingir autonomia (★★☆) em Java/Spring Boot | Sprint 0-4 | 🔴 Alta |

### 7.2 Capacitação

| # | Treinamento | Público | Quando | Duração |
|:---|:---|:---|:---|:---:|
| T1 | Workshop Multi-Tenant Design (modelo lógico, `tenant_id`, Row-Level Security) | Todo time backend + DB | Sprint 0 | 4h |
| T2 | GraalVM Native Image — diagnóstico de erros AOT | Backend team | Sprint 0 | 2h |
| T3 | OpenTelemetry — spans manuais e tracing | Backend team | Sprint 1 | 2h |
| T4 | SAML 2.0 Fundamentals | Francisco (TL) | Sprint 1 | 8h (auto-estudo + pair com Gertrudes) |
| T5 | JMeter/k6 Básico | Maria (BE) | Sprint 2 | 4h |
| T6 | Mentoria Java/Spring Boot — aceleração Maria Madalena | Maria (BE) + Francisco/Bolismar | Sprint 0-4 | 2h/semana (pair programming) |
| T7 | Grafana Dashboards Básico | Bolismar (FS) | Sprint 1 | 2h (pair com Davi) |

### 7.3 Ações Imediatas (Sprint 0)

1. **Tech Lead (Francisco):** Criar template de code review com checklist multi-tenant. Designar tasks iniciais para Maria com escopo bem delimitado.
2. **Full-Stack (Bolismar):** Setup inicial do frontend (Next.js + Tailwind) para EP-0001. Cobrir frontend até chegada de Tom.
3. **IAM (Gertrudes):** Documentar arquitetura SAML 2.0 + OIDC, publicar ADR de autenticação. Parear com Francisco em SAML.
4. **DevOps (Davi):** Configurar infraestrutura local (Docker Compose com Keycloak + PostgreSQL). Auto-instrumentação OpenTelemetry. Setup Grafana para dashboards de monitoramento. Documentar configurações Hostinger/Digital-Ocean.
5. **Arquiteto (Bruno):** Revisão de design multi-tenant. Publicar ADRs de arquitetura. Code review de queries com `tenant_id`.
6. **DB (Carlos):** Definir schema base, configurar Flyway, criar índices para multi-tenant.
7. **QA (Felipe):** Setup Playwright + JMeter. Definir cenários de teste para EP-0001.
8. **BA (Mauro):** Preparar massa de dados de teste. Validar critérios de aceitação do EP-0001.
9. **BE Temp (Judith):** Code review de setup inicial. Contribuir com testes e documentação (27-31/07).

---

## 8. Referências

| Documento | Relação |
|:---|:---|
| [PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY.md) | Nomes, contatos, horas/semana de cada profissional |
| [PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-CAPACITY-EXCEPTIONS.md) | Exceções de capacidade (ausências, reduções) |
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Stack tecnológica completa e decisões de arquitetura |
| [01-PROJECT-CHARTER](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | Escopo, entregas D1-D7, marcos M1-M7 |
| [04-FEATURES](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | 18 funcionalidades, 58 user stories, datas-alvo |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: matriz de skills completa (11 perfis, 8 categorias, 35+ tecnologias), gap analysis com 9 gaps, recomendações de capacitação | Time de Arquitetura |
| 1.1 | 25/07/2026 | Correções pós-validação humana: ajuste de proficiências por profissional (Bolismar expandido, Maria junior ★☆☆, Carlos restrito a DB, Mauro com Playwright ★★★, Bruno conceitual ★☆☆, Gertrudes hiperespecialista, Judith backend senior ★★★, Felipe restrito ao seu escopo). Gap de Multi-Tenant Design resolvido. Adicionado gap crítico de Maria Madalena (100% junior). Ações da Sprint 0 detalhadas por profissional. | Time de Arquitetura |
| 1.2 | 25/07/2026 | Ajuste perfil Davi Silva: hiperespecialista DevOps/Cloud. ★★★ em Docker, Kubernetes, AWS, Hostinger, Digital-Ocean, CI/CD, OpenTelemetry, Micrometer, Grafana, Bash, Logging. Demais tecnologias: sem conhecimento (—). Adicionadas categorias Cloud & Hosting e tecnologias Hostinger, Digital-Ocean, Grafana. Novo gap moderado: Hostinger/Digital-Ocean concentrados. Adicionado T7 (Grafana para Bolismar). | Time de Arquitetura |
| 1.3 | 26/07/2026 | Adição Kong API Gateway: Francisco Oliveira ★★★, Davi Silva ★★★. Bolismar Oliveira ★★☆. Demais sem conhecimento (—). | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Validado e ajustado pelo Time de Arquitetura em 25/07/2026. Status: COMPLIANCE.*

| 3.0 | 28/07/2026 | Renomeado para TEAM-SKILLS-MAP; adicionada Seção 2 — Discovery Team com 9 papéis de Upstream Architecture (TOGAF); seções renumeradas (2→3, 3→4, 4→5, 5→6, 6→7, 7→8); adicionada coluna "Profissional Designado" no mapeamento de papéis | Time de Arquitetura |
