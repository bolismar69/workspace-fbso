package com.fbso.platform.admin.entity;

import com.fbso.platform.admin.common.BaseEntity;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa uma Unidade de Negócio (Business Unit) de um tenant.
 *
 * <p>Mapeia a tabela {@code fbso_platform.business_unit} (V001).
 * RLS ativo (V003) — filtrado por {@code tenant_id}.</p>
 *
 * <h3>Hierarquia</h3>
 * <ul>
 *   <li>{@code parent_id = NULL} → Matriz (raiz da hierarquia)</li>
 *   <li>{@code parent_id != NULL} → Filial</li>
 * </ul>
 *
 * <p><b>Nota:</b> Entity mínima criada na Sprint 4 para referência no
 * {@code UserPermission}. Flag {@code isMatrix} adicionada na Sprint 5 (DT-107/T-142).
 * O CRUD completo será implementado na Sprint 6 (M6).</p>
 */
public class BusinessUnit extends BaseEntity {

    private UUID id;
    private UUID tenantId;
    private String name;
    private String cnpj;
    private String hierarchyType;
    private UUID parentId;
    private boolean isMatrix;

    public BusinessUnit() {
        super();
    }

    // -- Getters / Setters --

    public UUID getTenantId() {
        return tenantId;
    }

    public void setTenantId(UUID tenantId) {
        this.tenantId = tenantId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCnpj() {
        return cnpj;
    }

    public void setCnpj(String cnpj) {
        this.cnpj = cnpj;
    }

    public String getHierarchyType() {
        return hierarchyType;
    }

    public void setHierarchyType(String hierarchyType) {
        this.hierarchyType = hierarchyType;
    }

    public UUID getParentId() {
        return parentId;
    }

    public void setParentId(UUID parentId) {
        this.parentId = parentId;
    }

    public boolean isMatrix() {
        return isMatrix;
    }

    public void setMatrix(boolean matrix) {
        isMatrix = matrix;
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
        columns.put("name", name);
        columns.put("cnpj", cnpj);
        columns.put("hierarchy_type", hierarchyType);
        columns.put("parent_id", parentId);
        columns.put("is_matrix", isMatrix);
        return columns;
    }

    @Override
    public String toString() {
        return "BusinessUnit{id=" + id + ", name='" + name + '\'' + ", cnpj='" + cnpj + "'}";
    }
}
