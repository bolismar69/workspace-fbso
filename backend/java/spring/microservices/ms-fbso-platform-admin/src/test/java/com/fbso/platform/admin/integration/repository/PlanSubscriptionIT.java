package com.fbso.platform.admin.integration.repository;

import com.fbso.platform.admin.dto.request.PlanCreateRequest;
import com.fbso.platform.admin.dto.request.PlanUpdateRequest;
import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.Recurrence;
import com.fbso.platform.admin.enums.TenantSegment;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.exception.PlanHasActiveSubscribersException;
import com.fbso.platform.admin.integration.BaseIntegrationTest;
import com.fbso.platform.admin.repository.PlanRepository;
import com.fbso.platform.admin.repository.SubscriptionRepository;
import com.fbso.platform.admin.repository.TenantRepository;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.PlanService;
import com.fbso.platform.admin.service.SubscriptionService;
import org.flywaydb.core.Flyway;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.DriverManagerDataSource;

import java.math.BigDecimal;
import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

@DisplayName("Plan+Subscription — Integração PostgreSQL")
class PlanSubscriptionIT extends BaseIntegrationTest {

    private JdbcTemplate jdbc;
    private PlanService planService;
    private SubscriptionService subService;
    private TenantRepository tenantRepo;
    private UUID tenantId;

    @BeforeEach
    void setUp() {
        var ds = new DriverManagerDataSource(
                postgres.getJdbcUrl(), postgres.getUsername(), postgres.getPassword());
        this.jdbc = new JdbcTemplate(ds);
        Flyway.configure().dataSource(ds).schemas("fbso_platform")
                .locations("classpath:db/migration").load().migrate();
        jdbc.update("DELETE FROM fbso_platform.subscription");
        jdbc.update("DELETE FROM fbso_platform.plan");
        jdbc.update("DELETE FROM fbso_platform.tenant");

        this.tenantRepo = new TenantRepository(jdbc);
        var planRepo = new PlanRepository(jdbc);
        var subRepo = new SubscriptionRepository(jdbc);
        this.planService = new PlanService(planRepo);
        this.subService = new SubscriptionService(subRepo, planRepo);

        // Seed tenant
        Tenant t = new Tenant();
        t.setNameCorporate("Tenant Teste Ltda");
        t.setSegment(TenantSegment.RETAIL);
        t.setStatus(TenantStatus.ACTIVE);
        tenantRepo.save(t);
        this.tenantId = t.getId();

        // TenantContext.tenantId precisa ser o ID real do tenant criado,
        // pois BaseRepository.save() usa TenantContext.getTenantId() para
        // popular tenant_id em tabelas com hasTenantColumn=true (RLS)
        TenantContext.set(tenantId, UUID.randomUUID(),
                List.of("ADMIN_TENANT"), List.of(), List.of());
    }

    @Nested
    @DisplayName("Plan service")
    class PlanTests {

        @Test
        @DisplayName("deve criar e desativar plano (com plano reserva)")
        void shouldCreateAndDeactivate() {
            // Precisa de 2+ planos para poder desativar (RN06-03)
            planService.create(new PlanCreateRequest(
                    "Plano Reserva", null, new BigDecimal("999.00"), Recurrence.YEARLY));
            var plan = planService.create(new PlanCreateRequest(
                    "Plano Teste", "Desc", new BigDecimal("49.00"), Recurrence.MONTHLY));
            assertThat(plan.status()).isEqualTo("ACTIVE");

            var deactivated = planService.deactivate(plan.id());
            assertThat(deactivated.status()).isEqualTo("DISCONTINUED");
        }

        @Test
        @DisplayName("deve versionar quando preço altera")
        void shouldVersionOnPriceChange() {
            var plan = planService.create(new PlanCreateRequest(
                    "Versão Teste", null, new BigDecimal("99.00"), Recurrence.QUARTERLY));

            var updated = planService.update(plan.id(), new PlanUpdateRequest(
                    null, null, new BigDecimal("149.00"), null));

            assertThat(updated.version()).isEqualTo(2);
            assertThat(updated.price()).isEqualByComparingTo("149.00");
        }

        @Test
        @org.junit.jupiter.api.Disabled("Transactional sem Spring Context") @DisplayName("não deve permitir desativar plano com assinantes ativos (RN06-01)")
        void shouldBlockDeactivationWithSubscribers() {
            // Cria segundo plano para evitar o erro "último plano ativo"
            planService.create(new PlanCreateRequest(
                    "Plano Reserva", null, new BigDecimal("999.00"), Recurrence.YEARLY));
            var plan = planService.create(new PlanCreateRequest(
                    "Com Assinantes", null, new BigDecimal("299.00"), Recurrence.MONTHLY));
            subService.create(tenantId, plan.id());

            assertThatThrownBy(() -> planService.deactivate(plan.id()))
                    .isInstanceOf(PlanHasActiveSubscribersException.class);
        }
    }

    @Nested
    @DisplayName("Subscription service")
    class SubscriptionTests {

        @Test
        @DisplayName("deve criar assinatura com locked_price (DT-009)")
        void shouldCreateWithLockedPrice() {
            var plan = planService.create(new PlanCreateRequest(
                    "Locked Price Test", null, new BigDecimal("199.00"), Recurrence.MONTHLY));

            var sub = subService.create(tenantId, plan.id());

            assertThat(sub.lockedPrice()).isEqualByComparingTo("199.00");
            assertThat(sub.lockedRecurrence()).isEqualTo("MONTHLY");
        }

        @Test
        @DisplayName("deve rejeitar segunda assinatura ativa (RN07-01)")
        void shouldRejectSecondActive() {
            var plan = planService.create(new PlanCreateRequest(
                    "Única Assinatura", null, new BigDecimal("79.00"), Recurrence.MONTHLY));
            subService.create(tenantId, plan.id());

            assertThatThrownBy(() -> subService.create(tenantId, plan.id()))
                    .isInstanceOf(IllegalStateException.class)
                    .hasMessageContaining("já possui assinatura ativa");
        }

        @Test
        @DisplayName("change-plan deve criar nova e finalizar anterior (RN07-02)")
        void shouldChangePlanWithoutGap() {
            var plan1 = planService.create(new PlanCreateRequest(
                    "Plano Antigo", null, new BigDecimal("99.00"), Recurrence.MONTHLY));
            var plan2 = planService.create(new PlanCreateRequest(
                    "Plano Novo", null, new BigDecimal("299.00"), Recurrence.MONTHLY));

            var sub = subService.create(tenantId, plan1.id());
            var changed = subService.changePlan(sub.id(), plan2.id());

            assertThat(changed.planId()).isEqualTo(plan2.id());
            assertThat(changed.lockedPrice()).isEqualByComparingTo("299.00");
        }
    }
}
