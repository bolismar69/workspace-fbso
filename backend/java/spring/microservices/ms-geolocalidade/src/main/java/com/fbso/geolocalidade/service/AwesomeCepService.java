package com.fbso.geolocalidade.service;

import com.fbso.geolocalidade.config.AwesomeApiProperties;
import com.fbso.geolocalidade.config.CacheConfig;
import com.fbso.geolocalidade.dto.AwesomeCepDTO;
import com.fbso.geolocalidade.dto.AwesomeSearchResponseDTO;
import com.fbso.geolocalidade.exception.AwesomeApiException;
import com.fbso.geolocalidade.utils.CepUtils;
import java.util.Collections;
import java.util.List;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;
import org.springframework.web.client.RestClientResponseException;

@Service
public class AwesomeCepService {

  private static final Logger log = LoggerFactory.getLogger(AwesomeCepService.class);

  private final RestClient restClient;
  private final AwesomeApiProperties props;

  public AwesomeCepService(RestClient restClient, AwesomeApiProperties props) {
    this.restClient = restClient;
    this.props = props;
  }

  @Cacheable(cacheNames = CacheConfig.AWESOME_CEP_CACHE, key = "T(com.fbso.geolocalidade.utils.CepUtils).normalizeCep(#cep)")
  public AwesomeCepDTO obterCoordenadas(String cep) {
    String normalized = CepUtils.normalizeCep(cep);
    String baseUrl = props.baseUrlOrDefault();

    String url = baseUrl + "/json/" + normalized;
    if (props.token() != null && !props.token().isBlank()) {
      url = url + "?token=" + props.token();
    }

    long startNs = System.nanoTime();
    log.info("AwesomeAPI request: GET /json/{{cep}} cep={} url={}", normalized, sanitizeUrl(url));
    try {
      ResponseEntity<AwesomeCepDTO> entity = restClient.get().uri(url).retrieve().toEntity(AwesomeCepDTO.class);
      long tookMs = tookMs(startNs);
      AwesomeCepDTO body = entity.getBody();
      log.info(
          "AwesomeAPI response: GET /json/{{cep}} status={} tookMs={} city_ibge={} lat={} lng={}",
          entity.getStatusCode().value(),
          tookMs,
          body == null ? null : body.city_ibge(),
          body == null ? null : body.lat(),
          body == null ? null : body.lng());
      return body;
    } catch (RestClientResponseException ex) {
      long tookMs = tookMs(startNs);
      log.warn(
          "AwesomeAPI error: GET /json/{{cep}} status={} tookMs={} url={} body={}",
          ex.getStatusCode().value(),
          tookMs,
          sanitizeUrl(url),
          truncate(ex.getResponseBodyAsString(), 800));
      throw new AwesomeApiException("Erro na geocodificação do CEP: status=" + ex.getStatusCode().value(), ex);
    } catch (Exception ex) {
      long tookMs = tookMs(startNs);
      log.warn(
          "AwesomeAPI error: GET /json/{{cep}} tookMs={} url={} exType={} msg={}",
          tookMs,
          sanitizeUrl(url),
          ex.getClass().getSimpleName(),
          truncate(ex.getMessage(), 800));
      throw new AwesomeApiException("Erro na geocodificação do CEP: " + ex.getClass().getSimpleName(), ex);
    }
  }

  public List<AwesomeCepDTO> buscarCepsVizinhosNoRaio(String lat, String lng, Double raioKm) {
    if ((props.token() == null || props.token().isBlank()) && (props.key() == null || props.key().isBlank())) {
      throw new AwesomeApiException(
          "AwesomeAPI: credenciais não configuradas (defina AWESOME_API_TOKEN ou AWESOME_API_KEY)");
    }

    String baseUrl = props.baseUrlOrDefault();
    String url = baseUrl + "/search?lat=" + lat + "&lng=" + lng + "&d=" + raioKm;

    // A AwesomeAPI aceita autenticação por token (query param) OU x-api-key (header).
    // Preferimos token, pois é o mesmo padrão usado no endpoint /json/{cep}.
    if (props.token() != null && !props.token().isBlank()) {
      url = url + "&token=" + props.token();
    }

    String authMode = (props.token() != null && !props.token().isBlank()) ? "token(query)" : "x-api-key(header)";
    long startNs = System.nanoTime();
    log.info(
        "AwesomeAPI request: GET /search lat={} lng={} d={} authMode={} url={}",
        lat,
        lng,
        raioKm,
        authMode,
        sanitizeUrl(url));
    try {
      var request = restClient.get().uri(url);

      // Só envia x-api-key quando NÃO há token e a key está configurada.
      if ((props.token() == null || props.token().isBlank()) && props.key() != null && !props.key().isBlank()) {
        request = request.header("x-api-key", props.key());
      }

      ResponseEntity<AwesomeSearchResponseDTO> entity = request.retrieve().toEntity(AwesomeSearchResponseDTO.class);
      long tookMs = tookMs(startNs);
      AwesomeSearchResponseDTO response = entity.getBody();

      log.info(
          "AwesomeAPI response: GET /search status={} tookMs={} found={} resultsSize={}",
          entity.getStatusCode().value(),
          tookMs,
          response == null ? null : response.found(),
          (response == null || response.results() == null) ? 0 : response.results().size());

      if (response == null) {
        return Collections.emptyList();
      }
      if (response.results() == null) {
        return Collections.emptyList();
      }
      return response.results();
    } catch (RestClientResponseException ex) {
      long tookMs = tookMs(startNs);
      log.warn(
          "AwesomeAPI error: GET /search status={} tookMs={} url={} body={}",
          ex.getStatusCode().value(),
          tookMs,
          sanitizeUrl(url),
          truncate(ex.getResponseBodyAsString(), 1200));
      throw new AwesomeApiException("Erro na busca por raio: status=" + ex.getStatusCode().value(), ex);
    } catch (Exception ex) {
      long tookMs = tookMs(startNs);
      String msg = ex.getMessage();
      log.warn(
          "AwesomeAPI error: GET /search tookMs={} url={} exType={} msg={}",
          tookMs,
          sanitizeUrl(url),
          ex.getClass().getSimpleName(),
          truncate(msg, 1200));
      throw new AwesomeApiException(
          "Erro na busca por raio: " + ex.getClass().getSimpleName() + (msg == null || msg.isBlank() ? "" : " - " + msg),
          ex);
    }
  }

  private static long tookMs(long startNs) {
    return (System.nanoTime() - startNs) / 1_000_000;
  }

  private static String sanitizeUrl(String url) {
    if (url == null) {
      return null;
    }
    // Remove token from query string to avoid leaking credentials.
    return url.replaceAll("([?&])token=[^&]+", "$1token=<redacted>");
  }

  private static String truncate(String value, int maxLen) {
    if (value == null) {
      return null;
    }
    if (value.length() <= maxLen) {
      return value;
    }
    return value.substring(0, Math.max(0, maxLen - 3)) + "...";
  }
}
