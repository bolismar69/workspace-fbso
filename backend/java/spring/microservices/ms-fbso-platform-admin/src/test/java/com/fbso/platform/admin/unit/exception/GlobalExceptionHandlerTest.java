package com.fbso.platform.admin.unit.exception;

import com.fbso.platform.admin.dto.response.ErrorResponse;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.exception.DuplicateCnpjException;
import com.fbso.platform.admin.exception.GlobalExceptionHandler;
import com.fbso.platform.admin.exception.InvalidStatusTransitionException;
import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.exception.PlanHasActiveSubscribersException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
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

        @Test
        @DisplayName("TC-T-054: 403 segue EXATAMENTE o formato RFC 7807 com type, title, status, detail")
        void shouldReturn403ExactRfc7807Format() {
            var ex = new PermissionDeniedException();

            ResponseEntity<ErrorResponse> response = handler.handlePermissionDenied(ex);
            ErrorResponse body = response.getBody();

            assertThat(body).isNotNull();
            // RFC 7807 requer campo type (URI identificando o tipo de erro)
            assertThat(body.type()).isEqualTo("https://api.fbso.org/errors/access-denied");
            // Título em PT-BR
            assertThat(body.title()).isEqualTo("Acesso negado");
            // Status HTTP
            assertThat(body.status()).isEqualTo(403);
            // Detalhe amigável (RN12-02)
            assertThat(body.detail()).isEqualTo("Você não tem permissão para acessar esta área.");
            // fields deve ser null (não é erro de validação)
            assertThat(body.fields()).isNull();
        }

        @Test
        @DisplayName("TC-T-054: Spring AccessDeniedException também retorna 403 RFC 7807")
        void springAccessDeniedAlsoReturnsRfc7807() {
            var ex = new org.springframework.security.access.AccessDeniedException("Access Denied");

            ResponseEntity<ErrorResponse> response = handler.handleSpringAccessDenied(ex);
            ErrorResponse body = response.getBody();

            assertThat(body).isNotNull();
            assertThat(body.status()).isEqualTo(403);
            assertThat(body.type()).contains("access-denied");
            assertThat(body.title()).isEqualTo("Acesso negado");
            assertThat(body.detail()).contains("Você não tem permissão");
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

    // ---- TenantNotFoundException → 404 ----

    @Nested
    @DisplayName("TenantNotFoundException → 404")
    class TenantNotFoundTest {

        @Test
        @DisplayName("deve retornar 404 com código do erro")
        void shouldReturn404() {
            var ex = new TenantNotFoundException("tenant-123");

            ResponseEntity<ErrorResponse> response = handler.handleTenantNotFound(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(404);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().type()).contains("tenant-not-found");
            assertThat(response.getBody().title()).contains("tenant-123");
        }
    }

    // ---- DuplicateCnpjException → 409 ----

    @Nested
    @DisplayName("DuplicateCnpjException → 409")
    class DuplicateCnpjTest {

        @Test
        @DisplayName("deve retornar 409 com CNPJ informado")
        void shouldReturn409() {
            var ex = new DuplicateCnpjException("12345678000199");

            ResponseEntity<ErrorResponse> response = handler.handleDuplicateCnpj(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(409);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().type()).contains("duplicate-cnpj");
            assertThat(response.getBody().title()).contains("12345678000199");
        }
    }

    // ---- Subclasses de BusinessException ----

    @Nested
    @DisplayName("Subclasses de BusinessException → 422")
    class BusinessSubclassTest {

        @Test
        @DisplayName("InvalidStatusTransitionException deve retornar 422")
        void invalidStatusTransitionShouldReturn422() {
            var ex = new InvalidStatusTransitionException("ACTIVE", "PENDING_ONBOARDING");

            ResponseEntity<ErrorResponse> response = handler.handleBusinessException(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(422);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().title()).contains("ACTIVE");
        }

        @Test
        @DisplayName("PlanHasActiveSubscribersException deve retornar 422")
        void planHasActiveSubscribersShouldReturn422() {
            var ex = new PlanHasActiveSubscribersException("Plano Enterprise");

            ResponseEntity<ErrorResponse> response = handler.handleBusinessException(ex);

            assertThat(response.getStatusCode().value()).isEqualTo(422);
            assertThat(response.getBody()).isNotNull();
            assertThat(response.getBody().title()).contains("Plano Enterprise");
        }
    }
}
