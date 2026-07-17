package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.Plan;
import com.fbso.platform.admin.enums.Recurrence;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

public class PlanRowMapper implements RowMapper<Plan> {

    @Override
    public Plan mapRow(ResultSet rs, int rowNum) throws SQLException {
        Plan p = new Plan();
        p.setId(rs.getObject("id", UUID.class));
        p.setName(rs.getString("name"));
        p.setDescription(rs.getString("description"));
        p.setPrice(rs.getBigDecimal("price"));

        String recStr = rs.getString("recurrence");
        if (recStr != null) {
            p.setRecurrence(Recurrence.valueOf(recStr));
        }

        p.setStatus(rs.getString("status"));
        p.setVersion(rs.getInt("version"));

        p.setCreatedDt(rs.getObject("created_dt", java.time.OffsetDateTime.class));
        p.setUpdatedDt(rs.getObject("updated_dt", java.time.OffsetDateTime.class));
        p.setCreatedBy(rs.getObject("created_by", UUID.class));
        p.setUpdatedBy(rs.getObject("updated_by", UUID.class));
        p.setDeletedDt(rs.getObject("deleted_dt", java.time.OffsetDateTime.class));
        p.setDeletedBy(rs.getObject("deleted_by", UUID.class));

        return p;
    }
}
