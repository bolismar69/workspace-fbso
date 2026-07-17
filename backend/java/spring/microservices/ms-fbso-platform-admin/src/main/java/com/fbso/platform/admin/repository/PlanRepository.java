package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.Plan;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.PlanRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

/**
 * Repository para a entidade {@link Plan} — tabela {@code fbso_platform.plan}.
 *
 * <p>Tabela global (sem tenant_id) — visível por todos os tenants.</p>
 */
@Repository
public class PlanRepository extends BaseRepository<Plan> {

    private static final PlanRowMapper ROW_MAPPER = new PlanRowMapper();

    public PlanRepository(JdbcTemplate jdbc) {
        super(jdbc, "plan", ROW_MAPPER, false); // global table
    }

    /**
     * Lista todos os planos ativos (para exibição pública e assinatura).
     */
    public List<Plan> findAllActive() {
        String sql = "SELECT * FROM fbso_platform.plan"
                   + " WHERE deleted_dt IS NULL AND status = 'ACTIVE'"
                   + " ORDER BY price ASC";
        return jdbc.query(sql, ROW_MAPPER);
    }

    /**
     * Verifica se existem assinaturas ativas vinculadas a este plano (RN06-01).
     */
    public boolean hasActiveSubscribers(UUID planId) {
        String sql = """
            SELECT COUNT(*) FROM fbso_platform.subscription
            WHERE plan_id = ? AND status = 'ACTIVE' AND deleted_dt IS NULL
            """;
        Integer count = jdbc.queryForObject(sql, Integer.class, planId);
        return count != null && count > 0;
    }

    /**
     * Conta quantos planos ativos existem (RN06-03 — mínimo 1 ativo).
     */
    public int countActive() {
        String sql = "SELECT COUNT(*) FROM fbso_platform.plan"
                   + " WHERE deleted_dt IS NULL AND status = 'ACTIVE'";
        Integer count = jdbc.queryForObject(sql, Integer.class);
        return count != null ? count : 0;
    }
}
