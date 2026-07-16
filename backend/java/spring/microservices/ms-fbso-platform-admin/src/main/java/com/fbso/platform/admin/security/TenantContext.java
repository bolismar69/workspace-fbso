package com.fbso.platform.admin.security;

import java.util.List;
import java.util.UUID;

/**
 * Contexto da requisição — armazena claims do JWT em ThreadLocal.
 * <p>
 * Populado pelo {@code JwtAuthenticationFilter} no início de cada requisição.
 * Limpo no {@code finally} do filter para evitar vazamento entre threads.
 * <p>
 * Métodos estáticos delegam para um {@code ThreadLocal<TenantContextData>} interno.
 */
public final class TenantContext {

    private static final ThreadLocal<TenantContextData> CONTEXT = new ThreadLocal<>();

    private TenantContext() {}

    /**
     * Seta o contexto da requisição atual.
     */
    public static void set(UUID tenantId, UUID userId, List<String> roles,
                           List<UUID> businessUnitIds, List<String> modules) {
        CONTEXT.set(new TenantContextData(tenantId, userId, roles, businessUnitIds, modules));
    }

    /**
     * @return tenant_id extraído do JWT
     * @throws IllegalStateException se o contexto não foi inicializado
     */
    public static UUID getTenantId() {
        TenantContextData data = CONTEXT.get();
        if (data == null) {
            throw new IllegalStateException(
                "TenantContext não inicializado — JwtAuthenticationFilter não foi executado?");
        }
        return data.tenantId;
    }

    /**
     * @return user_id extraído do JWT
     */
    public static UUID getUserId() {
        TenantContextData data = CONTEXT.get();
        return data != null ? data.userId : null;
    }

    /**
     * @return roles extraídos do JWT
     */
    public static List<String> getRoles() {
        TenantContextData data = CONTEXT.get();
        return data != null ? data.roles : List.of();
    }

    /**
     * @return business_unit_ids extraídos do JWT
     */
    public static List<UUID> getBusinessUnitIds() {
        TenantContextData data = CONTEXT.get();
        return data != null ? data.businessUnitIds : List.of();
    }

    /**
     * @return modules extraídos do JWT
     */
    public static List<String> getModules() {
        TenantContextData data = CONTEXT.get();
        return data != null ? data.modules : List.of();
    }

    /**
     * Retorna o tenant_id como String sem lançar exceção.
     * Usado pelo {@code TenantAwareDataSource} para configurar a sessão PostgreSQL.
     *
     * @return tenant_id como string, ou {@code null} se o contexto não foi inicializado
     *         (ex: health check, Admin FBSO global)
     */
    public static String getTenantIdQuietly() {
        TenantContextData data = CONTEXT.get();
        return data != null && data.tenantId != null ? data.tenantId.toString() : null;
    }

    /**
     * Limpa o contexto. DEVE ser chamado no finally do filter.
     */
    public static void clear() {
        CONTEXT.remove();
    }

    // ---- Data Holder ----

    private record TenantContextData(
            UUID tenantId,
            UUID userId,
            List<String> roles,
            List<UUID> businessUnitIds,
            List<String> modules
    ) {}
}
