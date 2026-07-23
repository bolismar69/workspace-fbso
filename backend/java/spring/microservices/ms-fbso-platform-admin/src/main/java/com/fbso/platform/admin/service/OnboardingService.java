package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.OnboardingStep1Request;
import com.fbso.platform.admin.dto.request.OnboardingStep2Request;
import com.fbso.platform.admin.dto.request.OnboardingStep3Request;
import com.fbso.platform.admin.utils.CnpjValidator;
import com.fbso.platform.admin.dto.response.OnboardingStatusResponse;
import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.repository.TenantRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

/**
 * Serviço de Onboarding — gerencia o fluxo de 4 passos do primeiro acesso.
 *
 * <p>Máquina de Estados (ARCHITECTURE.md §8.2): NOT_STARTED → STEP1_DONE → STEP2_DONE → STEP3_DONE → COMPLETED.</p>
 *
 * <p>RNs: RN14-01 (ordem), RN14-02 (primeira BU=Matriz), RN14-03 (todos passos OK), RN14-04 (Tenant→ACTIVE).</p>
 */
@Service
public class OnboardingService {

    private static final Logger log = LoggerFactory.getLogger(OnboardingService.class);

    public static final String NOT_STARTED = "NOT_STARTED";
    public static final String STEP1_DONE = "STEP1_DONE";
    public static final String STEP2_DONE = "STEP2_DONE";
    public static final String STEP3_DONE = "STEP3_DONE";
    public static final String COMPLETED = "COMPLETED";

    private final TenantRepository tenantRepo;

    public OnboardingService(TenantRepository tenantRepo) {
        this.tenantRepo = tenantRepo;
    }

    public OnboardingStatusResponse getStatus(UUID tenantId) {
        return buildStatusResponse(findTenant(tenantId));
    }

    @Transactional
    public OnboardingStatusResponse completeStep1(UUID tenantId, OnboardingStep1Request request) {
        Tenant tenant = findTenant(tenantId);
        requireStep(tenant, NOT_STARTED, "Step 1");
        tenant.setNameCorporate(request.nameCorporate());
        if (request.nameFantasy() != null) tenant.setNameFantasy(request.nameFantasy());
        tenant.setOnboardingStep(STEP1_DONE);
        tenantRepo.update(tenant);
        log.info("Onboarding Step-1 concluído: tenantId={}", tenantId);
        return buildStatusResponse(tenant);
    }

    @Transactional
    public OnboardingStatusResponse completeStep2(UUID tenantId, OnboardingStep2Request request) {
        Tenant tenant = findTenant(tenantId);
        requireStep(tenant, STEP1_DONE, "Step 2");
        if (!CnpjValidator.isValid(request.cnpj())) {
            throw new BusinessException("CNPJ_INVALIDO", "CNPJ inválido. Verifique e tente novamente.");
        }
        tenant.setOnboardingStep(STEP2_DONE);
        tenantRepo.update(tenant);
        log.info("Onboarding Step-2 concluído: tenantId={}", tenantId);
        return buildStatusResponse(tenant);
    }

    @Transactional
    public OnboardingStatusResponse completeStep3(UUID tenantId, OnboardingStep3Request request) {
        Tenant tenant = findTenant(tenantId);
        requireStep(tenant, STEP2_DONE, "Step 3");
        tenant.setOnboardingStep(STEP3_DONE);
        tenantRepo.update(tenant);
        log.info("Onboarding Step-3 concluído: tenantId={}", tenantId);
        return buildStatusResponse(tenant);
    }

    @Transactional
    public OnboardingStatusResponse complete(UUID tenantId) {
        Tenant tenant = findTenant(tenantId);
        requireStep(tenant, STEP3_DONE, "Complete");

        if (tenant.getStatus() == TenantStatus.ACTIVE) {
            throw new BusinessException("ONBOARDING_JA_CONCLUIDO", "Este tenant já concluiu o onboarding.");
        }
        if (tenant.getStatus() != TenantStatus.PENDING_ONBOARDING) {
            throw new BusinessException("STATUS_INVALIDO",
                "Tenant deve estar em PENDING_ONBOARDING. Status atual: " + tenant.getStatus());
        }

        tenant.setStatus(TenantStatus.ACTIVE);
        tenant.setOnboardingStep(COMPLETED);
        tenantRepo.update(tenant);
        log.info("Onboarding COMPLETO: tenantId={} → ACTIVE", tenantId);
        return OnboardingStatusResponse.completed();
    }

    // ---- Helpers ----

    private Tenant findTenant(UUID tenantId) {
        return tenantRepo.findById(tenantId)
            .orElseThrow(() -> new BusinessException("TENANT_NAO_ENCONTRADO",
                "Tenant id=" + tenantId + " não encontrado."));
    }

    private void requireStep(Tenant tenant, String expectedStep, String stepName) {
        String current = tenant.getOnboardingStep() != null ? tenant.getOnboardingStep() : NOT_STARTED;
        if (!current.equals(expectedStep)) {
            throw new BusinessException("ORDEM_ONBOARDING",
                String.format("Conclua o passo anterior. Passo atual: %s, esperado: %s, tentou: %s.",
                    current, expectedStep, stepName));
        }
    }

    private OnboardingStatusResponse buildStatusResponse(Tenant tenant) {
        String step = tenant.getOnboardingStep() != null ? tenant.getOnboardingStep() : NOT_STARTED;
        List<String> completed = new ArrayList<>();
        if (List.of(STEP1_DONE, STEP2_DONE, STEP3_DONE, COMPLETED).contains(step)) completed.add("STEP1");
        if (List.of(STEP2_DONE, STEP3_DONE, COMPLETED).contains(step)) completed.add("STEP2");
        if (List.of(STEP3_DONE, COMPLETED).contains(step)) completed.add("STEP3");
        int progress = switch (step) {
            case NOT_STARTED -> 0; case STEP1_DONE -> 25; case STEP2_DONE -> 50;
            case STEP3_DONE -> 75; case COMPLETED -> 100; default -> 0;
        };
        return OnboardingStatusResponse.of(step, completed, progress,
            tenant.getStatus() != null ? tenant.getStatus().name() : null);
    }
}
