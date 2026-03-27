package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public record CepInfoDTO(
    AwesomeCepResponseDTO cepInfo
) {}
