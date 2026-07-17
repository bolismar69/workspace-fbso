package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.AuditEntry;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

public class AuditEntryRowMapper implements RowMapper<AuditEntry> {

    @Override
    public AuditEntry mapRow(ResultSet rs, int rowNum) throws SQLException {
        AuditEntry e = new AuditEntry();
        e.setId(rs.getObject("id", UUID.class));
        e.setTimestamp(rs.getObject("timestamp", java.time.OffsetDateTime.class));
        e.setTenantId(rs.getObject("tenant_id", UUID.class));
        e.setAction(rs.getString("action"));
        e.setEntityType(rs.getString("entity_type"));
        e.setEntityId(rs.getObject("entity_id", UUID.class));
        e.setActorId(rs.getObject("actor_id", UUID.class));
        e.setActorName(rs.getString("actor_name"));
        e.setPreviousValue(rs.getString("previous_value"));
        e.setNewValue(rs.getString("new_value"));
        e.setReason(rs.getString("reason"));
        e.setCreatedDt(rs.getObject("created_dt", java.time.OffsetDateTime.class));
        return e;
    }
}
