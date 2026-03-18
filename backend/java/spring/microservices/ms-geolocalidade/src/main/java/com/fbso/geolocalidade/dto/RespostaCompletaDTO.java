package com.fbso.geolocalidade.dto;

import java.util.List;

public record RespostaCompletaDTO(
    AwesomeCepDTO viacep,
    LocalidadeDetalhadaDTO localidade,
    List<VizinhoEnriquecidoDTO> cidadesProximas
) {}
