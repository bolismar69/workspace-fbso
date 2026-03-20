package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record AwesomeCepDTO(
    // @JsonProperty("zipcode")
    String cep,
    String address_type,
    String address_name,
    String address,
    String state,
    String city,
    String city_ibge,
    String lat,
    String lng,
    String district,
    Double distance_km
) {}
