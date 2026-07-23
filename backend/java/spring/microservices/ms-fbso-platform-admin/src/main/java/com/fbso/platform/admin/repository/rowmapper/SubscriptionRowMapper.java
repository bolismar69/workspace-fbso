package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.Subscription;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

public class SubscriptionRowMapper implements RowMapper<Subscription> {

    @Override
    public Subscription mapRow(ResultSet rs, int rowNum) throws SQLException {
        Subscription s = new Subscription();
        s.setId(rs.getObject("id", UUID.class));
        s.setTenantId(rs.getObject("tenant_id", UUID.class));
        s.setPlanId(rs.getObject("plan_id", UUID.class));
        s.setStartDate(rs.getObject("start_date", java.time.OffsetDateTime.class));
        s.setEndDate(rs.getObject("end_date", java.time.OffsetDateTime.class));
        s.setStatus(rs.getString("status"));
        s.setLockedPrice(rs.getBigDecimal("locked_price"));

        String lr = rs.getString("locked_recurrence");
        if (lr != null) s.setLockedRecurrence(lr);

        // Campos de auditoria — delegado ao helper (DT-086)
        AuditFieldsRowMapper.mapAuditFields(rs, s);

        return s;
    }
}
