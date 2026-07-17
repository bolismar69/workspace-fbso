package com.fbso.platform.admin.integration.repository;

import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantSegment;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.integration.BaseIntegrationTest;
import com.fbso.platform.admin.repository.DashboardRepository;
import com.fbso.platform.admin.repository.TenantRepository;
import org.flywaydb.core.Flyway;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.DriverManagerDataSource;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

/**
 * Testes de integração do {@link DashboardRepository} com PostgreSQL real via Testcontainers.
 *
 * <h3>T-023 — 6 cenários de integração</h3>
 * <ul>
 *   <li>TC-F01-01-004: summary com dados reais</li>
 *   <li>TC-F01-01-005: soft-deleted tenants excluídos das métricas</li>
 *   <li>TC-F01-02-003: findAllPaginated com ordenação</li>
 *   <li>TC-F01-02-004: busca textual &lt;3 chars ignora</li>
 *   <li>TC-F01-03-003: alertas — onboarding >48h + assinatura suspensa</li>
 *   <li>TC-F02-04-010: locked_price preserva preço da assinatura original (DT-009)</li>
 * </ul>
 *
 * <p><b>Por que não usa @SpringBootTest?</b> O contexto completo da aplicação
 * carrega {@code SecurityConfig} que tem dependência circular com
 * {@code JwtAuthenticationFilter}. Para testes de repository puro, criamos
 * JdbcTemplate e repositories manualmente a partir do container PostgreSQL
 * gerenciado por {@link BaseIntegrationTest}. Isso é mais rápido e mais isolado.</p>
 *
 * <p>Herda {@link BaseIntegrationTest} que sobe PostgreSQL 17 via Testcontainers
 * e configura Flyway automaticamente.</p>
 *
 * <p><b>Pré-requisito:</b> Docker Engine rodando.</p>
 *
 * @see BaseIntegrationTest
 * @see DashboardRepository
 */
@DisplayName("DashboardRepository — Integração PostgreSQL (T-023)")
class DashboardRepositoryIT extends BaseIntegrationTest {

    private JdbcTemplate jdbc;
    private DashboardRepository dashboardRepo;
    private TenantRepository tenantRepo;

    // IDs fixos usados no seed (referenciados nos asserts)
    private UUID planBasicoId;
    private UUID planAvancadoId;
    private UUID planEnterpriseId;
    private UUID tenant1ActiveId;
    private UUID tenant3PendingId;

    // =========================================================================
    // Setup — cria JdbcTemplate e repositories manualmente (sem Spring Context)
    // =========================================================================

    @BeforeEach
    void setUp() {
        // 1. Criar JdbcTemplate apontando para o PostgreSQL do Testcontainers
        // fbso_test é owner das tabelas → bypass RLS automático (PostgreSQL default)
        var dataSource = new DriverManagerDataSource(
                postgres.getJdbcUrl(), postgres.getUsername(), postgres.getPassword());
        this.jdbc = new JdbcTemplate(dataSource);

        // 2. Executar Flyway migrations (não há Spring Context para auto-executar)
        // Flyway é idempotente — se migrations já foram aplicadas, apenas verifica
        Flyway.configure()
                .dataSource(dataSource)
                .schemas("fbso_platform")
                .locations("classpath:db/migration")
                .load()
                .migrate();

        // 3. Limpar dados de execuções anteriores (ordem reversa das FKs)
        jdbc.update("DELETE FROM fbso_platform.subscription");
        jdbc.update("DELETE FROM fbso_platform.plan");
        jdbc.update("DELETE FROM fbso_platform.tenant");

        // 4. Instanciar repositories manualmente (injeção pura, sem Spring)
        this.dashboardRepo = new DashboardRepository(jdbc);
        this.tenantRepo = new TenantRepository(jdbc);

        // 5. Seed data
        seedData();
    }

    // =========================================================================
    // Seed — 11 tenants com distribuição realista
    // =========================================================================

    private void seedData() {
        OffsetDateTime now = OffsetDateTime.now();

        // Planos (pré-requisito para subscriptions)
        planBasicoId = insertPlan("Plano Básico", new BigDecimal("99.00"), "MONTHLY");
        planAvancadoId = insertPlan("Plano Avançado", new BigDecimal("299.00"), "MONTHLY");
        planEnterpriseId = insertPlan("Plano Enterprise", new BigDecimal("999.00"), "YEARLY");

        // ACTIVE (5)
        tenant1ActiveId = insertTenant("Mercado Silva Ltda", "Mercado Silva",
                TenantStatus.ACTIVE, TenantSegment.RETAIL, now.minusDays(10));
        var tenant2ActiveId = insertTenant("Distribuidora Costa S.A.", "Distribuidora Costa",
                TenantStatus.ACTIVE, TenantSegment.SERVICES, now.minusDays(8));
        var tenant3ActiveId = insertTenant("Farmácia Saúde Total Ltda", "Farmácia Saúde",
                TenantStatus.ACTIVE, TenantSegment.HEALTHCARE, now.minusDays(6));
        var tenant4ActiveId = insertTenant("Indústria Tech Paraná S.A.", "Tech PR",
                TenantStatus.ACTIVE, TenantSegment.INDUSTRY, now.minusDays(4));
        var tenant5ActiveId = insertTenant("Serviços Gerais Bom Jesus Eireli", "Bom Jesus",
                TenantStatus.ACTIVE, TenantSegment.SERVICES, now.minusDays(2));

        // PENDING_ONBOARDING (3) — um deles criado há >48h para teste de alerta
        tenant3PendingId = insertTenant("Onboarding Lento ME", "Onboarding Lento",
                TenantStatus.PENDING_ONBOARDING, TenantSegment.RETAIL, now.minusDays(5));
        insertTenant("Nova Empresa ABC Ltda", "Nova ABC",
                TenantStatus.PENDING_ONBOARDING, TenantSegment.FINANCIAL, now.minusHours(6));
        insertTenant("Startup XYZ Tecnologia S.A.", "Startup XYZ",
                TenantStatus.PENDING_ONBOARDING, TenantSegment.TECHNOLOGY, now.minusHours(2));

        // SUSPENDED (2)
        var tenant4SuspendedId = insertTenant("Comércio Inadimplente Ltda", "Comércio Inadimplente",
                TenantStatus.SUSPENDED, TenantSegment.RETAIL, now.minusDays(15));
        insertTenant("Distribuidora Suspensa Eireli", "Distribuidora Suspensa",
                TenantStatus.SUSPENDED, TenantSegment.EDUCATION, now.minusDays(12));

        // Soft-deleted (1) — NÃO deve aparecer nas métricas (TC-F01-01-005)
        var tenant5DeletedId = insertTenant("Empresa Excluída Ltda", "Empresa Excluída",
                TenantStatus.ACTIVE, TenantSegment.RETAIL, now.minusDays(20));
        softDeleteTenant(tenant5DeletedId);

        // Subscriptions — vinculando tenants ativos a planos (5 ACTIVE)
        insertSubscription(tenant1ActiveId, planBasicoId, "ACTIVE",
                new BigDecimal("99.00"), "MONTHLY", now.minusDays(10));
        insertSubscription(tenant2ActiveId, planAvancadoId, "ACTIVE",
                new BigDecimal("299.00"), "MONTHLY", now.minusDays(8));
        insertSubscription(tenant3ActiveId, planAvancadoId, "ACTIVE",
                new BigDecimal("299.00"), "MONTHLY", now.minusDays(6));
        insertSubscription(tenant4ActiveId, planEnterpriseId, "ACTIVE",
                new BigDecimal("999.00"), "YEARLY", now.minusDays(4));
        insertSubscription(tenant5ActiveId, planBasicoId, "ACTIVE",
                new BigDecimal("99.00"), "MONTHLY", now.minusDays(2));

        // Subscription SUSPENDED — para alerta CRITICAL (TC-F01-03-003)
        insertSubscription(tenant4SuspendedId, planBasicoId, "SUSPENDED",
                new BigDecimal("99.00"), "MONTHLY", now.minusDays(15));

        // Subscription com locked_price diferente (DT-009 — TC-F02-04-010)
        insertSubscription(tenant1ActiveId, planEnterpriseId, "CANCELED",
                new BigDecimal("799.00"), "MONTHLY", now.minusDays(30));
    }

    // =========================================================================
    // TC-F01-01-004: Dashboard summary carrega indicadores corretos
    // =========================================================================

    @Nested
    @DisplayName("TC-F01-01-004 — Dashboard Summary")
    class DashboardSummary {

        @Test
        @DisplayName("countActive — deve contar apenas tenants não deletados (10)")
        void countActiveShouldExcludeSoftDeleted() {
            assertThat(dashboardRepo.countActive()).isEqualTo(10);
        }

        @Test
        @DisplayName("countByStatus — ACTIVE deve excluir soft-deleted (5 = 6 - 1)")
        void countByStatusActiveShouldExcludeSoftDeleted() {
            // 6 ACTIVE inseridos (5 ativos + 1 "Empresa Excluída" para soft-delete)
            // Após soft-delete: 6 - 1 = 5 ACTIVE com deleted_dt IS NULL
            assertThat(dashboardRepo.countByStatus("ACTIVE")).isEqualTo(5);
        }

        @Test
        @DisplayName("countByStatus — PENDING_ONBOARDING deve retornar 3")
        void countByStatusPendingShouldReturn3() {
            assertThat(dashboardRepo.countByStatus("PENDING_ONBOARDING")).isEqualTo(3);
        }

        @Test
        @DisplayName("countByStatus — SUSPENDED deve retornar 2")
        void countByStatusSuspendedShouldReturn2() {
            assertThat(dashboardRepo.countByStatus("SUSPENDED")).isEqualTo(2);
        }

        @Test
        @DisplayName("accountsByPlan — 3 planos: Básico=2, Avançado=2, Enterprise=1")
        void accountsByPlanShouldReturnCorrectDistribution() {
            var plans = dashboardRepo.accountsByPlan();

            assertThat(plans).hasSize(3);
            var counts = toPlanCountMap(plans);
            assertThat(counts)
                    .containsEntry("Plano Básico", 2L)
                    .containsEntry("Plano Avançado", 2L)
                    .containsEntry("Plano Enterprise", 1L);
        }

        @Test
        @DisplayName("monthlyRevenue — deve somar assinaturas ativas (2×99 + 2×299 + 1×999 = 1795)")
        void monthlyRevenueShouldSumActiveSubscriptions() {
            var revenue = dashboardRepo.monthlyRevenue();

            // SUSPENDED e CANCELED não entram na soma
            assertThat(revenue).isEqualByComparingTo(new BigDecimal("1795.00"));
        }
    }

    // =========================================================================
    // TC-F01-01-005: Soft-deleted tenants excluídos das métricas
    // =========================================================================

    @Nested
    @DisplayName("TC-F01-01-005 — Soft Delete Exclui das Métricas")
    class SoftDeleteExclusion {

        @Test
        @DisplayName("accountsByStatus — total = 10 (5 ACTIVE + 3 PENDING + 2 SUSPENDED)")
        void accountsByStatusShouldExcludeSoftDeleted() {
            // Seed: 6 ACTIVE + 3 PENDING + 2 SUSPENDED = 11 total
            // Após soft-delete: 5 ACTIVE + 3 PENDING + 2 SUSPENDED = 10
            var statuses = dashboardRepo.accountsByStatus();

            long total = statuses.stream()
                    .mapToLong(r -> ((Number) r.get("count")).longValue())
                    .sum();
            assertThat(total).isEqualTo(10);

            var counts = toStatusCountMap(statuses);
            assertThat(counts).containsEntry("ACTIVE", 5L)
                             .containsEntry("PENDING_ONBOARDING", 3L)
                             .containsEntry("SUSPENDED", 2L);
        }

        @Test
        @DisplayName("evolutionByDay — total deve ser 10 (soft-deleted excluído)")
        void evolutionByDayShouldExcludeSoftDeleted() {
            var evolution = dashboardRepo.evolutionByDay(OffsetDateTime.now().minusDays(30));

            long total = evolution.stream()
                    .mapToLong(r -> ((Number) r.get("count")).longValue())
                    .sum();
            assertThat(total).isEqualTo(10);
        }
    }

    // =========================================================================
    // TC-F01-02-003: findAllPaginated — paginação, ordenação, filtros
    // =========================================================================

    @Nested
    @DisplayName("TC-F01-02-003 — Paginação e Ordenação")
    class PaginationAndSorting {

        @Test
        @DisplayName("page=0, size=25, sem filtros → 10 tenants")
        void shouldReturnAllNonDeletedTenants() {
            var tenants = tenantRepo.findAllPaginated(0, 25, null, null, null);
            assertThat(tenants).hasSize(10);
        }

        @Test
        @DisplayName("ordenação por created_dt DESC")
        void shouldOrderByCreatedAtDesc() {
            var tenants = tenantRepo.findAllPaginated(0, 25, null, null, null);

            for (int i = 0; i < tenants.size() - 1; i++) {
                assertThat(tenants.get(i).getCreatedDt())
                        .as("Tenant[%d] >= Tenant[%d]", i, i + 1)
                        .isAfterOrEqualTo(tenants.get(i + 1).getCreatedDt());
            }
        }

        @Test
        @DisplayName("filtro status=SUSPENDED → 2 resultados")
        void shouldFilterByStatus() {
            var tenants = tenantRepo.findAllPaginated(0, 25, "SUSPENDED", null, null);
            assertThat(tenants).hasSize(2);
            assertThat(tenants).allMatch(t -> t.getStatus() == TenantStatus.SUSPENDED);
        }

        @Test
        @DisplayName("filtro plano='Plano Básico' → 2 tenants")
        void shouldFilterByPlan() {
            var tenants = tenantRepo.findAllPaginated(0, 25, null, "Plano Básico", null);
            assertThat(tenants).hasSize(2);
        }

        @Test
        @DisplayName("countFiltered deve bater com findAllPaginated para filtro ACTIVE")
        void countFilteredShouldMatchFindAll() {
            int count = tenantRepo.countFiltered("ACTIVE", null, null);
            var tenants = tenantRepo.findAllPaginated(0, 25, "ACTIVE", null, null);
            assertThat(tenants).hasSize(count);
        }
    }

    // =========================================================================
    // TC-F01-02-004: Busca textual < 3 chars ignora
    // =========================================================================

    @Nested
    @DisplayName("TC-F01-02-004 — Busca Textual")
    class TextualSearch {

        @Test
        @DisplayName("Busca 'Me' (2 chars) → ignora, retorna todos os 10")
        void searchWithLessThan3CharsShouldBeIgnored() {
            var tenants = tenantRepo.findAllPaginated(0, 25, null, null, "Me");
            assertThat(tenants).hasSize(10);
        }

        @Test
        @DisplayName("Busca 'Mercado' (≥3 chars) → 1 resultado exato")
        void searchWith3OrMoreCharsShouldFilter() {
            var tenants = tenantRepo.findAllPaginated(0, 25, null, null, "Mercado");
            assertThat(tenants).hasSize(1);
            assertThat(tenants.get(0).getNameCorporate()).isEqualTo("Mercado Silva Ltda");
        }

        @Test
        @DisplayName("Busca 'Distribuidora' → 2 tenants (ACTIVE + SUSPENDED)")
        void searchShouldMatchBothActiveAndSuspended() {
            var tenants = tenantRepo.findAllPaginated(0, 25, null, null, "Distribuidora");
            assertThat(tenants).hasSize(2);
            assertThat(tenants).extracting(Tenant::getNameCorporate)
                    .containsExactlyInAnyOrder(
                            "Distribuidora Costa S.A.",
                            "Distribuidora Suspensa Eireli");
        }

        @Test
        @DisplayName("ILIKE — 'mercado' e 'MERCADO' retornam o mesmo")
        void searchShouldBeCaseInsensitive() {
            var lower = tenantRepo.findAllPaginated(0, 25, null, null, "mercado");
            var upper = tenantRepo.findAllPaginated(0, 25, null, null, "MERCADO");

            assertThat(lower).hasSize(1);
            assertThat(upper).hasSize(1);
            assertThat(lower.get(0).getId()).isEqualTo(upper.get(0).getId());
        }
    }

    // =========================================================================
    // TC-F01-03-003: Alertas — onboarding >48h + assinatura suspensa
    // =========================================================================

    @Nested
    @DisplayName("TC-F01-03-003 — Alertas do Dashboard")
    class DashboardAlerts {

        @Test
        @DisplayName("onboardingStalled — 1 tenant PENDING há 5 dias → WARNING")
        void onboardingStalledShouldDetectTenantPendingForOver48h() {
            var stalled = dashboardRepo.onboardingStalled();

            assertThat(stalled).hasSize(1);
            assertThat(stalled.get(0).get("name_corporate")).isEqualTo("Onboarding Lento ME");
            assertThat(stalled.get(0).get("status")).isEqualTo("PENDING_ONBOARDING");
        }

        @Test
        @DisplayName("onboardingStalled — tenants PENDING há 6h e 2h NÃO aparecem")
        void onboardingStalledShouldNotIncludeRecentPending() {
            var stalled = dashboardRepo.onboardingStalled();

            // Apenas 1 de 3 PENDING foi criado há >48h
            assertThat(stalled).hasSize(1);
        }

        @Test
        @DisplayName("suspendedSubscriptions — 1 resultado: Comércio Inadimplente / Plano Básico")
        void suspendedSubscriptionsShouldReturnSuspendedWithDetails() {
            var suspended = dashboardRepo.suspendedSubscriptions();

            assertThat(suspended).hasSize(1);
            assertThat(suspended.get(0).get("name_corporate"))
                    .isEqualTo("Comércio Inadimplente Ltda");
            assertThat(suspended.get(0).get("plan_name")).isEqualTo("Plano Básico");
            assertThat(suspended.get(0).get("status")).isEqualTo("SUSPENDED");
        }

        @Test
        @DisplayName("suspendedSubscriptions — CANCELED não conta como suspensa")
        void suspendedSubscriptionsShouldNotIncludeCanceled() {
            var suspended = dashboardRepo.suspendedSubscriptions();
            assertThat(suspended).hasSize(1);
        }
    }

    // =========================================================================
    // TC-F02-04-010: locked_price preserva preço original (DT-009)
    // =========================================================================

    @Nested
    @DisplayName("TC-F02-04-010 — locked_price (DT-009)")
    class LockedPricePreservation {

        @Test
        @DisplayName("monthlyRevenue — locked_price=149.00 prevalece sobre price=99.00 do plano")
        void monthlyRevenueShouldUseLockedPriceWhenPresent() {
            // Insere subscription ACTIVE com locked_price=149.00 (plano tem 99.00)
            var testTenantId = insertTenant("Test Locked Price Ltda", "Test Locked",
                    TenantStatus.ACTIVE, TenantSegment.TECHNOLOGY, OffsetDateTime.now().minusDays(1));
            insertSubscription(testTenantId, planBasicoId, "ACTIVE",
                    new BigDecimal("149.00"), "MONTHLY", OffsetDateTime.now().minusDays(1));

            var revenue = dashboardRepo.monthlyRevenue();

            // 1795.00 (anteriores) + 149.00 (locked_price, não 99.00)
            assertThat(revenue).isEqualByComparingTo(new BigDecimal("1944.00"));
        }

        @Test
        @DisplayName("V005 — colunas locked_price e locked_recurrence existem")
        void lockedPriceColumnsExist() {
            var columns = jdbc.queryForList("""
                SELECT column_name
                FROM information_schema.columns
                WHERE table_schema = 'fbso_platform'
                  AND table_name = 'subscription'
                  AND column_name IN ('locked_price', 'locked_recurrence')
                ORDER BY column_name
                """);

            assertThat(columns).hasSize(2);
            assertThat(columns).extracting(r -> r.get("column_name"))
                    .containsExactlyInAnyOrder("locked_price", "locked_recurrence");
        }
    }

    // =========================================================================
    // Helpers — inserção de dados seed (JDBC direto)
    // =========================================================================

    private UUID insertTenant(String nameCorporate, String nameFantasy,
                              TenantStatus status, TenantSegment segment,
                              OffsetDateTime createdDt) {
        var id = UUID.randomUUID();
        jdbc.update("""
            INSERT INTO fbso_platform.tenant
                (id, name_corporate, name_fantasy, segment, status, created_dt, updated_dt)
            VALUES (?::uuid, ?, ?, ?, ?, ?, ?)
            """, id.toString(), nameCorporate, nameFantasy, segment.name(), status.name(),
                createdDt, createdDt);
        return id;
    }

    private UUID insertPlan(String name, BigDecimal price, String recurrence) {
        var id = UUID.randomUUID();
        jdbc.update("""
            INSERT INTO fbso_platform.plan
                (id, name, description, price, recurrence, status, version, created_dt, updated_dt)
            VALUES (?::uuid, ?, ?, ?::numeric, ?, 'ACTIVE', 1, NOW(), NOW())
            """, id.toString(), name, "Descrição do " + name, price.toString(), recurrence);
        return id;
    }

    private void insertSubscription(UUID tenantId, UUID planId, String status,
                                     BigDecimal lockedPrice, String lockedRecurrence,
                                     OffsetDateTime startDate) {
        jdbc.update("""
            INSERT INTO fbso_platform.subscription
                (id, tenant_id, plan_id, start_date, status,
                 locked_price, locked_recurrence, created_dt, updated_dt)
            VALUES (?::uuid, ?::uuid, ?::uuid, ?, ?,
                    ?::numeric, ?, ?, ?)
            """, UUID.randomUUID().toString(), tenantId.toString(), planId.toString(),
                startDate, status,
                lockedPrice.toString(), lockedRecurrence, startDate, startDate);
    }

    private void softDeleteTenant(UUID tenantId) {
        jdbc.update("""
            UPDATE fbso_platform.tenant
            SET deleted_dt = NOW(), deleted_by = ?::uuid
            WHERE id = ?::uuid
            """, "00000000-0000-0000-0000-000000000000", tenantId.toString());
    }

    // ---- Utilitários para asserts ----

    private static Map<String, Long> toPlanCountMap(List<Map<String, Object>> rows) {
        var result = new java.util.LinkedHashMap<String, Long>();
        for (var row : rows) {
            result.put((String) row.get("plan_name"),
                    ((Number) row.get("count")).longValue());
        }
        return result;
    }

    private static Map<String, Long> toStatusCountMap(List<Map<String, Object>> rows) {
        var result = new java.util.LinkedHashMap<String, Long>();
        for (var row : rows) {
            result.put((String) row.get("status"),
                    ((Number) row.get("count")).longValue());
        }
        return result;
    }
}
