package com.fbso.platform.admin.exception;

/**
 * Exceção base para erros de regra de negócio.
 * <p>
 * Mapeada para HTTP 422 (Unprocessable Entity) pelo {@link GlobalExceptionHandler}.
 * Subclasses representam erros específicos:
 * <ul>
 *   <li>{@code DuplicateCnpjException}</li>
 *   <li>{@code InvalidStatusTransitionException}</li>
 *   <li>{@code PlanHasActiveSubscribersException}</li>
 *   <li>{@code TenantNotFoundException}</li>
 * </ul>
 */
public class BusinessException extends RuntimeException {

    private final String errorCode;

    public BusinessException(String message) {
        super(message);
        this.errorCode = "business-error";
    }

    public BusinessException(String errorCode, String message) {
        super(message);
        this.errorCode = errorCode;
    }

    public String getErrorCode() {
        return errorCode;
    }
}
