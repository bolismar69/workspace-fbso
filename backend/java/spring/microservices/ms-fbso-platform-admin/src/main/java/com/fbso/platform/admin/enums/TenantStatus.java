package com.fbso.platform.admin.enums;

/**
 * Status do ciclo de vida de um Tenant.
 * <p>
 * Transições permitidas (RN05-01):
 * <pre>
 * PENDING_ONBOARDING → ACTIVE
 * ACTIVE ↔ SUSPENDED
 * ACTIVE ↔ INACTIVE
 * </pre>
 */
public enum TenantStatus {
    PENDING_ONBOARDING,
    ACTIVE,
    SUSPENDED,
    INACTIVE
}
