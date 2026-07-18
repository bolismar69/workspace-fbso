package com.fbso.platform.admin.unit.entity;

import com.fbso.platform.admin.entity.ResourceAction;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

@DisplayName("ResourceAction (entity)")
class ResourceActionTest {

    @Nested
    @DisplayName("toColumnMap")
    class ToColumnMap {

        @Test
        @DisplayName("deve mapear resource_name e action")
        void shouldMapResourceAndAction() {
            ResourceAction ra = new ResourceAction();
            ra.setResourceName("DASHBOARD");
            ra.setAction("view");

            Map<String, Object> columns = ra.toColumnMap();

            assertThat(columns)
                    .containsEntry("resource_name", "DASHBOARD")
                    .containsEntry("action", "view");
        }
    }

    @Nested
    @DisplayName("getId / setId")
    class IdManagement {

        @Test
        @DisplayName("deve persistir ID via setter e retornar via getter")
        void shouldPersistAndReturnId() {
            ResourceAction ra = new ResourceAction();
            UUID id = UUID.randomUUID();
            ra.setId(id);

            assertThat(ra.getId()).isEqualTo(id);
        }
    }
}
