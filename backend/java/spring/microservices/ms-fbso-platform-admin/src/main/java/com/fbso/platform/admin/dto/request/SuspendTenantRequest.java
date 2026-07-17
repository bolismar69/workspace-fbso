package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.NotBlank;

/**
 * DTO para suspensão de Tenant (F02-02, RN05-02).
 * Exige motivo da suspensão.
 */
public record SuspendTenantRequest(
        @NotBlank(message = "Motivo da suspensão é obrigatório")
        String reason
) {}
