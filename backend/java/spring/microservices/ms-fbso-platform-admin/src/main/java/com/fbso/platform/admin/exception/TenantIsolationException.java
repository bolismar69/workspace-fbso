package com.fbso.platform.admin.exception;

/**
 * Exceção lançada quando ocorre falha no isolamento multi-tenant.
 * <p>
 * Indica que não foi possível configurar o contexto de tenant na conexão
 * com o banco de dados — a conexão NÃO pode retornar ao pool.
 * <p>
 * Mapeada para HTTP 500 (Internal Server Error) pelo {@link GlobalExceptionHandler}.
 *
 * @see com.fbso.platform.admin.config.TenantAwareDataSource
 */
public class TenantIsolationException extends RuntimeException {

    public TenantIsolationException(String message, Throwable cause) {
        super(message, cause);
    }
}
