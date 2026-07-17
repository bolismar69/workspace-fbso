package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;

/**
 * DTO para criação/convite de usuário (F03-01).
 *
 * <p>RN09-02: Email único por tenant ativo — validado no {@code UserService}.
 */
public record UserCreateRequest(
        @NotBlank(message = "Nome é obrigatório")
        String name,

        @NotBlank(message = "Email é obrigatório")
        @Email(message = "Formato de email inválido")
        String email
) {}
