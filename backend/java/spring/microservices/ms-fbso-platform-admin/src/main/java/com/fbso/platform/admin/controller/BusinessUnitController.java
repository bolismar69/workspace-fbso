package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.BusinessUnitCreateRequest;
import com.fbso.platform.admin.dto.request.BusinessUnitUpdateRequest;
import com.fbso.platform.admin.dto.response.BusinessUnitResponse;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.BusinessUnitService;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

/**
 * Controller REST para Unidades de Negócio (F04-05).
 *
 * <h3>RBAC (RN10-01)</h3>
 * <ul>
 *   <li>ADMIN_TENANT: acesso total</li>
 *   <li>MANAGER_BU: create, edit (BU autorizadas)</li>
 *   <li>OPERATOR_BU: view (BU autorizadas)</li>
 *   <li>AUDITOR: sem acesso</li>
 * </ul>
 */
@RestController
@RequestMapping("/api/v1/business-units")
public class BusinessUnitController {

    private final BusinessUnitService buService;

    public BusinessUnitController(BusinessUnitService buService) {
        this.buService = buService;
    }

    @GetMapping
    @RequiresPermission(resource = "BUSINESS_UNIT", action = "view")
    public ResponseEntity<List<BusinessUnitResponse>> list() {
        return ResponseEntity.ok(buService.findAccessible());
    }

    @GetMapping("/tree")
    @RequiresPermission(resource = "BUSINESS_UNIT", action = "view")
    public ResponseEntity<List<BusinessUnitResponse>> tree() {
        return ResponseEntity.ok(buService.findTree());
    }

    @GetMapping("/{id}")
    @RequiresPermission(resource = "BUSINESS_UNIT", action = "view")
    public ResponseEntity<BusinessUnitResponse> getById(@PathVariable UUID id) {
        return ResponseEntity.ok(buService.findById(id));
    }

    @PostMapping
    @RequiresPermission(resource = "BUSINESS_UNIT", action = "create")
    public ResponseEntity<BusinessUnitResponse> create(
            @Valid @RequestBody BusinessUnitCreateRequest req) {
        BusinessUnitResponse response = buService.create(req);
        return ResponseEntity.status(HttpStatus.CREATED).body(response);
    }

    @PatchMapping("/{id}")
    @RequiresPermission(resource = "BUSINESS_UNIT", action = "edit")
    public ResponseEntity<BusinessUnitResponse> update(
            @PathVariable UUID id,
            @Valid @RequestBody BusinessUnitUpdateRequest req) {
        return ResponseEntity.ok(buService.update(id, req));
    }

    @PostMapping("/{id}/deactivate")
    @RequiresPermission(resource = "BUSINESS_UNIT", action = "delete")
    public ResponseEntity<BusinessUnitResponse> deactivate(@PathVariable UUID id) {
        return ResponseEntity.ok(buService.deactivate(id));
    }
}
