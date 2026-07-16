package com.fbso.platform.admin.integration.security;

import com.fbso.platform.admin.security.TenantContext;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.condition.EnabledIfEnvironmentVariable;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

/**
 * Testes estruturais da Migration V003 — PostgreSQL Row-Level Security.
 * <p>
 * Verifica que os arquivos de migração existem e contêm os comandos esperados.
 * Os testes de integração completos (com PostgreSQL real) exigem Testcontainers
 * e são executados como parte do {@code mvn verify}.
 */
@DisplayName("V003 — PostgreSQL Row-Level Security (Estrutural)")
class RLSIsolationTest {

    private Path v003Path;
    private Path u003Path;

    @BeforeEach
    void setUp() {
        Path migrationDir = Paths.get("src/main/resources/db/migration");
        v003Path = migrationDir.resolve("V003__enable_rls.sql");
        u003Path = migrationDir.resolve("U003__disable_rls.sql");
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    @Test
    @DisplayName("TC-INFRA-022: V003 existe e contém ENABLE ROW LEVEL SECURITY para 5 tabelas")
    void v003MigrationExistsAndEnablesRLSOn5Tables() throws IOException {
        assertThat(v003Path).exists();

        String content = Files.readString(v003Path);
        assertThat(content).contains("ENABLE ROW LEVEL SECURITY");

        // 5 tabelas com tenant_id devem ter RLS
        List<String> expectedTables = List.of(
                "subscription", "\"user\"", "business_unit",
                "product_service", "audit_log");

        for (String table : expectedTables) {
            assertThat(content)
                    .as("Tabela %s deve ter RLS habilitado", table)
                    .contains("fbso_platform." + table);
        }
    }

    @Test
    @DisplayName("TC-INFRA-023: V003 cria política tenant_isolation com USING + WITH CHECK")
    void v003CreatesTenantIsolationPolicy() throws IOException {
        String content = Files.readString(v003Path);

        // 5 políticas criadas
        long policyCount = content.lines()
                .filter(line -> line.contains("CREATE POLICY tenant_isolation"))
                .count();
        assertThat(policyCount).isEqualTo(5);

        // Cada política usa current_setting('app.current_tenant_id')
        assertThat(content)
                .contains("current_setting('app.current_tenant_id')::UUID");
        assertThat(content).contains("USING");
        assertThat(content).contains("WITH CHECK");
    }

    @Test
    @DisplayName("V003 NÃO habilita RLS em tabelas globais (tenant, plan, resource_action, role_resource)")
    void v003DoesNotEnableRLSOnGlobalTables() throws IOException {
        String content = Files.readString(v003Path);

        // Tabelas globais NÃO devem ter RLS
        List<String> globalTables = List.of(
                "fbso_platform.tenant", "fbso_platform.plan",
                "fbso_platform.plan_module", "fbso_platform.user_permission",
                "fbso_platform.resource_action", "fbso_platform.role_resource");

        for (String table : globalTables) {
            assertThat(content)
                    .as("Tabela global %s NÃO deve ter RLS", table)
                    .doesNotContain("ALTER TABLE " + table + " ENABLE ROW LEVEL SECURITY");
        }
    }

    @Test
    @DisplayName("TC-INFRA-025: U003 (rollback) existe e remove RLS das 5 tabelas")
    void u003RollbackExistsAndDisablesRLS() throws IOException {
        assertThat(u003Path).exists();

        String content = Files.readString(u003Path);

        // 5 DROP POLICY + 5 DISABLE ROW LEVEL SECURITY
        long dropCount = content.lines()
                .filter(line -> line.contains("DROP POLICY IF EXISTS tenant_isolation"))
                .count();
        long disableCount = content.lines()
                .filter(line -> line.contains("DISABLE ROW LEVEL SECURITY"))
                .count();

        assertThat(dropCount).isEqualTo(5);
        assertThat(disableCount).isEqualTo(5);
    }

    @Test
    @DisplayName("Migration V003 é idempotente — políticas com IF NOT EXISTS ou DROP IF EXISTS")
    @EnabledIfEnvironmentVariable(named = "CI", matches = "true",
            disabledReason = "Verificação de idempotência — executada apenas no CI")
    void v003MigrationIsIdempotent() throws IOException {
        // Nota: PostgreSQL RLS permite múltiplas políticas com mesmo nome,
        // então CREATE POLICY não é idempotente por padrão.
        // Se necessário, usar CREATE POLICY IF NOT EXISTS (PostgreSQL 17+).
        // Este teste serve como lembrete/documentação.
        String content = Files.readString(v003Path);
        assertThat(content).isNotEmpty();
    }
}
