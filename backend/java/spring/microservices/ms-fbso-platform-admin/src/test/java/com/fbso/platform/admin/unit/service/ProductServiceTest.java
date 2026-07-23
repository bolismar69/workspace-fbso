package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.request.ProductCreateRequest;
import com.fbso.platform.admin.dto.response.ProductResponse;
import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
import com.fbso.platform.admin.repository.BusinessUnitRepository;
import com.fbso.platform.admin.repository.ProductRepository;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.ProductService;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class ProductServiceTest {

    @Mock private ProductRepository productRepo;
    @Mock private BusinessUnitRepository buRepo;
    private ProductService service;
    private final UUID tenantId = UUID.randomUUID();
    private final UUID buId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        service = new ProductService(productRepo, buRepo);
        TenantContext.set(tenantId, UUID.randomUUID(), List.of("ADMIN_TENANT"), List.of(), List.of());
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    @Test
    void create_shouldCreateProductWithValidData() {
        BusinessUnit bu = new BusinessUnit();
        bu.setId(buId);
        bu.setTenantId(tenantId);
        bu.setStatus("ACTIVE");
        when(buRepo.findById(buId)).thenReturn(Optional.of(bu));
        when(productRepo.existsBySku(any(), any())).thenReturn(false);

        ProductCreateRequest req = new ProductCreateRequest(
                buId, "Produto Teste", "SKU-001", "PRODUCT", "Descrição");

        ProductResponse resp = service.create(req);

        assertThat(resp.name()).isEqualTo("Produto Teste");
        assertThat(resp.sku()).isEqualTo("SKU-001");
        assertThat(resp.type()).isEqualTo("PRODUCT");
        assertThat(resp.fiscalMappingStatus()).isEqualTo("NOT_MAPPED");
        verify(productRepo).save(any());
    }

    @Test
    void create_shouldRejectDuplicateSku() {
        BusinessUnit bu = new BusinessUnit();
        bu.setId(buId);
        bu.setStatus("ACTIVE");
        when(buRepo.findById(buId)).thenReturn(Optional.of(bu));
        when(productRepo.existsBySku("SKU-001", buId)).thenReturn(true);

        ProductCreateRequest req = new ProductCreateRequest(
                buId, "Produto", "SKU-001", "PRODUCT", null);

        assertThatThrownBy(() -> service.create(req))
                .isInstanceOf(BusinessException.class)
                .hasMessageContaining("SKU");
    }

    @Test
    void create_shouldAllowNullSku() {
        BusinessUnit bu = new BusinessUnit();
        bu.setId(buId);
        bu.setStatus("ACTIVE");
        when(buRepo.findById(buId)).thenReturn(Optional.of(bu));

        ProductCreateRequest req = new ProductCreateRequest(
                buId, "Produto sem SKU", null, "SERVICE", null);

        ProductResponse resp = service.create(req);
        assertThat(resp.sku()).isNull();
        assertThat(resp.fiscalMappingStatus()).isEqualTo("NOT_MAPPED"); // RN18-03
    }

    @Test
    void create_shouldRejectNonExistentBu() {
        when(buRepo.findById(any())).thenReturn(Optional.empty());

        ProductCreateRequest req = new ProductCreateRequest(
                UUID.randomUUID(), "Produto", null, "PRODUCT", null);

        assertThatThrownBy(() -> service.create(req))
                .isInstanceOf(TenantNotFoundException.class);
    }

    @Test
    void deactivate_shouldSoftDelete() {
        com.fbso.platform.admin.entity.ProductService ps =
                new com.fbso.platform.admin.entity.ProductService();
        ps.setId(UUID.randomUUID());
        ps.setName("Produto a desativar");
        ps.setStatus("ACTIVE");
        when(productRepo.findById(ps.getId())).thenReturn(Optional.of(ps));

        service.deactivate(ps.getId());
        verify(productRepo).softDelete(eq(ps.getId()), any());
    }

    @Test
    void findById_shouldThrowWhenNotFound() {
        when(productRepo.findById(any())).thenReturn(Optional.empty());
        assertThatThrownBy(() -> service.findById(UUID.randomUUID()))
                .isInstanceOf(TenantNotFoundException.class);
    }
}
