package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.UserPermission;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

/**
 * RowMapper para a entidade {@link UserPermission} — tabela {@code fbso_platform.user_permission}.
 */
public class UserPermissionRowMapper implements RowMapper<UserPermission> {

    @Override
    public UserPermission mapRow(ResultSet rs, int rowNum) throws SQLException {
        UserPermission up = new UserPermission();
        up.setId(rs.getObject("id", UUID.class));
        up.setUserId(rs.getObject("user_id", UUID.class));
        up.setBusinessUnitId(rs.getObject("business_unit_id", UUID.class));
        up.setRole(rs.getString("role"));
        return up;
    }
}
