package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.ProductService;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.ProductServiceRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

/**
 * Repository para ProductService — catálogo de produtos/serviços (F04-06).
 *
 * <p>V009 (Sprint 6 F1): {@code hasTenantColumn=true} — tenant_id
 * adicionado à tabela com RLS FORCE.</p>
 */
@Repository
public class ProductRepository extends BaseRepository<ProductService> {

    public ProductRepository(JdbcTemplate jdbc) {
        super(jdbc, "product_service", new ProductServiceRowMapper(), true);
    }

    /**
     * Lista produtos de uma BU específica (RN18-01 — catálogo segmentado).
     */
    public List<ProductService> findByBusinessUnit(UUID businessUnitId) {
        String sql = """
            SELECT * FROM fbso_platform.product_service
            WHERE business_unit_id = ? AND deleted_dt IS NULL
            ORDER BY name
            """;
        return jdbc.query(sql, rowMapper, businessUnitId);
    }

    /**
     * Verifica se SKU já existe na BU (RN18-02 — único por BU ativo).
     */
    public boolean existsBySku(String sku, UUID businessUnitId) {
        if (sku == null || sku.isBlank()) return false;
        String sql = """
            SELECT COUNT(*) FROM fbso_platform.product_service
            WHERE business_unit_id = ? AND sku = ? AND deleted_dt IS NULL
            """;
        Integer count = jdbc.queryForObject(sql, Integer.class, businessUnitId, sku);
        return count != null && count > 0;
    }
}
