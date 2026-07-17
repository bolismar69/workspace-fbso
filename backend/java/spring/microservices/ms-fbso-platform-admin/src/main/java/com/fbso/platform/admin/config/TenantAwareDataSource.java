package com.fbso.platform.admin.config;

import com.fbso.platform.admin.exception.TenantIsolationException;
import com.fbso.platform.admin.security.TenantContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.jdbc.datasource.DelegatingDataSource;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.UUID;

/**
 * DataSource proxy que configura {@code app.current_tenant_id} no PostgreSQL
 * sempre que uma conexão é emprestada do pool HikariCP.
 * <p>
 * Isso garante que o PostgreSQL Row-Level Security (ADR-L07) funcione
 * corretamente mesmo com connection pooling — cada query executada em
 * qualquer conexão do pool terá o tenant_id correto da requisição atual.
 * <p>
 * O valor é lido do {@link TenantContext} (ThreadLocal), que por sua vez
 * é populado pelo {@code JwtAuthenticationFilter} a partir do token JWT.
 *
 * @see <a href="ARCHITECTURE.md#4.3">ARCHITECTURE.md §4.3 — Defesa em Profundidade</a>
 */
public class TenantAwareDataSource extends DelegatingDataSource {

    private static final Logger log = LoggerFactory.getLogger(TenantAwareDataSource.class);

    public TenantAwareDataSource(DataSource targetDataSource) {
        super(targetDataSource);
    }

    @Override
    public Connection getConnection() throws SQLException {
        Connection conn = super.getConnection();
        applyTenantContext(conn);
        return conn;
    }

    @Override
    public Connection getConnection(String username, String password) throws SQLException {
        Connection conn = super.getConnection(username, password);
        applyTenantContext(conn);
        return conn;
    }

    /**
     * Configura a variável de sessão PostgreSQL usada pelas políticas RLS.
     * Se não há tenant_id no contexto (ex: health check ou Admin FBSO global),
     * a variável é resetada para garantir que não haja vazamento de conexão anterior.
     */
    private void applyTenantContext(Connection conn) {
        String tenantId = TenantContext.getTenantIdQuietly();
        try {
            if (tenantId != null && !tenantId.isEmpty()) {
                try (PreparedStatement pstmt = conn.prepareStatement(
                        "SET app.current_tenant_id = ?")) {
                    pstmt.setObject(1, UUID.fromString(tenantId));
                    pstmt.execute();
                }
            } else {
                // Admin FBSO (visão global) ou health check — limpar variável residual
                try (Statement stmt = conn.createStatement()) {
                    stmt.execute("RESET app.current_tenant_id");
                }
            }
        } catch (SQLException e) {
            log.error("Falha ao configurar app.current_tenant_id na conexão: {}", e.getMessage(), e);
            throw new TenantIsolationException(
                    "Falha ao configurar isolamento multi-tenant na conexão", e);
        }
    }
}
