package com.fbso.platform.admin.unit.controller;

import com.fbso.platform.admin.controller.PermissionController;
import com.fbso.platform.admin.dto.response.PermissionResponse;
import com.fbso.platform.admin.exception.GlobalExceptionHandler;
import com.fbso.platform.admin.exception.UserNotFoundException;
import com.fbso.platform.admin.service.PermissionService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;

import java.util.List;
import java.util.UUID;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@ExtendWith(MockitoExtension.class)
@DisplayName("PermissionController")
class PermissionControllerTest {

    @Mock
    private PermissionService permissionService;

    private MockMvc mockMvc;

    private final UUID userId = UUID.randomUUID();
    private final UUID buId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        mockMvc = MockMvcBuilders.standaloneSetup(new PermissionController(permissionService))
                .setControllerAdvice(new GlobalExceptionHandler())
                .build();
    }

    @Nested
    @DisplayName("GET /api/v1/users/{userId}/permissions")
    class ListPermissions {

        @Test
        @DisplayName("deve retornar 200 com lista de permissões")
        void shouldReturnPermissions() throws Exception {
            PermissionResponse p = new PermissionResponse(userId, buId, "MANAGER_BU");
            when(permissionService.getUserPermissions(userId)).thenReturn(List.of(p));

            mockMvc.perform(get("/api/v1/users/" + userId + "/permissions"))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$[0].role").value("MANAGER_BU"))
                    .andExpect(jsonPath("$[0].businessUnitId").value(buId.toString()));
        }

        @Test
        @DisplayName("deve retornar lista vazia quando sem permissões")
        void shouldReturnEmptyList() throws Exception {
            when(permissionService.getUserPermissions(userId)).thenReturn(List.of());

            mockMvc.perform(get("/api/v1/users/" + userId + "/permissions"))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$").isArray())
                    .andExpect(jsonPath("$").isEmpty());
        }

        @Test
        @DisplayName("deve retornar 404 quando usuário não pertence ao tenant")
        void shouldReturn404() throws Exception {
            when(permissionService.getUserPermissions(userId))
                    .thenThrow(new UserNotFoundException(userId));

            mockMvc.perform(get("/api/v1/users/" + userId + "/permissions"))
                    .andExpect(status().isNotFound());
        }
    }

    @Nested
    @DisplayName("PUT /api/v1/users/{userId}/permissions")
    class UpdatePermissions {

        @Test
        @DisplayName("deve retornar 200 com permissões atualizadas")
        void shouldUpdatePermissions() throws Exception {
            PermissionResponse p = new PermissionResponse(userId, buId, "OPERATOR_BU");
            when(permissionService.updateUserPermissions(eq(userId), any()))
                    .thenReturn(List.of(p));

            mockMvc.perform(put("/api/v1/users/" + userId + "/permissions")
                            .contentType(MediaType.APPLICATION_JSON)
                            .content("""
                                    {"permissions":[{"businessUnitId":"%s","role":"OPERATOR_BU"}]}
                                    """.formatted(buId)))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$[0].role").value("OPERATOR_BU"));
        }

        @Test
        @DisplayName("deve retornar 400 quando lista de permissões vazia")
        void shouldReturn400ForEmptyPermissions() throws Exception {
            mockMvc.perform(put("/api/v1/users/" + userId + "/permissions")
                            .contentType(MediaType.APPLICATION_JSON)
                            .content("{\"permissions\":[]}"))
                    .andExpect(status().isBadRequest());
        }

        @Test
        @DisplayName("deve retornar 400 quando role está em branco")
        void shouldReturn400ForBlankRole() throws Exception {
            mockMvc.perform(put("/api/v1/users/" + userId + "/permissions")
                            .contentType(MediaType.APPLICATION_JSON)
                            .content("""
                                    {"permissions":[{"businessUnitId":"%s","role":""}]}
                                    """.formatted(buId)))
                    .andExpect(status().isBadRequest());
        }
    }
}
