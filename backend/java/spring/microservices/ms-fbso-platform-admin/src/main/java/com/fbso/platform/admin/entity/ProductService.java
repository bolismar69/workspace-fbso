package com.fbso.platform.admin.entity;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.enums.ProductType;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa um produto ou serviço do catálogo (F04-06).
 *
 * <p>Mapeia a tabela {@code fbso_platform.product_service} (V001).</p>
 *
 * <h3>Isolamento Multi-Tenant</h3>
 * <p>Esta tabela NÃO possui coluna {@code tenant_id} própria.
 * O isolamento entre tenants é feito via JOIN com
 * {@code business_unit.tenant_id}. O repositório deve usar
 * {@code hasTenantColumn=false} e filtrar manualmente.</p>
 *
 * <p><b>DT-127 (Sprint 6):</b> Entity criada — tabela existia desde V001
 * mas não havia classe Java correspondente.</p>
 */
public class ProductService extends BaseEntity {

    private UUID id;
    private UUID businessUnitId;
    private String name;
    private String sku;
    private ProductType type;
    private String description;
    private String status;

    public ProductService() {
        super();
        this.type = ProductType.SERVICE;
        this.status = "ACTIVE";
    }

    // -- Getters / Setters --

    public UUID getBusinessUnitId() {
        return businessUnitId;
    }

    public void setBusinessUnitId(UUID businessUnitId) {
        this.businessUnitId = businessUnitId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getSku() {
        return sku;
    }

    public void setSku(String sku) {
        this.sku = sku;
    }

    public ProductType getType() {
        return type;
    }

    public void setType(ProductType type) {
        this.type = type;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
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

    /**
     * Mapeia as colunas de domínio para INSERT/UPDATE no BaseRepository.
     * Apenas colunas de domínio — auditoria é gerenciada pelo BaseRepository.
     *
     * @return mapa ordenado coluna → valor
     */
    @Override
    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("business_unit_id", businessUnitId);
        columns.put("name", name);
        columns.put("sku", sku);
        columns.put("type", type != null ? type.name() : null);
        columns.put("description", description);
        columns.put("status", status);
        return columns;
    }

    public boolean isActive() {
        return "ACTIVE".equals(status) && super.isActive();
    }

    @Override
    public String toString() {
        return "ProductService{id=" + id + ", name='" + name + "', sku='"
                + sku + "', type=" + type + ", status=" + status + '}';
    }
}
