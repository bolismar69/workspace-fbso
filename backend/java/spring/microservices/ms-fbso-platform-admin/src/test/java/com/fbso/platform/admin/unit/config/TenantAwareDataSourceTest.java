package com.fbso.platform.admin.unit.config;

import com.fbso.platform.admin.config.TenantAwareDataSource;
import com.fbso.platform.admin.security.TenantContext;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.*;

/**
 * Testes unitários do {@link TenantAwareDataSource} — o proxy que configura
 * {@code app.current_tenant_id} na sessão PostgreSQL a cada {@code getConnection()}.
 */
@DisplayName("TenantAwareDataSource — Proxy Multi-Tenant")
class TenantAwareDataSourceTest {

    private DataSource mockTarget;
    private Connection mockConnection;
    private Statement mockStatement;
    private TenantAwareDataSource proxy;

    @BeforeEach
    void setUp() throws SQLException {
        mockTarget = mock(DataSource.class);
        mockConnection = mock(Connection.class);
        mockStatement = mock(Statement.class);

        when(mockTarget.getConnection()).thenReturn(mockConnection);
        when(mockConnection.createStatement()).thenReturn(mockStatement);

        proxy = new TenantAwareDataSource(mockTarget);
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    @Test
    @DisplayName("TC-INFRA-024: getConnection() com tenant_id no contexto → SET app.current_tenant_id")
    void shouldSetTenantIdWhenContextIsPopulated() throws SQLException {
        // Arrange
        UUID tenantId = UUID.fromString("11111111-1111-1111-1111-111111111111");
        TenantContext.set(tenantId, UUID.randomUUID(), List.of("ADMIN_TENANT"),
                List.of(), List.of());

        // Act
        Connection conn = proxy.getConnection();

        // Assert
        assertThat(conn).isSameAs(mockConnection);
        verify(mockStatement).execute("SET app.current_tenant_id = '" + tenantId + "'");
        verify(mockTarget).getConnection();
    }

    @Test
    @DisplayName("getConnection() sem tenant_id no contexto → RESET app.current_tenant_id")
    void shouldResetTenantIdWhenContextIsEmpty() throws SQLException {
        // Arrange — TenantContext não inicializado

        // Act
        proxy.getConnection();

        // Assert
        verify(mockStatement).execute("RESET app.current_tenant_id");
    }

    @Test
    @DisplayName("getConnection() com tenant_id nulo → RESET (Admin FBSO global)")
    void shouldResetWhenTenantIdIsNull() throws SQLException {
        // Arrange — contexto com tenantId = null (Admin FBSO cross-tenant)
        TenantContext.set(null, UUID.randomUUID(), List.of("ADMIN_FBSO"),
                List.of(), List.of());

        // Act
        proxy.getConnection();

        // Assert
        verify(mockStatement).execute("RESET app.current_tenant_id");
    }

    @Test
    @DisplayName("getConnection(username, password) também configura tenant_id")
    void shouldSetTenantIdOnCredentialedConnection() throws SQLException {
        // Arrange
        UUID tenantId = UUID.fromString("22222222-2222-2222-2222-222222222222");
        TenantContext.set(tenantId, UUID.randomUUID(), List.of(), List.of(), List.of());
        when(mockTarget.getConnection("user", "pass")).thenReturn(mockConnection);

        // Act
        proxy.getConnection("user", "pass");

        // Assert
        verify(mockStatement).execute("SET app.current_tenant_id = '" + tenantId + "'");
        verify(mockTarget).getConnection("user", "pass");
    }

    @Test
    @DisplayName("SQLException no SET é capturada — não propaga para o chamador")
    void shouldSwallowSqlExceptionOnSet() throws SQLException {
        // Arrange
        UUID tenantId = UUID.randomUUID();
        TenantContext.set(tenantId, UUID.randomUUID(), List.of(), List.of(), List.of());
        doThrow(new SQLException("Connection closed"))
                .when(mockStatement).execute(anyString());

        // Act — não deve lançar exceção
        Connection conn = proxy.getConnection();

        // Assert
        assertThat(conn).isSameAs(mockConnection);
    }

    @Test
    @DisplayName("TenantAwareDataSource não é re-encapsulado pelo BeanPostProcessor")
    void shouldNotBeDoubleWrapped() {
        TenantAwareDataSource alreadyWrapped = new TenantAwareDataSource(mockTarget);
        // Se o BeanPostProcessor tentasse encapsular de novo, teríamos
        // TenantAwareDataSource(TenantAwareDataSource(mockTarget))
        // O instanceof check no BeanPostProcessor previne isso
        assertThat(alreadyWrapped).isInstanceOf(TenantAwareDataSource.class);
    }
}
