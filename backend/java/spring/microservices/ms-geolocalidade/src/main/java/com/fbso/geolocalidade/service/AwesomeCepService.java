package com.fbso.geolocalidade.service;

import com.fbso.geolocalidade.config.AwesomeApiProperties;
import com.fbso.geolocalidade.config.CacheConfig;
import com.fbso.geolocalidade.dto.AwesomeCepDTO;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.http.HttpHeaders;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;
import org.springframework.web.client.RestClientResponseException;

@Service
public class AwesomeCepService {

  private final RestClient restClient;
  private final AwesomeApiProperties props;

  public AwesomeCepService(RestClient restClient, AwesomeApiProperties props) {
    this.restClient = restClient;
    this.props = props;
  }

  @Cacheable(cacheNames = CacheConfig.AWESOME_CEP_CACHE, key = "T(com.fbso.geolocalidade.service.CepUtils).normalizeCep(#cep)")
  public AwesomeCepDTO obterCoordenadas(String cep) {
    String normalized = CepUtils.normalizeCep(cep);
    String baseUrl = props.baseUrlOrDefault();

    String url = baseUrl + "/json/" + normalized;
    if (props.token() != null && !props.token().isBlank()) {
      url = url + "?token=" + props.token();
    }

    try {
      return restClient.get().uri(url).retrieve().body(AwesomeCepDTO.class);
    } catch (RestClientResponseException ex) {
      throw new AwesomeApiException("Erro na geocodificação do CEP: status=" + ex.getStatusCode().value(), ex);
    } catch (Exception ex) {
      throw new AwesomeApiException("Erro na geocodificação do CEP: " + ex.getClass().getSimpleName(), ex);
    }
  }

  public List<AwesomeCepDTO> buscarVizinhosNoRaio(String lat, String lng, Double raioKm) {
    String baseUrl = props.baseUrlOrDefault();
    String url = baseUrl + "/search?lat=" + lat + "&lng=" + lng + "&d=" + raioKm;

    try {
      AwesomeCepDTO[] response = restClient.get()
          .uri(url)
          .header("x-api-key", props.key() == null ? "" : props.key())
          .retrieve()
          .body(AwesomeCepDTO[].class);

      if (response == null) {
        return Collections.emptyList();
      }
      return Arrays.asList(response);
    } catch (RestClientResponseException ex) {
      throw new AwesomeApiException("Erro na busca por raio: status=" + ex.getStatusCode().value(), ex);
    } catch (Exception ex) {
      throw new AwesomeApiException("Erro na busca por raio: " + ex.getClass().getSimpleName(), ex);
    }
  }
}
