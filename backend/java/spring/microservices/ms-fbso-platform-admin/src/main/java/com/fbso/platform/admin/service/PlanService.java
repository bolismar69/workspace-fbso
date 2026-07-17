package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.PlanCreateRequest;
import com.fbso.platform.admin.dto.request.PlanUpdateRequest;
import com.fbso.platform.admin.dto.response.PlanResponse;
import com.fbso.platform.admin.entity.Plan;
import com.fbso.platform.admin.exception.PlanHasActiveSubscribersException;
import com.fbso.platform.admin.repository.PlanRepository;
import com.fbso.platform.admin.security.annotation.Auditable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.UUID;

/**
 * Serviço de gestão de Planos (F02-03).
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN06-01: Plano com assinantes ativos não pode ser desativado</li>
 *   <li>RN06-02: Alteração de preço não afeta assinaturas existentes (locked_price)</li>
 *   <li>RN06-03: Mínimo 1 plano ativo</li>
 * </ul>
 */
@Service
public class PlanService {

    private static final Logger log = LoggerFactory.getLogger(PlanService.class);

    private final PlanRepository planRepo;

    public PlanService(PlanRepository planRepo) {
        this.planRepo = planRepo;
    }

    // ---- Create ----

    @Auditable(entityType = "PLAN", action = "CREATED")
    @Transactional
    public PlanResponse create(PlanCreateRequest req) {
        Plan plan = new Plan();
        plan.setName(req.name());
        plan.setDescription(req.description());
        plan.setPrice(req.price());
        plan.setRecurrence(req.recurrence());
        plan.setVersion(1);

        planRepo.save(plan);
        log.info("Plano criado: id={}, name={}, price={}", plan.getId(), plan.getName(), plan.getPrice());
        return PlanResponse.from(plan);
    }

    // ---- Update (versionamento — RN06-02) ----

    @Auditable(entityType = "PLAN", action = "UPDATED")
    @Transactional
    public PlanResponse update(UUID id, PlanUpdateRequest req) {
        Plan existing = getById(id);

        // Se o preço mudou, versiona (RN06-02: assinaturas existentes mantêm locked_price)
        boolean priceChanged = req.price() != null
                && req.price().compareTo(existing.getPrice()) != 0;

        if (req.name() != null) existing.setName(req.name());
        if (req.description() != null) existing.setDescription(req.description());
        if (req.price() != null) existing.setPrice(req.price());
        if (req.recurrence() != null) existing.setRecurrence(req.recurrence());

        if (priceChanged) {
            existing.setVersion(existing.getVersion() + 1);
        }

        planRepo.update(existing);
        log.info("Plano atualizado: id={}, version={}", id, existing.getVersion());
        return PlanResponse.from(existing);
    }

    // ---- Deactivate (RN06-01, RN06-03) ----

    @Auditable(entityType = "PLAN", action = "DEACTIVATED")
    @Transactional
    public PlanResponse deactivate(UUID id) {
        Plan plan = getById(id);

        // RN06-01: não pode desativar plano com assinantes ativos
        if (planRepo.hasActiveSubscribers(id)) {
            throw new PlanHasActiveSubscribersException(plan.getName());
        }

        // RN06-03: mínimo 1 plano ativo
        if (planRepo.countActive() <= 1) {
            throw new com.fbso.platform.admin.exception.BusinessException(
                    "last-plan-active",
                    "Não é possível desativar o último plano ativo. Mínimo 1 plano deve estar ativo.");
        }

        plan.setStatus("DISCONTINUED");
        planRepo.update(plan);
        log.warn("Plano desativado: id={}, name={}", id, plan.getName());
        return PlanResponse.from(plan);
    }

    // ---- Queries ----

    public List<PlanResponse> listAll() {
        return planRepo.findAllActive().stream().map(PlanResponse::from).toList();
    }

    public List<PlanResponse> listAdmin() {
        return planRepo.findAll(0, 100, "name").stream().map(PlanResponse::from).toList();
    }

    public PlanResponse getPlan(UUID id) {
        return PlanResponse.from(getById(id));
    }

    Plan getById(UUID id) {
        return planRepo.findById(id)
                .orElseThrow(() -> new IllegalArgumentException("Plano não encontrado: " + id));
    }
}
