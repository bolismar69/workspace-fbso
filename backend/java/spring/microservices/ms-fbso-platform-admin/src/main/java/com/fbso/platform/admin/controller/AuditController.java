package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.response.AuditEntryResponse;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.AuditService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

/**
 * Endpoint de consulta de auditoria (F02-05).
 *
 * <p>GET /api/v1/audit — registros imutáveis, filtros por período/ação/entidade.
 * Acesso: Admin FBSO (tudo) ou Auditor (leitura).</p>
 */
@RestController
@RequestMapping("/api/v1/audit")
@Tag(name = "Auditoria", description = "Consulta de trilha de auditoria — Admin/Auditor")
public class AuditController {

    private final AuditService auditService;

    public AuditController(AuditService auditService) {
        this.auditService = auditService;
    }

    @GetMapping
    @Operation(summary = "Consultar registros de auditoria com filtros")
    @RequiresPermission(resource = "AUDIT", action = "view")
    public ResponseEntity<List<AuditEntryResponse>> search(
            @RequestParam(required = false) String startDate,
            @RequestParam(required = false) String endDate,
            @RequestParam(required = false) String action,
            @RequestParam(required = false) String entityType,
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "25") int size) {

        List<AuditEntryResponse> results = auditService.search(
                startDate, endDate, action, entityType, page, size);
        return ResponseEntity.ok(results);
    }
}
