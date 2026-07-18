package com.fbso.platform.admin.dto.request;

import jakarta.validation.Valid;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;

import java.util.List;
import java.util.UUID;

/**
 * DTO para atualização em lote de permissões (F03-02, F03-03).
 *
 * <p>Substitui completamente as permissões do usuário: as permissões atuais
 * são removidas e as novas são inseridas em uma única transação.
 */
public record PermissionUpdateRequest(
        @NotEmpty(message = "A lista de permissões não pode estar vazia")
        List<@Valid PermissionAssignment> permissions
) {
    /**
     * Par (businessUnitId, role) que compõe uma atribuição de permissão.
     */
    public record PermissionAssignment(
            @NotNull(message = "businessUnitId é obrigatório")
            UUID businessUnitId,

            @NotBlank(message = "role é obrigatório")
            String role
    ) {}
}
