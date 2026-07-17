package com.fbso.platform.admin.exception;

/**
 * Exceção lançada quando há tentativa de cadastrar um CNPJ duplicado.
 * <p>
 * Mapeada para HTTP 409 (Conflict) pelo {@link GlobalExceptionHandler}.
 */
public class DuplicateCnpjException extends BusinessException {

    public DuplicateCnpjException(String cnpj) {
        super("duplicate-cnpj", "Já existe uma conta ativa com o CNPJ " + cnpj);
    }
}
