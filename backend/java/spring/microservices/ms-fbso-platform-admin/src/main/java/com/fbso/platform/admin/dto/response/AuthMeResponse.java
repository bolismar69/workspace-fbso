package com.fbso.platform.admin.dto.response;

import java.util.List;
import java.util.UUID;

public record AuthMeResponse(
    UUID id, String name, String email, String role,
    List<UUID> businessUnitIds, List<String> modules,
    UUID tenantId, String onboardingStatus
) {}
