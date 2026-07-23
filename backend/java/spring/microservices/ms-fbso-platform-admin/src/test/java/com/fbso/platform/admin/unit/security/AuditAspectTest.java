package com.fbso.platform.admin.unit.security;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.Auditable;
import com.fbso.platform.admin.security.aspect.AuditAspect;
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
import org.mockito.quality.Strictness;
import org.mockito.junit.jupiter.MockitoSettings;
import org.springframework.core.task.TaskExecutor;
import org.springframework.jdbc.core.JdbcTemplate;

import java.lang.reflect.Method;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.doAnswer;
import static org.mockito.Mockito.lenient;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoInteractions;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@MockitoSettings(strictness = Strictness.LENIENT)
@DisplayName("AuditAspect")
class AuditAspectTest {

    @Mock private JdbcTemplate jdbc;
    @Mock private TaskExecutor taskExecutor;
    @Mock private ProceedingJoinPoint joinPoint;
    @Mock private MethodSignature methodSignature;

    private final ObjectMapper objectMapper = new ObjectMapper();

    private AuditAspect aspect;

    private static final UUID TENANT_ID = UUID.fromString("00000000-0000-0000-0000-000000000001");
    private static final UUID USER_ID = UUID.fromString("00000000-0000-0000-0000-0000000000AA");

    @BeforeEach
    void setUp() {
        aspect = new AuditAspect(jdbc, taskExecutor, objectMapper);
        TenantContext.set(TENANT_ID, USER_ID, List.of("ADMIN_TENANT"), List.of(), List.of());

        // Faz o TaskExecutor executar sincronamente nos testes
        doAnswer(invocation -> {
            ((Runnable) invocation.getArgument(0)).run();
            return null;
        }).when(taskExecutor).execute(any(Runnable.class));

        // Mock proceed() para @Around — retorna null (void methods)
        try {
            lenient().when(joinPoint.proceed()).thenReturn(null);
            lenient().when(joinPoint.proceed(any())).thenReturn(null);
        } catch (Throwable ignored) {}
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    // ---- Anotações de teste (private, acessadas via reflection) ----

    @Auditable(entityType = "TENANT", action = "CREATED")
    private void createTenant(UUID tenantId) {}

    @Auditable(entityType = "SUBSCRIPTION", action = "PLAN_CHANGED", idParamName = "subscriptionId")
    private void changePlan(String reason, UUID subscriptionId) {}

    // ---- Helpers ----

    private Method getTestMethod(String name, Class<?>... paramTypes) throws NoSuchMethodException {
        return AuditAspectTest.class.getDeclaredMethod(name, paramTypes);
    }

    // ---- Testes ----

    @Nested
    @DisplayName("Auditoria com contexto válido")
    class ValidContext {

        @Test
        @DisplayName("deve gravar registro com tenantId e userId corretos (DT-002)")
        void shouldRecordAuditWithCorrectTenantAndUser() throws Throwable {
            UUID entityId = UUID.randomUUID();
            Method method = getTestMethod("createTenant", UUID.class);
            when(joinPoint.getSignature()).thenReturn(methodSignature);
            when(methodSignature.getMethod()).thenReturn(method);
            when(joinPoint.getArgs()).thenReturn(new Object[]{entityId});
            // DT-021: INSERT agora inclui previous_value e new_value (8 params)
            lenient().when(jdbc.update(anyString(), any(OffsetDateTime.class), eq(TENANT_ID),
                    anyString(), anyString(), eq(entityId), eq(USER_ID),
                    any(), any())).thenReturn(1);

            Auditable annotation = method.getAnnotation(Auditable.class);
            aspect.audit(joinPoint, annotation);

            verify(jdbc).update(anyString(), any(OffsetDateTime.class), eq(TENANT_ID),
                    anyString(), anyString(), eq(entityId), eq(USER_ID), any(), any());
        }

        @Test
        @DisplayName("deve extrair entityId via idParamName (DT-008)")
        void shouldExtractEntityIdViaNamedParam() throws Throwable {
            UUID subId = UUID.randomUUID();
            Method method = getTestMethod("changePlan", String.class, UUID.class);
            when(joinPoint.getSignature()).thenReturn(methodSignature);
            when(methodSignature.getMethod()).thenReturn(method);
            when(joinPoint.getArgs()).thenReturn(new Object[]{"upgrade requested", subId});
            lenient().when(jdbc.update(anyString(), any(OffsetDateTime.class), any(UUID.class),
                    anyString(), anyString(), eq(subId), any(UUID.class),
                    any(), any())).thenReturn(1);

            Auditable annotation = method.getAnnotation(Auditable.class);
            aspect.audit(joinPoint, annotation);

            verify(jdbc).update(anyString(), any(OffsetDateTime.class), any(UUID.class),
                    anyString(), anyString(), eq(subId), any(UUID.class), any(), any());
        }
    }

    @Nested
    @DisplayName("Auditoria sem contexto")
    class MissingContext {

        @Test
        @DisplayName("não deve gravar auditoria quando TenantContext falha (DT-002)")
        void shouldSkipAuditWhenTenantContextFails() throws Throwable {
            // Limpa o contexto para simular falha
            TenantContext.clear();

            UUID entityId = UUID.randomUUID();
            Method method = getTestMethod("createTenant", UUID.class);
            // Usar lenient() porque o aspecto retorna antes de usar estes stubs
            lenient().when(joinPoint.getSignature()).thenReturn(methodSignature);
            lenient().when(methodSignature.getMethod()).thenReturn(method);
            lenient().when(joinPoint.getArgs()).thenReturn(new Object[]{entityId});

            Auditable annotation = method.getAnnotation(Auditable.class);
            aspect.audit(joinPoint, annotation);

            // Não deve interagir com JDBC — auditoria pulada
            verifyNoInteractions(jdbc);
        }
    }
}
