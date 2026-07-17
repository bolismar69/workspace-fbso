package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonInclude;

import java.util.List;

/**
 * Distribuição de contas por plano comercial.
 * <p>
 * F01-01: Agrupamento de tenants por plano contratado (via subscription ativa).
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
public record AccountsByPlanResponse(
        List<PlanCount> plans
) {
    public record PlanCount(String planName, int count) {}

    public static AccountsByPlanResponse of(List<PlanCount> plans) {
        return new AccountsByPlanResponse(plans);
    }
}
