package com.fbso.platform.admin.service;

import com.fbso.platform.admin.enums.Role;
import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.security.TenantContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Service;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Serviço de permissões RBAC — fonte autoritativa para verificação de acesso.
 * <p>
 * Substitui a matriz hardcoded do {@code RbacAspect} (Fase 0) por consultas
 * ao banco ({@code user_permission}, {@code role_resource}, {@code resource_action}).
 * <p>
 * <b>Estratégia de merge JWT×DB (DT-050):</b> O banco é a fonte primária.
 * Roles são obtidas de {@code user_permission} pelo {@code user_id} do contexto,
 * NÃO do JWT. Isso garante RN11-03 ("efeito imediato") — alterações de permissão
 * têm efeito na próxima requisição, sem esperar refresh do token.
 * <p>
 * <b>Estratégia de cache (DT-051):</b> A matriz RN10-01 tem tamanho fixo
 * (4 roles × ~28 resource_actions = ~112 linhas). Carregada em memória no
 * startup via {@code @PostConstruct} e recarregada sob demanda. Sem TTL —
 * a matriz só muda com deploy de nova migration de seed.
 *
 * @see com.fbso.platform.admin.security.aspect.RbacAspect
 */
@Service
public class PermissionService {

    private static final Logger log = LoggerFactory.getLogger(PermissionService.class);

    private final JdbcTemplate jdbc;

    /**
     * Matriz de permissões carregada do banco.
     * Key: role name (ex: "MANAGER_BU")
     * Value: set de "RESOURCE:action" (ex: "BUSINESS_UNIT:edit")
     */
    private final Map<String, Set<String>> permissionMatrix = new ConcurrentHashMap<>();

    public PermissionService(JdbcTemplate jdbc) {
        this.jdbc = jdbc;
        loadPermissionMatrix();
    }

    // ---- Carregamento da Matriz ----

    /**
     * Carrega a matriz RN10-01 completa do banco.
     * Executado no startup e pode ser re-executado para recarregar após nova seed.
     */
    public void loadPermissionMatrix() {
        String sql = """
            SELECT rr.role, ra.resource_name, ra.action
            FROM fbso_platform.role_resource rr
            JOIN fbso_platform.resource_action ra ON rr.resource_action_id = ra.id
            ORDER BY rr.role, ra.resource_name, ra.action
            """;

        Map<String, Set<String>> matrix = new ConcurrentHashMap<>();
        jdbc.query(sql, rs -> {
            String role = rs.getString("role");
            String resource = rs.getString("resource_name");
            String action = rs.getString("action");
            matrix.computeIfAbsent(role, k -> new HashSet<>())
                  .add(resource + ":" + action);
        });

        permissionMatrix.clear();
        permissionMatrix.putAll(matrix);
        log.info("Matriz RBAC carregada: {} roles, {} permissões totais",
                matrix.size(), matrix.values().stream().mapToInt(Set::size).sum());
    }

    // ---- Verificação de Permissão ----

    /**
     * Verifica se o usuário atual tem permissão para o recurso + ação.
     * <p>
     * ADMIN_TENANT tem acesso implícito total (não requer registros em user_permission).
     * Demais roles são verificadas contra a matriz carregada do banco.
     *
     * @param resource nome do recurso (ex: "TENANT")
     * @param action   ação solicitada (ex: "create")
     * @throws PermissionDeniedException se o acesso for negado
     */
    public void checkPermission(String resource, String action) {
        List<String> roles = getUserRoles();

        if (roles.isEmpty()) {
            log.warn("RBAC: sem roles no contexto — acesso negado a {}:{}", resource, action);
            throw new PermissionDeniedException();
        }

        // Admin tem acesso total implícito
        if (roles.contains(Role.ADMIN_TENANT.name())) {
            log.debug("RBAC: admin acessa {}:{}", resource, action);
            return;
        }

        // Verificar cada role do usuário contra a matriz
        String permissionKey = resource + ":" + action;
        boolean granted = roles.stream().anyMatch(role -> {
            Set<String> permissions = permissionMatrix.get(role);
            return permissions != null && permissions.contains(permissionKey);
        });

        if (!granted) {
            log.warn("RBAC: acesso negado — roles={}, resource={}, action={}",
                    roles, resource, action);
            throw new PermissionDeniedException();
        }

        log.debug("RBAC: acesso permitido — roles={}, resource={}, action={}",
                roles, resource, action);
    }

    /**
     * Obtém os roles do usuário atual a partir do banco ({@code user_permission}).
     * <p>
     * Fonte primária (DT-050): banco, não JWT. Isso garante que alterações
     * de permissão tenham efeito imediato (RN11-03).
     */
    public List<String> getUserRoles() {
        UUID userId = TenantContext.getUserIdQuietly();
        if (userId == null) {
            return List.of();
        }

        // ADMIN_TENANT: acesso implícito total — não requer registros em user_permission
        // Para outros papéis, consultar user_permission
        String sql = """
            SELECT up.role
            FROM fbso_platform.user_permission up
            WHERE up.user_id = ? AND up.deleted_dt IS NULL
            """;

        List<String> roles = jdbc.query(sql,
                (rs, rowNum) -> rs.getString("role"),
                userId);

        // Se não há registros em user_permission, verificar se é admin via JWT
        // (fallback para transição gradual JWT→DB)
        if (roles.isEmpty()) {
            List<String> jwtRoles = TenantContext.getRoles();
            if (jwtRoles.contains(Role.ADMIN_TENANT.name())) {
                return List.of(Role.ADMIN_TENANT.name());
            }
        }

        return roles;
    }

    // ---- Validação de Escopo por Business Unit (DT-067) ----

    /**
     * Valida que o usuário tem acesso à Business Unit especificada.
     * <p>
     * ADMIN_TENANT: acesso total (bypass).
     * MANAGER_BU e OPERATOR_BU: deve ter a BU na lista de BUs designadas do JWT.
     *
     * @param businessUnitId ID da BU sendo acessada
     * @throws PermissionDeniedException se o usuário não tiver acesso à BU
     */
    public void validateBusinessUnitAccess(UUID businessUnitId) {
        List<String> roles = getUserRoles();

        // Admin tem acesso total
        if (roles.contains(Role.ADMIN_TENANT.name())) {
            return;
        }

        List<UUID> assignedBus = TenantContext.getBusinessUnitIds();
        if (assignedBus.isEmpty() || !assignedBus.contains(businessUnitId)) {
            log.warn("RBAC: acesso negado à BU {} — roles={}, BUs designadas={}",
                    businessUnitId, roles, assignedBus);
            throw new PermissionDeniedException();
        }
    }
}
