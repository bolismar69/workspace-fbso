package com.fbso.platform.admin.unit.security;

import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.security.aspect.RbacAspect;
import com.fbso.platform.admin.service.PermissionService;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.reflect.MethodSignature;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.doThrow;
import static org.mockito.Mockito.doNothing;

@ExtendWith(MockitoExtension.class)
@DisplayName("RbacAspect (DB-backed via PermissionService)")
class RbacAspectTest {

    @Mock private ProceedingJoinPoint joinPoint;
    @Mock private MethodSignature methodSignature;
    @Mock private PermissionService permissionService;

    private RbacAspect aspect;

    // Métodos anotados para teste
    @RequiresPermission(resource = "PRODUCT_SERVICE", action = "edit")
    public void annotatedEditMethod() {}

    @RequiresPermission(resource = "TENANT", action = "create")
    public void annotatedCreateMethod() {}

    @RequiresPermission(resource = "AUDIT", action = "view")
    public void annotatedAuditMethod() {}

    @BeforeEach
    void setUp() {
        aspect = new RbacAspect(permissionService);
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    // ---- ADMIN acesso total (bypass por role) ----

    @Nested
    @DisplayName("ADMIN_TENANT — acesso total")
    class AdminGranted {

        @Test
        @DisplayName("ADMIN_TENANT pode editar produto (PermissionService não lança)")
        void shouldAllowAdminToEditProduct() throws Throwable {
            doNothing().when(permissionService).checkPermission("PRODUCT_SERVICE", "edit");

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedEditMethod")
                    .getAnnotation(RequiresPermission.class);

            aspect.checkPermission(joinPoint, annotation);
            // Não lança exceção → sucesso
        }
    }

    // ---- OPERATOR sem permissão → 403 ----

    @Nested
    @DisplayName("OPERATOR_BU — acesso negado")
    class OperatorDenied {

        @Test
        @DisplayName("OPERATOR tentando editar produto → PermissionDeniedException")
        void shouldDenyOperatorEditingProduct() throws Throwable {
            doThrow(new PermissionDeniedException())
                    .when(permissionService).checkPermission("PRODUCT_SERVICE", "edit");

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedEditMethod")
                    .getAnnotation(RequiresPermission.class);

            assertThatThrownBy(() -> aspect.checkPermission(joinPoint, annotation))
                    .isInstanceOf(PermissionDeniedException.class);
        }

        @Test
        @DisplayName("OPERATOR tentando criar tenant → PermissionDeniedException")
        void shouldDenyOperatorCreatingTenant() throws Throwable {
            doThrow(new PermissionDeniedException())
                    .when(permissionService).checkPermission("TENANT", "create");

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedCreateMethod")
                    .getAnnotation(RequiresPermission.class);

            assertThatThrownBy(() -> aspect.checkPermission(joinPoint, annotation))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ---- AUDITOR acesso restrito ----

    @Nested
    @DisplayName("AUDITOR — acesso restrito")
    class AuditorAccess {

        @Test
        @DisplayName("AUDITOR pode ver auditoria (PermissionService não lança)")
        void shouldAllowAuditorToViewAudit() throws Throwable {
            doNothing().when(permissionService).checkPermission("AUDIT", "view");

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedAuditMethod")
                    .getAnnotation(RequiresPermission.class);

            aspect.checkPermission(joinPoint, annotation);
            // Não lança exceção → sucesso
        }

        @Test
        @DisplayName("AUDITOR não pode criar tenant")
        void shouldDenyAuditorCreatingTenant() throws Throwable {
            doThrow(new PermissionDeniedException())
                    .when(permissionService).checkPermission("TENANT", "create");

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedCreateMethod")
                    .getAnnotation(RequiresPermission.class);

            assertThatThrownBy(() -> aspect.checkPermission(joinPoint, annotation))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }
}
