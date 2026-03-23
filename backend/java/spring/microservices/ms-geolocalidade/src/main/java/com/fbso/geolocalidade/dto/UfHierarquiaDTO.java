package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fbso.geolocalidade.dto.MunicipioDTO.DistritoDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.SubdistritoDTO;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@JsonInclude(JsonInclude.Include.NON_EMPTY)
public record UfHierarquiaDTO(
    List<UfDTO> uf
) {

  public static UfHierarquiaDTO fromMunicipios(List<MunicipioDTO> municipios) {
    Map<String, UfDTO> ufById = new LinkedHashMap<>();
    Map<String, RegiaoIntermediariaDTO> riByKey = new LinkedHashMap<>();
    Map<String, RegiaoImediataDTO> rImByKey = new LinkedHashMap<>();
    Map<String, MunicipioHierarquiaDTO> municipioByKey = new LinkedHashMap<>();

    for (MunicipioDTO m : municipios) {
      if (m == null || m.ufId() == null) {
        continue;
      }

      UfDTO ufNode = ufById.computeIfAbsent(
          m.ufId(),
          id -> new UfDTO(id, m.ufSigla(), m.ufNome(), new ArrayList<>()));

      String riKey = m.ufId() + ":" + nullToEmpty(m.regiaoIntermediariaId());
      RegiaoIntermediariaDTO riNode = riByKey.computeIfAbsent(
          riKey,
          k -> {
            RegiaoIntermediariaDTO node = new RegiaoIntermediariaDTO(
                m.regiaoIntermediariaId(),
                m.regiaoIntermediariaNome(),
                new ArrayList<>());
            ufNode.regiaoIntermediaria().add(node);
            return node;
          });

      String rImKey = riKey + ":" + nullToEmpty(m.regiaoImediataId());
      RegiaoImediataDTO rImNode = rImByKey.computeIfAbsent(
          rImKey,
          k -> {
            RegiaoImediataDTO node = new RegiaoImediataDTO(
                m.regiaoImediataId(),
                m.regiaoImediataNome(),
                new ArrayList<>());
            riNode.regiaoImediata().add(node);
            return node;
          });

      String municipioKey = rImKey + ":" + nullToEmpty(m.municipioId());
      MunicipioHierarquiaDTO municipioNode = municipioByKey.computeIfAbsent(
          municipioKey,
          k -> {
            MunicipioHierarquiaDTO node = new MunicipioHierarquiaDTO(
                m.municipioId(),
                m.municipioCodigo(),
                m.municipioNome(),
                new ArrayList<>());
            rImNode.municipio().add(node);
            return node;
          });

      List<DistritoHierarquiaDTO> distritos = new ArrayList<>();
      if (m.distrito() != null) {
        for (DistritoDTO d : m.distrito()) {
          if (d == null) {
            continue;
          }

          List<SubDistritoDTO> subDistritos = new ArrayList<>();
          if (d.subDistrito() != null) {
            for (SubdistritoDTO s : d.subDistrito()) {
              if (s == null) {
                continue;
              }
              subDistritos.add(new SubDistritoDTO(s.id(), s.codigo(), s.nome()));
            }
          }

          distritos.add(new DistritoHierarquiaDTO(d.id(), d.codigo(), d.nome(), subDistritos));
        }
      }

      if (!distritos.isEmpty()) {
        municipioNode.distrito().addAll(distritos);
      }
    }

    return new UfHierarquiaDTO(new ArrayList<>(ufById.values()));
  }

  public static UfHierarquiaDTO fromUfs(List<UfDTO> ufs) {
    return new UfHierarquiaDTO(ufs == null ? List.of() : ufs);
  }

  private static String nullToEmpty(String v) {
    return v == null ? "" : v;
  }

  @JsonInclude(JsonInclude.Include.NON_EMPTY)
  public record UfDTO(
      String id,
      String sigla,
      String nome,
      List<RegiaoIntermediariaDTO> regiaoIntermediaria
  ) {}

  @JsonInclude(JsonInclude.Include.NON_EMPTY)
  public record RegiaoIntermediariaDTO(
      String id,
      String nome,
      List<RegiaoImediataDTO> regiaoImediata
  ) {}

  @JsonInclude(JsonInclude.Include.NON_EMPTY)
  public record RegiaoImediataDTO(
      String id,
      String nome,
      List<MunicipioHierarquiaDTO> municipio
  ) {}

  @JsonInclude(JsonInclude.Include.NON_EMPTY)
  public record MunicipioHierarquiaDTO(
      String id,
      String codigo,
      String nome,
      List<DistritoHierarquiaDTO> distrito
  ) {}

  @JsonInclude(JsonInclude.Include.NON_EMPTY)
  public record DistritoHierarquiaDTO(
      String id,
      String codigo,
      String nome,
      List<SubDistritoDTO> subDistrito
  ) {}

  public record SubDistritoDTO(
      String id,
      String codigo,
      String nome
  ) {}
}
