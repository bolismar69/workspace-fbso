package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.entity.Plan;
import com.fbso.platform.admin.entity.Subscription;
import com.fbso.platform.admin.repository.PlanRepository;
import com.fbso.platform.admin.repository.SubscriptionRepository;
import com.fbso.platform.admin.service.SubscriptionService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
@org.mockito.junit.jupiter.MockitoSettings(strictness = org.mockito.quality.Strictness.LENIENT)
@DisplayName("SubscriptionService")
class SubscriptionServiceTest {

    @Mock private SubscriptionRepository subRepo;
    @Mock private PlanRepository planRepo;
    private SubscriptionService service;

    @BeforeEach
    void setUp() {
        service = new SubscriptionService(subRepo, planRepo);
    }

    @Nested
    @DisplayName("create")
    class Create {

        @Test
        @DisplayName("deve criar assinatura com locked_price do plano (DT-009)")
        void shouldCreateWithLockedPrice() {
            UUID tenantId = UUID.randomUUID();
            UUID planId = UUID.randomUUID();

            Plan plan = new Plan();
            plan.setId(planId);
            plan.setName("Básico");
            plan.setPrice(new BigDecimal("99.00"));

            when(subRepo.findActiveByTenantId(tenantId)).thenReturn(Optional.empty());
            when(planRepo.findById(planId)).thenReturn(Optional.of(plan));
            doNothing().when(subRepo).save(any());

            var result = service.create(tenantId, planId);

            assertThat(result.status()).isEqualTo("ACTIVE");
            assertThat(result.lockedPrice()).isEqualByComparingTo("99.00");
        }

        @Test
        @DisplayName("deve lançar exceção quando já existe assinatura ativa (RN07-01)")
        void shouldThrowWhenActiveExists() {
            UUID tenantId = UUID.randomUUID();
            UUID planId = UUID.randomUUID();

            when(subRepo.findActiveByTenantId(tenantId))
                    .thenReturn(Optional.of(new Subscription()));

            assertThatThrownBy(() -> service.create(tenantId, planId))
                    .isInstanceOf(IllegalStateException.class)
                    .hasMessageContaining("já possui assinatura ativa");
        }
    }

    @Nested
    @DisplayName("changePlan (RN07-02)")
    class ChangePlan {

        @Test
        @DisplayName("deve finalizar atual e criar nova sem gap")
        void shouldChangePlanWithoutGap() {
            UUID subId = UUID.randomUUID();
            UUID newPlanId = UUID.randomUUID();

            Subscription current = new Subscription();
            current.setId(subId);
            current.setTenantId(UUID.randomUUID());
            current.setPlanId(UUID.randomUUID());
            current.setStatus("ACTIVE");

            Plan newPlan = new Plan();
            newPlan.setId(newPlanId);
            newPlan.setName("Enterprise");
            newPlan.setPrice(new BigDecimal("999.00"));

            when(subRepo.findById(subId)).thenReturn(Optional.of(current));
            when(planRepo.findById(newPlanId)).thenReturn(Optional.of(newPlan));
            doNothing().when(subRepo).update(any());
            doNothing().when(subRepo).save(any());

            var result = service.changePlan(subId, newPlanId);

            assertThat(result.status()).isEqualTo("ACTIVE");
            assertThat(result.planId()).isEqualTo(newPlanId);
        }
    }

    @Nested
    @DisplayName("suspend")
    class Suspend {

        @Test
        @DisplayName("deve suspender assinatura ativa")
        void shouldSuspend() {
            UUID subId = UUID.randomUUID();
            Subscription sub = new Subscription();
            sub.setId(subId);
            sub.setStatus("ACTIVE");

            when(subRepo.findById(subId)).thenReturn(Optional.of(sub));
            doNothing().when(subRepo).update(any());

            var result = service.suspend(subId);

            assertThat(result.status()).isEqualTo("SUSPENDED");
        }
    }
}
