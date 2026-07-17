package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.response.AccountsByPlanResponse;
import com.fbso.platform.admin.dto.response.AccountsByStatusResponse;
import com.fbso.platform.admin.dto.response.AlertResponse;
import com.fbso.platform.admin.dto.response.DashboardSummaryResponse;
import com.fbso.platform.admin.dto.response.EvolutionResponse;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.DashboardService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

/**
 * Endpoints do Dashboard Administrativo (F01-01, F01-03).
 * <p>
 * Acessível exclusivamente pelo Admin FBSO (papel ADMIN_TENANT).
 * Todos os endpoints exigem permissão {@code DASHBOARD:view}.
 * <p>
 * Performance: p95 ≤ 3s com 1000 tenants (BR-NFR05).
 *
 * @see DashboardService
 */
@RestController
@RequestMapping("/api/v1/dashboard/admin")
@Tag(name = "Dashboard Admin", description = "Métricas operacionais do SaaS — Admin FBSO")
public class DashboardController {

    private static final Logger log = LoggerFactory.getLogger(DashboardController.class);

    private final DashboardService service;

    public DashboardController(DashboardService service) {
        this.service = service;
    }

    /**
     * Indicadores consolidados: total de contas, ativas, pendentes, suspensas,
     * distribuição por plano e receita mensal estimada.
     */
    @GetMapping("/summary")
    @Operation(summary = "Indicadores consolidados do SaaS")
    @RequiresPermission(resource = "DASHBOARD", action = "view")
    public ResponseEntity<DashboardSummaryResponse> getSummary() {
        log.debug("Dashboard: carregando summary");
        DashboardSummaryResponse response = service.getSummary();
        return ResponseEntity.ok(response);
    }

    /**
     * Evolução temporal da base de tenants.
     *
     * @param period período opcional (7d, 30d, 90d, mes_atual, ano_atual).
     *               Padrão: mês atual (RN01-02)
     */
    @GetMapping("/evolution")
    @Operation(summary = "Evolução temporal da base de tenants")
    @RequiresPermission(resource = "DASHBOARD", action = "view")
    public ResponseEntity<EvolutionResponse> getEvolution(
            @Parameter(description = "Período: 7d, 30d, 90d, mes_atual, ano_atual")
            @RequestParam(defaultValue = "mes_atual") String period) {
        log.debug("Dashboard: carregando evolution — period={}", period);
        EvolutionResponse response = service.getEvolution(period);
        return ResponseEntity.ok(response);
    }

    /**
     * Distribuição de contas por status (ACTIVE, PENDING_ONBOARDING, SUSPENDED, INACTIVE).
     */
    @GetMapping("/accounts-by-status")
    @Operation(summary = "Distribuição de contas por status")
    @RequiresPermission(resource = "DASHBOARD", action = "view")
    public ResponseEntity<AccountsByStatusResponse> getAccountsByStatus() {
        log.debug("Dashboard: carregando accounts-by-status");
        AccountsByStatusResponse response = service.getAccountsByStatus();
        return ResponseEntity.ok(response);
    }

    /**
     * Distribuição de contas por plano comercial.
     */
    @GetMapping("/accounts-by-plan")
    @Operation(summary = "Distribuição de contas por plano")
    @RequiresPermission(resource = "DASHBOARD", action = "view")
    public ResponseEntity<AccountsByPlanResponse> getAccountsByPlan() {
        log.debug("Dashboard: carregando accounts-by-plan");
        AccountsByPlanResponse response = service.getAccountsByPlan();
        return ResponseEntity.ok(response);
    }

    /**
     * Alertas do dashboard: onboarding >48h (WARNING) e assinaturas suspensas (CRITICAL).
     */
    @GetMapping("/alerts")
    @Operation(summary = "Alertas do dashboard administrativo")
    @RequiresPermission(resource = "DASHBOARD", action = "view")
    public ResponseEntity<AlertResponse> getAlerts() {
        log.debug("Dashboard: carregando alerts");
        AlertResponse response = service.getAlerts();
        return ResponseEntity.ok(response);
    }
}
