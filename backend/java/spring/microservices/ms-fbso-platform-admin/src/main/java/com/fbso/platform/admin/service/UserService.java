package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.UserCreateRequest;
import com.fbso.platform.admin.dto.response.UserResponse;
import com.fbso.platform.admin.entity.User;
import com.fbso.platform.admin.enums.UserStatus;
import com.fbso.platform.admin.exception.DuplicateEmailException;
import com.fbso.platform.admin.exception.SelfDeactivationException;
import com.fbso.platform.admin.exception.UserNotFoundException;
import com.fbso.platform.admin.repository.UserRepository;
import com.fbso.platform.admin.security.TenantContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.UUID;

/**
 * Serviço de gestão de Usuários (F03-01).
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN09-01: Convite expira em 7 dias (campo {@code invited_dt})</li>
 *   <li>RN09-02: Email único por tenant ativo — índice parcial {@code unique_email_active}</li>
 *   <li>RN09-03: Admin não pode desativar a si mesmo</li>
 * </ul>
 */
@Service
public class UserService {

    private static final Logger log = LoggerFactory.getLogger(UserService.class);

    private final UserRepository userRepo;

    public UserService(UserRepository userRepo) {
        this.userRepo = userRepo;
    }

    /**
     * Convida um novo usuário para o tenant.
     *
     * <p>RN09-01: O convite expira em 7 dias a partir de {@code invited_dt}.
     * <p>RN09-02: Email deve ser único entre usuários ativos do tenant.
     *
     * @param request  dados do usuário (nome, email)
     * @param tenantId tenant do contexto
     * @return o usuário criado como DTO de resposta
     * @throws DuplicateEmailException se o email já estiver em uso no tenant
     */
    @Transactional
    public UserResponse invite(UserCreateRequest request, UUID tenantId) {
        // RN09-02: email único por tenant ativo
        userRepo.findByEmailAndTenant(request.email(), tenantId).ifPresent(existing -> {
            log.warn("Convite recusado: email {} já existe no tenant {}", request.email(), tenantId);
            throw new DuplicateEmailException(request.email());
        });

        User user = new User();
        user.setTenantId(tenantId);
        user.setName(request.name());
        user.setEmail(request.email());
        user.setStatus(UserStatus.INVITE_PENDING);
        // RN09-01: convite expira em 7 dias
        user.setInvitedDt(OffsetDateTime.now(java.time.ZoneOffset.UTC));

        userRepo.save(user);
        log.info("Usuário convidado: id={}, email={}, tenant={}",
                user.getId(), maskEmail(user.getEmail()), tenantId);

        return UserResponse.from(user);
    }

    /**
     * Desativa um usuário (soft delete).
     *
     * <p>RN09-03: O admin não pode desativar a si mesmo.
     *
     * @param userId ID do usuário a desativar
     * @throws UserNotFoundException    se o usuário não existir
     * @throws SelfDeactivationException se tentar desativar a si mesmo
     */
    @Transactional
    public void deactivate(UUID userId) {
        // RN09-03: não permitir autodesativação
        UUID currentUserId = TenantContext.getUserIdQuietly();
        if (currentUserId != null && currentUserId.equals(userId)) {
            log.warn("Tentativa de autodesativação bloqueada: userId={}", userId);
            throw new SelfDeactivationException();
        }

        User user = userRepo.findById(userId)
                .orElseThrow(() -> new UserNotFoundException(userId));

        // Defesa em profundidade: verificar tenant do usuário vs contexto
        UUID tenantId = TenantContext.getTenantId();
        if (!tenantId.equals(user.getTenantId())) {
            log.warn("Tentativa de acesso cross-tenant bloqueada: userId={}, tenant={}", userId, tenantId);
            throw new UserNotFoundException(userId);
        }

        UUID deletedBy = TenantContext.getUserIdQuietly();
        userRepo.softDelete(userId, deletedBy != null ? deletedBy : userId);
        log.info("Usuário desativado: id={}, email={}", userId, maskEmail(user.getEmail()));
    }

    /**
     * Reativa um usuário previamente desativado.
     *
     * @param userId ID do usuário a reativar
     * @return o usuário reativado como DTO de resposta
     * @throws UserNotFoundException se o usuário não existir
     */
    @Transactional
    public UserResponse reactivate(UUID userId) {
        User user = userRepo.findById(userId)
                .orElseThrow(() -> new UserNotFoundException(userId));

        // Defesa em profundidade: verificar tenant do usuário vs contexto
        UUID tenantId = TenantContext.getTenantId();
        if (!tenantId.equals(user.getTenantId())) {
            log.warn("Tentativa de acesso cross-tenant bloqueada: userId={}, tenant={}", userId, tenantId);
            throw new UserNotFoundException(userId);
        }

        user.setDeletedDt(null);
        user.setDeletedBy(null);
        user.setStatus(UserStatus.ACTIVE);
        userRepo.update(user);

        log.info("Usuário reativado: id={}, email={}", userId, maskEmail(user.getEmail()));
        return UserResponse.from(user);
    }

    /**
     * Lista todos os usuários ativos do tenant.
     *
     * @param tenantId tenant do contexto
     * @return lista de usuários como DTOs de resposta
     */
    public List<UserResponse> findAll(UUID tenantId) {
        return userRepo.findAllByTenant(tenantId).stream()
                .map(UserResponse::from)
                .toList();
    }

    /**
     * Busca um usuário por ID.
     *
     * @param userId ID do usuário
     * @return o usuário como DTO de resposta
     * @throws UserNotFoundException se o usuário não existir
     */
    public UserResponse findById(UUID userId) {
        return userRepo.findById(userId)
                .map(UserResponse::from)
                .orElseThrow(() -> new UserNotFoundException(userId));
    }

    /** Mascara email para logs (LGPD). */
    private static String maskEmail(String email) {
        if (email == null || !email.contains("@")) return "***";
        int at = email.indexOf('@');
        return email.charAt(0) + "***" + email.substring(at);
    }
}
