package com.fbso.geolocalidade.dto;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonAlias;
import com.fasterxml.jackson.annotation.JsonProperty;

public record RespostaCompletaDTO(
    @JsonProperty("cepInfo")
    @JsonAlias("zipcodeInfo")
    AwesomeCepDTO cepInfo,
    // LocalidadeDetalhadaDTO localidade,
    List<VizinhoEnriquecidoDTO> cidadesProximas
) {}
