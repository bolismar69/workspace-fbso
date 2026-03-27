package com.fbso.geolocalidade.service;

import com.fbso.geolocalidade.dto.LocalidadeDetalhadaDTO;
import com.fbso.geolocalidade.dto.RespostaCompletaDTO;
import com.fbso.geolocalidade.dto.VizinhoEnriquecidoDTO;
import com.fbso.geolocalidade.exception.AwesomeApiException;
import com.fbso.geolocalidade.repository.MunicipioRepository;
import com.fbso.geolocalidade.repository.SubdistritoRepository;
import java.util.List;
import org.springframework.stereotype.Service;

@Service
public class LocalidadeService {

  private final AwesomeCepService awesomeService;
  private final MunicipioRepository municipioRepository;
  private final SubdistritoRepository subdistritoRepository;

  public LocalidadeService(
      AwesomeCepService awesomeService,
      MunicipioRepository municipioRepository,
      SubdistritoRepository subdistritoRepository) {
    this.awesomeService = awesomeService;
    this.municipioRepository = municipioRepository;
    this.subdistritoRepository = subdistritoRepository;
  }

  public RespostaCompletaDTO processarBuscaPorCepsProximos(String cep, Double raioKm) {
    var origemApi = awesomeService.obterCoordenadas(cep);

    if (origemApi == null) {
      throw new AwesomeApiException("AwesomeAPI retornou resposta vazia para o CEP informado");
    }
    if (origemApi.lat() == null || origemApi.lat().isBlank() || origemApi.lng() == null || origemApi.lng().isBlank()) {
      throw new AwesomeApiException("AwesomeAPI não retornou coordenadas válidas para o CEP informado");
    }

    var cepsVizinhos = awesomeService.buscarCepsVizinhosNoRaio(origemApi.lat(), origemApi.lng(), raioKm);

    List<VizinhoEnriquecidoDTO> cepsVizinhosEnriquecidos = cepsVizinhos.stream()
        .map(v -> {
        String cityIbge = v == null ? null : v.city_ibge();

        String subdistritoOficial = (cityIbge == null || cityIbge.isBlank())
          ? null
          : subdistritoRepository.findNomeByCodigo(cityIbge)
            .orElse(null); // Se não encontrar subdistrito, deixa como null para diferenciar "Sede" de "Não Informado"

          return new VizinhoEnriquecidoDTO(
              v.cep(),
              v.city(),
              v.city_ibge(),
              v.district(),
              subdistritoOficial,
              v.distance_km());
        })
        .toList();

    // LocalidadeDetalhadaDTO localidadeInfo = null;
    // if (origemApi.city_ibge() != null && !origemApi.city_ibge().isBlank()) {
    //   localidadeInfo = municipioRepository.findById(origemApi.city_ibge())
    //       .map(m -> new LocalidadeDetalhadaDTO(
    //           m.getCodigoIbge7(),
    //           m.getNomeMunicipio(),
    //           m.getUfSigla()
    //       ))
    //       .orElse(null);
    // }

    return new RespostaCompletaDTO(origemApi, cepsVizinhosEnriquecidos);
  }
}
