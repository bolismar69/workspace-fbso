
### 1\. O Blueprint Java (`/architecture/blueprints/java/`)

Este blueprint não contém código executável. Ele define a **estrutura de pastas (Clean Architecture)** e as configurações globais de qualidade que todos devem seguir.

### 🏗️ Hierarquia de Blueprints Java (`/architecture/blueprints/java`)

```text
workspace-fbso/
└── architecture/
    └── blueprints/
        └── java/               <-- Ponto de entrada (README.md aqui)
            └── java-generic/   <-- DNA (Estrutura de pastas e padrões)
                ├── .editorconfig            # Padronização de indentação/espaços
                ├── checkstyle.xml           # Regras de estilo de código Java
                ├── pom-base.xml             # Definição de versões de libs comuns (Lombok, JUnit)
                └── src/
                    └── main/java/com/fbso/core/
                        ├── domain/          # Entidades e Regras de Negócio (Puro Java)
                        ├── application/     # Casos de Uso (Interfaces)
                        ├── infrastructure/  # Adaptadores (DB, APIs externas)
                        └── interfaces/      # Controllers / Entrypoints
            ├── java-quarkus/   <-- Especialização para Framework Quarkus
            └── java-spring/    <-- Especialização para Framework Spring
```

-----

### 2\. Generic (`/architecture/blueprints/java/java-generic`)

Define a Clean Architecture (Domain, Application, Infrastructure) e padrões de código agnósticos a framework.

## 📋 Quando usar cada um?
- Para APIs e Microservices: **Quarkus**
- Para Processamento de Dados (Batch): **Spring Boot**
- Para bibliotecas compartilhadas: Seguir o padrão **Generic**

-----

### 3\. Especialização: Quarkus (`/architecture/blueprints/java/java-quarkus`)

Aqui, pegamos o DNA genérico e adicionamos as ferramentas específicas do Quarkus para alta performance e Cloud Native.

**Diferenciais do Blueprint Quarkus:**

  * **Extensions:** Inclusão de `quarkus-resteasy-reactive`, `quarkus-hibernate-panache`.
  * **Dockerfiles:** Otimizados para `native-image` (GraalVM).
  * **Config:** `application.properties` pré-configurado para ambiente de Dev/Prod.

-----

### 4\. Especialização: Spring Boot (`/architecture/blueprints/java/java-spring`)

Focado em serviços que exigem o ecossistema robusto do Spring (como Batch ou integrações complexas).

**Diferenciais do Blueprint Spring:**

  * **Starters:** Inclusão de `spring-boot-starter-web`, `spring-boot-starter-data-jpa`.
  * **Actuator:** Pré-configurado para métricas de saúde (Health Checks) padrão da empresa.
  * **Profiles:** `application-dev.yml` e `application-prod.yml` padronizados.

-----

### 💡 Vide:

- Próximo passo é verificar o **ADR (Architecture Decision Record)** dentro de `architecture/adr/` onde se explica **por que e quando. ** adotar **Quarkus** ou **Spring**.
