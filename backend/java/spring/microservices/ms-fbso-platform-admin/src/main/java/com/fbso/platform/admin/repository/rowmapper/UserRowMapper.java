package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.User;
import com.fbso.platform.admin.enums.UserStatus;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

/**
 * RowMapper para a entidade {@link User} — tabela {@code fbso_platform."user"}.
 */
public class UserRowMapper implements RowMapper<User> {

    @Override
    public User mapRow(ResultSet rs, int rowNum) throws SQLException {
        User u = new User();
        u.setId(rs.getObject("id", UUID.class));
        u.setTenantId(rs.getObject("tenant_id", UUID.class));
        u.setExternalKeycloakId(rs.getObject("external_keycloak_id", UUID.class));
        u.setEmail(rs.getString("email"));
        u.setName(rs.getString("name"));
        String statusStr = rs.getString("status");
        if (statusStr != null) {
            u.setStatus(UserStatus.valueOf(statusStr));
        }
        // Auditoria
        u.setCreatedDt(rs.getObject("created_dt", java.time.OffsetDateTime.class));
        u.setUpdatedDt(rs.getObject("updated_dt", java.time.OffsetDateTime.class));
        u.setCreatedBy(rs.getObject("created_by", UUID.class));
        u.setUpdatedBy(rs.getObject("updated_by", UUID.class));
        u.setDeletedDt(rs.getObject("deleted_dt", java.time.OffsetDateTime.class));
        u.setDeletedBy(rs.getObject("deleted_by", UUID.class));
        return u;
    }
}
