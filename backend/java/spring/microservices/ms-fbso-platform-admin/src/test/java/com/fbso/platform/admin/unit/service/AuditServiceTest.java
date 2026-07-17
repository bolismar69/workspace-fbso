package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.entity.AuditEntry;
import com.fbso.platform.admin.repository.AuditRepository;
import com.fbso.platform.admin.service.AuditService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("AuditService")
class AuditServiceTest {

    @Mock private AuditRepository repo;
    private AuditService service;

    @BeforeEach
    void setUp() {
        service = new AuditService(repo);
    }

    @Nested
    @DisplayName("search")
    class Search {

        @Test
        @DisplayName("deve retornar registros com filtros")
        void shouldReturnFilteredResults() {
            AuditEntry entry = new AuditEntry();
            entry.setId(UUID.randomUUID());
            entry.setAction("CREATED");
            entry.setEntityType("TENANT");
            entry.setEntityId(UUID.randomUUID());

            when(repo.findByFilters(any(), any(), any(), any(), anyInt(), anyInt()))
                    .thenReturn(List.of(entry));

            var results = service.search(null, null, null, null, 0, 25);

            assertThat(results).hasSize(1);
            assertThat(results.get(0).action()).isEqualTo("CREATED");
        }

        @Test
        @DisplayName("deve retornar lista vazia quando não há registros")
        void shouldReturnEmptyWhenNoResults() {
            when(repo.findByFilters(any(), any(), any(), any(), anyInt(), anyInt()))
                    .thenReturn(List.of());

            var results = service.search(null, null, null, null, 0, 25);

            assertThat(results).isEmpty();
        }

        @Test
        @DisplayName("deve parsear datas ISO para filtro")
        void shouldParseIsoDates() {
            when(repo.findByFilters(any(), any(), eq("CREATED"), any(), anyInt(), anyInt()))
                    .thenReturn(List.of());

            var results = service.search(
                    "2026-07-01T00:00:00Z", "2026-07-17T00:00:00Z", "CREATED", null, 0, 25);

            assertThat(results).isEmpty();
        }
    }
}
