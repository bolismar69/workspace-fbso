package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public record MunicipioDTO(
    @JsonProperty("localidade") LocalidadeDTO localidade
) {

  public record LocalidadeDTO(
      @JsonProperty("uf") UfDTO uf,
      @JsonProperty("regiao_intermediaria") RegiaoIntermediariaDTO regiaoIntermediaria,
      @JsonProperty("regiao_imediata") RegiaoImediataDTO regiaoImediata,
      @JsonProperty("municipio") MunicipioInfoDTO municipio,
      @JsonProperty("distrito") DistritoDTO distrito,
      @JsonProperty("subdistrito") SubdistritoDTO subdistrito
  ) {}

  public record UfDTO(
      String id,
      String sigla,
      String nome
  ) {}

  public record RegiaoIntermediariaDTO(
      String id,
      String nome
  ) {}

  public record RegiaoImediataDTO(
      String id,
      String nome
  ) {}

  public record MunicipioInfoDTO(
      String id,
      String codigo,
      String nome
  ) {}

  public record DistritoDTO(
      String id,
      String codigo,
      String nome
  ) {}

  public record SubdistritoDTO(
      String id,
      String codigo,
      String nome
  ) {}
}
