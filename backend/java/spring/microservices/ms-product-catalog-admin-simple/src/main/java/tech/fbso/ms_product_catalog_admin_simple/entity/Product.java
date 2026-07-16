package tech.fbso.ms_product_catalog_admin_simple.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import java.math.BigDecimal;
import java.time.OffsetDateTime;

@Entity
@Table(name = "product", schema = "product_catalog_simple")
@Getter
@Setter
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