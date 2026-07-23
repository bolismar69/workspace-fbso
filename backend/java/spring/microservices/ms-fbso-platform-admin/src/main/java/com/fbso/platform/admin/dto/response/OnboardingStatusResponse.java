package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonInclude;

import java.util.List;

@JsonInclude(JsonInclude.Include.NON_NULL)
public record OnboardingStatusResponse(
    String currentStep,
    List<String> completedSteps,
    int progressPercent,
    String tenantStatus,
    String message
) {
    public static OnboardingStatusResponse of(String currentStep, List<String> completedSteps,
                                               int progressPercent, String tenantStatus) {
        return new OnboardingStatusResponse(currentStep, completedSteps, progressPercent, tenantStatus, null);
    }

    public static OnboardingStatusResponse completed() {
        return new OnboardingStatusResponse("COMPLETED", List.of("STEP1", "STEP2", "STEP3"), 100, "ACTIVE",
            "Onboarding concluído com sucesso!");
    }
}
