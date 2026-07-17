package com.fbso.platform.admin.entity;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa um recurso + ação no sistema RBAC.
 *
 * <p>Mapeia a tabela {@code fbso_platform.resource_action} (V001).
 * Tabela GLOBAL — sem tenant_id (matriz RBAC compartilhada).</p>
 *
 * <h3>Exemplos</h3>
 * <ul>
 *   <li>DASHBOARD + view</li>
 *   <li>TENANT + create</li>
 *   <li>PRODUCT_SERVICE + edit</li>
 * </ul>
 *
 * @see RoleResource
 */
public class ResourceAction {

    private UUID id;
    private String resourceName;
    private String action;

    public ResourceAction() {
    }

    public UUID getId() { return id; }
    public void setId(UUID id) { this.id = id; }

    public String getResourceName() { return resourceName; }
    public void setResourceName(String resourceName) { this.resourceName = resourceName; }

    public String getAction() { return action; }
    public void setAction(String action) { this.action = action; }

    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("resource_name", resourceName);
        columns.put("action", action);
        return columns;
    }

    @Override
    public String toString() {
        return "ResourceAction{id=" + id + ", resource='" + resourceName + "', action='" + action + "'}";
    }
}
