package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonAlias;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record AwesomeCepDTO(
    // @JsonProperty("cep")
    String cep,
    @JsonAlias("address_type")
    String address_type,
    @JsonAlias("address_name")
    String address_name,
    String address,
    String state,
    String city,
    @JsonAlias("city_ibge")
    String city_ibge,
    String ddd,
    String lat,
    String lng,
    String district,
    @JsonAlias("distance_km")
    Double distance_km
) {}
