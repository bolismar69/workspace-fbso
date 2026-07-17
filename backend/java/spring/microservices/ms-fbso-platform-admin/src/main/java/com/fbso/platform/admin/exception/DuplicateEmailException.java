package com.fbso.platform.admin.exception;

/**
 * Exceção lançada quando há tentativa de cadastrar um email duplicado no mesmo tenant.
 * <p>
 * Mapeada para HTTP 409 (Conflict) pelo {@link GlobalExceptionHandler}.
 * Implementa RN09-02: Email único por tenant ativo.
 */
public class DuplicateEmailException extends BusinessException {

    public DuplicateEmailException(String email) {
        super("duplicate-email", "Este email já está em uso no tenant: " + email);
    }
}
