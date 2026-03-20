package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public record VizinhoEnriquecidoDTO(
    // @JsonProperty("zipcode")
    String cep,
    String city,
    String city_ibge,
    String district,
    String subdistrict,
    Double distance_km
) {}
