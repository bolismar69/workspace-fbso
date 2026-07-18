package com.fbso.platform.admin.security.aspect;

import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.PermissionService;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

/**
 * Aspecto RBAC — intercepta métodos anotados com {@link RequiresPermission}
 * e verifica se o papel do usuário tem acesso ao recurso + ação.
 * <p>
 * Pipeline: JWT Filter → RbacAspect → Controller
 * <p>
 * <b>Sprint 4 (M4 — RBAC):</b> Matriz carregada do banco via
 * {@link PermissionService} (tabelas {@code role_resource} +
 * {@code resource_action}). Roles do usuário via {@code user_permission}
 * (fonte primária), com fallback para JWT durante transição.
 * <p>
 * ADMIN_TENANT tem acesso implícito total — não requer registros em
 * {@code user_permission}. Demais roles são verificadas contra a matriz
 * carregada no startup e recarregável sob demanda.
 *
 * @see RequiresPermission
 * @see PermissionService
 */
@Aspect
@Component
public class RbacAspect {

    private static final Logger log = LoggerFactory.getLogger(RbacAspect.class);

    private final PermissionService permissionService;

    public RbacAspect(PermissionService permissionService) {
        this.permissionService = permissionService;
    }

    @Around("@annotation(requiresPermission)")
    public Object checkPermission(ProceedingJoinPoint joinPoint,
                                   RequiresPermission requiresPermission) throws Throwable {

        String resource = requiresPermission.resource();
        String action = requiresPermission.action();

        try {
            permissionService.checkPermission(resource, action);
        } catch (PermissionDeniedException e) {
            log.warn("RBAC: acesso negado — resource={}, action={}", resource, action);
            throw e;
        }

        log.debug("RBAC: acesso permitido — resource={}, action={}", resource, action);
        return joinPoint.proceed();
    }
}
