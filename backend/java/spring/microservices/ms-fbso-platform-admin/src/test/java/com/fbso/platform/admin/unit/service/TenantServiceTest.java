package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.request.TenantCreateRequest;
import com.fbso.platform.admin.dto.request.TenantUpdateRequest;
import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantSegment;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.exception.DuplicateCnpjException;
import com.fbso.platform.admin.exception.InvalidStatusTransitionException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
import com.fbso.platform.admin.repository.TenantRepository;
import com.fbso.platform.admin.service.TenantService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
@org.mockito.junit.jupiter.MockitoSettings(strictness = org.mockito.quality.Strictness.LENIENT)
@DisplayName("TenantService")
class TenantServiceTest {

    @Mock private TenantRepository repo;
    private TenantService service;

    @BeforeEach
    void setUp() {
        service = new TenantService(repo);
    }

    @Nested
    @DisplayName("create")
    class Create {

        @Test
        @DisplayName("deve criar tenant com status PENDING_ONBOARDING (RN04-01)")
        void shouldCreateWithPendingStatus() {
            when(repo.findByNameCorporate("Mercado Ltda")).thenReturn(Optional.empty());
            doNothing().when(repo).save(any());

            var req = new TenantCreateRequest("Mercado Ltda", "Mercado", TenantSegment.RETAIL);
            var result = service.create(req);

            assertThat(result.status()).isEqualTo(TenantStatus.PENDING_ONBOARDING);
            assertThat(result.nameCorporate()).isEqualTo("Mercado Ltda");
        }

        @Test
        @DisplayName("deve lançar DuplicateCnpjException quando razão social duplicada (RN04-02)")
        void shouldThrowOnDuplicateName() {
            when(repo.findByNameCorporate("Mercado Ltda"))
                    .thenReturn(Optional.of(new Tenant()));

            var req = new TenantCreateRequest("Mercado Ltda", null, TenantSegment.RETAIL);

            assertThatThrownBy(() -> service.create(req))
                    .isInstanceOf(DuplicateCnpjException.class);
        }
    }

    @Nested
    @DisplayName("update")
    class Update {

        @Test
        @DisplayName("deve atualizar campos editáveis")
        void shouldUpdateFields() {
            UUID id = UUID.randomUUID();
            Tenant existing = new Tenant();
            existing.setId(id);
            existing.setNameCorporate("Old Name");
            existing.setStatus(TenantStatus.ACTIVE);

            when(repo.findById(id)).thenReturn(Optional.of(existing));
            doNothing().when(repo).update(any());

            var req = new TenantUpdateRequest("New Fantasy", TenantSegment.SERVICES);
            var result = service.update(id, req);

            assertThat(result.nameFantasy()).isEqualTo("New Fantasy");
            assertThat(result.segment()).isEqualTo(TenantSegment.SERVICES);
        }

        @Test
        @DisplayName("deve lançar TenantNotFoundException quando ID não existe")
        void shouldThrowOnNotFound() {
            UUID id = UUID.randomUUID();
            when(repo.findById(id)).thenReturn(Optional.empty());

            var req = new TenantUpdateRequest("X", null);
            assertThatThrownBy(() -> service.update(id, req))
                    .isInstanceOf(TenantNotFoundException.class);
        }
    }

    @Nested
    @DisplayName("suspend (RN05-02)")
    class Suspend {

        @Test
        @DisplayName("deve suspender tenant ativo com motivo")
        void shouldSuspendActiveTenant() {
            UUID id = UUID.randomUUID();
            Tenant tenant = new Tenant();
            tenant.setId(id);
            tenant.setNameCorporate("Test");
            tenant.setStatus(TenantStatus.ACTIVE);

            when(repo.findById(id)).thenReturn(Optional.of(tenant));
            doNothing().when(repo).update(any());

            var result = service.suspend(id, "Inadimplência");

            assertThat(result.status()).isEqualTo(TenantStatus.SUSPENDED);
        }

        @Test
        @DisplayName("deve lançar exceção quando motivo é vazio")
        void shouldThrowOnEmptyReason() {
            assertThatThrownBy(() -> service.suspend(UUID.randomUUID(), ""))
                    .isInstanceOf(IllegalArgumentException.class);
        }

        @Test
        @DisplayName("deve lançar InvalidStatusTransitionException para transição inválida")
        void shouldThrowOnInvalidTransition() {
            UUID id = UUID.randomUUID();
            Tenant tenant = new Tenant();
            tenant.setId(id);
            tenant.setNameCorporate("Test");
            tenant.setStatus(TenantStatus.PENDING_ONBOARDING);

            when(repo.findById(id)).thenReturn(Optional.of(tenant));

            assertThatThrownBy(() -> service.suspend(id, "Motivo válido"))
                    .isInstanceOf(InvalidStatusTransitionException.class);
        }
    }

    @Nested
    @DisplayName("reactivate (RN05-03)")
    class Reactivate {

        @Test
        @DisplayName("deve reativar tenant suspenso")
        void shouldReactivateSuspended() {
            UUID id = UUID.randomUUID();
            Tenant tenant = new Tenant();
            tenant.setId(id);
            tenant.setNameCorporate("Test");
            tenant.setStatus(TenantStatus.SUSPENDED);

            when(repo.findById(id)).thenReturn(Optional.of(tenant));
            doNothing().when(repo).update(any());

            var result = service.reactivate(id);

            assertThat(result.status()).isEqualTo(TenantStatus.ACTIVE);
        }
    }

    @Nested
    @DisplayName("validateTransition")
    class ValidateTransition {

        @Test
        @DisplayName("PENDING → ACTIVE permitida")
        void pendingToActiveOk() {
            TenantService.validateTransition(
                    TenantStatus.PENDING_ONBOARDING, TenantStatus.ACTIVE, "test");
        }

        @Test
        @DisplayName("ACTIVE → PENDING proibida")
        void activeToPendingInvalid() {
            assertThatThrownBy(() -> TenantService.validateTransition(
                    TenantStatus.ACTIVE, TenantStatus.PENDING_ONBOARDING, "test"))
                    .isInstanceOf(InvalidStatusTransitionException.class);
        }
    }
}
