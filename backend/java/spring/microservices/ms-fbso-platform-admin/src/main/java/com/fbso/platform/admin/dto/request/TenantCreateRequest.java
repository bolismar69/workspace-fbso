package com.fbso.platform.admin.dto.request;

import com.fbso.platform.admin.enums.TenantSegment;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

/**
 * DTO para criação de Tenant (F02-01).
 */
public record TenantCreateRequest(
        @NotBlank(message = "Razão social é obrigatória")
        String nameCorporate,

        String nameFantasy,

        @NotNull(message = "Segmento é obrigatório")
        TenantSegment segment
) {}
