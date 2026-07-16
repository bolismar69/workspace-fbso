package com.fbso.platform.admin.unit.repository;

import com.fbso.platform.admin.common.BaseEntity;
import com.fbso.platform.admin.repository.common.BaseRepository;
import com.fbso.platform.admin.security.TenantContext;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyInt;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("BaseRepository")
class BaseRepositoryTest {

    @Mock
    private JdbcTemplate jdbc;

    @Mock
    private RowMapper<TestEntity> rowMapper;

    private TestRepository repository;

    // ---- Test Entity ----

    static class TestEntity extends BaseEntity {
        private UUID id;
        private String name;
        public UUID getId() { return id; }
        public void setId(UUID id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
    }

    static class TestRepository extends BaseRepository<TestEntity> {
        TestRepository(JdbcTemplate jdbc, RowMapper<TestEntity> rowMapper, boolean hasTenant) {
            super(jdbc, "test_table", rowMapper, hasTenant);
        }
    }

    // ---- Setup / Teardown ----

    @BeforeEach
    void setUp() {
        repository = new TestRepository(jdbc, rowMapper, true);
        TenantContext.set(
                UUID.randomUUID(),        // tenantId
                UUID.randomUUID(),        // userId
                List.of("ADMIN_TENANT"),  // roles
                List.of(),                // businessUnitIds
                List.of()                 // modules
        );
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    // ---- findAll ----

    @Nested
    @DisplayName("findAll")
    class FindAll {

        @Test
        @DisplayName("deve injetar WHERE deleted_dt IS NULL e tenant_id")
        void shouldInjectSoftDeleteAndTenantFilter() {
            when(jdbc.query(anyString(), eq(rowMapper), any(UUID.class), anyInt(), anyInt()))
                    .thenReturn(List.of());

            repository.findAll(0, 25, "created_dt");

            verify(jdbc).query(
                    anyString(),
                    eq(rowMapper),
                    eq(TenantContext.getTenantId()),
                    eq(25),
                    eq(0)
            );
        }
    }

    // ---- findById ----

    @Nested
    @DisplayName("findById")
    class FindById {

        @Test
        @DisplayName("deve retornar entidade quando encontrada")
        void shouldReturnEntityWhenFound() {
            UUID id = UUID.randomUUID();
            TestEntity entity = new TestEntity();
            entity.setId(id);

            when(jdbc.query(anyString(), eq(rowMapper), eq(id), any(UUID.class)))
                    .thenReturn(List.of(entity));

            Optional<TestEntity> result = repository.findById(id);

            assertThat(result).isPresent();
            assertThat(result.get().getId()).isEqualTo(id);
        }

        @Test
        @DisplayName("deve retornar empty quando não encontrada")
        void shouldReturnEmptyWhenNotFound() {
            UUID id = UUID.randomUUID();

            when(jdbc.query(anyString(), eq(rowMapper), eq(id), any(UUID.class)))
                    .thenReturn(List.of());

            Optional<TestEntity> result = repository.findById(id);

            assertThat(result).isEmpty();
        }
    }

    // ---- softDelete ----

    @Nested
    @DisplayName("softDelete")
    class SoftDelete {

        @Test
        @DisplayName("deve setar deleted_dt e deleted_by")
        void shouldSetDeletedFields() {
            UUID id = UUID.randomUUID();
            UUID deletedBy = UUID.randomUUID();

            when(jdbc.update(anyString(), any(OffsetDateTime.class), eq(deletedBy),
                    eq(id), any(UUID.class)))
                    .thenReturn(1);

            repository.softDelete(id, deletedBy);

            verify(jdbc).update(
                    anyString(),
                    any(OffsetDateTime.class),
                    eq(deletedBy),
                    eq(id),
                    any(UUID.class)
            );
        }

        @Test
        @DisplayName("deve lançar exceção quando registro não encontrado")
        void shouldThrowWhenRecordNotFound() {
            UUID id = UUID.randomUUID();
            UUID deletedBy = UUID.randomUUID();

            when(jdbc.update(anyString(), any(OffsetDateTime.class), eq(deletedBy),
                    eq(id), any(UUID.class)))
                    .thenReturn(0);

            assertThatThrownBy(() -> repository.softDelete(id, deletedBy))
                    .isInstanceOf(IllegalStateException.class)
                    .hasMessageContaining("test_table");
        }
    }

    // ---- count ----

    @Nested
    @DisplayName("count")
    class Count {

        @Test
        @DisplayName("deve retornar total de registros ativos com tenant filter")
        void shouldReturnActiveCountWithTenantFilter() {
            when(jdbc.queryForObject(anyString(), eq(Integer.class), any(UUID.class)))
                    .thenReturn(42);

            int result = repository.count();

            assertThat(result).isEqualTo(42);
        }
    }

    // ---- sanitize ----

    @Nested
    @DisplayName("sanitize de coluna")
    class SanitizeColumn {

        @Test
        @DisplayName("deve rejeitar nome de coluna com SQL injection")
        void shouldRejectInvalidColumnName() {
            // nome inválido é passado via findAll que chama sanitizeColumn internamente
            // O teste verifica que caracteres especiais NÃO são aceitos
            assertThatThrownBy(() -> repository.findAll(0, 10, "created_dt; DROP TABLE test;--"))
                    .isInstanceOf(IllegalArgumentException.class)
                    .hasMessageContaining("Nome de coluna");
        }
    }
}
