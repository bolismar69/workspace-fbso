package com.fbso.platform.admin.exception;

/**
 * Exceção lançada quando o usuário não tem permissão para acessar um recurso.
 * <p>
 * Mapeada para HTTP 403 (Forbidden) pelo {@link GlobalExceptionHandler}.
 * Mensagem amigável em PT-BR, sem detalhes técnicos (RN12-01, RN12-02).
 */
public class PermissionDeniedException extends RuntimeException {

    private static final String DEFAULT_MESSAGE = "Você não tem permissão para executar esta operação.";

    public PermissionDeniedException() {
        super(DEFAULT_MESSAGE);
    }

    public PermissionDeniedException(String message) {
        super(message);
    }
}
