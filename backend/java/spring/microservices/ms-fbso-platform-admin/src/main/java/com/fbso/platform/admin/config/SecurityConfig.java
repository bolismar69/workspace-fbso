package com.fbso.platform.admin.config;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fbso.platform.admin.security.FbsoJwtAuthenticationConverter;
import com.fbso.platform.admin.security.JwtAuthenticationFilter;
import com.fbso.platform.admin.security.RateLimitFilter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.annotation.Order;
import org.springframework.security.config.annotation.method.configuration.EnableMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.oauth2.jwt.JwtDecoder;
import org.springframework.security.oauth2.jwt.NimbusJwtDecoder;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.CorsConfigurationSource;
import org.springframework.web.cors.UrlBasedCorsConfigurationSource;

import java.util.Arrays;
import java.util.List;

/**
 * Configuração central de segurança do microsserviço.
 * <p>
 * Dois {@link SecurityFilterChain} beans com ordens distintas:
 * <ol>
 *   <li>{@code oauth2LoginFilterChain} ({@code @Order(1)}) — OAuth2 Client
 *       Authorization Code Flow para login (Keycloak). Aplica-se apenas a
 *       {@code /auth/**, /login, /oauth2/**}.</li>
 *   <li>{@code apiFilterChain} ({@code @Order(2)}) — Resource Server JWT
 *       para APIs REST. Stateless, sem sessão, sem CSRF.</li>
 * </ol>
 * <p>
 * Pipeline da API (stateless):
 * <ol>
 *   <li>JwtAuthenticationFilter — extrai e valida JWT (Keycloak RS256)</li>
 *   <li>RbacAspect — verifica @RequiresPermission contra matriz RN10-01</li>
 *   <li>TenantIsolationAspect — injeta WHERE tenant_id = ?</li>
 *   <li>AuditAspect — registra @Auditable de forma assíncrona</li>
 * </ol>
 *
 * @see <a href="ARCHITECTURE.md#3">ARCHITECTURE.md §3 — Pipeline de Segurança</a>
 */
@Configuration
@EnableWebSecurity
@EnableMethodSecurity
public class SecurityConfig {

    private final JwtAuthenticationFilter jwtAuthenticationFilter;
    private final ObjectMapper objectMapper;

    @Value("${spring.security.oauth2.resourceserver.jwt.jwk-set-uri}")
    private String jwkSetUri;

    @Value("${spring.security.oauth2.resourceserver.jwt.issuer-uri}")
    private String issuerUri;

    @Value("${app.cors.allowed-origins:http://localhost:3000,https://app.fbso.org}")
    private String allowedOrigins;

    @Value("${app.rate-limit.trusted-proxy-ips:127.0.0.1,0:0:0:0:0:0:0:1}")
    private String trustedProxyIpsCsv;

    public SecurityConfig(JwtAuthenticationFilter jwtAuthenticationFilter,
                          ObjectMapper objectMapper) {
        this.jwtAuthenticationFilter = jwtAuthenticationFilter;
        this.objectMapper = objectMapper;
    }

    /**
     * Filter chain para OAuth2 Login (Authorization Code Flow).
     * <p>
     * Aplica-se apenas a {@code /auth/**, /login, /oauth2/**}.
     * Usa sessão para manter o estado do fluxo OAuth2 (necessário).
     * Após login bem-sucedido, o frontend recebe o token JWT e o utiliza
     * nas chamadas à API (Resource Server, stateless).
     */
    @Bean
    @Order(1)
    public SecurityFilterChain oauth2LoginFilterChain(HttpSecurity http) throws Exception {
        http
            .securityMatcher("/auth/**", "/login", "/oauth2/**")
            .authorizeHttpRequests(auth -> auth
                .anyRequest().authenticated()
            )
            .oauth2Login(oauth2 -> oauth2
                .defaultSuccessUrl("/api/v1/auth/me", true)
            )
            .csrf(csrf -> csrf.disable())
            .sessionManagement(session -> session
                .sessionCreationPolicy(SessionCreationPolicy.IF_REQUIRED)
            );

        return http.build();
    }

    /**
     * Filter chain para API REST — Resource Server JWT (stateless).
     * <p>
     * Aplica-se a todos os endpoints exceto os capturados pelo
     * {@code oauth2LoginFilterChain} (que tem precedência via {@code @Order(1)}).
     */
    @Bean
    @Order(2)
    public SecurityFilterChain apiFilterChain(HttpSecurity http) throws Exception {
        http
            // ---- JWT Resource Server (Keycloak RS256) ----
            // O FbsoJwtAuthenticationConverter extrai claims customizadas
            // (tenant_id, roles, business_unit_ids, modules) durante a
            // decodificação do Resource Server, eliminando dupla decodificação
            // (DT-076/DT-102).
            .oauth2ResourceServer(oauth2 -> oauth2
                .jwt(jwt -> jwt
                    .decoder(jwtDecoder())
                    .jwtAuthenticationConverter(fbsoJwtAuthenticationConverter())
                )
            )

            // ---- Rate limiting (antes da autenticação) ----
            .addFilterBefore(rateLimitFilter(), JwtAuthenticationFilter.class)

            // ---- Injetar filtro customizado ANTES do UsernamePasswordAuthenticationFilter ----
            .addFilterBefore(jwtAuthenticationFilter, UsernamePasswordAuthenticationFilter.class)

            // ---- Stateless (API REST — sem sessão HTTP) ----
            .sessionManagement(session -> session
                .sessionCreationPolicy(SessionCreationPolicy.STATELESS)
            )

            // ---- CSRF desabilitado (JWT substitui) ----
            .csrf(csrf -> csrf.disable())

            // ---- CORS configurado para o frontend ----
            .cors(cors -> cors.configurationSource(corsConfigurationSource()))

            // ---- Regras de acesso ----
            .authorizeHttpRequests(auth -> auth
                // Health check público (K8s liveness/readiness)
                .requestMatchers("/actuator/health").permitAll()
                // Demais endpoints exigem autenticação
                .anyRequest().authenticated()
            );

        return http.build();
    }

    /**
     * Decodificador JWT — valida assinatura RS256 com chave pública do Keycloak.
     * <p>
     * A URI do JWKS é configurada em {@code application.yml}
     * ({@code spring.security.oauth2.resourceserver.jwt.jwk-set-uri}).
     */
    @Bean
    public JwtDecoder jwtDecoder() {
        NimbusJwtDecoder decoder = NimbusJwtDecoder
                .withJwkSetUri(jwkSetUri)
                .build();
        decoder.setJwtValidator(
                org.springframework.security.oauth2.jwt.JwtValidators
                        .createDefaultWithIssuer(issuerUri));
        return decoder;
    }

    /**
     * Rate limiting para tentativas de login (5 tentativas → 15min bloqueio).
     * <p>
     * Usa Caffeine Cache local. Single-instance na Fase 0.
     * Migrar para Redis quando {@code INSTANCE_COUNT > 1}.
     */
    /**
     * Converter JWT customizado — extrai claims (tenant_id, roles, BU ids, modules)
     * durante a decodificação do Resource Server. Elimina dupla decodificação.
     */
    @Bean
    public FbsoJwtAuthenticationConverter fbsoJwtAuthenticationConverter() {
        return new FbsoJwtAuthenticationConverter();
    }

    @Bean
    public RateLimitFilter rateLimitFilter() {
        List<String> trustedProxyIps = trustedProxyIpsCsv.isBlank()
                ? List.of()
                : Arrays.asList(trustedProxyIpsCsv.split(","));
        return new RateLimitFilter(objectMapper, trustedProxyIps);
    }

    /**
     * Configuração CORS — apenas a origem do frontend pode chamar a API.
     */
    @Bean
    public CorsConfigurationSource corsConfigurationSource() {
        CorsConfiguration configuration = new CorsConfiguration();
        configuration.setAllowedOrigins(Arrays.asList(allowedOrigins.split(",")));
        configuration.setAllowedMethods(List.of("GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"));
        configuration.setAllowedHeaders(List.of("Authorization", "Content-Type", "Accept"));
        configuration.setAllowCredentials(true);
        configuration.setMaxAge(3600L);

        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/api/**", configuration);
        return source;
    }
}
