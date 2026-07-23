package com.fbso.platform.admin.repository.rowmapper;

import com.fbso.platform.admin.entity.ProductService;
import com.fbso.platform.admin.enums.ProductType;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

/**
 * RowMapper para ProductService — mapeia {@code fbso_platform.product_service}.
 */
public class ProductServiceRowMapper implements RowMapper<ProductService> {

    @Override
    public ProductService mapRow(ResultSet rs, int rowNum) throws SQLException {
        ProductService ps = new ProductService();

        ps.setId(rs.getObject("id", UUID.class));
        ps.setTenantId(rs.getObject("tenant_id", UUID.class));
        ps.setBusinessUnitId(rs.getObject("business_unit_id", UUID.class));
        ps.setName(rs.getString("name"));
        ps.setSku(rs.getString("sku"));

        String typeStr = rs.getString("type");
        if (typeStr != null) {
            ps.setType(ProductType.valueOf(typeStr));
        }

        ps.setDescription(rs.getString("description"));
        ps.setStatus(rs.getString("status"));

        AuditFieldsRowMapper.mapAuditFields(rs, ps);
        return ps;
    }
}
