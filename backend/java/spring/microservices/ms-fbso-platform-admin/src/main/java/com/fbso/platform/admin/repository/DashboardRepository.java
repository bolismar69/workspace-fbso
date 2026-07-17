package com.fbso.platform.admin.repository;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

/**
 * Repository de queries agregadas para o Dashboard Administrativo (F01-01).
 * <p>
 * Todas as queries operam sobre a tabela {@code fbso_platform.tenant} e
 * respeitam soft delete ({@code deleted_dt IS NULL}).
 * <p>
 * A tabela tenant é GLOBAL — não possui {@code tenant_id}, portanto
 * não herda de {@link com.fbso.platform.admin.repository.common.BaseRepository}.
 *
 * @see com.fbso.platform.admin.service.DashboardService
 */
@Repository
public class DashboardRepository {

    private final JdbcTemplate jdbc;

    public DashboardRepository(JdbcTemplate jdbc) {
        this.jdbc = jdbc;
    }

    // ---- Counts ----

    /**
     * @return total de tenants ativos (não soft-deletados)
     */
    public int countActive() {
        String sql = "SELECT COUNT(*) FROM fbso_platform.tenant WHERE deleted_dt IS NULL";
        Integer result = jdbc.queryForObject(sql, Integer.class);
        return result != null ? result : 0;
    }

    /**
     * @return total de tenants com o status informado
     */
    public int countByStatus(String status) {
        String sql = "SELECT COUNT(*) FROM fbso_platform.tenant"
                   + " WHERE deleted_dt IS NULL AND status = ?";
        Integer result = jdbc.queryForObject(sql, Integer.class, status);
        return result != null ? result : 0;
    }

    // ---- Agregações ----

    /**
     * Agrupa tenants por status (ACTIVE, PENDING_ONBOARDING, SUSPENDED, INACTIVE).
     */
    public List<Map<String, Object>> accountsByStatus() {
        String sql = """
            SELECT status, COUNT(*) AS count
            FROM fbso_platform.tenant
            WHERE deleted_dt IS NULL
            GROUP BY status
            ORDER BY count DESC
            """;
        return jdbc.queryForList(sql);
    }

    /**
     * Agrupa tenants por plano contratado (via subscription ativa + plan).
     */
    public List<Map<String, Object>> accountsByPlan() {
        String sql = """
            SELECT p.name AS plan_name, COUNT(DISTINCT s.tenant_id) AS count
            FROM fbso_platform.subscription s
            JOIN fbso_platform.plan p ON s.plan_id = p.id
            WHERE s.status = 'ACTIVE'
              AND s.deleted_dt IS NULL
              AND p.deleted_dt IS NULL
            GROUP BY p.name
            ORDER BY count DESC
            """;
        return jdbc.queryForList(sql);
    }

    /**
     * Receita mensal estimada: soma de price das assinaturas ativas
     * (usa locked_price quando disponível — DT-009, caso contrário price do plano).
     */
    public BigDecimal monthlyRevenue() {
        String sql = """
            SELECT COALESCE(SUM(
                CASE WHEN s.locked_price IS NOT NULL THEN s.locked_price
                     ELSE p.price
                END
            ), 0)
            FROM fbso_platform.subscription s
            JOIN fbso_platform.plan p ON s.plan_id = p.id
            WHERE s.status = 'ACTIVE'
              AND s.deleted_dt IS NULL
            """;
        return jdbc.queryForObject(sql, BigDecimal.class);
    }

    // ---- Evolução Temporal ----

    /**
     * Contagem de tenants criados desde {@code since}, agrupados por dia.
     */
    public List<Map<String, Object>> evolutionByDay(OffsetDateTime since) {
        String sql = """
            SELECT DATE(created_dt) AS date, COUNT(*) AS count
            FROM fbso_platform.tenant
            WHERE deleted_dt IS NULL
              AND created_dt >= ?
            GROUP BY DATE(created_dt)
            ORDER BY date
            """;
        return jdbc.queryForList(sql, since);
    }

    // ---- Alertas (F01-03) ----

    /**
     * Tenants em onboarding há mais de 48h (RN03-01).
     */
    public List<Map<String, Object>> onboardingStalled() {
        String sql = """
            SELECT id, name_corporate, status, created_dt
            FROM fbso_platform.tenant
            WHERE deleted_dt IS NULL
              AND status = 'PENDING_ONBOARDING'
              AND created_dt < NOW() - INTERVAL '48 hours'
            ORDER BY created_dt ASC
            """;
        return jdbc.queryForList(sql);
    }

    /**
     * Assinaturas suspensas com nome do tenant e plano (RN03-02).
     */
    public List<Map<String, Object>> suspendedSubscriptions() {
        String sql = """
            SELECT s.id AS subscription_id, s.tenant_id,
                   t.name_corporate, p.name AS plan_name, s.status
            FROM fbso_platform.subscription s
            JOIN fbso_platform.tenant t ON s.tenant_id = t.id
            JOIN fbso_platform.plan p ON s.plan_id = p.id
            WHERE s.status = 'SUSPENDED'
              AND s.deleted_dt IS NULL
              AND t.deleted_dt IS NULL
            ORDER BY t.name_corporate
            """;
        return jdbc.queryForList(sql);
    }
}
