package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.ALWAYS)
public record AwesomeCepResponseDTO(
    String cep,
    String addressType,
    String addressName,
    String address,
    String state,
    String city,
    String cityIbge,
    String ddd,
    String lat,
    String lng,
    String district,
    Double distanceKm
) {

  public static AwesomeCepResponseDTO from(AwesomeCepDTO dto) {
    if (dto == null) {
      return null;
    }

    return new AwesomeCepResponseDTO(
        dto.cep(),
        dto.address_type(),
        dto.address_name(),
        dto.address(),
        dto.state(),
        dto.city(),
        dto.city_ibge(),
        dto.ddd(),
        dto.lat(),
        dto.lng(),
        dto.district(),
        dto.distance_km());
  }
}
