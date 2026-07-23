package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.NotBlank;

public record OnboardingStep2Request(
    @NotBlank String cnpj,
    @NotBlank String taxRegime
) {}
