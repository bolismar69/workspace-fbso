package com.fbso.platform.admin.utils;

import org.springframework.security.oauth2.jwt.Jwt;

import java.util.List;
import java.util.UUID;

/**
 * Utilitário para extração de claims customizadas do JWT (Keycloak).
 * <p>
 * Claims esperadas no token:
 * <ul>
 *   <li>{@code tenant_id} — UUID do tenant (obrigatório)</li>
 *   <li>{@code user_id} — UUID do usuário (sub claim padrão)</li>
 *   <li>{@code roles} — lista de papéis (realm_access.roles)</li>
 *   <li>{@code business_unit_ids} — lista de UUIDs das BUs autorizadas</li>
 *   <li>{@code modules} — lista de módulos contratados no plano</li>
 * </ul>
 */
public final class JwtUtils {

    private JwtUtils() {}

    /**
     * Extrai o tenant_id do JWT.
     *
     * @throws IllegalArgumentException se a claim estiver ausente
     */
    public static UUID getTenantId(Jwt jwt) {
        String claim = jwt.getClaimAsString("tenant_id");
        if (claim == null || claim.isBlank()) {
            throw new IllegalArgumentException("JWT não contém tenant_id");
        }
        return UUID.fromString(claim);
    }

    /**
     * Extrai o user_id do JWT (claim {@code sub}).
     */
    public static UUID getUserId(Jwt jwt) {
        String subject = jwt.getSubject();
        if (subject == null || subject.isBlank()) {
            throw new IllegalArgumentException("JWT não contém sub (user_id)");
        }
        return UUID.fromString(subject);
    }

    /**
     * Extrai os roles do JWT.
     * <p>
     * Busca primeiro em {@code realm_access.roles}, depois em {@code roles} direto.
     */
    @SuppressWarnings("unchecked")
    public static List<String> getRoles(Jwt jwt) {
        // Tenta realm_access.roles (padrão Keycloak)
        var realmAccess = jwt.getClaimAsMap("realm_access");
        if (realmAccess != null && realmAccess.get("roles") instanceof List<?> roles) {
            return (List<String>) roles;
        }
        // Fallback: claim "roles" direta
        var directRoles = jwt.getClaimAsStringList("roles");
        if (directRoles != null) {
            return directRoles;
        }
        return List.of();
    }

    /**
     * Extrai os business_unit_ids do JWT.
     */
    public static List<UUID> getBusinessUnitIds(Jwt jwt) {
        List<String> raw = jwt.getClaimAsStringList("business_unit_ids");
        if (raw == null) {
            return List.of();
        }
        return raw.stream().map(UUID::fromString).toList();
    }

    /**
     * Extrai os módulos contratados do JWT.
     */
    public static List<String> getModules(Jwt jwt) {
        List<String> modules = jwt.getClaimAsStringList("modules");
        return modules != null ? modules : List.of();
    }
}
