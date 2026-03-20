package com.fbso.geolocalidade.dto;

import java.util.List;

public record RespostaCompletaDTO(
    AwesomeCepDTO zipcodeInfo,
    LocalidadeDetalhadaDTO localidade,
    List<VizinhoEnriquecidoDTO> cidadesProximas
) {}
