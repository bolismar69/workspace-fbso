package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.Email;

/**
 * DTO para edição de usuário (F03-01).
 * Todos os campos são opcionais — apenas os informados são atualizados.
 */
public record UserUpdateRequest(
        String name,

        @Email(message = "Formato de email inválido")
        String email
) {}
