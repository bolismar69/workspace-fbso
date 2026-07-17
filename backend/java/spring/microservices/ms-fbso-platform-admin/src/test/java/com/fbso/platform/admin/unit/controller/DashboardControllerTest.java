package com.fbso.platform.admin.unit.controller;

import com.fbso.platform.admin.controller.DashboardController;
import com.fbso.platform.admin.dto.response.*;
import com.fbso.platform.admin.service.DashboardService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.UUID;

import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@ExtendWith(MockitoExtension.class)
@DisplayName("DashboardController")
class DashboardControllerTest {

    @Mock private DashboardService service;
    private MockMvc mockMvc;

    @BeforeEach
    void setUp() {
        // Standalone setup — sem Spring Security (testa só o controller)
        mockMvc = MockMvcBuilders.standaloneSetup(new DashboardController(service)).build();
    }

    @Nested
    @DisplayName("GET /api/v1/dashboard/admin/summary")
    class GetSummary {

        @Test
        @DisplayName("deve retornar 200 com indicadores consolidados")
        void shouldReturnSummary() throws Exception {
            var response = DashboardSummaryResponse.of(
                    50, 35, 10, 5,
                    Map.of("Básico", 25, "Avançado", 10),
                    new BigDecimal("14990.00"), "mes_atual");
            when(service.getSummary()).thenReturn(response);

            mockMvc.perform(get("/api/v1/dashboard/admin/summary")
                            .accept(MediaType.APPLICATION_JSON))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.totalAccounts").value(50))
                    .andExpect(jsonPath("$.activeAccounts").value(35))
                    .andExpect(jsonPath("$.pendingAccounts").value(10))
                    .andExpect(jsonPath("$.suspendedAccounts").value(5))
                    .andExpect(jsonPath("$.monthlyRevenue").value(14990.00))
                    .andExpect(jsonPath("$.period").value("mes_atual"));
        }
    }

    @Nested
    @DisplayName("GET /api/v1/dashboard/admin/evolution")
    class GetEvolution {

        @Test
        @DisplayName("deve retornar 200 com evolução para 30d")
        void shouldReturnEvolution30d() throws Exception {
            var dataPoints = List.of(
                    new EvolutionResponse.DataPoint(
                            java.time.LocalDate.of(2026, 7, 15), 5),
                    new EvolutionResponse.DataPoint(
                            java.time.LocalDate.of(2026, 7, 16), 8));
            var response = EvolutionResponse.of("30d", dataPoints);
            when(service.getEvolution("30d")).thenReturn(response);

            mockMvc.perform(get("/api/v1/dashboard/admin/evolution")
                            .param("period", "30d")
                            .accept(MediaType.APPLICATION_JSON))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.period").value("30d"))
                    .andExpect(jsonPath("$.dataPoints[0].count").value(5))
                    .andExpect(jsonPath("$.dataPoints[1].count").value(8));
        }

        @Test
        @DisplayName("deve usar mes_atual como padrão quando período não informado")
        void shouldDefaultToMesAtual() throws Exception {
            var response = EvolutionResponse.of("mes_atual", List.of());
            when(service.getEvolution("mes_atual")).thenReturn(response);

            mockMvc.perform(get("/api/v1/dashboard/admin/evolution")
                            .accept(MediaType.APPLICATION_JSON))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.period").value("mes_atual"));
        }
    }

    @Nested
    @DisplayName("GET /api/v1/dashboard/admin/accounts-by-status")
    class GetAccountsByStatus {

        @Test
        @DisplayName("deve retornar 200 com distribuição por status")
        void shouldReturnByStatus() throws Exception {
            var statuses = List.of(
                    new AccountsByStatusResponse.StatusCount("ACTIVE", 35),
                    new AccountsByStatusResponse.StatusCount("PENDING_ONBOARDING", 10));
            var response = AccountsByStatusResponse.of(statuses);
            when(service.getAccountsByStatus()).thenReturn(response);

            mockMvc.perform(get("/api/v1/dashboard/admin/accounts-by-status")
                            .accept(MediaType.APPLICATION_JSON))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.statuses[0].status").value("ACTIVE"))
                    .andExpect(jsonPath("$.statuses[0].count").value(35));
        }
    }

    @Nested
    @DisplayName("GET /api/v1/dashboard/admin/accounts-by-plan")
    class GetAccountsByPlan {

        @Test
        @DisplayName("deve retornar 200 com distribuição por plano")
        void shouldReturnByPlan() throws Exception {
            var plans = List.of(
                    new AccountsByPlanResponse.PlanCount("Enterprise", 15),
                    new AccountsByPlanResponse.PlanCount("Básico", 10));
            var response = AccountsByPlanResponse.of(plans);
            when(service.getAccountsByPlan()).thenReturn(response);

            mockMvc.perform(get("/api/v1/dashboard/admin/accounts-by-plan")
                            .accept(MediaType.APPLICATION_JSON))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.plans[0].planName").value("Enterprise"))
                    .andExpect(jsonPath("$.plans[0].count").value(15));
        }
    }

    @Nested
    @DisplayName("GET /api/v1/dashboard/admin/alerts")
    class GetAlerts {

        @Test
        @DisplayName("deve retornar 200 com alertas WARNING e CRITICAL")
        void shouldReturnAlerts() throws Exception {
            UUID tenantId = UUID.randomUUID();
            var alerts = List.of(
                    new AlertResponse.Alert(AlertResponse.AlertType.WARNING,
                            "Onboarding pendente", tenantId, "TENANT"),
                    new AlertResponse.Alert(AlertResponse.AlertType.CRITICAL,
                            "Assinatura suspensa", UUID.randomUUID(), "SUBSCRIPTION"));
            var response = AlertResponse.of(alerts);
            when(service.getAlerts()).thenReturn(response);

            mockMvc.perform(get("/api/v1/dashboard/admin/alerts")
                            .accept(MediaType.APPLICATION_JSON))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.alerts[0].type").value("WARNING"))
                    .andExpect(jsonPath("$.alerts[0].entityType").value("TENANT"))
                    .andExpect(jsonPath("$.alerts[1].type").value("CRITICAL"));
        }

        @Test
        @DisplayName("deve retornar lista vazia quando não há alertas")
        void shouldReturnEmptyAlerts() throws Exception {
            when(service.getAlerts()).thenReturn(AlertResponse.of(List.of()));

            mockMvc.perform(get("/api/v1/dashboard/admin/alerts")
                            .accept(MediaType.APPLICATION_JSON))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.alerts").isEmpty());
        }
    }
}
