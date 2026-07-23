package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.NotBlank;

public record OnboardingStep1Request(
    @NotBlank String nameCorporate,
    String nameFantasy,
    @NotBlank String segment
) {}
