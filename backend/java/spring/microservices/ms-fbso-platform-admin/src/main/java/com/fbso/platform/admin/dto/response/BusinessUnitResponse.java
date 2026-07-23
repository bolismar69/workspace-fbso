package com.fbso.platform.admin.dto.response;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.UUID;

/**
 * Response para Unidade de Negócio (F04-05).
 *
 * <p>Inclui sub-árvore hierárquica (children) para navegação.</p>
 */
public record BusinessUnitResponse(
        UUID id,
        UUID tenantId,
        UUID parentId,
        String cnpj,
        String corporateName,
        String taxRegime,
        String street,
        String number,
        String complement,
        String neighborhood,
        String city,
        String state,
        String zipCode,
        String status,
        boolean isMatrix,
        List<BusinessUnitResponse> children,
        OffsetDateTime createdDt,
        OffsetDateTime updatedDt
) {}
