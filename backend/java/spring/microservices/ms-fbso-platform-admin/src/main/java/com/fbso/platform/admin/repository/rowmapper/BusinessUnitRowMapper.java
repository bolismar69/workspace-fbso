package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.BusinessUnit;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

/**
 * {@link RowMapper} para a entidade {@link BusinessUnit} —
 * mapeia {@code fbso_platform.business_unit}.
 *
 * <p>Mapeia 15 colunas de domínio + 6 colunas de auditoria (via
 * {@link AuditFieldsRowMapper}). Compatível com schema V001 + V007
 * (is_matrix).</p>
 *
 * <p><b>DT-134 (Sprint 6 F1):</b> Criado junto com
 * {@code BusinessUnitRepository.findTree()} (ADR-L08).</p>
 */
public class BusinessUnitRowMapper implements RowMapper<BusinessUnit> {

    @Override
    public BusinessUnit mapRow(ResultSet rs, int rowNum) throws SQLException {
        BusinessUnit bu = new BusinessUnit();

        bu.setId(rs.getObject("id", UUID.class));
        bu.setTenantId(rs.getObject("tenant_id", UUID.class));
        bu.setParentId(rs.getObject("parent_id", UUID.class));
        bu.setCnpj(rs.getString("cnpj"));
        bu.setCorporateName(rs.getString("corporate_name"));
        bu.setTaxRegime(rs.getString("tax_regime"));
        bu.setStreet(rs.getString("street"));
        bu.setNumber(rs.getString("number"));
        bu.setComplement(rs.getString("complement"));
        bu.setNeighborhood(rs.getString("neighborhood"));
        bu.setCity(rs.getString("city"));
        bu.setState(rs.getString("state"));
        bu.setZipCode(rs.getString("zip_code"));
        bu.setStatus(rs.getString("status"));
        bu.setMatrix(rs.getBoolean("is_matrix"));

        // Campos de auditoria (BaseEntity) — delegado ao helper (DT-086)
        AuditFieldsRowMapper.mapAuditFields(rs, bu);

        return bu;
    }
}
