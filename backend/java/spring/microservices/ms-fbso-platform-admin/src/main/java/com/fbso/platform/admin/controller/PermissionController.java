package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.PermissionUpdateRequest;
import com.fbso.platform.admin.dto.response.PermissionResponse;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.PermissionService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

/**
 * Endpoints de gestão de Permissões (F03-02, F03-03).
 *
 * <p>2 endpoints REST:
 * <ul>
 *   <li>GET /api/v1/users/{userId}/permissions — lista permissões do usuário</li>
 *   <li>PUT /api/v1/users/{userId}/permissions — atualiza vínculos (batch)</li>
 * </ul>
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN11-01: Usuário requer ≥1 BU — validado via {@code PermissionUpdateRequest.permissions}</li>
 *   <li>RN11-03: Efeito imediato — batch update em uma transação</li>
 * </ul>
 */
@RestController
@RequestMapping("/api/v1/users/{userId}/permissions")
@Tag(name = "Permissions", description = "Gestão de permissões RBAC — Admin Tenant")
public class PermissionController {

    private static final Logger log = LoggerFactory.getLogger(PermissionController.class);

    private final PermissionService permissionService;

    public PermissionController(PermissionService permissionService) {
        this.permissionService = permissionService;
    }

    @GetMapping
    @Operation(summary = "Listar permissões do usuário")
    @RequiresPermission(resource = "PERMISSION", action = "view")
    public ResponseEntity<List<PermissionResponse>> list(@PathVariable UUID userId) {
        log.debug("Listando permissões do usuário {}", userId);
        List<PermissionResponse> permissions = permissionService.getUserPermissions(userId);
        return ResponseEntity.ok(permissions);
    }

    @PutMapping
    @Operation(summary = "Atualizar permissões do usuário (batch)")
    @RequiresPermission(resource = "PERMISSION", action = "edit")
    public ResponseEntity<List<PermissionResponse>> update(
            @PathVariable UUID userId,
            @Valid @RequestBody PermissionUpdateRequest request) {
        log.info("Atualizando permissões do usuário {}: {} assignments", userId,
                request.permissions().size());
        List<PermissionResponse> updated = permissionService.updateUserPermissions(userId, request);
        return ResponseEntity.ok(updated);
    }
}
