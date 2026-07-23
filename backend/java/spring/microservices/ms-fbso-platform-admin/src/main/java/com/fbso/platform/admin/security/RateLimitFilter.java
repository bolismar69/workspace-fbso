package com.fbso.platform.admin.security;

import com.fbso.platform.admin.dto.response.ErrorResponse;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.benmanes.caffeine.cache.Cache;
import com.github.benmanes.caffeine.cache.Caffeine;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.time.Duration;
import java.time.Instant;

/**
 * Rate limiting para tentativas de login — Filter + Caffeine Cache.
 * <p>
 * Política: 5 tentativas de login com falha → bloqueio de 15 minutos.
 * A chave do cache é o username (extraído do request body) ou o IP
 * remoto como fallback. Aplica-se apenas a {@code POST /api/v1/auth/login}.
 * <p>
 * <b>Decisão de design (DT-110):</b> Filter (não @Aspect).
 * Filter é mais idiomático no Spring Security para preocupações de
 * infraestrutura. Caffeine é local (por instância) — adequado para
 * Fase 0 single-instance. Trigger para Redis: {@code INSTANCE_COUNT > 1}.
 *
 * @see <a href="ARCHITECTURE.md#4">ARCHITECTURE.md §4 — Pipeline de Segurança</a>
 */
public class RateLimitFilter extends OncePerRequestFilter {

    private static final Logger log = LoggerFactory.getLogger(RateLimitFilter.class);

    static final int MAX_ATTEMPTS = 5;
    static final Duration BLOCK_DURATION = Duration.ofMinutes(15);

    private final Cache<String, RateLimitEntry> cache;
    private final ObjectMapper objectMapper;

    public RateLimitFilter(ObjectMapper objectMapper) {
        this.cache = Caffeine.newBuilder()
                .expireAfterWrite(BLOCK_DURATION)
                .maximumSize(10_000)
                .build();
        this.objectMapper = objectMapper;
    }

    @Override
    protected void doFilterInternal(HttpServletRequest request,
                                    HttpServletResponse response,
                                    FilterChain filterChain)
            throws ServletException, IOException {

        // Aplica apenas a POST /api/v1/auth/login
        if (!isLoginEndpoint(request)) {
            filterChain.doFilter(request, response);
            return;
        }

        String key = extractKey(request);
        RateLimitEntry entry = cache.getIfPresent(key);

        if (entry != null && entry.isBlocked()) {
            long secondsLeft = entry.secondsUntilUnblock();
            log.warn("Rate limit: key={} bloqueada por mais {}s", key, secondsLeft);
            sendRateLimitResponse(response, secondsLeft);
            return;
        }

        // Continua a cadeia — o filtro de autenticação registrará falha/sucesso
        filterChain.doFilter(request, response);

        // Após a resposta: se foi 401, incrementa o contador
        if (response.getStatus() == HttpStatus.UNAUTHORIZED.value()) {
            recordFailedAttempt(key);
        } else if (response.getStatus() == HttpStatus.OK.value()) {
            // Login bem-sucedido: limpa o contador
            cache.invalidate(key);
        }
    }

    /**
     * Registra uma tentativa de login com falha.
     * Na 5ª falha, bloqueia por 15 minutos.
     */
    void recordFailedAttempt(String key) {
        RateLimitEntry entry = cache.get(key, k -> new RateLimitEntry());
        entry.incrementFailures();
        cache.put(key, entry);

        if (entry.isBlocked()) {
            log.warn("Rate limit: key={} bloqueada por {} tentativas falhas ({}min)",
                    key, MAX_ATTEMPTS, BLOCK_DURATION.toMinutes());
        }
    }

    /**
     * Responde com 429 Too Many Requests — RFC 7807.
     */
    private void sendRateLimitResponse(HttpServletResponse response, long secondsLeft)
            throws IOException {
        response.setStatus(HttpStatus.TOO_MANY_REQUESTS.value());
        response.setContentType(MediaType.APPLICATION_JSON_VALUE);

        ErrorResponse error = ErrorResponse.of(
                "https://api.fbso.org/errors/rate-limit",
                "Muitas tentativas de login",
                429,
                "Conta bloqueada por " + BLOCK_DURATION.toMinutes()
                        + " minutos. Tente novamente em " + secondsLeft + " segundos."
        );

        objectMapper.writeValue(response.getWriter(), error);
    }

    private boolean isLoginEndpoint(HttpServletRequest request) {
        return "POST".equalsIgnoreCase(request.getMethod())
                && request.getRequestURI().equals("/api/v1/auth/login");
    }

    /**
     * Extrai a chave de rate limiting: IP remoto da requisição.
     * <p>
     * Nota: X-Forwarded-For NÃO é utilizado por risco de spoofing.
     * Quando houver proxy reverso confiável (ex: Nginx com set_real_ip_from),
     * usar o IP do proxy configurado em application.yml.
     * // ponytail: externalizar trusted-proxy-ips para application.yml na Sprint 6
     */
    private String extractKey(HttpServletRequest request) {
        return request.getRemoteAddr();
    }

    // -- Inner class --

    /**
     * Entrada no cache de rate limiting.
     */
    static class RateLimitEntry {
        private int failures;
        private Instant blockedUntil;

        void incrementFailures() {
            failures++;
            if (failures >= MAX_ATTEMPTS) {
                blockedUntil = Instant.now().plus(BLOCK_DURATION);
            }
        }

        boolean isBlocked() {
            return blockedUntil != null && Instant.now().isBefore(blockedUntil);
        }

        long secondsUntilUnblock() {
            if (blockedUntil == null) return 0;
            return Math.max(0, Duration.between(Instant.now(), blockedUntil).getSeconds());
        }
    }
}
