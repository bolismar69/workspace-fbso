package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.request.BusinessUnitCreateRequest;
import com.fbso.platform.admin.dto.request.BusinessUnitUpdateRequest;
import com.fbso.platform.admin.dto.response.BusinessUnitResponse;
import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.exception.DuplicateCnpjException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
import com.fbso.platform.admin.repository.BusinessUnitRepository;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.BusinessUnitService;
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
class BusinessUnitServiceTest {

    @Mock private BusinessUnitRepository buRepo;
    private BusinessUnitService service;
    private final UUID tenantId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        service = new BusinessUnitService(buRepo);
        TenantContext.set(tenantId, UUID.randomUUID(), List.of("ADMIN_TENANT"), List.of(), List.of());
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    @Test
    void create_shouldCreateBuWithValidCnpj() {
        when(buRepo.existsByCnpj(any(), any())).thenReturn(false);

        BusinessUnitCreateRequest req = new BusinessUnitCreateRequest(
                "11222333000181", "Matriz FBSO Ltda", "SIMPLES_NACIONAL",
                null, null, null, null, null, null, null, null, true);

        BusinessUnitResponse resp = service.create(req);

        assertThat(resp.cnpj()).isEqualTo("11222333000181");
        assertThat(resp.corporateName()).isEqualTo("Matriz FBSO Ltda");
        assertThat(resp.isMatrix()).isTrue();
        verify(buRepo).save(any(BusinessUnit.class));
    }

    @Test
    void create_shouldRejectDuplicateCnpj() {
        when(buRepo.existsByCnpj("11222333000181", tenantId)).thenReturn(true);

        BusinessUnitCreateRequest req = new BusinessUnitCreateRequest(
                "11222333000181", "Outra Empresa", "SIMPLES_NACIONAL",
                null, null, null, null, null, null, null, null, false);

        assertThatThrownBy(() -> service.create(req))
                .isInstanceOf(DuplicateCnpjException.class);
    }

    @Test
    void create_shouldRejectInvalidCnpj() {
        BusinessUnitCreateRequest req = new BusinessUnitCreateRequest(
                "00.000.000/0000-00", "Empresa Falsa", "SIMPLES_NACIONAL",
                null, null, null, null, null, null, null, null, false);

        assertThatThrownBy(() -> service.create(req))
                .isInstanceOf(BusinessException.class);
    }

    @Test
    void create_shouldRejectInactiveParent() {
        UUID inactiveParentId = UUID.randomUUID();
        BusinessUnit inactiveParent = new BusinessUnit();
        inactiveParent.setId(inactiveParentId);
        inactiveParent.setStatus("INACTIVE");

        when(buRepo.existsByCnpj(any(), any())).thenReturn(false);
        when(buRepo.findById(inactiveParentId)).thenReturn(Optional.of(inactiveParent));

        BusinessUnitCreateRequest req = new BusinessUnitCreateRequest(
                "11222333000181", "Filial", "SIMPLES_NACIONAL",
                inactiveParentId, null, null, null, null, null, null, null, false);

        assertThatThrownBy(() -> service.create(req))
                .isInstanceOf(BusinessException.class);
    }

    @Test
    void update_shouldNotChangeCnpj() {
        UUID buId = UUID.randomUUID();
        BusinessUnit existing = new BusinessUnit();
        existing.setId(buId);
        existing.setCnpj("11222333000181");
        existing.setCorporateName("Original");
        existing.setTaxRegime("SIMPLES_NACIONAL");
        existing.setStatus("ACTIVE");

        when(buRepo.findById(buId)).thenReturn(Optional.of(existing));

        BusinessUnitUpdateRequest req = new BusinessUnitUpdateRequest(
                "Novo Nome", null, null, null, null, null, null, null, null, null);

        BusinessUnitResponse resp = service.update(buId, req);

        assertThat(resp.cnpj()).isEqualTo("11222333000181"); // CNPJ imutável
        assertThat(resp.corporateName()).isEqualTo("Novo Nome");
    }

    @Test
    void deactivate_shouldSoftDelete() {
        UUID buId = UUID.randomUUID();
        BusinessUnit existing = new BusinessUnit();
        existing.setId(buId);
        existing.setCnpj("11222333000181");
        existing.setStatus("ACTIVE");

        when(buRepo.findById(buId)).thenReturn(Optional.of(existing));

        service.deactivate(buId);

        verify(buRepo).softDelete(eq(buId), any());
    }

    @Test
    void findById_shouldThrowWhenNotFound() {
        when(buRepo.findById(any())).thenReturn(Optional.empty());

        assertThatThrownBy(() -> service.findById(UUID.randomUUID()))
                .isInstanceOf(TenantNotFoundException.class);
    }

    @Test
    void findTree_shouldReturnHierarchy() {
        BusinessUnit matrix = new BusinessUnit();
        matrix.setId(UUID.randomUUID());
        matrix.setCorporateName("Matriz");
        matrix.setCnpj("11222333000181");
        matrix.setStatus("ACTIVE");
        matrix.setTaxRegime("SIMPLES_NACIONAL");

        when(buRepo.findTree(tenantId)).thenReturn(List.of(matrix));

        List<BusinessUnitResponse> tree = service.findTree();
        assertThat(tree).hasSize(1);
        assertThat(tree.get(0).corporateName()).isEqualTo("Matriz");
    }
}
