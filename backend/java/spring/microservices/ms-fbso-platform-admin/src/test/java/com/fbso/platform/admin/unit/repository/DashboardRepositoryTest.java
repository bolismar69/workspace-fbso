package com.fbso.platform.admin.unit.repository;

import com.fbso.platform.admin.repository.DashboardRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;

import java.math.BigDecimal;
import java.sql.Date;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("DashboardRepository")
class DashboardRepositoryTest {

    @Mock private JdbcTemplate jdbc;
    private DashboardRepository repo;

    @BeforeEach
    void setUp() {
        repo = new DashboardRepository(jdbc);
    }

    @Nested
    @DisplayName("countActive")
    class CountActive {

        @Test
        @DisplayName("deve retornar total de tenants não deletados")
        void shouldReturnActiveCount() {
            when(jdbc.queryForObject(anyString(), eq(Integer.class)))
                    .thenReturn(10);

            assertThat(repo.countActive()).isEqualTo(10);
        }

        @Test
        @DisplayName("deve retornar 0 quando queryForObject retorna null")
        void shouldReturnZeroWhenNull() {
            when(jdbc.queryForObject(anyString(), eq(Integer.class)))
                    .thenReturn(null);

            assertThat(repo.countActive()).isZero();
        }
    }

    @Nested
    @DisplayName("countByStatus")
    class CountByStatus {

        @Test
        @DisplayName("deve contar tenants pelo status")
        void shouldCountByStatus() {
            when(jdbc.queryForObject(anyString(), eq(Integer.class), eq("ACTIVE")))
                    .thenReturn(5);

            assertThat(repo.countByStatus("ACTIVE")).isEqualTo(5);
        }

        @Test
        @DisplayName("deve retornar 0 quando status não encontrado")
        void shouldReturnZeroForUnknownStatus() {
            when(jdbc.queryForObject(anyString(), eq(Integer.class), eq("INACTIVE")))
                    .thenReturn(null);

            assertThat(repo.countByStatus("INACTIVE")).isZero();
        }
    }

    @Nested
    @DisplayName("accountsByStatus")
    class AccountsByStatus {

        @Test
        @DisplayName("deve retornar agrupamento por status")
        void shouldReturnGroupedByStatus() {
            when(jdbc.queryForList(anyString()))
                    .thenReturn(List.of(
                            Map.of("status", "ACTIVE", "count", 5L),
                            Map.of("status", "PENDING_ONBOARDING", "count", 3L)));

            var result = repo.accountsByStatus();

            assertThat(result).hasSize(2);
            assertThat(result.get(0)).containsEntry("status", "ACTIVE");
        }
    }

    @Nested
    @DisplayName("accountsByPlan")
    class AccountsByPlan {

        @Test
        @DisplayName("deve retornar distribuição por plano")
        void shouldReturnByPlan() {
            when(jdbc.queryForList(anyString()))
                    .thenReturn(List.of(
                            Map.of("plan_name", "Básico", "count", 3L),
                            Map.of("plan_name", "Enterprise", "count", 2L)));

            var result = repo.accountsByPlan();

            assertThat(result).hasSize(2);
            assertThat(result.get(0).get("plan_name")).isEqualTo("Básico");
        }
    }

    @Nested
    @DisplayName("monthlyRevenue")
    class MonthlyRevenue {

        @Test
        @DisplayName("deve retornar receita mensal somada")
        void shouldReturnRevenue() {
            when(jdbc.queryForObject(anyString(), eq(BigDecimal.class)))
                    .thenReturn(new BigDecimal("1795.00"));

            var result = repo.monthlyRevenue();

            assertThat(result).isEqualByComparingTo("1795.00");
        }
    }

    @Nested
    @DisplayName("evolutionByDay")
    class EvolutionByDay {

        @Test
        @DisplayName("deve retornar evolução temporal desde a data informada")
        void shouldReturnEvolution() {
            var since = OffsetDateTime.now().minusDays(30);
            when(jdbc.queryForList(anyString(), eq(since)))
                    .thenReturn(List.of(
                            Map.of("date", Date.valueOf("2026-07-01"), "count", 3L),
                            Map.of("date", Date.valueOf("2026-07-02"), "count", 5L)));

            var result = repo.evolutionByDay(since);

            assertThat(result).hasSize(2);
        }
    }

    @Nested
    @DisplayName("onboardingStalled")
    class OnboardingStalled {

        @Test
        @DisplayName("deve retornar tenants com onboarding >48h")
        void shouldReturnStalledOnboarding() {
            when(jdbc.queryForList(anyString()))
                    .thenReturn(List.of(
                            Map.of("id", java.util.UUID.randomUUID(),
                                    "name_corporate", "Empresa Lenta",
                                    "status", "PENDING_ONBOARDING",
                                    "created_dt", OffsetDateTime.now().minusDays(5))));

            var result = repo.onboardingStalled();

            assertThat(result).hasSize(1);
            assertThat(result.get(0).get("name_corporate")).isEqualTo("Empresa Lenta");
        }

        @Test
        @DisplayName("deve retornar lista vazia quando não há pendentes >48h")
        void shouldReturnEmptyWhenNone() {
            when(jdbc.queryForList(anyString())).thenReturn(List.of());

            assertThat(repo.onboardingStalled()).isEmpty();
        }
    }

    @Nested
    @DisplayName("suspendedSubscriptions")
    class SuspendedSubscriptions {

        @Test
        @DisplayName("deve retornar assinaturas suspensas com dados do tenant e plano")
        void shouldReturnSuspended() {
            when(jdbc.queryForList(anyString()))
                    .thenReturn(List.of(
                            Map.of("subscription_id", java.util.UUID.randomUUID(),
                                    "tenant_id", java.util.UUID.randomUUID(),
                                    "name_corporate", "Empresa Suspensa",
                                    "plan_name", "Básico",
                                    "status", "SUSPENDED")));

            var result = repo.suspendedSubscriptions();

            assertThat(result).hasSize(1);
            assertThat(result.get(0).get("plan_name")).isEqualTo("Básico");
        }

        @Test
        @DisplayName("deve retornar lista vazia quando não há suspensas")
        void shouldReturnEmptyWhenNone() {
            when(jdbc.queryForList(anyString())).thenReturn(List.of());

            assertThat(repo.suspendedSubscriptions()).isEmpty();
        }
    }
}
