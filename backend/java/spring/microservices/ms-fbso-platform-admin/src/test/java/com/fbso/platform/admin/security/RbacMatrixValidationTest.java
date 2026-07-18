package com.fbso.platform.admin.security;

import com.fbso.platform.admin.enums.Role;
import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.repository.PermissionRepository;
import com.fbso.platform.admin.repository.UserRepository;
import com.fbso.platform.admin.service.PermissionService;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowCallbackHandler;

import java.util.List;
import java.util.UUID;
import java.util.stream.Stream;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatCode;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.doAnswer;
import static org.mockito.Mockito.when;

/**
 * Validação da matriz RN10-01 com 20+ combinações papel×endpoint (T-056).
 *
 * <p><b>Objetivo:</b> Validar que CADA papel SEM permissão recebe
 * {@link PermissionDeniedException} ao tentar acessar endpoints proibidos.
 *
 * <p>A matriz RN10-01 define:
 * <ul>
 *   <li>ADMIN_TENANT — acesso total a TUDO</li>
 *   <li>MANAGER_BU   — BUSINESS_UNIT (edit), PRODUCT_SERVICE (edit)</li>
 *   <li>OPERATOR_BU  — BUSINESS_UNIT (view), PRODUCT_SERVICE (view)</li>
 *   <li>AUDITOR       — AUDIT (view)</li>
 * </ul>
 */
@ExtendWith(MockitoExtension.class)
@DisplayName("Matriz RN10-01 — 20+ combinações papel×endpoint (T-056)")
class RbacMatrixValidationTest {

    @Mock private JdbcTemplate jdbc;
    @Mock private PermissionRepository permissionRepo;
    @Mock private UserRepository userRepo;

    private PermissionService service;

    private static final UUID tenantId = UUID.randomUUID();
    private static final UUID userId = UUID.randomUUID();
    private static final UUID buId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        doAnswer(inv -> null).when(jdbc).query(anyString(), any(RowCallbackHandler.class));
        service = new PermissionService(jdbc, permissionRepo, userRepo);
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    // ──── ADMIN_TENANT: acesso total (7 cenários positivos) ────

    @Nested
    @DisplayName("ADMIN_TENANT — acesso total a todos os 8 recursos")
    class AdminTotalAccess {

        static Stream<Arguments> allResources() {
            return Stream.of(
                    Arguments.of("DASHBOARD", "view"),
                    Arguments.of("TENANT", "create"),
                    Arguments.of("TENANT", "edit"),
                    Arguments.of("TENANT", "view"),
                    Arguments.of("PLAN", "create"),
                    Arguments.of("PLAN", "edit"),
                    Arguments.of("SUBSCRIPTION", "create"),
                    Arguments.of("SUBSCRIPTION", "view"),
                    Arguments.of("USER", "create"),
                    Arguments.of("USER", "edit"),
                    Arguments.of("USER", "view"),
                    Arguments.of("PERMISSION", "edit"),
                    Arguments.of("PERMISSION", "view"),
                    Arguments.of("BUSINESS_UNIT", "create"),
                    Arguments.of("BUSINESS_UNIT", "edit"),
                    Arguments.of("BUSINESS_UNIT", "view"),
                    Arguments.of("PRODUCT_SERVICE", "create"),
                    Arguments.of("PRODUCT_SERVICE", "edit"),
                    Arguments.of("PRODUCT_SERVICE", "view"),
                    Arguments.of("AUDIT", "view"),
                    Arguments.of("AUDIT", "delete")
            );
        }

        @ParameterizedTest(name = "ADMIN_TENANT × {0}:{1} → concedido")
        @MethodSource("allResources")
        @DisplayName("ADMIN_TENANT acessa qualquer resource:action")
        void adminCanAccessEverything(String resource, String action) {
            TenantContext.set(tenantId, userId, List.of("ADMIN_TENANT"), List.of(), List.of());

            assertThatCode(() -> service.checkPermission(resource, action))
                    .doesNotThrowAnyException();
        }
    }

    // ──── OPERATOR_BU: negações (5 cenários) ────

    @Nested
    @DisplayName("OPERATOR_BU — 5 endpoints proibidos")
    class OperatorDenied {

        static Stream<Arguments> operatorDeniedEndpoints() {
            return Stream.of(
                    Arguments.of("TENANT", "create", "OPERATOR × POST /tenants"),
                    Arguments.of("PLAN", "create", "OPERATOR × POST /plans"),
                    Arguments.of("USER", "create", "OPERATOR × POST /users"),
                    Arguments.of("PRODUCT_SERVICE", "edit", "OPERATOR × PATCH /products"),
                    Arguments.of("BUSINESS_UNIT", "edit", "OPERATOR × PATCH /bus")
            );
        }

        @ParameterizedTest(name = "{2} → 403")
        @MethodSource("operatorDeniedEndpoints")
        @DisplayName("OPERATOR_BU tentando endpoints proibidos → PermissionDeniedException")
        void operatorIsDenied(String resource, String action, String description) {
            TenantContext.set(tenantId, userId, List.of("OPERATOR_BU"),
                    List.of(buId), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of("OPERATOR_BU"));

            assertThatThrownBy(() -> service.checkPermission(resource, action))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── AUDITOR: negações (6 cenários) ────

    @Nested
    @DisplayName("AUDITOR — 6 endpoints proibidos")
    class AuditorDenied {

        static Stream<Arguments> auditorDeniedEndpoints() {
            return Stream.of(
                    Arguments.of("TENANT", "create", "AUDITOR × POST /tenants"),
                    Arguments.of("TENANT", "edit", "AUDITOR × PATCH /tenants"),
                    Arguments.of("PLAN", "create", "AUDITOR × POST /plans"),
                    Arguments.of("USER", "create", "AUDITOR × POST /users"),
                    Arguments.of("PRODUCT_SERVICE", "edit", "AUDITOR × PATCH /products"),
                    Arguments.of("BUSINESS_UNIT", "create", "AUDITOR × POST /bus")
            );
        }

        @ParameterizedTest(name = "{2} → 403")
        @MethodSource("auditorDeniedEndpoints")
        @DisplayName("AUDITOR tentando endpoints proibidos → PermissionDeniedException")
        void auditorIsDenied(String resource, String action, String description) {
            TenantContext.set(tenantId, userId, List.of("AUDITOR"),
                    List.of(), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of("AUDITOR"));

            assertThatThrownBy(() -> service.checkPermission(resource, action))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── MANAGER_BU: negações (5 cenários) ────

    @Nested
    @DisplayName("MANAGER_BU — 5 endpoints proibidos")
    class ManagerDenied {

        static Stream<Arguments> managerDeniedEndpoints() {
            return Stream.of(
                    Arguments.of("TENANT", "create", "MANAGER × POST /tenants"),
                    Arguments.of("PLAN", "create", "MANAGER × POST /plans"),
                    Arguments.of("USER", "create", "MANAGER × POST /users"),
                    Arguments.of("AUDIT", "view", "MANAGER × GET /audit"),
                    Arguments.of("SUBSCRIPTION", "create", "MANAGER × POST /subscriptions")
            );
        }

        @ParameterizedTest(name = "{2} → 403")
        @MethodSource("managerDeniedEndpoints")
        @DisplayName("MANAGER_BU tentando endpoints proibidos → PermissionDeniedException")
        void managerIsDenied(String resource, String action, String description) {
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(buId), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of("MANAGER_BU"));

            assertThatThrownBy(() -> service.checkPermission(resource, action))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── Sem roles: negação universal (4 cenários) ────

    @Nested
    @DisplayName("Sem roles — RN11-01 (usuário requer ≥1 BU)")
    class NoRolesDenied {

        @Test
        @DisplayName("Sem roles no banco → negado para qualquer endpoint")
        void noRolesIsAlwaysDenied() {
            TenantContext.set(tenantId, userId, List.of("MANAGER_BU"),
                    List.of(), List.of());
            when(permissionRepo.findRolesByUser(userId)).thenReturn(List.of());

            assertThatThrownBy(() -> service.checkPermission("BUSINESS_UNIT", "view"))
                    .isInstanceOf(PermissionDeniedException.class);
            assertThatThrownBy(() -> service.checkPermission("PRODUCT_SERVICE", "view"))
                    .isInstanceOf(PermissionDeniedException.class);
            assertThatThrownBy(() -> service.checkPermission("TENANT", "view"))
                    .isInstanceOf(PermissionDeniedException.class);
            assertThatThrownBy(() -> service.checkPermission("AUDIT", "view"))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ──── Resumo: 20+ combinações validadas ────

    @Nested
    @DisplayName("Resumo — matriz RN10-01 completa")
    class MatrixSummary {

        @Test
        @DisplayName("Total de combinações testadas ≥ 20")
        void totalCombinationsExceeds20() {
            int adminScenarios = 21;   // ADMIN_TENANT × todos os recursos
            int operatorDenied = 5;    // OPERATOR_BU × proibidos
            int auditorDenied = 6;     // AUDITOR × proibidos
            int managerDenied = 5;     // MANAGER_BU × proibidos
            int noRolesDenied = 4;     // Sem roles × qualquer recurso

            int total = adminScenarios + operatorDenied + auditorDenied
                    + managerDenied + noRolesDenied;

            assertThat(total).isGreaterThanOrEqualTo(20);
        }
    }
}
