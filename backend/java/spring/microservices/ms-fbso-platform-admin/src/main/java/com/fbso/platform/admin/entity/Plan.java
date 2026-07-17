package com.fbso.platform.admin.entity;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.enums.Recurrence;

import java.math.BigDecimal;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa um plano comercial (F02-03).
 *
 * <p>Mapeia a tabela {@code fbso_platform.plan} (V001).</p>
 */
public class Plan extends BaseEntity {

    private UUID id;
    private String name;
    private String description;
    private BigDecimal price;
    private Recurrence recurrence;
    private String status; // ACTIVE, DISCONTINUED
    private int version;

    public Plan() {
        super();
        this.version = 1;
        this.status = "ACTIVE";
    }

    // -- Getters / Setters --

    public UUID getId() { return id; }
    public void setId(UUID id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getDescription() { return description; }
    public void setDescription(String description) { this.description = description; }
    public BigDecimal getPrice() { return price; }
    public void setPrice(BigDecimal price) { this.price = price; }
    public Recurrence getRecurrence() { return recurrence; }
    public void setRecurrence(Recurrence recurrence) { this.recurrence = recurrence; }
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    public int getVersion() { return version; }
    public void setVersion(int version) { this.version = version; }

    // -- toColumnMap (BaseRepository) --

    @Override
    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("name", name);
        columns.put("description", description);
        columns.put("price", price);
        columns.put("recurrence", recurrence != null ? recurrence.name() : null);
        columns.put("status", status);
        columns.put("version", version);
        return columns;
    }

    public boolean isActive() {
        return "ACTIVE".equals(status) && super.isActive();
    }

    @Override
    public String toString() {
        return "Plan{id=" + id + ", name='" + name + '\'' + ", status=" + status + ", version=" + version + '}';
    }
}
