package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.TenantRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

/**
 * Repository para a entidade {@link Tenant} — tabela {@code fbso_platform.tenant}.
 * <p>
 * A tabela tenant é GLOBAL (visível apenas pelo Admin FBSO), sem tenant_id próprio.
 * {@code hasTenantColumn = false}.
 * <p>
 * F01-02: Lista paginada com filtros de status, plano e busca textual (≥3 chars).
 *
 * @see Tenant
 * @see TenantRowMapper
 */
@Repository
public class TenantRepository extends BaseRepository<Tenant> {

    private static final TenantRowMapper ROW_MAPPER = new TenantRowMapper();

    public TenantRepository(JdbcTemplate jdbc) {
        super(jdbc, "tenant", ROW_MAPPER, false); // tenant não tem tenant_id
    }

    /**
     * Busca paginada com filtros opcionais de status, plano e busca textual.
     *
     * @param page   página (0-based)
     * @param size   itens por página (padrão 25)
     * @param status filtrar por status (opcional)
     * @param plan   filtrar por nome do plano (opcional, via subscription ativa)
     * @param search busca textual em name_corporate e name_fantasy (≥3 chars, ILIKE)
     * @return lista de tenants ordenados por created_dt DESC
     */
    public List<Tenant> findAllPaginated(int page, int size, String status,
                                          String plan, String search) {
        StringBuilder sql = new StringBuilder(
                "SELECT DISTINCT t.* FROM fbso_platform.tenant t");

        // JOIN com subscription + plan se filtro de plano estiver ativo
        if (plan != null && !plan.isBlank()) {
            sql.append(" JOIN fbso_platform.subscription s ON t.id = s.tenant_id"
                     + " AND s.status = 'ACTIVE' AND s.deleted_dt IS NULL"
                     + " JOIN fbso_platform.plan p ON s.plan_id = p.id"
                     + " AND p.deleted_dt IS NULL");
        }

        sql.append(" WHERE t.deleted_dt IS NULL");

        // Filtro de status
        if (status != null && !status.isBlank()) {
            sql.append(" AND t.status = ?");
        }

        // Filtro de plano
        if (plan != null && !plan.isBlank()) {
            sql.append(" AND p.name = ?");
        }

        // Busca textual (case-insensitive, ≥3 caracteres)
        if (search != null && search.trim().length() >= 3) {
            sql.append(" AND (t.name_corporate ILIKE ? OR t.name_fantasy ILIKE ?)");
        }

        sql.append(" ORDER BY t.created_dt DESC");
        sql.append(" LIMIT ? OFFSET ?");

        // Construir array de parâmetros
        Object[] params = buildFilterParams(status, plan, search, size, page * size);
        return jdbc.query(sql.toString(), ROW_MAPPER, params);
    }

    /**
     * Conta total com os mesmos filtros (para paginação).
     */
    public int countFiltered(String status, String plan, String search) {
        StringBuilder sql = new StringBuilder(
                "SELECT COUNT(DISTINCT t.id) FROM fbso_platform.tenant t");

        if (plan != null && !plan.isBlank()) {
            sql.append(" JOIN fbso_platform.subscription s ON t.id = s.tenant_id"
                     + " AND s.status = 'ACTIVE' AND s.deleted_dt IS NULL"
                     + " JOIN fbso_platform.plan p ON s.plan_id = p.id"
                     + " AND p.deleted_dt IS NULL");
        }

        sql.append(" WHERE t.deleted_dt IS NULL");

        if (status != null && !status.isBlank()) {
            sql.append(" AND t.status = ?");
        }
        if (plan != null && !plan.isBlank()) {
            sql.append(" AND p.name = ?");
        }
        if (search != null && search.trim().length() >= 3) {
            sql.append(" AND (t.name_corporate ILIKE ? OR t.name_fantasy ILIKE ?)");
        }

        Object[] params = buildFilterParams(status, plan, search);
        Integer result = jdbc.queryForObject(sql.toString(), Integer.class, params);
        return result != null ? result : 0;
    }

    /**
     * Busca tenant por razão social exata (para validação de duplicidade — RN04-02).
     */
    public Optional<Tenant> findByNameCorporate(String nameCorporate) {
        String sql = "SELECT * FROM fbso_platform.tenant"
                   + " WHERE name_corporate = ? AND deleted_dt IS NULL";
        List<Tenant> results = jdbc.query(sql, ROW_MAPPER, nameCorporate);
        return results.isEmpty() ? Optional.empty() : Optional.of(results.get(0));
    }

    // ---- Helpers ----

    /**
     * Constrói o array de parâmetros para queries com filtros.
     * Evita duplicação de lógica entre findAllPaginated e countFiltered.
     */
    private Object[] buildFilterParams(String status, String plan, String search,
                                        Object... extraParams) {
        // Contar quantos parâmetros teremos
        int count = 0;
        if (status != null && !status.isBlank()) count++;
        if (plan != null && !plan.isBlank()) count++;
        if (search != null && search.trim().length() >= 3) count += 2; // LIKE para 2 colunas
        count += extraParams.length;

        Object[] params = new Object[count];
        int idx = 0;

        if (status != null && !status.isBlank()) {
            params[idx++] = status;
        }
        if (plan != null && !plan.isBlank()) {
            params[idx++] = plan;
        }
        if (search != null && search.trim().length() >= 3) {
            String like = "%" + search.trim() + "%";
            params[idx++] = like;
            params[idx++] = like;
        }
        for (Object extra : extraParams) {
            params[idx++] = extra;
        }
        return params;
    }
}
