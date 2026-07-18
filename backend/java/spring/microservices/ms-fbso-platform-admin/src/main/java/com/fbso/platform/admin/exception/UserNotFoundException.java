package com.fbso.platform.admin.exception;

import java.util.UUID;

/**
 * Exceção lançada quando um usuário não é encontrado (ou está soft-deleted).
 * <p>
 * Mapeada para HTTP 404 (Not Found) pelo {@link GlobalExceptionHandler}.
 */
public class UserNotFoundException extends BusinessException {

    public UserNotFoundException(UUID userId) {
        super("user-not-found", "Usuário não encontrado: " + userId);
    }
}
