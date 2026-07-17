package com.fbso.platform.admin.exception;

import com.fbso.platform.admin.dto.response.ErrorResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

/**
 * Handler global de exceções — converte exceções em respostas RFC 7807.
 * <p>
 * Garantias:
 * <ul>
 *   <li>Toda resposta de erro contém {@code type}, {@code title}, {@code status}, {@code detail}</li>
 *   <li>NUNCA expõe stack traces em respostas HTTP</li>
 *   <li>Mensagens em PT-BR</li>
 *   <li>Exceções não mapeadas → 500 genérico (BR-NFR07)</li>
 * </ul>
 *
 * @see <a href="ARCHITECTURE.md#6">ARCHITECTURE.md §6 — Tratamento de Erros</a>
 */
@ControllerAdvice
public class GlobalExceptionHandler {

    private static final Logger log = LoggerFactory.getLogger(GlobalExceptionHandler.class);

    // ---- 422 — Regras de Negócio ----

    @ExceptionHandler(BusinessException.class)
    public ResponseEntity<ErrorResponse> handleBusinessException(BusinessException ex) {
        log.warn("Erro de negócio: [{}] {}", ex.getErrorCode(), ex.getMessage());
        return ResponseEntity.unprocessableEntity().body(
                ErrorResponse.of(
                        "https://api.fbso.org/errors/" + ex.getErrorCode(),
                        ex.getMessage(),
                        422,
                        null
                )
        );
    }

    // ---- 404 — Recurso Não Encontrado ----

    @ExceptionHandler(TenantNotFoundException.class)
    public ResponseEntity<ErrorResponse> handleTenantNotFound(TenantNotFoundException ex) {
        log.warn("Tenant não encontrado: {}", ex.getMessage());
        ErrorResponse body = ErrorResponse.of(
                "https://api.fbso.org/errors/" + ex.getErrorCode(),
                ex.getMessage(),
                404,
                null
        );
        return new ResponseEntity<>(body, HttpStatus.NOT_FOUND);
    }

    // ---- 409 — Conflito (CNPJ duplicado) ----

    @ExceptionHandler(DuplicateCnpjException.class)
    public ResponseEntity<ErrorResponse> handleDuplicateCnpj(DuplicateCnpjException ex) {
        log.warn("CNPJ duplicado: {}", ex.getMessage());
        ErrorResponse body = ErrorResponse.of(
                "https://api.fbso.org/errors/" + ex.getErrorCode(),
                ex.getMessage(),
                409,
                null
        );
        return new ResponseEntity<>(body, HttpStatus.CONFLICT);
    }

    // ---- 403 — Acesso Negado (Spring Security) ----

    @ExceptionHandler(org.springframework.security.access.AccessDeniedException.class)
    public ResponseEntity<ErrorResponse> handleSpringAccessDenied(
            org.springframework.security.access.AccessDeniedException ex) {
        log.warn("Acesso negado (Spring Security): {}", ex.getMessage());
        return ResponseEntity.status(403).body(
                ErrorResponse.of(
                        "https://api.fbso.org/errors/access-denied",
                        "Acesso negado",
                        403,
                        "Você não tem permissão para acessar esta área."
                )
        );
    }

    // ---- 403 — Acesso Negado (Custom) ----

    @ExceptionHandler(PermissionDeniedException.class)
    public ResponseEntity<ErrorResponse> handlePermissionDenied(PermissionDeniedException ex) {
        log.warn("Acesso negado: {}", ex.getMessage());
        return ResponseEntity.status(403).body(
                ErrorResponse.of(
                        "https://api.fbso.org/errors/access-denied",
                        "Acesso negado",
                        403,
                        "Você não tem permissão para acessar esta área."
                )
        );
    }

    // ---- 400 — Validação (Bean Validation) ----

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ErrorResponse> handleValidation(MethodArgumentNotValidException ex) {
        var fieldErrors = ex.getBindingResult().getFieldErrors().stream()
                .map(fe -> new ErrorResponse.FieldError(fe.getField(), fe.getDefaultMessage()))
                .toList();

        log.debug("Erro de validação: {} campos", fieldErrors.size());
        return ResponseEntity.badRequest().body(
                ErrorResponse.validation("Erro de validação", fieldErrors)
        );
    }

    // ---- 401 — Segurança (Tenant Isolation) ----

    @ExceptionHandler(SecurityException.class)
    public ResponseEntity<ErrorResponse> handleSecurityException(SecurityException ex) {
        log.error("Violação de segurança: {}", ex.getMessage());
        return ResponseEntity.status(401).body(
                ErrorResponse.of(
                        "https://api.fbso.org/errors/unauthorized",
                        "Token de acesso não informado",
                        401,
                        null
                )
        );
    }

    // ---- 500 — Erro Interno (genérico, SEM stack trace) ----

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleGenericException(Exception ex) {
        log.error("Erro interno não tratado: {}", ex.getClass().getName(), ex);
        return ResponseEntity.internalServerError().body(
                ErrorResponse.of(
                        "https://api.fbso.org/errors/internal-error",
                        "Erro interno do servidor",
                        500,
                        "Ocorreu um erro inesperado. Por favor, tente novamente."
                )
        );
    }
}
