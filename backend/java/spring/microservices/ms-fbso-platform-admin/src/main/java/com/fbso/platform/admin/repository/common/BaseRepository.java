package com.fbso.platform.admin.repository.common;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.security.TenantContext;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.jdbc.core.namedparam.MapSqlParameterSource;
import org.springframework.jdbc.core.namedparam.NamedParameterJdbcTemplate;
import org.springframework.jdbc.support.GeneratedKeyHolder;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;
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
        String sql = "SELECT * FROM fbso_platform." + tableName
                   + " WHERE deleted_dt IS NULL"
                   + tenantClause()
                   + " ORDER BY " + sanitizeColumn(sortColumn) + " DESC"
                   + " LIMIT ? OFFSET ?";

        Object[] params = buildParams(size, page * size);
        return jdbc.query(sql, rowMapper, params);
    }

    /**
     * Busca registro ativo por ID, com tenant filter se aplicável.
     */
    public Optional<T> findById(UUID id) {
        String sql = "SELECT * FROM fbso_platform." + tableName
                   + " WHERE id = ? AND deleted_dt IS NULL"
                   + tenantClause();

        Object[] params = buildParams(id);
        List<T> results = jdbc.query(sql, rowMapper, params);
        return results.isEmpty() ? Optional.empty() : Optional.of(results.get(0));
    }

    /**
     * Conta total de registros ativos (para paginação).
     */
    public int count() {
        String sql = "SELECT COUNT(*) FROM fbso_platform." + tableName
                   + " WHERE deleted_dt IS NULL" + tenantClause();

        Object[] params = buildParams();
        Integer result = jdbc.queryForObject(sql, Integer.class, params);
        return result != null ? result : 0;
    }

    // ---- Soft Delete ----

    /**
     * Soft delete: seta {@code deleted_dt = NOW()} e {@code deleted_by}.
     * O registro permanece no banco mas não aparece em queries.
     */
    public void softDelete(UUID id, UUID deletedBy) {
        String sql = "UPDATE fbso_platform." + tableName
                   + " SET deleted_dt = ?, deleted_by = ?"
                   + " WHERE id = ? AND deleted_dt IS NULL"
                   + tenantClause();

        // DT-029: buildParams() centraliza o branching hasTenantColumn
        Object[] params = buildParams(OffsetDateTime.now(), deletedBy, id);

        int updated = jdbc.update(sql, params);
        if (updated == 0) {
            throw new IllegalStateException(
                "Registro não encontrado ou já excluído: " + tableName + "." + id);
        }
    }

    // ---- INSERT / UPDATE (T-015.4.DT-003) ----

    /**
     * Insere um novo registro com preenchimento automático de
     * {@code id}, {@code created_dt}, {@code updated_dt}, {@code created_by},
     * {@code updated_by} e {@code tenant_id} (se aplicável).
     *
     * @param entity entidade a persistir
     */
    public void save(T entity) {
        Map<String, Object> columns = entity.toColumnMap();
        UUID id = entity.getId() != null ? entity.getId() : UUID.randomUUID();
        UUID currentUser = TenantContext.getUserIdQuietly();
        OffsetDateTime now = OffsetDateTime.now();

        // build INSERT
        StringBuilder sql = new StringBuilder("INSERT INTO fbso_platform.")
                .append(tableName).append(" (id");
        StringBuilder values = new StringBuilder(" VALUES (?");
        for (String col : columns.keySet()) {
            sql.append(", ").append(col);
            values.append(", ?");
        }
        sql.append(", created_dt, updated_dt, created_by, updated_by");
        values.append(", ?, ?, ?, ?");
        if (hasTenantColumn) {
            sql.append(", tenant_id");
            values.append(", ?");
        }
        sql.append(")").append(values).append(")");

        // build params: 1 (id) + columns + 4 (audit: created_dt, updated_dt, created_by, updated_by) + tenant
        Object[] params = new Object[5 + columns.size() + (hasTenantColumn ? 1 : 0)];
        int idx = 0;
        params[idx++] = id;
        for (Object value : columns.values()) {
            params[idx++] = value;
        }
        params[idx++] = now;          // created_dt
        params[idx++] = now;          // updated_dt
        params[idx++] = currentUser;  // created_by
        params[idx++] = currentUser;  // updated_by
        if (hasTenantColumn) {
            params[idx] = TenantContext.getTenantId();
        }

        jdbc.update(sql.toString(), params);

        // set ID back on entity for chaining
        if (entity.getId() == null) {
            entity.setId(id);
        }
    }

    /**
     * Atualiza um registro existente com preenchimento automático de
     * {@code updated_dt} e {@code updated_by}.
     *
     * @param entity entidade a atualizar (deve ter ID preenchido)
     * @throws IllegalStateException se nenhuma linha foi atualizada
     */
    public void update(T entity) {
        Map<String, Object> columns = entity.toColumnMap();
        if (columns.isEmpty()) {
            return; // nada para atualizar
        }

        UUID currentUser = TenantContext.getUserIdQuietly();
        OffsetDateTime now = OffsetDateTime.now();

        // build UPDATE
        StringBuilder sql = new StringBuilder("UPDATE fbso_platform.")
                .append(tableName).append(" SET ");
        for (String col : columns.keySet()) {
            sql.append(col).append(" = ?, ");
        }
        sql.append("updated_dt = ?, updated_by = ?")
           .append(" WHERE id = ? AND deleted_dt IS NULL")
           .append(tenantClause());

        // build params
        Object[] params = new Object[columns.size() + 2 + 1 + (hasTenantColumn ? 1 : 0)];
        int idx = 0;
        for (Object value : columns.values()) {
            params[idx++] = value;
        }
        params[idx++] = now;          // updated_dt
        params[idx++] = currentUser;  // updated_by
        params[idx++] = entity.getId();
        if (hasTenantColumn) {
            params[idx] = TenantContext.getTenantId();
        }

        int updated = jdbc.update(sql.toString(), params);
        if (updated == 0) {
            throw new IllegalStateException(
                "Registro não encontrado ou já excluído: " + tableName + "." + entity.getId());
        }
    }

    // ---- Helpers ----

    /**
     * Retorna cláusula SQL de tenant filter, ou string vazia se a tabela não tem tenant_id.
     * Centraliza a lógica de branching para evitar duplicação (DT-029).
     */
    private String tenantClause() {
        return hasTenantColumn ? " AND tenant_id = ?" : "";
    }

    /**
     * Constrói array de parâmetros para queries, anexando tenant_id se aplicável.
     * Centraliza a lógica de branching para evitar duplicação (DT-029).
     */
    private Object[] buildParams(Object... baseParams) {
        if (!hasTenantColumn) {
            return baseParams;
        }
        Object[] params = new Object[baseParams.length + 1];
        System.arraycopy(baseParams, 0, params, 0, baseParams.length);
        params[baseParams.length] = TenantContext.getTenantId();
        return params;
    }

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
