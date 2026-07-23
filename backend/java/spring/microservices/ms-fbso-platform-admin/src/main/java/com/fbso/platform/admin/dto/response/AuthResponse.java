package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public record AuthResponse(
    String redirectUrl,
    String message,
    String tokenType,
    Integer expiresIn
) {
    public static AuthResponse loginRedirect(String redirectUrl) {
        return new AuthResponse(redirectUrl, "Redirecionando para autenticação", null, null);
    }

    public static AuthResponse forgotPassword(String message) {
        return new AuthResponse(null, message, null, null);
    }

    public static AuthResponse resetPassword(String message) {
        return new AuthResponse(null, message, null, null);
    }
}
