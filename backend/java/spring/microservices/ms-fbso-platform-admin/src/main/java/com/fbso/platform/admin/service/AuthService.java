package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.ForgotPasswordRequest;
import com.fbso.platform.admin.dto.request.ResetPasswordRequest;
import com.fbso.platform.admin.dto.response.AuthResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import com.github.benmanes.caffeine.cache.Cache;
import com.github.benmanes.caffeine.cache.Caffeine;

import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.UUID;

/**
 * Serviço de autenticação — delega fluxos OAuth2 ao Keycloak.
 * <p>
 * Login: redireciona para Authorization Code Flow (OAuth2 Client no SecurityConfig).
 * Forgot/Reset Password: proxy para Keycloak Admin REST API.
 */
@Service
public class AuthService {

    private static final Logger log = LoggerFactory.getLogger(AuthService.class);

    private final RestTemplate restTemplate;

    /** Cache de tokens de reset: token → userId. TTL = 1h (RN13-03). */
    private final Cache<String, UUID> resetTokenCache = Caffeine.newBuilder()
            .expireAfterWrite(Duration.ofHours(1))
            .maximumSize(10_000)
            .build();

    @Value("${spring.security.oauth2.client.provider.keycloak.issuer-uri:http://localhost:8081/realms/fbso-platform}")
    private String issuerUri;

    @Value("${spring.security.oauth2.client.registration.keycloak.client-secret:changeme}")
    private String clientSecret;

    public AuthService() {
        this.restTemplate = new RestTemplate();
    }

    /**
     * Login — redireciona para Keycloak Authorization Code Flow.
     * O fluxo OAuth2 é gerenciado pelo Spring Security OAuth2 Client.
     */
    public AuthResponse login() {
        String redirectUrl = "/oauth2/authorization/keycloak";
        log.info("Auth: redirecionando para Keycloak OIDC — {}", redirectUrl);
        return AuthResponse.loginRedirect(redirectUrl);
    }

    /**
     * Forgot Password — envia email com link de reset via Keycloak Admin API.
     * <p>
     * RN13-03: link de reset expira em 1h (configurado no realm: actionTokenGeneratedByUserLifespan = 3600).
     */
    public AuthResponse forgotPassword(ForgotPasswordRequest request) {
        String email = request.email();
        log.info("Auth: solicitação de reset de senha para email={}", maskEmail(email));

        try {
            String adminToken = getAdminToken();
            String userId = findUserByEmail(adminToken, email);

            if (userId != null) {
                // Gerar token de reset aleatório (UUID v4) e armazenar no cache com TTL 1h
                String resetToken = UUID.randomUUID().toString();
                resetTokenCache.put(resetToken, UUID.fromString(userId));
                sendResetPasswordEmail(adminToken, userId);
                log.info("Auth: token de reset gerado e email enviado para userId={}", userId);
            } else {
                log.info("Auth: email de reset solicitado para email não encontrado — email={}", maskEmail(email));
            }

            return AuthResponse.forgotPassword(
                "Se o email informado estiver cadastrado, um link de redefinição será enviado."
            );
        } catch (Exception e) {
            log.warn("Auth: falha ao processar forgot-password para {}: {}", maskEmail(email), e.getMessage());
            return AuthResponse.forgotPassword(
                "Se o email informado estiver cadastrado, um link de redefinição será enviado."
            );
        }
    }

    /**
     * Reset Password — redefine a senha usando token de ação do Keycloak.
     * <p>
     * RN13-01: senha deve ter 8+ caracteres, letra + número (validado no Keycloak password policy).
     */
    public AuthResponse resetPassword(ResetPasswordRequest request) {
        log.info("Auth: reset de senha solicitado");

        // Validar token de reset contra o cache (one-time token, TTL 1h — RN13-03)
        UUID userId = resetTokenCache.getIfPresent(request.token());
        if (userId == null) {
            log.warn("Auth: token de reset inválido ou expirado");
            throw new RuntimeException("Token de redefinição inválido ou expirado. Solicite um novo link.");
        }

        try {
            String adminToken = getAdminToken();

            // Keycloak Admin API: PUT /admin/realms/{realm}/users/{id}/reset-password
            String realmUrl = issuerUri.replace("/realms/fbso-platform", "");
            String url = realmUrl + "/admin/realms/fbso-platform/users/" + userId + "/reset-password";

            HttpHeaders headers = new HttpHeaders();
            headers.setBearerAuth(adminToken);
            headers.setContentType(MediaType.APPLICATION_JSON);

            Map<String, Object> body = Map.of(
                "type", "password",
                "value", request.newPassword(),
                "temporary", false
            );

            HttpEntity<Map<String, Object>> entity = new HttpEntity<>(body, headers);
            restTemplate.exchange(url, HttpMethod.PUT, entity, String.class);

            log.info("Auth: senha redefinida com sucesso");
            resetTokenCache.invalidate(request.token()); // one-time token
            return AuthResponse.resetPassword("Senha redefinida com sucesso. Faça login com a nova senha.");

        } catch (Exception e) {
            log.error("Auth: falha ao resetar senha: {}", e.getMessage());
            throw new RuntimeException("Não foi possível redefinir a senha. O link pode ter expirado.");
        }
    }

    // ---- Helpers ----

    private String getAdminToken() {
        String realmUrl = issuerUri.replace("/realms/fbso-platform", "");
        String tokenUrl = realmUrl + "/realms/master/protocol/openid-connect/token";

        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_FORM_URLENCODED);

        String body = "grant_type=client_credentials"
                    + "&client_id=admin-cli"
                    + "&client_secret=" + URLEncoder.encode(clientSecret, StandardCharsets.UTF_8);

        HttpEntity<String> entity = new HttpEntity<>(body, headers);
        ResponseEntity<Map> response = restTemplate.exchange(tokenUrl, HttpMethod.POST, entity, Map.class);

        return (String) response.getBody().get("access_token");
    }

    @SuppressWarnings("unchecked")
    private String findUserByEmail(String adminToken, String email) {
        String realmUrl = issuerUri.replace("/realms/fbso-platform", "");
        String encodedEmail = URLEncoder.encode(email, StandardCharsets.UTF_8);
        String url = realmUrl + "/admin/realms/fbso-platform/users?email=" + encodedEmail + "&exact=true";

        HttpHeaders headers = new HttpHeaders();
        headers.setBearerAuth(adminToken);

        HttpEntity<Void> entity = new HttpEntity<>(headers);
        ResponseEntity<List> response = restTemplate.exchange(url, HttpMethod.GET, entity, List.class);

        List<Map<String, Object>> users = response.getBody();
        if (users != null && !users.isEmpty()) {
            return (String) users.get(0).get("id");
        }
        return null;
    }

    private void sendResetPasswordEmail(String adminToken, String userId) {
        String realmUrl = issuerUri.replace("/realms/fbso-platform", "");
        // userId já validado como UUID em resetPassword() — seguro para path segment
        String url = realmUrl + "/admin/realms/fbso-platform/users/" + userId + "/execute-actions-email";

        HttpHeaders headers = new HttpHeaders();
        headers.setBearerAuth(adminToken);
        headers.setContentType(MediaType.APPLICATION_JSON);

        List<String> actions = List.of("UPDATE_PASSWORD");

        HttpEntity<List<String>> entity = new HttpEntity<>(actions, headers);
        restTemplate.exchange(url, HttpMethod.PUT, entity, String.class);
    }

    private String maskEmail(String email) {
        if (email == null || !email.contains("@")) return "***";
        int at = email.indexOf('@');
        return email.charAt(0) + "***" + email.substring(at);
    }
}
