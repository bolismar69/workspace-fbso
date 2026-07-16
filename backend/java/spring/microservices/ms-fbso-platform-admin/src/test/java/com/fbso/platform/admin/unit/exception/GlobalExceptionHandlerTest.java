package com.fbso.platform.admin.unit.exception;

import com.fbso.platform.admin.dto.response.ErrorResponse;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.exception.GlobalExceptionHandler;
import com.fbso.platform.admin.exception.PermissionDeniedException;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.springframework.http.ResponseEntity;

import static org.assertj.core.api.Assertions.assertThat;

@DisplayName("GlobalExceptionHandler")
class GlobalExceptionHandlerTest {

    private final GlobalExceptionHandler handler = new GlobalExceptionHandler();

    // ---- TC-S2-018: BusinessException → 422 ----

    @Nested
    @DisplayName("TC-S2-018 — BusinessException")
    class BusinessExceptionTest {

        @Test
        @DisplayName("deve retornar 422 com corpo RFC 7807")
        void shouldReturn422WithRfc7807() {
            var ex = new BusinessException("duplicate-cnpj", "CNPJ já cadastrado");

            ResponseEntity<ErrorResponse> response = handler.handleBusinessException(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(422);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().title()).isEqualTo("CNPJ já cadastrado");
            assertThat(response.getBody().status()).isEqualTo(422);
            assertThat(response.getBody().type()).contains("duplicate-cnpj");
        }
    }

    // ---- TC-S2-019: PermissionDeniedException → 403 ----

    @Nested
    @DisplayName("TC-S2-019 — PermissionDeniedException")
    class PermissionDeniedTest {

        @Test
        @DisplayName("deve retornar 403 com mensagem amigável em PT-BR")
        void shouldReturn403WithFriendlyMessage() {
            var ex = new PermissionDeniedException();

            ResponseEntity<ErrorResponse> response = handler.handlePermissionDenied(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(403);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().title()).isEqualTo("Acesso negado");
            assertThat(response.getBody().status()).isEqualTo(403);
            // Mensagem amigável — sem detalhes técnicos
            assertThat(response.getBody().detail()).contains("Você não tem permissão");
        }
    }

    // ---- TC-S2-020: Exception → 500 sem stack trace ----

    @Nested
    @DisplayName("TC-S2-020 — Erro genérico 500")
    class GenericExceptionTest {

        @Test
        @DisplayName("deve retornar 500 SEM stack trace")
        void shouldReturn500WithoutStackTrace() {
            var ex = new RuntimeException("Detalhe interno sensível");

            ResponseEntity<ErrorResponse> response = handler.handleGenericException(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(500);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().title()).isEqualTo("Erro interno do servidor");
            // NUNCA expor a mensagem da exceção original
            assertThat(response.getBody().detail()).doesNotContain("Detalhe interno");
        }
    }

    // ---- TC-S2-021: SecurityException → 401 ----

    @Nested
    @DisplayName("TC-S2-021 — SecurityException")
    class SecurityExceptionTest {

        @Test
        @DisplayName("deve retornar 401 para violação de segurança")
        void shouldReturn401ForSecurityViolation() {
            var ex = new SecurityException("Token ausente");

            ResponseEntity<ErrorResponse> response = handler.handleSecurityException(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(401);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().title()).isEqualTo("Token de acesso não informado");
        }
    }
}
