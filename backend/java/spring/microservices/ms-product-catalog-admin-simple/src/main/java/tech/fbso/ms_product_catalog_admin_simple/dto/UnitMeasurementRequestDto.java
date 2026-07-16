package tech.fbso.ms_product_catalog_admin_simple.dto;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;

public record UnitMeasurementRequestDto(
    @NotBlank(message = "Acronym is required")
    @Size(min = 1, max = 10, message = "Acronym must be between 1 and 10 characters")
    String acronym,
    
    @NotBlank(message = "Description is required")
    @Size(min = 1, max = 255, message = "Description must be between 1 and 255 characters")
    String description
) {}
