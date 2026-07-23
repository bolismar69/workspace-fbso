package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.common.BaseEntity;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

/**
 * Helper que centraliza o mapeamento dos 6 campos de auditoria
 * ({@code created_dt, updated_dt, created_by, updated_by, deleted_dt, deleted_by})
 * do {@link ResultSet} para {@link BaseEntity}.
 *
 * <p>Elimina duplicação em 4 RowMappers (Tenant, Plan, Subscription, User)
 * que repetiam exatamente as mesmas 6 linhas de mapeamento.</p>
 *
 * <p>Uso:</p>
 * <pre>{@code
 *   AuditFieldsRowMapper.mapAuditFields(rs, entity);
 * }</pre>
 *
 * @see BaseEntity
 */
public final class AuditFieldsRowMapper {

    private AuditFieldsRowMapper() {
        // Classe utilitária — não instanciável
    }

    /**
     * Popula os 6 campos de auditoria da {@link BaseEntity} a partir
     * das colunas do {@link ResultSet}.
     *
     * @param rs     ResultSet posicionado na linha atual
     * @param entity entidade que receberá os valores de auditoria
     * @throws SQLException se alguma coluna não existir no ResultSet
     */
    public static void mapAuditFields(ResultSet rs, BaseEntity entity) throws SQLException {
        entity.setCreatedDt(rs.getObject("created_dt", java.time.OffsetDateTime.class));
        entity.setUpdatedDt(rs.getObject("updated_dt", java.time.OffsetDateTime.class));
        entity.setCreatedBy(rs.getObject("created_by", UUID.class));
        entity.setUpdatedBy(rs.getObject("updated_by", UUID.class));
        entity.setDeletedDt(rs.getObject("deleted_dt", java.time.OffsetDateTime.class));
        entity.setDeletedBy(rs.getObject("deleted_by", UUID.class));
    }
}
