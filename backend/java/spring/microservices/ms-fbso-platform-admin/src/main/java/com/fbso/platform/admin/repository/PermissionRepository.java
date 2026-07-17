package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.UserPermission;
import com.fbso.platform.admin.repository.rowmapper.UserPermissionRowMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Repository para a tabela {@code fbso_platform.user_permission}.
 *
 * <p>Não estende {@code BaseRepository} porque {@code user_permission}:
 * <ul>
 *   <li>Não tem soft delete — remoção é física (revogação de permissão)</li>
 *   <li>Não tem tenant_id próprio — o tenant é herdado da BU vinculada</li>
 * </ul>
 *
 * <h3>Constraints</h3>
 * <ul>
 *   <li>UNIQUE(user_id, business_unit_id) — validado no banco</li>
 *   <li>FK user_id → "user"(id)</li>
 *   <li>FK business_unit_id → business_unit(id) — V006</li>
 * </ul>
 */
@Repository
public class PermissionRepository {

    private static final Logger log = LoggerFactory.getLogger(PermissionRepository.class);
    private static final UserPermissionRowMapper ROW_MAPPER = new UserPermissionRowMapper();

    private final JdbcTemplate jdbc;

    public PermissionRepository(JdbcTemplate jdbc) {
        this.jdbc = jdbc;
    }

    /**
     * Lista todas as permissões de um usuário.
     */
    public List<UserPermission> findByUser(UUID userId) {
        String sql = "SELECT * FROM fbso_platform.user_permission WHERE user_id = ?";
        return jdbc.query(sql, ROW_MAPPER, userId);
    }

    /**
     * Busca permissão específica usuário×BU.
     */
    public Optional<UserPermission> findByUserAndBu(UUID userId, UUID businessUnitId) {
        String sql = "SELECT * FROM fbso_platform.user_permission"
                   + " WHERE user_id = ? AND business_unit_id = ?";
        List<UserPermission> results = jdbc.query(sql, ROW_MAPPER, userId, businessUnitId);
        return results.isEmpty() ? Optional.empty() : Optional.of(results.get(0));
    }

    /**
     * Atribui uma permissão (role) a um usuário para uma BU.
     * <p>
     * Se já existir (UNIQUE constraint), atualiza o role.
     */
    public void assign(UUID userId, UUID businessUnitId, String role) {
        jdbc.update("""
            INSERT INTO fbso_platform.user_permission (user_id, business_unit_id, role)
            VALUES (?, ?, ?)
            ON CONFLICT (user_id, business_unit_id)
            DO UPDATE SET role = EXCLUDED.role
            """, userId, businessUnitId, role);
        log.debug("Permissão atribuída: userId={}, buId={}, role={}", userId, businessUnitId, role);
    }

    /**
     * Revoga (remove) uma permissão de usuário para uma BU.
     */
    public void revoke(UUID userId, UUID businessUnitId) {
        int deleted = jdbc.update(
                "DELETE FROM fbso_platform.user_permission WHERE user_id = ? AND business_unit_id = ?",
                userId, businessUnitId);
        log.debug("Permissão revogada: userId={}, buId={}, rows={}", userId, businessUnitId, deleted);
    }

    /**
     * Lista os roles de um usuário (para o PermissionService.getUserRoles()).
     */
    public List<String> findRolesByUser(UUID userId) {
        String sql = "SELECT role FROM fbso_platform.user_permission WHERE user_id = ?";
        return jdbc.query(sql, (rs, rowNum) -> rs.getString("role"), userId);
    }
}
