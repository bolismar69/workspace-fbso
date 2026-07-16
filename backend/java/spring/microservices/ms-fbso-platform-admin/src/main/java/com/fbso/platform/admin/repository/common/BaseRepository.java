package com.fbso.platform.admin.repository.common;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.security.TenantContext;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Template base para todos os repositories do projeto.
 * <p>
 * Implementa automaticamente:
 * <ul>
 *   <li><strong>Soft Delete:</strong> queries incluem {@code WHERE deleted_dt IS NULL}</li>
 *   <li><strong>Tenant Filter:</strong> queries incluem {@code WHERE tenant_id = ?} quando aplicável</li>
 *   <li><strong>Auditoria:</strong> {@code save()} preenche {@code created_by/updated_by}</li>
 * </ul>
 * <p>
 * ADR-L01: JDBC Template (não JPA/Hibernate) — controle total sobre SQL.
 *
 * @param <T> tipo da entidade (deve estender {@link BaseEntity})
 * @see <a href="ARCHITECTURE.md#5.1">ARCHITECTURE.md §5.1</a>
 */
public abstract class BaseRepository<T extends BaseEntity> {

    protected final JdbcTemplate jdbc;
    protected final String tableName;
    protected final RowMapper<T> rowMapper;
    protected final boolean hasTenantColumn;

    protected BaseRepository(JdbcTemplate jdbc, String tableName,
                             RowMapper<T> rowMapper, boolean hasTenantColumn) {
        this.jdbc = jdbc;
        this.tableName = tableName;
        this.rowMapper = rowMapper;
        this.hasTenantColumn = hasTenantColumn;
    }

    // ---- Queries Base (Soft Delete + Tenant Filter) ----

    /**
     * Busca todos os registros ativos, com tenant filter se aplicável.
     */
    public List<T> findAll(int page, int size, String sortColumn) {
        String tenantClause = hasTenantColumn ? " AND tenant_id = ?" : "";
        String sql = "SELECT * FROM fbso_platform." + tableName
                   + " WHERE deleted_dt IS NULL"
                   + tenantClause
                   + " ORDER BY " + sanitizeColumn(sortColumn) + " DESC"
                   + " LIMIT ? OFFSET ?";

        if (hasTenantColumn) {
            return jdbc.query(sql, rowMapper,
                    TenantContext.getTenantId(), size, page * size);
        }
        return jdbc.query(sql, rowMapper, size, page * size);
    }

    /**
     * Busca registro ativo por ID, com tenant filter se aplicável.
     */
    public Optional<T> findById(UUID id) {
        String tenantClause = hasTenantColumn ? " AND tenant_id = ?" : "";
        String sql = "SELECT * FROM fbso_platform." + tableName
                   + " WHERE id = ? AND deleted_dt IS NULL"
                   + tenantClause;

        List<T> results;
        if (hasTenantColumn) {
            results = jdbc.query(sql, rowMapper, id, TenantContext.getTenantId());
        } else {
            results = jdbc.query(sql, rowMapper, id);
        }
        return results.isEmpty() ? Optional.empty() : Optional.of(results.get(0));
    }

    /**
     * Conta total de registros ativos (para paginação).
     */
    public int count() {
        String tenantClause = hasTenantColumn ? " AND tenant_id = ?" : "";
        String sql = "SELECT COUNT(*) FROM fbso_platform." + tableName
                   + " WHERE deleted_dt IS NULL" + tenantClause;

        Integer result;
        if (hasTenantColumn) {
            result = jdbc.queryForObject(sql, Integer.class, TenantContext.getTenantId());
        } else {
            result = jdbc.queryForObject(sql, Integer.class);
        }
        return result != null ? result : 0;
    }

    // ---- Soft Delete ----

    /**
     * Soft delete: seta {@code deleted_dt = NOW()} e {@code deleted_by}.
     * O registro permanece no banco mas não aparece em queries.
     */
    public void softDelete(UUID id, UUID deletedBy) {
        String tenantClause = hasTenantColumn ? " AND tenant_id = ?" : "";
        String sql = "UPDATE fbso_platform." + tableName
                   + " SET deleted_dt = ?, deleted_by = ?"
                   + " WHERE id = ? AND deleted_dt IS NULL"
                   + tenantClause;

        int updated;
        if (hasTenantColumn) {
            updated = jdbc.update(sql, OffsetDateTime.now(), deletedBy,
                    id, TenantContext.getTenantId());
        } else {
            updated = jdbc.update(sql, OffsetDateTime.now(), deletedBy, id);
        }

        if (updated == 0) {
            throw new IllegalStateException(
                "Registro não encontrado ou já excluído: " + tableName + "." + id);
        }
    }

    // ---- Helpers ----

    /**
     * Sanitiza nome de coluna para evitar SQL injection em ORDER BY.
     */
    private String sanitizeColumn(String column) {
        if (column == null || column.isBlank()) {
            return "created_dt";
        }
        // permite apenas letras, números e underscore
        if (!column.matches("^[a-zA-Z0-9_]+$")) {
            throw new IllegalArgumentException("Nome de coluna inválido: " + column);
        }
        return column;
    }
}
