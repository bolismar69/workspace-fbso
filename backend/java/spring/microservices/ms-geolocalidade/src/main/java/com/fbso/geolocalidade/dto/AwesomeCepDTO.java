package com.fbso.geolocalidade.dto;

public record AwesomeCepDTO(
    String cep,
    String address,
    String state,
    String city,
    String city_ibge,
    String lat,
    String lng,
    String district,
    String d
) {}
