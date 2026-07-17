package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.TenantCreateRequest;
import com.fbso.platform.admin.dto.request.TenantUpdateRequest;
import com.fbso.platform.admin.dto.response.TenantResponse;
import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.exception.DuplicateCnpjException;
import com.fbso.platform.admin.exception.InvalidStatusTransitionException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
import com.fbso.platform.admin.repository.TenantRepository;
import com.fbso.platform.admin.security.annotation.Auditable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Map;
import java.util.Set;
import java.util.UUID;

/**
 * Serviço de gestão de Tenants (F02-01, F02-02).
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN04-01: Status inicial PENDING_ONBOARDING, auditoria registrada</li>
 *   <li>RN04-02: Razão social única entre ativos</li>
 *   <li>RN05-01: Transições de status</li>
 *   <li>RN05-02: Suspensão exige motivo</li>
 *   <li>RN05-03: Reativação restaura permissões</li>
 * </ul>
 */
@Service
public class TenantService {

    private static final Logger log = LoggerFactory.getLogger(TenantService.class);

    /** Mapa de transições válidas (RN05-01) */
    static final Map<TenantStatus, Set<TenantStatus>> VALID_TRANSITIONS = Map.of(
            TenantStatus.PENDING_ONBOARDING, Set.of(TenantStatus.ACTIVE),
            TenantStatus.ACTIVE, Set.of(TenantStatus.SUSPENDED, TenantStatus.INACTIVE),
            TenantStatus.SUSPENDED, Set.of(TenantStatus.ACTIVE),
            TenantStatus.INACTIVE, Set.of(TenantStatus.ACTIVE)
    );

    private final TenantRepository tenantRepo;

    public TenantService(TenantRepository tenantRepo) {
        this.tenantRepo = tenantRepo;
    }

    // ---- Create (F02-01) ----

    @Auditable(entityType = "TENANT", action = "CREATED")
    @Transactional
    public TenantResponse create(TenantCreateRequest req) {
        // RN04-02: validação de razão social única
        if (tenantRepo.findByNameCorporate(req.nameCorporate()).isPresent()) {
            throw new DuplicateCnpjException(req.nameCorporate());
        }

        Tenant tenant = new Tenant();
        tenant.setNameCorporate(req.nameCorporate());
        tenant.setNameFantasy(req.nameFantasy());
        tenant.setSegment(req.segment());
        tenant.setStatus(TenantStatus.PENDING_ONBOARDING); // RN04-01

        tenantRepo.save(tenant);
        log.info("Tenant criado: id={}, nameCorporate={}", tenant.getId(), tenant.getNameCorporate());
        return TenantResponse.from(tenant);
    }

    // ---- Update (F02-01) ----

    @Auditable(entityType = "TENANT", action = "UPDATED")
    @Transactional
    public TenantResponse update(UUID id, TenantUpdateRequest req) {
        Tenant tenant = getById(id);

        if (req.nameFantasy() != null) {
            tenant.setNameFantasy(req.nameFantasy());
        }
        if (req.segment() != null) {
            tenant.setSegment(req.segment());
        }

        tenantRepo.update(tenant);
        log.info("Tenant atualizado: id={}", id);
        return TenantResponse.from(tenant);
    }

    // ---- Suspend (F02-02, RN05-02) ----

    @Auditable(entityType = "TENANT", action = "SUSPENDED")
    @Transactional
    public TenantResponse suspend(UUID id, String reason) {
        if (reason == null || reason.isBlank()) {
            throw new IllegalArgumentException("Motivo da suspensão é obrigatório");
        }

        Tenant tenant = getById(id);
        validateTransition(tenant.getStatus(), TenantStatus.SUSPENDED,
                tenant.getNameCorporate());

        tenant.setStatus(TenantStatus.SUSPENDED);
        tenantRepo.update(tenant);
        log.warn("Tenant suspenso: id={}, reason={}", id, reason);
        return TenantResponse.from(tenant);
    }

    // ---- Reactivate (F02-02, RN05-03) ----

    @Auditable(entityType = "TENANT", action = "REACTIVATED")
    @Transactional
    public TenantResponse reactivate(UUID id) {
        Tenant tenant = getById(id);
        validateTransition(tenant.getStatus(), TenantStatus.ACTIVE,
                tenant.getNameCorporate());

        tenant.setStatus(TenantStatus.ACTIVE);
        tenantRepo.update(tenant);
        log.info("Tenant reativado: id={}", id);
        return TenantResponse.from(tenant);
    }

    // ---- Query helpers ----

    public Tenant getById(UUID id) {
        return tenantRepo.findById(id)
                .orElseThrow(() -> new TenantNotFoundException(id));
    }

    /**
     * Lista paginada com filtros — delegada ao repository.
     * Centralizada no service para garantir que o {@code @RequiresPermission}
     * do controller seja aplicado uniformemente (sem bypass do repository).
     */
    public List<Tenant> findAllPaginated(int page, int size, String status,
                                          String plan, String search) {
        return tenantRepo.findAllPaginated(page, size, status, plan, search);
    }

    // ---- Transition validation (RN05-01) ----

    public static void validateTransition(TenantStatus from, TenantStatus to, String name) {
        Set<TenantStatus> allowed = VALID_TRANSITIONS.get(from);
        if (allowed == null || !allowed.contains(to)) {
            throw new InvalidStatusTransitionException(from.name(), to.name());
        }
    }
}
