package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonInclude;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

/**
 * Resposta do dashboard administrativo — indicadores consolidados do SaaS.
 * <p>
 * F01-01: Métricas operacionais — contas ativas, por status, por plano, receita mensal.
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
public record DashboardSummaryResponse(
        int totalAccounts,
        int activeAccounts,
        int pendingAccounts,
        int suspendedAccounts,
        Map<String, Integer> accountsByPlan,
        BigDecimal monthlyRevenue,
        String period
) {
    public static DashboardSummaryResponse of(
            int total, int active, int pending, int suspended,
            Map<String, Integer> byPlan, BigDecimal revenue, String period) {
        return new DashboardSummaryResponse(total, active, pending, suspended,
                byPlan, revenue, period);
    }
}
