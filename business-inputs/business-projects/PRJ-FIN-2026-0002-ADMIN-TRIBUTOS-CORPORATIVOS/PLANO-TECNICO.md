# PLANO TÉCNICO — Análise de Organização das Demandas Técnicas

- **Projeto:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
- **Programa Pai:** PRJ-FIN-2026-0001 — Adequação Corporativa à Reforma Tributária Nacional
- **Data da Análise:** 08 de Julho de 2026
- **Analistas:** Time de Arquitetos de Solução, Engenheiros de Sistemas, TechLeads, Especialista Fullstack, PM
- **Status:** Análise Inicial — norteará os próximos passos de planejamento fino
- **Versão:** 1.2
- **Última Atualização:** 10 de Julho de 2026 (revisão do sequenciamento com fases de UX/UI, definições transversais, e acréscimo da estrutura de times)

---

## 1. Contexto da Análise

O projeto PRJ-FIN-2026-0002 tem como objetivo central **prover ao time de Finanças um Portal Corporativo de Gestão Tributária**. Para viabilizar tecnicamente essa entrega, três demandas técnicas foram identificadas:

| Demanda | Descrição | Stack Alvo | Dependências |
|:---|:---|:---|:---|
| DT-1 | Novo microserviço CRUD de administração tributária (`ms-billing-admin-tax-rates`) | Java 21 + Spring Boot 4.0.1 | Nenhuma (inaugural) |
| DT-2 | Novo frontend do portal (`web_app-billing-admin-tax-rates`) | React 19 + Vite | Contrato de API com DT-1 |
| DT-3 | Refatoração do microserviço existente (`ms-billing-engine-tax-rates`) para remoção de funcionalidades "admin" | Go + Fiber | DT-1 concluída |

### 1.1 Referências de Arquitetura Java (Fundações para DT-1)

A stack Java do projeto se apoia em dois pilares de arquitetura corporativa já estabelecidos no workspace:

#### 1.1.1 ADR — Decisões Arquiteturais

O diretório [`architecture/adr/java/`](../../../architecture/adr/java/) contém os _Architecture Decision Records_ que fundamentam as escolhas técnicas da stack Java:

| ADR | Decisão | Impacto no DT-1 |
|:---|:---|:---|
| [`0001-runtime-framework-selection.md`](../../../architecture/adr/java/0001-runtime-framework-selection.md) | Estratégia _Dual-Framework_ (Quarkus para microsserviços Cloud Native, Spring Boot para serviços Batch e integrações complexas) | O DT-1 se enquadra no cenário de **Spring Boot**, por se tratar de um microserviço CRUD com integração a banco de dados existente e ecossistema corporativo estabelecido |

#### 1.1.2 Blueprints — Modelos de Referência

O diretório [`architecture/blueprints/backend/java/`](../../../architecture/blueprints/backend/java/) provê os artefatos reutilizáveis que aceleram a criação de novos microserviços:

| Blueprint | Caminho | Uso no DT-1 |
|:---|:---|:---|
| **Parent POM v21** | [`spring/microservices/v21/pom.xml`](../../../architecture/blueprints/backend/java/spring/microservices/v21/pom.xml) | **Obrigatório** — o `pom.xml` do `ms-billing-admin-tax-rates` deve declarar este artefato como `<parent>`. Ele já pré-configura: Java 21, Spring Boot 4.0.1, spring-boot-starter-web, spring-boot-starter-actuator, spring-boot-starter-validation, maven-checkstyle-plugin (Google Checks), JaCoCo 0.8.12, Surefire/Failsafe 3.2.5, e o GraalVM Native Build Tools (`native-maven-plugin`) |
| **Dockerfile JVM** | [`spring/microservices/v21/Dockerfile.jvm.21`](../../../architecture/blueprints/backend/java/spring/microservices/v21/Dockerfile.jvm.21) |Alternativo — build com Eclipse Temurin 21 JDK → runtime JRE 21 |
| **Dockerfile Native** | [`spring/microservices/v21/Dockerfile.native.21`](../../../architecture/blueprints/backend/java/spring/microservices/v21/Dockerfile.native.21) | **Preferencial** — build com GraalVM Native Image Community → runtime distroless cc-debian12. Gera imagem nativa com inicialização sub-segundo e baixíssimo consumo de memória |
| **Checkstyle** | [`generic/Checkstyle.xml`](../../../architecture/blueprints/backend/java/generic/Checkstyle.xml) | Regras de estilo e qualidade de código Java (referenciado pelo parent POM) |
| **EditorConfig** | [`generic/.editorconfig`](../../../architecture/blueprints/backend/java/generic/.editorconfig) | Configuração de indentação e charset para IDEs |
| **Sonar** | [`generic/Sonar-project.yaml`](../../../architecture/blueprints/backend/java/generic/Sonar-project.yaml) | Configuração base para análise SonarQube |

**Regras para o DT-1:**
- O `pom.xml` do `ms-billing-admin-tax-rates` **deve** declarar como parent o artefato `com.fbso.blueprints.backend.java.spring:microservices-parent-v21:1.0.0`
- O build de produção **deve** usar `Dockerfile.native.21` para geração de imagem GraalVM Native Image
- O `Dockerfile.jvm.21` pode ser usado durante desenvolvimento ou se houver incompatibilidade com GraalVM Native Image
- A propriedade `<native.mainClass>` no `pom.xml` do microserviço **precisa** ser definida apontando para a classe `@SpringBootApplication`
- O `Checkstyle.xml` e `.editorconfig` do blueprint genérico devem ser copiados para a raiz do projeto DT-1

---

## 2. Análise das Possibilidades de Organização dos Trabalhos

### 2.1 Cenários de Organização

#### Cenário A — Sequencial Puro (NÃO RECOMENDADO)

```
DT-1 (backend CRUD) → DT-3 (cleanup Go) → DT-2 (frontend)
```

**Vantagem:** Sem risco de conflito de contrato de API — o frontend só começa quando o backend está pronto.  
**Desvantagem:** Tempo total de entrega é a soma dos três. DT-2 (frontend) fica ocioso esperando DT-1 concluir.

#### Cenário B — Backend-First com Frontend em Onda (NÃO RECOMENDADO)

```
DT-1 (backend CRUD) → DT-2 (frontend, inicia após DT-1 definir API)
                    ↘ DT-3 (cleanup Go, após DT-1 em produção)
```

**Vantagem:** DT-3 pode iniciar assim que DT-1 estiver em produção.  
**Desvantagem:** Frontend ainda espera backend definir API para começar.

#### Cenário C — Paralelo com Contrato de API como Ponto de Sincronização (RECOMENDADO)

```
Sprint 0: API Contract + Scaffolds + Criação UX/UI DT-2
Sprint 1: Aprovação UX/UI DT-2 + Definições Transversais DT-1
Sprint 2..N: Desenvolvimento Paralelo DT-1 + DT-2
Sprint N+1..M: Integração + DT-3 Cleanup + Revogação de Grants
Sprint M+1..: Estabilização e Go-Live

        ┌──────────────────────────────────────────────────────┐
        │  FASE 0 (Sprint 0): Fundação                        │
        │  Artefatos: API-CONTRACTS.md, INTEGRATION-MAP.md,    │
        │  OpenAPI YAML, Scaffolds DT-1/DT-2, Criação UX DT-2 │
        └──────────────────────────────────────────────────────┘
                │               │                       │
                ▼               ▼                       ▼
        ┌──────────────┐ ┌──────────────┐ ┌──────────────────────┐
        │  FASE 1      │ │  FASE 1      │ │  FASE 1              │
        │  Sprint 1    │ │  Sprint 1    │ │  Sprint 1            │
        │  DT-1:       │ │  DT-2:       │ │  Times Transversais: │
        │  Definições  │ │  UX Aprovação│ │  Arquitetura, DevOps, │
        │  DB, K8S,    │ │  + Definições│ │  DevSecOps, Infra     │
        │  Segredos    │ │  Arquitetura,│ │  Premissas App/Site   │
        │              │ │  CI/CD, Seg  │ │                       │
        └──────────────┘ └──────────────┘ └──────────────────────┘
                │               │
                ▼               ▼
        ┌──────────────────────────────────────────────────────┐
        │  FASE 2 (Sprints 2..N): Desenvolvimento Paralelo     │
        │  DT-1: Backend Java 21/Spring ◄──MSW Mock──► DT-2:  │
        │  Frontend React 19/Vite (UX aprovada)                │
        └──────────────────────────────────────────────────────┘
                │
                ▼ (após deploy de DT-1 em staging/homolog)
        ┌──────────────────────────────────────────────────────┐
        │  FASE 3 (Sprints N+1..M): Integração + Cleanup       │
        │  DT-1 + DT-2: Integração frontend ↔ backend, E2E     │
        │  DT-3: Remoção funcionalidades admin (Go/Fiber)      │
        │  DT-3: Revogação grants transacionais (mantém SELECT) │
        └──────────────────────────────────────────────────────┘
                │
                ▼
        ┌──────────────────────────────────────────────────────┐
        │  FASE 4 (Sprints M+1..): Estabilização e Go-Live     │
        │  Regressão, Aceitação, Deploy Produção, Cutover       │
        └──────────────────────────────────────────────────────┘
```

**Vantagens:**
- DT-1 e DT-2 avançam em paralelo a partir do Sprint 2, com UX aprovada e contrato de API como "fonte da verdade"
- Sprint 0 concentra todos os scaffolds e o design UX/UI — times de negócio (UX) e técnicos trabalham simultaneamente
- Sprint 1 destrava todas as definições transversais (DB, DevOps, DevSecOps, Arquitetura, Infraestrutura, K8S) para ambos os times simultaneamente
- DT-2 pode usar mock server (MSW) baseado no contrato OpenAPI para desenvolvimento independente desde o Sprint 0
- DT-3 inicia quando DT-1 atinge maturidade suficiente (deploy em staging), sem bloquear as outras frentes
- A revogação de grants transacionais no DT-3 é tratada como parte do deploy de remoção (least-privilege reverso)
- Tempo total de entrega é aproximadamente `2 + max(duração DT-1, duração DT-2) + duração DT-3`

**Desvantagens:**
- Se o contrato de API mudar durante o desenvolvimento, ambos os times precisam se realinhar
- Requer disciplina de governança de contrato (versionamento de API, comunicação entre times)
- Sprint 1 concentra muitas definições transversais — times de Arquitetura, DevOps, DevSecOps e Infraestrutura são compartilhados entre DT-1 e DT-2 e podem se tornar gargalo

#### Cenário D — Três Trilhas Paralelas com Integração Contínua (ALTERNATIVA AVANÇADA)

```
Sprint 0: API Contract + Scaffold dos 3 projetos
Sprint 1+: DT-1, DT-2, DT-3 em paralelo (DT-3 usa feature flags)
```

**Vantagem:** Máximo paralelismo, menor tempo total.  
**Desvantagem:** DT-3 depende tecnicamente de DT-1 estar operacional — risco de retrabalho se DT-1 atrasar. Requer maturidade alta do time.

### 2.2 Recomendação

**Recomenda-se o Cenário C** (Paralelo com API Contract First) pelos seguintes motivos:

1. Melhor equilíbrio entre paralelismo e risco
2. Adequado ao nível de maturidade de times que trabalham com múltiplas stacks (Go, Java, React)
3. O contrato de API como artefato inicial resolve o principal ponto de acoplamento entre DT-1 e DT-2
4. DT-3 inicia com segurança após DT-1 estar funcional em staging

---

## 3. Definição de Artefatos por Solução Técnica

### 3.1 Matriz de Artefatos

Cada solução técnica terá seu próprio conjunto de artefatos, seguindo os padrões já estabelecidos no workspace. A tabela abaixo define **onde** cada artefato reside:

| Artefato | Projeto de Negócio (PRJ-FIN-2026-0002) | DT-1: Backend Java | DT-2: Frontend React | DT-3: Go Refactor |
|:---|:---|:---|:---|:---|
| **PROJECT-CHARTER** | ✅ (já existe) | — | — | — |
| **BUSINESS-REQUIREMENTS** | ✅ (já existe) | — | — | — |
| **EPICS** | ✅ (já existe) | — | — | — |
| **FEATURES** | ✅ (já existe) | — | — | — |
| **USER-STORYS** | ✅ (já existe) | — | — | — |
| **PLANO-TECNICO.md** | ✅ (este documento) | — | — | — |
| **API-CONTRACTS.md** | ✅ (nível projeto) | 🔗 referencia | 🔗 referencia | 🔗 referencia |
| **INTEGRATION-MAP.md** | ✅ (nível projeto) | 🔗 referencia | 🔗 referencia | 🔗 referencia |
| **SPECS.md** | — | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/SPECS.md` | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/SPECS.md` | 🔗 mantém o existente |
| **TASKS.md** | — | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/TASKS.md` | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/TASKS.md` | 🔗 atualiza o existente |
| **TEST_PLAN.md** | — | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/TEST_PLAN.md` | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/TEST_PLAN.md` | 🔗 atualiza o existente |
| **DEFINITION-OF-DONE.md** | ✅ (nível projeto, já existe) | 🔗 referencia | 🔗 referencia | 🔗 referencia |
| **ENVIRONMENTS.md** | — | ✅ `.specs/governance/ENVIRONMENTS.md` | ✅ `.specs/governance/ENVIRONMENTS.md` | 🔗 atualiza o existente |
| **ARCHITECTURE.md** | — | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/ARCHITECTURE.md` | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/ARCHITECTURE.md` | 🔗 atualiza o existente |
| **PRD.md** | — | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/PRD.md` | ✅ `.specs/business-projects/PRJ-FIN-2026-0002/PRD.md` | 🔗 mantém o existente |
| **OpenAPI YAML** | — | ✅ `.specs/api/tax-admin-api.yaml` | 🔗 copia/consome o mesmo | — |

### 3.2 Fundamentação

O padrão atual do workspace estabelece que:

1. **Documentos de negócio** (Project Charter, Épicos, Features, User Stories) residem na pasta `business-inputs/business-projects/PRJ-FIN-XXXX-.../`
2. **Especificações técnicas** residem no `.specs/` de cada solução técnica (microserviço ou frontend)
3. **Artefatos de projeto** (PRD, SPECS, ARCHITECTURE, TASKS, TEST_PLAN) que conectam negócio e técnica residem em `.specs/business-projects/PRJ-FIN-XXXX/` — este padrão foi estabelecido pelo PRJ-FIN-2026-0001 e deve ser mantido

Cada solução técnica (DT-1, DT-2) conterá **exclusivamente** seus próprios artefatos técnicos em seu próprio diretório `.specs/`. DT-3 opera sobre o microserviço existente e apenas **atualiza** os artefatos já existentes.

### 3.3 Árvore de Diretórios Esperada

```
workspace-fbso/
│
├── business-inputs/business-projects/
│   └── PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/
│       ├── 01-PROJECT-CHARTER.md          ← já existe
│       ├── 02-BUSINESS-REQUIREMENTS.md    ← já existe
│       ├── 03-EPICS.md                    ← já existe
│       ├── 04-FEATURES.md                 ← já existe
│       ├── 05-USER-STORYS-*.md            ← já existe (10 arquivos)
│       ├── DEFINITION_OF_DONE.md          ← já existe (nível negócio)
│       ├── MATRIZ-KPI.md                  ← já existe
│       ├── GLOSSARY.md                    ← já existe
│       ├── STAKEHOLDER-MAP.md             ← já existe
│       ├── PLANO-TECNICO.md               ← ESTE DOCUMENTO
│       ├── API-CONTRACTS.md               ← A CRIAR (nível projeto)
│       └── INTEGRATION-MAP.md             ← A CRIAR (nível projeto)
│
├── backend/java/spring/microservices/
│   └── ms-billing-admin-tax-rates/        ← DT-1: A CRIAR
│       ├── pom.xml                         ← parent: microservices-parent-v21:1.0.0
│       ├── Dockerfile                      ← cópia de Dockerfile.native.21 (preferencial)
│       ├── Dockerfile.jvm                  ← cópia de Dockerfile.jvm.21 (alternativo)
│       ├── .editorconfig                   ← cópia do blueprint genérico
│       ├── Checkstyle.xml                  ← cópia do blueprint genérico
│       ├── src/main/java/...
│       ├── src/main/resources/
│       │   └── application.yml
│       ├── src/test/java/...
│       ├── README.md
│       └── .specs/
│           ├── INDEX.md
│           ├── api/
│           │   └── tax-admin-api.yaml     ← OpenAPI do CRUD
│           ├── architecture/
│           │   ├── architecture.md
│           │   ├── c4-context.md
│           │   ├── c4-containers.md
│           │   ├── erd.md
│           │   ├── integrations.md
│           │   └── data-dictionary.md
│           ├── domain/
│           │   └── domain.md
│           ├── engineering/
│           │   ├── code-analysis.md
│           │   └── api-guidelines.md
│           ├── product/
│           │   ├── requirements.md
│           │   └── feature-roadmap.md
│           ├── governance/
│           │   ├── inventory.md
│           │   ├── confidence-report.md
│           │   └── ENVIRONMENTS.md        ← NOVO: ambientes de deploy
│           ├── security/
│           │   └── SECURITY.md
│           └── business-projects/
│               ├── README.md
│               └── PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/
│                   ├── PRD.md
│                   ├── SPECS.md
│                   ├── ARCHITECTURE.md
│                   ├── TASKS.md
│                   └── TEST_PLAN.md
│
├── frontend/javascript/react/web_apps/
│   └── web_app-billing-admin-tax-rates/   ← DT-2: A CRIAR
│       ├── src/...
│       ├── public/...
│       ├── README.md
│       └── .specs/
│           ├── INDEX.md
│           ├── api/
│           │   └── tax-admin-api.yaml     ← cópia do contrato
│           ├── architecture/
│           │   ├── architecture.md
│           │   ├── c4-context.md
│           │   ├── c4-containers.md
│           │   └── integrations.md
│           ├── domain/
│           │   └── domain.md
│           ├── engineering/
│           │   ├── code-analysis.md
│           │   └── api-guidelines.md
│           ├── design/
│           │   ├── DESIGN.md
│           │   ├── design-tokens.md
│           │   └── components.md
│           ├── frontend/
│           │   ├── audit-report.md
│           │   ├── ux-critique.md
│           │   └── ...
│           ├── product/
│           │   ├── requirements.md
│           │   └── feature-roadmap.md
│           ├── governance/
│           │   ├── inventory.md
│           │   ├── confidence-report.md
│           │   └── ENVIRONMENTS.md        ← NOVO: ambientes de deploy
│           └── business-projects/
│               ├── README.md
│               └── PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/
│                   ├── PRD.md
│                   ├── SPECS.md
│                   ├── ARCHITECTURE.md
│                   ├── TASKS.md
│                   └── TEST_PLAN.md
│
└── backend/go/fiber/microservices/
    └── ms-billing-engine-tax-rates/       ← DT-3: ALTERAR (existente)
        └── .specs/
            ├── INDEX.md                   ← atualizar removendo refs a "admin"
            ├── api/
            │   └── tax-rates-api.yaml     ← remover endpoints admin
            ├── architecture/
            │   └── ...                    ← atualizar docs de integração
            ├── business-projects/
            │   └── PRJ-FIN-2026-0001-.../ ← mantido (histórico)
            └── ...
```

---

## 4. Definição da Localização dos Artefatos Transversais

### 4.1 API-CONTRACTS.md

**Localização:** `business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/API-CONTRACTS.md`

**Justificativa:**
- O contrato de API entre DT-1 (backend) e DT-2 (frontend) é um artefato de **nível projeto**, não de uma solução técnica específica
- Ele define a "fonte da verdade" que ambos os times consomem
- Centralizá-lo no projeto de negócio evita duplicação e divergência entre os contratos nas duas soluções
- O OpenAPI YAML técnico (`tax-admin-api.yaml`) reside em cada solução como **cópia derivada**, mas o documento descritivo e as decisões de design da API ficam no projeto

**Conteúdo esperado:**
- Visão geral dos endpoints REST (recursos, métodos, paths)
- Modelos de request/response (schemas)
- Regras de autenticação/autorização por endpoint
- Estratégia de versionamento (URI vs header)
- Política de erros e códigos HTTP
- Exemplos de fluxos de consumo (frontend → backend)

### 4.2 INTEGRATION-MAP.md

**Localização:** `business-inputs/business-projects/PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS/INTEGRATION-MAP.md`

**Justificativa:**
- O mapa de integrações abrange **múltiplas soluções técnicas**: como o frontend se comunica com o backend, como o backend se conecta ao banco de dados, como o novo microserviço (DT-1) se relaciona com o existente (DT-3 após refatoração)
- É um artefato de visão sistêmica, essencial para arquitetos e tech leads
- Colocá-lo em uma única solução técnica (ex: apenas no backend) deixaria o time de frontend sem visibilidade oficial das dependências

**Conteúdo esperado:**
- Diagrama de comunicação entre todos os componentes do sistema
- Protocolos e formatos de cada integração (REST/JSON, gRPC, mensageria, etc.)
- Dependências externas (bancos de dados, serviços de autenticação, etc.)
- Direção e cardinalidade de cada integração
- Requisitos de segurança por canal de comunicação
- SLAs e timeouts esperados

### 4.3 OpenAPI YAML (tax-admin-api.yaml)

**Localização primária (fonte da verdade):** `.specs/api/tax-admin-api.yaml` no microserviço DT-1 (`ms-billing-admin-tax-rates`)

**Localização secundária (cópia de consumo):** `.specs/api/tax-admin-api.yaml` no frontend DT-2 (`web_app-billing-admin-tax-rates`)

**Justificativa:**
- O backend é o **provedor** da API — portanto, detém a fonte canônica
- O frontend mantém uma cópia para geração de mocks e validação de tipos (ex: gerar tipos TypeScript a partir do OpenAPI)
- O documento `API-CONTRACTS.md` no nível do projeto referencia ambos e define o processo de atualização (ex: "toda alteração no OpenAPI do backend deve ser refletida no frontend em até 1 sprint")

---

## 5. Governança dos Artefatos

### 5.1 Fluxo de Atualização

```
Mudança de requisito de negócio
        │
        ▼
[Projeto] API-CONTRACTS.md atualizado (PM + Arquiteto)
        │
        ▼
[DT-1] OpenAPI YAML atualizado (Time Backend)
        │
        ├──► [DT-1] Código Java implementa novo endpoint/campo
        │
        └──► [DT-2] OpenAPI YAML copiado → Tipos TypeScript regenerados → Código React atualizado
```

### 5.2 Responsáveis por Artefato

| Artefato | Responsável Primário | Revisor | Aprovador |
|:---|:---|:---|:---|
| PLANO-TECNICO.md | Arquiteto de Solução | Tech Leads, Engenheiros | PM + Arquiteto |
| API-CONTRACTS.md | Arquiteto de Solução + Tech Lead Backend | Tech Lead Frontend | PM |
| INTEGRATION-MAP.md | Arquiteto de Solução | Engenheiros de Sistemas | Tech Leads |
| SPECS.md (por solução) | Tech Lead da solução | Arquiteto | PM |
| TASKS.md (por solução) | Tech Lead da solução | PM | — |
| TEST_PLAN.md (por solução) | QA + Tech Lead | Arquiteto | PM |
| ENVIRONMENTS.md (por solução) | DevOps / Engenheiro | Tech Lead | Arquiteto |
| ARCHITECTURE.md (por solução) | Arquiteto de Solução | Tech Leads | — |
| OpenAPI YAML | Tech Lead Backend (DT-1) | Tech Lead Frontend (DT-2) | Arquiteto |

### 5.3 Versionamento

- Todos os artefatos seguem versionamento semântico (MAJOR.MINOR) no frontmatter
- O `CHANGELOG.md` de cada solução registra as alterações
- O `API-CONTRACTS.md` referencia a versão exata do OpenAPI YAML que está em vigor

---

## 6. Marcos e Sequenciamento Proposto

### 6.1 Fase 0 — Fundação (Sprint 0)

| Atividade | Responsável | Entregável |
|:---|:---|:---|
| Criação do PLANO-TECNICO.md | Arquiteto + Time | Este documento |
| Definição do API-CONTRACTS.md | Arquiteto + Tech Leads | Contrato de API aprovado |
| Definição do INTEGRATION-MAP.md | Arquiteto + Engenheiros | Mapa de integrações |
| Criação do OpenAPI YAML inicial | Tech Lead Backend | `tax-admin-api.yaml` v1.0 |
| Scaffold do projeto Java/Spring (DT-1) — `pom.xml` com parent `microservices-parent-v21:1.0.0` | Time Backend | Estrutura Maven + `native.mainClass` definida |
| Scaffold do projeto Java/Spring (DT-1) — copiar `Dockerfile.native.21` do blueprint v21 como `Dockerfile` | Time Backend | Dockerfile Native (preferencial) |
| Scaffold do projeto Java/Spring (DT-1) — copiar `Dockerfile.jvm.21` do blueprint v21 como `Dockerfile.jvm` | Time Backend | Dockerfile JVM (alternativo/dev) |
| Scaffold do projeto Java/Spring (DT-1) — copiar `.editorconfig` e `Checkstyle.xml` dos blueprints genéricos | Time Backend | Qualidade de código padronizada |
| Scaffold do projeto Java/Spring (DT-1) — criar estrutura `.specs/` conforme Seção 3.3 | Time Backend | Documentação técnica |
| Scaffold do projeto React/Vite (DT-2) | Time Frontend | Estrutura inicial + CI/CD |
| Setup de mock server para DT-2 (MSW local) | Time Frontend | Mock baseado no OpenAPI YAML |
| **Criação do design UX/UI do portal (DT-2)** — layout, modelagem, prototipação | **Time UX/UI** | **Wireframes, protótipos, design system inicial** |
| **Reunião conjunta de alinhamento DT-1 + DT-2** — visão geral do projeto, escopo, cronograma, dependências entre times | **PM + PO + Tech Leads + Times** | Ata de alinhamento, entendimento compartilhado, matriz de comunicação |

**`★ Insight ─────────────────────────────────────`**
- Todas as atividades do Sprint 0 são executadas por **times diferentes** (Arquitetura, Backend Java, Frontend React, UX/UI) — não há contenção de recursos
- A **reunião conjunta DT-1 + DT-2** é crítica para estabelecer um entendimento compartilhado do projeto desde o início, reduzindo ruídos de comunicação e prevenindo desalinhamentos futuros
- O time de UX/UI **não depende do API Contract** para iniciar o design visual — os requisitos de negócio (FEATURES.md, USER-STORYS.md) já são suficientes
- O risco é o time de Arquitetura estar sobrecarregado (API Contract + INTEGRATION-MAP + apoiar scaffolds simultaneamente)
- **Recomendação:** Agendar já no Sprint 0 as sessões de definição com os times transversais para o Sprint 1, garantindo disponibilidade de DBA, DevOps, DevSecOps e Infraestrutura
`─────────────────────────────────────────────────`

### 6.2 Fase 1 — Design e Definições Transversais (Sprint 1)

| Atividade | Responsável | Entregável |
|:---|:---|:---|
| **DT-2:** Montar apresentação do design UX/UI e obter aprovação formal | Time UX/UI + PO + Stakeholders | Design aprovado |
| **DT-2:** Definição com Arquitetura sobre desenho arquitetural da solução frontend | Time Frontend + Arquiteto | Desenho arquitetural frontend |
| **DT-2:** Definição com DevOps sobre processos de CI/CD | Time Frontend + DevOps | Pipeline CI/CD definido |
| **DT-2:** Definição com DevSecOps sobre processos de segurança | Time Frontend + DevSecOps | Requisitos de segurança mapeados |
| **DT-2:** Definição com Infraestrutura sobre plataforma de hospedagem | Time Frontend + Infra | Plataforma definida |
| **DT-1:** Definição com time de Banco de Dados — criação de usuário, grants, acessos | Time Backend + DBA | Acessos ao banco provisionados |
| **DT-1:** Definição com DevOps e DevSecOps — obtenção de senhas e chaves para acesso ao banco | Time Backend + DevOps + DevSecOps | Estratégia de segredos definida (Vault/Secrets Manager) |
| **DT-1:** Definição com Arquitetura sobre desenho arquitetural da solução | Time Backend + Arquiteto | Desenho arquitetural backend |
| **DT-1:** Definição com DevOps, DevSecOps e Infraestrutura — hospedagem do microserviço em K8S | Time Backend + DevOps + DevSecOps + Infra | Configuração K8S definida |
| **Times Transversais:** Definição entre Arquitetura, DevOps, DevSecOps e Infraestrutura sobre premissas de criação e existência da nova aplicação e site corporativo | Arquitetura + DevOps + DevSecOps + Infra | Documento de premissas da plataforma |

> ⚠️ **Atenção:** O Sprint 1 concentra muitas definições com times transversais. Os times de DBA, DevOps, DevSecOps, Infraestrutura e Arquitetura são **compartilhados entre DT-1 e DT-2**. Recomenda-se:
> - Agendar as sessões com antecedência (idealmente durante o Sprint 0)
> - Consolidar sessões quando possível (ex: uma única reunião de Arquitetura cobrindo desenhos de DT-1 e DT-2)
> - Se houver conflito de agenda, priorizar as definições que bloqueiam o Sprint 2

### 6.3 Fase 2 — Desenvolvimento Paralelo (Sprints 2..N)

| Stream | Time | Foco |
|:---|:---|:---|
| DT-1 Backend | Time Java/Spring | Implementar endpoints CRUD conforme OpenAPI, integração com banco de dados, autenticação SAML 2.0 via Keycloak |
| DT-2 Frontend | Time React + UX/UI | Implementar módulos do portal consumindo mock API (MSW) → migrar para API real quando disponível |

**Pré-condições para iniciar a Fase 2:**
- ✅ UX/UI do DT-2 aprovada (Sprint 1)
- ✅ Definições transversais do DT-1 concluídas (Sprint 1)
- ✅ Acessos ao banco de dados provisionados (Sprint 1)
- ✅ Contrato de API (OpenAPI) estável

**Ponto de sincronização:** Ao final de cada sprint, verificar aderência ao contrato de API. O mock server (MSW) é regenerado automaticamente a partir do OpenAPI YAML se houver alterações.

### 6.4 Fase 3 — Integração, Cleanup e Revogação (Sprints N+1..M)

| Stream | Time | Foco |
|:---|:---|:---|
| DT-1 + DT-2 | Ambos os times | Integração frontend ↔ backend real, testes E2E, ajustes de contrato |
| DT-3 Cleanup | Time Go/Fiber | Remover funcionalidades admin do `ms-billing-engine-tax-rates` |
| DT-3 Revogação | DevOps + DevSecOps + Infra + DBA | Retirada dos acessos administrativos do microserviço existente (ver Seção 6.4.1) |

#### 6.4.1 Revogação de Grants — DT-3 (Processo de Least-Privilege Reverso)

No deploy de remoção das operações de administração do `ms-billing-engine-tax-rates`, o time de Infraestrutura, DevOps, DevSecOps e DBA devem executar as seguintes tarefas de revogação:

| Ação | Responsável | Descrição |
|:---|:---|:---|
| Remover grants transacionais | DBA + DevOps | Revogar permissões de `INSERT`, `UPDATE`, `DELETE` nas tabelas administrativas para o usuário do motor Go |
| Manter grant de `SELECT` | DBA | Preservar permissão de leitura nas tabelas administrativas (motor antigo torna-se somente consulta) |
| Remover acessos administrativos do microserviço | Infra + DevOps | Remover roles/permissões de admin do serviço no K8S |
| Auditoria de segurança | DevSecOps | Validar que o serviço não retém acessos além do necessário após a revogação |
| Atualizar documentação de acessos | DBA + DevSecOps | Registrar o novo perfil de acessos (read-only) na documentação de segurança |

### 6.5 Fase 4 — Estabilização e Go-Live (Sprints M+1..)

- Testes de regressão no motor de cálculo (DT-3) — garantir que apenas operações de leitura permanecem
- Testes de aceitação do portal (DT-1 + DT-2)
- Homologação pelo time de negócio (Finanças)
- Deploy em produção do novo portal
- Cutover: time de Finanças migra para o novo portal
- Desativação das rotas admin no motor antigo
- Validação pós-go-live: monitoramento, logs, rollback plan

---

## 7. Riscos Identificados

| Risco | Probabilidade | Impacto | Mitigação |
|:---|:---|:---|:---|
| Alteração no contrato de API durante desenvolvimento | Média | Alto | API-CONTRACTS.md aprovado por ambos os times; mudanças seguem processo formal de amendment |
| DT-1 atrasar e bloquear DT-3 | Média | Médio | DT-3 pode iniciar preparação (identificar código a remover) antes de DT-1 estar pronto |
| Divergência entre mock e API real | Média | Médio | Mock server (MSW) gerado automaticamente a partir do OpenAPI YAML |
| Conflito de schemas entre DB do DT-1 e DB do motor existente | Alta | Alto | INTEGRATION-MAP.md deve definir claramente quais tabelas são compartilhadas e quais são exclusivas |
| Complexidade do Período Híbrido (2029-2032) não ser capturada no contrato de API | Baixa | Alto | Incluir cenários de Período Híbrido nos exemplos do API-CONTRACTS.md |
| Indisponibilidade de times transversais (DBA, DevOps, DevSecOps, Infra, Arquitetura) no Sprint 1 | Média | Alto | Agendar sessões de definição com antecedência durante o Sprint 0; consolidar reuniões de DT-1 e DT-2 quando possível (ex: única sessão de Arquitetura); PM deve garantir prioridade na agenda dos times transversais |
| Atraso na aprovação do design UX/UI (DT-2) | Média | Alto | Envolver PO e stakeholders desde o Sprint 0 com checkpoints intermediários; o design usa os mesmos requisitos de negócio já validados; ter critérios de aprovação claros definidos antes da apresentação formal |
| Times de DT-1 e DT-2 com velocidades diferentes causando desalinhamento | Baixa | Médio | Sincronização semanal entre Tech Leads; API Contract estável como ponto de ancoragem; integração contínua desde o primeiro endpoint disponível |
| Regressão no motor de cálculo (DT-3) após remoção de funcionalidades admin | Baixa | Alto | Testes de regressão automatizados antes e depois do cleanup; deploy em staging primeiro; rollback plan documentado |
| Revogação incorreta de grants no DT-3 (remoção de SELECT indevido) | Baixa | Alto | Checklist de revogação revisado por DBA e DevSecOps; validação em ambiente de staging antes de produção; dupla checagem dos grants mantidos |

---

## 8. Perguntas em Aberto

> **Status: TODAS RESPONDIDAS em 08/07/2026** ✅

1. **Banco de dados:** DT-1 (`ms-billing-admin-tax-rates`) usará o mesmo banco de dados do motor de cálculo (`ms-billing-engine-tax-rates`) ou terá seu próprio banco? Se compartilhar, como garantir isolamento de responsabilidades (admin vs engine)?
   - **Resposta:** Será o mesmo banco. Como o compartilhamento será durante a fase de convivência que tende a ser muito rápida os riscos serão mínimos, e não teremos que nos preocupar em realizar sincronização de base de dados.

2. **Autenticação/Autorização:** O portal usará o mesmo mecanismo de autenticação do ecossistema corporativo? Qual o protocolo (OAuth2, JWT, SAML)?
   - **Resposta:** Usaremos SAML 2.0 através do Keycloak.

3. **Padrão de nomenclatura de artefatos na stack Java:** O microserviço Java existente (`ms-product-catalog-admin-simple`) usa `docs/` em vez de `.specs/`. Devemos padronizar para `.specs/` no novo microserviço, mantendo coerência com o ecossistema Go e React?
   - **Resposta:** Usaremos `.specs/`.

4. **ENVIRONMENTS.md:** Quais ambientes serão provisionados? Dev, Staging, Production? Haverá ambiente de teste integrado (DT-1 + DT-2)?
   - **Resposta:** Ambientes Dev, Staging, Production provisionados via Kubernetes. Para desenvolvimento local será usado `docker` e `docker compose`.

5. **Estratégia de Mock:** O mock server do frontend será local (MSW) ou compartilhado (WireMock em servidor)? Impacta o pipeline de CI/CD.
   - **Resposta:** Será local (MSW — Mock Service Worker).

6. **Migração de dados:** As tabelas fiscais atualmente gerenciadas via operações admin no motor Go precisarão ser migradas para o novo banco/ schema do microserviço Java? Ou o novo microserviço apenas passará a escrever nas mesmas tabelas?
   - **Resposta:** Escreverá nas mesmas tabelas.

---

## 9. Próximos Passos

1. **[Imediato]** Revisar e aprovar este PLANO-TECNICO.md com todos os stakeholders
2. **[Imediato]** ~~Responder às 6 perguntas em aberto~~ ✅ **CONCLUÍDO** — Seção 8 (08/07/2026)
3. **[Imediato]** Validar e aprovar a estrutura de times proposta (Seção 11)
4. **[Pré-Sprint 0]** Agendar sessões de definição com times transversais (DBA, DevOps, DevSecOps, Infra, Arquitetura) para o Sprint 1
5. **[Sprint 0]** Criar `API-CONTRACTS.md` no nível do projeto
6. **[Sprint 0]** Criar `INTEGRATION-MAP.md` no nível do projeto
7. **[Sprint 0]** Criar `tax-admin-api.yaml` (OpenAPI) com endpoints CRUD essenciais
8. **[Sprint 0]** Scaffold do `ms-billing-admin-tax-rates` (Java 21 + Spring Boot 4.0.1):
   - Criar `pom.xml` com `<parent>` apontando para `com.fbso.blueprints.backend.java.spring:microservices-parent-v21:1.0.0`
   - Definir `<native.mainClass>` com a classe `@SpringBootApplication`
   - Copiar `Dockerfile.native.21` do blueprint v21 como `Dockerfile` (build de produção)
   - Copiar `Dockerfile.jvm.21` do blueprint v21 como `Dockerfile.jvm` (alternativa dev)
   - Copiar `.editorconfig` e `Checkstyle.xml` dos blueprints genéricos
   - Criar estrutura `.specs/` completa conforme Seção 3.3
9. **[Sprint 0]** Scaffold do `web_app-billing-admin-tax-rates` (React 19 + Vite)
10. **[Sprint 0]** Setup de mock server (MSW) baseado no OpenAPI YAML
11. **[Sprint 0]** Time de UX/UI: criar wireframes, protótipos e design system inicial do portal
12. **[Sprint 1]** Time de UX/UI: apresentar design e obter aprovação formal (PO + Stakeholders)
13. **[Sprint 1]** Executar todas as definições transversais de DT-1 e DT-2 (DB, K8S, CI/CD, Segurança, Arquitetura)
14. **[Sprint 2+]** Iniciar desenvolvimento paralelo DT-1 e DT-2 com UX aprovada e mock server
15. **[Sprint N+1]** Iniciar DT-3 (cleanup + revogação de grants) após DT-1 em staging
16. **[Sprint M+1]** Fase de estabilização, homologação e go-live

---

## 10. Aprovações

| Papel | Nome | Data | Assinatura |
|:---|:---|:---|:---|
| Product Manager |  |  |  |
| Arquiteto de Solução |  |  |  |
| Tech Lead Backend (Java) |  |  |  |
| Tech Lead Frontend (React) |  |  |  |
| Tech Lead Motor (Go) |  |  |  |
| Engenheiro de Sistemas |  |  |  |

---

## 11. Estrutura Proposta de Times

### 11.1 Visão Geral

O projeto demanda times multidisciplinares organizados em três frentes de desenvolvimento (DT-1, DT-2, DT-3), apoiadas por times transversais, gestão e qualidade. A tabela abaixo apresenta a lotação proposta:

| Papel | Frente | Qtd. | Sprint 0 | Sprint 1 | Sprints 2..N | Sprints N+1..M | Sprints M+1.. |
|:---|:---|:---:|:---:|:---:|:---:|:---:|:---:|
| **Product Manager (PM)** | Gestão | 1 | 100% | 100% | 50% | 50% | 50% |
| **Product Owner (PO)** | Gestão | 1 | 50% | 100% | 50% | 50% | 50% |
| **Tech Lead Java/Spring** | DT-1 | 1 | 100% | 100% | 100% | 80% | 50% |
| **Dev Java/Spring Sênior** | DT-1 | 2 | 80% | 80% | 100% | 100% | 50% |
| **Dev Java/Spring Pleno** | DT-1 | 1-2 | 50% | 50% | 100% | 100% | 50% |
| **Tech Lead Frontend/React** | DT-2 | 1 | 100% | 100% | 100% | 80% | 50% |
| **Dev React Sênior** | DT-2 | 1-2 | 80% | 50% | 100% | 100% | 50% |
| **Dev React Pleno** | DT-2 | 1 | 50% | 50% | 100% | 100% | 50% |
| **UX/UI Designer** | DT-2 | 1-2 | 100% | 100% | 30%¹ | 20%¹ | — |
| **Tech Lead Go/Fiber** | DT-3 | 1 | — | — | 20%² | 100% | 80% |
| **Dev Go/Fiber Sênior** | DT-3 | 1 | — | — | — | 100% | 80% |
| **Arquiteto de Solução** | Transversal | 1 | 100% | 100% | 30% | 30% | 20% |
| **DBA** | Transversal | 1 | — | 80% | 20% | 50% | 20% |
| **DevOps Engineer** | Transversal | 1 | 30% | 100%³ | 30% | 80% | 50% |
| **DevSecOps Engineer** | Transversal | 1 | 20% | 100%³ | 20% | 80% | 50% |
| **Infraestrutura/K8S Engineer** | Transversal | 1 | 20% | 100%³ | 20% | 50% | 50% |
| **QA / Test Engineer** | Qualidade | 1-2 | — | 30% | 50% | 100% | 80% |
| **Time de Homologação (Negócio)** | Qualidade | 1-2 | — | 20% | 20% | 30% | 100% |

> ¹ UX/UI Designer mantém participação reduzida durante o desenvolvimento para refinamentos e ajustes de design.  
> ² Tech Lead Go participa de alinhamentos sobre o escopo da refatoração durante o desenvolvimento do DT-1.  
> ³ Times transversais com pico de dedicação no Sprint 1 devido à concentração de definições (ver Seção 6.2).

### 11.2 Skills por Frente

> **Legenda da Notação de Proficiência:**
> 
> | Notação | Nível | Descrição |
> |:---:|:---|:---|
> | ★★★ | **Avançado / Autônomo** | Domina a tecnologia com profundidade. Capaz de definir padrões, tomar decisões arquiteturais, resolver problemas complexos e mentorar outros membros do time. **Atua como referência técnica na skill.** |
> | ★★☆ | **Intermediário / Produtivo** | Utiliza a tecnologia com fluência no dia a dia. Consegue implementar tarefas típicas com autonomia, mas pode precisar de apoio em cenários complexos ou atípicos. |
> | ★☆☆ | **Básico / Assistido** | Conhecimento fundamental. Consegue trabalhar com supervisão ou consultando documentação. Não atua de forma independente em decisões de design ou problemas complexos. |
> | — | **Não requerido** | A skill não é necessária para este papel. O profissional não precisa ter conhecimento neste domínio. |

#### 11.2.1 DT-1 — Backend Java/Spring (`ms-billing-admin-tax-rates`)

| Skill | Tech Lead | Dev Sênior | Dev Pleno |
|:---|:---:|:---:|:---:|
| **Java 21** | ★★★ | ★★★ | ★★☆ |
| **Spring Boot 4.0.1** (Web, Actuator, Validation) | ★★★ | ★★★ | ★★☆ |
| **Spring Data JDBC / JDBC Template** | ★★★ | ★★★ | ★★☆ |
| **PostgreSQL** (modelagem, queries, migrations Flyway) | ★★★ | ★★☆ | ★★☆ |
| **GraalVM Native Image** (build, troubleshooting) | ★★★ | ★★☆ | ★☆☆ |
| **Docker / Kubernetes** | ★★☆ | ★★☆ | ★☆☆ |
| **Keycloak / SAML 2.0** | ★★★ | ★★☆ | ★☆☆ |
| **OpenAPI / Swagger** (contrato, documentação) | ★★★ | ★★☆ | ★★☆ |
| **Maven** (parent POM, plugins, multi-module) | ★★★ | ★★★ | ★★☆ |
| **Testes** (JUnit 5, Mockito, Testcontainers, JaCoCo) | ★★★ | ★★★ | ★★☆ |
| **Checkstyle / Google Checks** | ★★☆ | ★★☆ | ★★☆ |

#### 11.2.2 DT-2 — Frontend React (`web_app-billing-admin-tax-rates`)

| Skill | Tech Lead | Dev Sênior | Dev Pleno | UX/UI Designer |
|:---|:---:|:---:|:---:|:---:|
| **React 19** (hooks, context, concurrent features) | ★★★ | ★★★ | ★★☆ | — |
| **TypeScript** | ★★★ | ★★★ | ★★☆ | — |
| **Vite** (build tooling, configuração) | ★★★ | ★★☆ | ★★☆ | — |
| **Design System / Component Library** | ★★☆ | ★★☆ | ★☆☆ | ★★★ |
| **MSW (Mock Service Worker)** | ★★★ | ★★☆ | ★★☆ | — |
| **OpenAPI / Geração de Tipos TypeScript** | ★★★ | ★★☆ | ★★☆ | — |
| **Consumo de APIs REST** (fetch, axios, react-query) | ★★★ | ★★★ | ★★☆ | — |
| **Autenticação SAML 2.0 / Keycloak** (frontend) | ★★☆ | ★★☆ | ★☆☆ | — |
| **Testes** (Vitest, React Testing Library, Playwright) | ★★★ | ★★★ | ★★☆ | — |
| **Acessibilidade (WCAG 2.1 AA)** | ★★☆ | ★★☆ | ★☆☆ | ★★★ |
| **Figma / Prototipação** | ★☆☆ | — | — | ★★★ |
| **UX Research / Testes de Usabilidade** | — | — | — | ★★★ |
| **CSS / Design Tokens / Temas** | ★★☆ | ★★☆ | ★★☆ | ★★★ |
| **DevOps Frontend** (CI/CD pipelines, deploy estático) | ★★☆ | ★★☆ | ★☆☆ | — |

#### 11.2.3 DT-3 — Refatoração Go/Fiber (`ms-billing-engine-tax-rates`)

| Skill | Tech Lead | Dev Sênior |
|:---|:---:|:---:|
| **Go** (1.21+) | ★★★ | ★★★ |
| **Fiber Framework** | ★★★ | ★★★ |
| **PostgreSQL** (consultas, migrations) | ★★★ | ★★☆ |
| **Refatoração / Análise de Código Legado** | ★★★ | ★★☆ |
| **Testes de Regressão** (Go testing, testify) | ★★★ | ★★☆ |
| **Docker / Kubernetes** | ★★☆ | ★★☆ |
| **OpenAPI / Contratos de API** | ★★☆ | ★★☆ |
| **Grants e Permissionamento DB** | ★★☆ | ★★☆ |

#### 11.2.4 Times Transversais

| Skill | Arquiteto | DBA | DevOps | DevSecOps | Infra/K8S |
|:---|:---:|:---:|:---:|:---:|:---:|
| **Arquitetura de Microsserviços** | ★★★ | — | ★★☆ | ★★☆ | ★★☆ |
| **Desenho de APIs REST / Contratos** | ★★★ | — | ★★☆ | — | — |
| **Modelagem de Dados / PostgreSQL** | ★★☆ | ★★★ | — | — | — |
| **Grants, Roles e Segurança de DB** | ★★☆ | ★★★ | ★★☆ | ★★★ | — |
| **Kubernetes (K8S)** | ★★☆ | — | ★★★ | ★★☆ | ★★★ |
| **CI/CD Pipelines** | ★★☆ | — | ★★★ | ★★☆ | ★★☆ |
| **Gestão de Segredos (Vault/Secrets Manager)** | ★★☆ | — | ★★★ | ★★★ | — |
| **Segurança de Aplicações (SAST, SCA)** | ★★☆ | — | ★★☆ | ★★★ | — |
| **SAML 2.0 / Keycloak** | ★★★ | — | ★★☆ | ★★★ | — |
| **Docker / Containers** | ★★☆ | — | ★★★ | ★★☆ | ★★★ |
| **GraalVM Native Image** | ★★☆ | — | ★★☆ | — | ★★☆ |
| **Infraestrutura como Código (IaC)** | ★★☆ | — | ★★★ | ★★☆ | ★★★ |
| **Monitoramento e Observabilidade** | ★★☆ | ★☆☆ | ★★★ | ★★☆ | ★★★ |

#### 11.2.5 Gestão e Qualidade

| Skill | PM | PO | QA/Test | Homologação |
|:---|:---:|:---:|:---:|:---:|
| **Gestão de Projetos Ágeis (Scrum/Kanban)** | ★★★ | ★★☆ | ★★☆ | — |
| **Domínio de Negócio Tributário** | ★★☆ | ★★★ | ★☆☆ | ★★★ |
| **Gestão de Stakeholders** | ★★★ | ★★★ | — | — |
| **Testes Manuais / Exploratórios** | — | — | ★★★ | ★★★ |
| **Testes Automatizados (E2E)** | — | — | ★★★ | ★☆☆ |
| **Testes de Regressão** | — | — | ★★★ | ★★★ |
| **Homologação / UAT** | — | ★★☆ | ★★☆ | ★★★ |
| **Análise de Requisitos / Critérios de Aceite** | ★★☆ | ★★★ | ★★☆ | ★★★ |
| **Comunicação e Facilitação** | ★★★ | ★★★ | ★☆☆ | ★☆☆ |

### 11.3 Dimensionamento Total

| Categoria | Papéis | Headcount Estimado |
|:---|:---|:---:|
| **Gestão** | PM, PO | 2 |
| **DT-1 (Backend Java)** | Tech Lead, 2 Devs Sênior, 1-2 Devs Pleno | 4-5 |
| **DT-2 (Frontend React)** | Tech Lead, 1-2 Devs Sênior, 1 Dev Pleno, 1-2 UX/UI | 4-6 |
| **DT-3 (Go Refactor)** | Tech Lead, 1 Dev Sênior | 2 |
| **Times Transversais** | Arquiteto, DBA, DevOps, DevSecOps, Infra/K8S | 5 |
| **Qualidade** | QA/Test, Homologação | 2-4 |
| **Total** | | **19-24** |

> **Nota:** Os times transversais (Arquitetura, DBA, DevOps, DevSecOps, Infra) são **compartilhados com outros projetos** da organização. O headcount indicado reflete a dedicação parcial necessária para este projeto, não necessariamente contratações exclusivas.

### 11.4 Premissas de Composição dos Times

1. **Times de DT-1 e DT-2 são independentes** — composto por pessoas diferentes, permitindo paralelismo real a partir do Sprint 2
2. **Tech Leads são dedicados em tempo integral** durante as fases de desenvolvimento — um por frente técnica
3. **Times transversais são compartilhados** — DBA, DevOps, DevSecOps, Infraestrutura e Arquitetura atendem múltiplos projetos; o pico de dedicação ocorre no Sprint 1 (definições) e nos Sprints N+1..M (revogação de grants, deploy)
4. **Time de DT-3 é acionado sob demanda** — o Tech Lead Go participa de alinhamentos durante a Fase 2 (Sprint 2..N), mas a lotação completa só ocorre na Fase 3
5. **UX/UI Designer tem dedicação concentrada** — 100% nos Sprints 0-1, depois reduz para refinamentos pontuais durante o desenvolvimento
6. **Time de Homologação (negócio)** — idealmente composto por usuários-chave do time de Finanças, com dedicação mais intensa na Fase 4 (Go-Live)

---

🤖 *Análise gerada com apoio de Claude Code (Anthropic), em 08 de Julho de 2026, como etapa preparatória para o planejamento fino das demandas técnicas DT-1, DT-2 e DT-3.*
