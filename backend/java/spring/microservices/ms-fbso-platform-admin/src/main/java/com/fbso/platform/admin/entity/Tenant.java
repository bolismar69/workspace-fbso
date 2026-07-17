package com.fbso.platform.admin.entity;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.enums.TenantSegment;
import com.fbso.platform.admin.enums.TenantStatus;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Representa uma conta corporativa (Tenant) na plataforma FBSO.
 *
 * <p>Mapeia a tabela {@code fbso_platform.tenant} (V001).
 * Esta é uma tabela GLOBAL — visível apenas pelo Admin FBSO, sem RLS.</p>
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN04-02: Razão social obrigatória, única entre ativos (validado no Service)</li>
 *   <li>RN05-01: Transições de status (máquina de estados)</li>
 * </ul>
 *
 * @see TenantStatus
 * @see TenantSegment
 */
public class Tenant extends BaseEntity {

    private UUID id;
    private String nameCorporate;
    private String nameFantasy;
    private TenantSegment segment;
    private TenantStatus status;

    public Tenant() {
        super();
    }

    // -- Getters / Setters --

    public String getNameCorporate() {
        return nameCorporate;
    }

    public void setNameCorporate(String nameCorporate) {
        this.nameCorporate = nameCorporate;
    }

    public String getNameFantasy() {
        return nameFantasy;
    }

    public void setNameFantasy(String nameFantasy) {
        this.nameFantasy = nameFantasy;
    }

    public TenantSegment getSegment() {
        return segment;
    }

    public void setSegment(TenantSegment segment) {
        this.segment = segment;
    }

    public TenantStatus getStatus() {
        return status;
    }

    public void setStatus(TenantStatus status) {
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

    @Override
    public Map<String, Object> toColumnMap() {
        Map<String, Object> columns = new LinkedHashMap<>();
        columns.put("name_corporate", nameCorporate);
        columns.put("name_fantasy", nameFantasy);
        columns.put("segment", segment != null ? segment.name() : null);
        columns.put("status", status != null ? status.name() : null);
        return columns;
    }

    // -- Métodos de domínio (RN05-01: máquina de estados) --

    /**
     * @return {@code true} se este tenant está ativo (não suspenso, não inativo, não deletado)
     */
    public boolean isOperational() {
        return isActive() && status == TenantStatus.ACTIVE;
    }

    @Override
    public String toString() {
        return "Tenant{id=" + id + ", nameCorporate='" + nameCorporate + '\'' + ", status=" + status + '}';
    }
}
