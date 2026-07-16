package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonInclude;

import java.util.List;

/**
 * DTO de resposta de erro — RFC 7807 (Problem Details).
 * <p>
 * Campos:
 * <ul>
 *   <li>{@code type} — URI identificando o tipo de erro</li>
 *   <li>{@code title} — título legível em PT-BR</li>
 *   <li>{@code status} — código HTTP</li>
 *   <li>{@code detail} — descrição adicional (opcional)</li>
 *   <li>{@code fields} — erros de validação por campo (opcional)</li>
 * </ul>
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
public record ErrorResponse(
        String type,
        String title,
        int status,
        String detail,
        List<FieldError> fields
) {

    /**
     * Erro de campo individual (validação Bean Validation).
     */
    @JsonInclude(JsonInclude.Include.NON_NULL)
    public record FieldError(String field, String message) {}

    // ---- Factory Methods ----

    public static ErrorResponse of(String type, String title, int status, String detail) {
        return new ErrorResponse(type, title, status, detail, null);
    }

    public static ErrorResponse validation(String title, List<FieldError> fields) {
        return new ErrorResponse(
                "https://api.fbso.org/errors/validation-error",
                title != null ? title : "Erro de validação",
                400,
                "Um ou mais campos contêm valores inválidos",
                fields
        );
    }
}
