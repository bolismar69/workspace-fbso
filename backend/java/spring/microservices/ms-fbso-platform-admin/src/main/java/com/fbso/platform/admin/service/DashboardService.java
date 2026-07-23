package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.response.AccountsByPlanResponse;
import com.fbso.platform.admin.dto.response.AccountsByStatusResponse;
import com.fbso.platform.admin.dto.response.AlertResponse;
import com.fbso.platform.admin.dto.response.DashboardSummaryResponse;
import com.fbso.platform.admin.dto.response.EvolutionResponse;
import com.fbso.platform.admin.repository.DashboardRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.OffsetDateTime;
import java.time.YearMonth;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;

/**
 * Serviço de métricas do Dashboard Administrativo (F01-01, F01-03).
 * <p>
 * Calcula indicadores consolidados do SaaS a partir de queries agregadas
 * no {@link DashboardRepository}.
 * <p>
 * RN01-02: Período padrão = mês atual.
 *
 * @see DashboardRepository
 */
@Service
public class DashboardService {

    private static final Logger log = LoggerFactory.getLogger(DashboardService.class);

    private final DashboardRepository repo;

    public DashboardService(DashboardRepository repo) {
        this.repo = repo;
    }

    // ---- Summary ----

    /**
     * Indicadores consolidados do SaaS — total, ativos, pendentes, suspensos,
     * distribuição por plano e receita mensal estimada.
     */
    public DashboardSummaryResponse getSummary() {
        int total = repo.countActive();
        int active = repo.countByStatus("ACTIVE");
        int pending = repo.countByStatus("PENDING_ONBOARDING");
        int suspended = repo.countByStatus("SUSPENDED");

        Map<String, Integer> byPlan = new LinkedHashMap<>();
        for (var row : repo.accountsByPlan()) {
            byPlan.put((String) row.get("plan_name"),
                    ((Number) row.get("count")).intValue());
        }

        BigDecimal revenue = repo.monthlyRevenue();

        return DashboardSummaryResponse.of(total, active, pending, suspended,
                byPlan, revenue, "mes_atual");
    }

    // ---- Evolution ----

    /**
     * Evolução temporal da base de tenants.
     * <p>
     * Períodos suportados: 7d, 30d, 90d, mes_atual, ano_atual.
     * Se inválido → assume mês atual (RN01-02).
     */
    public EvolutionResponse getEvolution(String period) {
        String normalized = normalizePeriod(period);
        OffsetDateTime since = resolveSince(normalized);

        List<EvolutionResponse.DataPoint> dataPoints = new ArrayList<>();
        for (var row : repo.evolutionByDay(since)) {
            dataPoints.add(new EvolutionResponse.DataPoint(
                    ((java.sql.Date) row.get("date")).toLocalDate(),
                    ((Number) row.get("count")).intValue()));
        }

        return EvolutionResponse.of(normalized, dataPoints);
    }

    // ---- Accounts by Status ----

    public AccountsByStatusResponse getAccountsByStatus() {
        List<AccountsByStatusResponse.StatusCount> statuses = new ArrayList<>();
        for (var row : repo.accountsByStatus()) {
            statuses.add(new AccountsByStatusResponse.StatusCount(
                    (String) row.get("status"),
                    ((Number) row.get("count")).intValue()));
        }
        return AccountsByStatusResponse.of(statuses);
    }

    // ---- Accounts by Plan ----

    public AccountsByPlanResponse getAccountsByPlan() {
        List<AccountsByPlanResponse.PlanCount> plans = new ArrayList<>();
        for (var row : repo.accountsByPlan()) {
            plans.add(new AccountsByPlanResponse.PlanCount(
                    (String) row.get("plan_name"),
                    ((Number) row.get("count")).intValue()));
        }
        return AccountsByPlanResponse.of(plans);
    }

    // ---- Alertas (F01-03) ----

    /**
     * Alertas do dashboard: onboarding >48h (WARNING) + assinatura suspensa (CRITICAL).
     */
    public AlertResponse getAlerts() {
        List<AlertResponse.Alert> alerts = new ArrayList<>();

        // RN03-01: Onboarding >48h → WARNING
        for (var row : repo.onboardingStalled()) {
            alerts.add(new AlertResponse.Alert(
                    AlertResponse.AlertType.WARNING,
                    "Onboarding pendente há mais de 48h: "
                            + row.get("name_corporate"),
                    (UUID) row.get("id"),
                    "TENANT"));
        }

        // RN03-02: Assinatura suspensa → CRITICAL
        for (var row : repo.suspendedSubscriptions()) {
            alerts.add(new AlertResponse.Alert(
                    AlertResponse.AlertType.CRITICAL,
                    "Assinatura suspensa: " + row.get("name_corporate")
                            + " — Plano: " + row.get("plan_name"),
                    (UUID) row.get("tenant_id"),
                    "SUBSCRIPTION"));
        }

        log.debug("Alertas carregados: {} warnings, {} críticos",
                alerts.stream().filter(a -> a.type() == AlertResponse.AlertType.WARNING).count(),
                alerts.stream().filter(a -> a.type() == AlertResponse.AlertType.CRITICAL).count());

        return AlertResponse.of(alerts);
    }

    // ---- Helpers ----

    /**
     * Normaliza o período para um valor conhecido.
     * Se inválido → mês atual (RN01-02).
     */
    private String normalizePeriod(String period) {
        if (period == null || period.isBlank()) return "mes_atual";
        return switch (period.toLowerCase()) {
            case "7d", "30d", "90d", "mes_atual", "ano_atual" -> period.toLowerCase();
            default -> "mes_atual";
        };
    }

    /**
     * Calcula a data de início com base no período.
     */
    private OffsetDateTime resolveSince(String period) {
        OffsetDateTime now = OffsetDateTime.now(java.time.ZoneOffset.UTC);
        return switch (period) {
            case "7d" -> now.minusDays(7);
            case "30d" -> now.minusDays(30);
            case "90d" -> now.minusDays(90);
            case "ano_atual" -> now.withDayOfYear(1).withHour(0).withMinute(0).withSecond(0).withNano(0);
            case "mes_atual" -> YearMonth.now().atDay(1).atStartOfDay()
                    .atOffset(now.getOffset());
            default -> YearMonth.now().atDay(1).atStartOfDay()
                    .atOffset(now.getOffset());
        };
    }
}
