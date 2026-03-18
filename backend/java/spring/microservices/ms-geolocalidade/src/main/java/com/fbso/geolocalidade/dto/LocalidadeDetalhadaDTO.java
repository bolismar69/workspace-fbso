package com.fbso.geolocalidade.dto;

public record LocalidadeDetalhadaDTO(
    String codigoIbge,
    String municipio,
    String uf,
    String lat,
    String lng
) {}
