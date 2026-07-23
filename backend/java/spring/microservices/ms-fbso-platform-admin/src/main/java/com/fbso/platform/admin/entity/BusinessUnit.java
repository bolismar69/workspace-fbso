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
 *   <li>{@code parentId = null} → Matriz (raiz da hierarquia)</li>
 *   <li>{@code parentId != null} → Filial</li>
 *   <li>{@code isMatrix = true} → flag explícita de Matriz (V007)</li>
 * </ul>
 *
 * <h3>Campos de Endereço</h3>
 * <p>Embedded diretamente na tabela (street, number, complement,
 * neighborhood, city, state, zipCode). O Value Object {@code Address.java}
 * existe no classpath mas não é utilizado por esta entidade.</p>
 *
 * <p><b>DT-126 (Sprint 6):</b> Entity reescrita para alinhar 100% com o
 * schema V001. Removidos: {@code name} (→ corporateName), {@code hierarchyType}
 * (sem coluna no DB). Adicionados: corporateName, taxRegime, street, number,
 * complement, neighborhood, city, state, zipCode, status.</p>
 */
public class BusinessUnit extends BaseEntity {

    private UUID id;
    private UUID tenantId;
    private UUID parentId;
    private String cnpj;
    private String corporateName;
    private String taxRegime;
    private String street;
    private String number;
    private String complement;
    private String neighborhood;
    private String city;
    private String state;
    private String zipCode;
    private String status;
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

    public UUID getParentId() {
        return parentId;
    }

    public void setParentId(UUID parentId) {
        this.parentId = parentId;
    }

    public String getCnpj() {
        return cnpj;
    }

    public void setCnpj(String cnpj) {
        this.cnpj = cnpj;
    }

    public String getCorporateName() {
        return corporateName;
    }

    public void setCorporateName(String corporateName) {
        this.corporateName = corporateName;
    }

    public String getTaxRegime() {
        return taxRegime;
    }

    public void setTaxRegime(String taxRegime) {
        this.taxRegime = taxRegime;
    }

    public String getStreet() {
        return street;
    }

    public void setStreet(String street) {
        this.street = street;
    }

    public String getNumber() {
        return number;
    }

    public void setNumber(String number) {
        this.number = number;
    }

    public String getComplement() {
        return complement;
    }

    public void setComplement(String complement) {
        this.complement = complement;
    }

    public String getNeighborhood() {
        return neighborhood;
    }

    public void setNeighborhood(String neighborhood) {
        this.neighborhood = neighborhood;
    }

    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
    }

    public String getState() {
        return state;
    }

    public void setState(String state) {
        this.state = state;
    }

    public String getZipCode() {
        return zipCode;
    }

    public void setZipCode(String zipCode) {
        this.zipCode = zipCode;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
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

    /**
     * Mapeia as colunas de domínio para INSERT/UPDATE no BaseRepository.
     * Colunas de auditoria (created_dt, updated_dt, created_by, updated_by,
     * deleted_dt, deleted_by) são gerenciadas automaticamente pelo BaseRepository.
     *
     * @return mapa ordenado coluna → valor (apenas colunas de domínio)
     */
    @Override
    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("tenant_id", tenantId);
        columns.put("parent_id", parentId);
        columns.put("cnpj", cnpj);
        columns.put("corporate_name", corporateName);
        columns.put("tax_regime", taxRegime);
        columns.put("street", street);
        columns.put("number", number);
        columns.put("complement", complement);
        columns.put("neighborhood", neighborhood);
        columns.put("city", city);
        columns.put("state", state);
        columns.put("zip_code", zipCode);
        columns.put("status", status);
        columns.put("is_matrix", isMatrix);
        return columns;
    }

    @Override
    public String toString() {
        return "BusinessUnit{id=" + id + ", corporateName='" + corporateName
                + "', cnpj='" + cnpj + "', isMatrix=" + isMatrix + "}";
    }
}
