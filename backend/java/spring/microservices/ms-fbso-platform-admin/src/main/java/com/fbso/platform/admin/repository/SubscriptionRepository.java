package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.Subscription;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.SubscriptionRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Repository para {@link Subscription} — tabela com RLS (hasTenantColumn=true).
 */
@Repository
public class SubscriptionRepository extends BaseRepository<Subscription> {

    private static final SubscriptionRowMapper ROW_MAPPER = new SubscriptionRowMapper();

    public SubscriptionRepository(JdbcTemplate jdbc) {
        super(jdbc, "subscription", ROW_MAPPER, true);
    }

    /**
     * Busca a assinatura ativa de um tenant (RN07-01).
     */
    public Optional<Subscription> findActiveByTenantId(UUID tenantId) {
        String sql = """
            SELECT * FROM fbso_platform.subscription
            WHERE tenant_id = ? AND status = 'ACTIVE' AND deleted_dt IS NULL
            """;
        List<Subscription> results = jdbc.query(sql, ROW_MAPPER, tenantId);
        return results.isEmpty() ? Optional.empty() : Optional.of(results.get(0));
    }

    /**
     * Histórico de assinaturas de um tenant.
     */
    public List<Subscription> findByTenantId(UUID tenantId) {
        String sql = """
            SELECT * FROM fbso_platform.subscription
            WHERE tenant_id = ? AND deleted_dt IS NULL
            ORDER BY start_date DESC
            """;
        return jdbc.query(sql, ROW_MAPPER, tenantId);
    }
}
