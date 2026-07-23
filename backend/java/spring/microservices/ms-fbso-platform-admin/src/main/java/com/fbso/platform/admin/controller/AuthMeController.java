package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.response.AuthMeResponse;
import com.fbso.platform.admin.security.TenantContext;
import org.springframework.http.ResponseEntity;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

/**
 * Endpoint /auth/me — retorna dados do usuário logado extraídos do JWT (stateless).
 * Sem @RequiresPermission — autenticação já validada no JwtAuthenticationFilter.
 */
@RestController
@RequestMapping("/api/v1/auth")
public class AuthMeController {

    private final JdbcTemplate jdbc;

    public AuthMeController(JdbcTemplate jdbc) {
        this.jdbc = jdbc;
    }

    @GetMapping("/me")
    public ResponseEntity<AuthMeResponse> me() {
        UUID userId = TenantContext.getUserId();
        UUID tenantId = TenantContext.getTenantId();
        List<String> roles = TenantContext.getRoles();
        List<UUID> buIds = TenantContext.getBusinessUnitIds();
        List<String> modules = TenantContext.getModules();

        // RN16-01, RN16-02: Placeholder "FBSO Platform" se modules vazio
        if (modules == null || modules.isEmpty()) {
            modules = List.of("FBSO Platform");
        }

        // Nome e email podem vir do JWT; Fase 0 usa defaults do TenantContext
        String role = (roles != null && !roles.isEmpty()) ? roles.get(0) : "UNKNOWN";
        String onboardingStatus = getOnboardingStatus(tenantId);

        return ResponseEntity.ok(new AuthMeResponse(
            userId, "Usuário FBSO", "usuario@fbso.org",
            role, buIds != null ? buIds : List.of(), modules,
            tenantId, onboardingStatus
        ));
    }

    private String getOnboardingStatus(UUID tenantId) {
        try {
            String step = jdbc.queryForObject(
                "SELECT onboarding_step FROM fbso_platform.tenant WHERE id = ? AND deleted_dt IS NULL",
                String.class, tenantId);
            return step != null ? step : "NOT_STARTED";
        } catch (Exception e) {
            return "NOT_STARTED";
        }
    }
}
