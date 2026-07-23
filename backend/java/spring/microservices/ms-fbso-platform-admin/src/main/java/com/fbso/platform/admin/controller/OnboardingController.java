package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.OnboardingStep1Request;
import com.fbso.platform.admin.dto.request.OnboardingStep2Request;
import com.fbso.platform.admin.dto.request.OnboardingStep3Request;
import com.fbso.platform.admin.dto.response.OnboardingStatusResponse;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.OnboardingService;
import jakarta.validation.Valid;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.UUID;

@RestController
@RequestMapping("/api/v1/onboarding")
public class OnboardingController {

    private final OnboardingService onboardingService;

    public OnboardingController(OnboardingService onboardingService) {
        this.onboardingService = onboardingService;
    }

    @GetMapping("/status")
    public ResponseEntity<OnboardingStatusResponse> getStatus() {
        UUID tenantId = TenantContext.getTenantId();
        return ResponseEntity.ok(onboardingService.getStatus(tenantId));
    }

    @PatchMapping("/step-1")
    public ResponseEntity<OnboardingStatusResponse> completeStep1(@Valid @RequestBody OnboardingStep1Request request) {
        UUID tenantId = TenantContext.getTenantId();
        return ResponseEntity.ok(onboardingService.completeStep1(tenantId, request));
    }

    @PostMapping("/step-2")
    public ResponseEntity<OnboardingStatusResponse> completeStep2(@Valid @RequestBody OnboardingStep2Request request) {
        UUID tenantId = TenantContext.getTenantId();
        return ResponseEntity.status(201).body(onboardingService.completeStep2(tenantId, request));
    }

    @PostMapping("/complete")
    public ResponseEntity<OnboardingStatusResponse> complete() {
        UUID tenantId = TenantContext.getTenantId();
        return ResponseEntity.ok(onboardingService.complete(tenantId));
    }
}
