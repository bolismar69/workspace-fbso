package tech.fbso.ms_product_catalog_admin_simple.dto;

import java.math.BigDecimal;
import java.time.OffsetDateTime;

public record ProductResponseDto(
    Long id,
    String sku,
    String barcode,
    String name,
    BigDecimal price,
    BigDecimal quantityInStock,
    Long categoryId,
    String categoryName,
    Long unitMeasurementId,
    String unitMeasurementAcronym,
    OffsetDateTime createdAt,
    OffsetDateTime updatedAt
) {}
