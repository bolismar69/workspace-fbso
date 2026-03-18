package com.fbso.geolocalidade.config;

import com.github.benmanes.caffeine.cache.Caffeine;
import java.time.Duration;
import org.springframework.cache.CacheManager;
import org.springframework.cache.caffeine.CaffeineCacheManager;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.cache.annotation.EnableCaching;

@Configuration
@EnableCaching
public class CacheConfig {
  public static final String AWESOME_CEP_CACHE = "awesomeCep";

  @Bean
  public CacheManager cacheManager(CacheProperties props) {
    var cache = Caffeine.newBuilder()
        .maximumSize(Math.max(1, props.maxSize()))
        .expireAfterWrite(Duration.ofSeconds(Math.max(1, props.ttlSeconds())));

    var manager = new CaffeineCacheManager(AWESOME_CEP_CACHE);
    manager.setCaffeine(cache);
    return manager;
  }
}
