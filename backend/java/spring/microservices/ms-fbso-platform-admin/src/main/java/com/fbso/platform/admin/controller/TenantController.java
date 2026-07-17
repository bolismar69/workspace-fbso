package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.SuspendTenantRequest;
import com.fbso.platform.admin.dto.request.TenantCreateRequest;
import com.fbso.platform.admin.dto.request.TenantUpdateRequest;
import com.fbso.platform.admin.dto.response.TenantResponse;
import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.repository.TenantRepository;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.TenantService;
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

/**
 * Endpoints de gestão de Tenants (F02-01, F02-02).
 *
 * <p>7 endpoints REST:
 * <ul>
 *   <li>GET /api/v1/tenants — lista paginada com filtros</li>
 *   <li>GET /api/v1/tenants/{id} — busca por ID</li>
 *   <li>POST /api/v1/tenants — criar (status PENDING_ONBOARDING)</li>
 *   <li>PATCH /api/v1/tenants/{id} — editar dados cadastrais</li>
 *   <li>POST /api/v1/tenants/{id}/suspend — suspender (exige motivo)</li>
 *   <li>POST /api/v1/tenants/{id}/reactivate — reativar</li>
 *   <li>POST /api/v1/tenants/{id}/resend-invite — reenviar convite</li>
 * </ul>
 */
@RestController
@RequestMapping("/api/v1/tenants")
@Tag(name = "Tenants", description = "Gestão de contas corporativas — Admin FBSO")
public class TenantController {

    private static final Logger log = LoggerFactory.getLogger(TenantController.class);

    private final TenantService tenantService;
    private final TenantRepository tenantRepo;

    public TenantController(TenantService tenantService, TenantRepository tenantRepo) {
        this.tenantService = tenantService;
        this.tenantRepo = tenantRepo;
    }

    // ---- List (F01-02 herdado) ----

    @GetMapping
    @Operation(summary = "Listar tenants paginado com filtros")
    @RequiresPermission(resource = "TENANT", action = "view")
    public ResponseEntity<List<TenantResponse>> list(
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "25") int size,
            @RequestParam(required = false) String status,
            @RequestParam(required = false) String plan,
            @RequestParam(required = false) String search) {

        List<Tenant> tenants = tenantRepo.findAllPaginated(page, size, status, plan, search);
        List<TenantResponse> response = tenants.stream().map(TenantResponse::from).toList();
        return ResponseEntity.ok(response);
    }

    // ---- Get by ID ----

    @GetMapping("/{id}")
    @Operation(summary = "Buscar tenant por ID")
    @RequiresPermission(resource = "TENANT", action = "view")
    public ResponseEntity<TenantResponse> getById(@PathVariable UUID id) {
        Tenant tenant = tenantService.getById(id);
        return ResponseEntity.ok(TenantResponse.from(tenant));
    }

    // ---- Create (F02-01) ----

    @PostMapping
    @Operation(summary = "Criar novo tenant (status PENDING_ONBOARDING)")
    @RequiresPermission(resource = "TENANT", action = "create")
    public ResponseEntity<TenantResponse> create(@Valid @RequestBody TenantCreateRequest req) {
        log.info("Criando tenant: nameCorporate={}", req.nameCorporate());
        TenantResponse response = tenantService.create(req);
        return ResponseEntity.created(URI.create("/api/v1/tenants/" + response.id()))
                .body(response);
    }

    // ---- Update (F02-01) ----

    @PatchMapping("/{id}")
    @Operation(summary = "Editar dados cadastrais do tenant")
    @RequiresPermission(resource = "TENANT", action = "edit")
    public ResponseEntity<TenantResponse> update(@PathVariable UUID id,
                                                  @Valid @RequestBody TenantUpdateRequest req) {
        log.info("Atualizando tenant: id={}", id);
        TenantResponse response = tenantService.update(id, req);
        return ResponseEntity.ok(response);
    }

    // ---- Suspend (F02-02, RN05-02) ----

    @PostMapping("/{id}/suspend")
    @Operation(summary = "Suspender tenant (exige motivo)")
    @RequiresPermission(resource = "TENANT", action = "suspend")
    public ResponseEntity<TenantResponse> suspend(@PathVariable UUID id,
                                                   @Valid @RequestBody SuspendTenantRequest req) {
        log.warn("Suspendendo tenant: id={}, reason={}", id, req.reason());
        TenantResponse response = tenantService.suspend(id, req.reason());
        return ResponseEntity.ok(response);
    }

    // ---- Reactivate (F02-02, RN05-03) ----

    @PostMapping("/{id}/reactivate")
    @Operation(summary = "Reativar tenant suspenso")
    @RequiresPermission(resource = "TENANT", action = "reactivate")
    public ResponseEntity<TenantResponse> reactivate(@PathVariable UUID id) {
        log.info("Reativando tenant: id={}", id);
        TenantResponse response = tenantService.reactivate(id);
        return ResponseEntity.ok(response);
    }

    // ---- Resend Invite ----

    @PostMapping("/{id}/resend-invite")
    @Operation(summary = "Reenviar convite de ativação")
    @RequiresPermission(resource = "TENANT", action = "edit")
    public ResponseEntity<Void> resendInvite(@PathVariable UUID id) {
        log.info("Reenviando convite para tenant: id={}", id);
        // T-028: Email integration
        return ResponseEntity.accepted().build();
    }
}
