package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.BusinessUnitRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Repository para {@link BusinessUnit} — extende {@link BaseRepository}
 * com suporte a queries hierárquicas e validação de CNPJ.
 *
 * <h3>ADR-L08 — PostgreSQL WITH RECURSIVE</h3>
 * <p>A query {@link #findTree(UUID)} usa Common Table Expression (CTE)
 * recursiva do PostgreSQL. ADR-L08 documentado em ARCHITECTURE.md §10.</p>
 *
 * <p><b>DT-134 (Sprint 6 F1):</b> findTree() + findChildren().
 * <b>M6 (Sprint 6 F2):</b> existsByCnpj() + findByCnpj() para RN17-01.</p>
 */
@Repository
public class BusinessUnitRepository extends BaseRepository<BusinessUnit> {

    public BusinessUnitRepository(JdbcTemplate jdbc) {
        super(jdbc, "business_unit", new BusinessUnitRowMapper(), true);
    }

    /**
     * Retorna a árvore hierárquica completa de BUs do tenant (ADR-L08).
     */
    public List<BusinessUnit> findTree(UUID tenantId) {
        String sql = """
            WITH RECURSIVE bu_tree AS (
                SELECT bu.*, 0 AS depth, bu.corporate_name AS sort_path
                FROM fbso_platform.business_unit bu
                WHERE bu.tenant_id = ? AND bu.parent_id IS NULL AND bu.deleted_dt IS NULL
                UNION ALL
                SELECT child.*, bt.depth + 1,
                       bt.sort_path || ' > ' || child.corporate_name
                FROM fbso_platform.business_unit child
                INNER JOIN bu_tree bt ON child.parent_id = bt.id
                WHERE child.deleted_dt IS NULL
            )
            SELECT * FROM bu_tree ORDER BY sort_path
            """;
        return jdbc.query(sql, rowMapper, tenantId);
    }

    /**
     * Retorna BUs filhas diretas de uma BU pai.
     */
    public List<BusinessUnit> findChildren(UUID parentId, UUID tenantId) {
        if (parentId == null) {
            return jdbc.query("""
                SELECT * FROM fbso_platform.business_unit
                WHERE tenant_id = ? AND parent_id IS NULL AND deleted_dt IS NULL
                ORDER BY corporate_name
                """, rowMapper, tenantId);
        }
        return jdbc.query("""
            SELECT * FROM fbso_platform.business_unit
            WHERE tenant_id = ? AND parent_id = ? AND deleted_dt IS NULL
            ORDER BY corporate_name
            """, rowMapper, tenantId, parentId);
    }

    /**
     * Verifica se CNPJ já existe entre BUs ativas do tenant (RN17-01).
     */
    public boolean existsByCnpj(String cnpj, UUID tenantId) {
        String sql = """
            SELECT COUNT(*) FROM fbso_platform.business_unit
            WHERE tenant_id = ? AND cnpj = ? AND deleted_dt IS NULL
            """;
        Integer count = jdbc.queryForObject(sql, Integer.class, tenantId, cnpj);
        return count != null && count > 0;
    }

    /**
     * Busca BU por CNPJ no tenant.
     */
    public Optional<BusinessUnit> findByCnpj(String cnpj, UUID tenantId) {
        String sql = """
            SELECT * FROM fbso_platform.business_unit
            WHERE tenant_id = ? AND cnpj = ? AND deleted_dt IS NULL
            """;
        List<BusinessUnit> results = jdbc.query(sql, rowMapper, tenantId, cnpj);
        return results.isEmpty() ? Optional.empty() : Optional.of(results.get(0));
    }
}
