package com.fbso.platform.admin.entity;

import com.fbso.platform.admin.common.BaseEntity;

import java.time.OffsetDateTime;
import java.util.Map;
import java.util.UUID;

/**
 * Entidade imutável de auditoria (F02-05).
 *
 * <p>Mapeia a tabela {@code fbso_platform.audit_log} (V001).
 * Registros são IMUTÁVEIS — apenas INSERT, nunca UPDATE ou DELETE.</p>
 */
public class AuditEntry extends BaseEntity {

    private UUID id;
    private OffsetDateTime timestamp;
    private UUID tenantId;
    private String action;
    private String entityType;
    private UUID entityId;
    private UUID actorId;
    private String actorName;
    private String previousValue;
    private String newValue;
    private String reason;

    public AuditEntry() {
        super();
        this.timestamp = OffsetDateTime.now(java.time.ZoneOffset.UTC);
    }

    // -- Getters / Setters --

    public UUID getId() { return id; }
    public void setId(UUID id) { this.id = id; }
    public OffsetDateTime getTimestamp() { return timestamp; }
    public void setTimestamp(OffsetDateTime timestamp) { this.timestamp = timestamp; }
    public UUID getTenantId() { return tenantId; }
    public void setTenantId(UUID tenantId) { this.tenantId = tenantId; }
    public String getAction() { return action; }
    public void setAction(String action) { this.action = action; }
    public String getEntityType() { return entityType; }
    public void setEntityType(String entityType) { this.entityType = entityType; }
    public UUID getEntityId() { return entityId; }
    public void setEntityId(UUID entityId) { this.entityId = entityId; }
    public UUID getActorId() { return actorId; }
    public void setActorId(UUID actorId) { this.actorId = actorId; }
    public String getActorName() { return actorName; }
    public void setActorName(String actorName) { this.actorName = actorName; }
    public String getPreviousValue() { return previousValue; }
    public void setPreviousValue(String previousValue) { this.previousValue = previousValue; }
    public String getNewValue() { return newValue; }
    public void setNewValue(String newValue) { this.newValue = newValue; }
    public String getReason() { return reason; }
    public void setReason(String reason) { this.reason = reason; }

    @Override
    public Map<String, Object> toColumnMap() {
        return Map.of(); // audit entries are created by AuditAspect, not via BaseRepository.save()
    }

    @Override
    public String toString() {
        return "AuditEntry{id=" + id + ", action=" + action + ", entityType=" + entityType + ", entityId=" + entityId + '}';
    }
}
