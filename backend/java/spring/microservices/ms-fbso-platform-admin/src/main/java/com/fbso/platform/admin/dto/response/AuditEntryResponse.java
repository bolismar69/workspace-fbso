package com.fbso.platform.admin.dto.response;

import com.fbso.platform.admin.entity.AuditEntry;

import java.time.OffsetDateTime;
import java.util.UUID;

public record AuditEntryResponse(
        UUID id,
        OffsetDateTime timestamp,
        UUID tenantId,
        String action,
        String entityType,
        UUID entityId,
        UUID actorId,
        String actorName,
        String previousValue,
        String newValue,
        String reason
) {
    public static AuditEntryResponse from(AuditEntry e) {
        return new AuditEntryResponse(e.getId(), e.getTimestamp(), e.getTenantId(),
                e.getAction(), e.getEntityType(), e.getEntityId(),
                e.getActorId(), e.getActorName(),
                e.getPreviousValue(), e.getNewValue(), e.getReason());
    }
}
