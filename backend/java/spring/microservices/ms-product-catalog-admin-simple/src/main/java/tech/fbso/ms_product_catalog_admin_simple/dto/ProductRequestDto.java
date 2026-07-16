package tech.fbso.ms_product_catalog_admin_simple.dto;

import jakarta.validation.constraints.*;
import java.math.BigDecimal;

public record ProductRequestDto(
    @NotBlank(message = "Product name is required")
    @Size(min = 1, max = 255, message = "Product name must be between 1 and 255 characters")
    String name,
    
    @NotBlank(message = "Barcode is required")
    @Size(min = 1, max = 50, message = "Barcode must be between 1 and 50 characters")
    String barcode,
    
    @NotNull(message = "Price is required")
    @DecimalMin(value = "0.0", inclusive = false, message = "Price must be greater than 0")
    BigDecimal price,
    
    @NotNull(message = "Quantity in stock is required")
    @DecimalMin(value = "0.0", message = "Quantity in stock cannot be negative")
    BigDecimal quantityInStock,
    
    @NotNull(message = "Category ID is required")
    @Positive(message = "Category ID must be positive")
    Long categoryId,
    
    @NotNull(message = "Unit Measurement ID is required")
    @Positive(message = "Unit Measurement ID must be positive")
    Long unitMeasurementId
) {}
