package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.PermissionUpdateRequest;
import com.fbso.platform.admin.dto.response.PermissionResponse;
import com.fbso.platform.admin.entity.UserPermission;
import com.fbso.platform.admin.enums.Role;
import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.exception.UserNotFoundException;
import com.fbso.platform.admin.repository.PermissionRepository;
import com.fbso.platform.admin.repository.UserRepository;
import com.fbso.platform.admin.security.TenantContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

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
    private final PermissionRepository permissionRepo;
    private final UserRepository userRepo;

    /**
     * Matriz de permissões carregada do banco.
     * Key: role name (ex: "MANAGER_BU")
     * Value: set de "RESOURCE:action" (ex: "BUSINESS_UNIT:edit")
     */
    private final Map<String, Set<String>> permissionMatrix = new ConcurrentHashMap<>();

    public PermissionService(JdbcTemplate jdbc, PermissionRepository permissionRepo,
                             UserRepository userRepo) {
        this.jdbc = jdbc;
        this.permissionRepo = permissionRepo;
        this.userRepo = userRepo;
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
     * <b>ADMIN_TENANT</b> tem acesso implícito total — verificado via JWT
     * (Keycloak é a fonte autoritativa de autenticação). Não requer registros
     * em {@code user_permission}.
     * <p>
     * <b>Demais roles</b> (MANAGER_BU, OPERATOR_BU, AUDITOR) são verificadas
     * contra a matriz carregada do banco ({@code user_permission} +
     * {@code role_resource}). Isso garante RN11-03 ("efeito imediato") —
     * alterações de permissão têm efeito na próxima requisição, sem esperar
     * refresh do token JWT.
     *
     * @param resource nome do recurso (ex: "TENANT")
     * @param action   ação solicitada (ex: "create")
     * @throws PermissionDeniedException se o acesso for negado
     */
    public void checkPermission(String resource, String action) {
        // ADMIN_TENANT: acesso implícito total via JWT
        // Keycloak é a fonte autoritativa de autenticação — não requer
        // registros em user_permission. Isso permite que o admin inicial
        // (criado no Keycloak) funcione sem seed manual de permissões.
        if (isAdmin()) {
            log.debug("RBAC: admin acessa {}:{}", resource, action);
            return;
        }

        // Demais roles: banco como fonte primária (DT-050)
        List<String> dbRoles = getUserRoles();
        if (dbRoles.isEmpty()) {
            log.warn("RBAC: sem roles no banco — acesso negado a {}:{}", resource, action);
            throw new PermissionDeniedException();
        }

        // Verificar cada role do usuário contra a matriz RN10-01
        String permissionKey = resource + ":" + action;
        boolean granted = dbRoles.stream().anyMatch(role -> {
            Set<String> permissions = permissionMatrix.get(role);
            return permissions != null && permissions.contains(permissionKey);
        });

        if (!granted) {
            log.warn("RBAC: acesso negado — roles={}, resource={}, action={}",
                    dbRoles, resource, action);
            throw new PermissionDeniedException();
        }

        log.debug("RBAC: acesso permitido — roles={}, resource={}, action={}",
                dbRoles, resource, action);
    }

    /**
     * Verifica se o usuário atual é ADMIN_TENANT (via JWT/Keycloak).
     * <p>
     * Keycloak é a fonte autoritativa de autenticação. ADMIN_TENANT é um papel
     * de autenticação que concede acesso implícito total — não requer registros
     * em {@code user_permission}.
     */
    private boolean isAdmin() {
        return TenantContext.getRoles().contains(Role.ADMIN_TENANT.name());
    }

    /**
     * Obtém os roles do usuário atual a partir do banco ({@code user_permission}).
     * <p>
     * <b>Fonte primária (DT-050):</b> banco, não JWT. Isso garante que alterações
     * de permissão tenham efeito imediato (RN11-03).
     * <p>
     * <b>ADMIN_TENANT NÃO é retornado por este método.</b> Admin é verificado
     * separadamente via {@link #isAdmin()} no {@link #checkPermission(String, String)}.
     * Isso mantém a separação clara: Keycloak para autenticação, banco para autorização.
     *
     * @return roles do banco ({@code user_permission}), ou lista vazia se o
     *         usuário não tiver permissões explícitas
     */
    public List<String> getUserRoles() {
        UUID userId = TenantContext.getUserIdQuietly();
        if (userId == null) {
            return List.of();
        }
        return permissionRepo.findRolesByUser(userId);
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
        // Admin tem acesso total implícito a todas as BUs
        if (isAdmin()) {
            return;
        }

        List<UUID> assignedBus = TenantContext.getBusinessUnitIds();
        if (assignedBus.isEmpty() || !assignedBus.contains(businessUnitId)) {
            log.warn("RBAC: acesso negado à BU {} — BUs designadas={}",
                    businessUnitId, assignedBus);
            throw new PermissionDeniedException();
        }
    }

    // ---- Gestão de Permissões (Frente 2 — T-050) ----

    /**
     * Atribui um papel a um usuário para uma Business Unit.
     * <p>
     * Efeito imediato (RN11-03): a próxima chamada a {@code getUserRoles()}
     * já retorna o novo papel.
     *
     * @param userId         ID do usuário
     * @param businessUnitId ID da BU
     * @param role           papel a atribuir (ex: "MANAGER_BU")
     * @throws UserNotFoundException se o usuário não pertencer ao tenant atual
     */
    public void assignRole(UUID userId, UUID businessUnitId, String role) {
        validateUserTenant(userId);
        // TODO Frente 3/Sprint 6: validateBusinessUnitTenant(businessUnitId)
        //   quando BusinessUnitRepository for criado
        permissionRepo.assign(userId, businessUnitId, role);
        log.info("Role atribuído: userId={}, buId={}, role={}", userId, businessUnitId, role);
    }

    /**
     * Revoga um papel de um usuário para uma Business Unit.
     * <p>
     * Efeito imediato (RN11-03).
     *
     * @param userId         ID do usuário
     * @param businessUnitId ID da BU
     * @throws UserNotFoundException se o usuário não pertencer ao tenant atual
     */
    public void revokeRole(UUID userId, UUID businessUnitId) {
        validateUserTenant(userId);
        permissionRepo.revoke(userId, businessUnitId);
        log.info("Role revogado: userId={}, buId={}", userId, businessUnitId);
    }

    // ---- Consulta de Permissões (Frente 3 — T-051) ----

    /**
     * Lista as permissões atuais de um usuário.
     *
     * @param userId ID do usuário
     * @return lista de permissões (userId, businessUnitId, role)
     * @throws UserNotFoundException se o usuário não pertencer ao tenant
     */
    public List<PermissionResponse> getUserPermissions(UUID userId) {
        validateUserTenant(userId);
        return permissionRepo.findByUser(userId).stream()
                .map(PermissionResponse::from)
                .toList();
    }

    /**
     * Atualiza em lote as permissões de um usuário (substituição completa).
     * <p>
     * Remove todas as permissões atuais e insere as novas em uma única transação.
     * Efeito imediato (RN11-03).
     *
     * @param userId  ID do usuário
     * @param request lista de assignments [{businessUnitId, role}]
     * @throws UserNotFoundException se o usuário não pertencer ao tenant
     */
    @Transactional
    public List<PermissionResponse> updateUserPermissions(UUID userId,
                                                          PermissionUpdateRequest request) {
        validateUserTenant(userId);

        // 1. Revogar todas as permissões atuais
        List<UserPermission> current = permissionRepo.findByUser(userId);
        for (UserPermission up : current) {
            permissionRepo.revoke(userId, up.getBusinessUnitId());
        }

        // 2. Atribuir as novas permissões
        for (var assignment : request.permissions()) {
            assignRole(userId, assignment.businessUnitId(), assignment.role());
        }

        log.info("Permissões atualizadas em lote: userId={}, novas={}", userId,
                request.permissions().size());

        // 3. Retornar estado atual
        return getUserPermissions(userId);
    }

    /**
     * Lista as BUs que o usuário tem acesso.
     * <p>
     * ADMIN_TENANT: retorna lista vazia (acesso implícito a todas).
     * Demais papéis: retorna BUs de {@code user_permission}.
     */
    public List<UUID> getUserBusinessUnits(UUID userId) {
        if (isAdmin()) {
            return List.of(); // acesso implícito a todas
        }
        return permissionRepo.findByUser(userId).stream()
                .map(UserPermission::getBusinessUnitId)
                .distinct()
                .toList();
    }

    /** Valida que o usuário pertence ao tenant do contexto (defesa contra IDOR cross-tenant). */
    private void validateUserTenant(UUID userId) {
        UUID tenantId = TenantContext.getTenantId();
        userRepo.findById(userId).ifPresent(user -> {
            if (!tenantId.equals(user.getTenantId())) {
                log.warn("Tentativa de assign/revoke cross-tenant bloqueada: userId={}, tenant={}",
                        userId, tenantId);
                throw new UserNotFoundException(userId);
            }
        });
    }
}
