package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.NotNull;
import java.util.UUID;

public record ChangePlanRequest(
        @NotNull UUID newPlanId
) {}
