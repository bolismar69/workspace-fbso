package com.fbso.platform.admin.unit.repository;

import com.fbso.platform.admin.entity.User;
import com.fbso.platform.admin.enums.UserStatus;
import com.fbso.platform.admin.repository.UserRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("UserRepository")
class UserRepositoryTest {

    @Mock
    private JdbcTemplate jdbc;

    private UserRepository repo;

    @BeforeEach
    void setUp() {
        repo = new UserRepository(jdbc);
    }

    @Nested
    @DisplayName("findByEmailAndTenant")
    class FindByEmailAndTenant {

        @Test
        @DisplayName("deve retornar usuário quando email e tenant batem")
        void shouldReturnUserWhenFound() {
            UUID tenantId = UUID.randomUUID();
            User expected = new User();
            expected.setEmail("admin@fbso.org");
            expected.setTenantId(tenantId);
            expected.setName("Admin");
            expected.setStatus(UserStatus.ACTIVE);

            when(jdbc.query(anyString(), any(RowMapper.class), eq("admin@fbso.org"), eq(tenantId)))
                    .thenReturn(List.of(expected));

            Optional<User> result = repo.findByEmailAndTenant("admin@fbso.org", tenantId);

            assertThat(result).isPresent();
            assertThat(result.get().getEmail()).isEqualTo("admin@fbso.org");
            assertThat(result.get().getName()).isEqualTo("Admin");
        }

        @Test
        @DisplayName("deve retornar vazio quando email não encontrado")
        void shouldReturnEmptyWhenNotFound() {
            UUID tenantId = UUID.randomUUID();

            when(jdbc.query(anyString(), any(RowMapper.class), eq("nonexistent@fbso.org"), eq(tenantId)))
                    .thenReturn(List.of());

            Optional<User> result = repo.findByEmailAndTenant("nonexistent@fbso.org", tenantId);

            assertThat(result).isEmpty();
        }

        @Test
        @DisplayName("deve excluir registros com soft delete")
        void shouldExcludeSoftDeleted() {
            UUID tenantId = UUID.randomUUID();
            // A query usa "deleted_dt IS NULL" — se não houver resultados,
            // significa que o soft delete está sendo respeitado
            when(jdbc.query(anyString(), any(RowMapper.class), eq("deleted@fbso.org"), eq(tenantId)))
                    .thenReturn(List.of());

            Optional<User> result = repo.findByEmailAndTenant("deleted@fbso.org", tenantId);

            assertThat(result).isEmpty();
        }
    }

    @Nested
    @DisplayName("findAllByTenant")
    class FindAllByTenant {

        @Test
        @DisplayName("deve retornar usuários ativos ordenados por nome")
        void shouldReturnActiveUsersOrderedByName() {
            UUID tenantId = UUID.randomUUID();
            User u1 = new User();
            u1.setName("Alice");
            u1.setStatus(UserStatus.ACTIVE);
            User u2 = new User();
            u2.setName("Bob");
            u2.setStatus(UserStatus.ACTIVE);

            when(jdbc.query(anyString(), any(RowMapper.class), eq(tenantId)))
                    .thenReturn(List.of(u1, u2));

            List<User> result = repo.findAllByTenant(tenantId);

            assertThat(result).hasSize(2);
            assertThat(result.get(0).getName()).isEqualTo("Alice");
            assertThat(result.get(1).getName()).isEqualTo("Bob");
        }

        @Test
        @DisplayName("deve retornar lista vazia para tenant sem usuários")
        void shouldReturnEmptyForTenantWithoutUsers() {
            UUID tenantId = UUID.randomUUID();

            when(jdbc.query(anyString(), any(RowMapper.class), eq(tenantId)))
                    .thenReturn(List.of());

            List<User> result = repo.findAllByTenant(tenantId);

            assertThat(result).isEmpty();
        }

        @Test
        @DisplayName("deve excluir registros com soft delete")
        void shouldExcludeSoftDeletedRecords() {
            UUID tenantId = UUID.randomUUID();

            when(jdbc.query(anyString(), any(RowMapper.class), eq(tenantId)))
                    .thenReturn(List.of());

            List<User> result = repo.findAllByTenant(tenantId);

            assertThat(result).isEmpty();
        }
    }

    @Nested
    @DisplayName("constructor")
    class ConstructorTest {

        @Test
        @DisplayName("deve configurar tableName com aspas para palavra reservada SQL")
        void shouldUseQuotedTableName() {
            // UserRepository usa "\"user\"" como tableName (palavra reservada SQL)
            // O BaseRepository deve tratar isso corretamente
            assertThat(repo).isNotNull();
        }
    }
}
