package com.fbso.platform.admin.entity;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.enums.UserStatus;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa um usuário do sistema vinculado a um tenant.
 *
 * <p>Mapeia a tabela {@code fbso_platform."user"} (V001).
 * RLS ativo (V003) — filtrado por {@code tenant_id}.</p>
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN09-01: Convite expira em 7 dias</li>
 *   <li>RN09-02: Email único por tenant ativo (índice parcial V002)</li>
 *   <li>RN09-03: Admin não desativa a si mesmo (validado no Service)</li>
 * </ul>
 *
 * @see UserStatus
 */
public class User extends BaseEntity {

    private UUID id;
    private UUID tenantId;
    private UUID externalKeycloakId;
    private String email;
    private String name;
    private UserStatus status;
    private java.time.OffsetDateTime invitedDt;

    public User() {
        super();
    }

    // -- Getters / Setters --

    public UUID getTenantId() {
        return tenantId;
    }

    public void setTenantId(UUID tenantId) {
        this.tenantId = tenantId;
    }

    public UUID getExternalKeycloakId() {
        return externalKeycloakId;
    }

    public void setExternalKeycloakId(UUID externalKeycloakId) {
        this.externalKeycloakId = externalKeycloakId;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public UserStatus getStatus() {
        return status;
    }

    public void setStatus(UserStatus status) {
        this.status = status;
    }

    public java.time.OffsetDateTime getInvitedDt() {
        return invitedDt;
    }

    public void setInvitedDt(java.time.OffsetDateTime invitedDt) {
        this.invitedDt = invitedDt;
    }

    // -- Métodos de infraestrutura (BaseRepository) --

    @Override
    public UUID getId() {
        return id;
    }

    @Override
    public void setId(UUID id) {
        this.id = id;
    }

    @Override
    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("tenant_id", tenantId);
        columns.put("external_keycloak_id", externalKeycloakId);
        columns.put("email", email);
        columns.put("name", name);
        columns.put("status", status != null ? status.name() : null);
        columns.put("invited_dt", invitedDt);
        return columns;
    }

    // -- Métodos de domínio --

    public boolean isInvitePending() {
        return status == UserStatus.INVITE_PENDING;
    }

    @Override
    public String toString() {
        return "User{id=" + id + ", email='" + email + '\'' + ", status=" + status + '}';
    }
}
