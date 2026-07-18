package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.request.UserCreateRequest;
import com.fbso.platform.admin.dto.response.UserResponse;
import com.fbso.platform.admin.entity.User;
import com.fbso.platform.admin.enums.UserStatus;
import com.fbso.platform.admin.exception.DuplicateEmailException;
import com.fbso.platform.admin.exception.SelfDeactivationException;
import com.fbso.platform.admin.exception.UserNotFoundException;
import com.fbso.platform.admin.repository.UserRepository;
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

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
@DisplayName("UserService")
class UserServiceTest {

    @Mock
    private UserRepository userRepo;

    private UserService service;

    private final UUID tenantId = UUID.randomUUID();
    private final UUID userId = UUID.randomUUID();

    @BeforeEach
    void setUp() {
        service = new UserService(userRepo);
    }

    @AfterEach
    void tearDown() {
        TenantContext.clear();
    }

    @Nested
    @DisplayName("invite — RN09-01, RN09-02")
    class InviteTests {

        @Test
        @DisplayName("deve criar usuário com status INVITE_PENDING e invitedDt")
        void shouldCreateUserWithInvitePending() {
            UserCreateRequest request = new UserCreateRequest("Admin FBSO", "admin@fbso.org");
            when(userRepo.findByEmailAndTenant("admin@fbso.org", tenantId)).thenReturn(Optional.empty());
            doAnswer(inv -> {
                User u = inv.getArgument(0);
                u.setId(userId);
                return null; // save() returns void
            }).when(userRepo).save(any(User.class));

            UserResponse result = service.invite(request, tenantId);

            assertThat(result.status()).isEqualTo("INVITE_PENDING");
            assertThat(result.name()).isEqualTo("Admin FBSO");
            assertThat(result.email()).isEqualTo("admin@fbso.org");
            assertThat(result.invitedDt()).isNotNull(); // RN09-01
        }

        @Test
        @DisplayName("deve lançar DuplicateEmailException quando email já existe (RN09-02)")
        void shouldThrowWhenEmailExists() {
            UserCreateRequest request = new UserCreateRequest("Outro", "admin@fbso.org");
            User existing = new User();
            existing.setEmail("admin@fbso.org");
            when(userRepo.findByEmailAndTenant("admin@fbso.org", tenantId))
                    .thenReturn(Optional.of(existing));

            assertThatThrownBy(() -> service.invite(request, tenantId))
                    .isInstanceOf(DuplicateEmailException.class);
            verify(userRepo, never()).save(any());
        }
    }

    @Nested
    @DisplayName("deactivate — RN09-03")
    class DeactivateTests {

        @Test
        @DisplayName("deve desativar usuário normalmente")
        void shouldDeactivateUser() {
            TenantContext.set(tenantId, UUID.randomUUID(),
                    List.of("ADMIN_TENANT"), List.of(), List.of());
            User user = new User();
            user.setId(userId);
            user.setTenantId(tenantId);
            user.setEmail("other@fbso.org");
            when(userRepo.findById(userId)).thenReturn(Optional.of(user));

            service.deactivate(userId);

            verify(userRepo).softDelete(eq(userId), any());
        }

        @Test
        @DisplayName("deve lançar SelfDeactivationException ao tentar autodesativar (RN09-03)")
        void shouldThrowWhenSelfDeactivating() {
            TenantContext.set(tenantId, userId, List.of("ADMIN_TENANT"), List.of(), List.of());

            assertThatThrownBy(() -> service.deactivate(userId))
                    .isInstanceOf(SelfDeactivationException.class);
            verify(userRepo, never()).softDelete(any(), any());
        }

        @Test
        @DisplayName("deve lançar UserNotFoundException para usuário inexistente")
        void shouldThrowWhenUserNotFound() {
            when(userRepo.findById(userId)).thenReturn(Optional.empty());

            assertThatThrownBy(() -> service.deactivate(userId))
                    .isInstanceOf(UserNotFoundException.class);
        }
    }

    @Nested
    @DisplayName("reactivate")
    class ReactivateTests {

        @Test
        @DisplayName("deve reativar usuário restaurando status ACTIVE")
        void shouldReactivateUser() {
            TenantContext.set(tenantId, UUID.randomUUID(),
                    List.of("ADMIN_TENANT"), List.of(), List.of());
            User user = new User();
            user.setId(userId);
            user.setTenantId(tenantId);
            user.setEmail("inactive@fbso.org");
            user.setStatus(UserStatus.INACTIVE);
            when(userRepo.findById(userId)).thenReturn(Optional.of(user));

            UserResponse result = service.reactivate(userId);

            assertThat(result.status()).isEqualTo("ACTIVE");
            verify(userRepo).update(user);
        }

        @Test
        @DisplayName("deve lançar UserNotFoundException para usuário inexistente")
        void shouldThrowWhenUserNotFound() {
            when(userRepo.findById(userId)).thenReturn(Optional.empty());

            assertThatThrownBy(() -> service.reactivate(userId))
                    .isInstanceOf(UserNotFoundException.class);
        }
    }

    @Nested
    @DisplayName("findAll")
    class FindAllTests {

        @Test
        @DisplayName("deve retornar lista de usuários ativos")
        void shouldReturnActiveUsers() {
            User u1 = new User();
            u1.setId(UUID.randomUUID());
            u1.setName("Alice");
            u1.setEmail("alice@fbso.org");
            u1.setStatus(UserStatus.ACTIVE);
            User u2 = new User();
            u2.setId(UUID.randomUUID());
            u2.setName("Bob");
            u2.setEmail("bob@fbso.org");
            u2.setStatus(UserStatus.INVITE_PENDING);

            when(userRepo.findAllByTenant(tenantId)).thenReturn(List.of(u1, u2));

            List<UserResponse> result = service.findAll(tenantId);

            assertThat(result).hasSize(2);
            assertThat(result.get(0).name()).isEqualTo("Alice");
            assertThat(result.get(1).name()).isEqualTo("Bob");
        }
    }

    @Nested
    @DisplayName("findById")
    class FindByIdTests {

        @Test
        @DisplayName("deve retornar usuário por ID")
        void shouldReturnUserById() {
            User user = new User();
            user.setId(userId);
            user.setName("Admin");
            user.setEmail("admin@fbso.org");
            user.setStatus(UserStatus.ACTIVE);
            when(userRepo.findById(userId)).thenReturn(Optional.of(user));

            UserResponse result = service.findById(userId);

            assertThat(result.name()).isEqualTo("Admin");
            assertThat(result.email()).isEqualTo("admin@fbso.org");
        }

        @Test
        @DisplayName("deve lançar UserNotFoundException para ID inexistente")
        void shouldThrowWhenNotFound() {
            when(userRepo.findById(userId)).thenReturn(Optional.empty());

            assertThatThrownBy(() -> service.findById(userId))
                    .isInstanceOf(UserNotFoundException.class);
        }
    }
}
