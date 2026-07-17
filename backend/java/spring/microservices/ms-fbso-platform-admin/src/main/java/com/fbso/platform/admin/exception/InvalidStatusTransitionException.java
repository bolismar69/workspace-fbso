package com.fbso.platform.admin.exception;

/**
 * Exceção lançada quando uma transição de status inválida é solicitada.
 * <p>
 * Mapeada para HTTP 422 (Unprocessable Entity) pelo {@link GlobalExceptionHandler}.
 */
public class InvalidStatusTransitionException extends BusinessException {

    public InvalidStatusTransitionException(String from, String to) {
        super("invalid-status-transition",
                "Transição de status inválida: " + from + " → " + to);
    }
}
