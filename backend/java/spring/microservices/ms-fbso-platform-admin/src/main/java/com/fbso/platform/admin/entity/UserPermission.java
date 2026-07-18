package com.fbso.platform.admin.entity;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Vínculo Usuário × Unidade de Negócio × Role (F03-02, F03-03).
 *
 * <p>Mapeia a tabela {@code fbso_platform.user_permission} (V001).
 * Tabela MULTI-TENANT indireta — o tenant da BU é verificado via FK.</p>
 *
 * <h3>Constraints</h3>
 * <ul>
 *   <li>UNIQUE(user_id, business_unit_id) — um usuário não pode ter 2 roles na mesma BU</li>
 *   <li>FK user_id → "user"(id)</li>
 *   <li>FK business_unit_id → business_unit(id) — V006</li>
 * </ul>
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN11-01: Usuário requer ≥1 BU vinculada</li>
 *   <li>RN11-02: Usuário requer ≥1 Módulo</li>
 *   <li>RN11-03: Efeito imediato — sem cache TTL</li>
 * </ul>
 */
public class UserPermission {

    private UUID id;
    private UUID userId;
    private UUID businessUnitId;
    private String role;

    public UserPermission() {
    }

    public UUID getId() { return id; }
    public void setId(UUID id) { this.id = id; }

    public UUID getUserId() { return userId; }
    public void setUserId(UUID userId) { this.userId = userId; }

    public UUID getBusinessUnitId() { return businessUnitId; }
    public void setBusinessUnitId(UUID businessUnitId) { this.businessUnitId = businessUnitId; }

    public String getRole() { return role; }
    public void setRole(String role) { this.role = role; }

    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("user_id", userId);
        columns.put("business_unit_id", businessUnitId);
        columns.put("role", role);
        return columns;
    }

    @Override
    public String toString() {
        return "UserPermission{id=" + id + ", userId=" + userId
                + ", buId=" + businessUnitId + ", role='" + role + "'}";
    }
}
