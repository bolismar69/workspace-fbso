package com.fbso.platform.admin.unit.controller;

import com.fbso.platform.admin.controller.OnboardingController;
import com.fbso.platform.admin.dto.response.OnboardingStatusResponse;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.OnboardingService;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.HttpStatus;

import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("OnboardingController")
class OnboardingControllerTest {

    @Mock private OnboardingService onboardingService;
    private OnboardingController controller;
    private UUID tenantId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        controller = new OnboardingController(onboardingService);
        TenantContext.set(tenantId, UUID.randomUUID(), List.of("ADMIN_TENANT"), List.of(), List.of());
    }

    @AfterEach
    void tearDown() { TenantContext.clear(); }

    @Test
    @DisplayName("GET /status deve retornar 200 com progresso")
    void shouldReturnStatus() {
        when(onboardingService.getStatus(tenantId)).thenReturn(
            OnboardingStatusResponse.of("STEP1_DONE", List.of("STEP1"), 25, "PENDING_ONBOARDING"));

        var response = controller.getStatus();
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.OK);
        assertThat(response.getBody().currentStep()).isEqualTo("STEP1_DONE");
        assertThat(response.getBody().progressPercent()).isEqualTo(25);
    }

    @Test
    @DisplayName("POST /complete deve retornar tenant ACTIVE")
    void shouldCompleteOnboarding() {
        when(onboardingService.complete(tenantId)).thenReturn(OnboardingStatusResponse.completed());

        var response = controller.complete();
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.OK);
        assertThat(response.getBody().currentStep()).isEqualTo("COMPLETED");
        assertThat(response.getBody().tenantStatus()).isEqualTo("ACTIVE");
    }
}
