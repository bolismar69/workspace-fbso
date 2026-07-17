package com.fbso.platform.admin.dto.response;

import com.fbso.platform.admin.entity.Plan;
import com.fbso.platform.admin.enums.Recurrence;

import java.math.BigDecimal;
import java.util.UUID;

public record PlanResponse(
        UUID id,
        String name,
        String description,
        BigDecimal price,
        Recurrence recurrence,
        String status,
        int version
) {
    public static PlanResponse from(Plan p) {
        return new PlanResponse(p.getId(), p.getName(), p.getDescription(),
                p.getPrice(), p.getRecurrence(), p.getStatus(), p.getVersion());
    }
}
