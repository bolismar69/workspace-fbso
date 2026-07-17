package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.response.AccountsByPlanResponse;
import com.fbso.platform.admin.dto.response.AccountsByStatusResponse;
import com.fbso.platform.admin.dto.response.AlertResponse;
import com.fbso.platform.admin.dto.response.DashboardSummaryResponse;
import com.fbso.platform.admin.dto.response.EvolutionResponse;
import com.fbso.platform.admin.repository.DashboardRepository;
import com.fbso.platform.admin.service.DashboardService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("DashboardService")
class DashboardServiceTest {

    @Mock private DashboardRepository repo;
    private DashboardService service;

    @BeforeEach
    void setUp() {
        service = new DashboardService(repo);
    }

    @Nested
    @DisplayName("getSummary")
    class Summary {

        @Test
        @DisplayName("deve retornar indicadores consolidados corretos")
        void shouldReturnCorrectSummary() {
            when(repo.countActive()).thenReturn(50);
            when(repo.countByStatus("ACTIVE")).thenReturn(35);
            when(repo.countByStatus("PENDING_ONBOARDING")).thenReturn(10);
            when(repo.countByStatus("SUSPENDED")).thenReturn(5);
            when(repo.accountsByPlan()).thenReturn(List.of(
                    Map.of("plan_name", "Básico", "count", 25L),
                    Map.of("plan_name", "Avançado", "count", 10L)));
            when(repo.monthlyRevenue()).thenReturn(new BigDecimal("14990.00"));

            DashboardSummaryResponse result = service.getSummary();

            assertThat(result.totalAccounts()).isEqualTo(50);
            assertThat(result.activeAccounts()).isEqualTo(35);
            assertThat(result.pendingAccounts()).isEqualTo(10);
            assertThat(result.suspendedAccounts()).isEqualTo(5);
            assertThat(result.accountsByPlan()).containsKeys("Básico", "Avançado");
            assertThat(result.monthlyRevenue()).isEqualByComparingTo("14990.00");
            assertThat(result.period()).isEqualTo("mes_atual");
        }
    }

    @Nested
    @DisplayName("getEvolution")
    class Evolution {

        @Test
        @DisplayName("deve retornar dados de evolução para período 30d")
        void shouldReturnEvolutionFor30d() {
            when(repo.evolutionByDay(org.mockito.ArgumentMatchers.any()))
                    .thenReturn(List.of(
                            Map.of("date", java.sql.Date.valueOf("2026-06-17"), "count", 5L),
                            Map.of("date", java.sql.Date.valueOf("2026-06-18"), "count", 8L)));

            EvolutionResponse result = service.getEvolution("30d");

            assertThat(result.period()).isEqualTo("30d");
            assertThat(result.dataPoints()).hasSize(2);
        }

        @Test
        @DisplayName("período inválido deve assumir mês atual (RN01-02)")
        void shouldDefaultToCurrentMonthForInvalidPeriod() {
            when(repo.evolutionByDay(org.mockito.ArgumentMatchers.any()))
                    .thenReturn(List.of());

            EvolutionResponse result = service.getEvolution("invalido");

            assertThat(result.period()).isEqualTo("mes_atual");
        }

        @Test
        @DisplayName("período nulo deve assumir mês atual")
        void shouldDefaultToCurrentMonthForNull() {
            when(repo.evolutionByDay(org.mockito.ArgumentMatchers.any()))
                    .thenReturn(List.of());

            EvolutionResponse result = service.getEvolution(null);

            assertThat(result.period()).isEqualTo("mes_atual");
        }
    }

    @Nested
    @DisplayName("getAccountsByPlan")
    class ByPlan {

        @Test
        @DisplayName("deve retornar distribuição por plano")
        void shouldReturnPlanDistribution() {
            when(repo.accountsByPlan()).thenReturn(List.of(
                    Map.of("plan_name", "Enterprise", "count", 15L)));

            AccountsByPlanResponse result = service.getAccountsByPlan();

            assertThat(result.plans()).hasSize(1);
            assertThat(result.plans().get(0).planName()).isEqualTo("Enterprise");
            assertThat(result.plans().get(0).count()).isEqualTo(15);
        }
    }

    @Nested
    @DisplayName("getAccountsByStatus")
    class ByStatus {

        @Test
        @DisplayName("deve retornar distribuição por status")
        void shouldReturnStatusDistribution() {
            when(repo.accountsByStatus()).thenReturn(List.of(
                    Map.of("status", "ACTIVE", "count", 35L)));

            AccountsByStatusResponse result = service.getAccountsByStatus();

            assertThat(result.statuses()).hasSize(1);
            assertThat(result.statuses().get(0).status()).isEqualTo("ACTIVE");
            assertThat(result.statuses().get(0).count()).isEqualTo(35);
        }
    }

    @Nested
    @DisplayName("getAlerts")
    class Alerts {

        @Test
        @DisplayName("deve retornar alertas WARNING para onboarding >48h")
        void shouldReturnWarningAlerts() {
            UUID tenantId = UUID.randomUUID();
            when(repo.onboardingStalled()).thenReturn(List.of(
                    Map.of("id", tenantId, "name_corporate", "Empresa XPTO",
                            "status", "PENDING_ONBOARDING")));
            when(repo.suspendedSubscriptions()).thenReturn(List.of());

            AlertResponse result = service.getAlerts();

            assertThat(result.alerts()).hasSize(1);
            assertThat(result.alerts().get(0).type()).isEqualTo(AlertResponse.AlertType.WARNING);
            assertThat(result.alerts().get(0).entityId()).isEqualTo(tenantId);
        }

        @Test
        @DisplayName("deve retornar alertas CRITICAL para assinaturas suspensas")
        void shouldReturnCriticalAlerts() {
            UUID tenantId = UUID.randomUUID();
            when(repo.onboardingStalled()).thenReturn(List.of());
            when(repo.suspendedSubscriptions()).thenReturn(List.of(
                    Map.of("subscription_id", UUID.randomUUID(),
                            "tenant_id", tenantId,
                            "name_corporate", "Empresa Inadimplente",
                            "plan_name", "Básico",
                            "status", "SUSPENDED")));

            AlertResponse result = service.getAlerts();

            assertThat(result.alerts()).hasSize(1);
            assertThat(result.alerts().get(0).type()).isEqualTo(AlertResponse.AlertType.CRITICAL);
        }
    }
}
