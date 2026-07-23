package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.request.OnboardingStep1Request;
import com.fbso.platform.admin.dto.request.OnboardingStep2Request;
import com.fbso.platform.admin.dto.request.OnboardingStep3Request;
import com.fbso.platform.admin.dto.response.OnboardingStatusResponse;
import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.repository.TenantRepository;
import com.fbso.platform.admin.service.OnboardingService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("OnboardingService")
class OnboardingServiceTest {

    @Mock private TenantRepository tenantRepo;
    private OnboardingService service;
    private UUID tenantId;
    private Tenant tenant;

    @BeforeEach
    void setUp() {
        service = new OnboardingService(tenantRepo);
        tenantId = UUID.randomUUID();
        tenant = new Tenant();
        tenant.setId(tenantId);
        tenant.setStatus(TenantStatus.PENDING_ONBOARDING);
        tenant.setOnboardingStep(OnboardingService.NOT_STARTED);
        when(tenantRepo.findById(tenantId)).thenReturn(Optional.of(tenant));
    }

    @Nested
    @DisplayName("Step 1 — Confirmar Dados")
    class Step1 {
        @Test
        @DisplayName("deve avançar de NOT_STARTED para STEP1_DONE")
        void shouldAdvanceToStep1Done() {
            var req = new OnboardingStep1Request("Empresa Ltda", "Fantasia", "RETAIL");
            var result = service.completeStep1(tenantId, req);
            assertThat(result.currentStep()).isEqualTo(OnboardingService.STEP1_DONE);
            assertThat(result.progressPercent()).isEqualTo(25);
            verify(tenantRepo).update(any());
        }

        @Test
        @DisplayName("deve rejeitar se já concluiu step 1 — EC-3")
        void shouldRejectIfAlreadyDone() {
            tenant.setOnboardingStep(OnboardingService.STEP1_DONE);
            var req = new OnboardingStep1Request("Empresa Ltda", null, "RETAIL");
            assertThatThrownBy(() -> service.completeStep1(tenantId, req))
                .isInstanceOf(BusinessException.class)
                .hasMessageContaining("Conclua o passo anterior");
        }
    }

    @Nested
    @DisplayName("Step 2 — CNPJ + Matriz")
    class Step2 {
        @Test
        @DisplayName("deve validar CNPJ (14 dígitos)")
        void shouldValidateCnpj() {
            tenant.setOnboardingStep(OnboardingService.STEP1_DONE);
            var req = new OnboardingStep2Request("123", "SIMPLES_NACIONAL");
            assertThatThrownBy(() -> service.completeStep2(tenantId, req))
                .isInstanceOf(BusinessException.class)
                .hasMessageContaining("CNPJ inválido");
        }

        @Test
        @DisplayName("CNPJ inválido não avança estado — EC-2")
        void shouldStayOnStep1WhenCnpjInvalid() {
            tenant.setOnboardingStep(OnboardingService.STEP1_DONE);
            var req = new OnboardingStep2Request("123", "SIMPLES_NACIONAL");
            try { service.completeStep2(tenantId, req); } catch (BusinessException ignored) {}
            assertThat(tenant.getOnboardingStep()).isEqualTo(OnboardingService.STEP1_DONE);
        }

        @Test
        @DisplayName("deve avançar para STEP2_DONE com CNPJ válido")
        void shouldAdvanceWithValidCnpj() {
            tenant.setOnboardingStep(OnboardingService.STEP1_DONE);
            var req = new OnboardingStep2Request("11222333444455", "SIMPLES_NACIONAL");
            var result = service.completeStep2(tenantId, req);
            assertThat(result.currentStep()).isEqualTo(OnboardingService.STEP2_DONE);
            assertThat(result.progressPercent()).isEqualTo(50);
        }
    }

    @Nested
    @DisplayName("Complete — Finalizar Onboarding")
    class Complete {
        @Test
        @DisplayName("deve transitar tenant para ACTIVE — RN14-04")
        void shouldActivateTenant() {
            tenant.setOnboardingStep(OnboardingService.STEP3_DONE);
            var result = service.complete(tenantId);
            assertThat(result.currentStep()).isEqualTo("COMPLETED");
            assertThat(result.tenantStatus()).isEqualTo("ACTIVE");
            verify(tenantRepo).update(any());
        }

        @Test
        @DisplayName("deve rejeitar se tenant já está ACTIVE — EC-5")
        void shouldRejectAlreadyActive() {
            tenant.setOnboardingStep(OnboardingService.STEP3_DONE);
            tenant.setStatus(TenantStatus.ACTIVE);
            assertThatThrownBy(() -> service.complete(tenantId))
                .isInstanceOf(BusinessException.class)
                .hasMessageContaining("já concluiu o onboarding");
        }

        @Test
        @DisplayName("deve rejeitar se status não é PENDING_ONBOARDING — EC-6")
        void shouldRejectWrongStatus() {
            tenant.setOnboardingStep(OnboardingService.STEP3_DONE);
            tenant.setStatus(TenantStatus.SUSPENDED);
            assertThatThrownBy(() -> service.complete(tenantId))
                .isInstanceOf(BusinessException.class)
                .hasMessageContaining("PENDING_ONBOARDING");
        }
    }

    @Nested
    @DisplayName("Status")
    class Status {
        @Test
        @DisplayName("deve retornar progresso 0% para NOT_STARTED")
        void shouldReturnZeroProgress() {
            var result = service.getStatus(tenantId);
            assertThat(result.currentStep()).isEqualTo(OnboardingService.NOT_STARTED);
            assertThat(result.progressPercent()).isZero();
        }

        @Test
        @DisplayName("deve retornar 100% para COMPLETED")
        void shouldReturnFullProgress() {
            tenant.setOnboardingStep(OnboardingService.COMPLETED);
            when(tenantRepo.findById(tenantId)).thenReturn(Optional.of(tenant));
            var result = service.getStatus(tenantId);
            assertThat(result.progressPercent()).isEqualTo(100);
        }
    }
}
