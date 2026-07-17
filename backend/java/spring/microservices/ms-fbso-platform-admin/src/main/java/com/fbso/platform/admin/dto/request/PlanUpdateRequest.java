package com.fbso.platform.admin.dto.request;

import com.fbso.platform.admin.enums.Recurrence;
import jakarta.validation.constraints.Positive;

import java.math.BigDecimal;

public record PlanUpdateRequest(
        String name,
        String description,
        @Positive BigDecimal price,
        Recurrence recurrence
) {}
