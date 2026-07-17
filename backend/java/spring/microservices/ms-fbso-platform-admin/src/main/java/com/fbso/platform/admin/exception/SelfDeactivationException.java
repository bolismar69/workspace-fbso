package com.fbso.platform.admin.exception;

/**
 * Exceção lançada quando um administrador tenta desativar a si mesmo.
 * <p>
 * Mapeada para HTTP 422 (Unprocessable Entity) pelo {@link GlobalExceptionHandler}.
 * Implementa RN09-03: Um administrador não pode desativar a si mesmo.
 */
public class SelfDeactivationException extends BusinessException {

    public SelfDeactivationException() {
        super("self-deactivation", "Um administrador não pode desativar a si mesmo.");
    }
}
