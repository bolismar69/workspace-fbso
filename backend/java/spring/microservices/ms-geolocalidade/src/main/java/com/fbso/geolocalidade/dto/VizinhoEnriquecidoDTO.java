package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public record VizinhoEnriquecidoDTO(
    // @JsonProperty("cep")
    String cep,
    String city,
    String cityIbge,
    String district,
    String subdistrict,
    Double distanceKm
) {}
