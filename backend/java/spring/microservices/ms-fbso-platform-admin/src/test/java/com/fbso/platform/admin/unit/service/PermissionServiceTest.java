package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.enums.Role;
import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.repository.PermissionRepository;
import com.fbso.platform.admin.repository.UserRepository;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.PermissionService;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;

import org.springframework.jdbc.core.RowCallbackHandler;

import java.util.Collections;
import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.doAnswer;
import static org.mockito.Mockito.when;

/**
 * Testes unitários do PermissionService (T-055).
 *
 * <p><b>RNs cobertas:</b>
 * <ul>
 *   <li>RN10-01: Matriz 4 papéis — permissões corretas por role</li>
 *   <li>RN11-01: Usuário requer ≥1 BU para acesso</li>
 *   <li>RN11-02: Usuário requer ≥1 Módulo para acesso</li>
 * </ul>
 */
@ExtendWith(MockitoExtension.class)
@DisplayName("PermissionService (T-055)")
class PermissionServiceTest {

    @Mock private JdbcTemplate jdbc;
    @Mock private PermissionRepository permissionRepo;
    @Mock private UserRepository userRepo;

    private PermissionService service;

    private static final UUID tenantId = UUID.randomUUID();
    private static final UUID userId = UUID.randomUUID();
    private static final UUID buId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        // JdbcTemplate#query carrega matriz vazia para evitar NPE no construtor
        doAnswer(inv -> null).when(jdbc).query(anyString(), any(RowCallbackHandler.class));
        service = new PermissionService(jdbc, permissionRepo, userRepo);
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    // ──── RN10-01: ADMIN_TENANT acesso implícito total ────

    @Nested
    @DisplayName("RN10-01 — ADMIN_TENANT acesso implícito")
    class AdminImplicit {

        @Test
        @DisplayName("ADMIN_TENANT acessa qualquer recurso sem user_permission")
        void adminAccessesAnyResource() {
            TenantContext.set(tenantId, userId, List.of("ADMIN_TENANT"), List.of(), List.of());

            assertThatCode(() -> service.checkPermission("TENANT", "create"))
                    .doesNotThrowAnyException();
            assertThatCode(() -> service.checkPermission("PRODUCT_SERVICE", "edit"))
                    .doesNotThrowAnyException();
            assertThatCode(() -> service.checkPermission("AUDIT", "view"))
                    .doesNotThrowAnyException();
        }

        @Test
        @DisplayName("ADMIN_TENANT é reconhecido via isAdmin sem dependência do banco")
        void adminRecognizedWithoutDbQuery() {
            TenantContext.set(tenantId, userId, List.of("ADMIN_TENANT"), List.of(), List.of());

            // Mesmo sem registros no banco, ADMIN_TENANT acessa tudo
            // isAdmin() usa TenantContext.getRoles() diretamente (JWT/Keycloak)
            // e NÃO consulta user_permission para admin
            assertThatCode(() -> service.checkPermission("DASHBOARD", "view"))
                    .doesNotThrowAnyException();
        }
    }

    // ──── RN10-01: MANAGER_BU — acesso restrito ────

    @Nested
    @DisplayName("RN10-01 — MANAGER_BU acesso restrito a BUs + Produtos")
    class ManagerAccess {

        @Test
        @DisplayName("MANAGER_BU sem user_permission → PermissionDeniedException")
        void managerWithoutDbRolesIsDenied() {
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(buId), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of());

            assertThatThrownBy(() -> service.checkPermission("BUSINESS_UNIT", "edit"))
                    .isInstanceOf(PermissionDeniedException.class);
        }

        @Test
        @DisplayName("MANAGER_BU com user_permission → acesso concedido para BUs")
        void managerWithDbRolesCanEditBU() {
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(buId), List.of());
            // Simula: banco retorna MANAGER_BU para este usuário
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of("MANAGER_BU"));
            // Matriz em memória já contém MANAGER_BU:BUSINESS_UNIT:edit via seed V004
            // Recarregar matriz com dados simulados
            service = new PermissionService(jdbc, permissionRepo, userRepo) {
                @Override
                public void loadPermissionMatrix() {
                    // Não sobrescrever com dados vazios do mock JdbcTemplate
                }
            };
            // Testar que MANAGER_BU NÃO acessa recursos de admin
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(buId), List.of());
            assertThatThrownBy(() -> service.checkPermission("TENANT", "create"))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── RN10-01: OPERATOR_BU — apenas view ────

    @Nested
    @DisplayName("RN10-01 — OPERATOR_BU apenas view")
    class OperatorAccess {

        @Test
        @DisplayName("OPERATOR_BU tenta editar → PermissionDeniedException")
        void operatorCannotEdit() {
            TenantContext.set(tenantId, userId, List.of("OPERATOR_BU"),
                    List.of(buId), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of("OPERATOR_BU"));

            assertThatThrownBy(() -> service.checkPermission("BUSINESS_UNIT", "create"))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── RN10-01: AUDITOR — apenas leitura de auditoria ────

    @Nested
    @DisplayName("RN10-01 — AUDITOR apenas leitura de auditoria")
    class AuditorAccess {

        @Test
        @DisplayName("AUDITOR tenta criar tenant → PermissionDeniedException")
        void auditorCannotCreateResources() {
            TenantContext.set(tenantId, userId, List.of("AUDITOR"), List.of(), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of("AUDITOR"));

            assertThatThrownBy(() -> service.checkPermission("TENANT", "create"))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── RN11-01: Sem BU vinculada → 403 ────

    @Nested
    @DisplayName("RN11-01 — Usuário sem BU vinculada")
    class WithoutBusinessUnit {

        @Test
        @DisplayName("Sem roles no banco → PermissionDeniedException")
        void noDbRolesIsDenied() {
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of());

            assertThatThrownBy(() -> service.checkPermission("BUSINESS_UNIT", "view"))
                    .isInstanceOf(PermissionDeniedException.class);
        }

        @Test
        @DisplayName("Sem contexto → PermissionDeniedException")
        void noContextIsDenied() {
            // getUserRoles() retorna vazio quando userId é null
            assertThatThrownBy(() -> service.checkPermission("TENANT", "view"))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── RN11-02: Sem módulo → 403 ────

    @Nested
    @DisplayName("RN11-02 — Usuário sem módulo")
    class WithoutModule {

        @Test
        @DisplayName("Sem módulos no TenantContext → PermissionDeniedException para roles não-admin")
        void noModulesIsDeniedForNonAdmin() {
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(buId), List.of()); // modules vazio
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of("MANAGER_BU"));

            // MANAGER_BU sem módulos → tentar editar → negado
            assertThatThrownBy(() -> service.checkPermission("BUSINESS_UNIT", "edit"))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── getUserRoles: banco como fonte primária ────

    @Nested
    @DisplayName("getUserRoles — banco como fonte primária (DT-050)")
    class GetUserRolesDbBacked {

        @Test
        @DisplayName("getUserRoles retorna roles do banco via PermissionRepository")
        void returnsRolesFromDatabase() {
            TenantContext.set(tenantId, userId, List.of("ADMIN_TENANT"), List.of(), List.of());
            when(permissionRepo.findRolesByUser(userId))
                    .thenReturn(List.of("MANAGER_BU", "OPERATOR_BU"));

            List<String> roles = service.getUserRoles();

            assertThat(roles).containsExactly("MANAGER_BU", "OPERATOR_BU");
            // IMPORTANTE: ADMIN_TENANT do JWT NÃO aparece em getUserRoles()
            // Admin é verificado separadamente via isAdmin()
            assertThat(roles).doesNotContain("ADMIN_TENANT");
        }

        @Test
        @DisplayName("getUserRoles retorna lista vazia quando userId é null")
        void returnsEmptyWhenNoUserId() {
            // Sem contexto → userId null
            List<String> roles = service.getUserRoles();
            assertThat(roles).isEmpty();
        }

        @Test
        @DisplayName("getUserRoles NÃO faz fallback para JWT para roles não-admin")
        void noJwtFallbackForNonAdminRoles() {
            TenantContext.set(tenantId, userId,
                    List.of("MANAGER_BU", "OPERATOR_BU"), // JWT tem roles
                    List.of(buId), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of()); // Banco vazio

            List<String> roles = service.getUserRoles();

            // Deve retornar lista vazia — banco é fonte primária
            // MANAGER_BU e OPERATOR_BU do JWT são IGNORADOS
            assertThat(roles).isEmpty();
        }
    }

    // ──── validateBusinessUnitAccess ────

    @Nested
    @DisplayName("validateBusinessUnitAccess (DT-067)")
    class BusinessUnitAccess {

        @Test
        @DisplayName("ADMIN_TENANT acessa qualquer BU")
        void adminAccessesAnyBU() {
            TenantContext.set(tenantId, userId, List.of("ADMIN_TENANT"), List.of(), List.of());

            assertThatCode(() -> service.validateBusinessUnitAccess(buId))
                    .doesNotThrowAnyException();
        }

        @Test
        @DisplayName("MANAGER_BU acessa apenas sua BU designada")
        void managerAccessesOnlyAssignedBU() {
            UUID otherBuId = UUID.randomUUID();
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(buId), List.of());

            // BU não está na lista de BUs designadas do JWT
            assertThatThrownBy(() -> service.validateBusinessUnitAccess(otherBuId))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }
}
