package com.fbso.platform.admin.repository;

import com.fbso.platform.admin.entity.User;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.repository.rowmapper.UserRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * Repository para a entidade {@link User} — tabela {@code fbso_platform."user"}.
 * <p>
 * RLS ativo (V003) — filtrado por {@code tenant_id} automaticamente
 * via {@link BaseRepository#tenantClause()}.
 * <p>
 * F03-01: CRUD de usuários com email único por tenant ativo (RN09-02).
 *
 * @see User
 * @see UserRowMapper
 */
@Repository
public class UserRepository extends BaseRepository<User> {

    private static final UserRowMapper ROW_MAPPER = new UserRowMapper();

    public UserRepository(JdbcTemplate jdbc) {
        super(jdbc, "\"user\"", ROW_MAPPER, true);
    }

    /**
     * Busca usuário por email + tenant (apenas ativos — sem soft delete).
     * Usa o índice parcial {@code unique_email_active} (V002).
     *
     * @param email    email do usuário
     * @param tenantId tenant do contexto
     * @return usuário se encontrado, vazio caso contrário
     */
    public Optional<User> findByEmailAndTenant(String email, UUID tenantId) {
        String sql = "SELECT * FROM fbso_platform.\"user\""
                   + " WHERE email = ? AND tenant_id = ? AND deleted_dt IS NULL";
        List<User> results = jdbc.query(sql, ROW_MAPPER, email, tenantId);
        return results.isEmpty() ? Optional.empty() : Optional.of(results.get(0));
    }

    /**
     * Lista usuários ativos do tenant (com soft delete filter).
     *
     * @param tenantId tenant do contexto
     * @return lista de usuários ordenados por nome
     */
    public List<User> findAllByTenant(UUID tenantId) {
        String sql = "SELECT * FROM fbso_platform.\"user\""
                   + " WHERE tenant_id = ? AND deleted_dt IS NULL"
                   + " ORDER BY name ASC";
        return jdbc.query(sql, ROW_MAPPER, tenantId);
    }
}
