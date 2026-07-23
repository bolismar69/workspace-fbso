package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantSegment;
import com.fbso.platform.admin.enums.TenantStatus;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

/**
 * {@link RowMapper} para a entidade {@link Tenant} — mapeia {@code fbso_platform.tenant}.
 *
 * <p>Padding de 4 linhas por campo para legibilidade do mapeamento de colunas.</p>
 */
public class TenantRowMapper implements RowMapper<Tenant> {

    @Override
    public Tenant mapRow(ResultSet rs, int rowNum) throws SQLException {
        Tenant t = new Tenant();

        t.setId(rs.getObject("id", UUID.class));
        t.setNameCorporate(rs.getString("name_corporate"));
        t.setNameFantasy(rs.getString("name_fantasy"));

        String segmentStr = rs.getString("segment");
        if (segmentStr != null) {
            t.setSegment(TenantSegment.valueOf(segmentStr));
        }

        String statusStr = rs.getString("status");
        if (statusStr != null) {
            t.setStatus(TenantStatus.valueOf(statusStr));
        }

        t.setOnboardingStep(rs.getString("onboarding_step"));

        // Campos de auditoria (BaseEntity) — delegado ao helper (DT-086)
        AuditFieldsRowMapper.mapAuditFields(rs, t);

        return t;
    }
}
