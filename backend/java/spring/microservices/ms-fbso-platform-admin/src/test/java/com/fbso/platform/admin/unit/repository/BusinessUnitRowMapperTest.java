package com.fbso.platform.admin.unit.repository;

import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.repository.rowmapper.AuditFieldsRowMapper;
import com.fbso.platform.admin.repository.rowmapper.BusinessUnitRowMapper;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Timestamp;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.when;

/**
 * Testes unitários para {@link BusinessUnitRowMapper}.
 *
 * <p>Verifica mapeamento de todas as 15 colunas de domínio
 * + 6 colunas de auditoria (via {@link AuditFieldsRowMapper}).</p>
 */
@ExtendWith(MockitoExtension.class)
class BusinessUnitRowMapperTest {

    @Mock
    private ResultSet resultSet;

    private final BusinessUnitRowMapper rowMapper = new BusinessUnitRowMapper();

    @Test
    void mapRow_shouldMapAllDomainFields() throws SQLException {
        // Arrange
        UUID id = UUID.randomUUID();
        UUID tenantId = UUID.randomUUID();
        UUID parentId = UUID.randomUUID();
        OffsetDateTime now = OffsetDateTime.now(ZoneOffset.UTC);

        when(resultSet.getObject("id", UUID.class)).thenReturn(id);
        when(resultSet.getObject("tenant_id", UUID.class)).thenReturn(tenantId);
        when(resultSet.getObject("parent_id", UUID.class)).thenReturn(parentId);
        when(resultSet.getString("cnpj")).thenReturn("11.222.333/0001-81");
        when(resultSet.getString("corporate_name")).thenReturn("Matriz FBSO Ltda");
        when(resultSet.getString("tax_regime")).thenReturn("SIMPLES_NACIONAL");
        when(resultSet.getString("street")).thenReturn("Rua Principal");
        when(resultSet.getString("number")).thenReturn("100");
        when(resultSet.getString("complement")).thenReturn("Sala 1");
        when(resultSet.getString("neighborhood")).thenReturn("Centro");
        when(resultSet.getString("city")).thenReturn("São Paulo");
        when(resultSet.getString("state")).thenReturn("SP");
        when(resultSet.getString("zip_code")).thenReturn("01001-000");
        when(resultSet.getString("status")).thenReturn("ACTIVE");
        when(resultSet.getBoolean("is_matrix")).thenReturn(true);

        // Auditoria
        when(resultSet.getObject("created_dt", OffsetDateTime.class)).thenReturn(now);
        when(resultSet.getObject("updated_dt", OffsetDateTime.class)).thenReturn(now);
        when(resultSet.getObject("created_by", UUID.class)).thenReturn(UUID.randomUUID());
        when(resultSet.getObject("updated_by", UUID.class)).thenReturn(UUID.randomUUID());
        when(resultSet.getObject("deleted_dt", OffsetDateTime.class)).thenReturn(null);
        when(resultSet.getObject("deleted_by", UUID.class)).thenReturn(null);

        // Act
        BusinessUnit bu = rowMapper.mapRow(resultSet, 0);

        // Assert — domain fields
        assertThat(bu.getId()).isEqualTo(id);
        assertThat(bu.getTenantId()).isEqualTo(tenantId);
        assertThat(bu.getParentId()).isEqualTo(parentId);
        assertThat(bu.getCnpj()).isEqualTo("11.222.333/0001-81");
        assertThat(bu.getCorporateName()).isEqualTo("Matriz FBSO Ltda");
        assertThat(bu.getTaxRegime()).isEqualTo("SIMPLES_NACIONAL");
        assertThat(bu.getStreet()).isEqualTo("Rua Principal");
        assertThat(bu.getNumber()).isEqualTo("100");
        assertThat(bu.getComplement()).isEqualTo("Sala 1");
        assertThat(bu.getNeighborhood()).isEqualTo("Centro");
        assertThat(bu.getCity()).isEqualTo("São Paulo");
        assertThat(bu.getState()).isEqualTo("SP");
        assertThat(bu.getZipCode()).isEqualTo("01001-000");
        assertThat(bu.getStatus()).isEqualTo("ACTIVE");
        assertThat(bu.isMatrix()).isTrue();

        // Audit fields (via AuditFieldsRowMapper)
        assertThat(bu.getCreatedDt()).isEqualTo(now);
        assertThat(bu.getUpdatedDt()).isEqualTo(now);
        assertThat(bu.getDeletedDt()).isNull();
    }

    @Test
    void mapRow_shouldMapNonMatrixBranch() throws SQLException {
        // Arrange
        UUID id = UUID.randomUUID();
        OffsetDateTime now = OffsetDateTime.now(ZoneOffset.UTC);
        when(resultSet.getObject("id", UUID.class)).thenReturn(id);
        when(resultSet.getObject("tenant_id", UUID.class)).thenReturn(UUID.randomUUID());
        when(resultSet.getObject("parent_id", UUID.class)).thenReturn(UUID.randomUUID());
        when(resultSet.getString("cnpj")).thenReturn("22.333.444/0001-55");
        when(resultSet.getString("corporate_name")).thenReturn("Filial Sul");
        when(resultSet.getString("tax_regime")).thenReturn("LUCRO_REAL");
        when(resultSet.getString("street")).thenReturn(null);
        when(resultSet.getString("number")).thenReturn(null);
        when(resultSet.getString("complement")).thenReturn(null);
        when(resultSet.getString("neighborhood")).thenReturn(null);
        when(resultSet.getString("city")).thenReturn(null);
        when(resultSet.getString("state")).thenReturn(null);
        when(resultSet.getString("zip_code")).thenReturn(null);
        when(resultSet.getString("status")).thenReturn("ACTIVE");
        when(resultSet.getBoolean("is_matrix")).thenReturn(false);
        when(resultSet.getObject("created_dt", OffsetDateTime.class)).thenReturn(now);
        when(resultSet.getObject("updated_dt", OffsetDateTime.class)).thenReturn(now);
        when(resultSet.getObject("created_by", UUID.class)).thenReturn(null);
        when(resultSet.getObject("updated_by", UUID.class)).thenReturn(null);
        when(resultSet.getObject("deleted_dt", OffsetDateTime.class)).thenReturn(null);
        when(resultSet.getObject("deleted_by", UUID.class)).thenReturn(null);

        // Act
        BusinessUnit bu = rowMapper.mapRow(resultSet, 1);

        // Assert
        assertThat(bu.isMatrix()).isFalse();
        assertThat(bu.getCnpj()).isEqualTo("22.333.444/0001-55");
        assertThat(bu.getCorporateName()).isEqualTo("Filial Sul");
        assertThat(bu.getTaxRegime()).isEqualTo("LUCRO_REAL");
        assertThat(bu.getStreet()).isNull();
        assertThat(bu.getCity()).isNull();
    }
}
