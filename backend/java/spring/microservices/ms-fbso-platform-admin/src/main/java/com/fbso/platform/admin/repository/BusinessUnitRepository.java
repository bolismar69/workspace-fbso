package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.BusinessUnitRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

/**
 * Repository para {@link BusinessUnit} — extende {@link BaseRepository}
 * com suporte a queries hierárquicas.
 *
 * <h3>ADR-L08 — PostgreSQL WITH RECURSIVE</h3>
 * <p>A query {@link #findTree(UUID)} usa Common Table Expression (CTE)
 * recursiva do PostgreSQL para retornar a árvore hierárquica completa
 * em uma única query. Performance O(n) em vez de O(n²) com N+1 queries.
 * Sem limite de profundidade (RN17-04).</p>
 *
 * <p><b>DT-134 (Sprint 6 F1):</b> Repository criado com findTree().
 * ADR-L08 documentado em ARCHITECTURE.md §10.</p>
 *
 * @see <a href="ARCHITECTURE.md#10">ARCHITECTURE.md §10 — ADR-L08</a>
 */
@Repository
public class BusinessUnitRepository extends BaseRepository<BusinessUnit> {

    public BusinessUnitRepository(JdbcTemplate jdbc) {
        super(jdbc, "business_unit", new BusinessUnitRowMapper(), true);
    }

    /**
     * Retorna a árvore hierárquica completa de BUs do tenant, ordenada
     * com a Matriz primeiro e filhas por nível de profundidade.
     *
     * <p><b>ADR-L08:</b> Usa PostgreSQL {@code WITH RECURSIVE} CTE.
     * A CTE percorre a hierarquia a partir das raízes
     * ({@code parent_id IS NULL}) e desce recursivamente por
     * {@code parent_id}, acumulando profundidade e caminho de
     * ordenação.</p>
     *
     * <p><b>Ordenação:</b> As BUs são retornadas em ordem de caminho
     * (sort_path), que concatena os nomes corporativos ao longo da
     * hierarquia. Isso garante que a Matriz apareça primeiro, seguida
     * por suas filiais e sub-filiais em ordem alfabética por nível.</p>
     *
     * <p><b>Segurança:</b> O RLS do PostgreSQL (V003) filtra
     * automaticamente por {@code tenant_id = current_setting(...)}.
     * O parâmetro {@code tenantId} fornecido é usado como condição
     * base da CTE para eficiência — o RLS atua como segunda camada
     * de defesa.</p>
     *
     * @param tenantId UUID do tenant cuja árvore de BUs será carregada
     * @return lista de BusinessUnit ordenada hierarquicamente
     */
    public List<BusinessUnit> findTree(UUID tenantId) {
        String sql = """
            WITH RECURSIVE bu_tree AS (
                -- Caso base: raízes (Matriz e filhas soltas)
                SELECT
                    bu.*,
                    0 AS depth,
                    bu.corporate_name AS sort_path
                FROM fbso_platform.business_unit bu
                WHERE bu.tenant_id = ?
                  AND bu.parent_id IS NULL
                  AND bu.deleted_dt IS NULL

                UNION ALL

                -- Passo recursivo: desce um nível na hierarquia
                SELECT
                    child.*,
                    bt.depth + 1,
                    bt.sort_path || ' > ' || child.corporate_name
                FROM fbso_platform.business_unit child
                INNER JOIN bu_tree bt ON child.parent_id = bt.id
                WHERE child.deleted_dt IS NULL
            )
            SELECT * FROM bu_tree
            ORDER BY sort_path
            """;
        return jdbc.query(sql, rowMapper, tenantId);
    }

    /**
     * Retorna BUs filhas diretas de uma BU pai (1 nível de profundidade).
     * Útil para navegação incremental ou lazy loading da árvore.
     *
     * <p><b>Segurança:</b> Inclui filtro {@code tenant_id} explícito como
     * camada 2 de defesa em profundidade (ADR-L07). O RLS do PostgreSQL
     * atua como camada 1.</p>
     *
     * @param parentId UUID da BU pai (ou null para raízes)
     * @param tenantId UUID do tenant para filtro de isolamento
     * @return lista de BusinessUnit filhas diretas
     */
    public List<BusinessUnit> findChildren(UUID parentId, UUID tenantId) {
        if (parentId == null) {
            String sql = """
                SELECT * FROM fbso_platform.business_unit
                WHERE tenant_id = ?
                  AND parent_id IS NULL
                  AND deleted_dt IS NULL
                ORDER BY corporate_name
                """;
            return jdbc.query(sql, rowMapper, tenantId);
        }
        String sql = """
            SELECT * FROM fbso_platform.business_unit
            WHERE tenant_id = ?
              AND parent_id = ?
              AND deleted_dt IS NULL
            ORDER BY corporate_name
            """;
        return jdbc.query(sql, rowMapper, tenantId, parentId);
    }
}
