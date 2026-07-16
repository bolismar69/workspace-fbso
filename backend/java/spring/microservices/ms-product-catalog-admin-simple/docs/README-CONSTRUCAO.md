Com o DER e a arquitetura definidos, vamos estruturar o **ms-product-catalog-admin-simple**. O objetivo deste serviço é garantir a integridade dos dados ao inserir ou atualizar produtos no PostgreSQL.

Como usaremos **Java 21**, **Spring Boot 3** e foco em **GraalVM**, utilizaremos as melhores práticas de imutabilidade e performance.

---

### 1. Estrutura de Pastas Sugerida

```text
com.supermarket.admin
├── config              # Configuração do Jackson (JSON) e DB
├── controller          # Endpoints REST
├── dto                 # Records para Request e Response
├── exception           # Tratamento de erros (Custom Handler)
├── model               # Entidades JPA
├── repository          # Interfaces Spring Data JPA
└── service             # Lógica de negócio

```

---

### 2. Entidades JPA (Model)

Aqui está a entidade principal `Product`. Note o uso do `schema` no mapeamento e as colunas de auditoria que o banco preencherá.

```java
package com.supermarket.admin.model;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import java.math.BigDecimal;
import java.time.OffsetDateTime;

@Entity
@Table(name = "product", schema = "product_catalog_simple")
@Getter @Setter
public class Product {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    // Insertable/Updatable false pois a Trigger do banco gera o SKU
    @Column(insertable = false, updatable = false)
    private String sku;

    private String barcode;
    private String name;
    private BigDecimal price;

    @Column(name = "quantity_in_stock")
    private BigDecimal quantityInStock;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "category_id")
    private Category category;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "unit_measurement_id")
    private UnitMeasurement unitMeasurement;

    @Column(name = "created_at", insertable = false, updatable = false)
    private OffsetDateTime createdAt;

    @Column(name = "updated_at", insertable = false, updatable = false)
    private OffsetDateTime updatedAt;
}

```

---

### 3. Data Transfer Objects (Records)

Usar `records` é ideal para a compilação nativa e garante imutabilidade.

```java
package com.supermarket.admin.dto;

import java.math.BigDecimal;

public record ProductRequest(
    String name,
    String barcode,
    BigDecimal price,
    BigDecimal quantityInStock,
    Integer categoryId,
    Integer unitMeasurementId
) {}

public record ProductResponse(
    Long id,
    String sku,
    String name,
    BigDecimal price,
    BigDecimal quantityInStock
) {}

```

---

### 4. Service (Lógica de Escrita)

O serviço é simples, pois as regras de SKU e Timestamps estão delegadas ao PostgreSQL via Triggers.

```java
package com.supermarket.admin.service;

import com.supermarket.admin.dto.ProductRequest;
import com.supermarket.admin.model.Product;
import com.supermarket.admin.repository.ProductRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class ProductAdminService {

    private final ProductRepository repository;

    public ProductAdminService(ProductRepository repository) {
        this.repository = repository;
    }

    @Transactional
    public Product create(ProductRequest request) {
        Product product = new Product();
        product.setName(request.name());
        product.setBarcode(request.barcode());
        product.setPrice(request.price());
        product.setQuantityInStock(request.quantityInStock());
        
        // Aqui buscaríamos as referências de Categoria e Unidade
        // product.setCategory(categoryRepo.getReferenceById(request.categoryId()));
        
        return repository.save(product);
    }
}

```

---

### 5. Configuração para Imagem Nativa (GraalVM)

No seu arquivo `pom.xml` (se usar Maven), você deve incluir o plugin do Spring Boot configurado para o perfil nativo:

```xml
<build>
    <plugins>
        <plugin>
            <groupId>org.graalvm.buildtools</groupId>
            <artifactId>native-maven-plugin</artifactId>
        </plugin>
        <plugin>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-maven-plugin</artifactId>
        </plugin>
    </plugins>
</build>

```

### O que acontece agora?

Para validar este microserviço, precisamos garantir que o banco de dados esteja rodando com o Schema que criamos.

**Sugiro que o próximo passo seja o `docker-compose.yml` para subirmos o ambiente completo (Postgres + Redis + Admin App) para testes de integração.** Aceita esta abordagem?