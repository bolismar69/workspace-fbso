package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.response.DashboardClientResponse;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.DashboardClientService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.UUID;

@RestController
@RequestMapping("/api/v1/dashboard/client")
public class DashboardClientController {

    private final DashboardClientService dashboardClientService;

    public DashboardClientController(DashboardClientService dashboardClientService) {
        this.dashboardClientService = dashboardClientService;
    }

    @GetMapping("/summary")
    public ResponseEntity<DashboardClientResponse> getSummary(@RequestParam(required = false) String moduleId) {
        UUID tenantId = TenantContext.getTenantId();
        return ResponseEntity.ok(dashboardClientService.getSummary(tenantId));
    }

    @GetMapping("/notifications")
    public ResponseEntity<java.util.List<?>> getNotifications() {
        return ResponseEntity.ok(java.util.List.of());
    }
}
