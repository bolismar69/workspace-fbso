package com.fbso.platform.admin.entity;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa o vínculo entre um papel (Role) e um ResourceAction na matriz RBAC.
 *
 * <p>Mapeia a tabela {@code fbso_platform.role_resource} (V001).
 * Tabela GLOBAL — sem tenant_id (a matriz RBAC é compartilhada entre tenants).</p>
 *
 * <h3>Exemplo</h3>
 * <pre>
 *   ADMIN_TENANT → DASHBOARD:view
 *   ADMIN_TENANT → TENANT:create
 *   MANAGER_BU   → BUSINESS_UNIT:edit
 * </pre>
 *
 * @see ResourceAction
 * @see com.fbso.platform.admin.enums.Role
 */
public class RoleResource {

    private UUID id;
    private String role;
    private UUID resourceActionId;

    public RoleResource() {
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getRole() {
        return role;
    }

    public void setRole(String role) {
        this.role = role;
    }

    public UUID getResourceActionId() {
        return resourceActionId;
    }

    public void setResourceActionId(UUID resourceActionId) {
        this.resourceActionId = resourceActionId;
    }

    /**
     * Mapeia colunas para INSERT/UPDATE via JDBC.
     */
    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("role", role);
        columns.put("resource_action_id", resourceActionId);
        return columns;
    }

    @Override
    public String toString() {
        return "RoleResource{id=" + id + ", role='" + role + "', resourceActionId=" + resourceActionId + '}';
    }
}
