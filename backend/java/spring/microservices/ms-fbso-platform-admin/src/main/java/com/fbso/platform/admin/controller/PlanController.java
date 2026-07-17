package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.PlanCreateRequest;
import com.fbso.platform.admin.dto.request.PlanUpdateRequest;
import com.fbso.platform.admin.dto.response.PlanResponse;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.PlanService;
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
@RequestMapping("/api/v1/plans")
@Tag(name = "Planos", description = "Gestão de planos comerciais — Admin FBSO")
public class PlanController {

    private static final Logger log = LoggerFactory.getLogger(PlanController.class);

    private final PlanService planService;

    public PlanController(PlanService planService) {
        this.planService = planService;
    }

    @GetMapping
    @Operation(summary = "Listar planos ativos (para assinatura)")
    @RequiresPermission(resource = "PLAN", action = "view")
    public ResponseEntity<List<PlanResponse>> listActive() {
        return ResponseEntity.ok(planService.listAll());
    }

    @GetMapping("/admin")
    @Operation(summary = "Listar todos os planos (admin)")
    @RequiresPermission(resource = "PLAN", action = "view")
    public ResponseEntity<List<PlanResponse>> listAdmin() {
        return ResponseEntity.ok(planService.listAdmin());
    }

    @GetMapping("/{id}")
    @Operation(summary = "Buscar plano por ID")
    @RequiresPermission(resource = "PLAN", action = "view")
    public ResponseEntity<PlanResponse> getById(@PathVariable UUID id) {
        return ResponseEntity.ok(planService.getPlan(id));
    }

    @PostMapping
    @Operation(summary = "Criar novo plano")
    @RequiresPermission(resource = "PLAN", action = "create")
    public ResponseEntity<PlanResponse> create(@Valid @RequestBody PlanCreateRequest req) {
        PlanResponse response = planService.create(req);
        return ResponseEntity.created(URI.create("/api/v1/plans/" + response.id())).body(response);
    }

    @PatchMapping("/{id}")
    @Operation(summary = "Editar plano (versiona se preço alterado)")
    @RequiresPermission(resource = "PLAN", action = "edit")
    public ResponseEntity<PlanResponse> update(@PathVariable UUID id,
                                                @Valid @RequestBody PlanUpdateRequest req) {
        return ResponseEntity.ok(planService.update(id, req));
    }

    @PostMapping("/{id}/deactivate")
    @Operation(summary = "Desativar plano (bloqueado se com assinantes)")
    @RequiresPermission(resource = "PLAN", action = "deactivate")
    public ResponseEntity<PlanResponse> deactivate(@PathVariable UUID id) {
        return ResponseEntity.ok(planService.deactivate(id));
    }
}
