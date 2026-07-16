Abaixo, apresento os diagramas em formato **Mermaid** (que é renderizado nativamente em arquivos Markdown), estruturando a solução com os dois microserviços, o banco de dados PostgreSQL e o cache Redis.

---

# 🏗️ Documentação de Arquitetura C4 - Catálogo Simples

## 1. Diagrama de Contexto (Nível 1)

O objetivo aqui é mostrar como o sistema de catálogo interage com os usuários e o mundo externo.

```mermaid
graph TD
    User([Cliente / Gerente]) -- Consulta produtos e preços --> System[Sistema de Catálogo de Produtos]
    System -- Armazena dados de inventário --> DB[(PostgreSQL)]
    
    style System fill:#1168bd,color:#fff

```

---

## 2. Diagrama de Contêineres (Nível 2)

Aqui detalhamos a separação entre o Admin e o Engine de Consulta, além da infraestrutura de dados.

```mermaid
graph TB
    Client([Usuários/Apps]) -- "REST (Escrita)" --> Admin[ms-product-catalog-admin-simple]
    Client -- "REST (Leitura)" --> Engine[ms-product-catalog-engine-simple]

    subgraph "Infraestrutura de Dados"
        DB[(PostgreSQL)]
        Cache[(Redis Cache)]
    end

    Admin -- "Salva/Altera" --> DB
    Engine -- "Lê dados (Fallback)" --> DB
    Engine -- "Check/Set Cache (TTL 5m)" --> Cache

    style Admin fill:#1168bd,color:#fff
    style Engine fill:#1168bd,color:#fff
    style DB fill:#2b2b2b,color:#fff
    style Cache fill:#a41e11,color:#fff

```

---

## 3. Diagrama de Componentes (Nível 3)

Focaremos no **ms-product-catalog-engine-simple**, detalhando como o Spring Boot organiza a lógica de cache.

```mermaid
graph TB
    subgraph "Engine Microservice (Java 21 / Spring Boot)"
        Controller[Catalog Controller] --> Service[Catalog Service]
        Service --> Repository[Product Repository]
        Service --> CacheManager[Spring Cache Manager]
    end

    Repository -- "JDBC/SQL" --> DB[(PostgreSQL)]
    CacheManager -- "Lettuce Driver" --> Redis[(Redis)]

    style Service fill:#1168bd,color:#fff

```

---

## 4. Diagrama de Código (Nível 4)

Representação simplificada das classes principais e da lógica de decisão do Cache (Padrão *Cache-Aside*).

```mermaid
classDiagram
    class Product {
        +Long id
        +String sku
        +String name
        +BigDecimal price
    }
    class CatalogService {
        +findBySku(String sku) ProductResponse
    }
    class ProductRepository {
        <<interface>>
        +findBySku(String sku)
    }

    CatalogService ..> ProductRepository : usa
    CatalogService ..> Redis : 1. Verifica Cache
    CatalogService ..> ProductRepository : 2. Se vazio, busca DB

```

---

## 📝 Resumo da Decisão Arquitetural

* **Separação de Preocupações:** O serviço de **Admin** detém a soberania da escrita. O serviço de **Engine** detém a responsabilidade pela performance e entrega de leitura.
* **Estratégia de Cache:** Baseada em **TTL (Time To Live)** de 300 segundos (5 minutos). Não há acoplamento de mensagens entre os serviços nesta fase.
* **Persistência:** PostgreSQL utilizando o schema `product_catalog_simple` com as triggers de SKU e Auditoria que desenvolvemos.
* **Eficiência:** Ambos os serviços desenhados para compilação nativa via **GraalVM**, garantindo baixo uso de memória e inicialização rápida.

---
