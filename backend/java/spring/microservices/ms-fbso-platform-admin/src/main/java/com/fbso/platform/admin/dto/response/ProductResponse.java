package com.fbso.platform.admin.dto.response;

import java.time.OffsetDateTime;
import java.util.UUID;

/**
 * Response para Produto/Serviço (F04-06).
 *
 * <p>RN18-03: {@code fiscalMappingStatus} default "NOT_MAPPED" (placeholder fiscal).</p>
 */
public record ProductResponse(
        UUID id,
        UUID tenantId,
        UUID businessUnitId,
        String name,
        String sku,
        String type,
        String description,
        String status,
        String fiscalMappingStatus,
        OffsetDateTime createdDt,
        OffsetDateTime updatedDt
) {}
