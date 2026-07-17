package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.ChangePlanRequest;
import com.fbso.platform.admin.dto.request.SubscriptionCreateRequest;
import com.fbso.platform.admin.dto.response.SubscriptionResponse;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.SubscriptionService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.net.URI;
import java.util.List;
import java.util.UUID;

@RestController
@Tag(name = "Assinaturas", description = "Gestão de assinaturas Tenant × Plano")
public class SubscriptionController {

    private static final Logger log = LoggerFactory.getLogger(SubscriptionController.class);

    private final SubscriptionService subService;

    public SubscriptionController(SubscriptionService subService) {
        this.subService = subService;
    }

    @GetMapping("/api/v1/tenants/{tid}/subscriptions")
    @Operation(summary = "Histórico de assinaturas do tenant")
    @RequiresPermission(resource = "SUBSCRIPTION", action = "view")
    public ResponseEntity<List<SubscriptionResponse>> listByTenant(@PathVariable UUID tid) {
        return ResponseEntity.ok(subService.findByTenant(tid));
    }

    @PostMapping("/api/v1/tenants/{tid}/subscriptions")
    @Operation(summary = "Criar assinatura para tenant")
    @RequiresPermission(resource = "SUBSCRIPTION", action = "create")
    public ResponseEntity<SubscriptionResponse> create(@PathVariable UUID tid,
                                                        @Valid @RequestBody SubscriptionCreateRequest req) {
        SubscriptionResponse response = subService.create(tid, req.planId());
        return ResponseEntity.created(URI.create("/api/v1/tenants/" + tid + "/subscriptions/" + response.id()))
                .body(response);
    }

    @PostMapping("/api/v1/subscriptions/{id}/change-plan")
    @Operation(summary = "Mudar plano da assinatura (atômico, sem gap)")
    @RequiresPermission(resource = "SUBSCRIPTION", action = "edit")
    public ResponseEntity<SubscriptionResponse> changePlan(@PathVariable UUID id,
                                                            @Valid @RequestBody ChangePlanRequest req) {
        return ResponseEntity.ok(subService.changePlan(id, req.newPlanId()));
    }

    @PostMapping("/api/v1/subscriptions/{id}/suspend")
    @Operation(summary = "Suspender assinatura")
    @RequiresPermission(resource = "SUBSCRIPTION", action = "edit")
    public ResponseEntity<SubscriptionResponse> suspend(@PathVariable UUID id) {
        return ResponseEntity.ok(subService.suspend(id));
    }
}
