package com.fbso.platform.admin.security;

import com.fbso.platform.admin.dto.response.ErrorResponse;
import com.fbso.platform.admin.utils.JwtUtils;
import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.oauth2.jwt.Jwt;
import org.springframework.security.oauth2.jwt.JwtDecoder;
import org.springframework.security.oauth2.jwt.JwtException;
import org.springframework.security.web.authentication.preauth.PreAuthenticatedAuthenticationToken;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.List;

/**
 * Filtro de autenticação JWT — executado em TODA requisição (exceto health check).
 * <p>
 * Pipeline:
 * <ol>
 *   <li>Extrai token do header {@code Authorization: Bearer <jwt>}</li>
 *   <li>Valida assinatura RS256 com chave pública do Keycloak (JWKS)</li>
 *   <li>Valida expiração ({@code exp})</li>
 *   <li>Extrai claims customizadas (tenant_id, user_id, roles, business_unit_ids, modules)</li>
 *   <li>Popula {@link TenantContext} e {@link SecurityContextHolder}</li>
 *   <li>Limpa {@link TenantContext} no {@code finally}</li>
 * </ol>
 *
 * @see <a href="ARCHITECTURE.md#3">ARCHITECTURE.md §3 — Pipeline de Segurança</a>
 */
@Component
public class JwtAuthenticationFilter extends OncePerRequestFilter {

    private static final Logger log = LoggerFactory.getLogger(JwtAuthenticationFilter.class);

    private static final String BEARER_PREFIX = "Bearer ";
    private static final String HEADER_AUTHORIZATION = "Authorization";

    private final JwtDecoder jwtDecoder;
    private final ObjectMapper objectMapper;

    public JwtAuthenticationFilter(JwtDecoder jwtDecoder, ObjectMapper objectMapper) {
        this.jwtDecoder = jwtDecoder;
        this.objectMapper = objectMapper;
    }

    @Override
    protected boolean shouldNotFilter(HttpServletRequest request) {
        String path = request.getRequestURI();
        // Health check do K8s não precisa de autenticação
        return path.equals("/actuator/health");
    }

    @Override
    protected void doFilterInternal(HttpServletRequest request,
                                     HttpServletResponse response,
                                     FilterChain filterChain)
            throws ServletException, IOException {

        String authHeader = request.getHeader(HEADER_AUTHORIZATION);

        // ---- Sem token → 401 ----
        if (authHeader == null || !authHeader.startsWith(BEARER_PREFIX)) {
            sendUnauthorized(response, "Token de acesso não informado");
            return;
        }

        String token = authHeader.substring(BEARER_PREFIX.length()).trim();

        try {
            // ---- Validar assinatura + expiração (Keycloak RS256 JWKS) ----
            Jwt jwt = jwtDecoder.decode(token);

            // ---- Extrair claims ----
            var tenantId = JwtUtils.getTenantId(jwt);
            var userId = JwtUtils.getUserId(jwt);
            var roles = JwtUtils.getRoles(jwt);
            var buIds = JwtUtils.getBusinessUnitIds(jwt);
            var modules = JwtUtils.getModules(jwt);

            // ---- Popular TenantContext (ThreadLocal) ----
            // O TenantAwareDataSource lê este valor e configura app.current_tenant_id
            // na sessão PostgreSQL automaticamente em cada getConnection() (ADR-L07)
            TenantContext.set(tenantId, userId, roles, buIds, modules);

            // ---- Popular SecurityContext ----
            var authorities = roles.stream()
                    .map(r -> new SimpleGrantedAuthority("ROLE_" + r))
                    .toList();

            PreAuthenticatedAuthenticationToken authToken =
                    new PreAuthenticatedAuthenticationToken(userId, jwt, authorities);
            SecurityContextHolder.getContext().setAuthentication(authToken);

            log.debug("JWT validado: tenant={}, user={}, roles={}", tenantId, userId, roles);

            // ---- Prosseguir na cadeia ----
            filterChain.doFilter(request, response);

        } catch (IllegalArgumentException e) {
            // Claims ausentes ou inválidas
            log.warn("JWT com claims inválidas: {}", e.getMessage());
            sendUnauthorized(response, "Token inválido: " + e.getMessage());
        } catch (JwtException e) {
            // Assinatura inválida, token expirado, alg=none, etc.
            log.warn("JWT inválido: {}", e.getMessage());
            sendUnauthorized(response, "Token inválido ou expirado");
        } finally {
            // ---- SEMPRE limpar ThreadLocal para evitar vazamento entre requisições ----
            // A variável de sessão PostgreSQL (app.current_tenant_id) é gerenciada
            // pelo TenantAwareDataSource em cada getConnection() — não requer cleanup aqui
            TenantContext.clear();
        }
    }

    /**
     * Retorna 401 Unauthorized com corpo JSON RFC 7807 (Problem Details).
     * Usa ObjectMapper para garantir escape JSON correto e consistência com o resto da API.
     */
    private void sendUnauthorized(HttpServletResponse response, String message) throws IOException {
        response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
        response.setContentType("application/json; charset=UTF-8");
        objectMapper.writeValue(response.getWriter(),
                ErrorResponse.of(
                        "https://api.fbso.org/errors/unauthorized",
                        message,
                        401,
                        null));
    }
}
