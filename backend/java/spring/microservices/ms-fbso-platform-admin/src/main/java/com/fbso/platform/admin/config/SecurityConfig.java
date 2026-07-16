package com.fbso.platform.admin.config;

import com.fbso.platform.admin.security.JwtAuthenticationFilter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
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

import java.util.List;

/**
 * Configuração central de segurança do microsserviço.
 * <p>
 * Pipeline por requisição:
 * <ol>
 *   <li>JwtAuthenticationFilter — extrai e valida JWT (Keycloak RS256)</li>
 *   <li>RbacAspect — verifica @RequiresPermission contra matriz RN10-01</li>
 *   <li>TenantIsolationAspect — injeta WHERE tenant_id = ?</li>
 *   <li>AuditAspect — registra @Auditable de forma assíncrona</li>
 * </ol>
 * <p>
 * API stateless — sem sessão, sem CSRF. JWT substitui ambas.
 *
 * @see <a href="ARCHITECTURE.md#3">ARCHITECTURE.md §3 — Pipeline de Segurança</a>
 */
@Configuration
@EnableWebSecurity
@EnableMethodSecurity
public class SecurityConfig {

    private final JwtAuthenticationFilter jwtAuthenticationFilter;

    @Value("${spring.security.oauth2.resourceserver.jwt.jwk-set-uri}")
    private String jwkSetUri;

    public SecurityConfig(JwtAuthenticationFilter jwtAuthenticationFilter) {
        this.jwtAuthenticationFilter = jwtAuthenticationFilter;
    }

    /**
     * Configura a cadeia de filtros de segurança.
     */
    @Bean
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {
        http
            // ---- JWT Resource Server (Keycloak RS256) ----
            .oauth2ResourceServer(oauth2 -> oauth2
                .jwt(jwt -> jwt.decoder(jwtDecoder()))
            )

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
        return NimbusJwtDecoder
                .withJwkSetUri(jwkSetUri)
                .build();
    }

    /**
     * Configuração CORS — apenas a origem do frontend pode chamar a API.
     */
    @Bean
    public CorsConfigurationSource corsConfigurationSource() {
        CorsConfiguration configuration = new CorsConfiguration();
        configuration.setAllowedOrigins(List.of(
            "http://localhost:3000",   // dev frontend
            "https://app.fbso.org"     // prod frontend
        ));
        configuration.setAllowedMethods(List.of("GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"));
        configuration.setAllowedHeaders(List.of("Authorization", "Content-Type", "Accept"));
        configuration.setAllowCredentials(true);
        configuration.setMaxAge(3600L);

        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/api/**", configuration);
        return source;
    }
}
