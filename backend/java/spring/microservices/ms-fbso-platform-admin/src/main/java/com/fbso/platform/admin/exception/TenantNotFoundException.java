package com.fbso.platform.admin.exception;

import java.util.UUID;

/**
 * Exceção lançada quando um Tenant não é encontrado.
 * <p>
 * Mapeada para HTTP 404 (Not Found) pelo {@link GlobalExceptionHandler}.
 */
public class TenantNotFoundException extends BusinessException {

    public TenantNotFoundException(UUID tenantId) {
        super("tenant-not-found", "Tenant não encontrado: " + tenantId);
    }

    public TenantNotFoundException(String identifier) {
        super("tenant-not-found", "Tenant não encontrado: " + identifier);
    }
}
