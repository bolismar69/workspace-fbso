package com.fbso.platform.admin.unit.security;

import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.security.aspect.RbacAspect;
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

import java.lang.reflect.Method;
import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("RbacAspect")
class RbacAspectTest {

    @Mock private ProceedingJoinPoint joinPoint;
    @Mock private MethodSignature methodSignature;

    private final RbacAspect aspect = new RbacAspect();

    // Método anotado para teste
    @RequiresPermission(resource = "PRODUCT_SERVICE", action = "edit")
    public void annotatedEditMethod() {}

    @RequiresPermission(resource = "TENANT", action = "create")
    public void annotatedCreateMethod() {}

    @RequiresPermission(resource = "AUDIT", action = "view")
    public void annotatedAuditMethod() {}

    @BeforeEach
    void setUp() throws NoSuchMethodException {
        // signature é usado apenas por testes que chamam joinPoint.proceed()
        // (ex: AdminGranted). Demais testes não chegam ao proceed().
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    // ---- TC-S2-012: OPERATOR sem permissão → 403 ----

    @Nested
    @DisplayName("TC-S2-012 — OPERATOR sem permissão")
    class OperatorDenied {

        @Test
        @DisplayName("OPERATOR tentando editar produto → PermissionDeniedException")
        void shouldDenyOperatorEditingProduct() throws Throwable {
            TenantContext.set(UUID.randomUUID(), UUID.randomUUID(),
                    List.of("OPERATOR_BU"), List.of(), List.of());

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedEditMethod")
                    .getAnnotation(RequiresPermission.class);

            assertThatThrownBy(() -> aspect.checkPermission(joinPoint, annotation))
                    .isInstanceOf(PermissionDeniedException.class);
        }

        @Test
        @DisplayName("OPERATOR tentando criar tenant → PermissionDeniedException")
        void shouldDenyOperatorCreatingTenant() throws Throwable {
            TenantContext.set(UUID.randomUUID(), UUID.randomUUID(),
                    List.of("OPERATOR_BU"), List.of(), List.of());

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedCreateMethod")
                    .getAnnotation(RequiresPermission.class);

            assertThatThrownBy(() -> aspect.checkPermission(joinPoint, annotation))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }

    // ---- TC-S2-013: ADMIN acesso total ----

    @Nested
    @DisplayName("TC-S2-013 — ADMIN acesso total")
    class AdminGranted {

        @Test
        @DisplayName("ADMIN_TENANT pode editar produto")
        void shouldAllowAdminToEditProduct() throws Throwable {
            TenantContext.set(UUID.randomUUID(), UUID.randomUUID(),
                    List.of("ADMIN_TENANT"), List.of(), List.of());

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedEditMethod")
                    .getAnnotation(RequiresPermission.class);

            // Não deve lançar exceção
            aspect.checkPermission(joinPoint, annotation);
        }
    }

    // ---- TC-S2-014: AUDITOR acesso apenas leitura de auditoria ----

    @Nested
    @DisplayName("TC-S2-014 — AUDITOR acesso restrito")
    class AuditorGranted {

        @Test
        @DisplayName("AUDITOR pode ver auditoria")
        void shouldAllowAuditorToViewAudit() throws Throwable {
            TenantContext.set(UUID.randomUUID(), UUID.randomUUID(),
                    List.of("AUDITOR"), List.of(), List.of());

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedAuditMethod")
                    .getAnnotation(RequiresPermission.class);

            aspect.checkPermission(joinPoint, annotation);
        }

        @Test
        @DisplayName("AUDITOR não pode criar tenant")
        void shouldDenyAuditorCreatingTenant() throws Throwable {
            TenantContext.set(UUID.randomUUID(), UUID.randomUUID(),
                    List.of("AUDITOR"), List.of(), List.of());

            RequiresPermission annotation = RbacAspectTest.class
                    .getMethod("annotatedCreateMethod")
                    .getAnnotation(RequiresPermission.class);

            assertThatThrownBy(() -> aspect.checkPermission(joinPoint, annotation))
                    .isInstanceOf(PermissionDeniedException.class);
        }
    }
}
