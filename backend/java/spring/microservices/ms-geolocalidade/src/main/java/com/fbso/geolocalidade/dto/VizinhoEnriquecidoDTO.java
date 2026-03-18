package com.fbso.geolocalidade.dto;

public record VizinhoEnriquecidoDTO(
    String cep,
    String cidade,
    String ibge,
    String distrito,
    String subdistrito,
    Double distanciaKm
) {}
