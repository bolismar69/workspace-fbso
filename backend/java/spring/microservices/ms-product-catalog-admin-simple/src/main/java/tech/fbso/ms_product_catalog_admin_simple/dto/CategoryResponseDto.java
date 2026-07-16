package tech.fbso.ms_product_catalog_admin_simple.dto;

import java.time.OffsetDateTime;

public record CategoryResponseDto(
    Long id,
    String acronym,
    String description,
    OffsetDateTime createdAt,
    OffsetDateTime updatedAt
) {}
