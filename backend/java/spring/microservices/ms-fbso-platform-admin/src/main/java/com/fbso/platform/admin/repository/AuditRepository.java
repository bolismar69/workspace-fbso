package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.AuditEntry;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.AuditEntryRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

/**
 * Repository somente-leitura para {@link AuditEntry} (F02-05).
 *
 * <p>Registros de auditoria são IMUTÁVEIS — este repository não expõe
 * save/update/delete. Apenas queries de consulta com filtros.</p>
 */
@Repository
public class AuditRepository extends BaseRepository<AuditEntry> {

    private static final AuditEntryRowMapper ROW_MAPPER = new AuditEntryRowMapper();

    public AuditRepository(JdbcTemplate jdbc) {
        super(jdbc, "audit_log", ROW_MAPPER, true);
    }

    /**
     * Consulta registros de auditoria com filtros opcionais.
     *
     * @param startDate  data inicial (opcional)
     * @param endDate    data final (opcional)
     * @param action     ação: CREATED, UPDATED, SUSPENDED, REACTIVATED, PLAN_CHANGED (opcional)
     * @param entityType tipo de entidade: TENANT, PLAN, SUBSCRIPTION (opcional)
     * @param page       página (0-based)
     * @param size       registros por página (padrão 25, max 100)
     * @return lista paginada ordenada por timestamp DESC
     */
    public List<AuditEntry> findByFilters(OffsetDateTime startDate, OffsetDateTime endDate,
                                           String action, String entityType,
                                           int page, int size) {
        StringBuilder sql = new StringBuilder(
                "SELECT * FROM fbso_platform.audit_log WHERE deleted_dt IS NULL");
        List<Object> params = new ArrayList<>();

        if (startDate != null) {
            sql.append(" AND timestamp >= ?");
            params.add(startDate);
        }
        if (endDate != null) {
            sql.append(" AND timestamp <= ?");
            params.add(endDate);
        }
        if (action != null && !action.isBlank()) {
            sql.append(" AND action = ?");
            params.add(action);
        }
        if (entityType != null && !entityType.isBlank()) {
            sql.append(" AND entity_type = ?");
            params.add(entityType);
        }

        sql.append(" ORDER BY timestamp DESC LIMIT ? OFFSET ?");
        int cappedSize = Math.min(size, 100);
        params.add(cappedSize);
        params.add(page * cappedSize);

        return jdbc.query(sql.toString(), ROW_MAPPER, params.toArray());
    }

    /**
     * Conta total de registros com os mesmos filtros (para paginação).
     */
    public int countFiltered(OffsetDateTime startDate, OffsetDateTime endDate,
                             String action, String entityType) {
        StringBuilder sql = new StringBuilder(
                "SELECT COUNT(*) FROM fbso_platform.audit_log WHERE deleted_dt IS NULL");
        List<Object> params = new ArrayList<>();

        if (startDate != null) {
            sql.append(" AND timestamp >= ?");
            params.add(startDate);
        }
        if (endDate != null) {
            sql.append(" AND timestamp <= ?");
            params.add(endDate);
        }
        if (action != null && !action.isBlank()) {
            sql.append(" AND action = ?");
            params.add(action);
        }
        if (entityType != null && !entityType.isBlank()) {
            sql.append(" AND entity_type = ?");
            params.add(entityType);
        }

        Integer result = jdbc.queryForObject(sql.toString(), Integer.class, params.toArray());
        return result != null ? result : 0;
    }
}
