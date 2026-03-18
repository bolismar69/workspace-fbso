package com.fbso.geolocalidade.service;

import com.fbso.geolocalidade.dto.LocalidadeDetalhadaDTO;
import com.fbso.geolocalidade.dto.RespostaCompletaDTO;
import com.fbso.geolocalidade.dto.VizinhoEnriquecidoDTO;
import com.fbso.geolocalidade.repository.MunicipioRepository;
import com.fbso.geolocalidade.repository.SubdistritoRepository;
import java.util.List;
import org.springframework.stereotype.Service;

@Service
public class LocalidadeOrchestrator {

  private final AwesomeCepService awesomeService;
  private final MunicipioRepository municipioRepository;
  private final SubdistritoRepository subdistritoRepository;

  public LocalidadeOrchestrator(
      AwesomeCepService awesomeService,
      MunicipioRepository municipioRepository,
      SubdistritoRepository subdistritoRepository) {
    this.awesomeService = awesomeService;
    this.municipioRepository = municipioRepository;
    this.subdistritoRepository = subdistritoRepository;
  }

  public RespostaCompletaDTO processarBuscaPorCep(String cep, Double raioKm) {
    var origemApi = awesomeService.obterCoordenadas(cep);
    var vizinhosApi = awesomeService.buscarVizinhosNoRaio(origemApi.lat(), origemApi.lng(), raioKm);

    List<VizinhoEnriquecidoDTO> vizinhosEnriquecidos = vizinhosApi.stream()
        .map(v -> {
          String subdistritoOficial = subdistritoRepository
              .findNomeByCodigo(v.city_ibge() == null ? "" : v.city_ibge())
              .orElse("Sede / Não Informado");

          return new VizinhoEnriquecidoDTO(
              v.cep(),
              v.city(),
              v.city_ibge(),
              v.district(),
              subdistritoOficial,
              parseDoubleOrNull(v.d()));
        })
        .toList();

    LocalidadeDetalhadaDTO localidadeInfo = null;
    if (origemApi.city_ibge() != null && !origemApi.city_ibge().isBlank()) {
      localidadeInfo = municipioRepository.findById(origemApi.city_ibge())
          .map(m -> new LocalidadeDetalhadaDTO(
              m.getCodigoIbge7(),
              m.getNomeMunicipio(),
              m.getUfSigla(),
              origemApi.lat(),
              origemApi.lng()))
          .orElse(null);
    }

    return new RespostaCompletaDTO(origemApi, localidadeInfo, vizinhosEnriquecidos);
  }

  private static Double parseDoubleOrNull(String value) {
    if (value == null || value.isBlank()) {
      return null;
    }
    try {
      return Double.parseDouble(value);
    } catch (NumberFormatException ex) {
      return null;
    }
  }
}
