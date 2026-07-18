package com.fbso.platform.admin.unit.entity;

import com.fbso.platform.admin.entity.UserPermission;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

@DisplayName("UserPermission (entity)")
class UserPermissionTest {

    @Nested
    @DisplayName("toColumnMap")
    class ToColumnMap {

        @Test
        @DisplayName("deve mapear user_id, business_unit_id e role")
        void shouldMapAllFields() {
            UserPermission up = new UserPermission();
            UUID userId = UUID.randomUUID();
            UUID buId = UUID.randomUUID();
            up.setUserId(userId);
            up.setBusinessUnitId(buId);
            up.setRole("MANAGER_BU");

            Map<String, Object> columns = up.toColumnMap();

            assertThat(columns)
                    .containsEntry("user_id", userId)
                    .containsEntry("business_unit_id", buId)
                    .containsEntry("role", "MANAGER_BU");
        }
    }

    @Nested
    @DisplayName("getId / setId")
    class IdManagement {

        @Test
        @DisplayName("deve persistir ID")
        void shouldPersistId() {
            UserPermission up = new UserPermission();
            UUID id = UUID.randomUUID();
            up.setId(id);
            assertThat(up.getId()).isEqualTo(id);
        }
    }
}
