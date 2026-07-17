package com.fbso.platform.admin.security;

import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.servlet.FilterChain;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.oauth2.jwt.Jwt;
import org.springframework.security.oauth2.jwt.JwtDecoder;
import org.springframework.security.oauth2.jwt.JwtException;

import java.io.PrintWriter;
import java.time.Instant;
import java.util.List;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("JwtAuthenticationFilter")
class JwtAuthenticationFilterTest {

    @Mock private JwtDecoder jwtDecoder;
    @Mock private HttpServletRequest request;
    @Mock private HttpServletResponse response;
    @Mock private FilterChain filterChain;
    @Mock private PrintWriter writer;

    private JwtAuthenticationFilter filter;
    private final ObjectMapper objectMapper = new ObjectMapper();

    @BeforeEach
    void setUp() {
        filter = new JwtAuthenticationFilter(jwtDecoder, objectMapper);
    }

    /** Configura o mock do writer apenas para testes que escrevem resposta de erro. */
    private void stubWriter() throws Exception {
        when(response.getWriter()).thenReturn(writer);
    }

    @AfterEach
    void tearDown() {
        // TenantContext é limpo no finally do filter
    }

    // ---- TC-S2-001: Sem header → 401 ----

    @Nested
    @DisplayName("TC-S2-001 — Sem token")
    class NoToken {

        @Test
        @DisplayName("deve retornar 401 quando header Authorization ausente")
        void shouldReturn401WhenNoAuthHeader() throws Exception {
            stubWriter();
            when(request.getHeader("Authorization")).thenReturn(null);

            filter.doFilterInternal(request, response, filterChain);

            verify(response).setStatus(401);
            verify(filterChain, never()).doFilter(request, response);
        }

        @Test
        @DisplayName("deve retornar 401 quando header não começa com Bearer")
        void shouldReturn401WhenNotBearer() throws Exception {
            stubWriter();
            when(request.getHeader("Authorization")).thenReturn("Basic dXNlcjpwYXNz");

            filter.doFilterInternal(request, response, filterChain);

            verify(response).setStatus(401);
            verify(filterChain, never()).doFilter(request, response);
        }
    }

    // ---- TC-S2-002: JWT assinatura inválida → 401 ----

    @Nested
    @DisplayName("TC-S2-002 — Token inválido")
    class InvalidToken {

        @Test
        @DisplayName("deve retornar 401 quando assinatura é inválida")
        void shouldReturn401WhenSignatureInvalid() throws Exception {
            stubWriter();
            when(request.getHeader("Authorization")).thenReturn("Bearer invalid.token.here");
            when(jwtDecoder.decode(anyString())).thenThrow(new JwtException("Invalid signature"));

            filter.doFilterInternal(request, response, filterChain);

            verify(response).setStatus(401);
        }
    }

    // ---- TC-S2-003: JWT expirado → 401 ----

    @Nested
    @DisplayName("TC-S2-003 — Token expirado")
    class ExpiredToken {

        @Test
        @DisplayName("deve retornar 401 quando token está expirado")
        void shouldReturn401WhenTokenExpired() throws Exception {
            stubWriter();
            when(request.getHeader("Authorization")).thenReturn("Bearer expired.token.here");
            when(jwtDecoder.decode(anyString())).thenThrow(new JwtException("Token expired"));

            filter.doFilterInternal(request, response, filterChain);

            verify(response).setStatus(401);
        }
    }

    // ---- TC-S2-005: JWT válido → extrai claims ----

    @Nested
    @DisplayName("TC-S2-005 — Token válido")
    class ValidToken {

        @Test
        @DisplayName("deve extrair claims e prosseguir na cadeia")
        void shouldExtractClaimsAndProceed() throws Exception {
            when(request.getHeader("Authorization"))
                    .thenReturn("Bearer valid.jwt.token");

            Jwt jwt = Jwt.withTokenValue("valid.jwt.token")
                    .header("alg", "RS256")
                    .claim("tenant_id", "00000000-0000-0000-0000-000000000001")
                    .claim("roles", List.of("ADMIN_TENANT"))
                    .claim("business_unit_ids", List.of())
                    .claim("modules", List.of("FBSO Platform"))
                    .subject("00000000-0000-0000-0000-0000000000AA")
                    .issuedAt(Instant.now())
                    .expiresAt(Instant.now().plusSeconds(3600))
                    .build();

            when(jwtDecoder.decode("valid.jwt.token")).thenReturn(jwt);

            filter.doFilterInternal(request, response, filterChain);

            verify(filterChain).doFilter(request, response);
            verify(response, never()).setStatus(any(Integer.class));
        }
    }

    // ---- TC-S2-006: Health check não passa pelo filter ----

    @Nested
    @DisplayName("TC-S2-006 — Health check")
    class HealthCheck {

        @Test
        @DisplayName("não deve filtrar /actuator/health")
        void shouldNotFilterHealthEndpoint() {
            when(request.getRequestURI()).thenReturn("/actuator/health");

            boolean skip = filter.shouldNotFilter(request);

            org.assertj.core.api.Assertions.assertThat(skip).isTrue();
        }
    }
}
