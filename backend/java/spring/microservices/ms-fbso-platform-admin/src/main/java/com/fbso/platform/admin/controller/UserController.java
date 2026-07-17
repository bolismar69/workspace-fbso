package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.UserCreateRequest;
import com.fbso.platform.admin.dto.request.UserUpdateRequest;
import com.fbso.platform.admin.dto.response.UserResponse;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.UserService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.net.URI;
import java.util.List;
import java.util.UUID;

/**
 * Endpoints de gestão de Usuários (F03-01).
 *
 * <p>6 endpoints REST:
 * <ul>
 *   <li>GET /api/v1/users — lista paginada</li>
 *   <li>GET /api/v1/users/{id} — busca por ID</li>
 *   <li>POST /api/v1/users — convidar (status INVITE_PENDING)</li>
 *   <li>PATCH /api/v1/users/{id} — editar dados</li>
 *   <li>POST /api/v1/users/{id}/deactivate — desativar (soft delete)</li>
 *   <li>POST /api/v1/users/{id}/reactivate — reativar</li>
 * </ul>
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN09-02: Email único — validado no {@code UserService.invite()}</li>
 *   <li>RN09-03: Admin não desativa a si mesmo — validado no {@code UserService.deactivate()}</li>
 * </ul>
 */
@RestController
@RequestMapping("/api/v1/users")
@Tag(name = "Users", description = "Gestão de usuários — Admin Tenant")
public class UserController {

    private static final Logger log = LoggerFactory.getLogger(UserController.class);

    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @GetMapping
    @Operation(summary = "Listar usuários do tenant")
    @RequiresPermission(resource = "USER", action = "view")
    public ResponseEntity<List<UserResponse>> list() {
        UUID tenantId = TenantContext.getTenantId();
        log.debug("Listando usuários do tenant {}", tenantId);
        List<UserResponse> users = userService.findAll(tenantId);
        return ResponseEntity.ok(users);
    }

    @GetMapping("/{id}")
    @Operation(summary = "Buscar usuário por ID")
    @RequiresPermission(resource = "USER", action = "view")
    public ResponseEntity<UserResponse> getById(@PathVariable UUID id) {
        log.debug("Buscando usuário {}", id);
        return ResponseEntity.ok(userService.findById(id));
    }

    @PostMapping
    @Operation(summary = "Convidar novo usuário")
    @RequiresPermission(resource = "USER", action = "create")
    public ResponseEntity<UserResponse> create(@Valid @RequestBody UserCreateRequest request) {
        UUID tenantId = TenantContext.getTenantId();
        log.info("Convite de usuário: email={}, tenant={}", request.email(), tenantId);
        UserResponse created = userService.invite(request, tenantId);
        return ResponseEntity.created(URI.create("/api/v1/users/" + created.id()))
                .body(created);
    }

    @PatchMapping("/{id}")
    @Operation(summary = "Editar dados do usuário")
    @RequiresPermission(resource = "USER", action = "edit")
    public ResponseEntity<UserResponse> update(
            @PathVariable UUID id,
            @Valid @RequestBody UserUpdateRequest request) {
        log.info("Editando usuário {}", id);
        UserResponse current = userService.findById(id);

        // Atualização parcial — apenas campos informados
        com.fbso.platform.admin.entity.User entity =
                new com.fbso.platform.admin.entity.User();
        entity.setId(id);
        if (request.name() != null) entity.setName(request.name());
        if (request.email() != null) entity.setEmail(request.email());
        // reusar status atual
        if (current.status() != null) {
            entity.setStatus(com.fbso.platform.admin.enums.UserStatus.valueOf(current.status()));
        }

        // TODO Frente 2: mover lógica de update para UserService.update()
        return ResponseEntity.ok(current); // placeholder — retorna estado atual
    }

    @PostMapping("/{id}/deactivate")
    @Operation(summary = "Desativar usuário (soft delete)")
    @RequiresPermission(resource = "USER", action = "delete")
    public ResponseEntity<Void> deactivate(@PathVariable UUID id) {
        log.info("Desativando usuário {}", id);
        userService.deactivate(id);
        return ResponseEntity.noContent().build();
    }

    @PostMapping("/{id}/reactivate")
    @Operation(summary = "Reativar usuário")
    @RequiresPermission(resource = "USER", action = "edit")
    public ResponseEntity<UserResponse> reactivate(@PathVariable UUID id) {
        log.info("Reativando usuário {}", id);
        return ResponseEntity.ok(userService.reactivate(id));
    }
}
