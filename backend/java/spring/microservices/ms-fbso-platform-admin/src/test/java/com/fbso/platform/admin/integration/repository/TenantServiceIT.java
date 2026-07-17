package com.fbso.platform.admin.integration.repository;

import com.fbso.platform.admin.dto.request.TenantCreateRequest;
import com.fbso.platform.admin.dto.request.TenantUpdateRequest;
import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantSegment;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.exception.DuplicateCnpjException;
import com.fbso.platform.admin.exception.InvalidStatusTransitionException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
import com.fbso.platform.admin.integration.BaseIntegrationTest;
import com.fbso.platform.admin.repository.TenantRepository;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.TenantService;
import org.flywaydb.core.Flyway;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.DriverManagerDataSource;

import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

@DisplayName("TenantService — Integração PostgreSQL")
class TenantServiceIT extends BaseIntegrationTest {

    private JdbcTemplate jdbc;
    private TenantService service;
    private TenantRepository repo;

    @BeforeEach
    void setUp() {
        var ds = new DriverManagerDataSource(
                postgres.getJdbcUrl(), postgres.getUsername(), postgres.getPassword());
        this.jdbc = new JdbcTemplate(ds);
        Flyway.configure().dataSource(ds).schemas("fbso_platform")
                .locations("classpath:db/migration").load().migrate();
        jdbc.update("DELETE FROM fbso_platform.subscription");
        jdbc.update("DELETE FROM fbso_platform.tenant");

        // tenant table is global (hasTenantColumn=false), so random tenantId is fine
        TenantContext.set(UUID.randomUUID(), UUID.randomUUID(),
                List.of("ADMIN_TENANT"), List.of(), List.of());
        this.repo = new TenantRepository(jdbc);
        this.service = new TenantService(repo);
    }

    @Nested
    @DisplayName("create")
    class Create {

        @Test
        @DisplayName("deve criar tenant PENDING_ONBOARDING")
        void shouldCreateTenant() {
            var req = new TenantCreateRequest("Empresa Teste Ltda", "Teste",
                    TenantSegment.RETAIL);
            var result = service.create(req);

            assertThat(result.status()).isEqualTo(TenantStatus.PENDING_ONBOARDING);
            assertThat(result.id()).isNotNull();
        }

        @Test
        @DisplayName("deve lançar DuplicateCnpjException ao duplicar razão social")
        void shouldRejectDuplicate() {
            service.create(new TenantCreateRequest("Única Ltda", null, TenantSegment.TECHNOLOGY));

            assertThatThrownBy(() -> service.create(
                    new TenantCreateRequest("Única Ltda", null, TenantSegment.TECHNOLOGY)))
                    .isInstanceOf(DuplicateCnpjException.class);
        }
    }

    @Nested
    @DisplayName("transições de status")
    class StatusTransitions {

        @Test
        @DisplayName("ciclo PENDING → ACTIVE → SUSPENDED → ACTIVE")
        void shouldCompleteFullCycle() {
            var created = service.create(new TenantCreateRequest(
                    "Ciclo Ltda", "Ciclo", TenantSegment.SERVICES));

            // PENDING → ACTIVE (via update + status change)
            var tenant = repo.findById(created.id()).orElseThrow();
            tenant.setStatus(TenantStatus.ACTIVE);
            repo.update(tenant);

            // ACTIVE → SUSPENDED
            var suspended = service.suspend(created.id(), "Teste de suspensão");
            assertThat(suspended.status()).isEqualTo(TenantStatus.SUSPENDED);

            // SUSPENDED → ACTIVE
            var reactivated = service.reactivate(created.id());
            assertThat(reactivated.status()).isEqualTo(TenantStatus.ACTIVE);
        }

        @Test
        @DisplayName("PENDING → SUSPENDED deve lançar InvalidStatusTransitionException")
        void shouldRejectInvalidTransition() {
            var created = service.create(new TenantCreateRequest(
                    "Inválida Ltda", "Inválida", TenantSegment.EDUCATION));

            assertThatThrownBy(() -> service.suspend(created.id(), "Forçando erro"))
                    .isInstanceOf(InvalidStatusTransitionException.class);
        }

        @Test
        @DisplayName("suspend sem motivo deve lançar IllegalArgumentException")
        void shouldRequireReason() {
            var created = service.create(new TenantCreateRequest(
                    "Sem Motivo Ltda", "Sem Motivo", TenantSegment.HEALTHCARE));
            var tenant = repo.findById(created.id()).orElseThrow();
            tenant.setStatus(TenantStatus.ACTIVE);
            repo.update(tenant);

            assertThatThrownBy(() -> service.suspend(created.id(), ""))
                    .isInstanceOf(IllegalArgumentException.class);
        }

        @Test
        @DisplayName("TenantNotFoundException para ID inexistente")
        void shouldThrowOnInvalidId() {
            assertThatThrownBy(() -> service.suspend(UUID.randomUUID(), "Motivo"))
                    .isInstanceOf(TenantNotFoundException.class);
        }
    }
}
