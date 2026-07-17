package com.fbso.platform.admin.dto.response;

import com.fbso.platform.admin.entity.Subscription;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.UUID;

public record SubscriptionResponse(
        UUID id,
        UUID tenantId,
        UUID planId,
        OffsetDateTime startDate,
        OffsetDateTime endDate,
        String status,
        BigDecimal lockedPrice,
        String lockedRecurrence
) {
    public static SubscriptionResponse from(Subscription s) {
        return new SubscriptionResponse(s.getId(), s.getTenantId(), s.getPlanId(),
                s.getStartDate(), s.getEndDate(), s.getStatus(),
                s.getLockedPrice(), s.getLockedRecurrence());
    }
}
