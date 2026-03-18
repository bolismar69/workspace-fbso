package com.fbso.geolocalidade.config;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "cache.awesome-cep")
public record CacheProperties(long ttlSeconds, long maxSize) {}
