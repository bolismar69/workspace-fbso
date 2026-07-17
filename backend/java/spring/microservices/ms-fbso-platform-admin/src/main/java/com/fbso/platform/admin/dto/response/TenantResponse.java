package com.fbso.platform.admin.dto.response;

import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantSegment;
import com.fbso.platform.admin.enums.TenantStatus;

import java.time.OffsetDateTime;
import java.util.UUID;

/**
 * DTO de resposta para Tenant (F02-01, F02-02).
 */
public record TenantResponse(
        UUID id,
        String nameCorporate,
        String nameFantasy,
        TenantSegment segment,
        TenantStatus status,
        OffsetDateTime createdDt
) {
    public static TenantResponse from(Tenant t) {
        return new TenantResponse(
                t.getId(), t.getNameCorporate(), t.getNameFantasy(),
                t.getSegment(), t.getStatus(), t.getCreatedDt());
    }
}
