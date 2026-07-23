package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.response.DashboardClientResponse;
import com.fbso.platform.admin.service.DashboardClientService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;

import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("DashboardClientService")
class DashboardClientServiceTest {

    @Mock private JdbcTemplate jdbc;
    private DashboardClientService service;
    private UUID tenantId = UUID.randomUUID();

    @BeforeEach
    void setUp() { service = new DashboardClientService(jdbc); }

    @Test
    @DisplayName("deve retornar dashboard com 4 cards")
    void shouldReturnDashboardWithFourCards() {
        when(jdbc.queryForObject(contains("business_unit"), eq(Integer.class), eq(tenantId))).thenReturn(3);
        when(jdbc.queryForObject(contains("product_service"), eq(Integer.class), eq(tenantId))).thenReturn(12);

        DashboardClientResponse result = service.getSummary(tenantId);

        assertThat(result.activeUnits()).isEqualTo(3);
        assertThat(result.productCount()).isEqualTo(12);
        assertThat(result.activeUnitsLink()).isEqualTo("/business-units");
        assertThat(result.productsLink()).isEqualTo("/products");
        assertThat(result.notifications()).isEmpty();
    }

    @Test
    @DisplayName("deve usar placeholder quando tenant sem assinatura")
    void shouldUsePlaceholderWhenNoSubscription() {
        when(jdbc.queryForObject(contains("business_unit"), eq(Integer.class), eq(tenantId))).thenReturn(0);
        when(jdbc.queryForObject(contains("product_service"), eq(Integer.class), eq(tenantId))).thenReturn(0);
        when(jdbc.queryForList(contains("subscription"), eq(tenantId))).thenReturn(java.util.List.of());

        DashboardClientResponse result = service.getSummary(tenantId);

        assertThat(result.activeUnits()).isZero();
        assertThat(result.planName()).isEqualTo("FBSO Platform");
    }
}
