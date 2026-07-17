package com.fbso.platform.admin.dto.request;

import com.fbso.platform.admin.enums.TenantSegment;

/**
 * DTO para edição de Tenant (F02-01).
 * Apenas campos editáveis após criação.
 */
public record TenantUpdateRequest(
        String nameFantasy,
        TenantSegment segment
) {}
