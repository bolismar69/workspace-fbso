package com.fbso.platform.admin.unit.entity;

import com.fbso.platform.admin.entity.User;
import com.fbso.platform.admin.enums.UserStatus;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

@DisplayName("User (entity)")
class UserTest {

    @Nested
    @DisplayName("toColumnMap")
    class ToColumnMap {

        @Test
        @DisplayName("deve mapear todos os campos de domínio")
        void shouldMapAllDomainFields() {
            User u = new User();
            UUID tenantId = UUID.randomUUID();
            UUID keycloakId = UUID.randomUUID();
            u.setTenantId(tenantId);
            u.setExternalKeycloakId(keycloakId);
            u.setEmail("admin@fbso.org");
            u.setName("Admin FBSO");
            u.setStatus(UserStatus.INVITE_PENDING);

            Map<String, Object> columns = u.toColumnMap();

            assertThat(columns)
                    .containsEntry("tenant_id", tenantId)
                    .containsEntry("external_keycloak_id", keycloakId)
                    .containsEntry("email", "admin@fbso.org")
                    .containsEntry("name", "Admin FBSO")
                    .containsEntry("status", "INVITE_PENDING");
        }

        @Test
        @DisplayName("deve mapear status como null quando status é null")
        void shouldMapNullStatus() {
            User u = new User();
            u.setStatus(null);

            Map<String, Object> columns = u.toColumnMap();

            assertThat(columns.get("status")).isNull();
        }
    }

    @Nested
    @DisplayName("isInvitePending")
    class IsInvitePending {

        @Test
        @DisplayName("deve retornar true quando status é INVITE_PENDING")
        void shouldReturnTrueForInvitePending() {
            User u = new User();
            u.setStatus(UserStatus.INVITE_PENDING);

            assertThat(u.isInvitePending()).isTrue();
        }

        @Test
        @DisplayName("deve retornar false quando status é ACTIVE")
        void shouldReturnFalseForActive() {
            User u = new User();
            u.setStatus(UserStatus.ACTIVE);

            assertThat(u.isInvitePending()).isFalse();
        }

        @Test
        @DisplayName("deve retornar false quando status é INACTIVE")
        void shouldReturnFalseForInactive() {
            User u = new User();
            u.setStatus(UserStatus.INACTIVE);

            assertThat(u.isInvitePending()).isFalse();
        }
    }

    @Nested
    @DisplayName("getId / setId")
    class IdManagement {

        @Test
        @DisplayName("deve persistir ID via setter e retornar via getter")
        void shouldPersistAndReturnId() {
            User u = new User();
            UUID id = UUID.randomUUID();
            u.setId(id);

            assertThat(u.getId()).isEqualTo(id);
        }
    }
}
