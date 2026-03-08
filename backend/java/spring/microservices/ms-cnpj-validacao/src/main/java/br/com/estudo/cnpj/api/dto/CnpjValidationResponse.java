package br.com.estudo.cnpj.api.dto;

public record CnpjValidationResponse(
        String input,
        String normalized,
        boolean valid
) {
}
