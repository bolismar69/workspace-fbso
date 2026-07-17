package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonInclude;

import java.util.List;

/**
 * Distribuição de contas por status.
 * <p>
 * F01-01: Agrupamento de tenants por status (ACTIVE, PENDING_ONBOARDING, SUSPENDED, INACTIVE).
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
public record AccountsByStatusResponse(
        List<StatusCount> statuses
) {
    public record StatusCount(String status, int count) {}

    public static AccountsByStatusResponse of(List<StatusCount> statuses) {
        return new AccountsByStatusResponse(statuses);
    }
}
