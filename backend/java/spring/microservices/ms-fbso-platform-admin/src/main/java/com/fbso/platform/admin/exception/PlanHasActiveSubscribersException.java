package com.fbso.platform.admin.exception;

/**
 * Exceção lançada ao tentar desativar um plano que possui assinantes ativos.
 * <p>
 * Mapeada para HTTP 422 (Unprocessable Entity) pelo {@link GlobalExceptionHandler}.
 */
public class PlanHasActiveSubscribersException extends BusinessException {

    public PlanHasActiveSubscribersException(String planName) {
        super("plan-has-active-subscribers",
                "Não é possível desativar o plano '" + planName + "' pois existem assinantes ativos");
    }
}
