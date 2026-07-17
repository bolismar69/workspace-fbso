package com.fbso.platform.admin.unit.controller;

import com.fbso.platform.admin.controller.UserController;
import com.fbso.platform.admin.dto.request.UserCreateRequest;
import com.fbso.platform.admin.dto.response.UserResponse;
import com.fbso.platform.admin.exception.DuplicateEmailException;
import com.fbso.platform.admin.exception.GlobalExceptionHandler;
import com.fbso.platform.admin.exception.SelfDeactivationException;
import com.fbso.platform.admin.exception.UserNotFoundException;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.service.UserService;
import org.junit.jupiter.api.AfterEach;
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

import java.time.OffsetDateTime;
import java.util.List;
import java.util.UUID;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.*;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@ExtendWith(MockitoExtension.class)
@DisplayName("UserController")
class UserControllerTest {

    @Mock
    private UserService userService;

    private MockMvc mockMvc;

    private final UUID tenantId = UUID.randomUUID();
    private final UUID userId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        // Standalone setup — sem Spring Security (testa só o controller)
        mockMvc = MockMvcBuilders.standaloneSetup(new UserController(userService))
                .setControllerAdvice(new GlobalExceptionHandler())
                .build();
        TenantContext.set(tenantId, UUID.randomUUID(),
                List.of("ADMIN_TENANT"), List.of(), List.of());
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    @Nested
    @DisplayName("GET /api/v1/users")
    class ListUsers {

        @Test
        @DisplayName("deve retornar 200 com lista de usuários")
        void shouldReturnUserList() throws Exception {
            UserResponse u1 = new UserResponse(userId, "Admin", "admin@fbso.org",
                    "ACTIVE", List.of(), List.of(), null, OffsetDateTime.now());
            when(userService.findAll(tenantId)).thenReturn(List.of(u1));

            mockMvc.perform(get("/api/v1/users"))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$[0].name").value("Admin"))
                    .andExpect(jsonPath("$[0].email").value("admin@fbso.org"));
        }
    }

    @Nested
    @DisplayName("GET /api/v1/users/{id}")
    class GetUserById {

        @Test
        @DisplayName("deve retornar 200 com usuário")
        void shouldReturnUser() throws Exception {
            UserResponse u = new UserResponse(userId, "Admin", "admin@fbso.org",
                    "ACTIVE", List.of(), List.of(), null, OffsetDateTime.now());
            when(userService.findById(userId)).thenReturn(u);

            mockMvc.perform(get("/api/v1/users/" + userId))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.name").value("Admin"));
        }

        @Test
        @DisplayName("deve retornar 404 quando usuário não encontrado")
        void shouldReturn404() throws Exception {
            when(userService.findById(userId)).thenThrow(new UserNotFoundException(userId));

            mockMvc.perform(get("/api/v1/users/" + userId))
                    .andExpect(status().isNotFound());
        }
    }

    @Nested
    @DisplayName("POST /api/v1/users")
    class CreateUser {

        @Test
        @DisplayName("deve retornar 201 com usuário criado")
        void shouldCreateUser() throws Exception {
            UserResponse created = new UserResponse(userId, "Novo", "novo@fbso.org",
                    "INVITE_PENDING", List.of(), List.of(), OffsetDateTime.now(), OffsetDateTime.now());
            when(userService.invite(any(UserCreateRequest.class), eq(tenantId))).thenReturn(created);

            mockMvc.perform(post("/api/v1/users")
                            .contentType(MediaType.APPLICATION_JSON)
                            .content("{\"name\":\"Novo\",\"email\":\"novo@fbso.org\"}"))
                    .andExpect(status().isCreated())
                    .andExpect(jsonPath("$.status").value("INVITE_PENDING"));
        }

        @Test
        @DisplayName("deve retornar 409 quando email duplicado (RN09-02)")
        void shouldReturn409() throws Exception {
            when(userService.invite(any(UserCreateRequest.class), eq(tenantId)))
                    .thenThrow(new DuplicateEmailException("admin@fbso.org"));

            mockMvc.perform(post("/api/v1/users")
                            .contentType(MediaType.APPLICATION_JSON)
                            .content("{\"name\":\"Outro\",\"email\":\"admin@fbso.org\"}"))
                    .andExpect(status().isConflict());
        }

        @Test
        @DisplayName("deve retornar 400 quando body inválido")
        void shouldReturn400() throws Exception {
            mockMvc.perform(post("/api/v1/users")
                            .contentType(MediaType.APPLICATION_JSON)
                            .content("{\"name\":\"\",\"email\":\"invalido\"}"))
                    .andExpect(status().isBadRequest());
        }
    }

    @Nested
    @DisplayName("POST /api/v1/users/{id}/deactivate")
    class DeactivateUser {

        @Test
        @DisplayName("deve retornar 204 ao desativar")
        void shouldReturn204() throws Exception {
            doNothing().when(userService).deactivate(userId);

            mockMvc.perform(post("/api/v1/users/" + userId + "/deactivate"))
                    .andExpect(status().isNoContent());
        }

        @Test
        @DisplayName("deve retornar 422 ao tentar autodesativar (RN09-03)")
        void shouldReturn422() throws Exception {
            doThrow(new SelfDeactivationException()).when(userService).deactivate(userId);

            mockMvc.perform(post("/api/v1/users/" + userId + "/deactivate"))
                    .andExpect(status().isUnprocessableEntity());
        }
    }

    @Nested
    @DisplayName("POST /api/v1/users/{id}/reactivate")
    class ReactivateUser {

        @Test
        @DisplayName("deve retornar 200 com usuário reativado")
        void shouldReturn200() throws Exception {
            UserResponse reactivated = new UserResponse(userId, "Voltou", "voltou@fbso.org",
                    "ACTIVE", List.of(), List.of(), null, OffsetDateTime.now());
            when(userService.reactivate(userId)).thenReturn(reactivated);

            mockMvc.perform(post("/api/v1/users/" + userId + "/reactivate"))
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.status").value("ACTIVE"));
        }
    }
}
