package com.fbso.platform.admin.unit.entity;

import com.fbso.platform.admin.entity.ProductService;
import com.fbso.platform.admin.enums.ProductType;
import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

/**
 * Testes unitários para {@link ProductService} entity.
 *
 * <p><b>DT-130 (Sprint 6 F1):</b> Verifica que tenantId está
 * presente no toColumnMap() após V009 migration.</p>
 */
class ProductServiceTest {

    @Test
    void toColumnMap_shouldIncludeTenantId() {
        // Arrange
        ProductService ps = new ProductService();
        UUID tenantId = UUID.randomUUID();
        UUID businessUnitId = UUID.randomUUID();

        ps.setTenantId(tenantId);
        ps.setBusinessUnitId(businessUnitId);
        ps.setName("Produto Teste");
        ps.setSku("SKU-001");
        ps.setType(ProductType.PRODUCT);
        ps.setDescription("Descrição");
        ps.setStatus("ACTIVE");

        // Act
        Map<String, Object> columns = ps.toColumnMap();

        // Assert
        assertThat(columns).containsKey("tenant_id");
        assertThat(columns.get("tenant_id")).isEqualTo(tenantId);
        assertThat(columns).containsKey("business_unit_id");
        assertThat(columns.get("business_unit_id")).isEqualTo(businessUnitId);
        assertThat(columns).containsKey("name");
        assertThat(columns.get("name")).isEqualTo("Produto Teste");
        assertThat(columns).containsKey("sku");
        assertThat(columns.get("sku")).isEqualTo("SKU-001");
        assertThat(columns).containsKey("type");
        assertThat(columns.get("type")).isEqualTo("PRODUCT");
        assertThat(columns).containsKey("description");
        assertThat(columns).containsKey("status");
    }

    @Test
    void tenantId_shouldBeAccessible() {
        // Arrange
        ProductService ps = new ProductService();
        UUID tenantId = UUID.randomUUID();

        // Act
        ps.setTenantId(tenantId);

        // Assert
        assertThat(ps.getTenantId()).isEqualTo(tenantId);
    }

    @Test
    void defaultConstructor_shouldSetDefaultValues() {
        // Act
        ProductService ps = new ProductService();

        // Assert
        assertThat(ps.getTenantId()).isNull(); // não inicializado
        assertThat(ps.getType()).isEqualTo(ProductType.SERVICE);
        assertThat(ps.getStatus()).isEqualTo("ACTIVE");
    }
}
