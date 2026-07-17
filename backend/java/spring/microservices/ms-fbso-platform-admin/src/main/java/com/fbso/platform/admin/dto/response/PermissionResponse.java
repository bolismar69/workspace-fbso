package com.fbso.platform.admin.dto.response;

import com.fbso.platform.admin.entity.UserPermission;

import java.util.UUID;

/**
 * DTO de resposta com dados de uma permissão (F03-02, F03-03).
 */
public record PermissionResponse(
        UUID userId,
        UUID businessUnitId,
        String role
) {
    public static PermissionResponse from(UserPermission up) {
        return new PermissionResponse(
                up.getUserId(),
                up.getBusinessUnitId(),
                up.getRole()
        );
    }
}
