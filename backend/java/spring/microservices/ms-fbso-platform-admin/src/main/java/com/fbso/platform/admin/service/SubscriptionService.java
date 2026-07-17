package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.response.SubscriptionResponse;
import com.fbso.platform.admin.entity.Plan;
import com.fbso.platform.admin.entity.Subscription;
import com.fbso.platform.admin.repository.PlanRepository;
import com.fbso.platform.admin.repository.SubscriptionRepository;
import com.fbso.platform.admin.security.annotation.Auditable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.UUID;

/**
 * Serviço de gestão de Assinaturas (F02-04).
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN07-01: 1 assinatura ativa por tenant</li>
 *   <li>RN07-02: Change-plan sem gap (transação atômica)</li>
 *   <li>DT-009: locked_price preserva preço original</li>
 * </ul>
 */
@Service
public class SubscriptionService {

    private static final Logger log = LoggerFactory.getLogger(SubscriptionService.class);

    private final SubscriptionRepository subRepo;
    private final PlanRepository planRepo;

    public SubscriptionService(SubscriptionRepository subRepo, PlanRepository planRepo) {
        this.subRepo = subRepo;
        this.planRepo = planRepo;
    }

    // ---- Create (RN07-01) ----

    @Auditable(entityType = "SUBSCRIPTION", action = "CREATED")
    @Transactional
    public SubscriptionResponse create(UUID tenantId, UUID planId) {
        // RN07-01: validar se já existe assinatura ativa
        subRepo.findActiveByTenantId(tenantId).ifPresent(existing -> {
            throw new IllegalStateException(
                    "Tenant já possui assinatura ativa: " + existing.getId());
        });

        Plan plan = planRepo.findById(planId)
                .orElseThrow(() -> new IllegalArgumentException("Plano não encontrado: " + planId));

        if (!plan.isActive()) {
            throw new com.fbso.platform.admin.exception.BusinessException(
                    "plan-inactive",
                    "Plano " + plan.getName() + " não está ativo (status=" + plan.getStatus() + ")");
        }

        Subscription sub = new Subscription();
        sub.setTenantId(tenantId);
        sub.setPlanId(planId);
        sub.setStatus("ACTIVE");
        sub.setStartDate(OffsetDateTime.now());
        // DT-009: travar preço e recorrência no momento da assinatura
        sub.setLockedPrice(plan.getPrice());
        sub.setLockedRecurrence(plan.getRecurrence() != null ? plan.getRecurrence().name() : null);

        subRepo.save(sub);
        log.info("Assinatura criada: id={}, tenantId={}, planId={}, lockedPrice={}",
                sub.getId(), tenantId, planId, sub.getLockedPrice());
        return SubscriptionResponse.from(sub);
    }

    // ---- Change Plan (RN07-02) ----

    @Auditable(entityType = "SUBSCRIPTION", action = "PLAN_CHANGED")
    @Transactional
    public SubscriptionResponse changePlan(UUID subscriptionId, UUID newPlanId) {
        Subscription current = getById(subscriptionId);

        if (!"ACTIVE".equals(current.getStatus())) {
            throw new IllegalStateException("Assinatura não está ativa: " + current.getStatus());
        }

        Plan newPlan = planRepo.findById(newPlanId)
                .orElseThrow(() -> new IllegalArgumentException("Plano não encontrado: " + newPlanId));

        if (!newPlan.isActive()) {
            throw new com.fbso.platform.admin.exception.BusinessException(
                    "plan-inactive",
                    "Plano " + newPlan.getName() + " não está ativo (status=" + newPlan.getStatus() + ")");
        }

        OffsetDateTime now = OffsetDateTime.now();

        // Finalizar assinatura atual (RN07-02: sem gap)
        current.setEndDate(now);
        current.setStatus("CANCELED");
        subRepo.update(current);

        // Criar nova assinatura (RN07-02: transição sem gap)
        Subscription newSub = new Subscription();
        newSub.setTenantId(current.getTenantId());
        newSub.setPlanId(newPlanId);
        newSub.setStatus("ACTIVE");
        newSub.setStartDate(now);
        newSub.setLockedPrice(newPlan.getPrice());
        newSub.setLockedRecurrence(newPlan.getRecurrence() != null ? newPlan.getRecurrence().name() : null);

        subRepo.save(newSub);
        log.info("Change-plan: subscriptionId={}, oldPlanId={}, newPlanId={}",
                subscriptionId, current.getPlanId(), newPlanId);
        return SubscriptionResponse.from(newSub);
    }

    // ---- Suspend ----

    @Auditable(entityType = "SUBSCRIPTION", action = "SUSPENDED")
    @Transactional
    public SubscriptionResponse suspend(UUID subscriptionId) {
        Subscription sub = getById(subscriptionId);
        sub.setStatus("SUSPENDED");
        subRepo.update(sub);
        log.warn("Assinatura suspensa: id={}", subscriptionId);
        return SubscriptionResponse.from(sub);
    }

    // ---- Query helpers ----

    public List<SubscriptionResponse> findByTenant(UUID tenantId) {
        return subRepo.findByTenantId(tenantId).stream()
                .map(SubscriptionResponse::from).toList();
    }

    private Subscription getById(UUID id) {
        return subRepo.findById(id)
                .orElseThrow(() -> new IllegalArgumentException("Assinatura não encontrada: " + id));
    }
}
