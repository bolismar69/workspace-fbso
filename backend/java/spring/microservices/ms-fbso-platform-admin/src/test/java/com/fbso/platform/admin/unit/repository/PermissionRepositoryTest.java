package com.fbso.platform.admin.unit.repository;

import com.fbso.platform.admin.entity.UserPermission;
import com.fbso.platform.admin.repository.PermissionRepository;
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
import static org.mockito.Mockito.doReturn;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("PermissionRepository")
class PermissionRepositoryTest {

    @Mock
    private JdbcTemplate jdbc;

    private PermissionRepository repo;

    @BeforeEach
    void setUp() {
        repo = new PermissionRepository(jdbc);
    }

    @Nested
    @DisplayName("findByUser")
    class FindByUser {

        @Test
        @DisplayName("deve retornar permissões do usuário")
        void shouldReturnUserPermissions() {
            UUID userId = UUID.randomUUID();
            UserPermission up = new UserPermission();
            up.setUserId(userId);
            up.setRole("MANAGER_BU");

            when(jdbc.query(anyString(), any(RowMapper.class), eq(userId)))
                    .thenReturn(List.of(up));

            List<UserPermission> result = repo.findByUser(userId);

            assertThat(result).hasSize(1);
            assertThat(result.get(0).getRole()).isEqualTo("MANAGER_BU");
        }
    }

    @Nested
    @DisplayName("findByUserAndBu")
    class FindByUserAndBu {

        @Test
        @DisplayName("deve retornar permissão específica")
        void shouldReturnSpecificPermission() {
            UUID userId = UUID.randomUUID();
            UUID buId = UUID.randomUUID();
            UserPermission up = new UserPermission();
            up.setUserId(userId);
            up.setBusinessUnitId(buId);
            up.setRole("OPERATOR_BU");

            when(jdbc.query(anyString(), any(RowMapper.class), eq(userId), eq(buId)))
                    .thenReturn(List.of(up));

            Optional<UserPermission> result = repo.findByUserAndBu(userId, buId);

            assertThat(result).isPresent();
            assertThat(result.get().getRole()).isEqualTo("OPERATOR_BU");
        }

        @Test
        @DisplayName("deve retornar vazio quando não encontrado")
        void shouldReturnEmpty() {
            when(jdbc.query(anyString(), any(RowMapper.class), any(), any()))
                    .thenReturn(List.of());

            Optional<UserPermission> result = repo.findByUserAndBu(
                    UUID.randomUUID(), UUID.randomUUID());

            assertThat(result).isEmpty();
        }
    }

    @Nested
    @DisplayName("assign")
    class Assign {

        @Test
        @DisplayName("deve executar upsert")
        void shouldUpsert() {
            UUID userId = UUID.randomUUID();
            UUID buId = UUID.randomUUID();

            when(jdbc.update(anyString(), any(), any(), any())).thenReturn(1);

            repo.assign(userId, buId, "MANAGER_BU");

            verify(jdbc).update(anyString(), eq(userId), eq(buId), eq("MANAGER_BU"));
        }
    }

    @Nested
    @DisplayName("revoke")
    class Revoke {

        @Test
        @DisplayName("deve executar delete")
        void shouldDelete() {
            UUID userId = UUID.randomUUID();
            UUID buId = UUID.randomUUID();

            doReturn(1).when(jdbc).update(anyString(), eq(userId), eq(buId));

            repo.revoke(userId, buId);

            verify(jdbc).update(anyString(), eq(userId), eq(buId));
        }
    }

    @Nested
    @DisplayName("findRolesByUser")
    class FindRolesByUser {

        @Test
        @DisplayName("deve retornar lista de roles")
        void shouldReturnRoles() {
            UUID userId = UUID.randomUUID();

            when(jdbc.query(anyString(), any(RowMapper.class), eq(userId)))
                    .thenReturn(List.of("MANAGER_BU", "OPERATOR_BU"));

            List<String> roles = repo.findRolesByUser(userId);

            assertThat(roles).containsExactly("MANAGER_BU", "OPERATOR_BU");
        }
    }
}
