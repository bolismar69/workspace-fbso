package com.fbso.platform.admin.dto.response;

import com.fbso.platform.admin.entity.User;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.UUID;

/**
 * DTO de resposta com dados do usuário (F03-01).
 *
 * <p>Inclui nome, email, status, papel, BUs vinculadas e datas.
 */
public record UserResponse(
        UUID id,
        String name,
        String email,
        String status,
        List<String> roles,
        List<UUID> businessUnitIds,
        OffsetDateTime invitedDt,
        OffsetDateTime createdAt
) {
    /**
     * Factory method a partir da entity User.
     * <p>
     * {@code businessUnitIds} é populado como lista vazia na Frente 1 —
     * será preenchido com dados reais de {@code user_permission} na Frente 3.
     */
    public static UserResponse from(User user) {
        return new UserResponse(
                user.getId(),
                user.getName(),
                user.getEmail(),
                user.getStatus() != null ? user.getStatus().name() : null,
                List.of(),   // Frente 3: popular via user_permission
                List.of(),   // Frente 3: popular via user_permission
                user.getInvitedDt(),
                user.getCreatedDt()
        );
    }
}
