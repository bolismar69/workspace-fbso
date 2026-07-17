package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonInclude;

import java.util.List;
import java.util.UUID;

/**
 * Alertas do dashboard administrativo.
 * <p>
 * F01-03: Cards coloridos — WARNING (atenção) e CRITICAL (ação urgente).
 * Alertas: onboarding >48h (WARNING), assinatura suspensa (CRITICAL).
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
public record AlertResponse(
        List<Alert> alerts
) {
    public enum AlertType { WARNING, CRITICAL }

    public record Alert(
            AlertType type,
            String message,
            UUID entityId,
            String entityType
    ) {}

    public static AlertResponse of(List<Alert> alerts) {
        return new AlertResponse(alerts);
    }
}
