package com.fbso.platform.admin.common;

import java.time.OffsetDateTime;
import java.util.Collections;
import java.util.Map;
import java.util.UUID;

/**
 * Base entity com campos de auditoria compartilhados por todas as tabelas.
 * <p>
 * Campos obrigatórios em TODA tabela do schema {@code fbso_platform}:
 * <ul>
 *   <li>{@code created_dt} — data/hora de criação (NOT NULL)</li>
 *   <li>{@code updated_dt} — data/hora da última atualização (NOT NULL)</li>
 *   <li>{@code created_by} — UUID do usuário criador</li>
 *   <li>{@code updated_by} — UUID do usuário da última atualização</li>
 *   <li>{@code deleted_dt} — soft delete (NULL = ativo)</li>
 *   <li>{@code deleted_by} — UUID do usuário que excluiu</li>
 * </ul>
 *
 * @see <a href="SPECS.md#6.3">SPECS.md §6.3 — Campos de Auditoria</a>
 */
public abstract class BaseEntity {

    private OffsetDateTime createdDt;
    private OffsetDateTime updatedDt;
    private UUID createdBy;
    private UUID updatedBy;
    private OffsetDateTime deletedDt;
    private UUID deletedBy;

    protected BaseEntity() {
        this.createdDt = OffsetDateTime.now(java.time.ZoneOffset.UTC);
        this.updatedDt = OffsetDateTime.now(java.time.ZoneOffset.UTC);
    }

    public OffsetDateTime getCreatedDt() {
        return createdDt;
    }

    public void setCreatedDt(OffsetDateTime createdDt) {
        this.createdDt = createdDt;
    }

    public OffsetDateTime getUpdatedDt() {
        return updatedDt;
    }

    public void setUpdatedDt(OffsetDateTime updatedDt) {
        this.updatedDt = updatedDt;
    }

    public UUID getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(UUID createdBy) {
        this.createdBy = createdBy;
    }

    public UUID getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(UUID updatedBy) {
        this.updatedBy = updatedBy;
    }

    public OffsetDateTime getDeletedDt() {
        return deletedDt;
    }

    public void setDeletedDt(OffsetDateTime deletedDt) {
        this.deletedDt = deletedDt;
    }

    public UUID getDeletedBy() {
        return deletedBy;
    }

    public void setDeletedBy(UUID deletedBy) {
        this.deletedBy = deletedBy;
    }

    /**
     * @return o ID da entidade (implementado por cada subclasse)
     */
    public abstract UUID getId();

    /**
     * Define o ID da entidade (usado pelo BaseRepository após INSERT).
     */
    public abstract void setId(UUID id);

    /**
     * @return mapa coluna→valor com os campos específicos da entidade
     *         (excluindo id, created_dt, updated_dt, created_by, updated_by,
     *          deleted_dt, deleted_by — que são gerenciados pelo BaseRepository)
     */
    public Map<String, Object> toColumnMap() {
        return Collections.emptyMap();
    }

    /**
     * @return {@code true} se o registro está ativo (não sofreu soft delete)
     */
    public boolean isActive() {
        return deletedDt == null;
    }

    /**
     * Marca o registro como excluído (soft delete).
     *
     * @param deletedBy UUID do usuário que executou a exclusão
     */
    public void markAsDeleted(UUID deletedBy) {
        this.deletedDt = OffsetDateTime.now(java.time.ZoneOffset.UTC);
        this.deletedBy = deletedBy;
    }
}
