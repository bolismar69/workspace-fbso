package com.fbso.platform.admin.unit.repository;

import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.repository.BusinessUnitRepository;
import com.fbso.platform.admin.repository.rowmapper.BusinessUnitRowMapper;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.util.Collections;
import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

/**
 * Testes unitários para {@link BusinessUnitRepository}.
 *
 * <p>Verifica queries hierárquicas (ADR-L08 — WITH RECURSIVE)
 * e queries de filhos diretos (findChildren).</p>
 */
@ExtendWith(MockitoExtension.class)
class BusinessUnitRepositoryTest {

    @Mock
    private JdbcTemplate jdbcTemplate;

    private BusinessUnitRepository repository;

    private final UUID tenantId = UUID.randomUUID();
    private final UUID parentId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        repository = new BusinessUnitRepository(jdbcTemplate);
    }

    @Test
    void findTree_shouldReturnHierarchicalTree() {
        // Arrange
        BusinessUnit matrix = createBu("Matriz SA", null);
        BusinessUnit filialA = createBu("Filial A", parentId);
        List<BusinessUnit> expected = List.of(matrix, filialA);

        when(jdbcTemplate.query(anyString(), any(RowMapper.class), eq(tenantId)))
                .thenReturn(expected);

        // Act
        List<BusinessUnit> result = repository.findTree(tenantId);

        // Assert
        assertThat(result).hasSize(2);
        assertThat(result.get(0).getCorporateName()).isEqualTo("Matriz SA");
        assertThat(result.get(1).getCorporateName()).isEqualTo("Filial A");

        // Verify SQL uses WITH RECURSIVE CTE
        ArgumentCaptor<String> sqlCaptor = ArgumentCaptor.forClass(String.class);
        verify(jdbcTemplate).query(sqlCaptor.capture(), any(RowMapper.class), eq(tenantId));
        assertThat(sqlCaptor.getValue()).contains("WITH RECURSIVE bu_tree");
        assertThat(sqlCaptor.getValue()).contains("UNION ALL");
    }

    @Test
    void findTree_shouldReturnEmptyListWhenNoBUs() {
        // Arrange
        when(jdbcTemplate.query(anyString(), any(RowMapper.class), eq(tenantId)))
                .thenReturn(Collections.emptyList());

        // Act
        List<BusinessUnit> result = repository.findTree(tenantId);

        // Assert
        assertThat(result).isEmpty();
    }

    @Test
    void findTree_shouldReturnDeepHierarchy() {
        // Arrange — 3 níveis: Matriz → Filial → Sub-filial
        BusinessUnit matrix = createBu("Matriz SA", null);
        BusinessUnit filial = createBu("Filial Norte", UUID.randomUUID());
        BusinessUnit subFilial = createBu("Sub-Filial Norte-1", UUID.randomUUID());
        List<BusinessUnit> expected = List.of(matrix, filial, subFilial);

        when(jdbcTemplate.query(anyString(), any(RowMapper.class), eq(tenantId)))
                .thenReturn(expected);

        // Act
        List<BusinessUnit> result = repository.findTree(tenantId);

        // Assert
        assertThat(result).hasSize(3);
    }

    @Test
    void findChildren_shouldReturnDirectChildren() {
        // Arrange
        BusinessUnit filialA = createBu("Filial A", parentId);
        BusinessUnit filialB = createBu("Filial B", parentId);
        List<BusinessUnit> expected = List.of(filialA, filialB);

        when(jdbcTemplate.query(anyString(), any(RowMapper.class), eq(tenantId), eq(parentId)))
                .thenReturn(expected);

        // Act
        List<BusinessUnit> result = repository.findChildren(parentId, tenantId);

        // Assert
        assertThat(result).hasSize(2);

        ArgumentCaptor<String> sqlCaptor = ArgumentCaptor.forClass(String.class);
        verify(jdbcTemplate).query(sqlCaptor.capture(), any(RowMapper.class), eq(tenantId), eq(parentId));
        assertThat(sqlCaptor.getValue()).contains("tenant_id = ?");
        assertThat(sqlCaptor.getValue()).contains("parent_id = ?");
    }

    @Test
    void findChildren_shouldReturnRootBUsWhenParentIdIsNull() {
        // Arrange
        BusinessUnit matrix = createBu("Matriz SA", null);
        List<BusinessUnit> expected = List.of(matrix);

        when(jdbcTemplate.query(anyString(), any(RowMapper.class), eq(tenantId)))
                .thenReturn(expected);

        // Act
        List<BusinessUnit> result = repository.findChildren(null, tenantId);

        // Assert
        assertThat(result).hasSize(1);
        assertThat(result.get(0).getCorporateName()).isEqualTo("Matriz SA");

        ArgumentCaptor<String> sqlCaptor = ArgumentCaptor.forClass(String.class);
        verify(jdbcTemplate).query(sqlCaptor.capture(), any(RowMapper.class), eq(tenantId));
        assertThat(sqlCaptor.getValue()).contains("tenant_id = ?");
        assertThat(sqlCaptor.getValue()).contains("parent_id IS NULL");
    }

    // -- helpers --

    private BusinessUnit createBu(String corporateName, UUID parentId) {
        BusinessUnit bu = new BusinessUnit();
        bu.setId(UUID.randomUUID());
        bu.setTenantId(tenantId);
        bu.setParentId(parentId);
        bu.setCnpj("11222333000181");
        bu.setCorporateName(corporateName);
        bu.setTaxRegime("SIMPLES_NACIONAL");
        bu.setStatus("ACTIVE");
        bu.setMatrix(parentId == null);
        return bu;
    }
}
