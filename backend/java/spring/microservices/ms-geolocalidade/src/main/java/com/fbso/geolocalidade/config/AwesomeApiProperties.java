package com.fbso.geolocalidade.config;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "awesomeapi")
public record AwesomeApiProperties(String token, String key, String baseUrl) {
  public String baseUrlOrDefault() {
    return (baseUrl == null || baseUrl.isBlank()) ? "https://cep.awesomeapi.com.br" : baseUrl;
  }
}
