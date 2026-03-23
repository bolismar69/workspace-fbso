package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public record MunicipioDTO(
    String municipioId,
    String municipioNome,
    String municipioCodigo,
    String ufId,
    String ufNome,
    String ufSigla,
    String regiaoIntermediariaId,
    String regiaoIntermediariaNome,
    String regiaoImediataId,
    String regiaoImediataNome,
    List<DistritoDTO> distrito
) {

  public record DistritoDTO(
      String id,
      String codigo,
      String nome,
      @JsonInclude(JsonInclude.Include.NON_EMPTY)
      List<SubdistritoDTO> subDistrito
  ) {}

  public record SubdistritoDTO(
      String id,
      String codigo,
      String nome
  ) {}
}
