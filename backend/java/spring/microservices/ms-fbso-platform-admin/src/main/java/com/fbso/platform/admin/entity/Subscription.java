package com.fbso.platform.admin.entity;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.enums.Recurrence;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa uma assinatura de Tenant a um Plano (F02-04).
 *
 * <p>Mapeia a tabela {@code fbso_platform.subscription} (V001).
 * Tabela com RLS — {@code hasTenantColumn = true}.</p>
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN07-01: 1 assinatura ativa por tenant</li>
 *   <li>DT-009: locked_price + locked_recurrence preservam preço original</li>
 * </ul>
 */
public class Subscription extends BaseEntity {

    private UUID id;
    private UUID tenantId;
    private UUID planId;
    private OffsetDateTime startDate;
    private OffsetDateTime endDate;
    private String status; // ACTIVE, SUSPENDED, CANCELED
    private BigDecimal lockedPrice;
    private String lockedRecurrence;

    public Subscription() {
        super();
        this.startDate = OffsetDateTime.now(java.time.ZoneOffset.UTC);
        this.status = "ACTIVE";
    }

    // -- Getters / Setters --

    public UUID getId() { return id; }
    public void setId(UUID id) { this.id = id; }
    public UUID getTenantId() { return tenantId; }
    public void setTenantId(UUID tenantId) { this.tenantId = tenantId; }
    public UUID getPlanId() { return planId; }
    public void setPlanId(UUID planId) { this.planId = planId; }
    public OffsetDateTime getStartDate() { return startDate; }
    public void setStartDate(OffsetDateTime startDate) { this.startDate = startDate; }
    public OffsetDateTime getEndDate() { return endDate; }
    public void setEndDate(OffsetDateTime endDate) { this.endDate = endDate; }
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    public BigDecimal getLockedPrice() { return lockedPrice; }
    public void setLockedPrice(BigDecimal lockedPrice) { this.lockedPrice = lockedPrice; }
    public String getLockedRecurrence() { return lockedRecurrence; }
    public void setLockedRecurrence(String lockedRecurrence) { this.lockedRecurrence = lockedRecurrence; }

    @Override
    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        // tenant_id é gerenciado pelo BaseRepository (hasTenantColumn=true)
        columns.put("plan_id", planId);
        columns.put("start_date", startDate);
        columns.put("end_date", endDate);
        columns.put("status", status);
        columns.put("locked_price", lockedPrice);
        columns.put("locked_recurrence", lockedRecurrence);
        return columns;
    }

    @Override
    public String toString() {
        return "Subscription{id=" + id + ", tenantId=" + tenantId + ", planId=" + planId + ", status=" + status + '}';
    }
}
