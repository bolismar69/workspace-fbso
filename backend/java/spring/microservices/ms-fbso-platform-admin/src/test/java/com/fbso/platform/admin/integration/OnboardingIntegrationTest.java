package com.fbso.platform.admin.integration;

import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.OnboardingService;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.jdbc.core.JdbcTemplate;

import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
@DisplayName("Onboarding Integration")
class OnboardingIntegrationTest extends BaseIntegrationTest {

    @Autowired private OnboardingService onboardingService;
    @Autowired private JdbcTemplate jdbc;

    private UUID tenantId;

    @BeforeEach
    void setUp() {
        tenantId = UUID.randomUUID();
        // Criar tenant PENDING_ONBOARDING via JDBC direto
        jdbc.update("INSERT INTO fbso_platform.tenant (id, name_corporate, segment, status, onboarding_step, created_dt, updated_dt) " +
            "VALUES (?, 'Test Tenant', 'RETAIL', 'PENDING_ONBOARDING', 'NOT_STARTED', NOW(), NOW())",
            tenantId);
        TenantContext.set(tenantId, UUID.randomUUID(), List.of("ADMIN_TENANT"), List.of(), List.of());
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
        jdbc.update("DELETE FROM fbso_platform.tenant WHERE id = ?", tenantId);
    }

    @Test
    @DisplayName("Fluxo completo: NOT_STARTED → STEP1 → STEP2 → STEP3 → COMPLETED → ACTIVE")
    void shouldCompleteFullOnboardingFlow() {
        // Step 1
        var step1 = onboardingService.completeStep1(tenantId,
            new com.fbso.platform.admin.dto.request.OnboardingStep1Request("Empresa Teste", "Fantasia", "RETAIL"));
        assertThat(step1.currentStep()).isEqualTo(OnboardingService.STEP1_DONE);
        assertThat(step1.progressPercent()).isEqualTo(25);

        // Step 2
        var step2 = onboardingService.completeStep2(tenantId,
            new com.fbso.platform.admin.dto.request.OnboardingStep2Request("11222333444455", "SIMPLES_NACIONAL"));
        assertThat(step2.currentStep()).isEqualTo(OnboardingService.STEP2_DONE);
        assertThat(step2.progressPercent()).isEqualTo(50);

        // Step 3
        var step3 = onboardingService.completeStep3(tenantId,
            new com.fbso.platform.admin.dto.request.OnboardingStep3Request(null));
        assertThat(step3.currentStep()).isEqualTo(OnboardingService.STEP3_DONE);

        // Complete
        var complete = onboardingService.complete(tenantId);
        assertThat(complete.currentStep()).isEqualTo("COMPLETED");
        assertThat(complete.progressPercent()).isEqualTo(100);

        // Verificar no banco
        var tenant = jdbc.queryForMap("SELECT status, onboarding_step FROM fbso_platform.tenant WHERE id = ?", tenantId);
        assertThat(tenant.get("status").toString()).isEqualTo("ACTIVE");
        assertThat(tenant.get("onboarding_step").toString()).isEqualTo("COMPLETED");
    }
}
