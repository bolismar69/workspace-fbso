package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.response.DashboardClientResponse;
import com.fbso.platform.admin.dto.response.NotificationResponse;
import com.fbso.platform.admin.security.TenantContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

/**
 * Serviço de dashboard do cliente — agrega dados do tenant autenticado.
 * <p>Fase 0: dashboard genérico (RN15-02). Placeholder "FBSO Platform" (RN16-02).</p>
 */
@Service
public class DashboardClientService {

    private static final Logger log = LoggerFactory.getLogger(DashboardClientService.class);
    private final JdbcTemplate jdbc;

    public DashboardClientService(JdbcTemplate jdbc) {
        this.jdbc = jdbc;
    }

    public DashboardClientResponse getSummary(UUID tenantId) {
        int activeUnits = countActiveUnits(tenantId);
        int productCount = countProducts(tenantId);

        String planName = "FBSO Platform";
        String planStatus = "ACTIVE";
        try {
            var rows = jdbc.queryForList(
                "SELECT p.name, s.status FROM fbso_platform.subscription s " +
                "JOIN fbso_platform.plan p ON s.plan_id = p.id " +
                "WHERE s.tenant_id = ? AND s.status = 'ACTIVE' AND s.deleted_dt IS NULL " +
                "LIMIT 1", tenantId);
            if (!rows.isEmpty()) {
                planName = (String) rows.get(0).get("name");
                planStatus = (String) rows.get(0).get("status");
            }
        } catch (Exception e) {
            log.debug("Dashboard: tenant sem assinatura ativa — usando placeholder. tenantId={}", tenantId);
        }

        List<NotificationResponse> notifications = List.of();

        return new DashboardClientResponse(
            activeUnits, "/business-units",
            productCount, "/products",
            planName, planStatus, "/plan",
            notifications, "/notifications"
        );
    }

    private int countActiveUnits(UUID tenantId) {
        Integer count = jdbc.queryForObject(
            "SELECT COUNT(*) FROM fbso_platform.business_unit WHERE tenant_id = ? AND deleted_dt IS NULL AND status = 'ACTIVE'",
            Integer.class, tenantId);
        return count != null ? count : 0;
    }

    private int countProducts(UUID tenantId) {
        Integer count = jdbc.queryForObject(
            "SELECT COUNT(*) FROM fbso_platform.product_service ps " +
            "JOIN fbso_platform.business_unit bu ON ps.business_unit_id = bu.id " +
            "WHERE bu.tenant_id = ? AND ps.deleted_dt IS NULL AND ps.status = 'ACTIVE'",
            Integer.class, tenantId);
        return count != null ? count : 0;
    }
}
