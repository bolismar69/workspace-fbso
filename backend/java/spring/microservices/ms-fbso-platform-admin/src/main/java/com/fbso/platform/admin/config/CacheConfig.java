package com.fbso.platform.admin.config;

import com.github.benmanes.caffeine.cache.Caffeine;
import org.springframework.cache.CacheManager;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.cache.caffeine.CaffeineCacheManager;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.concurrent.TimeUnit;

/**
 * Configuração de cache com Caffeine (alta performance, in-memory).
 * <p>
 * Usado pela Sprint 4 para:
 * <ul>
 *   <li>Cache da matriz de permissões RBAC (se necessário no futuro)</li>
 *   <li>Rate limiting por tenant</li>
 * </ul>
 * <p>
 * TTL padrão: 5 minutos. A matriz RBAC (32 linhas) atualmente usa
 * {@code findAll()} indexado (&lt;1ms) — cache não é necessário para
 * esse volume, mas a infraestrutura fica disponível para uso futuro.
 *
 * @see com.fbso.platform.admin.security.aspect.RbacAspect
 */
@Configuration
@EnableCaching
public class CacheConfig {

    /**
     * CacheManager baseado em Caffeine com TTL padrão de 5 minutos.
     * Caches específicos podem ser configurados com TTLs diferentes
     * via {@code Caffeine.newBuilder().expireAfterWrite(...)}.
     */
    @Bean
    public CacheManager cacheManager() {
        CaffeineCacheManager cacheManager = new CaffeineCacheManager();
        cacheManager.setCaffeine(Caffeine.newBuilder()
                .expireAfterWrite(5, TimeUnit.MINUTES)
                .maximumSize(1000));
        return cacheManager;
    }
}
