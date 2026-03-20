package com.fbso.geolocalidade.dto;

import java.time.LocalDateTime;

public record ResponseStatusDTO(
    int code,             // Ex: 200, 400, 500
    String status,       // Ex: "OK", "Bad Request"
    String message,      // Mensagem amigável
    LocalDateTime timestamp
) {
    public static ResponseStatusDTO success(int code, String message) {
        return new ResponseStatusDTO(code, "SUCCESS", message, LocalDateTime.now());
    }

    public static ResponseStatusDTO error(int code, String message) {
        return new ResponseStatusDTO(code, "ERROR", message, LocalDateTime.now());
    }

    public static ResponseStatusDTO fromHttpStatus(org.springframework.http.HttpStatus httpStatus, String message) {
        return new ResponseStatusDTO(httpStatus.value(), httpStatus.getReasonPhrase(), message, LocalDateTime.now());
    }

    public static ResponseStatusDTO fromException(Exception ex) {
        // Aqui você pode mapear diferentes tipos de exceções para códigos e mensagens específicas
        if (ex instanceof IllegalArgumentException) {
            return error(400, ex.getMessage());
        } else if (ex instanceof NullPointerException) {
            return error(500, "Ocorreu um erro interno.");
        } else {
            return error(500, "Erro desconhecido: " + ex.getMessage());
        }
    }
    
    // Você pode criar outros como error() ou a partir de um HttpStatus
}
