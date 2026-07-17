package com.fbso.platform.admin.unit.entity;

import com.fbso.platform.admin.entity.RoleResource;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

@DisplayName("RoleResource (entity)")
class RoleResourceTest {

    @Nested
    @DisplayName("toColumnMap")
    class ToColumnMap {

        @Test
        @DisplayName("deve mapear role e resource_action_id")
        void shouldMapRoleAndResourceActionId() {
            RoleResource rr = new RoleResource();
            rr.setRole("MANAGER_BU");
            UUID raId = UUID.randomUUID();
            rr.setResourceActionId(raId);

            Map<String, Object> columns = rr.toColumnMap();

            assertThat(columns)
                    .containsEntry("role", "MANAGER_BU")
                    .containsEntry("resource_action_id", raId);
        }

        @Test
        @DisplayName("deve mapear ADMIN_TENANT role")
        void shouldMapAdminTenantRole() {
            RoleResource rr = new RoleResource();
            rr.setRole("ADMIN_TENANT");
            rr.setResourceActionId(UUID.randomUUID());

            Map<String, Object> columns = rr.toColumnMap();

            assertThat(columns).containsEntry("role", "ADMIN_TENANT");
        }
    }

    @Nested
    @DisplayName("getId / setId")
    class IdManagement {

        @Test
        @DisplayName("deve persistir ID via setter e retornar via getter")
        void shouldPersistAndReturnId() {
            RoleResource rr = new RoleResource();
            UUID id = UUID.randomUUID();
            rr.setId(id);

            assertThat(rr.getId()).isEqualTo(id);
        }
    }
}
